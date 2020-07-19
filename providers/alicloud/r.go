package alicloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "alicloud_actiontrail",
			Category:         "Actiontrail",
			ShortDescription: `Provides Alibaba Cloud ActionTrail Resource`,
			Description:      ``,
			Keywords: []string{
				"actiontrail",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) The name of the trail to be created, which must be unique for an account.`,
				},
				resource.Attribute{
					Name:        "event_rw",
					Description: `(Optional) Indicates whether the event is a read or a write event. Valid values: Read, Write, and All. Default value: Write.`,
				},
				resource.Attribute{
					Name:        "oss_bucket_name",
					Description: `(Required) The OSS bucket to which the trail delivers logs. Ensure that this is an existing OSS bucket.`,
				},
				resource.Attribute{
					Name:        "role_name",
					Description: `(Required) The RAM role in ActionTrail permitted by the user.`,
				},
				resource.Attribute{
					Name:        "oss_key_prefix",
					Description: `(Optional) The prefix of the specified OSS bucket name. This parameter can be left empty.`,
				},
				resource.Attribute{
					Name:        "sls_project_arn",
					Description: `(Optional) The unique ARN of the Log Service project.`,
				},
				resource.Attribute{
					Name:        "sls_write_role_arn",
					Description: `(Optional) The unique ARN of the Log Service role. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The action trail id. The value is same as its name. ## Import Action trail can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_actiontrail.foo abc12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The action trail id. The value is same as its name. ## Import Action trail can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_actiontrail.foo abc12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_adb_account",
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
					Description: `The current account resource ID. Composed of instance ID and account name with format ` + "`" + `<instance_id>:<name>` + "`" + `. ## Import ADB account can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_adb_account.example "am-12345:tf_account" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The current account resource ID. Composed of instance ID and account name with format ` + "`" + `<instance_id>:<name>` + "`" + `. ## Import ADB account can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_adb_account.example "am-12345:tf_account" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_adb_backup_policy",
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
					Description: `Cluster backup retention days, Fixed for 7 days, not modified. ## Import ADB backup policy can be imported using the id or cluster id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_adb_backup_policy.example "am-12345678" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The current backup policy resource ID. It is same as 'db_cluster_id'.`,
				},
				resource.Attribute{
					Name:        "backup_retention_period",
					Description: `Cluster backup retention days, Fixed for 7 days, not modified. ## Import ADB backup policy can be imported using the id or cluster id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_adb_backup_policy.example "am-12345678" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_adb_cluster",
			Category:         "AnalyticDB for MySQL (ADB)",
			ShortDescription: `Provides a ADB cluster resource.`,
			Description:      ``,
			Keywords: []string{
				"analyticdb",
				"for",
				"mysql",
				"adb",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "db_cluster_version",
					Description: `(Optional, ForceNew) Cluster version. Value options: ` + "`" + `3.0` + "`" + `, Default to ` + "`" + `3.0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "db_cluster_category",
					Description: `(Required, ForceNew) Cluster category. Value options: ` + "`" + `Basic` + "`" + `, ` + "`" + `Cluster` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "db_node_class",
					Description: `(Required) The db_node_class of cluster node.`,
				},
				resource.Attribute{
					Name:        "db_node_count",
					Description: `(Required) The db_node_count of cluster node.`,
				},
				resource.Attribute{
					Name:        "db_node_storage",
					Description: `(Required) The db_node_storage of cluster node.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Optional) The Zone to launch the DB cluster.`,
				},
				resource.Attribute{
					Name:        "pay_type",
					Description: `(Optional, ForceNew) Valid values are ` + "`" + `PrePaid` + "`" + `, ` + "`" + `PostPaid` + "`" + `, Default to ` + "`" + `PostPaid` + "`" + `. Currently, the resource can not supports change pay type.`,
				},
				resource.Attribute{
					Name:        "renewal_status",
					Description: `(Optional) Valid values are ` + "`" + `AutoRenewal` + "`" + `, ` + "`" + `Normal` + "`" + `, ` + "`" + `NotRenewal` + "`" + `, Default to ` + "`" + `NotRenewal` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "auto_renew_period",
					Description: `(Optional) Auto-renewal period of an cluster, in the unit of the month. It is valid when pay_type is ` + "`" + `PrePaid` + "`" + `. Valid value:1, 2, 3, 6, 12, 24, 36, Default to 1.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) The duration that you will buy DB cluster (in month). It is valid when pay_type is ` + "`" + `PrePaid` + "`" + `. Valid values: [1~9], 12, 24, 36. Default to 1.`,
				},
				resource.Attribute{
					Name:        "security_ips",
					Description: `(Optional) List of IP addresses allowed to access all databases of an cluster. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Required, ForceNew) The virtual switch ID to launch DB instances in one VPC.`,
				},
				resource.Attribute{
					Name:        "maintain_time",
					Description: `(Optional) Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of cluster.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string. - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ADB cluster ID. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 50 mins) Used when creating the adb cluster (until it reaches the initial ` + "`" + `Running` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 72 hours) Used when updating the adb cluster (until it reaches the initial ` + "`" + `Running` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 10 mins) Used when terminating the adb cluster. ## Import ADB cluster can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_adb_cluster.example am-abc12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ADB cluster ID. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 50 mins) Used when creating the adb cluster (until it reaches the initial ` + "`" + `Running` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 72 hours) Used when updating the adb cluster (until it reaches the initial ` + "`" + `Running` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 10 mins) Used when terminating the adb cluster. ## Import ADB cluster can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_adb_cluster.example am-abc12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_adb_connection",
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
					Description: `(ForceNew) Prefix of an Internet connection string. It must be checked for uniqueness. It may consist of lowercase letters, numbers, and underlines, and must start with a letter and have no more than 30 characters. Default to <db_cluster_id> + 'tf'. ## Attributes Reference The following attributes are exported:`,
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
					Description: `The ip address of connection string. ## Import ADB connection can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_adb_connection.example am-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `The ip address of connection string. ## Import ADB connection can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_adb_connection.example am-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_alidns_domain_group",
			Category:         "DNS",
			ShortDescription: `Provides a Alidns Domain Group resource.`,
			Description:      ``,
			Keywords: []string{
				"dns",
				"alidns",
				"domain",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group_name",
					Description: `(Required) Name of the domain group.`,
				},
				resource.Attribute{
					Name:        "lang",
					Description: `(Optional) User language. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `This ID of this domain group resource. ## Import Alidns domain group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_alidns_domain_group.example 0932eb3ddee7499085c4d13d45`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `This ID of this domain group resource. ## Import Alidns domain group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_alidns_domain_group.example 0932eb3ddee7499085c4d13d45`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_alidns_record",
			Category:         "DNS",
			ShortDescription: `Provides a Alidns Domain Record resource.`,
			Description:      ``,
			Keywords: []string{
				"dns",
				"alidns",
				"record",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain_name",
					Description: `(Required, ForceNew) Name of the domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix ` + "`" + `.sh` + "`" + ` and ` + "`" + `.tel` + "`" + ` are not supported.`,
				},
				resource.Attribute{
					Name:        "rr",
					Description: `(Required) Host record for the domain record. This host_record can have at most 253 characters, and each part split with ` + "`" + `.` + "`" + ` can have at most 63 characters, and must contain only alphanumeric characters or hyphens, such as ` + "`" + `-` + "`" + `, ` + "`" + `.` + "`" + `, ` + "`" + ``,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of domain record. Valid values: ` + "`" + `A` + "`" + `,` + "`" + `NS` + "`" + `,` + "`" + `MX` + "`" + `,` + "`" + `TXT` + "`" + `,` + "`" + `CNAME` + "`" + `,` + "`" + `SRV` + "`" + `,` + "`" + `AAAA` + "`" + `,` + "`" + `CAA` + "`" + `, ` + "`" + `REDIRECT_URL` + "`" + ` and ` + "`" + `FORWORD_URL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of domain record, When the ` + "`" + `type` + "`" + ` is ` + "`" + `MX` + "`" + `,` + "`" + `NS` + "`" + `,` + "`" + `CNAME` + "`" + `,` + "`" + `SRV` + "`" + `, the server will treat the ` + "`" + `value` + "`" + ` as a fully qualified domain name, so it's no need to add a ` + "`" + `.` + "`" + ` at the end.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The effective time of domain record. Its scope depends on the edition of the cloud resolution. Free is ` + "`" + `[600, 86400]` + "`" + `, Basic is ` + "`" + `[120, 86400]` + "`" + `, Standard is ` + "`" + `[60, 86400]` + "`" + `, Ultimate is ` + "`" + `[10, 86400]` + "`" + `, Exclusive is ` + "`" + `[1, 86400]` + "`" + `. Default value is ` + "`" + `600` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) The priority of domain record. Valid values: ` + "`" + `[1-10]` + "`" + `. When the ` + "`" + `type` + "`" + ` is ` + "`" + `MX` + "`" + `, this parameter is required.`,
				},
				resource.Attribute{
					Name:        "line",
					Description: `(Optional) The resolution line of domain record. Valid values: ` + "`" + `default` + "`" + `, ` + "`" + `telecom` + "`" + `, ` + "`" + `unicom` + "`" + `, ` + "`" + `mobile` + "`" + `, ` + "`" + `oversea` + "`" + `, ` + "`" + `edu` + "`" + `, ` + "`" + `drpeng` + "`" + `, ` + "`" + `btvn` + "`" + `. When the ` + "`" + `type` + "`" + ` is ` + "`" + `FORWORD_URL` + "`" + `, this parameter must be ` + "`" + `default` + "`" + `. Default value is ` + "`" + `default` + "`" + `. For checking all resolution lines enumeration please visit [Alibaba Cloud DNS doc](https://www.alibabacloud.com/help/doc-detail/34339.htm) or using alicloud_dns_resolution_lines in data source to get the value.`,
				},
				resource.Attribute{
					Name:        "lang",
					Description: `(Optional) User language.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `(Optional) The remark of the domain record.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The status of the domain record. Valid values: ` + "`" + `ENABLE` + "`" + `,` + "`" + `DISABLE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user_client_ip",
					Description: `(Optional) The IP address of the client. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of Domain Record. ## Import Alidns Domain Record can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_alidns_record.example abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of Domain Record. ## Import Alidns Domain Record can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_alidns_record.example abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_alikafka_consumer_group",
			Category:         "Alikafka",
			ShortDescription: `Provides a Alicloud Alikafka Consumer Group resource.`,
			Description:      ``,
			Keywords: []string{
				"alikafka",
				"consumer",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, ForceNew) ID of the ALIKAFKA Instance that owns the groups.`,
				},
				resource.Attribute{
					Name:        "consumer_id",
					Description: `(Required, ForceNew) ID of the consumer group. The length cannot exceed 64 characters.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Available in v1.63.0+) A mapping of tags to assign to the resource. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `key` + "`" + ` of the resource supplied above. The value is formulated as ` + "`" + `<instance_id>:<consumer_id>` + "`" + `. ## Import ALIKAFKA GROUP can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_alikafka_consumer_group.group alikafka_post-cn-123455abc:consumerId ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `key` + "`" + ` of the resource supplied above. The value is formulated as ` + "`" + `<instance_id>:<consumer_id>` + "`" + `. ## Import ALIKAFKA GROUP can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_alikafka_consumer_group.group alikafka_post-cn-123455abc:consumerId ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_alikafka_instance",
			Category:         "Alikafka",
			ShortDescription: `Provides a Alicloud ALIKAFKA Instance resource.`,
			Description:      ``,
			Keywords: []string{
				"alikafka",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of your Kafka instance. The length should between 3 and 64 characters. If not set, will use instance id as instance name.`,
				},
				resource.Attribute{
					Name:        "topic_quota",
					Description: `(Required) The max num of topic can be create of the instance. When modify this value, it only adjust to a greater value.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `(Required, ForceNew) The disk type of the instance. 0: efficient cloud disk , 1: SSD.`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `(Required) The disk size of the instance. When modify this value, it only support adjust to a greater value.`,
				},
				resource.Attribute{
					Name:        "deploy_type",
					Description: `(Required, ForceNew) The deploy type of the instance. Currently only support two deploy type, 4: eip/vpc instance, 5: vpc instance.`,
				},
				resource.Attribute{
					Name:        "io_max",
					Description: `(Required) The max value of io of the instance. When modify this value, it only support adjust to a greater value.`,
				},
				resource.Attribute{
					Name:        "eip_max",
					Description: `(Optional) The max bandwidth of the instance. When modify this value, it only support adjust to a greater value.`,
				},
				resource.Attribute{
					Name:        "paid_type",
					Description: `(Optional) The paid type of the instance. Support two type, "PrePaid": pre paid type instance, "PostPaid": post paid type instance. Default is PostPaid. When modify this value, it only support adjust from post pay to pre pay.`,
				},
				resource.Attribute{
					Name:        "spec_type",
					Description: `(Optional) The spec type of the instance. Support two type, "normal": normal version instance, "professional": professional version instance. Default is normal. When modify this value, it only support adjust from normal to professional. Note only pre paid type instance support professional specific type.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Required, ForceNew) The ID of attaching vswitch to instance.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Available in v1.63.0+) A mapping of tags to assign to the resource. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `key` + "`" + ` of the resource supplied above, also call instance id.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of attaching VPC to instance.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `The Zone to launch the kafka instance. ## Import ALIKAFKA TOPIC can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_alikafka_instance.instance alikafka_post-cn-123455abc ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `key` + "`" + ` of the resource supplied above, also call instance id.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of attaching VPC to instance.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `The Zone to launch the kafka instance. ## Import ALIKAFKA TOPIC can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_alikafka_instance.instance alikafka_post-cn-123455abc ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_alikafka_sasl_acl",
			Category:         "Alikafka",
			ShortDescription: `Provides a Alicloud Alikafka Sasl Acl resource.`,
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
					Description: `The host of the acl. ## Import ALIKAFKA GROUP can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_alikafka_sasl_acl.acl alikafka_post-cn-123455abc:username:Topic:test-topic:LITERAL:Write ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `key` + "`" + ` of the resource supplied above. The value is formulated as ` + "`" + `<instance_id>:<username>:<acl_resource_type>:<acl_resource_name>:<acl_resource_pattern_type>:<acl_operation_type>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `The host of the acl. ## Import ALIKAFKA GROUP can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_alikafka_sasl_acl.acl alikafka_post-cn-123455abc:username:Topic:test-topic:LITERAL:Write ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_alikafka_sasl_user",
			Category:         "Alikafka",
			ShortDescription: `Provides a Alicloud Alikafka Sasl User resource.`,
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
					Description: `(Optional, MapString) An KMS encryption context used to decrypt ` + "`" + `kms_encrypted_password` + "`" + ` before creating or updating a user with ` + "`" + `kms_encrypted_password` + "`" + `. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when ` + "`" + `kms_encrypted_password` + "`" + ` is set. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `key` + "`" + ` of the resource supplied above. The value is formulated as ` + "`" + `<instance_id>:<username>` + "`" + `. ## Import ALIKAFKA GROUP can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_alikafka_sasl_user.user alikafka_post-cn-123455abc:username ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `key` + "`" + ` of the resource supplied above. The value is formulated as ` + "`" + `<instance_id>:<username>` + "`" + `. ## Import ALIKAFKA GROUP can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_alikafka_sasl_user.user alikafka_post-cn-123455abc:username ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_alikafka_topic",
			Category:         "Alikafka",
			ShortDescription: `Provides a Alicloud ALIKAFKA Topic resource.`,
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
					Description: `The ` + "`" + `key` + "`" + ` of the resource supplied above. The value is formulated as ` + "`" + `<instance_id>:<topic>` + "`" + `. ## Import ALIKAFKA TOPIC can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_alikafka_topic.topic alikafka_post-cn-123455abc:topicName ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `key` + "`" + ` of the resource supplied above. The value is formulated as ` + "`" + `<instance_id>:<topic>` + "`" + `. ## Import ALIKAFKA TOPIC can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_alikafka_topic.topic alikafka_post-cn-123455abc:topicName ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_api_gateway_api",
			Category:         "API Gateway",
			ShortDescription: `Provides a Alicloud Api Gateway Api Resource.`,
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
					Name:        "http_service_config",
					Description: `(Optional, Type: list) http_service_config defines the config when service_type selected 'HTTP'.`,
				},
				resource.Attribute{
					Name:        "http_vpc_service_config",
					Description: `(Optional, Type: list) http_vpc_service_config defines the config when service_type selected 'HTTP-VPC'.`,
				},
				resource.Attribute{
					Name:        "fc_service_config",
					Description: `(Optional, Type: list) fc_service_config defines the config when service_type selected 'FunctionCompute'.`,
				},
				resource.Attribute{
					Name:        "mock_service_config",
					Description: `(Optional, Type: list) http_service_config defines the config when service_type selected 'MOCK'.`,
				},
				resource.Attribute{
					Name:        "request_parameters",
					Description: `(Required, Type: list) request_parameters defines the request parameters of the api.`,
				},
				resource.Attribute{
					Name:        "constant_parameters",
					Description: `(Required, Type: list) constant_parameters defines the constant parameters of the api.`,
				},
				resource.Attribute{
					Name:        "system_parameters",
					Description: `(Required, Type: list) system_parameters defines the system parameters of the api.`,
				},
				resource.Attribute{
					Name:        "stage_names",
					Description: `(Optional, Type: list) Stages that the api need to be deployed. Valid value: RELEASE | PRE | TEST. ### Block request_config The request_config mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The protocol of api which supports values of 'HTTP','HTTPS' or 'HTTP,HTTPS'`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Required) The method of the api, including 'GET','POST','PUT' and etc..`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The request path of the api.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Required) The mode of the parameters between request parameters and service parameters, which support the values of 'MAPPING' and 'PASSTHROUGH'`,
				},
				resource.Attribute{
					Name:        "body_format",
					Description: `(Optional) The body format of the api, which support the values of 'STREAM' and 'FORM' ### Block http_service_config The http_service_config mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) The address of backend service.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The path of backend service.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Required) The http method of backend service.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Required) Backend service time-out time; unit: millisecond. ### Block http_vpc_service_config The http_vpc_service_config mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of vpc instance.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The path of backend service.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Required) The http method of backend service.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Required) Backend service time-out time; unit: millisecond. ### Block fc_vpc_service_config The fc_service_config mapping supports the following:`,
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
					Description: `(Required) Backend service time-out time; unit: millisecond. ### Block mock_service_config The mock_service_config mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "result",
					Description: `(Required) The result of the mock service. ### Block request_parameters The request_parameters mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Request's parameter name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Parameter type which supports values of 'STRING','INT','BOOLEAN','LONG',"FLOAT" and "DOUBLE"`,
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
					Description: `(Required) Backend service's parameter name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of parameter.`,
				},
				resource.Attribute{
					Name:        "default_value",
					Description: `(Optional) The default value of the parameter. ### Block constant_parameters The constant_parameters mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Constant parameter name.`,
				},
				resource.Attribute{
					Name:        "in",
					Description: `(Required) Constant parameter location; values: 'HEAD' and 'QUERY'.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Constant parameter value.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of Constant parameter. ### Block system_parameters The system_parameters mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) System parameter name which supports values including in [system parameter list](https://www.alibabacloud.com/help/doc-detail/43677.html)`,
				},
				resource.Attribute{
					Name:        "in",
					Description: `(Required) System parameter location; values: 'HEAD' and 'QUERY'.`,
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
			Type:             "alicloud_api_gateway_app",
			Category:         "API Gateway",
			ShortDescription: `Provides a Alicloud Api Gateway App Resource.`,
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
			Type:             "alicloud_api_gateway_app_attachment",
			Category:         "API Gateway",
			ShortDescription: `Provides a Alicloud Api Gateway App Attachment Resource.`,
			Description:      ``,
			Keywords: []string{
				"api",
				"gateway",
				"app",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "api_id",
					Description: `(Required，ForceNew) The api_id that app apply to access.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required，ForceNew) The group that the api belongs to.`,
				},
				resource.Attribute{
					Name:        "app_id",
					Description: `(Required，ForceNew) The app that apply to the authorization.`,
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
			Type:             "alicloud_api_gateway_group",
			Category:         "API Gateway",
			ShortDescription: `Provides a Alicloud Api Gateway Group Resource.`,
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
					Description: `(Available in 1.69.0+) Second-level VPC domain name automatically assigned to the API group. ## Import Api gateway group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_api_gateway_group.example "ab2351f2ce904edaa8d92a0510832b91" ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Available in 1.69.0+) Second-level VPC domain name automatically assigned to the API group. ## Import Api gateway group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_api_gateway_group.example "ab2351f2ce904edaa8d92a0510832b91" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_api_gateway_vpc_access",
			Category:         "API Gateway",
			ShortDescription: `Provides a Alicloud Api Gateway vpc authorization Resource.`,
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
					Description: `The ID of the vpc authorization of api gateway. ## Import Api gateway app can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_api_gateway_vpc_access.example "APiGatewayVpc:vpc-aswcj19ajsz:i-ajdjfsdlf:8080" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the vpc authorization of api gateway. ## Import Api gateway app can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_api_gateway_vpc_access.example "APiGatewayVpc:vpc-aswcj19ajsz:i-ajdjfsdlf:8080" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_auto_provisioning_group",
			Category:         "ECS",
			ShortDescription: `Provides a ECS Auto Provisioning group resource.`,
			Description:      ``,
			Keywords: []string{
				"ecs",
				"auto",
				"provisioning",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "launch_template_id",
					Description: `(Required) The ID of the instance launch template associated with the auto provisioning group.`,
				},
				resource.Attribute{
					Name:        "total_target_capacity",
					Description: `(Required) The total target capacity of the auto provisioning group. The target capacity consists of the following three parts:PayAsYouGoTargetCapacity,SpotTargetCapacity and the supplemental capacity besides PayAsYouGoTargetCapacity and SpotTargetCapacity.`,
				},
				resource.Attribute{
					Name:        "auto_Provisioning_group_name",
					Description: `(Optional) The name of the auto provisioning group to be created. It must be 2 to 128 characters in length. It must start with a letter but cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-)`,
				},
				resource.Attribute{
					Name:        "auto_Provisioning_group_type",
					Description: `(Optional) The type of the auto provisioning group. Valid values:` + "`" + `request` + "`" + ` and ` + "`" + `maintain` + "`" + `,Default value: ` + "`" + `maintain` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "spot_allocation_strategy",
					Description: `(Optional) The scale-out policy for preemptible instances. Valid values:` + "`" + `lowest-price` + "`" + ` and ` + "`" + `diversified` + "`" + `,Default value: ` + "`" + `lowest-price` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "spot_target_capacity",
					Description: `(Optional) The target capacity of preemptible instances in the auto provisioning group.`,
				},
				resource.Attribute{
					Name:        "spot_instance_interruption_behavior",
					Description: `(Optional) The default behavior after preemptible instances are shut down. Value values: ` + "`" + `stop` + "`" + ` and ` + "`" + `terminate` + "`" + `,Default value: ` + "`" + `stop` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "spot_instance_pools_to_use_count",
					Description: `(Optional) This parameter takes effect when the ` + "`" + `SpotAllocationStrategy` + "`" + ` parameter is set to ` + "`" + `lowest-price` + "`" + `. The auto provisioning group selects instance types of the lowest cost to create instances.`,
				},
				resource.Attribute{
					Name:        "pay_as_you_go_allocation_strategy",
					Description: `(Optional) The scale-out policy for pay-as-you-go instances. Valid values: ` + "`" + `lowest-price` + "`" + ` and ` + "`" + `prioritized` + "`" + `,Default value: ` + "`" + `lowest-price` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "pay_as_you_go_target_capacity",
					Description: `(Optional) The target capacity of pay-as-you-go instances in the auto provisioning group.`,
				},
				resource.Attribute{
					Name:        "default_target_capacity_type",
					Description: `(Optional) The type of supplemental instances. When the total value of ` + "`" + `PayAsYouGoTargetCapacity` + "`" + ` and ` + "`" + `SpotTargetCapacity` + "`" + ` is smaller than the value of TotalTargetCapacity, the auto provisioning group will create instances of the specified type to meet the capacity requirements. Valid values:` + "`" + `PayAsYouGo` + "`" + `: Pay-as-you-go instances; ` + "`" + `Spot` + "`" + `: Preemptible instances, Default value: ` + "`" + `Spot` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "launch_template_version",
					Description: `(Optional) The version of the instance launch template associated with the auto provisioning group.`,
				},
				resource.Attribute{
					Name:        "excess_capacity_termination_policy",
					Description: `(Optional) The shutdown policy for excess preemptible instances followed when the capacity of the auto provisioning group exceeds the target capacity. Valid values: ` + "`" + `no-termination` + "`" + ` and ` + "`" + `termination` + "`" + `,Default value: ` + "`" + `no-termination` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "terminate_instances_with_expiration",
					Description: `(Optional) The shutdown policy for preemptible instances when the auto provisioning group expires. Valid values: ` + "`" + `false` + "`" + ` and ` + "`" + `true` + "`" + `, default value: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "terminate_instances",
					Description: `(Optional) Specifies whether to release instances of the auto provisioning group. Valid values:` + "`" + `false` + "`" + ` and ` + "`" + `true` + "`" + `, default value: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the auto provisioning group.`,
				},
				resource.Attribute{
					Name:        "max_spot_price",
					Description: `(Optional) The global maximum price for preemptible instances in the auto provisioning group. If both the ` + "`" + `MaxSpotPrice` + "`" + ` and ` + "`" + `LaunchTemplateConfig.N.MaxPrice` + "`" + ` parameters are specified, the maximum price is the lower value of the two.`,
				},
				resource.Attribute{
					Name:        "valid_from",
					Description: `(Optional) The time when the auto provisioning group is started. The period of time between this point in time and the point in time specified by the ` + "`" + `valid_until` + "`" + ` parameter is the effective time period of the auto provisioning group.By default, an auto provisioning group is immediately started after creation.`,
				},
				resource.Attribute{
					Name:        "valid_until",
					Description: `(Optional) The time when the auto provisioning group expires. The period of time between this point in time and the point in time specified by the ` + "`" + `valid_from` + "`" + ` parameter is the effective time period of the auto provisioning group.By default, an auto provisioning group never expires.`,
				},
				resource.Attribute{
					Name:        "launch_template_config",
					Description: `(Optional) DataDisk mappings to attach to ecs instance. See [Block config](#block-config) below for details. ## Block config The config mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Optional) The instance type of the Nth extended configurations of the launch template.`,
				},
				resource.Attribute{
					Name:        "max_price",
					Description: `(Required) The maximum price of the instance type specified in the Nth extended configurations of the launch template.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Required) The ID of the VSwitch in the Nth extended configurations of the launch template.`,
				},
				resource.Attribute{
					Name:        "weighted_capacity",
					Description: `(Optional) The weight of the instance type specified in the Nth extended configurations of the launch template.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) The priority of the instance type specified in the Nth extended configurations of the launch template. A value of 0 indicates the highest priority. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the auto provisioning group ## Import ECS auto provisioning group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_auto_provisioning_group.example asg-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the auto provisioning group ## Import ECS auto provisioning group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_auto_provisioning_group.example asg-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cas_certificate",
			Category:         "SSL Certificates",
			ShortDescription: `Provides a CAS Certificate resource.`,
			Description:      ``,
			Keywords: []string{
				"ssl",
				"certificates",
				"cas",
				"certificate",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForcesNew) Name of the Certificate. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix ` + "`" + `.sh` + "`" + ` and ` + "`" + `.tel` + "`" + ` are not supported.`,
				},
				resource.Attribute{
					Name:        "cert",
					Description: `(Required, ForcesNew) Cert of the Certificate in which the Certificate will add.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required, ForcesNew) Key of the Certificate in which the Certificate will add. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The cert id.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The cert id.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cassandra_cluster",
			Category:         "Cassandra",
			ShortDescription: `Provides a Cassandra cluster resource.`,
			Description:      ``,
			Keywords: []string{
				"cassandra",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Required) Cassandra cluster name. Length must be 2~128 characters long. Only Chinese characters, English letters, numbers, period ` + "`" + `.` + "`" + `, underline ` + "`" + `_` + "`" + `, or dash ` + "`" + `-` + "`" + ` are permitted.`,
				},
				resource.Attribute{
					Name:        "data_center_name",
					Description: `(Required) Cassandra dataCenter-1 name. Length must be 2~128 characters long. Only Chinese characters, English letters, numbers, period ` + "`" + `.` + "`" + `, underline ` + "`" + `_` + "`" + `, or dash ` + "`" + `-` + "`" + ` are permitted.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Optional, ForceNew) The Zone to launch the Cassandra cluster. If vswitch_id is not empty, this zone_id can be "" or consistent.`,
				},
				resource.Attribute{
					Name:        "major_version",
					Description: `(Required, ForceNew) Cassandra major version. Now only support version ` + "`" + `3.11` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Required) Instance specification. See [Instance specifications](https://help.aliyun.com/document_detail/157445.html). Or you can call describeInstanceType api.`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `(Optional) User-defined Cassandra dataCenter-1 one node's storage space.Unit: GB. Value range: - Custom storage space; value range: [160, 2000]. - 80-GB increments.`,
				},
				resource.Attribute{
					Name:        "pay_type",
					Description: `(Optional, ForceNew) The pay type of Cassandra dataCenter-1. Valid values are ` + "`" + `Subscription` + "`" + `, ` + "`" + `PayAsYouGo` + "`" + `,System default to ` + "`" + `PayAsYouGo` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "auto_renew",
					Description: `(Optional, ForceNew) Auto renew of dataCenter-1,` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `. System default to ` + "`" + `false` + "`" + `, valid when pay_type = PrePaid.`,
				},
				resource.Attribute{
					Name:        "auto_renew_period",
					Description: `(Optional, ForceNew) Period of dataCenter-1 auto renew, if auto renew is ` + "`" + `true` + "`" + `, one of ` + "`" + `1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 60` + "`" + `, valid when pay_type = Subscription. Unit: month.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Optional, ForceNew) The vswitch_id of dataCenter-1, can not empty.`,
				},
				resource.Attribute{
					Name:        "maintain_start_time",
					Description: `(Optional) The start time of the operation and maintenance time period of the cluster, in the format of HH:mmZ (UTC time).`,
				},
				resource.Attribute{
					Name:        "maintain_end_time",
					Description: `(Optional) The end time of the operation and maintenance time period of the cluster, in the format of HH:mmZ (UTC time).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "ip_white",
					Description: `(Optional) Set the instance's IP whitelist in VPC network.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Optional) A list of security group ids to associate with. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Cassandra. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 30 mins) Used when creating the Cassandra cluster (until it reaches the initial ` + "`" + `Running` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 30 mins) Used when updating the Cassandra cluster (until it reaches the initial ` + "`" + `Running` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 5 mins) Used when terminating the Cassandra cluster. ## Import Cassandra cluster can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cassandra_cluster.example cds-wz9sr400dd7xxxxx ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Cassandra. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 30 mins) Used when creating the Cassandra cluster (until it reaches the initial ` + "`" + `Running` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 30 mins) Used when updating the Cassandra cluster (until it reaches the initial ` + "`" + `Running` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 5 mins) Used when terminating the Cassandra cluster. ## Import Cassandra cluster can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cassandra_cluster.example cds-wz9sr400dd7xxxxx ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cassandra_data_center",
			Category:         "Cassandra",
			ShortDescription: `Provides a Cassandra dataCenter resource.`,
			Description:      ``,
			Keywords: []string{
				"cassandra",
				"data",
				"center",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) Cassandra cluster id of dataCenter-2 belongs to.`,
				},
				resource.Attribute{
					Name:        "data_center_name",
					Description: `(Required) Cassandra dataCenter-2 name. Length must be 2~128 characters long. Only Chinese characters, English letters, numbers, period ` + "`" + `.` + "`" + `, underline ` + "`" + `_` + "`" + `, or dash ` + "`" + `-` + "`" + ` are permitted.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Optional, ForceNew) The Zone to launch the Cassandra dataCenter-2. If vswitch_id is not empty, this zone_id can be "" or consistent.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Required) Instance specification. See [Instance specifications](https://help.aliyun.com/document_detail/157445.html). Or you can call describeInstanceType api.`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `(Optional) User-defined Cassandra dataCenter one core node's storage space.Unit: GB. Value range: - Custom storage space; value range: [160, 2000]. - 80-GB increments.`,
				},
				resource.Attribute{
					Name:        "pay_type",
					Description: `(Optional, ForceNew) The pay type of Cassandra dataCenter-2. Valid values are ` + "`" + `Subscription` + "`" + `, ` + "`" + `PayAsYouGo` + "`" + `. System default to ` + "`" + `PayAsYouGo` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "auto_renew",
					Description: `(Optional, ForceNew) Auto renew of dataCenter-2,` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `. System default to ` + "`" + `false` + "`" + `, valid when pay_type = Subscription.`,
				},
				resource.Attribute{
					Name:        "auto_renew_period",
					Description: `(Optional, ForceNew) Period of dataCenter-2 auto renew, if auto renew is ` + "`" + `true` + "`" + `, one of ` + "`" + `1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 60` + "`" + `, valid when pay_type = Subscription. Unit: month.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Optional, ForceNew) The vswitch_id of dataCenter-2, mast different of vswitch_id(dc-1), can not empty. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Cassandra. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 30 mins) Used when creating the Cassandra dataCenter (until it reaches the initial ` + "`" + `Running` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 30 mins) Used when updating the Cassandra dataCenter (until it reaches the initial ` + "`" + `Running` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 5 mins) Used when terminating the Cassandra dataCenter. ## Import If you need full function, please import Cassandra cluster first. Cassandra dataCenter can be imported using the dcId:clusterId, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cassandra_data_center.dc_2 cn-shenxxxx-x:cds-wz933ryoaurxxxxx ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Cassandra. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 30 mins) Used when creating the Cassandra dataCenter (until it reaches the initial ` + "`" + `Running` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 30 mins) Used when updating the Cassandra dataCenter (until it reaches the initial ` + "`" + `Running` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 5 mins) Used when terminating the Cassandra dataCenter. ## Import If you need full function, please import Cassandra cluster first. Cassandra dataCenter can be imported using the dcId:clusterId, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cassandra_data_center.dc_2 cn-shenxxxx-x:cds-wz933ryoaurxxxxx ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cdn_domain",
			Category:         "CDN",
			ShortDescription: `Provides a CDN Accelerated Domain resource.`,
			Description:      ``,
			Keywords: []string{
				"cdn",
				"domain",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain_name",
					Description: `(Required) Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix ` + "`" + `.sh` + "`" + ` and ` + "`" + `.tel` + "`" + ` are not supported.`,
				},
				resource.Attribute{
					Name:        "cdn_type",
					Description: `(Required) Cdn type of the accelerated domain. Valid values are ` + "`" + `web` + "`" + `, ` + "`" + `download` + "`" + `, ` + "`" + `video` + "`" + `, ` + "`" + `liveStream` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `(Optional) Source type of the accelerated domain. Valid values are ` + "`" + `ipaddr` + "`" + `, ` + "`" + `domain` + "`" + `, ` + "`" + `oss` + "`" + `. You must set this parameter when ` + "`" + `cdn_type` + "`" + ` value is not ` + "`" + `liveStream` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "source_port",
					Description: `(Optional) Source port of the accelerated domain. Valid values are ` + "`" + `80` + "`" + ` and ` + "`" + `443` + "`" + `. Default value is ` + "`" + `80` + "`" + `. You must use ` + "`" + `80` + "`" + ` when the ` + "`" + `source_type` + "`" + ` is ` + "`" + `oss` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sources",
					Description: `(Optional, Type: list) Sources of the accelerated domain. It's a list of domain names or IP address and consists of at most 20 items. You must set this parameter when ` + "`" + `cdn_type` + "`" + ` value is not ` + "`" + `liveStream` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) Scope of the accelerated domain. Valid values are ` + "`" + `domestic` + "`" + `, ` + "`" + `overseas` + "`" + `, ` + "`" + `global` + "`" + `. Default value is ` + "`" + `domestic` + "`" + `. This parameter's setting is valid Only for the international users and domestic L3 and above users . #### Domain config The config supports the following:`,
				},
				resource.Attribute{
					Name:        "optimize_enable",
					Description: `(Optional) Page Optimize config of the accelerated domain. Valid values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. Default value is ` + "`" + `off` + "`" + `. It can effectively remove the page redundant content, reduce the file size and improve the speed of distribution when this parameter value is ` + "`" + `on` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "page_compress_enable",
					Description: `(Optional) Page Compress config of the accelerated domain. Valid values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. Default value is ` + "`" + `off` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "range_enable",
					Description: `(Optional) Range Source config of the accelerated domain. Valid values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. Default value is ` + "`" + `off` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "video_seek_enable",
					Description: `(Optional) Video Seek config of the accelerated domain. Valid values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. Default value is ` + "`" + `off` + "`" + `. ### Block parameter_filter_config ` + "`" + `parameter_filter_config` + "`" + ` - (Optional, Type: set) Parameter filter config of the accelerated domain. It's a set and consists of at most one item.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional) This parameter indicates whether or not the ` + "`" + `parameter_filter_config` + "`" + ` is enable. Valid values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. Default value is ` + "`" + `off` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "hash_key_args",
					Description: `(Optional, Type: list) Reserved parameters of ` + "`" + `parameter_filter_config` + "`" + `. It's a list of string and consists of at most 10 items. ### Block page_404_config ` + "`" + `page_404_config` + "`" + ` - (Optional, Type: set) Error Page config of the accelerated domain. It's a set and consists of at most one item.`,
				},
				resource.Attribute{
					Name:        "page_type",
					Description: `(Optional) Page type of the error page. Valid values are ` + "`" + `default` + "`" + `, ` + "`" + `charity` + "`" + `, ` + "`" + `other` + "`" + `. Default value is ` + "`" + `default` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "custom_page_url",
					Description: `(Optional) Custom page url of the error page. It must be the full path under the accelerated domain name. It's value must be ` + "`" + `http://promotion.alicdn.com/help/oss/error.html` + "`" + ` when ` + "`" + `page_type` + "`" + ` value is ` + "`" + `charity` + "`" + ` and It can not be set when ` + "`" + `page_type` + "`" + ` value is ` + "`" + `default` + "`" + `. ### Block refer_config ` + "`" + `refer_config` + "`" + ` - (Optional, Type: set) Refer anti-theft chain config of the accelerated domain. It's a set and consists of at most 1 item.`,
				},
				resource.Attribute{
					Name:        "refer_type",
					Description: `(Optional) Refer type of the refer config. Valid values are ` + "`" + `block` + "`" + ` and ` + "`" + `allow` + "`" + `. Default value is ` + "`" + `block` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "refer_list",
					Description: `(Required, Type: list) A list of domain names of the refer config.`,
				},
				resource.Attribute{
					Name:        "allow_empty",
					Description: `(Optional) This parameter indicates whether or not to allow empty refer access. Valid values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. Default value is ` + "`" + `on` + "`" + `. ### Block auth_config ` + "`" + `auth_config` + "`" + ` - (Optional, Type: set) Auth config of the accelerated domain. It's a set and consist of at most 1 item.`,
				},
				resource.Attribute{
					Name:        "auth_type",
					Description: `(Optional) Auth type of the auth config. Valid values are ` + "`" + `no_auth` + "`" + `, ` + "`" + `type_a` + "`" + `, ` + "`" + `type_b` + "`" + ` and ` + "`" + `type_c` + "`" + `. Default value is ` + "`" + `no_auth` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "master_key",
					Description: `(Optional) Master authentication key of the auth config. This parameter can have a string of 6 to 32 characters and must contain only alphanumeric characters.`,
				},
				resource.Attribute{
					Name:        "slave_key",
					Description: `(Optional) Slave authentication key of the auth config. This parameter can have a string of 6 to 32 characters and must contain only alphanumeric characters.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional, Type: int) Authentication cache time of the auth config. Default value is ` + "`" + `1800` + "`" + `. It's value is valid only when the ` + "`" + `auth_type` + "`" + ` is ` + "`" + `type_b` + "`" + ` or ` + "`" + `type_c` + "`" + `. ### Block certificate_config ` + "`" + `certificate_config` + "`" + ` - (Optional, Type: set) Certificate config of the accelerated domain. It's a set and consist of at most 1 item.`,
				},
				resource.Attribute{
					Name:        "server_certificate_status",
					Description: `(Optional) This parameter indicates whether or not enable https. Valid values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. Default value is ` + "`" + `on` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "server_certificate",
					Description: `(Optional) The SSL server certificate string. This is required if ` + "`" + `server_certificate_status` + "`" + ` is ` + "`" + `on` + "`" + ``,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Optional) The SSL private key. This is required if ` + "`" + `server_certificate_status` + "`" + ` is ` + "`" + `on` + "`" + ` ### Block http_header_config ` + "`" + `http_header_config` + "`" + ` - (Optional, Type: set) Http header config of the accelerated domain. It's a set and consist of at most 8 items. The ` + "`" + `header_key` + "`" + ` for each item can not be repeated.`,
				},
				resource.Attribute{
					Name:        "header_key",
					Description: `(Required) Header key of the http header. Valid values are ` + "`" + `Content-Type` + "`" + `, ` + "`" + `Cache-Control` + "`" + `, ` + "`" + `Content-Disposition` + "`" + `, ` + "`" + `Content-Language` + "`" + `，` + "`" + `Expires` + "`" + `, ` + "`" + `Access-Control-Allow-Origin` + "`" + `, ` + "`" + `Access-Control-Allow-Methods` + "`" + ` and ` + "`" + `Access-Control-Max-Age` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "header_value",
					Description: `(Required) Header value of the http header. ### Block cache_config ` + "`" + `cache_config` + "`" + ` - (Optional, Type: set) Cache config of the accelerated domain. It's a set and each item's ` + "`" + `cache_content` + "`" + ` can not be repeated.`,
				},
				resource.Attribute{
					Name:        "cache_type",
					Description: `(Required) Cache type of the cache config. Valid values are ` + "`" + `suffix` + "`" + ` and ` + "`" + `path` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cache_content",
					Description: `(Required) Cache content of the cache config. It's value is a path string when the ` + "`" + `cache_type` + "`" + ` is ` + "`" + `path` + "`" + `. When the ` + "`" + `cache_type` + "`" + ` is ` + "`" + `suffix` + "`" + `, it's value is a string which contains multiple file suffixes separated by commas.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Required, Type: int) Cache time of the cache config.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional, Type: int) Weight of the cache config. This parameter's value is between 1 and 99. Default value is ` + "`" + `1` + "`" + `. The higher the value, the higher the priority. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `The accelerated domain name.`,
				},
				resource.Attribute{
					Name:        "sources",
					Description: `The accelerated domain sources.`,
				},
				resource.Attribute{
					Name:        "cdn_type",
					Description: `The cdn type of the accelerated domain.`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `The source type ot the accelerated domain.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `The accelerated domain scope.`,
				},
				resource.Attribute{
					Name:        "optimize_enable",
					Description: `The page optimize config of the accelerated domain.`,
				},
				resource.Attribute{
					Name:        "page_compress_enable",
					Description: `The page compress config of the accelerated domain.`,
				},
				resource.Attribute{
					Name:        "range_enable",
					Description: `The range source config of the accelerated domain.`,
				},
				resource.Attribute{
					Name:        "video_seek_enable",
					Description: `The video seek config of the accelerated domain.`,
				},
				resource.Attribute{
					Name:        "parameter_filter_config",
					Description: `The parameter filter config of the accelerated domain.`,
				},
				resource.Attribute{
					Name:        "page_404_config",
					Description: `The error page config of the accelerated domain.`,
				},
				resource.Attribute{
					Name:        "refer_config",
					Description: `The refer config of the accelerated domain.`,
				},
				resource.Attribute{
					Name:        "auth_config",
					Description: `The auth config of the accelerated domain.`,
				},
				resource.Attribute{
					Name:        "http_header_config",
					Description: `The http header configs of the accelerated domain.`,
				},
				resource.Attribute{
					Name:        "cache_config",
					Description: `The cache configs of the accelerated domain.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "domain_name",
					Description: `The accelerated domain name.`,
				},
				resource.Attribute{
					Name:        "sources",
					Description: `The accelerated domain sources.`,
				},
				resource.Attribute{
					Name:        "cdn_type",
					Description: `The cdn type of the accelerated domain.`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `The source type ot the accelerated domain.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `The accelerated domain scope.`,
				},
				resource.Attribute{
					Name:        "optimize_enable",
					Description: `The page optimize config of the accelerated domain.`,
				},
				resource.Attribute{
					Name:        "page_compress_enable",
					Description: `The page compress config of the accelerated domain.`,
				},
				resource.Attribute{
					Name:        "range_enable",
					Description: `The range source config of the accelerated domain.`,
				},
				resource.Attribute{
					Name:        "video_seek_enable",
					Description: `The video seek config of the accelerated domain.`,
				},
				resource.Attribute{
					Name:        "parameter_filter_config",
					Description: `The parameter filter config of the accelerated domain.`,
				},
				resource.Attribute{
					Name:        "page_404_config",
					Description: `The error page config of the accelerated domain.`,
				},
				resource.Attribute{
					Name:        "refer_config",
					Description: `The refer config of the accelerated domain.`,
				},
				resource.Attribute{
					Name:        "auth_config",
					Description: `The auth config of the accelerated domain.`,
				},
				resource.Attribute{
					Name:        "http_header_config",
					Description: `The http header configs of the accelerated domain.`,
				},
				resource.Attribute{
					Name:        "cache_config",
					Description: `The cache configs of the accelerated domain.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cdn_domain_config",
			Category:         "CDN",
			ShortDescription: `Provides a Alicloud Cdn domain config Resource.`,
			Description:      ``,
			Keywords: []string{
				"cdn",
				"domain",
				"config",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain_name",
					Description: `(Required, ForceNew) Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix ` + "`" + `.sh` + "`" + ` and ` + "`" + `.tel` + "`" + ` are not supported.`,
				},
				resource.Attribute{
					Name:        "function_name",
					Description: `(Required, ForceNew) The name of the domain config.`,
				},
				resource.Attribute{
					Name:        "function_args",
					Description: `(Required, ForceNew, Type: list) The args of the domain config. ### Block function_args The ` + "`" + `function_args` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "arg_name",
					Description: `(Required) The name of arg.`,
				},
				resource.Attribute{
					Name:        "arg_value",
					Description: `(Required) The value of arg. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the domain config. The value is formate as ` + "`" + `<domain_name>:<function_name>` + "`" + `. ## Import CDN domain config can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import alicloud_cdn_domain_config.example cdn:config-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the domain config. The value is formate as ` + "`" + `<domain_name>:<function_name>` + "`" + `. ## Import CDN domain config can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import alicloud_cdn_domain_config.example cdn:config-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cdn_domain_new",
			Category:         "CDN",
			ShortDescription: `Provides a Alicloud Cdn Domain New Resource.`,
			Description:      ``,
			Keywords: []string{
				"cdn",
				"domain",
				"new",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain_name",
					Description: `(Required) Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix ` + "`" + `.sh` + "`" + ` and ` + "`" + `.tel` + "`" + ` are not supported.`,
				},
				resource.Attribute{
					Name:        "cdn_type",
					Description: `(Required, ForceNew) Cdn type of the accelerated domain. Valid values are ` + "`" + `web` + "`" + `, ` + "`" + `download` + "`" + `, ` + "`" + `video` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) Scope of the accelerated domain. Valid values are ` + "`" + `domestic` + "`" + `, ` + "`" + `overseas` + "`" + `, ` + "`" + `global` + "`" + `. Default value is ` + "`" + `domestic` + "`" + `. This parameter's setting is valid Only for the international users and domestic L3 and above users .`,
				},
				resource.Attribute{
					Name:        "sources",
					Description: `(Optional, Type: list) The source address list of the accelerated domain. Defaults to null. See Block Sources.`,
				},
				resource.Attribute{
					Name:        "certificate_config",
					Description: `(Optional, Type: list, Available in 1.52.0+) Certificate config of the accelerated domain. It's a list and consist of at most 1 item.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional, Available in v1.67.0+) Resource group ID. ### Block sources The ` + "`" + `sources` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Required) The adress of source. Valid values can be ip or doaminName. Each item's ` + "`" + `content` + "`" + ` can not be repeated.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of the source. Valid values are ` + "`" + `ipaddr` + "`" + `, ` + "`" + `domain` + "`" + ` and ` + "`" + `oss` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional, Type: int) The port of source. Valid values are ` + "`" + `443` + "`" + ` and ` + "`" + `80` + "`" + `. Default value is ` + "`" + `80` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional, Type: int) Priority of the source. Valid values are ` + "`" + `0` + "`" + ` and ` + "`" + `100` + "`" + `. Default value is ` + "`" + `20` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional, Type: int) Weight of the source. Valid values are from ` + "`" + `0` + "`" + ` to ` + "`" + `100` + "`" + `. Default value is ` + "`" + `10` + "`" + `, but if type is ` + "`" + `ipaddr` + "`" + `, the value can only be ` + "`" + `10` + "`" + `. ### Block certificate_config The ` + "`" + `certificate_config` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "server_certificate_status",
					Description: `(Optional) This parameter indicates whether or not enable https. Valid values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. Default value is ` + "`" + `on` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "server_certificate",
					Description: `(Optional) The SSL server certificate string. This is required if ` + "`" + `server_certificate_status` + "`" + ` is ` + "`" + `on` + "`" + ``,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Optional) The SSL private key. This is required if ` + "`" + `server_certificate_status` + "`" + ` is ` + "`" + `on` + "`" + ``,
				},
				resource.Attribute{
					Name:        "force_set",
					Description: `(Optional) Set ` + "`" + `1` + "`" + ` to ignore the repeated verification for certificate name, and cover the information of the origin certificate (with the same name). Set ` + "`" + `0` + "`" + ` to work the verification.`,
				},
				resource.Attribute{
					Name:        "cert_name",
					Description: `(Optional) The SSL certificate name.`,
				},
				resource.Attribute{
					Name:        "cert_type",
					Description: `(Optional) The SSL certificate type, can be "upload", "cas" and "free".`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Available in v1.55.2+) A mapping of tags to assign to the resource. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The cdn domain id. The value is same as the domain name.`,
				},
				resource.Attribute{
					Name:        "cname",
					Description: `(Available in v1.90.0+) The CNAME of the CDN domain. ## Import CDN domain can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import alicloud_cdn_domain_new.example xxxx.com ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The cdn domain id. The value is same as the domain name.`,
				},
				resource.Attribute{
					Name:        "cname",
					Description: `(Available in v1.90.0+) The CNAME of the CDN domain. ## Import CDN domain can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import alicloud_cdn_domain_new.example xxxx.com ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cen_bandwidth_limit",
			Category:         "Cloud Enterprise Network (CEN)",
			ShortDescription: `Provides a Alicloud CEN cross-regional interconnection bandwidth configuration resource.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"enterprise",
				"network",
				"cen",
				"bandwidth",
				"limit",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, ForceNew) The ID of the CEN.`,
				},
				resource.Attribute{
					Name:        "region_ids",
					Description: `(Required, ForceNew) List of the two regions to interconnect. Must be two different regions.`,
				},
				resource.Attribute{
					Name:        "bandwidth_limit",
					Description: `(Required) The bandwidth configured for the interconnected regions communication. ->`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 10 mins) Used when activating the cen bandwidth limit when necessary during update - when changing bandwidth limit.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 10 mins) Used when terminating the cen bandwidth limit. ## Attributes Reference The following attributes are exported: - ` + "`" + `id` + "`" + ` - ID of the resource, formatted as ` + "`" + `<instance_id>:<region_id_1>:<region_id_2>` + "`" + `. ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cen_bandwidth_package",
			Category:         "Cloud Enterprise Network (CEN)",
			ShortDescription: `Provides a Alicloud CEN bandwidth package resource.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"enterprise",
				"network",
				"cen",
				"bandwidth",
				"package",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Required) The bandwidth in Mbps of the bandwidth package. Cannot be less than 2Mbps.`,
				},
				resource.Attribute{
					Name:        "geographic_region_ids",
					Description: `(Required) List of the two areas to connect. Valid value: China | North-America | Asia-Pacific | Europe | Middle-East | Australia.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the bandwidth package. Defaults to null.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the bandwidth package. Default to null.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `(Optional) The billing method. Valid value: PostPaid | PrePaid. Default to PostPaid. If set to PrePaid, the bandwidth package can't be deleted before expired time.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) The purchase period in month. Valid value: 1, 2, 3, 6, 12. Default to 1. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the bandwidth package.`,
				},
				resource.Attribute{
					Name:        "expired_time",
					Description: `The time of the bandwidth package to expire.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the bandwidth, including "InUse" and "Idle". ## Import CEN bandwidth package can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cen_bandwidth_package.example cenbwp-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the bandwidth package.`,
				},
				resource.Attribute{
					Name:        "expired_time",
					Description: `The time of the bandwidth package to expire.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the bandwidth, including "InUse" and "Idle". ## Import CEN bandwidth package can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cen_bandwidth_package.example cenbwp-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cen_bandwidth_package_attachment",
			Category:         "Cloud Enterprise Network (CEN)",
			ShortDescription: `Provides a Alicloud CEN bandwidth package attachment resource.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"enterprise",
				"network",
				"cen",
				"bandwidth",
				"package",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, ForceNew) The ID of the CEN.`,
				},
				resource.Attribute{
					Name:        "bandwidth_package_id",
					Description: `(Required, ForceNew) The ID of the bandwidth package. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource, the same as bandwidth_package_id. ## Import CEN bandwidth package attachment resource can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $terraform import alicloud_cen_bandwidth_package_attachment.example bwp-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource, the same as bandwidth_package_id. ## Import CEN bandwidth package attachment resource can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $terraform import alicloud_cen_bandwidth_package_attachment.example bwp-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cen_flowlog",
			Category:         "Cloud Enterprise Network (CEN)",
			ShortDescription: `Provides a Alicloud CEN manage route entried resource.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"enterprise",
				"network",
				"cen",
				"flowlog",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cen_id",
					Description: `(Required, ForceNew) The ID of the CEN Instance.`,
				},
				resource.Attribute{
					Name:        "project_name",
					Description: `(Required, ForceNew) The name of the SLS project.`,
				},
				resource.Attribute{
					Name:        "log_store_name",
					Description: `(Required, ForceNew) The name of the log store which is in the ` + "`" + `project_name` + "`" + ` SLS project.`,
				},
				resource.Attribute{
					Name:        "flow_log_name",
					Description: `(Optional) The name of flowlog.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of flowlog.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The status of flowlog. Valid values: ["Active", "Inactive"]. Default to "Active". ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the flowlog. ## Import CEN flowlog can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cen_flowlog.default flowlog-tig1xxxxxx ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the flowlog. ## Import CEN flowlog can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cen_flowlog.default flowlog-tig1xxxxxx ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cen_instance",
			Category:         "Cloud Enterprise Network (CEN)",
			ShortDescription: `Provides a Alicloud CEN instance resource.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"enterprise",
				"network",
				"cen",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the CEN instance. Defaults to null. The name must be 2 to 128 characters in length and can contain letters, numbers, periods (.), underscores (_), and hyphens (-). The name must start with a letter, but cannot start with http:// or https://.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the CEN instance. Defaults to null. The description must be 2 to 256 characters in length. It must start with a letter, and cannot start with http:// or https://.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Available in v1.80.0+) A mapping of tags to assign to the resource. ### Timeouts ->`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 6 mins) Used when creating the cen instance (until it reaches the initial ` + "`" + `Active` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 6 mins) Used when terminating the cen instance. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the CEN instance.`,
				},
				resource.Attribute{
					Name:        "protection_level",
					Description: `(Available in 1.76.0+) Indicates the allowed level of CIDR block overlapping.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The Cen Instance current status. ## Import CEN instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cen_instance.example cen-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the CEN instance.`,
				},
				resource.Attribute{
					Name:        "protection_level",
					Description: `(Available in 1.76.0+) Indicates the allowed level of CIDR block overlapping.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The Cen Instance current status. ## Import CEN instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cen_instance.example cen-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cen_instance_attachment",
			Category:         "Cloud Enterprise Network (CEN)",
			ShortDescription: `Provides a Alicloud CEN child instance attachment resource.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"enterprise",
				"network",
				"cen",
				"instance",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, ForceNew) The ID of the CEN.`,
				},
				resource.Attribute{
					Name:        "child_instance_id",
					Description: `(Required, ForceNew) The ID of the child instance to attach.`,
				},
				resource.Attribute{
					Name:        "child_instance_region_id",
					Description: `(Required, ForceNew) The region ID of the child instance to attach.`,
				},
				resource.Attribute{
					Name:        "child_instance_owner_id",
					Description: `(Optional, Available in 1.42.0+) The uid of the child instance. Only used when attach a child instance of other account. ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cen_instance_grant",
			Category:         "Cloud Enterprise Network (CEN)",
			ShortDescription: `Provides a Alicloud CEN child instance grant resource.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"enterprise",
				"network",
				"cen",
				"instance",
				"grant",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cen_id",
					Description: `(Required) The ID of the CEN.`,
				},
				resource.Attribute{
					Name:        "child_instance_id",
					Description: `(Required) The ID of the child instance to grant.`,
				},
				resource.Attribute{
					Name:        "cen_owner_id",
					Description: `(Required) The owner UID of the CEN which the child instance granted to. ## Attributes Reference The following attributes are exported: - ` + "`" + `id` + "`" + ` - ID of the resource, formatted as ` + "`" + `<cen_id>:<child_instance_id>:<cen_owner_id>` + "`" + `. ## Import CEN instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cen_instance_grant.example cen-abc123456:vpc-abc123456:uid123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cen_private_zone",
			Category:         "Cloud Enterprise Network (CEN)",
			ShortDescription: `Provides a Alicloud CEN private zone resource.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"enterprise",
				"network",
				"cen",
				"private",
				"zone",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cen_id",
					Description: `(Required, ForceNew) The ID of the CEN instance.`,
				},
				resource.Attribute{
					Name:        "access_region_id",
					Description: `(Required, ForceNew) The access region. The access region is the region of the cloud resource that accesses the PrivateZone service through CEN.`,
				},
				resource.Attribute{
					Name:        "host_region_id",
					Description: `(Required, ForceNew) The service region. The service region is the target region of the PrivateZone service to be accessed through CEN.`,
				},
				resource.Attribute{
					Name:        "host_vpc_id",
					Description: `(Required, ForceNew) The VPC that belongs to the service region. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource, formatted as ` + "`" + `<cen_id>:<access_region_id>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the PrivateZone service. Valid values: ["Creating", "Active", "Deleting"]. ## Import CEN Private Zone can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cen_private_zone.example cen-abc123456:cn-hangzhou ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource, formatted as ` + "`" + `<cen_id>:<access_region_id>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the PrivateZone service. Valid values: ["Creating", "Active", "Deleting"]. ## Import CEN Private Zone can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cen_private_zone.example cen-abc123456:cn-hangzhou ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cen_route_entry",
			Category:         "Cloud Enterprise Network (CEN)",
			ShortDescription: `Provides a Alicloud CEN manage route entried resource.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"enterprise",
				"network",
				"cen",
				"route",
				"entry",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, ForceNew) The ID of the CEN.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `(Required, ForceNew) The route table of the attached VBR or VPC.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Required, ForceNew) The destination CIDR block of the route entry to publish. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource, formatted as ` + "`" + `<instance_id>:<route_table_id>:<cidr_block>` + "`" + `. ## Import CEN instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cen_route_entry.example cen-abc123456:vtb-abc123:192.168.0.0/24 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource, formatted as ` + "`" + `<instance_id>:<route_table_id>:<cidr_block>` + "`" + `. ## Import CEN instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cen_route_entry.example cen-abc123456:vtb-abc123:192.168.0.0/24 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cen_route_map",
			Category:         "Cloud Enterprise Network (CEN)",
			ShortDescription: `Provides a Alicloud CEN manage route map resource.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"enterprise",
				"network",
				"cen",
				"route",
				"map",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cen_id",
					Description: `(Required, ForceNew) The ID of the CEN instance.`,
				},
				resource.Attribute{
					Name:        "cen_region_id",
					Description: `(Required) The ID of the region to which the CEN instance belongs.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Required) The priority of the route map. Value range: 1 to 100. A lower value indicates a higher priority.`,
				},
				resource.Attribute{
					Name:        "transmit_direction",
					Description: `(Required, ForceNew) The direction in which the route map is applied. Valid values: ["RegionIn", "RegionOut"].`,
				},
				resource.Attribute{
					Name:        "map_result",
					Description: `(Required) The action that is performed to a route if the route matches all the match conditions. Valid values: ["Permit", "Deny"].`,
				},
				resource.Attribute{
					Name:        "next_priority",
					Description: `(Optional) The priority of the next route map that is associated with the current route map. Value range: 1 to 100.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the route map.`,
				},
				resource.Attribute{
					Name:        "source_region_ids",
					Description: `(Optional) A match statement that indicates the list of IDs of the source regions. You can enter a maximum of 32 region IDs.`,
				},
				resource.Attribute{
					Name:        "source_instance_ids",
					Description: `(Optional) A match statement that indicates the list of IDs of the source instances.`,
				},
				resource.Attribute{
					Name:        "source_instance_ids_reverse_match",
					Description: `(Optional) Indicates whether to enable the reverse match method for the SourceInstanceIds match condition. Valid values: ["false", "true"]. Default to "false".`,
				},
				resource.Attribute{
					Name:        "destination_instance_ids",
					Description: `(Optional) A match statement that indicates the list of IDs of the destination instances.`,
				},
				resource.Attribute{
					Name:        "destination_instance_ids_reverse_match",
					Description: `(Optional) Indicates whether to enable the reverse match method for the DestinationInstanceIds match condition. Valid values: ["false", "true"]. Default to "false".`,
				},
				resource.Attribute{
					Name:        "source_route_table_ids",
					Description: `(Optional) A match statement that indicates the list of IDs of the source route tables. You can enter a maximum of 32 route table IDs.`,
				},
				resource.Attribute{
					Name:        "destination_route_table_ids",
					Description: `(Optional) A match statement that indicates the list of IDs of the destination route tables. You can enter a maximum of 32 route table IDs.`,
				},
				resource.Attribute{
					Name:        "source_child_instance_types",
					Description: `(Optional) A match statement that indicates the list of source instance types. Valid values: ["VPC", "VBR", "CCN"].`,
				},
				resource.Attribute{
					Name:        "destination_child_instance_types",
					Description: `(Optional) A match statement that indicates the list of destination instance types. Valid values: ["VPC", "VBR", "CCN"].`,
				},
				resource.Attribute{
					Name:        "destination_cidr_blocks",
					Description: `(Optional) A match statement that indicates the prefix list. The prefix is in the CIDR format. You can enter a maximum of 32 CIDR blocks.`,
				},
				resource.Attribute{
					Name:        "cidr_match_mode",
					Description: `(Optional) A match statement. It indicates the mode in which the prefix attribute is matched. Valid values: ["Include", "Complete"].`,
				},
				resource.Attribute{
					Name:        "route_types",
					Description: `(Optional) A match statement that indicates the list of route types. Valid values: ["System", "Custom", "BGP"].`,
				},
				resource.Attribute{
					Name:        "match_asns",
					Description: `(Optional) A match statement that indicates the AS path list. The AS path is a well-known mandatory attribute, which describes the numbers of the ASs that a BGP route passes through during transmission.`,
				},
				resource.Attribute{
					Name:        "as_path_match_mode",
					Description: `(Optional) A match statement. It indicates the mode in which the AS path attribute is matched. Valid values: ["Include", "Complete"].`,
				},
				resource.Attribute{
					Name:        "match_community_set",
					Description: `(Optional) A match statement that indicates the community set. The format of each community is nn:nn, which ranges from 1 to 65535. You can enter a maximum of 32 communities. Communities must comply with RFC 1997. Large communities (RFC 8092) are not supported.`,
				},
				resource.Attribute{
					Name:        "community_match_mode",
					Description: `(Optional) A match statement. It indicates the mode in which the community attribute is matched. Valid values: ["Include", "Complete"].`,
				},
				resource.Attribute{
					Name:        "community_operate_mode",
					Description: `(Optional) An action statement. It indicates the mode in which the community attribute is operated. Valid values: ["Additive", "Replace"].`,
				},
				resource.Attribute{
					Name:        "operate_community_set",
					Description: `(Optional) An action statement that operates the community attribute. The format of each community is nn:nn, which ranges from 1 to 65535. You can enter a maximum of 32 communities. Communities must comply with RFC 1997. Large communities (RFC 8092) are not supported.`,
				},
				resource.Attribute{
					Name:        "preference",
					Description: `(Optional) An action statement that modifies the priority of the route. Value range: 1 to 100. The default priority of a route is 50. A lower value indicates a higher preference.`,
				},
				resource.Attribute{
					Name:        "prepend_as_path",
					Description: `(Optional) An action statement that indicates an AS path is prepended when the regional gateway receives or advertises a route. ## Attributes Reference The RouteMapId attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the RouteMap.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Computed) The status of route map. Valid values: ["Creating", "Active", "Deleting"]. ## Import CEN RouteMap can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cen_route_map.default cenrmap-tig1xxxxxx ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the RouteMap.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Computed) The status of route map. Valid values: ["Creating", "Active", "Deleting"]. ## Import CEN RouteMap can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cen_route_map.default cenrmap-tig1xxxxxx ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cen_vbr_health_check",
			Category:         "Cloud Enterprise Network (CEN)",
			ShortDescription: `Provides a Alicloud CEN VBR HealthCheck resource.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"enterprise",
				"network",
				"cen",
				"vbr",
				"health",
				"check",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cen_id",
					Description: `(Required, ForceNew) The ID of the CEN instance.`,
				},
				resource.Attribute{
					Name:        "health_check_interval",
					Description: `(Optional, Default) Specifies the interval at which the health check sends continuous detection packets. Default value: 2. Value range: 2 to 3.`,
				},
				resource.Attribute{
					Name:        "health_check_source_ip",
					Description: `(Optional) The source IP address of health checks.`,
				},
				resource.Attribute{
					Name:        "health_check_target_ip",
					Description: `(Required) The destination IP address of health checks.`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `(Optional, Default) Specifies the number of probe messages sent by the health check. Default value: 8. Value range: 3 to 8.`,
				},
				resource.Attribute{
					Name:        "vbr_instance_id",
					Description: `(Required, ForceNew) The ID of the VBR.`,
				},
				resource.Attribute{
					Name:        "vbr_instance_owner_id",
					Description: `(Optional) The ID of the account to which the VBR belongs.`,
				},
				resource.Attribute{
					Name:        "vbr_instance_region_id",
					Description: `(Required, ForceNew) The ID of the region to which the VBR belongs. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource, formatted as ` + "`" + `<vbr_instance_id>:<vbr_instance_region_id>` + "`" + `. ## Import CEN VBR HealthCheck can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cen_vbr_health_check.example vbr-xxxxx:cn-hangzhou ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource, formatted as ` + "`" + `<vbr_instance_id>:<vbr_instance_region_id>` + "`" + `. ## Import CEN VBR HealthCheck can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cen_vbr_health_check.example vbr-xxxxx:cn-hangzhou ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cloud_connect_network",
			Category:         "Cloud Connect Network",
			ShortDescription: `Provides a Alicloud Cloud Connect Network resource.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"connect",
				"network",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the CCN instance. The name can contain 2 to 128 characters including a-z, A-Z, 0-9, periods, underlines, and hyphens. The name must start with an English letter, but cannot start with http:// or https://.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the CCN instance. The description can contain 2 to 256 characters. The description must start with English letters, but cannot start with http:// or https://.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Optional) The CidrBlock of the CCN instance. Defaults to null.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `(Required) Created by default. If the client does not have ccn in the binding, it will create a ccn for the user to replace. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The CcnId of the CCN instance. For example "ccn-xxx". ## Import The cloud connect network instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cloud_connect_network.example ccn-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The CcnId of the CCN instance. For example "ccn-xxx". ## Import The cloud connect network instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cloud_connect_network.example ccn-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cloud_connect_network_attachment",
			Category:         "Cloud Connect Network",
			ShortDescription: `Provides a Alicloud Cloud Connect Network Attachment resource.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"connect",
				"network",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ccn_id",
					Description: `(Required,ForceNew) The ID of the CCN instance.`,
				},
				resource.Attribute{
					Name:        "sag_id",
					Description: `(Required,ForceNew) The ID of the Smart Access Gateway instance. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Cloud Connect Network Attachment Id and formates as ` + "`" + `<ccn_id>:<sag_id>` + "`" + `. ## Import The Cloud Connect Network Attachment can be imported using the instance_id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cloud_connect_network_attachment.example ccn-abc123456:sag-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Cloud Connect Network Attachment Id and formates as ` + "`" + `<ccn_id>:<sag_id>` + "`" + `. ## Import The Cloud Connect Network Attachment can be imported using the instance_id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cloud_connect_network_attachment.example ccn-abc123456:sag-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cloud_connect_network_grant",
			Category:         "Cloud Connect Network",
			ShortDescription: `Provides a Alicloud Cloud Connect Network Grant resource.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"connect",
				"network",
				"grant",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ccn_id",
					Description: `(Required,ForceNew) The ID of the CCN instance.`,
				},
				resource.Attribute{
					Name:        "cen_id",
					Description: `(Required,ForceNew) The ID of the CEN instance.`,
				},
				resource.Attribute{
					Name:        "cen_uid",
					Description: `(Required,ForceNew) The ID of the account to which the CEN instance belongs. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Cloud Connect Network grant Id and formates as ` + "`" + `<ccn_id>:<cen_id>` + "`" + `. ## Import The Cloud Connect Network Grant can be imported using the instance_id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cloud_connect_network_grant.example ccn-abc123456:cen-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Cloud Connect Network grant Id and formates as ` + "`" + `<ccn_id>:<cen_id>` + "`" + `. ## Import The Cloud Connect Network Grant can be imported using the instance_id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cloud_connect_network_grant.example ccn-abc123456:cen-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cms_alarm",
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
					Description: `Index query cycle, which must be consistent with that defined for metrics. Default to 300, in seconds.`,
				},
				resource.Attribute{
					Name:        "statistics",
					Description: `Statistical method. It must be consistent with that defined for metrics. Valid values: ["Average", "Minimum", "Maximum"]. Default to "Average".`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `Alarm comparison operator. Valid values: ["<=", "<", ">", ">=", "==", "!="]. Default to "==".`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `(Required) Alarm threshold value, which must be a numeric value currently.`,
				},
				resource.Attribute{
					Name:        "triggered_count",
					Description: `Number of consecutive times it has been detected that the values exceed the threshold. Default to 3.`,
				},
				resource.Attribute{
					Name:        "contact_groups",
					Description: `(Required) List contact groups of the alarm rule, which must have been created on the console.`,
				},
				resource.Attribute{
					Name:        "effective_interval",
					Description: `(Available in 1.50.0+) The interval of effecting alarm rule. It foramt as "hh:mm-hh:mm", like "0:00-4:00". Default to "00:00-23:59".`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `It has been deprecated from provider version 1.50.0 and 'effective_interval' instead.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `It has been deprecated from provider version 1.50.0 and 'effective_interval' instead.`,
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
					Name:        "id",
					Description: `The ID of the alarm rule.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current alarm rule status. ## Import Alarm rule can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cms_alarm.alarm abc12345 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the alarm rule.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current alarm rule status. ## Import Alarm rule can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cms_alarm.alarm abc12345 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cms_site_monitor",
			Category:         "Cloud Monitor",
			ShortDescription: `Provides a resource to build a site monitor rule for cloud monitor.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"monitor",
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
					Description: `The detection points in a JSON array. For example, ` + "`" + `[{"city":"546","isp":"465"},{"city":"572","isp":"465"},{"city":"738","isp":"465"}]` + "`" + ` indicates the detection points in Beijing, Hangzhou, and Qingdao respectively. You can call the [DescribeSiteMonitorISPCityList](https://www.alibabacloud.com/help/en/doc-detail/115045.htm?spm=a2c63.p38356.b99.238.5fec36962UlFG6) operation to query detection point information. If this parameter is not specified, three detection points will be chosen randomly for monitoring.`,
				},
				resource.Attribute{
					Name:        "options_json",
					Description: `The extended options of the protocol of the site monitoring task. The options vary according to the protocol. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the site monitor rule. ## Import Alarm rule can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cms_site_monitor.alarm abc12345 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the site monitor rule. ## Import Alarm rule can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cms_site_monitor.alarm abc12345 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_common_bandwidth_package",
			Category:         "VPC",
			ShortDescription: `Provides a Alicloud Common Bandwidth Package resource.`,
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
					Name:        "internet_charge_type",
					Description: `(Optional, ForceNew) The billing method of the common bandwidth package. Valid values are "PayByBandwidth" and "PayBy95" and "PayByTraffic". "PayBy95" is pay by classic 95th percentile pricing. International Account doesn't supports "PayByBandwidth" and "PayBy95". Default to "PayByTraffic".`,
				},
				resource.Attribute{
					Name:        "ratio",
					Description: `(Optional, ForceNew Available in 1.55.3+) Ratio of the common bandwidth package. It is valid when ` + "`" + `internet_charge_type` + "`" + ` is ` + "`" + `PayBy95` + "`" + `. Default to 100. Valid values: [10-100].`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the common bandwidth package.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the common bandwidth package instance.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(ForceNew, Available in 1.58.0+) The Id of resource group which the common bandwidth package belongs.`,
				},
				resource.Attribute{
					Name:        "isp",
					Description: `(Optional, Available in 1.90.1+) The type of the Internet Service Provider. Default to ` + "`" + `BGP` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the common bandwidth package instance id. ## Import The common bandwidth package can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_common_bandwidth_package.foo cbwp-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the common bandwidth package instance id. ## Import The common bandwidth package can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_common_bandwidth_package.foo cbwp-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_common_bandwidth_package_attachment",
			Category:         "VPC",
			ShortDescription: `Provides an Alicloud Common Attachment resource.`,
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
					Description: `The ID of the common bandwidth package attachment id and formates as ` + "`" + `<bandwidth_package_id>:<instance_id>` + "`" + `. ## Import The common bandwidth package attachemnt can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_common_bandwidth_package_attachment.foo cbwp-abc123456:eip-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the common bandwidth package attachment id and formates as ` + "`" + `<bandwidth_package_id>:<instance_id>` + "`" + `. ## Import The common bandwidth package attachemnt can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_common_bandwidth_package_attachment.foo cbwp-abc123456:eip-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_container_cluster",
			Category:         "Container Service (CS)",
			ShortDescription: `Provides a Alicloud container cluster resource.`,
			Description:      ``,
			Keywords: []string{
				"container",
				"service",
				"cs",
				"cluster",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cr_ee_namespace",
			Category:         "Container Registry (CR)",
			ShortDescription: `Provides a Alicloud resource to manage Container Registry Enterprise Edition namespaces.`,
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
					Description: `ID of Container Registry Enterprise Edition namespace. The value is in format ` + "`" + `{instance_id}:{namespace}` + "`" + ` . ## Import Container Registry Enterprise Edition namespace can be imported using the ` + "`" + `{instance_id}:{namespace}` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cr_ee_namespace.default cri-xxx:my-namespace ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of Container Registry Enterprise Edition namespace. The value is in format ` + "`" + `{instance_id}:{namespace}` + "`" + ` . ## Import Container Registry Enterprise Edition namespace can be imported using the ` + "`" + `{instance_id}:{namespace}` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cr_ee_namespace.default cri-xxx:my-namespace ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cr_ee_repo",
			Category:         "Container Registry (CR)",
			ShortDescription: `Provides a Alicloud resource to manage Container Registry Enterprise Edition repositories.`,
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
					Description: `The uuid of Container Registry Enterprise Edition repository. ## Import Container Registry Enterprise Edition repository can be imported using the ` + "`" + `{instance_id}:{namespace}:{repository}` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cr_ee_repo.default ` + "`" + `cri-xxx:my-namespace:my-repo` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource id of Container Registry Enterprise Edition repository. The value is in format ` + "`" + `{instance_id}:{namespace}:{repository}` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "repo_id",
					Description: `The uuid of Container Registry Enterprise Edition repository. ## Import Container Registry Enterprise Edition repository can be imported using the ` + "`" + `{instance_id}:{namespace}:{repository}` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cr_ee_repo.default ` + "`" + `cri-xxx:my-namespace:my-repo` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cr_ee_sync_rule",
			Category:         "Container Registry (CR)",
			ShortDescription: `Provides a Alicloud resource to manage Container Registry Enterprise Edition sync rules.`,
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
					Description: `` + "`" + `REPO` + "`" + ` or ` + "`" + `NAMESPACE` + "`" + `,the scope that the synchronization rule applies. ## Import Container Registry Enterprise Edition sync rule can be imported using the id. Format to ` + "`" + `{instance_id}:{namespace_name}:{rule_id}` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cr_ee_sync_rule.default ` + "`" + `cri-xxx:my-namespace:crsr-yyy` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `` + "`" + `REPO` + "`" + ` or ` + "`" + `NAMESPACE` + "`" + `,the scope that the synchronization rule applies. ## Import Container Registry Enterprise Edition sync rule can be imported using the id. Format to ` + "`" + `{instance_id}:{namespace_name}:{rule_id}` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cr_ee_sync_rule.default ` + "`" + `cri-xxx:my-namespace:crsr-yyy` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cr_namespace",
			Category:         "Container Registry (CR)",
			ShortDescription: `Provides a Alicloud resource to manage Container Registry namespaces.`,
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
					Description: `The id of Container Registry namespace. The value is same as its name. ## Import Container Registry namespace can be imported using the namespace, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cr_namespace.default my-namespace ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of Container Registry namespace. The value is same as its name. ## Import Container Registry namespace can be imported using the namespace, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cr_namespace.default my-namespace ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cr_repo",
			Category:         "Container Registry (CR)",
			ShortDescription: `Provides a Alicloud resource to manage Container Registry repositories.`,
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
					Description: `Domain of vpc endpoint. ## Import Container Registry repository can be imported using the ` + "`" + `namespace/repository` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cr_repo.default ` + "`" + `my-namespace/my-repo` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `Domain of vpc endpoint. ## Import Container Registry repository can be imported using the ` + "`" + `namespace/repository` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cr_repo.default ` + "`" + `my-namespace/my-repo` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cs_application",
			Category:         "Container Service (CS)",
			ShortDescription: `Provides a resource to deploy application in one container cluster.`,
			Description:      ``,
			Keywords: []string{
				"container",
				"service",
				"cs",
				"application",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Required, ForceNew) The swarm cluster's name.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) The application name. It should be 1-64 characters long, and can contain numbers, English letters and hyphens, but cannot start with hyphens.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of application.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The application deploying version. Each updating, it must be different with current. Default to "1.0"`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Required) The application deployment template and it must be [Docker Compose format](https://docs.docker.com/compose/).`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `A key/value map used to replace the variable parameter in the Compose template.`,
				},
				resource.Attribute{
					Name:        "latest_image",
					Description: `Whether to use latest docker image while each updating application. Default to false.`,
				},
				resource.Attribute{
					Name:        "blue_green",
					Description: `Wherther to use "Blue Green" method when release a new version. Default to false.`,
				},
				resource.Attribute{
					Name:        "blue_green_confirm",
					Description: `Whether to confirm a "Blue Green" application. Default to false. It will be ignored when ` + "`" + `blue_green` + "`" + ` is false. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the container application. It's formate is ` + "`" + `<cluster_name>:<name>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `The name of the container cluster.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The application name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The application description.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `The application deploying template.`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `The application environment variables.`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `List of services in the application. It contains several attributes to ` + "`" + `Block Nodes` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "default_domain",
					Description: `The application default domain and it can be used to configure routing service. ### Block Nodes`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the service.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Service name.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current status of service.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The current version of service. ## Import Swarm application can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cs_application.app my-first-swarm:wordpress ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the container application. It's formate is ` + "`" + `<cluster_name>:<name>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `The name of the container cluster.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The application name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The application description.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `The application deploying template.`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `The application environment variables.`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `List of services in the application. It contains several attributes to ` + "`" + `Block Nodes` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "default_domain",
					Description: `The application default domain and it can be used to configure routing service. ### Block Nodes`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the service.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Service name.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current status of service.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The current version of service. ## Import Swarm application can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cs_application.app my-first-swarm:wordpress ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cs_kubernetes",
			Category:         "Container Service (CS)",
			ShortDescription: `Provides a Alicloud resource to manage container kubernetes cluster.`,
			Description:      ``,
			Keywords: []string{
				"container",
				"service",
				"cs",
				"kubernetes",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The kubernetes cluster's name. It is unique in one Alicloud account.`,
				},
				resource.Attribute{
					Name:        "name_prefix",
					Description: `(Optional) The kubernetes cluster name's prefix. It is conflict with ` + "`" + `name` + "`" + `. If it is specified, terraform will using it to build the only cluster name. Default to "Terraform-Creation".`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional, Available since 1.70.1) Desired Kubernetes version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except you set a higher version number. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by ACK.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required, Sensitive) The password of ssh login cluster node. You have to specify one of ` + "`" + `password` + "`" + ` ` + "`" + `key_name` + "`" + ` ` + "`" + `kms_encrypted_password` + "`" + ` fields.`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `(Required) The keypair of ssh login cluster node, you have to create it first. You have to specify one of ` + "`" + `password` + "`" + ` ` + "`" + `key_name` + "`" + ` ` + "`" + `kms_encrypted_password` + "`" + ` fields.`,
				},
				resource.Attribute{
					Name:        "kms_encrypted_password",
					Description: `(Required, Available in 1.57.1+) An KMS encrypts password used to a cs kubernetes. You have to specify one of ` + "`" + `password` + "`" + ` ` + "`" + `key_name` + "`" + ` ` + "`" + `kms_encrypted_password` + "`" + ` fields.`,
				},
				resource.Attribute{
					Name:        "kms_encryption_context",
					Description: `(Optional, MapString, Available in 1.57.1+) An KMS encryption context used to decrypt ` + "`" + `kms_encrypted_password` + "`" + ` before creating or updating a cs kubernetes with ` + "`" + `kms_encrypted_password` + "`" + `. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when ` + "`" + `kms_encrypted_password` + "`" + ` is set.`,
				},
				resource.Attribute{
					Name:        "user_ca",
					Description: `(Optional, ForceNew) The path of customized CA cert, you can use this CA to sign client certs to connect your cluster.`,
				},
				resource.Attribute{
					Name:        "enable_ssh",
					Description: `(Optional) Enable login to the node through SSH. default: false`,
				},
				resource.Attribute{
					Name:        "install_cloud_monitor",
					Description: `(Optional) Install cloud monitor agent on ECS. default: true`,
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
					Name:        "image_id",
					Description: `Custom Image support. Must based on CentOS7 or AliyunLinux2.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional, Available in 1.81.0+) Windows instances support batch and PowerShell scripts. If your script file is larger than 1 KB, we recommend that you upload the script to Object Storage Service (OSS) and pull it through the internal endpoint of your OSS bucket.`,
				},
				resource.Attribute{
					Name:        "exclude_autoscaler_nodes",
					Description: `(Optional, Available in 1.88.0+) Exclude autoscaler nodes from ` + "`" + `worker_nodes` + "`" + `. default: false`,
				},
				resource.Attribute{
					Name:        "node_name_mode",
					Description: `(Optional, Available in 1.88.0+) Each node name consists of a prefix, an IP substring, and a suffix. For example, if the node IP address is 192.168.0.55, the prefix is aliyun.com, IP substring length is 5, and the suffix is test, the node name will be aliyun.com00055test. #### Addons It is a new field since 1.75.0. You can specific network plugin,log component,ingress component and so on. ` + "`" + `` + "`" + `` + "`" + `$xslt main.tf dynamic "addons" { for_each = var.cluster_addons content { name = lookup(addons.value, "name", var.cluster_addons) config = lookup(addons.value, "config", var.cluster_addons) } } ` + "`" + `` + "`" + `` + "`" + ` ` + "`" + `` + "`" + `` + "`" + `$xslt varibales.tf // Network-flannel variable "cluster_addons" { description = "Addon components in kubernetes cluster" type = list(object({ name = string config = string })) default = [ { "name" = "flannel", "config" = "", } ] } // Network-terway variable "cluster_addons" { type = list(object({ name = string config = string })) default = [ { "name" = "terway-eniip", "config" = "", } ] } // Storage-csi variable "cluster_addons" { type = list(object({ name = string config = string })) default = [ { "name" = "csi-plugin", "config" = "", }, { "name" = "csi-provisioner", "config" = "", } ] } // Storage-flexvolume variable "cluster_addons" { type = list(object({ name = string config = string })) default = [ { "name" = "flexvolume", "config" = "", } ] } // Log variable "cluster_addons" { type = list(object({ name = string config = string })) default = [ { "name" = "logtail-ds", "config" = "{\"IngressDashboardEnabled\":\"true\",\"sls_project_name\":\"your-sls-project-name\"}", } ] } // Ingress variable "cluster_addons" { type = list(object({ name = string config = string })) default = [ { "name" = "nginx-ingress-controller", "config" = "{\"IngressSlbNetworkType\":\"internet\"}", } ] } // Ingress-Disable variable "cluster_addons" { type = list(object({ name = string config = string disabled = bool })) default = [ { "name" = "nginx-ingress-controller", "config" = "", "disabled": true, } ] } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "logtail-ds",
					Description: `You can specific ` + "`" + `IngressDashboardEnabled` + "`" + ` and ` + "`" + `sls_project_name` + "`" + ` in config. If you switch on ` + "`" + `IngressDashboardEnabled` + "`" + ` and ` + "`" + `sls_project_name` + "`" + `,then logtail-ds would use ` + "`" + `sls_project_name` + "`" + ` as default log store.`,
				},
				resource.Attribute{
					Name:        "nginx-ingress-controller",
					Description: `You can specific ` + "`" + `IngressSlbNetworkType` + "`" + ` in config. Options: internet|intranet. You can get more information about addons on ACK web console. When you create a ACK cluster. You can get openapi-spec before creating the cluster on submission page. #### Network`,
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
					Description: `(Optional) Whether to create a new nat gateway while creating kubernetes cluster. Default to true. Then openapi in Alibaba Cloud are not all on intranet, So turn this option on is a good choice.`,
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
					Description: `(Optional) Whether to create internet load balancer for API Server. Default to true. If you want to use ` + "`" + `Terway` + "`" + ` as CNI network plugin, You need to specific the ` + "`" + `pod_vswitch_ids` + "`" + ` field and addons with ` + "`" + `terway-eniip` + "`" + `. If you want to use ` + "`" + `Flannel` + "`" + ` as CNI network plugin, You need to specific the ` + "`" + `pod_cidr` + "`" + ` field and addons with ` + "`" + `flannel` + "`" + `. #### Master params`,
				},
				resource.Attribute{
					Name:        "master_vswtich_ids",
					Description: `(Required) The vswitches used by master, you can specific 3 or 5 vswitches because of the amount of masters. You can also specific`,
				},
				resource.Attribute{
					Name:        "master_instance_types",
					Description: `(Required) The instance type of master node. Specify one type for single AZ Cluster, three types for MultiAZ Cluster.`,
				},
				resource.Attribute{
					Name:        "master_instance_charge_type",
					Description: `(Optional) Master payment type. ` + "`" + `PrePaid` + "`" + ` or ` + "`" + `PostPaid` + "`" + `, defaults to ` + "`" + `PostPaid` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "master_period_unit",
					Description: `(Optional) Master payment period unit. ` + "`" + `Month` + "`" + ` or ` + "`" + `Week` + "`" + `, defaults to ` + "`" + `Month` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "master_period",
					Description: `(Optional) Master payment period. When period unit is ` + "`" + `Month` + "`" + `, it can be one of { “1”, “2”, “3”, “4”, “5”, “6”, “7”, “8”, “9”, “12”, “24”, “36”,”48”,”60”}. When period unit is ` + "`" + `Week` + "`" + `, it can be one of {“1”, “2”, “3”, “4”}.`,
				},
				resource.Attribute{
					Name:        "master_auto_renew",
					Description: `(Optional) Enable master payment auto-renew, defaults to false.`,
				},
				resource.Attribute{
					Name:        "master_auto_renew_period",
					Description: `(Optional) Master payment auto-renew period. When period unit is ` + "`" + `Month` + "`" + `, it can be one of {“1”, “2”, “3”, “6”, “12”}. When period unit is ` + "`" + `Week` + "`" + `, it can be one of {“1”, “2”, “3”}.`,
				},
				resource.Attribute{
					Name:        "master_disk_category",
					Description: `(Optional) The system disk category of master node. Its valid value are ` + "`" + `cloud_ssd` + "`" + ` and ` + "`" + `cloud_efficiency` + "`" + `. Default to ` + "`" + `cloud_efficiency` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "master_disk_size",
					Description: `(Optional) The system disk size of master node. Its valid value range [20~500] in GB. Default to 20. #### Worker params`,
				},
				resource.Attribute{
					Name:        "worker_number",
					Description: `(Required) The worker node number of the kubernetes cluster. Default to 3. It is limited up to 50 and if you want to enlarge it, please apply white list or contact with us.`,
				},
				resource.Attribute{
					Name:        "worker_vswtich_ids",
					Description: `(Required) The vswitches used by workers.`,
				},
				resource.Attribute{
					Name:        "worker_instance_types",
					Description: `(Required, ForceNew) The instance type of worker node. Specify one type for single AZ Cluster, three types for MultiAZ Cluster.`,
				},
				resource.Attribute{
					Name:        "worker_instance_charge_type",
					Description: `(Optional, Force new resource) Worker payment type. ` + "`" + `PrePaid` + "`" + ` or ` + "`" + `PostPaid` + "`" + `, defaults to ` + "`" + `PostPaid` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "worker_period_unit",
					Description: `(Optional) Worker payment period unit. ` + "`" + `Month` + "`" + ` or ` + "`" + `Week` + "`" + `, defaults to ` + "`" + `Month` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "worker_period",
					Description: `(Optional) Worker payment period. When period unit is ` + "`" + `Month` + "`" + `, it can be one of { “1”, “2”, “3”, “4”, “5”, “6”, “7”, “8”, “9”, “12”, “24”, “36”,”48”,”60”}. When period unit is ` + "`" + `Week` + "`" + `, it can be one of {“1”, “2”, “3”, “4”}.`,
				},
				resource.Attribute{
					Name:        "worker_auto_renew",
					Description: `(Optional) Enable worker payment auto-renew, defaults to false.`,
				},
				resource.Attribute{
					Name:        "worker_auto_renew_period",
					Description: `(Optional) Worker payment auto-renew period. When period unit is ` + "`" + `Month` + "`" + `, it can be one of {“1”, “2”, “3”, “6”, “12”}. When period unit is ` + "`" + `Week` + "`" + `, it can be one of {“1”, “2”, “3”}.`,
				},
				resource.Attribute{
					Name:        "worker_disk_category",
					Description: `(Optional) The system disk category of worker node. Its valid value are ` + "`" + `cloud_ssd` + "`" + ` and ` + "`" + `cloud_efficiency` + "`" + `. Default to ` + "`" + `cloud_efficiency` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "worker_disk_size",
					Description: `(Optional) The system disk size of worker node. Its valid value range [20~32768] in GB. Default to 40. #### Computed params (No need to configure)`,
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
					Description: `(Optional) The Zone where new kubernetes cluster will be located. If it is not be specified, the ` + "`" + `vswitch_ids` + "`" + ` should be set, its value will be vswitch's zone. #### Removed params (Never Supported)`,
				},
				resource.Attribute{
					Name:        "master_instance_type",
					Description: `(Deprecated from version 1.16.0)(Required, Force new resource) The instance type of master node.`,
				},
				resource.Attribute{
					Name:        "worker_instance_type",
					Description: `(Deprecated from version 1.16.0)(Required, Force new resource) The instance type of worker node.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Deprecated from version 1.16.0)(Force new resource) The vswitch where new kubernetes cluster will be located. If it is not specified, a new VPC and VSwicth will be built. It must be in the zone which ` + "`" + `availability_zone` + "`" + ` specified.`,
				},
				resource.Attribute{
					Name:        "vswitch_ids",
					Description: `(Required, ForceNew) The vswitch where new kubernetes cluster will be located. Specify one or more vswitch's id. It must be in the zone which ` + "`" + `availability_zone` + "`" + ` specified.`,
				},
				resource.Attribute{
					Name:        "force_update",
					Description: `(Optional, Available in 1.50.0+) Whether to force the update of kubernetes cluster arguments. Default to false.`,
				},
				resource.Attribute{
					Name:        "is_outdated",
					Description: `(Optional) Whether to use outdated instance type. Default to false.`,
				},
				resource.Attribute{
					Name:        "log_config",
					Description: `(Optional, ForceNew) A list of one element containing information about the associated log store. It contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of collecting logs, only ` + "`" + `SLS` + "`" + ` are supported currently.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `Log Service project name, cluster logs will output to this project.`,
				},
				resource.Attribute{
					Name:        "cluster_network_type",
					Description: `(Optional) The network that cluster uses, use ` + "`" + `flannel` + "`" + ` or ` + "`" + `terway` + "`" + `. ### Timeouts ->`,
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
					Description: `The RamRole Name attached to worker node. ### Block Nodes`,
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
					Description: `The private IP address of node.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Deprecated from version 1.9.4) ### Block Connections`,
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
					Description: `Service Access Domain. ## Import Kubernetes cluster can be imported using the id, e.g. Then complete the main.tf accords to the result of ` + "`" + `terraform plan` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cs_kubernetes.main cluster-id ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `The RamRole Name attached to worker node. ### Block Nodes`,
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
					Description: `The private IP address of node.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Deprecated from version 1.9.4) ### Block Connections`,
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
					Description: `Service Access Domain. ## Import Kubernetes cluster can be imported using the id, e.g. Then complete the main.tf accords to the result of ` + "`" + `terraform plan` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cs_kubernetes.main cluster-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cs_kubernetes_autoscaler",
			Category:         "Container Service (CS)",
			ShortDescription: `Provides a Alicloud resource to manage container kubernetes cluster-autoscaler.`,
			Description:      ``,
			Keywords: []string{
				"container",
				"service",
				"cs",
				"kubernetes",
				"autoscaler",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) The id of kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "nodepools",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "nodepools.id",
					Description: `(Required) The scaling group id of the groups configured for cluster-autoscaler.`,
				},
				resource.Attribute{
					Name:        "nodepools.taints",
					Description: `(Required) The taints for the nodes in scaling group.`,
				},
				resource.Attribute{
					Name:        "nodepools.labels",
					Description: `(Required) The labels for the nodes in scaling group.`,
				},
				resource.Attribute{
					Name:        "utilization",
					Description: `(Required) The utilization option of cluster-autoscaler.`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 90 mins) Used when creating cluster-autoscaler in the kubernetes cluster (until it reaches the initial ` + "`" + `running` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 60 mins) Used when activating the cluster-autoscaler in the kubernetes cluster when necessary during update.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 60 mins) Used when deleting cluster-autoscaler in kubernetes cluster.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cs_managed_kubernetes",
			Category:         "Container Service (CS)",
			ShortDescription: `Provides a Alicloud resource to manage container managed kubernetes cluster.`,
			Description:      ``,
			Keywords: []string{
				"container",
				"service",
				"cs",
				"managed",
				"kubernetes",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The kubernetes cluster's name. It is unique in one Alicloud account.`,
				},
				resource.Attribute{
					Name:        "name_prefix",
					Description: `(Optional) The kubernetes cluster name's prefix. It is conflict with ` + "`" + `name` + "`" + `. If it is specified, terraform will using it to build the only cluster name. Default to "Terraform-Creation".`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional, Available since 1.70.1) Desired Kubernetes version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except you set a higher version number. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by ACK.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required, Sensitive) The password of ssh login cluster node. You have to specify one of ` + "`" + `password` + "`" + ` ` + "`" + `key_name` + "`" + ` ` + "`" + `kms_encrypted_password` + "`" + ` fields.`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `(Required) The keypair of ssh login cluster node, you have to create it first. You have to specify one of ` + "`" + `password` + "`" + ` ` + "`" + `key_name` + "`" + ` ` + "`" + `kms_encrypted_password` + "`" + ` fields.`,
				},
				resource.Attribute{
					Name:        "kms_encrypted_password",
					Description: `(Required, Available in 1.57.1+) An KMS encrypts password used to a cs kubernetes. You have to specify one of ` + "`" + `password` + "`" + ` ` + "`" + `key_name` + "`" + ` ` + "`" + `kms_encrypted_password` + "`" + ` fields.`,
				},
				resource.Attribute{
					Name:        "kms_encryption_context",
					Description: `(Optional, MapString, Available in 1.57.1+) An KMS encryption context used to decrypt ` + "`" + `kms_encrypted_password` + "`" + ` before creating or updating a cs kubernetes with ` + "`" + `kms_encrypted_password` + "`" + `. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when ` + "`" + `kms_encrypted_password` + "`" + ` is set.`,
				},
				resource.Attribute{
					Name:        "user_ca",
					Description: `(Optional, ForceNew) The path of customized CA cert, you can use this CA to sign client certs to connect your cluster.`,
				},
				resource.Attribute{
					Name:        "enable_ssh",
					Description: `(Optional) Enable login to the node through SSH. default: false`,
				},
				resource.Attribute{
					Name:        "install_cloud_monitor",
					Description: `(Optional) Install cloud monitor agent on ECS. default: true`,
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
					Name:        "image_id",
					Description: `Custom Image support. Must based on CentOS7 or AliyunLinux2.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional, Available in 1.81.0+) Windows instances support batch and PowerShell scripts. If your script file is larger than 1 KB, we recommend that you upload the script to Object Storage Service (OSS) and pull it through the internal endpoint of your OSS bucket.`,
				},
				resource.Attribute{
					Name:        "exclude_autoscaler_nodes",
					Description: `(Optional, Available in 1.88.0+) Exclude autoscaler nodes from ` + "`" + `worker_nodes` + "`" + `. default: false`,
				},
				resource.Attribute{
					Name:        "node_name_mode",
					Description: `(Optional, Available in 1.88.0+) Each node name consists of a prefix, an IP substring, and a suffix. For example, if the node IP address is 192.168.0.55, the prefix is aliyun.com, IP substring length is 5, and the suffix is test, the node name will be aliyun.com00055test. #### Addons It is a new field since 1.75.0. You can specific network plugin,log component,ingress component and so on. ` + "`" + `` + "`" + `` + "`" + `$xslt main.tf dynamic "addons" { for_each = var.cluster_addons content { name = lookup(addons.value, "name", var.cluster_addons) config = lookup(addons.value, "config", var.cluster_addons) } } ` + "`" + `` + "`" + `` + "`" + ` ` + "`" + `` + "`" + `` + "`" + `$xslt varibales.tf // Network-flannel variable "cluster_addons" { description = "Addon components in kubernetes cluster" type = list(object({ name = string config = string })) default = [ { "name" = "flannel", "config" = "", } ] } // Network-terway variable "cluster_addons" { type = list(object({ name = string config = string })) default = [ { "name" = "terway-eniip", "config" = "", } ] } // Storage-csi variable "cluster_addons" { type = list(object({ name = string config = string })) default = [ { "name" = "csi-plugin", "config" = "", }, { "name" = "csi-provisioner", "config" = "", } ] } // Storage-flexvolume variable "cluster_addons" { type = list(object({ name = string config = string })) default = [ { "name" = "flexvolume", "config" = "", } ] } // Log variable "cluster_addons" { type = list(object({ name = string config = string })) default = [ { "name" = "logtail-ds", "config" = "{\"IngressDashboardEnabled\":\"true\",\"sls_project_name\":\"your-sls-project-name\"}", } ] } // Ingress variable "cluster_addons" { type = list(object({ name = string config = string })) default = [ { "name" = "nginx-ingress-controller", "config" = "{\"IngressSlbNetworkType\":\"internet\"}", } ] } // Ingress-Disable variable "cluster_addons" { type = list(object({ name = string config = string disabled = bool })) default = [ { "name" = "nginx-ingress-controller", "config" = "", "disabled": true, } ] } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "logtail-ds",
					Description: `You can specific ` + "`" + `IngressDashboardEnabled` + "`" + ` and ` + "`" + `sls_project_name` + "`" + ` in config. If you switch on ` + "`" + `IngressDashboardEnabled` + "`" + ` and ` + "`" + `sls_project_name` + "`" + `,then logtail-ds would use ` + "`" + `sls_project_name` + "`" + ` as default log store.`,
				},
				resource.Attribute{
					Name:        "nginx-ingress-controller",
					Description: `You can specific ` + "`" + `IngressSlbNetworkType` + "`" + ` in config. Options: internet|intranet. You can get more information about addons on ACK web console. When you create a ACK cluster. You can get openapi-spec before creating the cluster on submission page. #### Network`,
				},
				resource.Attribute{
					Name:        "pod_cidr",
					Description: `(Required) [Flannel Specific] The CIDR block for the pod network when using Flannel.`,
				},
				resource.Attribute{
					Name:        "pod_vswitch_ids",
					Description: `(Required) [Terway Specific] The vswitches for the pod network when using Terway.Be careful the ` + "`" + `pod_vswitch_ids` + "`" + ` can not equal to ` + "`" + `worker_vswtich_ids` + "`" + `.but must be in same availability zones.`,
				},
				resource.Attribute{
					Name:        "new_nat_gateway",
					Description: `(Optional) Whether to create a new nat gateway while creating kubernetes cluster. Default to true. Then openapi in Alibaba Cloud are not all on intranet, So turn this option on is a good choice.`,
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
					Description: `(Optional) Whether to create internet load balancer for API Server. Default to true. If you want to use ` + "`" + `Terway` + "`" + ` as CNI network plugin, You need to specific the ` + "`" + `pod_vswitch_ids` + "`" + ` field and addons with ` + "`" + `terway-eniip` + "`" + ` or ` + "`" + `terway-eni` + "`" + `. The ` + "`" + `terway-eni` + "`" + ` mode for pod with one exclude ENI, the ` + "`" + `terway-eniip` + "`" + ` mode for pods share ENI. If you want to use ` + "`" + `Flannel` + "`" + ` as CNI network plugin, You need to specific the ` + "`" + `pod_cidr` + "`" + ` field and addons with ` + "`" + `flannel` + "`" + `. #### Worker params`,
				},
				resource.Attribute{
					Name:        "worker_number",
					Description: `(Required) The worker node number of the kubernetes cluster. Default to 3. It is limited up to 50 and if you want to enlarge it, please apply white list or contact with us.`,
				},
				resource.Attribute{
					Name:        "worker_vswtich_ids",
					Description: `(Required) The vswitches used by workers.`,
				},
				resource.Attribute{
					Name:        "worker_instance_types",
					Description: `(Required, ForceNew) The instance type of worker node. Specify one type for single AZ Cluster, three types for MultiAZ Cluster.`,
				},
				resource.Attribute{
					Name:        "worker_instance_charge_type",
					Description: `(Optional, Force new resource) Worker payment type. ` + "`" + `PrePaid` + "`" + ` or ` + "`" + `PostPaid` + "`" + `, defaults to ` + "`" + `PostPaid` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "worker_period_unit",
					Description: `(Optional) Worker payment period unit. ` + "`" + `Month` + "`" + ` or ` + "`" + `Week` + "`" + `, defaults to ` + "`" + `Month` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "worker_period",
					Description: `(Optional) Worker payment period. When period unit is ` + "`" + `Month` + "`" + `, it can be one of { “1”, “2”, “3”, “4”, “5”, “6”, “7”, “8”, “9”, “12”, “24”, “36”,”48”,”60”}. When period unit is ` + "`" + `Week` + "`" + `, it can be one of {“1”, “2”, “3”, “4”}.`,
				},
				resource.Attribute{
					Name:        "worker_auto_renew",
					Description: `(Optional) Enable worker payment auto-renew, defaults to false.`,
				},
				resource.Attribute{
					Name:        "worker_auto_renew_period",
					Description: `(Optional) Worker payment auto-renew period. When period unit is ` + "`" + `Month` + "`" + `, it can be one of {“1”, “2”, “3”, “6”, “12”}. When period unit is ` + "`" + `Week` + "`" + `, it can be one of {“1”, “2”, “3”}.`,
				},
				resource.Attribute{
					Name:        "worker_disk_category",
					Description: `(Optional) The system disk category of worker node. Its valid value are ` + "`" + `cloud_ssd` + "`" + ` and ` + "`" + `cloud_efficiency` + "`" + `. Default to ` + "`" + `cloud_efficiency` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "worker_disk_size",
					Description: `(Optional) The system disk size of worker node. Its valid value range [20~32768] in GB. Default to 40. #### Computed params (No need to configure)`,
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
					Description: `(Optional) The Zone where new kubernetes cluster will be located. If it is not be specified, the ` + "`" + `vswitch_ids` + "`" + ` should be set, its value will be vswitch's zone. #### Removed params (Never Supported)`,
				},
				resource.Attribute{
					Name:        "worker_instance_type",
					Description: `(Deprecated from version 1.16.0)(Required, Force new resource) The instance type of worker node.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Deprecated from version 1.16.0)(Force new resource) The vswitch where new kubernetes cluster will be located. If it is not specified, a new VPC and VSwicth will be built. It must be in the zone which ` + "`" + `availability_zone` + "`" + ` specified.`,
				},
				resource.Attribute{
					Name:        "vswitch_ids",
					Description: `(Required, ForceNew) The vswitch where new kubernetes cluster will be located. Specify one or more vswitch's id. It must be in the zone which ` + "`" + `availability_zone` + "`" + ` specified.`,
				},
				resource.Attribute{
					Name:        "force_update",
					Description: `(Optional, Available in 1.50.0+) Whether to force the update of kubernetes cluster arguments. Default to false.`,
				},
				resource.Attribute{
					Name:        "is_outdated",
					Description: `(Optional) Whether to use outdated instance type. Default to false.`,
				},
				resource.Attribute{
					Name:        "log_config",
					Description: `(Optional, ForceNew) A list of one element containing information about the associated log store. It contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of collecting logs, only ` + "`" + `SLS` + "`" + ` are supported currently.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `Log Service project name, cluster logs will output to this project.`,
				},
				resource.Attribute{
					Name:        "cluster_network_type",
					Description: `(Optional) The network that cluster uses, use ` + "`" + `flannel` + "`" + ` or ` + "`" + `terway` + "`" + `. ### Timeouts ->`,
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
					Description: `The RamRole Name attached to worker node. ### Block Nodes`,
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
					Description: `The private IP address of node.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Deprecated from version 1.9.4) ### Block Connections`,
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
					Name:        "service_domain",
					Description: `Service Access Domain. ## Import Kubernetes cluster can be imported using the id, e.g. Then complete the main.tf accords to the result of ` + "`" + `terraform plan` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cs_managed_kubernetes.main cluster-id ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `The RamRole Name attached to worker node. ### Block Nodes`,
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
					Description: `The private IP address of node.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Deprecated from version 1.9.4) ### Block Connections`,
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
					Name:        "service_domain",
					Description: `Service Access Domain. ## Import Kubernetes cluster can be imported using the id, e.g. Then complete the main.tf accords to the result of ` + "`" + `terraform plan` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cs_managed_kubernetes.main cluster-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cs_serverless_kubernetes",
			Category:         "Container Service (CS)",
			ShortDescription: `Provides a Alicloud resource to manage container serverless kubernetes cluster.`,
			Description:      ``,
			Keywords: []string{
				"container",
				"service",
				"cs",
				"serverless",
				"kubernetes",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The kubernetes cluster's name. It is the only in one Alicloud account.`,
				},
				resource.Attribute{
					Name:        "name_prefix",
					Description: `(Optional) The kubernetes cluster name's prefix. It is conflict with ` + "`" + `name` + "`" + `. If it is specified, terraform will using it to build the only cluster name. Default to "Terraform-Creation".`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, ForceNew) The vpc where new kubernetes cluster will be located. Specify one vpc's id, if it is not specified, a new VPC will be built.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Required, ForceNew) The vswitch where new kubernetes cluster will be located. Specify one vswitch's id, if it is not specified, a new VPC and VSwicth will be built. It must be in the zone which ` + "`" + `availability_zone` + "`" + ` specified.`,
				},
				resource.Attribute{
					Name:        "new_nat_gateway",
					Description: `(Optional) Whether to create a new nat gateway while creating kubernetes cluster. Default to true.`,
				},
				resource.Attribute{
					Name:        "endpoint_public_access_enabled",
					Description: `(Optional, ForceNew) Whether to create internet eip for API Server. Default to false.`,
				},
				resource.Attribute{
					Name:        "private_zone",
					Description: `(Optional, ForceNew) Enable Privatezone if you need to use the service discovery feature within the serverless cluster. Default to false.`,
				},
				resource.Attribute{
					Name:        "deletion_protection",
					Description: `(Optional, ForceNew) Whether enable the deletion protection or not. - true: Enable deletion protection. - false: Disable deletion protection.`,
				},
				resource.Attribute{
					Name:        "force_update",
					Description: `(Optional) Default false, when you want to change ` + "`" + `vpc_id` + "`" + ` and ` + "`" + `vswitch_id` + "`" + `, you have to set this field to true, then the cluster will be recreated.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Default nil, A map of tags assigned to the kubernetes cluster .`,
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
					Description: `(Optional) The path of cluster ca certificate, like ` + "`" + `~/.kube/cluster-ca-cert.pem` + "`" + ` ### Timeouts ->`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 60 mins) Used when creating the kubernetes cluster (until it reaches the initial ` + "`" + `running` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 30 mins) Used when terminating the kubernetes cluster. ## Attributes Reference The following attributes are exported:`,
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
					Name:        "vpc_id",
					Description: `The ID of VPC where the current cluster is located.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `The ID of VSwicth where the current cluster is located.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of security group where the current cluster worker node is located.`,
				},
				resource.Attribute{
					Name:        "deletion_protection",
					Description: `Whether enable the deletion protection or not. ## Import Serverless Kubernetes cluster can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cs_serverless_kubernetes.main ce4273f9156874b46bb ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "vpc_id",
					Description: `The ID of VPC where the current cluster is located.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `The ID of VSwicth where the current cluster is located.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of security group where the current cluster worker node is located.`,
				},
				resource.Attribute{
					Name:        "deletion_protection",
					Description: `Whether enable the deletion protection or not. ## Import Serverless Kubernetes cluster can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cs_serverless_kubernetes.main ce4273f9156874b46bb ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cs_swarm",
			Category:         "Container Service (CS)",
			ShortDescription: `Provides a Alicloud container swarm cluster resource.`,
			Description:      ``,
			Keywords: []string{
				"container",
				"service",
				"cs",
				"swarm",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The container cluster's name. It is the only in one Alicloud account.`,
				},
				resource.Attribute{
					Name:        "name_prefix",
					Description: `The container cluster name's prefix. It is conflict with ` + "`" + `name` + "`" + `. If it is specified, terraform will using it to build the only cluster name. Default to 'Terraform-Creation'.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Field 'size' has been deprecated from provider version 1.9.1. New field 'node_number' replaces it.`,
				},
				resource.Attribute{
					Name:        "node_number",
					Description: `The ECS node number of the container cluster. Its value choices are 1~50, and default to 1.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Required, ForceNew) The CIDR block for the Container. It can not be same as the CIDR used by the VPC. Valid value: - 192.168.0.0/16 - 172.19-30.0.0/16 - 10.0.0.0/16 System reserved private network address: 172.16/17/18/31.0.0/16. Maximum number of hosts allowed in the cluster: 256.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(ForceNew) The image ID of ECS instance node used. Default to System automate allocated.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Required, ForceNew) The type of ECS instance node.`,
				},
				resource.Attribute{
					Name:        "is_outdated",
					Description: `(Optional) Whether to use outdated instance type. Default to false.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required, ForceNew, Sensitive) The password of ECS instance node.`,
				},
				resource.Attribute{
					Name:        "disk_category",
					Description: `(ForceNew) The data disk category of ECS instance node. Its valid value are ` + "`" + `cloud` + "`" + `, ` + "`" + `cloud_ssd` + "`" + `, ` + "`" + `cloud_essd` + "`" + `, ` + "`" + `ephemeral_essd` + "`" + ` and ` + "`" + `cloud_efficiency` + "`" + `. Default to ` + "`" + `cloud_efficiency` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `(ForceNew) The data disk size of ECS instance node. Its valid value is 20~32768 GB. Default to 20.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Required, ForceNew) The password of ECS instance node. If it is not specified, the container cluster's network mode will be ` + "`" + `Classic` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "release_eip",
					Description: `Whether to release EIP after creating swarm cluster successfully. Default to false.`,
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
					Name:        "size",
					Description: `It has been deprecated from provider version 1.9.1. New field 'node_number' replaces it.`,
				},
				resource.Attribute{
					Name:        "node_number",
					Description: `The node number.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of VPC where the current cluster is located.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `The ID of VSwitch where the current cluster is located.`,
				},
				resource.Attribute{
					Name:        "slb_id",
					Description: `The ID of load balancer where the current cluster worker node is located.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of security group where the current cluster worker node is located.`,
				},
				resource.Attribute{
					Name:        "agent_version",
					Description: `The nodes agent version.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `The instance type of nodes.`,
				},
				resource.Attribute{
					Name:        "disk_category",
					Description: `The data disk category of nodes.`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `The data disk size of nodes.`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `List of cluster nodes. It contains several attributes to ` + "`" + `Block Nodes` + "`" + `. ### Block Nodes`,
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
					Description: `The private IP address of node.`,
				},
				resource.Attribute{
					Name:        "eip",
					Description: `The Elastic IP address of node.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The node current status. It is different with instance status. ## Import Swarm cluster can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cs_swarm.foo cf123456789 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "size",
					Description: `It has been deprecated from provider version 1.9.1. New field 'node_number' replaces it.`,
				},
				resource.Attribute{
					Name:        "node_number",
					Description: `The node number.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of VPC where the current cluster is located.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `The ID of VSwitch where the current cluster is located.`,
				},
				resource.Attribute{
					Name:        "slb_id",
					Description: `The ID of load balancer where the current cluster worker node is located.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of security group where the current cluster worker node is located.`,
				},
				resource.Attribute{
					Name:        "agent_version",
					Description: `The nodes agent version.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `The instance type of nodes.`,
				},
				resource.Attribute{
					Name:        "disk_category",
					Description: `The data disk category of nodes.`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `The data disk size of nodes.`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `List of cluster nodes. It contains several attributes to ` + "`" + `Block Nodes` + "`" + `. ### Block Nodes`,
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
					Description: `The private IP address of node.`,
				},
				resource.Attribute{
					Name:        "eip",
					Description: `The Elastic IP address of node.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The node current status. It is different with instance status. ## Import Swarm cluster can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cs_swarm.foo cf123456789 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_datahub_project",
			Category:         "Datahub Service",
			ShortDescription: `Provides a Alicloud datahub project resource.`,
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
			Type:             "alicloud_datahub_subscription",
			Category:         "Datahub Service",
			ShortDescription: `Provides a Alicloud datahub subscription resource.`,
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
					Description: `The ID of the datahub subscritpion as terraform resource. It was composed of project name, topic name and practical subscription ID generated from server side. Format to ` + "`" + `<project_name>:<topic_name>:<sub_id>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sub_id",
					Description: `The identidy of the subscritpion, generate from server side.`,
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
					Description: `The ID of the datahub subscritpion as terraform resource. It was composed of project name, topic name and practical subscription ID generated from server side. Format to ` + "`" + `<project_name>:<topic_name>:<sub_id>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sub_id",
					Description: `The identidy of the subscritpion, generate from server side.`,
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
			Type:             "alicloud_datahub_topic",
			Category:         "Datahub Service",
			ShortDescription: `Provides a Alicloud datahub topic resource.`,
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
			Type:             "alicloud_db_account",
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
					Description: `(Optional, Available in 1.57.1+) An KMS encrypts password used to a db account. If the ` + "`" + `password` + "`" + ` is filled in, this field will be ignored.`,
				},
				resource.Attribute{
					Name:        "kms_encryption_context",
					Description: `(Optional, MapString, Available in 1.57.1+) An KMS encryption context used to decrypt ` + "`" + `kms_encrypted_password` + "`" + ` before creating or updating a db account with ` + "`" + `kms_encrypted_password` + "`" + `. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when ` + "`" + `kms_encrypted_password` + "`" + ` is set.`,
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
					Description: `The current account resource ID. Composed of instance ID and account name with format ` + "`" + `<instance_id>:<name>` + "`" + `. ## Import RDS account can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_db_account.example "rm-12345:tf_account" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The current account resource ID. Composed of instance ID and account name with format ` + "`" + `<instance_id>:<name>` + "`" + `. ## Import RDS account can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_db_account.example "rm-12345:tf_account" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_db_account_privilege",
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
					Description: `The privilege of one account access database. Valid values: - ReadOnly: This value is only for MySQL, MariaDB and SQL Server - ReadWrite: This value is only for MySQL, MariaDB and SQL Server - DDLOnly: (Available in 1.64.0+) This value is only for MySQL and MariaDB - DMLOnly: (Available in 1.64.0+) This value is only for MySQL and MariaDB - DBOwner: (Available in 1.64.0+) This value is only for SQL Server and PostgreSQL. Default to "ReadOnly".`,
				},
				resource.Attribute{
					Name:        "db_names",
					Description: `(Required) List of specified database name. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The current account resource ID. Composed of instance ID, account name and privilege with format ` + "`" + `<instance_id>:<name>:<privilege>` + "`" + `. ## Import RDS account privilege can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_db_account_privilege.example "rm-12345:tf_account:ReadOnly" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The current account resource ID. Composed of instance ID, account name and privilege with format ` + "`" + `<instance_id>:<name>:<privilege>` + "`" + `. ## Import RDS account privilege can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_db_account_privilege.example "rm-12345:tf_account:ReadOnly" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_db_backup_policy",
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
					Name:        "backup_period",
					Description: `(Deprecated) It has been deprecated from version 1.69.0, and use field 'preferred_backup_period' instead.`,
				},
				resource.Attribute{
					Name:        "preferred_backup_period",
					Description: `(Optional, available in 1.69.0+) DB Instance backup period. Please set at least two days to ensure backing up at least twice a week. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]. Default to ["Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"].`,
				},
				resource.Attribute{
					Name:        "backup_time",
					Description: `(Deprecated) It has been deprecated from version 1.69.0, and use field 'preferred_backup_time' instead.`,
				},
				resource.Attribute{
					Name:        "preferred_backup_time",
					Description: `(Optional, available in 1.69.0+) DB instance backup time, in the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. Default to "02:00Z-03:00Z". China time is 8 hours behind it.`,
				},
				resource.Attribute{
					Name:        "retention_period",
					Description: `(Deprecated) It has been deprecated from version 1.69.0, and use field 'backup_retention_period' instead.`,
				},
				resource.Attribute{
					Name:        "backup_retention_period",
					Description: `(Optional, available in 1.69.0+) Instance backup retention days. Valid values: [7-730]. Default to 7. But mysql local disk is unlimited.`,
				},
				resource.Attribute{
					Name:        "log_backup",
					Description: `(Deprecated) It has been deprecated from version 1.68.0, and use field 'enable_backup_log' instead.`,
				},
				resource.Attribute{
					Name:        "enable_backup_log",
					Description: `(Optional, available in 1.68.0+) Whether to backup instance log. Valid values are ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + `, Default to ` + "`" + `true` + "`" + `. Note: The 'Basic Edition' category Rds instance does not support setting log backup. [What is Basic Edition](https://www.alibabacloud.com/help/doc-detail/48980.htm).`,
				},
				resource.Attribute{
					Name:        "log_retention_period",
					Description: `(Deprecated) It has been deprecated from version 1.69.0, and use field 'log_backup_retention_period' instead.`,
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
					Description: `The current backup policy resource ID. It is same as 'instance_id'. ## Import RDS backup policy can be imported using the id or instance id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_db_backup_policy.example "rm-12345678" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The current backup policy resource ID. It is same as 'instance_id'. ## Import RDS backup policy can be imported using the id or instance id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_db_backup_policy.example "rm-12345678" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_db_connection",
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
					Description: `The ip address of connection string. ## Import RDS connection can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_db_connection.example abc12345678 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `The ip address of connection string. ## Import RDS connection can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_db_connection.example abc12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_db_database",
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
					Description: `(ForceNew) Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The current database resource ID. Composed of instance ID and database name with format ` + "`" + `<instance_id>:<name>` + "`" + `. ## Import RDS database can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_db_database.example "rm-12345:tf_database" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The current database resource ID. Composed of instance ID and database name with format ` + "`" + `<instance_id>:<name>` + "`" + `. ## Import RDS database can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_db_database.example "rm-12345:tf_database" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_db_instance",
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
					Name:        "db_instance_storage_type",
					Description: `(Optional, Available in 1.68.0+) The storage type of the instance. Valid values: - local_ssd: specifies to use local SSDs. This value is recommended. - cloud_ssd: specifies to use standard SSDs. - cloud_essd: specifies to use enhanced SSDs (ESSDs). - cloud_essd2: specifies to use enhanced SSDs (ESSDs). - cloud_essd3: specifies to use enhanced SSDs (ESSDs).`,
				},
				resource.Attribute{
					Name:        "sql_collector_status",
					Description: `(Optional, Available in 1.70.0+) The sql collector status of the instance. Valid values are ` + "`" + `Enabled` + "`" + `, ` + "`" + `Disabled` + "`" + `, Default to ` + "`" + `Disabled` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sql_collector_config_value",
					Description: `(Optional, Available in 1.70.0+) The sql collector keep time of the instance. Valid values are ` + "`" + `30` + "`" + `, ` + "`" + `180` + "`" + `, ` + "`" + `365` + "`" + `, ` + "`" + `1095` + "`" + `, ` + "`" + `1825` + "`" + `, Default to ` + "`" + `30` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Optional) The name of DB instance. It a string of 2 to 256 characters.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `(Optional) Valid values are ` + "`" + `Prepaid` + "`" + `, ` + "`" + `Postpaid` + "`" + `, Default to ` + "`" + `Postpaid` + "`" + `. Currently, the resource only supports PostPaid to PrePaid.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) The duration that you will buy DB instance (in month). It is valid when instance_charge_type is ` + "`" + `PrePaid` + "`" + `. Valid values: [1~9], 12, 24, 36. Default to 1.`,
				},
				resource.Attribute{
					Name:        "monitoring_period",
					Description: `(Optional) The monitoring frequency in seconds. Valid values are 5, 60, 300. Defaults to 300.`,
				},
				resource.Attribute{
					Name:        "auto_renew",
					Description: `(Optional, Available in 1.34.0+) Whether to renewal a DB instance automatically or not. It is valid when instance_charge_type is ` + "`" + `PrePaid` + "`" + `. Default to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "auto_renew_period",
					Description: `(Optional, Available in 1.34.0+) Auto-renewal period of an instance, in the unit of the month. It is valid when instance_charge_type is ` + "`" + `PrePaid` + "`" + `. Valid value:[1~12], Default to 1.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(ForceNew) The Zone to launch the DB instance. From version 1.8.1, it supports multiple zone. If it is a multi-zone and ` + "`" + `vswitch_id` + "`" + ` is specified, the vswitch must in the one of them. The multiple zone ID can be retrieved by setting ` + "`" + `multi` + "`" + ` to "true" in the data source ` + "`" + `alicloud_zones` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(ForceNew) The virtual switch ID to launch DB instances in one VPC. If there are multiple vswitches, separate them with commas.`,
				},
				resource.Attribute{
					Name:        "security_ips",
					Description: `(Optional) List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).`,
				},
				resource.Attribute{
					Name:        "security_ip_mode",
					Description: `(Optional, Available in 1.62.1+) Valid values are ` + "`" + `normal` + "`" + `, ` + "`" + `safety` + "`" + `, Default to ` + "`" + `normal` + "`" + `. support ` + "`" + `safety` + "`" + ` switch to high security access mode`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `(Optional) Set of parameters needs to be set after DB instance was launched. Available parameters can refer to the latest docs [View database parameter templates](https://www.alibabacloud.com/help/doc-detail/26284.htm) .`,
				},
				resource.Attribute{
					Name:        "force_restart",
					Description: `(Optional, Available in 1.75.0+) Set it to true to make some parameter efficient when modifying them. Default to false.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string. - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string. Note: From 1.63.0, the tag key and value are case sensitive. Before that, they are not case sensitive.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Deprecated) It has been deprecated from 1.69.0 and use ` + "`" + `security_group_ids` + "`" + ` instead.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Optional, List(string), Available in 1.69.0+) The list IDs to join ECS Security Group. At most supports three security groups.`,
				},
				resource.Attribute{
					Name:        "maintain_time",
					Description: `(Optional, Available in 1.56.0+) Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)`,
				},
				resource.Attribute{
					Name:        "auto_upgrade_minor_version",
					Description: `(Optional, Available in 1.62.1+) The upgrade method to use. Valid values: - Auto: Instances are automatically upgraded to a higher minor version. - Manual: Instances are forcibly upgraded to a higher minor version when the current version is unpublished. Default to "Manual". See more [details and limitation](https://www.alibabacloud.com/help/doc-detail/123605.htm).`,
				},
				resource.Attribute{
					Name:        "ssl_action",
					Description: `(Optional, Available in v1.90.0+) Actions performed on SSL functions, Valid values: ` + "`" + `Open` + "`" + `: turn on SSL encryption; ` + "`" + `Close` + "`" + `: turn off SSL encryption; ` + "`" + `Update` + "`" + `: update SSL certificate. See more [engine and engineVersion limitation](https://www.alibabacloud.com/help/zh/doc-detail/26254.htm).`,
				},
				resource.Attribute{
					Name:        "tde_status",
					Description: `(Optional, ForceNew, Available in 1.90.0+) The TDE(Transparent Data Encryption) status. See more [engine and engineVersion limitation](https://www.alibabacloud.com/help/zh/doc-detail/26256.htm). ->`,
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
					Description: `RDS database connection string.`,
				},
				resource.Attribute{
					Name:        "ssl_status",
					Description: `Status of the SSL feature. ` + "`" + `Yes` + "`" + `: SSL is turned on; ` + "`" + `No` + "`" + `: SSL is turned off. ### Timeouts ->`,
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
					Description: `(Defaults to 20 mins) Used when terminating the db instance. ## Import RDS instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_db_instance.example rm-abc12345678 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `RDS database connection string.`,
				},
				resource.Attribute{
					Name:        "ssl_status",
					Description: `Status of the SSL feature. ` + "`" + `Yes` + "`" + `: SSL is turned on; ` + "`" + `No` + "`" + `: SSL is turned off. ### Timeouts ->`,
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
					Description: `(Defaults to 20 mins) Used when terminating the db instance. ## Import RDS instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_db_instance.example rm-abc12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_db_read_write_splitting_connection",
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
					Description: `Connection instance string. ## Import RDS read write splitting connection can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_db_read_write_splitting_connection.example abc12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Id of DB instance.`,
				},
				resource.Attribute{
					Name:        "connection_string",
					Description: `Connection instance string. ## Import RDS read write splitting connection can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_db_read_write_splitting_connection.example abc12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_db_readonly_instance",
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
					Description: `(Optional, Available in 1.68.0+) A mapping of tags to assign to the resource. - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string. - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string. ->`,
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
					Description: `RDS database connection string. ### Timeouts ->`,
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
					Description: `(Defaults to 20 mins) Used when terminating the db instance. ## Import RDS readonly instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_db_readonly_instance.example rm-abc12345678 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `RDS database connection string. ### Timeouts ->`,
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
					Description: `(Defaults to 20 mins) Used when terminating the db instance. ## Import RDS readonly instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_db_readonly_instance.example rm-abc12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ddosbgp_instance",
			Category:         "BGP-Line Anti-DDoS Resources",
			ShortDescription: `Provides a Alicloud Anti-DDoS Advanced(Ddosbgp) Instance Resource.`,
			Description:      ``,
			Keywords: []string{
				"bgp",
				"line",
				"anti",
				"ddos",
				"ddosbgp",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `(Required, ForceNew) Type of the instance. Valid values: Enterprise,Professional. Default to ` + "`" + `Enterprise` + "`" + ``,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the instance. This name can have a string of 1 to 63 characters.`,
				},
				resource.Attribute{
					Name:        "base_bandwidth",
					Description: `(Optional, ForceNew) Base defend bandwidth of the instance. Valid values: 20. The unit is Gbps. Default to ` + "`" + `20` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Required, ForceNew) Elastic defend bandwidth of the instance. This value must be larger than the base defend bandwidth. Valid values: 51,91,101,201,301. The unit is Gbps.`,
				},
				resource.Attribute{
					Name:        "ip_count",
					Description: `(Required, ForceNew) IP count of the instance. Valid values: 100.`,
				},
				resource.Attribute{
					Name:        "ip_type",
					Description: `(Required, ForceNew) IP version of the instance. Valid values: IPv4,IPv6.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional, ForceNew) The duration that you will buy Ddosbgp instance (in month). Valid values: [1~9], 12, 24, 36. Default to 12. At present, the provider does not support modify "period". ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the instance resource of Ddosbgp. ## Import Ddosbgp instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ddosbgp.example ddosbgp-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the instance resource of Ddosbgp. ## Import Ddosbgp instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ddosbgp.example ddosbgp-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ddoscoo_instance",
			Category:         "BGP-Line Anti-DDoS Resources",
			ShortDescription: `Provides a Alicloud BGP-line Anti-DDoS Pro(Ddoscoo) Instance Resource.`,
			Description:      ``,
			Keywords: []string{
				"bgp",
				"line",
				"anti",
				"ddos",
				"ddoscoo",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the instance. This name can have a string of 1 to 63 characters.`,
				},
				resource.Attribute{
					Name:        "base_bandwidth",
					Description: `(Required) Base defend bandwidth of the instance. Valid values: 30, 60, 100, 300, 400, 500, 600. The unit is Gbps. Only support upgrade.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Required) Elastic defend bandwidth of the instance. This value must be larger than the base defend bandwidth. Valid values: 30, 60, 100, 300, 400, 500, 600. The unit is Gbps. Only support upgrade.`,
				},
				resource.Attribute{
					Name:        "service_bandwidth",
					Description: `(Required) Business bandwidth of the instance. At leaset 100. Increased 100 per step, such as 100, 200, 300. The unit is Mbps. Only support upgrade.`,
				},
				resource.Attribute{
					Name:        "port_count",
					Description: `(Required) Port retransmission rule count of the instance. At least 50. Increase 5 per step, such as 55, 60, 65. Only support upgrade.`,
				},
				resource.Attribute{
					Name:        "domain_count",
					Description: `(Required) Domain retransmission rule count of the instance. At least 50. Increase 5 per step, such as 55, 60, 65. Only support upgrade.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional, ForceNew) The duration that you will buy Ddoscoo instance (in month). Valid values: [1~9], 12, 24, 36. Default to 1. At present, the provider does not support modify "period". ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the instance resource of Ddoscoo. ## Import Ddoscoo instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ddoscoo_instance.example ddoscoo-cn-123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the instance resource of Ddoscoo. ## Import Ddoscoo instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ddoscoo_instance.example ddoscoo-cn-123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ddoscoo_scheduler_rule",
			Category:         "BGP-Line Anti-DDoS Resources",
			ShortDescription: `Provides a Alicloud DdosCoo Scheduler Rule resource.`,
			Description:      ``,
			Keywords: []string{
				"bgp",
				"line",
				"anti",
				"ddos",
				"ddoscoo",
				"scheduler",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "rule_name",
					Description: `(Required, ForceNew) The name of the rule.`,
				},
				resource.Attribute{
					Name:        "rule_type",
					Description: `(Required) The rule type. Valid values: ` + "`" + `2` + "`" + `: tiered protection. ` + "`" + `3` + "`" + `: globalization acceleration. ` + "`" + `6` + "`" + `: Cloud product interaction.`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `(Required) The details of the common filter interaction rule, expressed as a JSON string. The structure is as follows: ` + "`" + `Type` + "`" + `: String type, required, the address format of the linkage resource. Valid values: ` + "`" + `A` + "`" + `: IP address. ` + "`" + `CNAME` + "`" + `: Domain name. ` + "`" + `Value` + "`" + `: String type, required, link address of resource. ` + "`" + `Priority` + "`" + `: the priority of the rule. This parameter is required and of Integer type. Valid values: 0~100 the larger the value, the higher the priority. ` + "`" + `ValueType` + "`" + `: Required. The type of the linked resource. It is an Integer. Valid values: ` + "`" + `1` + "`" + `: Anti-DDoS Pro. ` + "`" + `2` + "`" + `: (Tiered protection) cloud resource IP. ` + "`" + `3` + "`" + `: (sea acceleration) MCA IP address. ` + "`" + `6` + "`" + `: (Cloud product linkage) cloud resource IP. ` + "`" + `RegionId` + "`" + `: String type, optional (Required when ValueType is 2) the ID of the region.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional) The ID of the resource group to which the anti-DDoS pro instance belongs in resource management. By default, no value is specified, indicating that the domains in the default resource group are listed. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID of scheduler rule. The value is ` + "`" + `<rule_name>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cname",
					Description: `The cname is the traffic scheduler corresponding to rules. ## Import DdosCoo Scheduler Rule can be imported using the id or the rule name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ddoscoo_scheduler_rule.example fbb20dc77e8fc`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID of scheduler rule. The value is ` + "`" + `<rule_name>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cname",
					Description: `The cname is the traffic scheduler corresponding to rules. ## Import DdosCoo Scheduler Rule can be imported using the id or the rule name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ddoscoo_scheduler_rule.example fbb20dc77e8fc`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_disk",
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
					Description: `(Optional, ForceNew) Category of the disk. Valid values are ` + "`" + `cloud` + "`" + `, ` + "`" + `cloud_efficiency` + "`" + `, ` + "`" + `cloud_ssd` + "`" + `, ` + "`" + `cloud_essd` + "`" + `. Default is ` + "`" + `cloud_efficiency` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) The size of the disk in GiBs. When resize the disk, the new size must be greater than the former value, or you would get an error ` + "`" + `InvalidDiskSize.TooSmall` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) A snapshot to base the disk off of. If the disk size required by snapshot is greater than ` + "`" + `size` + "`" + `, the ` + "`" + `size` + "`" + ` will be ignored, conflict with ` + "`" + `encrypted` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "kms_key_id",
					Description: `(Optional, Available in 1.89.0+) The ID of the KMS key corresponding to the data disk, The specified parameter ` + "`" + `Encrypted` + "`" + ` must be ` + "`" + `true` + "`" + ` when KmsKeyId is not empty.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `(Optional, ForceNew) If true, the disk will be encrypted, conflict with ` + "`" + `snapshot_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "delete_auto_snapshot",
					Description: `(Optional Available in 1.53.0+) Indicates whether the automatic snapshot is deleted when the disk is released. Default value: false.`,
				},
				resource.Attribute{
					Name:        "delete_with_instance",
					Description: `(Optional Available in 1.53.0+) Indicates whether the disk is released together with the instance: Default value: false.`,
				},
				resource.Attribute{
					Name:        "enable_auto_snapshot",
					Description: `(Optional Available in 1.53.0+) Indicates whether to apply a created automatic snapshot policy to the disk. Default value: false.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(ForceNew, ForceNew, Available in 1.57.0+) The Id of resource group which the disk belongs. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the disk.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The disk status. ## Import Cloud disk can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_disk.example d-abc12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the disk.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The disk status. ## Import Cloud disk can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_disk.example d-abc12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_disk_attachment",
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
					Description: `(Deprecated) The device name has been deprecated, and when attaching disk, it will be allocated automatically by system according to default order from /dev/xvdb to /dev/xvdz. ## Attributes Reference The following attributes are exported:`,
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
			Type:             "alicloud_dms_enterprise_instance",
			Category:         "DMS-Enterprise",
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
					Description: `(Required) Instance alias, to help users quickly distinguish positioning.`,
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
					Description: `The instance status. ## Import DMS Enterprise can be imported using host and port, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_dms_enterprise_instance.example rm-uf648hgs7874xxxx.mysql.rds.aliyuncs.com:3306 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `The instance status. ## Import DMS Enterprise can be imported using host and port, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_dms_enterprise_instance.example rm-uf648hgs7874xxxx.mysql.rds.aliyuncs.com:3306 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_dms_enterprise_user",
			Category:         "DMS-Enterprise",
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
					Description: `The state of DMS Enterprise User. ## Import DMS Enterprise User can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_dms_enterprise_user.example 24356xxx ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `The state of DMS Enterprise User. ## Import DMS Enterprise User can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_dms_enterprise_user.example 24356xxx ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_dns",
			Category:         "DNS",
			ShortDescription: `Provides a DNS resource.`,
			Description:      ``,
			Keywords: []string{
				"dns",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Name of the domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix ` + "`" + `.sh` + "`" + ` and ` + "`" + `.tel` + "`" + ` are not supported.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Optional) Id of the group in which the domain will add. If not supplied, then use default group.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional, ForceNew, Available in 1.59.0+) The Id of resource group which the dns belongs. ## Attributes Reference The following attributes are exported:`,
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
					Name:        "name",
					Description: `The domain name.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `The group id of domain.`,
				},
				resource.Attribute{
					Name:        "dns_server",
					Description: `A list of the dns server name. ## Import DNS can be imported using the id or domain name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_dns.example "aliyun.com" ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "name",
					Description: `The domain name.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `The group id of domain.`,
				},
				resource.Attribute{
					Name:        "dns_server",
					Description: `A list of the dns server name. ## Import DNS can be imported using the id or domain name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_dns.example "aliyun.com" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_dns_domain",
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
					Description: `(Optional) Remarks information for your domain name.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string. - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string. ## Attributes Reference The following attributes are exported:`,
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
					Description: `A list of the dns server name. ## Import DNS domain can be imported using the id or domain name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_dns_domain.example aliyun.com ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `A list of the dns server name. ## Import DNS domain can be imported using the id or domain name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_dns_domain.example aliyun.com ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_dns_domain_attachment",
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
					Description: `Domain names bound to DNS instance. ## Import DNS domain attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_dns_domain_attachment.example dns-cn-v0h1ldjhxxx ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `This ID of this resource. The value is same as ` + "`" + `instance_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "domain_names",
					Description: `Domain names bound to DNS instance. ## Import DNS domain attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_dns_domain_attachment.example dns-cn-v0h1ldjhxxx ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_dns_group",
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
			Type:             "alicloud_dns_instance",
			Category:         "DNS",
			ShortDescription: `Provides a Alicloud DNS Instance resource.`,
			Description:      ``,
			Keywords: []string{
				"dns",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dns_security",
					Description: `(Required, ForceNew) DNS security level. Valid values: ` + "`" + `no` + "`" + `, ` + "`" + `basic` + "`" + `, ` + "`" + `advanced` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "domain_numbers",
					Description: `(Required, ForceNew) Number of domain names bound.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional, ForceNew) Creating a pre-paid instance, it must be set, the unit is month, please enter an integer multiple of 12 for annually paid products.`,
				},
				resource.Attribute{
					Name:        "renew_period",
					Description: `(Optional, ForceNew) Automatic renewal period, the unit is month. When setting RenewalStatus to AutoRenewal, it must be set.`,
				},
				resource.Attribute{
					Name:        "renewal_status",
					Description: `(Optional, ForceNew) Automatic renewal status. Valid values: ` + "`" + `AutoRenewal` + "`" + `, ` + "`" + `ManualRenewal` + "`" + `, default to ` + "`" + `ManualRenewal` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "version_code",
					Description: `(Required, ForceNew) Paid package version. Valid values: ` + "`" + `version_personal` + "`" + `, ` + "`" + `version_enterprise_basic` + "`" + `, ` + "`" + `version_enterprise_advanced` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the DNS instance.`,
				},
				resource.Attribute{
					Name:        "version_name",
					Description: `Paid package version name. ## Import DNS instance be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_dns_instance.example dns-cn-v0h1ldjhfff ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the DNS instance.`,
				},
				resource.Attribute{
					Name:        "version_name",
					Description: `Paid package version name. ## Import DNS instance be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_dns_instance.example dns-cn-v0h1ldjhfff ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_dns_record",
			Category:         "DNS",
			ShortDescription: `Provides a DNS Record resource.`,
			Description:      ``,
			Keywords: []string{
				"dns",
				"record",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix ` + "`" + `.sh` + "`" + ` and ` + "`" + `.tel` + "`" + ` are not supported.`,
				},
				resource.Attribute{
					Name:        "host_record",
					Description: `(Required) Host record for the domain record. This host_record can have at most 253 characters, and each part split with "." can have at most 63 characters, and must contain only alphanumeric characters or hyphens, such as "-",".","`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of domain record. Valid values are ` + "`" + `A` + "`" + `,` + "`" + `NS` + "`" + `,` + "`" + `MX` + "`" + `,` + "`" + `TXT` + "`" + `,` + "`" + `CNAME` + "`" + `,` + "`" + `SRV` + "`" + `,` + "`" + `AAAA` + "`" + `,` + "`" + `CAA` + "`" + `, ` + "`" + `REDIRECT_URL` + "`" + ` and ` + "`" + `FORWORD_URL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of domain record, When the ` + "`" + `type` + "`" + ` is ` + "`" + `MX` + "`" + `,` + "`" + `NS` + "`" + `,` + "`" + `CNAME` + "`" + `,` + "`" + `SRV` + "`" + `, the server will treat the ` + "`" + `value` + "`" + ` as a fully qualified domain name, so it's no need to add a ` + "`" + `.` + "`" + ` at the end.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The effective time of domain record. Its scope depends on the edition of the cloud resolution. Free is ` + "`" + `[600, 86400]` + "`" + `, Basic is ` + "`" + `[120, 86400]` + "`" + `, Standard is ` + "`" + `[60, 86400]` + "`" + `, Ultimate is ` + "`" + `[10, 86400]` + "`" + `, Exclusive is ` + "`" + `[1, 86400]` + "`" + `. Default value is ` + "`" + `600` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) The priority of domain record. Valid values are ` + "`" + `[1-10]` + "`" + `. When the ` + "`" + `type` + "`" + ` is ` + "`" + `MX` + "`" + `, this parameter is required.`,
				},
				resource.Attribute{
					Name:        "routing",
					Description: `(Optional) The resolution line of domain record. Valid values are ` + "`" + `default` + "`" + `, ` + "`" + `telecom` + "`" + `, ` + "`" + `unicom` + "`" + `, ` + "`" + `mobile` + "`" + `, ` + "`" + `oversea` + "`" + `, ` + "`" + `edu` + "`" + `, ` + "`" + `drpeng` + "`" + `, ` + "`" + `btvn` + "`" + `, .etc. When the ` + "`" + `type` + "`" + ` is ` + "`" + `FORWORD_URL` + "`" + `, this parameter must be ` + "`" + `default` + "`" + `. Default value is ` + "`" + `default` + "`" + `. For checking all resolution lines enumeration please visit [Alibaba Cloud DNS doc](https://www.alibabacloud.com/help/doc-detail/34339.htm) or using alicloud_dns_resolution_lines in data source to get the value. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The record id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The record domain name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The record type.`,
				},
				resource.Attribute{
					Name:        "host_record",
					Description: `The host record of record.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The record value.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `The record effective time.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The record priority.`,
				},
				resource.Attribute{
					Name:        "routing",
					Description: `The record resolution line.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The record status. ` + "`" + `Enable` + "`" + ` or ` + "`" + `Disable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "Locked",
					Description: `The record locked state. ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `. ## Import RDS record can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_dns_record.example abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The record id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The record domain name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The record type.`,
				},
				resource.Attribute{
					Name:        "host_record",
					Description: `The host record of record.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The record value.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `The record effective time.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The record priority.`,
				},
				resource.Attribute{
					Name:        "routing",
					Description: `The record resolution line.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The record status. ` + "`" + `Enable` + "`" + ` or ` + "`" + `Disable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "Locked",
					Description: `The record locked state. ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `. ## Import RDS record can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_dns_record.example abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_drds_instance",
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
					Description: `(Optional, ForceNew) The Zone to launch the DRDS instance.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `(Optional, ForceNew) Valid values are ` + "`" + `PrePaid` + "`" + `, ` + "`" + `PostPaid` + "`" + `, Default to ` + "`" + `PostPaid` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Optional, ForceNew) The VSwitch ID to launch in.`,
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
					Description: `The DRDS instance ID. ## Import Distributed Relational Database Service (DRDS) can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_drds_instance.example drds-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The DRDS instance ID. ## Import Distributed Relational Database Service (DRDS) can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_drds_instance.example drds-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_eci_image_cache",
			Category:         "Elastic Container Instance (ECI)",
			ShortDescription: `Provides an Alicloud ECI Image Cache resource.`,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"container",
				"instance",
				"eci",
				"image",
				"cache",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "image_cache_name",
					Description: `(Required, ForceNew) The name of the image cache.`,
				},
				resource.Attribute{
					Name:        "images",
					Description: `(Required, ForceNew) The images to be cached. The image name must be versioned.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Required, ForceNew) The ID of the security group. You do not need to specify the same security group as the container group.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Required, ForceNew) The ID of the VSwitch. You do not need to specify the same VSwitch as the container group.`,
				},
				resource.Attribute{
					Name:        "eip_instance_id",
					Description: `(Optional, ForceNew) The instance ID of the Elastic IP Address (EIP). If you want to pull images from the Internet, you must specify an EIP to make sure that the container group can access the Internet. You can also configure the network address translation (NAT) gateway. We recommend that you configure the NAT gateway for the Internet access. Refer to [Public Network Access Method](https://help.aliyun.com/document_detail/99146.html)`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional, ForceNew) The ID of the resource group.`,
				},
				resource.Attribute{
					Name:        "retention_days",
					Description: `(Optional, ForceNew) The retention days of the image cache. Once the image cache expires, it will be cleared. By default, the image cache never expires. Note: The image cache that fails to be created is retained for only one day.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Optional, ForceNew) The zone id to cache image.`,
				},
				resource.Attribute{
					Name:        "image_registry_credential",
					Description: `(Optional, ForceNew) The Image Registry parameters about the image to be cached. ### Block image_registry_credential`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `(Optional) The address of Image Registry without ` + "`" + `http://` + "`" + ` or ` + "`" + `https://` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `(Optional) The user name of Image Registry.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The password of the Image Registry. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "container_group_id",
					Description: `The ID of the container group job that is used to create the image cache.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "container_group_id",
					Description: `The ID of the container group job that is used to create the image cache.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_edas_application",
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
					Description: `The ` + "`" + `key` + "`" + ` of the resource supplied above. The value is formulated as ` + "`" + `app_Id` + "`" + `. ## Import EDAS application can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_edas_application.app app_Id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `key` + "`" + ` of the resource supplied above. The value is formulated as ` + "`" + `app_Id` + "`" + `. ## Import EDAS application can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_edas_application.app app_Id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_edas_application_deployment",
			Category:         "EDAS",
			ShortDescription: `Creates an EDAS application deployment resource.`,
			Description:      ``,
			Keywords: []string{
				"edas",
				"application",
				"deployment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "app_id",
					Description: `(Required, ForceNew) The ID of the application that you want to deploy.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required, ForceNew) The ID of the instance group where the application is going to be deployed. Set this parameter to all if you want to deploy the application to all groups.`,
				},
				resource.Attribute{
					Name:        "package_version",
					Description: `(Optional, ForceNew) The version of the application that you want to deploy. It must be unique for every application. The length cannot exceed 64 characters. We recommended you to use a timestamp.`,
				},
				resource.Attribute{
					Name:        "war_url",
					Description: `(Required, ForceNew) The address to store the uploaded web application (WAR) package for application deployment. This parameter is required when the deployType parameter is set as url. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `key` + "`" + ` of the resource supplied above. The value is formulated as ` + "`" + `<app_id>:<package_version>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "last_package_version",
					Description: `Last package version deployed.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `key` + "`" + ` of the resource supplied above. The value is formulated as ` + "`" + `<app_id>:<package_version>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "last_package_version",
					Description: `Last package version deployed.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_edas_application_scale",
			Category:         "EDAS",
			ShortDescription: `This operation is provided to scale out an EDAS application.`,
			Description:      ``,
			Keywords: []string{
				"edas",
				"application",
				"scale",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "app_id",
					Description: `(Required, ForceNew) The ID of the application that you want to deploy.`,
				},
				resource.Attribute{
					Name:        "deploy_group",
					Description: `(Required, ForceNew) The ID of the instance group to which you want to add ECS instances to scale out the application.`,
				},
				resource.Attribute{
					Name:        "ecu_info",
					Description: `(Required, ForceNew) The IDs of the Elastic Compute Unit (ECU) where you want to deploy the application. Type: List.`,
				},
				resource.Attribute{
					Name:        "force_status",
					Description: `(Optional) This parameter specifies whether to forcibly remove an ECS instance where the application is deployed. It is set as true only after the ECS instance expires. In normal cases, this parameter do not need to be specified. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `key` + "`" + ` of the resource supplied above. The value is formulated as ` + "`" + `<app_id>:<ecu1,ecu2>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ecc_info",
					Description: `The ecc information of the resource supplied above. The value is formulated as ` + "`" + `<ecc1,ecc2>` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `key` + "`" + ` of the resource supplied above. The value is formulated as ` + "`" + `<app_id>:<ecu1,ecu2>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ecc_info",
					Description: `The ecc information of the resource supplied above. The value is formulated as ` + "`" + `<ecc1,ecc2>` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_edas_cluster",
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
					Description: `The ` + "`" + `key` + "`" + ` of the resource supplied above. The value is formulated as ` + "`" + `<cluster_id>` + "`" + `. ## Import EDAS cluster can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_edas_cluster.cluster cluster_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `key` + "`" + ` of the resource supplied above. The value is formulated as ` + "`" + `<cluster_id>` + "`" + `. ## Import EDAS cluster can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_edas_cluster.cluster cluster_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_edas_deploy_group",
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
					Description: `The type of the instance group that you want to create. Valid values: 0: Default group. 1: Phased release is disabled for traffic management. 2: Phased release is enabled for traffic management. ## Import EDAS deploy group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_edas_deploy_group.group app_id:group_name:group_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `key` + "`" + ` of the resource supplied above. The value is formulated as ` + "`" + `<app_id>:<group_name>:<group_id>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "group_type",
					Description: `The type of the instance group that you want to create. Valid values: 0: Default group. 1: Phased release is disabled for traffic management. 2: Phased release is enabled for traffic management. ## Import EDAS deploy group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_edas_deploy_group.group app_id:group_name:group_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_edas_instance_cluster_attachment",
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
			Type:             "alicloud_edas_slb_attachment",
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
					Description: `(Required, ForceNew) The ID of the applicaton to which you want to bind an SLB instance.`,
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
					Description: `Running Status of SLB instance. Inactive：The instance is stopped, and listener will not monitor and foward traffic. Active：The instance is running. After the instance is created, the default state is active. Locked：The instance is locked, the instance has been owed or locked by Alibaba Cloud. Expired: The instance has expired.`,
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
					Description: `Running Status of SLB instance. Inactive：The instance is stopped, and listener will not monitor and foward traffic. Active：The instance is running. After the instance is created, the default state is active. Locked：The instance is locked, the instance has been owed or locked by Alibaba Cloud. Expired: The instance has expired.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `VPC related vswitch ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_eip",
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
					Name:        "internet_charge_type",
					Description: `(Optional, ForceNew) Internet charge type of the EIP, Valid values are ` + "`" + `PayByBandwidth` + "`" + `, ` + "`" + `PayByTraffic` + "`" + `. Default to ` + "`" + `PayByBandwidth` + "`" + `. From version ` + "`" + `1.7.1` + "`" + `, default to ` + "`" + `PayByTraffic` + "`" + `. It is only PayByBandwidth when ` + "`" + `instance_charge_type` + "`" + ` is PrePaid.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `(Optional, ForceNew) Elastic IP instance charge type. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid".`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional, ForceNew) The duration that you will buy the resource, in month. It is valid when ` + "`" + `instance_charge_type` + "`" + ` is ` + "`" + `PrePaid` + "`" + `. Default to 1. Valid values: [1-9, 12, 24, 36]. At present, the provider does not support modify "period" and you can do that via web console.`,
				},
				resource.Attribute{
					Name:        "isp",
					Description: `(Optional, ForceNew, Available in 1.47.0+) The line type of the Elastic IP instance. Default to ` + "`" + `BGP` + "`" + `. Other type of the isp need to open a whitelist.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Available in v1.55.3+) A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional, ForceNew, Available in 1.58.0+) The Id of resource group which the eip belongs. ## Attributes Reference The following attributes are exported:`,
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
					Name:        "internet_charge_type",
					Description: `The EIP internet charge type.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The EIP current status.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The elastic ip address ## Import Elastic IP address can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_eip.example eip-abc12345678 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "internet_charge_type",
					Description: `The EIP internet charge type.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The EIP current status.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The elastic ip address ## Import Elastic IP address can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_eip.example eip-abc12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_eip_association",
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
					Description: `(Optional, ForceNew, Available in 1.46.0+) The type of cloud product that the eip instance to bind.`,
				},
				resource.Attribute{
					Name:        "private_ip_address",
					Description: `(Optional, ForceNew, Available in 1.52.2+) The private IP address in the network segment of the vswitch which has been assigned. ## Attributes Reference The following attributes are exported:`,
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
			Type:             "alicloud_elasticsearch_instance",
			Category:         "Elasticsearch",
			ShortDescription: `Provides a Alicloud Elasticsearch instance resource.`,
			Description:      ``,
			Keywords: []string{
				"elasticsearch",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of instance. It a string of 0 to 30 characters.`,
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
					Description: `(Required, ForceNew) Elasticsearch version. Supported values: ` + "`" + `5.5.3_with_X-Pack` + "`" + `, ` + "`" + `6.3_with_X-Pack` + "`" + ` and ` + "`" + `6.7_with_X-Pack` + "`" + `.`,
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
					Name:        "zone_count",
					Description: `(Optional, Available in 1.44.0+) The Multi-AZ supported for Elasticsearch, between 1 and 3. The ` + "`" + `data_node_amount` + "`" + ` value must be an integral multiple of the ` + "`" + `zone_count` + "`" + ` value.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Available in v1.73.0+) A mapping of tags to assign to the resource. - key: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:". It cannot contain "http://" and "https://". It cannot be a null string. - value: It can be up to 128 characters in length. It cannot contain "http://" and "https://". It can be a null string.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional, ForceNew, Computed Available in v1.86.0+) The Id of resource group which the Elasticsearch instance belongs. ### Timeouts ->`,
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
					Description: `The Elasticsearch instance status. Includes ` + "`" + `active` + "`" + `, ` + "`" + `activating` + "`" + `, ` + "`" + `inactive` + "`" + `. Some operations are denied when status is not ` + "`" + `active` + "`" + `. ## Import Elasticsearch can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_elasticsearch_instance.example es-cn-abcde123456 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `The Elasticsearch instance status. Includes ` + "`" + `active` + "`" + `, ` + "`" + `activating` + "`" + `, ` + "`" + `inactive` + "`" + `. Some operations are denied when status is not ` + "`" + `active` + "`" + `. ## Import Elasticsearch can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_elasticsearch_instance.example es-cn-abcde123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_emr_cluster",
			Category:         "E-MapReduce",
			ShortDescription: `Provides a EMR Cluster resource.`,
			Description:      ``,
			Keywords: []string{
				"e",
				"mapreduce",
				"emr",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of emr cluster. The name length must be less than 64. Supported characters: chinese character, english character, number, "-", "_".`,
				},
				resource.Attribute{
					Name:        "emr_ver",
					Description: `(Required, ForceNew) EMR Version, e.g. EMR-3.22.0. You can find the all valid EMR Version in emr web console.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `(Required, ForceNew) EMR Cluster Type, e.g. HADOOP, KAFKA, DRUID, GATEWAY etc. You can find all valid EMR cluster type in emr web console. Supported 'GATEWAY' available in 1.61.0+.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `(Required, ForceNew) Charge Type for this cluster. Supported value: PostPaid or PrePaid. Default value: PostPaid.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Required, ForceNew) Zone ID, e.g. cn-huhehaote-a`,
				},
				resource.Attribute{
					Name:        "host_group",
					Description: `(Optional) Groups of Host, You can specify MASTER as a group, CORE as a group (just like the above example).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Available in v1.67.0+) A mapping of tags to assign to the resource. #### Block host_group The host_group mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "host_group_name",
					Description: `(Required, ForceNew) host group name.`,
				},
				resource.Attribute{
					Name:        "host_group_type",
					Description: `(Required) host group type, supported value: MASTER, CORE or TASK, supported 'GATEWAY' available in 1.61.0+.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `(Optional) Charge Type for this group of hosts: PostPaid or PrePaid. If this is not specified, charge type will follow global charge_type value.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) If charge type is PrePaid, this should be specified, unit is month. Supported value: 1、2、3、4、5、6、7、8、9、12、24、36.`,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `(Required) Host number in this group.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Required) Host Ecs instance type.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `(Required) Data disk type. Supported value: cloud,cloud_efficiency,cloud_ssd,local_disk,cloud_essd.`,
				},
				resource.Attribute{
					Name:        "disk_capacity",
					Description: `(Required) Data disk capacity.`,
				},
				resource.Attribute{
					Name:        "disk_count",
					Description: `(Required) Data disk count.`,
				},
				resource.Attribute{
					Name:        "sys_disk_type",
					Description: `(Required) System disk type. Supported value: cloud,cloud_efficiency,cloud_ssd,cloud_essd.`,
				},
				resource.Attribute{
					Name:        "sys_disk_capacity",
					Description: `(Required) System disk capacity.`,
				},
				resource.Attribute{
					Name:        "auto_renew",
					Description: `(Optional) Auto renew for prepaid, true of false. Default is false.`,
				},
				resource.Attribute{
					Name:        "instance_list",
					Description: `(Optional) Instance list for cluster scale down. This value follows the json format, e.g. ["instance_id1","instance_id2"]. escape character for " is \". #### Block bootstrap_action The bootstrap_action mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, Available in 1.71.2+) bootstrap action name.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional, Available in 1.71.2+) bootstrap action path, e.g. "oss://bucket/path".`,
				},
				resource.Attribute{
					Name:        "arg",
					Description: `(Optional, Available in 1.71.2+) bootstrap action args, e.g. "--a=b". #### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 20 mins) Used when creating the cluster (until it reaches the initial ` + "`" + `IDLE` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 10 mins) Used when terminating the instance. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The cluster ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The cluster ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ess_alarm",
			Category:         "Auto Scaling (ESS)",
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
					Description: `(Optional, Available in 1.48.0+) Whether to enable specific ess alarm. Default to true.`,
				},
				resource.Attribute{
					Name:        "alarm_actions",
					Description: `(Required) The list of actions to execute when this alarm transition into an ALARM state. Each action is specified as ess scaling rule ari.`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `(Required, ForceNew) The scaling group associated with this alarm, the 'ForceNew' attribute is available in 1.56.0+.`,
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
					Description: `The state of specified alarm. ## Import Ess alarm can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ess_alarm.example asg-2ze500_045efffe-4d05 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id for ess alarm.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of specified alarm. ## Import Ess alarm can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ess_alarm.example asg-2ze500_045efffe-4d05 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ess_attachment",
			Category:         "Auto Scaling (ESS)",
			ShortDescription: `Provides a ESS Attachment resource to attach or remove ECS instances.`,
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
					Description: `Whether to delete "AutoCreated" ECS instances. ## Import ESS attachment can be imported using the id or scaling group id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ess_attachment.example asg-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `Whether to delete "AutoCreated" ECS instances. ## Import ESS attachment can be imported using the id or scaling group id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ess_attachment.example asg-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ess_notification",
			Category:         "Auto Scaling (ESS)",
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
					Description: `(Required, ForceNew) The Alibaba Cloud Resource Name (ARN) of the notification object, The value must be in ` + "`" + `acs:ess:{region}:{account-id}:{resource-relative-id}` + "`" + ` format.`,
				},
				resource.Attribute{
					Name:        "notification_types",
					Description: `(Required) The notification types of Auto Scaling events and resource changes. Supported notification types: 'AUTOSCALING:SCALE_OUT_SUCCESS', 'AUTOSCALING:SCALE_IN_SUCCESS', 'AUTOSCALING:SCALE_OUT_ERROR', 'AUTOSCALING:SCALE_IN_ERROR', 'AUTOSCALING:SCALE_REJECT', 'AUTOSCALING:SCALE_OUT_START', 'AUTOSCALING:SCALE_IN_START', 'AUTOSCALING:SCHEDULE_TASK_EXPIRING'. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of notification resource, which is composed of 'scaling_group_id' and 'notification_arn' in the format of '<scaling_group_id>:<notification_arn>'. ## Import Ess notification can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ess_notification.example 'scaling_group_id:notification_arn' ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of notification resource, which is composed of 'scaling_group_id' and 'notification_arn' in the format of '<scaling_group_id>:<notification_arn>'. ## Import Ess notification can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ess_notification.example 'scaling_group_id:notification_arn' ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ess_scaling_configuration",
			Category:         "Auto Scaling (ESS)",
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
					Description: `(Required) ID of an image file, indicating the image resource selected when an instance is enabled.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Optional) Resource type of an ECS instance.`,
				},
				resource.Attribute{
					Name:        "instance_types",
					Description: `(Optional, Available in 1.46.0+) Resource types of an ECS instance.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Optional) Name of an ECS instance. Default to "ESS-Instance". It is valid from version 1.7.1.`,
				},
				resource.Attribute{
					Name:        "io_optimized",
					Description: `(Deprecated) It has been deprecated on instance resource. All the launched alicloud instances will be I/O optimized.`,
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
					Description: `(Optional, Available in 1.43.0+) List IDs of the security group used to create new instances. It is conflict with ` + "`" + `security_group_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scaling_configuration_name",
					Description: `(Optional) Name shown for the scheduled task. which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores ` + "`" + `_` + "`" + `, hypens ` + "`" + `-` + "`" + `, and decimal point ` + "`" + `.` + "`" + `. If this parameter value is not specified, the default value is ScalingConfigurationId.`,
				},
				resource.Attribute{
					Name:        "internet_charge_type",
					Description: `(Optional) Network billing type, Values: PayByBandwidth or PayByTraffic. Default to ` + "`" + `PayByBandwidth` + "`" + `.`,
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
					Name:        "system_disk_category",
					Description: `(Optional) Category of the system disk. The parameter value options are ` + "`" + `ephemeral_ssd` + "`" + `, ` + "`" + `cloud_efficiency` + "`" + `, ` + "`" + `cloud_ssd` + "`" + `, ` + "`" + `cloud_essd` + "`" + ` and ` + "`" + `cloud` + "`" + `. ` + "`" + `cloud` + "`" + ` only is used to some no I/O optimized instance. Default to ` + "`" + `cloud_efficiency` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "system_disk_size",
					Description: `(Optional) Size of system disk, in GiB. Optional values: cloud: 20-500, cloud_efficiency: 20-500, cloud_ssd: 20-500, ephemeral_ssd: 20-500 The default value is max{40, ImageSize}. If this parameter is set, the system disk size must be greater than or equal to max{40, ImageSize}.`,
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
					Description: `(Optional) Instance RAM role name. The name is provided and maintained by RAM. You can use ` + "`" + `alicloud_ram_role` + "`" + ` to create a new one.`,
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
					Name:        "instance_ids",
					Description: `(Deprecated) It has been deprecated from version 1.6.0. New resource ` + "`" + `alicloud_ess_attachment` + "`" + ` replaces it.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. It will be applied for ECS instances finally. - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "http://", or "https://". It cannot be a null string. - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "http://", or "https://" It can be a null string.`,
				},
				resource.Attribute{
					Name:        "override",
					Description: `(Optional, Available in 1.46.0+) Indicates whether to overwrite the existing data. Default to false.`,
				},
				resource.Attribute{
					Name:        "password_inherit",
					Description: `(Optional, Available in 1.62.0+) Specifies whether to use the password that is predefined in the image. If the PasswordInherit parameter is set to true, the ` + "`" + `password` + "`" + ` and ` + "`" + `kms_encrypted_password` + "`" + ` will be ignored. You must ensure that the selected image has a password configured.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional, ForceNew, Available in 1.60.0+) The password of the ECS instance. The password must be 8 to 30 characters in length. It must contains at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include ` + "`" + `() ~!@#$%^&`,
				},
				resource.Attribute{
					Name:        "kms_encrypted_password",
					Description: `(Optional, ForceNew, Available in 1.60.0+) An KMS encrypts password used to a db account. If the ` + "`" + `password` + "`" + ` is filled in, this field will be ignored.`,
				},
				resource.Attribute{
					Name:        "kms_encryption_context",
					Description: `(Optional, MapString, Available in 1.60.0+) An KMS encryption context used to decrypt ` + "`" + `kms_encrypted_password` + "`" + ` before creating or updating a db account with ` + "`" + `kms_encrypted_password` + "`" + `. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when ` + "`" + `kms_encrypted_password` + "`" + ` is set. ->`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) Size of data disk, in GB. The value ranges [5,2000] for a cloud disk, [5,1024] for an ephemeral disk, [5,800] for an ephemeral_ssd disk, [20,32768] for cloud_efficiency, cloud_ssd, cloud_essd disk.`,
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
					Description: `(Optional) Whether to delete data disks attached on ecs when release ecs instance. Optional value: ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `, default to ` + "`" + `true` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The scaling configuration ID. ## Import ESS scaling configuration can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ess_scaling_configuration.example asg-abc123456 ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The scaling configuration ID. ## Import ESS scaling configuration can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ess_scaling_configuration.example asg-abc123456 ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ess_scaling_group",
			Category:         "Auto Scaling (ESS)",
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
					Description: `(Required) Minimum number of ECS instances in the scaling group. Value range: [0, 1000].`,
				},
				resource.Attribute{
					Name:        "max_size",
					Description: `(Required) Maximum number of ECS instances in the scaling group. Value range: [0, 1000].`,
				},
				resource.Attribute{
					Name:        "desired_capacity",
					Description: `(Optional,Available in 1.76.0+) Expected number of ECS instances in the scaling group. Value range: [min_size, max_size].`,
				},
				resource.Attribute{
					Name:        "scaling_group_name",
					Description: `(Optional) Name shown for the scaling group, which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain numbers, underscores ` + "`" + `_` + "`" + `, hyphens ` + "`" + `-` + "`" + `, and decimal points ` + "`" + `.` + "`" + `. If this parameter is not specified, the default value is ScalingGroupId.`,
				},
				resource.Attribute{
					Name:        "default_cooldown",
					Description: `(Optional) Default cool-down time (in seconds) of the scaling group. Value range: [0, 86400]. The default value is 300s.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Deprecated) It has been deprecated from version 1.7.1 and new field 'vswitch_ids' replaces it.`,
				},
				resource.Attribute{
					Name:        "vswitch_ids",
					Description: `(Optional) List of virtual switch IDs in which the ecs instances to be launched.`,
				},
				resource.Attribute{
					Name:        "removal_policies",
					Description: `(Optional) RemovalPolicy is used to select the ECS instances you want to remove from the scaling group when multiple candidates for removal exist. Optional values: - OldestInstance: removes the ECS instance that is added to the scaling group at the earliest point in time. - NewestInstance: removes the ECS instance that is added to the scaling group at the latest point in time. - OldestScalingConfiguration: removes the ECS instance that is created based on the earliest scaling configuration. - Default values: Default value of RemovalPolicy.1: OldestScalingConfiguration. Default value of RemovalPolicy.2: OldestInstance.`,
				},
				resource.Attribute{
					Name:        "db_instance_ids",
					Description: `(Optional) If an RDS instance is specified in the scaling group, the scaling group automatically attaches the Intranet IP addresses of its ECS instances to the RDS access whitelist. - The specified RDS instance must be in running status. - The specified RDS instance’s whitelist must have room for more IP addresses.`,
				},
				resource.Attribute{
					Name:        "loadbalancer_ids",
					Description: `(Optional) If a Server Load Balancer instance is specified in the scaling group, the scaling group automatically attaches its ECS instances to the Server Load Balancer instance. - The Server Load Balancer instance must be enabled. - At least one listener must be configured for each Server Load Balancer and it HealthCheck must be on. Otherwise, creation will fail (it may be useful to add a ` + "`" + `depends_on` + "`" + ` argument targeting your ` + "`" + `alicloud_slb_listener` + "`" + ` in order to make sure the listener with its HealthCheck configuration is ready before creating your scaling group). - The Server Load Balancer instance attached with VPC-type ECS instances cannot be attached to the scaling group. - The default weight of an ECS instance attached to the Server Load Balancer instance is 50.`,
				},
				resource.Attribute{
					Name:        "multi_az_policy",
					Description: `(Optional, ForceNew) Multi-AZ scaling group ECS instance expansion and contraction strategy. PRIORITY, BALANCE or COST_OPTIMIZED(Available in 1.54.0+).`,
				},
				resource.Attribute{
					Name:        "on_demand_base_capacity",
					Description: `(Optional, Available in 1.54.0+) The minimum amount of the Auto Scaling group's capacity that must be fulfilled by On-Demand Instances. This base portion is provisioned first as your group scales.`,
				},
				resource.Attribute{
					Name:        "on_demand_percentage_above_base_capacity",
					Description: `(Optional, Available in 1.54.0+) Controls the percentages of On-Demand Instances and Spot Instances for your additional capacity beyond OnDemandBaseCapacity.`,
				},
				resource.Attribute{
					Name:        "spot_instance_pools",
					Description: `(Optional, Available in 1.54.0+) The number of Spot pools to use to allocate your Spot capacity. The Spot pools is composed of instance types of lowest price.`,
				},
				resource.Attribute{
					Name:        "spot_instance_remedy",
					Description: `(Optional, Available in 1.54.0+) Whether to replace spot instances with newly created spot/onDemand instance when receive a spot recycling message. ->`,
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
					Description: `The vswitches id in which the ECS instance launched. ## Import ESS scaling group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ess_scaling_group.example asg-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `The vswitches id in which the ECS instance launched. ## Import ESS scaling group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ess_scaling_group.example asg-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ess_lifecycle_hook",
			Category:         "Auto Scaling (ESS)",
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
					Description: `The arn of notification target. ## Import Ess lifecycle hook can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ess_lifecycle_hook.example ash-l12345 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `The arn of notification target. ## Import Ess lifecycle hook can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ess_lifecycle_hook.example ash-l12345 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ess_scaling_rule",
			Category:         "Auto Scaling (ESS)",
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
					Description: `(Optional) Adjusted value of a scaling rule. Value range: - QuantityChangeInCapacity：(0, 500] U (-500, 0] - PercentChangeInCapacity：[0, 10000] U [-100, 0] - TotalCapacity：[0, 1000]`,
				},
				resource.Attribute{
					Name:        "scaling_rule_name",
					Description: `(Optional) Name shown for the scaling rule, which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores ` + "`" + `_` + "`" + `, hypens ` + "`" + `-` + "`" + `, and decimal point ` + "`" + `.` + "`" + `. If this parameter value is not specified, the default value is scaling rule id.`,
				},
				resource.Attribute{
					Name:        "cooldown",
					Description: `(Optional) Cool-down time of a scaling rule. Value range: [0, 86,400], in seconds. The default value is empty，if not set, the return value will be 0, which is the default value of integer.`,
				},
				resource.Attribute{
					Name:        "scaling_rule_type",
					Description: `(Optional, Available in 1.58.0+) The scaling rule type, either "SimpleScalingRule", "TargetTrackingScalingRule", "StepScalingRule". Default to "SimpleScalingRule".`,
				},
				resource.Attribute{
					Name:        "estimated_instance_warmup",
					Description: `(Optional, Available in 1.58.0+) The estimated time, in seconds, until a newly launched instance will contribute CloudMonitor metrics. Default to 300.`,
				},
				resource.Attribute{
					Name:        "metric_name",
					Description: `(Optional, Available in 1.58.0+) A CloudMonitor metric name.`,
				},
				resource.Attribute{
					Name:        "target_value",
					Description: `(Optional, Available in 1.58.0+) The target value for the metric.`,
				},
				resource.Attribute{
					Name:        "disable_scale_in",
					Description: `(Optional, Available in 1.58.0+) Indicates whether scale in by the target tracking policy is disabled. Default to false.`,
				},
				resource.Attribute{
					Name:        "step_adjustment",
					Description: `(Optional, Available in 1.58.0+) Steps for StepScalingRule. See [Block stepAdjustment](#block-stepAdjustment) below for details. ## Block stepAdjustment The stepAdjustment mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "metric_interval_lower_bound",
					Description: `(Optional) The lower bound of step.`,
				},
				resource.Attribute{
					Name:        "metric_interval_upper_bound",
					Description: `(Optional) The upper bound of step.`,
				},
				resource.Attribute{
					Name:        "scaling_adjustment",
					Description: `(Optional) The adjust value of step. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The scaling rule ID. ## Import ESS scaling rule can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ess_scaling_rule.example abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The scaling rule ID. ## Import ESS scaling rule can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ess_scaling_rule.example abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ess_scalinggroup_vserver_groups",
			Category:         "Auto Scaling (ESS)",
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
					Description: `(Optional, Available in 1.64.0+) If instances of scaling group are attached/removed from slb backend server when attach/detach vserver group from scaling group. Default to true. ## Block vserver_group the vserver_group supports the following:`,
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
					Description: `(Required, ForceNew) The ESS vserver groups attachment resource ID. ## Import ESS vserver groups can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ess_vserver_groups.example abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required, ForceNew) The ESS vserver groups attachment resource ID. ## Import ESS vserver groups can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ess_vserver_groups.example abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ess_schedule",
			Category:         "Auto Scaling (ESS)",
			ShortDescription: `Provides a ESS schedule resource.`,
			Description:      ``,
			Keywords: []string{
				"auto",
				"scaling",
				"ess",
				"schedule",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ess_scheduled_task",
			Category:         "Auto Scaling (ESS)",
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
					Description: `(Optional) The operation to be performed when a scheduled task is triggered. Enter the unique identifier of a scaling rule.`,
				},
				resource.Attribute{
					Name:        "scheduled_task_name",
					Description: `(Optional) Display name of the scheduled task, which must be 2-40 characters (English or Chinese) long.`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `(Optional, Available in 1.90.0+) The ID of the scaling group where the number of instances is modified when the scheduled task is triggered. After the ` + "`" + `ScalingGroupId` + "`" + ` parameter is specified, the scaling method of the scheduled task is to specify the number of instances in a scaling group. You must specify at least one of the following parameters: ` + "`" + `MinValue` + "`" + `, ` + "`" + `MaxValue` + "`" + `, and ` + "`" + `DesiredCapacity` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "min_value",
					Description: `(Optional, Available in 1.90.0+) The minimum number of instances in a scaling group when the scaling method of the scheduled task is to specify the number of instances in a scaling group.`,
				},
				resource.Attribute{
					Name:        "max_value",
					Description: `(Optional, Available in 1.90.0+) The maximum number of instances in a scaling group when the scaling method of the scheduled task is to specify the number of instances in a scaling group.`,
				},
				resource.Attribute{
					Name:        "desired_capacity",
					Description: `(Optional, Available in 1.90.0+) The expected number of instances in a scaling group when the scaling method of the scheduled task is to specify the number of instances in a scaling group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the scheduled task, which is 2-200 characters (English or Chinese) long.`,
				},
				resource.Attribute{
					Name:        "launch_time",
					Description: `(Optional) The time at which the scheduled task is triggered. Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time must be in UTC. You cannot enter a time point later than 90 days from the date of scheduled task creation. If the ` + "`" + `recurrence_type` + "`" + ` parameter is specified, the task is executed repeatedly at the time specified by LaunchTime. Otherwise, the task is only executed once at the date and time specified by LaunchTime.`,
				},
				resource.Attribute{
					Name:        "launch_expiration_time",
					Description: `(Optional) The time period during which a failed scheduled task is retried. Unit: seconds. Valid values: 0 to 21600. Default value: 600`,
				},
				resource.Attribute{
					Name:        "recurrence_type",
					Description: `(Optional) Specifies the recurrence type of the scheduled task.`,
				},
				resource.Attribute{
					Name:        "recurrence_value",
					Description: `(Optional) Specifies how often a scheduled task recurs.`,
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
					Description: `The schedule task ID. ## Import ESS schedule task can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ess_scheduled_task.example abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The schedule task ID. ## Import ESS schedule task can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ess_scheduled_task.example abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_fc_function",
			Category:         "Function Compute Service",
			ShortDescription: `Provides a Alicloud Function Compute Function resource. Function allows you to trigger execution of code in response to events in Alibaba Cloud. The Function itself includes source code and runtime configuration.`,
			Description:      ``,
			Keywords: []string{
				"function",
				"compute",
				"service",
				"fc",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service",
					Description: `(Required, ForceNew) The Function Compute service name.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, ForceNew) The Function Compute function name. It is the only in one service and is conflict with "name_prefix".`,
				},
				resource.Attribute{
					Name:        "name_prefix",
					Description: `(Optional, ForceNew) Setting a prefix to get a only function name. It is conflict with "name".`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The Function Compute function description.`,
				},
				resource.Attribute{
					Name:        "filename",
					Description: `(Optional) The path to the function's deployment package within the local filesystem. It is conflict with the ` + "`" + `oss_` + "`" + `-prefixed options.`,
				},
				resource.Attribute{
					Name:        "oss_bucket",
					Description: `(Optional) The OSS bucket location containing the function's deployment package. Conflicts with ` + "`" + `filename` + "`" + `. This bucket must reside in the same Alibaba Cloud region where you are creating the function.`,
				},
				resource.Attribute{
					Name:        "oss_key",
					Description: `(Optional) The OSS key of an object containing the function's deployment package. Conflicts with ` + "`" + `filename` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "handler",
					Description: `(Required) The function [entry point](https://www.alibabacloud.com/help/doc-detail/62213.htm) in your code.`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `(Optional) Amount of memory in MB your Function can use at runtime. Defaults to ` + "`" + `128` + "`" + `. Limits to [128, 3072].`,
				},
				resource.Attribute{
					Name:        "runtime",
					Description: `(Required) See [Runtimes][https://www.alibabacloud.com/help/doc-detail/52077.htm] for valid values.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) The amount of time your Function has to run in seconds.`,
				},
				resource.Attribute{
					Name:        "environment_variables",
					Description: `(Optional, Available in 1.36.0+) A map that defines environment variables for the function.`,
				},
				resource.Attribute{
					Name:        "code_checksum",
					Description: `(Optional, Available in 1.59.0+) The checksum (crc64) of the function code.The value can be generated by data source [alicloud_file_crc64_checksum](https://www.terraform.io/docs/providers/alicloud/d/file_crc64_checksum.html). ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the function. The value is formate as ` + "`" + `<service>:<name>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `The date this resource was last modified.`,
				},
				resource.Attribute{
					Name:        "function_id",
					Description: `The Function Compute service ID.`,
				},
				resource.Attribute{
					Name:        "code_checksum",
					Description: `The checksum (crc64) of the function code. ## Import Function Compute function can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_fc_function.foo my-fc-service:hello-world ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the function. The value is formate as ` + "`" + `<service>:<name>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `The date this resource was last modified.`,
				},
				resource.Attribute{
					Name:        "function_id",
					Description: `The Function Compute service ID.`,
				},
				resource.Attribute{
					Name:        "code_checksum",
					Description: `The checksum (crc64) of the function code. ## Import Function Compute function can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_fc_function.foo my-fc-service:hello-world ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_fc_service",
			Category:         "Function Compute Service",
			ShortDescription: `Provides a Alicloud Function Compute Service resource. The resource is the base of launching Function and Trigger configuration.`,
			Description:      ``,
			Keywords: []string{
				"function",
				"compute",
				"service",
				"fc",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(ForceNew) The Function Compute service name. It is the only in one Alicloud account and is conflict with "name_prefix".`,
				},
				resource.Attribute{
					Name:        "name_prefix",
					Description: `(ForceNew) Setting a prefix to get a only name. It is conflict with "name".`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The function compute service description.`,
				},
				resource.Attribute{
					Name:        "internet_access",
					Description: `(Optional) Whether to allow the service to access Internet. Default to "true".`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Optional) RAM role arn attached to the Function Compute service. This governs both who / what can invoke your Function, as well as what resources our Function has access to. See [User Permissions](https://www.alibabacloud.com/help/doc-detail/52885.htm) for more details.`,
				},
				resource.Attribute{
					Name:        "log_config",
					Description: `(Optional) Provide this to store your FC service logs. Fields documented below. See [Create a Service](https://www.alibabacloud.com/help/doc-detail/51924.htm).`,
				},
				resource.Attribute{
					Name:        "vpc_config",
					Description: `(Optional) Provide this to allow your FC service to access your VPC. Fields documented below. See [Function Compute Service in VPC](https://www.alibabacloud.com/help/faq-detail/72959.htm).`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Required) The project name of Logs service.`,
				},
				resource.Attribute{
					Name:        "logstore",
					Description: `(Required) The log store name of Logs service. ->`,
				},
				resource.Attribute{
					Name:        "vswitch_ids",
					Description: `(Required) A list of vswitch IDs associated with the FC service.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Required) A security group ID associated with the FC service. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the FC service. The value is same as name.`,
				},
				resource.Attribute{
					Name:        "service_id",
					Description: `The Function Compute service ID.`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `The date this resource was last modified. ## Import Function Compute Service can be imported using the id or name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_fc_service.foo my-fc-service ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the FC service. The value is same as name.`,
				},
				resource.Attribute{
					Name:        "service_id",
					Description: `The Function Compute service ID.`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `The date this resource was last modified. ## Import Function Compute Service can be imported using the id or name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_fc_service.foo my-fc-service ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_fc_trigger",
			Category:         "Function Compute Service",
			ShortDescription: `Provides a Alicloud Function Compute Trigger resource.`,
			Description:      ``,
			Keywords: []string{
				"function",
				"compute",
				"service",
				"fc",
				"trigger",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service",
					Description: `(Required, ForceNew) The Function Compute service name.`,
				},
				resource.Attribute{
					Name:        "function",
					Description: `(Required, ForceNew) The Function Compute function name.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(ForceNew) The Function Compute trigger name. It is the only in one service and is conflict with "name_prefix".`,
				},
				resource.Attribute{
					Name:        "name_prefix",
					Description: `(ForceNew) Setting a prefix to get a only trigger name. It is conflict with "name".`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Optional) RAM role arn attached to the Function Compute trigger. Role used by the event source to call the function. The value format is "acs:ram::$account-id:role/$role-name". See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.`,
				},
				resource.Attribute{
					Name:        "source_arn",
					Description: `(Optional, ForceNew) Event source resource address. See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `(Optional) The config of Function Compute trigger.It is valid when ` + "`" + `type` + "`" + ` is not "mns_topic".See [Configure triggers and events](https://www.alibabacloud.com/help/doc-detail/70140.htm) for more details.`,
				},
				resource.Attribute{
					Name:        "config_mns",
					Description: `(Optional, ForceNew, Available in 1.41.0) The config of Function Compute trigger when the type is "mns_topic".It is conflict with ` + "`" + `config` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required, ForceNew) The Type of the trigger. Valid values: ["oss", "log", "timer", "http", "mns_topic", "cdn_events"]. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the function. The value is formate as ` + "`" + `<service>:<function>:<name>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `The date this resource was last modified.`,
				},
				resource.Attribute{
					Name:        "trigger_id",
					Description: `The Function Compute trigger ID. ## Import Function Compute trigger can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_fc_trigger.foo my-fc-service:hello-world:hello-trigger ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the function. The value is formate as ` + "`" + `<service>:<function>:<name>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `The date this resource was last modified.`,
				},
				resource.Attribute{
					Name:        "trigger_id",
					Description: `The Function Compute trigger ID. ## Import Function Compute trigger can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_fc_trigger.foo my-fc-service:hello-world:hello-trigger ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_forward_entry",
			Category:         "VPC",
			ShortDescription: `Provides a Alicloud forward resource.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"forward",
				"entry",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "forward_table_id",
					Description: `(Required, ForceNew) The value can get from ` + "`" + `alicloud_nat_gateway` + "`" + ` Attributes "forward_table_ids".`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, Available in 1.44.0+) The name of forward entry.`,
				},
				resource.Attribute{
					Name:        "external_ip",
					Description: `(Required, ForceNew) The external ip address, the ip must along bandwidth package public ip which ` + "`" + `alicloud_nat_gateway` + "`" + ` argument ` + "`" + `bandwidth_packages` + "`" + `.`,
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
			Type:             "alicloud_gpdb_connection",
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
					Description: `(Optional) Internet connection port. Valid value: [3200-3999]. Default to 3306. ### Timeouts ->`,
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
					Description: `The ip address of connection string. ## Import AnalyticDB for PostgreSQL's connection can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_gpdb_connection.example abc12345678 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `The ip address of connection string. ## Import AnalyticDB for PostgreSQL's connection can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_gpdb_connection.example abc12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_gpdb_instance",
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
					Description: `(Optional, ForceNew) The Zone to launch the DB instance. it supports multiple zone. If it is a multi-zone and ` + "`" + `vswitch_id` + "`" + ` is specified, the vswitch must in one of them. The multiple zone ID can be retrieved by setting ` + "`" + `multi` + "`" + ` to "true" in the data source ` + "`" + `alicloud_zones` + "`" + `.`,
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
					Description: `(Optional, Available in v1.55.3+) A mapping of tags to assign to the resource. ### Timeouts ->`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 30 mins) Used when creating the DB instance (until it reaches the initial ` + "`" + `Running` + "`" + ` status). ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Instance. ## Import AnalyticDB for PostgreSQL can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_gpdb_instance.example gp-bp1291daeda44194 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Instance. ## Import AnalyticDB for PostgreSQL can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_gpdb_instance.example gp-bp1291daeda44194 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_hbase_instance",
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
					Description: `(Optional, ForceNew) The Zone to launch the HBase instance. if vswitch_id is not empty, this zone_id can be "" or consistent.`,
				},
				resource.Attribute{
					Name:        "hbase",
					Description: `(Optional, ForceNew) "hbase/hbaseue/bds", The following types are supported after v1.73.0: ` + "`" + `hbaseue` + "`" + ` and ` + "`" + `bds ` + "`" + ``,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `(Required, ForceNew) hbase major version. hbase:1.1/2.0, hbaseue:2.0, bds:1.0, unsupport other engine temporarily. Value options can refer to the latest docs [CreateInstance](https://help.aliyun.com/document_detail/144607.html).`,
				},
				resource.Attribute{
					Name:        "core_disk_size",
					Description: `(Optional, ForceNew) User-defined HBase instance one core node's storage space.Unit: GB. Value range: - Custom storage space; value range: [400, 8000] - 40-GB increments.`,
				},
				resource.Attribute{
					Name:        "pay_type",
					Description: `(Optional, ForceNew) Valid values are ` + "`" + `PrePaid` + "`" + `, ` + "`" + `PostPaid` + "`" + `,System default to ` + "`" + `PostPaid` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `(Optional, ForceNew) 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 60, valid when pay_type = PrePaid. unit: month.`,
				},
				resource.Attribute{
					Name:        "auto_renew",
					Description: `(Optional, ForceNew) ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + `, System default to ` + "`" + `false` + "`" + `, valid when pay_type = PrePaid.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Optional, ForceNew) if vswitch_id is not empty, that mean net_type = vpc and has a same region. if vswitch_id is empty, net_type_classic`,
				},
				resource.Attribute{
					Name:        "cold_storage_size",
					Description: `(Optional, ForceNew) 0 or 0+. 0 means is_cold_storage = false. 0+ means is_cold_storage = true`,
				},
				resource.Attribute{
					Name:        "maintain_start_time",
					Description: `(Optional, Available in 1.73.0) The start time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).`,
				},
				resource.Attribute{
					Name:        "maintain_end_time",
					Description: `(Optional, Available in 1.73.0) The end time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).`,
				},
				resource.Attribute{
					Name:        "deletion_protection",
					Description: `(Optional, Available in 1.73.0) the switch of delete protection. true: delete protect, false: no delete protect. you must set false when you want to delete cluster.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Available in 1.73.0) A mapping of tags to assign to the resource. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the HBase. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 30 mins) Used when creating the HBase instance (until it reaches the initial ` + "`" + `ACTIVATION` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 30 mins) Used when terminating the HBase instance. ## Import HBase can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_hbase_instance.example hb-wz96815u13k659fvd ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the HBase. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 30 mins) Used when creating the HBase instance (until it reaches the initial ` + "`" + `ACTIVATION` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 30 mins) Used when terminating the HBase instance. ## Import HBase can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_hbase_instance.example hb-wz96815u13k659fvd ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_image",
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
					Name:        "architecture",
					Description: `(Optional, ForceNew) Specifies the architecture of the system disk after you specify a data disk snapshot as the data source of the system disk for creating an image. Valid values: ` + "`" + `i386` + "`" + ` , Default is ` + "`" + `x86_64` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "platform",
					Description: `(Optional, ForceNew) Specifies the operating system platform of the system disk after you specify a data disk snapshot as the data source of the system disk for creating an image. Valid values: ` + "`" + `CentOS` + "`" + `, ` + "`" + `Ubuntu` + "`" + `, ` + "`" + `SUSE` + "`" + `, ` + "`" + `OpenSUSE` + "`" + `, ` + "`" + `RedHat` + "`" + `, ` + "`" + `Debian` + "`" + `, ` + "`" + `CoreOS` + "`" + `, ` + "`" + `Aliyun Linux` + "`" + `, ` + "`" + `Windows Server 2003` + "`" + `, ` + "`" + `Windows Server 2008` + "`" + `, ` + "`" + `Windows Server 2012` + "`" + `, ` + "`" + `Windows 7` + "`" + `, Default is ` + "`" + `Others Linux` + "`" + `, ` + "`" + `Customized Linux` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The tag value of an image. The value of N ranges from 1 to 20.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional, ForceNew) The ID of the enterprise resource group to which a custom image belongs`,
				},
				resource.Attribute{
					Name:        "disk_device_mapping",
					Description: `(Optional, ForceNew, Conflict with ` + "`" + `snapshot_id ` + "`" + ` and ` + "`" + `instance_id ` + "`" + `) Description of the system with disks and snapshots under the image.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `(Optional, ForceNew) Specifies the type of a disk in the combined custom image. If you specify this parameter, you can use a data disk snapshot as the data source of a system disk for creating an image. If it is not specified, the disk type is determined by the corresponding snapshot. Valid values: ` + "`" + `system` + "`" + `, ` + "`" + `data` + "`" + `,`,
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
					Name:        "device",
					Description: `(Optional, ForceNew)Specifies the name of a disk in the combined custom image. Value range: /dev/xvda to /dev/xvdz.`,
				},
				resource.Attribute{
					Name:        "force",
					Description: `(Optional) Indicates whether to force delete the custom image, Default is ` + "`" + `false` + "`" + `. - true：Force deletes the custom image, regardless of whether the image is currently being used by other instances. - false：Verifies that the image is not currently in use by any other instances before deleting the image. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 10 mins) Used when creating the image (until it reaches the initial ` + "`" + `Available` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 10 mins) Used when terminating the image. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the image. ## Import image can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_image.default m-uf66871ape`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the image. ## Import image can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_image.default m-uf66871ape`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_image_copy",
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
					Description: `(Required, ForceNew) The ID of the region to which the source custom image belongs. You can call [DescribeRegions](https://www.alibabacloud.com/help/doc-detail/25609.htm) to view the latest regions of Alibaba Cloud.`,
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
					Description: `ID of the image. ## Import image can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_image_copy.default m-uf66871ape`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the image. ## Import image can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_image_copy.default m-uf66871ape`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_image_export",
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
			Type:             "alicloud_image_import",
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
					Description: `ID of the image. ## Import image can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_image_import.default m-uf66871ape`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the image. ## Import image can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_image_import.default m-uf66871ape`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_image_share_permission",
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
					Description: `(Required, ForceNew) Alibaba Cloud Account ID. It is used to share images. ## Attributes Reference0 The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the image. It formats as ` + "`" + `<image_id>:<account_id>` + "`" + ` ## Import image can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_image_share_permission.default m-uf66yg1q:123456789 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the image. It formats as ` + "`" + `<image_id>:<account_id>` + "`" + ` ## Import image can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_image_share_permission.default m-uf66yg1q:123456789 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_instance",
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
					Description: `(Required) The Image to use for the instance. ECS instance's image can be replaced via changing ` + "`" + `image_id` + "`" + `. When it is changed, the instance will reboot to make the change take effect.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Required) The type of instance to start. When it is changed, the instance will reboot to make the change take effect.`,
				},
				resource.Attribute{
					Name:        "io_optimized",
					Description: `(Deprecated) It has been deprecated on instance resource. All the launched alicloud instances will be I/O optimized.`,
				},
				resource.Attribute{
					Name:        "is_outdated",
					Description: `(Optional) Whether to use outdated instance type. Default to false.`,
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
					Description: `(Optional) The name of the ECS. This instance_name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin with a hyphen, and must not begin with http:// or https://. If not specified, Terraform will autogenerate a default name is ` + "`" + `ECS-Instance` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "allocate_public_ip",
					Description: `(Deprecated) It has been deprecated from version "1.7.0". Setting "internet_max_bandwidth_out" larger than 0 can allocate a public ip address for an instance.`,
				},
				resource.Attribute{
					Name:        "system_disk_category",
					Description: `(Optional,ForceNew) Valid values are ` + "`" + `ephemeral_ssd` + "`" + `, ` + "`" + `cloud_efficiency` + "`" + `, ` + "`" + `cloud_ssd` + "`" + `, ` + "`" + `cloud_essd` + "`" + `, ` + "`" + `cloud` + "`" + `. ` + "`" + `cloud` + "`" + ` only is used to some none I/O optimized instance. Default to ` + "`" + `cloud_efficiency` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "system_disk_size",
					Description: `(Optional) Size of the system disk, measured in GiB. Value range: [20, 500]. The specified value must be equal to or greater than max{20, Imagesize}. Default value: max{40, ImageSize}.`,
				},
				resource.Attribute{
					Name:        "system_disk_auto_snapshot_policy_id",
					Description: `(Optional, ForceNew, Available in 1.73.0+) The ID of the automatic snapshot policy applied to the system disk.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the instance, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.`,
				},
				resource.Attribute{
					Name:        "internet_charge_type",
					Description: `(Optional) Internet charge type of the instance, Valid values are ` + "`" + `PayByBandwidth` + "`" + `, ` + "`" + `PayByTraffic` + "`" + `. Default is ` + "`" + `PayByTraffic` + "`" + `. At present, 'PrePaid' instance cannot change the value to "PayByBandwidth" from "PayByTraffic".`,
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
					Description: `(Optional) Host name of the ECS, which is a string of at least two characters. “hostname” cannot start or end with “.” or “-“. In addition, two or more consecutive “.” or “-“ symbols are not allowed. On Windows, the host name can contain a maximum of 15 characters, which can be a combination of uppercase/lowercase letters, numerals, and “-“. The host name cannot contain dots (“.”) or contain only numeric characters. When it is changed, the instance will reboot to make the change take effect. On other OSs such as Linux, the host name can contain a maximum of 64 characters, which can be segments separated by dots (“.”), where each segment can contain uppercase/lowercase letters, numerals, or “_“. When it is changed, the instance will reboot to make the change take effect.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional, Sensitive) Password to an instance is a string of 8 to 30 characters. It must contain uppercase/lowercase letters and numerals, but cannot contain special symbols. When it is changed, the instance will reboot to make the change take effect.`,
				},
				resource.Attribute{
					Name:        "kms_encrypted_password",
					Description: `(Optional, Available in 1.57.1+) An KMS encrypts password used to an instance. If the ` + "`" + `password` + "`" + ` is filled in, this field will be ignored. When it is changed, the instance will reboot to make the change take effect.`,
				},
				resource.Attribute{
					Name:        "kms_encryption_context",
					Description: `(Optional, MapString, Available in 1.57.1+) An KMS encryption context used to decrypt ` + "`" + `kms_encrypted_password` + "`" + ` before creating or updating an instance with ` + "`" + `kms_encrypted_password` + "`" + `. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when ` + "`" + `kms_encrypted_password` + "`" + ` is set. When it is changed, the instance will reboot to make the change take effect.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Optional) The virtual switch ID to launch in VPC. This parameter must be set unless you can create classic network instances. When it is changed, the instance will reboot to make the change take effect.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `(Optional) Valid values are ` + "`" + `PrePaid` + "`" + `, ` + "`" + `PostPaid` + "`" + `, The default is ` + "`" + `PostPaid` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(ForceNew, ForceNew, Available in 1.57.0+) The Id of resource group which the instance belongs.`,
				},
				resource.Attribute{
					Name:        "period_unit",
					Description: `(Optional) The duration unit that you will buy the resource. It is valid when ` + "`" + `instance_charge_type` + "`" + ` is 'PrePaid'. Valid value: ["Week", "Month"]. Default to "Month".`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) The duration that you will buy the resource, in month. It is valid when ` + "`" + `instance_charge_type` + "`" + ` is ` + "`" + `PrePaid` + "`" + `. Default to 1. Valid values: - [1-9, 12, 24, 36, 48, 60] when ` + "`" + `period_unit` + "`" + ` in "Month" - [1-3] when ` + "`" + `period_unit` + "`" + ` in "Week"`,
				},
				resource.Attribute{
					Name:        "renewal_status",
					Description: `(Optional) Whether to renew an ECS instance automatically or not. It is valid when ` + "`" + `instance_charge_type` + "`" + ` is ` + "`" + `PrePaid` + "`" + `. Default to "Normal". Valid values: - ` + "`" + `AutoRenewal` + "`" + `: Enable auto renewal. - ` + "`" + `Normal` + "`" + `: Disable auto renewal. - ` + "`" + `NotRenewal` + "`" + `: No renewal any longer. After you specify this value, Alibaba Cloud stop sending notification of instance expiry, and only gives a brief reminder on the third day before the instance expiry.`,
				},
				resource.Attribute{
					Name:        "auto_renew_period",
					Description: `(Optional) Auto renewal period of an instance, in the unit of month. It is valid when ` + "`" + `instance_charge_type` + "`" + ` is ` + "`" + `PrePaid` + "`" + `. Default to 1. Valid value: - [1, 2, 3, 6, 12] when ` + "`" + `period_unit` + "`" + ` in "Month" - [1, 2, 3] when ` + "`" + `period_unit` + "`" + ` in "Week"`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string. - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.`,
				},
				resource.Attribute{
					Name:        "volume_tags",
					Description: `(Optional) A mapping of tags to assign to the devices created by the instance at launch time. - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string. - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional) User-defined data to customize the startup behaviors of an ECS instance and to pass data into an ECS instance. From version 1.60.0, it can be update in-place. If updated, the instance will reboot to make the change take effect. Note: Not all of changes will take effect and it depends on [cloud-init module type](https://cloudinit.readthedocs.io/en/latest/topics/modules.html).`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `(Optional, Force new resource) The name of key pair that can login ECS instance successfully without password. If it is specified, the password would be invalid.`,
				},
				resource.Attribute{
					Name:        "role_name",
					Description: `(Optional, Force new resource) Instance RAM role name. The name is provided and maintained by RAM. You can use ` + "`" + `alicloud_ram_role` + "`" + ` to create a new one.`,
				},
				resource.Attribute{
					Name:        "include_data_disks",
					Description: `(Optional) Whether to change instance disks charge type when changing instance charge type.`,
				},
				resource.Attribute{
					Name:        "dry_run",
					Description: `(Optional) Specifies whether to send a dry-run request. Default to false. - true: Only a dry-run request is sent and no instance is created. The system checks whether the required parameters are set, and validates the request format, service permissions, and available ECS instances. If the validation fails, the corresponding error code is returned. If the validation succeeds, the ` + "`" + `DryRunOperation` + "`" + ` error code is returned. - false: A request is sent. If the validation succeeds, the instance is created.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `(Optional) Instance private IP address can be specified when you creating new instance. It is valid when ` + "`" + `vswitch_id` + "`" + ` is specified. When it is changed, the instance will reboot to make the change take effect.`,
				},
				resource.Attribute{
					Name:        "credit_specification",
					Description: `(Optional, Available in 1.57.1+) Performance mode of the t5 burstable instance. Valid values: 'Standard', 'Unlimited'.`,
				},
				resource.Attribute{
					Name:        "spot_strategy",
					Description: `(Optional, ForceNew) The spot strategy of a Pay-As-You-Go instance, and it takes effect only when parameter ` + "`" + `instance_charge_type` + "`" + ` is 'PostPaid'. Value range: - NoSpot: A regular Pay-As-You-Go instance. - SpotWithPriceLimit: A price threshold for a spot instance - SpotAsPriceGo: A price that is based on the highest Pay-As-You-Go instance Default to NoSpot. Note: Currently, the spot instance only supports domestic site account.`,
				},
				resource.Attribute{
					Name:        "spot_price_limit",
					Description: `(Optional, Float, ForceNew) The hourly price threshold of a instance, and it takes effect only when parameter 'spot_strategy' is 'SpotWithPriceLimit'. Three decimals is allowed at most.`,
				},
				resource.Attribute{
					Name:        "deletion_protection",
					Description: `(Optional, true) Whether enable the deletion protection or not. Default value: ` + "`" + `false` + "`" + `. - true: Enable deletion protection. - false: Disable deletion protection.`,
				},
				resource.Attribute{
					Name:        "force_delete",
					Description: `(Optional, Available in 1.18.0+) If it is true, the "PrePaid" instance will be change to "PostPaid" and then deleted forcibly. However, because of changing instance charge type has CPU core count quota limitation, so strongly recommand that "Don't modify instance charge type frequentlly in one month".`,
				},
				resource.Attribute{
					Name:        "auto_release_time",
					Description: `(Optional, Available in 1.70.0+) The automatic release time of the ` + "`" + `PostPaid` + "`" + ` instance. The time follows the ISO 8601 standard and is in UTC time. Format: yyyy-MM-ddTHH:mm:ssZ. It must be at least half an hour later than the current time and less than 3 years since the current time. Set it to null can cancel automatic release attribute and the ECS instance will not be released automatically.`,
				},
				resource.Attribute{
					Name:        "security_enhancement_strategy",
					Description: `(Optional, ForceNew) The security enhancement strategy. - Active: Enable security enhancement strategy, it only works on system images. - Deactive: Disable security enhancement strategy, it works on all images.`,
				},
				resource.Attribute{
					Name:        "data_disks",
					Description: `(Optional, ForceNew, Available 1.23.1+) The list of data disks created with instance.`,
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
					Description: `(Optional, ForceNew) The category of the disk: - ` + "`" + `cloud` + "`" + `: The general cloud disk. - ` + "`" + `cloud_efficiency` + "`" + `: The efficiency cloud disk. - ` + "`" + `cloud_ssd` + "`" + `: The SSD cloud disk. - ` + "`" + `cloud_essd` + "`" + `: The ESSD cloud disk. - ` + "`" + `ephemeral_ssd` + "`" + `: The local SSD disk. Default to ` + "`" + `cloud_efficiency` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "kms_key_id",
					Description: `(Optional, Available in 1.90.1+) The KMS key ID corresponding to the Nth data disk.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional, ForceNew) The snapshot ID used to initialize the data disk. If the size specified by snapshot is greater that the size of the disk, use the size specified by snapshot as the size of the data disk.`,
				},
				resource.Attribute{
					Name:        "auto_snapshot_policy_id",
					Description: `(Optional, ForceNew, Available in 1.73.0+) The ID of the automatic snapshot policy applied to the system disk.`,
				},
				resource.Attribute{
					Name:        "delete_with_instance",
					Description: `(Optional, ForceNew) Delete this data disk when the instance is destroyed. It only works on cloud, cloud_efficiency, cloud_essd, cloud_ssd disk. If the category of this data disk was ephemeral_ssd, please don't set this param. Default value: ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, ForceNew) The description of the data disk.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional 1.85.0+) The instance status. Valid values: ["Running", "Stopped"]. You can control the instance start and stop through this parameter. Default to ` + "`" + `Running` + "`" + `. ->`,
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
					Name:        "public_ip",
					Description: `The instance public ip. ## Import Instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_instance.example i-abc12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The instance ID.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The instance public ip. ## Import Instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_instance.example i-abc12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_key_pair",
			Category:         "ECS",
			ShortDescription: `Provides a Alicloud key pair resource.`,
			Description:      ``,
			Keywords: []string{
				"ecs",
				"key",
				"pair",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key_name",
					Description: `(ForceNew) The key pair's name. It is the only in one Alicloud account.`,
				},
				resource.Attribute{
					Name:        "key_name_prefix",
					Description: `(ForceNew) The key pair name's prefix. It is conflict with ` + "`" + `key_name` + "`" + `. If it is specified, terraform will using it to build the only key name.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(ForceNew) You can import an existing public key and using Alicloud key pair to manage it.`,
				},
				resource.Attribute{
					Name:        "key_file",
					Description: `(ForceNew) The name of file to save your new key pair's private key. Strongly suggest you to specified it when you creating key pair, otherwise, you wouldn't get its private key ever.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(ForceNew, Available in 1.57.0+) The Id of resource group which the key pair belongs.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Available in v1.66.0+) A mapping of tags to assign to the resource. ->`,
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
			Type:             "alicloud_key_pair_attachment",
			Category:         "ECS",
			ShortDescription: `Provides a Alicloud key pair attachment resource to bind key pair for several ECS instances.`,
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
			Type:             "alicloud_kms_alias",
			Category:         "KMS",
			ShortDescription: `Provides a Alicloud KMS Alias resource.`,
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
					Description: `The ID of the alias. ## Import KMS alias can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_kms_alias.example alias/test_kms_alias ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the alias. ## Import KMS alias can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_kms_alias.example alias/test_kms_alias ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_kms_ciphertext",
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
					Description: `(Optional, ForceNew) The Encryption context. If you specify this parameter here, it is also required when you call the Decrypt API operation. For more information, see [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
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
			Type:             "alicloud_kms_key",
			Category:         "KMS",
			ShortDescription: `Provides a Alikms key resource.`,
			Description:      ``,
			Keywords: []string{
				"kms",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the key as viewed in Alicloud console.`,
				},
				resource.Attribute{
					Name:        "key_usage",
					Description: `(Optional, ForceNew) Specifies the usage of CMK. Currently, default to 'ENCRYPT/DECRYPT', indicating that CMK is used for encryption and decryption.`,
				},
				resource.Attribute{
					Name:        "deletion_window_in_days",
					Description: `(Optional) Field 'deletion_window_in_days' has been deprecated from provider version 1.85.0. New field 'pending_window_in_days' instead.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `(Optional) Field 'is_enabled' has been deprecated from provider version 1.85.0. New field 'key_state' instead.`,
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
					Description: `The Alicloud Resource Name (ARN) of the key.`,
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
					Description: `The ID of the current primary key version of the symmetric CMK. ## Import Alikms key can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_kms_key.example abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the key.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The Alicloud Resource Name (ARN) of the key.`,
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
					Description: `The ID of the current primary key version of the symmetric CMK. ## Import Alikms key can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_kms_key.example abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_kms_key_version",
			Category:         "KMS",
			ShortDescription: `Provides a Alikms key version resource.`,
			Description:      ``,
			Keywords: []string{
				"kms",
				"key",
				"version",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key_id",
					Description: `(Required, ForceNew) The id of the master key (CMK). ->`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `The date and time (UTC time) when the Alikms key version was created.`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `The id of the master key (CMK).`,
				},
				resource.Attribute{
					Name:        "key_version_id",
					Description: `The id of the Alikms key version. ## Import Alikms key version can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_kms_key_version.example 72da539a-2fa8-4f2d-b854-`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "creation_date",
					Description: `The date and time (UTC time) when the Alikms key version was created.`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `The id of the master key (CMK).`,
				},
				resource.Attribute{
					Name:        "key_version_id",
					Description: `The id of the Alikms key version. ## Import Alikms key version can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_kms_key_version.example 72da539a-2fa8-4f2d-b854-`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_kms_secret",
			Category:         "KMS",
			ShortDescription: `Provides a Alibaba Cloud kms secret resource.`,
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
					Description: `The Alicloud Resource Name (ARN) of the secret.`,
				},
				resource.Attribute{
					Name:        "planned_delete_time",
					Description: `The time when the secret is scheduled to be deleted. ## Import KMS secret can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_kms_secret.default secret-foo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the secret. It same with ` + "`" + `secret_name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The Alicloud Resource Name (ARN) of the secret.`,
				},
				resource.Attribute{
					Name:        "planned_delete_time",
					Description: `The time when the secret is scheduled to be deleted. ## Import KMS secret can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_kms_secret.default secret-foo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_kvstore_account",
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
					Description: `The current account resource ID. Composed of instance ID and account name with format ` + "`" + `<instance_id>:<name>` + "`" + `. ## Import kvstore account can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_KVStore_account.example "rm-12345:tf_account" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The current account resource ID. Composed of instance ID and account name with format ` + "`" + `<instance_id>:<name>` + "`" + `. ## Import kvstore account can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_KVStore_account.example "rm-12345:tf_account" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_kvstore_backup_policy",
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
					Description: `Backup Cycle. Allowed values: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday ## Import KVStore backup policy can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_kvstore_backup_policy.example r-abc12345678 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `Backup Cycle. Allowed values: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday ## Import KVStore backup policy can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_kvstore_backup_policy.example r-abc12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_kvstore_instance",
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
					Description: `(Optional, Available in 1.57.1+) An KMS encrypts password used to a instance. If the ` + "`" + `password` + "`" + ` is filled in, this field will be ignored.`,
				},
				resource.Attribute{
					Name:        "kms_encryption_context",
					Description: `(Optional, MapString, Available in 1.57.1+) An KMS encryption context used to decrypt ` + "`" + `kms_encrypted_password` + "`" + ` before creating or updating instance with ` + "`" + `kms_encrypted_password` + "`" + `. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when ` + "`" + `kms_encrypted_password` + "`" + ` is set.`,
				},
				resource.Attribute{
					Name:        "instance_class",
					Description: `(Required) Type of the applied ApsaraDB for Redis instance. It can be retrieved by data source [` + "`" + `alicloud_kvstore_instance_classes` + "`" + `](https://www.terraform.io/docs/providers/alicloud/d/kvstore_instance_classes.html) or referring to help-docs [Instance type table](https://www.alibabacloud.com/help/doc-detail/26350.htm).`,
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
					Name:        "auto_renew",
					Description: `(Optional, Available in 1.36.0+) Whether to renewal a DB instance automatically or not. It is valid when instance_charge_type is ` + "`" + `PrePaid` + "`" + `. Default to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "auto_renew_period",
					Description: `(Optional, Available in 1.36.0+) Auto-renewal period of an instance, in the unit of the month. It is valid when instance_charge_type is ` + "`" + `PrePaid` + "`" + `. Valid value:[1~12], Default to 1.`,
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
					Description: `(Optional, Available in 1.76.0+) The Security Group ID of ECS.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `(Optional) Set of parameters needs to be set after instance was launched. Available parameters can refer to the latest docs [Instance configurations table](https://www.alibabacloud.com/help/doc-detail/61209.htm) .`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Available in v1.55.3+) A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "maintain_start_time",
					Description: `(Optional, Available in v1.56.0+) The start time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).`,
				},
				resource.Attribute{
					Name:        "maintain_end_time",
					Description: `(Optional, Available in v1.56.0+) The end time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional, ForceNew, Available in v1.86.0+) The ID of resource group which the resource belongs. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The KVStore instance ID.`,
				},
				resource.Attribute{
					Name:        "connection_domain",
					Description: `Instance connection domain (only Intranet access supported). ### Timeouts ->`,
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
					Description: `(Defaults to 20 mins) Used when terminating the KVStore instance. ## Import KVStore instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_kvstore_instance.example r-abc12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The KVStore instance ID.`,
				},
				resource.Attribute{
					Name:        "connection_domain",
					Description: `Instance connection domain (only Intranet access supported). ### Timeouts ->`,
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
					Description: `(Defaults to 20 mins) Used when terminating the KVStore instance. ## Import KVStore instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_kvstore_instance.example r-abc12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_launch_template",
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
					Description: `(Required, ForceNew) Instance launch template name. Can contain [2, 128] characters in length. It must start with an English letter or Chinese, can contain numbers, periods (.), colons (:), underscores (_), and hyphens (-). It cannot start with "http://" or "https://".`,
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
					Name:        "instance_charge_type",
					Description: `(Optional)Billing methods. Optional values: - PrePaid: Monthly, or annual subscription. Make sure that your registered credit card is invalid or you have insufficient balance in your PayPal account. Otherwise, InvalidPayMethod error may occur. - PostPaid: Pay-As-You-Go. Default value: PostPaid.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Optional) Instance type. For more information, call resource_alicloud_instances to obtain the latest instance type list.`,
				},
				resource.Attribute{
					Name:        "auto_release_time",
					Description: `(Optional) Instance auto release time. The time is presented using the ISO8601 standard and in UTC time. The format is YYYY-MM-DDTHH:MM:SSZ.`,
				},
				resource.Attribute{
					Name:        "internet_charge_type",
					Description: `(Optional) Internet bandwidth billing method. Optional values: ` + "`" + `PayByTraffic` + "`" + ` | ` + "`" + `PayByBandwidth` + "`" + `.`,
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
					Description: `(Optional) Network type of the instance. Value options: ` + "`" + `classic` + "`" + ` | ` + "`" + `vpc` + "`" + `.`,
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
					Description: `The Launch Template ID. ## Import Launch Template can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_launch_template.lt lt-abc1234567890000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Launch Template ID. ## Import Launch Template can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_launch_template.lt lt-abc1234567890000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_log_alert",
			Category:         "Log Service (SLS)",
			ShortDescription: `Provides a Alicloud log alert resource.`,
			Description:      ``,
			Keywords: []string{
				"log",
				"service",
				"sls",
				"alert",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_name",
					Description: `(Required, ForceNew) The project name.`,
				},
				resource.Attribute{
					Name:        "alert_name",
					Description: `(Required, ForceNew) Name of logstore for configuring alarm service.`,
				},
				resource.Attribute{
					Name:        "alert_displayname",
					Description: `(Required) Alert displayname.`,
				},
				resource.Attribute{
					Name:        "alert_description",
					Description: `(Optional) Alert description.`,
				},
				resource.Attribute{
					Name:        "condition",
					Description: `(Required) Conditional expression, such as: count> 100.`,
				},
				resource.Attribute{
					Name:        "dashboard",
					Description: `(Required) The name of the dashboard associated with the alarm. The name of the instrument cluster associated with the alarm. If there is no such instrument cluster, terraform will help you create an empty instrument cluster.`,
				},
				resource.Attribute{
					Name:        "mute_until",
					Description: `(Optional) Timestamp, notifications before closing again.`,
				},
				resource.Attribute{
					Name:        "throttling",
					Description: `(Optional) Notification interval, default is no interval. Support number + unit type, for example 60s, 1h.`,
				},
				resource.Attribute{
					Name:        "notify_threshold",
					Description: `(Optional) Notification threshold, which is not notified until the number of triggers is reached. The default is 1.`,
				},
				resource.Attribute{
					Name:        "query_list",
					Description: `(Required) Multiple conditions for configured alarm query.`,
				},
				resource.Attribute{
					Name:        "chart_title",
					Description: `(Required) chart title`,
				},
				resource.Attribute{
					Name:        "logstore",
					Description: `(Required) Query logstore`,
				},
				resource.Attribute{
					Name:        "query",
					Description: `(Required) query corresponding to chart. example:`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `(Required) begin time. example: -60s.`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `(Required) end time. example: 20s.`,
				},
				resource.Attribute{
					Name:        "time_span_type",
					Description: `(Optional) default Custom. No need to configure this parameter.`,
				},
				resource.Attribute{
					Name:        "notification_list",
					Description: `(Required) Alarm information notification list.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Notification type. support Email, SMS, DingTalk.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Required) Notice content of alarm.`,
				},
				resource.Attribute{
					Name:        "service_uri",
					Description: `(Optional) Request address.`,
				},
				resource.Attribute{
					Name:        "mobile_list",
					Description: `(Optional) SMS sending mobile number.`,
				},
				resource.Attribute{
					Name:        "email_list",
					Description: `(Optional) Email address list.`,
				},
				resource.Attribute{
					Name:        "schedule_interval",
					Description: `(Optional) Execution interval. 60 seconds minimum, such as 60s, 1h.`,
				},
				resource.Attribute{
					Name:        "schedule_type",
					Description: `(Optional) Default FixedRate. No need to configure this parameter. ## Attributes Reference The following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_log_audit",
			Category:         "Log Service (SLS)",
			ShortDescription: `Provides a Alicloud log audit resource.`,
			Description:      ``,
			Keywords: []string{
				"log",
				"service",
				"sls",
				"audit",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required, ForceNew) Name of SLS log audit.`,
				},
				resource.Attribute{
					Name:        "aliuid",
					Description: `(Required, ForceNew) Aliuid value of your account.`,
				},
				resource.Attribute{
					Name:        "variable_map",
					Description: `(Required) Log audit detailed configuration.`,
				},
				resource.Attribute{
					Name:        "actiontrail_enabled",
					Description: `(Optional) Notification type. Support Email, SMS, DingTalk. Default true.`,
				},
				resource.Attribute{
					Name:        "actiontrail_ttl",
					Description: `(Optional) Actiontril action log TTL. Default 180.`,
				},
				resource.Attribute{
					Name:        "oss_access_enabled",
					Description: `(Optional) Access log switch of OSS. Default true.`,
				},
				resource.Attribute{
					Name:        "oss_access_ttl",
					Description: `(Optional) Access log TTL of OSS. Default 180.`,
				},
				resource.Attribute{
					Name:        "oss_sync_enabled",
					Description: `(Optional) OSS synchronization to central configuration switch. Default true.`,
				},
				resource.Attribute{
					Name:        "oss_sync_ttl",
					Description: `(Optional) OSS synchronization to central TTL. Default 180.`,
				},
				resource.Attribute{
					Name:        "oss_metering_enabled",
					Description: `(Optional) OSS metering log switch.Default true.`,
				},
				resource.Attribute{
					Name:        "oss_metering_ttl",
					Description: `(Optional) OSS measurement log TTL. Default 180.`,
				},
				resource.Attribute{
					Name:        "rds_enabled",
					Description: `(Optional) RDS audit log switch. Default true.`,
				},
				resource.Attribute{
					Name:        "rds_ttl",
					Description: `(Optional) Dds log centralization ttl. Default 180.`,
				},
				resource.Attribute{
					Name:        "slb_access_enabled",
					Description: `(Optional) Slb log switch. Default true.`,
				},
				resource.Attribute{
					Name:        "slb_access_ttl",
					Description: `(Optional) Slb centralized ttl. Default 180.`,
				},
				resource.Attribute{
					Name:        "slb_sync_enabled",
					Description: `(Optional) Slb sync to center switch. Default true.`,
				},
				resource.Attribute{
					Name:        "slb_sync_ttl",
					Description: `(Optional) Slb sync to center switch. Default 180.`,
				},
				resource.Attribute{
					Name:        "bastion_enabled",
					Description: `(Optional) Fortress machine operation log switch.Default true.`,
				},
				resource.Attribute{
					Name:        "bastion_ttl",
					Description: `(Optional) Fort machine centralized ttl. Default 180.`,
				},
				resource.Attribute{
					Name:        "waf_enabled",
					Description: `(Optional) Waf log switch. Default true.`,
				},
				resource.Attribute{
					Name:        "waf_ttl",
					Description: `(Optional) Waf centralized ttl. Default true.`,
				},
				resource.Attribute{
					Name:        "cloudfirewall_ttl",
					Description: `(Optional) Cloud firewall switch.Default true.`,
				},
				resource.Attribute{
					Name:        "sas_ttl",
					Description: `(Optional) Cloud Security Center centralized ttl. Default 180.`,
				},
				resource.Attribute{
					Name:        "sas_process_enabled",
					Description: `(Optional) Cloud Security Center process startup log switch. Default false.`,
				},
				resource.Attribute{
					Name:        "sas_network_enabled",
					Description: `(Optional) Cloud security center network connection log switch. Default false.`,
				},
				resource.Attribute{
					Name:        "sas_login_enabled",
					Description: `(Optional) Cloud security center login flow log switch. Default false.`,
				},
				resource.Attribute{
					Name:        "sas_crack_enabled",
					Description: `(Optional) Cloud Security Center Brute Force Log Switch. Default false.`,
				},
				resource.Attribute{
					Name:        "sas_snapshot_process_enabled",
					Description: `(Optional) Cloud Security Center process snapshot switch. Default false.`,
				},
				resource.Attribute{
					Name:        "sas_snapshot_account_enabled",
					Description: `(Optional) Cloud Security Center account snapshot switch. Default false.`,
				},
				resource.Attribute{
					Name:        "sas_snapshot_port_enabled",
					Description: `(Optional) Cloud Security Center Port Snapshot Switch. Default false.`,
				},
				resource.Attribute{
					Name:        "sas_dns_enabled",
					Description: `(Optional) Cloud Security Center DNS resolution log switch. Default false.`,
				},
				resource.Attribute{
					Name:        "sas_local_dns_enabled",
					Description: `(Optional) Cloud security center local DNS log switch. Default false.`,
				},
				resource.Attribute{
					Name:        "sas_session_enabled",
					Description: `(Optional) Cloud security center network session log switch.Default false.`,
				},
				resource.Attribute{
					Name:        "sas_http_enabled",
					Description: `(Optional). Cloud Security Center WEB access log switch. Default false.`,
				},
				resource.Attribute{
					Name:        "sas_security_vul_enabled",
					Description: `(Optional) Cloud Security Center Vulnerability Log Switch.Default false.`,
				},
				resource.Attribute{
					Name:        "sas_security_hc_enabled",
					Description: `(Optional) Cloud Security Center Baseline Log Switch. Default false.`,
				},
				resource.Attribute{
					Name:        "sas_security_alert_enabled",
					Description: `(Optional) Cloud Security Center Security Alarm Log Switch .Default false.`,
				},
				resource.Attribute{
					Name:        "apigateway_enabled",
					Description: `(Optional) API Gateway Log Switch. Default true.`,
				},
				resource.Attribute{
					Name:        "apigateway_ttl",
					Description: `(Optional) API Gateway ttl. Default 180.`,
				},
				resource.Attribute{
					Name:        "nas_enabled",
					Description: `(Optional) Nas log switch. Default true.`,
				},
				resource.Attribute{
					Name:        "nas_ttl",
					Description: `(Optional) Nas centralized ttl. Default 180.`,
				},
				resource.Attribute{
					Name:        "cps_enabled",
					Description: `(Optional) Mobile push log switch. Default true.`,
				},
				resource.Attribute{
					Name:        "cps_ttl",
					Description: `(Optional) Mobile push ttl. Default 180.`,
				},
				resource.Attribute{
					Name:        "multi_account",
					Description: `(Optional) Multi-account configuration, please fill in multiple aliuid. ## Attributes Reference The following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_log_dashboard",
			Category:         "Log Service (SLS)",
			ShortDescription: `Provides a Alicloud Log Dashboard resource.`,
			Description:      ``,
			Keywords: []string{
				"log",
				"service",
				"sls",
				"dashboard",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_name",
					Description: `(Required, ForceNew) The name of the log project. It is the only in one Alicloud account.`,
				},
				resource.Attribute{
					Name:        "dashboard_name",
					Description: `(Required, ForceNew) The name of the Log Dashboard.`,
				},
				resource.Attribute{
					Name:        "char_list",
					Description: `(Required) Configuration of charts in the dashboard.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Dashboard alias. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Log Dashboard. It sames as its name. ## Import Log Dashboard can be imported using the id or name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_log_dashboard.example tf-project:tf-logstore:tf-dashboard ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Log Dashboard. It sames as its name. ## Import Log Dashboard can be imported using the id or name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_log_dashboard.example tf-project:tf-logstore:tf-dashboard ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_log_machine_group",
			Category:         "Log Service (SLS)",
			ShortDescription: `Provides a Alicloud log tail machine group resource.`,
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
					Description: `The machine group topic. ## Import Log machine group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_log_machine_group.example tf-log:tf-machine-group ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `The machine group topic. ## Import Log machine group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_log_machine_group.example tf-log:tf-machine-group ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_log_project",
			Category:         "Log Service (SLS)",
			ShortDescription: `Provides a Alicloud log project resource.`,
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
					Description: `(Required, ForceNew) The name of the log project. It is the only in one Alicloud account.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the log project. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the log project. It sames as its name.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Log project name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Log project description. ## Import Log project can be imported using the id or name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_log_project.example tf-log ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the log project. It sames as its name.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Log project name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Log project description. ## Import Log project can be imported using the id or name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_log_project.example tf-log ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_log_store",
			Category:         "Log Service (SLS)",
			ShortDescription: `Provides a Alicloud log store resource.`,
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
					Description: `(Optional) The data retention time (in days). Valid values: [1-3650]. Default to 30. Log store data will be stored permanently when the value is "3650".`,
				},
				resource.Attribute{
					Name:        "shard_count",
					Description: `(Optional) The number of shards in this log store. Default to 2. You can modify it by "Split" or "Merge" operations. [Refer to details](https://www.alibabacloud.com/help/doc-detail/28976.htm)`,
				},
				resource.Attribute{
					Name:        "auto_split",
					Description: `(Optional) Determines whether to automatically split a shard. Default to true.`,
				},
				resource.Attribute{
					Name:        "max_split_shard_count",
					Description: `(Optional) The maximum number of shards for automatic split, which is in the range of 1 to 64. You must specify this parameter when autoSplit is true.`,
				},
				resource.Attribute{
					Name:        "append_meta",
					Description: `(Optional) Determines whether to append log meta automatically. The meta includes log receive time and client IP address. Default to true.`,
				},
				resource.Attribute{
					Name:        "enable_web_tracking",
					Description: `(Optional) Determines whether to enable Web Tracking. Default false. ## Attributes Reference The following attributes are exported:`,
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
					Description: `Determines whether to enable Web Tracking. ## Import Log store can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_log_store.example tf-log:tf-log-store ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `Determines whether to enable Web Tracking. ## Import Log store can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_log_store.example tf-log:tf-log-store ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_log_store_index",
			Category:         "Log Service (SLS)",
			ShortDescription: `Provides a Alicloud log store index resource.`,
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
					Description: `The ID of the log store index. It formats of ` + "`" + `<project>:<logstore>` + "`" + `. ## Import Log store index can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_log_store_index.example tf-log:tf-log-store ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the log store index. It formats of ` + "`" + `<project>:<logstore>` + "`" + `. ## Import Log store index can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_log_store_index.example tf-log:tf-log-store ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_logtail_attachment",
			Category:         "Log Service (SLS)",
			ShortDescription: `Provides a Alicloud logtail attachment resource.`,
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
					Description: `The ID of the logtail to machine group. It formats of ` + "`" + `<project>:<logtail_config_name>:<machine_group_name>` + "`" + `. ## Import Logtial to machine group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_logtail_to_machine_group.example tf-log:tf-log-config:tf-log-machine-group ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the logtail to machine group. It formats of ` + "`" + `<project>:<logtail_config_name>:<machine_group_name>` + "`" + `. ## Import Logtial to machine group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_logtail_to_machine_group.example tf-log:tf-log-config:tf-log-machine-group ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_logtail_config",
			Category:         "Log Service (SLS)",
			ShortDescription: `Provides a Alicloud logtail config resource.`,
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
					Description: `The ID of the log store index. It formats of ` + "`" + `<project>:<logstore>:<config_name>` + "`" + `. ## Import Logtial config can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_logtail_config.example tf-log:tf-log-store:tf-log-config ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the log store index. It formats of ` + "`" + `<project>:<logstore>:<config_name>` + "`" + `. ## Import Logtial config can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_logtail_config.example tf-log:tf-log-store:tf-log-config ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_market_order",
			Category:         "Market Place",
			ShortDescription: `Provides a market order resource.`,
			Description:      ``,
			Keywords: []string{
				"market",
				"place",
				"order",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "product_code",
					Description: `(Required, ForceNew) The product_code of market place product.`,
				},
				resource.Attribute{
					Name:        "pay_type",
					Description: `(Optional, ForceNew) Valid values are ` + "`" + `PrePaid` + "`" + `, ` + "`" + `PostPaid` + "`" + `,System default to ` + "`" + `PostPaid` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `(Optional, ForceNew) The number of purchase cycles.`,
				},
				resource.Attribute{
					Name:        "pricing_cycle",
					Description: `(Required, ForceNew) The purchase cycle of the product, valid values are ` + "`" + `Day` + "`" + `, ` + "`" + `Month` + "`" + ` and ` + "`" + `Year` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "package_version",
					Description: `(Required, ForceNew) The package version of the market product.`,
				},
				resource.Attribute{
					Name:        "quantity",
					Description: `(Optional, ForceNew) The quantity of the market product will be purchased.`,
				},
				resource.Attribute{
					Name:        "coupon_id",
					Description: `(Optional, ForceNew) The coupon id of the market product.`,
				},
				resource.Attribute{
					Name:        "components",
					Description: `(Optional, ForceNew) Service providers customize additional components. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the market order. ## Import Market order can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_market_order.order your-order-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the market order. ## Import Market order can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_market_order.order your-order-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_maxcompute_project",
			Category:         "MaxCompute Resources",
			ShortDescription: `Provides a Alicloud maxcompute project resource.`,
			Description:      ``,
			Keywords: []string{
				"maxcompute",
				"project",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) The name of the maxcompute project.`,
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
			Type:             "alicloud_mns_queue",
			Category:         "Message Notification Service (MNS)",
			ShortDescription: `Provides a Alicloud MNS Queue resource.`,
			Description:      ``,
			Keywords: []string{
				"message",
				"notification",
				"service",
				"mns",
				"queue",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForcesNew)Two queues on a single account in the same region cannot have the same name. A queue name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 256 characters .`,
				},
				resource.Attribute{
					Name:        "delay_seconds",
					Description: `(Optional)This attribute defines the length of time, in seconds, after which every message sent to the queue is dequeued. Valid value range: 0-604800 seconds, i.e., 0 to 7 days. Default value to 0.`,
				},
				resource.Attribute{
					Name:        "maximum_message_size",
					Description: `(Optional)This indicates the maximum length, in bytes, of any message body sent to the queue. Valid value range: 1024-65536, i.e., 1K to 64K. Default value to 65536.`,
				},
				resource.Attribute{
					Name:        "message_retention_period",
					Description: `(Optional) Messages are deleted from the queue after a specified length of time, whether they have been activated or not. This attribute defines the viability period, in seconds, for every message in the queue. Valid value range: 60-604800 seconds, i.e., 1 minutes to 7 days. Default value to 345600.`,
				},
				resource.Attribute{
					Name:        "visibility_timeout",
					Description: `(Optional) The VisibilityTimeout attribute of the queue. A dequeued messages will change from active (visible) status to inactive (invisible) status, and this attribute defines the length of time, in seconds, that messages remain invisible. Messages return to active status after the set period. Valid value range: 1-43200 seconds, i.e., 1 seconds to 12 hours. Default value to 30.`,
				},
				resource.Attribute{
					Name:        "polling_wait_seconds",
					Description: `(Optional) Long polling is measured in seconds. When this attribute is set to 0, long polling is disabled. When it is not set to 0, long polling is enabled and message dequeue requests will be processed only when valid messages are received or when long polling times out. Valid value range: 0-30 seconds. Default value to 0. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the queue is equal to name. ## Import MNS QUEUE can be imported using the id or name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_mns_queue.queue queuename ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the queue is equal to name. ## Import MNS QUEUE can be imported using the id or name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_mns_queue.queue queuename ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_mns_topic",
			Category:         "Message Notification Service (MNS)",
			ShortDescription: `Provides a Alicloud MNS Topic resource.`,
			Description:      ``,
			Keywords: []string{
				"message",
				"notification",
				"service",
				"mns",
				"topic",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew)Two topics on a single account in the same region cannot have the same name. A topic name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 256 characters.`,
				},
				resource.Attribute{
					Name:        "maximum_message_size",
					Description: `(Optional)This indicates the maximum length, in bytes, of any message body sent to the topic. Valid value range: 1024-65536, i.e., 1K to 64K. Default value to 65536.`,
				},
				resource.Attribute{
					Name:        "logging_enabled",
					Description: `(Optional) Is logging enabled? true or false. Default value to false. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the topic is equal to name. ## Import MNS Topic can be imported using the id or name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_mns_topic.topic topicName ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the topic is equal to name. ## Import MNS Topic can be imported using the id or name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_mns_topic.topic topicName ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_mns_topic_subscription",
			Category:         "Message Notification Service (MNS)",
			ShortDescription: `Provides a Alicloud MNS Topic Subscription resource.`,
			Description:      ``,
			Keywords: []string{
				"message",
				"notification",
				"service",
				"mns",
				"topic",
				"subscription",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Two topics subscription on a single account in the same topic cannot have the same name. A topic subscription name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 256 characters.`,
				},
				resource.Attribute{
					Name:        "notify_strategy",
					Description: `(Optional) The NotifyStrategy attribute of Subscription. This attribute specifies the retry strategy when message sending fails. the attribute has two value EXPONENTIAL_DECAY_RETR or BACKOFF_RETRY. Default value to BACKOFF_RETRY .`,
				},
				resource.Attribute{
					Name:        "notify_content_format",
					Description: `(Optional, ForceNew) The NotifyContentFormat attribute of Subscription. This attribute specifies the content format of the messages pushed to users. The valid values: 'SIMPLIFIED', 'XML' and 'JSON'. Default to 'SIMPLIFIED'.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Required, ForceNew) The endpoint has three format. Available values format: - HTTP Format: http://xxx.com/xxx - Queue Format: acs:mns:{REGION}:{AccountID}:queues/{QueueName} - Email Format: mail:directmail:{MailAddress}`,
				},
				resource.Attribute{
					Name:        "filter_tag",
					Description: `(Optional, ForceNew) The length should be shorter than 16. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the topic subscription.Format to topic_name:name ## Import MNS Topic subscription can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_mns_topic_subscription.subscription tf-example-mnstopic:tf-example-mnstopic-sub ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the topic subscription.Format to topic_name:name ## Import MNS Topic subscription can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_mns_topic_subscription.subscription tf-example-mnstopic:tf-example-mnstopic-sub ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_mongodb_instance",
			Category:         "MongoDB",
			ShortDescription: `Provides a MongoDB instance resource.`,
			Description:      ``,
			Keywords: []string{
				"mongodb",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "engine_version",
					Description: `(Required, ForceNew) Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/doc-detail/61763.htm) ` + "`" + `EngineVersion` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "db_instance_class",
					Description: `(Required) Instance specification. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/57141.htm).`,
				},
				resource.Attribute{
					Name:        "db_instance_storage",
					Description: `(Required) User-defined DB instance storage space.Unit: GB. Value range: - Custom storage space; value range: [10,2000] - 10-GB increments.`,
				},
				resource.Attribute{
					Name:        "replication_factor",
					Description: `(Optional) Number of replica set nodes. Valid values: [3, 5, 7]`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of DB instance. It a string of 2 to 256 characters.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `(Optional) Valid values are ` + "`" + `PrePaid` + "`" + `, ` + "`" + `PostPaid` + "`" + `, System default to ` + "`" + `PostPaid` + "`" + `. It can be modified from ` + "`" + `PostPaid` + "`" + ` to ` + "`" + `PrePaid` + "`" + ` after version 1.63.0.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) The duration that you will buy DB instance (in month). It is valid when instance_charge_type is ` + "`" + `PrePaid` + "`" + `. Valid values: [1~9], 12, 24, 36. System default to 1.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Optional, ForceNew) The Zone to launch the DB instance. it supports multiple zone. If it is a multi-zone and ` + "`" + `vswitch_id` + "`" + ` is specified, the vswitch must in one of them. The multiple zone ID can be retrieved by setting ` + "`" + `multi` + "`" + ` to "true" in the data source ` + "`" + `alicloud_zones` + "`" + `.`,
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
					Description: `(Optional) List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Optional, Available in 1.73.0+) The Security Group ID of ECS.`,
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
					Name:        "tde_status",
					Description: `(Optional, ForceNew, Available in 1.73.0+) The TDE(Transparent Data Encryption) status.`,
				},
				resource.Attribute{
					Name:        "maintain_start_time",
					Description: `(Optional, Available in v1.56.0+) The start time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).`,
				},
				resource.Attribute{
					Name:        "maintain_end_time",
					Description: `(Optional, Available in v1.56.0+) The end time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time).`,
				},
				resource.Attribute{
					Name:        "ssl_action",
					Description: `(Optional, Available in v1.78.0+) Actions performed on SSL functions, Valid values: ` + "`" + `Open` + "`" + `: turn on SSL encryption; ` + "`" + `Close` + "`" + `: turn off SSL encryption; ` + "`" + `Update` + "`" + `: update SSL certificate.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Available in v1.66.0+) A mapping of tags to assign to the resource. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the MongoDB.`,
				},
				resource.Attribute{
					Name:        "retention_period",
					Description: `Instance log backup retention days. Available in 1.42.0+.`,
				},
				resource.Attribute{
					Name:        "replica_set_name",
					Description: `The name of the mongo replica set`,
				},
				resource.Attribute{
					Name:        "ssl_status",
					Description: `Status of the SSL feature. ` + "`" + `Open` + "`" + `: SSL is turned on; ` + "`" + `Closed` + "`" + `: SSL is turned off. ### Timeouts ->`,
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
					Description: `(Defaults to 30 mins) Used when terminating the MongoDB instance. ## Import MongoDB can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_mongodb_instance.example dds-bp1291daeda44194 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the MongoDB.`,
				},
				resource.Attribute{
					Name:        "retention_period",
					Description: `Instance log backup retention days. Available in 1.42.0+.`,
				},
				resource.Attribute{
					Name:        "replica_set_name",
					Description: `The name of the mongo replica set`,
				},
				resource.Attribute{
					Name:        "ssl_status",
					Description: `Status of the SSL feature. ` + "`" + `Open` + "`" + `: SSL is turned on; ` + "`" + `Closed` + "`" + `: SSL is turned off. ### Timeouts ->`,
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
					Description: `(Defaults to 30 mins) Used when terminating the MongoDB instance. ## Import MongoDB can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_mongodb_instance.example dds-bp1291daeda44194 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodb_sharding_instance",
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
					Description: `(Optional, ForceNew) Valid values are ` + "`" + `PrePaid` + "`" + `, ` + "`" + `PostPaid` + "`" + `,System default to ` + "`" + `PostPaid` + "`" + `.`,
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
					Name:        "backup_period",
					Description: `(Optional, Available in 1.42.0+) MongoDB Instance backup period. It is required when ` + "`" + `backup_time` + "`" + ` was existed. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]. Default to [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]`,
				},
				resource.Attribute{
					Name:        "backup_time",
					Description: `(Optional, Available in 1.42.0+) MongoDB instance backup time. It is required when ` + "`" + `backup_period` + "`" + ` was existed. In the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. If not set, the system will return a default, like "23:00Z-24:00Z". ## Attributes Reference The following attributes are exported:`,
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
					Description: `Instance log backup retention days. Available in 1.42.0+. ## Import MongoDB can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_mongodb_sharding_instance.example dds-bp1291daeda44195 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `Instance log backup retention days. Available in 1.42.0+. ## Import MongoDB can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_mongodb_sharding_instance.example dds-bp1291daeda44195 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_nas_access_group",
			Category:         "Network Attached Storage (NAS)",
			ShortDescription: `Provides a Alicloud Nas Access Group resource.`,
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
					Description: `(Required, ForceNew) A Name of one Access Group.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required, ForceNew) A Type of one Access Group. Valid values: ` + "`" + `Vpc` + "`" + ` and ` + "`" + `Classic` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The Access Group description. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Access Group. ## Import Nas Access Group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_nas_access_group.foo tf_testAccNasConfig ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Access Group. ## Import Nas Access Group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_nas_access_group.foo tf_testAccNasConfig ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_nas_access_rule",
			Category:         "Network Attached Storage (NAS)",
			ShortDescription: `Provides a Alicloud Nas Access Rule resource.`,
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
					Description: `(Optional) Read-write permission type: RDWR (default), RDONLY.`,
				},
				resource.Attribute{
					Name:        "user_access_type",
					Description: `(Optional) User permission type: no_squash (default), root_squash, all_squash.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Priority level. Range: 1-100. Default value: 1. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `This ID of this resource. The value is formate as ` + "`" + `<access_group_name>:<access rule id>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "access_rule_id",
					Description: `The nas access rule ID. ## Import Nas Access Rule can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_nas_access_rule.foo tf-testAccNasConfigName:1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `This ID of this resource. The value is formate as ` + "`" + `<access_group_name>:<access rule id>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "access_rule_id",
					Description: `The nas access rule ID. ## Import Nas Access Rule can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_nas_access_rule.foo tf-testAccNasConfigName:1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_nas_file_system",
			Category:         "Network Attached Storage (NAS)",
			ShortDescription: `Provides a Alicloud NAS File System resource.`,
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
					Name:        "protocol_type",
					Description: `(Required, ForceNew) The Protocol Type of a File System. Valid values: ` + "`" + `NFS` + "`" + ` and ` + "`" + `SMB` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `(Required, ForceNew) The Storage Type of a File System. Valid values: ` + "`" + `Capacity` + "`" + ` and ` + "`" + `Performance` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The File System description. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the File System. ## Import Nas File System can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_nas_file_system.foo 1337849c59 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the File System. ## Import Nas File System can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_nas_file_system.foo 1337849c59 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_nas_mount_target",
			Category:         "Network Attached Storage (NAS)",
			ShortDescription: `Provides a Alicloud Nas MountTarget resource.`,
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
					Description: `(Required, ForceNew) File system ID.`,
				},
				resource.Attribute{
					Name:        "access_group_name",
					Description: `(Required) Permission group name.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Optional, ForceNew) VSwitch ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Whether the MountTarget is active. An inactive MountTarget is inusable. Valid values are Active(default) and Inactive. ## Attributes Reference The following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_nat_gateway",
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
					Name:        "spec",
					Description: `(Deprecated) It has been deprecated from provider version 1.7.1, and new field 'specification' can replace it.`,
				},
				resource.Attribute{
					Name:        "specification",
					Description: `(Optional) The specification of the nat gateway. Valid values are ` + "`" + `Small` + "`" + `, ` + "`" + `Middle` + "`" + ` and ` + "`" + `Large` + "`" + `. Default to ` + "`" + `Small` + "`" + `. Details refer to [Nat Gateway Specification](https://www.alibabacloud.com/help/doc-detail/42757.htm).`,
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
					Description: `(Optional) A list of bandwidth packages for the nat gatway. Only support nat gateway created before 00:00 on November 4, 2017. Available in v1.13.0+ and v1.7.1-.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `(Optional, ForceNew, Available in 1.45.0+) The billing method of the nat gateway. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid".`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional, ForceNew, Available in 1.45.0+) The duration that you will buy the resource, in month. It is valid when ` + "`" + `instance_charge_type` + "`" + ` is ` + "`" + `PrePaid` + "`" + `. Default to 1. Valid values: [1-9, 12, 24, 36]. At present, the provider does not support modify "period" and you can do that via web console. ## Block bandwidth packages The bandwidth package mapping supports the following:`,
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
					Name:        "spec",
					Description: `It has been deprecated from provider version 1.7.1.`,
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
					Description: `The nat gateway will auto create a snap and forward item, the ` + "`" + `forward_table_ids` + "`" + ` is the created one. ## Import Nat gateway can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_nat_gateway.example ngw-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "spec",
					Description: `It has been deprecated from provider version 1.7.1.`,
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
					Description: `The nat gateway will auto create a snap and forward item, the ` + "`" + `forward_table_ids` + "`" + ` is the created one. ## Import Nat gateway can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_nat_gateway.example ngw-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_network_acl",
			Category:         "VPC",
			ShortDescription: `Provides a Alicloud Network Acl resource.`,
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
					Description: `(Optional) The name of the network acl.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the network acl instance. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the network acl instance id. ## Import The network acl can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_network_acl.default nacl-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the network acl instance id. ## Import The network acl can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_network_acl.default nacl-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_network_acl_attachment",
			Category:         "VPC",
			ShortDescription: `Provides a Alicloud Network Acl Attachment resource.`,
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
			Type:             "alicloud_network_acl_entries",
			Category:         "VPC",
			ShortDescription: `Provides a Alicloud Network Acl Entries resource.`,
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
			Type:             "alicloud_network_interface",
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
					Description: `(Require) A list of security group ids to associate with.`,
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
					Description: `(Optional) A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(ForceNew, ForceNew, Available in 1.57.0+) The Id of resource group which the network interface belongs. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ENI ID.`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `(Available in 1.54.0+) The MAC address of an ENI. ## Import ENI can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_network_interface.eni eni-abc1234567890000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ENI ID.`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `(Available in 1.54.0+) The MAC address of an ENI. ## Import ENI can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_network_interface.eni eni-abc1234567890000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_network_interface_attachment",
			Category:         "ECS",
			ShortDescription: `Provides an Alicloud ECS Elastic Network Interface Attachment as a resource to attach ENI to or detach ENI from ECS Instances.`,
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
					Description: `The ID of the resource, formatted as ` + "`" + `<network_interface_id>:<instance_id>` + "`" + `. ## Import Network Interfaces Attachment resource can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_network_interface_attachment.eni eni-abc123456789000:i-abc123456789000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the resource, formatted as ` + "`" + `<network_interface_id>:<instance_id>` + "`" + `. ## Import Network Interfaces Attachment resource can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_network_interface_attachment.eni eni-abc123456789000:i-abc123456789000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ons_group",
			Category:         "RocketMQ",
			ShortDescription: `Provides a Alicloud ONS Group resource.`,
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
					Description: `(Optional) This attribute is a concise description of group. The length cannot exceed 256.`,
				},
				resource.Attribute{
					Name:        "read_enable",
					Description: `(Optional) This attribute is used to set the message reading enabled or disabled. It can only be set after the group is used by the client. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `key` + "`" + ` of the resource supplied above. The value is formulated as ` + "`" + `<instance_id>:<group_id>` + "`" + `. ## Import ONS GROUP can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ons_group.group MQ_INST_1234567890_Baso1234567:GID-onsGroupDemo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `key` + "`" + ` of the resource supplied above. The value is formulated as ` + "`" + `<instance_id>:<group_id>` + "`" + `. ## Import ONS GROUP can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ons_group.group MQ_INST_1234567890_Baso1234567:GID-onsGroupDemo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ons_instance",
			Category:         "RocketMQ",
			ShortDescription: `Provides a Alicloud ONS Instance resource.`,
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
					Description: `Platinum edition instance expiration time. ## Import ONS INSTANCE can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ons_instance.instance MQ_INST_1234567890_Baso1234567 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `Platinum edition instance expiration time. ## Import ONS INSTANCE can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ons_instance.instance MQ_INST_1234567890_Baso1234567 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ons_topic",
			Category:         "RocketMQ",
			ShortDescription: `Provides a Alicloud ONS Topic resource.`,
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
					Description: `(Required) The type of the message. Read [Ons Topic Create](https://www.alibabacloud.com/help/doc-detail/29591.html) for further details.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `(Optional) This attribute is a concise description of topic. The length cannot exceed 128.`,
				},
				resource.Attribute{
					Name:        "perm",
					Description: `(Optional) This attribute is used to set the read-write mode for the topic. Read [Request parameters](https://www.alibabacloud.com/help/doc-detail/56880.html) for further details. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `key` + "`" + ` of the resource supplied above. The value is formulated as ` + "`" + `<instance_id>:<topic>` + "`" + `. ## Import ONS TOPIC can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ons_topic.topic MQ_INST_1234567890_Baso1234567:onsTopicDemo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `key` + "`" + ` of the resource supplied above. The value is formulated as ` + "`" + `<instance_id>:<topic>` + "`" + `. ## Import ONS TOPIC can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ons_topic.topic MQ_INST_1234567890_Baso1234567:onsTopicDemo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_oss_bucket",
			Category:         "OSS",
			ShortDescription: `Provides a resource to create a oss bucket.`,
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
					Description: `(Optional) The [canned ACL](https://www.alibabacloud.com/help/doc-detail/31898.htm) to apply. Can be "private", "public-read" and "public-read-write". Defaults to "private".`,
				},
				resource.Attribute{
					Name:        "cors_rule",
					Description: `(Optional) A rule of [Cross-Origin Resource Sharing](https://www.alibabacloud.com/help/doc-detail/31903.htm) (documented below). The items of core rule are no more than 10 for every OSS bucket.`,
				},
				resource.Attribute{
					Name:        "website",
					Description: `(Optional) A website object(documented below).`,
				},
				resource.Attribute{
					Name:        "logging",
					Description: `(Optional) A Settings of [bucket logging](https://www.alibabacloud.com/help/doc-detail/31900.htm) (documented below).`,
				},
				resource.Attribute{
					Name:        "logging_isenable",
					Description: `(Optional) The flag of using logging enable container. Defaults true.`,
				},
				resource.Attribute{
					Name:        "referer_config",
					Description: `(Optional) The configuration of [referer](https://www.alibabacloud.com/help/doc-detail/31901.htm) (documented below).`,
				},
				resource.Attribute{
					Name:        "lifecycle_rule",
					Description: `(Optional) A configuration of [object lifecycle management](https://www.alibabacloud.com/help/doc-detail/31904.htm) (documented below).`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Optional, Available in 1.41.0) Json format text of bucket policy [bucket policy management](https://www.alibabacloud.com/help/doc-detail/100680.htm).`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `(Optional, ForceNew) The [storage class](https://www.alibabacloud.com/help/doc-detail/51374.htm) to apply. Can be "Standard", "IA" and "Archive". Defaults to "Standard".`,
				},
				resource.Attribute{
					Name:        "server_side_encryption_rule",
					Description: `(Optional, Available in 1.45.0+) A configuration of server-side encryption (documented below).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Available in 1.45.0+) A mapping of tags to assign to the bucket. The items are no more than 10 for a bucket.`,
				},
				resource.Attribute{
					Name:        "versioning",
					Description: `(Optional, Available in 1.45.0+) A state of versioning (documented below).`,
				},
				resource.Attribute{
					Name:        "force_destroy",
					Description: `(Optional, Available in 1.45.0+) A boolean that indicates all objects should be deleted from the bucket so that the bucket can be destroyed without error. These objects are not recoverable. Defaults to "false". #### Block cors_rule The cors_rule mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "allowed_headers",
					Description: `(Optional) Specifies which headers are allowed.`,
				},
				resource.Attribute{
					Name:        "allowed_methods",
					Description: `(Required) Specifies which methods are allowed. Can be GET, PUT, POST, DELETE or HEAD.`,
				},
				resource.Attribute{
					Name:        "allowed_origins",
					Description: `(Required) Specifies which origins are allowed.`,
				},
				resource.Attribute{
					Name:        "expose_headers",
					Description: `(Optional) Specifies expose header in the response.`,
				},
				resource.Attribute{
					Name:        "max_age_seconds",
					Description: `(Optional) Specifies time in seconds that browser can cache the response for a preflight request. #### Block website The website mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "index_document",
					Description: `(Required) Alicloud OSS returns this index document when requests are made to the root domain or any of the subfolders.`,
				},
				resource.Attribute{
					Name:        "error_document",
					Description: `(Optional) An absolute path to the document to return in case of a 4XX error. #### Block logging The logging object supports the following:`,
				},
				resource.Attribute{
					Name:        "target_bucket",
					Description: `(Required) The name of the bucket that will receive the log objects.`,
				},
				resource.Attribute{
					Name:        "target_prefix",
					Description: `(Optional) To specify a key prefix for log objects. #### Block referer configuration The referer configuration supports the following:`,
				},
				resource.Attribute{
					Name:        "allow_empty",
					Description: `(Optional, Type: bool) Allows referer to be empty. Defaults false.`,
				},
				resource.Attribute{
					Name:        "referers",
					Description: `(Required, Type: list) The list of referer. #### Block lifecycle_rule The lifecycle_rule object supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Unique identifier for the rule. If omitted, OSS bucket will assign a unique name.`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Optional, Available in v1.90.0+) Object key prefix identifying one or more objects to which the rule applies. Default value is null, the rule applies to all objects in a bucket.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required, Type: bool) Specifies lifecycle rule status.`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `(Optional, Type: set) Specifies a period in the object's expire (documented below).`,
				},
				resource.Attribute{
					Name:        "transitions",
					Description: `(Optional, Type: set, Available in 1.62.1+) Specifies the time when an object is converted to the IA or archive storage class during a valid life cycle. (documented below). ` + "`" + `NOTE` + "`" + `: At least one of expiration and transitions should be configured. #### Block expiration The lifecycle_rule expiration object supports the following:`,
				},
				resource.Attribute{
					Name:        "date",
					Description: `(Optional) Specifies the date after which you want the corresponding action to take effect. The value obeys ISO8601 format like ` + "`" + `2017-03-09` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "days",
					Description: `(Optional, Type: int) Specifies the number of days after object creation when the specific rule action takes effect. ` + "`" + `NOTE` + "`" + `: One and only one of "date" and "days" can be specified in one expiration configuration. #### Block transitions The lifecycle_rule transitions object supports the following:`,
				},
				resource.Attribute{
					Name:        "created_before_date",
					Description: `(Optional) Specifies the time before which the rules take effect. The date must conform to the ISO8601 format and always be UTC 00:00. For example: 2002-10-11T00:00:00.000Z indicates that objects updated before 2002-10-11T00:00:00.000Z are deleted or converted to another storage class, and objects updated after this time (including this time) are not deleted or converted.`,
				},
				resource.Attribute{
					Name:        "days",
					Description: `(Optional, Type: int) Specifies the number of days after object creation when the specific rule action takes effect.`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `(Required) Specifies the storage class that objects that conform to the rule are converted into. The storage class of the objects in a bucket of the IA storage class can be converted into Archive but cannot be converted into Standard. Values: ` + "`" + `IA` + "`" + `, ` + "`" + `Archive` + "`" + `. ` + "`" + `NOTE` + "`" + `: One and only one of "created_before_date" and "days" can be specified in one transition configuration. #### Block server-side encryption rule The server-side encryption rule supports the following:`,
				},
				resource.Attribute{
					Name:        "sse_algorithm",
					Description: `(Required) The server-side encryption algorithm to use. Possible values: ` + "`" + `AES256` + "`" + ` and ` + "`" + `KMS` + "`" + `. #### Block versioning The versioning supports the following:`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Required) Specifies the versioning state of a bucket. Valid values: ` + "`" + `Enabled` + "`" + ` and ` + "`" + `Suspended` + "`" + `. ` + "`" + `NOTE` + "`" + `: Currently, the ` + "`" + `versioning` + "`" + ` feature is only available in ap-south-1 and with white list. If you want to use it, please contact us. ## Attributes Reference The following attributes are exported:`,
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
					Description: `The bucket owner. ## Import OSS bucket can be imported using the bucket name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_oss_bucket.bucket bucket-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `The bucket owner. ## Import OSS bucket can be imported using the bucket name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_oss_bucket.bucket bucket-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_oss_bucket_object",
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
					Name:        "content_length",
					Description: `the content length of request.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `the ETag generated for the object (an MD5 sum of the object content).`,
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
					Name:        "content_length",
					Description: `the content length of request.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `the ETag generated for the object (an MD5 sum of the object content).`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `A unique version ID value for the object, if bucket versioning is enabled.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ots_instance",
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
					Description: `The instance tags. ## Import OTS instance can be imported using instance id or name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ots_instance.foo "my-ots-instance" ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `The instance tags. ## Import OTS instance can be imported using instance id or name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ots_instance.foo "my-ots-instance" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ots_instance_attachment",
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
			Type:             "alicloud_ots_table",
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
					Description: `The max version offset of the table. ## Import OTS table can be imported using id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ots_table.table "my-ots:ots_table" ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `The max version offset of the table. ## Import OTS table can be imported using id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ots_table.table "my-ots:ots_table" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_polardb_account",
			Category:         "PolarDB",
			ShortDescription: `Provides a RDS account resource.`,
			Description:      ``,
			Keywords: []string{
				"polardb",
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
					Description: `(Required) Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters.`,
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
					Description: `(Optional) Account description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.`,
				},
				resource.Attribute{
					Name:        "account_type",
					Description: `(Optional, ForceNew) Account type, Valid values are ` + "`" + `Normal` + "`" + `, ` + "`" + `Super` + "`" + `, Default to ` + "`" + `Normal` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The current account resource ID. Composed of instance ID and account name with format ` + "`" + `<instance_id>:<name>` + "`" + `. ## Import PolarDB account can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_polardb_account.example "pc-12345:tf_account" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The current account resource ID. Composed of instance ID and account name with format ` + "`" + `<instance_id>:<name>` + "`" + `. ## Import PolarDB account can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_polardb_account.example "pc-12345:tf_account" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_polardb_account_privilege",
			Category:         "PolarDB",
			ShortDescription: `Provides a PolarDB account privilege resource.`,
			Description:      ``,
			Keywords: []string{
				"polardb",
				"account",
				"privilege",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "db_cluster_id",
					Description: `(Required, ForceNew) The Id of cluster in which account belongs.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `(Required, ForceNew) A specified account name.`,
				},
				resource.Attribute{
					Name:        "account_privilege",
					Description: `(Optional, ForceNew) The privilege of one account access database. Valid values: ["ReadOnly", "ReadWrite"]. Default to "ReadOnly".`,
				},
				resource.Attribute{
					Name:        "db_names",
					Description: `(Required) List of specified database name. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The current account resource ID. Composed of instance ID, account name and privilege with format ` + "`" + `<db_cluster_id>:<account_name>:<account_privilege>` + "`" + `. ## Import PolarDB account privilege can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_polardb_account_privilege.example "pc-12345:tf_account:ReadOnly" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The current account resource ID. Composed of instance ID, account name and privilege with format ` + "`" + `<db_cluster_id>:<account_name>:<account_privilege>` + "`" + `. ## Import PolarDB account privilege can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_polardb_account_privilege.example "pc-12345:tf_account:ReadOnly" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_polardb_backup_policy",
			Category:         "PolarDB",
			ShortDescription: `Provides a PolarDB backup policy resource.`,
			Description:      ``,
			Keywords: []string{
				"polardb",
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
					Description: `(Optional) PolarDB Cluster backup period. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]. Default to ["Tuesday", "Thursday", "Saturday"].`,
				},
				resource.Attribute{
					Name:        "preferred_backup_time",
					Description: `(Optional) PolarDB Cluster backup time, in the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. Default to "02:00Z-03:00Z". China time is 8 hours behind it. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The current backup policy resource ID. It is same as 'db_cluster_id'.`,
				},
				resource.Attribute{
					Name:        "backup_retention_period",
					Description: `Cluster backup retention days, Fixed for 7 days, not modified. ## Import PolarDB backup policy can be imported using the id or cluster id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_polardb_backup_policy.example "rm-12345678" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The current backup policy resource ID. It is same as 'db_cluster_id'.`,
				},
				resource.Attribute{
					Name:        "backup_retention_period",
					Description: `Cluster backup retention days, Fixed for 7 days, not modified. ## Import PolarDB backup policy can be imported using the id or cluster id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_polardb_backup_policy.example "rm-12345678" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_polardb_cluster",
			Category:         "PolarDB",
			ShortDescription: `Provides a PolarDB cluster resource.`,
			Description:      ``,
			Keywords: []string{
				"polardb",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "db_type",
					Description: `(Required,ForceNew) Database type. Value options: MySQL, Oracle, PostgreSQL.`,
				},
				resource.Attribute{
					Name:        "db_version",
					Description: `(Required,ForceNew) Database version. Value options can refer to the latest docs [CreateDBCluster](https://help.aliyun.com/document_detail/98169.html) ` + "`" + `DBVersion` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "db_node_class",
					Description: `(Required) The db_node_class of cluster node.`,
				},
				resource.Attribute{
					Name:        "modify_type",
					Description: `(Optional, Available in 1.71.2+) Use as ` + "`" + `db_node_class` + "`" + ` change class , define upgrade or downgrade. Valid values are ` + "`" + `Upgrade` + "`" + `, ` + "`" + `Downgrade` + "`" + `, Default to ` + "`" + `Upgrade` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Optional) The Zone to launch the DB cluster. it supports multiple zone.`,
				},
				resource.Attribute{
					Name:        "pay_type",
					Description: `(Optional,ForceNew) Valid values are ` + "`" + `PrePaid` + "`" + `, ` + "`" + `PostPaid` + "`" + `, Default to ` + "`" + `PostPaid` + "`" + `. Currently, the resource can not supports change pay type.`,
				},
				resource.Attribute{
					Name:        "renewal_status",
					Description: `(Optional) Valid values are ` + "`" + `AutoRenewal` + "`" + `, ` + "`" + `Normal` + "`" + `, ` + "`" + `NotRenewal` + "`" + `, Default to ` + "`" + `NotRenewal` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "auto_renew_period",
					Description: `(Optional) Auto-renewal period of an cluster, in the unit of the month. It is valid when pay_type is ` + "`" + `PrePaid` + "`" + `. Valid value:1, 2, 3, 6, 12, 24, 36, Default to 1.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) The duration that you will buy DB cluster (in month). It is valid when pay_type is ` + "`" + `PrePaid` + "`" + `. Valid values: [1~9], 12, 24, 36. Default to 1.`,
				},
				resource.Attribute{
					Name:        "security_ips",
					Description: `(Optional) List of IP addresses allowed to access all databases of an cluster. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Optional) The virtual switch ID to launch DB instances in one VPC.`,
				},
				resource.Attribute{
					Name:        "maintain_time",
					Description: `(Optional) Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of cluster.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `(Optional) Set of parameters needs to be set after DB cluster was launched. Available parameters can refer to the latest docs [View database parameter templates](https://www.alibabacloud.com/help/doc-detail/26284.htm) .`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Available in v1.68.0+) A mapping of tags to assign to the resource. - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string. - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The PolarDB cluster ID.`,
				},
				resource.Attribute{
					Name:        "connection_string",
					Description: `(Available in 1.81.0+) PolarDB cluster connection string. When security_ips is configured, the address of cluster type endpoint will be returned, and if only "127.0.0.1" is configured, it will also be an empty string. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 30 mins) Used when creating the polardb cluster (until it reaches the initial ` + "`" + `Running` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 30 mins) Used when updating the polardb cluster (until it reaches the initial ` + "`" + `Running` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 10 mins) Used when terminating the polardb cluster. ## Import PolarDB cluster can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_polardb_cluster.example pc-abc12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The PolarDB cluster ID.`,
				},
				resource.Attribute{
					Name:        "connection_string",
					Description: `(Available in 1.81.0+) PolarDB cluster connection string. When security_ips is configured, the address of cluster type endpoint will be returned, and if only "127.0.0.1" is configured, it will also be an empty string. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 30 mins) Used when creating the polardb cluster (until it reaches the initial ` + "`" + `Running` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 30 mins) Used when updating the polardb cluster (until it reaches the initial ` + "`" + `Running` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 10 mins) Used when terminating the polardb cluster. ## Import PolarDB cluster can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_polardb_cluster.example pc-abc12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_polardb_database",
			Category:         "PolarDB",
			ShortDescription: `Provides a PolarDB database resource.`,
			Description:      ``,
			Keywords: []string{
				"polardb",
				"database",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "db_cluster_id",
					Description: `(Required, ForceNew) The Id of cluster that can run database.`,
				},
				resource.Attribute{
					Name:        "db_name",
					Description: `(Required, ForceNew) Name of the database requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letterand have no more than 64 characters.`,
				},
				resource.Attribute{
					Name:        "character_set_name",
					Description: `(Optional,ForceNew) Character set. The value range is limited to the following: [ utf8, gbk, latin1, utf8mb4, Chinese_PRC_CI_AS, Chinese_PRC_CS_AS, SQL_Latin1_General_CP1_CI_AS, SQL_Latin1_General_CP1_CS_AS, Chinese_PRC_BIN ], default is "utf8" \(` + "`" + `utf8mb4` + "`" + ` only supports versions 5.5 and 5.6\).`,
				},
				resource.Attribute{
					Name:        "db_description",
					Description: `(Optional) Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The current database resource ID. Composed of cluster ID and database name with format ` + "`" + `<cluster_id>:<name>` + "`" + `. ## Import PolarDB database can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_polardb_database.example "pc-12345:tf_database" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The current database resource ID. Composed of cluster ID and database name with format ` + "`" + `<cluster_id>:<name>` + "`" + `. ## Import PolarDB database can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_polardb_database.example "pc-12345:tf_database" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_polardb_endpoint",
			Category:         "PolarDB",
			ShortDescription: `Provides a PolarDB instance endpoint resource.`,
			Description:      ``,
			Keywords: []string{
				"polardb",
				"endpoint",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "db_cluster_id",
					Description: `(Required, ForceNew) The Id of cluster that can run database.`,
				},
				resource.Attribute{
					Name:        "endpoint_type",
					Description: `(Required, ForceNew) Type of endpoint. Valid value: ` + "`" + `Custom` + "`" + `. Currently supported only ` + "`" + `Custom` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "read_write_mode",
					Description: `(Optional) Read or write mode. Valid values are ` + "`" + `ReadWrite` + "`" + `, ` + "`" + `ReadOnly` + "`" + `. Default to ` + "`" + `ReadOnly` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `(Optional) Node id list for endpoint configuration. At least 2 nodes if specified, or if the cluster has more than 3 nodes, read-only endpoint is allowed to mount only one node. Default is all nodes.`,
				},
				resource.Attribute{
					Name:        "auto_add_new_nodes",
					Description: `(Optional) Whether the new node automatically joins the default cluster address. Valid values are ` + "`" + `Enable` + "`" + `, ` + "`" + `Disable` + "`" + `. Default to ` + "`" + `Disable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "endpoint_config",
					Description: `(Optional) Advanced configuration of the cluster address. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The current instance connection resource ID. Composed of instance ID and connection string with format ` + "`" + `<db_cluster_id>:<db_endpoint_id>` + "`" + `. ## Import PolarDB endpoint can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_polardb_endpoint.example pc-abc123456:pe-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The current instance connection resource ID. Composed of instance ID and connection string with format ` + "`" + `<db_cluster_id>:<db_endpoint_id>` + "`" + `. ## Import PolarDB endpoint can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_polardb_endpoint.example pc-abc123456:pe-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_polardb_endpoint_address",
			Category:         "PolarDB",
			ShortDescription: `Provides a PolarDB instance endpoint resource.`,
			Description:      ``,
			Keywords: []string{
				"polardb",
				"endpoint",
				"address",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "db_cluster_id",
					Description: `(Required, ForceNew) The Id of cluster that can run database.`,
				},
				resource.Attribute{
					Name:        "db_endpoint_id",
					Description: `(Required, ForceNew) The Id of endpoint that can run database.`,
				},
				resource.Attribute{
					Name:        "connection_prefix",
					Description: `(Optional) Prefix of an Internet connection string. It must be checked for uniqueness. It may consist of lowercase letters, numbers, and underlines, and must start with a letter and have no more than 30 characters. Default to <db_endpoint_id> + 'tf'.`,
				},
				resource.Attribute{
					Name:        "net_type",
					Description: `(Optional, ForceNew) Internet connection net type. Valid value: ` + "`" + `Public` + "`" + `. Default to ` + "`" + `Public` + "`" + `. Currently supported only ` + "`" + `Public` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The current instance connection resource ID. Composed of instance ID and connection string with format ` + "`" + `<db_cluster_id>:<db_endpoint_id>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Connection cluster or endpoint port.`,
				},
				resource.Attribute{
					Name:        "connection_string",
					Description: `Connection cluster or endpoint string.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The ip address of connection string. ## Import PolarDB endpoint address can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_polardb_endpoint_address.example pc-abc123456:pe-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The current instance connection resource ID. Composed of instance ID and connection string with format ` + "`" + `<db_cluster_id>:<db_endpoint_id>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Connection cluster or endpoint port.`,
				},
				resource.Attribute{
					Name:        "connection_string",
					Description: `Connection cluster or endpoint string.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The ip address of connection string. ## Import PolarDB endpoint address can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_polardb_endpoint_address.example pc-abc123456:pe-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_pvtz_zone",
			Category:         "Private Zone",
			ShortDescription: `Provides a Alicloud Private Zone resource.`,
			Description:      ``,
			Keywords: []string{
				"private",
				"zone",
				"pvtz",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, ForceNew) The name of the Private Zone.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `(Optional) The remark of the Private Zone.`,
				},
				resource.Attribute{
					Name:        "proxy_pattern",
					Description: `(Optional, Available in 1.69.0+) The recursive DNS proxy. Valid values: - ZONE: indicates that the recursive DNS proxy is disabled. - RECORD: indicates that the recursive DNS proxy is enabled. Default to "ZONE"`,
				},
				resource.Attribute{
					Name:        "user_client_ip",
					Description: `(Optional, Available in 1.69.0+) The IP address of the client.`,
				},
				resource.Attribute{
					Name:        "lang",
					Description: `(Optional, Available in 1.69.0+) The language. Valid values: "zh", "en", "jp".`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional, ForceNew, Available in v1.86.0+) The Id of resource group which the Private Zone belongs. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Private Zone.`,
				},
				resource.Attribute{
					Name:        "record_count",
					Description: `The count of the Private Zone Record. ## Import Private Zone can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_pvtz_zone.example abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Private Zone.`,
				},
				resource.Attribute{
					Name:        "record_count",
					Description: `The count of the Private Zone Record. ## Import Private Zone can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_pvtz_zone.example abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_pvtz_zone_attachment",
			Category:         "Private Zone",
			ShortDescription: `Provides vpcs bound to Alicloud Private Zone resource.`,
			Description:      ``,
			Keywords: []string{
				"private",
				"zone",
				"pvtz",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Required, ForceNew) The name of the Private Zone Record.`,
				},
				resource.Attribute{
					Name:        "vpc_ids",
					Description: `(Optional, Conflict with ` + "`" + `vpcs` + "`" + `) The id List of the VPC with the same region, for example:["vpc-1","vpc-2"].`,
				},
				resource.Attribute{
					Name:        "vpcs",
					Description: `(Optional, Conflict with ` + "`" + `vpc_ids` + "`" + `, Available in 1.62.1+) The List of the VPC:`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) The Id of the vpc.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `(Option) The region of the vpc. If not set, the current region will instead of. Recommend to use ` + "`" + `vpcs` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lang",
					Description: `(Optional, Available in 1.62.1+) The language of code.`,
				},
				resource.Attribute{
					Name:        "user_client_ip",
					Description: `(Optional, Available in 1.62.1+) The user custom IP address. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Private Zone VPC Attachment. It sames with ` + "`" + `zone_id` + "`" + `. ## Import Private Zone attachment can be imported using the id(same with ` + "`" + `zone_id` + "`" + `), e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_pvtz_zone_attachment.example abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Private Zone VPC Attachment. It sames with ` + "`" + `zone_id` + "`" + `. ## Import Private Zone attachment can be imported using the id(same with ` + "`" + `zone_id` + "`" + `), e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_pvtz_zone_attachment.example abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_pvtz_zone_record",
			Category:         "Private Zone",
			ShortDescription: `Provides a Alicloud Private Zone Record resource.`,
			Description:      ``,
			Keywords: []string{
				"private",
				"zone",
				"pvtz",
				"record",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Required, ForceNew) The name of the Private Zone Record.`,
				},
				resource.Attribute{
					Name:        "resource_record",
					Description: `(Required, ForceNew) The resource record of the Private Zone Record.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of the Private Zone Record. Valid values: A, CNAME, TXT, MX, PTR.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of the Private Zone Record.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The ttl of the Private Zone Record.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) The priority of the Private Zone Record. At present, only can "MX" record support it. Valid values: [1-50]. Default to 1. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `This ID of this resource. The value is formate as ` + "`" + `<record_id>:<zone_id>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "record_id",
					Description: `The Private Zone Record ID. ## Import Private Zone Record can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_pvtz_zone_record.example abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `This ID of this resource. The value is formate as ` + "`" + `<record_id>:<zone_id>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "record_id",
					Description: `The Private Zone Record ID. ## Import Private Zone Record can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_pvtz_zone_record.example abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ram_access_key",
			Category:         "RAM",
			ShortDescription: `Provides a RAM User access key resource.`,
			Description:      ``,
			Keywords: []string{
				"ram",
				"access",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "user_name",
					Description: `(Optional, ForceNew) Name of the RAM user. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin with a hyphen.`,
				},
				resource.Attribute{
					Name:        "secret_file",
					Description: `(Optional, ForceNew) The name of file that can save access key id and access key secret. Strongly suggest you to specified it when you creating access key, otherwise, you wouldn't get its secret ever.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Status of access key. It must be ` + "`" + `Active` + "`" + ` or ` + "`" + `Inactive` + "`" + `. Default value is ` + "`" + `Active` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "pgp_key",
					Description: `(Optional, Available in 1.47.0+) Either a base-64 encoded PGP public key, or a keybase username in the form ` + "`" + `keybase:some_person_that_exists` + "`" + ` ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The access key ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The access key status.`,
				},
				resource.Attribute{
					Name:        "key_fingerprint",
					Description: `The fingerprint of the PGP key used to encrypt the secret`,
				},
				resource.Attribute{
					Name:        "encrypted_secret",
					Description: `The encrypted secret, base64 encoded. ~> NOTE: The encrypted secret may be decrypted using the command line, for example: ` + "`" + `terraform output encrypted_secret | base64 --decode | keybase pgp decrypt` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The access key ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The access key status.`,
				},
				resource.Attribute{
					Name:        "key_fingerprint",
					Description: `The fingerprint of the PGP key used to encrypt the secret`,
				},
				resource.Attribute{
					Name:        "encrypted_secret",
					Description: `The encrypted secret, base64 encoded. ~> NOTE: The encrypted secret may be decrypted using the command line, for example: ` + "`" + `terraform output encrypted_secret | base64 --decode | keybase pgp decrypt` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ram_account_alias",
			Category:         "RAM",
			ShortDescription: `Provides a RAM cloud account alias.`,
			Description:      ``,
			Keywords: []string{
				"ram",
				"account",
				"alias",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_alias",
					Description: `(Required, ForceNew) Alias of cloud account. This name can have a string of 3 to 32 characters, must contain only alphanumeric characters or hyphens, such as "-", and must not begin with a hyphen. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The account alias ID, it's set to ` + "`" + `account_alias` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "account_alias",
					Description: `The account alias. ## Import RAM account alias can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ram_account_alias.example my-alias ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The account alias ID, it's set to ` + "`" + `account_alias` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "account_alias",
					Description: `The account alias. ## Import RAM account alias can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ram_account_alias.example my-alias ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ram_account_password_policy",
			Category:         "RAM",
			ShortDescription: `Provides a RAM password policy configuration for entire account.`,
			Description:      ``,
			Keywords: []string{
				"ram",
				"account",
				"password",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "minimum_password_length",
					Description: `(Optional) Minimal required length of password for a user. Valid value range: [8-32]. Default to 12.`,
				},
				resource.Attribute{
					Name:        "require_lowercase_characters",
					Description: `(Optional) Specifies if the occurrence of a lowercase character in the password is mandatory. Default to true.`,
				},
				resource.Attribute{
					Name:        "require_uppercase_characters",
					Description: `(Optional) Specifies if the occurrence of an uppercase character in the password is mandatory. Default to true.`,
				},
				resource.Attribute{
					Name:        "require_numbers",
					Description: `(Optional) Specifies if the occurrence of a number in the password is mandatory. Default to true.`,
				},
				resource.Attribute{
					Name:        "require_symbols",
					Description: `(Optional Specifies if the occurrence of a special character in the password is mandatory. Default to true.`,
				},
				resource.Attribute{
					Name:        "hard_expiry",
					Description: `(Optional) Specifies if a password can expire in a hard way. Default to false.`,
				},
				resource.Attribute{
					Name:        "max_password_age",
					Description: `(Optional) The number of days after which password expires. A value of 0 indicates that the password never expires. Valid value range: [0-1095]. Default to 0.`,
				},
				resource.Attribute{
					Name:        "password_reuse_prevention",
					Description: `(Optional) User is not allowed to use the latest number of passwords specified in this parameter. A value of 0 indicates the password history check policy is disabled. Valid value range: [0-24]. Default to 0.`,
				},
				resource.Attribute{
					Name:        "max_login_attempts",
					Description: `(Optional, Type: int) Maximum logon attempts with an incorrect password within an hour. Valid value range: [0-32]. Default to 5. ## Import RAM account password policy can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import alicloud_ram_account_password_policy.example ram-account-password-policy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ram_alias",
			Category:         "RAM",
			ShortDescription: `Provides a RAM cloud account alias.`,
			Description:      ``,
			Keywords: []string{
				"ram",
				"alias",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ram_group",
			Category:         "RAM",
			ShortDescription: `Provides a RAM Group resource.`,
			Description:      ``,
			Keywords: []string{
				"ram",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Name of the RAM group. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comment of the RAM group. This parameter can have a string of 1 to 128 characters.`,
				},
				resource.Attribute{
					Name:        "force",
					Description: `(Optional) This parameter is used for resource destroy. Default value is ` + "`" + `false` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The group ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The group name.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `The group comments. ## Import RAM group can be imported using the id or name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ram_group.example my-group ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The group ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The group name.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `The group comments. ## Import RAM group can be imported using the id or name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ram_group.example my-group ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ram_group_membership",
			Category:         "RAM",
			ShortDescription: `Provides a RAM Group membership resource.`,
			Description:      ``,
			Keywords: []string{
				"ram",
				"group",
				"membership",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group_name",
					Description: `(Required, ForceNew) Name of the RAM group. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.`,
				},
				resource.Attribute{
					Name:        "user_names",
					Description: `(Required) Set of user name which will be added to group. Each name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin with a hyphen. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The membership ID, it's set to ` + "`" + `group_name` + "`" + ``,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `The group name.`,
				},
				resource.Attribute{
					Name:        "user_names",
					Description: `The list of names of users which in the group. ## Import RAM Group membership can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ram_group_membership.example my-group ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The membership ID, it's set to ` + "`" + `group_name` + "`" + ``,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `The group name.`,
				},
				resource.Attribute{
					Name:        "user_names",
					Description: `The list of names of users which in the group. ## Import RAM Group membership can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ram_group_membership.example my-group ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ram_group_policy_attachment",
			Category:         "RAM",
			ShortDescription: `Provides a RAM Group Policy attachment resource.`,
			Description:      ``,
			Keywords: []string{
				"ram",
				"group",
				"policy",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group_name",
					Description: `(Required, ForceNew) Name of the RAM group. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `(Required, ForceNew) Name of the RAM policy. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `(Required, ForceNew) Type of the RAM policy. It must be ` + "`" + `Custom` + "`" + ` or ` + "`" + `System` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The attachment ID. Composed of policy name, policy type and group name with format ` + "`" + `group:<policy_name>:<policy_type>:<group_name>` + "`" + `. ## Import RAM Group Policy attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ram_group_policy_attachment.example group:my-policy:Custom:my-group ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The attachment ID. Composed of policy name, policy type and group name with format ` + "`" + `group:<policy_name>:<policy_type>:<group_name>` + "`" + `. ## Import RAM Group Policy attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ram_group_policy_attachment.example group:my-policy:Custom:my-group ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ram_login_profile",
			Category:         "RAM",
			ShortDescription: `Provides a RAM User Login Profile resource.`,
			Description:      ``,
			Keywords: []string{
				"ram",
				"login",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "user_name",
					Description: `(Required, ForceNew) Name of the RAM user. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin with a hyphen.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required, Sensitive) Password of the RAM user.`,
				},
				resource.Attribute{
					Name:        "mfa_bind_required",
					Description: `(Optional) This parameter indicates whether the MFA needs to be bind when the user first logs in. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "password_reset_required",
					Description: `(Optional) This parameter indicates whether the password needs to be reset when the user first logs in. Default value is ` + "`" + `false` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The login profile ID.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `The user name.`,
				},
				resource.Attribute{
					Name:        "mfa_bind_required",
					Description: `The parameter which indicates whether the MFA needs to be bind when the user first logs in.`,
				},
				resource.Attribute{
					Name:        "password_reset_required",
					Description: `The parameter which indicates whether the password needs to be reset when the user first logs in. ## Import RAM login profile can be imported using the id or user name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ram_login_profile.example my-login ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The login profile ID.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `The user name.`,
				},
				resource.Attribute{
					Name:        "mfa_bind_required",
					Description: `The parameter which indicates whether the MFA needs to be bind when the user first logs in.`,
				},
				resource.Attribute{
					Name:        "password_reset_required",
					Description: `The parameter which indicates whether the password needs to be reset when the user first logs in. ## Import RAM login profile can be imported using the id or user name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ram_login_profile.example my-login ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ram_policy",
			Category:         "RAM",
			ShortDescription: `Provides a RAM Policy resource.`,
			Description:      ``,
			Keywords: []string{
				"ram",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Name of the RAM policy. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.`,
				},
				resource.Attribute{
					Name:        "statement",
					Description: `(Deprecated, Optional, Type: list, Conflicts with ` + "`" + `document` + "`" + `) (It has been deprecated from version 1.49.0, and use field 'document' to replace.) Statements of the RAM policy document. It is required when the ` + "`" + `document` + "`" + ` is not specified.`,
				},
				resource.Attribute{
					Name:        "resource",
					Description: `(Deprecated, Required, Type: list) (It has been deprecated from version 1.49.0, and use field 'document' to replace.) List of specific objects which will be authorized. The format of each item in this list is ` + "`" + `acs:${service}:${region}:${account_id}:${relative_id}` + "`" + `, such as ` + "`" + `acs:ecs:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Deprecated, Required, Type: list) (It has been deprecated from version 1.49.0, and use field 'document' to replace.) List of operations for the ` + "`" + `resource` + "`" + `. The format of each item in this list is ` + "`" + `${service}:${action_name}` + "`" + `, such as ` + "`" + `oss:ListBuckets` + "`" + ` and ` + "`" + `ecs:Describe`,
				},
				resource.Attribute{
					Name:        "effect",
					Description: `(Deprecated, Required) (It has been deprecated from version 1.49.0, and use field 'document' to replace.) This parameter indicates whether or not the ` + "`" + `action` + "`" + ` is allowed. Valid values are ` + "`" + `Allow` + "`" + ` and ` + "`" + `Deny` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Deprecated, Optional, Conflicts with ` + "`" + `document` + "`" + `) (It has been deprecated from version 1.49.0, and use field 'document' to replace.) Version of the RAM policy document. Valid value is ` + "`" + `1` + "`" + `. Default value is ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "document",
					Description: `(Optional, Conflicts with ` + "`" + `statement` + "`" + ` and ` + "`" + `version` + "`" + `) Document of the RAM policy. It is required when the ` + "`" + `statement` + "`" + ` is not specified.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, ForceNew) Description of the RAM policy. This name can have a string of 1 to 1024 characters.`,
				},
				resource.Attribute{
					Name:        "force",
					Description: `(Optional) This parameter is used for resource destroy. Default value is ` + "`" + `false` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The policy ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The policy name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The policy type.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The policy description.`,
				},
				resource.Attribute{
					Name:        "statement",
					Description: `List of statement of the policy document.`,
				},
				resource.Attribute{
					Name:        "document",
					Description: `The policy document.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The policy document version.`,
				},
				resource.Attribute{
					Name:        "attachment_count",
					Description: `The policy attachment count. ## Import RAM policy can be imported using the id or name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ram_policy.example my-policy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The policy ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The policy name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The policy type.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The policy description.`,
				},
				resource.Attribute{
					Name:        "statement",
					Description: `List of statement of the policy document.`,
				},
				resource.Attribute{
					Name:        "document",
					Description: `The policy document.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The policy document version.`,
				},
				resource.Attribute{
					Name:        "attachment_count",
					Description: `The policy attachment count. ## Import RAM policy can be imported using the id or name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ram_policy.example my-policy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ram_role",
			Category:         "RAM",
			ShortDescription: `Provides a RAM Role resource.`,
			Description:      ``,
			Keywords: []string{
				"ram",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Name of the RAM role. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-", "_", and must not begin with a hyphen.`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `(Deprecated, Optional, Type: list, Conflicts with ` + "`" + `document` + "`" + `) (It has been deprecated from version 1.49.0, and use field 'document' to replace.) List of services which can assume the RAM role. The format of each item in this list is ` + "`" + `${service}.aliyuncs.com` + "`" + ` or ` + "`" + `${account_id}@${service}.aliyuncs.com` + "`" + `, such as ` + "`" + `ecs.aliyuncs.com` + "`" + ` and ` + "`" + `1234567890000@ots.aliyuncs.com` + "`" + `. The ` + "`" + `${service}` + "`" + ` can be ` + "`" + `ecs` + "`" + `, ` + "`" + `log` + "`" + `, ` + "`" + `apigateway` + "`" + ` and so on, the ` + "`" + `${account_id}` + "`" + ` refers to someone's Alicloud account id.`,
				},
				resource.Attribute{
					Name:        "ram_users",
					Description: `(Deprecated, Optional, Type: list, Conflicts with ` + "`" + `document` + "`" + `) (It has been deprecated from version 1.49.0, and use field 'document' to replace.) List of ram users who can assume the RAM role. The format of each item in this list is ` + "`" + `acs:ram::${account_id}:root` + "`" + ` or ` + "`" + `acs:ram::${account_id}:user/${user_name}` + "`" + `, such as ` + "`" + `acs:ram::1234567890000:root` + "`" + ` and ` + "`" + `acs:ram::1234567890001:user/Mary` + "`" + `. The ` + "`" + `${user_name}` + "`" + ` is the name of a RAM user which must exists in the Alicloud account indicated by the ` + "`" + `${account_id}` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Deprecated, Optional, Conflicts with ` + "`" + `document` + "`" + `) (It has been deprecated from version 1.49.0, and use field 'document' to replace.) Version of the RAM role policy document. Valid value is ` + "`" + `1` + "`" + `. Default value is ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "document",
					Description: `(Optional, Conflicts with ` + "`" + `services` + "`" + `, ` + "`" + `ram_users` + "`" + ` and ` + "`" + `version` + "`" + `) Authorization strategy of the RAM role. It is required when the ` + "`" + `services` + "`" + ` and ` + "`" + `ram_users` + "`" + ` are not specified.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, Forces new resource) Description of the RAM role. This name can have a string of 1 to 1024 characters.`,
				},
				resource.Attribute{
					Name:        "force",
					Description: `(Optional) This parameter is used for resource destroy. Default value is ` + "`" + `false` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `This ID of this resource. The value is set to ` + "`" + `role_name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `The role ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The role name.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The role arn.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The role description.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The role policy document version.`,
				},
				resource.Attribute{
					Name:        "document",
					Description: `Authorization strategy of the role.`,
				},
				resource.Attribute{
					Name:        "ram_users",
					Description: `List of services which can assume the RAM role.`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `List of services which can assume the RAM role. ## Import RAM role can be imported using the id or name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ram_role.example my-role ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `This ID of this resource. The value is set to ` + "`" + `role_name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `The role ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The role name.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The role arn.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The role description.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The role policy document version.`,
				},
				resource.Attribute{
					Name:        "document",
					Description: `Authorization strategy of the role.`,
				},
				resource.Attribute{
					Name:        "ram_users",
					Description: `List of services which can assume the RAM role.`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `List of services which can assume the RAM role. ## Import RAM role can be imported using the id or name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ram_role.example my-role ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ram_role_attachment",
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
			Type:             "alicloud_ram_role_policy_attachment",
			Category:         "RAM",
			ShortDescription: `Provides a RAM Role Policy attachment resource.`,
			Description:      ``,
			Keywords: []string{
				"ram",
				"role",
				"policy",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "role_name",
					Description: `(Required, ForceNew) Name of the RAM Role. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-", "_", and must not begin with a hyphen.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `(Required, ForceNew) Name of the RAM policy. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `(Required, ForceNew) Type of the RAM policy. It must be ` + "`" + `Custom` + "`" + ` or ` + "`" + `System` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The attachment ID. Composed of policy name, policy type and role name with format ` + "`" + `role:<policy_name>:<policy_type>:<role_name>` + "`" + `. ## Import RAM Role Policy attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ram_role_policy_attachment.example role:my-policy:Custom:my-role ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The attachment ID. Composed of policy name, policy type and role name with format ` + "`" + `role:<policy_name>:<policy_type>:<role_name>` + "`" + `. ## Import RAM Role Policy attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ram_role_policy_attachment.example role:my-policy:Custom:my-role ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ram_user",
			Category:         "RAM",
			ShortDescription: `Provides a RAM User resource.`,
			Description:      ``,
			Keywords: []string{
				"ram",
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the RAM user. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin with a hyphen.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Name of the RAM user which for display. This name can have a string of 1 to 128 characters or Chinese characters, must contain only alphanumeric characters or Chinese characters or hyphens, such as "-",".", and must not end with a hyphen.`,
				},
				resource.Attribute{
					Name:        "mobile",
					Description: `(Optional) Phone number of the RAM user. This number must contain an international area code prefix, just look like this: 86-18600008888.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) Email of the RAM user.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comment of the RAM user. This parameter can have a string of 1 to 128 characters.`,
				},
				resource.Attribute{
					Name:        "force",
					Description: `(Optional) This parameter is used for resource destroy. Default value is ` + "`" + `false` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ram user id. ## Import RAM user can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ram_user.example 123456789xxx ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ram user id. ## Import RAM user can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ram_user.example 123456789xxx ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ram_user_policy_attachment",
			Category:         "RAM",
			ShortDescription: `Provides a RAM User Policy attachment resource.`,
			Description:      ``,
			Keywords: []string{
				"ram",
				"user",
				"policy",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "user_name",
					Description: `(Required, ForceNew) Name of the RAM user. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin with a hyphen.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `(Required, ForceNew) Name of the RAM policy. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `(Required, ForceNew) Type of the RAM policy. It must be ` + "`" + `Custom` + "`" + ` or ` + "`" + `System` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The attachment ID. Composed of policy name, policy type and user name with format ` + "`" + `user:<policy_name>:<policy_type>:<user_name>` + "`" + `. ## Import RAM User Policy attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ram_user_policy_attachment.example user:my-policy:Custom:my-user ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The attachment ID. Composed of policy name, policy type and user name with format ` + "`" + `user:<policy_name>:<policy_type>:<user_name>` + "`" + `. ## Import RAM User Policy attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ram_user_policy_attachment.example user:my-policy:Custom:my-user ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_reserved_instance",
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
					Description: `(Optional, ForceNew) The operating system type of the image used by the instance. Optional values: ` + "`" + `Windows` + "`" + `, ` + "`" + `Linux` + "`" + `. Default is ` + "`" + `Linux` + "`" + `. ### Removing alicloud_reserved_instance from your configuration The alicloud_reserved_instance resource allows you to manage your ReservedInstance, but Terraform cannot destroy it. Removing this resource from your configuration will remove it from your statefile and management, but will not destroy the ReservedInstance. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the ReservedInstance. ## Import reservedInstance can be imported using id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_reserved_instance.default ecsri-uf6df4xm0h3licit`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the ReservedInstance. ## Import reservedInstance can be imported using id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_reserved_instance.default ecsri-uf6df4xm0h3licit`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_resource_manager_account",
			Category:         "Resource Manager",
			ShortDescription: `Provides a Resource Manager Account resource.`,
			Description:      ``,
			Keywords: []string{
				"resource",
				"manager",
				"account",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) Member name. The length is 2 ~ 50 characters or Chinese characters, which can include Chinese characters, English letters, numbers, underscores (_), dots (.) And dashes (-).`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) The ID of the parent folder.`,
				},
				resource.Attribute{
					Name:        "payer_account_id",
					Description: `(Optional, ForceNew) Settlement account ID. If the value is empty, the current account will be used for settlement. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `This ID of Resource Manager Account.`,
				},
				resource.Attribute{
					Name:        "join_method",
					Description: `Ways for members to join the resource directory. Valid values: ` + "`" + `invited` + "`" + `, ` + "`" + `created` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "join_time",
					Description: `The time when the member joined the resource directory.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `The modification time of the invitation.`,
				},
				resource.Attribute{
					Name:        "resource_directory_id",
					Description: `Resource directory ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Member joining status. Valid values: ` + "`" + `CreateSuccess` + "`" + `,` + "`" + `CreateVerifying` + "`" + `,` + "`" + `CreateFailed` + "`" + `,` + "`" + `CreateExpired` + "`" + `,` + "`" + `CreateCancelled` + "`" + `,` + "`" + `PromoteVerifying` + "`" + `,` + "`" + `PromoteFailed` + "`" + `,` + "`" + `PromoteExpired` + "`" + `,` + "`" + `PromoteCancelled` + "`" + `,` + "`" + `PromoteSuccess` + "`" + `,` + "`" + `InviteSuccess` + "`" + `,` + "`" + `Removed` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Member type. The value of ` + "`" + `ResourceAccount` + "`" + ` indicates the resource account. ## Import Resource Manager Account can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_resource_manager_account.example 13148890145`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `This ID of Resource Manager Account.`,
				},
				resource.Attribute{
					Name:        "join_method",
					Description: `Ways for members to join the resource directory. Valid values: ` + "`" + `invited` + "`" + `, ` + "`" + `created` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "join_time",
					Description: `The time when the member joined the resource directory.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `The modification time of the invitation.`,
				},
				resource.Attribute{
					Name:        "resource_directory_id",
					Description: `Resource directory ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Member joining status. Valid values: ` + "`" + `CreateSuccess` + "`" + `,` + "`" + `CreateVerifying` + "`" + `,` + "`" + `CreateFailed` + "`" + `,` + "`" + `CreateExpired` + "`" + `,` + "`" + `CreateCancelled` + "`" + `,` + "`" + `PromoteVerifying` + "`" + `,` + "`" + `PromoteFailed` + "`" + `,` + "`" + `PromoteExpired` + "`" + `,` + "`" + `PromoteCancelled` + "`" + `,` + "`" + `PromoteSuccess` + "`" + `,` + "`" + `InviteSuccess` + "`" + `,` + "`" + `Removed` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Member type. The value of ` + "`" + `ResourceAccount` + "`" + ` indicates the resource account. ## Import Resource Manager Account can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_resource_manager_account.example 13148890145`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_resource_manager_folder",
			Category:         "Resource Manager",
			ShortDescription: `Provides a Alicloud Resource Manager Folder resource.`,
			Description:      ``,
			Keywords: []string{
				"resource",
				"manager",
				"folder",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "folder_name",
					Description: `(Required) The name of the folder. The name must be 1 to 24 characters in length and can contain letters, digits, underscores (_), periods (.), and hyphens (-).`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the folder. ## Import Resource Manager Folder can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_resource_manager_folder.example fd-u8B321`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the folder. ## Import Resource Manager Folder can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_resource_manager_folder.example fd-u8B321`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_resource_manager_handshake",
			Category:         "Resource Manager",
			ShortDescription: `Provides a Resource Manager handshake resource.`,
			Description:      ``,
			Keywords: []string{
				"resource",
				"manager",
				"handshake",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "target_entity",
					Description: `(Required, ForceNew) Invited account ID or login email.`,
				},
				resource.Attribute{
					Name:        "target_type",
					Description: `(Required, ForceNew) Type of account being invited. Valid values: ` + "`" + `Account` + "`" + `, ` + "`" + `Email` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "note",
					Description: `(Optional, ForceNew) Remarks. The maximum length is 1024 characters. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `This ID of Resource Manager handshake.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `The expiration time of the invitation.`,
				},
				resource.Attribute{
					Name:        "master_account_id",
					Description: `Resource account master account ID.`,
				},
				resource.Attribute{
					Name:        "master_account_name",
					Description: `The name of the main account of the resource directory.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `The modification time of the invitation.`,
				},
				resource.Attribute{
					Name:        "resource_directory_id",
					Description: `Resource directory ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Invitation status. Valid values: ` + "`" + `Pending` + "`" + ` waiting for confirmation, ` + "`" + `Accepted` + "`" + `, ` + "`" + `Cancelled` + "`" + `, ` + "`" + `Declined` + "`" + `, ` + "`" + `Expired` + "`" + `. ## Import Resource Manager handshake can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_resource_manager_handshake.example h-QmdexeFm1kE`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `This ID of Resource Manager handshake.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `The expiration time of the invitation.`,
				},
				resource.Attribute{
					Name:        "master_account_id",
					Description: `Resource account master account ID.`,
				},
				resource.Attribute{
					Name:        "master_account_name",
					Description: `The name of the main account of the resource directory.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `The modification time of the invitation.`,
				},
				resource.Attribute{
					Name:        "resource_directory_id",
					Description: `Resource directory ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Invitation status. Valid values: ` + "`" + `Pending` + "`" + ` waiting for confirmation, ` + "`" + `Accepted` + "`" + `, ` + "`" + `Cancelled` + "`" + `, ` + "`" + `Declined` + "`" + `, ` + "`" + `Expired` + "`" + `. ## Import Resource Manager handshake can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_resource_manager_handshake.example h-QmdexeFm1kE`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_resource_manager_policy",
			Category:         "Resource Manager",
			ShortDescription: `Provides a Alicloud Resource Manager Policy resource.`,
			Description:      ``,
			Keywords: []string{
				"resource",
				"manager",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_name",
					Description: `(Required, ForceNew) The name of the policy. name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).`,
				},
				resource.Attribute{
					Name:        "policy_document",
					Description: `(Required, ForceNew) The content of the policy. The content must be 1 to 2,048 characters in length.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, ForceNew) The description of the policy. The description must be 1 to 1,024 characters in length.`,
				},
				resource.Attribute{
					Name:        "default_version",
					Description: `(Optional, Computed, Deprecated from version 1.90.0) The version of the policy. Default to v1. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID of policy. The value is same as ` + "`" + `policy_name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "create_date",
					Description: `The time when the policy was created.`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `The type of the policy. Valid values: ` + "`" + `Custom` + "`" + `, ` + "`" + `System` + "`" + `. ## Import Resource Manager Policy can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_resource_policy.example abc12345 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID of policy. The value is same as ` + "`" + `policy_name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "create_date",
					Description: `The time when the policy was created.`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `The type of the policy. Valid values: ` + "`" + `Custom` + "`" + `, ` + "`" + `System` + "`" + `. ## Import Resource Manager Policy can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_resource_policy.example abc12345 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_resource_manager_policy_version",
			Category:         "Resource Manager",
			ShortDescription: `Provides a Alicloud Resource Manager Policy Version resource.`,
			Description:      ``,
			Keywords: []string{
				"resource",
				"manager",
				"policy",
				"version",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_name",
					Description: `(Required, ForceNew) The name of the policy. Name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).`,
				},
				resource.Attribute{
					Name:        "policy_document",
					Description: `(Required, ForceNew) The content of the policy. The content must be 1 to 2,048 characters in length.`,
				},
				resource.Attribute{
					Name:        "is_default_version",
					Description: `(Optional, Deprecated from version 1.90.0) Specifies whether to set the policy version as the default version. Default to ` + "`" + `false` + "`" + `. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID of policy version. The value is "` + "`" + `<policy_name>` + "`" + `:` + "`" + `<version_id>` + "`" + `".`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `The ID of the policy version.`,
				},
				resource.Attribute{
					Name:        "create_date",
					Description: `The time when the policy version was created. ## Import Resource Manager Policy Version can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_resource_manager_policy_version.example tftest:v2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID of policy version. The value is "` + "`" + `<policy_name>` + "`" + `:` + "`" + `<version_id>` + "`" + `".`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `The ID of the policy version.`,
				},
				resource.Attribute{
					Name:        "create_date",
					Description: `The time when the policy version was created. ## Import Resource Manager Policy Version can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_resource_manager_policy_version.example tftest:v2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_resource_manager_resource_directory",
			Category:         "Resource Manager",
			ShortDescription: `Provides a Alicloud Resource Manager Resource Directory resource.`,
			Description:      ``,
			Keywords: []string{
				"resource",
				"manager",
				"directory",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the resource directory.`,
				},
				resource.Attribute{
					Name:        "master_account_id",
					Description: `The ID of the master account.`,
				},
				resource.Attribute{
					Name:        "master_account_name",
					Description: `The name of the master account.`,
				},
				resource.Attribute{
					Name:        "root_folder_id",
					Description: `The ID of the root folder. ## Import Resource Manager Resource Directory can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_resource_manager_resource_directory.example rd-s3`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the resource directory.`,
				},
				resource.Attribute{
					Name:        "master_account_id",
					Description: `The ID of the master account.`,
				},
				resource.Attribute{
					Name:        "master_account_name",
					Description: `The name of the master account.`,
				},
				resource.Attribute{
					Name:        "root_folder_id",
					Description: `The ID of the root folder. ## Import Resource Manager Resource Directory can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_resource_manager_resource_directory.example rd-s3`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_resource_manager_resource_group",
			Category:         "Resource Manager",
			ShortDescription: `Provides a Alicloud Resource Manager Resource Group resource.`,
			Description:      ``,
			Keywords: []string{
				"resource",
				"manager",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) The unique identifier of the resource group.The identifier must be 3 to 12 characters in length and can contain letters, digits, periods (.), hyphens (-), and underscores (_). The identifier must start with a letter.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) The display name of the resource group. The name must be 1 to 30 characters in length and can contain letters, digits, periods (.), at signs (@), and hyphens (-). ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the resource group.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the resource group.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The ID of the Alibaba Cloud account to which the resource group belongs.`,
				},
				resource.Attribute{
					Name:        "create_date",
					Description: `The time when the resource group was created.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the resource group.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the resource group.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The ID of the Alibaba Cloud account to which the resource group belongs.`,
				},
				resource.Attribute{
					Name:        "create_date",
					Description: `The time when the resource group was created.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_resource_manager_role",
			Category:         "Resource Manager",
			ShortDescription: `Provides a Resource Manager role resource.`,
			Description:      ``,
			Keywords: []string{
				"resource",
				"manager",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "assume_role_policy_document",
					Description: `(Required) The content of the permissions strategy that plays a role.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, ForceNew) The description of the Resource Manager role.`,
				},
				resource.Attribute{
					Name:        "max_session_duration",
					Description: `(Optional, ForceNew) Role maximum session time. Valid values: [3600-43200]. Default to ` + "`" + `3600` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "role_name",
					Description: `(Required, ForceNew) Role Name. The length is 1 ~ 64 characters, which can include English letters, numbers, dots "." and dashes "-". ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `This ID of Resource Manager role. The value is set to ` + "`" + `role_name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The resource descriptor of the role.`,
				},
				resource.Attribute{
					Name:        "create_date",
					Description: `Role creation time.`,
				},
				resource.Attribute{
					Name:        "update_date",
					Description: `Role update time. ## Import Resource Manager can be imported using the id or role_name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_resource_manager_role.example testrd ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `This ID of Resource Manager role. The value is set to ` + "`" + `role_name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The resource descriptor of the role.`,
				},
				resource.Attribute{
					Name:        "create_date",
					Description: `Role creation time.`,
				},
				resource.Attribute{
					Name:        "update_date",
					Description: `Role update time. ## Import Resource Manager can be imported using the id or role_name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_resource_manager_role.example testrd ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_route_entry",
			Category:         "VPC",
			ShortDescription: `Provides a Alicloud Route Entry resource.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"route",
				"entry",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "router_id",
					Description: `(Deprecated) This argument has beeb deprecated. Please use other arguments to launch a custom route entry.`,
				},
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
					Description: `(ForceNew) The route entry's next hop. ECS instance ID or VPC router interface ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, ForceNew, Available in 1.55.1+) The name of the route entry. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The route entry id,it formats of ` + "`" + `<route_table_id:router_id:destination_cidrblock:nexthop_type:nexthop_id>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `The ID of the virtual router attached to Vpc.`,
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
					Description: `The route entry's next hop. ## Import Router entry can be imported using the id, e.g (formatted as<route_table_id:router_id:destination_cidrblock:nexthop_type:nexthop_id>). ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_route_entry.example vtb-123456:vrt-123456:0.0.0.0/0:NatGateway:ngw-123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The route entry id,it formats of ` + "`" + `<route_table_id:router_id:destination_cidrblock:nexthop_type:nexthop_id>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `The ID of the virtual router attached to Vpc.`,
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
					Description: `The route entry's next hop. ## Import Router entry can be imported using the id, e.g (formatted as<route_table_id:router_id:destination_cidrblock:nexthop_type:nexthop_id>). ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_route_entry.example vtb-123456:vrt-123456:0.0.0.0/0:NatGateway:ngw-123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_route_table",
			Category:         "VPC",
			ShortDescription: `Provides a Alicloud Route Table resource.`,
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
					Description: `(Optional, Available in v1.55.3+) A mapping of tags to assign to the resource. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the route table instance id. ## Import The route table can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_route_table.foo vtb-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the route table instance id. ## Import The route table can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_route_table.foo vtb-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_route_table_attachment",
			Category:         "VPC",
			ShortDescription: `Provides an Alicloud Route Table Attachment resource.`,
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
					Description: `The ID of the route table attachment id and formates as ` + "`" + `<route_table_id>:<vswitch_id>` + "`" + `. ## Import The route table attachemnt can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_route_table_attachment.foo vtb-abc123456:vsw-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the route table attachment id and formates as ` + "`" + `<route_table_id>:<vswitch_id>` + "`" + `. ## Import The route table attachemnt can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_route_table_attachment.foo vtb-abc123456:vsw-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_router_interface",
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
					Name:        "opposite_router_type",
					Description: `(Deprecated) It has been deprecated from version 1.11.0. resource alicloud_router_interface_connection's 'opposite_router_type' instead.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `(Required, ForceNew) The Router ID.`,
				},
				resource.Attribute{
					Name:        "opposite_router_id",
					Description: `(Deprecated) It has been deprecated from version 1.11.0. Use resource alicloud_router_interface_connection's 'opposite_router_id' instead.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required, ForceNew) The role the router interface plays. Optional value: ` + "`" + `InitiatingSide` + "`" + `, ` + "`" + `AcceptingSide` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "specification",
					Description: `(Optional) Specification of router interfaces. It is valid when ` + "`" + `role` + "`" + ` is ` + "`" + `InitiatingSide` + "`" + `. Accepting side's role is default to set as 'Negative'. For more about the specification, refer to [Router interface specification](https://www.alibabacloud.com/help/doc-detail/36037.htm).`,
				},
				resource.Attribute{
					Name:        "access_point_id",
					Description: `(Deprecated) It has been deprecated from version 1.11.0.`,
				},
				resource.Attribute{
					Name:        "opposite_access_point_id",
					Description: `(Deprecated) It has been deprecated from version 1.11.0.`,
				},
				resource.Attribute{
					Name:        "opposite_interface_id",
					Description: `(Deprecated) It has been deprecated from version 1.11.0. Use resource alicloud_router_interface_connection's 'opposite_router_id' instead.`,
				},
				resource.Attribute{
					Name:        "opposite_interface_owner_id",
					Description: `(Deprecated) It has been deprecated from version 1.11.0. Use resource alicloud_router_interface_connection's 'opposite_interface_id' instead.`,
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
					Description: `(Optional) Used as the Packet Target IP of health check for disaster recovery or ECMP. It is only valid when ` + "`" + `router_type` + "`" + ` is ` + "`" + `VBR` + "`" + `. The IP must be an unused IP in the local VPC. It and ` + "`" + `health_check_source_ip` + "`" + ` must be specified at the same time.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `(Optional, ForceNew) The billing method of the router interface. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid". Router Interface doesn't support "PrePaid" when region and opposite_region are the same.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional, ForceNew) The duration that you will buy the resource, in month. It is valid when ` + "`" + `instance_charge_type` + "`" + ` is ` + "`" + `PrePaid` + "`" + `. Default to 1. Valid values: [1-9, 12, 24, 36]. At present, the provider does not support modify "period" and you can do that via web console. ## Attributes Reference The following attributes are exported:`,
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
					Description: `(Deprecated) It has been deprecated from version 1.11.0.`,
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
					Description: `Target IP of Packet of Line HealthCheck. ## Import The router interface can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_router_interface.interface ri-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Deprecated) It has been deprecated from version 1.11.0.`,
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
					Description: `Target IP of Packet of Line HealthCheck. ## Import The router interface can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_router_interface.interface ri-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_router_interface_connection",
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
					Description: `(Optional, ForceNew) Another side router interface account ID. Log on to the Alibaba Cloud console, select User Info > Account Management to check the account ID. Default to [Provider account_id](https://www.terraform.io/docs/providers/alicloud/index.html#account_id).`,
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
					Description: `Router interface ID. The value is equal to "interface_id". ## Import The router interface connection can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_router_interface_connection.foo ri-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Router interface ID. The value is equal to "interface_id". ## Import The router interface connection can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_router_interface_connection.foo ri-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_sag_acl",
			Category:         "Smart Access Gateway",
			ShortDescription: `Provides a Sag Acl resource.`,
			Description:      ``,
			Keywords: []string{
				"smart",
				"access",
				"gateway",
				"sag",
				"acl",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the ACL instance. The name can contain 2 to 128 characters including a-z, A-Z, 0-9, periods, underlines, and hyphens. The name must start with an English letter, but cannot start with http:// or https://. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the ACL. For example "acl-xxx". ## Import The Sag Acl can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_sag_acl.example acl-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the ACL. For example "acl-xxx". ## Import The Sag Acl can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_sag_acl.example acl-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_sag_acl_rule",
			Category:         "Smart Access Gateway",
			ShortDescription: `Provides a Sag Acl Rule resource.`,
			Description:      ``,
			Keywords: []string{
				"smart",
				"access",
				"gateway",
				"sag",
				"acl",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "acl_id",
					Description: `(Required) The ID of the ACL.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the ACL rule. It must be 1 to 512 characters in length.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Required) The policy used by the ACL rule. Valid values: accept|drop.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `(Required) The protocol used by the ACL rule. The value is not case sensitive.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Required) The direction of the ACL rule. Valid values: in|out.`,
				},
				resource.Attribute{
					Name:        "source_cidr",
					Description: `(Required) The source address. It is an IPv4 address range in the CIDR format. Default value: 0.0.0.0/0.`,
				},
				resource.Attribute{
					Name:        "source_port_range",
					Description: `(Required) The range of the source port. Valid value: 80/80.`,
				},
				resource.Attribute{
					Name:        "dest_cidr",
					Description: `(Required) The destination address. It is an IPv4 address range in CIDR format. Default value: 0.0.0.0/0.`,
				},
				resource.Attribute{
					Name:        "dest_port_range",
					Description: `(Required) The range of the destination port. Valid value: 80/80.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) The priority of the ACL rule. Value range: 1 to 100. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the ACL rule. For example "acr-xxx". ## Import The Sag Acl Rule can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_sag_acl_rule.example acr-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the ACL rule. For example "acr-xxx". ## Import The Sag Acl Rule can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_sag_acl_rule.example acr-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_sag_client_user",
			Category:         "Smart Access Gateway",
			ShortDescription: `Provides a Sag ClientUser resource.`,
			Description:      ``,
			Keywords: []string{
				"smart",
				"access",
				"gateway",
				"sag",
				"client",
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "sag_id",
					Description: `(Required,ForceNew) The ID of the SAG instance created for the SAG APP.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Required) The SAG APP bandwidth that the user can use. Unit: Kbit/s. Maximum value: 2000 Kbit/s.`,
				},
				resource.Attribute{
					Name:        "user_mail",
					Description: `(Required,ForceNew) The email address of the user. The administrator uses this address to send the account information for logging on to the APP to the user.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `(Optional,ForceNew) The user name. User names in the same SAG APP must be unique.Both the user name and the password must be specified. If you specify the user name, the password must be specified, too.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional,ForceNew) The password used to log on to the SAG APP.Both the user name and the password must be specified. If you specify the user name, the password must be specified, too.`,
				},
				resource.Attribute{
					Name:        "client_ip",
					Description: `(Optional,ForceNew) The IP address of the SAG APP. If you specify this parameter, the current account always uses the specified IP address.Note The IP address must be in the private CIDR block of the SAG client.If you do not specify this parameter, the system automatically allocates an IP address from the private CIDR block of the SAG client. In this case, each re-connection uses a different IP address. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Sag Id and formates as ` + "`" + `<sag_id>:<user_name>` + "`" + `. ## Import The Sag ClientUser can be imported using the name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_sag_client_user.example sag-abc123456:tf-username-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Sag Id and formates as ` + "`" + `<sag_id>:<user_name>` + "`" + `. ## Import The Sag ClientUser can be imported using the name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_sag_client_user.example sag-abc123456:tf-username-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_sag_dnat_entry",
			Category:         "Smart Access Gateway",
			ShortDescription: `Provides a Sag DnatEntry resource.`,
			Description:      ``,
			Keywords: []string{
				"smart",
				"access",
				"gateway",
				"sag",
				"dnat",
				"entry",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "sag_id",
					Description: `(Required) The ID of the SAG instance.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The DNAT type. Valid values: Intranet: DNAT of private IP addresses. Internet: DNAT of public IP addresses`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `(Required) The protocol type. Valid values: TCP: Forwards packets of the TCP protocol. UDP: Forwards packets of the UDP protocol. Any: Forwards packets of all protocols.`,
				},
				resource.Attribute{
					Name:        "external_ip",
					Description: `(Optional) The external public IP address.when "type" is "Internet",automatically identify the external ip.`,
				},
				resource.Attribute{
					Name:        "external_port",
					Description: `(Required) The public port.Value range: 1 to 65535 or "any".`,
				},
				resource.Attribute{
					Name:        "internal_ip",
					Description: `(Required) The destination private IP address.`,
				},
				resource.Attribute{
					Name:        "internal_port",
					Description: `(Required) The destination private port.Value range: 1 to 65535 or "any". ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the DNAT entry Id and formates as ` + "`" + `<sag_id>:<dnat_id>` + "`" + `. ## Import The Sag DnatEntry can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_sag_dnat_entry.example sag-abc123456:dnat-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the DNAT entry Id and formates as ` + "`" + `<sag_id>:<dnat_id>` + "`" + `. ## Import The Sag DnatEntry can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_sag_dnat_entry.example sag-abc123456:dnat-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_sag_qos",
			Category:         "Smart Access Gateway",
			ShortDescription: `Provides a Sag Qos resource.`,
			Description:      ``,
			Keywords: []string{
				"smart",
				"access",
				"gateway",
				"sag",
				"qos",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the QoS policy to be created. The name can contain 2 to 128 characters including a-z, A-Z, 0-9, periods, underlines, and hyphens. The name must start with an English letter, but cannot start with http:// or https://. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Qos. For example "qos-xxx". ## Import The Sag Qos can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_sag_qos.example qos-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Qos. For example "qos-xxx". ## Import The Sag Qos can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_sag_qos.example qos-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_sag_qos_car",
			Category:         "Smart Access Gateway",
			ShortDescription: `Provides a Sag Qos Car resource.`,
			Description:      ``,
			Keywords: []string{
				"smart",
				"access",
				"gateway",
				"sag",
				"qos",
				"car",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "qos_id",
					Description: `(Required) The instance ID of the QoS.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the QoS speed limiting rule..`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the QoS speed limiting rule.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Required) The priority of the specified stream.`,
				},
				resource.Attribute{
					Name:        "limit_type",
					Description: `(Required) The speed limiting method. Valid values: Absolute, Percent.`,
				},
				resource.Attribute{
					Name:        "min_bandwidth_abs",
					Description: `(Optional) The minimum bandwidth allowed for the stream specified in the quintuple rule. This parameter is required when the value of the LimitType parameter is Absolute.`,
				},
				resource.Attribute{
					Name:        "max_bandwidth_abs",
					Description: `(Optional) The maximum bandwidth allowed for the stream specified in the quintuple rule. This parameter is required when the value of the LimitType is Absolute.`,
				},
				resource.Attribute{
					Name:        "min_bandwidth_percent",
					Description: `(Optional) The minimum bandwidth percentage allowed for the stream specified in the quintuple rule. It is based on the maximum upstream bandwidth you set for the associated SAG instance.This parameter is required when the value of the LimitType parameter is Percent.`,
				},
				resource.Attribute{
					Name:        "max_bandwidth_percent",
					Description: `(Optional) The maximum bandwidth percentage allowed for the stream specified in the quintuple rule. It is based on the maximum upstream bandwidth you set for the associated Smart Access Gateway (SAG) instance.This parameter is required when the value of the LimitType parameter is Percent.`,
				},
				resource.Attribute{
					Name:        "percent_source_type",
					Description: `(Optional) The bandwidth type when the speed is limited based on percentage. Valid values: CcnBandwidth, InternetUpBandwidth.The default value is InternetUpBandwidth. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Qos Car id and formates as ` + "`" + `<qos_id>:<qos_car_id>` + "`" + `. ## Import The Sag Qos Car can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_sag_qos_car.example qos-abc123456:qoscar-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Qos Car id and formates as ` + "`" + `<qos_id>:<qos_car_id>` + "`" + `. ## Import The Sag Qos Car can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_sag_qos_car.example qos-abc123456:qoscar-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_sag_qos_policy",
			Category:         "Smart Access Gateway",
			ShortDescription: `Provides a Sag Qos Policy resource.`,
			Description:      ``,
			Keywords: []string{
				"smart",
				"access",
				"gateway",
				"sag",
				"qos",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "qos_id",
					Description: `(Required) The instance ID of the QoS policy to which the quintuple rule is created.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the QoS policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the QoS policy.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Required) The priority of the quintuple rule. A smaller value indicates a higher priority. If the priorities of two quintuple rules are the same, the rule created earlier is applied first.Value range: 1 to 7.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `(Required) The transport layer protocol.`,
				},
				resource.Attribute{
					Name:        "source_cidr",
					Description: `(Required) The source CIDR block.`,
				},
				resource.Attribute{
					Name:        "source_port_range",
					Description: `(Required) The source port range of the transport layer.`,
				},
				resource.Attribute{
					Name:        "dest_cidr",
					Description: `(Required) The destination CIDR block.`,
				},
				resource.Attribute{
					Name:        "dest_port_range",
					Description: `(Required) The destination port range.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `(Optional) The time when the quintuple rule takes effect.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `(Optional) The expiration time of the quintuple rule. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Qos Policy id and formates as ` + "`" + `<qos_id>:<qos_policy_id>` + "`" + `. ## Import The Sag Qos Policy can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_sag_qos_policy.example qos-abc123456:qospy-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Qos Policy id and formates as ` + "`" + `<qos_id>:<qos_policy_id>` + "`" + `. ## Import The Sag Qos Policy can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_sag_qos_policy.example qos-abc123456:qospy-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_sag_snat_entry",
			Category:         "Smart Access Gateway",
			ShortDescription: `Provides a Sag SnatEntry resource.`,
			Description:      ``,
			Keywords: []string{
				"smart",
				"access",
				"gateway",
				"sag",
				"snat",
				"entry",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "sag_id",
					Description: `(Required) The ID of the SAG instance.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Required) The destination CIDR block.`,
				},
				resource.Attribute{
					Name:        "snat_ip",
					Description: `(Required) The public IP address. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the SNAT entry Id and formates as ` + "`" + `<sag_id>:<snat_id>` + "`" + `. ## Import The Sag SnatEntry can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_sag_snat_entry.example sag-abc123456:snat-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the SNAT entry Id and formates as ` + "`" + `<sag_id>:<snat_id>` + "`" + `. ## Import The Sag SnatEntry can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_sag_snat_entry.example sag-abc123456:snat-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_security_group",
			Category:         "ECS",
			ShortDescription: `Provides a Alicloud Security Group resource.`,
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
					Name:        "resource_group_id",
					Description: `(Optional, ForceNew, Available in 1.58.0+) The Id of resource group which the security_group belongs.`,
				},
				resource.Attribute{
					Name:        "security_group_type",
					Description: `(Optional, ForceNew, Available in 1.58.0+) The type of the security group. Valid values: ` + "`" + `normal` + "`" + `: basic security group. ` + "`" + `enterprise` + "`" + `: advanced security group For more information.`,
				},
				resource.Attribute{
					Name:        "inner_access",
					Description: `(Deprecated) Field 'inner_access' has been deprecated from provider version 1.55.3. Use 'inner_access_policy' replaces it.`,
				},
				resource.Attribute{
					Name:        "inner_access_policy",
					Description: `(Optional, Available in 1.55.3+) Whether to allow both machines to access each other on all ports in the same security group. Valid values: ["Accept", "Drop"]`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. Combining security group rules, the policy can define multiple application scenario. Default to true. It is valid from verison ` + "`" + `1.7.2` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the security group ## Import Security Group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_security_group.example sg-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the security group ## Import Security Group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_security_group.example sg-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_security_group_rule",
			Category:         "ECS",
			ShortDescription: `Provides a Alicloud Security Group Rule resource.`,
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
					Description: `(Optional, ForceNew) The Alibaba Cloud user account Id of the target security group when security groups are authorized across accounts. This parameter is invalid if ` + "`" + `cidr_ip` + "`" + ` has already been set.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the security group rule. The description can be up to 1 to 512 characters in length. Defaults to null. ->`,
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
			Type:             "alicloud_slb",
			Category:         "Server Load Balancer (SLB)",
			ShortDescription: `Provides an Application Load Banlancer resource.`,
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
					Description: `(Optional) The name of the SLB. This name must be unique within your AliCloud account, can have a maximum of 80 characters, must contain only alphanumeric characters or hyphens, such as "-","/",".","_", and must not begin or end with a hyphen. If not specified, Terraform will autogenerate a name beginning with ` + "`" + `tf-lb` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "internet",
					Description: `(Deprecated) Field 'internet' has been deprecated from provider version 1.55.3. Use 'address_type' replaces it.`,
				},
				resource.Attribute{
					Name:        "address_type",
					Description: `(Optional, ForceNew, Available in 1.55.3+) The network type of the SLB instance. Valid values: ["internet", "intranet"]. If load balancer launched in VPC, this value must be "intranet". - internet: After an Internet SLB instance is created, the system allocates a public IP address so that the instance can forward requests from the Internet. - intranet: After an intranet SLB instance is created, the system allocates an intranet IP address so that the instance can only forward intranet requests.`,
				},
				resource.Attribute{
					Name:        "internet_charge_type",
					Description: `(Optional, ForceNew) Valid values are ` + "`" + `PayByBandwidth` + "`" + `, ` + "`" + `PayByTraffic` + "`" + `. If this value is "PayByBandwidth", then argument "internet" must be "true". Default is "PayByTraffic". If load balancer launched in VPC, this value must be "PayByTraffic". Before version 1.10.1, the valid values are "paybybandwidth" and "paybytraffic".`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Optional) Valid value is between 1 and 1000, If argument "internet_charge_type" is "paybytraffic", then this value will be ignore.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Required for a VPC SLB, Forces New Resource) The VSwitch ID to launch in. If ` + "`" + `address_type` + "`" + ` is internet, it will be ignore.`,
				},
				resource.Attribute{
					Name:        "specification",
					Description: `(Optional) The specification of the Server Load Balancer instance. Default to empty string indicating it is "Shared-Performance" instance. Launching "[Performance-guaranteed](https://www.alibabacloud.com/help/doc-detail/27657.htm)" instance, it is must be specified and it valid values are: "slb.s1.small", "slb.s2.small", "slb.s2.medium", "slb.s3.small", "slb.s3.medium", "slb.s3.large" and "slb.s4.large".`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. The ` + "`" + `tags` + "`" + ` can have a maximum of 10 tag for every load balancer instance.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `(Optional, Available in v1.34.0+) The billing method of the load balancer. Valid values are "PrePaid" and "PostPaid". Default to "PostPaid".`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional, Available in v1.34.0+) The duration that you will buy the resource, in month. It is valid when ` + "`" + `instance_charge_type` + "`" + ` is ` + "`" + `PrePaid` + "`" + `. Default to 1. Valid values: [1-9, 12, 24, 36].`,
				},
				resource.Attribute{
					Name:        "master_zone_id",
					Description: `(Optional, ForceNew, Available in v1.36.0+) The primary zone ID of the SLB instance. If not specified, the system will be randomly assigned. You can query the primary and standby zones in a region by calling the DescribeZone API.`,
				},
				resource.Attribute{
					Name:        "slave_zone_id",
					Description: `(Optional, ForceNew, Available in v1.36.0+) The standby zone ID of the SLB instance. If not specified, the system will be randomly assigned. You can query the primary and standby zones in a region by calling the DescribeZone API.`,
				},
				resource.Attribute{
					Name:        "delete_protection",
					Description: `(Optional, Available in v1.51.0+) Whether enable the deletion protection or not. on: Enable deletion protection. off: Disable deletion protection. Default to off. Only postpaid instance support this function.`,
				},
				resource.Attribute{
					Name:        "address_ip_version",
					Description: `(Optional, Available in v1.55.2+) The IP version of the SLB instance to be created, which can be set to ipv4 or ipv6 . Default to "ipv4". Now, only internet instance support ipv6 address.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Optional, Available in v1.55.2+) Specify the IP address of the private network for the SLB instance, which must be in the destination CIDR block of the correspond ing switch.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional, ForceNew, Available in v1.55.3+) The Id of resource group which the SLB belongs. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the load balancer.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The IP address of the load balancer. ## Import Load balancer can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_slb.example lb-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the load balancer.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The IP address of the load balancer. ## Import Load balancer can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_slb.example lb-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_slb_acl",
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
					Description: `The Id of the access control list. ## Import Server Load balancer access control list can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_slb_acl.example acl-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Id of the access control list. ## Import Server Load balancer access control list can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_slb_acl.example acl-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_slb_attachment",
			Category:         "Server Load Balancer (SLB)",
			ShortDescription: `Provides an Application Load Banlancer Attachment resource.`,
			Description:      ``,
			Keywords: []string{
				"server",
				"load",
				"balancer",
				"slb",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "load_balancer_id",
					Description: `(Required) ID of the load balancer.`,
				},
				resource.Attribute{
					Name:        "instance_ids",
					Description: `(Required) A list of instance ids to added backend server in the SLB.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) Weight of the instances. Valid value range: [0-100]. Default to 100.`,
				},
				resource.Attribute{
					Name:        "server_type",
					Description: `(Optional, Available in 1.60.0+) Type of the instances. Valid value ecs, eni. Default to ecs.`,
				},
				resource.Attribute{
					Name:        "delete_protection_validation",
					Description: `(Optional, Available in 1.63.0+) Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "load_balancer_id",
					Description: `ID of the load balancer.`,
				},
				resource.Attribute{
					Name:        "instance_ids",
					Description: `A list of instance ids that have been added in the SLB.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Weight of the instances.`,
				},
				resource.Attribute{
					Name:        "backend_servers",
					Description: `The backend servers of the load balancer.`,
				},
				resource.Attribute{
					Name:        "server_type",
					Description: `Type of the instances. ## Import Load balancer attachment can be imported using the id or load balancer id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_slb_attachment.example lb-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "load_balancer_id",
					Description: `ID of the load balancer.`,
				},
				resource.Attribute{
					Name:        "instance_ids",
					Description: `A list of instance ids that have been added in the SLB.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Weight of the instances.`,
				},
				resource.Attribute{
					Name:        "backend_servers",
					Description: `The backend servers of the load balancer.`,
				},
				resource.Attribute{
					Name:        "server_type",
					Description: `Type of the instances. ## Import Load balancer attachment can be imported using the id or load balancer id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_slb_attachment.example lb-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_slb_backend_server",
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
					Description: `(Required) A list of instances to added backend server in the SLB. It contains three sub-fields as ` + "`" + `Block server` + "`" + ` follows.`,
				},
				resource.Attribute{
					Name:        "delete_protection_validation",
					Description: `(Optional, Available in 1.63.0+) Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false. ## Block servers The servers mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "server_id",
					Description: `(Required) A list backend server ID (ECS instance ID).`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) Weight of the backend server. Valid value range: [0-100].`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type of the backend server. Valid value ecs, eni. Default to eni. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the resource and the value same as load balancer id. ## Import Load balancer backend server can be imported using the load balancer id. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_slb_backend_server.example lb-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the resource and the value same as load balancer id. ## Import Load balancer backend server can be imported using the load balancer id. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_slb_backend_server.example lb-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_slb_ca_certificate",
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
					Description: `(Required, ForceNew) the content of the CA certificate.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional, ForceNew, Available in 1.58.0+) The Id of resource group which the slb_ca certificate belongs.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Available in v1.66.0+) A mapping of tags to assign to the resource. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Id of CA Certificate . ## Import Server Load balancer CA Certificate can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_slb_ca_certificate.example abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Id of CA Certificate . ## Import Server Load balancer CA Certificate can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_slb_ca_certificate.example abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_slb_domain_extension",
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
					Description: `The ID of the domain extension. ## Import Load balancer domain_extension can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_slb_domain_extension.example de-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the domain extension. ## Import Load balancer domain_extension can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_slb_domain_extension.example de-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_slb_listener",
			Category:         "Server Load Balancer (SLB)",
			ShortDescription: `Provides an Application Load Banlancer resource.`,
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
					Description: `(Optional, ForceNew) Port used by the Server Load Balancer instance backend. Valid value range: [1-65535].`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required, ForceNew) The protocol to listen on. Valid values are [` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + `, ` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `].`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Optional) Bandwidth peak of Listener. For the public network instance charged per traffic consumed, the Bandwidth on Listener can be set to -1, indicating the bandwidth peak is unlimited. Valid values are [-1, 1-1000] in Mbps.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, Available in 1.69.0+) The description of slb listener. This description can have a string of 1 to 80 characters. Default value: null.`,
				},
				resource.Attribute{
					Name:        "scheduler",
					Description: `(Optional) Scheduling algorithm, Valid values are ` + "`" + `wrr` + "`" + `, ` + "`" + `rr` + "`" + ` and ` + "`" + `wlc` + "`" + `. Default to "wrr".`,
				},
				resource.Attribute{
					Name:        "sticky_session",
					Description: `(Optional) Whether to enable session persistence, Valid values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. Default to ` + "`" + `off` + "`" + `.`,
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
					Description: `(Optional) Whether to enable health check. Valid values are` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. TCP and UDP listener's HealthCheck is always on, so it will be ignore when launching TCP or UDP listener.`,
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
					Description: `(Optional, Available in 1.70.0+) HealthCheckMethod used for health check.` + "`" + `http` + "`" + ` and ` + "`" + `https` + "`" + ` support regions ap-northeast-1, ap-southeast-1, ap-southeast-2, ap-southeast-3, us-east-1, us-west-1, eu-central-1, ap-south-1, me-east-1, cn-huhehaote, cn-zhangjiakou, ap-southeast-5, cn-shenzhen, cn-hongkong, cn-qingdao, cn-chengdu, eu-west-1, cn-hangzhou", cn-beijing, cn-shanghai.This function does not support the TCP protocol .`,
				},
				resource.Attribute{
					Name:        "ssl_certificate_id",
					Description: `(Deprecated) It has been deprecated from 1.59.0 and using ` + "`" + `server_certificate_id` + "`" + ` instead.`,
				},
				resource.Attribute{
					Name:        "server_certificate_id",
					Description: `(Optional, Available in 1.59.0+) SLB Server certificate ID. It is required when ` + "`" + `protocol` + "`" + ` is ` + "`" + `https` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "gzip",
					Description: `(Optional) Whether to enable "Gzip Compression". If enabled, files of specific file types will be compressed, otherwise, no files will be compressed. Default to true. Available in v1.13.0+.`,
				},
				resource.Attribute{
					Name:        "x_forwarded_for",
					Description: `(Optional) Whether to set additional HTTP Header field "X-Forwarded-For" (documented below). Available in v1.13.0+.`,
				},
				resource.Attribute{
					Name:        "acl_status",
					Description: `(Optional) Whether to enable "acl(access control list)", the acl is specified by ` + "`" + `acl_id` + "`" + `. Valid values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. Default to ` + "`" + `off` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "acl_type",
					Description: `(Optional) Mode for handling the acl specified by acl_id. If ` + "`" + `acl_status` + "`" + ` is "on", it is mandatory. Otherwise, it will be ignored. Valid values are ` + "`" + `white` + "`" + ` and ` + "`" + `black` + "`" + `. ` + "`" + `white` + "`" + ` means the Listener can only be accessed by client ip belongs to the acl; ` + "`" + `black` + "`" + ` means the Listener can not be accessed by client ip belongs to the acl.`,
				},
				resource.Attribute{
					Name:        "acl_id",
					Description: `(Optional) the id of access control list to be apply on the listener, is the id of resource alicloud_slb_acl. If ` + "`" + `acl_status` + "`" + ` is "on", it is mandatory. Otherwise, it will be ignored.`,
				},
				resource.Attribute{
					Name:        "established_timeout",
					Description: `(Optional) Timeout of tcp listener established connection idle timeout. Valid value range: [10-900] in seconds. Default to 900.`,
				},
				resource.Attribute{
					Name:        "idle_timeout",
					Description: `(Optional) Timeout of http or https listener established connection idle timeout. Valid value range: [1-60] in seconds. Default to 15.`,
				},
				resource.Attribute{
					Name:        "request_timeout",
					Description: `(Optional) Timeout of http or https listener request (which does not get response from backend) timeout. Valid value range: [1-180] in seconds. Default to 60.`,
				},
				resource.Attribute{
					Name:        "enable_http2",
					Description: `(Optional) Whether to enable https listener support http2 or not. Valid values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. Default to ` + "`" + `on` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tls_cipher_policy",
					Description: `(Optional) Https listener TLS cipher policy. Valid values are ` + "`" + `tls_cipher_policy_1_0` + "`" + `, ` + "`" + `tls_cipher_policy_1_1` + "`" + `, ` + "`" + `tls_cipher_policy_1_2` + "`" + `, ` + "`" + `tls_cipher_policy_1_2_strict` + "`" + `. Default to ` + "`" + `tls_cipher_policy_1_0` + "`" + `. Currently the ` + "`" + `tls_cipher_policy` + "`" + ` can not be updated when load balancer instance is "Shared-Performance".`,
				},
				resource.Attribute{
					Name:        "server_group_id",
					Description: `(Optional) the id of server group to be apply on the listener, is the id of resource ` + "`" + `alicloud_slb_server_group` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "listener_forward",
					Description: `(Optional, ForceNew, Available in 1.40.0+) Whether to enable http redirect to https, Valid values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. Default to ` + "`" + `off` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "forward_port",
					Description: `(Optional, ForceNew, Available in 1.40.0+) The port that http redirect to https.`,
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
					Description: `(Optional) Whether to use the XForwardedFor_proto header to obtain the protocol used by the listener. Default to false. ## Listener fields and protocol mapping load balance support 4 protocal to listen on, they are ` + "`" + `http` + "`" + `,` + "`" + `https` + "`" + `,` + "`" + `tcp` + "`" + `,` + "`" + `udp` + "`" + `, the every listener support which portocal following: listener parameter | support protocol | value range | ------------- | ------------- | ------------- | backend_port | http & https & tcp & udp | 1-65535 | frontend_port | http & https & tcp & udp | 1-65535 | protocol | http & https & tcp & udp | bandwidth | http & https & tcp & udp | -1 / 1-1000 | scheduler | http & https & tcp & udp | wrr rr or wlc | sticky_session | http & https | on or off | sticky_session_type | http & https | insert or server | cookie_timeout | http & https | 1-86400 | cookie | http & https | | persistence_timeout | tcp & udp | 0-3600 | health_check | http & https | on or off | health_check_type | tcp | tcp or http | health_check_domain | http & https & tcp | health_check_method | http & https & tcp | health_check_uri | http & https & tcp | | health_check_connect_port | http & https & tcp & udp | 1-65535 or -520 | healthy_threshold | http & https & tcp & udp | 1-10 | unhealthy_threshold | http & https & tcp & udp | 1-10 | health_check_timeout | http & https & tcp & udp | 1-300 | health_check_interval | http & https & tcp & udp | 1-50 | health_check_http_code | http & https & tcp | http_2xx,http_3xx,http_4xx,http_5xx | server_certificate_id | https | | gzip | http & https | true or false | x_forwarded_for | http & https | | acl_status | http & https & tcp & udp | on or off | acl_type | http & https & tcp & udp | white or black | acl_id | http & https & tcp & udp | the id of resource alicloud_slb_acl| established_timeout | tcp | 10-900| idle_timeout |http & https | 1-60 | request_timeout |http & https | 1-180 | enable_http2 |https | on or off | tls_cipher_policy |https | tls_cipher_policy_1_0, tls_cipher_policy_1_1, tls_cipher_policy_1_2, tls_cipher_policy_1_2_strict | server_group_id | http & https & tcp & udp | the id of resource alicloud_slb_server_group | The listener mapping supports the following: ## Attributes Reference The following attributes are exported:`,
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
					Description: `(Optional) Security certificate ID. ## Import Load balancer listener can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_slb_listener.example "lb-abc123456:22" ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Optional) Security certificate ID. ## Import Load balancer listener can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_slb_listener.example "lb-abc123456:22" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_slb_master_slave_server_group",
			Category:         "Server Load Balancer (SLB)",
			ShortDescription: `Provides a Load Banlancer Master Slave Server Group resource.`,
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
					Description: `(Required, ForceNew) Name of the master slave server group.`,
				},
				resource.Attribute{
					Name:        "servers",
					Description: `(Optional, ForceNew) A list of ECS instances to be added. Only two ECS instances can be supported in one resource. It contains six sub-fields as ` + "`" + `Block server` + "`" + ` follows.`,
				},
				resource.Attribute{
					Name:        "delete_protection_validation",
					Description: `(Optional, Available in 1.63.0+) Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false. ## Block servers The servers mapping supports the following:`,
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
					Description: `(Optional, Available in 1.51.0+) Type of the backend server. Valid value ecs, eni. Default to eni.`,
				},
				resource.Attribute{
					Name:        "server_type",
					Description: `(Optional) The server type of the backend server. Valid value Master, Slave.`,
				},
				resource.Attribute{
					Name:        "is_backup",
					Description: `(Removed from v1.63.0) Determine if the server is executing. Valid value 0, 1. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the master slave server group. ## Import Load balancer master slave server group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_slb_master_slave_server_group.example abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the master slave server group. ## Import Load balancer master slave server group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_slb_master_slave_server_group.example abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_slb_rule",
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
					Description: `(Optional) Name of the forwarding rule. Our plugin provides a default name: "tf-slb-rule".`,
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
					Description: `(Optional, Available in v1.51.0+) Scheduling algorithm, Valid values are ` + "`" + `wrr` + "`" + `, ` + "`" + `rr` + "`" + ` and ` + "`" + `wlc` + "`" + `. Default to "wrr". This parameter is required and takes effect only when ListenerSync is set to off.`,
				},
				resource.Attribute{
					Name:        "sticky_session",
					Description: `(Optional, Available in v1.51.0+) Whether to enable session persistence, Valid values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. Default to ` + "`" + `off` + "`" + `. This parameter is required and takes effect only when ListenerSync is set to off.`,
				},
				resource.Attribute{
					Name:        "sticky_session_type",
					Description: `(Optional, Available in v1.51.0+) Mode for handling the cookie. If ` + "`" + `sticky_session` + "`" + ` is "on", it is mandatory. Otherwise, it will be ignored. Valid values are ` + "`" + `insert` + "`" + ` and ` + "`" + `server` + "`" + `. ` + "`" + `insert` + "`" + ` means it is inserted from Server Load Balancer; ` + "`" + `server` + "`" + ` means the Server Load Balancer learns from the backend server.`,
				},
				resource.Attribute{
					Name:        "cookie_timeout",
					Description: `(Optional, Available in v1.51.0+) Cookie timeout. It is mandatory when ` + "`" + `sticky_session` + "`" + ` is "on" and ` + "`" + `sticky_session_type` + "`" + ` is "insert". Otherwise, it will be ignored. Valid value range: [1-86400] in seconds.`,
				},
				resource.Attribute{
					Name:        "cookie",
					Description: `(Optional, Available in v1.51.0+) The cookie configured on the server. It is mandatory when ` + "`" + `sticky_session` + "`" + ` is "on" and ` + "`" + `sticky_session_type` + "`" + ` is "server". Otherwise, it will be ignored. Valid value：String in line with RFC 2965, with length being 1- 200. It only contains characters such as ASCII codes, English letters and digits instead of the comma, semicolon or spacing, and it cannot start with $.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `(Optional, Available in v1.51.0+) Whether to enable health check. Valid values are` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. TCP and UDP listener's HealthCheck is always on, so it will be ignore when launching TCP or UDP listener. This parameter is required and takes effect only when ListenerSync is set to off.`,
				},
				resource.Attribute{
					Name:        "health_check_domain",
					Description: `(Optional, Available in v1.51.0+) Domain name used for health check. When it used to launch TCP listener, ` + "`" + `health_check_type` + "`" + ` must be "http". Its length is limited to 1-80 and only characters such as letters, digits, ‘-‘ and ‘.’ are allowed. When it is not set or empty, Server Load Balancer uses the private network IP address of each backend server as Domain used for health check.`,
				},
				resource.Attribute{
					Name:        "health_check_uri",
					Description: `(Optional, Available in v1.51.0+) URI used for health check. When it used to launch TCP listener, ` + "`" + `health_check_type` + "`" + ` must be "http". Its length is limited to 1-80 and it must start with /. Only characters such as letters, digits, ‘-’, ‘/’, ‘.’, ‘%’, ‘?’, #’ and ‘&’ are allowed.`,
				},
				resource.Attribute{
					Name:        "health_check_connect_port",
					Description: `(Optional, Available in v1.51.0+) Port used for health check. Valid value range: [1-65535]. Default to "None" means the backend server port is used.`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `(Optional, Available in v1.51.0+) Threshold determining the result of the health check is success. It is required when ` + "`" + `health_check` + "`" + ` is on. Valid value range: [1-10] in seconds. Default to 3.`,
				},
				resource.Attribute{
					Name:        "unhealthy_threshold",
					Description: `(Optional, Available in v1.51.0+) Threshold determining the result of the health check is fail. It is required when ` + "`" + `health_check` + "`" + ` is on. Valid value range: [1-10] in seconds. Default to 3.`,
				},
				resource.Attribute{
					Name:        "health_check_timeout",
					Description: `(Optional, Available in v1.51.0+) Maximum timeout of each health check response. It is required when ` + "`" + `health_check` + "`" + ` is on. Valid value range: [1-300] in seconds. Default to 5. Note: If ` + "`" + `health_check_timeout` + "`" + ` < ` + "`" + `health_check_interval` + "`" + `, its will be replaced by ` + "`" + `health_check_interval` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_check_interval",
					Description: `(Optional, Available in v1.51.0+) Time interval of health checks. It is required when ` + "`" + `health_check` + "`" + ` is on. Valid value range: [1-50] in seconds. Default to 2.`,
				},
				resource.Attribute{
					Name:        "health_check_http_code",
					Description: `(Optional, Available in v1.51.0+) Regular health check HTTP status code. Multiple codes are segmented by “,”. It is required when ` + "`" + `health_check` + "`" + ` is on. Default to ` + "`" + `http_2xx` + "`" + `. Valid values are: ` + "`" + `http_2xx` + "`" + `, ` + "`" + `http_3xx` + "`" + `, ` + "`" + `http_4xx` + "`" + ` and ` + "`" + `http_5xx` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "listener_sync",
					Description: `(Optional, Available in v1.51.0+) Indicates whether a forwarding rule inherits the settings of a health check , session persistence, and scheduling algorithm from a listener. Default to on.`,
				},
				resource.Attribute{
					Name:        "delete_protection_validation",
					Description: `(Optional, Available in 1.63.0+) Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the forwarding rule. ## Import Load balancer forwarding rule can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_slb_rule.example rule-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the forwarding rule. ## Import Load balancer forwarding rule can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_slb_rule.example rule-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_slb_server_certificate",
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
					Description: `(Optional, ForceNew) the content of the ssl certificate. where ` + "`" + `alicloud_certificate_id` + "`" + ` is null, it is required, otherwise it is ignored.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Optional, ForceNew) the content of privat key of the ssl certificate specified by ` + "`" + `server_certificate` + "`" + `. where ` + "`" + `alicloud_certificate_id` + "`" + ` is null, it is required, otherwise it is ignored.`,
				},
				resource.Attribute{
					Name:        "alicloud_certificate_id",
					Description: `(Optional, ForceNew) an id of server certificate ssued/proxied by alibaba cloud. but it is not supported on the international site of alibaba cloud now.`,
				},
				resource.Attribute{
					Name:        "alicloud_certificate_name",
					Description: `(Optional, ForceNew) the name of the certificate specified by ` + "`" + `alicloud_certificate_id` + "`" + `.but it is not supported on the international site of alibaba cloud now.`,
				},
				resource.Attribute{
					Name:        "alicloud_certificate_region_id",
					Description: `(Optional, ForceNew, Available in 1.69.0+) the region of the certificate specified by ` + "`" + `alicloud_certificate_id` + "`" + `. but it is not supported on the international site of alibaba cloud now.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional, ForceNew, Available in 1.58.0+) The Id of resource group which the slb server certificate belongs.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Available in v1.66.0+) A mapping of tags to assign to the resource. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Id of Server Certificate (SSL Certificate). ## Import Server Load balancer Server Certificate can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_slb_server_certificate.example abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Id of Server Certificate (SSL Certificate). ## Import Server Load balancer Server Certificate can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_slb_server_certificate.example abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_slb_server_group",
			Category:         "Server Load Balancer (SLB)",
			ShortDescription: `Provides a Load Banlancer Virtual Backend Server Group resource.`,
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
					Description: `(Optional, Available in 1.63.0+) Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false. ## Block servers The servers mapping supports the following:`,
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
					Description: `(Optional, Available in 1.51.0+) Type of the backend server. Valid value ecs, eni. Default to eni. ## Attributes Reference The following attributes are exported:`,
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
					Description: `A list of ECS instances that have be added. ## Import Load balancer backend server group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_slb_server_group.example abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `A list of ECS instances that have be added. ## Import Load balancer backend server group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_slb_server_group.example abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_snapshot",
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
					Description: `(Required, ForceNew) The source disk ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, ForceNew) Name of the snapshot. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-", ".", "_", and must not begin or end with a hyphen, and must not begin with http:// or https://. Default value is null.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, ForceNew) Description of the snapshot. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. ### Timeouts ->`,
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
					Description: `The snapshot ID. ## Import Snapshot can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_snapshot.snapshot s-abc1234567890000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The snapshot ID. ## Import Snapshot can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_snapshot.snapshot s-abc1234567890000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_snapshot_policy",
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
					Description: `(Optional) The snapshot policy name.`,
				},
				resource.Attribute{
					Name:        "repeat_weekdays",
					Description: `(Required) The automatic snapshot repetition dates. The unit of measurement is day and the repeating cycle is a week. Value range: [1, 7], which represents days starting from Monday to Sunday, for example 1 indicates Monday. When you want to schedule multiple automatic snapshot tasks for a disk in a week, you can set the RepeatWeekdays to an array. - A maximum of seven time points can be selected. - The format is an JSON array of ["1", "2", … "7"] and the time points are separated by commas (,).`,
				},
				resource.Attribute{
					Name:        "retention_days",
					Description: `(Required) The snapshot retention time, and the unit of measurement is day. Optional values: - -1: The automatic snapshots are retained permanently. - [1, 65536]: The number of days retained. Default value: -1.`,
				},
				resource.Attribute{
					Name:        "time_points",
					Description: `(Required) The automatic snapshot creation schedule, and the unit of measurement is hour. Value range: [0, 23], which represents from 00:00 to 24:00, for example 1 indicates 01:00. When you want to schedule multiple automatic snapshot tasks for a disk in a day, you can set the TimePoints to an array. - A maximum of 24 time points can be selected. - The format is an JSON array of ["0", "1", … "23"] and the time points are separated by commas (,). ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The snapshot policy ID. ## Import Snapshot can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_snapshot.snapshot s-abc1234567890000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The snapshot policy ID. ## Import Snapshot can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_snapshot.snapshot s-abc1234567890000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_snat_entry",
			Category:         "VPC",
			ShortDescription: `Provides a Alicloud snat resource.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"snat",
				"entry",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "snat_table_id",
					Description: `(Required, ForceNew) The value can get from ` + "`" + `alicloud_nat_gateway` + "`" + ` Attributes "snat_table_ids".`,
				},
				resource.Attribute{
					Name:        "source_vswitch_id",
					Description: `(Optional, ForceNew) The vswitch ID.`,
				},
				resource.Attribute{
					Name:        "source_cidr",
					Description: `(Optional, ForceNew, Available in 1.71.1+) The private network segment of Ecs. This parameter and the ` + "`" + `source_vswitch_id` + "`" + ` parameter are mutually exclusive and cannot appear at the same time.`,
				},
				resource.Attribute{
					Name:        "snat_entry_name",
					Description: `(Optional, Available in 1.71.2+) The name of snat entry.`,
				},
				resource.Attribute{
					Name:        "snat_ip",
					Description: `(Required) The SNAT ip address, the ip must along bandwidth package public ip which ` + "`" + `alicloud_nat_gateway` + "`" + ` argument ` + "`" + `bandwidth_packages` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the snat entry. The value formats as ` + "`" + `<snat_table_id>:<snat_entry_id>` + "`" + ``,
				},
				resource.Attribute{
					Name:        "snat_entry_id",
					Description: `The id of the snat entry on the server. ## Import Snat Entry can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_snat_entry.foo stb-1aece3:snat-232ce2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the snat entry. The value formats as ` + "`" + `<snat_table_id>:<snat_entry_id>` + "`" + ``,
				},
				resource.Attribute{
					Name:        "snat_entry_id",
					Description: `The id of the snat entry on the server. ## Import Snat Entry can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_snat_entry.foo stb-1aece3:snat-232ce2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ssl_vpn_client_cert",
			Category:         "VPN",
			ShortDescription: `Provides a Alicloud SSL VPN Client Cert resource.`,
			Description:      ``,
			Keywords: []string{
				"vpn",
				"ssl",
				"client",
				"cert",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ssl_vpn_server",
			Category:         "VPN",
			ShortDescription: `Provides a Alicloud SSL VPN server resource.`,
			Description:      ``,
			Keywords: []string{
				"vpn",
				"ssl",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the SSL-VPN server.`,
				},
				resource.Attribute{
					Name:        "vpn_gateway_id",
					Description: `(Required, ForceNew) The ID of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "client_ip_pool",
					Description: `(Required) The CIDR block from which access addresses are allocated to the virtual network interface card of the client.`,
				},
				resource.Attribute{
					Name:        "local_subnet",
					Description: `(Required) The CIDR block to be accessed by the client through the SSL-VPN connection. It supports to set multi CIDRs by comma join ways, like ` + "`" + `10.0.1.0/24,10.0.2.0/24,10.0.3.0/24` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) The protocol used by the SSL-VPN server. Valid value: UDP(default) |TCP`,
				},
				resource.Attribute{
					Name:        "cipher",
					Description: `(Optional) The encryption algorithm used by the SSL-VPN server. Valid value: AES-128-CBC (default)| AES-192-CBC | AES-256-CBC | none`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The port used by the SSL-VPN server. The default value is 1194.The following ports cannot be used: [22, 2222, 22222, 9000, 9001, 9002, 7505, 80, 443, 53, 68, 123, 4510, 4560, 500, 4500].`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the SSL-VPN server.`,
				},
				resource.Attribute{
					Name:        "internet_ip",
					Description: `The internet IP of the SSL-VPN server.`,
				},
				resource.Attribute{
					Name:        "connections",
					Description: `The number of current connections.`,
				},
				resource.Attribute{
					Name:        "max_connections",
					Description: `The maximum number of connections. ## Import SSL-VPN server can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ssl_vpn_server.example vss-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the SSL-VPN server.`,
				},
				resource.Attribute{
					Name:        "internet_ip",
					Description: `The internet IP of the SSL-VPN server.`,
				},
				resource.Attribute{
					Name:        "connections",
					Description: `The number of current connections.`,
				},
				resource.Attribute{
					Name:        "max_connections",
					Description: `The maximum number of connections. ## Import SSL-VPN server can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ssl_vpn_server.example vss-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_vpc",
			Category:         "VPC",
			ShortDescription: `Provides a Alicloud VPC resource.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Required, ForceNew) The CIDR block for the VPC.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the VPC. Defaults to null.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The VPC description. Defaults to null.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional, ForceNew, Available in 1.40.0+) The Id of resource group which the VPC belongs.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Available in v1.55.3+) A mapping of tags to assign to the resource. ### Timeouts ->`,
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
					Description: `The route table ID of the router created by default on VPC creation. ## Import VPC can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_vpc.example vpc-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `The route table ID of the router created by default on VPC creation. ## Import VPC can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_vpc.example vpc-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_vpn_connection",
			Category:         "VPN",
			ShortDescription: `Provides a Alicloud VPN connection resource.`,
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
					Description: `The configurations of phase-two negotiation. ## Import VPN connection can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_vpn_connection.example vco-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `The configurations of phase-two negotiation. ## Import VPN connection can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_vpn_connection.example vco-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_vpn_customer_gateway",
			Category:         "VPN",
			ShortDescription: `Provides a Alicloud VPN customer gateway resource.`,
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
					Description: `The ID of the VPN customer gateway instance id. ## Import VPN customer gateway can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_vpn_customer_gateway.example cgw-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the VPN customer gateway instance id. ## Import VPN customer gateway can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_vpn_customer_gateway.example cgw-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_vpn_gateway",
			Category:         "VPN",
			ShortDescription: `Provides a Alicloud VPN gateway resource.`,
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
					Description: `(Optional, ForceNew, Available in v1.56.0+) The VPN belongs the vswitch_id, the field can't be changed. ## Attributes Reference The following attributes are exported:`,
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
					Description: `The business status of the VPN gateway. ## Import VPN gateway can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_vpn_gateway.example vpn-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `The business status of the VPN gateway. ## Import VPN gateway can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_vpn_gateway.example vpn-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_vpn_route_entry",
			Category:         "VPN",
			ShortDescription: `Provides a Alicloud VPN Route Entry resource.`,
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
					Description: `The combination id of the vpn route entry. ## Import VPN route entry can be imported using the id(VpnGatewayId +":"+ NextHop +":"+ RouteDest), e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_vpn_route_entry.example vpn-abc123456:vco-abc123456:10.0.0.10/24 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The combination id of the vpn route entry. ## Import VPN route entry can be imported using the id(VpnGatewayId +":"+ NextHop +":"+ RouteDest), e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_vpn_route_entry.example vpn-abc123456:vco-abc123456:10.0.0.10/24 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_vswitch",
			Category:         "VPC",
			ShortDescription: `Provides a Alicloud VPC switch resource.`,
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
					Description: `(Optional) The switch description. Defaults to null.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Available in v1.55.3+) A mapping of tags to assign to the resource. ### Timeouts ->`,
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
					Description: `The description of the switch. ## Import Vswitch can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_vswitch.example vsw-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `The description of the switch. ## Import Vswitch can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_vswitch.example vsw-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_waf_domain",
			Category:         "Web Application Firewall(WAF)",
			ShortDescription: `Provides a Web Application Firewall Domain resource.`,
			Description:      ``,
			Keywords: []string{
				"web",
				"application",
				"firewall",
				"waf",
				"domain",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_type",
					Description: `(Optional) The type of the WAF cluster. Valid values: "PhysicalCluster" and "VirtualCluster". Default to "PhysicalCluster".`,
				},
				resource.Attribute{
					Name:        "connection_time",
					Description: `(Optional) The connection timeout for WAF exclusive clusters. Unit: seconds.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Required, ForceNew) The domain that you want to add to WAF.`,
				},
				resource.Attribute{
					Name:        "http2_port",
					Description: `(Optional) List of the HTTP 2.0 ports.`,
				},
				resource.Attribute{
					Name:        "http_port",
					Description: `(Optional) List of the HTTP ports`,
				},
				resource.Attribute{
					Name:        "http_to_user_ip",
					Description: `(Optional) Specifies whether to enable the HTTP back-to-origin feature. After this feature is enabled, the WAF instance can use HTTP to forward HTTPS requests to the origin server. By default, port 80 is used to forward the requests to the origin server. Valid values: "On" and "Off". Default to "Off".`,
				},
				resource.Attribute{
					Name:        "https_port",
					Description: `(Optional) List of the HTTPS ports`,
				},
				resource.Attribute{
					Name:        "https_redirect",
					Description: `(Optional) Specifies whether to redirect HTTP requests as HTTPS requests. Valid values: "On" and "Off". Default to "Off".`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, ForceNew) The ID of the WAF instance.`,
				},
				resource.Attribute{
					Name:        "is_access_product",
					Description: `(Required) Specifies whether to configure a Layer-7 proxy, such as Anti-DDoS Pro or CDN, to filter the inbound traffic before it is forwarded to WAF. Valid values: "On" and "Off". Default to "Off".`,
				},
				resource.Attribute{
					Name:        "load_balancing",
					Description: `(Optional) The load balancing algorithm that is used to forward requests to the origin. Valid values: "IpHash" and "RoundRobin". Default to "IpHash".`,
				},
				resource.Attribute{
					Name:        "log_headers",
					Description: `(Optional) The key-value pair that is used to mark the traffic that flows through WAF to the domain. Each item contains two field:`,
				},
				resource.Attribute{
					Name:        "read_time",
					Description: `(Optional) The read timeout of a WAF exclusive cluster. Unit: seconds.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional) The ID of the resource group to which the queried domain belongs in Resource Management. By default, no value is specified, indicating that the domain belongs to the default resource group.`,
				},
				resource.Attribute{
					Name:        "source_ips",
					Description: `(Optional) List of the IP address or domain of the origin server to which the specified domain points.`,
				},
				resource.Attribute{
					Name:        "write_time",
					Description: `(Optional) The timeout period for a WAF exclusive cluster write connection. Unit: seconds. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `This resource id. It formats as ` + "`" + `<instance_id>:<domain>` + "`" + ``,
				},
				resource.Attribute{
					Name:        "cname",
					Description: `The CNAME record assigned by the WAF instance to the specified domain. ## Import WAF domain can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_waf_domain.domain waf-132435:www.domain.com ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `This resource id. It formats as ` + "`" + `<instance_id>:<domain>` + "`" + ``,
				},
				resource.Attribute{
					Name:        "cname",
					Description: `The CNAME record assigned by the WAF instance to the specified domain. ## Import WAF domain can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_waf_domain.domain waf-132435:www.domain.com ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_waf_instance",
			Category:         "Web Application Firewall(WAF)",
			ShortDescription: `Provides a Web Application Firewall Instance resource.`,
			Description:      ``,
			Keywords: []string{
				"web",
				"application",
				"firewall",
				"waf",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "big_screen",
					Description: `(Required, String) Specify whether big screen is supported. Valid values: ["0", "1"]. "0" for false and "1" for true.`,
				},
				resource.Attribute{
					Name:        "exclusive_ip_package",
					Description: `(Required, String) Specify the number of exclusive WAF IP addresses.`,
				},
				resource.Attribute{
					Name:        "ext_bandwidth",
					Description: `(Required, String) The extra bandwidth. Unit: Mbit/s.`,
				},
				resource.Attribute{
					Name:        "ext_domain_package",
					Description: `(Required, String) The number of extra domains.`,
				},
				resource.Attribute{
					Name:        "log_storage",
					Description: `(Required, String) Log storage size. Unit: T. Valid values: [3, 5, 10, 20, 50].`,
				},
				resource.Attribute{
					Name:        "log_time",
					Description: `(Required, String) Log storage period. Unit: day. Valid values: [180, 360].`,
				},
				resource.Attribute{
					Name:        "modify_type",
					Description: `(Optional) Type of configuration change. Valid value: Upgrade.`,
				},
				resource.Attribute{
					Name:        "package_code",
					Description: `(Required, String) Subscription plan:`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(ForceNew) Service time of Web Application Firewall.`,
				},
				resource.Attribute{
					Name:        "prefessional_service",
					Description: `(Required, String) Specify whether professional service is supported. Valid values: ["true", "false"]`,
				},
				resource.Attribute{
					Name:        "renew_period",
					Description: `(ForceNew) Renewal period of WAF service. Unit: month`,
				},
				resource.Attribute{
					Name:        "renewal_status",
					Description: `(ForceNew) Renewal status of WAF service. Valid values:`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional) The resource group ID.`,
				},
				resource.Attribute{
					Name:        "subscription_type",
					Description: `(Required, String) Subscription of WAF service. Valid values: ["Subscription", "PayAsYouGo"].`,
				},
				resource.Attribute{
					Name:        "waf_log",
					Description: `(Required, String) Specify whether Log service is supported. Valid values: ["true", "false"] ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `This resource instance id.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the instance. ## Import WAF instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_waf_instance.default waf-cn-132435 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `This resource instance id.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the instance. ## Import WAF instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_waf_instance.default waf-cn-132435 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_yundun_bastionhost_instance",
			Category:         "Cloud Bastionhost",
			ShortDescription: `Provides a Alicloud Cloud Bastionhost Instance Resource.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"bastionhost",
				"yundun",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "plan_code",
					Description: `(Required) Plan code of the Cloud Bastionhost to produce. (alpha.professional, alpha.basic, alpha.premium)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) Description of the instance. This name can have a string of 1 to 63 characters.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(ForceNew) Duration for initially producing the instance. Valid values: [1~9], 12, 24, 36. Default to 1. At present, the provider does not support modify "period".`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Required, ForceNew) vSwtich ID configured to bastionhost`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Required) security group IDs configured to bastionhost`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Available in v1.67.0+) A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional, Available in v1.87.0+) The Id of resource group which the Bastionhost Instance belongs. If not set, the resource is created in the default resource group. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the instance resource of Yundun_bastionhost. ## Import Yundun_bastionhost instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_yundun_bastionhost.example bastionhost-exampe123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the instance resource of Yundun_bastionhost. ## Import Yundun_bastionhost instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_yundun_bastionhost.example bastionhost-exampe123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_yundun_dbaudit_instance",
			Category:         "Cloud DBAudit",
			ShortDescription: `Provides a Alicloud Cloud DBaudit Instance Resource.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"dbaudit",
				"yundun",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "plan_code",
					Description: `(Required) Plan code of the Cloud DBAudit to produce. (alpha.professional, alpha.basic, alpha.premium)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) Description of the instance. This name can have a string of 1 to 63 characters.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Required, ForceNew) Duration for initially producing the instance. Valid values: [1~9], 12, 24, 36. Default to 12. At present, the provider does not support modify "period".`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Required, ForceNew) vSwtich ID configured to audit`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Available in v1.67.0+) A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional, Available in v1.87.0+) The Id of resource group which the DBaudit Instance belongs. If not set, the resource is created in the default resource group. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the instance resource of Yundun_dbaudit. ## Import Yundun_dbaudit instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_yundun_dbaudit_instance.example dbaudit-exampe123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the instance resource of Yundun_dbaudit. ## Import Yundun_dbaudit instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_yundun_dbaudit_instance.example dbaudit-exampe123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"alicloud_actiontrail":                         0,
		"alicloud_adb_account":                         1,
		"alicloud_adb_backup_policy":                   2,
		"alicloud_adb_cluster":                         3,
		"alicloud_adb_connection":                      4,
		"alicloud_alidns_domain_group":                 5,
		"alicloud_alidns_record":                       6,
		"alicloud_alikafka_consumer_group":             7,
		"alicloud_alikafka_instance":                   8,
		"alicloud_alikafka_sasl_acl":                   9,
		"alicloud_alikafka_sasl_user":                  10,
		"alicloud_alikafka_topic":                      11,
		"alicloud_api_gateway_api":                     12,
		"alicloud_api_gateway_app":                     13,
		"alicloud_api_gateway_app_attachment":          14,
		"alicloud_api_gateway_group":                   15,
		"alicloud_api_gateway_vpc_access":              16,
		"alicloud_auto_provisioning_group":             17,
		"alicloud_cas_certificate":                     18,
		"alicloud_cassandra_cluster":                   19,
		"alicloud_cassandra_data_center":               20,
		"alicloud_cdn_domain":                          21,
		"alicloud_cdn_domain_config":                   22,
		"alicloud_cdn_domain_new":                      23,
		"alicloud_cen_bandwidth_limit":                 24,
		"alicloud_cen_bandwidth_package":               25,
		"alicloud_cen_bandwidth_package_attachment":    26,
		"alicloud_cen_flowlog":                         27,
		"alicloud_cen_instance":                        28,
		"alicloud_cen_instance_attachment":             29,
		"alicloud_cen_instance_grant":                  30,
		"alicloud_cen_private_zone":                    31,
		"alicloud_cen_route_entry":                     32,
		"alicloud_cen_route_map":                       33,
		"alicloud_cen_vbr_health_check":                34,
		"alicloud_cloud_connect_network":               35,
		"alicloud_cloud_connect_network_attachment":    36,
		"alicloud_cloud_connect_network_grant":         37,
		"alicloud_cms_alarm":                           38,
		"alicloud_cms_site_monitor":                    39,
		"alicloud_common_bandwidth_package":            40,
		"alicloud_common_bandwidth_package_attachment": 41,
		"alicloud_container_cluster":                   42,
		"alicloud_cr_ee_namespace":                     43,
		"alicloud_cr_ee_repo":                          44,
		"alicloud_cr_ee_sync_rule":                     45,
		"alicloud_cr_namespace":                        46,
		"alicloud_cr_repo":                             47,
		"alicloud_cs_application":                      48,
		"alicloud_cs_kubernetes":                       49,
		"alicloud_cs_kubernetes_autoscaler":            50,
		"alicloud_cs_managed_kubernetes":               51,
		"alicloud_cs_serverless_kubernetes":            52,
		"alicloud_cs_swarm":                            53,
		"alicloud_datahub_project":                     54,
		"alicloud_datahub_subscription":                55,
		"alicloud_datahub_topic":                       56,
		"alicloud_db_account":                          57,
		"alicloud_db_account_privilege":                58,
		"alicloud_db_backup_policy":                    59,
		"alicloud_db_connection":                       60,
		"alicloud_db_database":                         61,
		"alicloud_db_instance":                         62,
		"alicloud_db_read_write_splitting_connection":  63,
		"alicloud_db_readonly_instance":                64,
		"alicloud_ddosbgp_instance":                    65,
		"alicloud_ddoscoo_instance":                    66,
		"alicloud_ddoscoo_scheduler_rule":              67,
		"alicloud_disk":                                68,
		"alicloud_disk_attachment":                     69,
		"alicloud_dms_enterprise_instance":             70,
		"alicloud_dms_enterprise_user":                 71,
		"alicloud_dns":                                 72,
		"alicloud_dns_domain":                          73,
		"alicloud_dns_domain_attachment":               74,
		"alicloud_dns_group":                           75,
		"alicloud_dns_instance":                        76,
		"alicloud_dns_record":                          77,
		"alicloud_drds_instance":                       78,
		"alicloud_eci_image_cache":                     79,
		"alicloud_edas_application":                    80,
		"alicloud_edas_application_deployment":         81,
		"alicloud_edas_application_scale":              82,
		"alicloud_edas_cluster":                        83,
		"alicloud_edas_deploy_group":                   84,
		"alicloud_edas_instance_cluster_attachment":    85,
		"alicloud_edas_slb_attachment":                 86,
		"alicloud_eip":                                 87,
		"alicloud_eip_association":                     88,
		"alicloud_elasticsearch_instance":              89,
		"alicloud_emr_cluster":                         90,
		"alicloud_ess_alarm":                           91,
		"alicloud_ess_attachment":                      92,
		"alicloud_ess_notification":                    93,
		"alicloud_ess_scaling_configuration":           94,
		"alicloud_ess_scaling_group":                   95,
		"alicloud_ess_lifecycle_hook":                  96,
		"alicloud_ess_scaling_rule":                    97,
		"alicloud_ess_scalinggroup_vserver_groups":     98,
		"alicloud_ess_schedule":                        99,
		"alicloud_ess_scheduled_task":                  100,
		"alicloud_fc_function":                         101,
		"alicloud_fc_service":                          102,
		"alicloud_fc_trigger":                          103,
		"alicloud_forward_entry":                       104,
		"alicloud_gpdb_connection":                     105,
		"alicloud_gpdb_instance":                       106,
		"alicloud_hbase_instance":                      107,
		"alicloud_image":                               108,
		"alicloud_image_copy":                          109,
		"alicloud_image_export":                        110,
		"alicloud_image_import":                        111,
		"alicloud_image_share_permission":              112,
		"alicloud_instance":                            113,
		"alicloud_key_pair":                            114,
		"alicloud_key_pair_attachment":                 115,
		"alicloud_kms_alias":                           116,
		"alicloud_kms_ciphertext":                      117,
		"alicloud_kms_key":                             118,
		"alicloud_kms_key_version":                     119,
		"alicloud_kms_secret":                          120,
		"alicloud_kvstore_account":                     121,
		"alicloud_kvstore_backup_policy":               122,
		"alicloud_kvstore_instance":                    123,
		"alicloud_launch_template":                     124,
		"alicloud_log_alert":                           125,
		"alicloud_log_audit":                           126,
		"alicloud_log_dashboard":                       127,
		"alicloud_log_machine_group":                   128,
		"alicloud_log_project":                         129,
		"alicloud_log_store":                           130,
		"alicloud_log_store_index":                     131,
		"alicloud_logtail_attachment":                  132,
		"alicloud_logtail_config":                      133,
		"alicloud_market_order":                        134,
		"alicloud_maxcompute_project":                  135,
		"alicloud_mns_queue":                           136,
		"alicloud_mns_topic":                           137,
		"alicloud_mns_topic_subscription":              138,
		"alicloud_mongodb_instance":                    139,
		"mongodb_sharding_instance":                    140,
		"alicloud_nas_access_group":                    141,
		"alicloud_nas_access_rule":                     142,
		"alicloud_nas_file_system":                     143,
		"alicloud_nas_mount_target":                    144,
		"alicloud_nat_gateway":                         145,
		"alicloud_network_acl":                         146,
		"alicloud_network_acl_attachment":              147,
		"alicloud_network_acl_entries":                 148,
		"alicloud_network_interface":                   149,
		"alicloud_network_interface_attachment":        150,
		"alicloud_ons_group":                           151,
		"alicloud_ons_instance":                        152,
		"alicloud_ons_topic":                           153,
		"alicloud_oss_bucket":                          154,
		"alicloud_oss_bucket_object":                   155,
		"alicloud_ots_instance":                        156,
		"alicloud_ots_instance_attachment":             157,
		"alicloud_ots_table":                           158,
		"alicloud_polardb_account":                     159,
		"alicloud_polardb_account_privilege":           160,
		"alicloud_polardb_backup_policy":               161,
		"alicloud_polardb_cluster":                     162,
		"alicloud_polardb_database":                    163,
		"alicloud_polardb_endpoint":                    164,
		"alicloud_polardb_endpoint_address":            165,
		"alicloud_pvtz_zone":                           166,
		"alicloud_pvtz_zone_attachment":                167,
		"alicloud_pvtz_zone_record":                    168,
		"alicloud_ram_access_key":                      169,
		"alicloud_ram_account_alias":                   170,
		"alicloud_ram_account_password_policy":         171,
		"alicloud_ram_alias":                           172,
		"alicloud_ram_group":                           173,
		"alicloud_ram_group_membership":                174,
		"alicloud_ram_group_policy_attachment":         175,
		"alicloud_ram_login_profile":                   176,
		"alicloud_ram_policy":                          177,
		"alicloud_ram_role":                            178,
		"alicloud_ram_role_attachment":                 179,
		"alicloud_ram_role_policy_attachment":          180,
		"alicloud_ram_user":                            181,
		"alicloud_ram_user_policy_attachment":          182,
		"alicloud_reserved_instance":                   183,
		"alicloud_resource_manager_account":            184,
		"alicloud_resource_manager_folder":             185,
		"alicloud_resource_manager_handshake":          186,
		"alicloud_resource_manager_policy":             187,
		"alicloud_resource_manager_policy_version":     188,
		"alicloud_resource_manager_resource_directory": 189,
		"alicloud_resource_manager_resource_group":     190,
		"alicloud_resource_manager_role":               191,
		"alicloud_route_entry":                         192,
		"alicloud_route_table":                         193,
		"alicloud_route_table_attachment":              194,
		"alicloud_router_interface":                    195,
		"alicloud_router_interface_connection":         196,
		"alicloud_sag_acl":                             197,
		"alicloud_sag_acl_rule":                        198,
		"alicloud_sag_client_user":                     199,
		"alicloud_sag_dnat_entry":                      200,
		"alicloud_sag_qos":                             201,
		"alicloud_sag_qos_car":                         202,
		"alicloud_sag_qos_policy":                      203,
		"alicloud_sag_snat_entry":                      204,
		"alicloud_security_group":                      205,
		"alicloud_security_group_rule":                 206,
		"alicloud_slb":                                 207,
		"alicloud_slb_acl":                             208,
		"alicloud_slb_attachment":                      209,
		"alicloud_slb_backend_server":                  210,
		"alicloud_slb_ca_certificate":                  211,
		"alicloud_slb_domain_extension":                212,
		"alicloud_slb_listener":                        213,
		"alicloud_slb_master_slave_server_group":       214,
		"alicloud_slb_rule":                            215,
		"alicloud_slb_server_certificate":              216,
		"alicloud_slb_server_group":                    217,
		"alicloud_snapshot":                            218,
		"alicloud_snapshot_policy":                     219,
		"alicloud_snat_entry":                          220,
		"alicloud_ssl_vpn_client_cert":                 221,
		"alicloud_ssl_vpn_server":                      222,
		"alicloud_vpc":                                 223,
		"alicloud_vpn_connection":                      224,
		"alicloud_vpn_customer_gateway":                225,
		"alicloud_vpn_gateway":                         226,
		"alicloud_vpn_route_entry":                     227,
		"alicloud_vswitch":                             228,
		"alicloud_waf_domain":                          229,
		"alicloud_waf_instance":                        230,
		"alicloud_yundun_bastionhost_instance":         231,
		"alicloud_yundun_dbaudit_instance":             232,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
