package huaweicloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_api_gateway_api",
			Category:         "API Gateway (APIG)",
			ShortDescription: `Provides an API gateway API resource.`,
			Description:      ``,
			Keywords: []string{
				"api",
				"gateway",
				"apig",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the API. An API name consists of 3–64 characters, starting with a letter. Only letters, digits, and underscores (_) are allowed.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required) Specifies the ID of the API group. Changing this creates a new resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Specifies the description of the API. The description cannot exceed 255 characters.`,
				},
				resource.Attribute{
					Name:        "visibility",
					Description: `(Optional) Specifies whether the API is available to the public. The value can be 1 (public) and 2 (private). Defaults to 2.`,
				},
				resource.Attribute{
					Name:        "auth_type",
					Description: `(Required) Specifies the security authentication mode. The value can be 'App', 'IAM', and 'NONE'.`,
				},
				resource.Attribute{
					Name:        "request_protocol",
					Description: `(Optional) Specifies the request protocol. The value can be 'HTTP', 'HTTPS', and 'BOTH' which means the API can be accessed through both 'HTTP' and 'HTTPS'. Defaults to 'HTTPS'.`,
				},
				resource.Attribute{
					Name:        "request_method",
					Description: `(Required) Specifies the request method, including 'GET','POST','PUT' and etc..`,
				},
				resource.Attribute{
					Name:        "request_uri",
					Description: `(Required) Specifies the request path of the API. The value must comply with URI specifications.`,
				},
				resource.Attribute{
					Name:        "backend_type",
					Description: `(Required) Specifies the service backend type. The value can be: - 'HTTP': the web service backend - 'FUNCTION': the FunctionGraph service backend - 'MOCK': the Mock service backend`,
				},
				resource.Attribute{
					Name:        "http_backend",
					Description: `(Optional) Specifies the configuration when backend_type selected 'HTTP' (documented below).`,
				},
				resource.Attribute{
					Name:        "function_backend",
					Description: `(Optional) Specifies the configuration when backend_type selected 'FUNCTION' (documented below).`,
				},
				resource.Attribute{
					Name:        "mock_backend",
					Description: `(Optional) Specifies the configuration when backend_type selected 'MOCK' (documented below).`,
				},
				resource.Attribute{
					Name:        "request_parameter",
					Description: `(Optional) the request parameter list (documented below).`,
				},
				resource.Attribute{
					Name:        "backend_parameter",
					Description: `(Optional) the backend parameter list (documented below).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) the tags of API in format of string list.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) Specifies the version of the API. A maximum of 16 characters are allowed.`,
				},
				resource.Attribute{
					Name:        "cors",
					Description: `(Optional) Specifies whether CORS is supported or not.`,
				},
				resource.Attribute{
					Name:        "example_success_response",
					Description: `(Required) Specifies the example response for a successful request. The length cannot exceed 20,480 characters.`,
				},
				resource.Attribute{
					Name:        "example_failure_response",
					Description: `(Optional) Specifies the example response for a failed request The length cannot exceed 20,480 characters. The ` + "`" + `http_backend` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Specifies the backend request protocol. The value can be 'HTTP' and 'HTTPS'.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Optional) Specifies the backend request method, including 'GET','POST','PUT' and etc..`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `(Required) Specifies the backend request path. The value must comply with URI specifications.`,
				},
				resource.Attribute{
					Name:        "vpc_channel",
					Description: `(Optional) Specifies the VPC channel ID. This parameter and ` + "`" + `url_domain` + "`" + ` are alternative.`,
				},
				resource.Attribute{
					Name:        "url_domain",
					Description: `(Optional) Specifies the backend service address. An endpoint URL is in the format of "domain name (or IP address):port number", with up to 255 characters. This parameter and ` + "`" + `vpc_channel` + "`" + ` are alternative.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) Timeout duration (in ms) for API Gateway to request for the backend service. Defaults to 50000. The ` + "`" + `function_backend` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "function_urn",
					Description: `(Required) Specifies the function URN.`,
				},
				resource.Attribute{
					Name:        "invocation_type",
					Description: `(Required) Specifies the invocation mode, which can be 'async' or 'sync'.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) Specifies the function version.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) Timeout duration (in ms) for API Gateway to request for FunctionGraph. Defaults to 50000. The ` + "`" + `mock_backend` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the input parameter name. A parameter name consists of 1–32 characters, starting with a letter. Only letters, digits, periods (.), hyphens (-), and underscores (_) are allowed.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Required) Specifies the input parameter location, which can be 'PATH', 'QUERY' or 'HEADER'.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Specifies the input parameter type, which can be 'STRING' or 'NUMBER'.`,
				},
				resource.Attribute{
					Name:        "required",
					Description: `(Optional) Specifies whether the parameter is mandatory or not.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `(Optional) Specifies the default value when the parameter is optional.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Specifies the description of the parameter. The description cannot exceed 255 characters. The ` + "`" + `backend_parameter` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the parameter name. A parameter name consists of 1–32 characters, starting with a letter. Only letters, digits, periods (.), hyphens (-), and underscores (_) are allowed.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Required) Specifies the parameter location, which can be 'PATH', 'QUERY' or 'HEADER'.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Specifies the parameter type, which can be 'REQUEST', 'CONSTANT', or 'SYSTEM'.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Specifies the parameter value, which is a string of not more than 255 characters. The value varies depending on the parameter type: - 'REQUEST': parameter name in ` + "`" + `request_parameter` + "`" + ` - 'CONSTANT': real value of the parameter - 'SYSTEM': gateway parameter name`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Specifies the description of the parameter. The description cannot exceed 255 characters. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the API.`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `The name of the API group to which the API belongs. ## Import API can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_api_gateway_api.api "774438a28a574ac8a496325d1bf51807" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the API.`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `The name of the API group to which the API belongs. ## Import API can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_api_gateway_api.api "774438a28a574ac8a496325d1bf51807" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_api_gateway_group",
			Category:         "API Gateway (APIG)",
			ShortDescription: `Provides an API gateway group resource.`,
			Description:      ``,
			Keywords: []string{
				"api",
				"gateway",
				"apig",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the API group. An API group name consists of 3–64 characters, starting with a letter. Only letters, digits, and underscores (_) are allowed.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Specifies the description of the API group. The description cannot exceed 255 characters. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the API group.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the API group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the API group.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the API group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_as_configuration_v1",
			Category:         "Auto Scaling",
			ShortDescription: `Manages a V1 AS Configuration resource within HuaweiCloud.`,
			Description:      ``,
			Keywords: []string{
				"auto",
				"scaling",
				"as",
				"configuration",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to create the AS configuration. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new AS configuration.`,
				},
				resource.Attribute{
					Name:        "scaling_configuration_name",
					Description: `(Required) The name of the AS configuration. The name can contain letters, digits, underscores(_), and hyphens(-), and cannot exceed 64 characters.`,
				},
				resource.Attribute{
					Name:        "instance_config",
					Description: `(Required) The information about instance configurations. The instance_config dictionary data structure is documented below. The ` + "`" + `instance_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Optional) When using the existing instance specifications as the template to create AS configurations, specify this argument. In this case, flavor, image, and disk arguments do not take effect. If the instance_id argument is not specified, flavor, image, and disk arguments are mandatory.`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `(Optional) The flavor ID.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Optional) The image ID.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `(Optional) The disk group information. System disks are mandatory and data disks are optional. The dick structure is described below.`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `(Required) The name of the SSH key pair used to log in to the instance.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional) The user data to provide when launching the instance. The file content must be encoded with Base64.`,
				},
				resource.Attribute{
					Name:        "personality",
					Description: `(Optional) Customize the personality of an instance by defining one or more files and their contents. The personality structure is described below.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `(Optional) The elastic IP address of the instance. The public_ip structure is described below.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) Metadata key/value pairs to make available from within the instance. The ` + "`" + `disk` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) The disk size. The unit is GB. The system disk size ranges from 40 to 32768, and the data disk size ranges from 10 to 32768.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `(Required) The disk type, which must be the same as the disk type available in the system. The options include ` + "`" + `SATA` + "`" + ` (common I/O disk type) and ` + "`" + `SSD` + "`" + ` (ultra-high I/O disk type).`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `(Required) Whether the disk is a system disk or a data disk. Option ` + "`" + `DATA` + "`" + ` indicates a data disk. option ` + "`" + `SYS` + "`" + ` indicates a system disk.`,
				},
				resource.Attribute{
					Name:        "kms_id",
					Description: `(Optional) The Encryption KMS ID of the data disk. The ` + "`" + `personality` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The absolute path of the destination file.`,
				},
				resource.Attribute{
					Name:        "contents",
					Description: `(Required) The content of the injected file, which must be encoded with base64. The ` + "`" + `public_ip` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "eip",
					Description: `(Required) The configuration parameter for creating an elastic IP address that will be automatically assigned to the instance. The eip structure is described below. The ` + "`" + `eip` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "ip_type",
					Description: `(Required) The IP address type. The system only supports ` + "`" + `5_bgp` + "`" + ` (indicates dynamic BGP).`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Required) The bandwidth information. The structure is described below. The ` + "`" + `bandwidth` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) The bandwidth (Mbit/s). The value range is 1 to 300.`,
				},
				resource.Attribute{
					Name:        "share_type",
					Description: `(Required) The bandwidth sharing type. The system only supports ` + "`" + `PER` + "`" + ` (indicates exclusive bandwidth).`,
				},
				resource.Attribute{
					Name:        "charging_mode",
					Description: `(Required) The bandwidth charging mode. The system only supports ` + "`" + `traffic` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_as_group_v1",
			Category:         "Auto Scaling",
			ShortDescription: `Manages a V1 Autoscaling Group resource within HuaweiCloud.`,
			Description:      ``,
			Keywords: []string{
				"auto",
				"scaling",
				"as",
				"group",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to create the AS group. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new AS group.`,
				},
				resource.Attribute{
					Name:        "scaling_group_name",
					Description: `(Required) The name of the scaling group. The name can contain letters, digits, underscores(_), and hyphens(-),and cannot exceed 64 characters.`,
				},
				resource.Attribute{
					Name:        "scaling_configuration_id",
					Description: `(Optional) The configuration ID which defines configurations of instances in the AS group.`,
				},
				resource.Attribute{
					Name:        "desire_instance_number",
					Description: `(Optional) The expected number of instances. The default value is the minimum number of instances. The value ranges from the minimum number of instances to the maximum number of instances.`,
				},
				resource.Attribute{
					Name:        "min_instance_number",
					Description: `(Optional) The minimum number of instances. The default value is 0.`,
				},
				resource.Attribute{
					Name:        "max_instance_number",
					Description: `(Optional) The maximum number of instances. The default value is 0.`,
				},
				resource.Attribute{
					Name:        "cool_down_time",
					Description: `(Optional) The cooling duration (in seconds). The value ranges from 0 to 86400, and is 300 by default.`,
				},
				resource.Attribute{
					Name:        "lb_listener_id",
					Description: `(Optional) The ELB listener IDs. The system supports up to three ELB listeners, the IDs of which are separated using a comma (,).`,
				},
				resource.Attribute{
					Name:        "lbaas_listeners",
					Description: `(Optional) An array of one or more enhanced load balancer. The system supports the binding of up to three load balancers. The field is alternative to lb_listener_id. The lbaas_listeners object structure is documented below.`,
				},
				resource.Attribute{
					Name:        "available_zones",
					Description: `(Optional) The availability zones in which to create the instances in the autoscaling group.`,
				},
				resource.Attribute{
					Name:        "networks",
					Description: `(Required) An array of one or more network IDs. The system supports up to five networks. The networks object structure is documented below.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Required) An array of one or more security group IDs to associate with the group. The security_groups object structure is documented below.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) The VPC ID. Changing this creates a new group.`,
				},
				resource.Attribute{
					Name:        "health_periodic_audit_method",
					Description: `(Optional) The health check method for instances in the AS group. The health check methods include ` + "`" + `ELB_AUDIT` + "`" + ` and ` + "`" + `NOVA_AUDIT` + "`" + `. If load balancing is configured, the default value of this parameter is ` + "`" + `ELB_AUDIT` + "`" + `. Otherwise, the default value is ` + "`" + `NOVA_AUDIT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_periodic_audit_time",
					Description: `(Optional) The health check period for instances. The period has four options: 5 minutes (default), 15 minutes, 60 minutes, and 180 minutes.`,
				},
				resource.Attribute{
					Name:        "instance_terminate_policy",
					Description: `(Optional) The instance removal policy. The policy has four options: ` + "`" + `OLD_CONFIG_OLD_INSTANCE` + "`" + ` (default), ` + "`" + `OLD_CONFIG_NEW_INSTANCE` + "`" + `, ` + "`" + `OLD_INSTANCE` + "`" + `, and ` + "`" + `NEW_INSTANCE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value pairs to associate with the scaling group.`,
				},
				resource.Attribute{
					Name:        "notifications",
					Description: `(Optional) The notification mode. The system only supports ` + "`" + `EMAIL` + "`" + ` mode which refers to notification by email.`,
				},
				resource.Attribute{
					Name:        "delete_publicip",
					Description: `(Optional) Whether to delete the elastic IP address bound to the instances of AS group when deleting the instances. The options are ` + "`" + `true` + "`" + ` and ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "delete_instances",
					Description: `(Optional) Whether to delete the instances in the AS group when deleting the AS group. The options are ` + "`" + `yes` + "`" + ` and ` + "`" + `no` + "`" + `. The ` + "`" + `networks` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The network UUID. The ` + "`" + `security_groups` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The UUID of the security group. The ` + "`" + `lbaas_listeners` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "pool_id",
					Description: `(Required) Specifies the backend ECS group ID.`,
				},
				resource.Attribute{
					Name:        "protocol_port",
					Description: `(Required) Specifies the backend protocol, which is the port on which a backend ECS listens for traffic. The number of the port ranges from 1 to 65535.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) Specifies the weight, which determines the portion of requests a backend ECS processes compared to other backend ECSs added to the same listener. The value of this parameter ranges from 0 to 100. The default value is 1. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "scaling_group_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "desire_instance_number",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "min_instance_number",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "max_instance_number",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cool_down_time",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "lb_listener_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "health_periodic_audit_method",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "health_periodic_audit_time",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "instance_terminate_policy",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "scaling_configuration_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "delete_publicip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "notifications",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `The instances IDs of the AS group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "scaling_group_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "desire_instance_number",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "min_instance_number",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "max_instance_number",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cool_down_time",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "lb_listener_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "health_periodic_audit_method",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "health_periodic_audit_time",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "instance_terminate_policy",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "scaling_configuration_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "delete_publicip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "notifications",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `The instances IDs of the AS group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_as_policy_v1",
			Category:         "Auto Scaling",
			ShortDescription: `Manages a V1 AS Policy resource within HuaweiCloud.`,
			Description:      ``,
			Keywords: []string{
				"auto",
				"scaling",
				"as",
				"policy",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to create the AS policy. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new AS policy.`,
				},
				resource.Attribute{
					Name:        "scaling_policy_name",
					Description: `(Required) The name of the AS policy. The name can contain letters, digits, underscores(_), and hyphens(-),and cannot exceed 64 characters.`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `(Required) The AS group ID. Changing this creates a new AS policy.`,
				},
				resource.Attribute{
					Name:        "scaling_policy_type",
					Description: `(Required) The AS policy type. The values can be ` + "`" + `ALARM` + "`" + `, ` + "`" + `SCHEDULED` + "`" + `, and ` + "`" + `RECURRENCE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "alarm_id",
					Description: `(Optional) The alarm rule ID. This argument is mandatory when ` + "`" + `scaling_policy_type` + "`" + ` is set to ` + "`" + `ALARM` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scheduled_policy",
					Description: `(Optional) The periodic or scheduled AS policy. This argument is mandatory when ` + "`" + `scaling_policy_type` + "`" + ` is set to ` + "`" + `SCHEDULED` + "`" + ` or ` + "`" + `RECURRENCE` + "`" + `. The scheduled_policy structure is documented below.`,
				},
				resource.Attribute{
					Name:        "scaling_policy_action",
					Description: `(Optional) The action of the AS policy. The scaling_policy_action structure is documented below.`,
				},
				resource.Attribute{
					Name:        "cool_down_time",
					Description: `(Optional) The cooling duration (in seconds), and is 900 by default. The ` + "`" + `scheduled_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "launch_time",
					Description: `(Required) The time when the scaling action is triggered. If ` + "`" + `scaling_policy_type` + "`" + ` is set to ` + "`" + `SCHEDULED` + "`" + `, the time format is YYYY-MM-DDThh:mmZ. If ` + "`" + `scaling_policy_type` + "`" + ` is set to ` + "`" + `RECURRENCE` + "`" + `, the time format is hh:mm.`,
				},
				resource.Attribute{
					Name:        "recurrence_type",
					Description: `(Optional) The periodic triggering type. This argument is mandatory when ` + "`" + `scaling_policy_type` + "`" + ` is set to ` + "`" + `RECURRENCE` + "`" + `. The options include ` + "`" + `Daily` + "`" + `, ` + "`" + `Weekly` + "`" + `, and ` + "`" + `Monthly` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "recurrence_value",
					Description: `(Optional) The frequency at which scaling actions are triggered.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `(Optional) The start time of the scaling action triggered periodically. The time format complies with UTC. The current time is used by default. The time format is YYYY-MM-DDThh:mmZ.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `(Optional) The end time of the scaling action triggered periodically. The time format complies with UTC. This argument is mandatory when ` + "`" + `scaling_policy_type` + "`" + ` is set to ` + "`" + `RECURRENCE` + "`" + `. The time format is YYYY-MM-DDThh:mmZ. The ` + "`" + `scaling_policy_action` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "operation",
					Description: `(Optional) The operation to be performed. The options include ` + "`" + `ADD` + "`" + ` (default), ` + "`" + `REMOVE` + "`" + `, and ` + "`" + `SET` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_number",
					Description: `(Optional) The number of instances to be operated. The default number is 1. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "scaling_policy_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "scaling_policy_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "alarm_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cool_down_time",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "scaling_policy_action/operation",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "scaling_policy_action/instance_number",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "scheduled_policy/launch_time",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "scheduled_policy/recurrence_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "scheduled_policy/recurrence_value",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "scheduled_policy/start_time",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "scheduled_policy/end_time",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "scaling_policy_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "scaling_policy_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "alarm_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cool_down_time",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "scaling_policy_action/operation",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "scaling_policy_action/instance_number",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "scheduled_policy/launch_time",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "scheduled_policy/recurrence_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "scheduled_policy/recurrence_value",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "scheduled_policy/start_time",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "scheduled_policy/end_time",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_blockstorage_volume_v2",
			Category:         "Elastic Volume Service (EVS)",
			ShortDescription: `Manages a V2 volume resource within HuaweiCloud.`,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"volume",
				"service",
				"evs",
				"blockstorage",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to create the volume. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new volume.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) The size of the volume to create (in gigabytes).`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) The availability zone for the volume. Changing this creates a new volume.`,
				},
				resource.Attribute{
					Name:        "consistency_group_id",
					Description: `(Optional) The consistency group to place the volume in.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the volume. Changing this updates the volume's description.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Optional) The image ID from which to create the volume. Changing this creates a new volume.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) Metadata key/value pairs to associate with the volume. Changing this updates the existing volume metadata.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A unique name for the volume. Changing this updates the volume's name.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) The snapshot ID from which to create the volume. Changing this creates a new volume.`,
				},
				resource.Attribute{
					Name:        "source_replica",
					Description: `(Optional) The volume ID to replicate with.`,
				},
				resource.Attribute{
					Name:        "source_vol_id",
					Description: `(Optional) The volume ID from which to create the volume. Changing this creates a new volume.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `(Optional) The type of volume to create. Available types are ` + "`" + `SSD` + "`" + `, ` + "`" + `SAS` + "`" + ` and ` + "`" + `SATA` + "`" + `. Changing this creates a new volume.`,
				},
				resource.Attribute{
					Name:        "cascade",
					Description: `(Optional, Default:false) Specifies to delete all snapshots associated with the EVS disk. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "source_vol_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "attachment",
					Description: `If a volume is attached to an instance, this attribute will display the Attachment ID, Instance ID, and the Device as the Instance sees it. ## Import Volumes can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_blockstorage_volume_v2.volume_1 ea257959-eeb1-4c10-8d33-26f0409a755d ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "source_vol_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "attachment",
					Description: `If a volume is attached to an instance, this attribute will display the Attachment ID, Instance ID, and the Device as the Instance sees it. ## Import Volumes can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_blockstorage_volume_v2.volume_1 ea257959-eeb1-4c10-8d33-26f0409a755d ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_cce_cluster_v3",
			Category:         "Cloud Container Engine (CCE)",
			ShortDescription: `Provides Cloud Container Engine(CCE) resource.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"container",
				"engine",
				"cce",
				"cluster",
				"v3",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Cluster name. Changing this parameter will create a new cluster resource.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Cluster tag, key/value pair format. Changing this parameter will create a new cluster resource.`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional) Cluster annotation, key/value pair format. Changing this parameter will create a new cluster resource.`,
				},
				resource.Attribute{
					Name:        "flavor_id",
					Description: `(Required) Cluster specifications. Changing this parameter will create a new cluster resource. Possible values:`,
				},
				resource.Attribute{
					Name:        "cce.s1.small",
					Description: `small-scale single cluster (up to 50 nodes).`,
				},
				resource.Attribute{
					Name:        "cce.s1.medium",
					Description: `medium-scale single cluster (up to 200 nodes).`,
				},
				resource.Attribute{
					Name:        "cce.s1.large",
					Description: `large-scale single cluster (up to 1000 nodes).`,
				},
				resource.Attribute{
					Name:        "cce.s2.small",
					Description: `small-scale HA cluster (up to 50 nodes).`,
				},
				resource.Attribute{
					Name:        "cce.s2.medium",
					Description: `medium-scale HA cluster (up to 200 nodes).`,
				},
				resource.Attribute{
					Name:        "cce.s2.large",
					Description: `large-scale HA cluster (up to 1000 nodes).`,
				},
				resource.Attribute{
					Name:        "cce.t1.small",
					Description: `small-scale single physical machine cluster (up to 10 nodes).`,
				},
				resource.Attribute{
					Name:        "cce.t1.medium",
					Description: `medium-scale single physical machine cluster (up to 100 nodes).`,
				},
				resource.Attribute{
					Name:        "cce.t1.large",
					Description: `large-scale single physical machine cluster (up to 500 nodes).`,
				},
				resource.Attribute{
					Name:        "cce.t2.small",
					Description: `small-scale HA physical machine cluster (up to 10 nodes).`,
				},
				resource.Attribute{
					Name:        "cce.t2.medium",
					Description: `medium-scale HA physical machine cluster (up to 100 nodes).`,
				},
				resource.Attribute{
					Name:        "cce.t2.large",
					Description: `large-scale HA physical machine cluster (up to 500 nodes).`,
				},
				resource.Attribute{
					Name:        "cluster_version",
					Description: `(Optional) For the cluster version, possible values are v1.7.3-r10 or v1.9.2-r1.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `(Required) Cluster Type, possible values are VirtualMachine and BareMetal. Changing this parameter will create a new cluster resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Cluster description.`,
				},
				resource.Attribute{
					Name:        "billing_mode",
					Description: `(Optional) Charging mode of the cluster, which is 0 (on demand). Changing this parameter will create a new cluster resource.`,
				},
				resource.Attribute{
					Name:        "extend_param",
					Description: `(Optional) Extended parameter. Changing this parameter will create a new cluster resource.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) The ID of the VPC used to create the node. Changing this parameter will create a new cluster resource.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) The ID of the subnet used to create the node. Changing this parameter will create a new cluster resource.`,
				},
				resource.Attribute{
					Name:        "highway_subnet_id",
					Description: `(Optional) The ID of the high speed network used to create bare metal nodes. Changing this parameter will create a new cluster resource.`,
				},
				resource.Attribute{
					Name:        "container_network_type",
					Description: `(Required) Container network parameters. Possible values:`,
				},
				resource.Attribute{
					Name:        "overlay_l2",
					Description: `An overlay_l2 network built for containers by using Open vSwitch(OVS)`,
				},
				resource.Attribute{
					Name:        "underlay_ipvlan",
					Description: `An underlay_ipvlan network built for bare metal servers by using ipvlan.`,
				},
				resource.Attribute{
					Name:        "vpc-router",
					Description: `An vpc-router network built for containers by using ipvlan and custom VPC routes.`,
				},
				resource.Attribute{
					Name:        "container_network_cidr",
					Description: `(Optional) Container network segment. Changing this parameter will create a new cluster resource.`,
				},
				resource.Attribute{
					Name:        "authentication_mode",
					Description: `(Optional) Authentication mode of the cluster, possible values are x509 and rbac. Defaults to x509. Changing this parameter will create a new cluster resource.`,
				},
				resource.Attribute{
					Name:        "authenticating_proxy_ca",
					Description: `(Optional) CA root certificate provided in the authenticating_proxy mode. The CA root certificate is encoded to the Base64 format. Changing this parameter will create a new cluster resource.`,
				},
				resource.Attribute{
					Name:        "multi_az",
					Description: `(Optional) Enable multiple AZs for the cluster, only when using HA flavors. Changing this parameter will create a new cluster resource.`,
				},
				resource.Attribute{
					Name:        "eip",
					Description: `(Optional) EIP address of the cluster. Changing this parameter will create a new cluster resource. ## Attributes Reference All above argument parameters can be exported as attribute parameters along with attribute reference.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of the cluster resource.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Cluster status information.`,
				},
				resource.Attribute{
					Name:        "certificate_clusters/name",
					Description: `The cluster name.`,
				},
				resource.Attribute{
					Name:        "certificate_clusters/server",
					Description: `The server IP address.`,
				},
				resource.Attribute{
					Name:        "certificate_clusters/certificate_authority_data",
					Description: `The certificate data.`,
				},
				resource.Attribute{
					Name:        "certificate_users/name",
					Description: `The user name.`,
				},
				resource.Attribute{
					Name:        "certificate_users/client_certificate_data",
					Description: `The client certificate data.`,
				},
				resource.Attribute{
					Name:        "certificate_users/client_key_data",
					Description: `The client key data. ## Import Cluster can be imported using the cluster id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_cce_cluster_v3.cluster_1 4779ab1c-7c1a-44b1-a02e-93dfc361b32d ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Id of the cluster resource.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Cluster status information.`,
				},
				resource.Attribute{
					Name:        "certificate_clusters/name",
					Description: `The cluster name.`,
				},
				resource.Attribute{
					Name:        "certificate_clusters/server",
					Description: `The server IP address.`,
				},
				resource.Attribute{
					Name:        "certificate_clusters/certificate_authority_data",
					Description: `The certificate data.`,
				},
				resource.Attribute{
					Name:        "certificate_users/name",
					Description: `The user name.`,
				},
				resource.Attribute{
					Name:        "certificate_users/client_certificate_data",
					Description: `The client certificate data.`,
				},
				resource.Attribute{
					Name:        "certificate_users/client_key_data",
					Description: `The client key data. ## Import Cluster can be imported using the cluster id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_cce_cluster_v3.cluster_1 4779ab1c-7c1a-44b1-a02e-93dfc361b32d ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_cce_nodes_v3",
			Category:         "Cloud Container Engine (CCE)",
			ShortDescription: `Add a node to a container cluster.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"container",
				"engine",
				"cce",
				"nodes",
				"v3",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) ID of the cluster. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "billing_mode",
					Description: `(Optional) Node's billing mode: The value is 0 (on demand). Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Node Name.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Node tag, key/value pair format. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional) Node annotation, key/value pair format. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "flavor_id",
					Description: `(Required) Specifies the flavor id. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required) specify the name of the available partition (AZ). Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "os",
					Description: `(Optional) Operating System of the node, possible values are EulerOS 2.2 and CentOS 7.1. Defaults to EulerOS 2.2. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "key_pair",
					Description: `(Optional) Key pair name when logging in to select the key pair mode. This parameter and ` + "`" + `password` + "`" + ` are alternative. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) root password when logging in to select the password mode. This parameter must be salted and alternative to ` + "`" + `key_pair` + "`" + `. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "eip_ids",
					Description: `(Optional) List of existing elastic IP IDs. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "eip_count",
					Description: `(Optional) Number of elastic IPs to be dynamically created. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "iptype",
					Description: `(Optional) Elastic IP type. Default is 5_bgp. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "bandwidth_charge_mode",
					Description: `(Optional) Bandwidth billing type. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "sharetype",
					Description: `(Optional) Bandwidth sharing type. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "bandwidth_size",
					Description: `(Optional) Bandwidth size. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "extend_param_charging_mode",
					Description: `(Optional) Node charging mode, 0 is on-demand charging. Changing this parameter will create a new cluster resource.`,
				},
				resource.Attribute{
					Name:        "ecs_performance_type",
					Description: `(Optional) Classification of cloud server specifications. Changing this parameter will create a new cluster resource.`,
				},
				resource.Attribute{
					Name:        "order_id",
					Description: `(Optional) Order ID, mandatory when the node payment type is the automatic payment package period type. Changing this parameter will create a new cluster resource.`,
				},
				resource.Attribute{
					Name:        "product_id",
					Description: `(Optional) The Product ID. Changing this parameter will create a new cluster resource.`,
				},
				resource.Attribute{
					Name:        "max_pods",
					Description: `(Optional) The maximum number of instances a node is allowed to create. Changing this parameter will create a new cluster resource.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(Optional) The Public key. Changing this parameter will create a new cluster resource.`,
				},
				resource.Attribute{
					Name:        "preinstall",
					Description: `(Optional) Script required before installation. The input value can be a Base64 encoded string or not. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "postinstall",
					Description: `(Optional) Script required after installation. The input value can be a Base64 encoded string or not. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) The ID of the subnet to which the NIC belongs. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) Disk size in GB.`,
				},
				resource.Attribute{
					Name:        "volumetype",
					Description: `(Required) Disk type.`,
				},
				resource.Attribute{
					Name:        "extend_param",
					Description: `(Optional) Disk expansion parameters.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) Disk size in GB.`,
				},
				resource.Attribute{
					Name:        "volumetype",
					Description: `(Required) Disk type.`,
				},
				resource.Attribute{
					Name:        "extend_param",
					Description: `(Optional) Disk expansion parameters. ## Attributes Reference All above argument parameters can be exported as attribute parameters along with attribute reference.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Node status information.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Private IP of the CCE node.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP of the CCE node.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `Node status information.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Private IP of the CCE node.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP of the CCE node.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_cci_network_v1",
			Category:         "Cloud Container Instance (CCI)",
			ShortDescription: `Provides Cloud Container Instance(CCI) resource.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"container",
				"instance",
				"cci",
				"network",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) CCI Network name. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Required) CCI Network namespace. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "security_group",
					Description: `(Required) ID of the security group to which the subnet of the network belongs. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) Project ID of the tenant. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Required) Domain ID of the tenant. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) ID of the VPC to which the network belongs. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required) Network ID of the VPC subnet in which the network belongs. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) ID of the VPC subnet to which the network belongs. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "available_zone",
					Description: `(Required) AZ to which the VPC subnet of the network belongs. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Required) Network segment of the VPC subnet to which the network belongs. Changing this parameter will create a new resource. ## Attributes Reference All above argument parameters can be exported as attribute parameters along with attribute reference.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of the instance resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Id of the instance resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_cdm_cluster_v1",
			Category:         "Cloud Data Migration (CDM)",
			ShortDescription: `cdm cluster management`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"data",
				"migration",
				"cdm",
				"cluster",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required) Available zone. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "flavor_id",
					Description: `(Required) Flavor id. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Cluster name. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Required) Security group ID. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) Subnet ID. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) Cluster version. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) VPC ID. Changing this parameter will create a new resource. - - -`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) Notification email addresses. The max number is 5. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "enterprise_project_id",
					Description: `(Optional) The enterprise project id. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "is_auto_off",
					Description: `(Optional) Whether to automatically shut down. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "phone_num",
					Description: `(Optional) Notification phone numbers. The max number is 5. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "schedule_boot_time",
					Description: `(Optional) Timed boot time. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "schedule_off_time",
					Description: `(Optional) Timed shutdown time. Changing this parameter will create a new resource. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `Create time.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `Instance list. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "publid_ip",
					Description: `Public ip. The ` + "`" + `instances` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Instance ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Instance name.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `Role.`,
				},
				resource.Attribute{
					Name:        "traffic_ip",
					Description: `Traffic IP.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Instance type. ## Timeouts This resource provides the following timeouts configuration options: - ` + "`" + `create` + "`" + ` - Default is 30 minute.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created",
					Description: `Create time.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `Instance list. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "publid_ip",
					Description: `Public ip. The ` + "`" + `instances` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Instance ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Instance name.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `Role.`,
				},
				resource.Attribute{
					Name:        "traffic_ip",
					Description: `Traffic IP.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Instance type. ## Timeouts This resource provides the following timeouts configuration options: - ` + "`" + `create` + "`" + ` - Default is 30 minute.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_cdn_domain_v1",
			Category:         "Content Delivery Network (CDN)",
			ShortDescription: `cdn domain management`,
			Description:      ``,
			Keywords: []string{
				"content",
				"delivery",
				"network",
				"cdn",
				"domain",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The acceleration domain name. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The service type. The valid values are 'web', 'download' and 'video'. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "sources",
					Description: `(Required) An array of one or more objects specifies the domain name of the origin server. The sources object structure is documented below.`,
				},
				resource.Attribute{
					Name:        "enterprise_project_id",
					Description: `(Optional) The enterprise project id. Changing this parameter will create a new resource. The ` + "`" + `sources` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "origin",
					Description: `(Required) The domain name or IP address of the origin server.`,
				},
				resource.Attribute{
					Name:        "origin_type",
					Description: `(Required) The origin server type. The valid values are 'ipaddr', 'domain', and 'obs_bucket'.`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `(Optional) Whether an origin server is active or standby (1: active; 0: standby). The default value is 1. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "enterprise_project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "sources/origin",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "sources/origin_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "sources/active",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The acceleration domain name ID.`,
				},
				resource.Attribute{
					Name:        "cname",
					Description: `The CNAME of the acceleration domain name.`,
				},
				resource.Attribute{
					Name:        "domain_status",
					Description: `The status of the acceleration domain name. The available values are 'online', 'offline', 'configuring', 'configure_failed', 'checking', 'check_failed' and 'deleting.'`,
				},
				resource.Attribute{
					Name:        "service_area",
					Description: `The area covered by the acceleration service. ## Import Domains can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_cdn_domain_v1.domain_1 fe2462fac09a4a42a76ecc4a1ef542f1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "enterprise_project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "sources/origin",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "sources/origin_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "sources/active",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The acceleration domain name ID.`,
				},
				resource.Attribute{
					Name:        "cname",
					Description: `The CNAME of the acceleration domain name.`,
				},
				resource.Attribute{
					Name:        "domain_status",
					Description: `The status of the acceleration domain name. The available values are 'online', 'offline', 'configuring', 'configure_failed', 'checking', 'check_failed' and 'deleting.'`,
				},
				resource.Attribute{
					Name:        "service_area",
					Description: `The area covered by the acceleration service. ## Import Domains can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_cdn_domain_v1.domain_1 fe2462fac09a4a42a76ecc4a1ef542f1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_ces-alarmrule",
			Category:         "Cloud Eye",
			ShortDescription: `Manages a V2 topic resource within HuaweiCloud.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"eye",
				"ces-alarmrule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "alarm_name",
					Description: `(Required) Specifies the name of an alarm rule. The value can be a string of 1 to 128 characters that can consist of numbers, lowercase letters, uppercase letters, underscores (_), or hyphens (-).`,
				},
				resource.Attribute{
					Name:        "alarm_description",
					Description: `(Optional) The value can be a string of 0 to 256 characters.`,
				},
				resource.Attribute{
					Name:        "metric",
					Description: `(Required) Specifies the alarm metrics. The structure is described below.`,
				},
				resource.Attribute{
					Name:        "condition",
					Description: `(Required) Specifies the alarm triggering condition. The structure is described below.`,
				},
				resource.Attribute{
					Name:        "alarm_actions",
					Description: `(Optional) Specifies the action triggered by an alarm. The structure is described below.`,
				},
				resource.Attribute{
					Name:        "insufficientdata_actions",
					Description: `(Optional) Specifies the action triggered by data insufficiency. The structure is described below.`,
				},
				resource.Attribute{
					Name:        "ok_actions",
					Description: `(Optional) Specifies the action triggered by the clearing of an alarm. The structure is described below.`,
				},
				resource.Attribute{
					Name:        "alarm_enabled",
					Description: `(Optional) Specifies whether to enable the alarm. The default value is true.`,
				},
				resource.Attribute{
					Name:        "alarm_action_enabled",
					Description: `(Optional) Specifies whether to enable the action to be triggered by an alarm. The default value is true. Note: If alarm_action_enabled is set to true, at least one of the following parameters alarm_actions, insufficientdata_actions, and ok_actions cannot be empty. If alarm_actions, insufficientdata_actions, and ok_actions coexist, their corresponding notification_list must be of the same value. The ` + "`" + `metric` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Required) Specifies the namespace in service.item format. service.item can be a string of 3 to 32 characters that must start with a letter and can consists of uppercase letters, lowercase letters, numbers, or underscores (_).`,
				},
				resource.Attribute{
					Name:        "metric_name",
					Description: `(Required) Specifies the metric name. The value can be a string of 1 to 64 characters that must start with a letter and can consists of uppercase letters, lowercase letters, numbers, or underscores (_).`,
				},
				resource.Attribute{
					Name:        "dimensions",
					Description: `(Required) Specifies the list of metric dimensions. Currently, the maximum length of the dimesion list that are supported is 3. The structure is described below. The ` + "`" + `dimensions` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the dimension name. The value can be a string of 1 to 32 characters that must start with a letter and can consists of uppercase letters, lowercase letters, numbers, underscores (_), or hyphens (-).`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Specifies the dimension value. The value can be a string of 1 to 64 characters that must start with a letter or a number and can consists of uppercase letters, lowercase letters, numbers, underscores (_), or hyphens (-). The ` + "`" + `condition` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Required) Specifies the alarm checking period in seconds. The value can be 1, 300, 1200, 3600, 14400, and 86400. Note: If period is set to 1, the raw metric data is used to determine whether to generate an alarm.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Required) Specifies the data rollup methods. The value can be max, min, average, sum, and vaiance.`,
				},
				resource.Attribute{
					Name:        "comparison_operator",
					Description: `(Required) Specifies the comparison condition of alarm thresholds. The value can be >, =, <, >=, or <=.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Specifies the alarm threshold. The value ranges from 0 to Number of 1.7976931348623157e+308.`,
				},
				resource.Attribute{
					Name:        "unit",
					Description: `(Optional) Specifies the data unit.`,
				},
				resource.Attribute{
					Name:        "count",
					Description: `(Required) Specifies the number of consecutive occurrence times. The value ranges from 1 to 5. the ` + "`" + `alarm_actions` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) specifies the type of action triggered by an alarm. the value can be notification or autoscaling. notification: indicates that a notification will be sent to the user. autoscaling: indicates that a scaling action will be triggered.`,
				},
				resource.Attribute{
					Name:        "notification_list",
					Description: `(Optional) specifies the topic urn list of the target notification objects. the maximum length is 5. the topic urn list can be obtained from simple message notification (smn) and in the following format: urn: smn:([a-z]|[a-z]|[0-9]|\-){1,32}:([a-z]|[a-z]|[0-9]){32}:([a-z]|[a-z]|[0-9]|\-|\_){1,256}. if type is set to notification, the value of notification_list cannot be empty. if type is set to autoscaling, the value of notification_list must be [] and the value of namespace must be sys.as. Note: to enable the as alarm rules take effect, you must bind scaling policies. for details, see the auto scaling api reference. the ` + "`" + `insufficientdata_actions` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) specifies the type of action triggered by an alarm. the value is notification. notification: indicates that a notification will be sent to the user.`,
				},
				resource.Attribute{
					Name:        "notification_list",
					Description: `(Optional) indicates the list of objects to be notified if the alarm status changes. the maximum length is 5. the ` + "`" + `ok_actions` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) specifies the type of action triggered by an alarm. the value is notification. notification: indicates that a notification will be sent to the user.`,
				},
				resource.Attribute{
					Name:        "notification_list",
					Description: `(Optional) indicates the list of objects to be notified if the alarm status changes. the maximum length is 5. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "alarm_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "alarm_description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "metric",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "condition",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "alarm_actions",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "insufficientdata_actions",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ok_actions",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "alarm_enabled",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "alarm_action_enabled",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Specifies the alarm rule ID.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Specifies the time when the alarm status changed. The value is a UNIX timestamp and the unit is ms.`,
				},
				resource.Attribute{
					Name:        "alarm_state",
					Description: `Specifies the alarm status. The value can be: ok: The alarm status is normal, alarm: An alarm is generated, insufficient_data: The required data is insufficient.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "alarm_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "alarm_description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "metric",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "condition",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "alarm_actions",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "insufficientdata_actions",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ok_actions",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "alarm_enabled",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "alarm_action_enabled",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Specifies the alarm rule ID.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Specifies the time when the alarm status changed. The value is a UNIX timestamp and the unit is ms.`,
				},
				resource.Attribute{
					Name:        "alarm_state",
					Description: `Specifies the alarm status. The value can be: ok: The alarm status is normal, alarm: An alarm is generated, insufficient_data: The required data is insufficient.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_cloudtable_cluster_v2",
			Category:         "Cloud Table",
			ShortDescription: `cloud table cluster management`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"table",
				"cloudtable",
				"cluster",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required) Availability zone (AZ). Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Cluster name. The value must be between 4 and 64 characters long and start with a letter. Only letters, digits, and hyphens (-) are allowed. It is case insensitive. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "rs_num",
					Description: `(Required) Number of computing units. Value range: 2 to 10. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Required) Security group ID. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `(Required) Storage I/O type. The value are ULTRAHIGH and COMMON. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) Subnet ID. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) VPC of the cluster. Changing this parameter will create a new resource. - - -`,
				},
				resource.Attribute{
					Name:        "enable_iam_auth",
					Description: `(Optional) Whether to enable IAM authentication for OpenTSDB. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "lemon_num",
					Description: `(Optional) Number of Lemon nodes Value range: 2 to 10. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "opentsdb_num",
					Description: `(Optional) Number of OpenTSDB nodes Value range: 2 to 10. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Enterprise project information. Structure is documented below. Changing this parameter will create a new resource. The ` + "`" + `tags` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) Tag value. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Tag key. Changing this parameter will create a new resource. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `Time when the cluster was created.`,
				},
				resource.Attribute{
					Name:        "hbase_public_endpoint",
					Description: `HBase link of the public network.`,
				},
				resource.Attribute{
					Name:        "lemon_link",
					Description: `Lemon link of the intranet.`,
				},
				resource.Attribute{
					Name:        "open_tsdb_link",
					Description: `OpenTSDB link of the intranet.`,
				},
				resource.Attribute{
					Name:        "opentsdb_public_endpoint",
					Description: `OpenTSDB link of the public network.`,
				},
				resource.Attribute{
					Name:        "storage_quota",
					Description: `Storage quota.`,
				},
				resource.Attribute{
					Name:        "used_storage_size",
					Description: `Used storage space.`,
				},
				resource.Attribute{
					Name:        "zookeeper_link",
					Description: `ZooKeeper link of the intranet. ## Timeouts This resource provides the following timeouts configuration options: - ` + "`" + `create` + "`" + ` - Default is 30 minute.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created",
					Description: `Time when the cluster was created.`,
				},
				resource.Attribute{
					Name:        "hbase_public_endpoint",
					Description: `HBase link of the public network.`,
				},
				resource.Attribute{
					Name:        "lemon_link",
					Description: `Lemon link of the intranet.`,
				},
				resource.Attribute{
					Name:        "open_tsdb_link",
					Description: `OpenTSDB link of the intranet.`,
				},
				resource.Attribute{
					Name:        "opentsdb_public_endpoint",
					Description: `OpenTSDB link of the public network.`,
				},
				resource.Attribute{
					Name:        "storage_quota",
					Description: `Storage quota.`,
				},
				resource.Attribute{
					Name:        "used_storage_size",
					Description: `Used storage space.`,
				},
				resource.Attribute{
					Name:        "zookeeper_link",
					Description: `ZooKeeper link of the intranet. ## Timeouts This resource provides the following timeouts configuration options: - ` + "`" + `create` + "`" + ` - Default is 30 minute.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_compute_floatingip_associate_v2",
			Category:         "Elastic Cloud Server (ECS)",
			ShortDescription: `Associate a floating IP to an instance`,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"cloud",
				"server",
				"ecs",
				"compute",
				"floatingip",
				"associate",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Compute client. Keypairs are associated with accounts, but a Compute client is needed to create one. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new floatingip_associate.`,
				},
				resource.Attribute{
					Name:        "floating_ip",
					Description: `(Required) The floating IP to associate.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) The instance to associte the floating IP with.`,
				},
				resource.Attribute{
					Name:        "fixed_ip",
					Description: `(Optional) The specific IP address to direct traffic to. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "floating_ip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "fixed_ip",
					Description: `See Argument Reference above. ## Import This resource can be imported by specifying all three arguments, separated by a forward slash: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_compute_floatingip_associate_v2.fip_1 <floating_ip>/<instance_id>/<fixed_ip> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "floating_ip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "fixed_ip",
					Description: `See Argument Reference above. ## Import This resource can be imported by specifying all three arguments, separated by a forward slash: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_compute_floatingip_associate_v2.fip_1 <floating_ip>/<instance_id>/<fixed_ip> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_compute_floatingip_v2",
			Category:         "Elastic Cloud Server (ECS)",
			ShortDescription: `Manages a V2 floating IP resource within HuaweiCloud Nova (compute).`,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"cloud",
				"server",
				"ecs",
				"compute",
				"floatingip",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Compute client. A Compute client is needed to create a floating IP that can be used with a compute instance. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new floating IP (which may or may not have a different address).`,
				},
				resource.Attribute{
					Name:        "pool",
					Description: `(Optional) The name of the pool from which to obtain the floating IP. Only admin_external_net is valid. Changing this creates a new floating IP. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "pool",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The actual floating IP address itself.`,
				},
				resource.Attribute{
					Name:        "fixed_ip",
					Description: `The fixed IP address corresponding to the floating IP.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `UUID of the compute instance associated with the floating IP. ## Import Floating IPs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_compute_floatingip_v2.floatip_1 89c60255-9bd6-460c-822a-e2b959ede9d2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "pool",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The actual floating IP address itself.`,
				},
				resource.Attribute{
					Name:        "fixed_ip",
					Description: `The fixed IP address corresponding to the floating IP.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `UUID of the compute instance associated with the floating IP. ## Import Floating IPs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_compute_floatingip_v2.floatip_1 89c60255-9bd6-460c-822a-e2b959ede9d2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_compute_instance_v2",
			Category:         "Elastic Cloud Server (ECS)",
			ShortDescription: `Manages a V2 VM instance resource within HuaweiCloud.`,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"cloud",
				"server",
				"ecs",
				"compute",
				"instance",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to create the server instance. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the resource.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Optional; Required if ` + "`" + `image_name` + "`" + ` is empty and not booting from a volume. Do not specify if booting from a volume.) The image ID of the desired image for the server. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `(Optional; Required if ` + "`" + `image_id` + "`" + ` is empty and not booting from a volume. Do not specify if booting from a volume.) The name of the desired image for the server. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "flavor_id",
					Description: `(Optional; Required if ` + "`" + `flavor_name` + "`" + ` is empty) The flavor ID of the desired flavor for the server. Changing this resizes the existing server.`,
				},
				resource.Attribute{
					Name:        "flavor_name",
					Description: `(Optional; Required if ` + "`" + `flavor_id` + "`" + ` is empty) The name of the desired flavor for the server. Changing this resizes the existing server.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional) The user data to provide when launching the instance. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Optional) An array of one or more security group names to associate with the server. Changing this results in adding/removing security groups from the existing server.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required) The availability zone in which to create the server. Please refer to https://developer.huaweicloud.com/endpoint for the values. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Optional) An array of one or more networks to attach to the instance. The network object structure is documented below. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) Metadata key/value pairs to make available from within the instance. Changing this updates the existing server metadata.`,
				},
				resource.Attribute{
					Name:        "config_drive",
					Description: `(Optional) Whether to use the config_drive feature to configure the instance. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "admin_pass",
					Description: `(Optional) The administrative password to assign to the server. Changing this changes the root password on the existing server.`,
				},
				resource.Attribute{
					Name:        "key_pair",
					Description: `(Optional) The name of a key pair to put on the server. The key pair must already be created and associated with the tenant's account. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "block_device",
					Description: `(Optional) Configuration of block devices. The block_device structure is documented below. Changing this creates a new server. You can specify multiple block devices which will create an instance with multiple disks. This configuration is very flexible, so please see the following [reference](http://docs.openstack.org/developer/nova/block_device_mapping.html) for more information.`,
				},
				resource.Attribute{
					Name:        "scheduler_hints",
					Description: `(Optional) Provide the Nova scheduler with hints on how the instance should be launched. The available hints are described below.`,
				},
				resource.Attribute{
					Name:        "stop_before_destroy",
					Description: `(Optional) Whether to try stop instance gracefully before destroying it, thus giving chance for guest OS daemons to stop correctly. If instance doesn't stop within timeout, it will be destroyed anyway. The ` + "`" + `network` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `(Required unless ` + "`" + `port` + "`" + ` or ` + "`" + `name` + "`" + ` is provided) The network UUID to attach to the server. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required unless ` + "`" + `uuid` + "`" + ` or ` + "`" + `port` + "`" + ` is provided) The human-readable name of the network. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required unless ` + "`" + `uuid` + "`" + ` or ` + "`" + `name` + "`" + ` is provided) The port UUID of a network to attach to the server. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "fixed_ip_v4",
					Description: `(Optional) Specifies a fixed IPv4 address to be used on this network. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "fixed_ip_v6",
					Description: `(Optional) Specifies a fixed IPv6 address to be used on this network. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "access_network",
					Description: `(Optional) Specifies if this network should be used for provisioning access. Accepts true or false. Defaults to false. The ` + "`" + `block_device` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `(Required unless ` + "`" + `source_type` + "`" + ` is set to ` + "`" + `"blank"` + "`" + ` ) The UUID of the image, volume, or snapshot. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) The source type of the device. Must be one of "blank", "image", "volume", or "snapshot". Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `The size of the volume to create (in gigabytes). Required in the following combinations: source=image and destination=volume, source=blank and destination=local, and source=blank and destination=volume. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "boot_index",
					Description: `(Optional) The boot index of the volume. It defaults to 0. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "destination_type",
					Description: `(Optional) The type that gets created. Possible values are "volume" and "local". Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "delete_on_termination",
					Description: `(Optional) Delete the volume / block device upon termination of the instance. Defaults to false. Changing this creates a new server. The ` + "`" + `scheduler_hints` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `(Optional) A UUID of a Server Group. The instance will be placed into that group.`,
				},
				resource.Attribute{
					Name:        "different_host",
					Description: `(Optional) A list of instance UUIDs. The instance will be scheduled on a different host than all other instances.`,
				},
				resource.Attribute{
					Name:        "same_host",
					Description: `(Optional) A list of instance UUIDs. The instance will be scheduled on the same host of those specified.`,
				},
				resource.Attribute{
					Name:        "query",
					Description: `(Optional) A conditional query that a compute node must pass in order to host an instance.`,
				},
				resource.Attribute{
					Name:        "target_cell",
					Description: `(Optional) The name of a cell to host the instance.`,
				},
				resource.Attribute{
					Name:        "build_near_host_ip",
					Description: `(Optional) An IP Address in CIDR form. The instance will be placed on a compute node that is in the same subnet.`,
				},
				resource.Attribute{
					Name:        "tenancy",
					Description: `(Optional) The tenancy specifies whether the ECS is to be created on a Dedicated Host (DeH) or in a shared pool.`,
				},
				resource.Attribute{
					Name:        "deh_id",
					Description: `(Optional) The ID of DeH. This parameter takes effect only when the value of tenancy is dedicated. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "access_ip_v4",
					Description: `The first detected Fixed IPv4 address _or_ the Floating IP.`,
				},
				resource.Attribute{
					Name:        "access_ip_v6",
					Description: `The first detected Fixed IPv6 address.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "flavor_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "flavor_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/uuid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/fixed_ip_v4",
					Description: `The Fixed IPv4 address of the Instance on that network.`,
				},
				resource.Attribute{
					Name:        "network/fixed_ip_v6",
					Description: `The Fixed IPv6 address of the Instance on that network.`,
				},
				resource.Attribute{
					Name:        "network/mac",
					Description: `The MAC address of the NIC on that network.`,
				},
				resource.Attribute{
					Name:        "volume_attached/volume_id",
					Description: `The volume id on that attachment.`,
				},
				resource.Attribute{
					Name:        "volume_attached/pci_address",
					Description: `The volume pci address on that attachment.`,
				},
				resource.Attribute{
					Name:        "volume_attached/boot_index",
					Description: `The volume boot index on that attachment.`,
				},
				resource.Attribute{
					Name:        "volume_attached/size",
					Description: `The volume size on that attachment.`,
				},
				resource.Attribute{
					Name:        "all_metadata",
					Description: `Contains all instance metadata, even metadata not set by Terraform. ## Notes ### Multiple Ephemeral Disks It's possible to specify multiple ` + "`" + `block_device` + "`" + ` entries to create an instance with multiple ephemeral (local) disks. In order to create multiple ephemeral disks, the sum of the total amount of ephemeral space must be less than or equal to what the chosen flavor supports. The following example shows how to create an instance with multiple ephemeral disks: ` + "`" + `` + "`" + `` + "`" + ` resource "huaweicloud_compute_instance_v2" "foo" { name = "terraform-test" security_groups = ["default"] block_device { boot_index = 0 delete_on_termination = true destination_type = "local" source_type = "image" uuid = "<image uuid>" } block_device { boot_index = -1 delete_on_termination = true destination_type = "local" source_type = "blank" volume_size = 1 } block_device { boot_index = -1 delete_on_termination = true destination_type = "local" source_type = "blank" volume_size = 1 } } ` + "`" + `` + "`" + `` + "`" + ` ### Instances and Ports Neutron Ports are a great feature and provide a lot of functionality. However, there are some notes to be aware of when mixing Instances and Ports:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "access_ip_v4",
					Description: `The first detected Fixed IPv4 address _or_ the Floating IP.`,
				},
				resource.Attribute{
					Name:        "access_ip_v6",
					Description: `The first detected Fixed IPv6 address.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "flavor_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "flavor_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/uuid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/fixed_ip_v4",
					Description: `The Fixed IPv4 address of the Instance on that network.`,
				},
				resource.Attribute{
					Name:        "network/fixed_ip_v6",
					Description: `The Fixed IPv6 address of the Instance on that network.`,
				},
				resource.Attribute{
					Name:        "network/mac",
					Description: `The MAC address of the NIC on that network.`,
				},
				resource.Attribute{
					Name:        "volume_attached/volume_id",
					Description: `The volume id on that attachment.`,
				},
				resource.Attribute{
					Name:        "volume_attached/pci_address",
					Description: `The volume pci address on that attachment.`,
				},
				resource.Attribute{
					Name:        "volume_attached/boot_index",
					Description: `The volume boot index on that attachment.`,
				},
				resource.Attribute{
					Name:        "volume_attached/size",
					Description: `The volume size on that attachment.`,
				},
				resource.Attribute{
					Name:        "all_metadata",
					Description: `Contains all instance metadata, even metadata not set by Terraform. ## Notes ### Multiple Ephemeral Disks It's possible to specify multiple ` + "`" + `block_device` + "`" + ` entries to create an instance with multiple ephemeral (local) disks. In order to create multiple ephemeral disks, the sum of the total amount of ephemeral space must be less than or equal to what the chosen flavor supports. The following example shows how to create an instance with multiple ephemeral disks: ` + "`" + `` + "`" + `` + "`" + ` resource "huaweicloud_compute_instance_v2" "foo" { name = "terraform-test" security_groups = ["default"] block_device { boot_index = 0 delete_on_termination = true destination_type = "local" source_type = "image" uuid = "<image uuid>" } block_device { boot_index = -1 delete_on_termination = true destination_type = "local" source_type = "blank" volume_size = 1 } block_device { boot_index = -1 delete_on_termination = true destination_type = "local" source_type = "blank" volume_size = 1 } } ` + "`" + `` + "`" + `` + "`" + ` ### Instances and Ports Neutron Ports are a great feature and provide a lot of functionality. However, there are some notes to be aware of when mixing Instances and Ports:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_compute_interface_attach_v2",
			Category:         "Elastic Cloud Server (ECS)",
			ShortDescription: `Attaches a Network Interface to an Instance.`,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"cloud",
				"server",
				"ecs",
				"compute",
				"interface",
				"attach",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to create the interface attachment. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new attachment.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) The ID of the Instance to attach the Port or Network to.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `(Optional) The ID of the Port to attach to an Instance. _NOTE_: This option and ` + "`" + `network_id` + "`" + ` are mutually exclusive.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Optional) The ID of the Network to attach to an Instance. A port will be created automatically. _NOTE_: This option and ` + "`" + `port_id` + "`" + ` are mutually exclusive.`,
				},
				resource.Attribute{
					Name:        "fixed_ip",
					Description: `(Optional) An IP address to assosciate with the port. _NOTE_: This option cannot be used with port_id. You must specifiy a network_id. The IP address must lie in a range on the supplied network. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_compute_keypair_v2",
			Category:         "Elastic Cloud Server (ECS)",
			ShortDescription: `Manages a V2 keypair resource within HuaweiCloud.`,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"cloud",
				"server",
				"ecs",
				"compute",
				"keypair",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Compute client. Keypairs are associated with accounts, but a Compute client is needed to create one. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new keypair.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the keypair. Changing this creates a new keypair.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(Required) A pregenerated OpenSSH-formatted public key. Changing this creates a new keypair.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional options. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `See Argument Reference above. ## Import Keypairs can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_compute_keypair_v2.my-keypair test-keypair ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `See Argument Reference above. ## Import Keypairs can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_compute_keypair_v2.my-keypair test-keypair ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_compute_secgroup_v2",
			Category:         "Elastic Cloud Server (ECS)",
			ShortDescription: `Manages a V2 security group resource within HuaweiCloud.`,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"cloud",
				"server",
				"ecs",
				"compute",
				"secgroup",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Compute client. A Compute client is needed to create a security group. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new security group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the security group. Changing this updates the ` + "`" + `name` + "`" + ` of an existing security group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) A description for the security group. Changing this updates the ` + "`" + `description` + "`" + ` of an existing security group.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Optional) A rule describing how the security group operates. The rule object structure is documented below. Changing this updates the security group rules. As shown in the example above, multiple rule blocks may be used. The ` + "`" + `rule` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "from_port",
					Description: `(Required) An integer representing the lower bound of the port range to open. Changing this creates a new security group rule.`,
				},
				resource.Attribute{
					Name:        "to_port",
					Description: `(Required) An integer representing the upper bound of the port range to open. Changing this creates a new security group rule.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `(Required) The protocol type that will be allowed. Changing this creates a new security group rule.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Optional) Required if ` + "`" + `from_group_id` + "`" + ` or ` + "`" + `self` + "`" + ` is empty. The IP range that will be the source of network traffic to the security group. Use 0.0.0.0/0 to allow all IP addresses. Changing this creates a new security group rule. Cannot be combined with ` + "`" + `from_group_id` + "`" + ` or ` + "`" + `self` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "from_group_id",
					Description: `(Optional) Required if ` + "`" + `cidr` + "`" + ` or ` + "`" + `self` + "`" + ` is empty. The ID of a group from which to forward traffic to the parent group. Changing this creates a new security group rule. Cannot be combined with ` + "`" + `cidr` + "`" + ` or ` + "`" + `self` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "self",
					Description: `(Optional) Required if ` + "`" + `cidr` + "`" + ` and ` + "`" + `from_group_id` + "`" + ` is empty. If true, the security group itself will be added as a source to this ingress rule. Cannot be combined with ` + "`" + `cidr` + "`" + ` or ` + "`" + `from_group_id` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `See Argument Reference above. ## Notes ### ICMP Rules When using ICMP as the ` + "`" + `ip_protocol` + "`" + `, the ` + "`" + `from_port` + "`" + ` sets the ICMP _type_ and the ` + "`" + `to_port` + "`" + ` sets the ICMP _code_. To allow all ICMP types, set each value to ` + "`" + `-1` + "`" + `, like so: ` + "`" + `` + "`" + `` + "`" + `hcl rule { from_port = -1 to_port = -1 ip_protocol = "icmp" cidr = "0.0.0.0/0" } ` + "`" + `` + "`" + `` + "`" + ` A list of ICMP types and codes can be found [here](https://en.wikipedia.org/wiki/Internet_Control_Message_Protocol#Control_messages). ### Referencing Security Groups When referencing a security group in a configuration (for example, a configuration creates a new security group and then needs to apply it to an instance being created in the same configuration), it is currently recommended to reference the security group by name and not by ID, like this: ` + "`" + `` + "`" + `` + "`" + `hcl resource "huaweicloud_compute_instance_v2" "test-server" { name = "tf-test" image_id = "ad091b52-742f-469e-8f3c-fd81cadf0743" flavor_id = "3" key_pair = "my_key_pair_name" security_groups = ["${huaweicloud_compute_secgroup_v2.secgroup_1.name}"] } ` + "`" + `` + "`" + `` + "`" + ` ## Import Security Groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_compute_secgroup_v2.my_secgroup 1bc30ee9-9d5b-4c30-bdd5-7f1e663f5edf ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `See Argument Reference above. ## Notes ### ICMP Rules When using ICMP as the ` + "`" + `ip_protocol` + "`" + `, the ` + "`" + `from_port` + "`" + ` sets the ICMP _type_ and the ` + "`" + `to_port` + "`" + ` sets the ICMP _code_. To allow all ICMP types, set each value to ` + "`" + `-1` + "`" + `, like so: ` + "`" + `` + "`" + `` + "`" + `hcl rule { from_port = -1 to_port = -1 ip_protocol = "icmp" cidr = "0.0.0.0/0" } ` + "`" + `` + "`" + `` + "`" + ` A list of ICMP types and codes can be found [here](https://en.wikipedia.org/wiki/Internet_Control_Message_Protocol#Control_messages). ### Referencing Security Groups When referencing a security group in a configuration (for example, a configuration creates a new security group and then needs to apply it to an instance being created in the same configuration), it is currently recommended to reference the security group by name and not by ID, like this: ` + "`" + `` + "`" + `` + "`" + `hcl resource "huaweicloud_compute_instance_v2" "test-server" { name = "tf-test" image_id = "ad091b52-742f-469e-8f3c-fd81cadf0743" flavor_id = "3" key_pair = "my_key_pair_name" security_groups = ["${huaweicloud_compute_secgroup_v2.secgroup_1.name}"] } ` + "`" + `` + "`" + `` + "`" + ` ## Import Security Groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_compute_secgroup_v2.my_secgroup 1bc30ee9-9d5b-4c30-bdd5-7f1e663f5edf ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_compute_servergroup_v2",
			Category:         "Elastic Cloud Server (ECS)",
			ShortDescription: `Manages a V2 Server Group resource within HuaweiCloud.`,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"cloud",
				"server",
				"ecs",
				"compute",
				"servergroup",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Compute client. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new server group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the server group. Changing this creates a new server group.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Required) The set of policies for the server group. Only two two policies are available right now, and both are mutually exclusive. See the Policies section for more information. Changing this creates a new server group.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional options. ## Policies`,
				},
				resource.Attribute{
					Name:        "affinity",
					Description: `All instances/servers launched in this group will be hosted on the same compute node.`,
				},
				resource.Attribute{
					Name:        "anti-affinity",
					Description: `All instances/servers launched in this group will be hosted on different compute nodes. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `The instances that are part of this server group. ## Import Server Groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_compute_servergroup_v2.test-sg 1bc30ee9-9d5b-4c30-bdd5-7f1e663f5edf ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `The instances that are part of this server group. ## Import Server Groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_compute_servergroup_v2.test-sg 1bc30ee9-9d5b-4c30-bdd5-7f1e663f5edf ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_compute_volume_attach_v2",
			Category:         "Elastic Cloud Server (ECS)",
			ShortDescription: `Attaches a Block Storage Volume to an Instance.`,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"cloud",
				"server",
				"ecs",
				"compute",
				"volume",
				"attach",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Compute client. A Compute client is needed to create a volume attachment. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new volume attachment.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) The ID of the Instance to attach the Volume to.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `(Required) The ID of the Volume to attach to an Instance.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `(Optional) The device of the volume attachment (ex: ` + "`" + `/dev/vdc` + "`" + `). _NOTE_: Being able to specify a device is dependent upon the hypervisor in use. There is a chance that the device specified in Terraform will not be the same device the hypervisor chose. If this happens, Terraform will wish to update the device upon subsequent applying which will cause the volume to be detached and reattached indefinitely. Please use with caution. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `See Argument Reference above. _NOTE_: The correctness of this information is dependent upon the hypervisor in use. In some cases, this should not be used as an authoritative piece of information.`,
				},
				resource.Attribute{
					Name:        "pci_address",
					Description: `PCI address of the block device. ## Import Volume Attachments can be imported using the Instance ID and Volume ID separated by a slash, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_compute_volume_attach_v2.va_1 89c60255-9bd6-460c-822a-e2b959ede9d2/45670584-225f-46c3-b33e-6707b589b666 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `See Argument Reference above. _NOTE_: The correctness of this information is dependent upon the hypervisor in use. In some cases, this should not be used as an authoritative piece of information.`,
				},
				resource.Attribute{
					Name:        "pci_address",
					Description: `PCI address of the block device. ## Import Volume Attachments can be imported using the Instance ID and Volume ID separated by a slash, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_compute_volume_attach_v2.va_1 89c60255-9bd6-460c-822a-e2b959ede9d2/45670584-225f-46c3-b33e-6707b589b666 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_cs_cluster_v1",
			Category:         "Cloud Stream",
			ShortDescription: `Cloud Stream Service cluster management`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"stream",
				"cs",
				"cluster",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Cluster name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) cluster description.`,
				},
				resource.Attribute{
					Name:        "max_spu_num",
					Description: `(Optional) Cluster maximum SPU number.`,
				},
				resource.Attribute{
					Name:        "subnet_cidr",
					Description: `(Optional) Cluster sub segment. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "subnet_gateway",
					Description: `(Optional) Cluster subnet gateway. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "vpc_cidr",
					Description: `(Optional) Cluster VPC network segment. Changing this parameter will create a new resource. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Cluster creation time.`,
				},
				resource.Attribute{
					Name:        "manager_node_spu_num",
					Description: `Cluster management node SPU number.`,
				},
				resource.Attribute{
					Name:        "used_spu_num",
					Description: `The used SPU number of Cluster. ## Timeouts This resource provides the following timeouts configuration options: - ` + "`" + `create` + "`" + ` - Default is 30 minute. - ` + "`" + `delete` + "`" + ` - Default is 30 minute.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `Cluster creation time.`,
				},
				resource.Attribute{
					Name:        "manager_node_spu_num",
					Description: `Cluster management node SPU number.`,
				},
				resource.Attribute{
					Name:        "used_spu_num",
					Description: `The used SPU number of Cluster. ## Timeouts This resource provides the following timeouts configuration options: - ` + "`" + `create` + "`" + ` - Default is 30 minute. - ` + "`" + `delete` + "`" + ` - Default is 30 minute.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_cs_peering_connect_v1",
			Category:         "Cloud Stream",
			ShortDescription: `Cloud Stream Service cluster peering connect management`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"stream",
				"cs",
				"peering",
				"connect",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) The id of cloud stream cluster. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of peering connection. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "target_vpc_info",
					Description: `(Optional) The information of target vpc. Structure is documented below. Changing this parameter will create a new resource. The ` + "`" + `target_vpc_info` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) The project ID to which target vpc belongs. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) The VPC ID. Changing this parameter will create a new resource. ## Timeouts This resource provides the following timeouts configuration options: - ` + "`" + `create` + "`" + ` - Default is 30 minute.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_cs_route_v1",
			Category:         "Cloud Stream",
			ShortDescription: `Cloud Stream Service cluster peering connect route management`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"stream",
				"cs",
				"route",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) The id of cloud stream cluster. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `(Required) Routing destination CIDR. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "peering_id",
					Description: `(Required) The peering connection id of cloud stream cluster. Changing this parameter will create a new resource.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_csbs_backup_policy_v1",
			Category:         "Cloud Server Backup Service (CSBS)",
			ShortDescription: `Provides an HuaweiCloud Backup Policy of Resource.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"server",
				"backup",
				"service",
				"csbs",
				"policy",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of backup policy. The value consists of 1 to 255 characters and can contain only letters, digits, underscores (_), and hyphens (-).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Backup policy description. The value consists of 0 to 255 characters and must not contain a greater-than sign (>) or less-than sign (<).`,
				},
				resource.Attribute{
					Name:        "provider_id",
					Description: `(Required) Specifies backup provider ID. Default value is`,
				},
				resource.Attribute{
					Name:        "common",
					Description: `(Optional) General backup policy parameters, which are blank by default.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Specifies Scheduling period name.The value consists of 1 to 255 characters and can contain only letters, digits, underscores (_), and hyphens (-).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Specifies Scheduling period description.The value consists of 0 to 255 characters and must not contain a greater-than sign (>) or less-than sign (<).`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Specifies whether the scheduling period is enabled. Default value is`,
				},
				resource.Attribute{
					Name:        "max_backups",
					Description: `(Optional) Specifies maximum number of backups that can be automatically created for a backup object.`,
				},
				resource.Attribute{
					Name:        "retention_duration_days",
					Description: `(Optional) Specifies duration of retaining a backup, in days.`,
				},
				resource.Attribute{
					Name:        "permanent",
					Description: `(Optional) Specifies whether backups are permanently retained.`,
				},
				resource.Attribute{
					Name:        "trigger_pattern",
					Description: `(Required) Specifies Scheduling policy of the scheduler.`,
				},
				resource.Attribute{
					Name:        "operation_type",
					Description: `(Required) Specifies Operation type, which can be backup.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) Specifies the ID of the object to be backed up.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Entity object type of the backup object. If the type is VMs, the value is`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies backup object name. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of Backup Policy.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Backup Policy ID.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Specifies Scheduling period ID.`,
				},
				resource.Attribute{
					Name:        "trigger_id",
					Description: `Specifies Scheduler ID.`,
				},
				resource.Attribute{
					Name:        "trigger_name",
					Description: `Specifies Scheduler name.`,
				},
				resource.Attribute{
					Name:        "trigger_type",
					Description: `Specifies Scheduler type. ## Import Backup Policy can be imported using ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_csbs_backup_policy_v1.backup_policy_v1 7056d636-ac60-4663-8a6c-82d3c32c1c64 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `Status of Backup Policy.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Backup Policy ID.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Specifies Scheduling period ID.`,
				},
				resource.Attribute{
					Name:        "trigger_id",
					Description: `Specifies Scheduler ID.`,
				},
				resource.Attribute{
					Name:        "trigger_name",
					Description: `Specifies Scheduler name.`,
				},
				resource.Attribute{
					Name:        "trigger_type",
					Description: `Specifies Scheduler type. ## Import Backup Policy can be imported using ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_csbs_backup_policy_v1.backup_policy_v1 7056d636-ac60-4663-8a6c-82d3c32c1c64 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_csbs_backup_v1",
			Category:         "Cloud Server Backup Service (CSBS)",
			ShortDescription: `Provides an HuaweiCloud Backup of Resources.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"server",
				"backup",
				"service",
				"csbs",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "backup_name",
					Description: `(Optional) Name for the backup. The value consists of 1 to 255 characters and can contain only letters, digits, underscores (_), and hyphens (-). Changing backup_name creates a new backup.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Backup description. The value consists of 0 to 255 characters and must not contain a greater-than sign (>) or less-than sign (<). Changing description creates a new backup.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required) ID of the target to which the backup is restored. Changing this creates a new backup.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `(Optional) Type of the target to which the backup is restored. The default value is`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `It specifies the status of backup.`,
				},
				resource.Attribute{
					Name:        "backup_record_id",
					Description: `Specifies backup record ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of backup Volume.`,
				},
				resource.Attribute{
					Name:        "space_saving_ratio",
					Description: `Specifies space saving rate.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `It gives EVS disk backup name.`,
				},
				resource.Attribute{
					Name:        "bootable",
					Description: `Specifies whether the disk is bootable.`,
				},
				resource.Attribute{
					Name:        "average_speed",
					Description: `Specifies the average speed.`,
				},
				resource.Attribute{
					Name:        "source_volume_size",
					Description: `Shows source volume size in GB.`,
				},
				resource.Attribute{
					Name:        "source_volume_id",
					Description: `It specifies source volume ID.`,
				},
				resource.Attribute{
					Name:        "incremental",
					Description: `Shows whether incremental backup is used.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `ID of snapshot.`,
				},
				resource.Attribute{
					Name:        "source_volume_name",
					Description: `Specifies source volume name.`,
				},
				resource.Attribute{
					Name:        "image_type",
					Description: `It specifies backup. The default value is backup.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Specifies Cinder backup ID.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Specifies accumulated size (MB) of backups.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of backup data.`,
				},
				resource.Attribute{
					Name:        "eip",
					Description: `Specifies elastic IP address of the ECS.`,
				},
				resource.Attribute{
					Name:        "cloud_service_type",
					Description: `Specifies ECS type.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `Specifies memory size of the ECS, in MB.`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `Specifies CPU cores corresponding to the ECS.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `It specifies internal IP address of the ECS.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `Shows system disk size corresponding to the ECS specifications.`,
				},
				resource.Attribute{
					Name:        "image_type",
					Description: `Specifies image type. ## Import Backup can be imported using ` + "`" + `backup_record_id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_csbs_backup_v1.backup_v1.backup_v1 7056d636-ac60-4663-8a6c-82d3c32c1c64 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `It specifies the status of backup.`,
				},
				resource.Attribute{
					Name:        "backup_record_id",
					Description: `Specifies backup record ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of backup Volume.`,
				},
				resource.Attribute{
					Name:        "space_saving_ratio",
					Description: `Specifies space saving rate.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `It gives EVS disk backup name.`,
				},
				resource.Attribute{
					Name:        "bootable",
					Description: `Specifies whether the disk is bootable.`,
				},
				resource.Attribute{
					Name:        "average_speed",
					Description: `Specifies the average speed.`,
				},
				resource.Attribute{
					Name:        "source_volume_size",
					Description: `Shows source volume size in GB.`,
				},
				resource.Attribute{
					Name:        "source_volume_id",
					Description: `It specifies source volume ID.`,
				},
				resource.Attribute{
					Name:        "incremental",
					Description: `Shows whether incremental backup is used.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `ID of snapshot.`,
				},
				resource.Attribute{
					Name:        "source_volume_name",
					Description: `Specifies source volume name.`,
				},
				resource.Attribute{
					Name:        "image_type",
					Description: `It specifies backup. The default value is backup.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Specifies Cinder backup ID.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Specifies accumulated size (MB) of backups.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of backup data.`,
				},
				resource.Attribute{
					Name:        "eip",
					Description: `Specifies elastic IP address of the ECS.`,
				},
				resource.Attribute{
					Name:        "cloud_service_type",
					Description: `Specifies ECS type.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `Specifies memory size of the ECS, in MB.`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `Specifies CPU cores corresponding to the ECS.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `It specifies internal IP address of the ECS.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `Shows system disk size corresponding to the ECS specifications.`,
				},
				resource.Attribute{
					Name:        "image_type",
					Description: `Specifies image type. ## Import Backup can be imported using ` + "`" + `backup_record_id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_csbs_backup_v1.backup_v1.backup_v1 7056d636-ac60-4663-8a6c-82d3c32c1c64 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_css_cluster_v1",
			Category:         "Cloud Search Service (CSS)",
			ShortDescription: `cluster management`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"search",
				"service",
				"css",
				"cluster",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "engine_version",
					Description: `(Required) Engine version. Versions 5.5.1 and 6.2.3 are supported. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Cluster name. It contains 4 to 32 characters. Only letters, digits, hyphens (-), and underscores (_) are allowed. The value must start with a letter. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "node_config",
					Description: `(Required) Node configuration. Structure is documented below. Changing this parameter will create a new resource. The ` + "`" + `node_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) Availability zone (AZ). Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `(Required) Instance flavor name. Value range of flavor ess.spec-1u8g: 40 GB to 640 GB Value range of flavor ess.spec-2u16g: 40 GB to 1280 GB Value range of flavor ess.spec-4u32g: 40 GB to 2560 GB Value range of flavor ess.spec-8u64g: 80 GB to 5120 GB Value range of flavor ess.spec-16u128g: 160 GB to 10240 GB. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "network_info",
					Description: `(Required) Network information. Structure is documented below. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "volume",
					Description: `(Required) Information about the volume. Structure is documented below. Changing this parameter will create a new resource. The ` + "`" + `network_info` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Required) Security group ID. All instances in a cluster must have the same subnets and security groups. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) Subnet ID. All instances in a cluster must have the same subnets and security groups. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) VPC ID, which is used for configuring cluster network. Changing this parameter will create a new resource. The ` + "`" + `volume` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) Volume size, which must be a multiple of 4 and 10. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `(Required) COMMON: Common I/O. The SATA disk is used. HIGH: High I/O. The SAS disk is used. ULTRAHIGH: Ultra-high I/O. The solid-state drive (SSD) is used. Changing this parameter will create a new resource. - - -`,
				},
				resource.Attribute{
					Name:        "engine_type",
					Description: `(Optional) Engine type. The default value is elasticsearch. Currently, the value can only be elasticsearch. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "expect_node_num",
					Description: `(Optional) Number of cluster instances. The value range is 1 to 32. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `Time when a cluster is created. The format is ISO8601: CCYY-MM-DDThh:mm:ss.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `Indicates the IP address and port number.`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `List of node objects. Structure is documented below. The ` + "`" + `nodes` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Instance ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Instance name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Supported type: ess (indicating the Elasticsearch node). ## Timeouts This resource provides the following timeouts configuration options: - ` + "`" + `create` + "`" + ` - Default is 30 minute. - ` + "`" + `update` + "`" + ` - Default is 30 minute.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created",
					Description: `Time when a cluster is created. The format is ISO8601: CCYY-MM-DDThh:mm:ss.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `Indicates the IP address and port number.`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `List of node objects. Structure is documented below. The ` + "`" + `nodes` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Instance ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Instance name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Supported type: ess (indicating the Elasticsearch node). ## Timeouts This resource provides the following timeouts configuration options: - ` + "`" + `create` + "`" + ` - Default is 30 minute. - ` + "`" + `update` + "`" + ` - Default is 30 minute.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_cts_tracker_v1",
			Category:         "Cloud Trace Service (CTS)",
			ShortDescription: `CTS tracker allows you to collect, store, and query cloud resource operation records and use these records for security analysis, compliance auditing, resource tracking, and fault locating.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"trace",
				"service",
				"cts",
				"tracker",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket_name",
					Description: `(Required) The OBS bucket name for a tracker.`,
				},
				resource.Attribute{
					Name:        "file_prefix_name",
					Description: `(Optional) The prefix of a log that needs to be stored in an OBS bucket.`,
				},
				resource.Attribute{
					Name:        "is_support_smn",
					Description: `(Required) Specifies whether SMN is supported. When the value is false, topic_id and operations can be left empty.`,
				},
				resource.Attribute{
					Name:        "topic_id",
					Description: `(Required)The theme of the SMN service, Is obtained from SMN and in the format of`,
				},
				resource.Attribute{
					Name:        "operations",
					Description: `(Required) Trigger conditions for sending a notification.`,
				},
				resource.Attribute{
					Name:        "is_send_all_key_operation",
					Description: `(Required) When the value is`,
				},
				resource.Attribute{
					Name:        "need_notify_user_list",
					Description: `(Optional) The users using the login function. When these users log in, notifications will be sent. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of a tracker. The value is`,
				},
				resource.Attribute{
					Name:        "tracker_name",
					Description: `The tracker name. Currently, only tracker`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `The status of a tracker. The value is`,
				},
				resource.Attribute{
					Name:        "tracker_name",
					Description: `The tracker name. Currently, only tracker`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_dcs_instance_v1",
			Category:         "Distributed Cache Service",
			ShortDescription: `Manages a DCS instance in the huaweicloud DCS Service`,
			Description:      ``,
			Keywords: []string{
				"distributed",
				"cache",
				"service",
				"dcs",
				"instance",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Indicates the name of an instance. An instance name starts with a letter, consists of 4 to 64 characters, and supports only letters, digits, and hyphens (-).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Indicates the description of an instance. It is a character string containing not more than 1024 characters.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `(Optional) Indicates a cache engine. Options: Redis and Memcached. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `(Optional) Indicates the version of a message engine. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `(Required) Indicates the Cache capacity. Unit: GB. For a DCS Redis or Memcached instance in single-node or master/standby mode, the cache capacity can be 2 GB, 4 GB, 8 GB, 16 GB, 32 GB, or 64 GB. For a DCS Redis instance in cluster mode, the cache capacity can be 64, 128, 256, 512, or 1024 GB. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "partition_num",
					Description: `(Optional) This parameter is mandatory when a Kafka instance is created. Indicates the maximum number of topics in a Kafka instance. When specification is 300 MB: 900 When specification is 600 MB: 1800 When specification is 1200 MB: 1800`,
				},
				resource.Attribute{
					Name:        "access_user",
					Description: `(Optional) Username used for accessing a DCS instance after password authentication. A username starts with a letter, consists of 1 to 64 characters, and supports only letters, digits, and hyphens (-). Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Password of a DCS instance. The password of a DCS Redis instance must meet the following complexity requirements: Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) Tenant's VPC ID. For details on how to create VPCs, see the Virtual Private Cloud API Reference. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Required) Tenant's security group ID. For details on how to create security groups, see the Virtual Private Cloud API Reference.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) Subnet ID. For details on how to create subnets, see the Virtual Private Cloud API Reference. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "available_zones",
					Description: `(Required) IDs of the AZs where cache nodes reside. For details on how to query AZs, see Querying AZ Information. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "product_id",
					Description: `(Required) Product ID used to differentiate DCS instance types. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "maintain_begin",
					Description: `(Optional) Indicates the time at which a maintenance time window starts. Format: HH:mm:ss. The start time and end time of a maintenance time window must indicate the time segment of a supported maintenance time window. For details, see section Querying Maintenance Time Windows. The start time must be set to 22:00, 02:00, 06:00, 10:00, 14:00, or 18:00. Parameters maintain_begin and maintain_end must be set in pairs. If parameter maintain_begin is left blank, parameter maintain_end is also blank. In this case, the system automatically allocates the default start time 02:00.`,
				},
				resource.Attribute{
					Name:        "maintain_end",
					Description: `(Optional) Indicates the time at which a maintenance time window ends. Format: HH:mm:ss. The start time and end time of a maintenance time window must indicate the time segment of a supported maintenance time window. For details, see section Querying Maintenance Time Windows. The end time is four hours later than the start time. For example, if the start time is 22:00, the end time is 02:00. Parameters maintain_begin and maintain_end must be set in pairs. If parameter maintain_end is left blank, parameter maintain_begin is also blank. In this case, the system automatically allocates the default end time 06:00.`,
				},
				resource.Attribute{
					Name:        "save_days",
					Description: `(Optional) Retention time. Unit: day. Range: 1–7. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "backup_type",
					Description: `(Optional) Backup type. Options: auto: automatic backup. manual: manual backup. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "begin_at",
					Description: `(Optional) Time at which backup starts. "00:00-01:00" indicates that backup starts at 00:00:00. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "period_type",
					Description: `(Optional) Interval at which backup is performed. Currently, only weekly backup is supported. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "backup_at",
					Description: `(Optional) Day in a week on which backup starts. Range: 1–7. Where: 1 indicates Monday; 7 indicates Sunday. Changing this creates a new instance. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "access_user",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vpc_name",
					Description: `Indicates the name of a vpc.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `Indicates the name of a security group.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subnet_name",
					Description: `Indicates the name of a subnet.`,
				},
				resource.Attribute{
					Name:        "available_zones",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "product_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "maintain_begin",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "maintain_end",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "save_days",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "backup_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "begin_at",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "period_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "backup_at",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "order_id",
					Description: `An order ID is generated only in the monthly or yearly billing mode. In other billing modes, no value is returned for this parameter.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port of the cache node.`,
				},
				resource.Attribute{
					Name:        "resource_spec_code",
					Description: `Resource specifications. dcs.single_node: indicates a DCS instance in single-node mode. dcs.master_standby: indicates a DCS instance in master/standby mode. dcs.cluster: indicates a DCS instance in cluster mode.`,
				},
				resource.Attribute{
					Name:        "used_memory",
					Description: `Size of the used memory. Unit: MB.`,
				},
				resource.Attribute{
					Name:        "internal_version",
					Description: `Internal DCS version.`,
				},
				resource.Attribute{
					Name:        "max_memory",
					Description: `Overall memory size. Unit: MB.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `Indicates a user ID.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Cache node's IP address in tenant's VPC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "access_user",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vpc_name",
					Description: `Indicates the name of a vpc.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `Indicates the name of a security group.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subnet_name",
					Description: `Indicates the name of a subnet.`,
				},
				resource.Attribute{
					Name:        "available_zones",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "product_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "maintain_begin",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "maintain_end",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "save_days",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "backup_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "begin_at",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "period_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "backup_at",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "order_id",
					Description: `An order ID is generated only in the monthly or yearly billing mode. In other billing modes, no value is returned for this parameter.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port of the cache node.`,
				},
				resource.Attribute{
					Name:        "resource_spec_code",
					Description: `Resource specifications. dcs.single_node: indicates a DCS instance in single-node mode. dcs.master_standby: indicates a DCS instance in master/standby mode. dcs.cluster: indicates a DCS instance in cluster mode.`,
				},
				resource.Attribute{
					Name:        "used_memory",
					Description: `Size of the used memory. Unit: MB.`,
				},
				resource.Attribute{
					Name:        "internal_version",
					Description: `Internal DCS version.`,
				},
				resource.Attribute{
					Name:        "max_memory",
					Description: `Overall memory size. Unit: MB.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `Indicates a user ID.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Cache node's IP address in tenant's VPC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_dds_instance_v3",
			Category:         "Document Database Service (DDS)",
			ShortDescription: `Manages dds instance resource within HuaweiCloud`,
			Description:      ``,
			Keywords: []string{
				"document",
				"database",
				"service",
				"dds",
				"instance",
				"v3",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Specifies the region of the DDS instance. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the DB instance name. The DB instance name of the same type is unique in the same tenant. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "datastore",
					Description: `(Required) Specifies database information. The structure is described below. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required) Specifies the ID of the availability zone. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) Specifies the VPC ID. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) Specifies the subnet Network ID. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Required) Specifies the security group ID of the DDS instance. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Specifies the Administrator password of the database instance. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "disk_encryption_id",
					Description: `(Required) Specifies the disk encryption ID of the instance. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Required) Specifies the mode of the database instance. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `(Required) Specifies the flavors information. The structure is described below. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "backup_strategy",
					Description: `(Optional) Specifies the advanced backup policy. The structure is described below. Changing this creates a new instance. The ` + "`" + `datastore` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Specifies the DB engine. 'DDS-Community' and 'DDS-Enhanced' are supported.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) Specifies the DB instance version. For the Community Edition, the valid values are 3.2, 3.4, or 4.0. For the Enhanced Edition, only 3.4 is supported now.`,
				},
				resource.Attribute{
					Name:        "storage_engine",
					Description: `(Optional) Specifies the storage engine of the DB instance. DDS Community Edition supports wiredTiger engine, and the Enhanced Edition supports rocksDB engine. The ` + "`" + `flavor` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Specifies the node type. Valid value:`,
				},
				resource.Attribute{
					Name:        "num",
					Description: `(Required) Specifies the node quantity. Valid value:`,
				},
				resource.Attribute{
					Name:        "storage",
					Description: `(Optional) Specifies the disk type. Valid value: ULTRAHIGH which indicates the type SSD.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) Specifies the disk size. The value must be a multiple of 10. The unit is GB. This parameter is mandatory for nodes except mongos and invalid for mongos.`,
				},
				resource.Attribute{
					Name:        "spec_code",
					Description: `(Required) Specifies the resource specification code. In a cluster instance, multiple specifications need to be specified. All specifications must be of the same series, that is, general-purpose (s6), enhanced (c3), or enhanced II (c6). For example:`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `(Required) Specifies the backup time window. Automated backups will be triggered during the backup time window. The value cannot be empty. It must be a valid value in the "hh:mm-HH:MM" format. The current time is in the UTC format.`,
				},
				resource.Attribute{
					Name:        "keep_days",
					Description: `(Optional) Specifies the number of days to retain the generated backup files. The value range is from 0 to 732.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "datastore",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "disk_encryption_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "backup_strategy",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "db_username",
					Description: `Indicates the DB Administator name.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates the the DB instance status`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "datastore",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "disk_encryption_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "backup_strategy",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "db_username",
					Description: `Indicates the DB Administator name.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates the the DB instance status`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_dis_stream_v2",
			Category:         "Data Ingestion Service (DIS)",
			ShortDescription: `DIS Stream management`,
			Description:      ``,
			Keywords: []string{
				"data",
				"ingestion",
				"service",
				"dis",
				"stream",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "partition_count",
					Description: `(Required) Number of the expect partitions. NOTE: Each stream can be scaled up and down a total of five times within one hour. After the stream is successfully scaled up or down, it cannot be scaled up or down again within the next one hour.`,
				},
				resource.Attribute{
					Name:        "stream_name",
					Description: `(Required) Name of the DIS stream to be created. - - -`,
				},
				resource.Attribute{
					Name:        "auto_scale_max_partition_count",
					Description: `(Optional) Maximum number of partition for automatic scaling. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "auto_scale_min_partition_count",
					Description: `(Optional) Minimum number of partition for automatic scaling. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "compression_format",
					Description: `(Optional) Data compression type. The value is one of snappy, gzip and zip. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "csv_delimiter",
					Description: `(Optional) Field separator for CSV file. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "data_schema",
					Description: `(Optional) User's JOSN, CSV format data schema, described with Avro schema. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "data_type",
					Description: `(Optional) Data type of the data putting into the stream. The value is one of BLOB, JSON and CSV. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "retention_period",
					Description: `(Optional) The number of hours for which data from the stream will be retained in DIS. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "stream_type",
					Description: `(Optional) Stream Type. The value is COMMON(means 1M bandwidth) or ADVANCED(means 5M bandwidth). Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) List of tags for the newly created DIS stream. Structure is documented below. Changing this parameter will create a new resource. The ` + "`" + `tags` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) The key of tag. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) The value of tag. Changing this parameter will create a new resource. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `Timestamp at which the DIS stream was created.`,
				},
				resource.Attribute{
					Name:        "readable_partition_count",
					Description: `Total number of readable partitions (including partitions in ACTIVE state only).`,
				},
				resource.Attribute{
					Name:        "writable_partition_count",
					Description: `Total number of writable partitions (including partitions in ACTIVE and DELETED states).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created",
					Description: `Timestamp at which the DIS stream was created.`,
				},
				resource.Attribute{
					Name:        "readable_partition_count",
					Description: `Total number of readable partitions (including partitions in ACTIVE state only).`,
				},
				resource.Attribute{
					Name:        "writable_partition_count",
					Description: `Total number of writable partitions (including partitions in ACTIVE and DELETED states).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_dli_queue_v1",
			Category:         "Data Lake Insight (DLI)",
			ShortDescription: `queue management`,
			Description:      ``,
			Keywords: []string{
				"data",
				"lake",
				"insight",
				"dli",
				"queue",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cu_count",
					Description: `(Required) Minimum number of CUs that are bound to a queue. The value can be 4, 16, or 64. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of a queue. The name can contain only digits, letters, and underscores (_), but cannot contain only digits or start with an underscore (_). Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of a queue. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "management_subnet_cidr",
					Description: `(Optional) CIDR of the management subnet. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "subnet_cidr",
					Description: `(Optional) Subnet CIDR. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "vpc_cidr",
					Description: `(Optional) VPC CIDR. Changing this parameter will create a new resource. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Time when a queue is created.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "create_time",
					Description: `Time when a queue is created.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_dms_group_v1",
			Category:         "Distributed Message Service (DMS)",
			ShortDescription: `Manages a DMS group in the huaweicloud DMS Service`,
			Description:      ``,
			Keywords: []string{
				"distributed",
				"message",
				"service",
				"dms",
				"group",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Indicates the unique name of a group. A string of 1 to 64 characters that contain a-z, A-Z, 0-9, hyphens (-), and underscores (_). The name cannot be modified once specified.`,
				},
				resource.Attribute{
					Name:        "queue_id",
					Description: `(Required) Indicates the ID of a specified queue. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "queue_id",
					Description: `Indicates the ID of a queue.`,
				},
				resource.Attribute{
					Name:        "redrive_policy",
					Description: `Indicates whether to enable dead letter messages.`,
				},
				resource.Attribute{
					Name:        "produced_messages",
					Description: `Indicates the total number of messages (not including the messages that have expired and been deleted) in a queue.`,
				},
				resource.Attribute{
					Name:        "consumed_messages",
					Description: `Indicates the total number of messages that are successfully consumed.`,
				},
				resource.Attribute{
					Name:        "available_messages",
					Description: `Indicates the accumulated number of messages that can be consumed.`,
				},
				resource.Attribute{
					Name:        "produced_deadletters",
					Description: `Indicates the total number of dead letter messages generated by the consumer group.`,
				},
				resource.Attribute{
					Name:        "available_deadletters",
					Description: `Indicates the accumulated number of dead letter messages that have not been consumed.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "queue_id",
					Description: `Indicates the ID of a queue.`,
				},
				resource.Attribute{
					Name:        "redrive_policy",
					Description: `Indicates whether to enable dead letter messages.`,
				},
				resource.Attribute{
					Name:        "produced_messages",
					Description: `Indicates the total number of messages (not including the messages that have expired and been deleted) in a queue.`,
				},
				resource.Attribute{
					Name:        "consumed_messages",
					Description: `Indicates the total number of messages that are successfully consumed.`,
				},
				resource.Attribute{
					Name:        "available_messages",
					Description: `Indicates the accumulated number of messages that can be consumed.`,
				},
				resource.Attribute{
					Name:        "produced_deadletters",
					Description: `Indicates the total number of dead letter messages generated by the consumer group.`,
				},
				resource.Attribute{
					Name:        "available_deadletters",
					Description: `Indicates the accumulated number of dead letter messages that have not been consumed.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_dms_instance_v1",
			Category:         "Distributed Message Service (DMS)",
			ShortDescription: `Manages a DMS instance in the huaweicloud DMS Service`,
			Description:      ``,
			Keywords: []string{
				"distributed",
				"message",
				"service",
				"dms",
				"instance",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Indicates the name of an instance. An instance name starts with a letter, consists of 4 to 64 characters, and supports only letters, digits, and hyphens (-).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Indicates the description of an instance. It is a character string containing not more than 1024 characters.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `(Optional) Indicates a message engine. Options: rabbitmq and kafka.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `(Optional) Indicates the version of a message engine.`,
				},
				resource.Attribute{
					Name:        "specification",
					Description: `(Optional) This parameter is mandatory if the engine is kafka. Indicates the baseline bandwidth of a Kafka instance, that is, the maximum amount of data transferred per unit time. Unit: byte/s. Options: 300 MB, 600 MB, 1200 MB.`,
				},
				resource.Attribute{
					Name:        "storage_space",
					Description: `(Required) Indicates the message storage space. Value range: Single-node RabbitMQ instance: 100–90000 GB Cluster RabbitMQ instance: 100 GB x Number of nodes to 90000 GB, 200 GB x Number of nodes to 90000 GB, 300 GB x Number of nodes to 90000 GB Kafka instance with specification being 300 MB: 1200–90000 GB Kafka instance with specification being 600 MB: 2400–90000 GB Kafka instance with specification being 1200 MB: 4800–90000 GB`,
				},
				resource.Attribute{
					Name:        "partition_num",
					Description: `(Optional) This parameter is mandatory when a Kafka instance is created. Indicates the maximum number of topics in a Kafka instance. When specification is 300 MB: 900 When specification is 600 MB: 1800 When specification is 1200 MB: 1800`,
				},
				resource.Attribute{
					Name:        "access_user",
					Description: `(Optional) Indicates a username. If the engine is rabbitmq, this parameter is mandatory. If the engine is kafka, this parameter is optional. A username consists of 4 to 64 characters and supports only letters, digits, and hyphens (-).`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) If the engine is rabbitmq, this parameter is mandatory. If the engine is kafka, this parameter is mandatory when ssl_enable is true and is invalid when ssl_enable is false. Indicates the password of an instance. An instance password must meet the following complexity requirements: Must be 8 to 32 characters long. Must contain at least 2 of the following character types: lowercase letters, uppercase letters, digits, and special characters (` + "`" + `~!@#$%^&`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) Indicates the ID of a VPC.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Required) Indicates the ID of a security group.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) Indicates the ID of a subnet.`,
				},
				resource.Attribute{
					Name:        "available_zones",
					Description: `(Required) Indicates the ID of an AZ. The parameter value can not be left blank or an empty array. For details, see section Querying AZ Information.`,
				},
				resource.Attribute{
					Name:        "product_id",
					Description: `(Required) Indicates a product ID.`,
				},
				resource.Attribute{
					Name:        "maintain_begin",
					Description: `(Optional) Indicates the time at which a maintenance time window starts. Format: HH:mm:ss. The start time and end time of a maintenance time window must indicate the time segment of a supported maintenance time window. For details, see section Querying Maintenance Time Windows. The start time must be set to 22:00, 02:00, 06:00, 10:00, 14:00, or 18:00. Parameters maintain_begin and maintain_end must be set in pairs. If parameter maintain_begin is left blank, parameter maintain_end is also blank. In this case, the system automatically allocates the default start time 02:00.`,
				},
				resource.Attribute{
					Name:        "maintain_end",
					Description: `(Optional) Indicates the time at which a maintenance time window ends. Format: HH:mm:ss. The start time and end time of a maintenance time window must indicate the time segment of a supported maintenance time window. For details, see section Querying Maintenance Time Windows. The end time is four hours later than the start time. For example, if the start time is 22:00, the end time is 02:00. Parameters maintain_begin and maintain_end must be set in pairs. If parameter maintain_end is left blank, parameter maintain_begin is also blank. In this case, the system automatically allocates the default end time 06:00.`,
				},
				resource.Attribute{
					Name:        "enable_publicip",
					Description: `(Optional) Indicates whether to enable public access to a RabbitMQ instance. true: enable, false: disable`,
				},
				resource.Attribute{
					Name:        "publicip_id",
					Description: `(Optional) Indicates the ID of the elastic IP address (EIP) bound to a RabbitMQ instance. This parameter is mandatory if public access is enabled (that is, enable_publicip is set to true).`,
				},
				resource.Attribute{
					Name:        "storage_spec_code",
					Description: `(Optional) Indicates the storage I/O specification. For details on how to select a disk type, see Disk Types and Disk Performance. Options for a RabbitMQ instance: dms.physical.storage.normal dms.physical.storage.high dms.physical.storage.ultra Options for a Kafka instance: When specification is 300 MB: dms.physical.storage.high or dms.physical.storage.ultra When specification is 600 MB: dms.physical.storage.ultra When specification is 1200 MB: dms.physical.storage.ultra ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "specification",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "storage_space",
					Description: `Indicates the time when a instance is created.`,
				},
				resource.Attribute{
					Name:        "partition_num",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "access_user",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `Indicates the name of a security group.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subnet_name",
					Description: `Indicates the name of a subnet.`,
				},
				resource.Attribute{
					Name:        "subnet_cidr",
					Description: `Indicates a subnet segment.`,
				},
				resource.Attribute{
					Name:        "available_zones",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "product_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "maintain_begin",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "maintain_end",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "enable_publicip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "publicip_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "storage_spec_code",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "used_storage_space",
					Description: `Indicates the used message storage space. Unit: GB`,
				},
				resource.Attribute{
					Name:        "connect_address",
					Description: `Indicates the IP address of an instance.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Indicates the port number of an instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates the status of an instance. For details, see section Instance Status.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `Indicates the ID of an instance.`,
				},
				resource.Attribute{
					Name:        "resource_spec_code",
					Description: `Indicates a resource specifications identifier.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Indicates an instance type. Options: "single" and "cluster"`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Indicates the time when an instance is created. The time is in the format of timestamp, that is, the offset milliseconds from 1970-01-01 00:00:00 UTC to the specified time.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `Indicates a user ID.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `Indicates a username.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "specification",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "storage_space",
					Description: `Indicates the time when a instance is created.`,
				},
				resource.Attribute{
					Name:        "partition_num",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "access_user",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `Indicates the name of a security group.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subnet_name",
					Description: `Indicates the name of a subnet.`,
				},
				resource.Attribute{
					Name:        "subnet_cidr",
					Description: `Indicates a subnet segment.`,
				},
				resource.Attribute{
					Name:        "available_zones",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "product_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "maintain_begin",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "maintain_end",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "enable_publicip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "publicip_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "storage_spec_code",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "used_storage_space",
					Description: `Indicates the used message storage space. Unit: GB`,
				},
				resource.Attribute{
					Name:        "connect_address",
					Description: `Indicates the IP address of an instance.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Indicates the port number of an instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates the status of an instance. For details, see section Instance Status.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `Indicates the ID of an instance.`,
				},
				resource.Attribute{
					Name:        "resource_spec_code",
					Description: `Indicates a resource specifications identifier.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Indicates an instance type. Options: "single" and "cluster"`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Indicates the time when an instance is created. The time is in the format of timestamp, that is, the offset milliseconds from 1970-01-01 00:00:00 UTC to the specified time.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `Indicates a user ID.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `Indicates a username.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_dms_queue_v1",
			Category:         "Distributed Message Service (DMS)",
			ShortDescription: `Manages a DMS queue in the huaweicloud DMS Service`,
			Description:      ``,
			Keywords: []string{
				"distributed",
				"message",
				"service",
				"dms",
				"queue",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Indicates the unique name of a queue. A string of 1 to 64 characters that contain a-z, A-Z, 0-9, hyphens (-), and underscores (_). The name cannot be modified once specified.`,
				},
				resource.Attribute{
					Name:        "queue_mode",
					Description: `(Optional) Indicates the queue type. It only support 'NORMAL' and 'FIFO'. NORMAL: Standard queue. Best-effort ordering. Messages might be retrieved in an order different from which they were sent. Select standard queues when throughput is important. FIFO: First-ln-First-out (FIFO) queue. FIFO delivery. Messages are retrieved in the order they were sent. Select FIFO queues when the order of messages is important. Default value: NORMAL.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Indicates the basic information about a queue. The queue description must be 0 to 160 characters in length, and does not contain angle brackets (<) and (>).`,
				},
				resource.Attribute{
					Name:        "redrive_policy",
					Description: `(Optional) Indicates whether to enable dead letter messages. Dead letter messages indicate messages that cannot be normally consumed. The redrive_policy should be set to 'enable' or 'disable'. The default value is 'disable'.`,
				},
				resource.Attribute{
					Name:        "max_consume_count",
					Description: `(Optional) This parameter is mandatory only when redrive_policy is set to enable. This parameter indicates the maximum number of allowed message consumption failures. When a message fails to be consumed after the number of consumption attempts of this message reaches this value, DMS stores this message into the dead letter queue. The max_consume_count value range is 1–100. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "queue_mode",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "redrive_policy",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "max_consume_count",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `Indicates the time when a queue is created.`,
				},
				resource.Attribute{
					Name:        "reservation",
					Description: `Indicates the retention period (unit: min) of a message in a queue.`,
				},
				resource.Attribute{
					Name:        "max_msg_size_byte",
					Description: `Indicates the maximum message size (unit: byte) that is allowed in queue.`,
				},
				resource.Attribute{
					Name:        "produced_messages",
					Description: `Indicates the total number of messages (not including the messages that have expired and been deleted) in a queue.`,
				},
				resource.Attribute{
					Name:        "group_count",
					Description: `Indicates the total number of consumer groups in a queue.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "queue_mode",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "redrive_policy",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "max_consume_count",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `Indicates the time when a queue is created.`,
				},
				resource.Attribute{
					Name:        "reservation",
					Description: `Indicates the retention period (unit: min) of a message in a queue.`,
				},
				resource.Attribute{
					Name:        "max_msg_size_byte",
					Description: `Indicates the maximum message size (unit: byte) that is allowed in queue.`,
				},
				resource.Attribute{
					Name:        "produced_messages",
					Description: `Indicates the total number of messages (not including the messages that have expired and been deleted) in a queue.`,
				},
				resource.Attribute{
					Name:        "group_count",
					Description: `Indicates the total number of consumer groups in a queue.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_dns_ptrrecord_v2",
			Category:         "Domain Name Service (DNS)",
			ShortDescription: `Manages a DNS PTR record in the HuaweiCloud DNS Service`,
			Description:      ``,
			Keywords: []string{
				"domain",
				"name",
				"service",
				"dns",
				"ptrrecord",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Domain name of the PTR record. A domain name is case insensitive. Uppercase letters will also be converted into lowercase letters.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the PTR record.`,
				},
				resource.Attribute{
					Name:        "floatingip_id",
					Description: `(Required) The ID of the FloatingIP/EIP.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The time to live (TTL) of the record set (in seconds). The value range is 300–2147483647. The default value is 300.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags key/value pairs to associate with the PTR record. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The PTR record ID, which is in {region}:{floatingip_id} format.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "floatingip_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The address of the FloatingIP/EIP. ## Import PTR records can be imported using region and floatingip/eip ID, separated by a colon(:), e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_dns_ptrrecord_v2.ptr_1 cn-north-1:d90ce693-5ccf-4136-a0ed-152ce412b6b9 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The PTR record ID, which is in {region}:{floatingip_id} format.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "floatingip_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The address of the FloatingIP/EIP. ## Import PTR records can be imported using region and floatingip/eip ID, separated by a colon(:), e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_dns_ptrrecord_v2.ptr_1 cn-north-1:d90ce693-5ccf-4136-a0ed-152ce412b6b9 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_dns_recordset_v2",
			Category:         "Domain Name Service (DNS)",
			ShortDescription: `Manages a DNS record set in the HuaweiCloud DNS Service`,
			Description:      ``,
			Keywords: []string{
				"domain",
				"name",
				"service",
				"dns",
				"recordset",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 DNS client. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new DNS record set.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Required) The ID of the zone in which to create the record set. Changing this creates a new DNS record set.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the record set. Note the ` + "`" + `.` + "`" + ` at the end of the name. Changing this creates a new DNS record set.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type of record set. The options include ` + "`" + `A` + "`" + `, ` + "`" + `AAAA` + "`" + `, ` + "`" + `MX` + "`" + `, ` + "`" + `CNAME` + "`" + `, ` + "`" + `TXT` + "`" + `, ` + "`" + `NS` + "`" + `, ` + "`" + `SRV` + "`" + `, and ` + "`" + `PTR` + "`" + `. Changing this creates a new DNS record set.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The time to live (TTL) of the record set (in seconds). The value range is 300–2147483647. The default value is 300.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the record set.`,
				},
				resource.Attribute{
					Name:        "records",
					Description: `(Required) An array of DNS records.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional options. Changing this creates a new record set. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "records",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `See Argument Reference above. ## Import This resource can be imported by specifying the zone ID and recordset ID, separated by a forward slash. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_dns_recordset_v2.recordset_1 <zone_id>/<recordset_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "records",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `See Argument Reference above. ## Import This resource can be imported by specifying the zone ID and recordset ID, separated by a forward slash. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_dns_recordset_v2.recordset_1 <zone_id>/<recordset_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_dns_zone_v2",
			Category:         "Domain Name Service (DNS)",
			ShortDescription: `Manages a DNS zone in the HuaweiCloud DNS Service`,
			Description:      ``,
			Keywords: []string{
				"domain",
				"name",
				"service",
				"dns",
				"zone",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Compute client. Keypairs are associated with accounts, but a Compute client is needed to create one. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new DNS zone. Changing this creates a new DNS zone.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the zone. Note the ` + "`" + `.` + "`" + ` at the end of the name. Changing this creates a new DNS zone.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) The email contact for the zone record.`,
				},
				resource.Attribute{
					Name:        "zone_type",
					Description: `(Optional) The type of zone. Can either be ` + "`" + `public` + "`" + ` or ` + "`" + `private` + "`" + `. Changing this creates a new DNS zone.`,
				},
				resource.Attribute{
					Name:        "router",
					Description: `(Optional) Router configuration block which is required if zone_type is private. The router structure is documented below.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The time to live (TTL) of the zone.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the zone.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional options. Changing this creates a new DNS zone. The ` + "`" + `router` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `(Required) The router UUID.`,
				},
				resource.Attribute{
					Name:        "router_region",
					Description: `(Required) The region of the router. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "zone_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "masters",
					Description: `An array of master DNS servers.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `See Argument Reference above. ## Import This resource can be imported by specifying the zone ID: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_dns_zone_v2.zone_1 <zone_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "zone_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "masters",
					Description: `An array of master DNS servers.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `See Argument Reference above. ## Import This resource can be imported by specifying the zone ID: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_dns_zone_v2.zone_1 <zone_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_dws_cluster",
			Category:         "Data Warehouse Service (DWS)",
			ShortDescription: `cluster management`,
			Description:      ``,
			Keywords: []string{
				"data",
				"warehouse",
				"service",
				"dws",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Cluster name, which must be unique and contains 4 to 64 characters, which consist of letters, digits, hyphens (-), or underscores (_) only and must start with a letter.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required) Network ID, which is used for configuring cluster network`,
				},
				resource.Attribute{
					Name:        "node_type",
					Description: `(Required) Node type`,
				},
				resource.Attribute{
					Name:        "number_of_node",
					Description: `(Required) Number of nodes in a cluster. The value ranges from 3 to 32`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Required) ID of a security group. The ID is used for configuring cluster network`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `(Required) Administrator username for logging in to a data warehouse cluster The administrator username must: Consist of lowercase letters, digits, or underscores. Start with a lowercase letter or an underscore. Contain 1 to 63 characters. Cannot be a keyword of the DWS database.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) VPC ID, which is used for configuring cluster network`,
				},
				resource.Attribute{
					Name:        "user_pwd",
					Description: `(Required) Administrator password for logging in to a data warehouse cluster A password must conform to the following rules: Contains 8 to 32 characters. Cannot be the same as the username or the username written in reverse order. Contains three types of the following: Lowercase letters Uppercase letters Digits Special characters ~!@#%^&`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) AZ in a cluster`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Service port of a cluster (8000 to 10000). The default value is 8000`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `(Optional) A nested object resource Structure is documented below. The ` + "`" + `public_ip` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "eip_id",
					Description: `(Optional) EIP ID`,
				},
				resource.Attribute{
					Name:        "public_bind_type",
					Description: `(Optional) Binding type of an EIP. The value can be either of the following: auto_assign not_use bind_existing The default value is not_use. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `Cluster creation time. The format is ISO8601:YYYY-MM-DDThh:mm:ssZ`,
				},
				resource.Attribute{
					Name:        "endpoints",
					Description: `View the private network connection information about the cluster. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Cluster ID`,
				},
				resource.Attribute{
					Name:        "public_endpoints",
					Description: `Public network connection information about the cluster. If the value is not specified, the public network connection information is not used by default Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "recent_event",
					Description: `The recent event number`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Cluster status, which can be one of the following: CREATING AVAILABLE UNAVAILABLE CREATION FAILED`,
				},
				resource.Attribute{
					Name:        "sub_status",
					Description: `Sub-status of clusters in the AVAILABLE state. The value can be one of the following: NORMAL READONLY REDISTRIBUTING REDISTRIBUTION-FAILURE UNBALANCED UNBALANCED | READONLY DEGRADED DEGRADED | READONLY DEGRADED | UNBALANCED UNBALANCED | REDISTRIBUTING UNBALANCED | REDISTRIBUTION-FAILURE READONLY | REDISTRIBUTION-FAILURE UNBALANCED | READONLY | REDISTRIBUTION-FAILURE DEGRADED | REDISTRIBUTION-FAILURE DEGRADED | UNBALANCED | REDISTRIBUTION-FAILURE DEGRADED | UNBALANCED | READONLY | REDISTRIBUTION-FAILURE DEGRADED | UNBALANCED | READONLY`,
				},
				resource.Attribute{
					Name:        "task_status",
					Description: `Cluster management task. The value can be one of the following: RESTORING SNAPSHOTTING GROWING REBOOTING SETTING_CONFIGURATION CONFIGURING_EXT_DATASOURCE DELETING_EXT_DATASOURCE REBOOT_FAILURE RESIZE_FAILURE`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `Last modification time of a cluster. The format is ISO8601:YYYY-MM-DDThh:mm:ssZ`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Data warehouse version The ` + "`" + `endpoints` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "connect_info",
					Description: `(Optional) Private network connection information`,
				},
				resource.Attribute{
					Name:        "jdbc_url",
					Description: `(Optional) JDBC URL. The following is the default format: jdbc:postgresql://< connect_info>/<YOUR_DATABASE_NAME> The ` + "`" + `public_endpoints` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "jdbc_url",
					Description: `(Optional) JDBC URL. The following is the default format: jdbc:postgresql://< public_connect_info>/<YOUR_DATABASE_NAME>`,
				},
				resource.Attribute{
					Name:        "public_connect_info",
					Description: `(Optional) Public network connection information ## Timeouts This resource provides the following timeouts configuration options: - ` + "`" + `create` + "`" + ` - Default is 10 minute. - ` + "`" + `delete` + "`" + ` - Default is 10 minute. ## Import Cluster can be imported using the following format: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_dws_cluster.default {{ resource id}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created",
					Description: `Cluster creation time. The format is ISO8601:YYYY-MM-DDThh:mm:ssZ`,
				},
				resource.Attribute{
					Name:        "endpoints",
					Description: `View the private network connection information about the cluster. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Cluster ID`,
				},
				resource.Attribute{
					Name:        "public_endpoints",
					Description: `Public network connection information about the cluster. If the value is not specified, the public network connection information is not used by default Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "recent_event",
					Description: `The recent event number`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Cluster status, which can be one of the following: CREATING AVAILABLE UNAVAILABLE CREATION FAILED`,
				},
				resource.Attribute{
					Name:        "sub_status",
					Description: `Sub-status of clusters in the AVAILABLE state. The value can be one of the following: NORMAL READONLY REDISTRIBUTING REDISTRIBUTION-FAILURE UNBALANCED UNBALANCED | READONLY DEGRADED DEGRADED | READONLY DEGRADED | UNBALANCED UNBALANCED | REDISTRIBUTING UNBALANCED | REDISTRIBUTION-FAILURE READONLY | REDISTRIBUTION-FAILURE UNBALANCED | READONLY | REDISTRIBUTION-FAILURE DEGRADED | REDISTRIBUTION-FAILURE DEGRADED | UNBALANCED | REDISTRIBUTION-FAILURE DEGRADED | UNBALANCED | READONLY | REDISTRIBUTION-FAILURE DEGRADED | UNBALANCED | READONLY`,
				},
				resource.Attribute{
					Name:        "task_status",
					Description: `Cluster management task. The value can be one of the following: RESTORING SNAPSHOTTING GROWING REBOOTING SETTING_CONFIGURATION CONFIGURING_EXT_DATASOURCE DELETING_EXT_DATASOURCE REBOOT_FAILURE RESIZE_FAILURE`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `Last modification time of a cluster. The format is ISO8601:YYYY-MM-DDThh:mm:ssZ`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Data warehouse version The ` + "`" + `endpoints` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "connect_info",
					Description: `(Optional) Private network connection information`,
				},
				resource.Attribute{
					Name:        "jdbc_url",
					Description: `(Optional) JDBC URL. The following is the default format: jdbc:postgresql://< connect_info>/<YOUR_DATABASE_NAME> The ` + "`" + `public_endpoints` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "jdbc_url",
					Description: `(Optional) JDBC URL. The following is the default format: jdbc:postgresql://< public_connect_info>/<YOUR_DATABASE_NAME>`,
				},
				resource.Attribute{
					Name:        "public_connect_info",
					Description: `(Optional) Public network connection information ## Timeouts This resource provides the following timeouts configuration options: - ` + "`" + `create` + "`" + ` - Default is 10 minute. - ` + "`" + `delete` + "`" + ` - Default is 10 minute. ## Import Cluster can be imported using the following format: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_dws_cluster.default {{ resource id}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_ecs_instance_v1",
			Category:         "Elastic Cloud Server (ECS)",
			ShortDescription: `Manages a V1 ECS instance resource within HuaweiCloud.`,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"cloud",
				"server",
				"ecs",
				"instance",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the instance.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Required) The ID of the desired image for the server. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `(Required) The name of the desired flavor for the server. Changing this resizes the existing server.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional) The user data to provide when launching the instance. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The administrative password to assign to the server. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `(Optional) The name of a key pair to put on the server. The key pair must already be created and associated with the tenant's account. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) The ID of the desired VPC for the server. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "nics",
					Description: `(Optional) An array of one or more networks to attach to the instance. The nics object structure is documented below. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "system_disk_type",
					Description: `(Optional) The system disk type of the server. For HANA, HL1, and HL2 ECSs use co-p1 and uh-l1 disks. Changing this creates a new server. Available options are:`,
				},
				resource.Attribute{
					Name:        "system_disk_size",
					Description: `(Optional) The system disk size in GB, The value range is 1 to 1024. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "data_disks",
					Description: `(Optional) An array of one or more data disks to attach to the instance. The data_disks object structure is documented below. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Optional) An array of one or more security group names to associate with the server. Changing this results in adding/removing security groups from the existing server.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required) The availability zone in which to create the server. Please refer to https://developer.huaweicloud.com/endpoint for the values. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "charging_mode",
					Description: `(Optional) The charging mode of the instance. Valid options are: prePaid and postPaid, defaults to postPaid. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "period_unit",
					Description: `(Optional) The charging period unit of the instance. Valid options are: month and year, defaults to month. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) The charging period of the instance. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "auto_recovery",
					Description: `(Optional) Whether configure automatic recovery of an instance.`,
				},
				resource.Attribute{
					Name:        "delete_disks_on_termination",
					Description: `(Optional) Delete the data disks upon termination of the instance. Defaults to false. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "enterprise_project_id",
					Description: `(Optional) The enterprise project id. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags key/value pairs to associate with the instance.`,
				},
				resource.Attribute{
					Name:        "op_svc_userid",
					Description: `(Optional) User ID, required when using key_name. Changing this creates a new server. The ` + "`" + `nics` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required) The network UUID to attach to the server. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) Specifies a fixed IPv4 address to be used on this network. Changing this creates a new server. The ` + "`" + `data_disks` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The data disk type of the server. For HANA, HL1, and HL2 ECSs use co-p1 and uh-l1 disks. Changing this creates a new server. Available options are:`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) The size of the data disk in GB. The value range is 10 to 32768. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) Specifies the snapshot ID or ID of the original data disk contained in the full-ECS image. Changing this creates a new server. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the server.`,
				},
				resource.Attribute{
					Name:        "nics/mac_address",
					Description: `The MAC address of the NIC on that network.`,
				},
				resource.Attribute{
					Name:        "nics/port_id",
					Description: `The port ID of the NIC on that network. ## Import Instances can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_ecs_instance_v1.instance_1 d90ce693-5ccf-4136-a0ed-152ce412b6b9`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the server.`,
				},
				resource.Attribute{
					Name:        "nics/mac_address",
					Description: `The MAC address of the NIC on that network.`,
				},
				resource.Attribute{
					Name:        "nics/port_id",
					Description: `The port ID of the NIC on that network. ## Import Instances can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_ecs_instance_v1.instance_1 d90ce693-5ccf-4136-a0ed-152ce412b6b9`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_elb_backendecs",
			Category:         "Elastic Load Balance (Classic)",
			ShortDescription: `Manages an elastic loadbalancer backendecs resource within huawei cloud.`,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"load",
				"balance",
				"classic",
				"elb",
				"backendecs",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Required) Specifies the listener ID.`,
				},
				resource.Attribute{
					Name:        "server_id",
					Description: `(Required) Specifies the backend member ID.`,
				},
				resource.Attribute{
					Name:        "private_address",
					Description: `(Required) Specifies the private IP address of the backend member. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "server_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "private_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "public_address",
					Description: `Specifies the floating IP address assigned to the backend member.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Specifies the backend member ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Specifies the backend ECS status. The value is ACTIVE, PENDING, or ERROR.`,
				},
				resource.Attribute{
					Name:        "health_status",
					Description: `Specifies the health check status. The value is NORMAL, ABNORMAL, or UNAVAILABLE.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Specifies the time when information about the backend member was updated.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Specifies the time when the backend member was created.`,
				},
				resource.Attribute{
					Name:        "server_name",
					Description: `Specifies the backend member name.`,
				},
				resource.Attribute{
					Name:        "listeners",
					Description: `Specifies the listener to which the backend member belongs.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "listener_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "server_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "private_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "public_address",
					Description: `Specifies the floating IP address assigned to the backend member.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Specifies the backend member ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Specifies the backend ECS status. The value is ACTIVE, PENDING, or ERROR.`,
				},
				resource.Attribute{
					Name:        "health_status",
					Description: `Specifies the health check status. The value is NORMAL, ABNORMAL, or UNAVAILABLE.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Specifies the time when information about the backend member was updated.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Specifies the time when the backend member was created.`,
				},
				resource.Attribute{
					Name:        "server_name",
					Description: `Specifies the backend member name.`,
				},
				resource.Attribute{
					Name:        "listeners",
					Description: `Specifies the listener to which the backend member belongs.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_elb_healthcheck",
			Category:         "Elastic Load Balance (Classic)",
			ShortDescription: `Manages an elastic loadbalancer healthcheck resource within huawei cloud.`,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"load",
				"balance",
				"classic",
				"elb",
				"healthcheck",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Required) Specifies the ID of the listener to which the health check task belongs.`,
				},
				resource.Attribute{
					Name:        "healthcheck_protocol",
					Description: `(Optional) Specifies the protocol used for the health check. The value can be HTTP or TCP (case-insensitive).`,
				},
				resource.Attribute{
					Name:        "healthcheck_uri",
					Description: `(Optional) Specifies the URI for health check. This parameter is valid when healthcheck_ protocol is HTTP. The value is a string of 1 to 80 characters that must start with a slash (/) and can only contain letters, digits, and special characters, such as -/.%?#&.`,
				},
				resource.Attribute{
					Name:        "healthcheck_connect_port",
					Description: `(Optional) Specifies the port used for the health check. The value ranges from 1 to 65535.`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `(Optional) Specifies the threshold at which the health check result is success, that is, the number of consecutive successful health checks when the health check result of the backend server changes from fail to success. The value ranges from 1 to 10.`,
				},
				resource.Attribute{
					Name:        "unhealthy_threshold",
					Description: `(Optional) Specifies the threshold at which the health check result is fail, that is, the number of consecutive failed health checks when the health check result of the backend server changes from success to fail. The value ranges from 1 to 10.`,
				},
				resource.Attribute{
					Name:        "healthcheck_timeout",
					Description: `(Optional) Specifies the maximum timeout duration (s) for the health check. The value ranges from 1 to 50.`,
				},
				resource.Attribute{
					Name:        "healthcheck_interval",
					Description: `(Optional) Specifies the maximum interval (s) for health check. The value ranges from 1 to 5. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "healthcheck_protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "healthcheck_uri",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "healthcheck_connect_port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "unhealthy_threshold",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "healthcheck_timeout",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "healthcheck_interval",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Specifies the health check task ID.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Specifies the time when information about the health check task was updated.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Specifies the time when the health check task was created.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "listener_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "healthcheck_protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "healthcheck_uri",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "healthcheck_connect_port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "unhealthy_threshold",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "healthcheck_timeout",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "healthcheck_interval",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Specifies the health check task ID.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Specifies the time when information about the health check task was updated.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Specifies the time when the health check task was created.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_elb_listener",
			Category:         "Elastic Load Balance (Classic)",
			ShortDescription: `Manages an elastic loadbalancer listener resource within huawei cloud.`,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"load",
				"balance",
				"classic",
				"elb",
				"listener",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the load balancer name. The name is a string of 1 to 64 characters that consist of letters, digits, underscores (_), and hyphens (-).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Provides supplementary information about the listener. The value is a string of 0 to 128 characters and cannot be <>.`,
				},
				resource.Attribute{
					Name:        "loadbalancer_id",
					Description: `(Required) Specifies the ID of the load balancer to which the listener belongs.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Specifies the listening protocol used for layer 4 or 7. The value can be HTTP, TCP, HTTPS, or UDP.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) Specifies the listening port. The value ranges from 1 to 65535.`,
				},
				resource.Attribute{
					Name:        "backend_protocol",
					Description: `(Required) Specifies the backend protocol. If the value of protocol is UDP, the value of this parameter can only be UDP. The value can be HTTP, TCP, or UDP.`,
				},
				resource.Attribute{
					Name:        "backend_port",
					Description: `(Required) Specifies the backend port. The value ranges from 1 to 65535.`,
				},
				resource.Attribute{
					Name:        "lb_algorithm",
					Description: `(Required) Specifies the load balancing algorithm for the listener. The value can be roundrobin, leastconn, or source.`,
				},
				resource.Attribute{
					Name:        "session_sticky",
					Description: `(Optional) Specifies whether to enable sticky session. The value can be true or false. The Sticky session is enabled when the value is true, and is disabled when the value is false. If the value of protocol is HTTP, HTTPS, or TCP, and the value of lb_algorithm is not roundrobin, the value of this parameter can only be false.`,
				},
				resource.Attribute{
					Name:        "sticky_session_type",
					Description: `(Optional) Specifies the cookie processing method. The value is insert. insert indicates that the cookie is inserted by the load balancer. This parameter is valid when protocol is set to HTTP, and session_sticky to true. The default value is insert. This parameter is invalid when protocol is set to TCP or UDP, which means the parameter is empty.`,
				},
				resource.Attribute{
					Name:        "cookie_timeout",
					Description: `(Optional) Specifies the cookie timeout period (minutes). This parameter is valid when protocol is set to HTTP, session_sticky to true, and sticky_session_type to insert. This parameter is invalid when protocol is set to TCP or UDP. The value ranges from 1 to 1440.`,
				},
				resource.Attribute{
					Name:        "tcp_timeout",
					Description: `(Optional) Specifies the TCP timeout period (minutes). This parameter is valid when protocol is set to TCP. The value ranges from 1 to 5.`,
				},
				resource.Attribute{
					Name:        "tcp_draining",
					Description: `(Optional) Specifies whether to maintain the TCP connection to the backend ECS after the ECS is deleted. This parameter is valid when protocol is set to TCP. The value can be true or false.`,
				},
				resource.Attribute{
					Name:        "tcp_draining_timeout",
					Description: `(Optional) Specifies the timeout duration (minutes) for the TCP connection to the backend ECS after the ECS is deleted. This parameter is valid when protocol is set to TCP, and tcp_draining to true. The value ranges from 0 to 60.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `(Optional) Specifies the ID of the SSL certificate used for security authentication when HTTPS is used to make API calls. This parameter is mandatory if the value of protocol is HTTPS. The value can be obtained by viewing the details of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "udp_timeout",
					Description: `(Optional) Specifies the UDP timeout duration (minutes). This parameter is valid when protocol is set to UDP. The value ranges from 1 to 1440.`,
				},
				resource.Attribute{
					Name:        "ssl_protocols",
					Description: `(Optional) Specifies the SSL protocol standard supported by a tracker, which is used for enabling specified encryption protocols. This parameter is valid only when the value of protocol is set to HTTPS. The value is TLSv1.2 or TLSv1.2 TLSv1.1 TLSv1. The default value is TLSv1.2.`,
				},
				resource.Attribute{
					Name:        "ssl_ciphers",
					Description: `(Optional) Specifies the cipher suite of an encryption protocol. This parameter is valid only when the value of protocol is set to HTTPS. The value is Default, Extended, or Strict. The default value is Default. The value can only be set to Extended if the value of ssl_protocols is set to TLSv1.2 TLSv1.1 TLSv1. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "loadbalancer_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "backend_protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "backend_port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "lb_algorithm",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "session_sticky",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "sticky_session_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cookie_timeout",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tcp_timeout",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tcp_draining",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tcp_draining_timeout",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "udp_timeout",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ssl_protocols",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ssl_ciphers",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Specifies the time when information about the listener was updated.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Specifies the listener ID.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Specifies the time when the listener was created.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Specifies the listener status. The value can be ACTIVE, PENDING_CREATE, or ERROR.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `Specifies the status of the load balancer. Value range: false: The load balancer is disabled. true: The load balancer runs properly.`,
				},
				resource.Attribute{
					Name:        "member_number",
					Description: `Specifies the number of backend members.`,
				},
				resource.Attribute{
					Name:        "healthcheck_id",
					Description: `Specifies the health check task ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "loadbalancer_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "backend_protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "backend_port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "lb_algorithm",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "session_sticky",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "sticky_session_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cookie_timeout",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tcp_timeout",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tcp_draining",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tcp_draining_timeout",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "udp_timeout",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ssl_protocols",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ssl_ciphers",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Specifies the time when information about the listener was updated.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Specifies the listener ID.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Specifies the time when the listener was created.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Specifies the listener status. The value can be ACTIVE, PENDING_CREATE, or ERROR.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `Specifies the status of the load balancer. Value range: false: The load balancer is disabled. true: The load balancer runs properly.`,
				},
				resource.Attribute{
					Name:        "member_number",
					Description: `Specifies the number of backend members.`,
				},
				resource.Attribute{
					Name:        "healthcheck_id",
					Description: `Specifies the health check task ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_elb_loadbalancer",
			Category:         "Elastic Load Balance (Classic)",
			ShortDescription: `Manages an elastic loadbalancer resource within huawei cloud.`,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"load",
				"balance",
				"classic",
				"elb",
				"loadbalancer",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the load balancer name. The name is a string of 1 to 64 characters that consist of letters, digits, underscores (_), and hyphens (-).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Provides supplementary information about the listener. The value is a string of 0 to 128 characters and cannot be <>.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) Specifies the VPC ID.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Optional) Specifies the bandwidth (Mbit/s). This parameter is mandatory when type is set to External, and it is invalid when type is set to Internal. The value ranges from 1 to 300.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Specifies the load balancer type. The value can be Internal or External.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Required) Specifies the status of the load balancer. Value range: 0 or false: indicates that the load balancer is stopped. Only tenants are allowed to enter these two values. 1 or true: indicates that the load balancer is running properly. 2 or false: indicates that the load balancer is frozen. Only tenants are allowed to enter these two values.`,
				},
				resource.Attribute{
					Name:        "vip_subnet_id",
					Description: `(Optional) Specifies the ID of the private network to be added. This parameter is mandatory when type is set to Internal, and it is invalid when type is set to External.`,
				},
				resource.Attribute{
					Name:        "az",
					Description: `(Optional) Specifies the ID of the availability zone (AZ). This parameter is mandatory when type is set to Internal, and it is invalid when type is set to External.`,
				},
				resource.Attribute{
					Name:        "charge_mode",
					Description: `(Optional) This is a reserved field. If the system supports charging by traffic and this field is specified, then you are charged by traffic for elastic IP addresses. The value is traffic.`,
				},
				resource.Attribute{
					Name:        "eip_type",
					Description: `(Optional) This parameter is reserved.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Optional) Specifies the security group ID. The value is a string of 1 to 200 characters that consists of uppercase and lowercase letters, digits, and hyphens (-). This parameter is mandatory only when type is set to Internal.`,
				},
				resource.Attribute{
					Name:        "vip_address",
					Description: `(Optional) Specifies the IP address provided by ELB. When typeis set to External, the value of this parameter is the elastic IP address. When type is set to Internal, the value of this parameter is the private network IP address. You can select an existing elastic IP address and create a public network load balancer. When this parameter is configured, parameters bandwidth, charge_mode, and eip_type are invalid.`,
				},
				resource.Attribute{
					Name:        "tenantid",
					Description: `(Optional) Specifies the tenant ID. This parameter is mandatory only when type is set to Internal. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vip_subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "az",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "charge_mode",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "eip_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vip_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenantid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Specifies the time when information about the load balancer was updated.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Specifies the time when the load balancer was created.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Specifies the load balancer ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Specifies the status of the load balancer. The value can be ACTIVE, PENDING_CREATE, or ERROR.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vip_subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "az",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "charge_mode",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "eip_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vip_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenantid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Specifies the time when information about the load balancer was updated.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Specifies the time when the load balancer was created.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Specifies the load balancer ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Specifies the status of the load balancer. The value can be ACTIVE, PENDING_CREATE, or ERROR.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_evs_snapshot",
			Category:         "Elastic Volume Service (EVS)",
			ShortDescription: `Provides an EVS snapshot resource.`,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"volume",
				"service",
				"evs",
				"snapshot",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_fgs_function_v2",
			Category:         "FunctionGraph",
			ShortDescription: `Manages a V2 function resource within HuaweiCloud.`,
			Description:      ``,
			Keywords: []string{
				"functiongraph",
				"fgs",
				"function",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the function. Changing this creates a new function.`,
				},
				resource.Attribute{
					Name:        "package",
					Description: `(Required) Group to which the function belongs. Changing this creates a new function.`,
				},
				resource.Attribute{
					Name:        "code_type",
					Description: `(Required) Function code type, which can be inline: inline code, zip: ZIP file, jar: JAR file or java functions, obs: function code stored in an OBS bucket. Changing this creates a new function.`,
				},
				resource.Attribute{
					Name:        "code_url",
					Description: `(Optional) This parameter is mandatory when code_type is set to obs. Changing this creates a new function.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the function. Changing this creates a new function.`,
				},
				resource.Attribute{
					Name:        "code_filename",
					Description: `(Optional) Name of a function file, This field is mandatory only when coe_type is set to jar or zip. Changing this creates a new function.`,
				},
				resource.Attribute{
					Name:        "handler",
					Description: `(Required) Entry point of the function. Changing this creates a new function.`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `(Required) Memory size(MB) allocated to the function. Changing this creates a new function.`,
				},
				resource.Attribute{
					Name:        "runtime",
					Description: `(Required) Environment for executing the function. Changing this creates a new function.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Required) Timeout interval of the function, ranges from 3s to 900s. Changing this creates a new function.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional) Key/Value information defined for the function. Changing this creates a new function.`,
				},
				resource.Attribute{
					Name:        "xrole",
					Description: `(Optional) This parameter is mandatory if the function needs to access other cloud services. Changing this creates a new function.`,
				},
				resource.Attribute{
					Name:        "func_code",
					Description: `(Required) Function code. When code_type is set to inline, zip, or jar, this parameter is mandatory, and the code must be encoded using Base64. Changing this creates a new function. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "package",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "code_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "code_url",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "code_filename",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "handler",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "runtime",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "xrole",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "func_code",
					Description: `See Argument Reference above. ## Import Functions can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_fgs_function_v2.my-func 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "package",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "code_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "code_url",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "code_filename",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "handler",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "runtime",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "xrole",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "func_code",
					Description: `See Argument Reference above. ## Import Functions can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_fgs_function_v2.my-func 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_fw_firewall_group_v2",
			Category:         "Network ACL",
			ShortDescription: `Manages a v2 firewall group resource within HuaweiCloud.`,
			Description:      ``,
			Keywords: []string{
				"network",
				"acl",
				"fw",
				"firewall",
				"group",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the v2 networking client. A networking client is needed to create a firewall group. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new firewall group.`,
				},
				resource.Attribute{
					Name:        "ingress_policy_id",
					Description: `The ingress policy resource id for the firewall group. Changing this updates the ` + "`" + `ingress_policy_id` + "`" + ` of an existing firewall group.`,
				},
				resource.Attribute{
					Name:        "egress_policy_id",
					Description: `The egress policy resource id for the firewall group. Changing this updates the ` + "`" + `egress_policy_id` + "`" + ` of an existing firewall group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A name for the firewall group. Changing this updates the ` + "`" + `name` + "`" + ` of an existing firewall group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) A description for the firewall group. Changing this updates the ` + "`" + `description` + "`" + ` of an existing firewall group.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) Administrative up/down status for the firewall group (must be "true" or "false" if provided - defaults to "true"). Changing this updates the ` + "`" + `admin_state_up` + "`" + ` of an existing firewall group.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the floating IP. Required if admin wants to create a firewall group for another tenant. Changing this creates a new firewall group.`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `(Optional) Port(s) to associate this firewall group instance with. Must be a list of strings. Changing this updates the associated routers of an existing firewall group.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional options. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `See Argument Reference above. ## Import Firewall Groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_fw_firewall_group_v2.firewall_group_1 c9e39fb2-ce20-46c8-a964-25f3898c7a97 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `See Argument Reference above. ## Import Firewall Groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_fw_firewall_group_v2.firewall_group_1 c9e39fb2-ce20-46c8-a964-25f3898c7a97 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_fw_policy_v2",
			Category:         "Network ACL",
			ShortDescription: `Manages a v2 firewall policy resource within HuaweiCloud.`,
			Description:      ``,
			Keywords: []string{
				"network",
				"acl",
				"fw",
				"policy",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the v2 networking client. A networking client is needed to create a firewall policy. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new firewall policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A name for the firewall policy. Changing this updates the ` + "`" + `name` + "`" + ` of an existing firewall policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description for the firewall policy. Changing this updates the ` + "`" + `description` + "`" + ` of an existing firewall policy.`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `(Optional) An array of one or more firewall rules that comprise the policy. Changing this results in adding/removing rules from the existing firewall policy.`,
				},
				resource.Attribute{
					Name:        "audited",
					Description: `(Optional) Audit status of the firewall policy (must be "true" or "false" if provided - defaults to "false"). This status is set to "false" whenever the firewall policy or any of its rules are changed. Changing this updates the ` + "`" + `audited` + "`" + ` status of an existing firewall policy.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `(Optional) Sharing status of the firewall policy (must be "true" or "false" if provided). If this is "true" the policy is visible to, and can be used in, firewalls in other tenants. Changing this updates the ` + "`" + `shared` + "`" + ` status of an existing firewall policy. Only administrative users can specify if the policy should be shared.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the firewall policy. Required if admin wants to create a firewall policy for another tenant. Changing this creates a new firewall policy.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional options. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "audited",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above. ## Import Firewall Policies can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_fw_policy_v2.policy_1 07f422e6-c596-474b-8b94-fe2c12506ce0 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "audited",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above. ## Import Firewall Policies can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_fw_policy_v2.policy_1 07f422e6-c596-474b-8b94-fe2c12506ce0 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_fw_rule_v2",
			Category:         "Network ACL",
			ShortDescription: `Manages a v2 firewall group rule resource within HuaweiCloud.`,
			Description:      ``,
			Keywords: []string{
				"network",
				"acl",
				"fw",
				"rule",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the v2 Networking client. A Compute client is needed to create a firewall rule. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new firewall rule.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A unique name for the firewall rule. Changing this updates the ` + "`" + `name` + "`" + ` of an existing firewall rule.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description for the firewall rule. Changing this updates the ` + "`" + `description` + "`" + ` of an existing firewall rule.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The protocol type on which the firewall rule operates. Valid values are: ` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `, and ` + "`" + `any` + "`" + `. Changing this updates the ` + "`" + `protocol` + "`" + ` of an existing firewall rule.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) Action to be taken ( must be "allow" or "deny") when the firewall rule matches. Changing this updates the ` + "`" + `action` + "`" + ` of an existing firewall rule.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `(Optional) IP version, either 4 (default) or 6. Changing this updates the ` + "`" + `ip_version` + "`" + ` of an existing firewall rule.`,
				},
				resource.Attribute{
					Name:        "source_ip_address",
					Description: `(Optional) The source IP address on which the firewall rule operates. Changing this updates the ` + "`" + `source_ip_address` + "`" + ` of an existing firewall rule.`,
				},
				resource.Attribute{
					Name:        "destination_ip_address",
					Description: `(Optional) The destination IP address on which the firewall rule operates. Changing this updates the ` + "`" + `destination_ip_address` + "`" + ` of an existing firewall rule.`,
				},
				resource.Attribute{
					Name:        "source_port",
					Description: `(Optional) The source port on which the firewall rule operates. Changing this updates the ` + "`" + `source_port` + "`" + ` of an existing firewall rule.`,
				},
				resource.Attribute{
					Name:        "destination_port",
					Description: `(Optional) The destination port on which the firewall rule operates. Changing this updates the ` + "`" + `destination_port` + "`" + ` of an existing firewall rule.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Enabled status for the firewall rule (must be "true" or "false" if provided - defaults to "true"). Changing this updates the ` + "`" + `enabled` + "`" + ` status of an existing firewall rule.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the firewall rule. Required if admin wants to create a firewall rule for another tenant. Changing this creates a new firewall rule.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional options. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "source_ip_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "destination_ip_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "source_port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "destination_port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above. ## Import Firewall Rules can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_fw_rule_v2.rule_1 8dbc0c28-e49c-463f-b712-5c5d1bbac327 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "source_ip_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "destination_ip_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "source_port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "destination_port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above. ## Import Firewall Rules can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_fw_rule_v2.rule_1 8dbc0c28-e49c-463f-b712-5c5d1bbac327 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_ges_graph_v1",
			Category:         "Graph Engine Service (GES)",
			ShortDescription: `graph management`,
			Description:      ``,
			Keywords: []string{
				"graph",
				"engine",
				"service",
				"ges",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required) Indicates availability zone. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "graph_size_type",
					Description: `(Required) Indicates the graph size type. 0: indicates 10 thousand edges. 1: indicates 1 million edges. 2: indicates 10 million edges. 3: indicates 100 million edges. 4: indicates 1 billion edges. 5: indicates 10 billion edges. 6: indicates 100 billion edges. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Indicates the graph name. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Indicates the region code. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Required) Indicates the security group ID. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) Indicates the subnet ID in the specified VPC. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) Indicates the VPC ID. Changing this parameter will create a new resource. - - -`,
				},
				resource.Attribute{
					Name:        "auto_assign",
					Description: `(Optional) Indicates whether to assign a new eip to the graph automatically. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "eip_id",
					Description: `(Optional) Indicates the ID of an EIP. Changing this parameter will create a new resource. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `Indicates the time when a graph is created.`,
				},
				resource.Attribute{
					Name:        "edgeset_path",
					Description: `Indicates the OBS path of the edge data set. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Indicates the private network access address of a graph instance. Users can access the instance using the IP address through the ECS deployed on the private network.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Indicates the public network access address of a graph instance. Users can access the instance using the IP address from the Internet.`,
				},
				resource.Attribute{
					Name:        "schema_path",
					Description: `Indicates the path for storing the metadata file. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Indicates the graph version.`,
				},
				resource.Attribute{
					Name:        "vertexset_path",
					Description: `Indicates the OBS path of the vertex data set. Structure is documented below. The ` + "`" + `edgeset_path` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `Indicates the OBS storage path, excluding OBS endpoint.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates the OBS file import status: success: Imported successfully. partiallyFailed: Partially failed. failed: Failed to import the file. The ` + "`" + `schema_path` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `Indicates the OBS storage path, excluding OBS endpoint.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates the OBS file import status: success: Imported successfully. partiallyFailed: Partially failed. failed: Failed to import the file. The ` + "`" + `vertexset_path` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `Indicates the OBS storage path, excluding OBS endpoint.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates the OBS file import status: success: Imported successfully. partiallyFailed: Partially failed. failed: Failed to import the file. ## Timeouts This resource provides the following timeouts configuration options: - ` + "`" + `create` + "`" + ` - Default is 30 minute. - ` + "`" + `delete` + "`" + ` - Default is 30 minute.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created",
					Description: `Indicates the time when a graph is created.`,
				},
				resource.Attribute{
					Name:        "edgeset_path",
					Description: `Indicates the OBS path of the edge data set. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Indicates the private network access address of a graph instance. Users can access the instance using the IP address through the ECS deployed on the private network.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Indicates the public network access address of a graph instance. Users can access the instance using the IP address from the Internet.`,
				},
				resource.Attribute{
					Name:        "schema_path",
					Description: `Indicates the path for storing the metadata file. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Indicates the graph version.`,
				},
				resource.Attribute{
					Name:        "vertexset_path",
					Description: `Indicates the OBS path of the vertex data set. Structure is documented below. The ` + "`" + `edgeset_path` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `Indicates the OBS storage path, excluding OBS endpoint.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates the OBS file import status: success: Imported successfully. partiallyFailed: Partially failed. failed: Failed to import the file. The ` + "`" + `schema_path` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `Indicates the OBS storage path, excluding OBS endpoint.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates the OBS file import status: success: Imported successfully. partiallyFailed: Partially failed. failed: Failed to import the file. The ` + "`" + `vertexset_path` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `Indicates the OBS storage path, excluding OBS endpoint.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates the OBS file import status: success: Imported successfully. partiallyFailed: Partially failed. failed: Failed to import the file. ## Timeouts This resource provides the following timeouts configuration options: - ` + "`" + `create` + "`" + ` - Default is 30 minute. - ` + "`" + `delete` + "`" + ` - Default is 30 minute.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_iam_agency_v3",
			Category:         "Identity and Access Management (IAM)",
			ShortDescription: `Manages an agency resource within huawei cloud.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"and",
				"access",
				"management",
				"iam",
				"agency",
				"v3",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of agency. The name is a string of 1 to 64 characters.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Provides supplementary information about the agency. The value is a string of 0 to 255 characters.`,
				},
				resource.Attribute{
					Name:        "delegated_domain_name",
					Description: `(Required) The name of delegated domain.`,
				},
				resource.Attribute{
					Name:        "project_role",
					Description: `(Optional) An array of roles and projects which are used to grant permissions to agency on project. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "domain_roles",
					Description: `(optional) An array of role names which stand for the permissionis to be granted to agency on domain. The ` + "`" + `project_role` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Required) The name of project`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `(Required) An array of role names`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "delegated_domain_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_role",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "domain_roles",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `Validity period of an agency. The default value is null, indicating that the agency is permanently valid.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `The expiration time of agency`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time when the agency was created.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The agency ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "delegated_domain_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_role",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "domain_roles",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `Validity period of an agency. The default value is null, indicating that the agency is permanently valid.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `The expiration time of agency`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time when the agency was created.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The agency ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_identity_group_membership_v3",
			Category:         "Identity and Access Management (IAM)",
			ShortDescription: `Manages the membership combine User Group resource and User resource within HuaweiCloud IAM service.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"and",
				"access",
				"management",
				"iam",
				"group",
				"membership",
				"v3",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group",
					Description: `(Required) The group ID of this membership.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `(Required) A List of user IDs to associate to the group. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "group",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_identity_group_v3",
			Category:         "Identity and Access Management (IAM)",
			ShortDescription: `Manages a User Group resource within HuaweiCloud IAM service.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"and",
				"access",
				"management",
				"iam",
				"group",
				"v3",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the group.The length is less than or equal to 64 bytes`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the group.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Optional) The domain this group belongs to.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V3 Keystone client. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new User Group. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `See Argument Reference above. ## Import Groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_identity_group_v3.group_1 89c60255-9bd6-460c-822a-e2b959ede9d2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "domain_id",
					Description: `See Argument Reference above. ## Import Groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_identity_group_v3.group_1 89c60255-9bd6-460c-822a-e2b959ede9d2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_identity_project_v3",
			Category:         "Identity and Access Management (IAM)",
			ShortDescription: `Manages a Project resource within HuaweiCloud Keystone.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"and",
				"access",
				"management",
				"iam",
				"project",
				"v3",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the project. it must start with ID of an existing region_ and be less than or equal to 64 characters. Example: eu-de_project1.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the project.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Optional) The domain this project belongs to. Changing this creates a new Project.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `(Optional) The parent of this project. Changing this creates a new Project.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the IAM client. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new Project. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `See Argument Reference above. ## Import Projects can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_identity_project_v3.project_1 89c60255-9bd6-460c-822a-e2b959ede9d2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "domain_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `See Argument Reference above. ## Import Projects can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_identity_project_v3.project_1 89c60255-9bd6-460c-822a-e2b959ede9d2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_identity_role_assignment_v3",
			Category:         "Identity and Access Management (IAM)",
			ShortDescription: `Manages a V3 Policy assignment within HuaweiCloud IAM Service.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"and",
				"access",
				"management",
				"iam",
				"role",
				"assignment",
				"v3",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Optional; Required if ` + "`" + `project_id` + "`" + ` is empty) The domain to assign the role in.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Optional; Required if ` + "`" + `user_id` + "`" + ` is empty) The group to assign the role to.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional; Required if ` + "`" + `domain_id` + "`" + ` is empty) The project to assign the role in.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `(Required) The role to assign. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "domain_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_identity_user_v3",
			Category:         "Identity and Access Management (IAM)",
			ShortDescription: `Manages a User resource within HuaweiCloud IAM service.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"and",
				"access",
				"management",
				"iam",
				"user",
				"v3",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the user. The user name consists of 5 to 32 characters. It can contain only uppercase letters, lowercase letters, digits, spaces, and special characters (-_) and cannot start with a digit.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the user.`,
				},
				resource.Attribute{
					Name:        "default_project_id",
					Description: `(Optional) The default project this user belongs to.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Optional) The domain this user belongs to.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Whether the user is enabled or disabled. Valid values are ` + "`" + `true` + "`" + ` and ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The password for the user. It must contain at least two of the following character types: uppercase letters, lowercase letters, digits, and special characters.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V3 Keystone client. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new User. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `See Argument Reference above. ## Import Users can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_identity_user_v3.user_1 89c60255-9bd6-460c-822a-e2b959ede9d2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "domain_id",
					Description: `See Argument Reference above. ## Import Users can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_identity_user_v3.user_1 89c60255-9bd6-460c-822a-e2b959ede9d2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_images_image_v2",
			Category:         "Image Management Service (IMS)",
			ShortDescription: `Manages a V2 Image resource within HuaweiCloud IMS.`,
			Description:      ``,
			Keywords: []string{
				"image",
				"management",
				"service",
				"ims",
				"images",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "container_format",
					Description: `(Required) The container format. Must be "bare".`,
				},
				resource.Attribute{
					Name:        "disk_format",
					Description: `(Required) The disk format. Must be one of "qcow2", "vhd".`,
				},
				resource.Attribute{
					Name:        "local_file_path",
					Description: `(Optional) This is the filepath of the raw image file that will be uploaded to Glance. Conflicts with ` + "`" + `image_source_url` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "image_cache_path",
					Description: `(Optional) This is the directory where the images will be downloaded. Images will be stored with a filename corresponding to the url's md5 hash. Defaults to "$HOME/.terraform/image_cache"`,
				},
				resource.Attribute{
					Name:        "image_source_url",
					Description: `(Optional) This is the url of the raw image that will be downloaded in the ` + "`" + `image_cache_path` + "`" + ` before being uploaded to Glance. Glance is able to download image from internet but the ` + "`" + `golangsdk` + "`" + ` library does not yet provide a way to do so. Conflicts with ` + "`" + `local_file_path` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "min_disk_gb",
					Description: `(Optional) Amount of disk space (in GB) required to boot image. Defaults to 0.`,
				},
				resource.Attribute{
					Name:        "min_ram_mb",
					Description: `(Optional) Amount of ram (in MB) required to boot image. Defauts to 0.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the image.`,
				},
				resource.Attribute{
					Name:        "protected",
					Description: `(Optional) If true, image will not be deletable. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Glance client. A Glance client is needed to create an Image that can be used with a compute instance. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new Image.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The tags of the image. It must be a list of strings. At this time, it is not possible to delete all tags of an image.`,
				},
				resource.Attribute{
					Name:        "visibility",
					Description: `(Optional) The visibility of the image. Must be "private". The ability to set the visibility depends upon the configuration of the HuaweiCloud cloud. Note: The ` + "`" + `properties` + "`" + ` attribute handling in the golangsdk library is currently buggy and needs to be fixed before being implemented in this resource. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "checksum",
					Description: `The checksum of the data associated with the image.`,
				},
				resource.Attribute{
					Name:        "container_format",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date the image was created.`,
				},
				resource.Attribute{
					Name:        "disk_format",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "file",
					Description: `the trailing path after the glance endpoint that represent the location of the image or the path to retrieve it.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A unique ID assigned by Glance.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `The metadata associated with the image. Image metadata allow for meaningfully define the image properties and tags. See http://docs.openstack.org/developer/glance/metadefs-concepts.html.`,
				},
				resource.Attribute{
					Name:        "min_disk_gb",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "min_ram_mb",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `The id of the huaweicloud user who owns the image.`,
				},
				resource.Attribute{
					Name:        "protected",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "schema",
					Description: `The path to the JSON-schema that represent the image or image`,
				},
				resource.Attribute{
					Name:        "size_bytes",
					Description: `The size in bytes of the data associated with the image.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the image. It can be "queued", "active" or "saving".`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "update_at",
					Description: `The date the image was last updated.`,
				},
				resource.Attribute{
					Name:        "visibility",
					Description: `See Argument Reference above. ## Import Images can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_images_image_v2.rancheros 89c60255-9bd6-460c-822a-e2b959ede9d2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "checksum",
					Description: `The checksum of the data associated with the image.`,
				},
				resource.Attribute{
					Name:        "container_format",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date the image was created.`,
				},
				resource.Attribute{
					Name:        "disk_format",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "file",
					Description: `the trailing path after the glance endpoint that represent the location of the image or the path to retrieve it.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A unique ID assigned by Glance.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `The metadata associated with the image. Image metadata allow for meaningfully define the image properties and tags. See http://docs.openstack.org/developer/glance/metadefs-concepts.html.`,
				},
				resource.Attribute{
					Name:        "min_disk_gb",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "min_ram_mb",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `The id of the huaweicloud user who owns the image.`,
				},
				resource.Attribute{
					Name:        "protected",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "schema",
					Description: `The path to the JSON-schema that represent the image or image`,
				},
				resource.Attribute{
					Name:        "size_bytes",
					Description: `The size in bytes of the data associated with the image.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the image. It can be "queued", "active" or "saving".`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "update_at",
					Description: `The date the image was last updated.`,
				},
				resource.Attribute{
					Name:        "visibility",
					Description: `See Argument Reference above. ## Import Images can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_images_image_v2.rancheros 89c60255-9bd6-460c-822a-e2b959ede9d2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_kms_key_v1",
			Category:         "Key Management Service (KMS)",
			ShortDescription: `Manages a V1 key resource within KMS.`,
			Description:      ``,
			Keywords: []string{
				"key",
				"management",
				"service",
				"kms",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key_alias",
					Description: `(Required) The alias in which to create the key. It is required when we create a new key. Changing this updates the alias of key.`,
				},
				resource.Attribute{
					Name:        "key_description",
					Description: `(Optional) The description of the key as viewed in Huawei console. Changing this updates the description of key.`,
				},
				resource.Attribute{
					Name:        "realm",
					Description: `(Optional) Region where a key resides. Changing this creates a new key.`,
				},
				resource.Attribute{
					Name:        "pending_days",
					Description: `(Optional) Duration in days after which the key is deleted after destruction of the resource, must be between 7 and 1096 days. It doesn't have default value. It only be used when delete a key.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `(Optional) Specifies whether the key is enabled. Defaults to true. Changing this updates the state of existing key. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "key_alias",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "key_description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "realm",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `The globally unique identifier for the key.`,
				},
				resource.Attribute{
					Name:        "default_key_flag",
					Description: `Identification of a Master Key. The value 1 indicates a Default Master Key, and the value 0 indicates a key.`,
				},
				resource.Attribute{
					Name:        "scheduled_deletion_date",
					Description: `Scheduled deletion time (time stamp) of a key.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `ID of a user domain for the key.`,
				},
				resource.Attribute{
					Name:        "expiration_time",
					Description: `Expiration time.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `Creation time (time stamp) of a key.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `See Argument Reference above. ## Import KMS Keys can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_kms_key_v1.key_1 7056d636-ac60-4663-8a6c-82d3c32c1c64 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "key_alias",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "key_description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "realm",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `The globally unique identifier for the key.`,
				},
				resource.Attribute{
					Name:        "default_key_flag",
					Description: `Identification of a Master Key. The value 1 indicates a Default Master Key, and the value 0 indicates a key.`,
				},
				resource.Attribute{
					Name:        "scheduled_deletion_date",
					Description: `Scheduled deletion time (time stamp) of a key.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `ID of a user domain for the key.`,
				},
				resource.Attribute{
					Name:        "expiration_time",
					Description: `Expiration time.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `Creation time (time stamp) of a key.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `See Argument Reference above. ## Import KMS Keys can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_kms_key_v1.key_1 7056d636-ac60-4663-8a6c-82d3c32c1c64 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_lb_certificate_v2",
			Category:         "Elastic Load Balance (Enhanced)",
			ShortDescription: `Manages a V2 certificate resource within HuaweiCloud.`,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"load",
				"balance",
				"enhanced",
				"lb",
				"certificate",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create an LB certificate. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new LB certificate.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Human-readable name for the Certificate. Does not have to be unique.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-readable description for the Certificate.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional) The domain of the Certificate.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Required) The private encrypted key of the Certificate, PEM format.`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(Required) The public encrypted key of the Certificate, PEM format. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Indicates the update time.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Indicates the creation time.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Indicates the update time.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Indicates the creation time.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_lb_l7policy_v2",
			Category:         "Elastic Load Balance (Enhanced)",
			ShortDescription: `Manages a V2 L7 Policy resource within HuaweiCloud.`,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"load",
				"balance",
				"enhanced",
				"lb",
				"l7policy",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create an . If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new L7 Policy.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) Required for admins. The UUID of the tenant who owns the L7 Policy. Only administrative users can specify a tenant UUID other than their own. Changing this creates a new L7 Policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Human-readable name for the L7 Policy. Does not have to be unique.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-readable description for the L7 Policy.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) The L7 Policy action - can either be REDIRECT\_TO\_POOL, or REDIRECT\_TO\_LISTENER. Changing this creates a new L7 Policy.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Required) The Listener on which the L7 Policy will be associated with. Changing this creates a new L7 Policy.`,
				},
				resource.Attribute{
					Name:        "position",
					Description: `(Optional) The position of this policy on the listener. Positions start at 1. Changing this creates a new L7 Policy.`,
				},
				resource.Attribute{
					Name:        "redirect_pool_id",
					Description: `(Optional) Requests matching this policy will be redirected to the pool with this ID. Only valid if action is REDIRECT\_TO\_POOL.`,
				},
				resource.Attribute{
					Name:        "redirect_listener_id",
					Description: `(Optional) Requests matching this policy will be redirected to the listener with this ID. Only valid if action is REDIRECT\_TO\_LISTENER.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) The administrative state of the L7 Policy. This value can only be true (UP). ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the L7 {olicy.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "position",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "redirect_pool_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "redirect_listener_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above. ## Import Load Balancer L7 Policy can be imported using the L7 Policy ID, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_lb_l7policy_v2.l7policy_1 8a7a79c2-cf17-4e65-b2ae-ddc8bfcf6c74 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the L7 {olicy.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "position",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "redirect_pool_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "redirect_listener_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above. ## Import Load Balancer L7 Policy can be imported using the L7 Policy ID, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_lb_l7policy_v2.l7policy_1 8a7a79c2-cf17-4e65-b2ae-ddc8bfcf6c74 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_lb_l7rule_v2",
			Category:         "Elastic Load Balance (Enhanced)",
			ShortDescription: `Manages a V2 l7rule resource within HuaweiCloud.`,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"load",
				"balance",
				"enhanced",
				"lb",
				"l7rule",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create an . If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new L7 Rule.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) Required for admins. The UUID of the tenant who owns the L7 Rule. Only administrative users can specify a tenant UUID other than their own. Changing this creates a new L7 Rule.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-readable description for the L7 Rule.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The L7 Rule type - can either be HOST\_NAME or PATH. Changing this creates a new L7 Rule.`,
				},
				resource.Attribute{
					Name:        "compare_type",
					Description: `(Required) The comparison type for the L7 rule - can either be STARTS\_WITH, EQUAL_TO or REGEX`,
				},
				resource.Attribute{
					Name:        "l7policy_id",
					Description: `(Required) The ID of the L7 Policy to query. Changing this creates a new L7 Rule.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value to use for the comparison. For example, the file type to compare.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) The key to use for the comparison. For example, the name of the cookie to evaluate. Valid when ` + "`" + `type` + "`" + ` is set to COOKIE or HEADER. Changing this creates a new L7 Rule.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) The administrative state of the L7 Rule. The value can only be true (UP). ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the L7 Rule.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "compare_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "l7policy_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "invert",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `The ID of the Listener owning this resource. ## Import Load Balancer L7 Rule can be imported using the L7 Policy ID and L7 Rule ID separated by a slash, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_lb_l7rule_v2.l7rule_1 e0bd694a-abbe-450e-b329-0931fd1cc5eb/4086b0c9-b18c-4d1c-b6b8-4c56c3ad2a9e ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the L7 Rule.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "compare_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "l7policy_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "invert",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `The ID of the Listener owning this resource. ## Import Load Balancer L7 Rule can be imported using the L7 Policy ID and L7 Rule ID separated by a slash, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_lb_l7rule_v2.l7rule_1 e0bd694a-abbe-450e-b329-0931fd1cc5eb/4086b0c9-b18c-4d1c-b6b8-4c56c3ad2a9e ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_lb_listener_v2",
			Category:         "Elastic Load Balance (Enhanced)",
			ShortDescription: `Manages a V2 listener resource within HuaweiCloud.`,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"load",
				"balance",
				"enhanced",
				"lb",
				"listener",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create an . If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new Listener.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The protocol can either be TCP, HTTP, HTTPS or TERMINATED_HTTPS. Changing this creates a new Listener.`,
				},
				resource.Attribute{
					Name:        "protocol_port",
					Description: `(Required) The port on which to listen for client traffic. Changing this creates a new Listener.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) Required for admins. The UUID of the tenant who owns the Listener. Only administrative users can specify a tenant UUID other than their own. Changing this creates a new Listener.`,
				},
				resource.Attribute{
					Name:        "loadbalancer_id",
					Description: `(Required) The load balancer on which to provision this Listener. Changing this creates a new Listener.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Human-readable name for the Listener. Does not have to be unique.`,
				},
				resource.Attribute{
					Name:        "default_pool_id",
					Description: `(Optional) The ID of the default pool with which the Listener is associated. Changing this creates a new Listener.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-readable description for the Listener.`,
				},
				resource.Attribute{
					Name:        "connection_limit",
					Description: `(Optional) The maximum number of connections allowed for the Listener. The value ranges from -1 to 2,147,483,647. This parameter is reserved and has been not used. Only the administrator can specify the maximum number of connections.`,
				},
				resource.Attribute{
					Name:        "http2_enable",
					Description: `(Optional) Specifies whether to use HTTP/2. The default value is false. This parameter is valid only when the protocol is set to ` + "`" + `TERMINATED_HTTPS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "default_tls_container_ref",
					Description: `(Optional) Specifies the ID of the server certificate used by the listener. This parameter is mandatory when protocol is set to ` + "`" + `TERMINATED_HTTPS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sni_container_refs",
					Description: `(Optional) Lists the IDs of SNI certificates (server certificates with a domain name) used by the listener. This parameter is valid when protocol is set to ` + "`" + `TERMINATED_HTTPS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) The administrative state of the Listener. A valid value is true (UP) or false (DOWN). ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the Listener.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "protocol_port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "default_port_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "connection_limit",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "http2_enable",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "default_tls_container_ref",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "sni_container_refs",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the Listener.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "protocol_port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "default_port_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "connection_limit",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "http2_enable",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "default_tls_container_ref",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "sni_container_refs",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_lb_loadbalancer_v2",
			Category:         "Elastic Load Balance (Enhanced)",
			ShortDescription: `Manages a V2 loadbalancer resource within HuaweiCloud.`,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"load",
				"balance",
				"enhanced",
				"lb",
				"loadbalancer",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create an LB member. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new LB member.`,
				},
				resource.Attribute{
					Name:        "vip_subnet_id",
					Description: `(Required) The network on which to allocate the Loadbalancer's address. A tenant can only create Loadbalancers on networks authorized by policy (e.g. networks that belong to them or networks that are shared). Changing this creates a new loadbalancer.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Human-readable name for the Loadbalancer. Does not have to be unique.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-readable description for the Loadbalancer.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) Required for admins. The UUID of the tenant who owns the Loadbalancer. Only administrative users can specify a tenant UUID other than their own. Changing this creates a new loadbalancer.`,
				},
				resource.Attribute{
					Name:        "vip_address",
					Description: `(Optional) The ip address of the load balancer. Changing this creates a new loadbalancer.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) The administrative state of the Loadbalancer. A valid value is true (UP) or false (DOWN).`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `(Optional) The UUID of a flavor. Changing this creates a new loadbalancer.`,
				},
				resource.Attribute{
					Name:        "loadbalancer_provider",
					Description: `(Optional) The name of the provider. Changing this creates a new loadbalancer.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Optional) A list of security group IDs to apply to the loadbalancer. The security groups must be specified by ID and not name (as opposed to how they are configured with the Compute Instance). ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vip_subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vip_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "loadbalancer_provider",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vip_port_id",
					Description: `The Port ID of the Load Balancer IP.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vip_subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vip_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "loadbalancer_provider",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vip_port_id",
					Description: `The Port ID of the Load Balancer IP.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_lb_member_v2",
			Category:         "Elastic Load Balance (Enhanced)",
			ShortDescription: `Manages a V2 member resource within HuaweiCloud.`,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"load",
				"balance",
				"enhanced",
				"lb",
				"member",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create an . If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new member.`,
				},
				resource.Attribute{
					Name:        "pool_id",
					Description: `(Required) The id of the pool that this member will be assigned to.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) The subnet in which to access the member`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Human-readable name for the member.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) Required for admins. The UUID of the tenant who owns the member. Only administrative users can specify a tenant UUID other than their own. Changing this creates a new member.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) The IP address of the member to receive traffic from the load balancer. Changing this creates a new member.`,
				},
				resource.Attribute{
					Name:        "protocol_port",
					Description: `(Required) The port on which to listen for client traffic. Changing this creates a new member.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) A positive integer value that indicates the relative portion of traffic that this member should receive from the pool. For example, a member with a weight of 10 receives five times as much traffic as a member with a weight of 2.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) The administrative state of the member. A valid value is true (UP) or false (DOWN). ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the member.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "pool_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "protocol_port",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the member.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "pool_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "protocol_port",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_lb_monitor_v2",
			Category:         "Elastic Load Balance (Enhanced)",
			ShortDescription: `Manages a V2 monitor resource within HuaweiCloud.`,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"load",
				"balance",
				"enhanced",
				"lb",
				"monitor",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create an . If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new monitor.`,
				},
				resource.Attribute{
					Name:        "pool_id",
					Description: `(Required) The id of the pool that this monitor will be assigned to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The Name of the Monitor.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) Required for admins. The UUID of the tenant who owns the monitor. Only administrative users can specify a tenant UUID other than their own. Changing this creates a new monitor.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of probe, which is PING, TCP, HTTP, or HTTPS, that is sent by the load balancer to verify the member state. Changing this creates a new monitor.`,
				},
				resource.Attribute{
					Name:        "delay",
					Description: `(Required) The time, in seconds, between sending probes to members.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Required) Maximum number of seconds for a monitor to wait for a ping reply before it times out. The value must be less than the delay value.`,
				},
				resource.Attribute{
					Name:        "max_retries",
					Description: `(Required) Number of permissible ping failures before changing the member's status to INACTIVE. Must be a number between 1 and 10.`,
				},
				resource.Attribute{
					Name:        "url_path",
					Description: `(Optional) Required for HTTP(S) types. URI path that will be accessed if monitor type is HTTP or HTTPS.`,
				},
				resource.Attribute{
					Name:        "expected_codes",
					Description: `(Optional) Required for HTTP(S) types. Expected HTTP codes for a passing HTTP(S) monitor. You can either specify a single status like "200", or a range like "200-202".`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) The administrative state of the monitor. A valid value is true (UP) or false (DOWN). ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the monitor.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "delay",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "max_retries",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "url_path",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "http_method",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "expected_codes",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the monitor.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "delay",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "max_retries",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "url_path",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "http_method",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "expected_codes",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_lb_pool_v2",
			Category:         "Elastic Load Balance (Enhanced)",
			ShortDescription: `Manages a V2 pool resource within HuaweiCloud.`,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"load",
				"balance",
				"enhanced",
				"lb",
				"pool",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create an . If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new pool.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) Required for admins. The UUID of the tenant who owns the pool. Only administrative users can specify a tenant UUID other than their own. Changing this creates a new pool.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Human-readable name for the pool.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-readable description for the pool.`,
				},
				resource.Attribute{
					Name:        "loadbalancer_id",
					Description: `(Optional) The load balancer on which to provision this pool. Changing this creates a new pool. Note: One of LoadbalancerID or ListenerID must be provided.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Optional) The Listener on which the members of the pool will be associated with. Changing this creates a new pool. Note: One of LoadbalancerID or ListenerID must be provided.`,
				},
				resource.Attribute{
					Name:        "lb_method",
					Description: `(Required) The load balancing algorithm to distribute traffic to the pool's members. Must be one of ROUND_ROBIN, LEAST_CONNECTIONS, or SOURCE_IP.`,
				},
				resource.Attribute{
					Name:        "persistence",
					Description: `Omit this field to prevent session persistence. Indicates whether connections in the same session will be processed by the same Pool member or not. Changing this creates a new pool.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) The administrative state of the pool. A valid value is true (UP) or false (DOWN). The ` + "`" + `persistence` + "`" + ` argument supports:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of persistence mode. The current specification supports SOURCE_IP, HTTP_COOKIE, and APP_COOKIE.`,
				},
				resource.Attribute{
					Name:        "cookie_name",
					Description: `(Optional) The name of the cookie if persistence mode is set appropriately. Required if ` + "`" + `type = APP_COOKIE` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the pool.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "lb_method",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "persistence",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the pool.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "lb_method",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "persistence",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_lb_whitelist_v2",
			Category:         "Elastic Load Balance (Enhanced)",
			ShortDescription: `Manages a Load Balancer whitelist resource within HuaweiCloud.`,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"load",
				"balance",
				"enhanced",
				"lb",
				"whitelist",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) Required for admins. The UUID of the tenant who owns the whitelist. Only administrative users can specify a tenant UUID other than their own. Changing this creates a new whitelist.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Required) The Listener ID that the whitelist will be associated with. Changing this creates a new whitelist.`,
				},
				resource.Attribute{
					Name:        "enable_whitelist",
					Description: `(Optional) Specify whether to enable access control.`,
				},
				resource.Attribute{
					Name:        "whitelist",
					Description: `(Optional) Specifies the IP addresses in the whitelist. Use commas(,) to separate the multiple IP addresses. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the whitelist.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "enable_whitelist",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "whitelist",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the whitelist.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "enable_whitelist",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "whitelist",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_maas_task_v1",
			Category:         "Object Storage Migration Service",
			ShortDescription: `Manages resource task within HuaweiCloud MAAS.`,
			Description:      ``,
			Keywords: []string{
				"object",
				"storage",
				"migration",
				"service",
				"maas",
				"task",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "src_node",
					Description: `(Required) Specifies the source node information.`,
				},
				resource.Attribute{
					Name:        "dst_node",
					Description: `(Required) Specifies the destination node information.`,
				},
				resource.Attribute{
					Name:        "enable_kms",
					Description: `(Required) Specifies whether to use KMS encryption.`,
				},
				resource.Attribute{
					Name:        "thread_num",
					Description: `(Required) Specifies the number of threads used by the migration task. The value cannot exceed 50.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Specifies tasks description, which cannot exceed 255 characters. The following special characters are not allowed: <>()"&`,
				},
				resource.Attribute{
					Name:        "smn_info",
					Description: `(Optional) Specifies the field used for sending messages using the Simple Message Notification (SMN) service. The ` + "`" + `src_node` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Specifies the region where the source bucket locates.`,
				},
				resource.Attribute{
					Name:        "ak",
					Description: `(Required) Specifies the source bucket Access Key.`,
				},
				resource.Attribute{
					Name:        "sk",
					Description: `(Required) Specifies the source bucket Secret Key.`,
				},
				resource.Attribute{
					Name:        "object_key",
					Description: `(Required) Specifies the name of the object to be selected in the source bucket.`,
				},
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) Specifies the name of the source bucket.`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `(Optional) Specifies the source cloud vendor. Currently only Aliyun and AWS are supported. The default value is Aliyun. The ` + "`" + `dst_node` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Specifies the region where the destination bucket locates.`,
				},
				resource.Attribute{
					Name:        "ak",
					Description: `(Required) Specifies the destination bucket Access Key.`,
				},
				resource.Attribute{
					Name:        "sk",
					Description: `(Required) Specifies the destination bucket Secret Key.`,
				},
				resource.Attribute{
					Name:        "object_key",
					Description: `(Required) Specifies the name of the object to be selected in the destination bucket.`,
				},
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) Specifies the name of the destination bucket. The ` + "`" + `smn_info` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "topic_urn",
					Description: `(Required) Specifies the SMN message topic URN bound to a migration task.`,
				},
				resource.Attribute{
					Name:        "language",
					Description: `(Optional) Specifies the management console language used by the current users. Users can select en-us.`,
				},
				resource.Attribute{
					Name:        "trigger_conditions",
					Description: `(Required) Specifies the trigger conditions of sending messages using SMN. The value depending on the state of a migration task. The migration task status can be SUCCESS or FAIL. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "src_node",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dst_node",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "enable_kms",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "thread_num",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "smn_info",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name for a task.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Specifies the task status as follows: 0: Not started, 1: Waiting to migrate, 2: Migrating, 3: Migration paused, 4: Migration failed, 5: Migration succeeded.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "src_node",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dst_node",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "enable_kms",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "thread_num",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "smn_info",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name for a task.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Specifies the task status as follows: 0: Not started, 1: Waiting to migrate, 2: Migrating, 3: Migration paused, 4: Migration failed, 5: Migration succeeded.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_mls_instance",
			Category:         "Machine Learning Service (MLS)",
			ShortDescription: `mls instance`,
			Description:      ``,
			Keywords: []string{
				"machine",
				"learning",
				"service",
				"mls",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "flavor",
					Description: `(Required) Instance flavor`,
				},
				resource.Attribute{
					Name:        "mrs_cluster",
					Description: `(Required) A nested object resource Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Instance name. A tenant has a unique name of the instance of one type. Value range: An instance name must contain 4 to 64 characters and must start with a letter. The name is case insensitive and contains only letters, digits, and hyphens (-) or underscores (_), excluding other special characters.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Required) A nested object resource Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) Instance version The ` + "`" + `mrs_cluster` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) MRS cluster ID`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `(Optional) MRS cluster username. This parameter is mandatory only when the MRS cluster is in the security mode`,
				},
				resource.Attribute{
					Name:        "user_password",
					Description: `(Optional) Password of the MRS cluster user The ` + "`" + `network` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "available_zone",
					Description: `(Required) az`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required) ID of the subnet where the instance resides`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `(Required) A nested object resource Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Optional) ID of the security group of the instance`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) ID of the virtual private cloud (VPC) where the instance resides The ` + "`" + `public_ip` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "bind_type",
					Description: `(Required) Bind type. Possible values: auto_assign, not_use`,
				},
				resource.Attribute{
					Name:        "eip_id",
					Description: `EIP ID. This parameter value is returned only when bindType is set to auto_assign - - -`,
				},
				resource.Attribute{
					Name:        "agency",
					Description: `(Optional) Agency name. This parameter is mandatory only when you bind an instance to an elastic IP address (EIP). An instance must be bound to an EIP to grant MLS rights to obtain a tenant's token. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `Time when the instance is created. The parameter format is yyyy-mm-dd Thh:mm:ssZ. In the format, T indicates a time start point and Z specifies a UTC offset, for example, the Beijing time offset is +0800`,
				},
				resource.Attribute{
					Name:        "current_task",
					Description: `Instance task status. Possible values: UNFREEZING FREEZING RESTORING SNAPSHOTTING GROWING REBOOTING REBOOT_FAILURE RESIZE_FAILURE`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `instance id`,
				},
				resource.Attribute{
					Name:        "inner_endpoint",
					Description: `URL for accessing the instance. Only machines in the same VPC and subnet as the instance can access the URL`,
				},
				resource.Attribute{
					Name:        "public_endpoint",
					Description: `URL for accessing the instance. The URL can be accessed from the Internet. The URL is created only after the instance is bound to an EIP.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Instance status. Possible values: CREATING AVAILABLE FAILED CREATION FAILED`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `Time when the instance is updated. The parameter format is the same as the format of the created parameter ## Timeouts This resource provides the following timeouts configuration options: - ` + "`" + `create` + "`" + ` - Default is 10 minute. - ` + "`" + `delete` + "`" + ` - Default is 10 minute. ## Import Instance can be imported using the following format: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_mls_instance.default {{ resource id}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created",
					Description: `Time when the instance is created. The parameter format is yyyy-mm-dd Thh:mm:ssZ. In the format, T indicates a time start point and Z specifies a UTC offset, for example, the Beijing time offset is +0800`,
				},
				resource.Attribute{
					Name:        "current_task",
					Description: `Instance task status. Possible values: UNFREEZING FREEZING RESTORING SNAPSHOTTING GROWING REBOOTING REBOOT_FAILURE RESIZE_FAILURE`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `instance id`,
				},
				resource.Attribute{
					Name:        "inner_endpoint",
					Description: `URL for accessing the instance. Only machines in the same VPC and subnet as the instance can access the URL`,
				},
				resource.Attribute{
					Name:        "public_endpoint",
					Description: `URL for accessing the instance. The URL can be accessed from the Internet. The URL is created only after the instance is bound to an EIP.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Instance status. Possible values: CREATING AVAILABLE FAILED CREATION FAILED`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `Time when the instance is updated. The parameter format is the same as the format of the created parameter ## Timeouts This resource provides the following timeouts configuration options: - ` + "`" + `create` + "`" + ` - Default is 10 minute. - ` + "`" + `delete` + "`" + ` - Default is 10 minute. ## Import Instance can be imported using the following format: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_mls_instance.default {{ resource id}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_mrs_cluster_v1",
			Category:         "MapReduce Service (MRS)",
			ShortDescription: `Manages resource cluster within HuaweiCloud MRS.`,
			Description:      ``,
			Keywords: []string{
				"mapreduce",
				"service",
				"mrs",
				"cluster",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "billing_type",
					Description: `(Required) The value is 12, indicating on-demand payment.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Cluster region information. Obtain the value from Regions and Endpoints.`,
				},
				resource.Attribute{
					Name:        "master_node_num",
					Description: `(Required) Number of Master nodes The value is 2.`,
				},
				resource.Attribute{
					Name:        "master_node_size",
					Description: `(Required) Best match based on several years of commissioning experience. MRS supports specifications of hosts, and host specifications are determined by CPUs, memory, and disks space. MRS supports instance specifications detailed in [MRS specifications](https://support.huaweicloud.com/en-us/api-mrs/mrs_01_9006.html)`,
				},
				resource.Attribute{
					Name:        "core_node_num",
					Description: `(Required) Number of Core nodes Value range: 3 to 100 A maximum of 100 Core nodes are supported by default. If more than 100 Core nodes are required, contact technical support engineers or invoke background APIs to modify the database.`,
				},
				resource.Attribute{
					Name:        "core_node_size",
					Description: `(Required) Instance specification of a Core node Configuration method of this parameter is identical to that of master_node_size.`,
				},
				resource.Attribute{
					Name:        "available_zone_id",
					Description: `(Required) ID of an available zone. Obtain the value from Regions and Endpoints. North China AZ1 (cn-north-1a): ae04cf9d61544df3806a3feeb401b204, North China AZ2 (cn-north-1b): d573142f24894ef3bd3664de068b44b0, East China AZ1 (cn-east-2a): 72d50cedc49846b9b42c21495f38d81c, East China AZ2 (cn-east-2b): 38b0f7a602344246bcb0da47b5d548e7, East China AZ3 (cn-east-2c): 5547fd6bf8f84bb5a7f9db062ad3d015, South China AZ1(cn-south-1a): 34f5ff4865cf4ed6b270f15382ebdec5, South China AZ2(cn-south-2b): 043c7e39ecb347a08dc8fcb6c35a274e, South China AZ3(cn-south-1c): af1687643e8c4ec1b34b688e4e3b8901,`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Required) Cluster name, which is globally unique and contains only 1 to 64 letters, digits, hyphens (-), and underscores (_).`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) ID of the VPC where the subnet locates Obtain the VPC ID from the management console as follows: Register an account and log in to the management console. Click Virtual Private Cloud and select Virtual Private Cloud from the left list. On the Virtual Private Cloud page, obtain the VPC ID from the list.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) Subnet ID Obtain the subnet ID from the management console as follows: Register an account and log in to the management console. Click Virtual Private Cloud and select Virtual Private Cloud from the left list. On the Virtual Private Cloud page, obtain the subnet ID from the list.`,
				},
				resource.Attribute{
					Name:        "cluster_version",
					Description: `(Optional) Version of the clusters Currently, MRS 1.6.3, MRS 1.7.2 and MRS 1.8.1 are supported. The latest version of MRS is used by default. Currently, the latest version is MRS 1.8.1.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `(Optional) Type of clusters 0: analysis cluster 1: streaming cluster The default value is 0.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `(Required) Type of disks SATA and SSD are supported. SATA: common I/O SSD: super high-speed I/O`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `(Required) Data disk storage space of a Core node Users can add disks to expand storage capacity when creating a cluster. There are the following scenarios: Separation of data storage and computing: Data is stored in the OBS system. Costs of clusters are relatively low but computing performance is poor. The clusters can be deleted at any time. It is recommended when data computing is not frequently performed. Integration of data storage and computing: Data is stored in the HDFS system. Costs of clusters are relatively high but computing performance is good. The clusters cannot be deleted in a short term. It is recommended when data computing is frequently performed. Value range: 100 GB to 32000 GB`,
				},
				resource.Attribute{
					Name:        "node_public_cert_name",
					Description: `(Required) Name of a key pair You can use a key to log in to the Master node in the cluster.`,
				},
				resource.Attribute{
					Name:        "safe_mode",
					Description: `(Required) MRS cluster running mode 0: common mode The value indicates that the Kerberos authentication is disabled. Users can use all functions provided by the cluster. 1: safe mode The value indicates that the Kerberos authentication is enabled. Common users cannot use the file management or job management functions of an MRS cluster and cannot view cluster resource usage or the job records of Hadoop and Spark. To use these functions, the users must obtain the relevant permissions from the MRS Manager administrator. The request has the cluster_admin_secret parameter only when safe_mode is set to 1.`,
				},
				resource.Attribute{
					Name:        "cluster_admin_secret",
					Description: `(Optional) Indicates the password of the MRS Manager administrator. The password for MRS 1.5.0: Must contain 6 to 32 characters. Must contain at least two types of the following: Lowercase letters Uppercase letters Digits Special characters of ` + "`" + `~!@#$%^&`,
				},
				resource.Attribute{
					Name:        "log_collection",
					Description: `(Optional) Indicates whether logs are collected when cluster installation fails. 0: not collected 1: collected The default value is 0. If log_collection is set to 1, OBS buckets will be created to collect the MRS logs. These buckets will be charged.`,
				},
				resource.Attribute{
					Name:        "component_list",
					Description: `(Required) Service component list.`,
				},
				resource.Attribute{
					Name:        "add_jobs",
					Description: `(Optional) You can submit a job when you create a cluster to save time and use MRS easily. Only one job can be added. The ` + "`" + `component_list` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "component_name",
					Description: `(Required) Component name Currently, Hadoop, Spark, HBase, Hive, Hue, Loader, Flume, Kafka and Storm are supported. The ` + "`" + `add_jobs` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "job_type",
					Description: `(Required) Job type 1: MapReduce 2: Spark 3: Hive Script 4: HiveQL (not supported currently) 5: DistCp, importing and exporting data (not supported in this API currently). 6: Spark Script 7: Spark SQL, submitting Spark SQL statements (not supported in this API currently). NOTE: Spark and Hive jobs can be added to only clusters including Spark and Hive components.`,
				},
				resource.Attribute{
					Name:        "job_name",
					Description: `(Required) Job name It contains only 1 to 64 letters, digits, hyphens (-), and underscores (_). NOTE: Identical job names are allowed but not recommended.`,
				},
				resource.Attribute{
					Name:        "jar_path",
					Description: `(Required) Path of the .jar file or .sql file for program execution The parameter must meet the following requirements: Contains a maximum of 1023 characters, excluding special characters such as ;|&><'$. The address cannot be empty or full of spaces. Starts with / or s3a://. Spark Script must end with .sql; while MapReduce and Spark Jar must end with .jar. sql and jar are case-insensitive.`,
				},
				resource.Attribute{
					Name:        "arguments",
					Description: `(Optional) Key parameter for program execution The parameter is specified by the function of the user's program. MRS is only responsible for loading the parameter. The parameter contains a maximum of 2047 characters, excluding special characters such as ;|&>'<$, and can be empty.`,
				},
				resource.Attribute{
					Name:        "input",
					Description: `(Optional) Path for inputting data, which must start with / or s3a://. A correct OBS path is required. The parameter contains a maximum of 1023 characters, excluding special characters such as ;|&>'<$, and can be empty.`,
				},
				resource.Attribute{
					Name:        "output",
					Description: `(Optional) Path for outputting data, which must start with / or s3a://. A correct OBS path is required. If the path does not exist, the system automatically creates it. The parameter contains a maximum of 1023 characters, excluding special characters such as ;|&>'<$, and can be empty.`,
				},
				resource.Attribute{
					Name:        "job_log",
					Description: `(Optional) Path for storing job logs that record job running status. This path must start with / or s3a://. A correct OBS path is required. The parameter contains a maximum of 1023 characters, excluding special characters such as ;|&>'<$, and can be empty.`,
				},
				resource.Attribute{
					Name:        "shutdown_cluster",
					Description: `(Optional) Whether to delete the cluster after the jobs are complete true: Yes false: No`,
				},
				resource.Attribute{
					Name:        "file_action",
					Description: `(Optional) Data import and export import export`,
				},
				resource.Attribute{
					Name:        "submit_job_once_cluster_run",
					Description: `(Required) true: A job is submitted when a cluster is created. false: A job is submitted separately. The parameter is set to true in this example.`,
				},
				resource.Attribute{
					Name:        "hql",
					Description: `(Optional) HiveQL statement`,
				},
				resource.Attribute{
					Name:        "hive_script_path",
					Description: `(Optional) SQL program path This parameter is needed by Spark Script and Hive Script jobs only and must meet the following requirements: Contains a maximum of 1023 characters, excluding special characters such as ;|&><'$. The address cannot be empty or full of spaces. Starts with / or s3a://. Ends with .sql. sql is case-insensitive. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "billing_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "data_center",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "master_node_num",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "master_node_size",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "core_node_num",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "core_node_size",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "available_zone_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cluster_version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "node_public_cert_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "safe_mode",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cluster_admin_secret",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "log_collection",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "component_list",
					Description: `See Argument Reference below.`,
				},
				resource.Attribute{
					Name:        "add_jobs",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "order_id",
					Description: `Order ID for creating clusters.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `Cluster ID.`,
				},
				resource.Attribute{
					Name:        "available_zone_name",
					Description: `Name of an availability zone.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `Instance ID.`,
				},
				resource.Attribute{
					Name:        "hadoop_version",
					Description: `Hadoop version.`,
				},
				resource.Attribute{
					Name:        "master_node_ip",
					Description: `IP address of a Master node.`,
				},
				resource.Attribute{
					Name:        "externalIp",
					Description: `Internal IP address.`,
				},
				resource.Attribute{
					Name:        "private_ip_first",
					Description: `Primary private IP address.`,
				},
				resource.Attribute{
					Name:        "external_ip",
					Description: `External IP address.`,
				},
				resource.Attribute{
					Name:        "slave_security_groups_id",
					Description: `Standby security group ID.`,
				},
				resource.Attribute{
					Name:        "security_groups_id",
					Description: `Security group ID.`,
				},
				resource.Attribute{
					Name:        "external_alternate_ip",
					Description: `Backup external IP address.`,
				},
				resource.Attribute{
					Name:        "master_node_spec_id",
					Description: `Specification ID of a Master node.`,
				},
				resource.Attribute{
					Name:        "core_node_spec_id",
					Description: `Specification ID of a Core node.`,
				},
				resource.Attribute{
					Name:        "master_node_product_id",
					Description: `Product ID of a Master node.`,
				},
				resource.Attribute{
					Name:        "core_node_product_id",
					Description: `Product ID of a Core node.`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `Cluster subscription duration.`,
				},
				resource.Attribute{
					Name:        "vnc",
					Description: `URI address for remote login of the elastic cloud server.`,
				},
				resource.Attribute{
					Name:        "fee",
					Description: `Cluster creation fee, which is automatically calculated.`,
				},
				resource.Attribute{
					Name:        "deployment_id",
					Description: `Deployment ID of a cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_state",
					Description: `Cluster status Valid values include: existing history starting running terminated failed abnormal terminating rebooting shutdown frozen scaling-out scaling-in scaling-error.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `Project ID.`,
				},
				resource.Attribute{
					Name:        "create_at",
					Description: `Cluster creation time.`,
				},
				resource.Attribute{
					Name:        "update_at",
					Description: `Cluster update time.`,
				},
				resource.Attribute{
					Name:        "error_info",
					Description: `Error information.`,
				},
				resource.Attribute{
					Name:        "charging_start_time",
					Description: `Time when charging starts.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `Remarks of a cluster. The component_list attributes:`,
				},
				resource.Attribute{
					Name:        "component_name",
					Description: `(Required) Component name Currently, Hadoop, Spark, HBase, Hive, Hue, Loader, Flume, Kafka and Storm are supported.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "billing_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "data_center",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "master_node_num",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "master_node_size",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "core_node_num",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "core_node_size",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "available_zone_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cluster_version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "node_public_cert_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "safe_mode",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cluster_admin_secret",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "log_collection",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "component_list",
					Description: `See Argument Reference below.`,
				},
				resource.Attribute{
					Name:        "add_jobs",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "order_id",
					Description: `Order ID for creating clusters.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `Cluster ID.`,
				},
				resource.Attribute{
					Name:        "available_zone_name",
					Description: `Name of an availability zone.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `Instance ID.`,
				},
				resource.Attribute{
					Name:        "hadoop_version",
					Description: `Hadoop version.`,
				},
				resource.Attribute{
					Name:        "master_node_ip",
					Description: `IP address of a Master node.`,
				},
				resource.Attribute{
					Name:        "externalIp",
					Description: `Internal IP address.`,
				},
				resource.Attribute{
					Name:        "private_ip_first",
					Description: `Primary private IP address.`,
				},
				resource.Attribute{
					Name:        "external_ip",
					Description: `External IP address.`,
				},
				resource.Attribute{
					Name:        "slave_security_groups_id",
					Description: `Standby security group ID.`,
				},
				resource.Attribute{
					Name:        "security_groups_id",
					Description: `Security group ID.`,
				},
				resource.Attribute{
					Name:        "external_alternate_ip",
					Description: `Backup external IP address.`,
				},
				resource.Attribute{
					Name:        "master_node_spec_id",
					Description: `Specification ID of a Master node.`,
				},
				resource.Attribute{
					Name:        "core_node_spec_id",
					Description: `Specification ID of a Core node.`,
				},
				resource.Attribute{
					Name:        "master_node_product_id",
					Description: `Product ID of a Master node.`,
				},
				resource.Attribute{
					Name:        "core_node_product_id",
					Description: `Product ID of a Core node.`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `Cluster subscription duration.`,
				},
				resource.Attribute{
					Name:        "vnc",
					Description: `URI address for remote login of the elastic cloud server.`,
				},
				resource.Attribute{
					Name:        "fee",
					Description: `Cluster creation fee, which is automatically calculated.`,
				},
				resource.Attribute{
					Name:        "deployment_id",
					Description: `Deployment ID of a cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_state",
					Description: `Cluster status Valid values include: existing history starting running terminated failed abnormal terminating rebooting shutdown frozen scaling-out scaling-in scaling-error.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `Project ID.`,
				},
				resource.Attribute{
					Name:        "create_at",
					Description: `Cluster creation time.`,
				},
				resource.Attribute{
					Name:        "update_at",
					Description: `Cluster update time.`,
				},
				resource.Attribute{
					Name:        "error_info",
					Description: `Error information.`,
				},
				resource.Attribute{
					Name:        "charging_start_time",
					Description: `Time when charging starts.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `Remarks of a cluster. The component_list attributes:`,
				},
				resource.Attribute{
					Name:        "component_name",
					Description: `(Required) Component name Currently, Hadoop, Spark, HBase, Hive, Hue, Loader, Flume, Kafka and Storm are supported.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_mrs_job_v1",
			Category:         "MapReduce Service (MRS)",
			ShortDescription: `Manages resource job within HuaweiCloud MRS.`,
			Description:      ``,
			Keywords: []string{
				"mapreduce",
				"service",
				"mrs",
				"job",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "job_type",
					Description: `(Required) Job type 1: MapReduce 2: Spark 3: Hive Script 4: HiveQL (not supported currently) 5: DistCp, importing and exporting data. 6: Spark Script 7: Spark SQL, submitting Spark SQL statements. (not supported in this APIcurrently) NOTE: Spark and Hive jobs can be added to only clusters including Spark and Hive components.`,
				},
				resource.Attribute{
					Name:        "job_name",
					Description: `(Required) Job name Contains only 1 to 64 letters, digits, hyphens (-), and underscores (_). NOTE: Identical job names are allowed but not recommended.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) Cluster ID`,
				},
				resource.Attribute{
					Name:        "jar_path",
					Description: `(Required) Path of the .jar package or .sql file for program execution The parameter must meet the following requirements: Contains a maximum of 1023 characters, excluding special characters such as ;|&><'$. The address cannot be empty or full of spaces. Starts with / or s3a://. Spark Script must end with .sql; while MapReduce and Spark Jar must end with .jar. sql and jar are case-insensitive.`,
				},
				resource.Attribute{
					Name:        "arguments",
					Description: `(Optional) Key parameter for program execution. The parameter is specified by the function of the user's program. MRS is only responsible for loading the parameter. The parameter contains a maximum of 2047 characters, excluding special characters such as ;|&>'<$, and can be empty.`,
				},
				resource.Attribute{
					Name:        "input",
					Description: `(Optional) Path for inputting data, which must start with / or s3a://. A correct OBS path is required. The parameter contains a maximum of 1023 characters, excluding special characters such as ;|&>'<$, and can be empty.`,
				},
				resource.Attribute{
					Name:        "output",
					Description: `(Optional) Path for outputting data, which must start with / or s3a://. A correct OBS path is required. If the path does not exist, the system automatically creates it. The parameter contains a maximum of 1023 characters, excluding special characters such as ;|&>'<$, and can be empty.`,
				},
				resource.Attribute{
					Name:        "job_log",
					Description: `(Optional) Path for storing job logs that record job running status. This path must start with / or s3a://. A correct OBS path is required. The parameter contains a maximum of 1023 characters, excluding special characters such as ;|&>'<$, and can be empty.`,
				},
				resource.Attribute{
					Name:        "hive_script_path",
					Description: `(Optional) SQL program path This parameter is needed by Spark Script and Hive Script jobs only and must meet the following requirements: Contains a maximum of 1023 characters, excluding special characters such as ;|&><'$. The address cannot be empty or full of spaces. Starts with / or s3a://. Ends with .sql. sql is case-insensitive.`,
				},
				resource.Attribute{
					Name:        "is_protected",
					Description: `(Optional) Whether a job is protected true false The current version does not support this function.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `(Optional) Whether a job is public true false The current version does not support this function. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "job_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "job_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "jar_path",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "arguments",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "input",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "output",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "job_log",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "hive_script_path",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "is_protected",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "job_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "job_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "jar_path",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "arguments",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "input",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "output",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "job_log",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "hive_script_path",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "is_protected",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_nat_dnat_rule_v2",
			Category:         "NAT Gateway (NAT)",
			ShortDescription: `Manages a V2 dnat rule resource within HuaweiCloud Nat.`,
			Description:      ``,
			Keywords: []string{
				"nat",
				"gateway",
				"dnat",
				"rule",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "floating_ip_id",
					Description: `(Required) Specifies the ID of the floating IP address. Changing this creates a new resource.`,
				},
				resource.Attribute{
					Name:        "internal_service_port",
					Description: `(Required) Specifies port used by ECSs or BMSs to provide services for external systems. Changing this creates a new resource.`,
				},
				resource.Attribute{
					Name:        "nat_gateway_id",
					Description: `(Required) ID of the nat gateway this dnat rule belongs to. Changing this creates a new dnat rule.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `(Optional) Specifies the port ID of an ECS or a BMS. This parameter and private_ip are alternative. Changing this creates a new dnat rule.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `(Optional) Specifies the private IP address of a user, for example, the IP address of a VPC for dedicated connection. This parameter and port_id are alternative. Changing this creates a new dnat rule.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Specifies the protocol type. Currently, TCP, UDP, and ANY are supported. Changing this creates a new dnat rule.`,
				},
				resource.Attribute{
					Name:        "external_service_port",
					Description: `(Required) Specifies port used by ECSs or BMSs to provide services for external systems. Changing this creates a new dnat rule. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Dnat rule creation time.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Dnat rule status.`,
				},
				resource.Attribute{
					Name:        "floating_ip_address",
					Description: `The actual floating IP address. ## Import Dnat can be imported using the following format: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_nat_dnat_rule_v2.dnat_1 f4f783a7-b908-4215-b018-724960e5df4a ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `Dnat rule creation time.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Dnat rule status.`,
				},
				resource.Attribute{
					Name:        "floating_ip_address",
					Description: `The actual floating IP address. ## Import Dnat can be imported using the following format: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_nat_dnat_rule_v2.dnat_1 f4f783a7-b908-4215-b018-724960e5df4a ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_nat_gateway_v2",
			Category:         "NAT Gateway (NAT)",
			ShortDescription: `Manages a V2 nat gateway resource within HuaweiCloud Nat.`,
			Description:      ``,
			Keywords: []string{
				"nat",
				"gateway",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 nat client. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new nat gateway.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the nat gateway.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the nat gateway.`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `(Required) The specification of the nat gateway, valid values are "1", "2", "3", "4".`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The target tenant ID in which to allocate the nat gateway. Changing this creates a new nat gateway.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `(Required) ID of the router this nat gateway belongs to. Changing this creates a new nat gateway.`,
				},
				resource.Attribute{
					Name:        "internal_network_id",
					Description: `(Optional) ID of the network this nat gateway connects to. Changing this creates a new nat gateway. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "internal_network_id",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "internal_network_id",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_nat_snat_rule_v2",
			Category:         "NAT Gateway (NAT)",
			ShortDescription: `Manages a V2 snat rule resource within HuaweiCloud Nat.`,
			Description:      ``,
			Keywords: []string{
				"nat",
				"gateway",
				"snat",
				"rule",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 nat client. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new snat rule.`,
				},
				resource.Attribute{
					Name:        "nat_gateway_id",
					Description: `(Required) ID of the nat gateway this snat rule belongs to. Changing this creates a new snat rule.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required) ID of the network this snat rule connects to. Changing this creates a new snat rule.`,
				},
				resource.Attribute{
					Name:        "floating_ip_id",
					Description: `(Required) ID of the floating ip this snat rule connets to. Changing this creates a new snat rule. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "nat_gateway_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "floating_ip_id",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "nat_gateway_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "floating_ip_id",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_networking_floatingip_associate_v2",
			Category:         "Virtual Private Cloud (VPC)",
			ShortDescription: `Associates a Floating IP to a Port`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"networking",
				"floatingip",
				"associate",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create a floating IP that can be used with another networking resource, such as a load balancer. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new floating IP (which may or may not have a different address).`,
				},
				resource.Attribute{
					Name:        "floating_ip",
					Description: `(Required) IP Address of an existing floating IP.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `(Required) ID of an existing port with at least one IP address to associate with this floating IP. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "floating_ip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `See Argument Reference above. ## Import Floating IP associations can be imported using the ` + "`" + `id` + "`" + ` of the floating IP, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_networking_floatingip_associate_v2.fip 2c7f39f3-702b-48d1-940c-b50384177ee1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "floating_ip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `See Argument Reference above. ## Import Floating IP associations can be imported using the ` + "`" + `id` + "`" + ` of the floating IP, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_networking_floatingip_associate_v2.fip 2c7f39f3-702b-48d1-940c-b50384177ee1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_networking_floatingip_v2",
			Category:         "Virtual Private Cloud (VPC)",
			ShortDescription: `Manages a V2 floating IP resource within HuaweiCloud Neutron (networking).`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"networking",
				"floatingip",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create a floating IP that can be used with another networking resource, such as a load balancer. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new floating IP (which may or may not have a different address).`,
				},
				resource.Attribute{
					Name:        "pool",
					Description: `(Optional) The name of the pool from which to obtain the floating IP. Only admin_external_net is valid. Changing this creates a new floating IP.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `(Optional) ID of an existing port with at least one IP address to associate with this floating IP.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The target tenant ID in which to allocate the floating IP, if you specify this together with a port_id, make sure the target port belongs to the same tenant. Changing this creates a new floating IP (which may or may not have a different address)`,
				},
				resource.Attribute{
					Name:        "fixed_ip",
					Description: `Fixed IP of the port to associate with this floating IP. Required if the port has multiple fixed IPs.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional options. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "pool",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The actual floating IP address itself.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `ID of associated port.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `the ID of the tenant in which to create the floating IP.`,
				},
				resource.Attribute{
					Name:        "fixed_ip",
					Description: `The fixed IP which the floating IP maps to. ## Import Floating IPs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_networking_floatingip_v2.floatip_1 2c7f39f3-702b-48d1-940c-b50384177ee1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "pool",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The actual floating IP address itself.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `ID of associated port.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `the ID of the tenant in which to create the floating IP.`,
				},
				resource.Attribute{
					Name:        "fixed_ip",
					Description: `The fixed IP which the floating IP maps to. ## Import Floating IPs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_networking_floatingip_v2.floatip_1 2c7f39f3-702b-48d1-940c-b50384177ee1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_networking_network_v2",
			Category:         "Virtual Private Cloud (VPC)",
			ShortDescription: `Manages a V2 Neutron network resource within HuaweiCloud.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"networking",
				"network",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create a Neutron network. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new network.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the network. Changing this updates the name of the existing network.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `(Optional) Specifies whether the network resource can be accessed by any tenant or not. Changing this updates the sharing capabalities of the existing network.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the network. Required if admin wants to create a network for another tenant. Changing this creates a new network.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) The administrative state of the network. Acceptable values are "true" and "false". Changing this value updates the state of the existing network.`,
				},
				resource.Attribute{
					Name:        "segments",
					Description: `(Optional) An array of one or more provider segment objects.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional options. The ` + "`" + `segments` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "physical_network",
					Description: `The phisical network where this network is implemented.`,
				},
				resource.Attribute{
					Name:        "segmentation_id",
					Description: `An isolated segment on the physical network.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `The type of physical network. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above. ## Import Networks can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_networking_network_v2.network_1 d90ce693-5ccf-4136-a0ed-152ce412b6b9 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above. ## Import Networks can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_networking_network_v2.network_1 d90ce693-5ccf-4136-a0ed-152ce412b6b9 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_networking_port_v2",
			Category:         "Virtual Private Cloud (VPC)",
			ShortDescription: `Manages a V2 port resource within HuaweiCloud.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"networking",
				"port",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 networking client. A networking client is needed to create a port. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new port.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A unique name for the port. Changing this updates the ` + "`" + `name` + "`" + ` of an existing port.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required) The ID of the network to attach the port to. Changing this creates a new port.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) Administrative up/down status for the port (must be "true" or "false" if provided). Changing this updates the ` + "`" + `admin_state_up` + "`" + ` of an existing port.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `(Optional) Specify a specific MAC address for the port. Changing this creates a new port.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the Port. Required if admin wants to create a port for another tenant. Changing this creates a new port.`,
				},
				resource.Attribute{
					Name:        "device_owner",
					Description: `(Optional) The device owner of the Port. Changing this creates a new port.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Optional - Conflicts with ` + "`" + `no_security_groups` + "`" + `) A list of security group IDs to apply to the port. The security groups must be specified by ID and not name (as opposed to how they are configured with the Compute Instance).`,
				},
				resource.Attribute{
					Name:        "no_security_groups",
					Description: `(Optional - Conflicts with ` + "`" + `security_group_ids` + "`" + `) If set to ` + "`" + `true` + "`" + `, then no security groups are applied to the port. If set to ` + "`" + `false` + "`" + ` and no ` + "`" + `security_group_ids` + "`" + ` are specified, then the Port will yield to the default behavior of the Networking service, which is to usually apply the "default" security group.`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `(Optional) The ID of the device attached to the port. Changing this creates a new port.`,
				},
				resource.Attribute{
					Name:        "fixed_ip",
					Description: `(Optional) An array of desired IPs for this port. The structure is described below.`,
				},
				resource.Attribute{
					Name:        "allowed_address_pairs",
					Description: `(Optional) An IP/MAC Address pair of additional IP addresses that can be active on this port. The structure is described below.`,
				},
				resource.Attribute{
					Name:        "extra_dhcp_option",
					Description: `(Optional) An extra DHCP option that needs to be configured on the port. The structure is described below. Can be specified multiple times.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional options. The ` + "`" + `fixed_ip` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) Subnet in which to allocate IP address for this port.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) IP address desired in the subnet for this port. If you don't specify ` + "`" + `ip_address` + "`" + `, an available IP address from the specified subnet will be allocated to this port. This field will not be populated if it is left blank. To retrieve the assigned IP address, use the ` + "`" + `all_fixed_ips` + "`" + ` attribute. The ` + "`" + `allowed_address_pairs` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Required) The additional IP address.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `(Optional) The additional MAC address. The ` + "`" + `extra_dhcp_option` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the DHCP option.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Value of the DHCP option.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `(Optional) IP protocol version. Defaults to 4. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "device_owner",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "fixed_ip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "all_fixed_ips",
					Description: `The collection of Fixed IP addresses on the port in the order returned by the Network v2 API.`,
				},
				resource.Attribute{
					Name:        "all_security_group_ids",
					Description: `The collection of Security Group IDs on the port which have been explicitly and implicitly added.`,
				},
				resource.Attribute{
					Name:        "extra_dhcp_option",
					Description: `See Argument Reference above. ## Import Ports can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_networking_port_v2.port_1 eae26a3e-1c33-4cc1-9c31-0cd729c438a1 ` + "`" + `` + "`" + `` + "`" + ` ## Notes ### Ports and Instances There are some notes to consider when connecting Instances to networks using Ports. Please see the ` + "`" + `huaweicloud_compute_instance_v2` + "`" + ` documentation for further documentation.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "device_owner",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "fixed_ip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "all_fixed_ips",
					Description: `The collection of Fixed IP addresses on the port in the order returned by the Network v2 API.`,
				},
				resource.Attribute{
					Name:        "all_security_group_ids",
					Description: `The collection of Security Group IDs on the port which have been explicitly and implicitly added.`,
				},
				resource.Attribute{
					Name:        "extra_dhcp_option",
					Description: `See Argument Reference above. ## Import Ports can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_networking_port_v2.port_1 eae26a3e-1c33-4cc1-9c31-0cd729c438a1 ` + "`" + `` + "`" + `` + "`" + ` ## Notes ### Ports and Instances There are some notes to consider when connecting Instances to networks using Ports. Please see the ` + "`" + `huaweicloud_compute_instance_v2` + "`" + ` documentation for further documentation.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_networking_router_interface_v2",
			Category:         "Virtual Private Cloud (VPC)",
			ShortDescription: `Manages a V2 router interface resource within HuaweiCloud.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"networking",
				"router",
				"interface",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 networking client. A networking client is needed to create a router. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new router interface.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `(Required) ID of the router this interface belongs to. Changing this creates a new router interface.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of the subnet this interface connects to. Changing this creates a new router interface.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `ID of the port this interface connects to. Changing this creates a new router interface. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `See Argument Reference above. ## Import Router Interfaces can be imported using the port ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ openstack port list --router <router name or id> $ terraform import huaweicloud_networking_router_interface_v2.int_1 <port id from above output> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `See Argument Reference above. ## Import Router Interfaces can be imported using the port ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ openstack port list --router <router name or id> $ terraform import huaweicloud_networking_router_interface_v2.int_1 <port id from above output> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_networking_router_route_v2",
			Category:         "Virtual Private Cloud (VPC)",
			ShortDescription: `Creates a routing entry on a HuaweiCloud V2 router.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"networking",
				"router",
				"route",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 networking client. A networking client is needed to configure a routing entry on a router. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new routing entry.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `(Required) ID of the router this routing entry belongs to. Changing this creates a new routing entry.`,
				},
				resource.Attribute{
					Name:        "destination_cidr",
					Description: `(Required) CIDR block to match on the packet’s destination IP. Changing this creates a new routing entry.`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `(Required) IP address of the next hop gateway. Changing this creates a new routing entry. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "destination_cidr",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `See Argument Reference above. ## Notes The ` + "`" + `next_hop` + "`" + ` IP address must be directly reachable from the router at the ` + "`" + `` + "`" + `huaweicloud_networking_router_route_v2` + "`" + `` + "`" + ` resource creation time. You can ensure that by explicitly specifying a dependency on the ` + "`" + `` + "`" + `huaweicloud_networking_router_interface_v2` + "`" + `` + "`" + ` resource that connects the next hop to the router, as in the example above. ## Import Routing entries can be imported using a combined ID using the following format: ` + "`" + `` + "`" + `<router_id>-route-<destination_cidr>-<next_hop>` + "`" + `` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_networking_router_route_v2.router_route_1 686fe248-386c-4f70-9f6c-281607dad079-route-10.0.1.0/24-192.168.199.25 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "destination_cidr",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `See Argument Reference above. ## Notes The ` + "`" + `next_hop` + "`" + ` IP address must be directly reachable from the router at the ` + "`" + `` + "`" + `huaweicloud_networking_router_route_v2` + "`" + `` + "`" + ` resource creation time. You can ensure that by explicitly specifying a dependency on the ` + "`" + `` + "`" + `huaweicloud_networking_router_interface_v2` + "`" + `` + "`" + ` resource that connects the next hop to the router, as in the example above. ## Import Routing entries can be imported using a combined ID using the following format: ` + "`" + `` + "`" + `<router_id>-route-<destination_cidr>-<next_hop>` + "`" + `` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_networking_router_route_v2.router_route_1 686fe248-386c-4f70-9f6c-281607dad079-route-10.0.1.0/24-192.168.199.25 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_networking_router_v2",
			Category:         "Virtual Private Cloud (VPC)",
			ShortDescription: `Manages a V2 router resource within HuaweiCloud.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"networking",
				"router",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 networking client. A networking client is needed to create a router. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new router.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A unique name for the router. Changing this updates the ` + "`" + `name` + "`" + ` of an existing router.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) Administrative up/down status for the router (must be "true" or "false" if provided). Changing this updates the ` + "`" + `admin_state_up` + "`" + ` of an existing router.`,
				},
				resource.Attribute{
					Name:        "distributed",
					Description: `(Optional) Indicates whether or not to create a distributed router. The default policy setting in Neutron restricts usage of this property to administrative users only.`,
				},
				resource.Attribute{
					Name:        "external_network_id",
					Description: `(Optional) The network UUID of an external gateway for the router. A router with an external gateway is required if any compute instances or load balancers will be using floating IPs. Changing this updates the external gateway of the router.`,
				},
				resource.Attribute{
					Name:        "enable_snat",
					Description: `(Optional) Enable Source NAT for the router. Valid values are "true" or "false". An ` + "`" + `external_network_id` + "`" + ` has to be set in order to set this property. Changing this updates the ` + "`" + `enable_snat` + "`" + ` of the router.`,
				},
				resource.Attribute{
					Name:        "external_fixed_ip",
					Description: `(Optional) An external fixed IP for the router. This can be repeated. The structure is described below. An ` + "`" + `external_network_id` + "`" + ` has to be set in order to set this property. Changing this updates the external fixed IPs of the router.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the floating IP. Required if admin wants to create a router for another tenant. Changing this creates a new router.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional driver-specific options. The ` + "`" + `external_fixed_ip` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) Subnet in which the fixed IP belongs to.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) The IP address to set on the router. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the router.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "external_network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "enable_snat",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "external_fixed_ip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `See Argument Reference above. ## Import Routers can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_networking_router_v2.router_1 014395cd-89fc-4c9b-96b7-13d1ee79dad2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the router.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "external_network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "enable_snat",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "external_fixed_ip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `See Argument Reference above. ## Import Routers can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_networking_router_v2.router_1 014395cd-89fc-4c9b-96b7-13d1ee79dad2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_networking_secgroup_rule_v2",
			Category:         "Virtual Private Cloud (VPC)",
			ShortDescription: `Manages a V2 Neutron security group rule resource within HuaweiCloud.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"networking",
				"secgroup",
				"rule",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 networking client. A networking client is needed to create a port. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new security group rule.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Required) The direction of the rule, valid values are __ingress__ or __egress__. Changing this creates a new security group rule.`,
				},
				resource.Attribute{
					Name:        "ethertype",
					Description: `(Required) The layer 3 protocol type, valid values are __IPv4__ or __IPv6__. Changing this creates a new security group rule.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) The layer 4 protocol type, valid values are following. Changing this creates a new security group rule. This is required if you want to specify a port range.`,
				},
				resource.Attribute{
					Name:        "port_range_min",
					Description: `(Optional) The lower part of the allowed port range, valid integer value needs to be between 1 and 65535. Changing this creates a new security group rule.`,
				},
				resource.Attribute{
					Name:        "port_range_max",
					Description: `(Optional) The higher part of the allowed port range, valid integer value needs to be between 1 and 65535. Changing this creates a new security group rule.`,
				},
				resource.Attribute{
					Name:        "remote_ip_prefix",
					Description: `(Optional) The remote CIDR, the value needs to be a valid CIDR (i.e. 192.168.0.0/16). Changing this creates a new security group rule.`,
				},
				resource.Attribute{
					Name:        "remote_group_id",
					Description: `(Optional) The remote group id, the value needs to be an Openstack ID of a security group in the same tenant. Changing this creates a new security group rule.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Required) The security group id the rule should belong to, the value needs to be an Openstack ID of a security group in the same tenant. Changing this creates a new security group rule.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the security group. Required if admin wants to create a port for another tenant. Changing this creates a new security group rule. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ethertype",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_range_min",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_range_max",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "remote_ip_prefix",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "remote_group_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above. ## Import Security Group Rules can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_networking_secgroup_rule_v2.secgroup_rule_1 aeb68ee3-6e9d-4256-955c-9584a6212745 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ethertype",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_range_min",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_range_max",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "remote_ip_prefix",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "remote_group_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above. ## Import Security Group Rules can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_networking_secgroup_rule_v2.secgroup_rule_1 aeb68ee3-6e9d-4256-955c-9584a6212745 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_networking_secgroup_v2",
			Category:         "Virtual Private Cloud (VPC)",
			ShortDescription: `Manages a V2 Neutron security group resource within HuaweiCloud.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"networking",
				"secgroup",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 networking client. A networking client is needed to create a port. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new security group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the security group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A unique name for the security group.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the security group. Required if admin wants to create a port for another tenant. Changing this creates a new security group.`,
				},
				resource.Attribute{
					Name:        "delete_default_rules",
					Description: `(Optional) Whether or not to delete the default egress security rules. This is ` + "`" + `false` + "`" + ` by default. See the below note for more information. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above. ## Default Security Group Rules In most cases, HuaweiCloud will create some egress security group rules for each new security group. These security group rules will not be managed by Terraform, so if you prefer to have`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above. ## Default Security Group Rules In most cases, HuaweiCloud will create some egress security group rules for each new security group. These security group rules will not be managed by Terraform, so if you prefer to have`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_networking_subnet_v2",
			Category:         "Virtual Private Cloud (VPC)",
			ShortDescription: `Manages a V2 Neutron subnet resource within HuaweiCloud.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"networking",
				"subnet",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create a Neutron subnet. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new subnet.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required) The UUID of the parent network. Changing this creates a new subnet.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Required) CIDR representing IP range for this subnet, based on IP version. Changing this creates a new subnet.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `(Optional) IP version, either 4 (default) or 6. Changing this creates a new subnet.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the subnet. Changing this updates the name of the existing subnet.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the subnet. Required if admin wants to create a subnet for another tenant. Changing this creates a new subnet.`,
				},
				resource.Attribute{
					Name:        "allocation_pools",
					Description: `(Optional) An array of sub-ranges of CIDR available for dynamic allocation to ports. The allocation_pool object structure is documented below. Changing this creates a new subnet.`,
				},
				resource.Attribute{
					Name:        "gateway_ip",
					Description: `(Optional) Default gateway used by devices in this subnet. Leaving this blank and not setting ` + "`" + `no_gateway` + "`" + ` will cause a default gateway of ` + "`" + `.1` + "`" + ` to be used. Changing this updates the gateway IP of the existing subnet.`,
				},
				resource.Attribute{
					Name:        "no_gateway",
					Description: `(Optional) Do not set a gateway IP on this subnet. Changing this removes or adds a default gateway IP of the existing subnet.`,
				},
				resource.Attribute{
					Name:        "enable_dhcp",
					Description: `(Optional) The administrative state of the network. The value must be "true".`,
				},
				resource.Attribute{
					Name:        "dns_nameservers",
					Description: `(Optional) An array of DNS name server names used by hosts in this subnet. Changing this updates the DNS name servers for the existing subnet.`,
				},
				resource.Attribute{
					Name:        "host_routes",
					Description: `(Optional) An array of routes that should be used by devices with IPs from this subnet (not including local subnet route). The host_route object structure is documented below. Changing this updates the host routes for the existing subnet.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional options. The ` + "`" + `allocation_pools` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `(Required) The starting address.`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `(Required) The ending address. The ` + "`" + `host_routes` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "destination_cidr",
					Description: `(Required) The destination CIDR.`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `(Required) The next hop in the route. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "allocation_pools",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "gateway_ip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "enable_dhcp",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dns_nameservers",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "host_routes",
					Description: `See Argument Reference above. ## Import Subnets can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_networking_subnet_v2.subnet_1 da4faf16-5546-41e4-8330-4d0002b74048 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "allocation_pools",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "gateway_ip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "enable_dhcp",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dns_nameservers",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "host_routes",
					Description: `See Argument Reference above. ## Import Subnets can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_networking_subnet_v2.subnet_1 da4faf16-5546-41e4-8330-4d0002b74048 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_networking_vip_associate_v2",
			Category:         "Virtual Private Cloud (VPC)",
			ShortDescription: `Manages a V2 vip associate resource within HuaweiCloud.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"networking",
				"vip",
				"associate",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vip_id",
					Description: `(Required) The ID of vip to attach the port to. Changing this creates a new vip associate.`,
				},
				resource.Attribute{
					Name:        "port_ids",
					Description: `(Required) An array of one or more IDs of the ports to attach the vip to. Changing this creates a new vip associate. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "vip_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_ids",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vip_subnet_id",
					Description: `The ID of the subnet this vip connects to.`,
				},
				resource.Attribute{
					Name:        "vip_ip_address",
					Description: `The IP address in the subnet for this vip.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vip_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_ids",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vip_subnet_id",
					Description: `The ID of the subnet this vip connects to.`,
				},
				resource.Attribute{
					Name:        "vip_ip_address",
					Description: `The IP address in the subnet for this vip.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_networking_vip_v2",
			Category:         "Virtual Private Cloud (VPC)",
			ShortDescription: `Manages a V2 vip resource within HuaweiCloud.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"networking",
				"vip",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required) The ID of the network to attach the vip to. Changing this creates a new vip.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) Subnet in which to allocate IP address for this vip. Changing this creates a new vip.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) IP address desired in the subnet for this vip. If you don't specify ` + "`" + `ip_address` + "`" + `, an available IP address from the specified subnet will be allocated to this vip.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A unique name for the vip. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of vip.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the vip.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `The tenant ID of the vip.`,
				},
				resource.Attribute{
					Name:        "device_owner",
					Description: `The device owner of the vip.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of vip.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the vip.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `The tenant ID of the vip.`,
				},
				resource.Attribute{
					Name:        "device_owner",
					Description: `The device owner of the vip.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_obs_bucket",
			Category:         "Object Storage Service (OBS)",
			ShortDescription: `Provides an OBS bucket resource.`,
			Description:      ``,
			Keywords: []string{
				"object",
				"storage",
				"service",
				"obs",
				"bucket",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) Specifies the name of the bucket. Changing this parameter will create a new resource. A bucket must be named according to the globally applied DNS naming regulations as follows:`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `(Optional) Specifies the storage class of the bucket. OBS provides three storage classes: "STANDARD", "WARM" (Infrequent Access) and "COLD" (Archive). Defaults to ` + "`" + `STANDARD` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `(Optional) Specifies the ACL policy for a bucket. The predefined common policies are as follows: "private", "public-read", "public-read-write" and "log-delivery-write". Defaults to ` + "`" + `private` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the bucket. Each tag is represented by one key-value pair.`,
				},
				resource.Attribute{
					Name:        "versioning",
					Description: `(Optional) Whether enable versioning. Once you version-enable a bucket, it can never return to an unversioned state. You can, however, suspend versioning on that bucket.`,
				},
				resource.Attribute{
					Name:        "logging",
					Description: `(Optional) A settings of bucket logging (documented below).`,
				},
				resource.Attribute{
					Name:        "website",
					Description: `(Optional) A website object (documented below).`,
				},
				resource.Attribute{
					Name:        "cors_rule",
					Description: `(Optional) A rule of Cross-Origin Resource Sharing (documented below).`,
				},
				resource.Attribute{
					Name:        "lifecycle_rule",
					Description: `(Optional) A configuration of object lifecycle management (documented below).`,
				},
				resource.Attribute{
					Name:        "force_destroy",
					Description: `(Optional) A boolean that indicates all objects should be deleted from the bucket so that the bucket can be destroyed without error. Default to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Specified the region where this bucket will be created. If not specified, used the region by the provider. The ` + "`" + `logging` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "target_bucket",
					Description: `(Required) The name of the bucket that will receive the log objects. The acl policy of the target bucket should be ` + "`" + `log-delivery-write` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "target_prefix",
					Description: `(Optional) To specify a key prefix for log objects. The ` + "`" + `website` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "index_document",
					Description: `(Required, unless using ` + "`" + `redirect_all_requests_to` + "`" + `) Specifies the default homepage of the static website, only HTML web pages are supported. OBS only allows files such as ` + "`" + `index.html` + "`" + ` in the root directory of a bucket to function as the default homepage. That is to say, do not set the default homepage with a multi-level directory structure (for example, /page/index.html).`,
				},
				resource.Attribute{
					Name:        "error_document",
					Description: `(Optional) Specifies the error page returned when an error occurs during static website access. Only HTML, JPG, PNG, BMP, and WEBP files under the root directory are supported.`,
				},
				resource.Attribute{
					Name:        "redirect_all_requests_to",
					Description: `(Optional) A hostname to redirect all website requests for this bucket to. Hostname can optionally be prefixed with a protocol (` + "`" + `http://` + "`" + ` or ` + "`" + `https://` + "`" + `) to use when redirecting requests. The default is the protocol that is used in the original request.`,
				},
				resource.Attribute{
					Name:        "routing_rules",
					Description: `(Optional) A JSON or XML format containing routing rules describing redirect behavior and when redirects are applied. Each rule contains a ` + "`" + `Condition` + "`" + ` and a ` + "`" + `Redirect` + "`" + ` as shown in the following table: Parameter | Key -|- Condition | KeyPrefixEquals, HttpErrorCodeReturnedEquals Redirect | Protocol, HostName, ReplaceKeyPrefixWith, ReplaceKeyWith, HttpRedirectCode The ` + "`" + `cors_rule` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique identifier for lifecycle rules. The Rule Name contains a maximum of 255 characters.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Specifies lifecycle rule status.`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Optional) Object key prefix identifying one or more objects to which the rule applies. If omitted, all objects in the bucket will be managed by the lifecycle rule. The prefix cannot start or end with a slash (/), cannot have consecutive slashes (/), and cannot contain the following special characters: \:`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `(Optional) Specifies a period when objects that have been last updated are automatically deleted. (documented below).`,
				},
				resource.Attribute{
					Name:        "transition",
					Description: `(Optional) Specifies a period when objects that have been last updated are automatically transitioned to ` + "`" + `WARM` + "`" + ` or ` + "`" + `COLD` + "`" + ` storage class (documented below).`,
				},
				resource.Attribute{
					Name:        "noncurrent_version_expiration",
					Description: `(Optional) Specifies a period when noncurrent object versions are automatically deleted. (documented below).`,
				},
				resource.Attribute{
					Name:        "noncurrent_version_transition",
					Description: `(Optional) Specifies a period when noncurrent object versions are automatically transitioned to ` + "`" + `WARM` + "`" + ` or ` + "`" + `COLD` + "`" + ` storage class (documented below). At least one of ` + "`" + `expiration` + "`" + `, ` + "`" + `transition` + "`" + `, ` + "`" + `noncurrent_version_expiration` + "`" + `, ` + "`" + `noncurrent_version_transition` + "`" + ` must be specified. The ` + "`" + `expiration` + "`" + ` object supports the following`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `(Required) The class of storage used to store the object. Only ` + "`" + `WARM` + "`" + ` and ` + "`" + `COLD` + "`" + ` are supported. The ` + "`" + `noncurrent_version_expiration` + "`" + ` object supports the following`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `(Required) The class of storage used to store the object. Only ` + "`" + `WARM` + "`" + ` and ` + "`" + `COLD` + "`" + ` are supported. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The name of the bucket.`,
				},
				resource.Attribute{
					Name:        "bucket_domain_name",
					Description: `The bucket domain name. Will be of format ` + "`" + `bucketname.obs.region.myhuaweicloud.com` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region where this bucket resides in. ## Import OBS bucket can be imported using the ` + "`" + `bucket` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_obs_bucket.bucket bucket-name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The name of the bucket.`,
				},
				resource.Attribute{
					Name:        "bucket_domain_name",
					Description: `The bucket domain name. Will be of format ` + "`" + `bucketname.obs.region.myhuaweicloud.com` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region where this bucket resides in. ## Import OBS bucket can be imported using the ` + "`" + `bucket` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_obs_bucket.bucket bucket-name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_obs_bucket_object",
			Category:         "Object Storage Service (OBS)",
			ShortDescription: `Provides an OBS bucket object resource.`,
			Description:      ``,
			Keywords: []string{
				"object",
				"storage",
				"service",
				"obs",
				"bucket",
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
					Description: `(Optional) The literal content being uploaded to the bucket.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `(Optional) The ACL policy to apply. Defaults to ` + "`" + `private` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `(Optioanl) Specifies the storage class of the object. Defaults to ` + "`" + `STANDARD` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `(Optional) A standard MIME type describing the format of the object data, e.g. application/octet-stream. All Valid MIME Types are valid for this input.`,
				},
				resource.Attribute{
					Name:        "encryption",
					Description: `(Optional) Whether enable server-side encryption of the object in SSE-KMS mode.`,
				},
				resource.Attribute{
					Name:        "sse_kms_key_id",
					Description: `(Optional) The ID of the kms key. If omitted, the default master key will be used.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `(Optional) Specifies the unique identifier of the object content. It can be used to trigger updates. The only meaningful value is ` + "`" + `md5(file("path_to_file"))` + "`" + `. Either ` + "`" + `source` + "`" + ` or ` + "`" + `content` + "`" + ` must be provided to specify the bucket content. These two arguments are mutually-exclusive. ## Attributes Reference The following attributes are exported`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `the ` + "`" + `key` + "`" + ` of the resource supplied above.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `the ETag generated for the object (an MD5 sum of the object content). When the object is encrypted on the server side, the ETag value is not the MD5 value of the object, but the unique identifier calculated through the server-side encryption.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `the size of the object in bytes.`,
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
					Name:        "etag",
					Description: `the ETag generated for the object (an MD5 sum of the object content). When the object is encrypted on the server side, the ETag value is not the MD5 value of the object, but the unique identifier calculated through the server-side encryption.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `the size of the object in bytes.`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `A unique version ID value for the object, if bucket versioning is enabled.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_rds_instance_v1",
			Category:         "Relational Database Service (RDS)",
			ShortDescription: `Manages rds instance resource within HuaweiCloud`,
			Description:      ``,
			Keywords: []string{
				"relational",
				"database",
				"service",
				"rds",
				"instance",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the DB instance name. The DB instance name of the same type is unique in the same tenant.`,
				},
				resource.Attribute{
					Name:        "datastore",
					Description: `(Required) Specifies database information. The structure is described below.`,
				},
				resource.Attribute{
					Name:        "flavorref",
					Description: `(Required) Specifies the specification ID (flavors.id in the response message in Obtaining All DB Instance Specifications).`,
				},
				resource.Attribute{
					Name:        "volume",
					Description: `(Required) Specifies the volume information. The structure is described below.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Specifies the region ID.`,
				},
				resource.Attribute{
					Name:        "availabilityzone",
					Description: `(Required) Specifies the ID of the AZ.`,
				},
				resource.Attribute{
					Name:        "vpc",
					Description: `(Required) Specifies the VPC ID. For details about how to obtain this parameter value, see section "Virtual Private Cloud" in the Virtual Private Cloud API Reference.`,
				},
				resource.Attribute{
					Name:        "nics",
					Description: `(Required) Specifies the nics information. For details about how to obtain this parameter value, see section "Subnet" in the Virtual Private Cloud API Reference. The structure is described below.`,
				},
				resource.Attribute{
					Name:        "securitygroup",
					Description: `(Required) Specifies the security group which the RDS DB instance belongs to. The structure is described below.`,
				},
				resource.Attribute{
					Name:        "dbport",
					Description: `(Optional) Specifies the database port number.`,
				},
				resource.Attribute{
					Name:        "backupstrategy",
					Description: `(Optional) Specifies the advanced backup policy. The structure is described below.`,
				},
				resource.Attribute{
					Name:        "dbrtpd",
					Description: `(Required) Specifies the password for user root of the database.`,
				},
				resource.Attribute{
					Name:        "ha",
					Description: `(Optional) Specifies the parameters configured on HA and is used when creating HA DB instances. The structure is described below. NOTICE: RDS for Microsoft SQL Server does not support creating HA DB instances and this parameter is not involved. The ` + "`" + `datastore` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Specifies the DB engine. Currently, MySQL, PostgreSQL, and Microsoft SQL Server are supported. The value is MySQL, PostgreSQL, or SQLServer.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) Specifies the DB instance version.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Specifies the volume type. Valid value: It must be COMMON (SATA) or ULTRAHIGH (SSD) and is case-sensitive.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) Specifies the volume size. Its value must be a multiple of 10 and the value range is 100 GB to 2000 GB. The ` + "`" + `nics` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "subnetId",
					Description: `(Required) Specifies the subnet ID obtained from the VPC. The ` + "`" + `securitygroup ` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) Specifies the ID obtained from the securitygroup. The ` + "`" + `backupstrategy ` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "starttime",
					Description: `(Optional) Indicates the backup start time that has been set. The backup task will be triggered within one hour after the backup start time. Valid value: The value cannot be empty. It must use the hh:mm:ss format and must be valid. The current time is the UTC time.`,
				},
				resource.Attribute{
					Name:        "keepdays",
					Description: `(Optional) Specifies the number of days to retain the generated backup files. Its value range is 0 to 35. If this parameter is not specified or set to 0, the automated backup policy is disabled. The ` + "`" + `ha` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional) Specifies the configured parameters on the HA. Valid value: The value is true or false. The value true indicates creating HA DB instances. The value false indicates creating a single DB instance.`,
				},
				resource.Attribute{
					Name:        "replicationmode",
					Description: `(Optional) Specifies the replication mode for the standby DB instance. The value cannot be empty. For MySQL, the value is async or semisync. For PostgreSQL, the value is async or sync. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "flavorref",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "volume",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "availabilityzone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vpc",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "nics",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "securitygroup",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dbport",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "backupstrategy",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dbrtpd",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ha",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates the DB instance status.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `Indicates the instance connection address. It is a blank string.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Indicates the DB instance type, which can be master or readreplica.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `Indicates the creation time in the following format: yyyy-mm-dd Thh:mm:ssZ.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `Indicates the update time in the following format: yyyy-mm-dd Thh:mm:ssZ. ## Attributes Reference The following attributes can be updated:`,
				},
				resource.Attribute{
					Name:        "volume.size",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "flavorref",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "backupstrategy",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "flavorref",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "volume",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "availabilityzone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vpc",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "nics",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "securitygroup",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dbport",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "backupstrategy",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dbrtpd",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ha",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates the DB instance status.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `Indicates the instance connection address. It is a blank string.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Indicates the DB instance type, which can be master or readreplica.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `Indicates the creation time in the following format: yyyy-mm-dd Thh:mm:ssZ.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `Indicates the update time in the following format: yyyy-mm-dd Thh:mm:ssZ. ## Attributes Reference The following attributes can be updated:`,
				},
				resource.Attribute{
					Name:        "volume.size",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "flavorref",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "backupstrategy",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_rds_instance_v3",
			Category:         "Relational Database Service (RDS)",
			ShortDescription: `instance management`,
			Description:      ``,
			Keywords: []string{
				"relational",
				"database",
				"service",
				"rds",
				"instance",
				"v3",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required) Specifies the AZ name. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "db",
					Description: `(Required) Specifies the database information. Structure is documented below. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `(Required) Specifies the specification code.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the DB instance name. The DB instance name of the same type must be unique for the same tenant. The value must be 4 to 64 characters in length and start with a letter. It is case-sensitive and can contain only letters, digits, hyphens (-), and underscores (_). Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Required) Specifies the security group which the RDS DB instance belongs to. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) Specifies the network id of a subnet. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "volume",
					Description: `(Required) Specifies the volume information. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) Specifies the VPC ID. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "backup_strategy",
					Description: `(Optional) Specifies the advanced backup policy. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "ha_replication_mode",
					Description: `(Optional) Specifies the replication mode for the standby DB instance. For MySQL, the value is async or semisync. For PostgreSQL, the value is async or sync. For Microsoft SQL Server, the value is sync. NOTE: async indicates the asynchronous replication mode. semisync indicates the semi-synchronous replication mode. sync indicates the synchronous replication mode. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "param_group_id",
					Description: `(Optional) Specifies the parameter group ID. Changing this parameter will create a new resource. The ` + "`" + `db` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Specifies the database password. The value cannot be empty and should contain 8 to 32 characters, including uppercase and lowercase letters, digits, and the following special characters: ~!@#%^`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Specifies the database port information. The MySQL database port ranges from 1024 to 65535 (excluding 12017 and 33071, which are occupied by the RDS system and cannot be used). The PostgreSQL database port ranges from 2100 to 9500. The Microsoft SQL Server database port can be 1433 or ranges from 2100 to 9500, excluding 5355 and 5985. If this parameter is not set, the default value is as follows: For MySQL, the default value is 3306. For PostgreSQL, the default value is 5432. For Microsoft SQL Server, the default value is 1433. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Specifies the DB engine. Value: MySQL, PostgreSQL, SQLServer. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `Indicates the default user name of database.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) Specifies the database version. Changing this parameter will create a new resource. Available value for attributes: type | version ---- | --- MySQL| 5.6 <br>5.7 <br>8.0 <br>`,
				},
				resource.Attribute{
					Name:        "disk_encryption_id",
					Description: `(Optional) Specifies the key ID for disk encryption. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) Specifies the volume size. Its value range is from 40 GB to 4000 GB. The value must be a multiple of 10. Changing this resize the volume.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Specifies the volume type. Its value can be any of the following and is case-sensitive: ULTRAHIGH: indicates the SSD type. ULTRAHIGHPRO: indicates the ultra-high I/O (advanced), which supports ultra-high performance (advanced) DB instances. Changing this parameter will create a new resource. The ` + "`" + `backup_strategy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "keep_days",
					Description: `(Optional) Specifies the retention days for specific backup files. The value range is from 0 to 732. If this parameter is not specified or set to 0, the automated backup policy is disabled. NOTICE: Primary/standby DB instances of Microsoft SQL Server do not support disabling the automated backup policy.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `(Required) Specifies the backup time window. Automated backups will be triggered during the backup time window. It must be a valid value in the &quot;hh:mm-HH:MM&quot; format. The current time is in the UTC format. The HH value must be 1 greater than the hh value. The values of mm and MM must be the same and must be set to any of the following: 00, 15, 30, or 45. Example value: 08:15-09:15 23:00-00:00. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `Indicates the creation time.`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `Indicates the instance nodes information. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `Indicates the private IP address list. It is a blank string until an ECS is created.`,
				},
				resource.Attribute{
					Name:        "public_ips",
					Description: `Indicates the public IP address list. The ` + "`" + `nodes` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Indicates the AZ.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Indicates the node ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Indicates the node name.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `Indicates the node type. The value can be master or slave, indicating the primary node or standby node respectively.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates the node status. ## Timeouts This resource provides the following timeouts configuration options: - ` + "`" + `create` + "`" + ` - Default is 30 minute. ## Import RDS instance can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_rds_instance_v3.instance_1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ` But due to some attrubutes missing from the API response, it's required to ignore changes as below. ` + "`" + `` + "`" + `` + "`" + ` resource "huaweicloud_rds_instance_v3" "instance_1" { ... lifecycle { ignore_changes = [ "db", ] } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created",
					Description: `Indicates the creation time.`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `Indicates the instance nodes information. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `Indicates the private IP address list. It is a blank string until an ECS is created.`,
				},
				resource.Attribute{
					Name:        "public_ips",
					Description: `Indicates the public IP address list. The ` + "`" + `nodes` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Indicates the AZ.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Indicates the node ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Indicates the node name.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `Indicates the node type. The value can be master or slave, indicating the primary node or standby node respectively.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates the node status. ## Timeouts This resource provides the following timeouts configuration options: - ` + "`" + `create` + "`" + ` - Default is 30 minute. ## Import RDS instance can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_rds_instance_v3.instance_1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ` But due to some attrubutes missing from the API response, it's required to ignore changes as below. ` + "`" + `` + "`" + `` + "`" + ` resource "huaweicloud_rds_instance_v3" "instance_1" { ... lifecycle { ignore_changes = [ "db", ] } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_rds_parametergroup_v3",
			Category:         "Relational Database Service (RDS)",
			ShortDescription: `Manages a V3 RDS parametergroup resource within HuaweiCloud.`,
			Description:      ``,
			Keywords: []string{
				"relational",
				"database",
				"service",
				"rds",
				"parametergroup",
				"v3",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The parameter group name. It contains a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The parameter group description. It contains a maximum of 256 characters and cannot contain the following special characters:>!<"&'= the value is left blank by default.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Optional) Parameter group values key/value pairs defined by users based on the default parameter groups.`,
				},
				resource.Attribute{
					Name:        "datastore",
					Description: `(Required) Database object. The database object structure is documented below. Changing this creates a new parameter group. The ` + "`" + `datastore` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The DB engine. Currently, MySQL, PostgreSQL, and Microsoft SQL Server are supported. The value is case-insensitive and can be mysql, postgresql, or sqlserver.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) Specifies the database version.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the parameter group.`,
				},
				resource.Attribute{
					Name:        "configuration_parameters",
					Description: `Indicates the parameter configuration defined by users based on the default parameters groups.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Indicates the parameter name.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Indicates the parameter value.`,
				},
				resource.Attribute{
					Name:        "restart_required",
					Description: `Indicates whether a restart is required.`,
				},
				resource.Attribute{
					Name:        "readonly",
					Description: `Indicates whether the parameter is read-only.`,
				},
				resource.Attribute{
					Name:        "value_range",
					Description: `Indicates the parameter value range.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Indicates the parameter type.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Indicates the parameter description. ## Import Parameter groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_rds_parametergroup_v3.pg_1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the parameter group.`,
				},
				resource.Attribute{
					Name:        "configuration_parameters",
					Description: `Indicates the parameter configuration defined by users based on the default parameters groups.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Indicates the parameter name.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Indicates the parameter value.`,
				},
				resource.Attribute{
					Name:        "restart_required",
					Description: `Indicates whether a restart is required.`,
				},
				resource.Attribute{
					Name:        "readonly",
					Description: `Indicates whether the parameter is read-only.`,
				},
				resource.Attribute{
					Name:        "value_range",
					Description: `Indicates the parameter value range.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Indicates the parameter type.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Indicates the parameter description. ## Import Parameter groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_rds_parametergroup_v3.pg_1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_rts_software_config_v1",
			Category:         "Resources Template Service (RTS)",
			ShortDescription: `Provides an RTS software config resource.`,
			Description:      ``,
			Keywords: []string{
				"template",
				"service",
				"rts",
				"software",
				"config",
				"v1",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_rts_stack_v1",
			Category:         "Resources Template Service (RTS)",
			ShortDescription: `Provides an Huawei cloud RTS Stack resource.`,
			Description:      ``,
			Keywords: []string{
				"template",
				"service",
				"rts",
				"stack",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "stack_name",
					Description: `(Required) Specifies the stack name. The value must meet the regular expression rule (^[a-zA-Z][a-zA-Z0-9_.-]{0,254}$). Changing this will create a new stack.`,
				},
				resource.Attribute{
					Name:        "stack_id",
					Description: `(Required) Specifies the stack UUID.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Optional) Specifies the template. The template content must use the json syntax.`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `(Optional) Specifies the environment information about the stack.`,
				},
				resource.Attribute{
					Name:        "files",
					Description: `(Optional) Specifies files used in the environment.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `(Optional) Specifies parameter information of the stack.`,
				},
				resource.Attribute{
					Name:        "timeout_mins",
					Description: `(Optional) Specifies the timeout duration.`,
				},
				resource.Attribute{
					Name:        "template_url",
					Description: `(Optional) Specifies the template URL.`,
				},
				resource.Attribute{
					Name:        "disable_rollback",
					Description: `(Optional) Specifies whether to perform a rollback if the creation fails. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "outputs",
					Description: `A map of outputs from the stack.`,
				},
				resource.Attribute{
					Name:        "capabilities",
					Description: `List of stack capabilities for stack.`,
				},
				resource.Attribute{
					Name:        "notification_topics",
					Description: `List of notification topics for stack.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Specifies the stack status. ## Import RTS Stacks can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_rts_stack_v1.stack rts-stack ` + "`" + `` + "`" + `` + "`" + ` <a id="timeouts"></a> ## Timeouts ` + "`" + `huaweicloud_rts_stack_v1` + "`" + ` provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `30 minutes` + "`" + `) Used for Creating Stacks - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `30 minutes` + "`" + `) Used for Stack modifications - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `30 minutes` + "`" + `) Used for destroying stacks.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "outputs",
					Description: `A map of outputs from the stack.`,
				},
				resource.Attribute{
					Name:        "capabilities",
					Description: `List of stack capabilities for stack.`,
				},
				resource.Attribute{
					Name:        "notification_topics",
					Description: `List of notification topics for stack.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Specifies the stack status. ## Import RTS Stacks can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_rts_stack_v1.stack rts-stack ` + "`" + `` + "`" + `` + "`" + ` <a id="timeouts"></a> ## Timeouts ` + "`" + `huaweicloud_rts_stack_v1` + "`" + ` provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `30 minutes` + "`" + `) Used for Creating Stacks - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `30 minutes` + "`" + `) Used for Stack modifications - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `30 minutes` + "`" + `) Used for destroying stacks.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_s3_bucket",
			Category:         "S3",
			ShortDescription: `Provides a S3 bucket resource.`,
			Description:      ``,
			Keywords: []string{
				"s3",
				"bucket",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Optional, Forces new resource) The name of the bucket. If omitted, Terraform will assign a random, unique name.`,
				},
				resource.Attribute{
					Name:        "bucket_prefix",
					Description: `(Optional, Forces new resource) Creates a unique bucket name beginning with the specified prefix. Conflicts with ` + "`" + `bucket` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `(Optional) The [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) to apply. Defaults to "private".`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the bucket.`,
				},
				resource.Attribute{
					Name:        "force_destroy",
					Description: `(Optional, Default:false ) A boolean that indicates all objects should be deleted from the bucket so that the bucket can be destroyed without error. These objects are`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Optional) A valid [bucket policy](https://docs.aws.amazon.com/AmazonS3/latest/dev/example-bucket-policies.html) JSON document. Note that if the policy document is not specific enough (but still valid), Terraform may view the policy as constantly changing in a ` + "`" + `terraform plan` + "`" + `. In this case, please make sure you use the verbose/specific version of the policy.`,
				},
				resource.Attribute{
					Name:        "website",
					Description: `(Optional) A website object (documented below).`,
				},
				resource.Attribute{
					Name:        "cors_rule",
					Description: `(Optional) A rule of [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) (documented below).`,
				},
				resource.Attribute{
					Name:        "versioning",
					Description: `(Optional) A state of [versioning](https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html) (documented below)`,
				},
				resource.Attribute{
					Name:        "logging",
					Description: `(Optional) A settings of [bucket logging](https://docs.aws.amazon.com/AmazonS3/latest/UG/ManagingBucketLogging.html) (documented below).`,
				},
				resource.Attribute{
					Name:        "lifecycle_rule",
					Description: `(Optional) A configuration of [object lifecycle management](http://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) (documented below).`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) If specified, the region this bucket should reside in. Otherwise, the region used by the callee. The ` + "`" + `website` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "index_document",
					Description: `(Required, unless using ` + "`" + `redirect_all_requests_to` + "`" + `) Amazon S3 returns this index document when requests are made to the root domain or any of the subfolders.`,
				},
				resource.Attribute{
					Name:        "error_document",
					Description: `(Optional) An absolute path to the document to return in case of a 4XX error.`,
				},
				resource.Attribute{
					Name:        "redirect_all_requests_to",
					Description: `(Optional) A hostname to redirect all website requests for this bucket to. Hostname can optionally be prefixed with a protocol (` + "`" + `http://` + "`" + ` or ` + "`" + `https://` + "`" + `) to use when redirecting requests. The default is the protocol that is used in the original request.`,
				},
				resource.Attribute{
					Name:        "routing_rules",
					Description: `(Optional) A json array containing [routing rules](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules.html) describing redirect behavior and when redirects are applied. The ` + "`" + `CORS` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Enable versioning. Once you version-enable a bucket, it can never return to an unversioned state. You can, however, suspend versioning on that bucket.`,
				},
				resource.Attribute{
					Name:        "mfa_delete",
					Description: `(Optional) Enable MFA delete for either ` + "`" + `Change the versioning state of your bucket` + "`" + ` or ` + "`" + `Permanently delete an object version` + "`" + `. Default is ` + "`" + `false` + "`" + `. The ` + "`" + `logging` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "target_bucket",
					Description: `(Required) The name of the bucket that will receive the log objects.`,
				},
				resource.Attribute{
					Name:        "target_prefix",
					Description: `(Optional) To specify a key prefix for log objects. The ` + "`" + `lifecycle_rule` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Unique identifier for the rule.`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Optional) Object key prefix identifying one or more objects to which the rule applies.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Specifies lifecycle rule status.`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `(Optional) Specifies a period in the object's expire (documented below).`,
				},
				resource.Attribute{
					Name:        "noncurrent_version_expiration",
					Description: `(Optional) Specifies when noncurrent object versions expire (documented below). At least one of ` + "`" + `expiration` + "`" + `, ` + "`" + `noncurrent_version_expiration` + "`" + ` must be specified. The ` + "`" + `expiration` + "`" + ` object supports the following`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Unique identifier for the rule.`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `(Required) Specifies the destination for the rule (documented below).`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Required) Object keyname prefix identifying one or more objects to which the rule applies. Set as an empty string to replicate the whole bucket.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Required) The status of the rule. Either ` + "`" + `Enabled` + "`" + ` or ` + "`" + `Disabled` + "`" + `. The rule is ignored if status is not Enabled. The ` + "`" + `destination` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) The ARN of the S3 bucket where you want Amazon S3 to store replicas of the object identified by the rule.`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `(Optional) The class of storage used to store the object. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The name of the bucket.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The ARN of the bucket. Will be of format ` + "`" + `arn:aws:s3:::bucketname` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "bucket_domain_name",
					Description: `The bucket domain name. Will be of format ` + "`" + `bucketname.s3.amazonaws.com` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "hosted_zone_id",
					Description: `The [Route 53 Hosted Zone ID](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_website_region_endpoints) for this bucket's region.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region this bucket resides in.`,
				},
				resource.Attribute{
					Name:        "website_endpoint",
					Description: `The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.`,
				},
				resource.Attribute{
					Name:        "website_domain",
					Description: `The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string. This is used to create Route 53 alias records. ## Import S3 bucket can be imported using the ` + "`" + `bucket` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_s3_bucket.bucket bucket-name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The name of the bucket.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The ARN of the bucket. Will be of format ` + "`" + `arn:aws:s3:::bucketname` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "bucket_domain_name",
					Description: `The bucket domain name. Will be of format ` + "`" + `bucketname.s3.amazonaws.com` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "hosted_zone_id",
					Description: `The [Route 53 Hosted Zone ID](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_website_region_endpoints) for this bucket's region.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region this bucket resides in.`,
				},
				resource.Attribute{
					Name:        "website_endpoint",
					Description: `The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.`,
				},
				resource.Attribute{
					Name:        "website_domain",
					Description: `The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string. This is used to create Route 53 alias records. ## Import S3 bucket can be imported using the ` + "`" + `bucket` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_s3_bucket.bucket bucket-name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_s3_bucket_object",
			Category:         "S3",
			ShortDescription: `Provides a S3 bucket object resource.`,
			Description:      ``,
			Keywords: []string{
				"s3",
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
					Description: `(Required) The path to the source file being uploaded to the bucket.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Required unless ` + "`" + `source` + "`" + ` given) The literal content being uploaded to the bucket.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `(Optional) The [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) to apply. Defaults to "private".`,
				},
				resource.Attribute{
					Name:        "cache_control",
					Description: `(Optional) Specifies caching behavior along the request/reply chain Read [w3c cache_control](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9) for further details.`,
				},
				resource.Attribute{
					Name:        "content_disposition",
					Description: `(Optional) Specifies presentational information for the object. Read [wc3 content_disposition](http://www.w3.org/Protocols/rfc2616/rfc2616-sec19.html#sec19.5.1) for further information.`,
				},
				resource.Attribute{
					Name:        "content_encoding",
					Description: `(Optional) Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read [w3c content encoding](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.11) for further information.`,
				},
				resource.Attribute{
					Name:        "content_language",
					Description: `(Optional) The language the content is in e.g. en-US or en-GB.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `(Optional) A standard MIME type describing the format of the object data, e.g. application/octet-stream. All Valid MIME Types are valid for this input.`,
				},
				resource.Attribute{
					Name:        "website_redirect",
					Description: `(Optional) Specifies a target URL for [website redirect](http://docs.aws.amazon.com/AmazonS3/latest/dev/how-to-page-redirect.html).`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `(Optional) Used to trigger updates. The only meaningful value is ` + "`" + `${md5(file("path/to/file"))}` + "`" + `. This attribute is not compatible with ` + "`" + `kms_key_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "server_side_encryption",
					Description: `(Optional) Specifies server-side encryption of the object in S3. Valid values are "` + "`" + `AES256` + "`" + `" and "` + "`" + `aws:kms` + "`" + `".`,
				},
				resource.Attribute{
					Name:        "sse_kms_key_id",
					Description: `(Optional) The ID of the kms key. Either ` + "`" + `source` + "`" + ` or ` + "`" + `content` + "`" + ` must be provided to specify the bucket content. These two arguments are mutually-exclusive. ## Attributes Reference The following attributes are exported`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `the ` + "`" + `key` + "`" + ` of the resource supplied above`,
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
					Description: `the ` + "`" + `key` + "`" + ` of the resource supplied above`,
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
			Type:             "huaweicloud_s3_object_policy",
			Category:         "S3",
			ShortDescription: `Attaches a policy to an S3 bucket resource.`,
			Description:      ``,
			Keywords: []string{
				"s3",
				"object",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) The name of the bucket to which to apply the policy.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Required) The text of the policy.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_sfs_file_system_v2",
			Category:         "Scalable File Service (SFS)",
			ShortDescription: `Provides an Scalable File Resource (SFS) resource.`,
			Description:      ``,
			Keywords: []string{
				"scalable",
				"file",
				"service",
				"sfs",
				"system",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "size",
					Description: `(Required) The size (GB) of the shared file system.`,
				},
				resource.Attribute{
					Name:        "share_proto",
					Description: `(Optional) The protocol for sharing file systems. The default value is NFS.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the shared file system.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Describes the shared file system.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `(Optional) The level of visibility for the shared file system.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) Metadata key and value pairs as a dictionary of strings.Changing this will create a new resource.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) The availability zone name.Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "access_level",
					Description: `(Required) The access level of the shared file system. Changing this will create a new access rule.`,
				},
				resource.Attribute{
					Name:        "access_type",
					Description: `(Optional) The type of the share access rule. Changing this will create a new access rule.`,
				},
				resource.Attribute{
					Name:        "access_to",
					Description: `(Required) The access that the back end grants or denies. Changing this will create a new access rule ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the shared file system.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the shared file system.`,
				},
				resource.Attribute{
					Name:        "share_type",
					Description: `The storage service type assigned for the shared file system, such as high-performance storage (composed of SSDs) and large-capacity storage (composed of SATA disks).`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `The volume type.`,
				},
				resource.Attribute{
					Name:        "export_location",
					Description: `The address for accessing the shared file system.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `The host name of the shared file system.`,
				},
				resource.Attribute{
					Name:        "share_access_id",
					Description: `The UUID of the share access rule.`,
				},
				resource.Attribute{
					Name:        "access_rules_status",
					Description: `The status of the share access rule. ## Import SFS can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` > $ terraform import huaweicloud_sfs_file_system_v2 4779ab1c-7c1a-44b1-a02e-93dfc361b32d ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the shared file system.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the shared file system.`,
				},
				resource.Attribute{
					Name:        "share_type",
					Description: `The storage service type assigned for the shared file system, such as high-performance storage (composed of SSDs) and large-capacity storage (composed of SATA disks).`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `The volume type.`,
				},
				resource.Attribute{
					Name:        "export_location",
					Description: `The address for accessing the shared file system.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `The host name of the shared file system.`,
				},
				resource.Attribute{
					Name:        "share_access_id",
					Description: `The UUID of the share access rule.`,
				},
				resource.Attribute{
					Name:        "access_rules_status",
					Description: `The status of the share access rule. ## Import SFS can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` > $ terraform import huaweicloud_sfs_file_system_v2 4779ab1c-7c1a-44b1-a02e-93dfc361b32d ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_smn_subscription_v2",
			Category:         "Simple Message Notification (SMN)",
			ShortDescription: `Manages a V2 subscription resource within HuaweiCloud.`,
			Description:      ``,
			Keywords: []string{
				"simple",
				"message",
				"notification",
				"smn",
				"subscription",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "topic_urn",
					Description: `(Required) Resource identifier of a topic, which is unique.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Required) Message endpoint. For an HTTP subscription, the endpoint starts with http\://. For an HTTPS subscription, the endpoint starts with https\://. For an email subscription, the endpoint is a mail address. For an SMS message subscription, the endpoint is a phone number.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Protocol of the message endpoint. Currently, email, sms, http, and https are supported.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `(Optional) Remark information. The remarks must be a UTF-8-coded character string containing 128 bytes.`,
				},
				resource.Attribute{
					Name:        "subscription_urn",
					Description: `(Optional) Resource identifier of a subscription, which is unique.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `(Optional) Project ID of the topic creator.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Subscription status. 0 indicates that the subscription is not confirmed. 1 indicates that the subscription is confirmed. 3 indicates that the subscription is canceled. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "topic_urn",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subscription_urn",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "topic_urn",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subscription_urn",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_smn_topic_v2",
			Category:         "Simple Message Notification (SMN)",
			ShortDescription: `Manages a V2 topic resource within HuaweiCloud.`,
			Description:      ``,
			Keywords: []string{
				"simple",
				"message",
				"notification",
				"smn",
				"topic",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the topic to be created.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Topic display name, which is presented as the name of the email sender in an email message.`,
				},
				resource.Attribute{
					Name:        "topic_urn",
					Description: `(Optional) Resource identifier of a topic, which is unique.`,
				},
				resource.Attribute{
					Name:        "push_policy",
					Description: `(Optional) Message pushing policy. 0 indicates that the message sending fails and the message is cached in the queue. 1 indicates that the failed message is discarded.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `(Optional) Time when the topic was created.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `(Optional) Time when the topic was updated. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "topic_urn",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "push_policy",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "topic_urn",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "push_policy",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud-vbs-backup-policy-v2",
			Category:         "Volume Backup Service (VBS)",
			ShortDescription: `Provides an VBS Backup Policy resource.`,
			Description:      ``,
			Keywords: []string{
				"volume",
				"backup",
				"service",
				"vbs",
				"huaweicloud-vbs-backup-policy-v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the policy name. The value is a string of 1 to 64 characters that can contain letters, digits, underscores (_), and hyphens (-). It cannot start with`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `(Required) Specifies the start time(UTC) of the backup job. The value is in the HH:mm format. You need to set the execution time on a full hour. You can set multiple execution times, and use commas (,) to separate one time from another.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Specifies the backup policy status. Possible values are ON or OFF. Defaults to ON.`,
				},
				resource.Attribute{
					Name:        "retain_first_backup",
					Description: `(Required) Specifies whether to retain the first backup in the current month. Possible values are Y or N.`,
				},
				resource.Attribute{
					Name:        "rentention_num",
					Description: `(Optional) Specifies number of retained backups. Minimum value is 2. Either this field or ` + "`" + `rentention_day` + "`" + ` must be specified.`,
				},
				resource.Attribute{
					Name:        "rentention_day",
					Description: `(Optional) Specifies days of retained backups. Minimum value is 2. Either this field or ` + "`" + `rentention_num` + "`" + ` must be specified.`,
				},
				resource.Attribute{
					Name:        "frequency",
					Description: `(Optional) Specifies the backup interval. The value is in the range of 1 to 14 days. Either this field or ` + "`" + `week_frequency` + "`" + ` must be specified.`,
				},
				resource.Attribute{
					Name:        "week_frequency",
					Description: `(Optional) Specifies on which days of each week backup jobs are executed. The value can be one or more of the following: SUN, MON, TUE, WED, THU, FRI, SAT. Either this field or ` + "`" + `frequency` + "`" + ` must be specified.`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `(Optional) Specifies one or more volumes associated with the backup policy. Any previously associated backup policy will no longer apply.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Represents the list of tags to be configured for the backup policy.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Specifies the tag key. A tag key consists of up to 36 characters, chosen from letters, digits, hyphens (-), and underscores (_).`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Specifies the tag value. A tag value consists of 0 to 43 characters, chosen from letters, digits, hyphens (-), and underscores (_). ## Attributes Reference All of the argument attributes are also exported as result attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Specifies a backup policy ID.`,
				},
				resource.Attribute{
					Name:        "policy_resource_count",
					Description: `Specifies the number of volumes associated with the backup policy. ## Import Backup Policy can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_vbs_backup_policy_v2.vbs 4779ab1c-7c1a-44b1-a02e-93dfc361b32d ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Specifies a backup policy ID.`,
				},
				resource.Attribute{
					Name:        "policy_resource_count",
					Description: `Specifies the number of volumes associated with the backup policy. ## Import Backup Policy can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_vbs_backup_policy_v2.vbs 4779ab1c-7c1a-44b1-a02e-93dfc361b32d ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud-vbs-backup-v2",
			Category:         "Volume Backup Service (VBS)",
			ShortDescription: `Provides an VBS Backup resource.`,
			Description:      ``,
			Keywords: []string{
				"volume",
				"backup",
				"service",
				"vbs",
				"huaweicloud-vbs-backup-v2",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_vpc_bandwidth_v2",
			Category:         "Virtual Private Cloud (VPC)",
			ShortDescription: `Manages a V2 Shared Bandwidth resource within HuaweiCloud.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"bandwidth",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Shared Bandwidth.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) The size of the Shared Bandwidth. The value ranges from 5 to 2000 G.`,
				},
				resource.Attribute{
					Name:        "enterprise_project_id",
					Description: `(Optional) The enterprise project id of the Shared Bandwidth. Changing this creates a new bandwidth. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Shared Bandwidth.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "enterprise_project_id",
					Description: `See Argument Reference above. ## Import Shared Bandwidths can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_vpc_bandwidth_v2.bandwidth_1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Shared Bandwidth.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "enterprise_project_id",
					Description: `See Argument Reference above. ## Import Shared Bandwidths can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_vpc_bandwidth_v2.bandwidth_1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_vpc_eip_v1",
			Category:         "Virtual Private Cloud (VPC)",
			ShortDescription: `Manages a V1 EIP resource within Huawei Cloud VPC.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"eip",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to create the eip. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new eip.`,
				},
				resource.Attribute{
					Name:        "publicip",
					Description: `(Required) The elastic IP address object.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Required) The bandwidth object. The ` + "`" + `publicip` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The value must be a type supported by the system. Only ` + "`" + `5_bgp` + "`" + ` supported now. Changing this creates a new eip.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) The value must be a valid IP address in the available IP address segment. Changing this creates a new eip.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `(Optional) The port id which this eip will associate with. If the value is "" or this not specified, the eip will be in unbind state. The ` + "`" + `bandwidth` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The bandwidth name, which is a string of 1 to 64 characters that contain letters, digits, underscores (_), and hyphens (-).`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) The bandwidth size. The value ranges from 1 to 300 Mbit/s.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The share bandwidth id. Changing this creates a new eip.`,
				},
				resource.Attribute{
					Name:        "share_type",
					Description: `(Required) Whether the bandwidth is shared or exclusive. Changing this creates a new eip.`,
				},
				resource.Attribute{
					Name:        "charge_mode",
					Description: `(Optional) This is a reserved field. If the system supports charging by traffic and this field is specified, then you are charged by traffic for elastic IP addresses. Changing this creates a new eip. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "publicip/type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "publicip/ip_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "publicip/port_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "bandwidth/name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "bandwidth/size",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "bandwidth/share_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "bandwidth/charge_mode",
					Description: `See Argument Reference above. ## Import EIPs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_vpc_eip_v1.eip_1 2c7f39f3-702b-48d1-940c-b50384177ee1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "publicip/type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "publicip/ip_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "publicip/port_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "bandwidth/name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "bandwidth/size",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "bandwidth/share_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "bandwidth/charge_mode",
					Description: `See Argument Reference above. ## Import EIPs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_vpc_eip_v1.eip_1 2c7f39f3-702b-48d1-940c-b50384177ee1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_vpc_peering_connection_accepter_v2",
			Category:         "Virtual Private Cloud (VPC)",
			ShortDescription: `Manage the accepter's side of a VPC Peering Connection.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"peering",
				"connection",
				"accepter",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The VPC peering connection name.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The VPC peering connection ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The VPC peering connection status.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of requester VPC involved in a VPC peering connection.`,
				},
				resource.Attribute{
					Name:        "peer_vpc_id",
					Description: `The VPC ID of the accepter tenant.`,
				},
				resource.Attribute{
					Name:        "peer_tenant_id",
					Description: `The Tenant Id of the accepter tenant.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The VPC peering connection name.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The VPC peering connection ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The VPC peering connection status.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of requester VPC involved in a VPC peering connection.`,
				},
				resource.Attribute{
					Name:        "peer_vpc_id",
					Description: `The VPC ID of the accepter tenant.`,
				},
				resource.Attribute{
					Name:        "peer_tenant_id",
					Description: `The Tenant Id of the accepter tenant.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_vpc_peering_connection_v2",
			Category:         "Virtual Private Cloud (VPC)",
			ShortDescription: `Manage a VPC Peering Connection resource.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"peering",
				"connection",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The VPC peering connection ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The VPC peering connection status. The value can be PENDING_ACCEPTANCE, REJECTED, EXPIRED, DELETED, or ACTIVE. ## Notes If you create a VPC peering connection with another VPC of your own, the connection is created without the need for you to accept the connection. ## Import VPC Peering resources can be imported using the ` + "`" + `vpc peering id` + "`" + `, e.g. > $ terraform import huaweicloud_vpc_peering_connection_v2.test_connection 22b76469-08e3-4937-8c1d-7aad34892be1`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The VPC peering connection ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The VPC peering connection status. The value can be PENDING_ACCEPTANCE, REJECTED, EXPIRED, DELETED, or ACTIVE. ## Notes If you create a VPC peering connection with another VPC of your own, the connection is created without the need for you to accept the connection. ## Import VPC Peering resources can be imported using the ` + "`" + `vpc peering id` + "`" + `, e.g. > $ terraform import huaweicloud_vpc_peering_connection_v2.test_connection 22b76469-08e3-4937-8c1d-7aad34892be1`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_vpc_route_v2",
			Category:         "Virtual Private Cloud (VPC)",
			ShortDescription: `Provides an VPC route resource.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"route",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The route ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The route ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_vpc_subnet_v1",
			Category:         "Virtual Private Cloud (VPC)",
			ShortDescription: `Provides an VPC subnet resource.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"subnet",
				"v1",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_vpc_v1",
			Category:         "Virtual Private Cloud (VPC)",
			ShortDescription: `Manages a V1 VPC resource within HuaweiCloud.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cidr",
					Description: `(Required) The range of available subnets in the VPC. The value ranges from 10.0.0.0/8 to 10.255.255.0/24, 172.16.0.0/12 to 172.31.255.0/24, or 192.168.0.0/16 to 192.168.255.0/24.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V1 VPC client. A VPC client is needed to create a VPC. If omitted, the region argument of the provider is used. Changing this creates a new VPC.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the VPC. The name must be unique for a tenant. The value is a string of no more than 64 characters and can contain digits, letters, underscores (_), and hyphens (-). Changing this updates the name of the existing VPC.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value pairs to associate with the vpc. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the VPC.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current status of the desired VPC. Can be either CREATING, OK, DOWN, PENDING_UPDATE, PENDING_DELETE, or ERROR.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `Specifies whether the cross-tenant sharing is supported.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above. ## Import VPCs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_vpc_v1.vpc_v1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the VPC.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current status of the desired VPC. Can be either CREATING, OK, DOWN, PENDING_UPDATE, PENDING_DELETE, or ERROR.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `Specifies whether the cross-tenant sharing is supported.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above. ## Import VPCs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_vpc_v1.vpc_v1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_vpnaas_endpoint_group_v2",
			Category:         "Virtual Private Network (VPN)",
			ShortDescription: `Manages a V2 Endpoint Group resource within HuaweiCloud.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"network",
				"vpn",
				"vpnaas",
				"endpoint",
				"group",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create an endpoint group. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the group. Changing this updates the name of the existing group.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the group. Required if admin wants to create an endpoint group for another project. Changing this creates a new group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The human-readable description for the group. Changing this updates the description of the existing group.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the endpoints in the group. A valid value is subnet, cidr, network, router, or vlan. Changing this creates a new group.`,
				},
				resource.Attribute{
					Name:        "endpoints",
					Description: `List of endpoints of the same type, for the endpoint group. The values will depend on the type. Changing this creates a new group.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional options. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "endpoints",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `See Argument Reference above. ## Import Groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_vpnaas_endpoint_group_v2.group_1 832cb7f3-59fe-40cf-8f64-8350ffc03272 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "endpoints",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `See Argument Reference above. ## Import Groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_vpnaas_endpoint_group_v2.group_1 832cb7f3-59fe-40cf-8f64-8350ffc03272 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_vpnaas_ike_policy_v2",
			Category:         "Virtual Private Network (VPN)",
			ShortDescription: `Manages a V2 IKE policy resource within HuaweiCloud.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"network",
				"vpn",
				"vpnaas",
				"ike",
				"policy",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create a VPN service. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new service.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the policy. Changing this updates the name of the existing policy.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the policy. Required if admin wants to create a service for another policy. Changing this creates a new policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The human-readable description for the policy. Changing this updates the description of the existing policy.`,
				},
				resource.Attribute{
					Name:        "auth_algorithm",
					Description: `(Optional) The authentication hash algorithm. Valid values are md5, sha1, sha2-256, sha2-384, sha2-512. Default is sha1. Changing this updates the algorithm of the existing policy.`,
				},
				resource.Attribute{
					Name:        "encryption_algorithm",
					Description: `(Optional) The encryption algorithm. Valid values are 3des, aes-128, aes-192 and so on. The default value is aes-128. Changing this updates the existing policy.`,
				},
				resource.Attribute{
					Name:        "pfs",
					Description: `(Optional) The perfect forward secrecy mode. Valid values are Group2, Group5 and Group14. Default is Group5. Changing this updates the existing policy.`,
				},
				resource.Attribute{
					Name:        "phase1_negotiation_mode",
					Description: `(Optional) The IKE mode. A valid value is main, which is the default. Changing this updates the existing policy.`,
				},
				resource.Attribute{
					Name:        "ike_version",
					Description: `(Optional) The IKE mode. A valid value is v1 or v2. Default is v1. Changing this updates the existing policy.`,
				},
				resource.Attribute{
					Name:        "lifetime",
					Description: `(Optional) The lifetime of the security association. Consists of Unit and Value. - ` + "`" + `unit` + "`" + ` - (Optional) The units for the lifetime of the security association. Can be either seconds or kilobytes. Default is seconds. - ` + "`" + `value` + "`" + ` - (Optional) The value for the lifetime of the security association. Must be a positive integer. Default is 3600.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional options. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "auth_algorithm",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "encapsulation_mode",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "encryption_algorithm",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "pfs",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "transform_protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "lifetime",
					Description: `See Argument Reference above. - ` + "`" + `unit` + "`" + ` - See Argument Reference above. - ` + "`" + `value` + "`" + ` - See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `See Argument Reference above. ## Import Services can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_vpnaas_ike_policy_v2.policy_1 832cb7f3-59fe-40cf-8f64-8350ffc03272 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "auth_algorithm",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "encapsulation_mode",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "encryption_algorithm",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "pfs",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "transform_protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "lifetime",
					Description: `See Argument Reference above. - ` + "`" + `unit` + "`" + ` - See Argument Reference above. - ` + "`" + `value` + "`" + ` - See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `See Argument Reference above. ## Import Services can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_vpnaas_ike_policy_v2.policy_1 832cb7f3-59fe-40cf-8f64-8350ffc03272 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_vpnaas_ipsec_policy_v2",
			Category:         "Virtual Private Network (VPN)",
			ShortDescription: `Manages a V2 IPSec policy resource within HuaweiCloud.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"network",
				"vpn",
				"vpnaas",
				"ipsec",
				"policy",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create an IPSec policy. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the policy. Changing this updates the name of the existing policy.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the policy. Required if admin wants to create a policy for another project. Changing this creates a new policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The human-readable description for the policy. Changing this updates the description of the existing policy.`,
				},
				resource.Attribute{
					Name:        "auth_algorithm",
					Description: `(Optional) The authentication hash algorithm. Valid values are md5, sha1, sha2-256, sha2-384, sha2-512. Default is sha1. Changing this updates the algorithm of the existing policy.`,
				},
				resource.Attribute{
					Name:        "encapsulation_mode",
					Description: `(Optional) The encapsulation mode. Valid values are tunnel and transport. Default is tunnel. Changing this updates the existing policy.`,
				},
				resource.Attribute{
					Name:        "encryption_algorithm",
					Description: `(Optional) The encryption algorithm. Valid values are 3des, aes-128, aes-192 and so on. The default value is aes-128. Changing this updates the existing policy.`,
				},
				resource.Attribute{
					Name:        "pfs",
					Description: `(Optional) The perfect forward secrecy mode. Valid values are Group2, Group5 and Group14. Default is Group5. Changing this updates the existing policy.`,
				},
				resource.Attribute{
					Name:        "transform_protocol",
					Description: `(Optional) The transform protocol. Valid values are ESP, AH and AH-ESP. Changing this updates the existing policy. Default is ESP.`,
				},
				resource.Attribute{
					Name:        "lifetime",
					Description: `(Optional) The lifetime of the security association. Consists of Unit and Value. - ` + "`" + `unit` + "`" + ` - (Optional) The units for the lifetime of the security association. Can be either seconds or kilobytes. Default is seconds. - ` + "`" + `value` + "`" + ` - (Optional) The value for the lifetime of the security association. Must be a positive integer. Default is 3600.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional options. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "auth_algorithm",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "encapsulation_mode",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "encryption_algorithm",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "pfs",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "transform_protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "lifetime",
					Description: `See Argument Reference above. - ` + "`" + `unit` + "`" + ` - See Argument Reference above. - ` + "`" + `value` + "`" + ` - See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `See Argument Reference above. ## Import Policies can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_vpnaas_ipsec_policy_v2.policy_1 832cb7f3-59fe-40cf-8f64-8350ffc03272 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "auth_algorithm",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "encapsulation_mode",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "encryption_algorithm",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "pfs",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "transform_protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "lifetime",
					Description: `See Argument Reference above. - ` + "`" + `unit` + "`" + ` - See Argument Reference above. - ` + "`" + `value` + "`" + ` - See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `See Argument Reference above. ## Import Policies can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_vpnaas_ipsec_policy_v2.policy_1 832cb7f3-59fe-40cf-8f64-8350ffc03272 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_vpnaas_service_v2",
			Category:         "Virtual Private Network (VPN)",
			ShortDescription: `Manages a V2 VPN service resource within HuaweiCloud.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"network",
				"vpn",
				"vpnaas",
				"service",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create a VPN service. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new service.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the service. Changing this updates the name of the existing service.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the service. Required if admin wants to create a service for another project. Changing this creates a new service.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The human-readable description for the service. Changing this updates the description of the existing service.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) The administrative state of the resource. Can either be up(true) or down (false). Defaults to ` + "`" + `true` + "`" + `. Changing this updates the administrative state of the existing service.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) SubnetID is the ID of the subnet. Default is null.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `(Required) The ID of the router. Changing this creates a new service.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional options. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates whether IPsec VPN service is currently operational. Values are ACTIVE, DOWN, BUILD, ERROR, PENDING_CREATE, PENDING_UPDATE, or PENDING_DELETE.`,
				},
				resource.Attribute{
					Name:        "external_v6_ip",
					Description: `The read-only external (public) IPv6 address that is used for the VPN service.`,
				},
				resource.Attribute{
					Name:        "external_v4_ip",
					Description: `The read-only external (public) IPv4 address that is used for the VPN service.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `See Argument Reference above. ## Import Services can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_vpnaas_service_v2.service_1 832cb7f3-59fe-40cf-8f64-8350ffc03272 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates whether IPsec VPN service is currently operational. Values are ACTIVE, DOWN, BUILD, ERROR, PENDING_CREATE, PENDING_UPDATE, or PENDING_DELETE.`,
				},
				resource.Attribute{
					Name:        "external_v6_ip",
					Description: `The read-only external (public) IPv6 address that is used for the VPN service.`,
				},
				resource.Attribute{
					Name:        "external_v4_ip",
					Description: `The read-only external (public) IPv4 address that is used for the VPN service.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `See Argument Reference above. ## Import Services can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_vpnaas_service_v2.service_1 832cb7f3-59fe-40cf-8f64-8350ffc03272 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_vpnaas_site_connection_v2",
			Category:         "Virtual Private Network (VPN)",
			ShortDescription: `Manages a V2 IPSec site connection resource within HuaweiCloud.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"network",
				"vpn",
				"vpnaas",
				"site",
				"connection",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create an IPSec site connection. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new site connection.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the connection. Changing this updates the name of the existing connection.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the connection. Required if admin wants to create a connection for another project. Changing this creates a new connection.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The human-readable description for the connection. Changing this updates the description of the existing connection.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) The administrative state of the resource. Can either be up(true) or down(false). Defaults to ` + "`" + `true` + "`" + `. Changing this updates the administrative state of the existing connection.`,
				},
				resource.Attribute{
					Name:        "ikepolicy_id",
					Description: `(Required) The ID of the IKE policy. Changing this creates a new connection.`,
				},
				resource.Attribute{
					Name:        "vpnservice_id",
					Description: `(Required) The ID of the VPN service. Changing this creates a new connection.`,
				},
				resource.Attribute{
					Name:        "local_ep_group_id",
					Description: `(Optional) The ID for the endpoint group that contains private subnets for the local side of the connection. You must specify this parameter with the peer_ep_group_id parameter unless in backward- compatible mode where peer_cidrs is provided with a subnet_id for the VPN service. Changing this updates the existing connection.`,
				},
				resource.Attribute{
					Name:        "ipsecpolicy_id",
					Description: `(Required) The ID of the IPsec policy. Changing this creates a new connection.`,
				},
				resource.Attribute{
					Name:        "peer_id",
					Description: `(Required) The peer router identity for authentication. A valid value is an IPv4 address, IPv6 address, e-mail address, key ID, or FQDN. Typically, this value matches the peer_address value. Changing this updates the existing policy.`,
				},
				resource.Attribute{
					Name:        "peer_ep_group_id",
					Description: `(Optional) The ID for the endpoint group that contains private CIDRs in the form < net_address > / < prefix > for the peer side of the connection. You must specify this parameter with the local_ep_group_id parameter unless in backward-compatible mode where peer_cidrs is provided with a subnet_id for the VPN service.`,
				},
				resource.Attribute{
					Name:        "local_id",
					Description: `(Optional) An ID to be used instead of the external IP address for a virtual router used in traffic between instances on different networks in east-west traffic. Most often, local ID would be domain name, email address, etc. If this is not configured then the external IP address will be used as the ID.`,
				},
				resource.Attribute{
					Name:        "peer_address",
					Description: `(Required) The peer gateway public IPv4 or IPv6 address or FQDN.`,
				},
				resource.Attribute{
					Name:        "psk",
					Description: `(Required) The pre-shared key. A valid value is any string.`,
				},
				resource.Attribute{
					Name:        "initiator",
					Description: `(Optional) A valid value is response-only or bi-directional. Default is bi-directional.`,
				},
				resource.Attribute{
					Name:        "peer_cidrs",
					Description: `(Optional) Unique list of valid peer private CIDRs in the form < net_address > / < prefix > .`,
				},
				resource.Attribute{
					Name:        "dpd",
					Description: `(Optional) A dictionary with dead peer detection (DPD) protocol controls. - ` + "`" + `action` + "`" + ` - (Optional) The dead peer detection (DPD) action. A valid value is clear, hold, restart, disabled, or restart-by-peer. Default value is hold. - ` + "`" + `timeout` + "`" + ` - (Optional) The dead peer detection (DPD) timeout in seconds. A valid value is a positive integer that is greater than the DPD interval value. Default is 120. - ` + "`" + `interval` + "`" + ` - (Optional) The dead peer detection (DPD) interval, in seconds. A valid value is a positive integer. Default is 30.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `(Optional) The maximum transmission unit (MTU) value to address fragmentation. Minimum value is 68 for IPv4, and 1280 for IPv6.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional options. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dpd",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "psk",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "initiator",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "peer_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "peer_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "peer_cidrs",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "local_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "peer_ep_group_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ipsecpolicy_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vpnservice_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ikepolicy_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `See Argument Reference above. ## Import Site Connections can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_vpnaas_site_connection_v2.conn_1 832cb7f3-59fe-40cf-8f64-8350ffc03272 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dpd",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "psk",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "initiator",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "peer_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "peer_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "peer_cidrs",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "local_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "peer_ep_group_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ipsecpolicy_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vpnservice_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ikepolicy_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `See Argument Reference above. ## Import Site Connections can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloud_vpnaas_site_connection_v2.conn_1 832cb7f3-59fe-40cf-8f64-8350ffc03272 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"huaweicloud_api_gateway_api":                    0,
		"huaweicloud_api_gateway_group":                  1,
		"huaweicloud_as_configuration_v1":                2,
		"huaweicloud_as_group_v1":                        3,
		"huaweicloud_as_policy_v1":                       4,
		"huaweicloud_blockstorage_volume_v2":             5,
		"huaweicloud_cce_cluster_v3":                     6,
		"huaweicloud_cce_nodes_v3":                       7,
		"huaweicloud_cci_network_v1":                     8,
		"huaweicloud_cdm_cluster_v1":                     9,
		"huaweicloud_cdn_domain_v1":                      10,
		"huaweicloud_ces-alarmrule":                      11,
		"huaweicloud_cloudtable_cluster_v2":              12,
		"huaweicloud_compute_floatingip_associate_v2":    13,
		"huaweicloud_compute_floatingip_v2":              14,
		"huaweicloud_compute_instance_v2":                15,
		"huaweicloud_compute_interface_attach_v2":        16,
		"huaweicloud_compute_keypair_v2":                 17,
		"huaweicloud_compute_secgroup_v2":                18,
		"huaweicloud_compute_servergroup_v2":             19,
		"huaweicloud_compute_volume_attach_v2":           20,
		"huaweicloud_cs_cluster_v1":                      21,
		"huaweicloud_cs_peering_connect_v1":              22,
		"huaweicloud_cs_route_v1":                        23,
		"huaweicloud_csbs_backup_policy_v1":              24,
		"huaweicloud_csbs_backup_v1":                     25,
		"huaweicloud_css_cluster_v1":                     26,
		"huaweicloud_cts_tracker_v1":                     27,
		"huaweicloud_dcs_instance_v1":                    28,
		"huaweicloud_dds_instance_v3":                    29,
		"huaweicloud_dis_stream_v2":                      30,
		"huaweicloud_dli_queue_v1":                       31,
		"huaweicloud_dms_group_v1":                       32,
		"huaweicloud_dms_instance_v1":                    33,
		"huaweicloud_dms_queue_v1":                       34,
		"huaweicloud_dns_ptrrecord_v2":                   35,
		"huaweicloud_dns_recordset_v2":                   36,
		"huaweicloud_dns_zone_v2":                        37,
		"huaweicloud_dws_cluster":                        38,
		"huaweicloud_ecs_instance_v1":                    39,
		"huaweicloud_elb_backendecs":                     40,
		"huaweicloud_elb_healthcheck":                    41,
		"huaweicloud_elb_listener":                       42,
		"huaweicloud_elb_loadbalancer":                   43,
		"huaweicloud_evs_snapshot":                       44,
		"huaweicloud_fgs_function_v2":                    45,
		"huaweicloud_fw_firewall_group_v2":               46,
		"huaweicloud_fw_policy_v2":                       47,
		"huaweicloud_fw_rule_v2":                         48,
		"huaweicloud_ges_graph_v1":                       49,
		"huaweicloud_iam_agency_v3":                      50,
		"huaweicloud_identity_group_membership_v3":       51,
		"huaweicloud_identity_group_v3":                  52,
		"huaweicloud_identity_project_v3":                53,
		"huaweicloud_identity_role_assignment_v3":        54,
		"huaweicloud_identity_user_v3":                   55,
		"huaweicloud_images_image_v2":                    56,
		"huaweicloud_kms_key_v1":                         57,
		"huaweicloud_lb_certificate_v2":                  58,
		"huaweicloud_lb_l7policy_v2":                     59,
		"huaweicloud_lb_l7rule_v2":                       60,
		"huaweicloud_lb_listener_v2":                     61,
		"huaweicloud_lb_loadbalancer_v2":                 62,
		"huaweicloud_lb_member_v2":                       63,
		"huaweicloud_lb_monitor_v2":                      64,
		"huaweicloud_lb_pool_v2":                         65,
		"huaweicloud_lb_whitelist_v2":                    66,
		"huaweicloud_maas_task_v1":                       67,
		"huaweicloud_mls_instance":                       68,
		"huaweicloud_mrs_cluster_v1":                     69,
		"huaweicloud_mrs_job_v1":                         70,
		"huaweicloud_nat_dnat_rule_v2":                   71,
		"huaweicloud_nat_gateway_v2":                     72,
		"huaweicloud_nat_snat_rule_v2":                   73,
		"huaweicloud_networking_floatingip_associate_v2": 74,
		"huaweicloud_networking_floatingip_v2":           75,
		"huaweicloud_networking_network_v2":              76,
		"huaweicloud_networking_port_v2":                 77,
		"huaweicloud_networking_router_interface_v2":     78,
		"huaweicloud_networking_router_route_v2":         79,
		"huaweicloud_networking_router_v2":               80,
		"huaweicloud_networking_secgroup_rule_v2":        81,
		"huaweicloud_networking_secgroup_v2":             82,
		"huaweicloud_networking_subnet_v2":               83,
		"huaweicloud_networking_vip_associate_v2":        84,
		"huaweicloud_networking_vip_v2":                  85,
		"huaweicloud_obs_bucket":                         86,
		"huaweicloud_obs_bucket_object":                  87,
		"huaweicloud_rds_instance_v1":                    88,
		"huaweicloud_rds_instance_v3":                    89,
		"huaweicloud_rds_parametergroup_v3":              90,
		"huaweicloud_rts_software_config_v1":             91,
		"huaweicloud_rts_stack_v1":                       92,
		"huaweicloud_s3_bucket":                          93,
		"huaweicloud_s3_bucket_object":                   94,
		"huaweicloud_s3_object_policy":                   95,
		"huaweicloud_sfs_file_system_v2":                 96,
		"huaweicloud_smn_subscription_v2":                97,
		"huaweicloud_smn_topic_v2":                       98,
		"huaweicloud-vbs-backup-policy-v2":               99,
		"huaweicloud-vbs-backup-v2":                      100,
		"huaweicloud_vpc_bandwidth_v2":                   101,
		"huaweicloud_vpc_eip_v1":                         102,
		"huaweicloud_vpc_peering_connection_accepter_v2": 103,
		"huaweicloud_vpc_peering_connection_v2":          104,
		"huaweicloud_vpc_route_v2":                       105,
		"huaweicloud_vpc_subnet_v1":                      106,
		"huaweicloud_vpc_v1":                             107,
		"huaweicloud_vpnaas_endpoint_group_v2":           108,
		"huaweicloud_vpnaas_ike_policy_v2":               109,
		"huaweicloud_vpnaas_ipsec_policy_v2":             110,
		"huaweicloud_vpnaas_service_v2":                  111,
		"huaweicloud_vpnaas_site_connection_v2":          112,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
