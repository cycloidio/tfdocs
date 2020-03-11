package baiducloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "baiducloud_acl",
			Category:         "VPC Resources",
			ShortDescription: `Provide a resource to create an ACL Rule.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"acl",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "action",
					Description: `(Required) Action of the acl. Valid values are allow and deny.`,
				},
				resource.Attribute{
					Name:        "destination_ip_address",
					Description: `(Required) Destination ip address of the acl.`,
				},
				resource.Attribute{
					Name:        "destination_port",
					Description: `(Required) Destination port of the acl.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Required, ForceNew) Direction of the acl. Valid values are ingress and egress, respectively indicating the inbound of the rule and the outbound rule.`,
				},
				resource.Attribute{
					Name:        "position",
					Description: `(Required) Position of the acl, representing the priority of the acl rule. The value should be 1-5000 and cannot be duplicated with existing entries. The smaller the value, the higher the priority, and the rule matching order is to match the priority from high to low.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Protocol of the acl, available values are all, tcp, udp and icmp.`,
				},
				resource.Attribute{
					Name:        "source_ip_address",
					Description: `(Required) Source ip address of the acl.`,
				},
				resource.Attribute{
					Name:        "source_port",
					Description: `(Required) Source port of the acl.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required, ForceNew) Subnet ID of the acl.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the acl.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "baiducloud_appblb",
			Category:         "APPBLB Resources",
			ShortDescription: `Provide a resource to create an APPBLB.`,
			Description:      ``,
			Keywords: []string{
				"appblb",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required, ForceNew) The subnet ID to which the LoadBalance instance belongs`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, ForceNew) The VPC short ID to which the LoadBalance instance belongs`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) LoadBalance's description, length must be between 0 and 450 bytes, and support Chinese`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) LoadBalance instance's name, length must be between 1 and 65 bytes, and will be automatically generated if not set`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, ForceNew) Tags, do not support modify ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `LoadBalance instance's service IP, instance can be accessed through this IP`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `Cidr of the network where the LoadBalance instance reside`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `LoadBalance instance's create time`,
				},
				resource.Attribute{
					Name:        "listener",
					Description: `List of listeners mounted under the instance`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Listening port`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Listening protocol type`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `LoadBalance instance's public ip`,
				},
				resource.Attribute{
					Name:        "release_time",
					Description: `LoadBalance instance's auto release time`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `LoadBalance instance's status, see https://cloud.baidu.com/doc/BLB/s/Pjwvxnxdm/#blbstatus for detail`,
				},
				resource.Attribute{
					Name:        "subnet_cidr",
					Description: `Cidr of the subnet which the LoadBalance instance belongs`,
				},
				resource.Attribute{
					Name:        "subnet_name",
					Description: `The subnet name to which the LoadBalance instance belongs`,
				},
				resource.Attribute{
					Name:        "vpc_name",
					Description: `The VPC name to which the LoadBalance instance belongs ## Import APPBLB can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import baiducloud_appblb.default id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "address",
					Description: `LoadBalance instance's service IP, instance can be accessed through this IP`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `Cidr of the network where the LoadBalance instance reside`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `LoadBalance instance's create time`,
				},
				resource.Attribute{
					Name:        "listener",
					Description: `List of listeners mounted under the instance`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Listening port`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Listening protocol type`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `LoadBalance instance's public ip`,
				},
				resource.Attribute{
					Name:        "release_time",
					Description: `LoadBalance instance's auto release time`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `LoadBalance instance's status, see https://cloud.baidu.com/doc/BLB/s/Pjwvxnxdm/#blbstatus for detail`,
				},
				resource.Attribute{
					Name:        "subnet_cidr",
					Description: `Cidr of the subnet which the LoadBalance instance belongs`,
				},
				resource.Attribute{
					Name:        "subnet_name",
					Description: `The subnet name to which the LoadBalance instance belongs`,
				},
				resource.Attribute{
					Name:        "vpc_name",
					Description: `The VPC name to which the LoadBalance instance belongs ## Import APPBLB can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import baiducloud_appblb.default id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "baiducloud_appblb_listener",
			Category:         "APPBLB Resources",
			ShortDescription: `Provide a resource to create an APPBLB Listener.`,
			Description:      ``,
			Keywords: []string{
				"appblb",
				"listener",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "blb_id",
					Description: `(Required, ForceNew) ID of the Application LoadBalance instance`,
				},
				resource.Attribute{
					Name:        "listener_port",
					Description: `(Required, ForceNew) Listening port, range from 1-65535`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required, ForceNew) Listening protocol, support TCP/UDP/HTTP/HTTPS/SSL`,
				},
				resource.Attribute{
					Name:        "scheduler",
					Description: `(Required) Load balancing algorithm, support RoundRobin/LeastConnection/Hash, if protocol is HTTP/HTTPS, only support RoundRobin/LeastConnection`,
				},
				resource.Attribute{
					Name:        "cert_ids",
					Description: `(Optional) Listener bind certifications`,
				},
				resource.Attribute{
					Name:        "client_cert_ids",
					Description: `(Optional) Listener import cert list, only useful when dual_auth is true`,
				},
				resource.Attribute{
					Name:        "dual_auth",
					Description: `(Optional) Listener open dual authorization or not, default false`,
				},
				resource.Attribute{
					Name:        "encryption_protocols",
					Description: `(Optional) Listener encryption protocol, only useful when encryption_type is userDefind, support [sslv3, tlsv10, tlsv11, tlsv12]`,
				},
				resource.Attribute{
					Name:        "encryption_type",
					Description: `(Optional) Listener encryption option, support [compatibleIE, incompatibleIE, userDefind]`,
				},
				resource.Attribute{
					Name:        "ie6_compatible",
					Description: `(Optional) Listener support ie6 option, default true`,
				},
				resource.Attribute{
					Name:        "keep_session_cookie_name",
					Description: `(Optional) CookieName which need to covered, useful when keep_session_type is rewrite`,
				},
				resource.Attribute{
					Name:        "keep_session_timeout",
					Description: `(Optional) KeepSession Cookie timeout time(second), support in [1, 15552000], default 3600s`,
				},
				resource.Attribute{
					Name:        "keep_session_type",
					Description: `(Optional) KeepSessionType option, support insert/rewrite, default insert`,
				},
				resource.Attribute{
					Name:        "keep_session",
					Description: `(Optional) KeepSession or not`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Optional) Listener's policy`,
				},
				resource.Attribute{
					Name:        "redirect_port",
					Description: `(Optional) Redirect HTTP request to HTTPS Listener, HTTPS Listener port set by this parameter`,
				},
				resource.Attribute{
					Name:        "server_timeout",
					Description: `(Optional) Backend server maximum timeout time, only support in [1, 3600] second, default 30s`,
				},
				resource.Attribute{
					Name:        "tcp_session_timeout",
					Description: `(Optional) TCP Listener connection session timeout time(second), default 900, support 10-4000`,
				},
				resource.Attribute{
					Name:        "x_forwarded_for",
					Description: `(Optional) Listener xForwardedFor, determine get client real ip or not, default false The ` + "`" + `policies` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "app_server_group_id",
					Description: `(Required) Policy bind server group id`,
				},
				resource.Attribute{
					Name:        "backend_port",
					Description: `(Required) Backend port`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Required) Policy priority, support in [1, 32768]`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Policy's description`,
				},
				resource.Attribute{
					Name:        "rule_list",
					Description: `(Optional) Policy rule list`,
				},
				resource.Attribute{
					Name:        "app_server_group_name",
					Description: `Policy bind server group name`,
				},
				resource.Attribute{
					Name:        "frontend_port",
					Description: `Frontend port`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Policy's id`,
				},
				resource.Attribute{
					Name:        "port_type",
					Description: `Policy bind port protocol type The ` + "`" + `rule_list` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Rule key`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Rule value`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "baiducloud_appblb_server_group",
			Category:         "APPBLB Resources",
			ShortDescription: `Provide a resource to create an APPBLB Server Group.`,
			Description:      ``,
			Keywords: []string{
				"appblb",
				"server",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "blb_id",
					Description: `(Required, ForceNew) ID of the Application LoadBalance instance`,
				},
				resource.Attribute{
					Name:        "backend_server_list",
					Description: `(Optional) Server group bound backend server list`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Server Group's description, length must be between 0 and 450 bytes, and support Chinese`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the Server Group, length must be between 1 and 65 bytes, and will be automatically generated if not set`,
				},
				resource.Attribute{
					Name:        "port_list",
					Description: `(Optional) Server Group backend port list The ` + "`" + `backend_server_list` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) Backend server instance ID`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Required) Backend server instance weight in this group, range from 0-100`,
				},
				resource.Attribute{
					Name:        "port_list",
					Description: `Backend server open port list`,
				},
				resource.Attribute{
					Name:        "backend_port",
					Description: `Backend open port`,
				},
				resource.Attribute{
					Name:        "health_check_port_type",
					Description: `Health check port protocol type`,
				},
				resource.Attribute{
					Name:        "listener_port",
					Description: `Listener port`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Port bind policy id`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `Port id`,
				},
				resource.Attribute{
					Name:        "port_type",
					Description: `Port protocol type`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Port status, include Alive/Dead/Unknown`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Backend server instance bind private ip The ` + "`" + `port_list` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "backend_port",
					Description: `Backend open port`,
				},
				resource.Attribute{
					Name:        "health_check_port_type",
					Description: `Health check port protocol type`,
				},
				resource.Attribute{
					Name:        "listener_port",
					Description: `Listener port`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Port bind policy id`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `Port id`,
				},
				resource.Attribute{
					Name:        "port_type",
					Description: `Port protocol type`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Port status, include Alive/Dead/Unknown The ` + "`" + `port_list` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `(Required) Server Group port health check protocol, support TCP/UDP/HTTP, default same as port protocol type`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) App Server Group port, range from 1-65535`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Server Group port protocol type, support TCP/UDP/HTTP`,
				},
				resource.Attribute{
					Name:        "health_check_down_retry",
					Description: `(Optional) Server Group health check down retry time, support in [2, 5], default 3`,
				},
				resource.Attribute{
					Name:        "health_check_interval_in_second",
					Description: `(Optional) Server Group health check interval time(second), support in [1, 10], default 3`,
				},
				resource.Attribute{
					Name:        "health_check_normal_status",
					Description: `(Optional) Server Group health check normal http status code, only useful when health_check is HTTP`,
				},
				resource.Attribute{
					Name:        "health_check_port",
					Description: `(Optional) Server Group port health check port, default same as Server Group port`,
				},
				resource.Attribute{
					Name:        "health_check_timeout_in_second",
					Description: `(Optional) Server Group health check timeout(second), support in [1, 60], default 3`,
				},
				resource.Attribute{
					Name:        "health_check_up_retry",
					Description: `(Optional) Server Group health check up retry time, support in [2, 5], default 3`,
				},
				resource.Attribute{
					Name:        "health_check_url_path",
					Description: `(Optional) Server Group health check url path, only useful when health_check is HTTP`,
				},
				resource.Attribute{
					Name:        "udp_health_check_string",
					Description: `(Optional) Server Group udp health check string, if type is UDP, this parameter is required`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Server Group port id`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Server Group port status ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Server Group's status, see https://cloud.baidu.com/doc/BLB/s/Pjwvxnxdm/#blbstatus for detail`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `Server Group's status, see https://cloud.baidu.com/doc/BLB/s/Pjwvxnxdm/#blbstatus for detail`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "baiducloud_auto_snapshot_policy",
			Category:         "BCC Resources",
			ShortDescription: `Provide a resource to create an AutoSnapshotPolicy.`,
			Description:      ``,
			Keywords: []string{
				"bcc",
				"auto",
				"snapshot",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the automatic snapshot policy, which supports uppercase and lowercase letters, numbers, Chinese and special characters, such as "-","_","/",",".", and the value must start with a letter, length 1-65.`,
				},
				resource.Attribute{
					Name:        "repeat_weekdays",
					Description: `(Required) Repeat time of the automatic snapshot policy, supporting in range of [0, 6]`,
				},
				resource.Attribute{
					Name:        "retention_days",
					Description: `(Required) Number of days to retain the automatic snapshot, and -1 means permanently retained.`,
				},
				resource.Attribute{
					Name:        "time_points",
					Description: `(Required) Time point of generate snapshot in a day, the minimum unit is hour, supporting in range of [0, 23]`,
				},
				resource.Attribute{
					Name:        "volume_ids",
					Description: `(Optional) Volume id list to be attached of the automatic snapshot policy, these CDS volumes must be in-used. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "created_time",
					Description: `Creation time of the automatic snapshot policy.`,
				},
				resource.Attribute{
					Name:        "deleted_time",
					Description: `Deletion time of the automatic snapshot policy.`,
				},
				resource.Attribute{
					Name:        "last_execute_time",
					Description: `Last execution time of the automatic snapshot policy.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the automatic snapshot policy.`,
				},
				resource.Attribute{
					Name:        "updated_time",
					Description: `Update time of the automatic snapshot policy.`,
				},
				resource.Attribute{
					Name:        "volume_count",
					Description: `The count of volumes associated with the snapshot. ## Import AutoSnapshotPolicy can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import baiducloud_auto_snapshot_policy.default id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_time",
					Description: `Creation time of the automatic snapshot policy.`,
				},
				resource.Attribute{
					Name:        "deleted_time",
					Description: `Deletion time of the automatic snapshot policy.`,
				},
				resource.Attribute{
					Name:        "last_execute_time",
					Description: `Last execution time of the automatic snapshot policy.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the automatic snapshot policy.`,
				},
				resource.Attribute{
					Name:        "updated_time",
					Description: `Update time of the automatic snapshot policy.`,
				},
				resource.Attribute{
					Name:        "volume_count",
					Description: `The count of volumes associated with the snapshot. ## Import AutoSnapshotPolicy can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import baiducloud_auto_snapshot_policy.default id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "baiducloud_bos_bucket",
			Category:         "BOS Resources",
			ShortDescription: `Provide a resource to create a BOS Bucket.`,
			Description:      ``,
			Keywords: []string{
				"bos",
				"bucket",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required, ForceNew) Name of the bucket.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `(Optional) Canned ACL to apply, available values are private, public-read and public-read-write. Default to private.`,
				},
				resource.Attribute{
					Name:        "copyright_protection",
					Description: `(Optional) Configuration of the copyright protection.`,
				},
				resource.Attribute{
					Name:        "cors_rule",
					Description: `(Optional) Configuration of the Cross-Origin Resource Sharing. Up to 100 rules are allowed per bucket, if there are multiple configurations, the execution order is from top to bottom.`,
				},
				resource.Attribute{
					Name:        "force_destroy",
					Description: `(Optional) Whether to force delete the bucket and related objects when the bucket is not empty. Default to false.`,
				},
				resource.Attribute{
					Name:        "lifecycle_rule",
					Description: `(Optional) Configuration of object lifecycle management.`,
				},
				resource.Attribute{
					Name:        "logging",
					Description: `(Optional) Settings of the bucket logging.`,
				},
				resource.Attribute{
					Name:        "replication_configuration",
					Description: `(Optional) Replication configuration of the BOS bucket.`,
				},
				resource.Attribute{
					Name:        "server_side_encryption_rule",
					Description: `(Optional) Encryption rule for the server side, which can only be AES256 currently.`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `(Optional) Storage class of the BOS bucket, available values are STANDARD, STANDARD_IA, COLD or ARCHIVE.`,
				},
				resource.Attribute{
					Name:        "website",
					Description: `(Optional) Website of the BOS bucket. The ` + "`" + `copyright_protection` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "resource",
					Description: `(Required) The resources to be protected for copyright. The ` + "`" + `cors_rule` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "allowed_methods",
					Description: `(Required) Specifies which methods are allowed. Can be GET,PUT,DELETE,POST or HEAD.`,
				},
				resource.Attribute{
					Name:        "allowed_origins",
					Description: `(Required) Specifies which origins are allowed, containing up to one`,
				},
				resource.Attribute{
					Name:        "allowed_expose_headers",
					Description: `(Optional) Specifies which expose headers are allowed.`,
				},
				resource.Attribute{
					Name:        "allowed_headers",
					Description: `(Optional) Specifies which headers are allowed.`,
				},
				resource.Attribute{
					Name:        "max_age_seconds",
					Description: `(Optional) Specifies time in seconds that browser can cache the response for a preflight request. The ` + "`" + `lifecycle_rule` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) Action of the lifecycle rule.`,
				},
				resource.Attribute{
					Name:        "condition",
					Description: `(Required) Condition of the lifecycle rule, only the time form is supported currently.`,
				},
				resource.Attribute{
					Name:        "resource",
					Description: `(Required) Resource of the lifecycle rule. For example, samplebucket/prefix/`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Required) Status of the lifecycle rule, which can be enabled, disabled. The rule cannot take effect when the status is disabled.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the lifecycle rule. The id must be unique and cannot be repeated in the same bucket. The system will automatically generate one when it is not specified. The ` + "`" + `action` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Action name of the lifecycle rule, which can be Transition, DeleteObject and AbortMultipartUpload.`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `(Optional) When the action is Transition, it can be set to STANDARD_IA or COLD or ARCHIVE, indicating that it is changed from the original storage type to low frequency storage or cold storage or archive storage. The ` + "`" + `condition` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "time",
					Description: `(Required) The condition time, implemented by the date_greater_than. The ` + "`" + `time` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "date_greater_than",
					Description: `(Required) Support absolute time date and relative time days. The absolute time date format is yyyy-mm-ddThh:mm:ssZ,eg. 2019-09-07T00:00:00Z. The absolute time is UTC time, which must end at 00:00:00(UTC 0 point); the description of relative time days follows ISO8601, and the minimum time granularity supported is days, such as: $(lastModified)+P7D indicates the time of object 7 days after last-modified. The ` + "`" + `logging` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "target_bucket",
					Description: `(Required) Target bucket name that will receive the log data.`,
				},
				resource.Attribute{
					Name:        "target_prefix",
					Description: `(Optional) Target prefix for the log data. The ` + "`" + `replication_configuration` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `(Required) Destination of the replication configuration.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) ID of the replication configuration.`,
				},
				resource.Attribute{
					Name:        "replicate_deletes",
					Description: `(Required) Whether to enable the delete synchronization, which can be enabled, disabled.`,
				},
				resource.Attribute{
					Name:        "resource",
					Description: `(Required) Resource of the replication configuration. The configuration format of the resource is {$bucket_name/<effective object prefix>}, which must start with "$bucket_name"+"/"`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Required) Status of the replication configuration. Valid values are enabled and disabled.`,
				},
				resource.Attribute{
					Name:        "replicate_history",
					Description: `(Optional) Configuration of the replicate history. The bucket name in replicate history needs to be the same as the bucket name in the destination above. After the history file is copied, all the objects of the inventory are copied to the destination bucket synchronously. The history file copy range is not referenced to the resource. The ` + "`" + `destination` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) Destination bucket name of the replication configuration.`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `(Optional) Destination storage class of the replication configuration, the parameter does not need to be configured if it is consistent with the storage class of the source bucket, if you need to specify the storage class separately, it can be COLD, STANDARD, STANDARD_IA. The ` + "`" + `replicate_history` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) Destination bucket name of the replication configuration.`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `(Optional) Destination storage class of the replication configuration, the parameter does not need to be configured if it is consistent with the storage class of the source bucket, if you need to specify the storage class separately, it can be COLD, STANDARD, STANDARD_IA. The ` + "`" + `website` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "error_document",
					Description: `(Optional) An absolute path to the document to return in case of a 404 error.`,
				},
				resource.Attribute{
					Name:        "index_document",
					Description: `(Optional) Baiducloud BOS returns this index document when requests are made to the root domain or any of the subfolders. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `Creation date of the BOS bucket.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Location of the BOS bucket.`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `Owner ID of the BOS bucket.`,
				},
				resource.Attribute{
					Name:        "owner_name",
					Description: `Owner name of the BOS bucket. ## Import BOS bucket can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import baiducloud_bos_bucket.default bucket_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "creation_date",
					Description: `Creation date of the BOS bucket.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Location of the BOS bucket.`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `Owner ID of the BOS bucket.`,
				},
				resource.Attribute{
					Name:        "owner_name",
					Description: `Owner name of the BOS bucket. ## Import BOS bucket can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import baiducloud_bos_bucket.default bucket_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "baiducloud_bos_bucket_object",
			Category:         "BOS Resources",
			ShortDescription: `Provide a resource to create a BOS bucket object.`,
			Description:      ``,
			Keywords: []string{
				"bos",
				"bucket",
				"object",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required, ForceNew) Name of the bucket to put the file in.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required, ForceNew) Name of the object once it is in the bucket.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `(Optional) Canned ACL of the object, which can be private or public-read. If the value is not set, the object permission will be empty by default, and then the bucket permission as default.`,
				},
				resource.Attribute{
					Name:        "cache_control",
					Description: `(Optional) The caching behavior along the request/reply chain. Valid values are private, no-cache, max-age and must-revalidate. If not set, the value is empty.`,
				},
				resource.Attribute{
					Name:        "content_crc32",
					Description: `(Optional) Crc(cyclic redundancy check code) value of the object.`,
				},
				resource.Attribute{
					Name:        "content_disposition",
					Description: `(Optional) Specifies presentational information for the object, which can be inline or attachment. If not set, the value is empty.`,
				},
				resource.Attribute{
					Name:        "content_length",
					Description: `(Optional) Length of the content to be uploaded.`,
				},
				resource.Attribute{
					Name:        "content_md5",
					Description: `(Optional) MD5 digest of the HTTP request content defined in RFC2616 can be carried by the field to verify whether the file saved on the BOS side is consistent with the file expected by the user.`,
				},
				resource.Attribute{
					Name:        "content_sha256",
					Description: `(Optional) Sha256 value of the object, which is used to verify whether the file saved on the BOS side is consistent with the file expected by the user, the sha256 has higher verification accuracy, and the sha256 value of the transmitted data must match this, otherwise the object uploaded fails.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `(Optional) Type to describe the format of the object data.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Optional, ForceNew) The literal string value that will be uploaded as the object content.`,
				},
				resource.Attribute{
					Name:        "expires",
					Description: `(Optional) The expire date is used to set the cache expiration time when downloading object. If it is not set, the BOS will set the cache expiration time to three days by default.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional, ForceNew) The file path that will be read and uploaded as raw bytes for the object content.`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `(Optional) Storage class of the object, which can be COLD, STANDARD_IA, STANDARD or ARCHIVE. Default to STANDARD.`,
				},
				resource.Attribute{
					Name:        "user_meta",
					Description: `(Optional) The mapping of key/values to to provision metadata, which will be automatically prefixed by x-bce-meta-. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "content_encoding",
					Description: `Encoding of the object.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `Etag generated of the object.`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `Last modified date of the object.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "content_encoding",
					Description: `Encoding of the object.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `Etag generated of the object.`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `Last modified date of the object.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "baiducloud_cds",
			Category:         "BCC Resources",
			ShortDescription: `Provide a resource to create a CDS.`,
			Description:      ``,
			Keywords: []string{
				"bcc",
				"cds",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "payment_timing",
					Description: `(Required) payment method, support Prepaid or Postpaid`,
				},
				resource.Attribute{
					Name:        "auto_snapshot",
					Description: `(Optional) Delete relate auto snapshot when release this cds volume`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) CDS volume description`,
				},
				resource.Attribute{
					Name:        "disk_size_in_gb",
					Description: `(Optional) CDS disk size, support between 5 and 32765, if snapshot_id not set, this parameter is required.`,
				},
				resource.Attribute{
					Name:        "manual_snapshot",
					Description: `(Optional) Delete relate snapshot when release this cds volume`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) CDS volume name`,
				},
				resource.Attribute{
					Name:        "reservation_length",
					Description: `(Optional) Prepaid reservation length, support [1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36], only useful when payment_timing is Prepaid`,
				},
				resource.Attribute{
					Name:        "reservation_time_unit",
					Description: `(Optional) Prepaid reservation time unit, only support Month now`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional, ForceNew) Snapshot id, support create cds use snapshot, when set this parameter, cds_disk_size is ignored`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `(Optional) CDS dist storage type, support hp1, std1, cloud_hp1 and hdd, default hp1, see https://cloud.baidu.com/doc/BCC/s/6jwvyo0q2/#storagetype for detail`,
				},
				resource.Attribute{
					Name:        "zone_name",
					Description: `(Optional) Zone name ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "auto_snapshot_policy_id",
					Description: `CDS bind Auto Snapshot policy id`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `CDS volume create time`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `CDS volume expire time`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `CDS volume status`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `CDS volume type ## Import CDS can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import baiducloud_cds.default id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "auto_snapshot_policy_id",
					Description: `CDS bind Auto Snapshot policy id`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `CDS volume create time`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `CDS volume expire time`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `CDS volume status`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `CDS volume type ## Import CDS can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import baiducloud_cds.default id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "baiducloud_cds_attachment",
			Category:         "BCC Resources",
			ShortDescription: `Provide a resource to create a CDS attachment, can attach a CDS volume with instance.`,
			Description:      ``,
			Keywords: []string{
				"bcc",
				"cds",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cds_id",
					Description: `(Required, ForceNew) CDS volume ID`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, ForceNew) The ID of Instance which will attach CDS volume ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "attachment_device",
					Description: `CDS mount device path`,
				},
				resource.Attribute{
					Name:        "attachment_serial",
					Description: `CDS serial ## Import CDS attachment can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import baiducloud_cds_attachment.default id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "attachment_device",
					Description: `CDS mount device path`,
				},
				resource.Attribute{
					Name:        "attachment_serial",
					Description: `CDS serial ## Import CDS attachment can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import baiducloud_cds_attachment.default id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "baiducloud_cert",
			Category:         "CERT Resources",
			ShortDescription: `Provide a resource to Upload a cert.`,
			Description:      ``,
			Keywords: []string{
				"cert",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cert_name",
					Description: `(Required) Cert Name`,
				},
				resource.Attribute{
					Name:        "cert_private_data",
					Description: `(Required) Cert private key data, base64 encode`,
				},
				resource.Attribute{
					Name:        "cert_server_data",
					Description: `(Required) Server Cert data, base64 encode`,
				},
				resource.Attribute{
					Name:        "cert_link_data",
					Description: `(Optional) Cert lint data, base64 encode`,
				},
				resource.Attribute{
					Name:        "cert_type",
					Description: `(Optional) Cert type ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cert_common_name",
					Description: `Cert common name`,
				},
				resource.Attribute{
					Name:        "cert_create_time",
					Description: `Cert create time`,
				},
				resource.Attribute{
					Name:        "cert_start_time",
					Description: `Cert start time`,
				},
				resource.Attribute{
					Name:        "cert_stop_time",
					Description: `Cert stop time`,
				},
				resource.Attribute{
					Name:        "cert_update_time",
					Description: `Cert update time`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cert_common_name",
					Description: `Cert common name`,
				},
				resource.Attribute{
					Name:        "cert_create_time",
					Description: `Cert create time`,
				},
				resource.Attribute{
					Name:        "cert_start_time",
					Description: `Cert start time`,
				},
				resource.Attribute{
					Name:        "cert_stop_time",
					Description: `Cert stop time`,
				},
				resource.Attribute{
					Name:        "cert_update_time",
					Description: `Cert update time`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "baiducloud_cfc_alias",
			Category:         "CFC Resources",
			ShortDescription: `Provide a resource to create a CFC Function Alias.`,
			Description:      ``,
			Keywords: []string{
				"cfc",
				"alias",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "alias_name",
					Description: `(Required, ForceNew) CFC Function alias name`,
				},
				resource.Attribute{
					Name:        "function_name",
					Description: `(Required, ForceNew) CFC Function name`,
				},
				resource.Attribute{
					Name:        "function_version",
					Description: `(Required) CFC Function version this alias binded`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) CFC Function alias description ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "alias_arn",
					Description: `CFC Function alias arn`,
				},
				resource.Attribute{
					Name:        "alias_brn",
					Description: `CFC Function alias brn`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `CFC Function alias create time`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `CFC Function uid`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `CFC Function alias update time`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "alias_arn",
					Description: `CFC Function alias arn`,
				},
				resource.Attribute{
					Name:        "alias_brn",
					Description: `CFC Function alias brn`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `CFC Function alias create time`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `CFC Function uid`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `CFC Function alias update time`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "baiducloud_cfc_function",
			Category:         "CFC Resources",
			ShortDescription: `Provide a resource to create an CFC Function.`,
			Description:      ``,
			Keywords: []string{
				"cfc",
				"function",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "function_name",
					Description: `(Required, ForceNew) CFC function name, length must be between 1 and 64 bytes`,
				},
				resource.Attribute{
					Name:        "handler",
					Description: `(Required) CFC Function execution handler`,
				},
				resource.Attribute{
					Name:        "runtime",
					Description: `(Required) CFC Function runtime`,
				},
				resource.Attribute{
					Name:        "time_out",
					Description: `(Required) Function time out, support [1, 300]s`,
				},
				resource.Attribute{
					Name:        "code_bos_bucket",
					Description: `(Optional) CFC Function Code storage bos bucket name`,
				},
				resource.Attribute{
					Name:        "code_bos_object",
					Description: `(Optional) CFC Function Code storage bos object key`,
				},
				resource.Attribute{
					Name:        "code_file_dir",
					Description: `(Optional) CFC Function Code local file dir`,
				},
				resource.Attribute{
					Name:        "code_file_name",
					Description: `(Optional) CFC Function Code local zip file name`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Function description`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `(Optional) CFC Function environment variables`,
				},
				resource.Attribute{
					Name:        "log_bos_dir",
					Description: `(Optional) Log save dir if log type is bos`,
				},
				resource.Attribute{
					Name:        "log_type",
					Description: `(Optional) Log save type, support bos/none`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `(Optional) CFC Function memory size, should be an integer multiple of 128`,
				},
				resource.Attribute{
					Name:        "reserved_concurrent_executions",
					Description: `(Optional) Function reserved concurrent executions, support [0-90] ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "code_id",
					Description: `CFC Function code id`,
				},
				resource.Attribute{
					Name:        "code_sha256",
					Description: `Function code sha256`,
				},
				resource.Attribute{
					Name:        "code_size",
					Description: `Function code size`,
				},
				resource.Attribute{
					Name:        "code_storage",
					Description: `CFC Code storage information`,
				},
				resource.Attribute{
					Name:        "commit_id",
					Description: `Function commit id`,
				},
				resource.Attribute{
					Name:        "function_arn",
					Description: `The same as function brn`,
				},
				resource.Attribute{
					Name:        "function_brn",
					Description: `Function brn`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `The same as update_time`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Function region`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `Function exec role`,
				},
				resource.Attribute{
					Name:        "source_tag",
					Description: `CFC Function source tag`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `Function user uid`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Last update time`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Function version, should only be $LATEST ## Import CFC can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import baiducloud_cfc_function.default functionName ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "code_id",
					Description: `CFC Function code id`,
				},
				resource.Attribute{
					Name:        "code_sha256",
					Description: `Function code sha256`,
				},
				resource.Attribute{
					Name:        "code_size",
					Description: `Function code size`,
				},
				resource.Attribute{
					Name:        "code_storage",
					Description: `CFC Code storage information`,
				},
				resource.Attribute{
					Name:        "commit_id",
					Description: `Function commit id`,
				},
				resource.Attribute{
					Name:        "function_arn",
					Description: `The same as function brn`,
				},
				resource.Attribute{
					Name:        "function_brn",
					Description: `Function brn`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `The same as update_time`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Function region`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `Function exec role`,
				},
				resource.Attribute{
					Name:        "source_tag",
					Description: `CFC Function source tag`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `Function user uid`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Last update time`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Function version, should only be $LATEST ## Import CFC can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import baiducloud_cfc_function.default functionName ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "baiducloud_cfc_trigger",
			Category:         "CFC Resources",
			ShortDescription: `Provide a resource to create a CFC Function Trigger.`,
			Description:      ``,
			Keywords: []string{
				"cfc",
				"trigger",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required, ForceNew) CFC Funtion Trigger source type, support bos/http/crontab/dueros/duedge/cdn`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Required, ForceNew) CFC Function Trigger target, it should be function brn`,
				},
				resource.Attribute{
					Name:        "auth_type",
					Description: `(Optional) CFC Function Trigger auth type if source_type is http, support anonymous or iam`,
				},
				resource.Attribute{
					Name:        "bos_event_type",
					Description: `(Optional) CFC Function Trigger bos event type`,
				},
				resource.Attribute{
					Name:        "bucket",
					Description: `(Optional, ForceNew) CFC Function Trigger source bucket if source_type is bos`,
				},
				resource.Attribute{
					Name:        "cdn_event_type",
					Description: `(Optional) CFC Function Trigger cdn event type`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional) CFC Function Trigger domain if source_type is cdn`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) CFC Function Trigger enabled if source_type is crontab`,
				},
				resource.Attribute{
					Name:        "input",
					Description: `(Optional) CFC Function Trigger input if source_type is crontab`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Optional) CFC Function Trigger method if source_type is http`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) CFC Function Trigger name if source_type is bos or crontab`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `(Optional) CFC Function Trigger remark if source_type is cdn`,
				},
				resource.Attribute{
					Name:        "resource_path",
					Description: `(Optional) CFC Function Trigger resource path if source_type is http`,
				},
				resource.Attribute{
					Name:        "resource",
					Description: `(Optional) CFC Function Trigger resource if source_type is bos`,
				},
				resource.Attribute{
					Name:        "schedule_expression",
					Description: `(Optional) CFC Function Trigger schedule expression if source_type is crontab`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) CFC Funtion Trigger status if source_type is bos or cdn ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "relation_id",
					Description: `CFC Function Trigger relation id`,
				},
				resource.Attribute{
					Name:        "sid",
					Description: `CFC Funtion Trigger sid`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "relation_id",
					Description: `CFC Function Trigger relation id`,
				},
				resource.Attribute{
					Name:        "sid",
					Description: `CFC Funtion Trigger sid`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "baiducloud_cfc_version",
			Category:         "CFC Resources",
			ShortDescription: `Provide a resource to publish a CFC Function Version.`,
			Description:      ``,
			Keywords: []string{
				"cfc",
				"version",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "function_name",
					Description: `(Required, ForceNew) CFC function name, length must be between 1 and 64 bytes`,
				},
				resource.Attribute{
					Name:        "code_sha256",
					Description: `(Optional) Function code sha256`,
				},
				resource.Attribute{
					Name:        "log_bos_dir",
					Description: `(Optional) Log save dir if log type is bos`,
				},
				resource.Attribute{
					Name:        "log_type",
					Description: `(Optional) Log save type, support bos/none`,
				},
				resource.Attribute{
					Name:        "version_description",
					Description: `(Optional, ForceNew) Function version description ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "code_size",
					Description: `Function code size`,
				},
				resource.Attribute{
					Name:        "commit_id",
					Description: `Function commit id`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Function description`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `CFC Function environment variables`,
				},
				resource.Attribute{
					Name:        "function_arn",
					Description: `The same as function brn`,
				},
				resource.Attribute{
					Name:        "function_brn",
					Description: `Function brn`,
				},
				resource.Attribute{
					Name:        "handler",
					Description: `CFC Function execution handler`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `The same as update_time`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `CFC Function memory size`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `CFC Function bined VPC Subnet id list`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `Function exec role`,
				},
				resource.Attribute{
					Name:        "runtime",
					Description: `CFC Function runtime`,
				},
				resource.Attribute{
					Name:        "time_out",
					Description: `Function time out`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `Function user uid`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Last update time`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Function version`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "code_size",
					Description: `Function code size`,
				},
				resource.Attribute{
					Name:        "commit_id",
					Description: `Function commit id`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Function description`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `CFC Function environment variables`,
				},
				resource.Attribute{
					Name:        "function_arn",
					Description: `The same as function brn`,
				},
				resource.Attribute{
					Name:        "function_brn",
					Description: `Function brn`,
				},
				resource.Attribute{
					Name:        "handler",
					Description: `CFC Function execution handler`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `The same as update_time`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `CFC Function memory size`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `CFC Function bined VPC Subnet id list`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `Function exec role`,
				},
				resource.Attribute{
					Name:        "runtime",
					Description: `CFC Function runtime`,
				},
				resource.Attribute{
					Name:        "time_out",
					Description: `Function time out`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `Function user uid`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Last update time`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Function version`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "baiducloud_eip",
			Category:         "EIP Resources",
			ShortDescription: `Provide a resource to create an EIP.`,
			Description:      ``,
			Keywords: []string{
				"eip",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bandwidth_in_mbps",
					Description: `(Required) Eip bandwidth(Mbps), if payment_timing is Prepaid or billing_method is ByBandWidth, support between 1 and 200, if billing_method is ByTraffic, support between 1 and 1000`,
				},
				resource.Attribute{
					Name:        "billing_method",
					Description: `(Required, ForceNew) Eip billing method, support ByTraffic or ByBandwidth`,
				},
				resource.Attribute{
					Name:        "payment_timing",
					Description: `(Required, ForceNew) Eip payment timing, support Prepaid and Postpaid`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, ForceNew) Eip name, length must be between 1 and 65 bytes`,
				},
				resource.Attribute{
					Name:        "reservation_length",
					Description: `(Optional) Eip Prepaid billing reservation length, only useful when payment_timing is Prepaid`,
				},
				resource.Attribute{
					Name:        "reservation_time_unit",
					Description: `(Optional) Eip Prepaid billing reservation time unit, only useful when payment_timing is Prepaid`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, ForceNew) Tags, do not support modify ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Eip create time`,
				},
				resource.Attribute{
					Name:        "eip_instance_type",
					Description: `Eip instance type`,
				},
				resource.Attribute{
					Name:        "eip",
					Description: `Eip address`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `Eip expire time`,
				},
				resource.Attribute{
					Name:        "share_group_id",
					Description: `Eip share group id`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Eip status ## Import EIP can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import baiducloud_eip.default eip ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "create_time",
					Description: `Eip create time`,
				},
				resource.Attribute{
					Name:        "eip_instance_type",
					Description: `Eip instance type`,
				},
				resource.Attribute{
					Name:        "eip",
					Description: `Eip address`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `Eip expire time`,
				},
				resource.Attribute{
					Name:        "share_group_id",
					Description: `Eip share group id`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Eip status ## Import EIP can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import baiducloud_eip.default eip ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "baiducloud_eip_association",
			Category:         "EIP Resources",
			ShortDescription: `Provide a resource to create an EIP association, bind an EIP with instance.`,
			Description:      ``,
			Keywords: []string{
				"eip",
				"association",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "eip",
					Description: `(Required, ForceNew) EIP which need to associate with instance`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, ForceNew) Instance ID which need to associate with EIP`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Required, ForceNew) Instance type which need to associate with EIP, support BCC/BLB/NAT/VPN ## Import EIP association can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import baiducloud_eip_association.default eip ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "baiducloud_instance",
			Category:         "BCC Resources",
			ShortDescription: `Use this resource to get information about a BCC instance.`,
			Description:      ``,
			Keywords: []string{
				"bcc",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "billing",
					Description: `(Required) Billing information of the instance.`,
				},
				resource.Attribute{
					Name:        "cpu_count",
					Description: `(Required) Number of CPU cores to be created for the instance.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Required) ID of the image to be used for the instance.`,
				},
				resource.Attribute{
					Name:        "memory_capacity_in_gb",
					Description: `(Required) Memory capacity(GB) of the instance to be created.`,
				},
				resource.Attribute{
					Name:        "admin_pass",
					Description: `(Optional) Password of the instance to be started. This value should be 8-16 characters, and English, numbers and symbols must exist at the same time. The symbols is limited to "!@#$%^`,
				},
				resource.Attribute{
					Name:        "auto_renew_time_length",
					Description: `(Optional, ForceNew) The time length of automatic renewal. It is valid when payment_timing is Prepaid, and the value should be 1-9 when the auto_renew_time_unit is month and 1-3 when the auto_renew_time_unit is year. Default to 1.`,
				},
				resource.Attribute{
					Name:        "auto_renew_time_unit",
					Description: `(Optional, ForceNew) Time unit of automatic renewal, the value can be month or year. The default value is empty, indicating no automatic renewal. It is valid only when the payment_timing is Prepaid.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional, ForceNew) Availability zone to start the instance in.`,
				},
				resource.Attribute{
					Name:        "card_count",
					Description: `(Optional, ForceNew) Count of the GPU cards or FPGA cards to be carried for the instance to be created, it is valid only when the gpu_card or fpga_card field is not empty.`,
				},
				resource.Attribute{
					Name:        "cds_auto_renew",
					Description: `(Optional, ForceNew) Whether the cds is automatically renewed. It is valid when payment_timing is Prepaid. Default to false.`,
				},
				resource.Attribute{
					Name:        "cds_disks",
					Description: `(Optional) CDS disks of the instance.`,
				},
				resource.Attribute{
					Name:        "dedicate_host_id",
					Description: `(Optional, ForceNew) The ID of dedicated host.`,
				},
				resource.Attribute{
					Name:        "delete_cds_snapshot_flag",
					Description: `(Optional, ForceNew) Whether to release the cds disk snapshots, default to false. It is effective only when the related_release_flag is true.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the instance.`,
				},
				resource.Attribute{
					Name:        "ephemeral_disks",
					Description: `(Optional) Ephemeral disks of the instance.`,
				},
				resource.Attribute{
					Name:        "fpga_card",
					Description: `(Optional, ForceNew) FPGA card of the instance.`,
				},
				resource.Attribute{
					Name:        "gpu_card",
					Description: `(Optional, ForceNew) GPU card of the instance.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Optional, ForceNew) Type of the instance to start. Available values are N1, N2, N3, C1, C2, S1, G1, F1. Default to N3.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the instance. Support for uppercase and lowercase letters, numbers, Chinese and special characters, such as "-","_","/",".", the value must start with a letter, length 1-65.`,
				},
				resource.Attribute{
					Name:        "related_release_flag",
					Description: `(Optional, ForceNew) Whether to release the eip and data disks mounted by the current instance. Can only be released uniformly or not. Default to false.`,
				},
				resource.Attribute{
					Name:        "relation_tag",
					Description: `(Optional, ForceNew) The new instance associated with existing Tags or not, default false. The Tags should already exit if set true`,
				},
				resource.Attribute{
					Name:        "root_disk_size_in_gb",
					Description: `(Optional, ForceNew) System disk size(GB) of the instance to be created. The value range is [40,500]GB, Default to 40GB, and more than 40GB is charged according to the cloud disk price. Note that the specified system disk size needs to meet the minimum disk space limit of the mirror used.`,
				},
				resource.Attribute{
					Name:        "root_disk_storage_type",
					Description: `(Optional, ForceNew) System disk storage type of the instance. Available values are std1, hp1, cloud_hp1, local, sata, ssd. Default to cloud_hp1.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Optional) Security groups of the instance.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) The subnet ID of VPC. The default subnet will be used when it is empty. The instance will restart after changing the subnet.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, ForceNew) Tags, do not support modify The ` + "`" + `billing` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "payment_timing",
					Description: `(Required) Payment timing of billing, which can be Prepaid or Postpaid. The default is Postpaid.`,
				},
				resource.Attribute{
					Name:        "reservation",
					Description: `(Optional) Reservation of the instance. The ` + "`" + `reservation` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "reservation_length",
					Description: `(Required) The reservation length that you will pay for your resource. It is valid when payment_timing is Prepaid. Valid values: [1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36].`,
				},
				resource.Attribute{
					Name:        "reservation_time_unit",
					Description: `(Required) The reservation time unit that you will pay for your resource. It is valid when payment_timing is Prepaid. The value can only be month currently, which is also the default value. The ` + "`" + `cds_disks` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "cds_size_in_gb",
					Description: `(Optional, ForceNew) The size(GB) of CDS.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional, ForceNew) Snapshot ID of CDS.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `(Optional, ForceNew) Storage type of the CDS. The ` + "`" + `ephemeral_disks` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "size_in_gb",
					Description: `(Optional, ForceNew) Size(GB) of the ephemeral disk.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `(Optional, ForceNew) Storage type of the ephemeral disk. Available values are std1, hp1, cloud_hp1, local, sata, ssd. Default to cloud_hp1. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "auto_renew",
					Description: `Whether to automatically renew.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the instance.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `Expire time of the instance.`,
				},
				resource.Attribute{
					Name:        "internal_ip",
					Description: `Internal IP assigned to the instance.`,
				},
				resource.Attribute{
					Name:        "keypair_id",
					Description: `Key pair id of the instance.`,
				},
				resource.Attribute{
					Name:        "keypair_name",
					Description: `Key pair name of the instance.`,
				},
				resource.Attribute{
					Name:        "network_capacity_in_mbps",
					Description: `Public network bandwidth(Mbps) of the instance.`,
				},
				resource.Attribute{
					Name:        "placement_policy",
					Description: `The placement policy of the instance, which can be default or dedicatedHost.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the instance.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC ID of the instance. ## Import BCC instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import baiducloud_instance.my-server id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "auto_renew",
					Description: `Whether to automatically renew.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the instance.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `Expire time of the instance.`,
				},
				resource.Attribute{
					Name:        "internal_ip",
					Description: `Internal IP assigned to the instance.`,
				},
				resource.Attribute{
					Name:        "keypair_id",
					Description: `Key pair id of the instance.`,
				},
				resource.Attribute{
					Name:        "keypair_name",
					Description: `Key pair name of the instance.`,
				},
				resource.Attribute{
					Name:        "network_capacity_in_mbps",
					Description: `Public network bandwidth(Mbps) of the instance.`,
				},
				resource.Attribute{
					Name:        "placement_policy",
					Description: `The placement policy of the instance, which can be default or dedicatedHost.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the instance.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC ID of the instance. ## Import BCC instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import baiducloud_instance.my-server id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "baiducloud_nat_gateway",
			Category:         "VPC Resources",
			ShortDescription: `Provide a resource to create a NAT Gateway.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"nat",
				"gateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "billing",
					Description: `(Required) Billing information of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the NAT gateway, consisting of uppercase and lowercase letters、numbers and special characters, such as "-","_","/",".". The value must start with a letter, and the length should between 1-65.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, ForceNew) VPC ID of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `(Optional, ForceNew) Specification of the NAT gateway, available values are small(supports up to 5 public IPs), medium(up to 10 public IPs) and large(up to 15 public IPs). Default to small. The ` + "`" + `billing` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "payment_timing",
					Description: `(Required, ForceNew) Payment timing of the billing, which can be Prepaid or Postpaid. The default is Postpaid.`,
				},
				resource.Attribute{
					Name:        "reservation",
					Description: `(Optional) Reservation of the NAT gateway. The ` + "`" + `reservation` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "reservation_length",
					Description: `(Optional, ForceNew) Reservation length that you will pay for your resource. It is valid when payment_timing is Prepaid. Valid values: [1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36].`,
				},
				resource.Attribute{
					Name:        "reservation_time_unit",
					Description: `(Optional) Reservation time unit that you will pay for your resource. It is valid when payment_timing is Prepaid. The value can only be month currently, which is also the default value. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "eips",
					Description: `One public network EIP associated with the NAT gateway or one or more EIPs in the shared bandwidth.`,
				},
				resource.Attribute{
					Name:        "expired_time",
					Description: `Expired time of the NAT gateway, which will be empty when the payment_timing is Postpaid.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the NAT gateway. ## Import NAT Gateway instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import baiducloud_nat_gateway.default nat_gateway_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "eips",
					Description: `One public network EIP associated with the NAT gateway or one or more EIPs in the shared bandwidth.`,
				},
				resource.Attribute{
					Name:        "expired_time",
					Description: `Expired time of the NAT gateway, which will be empty when the payment_timing is Postpaid.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the NAT gateway. ## Import NAT Gateway instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import baiducloud_nat_gateway.default nat_gateway_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "baiducloud_peer_conn",
			Category:         "VPC Resources",
			ShortDescription: `Provide a resource to create a Peer Conn.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"peer",
				"conn",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bandwidth_in_mbps",
					Description: `(Required) Bandwidth(Mbps) of the peer connection.`,
				},
				resource.Attribute{
					Name:        "billing",
					Description: `(Required) Billing information of the peer connection.`,
				},
				resource.Attribute{
					Name:        "local_vpc_id",
					Description: `(Required, ForceNew) Local VPC ID of the peer connection.`,
				},
				resource.Attribute{
					Name:        "peer_region",
					Description: `(Required, ForceNew) Peer region of the peer connection.`,
				},
				resource.Attribute{
					Name:        "peer_vpc_id",
					Description: `(Required, ForceNew) Peer VPC ID of the peer connection.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the peer connection.`,
				},
				resource.Attribute{
					Name:        "dns_sync",
					Description: `(Optional) Whether to open the switch of dns synchronization.`,
				},
				resource.Attribute{
					Name:        "local_if_name",
					Description: `(Optional) Local interface name of the peer connection.`,
				},
				resource.Attribute{
					Name:        "peer_account_id",
					Description: `(Optional, ForceNew) Peer account ID of the peer VPC, which is required only when creating a peer connection across accounts.`,
				},
				resource.Attribute{
					Name:        "peer_if_name",
					Description: `(Optional) Peer interface name of the peer connection, which is allowed to be set only when the peer connection within this account. The ` + "`" + `billing` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "payment_timing",
					Description: `(Required, ForceNew) Payment timing of the billing, which can be Prepaid or Postpaid. The default is Postpaid.`,
				},
				resource.Attribute{
					Name:        "reservation",
					Description: `(Optional) Reservation of the peer connection. The ` + "`" + `reservation` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "reservation_length",
					Description: `(Optional, ForceNew) Reservation length that you will pay for your resource. It is valid when payment_timing is Prepaid. Valid values: [1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36].`,
				},
				resource.Attribute{
					Name:        "reservation_time_unit",
					Description: `(Optional) Reservation time unit that you will pay for your resource. It is valid when payment_timing is Prepaid. The value can only be month currently, which is also the default value. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "created_time",
					Description: `Created time of the peer connection.`,
				},
				resource.Attribute{
					Name:        "dns_status",
					Description: `DNS status of the peer connection.`,
				},
				resource.Attribute{
					Name:        "expired_time",
					Description: `Expired time of the peer connection, which will be empty when the payment_timing is Postpaid.`,
				},
				resource.Attribute{
					Name:        "local_if_id",
					Description: `Local interface ID of the peer connection.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `Role of the peer connection, which can be initiator or acceptor.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the peer connection. ## Import Peer Conn instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import baiducloud_peer_conn.default peer_conn_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_time",
					Description: `Created time of the peer connection.`,
				},
				resource.Attribute{
					Name:        "dns_status",
					Description: `DNS status of the peer connection.`,
				},
				resource.Attribute{
					Name:        "expired_time",
					Description: `Expired time of the peer connection, which will be empty when the payment_timing is Postpaid.`,
				},
				resource.Attribute{
					Name:        "local_if_id",
					Description: `Local interface ID of the peer connection.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `Role of the peer connection, which can be initiator or acceptor.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the peer connection. ## Import Peer Conn instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import baiducloud_peer_conn.default peer_conn_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "baiducloud_peer_conn_acceptor",
			Category:         "VPC Resources",
			ShortDescription: `Provide a resource to create a Peer Conn Acceptor.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"peer",
				"conn",
				"acceptor",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "peer_conn_id",
					Description: `(Required, ForceNew) ID of the peer connection.`,
				},
				resource.Attribute{
					Name:        "auto_accept",
					Description: `(Optional) Whether to accept the peer connection request, default to false.`,
				},
				resource.Attribute{
					Name:        "auto_reject",
					Description: `(Optional) Whether to reject the peer connection request, default to false.`,
				},
				resource.Attribute{
					Name:        "dns_sync",
					Description: `(Optional) Whether to open the switch of dns synchronization. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "bandwidth_in_mbps",
					Description: `Bandwidth(Mbps) of the peer connection.`,
				},
				resource.Attribute{
					Name:        "billing",
					Description: `Billing information of the peer connection.`,
				},
				resource.Attribute{
					Name:        "payment_timing",
					Description: `Payment timing of billing, which can be Prepaid or Postpaid. The default is Postpaid.`,
				},
				resource.Attribute{
					Name:        "created_time",
					Description: `Created time of the peer connection.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the peer connection.`,
				},
				resource.Attribute{
					Name:        "dns_status",
					Description: `DNS status of the peer connection.`,
				},
				resource.Attribute{
					Name:        "expired_time",
					Description: `Expired time of the peer connection, which will be empty when the payment_timing is Postpaid.`,
				},
				resource.Attribute{
					Name:        "local_if_id",
					Description: `Local interface ID of the peer connection.`,
				},
				resource.Attribute{
					Name:        "local_if_name",
					Description: `Local interface name of the peer connection.`,
				},
				resource.Attribute{
					Name:        "local_vpc_id",
					Description: `Local VPC ID of the peer connection.`,
				},
				resource.Attribute{
					Name:        "peer_account_id",
					Description: `Peer account ID of the peer VPC, which is required only when creating a peer connection across accounts.`,
				},
				resource.Attribute{
					Name:        "peer_if_name",
					Description: `Peer interface name of the peer connection, which is allowed to be set only when the peer connection within this account.`,
				},
				resource.Attribute{
					Name:        "peer_region",
					Description: `Peer region of the peer connection.`,
				},
				resource.Attribute{
					Name:        "peer_vpc_id",
					Description: `Peer VPC ID of the peer connection.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `Role of the peer connection, which can be initiator or acceptor.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the peer connection. ## Import Peer Conn Acceptor instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import baiducloud_peer_conn_acceptor.default peer_conn_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "bandwidth_in_mbps",
					Description: `Bandwidth(Mbps) of the peer connection.`,
				},
				resource.Attribute{
					Name:        "billing",
					Description: `Billing information of the peer connection.`,
				},
				resource.Attribute{
					Name:        "payment_timing",
					Description: `Payment timing of billing, which can be Prepaid or Postpaid. The default is Postpaid.`,
				},
				resource.Attribute{
					Name:        "created_time",
					Description: `Created time of the peer connection.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the peer connection.`,
				},
				resource.Attribute{
					Name:        "dns_status",
					Description: `DNS status of the peer connection.`,
				},
				resource.Attribute{
					Name:        "expired_time",
					Description: `Expired time of the peer connection, which will be empty when the payment_timing is Postpaid.`,
				},
				resource.Attribute{
					Name:        "local_if_id",
					Description: `Local interface ID of the peer connection.`,
				},
				resource.Attribute{
					Name:        "local_if_name",
					Description: `Local interface name of the peer connection.`,
				},
				resource.Attribute{
					Name:        "local_vpc_id",
					Description: `Local VPC ID of the peer connection.`,
				},
				resource.Attribute{
					Name:        "peer_account_id",
					Description: `Peer account ID of the peer VPC, which is required only when creating a peer connection across accounts.`,
				},
				resource.Attribute{
					Name:        "peer_if_name",
					Description: `Peer interface name of the peer connection, which is allowed to be set only when the peer connection within this account.`,
				},
				resource.Attribute{
					Name:        "peer_region",
					Description: `Peer region of the peer connection.`,
				},
				resource.Attribute{
					Name:        "peer_vpc_id",
					Description: `Peer VPC ID of the peer connection.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `Role of the peer connection, which can be initiator or acceptor.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the peer connection. ## Import Peer Conn Acceptor instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import baiducloud_peer_conn_acceptor.default peer_conn_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "baiducloud_route_rule",
			Category:         "VPC Resources",
			ShortDescription: `Provides a resource to create a VPC routing rule.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"route",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "destination_address",
					Description: `(Required, ForceNew) Destination CIDR block of the routing rule. The network segment can be 0.0.0.0/0, otherwise, the destination address cannot overlap with this VPC CIDR block(except when the destination network segment or the VPC CIDR is 0.0.0.0/0).`,
				},
				resource.Attribute{
					Name:        "next_hop_type",
					Description: `(Required, ForceNew) Type of the next hop, available values are custom、vpn and nat.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `(Required, ForceNew) ID of the routing table.`,
				},
				resource.Attribute{
					Name:        "source_address",
					Description: `(Required, ForceNew) Source CIDR block of the routing rule. The value can be all network segments 0.0.0.0/0, existing subnet segments in the VPC, or the network segment within the subnet.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, ForceNew) Description of the routing rule.`,
				},
				resource.Attribute{
					Name:        "next_hop_id",
					Description: `(Optional, ForceNew) ID of the next hop.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "baiducloud_security_group",
			Category:         "BCC Resources",
			ShortDescription: `Provide a resource to create a security group.`,
			Description:      ``,
			Keywords: []string{
				"bcc",
				"security",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) SecurityGroup name`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, ForceNew) SecurityGroup description`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, ForceNew) Tags, do not support modify`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional, ForceNew) SecurityGroup binded VPC id ## Import Bcc SecurityGroup can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import baiducloud_security_group.default security_group_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "baiducloud_security_group_rule",
			Category:         "BCC Resources",
			ShortDescription: `Provide a resource to create a security group rule.`,
			Description:      ``,
			Keywords: []string{
				"bcc",
				"security",
				"group",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "direction",
					Description: `(Required, ForceNew) SecurityGroup rule's direction, support ingress/egress`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Required, ForceNew) SecurityGroup rule's security group id`,
				},
				resource.Attribute{
					Name:        "dest_group_id",
					Description: `(Optional, ForceNew) SecurityGroup rule's destination group id, dest_group_id and dest_ip can not set in the same time`,
				},
				resource.Attribute{
					Name:        "dest_ip",
					Description: `(Optional, ForceNew) SecurityGroup rule's destination ip, dest_group_id and dest_ip can not set in the same time`,
				},
				resource.Attribute{
					Name:        "ether_type",
					Description: `(Optional, ForceNew) SecurityGroup rule's ether type, support IPv4/IPv6`,
				},
				resource.Attribute{
					Name:        "port_range",
					Description: `(Optional, ForceNew) SecurityGroup rule's port range, you can set single port like 80, or set a port range, like 1-65535, default 1-65535`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional, ForceNew) SecurityGroup rule's protocol, support tcp/udp/icmp/all, default all`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `(Optional, ForceNew) SecurityGroup rule's remark`,
				},
				resource.Attribute{
					Name:        "source_group_id",
					Description: `(Optional, ForceNew) SecurityGroup rule's source group id, source_group_id and source_ip can not set in the same time`,
				},
				resource.Attribute{
					Name:        "source_ip",
					Description: `(Optional, ForceNew) SecurityGroup rule's source ip, source_group_id and source_ip can not set in the same time`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "baiducloud_snapshot",
			Category:         "BCC Resources",
			ShortDescription: `Provide a resource to create a Snapshot.`,
			Description:      ``,
			Keywords: []string{
				"bcc",
				"snapshot",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Name of the snapshot, which supports uppercase and lowercase letters, numbers, Chinese and special characters, such as "-","_","/",",".", and the value must start with a letter, length 1-65.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `(Required, ForceNew) Volume id of the snapshot, this value will be nil if volume has been released.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, ForceNew) Description of the snapshot. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "create_method",
					Description: `Creation method of the snapshot.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the snapshot.`,
				},
				resource.Attribute{
					Name:        "size_in_gb",
					Description: `Size of the snapshot in GB.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the snapshot. ## Import Snapshot can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import baiducloud_snapshot.default id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "create_method",
					Description: `Creation method of the snapshot.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the snapshot.`,
				},
				resource.Attribute{
					Name:        "size_in_gb",
					Description: `Size of the snapshot in GB.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the snapshot. ## Import Snapshot can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import baiducloud_snapshot.default id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "baiducloud_subnet",
			Category:         "VPC Resources",
			ShortDescription: `Provide a resource to create a VPC subnet.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"subnet",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cidr",
					Description: `(Required, ForceNew) CIDR block of the subnet.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the subnet, which cannot take the value "default", the length is no more than 65 characters, and the value can be composed of numbers, characters and underscores.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, ForceNew) ID of the VPC.`,
				},
				resource.Attribute{
					Name:        "zone_name",
					Description: `(Required, ForceNew) The availability zone name within which the subnet should be created.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the subnet, and the value must be no more than 200 characters.`,
				},
				resource.Attribute{
					Name:        "subnet_type",
					Description: `(Optional, ForceNew) Type of the subnet, valid values are BCC, BCC_NAT and BBC. Default to BCC.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, ForceNew) Tags, do not support modify ## Import VPC subnet instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import baiducloud_subnet.default subnet_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "baiducloud_vpc",
			Category:         "VPC Resources",
			ShortDescription: `Provide a resource to create a VPC.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cidr",
					Description: `(Required, ForceNew) CIDR block for the VPC.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the VPC, which cannot take the value "default", the length is no more than 65 characters, and the value can be composed of numbers, characters and underscores.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the VPC. The value is no more than 200 characters.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, ForceNew) Tags, do not support modify ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `Route table ID created by default on VPC creation.`,
				},
				resource.Attribute{
					Name:        "secondary_cidrs",
					Description: `Secondary cidr list of the VPC. They will not be repeated. ## Import VPC instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import baiducloud_vpc.default vpc_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "route_table_id",
					Description: `Route table ID created by default on VPC creation.`,
				},
				resource.Attribute{
					Name:        "secondary_cidrs",
					Description: `Secondary cidr list of the VPC. They will not be repeated. ## Import VPC instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import baiducloud_vpc.default vpc_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"baiducloud_acl":                  0,
		"baiducloud_appblb":               1,
		"baiducloud_appblb_listener":      2,
		"baiducloud_appblb_server_group":  3,
		"baiducloud_auto_snapshot_policy": 4,
		"baiducloud_bos_bucket":           5,
		"baiducloud_bos_bucket_object":    6,
		"baiducloud_cds":                  7,
		"baiducloud_cds_attachment":       8,
		"baiducloud_cert":                 9,
		"baiducloud_cfc_alias":            10,
		"baiducloud_cfc_function":         11,
		"baiducloud_cfc_trigger":          12,
		"baiducloud_cfc_version":          13,
		"baiducloud_eip":                  14,
		"baiducloud_eip_association":      15,
		"baiducloud_instance":             16,
		"baiducloud_nat_gateway":          17,
		"baiducloud_peer_conn":            18,
		"baiducloud_peer_conn_acceptor":   19,
		"baiducloud_route_rule":           20,
		"baiducloud_security_group":       21,
		"baiducloud_security_group_rule":  22,
		"baiducloud_snapshot":             23,
		"baiducloud_subnet":               24,
		"baiducloud_vpc":                  25,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
