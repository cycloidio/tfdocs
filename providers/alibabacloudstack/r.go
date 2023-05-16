package alibabacloudstack

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_adb_account",
			Category:         "AnalyticDB for MySQL (ADB)",
			ShortDescription: `Provides a ADB account resource.`,
			Description:      ``,
			Keywords: []string{
				"analyticdb",
				"for",
				"mysql",
				"adb",
				"account",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "db_cluster_id",
					Description: `(Required, ForceNew) The Id of cluster in which account belongs.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `(Required, ForceNew) Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and have no more than 16 characters.`,
				},
				resource.Attribute{
					Name:        "account_password",
					Description: `(Optional) Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters. You have to specify one of ` + "`" + `account_password` + "`" + ` and ` + "`" + `kms_encrypted_password` + "`" + ` fields.`,
				},
				resource.Attribute{
					Name:        "kms_encrypted_password",
					Description: `(Optional) An KMS encrypts password used to a db account. If the ` + "`" + `account_password` + "`" + ` is filled in, this field will be ignored.`,
				},
				resource.Attribute{
					Name:        "kms_encryption_context",
					Description: `(Optional) An KMS encryption context used to decrypt ` + "`" + `kms_encrypted_password` + "`" + ` before creating or updating a db account with ` + "`" + `kms_encrypted_password` + "`" + `. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when ` + "`" + `kms_encrypted_password` + "`" + ` is set.`,
				},
				resource.Attribute{
					Name:        "account_description",
					Description: `(Optional) Account description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The current account resource ID. Composed of instance ID and account name with format ` + "`" + `<instance_id>:<name>` + "`" + `. ## Import ADB account can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_adb_account.example "am-12345:tf_account" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The current account resource ID. Composed of instance ID and account name with format ` + "`" + `<instance_id>:<name>` + "`" + `. ## Import ADB account can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_adb_account.example "am-12345:tf_account" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_adb_backup_policy",
			Category:         "AnalyticDB for MySQL (ADB)",
			ShortDescription: `Provides a ADB backup policy resource.`,
			Description:      ``,
			Keywords: []string{
				"analyticdb",
				"for",
				"mysql",
				"adb",
				"backup",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "db_cluster_id",
					Description: `(Required, ForceNew) The Id of cluster that can run database.`,
				},
				resource.Attribute{
					Name:        "preferred_backup_period",
					Description: `(Required) ADB Cluster backup period. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday].`,
				},
				resource.Attribute{
					Name:        "preferred_backup_time",
					Description: `(Required) ADB Cluster backup time, in the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. China time is 8 hours behind it. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The current backup policy resource ID. It is same as 'db_cluster_id'.`,
				},
				resource.Attribute{
					Name:        "backup_retention_period",
					Description: `Cluster backup retention days, Fixed for 7 days, not modified. ## Import ADB backup policy can be imported using the id or cluster id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_adb_backup_policy.example "am-12345678" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The current backup policy resource ID. It is same as 'db_cluster_id'.`,
				},
				resource.Attribute{
					Name:        "backup_retention_period",
					Description: `Cluster backup retention days, Fixed for 7 days, not modified. ## Import ADB backup policy can be imported using the id or cluster id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_adb_backup_policy.example "am-12345678" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_adb_connection",
			Category:         "AnalyticDB for MySQL (ADB)",
			ShortDescription: `Provides an ADB cluster connection resource.`,
			Description:      ``,
			Keywords: []string{
				"analyticdb",
				"for",
				"mysql",
				"adb",
				"connection",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "db_cluster_id",
					Description: `(Required, ForceNew) The Id of cluster that can run database.`,
				},
				resource.Attribute{
					Name:        "connection_prefix",
					Description: `(ForceNew) Prefix of the cluster public endpoint. The prefix must be 6 to 30 characters in length, and can contain lowercase letters, digits, and hyphens (-), must start with a letter and end with a digit or letter. Default to ` + "`" + `<db_cluster_id> + tf` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The current cluster connection resource ID. Composed of cluster ID and connection string with format ` + "`" + `<db_cluster_id>:<connection_prefix>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "connection_prefix",
					Description: `Prefix of a connection string.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Connection cluster port.`,
				},
				resource.Attribute{
					Name:        "connection_string",
					Description: `Connection cluster string.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The ip address of connection string. ## Import ADB connection can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_adb_connection.example am-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The current cluster connection resource ID. Composed of cluster ID and connection string with format ` + "`" + `<db_cluster_id>:<connection_prefix>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "connection_prefix",
					Description: `Prefix of a connection string.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Connection cluster port.`,
				},
				resource.Attribute{
					Name:        "connection_string",
					Description: `Connection cluster string.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The ip address of connection string. ## Import ADB connection can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_adb_connection.example am-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_adb_db_cluster",
			Category:         "AnalyticDB for MySQL (ADB)",
			ShortDescription: `Provides a Alibabacloudstack AnalyticDB for MySQL (ADB) DBCluster resource.`,
			Description:      ``,
			Keywords: []string{
				"analyticdb",
				"for",
				"mysql",
				"adb",
				"db",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "auto_renew_period",
					Description: `(Optional) Auto-renewal period of an cluster, in the unit of the month. It is valid when ` + "`" + `payment_type` + "`" + ` is ` + "`" + `Subscription` + "`" + `. Valid values: ` + "`" + `1` + "`" + `, ` + "`" + `2` + "`" + `, ` + "`" + `3` + "`" + `, ` + "`" + `6` + "`" + `, ` + "`" + `12` + "`" + `, ` + "`" + `24` + "`" + `, ` + "`" + `36` + "`" + `. Default to ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "compute_resource",
					Description: `(Optional) The specifications of computing resources in elastic mode. The increase of resources can speed up queries. AnalyticDB for MySQL automatically scales computing resources. For more information, see [ComputeResource](https://www.alibabacloud.com/help/en/doc-detail/144851.htm)`,
				},
				resource.Attribute{
					Name:        "db_cluster_category",
					Description: `(Required) The db cluster category. Valid values: ` + "`" + `Basic` + "`" + `, ` + "`" + `Cluster` + "`" + `, ` + "`" + `MixedStorage` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "db_cluster_class",
					Description: `(Deprecated) It duplicates with attribute db_node_class and is deprecated from 1.121.2.`,
				},
				resource.Attribute{
					Name:        "db_cluster_version",
					Description: `(Optional, ForceNew) The db cluster version. Value options: ` + "`" + `3.0` + "`" + `, Default to ` + "`" + `3.0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "db_node_class",
					Description: `(Optional, Computed) The db node class. For more information, see [DBClusterClass](https://help.aliyun.com/document_detail/190519.html)`,
				},
				resource.Attribute{
					Name:        "db_node_count",
					Description: `(Optional) The db node count.`,
				},
				resource.Attribute{
					Name:        "db_node_storage",
					Description: `(Optional) The db node storage.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, Computed) The description of DBCluster.`,
				},
				resource.Attribute{
					Name:        "maintain_time",
					Description: `(Optional, Computed) The maintenance window of the cluster. Format: hh:mmZ-hh:mmZ.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Required) The mode of the cluster. Valid values: ` + "`" + `reserver` + "`" + `, ` + "`" + `flexible` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "modify_type",
					Description: `(Optional) The modify type.`,
				},
				resource.Attribute{
					Name:        "pay_type",
					Description: `(Optional, Computed, ForceNew) Field ` + "`" + `pay_type` + "`" + ` has been deprecated. New field ` + "`" + `payment_type` + "`" + ` instead.`,
				},
				resource.Attribute{
					Name:        "payment_type",
					Description: `(Optional, Computed, ForceNew) The payment type of the resource. Valid values are ` + "`" + `PayAsYouGo` + "`" + ` and ` + "`" + `Subscription` + "`" + `. Default to ` + "`" + `PayAsYouGo` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) The duration that you will buy DB cluster (in month). It is valid when ` + "`" + `payment_type` + "`" + ` is ` + "`" + `Subscription` + "`" + `. Valid values: [1~9], 12, 24, 36. ->`,
				},
				resource.Attribute{
					Name:        "renewal_status",
					Description: `(Optional) Valid values are ` + "`" + `AutoRenewal` + "`" + `, ` + "`" + `Normal` + "`" + `, ` + "`" + `NotRenewal` + "`" + `, Default to ` + "`" + `NotRenewal` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional, Computed) The ID of the resource group.`,
				},
				resource.Attribute{
					Name:        "security_ips",
					Description: `(Optional, Computed) List of IP addresses allowed to access all databases of an cluster. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Optional, ForceNew) The vswitch id.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Optional, Computed, ForceNew) The zone ID of the resource.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `(Required) The cluster type of the resource. Valid values: ` + "`" + `analyticdb` + "`" + `, ` + "`" + `AnalyticdbOnPanguHybrid` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cpu_type",
					Description: `(Required) The cpu type of the resource.Valid values: ` + "`" + `intel` + "`" + `. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in terraform of DBCluster.`,
				},
				resource.Attribute{
					Name:        "connection_string",
					Description: `The endpoint of the cluster.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the resource. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 50 mins) Used when create the DBCluster.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 50 mins) Used when delete the DBCluster.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 72 mins) Used when update the DBCluster. ## Import AnalyticDB for MySQL (ADB) DBCluster can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_adb_db_cluster.example <id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in terraform of DBCluster.`,
				},
				resource.Attribute{
					Name:        "connection_string",
					Description: `The endpoint of the cluster.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the resource. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 50 mins) Used when create the DBCluster.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 50 mins) Used when delete the DBCluster.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 72 mins) Used when update the DBCluster. ## Import AnalyticDB for MySQL (ADB) DBCluster can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_adb_db_cluster.example <id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_alikafka_sasl_acl",
			Category:         "Alikafka",
			ShortDescription: `Provides a Alibabacloudstack Alikafka Sasl Acl resource.`,
			Description:      ``,
			Keywords: []string{
				"alikafka",
				"sasl",
				"acl",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, ForceNew) ID of the ALIKAFKA Instance that owns the groups.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required, ForceNew) Username for the sasl user. The length should between 1 to 64 characters. The user should be an existed sasl user.`,
				},
				resource.Attribute{
					Name:        "acl_resource_type",
					Description: `(Required, ForceNew) Resource type for this acl. The resource type can only be "Topic" and "Group".`,
				},
				resource.Attribute{
					Name:        "acl_resource_name",
					Description: `(Required, ForceNew) Resource name for this acl. The resource name should be a topic or consumer group name.`,
				},
				resource.Attribute{
					Name:        "acl_resource_pattern_type",
					Description: `(Required, ForceNew) Resource pattern type for this acl. The resource pattern support two types "LITERAL" and "PREFIXED". "LITERAL": A literal name defines the full name of a resource. The special wildcard character "`,
				},
				resource.Attribute{
					Name:        "acl_operation_type",
					Description: `(Required, ForceNew) Operation type for this acl. The operation type can only be "Write" and "Read". ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `key` + "`" + ` of the resource supplied above. The value is formulated as ` + "`" + `<instance_id>:<username>:<acl_resource_type>:<acl_resource_name>:<acl_resource_pattern_type>:<acl_operation_type>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `The host of the acl. ## Import ALIKAFKA GROUP can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_alikafka_sasl_acl.acl alikafka_post-cn-123455abc:username:Topic:test-topic:LITERAL:Write ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `key` + "`" + ` of the resource supplied above. The value is formulated as ` + "`" + `<instance_id>:<username>:<acl_resource_type>:<acl_resource_name>:<acl_resource_pattern_type>:<acl_operation_type>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `The host of the acl. ## Import ALIKAFKA GROUP can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_alikafka_sasl_acl.acl alikafka_post-cn-123455abc:username:Topic:test-topic:LITERAL:Write ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_alikafka_sasl_user",
			Category:         "Alikafka",
			ShortDescription: `Provides a Alibabacloudstack Alikafka Sasl User resource.`,
			Description:      ``,
			Keywords: []string{
				"alikafka",
				"sasl",
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, ForceNew) ID of the ALIKAFKA Instance that owns the groups.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required, ForceNew) Username for the sasl user. The length should between 1 to 64 characters. The characters can only contain 'a'-'z', 'A'-'Z', '0'-'9', '_' and '-'.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional, Sensitive) Operation password. It may consist of letters, digits, or underlines, with a length of 1 to 64 characters. You have to specify one of ` + "`" + `password` + "`" + ` and ` + "`" + `kms_encrypted_password` + "`" + ` fields.`,
				},
				resource.Attribute{
					Name:        "kms_encrypted_password",
					Description: `(Optional) An KMS encrypts password used to a db account. You have to specify one of ` + "`" + `password` + "`" + ` and ` + "`" + `kms_encrypted_password` + "`" + ` fields.`,
				},
				resource.Attribute{
					Name:        "kms_encryption_context",
					Description: `(Optional, MapString) An KMS encryption context used to decrypt ` + "`" + `kms_encrypted_password` + "`" + ` before creating or updating a user with ` + "`" + `kms_encrypted_password` + "`" + `. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when ` + "`" + `kms_encrypted_password` + "`" + ` is set.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional, ForceNew, Available in 1.159.0+) The authentication mechanism. Valid values: ` + "`" + `plain` + "`" + `, ` + "`" + `scram` + "`" + `. Default value: ` + "`" + `plain` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the resource. The value is formate as ` + "`" + `<instance_id>:<username>` + "`" + `. ## Import Alikafka Sasl User can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import alibabacloudstack_alikafka_sasl_user.example <instance_id>:<username> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the resource. The value is formate as ` + "`" + `<instance_id>:<username>` + "`" + `. ## Import Alikafka Sasl User can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import alibabacloudstack_alikafka_sasl_user.example <instance_id>:<username> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_alikafka_topic",
			Category:         "Alikafka",
			ShortDescription: `Provides a Alibabacloudstack ALIKAFKA Topic resource.`,
			Description:      ``,
			Keywords: []string{
				"alikafka",
				"topic",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, ForceNew) InstanceId of your Kafka resource, the topic will create in this instance.`,
				},
				resource.Attribute{
					Name:        "topic",
					Description: `(Required, ForceNew) Name of the topic. Two topics on a single instance cannot have the same name. The length cannot exceed 64 characters.`,
				},
				resource.Attribute{
					Name:        "local_topic",
					Description: `(Optional, ForceNew) Whether the topic is localTopic or not.`,
				},
				resource.Attribute{
					Name:        "compact_topic",
					Description: `(Optional, ForceNew) Whether the topic is compactTopic or not. Compact topic must be a localTopic.`,
				},
				resource.Attribute{
					Name:        "partition_num",
					Description: `(Optional) The number of partitions of the topic. The number should between 1 and 48.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `(Required) This attribute is a concise description of topic. The length cannot exceed 64.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Available in v1.63.0+) A mapping of tags to assign to the resource. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `key` + "`" + ` of the resource supplied above. The value is formulated as ` + "`" + `<instance_id>:<topic>` + "`" + `. ## Import ALIKAFKA TOPIC can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_alikafka_topic.topic alikafka_post-cn-123455abc:topicName ` + "`" + `` + "`" + `` + "`" + ` ### Timeouts ->`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 10 mins) Used when creating the topic (until it reaches the initial ` + "`" + `Running` + "`" + ` status).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `key` + "`" + ` of the resource supplied above. The value is formulated as ` + "`" + `<instance_id>:<topic>` + "`" + `. ## Import ALIKAFKA TOPIC can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_alikafka_topic.topic alikafka_post-cn-123455abc:topicName ` + "`" + `` + "`" + `` + "`" + ` ### Timeouts ->`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 10 mins) Used when creating the topic (until it reaches the initial ` + "`" + `Running` + "`" + ` status).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_api_gateway_api",
			Category:         "API Gateway",
			ShortDescription: `Provides a Alibabacloudstack Api Gateway Api Resource.`,
			Description:      ``,
			Keywords: []string{
				"api",
				"gateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the api gateway api. Defaults to null.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required, ForcesNew) The api gateway that the api belongs to. Defaults to null.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) The description of the api. Defaults to null.`,
				},
				resource.Attribute{
					Name:        "auth_type",
					Description: `(Required) The authorization Type including APP and ANONYMOUS. Defaults to null.`,
				},
				resource.Attribute{
					Name:        "request_config",
					Description: `(Required, Type: list) Request_config defines how users can send requests to your API.`,
				},
				resource.Attribute{
					Name:        "service_type",
					Description: `(Required) The type of backend service. Type including HTTP,VPC and MOCK. Defaults to null.`,
				},
				resource.Attribute{
					Name:        "fc_service_config",
					Description: `(Optional, Type: list) fc_service_config defines the config when service_type selected 'FunctionCompute'.`,
				},
				resource.Attribute{
					Name:        "request_parameters",
					Description: `(Required, Type: list) request_parameters defines the request parameters of the api. ### Block request_config The request_config mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The protocol of api which supports values of 'HTTP','HTTPS' or 'HTTP,HTTPS'.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Required) The method of the api, including 'GET','POST','PUT' etc.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The request path of the api.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Required) The mode of the parameters between request parameters and service parameters, which support the values of 'MAPPING' and 'PASSTHROUGH'. ### Block fc_vpc_service_config The fc_service_config mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The region that the function compute service belongs to.`,
				},
				resource.Attribute{
					Name:        "function_name",
					Description: `(Required) The function name of function compute service.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The service name of function compute service.`,
				},
				resource.Attribute{
					Name:        "arn_role",
					Description: `(Optional) RAM role arn attached to the Function Compute service. This governs both who / what can invoke your Function, as well as what resources our Function has access to. See [User Permissions](https://www.alibabacloud.com/help/doc-detail/52885.htm) for more details.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Required) Backend service time-out time; unit: millisecond. ### Block request_parameters The request_parameters mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Request's parameter name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Parameter type which supports values of 'STRING','INT','BOOLEAN','LONG',"FLOAT" and "DOUBLE".`,
				},
				resource.Attribute{
					Name:        "required",
					Description: `(Required) Parameter required or not; values: REQUIRED and OPTIONAL.`,
				},
				resource.Attribute{
					Name:        "in",
					Description: `(Required) Request's parameter location; values: BODY, HEAD, QUERY, and PATH.`,
				},
				resource.Attribute{
					Name:        "in_service",
					Description: `(Required) Backend service's parameter location; values: BODY, HEAD, QUERY, and PATH.`,
				},
				resource.Attribute{
					Name:        "name_service",
					Description: `(Required) Backend service's parameter name. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the api resource of api gateway.`,
				},
				resource.Attribute{
					Name:        "api_id",
					Description: `The ID of the api of api gateway. ## Import Api gateway api can be imported using the id.Format to ` + "`" + `<API Group Id>:<API Id>` + "`" + ` e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_api_gateway_api.example "ab2351f2ce904edaa8d92a0510832b91:e4f728fca5a94148b023b99a3e5d0b62" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the api resource of api gateway.`,
				},
				resource.Attribute{
					Name:        "api_id",
					Description: `The ID of the api of api gateway. ## Import Api gateway api can be imported using the id.Format to ` + "`" + `<API Group Id>:<API Id>` + "`" + ` e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_api_gateway_api.example "ab2351f2ce904edaa8d92a0510832b91:e4f728fca5a94148b023b99a3e5d0b62" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_api_gateway_app",
			Category:         "API Gateway",
			ShortDescription: `Provides a Alibabacloudstack Api Gateway App Resource.`,
			Description:      ``,
			Keywords: []string{
				"api",
				"gateway",
				"app",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the app.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the app. Defaults to null.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Available in v1.55.3+) A mapping of tags to assign to the resource. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the app of api gateway. ## Import Api gateway app can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_api_gateway_app.example "7379660" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the app of api gateway. ## Import Api gateway app can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_api_gateway_app.example "7379660" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_api_gateway_app_attachment",
			Category:         "API Gateway",
			ShortDescription: `Provides a Alibabacloudstack Api Gateway App Attachment Resource.`,
			Description:      ``,
			Keywords: []string{
				"api",
				"gateway",
				"app",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "app_id",
					Description: `(Required，ForceNew) The app that apply to the authorization.`,
				},
				resource.Attribute{
					Name:        "api_id",
					Description: `(Required，ForceNew) The api_id that app apply to access.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required，ForceNew) The group that the api belongs to.`,
				},
				resource.Attribute{
					Name:        "stage_name",
					Description: `(Required，ForceNew) Stage that the app apply to access. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the app attachment of api gateway., formatted as ` + "`" + `<group_id>:<api_id>:<app_id>:<stage_name>` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the app attachment of api gateway., formatted as ` + "`" + `<group_id>:<api_id>:<app_id>:<stage_name>` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_api_gateway_group",
			Category:         "API Gateway",
			ShortDescription: `Provides a Alibabacloudstack Api Gateway Group Resource.`,
			Description:      ``,
			Keywords: []string{
				"api",
				"gateway",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the api gateway group. Defaults to null.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) The description of the api gateway group. Defaults to null. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the api group of api gateway.`,
				},
				resource.Attribute{
					Name:        "sub_domain",
					Description: `(Available in 1.69.0+) Second-level domain name automatically assigned to the API group.`,
				},
				resource.Attribute{
					Name:        "vpc_domain",
					Description: `(Available in 1.69.0+) Second-level VPC domain name automatically assigned to the API group. ## Import Api gateway group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_api_gateway_group.example "ab2351f2ce904edaa8d92a0510832b91" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the api group of api gateway.`,
				},
				resource.Attribute{
					Name:        "sub_domain",
					Description: `(Available in 1.69.0+) Second-level domain name automatically assigned to the API group.`,
				},
				resource.Attribute{
					Name:        "vpc_domain",
					Description: `(Available in 1.69.0+) Second-level VPC domain name automatically assigned to the API group. ## Import Api gateway group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_api_gateway_group.example "ab2351f2ce904edaa8d92a0510832b91" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_api_gateway_vpc_access",
			Category:         "API Gateway",
			ShortDescription: `Provides a Alibabacloudstack Api Gateway vpc authorization Resource.`,
			Description:      ``,
			Keywords: []string{
				"api",
				"gateway",
				"vpc",
				"access",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required，ForceNew) The name of the vpc authorization.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required，ForceNew) The vpc id of the vpc authorization.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required，ForceNew) ID of the instance in VPC (ECS/Server Load Balance).`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required，ForceNew) ID of the port corresponding to the instance. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the vpc authorization of api gateway. ## Import Api gateway app can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_api_gateway_vpc_access.example "APiGatewayVpc:vpc-aswcj19ajsz:i-ajdjfsdlf:8080" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the vpc authorization of api gateway. ## Import Api gateway app can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_api_gateway_vpc_access.example "APiGatewayVpc:vpc-aswcj19ajsz:i-ajdjfsdlf:8080" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ascm_custom_role",
			Category:         "ASCM",
			ShortDescription: `Provides Ascm custom role.`,
			Description:      ``,
			Keywords: []string{
				"ascm",
				"custom",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "role_name",
					Description: `(Required) Custom Role name.`,
				},
				resource.Attribute{
					Name:        "organization_visibility",
					Description: `(Required) organization visibility. Valid Values are - "organizationVisibility.organization", "organizationVisibility.orgAndSubOrgs" and "organizationVisibility.global".`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for the custom role. Note - It should not contain any spaces.`,
				},
				resource.Attribute{
					Name:        "role_range",
					Description: `(Required) Role Range for the custom role.`,
				},
				resource.Attribute{
					Name:        "privileges",
					Description: `(Required) Privileges assign to that role. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Custom Role Name and ID of the user.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `The ID of the custom role.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Custom Role Name and ID of the user.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `The ID of the custom role.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ascm_logon_policy",
			Category:         "ASCM",
			ShortDescription: `Provides a Alibabacloudstack Logon Policy resource.`,
			Description:      ``,
			Keywords: []string{
				"ascm",
				"logon",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Logon Policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The Logon Policy description.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Optional) The rules of the logon policy. Valid values: Allow and Deny. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Logon Policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The Logon Policy description.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `The Rule for the Logon Policy.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `The ID of the logon policy created.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Logon Policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The Logon Policy description.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `The Rule for the Logon Policy.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `The ID of the logon policy created.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ascm_organization",
			Category:         "ASCM",
			ShortDescription: `Provides an Ascm organization resource.`,
			Description:      ``,
			Keywords: []string{
				"ascm",
				"organization",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the organization. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `(Optional) ID of its Parent organization. Default value for a parent_id is "1". For normal user (not an admin) parent_id will be its organization ID.`,
				},
				resource.Attribute{
					Name:        "person_num",
					Description: `(Optional) A reserved parameter.`,
				},
				resource.Attribute{
					Name:        "resource_group_num",
					Description: `(Optional) A reserved parameter. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Name and ID of the organization. The value is in format ` + "`" + `Name:ID` + "`" + ``,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `The ID of the organization.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Name and ID of the organization. The value is in format ` + "`" + `Name:ID` + "`" + ``,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `The ID of the organization.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ascm_password_policy",
			Category:         "ASCM",
			ShortDescription: `Provides a Ascm password policy configuration.`,
			Description:      ``,
			Keywords: []string{
				"ascm",
				"password",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "hard_expiry",
					Description: `(Optional) Specifies whether to disable logon after the password expires.`,
				},
				resource.Attribute{
					Name:        "require_numbers",
					Description: `(Optional) Specifies whether digits are required.`,
				},
				resource.Attribute{
					Name:        "require_symbols",
					Description: `(Optional) Specifies whether special characters are required.`,
				},
				resource.Attribute{
					Name:        "require_lowercase_characters",
					Description: `(Optional) Specifies whether lowercase letters are required.`,
				},
				resource.Attribute{
					Name:        "require_uppercase_characters",
					Description: `(Optional) Specifies whether uppercase letters are required.`,
				},
				resource.Attribute{
					Name:        "max_login_attempts",
					Description: `(Optional) The maximum number of allowed logon attempts`,
				},
				resource.Attribute{
					Name:        "max_password_age",
					Description: `(Optional) The validity period of the password.`,
				},
				resource.Attribute{
					Name:        "minimum_password_length",
					Description: `(Optional) The minimum length of the password.Valid value range: [8-32].`,
				},
				resource.Attribute{
					Name:        "password_reuse_prevention",
					Description: `(Optional) The maximum number of allowed password reuse attempts.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ascm_quota",
			Category:         "ASCM",
			ShortDescription: `Provides a Ascm quota resource.`,
			Description:      ``,
			Keywords: []string{
				"ascm",
				"quota",
			},
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
					Name:        "total_cpu",
					Description: `(Optional) This reserved parameter is optional and can be left empty.`,
				},
				resource.Attribute{
					Name:        "total_mem",
					Description: `(Optional) This reserved parameter is optional and can be left empty.`,
				},
				resource.Attribute{
					Name:        "total_gpu",
					Description: `(Optional) This reserved parameter is optional and can be left empty.`,
				},
				resource.Attribute{
					Name:        "total_disk_cloud_ssd",
					Description: `(Optional) This reserved parameter is optional and can be left empty.`,
				},
				resource.Attribute{
					Name:        "total_disk_cloud_efficiency",
					Description: `(Optional) This reserved parameter is optional and can be left empty.`,
				},
				resource.Attribute{
					Name:        "total_vip_internal",
					Description: `(Optional) This reserved parameter is optional and can be left empty.`,
				},
				resource.Attribute{
					Name:        "total_vip_public",
					Description: `(Optional) This reserved parameter is optional and can be left empty.`,
				},
				resource.Attribute{
					Name:        "total_vpc",
					Description: `(Optional) This reserved parameter is optional and can be left empty.`,
				},
				resource.Attribute{
					Name:        "total_amount",
					Description: `(Optional) This reserved parameter is optional and can be left empty.`,
				},
				resource.Attribute{
					Name:        "total_eip",
					Description: `(Optional) This reserved parameter is optional and can be left empty.`,
				},
				resource.Attribute{
					Name:        "total_disk",
					Description: `(Optional) This reserved parameter is optional and can be left empty.`,
				},
				resource.Attribute{
					Name:        "total_cu",
					Description: `(Optional) This reserved parameter is optional and can be left empty.`,
				},
				resource.Attribute{
					Name:        "target_type",
					Description: `(Optional) This reserved parameter is optional and can be left empty. It will be used only for some products. Products where target_type are required with their values - RDS ("MySql"), R-KVSTORE ("redis") and DDS ("mongodb"). You can call this operation to create a quota. Use parameters according to the product name. Sample for the product. ECS`,
				},
				resource.Attribute{
					Name:        "quota_id",
					Description: `ID of the quota.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ProductName, QuotaType and QuotaTypeId of the Service. The value is in format ` + "`" + `ProductName:QuotaType:QuotaTypeId` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "quota_id",
					Description: `ID of the quota.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ProductName, QuotaType and QuotaTypeId of the Service. The value is in format ` + "`" + `ProductName:QuotaType:QuotaTypeId` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ascm_ram_policy",
			Category:         "ASCM",
			ShortDescription: `Provides Ascm ram policy.`,
			Description:      ``,
			Keywords: []string{
				"ascm",
				"ram",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Ram Policy name.`,
				},
				resource.Attribute{
					Name:        "policy_document",
					Description: `(Required) Policy document of the policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for the ram policy. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Ram policy Name of the user.`,
				},
				resource.Attribute{
					Name:        "ram_id",
					Description: `The ID of the ram policy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Ram policy Name of the user.`,
				},
				resource.Attribute{
					Name:        "ram_id",
					Description: `The ID of the ram policy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ascm_ram_policy_for_role",
			Category:         "ASCM",
			ShortDescription: `Provides Ascm ram policy for role resource.`,
			Description:      ``,
			Keywords: []string{
				"ascm",
				"ram",
				"policy",
				"for",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ram_policy_id",
					Description: `(Required) ID of the ram_policy_id which will be used to bind.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `(Required, ForceNew) ID of the role which will be used to bind.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ascm_ram_role",
			Category:         "ASCM",
			ShortDescription: `Provides Ascm ram role.`,
			Description:      ``,
			Keywords: []string{
				"ascm",
				"ram",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "role_name",
					Description: `(Required) Role name.`,
				},
				resource.Attribute{
					Name:        "organization_visibility",
					Description: `(Required) organization visibility. Valid Values are - "organizationVisibility.organization", "organizationVisibility.orgAndSubOrgs" and "organizationVisibility.global".`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for the ram role. Note - It should not contain any spaces. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Ram Role Name of the user.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `The ID of the ram role.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Ram Role Name of the user.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `The ID of the ram role.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ascm_resource_group",
			Category:         "ASCM",
			ShortDescription: `Provides Ascm resource group resource.`,
			Description:      ``,
			Keywords: []string{
				"ascm",
				"resource",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the resource group. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.`,
				},
				resource.Attribute{
					Name:        "organization_id",
					Description: `(Required) ID of an Organization. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Name and ID of the resource group. The value is in format ` + "`" + `Name:ID` + "`" + ``,
				},
				resource.Attribute{
					Name:        "rg_id",
					Description: `The ID of the resource group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Name and ID of the resource group. The value is in format ` + "`" + `Name:ID` + "`" + ``,
				},
				resource.Attribute{
					Name:        "rg_id",
					Description: `The ID of the resource group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ascm_user",
			Category:         "ASCM",
			ShortDescription: `Provides a Ascm user resource.`,
			Description:      ``,
			Keywords: []string{
				"ascm",
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "login_name",
					Description: `(Required) User login name.`,
				},
				resource.Attribute{
					Name:        "cell_phone_number",
					Description: `(Required) Cellphone Number of a user.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) Display name of a user.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Required) Email ID of a user.`,
				},
				resource.Attribute{
					Name:        "mobile_nation_code",
					Description: `(Required) Mobile Nation Code of a user, where user belongs to.`,
				},
				resource.Attribute{
					Name:        "organization_id",
					Description: `(Required) User Organization ID.`,
				},
				resource.Attribute{
					Name:        "login_policy_id",
					Description: `(Optional) User login policy ID. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Login Name of the user.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `The ID of the user.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Login Name of the user.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `The ID of the user.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ascm_user_group",
			Category:         "ASCM",
			ShortDescription: `Provides a Ascm user group resource.`,
			Description:      ``,
			Keywords: []string{
				"ascm",
				"user",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group_name",
					Description: `(Required) group name.`,
				},
				resource.Attribute{
					Name:        "organization_id",
					Description: `(Required) User Organization ID.`,
				},
				resource.Attribute{
					Name:        "role_in_ids",
					Description: `(Optional) ascm role id. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Login Name of the user group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Login Name of the user group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ascm_user_group_resource_set_binding",
			Category:         "ASCM",
			ShortDescription: `Provides Ascm User Role Binding.`,
			Description:      ``,
			Keywords: []string{
				"ascm",
				"user",
				"group",
				"resource",
				"set",
				"binding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_set_id",
					Description: `(Required) List of resource group id.`,
				},
				resource.Attribute{
					Name:        "user_group_id",
					Description: `(Required) user group id.`,
				},
				resource.Attribute{
					Name:        "ascm_role_id",
					Description: `(Optional) ascm role id. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "resource_set_id",
					Description: `(Required) List of resource group id.`,
				},
				resource.Attribute{
					Name:        "user_group_id",
					Description: `(Required) user group id.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_set_id",
					Description: `(Required) List of resource group id.`,
				},
				resource.Attribute{
					Name:        "user_group_id",
					Description: `(Required) user group id.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ascm_user_group_role_binding",
			Category:         "ASCM",
			ShortDescription: `Provides Ascm User Role Binding.`,
			Description:      ``,
			Keywords: []string{
				"ascm",
				"user",
				"group",
				"role",
				"binding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "user_group_id",
					Description: `(Required) ID of user group.`,
				},
				resource.Attribute{
					Name:        "role_ids",
					Description: `(Required) User Role Id. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "user_group_id",
					Description: `(Required) ID of user group.`,
				},
				resource.Attribute{
					Name:        "role_ids",
					Description: `(Required) User Role Id.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "user_group_id",
					Description: `(Required) ID of user group.`,
				},
				resource.Attribute{
					Name:        "role_ids",
					Description: `(Required) User Role Id.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ascm_user_role_binding",
			Category:         "ASCM",
			ShortDescription: `Provides Ascm User Role Binding.`,
			Description:      ``,
			Keywords: []string{
				"ascm",
				"user",
				"role",
				"binding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "role_id",
					Description: `(Required) ID of the role which will be used to bind with user.`,
				},
				resource.Attribute{
					Name:        "login_name",
					Description: `(Required) Name of the User. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Name of the User.`,
				},
				resource.Attribute{
					Name:        "login_name",
					Description: `Name of the User.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `User Role Id.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Name of the User.`,
				},
				resource.Attribute{
					Name:        "login_name",
					Description: `Name of the User.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `User Role Id.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ascm_usergroup_user",
			Category:         "ASCM",
			ShortDescription: `Provides a Ascm usergroup_user resource.`,
			Description:      ``,
			Keywords: []string{
				"ascm",
				"usergroup",
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "user_group_id",
					Description: `(Required) group name.`,
				},
				resource.Attribute{
					Name:        "login_names",
					Description: `(Required) List of user login name. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Login Name of the usergroup_user.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Login Name of the usergroup_user.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_cms_alarm",
			Category:         "Cloud Monitor",
			ShortDescription: `Provides a resource to build a alarm rule for cloud monitor.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"monitor",
				"cms",
				"alarm",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The alarm rule name.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Required, ForceNew) Monitor project name, such as "acs_ecs_dashboard" and "acs_rds_dashboard". For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).`,
				},
				resource.Attribute{
					Name:        "metric",
					Description: `(Required, ForceNew) Name of the monitoring metrics corresponding to a project, such as "CPUUtilization" and "networkin_rate". For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).`,
				},
				resource.Attribute{
					Name:        "dimensions",
					Description: `(Required, ForceNew) Map of the resources associated with the alarm rule, such as "instanceId", "device" and "port". Each key's value is a string and it uses comma to split multiple items. For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `Index query cycle, which must be consistent with that defined for metrics.`,
				},
				resource.Attribute{
					Name:        "escalations_critical",
					Description: `(Optional) A configuration of critical alarm (documented below).`,
				},
				resource.Attribute{
					Name:        "escalations_warn",
					Description: `(Optional) A configuration of critical warn (documented below).`,
				},
				resource.Attribute{
					Name:        "escalations_info",
					Description: `(Optional) A configuration of critical info (documented below).`,
				},
				resource.Attribute{
					Name:        "contact_groups",
					Description: `(Required) List contact groups of the alarm rule, which must have been created on the console.`,
				},
				resource.Attribute{
					Name:        "effective_interval",
					Description: `(Available) The interval of effecting alarm rule. It foramt as "hh:mm-hh:mm", like "0:00-4:00". Default to "00:00-23:59".`,
				},
				resource.Attribute{
					Name:        "silence_time",
					Description: `Notification silence period in the alarm state, in seconds. Valid value range: [300, 86400]. Default to 86400`,
				},
				resource.Attribute{
					Name:        "notify_type",
					Description: `Notification type. Valid value [0, 1]. The value 0 indicates TradeManager+email, and the value 1 indicates that TradeManager+email+SMS`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Whether to enable alarm rule. Default to true.`,
				},
				resource.Attribute{
					Name:        "statistics",
					Description: `Critical level alarm statistics method.. It must be consistent with that defined for metrics. Valid values: ["Average", "Minimum", "Maximum"]. Default to "Average".`,
				},
				resource.Attribute{
					Name:        "comparison_operator",
					Description: `Critical level alarm comparison operator. Valid values: ["<=", "<", ">", ">=", "==", "!="]. Default to "==".`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `Critical level alarm threshold value, which must be a numeric value currently.`,
				},
				resource.Attribute{
					Name:        "times",
					Description: `Critical level alarm retry times. Default to 3. #### Block escalations warn alarm The escalations_warn supports the following:`,
				},
				resource.Attribute{
					Name:        "statistics",
					Description: `Critical level alarm statistics method.. It must be consistent with that defined for metrics. Valid values: ["Average", "Minimum", "Maximum"]. Default to "Average".`,
				},
				resource.Attribute{
					Name:        "comparison_operator",
					Description: `Critical level alarm comparison operator. Valid values: ["<=", "<", ">", ">=", "==", "!="]. Default to "==".`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `Critical level alarm threshold value, which must be a numeric value currently.`,
				},
				resource.Attribute{
					Name:        "times",
					Description: `Critical level alarm retry times. Default to 3. #### Block escalations info alarm The escalations_info supports the following:`,
				},
				resource.Attribute{
					Name:        "statistics",
					Description: `Critical level alarm statistics method.. It must be consistent with that defined for metrics. Valid values: ["Average", "Minimum", "Maximum"]. Default to "Average".`,
				},
				resource.Attribute{
					Name:        "comparison_operator",
					Description: `Critical level alarm comparison operator. Valid values: ["<=", "<", ">", ">=", "==", "!="]. Default to "==".`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `Critical level alarm threshold value, which must be a numeric value currently.`,
				},
				resource.Attribute{
					Name:        "times",
					Description: `Critical level alarm retry times. Default to 3. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the alarm rule.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current alarm rule status. ## Import Alarm rule can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_cms_alarm.alarm abc12345 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the alarm rule.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current alarm rule status. ## Import Alarm rule can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_cms_alarm.alarm abc12345 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_cms_alarm_contact",
			Category:         "Cloud Monitor",
			ShortDescription: `Provides a resource to add a alarm contact for cloud monitor.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"monitor",
				"cms",
				"alarm",
				"contact",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "alarm_contact_name",
					Description: `(Required, ForceNew) The name of the alarm contact.`,
				},
				resource.Attribute{
					Name:        "channels_aliim",
					Description: `(Optional) The TradeManager ID of the alarm contact.`,
				},
				resource.Attribute{
					Name:        "channels_ding_web_hook",
					Description: `(Optional) The webhook URL of the DingTalk chatbot.`,
				},
				resource.Attribute{
					Name:        "channels_mail",
					Description: `(Optional) The email address of the alarm contact. After you add or modify an email address, the recipient receives an email that contains an activation link. The system adds the recipient to the list of alarm contacts only after the recipient activates the email address.`,
				},
				resource.Attribute{
					Name:        "channels_sms",
					Description: `(Optional) The phone number of the alarm contact. After you add or modify an email address, the recipient receives an email that contains an activation link. The system adds the recipient to the list of alarm contacts only after the recipient activates the email address.`,
				},
				resource.Attribute{
					Name:        "describe",
					Description: `(Required) The description of the alarm contact.`,
				},
				resource.Attribute{
					Name:        "lang",
					Description: `(Optional) The language type of the alarm. Valid values: ` + "`" + `en` + "`" + `, ` + "`" + `zh-cn` + "`" + `. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the alarm contact. The value is same with ` + "`" + `alarm_contact_name` + "`" + `. ## Import Alarm contact can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_cms_alarm_contact.example abc12345 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the alarm contact. The value is same with ` + "`" + `alarm_contact_name` + "`" + `. ## Import Alarm contact can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_cms_alarm_contact.example abc12345 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_cms_alarm_contact_group",
			Category:         "Cloud Monitor Service",
			ShortDescription: `Provides a Alibabacloudstack CMS Alarm Contact Group resource.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"monitor",
				"service",
				"cms",
				"alarm",
				"contact",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "alarm_contact_group_name",
					Description: `(Required, ForceNew) The name of the alarm group.`,
				},
				resource.Attribute{
					Name:        "contacts",
					Description: `(Optional) The name of the alert contact.`,
				},
				resource.Attribute{
					Name:        "describe",
					Description: `(Optional) The description of the alert group.`,
				},
				resource.Attribute{
					Name:        "enable_subscribed",
					Description: `(Optional) Whether to open weekly subscription. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID of Alarm Contact Group. ## Import CMS Alarm Contact Group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_cms_alarm_contact_group.example tf-testacc123 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID of Alarm Contact Group. ## Import CMS Alarm Contact Group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_cms_alarm_contact_group.example tf-testacc123 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_cms_site_monitor",
			Category:         "Cloud Monitor Service",
			ShortDescription: `Provides a resource to build a site monitor rule for cloud monitor.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"monitor",
				"service",
				"cms",
				"site",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "address",
					Description: `(Required) The URL or IP address monitored by the site monitoring task.`,
				},
				resource.Attribute{
					Name:        "task_name",
					Description: `(Required) The name of the site monitoring task. The name must be 4 to 100 characters in length. The name can contain the following types of characters: letters, digits, and underscores.`,
				},
				resource.Attribute{
					Name:        "task_type",
					Description: `(Required, ForceNew) The protocol of the site monitoring task. Currently, site monitoring supports the following protocols: HTTP, Ping, TCP, UDP, DNS, SMTP, POP3, and FTP.`,
				},
				resource.Attribute{
					Name:        "alert_ids",
					Description: `The IDs of existing alert rules to be associated with the site monitoring task.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `The monitoring interval of the site monitoring task. Unit: minutes. Valid values: 1, 5, and 15. Default value: 1.`,
				},
				resource.Attribute{
					Name:        "isp_cities",
					Description: `The detection points in a JSON array. For example, ` + "`" + `[{"city":"546","isp":"465"},{"city":"572","isp":"465"},{"city":"738","isp":"465"}]` + "`" + ` indicates the detection points in Beijing, Hangzhou, and Qingdao respectively. You can call the [DescribeSiteMonitorISPCityList](https://www.alibabacloud.com/help/en/doc-detail/115045.htm) operation to query detection point information. If this parameter is not specified, three detection points will be chosen randomly for monitoring.`,
				},
				resource.Attribute{
					Name:        "options_json",
					Description: `The extended options of the protocol of the site monitoring task. The options vary according to the protocol. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the site monitor rule. ## Import Alarm rule can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_cms_site_monitor.alarm abc12345 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the site monitor rule. ## Import Alarm rule can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_cms_site_monitor.alarm abc12345 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_common_bandwidth_package",
			Category:         "VPC",
			ShortDescription: `Provides a Alibabacloudstack Common Bandwidth Package resource.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"common",
				"bandwidth",
				"package",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Required) The bandwidth of the common bandwidth package, in Mbps.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the common bandwidth package.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the common bandwidth package instance. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the common bandwidth package instance id.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the common bandwidth package instance id.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_common_bandwidth_package_attachment",
			Category:         "VPC",
			ShortDescription: `Provides an Alibabacloudstack Common Attachment resource.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"common",
				"bandwidth",
				"package",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bandwidth_package_id",
					Description: `(Required, ForceNew) The bandwidth_package_id of the common bandwidth package attachment, the field can't be changed.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, ForceNew) The instance_id of the common bandwidth package attachment, the field can't be changed. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the common bandwidth package attachment id and formates as ` + "`" + `<bandwidth_package_id>:<instance_id>` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the common bandwidth package attachment id and formates as ` + "`" + `<bandwidth_package_id>:<instance_id>` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_cr_ee_namespace",
			Category:         "Container Registry (CR)",
			ShortDescription: `Provides a Alibabacloudstack resource to manage Container Registry Enterprise Edition namespaces.`,
			Description:      ``,
			Keywords: []string{
				"container",
				"registry",
				"cr",
				"ee",
				"namespace",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, ForceNew) ID of Container Registry Enterprise Edition instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Name of Container Registry Enterprise Edition namespace. It can contain 2 to 30 characters.`,
				},
				resource.Attribute{
					Name:        "auto_create",
					Description: `(Required) Boolean, when it set to true, repositories are automatically created when pushing new images. If it set to false, you create repository for images before pushing.`,
				},
				resource.Attribute{
					Name:        "default_visibility",
					Description: `(Required) ` + "`" + `PUBLIC` + "`" + ` or ` + "`" + `PRIVATE` + "`" + `, default repository visibility in this namespace. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of Container Registry Enterprise Edition namespace. The value is in format ` + "`" + `{instance_id}:{namespace}` + "`" + ` . ## Import Container Registry Enterprise Edition namespace can be imported using the ` + "`" + `{instance_id}:{namespace}` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_cr_ee_namespace.default cri-xxx:my-namespace ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of Container Registry Enterprise Edition namespace. The value is in format ` + "`" + `{instance_id}:{namespace}` + "`" + ` . ## Import Container Registry Enterprise Edition namespace can be imported using the ` + "`" + `{instance_id}:{namespace}` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_cr_ee_namespace.default cri-xxx:my-namespace ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_cr_ee_repo",
			Category:         "Container Registry (CR)",
			ShortDescription: `Provides a Alibabacloudstack resource to manage Container Registry Enterprise Edition repositories.`,
			Description:      ``,
			Keywords: []string{
				"container",
				"registry",
				"cr",
				"ee",
				"repo",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, ForceNew) ID of Container Registry Enterprise Edition instance.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Required, ForceNew) Name of Container Registry Enterprise Edition namespace where repository is located. It can contain 2 to 30 characters.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Name of Container Registry Enterprise Edition repository. It can contain 2 to 64 characters.`,
				},
				resource.Attribute{
					Name:        "summary",
					Description: `(Required) The repository general information. It can contain 1 to 100 characters.`,
				},
				resource.Attribute{
					Name:        "repo_type",
					Description: `(Required) ` + "`" + `PUBLIC` + "`" + ` or ` + "`" + `PRIVATE` + "`" + `, repo's visibility.`,
				},
				resource.Attribute{
					Name:        "detail",
					Description: `(Optional) The repository specific information. MarkDown format is supported, and the length limit is 2000. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource id of Container Registry Enterprise Edition repository. The value is in format ` + "`" + `{instance_id}:{namespace}:{repository}` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "repo_id",
					Description: `The uuid of Container Registry Enterprise Edition repository. ## Import Container Registry Enterprise Edition repository can be imported using the ` + "`" + `{instance_id}:{namespace}:{repository}` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_cr_ee_repo.default ` + "`" + `cri-xxx:my-namespace:my-repo` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource id of Container Registry Enterprise Edition repository. The value is in format ` + "`" + `{instance_id}:{namespace}:{repository}` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "repo_id",
					Description: `The uuid of Container Registry Enterprise Edition repository. ## Import Container Registry Enterprise Edition repository can be imported using the ` + "`" + `{instance_id}:{namespace}:{repository}` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_cr_ee_repo.default ` + "`" + `cri-xxx:my-namespace:my-repo` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_cr_ee_sync_rule",
			Category:         "Container Registry (CR)",
			ShortDescription: `Provides a Alibabacloudstack resource to manage Container Registry Enterprise Edition sync rules.`,
			Description:      ``,
			Keywords: []string{
				"container",
				"registry",
				"cr",
				"ee",
				"sync",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, ForceNew) ID of Container Registry Enterprise Edition source instance.`,
				},
				resource.Attribute{
					Name:        "namespace_name",
					Description: `(Required, ForceNew) Name of Container Registry Enterprise Edition source namespace. It can contain 2 to 30 characters.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Name of Container Registry Enterprise Edition sync rule.`,
				},
				resource.Attribute{
					Name:        "target_region_id",
					Description: `(Required, ForceNew) The target region to be synchronized.`,
				},
				resource.Attribute{
					Name:        "target_instance_id",
					Description: `(Required, ForceNew) ID of Container Registry Enterprise Edition target instance to be synchronized.`,
				},
				resource.Attribute{
					Name:        "target_namespace_name",
					Description: `(Required, ForceNew) Name of Container Registry Enterprise Edition target namespace to be synchronized. It can contain 2 to 30 characters.`,
				},
				resource.Attribute{
					Name:        "tag_filter",
					Description: `(Required, ForceNew) The regular expression used to filter image tags for synchronization in the source repository.`,
				},
				resource.Attribute{
					Name:        "repo_name",
					Description: `(Optional, ForceNew) Name of the source repository which should be set together with ` + "`" + `target_repo_name` + "`" + `, if empty means that the synchronization scope is the entire namespace level.`,
				},
				resource.Attribute{
					Name:        "target_repo_name",
					Description: `(Optional, ForceNew) Name of the target repository. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource id of Container Registry Enterprise Edition sync rule. The value is in format ` + "`" + `{instance_id}:{namespace_name}:{rule_id}` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `The uuid of Container Registry Enterprise Edition sync rule.`,
				},
				resource.Attribute{
					Name:        "sync_direction",
					Description: `` + "`" + `FROM` + "`" + ` or ` + "`" + `TO` + "`" + `, the direction of synchronization. ` + "`" + `FROM` + "`" + ` means source instance, ` + "`" + `TO` + "`" + ` means target instance.`,
				},
				resource.Attribute{
					Name:        "sync_scope",
					Description: `` + "`" + `REPO` + "`" + ` or ` + "`" + `NAMESPACE` + "`" + `,the scope that the synchronization rule applies. ## Import Container Registry Enterprise Edition sync rule can be imported using the id. Format to ` + "`" + `{instance_id}:{namespace_name}:{rule_id}` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_cr_ee_sync_rule.default ` + "`" + `cri-xxx:my-namespace:crsr-yyy` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource id of Container Registry Enterprise Edition sync rule. The value is in format ` + "`" + `{instance_id}:{namespace_name}:{rule_id}` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `The uuid of Container Registry Enterprise Edition sync rule.`,
				},
				resource.Attribute{
					Name:        "sync_direction",
					Description: `` + "`" + `FROM` + "`" + ` or ` + "`" + `TO` + "`" + `, the direction of synchronization. ` + "`" + `FROM` + "`" + ` means source instance, ` + "`" + `TO` + "`" + ` means target instance.`,
				},
				resource.Attribute{
					Name:        "sync_scope",
					Description: `` + "`" + `REPO` + "`" + ` or ` + "`" + `NAMESPACE` + "`" + `,the scope that the synchronization rule applies. ## Import Container Registry Enterprise Edition sync rule can be imported using the id. Format to ` + "`" + `{instance_id}:{namespace_name}:{rule_id}` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_cr_ee_sync_rule.default ` + "`" + `cri-xxx:my-namespace:crsr-yyy` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_cr_namespace",
			Category:         "Container Registry (CR)",
			ShortDescription: `Provides a Alibabacloudstack resource to manage Container Registry namespaces.`,
			Description:      ``,
			Keywords: []string{
				"container",
				"registry",
				"cr",
				"namespace",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Name of Container Registry namespace.`,
				},
				resource.Attribute{
					Name:        "auto_create",
					Description: `(Required) Boolean, when it set to true, repositories are automatically created when pushing new images. If it set to false, you create repository for images before pushing.`,
				},
				resource.Attribute{
					Name:        "default_visibility",
					Description: `(Required) ` + "`" + `PUBLIC` + "`" + ` or ` + "`" + `PRIVATE` + "`" + `, default repository visibility in this namespace. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of Container Registry namespace. The value is same as its name.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of Container Registry namespace. The value is same as its name.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_cr_repo",
			Category:         "Container Registry (CR)",
			ShortDescription: `Provides a Alibabacloudstack resource to manage Container Registry repositories.`,
			Description:      ``,
			Keywords: []string{
				"container",
				"registry",
				"cr",
				"repo",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Required, ForceNew) Name of container registry namespace where repository is located.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Name of container registry repository.`,
				},
				resource.Attribute{
					Name:        "summary",
					Description: `(Required) The repository general information. It can contain 1 to 80 characters.`,
				},
				resource.Attribute{
					Name:        "repo_type",
					Description: `(Required) ` + "`" + `PUBLIC` + "`" + ` or ` + "`" + `PRIVATE` + "`" + `, repo's visibility.`,
				},
				resource.Attribute{
					Name:        "detail",
					Description: `(Optional) The repository specific information. MarkDown format is supported, and the length limit is 2000. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of Container Registry repository. The value is in format ` + "`" + `namespace/repository` + "`" + `.`,
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
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of Container Registry repository. The value is in format ` + "`" + `namespace/repository` + "`" + `.`,
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
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_cs_kubernetes",
			Category:         "Container Service for Kubernetes (ACK)",
			ShortDescription: `Provides a Alibabacloudstack resource to manage container kubernetes cluster.`,
			Description:      ``,
			Keywords: []string{
				"container",
				"service",
				"for",
				"kubernetes",
				"ack",
				"cs",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The kubernetes cluster's name. It is unique in one Alibabacloudstack account.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) Desired Kubernetes version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except you set a higher version number. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by ACK.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required, Sensitive) The password of ssh login cluster node. You have to specify one of ` + "`" + `password` + "`" + ` ` + "`" + `key_name` + "`" + ` ` + "`" + `kms_encrypted_password` + "`" + ` fields.`,
				},
				resource.Attribute{
					Name:        "kms_encrypted_password",
					Description: `(Required) An KMS encrypts password used to a cs kubernetes. You have to specify one of ` + "`" + `password` + "`" + ` ` + "`" + `key_name` + "`" + ` ` + "`" + `kms_encrypted_password` + "`" + ` fields.`,
				},
				resource.Attribute{
					Name:        "enable_ssh",
					Description: `(Optional) Enable login to the node through SSH. default: false`,
				},
				resource.Attribute{
					Name:        "cpu_policy",
					Description: `kubelet cpu policy. options: static|none. default: none.`,
				},
				resource.Attribute{
					Name:        "proxy_mode",
					Description: `Proxy mode is option of kube-proxy. options: iptables|ipvs. default: ipvs.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional) Windows instances support batch and PowerShell scripts. If your script file is larger than 1 KB, we recommend that you upload the script to Object Storage Service (OSS) and pull it through the internal endpoint of your OSS bucket.`,
				},
				resource.Attribute{
					Name:        "platform",
					Description: `(Optional) The architecture of the nodes that run pods. Default to CentOS.`,
				},
				resource.Attribute{
					Name:        "pod_cidr",
					Description: `(Required) [Flannel Specific] The CIDR block for the pod network when using Flannel.`,
				},
				resource.Attribute{
					Name:        "pod_vswitch_ids",
					Description: `(Required) [Terway Specific] The vswitches for the pod network when using Terway.Be careful the ` + "`" + `pod_vswitch_ids` + "`" + ` can not equal to ` + "`" + `worker_vswtich_ids` + "`" + ` or ` + "`" + `master_vswtich_ids` + "`" + ` but must be in same availability zones.`,
				},
				resource.Attribute{
					Name:        "new_nat_gateway",
					Description: `(Optional) Whether to create a new nat gateway while creating kubernetes cluster. Default to true. Then openapi in Alibabacloudstack are not all on intranet, So turn this option on is a good choice.`,
				},
				resource.Attribute{
					Name:        "service_cidr",
					Description: `(Optional) The CIDR block for the service network. It cannot be duplicated with the VPC CIDR and CIDR used by Kubernetes cluster in VPC, cannot be modified after creation.`,
				},
				resource.Attribute{
					Name:        "node_cidr_mask",
					Description: `(Optional) The node cidr block to specific how many pods can run on single node. 24-28 is allowed. 24 means 2^(32-24)-1=255 and the node can run at most 255 pods. default: 24`,
				},
				resource.Attribute{
					Name:        "slb_internet_enabled",
					Description: `(Optional) Whether to create internet load balancer for API Server. Default to true. If you want to use ` + "`" + `Terway` + "`" + ` as CNI network plugin, You need to specify the ` + "`" + `pod_vswitch_ids` + "`" + ` field and addons with ` + "`" + `csi-plugin` + "`" + `,` + "`" + `csi-provisioner` + "`" + `,` + "`" + `logtail-ds` + "`" + ` and ` + "`" + `nginx-ingress-controller` + "`" + `. If you want to use ` + "`" + `Flannel` + "`" + ` as CNI network plugin, You need to specify the ` + "`" + `pod_cidr` + "`" + ` field and addons with ` + "`" + `flannel` + "`" + `. #### Master params`,
				},
				resource.Attribute{
					Name:        "master_disk_category",
					Description: `(Optional) The system disk category of master node. Its valid value are ` + "`" + `cloud_ssd` + "`" + ` and ` + "`" + `cloud_efficiency` + "`" + `. Default to ` + "`" + `cloud_efficiency` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "master_disk_size",
					Description: `(Optional) The system disk size of master node. Its valid value range [20~500] in GB. Default to 20.`,
				},
				resource.Attribute{
					Name:        "master_vswtich_ids",
					Description: `(Required) The vswitches used by master, you can specific 3 or 5 vswitches because of the amount of masters. Detailed below.`,
				},
				resource.Attribute{
					Name:        "master_instance_types",
					Description: `(Required) The instance type of master node. Specify one type for single AZ Cluster, three types for MultiAZ Cluster.`,
				},
				resource.Attribute{
					Name:        "master_system_disk_performance_level",
					Description: `(Optional) #### Worker params`,
				},
				resource.Attribute{
					Name:        "num_of_nodes",
					Description: `(Required) The worker node number of the kubernetes cluster. Default to 3. It is limited up to 50 and if you want to enlarge it, please apply white list or contact with us.`,
				},
				resource.Attribute{
					Name:        "worker_vswtich_ids",
					Description: `(Required) The vswitches used by worker, you can specific 1 or more than 1 vswitches.`,
				},
				resource.Attribute{
					Name:        "worker_data_disks",
					Description: `(Optional) The configurations of the data disks that are mounted to worker nodes. Each configuration includes disk type and disk size.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `(Optional) The data disk category of worker node. Its valid value are cloud, cloud_ssd, cloud_essd and cloud_efficiency.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) The data disk size of worker node. Its valid value range [40~500] in GB.`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `(Optional) Specifies whether disk encryption is enabled.`,
				},
				resource.Attribute{
					Name:        "auto_snapshot_policy_id",
					Description: `(Optional) The ID of the policy that is used to back up the data disk.`,
				},
				resource.Attribute{
					Name:        "performance_level",
					Description: `(Optional) The ID of the policy that is used to back up the data disk.`,
				},
				resource.Attribute{
					Name:        "worker_disk_size",
					Description: `(Optional) The system disk size of worker node. Its valid value range [20~32768] in GB.`,
				},
				resource.Attribute{
					Name:        "worker_disk_category",
					Description: `(Optional) The system disk category of worker node. Its valid value are cloud, cloud_ssd, cloud_essd and cloud_efficiency.`,
				},
				resource.Attribute{
					Name:        "worker_instance_types",
					Description: `(Required) The instance type of worker node. Specify one type for single AZ Cluster, three types for MultiAZ Cluster.`,
				},
				resource.Attribute{
					Name:        "worker_system_disk_performance_level",
					Description: `(Optional) #### Computed params (No need to configure)`,
				},
				resource.Attribute{
					Name:        "kube_config",
					Description: `(Optional) The path of kube config, like ` + "`" + `~/.kube/config` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "client_cert",
					Description: `(Optional) The path of client certificate, like ` + "`" + `~/.kube/client-cert.pem` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: `(Optional) The path of client key, like ` + "`" + `~/.kube/client-key.pem` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cluster_ca_cert",
					Description: `(Optional) The path of cluster ca certificate, like ` + "`" + `~/.kube/cluster-ca-cert.pem` + "`" + ``,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) The Zone where new kubernetes cluster will be located. If it is not be specified, the ` + "`" + `vswitch_ids` + "`" + ` should be set, its value will be vswitch's zone. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 90 mins) Used when creating the kubernetes cluster (until it reaches the initial ` + "`" + `running` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 60 mins) Used when activating the kubernetes cluster when necessary during update.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 60 mins) Used when terminating the kubernetes cluster. ## Attributes Reference The following attributes are exported:`,
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
					Name:        "vpc_id",
					Description: `The ID of VPC where the current cluster is located.`,
				},
				resource.Attribute{
					Name:        "slb_intranet",
					Description: `The ID of private load balancer where the current cluster master node is located.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of security group where the current cluster worker node is located.`,
				},
				resource.Attribute{
					Name:        "nat_gateway_id",
					Description: `The ID of nat gateway used to launch kubernetes cluster.`,
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
					Name:        "version",
					Description: `The Kubernetes server version for the cluster.`,
				},
				resource.Attribute{
					Name:        "worker_ram_role_name",
					Description: `The RamRole Name attached to worker node.`,
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
					Name:        "vpc_id",
					Description: `The ID of VPC where the current cluster is located.`,
				},
				resource.Attribute{
					Name:        "slb_intranet",
					Description: `The ID of private load balancer where the current cluster master node is located.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of security group where the current cluster worker node is located.`,
				},
				resource.Attribute{
					Name:        "nat_gateway_id",
					Description: `The ID of nat gateway used to launch kubernetes cluster.`,
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
					Name:        "version",
					Description: `The Kubernetes server version for the cluster.`,
				},
				resource.Attribute{
					Name:        "worker_ram_role_name",
					Description: `The RamRole Name attached to worker node.`,
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
			Type:             "alibabacloudstack_cs_kubernetes_node_pool",
			Category:         "Container Service for Kubernetes (ACK)",
			ShortDescription: `Provides a Alibabacloudstack resource to manage container kubernetes node pool.`,
			Description:      ``,
			Keywords: []string{
				"container",
				"service",
				"for",
				"kubernetes",
				"ack",
				"cs",
				"node",
				"pool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) The id of kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of node pool.`,
				},
				resource.Attribute{
					Name:        "vswitch_ids",
					Description: `(Required) The vswitches used by node pool workers.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required, Sensitive) The password of ssh login cluster node. You have to specify one of ` + "`" + `password` + "`" + ` ` + "`" + `key_name` + "`" + ` ` + "`" + `kms_encrypted_password` + "`" + ` fields.`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `(Required) The keypair of ssh login cluster node, you have to create it first. You have to specify one of ` + "`" + `password` + "`" + ` ` + "`" + `key_name` + "`" + ` ` + "`" + `kms_encrypted_password` + "`" + ` fields. Only ` + "`" + `key_name` + "`" + ` is supported in the management node pool.`,
				},
				resource.Attribute{
					Name:        "kms_encrypted_password",
					Description: `(Required) An KMS encrypts password used to a cs kubernetes. You have to specify one of ` + "`" + `password` + "`" + ` ` + "`" + `key_name` + "`" + ` ` + "`" + `kms_encrypted_password` + "`" + ` fields.`,
				},
				resource.Attribute{
					Name:        "system_disk_category",
					Description: `(Optional) The system disk category of worker node. Its valid value are ` + "`" + `cloud_ssd` + "`" + ` and ` + "`" + `cloud_efficiency` + "`" + `. Default to ` + "`" + `cloud_efficiency` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "system_disk_size",
					Description: `(Optional) The system disk category of worker node. Its valid value range [40~500] in GB. Default to ` + "`" + `120` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "data_disks",
					Description: `(Optional) The data disk configurations of worker nodes, such as the disk type and disk size.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `The type of the data disks. Valid values:` + "`" + `cloud` + "`" + `, ` + "`" + `cloud_efficiency` + "`" + `, ` + "`" + `cloud_ssd` + "`" + ` and ` + "`" + `cloud_essd` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of a data disk, Its valid value range [40~32768] in GB. Default to ` + "`" + `40` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `Specifies whether to encrypt data disks. Valid values: true and false. Default to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "performance_level",
					Description: `(Optional, Available in 1.120.0+) Worker node data disk performance level, when ` + "`" + `category` + "`" + ` values ` + "`" + `cloud_essd` + "`" + `, the optional values are ` + "`" + `PL0` + "`" + `, ` + "`" + `PL1` + "`" + `, ` + "`" + `PL2` + "`" + ` or ` + "`" + `PL3` + "`" + `, but the specific performance level is related to the disk capacity. For more information, see [Enhanced SSDs](https://www.alibabacloud.com/help/doc-detail/122389.htm). Default is ` + "`" + `PL1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Optional) The security group id for worker node.`,
				},
				resource.Attribute{
					Name:        "platform",
					Description: `(Optional, Available in 1.127.0+) The platform. One of ` + "`" + `AliyunLinux` + "`" + `, ` + "`" + `Windows` + "`" + `, ` + "`" + `CentOS` + "`" + `, ` + "`" + `WindowsCore` + "`" + `. If you select ` + "`" + `Windows` + "`" + ` or ` + "`" + `WindowsCore` + "`" + `, the ` + "`" + `passord` + "`" + ` is required.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Optional) Custom Image support. Must based on CentOS7 or AliyunLinux2.`,
				},
				resource.Attribute{
					Name:        "node_name_mode",
					Description: `(Optional) Each node name consists of a prefix, an IP substring, and a suffix. For example "customized,aliyun.com,5,test", if the node IP address is 192.168.0.55, the prefix is aliyun.com, IP substring length is 5, and the suffix is test, the node name will be aliyun.com00055test.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional) Windows instances support batch and PowerShell scripts. If your script file is larger than 1 KB, we recommend that you upload the script to Object Storage Service (OSS) and pull it through the internal endpoint of your OSS bucket.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A Map of tags to assign to the resource. It will be applied for ECS instances finally.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) A List of Kubernetes labels to assign to the nodes . Only labels that are applied with the ACK API are managed by this argument.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The label key.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The label value.`,
				},
				resource.Attribute{
					Name:        "taints",
					Description: `(Optional) A List of Kubernetes taints to assign to the nodes.`,
				},
				resource.Attribute{
					Name:        "management",
					Description: `(Optional, Available in 1.109.1+) Managed node pool configuration. When using a managed node pool, the node key must use ` + "`" + `key_name` + "`" + `. Detailed below.`,
				},
				resource.Attribute{
					Name:        "scaling_policy",
					Description: `(Optional, Available in 1.127.0+) The scaling mode. Valid values: ` + "`" + `release` + "`" + `, ` + "`" + `recycle` + "`" + `, default is ` + "`" + `release` + "`" + `. Standard mode(release): Create and release ECS instances based on requests.Swift mode(recycle): Create, stop, and restart ECS instances based on needs. New ECS instances are only created when no stopped ECS instance is avalible. This mode further accelerates the scaling process. Apart from ECS instances that use local storage, when an ECS instance is stopped, you are only chatged for storage space.`,
				},
				resource.Attribute{
					Name:        "scaling_config",
					Description: `(Optional, Available in 1.111.0+) Auto scaling node pool configuration. For more details, see ` + "`" + `scaling_config` + "`" + `. With auto-scaling is enabled, the nodes in the node pool will be labeled with ` + "`" + `k8s.aliyun.com=true` + "`" + ` to prevent system pods such as coredns, metrics-servers from being scheduled to elastic nodes, and to prevent node shrinkage from causing business abnormalities.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional, ForceNew, Available in 1.123.1+) The ID of the resource group,by default these cloud resources are automatically assigned to the default resource group.`,
				},
				resource.Attribute{
					Name:        "internet_charge_type",
					Description: `(Optional, Available in 1.123.1+) The billing method for network usage. Valid values ` + "`" + `PayByBandwidth` + "`" + ` and ` + "`" + `PayByTraffic` + "`" + `. Conflict with ` + "`" + `eip_internet_charge_type` + "`" + `, EIP and public network IP can only choose one.`,
				},
				resource.Attribute{
					Name:        "internet_max_bandwidth_out",
					Description: `(Optional, Available in 1.123.1+) The maximum outbound bandwidth for the public network. Unit: Mbit/s. Valid values: 0 to 100.`,
				},
				resource.Attribute{
					Name:        "spot_strategy",
					Description: `(Optional, Available in 1.123.1+) The preemption policy for the pay-as-you-go instance. This parameter takes effect only when ` + "`" + `instance_charge_type` + "`" + ` is set to ` + "`" + `PostPaid` + "`" + `. Valid value ` + "`" + `SpotWithPriceLimit` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "spot_price_limit",
					Description: `(Optional, Available in 1.123.1+) The maximum hourly price of the instance. This parameter takes effect only when ` + "`" + `spot_strategy` + "`" + ` is set to ` + "`" + `SpotWithPriceLimit` + "`" + `. A maximum of three decimal places are allowed.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Optional, Available in 1.123.1+) Spot instance type.`,
				},
				resource.Attribute{
					Name:        "price_limit",
					Description: `(Optional, Available in 1.123.1+) The maximum hourly price of the spot instance.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `(Optional, Available in 1.127.0+) The instance list. Add existing nodes under the same cluster VPC to the node pool.`,
				},
				resource.Attribute{
					Name:        "keep_instance_name",
					Description: `(Optional, Available in 1.127.0+) Add an existing instance to the node pool, whether to keep the original instance name. It is recommended to set to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "format_disk",
					Description: `(Optional, Available in 1.127.0+) After you select this check box, if data disks have been attached to the specified ECS instances and the file system of the last data disk is uninitialized, the system automatically formats the last data disk to ext4 and mounts the data disk to /var/lib/docker and /var/lib/kubelet. The original data on the disk will be cleared. Make sure that you back up data in advance. If no data disk is mounted on the ECS instance, no new data disk will be purchased. Default is ` + "`" + `false` + "`" + `. #### tags The tags example： ` + "`" + `` + "`" + `` + "`" + ` tags { "key-a" = "value-a" "key-b" = "value-b" "env" = "prod" } ` + "`" + `` + "`" + `` + "`" + ` #### management The following arguments are supported in the ` + "`" + `management` + "`" + ` configuration block:`,
				},
				resource.Attribute{
					Name:        "auto_repair",
					Description: `(Optional) Whether automatic repair, Default to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "surge",
					Description: `(Optional) Number of additional nodes. You have to specify one of surge, surge_percentage.`,
				},
				resource.Attribute{
					Name:        "surge_percentage",
					Description: `(Optional) Proportion of additional nodes. You have to specify one of surge, surge_percentage.`,
				},
				resource.Attribute{
					Name:        "max_unavailable",
					Description: `(Required) Max number of unavailable nodes. Default to ` + "`" + `1` + "`" + `. #### scaling_config The following arguments are supported in the ` + "`" + `scaling_config` + "`" + ` configuration block:`,
				},
				resource.Attribute{
					Name:        "min_size",
					Description: `(Required, Available in 1.111.0+) Min number of instances in a auto scaling group, its valid value range [0~1000].`,
				},
				resource.Attribute{
					Name:        "max_size",
					Description: `(Required, Available in 1.111.0+) Max number of instances in a auto scaling group, its valid value range [0~1000]. ` + "`" + `max_size` + "`" + ` has to be greater than ` + "`" + `min_size` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional, Available in 1.111.0+) Instance classification, not required. Vaild value: ` + "`" + `cpu` + "`" + `, ` + "`" + `gpu` + "`" + `, ` + "`" + `gpushare` + "`" + ` and ` + "`" + `spot` + "`" + `. Default: ` + "`" + `cpu` + "`" + `. The actual instance type is determined by ` + "`" + `instance_types` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "is_bond_eip",
					Description: `(Optional, Available in 1.111.0+) Whether to bind EIP for an instance. Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "eip_internet_charge_type",
					Description: `(Optional, Available in 1.111.0+) EIP billing type. ` + "`" + `PayByBandwidth` + "`" + `: Charged at fixed bandwidth. ` + "`" + `PayByTraffic` + "`" + `: Billed as used traffic. Default: ` + "`" + `PayByBandwidth` + "`" + `. Conflict with ` + "`" + `internet_charge_type` + "`" + `, EIP and public network IP can only choose one.`,
				},
				resource.Attribute{
					Name:        "eip_bandwidth",
					Description: `(Optional, Available in 1.111.0+) Peak EIP bandwidth. Its valid value range [1~500] in Mbps. Default to ` + "`" + `5` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the node pool, format cluster_id:nodepool_id.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `The cluster id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the nodepool.`,
				},
				resource.Attribute{
					Name:        "vswitch_ids",
					Description: `The vswitches used by node pool workers.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The image used by node pool workers.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of security group where the current cluster worker node is located.`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `(Available in 1.105.0+) Id of the Scaling Group. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 90 mins) Used when creating node-pool in the kubernetes cluster (until it reaches the initial ` + "`" + `active` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 60 mins) Used when activating the node-pool in the kubernetes cluster when necessary during update.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 60 mins) Used when deleting node-pool in kubernetes cluster. ## Import Cluster nodepool can be imported using the id, e.g. Then complete the nodepool.tf accords to the result of ` + "`" + `terraform plan` + "`" + `. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_cs_node_pool.custom_nodepool cluster_id:nodepool_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the node pool, format cluster_id:nodepool_id.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `The cluster id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the nodepool.`,
				},
				resource.Attribute{
					Name:        "vswitch_ids",
					Description: `The vswitches used by node pool workers.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The image used by node pool workers.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of security group where the current cluster worker node is located.`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `(Available in 1.105.0+) Id of the Scaling Group. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 90 mins) Used when creating node-pool in the kubernetes cluster (until it reaches the initial ` + "`" + `active` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 60 mins) Used when activating the node-pool in the kubernetes cluster when necessary during update.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 60 mins) Used when deleting node-pool in kubernetes cluster. ## Import Cluster nodepool can be imported using the id, e.g. Then complete the nodepool.tf accords to the result of ` + "`" + `terraform plan` + "`" + `. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_cs_node_pool.custom_nodepool cluster_id:nodepool_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_csb_project",
			Category:         "CSB",
			ShortDescription: `Provides a Alibabacloudstack resource to manage CSB Project .`,
			Description:      ``,
			Keywords: []string{
				"csb",
				"project",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "data",
					Description: `(Required, ForceNew) Infomation of CSB Project.`,
				},
				resource.Attribute{
					Name:        "csb_id",
					Description: `(Required, ForceNew) id of CSB instance where repository is created.`,
				},
				resource.Attribute{
					Name:        "project_name",
					Description: `(Required, ForceNew) Name of CSB Project. It can contain 2 to 64 characters. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "csb_id",
					Description: `The id of CSB instance.`,
				},
				resource.Attribute{
					Name:        "project_name",
					Description: `The project name of CSB Project.`,
				},
				resource.Attribute{
					Name:        "project_owner_name",
					Description: `The project owner name of CSB Project.`,
				},
				resource.Attribute{
					Name:        "gmt_modified",
					Description: `The project modified time of CSB Project.`,
				},
				resource.Attribute{
					Name:        "gmt_create",
					Description: `The project create time of CSB Project.`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `The owner id of CSB Project.`,
				},
				resource.Attribute{
					Name:        "api_num",
					Description: `The api num of CSB Project.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `The user id of CSB Project.`,
				},
				resource.Attribute{
					Name:        "delete_flag",
					Description: `The delete flag of CSB Project.`,
				},
				resource.Attribute{
					Name:        "cs_id",
					Description: `The project id of CSB Project.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The project status of CSB Project.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "csb_id",
					Description: `The id of CSB instance.`,
				},
				resource.Attribute{
					Name:        "project_name",
					Description: `The project name of CSB Project.`,
				},
				resource.Attribute{
					Name:        "project_owner_name",
					Description: `The project owner name of CSB Project.`,
				},
				resource.Attribute{
					Name:        "gmt_modified",
					Description: `The project modified time of CSB Project.`,
				},
				resource.Attribute{
					Name:        "gmt_create",
					Description: `The project create time of CSB Project.`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `The owner id of CSB Project.`,
				},
				resource.Attribute{
					Name:        "api_num",
					Description: `The api num of CSB Project.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `The user id of CSB Project.`,
				},
				resource.Attribute{
					Name:        "delete_flag",
					Description: `The delete flag of CSB Project.`,
				},
				resource.Attribute{
					Name:        "cs_id",
					Description: `The project id of CSB Project.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The project status of CSB Project.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_data_works_connection",
			Category:         "Data Works",
			ShortDescription: `Provides a AlibabacloudStack Data Works Connection resource.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"works",
				"connection",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The ID of the project.`,
				},
				resource.Attribute{
					Name:        "connection_type",
					Description: `(Required) Type of connection.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Required) Details of the data source.`,
				},
				resource.Attribute{
					Name:        "env_type",
					Description: `(Required) The environment to which the data source belongs, including 0 (development environment) and 1 (production environment).`,
				},
				resource.Attribute{
					Name:        "sub_type",
					Description: `(Optional) Sub-types of strings, for scenarios where some parent types include sub-types. There are currently the following combinations:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the data source.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the connection. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "connection_id",
					Description: `The resource ID of Connection. The value formats as ` + "`" + `<connection_id>:<$.ProjectId>` + "`" + `. ## Import Data Works Connection can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_data_works_connection.example <connection_id>:<$.ProjectId> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "connection_id",
					Description: `The resource ID of Connection. The value formats as ` + "`" + `<connection_id>:<$.ProjectId>` + "`" + `. ## Import Data Works Connection can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_data_works_connection.example <connection_id>:<$.ProjectId> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_data_works_folder",
			Category:         "Data Works",
			ShortDescription: `Provides a AlibabacloudStack Data Works Folder resource.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"works",
				"folder",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "folder_path",
					Description: `(Required) Folder Path. The folder path composed with for part: ` + "`" + `Business Flow/{Business Flow Name}/[folderDi|folderMaxCompute|folderGeneral|folderJdbc|folderUserDefined]/{Directory Name}` + "`" + `. The first segment of path must be ` + "`" + `Business Flow` + "`" + `, and sencond segment of path must be a Business Flow Name within the project. The third part of path must be one of those keywords:` + "`" + `folderDi|folderMaxCompute|folderGeneral|folderJdbc|folderUserDefined` + "`" + `. Then the finial part of folder path can be specified in yourself.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required, ForceNew) The ID of the project. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID of Folder. The value formats as ` + "`" + `<folder_id>:<$.ProjectId>` + "`" + `. ## Import Data Works Folder can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_data_works_folder.example <folder_id>:<$.ProjectId> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID of Folder. The value formats as ` + "`" + `<folder_id>:<$.ProjectId>` + "`" + `. ## Import Data Works Folder can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_data_works_folder.example <folder_id>:<$.ProjectId> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_data_works_project",
			Category:         "Data Works",
			ShortDescription: `Provides a AlibabacloudStack Data Works Project resource.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"works",
				"project",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Computed) The ID of the project.`,
				},
				resource.Attribute{
					Name:        "project_name",
					Description: `(Required) The name of the project.`,
				},
				resource.Attribute{
					Name:        "task_auth_type",
					Description: `(Optional) The task auth type of the project, default value is PROJECT.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_data_works_remind",
			Category:         "Data Works",
			ShortDescription: `Provides a AlibabacloudStack Data Works Remind resource.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"works",
				"remind",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "alert_unit",
					Description: `(Required) The granularity of the receiving object, including OWNER (task owner) and OTHER (designated person).`,
				},
				resource.Attribute{
					Name:        "remind_name",
					Description: `(Required) The name of the custom rule cannot exceed 128 characters.`,
				},
				resource.Attribute{
					Name:        "remind_type",
					Description: `(Required) Trigger conditions, including FINISHED, UNFINISHED, ERROR, CYCLE_UNFINISHED and TIMEOUT.`,
				},
				resource.Attribute{
					Name:        "remind_unit",
					Description: `(Required) Types of objects, including NODE (task node), BASELINE (baseline), PROJECT (workspace) and BIZPROCESS (business process).`,
				},
				resource.Attribute{
					Name:        "dnd_end",
					Description: `(Optional) Do not disturb deadline, format HH: MM. The value range of hh is 0,23, and the value range of mm is 0,59.Default 00:00.`,
				},
				resource.Attribute{
					Name:        "node_ids",
					Description: `(Optional) The monitored task node id when the object type(remind_type) is NODE. Multiple IDs are separated by commas (,), and a rule can monitor up to 50 nodes.`,
				},
				resource.Attribute{
					Name:        "baseline_ids",
					Description: `(Optional) The monitored baseline id when the object type(remind_type) is BASELINE. Multiple IDs are separated by commas (,), and one rule can monitor up to 5 baselines.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) The monitored workspace id when the object type(remind_type) is PROJECT. A rule can only monitor one workspace.`,
				},
				resource.Attribute{
					Name:        "biz_process_ids",
					Description: `(Optional) The monitored business process id when the object type(remind_type) is BIZPROCESS. Multiple business process ids are separated by commas (,), and a rule can monitor up to 5 business processes.`,
				},
				resource.Attribute{
					Name:        "max_alert_times",
					Description: `(Optional) Maximum number of alarms. The minimum value is 1, the maximum value is 10 and the default value is 3.`,
				},
				resource.Attribute{
					Name:        "alert_interval",
					Description: `(Optional) Minimum alarm interval, in seconds. The minimum value is 1200 and the default value is 1800.`,
				},
				resource.Attribute{
					Name:        "detail",
					Description: `(Optional) The descriptions of different trigger conditions are as follows:`,
				},
				resource.Attribute{
					Name:        "alert_methods",
					Description: `(Optional) The alarm methods include MAIL (mail), SMS (short message) and PHONE (telephone, only supported by DataWorks Professional and above). Multiple alarm modes are separated by English commas (,).`,
				},
				resource.Attribute{
					Name:        "alert_targets",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "robot_urls",
					Description: `(Optional) The webhook addresses of DingTalk robots, and multiple webhook addresses are separated by English commas (,).`,
				},
				resource.Attribute{
					Name:        "use_flag",
					Description: `(Optional) TOpen and close rules, including true and false. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "remind_id",
					Description: `The resource ID of Remind. The value formats as ` + "`" + `<remind_id>:<$.ProjectId>` + "`" + `. ## Import Data Works Remind can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_data_works_remind.example <remind_id>:<$.ProjectId> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "remind_id",
					Description: `The resource ID of Remind. The value formats as ` + "`" + `<remind_id>:<$.ProjectId>` + "`" + `. ## Import Data Works Remind can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_data_works_remind.example <remind_id>:<$.ProjectId> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_data_works_user",
			Category:         "Data Works",
			ShortDescription: `Provides a AlibabacloudStack Data Works User resource.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"works",
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The ID of the project.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `(Required) User ID to be added.`,
				},
				resource.Attribute{
					Name:        "role_code",
					Description: `(Optional) If it is not blank, the user will be added to this role.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_data_works_user_role_binding",
			Category:         "Data Works",
			ShortDescription: `Provides a AlibabacloudStack Data Works UserRoleBinding resource.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"works",
				"user",
				"role",
				"binding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The ID of the project.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `(Required) Alibaba Cloud Account ID.`,
				},
				resource.Attribute{
					Name:        "role_code",
					Description: `(Required) Code of DataWorks workspace role.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_datahub_project",
			Category:         "Datahub Service",
			ShortDescription: `Provides a Alibabacloudstack datahub project resource.`,
			Description:      ``,
			Keywords: []string{
				"datahub",
				"service",
				"project",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) The name of the datahub project. Its length is limited to 3-32 and only characters such as letters, digits and '_' are allowed. It is case-insensitive.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Comment of the datahub project. It cannot be longer than 255 characters. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the datahub project. It is the same as its name.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the datahub project. It is a human-readable string rather than 64-bits UTC.`,
				},
				resource.Attribute{
					Name:        "last_modify_time",
					Description: `Last modify time of the datahub project. It is the same as`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the datahub project. It is the same as its name.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the datahub project. It is a human-readable string rather than 64-bits UTC.`,
				},
				resource.Attribute{
					Name:        "last_modify_time",
					Description: `Last modify time of the datahub project. It is the same as`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_datahub_subscription",
			Category:         "Datahub Service",
			ShortDescription: `Provides a Alibabacloudstack datahub subscription resource.`,
			Description:      ``,
			Keywords: []string{
				"datahub",
				"service",
				"subscription",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_name",
					Description: `(Required, ForceNew) The name of the datahub project that the subscription belongs to. Its length is limited to 3-32 and only characters such as letters, digits and '_' are allowed. It is case-insensitive.`,
				},
				resource.Attribute{
					Name:        "topic_name",
					Description: `(Required, ForceNew) The name of the datahub topic that the subscription belongs to. Its length is limited to 1-128 and only characters such as letters, digits and '_' are allowed. It is case-insensitive.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Comment of the datahub subscription. It cannot be longer than 255 characters. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the datahub subscription as terraform resource. It was composed of project name, topic name and practical subscription ID generated from server side. Format to ` + "`" + `<project_name>:<topic_name>:<sub_id>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sub_id",
					Description: `The identidy of the subscription, generate from server side.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the datahub subscription. It is a human-readable string rather than 64-bits UTC.`,
				},
				resource.Attribute{
					Name:        "last_modify_time",
					Description: `Last modify time of the datahub subscription. It is the same as`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the datahub subscription as terraform resource. It was composed of project name, topic name and practical subscription ID generated from server side. Format to ` + "`" + `<project_name>:<topic_name>:<sub_id>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sub_id",
					Description: `The identidy of the subscription, generate from server side.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the datahub subscription. It is a human-readable string rather than 64-bits UTC.`,
				},
				resource.Attribute{
					Name:        "last_modify_time",
					Description: `Last modify time of the datahub subscription. It is the same as`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_datahub_topic",
			Category:         "Datahub Service",
			ShortDescription: `Provides a Alibabacloudstack datahub topic resource.`,
			Description:      ``,
			Keywords: []string{
				"datahub",
				"service",
				"topic",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) The name of the datahub topic. Its length is limited to 1-128 and only characters such as letters, digits and '_' are allowed. It is case-insensitive.`,
				},
				resource.Attribute{
					Name:        "project_name",
					Description: `(Required, ForceNew) The name of the datahub project that this topic belongs to. It is case-insensitive.`,
				},
				resource.Attribute{
					Name:        "shard_count",
					Description: `(Optional, ForceNew) The number of shards this topic contains. The permitted range of values is [1, 10]. The default value is 1.`,
				},
				resource.Attribute{
					Name:        "life_cycle",
					Description: `(Optional) How many days this topic lives. The permitted range of values is [1, 7]. The default value is 3.`,
				},
				resource.Attribute{
					Name:        "record_type",
					Description: `(Optional, ForceNew) The type of this topic. Its value must be one of {BLOB, TUPLE}. For BLOB topic, data will be organized as binary and encoded by BASE64. For TUPLE topic, data has fixed schema. The default value is "TUPLE" with a schema {STRING}.`,
				},
				resource.Attribute{
					Name:        "record_schema",
					Description: `(Optional, ForceNew) Schema of this topic, required only for TUPLE topic. Supported data types (case-insensitive) are: - BIGINT - STRING - BOOLEAN - DOUBLE - TIMESTAMP`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Comment of the datahub topic. It cannot be longer than 255 characters.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the datahub topic. It was composed of project name and its name, and formats to ` + "`" + `<project_name>:<name>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the datahub topic. It is a human-readable string rather than 64-bits UTC.`,
				},
				resource.Attribute{
					Name:        "last_modify_time",
					Description: `Last modify time of the datahub topic. It is the same as`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the datahub topic. It was composed of project name and its name, and formats to ` + "`" + `<project_name>:<name>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the datahub topic. It is a human-readable string rather than 64-bits UTC.`,
				},
				resource.Attribute{
					Name:        "last_modify_time",
					Description: `Last modify time of the datahub topic. It is the same as`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_db_account",
			Category:         "RDS",
			ShortDescription: `Provides an RDS account resource.`,
			Description:      ``,
			Keywords: []string{
				"rds",
				"db",
				"account",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, ForceNew) The Id of instance in which account belongs.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and have no more than 16 characters.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional, Sensitive) Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters. You have to specify one of ` + "`" + `password` + "`" + ` and ` + "`" + `kms_encrypted_password` + "`" + ` fields.`,
				},
				resource.Attribute{
					Name:        "kms_encrypted_password",
					Description: `(Optional) An KMS encrypts password used to a db account. If the ` + "`" + `password` + "`" + ` is filled in, this field will be ignored.`,
				},
				resource.Attribute{
					Name:        "kms_encryption_context",
					Description: `(Optional) An KMS encryption context used to decrypt ` + "`" + `kms_encrypted_password` + "`" + ` before creating or updating a db account with ` + "`" + `kms_encrypted_password` + "`" + `. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when ` + "`" + `kms_encrypted_password` + "`" + ` is set.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional, ForceNew)Privilege type of account. - Normal: Common privilege. - Super: High privilege. Default to Normal. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The current account resource ID. Composed of instance ID and account name with format ` + "`" + `<instance_id>:<name>` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The current account resource ID. Composed of instance ID and account name with format ` + "`" + `<instance_id>:<name>` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_db_account_privilege",
			Category:         "RDS",
			ShortDescription: `Provides an RDS account privilege resource.`,
			Description:      ``,
			Keywords: []string{
				"rds",
				"db",
				"account",
				"privilege",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, ForceNew) The Id of instance in which account belongs.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `(Required, ForceNew) A specified account name.`,
				},
				resource.Attribute{
					Name:        "privilege",
					Description: `The privilege of one account access database. Valid values: - ReadOnly: This value is only for MySQL, MariaDB and SQL Server - ReadWrite: This value is only for MySQL, MariaDB and SQL Server Default to "ReadOnly".`,
				},
				resource.Attribute{
					Name:        "db_names",
					Description: `(Required) List of specified database name. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The current account resource ID. Composed of instance ID, account name and privilege with format ` + "`" + `<instance_id>:<name>:<privilege>` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The current account resource ID. Composed of instance ID, account name and privilege with format ` + "`" + `<instance_id>:<name>:<privilege>` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_db_backup_policy",
			Category:         "RDS",
			ShortDescription: `Provides an RDS backup policy resource.`,
			Description:      ``,
			Keywords: []string{
				"rds",
				"db",
				"backup",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, ForceNew) The Id of instance that can run database.`,
				},
				resource.Attribute{
					Name:        "preferred_backup_period",
					Description: `(Optional) DB Instance backup period. Please set at least two days to ensure backing up at least twice a week. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]. Default to ["Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"].`,
				},
				resource.Attribute{
					Name:        "preferred_backup_time",
					Description: `(Optional) DB instance backup time, in the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. Default to "02:00Z-03:00Z". China time is 8 hours behind it.`,
				},
				resource.Attribute{
					Name:        "backup_retention_period",
					Description: `(Optional, available in 1.69.0+) Instance backup retention days. Valid values: [7-730]. Default to 7. But mysql local disk is unlimited.`,
				},
				resource.Attribute{
					Name:        "enable_backup_log",
					Description: `(Optional, available in 1.68.0+) Whether to backup instance log. Valid values are ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + `, Default to ` + "`" + `true` + "`" + `. Note: The 'Basic Edition' category Rds instance does not support setting log backup. [What is Basic Edition](https://www.alibabacloud.com/help/doc-detail/48980.htm).`,
				},
				resource.Attribute{
					Name:        "log_backup_retention_period",
					Description: `(Optional, available in 1.69.0+) Instance log backup retention days. Valid when the ` + "`" + `enable_backup_log` + "`" + ` is ` + "`" + `1` + "`" + `. Valid values: [7-730]. Default to 7. It cannot be larger than ` + "`" + `backup_retention_period` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "local_log_retention_hours",
					Description: `(Optional, available in 1.69.0+) Instance log backup local retention hours. Valid when the ` + "`" + `enable_backup_log` + "`" + ` is ` + "`" + `true` + "`" + `. Valid values: [0-7`,
				},
				resource.Attribute{
					Name:        "local_log_retention_space",
					Description: `(Optional, available in 1.69.0+) Instance log backup local retention space. Valid when the ` + "`" + `enable_backup_log` + "`" + ` is ` + "`" + `true` + "`" + `. Valid values: [5-50].`,
				},
				resource.Attribute{
					Name:        "high_space_usage_protection",
					Description: `(Optional, available in 1.69.0+) Instance high space usage protection policy. Valid when the ` + "`" + `enable_backup_log` + "`" + ` is ` + "`" + `true` + "`" + `. Valid values are ` + "`" + `Enable` + "`" + `, ` + "`" + `Disable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "log_backup_frequency",
					Description: `(Optional, available in 1.69.0+) Instance log backup frequency. Valid when the instance engine is ` + "`" + `SQLServer` + "`" + `. Valid values are ` + "`" + `LogInterval` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "compress_type",
					Description: `(Optional, available in 1.69.0+) The compress type of instance policy. Valid values are ` + "`" + `1` + "`" + `, ` + "`" + `4` + "`" + `, ` + "`" + `8` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "archive_backup_retention_period",
					Description: `(Optional, available in 1.69.0+) Instance archive backup retention days. Valid when the ` + "`" + `enable_backup_log` + "`" + ` is ` + "`" + `true` + "`" + ` and instance is mysql local disk. Valid values: [30-1095], and ` + "`" + `archive_backup_retention_period` + "`" + ` must larger than ` + "`" + `backup_retention_period` + "`" + ` 730.`,
				},
				resource.Attribute{
					Name:        "archive_backup_keep_count",
					Description: `(Optional, available in 1.69.0+) Instance archive backup keep count. Valid when the ` + "`" + `enable_backup_log` + "`" + ` is ` + "`" + `true` + "`" + ` and instance is mysql local disk. When ` + "`" + `archive_backup_keep_policy` + "`" + ` is ` + "`" + `ByMonth` + "`" + ` Valid values: [1-31]. When ` + "`" + `archive_backup_keep_policy` + "`" + ` is ` + "`" + `ByWeek` + "`" + ` Valid values: [1-7].`,
				},
				resource.Attribute{
					Name:        "archive_backup_keep_policy",
					Description: `(Optional, available in 1.69.0+) Instance archive backup keep policy. Valid when the ` + "`" + `enable_backup_log` + "`" + ` is ` + "`" + `true` + "`" + ` and instance is mysql local disk. Valid values are ` + "`" + `ByMonth` + "`" + `, ` + "`" + `Disable` + "`" + `, ` + "`" + `KeepAll` + "`" + `. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The current backup policy resource ID. It is same as 'instance_id'.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The current backup policy resource ID. It is same as 'instance_id'.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_db_connection",
			Category:         "RDS",
			ShortDescription: `Provides an RDS instance connection resource.`,
			Description:      ``,
			Keywords: []string{
				"rds",
				"db",
				"connection",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, ForceNew) The Id of instance that can run database.`,
				},
				resource.Attribute{
					Name:        "connection_prefix",
					Description: `(ForceNew) Prefix of an Internet connection string. It must be checked for uniqueness. It may consist of lowercase letters, numbers, and underlines, and must start with a letter and have no more than 30 characters. Default to <instance_id> + 'tf'.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Internet connection port. Valid value: [3001-3999]. Default to 3306. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The current instance connection resource ID. Composed of instance ID and connection string with format ` + "`" + `<instance_id>:<connection_prefix>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "connection_prefix",
					Description: `Prefix of a connection string.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Connection instance port.`,
				},
				resource.Attribute{
					Name:        "connection_string",
					Description: `Connection instance string.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The ip address of connection string.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The current instance connection resource ID. Composed of instance ID and connection string with format ` + "`" + `<instance_id>:<connection_prefix>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "connection_prefix",
					Description: `Prefix of a connection string.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Connection instance port.`,
				},
				resource.Attribute{
					Name:        "connection_string",
					Description: `Connection instance string.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The ip address of connection string.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_db_database",
			Category:         "RDS",
			ShortDescription: `Provides an RDS database resource.`,
			Description:      ``,
			Keywords: []string{
				"rds",
				"db",
				"database",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, ForceNew) The Id of instance that can run database.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Name of the database requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and have no more than 64 characters.`,
				},
				resource.Attribute{
					Name:        "character_set",
					Description: `(Required) Character set. The value range is limited to the following: - MySQL: [ utf8, gbk, latin1, utf8mb4 ] \(` + "`" + `utf8mb4` + "`" + ` only supports versions 5.5 and 5.6\). - SQLServer: [ Chinese_PRC_CI_AS, Chinese_PRC_CS_AS, SQL_Latin1_General_CP1_CI_AS, SQL_Latin1_General_CP1_CS_AS, Chinese_PRC_BIN ] - PostgreSQL: [ KOI8U、UTF8、WIN866、WIN874、WIN1250、WIN1251、WIN1252、WIN1253、WIN1254、WIN1255、WIN1256、WIN1257、WIN1258、EUC_CN、EUC_KR、EUC_TW、EUC_JP、EUC_JIS_2004、KOI8R、MULE_INTERNAL、LATIN1、LATIN2、LATIN3、LATIN4、LATIN5、LATIN6、LATIN7、LATIN8、LATIN9、LATIN10、ISO_8859_5、ISO_8859_6、ISO_8859_7、ISO_8859_8、SQL_ASCII ] More details refer to [API Docs](https://www.alibabacloud.com/help/zh/doc-detail/26258.htm)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The current database resource ID. Composed of instance ID and database name with format ` + "`" + `<instance_id>:<name>` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The current database resource ID. Composed of instance ID and database name with format ` + "`" + `<instance_id>:<name>` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_db_instance",
			Category:         "RDS",
			ShortDescription: `Provides an RDS instance resource.`,
			Description:      ``,
			Keywords: []string{
				"rds",
				"db",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "engine",
					Description: `(Required,ForceNew) Database type. Value options: MySQL, SQLServer, PostgreSQL, and PPAS.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `(Required,ForceNew) Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/doc-detail/26228.htm) ` + "`" + `EngineVersion` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Required) DB Instance type. For details, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/26312.htm).`,
				},
				resource.Attribute{
					Name:        "instance_storage",
					Description: `(Required) User-defined DB instance storage space. Value range: - [5, 2000] for MySQL/PostgreSQL/PPAS HA dual node edition; - [20,1000] for MySQL 5.7 basic single node edition; - [10, 2000] for SQL Server 2008R2; - [20,2000] for SQL Server 2012 basic single node edition Increase progressively at a rate of 5 GB. For details, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/26312.htm). Note: There is extra 5 GB storage for SQL Server Instance and it is not in specified ` + "`" + `instance_storage` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `(Required) The type of storage media that is used for the instance.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Optional) The name of DB instance. It a string of 2 to 256 characters.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(ForceNew) The Zone to launch the DB instance.`,
				},
				resource.Attribute{
					Name:        "encryption_key",
					Description: `(Optional) Add encryptionkey to the DBInstance.`,
				},
				resource.Attribute{
					Name:        "zone_id_slave1",
					Description: `(Optional) The zone ID of the secondary instance.`,
				},
				resource.Attribute{
					Name:        "zone_id_slave",
					Description: `(Optional) The zone ID of the secondary instance.`,
				},
				resource.Attribute{
					Name:        "tde_status",
					Description: `(Optional) Enables the Transparent Data Encryption (TDE) function for an ApsaraDB for RDS instance.`,
				},
				resource.Attribute{
					Name:        "enable_ssl",
					Description: `(Optional) To enable the SSL encryption of an ApsaraDB RDS instance. If it is a multi-zone and ` + "`" + `vswitch_id` + "`" + ` is specified, the vswitch must in the one of them. The multiple zone ID can be retrieved by setting ` + "`" + `multi` + "`" + ` to "true" in the data source ` + "`" + `alibabacloudstack_zones` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(ForceNew) The virtual switch ID to launch DB instances in one VPC.`,
				},
				resource.Attribute{
					Name:        "security_ips",
					Description: `(Optional) List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]). ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The RDS instance ID.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `RDS database connection port.`,
				},
				resource.Attribute{
					Name:        "connection_string",
					Description: `RDS database connection string. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 20 mins) Used when creating the db instance (until it reaches the initial ` + "`" + `Running` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 30 mins) Used when updating the db instance (until it reaches the initial ` + "`" + `Running` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 20 mins) Used when terminating the db instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The RDS instance ID.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `RDS database connection port.`,
				},
				resource.Attribute{
					Name:        "connection_string",
					Description: `RDS database connection string. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 20 mins) Used when creating the db instance (until it reaches the initial ` + "`" + `Running` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 30 mins) Used when updating the db instance (until it reaches the initial ` + "`" + `Running` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 20 mins) Used when terminating the db instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_db_read_write_splitting_connection",
			Category:         "RDS",
			ShortDescription: `Provides an RDS instance read write splitting connection resource.`,
			Description:      ``,
			Keywords: []string{
				"rds",
				"db",
				"read",
				"write",
				"splitting",
				"connection",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, ForceNew) The Id of instance that can run database.`,
				},
				resource.Attribute{
					Name:        "distribution_type",
					Description: `(Required) Read weight distribution mode. Values are as follows: ` + "`" + `Standard` + "`" + ` indicates automatic weight distribution based on types, ` + "`" + `Custom` + "`" + ` indicates custom weight distribution.`,
				},
				resource.Attribute{
					Name:        "connection_prefix",
					Description: `(Optional, ForceNew) Prefix of an Internet connection string. It must be checked for uniqueness. It may consist of lowercase letters, numbers, and underlines, and must start with a letter and have no more than 30 characters. Default to <instance_id> + 'rw'.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Intranet connection port. Valid value: [3001-3999]. Default to 3306.`,
				},
				resource.Attribute{
					Name:        "max_delay_time",
					Description: `(Optional) Delay threshold, in seconds. The value range is 0 to 7200. Default to 30. Read requests are not routed to the read-only instances with a delay greater than the threshold.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) Read weight distribution. Read weights increase at a step of 100 up to 10,000. Enter weights in the following format: {"Instanceid":"Weight","Instanceid":"Weight"}. This parameter must be set when distribution_type is set to Custom. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Id of DB instance.`,
				},
				resource.Attribute{
					Name:        "connection_string",
					Description: `Connection instance string.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Id of DB instance.`,
				},
				resource.Attribute{
					Name:        "connection_string",
					Description: `Connection instance string.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_db_readonly_instance",
			Category:         "RDS",
			ShortDescription: `Provides an RDS readonly instance resource.`,
			Description:      ``,
			Keywords: []string{
				"rds",
				"db",
				"readonly",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "engine_version",
					Description: `(Required, ForceNew) Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/doc-detail/26228.htm) ` + "`" + `EngineVersion` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "master_db_instance_id",
					Description: `(Required) ID of the master instance.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Required) DB Instance type. For details, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/26312.htm).`,
				},
				resource.Attribute{
					Name:        "instance_storage",
					Description: `(Required) User-defined DB instance storage space. Value range: [5, 2000] for MySQL/SQL Server HA dual node edition. Increase progressively at a rate of 5 GB. For details, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/26312.htm).`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Optional) The name of DB instance. It a string of 2 to 256 characters.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `(Optional) Set of parameters needs to be set after DB instance was launched. Available parameters can refer to the latest docs [View database parameter templates](https://www.alibabacloud.com/help/doc-detail/26284.htm).`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Optional, ForceNew) The Zone to launch the DB instance.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Optional, ForceNew) The virtual switch ID to launch DB instances in one VPC.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string. - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.`,
				},
				resource.Attribute{
					Name:        "db_instance_storage_type",
					Description: `(Required) The storage type of the instance. Valid values: local_ssd: specifies to use local SSDs. This value is recommended. cloud_ssd: specifies to use standard SSDs. cloud_essd: specifies to use enhanced SSDs (ESSDs). cloud_essd2: specifies to use enhanced SSDs (ESSDs). cloud_essd3: specifies to use enhanced SSDs (ESSDs). ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The RDS instance ID.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `Database type.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `RDS database connection port.`,
				},
				resource.Attribute{
					Name:        "connection_string",
					Description: `RDS database connection string. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 60 mins) Used when creating the db instance (until it reaches the initial ` + "`" + `Running` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 30 mins) Used when updating the db instance (until it reaches the initial ` + "`" + `Running` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 20 mins) Used when terminating the db instance. ## Import RDS readonly instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_db_readonly_instance.example rm-abc12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The RDS instance ID.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `Database type.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `RDS database connection port.`,
				},
				resource.Attribute{
					Name:        "connection_string",
					Description: `RDS database connection string. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 60 mins) Used when creating the db instance (until it reaches the initial ` + "`" + `Running` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 30 mins) Used when updating the db instance (until it reaches the initial ` + "`" + `Running` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 20 mins) Used when terminating the db instance. ## Import RDS readonly instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_db_readonly_instance.example rm-abc12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_dbs_backup_plan",
			Category:         "Database Backup(DBS)",
			ShortDescription: `Provides a Alibabacloudstack DBS Backup Plan resource.`,
			Description:      ``,
			Keywords: []string{
				"database",
				"backup",
				"dbs",
				"plan",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "backup_method",
					Description: `(Required, ForceNew) Backup method. Valid values: ` + "`" + `logical` + "`" + `, ` + "`" + `physical` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "backup_plan_name",
					Description: `(Required, ForceNew) The name of the resource.`,
				},
				resource.Attribute{
					Name:        "database_type",
					Description: `(Required, ForceNew) Database type. Valid values: ` + "`" + `MySQL` + "`" + `, ` + "`" + `MSSQL` + "`" + `, ` + "`" + `Oracle` + "`" + `, ` + "`" + `MongoDB` + "`" + `, ` + "`" + `Redis` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_class",
					Description: `(Required, ForceNew) The instance class. Valid values: ` + "`" + `large` + "`" + `, ` + "`" + `small` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Optional) The instance type. Valid values: ` + "`" + `RDS` + "`" + `, ` + "`" + `PolarDB` + "`" + `, ` + "`" + `DDS` + "`" + `, ` + "`" + `Kvstore` + "`" + `, ` + "`" + `Other` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "database_region",
					Description: `(Optional) The region of the database.`,
				},
				resource.Attribute{
					Name:        "storage_region",
					Description: `(Optional) The storage region.`,
				},
				resource.Attribute{
					Name:        "from_app",
					Description: `(Optional) It is used to remark the request source. The default value is OpenAPI, and manual setting is unnecessary. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "backup_plan_id",
					Description: `The resource ID in terraform of Backup Plan. ## Import DBS Backup Plan can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_dbs_backup_plan.example <id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "backup_plan_id",
					Description: `The resource ID in terraform of Backup Plan. ## Import DBS Backup Plan can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_dbs_backup_plan.example <id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_disk",
			Category:         "ECS",
			ShortDescription: `Provides a ECS Disk resource.`,
			Description:      ``,
			Keywords: []string{
				"ecs",
				"disk",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required, ForceNew) The Zone to create the disk in.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the ECS disk. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the disk. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `(Optional, ForceNew) Category of the disk. Valid values are ` + "`" + `cloud` + "`" + `, ` + "`" + `cloud_efficiency` + "`" + `, ` + "`" + `cloud_ssd` + "`" + `. Default is ` + "`" + `cloud` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) The size of the disk in GiBs. When resize the disk, the new size must be greater than the former value, or you would get an error ` + "`" + `InvalidDiskSize.TooSmall` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) A snapshot to base the disk off of. If the disk size required by a snapshot is greater than ` + "`" + `size` + "`" + `, the ` + "`" + `size` + "`" + ` will be ignored, conflict with ` + "`" + `encrypted` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `(Optional) If true, the disk will be encrypted, conflict with ` + "`" + `snapshot_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "delete_auto_snapshot",
					Description: `(Optional) Indicates whether the automatic snapshot is deleted when the disk is released. Default value: false.`,
				},
				resource.Attribute{
					Name:        "delete_with_instance",
					Description: `(Optional) Indicates whether the disk is released together with the instance: Default value: false.`,
				},
				resource.Attribute{
					Name:        "enable_auto_snapshot",
					Description: `(Optional) Indicates whether to apply a created automatic snapshot policy to the disk. Default value: false. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the disk.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The disk status.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the disk.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The disk status.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_disk_attachment",
			Category:         "ECS",
			ShortDescription: `Provides a ECS Disk Attachment resource.`,
			Description:      ``,
			Keywords: []string{
				"ecs",
				"disk",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, Forces new resource) ID of the Instance to attach to.`,
				},
				resource.Attribute{
					Name:        "disk_id",
					Description: `(Required, Forces new resource) ID of the Disk to be attached.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `(Optional, Forces new resource) The device name which is used when attaching disk, it will be allocated automatically by system according to default order from /dev/xvdb to /dev/xvdz. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the Instance.`,
				},
				resource.Attribute{
					Name:        "disk_id",
					Description: `ID of the Disk.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `The device name exposed to the instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the Instance.`,
				},
				resource.Attribute{
					Name:        "disk_id",
					Description: `ID of the Disk.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `The device name exposed to the instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_dms_enterprise_instance",
			Category:         "DMS Enterprise",
			ShortDescription: `Provides a DMS Enterprise Instance resource.`,
			Description:      ``,
			Keywords: []string{
				"dms",
				"enterprise",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tid",
					Description: `(Optional) The tenant ID.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Required) Database type. Valid values: ` + "`" + `MySQL` + "`" + `, ` + "`" + `SQLServer` + "`" + `, ` + "`" + `PostgreSQL` + "`" + `, ` + "`" + `Oracle,` + "`" + ` ` + "`" + `DRDS` + "`" + `, ` + "`" + `OceanBase` + "`" + `, ` + "`" + `Mongo` + "`" + `, ` + "`" + `Redis` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_source",
					Description: `(Required) The source of the database instance. Valid values: ` + "`" + `PUBLIC_OWN` + "`" + `, ` + "`" + `RDS` + "`" + `, ` + "`" + `ECS_OWN` + "`" + `, ` + "`" + `VPC_IDC` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `(Required, ForceNew) Network type. Valid values: ` + "`" + `CLASSIC` + "`" + `, ` + "`" + `VPC` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "env_type",
					Description: `(Required) Environment type. Valid values: ` + "`" + `product` + "`" + ` production environment, ` + "`" + `dev` + "`" + ` development environment, ` + "`" + `pre` + "`" + ` pre-release environment, ` + "`" + `test` + "`" + ` test environment, ` + "`" + `sit` + "`" + ` SIT environment, ` + "`" + `uat` + "`" + ` UAT environment, ` + "`" + `pet` + "`" + ` pressure test environment, ` + "`" + `stag` + "`" + ` STAG environment.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Required, ForceNew) Host address of the target database.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required, ForceNew) Access port of the target database.`,
				},
				resource.Attribute{
					Name:        "database_user",
					Description: `(Required) Database access account.`,
				},
				resource.Attribute{
					Name:        "database_password",
					Description: `(Required) Database access password.`,
				},
				resource.Attribute{
					Name:        "instance_alias",
					Description: `It has been deprecated from provider version 1.100.0 and 'instance_name' instead.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Required) Instance name, to help users quickly distinguish positioning.`,
				},
				resource.Attribute{
					Name:        "dba_uid",
					Description: `(Required, ForceNew) The DBA of the instance is passed into the Alibaba Cloud uid of the DBA.`,
				},
				resource.Attribute{
					Name:        "safe_rule",
					Description: `(Required, ForceNew) The security rule of the instance is passed into the name of the security rule in the enterprise.`,
				},
				resource.Attribute{
					Name:        "query_timeout",
					Description: `(Required) Query timeout time, unit: s (seconds).`,
				},
				resource.Attribute{
					Name:        "export_timeout",
					Description: `(Required) Export timeout, unit: s (seconds).`,
				},
				resource.Attribute{
					Name:        "ecs_instance_id",
					Description: `(Optional, Computed) ECS instance ID. The value of InstanceSource is the ECS self-built library. This value must be passed.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) VPC ID. This value must be passed when the value of InstanceSource is VPC dedicated line IDC.`,
				},
				resource.Attribute{
					Name:        "ecs_region",
					Description: `(Optional) The region where the instance is located. This value must be passed when the value of InstanceSource is RDS, ECS self-built library, and VPC dedicated line IDC.`,
				},
				resource.Attribute{
					Name:        "sid",
					Description: `(Optional) The SID. This value must be passed when InstanceType is PostgreSQL or Oracle.`,
				},
				resource.Attribute{
					Name:        "data_link_name",
					Description: `(Optional) Cross-database query datalink name.`,
				},
				resource.Attribute{
					Name:        "ddl_online",
					Description: `(Optional) Whether to use online services, currently only supports MySQL and PolarDB. Valid values: ` + "`" + `0` + "`" + ` Not used, ` + "`" + `1` + "`" + ` Native online DDL priority, ` + "`" + `2` + "`" + ` DMS lock-free table structure change priority.`,
				},
				resource.Attribute{
					Name:        "use_dsql",
					Description: `(Optional) Whether to enable cross-instance query. Valid values: ` + "`" + `0` + "`" + ` not open, ` + "`" + `1` + "`" + ` open. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the DMS enterprise instance and format as ` + "`" + `<host>:<port>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dba_nick_name",
					Description: `The instance dba nickname.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `It has been deprecated from provider version 1.100.0 and 'status' instead.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The instance status. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 1 mins) Used when creating the DMS enterprise instance. ## Import DMS Enterprise can be imported using host and port, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_dms_enterprise_instance.example rm-uf648hgs7874xxxx.mysql.rds.aliyuncs.com:3306 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the DMS enterprise instance and format as ` + "`" + `<host>:<port>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dba_nick_name",
					Description: `The instance dba nickname.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `It has been deprecated from provider version 1.100.0 and 'status' instead.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The instance status. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 1 mins) Used when creating the DMS enterprise instance. ## Import DMS Enterprise can be imported using host and port, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_dms_enterprise_instance.example rm-uf648hgs7874xxxx.mysql.rds.aliyuncs.com:3306 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_dms_enterprise_user",
			Category:         "DMS Enterprise",
			ShortDescription: `Provides a DMS Enterprise User resource.`,
			Description:      ``,
			Keywords: []string{
				"dms",
				"enterprise",
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tid",
					Description: `(Optional) The tenant ID.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Required, ForceNew) The Alibaba Cloud unique ID (UID) of the user to add.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The state of DMS Enterprise User. Valid values: ` + "`" + `NORMAL` + "`" + `, ` + "`" + `DISABLE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "role_names",
					Description: `(Optional) The roles that the user plays.`,
				},
				resource.Attribute{
					Name:        "nick_name",
					Description: `(Optional) It has been deprecated and use ` + "`" + `user_name` + "`" + ` instead.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `(Optional) The nickname of the user.`,
				},
				resource.Attribute{
					Name:        "mobile",
					Description: `(Optional) The DingTalk number or mobile number of the user.`,
				},
				resource.Attribute{
					Name:        "max_result_count",
					Description: `(Optional) Query the maximum number of rows on the day.`,
				},
				resource.Attribute{
					Name:        "max_execute_count",
					Description: `(Optional) Maximum number of inquiries on the day. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Alibaba Cloud unique ID of the user. The value is same as the UID.`,
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
					Name:        "role_names",
					Description: `The list of roles that the user plays.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The state of DMS Enterprise User. ## Import DMS Enterprise User can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_dms_enterprise_user.example 24356xxx ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Alibaba Cloud unique ID of the user. The value is same as the UID.`,
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
					Name:        "role_names",
					Description: `The list of roles that the user plays.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The state of DMS Enterprise User. ## Import DMS Enterprise User can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_dms_enterprise_user.example 24356xxx ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_dns_domain",
			Category:         "DNS",
			ShortDescription: `Provides a DNS domain resource.`,
			Description:      ``,
			Keywords: []string{
				"dns",
				"domain",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain_name",
					Description: `(Required, ForceNew) Name of the domain. This name without suffix can have a string of 1 to 63 characters(domain name subject, excluding suffix), must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix ` + "`" + `.sh` + "`" + ` and ` + "`" + `.tel` + "`" + ` are not supported.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Optional) Id of the group in which the domain will add. If not supplied, then use default group.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional, ForceNew) The Id of resource group which the dns domain belongs.`,
				},
				resource.Attribute{
					Name:        "lang",
					Description: `(Optional) User language.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `(Optional) Remarks information for your domain name. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `This ID of this resource. The value is set to ` + "`" + `domain_name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `The domain ID.`,
				},
				resource.Attribute{
					Name:        "dns_server",
					Description: `A list of the dns server name.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `This ID of this resource. The value is set to ` + "`" + `domain_name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `The domain ID.`,
				},
				resource.Attribute{
					Name:        "dns_server",
					Description: `A list of the dns server name.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_dns_domain_attachment",
			Category:         "DNS",
			ShortDescription: `Provides bind the domain name to the DNS instance resource.`,
			Description:      ``,
			Keywords: []string{
				"dns",
				"domain",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, ForceNew) The id of the DNS instance.`,
				},
				resource.Attribute{
					Name:        "domain_names",
					Description: `(Required) The domain names bound to the DNS instance. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `This ID of this resource. The value is same as ` + "`" + `instance_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "domain_names",
					Description: `Domain names bound to DNS instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `This ID of this resource. The value is same as ` + "`" + `instance_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "domain_names",
					Description: `Domain names bound to DNS instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_dns_group",
			Category:         "DNS",
			ShortDescription: `Provides a DNS Group resource.`,
			Description:      ``,
			Keywords: []string{
				"dns",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the domain group. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The group id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The group name.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The group id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The group name.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_dns_record",
			Category:         "DNS",
			ShortDescription: `Provides a DNS Record resource.`,
			Description:      ``,
			Keywords: []string{
				"dns",
				"record",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Required) ID of the Dns Domain where this record belongs.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Host record for the domain record. This host_record can have at most 253 characters, and each part split with "." can have at most 63 characters, and must contain only alphanumeric characters or hyphens, such as "-",".","`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of domain record. Valid values are ` + "`" + `A` + "`" + `,` + "`" + `NS` + "`" + `,` + "`" + `MX` + "`" + `,` + "`" + `TXT` + "`" + `,` + "`" + `CNAME` + "`" + `,` + "`" + `SRV` + "`" + `,` + "`" + `AAAA` + "`" + `,` + "`" + `CAA` + "`" + `, ` + "`" + `REDIRECT_URL` + "`" + ` and ` + "`" + `FORWORD_URL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rr_set",
					Description: `(Optional) The value of domain record, When the ` + "`" + `type` + "`" + ` is ` + "`" + `MX` + "`" + `,` + "`" + `NS` + "`" + `,` + "`" + `CNAME` + "`" + `,` + "`" + `SRV` + "`" + `, the server will treat the ` + "`" + `value` + "`" + ` as a fully qualified domain name, so it's no need to add a ` + "`" + `.` + "`" + ` at the end.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The effective time of domain record. Its scope depends on the edition of the cloud resolution. Free is ` + "`" + `[600, 86400]` + "`" + `, Basic is ` + "`" + `[120, 86400]` + "`" + `, Standard is ` + "`" + `[60, 86400]` + "`" + `, Ultimate is ` + "`" + `[10, 86400]` + "`" + `, Exclusive is ` + "`" + `[1, 86400]` + "`" + `. Default value is ` + "`" + `300` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `(Optional) The effective time of domain record. Its scope depends on the edition of the cloud resolution. Free is ` + "`" + `[600, 86400]` + "`" + `, Basic is ` + "`" + `[120, 86400]` + "`" + `, Standard is ` + "`" + `[60, 86400]` + "`" + `, Ultimate is ` + "`" + `[10, 86400]` + "`" + `, Exclusive is ` + "`" + `[1, 86400]` + "`" + `. Default value is ` + "`" + `300` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "record_id",
					Description: `ID of the Dns Record.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The record type.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The host record of record.`,
				},
				resource.Attribute{
					Name:        "rr_set",
					Description: `The record value.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `The record effective time.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `ID of the Dns Domain where this record belongs.`,
				},
				resource.Attribute{
					Name:        "lba_strategy",
					Description: `The record strategy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "record_id",
					Description: `ID of the Dns Record.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The record type.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The host record of record.`,
				},
				resource.Attribute{
					Name:        "rr_set",
					Description: `The record value.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `The record effective time.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `ID of the Dns Domain where this record belongs.`,
				},
				resource.Attribute{
					Name:        "lba_strategy",
					Description: `The record strategy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_drds_instance",
			Category:         "Distributed Relational Database Service (DRDS)",
			ShortDescription: `Provides an DRDS instance resource.`,
			Description:      ``,
			Keywords: []string{
				"distributed",
				"relational",
				"database",
				"service",
				"drds",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Required) Description of the DRDS instance, This description can have a string of 2 to 256 characters.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Required, ForceNew) The Zone to launch the DRDS instance.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `(Optional, ForceNew) Valid values are ` + "`" + `PrePaid` + "`" + `, ` + "`" + `PostPaid` + "`" + `, Default to ` + "`" + `PostPaid` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Required, ForceNew) The VSwitch ID to launch in.`,
				},
				resource.Attribute{
					Name:        "instance_series",
					Description: `(Required, ForceNew) User-defined DRDS instance node spec. Value range: - ` + "`" + `drds.sn1.4c8g` + "`" + ` for DRDS instance Starter version; - ` + "`" + `drds.sn1.8c16g` + "`" + ` for DRDS instance Standard edition; - ` + "`" + `drds.sn1.16c32g` + "`" + ` for DRDS instance Enterprise Edition; - ` + "`" + `drds.sn1.32c64g` + "`" + ` for DRDS instance Extreme Edition;`,
				},
				resource.Attribute{
					Name:        "specification",
					Description: `(Required, ForceNew) User-defined DRDS instance specification. Value range: - ` + "`" + `drds.sn1.4c8g` + "`" + ` for DRDS instance Starter version; - value range : ` + "`" + `drds.sn1.4c8g.8c16g` + "`" + `, ` + "`" + `drds.sn1.4c8g.16c32g` + "`" + `, ` + "`" + `drds.sn1.4c8g.32c64g` + "`" + `, ` + "`" + `drds.sn1.4c8g.64c128g` + "`" + ` - ` + "`" + `drds.sn1.8c16g` + "`" + ` for DRDS instance Standard edition; - value range : ` + "`" + `drds.sn1.8c16g.16c32g` + "`" + `, ` + "`" + `drds.sn1.8c16g.32c64g` + "`" + `, ` + "`" + `drds.sn1.8c16g.64c128g` + "`" + ` - ` + "`" + `drds.sn1.16c32g` + "`" + ` for DRDS instance Enterprise Edition; - value range : ` + "`" + `drds.sn1.16c32g.32c64g` + "`" + `, ` + "`" + `drds.sn1.16c32g.64c128g` + "`" + ` - ` + "`" + `drds.sn1.32c64g` + "`" + ` for DRDS instance Extreme Edition; - value range : ` + "`" + `drds.sn1.32c64g.128c256g` + "`" + ` ### Timeouts ->`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 10 mins) Used when creating the drds instance (until it reaches running status).`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 10 mins) Used when terminating the drds instance. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The DRDS instance ID. ## Import Distributed Relational Database Service (DRDS) can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_drds_instance.example drds-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The DRDS instance ID. ## Import Distributed Relational Database Service (DRDS) can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_drds_instance.example drds-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_dts_subscription_job",
			Category:         "Data Transmission Service (DTS)",
			ShortDescription: `Provides a Alibabacloudstack DTS Subscription Job resource.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"transmission",
				"service",
				"dts",
				"subscription",
				"job",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dts_instance_id",
					Description: `(Computed, ForceNew) The ID of subscription instance.`,
				},
				resource.Attribute{
					Name:        "dts_job_name",
					Description: `(Optional) The name of subscription task.`,
				},
				resource.Attribute{
					Name:        "checkpoint",
					Description: `(Optional, OtherParam) Subscription start time in Unix timestamp format.`,
				},
				resource.Attribute{
					Name:        "compute_unit",
					Description: `(Optional, OtherParam) [ETL specifications](https://help.aliyun.com/document_detail/212324.html). The unit is the computing unit ComputeUnit (CU), 1CU=1vCPU+4 GB memory. The value range is an integer greater than or equal to 2.`,
				},
				resource.Attribute{
					Name:        "database_count",
					Description: `(Optional, OtherParam) The number of private customized RDS instances under PolarDB-X. The default value is 1. This parameter needs to be passed only when ` + "`" + `source_endpoint_engine_name` + "`" + ` equals ` + "`" + `drds` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "db_list",
					Description: `(Optional) Subscription object, in the format of JSON strings. For detailed definitions, please refer to the description of migration, synchronization or subscription objects [document](https://help.aliyun.com/document_detail/209545.html).`,
				},
				resource.Attribute{
					Name:        "delay_notice",
					Description: `(Optional) This parameter decides whether to monitor the delay status. Valid values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "delay_phone",
					Description: `(Optional) The mobile phone number of the contact who delayed the alarm. Multiple mobile phone numbers separated by English commas ` + "`" + `,` + "`" + `. This parameter currently only supports China stations, and only supports mainland mobile phone numbers, and up to 10 mobile phone numbers can be passed in.`,
				},
				resource.Attribute{
					Name:        "delay_rule_time",
					Description: `(Optional) When ` + "`" + `delay_notice` + "`" + ` is set to ` + "`" + `true` + "`" + `, this parameter must be passed in. The threshold for triggering the delay alarm. The unit is second and needs to be an integer. The threshold can be set according to business needs. It is recommended to set it above 10 seconds to avoid delay fluctuations caused by network and database load.`,
				},
				resource.Attribute{
					Name:        "destination_endpoint_engine_name",
					Description: `(Optional) The destination endpoint engine name. Valid values: ` + "`" + `ADS` + "`" + `, ` + "`" + `DB2` + "`" + `, ` + "`" + `DRDS` + "`" + `, ` + "`" + `DataHub` + "`" + `, ` + "`" + `Greenplum` + "`" + `, ` + "`" + `MSSQL` + "`" + `, ` + "`" + `MySQL` + "`" + `, ` + "`" + `PolarDB` + "`" + `, ` + "`" + `PostgreSQL` + "`" + `, ` + "`" + `Redis` + "`" + `, ` + "`" + `Tablestore` + "`" + `, ` + "`" + `as400` + "`" + `, ` + "`" + `clickhouse` + "`" + `, ` + "`" + `kafka` + "`" + `, ` + "`" + `mongodb` + "`" + `, ` + "`" + `odps` + "`" + `, ` + "`" + `oracle` + "`" + `, ` + "`" + `polardb_o` + "`" + `, ` + "`" + `polardb_pg` + "`" + `, ` + "`" + `tidb` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "destination_region",
					Description: `(Optional) The destination region. List of [supported regions](https://help.aliyun.com/document_detail/141033.html).`,
				},
				resource.Attribute{
					Name:        "error_notice",
					Description: `(Optional) This parameter decides whether to monitor abnormal status. Valid values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "error_phone",
					Description: `(Optional) The mobile phone number of the contact for abnormal alarm. Multiple mobile phone numbers separated by English commas ` + "`" + `,` + "`" + `. This parameter currently only supports China stations, and only supports mainland mobile phone numbers, and up to 10 mobile phone numbers can be passed in.`,
				},
				resource.Attribute{
					Name:        "instance_class",
					Description: `(Optional) The instance class. Valid values: ` + "`" + `large` + "`" + `, ` + "`" + `medium` + "`" + `, ` + "`" + `micro` + "`" + `, ` + "`" + `small` + "`" + `, ` + "`" + `xlarge` + "`" + `, ` + "`" + `xxlarge` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "payment_type",
					Description: `(Optional, ForceNew) The payment type of the resource. Valid values: ` + "`" + `Subscription` + "`" + `, ` + "`" + `PayAsYouGo` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "payment_duration_unit",
					Description: `(Optional) The payment duration unit. Valid values: ` + "`" + `Month` + "`" + `, ` + "`" + `Year` + "`" + `. When ` + "`" + `payment_type` + "`" + ` is ` + "`" + `Subscription` + "`" + `, this parameter is valid and must be passed in.`,
				},
				resource.Attribute{
					Name:        "payment_duration",
					Description: `(Optional) The duration of prepaid instance purchase. When ` + "`" + `payment_type` + "`" + ` is ` + "`" + `Subscription` + "`" + `, this parameter is valid and must be passed in.`,
				},
				resource.Attribute{
					Name:        "reserve",
					Description: `(Optional) DTS reserves parameters, the format is a JSON string, you can pass in this parameter to complete the source and target database information (such as the data storage format of the target Kafka database, the instance ID of the cloud enterprise network CEN). For more information, please refer to the parameter description of the [Reserve parameter](https://help.aliyun.com/document_detail/176470.html).`,
				},
				resource.Attribute{
					Name:        "source_endpoint_database_name",
					Description: `(Optional) To subscribe to the name of the database.`,
				},
				resource.Attribute{
					Name:        "source_endpoint_engine_name",
					Description: `(Optional) The source database type value is MySQL or Oracle. Valid values: ` + "`" + `MySQL` + "`" + `, ` + "`" + `Oracle` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "source_endpoint_instance_id",
					Description: `(Optional) The ID of source instance. Only when the type of source database instance was RDS MySQL, PolarDB-X 1.0, PolarDB MySQL, this parameter can be available and must be set.`,
				},
				resource.Attribute{
					Name:        "source_endpoint_instance_type",
					Description: `(Optional, ForceNew) The type of source instance. Valid values: ` + "`" + `RDS` + "`" + `, ` + "`" + `PolarDB` + "`" + `, ` + "`" + `DRDS` + "`" + `, ` + "`" + `LocalInstance` + "`" + `, ` + "`" + `ECS` + "`" + `, ` + "`" + `Express` + "`" + `, ` + "`" + `CEN` + "`" + `, ` + "`" + `dg` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "source_endpoint_ip",
					Description: `(Optional) The IP of source endpoint.`,
				},
				resource.Attribute{
					Name:        "source_endpoint_oracle_sid",
					Description: `(Optional) The SID of Oracle Database. When the source database is self-built Oracle and the Oracle database is a non-RAC instance, this parameter is available and must be passed in.`,
				},
				resource.Attribute{
					Name:        "source_endpoint_owner_id",
					Description: `(Optional) The Alibaba Cloud account ID to which the source instance belongs. This parameter is only available when configuring data subscriptions across Alibaba Cloud accounts and must be passed in.`,
				},
				resource.Attribute{
					Name:        "source_endpoint_user_name",
					Description: `(Optional) The username of source database instance account.`,
				},
				resource.Attribute{
					Name:        "source_endpoint_password",
					Description: `(Optional) The password of source database instance account.`,
				},
				resource.Attribute{
					Name:        "source_endpoint_port",
					Description: `(Optional) The port of source database.`,
				},
				resource.Attribute{
					Name:        "source_endpoint_region",
					Description: `(Optional) The region of source database.`,
				},
				resource.Attribute{
					Name:        "source_endpoint_role",
					Description: `(Optional) Both the authorization roles. When the source instance and configure subscriptions task of the Alibaba Cloud account is not the same as the need to pass the parameter, to specify the source of the authorization roles, to allow configuration subscription task of the Alibaba Cloud account to access the source of the source instance information.`,
				},
				resource.Attribute{
					Name:        "subscription_data_type_ddl",
					Description: `(Optional, Computed) Whether to subscribe the DDL type of data. Valid values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "subscription_data_type_dml",
					Description: `(Optional, Computed) Whether to subscribe the DML type of data. Valid values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "subscription_instance_network_type",
					Description: `(Optional) Subscription task type of network value: classic: classic Network. Virtual Private Cloud (vpc): a vpc. Valid values: ` + "`" + `classic` + "`" + `, ` + "`" + `vpc` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "subscription_instance_vpc_id",
					Description: `(Optional) The ID of subscription vpc instance. When the value of ` + "`" + `subscription_instance_network_type` + "`" + ` is vpc, this parameter is available and must be passed in.`,
				},
				resource.Attribute{
					Name:        "subscription_instance_vswitch_id",
					Description: `(Optional) The ID of subscription VSwitch instance. When the value of ` + "`" + `subscription_instance_network_type` + "`" + ` is vpc, this parameter is available and must be passed in.`,
				},
				resource.Attribute{
					Name:        "sync_architecture",
					Description: `(Optional) The sync architecture. Valid values: ` + "`" + `bidirectional` + "`" + `, ` + "`" + `oneway` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "synchronization_direction",
					Description: `(Optional) The synchronization direction. Valid values: ` + "`" + `Forward` + "`" + `, ` + "`" + `Reverse` + "`" + `. When the topology type of the data synchronization instance is bidirectional, it can be passed in to reverse to start the reverse synchronization link.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional, Computed) The status of the task. Valid values: ` + "`" + `Normal` + "`" + `, ` + "`" + `Abnormal` + "`" + `. When a task created, it is in this state of ` + "`" + `NotStarted` + "`" + `. You can specify this state to ` + "`" + `Normal` + "`" + ` to start the job, and specify this state of ` + "`" + `Abnormal` + "`" + ` to stop the job.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in terraform of Subscription Job. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 1 mins) Used when update the Subscription Job. ## Import DTS Subscription Job can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_dts_subscription_job.example <id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in terraform of Subscription Job. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 1 mins) Used when update the Subscription Job. ## Import DTS Subscription Job can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_dts_subscription_job.example <id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_dts_synchronization_instance",
			Category:         "Data Transmission Service (DTS)",
			ShortDescription: `Provides a Alibabacloudstack DTS Synchronization Instance resource.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"transmission",
				"service",
				"dts",
				"synchronization",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_class",
					Description: `(Required) The instance class. Valid values: ` + "`" + `large` + "`" + `, ` + "`" + `medium` + "`" + `, ` + "`" + `micro` + "`" + `, ` + "`" + `small` + "`" + `, ` + "`" + `xlarge` + "`" + `, ` + "`" + `xxlarge` + "`" + `. You can only upgrade the configuration, not downgrade the configuration. If you downgrade the instance, you need to [submit a ticket](https://selfservice.console.aliyun.com/ticket/category/dts/today).`,
				},
				resource.Attribute{
					Name:        "payment_type",
					Description: `(Required, ForceNew) The payment type of the resource. Valid values: ` + "`" + `Subscription` + "`" + `, ` + "`" + `PayAsYouGo` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "payment_duration_unit",
					Description: `(Optional) The payment duration unit. Valid values: ` + "`" + `Month` + "`" + `, ` + "`" + `Year` + "`" + `. When ` + "`" + `payment_type` + "`" + ` is ` + "`" + `Subscription` + "`" + `, this parameter is valid and must be passed in.`,
				},
				resource.Attribute{
					Name:        "payment_duration",
					Description: `(Required when ` + "`" + `payment_type` + "`" + ` equals ` + "`" + `Subscription` + "`" + `) The duration of prepaid instance purchase. When ` + "`" + `payment_type` + "`" + ` is ` + "`" + `Subscription` + "`" + `, this parameter is valid and must be passed in.`,
				},
				resource.Attribute{
					Name:        "source_endpoint_region",
					Description: `(Required, ForceNew) The region of source instance.`,
				},
				resource.Attribute{
					Name:        "source_endpoint_engine_name",
					Description: `(Required, ForceNew) The type of source endpoint engine. Valid values: ` + "`" + `ADS` + "`" + `, ` + "`" + `DB2` + "`" + `, ` + "`" + `DRDS` + "`" + `, ` + "`" + `DataHub` + "`" + `, ` + "`" + `Greenplum` + "`" + `, ` + "`" + `MSSQL` + "`" + `, ` + "`" + `MySQL` + "`" + `, ` + "`" + `PolarDB` + "`" + `, ` + "`" + `PostgreSQL` + "`" + `, ` + "`" + `Redis` + "`" + `, ` + "`" + `Tablestore` + "`" + `, ` + "`" + `as400` + "`" + `, ` + "`" + `clickhouse` + "`" + `, ` + "`" + `kafka` + "`" + `, ` + "`" + `mongodb` + "`" + `, ` + "`" + `odps` + "`" + `, ` + "`" + `oracle` + "`" + `, ` + "`" + `polardb_o` + "`" + `, ` + "`" + `polardb_pg` + "`" + `, ` + "`" + `tidb` + "`" + `. For the correspondence between the supported source and target libraries, see [Supported Databases, Synchronization Initialization Types and Synchronization Topologies](https://help.aliyun.com/document_detail/130744.html), [Supported Databases and Migration Types](https://help.aliyun.com/document_detail/26618.html).`,
				},
				resource.Attribute{
					Name:        "destination_endpoint_region",
					Description: `(Required, ForceNew) The region of destination instance. List of [supported regions](https://help.aliyun.com/document_detail/141033.html).`,
				},
				resource.Attribute{
					Name:        "destination_endpoint_engine_name",
					Description: `(Required, ForceNew) The type of destination engine. Valid values: ` + "`" + `ADS` + "`" + `, ` + "`" + `DB2` + "`" + `, ` + "`" + `DRDS` + "`" + `, ` + "`" + `DataHub` + "`" + `, ` + "`" + `Greenplum` + "`" + `, ` + "`" + `MSSQL` + "`" + `, ` + "`" + `MySQL` + "`" + `, ` + "`" + `PolarDB` + "`" + `, ` + "`" + `PostgreSQL` + "`" + `, ` + "`" + `Redis` + "`" + `, ` + "`" + `Tablestore` + "`" + `, ` + "`" + `as400` + "`" + `, ` + "`" + `clickhouse` + "`" + `, ` + "`" + `kafka` + "`" + `, ` + "`" + `mongodb` + "`" + `, ` + "`" + `odps` + "`" + `, ` + "`" + `oracle` + "`" + `, ` + "`" + `polardb_o` + "`" + `, ` + "`" + `polardb_pg` + "`" + `, ` + "`" + `tidb` + "`" + `. For the correspondence between the supported source and target libraries, see [Supported Databases, Synchronization Initialization Types and Synchronization Topologies](https://help.aliyun.com/document_detail/130744.html), [Supported Databases and Migration Types](https://help.aliyun.com/document_detail/26618.html).`,
				},
				resource.Attribute{
					Name:        "sync_architecture",
					Description: `(Optional, ForceNew) The sync architecture. Valid values: ` + "`" + `oneway` + "`" + `, ` + "`" + `bidirectional` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "compute_unit",
					Description: `(Optional) [ETL specifications](https://help.aliyun.com/document_detail/212324.html). The unit is the computing unit ComputeUnit (CU), 1CU=1vCPU+4 GB memory. The value range is an integer greater than or equal to 2.`,
				},
				resource.Attribute{
					Name:        "database_count",
					Description: `(Optional) The number of private customized RDS instances under PolarDB-X. The default value is 1. This parameter needs to be passed only when ` + "`" + `source_endpoint_engine_name` + "`" + ` equals ` + "`" + `drds` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "quantity",
					Description: `(Optional) The number of instances purchased. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID of Synchronization Instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status. ## Import DTS Synchronization Instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_dts_synchronization_instance.example <id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID of Synchronization Instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status. ## Import DTS Synchronization Instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_dts_synchronization_instance.example <id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_dts_synchronization_job",
			Category:         "Data Transmission Service (DTS)",
			ShortDescription: `Provides a Alibabacloudstack DTS Synchronization Job resource.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"transmission",
				"service",
				"dts",
				"synchronization",
				"job",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dts_instance_id",
					Description: `(Required, ForceNew) Synchronizing instance ID. The ID of ` + "`" + `alibabacloudstack_dts_synchronization_instance` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "synchronization_direction",
					Description: `(Optional, ForceNew) Synchronization direction. Valid values: ` + "`" + `Forward` + "`" + `, ` + "`" + `Reverse` + "`" + `. Only when the property ` + "`" + `sync_architecture` + "`" + ` of the ` + "`" + `alibabacloudstack_dts_synchronization_instance` + "`" + ` was ` + "`" + `bidirectional` + "`" + ` this parameter should be passed, otherwise this parameter should not be specified.`,
				},
				resource.Attribute{
					Name:        "dts_job_name",
					Description: `(Optional, Computed) The name of synchronization job.`,
				},
				resource.Attribute{
					Name:        "instance_class",
					Description: `(Optional, Computed) The instance class. Valid values: ` + "`" + `large` + "`" + `, ` + "`" + `medium` + "`" + `, ` + "`" + `micro` + "`" + `, ` + "`" + `small` + "`" + `, ` + "`" + `xlarge` + "`" + `, ` + "`" + `xxlarge` + "`" + `. You can only upgrade the configuration, not downgrade the configuration. If you downgrade the instance, you need to [submit a ticket](https://selfservice.console.aliyun.com/ticket/category/dts/today).`,
				},
				resource.Attribute{
					Name:        "checkpoint",
					Description: `(Optional, Computed, ForceNew) Start time in Unix timestamp format.`,
				},
				resource.Attribute{
					Name:        "data_initialization",
					Description: `(Required, ForceNew) Whether or not to execute DTS supports schema migration, full data migration, or full-data initialization values include:`,
				},
				resource.Attribute{
					Name:        "data_synchronization",
					Description: `(Required, ForceNew) Whether to perform incremental data migration for migration types or synchronization values include:`,
				},
				resource.Attribute{
					Name:        "structure_initialization",
					Description: `(Required, ForceNew) Whether to perform a database table structure to migrate or initialization values include:`,
				},
				resource.Attribute{
					Name:        "db_list",
					Description: `(Required, ForceNew) Migration object, in the format of JSON strings. For detailed definition instructions, please refer to [the description of migration, synchronization or subscription objects](https://help.aliyun.com/document_detail/209545.html).`,
				},
				resource.Attribute{
					Name:        "reserve",
					Description: `(Optional, ForceNew) DTS reserves parameters, the format is a JSON string, you can pass in this parameter to complete the source and target database information (such as the data storage format of the target Kafka database, the instance ID of the cloud enterprise network CEN). For more information, please refer to the parameter [description of the Reserve parameter](https://help.aliyun.com/document_detail/273111.html).`,
				},
				resource.Attribute{
					Name:        "source_endpoint_instance_type",
					Description: `(Required, ForceNew) The type of source instance. Valid values: ` + "`" + `CEN` + "`" + `, ` + "`" + `DG` + "`" + `, ` + "`" + `DISTRIBUTED_DMSLOGICDB` + "`" + `, ` + "`" + `ECS` + "`" + `, ` + "`" + `EXPRESS` + "`" + `, ` + "`" + `MONGODB` + "`" + `, ` + "`" + `OTHER` + "`" + `, ` + "`" + `PolarDB` + "`" + `, ` + "`" + `POLARDBX20` + "`" + `, ` + "`" + `RDS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "source_endpoint_engine_name",
					Description: `(Required, ForceNew) The type of source database. Valid values: ` + "`" + `AS400` + "`" + `, ` + "`" + `DB2` + "`" + `, ` + "`" + `DMSPOLARDB` + "`" + `, ` + "`" + `HBASE` + "`" + `, ` + "`" + `MONGODB` + "`" + `, ` + "`" + `MSSQL` + "`" + `, ` + "`" + `MySQL` + "`" + `, ` + "`" + `ORACLE` + "`" + `, ` + "`" + `PolarDB` + "`" + `, ` + "`" + `POLARDBX20` + "`" + `, ` + "`" + `POLARDB_O` + "`" + `, ` + "`" + `POSTGRESQL` + "`" + `, ` + "`" + `TERADATA` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "source_endpoint_instance_id",
					Description: `(Optional, ForceNew) The ID of source instance.`,
				},
				resource.Attribute{
					Name:        "source_endpoint_region",
					Description: `(Optional, ForceNew) The region of source instance.`,
				},
				resource.Attribute{
					Name:        "source_endpoint_ip",
					Description: `(Optional, ForceNew) The ip of source endpoint.`,
				},
				resource.Attribute{
					Name:        "source_endpoint_port",
					Description: `(Optional, ForceNew) The port of source endpoint.`,
				},
				resource.Attribute{
					Name:        "source_endpoint_oracle_sid",
					Description: `(Optional, ForceNew) The SID of Oracle database.`,
				},
				resource.Attribute{
					Name:        "source_endpoint_database_name",
					Description: `(Optional, ForceNew) The name of migrate the database.`,
				},
				resource.Attribute{
					Name:        "source_endpoint_user_name",
					Description: `(Optional, ForceNew) The username of database account.`,
				},
				resource.Attribute{
					Name:        "source_endpoint_password",
					Description: `(Optional) The password of database account.`,
				},
				resource.Attribute{
					Name:        "source_endpoint_owner_id",
					Description: `(Optional, ForceNew) The Alibaba Cloud account ID to which the source instance belongs.`,
				},
				resource.Attribute{
					Name:        "source_endpoint_role",
					Description: `(Optional, ForceNew) The name of the role configured for the cloud account to which the source instance belongs.`,
				},
				resource.Attribute{
					Name:        "destination_endpoint_instance_type",
					Description: `(Required, ForceNew) The type of destination instance. Valid values: ` + "`" + `ads` + "`" + `, ` + "`" + `CEN` + "`" + `, ` + "`" + `DATAHUB` + "`" + `, ` + "`" + `DG` + "`" + `, ` + "`" + `ECS` + "`" + `, ` + "`" + `EXPRESS` + "`" + `, ` + "`" + `GREENPLUM` + "`" + `, ` + "`" + `MONGODB` + "`" + `, ` + "`" + `OTHER` + "`" + `, ` + "`" + `PolarDB` + "`" + `, ` + "`" + `POLARDBX20` + "`" + `, ` + "`" + `RDS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "destination_endpoint_engine_name",
					Description: `(Required, ForceNew) The type of destination database. Valid values: ` + "`" + `ADB20` + "`" + `, ` + "`" + `ADB30` + "`" + `, ` + "`" + `AS400` + "`" + `, ` + "`" + `DATAHUB` + "`" + `, ` + "`" + `DB2` + "`" + `, ` + "`" + `GREENPLUM` + "`" + `, ` + "`" + `KAFKA` + "`" + `, ` + "`" + `MONGODB` + "`" + `, ` + "`" + `MSSQL` + "`" + `, ` + "`" + `MySQL` + "`" + `, ` + "`" + `ORACLE` + "`" + `, ` + "`" + `PolarDB` + "`" + `, ` + "`" + `POLARDBX20` + "`" + `, ` + "`" + `POLARDB_O` + "`" + `, ` + "`" + `PostgreSQL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "destination_endpoint_instance_id",
					Description: `(Optional, ForceNew) The ID of destination instance.`,
				},
				resource.Attribute{
					Name:        "destination_endpoint_region",
					Description: `(Optional, ForceNew) The region of destination instance.`,
				},
				resource.Attribute{
					Name:        "destination_endpoint_ip",
					Description: `(Optional, ForceNew) The ip of source endpoint.`,
				},
				resource.Attribute{
					Name:        "destination_endpoint_port",
					Description: `(Optional, ForceNew) The port of source endpoint.`,
				},
				resource.Attribute{
					Name:        "destination_endpoint_data_base_name",
					Description: `(Optional, ForceNew) The name of migrate the database.`,
				},
				resource.Attribute{
					Name:        "destination_endpoint_user_name",
					Description: `(Optional, ForceNew) The username of database account.`,
				},
				resource.Attribute{
					Name:        "destination_endpoint_password",
					Description: `(Optional, ForceNew) The password of database account.`,
				},
				resource.Attribute{
					Name:        "destination_endpoint_oracle_sid",
					Description: `(Optional, ForceNew) The SID of Oracle database.`,
				},
				resource.Attribute{
					Name:        "delay_notice",
					Description: `(Optional, ForceNew) The delay notice. Valid values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "delay_phone",
					Description: `(Optional, ForceNew) The delay phone. The mobile phone number of the contact who delayed the alarm. Multiple mobile phone numbers separated by English commas ` + "`" + `,` + "`" + `. This parameter currently only supports China stations, and only supports mainland mobile phone numbers, and up to 10 mobile phone numbers can be passed in.`,
				},
				resource.Attribute{
					Name:        "delay_rule_time",
					Description: `(Optional, ForceNew) The delay rule time. When ` + "`" + `delay_notice` + "`" + ` is set to ` + "`" + `true` + "`" + `, this parameter must be passed in. The threshold for triggering the delay alarm. The unit is second and needs to be an integer. The threshold can be set according to business needs. It is recommended to set it above 10 seconds to avoid delay fluctuations caused by network and database load.`,
				},
				resource.Attribute{
					Name:        "error_notice",
					Description: `(Optional, ForceNew) The error notice. Valid values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "error_phone",
					Description: `(Optional, ForceNew) The error phone. The mobile phone number of the contact who error the alarm. Multiple mobile phone numbers separated by English commas ` + "`" + `,` + "`" + `. This parameter currently only supports China stations, and only supports mainland mobile phone numbers, and up to 10 mobile phone numbers can be passed in.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional, Computed) The status of the resource. Valid values: ` + "`" + `Synchronizing` + "`" + `, ` + "`" + `Suspending` + "`" + `. You can stop the task by specifying ` + "`" + `Suspending` + "`" + ` and start the task by specifying ` + "`" + `Synchronizing` + "`" + `. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in terraform of Synchronization Job. ## Import DTS Synchronization Job can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_dts_synchronization_job.example <id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in terraform of Synchronization Job. ## Import DTS Synchronization Job can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_dts_synchronization_job.example <id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ecs_command",
			Category:         "ECS",
			ShortDescription: `Provides a Alibabacloudstack ECS Command resource.`,
			Description:      ``,
			Keywords: []string{
				"ecs",
				"command",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "command_content",
					Description: `(Required, ForceNew) The Base64-encoded content of the command.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, ForceNew) The description of command.`,
				},
				resource.Attribute{
					Name:        "enable_parameter",
					Description: `(Optional, ForceNew) Specifies whether to use custom parameters in the command to be created. Default to: false.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) The name of the command, which supports all character sets. It can be up to 128 characters in length.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional, ForceNew) The timeout period that is specified for the command to be run on ECS instances. Unit: seconds. Default to: ` + "`" + `60` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required, ForceNew) The command type. Valid Values: ` + "`" + `RunBatScript` + "`" + `, ` + "`" + `RunPowerShellScript` + "`" + ` and ` + "`" + `RunShellScript` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "working_dir",
					Description: `(Optional, ForceNew) The execution path of the command in the ECS instance. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in terraform of Command. ## Import ECS Command can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_ecs_command.example <id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in terraform of Command. ## Import ECS Command can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_ecs_command.example <id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ecs_dedicated_host",
			Category:         "ECS",
			ShortDescription: `Provides a Alibabacloudstack ecs dedicated host resource.`,
			Description:      ``,
			Keywords: []string{
				"ecs",
				"dedicated",
				"host",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "action_on_maintenance",
					Description: `(Optional) The policy used to migrate the instances from the dedicated host when the dedicated host fails or needs to be repaired online. Valid values: ` + "`" + `Migrate` + "`" + `, ` + "`" + `Stop` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "auto_placement",
					Description: `(Optional, Computed) Specifies whether to add the dedicated host to the resource pool for automatic deployment. If you do not specify the DedicatedHostId parameter when you create an instance on a dedicated host, Alibaba Cloud automatically selects a dedicated host from the resource pool to host the instance. Valid values: ` + "`" + `on` + "`" + `, ` + "`" + `off` + "`" + `. Default: ` + "`" + `on` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "auto_release_time",
					Description: `(Optional, Computed) The automatic release time of the dedicated host. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC+0.`,
				},
				resource.Attribute{
					Name:        "auto_renew",
					Description: `(Optional) Specifies whether to automatically renew the subscription dedicated host.`,
				},
				resource.Attribute{
					Name:        "auto_renew_period",
					Description: `(Optional) The auto-renewal period of the dedicated host. Unit: months. Valid values: ` + "`" + `1` + "`" + `, ` + "`" + `2` + "`" + `, ` + "`" + `3` + "`" + `, ` + "`" + `6` + "`" + `, and ` + "`" + `12` + "`" + `. takes effect and is required only when the AutoRenew parameter is set to true.`,
				},
				resource.Attribute{
					Name:        "dedicated_host_name",
					Description: `(Optional, Computed) The name of the dedicated host. The name must be 2 to 128 characters in length. It must start with a letter but cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).`,
				},
				resource.Attribute{
					Name:        "dedicated_host_type",
					Description: `(Required, ForceNew, Computed) The type of the dedicated host. You can call the [DescribeDedicatedHostTypes](https://www.alibabacloud.com/help/doc-detail/134240.htm) operation to obtain the most recent list of dedicated host types.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, Computed) The description of the dedicated host. The description must be 2 to 256 characters in length and cannot start with http:// or https://.`,
				},
				resource.Attribute{
					Name:        "detail_fee",
					Description: `(Optional) Specifies whether to return the billing details of the order when the billing method is changed from subscription to pay-as-you-go. Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dry_run",
					Description: `(Optional) Specifies whether to only validate the request. Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "expired_time",
					Description: `(Optional, Computed) The subscription period of the dedicated host. The Period parameter takes effect and is required only when the ChargeType parameter is set to PrePaid.`,
				},
				resource.Attribute{
					Name:        "network_attributes",
					Description: `(Optional) dedicated host network parameters. contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "slb_udp_timeout",
					Description: `The timeout period for a UDP session between Server Load Balancer (SLB) and the dedicated host. Unit: seconds. Valid values: 15 to 310.`,
				},
				resource.Attribute{
					Name:        "udp_timeout",
					Description: `The timeout period for a UDP session between a user and an Alibaba Cloud service on the dedicated host. Unit: seconds. Valid values: 15 to 310.`,
				},
				resource.Attribute{
					Name:        "payment_type",
					Description: `(Optional, Computed) The billing method of the dedicated host. Valid values: ` + "`" + `PrePaid` + "`" + `, ` + "`" + `PostPaid` + "`" + `. Default: ` + "`" + `PostPaid` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional, Computed) The ID of the resource group to which the dedicated host belongs.`,
				},
				resource.Attribute{
					Name:        "sale_cycle",
					Description: `(Optional, Computed) The unit of the subscription period of the dedicated host.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Optional, ForceNew, Computed) The zone ID of the dedicated host. This parameter is empty by default. If you do not specify this parameter, the system automatically selects a zone.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "cpu_over_commit_ratio",
					Description: `(Optional, Available in 1.123.1+) CPU oversold ratio. Only custom specifications g6s, c6s, r6s support setting the CPU oversold ratio.`,
				},
				resource.Attribute{
					Name:        "dedicated_host_cluster_id",
					Description: `(Optional, Available in 1.123.1+) The dedicated host cluster ID to which the dedicated host belongs.`,
				},
				resource.Attribute{
					Name:        "min_quantity",
					Description: `(Optional, Available in 1.123.1+) Specify the minimum purchase quantity of a dedicated host. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the dedicated host.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the dedicated host. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 11 mins) Used when create the dedicated host.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 1 mins) Used when delete the dedicated host.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 11 mins) Used when update the dedicated host. ## Import Ecs dedicated host can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_ecs_dedicated_host.default dh-2zedmxxxx ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the dedicated host.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the dedicated host. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 11 mins) Used when create the dedicated host.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 1 mins) Used when delete the dedicated host.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 11 mins) Used when update the dedicated host. ## Import Ecs dedicated host can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_ecs_dedicated_host.default dh-2zedmxxxx ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ecs_deployment_set",
			Category:         "ECS",
			ShortDescription: `Provides a Alibabacloudstack ECS Deployment Set resource.`,
			Description:      ``,
			Keywords: []string{
				"ecs",
				"deployment",
				"set",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "deployment_set_name",
					Description: `(Optional) The name of the deployment set. The name must be 2 to 128 characters in length and can contain letters, digits, colons (:), underscores (_), and hyphens (-). It must start with a letter and cannot start with ` + "`" + `http://` + "`" + ` or ` + "`" + `https://` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the deployment set. The description must be 2 to 256 characters in length and cannot start with ` + "`" + `http://` + "`" + ` or ` + "`" + `https://` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional, ForceNew) The deployment domain. Valid values: ` + "`" + `Default` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "granularity",
					Description: `(Optional, ForceNew) The deployment granularity. Valid values: ` + "`" + `Host` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "on_unable_to_redeploy_failed_instance",
					Description: `(Optional) The on unable to redeploy failed instance. Valid values: ` + "`" + `CancelMembershipAndStart` + "`" + `, ` + "`" + `KeepStopped` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "CancelMembershipAndStart",
					Description: `Removes the instances from the deployment set and restarts the instances immediately after the failover is complete.`,
				},
				resource.Attribute{
					Name:        "strategy",
					Description: `(Optional, ForceNew) The deployment strategy. Valid values: ` + "`" + `Availability` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in terraform of Deployment Set. ## Import ECS Deployment Set can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_ecs_deployment_set.example <id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in terraform of Deployment Set. ## Import ECS Deployment Set can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_ecs_deployment_set.example <id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ecs_hpc_cluster",
			Category:         "ECS",
			ShortDescription: `Provides a Alibabacloudstack ECS Hpc Cluster resource.`,
			Description:      ``,
			Keywords: []string{
				"ecs",
				"hpc",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of ECS Hpc Cluster.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of ECS Hpc Cluster. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in terraform of Hpc Cluster. ## Import ECS Hpc Cluster can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_ecs_hpc_cluster.example <id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in terraform of Hpc Cluster. ## Import ECS Hpc Cluster can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_ecs_hpc_cluster.example <id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ecs_key_pair",
			Category:         "ECS",
			ShortDescription: `Provides a Alibabacloudstack ECS Key Pair resource.`,
			Description:      ``,
			Keywords: []string{
				"ecs",
				"key",
				"pair",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key_file",
					Description: `(Optional, ForceNew) The key file.`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `(Optional, ForceNew) The key pair's name. It is the only in one Alicloud account.`,
				},
				resource.Attribute{
					Name:        "key_name_prefix",
					Description: `(Optional, ForceNew) The key pair name's prefix. It is conflict with ` + "`" + `key_pair_name` + "`" + `. If it is specified, terraform will using it to build the only key name.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(Optional, ForceNew) You can import an existing public key and using Alicloud key pair to manage it. If this parameter is specified, ` + "`" + `resource_group_id` + "`" + ` is the key pair belongs.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional) The Id of resource group which the key pair belongs.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in terraform of Key Pair. Value as ` + "`" + `key_pair_name` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in terraform of Key Pair. Value as ` + "`" + `key_pair_name` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ecs_key_pair_attachment",
			Category:         "ECS",
			ShortDescription: `Provides a Alibabacloudstack ECS Key Pair Attachment resource.`,
			Description:      ``,
			Keywords: []string{
				"ecs",
				"key",
				"pair",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key_pair_name",
					Description: `(Required, ForceNew) The name of key pair used to bind.`,
				},
				resource.Attribute{
					Name:        "force",
					Description: `(Optional, ForceNew) Set it to true and it will reboot instances which attached with the key pair to make key pair affect immediately.`,
				},
				resource.Attribute{
					Name:        "instance_ids",
					Description: `(Required, ForceNew) The list of ECS instance's IDs. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID of Key Pair Attachment. The value is formatted ` + "`" + `<key_pair_name>:<instance_ids>` + "`" + `. ## Import ECS Key Pair Attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ecs_key_pair_attachment.example <key_pair_name>:<instance_ids> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID of Key Pair Attachment. The value is formatted ` + "`" + `<key_pair_name>:<instance_ids>` + "`" + `. ## Import ECS Key Pair Attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ecs_key_pair_attachment.example <key_pair_name>:<instance_ids> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ecs_launch_template",
			Category:         "ECS",
			ShortDescription: `Provides a Alibabacloudstack ECS Launch Template resource.`,
			Description:      ``,
			Keywords: []string{
				"ecs",
				"launch",
				"template",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "auto_release_time",
					Description: `(Optional) Instance auto release time. The time is presented using the ISO8601 standard and in UTC time. The format is YYYY-MM-DDTHH:MM:SSZ.`,
				},
				resource.Attribute{
					Name:        "data_disks",
					Description: `(Optional) The list of data disks created with instance.`,
				},
				resource.Attribute{
					Name:        "deployment_set_id",
					Description: `(Optional) The Deployment Set Id.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of instance launch template version 1. It can be [2, 256] characters in length. It cannot start with "http://" or "https://". The default value is null.`,
				},
				resource.Attribute{
					Name:        "enable_vm_os_config",
					Description: `(Optional) Whether to enable the instance operating system configuration.`,
				},
				resource.Attribute{
					Name:        "host_name",
					Description: `(Optional) Instance host name.It cannot start or end with a period (.) or a hyphen (-) and it cannot have two or more consecutive periods (.) or hyphens (-).For Windows: The host name can be [2, 15] characters in length. It can contain A-Z, a-z, numbers, periods (.), and hyphens (-). It cannot only contain numbers. For other operating systems: The host name can be [2, 64] characters in length. It can be segments separated by periods (.). It can contain A-Z, a-z, numbers, and hyphens (-).`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Optional) The Image ID.`,
				},
				resource.Attribute{
					Name:        "image_owner_alias",
					Description: `(Optional) Mirror source. Valid values: ` + "`" + `system` + "`" + `, ` + "`" + `self` + "`" + `, ` + "`" + `others` + "`" + `, ` + "`" + `marketplace` + "`" + `, ` + "`" + `""` + "`" + `. Default to: ` + "`" + `""` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `(Optional) Billing methods. Valid values: ` + "`" + `PostPaid` + "`" + `, ` + "`" + `PrePaid` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Optional) Instance type. For more information, call resource_alicloud_instances to obtain the latest instance type list.`,
				},
				resource.Attribute{
					Name:        "internet_charge_type",
					Description: `(Optional) Internet bandwidth billing method. Valid values: ` + "`" + `PayByTraffic` + "`" + `, ` + "`" + `PayByBandwidth` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "internet_max_bandwidth_in",
					Description: `(Optional) The maximum inbound bandwidth from the Internet network, measured in Mbit/s. Value range: [1, 200].`,
				},
				resource.Attribute{
					Name:        "internet_max_bandwidth_out",
					Description: `(Optional) Maximum outbound bandwidth from the Internet, its unit of measurement is Mbit/s. Value range: [0, 100].`,
				},
				resource.Attribute{
					Name:        "io_optimized",
					Description: `(Optional) Whether it is an I/O-optimized instance or not. Valid values: ` + "`" + `none` + "`" + `, ` + "`" + `optimized` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "key_pair_name",
					Description: `(Optional) The name of the key pair. - Ignore this parameter for Windows instances. It is null by default. Even if you enter this parameter, only the Password content is used. - The password logon method for Linux instances is set to forbidden upon initialization.`,
				},
				resource.Attribute{
					Name:        "launch_template_name",
					Description: `(Optional, ForceNew) The name of Launch Template.`,
				},
				resource.Attribute{
					Name:        "network_interfaces",
					Description: `(Optional) The list of network interfaces created with instance.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `(Optional) Network type of the instance. Valid values: ` + "`" + `classic` + "`" + `, ` + "`" + `vpc` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "password_inherit",
					Description: `(Optional) Whether to use the password preset by the mirror.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) The subscription period of the instance. Unit: months. This parameter takes effect and is required only when InstanceChargeType is set to PrePaid. If the DedicatedHostId parameter is specified, the value of the Period parameter must be within the subscription period of the dedicated host. - When the PeriodUnit parameter is set to ` + "`" + `Week` + "`" + `, the valid values of the Period parameter are ` + "`" + `1` + "`" + `, ` + "`" + `2` + "`" + `, ` + "`" + `3` + "`" + `, and ` + "`" + `4` + "`" + `. - When the PeriodUnit parameter is set to ` + "`" + `Month` + "`" + `, the valid values of the Period parameter are ` + "`" + `1` + "`" + `, ` + "`" + `2` + "`" + `, ` + "`" + `3` + "`" + `, ` + "`" + `4` + "`" + `, ` + "`" + `5` + "`" + `, ` + "`" + `6` + "`" + `, ` + "`" + `7` + "`" + `, ` + "`" + `8` + "`" + `, ` + "`" + `9` + "`" + `, ` + "`" + `12` + "`" + `, ` + "`" + `24` + "`" + `, ` + "`" + `36` + "`" + `, ` + "`" + `48` + "`" + `, and ` + "`" + `60` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "private_ip_address",
					Description: `(Optional) The private IP address of the instance.`,
				},
				resource.Attribute{
					Name:        "ram_role_name",
					Description: `(Optional) The RAM role name of the instance. You can use the RAM API ListRoles to query instance RAM role names.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional) The ID of the resource group to which to assign the instance, Elastic Block Storage (EBS) device, and ENI.`,
				},
				resource.Attribute{
					Name:        "security_enhancement_strategy",
					Description: `(Optional) Whether or not to activate the security enhancement feature and install network security software free of charge. Valid values: ` + "`" + `Active` + "`" + `, ` + "`" + `Deactive` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Optional) The security group ID.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Optional) The ID of security group N to which to assign the instance.`,
				},
				resource.Attribute{
					Name:        "spot_duration",
					Description: `(Optional, Computed) The protection period of the preemptible instance. Unit: hours. Valid values: ` + "`" + `0` + "`" + `, ` + "`" + `1` + "`" + `, ` + "`" + `2` + "`" + `, ` + "`" + `3` + "`" + `, ` + "`" + `4` + "`" + `, ` + "`" + `5` + "`" + `, and ` + "`" + `6` + "`" + `. Default to: ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "spot_strategy",
					Description: `(Optional) The spot strategy for a Pay-As-You-Go instance. This parameter is valid and required only when InstanceChargeType is set to PostPaid. Valid values: ` + "`" + `NoSpot` + "`" + `, ` + "`" + `SpotAsPriceGo` + "`" + `, ` + "`" + `SpotWithPriceLimit` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "system_disk",
					Description: `(Optional) The System Disk.`,
				},
				resource.Attribute{
					Name:        "template_resource_group_id",
					Description: `(Optional, ForceNew) The template resource group id.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional, Computed) The User Data.`,
				},
				resource.Attribute{
					Name:        "version_description",
					Description: `(Optional) The description of the launch template version. The description must be 2 to 256 characters in length and cannot start with http:// or https://.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Optional) When creating a VPC-Connected instance, you must specify its VSwitch ID.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Optional) The zone ID of the instance.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to instance, block storage, and elastic network. - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string. - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.`,
				},
				resource.Attribute{
					Name:        "template_tags",
					Description: `(Optional) A mapping of tags to assign to the launch template. #### Block system_disk The system_disk supports the following:`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `(Optional, Computed) The category of the system disk. System disk type. Valid values: ` + "`" + `all` + "`" + `, ` + "`" + `cloud` + "`" + `, ` + "`" + `ephemeral_ssd` + "`" + `, ` + "`" + `cloud_essd` + "`" + `, ` + "`" + `cloud_efficiency` + "`" + `, ` + "`" + `cloud_ssd` + "`" + `, ` + "`" + `local_disk` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "delete_with_instance",
					Description: `(Optional) Specifies whether to release the system disk when the instance is released. Default to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, Computed) System disk description. It cannot begin with http:// or https://.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `(Optional) The Iops.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, Computed) System disk name. The name is a string of 2 to 128 characters. It must begin with an English or a Chinese character. It can contain A-Z, a-z, Chinese characters, numbers, periods (.), colons (:), underscores (_), and hyphens (-).`,
				},
				resource.Attribute{
					Name:        "performance_level",
					Description: `(Optional) The performance level of the ESSD used as the system disk. Valid Values: ` + "`" + `PL0` + "`" + `, ` + "`" + `PL1` + "`" + `, ` + "`" + `PL2` + "`" + `, and ` + "`" + `PL3` + "`" + `. Default to: ` + "`" + `PL0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional, Computed) Size of the system disk, measured in GB. Value range: [20, 500]. #### Block network_interfaces The network_interfaces supports the following:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The ENI description.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The ENI name.`,
				},
				resource.Attribute{
					Name:        "primary_ip",
					Description: `(Optional) The primary private IP address of the ENI.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Optional) The security group ID must be one in the same VPC.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Optional) The VSwitch ID for ENI. The instance must be in the same zone of the same VPC network as the ENI, but they may belong to different VSwitches. #### Block data_disks The data_disks supports the following:`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `(Optional) The category of the disk.`,
				},
				resource.Attribute{
					Name:        "delete_with_instance",
					Description: `(Optional) Indicates whether the data disk is released with the instance.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the data disk.`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `(Optional) Encrypted the data in this disk.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the data disk.`,
				},
				resource.Attribute{
					Name:        "performance_level",
					Description: `(Optional) The performance level of the ESSD used as the data disk.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) The size of the data disk.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) The snapshot ID used to initialize the data disk. If the size specified by snapshot is greater that the size of the disk, use the size specified by snapshot as the size of the data disk. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in terraform of Launch Template. ## Import ECS Launch Template can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ecs_launch_template.example <id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in terraform of Launch Template. ## Import ECS Launch Template can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ecs_launch_template.example <id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ecs_network_interface",
			Category:         "ECS",
			ShortDescription: `Provides a Alibabacloudstack ECS Network Interface resource.`,
			Description:      ``,
			Keywords: []string{
				"ecs",
				"network",
				"interface",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the ENI. The description must be 2 to 256 characters in length and cannot start with http:// or https://.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, Computed, Deprecated in v1.123.1+) Field ` + "`" + `name` + "`" + ` has been deprecated from provider version 1.123.1. New field ` + "`" + `network_interface_name` + "`" + ` instead`,
				},
				resource.Attribute{
					Name:        "network_interface_name",
					Description: `(Optional, Computed) The name of the ENI. The name must be 2 to 128 characters in length, and can contain letters, digits, colons (:), underscores (_), and hyphens (-). It must start with a letter and cannot start with http:// or https://.`,
				},
				resource.Attribute{
					Name:        "primary_ip_address",
					Description: `(Optional, Computed, ForceNew) The primary private IP address of the ENI. The specified IP address must be available within the CIDR block of the VSwitch. If this parameter is not specified, an available IP address is assigned from the VSwitch CIDR block at random.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `(Optional, Computed, ForceNew, Deprecated in v1.123.1+) Field ` + "`" + `private_ip` + "`" + ` has been deprecated from provider version 1.123.1. New field ` + "`" + `primary_ip_address` + "`" + ` instead`,
				},
				resource.Attribute{
					Name:        "private_ip_addresses",
					Description: `(Optional, Computed) Specifies secondary private IP address N of the ENI. This IP address must be an available IP address within the CIDR block of the VSwitch to which the ENI belongs.`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `(Optional, Computed, Deprecated in v1.123.1+) Field ` + "`" + `private_ips` + "`" + ` has been deprecated from provider version 1.123.1. New field ` + "`" + `private_ip_addresses` + "`" + ` instead`,
				},
				resource.Attribute{
					Name:        "private_ips_count",
					Description: `(Optional, Computed, Deprecated in v1.123.1+) Field ` + "`" + `private_ips_count` + "`" + ` has been deprecated from provider version 1.123.1. New field ` + "`" + `secondary_private_ip_address_count` + "`" + ` instead`,
				},
				resource.Attribute{
					Name:        "queue_number",
					Description: `(Optional, Computed) The queue number of the ENI.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional, ForceNew) The resource group id.`,
				},
				resource.Attribute{
					Name:        "secondary_private_ip_address_count",
					Description: `(Optional, Computed) The number of private IP addresses that can be automatically created by ECS.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Optional, Computed) The ID of security group N. The security groups and the ENI must belong to the same VPC. The valid values of N are based on the maximum number of security groups to which an ENI can be added.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Optional, Computed, Deprecated in v1.123.1+) Field ` + "`" + `security_groups` + "`" + ` has been deprecated from provider version 1.123.1. New field ` + "`" + `security_group_ids` + "`" + ` instead`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Required, ForceNew) The ID of the VSwitch in the specified VPC. The private IP addresses assigned to the ENI must be available IP addresses within the CIDR block of the VSwitch.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in terraform of Network Interface.`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `The MAC address of the ENI.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the ENI. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 2 mins) Used when create the Network Interface.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 1 mins) Used when delete the Network Interface.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 1 mins) Used when update the Network Interface. ## Import ECS Network Interface can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ecs_network_interface.example eni-abcd12345 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in terraform of Network Interface.`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `The MAC address of the ENI.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the ENI. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 2 mins) Used when create the Network Interface.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 1 mins) Used when delete the Network Interface.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 1 mins) Used when update the Network Interface. ## Import ECS Network Interface can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ecs_network_interface.example eni-abcd12345 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ecs_network_interface_attachment",
			Category:         "ECS",
			ShortDescription: `Provides a Alibabacloudstack ECS Network Interface Attachment resource.`,
			Description:      ``,
			Keywords: []string{
				"ecs",
				"network",
				"interface",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, ForceNew) The instance id.`,
				},
				resource.Attribute{
					Name:        "network_interface_id",
					Description: `(Required, ForceNew) The network interface id.`,
				},
				resource.Attribute{
					Name:        "trunk_network_instance_id",
					Description: `(Optional) The trunk network instance id.`,
				},
				resource.Attribute{
					Name:        "wait_for_network_configuration_ready",
					Description: `(Optional) The wait for network configuration ready. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID of Network Interface Attachment. The value is formatted ` + "`" + `<network_interface_id>:<instance_id>` + "`" + `. ## Import ECS Network Interface Attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ecs_network_interface_attachment.example eni-abcd1234:i-abcd1234 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID of Network Interface Attachment. The value is formatted ` + "`" + `<network_interface_id>:<instance_id>` + "`" + `. ## Import ECS Network Interface Attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ecs_network_interface_attachment.example eni-abcd1234:i-abcd1234 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_edas_application",
			Category:         "EDAS",
			ShortDescription: `Creates an EDAS ecs application on EDAS.`,
			Description:      ``,
			Keywords: []string{
				"edas",
				"application",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "application_name",
					Description: `(Required) Name of your EDAS application. Only letters '-' '_' and numbers are allowed. The length cannot exceed 36 characters.`,
				},
				resource.Attribute{
					Name:        "package_type",
					Description: `(Required) The type of the package for the deployment of the application that you want to create. The valid values are: WAR and JAR. We strongly recommend you to set this parameter when creating the application.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) The ID of the cluster that you want to create the application. The default cluster will be used if you do not specify this parameter.`,
				},
				resource.Attribute{
					Name:        "build_pack_id",
					Description: `(Optional) The package ID of Enterprise Distributed Application Service (EDAS) Container, which can be retrieved by calling container version list interface ListBuildPack or the "Pack ID" column in container version list. When creating High-speed Service Framework (HSF) application, this parameter is required.`,
				},
				resource.Attribute{
					Name:        "descriotion",
					Description: `(Optional) The description of the application that you want to create.`,
				},
				resource.Attribute{
					Name:        "health_check_url",
					Description: `(Optional) The URL for health checking of the application.`,
				},
				resource.Attribute{
					Name:        "logical_region_id",
					Description: `(Optional) The ID of the namespace where you want to create the application. You can call the ListUserDefineRegion operation to query the namespace ID.`,
				},
				resource.Attribute{
					Name:        "component_ids",
					Description: `(Optional) The ID of the component in the container where the application is going to be deployed. If the runtime environment is not specified when the application is created and the application is not deployed, you can set the parameter as fellow: when deploying a native Dubbo or Spring Cloud application using a WAR package for the first time, you must specify the version of the Apache Tomcat component based on the deployed application. You can call the ListClusterOperation interface to query the components. When deploying a non-native Dubbo or Spring Cloud application using a WAR package for the first time, you can leave this parameter empty.`,
				},
				resource.Attribute{
					Name:        "ecu_info",
					Description: `(Optional) The ID of the Elastic Compute Unit (ECU) where you want to deploy the application. Type: List.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Optional) The ID of the instance group where the application is going to be deployed. Set this parameter to all if you want to deploy the application to all groups.`,
				},
				resource.Attribute{
					Name:        "package_version",
					Description: `(Optional) The version of the application that you want to deploy. It must be unique for every application. The length cannot exceed 64 characters. We recommended you to use a timestamp.`,
				},
				resource.Attribute{
					Name:        "war_url",
					Description: `(Optional) The address to store the uploaded web application (WAR) package for application deployment. This parameter is required when the deployType parameter is set as url. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `key` + "`" + ` of the resource supplied above. The value is formulated as ` + "`" + `app_Id` + "`" + `. ## Import EDAS application can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_edas_application.app app_Id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `key` + "`" + ` of the resource supplied above. The value is formulated as ` + "`" + `app_Id` + "`" + `. ## Import EDAS application can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_edas_application.app app_Id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_edas_cluster",
			Category:         "EDAS",
			ShortDescription: `Provides an EDAS cluster resource.`,
			Description:      ``,
			Keywords: []string{
				"edas",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Required, ForceNew) The name of the cluster that you want to create.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `(Required, ForceNew) The type of the cluster that you want to create. Valid values only: 2: ECS cluster.`,
				},
				resource.Attribute{
					Name:        "network_mode",
					Description: `(Required, ForceNew) The network type of the cluster that you want to create. Valid values: 1: classic network. 2: VPC.`,
				},
				resource.Attribute{
					Name:        "logical_region_id",
					Description: `(Optional, ForceNew) The ID of the namespace where you want to create the application. You can call the ListUserDefineRegion operation to query the namespace ID.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional, ForceNew) The ID of the Virtual Private Cloud (VPC) for the cluster. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `key` + "`" + ` of the resource supplied above. The value is formulated as ` + "`" + `<cluster_id>` + "`" + `. ## Import EDAS cluster can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_edas_cluster.cluster cluster_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `key` + "`" + ` of the resource supplied above. The value is formulated as ` + "`" + `<cluster_id>` + "`" + `. ## Import EDAS cluster can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_edas_cluster.cluster cluster_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_edas_deploy_group",
			Category:         "EDAS",
			ShortDescription: `Provides an EDAS deploy group resource.`,
			Description:      ``,
			Keywords: []string{
				"edas",
				"deploy",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "app_id",
					Description: `(Required, ForceNew) The ID of the application that you want to deploy.`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `(Required, ForceNew) The name of the instance group that you want to create. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `key` + "`" + ` of the resource supplied above. The value is formulated as ` + "`" + `<app_id>:<group_name>:<group_id>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "group_type",
					Description: `The type of the instance group that you want to create. Valid values: 0: Default group. 1: Phased release is disabled for traffic management. 2: Phased release is enabled for traffic management. ## Import EDAS deploy group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_edas_deploy_group.group app_id:group_name:group_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `key` + "`" + ` of the resource supplied above. The value is formulated as ` + "`" + `<app_id>:<group_name>:<group_id>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "group_type",
					Description: `The type of the instance group that you want to create. Valid values: 0: Default group. 1: Phased release is disabled for traffic management. 2: Phased release is enabled for traffic management. ## Import EDAS deploy group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_edas_deploy_group.group app_id:group_name:group_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_edas_instance_cluster_attachment",
			Category:         "EDAS",
			ShortDescription: `Provides an EDAS instance cluster attachment resource.`,
			Description:      ``,
			Keywords: []string{
				"edas",
				"instance",
				"cluster",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required, ForceNew) The ID of the cluster that you want to create the application.`,
				},
				resource.Attribute{
					Name:        "instance_ids",
					Description: `(Required, ForceNew) The ID of instance. Type: list. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `key` + "`" + ` of the resource supplied above. The value is formulated as ` + "`" + `<cluster_id>:<instance_id1,instance_id2>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status_map",
					Description: `The status map of the resource supplied above. The key is instance_id and the values are 1(running) 0(converting) -1(failed) and -2(offline).`,
				},
				resource.Attribute{
					Name:        "ecu_map",
					Description: `The ecu map of the resource supplied above. The key is instance_id and the value is ecu_id.`,
				},
				resource.Attribute{
					Name:        "cluster_member_ids",
					Description: `The cluster members map of the resource supplied above. The key is instance_id and the value is cluster_member_id.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `key` + "`" + ` of the resource supplied above. The value is formulated as ` + "`" + `<cluster_id>:<instance_id1,instance_id2>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status_map",
					Description: `The status map of the resource supplied above. The key is instance_id and the values are 1(running) 0(converting) -1(failed) and -2(offline).`,
				},
				resource.Attribute{
					Name:        "ecu_map",
					Description: `The ecu map of the resource supplied above. The key is instance_id and the value is ecu_id.`,
				},
				resource.Attribute{
					Name:        "cluster_member_ids",
					Description: `The cluster members map of the resource supplied above. The key is instance_id and the value is cluster_member_id.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_edas_k8s_application",
			Category:         "EDAS",
			ShortDescription: `Provides an EDAS K8s cluster resource.`,
			Description:      ``,
			Keywords: []string{
				"edas",
				"k8s",
				"application",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "application_name",
					Description: `(Required, ForceNew) The name of the application you want to create. Must start with character,supports numbers, letters and dashes (-), supports up to 36 characters`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required, ForceNew) The ID of the alibabacloudstack container service kubernetes cluster that you want to import to. You can call the ListCluster operation to query.`,
				},
				resource.Attribute{
					Name:        "package_type",
					Description: `(Required, ForceNew) Application package type. Optional parameter values include: FatJar, WAR and Image.`,
				},
				resource.Attribute{
					Name:        "replicas",
					Description: `(Optional) Number of application instances.`,
				},
				resource.Attribute{
					Name:        "image_url",
					Description: `(Optional) Mirror address. When the package_type is set to 'Image', this parameter item is required.`,
				},
				resource.Attribute{
					Name:        "application_descriotion",
					Description: `(Optional) The description of the application`,
				},
				resource.Attribute{
					Name:        "package_url",
					Description: `(Optional) The url of the package to deploy.Applications deployed through FatJar or WAR packages need to configure it.`,
				},
				resource.Attribute{
					Name:        "package_version",
					Description: `(Optional) The version number of the deployment package. WAR and FatJar types are required. Please customize its meaning.`,
				},
				resource.Attribute{
					Name:        "jdk",
					Description: `(Optional, ForceNew) The JDK version that the deployed package depends on. The optional parameter values are Open JDK 7 and Open JDK 8. Image does not support this parameter.`,
				},
				resource.Attribute{
					Name:        "web_container",
					Description: `(Optional, ForceNew) The Tomcat version that the deployment package depends on. Applicable to Spring Cloud and Dubbo applications deployed through WAR packages. Image does not support this parameter.`,
				},
				resource.Attribute{
					Name:        "edas_container_version",
					Description: `(Optional) EDAS-Container version that the deployed package depends on. Image does not support this parameter.`,
				},
				resource.Attribute{
					Name:        "internet_target_port",
					Description: `(Optional, ForceNew) The private SLB back-end port, is also the service port of the application, ranging from 1 to 65535.`,
				},
				resource.Attribute{
					Name:        "internet_slb_port",
					Description: `(Optional, ForceNew) The public network SLB front-end port, range 1~65535.`,
				},
				resource.Attribute{
					Name:        "internet_slb_protocol",
					Description: `(Optional, ForceNew) The public network SLB protocol supports TCP, HTTP and HTTPS protocols.`,
				},
				resource.Attribute{
					Name:        "internet_slb_id",
					Description: `(Optional, ForceNew) Public network SLB ID. If not configured, EDAS will automatically purchase a new SLB for the user.`,
				},
				resource.Attribute{
					Name:        "limit_cpu",
					Description: `(Optional) During application operation, the CPU quota of the application instance, unit: number of cores.`,
				},
				resource.Attribute{
					Name:        "limit_mem",
					Description: `(Optional) The memory limit of the application instance during application operation, unit: M.`,
				},
				resource.Attribute{
					Name:        "requests_cpu",
					Description: `(Optional) When the application is created, the CPU quota of the application instance, unit: number of cores. When set to 0, it means unlimited.`,
				},
				resource.Attribute{
					Name:        "requests_mem",
					Description: `(Optional) When the application is created, the memory limit of the application instance, unit: M. When set to 0, it means unlimited.`,
				},
				resource.Attribute{
					Name:        "requests_m_cpu",
					Description: `(Optional) When the application is created, the CPU quota of the application instance, unit: number of millcores, similar to request_cpu`,
				},
				resource.Attribute{
					Name:        "limit_m_cpu",
					Description: `(Optional) The CPU quota of the application instance during application operation. Unit: Number of millcores, set to 0 means unlimited, similar to request_cpu.`,
				},
				resource.Attribute{
					Name:        "command",
					Description: `(Optional) The set command, if set, will replace the startup command in the mirror when the mirror is started.`,
				},
				resource.Attribute{
					Name:        "command_args",
					Description: `(Optional) Used in combination with the command, the parameter of the command is a JsonArray string in the format: ` + "`" + `[{"argument":"-c"},{"argument":"test"}]` + "`" + `. Among them, -c and test are two parameters that need to be set.`,
				},
				resource.Attribute{
					Name:        "envs",
					Description: `(Optional, ForceNew) Deployment environment variables, the format must conform to the JSON object array, such as: ` + "`" + `{"name":"x","value":"y"},{"name":"x2","value":"y2"}` + "`" + `, If you want to cancel the configuration, you need to set an empty JSON array "" to indicate no configuration.`,
				},
				resource.Attribute{
					Name:        "pre_stop",
					Description: `(Optional) Execute script before stopping`,
				},
				resource.Attribute{
					Name:        "post_start",
					Description: `(Optional) Execute script after startup`,
				},
				resource.Attribute{
					Name:        "liveness",
					Description: `(Optional) Container survival status monitoring, format such as: ` + "`" + `{"failureThreshold": 3,"initialDelaySeconds": 5,"successThreshold": 1,"timeoutSeconds": 1,"tcpSocket":{"host":"", "port":8080} }` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "readiness",
					Description: `(Optional) Container service status check. If the check fails, the traffic passing through K8s Service will not be transferred to the container. The format is: ` + "`" + `{"failureThreshold": 3,"initialDelaySeconds": 5,"successThreshold": 1,"timeoutSeconds": 1, "httpGet": {"path": "/consumer","port": 8080,"scheme": "HTTP","httpHeaders": [{"name": "test","value": "testvalue"} ]}}` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "nas_id",
					Description: `(Optional) The ID of the mounted NAS must be in the same region as the cluster. It must have an available mount point creation quota, or its mount point must be on a switch in the VPC. If it is not filled in and the mountDescs field exists, a NAS will be automatically purchased and mounted on the switch in the VPC by default.`,
				},
				resource.Attribute{
					Name:        "mount_descs",
					Description: `(Optional, ForceNew) Mount configuration description, as a serialized JSON. For example: ` + "`" + `[{"nasPath": "/k8s","mountPath": "/mnt"},{"nasPath": "/files","mountPath": "/app/files"}]` + "`" + `. Among them, nasPath refers to the file storage path; mountPath refers to the path mounted in the container.`,
				},
				resource.Attribute{
					Name:        "local_volume",
					Description: `(Optional, ForceNew) The configuration of the host file mounted to the container. For example: ` + "`" + `[{"type":"","nodePath":"/localfiles","mountPath":"/app/files"},{"type":"Directory","nodePath":"/mnt", "mountPath":"/app/storage"}]` + "`" + `. Among them, nodePath is the host path; mountPath is the path in the container; type is the mount type.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace of the K8s cluster, it will determine which K8s namespace your application is deployed in. The default is 'default'.`,
				},
				resource.Attribute{
					Name:        "logical_region_id",
					Description: `(Optional) The ID corresponding to the EDAS namespace, the non-default namespace must be filled in. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "application_name",
					Description: `The name of the application you want to create. Must start with character,supports numbers, letters and dashes (-), supports up to 36 characters`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `The ID of the alibabacloudstack container service kubernetes cluster that you want to import to. You can call the ListCluster operation to query.`,
				},
				resource.Attribute{
					Name:        "replicas",
					Description: `Number of application instances.`,
				},
				resource.Attribute{
					Name:        "package_type",
					Description: `Application package type. Optional parameter values include: FatJar, WAR and Image.`,
				},
				resource.Attribute{
					Name:        "image_url",
					Description: `Mirror address. When the package_type is set to 'Image', this parameter item is available. ## Import EDAS k8s application can be imported as below, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_edas_k8s_application.new_k8s_application application_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "application_name",
					Description: `The name of the application you want to create. Must start with character,supports numbers, letters and dashes (-), supports up to 36 characters`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `The ID of the alibabacloudstack container service kubernetes cluster that you want to import to. You can call the ListCluster operation to query.`,
				},
				resource.Attribute{
					Name:        "replicas",
					Description: `Number of application instances.`,
				},
				resource.Attribute{
					Name:        "package_type",
					Description: `Application package type. Optional parameter values include: FatJar, WAR and Image.`,
				},
				resource.Attribute{
					Name:        "image_url",
					Description: `Mirror address. When the package_type is set to 'Image', this parameter item is available. ## Import EDAS k8s application can be imported as below, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_edas_k8s_application.new_k8s_application application_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_edas_k8s_cluster",
			Category:         "EDAS",
			ShortDescription: `Provides an EDAS K8s cluster resource.`,
			Description:      ``,
			Keywords: []string{
				"edas",
				"k8s",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cs_cluster_id",
					Description: `(Required, ForceNew) The ID of the alibabacloudstack container service kubernetes cluster that you want to import.`,
				},
				resource.Attribute{
					Name:        "namespace_id",
					Description: `(Optional, ForceNew) The ID of the namespace where you want to import. You can call the [ListUserDefineRegion](https://www.alibabacloud.com/help/en/doc-detail/149377.htm?spm=a2c63.p38356.879954.34.331054faK2yNvC#doc-api-Edas-ListUserDefineRegion) operation to query the namespace ID. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `The name of the cluster that you want to create.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `The type of the cluster that you want to create. Valid values only: 5: K8s cluster.`,
				},
				resource.Attribute{
					Name:        "network_mode",
					Description: `The network type of the cluster that you want to create. Valid values: 1: classic network. 2: VPC.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `The ID of the region.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of the Virtual Private Cloud (VPC) for the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_import_status",
					Description: `The import status of cluster: ` + "`" + `1` + "`" + `: success. ` + "`" + `2` + "`" + `: failed. ` + "`" + `3` + "`" + `: importing. ` + "`" + `4` + "`" + `: deleted. ## Import EDAS cluster can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_edas_k8s_cluster.cluster cluster_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_name",
					Description: `The name of the cluster that you want to create.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `The type of the cluster that you want to create. Valid values only: 5: K8s cluster.`,
				},
				resource.Attribute{
					Name:        "network_mode",
					Description: `The network type of the cluster that you want to create. Valid values: 1: classic network. 2: VPC.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `The ID of the region.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of the Virtual Private Cloud (VPC) for the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_import_status",
					Description: `The import status of cluster: ` + "`" + `1` + "`" + `: success. ` + "`" + `2` + "`" + `: failed. ` + "`" + `3` + "`" + `: importing. ` + "`" + `4` + "`" + `: deleted. ## Import EDAS cluster can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_edas_k8s_cluster.cluster cluster_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_edas_slb_attachment",
			Category:         "EDAS",
			ShortDescription: `Binds SLB to an EDAS application.`,
			Description:      ``,
			Keywords: []string{
				"edas",
				"slb",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "app_id",
					Description: `(Required, ForceNew) The ID of the application to which you want to bind an SLB instance.`,
				},
				resource.Attribute{
					Name:        "slb_id",
					Description: `(Required, ForceNew) The ID of the SLB instance that is going to be bound.`,
				},
				resource.Attribute{
					Name:        "slb_ip",
					Description: `(Required, ForceNew) The IP address that is allocated to the bound SLB instance.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required, ForceNew) The type of the bound SLB instance.`,
				},
				resource.Attribute{
					Name:        "listener_port",
					Description: `(Optional, ForceNew) The listening port for the bound SLB instance.`,
				},
				resource.Attribute{
					Name:        "vserver_group_id",
					Description: `(Optional, ForceNew) The ID of the virtual server (VServer) group associated with the intranet SLB instance. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `key` + "`" + ` of the resource supplied above. The value is formulated as ` + "`" + `<app_id>:<slb_id>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "slb_status",
					Description: `Running Status of SLB instance. Inactive：The instance is stopped, and listener will not monitor and forward traffic. Active：The instance is running. After the instance is created, the default state is active. Locked：The instance is locked, the instance has been owed or locked by Alibaba Cloud. Expired: The instance has expired.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `VPC related vswitch ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `key` + "`" + ` of the resource supplied above. The value is formulated as ` + "`" + `<app_id>:<slb_id>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "slb_status",
					Description: `Running Status of SLB instance. Inactive：The instance is stopped, and listener will not monitor and forward traffic. Active：The instance is running. After the instance is created, the default state is active. Locked：The instance is locked, the instance has been owed or locked by Alibaba Cloud. Expired: The instance has expired.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `VPC related vswitch ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ehpc_job_template",
			Category:         "Elastic High Performance Computing(ehpc)",
			ShortDescription: `Provides a Alibabacloudstack Ehpc Job Template resource.`,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"high",
				"performance",
				"computing",
				"ehpc",
				"job",
				"template",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "array_request",
					Description: `(Optional) Queue Jobs, Is of the Form: 1-10:2.`,
				},
				resource.Attribute{
					Name:        "clock_time",
					Description: `(Optional) Job Maximum Run Time.`,
				},
				resource.Attribute{
					Name:        "command_line",
					Description: `(Required) Job Commands.`,
				},
				resource.Attribute{
					Name:        "gpu",
					Description: `(Optional) A Single Compute Node Using the GPU Number.Possible Values: 1~20000.`,
				},
				resource.Attribute{
					Name:        "job_template_name",
					Description: `(Required) A Job Template Name.`,
				},
				resource.Attribute{
					Name:        "mem",
					Description: `(Optional) A Single Compute Node Maximum Memory.`,
				},
				resource.Attribute{
					Name:        "node",
					Description: `(Optional) Submit a Task Is Required for Computing the Number of Data Nodes to Be. Possible Values: 1~5000 .`,
				},
				resource.Attribute{
					Name:        "package_path",
					Description: `(Optional) Job Commands the Directory.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) The Job Priority.`,
				},
				resource.Attribute{
					Name:        "queue",
					Description: `(Optional) The Job Queue.`,
				},
				resource.Attribute{
					Name:        "re_runable",
					Description: `(Optional) If the Job Is Support for the Re-Run.`,
				},
				resource.Attribute{
					Name:        "runas_user",
					Description: `(Optional) The name of the user who performed the job.`,
				},
				resource.Attribute{
					Name:        "stderr_redirect_path",
					Description: `(Optional) Error Output Path.`,
				},
				resource.Attribute{
					Name:        "stdout_redirect_path",
					Description: `(Optional) Standard Output Path and.`,
				},
				resource.Attribute{
					Name:        "task",
					Description: `(Optional) A Single Compute Node Required Number of Tasks. Possible Values: 1~20000 .`,
				},
				resource.Attribute{
					Name:        "thread",
					Description: `(Optional) A Single Task and the Number of Required Threads.`,
				},
				resource.Attribute{
					Name:        "variables",
					Description: `(Optional) The Job of the Environment Variable. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in terraform of Job Template. ## Import Ehpc Job Template can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_ehpc_job_template.example <id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in terraform of Job Template. ## Import Ehpc Job Template can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_ehpc_job_template.example <id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_eip",
			Category:         "VPC",
			ShortDescription: `Provides a ECS EIP resource.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"eip",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the EIP instance. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the EIP instance, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Optional) Maximum bandwidth to the elastic public network, measured in Mbps (Mega bit per second). If this value is not specified, then automatically sets it to 5 Mbps.`,
				},
				resource.Attribute{
					Name:        "isp",
					Description: `(Optional, ForceNew) The line type of the Elastic IP instance. Default to ` + "`" + `BGP` + "`" + `. Other type of the isp need to open a whitelist. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The EIP ID.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `The elastic public network bandwidth.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The EIP current status.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The elastic ip address`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The EIP ID.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `The elastic public network bandwidth.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The EIP current status.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The elastic ip address`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_eip_association",
			Category:         "VPC",
			ShortDescription: `Provides a ECS EIP Association resource.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"eip",
				"association",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "allocation_id",
					Description: `(Required, ForcesNew) The allocation EIP ID.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, ForcesNew) The ID of the ECS or SLB instance or Nat Gateway.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Optional, ForceNew) The type of cloud product that the eip instance to bind. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "allocation_id",
					Description: `As above.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `As above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "allocation_id",
					Description: `As above.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `As above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_elasticsearch",
			Category:         "Elasticsearch",
			ShortDescription: `Provides a Alibabacloudstack Elasticsearch instance resource.`,
			Description:      ``,
			Keywords: []string{
				"elasticsearch",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, Computed) The description of instance. It a string of 0 to 30 characters.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `(Optional) Valid values are ` + "`" + `PrePaid` + "`" + `, ` + "`" + `PostPaid` + "`" + `. Default to ` + "`" + `PostPaid` + "`" + `. From version 1.69.0, the Elasticsearch cluster allows you to update your instance_charge_ype from ` + "`" + `PostPaid` + "`" + ` to ` + "`" + `PrePaid` + "`" + `, the following attributes are required: ` + "`" + `period` + "`" + `. But, updating from ` + "`" + `PostPaid` + "`" + ` to ` + "`" + `PrePaid` + "`" + ` is not supported.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) The duration that you will buy Elasticsearch instance (in month). It is valid when instance_charge_type is ` + "`" + `PrePaid` + "`" + `. Valid values: [1~9], 12, 24, 36. Default to 1. From version 1.69.2, when to modify this value, the resource can renewal a ` + "`" + `PrePaid` + "`" + ` instance.`,
				},
				resource.Attribute{
					Name:        "data_node_amount",
					Description: `(Required) The Elasticsearch cluster's data node quantity, between 2 and 50.`,
				},
				resource.Attribute{
					Name:        "data_node_spec",
					Description: `(Required) The data node specifications of the Elasticsearch instance.`,
				},
				resource.Attribute{
					Name:        "data_node_disk_size",
					Description: `(Required) The single data node storage space. - ` + "`" + `cloud_ssd` + "`" + `: An SSD disk, supports a maximum of 2048 GiB (2 TB). - ` + "`" + `cloud_efficiency` + "`" + ` An ultra disk, supports a maximum of 5120 GiB (5 TB). If the data to be stored is larger than 2048 GiB, an ultra disk can only support the following data sizes (GiB): [` + "`" + `2560` + "`" + `, ` + "`" + `3072` + "`" + `, ` + "`" + `3584` + "`" + `, ` + "`" + `4096` + "`" + `, ` + "`" + `4608` + "`" + `, ` + "`" + `5120` + "`" + `].`,
				},
				resource.Attribute{
					Name:        "data_node_disk_type",
					Description: `(Required) The data node disk type. Supported values: cloud_ssd, cloud_efficiency.`,
				},
				resource.Attribute{
					Name:        "data_node_disk_encrypted",
					Description: `(Optional, ForceNew, Available in 1.86.0+) If encrypt the data node disk. Valid values are ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + `. Default to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Required, ForceNew) The ID of VSwitch.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional, Sensitive) The password of the instance. The password can be 8 to 30 characters in length and must contain three of the following conditions: uppercase letters, lowercase letters, numbers, and special characters (` + "`" + `!@#$%^&`,
				},
				resource.Attribute{
					Name:        "kms_encrypted_password",
					Description: `(Optional, Available in 1.57.1+) An KMS encrypts password used to a instance. If the ` + "`" + `password` + "`" + ` is filled in, this field will be ignored, but you have to specify one of ` + "`" + `password` + "`" + ` and ` + "`" + `kms_encrypted_password` + "`" + ` fields.`,
				},
				resource.Attribute{
					Name:        "kms_encryption_context",
					Description: `(Optional, MapString, Available in 1.57.1+) An KMS encryption context used to decrypt ` + "`" + `kms_encrypted_password` + "`" + ` before creating or updating instance with ` + "`" + `kms_encrypted_password` + "`" + `. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when ` + "`" + `kms_encrypted_password` + "`" + ` is set.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required, ForceNew) Elasticsearch version. Supported values: ` + "`" + `5.5.3_with_X-Pack` + "`" + `, ` + "`" + `6.3_with_X-Pack` + "`" + `, ` + "`" + `6.7_with_X-Pack` + "`" + `, ` + "`" + `6.8_with_X-Pack` + "`" + `, ` + "`" + `7.4_with_X-Pack` + "`" + ` and ` + "`" + `7.7_with_X-Pack` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "private_whitelist",
					Description: `(Optional) Set the instance's IP whitelist in VPC network.`,
				},
				resource.Attribute{
					Name:        "public_whitelist",
					Description: `(Optional) Set the instance's IP whitelist in internet network.`,
				},
				resource.Attribute{
					Name:        "enable_public",
					Description: `(Optional, Available in v1.87.0+) Bool, default to false. When it set to true, the instance can enable public network access。`,
				},
				resource.Attribute{
					Name:        "kibana_whitelist",
					Description: `(Optional) Set the Kibana's IP whitelist in internet network.`,
				},
				resource.Attribute{
					Name:        "enable_kibana_public_network",
					Description: `(Optional, Available in v1.87.0+) Bool, default to true. When it set to false, the instance can enable kibana public network access。`,
				},
				resource.Attribute{
					Name:        "kibana_private_whitelist",
					Description: `(Optional, Available in v1.87.0+) Set the Kibana's IP whitelist in private network.`,
				},
				resource.Attribute{
					Name:        "enable_kibana_private_network",
					Description: `(Optional, Available in v1.87.0+) Bool, default to false. When it set to true, the instance can close kibana private network access。`,
				},
				resource.Attribute{
					Name:        "master_node_spec",
					Description: `(Optional) The dedicated master node spec. If specified, dedicated master node will be created.`,
				},
				resource.Attribute{
					Name:        "client_node_amount",
					Description: `(Optional, Available in v1.101.0+) The Elasticsearch cluster's client node quantity, between 2 and 25.`,
				},
				resource.Attribute{
					Name:        "client_node_spec",
					Description: `(Optional, Available in v1.101.0+) The client node spec. If specified, client node will be created.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional, Available in v1.101.0+) Elasticsearch protocol. Supported values: ` + "`" + `HTTP` + "`" + `, ` + "`" + `HTTPS` + "`" + `.default is ` + "`" + `HTTP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "zone_count",
					Description: `(Optional, ForceNew, Available in 1.44.0+) The Multi-AZ supported for Elasticsearch, between 1 and 3. The ` + "`" + `data_node_amount` + "`" + ` value must be an integral multiple of the ` + "`" + `zone_count` + "`" + ` value.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Available in v1.73.0+) A mapping of tags to assign to the resource. - key: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:". It cannot contain "http://" and "https://". It cannot be a null string. - value: It can be up to 128 characters in length. It cannot contain "http://" and "https://". It can be a null string.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional, ForceNew, Computed Available in v1.86.0+) The Id of resource group which the Elasticsearch instance belongs.`,
				},
				resource.Attribute{
					Name:        "setting_config",
					Description: `(Optional, Computed Available in v1.125.0+) The YML configuration of the instance.[Detailed introduction](https://www.alibabacloud.com/help/doc-detail/61336.html). ### Timeouts ->`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 120 mins) Used when creating the elasticsearch instance (until it reaches the initial ` + "`" + `active` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 120 mins) Used when activating the elasticsearch instance when necessary during update - e.g. when changing elasticsearch instance description, whitelist, data node settings, master node spec and password.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 120 mins) Used when terminating the elasticsearch instance. ` + "`" + `Note` + "`" + `: There are 5 minutes to sleep to eusure the instance is deleted. It is not in the timeouts configure. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Elasticsearch instance.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Instance connection domain (only VPC network access supported).`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Instance connection port.`,
				},
				resource.Attribute{
					Name:        "kibana_domain",
					Description: `Kibana console domain (Internet access supported).`,
				},
				resource.Attribute{
					Name:        "kibana_port",
					Description: `Kibana console port.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The Elasticsearch instance status. Includes ` + "`" + `active` + "`" + `, ` + "`" + `activating` + "`" + `, ` + "`" + `inactive` + "`" + `. Some operations are denied when status is not ` + "`" + `active` + "`" + `. ## Import Elasticsearch can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_elasticsearch_instance.example es-cn-abcde123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Elasticsearch instance.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Instance connection domain (only VPC network access supported).`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Instance connection port.`,
				},
				resource.Attribute{
					Name:        "kibana_domain",
					Description: `Kibana console domain (Internet access supported).`,
				},
				resource.Attribute{
					Name:        "kibana_port",
					Description: `Kibana console port.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The Elasticsearch instance status. Includes ` + "`" + `active` + "`" + `, ` + "`" + `activating` + "`" + `, ` + "`" + `inactive` + "`" + `. Some operations are denied when status is not ` + "`" + `active` + "`" + `. ## Import Elasticsearch can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_elasticsearch_instance.example es-cn-abcde123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ess_alarm",
			Category:         "Auto Scaling(ESS)",
			ShortDescription: `Provides a ESS alarm task resource.`,
			Description:      ``,
			Keywords: []string{
				"auto",
				"scaling",
				"ess",
				"alarm",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name for ess alarm.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description for the alarm.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional) Whether to enable specific ess alarm. Default to true.`,
				},
				resource.Attribute{
					Name:        "alarm_actions",
					Description: `(Required) The list of actions to execute when this alarm transition into an ALARM state. Each action is specified as ess scaling rule ari.`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `(Required, ForceNew) The scaling group associated with this alarm.`,
				},
				resource.Attribute{
					Name:        "metric_type",
					Description: `(Optional, ForceNew) The type for the alarm's associated metric. Supported value: system, custom. "system" means the metric data is collected by Aliyun Cloud Monitor Service(CMS), "custom" means the metric data is upload to CMS by users. Defaults to system.`,
				},
				resource.Attribute{
					Name:        "metric_name",
					Description: `(Required) The name for the alarm's associated metric. See [Block_metricNames_and_dimensions](#block-metricnames_and_dimensions) below for details.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional, ForceNew) The period in seconds over which the specified statistic is applied. Supported value: 60, 120, 300, 900. Defaults to 300.`,
				},
				resource.Attribute{
					Name:        "statistics",
					Description: `(Optional) The statistic to apply to the alarm's associated metric. Supported value: Average, Minimum, Maximum. Defaults to Average.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `(Required) The value against which the specified statistics is compared.`,
				},
				resource.Attribute{
					Name:        "comparison_operator",
					Description: `(Optional) The arithmetic operation to use when comparing the specified Statistic and Threshold. The specified Statistic value is used as the first operand. Supported value: >=, <=, >, <. Defaults to >=.`,
				},
				resource.Attribute{
					Name:        "evaluation_count",
					Description: `(Optional) The number of times that needs to satisfies comparison condition before transition into ALARM state. Defaults to 3.`,
				},
				resource.Attribute{
					Name:        "cloud_monitor_group_id",
					Description: `(Optional) Defines the application group id defined by CMS which is assigned when you upload custom metric to CMS, only available for custom metirc.`,
				},
				resource.Attribute{
					Name:        "dimensions",
					Description: `(Optional) The dimension map for the alarm's associated metric (documented below). For all metrics, you can not set the dimension key as "scaling_group" or "userId", which is set by default, the second dimension for metric, such as "device" for "PackagesNetIn", need to be set by users. ## Block metricNames_and_dimensions Supported metric names and dimensions : | MetricName | Dimensions | | ------------------ | ---------------------------- | | CpuUtilization | user_id,scaling_group | | ClassicInternetRx | user_id,scaling_group | | ClassicInternetTx | user_id,scaling_group | | VpcInternetRx | user_id,scaling_group | | VpcInternetTx | user_id,scaling_group | | IntranetRx | user_id,scaling_group | | IntranetTx | user_id,scaling_group | | LoadAverage | user_id,scaling_group | | MemoryUtilization | user_id,scaling_group | | SystemDiskReadBps | user_id,scaling_group | | SystemDiskWriteBps | user_id,scaling_group | | SystemDiskReadOps | user_id,scaling_group | | SystemDiskWriteOps | user_id,scaling_group | | PackagesNetIn | user_id,scaling_group,device | | PackagesNetOut | user_id,scaling_group,device | | TcpConnection | user_id,scaling_group,state | ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id for ess alarm.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of specified alarm. ## Import Ess alarm can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_ess_alarm.example asg-2ze500_045efffe-4d05 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id for ess alarm.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of specified alarm. ## Import Ess alarm can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_ess_alarm.example asg-2ze500_045efffe-4d05 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ess_attachment",
			Category:         "Auto Scaling(ESS)",
			ShortDescription: `Provides an ESS Attachment resource to attach or remove ECS instances.`,
			Description:      ``,
			Keywords: []string{
				"auto",
				"scaling",
				"ess",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `(Required) ID of the scaling group of a scaling configuration.`,
				},
				resource.Attribute{
					Name:        "instance_ids",
					Description: `(Required) ID of the ECS instance to be attached to the scaling group. You can input up to 20 IDs.`,
				},
				resource.Attribute{
					Name:        "force",
					Description: `(Optional) Whether to remove forcibly "AutoCreated" ECS instances in order to release scaling group capacity "MaxSize" for attaching ECS instances. Default to false. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required, ForceNew) The ESS attachment resource ID.`,
				},
				resource.Attribute{
					Name:        "instance_ids",
					Description: `(Required)ID of list "Attached" ECS instance.`,
				},
				resource.Attribute{
					Name:        "force",
					Description: `Whether to delete "AutoCreated" ECS instances. ## Import ESS attachment can be imported using the id or scaling group id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_ess_attachment.example asg-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required, ForceNew) The ESS attachment resource ID.`,
				},
				resource.Attribute{
					Name:        "instance_ids",
					Description: `(Required)ID of list "Attached" ECS instance.`,
				},
				resource.Attribute{
					Name:        "force",
					Description: `Whether to delete "AutoCreated" ECS instances. ## Import ESS attachment can be imported using the id or scaling group id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_ess_attachment.example asg-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ess_notification",
			Category:         "Auto Scaling(ESS)",
			ShortDescription: `Provides a ESS notification resource.`,
			Description:      ``,
			Keywords: []string{
				"auto",
				"scaling",
				"ess",
				"notification",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `(Required, ForceNew) The ID of the Auto Scaling group.`,
				},
				resource.Attribute{
					Name:        "notification_arn",
					Description: `(Required, ForceNew) The Alibabacloudstack Cloud Resource Name (ARN) for the notification object. The format of ` + "`" + `notification_arn` + "`" + ` is acs:ess:{region}:{account-id}:{resource-relative-id}. Valid values for ` + "`" + `resource-relative-id` + "`" + `: 'cloudmonitor', 'queue/', 'topic/'.`,
				},
				resource.Attribute{
					Name:        "notification_types",
					Description: `(Required) The notification types of Auto Scaling events and resource changes. Supported notification types: 'AUTOSCALING:SCALE_OUT_SUCCESS', 'AUTOSCALING:SCALE_IN_SUCCESS', 'AUTOSCALING:SCALE_OUT_ERROR', 'AUTOSCALING:SCALE_IN_ERROR', 'AUTOSCALING:SCALE_REJECT', 'AUTOSCALING:SCALE_OUT_START', 'AUTOSCALING:SCALE_IN_START', 'AUTOSCALING:SCHEDULE_TASK_EXPIRING'. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of notification resource, which is composed of 'scaling_group_id' and 'notification_arn' in the format of '<scaling_group_id>:<notification_arn>'.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of notification resource, which is composed of 'scaling_group_id' and 'notification_arn' in the format of '<scaling_group_id>:<notification_arn>'.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ess_scaling_configuration",
			Category:         "Auto Scaling(ESS)",
			ShortDescription: `Provides a ESS scaling configuration resource.`,
			Description:      ``,
			Keywords: []string{
				"auto",
				"scaling",
				"ess",
				"configuration",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `(Required, ForceNew) ID of the scaling group of a scaling configuration.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Optional) ID of an image file, indicating the image resource selected when an instance is enabled.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `(Optional, ) Name of an image file, indicating the image resource selected when an instance is enabled.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Optional) Resource type of an ECS instance.`,
				},
				resource.Attribute{
					Name:        "instance_types",
					Description: `(Optional) Resource types of an ECS instance.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Optional) Name of an ECS instance. Default to "ESS-Instance".`,
				},
				resource.Attribute{
					Name:        "is_outdated",
					Description: `(Optional) Whether to use outdated instance type. Default to false.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Optional) ID of the security group used to create new instance. It is conflict with ` + "`" + `security_group_ids` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Optional) List IDs of the security group used to create new instances. It is conflict with ` + "`" + `security_group_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scaling_configuration_name",
					Description: `(Optional) Name shown for the scheduled task. which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores ` + "`" + `_` + "`" + `, hypens ` + "`" + `-` + "`" + `, and decimal point ` + "`" + `.` + "`" + `. If this parameter value is not specified, the default value is ScalingConfigurationId.`,
				},
				resource.Attribute{
					Name:        "internet_max_bandwidth_in",
					Description: `(Optional) Maximum incoming bandwidth from the public network, measured in Mbps (Mega bit per second). The value range is [1,200].`,
				},
				resource.Attribute{
					Name:        "internet_max_bandwidth_out",
					Description: `(Optional) Maximum outgoing bandwidth from the public network, measured in Mbps (Mega bit per second). The value range for PayByBandwidth is [0,100].`,
				},
				resource.Attribute{
					Name:        "credit_specification",
					Description: `(Optional) Performance mode of the t5 burstable instance. Valid values: 'Standard', 'Unlimited'.`,
				},
				resource.Attribute{
					Name:        "system_disk_category",
					Description: `(Optional) Category of the system disk. The parameter value options are ` + "`" + `ephemeral_ssd` + "`" + `, ` + "`" + `cloud_efficiency` + "`" + `, ` + "`" + `cloud_ssd` + "`" + `, ` + "`" + `cloud_essd` + "`" + ` and ` + "`" + `cloud` + "`" + `. ` + "`" + `cloud` + "`" + ` only is used to some no I/O optimized instance. Default to ` + "`" + `cloud_efficiency` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "system_disk_size",
					Description: `(Optional) Size of system disk, in GiB. Optional values: cloud: 20-500, cloud_efficiency: 20-500, cloud_ssd: 20-500, ephemeral_ssd: 20-500 The default value is max{40, ImageSize}. If this parameter is set, the system disk size must be greater than or equal to max{40, ImageSize}.`,
				},
				resource.Attribute{
					Name:        "system_disk_name",
					Description: `(Optional) The name of the system disk. It must be 2 to 128 characters in length. It must start with a letter and cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-). Default value: null.`,
				},
				resource.Attribute{
					Name:        "system_disk_description",
					Description: `(Optional) The description of the system disk. The description must be 2 to 256 characters in length and cannot start with http:// or https://.`,
				},
				resource.Attribute{
					Name:        "system_disk_auto_snapshot_policy_id",
					Description: `(Optional) The id of auto snapshot policy for system disk.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional) Whether enable the specified scaling group(make it active) to which the current scaling configuration belongs.`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `(Optional) Whether active current scaling configuration in the specified scaling group. Default to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "substitute",
					Description: `(Optional) The another scaling configuration which will be active automatically and replace current configuration when setting ` + "`" + `active` + "`" + ` to 'false'. It is invalid when ` + "`" + `active` + "`" + ` is 'true'.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional) User-defined data to customize the startup behaviors of the ECS instance and to pass data into the ECS instance.`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `(Optional) The name of key pair that can login ECS instance successfully without password. If it is specified, the password would be invalid.`,
				},
				resource.Attribute{
					Name:        "role_name",
					Description: `(Optional) Instance RAM role name. The name is provided and maintained by RAM. You can use ` + "`" + `alibabacloudstack_ram_role` + "`" + ` to create a new one.`,
				},
				resource.Attribute{
					Name:        "force_delete",
					Description: `(Optional) The last scaling configuration will be deleted forcibly with deleting its scaling group. Default to false.`,
				},
				resource.Attribute{
					Name:        "data_disk",
					Description: `(Optional) DataDisk mappings to attach to ecs instance. See [Block datadisk](#block-datadisk) below for details.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. It will be applied for ECS instances finally. - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "http://", or "https://". It cannot be a null string. - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "http://", or "https://" It can be a null string.`,
				},
				resource.Attribute{
					Name:        "override",
					Description: `(Optional) Indicates whether to overwrite the existing data. Default to false.`,
				},
				resource.Attribute{
					Name:        "password_inherit",
					Description: `(Optional) Specifies whether to use the password that is predefined in the image. If the PasswordInherit parameter is set to true, the ` + "`" + `password` + "`" + ` and ` + "`" + `kms_encrypted_password` + "`" + ` will be ignored. You must ensure that the selected image has a password configured.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional, ForceNew) The password of the ECS instance. The password must be 8 to 30 characters in length. It must contains at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include ` + "`" + `() ~!@#$%^&`,
				},
				resource.Attribute{
					Name:        "kms_encrypted_password",
					Description: `(Optional) An KMS encrypts password used to a db account. If the ` + "`" + `password` + "`" + ` is filled in, this field will be ignored.`,
				},
				resource.Attribute{
					Name:        "kms_encryption_context",
					Description: `(Optional, MapString) An KMS encryption context used to decrypt ` + "`" + `kms_encrypted_password` + "`" + ` before creating or updating a db account with ` + "`" + `kms_encrypted_password` + "`" + `. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when ` + "`" + `kms_encrypted_password` + "`" + ` is set. ->`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) Size of data disk, in GB. The value ranges [5,2000] for a cloud disk, [5,1024] for an ephemeral disk, [5,800] for an ephemeral_ssd disk, [20,32768] for cloud_efficiency, cloud_ssd, cloud_essd disk.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `(Optional) The mount point of data disk N. Valid values of N: 1 to 16. If this parameter is not specified, the system automatically allocates a mount point to created ECS instances. The name of the mount point ranges from /dev/xvdb to /dev/xvdz in alphabetical order.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `(Optional) Category of data disk. The parameter value options are ` + "`" + `ephemeral_ssd` + "`" + `, ` + "`" + `cloud_efficiency` + "`" + `, ` + "`" + `cloud_ssd` + "`" + ` and ` + "`" + `cloud` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) Snapshot used for creating the data disk. If this parameter is specified, the size parameter is neglected, and the size of the created disk is the size of the snapshot.`,
				},
				resource.Attribute{
					Name:        "delete_with_instance",
					Description: `(Optional) Whether to delete data disks attached on ecs when release ecs instance. Optional value: ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `, default to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `(Optional) Specifies whether data disk N is to be encrypted. Valid values of N: 1 to 16. Valid values: ` + "`" + `true` + "`" + `: encrypted, ` + "`" + `false` + "`" + `: not encrypted. Default value: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "kms_key_id",
					Description: `(Optional) The CMK ID for data disk N. Valid values of N: 1 to 16.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of data disk N. Valid values of N: 1 to 16. It must be 2 to 128 characters in length. It must start with a letter and cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-). Default value: null.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of data disk N. Valid values of N: 1 to 16. The description must be 2 to 256 characters in length and cannot start with http:// or https://.`,
				},
				resource.Attribute{
					Name:        "auto_snapshot_policy_id",
					Description: `(Optional) The id of auto snapshot policy for data disk. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The scaling configuration ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The scaling configuration ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ess_scaling_group",
			Category:         "Auto Scaling(ESS)",
			ShortDescription: `Provides a ESS scaling group resource.`,
			Description:      ``,
			Keywords: []string{
				"auto",
				"scaling",
				"ess",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "min_size",
					Description: `(Required) Minimum number of ECS instances in the scaling group. Value range: [0, 100].`,
				},
				resource.Attribute{
					Name:        "max_size",
					Description: `(Required) Maximum number of ECS instances in the scaling group. Value range: [0, 100].`,
				},
				resource.Attribute{
					Name:        "scaling_group_name",
					Description: `(Optional) Name shown for the scaling group, which must contain 2-40 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain numbers, underscores ` + "`" + `_` + "`" + `, hyphens ` + "`" + `-` + "`" + `, and decimal points ` + "`" + `.` + "`" + `. If this parameter is not specified, the default value is ScalingGroupId.`,
				},
				resource.Attribute{
					Name:        "default_cooldown",
					Description: `(Optional) Default cool-down time (in seconds) of the scaling group. Value range: [0, 86400]. The default value is 300s.`,
				},
				resource.Attribute{
					Name:        "vswitch_ids",
					Description: `(Optional) List of virtual switch IDs in which the ecs instances to be launched.`,
				},
				resource.Attribute{
					Name:        "removal_policies",
					Description: `(Optional) RemovalPolicy is used to select the ECS instances you want to remove from the scaling group when multiple candidates for removal exist. Optional values: - OldestInstance: removes the first ECS instance attached to the scaling group. - NewestInstance: removes the first ECS instance attached to the scaling group. - OldestScalingConfiguration: removes the ECS instance with the oldest scaling configuration. - Default values: OldestScalingConfiguration and OldestInstance. You can enter up to two removal policies.`,
				},
				resource.Attribute{
					Name:        "db_instance_ids",
					Description: `(Optional) If an RDS instance is specified in the scaling group, the scaling group automatically attaches the Intranet IP addresses of its ECS instances to the RDS access whitelist. - The specified RDS instance must be in running status. - The specified RDS instance’s whitelist must have room for more IP addresses.`,
				},
				resource.Attribute{
					Name:        "loadbalancer_ids",
					Description: `(Optional) If a Server Load Balancer instance is specified in the scaling group, the scaling group automatically attaches its ECS instances to the Server Load Balancer instance. - The Server Load Balancer instance must be enabled. - At least one listener must be configured for each Server Load Balancer and it HealthCheck must be on. Otherwise, creation will fail (it may be useful to add a ` + "`" + `depends_on` + "`" + ` argument targeting your ` + "`" + `alibabacloudstack_slb_listener` + "`" + ` in order to make sure the listener with its HealthCheck configuration is ready before creating your scaling group). - The Server Load Balancer instance attached with VPC-type ECS instances cannot be attached to the scaling group. - The default weight of an ECS instance attached to the Server Load Balancer instance is 50. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The scaling group ID.`,
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
					Name:        "scaling_group_name",
					Description: `The name of the scaling group.`,
				},
				resource.Attribute{
					Name:        "default_cooldown",
					Description: `The default cool-down of the scaling group.`,
				},
				resource.Attribute{
					Name:        "removal_policies",
					Description: `The removal policy used to select the ECS instance to remove from the scaling group.`,
				},
				resource.Attribute{
					Name:        "db_instance_ids",
					Description: `The db instances id which the ECS instance attached to.`,
				},
				resource.Attribute{
					Name:        "loadbalancer_ids",
					Description: `The slb instances id which the ECS instance attached to.`,
				},
				resource.Attribute{
					Name:        "vswitch_ids",
					Description: `The vswitches id in which the ECS instance launched.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The scaling group ID.`,
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
					Name:        "scaling_group_name",
					Description: `The name of the scaling group.`,
				},
				resource.Attribute{
					Name:        "default_cooldown",
					Description: `The default cool-down of the scaling group.`,
				},
				resource.Attribute{
					Name:        "removal_policies",
					Description: `The removal policy used to select the ECS instance to remove from the scaling group.`,
				},
				resource.Attribute{
					Name:        "db_instance_ids",
					Description: `The db instances id which the ECS instance attached to.`,
				},
				resource.Attribute{
					Name:        "loadbalancer_ids",
					Description: `The slb instances id which the ECS instance attached to.`,
				},
				resource.Attribute{
					Name:        "vswitch_ids",
					Description: `The vswitches id in which the ECS instance launched.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ess_scaling_lifecycle_hook",
			Category:         "Auto Scaling(ESS)",
			ShortDescription: `Provides a ESS lifecycle hook resource.`,
			Description:      ``,
			Keywords: []string{
				"auto",
				"scaling",
				"ess",
				"lifecycle",
				"hook",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `(Required, ForceNew) The ID of the Auto Scaling group to which you want to assign the lifecycle hook.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, ForceNew) The name of the lifecycle hook, which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores ` + "`" + `_` + "`" + `, hypens ` + "`" + `-` + "`" + `, and decimal point ` + "`" + `.` + "`" + `. If this parameter value is not specified, the default value is lifecycle hook id.`,
				},
				resource.Attribute{
					Name:        "lifecycle_transition",
					Description: `(Required) Type of Scaling activity attached to lifecycle hook. Supported value: SCALE_OUT, SCALE_IN.`,
				},
				resource.Attribute{
					Name:        "heartbeat_timeout",
					Description: `(Optional) Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. When the lifecycle hook times out, Auto Scaling performs the action defined in the default_result parameter. Default value: 600.`,
				},
				resource.Attribute{
					Name:        "default_result",
					Description: `(Optional) Defines the action the Auto Scaling group should take when the lifecycle hook timeout elapses. Applicable value: CONTINUE, ABANDON, default value: CONTINUE.`,
				},
				resource.Attribute{
					Name:        "notification_arn",
					Description: `(Optional) The Arn of notification target.`,
				},
				resource.Attribute{
					Name:        "notification_metadata",
					Description: `(Optional) Additional information that you want to include when Auto Scaling sends a message to the notification target. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of lifecycle hook.`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `The scalingGroupId to which lifecycle belongs.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of lifecycle hook.`,
				},
				resource.Attribute{
					Name:        "default_result",
					Description: `The action the Auto Scaling group should take when the lifecycle hook timeout elapses.`,
				},
				resource.Attribute{
					Name:        "heartbeat_timeout",
					Description: `The amount of time that can elapse before the lifecycle hook time out.`,
				},
				resource.Attribute{
					Name:        "lifecycle_transition",
					Description: `Type of Scaling activity attached to lifecycle hook.`,
				},
				resource.Attribute{
					Name:        "notification_metadata",
					Description: `Additional information that will be sent to notification target.`,
				},
				resource.Attribute{
					Name:        "notification_arn",
					Description: `The arn of notification target.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of lifecycle hook.`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `The scalingGroupId to which lifecycle belongs.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of lifecycle hook.`,
				},
				resource.Attribute{
					Name:        "default_result",
					Description: `The action the Auto Scaling group should take when the lifecycle hook timeout elapses.`,
				},
				resource.Attribute{
					Name:        "heartbeat_timeout",
					Description: `The amount of time that can elapse before the lifecycle hook time out.`,
				},
				resource.Attribute{
					Name:        "lifecycle_transition",
					Description: `Type of Scaling activity attached to lifecycle hook.`,
				},
				resource.Attribute{
					Name:        "notification_metadata",
					Description: `Additional information that will be sent to notification target.`,
				},
				resource.Attribute{
					Name:        "notification_arn",
					Description: `The arn of notification target.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ess_scaling_rule",
			Category:         "Auto Scaling(ESS)",
			ShortDescription: `Provides a ESS scaling rule resource.`,
			Description:      ``,
			Keywords: []string{
				"auto",
				"scaling",
				"ess",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `(Required) ID of the scaling group of a scaling rule.`,
				},
				resource.Attribute{
					Name:        "adjustment_type",
					Description: `(Optional) Adjustment mode of a scaling rule. Optional values: - QuantityChangeInCapacity: It is used to increase or decrease a specified number of ECS instances. - PercentChangeInCapacity: It is used to increase or decrease a specified proportion of ECS instances. - TotalCapacity: It is used to adjust the quantity of ECS instances in the current scaling group to a specified value.`,
				},
				resource.Attribute{
					Name:        "adjustment_value",
					Description: `(Optional) The number of ECS instances to be adjusted in the scaling rule. This parameter is required and applicable only to simple scaling rules. The number of ECS instances to be adjusted in a single scaling activity cannot exceed 500. Value range: - QuantityChangeInCapacity：(0, 500] U (-500, 0] - PercentChangeInCapacity：[0, 10000] U [-100, 0] - TotalCapacity：[0, 1000]`,
				},
				resource.Attribute{
					Name:        "scaling_rule_name",
					Description: `(Optional) Name shown for the scaling rule, which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores ` + "`" + `_` + "`" + `, hypens ` + "`" + `-` + "`" + `, and decimal point ` + "`" + `.` + "`" + `. If this parameter value is not specified, the default value is scaling rule id.`,
				},
				resource.Attribute{
					Name:        "cooldown",
					Description: `(Optional) The cooldown time of the scaling rule. This parameter is applicable only to simple scaling rules. Value range: [0, 86,400], in seconds. The default value is empty，if not set, the return value will be 0, which is the default value of integer. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The scaling rule ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The scaling rule ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ess_scalinggroup_vserver_groups",
			Category:         "Auto Scaling(ESS)",
			ShortDescription: `Provides a ESS Attachment resource to attach or remove vserver groups.`,
			Description:      ``,
			Keywords: []string{
				"auto",
				"scaling",
				"ess",
				"scalinggroup",
				"vserver",
				"groups",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `(Required) ID of the scaling group.`,
				},
				resource.Attribute{
					Name:        "vserver_groups",
					Description: `(Optional) A list of vserver groups attached on scaling group. See [Block vserver_group](#block-vserver_group) below for details.`,
				},
				resource.Attribute{
					Name:        "force",
					Description: `(Optional) If instances of scaling group are attached/removed from slb backend server when attach/detach vserver group from scaling group. Default to true. ## Block vserver_group the vserver_group supports the following:`,
				},
				resource.Attribute{
					Name:        "loadbalancer_id",
					Description: `(Required) Loadbalancer server ID of VServer Group.`,
				},
				resource.Attribute{
					Name:        "vserver_attributes",
					Description: `(Required) A list of VServer Group attributes. See [Block vserver_attribute](#block-vserver_attribute) below for details. ## Block vserver_attribute`,
				},
				resource.Attribute{
					Name:        "vserver_group_id",
					Description: `(Required) ID of VServer Group.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) - The port will be used for VServer Group backend server.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Required) The weight of an ECS instance attached to the VServer Group. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required, ForceNew) The ESS vserver groups attachment resource ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required, ForceNew) The ESS vserver groups attachment resource ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ess_scheduled_task",
			Category:         "Auto Scaling(ESS)",
			ShortDescription: `Provides a ESS schedule resource.`,
			Description:      ``,
			Keywords: []string{
				"auto",
				"scaling",
				"ess",
				"scheduled",
				"task",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "scheduled_action",
					Description: `(Required) The operation to be performed when a scheduled task is triggered. Enter the unique identifier of a scaling rule.`,
				},
				resource.Attribute{
					Name:        "scheduled_task_name",
					Description: `(Optional) Display name of the scheduled task, which must be 2-40 characters (English or Chinese) long.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the scheduled task, which is 2-200 characters (English or Chinese) long.`,
				},
				resource.Attribute{
					Name:        "launch_time",
					Description: `(Required) The time at which the scheduled task is triggered. Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time must be in UTC. You cannot enter a time point later than 90 days from the date of scheduled task creation. If the ` + "`" + `recurrence_type` + "`" + ` parameter is specified, the task is executed repeatedly at the time specified by LaunchTime. Otherwise, the task is only executed once at the date and time specified by LaunchTime.`,
				},
				resource.Attribute{
					Name:        "launch_expiration_time",
					Description: `(Optional) The time period during which a failed scheduled task is retried. Unit: seconds. Valid values: 0 to 21600. Default value: 600`,
				},
				resource.Attribute{
					Name:        "recurrence_type",
					Description: `(Optional) Specifies the recurrence type of the scheduled task. If set, both ` + "`" + `recurrence_value` + "`" + ` and ` + "`" + `recurrence_end_time` + "`" + ` must be set. Valid values: - Daily: The scheduled task is executed once every specified number of days. - Weekly: The scheduled task is executed on each specified day of a week. - Monthly: The scheduled task is executed on each specified day of a month. - Cron: The scheduled task is executed based on the specified cron expression.`,
				},
				resource.Attribute{
					Name:        "recurrence_value",
					Description: `(Optional) Specifies how often a scheduled task recurs. The valid value depends on ` + "`" + `recurrence_type` + "`" + ` - Daily: You can enter one value. Valid values: 1 to 31. - Weekly: You can enter multiple values and separate them with commas (,). For example, the values 0 to 6 correspond to the days of the week in sequence from Sunday to Saturday. - Monthly: You can enter two values in A-B format. Valid values of A and B: 1 to 31. The value of B must be greater than or equal to the value of A. - Cron: You can enter a cron expression which is written in UTC and consists of five fields: minute, hour, day of month (date), month, and day of week. The expression can contain wildcard characters including commas (,), question marks (?), hyphens (-), asterisks (`,
				},
				resource.Attribute{
					Name:        "recurrence_end_time",
					Description: `(Optional) Specifies the end time after which the scheduled task is no longer repeated. Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time must be in UTC. You cannot enter a time point later than 365 days from the date of scheduled task creation.`,
				},
				resource.Attribute{
					Name:        "task_enabled",
					Description: `(Optional) Specifies whether to start the scheduled task. Default to true. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The schedule task ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The schedule task ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_express_connect_physical_connection",
			Category:         "Express Connect",
			ShortDescription: `Provides a Alibabacloudstack Express Connect Physical Connection resource.`,
			Description:      ``,
			Keywords: []string{
				"express",
				"connect",
				"physical",
				"connection",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "access_point_id",
					Description: `(Required, ForceNew) The Physical Leased Line Access Point ID.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Optional, Computed) On the Bandwidth of the ECC Service and Physical Connection.`,
				},
				resource.Attribute{
					Name:        "circuit_code",
					Description: `(Optional) Operators for Physical Connection Circuit Provided Coding.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The Physical Connection to Which the Description.`,
				},
				resource.Attribute{
					Name:        "line_operator",
					Description: `(Required) Provides Access to the Physical Line Operator. Valid values:`,
				},
				resource.Attribute{
					Name:        "peer_location",
					Description: `(Required) and an on-Premises Data Center Location.`,
				},
				resource.Attribute{
					Name:        "physical_connection_name",
					Description: `(Optional) on Behalf of the Resource Name of the Resources-Attribute Field.`,
				},
				resource.Attribute{
					Name:        "port_type",
					Description: `(Optional) The Physical Leased Line Access Port Type. Valid value:`,
				},
				resource.Attribute{
					Name:        "redundant_physical_connection_id",
					Description: `(Optional) Redundant Physical Connection to Which the ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional, Computed) Resources on Behalf of a State of the Resource Attribute Field. Valid values: ` + "`" + `Canceled` + "`" + `, ` + "`" + `Enabled` + "`" + `, ` + "`" + `Terminated` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional, Computed, ForceNew) Physical Private Line of Type. Default Value: VPC. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in terraform of Physical Connection. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 2 mins) Used when create the Physical Connection. ## Import Express Connect Physical Connection can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_express_connect_physical_connection.example <id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in terraform of Physical Connection. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 2 mins) Used when create the Physical Connection. ## Import Express Connect Physical Connection can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_express_connect_physical_connection.example <id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_express_connect_virtual_border_router",
			Category:         "Express Connect",
			ShortDescription: `Provides a Alibabacloudstack Express Connect Virtual Border Router resource.`,
			Description:      ``,
			Keywords: []string{
				"express",
				"connect",
				"virtual",
				"border",
				"router",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "associated_physical_connections",
					Description: `(Optional) The associated physical connections.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Optional) The bandwidth.`,
				},
				resource.Attribute{
					Name:        "circuit_code",
					Description: `(Optional) Operators for physical connection circuit provided coding.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of VBR. Length is from 2 to 256 characters, must start with a letter or the Chinese at the beginning, but not at the http:// Or https:// at the beginning.`,
				},
				resource.Attribute{
					Name:        "detect_multiplier",
					Description: `(Optional, Computed) Detection time multiplier that recipient allows the sender to send a message of the maximum allowable connections for the number of packets, used to detect whether the link normal. Value: 3~10.`,
				},
				resource.Attribute{
					Name:        "enable_ipv6",
					Description: `(Optional, Computed) Whether to Enable IPv6. Valid values: ` + "`" + `false` + "`" + `, ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "local_gateway_ip",
					Description: `(Required) Alibaba Cloud-Connected IPv4 address.`,
				},
				resource.Attribute{
					Name:        "local_ipv6_gateway_ip",
					Description: `(Optional) Alibaba Cloud-Connected IPv6 Address.`,
				},
				resource.Attribute{
					Name:        "min_rx_interval",
					Description: `(Optional, Computed) Configure BFD packet reception interval of values include: 200~1000, unit: ms.`,
				},
				resource.Attribute{
					Name:        "min_tx_interval",
					Description: `(Optional, Computed) Configure BFD packet transmission interval maximum value: 200~1000, unit: ms.`,
				},
				resource.Attribute{
					Name:        "peer_gateway_ip",
					Description: `(Required) The Client-Side Interconnection IPv4 Address.`,
				},
				resource.Attribute{
					Name:        "peer_ipv6_gateway_ip",
					Description: `(Optional) The Client-Side Interconnection IPv6 Address.`,
				},
				resource.Attribute{
					Name:        "peering_ipv6_subnet_mask",
					Description: `(Optional) Alibaba Cloud-Connected IPv6 with Client-Side Interconnection IPv6 of Subnet Mask.`,
				},
				resource.Attribute{
					Name:        "peering_subnet_mask",
					Description: `(Required) Alibaba Cloud-Connected IPv4 and Client-Side Interconnection IPv4 of Subnet Mask.`,
				},
				resource.Attribute{
					Name:        "physical_connection_id",
					Description: `(Required, ForceNew) The ID of the Physical Connection to Which the ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional, Computed) The instance state. Valid values: ` + "`" + `active` + "`" + `, ` + "`" + `deleting` + "`" + `, ` + "`" + `recovering` + "`" + `, ` + "`" + `terminated` + "`" + `, ` + "`" + `terminating` + "`" + `, ` + "`" + `unconfirmed` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vbr_owner_id",
					Description: `(Optional) The vbr owner id.`,
				},
				resource.Attribute{
					Name:        "virtual_border_router_name",
					Description: `(Optional) The name of VBR. Length is from 2 to 128 characters, must start with a letter or the Chinese at the beginning can contain numbers, the underscore character (_) and dash (-). But do not start with http:// or https:// at the beginning.`,
				},
				resource.Attribute{
					Name:        "vlan_id",
					Description: `(Required) The VLAN ID of the VBR. Value range: 0~2999. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in terraform of Virtual Border Router.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The Route Table ID Of the Virtual Border Router. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 2 mins) Used when update the Virtual Border Router. ## Import Express Connect Virtual Border Router can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_express_connect_virtual_border_router.example <id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in terraform of Virtual Border Router.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The Route Table ID Of the Virtual Border Router. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 2 mins) Used when update the Virtual Border Router. ## Import Express Connect Virtual Border Router can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_express_connect_virtual_border_router.example <id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_forward_entry",
			Category:         "VPC",
			ShortDescription: `Provides a Alibabacloudstack forward resource.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"forward",
				"entry",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "forward_table_id",
					Description: `(Required, ForceNew) The value can get from ` + "`" + `alibabacloudstack_nat_gateway` + "`" + ` Attributes "forward_table_ids".`,
				},
				resource.Attribute{
					Name:        "external_ip",
					Description: `(Required, ForceNew) The external ip address, the ip must along bandwidth package public ip which ` + "`" + `alibabacloudstack_nat_gateway` + "`" + ` argument ` + "`" + `bandwidth_packages` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "external_port",
					Description: `(Required) The external port, valid value is 1~65535|any.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `(Required) The ip protocal, valid value is tcp|udp|any.`,
				},
				resource.Attribute{
					Name:        "internal_ip",
					Description: `(Required) The internal ip, must a private ip.`,
				},
				resource.Attribute{
					Name:        "internal_port",
					Description: `(Required) The internal port, valid value is 1~65535|any. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the forward entry. The value formats as ` + "`" + `<forward_table_id>:<forward_entry_id>` + "`" + ``,
				},
				resource.Attribute{
					Name:        "forward_entry_id",
					Description: `The id of the forward entry on the server.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the forward entry. The value formats as ` + "`" + `<forward_table_id>:<forward_entry_id>` + "`" + ``,
				},
				resource.Attribute{
					Name:        "forward_entry_id",
					Description: `The id of the forward entry on the server.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_gpdb_account",
			Category:         "AnalyticDB for PostgreSQL (GPDB)",
			ShortDescription: `Provides a Alibabacloudstack GPDB Account resource.`,
			Description:      ``,
			Keywords: []string{
				"analyticdb",
				"for",
				"postgresql",
				"gpdb",
				"account",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_description",
					Description: `(Optional, ForceNew) The description of the account.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `(Required, ForceNew) The name of the account. The account name must be unique and meet the following requirements:`,
				},
				resource.Attribute{
					Name:        "account_password",
					Description: `(Required) The password of the account. The password must be 8 to 32 characters in length and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include ` + "`" + `! @ # $ % ^ &`,
				},
				resource.Attribute{
					Name:        "db_instance_id",
					Description: `(Required, ForceNew) The ID of the instance. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID of Account. The value formats as ` + "`" + `<db_instance_id>:<account_name>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the account. Valid values: ` + "`" + `Active` + "`" + `, ` + "`" + `Creating` + "`" + ` and ` + "`" + `Deleting` + "`" + `. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 5 mins) Used when create the Account. ## Import GPDB Account can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_gpdb_account.example <db_instance_id>:<account_name> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID of Account. The value formats as ` + "`" + `<db_instance_id>:<account_name>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the account. Valid values: ` + "`" + `Active` + "`" + `, ` + "`" + `Creating` + "`" + ` and ` + "`" + `Deleting` + "`" + `. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 5 mins) Used when create the Account. ## Import GPDB Account can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_gpdb_account.example <db_instance_id>:<account_name> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_gpdb_connection",
			Category:         "AnalyticDB for PostgreSQL (GPDB)",
			ShortDescription: `Provides an AnalyticDB for PostgreSQL instance connection resource.`,
			Description:      ``,
			Keywords: []string{
				"analyticdb",
				"for",
				"postgresql",
				"gpdb",
				"connection",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, ForceNew) The Id of instance that can run database.`,
				},
				resource.Attribute{
					Name:        "connection_prefix",
					Description: `(ForceNew) Prefix of an Internet connection string. It must be checked for uniqueness. It may consist of lowercase letters, numbers, and underlines, and must start with a letter and have no more than 30 characters. Default to <instance_id> + '-tf'.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Internet connection port. Valid value: [3200-3999]. Default to 3306. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 10 mins) Used when creating the Internet connection (until DB instance reaches the initial ` + "`" + `Running` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 10 mins) Used when activating the DB instance during update.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 10 mins) Used when terminating the DB instance. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The current instance connection resource ID. Composed of instance ID and connection string with format ` + "`" + `<instance_id>:<connection_prefix>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "connection_string",
					Description: `Connection instance string.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The ip address of connection string. ## Import AnalyticDB for PostgreSQL's connection can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_gpdb_connection.example abc12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The current instance connection resource ID. Composed of instance ID and connection string with format ` + "`" + `<instance_id>:<connection_prefix>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "connection_string",
					Description: `Connection instance string.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The ip address of connection string. ## Import AnalyticDB for PostgreSQL's connection can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_gpdb_connection.example abc12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_gpdb_instance",
			Category:         "AnalyticDB for PostgreSQL (GPDB)",
			ShortDescription: `Provides a AnalyticDB for PostgreSQL instance resource.`,
			Description:      ``,
			Keywords: []string{
				"analyticdb",
				"for",
				"postgresql",
				"gpdb",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "engine_version",
					Description: `(Required, ForceNew) Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/doc-detail/86908.htm) ` + "`" + `EngineVersion` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_class",
					Description: `(Required) Instance specification. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/86942.htm).`,
				},
				resource.Attribute{
					Name:        "instance_group_count",
					Description: `(Required) The number of groups. Valid values: [2,4,8,16,32]`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The name of DB instance. It a string of 2 to 256 characters.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `(Optional, ForceNew) Valid values are ` + "`" + `PrePaid` + "`" + `, ` + "`" + `PostPaid` + "`" + `,System default to ` + "`" + `PostPaid` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Optional, ForceNew) The Zone to launch the DB instance. it supports multiple zone. If it is a multi-zone and ` + "`" + `vswitch_id` + "`" + ` is specified, the vswitch must in one of them. The multiple zone ID can be retrieved by setting ` + "`" + `multi` + "`" + ` to "true" in the data source ` + "`" + `alibabacloudstack_zones` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Optional, ForceNew) The virtual switch ID to launch DB instances in one VPC.`,
				},
				resource.Attribute{
					Name:        "security_ip_list",
					Description: `(Optional) List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 30 mins) Used when creating the DB instance (until it reaches the initial ` + "`" + `Running` + "`" + ` status). ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Instance. ## Import AnalyticDB for PostgreSQL can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_gpdb_instance.example gp-bp1291daeda44194 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Instance. ## Import AnalyticDB for PostgreSQL can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_gpdb_instance.example gp-bp1291daeda44194 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_hbase_instance",
			Category:         "HBase",
			ShortDescription: `Provides a HBase instance resource.`,
			Description:      ``,
			Keywords: []string{
				"hbase",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) HBase instance name. Length must be 2-128 characters long. Only Chinese characters, English letters, numbers, period (.), underline (_), or dash (-) are permitted.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Optional, ForceNew) The Zone to launch the HBase instance. If vswitch_id is not empty, this zone_id can be "" or consistent.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `(Optional, ForceNew) Valid values are "hbase/hbaseue/bds". The following types are supported after v1.73.0: ` + "`" + `hbaseue` + "`" + ` and ` + "`" + `bds` + "`" + `. Single hbase instance need to set engine=hbase, core_instance_quantity=1.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `(Required, ForceNew) HBase major version. hbase:1.1/2.0, hbaseue:2.0, bds:1.0, unsupport other engine temporarily. Value options can refer to the latest docs [CreateInstance](https://help.aliyun.com/document_detail/144607.html).`,
				},
				resource.Attribute{
					Name:        "core_disk_size",
					Description: `(Optional) User-defined HBase instance one core node's storage. Valid when engine=hbase/hbaseue. Bds engine no need core_disk_size, space.Unit: GB. Value range: - Custom storage space, value range: [20, 64000]. - Cluster [400, 64000], step:40-GB increments. - Single [20-500GB], step:1-GB increments.`,
				},
				resource.Attribute{
					Name:        "pay_type",
					Description: `(Optional) Valid values are ` + "`" + `PrePaid` + "`" + `, ` + "`" + `PostPaid` + "`" + `, System default to ` + "`" + `PostPaid` + "`" + `. You can also convert PostPaid to PrePaid. And support convert PrePaid to PostPaid from 1.115.0+.`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `(Optional, ForceNew) 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, valid when pay_type = PrePaid, unit: month. 12, 24, 36 mean 1, 2, 3 years.`,
				},
				resource.Attribute{
					Name:        "auto_renew",
					Description: `(Optional, ForceNew) Valid values are ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + `, system default to ` + "`" + `false` + "`" + `, valid when pay_type = PrePaid.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Optional, ForceNew) If vswitch_id is not empty, that mean net_type = vpc and has a same region. If vswitch_id is empty, net_type=classic. Intl site not support classic network.`,
				},
				resource.Attribute{
					Name:        "cold_storage_size",
					Description: `(Optional, ForceNew) 0 or [800, 1000000], step:10-GB increments. 0 means is_cold_storage = false. [800, 1000000] means is_cold_storage = true.`,
				},
				resource.Attribute{
					Name:        "maintain_start_time",
					Description: `(Optional, Available in 1.73.0) The start time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time), for example 02:00Z.`,
				},
				resource.Attribute{
					Name:        "maintain_end_time",
					Description: `(Optional, Available in 1.73.0) The end time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time), for example 04:00Z.`,
				},
				resource.Attribute{
					Name:        "deletion_protection",
					Description: `(Optional, Available in 1.73.0) The switch of delete protection. True: delete protect, False: no delete protect. You must set false when you want to delete cluster.`,
				},
				resource.Attribute{
					Name:        "immediate_delete_flag",
					Description: `(Optional, Available in 1.109.0) The switch of delete immediate. True: delete immediate, False: delete delay. You will not found the cluster no matter set true or false.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Available in 1.73.0) A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "account",
					Description: `(Optional, Available in 1.105.0+) The account of the cluster web ui. Size [0-128].`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional, Available in 1.105.0+) The password of the cluster web ui account. Size [0-128].`,
				},
				resource.Attribute{
					Name:        "ip_white",
					Description: `(Optional, Available in 1.105.0+) The white ip list of the cluster.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Optional, Available in 1.105.0+) The security group resource of the cluster. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the HBase.`,
				},
				resource.Attribute{
					Name:        "master_instance_quantity",
					Description: `Count nodes of the master node.`,
				},
				resource.Attribute{
					Name:        "ui_proxy_conn_addrs",
					Description: `(Available in 1.105.0+) The Web UI proxy addresses of the cluster.`,
				},
				resource.Attribute{
					Name:        "zk_conn_addrs",
					Description: `(Available in 1.105.0+) The zookeeper addresses of the cluster.`,
				},
				resource.Attribute{
					Name:        "slb_conn_addrs",
					Description: `(Available in 1.105.0+) The slb service addresses of the cluster. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 30 mins) Used when creating the HBase instance (until it reaches the initial ` + "`" + `ACTIVATION` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 30 mins) Used when updating the HBase instance (until it reaches the initial ` + "`" + `ACTIVATION` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 30 mins) Used when terminating the HBase instance. ## Import HBase can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_hbase_instance.example hb-wz96815u13k659fvd ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the HBase.`,
				},
				resource.Attribute{
					Name:        "master_instance_quantity",
					Description: `Count nodes of the master node.`,
				},
				resource.Attribute{
					Name:        "ui_proxy_conn_addrs",
					Description: `(Available in 1.105.0+) The Web UI proxy addresses of the cluster.`,
				},
				resource.Attribute{
					Name:        "zk_conn_addrs",
					Description: `(Available in 1.105.0+) The zookeeper addresses of the cluster.`,
				},
				resource.Attribute{
					Name:        "slb_conn_addrs",
					Description: `(Available in 1.105.0+) The slb service addresses of the cluster. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 30 mins) Used when creating the HBase instance (until it reaches the initial ` + "`" + `ACTIVATION` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 30 mins) Used when updating the HBase instance (until it reaches the initial ` + "`" + `ACTIVATION` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 30 mins) Used when terminating the HBase instance. ## Import HBase can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_hbase_instance.example hb-wz96815u13k659fvd ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_image",
			Category:         "ECS",
			ShortDescription: `Provides an ECS image resource.`,
			Description:      ``,
			Keywords: []string{
				"ecs",
				"image",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Optional, ForceNew, Conflict with ` + "`" + `snapshot_id ` + "`" + ` and ` + "`" + `disk_device_mapping ` + "`" + `) The instance ID.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `(Optional) The image name. It must be 2 to 128 characters in length, and must begin with a letter or Chinese character (beginning with http:// or https:// is not allowed). It can contain digits, colons (:), underscores (_), or hyphens (-). Default value: null.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the image. It must be 2 to 256 characters in length and must not start with http:// or https://. Default value: null.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional, ForceNew, Conflict with ` + "`" + `instance_id ` + "`" + ` and ` + "`" + `disk_device_mapping ` + "`" + `) Specifies a snapshot that is used to create a custom image.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The tag value of an image. The value of N ranges from 1 to 20.`,
				},
				resource.Attribute{
					Name:        "disk_device_mapping",
					Description: `(Optional, ForceNew, Conflict with ` + "`" + `snapshot_id ` + "`" + ` and ` + "`" + `instance_id ` + "`" + `) Description of the system with disks and snapshots under the image.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional, ForceNew) Specifies the size of a disk in the combined custom image, in GiB. Value range: 5 to 2000.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional, ForceNew) Specifies a snapshot that is used to create a combined custom image.`,
				},
				resource.Attribute{
					Name:        "force",
					Description: `(Optional) Indicates whether to force delete the custom image, Default is ` + "`" + `false` + "`" + `. - true：Force deletes the custom image, regardless of whether the image is currently being used by other instances. - false：Verifies that the image is not currently in use by any other instances before deleting the image. ### Timeouts`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 10 mins) Used when creating the image (until it reaches the initial ` + "`" + `Available` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 10 mins) Used when terminating the image. ### Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the image.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the image.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_image_copy",
			Category:         "ECS",
			ShortDescription: `Provides an ECS image copy resource.`,
			Description:      ``,
			Keywords: []string{
				"ecs",
				"image",
				"copy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "source_image_id",
					Description: `(Required, ForceNew) The source image ID.`,
				},
				resource.Attribute{
					Name:        "source_region_id",
					Description: `(Required, ForceNew) The ID of the region to which the source custom image belongs. You can call [DescribeRegions](https://www.Alibabacloudstackcloud.com/help/doc-detail/25609.htm) to view the latest regions of Alibabacloudstack Cloud.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `(Optional) The image name. It must be 2 to 128 characters in length, and must begin with a letter or Chinese character (beginning with http:// or https:// is not allowed). It can contain digits, colons (:), underscores (_), or hyphens (-). Default value: null.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the image. It must be 2 to 256 characters in length and must not start with http:// or https://. Default value: null.`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `(Optional, ForceNew) Indicates whether to encrypt the image.`,
				},
				resource.Attribute{
					Name:        "kms_key_id",
					Description: `(Optional, ForceNew) Key ID used to encrypt the image.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The tag value of an image. The value of N ranges from 1 to 20.`,
				},
				resource.Attribute{
					Name:        "force",
					Description: `(Optional) Indicates whether to force delete the custom image, Default is ` + "`" + `false` + "`" + `. - true：Force deletes the custom image, regardless of whether the image is currently being used by other instances. - false：Verifies that the image is not currently in use by any other instances before deleting the image. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 10 mins) Used when copying the image (until it reaches the initial ` + "`" + `Available` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 10 mins) Used when terminating the image. ## Attributes Reference0 The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the image.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the image.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_image_export",
			Category:         "ECS",
			ShortDescription: `Provides an ECS image export resource.`,
			Description:      ``,
			Keywords: []string{
				"ecs",
				"image",
				"export",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "image_id",
					Description: `(Required, ForceNew) The source image ID.`,
				},
				resource.Attribute{
					Name:        "oss_bucket",
					Description: `(Required, ForceNew) Save the exported OSS bucket.`,
				},
				resource.Attribute{
					Name:        "oss_prefix",
					Description: `(Optional, ForceNew) The prefix of your OSS Object. It can be composed of numbers or letters, and the character length is 1 ~ 30. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 20 mins) Used when exporting the image (until it reaches the initial ` + "`" + `Available` + "`" + ` status). ## Attributes Reference0 The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the image.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the image.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_image_import",
			Category:         "ECS",
			ShortDescription: `Provides an ECS image import resource.`,
			Description:      ``,
			Keywords: []string{
				"ecs",
				"image",
				"import",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "architecture",
					Description: `(Optional, ForceNew) Specifies the architecture of the system disk after you specify a data disk snapshot as the data source of the system disk for creating an image. Valid values: ` + "`" + `i386` + "`" + ` , Default is ` + "`" + `x86_64` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the image. The length is 2 to 256 English or Chinese characters, and cannot begin with http: // and https: //.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `(Optional) The image name. The length is 2 ~ 128 English or Chinese characters. Must start with a english letter or Chinese, and cannot start with http: // and https: //. Can contain numbers, colons (:), underscores (_), or hyphens (-).`,
				},
				resource.Attribute{
					Name:        "license_type",
					Description: `(Optional, ForceNew) The type of the license used to activate the operating system after the image is imported. Default value: ` + "`" + `Auto` + "`" + `. Valid values: ` + "`" + `Auto` + "`" + `,` + "`" + `Aliyun` + "`" + `,` + "`" + `BYOL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "platform",
					Description: `(Optional, ForceNew) Specifies the operating system platform of the system disk after you specify a data disk snapshot as the data source of the system disk for creating an image. Valid values: ` + "`" + `CentOS` + "`" + `, ` + "`" + `Ubuntu` + "`" + `, ` + "`" + `SUSE` + "`" + `, ` + "`" + `OpenSUSE` + "`" + `, ` + "`" + `Debian` + "`" + `, ` + "`" + `CoreOS` + "`" + `, ` + "`" + `Windows Server 2003` + "`" + `, ` + "`" + `Windows Server 2008` + "`" + `, ` + "`" + `Windows Server 2012` + "`" + `, ` + "`" + `Windows 7` + "`" + `, Default is ` + "`" + `Others Linux` + "`" + `, ` + "`" + `Customized Linux` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `(Optional, ForceNew) Operating system platform type. Valid values: ` + "`" + `windows` + "`" + `, Default is ` + "`" + `linux` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disk_device_mapping",
					Description: `(Optional, ForceNew) Description of the system with disks and snapshots under the image.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `(Optional, ForceNew) The name of disk N in the custom image.`,
				},
				resource.Attribute{
					Name:        "disk_image_size",
					Description: `(Optional, ForceNew) Resolution size. You must ensure that the system disk space ≥ file system space. Ranges: When n = 1, the system disk: 5 ~ 500GiB, When n = 2 ~ 17, that is, data disk: 5 ~ 1000GiB, When temporary is introduced, the system automatically detects the size, which is subject to the detection result.`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Optional, ForceNew) Image format. Value range: When the ` + "`" + `RAW` + "`" + `, ` + "`" + `VHD` + "`" + `, ` + "`" + `qcow2` + "`" + ` is imported into the image, the system automatically detects the image format, whichever comes first.`,
				},
				resource.Attribute{
					Name:        "oss_bucket",
					Description: `(Optional) Save the exported OSS bucket.`,
				},
				resource.Attribute{
					Name:        "oss_object",
					Description: `(Optional, ForceNew) The file name of your OSS Object. ->`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 20 mins) Used when copying the image (until it reaches the initial ` + "`" + `Available` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 20 mins) Used when terminating the image. ## Attributes Reference0 The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the image.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the image.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_image_share_permission",
			Category:         "ECS",
			ShortDescription: `Provides an ECS image share permission resource.`,
			Description:      ``,
			Keywords: []string{
				"ecs",
				"image",
				"share",
				"permission",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "image_id",
					Description: `(Required, ForceNew) The source image ID.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `(Required, ForceNew) Alibabacloudstack Account ID. It is used to share images. ### Attributes Reference0 The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the image. It formats as ` + "`" + `<image_id>:<account_id>` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the image. It formats as ` + "`" + `<image_id>:<account_id>` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_instance",
			Category:         "ECS",
			ShortDescription: `Provides an ECS instance resource.`,
			Description:      ``,
			Keywords: []string{
				"ecs",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "image_id",
					Description: `(Required) The Image to use for the instance. ECS instance's image can be replaced via changing 'image_id'. When it is changed, the instance will reboot to make the change take effect.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Required) The type of instance to start. When it is changed, the instance will reboot to make the change take effect.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Required) A list of security group ids to associate with.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) The Zone to start the instance in. It is ignored and will be computed when set ` + "`" + `vswitch_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Optional) The name of the ECS. This instance_name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://. If not specified, Terraform will autogenerate a default name is ` + "`" + `ECS-Instance` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "system_disk_category",
					Description: `(Optional) Valid values are ` + "`" + `ephemeral_ssd` + "`" + `, ` + "`" + `cloud_efficiency` + "`" + `, ` + "`" + `cloud_ssd` + "`" + `, ` + "`" + `cloud_essd` + "`" + `, ` + "`" + `cloud` + "`" + `. ` + "`" + `cloud` + "`" + ` only is used to some none I/O optimized instance. Default to ` + "`" + `cloud_efficiency` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "system_disk_size",
					Description: `(Optional) Size of the system disk, measured in GiB. Value range: [20, 500]. The specified value must be equal to or greater than max{20, Imagesize}. Default value: max{40, ImageSize}. ECS instance's system disk can be reset when replacing system disk. When it is changed, the instance will reboot to make the change take effect.`,
				},
				resource.Attribute{
					Name:        "system_disk_name",
					Description: `(Optional) Name of the system disk. The name must be 2 to 128 characters in length. It must start with a letter and can contain letters, digits, colons (:), underscores (_), and hyphens (-). It cannot start with http:// or https://. If not specified, this parameter is null. Default value: null`,
				},
				resource.Attribute{
					Name:        "system_disk_description",
					Description: `(Optional) Description of the system disk. The description must be 2 to 256 characters in length. It cannot start with http:// or https://. If not specified, this parameter is null. Default value: null`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the instance, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.`,
				},
				resource.Attribute{
					Name:        "internet_max_bandwidth_in",
					Description: `(Optional) Maximum incoming bandwidth from the public network, measured in Mbps (Mega bit per second). Value range: [1, 200]. If this value is not specified, then automatically sets it to 200 Mbps.`,
				},
				resource.Attribute{
					Name:        "internet_max_bandwidth_out",
					Description: `(Optional) Maximum outgoing bandwidth to the public network, measured in Mbps (Mega bit per second). Value range: [0, 100]. Default to 0 Mbps.`,
				},
				resource.Attribute{
					Name:        "host_name",
					Description: `(Optional) Host name of the ECS, which is a string of at least two characters. “hostname” cannot start or end with “.” or “-“. In addition, two or more consecutive “.” or “-“ symbols are not allowed. On Windows, the host name can contain a maximum of 15 characters, which can be a combination of uppercase/lowercase letters, numerals, and “-“. The host name cannot contain dots (“.”) or contain only numeric characters. When it is changed, the instance will reboot to make the change take effect. On other OSs such as Linux, the host name can contain a maximum of 30 characters, which can be segments separated by dots (“.”), where each segment can contain uppercase/lowercase letters, numerals, or “_“. When it is changed, the instance will reboot to make the change take effect.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional, Sensitive) Password to an instance is a string of 8 to 30 characters. It must contain uppercase/lowercase letters and numerals, but cannot contain special symbols. When it is changed, the instance will reboot to make the change take effect.`,
				},
				resource.Attribute{
					Name:        "kms_encrypted_password",
					Description: `(Optional) An KMS encrypts password used to an instance. If the ` + "`" + `password` + "`" + ` is filled in, this field will be ignored. When it is changed, the instance will reboot to make the change take effect.`,
				},
				resource.Attribute{
					Name:        "kms_encryption_context",
					Description: `(Optional) An KMS encryption context used to decrypt ` + "`" + `kms_encrypted_password` + "`" + ` before creating or updating an instance with ` + "`" + `kms_encrypted_password` + "`" + `. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when ` + "`" + `kms_encrypted_password` + "`" + ` is set. When it is changed, the instance will reboot to make the change take effect.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Optional) The virtual switch ID to launch in VPC. This parameter must be set unless you can create classic network instances. When it is changed, the instance will reboot to make the change take effect.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string. - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional) User-defined data to customize the startup behaviors of an ECS instance and to pass data into an ECS instance. If updated, the instance will reboot to make the change take effect. Note: Not all of changes will take effect and it depends on [cloud-init module type](https://cloudinit.readthedocs.io/en/latest/topics/modules.html).`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `(Optional, Force new resource) The name of key pair that can login ECS instance successfully without password. If it is specified, the password would be invalid.`,
				},
				resource.Attribute{
					Name:        "role_name",
					Description: `(Optional, Force new resource) Instance RAM role name. The name is provided and maintained by RAM. You can use ` + "`" + `alibabacloudstack_ram_role` + "`" + ` to create a new one.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `(Optional) Instance private IP address can be specified when you creating new instance. It is valid when ` + "`" + `vswitch_id` + "`" + ` is specified. When it is changed, the instance will reboot to make the change take effect. Default to NoSpot. Note: Currently, the spot instance only supports domestic site account. Default to false.`,
				},
				resource.Attribute{
					Name:        "security_enhancement_strategy",
					Description: `(Optional, ForceNew) The security enhancement strategy. - Active: Enable security enhancement strategy, it only works on system images. - Deactive: Disable security enhancement strategy, it works on all images.`,
				},
				resource.Attribute{
					Name:        "data_disks",
					Description: `(Optional, ForceNew) The list of data disks created with instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, ForceNew) The name of the data disk.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required, ForceNew) The size of the data disk. - cloud：[5, 2000] - cloud_efficiency：[20, 32768] - cloud_ssd：[20, 32768] - cloud_essd：[20, 32768] - ephemeral_ssd: [5, 800]`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `(Optional, ForceNew) The category of the disk: - ` + "`" + `cloud` + "`" + `: The general cloud disk. - ` + "`" + `cloud_efficiency` + "`" + `: The efficiency cloud disk. - ` + "`" + `cloud_ssd` + "`" + `: The SSD cloud disk. - ` + "`" + `ephemeral_ssd` + "`" + `: The local SSD disk. Default to ` + "`" + `cloud_efficiency` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional, ForceNew) The snapshot ID used to initialize the data disk. If the size specified by snapshot is greater that the size of the disk, use the size specified by snapshot as the size of the data disk.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, ForceNew) The description must be 2 to 256 characters in length.`,
				},
				resource.Attribute{
					Name:        "delete_with_instance",
					Description: `(Optional, ForceNew) Delete this data disk when the instance is destroyed. It only works on cloud, cloud_efficiency, cloud_essd, cloud_ssd disk. If the category of this data disk was ephemeral_ssd, please don't set this param. Default to true ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 10 mins) Used when creating the instance (until it reaches the initial ` + "`" + `Running` + "`" + ` status). ` + "`" + `Note` + "`" + `: There are extra at most 2 minutes used to retry to aviod some needless API errors and it is not in the timeouts configure.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 10 mins) Used when stopping and starting the instance when necessary during update - e.g. when changing instance type, password, image, vswitch and private IP.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 20 mins) Used when terminating the instance. ` + "`" + `Note` + "`" + `: There are extra at most 5 minutes used to retry to aviod some needless API errors and it is not in the timeouts configure. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The instance ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The instance status.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The instance private ip.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The instance ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The instance status.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The instance private ip.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_key_pair",
			Category:         "ECS",
			ShortDescription: `Provides a AlibabacloudStack key pair resource.`,
			Description:      ``,
			Keywords: []string{
				"ecs",
				"key",
				"pair",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key_name",
					Description: `(Required, ForceNew) The key pair's name.The name must be unique.`,
				},
				resource.Attribute{
					Name:        "key_name_prefix",
					Description: `(ForceNew) The key pair name's prefix. It is conflict with ` + "`" + `key_name` + "`" + `. If it is specified, terraform will using it to build the only key name.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(ForceNew) You can import an existing public key and using AlibabacloudStack key pair to manage it.`,
				},
				resource.Attribute{
					Name:        "key_file",
					Description: `(ForceNew) The name of file to save your new key pair's private key. Strongly suggest you to specified it when you creating key pair, otherwise, you wouldn't get its private key ever.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. ->`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `The name of the key pair.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "key_name",
					Description: `The name of the key pair.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_key_pair_attachment",
			Category:         "ECS",
			ShortDescription: `Provides a AlibabacloudStack key pair attachment resource to bind key pair for several ECS instances.`,
			Description:      ``,
			Keywords: []string{
				"ecs",
				"key",
				"pair",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key_name",
					Description: `(Required, ForceNew) The name of key pair used to bind.`,
				},
				resource.Attribute{
					Name:        "instance_ids",
					Description: `(Required, ForceNew) The list of ECS instance's IDs.`,
				},
				resource.Attribute{
					Name:        "force",
					Description: `(ForceNew) Set it to true and it will reboot instances which attached with the key pair to make key pair affect immediately. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `The name of the key pair.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "key_name",
					Description: `The name of the key pair.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_kms_alias",
			Category:         "KMS",
			ShortDescription: `Provides a AlibabacloudStack KMS Alias resource.`,
			Description:      ``,
			Keywords: []string{
				"kms",
				"alias",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "alias_name",
					Description: `(Required, ForceNew) The alias of CMK. ` + "`" + `Encrypt` + "`" + `、` + "`" + `GenerateDataKey` + "`" + `、` + "`" + `DescribeKey` + "`" + ` can be called using aliases. Length of characters other than prefixes: minimum length of 1 character and maximum length of 255 characters. Must contain prefix ` + "`" + `alias/` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `(Required) The id of the key. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the alias.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the alias.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_kms_ciphertext",
			Category:         "KMS",
			ShortDescription: `Encrypt data with KMS.`,
			Description:      ``,
			Keywords: []string{
				"kms",
				"ciphertext",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "plaintext",
					Description: `(ForceNew) The plaintext to be encrypted which must be encoded in Base64.`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `(ForceNew) The globally unique ID of the CMK.`,
				},
				resource.Attribute{
					Name:        "encryption_context",
					Description: `(Optional, ForceNew) The Encryption context. If you specify this parameter here, it is also required when you call the Decrypt API operation. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
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
			Type:             "alibabacloudstack_kms_key",
			Category:         "KMS",
			ShortDescription: `Provides a Alibabacloudstackms key resource.`,
			Description:      ``,
			Keywords: []string{
				"kms",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the key as viewed in Alibabacloudstack console.`,
				},
				resource.Attribute{
					Name:        "key_usage",
					Description: `(Optional, ForceNew) Specifies the usage of CMK. Currently, default to 'ENCRYPT/DECRYPT', indicating that CMK is used for encryption and decryption.`,
				},
				resource.Attribute{
					Name:        "automatic_rotation",
					Description: `(Optional) Specifies whether to enable automatic key rotation. Default:"Disabled".`,
				},
				resource.Attribute{
					Name:        "key_state",
					Description: `(Optional) The status of CMK. Defaults to Enabled.`,
				},
				resource.Attribute{
					Name:        "origin",
					Description: `(Optional, ForceNew) The source of the key material for the CMK. Defaults to "Aliyun_KMS".`,
				},
				resource.Attribute{
					Name:        "pending_window_in_days",
					Description: `(Optional) Duration in days after which the key is deleted after destruction of the resource, must be between 7 and 30 days. Defaults to 30 days.`,
				},
				resource.Attribute{
					Name:        "protection_level",
					Description: `(Optional, ForceNew) The protection level of the CMK. Defaults to "SOFTWARE".`,
				},
				resource.Attribute{
					Name:        "rotation_interval",
					Description: `(Optional) The period of automatic key rotation. Unit: seconds. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the key.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The Alibabacloudstack Resource Name (ARN) of the key.`,
				},
				resource.Attribute{
					Name:        "last_rotation_date",
					Description: `The date and time the last rotation was performed. The time is displayed in UTC.`,
				},
				resource.Attribute{
					Name:        "material_expire_time",
					Description: `The time and date the key material for the CMK expires. The time is displayed in UTC. If the value is empty, the key material for the CMK does not expire.`,
				},
				resource.Attribute{
					Name:        "next_rotation_date",
					Description: `The time the next rotation is scheduled for execution.`,
				},
				resource.Attribute{
					Name:        "primary_key_version",
					Description: `The ID of the current primary key version of the symmetric CMK.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the key.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The Alibabacloudstack Resource Name (ARN) of the key.`,
				},
				resource.Attribute{
					Name:        "last_rotation_date",
					Description: `The date and time the last rotation was performed. The time is displayed in UTC.`,
				},
				resource.Attribute{
					Name:        "material_expire_time",
					Description: `The time and date the key material for the CMK expires. The time is displayed in UTC. If the value is empty, the key material for the CMK does not expire.`,
				},
				resource.Attribute{
					Name:        "next_rotation_date",
					Description: `The time the next rotation is scheduled for execution.`,
				},
				resource.Attribute{
					Name:        "primary_key_version",
					Description: `The ID of the current primary key version of the symmetric CMK.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_kms_secret",
			Category:         "KMS",
			ShortDescription: `Provides a Alibabacloudstack Cloud kms secret resource.`,
			Description:      ``,
			Keywords: []string{
				"kms",
				"secret",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the secret.`,
				},
				resource.Attribute{
					Name:        "encryption_key_id",
					Description: `(Optional, ForceNew) The ID of the KMS CMK that is used to encrypt the secret value. If you do not specify this parameter, Secrets Manager automatically creates an encryption key to encrypt the secret.`,
				},
				resource.Attribute{
					Name:        "force_delete_without_recovery",
					Description: `(Optional) Specifies whether to forcibly delete the secret. If this parameter is set to true, the secret cannot be recovered. Valid values: true, false. Default to: false.`,
				},
				resource.Attribute{
					Name:        "recovery_window_in_days",
					Description: `(Optional) Specifies the recovery period of the secret if you do not forcibly delete it. Default value: 30. It will be ignored when ` + "`" + `force_delete_without_recovery` + "`" + ` is true.`,
				},
				resource.Attribute{
					Name:        "secret_data",
					Description: `(Required) The value of the secret that you want to create. Secrets Manager encrypts the secret value and stores it in the initial version.`,
				},
				resource.Attribute{
					Name:        "secret_data_type",
					Description: `(Optional) The type of the secret value. Valid values: text, binary. Default to "text".`,
				},
				resource.Attribute{
					Name:        "secret_name",
					Description: `(Required, ForceNew) The name of the secret.`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `(Required) The version number of the initial version. Version numbers are unique in each secret object.`,
				},
				resource.Attribute{
					Name:        "version_stages",
					Description: `(Optional, List(string)) The stage labels that mark the new secret version. If you do not specify this parameter, Secrets Manager marks it with "ACSCurrent".`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the secret. It same with ` + "`" + `secret_name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The Alibabacloudstack Resource Name (ARN) of the secret.`,
				},
				resource.Attribute{
					Name:        "planned_delete_time",
					Description: `The time when the secret is scheduled to be deleted.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the secret. It same with ` + "`" + `secret_name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The Alibabacloudstack Resource Name (ARN) of the secret.`,
				},
				resource.Attribute{
					Name:        "planned_delete_time",
					Description: `The time when the secret is scheduled to be deleted.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_kvstore_account",
			Category:         "Redis And Memcache (KVStore)",
			ShortDescription: `Provides a kvstore account resource.`,
			Description:      ``,
			Keywords: []string{
				"redis",
				"and",
				"memcache",
				"kvstore",
				"account",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, ForceNew) The Id of instance in which account belongs. (The engine version of instance must be 4.0 or 4.0+)`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `(Required, ForceNew) Operation account requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter and have no more than 16 characters.`,
				},
				resource.Attribute{
					Name:        "account_password",
					Description: `(Optional, Sensitive) Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters. You have to specify one of ` + "`" + `account_password` + "`" + ` and ` + "`" + `kms_encrypted_password` + "`" + ` fields.`,
				},
				resource.Attribute{
					Name:        "kms_encrypted_password",
					Description: `(Optional) An KMS encrypts password used to a KVStore account. If the ` + "`" + `account_password` + "`" + ` is filled in, this field will be ignored.`,
				},
				resource.Attribute{
					Name:        "kms_encryption_context",
					Description: `(Optional) An KMS encryption context used to decrypt ` + "`" + `kms_encrypted_password` + "`" + ` before creating or updating a KVStore account with ` + "`" + `kms_encrypted_password` + "`" + `. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when ` + "`" + `kms_encrypted_password` + "`" + ` is set.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.`,
				},
				resource.Attribute{
					Name:        "account_type",
					Description: `(Optional, ForceNew)Privilege type of account. - Normal: Common privilege. Default to Normal.`,
				},
				resource.Attribute{
					Name:        "account_privilege",
					Description: `(Optional) The privilege of account access database. Valid values: - RoleReadOnly: This value is only for Redis and Memcache - RoleReadWrite: This value is only for Redis and Memcache - RoleRepl: This value supports instance to read, write, and open SYNC / PSYNC commands. Only for Redis which engine version is 4.0 and architecture type is standard Default to "RoleReadWrite". ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The current account resource ID. Composed of instance ID and account name with format ` + "`" + `<instance_id>:<name>` + "`" + `. ## Import kvstore account can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_KVStore_account.example "rm-12345:tf_account" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The current account resource ID. Composed of instance ID and account name with format ` + "`" + `<instance_id>:<name>` + "`" + `. ## Import kvstore account can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_KVStore_account.example "rm-12345:tf_account" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_kvstore_backup_policy",
			Category:         "Redis And Memcache (KVStore)",
			ShortDescription: `Provides a backup policy for ApsaraDB Redis / Memcache instance resource.`,
			Description:      ``,
			Keywords: []string{
				"redis",
				"and",
				"memcache",
				"kvstore",
				"backup",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, ForceNew) The id of ApsaraDB for Redis or Memcache intance.`,
				},
				resource.Attribute{
					Name:        "backup_time",
					Description: `(Optional) Backup time, in the format of HH:mmZ- HH:mm Z`,
				},
				resource.Attribute{
					Name:        "backup_period",
					Description: `(Optional) Backup Cycle. Allowed values: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the backup policy.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The id of ApsaraDB for Redis or Memcache intance.`,
				},
				resource.Attribute{
					Name:        "backup_time",
					Description: `Backup time, in the format of HH:mmZ- HH:mm Z`,
				},
				resource.Attribute{
					Name:        "backup_period",
					Description: `Backup Cycle. Allowed values: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday ## Import KVStore backup policy can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_kvstore_backup_policy.example r-abc12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the backup policy.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The id of ApsaraDB for Redis or Memcache intance.`,
				},
				resource.Attribute{
					Name:        "backup_time",
					Description: `Backup time, in the format of HH:mmZ- HH:mm Z`,
				},
				resource.Attribute{
					Name:        "backup_period",
					Description: `Backup Cycle. Allowed values: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday ## Import KVStore backup policy can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_kvstore_backup_policy.example r-abc12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_kvstore_connection",
			Category:         "Redis And Memcache (KVStore)",
			ShortDescription: `Operate the public network ip of the specified resource.`,
			Description:      ``,
			Keywords: []string{
				"redis",
				"and",
				"memcache",
				"kvstore",
				"connection",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "connection_string_prefix",
					Description: `(Required) The prefix of the public endpoint. The prefix can be 8 to 64 characters in length, and can contain lowercase letters and digits. It must start with a lowercase letter.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) The service port number of the instance. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of KVStore DBInstance.`,
				},
				resource.Attribute{
					Name:        "connection_string",
					Description: `The public connection string of KVStore DBInstance. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 2 mins) Used when creating the KVStore connection (until it reaches the initial ` + "`" + `Normal` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 2 mins) Used when updating the KVStore connection (until it reaches the initial ` + "`" + `Normal` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 2 mins) Used when deleting the KVStore connection (until it reaches the initial ` + "`" + `Normal` + "`" + ` status). ## Import KVStore connection can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_kvstore_connection.example r-abc12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of KVStore DBInstance.`,
				},
				resource.Attribute{
					Name:        "connection_string",
					Description: `The public connection string of KVStore DBInstance. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 2 mins) Used when creating the KVStore connection (until it reaches the initial ` + "`" + `Normal` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 2 mins) Used when updating the KVStore connection (until it reaches the initial ` + "`" + `Normal` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 2 mins) Used when deleting the KVStore connection (until it reaches the initial ` + "`" + `Normal` + "`" + ` status). ## Import KVStore connection can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_kvstore_connection.example r-abc12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_kvstore_instance",
			Category:         "Redis And Memcache (KVStore)",
			ShortDescription: `Provides an ApsaraDB Redis / Memcache instance resource.`,
			Description:      ``,
			Keywords: []string{
				"redis",
				"and",
				"memcache",
				"kvstore",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Optional) The name of DB instance. It a string of 2 to 256 characters.`,
				},
				resource.Attribute{
					Name:        "kms_encrypted_password",
					Description: `(Optional) An KMS encrypts password used to a instance. If the ` + "`" + `password` + "`" + ` is filled in, this field will be ignored.`,
				},
				resource.Attribute{
					Name:        "kms_encryption_context",
					Description: `(Optional, MapString) An KMS encryption context used to decrypt ` + "`" + `kms_encrypted_password` + "`" + ` before creating or updating instance with ` + "`" + `kms_encrypted_password` + "`" + `. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when ` + "`" + `kms_encrypted_password` + "`" + ` is set.`,
				},
				resource.Attribute{
					Name:        "instance_class",
					Description: `(Required) Type of the applied ApsaraDB for Redis instance. It can be retrieved by data source [` + "`" + `alibabacloudstack_kvstore_instance_classes` + "`" + `](https://www.terraform.io/docs/providers/alibabacloudstack/d/kvstore_instance_classes.html)`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional, ForceNew) The Zone to launch the DB instance.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `(Optional) Valid values are ` + "`" + `PrePaid` + "`" + `, ` + "`" + `PostPaid` + "`" + `, Default to ` + "`" + `PostPaid` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) The duration that you will buy DB instance (in month). It is valid when instance_charge_type is ` + "`" + `PrePaid` + "`" + `. Valid values: [1~9], 12, 24, 36. Default to 1.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Optional, ForceNew) The engine to use: ` + "`" + `Redis` + "`" + ` or ` + "`" + `Memcache` + "`" + `. Defaults to ` + "`" + `Redis` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Optional, ForceNew) The ID of VSwitch.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Optional) The Security Group ID of ECS.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `(Optional) Set of parameters needs to be set after instance was launched. Available parameters can refer to the latest docs [Instance configurations table](https://www.alibabacloud.com/help/doc-detail/61209.htm) .`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "maintain_start_time",
					Description: `(Optional) The start time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).`,
				},
				resource.Attribute{
					Name:        "maintain_end_time",
					Description: `(Optional) The end time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).`,
				},
				resource.Attribute{
					Name:        "architecture_type",
					Description: `(Required) The architecture type of the resource. Valid values: ` + "`" + `standard` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cpu_type",
					Description: `(Required) The cpu type of the resource.Valid values: ` + "`" + `intel` + "`" + `. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The KVStore instance ID.`,
				},
				resource.Attribute{
					Name:        "connection_domain",
					Description: `Instance connection domain (only Intranet access supported). ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 20 mins) Used when creating the KVStore instance (until it reaches the initial ` + "`" + `Normal` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 30 mins) Used when updating the KVStore instance (until it reaches the initial ` + "`" + `Normal` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 20 mins) Used when terminating the KVStore instance. ## Import KVStore instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_kvstore_instance.example r-abc12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The KVStore instance ID.`,
				},
				resource.Attribute{
					Name:        "connection_domain",
					Description: `Instance connection domain (only Intranet access supported). ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 20 mins) Used when creating the KVStore instance (until it reaches the initial ` + "`" + `Normal` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 30 mins) Used when updating the KVStore instance (until it reaches the initial ` + "`" + `Normal` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 20 mins) Used when terminating the KVStore instance. ## Import KVStore instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_kvstore_instance.example r-abc12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_launch_template",
			Category:         "ECS",
			ShortDescription: `Provides an ECS Launch Template resource.`,
			Description:      ``,
			Keywords: []string{
				"ecs",
				"launch",
				"template",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, ForceNew) Instance launch template name. Can contain [2, 128] characters in length. It must start with an English letter (uppercase or lowercase) and can contain numbers, periods (.), colons (:), underscores (_), and hyphens (-). It cannot start with "http://" or "https://".`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of instance launch template version 1. It can be [2, 256] characters in length. It cannot start with "http://" or "https://". The default value is null.`,
				},
				resource.Attribute{
					Name:        "host_name",
					Description: `(Optional) Instance host name.It cannot start or end with a period (.) or a hyphen (-) and it cannot have two or more consecutive periods (.) or hyphens (-).For Windows: The host name can be [2, 15] characters in length. It can contain A-Z, a-z, numbers, periods (.), and hyphens (-). It cannot only contain numbers. For other operating systems: The host name can be [2, 64] characters in length. It can be segments separated by periods (.). It can contain A-Z, a-z, numbers, and hyphens (-).`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Optional) Image ID.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Optional) The name of the instance. The name is a string of 2 to 128 characters. It must begin with an English or a Chinese character. It can contain A-Z, a-z, Chinese characters, numbers, periods (.), colons (:), underscores (_), and hyphens (-).`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Optional) Instance type. For more information, call resource_alibabacloudstack_instances to obtain the latest instance type list.`,
				},
				resource.Attribute{
					Name:        "auto_release_time",
					Description: `(Optional) Instance auto release time. The time is presented using the ISO8601 standard and in UTC time. The format is YYYY-MM-DDTHH:MM:SSZ.`,
				},
				resource.Attribute{
					Name:        "internet_max_bandwidth_in",
					Description: `(Optional) The maximum inbound bandwidth from the Internet network, measured in Mbit/s. Value range: [1, 200].`,
				},
				resource.Attribute{
					Name:        "internet_max_bandwidth_out",
					Description: `(Optional) Maximum outbound bandwidth from the Internet, its unit of measurement is Mbit/s. Value range: [0, 100].`,
				},
				resource.Attribute{
					Name:        "io_optimized",
					Description: `(Optional) Whether it is an I/O-optimized instance or not. Optional values: - none - optimized`,
				},
				resource.Attribute{
					Name:        "key_pair_name",
					Description: `(Optional) The name of the key pair. - Ignore this parameter for Windows instances. It is null by default. Even if you enter this parameter, only the Password content is used. - The password logon method for Linux instances is set to forbidden upon initialization.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `(Optional) Network type of the instance. Value options: Classic | VPC.`,
				},
				resource.Attribute{
					Name:        "ram_role_name",
					Description: `(Optional) The RAM role name of the instance. You can use the RAM API ListRoles to query instance RAM role names.`,
				},
				resource.Attribute{
					Name:        "security_enhancement_strategy",
					Description: `(Optional) Whether or not to activate the security enhancement feature and install network security software free of charge. Optional values: Active | Deactive.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Optional) The security group ID.`,
				},
				resource.Attribute{
					Name:        "spot_strategy",
					Description: `(Optional) The spot strategy for a Pay-As-You-Go instance. This parameter is valid and required only when InstanceChargeType is set to PostPaid. Value range: - NoSpot: Normal Pay-As-You-Go instance. - SpotWithPriceLimit: Sets the maximum price for a spot instance. - SpotAsPriceGo: The system automatically calculates the price. The maximum value is the Pay-As-You-Go price.`,
				},
				resource.Attribute{
					Name:        "system_disk_category",
					Description: `(Optional) The category of the system disk. System disk type. Optional values: - cloud: Basic cloud disk. - cloud_efficiency: Ultra cloud disk. - cloud_ssd: SSD cloud Disks. - ephemeral_ssd: local SSD Disks - cloud_essd: ESSD cloud Disks.`,
				},
				resource.Attribute{
					Name:        "system_disk_description",
					Description: `(Optional) System disk description. It cannot begin with http:// or https://.`,
				},
				resource.Attribute{
					Name:        "system_disk_name",
					Description: `(Optional) System disk name. The name is a string of 2 to 128 characters. It must begin with an English or a Chinese character. It can contain A-Z, a-z, Chinese characters, numbers, periods (.), colons (:), underscores (_), and hyphens (-).`,
				},
				resource.Attribute{
					Name:        "system_disk_size",
					Description: `(Optional) Size of the system disk, measured in GB. Value range: [20, 500].`,
				},
				resource.Attribute{
					Name:        "userdata",
					Description: `(Optional) User data of the instance, which is Base64-encoded. Size of the raw data cannot exceed 16 KB.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Optional) When creating a VPC-Connected instance, you must specify its VSwitch ID.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Optional) The zone ID of the instance.`,
				},
				resource.Attribute{
					Name:        "network_interfaces",
					Description: `(Optional) The list of network interfaces created with instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) ENI name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The ENI description.`,
				},
				resource.Attribute{
					Name:        "primary_ip",
					Description: `(Optional) The primary private IP address of the ENI.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Optional) The security group ID must be one in the same VPC.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Optional) The VSwitch ID for ENI. The instance must be in the same zone of the same VPC network as the ENI, but they may belong to different VSwitches.`,
				},
				resource.Attribute{
					Name:        "data_disks",
					Description: `(Optional) The list of data disks created with instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the data disk.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) The size of the data disk. - cloud：[5, 2000] - cloud_efficiency：[20, 32768] - cloud_ssd：[20, 32768] - cloud_essd：[20, 32768] - ephemeral_ssd: [5, 800]`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `(Optional) The category of the disk: - cloud: Basic cloud disk. - cloud_efficiency: Ultra cloud disk. - cloud_ssd: SSD cloud Disks. - ephemeral_ssd: local SSD Disks - cloud_essd: ESSD cloud Disks. Default to ` + "`" + `cloud_efficiency` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) The snapshot ID used to initialize the data disk. If the size specified by snapshot is greater that the size of the disk, use the size specified by snapshot as the size of the data disk.`,
				},
				resource.Attribute{
					Name:        "delete_with_instance",
					Description: `(Optional) Delete this data disk when the instance is destroyed. It only works on cloud, cloud_efficiency, cloud_ssd and cloud_essd disk. If the category of this data disk was ephemeral_ssd, please don't set this param. Default to true`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the data disk.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string. - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Launch Template ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Launch Template ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_log_machine_group",
			Category:         "Log Service (SLS)",
			ShortDescription: `Provides a Alibabacloudstack log tail machine group resource.`,
			Description:      ``,
			Keywords: []string{
				"log",
				"service",
				"sls",
				"machine",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project",
					Description: `(Required, ForceNew) The project name to the machine group belongs.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) The machine group name, which is unique in the same project.`,
				},
				resource.Attribute{
					Name:        "identify_type",
					Description: `(Optional) The machine identification type, including IP and user-defined identity. Valid values are "ip" and "userdefined". Default to "ip".`,
				},
				resource.Attribute{
					Name:        "topic",
					Description: `(Optional) The topic of a machine group. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the log machine group. It formats of ` + "`" + `<project>:<name>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `The project name.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The machine group name.`,
				},
				resource.Attribute{
					Name:        "identify_type",
					Description: `The machine identification type.`,
				},
				resource.Attribute{
					Name:        "identify_list",
					Description: `The machine identification.`,
				},
				resource.Attribute{
					Name:        "topic",
					Description: `The machine group topic. ## Import Log machine group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_log_machine_group.example tf-log:tf-machine-group ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the log machine group. It formats of ` + "`" + `<project>:<name>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `The project name.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The machine group name.`,
				},
				resource.Attribute{
					Name:        "identify_type",
					Description: `The machine identification type.`,
				},
				resource.Attribute{
					Name:        "identify_list",
					Description: `The machine identification.`,
				},
				resource.Attribute{
					Name:        "topic",
					Description: `The machine group topic. ## Import Log machine group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_log_machine_group.example tf-log:tf-machine-group ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_log_project",
			Category:         "Log Service (SLS)",
			ShortDescription: `Provides a Alibabacloudstack log project resource.`,
			Description:      ``,
			Keywords: []string{
				"log",
				"service",
				"sls",
				"project",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) The name of the log project. It is the only in one Alibabacloudstack account.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the log project. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the log project. It same as its name.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Log project name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Log project description.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the log project. It same as its name.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Log project name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Log project description.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_log_store",
			Category:         "Log Service (SLS)",
			ShortDescription: `Provides a Alibabacloudstack log store resource.`,
			Description:      ``,
			Keywords: []string{
				"log",
				"service",
				"sls",
				"store",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project",
					Description: `(Required, ForceNew) The project name to the log store belongs.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) The log store, which is unique in the same project.`,
				},
				resource.Attribute{
					Name:        "retention_period",
					Description: `(Optional) The data retention time (in days). Valid values: [1-3650]. Default to ` + "`" + `30` + "`" + `. Log store data will be stored permanently when the value is ` + "`" + `3650` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "shard_count",
					Description: `(Optional) The number of shards in this log store. Default to 2. You can modify it by "Split" or "Merge" operations. [Refer to details](https://www.alibabacloud.com/help/doc-detail/28976.htm)`,
				},
				resource.Attribute{
					Name:        "auto_split",
					Description: `(Optional) Determines whether to automatically split a shard. Default to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max_split_shard_count",
					Description: `(Optional) The maximum number of shards for automatic split, which is in the range of 1 to 64. You must specify this parameter when autoSplit is true.`,
				},
				resource.Attribute{
					Name:        "append_meta",
					Description: `(Optional) Determines whether to append log meta automatically. The meta includes log receive time and client IP address. Default to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enable_web_tracking",
					Description: `(Optional) Determines whether to enable Web Tracking. Default ` + "`" + `false` + "`" + `. <<<<<<< HEAD`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the log project. It formats of ` + "`" + `<project>:<name>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `The project name.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Log store name.`,
				},
				resource.Attribute{
					Name:        "retention_period",
					Description: `The data retention time.`,
				},
				resource.Attribute{
					Name:        "shard_count",
					Description: `The number of shards.`,
				},
				resource.Attribute{
					Name:        "auto_split",
					Description: `Determines whether to automatically split a shard.`,
				},
				resource.Attribute{
					Name:        "max_split_shard_count",
					Description: `The maximum number of shards for automatic split.`,
				},
				resource.Attribute{
					Name:        "append_meta",
					Description: `Determines whether to append log meta automatically.`,
				},
				resource.Attribute{
					Name:        "enable_web_tracking",
					Description: `Determines whether to enable Web Tracking. <<<<<<< HEAD`,
				},
				resource.Attribute{
					Name:        "encryption",
					Description: `Determines whether to automatically encryption. =======`,
				},
				resource.Attribute{
					Name:        "encrypt_conf",
					Description: `Encryption configuration of logstore. >>>>>>> 153c5e75 (add 314 r doc) ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the log project. It formats of ` + "`" + `<project>:<name>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `The project name.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Log store name.`,
				},
				resource.Attribute{
					Name:        "retention_period",
					Description: `The data retention time.`,
				},
				resource.Attribute{
					Name:        "shard_count",
					Description: `The number of shards.`,
				},
				resource.Attribute{
					Name:        "auto_split",
					Description: `Determines whether to automatically split a shard.`,
				},
				resource.Attribute{
					Name:        "max_split_shard_count",
					Description: `The maximum number of shards for automatic split.`,
				},
				resource.Attribute{
					Name:        "append_meta",
					Description: `Determines whether to append log meta automatically.`,
				},
				resource.Attribute{
					Name:        "enable_web_tracking",
					Description: `Determines whether to enable Web Tracking. <<<<<<< HEAD`,
				},
				resource.Attribute{
					Name:        "encryption",
					Description: `Determines whether to automatically encryption. =======`,
				},
				resource.Attribute{
					Name:        "encrypt_conf",
					Description: `Encryption configuration of logstore. >>>>>>> 153c5e75 (add 314 r doc) ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_log_store_index",
			Category:         "Log Service (SLS)",
			ShortDescription: `Provides a Alibabacloudstack log store index resource.`,
			Description:      ``,
			Keywords: []string{
				"log",
				"service",
				"sls",
				"store",
				"index",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project",
					Description: `(Required, ForceNew) The project name to the log store belongs.`,
				},
				resource.Attribute{
					Name:        "logstore",
					Description: `(Required, ForceNew) The log store name to the query index belongs.`,
				},
				resource.Attribute{
					Name:        "full_text",
					Description: `The configuration of full text index. Valid item as follows:`,
				},
				resource.Attribute{
					Name:        "case_sensitive",
					Description: `(Optional) Whether the case sensitive. Default to false.`,
				},
				resource.Attribute{
					Name:        "include_chinese",
					Description: `(Optional) Whether includes the chinese. Default to false.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Optional) The string of several split words, like "\r", "#"`,
				},
				resource.Attribute{
					Name:        "field_search",
					Description: `List configurations of field search index. Valid item as follows:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The field name, which is unique in the same log store.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type of one field. Valid values: ["long", "text", "double", "json"]. Default to "long".`,
				},
				resource.Attribute{
					Name:        "alias",
					Description: `(Optional) The alias of one field`,
				},
				resource.Attribute{
					Name:        "case_sensitive",
					Description: `(Optional) Whether the case sensitive for the field. Default to false. It is valid when "type" is "text" or "json".`,
				},
				resource.Attribute{
					Name:        "include_chinese",
					Description: `(Optional) Whether includes the chinese for the field. Default to false. It is valid when "type" is "text" or "json".`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Optional) The string of several split words, like "\r", "#". It is valid when "type" is "text" or "json".`,
				},
				resource.Attribute{
					Name:        "enable_analytics",
					Description: `(Optional) Whether to enable field analytics. Default to true.`,
				},
				resource.Attribute{
					Name:        "json_keys",
					Description: `(Optional, Available in 1.66.0+) Use nested index when type is json`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) When using the json_keys field, this field is required.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type of one field. Valid values: ["long", "text", "double"]. Default to "long"`,
				},
				resource.Attribute{
					Name:        "alias",
					Description: `(Optional) The alias of one field.`,
				},
				resource.Attribute{
					Name:        "doc_value",
					Description: `(Optional) Whether to enable statistics. default to true. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the log store index. It formats of ` + "`" + `<project>:<logstore>` + "`" + `. ## Import Log store index can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_log_store_index.example tf-log:tf-log-store ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the log store index. It formats of ` + "`" + `<project>:<logstore>` + "`" + `. ## Import Log store index can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_log_store_index.example tf-log:tf-log-store ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_logtail_attachment",
			Category:         "Log Service (SLS)",
			ShortDescription: `Provides a Alibabacloudstack logtail attachment resource.`,
			Description:      ``,
			Keywords: []string{
				"log",
				"service",
				"sls",
				"logtail",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project",
					Description: `(Required, ForceNew) The project name to the log store belongs.`,
				},
				resource.Attribute{
					Name:        "logtail_config_name",
					Description: `(Required, ForceNew) The Logtail configuration name, which is unique in the same project.`,
				},
				resource.Attribute{
					Name:        "machine_group_name",
					Description: `(Required, ForceNew) The machine group name, which is unique in the same project. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the logtail to machine group. It formats of ` + "`" + `<project>:<logtail_config_name>:<machine_group_name>` + "`" + `. ## Import Logtial to machine group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_logtail_to_machine_group.example tf-log:tf-log-config:tf-log-machine-group ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the logtail to machine group. It formats of ` + "`" + `<project>:<logtail_config_name>:<machine_group_name>` + "`" + `. ## Import Logtial to machine group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_logtail_to_machine_group.example tf-log:tf-log-config:tf-log-machine-group ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_logtail_config",
			Category:         "Log Service (SLS)",
			ShortDescription: `Provides a Alibabacloudstack logtail config resource.`,
			Description:      ``,
			Keywords: []string{
				"log",
				"service",
				"sls",
				"logtail",
				"config",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project",
					Description: `(Required, ForceNew) The project name to the log store belongs.`,
				},
				resource.Attribute{
					Name:        "logstore",
					Description: `(Required, ForceNew) The log store name to the query index belongs.`,
				},
				resource.Attribute{
					Name:        "input_type",
					Description: `(Required) The input type. Currently only two types of files and plugin are supported.`,
				},
				resource.Attribute{
					Name:        "log_sample",
					Description: `（Optional）The log sample of the Logtail configuration. The log size cannot exceed 1,000 bytes.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) The Logtail configuration name, which is unique in the same project.`,
				},
				resource.Attribute{
					Name:        "output_type",
					Description: `(Required) The output type. Currently, only LogService is supported.`,
				},
				resource.Attribute{
					Name:        "input_detail",
					Description: `(Required) The logtail configure the required JSON files. ([Refer to details](https://www.alibabacloud.com/help/doc-detail/29058.htm)) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the log store index. It formats of ` + "`" + `<project>:<logstore>:<config_name>` + "`" + `. ## Import Logtial config can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_logtail_config.example tf-log:tf-log-store:tf-log-config ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the log store index. It formats of ` + "`" + `<project>:<logstore>:<config_name>` + "`" + `. ## Import Logtial config can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_logtail_config.example tf-log:tf-log-store:tf-log-config ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_maxcompute_cu",
			Category:         "MaxCompute",
			ShortDescription: `Provides a Alibabacloudstack maxcompute cu resource.`,
			Description:      ``,
			Keywords: []string{
				"maxcompute",
				"cu",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required, ForceNew) The id of the maxcompute cu.`,
				},
				resource.Attribute{
					Name:        "cu_name",
					Description: `(Required, ForceNew, Available in 1.110.0+) The name of the maxcompute cu.`,
				},
				resource.Attribute{
					Name:        "cu_num",
					Description: `(Required, ForceNew) The name of the maxcompute cu.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Required, ForceNew) The cluster name of the maxcompute cu. ## Import MaxCompute project can be imported using the`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_maxcompute_project",
			Category:         "MaxCompute",
			ShortDescription: `Provides a Alibabacloudstack maxcompute project resource.`,
			Description:      ``,
			Keywords: []string{
				"maxcompute",
				"project",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) It has been deprecated from provider version 1.110.0 and ` + "`" + `project_name` + "`" + ` instead.`,
				},
				resource.Attribute{
					Name:        "project_name",
					Description: `(Required, ForceNew, Available in 1.110.0+) The name of the maxcompute project.`,
				},
				resource.Attribute{
					Name:        "specification_type",
					Description: `(Required) The type of resource Specification, only ` + "`" + `OdpsStandard` + "`" + ` supported currently.`,
				},
				resource.Attribute{
					Name:        "order_type",
					Description: `(Required) The type of payment, only ` + "`" + `PayAsYouGo` + "`" + ` supported currently. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the maxcompute project. It is the same as its name. ## Import MaxCompute project can be imported using the`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the maxcompute project. It is the same as its name. ## Import MaxCompute project can be imported using the`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_maxcompute_user",
			Category:         "MaxCompute",
			ShortDescription: `Provides a Alibabacloudstack maxcompute user resource.`,
			Description:      ``,
			Keywords: []string{
				"maxcompute",
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "user_name",
					Description: `(Required, ForceNew) The name of the user that you want to create.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required, ForceNew) The description of the user that you want to create.`,
				},
				resource.Attribute{
					Name:        "organization_id",
					Description: `(Optional) The id of the organization. ## Attributes Reference ## Import MaxCompute project can be imported using the`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_mongodb_instance",
			Category:         "MongoDB",
			ShortDescription: `Provides a MongoDB instance resource supports replica set instances only. the MongoDB provides stable, reliable, and automatic scalable database services. It offers a full range of database solutions, such as disaster recovery, backup, recovery, monitoring, and alarms.`,
			Description:      ``,
			Keywords: []string{
				"mongodb",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "engine_version",
					Description: `(Required, ForceNew) The database engine version of the instance.`,
				},
				resource.Attribute{
					Name:        "db_instance_class",
					Description: `(Required) The instance type.`,
				},
				resource.Attribute{
					Name:        "db_instance_storage",
					Description: `(Required) The storage capacity of the instance. Valid values: 10 to 3000. The value must be a multiple of 10. Unit: GB.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Optional, ForceNew) The zone ID of the instance.`,
				},
				resource.Attribute{
					Name:        "backup_period",
					Description: `(Optional) MongoDB Instance backup period. It is required when backup_time was existed. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]. Default to [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday].`,
				},
				resource.Attribute{
					Name:        "backup_time",
					Description: `(Required, ForceNew) MongoDB instance backup time. It is required when backup_period was existed. In the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. If not set, the system will return a default, like "23:00Z-24:00Z".`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of DB instance. It's a string of 2 to 256 characters.`,
				},
				resource.Attribute{
					Name:        "audit_policy",
					Description: `(Optional) Enables or disables the audit log feature or set the log retention period for MongoDB instance.`,
				},
				resource.Attribute{
					Name:        "enable_audit_policy",
					Description: `(Required) Specifies whether the audit log feature is enabled.`,
				},
				resource.Attribute{
					Name:        "storage_period",
					Description: `(Optional) The number of days that audit logs are stored. Valid values: 1 to 365 days. Default value: 30 days.`,
				},
				resource.Attribute{
					Name:        "security_ip_list",
					Description: `(Optional) Description of the nat gateway, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Defaults to null.`,
				},
				resource.Attribute{
					Name:        "ssl_action",
					Description: `(Optional) A list of bandwidth packages for the nat gateway. Only support nat gateway created before 00:00 on November 4, 2017.`,
				},
				resource.Attribute{
					Name:        "tde_status",
					Description: `(Optional, ForceNew) The TDE(Transparent Data Encryption) status.`,
				},
				resource.Attribute{
					Name:        "connection_string",
					Description: `(Optional) The current connection string, which is to be modified.`,
				},
				resource.Attribute{
					Name:        "storage_engine",
					Description: `(Optional, ForceNew) Storage engine: WiredTiger or RocksDB. System Default value: WiredTiger.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `(Optional) Valid values are PrePaid, PostPaid, System default to PostPaid.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) The duration that you will buy DB instance (in month). It is valid when instance_charge_type is PrePaid. Valid values: [1~9], 12, 24, 36. System default to 1.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Optional, ForceNew) The virtual switch ID to launch DB instances in one VPC.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Optional) The Security Group ID of ECS.`,
				},
				resource.Attribute{
					Name:        "account_password",
					Description: `(Optional) Password of the root account. It is a string of 6 to 32 characters and is composed of letters, numbers, and underlines.`,
				},
				resource.Attribute{
					Name:        "kms_encrypted_password",
					Description: `(Optional) An KMS encrypts password used to a instance. If the account_password is filled in, this field will be ignored.`,
				},
				resource.Attribute{
					Name:        "kms_encryption_context",
					Description: `(Optional) An KMS encryption context used to decrypt kms_encrypted_password before creating or updating instance with kms_encrypted_password`,
				},
				resource.Attribute{
					Name:        "ssl_action",
					Description: `(Optional) Actions performed on SSL functions, Valid values: Open: turn on SSL encryption; Close: turn off SSL encryption; Update: update SSL certificate.`,
				},
				resource.Attribute{
					Name:        "maintain_start_time",
					Description: `(Optional) The start time of the maintenance window. Specify the time in the HH:mmZ format. The time must be in UTC.`,
				},
				resource.Attribute{
					Name:        "maintain_end_time",
					Description: `(Optional) The end time of the maintenance window. Specify the time in the HH:mmZ format. The time must be in UTC. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "retention_period",
					Description: `Instance log backup retention days.`,
				},
				resource.Attribute{
					Name:        "replica_set_name",
					Description: `The name of the mongo replica set.`,
				},
				resource.Attribute{
					Name:        "maintain_start_time",
					Description: `The start time of the maintenance window.`,
				},
				resource.Attribute{
					Name:        "maintain_end_time",
					Description: `The end time of the maintenance window.`,
				},
				resource.Attribute{
					Name:        "ssl_status",
					Description: `Status of the SSL feature. Open: SSL is turned on; Closed: SSL is turned off.`,
				},
				resource.Attribute{
					Name:        "connection_string",
					Description: `The current connection string.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `The zone ID of the instance.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `The virtual switch ID to launch DB instances in one VPC.`,
				},
				resource.Attribute{
					Name:        "security_ip_list",
					Description: `Description of the nat gateway, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Defaults to null.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The Security Group ID of ECS.`,
				},
				resource.Attribute{
					Name:        "backup_period",
					Description: `MongoDB Instance backup period.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "retention_period",
					Description: `Instance log backup retention days.`,
				},
				resource.Attribute{
					Name:        "replica_set_name",
					Description: `The name of the mongo replica set.`,
				},
				resource.Attribute{
					Name:        "maintain_start_time",
					Description: `The start time of the maintenance window.`,
				},
				resource.Attribute{
					Name:        "maintain_end_time",
					Description: `The end time of the maintenance window.`,
				},
				resource.Attribute{
					Name:        "ssl_status",
					Description: `Status of the SSL feature. Open: SSL is turned on; Closed: SSL is turned off.`,
				},
				resource.Attribute{
					Name:        "connection_string",
					Description: `The current connection string.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `The zone ID of the instance.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `The virtual switch ID to launch DB instances in one VPC.`,
				},
				resource.Attribute{
					Name:        "security_ip_list",
					Description: `Description of the nat gateway, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Defaults to null.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The Security Group ID of ECS.`,
				},
				resource.Attribute{
					Name:        "backup_period",
					Description: `MongoDB Instance backup period.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_mongodb_sharding_instance",
			Category:         "MongoDB",
			ShortDescription: `Provides a MongoDB sharding instance resource.`,
			Description:      ``,
			Keywords: []string{
				"mongodb",
				"sharding",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "engine_version",
					Description: `(Required, ForceNew) Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/zh/doc-detail/61884.htm) ` + "`" + `EngineVersion` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of DB instance. It a string of 2 to 256 characters.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `(Optional) Valid values are ` + "`" + `PrePaid` + "`" + `, ` + "`" + `PostPaid` + "`" + `,System default to ` + "`" + `PostPaid` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) The duration that you will buy DB instance (in month). It is valid when instance_charge_type is ` + "`" + `PrePaid` + "`" + `. Valid values: [1~9], 12, 24, 36. System default to 1.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Optional, ForceNew) The Zone to launch the DB instance. MongoDB sharding instance does not support multiple-zone. If it is a multi-zone and ` + "`" + `vswitch_id` + "`" + ` is specified, the vswitch must in one of them.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Optional, ForceNew) The virtual switch ID to launch DB instances in one VPC.`,
				},
				resource.Attribute{
					Name:        "account_password",
					Description: `(Optional, Sensitive) Password of the root account. It is a string of 6 to 32 characters and is composed of letters, numbers, and underlines.`,
				},
				resource.Attribute{
					Name:        "kms_encrypted_password",
					Description: `(Optional, Available in 1.57.1+) An KMS encrypts password used to a instance. If the ` + "`" + `account_password` + "`" + ` is filled in, this field will be ignored.`,
				},
				resource.Attribute{
					Name:        "kms_encryption_context",
					Description: `(Optional, MapString, Available in 1.57.1+) An KMS encryption context used to decrypt ` + "`" + `kms_encrypted_password` + "`" + ` before creating or updating instance with ` + "`" + `kms_encrypted_password` + "`" + `. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when ` + "`" + `kms_encrypted_password` + "`" + ` is set.`,
				},
				resource.Attribute{
					Name:        "security_ip_list",
					Description: `(Optional) List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]). System default to ` + "`" + `["127.0.0.1"]` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Optional, Available in 1.76.0+) The Security Group ID of ECS.`,
				},
				resource.Attribute{
					Name:        "tde_status",
					Description: `(Optional, ForceNew, Available in 1.76.0+) The TDE(Transparent Data Encryption) status.`,
				},
				resource.Attribute{
					Name:        "mongo_list",
					Description: `(Required) The mongo-node count can be purchased is in range of [2, 32].`,
				},
				resource.Attribute{
					Name:        "shard_list",
					Description: `(Required) the shard-node count can be purchased is in range of [2, 32].`,
				},
				resource.Attribute{
					Name:        "node_storage",
					Description: `(Required) - Custom storage space; value range: [10, 1,000] - 10-GB increments. Unit: GB.`,
				},
				resource.Attribute{
					Name:        "readonly_replicas",
					Description: `(Optional, Available in 1.126.0+) The number of read-only nodes in shard node. Valid values: 0 to 5. Default value: 0.`,
				},
				resource.Attribute{
					Name:        "backup_period",
					Description: `(Optional, Available in 1.42.0+) MongoDB Instance backup period. It is required when ` + "`" + `backup_time` + "`" + ` was existed. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]. Default to [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]`,
				},
				resource.Attribute{
					Name:        "backup_time",
					Description: `(Optional, Available in 1.42.0+) MongoDB instance backup time. It is required when ` + "`" + `backup_period` + "`" + ` was existed. In the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. If not set, the system will return a default, like "23:00Z-24:00Z".`,
				},
				resource.Attribute{
					Name:        "order_type",
					Description: `(Optional, Available in v1.134.0+) The type of configuration changes performed. Default value: DOWNGRADE. Valid values:`,
				},
				resource.Attribute{
					Name:        "auto_renew",
					Description: `(Optional, Available in v1.141.0+) Auto renew for prepaid, true of false. Default is false. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the MongoDB.`,
				},
				resource.Attribute{
					Name:        "node_id",
					Description: `The ID of the mongo-node.`,
				},
				resource.Attribute{
					Name:        "connect_string",
					Description: `Mongo node connection string`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Mongo node port`,
				},
				resource.Attribute{
					Name:        "node_id",
					Description: `The ID of the shard-node.`,
				},
				resource.Attribute{
					Name:        "retention_period",
					Description: `Instance log backup retention days.`,
				},
				resource.Attribute{
					Name:        "config_server_list",
					Description: `The node information list of config server. The details see Block ` + "`" + `config_server_list` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max_iops",
					Description: `The maximum IOPS of the Config Server node.`,
				},
				resource.Attribute{
					Name:        "connect_string",
					Description: `The connection address of the Config Server node.`,
				},
				resource.Attribute{
					Name:        "node_class",
					Description: `The node class of the Config Server node.`,
				},
				resource.Attribute{
					Name:        "max_connections",
					Description: `The max connections of the Config Server node.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The connection port of the Config Server node.`,
				},
				resource.Attribute{
					Name:        "node_description",
					Description: `The description of the Config Server node.`,
				},
				resource.Attribute{
					Name:        "node_id",
					Description: `The ID of the Config Server node.`,
				},
				resource.Attribute{
					Name:        "node_storage",
					Description: `The node storage of the Config Server node. ### Timeouts ->`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 30 mins) Used when creating the MongoDB instance (until it reaches the initial ` + "`" + `Running` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 30 mins) Used when updating the MongoDB instance (until it reaches the initial ` + "`" + `Running` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 30 mins) Used when terminating the MongoDB instance. ## Import MongoDB can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_mongodb_sharding_instance.example dds-bp1291daeda44195 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the MongoDB.`,
				},
				resource.Attribute{
					Name:        "node_id",
					Description: `The ID of the mongo-node.`,
				},
				resource.Attribute{
					Name:        "connect_string",
					Description: `Mongo node connection string`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Mongo node port`,
				},
				resource.Attribute{
					Name:        "node_id",
					Description: `The ID of the shard-node.`,
				},
				resource.Attribute{
					Name:        "retention_period",
					Description: `Instance log backup retention days.`,
				},
				resource.Attribute{
					Name:        "config_server_list",
					Description: `The node information list of config server. The details see Block ` + "`" + `config_server_list` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max_iops",
					Description: `The maximum IOPS of the Config Server node.`,
				},
				resource.Attribute{
					Name:        "connect_string",
					Description: `The connection address of the Config Server node.`,
				},
				resource.Attribute{
					Name:        "node_class",
					Description: `The node class of the Config Server node.`,
				},
				resource.Attribute{
					Name:        "max_connections",
					Description: `The max connections of the Config Server node.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The connection port of the Config Server node.`,
				},
				resource.Attribute{
					Name:        "node_description",
					Description: `The description of the Config Server node.`,
				},
				resource.Attribute{
					Name:        "node_id",
					Description: `The ID of the Config Server node.`,
				},
				resource.Attribute{
					Name:        "node_storage",
					Description: `The node storage of the Config Server node. ### Timeouts ->`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 30 mins) Used when creating the MongoDB instance (until it reaches the initial ` + "`" + `Running` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 30 mins) Used when updating the MongoDB instance (until it reaches the initial ` + "`" + `Running` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 30 mins) Used when terminating the MongoDB instance. ## Import MongoDB can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_mongodb_sharding_instance.example dds-bp1291daeda44195 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_nas_access_group",
			Category:         "Network Attached Storage (NAS)",
			ShortDescription: `Provides a Alibabacloudstack NAS Access Group resource.`,
			Description:      ``,
			Keywords: []string{
				"network",
				"attached",
				"storage",
				"nas",
				"access",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew, Deprecated in v1.92.0+) Replaced by ` + "`" + `access_group_name` + "`" + ` after version 1.92.0.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required, ForceNew, Deprecated in v1.92.0+) Replaced by ` + "`" + `access_group_type` + "`" + ` after version 1.92.0.`,
				},
				resource.Attribute{
					Name:        "access_group_name",
					Description: `(Required, ForceNew, Available in v1.92.0+) A Name of one Access Group.`,
				},
				resource.Attribute{
					Name:        "access_group_type",
					Description: `(Required, ForceNew, Available in v1.92.0+) A Type of one Access Group. Valid values: ` + "`" + `Vpc` + "`" + ` and ` + "`" + `Classic` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The Access Group description.`,
				},
				resource.Attribute{
					Name:        "file_system_type",
					Description: `(Optional, ForceNew, Available in v1.92.0+) The type of file system. Valid values: ` + "`" + `standard` + "`" + ` and ` + "`" + `extreme` + "`" + `. Default to ` + "`" + `standard` + "`" + `. Note that the extreme only support Vpc Network. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Access Group. The value as ` + "`" + `<access_group_name>` + "`" + `. ->`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Access Group. The value as ` + "`" + `<access_group_name>` + "`" + `. ->`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_nas_access_rule",
			Category:         "Network Attached Storage (NAS)",
			ShortDescription: `Provides a Alibabacloudstack Nas Access Rule resource.`,
			Description:      ``,
			Keywords: []string{
				"network",
				"attached",
				"storage",
				"nas",
				"access",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "access_group_name",
					Description: `(Required, ForceNew) Permission group name.`,
				},
				resource.Attribute{
					Name:        "source_cidr_ip",
					Description: `(Required) Address or address segment.`,
				},
				resource.Attribute{
					Name:        "rw_access_type",
					Description: `(Optional) Read-write permission type: ` + "`" + `RDWR` + "`" + ` (default), ` + "`" + `RDONLY` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user_access_type",
					Description: `(Optional) User permission type: ` + "`" + `no_squash` + "`" + ` (default), ` + "`" + `root_squash` + "`" + `, ` + "`" + `all_squash` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Priority level. Range: 1-100. Default value: ` + "`" + `1` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `This ID of this resource. The value is formate as ` + "`" + `<access_group_name>:<access rule id>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "access_rule_id",
					Description: `The nas access rule ID. ## Import Nas Access Rule can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_nas_access_rule.foo tf-testAccNasConfigName:1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `This ID of this resource. The value is formate as ` + "`" + `<access_group_name>:<access rule id>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "access_rule_id",
					Description: `The nas access rule ID. ## Import Nas Access Rule can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_nas_access_rule.foo tf-testAccNasConfigName:1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_nas_file_system",
			Category:         "Network Attached Storage (NAS)",
			ShortDescription: `Provides a Alibabacloudstack NAS File System resource.`,
			Description:      ``,
			Keywords: []string{
				"network",
				"attached",
				"storage",
				"nas",
				"file",
				"system",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "file_system_type",
					Description: `(Optional, Available in v1.140.0+) the type of the file system. Valid values: ` + "`" + `standard` + "`" + ` (Default), ` + "`" + `extreme` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "protocol_type",
					Description: `(Required, ForceNew) The protocol type of the file system. Valid values: ` + "`" + `NFS` + "`" + `, ` + "`" + `SMB` + "`" + ` (Available when the ` + "`" + `file_system_type` + "`" + ` is ` + "`" + `standard` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `(Required, ForceNew) The storage type of the file System.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The File System description.`,
				},
				resource.Attribute{
					Name:        "encrypt_type",
					Description: `(Optional, Available in v1.121.2+) Whether the file system is encrypted. Using kms service escrow key to encrypt and store the file system data. When reading and writing encrypted data, there is no need to decrypt.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `(Optional, Available in v1.140.0+ and when the ` + "`" + `file_system_type` + "`" + ` is ` + "`" + `extreme` + "`" + `) The capacity of the file system. The ` + "`" + `capacity` + "`" + ` is required when the ` + "`" + `file_system_type` + "`" + ` is ` + "`" + `extreme` + "`" + `. Unit: gib;`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Optional, Available in v1.140.0+) The available zones information that supports nas.When FileSystemType=standard, this parameter is not required.`,
				},
				resource.Attribute{
					Name:        "kms_key_id",
					Description: `(Optional, Available in v1.140.0+ and when the ` + "`" + `encrypt_type` + "`" + ` is ` + "`" + `2` + "`" + `) The id of the KMS key. The ` + "`" + `kms_key_id` + "`" + ` is required when the ` + "`" + `encrypt_type` + "`" + ` is ` + "`" + `2` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the File System. ## Import Nas File System can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_nas_file_system.foo 1337849c59 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the File System. ## Import Nas File System can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_nas_file_system.foo 1337849c59 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_nas_mount_target",
			Category:         "Network Attached Storage (NAS)",
			ShortDescription: `Provides a Alibabacloudstack NAS MountTarget resource.`,
			Description:      ``,
			Keywords: []string{
				"network",
				"attached",
				"storage",
				"nas",
				"mount",
				"target",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "file_system_id",
					Description: `(Required, ForceNew) The ID of the file system.`,
				},
				resource.Attribute{
					Name:        "access_group_name",
					Description: `(Required) The name of the permission group that applies to the mount target.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Optional, ForceNew) The ID of the VSwitch in the VPC where the mount target resides.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Whether the MountTarget is active. The status of the mount target. Valid values: ` + "`" + `Active` + "`" + ` and ` + "`" + `Inactive` + "`" + `, Default value is ` + "`" + `Active` + "`" + `. Before you mount a file system, make sure that the mount target is in the Active state.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Optional, ForceNew, Available in v1.95.0+.) The ID of security group. ## Attributes Reference The following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_nat_gateway",
			Category:         "VPC",
			ShortDescription: `Provides a resource to create a VPC NAT Gateway.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"nat",
				"gateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, ForceNew) The VPC ID.`,
				},
				resource.Attribute{
					Name:        "specification",
					Description: `(Optional) The specification of the nat gateway. Valid values are ` + "`" + `Small` + "`" + `, ` + "`" + `Middle` + "`" + ` and ` + "`" + `Large` + "`" + `. Default to ` + "`" + `Small` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the nat gateway. The value can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Defaults to null.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the nat gateway, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Defaults to null.`,
				},
				resource.Attribute{
					Name:        "bandwidth_packages",
					Description: `(Optional) A list of bandwidth packages for the nat gateway. Only support nat gateway created before 00:00 on November 4, 2017. ## Block bandwidth packages The bandwidth package mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "ip_count",
					Description: `(Required) The IP number of the current bandwidth package. Its value range from 1 to 50.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Required) The bandwidth value of the current bandwidth package. Its value range from 5 to 5000.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) The AZ for the current bandwidth. If this value is not specified, Terraform will set a random AZ.`,
				},
				resource.Attribute{
					Name:        "public_ip_addresses",
					Description: `(Computer) The public ip for bandwidth package. the public ip count equal ` + "`" + `ip_count` + "`" + `, multi ip would complex with ",", such as "10.0.0.1,10.0.0.2". ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the nat gateway.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the nat gateway.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the nat gateway.`,
				},
				resource.Attribute{
					Name:        "specification",
					Description: `The specification of the nat gateway.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The VPC ID for the nat gateway.`,
				},
				resource.Attribute{
					Name:        "bandwidth_package_ids",
					Description: `A list ID of the bandwidth packages, and split them with commas.`,
				},
				resource.Attribute{
					Name:        "snat_table_ids",
					Description: `The nat gateway will auto create a snap and forward item, the ` + "`" + `snat_table_ids` + "`" + ` is the created one.`,
				},
				resource.Attribute{
					Name:        "forward_table_ids",
					Description: `The nat gateway will auto create a snap and forward item, the ` + "`" + `forward_table_ids` + "`" + ` is the created one.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the nat gateway.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the nat gateway.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the nat gateway.`,
				},
				resource.Attribute{
					Name:        "specification",
					Description: `The specification of the nat gateway.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The VPC ID for the nat gateway.`,
				},
				resource.Attribute{
					Name:        "bandwidth_package_ids",
					Description: `A list ID of the bandwidth packages, and split them with commas.`,
				},
				resource.Attribute{
					Name:        "snat_table_ids",
					Description: `The nat gateway will auto create a snap and forward item, the ` + "`" + `snat_table_ids` + "`" + ` is the created one.`,
				},
				resource.Attribute{
					Name:        "forward_table_ids",
					Description: `The nat gateway will auto create a snap and forward item, the ` + "`" + `forward_table_ids` + "`" + ` is the created one.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_network_acl",
			Category:         "VPC",
			ShortDescription: `Provides a Alibabacloudstack Network Acl resource.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"network",
				"acl",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, ForceNew) The vpc_id of the network acl, the field can't be changed.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Field ` + "`" + `name` + "`" + ` has been deprecated from provider. New field ` + "`" + `network_acl_name` + "`" + ` instead.`,
				},
				resource.Attribute{
					Name:        "network_acl_name",
					Description: `(Optional) The name of the network acl.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the network acl instance.`,
				},
				resource.Attribute{
					Name:        "ingress_acl_entries",
					Description: `(Optional, Computed) List of the ingress entries of the network acl. The order of the ingress entries determines the priority. The details see Block ` + "`" + `ingress_acl_entries` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "egress_acl_entries",
					Description: `(Optional, Computed) List of the egress entries of the network acl. The order of the egress entries determines the priority. The details see Block ` + "`" + `egress_acl_entries` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `(Optional) The associated resources. ### Block ingress_acl_entries`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of ingress entries.`,
				},
				resource.Attribute{
					Name:        "network_acl_entry_name",
					Description: `(Optional) The entry name of ingress entries.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Optional) The policy of ingress entries. Valid values ` + "`" + `accept` + "`" + ` and ` + "`" + `drop` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The port of ingress entries.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) The protocol of ingress entries. Valid values ` + "`" + `icmp` + "`" + `,` + "`" + `gre` + "`" + `,` + "`" + `tcp` + "`" + `,` + "`" + `udp` + "`" + `, and ` + "`" + `all` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "source_cidr_ip",
					Description: `(Optional) The source cidr ip of ingress entries. ### Block egress_acl_entries`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of egress entries.`,
				},
				resource.Attribute{
					Name:        "network_acl_entry_name",
					Description: `(Optional) The entry name of egress entries.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Optional) The policy of egress entries. Valid values ` + "`" + `accept` + "`" + ` and ` + "`" + `drop` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The port of egress entries.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) The protocol of egress entries. Valid values ` + "`" + `icmp` + "`" + `,` + "`" + `gre` + "`" + `,` + "`" + `tcp` + "`" + `,` + "`" + `udp` + "`" + `, and ` + "`" + `all` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "destination_cidr_ip",
					Description: `(Optional) The destination cidr ip of egress entries. ### Block resources`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Optional) The ID of the associated resource.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `(Optional) The type of the associated resource. Valid values ` + "`" + `VSwitch` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the network acl instance id.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the network acl. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 10 mins) Used when creating the Network ACL. (until it reaches the initial ` + "`" + `Available` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 10 mins) Used when updating the Network ACL. (until it reaches the initial ` + "`" + `Available` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 10 mins) Used when terminating the Network ACL. ## Import The network acl can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_network_acl.default nacl-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the network acl instance id.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the network acl. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 10 mins) Used when creating the Network ACL. (until it reaches the initial ` + "`" + `Available` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 10 mins) Used when updating the Network ACL. (until it reaches the initial ` + "`" + `Available` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 10 mins) Used when terminating the Network ACL. ## Import The network acl can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_network_acl.default nacl-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_network_acl_attachment",
			Category:         "VPC",
			ShortDescription: `Provides a Alibabacloudstack Network Acl Attachment resource.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"network",
				"acl",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "network_acl_id",
					Description: `(Required, ForceNew) The id of the network acl, the field can't be changed.`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `(Required) List of the resources associated with the network acl. The details see Block Resources. ### Block Resources The resources mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required) The resource id that the network acl will associate with.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `(Required) The resource id that the network acl will associate with. Only support ` + "`" + `VSwitch` + "`" + ` now. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the network acl attachment. It is formatted as ` + "`" + `<network_acl_id>:<a unique id>` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the network acl attachment. It is formatted as ` + "`" + `<network_acl_id>:<a unique id>` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_network_acl_entries",
			Category:         "VPC",
			ShortDescription: `Provides a Alibabacloudstack Network Acl Entries resource.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"network",
				"acl",
				"entries",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "network_acl_id",
					Description: `(Required, ForceNew) The id of the network acl, the field can't be changed.`,
				},
				resource.Attribute{
					Name:        "ingress",
					Description: `(Optional) List of the ingress entries of the network acl. The order of the ingress entries determines the priority. The details see Block Ingress.`,
				},
				resource.Attribute{
					Name:        "egress",
					Description: `(Optional) List of the egress entries of the network acl. The order of the egress entries determines the priority. The details see Block Egress. ### Ingress Resources The resources mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the ingress entry.`,
				},
				resource.Attribute{
					Name:        "source_cidr_ip",
					Description: `(Optional) The source ip of the ingress entry.`,
				},
				resource.Attribute{
					Name:        "entry_type",
					Description: `(Optional) The entry type of the ingress entry. It must be ` + "`" + `custom` + "`" + ` or ` + "`" + `system` + "`" + `. Default value is ` + "`" + `custom` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the ingress entry.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Optional) The policy of the ingress entry. It must be ` + "`" + `accept` + "`" + ` or ` + "`" + `drop` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The port of the ingress entry.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) The protocol of the ingress entry. ### Egress Resources The resources mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the egress entry.`,
				},
				resource.Attribute{
					Name:        "destination_cidr_ip",
					Description: `(Optional) The destination ip of the egress entry.`,
				},
				resource.Attribute{
					Name:        "entry_type",
					Description: `(Optional) The entry type of the egress entry. It must be ` + "`" + `custom` + "`" + ` or ` + "`" + `system` + "`" + `. Default value is ` + "`" + `custom` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the egress entry.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Optional) The policy of the egress entry. It must be ` + "`" + `accept` + "`" + ` or ` + "`" + `drop` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The port of the egress entry.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) The protocol of the egress entry. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the network acl entries. It is formatted as ` + "`" + `<network_acl_id>:<a unique id>` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the network acl entries. It is formatted as ` + "`" + `<network_acl_id>:<a unique id>` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_network_interface",
			Category:         "ECS",
			ShortDescription: `Provides an ECS Elastic Network Interface resource.`,
			Description:      ``,
			Keywords: []string{
				"ecs",
				"network",
				"interface",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Required, ForceNew) The VSwitch to create the ENI in.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Required) A list of security group ids to associate with.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `(Optional, ForceNew) The primary private IP of the ENI.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the ENI. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-", ".", "_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the ENI. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.`,
				},
				resource.Attribute{
					Name:        "private_ips_count",
					Description: `(Optional) Number of secondary private IPs to assign to the ENI. Don't use both private_ips and private_ips_count in the same ENI resource block.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ENI ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ENI ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_network_interface_attachment",
			Category:         "ECS",
			ShortDescription: `Provides an alibabacloudstack ECS Elastic Network Interface Attachment as a resource to attach ENI to or detach ENI from ECS Instances.`,
			Description:      ``,
			Keywords: []string{
				"ecs",
				"network",
				"interface",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, ForceNew) The instance ID to attach.`,
				},
				resource.Attribute{
					Name:        "network_interface_id",
					Description: `(Required, ForceNew) The ENI ID to attach. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the resource, formatted as ` + "`" + `<network_interface_id>:<instance_id>` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the resource, formatted as ` + "`" + `<network_interface_id>:<instance_id>` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ons_group",
			Category:         "RocketMQ",
			ShortDescription: `Provides a alibabacloudstack ONS Group resource.`,
			Description:      ``,
			Keywords: []string{
				"rocketmq",
				"ons",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) ID of the ONS Instance that owns the groups.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required) Name of the group. Two groups on a single instance cannot have the same name. A ` + "`" + `group_id` + "`" + ` starts with "GID_" or "GID-", and contains letters, numbers, hyphens (-), and underscores (_).`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `(Required) This attribute is a concise description of group. The length cannot exceed 256.`,
				},
				resource.Attribute{
					Name:        "read_enable",
					Description: `(Optional) This attribute is used to set the message reading enabled or disabled. It can only be set after the group is used by the client. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `GroupID and InstanceID of the ONS Group. The value is in format ` + "`" + `GroupID:InstanceID` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `GroupID and InstanceID of the ONS Group. The value is in format ` + "`" + `GroupID:InstanceID` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ons_instance",
			Category:         "RocketMQ",
			ShortDescription: `Provides a alibabacloudstack ONS Instance resource.`,
			Description:      ``,
			Keywords: []string{
				"rocketmq",
				"ons",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required)Two instances on a single account in the same region cannot have the same name. The length must be 3 to 64 characters. Chinese characters, English letters digits and hyphen are allowed.`,
				},
				resource.Attribute{
					Name:        "tps_receive_max",
					Description: `(Required)This attribute is used to set the message receiving transactions per second (TPS) of the topic during a certain period of time.`,
				},
				resource.Attribute{
					Name:        "tps_send_max",
					Description: `(Required)This attribute is used to set the message sending transactions per second (TPS) of the topic during a certain period of time.`,
				},
				resource.Attribute{
					Name:        "topic_capacity",
					Description: `(Required)This attribute is used to set the topic capacity.`,
				},
				resource.Attribute{
					Name:        "independent_naming",
					Description: `(Required)This attribute is used to define an independent name or not. It takes only bool value.`,
				},
				resource.Attribute{
					Name:        "cluster",
					Description: `(Required)This attribute is a used to add cluster name.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `(Optional)This attribute is a concise description of instance. The length cannot exceed 128. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `key` + "`" + ` of the resource supplied above.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `The edition of instance. 1 represents the postPaid edition, and 2 represents the platinum edition.`,
				},
				resource.Attribute{
					Name:        "instance_status",
					Description: `The status of instance. 1 represents the platinum edition instance is in deployment. 2 represents the postpaid edition instance are overdue. 5 represents the postpaid or platinum edition instance is in service. 7 represents the platinum version instance is in upgrade and the service is available.`,
				},
				resource.Attribute{
					Name:        "release_time",
					Description: `Platinum edition instance expiration time.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `key` + "`" + ` of the resource supplied above.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `The edition of instance. 1 represents the postPaid edition, and 2 represents the platinum edition.`,
				},
				resource.Attribute{
					Name:        "instance_status",
					Description: `The status of instance. 1 represents the platinum edition instance is in deployment. 2 represents the postpaid edition instance are overdue. 5 represents the postpaid or platinum edition instance is in service. 7 represents the platinum version instance is in upgrade and the service is available.`,
				},
				resource.Attribute{
					Name:        "release_time",
					Description: `Platinum edition instance expiration time.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ons_topic",
			Category:         "RocketMQ",
			ShortDescription: `Provides a alibabacloudstack ONS Topic resource.`,
			Description:      ``,
			Keywords: []string{
				"rocketmq",
				"ons",
				"topic",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) ID of the ONS Instance that owns the topics.`,
				},
				resource.Attribute{
					Name:        "topic",
					Description: `(Required) Name of the topic. Two topics on a single instance cannot have the same name and the name cannot start with 'GID' or 'CID'. The length cannot exceed 64 characters.`,
				},
				resource.Attribute{
					Name:        "message_type",
					Description: `(Required) The type of the message.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `(Required) This attribute is a concise description of topic. The length cannot exceed 128.`,
				},
				resource.Attribute{
					Name:        "perm",
					Description: `(Optional) This attribute is used to set the read-write mode for the topic. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Topic and InstanceID of the ONS Topic. The value is in format ` + "`" + `Topic:InstanceID` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Topic and InstanceID of the ONS Topic. The value is in format ` + "`" + `Topic:InstanceID` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_oos_execution",
			Category:         "Operation Orchestration Service (OOS)",
			ShortDescription: `Provides a OOS Execution resource.`,
			Description:      ``,
			Keywords: []string{
				"operation",
				"orchestration",
				"service",
				"oos",
				"execution",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, ForceNew) The description of OOS Execution.`,
				},
				resource.Attribute{
					Name:        "loop_mode",
					Description: `(Optional, ForceNew) The loop mode of OOS Execution.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional, ForceNew) The mode of OOS Execution. Valid: ` + "`" + `Automatic` + "`" + `, ` + "`" + `Debug` + "`" + `. Default to ` + "`" + `Automatic` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `(Optional, ForceNew) The parameters required by the template. Default to ` + "`" + `{}` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "parent_execution_id",
					Description: `(Optional, ForceNew) The id of parent execution.`,
				},
				resource.Attribute{
					Name:        "safety_check",
					Description: `(Optional, ForceNew) The mode of safety check.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required, ForceNew) The name of execution template.`,
				},
				resource.Attribute{
					Name:        "template_version",
					Description: `(Optional, ForceNew) The version of execution template.`,
				},
				resource.Attribute{
					Name:        "template_content",
					Description: `(Optional, ForceNew, Available in v1.114.0+) The content of template. When the user selects an existing template to create and execute a task, it is not necessary to pass in this field. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of OOS Execution.`,
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
					Name:        "end_date",
					Description: `The time when the execution was ended.`,
				},
				resource.Attribute{
					Name:        "executed_by",
					Description: `The user who execute the template.`,
				},
				resource.Attribute{
					Name:        "is_parent",
					Description: `Whether to include subtasks.`,
				},
				resource.Attribute{
					Name:        "outputs",
					Description: `The outputs of OOS Execution.`,
				},
				resource.Attribute{
					Name:        "ram_role",
					Description: `The role that executes the current template.`,
				},
				resource.Attribute{
					Name:        "start_date",
					Description: `The time when the execution was started.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of OOS Execution.`,
				},
				resource.Attribute{
					Name:        "status_message",
					Description: `The message of status.`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `The id of template.`,
				},
				resource.Attribute{
					Name:        "update_date",
					Description: `The time when the execution was updated. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 11 mins) Used when creating the alibabacloudstack_oos_execution (until it reaches the initial ` + "`" + `Running` + "`" + ` status). ## Import OOS Execution can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_oos_execution.example exec-ef6xxxx ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of OOS Execution.`,
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
					Name:        "end_date",
					Description: `The time when the execution was ended.`,
				},
				resource.Attribute{
					Name:        "executed_by",
					Description: `The user who execute the template.`,
				},
				resource.Attribute{
					Name:        "is_parent",
					Description: `Whether to include subtasks.`,
				},
				resource.Attribute{
					Name:        "outputs",
					Description: `The outputs of OOS Execution.`,
				},
				resource.Attribute{
					Name:        "ram_role",
					Description: `The role that executes the current template.`,
				},
				resource.Attribute{
					Name:        "start_date",
					Description: `The time when the execution was started.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of OOS Execution.`,
				},
				resource.Attribute{
					Name:        "status_message",
					Description: `The message of status.`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `The id of template.`,
				},
				resource.Attribute{
					Name:        "update_date",
					Description: `The time when the execution was updated. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 11 mins) Used when creating the alibabacloudstack_oos_execution (until it reaches the initial ` + "`" + `Running` + "`" + ` status). ## Import OOS Execution can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_oos_execution.example exec-ef6xxxx ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_oos_template",
			Category:         "Operation Orchestration Service (OOS)",
			ShortDescription: `Provides a OOS Template resource.`,
			Description:      ``,
			Keywords: []string{
				"operation",
				"orchestration",
				"service",
				"oos",
				"template",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "content",
					Description: `(Required) The content of the template. The template must be in the JSON or YAML format. Maximum size: 64 KB.`,
				},
				resource.Attribute{
					Name:        "auto_delete_executions",
					Description: `(Optional) When deleting a template, whether to delete its related executions. Default to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required, ForceNew) The name of the template. The template name can be up to 200 characters in length. The name can contain letters, digits, hyphens (-), and underscores (_). It cannot start with ` + "`" + `ALIYUN` + "`" + `, ` + "`" + `ACS` + "`" + `, ` + "`" + `ALIBABA` + "`" + `, or ` + "`" + `ALICLOUD` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "version_name",
					Description: `(Optional) The name of template version.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the resource. It same with ` + "`" + `template_name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `The creator of the template.`,
				},
				resource.Attribute{
					Name:        "created_date",
					Description: `The time when the template is created.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the template.`,
				},
				resource.Attribute{
					Name:        "has_trigger",
					Description: `Is it triggered successfully.`,
				},
				resource.Attribute{
					Name:        "share_type",
					Description: `The sharing type of the template. The sharing type of templates created by users are set to Private. The sharing type of common templates provided by OOS are set to Public.`,
				},
				resource.Attribute{
					Name:        "template_format",
					Description: `The format of the template. The format can be JSON or YAML. The system automatically identifies the format.`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `The id of OOS Template.`,
				},
				resource.Attribute{
					Name:        "template_type",
					Description: `The type of OOS Template. ` + "`" + `Automation` + "`" + ` means the implementation of Alibaba Cloud API template, ` + "`" + `Package` + "`" + ` means represents a template for installing software.`,
				},
				resource.Attribute{
					Name:        "template_version",
					Description: `The version of OOS Template.`,
				},
				resource.Attribute{
					Name:        "updated_by",
					Description: `The user who updated the template.`,
				},
				resource.Attribute{
					Name:        "updated_date",
					Description: `The time when the template was updated. ## Import OOS Template can be imported using the id or template_name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_oos_template.example template_name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the resource. It same with ` + "`" + `template_name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `The creator of the template.`,
				},
				resource.Attribute{
					Name:        "created_date",
					Description: `The time when the template is created.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the template.`,
				},
				resource.Attribute{
					Name:        "has_trigger",
					Description: `Is it triggered successfully.`,
				},
				resource.Attribute{
					Name:        "share_type",
					Description: `The sharing type of the template. The sharing type of templates created by users are set to Private. The sharing type of common templates provided by OOS are set to Public.`,
				},
				resource.Attribute{
					Name:        "template_format",
					Description: `The format of the template. The format can be JSON or YAML. The system automatically identifies the format.`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `The id of OOS Template.`,
				},
				resource.Attribute{
					Name:        "template_type",
					Description: `The type of OOS Template. ` + "`" + `Automation` + "`" + ` means the implementation of Alibaba Cloud API template, ` + "`" + `Package` + "`" + ` means represents a template for installing software.`,
				},
				resource.Attribute{
					Name:        "template_version",
					Description: `The version of OOS Template.`,
				},
				resource.Attribute{
					Name:        "updated_by",
					Description: `The user who updated the template.`,
				},
				resource.Attribute{
					Name:        "updated_date",
					Description: `The time when the template was updated. ## Import OOS Template can be imported using the id or template_name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_oos_template.example template_name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_oss_bucket",
			Category:         "OSS",
			ShortDescription: `Provides a resource to create an oss bucket.`,
			Description:      ``,
			Keywords: []string{
				"oss",
				"bucket",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Optional, ForceNew) The name of the bucket. If omitted, Terraform will assign a random and unique name.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `(Optional) It Can be "private", "public-read" and "public-read-write". Defaults to "private".`,
				},
				resource.Attribute{
					Name:        "logging",
					Description: `(Optional) The logging object supports the following: - ` + "`" + `target_bucket` + "`" + ` - (Required) The name of the bucket that will receive the log objects. - ` + "`" + `target_prefix` + "`" + ` - (Optional) To specify a key prefix for log objects. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The name of the bucket.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `The acl of the bucket.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `The creation date of the bucket.`,
				},
				resource.Attribute{
					Name:        "extranet_endpoint",
					Description: `The extranet access endpoint of the bucket.`,
				},
				resource.Attribute{
					Name:        "intranet_endpoint",
					Description: `The intranet access endpoint of the bucket.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location of the bucket.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `The bucket owner.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The name of the bucket.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `The acl of the bucket.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `The creation date of the bucket.`,
				},
				resource.Attribute{
					Name:        "extranet_endpoint",
					Description: `The extranet access endpoint of the bucket.`,
				},
				resource.Attribute{
					Name:        "intranet_endpoint",
					Description: `The intranet access endpoint of the bucket.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location of the bucket.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `The bucket owner.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_oss_bucket_object",
			Category:         "OSS",
			ShortDescription: `Provides a resource to create a oss bucket object.`,
			Description:      ``,
			Keywords: []string{
				"oss",
				"bucket",
				"object",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) The name of the bucket to put the file in.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The name of the object once it is in the bucket.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) The path to the source file being uploaded to the bucket.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Optional unless ` + "`" + `source` + "`" + ` given) The literal content being uploaded to the bucket.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `(Optional) The [canned ACL](https://www.alibabacloud.com/help/doc-detail/52284.htm) to apply. Defaults to "private".`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `(Optional) A standard MIME type describing the format of the object data, e.g. application/octet-stream. All Valid MIME Types are valid for this input.`,
				},
				resource.Attribute{
					Name:        "cache_control",
					Description: `(Optional) Specifies caching behavior along the request/reply chain. Read [RFC2616 Cache-Control](https://www.ietf.org/rfc/rfc2616.txt) for further details.`,
				},
				resource.Attribute{
					Name:        "content_disposition",
					Description: `(Optional) Specifies presentational information for the object. Read [RFC2616 Content-Disposition](https://www.ietf.org/rfc/rfc2616.txt) for further details.`,
				},
				resource.Attribute{
					Name:        "content_encoding",
					Description: `(Optional) Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read [RFC2616 Content-Encoding](https://www.ietf.org/rfc/rfc2616.txt) for further details.`,
				},
				resource.Attribute{
					Name:        "content_md5",
					Description: `(Optional) The MD5 value of the content. Read [MD5](https://www.alibabacloud.com/help/doc-detail/31978.htm) for computing method.`,
				},
				resource.Attribute{
					Name:        "expires",
					Description: `(Optional) Specifies expire date for the the request/response. Read [RFC2616 Expires](https://www.ietf.org/rfc/rfc2616.txt) for further details.`,
				},
				resource.Attribute{
					Name:        "server_side_encryption",
					Description: `(Optional) Specifies server-side encryption of the object in OSS. Valid values are ` + "`" + `AES256` + "`" + `, ` + "`" + `KMS` + "`" + `. Default value is ` + "`" + `AES256` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "kms_key_id",
					Description: `(Optional, Available in 1.62.1+) Specifies the primary key managed by KMS. This parameter is valid when the value of ` + "`" + `server_side_encryption` + "`" + ` is set to KMS. Either ` + "`" + `source` + "`" + ` or ` + "`" + `content` + "`" + ` must be provided to specify the bucket content. These two arguments are mutually-exclusive. ## Attributes Reference The following attributes are exported`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `the ` + "`" + `key` + "`" + ` of the resource supplied above.`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `A unique version ID value for the object, if bucket versioning is enabled.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `the ` + "`" + `key` + "`" + ` of the resource supplied above.`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `A unique version ID value for the object, if bucket versioning is enabled.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ots_instance",
			Category:         "Table Store (OTS)",
			ShortDescription: `Provides an OTS (Open Table Service) instance resource.`,
			Description:      ``,
			Keywords: []string{
				"table",
				"store",
				"ots",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) The name of the instance.`,
				},
				resource.Attribute{
					Name:        "accessed_by",
					Description: `The network limitation of accessing instance. Valid values:`,
				},
				resource.Attribute{
					Name:        "Any",
					Description: `Allow all network to access the instance.`,
				},
				resource.Attribute{
					Name:        "Vpc",
					Description: `Only can the attached VPC allow to access the instance.`,
				},
				resource.Attribute{
					Name:        "ConsoleOrVpc",
					Description: `Allow web console or the attached VPC to access the instance. Default to "Any".`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(ForceNew) The type of instance. Valid values are "Capacity" and "HighPerformance". Default to "HighPerformance".`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, ForceNew) The description of the instance. Currently, it does not support modifying.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the instance. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID. The value is same as the "name".`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The instance name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The instance description.`,
				},
				resource.Attribute{
					Name:        "accessed_by",
					Description: `TThe network limitation of accessing instance.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `The instance type.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The instance tags. ## Import OTS instance can be imported using instance id or name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_ots_instance.foo "my-ots-instance" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID. The value is same as the "name".`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The instance name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The instance description.`,
				},
				resource.Attribute{
					Name:        "accessed_by",
					Description: `TThe network limitation of accessing instance.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `The instance type.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The instance tags. ## Import OTS instance can be imported using instance id or name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_ots_instance.foo "my-ots-instance" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ots_instance_attachment",
			Category:         "Table Store (OTS)",
			ShortDescription: `Provides an OTS (Open Table Service) resource to attach VPC to instance.`,
			Description:      ``,
			Keywords: []string{
				"table",
				"store",
				"ots",
				"instance",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Required, ForceNew) The name of the OTS instance.`,
				},
				resource.Attribute{
					Name:        "vpc_name",
					Description: `(Required, ForceNew) The name of attaching VPC to instance.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Required, ForceNew) The ID of attaching VSwitch to instance. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID. The value is same as "instance_name".`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `The instance name.`,
				},
				resource.Attribute{
					Name:        "vpc_name",
					Description: `The name of attaching VPC to instance.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `The ID of attaching VSwitch to instance.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of attaching VPC to instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID. The value is same as "instance_name".`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `The instance name.`,
				},
				resource.Attribute{
					Name:        "vpc_name",
					Description: `The name of attaching VPC to instance.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `The ID of attaching VSwitch to instance.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of attaching VPC to instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ots_table",
			Category:         "Table Store (OTS)",
			ShortDescription: `Provides an OTS (Open Table Service) table resource.`,
			Description:      ``,
			Keywords: []string{
				"table",
				"store",
				"ots",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Required, ForceNew) The name of the OTS instance in which table will located.`,
				},
				resource.Attribute{
					Name:        "table_name",
					Description: `(Required, ForceNew) The table name of the OTS instance. If changed, a new table would be created.`,
				},
				resource.Attribute{
					Name:        "primary_key",
					Description: `(Required, ForceNew) The property of ` + "`" + `TableMeta` + "`" + ` which indicates the structure information of a table. It describes the attribute value of primary key. The number of ` + "`" + `primary_key` + "`" + ` should not be less than one and not be more than four.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Name for primary key.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required, ForceNew) Type for primary key. Only ` + "`" + `Integer` + "`" + `, ` + "`" + `String` + "`" + ` or ` + "`" + `Binary` + "`" + ` is allowed.`,
				},
				resource.Attribute{
					Name:        "time_to_live",
					Description: `(Required) The retention time of data stored in this table (unit: second). The value maximum is 2147483647 and -1 means never expired.`,
				},
				resource.Attribute{
					Name:        "max_version",
					Description: `(Required) The maximum number of versions stored in this table. The valid value is 1-2147483647.`,
				},
				resource.Attribute{
					Name:        "deviation_cell_version_in_sec",
					Description: `(Optional, Available in 1.42.0+) The max version offset of the table. The valid value is 1-9223372036854775807. Defaults to 86400. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID. The value is ` + "`" + `<instance_name>:<table_name>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `The OTS instance name.`,
				},
				resource.Attribute{
					Name:        "table_name",
					Description: `The table name of the OTS which could not be changed.`,
				},
				resource.Attribute{
					Name:        "primary_key",
					Description: `The property of ` + "`" + `TableMeta` + "`" + ` which indicates the structure information of a table.`,
				},
				resource.Attribute{
					Name:        "time_to_live",
					Description: `The retention time of data stored in this table.`,
				},
				resource.Attribute{
					Name:        "max_version",
					Description: `The maximum number of versions stored in this table.`,
				},
				resource.Attribute{
					Name:        "deviation_cell_version_in_sec",
					Description: `The max version offset of the table. ## Import OTS table can be imported using id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_ots_table.table "my-ots:ots_table" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID. The value is ` + "`" + `<instance_name>:<table_name>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `The OTS instance name.`,
				},
				resource.Attribute{
					Name:        "table_name",
					Description: `The table name of the OTS which could not be changed.`,
				},
				resource.Attribute{
					Name:        "primary_key",
					Description: `The property of ` + "`" + `TableMeta` + "`" + ` which indicates the structure information of a table.`,
				},
				resource.Attribute{
					Name:        "time_to_live",
					Description: `The retention time of data stored in this table.`,
				},
				resource.Attribute{
					Name:        "max_version",
					Description: `The maximum number of versions stored in this table.`,
				},
				resource.Attribute{
					Name:        "deviation_cell_version_in_sec",
					Description: `The max version offset of the table. ## Import OTS table can be imported using id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_ots_table.table "my-ots:ots_table" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_quick_bi_user",
			Category:         "Quick BI",
			ShortDescription: `Provides a Alibabacloudstack Quick BI User resource.`,
			Description:      ``,
			Keywords: []string{
				"quick",
				"bi",
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_id",
					Description: `(Optional, ForceNew) Alibaba Cloud account ID.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `(Required) An Alibaba Cloud account, Alibaba Cloud name.`,
				},
				resource.Attribute{
					Name:        "admin_user",
					Description: `(Required) Whether it is the administrator. Valid values: ` + "`" + `true` + "`" + ` and ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "auth_admin_user",
					Description: `(Required) Whether this is a permissions administrator. Valid values: ` + "`" + `false` + "`" + `, ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "nick_name",
					Description: `(Required, ForceNew) The nickname of the user.`,
				},
				resource.Attribute{
					Name:        "user_type",
					Description: `(Required) The members of the organization of the type of role separately. Valid values: ` + "`" + `Analyst` + "`" + `, ` + "`" + `Developer` + "`" + ` and ` + "`" + `Visitor` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in terraform of User. ## Import Quick BI User can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_quick_bi_user.example <id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in terraform of User. ## Import Quick BI User can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_quick_bi_user.example <id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_quick_bi_user_group",
			Category:         "Quick BI",
			ShortDescription: `Provides a Alibabacloudstack Quick BI UserGroup resource.`,
			Description:      ``,
			Keywords: []string{
				"quick",
				"bi",
				"user",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "user_group_name",
					Description: `(Required) User group name.`,
				},
				resource.Attribute{
					Name:        "user_group_description",
					Description: `(Required) User group description.`,
				},
				resource.Attribute{
					Name:        "parent_user_group_id",
					Description: `(Required) Parent user group ID. You can add a new user group to this grouping.When you enter -1, the newly created user group will be added to the root directory. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "user_group_id",
					Description: `The resource ID in terraform of UserGroup. ## Import Quick BI UserGroup can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_quick_bi_user_group.example <id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "user_group_id",
					Description: `The resource ID in terraform of UserGroup. ## Import Quick BI UserGroup can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_quick_bi_user_group.example <id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_quick_bi_workspace",
			Category:         "Quick BI",
			ShortDescription: `Provides a Alibabacloudstack Quick BI Workspace resource.`,
			Description:      ``,
			Keywords: []string{
				"quick",
				"bi",
				"workspace",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "workspace_name",
					Description: `(Required) Workspace name.`,
				},
				resource.Attribute{
					Name:        "workspace_desc",
					Description: `(Optional) Workspace description.`,
				},
				resource.Attribute{
					Name:        "use_comment",
					Description: `(Optional) Do you want to use table comments when creating data sets (corresponding to preferences). Valid values: ` + "`" + `true` + "`" + ` and ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "allow_share",
					Description: `(Optional) Whether the report is allowed to be shared (corresponding function permission-works can be authorized). Valid values: ` + "`" + `false` + "`" + `, ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "allow_publish",
					Description: `(Optional) Whether the report is allowed to be made public (corresponding function permission-works can be made public).Valid values: ` + "`" + `false` + "`" + `, ` + "`" + `true` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "workspace_id",
					Description: `The resource ID in terraform of Workspace. ## Import Quick BI Workspace can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_quick_bi_workspace.example <id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "workspace_id",
					Description: `The resource ID in terraform of Workspace. ## Import Quick BI Workspace can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_quick_bi_workspace.example <id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ram_role_attachment",
			Category:         "RAM",
			ShortDescription: `Provides a RAM role attachment resource to bind role for several ECS instances.`,
			Description:      ``,
			Keywords: []string{
				"ram",
				"role",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "role_name",
					Description: `(Required, ForceNew) The name of role used to bind. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-", "_", and must not begin with a hyphen.`,
				},
				resource.Attribute{
					Name:        "instance_ids",
					Description: `(Required, ForceNew) The list of ECS instance's IDs. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "role_name",
					Description: `The name of the role.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "role_name",
					Description: `The name of the role.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_reserved_instance",
			Category:         "ECS",
			ShortDescription: `Provides an ECS Reserved Instance resource.`,
			Description:      ``,
			Keywords: []string{
				"ecs",
				"reserved",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "offering_type",
					Description: `(Required, ForceNew) Payment type of the RI. Optional values: ` + "`" + `No Upfront` + "`" + `: No upfront payment is required., ` + "`" + `Partial Upfront` + "`" + `: A portion of upfront payment is required.` + "`" + `All Upfront` + "`" + `: Full upfront payment is required.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Optional, ForceNew) ID of the zone to which the RI belongs. When Scope is set to Zone, this parameter is required. For information about the zone list, see [DescribeZones](https://www.alibabacloud.com/help/doc-detail/25610.html).`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional, ForceNew) Scope of the RI. Optional values: ` + "`" + `Region` + "`" + `: region-level, ` + "`" + `Zone` + "`" + `: zone-level. Default is ` + "`" + `Region` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Optional, ForceNew) Instance type of the RI. For more information, see [Instance type families](https://www.alibabacloud.com/help/doc-detail/25378.html).`,
				},
				resource.Attribute{
					Name:        "instance_amount",
					Description: `(Optional, ForceNew) Number of instances allocated to an RI (An RI is a coupon that includes one or more allocated instances.).`,
				},
				resource.Attribute{
					Name:        "Period",
					Description: `(Optional, ForceNew) Term of the RI. Unit: years. Optional values: 1 and 3.`,
				},
				resource.Attribute{
					Name:        "period_unit",
					Description: `(Optional, ForceNew) Term unit. Optional value: Year.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional, ForceNew) Resource group ID.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the RI. 2 to 256 English or Chinese characters. It cannot start with http:// or https://.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the RI. The name must be a string of 2 to 128 characters in length and can contain letters, numbers, colons (:), underscores (_), and hyphens. It must start with a letter. It cannot start with http:// or https://.`,
				},
				resource.Attribute{
					Name:        "platform",
					Description: `(Optional, ForceNew) The operating system type of the image used by the instance. Optional values: ` + "`" + `Windows` + "`" + `, ` + "`" + `Linux` + "`" + `. Default is ` + "`" + `Linux` + "`" + `. ### Removing alibabacloudstack_reserved_instance from your configuration The alibabacloudstack_reserved_instance resource allows you to manage your ReservedInstance, but Terraform cannot destroy it. Removing this resource from your configuration will remove it from your statefile and management, but will not destroy the ReservedInstance. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the ReservedInstance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the ReservedInstance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ros_stack",
			Category:         "ROS",
			ShortDescription: `Provides a Alibabacloudstack ROS Stack resource.`,
			Description:      ``,
			Keywords: []string{
				"ros",
				"stack",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "create_option",
					Description: `(Optional, ForceNew) Specifies whether to delete the stack after it is created.`,
				},
				resource.Attribute{
					Name:        "deletion_protection",
					Description: `(Optional, ForceNew) Specifies whether to enable deletion protection on the stack. Valid values: ` + "`" + `Disabled` + "`" + `, ` + "`" + `Enabled` + "`" + `. Default to: ` + "`" + `Disabled` + "`" + ``,
				},
				resource.Attribute{
					Name:        "disable_rollback",
					Description: `(Optional) Specifies whether to disable rollback on stack creation failure. Default to: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "notification_urls",
					Description: `(Optional, ForceNew) The callback URL for receiving stack event N. Only HTTP POST is supported. Maximum value of N: 5.`,
				},
				resource.Attribute{
					Name:        "ram_role_name",
					Description: `(Optional) The name of the RAM role. ROS assumes the specified RAM role to create the stack and call API operations by using the credentials of the role.`,
				},
				resource.Attribute{
					Name:        "replacement_option",
					Description: `(Optional) Specifies whether to enable replacement update after a resource attribute that does not support modification update is changed. Modification update keeps the physical ID of the resource unchanged. However, the resource is deleted and then recreated, and its physical ID is changed if replacement update is enabled.`,
				},
				resource.Attribute{
					Name:        "retain_all_resources",
					Description: `(Optional) The retain all resources.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `(Optional) The parameters. If the parameter name and value are not specified, ROS will use the default value specified in the template.`,
				},
				resource.Attribute{
					Name:        "retain_resources",
					Description: `(Optional) Specifies whether to retain the resources in the stack.`,
				},
				resource.Attribute{
					Name:        "stack_name",
					Description: `(Required, ForceNew) The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.`,
				},
				resource.Attribute{
					Name:        "stack_policy_body",
					Description: `(Optional) The structure that contains the stack policy body. The stack policy body must be 1 to 16,384 bytes in length.`,
				},
				resource.Attribute{
					Name:        "stack_policy_during_update_body",
					Description: `(Optional) The structure that contains the body of the temporary overriding stack policy. The stack policy body must be 1 to 16,384 bytes in length.`,
				},
				resource.Attribute{
					Name:        "stack_policy_during_update_url",
					Description: `(Optional) The URL of the file that contains the temporary overriding stack policy. The URL must point to a policy located in an HTTP or HTTPS web server or an Alibaba Cloud OSS bucket. Examples: oss://ros/stack-policy/demo and oss://ros/stack-policy/demo?RegionId=cn-hangzhou. The policy can be up to 16,384 bytes in length and the URL can be up to 1,350 bytes in length. If the region of the OSS bucket is not specified, the RegionId value is used by default.`,
				},
				resource.Attribute{
					Name:        "stack_policy_url",
					Description: `(Optional) The URL of the file that contains the stack policy. The URL must point to a policy located in an HTTP or HTTPS web server or an Alibaba Cloud OSS bucket. Examples: oss://ros/stack-policy/demo and oss://ros/stack-policy/demo?RegionId=cn-hangzhou. The policy can be up to 16,384 bytes in length and the URL can be up to 1,350 bytes in length. If the region of the OSS bucket is not specified, the RegionId value is used by default.`,
				},
				resource.Attribute{
					Name:        "template_body",
					Description: `(Optional) The structure that contains the template body. The template body must be 1 to 524,288 bytes in length. If the length of the template body is longer than required, we recommend that you add parameters to the HTTP POST request body to avoid request failures due to excessive length of URLs.`,
				},
				resource.Attribute{
					Name:        "template_url",
					Description: `(Optional) The URL of the file that contains the template body. The URL must point to a template located in an HTTP or HTTPS web server or an Alibaba Cloud OSS bucket. Examples: oss://ros/template/demo and oss://ros/template/demo?RegionId=cn-hangzhou. The template must be 1 to 524,288 bytes in length. If the region of the OSS bucket is not specified, the RegionId value is used by default.`,
				},
				resource.Attribute{
					Name:        "template_version",
					Description: `(Optional) The version of the template.`,
				},
				resource.Attribute{
					Name:        "timeout_in_minutes",
					Description: `(Optional) The timeout period that is specified for the stack creation request. Default to: ` + "`" + `60` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "use_previous_parameters",
					Description: `(Optional) Specifies whether to use the values that were passed last time for the parameters that you do not specify in the current request.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. #### Block parameters The parameters supports the following:`,
				},
				resource.Attribute{
					Name:        "parameter_key",
					Description: `(Required) The parameter key.`,
				},
				resource.Attribute{
					Name:        "parameter_value",
					Description: `(Required) The parameter value. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in terraform of Stack. Value as ` + "`" + `stack_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of Stack. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 11 mins) Used when create the Stack.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 6 mins) Used when delete the Stack.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 11 mins) Used when update the Stack. ## Import ROS Stack can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ros_stack.example <stack_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in terraform of Stack. Value as ` + "`" + `stack_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of Stack. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 11 mins) Used when create the Stack.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 6 mins) Used when delete the Stack.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 11 mins) Used when update the Stack. ## Import ROS Stack can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ros_stack.example <stack_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ros_template",
			Category:         "ROS",
			ShortDescription: `Provides a Alibabacloudstack ROS Template resource.`,
			Description:      ``,
			Keywords: []string{
				"ros",
				"template",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the template. The description can be up to 256 characters in length.`,
				},
				resource.Attribute{
					Name:        "template_body",
					Description: `(Optional) The structure that contains the template body. The template body must be 1 to 524,288 bytes in length. If the length of the template body is longer than required, we recommend that you add parameters to the HTTP POST request body to avoid request failures due to excessive length of URLs. You must specify one of the TemplateBody and TemplateURL parameters, but you cannot specify both of them.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) The name of the template. The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.`,
				},
				resource.Attribute{
					Name:        "template_url",
					Description: `(Optional) The template url.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in terraform of Template. Value as ` + "`" + `template_id` + "`" + `. ## Import ROS Template can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ros_template.example <template_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in terraform of Template. Value as ` + "`" + `template_id` + "`" + `. ## Import ROS Template can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ros_template.example <template_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_route_entry",
			Category:         "VPC",
			ShortDescription: `Provides a Alibabacloudstack Route Entry resource.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"route",
				"entry",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "route_table_id",
					Description: `(Required, ForceNew) The ID of the route table.`,
				},
				resource.Attribute{
					Name:        "destination_cidrblock",
					Description: `(ForceNew) The RouteEntry's target network segment.`,
				},
				resource.Attribute{
					Name:        "nexthop_type",
					Description: `(ForceNew) The next hop type. Available values: - ` + "`" + `Instance` + "`" + ` (Default): Route the traffic destined for the destination CIDR block to an ECS instance in the VPC. - ` + "`" + `RouterInterface` + "`" + `: Route the traffic destined for the destination CIDR block to a router interface. - ` + "`" + `VpnGateway` + "`" + `: Route the traffic destined for the destination CIDR block to a VPN Gateway. - ` + "`" + `HaVip` + "`" + `: Route the traffic destined for the destination CIDR block to an HAVIP. - ` + "`" + `NetworkInterface` + "`" + `: Route the traffic destined for the destination CIDR block to an NetworkInterface. - ` + "`" + `NatGateway` + "`" + `: Route the traffic destined for the destination CIDR block to an Nat Gateway.`,
				},
				resource.Attribute{
					Name:        "nexthop_id",
					Description: `(ForceNew) The route entry's next hop. ECS instance ID or VPC router interface ID. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The route entry id,it formats of ` + "`" + `<route_table_id:router_id:destination_cidrblock:nexthop_type:nexthop_id>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The ID of the route table.`,
				},
				resource.Attribute{
					Name:        "destination_cidrblock",
					Description: `The RouteEntry's target network segment.`,
				},
				resource.Attribute{
					Name:        "nexthop_type",
					Description: `The next hop type.`,
				},
				resource.Attribute{
					Name:        "nexthop_id",
					Description: `The route entry's next hop.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The route entry id,it formats of ` + "`" + `<route_table_id:router_id:destination_cidrblock:nexthop_type:nexthop_id>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The ID of the route table.`,
				},
				resource.Attribute{
					Name:        "destination_cidrblock",
					Description: `The RouteEntry's target network segment.`,
				},
				resource.Attribute{
					Name:        "nexthop_type",
					Description: `The next hop type.`,
				},
				resource.Attribute{
					Name:        "nexthop_id",
					Description: `The route entry's next hop.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_route_table",
			Category:         "VPC",
			ShortDescription: `Provides a Alibabacloudstack Route Table resource.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"route",
				"table",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, ForceNew) The vpc_id of the route table, the field can't be changed.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the route table.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the route table instance.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the route table instance id.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the route table instance id.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_route_table_attachment",
			Category:         "VPC",
			ShortDescription: `Provides an alibabacloudstack Route Table Attachment resource.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"route",
				"table",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Required, ForceNew) The vswitch_id of the route table attachment, the field can't be changed.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `(Required, ForceNew) The route_table_id of the route table attachment, the field can't be changed. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the route table attachment id and formates as ` + "`" + `<route_table_id>:<vswitch_id>` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the route table attachment id and formates as ` + "`" + `<route_table_id>:<vswitch_id>` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_router_interface",
			Category:         "VPC",
			ShortDescription: `Provides a VPC router interface resource to connect two VPCs.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"router",
				"interface",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "opposite_region",
					Description: `(Required, ForceNew) The Region of peer side.`,
				},
				resource.Attribute{
					Name:        "router_type",
					Description: `(Required, ForceNew) Router Type. Optional value: VRouter, VBR. Accepting side router interface type only be VRouter.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `(Required, ForceNew) The Router ID.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required, ForceNew) The role the router interface plays. Optional value: ` + "`" + `InitiatingSide` + "`" + `, ` + "`" + `AcceptingSide` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "specification",
					Description: `(Optional) Specification of router interfaces. It is valid when ` + "`" + `role` + "`" + ` is ` + "`" + `InitiatingSide` + "`" + `. Accepting side's role is default to set as 'Negative'.`,
				},
				resource.Attribute{
					Name:        "opposite_access_point_id",
					Description: `(Optional) The ID of the access point of the peer router interface.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the router interface. Length must be 2-80 characters long. Only Chinese characters, English letters, numbers, period (.), underline (_), or dash (-) are permitted. If it is not specified, the default value is interface ID. The name cannot start with http:// and https://.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the router interface. It can be 2-256 characters long or left blank. It cannot start with http:// and https://.`,
				},
				resource.Attribute{
					Name:        "health_check_source_ip",
					Description: `(Optional) Used as the Packet Source IP of health check for disaster recovery or ECMP. It is only valid when ` + "`" + `router_type` + "`" + ` is ` + "`" + `VBR` + "`" + `. The IP must be an unused IP in the local VPC. It and ` + "`" + `health_check_target_ip` + "`" + ` must be specified at the same time.`,
				},
				resource.Attribute{
					Name:        "health_check_target_ip",
					Description: `(Optional) Used as the Packet Target IP of health check for disaster recovery or ECMP. It is only valid when ` + "`" + `router_type` + "`" + ` is ` + "`" + `VBR` + "`" + `. The IP must be an unused IP in the local VPC. It and ` + "`" + `health_check_source_ip` + "`" + ` must be specified at the same time. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Router interface ID.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `Router ID.`,
				},
				resource.Attribute{
					Name:        "router_type",
					Description: `Router type.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `Router interface role.`,
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
					Name:        "specification",
					Description: `Router nterface specification.`,
				},
				resource.Attribute{
					Name:        "access_point_id",
					Description: `Access point of the router interface.`,
				},
				resource.Attribute{
					Name:        "opposite_access_point_id",
					Description: `ID of the access point of the peer`,
				},
				resource.Attribute{
					Name:        "opposite_router_type",
					Description: `Peer router type.`,
				},
				resource.Attribute{
					Name:        "opposite_router_id",
					Description: `Peer router ID.`,
				},
				resource.Attribute{
					Name:        "opposite_interface_id",
					Description: `Peer router interface ID.`,
				},
				resource.Attribute{
					Name:        "opposite_interface_owner_id",
					Description: `Peer account ID.`,
				},
				resource.Attribute{
					Name:        "health_check_source_ip",
					Description: `Source IP of Packet of Line HealthCheck.`,
				},
				resource.Attribute{
					Name:        "health_check_target_ip",
					Description: `Target IP of Packet of Line HealthCheck.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Router interface ID.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `Router ID.`,
				},
				resource.Attribute{
					Name:        "router_type",
					Description: `Router type.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `Router interface role.`,
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
					Name:        "specification",
					Description: `Router nterface specification.`,
				},
				resource.Attribute{
					Name:        "access_point_id",
					Description: `Access point of the router interface.`,
				},
				resource.Attribute{
					Name:        "opposite_access_point_id",
					Description: `ID of the access point of the peer`,
				},
				resource.Attribute{
					Name:        "opposite_router_type",
					Description: `Peer router type.`,
				},
				resource.Attribute{
					Name:        "opposite_router_id",
					Description: `Peer router ID.`,
				},
				resource.Attribute{
					Name:        "opposite_interface_id",
					Description: `Peer router interface ID.`,
				},
				resource.Attribute{
					Name:        "opposite_interface_owner_id",
					Description: `Peer account ID.`,
				},
				resource.Attribute{
					Name:        "health_check_source_ip",
					Description: `Source IP of Packet of Line HealthCheck.`,
				},
				resource.Attribute{
					Name:        "health_check_target_ip",
					Description: `Target IP of Packet of Line HealthCheck.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_router_interface_connection",
			Category:         "VPC",
			ShortDescription: `Provides a VPC router interface connection resource to connect two VPCs.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"router",
				"interface",
				"connection",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "interface_id",
					Description: `(Required, ForceNew) One side router interface ID.`,
				},
				resource.Attribute{
					Name:        "opposite_interface_id",
					Description: `(Required, ForceNew) Another side router interface ID. It must belong the specified "opposite_interface_owner_id" account.`,
				},
				resource.Attribute{
					Name:        "opposite_interface_owner_id",
					Description: `(Optional, ForceNew) Another side router interface account ID. Log on to the AlibabacloudStack Cloud console, select User Info > Account Management to check the account ID. Default to [Provider account_id](https://www.terraform.io/docs/providers/alibabacloudstack/index.html#account_id).`,
				},
				resource.Attribute{
					Name:        "opposite_router_id",
					Description: `(Optional, ForceNew) Another side router ID. It must belong the specified "opposite_interface_owner_id" account. It is valid when field "opposite_interface_owner_id" is specified.`,
				},
				resource.Attribute{
					Name:        "opposite_router_type",
					Description: `(Optional, ForceNew) Another side router Type. Optional value: VRouter, VBR. It is valid when field "opposite_interface_owner_id" is specified. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Router interface ID. The value is equal to "interface_id".`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Router interface ID. The value is equal to "interface_id".`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_security_group",
			Category:         "ECS",
			ShortDescription: `Provides a Alibabacloudstack Security Group resource.`,
			Description:      ``,
			Keywords: []string{
				"ecs",
				"security",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the security group. Defaults to null.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, Forces new resource) The security group description. Defaults to null.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional, ForceNew) The VPC ID.`,
				},
				resource.Attribute{
					Name:        "inner_access_policy",
					Description: `(Optional) Whether to allow both machines to access each other on all ports in the same security group. Valid values: ["Accept", "Drop"]`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the security group`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the security group`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_security_group_rule",
			Category:         "ECS",
			ShortDescription: `Provides a Alibabacloudstack Security Group Rule resource.`,
			Description:      ``,
			Keywords: []string{
				"ecs",
				"security",
				"group",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `(Required, ForceNew) The type of rule being created. Valid options are ` + "`" + `ingress` + "`" + ` (inbound) or ` + "`" + `egress` + "`" + ` (outbound).`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `(Required, ForceNew) The protocol. Can be ` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `, ` + "`" + `gre` + "`" + ` or ` + "`" + `all` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "port_range",
					Description: `(ForceNew) The range of port numbers relevant to the IP protocol. Default to "-1/-1". When the protocol is tcp or udp, each side port number range from 1 to 65535 and '-1/-1' will be invalid. For example, ` + "`" + `1/200` + "`" + ` means that the range of the port numbers is 1-200. Other protocols' 'port_range' can only be "-1/-1", and other values will be invalid.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Required, ForceNew) The security group to apply this rule to.`,
				},
				resource.Attribute{
					Name:        "nic_type",
					Description: `(Optional, ForceNew) Network type, can be either ` + "`" + `internet` + "`" + ` or ` + "`" + `intranet` + "`" + `, the default value is ` + "`" + `internet` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Optional, ForceNew) Authorization policy, can be either ` + "`" + `accept` + "`" + ` or ` + "`" + `drop` + "`" + `, the default value is ` + "`" + `accept` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional, ForceNew) Authorization policy priority, with parameter values: ` + "`" + `1-100` + "`" + `, default value: 1.`,
				},
				resource.Attribute{
					Name:        "cidr_ip",
					Description: `(Optional, ForceNew) The target IP address range. The default value is 0.0.0.0/0 (which means no restriction will be applied). Other supported formats include 10.159.6.18/12. Only IPv4 is supported.`,
				},
				resource.Attribute{
					Name:        "source_security_group_id",
					Description: `(Optional, ForceNew) The target security group ID within the same region. If this field is specified, the ` + "`" + `nic_type` + "`" + ` can only select ` + "`" + `intranet` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "source_group_owner_account",
					Description: `(Optional, ForceNew) The Alibaba Cloud user account Id of the target security group when security groups are authorized across accounts. This parameter is invalid if ` + "`" + `cidr_ip` + "`" + ` has already been set. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the security group rule`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of rule, ` + "`" + `ingress` + "`" + ` or ` + "`" + `egress` + "`" + ``,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the security group`,
				},
				resource.Attribute{
					Name:        "port_range",
					Description: `The range of port numbers`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `The protocol of the security group rule`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the security group rule`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of rule, ` + "`" + `ingress` + "`" + ` or ` + "`" + `egress` + "`" + ``,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the security group`,
				},
				resource.Attribute{
					Name:        "port_range",
					Description: `The range of port numbers`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `The protocol of the security group rule`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_slb",
			Category:         "Server Load Balancer (SLB)",
			ShortDescription: `Provides an Application Load Balancer resource.`,
			Description:      ``,
			Keywords: []string{
				"server",
				"load",
				"balancer",
				"slb",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the SLB. This name must be unique within your alibabacloudstack account, can have a maximum of 80 characters, must contain only alphanumeric characters or hyphens, such as "-","/",".","_", and must not begin or end with a hyphen. If not specified, Terraform will autogenerate a name beginning with ` + "`" + `tf-lb` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "address_type",
					Description: `(Optional, ForceNew) The network type of the SLB instance. Valid values: ["internet", "intranet"]. If load balancer launched in VPC, this value must be "intranet". - internet: After an Internet SLB instance is created, the system allocates a public IP address so that the instance can forward requests from the Internet. - intranet: After an intranet SLB instance is created, the system allocates an intranet IP address so that the instance can only forward intranet requests.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Required for a VPC SLB, Forces New Resource) The VSwitch ID to launch in. If ` + "`" + `address_type` + "`" + ` is internet, it will be ignore.`,
				},
				resource.Attribute{
					Name:        "specification",
					Description: `(Optional) The specification of the Server Load Balancer instance. Default to empty string indicating it is "Shared-Performance" instance.it is must be specified and it valid values are: ` + "`" + `slb.s1.small` + "`" + `, ` + "`" + `slb.s2.small` + "`" + `, ` + "`" + `slb.s2.medium` + "`" + `, ` + "`" + `slb.s3.small` + "`" + `, ` + "`" + `slb.s3.medium` + "`" + `, ` + "`" + `slb.s3.large` + "`" + ` and ` + "`" + `slb.s4.large` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the load balancer.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The IP address of the load balancer.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the load balancer.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The IP address of the load balancer.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_slb_acl",
			Category:         "Server Load Balancer (SLB)",
			ShortDescription: `Provides a Load Banlancer Access Control List resource.`,
			Description:      ``,
			Keywords: []string{
				"server",
				"load",
				"balancer",
				"slb",
				"acl",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the access control list.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `(Optional, ForceNew) The IP Version of access control list is the type of its entry (IP addresses or CIDR blocks). It values ipv4/ipv6. Our plugin provides a default ip_version: "ipv4".`,
				},
				resource.Attribute{
					Name:        "entry_list",
					Description: `(Optional) A list of entry (IP addresses or CIDR blocks) to be added. At most 50 etnry can be supported in one resource. It contains two sub-fields as ` + "`" + `Entry Block` + "`" + ` follows.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Available in v1.66.0+) A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional, ForceNew, Available in v1.67.0+) Resource group ID. ## Entry Block The entry mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "entry",
					Description: `(Required) An IP addresses or CIDR blocks.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) the comment of the entry. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Id of the access control list.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Id of the access control list.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_slb_backend_server",
			Category:         "Server Load Balancer (SLB)",
			ShortDescription: `Provides an Application Load Balancer Backend Server resource.`,
			Description:      ``,
			Keywords: []string{
				"server",
				"load",
				"balancer",
				"slb",
				"backend",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "load_balancer_id",
					Description: `(Required) ID of the load balancer.`,
				},
				resource.Attribute{
					Name:        "backend_servers",
					Description: `(Required) A list of instances to added backend server in the SLB. It contains two sub-fields as ` + "`" + `Block server` + "`" + ` follows.`,
				},
				resource.Attribute{
					Name:        "delete_protection_validation",
					Description: `(Optional) Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false. ## Block servers The servers mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "server_id",
					Description: `(Required) A list backend server ID (ECS instance ID).`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) Weight of the backend server. Valid value range: [0-100]. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the resource and the value same as load balancer id.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the resource and the value same as load balancer id.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_slb_ca_certificate",
			Category:         "Server Load Balancer (SLB)",
			ShortDescription: `Provides a Load Banlancer CA Certificate resource.`,
			Description:      ``,
			Keywords: []string{
				"server",
				"load",
				"balancer",
				"slb",
				"ca",
				"certificate",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the CA Certificate.`,
				},
				resource.Attribute{
					Name:        "ca_certificate",
					Description: `(Required, ForceNew) the content of the CA certificate. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Id of CA Certificate .`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Id of CA Certificate .`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_slb_domain_extension",
			Category:         "Server Load Balancer (SLB)",
			ShortDescription: `Provides a Load Banlancer domain extension Resource and add it to one Listener.`,
			Description:      ``,
			Keywords: []string{
				"server",
				"load",
				"balancer",
				"slb",
				"domain",
				"extension",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "load_balancer_id",
					Description: `(Required, ForceNew) The ID of the SLB instance.`,
				},
				resource.Attribute{
					Name:        "frontend_port",
					Description: `(Required, ForceNew) The frontend port used by the HTTPS listener of the SLB instance. Valid values: 1–65535.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional, ForceNew) The domain name,`,
				},
				resource.Attribute{
					Name:        "server_certificate_id",
					Description: `(Required) The ID of the certificate used by the domain name.`,
				},
				resource.Attribute{
					Name:        "delete_protection_validation",
					Description: `(Optional, Available in 1.63.0+) Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the domain extension.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the domain extension.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_slb_listener",
			Category:         "Server Load Balancer (SLB)",
			ShortDescription: `Provides an Application Load Balancer resource.`,
			Description:      ``,
			Keywords: []string{
				"server",
				"load",
				"balancer",
				"slb",
				"listener",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "load_balancer_id",
					Description: `(Required, ForceNew) The Load Balancer ID which is used to launch a new listener.`,
				},
				resource.Attribute{
					Name:        "frontend_port",
					Description: `(Required, ForceNew) Port used by the Server Load Balancer instance frontend. Valid value range: [1-65535].`,
				},
				resource.Attribute{
					Name:        "backend_port",
					Description: `(Required, ForceNew) Port used by the Server Load Balancer instance backend. Valid value range: [1-65535].`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional, ForceNew) The protocol to listen on. Valid values are [` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + `, ` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `].`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Required) Bandwidth peak of Listener. For the public network instance charged per traffic consumed, the Bandwidth on Listener can be set to -1, indicating the bandwidth peak is unlimited. Valid values are [-1, 1-5000] in Mbps.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of slb listener. This description can have a string of 1 to 80 characters. Default value: null.`,
				},
				resource.Attribute{
					Name:        "scheduler",
					Description: `(Optional) Scheduling algorithm, Valid values are ` + "`" + `wrr` + "`" + `, ` + "`" + `rr` + "`" + ` and ` + "`" + `wlc` + "`" + `. Default to "wrr".`,
				},
				resource.Attribute{
					Name:        "sticky_session",
					Description: `(Required) Whether to enable session persistence, Valid values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. Default to ` + "`" + `off` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sticky_session_type",
					Description: `(Optional) Mode for handling the cookie. If ` + "`" + `sticky_session` + "`" + ` is "on", it is mandatory. Otherwise, it will be ignored. Valid values are ` + "`" + `insert` + "`" + ` and ` + "`" + `server` + "`" + `. ` + "`" + `insert` + "`" + ` means it is inserted from Server Load Balancer; ` + "`" + `server` + "`" + ` means the Server Load Balancer learns from the backend server.`,
				},
				resource.Attribute{
					Name:        "cookie_timeout",
					Description: `(Optional) Cookie timeout. It is mandatory when ` + "`" + `sticky_session` + "`" + ` is "on" and ` + "`" + `sticky_session_type` + "`" + ` is "insert". Otherwise, it will be ignored. Valid value range: [1-86400] in seconds.`,
				},
				resource.Attribute{
					Name:        "cookie",
					Description: `(Optional) The cookie configured on the server. It is mandatory when ` + "`" + `sticky_session` + "`" + ` is "on" and ` + "`" + `sticky_session_type` + "`" + ` is "server". Otherwise, it will be ignored. Valid value：String in line with RFC 2965, with length being 1- 200. It only contains characters such as ASCII codes, English letters and digits instead of the comma, semicolon or spacing, and it cannot start with $.`,
				},
				resource.Attribute{
					Name:        "persistence_timeout",
					Description: `(Optional) Timeout of connection persistence. Valid value range: [0-3600] in seconds. Default to 0 and means closing it.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `(Required) Whether to enable health check. Valid values are` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. TCP and UDP listener's HealthCheck is always on, so it will be ignore when launching TCP or UDP listener.`,
				},
				resource.Attribute{
					Name:        "health_check_type",
					Description: `(Optional) Type of health check. Valid values are: ` + "`" + `tcp` + "`" + ` and ` + "`" + `http` + "`" + `. Default to ` + "`" + `tcp` + "`" + ` . TCP supports TCP and HTTP health check mode, you can select the particular mode depending on your application.`,
				},
				resource.Attribute{
					Name:        "health_check_domain",
					Description: `(Optional) Domain name used for health check. When it used to launch TCP listener, ` + "`" + `health_check_type` + "`" + ` must be "http". Its length is limited to 1-80 and only characters such as letters, digits, ‘-‘ and ‘.’ are allowed. When it is not set or empty, Server Load Balancer uses the private network IP address of each backend server as Domain used for health check.`,
				},
				resource.Attribute{
					Name:        "health_check_uri",
					Description: `(Optional) URI used for health check. When it used to launch TCP listener, ` + "`" + `health_check_type` + "`" + ` must be "http". Its length is limited to 1-80 and it must start with /. Only characters such as letters, digits, ‘-’, ‘/’, ‘.’, ‘%’, ‘?’, #’ and ‘&’ are allowed.`,
				},
				resource.Attribute{
					Name:        "health_check_connect_port",
					Description: `(Optional) Port used for health check. Valid value range: [1-65535]. Default to "None" means the backend server port is used.`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `(Optional) Threshold determining the result of the health check is success. It is required when ` + "`" + `health_check` + "`" + ` is on. Valid value range: [1-10] in seconds. Default to 3.`,
				},
				resource.Attribute{
					Name:        "unhealthy_threshold",
					Description: `(Optional) Threshold determining the result of the health check is fail. It is required when ` + "`" + `health_check` + "`" + ` is on. Valid value range: [1-10] in seconds. Default to 3.`,
				},
				resource.Attribute{
					Name:        "health_check_timeout",
					Description: `(Optional) Maximum timeout of each health check response. It is required when ` + "`" + `health_check` + "`" + ` is on. Valid value range: [1-300] in seconds. Default to 5. Note: If ` + "`" + `health_check_timeout` + "`" + ` < ` + "`" + `health_check_interval` + "`" + `, its will be replaced by ` + "`" + `health_check_interval` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_check_interval",
					Description: `(Optional) Time interval of health checks. It is required when ` + "`" + `health_check` + "`" + ` is on. Valid value range: [1-50] in seconds. Default to 2.`,
				},
				resource.Attribute{
					Name:        "health_check_http_code",
					Description: `(Optional) Regular health check HTTP status code. Multiple codes are segmented by “,”. It is required when ` + "`" + `health_check` + "`" + ` is on. Default to ` + "`" + `http_2xx` + "`" + `. Valid values are: ` + "`" + `http_2xx` + "`" + `, ` + "`" + `http_3xx` + "`" + `, ` + "`" + `http_4xx` + "`" + ` and ` + "`" + `http_5xx` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_check_method",
					Description: `(Optional) HealthCheckMethod used for health check.` + "`" + `http` + "`" + ` and ` + "`" + `https` + "`" + ` support regions ap-northeast-1, ap-southeast-1, ap-southeast-2, ap-southeast-3, us-east-1, us-west-1, eu-central-1, ap-south-1, me-east-1, cn-huhehaote, cn-zhangjiakou, ap-southeast-5, cn-shenzhen, cn-hongkong, cn-qingdao, cn-chengdu, eu-west-1, cn-hangzhou", cn-beijing, cn-shanghai.This function does not support the TCP protocol .`,
				},
				resource.Attribute{
					Name:        "server_certificate_id",
					Description: `(Required) SLB Server certificate ID. It is required when ` + "`" + `protocol` + "`" + ` is ` + "`" + `https` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "gzip",
					Description: `(Optional) Whether to enable "Gzip Compression". If enabled, files of specific file types will be compressed, otherwise, no files will be compressed. Default to true.`,
				},
				resource.Attribute{
					Name:        "x_forwarded_for",
					Description: `(Optional) Whether to set additional HTTP Header field "X-Forwarded-For" (documented below).`,
				},
				resource.Attribute{
					Name:        "established_timeout",
					Description: `(Optional) Timeout of tcp listener established connection idle timeout. Valid value range: [10-900] in seconds. Default to 900.`,
				},
				resource.Attribute{
					Name:        "server_group_id",
					Description: `(Optional) the id of server group to be apply on the listener, is the id of resource ` + "`" + `alibabacloudstack_slb_server_group` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "listener_forward",
					Description: `(Optional, ForceNew) Whether to enable http redirect to https, Valid values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. Default to ` + "`" + `off` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "forward_port",
					Description: `(Optional, ForceNew) The port that http redirect to https.`,
				},
				resource.Attribute{
					Name:        "health_check_method",
					Description: `(Optional, ForceNew, Available in 1.70.0+) The method of health check. Valid values: ["head", "get"].`,
				},
				resource.Attribute{
					Name:        "delete_protection_validation",
					Description: `(Optional, Available in 1.63.0+) Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false. ->`,
				},
				resource.Attribute{
					Name:        "retrive_slb_ip",
					Description: `(Optional) Whether to use the XForwardedFor_SLBIP header to obtain the public IP address of the SLB instance. Default to false.`,
				},
				resource.Attribute{
					Name:        "retrive_slb_id",
					Description: `(Optional) Whether to use the XForwardedFor header to obtain the ID of the SLB instance. Default to false.`,
				},
				resource.Attribute{
					Name:        "retrive_slb_proto",
					Description: `(Optional) Whether to use the XForwardedFor_proto header to obtain the protocol used by the listener. Default to false. ## Listener fields and protocol mapping load balance support 4 protocal to listen on, they are ` + "`" + `http` + "`" + `,` + "`" + `https` + "`" + `,` + "`" + `tcp` + "`" + `,` + "`" + `udp` + "`" + `, the every listener support which portocal following: listener parameter | support protocol | value range | ------------- | ------------- | ------------- | backend_port | http & https & tcp & udp | 1-65535 | frontend_port | http & https & tcp & udp | 1-65535 | protocol | http & https & tcp & udp | bandwidth | http & https & tcp & udp | -1 / 1-5000 | scheduler | http & https & tcp & udp | wrr rr or wlc | sticky_session | http & https | on or off | sticky_session_type | http & https | insert or server | cookie_timeout | http & https | 1-86400 | cookie | http & https | | persistence_timeout | tcp & udp | 0-3600 | health_check | http & https | on or off | health_check_type | tcp | tcp or http | health_check_domain | http & https & tcp | health_check_method | http & https & tcp | health_check_uri | http & https & tcp | | health_check_connect_port | http & https & tcp & udp | 1-65535 or -520 | healthy_threshold | http & https & tcp & udp | 1-10 | unhealthy_threshold | http & https & tcp & udp | 1-10 | health_check_timeout | http & https & tcp & udp | 1-300 | health_check_interval | http & https & tcp & udp | 1-50 | health_check_http_code | http & https & tcp | http_2xx,http_3xx,http_4xx,http_5xx | server_certificate_id | https | | gzip | http & https | true or false | x_forwarded_for | http & https | | established_timeout | tcp | 10-900| server_group_id | http & https & tcp & udp | the id of resource alibabacloudstack_slb_server_group | The listener mapping supports the following: ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the load balancer listener. Its format as ` + "`" + `<load_balancer_id>:<protocol>:<frontend_port>` + "`" + `. Before verson 1.57.1, the foramt as ` + "`" + `<load_balancer_id>:<frontend_port>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "load_balancer_id",
					Description: `The Load Balancer ID which is used to launch a new listener.`,
				},
				resource.Attribute{
					Name:        "frontend_port",
					Description: `Port used by the Server Load Balancer instance frontend.`,
				},
				resource.Attribute{
					Name:        "backend_port",
					Description: `Port used by the Server Load Balancer instance backend.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol to listen on.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `Bandwidth peak of Listener.`,
				},
				resource.Attribute{
					Name:        "scheduler",
					Description: `Scheduling algorithm.`,
				},
				resource.Attribute{
					Name:        "sticky_session",
					Description: `Whether to enable session persistence.`,
				},
				resource.Attribute{
					Name:        "sticky_session_type",
					Description: `Mode for handling the cookie.`,
				},
				resource.Attribute{
					Name:        "cookie_timeout",
					Description: `Cookie timeout.`,
				},
				resource.Attribute{
					Name:        "cookie",
					Description: `The cookie configured on the server.`,
				},
				resource.Attribute{
					Name:        "persistence_timeout",
					Description: `Timeout of connection persistence.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `Whether to enable health check.`,
				},
				resource.Attribute{
					Name:        "health_check_type",
					Description: `Type of health check.`,
				},
				resource.Attribute{
					Name:        "health_check_domain",
					Description: `Domain name used for health check.`,
				},
				resource.Attribute{
					Name:        "health_check_method",
					Description: `HealthCheckMethod used for health check.`,
				},
				resource.Attribute{
					Name:        "health_check_uri",
					Description: `URI used for health check.`,
				},
				resource.Attribute{
					Name:        "health_check_connect_port",
					Description: `Port used for health check.`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `Threshold determining the result of the health check is success.`,
				},
				resource.Attribute{
					Name:        "unhealthy_threshold",
					Description: `Threshold determining the result of the health check is fail.`,
				},
				resource.Attribute{
					Name:        "health_check_timeout",
					Description: `Maximum timeout of each health check response.`,
				},
				resource.Attribute{
					Name:        "health_check_interval",
					Description: `Time interval of health checks.`,
				},
				resource.Attribute{
					Name:        "health_check_http_code",
					Description: `Regular health check HTTP status code.`,
				},
				resource.Attribute{
					Name:        "server_certificate_id",
					Description: `(Optional) Security certificate ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the load balancer listener. Its format as ` + "`" + `<load_balancer_id>:<protocol>:<frontend_port>` + "`" + `. Before verson 1.57.1, the foramt as ` + "`" + `<load_balancer_id>:<frontend_port>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "load_balancer_id",
					Description: `The Load Balancer ID which is used to launch a new listener.`,
				},
				resource.Attribute{
					Name:        "frontend_port",
					Description: `Port used by the Server Load Balancer instance frontend.`,
				},
				resource.Attribute{
					Name:        "backend_port",
					Description: `Port used by the Server Load Balancer instance backend.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol to listen on.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `Bandwidth peak of Listener.`,
				},
				resource.Attribute{
					Name:        "scheduler",
					Description: `Scheduling algorithm.`,
				},
				resource.Attribute{
					Name:        "sticky_session",
					Description: `Whether to enable session persistence.`,
				},
				resource.Attribute{
					Name:        "sticky_session_type",
					Description: `Mode for handling the cookie.`,
				},
				resource.Attribute{
					Name:        "cookie_timeout",
					Description: `Cookie timeout.`,
				},
				resource.Attribute{
					Name:        "cookie",
					Description: `The cookie configured on the server.`,
				},
				resource.Attribute{
					Name:        "persistence_timeout",
					Description: `Timeout of connection persistence.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `Whether to enable health check.`,
				},
				resource.Attribute{
					Name:        "health_check_type",
					Description: `Type of health check.`,
				},
				resource.Attribute{
					Name:        "health_check_domain",
					Description: `Domain name used for health check.`,
				},
				resource.Attribute{
					Name:        "health_check_method",
					Description: `HealthCheckMethod used for health check.`,
				},
				resource.Attribute{
					Name:        "health_check_uri",
					Description: `URI used for health check.`,
				},
				resource.Attribute{
					Name:        "health_check_connect_port",
					Description: `Port used for health check.`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `Threshold determining the result of the health check is success.`,
				},
				resource.Attribute{
					Name:        "unhealthy_threshold",
					Description: `Threshold determining the result of the health check is fail.`,
				},
				resource.Attribute{
					Name:        "health_check_timeout",
					Description: `Maximum timeout of each health check response.`,
				},
				resource.Attribute{
					Name:        "health_check_interval",
					Description: `Time interval of health checks.`,
				},
				resource.Attribute{
					Name:        "health_check_http_code",
					Description: `Regular health check HTTP status code.`,
				},
				resource.Attribute{
					Name:        "server_certificate_id",
					Description: `(Optional) Security certificate ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_slb_master_slave_server_group",
			Category:         "Server Load Balancer (SLB)",
			ShortDescription: `Provides a Load Balancer Master Slave Server Group resource.`,
			Description:      ``,
			Keywords: []string{
				"server",
				"load",
				"balancer",
				"slb",
				"master",
				"slave",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "load_balancer_id",
					Description: `(Required, ForceNew) The Load Balancer ID which is used to launch a new master slave server group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, ForceNew) Name of the master slave server group.`,
				},
				resource.Attribute{
					Name:        "servers",
					Description: `(Optional, ForceNew) A list of ECS instances to be added. Only two ECS instances can be supported in one resource. It contains six sub-fields as ` + "`" + `Block server` + "`" + ` follows.`,
				},
				resource.Attribute{
					Name:        "delete_protection_validation",
					Description: `(Optional) Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false. ## Block servers The servers mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "server_ids",
					Description: `(Required) A list backend server ID (ECS instance ID).`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) The port used by the backend server. Valid value range: [1-65535].`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) Weight of the backend server. Valid value range: [0-100]. Default to 100.`,
				},
				resource.Attribute{
					Name:        "server_type",
					Description: `(Optional) The server type of the backend server. Valid value Master, Slave. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the master slave server group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the master slave server group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_slb_rule",
			Category:         "Server Load Balancer (SLB)",
			ShortDescription: `Provides a Load Banlancer Forwarding Rule Resource and add it to one Listener.`,
			Description:      ``,
			Keywords: []string{
				"server",
				"load",
				"balancer",
				"slb",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "load_balancer_id",
					Description: `(Required, ForceNew) The Load Balancer ID which is used to launch the new forwarding rule.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the forwarding rule.`,
				},
				resource.Attribute{
					Name:        "frontend_port",
					Description: `(Required, ForceNew) The listener frontend port which is used to launch the new forwarding rule. Valid range: [1-65535].`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional, ForceNew) Domain name of the forwarding rule. It can contain letters a-z, numbers 0-9, hyphens (-), and periods (.), and wildcard characters. The following two domain name formats are supported: - Standard domain name: www.test.com - Wildcard domain name:`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Optional, ForceNew) Domain of the forwarding rule. It must be 2-80 characters in length. Only letters a-z, numbers 0-9, and characters '-' '/' '?' '%' '#' and '&' are allowed. URLs must be started with the character '/', but cannot be '/' alone.`,
				},
				resource.Attribute{
					Name:        "server_group_id",
					Description: `(Required) ID of a virtual server group that will be forwarded.`,
				},
				resource.Attribute{
					Name:        "scheduler",
					Description: `(Optional) Scheduling algorithm, Valid values are ` + "`" + `wrr` + "`" + `, ` + "`" + `rr` + "`" + ` and ` + "`" + `wlc` + "`" + `. Default to "wrr". This parameter is required and takes effect only when ListenerSync is set to off.`,
				},
				resource.Attribute{
					Name:        "sticky_session",
					Description: `(Optional) Whether to enable session persistence, Valid values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. Default to ` + "`" + `off` + "`" + `. This parameter is required and takes effect only when ListenerSync is set to off.`,
				},
				resource.Attribute{
					Name:        "sticky_session_type",
					Description: `(Optional) Mode for handling the cookie. If ` + "`" + `sticky_session` + "`" + ` is "on", it is mandatory. Otherwise, it will be ignored. Valid values are ` + "`" + `insert` + "`" + ` and ` + "`" + `server` + "`" + `. ` + "`" + `insert` + "`" + ` means it is inserted from Server Load Balancer; ` + "`" + `server` + "`" + ` means the Server Load Balancer learns from the backend server.`,
				},
				resource.Attribute{
					Name:        "cookie_timeout",
					Description: `(Optional) Cookie timeout. It is mandatory when ` + "`" + `sticky_session` + "`" + ` is "on" and ` + "`" + `sticky_session_type` + "`" + ` is "insert". Otherwise, it will be ignored. Valid value range: [1-86400] in seconds.`,
				},
				resource.Attribute{
					Name:        "cookie",
					Description: `(Optional) The cookie configured on the server. It is mandatory when ` + "`" + `sticky_session` + "`" + ` is "on" and ` + "`" + `sticky_session_type` + "`" + ` is "server". Otherwise, it will be ignored. Valid value：String in line with RFC 2965, with length being 1- 200. It only contains characters such as ASCII codes, English letters and digits instead of the comma, semicolon or spacing, and it cannot start with $.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `(Optional) Whether to enable health check. Valid values are` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. TCP and UDP listener's HealthCheck is always on, so it will be ignore when launching TCP or UDP listener. This parameter is required and takes effect only when ListenerSync is set to off.`,
				},
				resource.Attribute{
					Name:        "health_check_domain",
					Description: `(Optional) Domain name used for health check. When it used to launch TCP listener, ` + "`" + `health_check_type` + "`" + ` must be "http". Its length is limited to 1-80 and only characters such as letters, digits, ‘-‘ and ‘.’ are allowed. When it is not set or empty, Server Load Balancer uses the private network IP address of each backend server as Domain used for health check.`,
				},
				resource.Attribute{
					Name:        "health_check_uri",
					Description: `(Optional) URI used for health check. When it used to launch TCP listener, ` + "`" + `health_check_type` + "`" + ` must be "http". Its length is limited to 1-80 and it must start with /. Only characters such as letters, digits, ‘-’, ‘/’, ‘.’, ‘%’, ‘?’, #’ and ‘&’ are allowed.`,
				},
				resource.Attribute{
					Name:        "health_check_connect_port",
					Description: `(Optional) Port used for health check. Valid value range: [1-65535]. Default to "None" means the backend server port is used.`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `(Optional) Threshold determining the result of the health check is success. It is required when ` + "`" + `health_check` + "`" + ` is on. Valid value range: [1-10] in seconds. Default to 3.`,
				},
				resource.Attribute{
					Name:        "unhealthy_threshold",
					Description: `(Optional) Threshold determining the result of the health check is fail. It is required when ` + "`" + `health_check` + "`" + ` is on. Valid value range: [1-10] in seconds. Default to 3.`,
				},
				resource.Attribute{
					Name:        "health_check_timeout",
					Description: `(Optional) Maximum timeout of each health check response. It is required when ` + "`" + `health_check` + "`" + ` is on. Valid value range: [1-300] in seconds. Default to 5. Note: If ` + "`" + `health_check_timeout` + "`" + ` < ` + "`" + `health_check_interval` + "`" + `, its will be replaced by ` + "`" + `health_check_interval` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_check_interval",
					Description: `(Optional) Time interval of health checks. It is required when ` + "`" + `health_check` + "`" + ` is on. Valid value range: [1-50] in seconds. Default to 2.`,
				},
				resource.Attribute{
					Name:        "health_check_http_code",
					Description: `(Optional) Regular health check HTTP status code. Multiple codes are segmented by “,”. It is required when ` + "`" + `health_check` + "`" + ` is on. Default to ` + "`" + `http_2xx` + "`" + `. Valid values are: ` + "`" + `http_2xx` + "`" + `, ` + "`" + `http_3xx` + "`" + `, ` + "`" + `http_4xx` + "`" + ` and ` + "`" + `http_5xx` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "listener_sync",
					Description: `(Optional) Indicates whether a forwarding rule inherits the settings of a health check , session persistence, and scheduling algorithm from a listener. Default to on.`,
				},
				resource.Attribute{
					Name:        "delete_protection_validation",
					Description: `(Optional) Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the forwarding rule.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the forwarding rule.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_slb_server_certificate",
			Category:         "Server Load Balancer (SLB)",
			ShortDescription: `Provides a Load Banlancer Server Certificate resource.`,
			Description:      ``,
			Keywords: []string{
				"server",
				"load",
				"balancer",
				"slb",
				"certificate",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the Server Certificate.`,
				},
				resource.Attribute{
					Name:        "server_certificate",
					Description: `(Optional, ForceNew) the content of the ssl certificate. where ` + "`" + `alibabacloudstack_certificate_id` + "`" + ` is null, it is required, otherwise it is ignored.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Optional, ForceNew) the content of privat key of the ssl certificate specified by ` + "`" + `server_certificate` + "`" + `. where ` + "`" + `alibabacloudstack_certificate_id` + "`" + ` is null, it is required, otherwise it is ignored. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Id of Server Certificate (SSL Certificate).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Id of Server Certificate (SSL Certificate).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_slb_server_group",
			Category:         "Server Load Balancer (SLB)",
			ShortDescription: `Provides a Load Balancer Virtual Backend Server Group resource.`,
			Description:      ``,
			Keywords: []string{
				"server",
				"load",
				"balancer",
				"slb",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "load_balancer_id",
					Description: `(Required, ForceNew) The Load Balancer ID which is used to launch a new virtual server group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the virtual server group. Our plugin provides a default name: "tf-server-group".`,
				},
				resource.Attribute{
					Name:        "servers",
					Description: `A list of ECS instances to be added. At most 20 ECS instances can be supported in one resource. It contains three sub-fields as ` + "`" + `Block server` + "`" + ` follows.`,
				},
				resource.Attribute{
					Name:        "delete_protection_validation",
					Description: `(Optional) Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false. ## Block servers The servers mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "server_ids",
					Description: `(Required) A list backend server ID (ECS instance ID).`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) The port used by the backend server. Valid value range: [1-65535].`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) Weight of the backend server. Valid value range: [0-100]. Default to 100.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type of the backend server. Valid value ecs, eni. Default to eni. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the virtual server group.`,
				},
				resource.Attribute{
					Name:        "load_balancer_id",
					Description: `The Load Balancer ID which is used to launch a new virtual server group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the virtual server group.`,
				},
				resource.Attribute{
					Name:        "servers",
					Description: `A list of ECS instances that have be added.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the virtual server group.`,
				},
				resource.Attribute{
					Name:        "load_balancer_id",
					Description: `The Load Balancer ID which is used to launch a new virtual server group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the virtual server group.`,
				},
				resource.Attribute{
					Name:        "servers",
					Description: `A list of ECS instances that have be added.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_snapshot",
			Category:         "ECS",
			ShortDescription: `Provides an ECS snapshot resource.`,
			Description:      ``,
			Keywords: []string{
				"ecs",
				"snapshot",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "disk_id",
					Description: `(Required, ForceNew) The ID of the disk.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, ForceNew) Name of the snapshot. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-", ".", "_", and must not begin or end with a hyphen, and must not begin with http:// or https://.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, ForceNew) Description of the snapshot. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. ### Timeouts`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 2 mins) Used when creating the snapshot (until it reaches the initial ` + "`" + `SnapshotCreatingAccomplished` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 2 mins) Used when terminating the snapshot. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The snapshot ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The snapshot ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_snapshot_policy",
			Category:         "ECS",
			ShortDescription: `Provides an ECS snapshot policy resource.`,
			Description:      ``,
			Keywords: []string{
				"ecs",
				"snapshot",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The snapshot policy name.`,
				},
				resource.Attribute{
					Name:        "repeat_weekdays",
					Description: `(Required) The automatic snapshot repetition dates. The unit of measurement is day and the repeating cycle is a week. Value range: [1, 7], which represents days starting from Monday to Sunday, for example 1 indicates Monday. When you want to schedule multiple automatic snapshot tasks for a disk in a week, you can set the RepeatWeekdays to an array. - A maximum of seven time points can be selected. - The format is an JSON array of ["1", "2", … "7"] and the time points are separated by commas (,). for example, ["1", "2", ... "7"].`,
				},
				resource.Attribute{
					Name:        "retention_days",
					Description: `(Optional) The snapshot retention time, and the unit of measurement is day. Optional values: - -1: The automatic snapshots are retained permanently. - [1, 65536]: The number of days retained. Default value: -1.`,
				},
				resource.Attribute{
					Name:        "time_points",
					Description: `(Required) The automatic snapshot creation schedule, and the unit of measurement is hour. Value range: [0, 23], which represents from 00:00 to 24:00, for example 1 indicates 01:00. When you want to schedule multiple automatic snapshot tasks for a disk in a day, you can set the TimePoints to an array. - A maximum of 24 time points can be selected. - The format is an JSON array of ["0", "1", … "23"] and the time points are separated by commas (,).`,
				},
				resource.Attribute{
					Name:        "disk_ids",
					Description: `(Optional) The IDs of the disks to apply an automatic snapshot policy.`,
				},
				resource.Attribute{
					Name:        "enable_automated_snapshot_policy",
					Description: `(Optional) To apply an automatic snapshot policy to one or more disks. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The snapshot policy ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The snapshot policy ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_snat",
			Category:         "VPC",
			ShortDescription: `Provides a Alibabacloudstack snat resource.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"snat",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "snat_table_id",
					Description: `(Required, ForceNew) The value can get from ` + "`" + `alibabacloudstack_nat_gateway` + "`" + ` Attributes "snat_table_ids".`,
				},
				resource.Attribute{
					Name:        "source_vswitch_id",
					Description: `(Optional, ForceNew) The vswitch ID.`,
				},
				resource.Attribute{
					Name:        "source_cidr",
					Description: `(Optional, ForceNew) The private network segment of Ecs. This parameter and the ` + "`" + `source_vswitch_id` + "`" + ` parameter are mutually exclusive and cannot appear at the same time.`,
				},
				resource.Attribute{
					Name:        "snat_ip",
					Description: `(Required) The SNAT ip address, the ip must along bandwidth package public ip which ` + "`" + `alibabacloudstack_nat_gateway` + "`" + ` argument ` + "`" + `bandwidth_packages` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the snat entry. The value formats as ` + "`" + `<snat_table_id>:<snat_entry_id>` + "`" + ``,
				},
				resource.Attribute{
					Name:        "snat_entry_id",
					Description: `The id of the snat entry on the server.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the snat entry. The value formats as ` + "`" + `<snat_table_id>:<snat_entry_id>` + "`" + ``,
				},
				resource.Attribute{
					Name:        "snat_entry_id",
					Description: `The id of the snat entry on the server.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_vpc",
			Category:         "VPC",
			ShortDescription: `Provides a Alibabacloudstack VPC resource.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Required, ForceNew) The CIDR block for the VPC. The ` + "`" + `cidr_block` + "`" + ` is Optional and default value is ` + "`" + `172.16.0.0/12` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vpc_name",
					Description: `(Optional) The name of the VPC. Defaults to null.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Field ` + "`" + `name` + "`" + ` has been deprecated from provider. New field ` + "`" + `vpc_name` + "`" + ` instead.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The VPC description. Defaults to null.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional) The Id of resource group which the VPC belongs.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "secondary_cidr_blocks",
					Description: `(Optional) The secondary CIDR blocks for the VPC.`,
				},
				resource.Attribute{
					Name:        "dry_run",
					Description: `(Optional, ForceNew) Specifies whether to precheck this request only. Valid values: ` + "`" + `true` + "`" + ` and ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user_cidrs",
					Description: `(Optional, ForceNew) The user cidrs of the VPC.`,
				},
				resource.Attribute{
					Name:        "enable_ipv6",
					Description: `(Optional, ForceNew) Specifies whether to enable the IPv6 CIDR block. Valid values: ` + "`" + `false` + "`" + ` (Default): disables IPv6 CIDR blocks. ` + "`" + `true` + "`" + `: enables IPv6 CIDR blocks. If the ` + "`" + `enable_ipv6` + "`" + ` is ` + "`" + `true` + "`" + `, the system will automatically create a free version of an IPv6 gateway for your private network and assign an IPv6 network segment assigned as /56. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 10 mins) Used when creating the vpc (until it reaches the initial ` + "`" + `Available` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 10 mins) Used when terminating the vpc. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the VPC.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The CIDR block for the VPC.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the VPC.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the VPC.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `The ID of the router created by default on VPC creation.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The route table ID of the router created by default on VPC creation.`,
				},
				resource.Attribute{
					Name:        "ipv6_cidr_block",
					Description: `The ipv6 cidr block of VPC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the VPC.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The CIDR block for the VPC.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the VPC.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the VPC.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `The ID of the router created by default on VPC creation.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The route table ID of the router created by default on VPC creation.`,
				},
				resource.Attribute{
					Name:        "ipv6_cidr_block",
					Description: `The ipv6 cidr block of VPC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_vpc_ipv6_egress_rule",
			Category:         "VPC",
			ShortDescription: `Provides a Alibabacloudstack VPC Ipv6 Egress Rule resource.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"ipv6",
				"egress",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, ForceNew) The description of the egress-only rule. The description must be ` + "`" + `2` + "`" + ` to ` + "`" + `256` + "`" + ` characters in length. It cannot start with ` + "`" + `http://` + "`" + ` or ` + "`" + `https://` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, ForceNew) The ID of the IPv6 address to which you want to apply the egress-only rule.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Optional, Computed, ForceNew) The type of instance to which you want to apply the egress-only rule. Valid values: ` + "`" + `Ipv6Address` + "`" + `. ` + "`" + `Ipv6Address` + "`" + ` (default): an IPv6 address.`,
				},
				resource.Attribute{
					Name:        "ipv6_egress_rule_name",
					Description: `(Optional, ForceNew) The name of the egress-only rule. The name must be ` + "`" + `2` + "`" + ` to ` + "`" + `128` + "`" + ` characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter but cannot start with ` + "`" + `http://` + "`" + ` or ` + "`" + `https://` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ipv6_gateway_id",
					Description: `(Required, ForceNew) The ID of the IPv6 gateway. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in terraform of Ipv6 Egress Rule. The value formats as ` + "`" + `<ipv6_gateway_id>:<ipv6_egress_rule_id>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the resource. Valid values: ` + "`" + `Available` + "`" + `, ` + "`" + `Pending` + "`" + ` and ` + "`" + `Deleting` + "`" + `. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 1 mins) Used when create the Ipv6 Egress Rule.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 1 mins) Used when delete the Ipv6 Egress Rule. ## Import VPC Ipv6 Egress Rule can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_vpc_ipv6_egress_rule.example <ipv6_gateway_id>:<ipv6_egress_rule_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in terraform of Ipv6 Egress Rule. The value formats as ` + "`" + `<ipv6_gateway_id>:<ipv6_egress_rule_id>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the resource. Valid values: ` + "`" + `Available` + "`" + `, ` + "`" + `Pending` + "`" + ` and ` + "`" + `Deleting` + "`" + `. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 1 mins) Used when create the Ipv6 Egress Rule.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 1 mins) Used when delete the Ipv6 Egress Rule. ## Import VPC Ipv6 Egress Rule can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_vpc_ipv6_egress_rule.example <ipv6_gateway_id>:<ipv6_egress_rule_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_vpc_ipv6_gateway",
			Category:         "VPC",
			ShortDescription: `Provides a Alibabacloudstack VPC Ipv6 Gateway resource.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"ipv6",
				"gateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the IPv6 gateway. The description must be ` + "`" + `2` + "`" + ` to ` + "`" + `256` + "`" + ` characters in length. It cannot start with ` + "`" + `http://` + "`" + ` or ` + "`" + `https://` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ipv6_gateway_name",
					Description: `(Optional) The name of the IPv6 gateway. The name must be ` + "`" + `2` + "`" + ` to ` + "`" + `128` + "`" + ` characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter but cannot start with ` + "`" + `http://` + "`" + ` or ` + "`" + `https://` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `(Optional, Computed) The edition of the IPv6 gateway. Valid values: ` + "`" + `Large` + "`" + `, ` + "`" + `Medium` + "`" + ` and ` + "`" + `Small` + "`" + `. ` + "`" + `Small` + "`" + ` (default): Free Edition. ` + "`" + `Medium` + "`" + `: Enterprise Edition . ` + "`" + `Large` + "`" + `: Enhanced Enterprise Edition. The throughput capacity of an IPv6 gateway varies based on the edition. For more information, see [Editions of IPv6 gateways](https://www.alibabacloud.com/help/doc-detail/98926.htm).`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, ForceNew) The ID of the virtual private cloud (VPC) for which you want to create the IPv6 gateway. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in terraform of Ipv6 Gateway.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the resource. Valid values: ` + "`" + `Available` + "`" + `, ` + "`" + `Pending` + "`" + ` and ` + "`" + `Deleting` + "`" + `. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 1 mins) Used when create the Ipv6 Gateway.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 1 mins) Used when update the Ipv6 Gateway.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 5 mins) Used when delete the Ipv6 Gateway. ## Import VPC Ipv6 Gateway can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_vpc_ipv6_gateway.example <id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in terraform of Ipv6 Gateway.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the resource. Valid values: ` + "`" + `Available` + "`" + `, ` + "`" + `Pending` + "`" + ` and ` + "`" + `Deleting` + "`" + `. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 1 mins) Used when create the Ipv6 Gateway.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 1 mins) Used when update the Ipv6 Gateway.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 5 mins) Used when delete the Ipv6 Gateway. ## Import VPC Ipv6 Gateway can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_vpc_ipv6_gateway.example <id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_vpc_ipv6_internet_bandwidth",
			Category:         "VPC",
			ShortDescription: `Provides a Alibabacloudstack VPC Ipv6 Internet Bandwidth resource.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"ipv6",
				"internet",
				"bandwidth",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Required) The amount of Internet bandwidth resources of the IPv6 address, Unit: ` + "`" + `Mbit/s` + "`" + `. Valid values: ` + "`" + `1` + "`" + ` to ` + "`" + `5000` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "internet_charge_type",
					Description: `(Optional, Computed, ForceNew) The metering method of the Internet bandwidth resources of the IPv6 gateway. Valid values: ` + "`" + `PayByBandwidth` + "`" + `, ` + "`" + `PayByTraffic` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ipv6_address_id",
					Description: `(Required, ForceNew) The ID of the IPv6 address.`,
				},
				resource.Attribute{
					Name:        "ipv6_gateway_id",
					Description: `(Required, ForceNew) The ID of the IPv6 gateway. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in terraform of Ipv6 Internet Bandwidth.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the resource.Valid values:` + "`" + `Normal` + "`" + `, ` + "`" + `FinancialLocked` + "`" + ` and ` + "`" + `SecurityLocked` + "`" + `. ## Import VPC Ipv6 Internet Bandwidth can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_vpc_ipv6_internet_bandwidth.example <id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in terraform of Ipv6 Internet Bandwidth.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the resource.Valid values:` + "`" + `Normal` + "`" + `, ` + "`" + `FinancialLocked` + "`" + ` and ` + "`" + `SecurityLocked` + "`" + `. ## Import VPC Ipv6 Internet Bandwidth can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_vpc_ipv6_internet_bandwidth.example <id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_vpn_connection",
			Category:         "VPN",
			ShortDescription: `Provides a Alibabacloudstack VPN connection resource.`,
			Description:      ``,
			Keywords: []string{
				"vpn",
				"connection",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the IPsec connection.`,
				},
				resource.Attribute{
					Name:        "vpn_gateway_id",
					Description: `(Required, ForceNew) The ID of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "customer_gateway_id",
					Description: `(Required, ForceNew) The ID of the customer gateway.`,
				},
				resource.Attribute{
					Name:        "local_subnet",
					Description: `(Required, Type:Set) The CIDR block of the VPC to be connected with the local data center. This parameter is used for phase-two negotiation.`,
				},
				resource.Attribute{
					Name:        "remote_subnet",
					Description: `(Required, Type:Set) The CIDR block of the local data center. This parameter is used for phase-two negotiation.`,
				},
				resource.Attribute{
					Name:        "effect_immediately",
					Description: `(Optional) Whether to delete a successfully negotiated IPsec tunnel and initiate a negotiation again. Valid value:true,false.`,
				},
				resource.Attribute{
					Name:        "ike_config",
					Description: `(Optional) The configurations of phase-one negotiation.`,
				},
				resource.Attribute{
					Name:        "ipsec_config",
					Description: `(Optional) The configurations of phase-two negotiation. ### Block ike_config The ike_config mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "psk",
					Description: `(Optional) Used for authentication between the IPsec VPN gateway and the customer gateway.`,
				},
				resource.Attribute{
					Name:        "ike_version",
					Description: `(Optional) The version of the IKE protocol. Valid value: ikev1 | ikev2. Default value: ikev1`,
				},
				resource.Attribute{
					Name:        "ike_mode",
					Description: `(Optional) The negotiation mode of IKE V1. Valid value: main (main mode) | aggressive (aggressive mode). Default value: main`,
				},
				resource.Attribute{
					Name:        "ike_enc_alg",
					Description: `(Optional) The encryption algorithm of phase-one negotiation. Valid value: aes | aes192 | aes256 | des | 3des. Default Valid value: aes`,
				},
				resource.Attribute{
					Name:        "ike_auth_alg",
					Description: `(Optional) The authentication algorithm of phase-one negotiation. Valid value: md5 | sha1 | sha256 | sha384 | sha512 |. Default value: sha1`,
				},
				resource.Attribute{
					Name:        "ike_pfs",
					Description: `(Optional) The Diffie-Hellman key exchange algorithm used by phase-one negotiation. Valid value: group1 | group2 | group5 | group14 | group24. Default value: group2`,
				},
				resource.Attribute{
					Name:        "ike_lifetime",
					Description: `(Optional) The SA lifecycle as the result of phase-one negotiation. The valid value of n is [0, 86400], the unit is second and the default value is 86400.`,
				},
				resource.Attribute{
					Name:        "ike_local_id",
					Description: `(Optional) The identification of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "ike_remote_id",
					Description: `(Optional) The identification of the customer gateway. ### Block ipsec_config The ipsec_config mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "ipsec_enc_alg",
					Description: `(Optional) The encryption algorithm of phase-two negotiation. Valid value: aes | aes192 | aes256 | des | 3des. Default value: aes`,
				},
				resource.Attribute{
					Name:        "ipsec_auth_alg",
					Description: `(Optional) The authentication algorithm of phase-two negotiation. Valid value: md5 | sha1 | sha256 | sha384 | sha512 |. Default value: sha1`,
				},
				resource.Attribute{
					Name:        "ipsec_pfs",
					Description: `(Optional) The Diffie-Hellman key exchange algorithm used by phase-two negotiation. Valid value: group1 | group2 | group5 | group14 | group24| disabled. Default value: group2`,
				},
				resource.Attribute{
					Name:        "ipsec_lifetime",
					Description: `(Optional) The SA lifecycle as the result of phase-two negotiation. The valid value is [0, 86400], the unit is second and the default value is 86400. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the VPN connection id.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of VPN connection.`,
				},
				resource.Attribute{
					Name:        "ike_config",
					Description: `The configurations of phase-one negotiation.`,
				},
				resource.Attribute{
					Name:        "ipsec_config",
					Description: `The configurations of phase-two negotiation. ## Import VPN connection can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_vpn_connection.example vco-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the VPN connection id.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of VPN connection.`,
				},
				resource.Attribute{
					Name:        "ike_config",
					Description: `The configurations of phase-one negotiation.`,
				},
				resource.Attribute{
					Name:        "ipsec_config",
					Description: `The configurations of phase-two negotiation. ## Import VPN connection can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_vpn_connection.example vco-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_vpn_customer_gateway",
			Category:         "VPN",
			ShortDescription: `Provides a Alibabacloudstack VPN customer gateway resource.`,
			Description:      ``,
			Keywords: []string{
				"vpn",
				"customer",
				"gateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the VPN customer gateway. Defaults to null.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Required, ForceNew) The IP address of the customer gateway.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the VPN customer gateway instance. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the VPN customer gateway instance id. ## Import VPN customer gateway can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_vpn_customer_gateway.example cgw-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the VPN customer gateway instance id. ## Import VPN customer gateway can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_vpn_customer_gateway.example cgw-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_vpn_gateway",
			Category:         "VPN",
			ShortDescription: `Provides a Alibabacloudstack VPN gateway resource.`,
			Description:      ``,
			Keywords: []string{
				"vpn",
				"gateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the VPN. Defaults to null.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, ForceNew) The VPN belongs the vpc_id, the field can't be changed.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `(ForceNew) The charge type for instance. If it is an international site account, the valid value is PostPaid, otherwise PrePaid. Default to PostPaid.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) The filed is only required while the InstanceChargeType is PrePaid. Valid values: [1-9, 12, 24, 36]. Default to 1.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Required) The value should be 10, 100, 200. if the user is postpaid, otherwise it can be 5, 10, 20, 50, 100, 200. It can't be changed by terraform.`,
				},
				resource.Attribute{
					Name:        "enable_ipsec",
					Description: `(Optional) Enable or Disable IPSec VPN. At least one type of VPN should be enabled.`,
				},
				resource.Attribute{
					Name:        "ssl_connections",
					Description: `(Optional) The max connections of SSL VPN. Default to 5. The number of connections supported by each account is different. This field is ignored when enable_ssl is false.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the VPN instance.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Optional, ForceNew) The VPN belongs the vswitch_id, the field can't be changed. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the VPN instance id.`,
				},
				resource.Attribute{
					Name:        "internet_ip",
					Description: `The internet ip of the VPN.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "business_status",
					Description: `The business status of the VPN gateway. ## Import VPN gateway can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_vpn_gateway.example vpn-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the VPN instance id.`,
				},
				resource.Attribute{
					Name:        "internet_ip",
					Description: `The internet ip of the VPN.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "business_status",
					Description: `The business status of the VPN gateway. ## Import VPN gateway can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_vpn_gateway.example vpn-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_vpn_route_entry",
			Category:         "VPN",
			ShortDescription: `Provides a Alibabacloudstack VPN Route Entry resource.`,
			Description:      ``,
			Keywords: []string{
				"vpn",
				"route",
				"entry",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vpn_gateway_id",
					Description: `(Required, ForceNew) The id of the vpn gateway.`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `(Required, ForceNew) The next hop of the destination route.`,
				},
				resource.Attribute{
					Name:        "publish_vpc",
					Description: `(Required) Whether to issue the destination route to the VPC.`,
				},
				resource.Attribute{
					Name:        "route_dest",
					Description: `(Required, ForceNew) The destination network segment of the destination route.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Required) The value should be 0 or 100. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The combination id of the vpn route entry. ## Import VPN route entry can be imported using the id(VpnGatewayId +":"+ NextHop +":"+ RouteDest), e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_vpn_route_entry.example vpn-abc123456:vco-abc123456:10.0.0.10/24 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The combination id of the vpn route entry. ## Import VPN route entry can be imported using the id(VpnGatewayId +":"+ NextHop +":"+ RouteDest), e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alibabacloudstack_vpn_route_entry.example vpn-abc123456:vco-abc123456:10.0.0.10/24 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_vswitch",
			Category:         "VPC",
			ShortDescription: `Provides a Alibabacloudstack VPC switch resource.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"vswitch",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required, ForceNew) The AZ for the switch.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, ForceNew) The VPC ID.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Required, ForceNew) The CIDR block for the switch.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the switch. Defaults to null.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The switch description. Defaults to null. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 10 mins) Used when creating the vswitch (until it reaches the initial ` + "`" + `Available` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 10 mins) Used when terminating the vswitch. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the switch.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The CIDR block for the switch.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The VPC ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the switch.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the switch.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the switch.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The CIDR block for the switch.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The VPC ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the switch.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the switch.`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"alibabacloudstack_adb_account":                           0,
		"alibabacloudstack_adb_backup_policy":                     1,
		"alibabacloudstack_adb_connection":                        2,
		"alibabacloudstack_adb_db_cluster":                        3,
		"alibabacloudstack_alikafka_sasl_acl":                     4,
		"alibabacloudstack_alikafka_sasl_user":                    5,
		"alibabacloudstack_alikafka_topic":                        6,
		"alibabacloudstack_api_gateway_api":                       7,
		"alibabacloudstack_api_gateway_app":                       8,
		"alibabacloudstack_api_gateway_app_attachment":            9,
		"alibabacloudstack_api_gateway_group":                     10,
		"alibabacloudstack_api_gateway_vpc_access":                11,
		"alibabacloudstack_ascm_custom_role":                      12,
		"alibabacloudstack_ascm_logon_policy":                     13,
		"alibabacloudstack_ascm_organization":                     14,
		"alibabacloudstack_ascm_password_policy":                  15,
		"alibabacloudstack_ascm_quota":                            16,
		"alibabacloudstack_ascm_ram_policy":                       17,
		"alibabacloudstack_ascm_ram_policy_for_role":              18,
		"alibabacloudstack_ascm_ram_role":                         19,
		"alibabacloudstack_ascm_resource_group":                   20,
		"alibabacloudstack_ascm_user":                             21,
		"alibabacloudstack_ascm_user_group":                       22,
		"alibabacloudstack_ascm_user_group_resource_set_binding":  23,
		"alibabacloudstack_ascm_user_group_role_binding":          24,
		"alibabacloudstack_ascm_user_role_binding":                25,
		"alibabacloudstack_ascm_usergroup_user":                   26,
		"alibabacloudstack_cms_alarm":                             27,
		"alibabacloudstack_cms_alarm_contact":                     28,
		"alibabacloudstack_cms_alarm_contact_group":               29,
		"alibabacloudstack_cms_site_monitor":                      30,
		"alibabacloudstack_common_bandwidth_package":              31,
		"alibabacloudstack_common_bandwidth_package_attachment":   32,
		"alibabacloudstack_cr_ee_namespace":                       33,
		"alibabacloudstack_cr_ee_repo":                            34,
		"alibabacloudstack_cr_ee_sync_rule":                       35,
		"alibabacloudstack_cr_namespace":                          36,
		"alibabacloudstack_cr_repo":                               37,
		"alibabacloudstack_cs_kubernetes":                         38,
		"alibabacloudstack_cs_kubernetes_node_pool":               39,
		"alibabacloudstack_csb_project":                           40,
		"alibabacloudstack_data_works_connection":                 41,
		"alibabacloudstack_data_works_folder":                     42,
		"alibabacloudstack_data_works_project":                    43,
		"alibabacloudstack_data_works_remind":                     44,
		"alibabacloudstack_data_works_user":                       45,
		"alibabacloudstack_data_works_user_role_binding":          46,
		"alibabacloudstack_datahub_project":                       47,
		"alibabacloudstack_datahub_subscription":                  48,
		"alibabacloudstack_datahub_topic":                         49,
		"alibabacloudstack_db_account":                            50,
		"alibabacloudstack_db_account_privilege":                  51,
		"alibabacloudstack_db_backup_policy":                      52,
		"alibabacloudstack_db_connection":                         53,
		"alibabacloudstack_db_database":                           54,
		"alibabacloudstack_db_instance":                           55,
		"alibabacloudstack_db_read_write_splitting_connection":    56,
		"alibabacloudstack_db_readonly_instance":                  57,
		"alibabacloudstack_dbs_backup_plan":                       58,
		"alibabacloudstack_disk":                                  59,
		"alibabacloudstack_disk_attachment":                       60,
		"alibabacloudstack_dms_enterprise_instance":               61,
		"alibabacloudstack_dms_enterprise_user":                   62,
		"alibabacloudstack_dns_domain":                            63,
		"alibabacloudstack_dns_domain_attachment":                 64,
		"alibabacloudstack_dns_group":                             65,
		"alibabacloudstack_dns_record":                            66,
		"alibabacloudstack_drds_instance":                         67,
		"alibabacloudstack_dts_subscription_job":                  68,
		"alibabacloudstack_dts_synchronization_instance":          69,
		"alibabacloudstack_dts_synchronization_job":               70,
		"alibabacloudstack_ecs_command":                           71,
		"alibabacloudstack_ecs_dedicated_host":                    72,
		"alibabacloudstack_ecs_deployment_set":                    73,
		"alibabacloudstack_ecs_hpc_cluster":                       74,
		"alibabacloudstack_ecs_key_pair":                          75,
		"alibabacloudstack_ecs_key_pair_attachment":               76,
		"alibabacloudstack_ecs_launch_template":                   77,
		"alibabacloudstack_ecs_network_interface":                 78,
		"alibabacloudstack_ecs_network_interface_attachment":      79,
		"alibabacloudstack_edas_application":                      80,
		"alibabacloudstack_edas_cluster":                          81,
		"alibabacloudstack_edas_deploy_group":                     82,
		"alibabacloudstack_edas_instance_cluster_attachment":      83,
		"alibabacloudstack_edas_k8s_application":                  84,
		"alibabacloudstack_edas_k8s_cluster":                      85,
		"alibabacloudstack_edas_slb_attachment":                   86,
		"alibabacloudstack_ehpc_job_template":                     87,
		"alibabacloudstack_eip":                                   88,
		"alibabacloudstack_eip_association":                       89,
		"alibabacloudstack_elasticsearch":                         90,
		"alibabacloudstack_ess_alarm":                             91,
		"alibabacloudstack_ess_attachment":                        92,
		"alibabacloudstack_ess_notification":                      93,
		"alibabacloudstack_ess_scaling_configuration":             94,
		"alibabacloudstack_ess_scaling_group":                     95,
		"alibabacloudstack_ess_scaling_lifecycle_hook":            96,
		"alibabacloudstack_ess_scaling_rule":                      97,
		"alibabacloudstack_ess_scalinggroup_vserver_groups":       98,
		"alibabacloudstack_ess_scheduled_task":                    99,
		"alibabacloudstack_express_connect_physical_connection":   100,
		"alibabacloudstack_express_connect_virtual_border_router": 101,
		"alibabacloudstack_forward_entry":                         102,
		"alibabacloudstack_gpdb_account":                          103,
		"alibabacloudstack_gpdb_connection":                       104,
		"alibabacloudstack_gpdb_instance":                         105,
		"alibabacloudstack_hbase_instance":                        106,
		"alibabacloudstack_image":                                 107,
		"alibabacloudstack_image_copy":                            108,
		"alibabacloudstack_image_export":                          109,
		"alibabacloudstack_image_import":                          110,
		"alibabacloudstack_image_share_permission":                111,
		"alibabacloudstack_instance":                              112,
		"alibabacloudstack_key_pair":                              113,
		"alibabacloudstack_key_pair_attachment":                   114,
		"alibabacloudstack_kms_alias":                             115,
		"alibabacloudstack_kms_ciphertext":                        116,
		"alibabacloudstack_kms_key":                               117,
		"alibabacloudstack_kms_secret":                            118,
		"alibabacloudstack_kvstore_account":                       119,
		"alibabacloudstack_kvstore_backup_policy":                 120,
		"alibabacloudstack_kvstore_connection":                    121,
		"alibabacloudstack_kvstore_instance":                      122,
		"alibabacloudstack_launch_template":                       123,
		"alibabacloudstack_log_machine_group":                     124,
		"alibabacloudstack_log_project":                           125,
		"alibabacloudstack_log_store":                             126,
		"alibabacloudstack_log_store_index":                       127,
		"alibabacloudstack_logtail_attachment":                    128,
		"alibabacloudstack_logtail_config":                        129,
		"alibabacloudstack_maxcompute_cu":                         130,
		"alibabacloudstack_maxcompute_project":                    131,
		"alibabacloudstack_maxcompute_user":                       132,
		"alibabacloudstack_mongodb_instance":                      133,
		"alibabacloudstack_mongodb_sharding_instance":             134,
		"alibabacloudstack_nas_access_group":                      135,
		"alibabacloudstack_nas_access_rule":                       136,
		"alibabacloudstack_nas_file_system":                       137,
		"alibabacloudstack_nas_mount_target":                      138,
		"alibabacloudstack_nat_gateway":                           139,
		"alibabacloudstack_network_acl":                           140,
		"alibabacloudstack_network_acl_attachment":                141,
		"alibabacloudstack_network_acl_entries":                   142,
		"alibabacloudstack_network_interface":                     143,
		"alibabacloudstack_network_interface_attachment":          144,
		"alibabacloudstack_ons_group":                             145,
		"alibabacloudstack_ons_instance":                          146,
		"alibabacloudstack_ons_topic":                             147,
		"alibabacloudstack_oos_execution":                         148,
		"alibabacloudstack_oos_template":                          149,
		"alibabacloudstack_oss_bucket":                            150,
		"alibabacloudstack_oss_bucket_object":                     151,
		"alibabacloudstack_ots_instance":                          152,
		"alibabacloudstack_ots_instance_attachment":               153,
		"alibabacloudstack_ots_table":                             154,
		"alibabacloudstack_quick_bi_user":                         155,
		"alibabacloudstack_quick_bi_user_group":                   156,
		"alibabacloudstack_quick_bi_workspace":                    157,
		"alibabacloudstack_ram_role_attachment":                   158,
		"alibabacloudstack_reserved_instance":                     159,
		"alibabacloudstack_ros_stack":                             160,
		"alibabacloudstack_ros_template":                          161,
		"alibabacloudstack_route_entry":                           162,
		"alibabacloudstack_route_table":                           163,
		"alibabacloudstack_route_table_attachment":                164,
		"alibabacloudstack_router_interface":                      165,
		"alibabacloudstack_router_interface_connection":           166,
		"alibabacloudstack_security_group":                        167,
		"alibabacloudstack_security_group_rule":                   168,
		"alibabacloudstack_slb":                                   169,
		"alibabacloudstack_slb_acl":                               170,
		"alibabacloudstack_slb_backend_server":                    171,
		"alibabacloudstack_slb_ca_certificate":                    172,
		"alibabacloudstack_slb_domain_extension":                  173,
		"alibabacloudstack_slb_listener":                          174,
		"alibabacloudstack_slb_master_slave_server_group":         175,
		"alibabacloudstack_slb_rule":                              176,
		"alibabacloudstack_slb_server_certificate":                177,
		"alibabacloudstack_slb_server_group":                      178,
		"alibabacloudstack_snapshot":                              179,
		"alibabacloudstack_snapshot_policy":                       180,
		"alibabacloudstack_snat":                                  181,
		"alibabacloudstack_vpc":                                   182,
		"alibabacloudstack_vpc_ipv6_egress_rule":                  183,
		"alibabacloudstack_vpc_ipv6_gateway":                      184,
		"alibabacloudstack_vpc_ipv6_internet_bandwidth":           185,
		"alibabacloudstack_vpn_connection":                        186,
		"alibabacloudstack_vpn_customer_gateway":                  187,
		"alibabacloudstack_vpn_gateway":                           188,
		"alibabacloudstack_vpn_route_entry":                       189,
		"alibabacloudstack_vswitch":                               190,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
