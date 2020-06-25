package baiducloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "baiducloud_acls",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query ACL list.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "acl_id",
					Description: `(Optional) ID of the ACL to retrieve.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional, ForceNew) only support filter string/int/bool value`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional, ForceNew) Output file for saving result.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) Subnet ID of the ACLs to retrieve.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) VPC ID of the ACLs to retrieve. The ` + "`" + `filter` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) filter variable name`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) filter variable value list ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "acls",
					Description: `List of the ACLs.`,
				},
				resource.Attribute{
					Name:        "acl_id",
					Description: `ID of the ACL.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action of the ACL.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the ACL.`,
				},
				resource.Attribute{
					Name:        "destination_ip_address",
					Description: `Destination IP address of the ACL.`,
				},
				resource.Attribute{
					Name:        "destination_port",
					Description: `Destination port of the ACL.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `Direction of the ACL.`,
				},
				resource.Attribute{
					Name:        "position",
					Description: `Position of the ACL.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol of the ACL.`,
				},
				resource.Attribute{
					Name:        "source_ip_address",
					Description: `Source IP address of the ACL.`,
				},
				resource.Attribute{
					Name:        "source_port",
					Description: `Source port of the ACL.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `Subnet ID of the ACL.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "acls",
					Description: `List of the ACLs.`,
				},
				resource.Attribute{
					Name:        "acl_id",
					Description: `ID of the ACL.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action of the ACL.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the ACL.`,
				},
				resource.Attribute{
					Name:        "destination_ip_address",
					Description: `Destination IP address of the ACL.`,
				},
				resource.Attribute{
					Name:        "destination_port",
					Description: `Destination port of the ACL.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `Direction of the ACL.`,
				},
				resource.Attribute{
					Name:        "position",
					Description: `Position of the ACL.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol of the ACL.`,
				},
				resource.Attribute{
					Name:        "source_ip_address",
					Description: `Source IP address of the ACL.`,
				},
				resource.Attribute{
					Name:        "source_port",
					Description: `Source port of the ACL.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `Subnet ID of the ACL.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "baiducloud_appblb_listeners",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query APPBLB Listener list.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "blb_id",
					Description: `(Required, ForceNew) ID of the Application LoadBalance instance`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional, ForceNew) only support filter string/int/bool value`,
				},
				resource.Attribute{
					Name:        "listener_port",
					Description: `(Optional, ForceNew) The port of the Listener to be queried`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional, ForceNew) Query result output file path`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional, ForceNew) Protocol of the Listener to be queried The ` + "`" + `filter` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) filter variable name`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) filter variable value list ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "listeners",
					Description: `A list of Application LoadBalance Listener`,
				},
				resource.Attribute{
					Name:        "cert_ids",
					Description: `Listener bind certifications`,
				},
				resource.Attribute{
					Name:        "client_cert_ids",
					Description: `Listener import cert list, only useful when dual_auth is true`,
				},
				resource.Attribute{
					Name:        "dual_auth",
					Description: `Listener open dual authorization or not, default false`,
				},
				resource.Attribute{
					Name:        "encryption_protocols",
					Description: `Listener encryption protocol`,
				},
				resource.Attribute{
					Name:        "encryption_type",
					Description: `Listener encryption option`,
				},
				resource.Attribute{
					Name:        "ie6_compatible",
					Description: `Listener support ie6 option, default true`,
				},
				resource.Attribute{
					Name:        "keep_session_cookie_name",
					Description: `Listener keepSeesionCookieName`,
				},
				resource.Attribute{
					Name:        "keep_session_timeout",
					Description: `Listener keepSessionTimeout value`,
				},
				resource.Attribute{
					Name:        "keep_session_type",
					Description: `Listener keepSessionType option`,
				},
				resource.Attribute{
					Name:        "keep_session",
					Description: `Listener keepSession or not`,
				},
				resource.Attribute{
					Name:        "listener_port",
					Description: `Listener bind port`,
				},
				resource.Attribute{
					Name:        "policys",
					Description: `Listener's policy`,
				},
				resource.Attribute{
					Name:        "app_server_group_id",
					Description: `Policy bind server group ID`,
				},
				resource.Attribute{
					Name:        "app_server_group_name",
					Description: `Policy bind server group name`,
				},
				resource.Attribute{
					Name:        "backend_port",
					Description: `Backend port`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Policy's description`,
				},
				resource.Attribute{
					Name:        "frontend_port",
					Description: `Frontend port`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Policy's ID`,
				},
				resource.Attribute{
					Name:        "port_type",
					Description: `Policy bind port protocol`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Policy priority`,
				},
				resource.Attribute{
					Name:        "rule_list",
					Description: `Policy rule list`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `Rule key`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Rule value`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Listener protocol`,
				},
				resource.Attribute{
					Name:        "redirect_port",
					Description: `Listener redirect request to https listener port`,
				},
				resource.Attribute{
					Name:        "scheduler",
					Description: `Load balancing algorithm`,
				},
				resource.Attribute{
					Name:        "server_timeout",
					Description: `Backend server maximum timeout time, only support in [1, 3600] second, default 30s`,
				},
				resource.Attribute{
					Name:        "tcp_session_timeout",
					Description: `TCP Listener connetion session timeout time`,
				},
				resource.Attribute{
					Name:        "x_forwarded_for",
					Description: `Listener xForwardedFor, determine get client real ip or not, default false`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "listeners",
					Description: `A list of Application LoadBalance Listener`,
				},
				resource.Attribute{
					Name:        "cert_ids",
					Description: `Listener bind certifications`,
				},
				resource.Attribute{
					Name:        "client_cert_ids",
					Description: `Listener import cert list, only useful when dual_auth is true`,
				},
				resource.Attribute{
					Name:        "dual_auth",
					Description: `Listener open dual authorization or not, default false`,
				},
				resource.Attribute{
					Name:        "encryption_protocols",
					Description: `Listener encryption protocol`,
				},
				resource.Attribute{
					Name:        "encryption_type",
					Description: `Listener encryption option`,
				},
				resource.Attribute{
					Name:        "ie6_compatible",
					Description: `Listener support ie6 option, default true`,
				},
				resource.Attribute{
					Name:        "keep_session_cookie_name",
					Description: `Listener keepSeesionCookieName`,
				},
				resource.Attribute{
					Name:        "keep_session_timeout",
					Description: `Listener keepSessionTimeout value`,
				},
				resource.Attribute{
					Name:        "keep_session_type",
					Description: `Listener keepSessionType option`,
				},
				resource.Attribute{
					Name:        "keep_session",
					Description: `Listener keepSession or not`,
				},
				resource.Attribute{
					Name:        "listener_port",
					Description: `Listener bind port`,
				},
				resource.Attribute{
					Name:        "policys",
					Description: `Listener's policy`,
				},
				resource.Attribute{
					Name:        "app_server_group_id",
					Description: `Policy bind server group ID`,
				},
				resource.Attribute{
					Name:        "app_server_group_name",
					Description: `Policy bind server group name`,
				},
				resource.Attribute{
					Name:        "backend_port",
					Description: `Backend port`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Policy's description`,
				},
				resource.Attribute{
					Name:        "frontend_port",
					Description: `Frontend port`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Policy's ID`,
				},
				resource.Attribute{
					Name:        "port_type",
					Description: `Policy bind port protocol`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Policy priority`,
				},
				resource.Attribute{
					Name:        "rule_list",
					Description: `Policy rule list`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `Rule key`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Rule value`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Listener protocol`,
				},
				resource.Attribute{
					Name:        "redirect_port",
					Description: `Listener redirect request to https listener port`,
				},
				resource.Attribute{
					Name:        "scheduler",
					Description: `Load balancing algorithm`,
				},
				resource.Attribute{
					Name:        "server_timeout",
					Description: `Backend server maximum timeout time, only support in [1, 3600] second, default 30s`,
				},
				resource.Attribute{
					Name:        "tcp_session_timeout",
					Description: `TCP Listener connetion session timeout time`,
				},
				resource.Attribute{
					Name:        "x_forwarded_for",
					Description: `Listener xForwardedFor, determine get client real ip or not, default false`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "baiducloud_appblb_server_groups",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query APPBLB Server Group list.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "blb_id",
					Description: `(Required) ID of the LoadBalance instance to be queried`,
				},
				resource.Attribute{
					Name:        "exactly_match",
					Description: `(Optional) Whether the name is an exact match or not, default false`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional, ForceNew) only support filter string/int/bool value`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the Server Group to be queried`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional, ForceNew) Query result output file path The ` + "`" + `filter` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) filter variable name`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) filter variable value list ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "server_groups",
					Description: `A list of Application LoadBalance Server Group`,
				},
				resource.Attribute{
					Name:        "backend_server_list",
					Description: `Server group bound backend server list`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `Backend server instance ID`,
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
					Description: `Port bind policy ID`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `Port ID`,
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
					Description: `Backend server instance bind private ip`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Backend server instance weight in this group`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Server Group's description`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Server Group's name`,
				},
				resource.Attribute{
					Name:        "port_list",
					Description: `Server Group backend port list`,
				},
				resource.Attribute{
					Name:        "health_check_down_retry",
					Description: `Server Group health check down retry time`,
				},
				resource.Attribute{
					Name:        "health_check_interval_in_second",
					Description: `Server Group health check interval time(second)`,
				},
				resource.Attribute{
					Name:        "health_check_normal_status",
					Description: `Server Group health check normal http status code`,
				},
				resource.Attribute{
					Name:        "health_check_port",
					Description: `Server Group port health check port`,
				},
				resource.Attribute{
					Name:        "health_check_timeout_in_second",
					Description: `Server Group health check timeout(second)`,
				},
				resource.Attribute{
					Name:        "health_check_up_retry",
					Description: `Server Group health check up retry time`,
				},
				resource.Attribute{
					Name:        "health_check_url_path",
					Description: `Server Group health check url path`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `Server Group port health check protocol`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Server Group port ID`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Server Group port`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Server Group port status`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Server Group port protocol type`,
				},
				resource.Attribute{
					Name:        "udp_health_check_string",
					Description: `Server Group udp health check string`,
				},
				resource.Attribute{
					Name:        "sg_id",
					Description: `Server Group's ID`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Server Group status`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "server_groups",
					Description: `A list of Application LoadBalance Server Group`,
				},
				resource.Attribute{
					Name:        "backend_server_list",
					Description: `Server group bound backend server list`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `Backend server instance ID`,
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
					Description: `Port bind policy ID`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `Port ID`,
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
					Description: `Backend server instance bind private ip`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Backend server instance weight in this group`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Server Group's description`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Server Group's name`,
				},
				resource.Attribute{
					Name:        "port_list",
					Description: `Server Group backend port list`,
				},
				resource.Attribute{
					Name:        "health_check_down_retry",
					Description: `Server Group health check down retry time`,
				},
				resource.Attribute{
					Name:        "health_check_interval_in_second",
					Description: `Server Group health check interval time(second)`,
				},
				resource.Attribute{
					Name:        "health_check_normal_status",
					Description: `Server Group health check normal http status code`,
				},
				resource.Attribute{
					Name:        "health_check_port",
					Description: `Server Group port health check port`,
				},
				resource.Attribute{
					Name:        "health_check_timeout_in_second",
					Description: `Server Group health check timeout(second)`,
				},
				resource.Attribute{
					Name:        "health_check_up_retry",
					Description: `Server Group health check up retry time`,
				},
				resource.Attribute{
					Name:        "health_check_url_path",
					Description: `Server Group health check url path`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `Server Group port health check protocol`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Server Group port ID`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Server Group port`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Server Group port status`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Server Group port protocol type`,
				},
				resource.Attribute{
					Name:        "udp_health_check_string",
					Description: `Server Group udp health check string`,
				},
				resource.Attribute{
					Name:        "sg_id",
					Description: `Server Group's ID`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Server Group status`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "baiducloud_appblbs",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query APPBLB list.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "address",
					Description: `(Optional) Address ip of the LoadBalance instance to be queried`,
				},
				resource.Attribute{
					Name:        "bcc_id",
					Description: `(Optional) ID of the BCC instance bound to the LoadBalance`,
				},
				resource.Attribute{
					Name:        "blb_id",
					Description: `(Optional) ID of the LoadBalance instance to be queried`,
				},
				resource.Attribute{
					Name:        "exactly_match",
					Description: `(Optional) Whether the query condition is an exact match or not, default false`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional, ForceNew) only support filter string/int/bool value`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the LoadBalance instance to be queried`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional, ForceNew) Query result output file path The ` + "`" + `filter` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) filter variable name`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) filter variable value list ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "appblbs",
					Description: `A list of Application LoadBalance Instance`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `LoadBalance instance's service IP, instance can be accessed through this IP`,
				},
				resource.Attribute{
					Name:        "blb_id",
					Description: `LoadBalance instance's ID`,
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
					Name:        "description",
					Description: `LoadBalance instance's description`,
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
					Name:        "name",
					Description: `LoadBalance instance's name`,
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
					Description: `LoadBalance instance's status`,
				},
				resource.Attribute{
					Name:        "subnet_cidr",
					Description: `Cidr of the subnet which the LoadBalance instance belongs`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The subnet ID to which the LoadBalance instance belongs`,
				},
				resource.Attribute{
					Name:        "subnet_name",
					Description: `The subnet name to which the LoadBalance instance belongs`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The VPC short ID to which the LoadBalance instance belongs`,
				},
				resource.Attribute{
					Name:        "vpc_name",
					Description: `The VPC name to which the LoadBalance instance belongs`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "appblbs",
					Description: `A list of Application LoadBalance Instance`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `LoadBalance instance's service IP, instance can be accessed through this IP`,
				},
				resource.Attribute{
					Name:        "blb_id",
					Description: `LoadBalance instance's ID`,
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
					Name:        "description",
					Description: `LoadBalance instance's description`,
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
					Name:        "name",
					Description: `LoadBalance instance's name`,
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
					Description: `LoadBalance instance's status`,
				},
				resource.Attribute{
					Name:        "subnet_cidr",
					Description: `Cidr of the subnet which the LoadBalance instance belongs`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The subnet ID to which the LoadBalance instance belongs`,
				},
				resource.Attribute{
					Name:        "subnet_name",
					Description: `The subnet name to which the LoadBalance instance belongs`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The VPC short ID to which the LoadBalance instance belongs`,
				},
				resource.Attribute{
					Name:        "vpc_name",
					Description: `The VPC name to which the LoadBalance instance belongs`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "baiducloud_auto_snapshot_policies",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query Auto Snapshot Policy list.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "asp_name",
					Description: `(Optional) Name of the automatic snapshot policy.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional, ForceNew) only support filter string/int/bool value`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional, ForceNew) Automatic snapshot policies search result output file.`,
				},
				resource.Attribute{
					Name:        "volume_name",
					Description: `(Optional) Name of the volume. The ` + "`" + `filter` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) filter variable name`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) filter variable value list ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "auto_snapshot_policies",
					Description: `The automatic snapshot policies search result list.`,
				},
				resource.Attribute{
					Name:        "created_time",
					Description: `The creation time of the automatic snapshot policy.`,
				},
				resource.Attribute{
					Name:        "deleted_time",
					Description: `The deletion time of the automatic snapshot policy.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the automatic snapshot policy.`,
				},
				resource.Attribute{
					Name:        "last_execute_time",
					Description: `The last execution time of the automatic snapshot policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the automatic snapshot policy.`,
				},
				resource.Attribute{
					Name:        "repeat_weekdays",
					Description: `The repeat weekdays of the automatic snapshot policy.`,
				},
				resource.Attribute{
					Name:        "retention_days",
					Description: `The retention days of the automatic snapshot policy.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the automatic snapshot policy.`,
				},
				resource.Attribute{
					Name:        "time_points",
					Description: `The time points of the automatic snapshot policy.`,
				},
				resource.Attribute{
					Name:        "updated_time",
					Description: `The updation time of the automatic snapshot policy.`,
				},
				resource.Attribute{
					Name:        "volume_count",
					Description: `The volume count of the automatic snapshot policy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "auto_snapshot_policies",
					Description: `The automatic snapshot policies search result list.`,
				},
				resource.Attribute{
					Name:        "created_time",
					Description: `The creation time of the automatic snapshot policy.`,
				},
				resource.Attribute{
					Name:        "deleted_time",
					Description: `The deletion time of the automatic snapshot policy.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the automatic snapshot policy.`,
				},
				resource.Attribute{
					Name:        "last_execute_time",
					Description: `The last execution time of the automatic snapshot policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the automatic snapshot policy.`,
				},
				resource.Attribute{
					Name:        "repeat_weekdays",
					Description: `The repeat weekdays of the automatic snapshot policy.`,
				},
				resource.Attribute{
					Name:        "retention_days",
					Description: `The retention days of the automatic snapshot policy.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the automatic snapshot policy.`,
				},
				resource.Attribute{
					Name:        "time_points",
					Description: `The time points of the automatic snapshot policy.`,
				},
				resource.Attribute{
					Name:        "updated_time",
					Description: `The updation time of the automatic snapshot policy.`,
				},
				resource.Attribute{
					Name:        "volume_count",
					Description: `The volume count of the automatic snapshot policy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "baiducloud_bos_bucket_objects",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query BOS bucket object list.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) Bucket name of the objects to retrieve.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional, ForceNew) only support filter string/int/bool value`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional, ForceNew) Output file for saving result.`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Optional) Prefix of the objects. The ` + "`" + `filter` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) filter variable name`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) filter variable value list ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "objects",
					Description: `List of the objects.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `Acl of the object.`,
				},
				resource.Attribute{
					Name:        "bucket",
					Description: `Bucket of the object.`,
				},
				resource.Attribute{
					Name:        "cache_control",
					Description: `Caching behavior of the object.`,
				},
				resource.Attribute{
					Name:        "content_crc32",
					Description: `Crc(cyclic redundancy check code) value of the object.`,
				},
				resource.Attribute{
					Name:        "content_disposition",
					Description: `Content disposition of the object.`,
				},
				resource.Attribute{
					Name:        "content_encoding",
					Description: `Encoding of the object.`,
				},
				resource.Attribute{
					Name:        "content_length",
					Description: `Content length of the object.`,
				},
				resource.Attribute{
					Name:        "content_md5",
					Description: `MD5 value of the object content defined in RFC2616.`,
				},
				resource.Attribute{
					Name:        "content_sha256",
					Description: `Sha256 value of the object.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `Content type of the object data.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `Etag of the object.`,
				},
				resource.Attribute{
					Name:        "expires",
					Description: `Expire date of the object.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `Key of the object.`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `Last modifyed time of the object.`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `Owner id of the object.`,
				},
				resource.Attribute{
					Name:        "owner_name",
					Description: `Owner name of the object.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Size of the object.`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `Storage class of the object.`,
				},
				resource.Attribute{
					Name:        "user_meta",
					Description: `Metadata of the object.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "objects",
					Description: `List of the objects.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `Acl of the object.`,
				},
				resource.Attribute{
					Name:        "bucket",
					Description: `Bucket of the object.`,
				},
				resource.Attribute{
					Name:        "cache_control",
					Description: `Caching behavior of the object.`,
				},
				resource.Attribute{
					Name:        "content_crc32",
					Description: `Crc(cyclic redundancy check code) value of the object.`,
				},
				resource.Attribute{
					Name:        "content_disposition",
					Description: `Content disposition of the object.`,
				},
				resource.Attribute{
					Name:        "content_encoding",
					Description: `Encoding of the object.`,
				},
				resource.Attribute{
					Name:        "content_length",
					Description: `Content length of the object.`,
				},
				resource.Attribute{
					Name:        "content_md5",
					Description: `MD5 value of the object content defined in RFC2616.`,
				},
				resource.Attribute{
					Name:        "content_sha256",
					Description: `Sha256 value of the object.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `Content type of the object data.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `Etag of the object.`,
				},
				resource.Attribute{
					Name:        "expires",
					Description: `Expire date of the object.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `Key of the object.`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `Last modifyed time of the object.`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `Owner id of the object.`,
				},
				resource.Attribute{
					Name:        "owner_name",
					Description: `Owner name of the object.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Size of the object.`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `Storage class of the object.`,
				},
				resource.Attribute{
					Name:        "user_meta",
					Description: `Metadata of the object.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "baiducloud_bos_buckets",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query BOS bucket list.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Optional) Name of the bucket to retrieve.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional, ForceNew) only support filter string/int/bool value`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional, ForceNew) Output file for saving result. The ` + "`" + `filter` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) filter variable name`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) filter variable value list ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "buckets",
					Description: `List of buckets.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `Acl of the bucket.`,
				},
				resource.Attribute{
					Name:        "bucket",
					Description: `Name of the bucket.`,
				},
				resource.Attribute{
					Name:        "copyright_protection",
					Description: `Configuration of the copyright protection.`,
				},
				resource.Attribute{
					Name:        "resource",
					Description: `Resources to be protected for copyright.`,
				},
				resource.Attribute{
					Name:        "cors_rule",
					Description: `Configuration of the Cross-Origin Resource Sharing.`,
				},
				resource.Attribute{
					Name:        "allowed_expose_headers",
					Description: `Indicate which expose headers are allowed.`,
				},
				resource.Attribute{
					Name:        "allowed_headers",
					Description: `Indicate which headers are allowed.`,
				},
				resource.Attribute{
					Name:        "allowed_methods",
					Description: `Indicate which methods are allowed.`,
				},
				resource.Attribute{
					Name:        "allowed_origins",
					Description: `Indicate which origins are allowed.`,
				},
				resource.Attribute{
					Name:        "max_age_seconds",
					Description: `Indicate time in seconds that browser can cache the response for a preflight request.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `Creation date of the bucket.`,
				},
				resource.Attribute{
					Name:        "lifecycle_rule",
					Description: `Configuration of object lifecycle management.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action of the lifecycle rule.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Action name of the lifecycle rule.`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `Storage class of the action.`,
				},
				resource.Attribute{
					Name:        "condition",
					Description: `Condition of the lifecycle rule.`,
				},
				resource.Attribute{
					Name:        "time",
					Description: `Condition time, implemented by the date_greater_than.`,
				},
				resource.Attribute{
					Name:        "date_greater_than",
					Description: `Support absolute time date and relative time days.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the lifecycle rule.`,
				},
				resource.Attribute{
					Name:        "resource",
					Description: `Resource of the lifecycle rule.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the lifecycle rule.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Bucket location of the bucket.`,
				},
				resource.Attribute{
					Name:        "logging",
					Description: `Logging of the bucket.`,
				},
				resource.Attribute{
					Name:        "target_bucket",
					Description: `Target bucket name of the logging.`,
				},
				resource.Attribute{
					Name:        "target_prefix",
					Description: `Target prefix of the logging.`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `Owner id of the bucket.`,
				},
				resource.Attribute{
					Name:        "owner_name",
					Description: `Owner name of the bucket.`,
				},
				resource.Attribute{
					Name:        "replication_configuration",
					Description: `Replication configuration of the bucket.`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `Destination of the replication configuration.`,
				},
				resource.Attribute{
					Name:        "bucket",
					Description: `Destination bucket name of the replication configuration.`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `Destination storage class of the replication configuration.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the replication configuration.`,
				},
				resource.Attribute{
					Name:        "replicate_deletes",
					Description: `Whether to enable the delete synchronization.`,
				},
				resource.Attribute{
					Name:        "replicate_history",
					Description: `Configuration of the replicate history.`,
				},
				resource.Attribute{
					Name:        "bucket",
					Description: `Destination bucket name of the replication configuration.`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `Destination storage class of the replication configuration.`,
				},
				resource.Attribute{
					Name:        "resource",
					Description: `Resource of the replication configuration.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the replication configuration.`,
				},
				resource.Attribute{
					Name:        "server_side_encryption_rule",
					Description: `Encryption of the bucket.`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `Storage class of the bucket.`,
				},
				resource.Attribute{
					Name:        "website",
					Description: `Website of the BOS bucket.`,
				},
				resource.Attribute{
					Name:        "error_document",
					Description: `An absolute path to the document to return in case of a 404 error.`,
				},
				resource.Attribute{
					Name:        "index_document",
					Description: `Baiducloud BOS returns this index document when requests are made to the root domain or any of the subfolders.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "buckets",
					Description: `List of buckets.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `Acl of the bucket.`,
				},
				resource.Attribute{
					Name:        "bucket",
					Description: `Name of the bucket.`,
				},
				resource.Attribute{
					Name:        "copyright_protection",
					Description: `Configuration of the copyright protection.`,
				},
				resource.Attribute{
					Name:        "resource",
					Description: `Resources to be protected for copyright.`,
				},
				resource.Attribute{
					Name:        "cors_rule",
					Description: `Configuration of the Cross-Origin Resource Sharing.`,
				},
				resource.Attribute{
					Name:        "allowed_expose_headers",
					Description: `Indicate which expose headers are allowed.`,
				},
				resource.Attribute{
					Name:        "allowed_headers",
					Description: `Indicate which headers are allowed.`,
				},
				resource.Attribute{
					Name:        "allowed_methods",
					Description: `Indicate which methods are allowed.`,
				},
				resource.Attribute{
					Name:        "allowed_origins",
					Description: `Indicate which origins are allowed.`,
				},
				resource.Attribute{
					Name:        "max_age_seconds",
					Description: `Indicate time in seconds that browser can cache the response for a preflight request.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `Creation date of the bucket.`,
				},
				resource.Attribute{
					Name:        "lifecycle_rule",
					Description: `Configuration of object lifecycle management.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action of the lifecycle rule.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Action name of the lifecycle rule.`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `Storage class of the action.`,
				},
				resource.Attribute{
					Name:        "condition",
					Description: `Condition of the lifecycle rule.`,
				},
				resource.Attribute{
					Name:        "time",
					Description: `Condition time, implemented by the date_greater_than.`,
				},
				resource.Attribute{
					Name:        "date_greater_than",
					Description: `Support absolute time date and relative time days.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the lifecycle rule.`,
				},
				resource.Attribute{
					Name:        "resource",
					Description: `Resource of the lifecycle rule.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the lifecycle rule.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Bucket location of the bucket.`,
				},
				resource.Attribute{
					Name:        "logging",
					Description: `Logging of the bucket.`,
				},
				resource.Attribute{
					Name:        "target_bucket",
					Description: `Target bucket name of the logging.`,
				},
				resource.Attribute{
					Name:        "target_prefix",
					Description: `Target prefix of the logging.`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `Owner id of the bucket.`,
				},
				resource.Attribute{
					Name:        "owner_name",
					Description: `Owner name of the bucket.`,
				},
				resource.Attribute{
					Name:        "replication_configuration",
					Description: `Replication configuration of the bucket.`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `Destination of the replication configuration.`,
				},
				resource.Attribute{
					Name:        "bucket",
					Description: `Destination bucket name of the replication configuration.`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `Destination storage class of the replication configuration.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the replication configuration.`,
				},
				resource.Attribute{
					Name:        "replicate_deletes",
					Description: `Whether to enable the delete synchronization.`,
				},
				resource.Attribute{
					Name:        "replicate_history",
					Description: `Configuration of the replicate history.`,
				},
				resource.Attribute{
					Name:        "bucket",
					Description: `Destination bucket name of the replication configuration.`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `Destination storage class of the replication configuration.`,
				},
				resource.Attribute{
					Name:        "resource",
					Description: `Resource of the replication configuration.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the replication configuration.`,
				},
				resource.Attribute{
					Name:        "server_side_encryption_rule",
					Description: `Encryption of the bucket.`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `Storage class of the bucket.`,
				},
				resource.Attribute{
					Name:        "website",
					Description: `Website of the BOS bucket.`,
				},
				resource.Attribute{
					Name:        "error_document",
					Description: `An absolute path to the document to return in case of a 404 error.`,
				},
				resource.Attribute{
					Name:        "index_document",
					Description: `Baiducloud BOS returns this index document when requests are made to the root domain or any of the subfolders.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "baiducloud_cdss",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query CDS list.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional, ForceNew) only support filter string/int/bool value`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Optional, ForceNew) CDS volume bind instance ID`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional, ForceNew) CDS volume search result output file`,
				},
				resource.Attribute{
					Name:        "zone_name",
					Description: `(Optional, ForceNew) CDS volume zone name The ` + "`" + `filter` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) filter variable name`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) filter variable value list ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cdss",
					Description: `CDS volume list`,
				},
				resource.Attribute{
					Name:        "attachments",
					Description: `CDS volume attachments`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `CDS attachment device path`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `CDS attachment instance id`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `CDS attachment serial`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `CDS attachment volume id`,
				},
				resource.Attribute{
					Name:        "auto_snapshot_policy",
					Description: `CDS volume bind auto snapshot policy info`,
				},
				resource.Attribute{
					Name:        "created_time",
					Description: `Auto Snapshot Policy created time`,
				},
				resource.Attribute{
					Name:        "deleted_time",
					Description: `Auto Snapshot Policy deleted time`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Auto Snapshot Policy ID`,
				},
				resource.Attribute{
					Name:        "last_execute_time",
					Description: `Auto Snapshot Policy last execute time`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Auto Snapshot Policy name`,
				},
				resource.Attribute{
					Name:        "repeat_weekdays",
					Description: `Auto Snapshot Policy repeat weekdays`,
				},
				resource.Attribute{
					Name:        "retention_days",
					Description: `Auto Snapshot Policy retention days`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Auto Snapshot Policy status`,
				},
				resource.Attribute{
					Name:        "time_points",
					Description: `Auto Snapshot Policy set snapshot create time points`,
				},
				resource.Attribute{
					Name:        "updated_time",
					Description: `Auto Snapshot Policy updated time`,
				},
				resource.Attribute{
					Name:        "volume_count",
					Description: `Auto Snapshot Policy volume count`,
				},
				resource.Attribute{
					Name:        "cds_id",
					Description: `CDS volume id`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `CDS disk create time`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `CDS description`,
				},
				resource.Attribute{
					Name:        "disk_size_in_gb",
					Description: `CDS disk size, should in [1, 32765], when snapshot_id not set, this parameter is required.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `CDS disk expire time`,
				},
				resource.Attribute{
					Name:        "is_system_volume",
					Description: `CDS disk is system volume or not`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `CDS disk name`,
				},
				resource.Attribute{
					Name:        "payment_timing",
					Description: `payment method, support Prepaid or Postpaid`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `CDS disk region id`,
				},
				resource.Attribute{
					Name:        "snapshot_num",
					Description: `CDS disk snapshot num`,
				},
				resource.Attribute{
					Name:        "source_snapshot_id",
					Description: `CDS disk create source snapshot id`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `CDS volume status`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `CDS dist storage type, support hp1 and std1, default hp1`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `CDS disk type`,
				},
				resource.Attribute{
					Name:        "zone_name",
					Description: `Zone name`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cdss",
					Description: `CDS volume list`,
				},
				resource.Attribute{
					Name:        "attachments",
					Description: `CDS volume attachments`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `CDS attachment device path`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `CDS attachment instance id`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `CDS attachment serial`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `CDS attachment volume id`,
				},
				resource.Attribute{
					Name:        "auto_snapshot_policy",
					Description: `CDS volume bind auto snapshot policy info`,
				},
				resource.Attribute{
					Name:        "created_time",
					Description: `Auto Snapshot Policy created time`,
				},
				resource.Attribute{
					Name:        "deleted_time",
					Description: `Auto Snapshot Policy deleted time`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Auto Snapshot Policy ID`,
				},
				resource.Attribute{
					Name:        "last_execute_time",
					Description: `Auto Snapshot Policy last execute time`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Auto Snapshot Policy name`,
				},
				resource.Attribute{
					Name:        "repeat_weekdays",
					Description: `Auto Snapshot Policy repeat weekdays`,
				},
				resource.Attribute{
					Name:        "retention_days",
					Description: `Auto Snapshot Policy retention days`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Auto Snapshot Policy status`,
				},
				resource.Attribute{
					Name:        "time_points",
					Description: `Auto Snapshot Policy set snapshot create time points`,
				},
				resource.Attribute{
					Name:        "updated_time",
					Description: `Auto Snapshot Policy updated time`,
				},
				resource.Attribute{
					Name:        "volume_count",
					Description: `Auto Snapshot Policy volume count`,
				},
				resource.Attribute{
					Name:        "cds_id",
					Description: `CDS volume id`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `CDS disk create time`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `CDS description`,
				},
				resource.Attribute{
					Name:        "disk_size_in_gb",
					Description: `CDS disk size, should in [1, 32765], when snapshot_id not set, this parameter is required.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `CDS disk expire time`,
				},
				resource.Attribute{
					Name:        "is_system_volume",
					Description: `CDS disk is system volume or not`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `CDS disk name`,
				},
				resource.Attribute{
					Name:        "payment_timing",
					Description: `payment method, support Prepaid or Postpaid`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `CDS disk region id`,
				},
				resource.Attribute{
					Name:        "snapshot_num",
					Description: `CDS disk snapshot num`,
				},
				resource.Attribute{
					Name:        "source_snapshot_id",
					Description: `CDS disk create source snapshot id`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `CDS volume status`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `CDS dist storage type, support hp1 and std1, default hp1`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `CDS disk type`,
				},
				resource.Attribute{
					Name:        "zone_name",
					Description: `Zone name`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "baiducloud_certs",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query CERT list.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cert_name",
					Description: `(Optional, ForceNew) Name of the Cert to be queried`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional, ForceNew) only support filter string/int/bool value`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional, ForceNew) Certs search result output file The ` + "`" + `filter` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) filter variable name`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) filter variable value list ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "certs",
					Description: `A list of Cert`,
				},
				resource.Attribute{
					Name:        "cert_common_name",
					Description: `Cert's common name`,
				},
				resource.Attribute{
					Name:        "cert_create_time",
					Description: `Cert's create time`,
				},
				resource.Attribute{
					Name:        "cert_id",
					Description: `Cert's ID`,
				},
				resource.Attribute{
					Name:        "cert_name",
					Description: `Cert's name`,
				},
				resource.Attribute{
					Name:        "cert_start_time",
					Description: `Cert's start time`,
				},
				resource.Attribute{
					Name:        "cert_stop_time",
					Description: `Cert's stop time`,
				},
				resource.Attribute{
					Name:        "cert_type",
					Description: `Cert's type`,
				},
				resource.Attribute{
					Name:        "cert_update_time",
					Description: `Cert's update time`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "certs",
					Description: `A list of Cert`,
				},
				resource.Attribute{
					Name:        "cert_common_name",
					Description: `Cert's common name`,
				},
				resource.Attribute{
					Name:        "cert_create_time",
					Description: `Cert's create time`,
				},
				resource.Attribute{
					Name:        "cert_id",
					Description: `Cert's ID`,
				},
				resource.Attribute{
					Name:        "cert_name",
					Description: `Cert's name`,
				},
				resource.Attribute{
					Name:        "cert_start_time",
					Description: `Cert's start time`,
				},
				resource.Attribute{
					Name:        "cert_stop_time",
					Description: `Cert's stop time`,
				},
				resource.Attribute{
					Name:        "cert_type",
					Description: `Cert's type`,
				},
				resource.Attribute{
					Name:        "cert_update_time",
					Description: `Cert's update time`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "baiducloud_cfc_function",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get a function.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "function_name",
					Description: `(Required, ForceNew) CFC function name, length must be between 1 and 64 bytes`,
				},
				resource.Attribute{
					Name:        "qualifier",
					Description: `(Optional) Function search qualifier ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
					Name:        "code_zip_file",
					Description: `CFC Function Code base64-encoded data`,
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
					Name:        "log_bos_dir",
					Description: `Log save dir if log type is bos`,
				},
				resource.Attribute{
					Name:        "log_type",
					Description: `Log save type, support bos/none`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `CFC Function memory size, should be an integer multiple of 128`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Function region`,
				},
				resource.Attribute{
					Name:        "reserved_concurrent_executions",
					Description: `Function reserved concurrent executions, support [0-90]`,
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
					Name:        "source_tag",
					Description: `CFC Function source tag`,
				},
				resource.Attribute{
					Name:        "time_out",
					Description: `Function time out, support [1, 300]s`,
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
					Description: `Function version, should only be $LATEST`,
				},
				resource.Attribute{
					Name:        "vpc_config",
					Description: `Function VPC Config`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `CFC Function binded Security group list`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `CFC Function bined VPC Subnet id list`,
				},
			},
			Attributes: []resource.Attribute{
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
					Name:        "code_zip_file",
					Description: `CFC Function Code base64-encoded data`,
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
					Name:        "log_bos_dir",
					Description: `Log save dir if log type is bos`,
				},
				resource.Attribute{
					Name:        "log_type",
					Description: `Log save type, support bos/none`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `CFC Function memory size, should be an integer multiple of 128`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Function region`,
				},
				resource.Attribute{
					Name:        "reserved_concurrent_executions",
					Description: `Function reserved concurrent executions, support [0-90]`,
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
					Name:        "source_tag",
					Description: `CFC Function source tag`,
				},
				resource.Attribute{
					Name:        "time_out",
					Description: `Function time out, support [1, 300]s`,
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
					Description: `Function version, should only be $LATEST`,
				},
				resource.Attribute{
					Name:        "vpc_config",
					Description: `Function VPC Config`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `CFC Function binded Security group list`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `CFC Function bined VPC Subnet id list`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "baiducloud_eips",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query EIP list.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "eip",
					Description: `(Optional, ForceNew) Eip address`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional, ForceNew) only support filter string/int/bool value`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Optional, ForceNew) Eip bind instance id`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Optional, ForceNew) Eip bind instance type`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional, ForceNew) Eips search result output file`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional, ForceNew) Eip status The ` + "`" + `filter` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) filter variable name`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) filter variable value list ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "eips",
					Description: `Eip list`,
				},
				resource.Attribute{
					Name:        "bandwidth_in_mbps",
					Description: `Eip bandwidth(Mbps)`,
				},
				resource.Attribute{
					Name:        "billing_method",
					Description: `Eip billing method`,
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
					Name:        "name",
					Description: `Eip name`,
				},
				resource.Attribute{
					Name:        "payment_timing",
					Description: `Eip payment timing`,
				},
				resource.Attribute{
					Name:        "share_group_id",
					Description: `Eip share group id`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Eip status`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "eips",
					Description: `Eip list`,
				},
				resource.Attribute{
					Name:        "bandwidth_in_mbps",
					Description: `Eip bandwidth(Mbps)`,
				},
				resource.Attribute{
					Name:        "billing_method",
					Description: `Eip billing method`,
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
					Name:        "name",
					Description: `Eip name`,
				},
				resource.Attribute{
					Name:        "payment_timing",
					Description: `Eip payment timing`,
				},
				resource.Attribute{
					Name:        "share_group_id",
					Description: `Eip share group id`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Eip status`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "baiducloud_images",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query image list.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional, ForceNew) only support filter string/int/bool value`,
				},
				resource.Attribute{
					Name:        "image_type",
					Description: `(Optional, ForceNew) Image type of the images to be queried, support ALL/System/Custom/Integration/Sharing/GpuBccSystem/GpuBccCustom/FpgaBccSystem/FpgaBccCustom`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional, ForceNew) Regex pattern of the search image name`,
				},
				resource.Attribute{
					Name:        "os_name",
					Description: `(Optional, ForceNew) Search image OS Name`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional, ForceNew) Images search result output file The ` + "`" + `filter` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) filter variable name`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) filter variable value list ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "images",
					Description: `Image list`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Image create time`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Image description`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Image id`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Image name`,
				},
				resource.Attribute{
					Name:        "os_arch",
					Description: `Image os arch`,
				},
				resource.Attribute{
					Name:        "os_build",
					Description: `Image os build`,
				},
				resource.Attribute{
					Name:        "os_name",
					Description: `Image os name`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `Image os type`,
				},
				resource.Attribute{
					Name:        "os_version",
					Description: `Image os version`,
				},
				resource.Attribute{
					Name:        "special_version",
					Description: `Image special version`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Image status`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Image type`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "images",
					Description: `Image list`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Image create time`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Image description`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Image id`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Image name`,
				},
				resource.Attribute{
					Name:        "os_arch",
					Description: `Image os arch`,
				},
				resource.Attribute{
					Name:        "os_build",
					Description: `Image os build`,
				},
				resource.Attribute{
					Name:        "os_name",
					Description: `Image os name`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `Image os type`,
				},
				resource.Attribute{
					Name:        "os_version",
					Description: `Image os version`,
				},
				resource.Attribute{
					Name:        "special_version",
					Description: `Image special version`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Image status`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Image type`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "baiducloud_instances",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query BCC Instance list.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dedicated_host_id",
					Description: `(Optional) Dedicated host id of the instance to retrieve.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional, ForceNew) only support filter string/int/bool value`,
				},
				resource.Attribute{
					Name:        "internal_ip",
					Description: `(Optional) Internal ip address of the instance to retrieve.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional, ForceNew) Output file of the instances search result`,
				},
				resource.Attribute{
					Name:        "zone_name",
					Description: `(Optional) Name of the available zone to which the instance belongs. The ` + "`" + `filter` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) filter variable name`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) filter variable value list ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `The result of the instances list.`,
				},
				resource.Attribute{
					Name:        "auto_renew",
					Description: `Whether to automatically renew.`,
				},
				resource.Attribute{
					Name:        "card_count",
					Description: `The card count of the instance.`,
				},
				resource.Attribute{
					Name:        "cpu_count",
					Description: `The cpu count of the instance.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The creation time of the instance.`,
				},
				resource.Attribute{
					Name:        "dedicated_host_id",
					Description: `The dedicated host id of the instance.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the instance.`,
				},
				resource.Attribute{
					Name:        "ephemeral_disks",
					Description: `The ephemeral disks of the instance.`,
				},
				resource.Attribute{
					Name:        "size_in_gb",
					Description: `The size(GB) of the ephemeral disk.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `The storage type of the ephemeral disk.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `The expire time of the instance.`,
				},
				resource.Attribute{
					Name:        "fpga_card",
					Description: `The fgpa card of the instance.`,
				},
				resource.Attribute{
					Name:        "gpu_card",
					Description: `The gpu card of the instance.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The image id of the instance.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of the instance.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `The type of the instance.`,
				},
				resource.Attribute{
					Name:        "internal_ip",
					Description: `The internal ip of the instance.`,
				},
				resource.Attribute{
					Name:        "keypair_id",
					Description: `The key pair id of the instance.`,
				},
				resource.Attribute{
					Name:        "keypair_name",
					Description: `The key pair name of the instance.`,
				},
				resource.Attribute{
					Name:        "memory_capacity_in_gb",
					Description: `The memory capacity in GB of the instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the instance.`,
				},
				resource.Attribute{
					Name:        "network_capacity_in_mbps",
					Description: `The network capacity in Mbps of the instance.`,
				},
				resource.Attribute{
					Name:        "payment_timing",
					Description: `The payment timing of the instance.`,
				},
				resource.Attribute{
					Name:        "placement_policy",
					Description: `The placement policy of the instance.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public ip of the instance.`,
				},
				resource.Attribute{
					Name:        "root_disk_size_in_gb",
					Description: `The system disk size in GB of the instance.`,
				},
				resource.Attribute{
					Name:        "root_disk_storage_type",
					Description: `The system disk storage type of the instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the instance.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The subnet ID of the instance.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The VPC ID of the instance.`,
				},
				resource.Attribute{
					Name:        "zone_name",
					Description: `The zone name of the instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instances",
					Description: `The result of the instances list.`,
				},
				resource.Attribute{
					Name:        "auto_renew",
					Description: `Whether to automatically renew.`,
				},
				resource.Attribute{
					Name:        "card_count",
					Description: `The card count of the instance.`,
				},
				resource.Attribute{
					Name:        "cpu_count",
					Description: `The cpu count of the instance.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The creation time of the instance.`,
				},
				resource.Attribute{
					Name:        "dedicated_host_id",
					Description: `The dedicated host id of the instance.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the instance.`,
				},
				resource.Attribute{
					Name:        "ephemeral_disks",
					Description: `The ephemeral disks of the instance.`,
				},
				resource.Attribute{
					Name:        "size_in_gb",
					Description: `The size(GB) of the ephemeral disk.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `The storage type of the ephemeral disk.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `The expire time of the instance.`,
				},
				resource.Attribute{
					Name:        "fpga_card",
					Description: `The fgpa card of the instance.`,
				},
				resource.Attribute{
					Name:        "gpu_card",
					Description: `The gpu card of the instance.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The image id of the instance.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of the instance.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `The type of the instance.`,
				},
				resource.Attribute{
					Name:        "internal_ip",
					Description: `The internal ip of the instance.`,
				},
				resource.Attribute{
					Name:        "keypair_id",
					Description: `The key pair id of the instance.`,
				},
				resource.Attribute{
					Name:        "keypair_name",
					Description: `The key pair name of the instance.`,
				},
				resource.Attribute{
					Name:        "memory_capacity_in_gb",
					Description: `The memory capacity in GB of the instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the instance.`,
				},
				resource.Attribute{
					Name:        "network_capacity_in_mbps",
					Description: `The network capacity in Mbps of the instance.`,
				},
				resource.Attribute{
					Name:        "payment_timing",
					Description: `The payment timing of the instance.`,
				},
				resource.Attribute{
					Name:        "placement_policy",
					Description: `The placement policy of the instance.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public ip of the instance.`,
				},
				resource.Attribute{
					Name:        "root_disk_size_in_gb",
					Description: `The system disk size in GB of the instance.`,
				},
				resource.Attribute{
					Name:        "root_disk_storage_type",
					Description: `The system disk storage type of the instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the instance.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The subnet ID of the instance.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The VPC ID of the instance.`,
				},
				resource.Attribute{
					Name:        "zone_name",
					Description: `The zone name of the instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "baiducloud_nat_gateways",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query NAT gateway list.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional, ForceNew) only support filter string/int/bool value`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Optional) Specify the EIP binded by the NAT gateway to retrieve.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "nat_id",
					Description: `(Optional) ID of the NAT gateway to retrieve.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional, ForceNew) Output file for saving result.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) VPC ID where the NAT gateways located. The ` + "`" + `filter` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) filter variable name`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) filter variable value list ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "nat_gateways",
					Description: `The list of NAT gateways.`,
				},
				resource.Attribute{
					Name:        "eips",
					Description: `EIP list of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "expired_time",
					Description: `Expired time of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "payment_timing",
					Description: `Payment timing of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `Spec of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC ID of the NAT gateway.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "nat_gateways",
					Description: `The list of NAT gateways.`,
				},
				resource.Attribute{
					Name:        "eips",
					Description: `EIP list of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "expired_time",
					Description: `Expired time of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "payment_timing",
					Description: `Payment timing of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `Spec of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC ID of the NAT gateway.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "baiducloud_peer_conns",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query Peer Conn list.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional, ForceNew) only support filter string/int/bool value`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional, ForceNew) Output file for saving result.`,
				},
				resource.Attribute{
					Name:        "peer_conn_id",
					Description: `(Optional) ID of the peer connection to retrieve.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) VPC ID where the peer connections located. The ` + "`" + `filter` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) filter variable name`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) filter variable value list ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "peer_conns",
					Description: `The list of the peer connections.`,
				},
				resource.Attribute{
					Name:        "bandwidth_in_mbps",
					Description: `Bandwidth(Mbps) of the peer connection.`,
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
					Description: `Expired time of the peer connection.`,
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
					Name:        "local_region",
					Description: `Local region of the peer connection.`,
				},
				resource.Attribute{
					Name:        "local_vpc_id",
					Description: `Local VPC ID of the peer connection.`,
				},
				resource.Attribute{
					Name:        "payment_timing",
					Description: `Payment timing of the peer connection.`,
				},
				resource.Attribute{
					Name:        "peer_account_id",
					Description: `Peer account ID of the peer connection.`,
				},
				resource.Attribute{
					Name:        "peer_conn_id",
					Description: `ID of the peer connection.`,
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
					Description: `Role of the peer connection.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the peer connection.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "peer_conns",
					Description: `The list of the peer connections.`,
				},
				resource.Attribute{
					Name:        "bandwidth_in_mbps",
					Description: `Bandwidth(Mbps) of the peer connection.`,
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
					Description: `Expired time of the peer connection.`,
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
					Name:        "local_region",
					Description: `Local region of the peer connection.`,
				},
				resource.Attribute{
					Name:        "local_vpc_id",
					Description: `Local VPC ID of the peer connection.`,
				},
				resource.Attribute{
					Name:        "payment_timing",
					Description: `Payment timing of the peer connection.`,
				},
				resource.Attribute{
					Name:        "peer_account_id",
					Description: `Peer account ID of the peer connection.`,
				},
				resource.Attribute{
					Name:        "peer_conn_id",
					Description: `ID of the peer connection.`,
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
					Description: `Role of the peer connection.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the peer connection.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "baiducloud_route_rules",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query route rule list.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional, ForceNew) only support filter string/int/bool value`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional, ForceNew) Output file for saving result.`,
				},
				resource.Attribute{
					Name:        "route_rule_id",
					Description: `(Optional) ID of the routing rule to be retrieved.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `(Optional) Routing table ID for the routing rules to retrieve.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) VPC ID for the routing rules. The ` + "`" + `filter` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) filter variable name`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) filter variable value list ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "route_rules",
					Description: `Result of the routing rules.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the routing rule.`,
				},
				resource.Attribute{
					Name:        "destination_address",
					Description: `Destination address of the routing rule.`,
				},
				resource.Attribute{
					Name:        "next_hop_id",
					Description: `Next hop ID of the routing rule.`,
				},
				resource.Attribute{
					Name:        "next_hop_type",
					Description: `Next hop type of the routing rule.`,
				},
				resource.Attribute{
					Name:        "route_rule_id",
					Description: `ID of the routing rule.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `Routing table ID of the routing rule.`,
				},
				resource.Attribute{
					Name:        "source_address",
					Description: `Source address of the routing rule.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "route_rules",
					Description: `Result of the routing rules.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the routing rule.`,
				},
				resource.Attribute{
					Name:        "destination_address",
					Description: `Destination address of the routing rule.`,
				},
				resource.Attribute{
					Name:        "next_hop_id",
					Description: `Next hop ID of the routing rule.`,
				},
				resource.Attribute{
					Name:        "next_hop_type",
					Description: `Next hop type of the routing rule.`,
				},
				resource.Attribute{
					Name:        "route_rule_id",
					Description: `ID of the routing rule.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `Routing table ID of the routing rule.`,
				},
				resource.Attribute{
					Name:        "source_address",
					Description: `Source address of the routing rule.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "baiducloud_security_group_rules",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query Security Group list.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Required) Security Group ID`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional, ForceNew) only support filter string/int/bool value`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Optional) Security Group attached instance ID`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional, ForceNew) Security Group search result output file`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) Security Group attached vpc id The ` + "`" + `filter` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) filter variable name`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) filter variable value list ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `Security Group rules`,
				},
				resource.Attribute{
					Name:        "dest_group_id",
					Description: `SecurityGroup rule's destination group id`,
				},
				resource.Attribute{
					Name:        "dest_ip",
					Description: `SecurityGroup rule's destination ip`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `SecurityGroup rule's direction`,
				},
				resource.Attribute{
					Name:        "ether_type",
					Description: `SecurityGroup rule's ether type`,
				},
				resource.Attribute{
					Name:        "port_range",
					Description: `SecurityGroup rule's port range`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `SecurityGroup rule's protocol`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `SecurityGroup rule's remark`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `SecurityGroup rule's security group id`,
				},
				resource.Attribute{
					Name:        "source_group_id",
					Description: `SecurityGroup rule's source group id`,
				},
				resource.Attribute{
					Name:        "source_ip",
					Description: `SecurityGroup rule's source ip`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "rules",
					Description: `Security Group rules`,
				},
				resource.Attribute{
					Name:        "dest_group_id",
					Description: `SecurityGroup rule's destination group id`,
				},
				resource.Attribute{
					Name:        "dest_ip",
					Description: `SecurityGroup rule's destination ip`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `SecurityGroup rule's direction`,
				},
				resource.Attribute{
					Name:        "ether_type",
					Description: `SecurityGroup rule's ether type`,
				},
				resource.Attribute{
					Name:        "port_range",
					Description: `SecurityGroup rule's port range`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `SecurityGroup rule's protocol`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `SecurityGroup rule's remark`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `SecurityGroup rule's security group id`,
				},
				resource.Attribute{
					Name:        "source_group_id",
					Description: `SecurityGroup rule's source group id`,
				},
				resource.Attribute{
					Name:        "source_ip",
					Description: `SecurityGroup rule's source ip`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "baiducloud_security_groups",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query Security Group list.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional, ForceNew) only support filter string/int/bool value`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Optional) Security Group attached instance ID`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional, ForceNew) Security Group search result output file`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) Security Group attached vpc id The ` + "`" + `filter` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) filter variable name`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) filter variable value list ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `Security Groups search result`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Security Group description`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Security Group ID`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Security Group name`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `Security Group vpc id`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "security_groups",
					Description: `Security Groups search result`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Security Group description`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Security Group ID`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Security Group name`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `Security Group vpc id`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "baiducloud_snapshots",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query Snapshot list.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional, ForceNew) only support filter string/int/bool value`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional, ForceNew) Snapshots search result output file.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `(Optional) Volume ID to be attached of snapshots, if volume is system disk, volume ID is instance ID The ` + "`" + `filter` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) filter variable name`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) filter variable value list ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "snapshots",
					Description: `The result of the snapshots list.`,
				},
				resource.Attribute{
					Name:        "create_method",
					Description: `The creation method of the snapshot.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The creation time of the snapshot.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the snapshot.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the snapshot.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the snapshot.`,
				},
				resource.Attribute{
					Name:        "size_in_gb",
					Description: `The size of the snapshot in GB.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the snapshot.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The volume id of the snapshot.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "snapshots",
					Description: `The result of the snapshots list.`,
				},
				resource.Attribute{
					Name:        "create_method",
					Description: `The creation method of the snapshot.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The creation time of the snapshot.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the snapshot.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the snapshot.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the snapshot.`,
				},
				resource.Attribute{
					Name:        "size_in_gb",
					Description: `The size of the snapshot in GB.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the snapshot.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The volume id of the snapshot.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "baiducloud_specs",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query spec list.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cpu_count",
					Description: `(Optional, ForceNew) Useful cpu count of the search spec`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional, ForceNew) only support filter string/int/bool value`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Optional, ForceNew) Instance type of the search spec`,
				},
				resource.Attribute{
					Name:        "memory_size_in_gb",
					Description: `(Optional, ForceNew) Useful memory size in GB of the search spec`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional, ForceNew) Regex pattern of the search spec name`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional, ForceNew) Output file for saving result. The ` + "`" + `filter` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) filter variable name`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) filter variable value list ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "specs",
					Description: `Useful spec list, when create a bcc instance, suggest use instance_type/cpu_count/memory_capacity_in_gb as bcc instance parameters`,
				},
				resource.Attribute{
					Name:        "cpu_count",
					Description: `Useful cpu count`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `Useful instance type`,
				},
				resource.Attribute{
					Name:        "local_disk_size_in_gb",
					Description: `Useful local disk size in GB`,
				},
				resource.Attribute{
					Name:        "memory_size_in_gb",
					Description: `Useful memory size in GB`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Spec name`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "specs",
					Description: `Useful spec list, when create a bcc instance, suggest use instance_type/cpu_count/memory_capacity_in_gb as bcc instance parameters`,
				},
				resource.Attribute{
					Name:        "cpu_count",
					Description: `Useful cpu count`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `Useful instance type`,
				},
				resource.Attribute{
					Name:        "local_disk_size_in_gb",
					Description: `Useful local disk size in GB`,
				},
				resource.Attribute{
					Name:        "memory_size_in_gb",
					Description: `Useful memory size in GB`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Spec name`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "baiducloud_subnets",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query subnet list.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional, ForceNew) only support filter string/int/bool value`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional, ForceNew) Output file for saving result.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional, ForceNew) ID of the subnet.`,
				},
				resource.Attribute{
					Name:        "subnet_type",
					Description: `(Optional, ForceNew) Specify the subnet type for subnets.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional, ForceNew) VPC ID for subnets to retrieve.`,
				},
				resource.Attribute{
					Name:        "zone_name",
					Description: `(Optional, ForceNew) Specify the zone name for subnets. The ` + "`" + `filter` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) filter variable name`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) filter variable value list ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `Result of the subnets.`,
				},
				resource.Attribute{
					Name:        "available_ip",
					Description: `Available IP address of the subnet.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `CIDR block of the subnet.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the subnet.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the subnet.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of the subnet.`,
				},
				resource.Attribute{
					Name:        "subnet_type",
					Description: `Type of the subnet.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC ID of the subnet.`,
				},
				resource.Attribute{
					Name:        "zone_name",
					Description: `Zone name of the subnet.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "subnets",
					Description: `Result of the subnets.`,
				},
				resource.Attribute{
					Name:        "available_ip",
					Description: `Available IP address of the subnet.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `CIDR block of the subnet.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the subnet.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the subnet.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of the subnet.`,
				},
				resource.Attribute{
					Name:        "subnet_type",
					Description: `Type of the subnet.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC ID of the subnet.`,
				},
				resource.Attribute{
					Name:        "zone_name",
					Description: `Zone name of the subnet.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "baiducloud_vpcs",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query vpc list.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional, ForceNew) only support filter string/int/bool value`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, ForceNew) Name of the specific VPC to retrieve.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional, ForceNew) Output file for saving result.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional, ForceNew) ID of the specific VPC to retrieve. The ` + "`" + `filter` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) filter variable name`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) filter variable value list ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "vpcs",
					Description: `Result of VPCs.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `CIDR block of the VPC.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the VPC.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Specify if it is the default VPC.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the VPC.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `Route table ID of the VPC.`,
				},
				resource.Attribute{
					Name:        "secondary_cidrs",
					Description: `The secondary cidr list of the VPC. They will not be repeated.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vpcs",
					Description: `Result of VPCs.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `CIDR block of the VPC.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the VPC.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Specify if it is the default VPC.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the VPC.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `Route table ID of the VPC.`,
				},
				resource.Attribute{
					Name:        "secondary_cidrs",
					Description: `The secondary cidr list of the VPC. They will not be repeated.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "baiducloud_zones",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query zone list.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional, ForceNew) only support filter string/int/bool value`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional, ForceNew) Regex pattern of the search zone name`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional, ForceNew) Output file for saving result. The ` + "`" + `filter` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) filter variable name`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) filter variable value list ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `Useful zone list`,
				},
				resource.Attribute{
					Name:        "zone_name",
					Description: `Useful zone name`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "zones",
					Description: `Useful zone list`,
				},
				resource.Attribute{
					Name:        "zone_name",
					Description: `Useful zone name`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"baiducloud_acls":                   0,
		"baiducloud_appblb_listeners":       1,
		"baiducloud_appblb_server_groups":   2,
		"baiducloud_appblbs":                3,
		"baiducloud_auto_snapshot_policies": 4,
		"baiducloud_bos_bucket_objects":     5,
		"baiducloud_bos_buckets":            6,
		"baiducloud_cdss":                   7,
		"baiducloud_certs":                  8,
		"baiducloud_cfc_function":           9,
		"baiducloud_eips":                   10,
		"baiducloud_images":                 11,
		"baiducloud_instances":              12,
		"baiducloud_nat_gateways":           13,
		"baiducloud_peer_conns":             14,
		"baiducloud_route_rules":            15,
		"baiducloud_security_group_rules":   16,
		"baiducloud_security_groups":        17,
		"baiducloud_snapshots":              18,
		"baiducloud_specs":                  19,
		"baiducloud_subnets":                20,
		"baiducloud_vpcs":                   21,
		"baiducloud_zones":                  22,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
