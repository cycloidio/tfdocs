package flexibleengine

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_antiddos_v1",
			Category:         "Anti-DDoS",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "Security-Anti-DDoS.svg",
			Keywords: []string{
				"anti",
				"ddos",
				"antiddos",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "floating_ip_id",
					Description: `(Required) The ID corresponding to the Elastic IP Address (EIP).`,
				},
				resource.Attribute{
					Name:        "enable_l7",
					Description: `(Required) Specifies whether to enable L7 defense.`,
				},
				resource.Attribute{
					Name:        "traffic_pos_id",
					Description: `(Required) The position ID of traffic. The value ranges from 1 to 9.`,
				},
				resource.Attribute{
					Name:        "http_request_pos_id",
					Description: `(Required) The position ID of number of HTTP requests. The value ranges from 1 to 15.`,
				},
				resource.Attribute{
					Name:        "cleaning_access_pos_id",
					Description: `(Required)The position ID of access limit during cleaning. The value ranges from 1 to 8.`,
				},
				resource.Attribute{
					Name:        "app_type_id",
					Description: `(Required) The application type ID. ## Attributes Reference All above argument parameters can be exported as attribute parameters. ## Import Antiddos can be imported using the floating_ip_id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_antiddos_v1.myantiddos c1881895-cdcb-4d23-96cb-032e6a3ee667 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_api_gateway_api",
			Category:         "API Gateway",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"api",
				"gateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String, ForceNew) The region in which to create the API resource. If omitted, the provider-level region will be used. Changing this creates a new API resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, String) Specifies the name of the API. An API name consists of 3–64 characters, starting with a letter. Only letters, digits, and underscores (_) are allowed.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required, String, ForceNew) Specifies the ID of the API group. Changing this creates a new resource.`,
				},
				resource.Attribute{
					Name:        "visibility",
					Description: `(Required, Int) Specifies whether the API is available to the public. The value can only be 1 (public).`,
				},
				resource.Attribute{
					Name:        "auth_type",
					Description: `(Required, String) Specifies the security authentication mode. The value can be 'APP', 'IAM', and ' NONE'.`,
				},
				resource.Attribute{
					Name:        "request_protocol",
					Description: `(Optional, String) Specifies the request protocol. The value can be 'HTTP', 'HTTPS', and 'BOTH' which means the API can be accessed through both 'HTTP' and 'HTTPS'. Defaults to 'HTTPS'.`,
				},
				resource.Attribute{
					Name:        "request_method",
					Description: `(Required, String) Specifies the request method, including 'GET','POST','PUT' and etc..`,
				},
				resource.Attribute{
					Name:        "request_uri",
					Description: `(Required, String) Specifies the request path of the API. The value must comply with URI specifications.`,
				},
				resource.Attribute{
					Name:        "backend_type",
					Description: `(Required, String) Specifies the service backend type. The value can be: + 'HTTP': the web service backend + 'FUNCTION': the FunctionGraph service backend + 'MOCK': the Mock service backend`,
				},
				resource.Attribute{
					Name:        "http_backend",
					Description: `(Optional, List) Specifies the configuration when backend_type selected 'HTTP' (documented below).`,
				},
				resource.Attribute{
					Name:        "function_backend",
					Description: `(Optional, List) Specifies the configuration when backend_type selected 'FUNCTION' (documented below).`,
				},
				resource.Attribute{
					Name:        "mock_backend",
					Description: `(Optional, List) Specifies the configuration when backend_type selected 'MOCK' (documented below).`,
				},
				resource.Attribute{
					Name:        "request_parameter",
					Description: `(Optional, List) the request parameter list (documented below).`,
				},
				resource.Attribute{
					Name:        "backend_parameter",
					Description: `(Optional, List) the backend parameter list (documented below).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, List) the tags of API in format of string list.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, String) Specifies the description of the API. The description cannot exceed 255 characters.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional, String) Specifies the version of the API. A maximum of 16 characters are allowed.`,
				},
				resource.Attribute{
					Name:        "cors",
					Description: `(Optional, Bool) Specifies whether CORS is supported or not.`,
				},
				resource.Attribute{
					Name:        "example_success_response",
					Description: `(Required, String) Specifies the example response for a successful request. The length cannot exceed 20,480 characters.`,
				},
				resource.Attribute{
					Name:        "example_failure_response",
					Description: `(Optional, String) Specifies the example response for a failed request The length cannot exceed 20,480 characters. The ` + "`" + `http_backend` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required, String) Specifies the backend request protocol. The value can be 'HTTP' and 'HTTPS'.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Optional, String) Specifies the backend request method, including 'GET','POST','PUT' and etc..`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `(Required, String) Specifies the backend request path. The value must comply with URI specifications.`,
				},
				resource.Attribute{
					Name:        "vpc_channel",
					Description: `(Optional, String) Specifies the VPC channel ID. This parameter and ` + "`" + `url_domain` + "`" + ` are alternative.`,
				},
				resource.Attribute{
					Name:        "url_domain",
					Description: `(Optional, String) Specifies the backend service address. An endpoint URL is in the format of "domain name (or IP address):port number", with up to 255 characters. This parameter and ` + "`" + `vpc_channel` + "`" + ` are alternative.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional, Int) Timeout duration (in ms) for API Gateway to request for the backend service. Defaults to 50000. The ` + "`" + `function_backend` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "function_urn",
					Description: `(Required, String) Specifies the function URN.`,
				},
				resource.Attribute{
					Name:        "invocation_type",
					Description: `(Required, String) Specifies the invocation mode, which can be 'async' or 'sync'.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional, String) Specifies the function version.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional, Int) Timeout duration (in ms) for API Gateway to request for FunctionGraph. Defaults to 50000. The ` + "`" + `mock_backend` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "result_content",
					Description: `(Optional, String) Specifies the return result.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional, String) Specifies the version of the Mock backend.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, String) Specifies the description of the Mock backend. The description cannot exceed 255 characters. The ` + "`" + `request_parameter` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, String) Specifies the input parameter name. A parameter name consists of 1–32 characters, starting with a letter. Only letters, digits, periods (.), hyphens (-), and underscores (_) are allowed.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Required, String) Specifies the input parameter location, which can be 'PATH', 'QUERY' or 'HEADER'.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required, String) Specifies the input parameter type, which can be 'STRING' or 'NUMBER'.`,
				},
				resource.Attribute{
					Name:        "required",
					Description: `(Optional, Bool) Specifies whether the parameter is mandatory or not.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `(Optional, String) Specifies the default value when the parameter is optional.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, String) Specifies the description of the parameter. The description cannot exceed 255 characters. The ` + "`" + `backend_parameter` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, String) Specifies the parameter name. A parameter name consists of 1–32 characters, starting with a letter. Only letters, digits, periods (.), hyphens (-), and underscores (_) are allowed.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Required, String) Specifies the parameter location, which can be 'PATH', 'QUERY' or 'HEADER'.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required, String) Specifies the parameter type, which can be 'REQUEST', 'CONSTANT', or 'SYSTEM'.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required, String) Specifies the parameter value, which is a string of not more than 255 characters. The value varies depending on the parameter type: + 'REQUEST': parameter name in ` + "`" + `request_parameter` + "`" + ` + 'CONSTANT': real value of the parameter + 'SYSTEM': gateway parameter name`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, String) Specifies the description of the parameter. The description cannot exceed 255 characters. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the API.`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `The name of the API group to which the API belongs. ## Timeouts This resource provides the following timeouts configuration options:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `Default is 10 minute.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `Default is 10 minute. ## Import API can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_api_gateway_api.api "774438a28a574ac8a496325d1bf51807" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the API.`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `The name of the API group to which the API belongs. ## Timeouts This resource provides the following timeouts configuration options:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `Default is 10 minute.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `Default is 10 minute. ## Import API can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_api_gateway_api.api "774438a28a574ac8a496325d1bf51807" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_api_gateway_group",
			Category:         "API Gateway",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"api",
				"gateway",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String, ForceNew) The region in which to create the API gateway group resource. If omitted, the provider-level region will be used. Changing this creates a new gateway group resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, String) Specifies the name of the API group. An API group name consists of 3–64 characters, starting with a letter. Only letters, digits, and underscores (_) are allowed.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, String) Specifies the description of the API group. The description cannot exceed 255 characters. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the API group.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the API group. ## Import API groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_api_gateway_group.apigw_group "c8738f7c-a4b0-4c5f-a202-bda7dc4018a4" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the API group.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the API group. ## Import API groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_api_gateway_group.apigw_group "c8738f7c-a4b0-4c5f-a202-bda7dc4018a4" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_as_configuration_v1",
			Category:         "Auto Scaling (AS)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "Computing-AS.svg",
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
					Description: `(Required) The disk size. The unit is GB. The system disk size ranges from 1 to 32768, and the data disk size ranges from 10 to 32768.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `(Required) The disk type, which must be the same as the disk type available in the system. The options include ` + "`" + `SATA` + "`" + ` (common I/O disk type) and ` + "`" + `SSD` + "`" + ` (ultra-high I/O disk type).`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `(Required) Whether the disk is a system disk or a data disk. Option ` + "`" + `DATA` + "`" + ` indicates a data disk. option ` + "`" + `SYS` + "`" + ` indicates a system disk. The ` + "`" + `personality` + "`" + ` block supports:`,
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
			Type:             "flexibleengine_as_group_v1",
			Category:         "Auto Scaling (AS)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "Computing-AS.svg",
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
					Description: `(Optional) The cooling duration (in seconds). The value ranges from 0 to 86400, and is 900 by default.`,
				},
				resource.Attribute{
					Name:        "lb_listener_id",
					Description: `(Optional) The ELB listener IDs. The system supports up to six ELB listeners, the IDs of which are separated using a comma (,).`,
				},
				resource.Attribute{
					Name:        "lbaas_listeners",
					Description: `(Optional) An array of one or more enhanced load balancer. The system supports the binding of up to six load balancers. The field is alternative to ` + "`" + `lb_listener_id` + "`" + `. The object structure is documented below.`,
				},
				resource.Attribute{
					Name:        "available_zones",
					Description: `(Optional) The availability zones in which to create the instances in the autoscaling group.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) The VPC ID. Changing this creates a new group.`,
				},
				resource.Attribute{
					Name:        "networks",
					Description: `(Required) An array of one or more network IDs. The system supports up to five networks. The networks object structure is documented below.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Required) An array of ` + "`" + `one` + "`" + ` security group ID to associate with the group. The security_groups object structure is documented below.`,
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
					Description: `(Optional, Map) The key/value pairs to associate with the scaling group.`,
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
					Description: `(Optional) Whether to delete the instances in the AS group when deleting the AS group. The options are ` + "`" + `yes` + "`" + ` and ` + "`" + `no` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "force_delete",
					Description: `(Optional) Whether to forcibly delete the AS group, remove the ECS instances and release them. The default value is ` + "`" + `false` + "`" + `. The ` + "`" + `networks` + "`" + ` block supports:`,
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
					Description: `(Optional) Specifies the weight, which determines the portion of requests a backend ECS processes compared to other backend ECSs added to the same listener. The value of this parameter ranges from 0 to 100. The default value is 1. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates the status of the AS group.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `The instances IDs of the AS group.`,
				},
				resource.Attribute{
					Name:        "current_instance_number",
					Description: `Indicates the number of current instances in the AS group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `Indicates the status of the AS group.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `The instances IDs of the AS group.`,
				},
				resource.Attribute{
					Name:        "current_instance_number",
					Description: `Indicates the number of current instances in the AS group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_as_lifecycle_hook_v1",
			Category:         "Auto Scaling (AS)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"auto",
				"scaling",
				"as",
				"lifecycle",
				"hook",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required, String) Specifies the lifecycle hook name. This parameter can contain a maximum of 32 characters, which may consist of letters, digits, underscores (_) and hyphens (-).`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `(Required, String, ForceNew) Specifies the ID of the AS group in UUID format. Changing this creates a new AS lifecycle hook.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required, String) Specifies the lifecycle hook type. The valid values are following strings:`,
				},
				resource.Attribute{
					Name:        "notification_topic_urn",
					Description: `(Required, String) Specifies a unique topic in SMN.`,
				},
				resource.Attribute{
					Name:        "default_result",
					Description: `(Optional, String) Specifies the default lifecycle hook callback operation. This operation is performed when the timeout duration expires. The valid values are`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional, Int) Specifies the lifecycle hook timeout duration, which ranges from 300 to 86400 in the unit of second, default to 3600.`,
				},
				resource.Attribute{
					Name:        "notification_message",
					Description: `(Optional, String) Specifies a customized notification. This parameter can contains a maximum of 256 characters, which cannot contain the following characters: <>&'(). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID.`,
				},
				resource.Attribute{
					Name:        "notification_topic_name",
					Description: `The topic name in SMN.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The server time in UTC format when the lifecycle hook is created. ## Import Lifecycle hooks can be imported using the AS group ID and hook ID separated by a slash, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_as_lifecycle_hook_v1.test <AS group ID>/<Lifecycle hook ID> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID.`,
				},
				resource.Attribute{
					Name:        "notification_topic_name",
					Description: `The topic name in SMN.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The server time in UTC format when the lifecycle hook is created. ## Import Lifecycle hooks can be imported using the AS group ID and hook ID separated by a slash, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_as_lifecycle_hook_v1.test <AS group ID>/<Lifecycle hook ID> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_as_policy_v1",
			Category:         "Auto Scaling (AS)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "Computing-AS.svg",
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
			Type:             "flexibleengine_blockstorage_volume_v2",
			Category:         "Elastic Volume Service (EVS)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "Storage.svg",
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
					Description: `(Optional) Metadata key/value pairs to associate with the volume. Changing this updates the existing volume metadata. The EVS encryption capability with KMS key can be set with the following parameters:`,
				},
				resource.Attribute{
					Name:        "__system__encrypted",
					Description: `The default value is set to '0', which means the volume is not encrypted, the value '1' indicates volume is encrypted.`,
				},
				resource.Attribute{
					Name:        "__system__cmkid",
					Description: `(Optional) The ID of the kms key.`,
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
					Description: `(Optional) The type of volume to create. Changing this creates a new volume.`,
				},
				resource.Attribute{
					Name:        "cascade",
					Description: `(Optional, Default:false) Specifies to delete all snapshots associated with the EVS disk.`,
				},
				resource.Attribute{
					Name:        "multiattach",
					Description: `(Optional) Specifies whether the EVS disk is shareable.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value pairs to associate with the volume. ## Attributes Reference The following attributes are exported:`,
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
					Name:        "multiattach",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "attachment",
					Description: `If a volume is attached to an instance, this attribute will display the Attachment ID, Instance ID, and the Device as the Instance sees it. ## Import Volumes can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_blockstorage_volume_v2.volume_1 ea257959-eeb1-4c10-8d33-26f0409a755d ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "multiattach",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "attachment",
					Description: `If a volume is attached to an instance, this attribute will display the Attachment ID, Instance ID, and the Device as the Instance sees it. ## Import Volumes can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_blockstorage_volume_v2.volume_1 ea257959-eeb1-4c10-8d33-26f0409a755d ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_cbr_policy",
			Category:         "Cloud Backup and Recovery (CBR)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"backup",
				"and",
				"recovery",
				"cbr",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String, ForceNew) Specifies the region in which to create the CBR policy. If omitted, the provider-level region will be used. Changing this will create a new policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, String) Specifies a unique name of the CBR policy. This parameter can contain a maximum of 64 characters, which may consist of chinese charactors, letters, digits, underscores(_) and hyphens (-).`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required, String, ForceNew) Specifies the protection type of the CBR policy. Valid values are`,
				},
				resource.Attribute{
					Name:        "backup_cycle",
					Description: `(Required, List) Specifies the scheduling rule for the CBR policy backup execution. The [object](#cbr_policy_backup_cycle) structure is documented below.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional, Bool) Specifies whether to enable the CBR policy. Default to`,
				},
				resource.Attribute{
					Name:        "destination_region",
					Description: `(Optional, String) Specifies the name of the replication destination region, which is mandatory for cross-region replication. Required if ` + "`" + `protection_type` + "`" + ` is`,
				},
				resource.Attribute{
					Name:        "destination_project_id",
					Description: `(Optional, String) Specifies the ID of the replication destination project, which is mandatory for cross-region replication. Required if ` + "`" + `protection_type` + "`" + ` is`,
				},
				resource.Attribute{
					Name:        "backup_quantity",
					Description: `(Optional, Int) Specifies the maximum number of retained backups. The value ranges from ` + "`" + `2` + "`" + ` to ` + "`" + `99,999` + "`" + `. This parameter and ` + "`" + `time_period` + "`" + ` are alternative.`,
				},
				resource.Attribute{
					Name:        "time_period",
					Description: `(Optional, Int) Specifies the duration (in days) for retained backups. The value ranges from ` + "`" + `2` + "`" + ` to ` + "`" + `99,999` + "`" + `. ->`,
				},
				resource.Attribute{
					Name:        "long_term_retention",
					Description: `(Optional, List) Specifies the long-term retention rules, which is an advanced options of the ` + "`" + `backup_quantity` + "`" + `. The [object](#cbr_policy_long_term_retention) structure is documented below. -> The configuration of ` + "`" + `long_term_retention` + "`" + ` and ` + "`" + `backup_quantity` + "`" + ` will take effect together. When the number of retained backups exceeds the preset value (number of ` + "`" + `backup_quantity` + "`" + `), the system automatically deletes the earliest backups. By default, the system automatically clears data every other day.`,
				},
				resource.Attribute{
					Name:        "time_zone",
					Description: `(Optional, String) Specifies the UTC time zone, e.g.: ` + "`" + `UTC+08:00` + "`" + `. Required if ` + "`" + `long_term_retention` + "`" + ` is set. <a name="cbr_policy_backup_cycle"></a> The ` + "`" + `backup_cycle` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "days",
					Description: `(Optional, String) Specifies the weekly backup day of backup schedule. It supports seven days a week (MO, TU, WE, TH, FR, SA, SU) and this parameter is separated by a comma (,) without spaces, between date and date during the configuration.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional, Int) Specifies the interval (in days) of backup schedule. The value range is ` + "`" + `1` + "`" + ` to ` + "`" + `30` + "`" + `. This parameter and ` + "`" + `days` + "`" + ` are alternative.`,
				},
				resource.Attribute{
					Name:        "execution_times",
					Description: `(Required, List) Specifies the backup time. Automated backups will be triggered at the backup time. The current time is in the UTC format (HH:MM). The minutes in the list must be set to`,
				},
				resource.Attribute{
					Name:        "daily",
					Description: `(Optional, Int) - Specifies the latest backup of each day is saved in the long term.`,
				},
				resource.Attribute{
					Name:        "weekly",
					Description: `(Optional, Int) - Specifies the latest backup of each week is saved in the long term.`,
				},
				resource.Attribute{
					Name:        "monthly",
					Description: `(Optional, Int) - Specifies the latest backup of each month is saved in the long term.`,
				},
				resource.Attribute{
					Name:        "yearly",
					Description: `(Optional, Int) - Specifies the latest backup of each year is saved in the long term. -> A maximum of 10 backups are retained for failed periodic backup tasks. They are retained for one month and can be manually deleted on the web console. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A resource ID in UUID format. ## Import Policies can be imported by their ` + "`" + `id` + "`" + `. For example, ` + "`" + `` + "`" + `` + "`" + ` terraform import flexibleengine_cbr_policy.test 4d2c2939-774f-42ef-ab15-e5b126b11ace ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `A resource ID in UUID format. ## Import Policies can be imported by their ` + "`" + `id` + "`" + `. For example, ` + "`" + `` + "`" + `` + "`" + ` terraform import flexibleengine_cbr_policy.test 4d2c2939-774f-42ef-ab15-e5b126b11ace ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_cbr_vault",
			Category:         "Cloud Backup and Recovery (CBR)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"backup",
				"and",
				"recovery",
				"cbr",
				"vault",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `A resource ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "allocated",
					Description: `The allocated capacity of the vault, in GB.`,
				},
				resource.Attribute{
					Name:        "used",
					Description: `The used capacity, in GB.`,
				},
				resource.Attribute{
					Name:        "spec_code",
					Description: `The specification code.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The vault status.`,
				},
				resource.Attribute{
					Name:        "storage",
					Description: `The name of the bucket for the vault. ## Import Vaults can be imported by their ` + "`" + `id` + "`" + `. For example, ` + "`" + `` + "`" + `` + "`" + ` terraform import flexibleengine_cbr_vault.test 01c33779-7c83-4182-8b6b-24a671fcedf8 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_cce_addon_v3",
			Category:         "Cloud Container Engine (CCE)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"container",
				"engine",
				"cce",
				"addon",
				"v3",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required, String, ForceNew) ID of the cluster. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required, String, ForceNew) Name of the addon template. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required, String, ForceNew) Version of the addon. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Optional, List, ForceNew) Add-on template installation parameters. These parameters vary depending on the add-on. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "basic",
					Description: `(Required, String, ForceNew) The basic parameters in json string format. Changing this will create a new resource.`,
				},
				resource.Attribute{
					Name:        "custom",
					Description: `(Optional, String, ForceNew) The custom parameters in json string format. Changing this will create a new resource.`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `(Optional, String, ForceNew) The flavor parameters in json string format. Changing this will create a new resource. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the addon instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Addon status information.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of addon instance. ## Timeouts This resource provides the following timeouts configuration options:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `Default is 10 minute.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `Default is 3 minute. ## Import CCE addon can be imported using the cluster ID and addon ID separated by a slash, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_cce_addon_v3.my_addon bb6923e4-b16e-11eb-b0cd-0255ac101da1/c7ecb230-b16f-11eb-b3b6-0255ac1015a3 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the addon instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Addon status information.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of addon instance. ## Timeouts This resource provides the following timeouts configuration options:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `Default is 10 minute.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `Default is 3 minute. ## Import CCE addon can be imported using the cluster ID and addon ID separated by a slash, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_cce_addon_v3.my_addon bb6923e4-b16e-11eb-b0cd-0255ac101da1/c7ecb230-b16f-11eb-b3b6-0255ac1015a3 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_cce_cluster_v3",
			Category:         "Cloud Container Engine (CCE)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "Computing-CCE.svg",
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
					Description: `(Required) Cluster specifications. Changing this parameter will create a new cluster resource.`,
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
					Description: `(Optional) For the cluster version, possible values are listed on the [CCE Cluster Version Release Notes](https://docs.prod-cloud-ocb.orange-business.com/usermanual2/cce/cce_01_0068.html). If this parameter is not set, the latest available version will be used.`,
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
					Description: `(Required) The NETWORK ID of the subnet used to create the node. Changing this parameter will create a new cluster resource.`,
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
					Name:        "service_network_cidr",
					Description: `(Optional) Service network segment. Changing this parameter will create a new cluster resource.`,
				},
				resource.Attribute{
					Name:        "authentication_mode",
					Description: `(Optional) Authentication mode of the cluster, possible values are x509 and rbac. Defaults to`,
				},
				resource.Attribute{
					Name:        "authenticating_proxy_ca",
					Description: `(Optional) CA root certificate provided in the authenticating_proxy mode. The CA root certificate is encoded to the Base64 format. Changing this parameter will create a new cluster resource.`,
				},
				resource.Attribute{
					Name:        "eip",
					Description: `(Optional) EIP address of the cluster. Changing this parameter will create a new cluster resource.`,
				},
				resource.Attribute{
					Name:        "kube_proxy_mode",
					Description: `(Optional, String, ForceNew) Service forwarding mode. Two modes are available: - iptables: Traditional kube-proxy uses iptables rules to implement service load balancing. In this mode, too many iptables rules will be generated when many services are deployed. In addition, non-incremental updates will cause a latency and even obvious performance issues in the case of heavy service traffic. - ipvs: Optimized kube-proxy mode with higher throughput and faster speed. This mode supports incremental updates and can keep connections uninterrupted during service updates. It is suitable for large-sized clusters.`,
				},
				resource.Attribute{
					Name:        "masters",
					Description: `(Optional, List, ForceNew) Advanced configuration of master nodes. Changing this creates a new cluster. The ` + "`" + `masters` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional, String, ForceNew) Specifies the availability zone of the master node. Changing this creates a new cluster. ## Attributes Reference All above argument parameters can be exported as attribute parameters along with attribute reference.`,
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
					Name:        "internal_endpoint",
					Description: `The internal network address.`,
				},
				resource.Attribute{
					Name:        "external_endpoint",
					Description: `The external network address.`,
				},
				resource.Attribute{
					Name:        "external_apig_endpoint",
					Description: `The endpoint of the cluster to be accessed through API Gateway.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `Security group ID of the cluster.`,
				},
				resource.Attribute{
					Name:        "certificate_clusters.name",
					Description: `The cluster name.`,
				},
				resource.Attribute{
					Name:        "certificate_clusters.server",
					Description: `The server IP address.`,
				},
				resource.Attribute{
					Name:        "certificate_clusters.certificate_authority_data",
					Description: `The certificate data.`,
				},
				resource.Attribute{
					Name:        "certificate_users.name",
					Description: `The user name.`,
				},
				resource.Attribute{
					Name:        "certificate_users.client_certificate_data",
					Description: `The client certificate data.`,
				},
				resource.Attribute{
					Name:        "certificate_users.client_key_data",
					Description: `The client key data. ## Import Cluster can be imported using the cluster id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_cce_cluster_v3.cluster_1 4779ab1c-7c1a-44b1-a02e-93dfc361b32d ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "internal_endpoint",
					Description: `The internal network address.`,
				},
				resource.Attribute{
					Name:        "external_endpoint",
					Description: `The external network address.`,
				},
				resource.Attribute{
					Name:        "external_apig_endpoint",
					Description: `The endpoint of the cluster to be accessed through API Gateway.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `Security group ID of the cluster.`,
				},
				resource.Attribute{
					Name:        "certificate_clusters.name",
					Description: `The cluster name.`,
				},
				resource.Attribute{
					Name:        "certificate_clusters.server",
					Description: `The server IP address.`,
				},
				resource.Attribute{
					Name:        "certificate_clusters.certificate_authority_data",
					Description: `The certificate data.`,
				},
				resource.Attribute{
					Name:        "certificate_users.name",
					Description: `The user name.`,
				},
				resource.Attribute{
					Name:        "certificate_users.client_certificate_data",
					Description: `The client certificate data.`,
				},
				resource.Attribute{
					Name:        "certificate_users.client_key_data",
					Description: `The client key data. ## Import Cluster can be imported using the cluster id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_cce_cluster_v3.cluster_1 4779ab1c-7c1a-44b1-a02e-93dfc361b32d ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_cce_namespace",
			Category:         "Cloud Container Engine (CCE)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"container",
				"engine",
				"cce",
				"namespace",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String, ForceNew) Specifies the region in which to create the namespace resource. If omitted, the provider-level region will be used. Changing this will create a new namespace resource.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required, String, ForceNew) Specifies the cluster ID to which the CCE namespace belongs. Changing this will create a new namespace resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, String, ForceNew) Specifies the unique name of the namespace. This parameter can contain a maximum of 63 characters, which may consist of lowercase letters, digits and hyphens (-), and must start and end with lowercase letters and digits. Changing this will create a new namespace resource. Exactly one of ` + "`" + `name` + "`" + ` or ` + "`" + `prefix` + "`" + ` must be provided.`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Optional, String, ForceNew) Specifies a prefix used by the server to generate a unique name. This parameter can contain a maximum of 63 characters, which may consist of lowercase letters, digits and hyphens (-), and must start and end with lowercase letters and digits. Changing this will create a new namespace resource. Exactly one of ` + "`" + `name` + "`" + ` or ` + "`" + `prefix` + "`" + ` must be provided.`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional, Map, ForceNew) Specifies an unstructured key value map for external parameters. Changing this will create a new namespace resource.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional, Map, ForceNew) Specifies the map of string keys and values for labels. Changing this will create a new namespace resource. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The namespace ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `The server time when namespace was created.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current phase of the namespace. ## Timeouts This resource provides the following timeouts configuration options:`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `Default is 5 minute. ## Import CCE namespace can be imported using the cluster ID and namespace name separated by a slash, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_cce_namespace.test bb6923e4-b16e-11eb-b0cd-0255ac101da1/test-namespace ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The namespace ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `The server time when namespace was created.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current phase of the namespace. ## Timeouts This resource provides the following timeouts configuration options:`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `Default is 5 minute. ## Import CCE namespace can be imported using the cluster ID and namespace name separated by a slash, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_cce_namespace.test bb6923e4-b16e-11eb-b0cd-0255ac101da1/test-namespace ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_cce_node_pool_v3",
			Category:         "Cloud Container Engine (CCE)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"container",
				"engine",
				"cce",
				"node",
				"pool",
				"v3",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required, String, ForceNew) ID of the cluster. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, String) Node Pool Name.`,
				},
				resource.Attribute{
					Name:        "initial_node_count",
					Description: `(Required, Int) Initial number of expected nodes in the node pool.`,
				},
				resource.Attribute{
					Name:        "flavor_id",
					Description: `(Required, String, ForceNew) Specifies the flavor id. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional, String, ForceNew) Node Pool type. Possible values are: "vm" and "ElasticBMS".`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional, String, ForceNew) specify the name of the available partition (AZ). Default value is random to create nodes in a random AZ in the node pool. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "os",
					Description: `(Optional, String) Operating System of the node. The value can be EulerOS 2.5 and CentOS 7.6. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "key_pair",
					Description: `(Optional, String, ForceNew) Key pair name when logging in to select the key pair mode. This parameter and ` + "`" + `password` + "`" + ` are alternative. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional, String, ForceNew) root password when logging in to select the password mode. This parameter must be`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional, String, ForceNew) The ID of the subnet to which the NIC belongs. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "max_pods",
					Description: `(Optional, Int, ForceNew) The maximum number of instances a node is allowed to create. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "preinstall",
					Description: `(Optional, String, ForceNew) Script required before installation. The input value can be a Base64 encoded string or not. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "postinstall",
					Description: `(Optional, String, ForceNew) Script required after the installation. The input value can be a Base64 encoded string or not. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "scall_enable",
					Description: `(Optional, Bool) Whether to enable auto scaling. If Autoscaler is enabled, install the autoscaler add-on to use the auto scaling feature.`,
				},
				resource.Attribute{
					Name:        "min_node_count",
					Description: `(Optional, Int) Minimum number of nodes allowed if auto scaling is enabled.`,
				},
				resource.Attribute{
					Name:        "max_node_count",
					Description: `(Optional, Int) Maximum number of nodes allowed if auto scaling is enabled.`,
				},
				resource.Attribute{
					Name:        "scale_down_cooldown_time",
					Description: `(Optional, Int) Interval between two scaling operations, in minutes.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional, Int) Weight of a node pool. A node pool with a higher weight has a higher priority during scaling.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional, Map) Tags of a Kubernetes node, key/value pair format.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Map) Tags of a VM node, key/value pair format.`,
				},
				resource.Attribute{
					Name:        "root_volume",
					Description: `(Required, List, ForceNew) It corresponds to the system disk related configuration. The object structure is documented below. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "data_volumes",
					Description: `(Required, List, ForceNew) Represents the data disk to be created. The object structure is documented below. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "taints",
					Description: `(Optional, List) You can add taints to created nodes to configure anti-affinity. The object structure is documented below. The ` + "`" + `root_volume` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required, Int) Specifies the disk size in GB.`,
				},
				resource.Attribute{
					Name:        "volumetype",
					Description: `(Required, String) Specifies the disk type.`,
				},
				resource.Attribute{
					Name:        "extend_params",
					Description: `(Optional, Map) Specifies the disk expansion parameters in key/value pair format. The ` + "`" + `data_volumes` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required, Int) Specifies the disk size in GB.`,
				},
				resource.Attribute{
					Name:        "volumetype",
					Description: `(Required, String) Specifies the disk type.`,
				},
				resource.Attribute{
					Name:        "kms_key_id",
					Description: `(Optional, String) Specifies the KMS key ID. This is used to encrypt the volume.`,
				},
				resource.Attribute{
					Name:        "extend_params",
					Description: `(Optional, Map) Specifies the disk expansion parameters in key/value pair format. The ` + "`" + `taints` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required, String) A key must contain 1 to 63 characters starting with a letter or digit. Only letters, digits, hyphens (-), underscores (_), and periods (.) are allowed. A DNS subdomain name can be used as the prefix of a key.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required, String) A value must start with a letter or digit and can contain a maximum of 63 characters, including letters, digits, hyphens (-), underscores (_), and periods (.).`,
				},
				resource.Attribute{
					Name:        "effect",
					Description: `(Required, String) Available options are`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Specifies a resource ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Node status information.`,
				},
				resource.Attribute{
					Name:        "billing_mode",
					Description: `Billing mode of a node. ## Timeouts This resource provides the following timeouts configuration options:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `Default is 20 minute.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `Default is 20 minute. ## Import Node_pool can be imported using the cluster and node_pool id, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import flexibleengine_cce_node_pool_v3.node_pool_1 <cluster-id>/<node_pool-id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Specifies a resource ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Node status information.`,
				},
				resource.Attribute{
					Name:        "billing_mode",
					Description: `Billing mode of a node. ## Timeouts This resource provides the following timeouts configuration options:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `Default is 20 minute.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `Default is 20 minute. ## Import Node_pool can be imported using the cluster and node_pool id, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import flexibleengine_cce_node_pool_v3.node_pool_1 <cluster-id>/<node_pool-id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_cce_node_v3",
			Category:         "Cloud Container Engine (CCE)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"container",
				"engine",
				"cce",
				"node",
				"v3",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) ID of the cluster. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Node Name.`,
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
					Name:        "key_pair",
					Description: `(Required) Key pair name when logging in to select the key pair mode. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "os",
					Description: `(Optional) Operating System of the node, possible values are EulerOS 2.2 and CentOS 7.6. Defaults to EulerOS 2.2. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Tags of a Kubernetes node, key/value pair format. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) VM tag, key/value pair format.`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional) Node annotation, key/value pair format. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "eip_ids",
					Description: `(Optional) List of existing elastic IP IDs. Changing this parameter will create a new resource. -> If the ` + "`" + `eip_ids` + "`" + ` parameter is configured, you do not need to configure the ` + "`" + `eip_count` + "`" + ` and bandwidth parameters: ` + "`" + `iptype` + "`" + `, ` + "`" + `bandwidth_charge_mode` + "`" + `, ` + "`" + `bandwidth_size` + "`" + ` and ` + "`" + `share_type` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "eip_count",
					Description: `(Optional) Number of elastic IPs to be dynamically created. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "iptype",
					Description: `(Optional) Elastic IP type.`,
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
					Name:        "ecs_performance_type",
					Description: `(Optional) Classification of cloud server specifications. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "product_id",
					Description: `(Optional) The Product ID. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "max_pods",
					Description: `(Optional) The maximum number of instances a node is allowed to create. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(Optional) The Public key. Changing this parameter will create a new resource.`,
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
					Name:        "extend_param",
					Description: `(Optional, Map, ForceNew) Extended parameter. Changing this parameter will create a new resource. Availiable keys: + ` + "`" + `agency_name` + "`" + ` - Specifies the agency name to provide temporary credentials for CCE node to access other cloud services. + ` + "`" + `dockerBaseSize` + "`" + ` - The available disk space of a single docker container on the node in device mapper mode. + ` + "`" + `DockerLVMConfigOverride` + "`" + ` - Docker data disk configurations. The following is an example default configuration: ` + "`" + `` + "`" + `` + "`" + `hcl extend_param = { DockerLVMConfigOverride = "dockerThinpool=vgpaas/90%VG;kubernetesLV=vgpaas/10%VG;diskType=evs;lvType=linear" } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) Specifies the disk size in GB.`,
				},
				resource.Attribute{
					Name:        "volumetype",
					Description: `(Required) Specifies the disk type.`,
				},
				resource.Attribute{
					Name:        "extend_params",
					Description: `(Optional) Specifies the disk expansion parameters in key/value pair format.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) Specifies the disk size in GB.`,
				},
				resource.Attribute{
					Name:        "volumetype",
					Description: `(Required) Specifies the disk type.`,
				},
				resource.Attribute{
					Name:        "extend_params",
					Description: `(Optional) Specifies the disk expansion parameters in key/value pair format.`,
				},
				resource.Attribute{
					Name:        "kms_key_id",
					Description: `(Optional) Specifies the ID of a KMS key. This is used to encrypt the volume. -> You need to create an agency (EVSAccessKMS) when disk encryption is used in the current project for the first time ever. The account and permission of the created agency are ` + "`" + `op_svc_evs` + "`" + ` and`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A key must contain 1 to 63 characters starting with a letter or digit. Only letters, digits, hyphens (-), underscores (_), and periods (.) are allowed. A DNS subdomain name can be used as the prefix of a key.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) A value must start with a letter or digit and can contain a maximum of 63 characters, including letters, digits, hyphens (-), underscores (_), and periods (.).`,
				},
				resource.Attribute{
					Name:        "effect",
					Description: `(Required) Available options are NoSchedule, PreferNoSchedule, and NoExecute. ## Attributes Reference All above argument parameters can be exported as attribute parameters along with attribute reference.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Node status information.`,
				},
				resource.Attribute{
					Name:        "server_id",
					Description: `ID of the ECS instance associated with the node.`,
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
					Name:        "id",
					Description: `The resource ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Node status information.`,
				},
				resource.Attribute{
					Name:        "server_id",
					Description: `ID of the ECS instance associated with the node.`,
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
			Type:             "flexibleengine_cce_pvc",
			Category:         "Cloud Container Engine (CCE)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"container",
				"engine",
				"cce",
				"pvc",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String, ForceNew) Specifies the region in which to create the PVC resource. If omitted, the provider-level region will be used. Changing this will create a new PVC resource.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required, String, ForceNew) Specifies the cluster ID to which the CCE PVC belongs.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Required, String, ForceNew) Specifies the namespace to logically divide your containers into different group. Changing this will create a new PVC resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, String, ForceNew) Specifies the unique name of the PVC resource. This parameter can contain a maximum of 63 characters, which may consist of lowercase letters, digits and hyphens (-), and must start and end with lowercase letters and digits. Changing this will create a new PVC resource.`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional, Map, ForceNew) Specifies the unstructured key value map for external parameters. Changing this will create a new PVC resource.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional, Map, ForceNew) Specifies the map of string keys and values for labels. Changing this will create a new PVC resource.`,
				},
				resource.Attribute{
					Name:        "storage_class_name",
					Description: `(Required, String, ForceNew) Specifies the type of the storage bound to the CCE PVC. The valid values are as follows: +`,
				},
				resource.Attribute{
					Name:        "access_modes",
					Description: `(Required, List, ForceNew) Specifies the desired access modes the volume should have. The valid values are as follows: +`,
				},
				resource.Attribute{
					Name:        "storage",
					Description: `(Required, String, ForceNew) Specifies the minimum amount of storage resources required. Changing this creates a new PVC resource. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The PVC ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `The server time when PVC was created.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current phase of the PVC. +`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `Default is 5 minute.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `Default is 3 minute. ## Import CCE PVC can be imported using the cluster ID, namespace and name separated by a slash, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_cce_pvc.test 5c20fdad-7288-11eb-b817-0255ac10158b/default/pvc_name ` + "`" + `` + "`" + `` + "`" + ` Note that the imported state may not be identical to your resource definition, due to some attributes missing from the API response, security or some other reason. The missing attributes include: ` + "`" + `annotations` + "`" + `. It is generally recommended running ` + "`" + `terraform plan` + "`" + ` after importing a PVC. You can then decide if changes should be applied to the PVC, or the resource definition should be updated to align with the PVC. Also you can ignore changes as below. ` + "`" + `` + "`" + `` + "`" + ` resource "flexibleengine_cce_pvc" "test" { ... lifecycle { ignore_changes = [ annotations, ] } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The PVC ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `The server time when PVC was created.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current phase of the PVC. +`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `Default is 5 minute.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `Default is 3 minute. ## Import CCE PVC can be imported using the cluster ID, namespace and name separated by a slash, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_cce_pvc.test 5c20fdad-7288-11eb-b817-0255ac10158b/default/pvc_name ` + "`" + `` + "`" + `` + "`" + ` Note that the imported state may not be identical to your resource definition, due to some attributes missing from the API response, security or some other reason. The missing attributes include: ` + "`" + `annotations` + "`" + `. It is generally recommended running ` + "`" + `terraform plan` + "`" + ` after importing a PVC. You can then decide if changes should be applied to the PVC, or the resource definition should be updated to align with the PVC. Also you can ignore changes as below. ` + "`" + `` + "`" + `` + "`" + ` resource "flexibleengine_cce_pvc" "test" { ... lifecycle { ignore_changes = [ annotations, ] } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_ces_alarmrule",
			Category:         "Cloud Eye (CES)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"eye",
				"ces",
				"alarmrule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "alarm_name",
					Description: `(Required, String) Specifies the name of an alarm rule. The value can be a string of 1 to 128 characters that can consist of numbers, lowercase letters, uppercase letters, underscores (_), or hyphens (-).`,
				},
				resource.Attribute{
					Name:        "metric",
					Description: `(Required, List) Specifies the alarm metrics. The structure is described below.`,
				},
				resource.Attribute{
					Name:        "condition",
					Description: `(Required, List) Specifies the alarm triggering condition. The structure is described below.`,
				},
				resource.Attribute{
					Name:        "alarm_description",
					Description: `(Optional, String) The value can be a string of 0 to 256 characters.`,
				},
				resource.Attribute{
					Name:        "alarm_enabled",
					Description: `(Optional, Bool) Specifies whether to enable the alarm. The default value is true.`,
				},
				resource.Attribute{
					Name:        "alarm_level",
					Description: `(Optional, Int, ForceNew) Specifies the alarm severity. The value can be 1, 2, 3 or 4, which indicates`,
				},
				resource.Attribute{
					Name:        "alarm_actions",
					Description: `(Optional, List) Specifies the action triggered by an alarm. The structure is described below.`,
				},
				resource.Attribute{
					Name:        "ok_actions",
					Description: `(Optional, List) Specifies the action triggered by the clearing of an alarm. The structure is described below.`,
				},
				resource.Attribute{
					Name:        "alarm_action_enabled",
					Description: `(Optional, Bool) Specifies whether to enable the action to be triggered by an alarm. The default value is true. ->`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Required, String) Specifies the namespace in`,
				},
				resource.Attribute{
					Name:        "metric_name",
					Description: `(Required, String) Specifies the metric name. The value can be a string of 1 to 64 characters that must start with a letter and can consists of uppercase letters, lowercase letters, numbers, or underscores (_).`,
				},
				resource.Attribute{
					Name:        "dimensions",
					Description: `(Required, List) Specifies the list of metric dimensions. Currently, the maximum length of the dimesion list that are supported is 3. The structure is described below. The ` + "`" + `dimensions` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, String) Specifies the dimension name. The value can be a string of 1 to 32 characters that must start with a letter and can consists of uppercase letters, lowercase letters, numbers, underscores (_), or hyphens (-).`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required, String) Specifies the dimension value. The value can be a string of 1 to 64 characters that must start with a letter or a number and can consists of uppercase letters, lowercase letters, numbers, underscores (_), or hyphens (-). The ` + "`" + `condition` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Required, Int) Specifies the alarm checking period in seconds. The value can be 1, 300, 1200, 3600, 14400, and 86400. Note: If period is set to 1, the raw metric data is used to determine whether to generate an alarm.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Required, String) Specifies the data rollup methods. The value can be max, min, average, sum, and vaiance.`,
				},
				resource.Attribute{
					Name:        "comparison_operator",
					Description: `(Required, String) Specifies the comparison condition of alarm thresholds. The value can be >, =, <, >=, or <=.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required, String) Specifies the alarm threshold. The value ranges from 0 to Number of 1.7976931348623157e+308.`,
				},
				resource.Attribute{
					Name:        "count",
					Description: `(Required, Int) Specifies the number of consecutive occurrence times. The value ranges from 1 to 5.`,
				},
				resource.Attribute{
					Name:        "unit",
					Description: `(Optional, String) Specifies the data unit. the ` + "`" + `alarm_actions` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional, String) specifies the type of action triggered by an alarm. the value can be`,
				},
				resource.Attribute{
					Name:        "notification_list",
					Description: `(Optional, List) specifies the list of objects to be notified if the alarm status changes, the maximum length is 5. if type is set to`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional, String) specifies the type of action triggered by an alarm. the value is notification. notification: indicates that a notification will be sent to the user.`,
				},
				resource.Attribute{
					Name:        "notification_list",
					Description: `(Optional, List) specifies the list of objects to be notified if the alarm status changes, the maximum length is 5. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Indicates the alarm rule ID.`,
				},
				resource.Attribute{
					Name:        "alarm_state",
					Description: `Indicates the alarm status. The value can be: - ok: The alarm status is normal; - alarm: An alarm is generated; - insufficient_data: The required data is insufficient.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Indicates the time when the alarm status changed. The value is a UNIX timestamp and the unit is ms. ## Import CES alarm rules can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_ces_alarmrule.alarm_rule al1619678242900OxEaaODM2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Indicates the alarm rule ID.`,
				},
				resource.Attribute{
					Name:        "alarm_state",
					Description: `Indicates the alarm status. The value can be: - ok: The alarm status is normal; - alarm: An alarm is generated; - insufficient_data: The required data is insufficient.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Indicates the time when the alarm status changed. The value is a UNIX timestamp and the unit is ms. ## Import CES alarm rules can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_ces_alarmrule.alarm_rule al1619678242900OxEaaODM2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_compute_bms_server_v2",
			Category:         "Bare Metal Server (BMS)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"bare",
				"metal",
				"server",
				"bms",
				"compute",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to create the bms server instance. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new bms server.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the BMS.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Optional; Required if ` + "`" + `image_name` + "`" + ` is empty.) Changing this creates a new bms server.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `(Optional; Required if ` + "`" + `image_id` + "`" + ` is empty.) The name of the desired image for the bms server. Changing this creates a new bms server.`,
				},
				resource.Attribute{
					Name:        "flavor_id",
					Description: `(Optional; Required if ` + "`" + `flavor_name` + "`" + ` is empty) The flavor ID of the desired flavor for the bms server. Changing this resizes the existing bms server.`,
				},
				resource.Attribute{
					Name:        "flavor_name",
					Description: `(Optional; Required if ` + "`" + `flavor_id` + "`" + ` is empty) The name of the desired flavor for the bms server. Changing this resizes the existing bms server.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional) The user data to provide when launching the instance. Changing this creates a new bms server.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Optional) An array of one or more security group names to associate with the bms server. Changing this results in adding/removing security groups from the existing bms server.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required) The availability zone in which to create the bms server.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Optional) An array of one or more networks to attach to the bms instance. Changing this creates a new bms server.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) Metadata key/value pairs to make available from within the instance. Changing this updates the existing bms server metadata.`,
				},
				resource.Attribute{
					Name:        "admin_pass",
					Description: `(Optional) The administrative password to assign to the bms server. Changing this changes the root password on the existing server.`,
				},
				resource.Attribute{
					Name:        "key_pair",
					Description: `(Optional) The name of a key pair to put on the bms server. The key pair must already be created and associated with the tenant's account. Changing this creates a new bms server.`,
				},
				resource.Attribute{
					Name:        "stop_before_destroy",
					Description: `(Optional) Whether to try stop instance gracefully before destroying it, thus giving chance for guest OS daemons to stop correctly. If instance doesn't stop within timeout, it will be destroyed anyway. The ` + "`" + `network` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `(Required unless ` + "`" + `port` + "`" + ` or ` + "`" + `name` + "`" + ` is provided) The network UUID to attach to the bms server. Changing this creates a new bms server.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required unless ` + "`" + `uuid` + "`" + ` or ` + "`" + `port` + "`" + ` is provided) The human-readable name of the network. Changing this creates a new bms server.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required unless ` + "`" + `uuid` + "`" + ` or ` + "`" + `name` + "`" + ` is provided) The port UUID of a network to attach to the bms server. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "fixed_ip_v4",
					Description: `(Optional) Specifies a fixed IPv4 address to be used on this network. Changing this creates a new bms server.`,
				},
				resource.Attribute{
					Name:        "fixed_ip_v6",
					Description: `(Optional) Specifies a fixed IPv6 address to be used on this network. Changing this creates a new bms server.`,
				},
				resource.Attribute{
					Name:        "access_network",
					Description: `(Optional) Specifies if this network should be used for provisioning access. Accepts true or false. Defaults to false. # Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the bms server.`,
				},
				resource.Attribute{
					Name:        "config_drive",
					Description: `Whether to use the config_drive feature to configure the instance.`,
				},
				resource.Attribute{
					Name:        "kernel_id",
					Description: `The UUID of the kernel image when the AMI image is used.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `The ID of the user to which the BMS belongs.`,
				},
				resource.Attribute{
					Name:        "host_status",
					Description: `The nova-compute status:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_compute_floatingip_associate_v2",
			Category:         "Elastic Cloud Server (ECS)",
			ShortDescription: ``,
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
					Description: `See Argument Reference above. ## Import This resource can be imported by specifying all three arguments, separated by a forward slash: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_compute_floatingip_associate_v2.fip_1 <floating_ip>/<instance_id>/<fixed_ip> ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `See Argument Reference above. ## Import This resource can be imported by specifying all three arguments, separated by a forward slash: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_compute_floatingip_associate_v2.fip_1 <floating_ip>/<instance_id>/<fixed_ip> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_compute_floatingip_v2",
			Category:         "Elastic Cloud Server (ECS)",
			ShortDescription: ``,
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
					Description: `(Optional) The name of the pool from which to obtain the floating IP. Default value is admin_external_net. Changing this creates a new floating IP. ## Attributes Reference The following attributes are exported:`,
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
					Description: `UUID of the compute instance associated with the floating IP. ## Import Floating IPs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_compute_floatingip_v2.floatip_1 89c60255-9bd6-460c-822a-e2b959ede9d2 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `UUID of the compute instance associated with the floating IP. ## Import Floating IPs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_compute_floatingip_v2.floatip_1 89c60255-9bd6-460c-822a-e2b959ede9d2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_compute_instance_v2",
			Category:         "Elastic Cloud Server (ECS)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "Computing-ECS.svg",
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
					Name:        "availability_zone",
					Description: `(Optional) The availability zone in which to create the server. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Optional) An array of one or more security group names to associate with the server. Changing this results in adding/removing security groups from the existing server.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Optional) An array of one or more networks to attach to the instance. The network object structure is documented below. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional) The user data to provide when launching the instance. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) The key/value pairs to associate with the instance. Changing this updates the existing server metadata.`,
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
					Description: `(Optional) Whether to try stop instance gracefully before destroying it, thus giving chance for guest OS daemons to stop correctly. If instance doesn't stop within timeout, it will be destroyed anyway.`,
				},
				resource.Attribute{
					Name:        "force_delete",
					Description: `(Optional) Whether to force the FlexibleEngine instance to be forcefully deleted. This is useful for environments that have reclaim / soft deletion enabled.`,
				},
				resource.Attribute{
					Name:        "auto_recovery",
					Description: `(Optional) Configures or deletes automatic recovery of an instance`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Specifies the key/value pairs to associate with the instance. The ` + "`" + `network` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `(Required unless ` + "`" + `port` + "`" + ` is provided) The network UUID to attach to the server. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required unless ` + "`" + `uuid` + "`" + ` is provided) The port UUID of a network to attach to the server. Changing this creates a new server.`,
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
					Name:        "volume_type",
					Description: `(Optional) Currently, the value can be ` + "`" + `SSD` + "`" + ` (ultra-I/O disk type), ` + "`" + `SAS` + "`" + ` (high I/O disk type), or ` + "`" + `SATA` + "`" + ` (common I/O disk type)`,
				},
				resource.Attribute{
					Name:        "boot_index",
					Description: `(Optional) The boot index of the volume. It defaults to 0, which indicates that it's a system disk. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "destination_type",
					Description: `(Optional) The type that gets created. Possible values are "volume" and "local". Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "delete_on_termination",
					Description: `(Optional) Delete the volume / block device upon termination of the instance. Defaults to false. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "disk_bus",
					Description: `(Optional) The low-level disk bus that will be used, for example,`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `(Optional) Specifies the`,
				},
				resource.Attribute{
					Name:        "tenancy",
					Description: `(Optional) Specifies whether the ECS is created on a Dedicated Host (DeH) or in a shared pool (default). The value can be`,
				},
				resource.Attribute{
					Name:        "deh_id",
					Description: `(Optional) Specifies the DeH ID. This parameter takes effect only when the value of tenancy is dedicated. If you do not specify this parameter, the system will automatically assign a DeH to you to deploy ECSs. ## Attributes Reference The following attributes are exported:`,
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
					Description: `The first detected Fixed IPv4 address`,
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
					Name:        "network/port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/name",
					Description: `The human-readable name of the network.`,
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
					Name:        "all_metadata",
					Description: `Contains all instance metadata, even metadata not set by Terraform.`,
				},
				resource.Attribute{
					Name:        "auto_recovery",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "floating_ip",
					Description: `The EIP address that is associted to the instance.`,
				},
				resource.Attribute{
					Name:        "system_disk_id",
					Description: `The system disk voume ID.`,
				},
				resource.Attribute{
					Name:        "volume_attached",
					Description: `An array of one or more disks to attach to the instance. The volume_attached object structure is documented below.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the instance. The ` + "`" + `volume_attached` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `The volume id on that attachment.`,
				},
				resource.Attribute{
					Name:        "boot_index",
					Description: `The volume boot index on that attachment.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The volume size on that attachment.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The volume type on that attachment.`,
				},
				resource.Attribute{
					Name:        "pci_address",
					Description: `The volume pci address on that attachment. ## Import Instances can be imported by their ` + "`" + `id` + "`" + `. For example, ` + "`" + `` + "`" + `` + "`" + ` terraform import flexibleengine_compute_instance_v2.my_instance b11b407c-e604-4e8d-8bc4-92398320b847 ` + "`" + `` + "`" + `` + "`" + ` Note that the imported state may not be identical to your resource definition, due to some attrubutes missing from the API response, security or some other reason. The missing attributes include: ` + "`" + `admin_pass` + "`" + `, ` + "`" + `config_drive` + "`" + `, ` + "`" + `user_data` + "`" + `, ` + "`" + `block_device` + "`" + `, ` + "`" + `scheduler_hints` + "`" + `, ` + "`" + `stop_before_destroy` + "`" + `, ` + "`" + `network/access_network` + "`" + ` and arguments for pre-paid. It is generally recommended running ` + "`" + `terraform plan` + "`" + ` after importing an instance. You can then decide if changes should be applied to the instance, or the resource definition should be updated to align with the instance. Also you can ignore changes as below. ` + "`" + `` + "`" + `` + "`" + ` resource "flexibleengine_compute_instance_v2" "my_instance" { ... lifecycle { ignore_changes = [ user_data, block_device, ] } } ` + "`" + `` + "`" + `` + "`" + ` ## Notes ### Multiple Ephemeral Disks It's possible to specify multiple ` + "`" + `block_device` + "`" + ` entries to create an instance with multiple ephemeral (local) disks. In order to create multiple ephemeral disks, the sum of the total amount of ephemeral space must be less than or equal to what the chosen flavor supports. The following example shows how to create an instance with multiple ephemeral disks: ` + "`" + `` + "`" + `` + "`" + ` resource "flexibleengine_compute_instance_v2" "foo" { name = "terraform-test" security_groups = ["default"] block_device { boot_index = 0 delete_on_termination = true destination_type = "local" source_type = "image" uuid = "<image uuid>" } block_device { boot_index = -1 delete_on_termination = true destination_type = "local" source_type = "blank" volume_size = 1 } block_device { boot_index = -1 delete_on_termination = true destination_type = "local" source_type = "blank" volume_size = 1 } } ` + "`" + `` + "`" + `` + "`" + ` ### Instances and Ports Neutron Ports are a great feature and provide a lot of functionality. However, there are some notes to be aware of when mixing Instances and Ports:`,
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
					Description: `The first detected Fixed IPv4 address`,
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
					Name:        "network/port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/name",
					Description: `The human-readable name of the network.`,
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
					Name:        "all_metadata",
					Description: `Contains all instance metadata, even metadata not set by Terraform.`,
				},
				resource.Attribute{
					Name:        "auto_recovery",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "floating_ip",
					Description: `The EIP address that is associted to the instance.`,
				},
				resource.Attribute{
					Name:        "system_disk_id",
					Description: `The system disk voume ID.`,
				},
				resource.Attribute{
					Name:        "volume_attached",
					Description: `An array of one or more disks to attach to the instance. The volume_attached object structure is documented below.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the instance. The ` + "`" + `volume_attached` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `The volume id on that attachment.`,
				},
				resource.Attribute{
					Name:        "boot_index",
					Description: `The volume boot index on that attachment.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The volume size on that attachment.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The volume type on that attachment.`,
				},
				resource.Attribute{
					Name:        "pci_address",
					Description: `The volume pci address on that attachment. ## Import Instances can be imported by their ` + "`" + `id` + "`" + `. For example, ` + "`" + `` + "`" + `` + "`" + ` terraform import flexibleengine_compute_instance_v2.my_instance b11b407c-e604-4e8d-8bc4-92398320b847 ` + "`" + `` + "`" + `` + "`" + ` Note that the imported state may not be identical to your resource definition, due to some attrubutes missing from the API response, security or some other reason. The missing attributes include: ` + "`" + `admin_pass` + "`" + `, ` + "`" + `config_drive` + "`" + `, ` + "`" + `user_data` + "`" + `, ` + "`" + `block_device` + "`" + `, ` + "`" + `scheduler_hints` + "`" + `, ` + "`" + `stop_before_destroy` + "`" + `, ` + "`" + `network/access_network` + "`" + ` and arguments for pre-paid. It is generally recommended running ` + "`" + `terraform plan` + "`" + ` after importing an instance. You can then decide if changes should be applied to the instance, or the resource definition should be updated to align with the instance. Also you can ignore changes as below. ` + "`" + `` + "`" + `` + "`" + ` resource "flexibleengine_compute_instance_v2" "my_instance" { ... lifecycle { ignore_changes = [ user_data, block_device, ] } } ` + "`" + `` + "`" + `` + "`" + ` ## Notes ### Multiple Ephemeral Disks It's possible to specify multiple ` + "`" + `block_device` + "`" + ` entries to create an instance with multiple ephemeral (local) disks. In order to create multiple ephemeral disks, the sum of the total amount of ephemeral space must be less than or equal to what the chosen flavor supports. The following example shows how to create an instance with multiple ephemeral disks: ` + "`" + `` + "`" + `` + "`" + ` resource "flexibleengine_compute_instance_v2" "foo" { name = "terraform-test" security_groups = ["default"] block_device { boot_index = 0 delete_on_termination = true destination_type = "local" source_type = "image" uuid = "<image uuid>" } block_device { boot_index = -1 delete_on_termination = true destination_type = "local" source_type = "blank" volume_size = 1 } block_device { boot_index = -1 delete_on_termination = true destination_type = "local" source_type = "blank" volume_size = 1 } } ` + "`" + `` + "`" + `` + "`" + ` ### Instances and Ports Neutron Ports are a great feature and provide a lot of functionality. However, there are some notes to be aware of when mixing Instances and Ports:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_compute_interface_attach_v2",
			Category:         "Elastic Cloud Server (ECS)",
			ShortDescription: ``,
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
			Type:             "flexibleengine_compute_keypair_v2",
			Category:         "Elastic Cloud Server (ECS)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "key-pair.svg",
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
					Description: `(Optional, String, ForceNew) Specifies the region in which to create the keypair resource. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new keypair.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, String, ForceNew) Specifies a unique name for the keypair. Changing this creates a new keypair.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(Optional, String, ForceNew) Specifies a imported OpenSSH-formatted public key. Changing this creates a new keypair.`,
				},
				resource.Attribute{
					Name:        "private_key_path",
					Description: `(Optional, String, ForceNew) Specifies the path of the created private key. The private key file (`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID which as same as keypair name. ## Import Keypairs can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_compute_keypair_v2.my-keypair test-keypair ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID which as same as keypair name. ## Import Keypairs can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_compute_keypair_v2.my-keypair test-keypair ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_compute_servergroup_v2",
			Category:         "Elastic Cloud Server (ECS)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "cloud-server-group.svg",
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
					Description: `(Required) The set of policies for the server group. Only`,
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
					Name:        "policies",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `The instances that are part of this server group. ## Import Server Groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_compute_servergroup_v2.test-sg 1bc30ee9-9d5b-4c30-bdd5-7f1e663f5edf ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `The instances that are part of this server group. ## Import Server Groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_compute_servergroup_v2.test-sg 1bc30ee9-9d5b-4c30-bdd5-7f1e663f5edf ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_compute_volume_attach_v2",
			Category:         "Elastic Cloud Server (ECS)",
			ShortDescription: ``,
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
					Description: `See Argument Reference above. _NOTE_: The correctness of this information is dependent upon the hypervisor in use. In some cases, this should not be used as an authoritative piece of information. ## Import Volume Attachments can be imported using the Instance ID and Volume ID separated by a slash, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_compute_volume_attach_v2.va_1 89c60255-9bd6-460c-822a-e2b959ede9d2/45670584-225f-46c3-b33e-6707b589b666 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `See Argument Reference above. _NOTE_: The correctness of this information is dependent upon the hypervisor in use. In some cases, this should not be used as an authoritative piece of information. ## Import Volume Attachments can be imported using the Instance ID and Volume ID separated by a slash, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_compute_volume_attach_v2.va_1 89c60255-9bd6-460c-822a-e2b959ede9d2/45670584-225f-46c3-b33e-6707b589b666 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_csbs_backup_policy_v1",
			Category:         "Cloud Server Backup Service (CSBS)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "Storage-CSBS.svg",
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
					Description: `Specifies Scheduler type. ## Import Backup Policy can be imported using ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_csbs_backup_policy_v1.backup_policy_v1 7056d636-ac60-4663-8a6c-82d3c32c1c64 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `Specifies Scheduler type. ## Import Backup Policy can be imported using ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_csbs_backup_policy_v1.backup_policy_v1 7056d636-ac60-4663-8a6c-82d3c32c1c64 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_csbs_backup_v1",
			Category:         "Cloud Server Backup Service (CSBS)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "Storage-CSBS.svg",
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
					Description: `Specifies image type. ## Import Backup can be imported using ` + "`" + `backup_record_id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_csbs_backup_v1.backup_v1.backup_v1 7056d636-ac60-4663-8a6c-82d3c32c1c64 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `Specifies image type. ## Import Backup can be imported using ` + "`" + `backup_record_id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_csbs_backup_v1.backup_v1.backup_v1 7056d636-ac60-4663-8a6c-82d3c32c1c64 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_css_cluster_v1",
			Category:         "Cloud Search Service (CSS)",
			ShortDescription: ``,
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
					Name:        "name",
					Description: `(Required) Cluster name. It contains 4 to 32 characters. Only letters, digits, hyphens (-), and underscores (_) are allowed. The value must start with a letter. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "engine_type",
					Description: `(Optional) Engine type. The default value is "elasticsearch". Currently, the value can only be "elasticsearch". Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `(Required) Engine version. Versions 6.5.4 and 7.1.1 are supported. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "node_number",
					Description: `(Optional) Number of cluster instances. The value range is 1 to 32. Defaults to 1.`,
				},
				resource.Attribute{
					Name:        "security_mode",
					Description: `(Optional) Whether to enable communication encryption and security authentication. Available values include`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Password of the cluster administrator admin in security mode. This parameter is mandatory only when security_mode is set to true. Changing this parameter will create a new resource. The administrator password must meet the following requirements: - The password can contain 8 to 32 characters. - The password must contain at least 3 of the following character types: uppercase letters, lowercase letters, digits, and special characters (~!@#$%^&`,
				},
				resource.Attribute{
					Name:        "node_config",
					Description: `(Required) Node configuration. Structure is documented below. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "backup_strategy",
					Description: `(Optional) Specifies the advanced backup policy. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value pairs to associate with the cluster. The ` + "`" + `node_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) Availability zone(s). You can set multiple vailability zones, and use commas (,) to separate one from another. Cluster instances will be evenly distributed to each AZ. The ` + "`" + `node_number` + "`" + ` should be greater than or equal to the number of available zones. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `(Required) Instance flavor name. For example: value range of flavor ess.spec-2u8g: 40 GB to 800 GB, value range of flavor ess.spec-4u16g: 40 GB to 1600 GB, value range of flavor ess.spec-8u32g: 80 GB to 3200 GB, value range of flavor ess.spec-16u64g: 100 GB to 6400 GB, value range of flavor ess.spec-32u128g: 100 GB to 10240 GB. Changing this parameter will create a new resource.`,
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
					Name:        "vpc_id",
					Description: `(Required) VPC ID, which is used for configuring cluster network. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Required) Security group ID. All instances in a cluster must have the same security group. Changing this parameter will create a new resource. The ` + "`" + `volume` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) Specifies volume size in GB, which must be a multiple of 10.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `(Required) Specifies the volume type. Changing this parameter will create a new resource. Supported value: - "COMMON": The SATA disk is used; - "HIGH": The SAS disk is used; - "ULTRAHIGH": The solid-state drive (SSD) is used. The ` + "`" + `backup_strategy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `(Required) Specifies the time when a snapshot is automatically created everyday. Snapshots can only be created on the hour. The time format is the time followed by the time zone, specifically,`,
				},
				resource.Attribute{
					Name:        "keep_days",
					Description: `(Optional) Specifies the number of days to retain the generated snapshots. Snapshots are reserved for seven days by default.`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Optional) Specifies the prefix of the snapshot that is automatically created. The default value is "snapshot". ->`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `Indicates the IP address and port number.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `Time when a cluster is created. The format is ISO8601: CCYY-MM-DDThh:mm:ss.`,
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
					Description: `Supported type: ess (indicating the Elasticsearch node). ## Timeouts This resource provides the following timeouts configuration options: - ` + "`" + `create` + "`" + ` - Default is 60 minute. - ` + "`" + `update` + "`" + ` - Default is 60 minute.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "endpoint",
					Description: `Indicates the IP address and port number.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `Time when a cluster is created. The format is ISO8601: CCYY-MM-DDThh:mm:ss.`,
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
					Description: `Supported type: ess (indicating the Elasticsearch node). ## Timeouts This resource provides the following timeouts configuration options: - ` + "`" + `create` + "`" + ` - Default is 60 minute. - ` + "`" + `update` + "`" + ` - Default is 60 minute.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_css_snapshot_v1",
			Category:         "Cloud Search Service (CSS)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"search",
				"service",
				"css",
				"snapshot",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the snapshot name. The snapshot name must start with a letter and contains 4 to 64 characters consisting of only lowercase letters, digits, hyphens (-), and underscores (_). Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) Specifies ID of the CSS cluster where index data is to be backed up. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "indices",
					Description: `(Optional) Specifies the name of the index to be backed up. Multiple index names are separated by commas (,). By default, data of all indices is backed up. You can use the asterisk (`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Specifies the description of a snapshot. The value contains 0 to 256 characters, and angle brackets (<) and (>) are not allowed. Changing this parameter will create a new resource. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates the snapshot status.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `Indicates the CSS cluster name.`,
				},
				resource.Attribute{
					Name:        "backup_type",
					Description: `Indicates the snapshot creation mode, the value should be "manual" or "automated". ## Import This resource can be imported by specifying the CSS cluster ID and snapshot ID separated by a slash, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_css_snapshot_v1.snapshot_1 <cluster_id>/<snapshot_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `Indicates the snapshot status.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `Indicates the CSS cluster name.`,
				},
				resource.Attribute{
					Name:        "backup_type",
					Description: `Indicates the snapshot creation mode, the value should be "manual" or "automated". ## Import This resource can be imported by specifying the CSS cluster ID and snapshot ID separated by a slash, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_css_snapshot_v1.snapshot_1 <cluster_id>/<snapshot_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_cts_tracker_v1",
			Category:         "Cloud Trace Service (CTS)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "Management&Deployment-CTS.svg",
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
					Name:        "status",
					Description: `The status of a tracker. The value should be`,
				},
				resource.Attribute{
					Name:        "tracker_name",
					Description: `The tracker name. Currently, only tracker`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "tracker_name",
					Description: `The tracker name. Currently, only tracker`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_dcs_instance_v1",
			Category:         "Distributed Cache Service (DCS)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "Database-DCS.svg",
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
					Description: `(Required) Indicates a cache engine. Only Redis is supported. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `(Required) Indicates the version of a cache engine, which is 3.0.7. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `(Required) Indicates the Cache capacity. Unit: GB. For a DCS Redis or Memcached instance in single-node or master/standby mode, the cache capacity can be 2 GB, 4 GB, 8 GB, 16 GB, 32 GB, or 64 GB. For a DCS Redis instance in cluster mode, the cache capacity can be 64, 128, 256, 512, or 1024 GB. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "access_user",
					Description: `(Optional) Username used for accessing a DCS instance after password authentication. A username starts with a letter, consists of 1 to 64 characters, and supports only letters, digits, and hyphens (-). Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Password of a DCS instance. The password of a DCS Redis instance must meet the following complexity requirements: Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) Specifies the id of the VPC. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required) Specifies the network id of the subnet. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Optional) Specifies the id of the security group which the instance belongs to. This parameter is mandatory for Memcached and Redis 3.0 versions.`,
				},
				resource.Attribute{
					Name:        "available_zones",
					Description: `(Required) IDs or Names of the AZs where cache nodes reside. For details on how to query AZs, see Querying AZ Information. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "product_id",
					Description: `(Optional) Product ID used to differentiate DCS instance types. For example now there are following values available: - dcs.memcached.master_standby-h, - dcs.memcached.master_standby-m, - dcs.memcached.single_node-h, - dcs.memcached.single_node-m, - dcs.master_standby-h, - dcs.cluster-h, - dcs.single_node-h, - redis.cluster.xu1.large.r1.4-h, - redis.cluster.xu1.large.r2.4-h, - redis.cluster.xu1.large.r1.8-h, - redis.cluster.xu1.large.r2.16-h, - redis.cluster.xu1.large.r1.16-h, - redis.cluster.xu1.large.r2.24-h, - redis.cluster.xu1.large.r2.32-h, - redis.cluster.xu1.large.r1.32-h, - redis.cluster.xu1.large.r2.48-h, - redis.cluster.xu1.large.r1.48-h - redis.cluster.xu1.large.r2.64-h - redis.cluster.xu1.large.r1.64-h - redis.cluster.xu1.large.r2.96-h - redis.cluster.xu1.large.r1.96-h - redis.cluster.xu1.large.r2.128-h - redis.cluster.xu1.large.r1.128-h - redis.cluster.xu1.large.r1.192-h - redis.cluster.xu1.large.r2.192-h - redis.cluster.xu1.large.r2.256-h - redis.cluster.xu1.large.r1.256-h - redis.cluster.xu1.large.r2.384-h - redis.cluster.xu1.large.r1.384-h - redis.cluster.xu1.large.r1.512-h - redis.cluster.xu1.large.r2.512-h ..... For the whole list and the specification of product id please check the DCS API DOC for querying: https://dcs.eu-west-0.prod-cloud-ocb.orange-business.com/v1.0/products Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Deprecated, Optional) DCS instance specification code. For example now there are following values available: - dcs.single_node - dcs.master_standby - dcs.cluster`,
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
					Description: `(Optional) Day in a week on which backup starts. Range: 1–7. Where: 1 indicates Monday; 7 indicates Sunday. Changing this creates a new instance. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the Cache instance.`,
				},
				resource.Attribute{
					Name:        "vpc_name",
					Description: `Indicates the name of a vpc.`,
				},
				resource.Attribute{
					Name:        "subnet_name",
					Description: `Indicates the name of a subnet.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `Indicates the name of a security group.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Cache node's IP address in tenant's VPC.`,
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
					Name:        "internal_version",
					Description: `Internal DCS version.`,
				},
				resource.Attribute{
					Name:        "max_memory",
					Description: `Overall memory size. Unit: MB.`,
				},
				resource.Attribute{
					Name:        "used_memory",
					Description: `Size of the used memory. Unit: MB.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `Indicates a user ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the Cache instance.`,
				},
				resource.Attribute{
					Name:        "vpc_name",
					Description: `Indicates the name of a vpc.`,
				},
				resource.Attribute{
					Name:        "subnet_name",
					Description: `Indicates the name of a subnet.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `Indicates the name of a security group.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Cache node's IP address in tenant's VPC.`,
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
					Name:        "internal_version",
					Description: `Internal DCS version.`,
				},
				resource.Attribute{
					Name:        "max_memory",
					Description: `Overall memory size. Unit: MB.`,
				},
				resource.Attribute{
					Name:        "used_memory",
					Description: `Size of the used memory. Unit: MB.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `Indicates a user ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_dds_instance_v3",
			Category:         "Document Database Service (DDS)",
			ShortDescription: ``,
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
					Description: `(Optional) Specifies the advanced backup policy. The structure is described below. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "ssl",
					Description: `(Optional) Specifies whether to enable or disable SSL. Defaults to true. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value pairs to associate with the DDS instance. The ` + "`" + `datastore` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Specifies the DB engine. Only DDS-Community is supported now.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) Specifies the DB instance version. Only 3.4 and 4.0 are supported now.`,
				},
				resource.Attribute{
					Name:        "storage_engine",
					Description: `(Optional) Specifies the storage engine of the DB instance. Only wiredTiger is supported now. The ` + "`" + `flavor` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Specifies the node type. Valid value: mongos, shard, config, replica.`,
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
					Description: `(Required) Specifies the resource specification code. Valid values: engine_name | type | vcpus | ram | speccode ---- | --- | --- DDS-Community | mongos | 1 | 4 | dds.mongodb.s3.medium.4.mongos DDS-Community | mongos | 2 | 8 | dds.mongodb.s3.large.4.mongos DDS-Community | mongos | 4 | 16 | dds.mongodb.s3.xlarge.4.mongos DDS-Community | mongos | 8 | 32 | dds.mongodb.s3.2xlarge.4.mongos DDS-Community | mongos | 16 | 64 | dds.mongodb.s3.4xlarge.4.mongos DDS-Community | shard | 1 | 4 | dds.mongodb.s3.medium.4.shard DDS-Community | shard | 2 | 8 | dds.mongodb.s3.large.4.shard DDS-Community | shard | 4 | 16 | dds.mongodb.s3.xlarge.4.shard DDS-Community | shard | 8 | 32 | dds.mongodb.s3.2xlarge.4.shard DDS-Community | shard | 16 | 64 | dds.mongodb.s3.4xlarge.4.shard DDS-Community | config | 2 | 4 | dds.mongodb.s3.large.2.config DDS-Community | replica | 1 | 4 | dds.mongodb.s3.medium.4.repset DDS-Community | replica | 2 | 8 | dds.mongodb.s3.large.4.repset DDS-Community | replica | 4 | 16 | dds.mongodb.s3.xlarge.4.repset DDS-Community | replica | 8 | 32 | dds.mongodb.s3.2xlarge.4.repset DDS-Community | replica | 16 | 64 | dds.mongodb.s3.4xlarge.4.repset The ` + "`" + `backup_strategy ` + "`" + ` block supports:`,
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
					Name:        "id",
					Description: `The resource ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates the the DB instance status.`,
				},
				resource.Attribute{
					Name:        "db_username",
					Description: `Indicates the DB Administator name.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Indicates the database port number. The port range is 2100 to 9500.`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `Indicates the instance nodes information. Structure is documented below. The ` + "`" + `nodes` + "`" + ` block contains: - ` + "`" + `id` + "`" + ` - Indicates the node ID. - ` + "`" + `name` + "`" + ` - Indicates the node name. - ` + "`" + `role` + "`" + ` - Indicates the node role. - ` + "`" + `type` + "`" + ` - Indicates the node type. - ` + "`" + `private_ip` + "`" + ` - Indicates the private IP address of a node. This parameter is valid only for mongos nodes, replica set instances, and single node instances. - ` + "`" + `public_ip` + "`" + ` - Indicates the EIP that has been bound on a node. This parameter is valid only for mongos nodes of cluster instances, primary nodes and secondary nodes of replica set instances, and single node instances. - ` + "`" + `status` + "`" + ` - Indicates the node status.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates the the DB instance status.`,
				},
				resource.Attribute{
					Name:        "db_username",
					Description: `Indicates the DB Administator name.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Indicates the database port number. The port range is 2100 to 9500.`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `Indicates the instance nodes information. Structure is documented below. The ` + "`" + `nodes` + "`" + ` block contains: - ` + "`" + `id` + "`" + ` - Indicates the node ID. - ` + "`" + `name` + "`" + ` - Indicates the node name. - ` + "`" + `role` + "`" + ` - Indicates the node role. - ` + "`" + `type` + "`" + ` - Indicates the node type. - ` + "`" + `private_ip` + "`" + ` - Indicates the private IP address of a node. This parameter is valid only for mongos nodes, replica set instances, and single node instances. - ` + "`" + `public_ip` + "`" + ` - Indicates the EIP that has been bound on a node. This parameter is valid only for mongos nodes of cluster instances, primary nodes and secondary nodes of replica set instances, and single node instances. - ` + "`" + `status` + "`" + ` - Indicates the node status.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_dis_stream",
			Category:         "Data Ingestion Service (DIS)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"data",
				"ingestion",
				"service",
				"dis",
				"stream",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required, String, ForceNew) Specifies the name of the DIS stream to be created. Changing this will create a new resource.`,
				},
				resource.Attribute{
					Name:        "partition_count",
					Description: `(Required, Int, ForceNew) Specifies the number of the expect partitions. Changing this will create a new resource.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional, String, ForceNew) Specifies the Stream type. The value can be`,
				},
				resource.Attribute{
					Name:        "retention_period",
					Description: `(Optional, Int, ForceNew) Specifies the number of hours for which data from the stream will be retained in DIS. The value ranges from 24 to 168 and defaults to 24. Changing this will create a new resource. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID which equals to stream name.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of stream: ` + "`" + `CREATING` + "`" + `,` + "`" + `RUNNING` + "`" + `,` + "`" + `TERMINATING` + "`" + `,` + "`" + `TERMINATED` + "`" + `,` + "`" + `FROZEN` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "partitions",
					Description: `The information of stream partitions. Structure is documented below. The ` + "`" + `partitions` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the partition.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the partition.`,
				},
				resource.Attribute{
					Name:        "hash_range",
					Description: `Possible value range of the hash key used by each partition.`,
				},
				resource.Attribute{
					Name:        "sequence_number_range",
					Description: `Sequence number range of each partition. ## Import Dis stream can be imported by ` + "`" + `name` + "`" + `. For example, ` + "`" + `` + "`" + `` + "`" + ` terraform import flexibleengine_dis_stream.example dis-demo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID which equals to stream name.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of stream: ` + "`" + `CREATING` + "`" + `,` + "`" + `RUNNING` + "`" + `,` + "`" + `TERMINATING` + "`" + `,` + "`" + `TERMINATED` + "`" + `,` + "`" + `FROZEN` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "partitions",
					Description: `The information of stream partitions. Structure is documented below. The ` + "`" + `partitions` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the partition.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the partition.`,
				},
				resource.Attribute{
					Name:        "hash_range",
					Description: `Possible value range of the hash key used by each partition.`,
				},
				resource.Attribute{
					Name:        "sequence_number_range",
					Description: `Sequence number range of each partition. ## Import Dis stream can be imported by ` + "`" + `name` + "`" + `. For example, ` + "`" + `` + "`" + `` + "`" + ` terraform import flexibleengine_dis_stream.example dis-demo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_dli_queue",
			Category:         "Data Lake Insight (DLI)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"data",
				"lake",
				"insight",
				"dli",
				"queue",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cu_count",
					Description: `(Required, Int) Minimum number of CUs that are bound to a queue. Initial value can be ` + "`" + `16` + "`" + `, ` + "`" + `64` + "`" + `, or ` + "`" + `256` + "`" + `. When scale_out or scale_in, the number must be a multiple of 16`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, String, ForceNew) Name of a queue. Name of a newly created resource queue. The name can contain only digits, letters, and underscores (_), but cannot contain only digits or start with an underscore (_). Length range: 1 to 128 characters. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, String, ForceNew) Description of a queue. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "queue_type",
					Description: `(Optional, String, ForceNew) Indicates the queue type. Changing this parameter will create a new resource. The options are as follows: - sql, - general > NOTE: If the type is not specified, the default value sql is used.`,
				},
				resource.Attribute{
					Name:        "resource_mode",
					Description: `(Optional, String, ForceNew) Queue resource mode. Changing this parameter will create a new resource. The options are as follows: - 0: indicates the shared resource mode. - 1: indicates the exclusive resource mode.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Map, ForceNew) Label of a queue. Changing this parameter will create a new resource. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Specifies a resource ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Time when a queue is created. ## Timeouts This resource provides the following timeouts configuration options: - ` + "`" + `update` + "`" + ` - Default is 45 minute. ## Import DLI queue can be imported by ` + "`" + `id` + "`" + `. For example, ` + "`" + `` + "`" + `` + "`" + ` terraform import flexibleengine_dli_queue.example abc123 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Specifies a resource ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Time when a queue is created. ## Timeouts This resource provides the following timeouts configuration options: - ` + "`" + `update` + "`" + ` - Default is 45 minute. ## Import DLI queue can be imported by ` + "`" + `id` + "`" + `. For example, ` + "`" + `` + "`" + `` + "`" + ` terraform import flexibleengine_dli_queue.example abc123 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_dms_kafka_instance",
			Category:         "Distributed Message Service (DMS)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"distributed",
				"message",
				"service",
				"dms",
				"kafka",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String, ForceNew) Specifies the region in which to create the DMS Kafka instance resource. If omitted, the provider-level region will be used. Changing this creates a new instance resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, String) Specifies the name of the DMS Kafka instance. An instance name starts with a letter, consists of 4 to 64 characters, and supports only letters, digits, hyphens (-) and underscores (_).`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Required, String, ForceNew) The baseline bandwidth of a Kafka instance, that is, the maximum amount of data transferred per unit time. The valid values are`,
				},
				resource.Attribute{
					Name:        "product_id",
					Description: `(Required, String, ForceNew) Specifies a product ID. You can get the value from id of [flexibleengine_dms_product](https://registry.terraform.io/providers/FlexibleEngineCloud/flexibleengine/latest/docs/data-sources/dms_product) data source. Changing this creates a new instance resource.`,
				},
				resource.Attribute{
					Name:        "storage_space",
					Description: `(Required, Int, ForceNew) Specifies the message storage capacity, the unit is GB. Value range: + When bandwidth is`,
				},
				resource.Attribute{
					Name:        "availability_zones",
					Description: `(Required, List, ForceNew) The names of the AZ where the Kafka instance resides. Changing this creates a new instance resource. ->`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, String, ForceNew) Specifies the ID of a VPC. Changing this creates a new instance resource.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required, String, ForceNew) Specifies the ID of a subnet. Changing this creates a new instance resource.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Required, String) Specifies the ID of a security group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, String) Specifies the description of the DMS Kafka instance. It is a character string containing not more than 1,024 characters.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `(Optional, String, ForceNew) Specifies the version of the Kafka engine. Valid values are "1.1.0" and "2.3.0". Defaults to`,
				},
				resource.Attribute{
					Name:        "storage_spec_code",
					Description: `(Optional, String, ForceNew) Specifies the storage I/O specification. Value range: + When bandwidth is`,
				},
				resource.Attribute{
					Name:        "manager_user",
					Description: `(Optional, String, ForceNew) Specifies the username for logging in to the Kafka Manager. The username consists of 4 to 64 characters and can contain letters, digits, hyphens (-), and underscores (_). Changing this creates a new instance resource.`,
				},
				resource.Attribute{
					Name:        "manager_password",
					Description: `(Optional, String, ForceNew) Specifies the password for logging in to the Kafka Manager. The password must meet the following complexity requirements: Must be 8 to 32 characters long. Must contain at least 2 of the following character types: lowercase letters, uppercase letters, digits, and special characters (` + "`" + `~!@#$%^&`,
				},
				resource.Attribute{
					Name:        "access_user",
					Description: `(Optional, String, ForceNew) Specifies a username who can accesse the instance with SASL authentication. A username consists of 4 to 64 characters and supports only letters, digits, and hyphens (-). Changing this creates a new instance resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional, String, ForceNew) Specifies the password of the access user. A password must meet the following complexity requirements: Must be 8 to 32 characters long. Must contain at least 2 of the following character types: lowercase letters, uppercase letters, digits, and special characters (` + "`" + `~!@#$%^&`,
				},
				resource.Attribute{
					Name:        "maintain_begin",
					Description: `(Optional, String) Specifies the time at which a maintenance time window starts. Format: HH:mm:ss. The start time must be set to 22:00:00, 02:00:00, 06:00:00, 10:00:00, 14:00:00, or 18:00:00. The system automatically allocates the default start time 02:00:00.`,
				},
				resource.Attribute{
					Name:        "maintain_end",
					Description: `(Optional, String) Specifies the time at which a maintenance time window ends. Format: HH:mm:ss. The end time is four hours later than the start time. For example, if the start time is 22:00:00, the end time is 02:00:00. The system automatically allocates the default end time 06:00:00. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Specifies a resource ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates the status of the DMS Kafka instance.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `Indicates the message engine, the value is "kafka".`,
				},
				resource.Attribute{
					Name:        "engine_type",
					Description: `Indicates the DMS Kafka instance type, the value is "cluster".`,
				},
				resource.Attribute{
					Name:        "product_spec_code",
					Description: `Indicates the DMS Kafka instance specification.`,
				},
				resource.Attribute{
					Name:        "partition_num",
					Description: `Indicates the maximum number of topics in the DMS Kafka instance.`,
				},
				resource.Attribute{
					Name:        "used_storage_space",
					Description: `Indicates the used message storage space. Unit: GB`,
				},
				resource.Attribute{
					Name:        "vpc_name",
					Description: `Indicates the name of a vpc.`,
				},
				resource.Attribute{
					Name:        "subnet_name",
					Description: `Indicates the name of a subnet.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `Indicates the name of a security group.`,
				},
				resource.Attribute{
					Name:        "node_num",
					Description: `Indicates the count of ECS instances.`,
				},
				resource.Attribute{
					Name:        "manegement_connect_address",
					Description: `Indicates the connection address of the Kafka Manager of a Kafka instance.`,
				},
				resource.Attribute{
					Name:        "connect_address",
					Description: `Indicates the IP addresses of the DMS Kafka instance.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Indicates the port number of the DMS Kafka instance.`,
				},
				resource.Attribute{
					Name:        "ssl_enable",
					Description: `Indicates whether the Kafka SASL_SSL is enabled.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Indicates the creation time of the DMS Kafka instance. ## Import DMS Kafka instance can be imported using the instance id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_dms_kafka_instance.instance_1 8d3c7938-dc47-4937-a30f-c80de381c5e3 ` + "`" + `` + "`" + `` + "`" + ` Note that the imported state may not be identical to your resource definition, because of ` + "`" + `access_user` + "`" + `, ` + "`" + `password` + "`" + `, ` + "`" + `manager_user` + "`" + ` and ` + "`" + `manager_password` + "`" + ` are missing from the API response due to security reason. It is generally recommended running ` + "`" + `terraform plan` + "`" + ` after importing a DMS Kafka instance. You can then decide if changes should be applied to the instance, or the resource definition should be updated to align with the instance. Also you can ignore changes as below. ` + "`" + `` + "`" + `` + "`" + ` resource "flexibleengine_dms_kafka_instance" "instance_1" { ... lifecycle { ignore_changes = [ access_user, password, manager_user, manager_password, ] } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Specifies a resource ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates the status of the DMS Kafka instance.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `Indicates the message engine, the value is "kafka".`,
				},
				resource.Attribute{
					Name:        "engine_type",
					Description: `Indicates the DMS Kafka instance type, the value is "cluster".`,
				},
				resource.Attribute{
					Name:        "product_spec_code",
					Description: `Indicates the DMS Kafka instance specification.`,
				},
				resource.Attribute{
					Name:        "partition_num",
					Description: `Indicates the maximum number of topics in the DMS Kafka instance.`,
				},
				resource.Attribute{
					Name:        "used_storage_space",
					Description: `Indicates the used message storage space. Unit: GB`,
				},
				resource.Attribute{
					Name:        "vpc_name",
					Description: `Indicates the name of a vpc.`,
				},
				resource.Attribute{
					Name:        "subnet_name",
					Description: `Indicates the name of a subnet.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `Indicates the name of a security group.`,
				},
				resource.Attribute{
					Name:        "node_num",
					Description: `Indicates the count of ECS instances.`,
				},
				resource.Attribute{
					Name:        "manegement_connect_address",
					Description: `Indicates the connection address of the Kafka Manager of a Kafka instance.`,
				},
				resource.Attribute{
					Name:        "connect_address",
					Description: `Indicates the IP addresses of the DMS Kafka instance.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Indicates the port number of the DMS Kafka instance.`,
				},
				resource.Attribute{
					Name:        "ssl_enable",
					Description: `Indicates whether the Kafka SASL_SSL is enabled.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Indicates the creation time of the DMS Kafka instance. ## Import DMS Kafka instance can be imported using the instance id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_dms_kafka_instance.instance_1 8d3c7938-dc47-4937-a30f-c80de381c5e3 ` + "`" + `` + "`" + `` + "`" + ` Note that the imported state may not be identical to your resource definition, because of ` + "`" + `access_user` + "`" + `, ` + "`" + `password` + "`" + `, ` + "`" + `manager_user` + "`" + ` and ` + "`" + `manager_password` + "`" + ` are missing from the API response due to security reason. It is generally recommended running ` + "`" + `terraform plan` + "`" + ` after importing a DMS Kafka instance. You can then decide if changes should be applied to the instance, or the resource definition should be updated to align with the instance. Also you can ignore changes as below. ` + "`" + `` + "`" + `` + "`" + ` resource "flexibleengine_dms_kafka_instance" "instance_1" { ... lifecycle { ignore_changes = [ access_user, password, manager_user, manager_password, ] } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_dms_kafka_topic",
			Category:         "Distributed Message Service (DMS)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"distributed",
				"message",
				"service",
				"dms",
				"kafka",
				"topic",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String, ForceNew) The region in which to create the DMS Kafka topic resource. If omitted, the provider-level region will be used. Changing this creates a new resource.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, String, ForceNew) Specifies the ID of the DMS Kafka instance to which the topic belongs. Changing this creates a new resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, String, ForceNew) Specifies the name of the topic. The name starts with a letter, consists of 4 to 64 characters, and supports only letters, digits, hyphens (-) and underscores (_). Changing this creates a new resource.`,
				},
				resource.Attribute{
					Name:        "partitions",
					Description: `(Optional, Int, ForceNew) Specifies the partition number. The value ranges from 1 to 50 and defaults to 3. Changing this creates a new resource.`,
				},
				resource.Attribute{
					Name:        "replicas",
					Description: `(Optional, Int, ForceNew) Specifies the replica number. The value ranges from 1 to 3 and defaults to 3. Changing this creates a new resource.`,
				},
				resource.Attribute{
					Name:        "aging_time",
					Description: `(Optional, Int, ForceNew) Specifies the aging time in hours. The value ranges from 1 to 720 and defaults to 72. Changing this creates a new resource.`,
				},
				resource.Attribute{
					Name:        "sync_replication",
					Description: `(Optional, Bool, ForceNew) Whether or not to enable synchronous replication. Changing this creates a new resource.`,
				},
				resource.Attribute{
					Name:        "sync_flushing",
					Description: `(Optional, Bool, ForceNew) Whether or not to enable synchronous flushing. Changing this creates a new resource. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID which equals to the topic name. ## Import DMS Kafka topics can be imported using the Kafka instance ID and topic name separated by a slash, e.g.: ` + "`" + `` + "`" + `` + "`" + `sh terraform import flexibleengine_dms_kafka_topic.topic c8057fe5-23a8-46ef-ad83-c0055b4e0c5c/topic_1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID which equals to the topic name. ## Import DMS Kafka topics can be imported using the Kafka instance ID and topic name separated by a slash, e.g.: ` + "`" + `` + "`" + `` + "`" + `sh terraform import flexibleengine_dms_kafka_topic.topic c8057fe5-23a8-46ef-ad83-c0055b4e0c5c/topic_1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_dns_ptrrecord_v2",
			Category:         "Domain Name Service (DNS)",
			ShortDescription: ``,
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
					Name:        "region",
					Description: `(Optional) The region in which to create the PTR record. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new PTR record.`,
				},
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
					Description: `(Required) The ID of the FloatingIP/EIP. Changing this creates a new PTR record.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The time to live (TTL) of the record set (in seconds). The value range is 300–2147483647. The default value is 300.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags key/value pairs to associate with the PTR record. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The PTR record ID, which is in {region}:{floatingip_id} format.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The address of the FloatingIP/EIP. ## Import PTR records can be imported using region and floatingip/eip ID, separated by a colon(:), e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_dns_ptrrecord_v2.ptr_1 eu-west-0:d90ce693-5ccf-4136-a0ed-152ce412b6b9 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The PTR record ID, which is in {region}:{floatingip_id} format.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The address of the FloatingIP/EIP. ## Import PTR records can be imported using region and floatingip/eip ID, separated by a colon(:), e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_dns_ptrrecord_v2.ptr_1 eu-west-0:d90ce693-5ccf-4136-a0ed-152ce412b6b9 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_dns_recordset_v2",
			Category:         "Domain Name Service (DNS)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "Network-DNS.svg",
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
					Description: `(Optional) The region in which to create the DNS record set. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new DNS record set.`,
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
					Description: `(Required) The type of record set. The options include ` + "`" + `A` + "`" + `, ` + "`" + `AAAA` + "`" + `, ` + "`" + `MX` + "`" + `, ` + "`" + `CNAME` + "`" + `, ` + "`" + `TXT` + "`" + `, ` + "`" + `NS` + "`" + `, ` + "`" + `SRV` + "`" + `, ` + "`" + `CAA` + "`" + `, and ` + "`" + `PTR` + "`" + `. Changing this creates a new DNS record set.`,
				},
				resource.Attribute{
					Name:        "records",
					Description: `(Required) An array of DNS records.`,
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
					Name:        "tags",
					Description: `(Optional) The key/value pairs to associate with the record set.`,
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
					Description: `See Argument Reference above. ## Import This resource can be imported by specifying the zone ID and recordset ID, separated by a forward slash. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_dns_recordset_v2.recordset_1 <zone_id>/<recordset_id> ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `See Argument Reference above. ## Import This resource can be imported by specifying the zone ID and recordset ID, separated by a forward slash. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_dns_recordset_v2.recordset_1 <zone_id>/<recordset_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_dns_zone_v2",
			Category:         "Domain Name Service (DNS)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "zone.svg",
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
					Description: `(Optional) The region in which to create the DNS zone. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new DNS zone.`,
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
					Name:        "tags",
					Description: `(Optional, Map) The key/value pairs to associate with the zone.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional options. Changing this creates a new DNS zone. The ` + "`" + `router` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `(Required) The VPC UUID.`,
				},
				resource.Attribute{
					Name:        "router_region",
					Description: `(Optional) The region of the VPC. Defaults to the ` + "`" + `region` + "`" + `. ## Attributes Reference The following attributes are exported:`,
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
					Description: `See Argument Reference above. ## Import This resource can be imported by specifying the zone ID: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_dns_zone_v2.zone_1 <zone_id> ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `See Argument Reference above. ## Import This resource can be imported by specifying the zone ID: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_dns_zone_v2.zone_1 <zone_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_drs_replication_v2",
			Category:         "Deprecated",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "drs.svg",
			Keywords: []string{
				"deprecated",
				"drs",
				"replication",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the EVS replication pair. The name can contain a maximum of 255 bytes.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the EVS replication pair. The description can contain a maximum of 255 bytes.`,
				},
				resource.Attribute{
					Name:        "volume_ids",
					Description: `(Required) An array of one or more IDs of the EVS disks used to create the EVS replication pair.`,
				},
				resource.Attribute{
					Name:        "priority_station",
					Description: `(Required) The primary AZ of the EVS replication pair. That is the AZ where the production disk belongs.`,
				},
				resource.Attribute{
					Name:        "replication_model",
					Description: `(Optional) The type of the EVS replication pair. Currently only type hypermetro is supported. ## Attributes Reference The following attributes are exported:`,
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
					Name:        "volume_ids",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "priority_station",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "replication_model",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the EVS replication pair.`,
				},
				resource.Attribute{
					Name:        "replication_consistency_group_id",
					Description: `The ID of the replication consistency group where the EVS replication pair belongs.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The creation time of the EVS replication pair.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The update time of the EVS replication pair.`,
				},
				resource.Attribute{
					Name:        "replication_status",
					Description: `The replication status of the EVS replication pair.`,
				},
				resource.Attribute{
					Name:        "progress",
					Description: `The synchronization progress of the EVS replication pair. Unit: %.`,
				},
				resource.Attribute{
					Name:        "failure_detail",
					Description: `The returned error code if the EVS replication pair status is error.`,
				},
				resource.Attribute{
					Name:        "record_metadata",
					Description: `The metadata of the EVS replication pair.`,
				},
				resource.Attribute{
					Name:        "fault_level",
					Description: `The fault level of the EVS replication pair.`,
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
					Name:        "volume_ids",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "priority_station",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "replication_model",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the EVS replication pair.`,
				},
				resource.Attribute{
					Name:        "replication_consistency_group_id",
					Description: `The ID of the replication consistency group where the EVS replication pair belongs.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The creation time of the EVS replication pair.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The update time of the EVS replication pair.`,
				},
				resource.Attribute{
					Name:        "replication_status",
					Description: `The replication status of the EVS replication pair.`,
				},
				resource.Attribute{
					Name:        "progress",
					Description: `The synchronization progress of the EVS replication pair. Unit: %.`,
				},
				resource.Attribute{
					Name:        "failure_detail",
					Description: `The returned error code if the EVS replication pair status is error.`,
				},
				resource.Attribute{
					Name:        "record_metadata",
					Description: `The metadata of the EVS replication pair.`,
				},
				resource.Attribute{
					Name:        "fault_level",
					Description: `The fault level of the EVS replication pair.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_drs_replicationconsistencygroup_v2",
			Category:         "Deprecated",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "drs.svg",
			Keywords: []string{
				"deprecated",
				"drs",
				"replicationconsistencygroup",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the replication consistency group. The name can contain a maximum of 255 bytes.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the replication consistency group. The description can contain a maximum of 255 bytes.`,
				},
				resource.Attribute{
					Name:        "replication_ids",
					Description: `(Required) An array of one or more IDs of the EVS replication pairs used to create the replication consistency group.`,
				},
				resource.Attribute{
					Name:        "priority_station",
					Description: `(Required) The primary AZ of the replication consistency group. That is the AZ where the production disk belongs.`,
				},
				resource.Attribute{
					Name:        "replication_model",
					Description: `(Optional) The type of the created replication consistency group. Currently only type hypermetro is supported. ## Attributes Reference The following attributes are exported:`,
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
					Name:        "replication_ids",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "priority_station",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "replication_model",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the replication consistency group.`,
				},
				resource.Attribute{
					Name:        "replication_status",
					Description: `The replication status of the replication consistency group.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The creation time of the replication consistency group.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The update time of the replication consistency group.`,
				},
				resource.Attribute{
					Name:        "failure_detail",
					Description: `The returned error code if the replication consistency group status is error.`,
				},
				resource.Attribute{
					Name:        "fault_level",
					Description: `The fault level of the replication consistency group.`,
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
					Name:        "replication_ids",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "priority_station",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "replication_model",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the replication consistency group.`,
				},
				resource.Attribute{
					Name:        "replication_status",
					Description: `The replication status of the replication consistency group.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The creation time of the replication consistency group.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The update time of the replication consistency group.`,
				},
				resource.Attribute{
					Name:        "failure_detail",
					Description: `The returned error code if the replication consistency group status is error.`,
				},
				resource.Attribute{
					Name:        "fault_level",
					Description: `The fault level of the replication consistency group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_dws_cluster_v1",
			Category:         "Data Warehouse Service (DWS)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "Data Analysis&AI-DWS.svg",
			Keywords: []string{
				"data",
				"warehouse",
				"service",
				"dws",
				"cluster",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Cluster name, which must be unique and contains 4 to 64 characters, which consist of letters, digits, hyphens (-), or underscores (_) only and must start with a letter.`,
				},
				resource.Attribute{
					Name:        "node_type",
					Description: `(Required) Node type.`,
				},
				resource.Attribute{
					Name:        "number_of_node",
					Description: `(Required) Number of nodes in a cluster. The value ranges from 3 to 32.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `(Required) Administrator username for logging in to a data warehouse cluster The administrator username must: - Consist of lowercase letters, digits, or underscores. - Start with a lowercase letter or an underscore. - Contain 1 to 63 characters. - Cannot be a keyword of the DWS database.`,
				},
				resource.Attribute{
					Name:        "user_pwd",
					Description: `(Required) Administrator password for logging in to a data warehouse cluster. A password must conform to the following rules: - Contains 8 to 32 characters. - Cannot be the same as the username or the username written in reverse order. - Contains three types of lowercase letters, uppercase letters, digits and special characters ~!@#%^&`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) VPC ID, which is used for configuring cluster network.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) Subnet ID, which is used for configuring cluster network.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Required) ID of a security group. The ID is used for configuring cluster network.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Service port of a cluster (8000 to 10000). The default value is 8000.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `(Optional) Public IP address. The object structure is documented below.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) AZ in a cluster The ` + "`" + `public_ip` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "eip_id",
					Description: `(Optional) EIP ID`,
				},
				resource.Attribute{
					Name:        "public_bind_type",
					Description: `(Optional) Binding type of an EIP. The value can be either of the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Cluster ID`,
				},
				resource.Attribute{
					Name:        "endpoints",
					Description: `The private network connection information about the cluster. The object structure is documented below.`,
				},
				resource.Attribute{
					Name:        "public_endpoints",
					Description: `The public network connection information about the cluster. The object structure is documented below.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Cluster status, which can be one of the following:`,
				},
				resource.Attribute{
					Name:        "sub_status",
					Description: `Sub-status of clusters in the AVAILABLE state.`,
				},
				resource.Attribute{
					Name:        "task_status",
					Description: `Cluster management task.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Data warehouse version`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `Cluster creation time. The format is ISO8601:YYYY-MM-DDThh:mm:ssZ.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `Last modification time of a cluster. The format is ISO8601:YYYY-MM-DDThh:mm:ssZ. The ` + "`" + `endpoints` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "connect_info",
					Description: `Private network connection information`,
				},
				resource.Attribute{
					Name:        "jdbc_url",
					Description: `JDBC URL. The following is the default format: jdbc:postgresql://< connect_info>/<YOUR_DATABASE_NAME> The ` + "`" + `public_endpoints` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "public_connect_info",
					Description: `Public network connection information`,
				},
				resource.Attribute{
					Name:        "jdbc_url",
					Description: `JDBC URL. The following is the default format: jdbc:postgresql://< public_connect_info>/<YOUR_DATABASE_NAME>`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Cluster ID`,
				},
				resource.Attribute{
					Name:        "endpoints",
					Description: `The private network connection information about the cluster. The object structure is documented below.`,
				},
				resource.Attribute{
					Name:        "public_endpoints",
					Description: `The public network connection information about the cluster. The object structure is documented below.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Cluster status, which can be one of the following:`,
				},
				resource.Attribute{
					Name:        "sub_status",
					Description: `Sub-status of clusters in the AVAILABLE state.`,
				},
				resource.Attribute{
					Name:        "task_status",
					Description: `Cluster management task.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Data warehouse version`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `Cluster creation time. The format is ISO8601:YYYY-MM-DDThh:mm:ssZ.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `Last modification time of a cluster. The format is ISO8601:YYYY-MM-DDThh:mm:ssZ. The ` + "`" + `endpoints` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "connect_info",
					Description: `Private network connection information`,
				},
				resource.Attribute{
					Name:        "jdbc_url",
					Description: `JDBC URL. The following is the default format: jdbc:postgresql://< connect_info>/<YOUR_DATABASE_NAME> The ` + "`" + `public_endpoints` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "public_connect_info",
					Description: `Public network connection information`,
				},
				resource.Attribute{
					Name:        "jdbc_url",
					Description: `JDBC URL. The following is the default format: jdbc:postgresql://< public_connect_info>/<YOUR_DATABASE_NAME>`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_elb_backend",
			Category:         "Deprecated",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "Network-ELB.svg",
			Keywords: []string{
				"deprecated",
				"elb",
				"backend",
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
					Name:        "address",
					Description: `(Required) Specifies the private IP address of the backend member. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Specifies the backend member ID.`,
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
					Name:        "address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "server_address",
					Description: `Specifies the floating IP address assigned to the backend member.`,
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
					Name:        "id",
					Description: `Specifies the backend member ID.`,
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
					Name:        "address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "server_address",
					Description: `Specifies the floating IP address assigned to the backend member.`,
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
			Type:             "flexibleengine_elb_certificate",
			Category:         "Elastic Load Balance (ELB)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"load",
				"balance",
				"elb",
				"certificate",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String, ForceNew) The region in which to create the ELB certificate resource. If omitted, the provider-level region will be used. Changing this creates a new certificate.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, String) Specifies the name of the certificate. Does not have to be unique.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, String) Specifies the description of the certificate.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional, String, ForceNew) Specifies the certificate type. The default value is "server". The value can be one of the following: + server: indicates the server certificate. + client: indicates the CA certificate.`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(Required, String) Specifies the public encrypted key of the certificate, PEM format.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Optional, String) Specifies the private encrypted key of the certificate, PEM format. This parameter is valid and mandatory only when ` + "`" + `type` + "`" + ` is set to "server".`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional, String) Specifies the domain of the certificate. The value contains a maximum of 100 characters. This parameter is valid only when ` + "`" + `type` + "`" + ` is set to "server". ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Specifies a resource ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Indicates the update time.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Indicates the creation time.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `Indicates the expire time. ## Timeouts This resource provides the following timeouts configuration options:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `Default is 10 minute.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `Default is 10 minute.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `Default is 5 minute. ## Import ELB certificate can be imported using the certificate ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_elb_certificate.certificate_1 5c20fdad-7288-11eb-b817-0255ac10158b ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Specifies a resource ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Indicates the update time.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Indicates the creation time.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `Indicates the expire time. ## Timeouts This resource provides the following timeouts configuration options:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `Default is 10 minute.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `Default is 10 minute.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `Default is 5 minute. ## Import ELB certificate can be imported using the certificate ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_elb_certificate.certificate_1 5c20fdad-7288-11eb-b817-0255ac10158b ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_elb_health",
			Category:         "Deprecated",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "Network-ELB.svg",
			Keywords: []string{
				"deprecated",
				"elb",
				"health",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to create the elb health. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new elb health.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Required) Specifies the ID of the listener to which the health check belongs.`,
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
					Name:        "id",
					Description: `Specifies the health check ID.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
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
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Specifies the health check ID.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
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
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_elb_ipgroup",
			Category:         "Elastic Load Balance (ELB)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"load",
				"balance",
				"elb",
				"ipgroup",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String, ForceNew) The region in which to create the ip group resource. If omitted, the provider-level region will be used. Changing this creates a new ip group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, String) Specifies the name of the ip group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, String) Specifies the description of the ip group.`,
				},
				resource.Attribute{
					Name:        "ip_list",
					Description: `(Required, List) Specifies an array of one or more ip addresses. The ip_list object structure is documented below. The ` + "`" + `ip_list` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required, String) IP address or CIDR block.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, String) Human-readable description for the ip. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The uuid of the ip group. ## Timeouts This resource provides the following timeouts configuration options:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `Default is 10 minute.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `Default is 10 minute.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `Default is 5 minute. ELB IP group can be imported using the IP group ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_elb_ipgroup.group_1 5c20fdad-7288-11eb-b817-0255ac10158b ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The uuid of the ip group. ## Timeouts This resource provides the following timeouts configuration options:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `Default is 10 minute.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `Default is 10 minute.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `Default is 5 minute. ELB IP group can be imported using the IP group ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_elb_ipgroup.group_1 5c20fdad-7288-11eb-b817-0255ac10158b ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_elb_listener",
			Category:         "Deprecated",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "Network-ELB.svg",
			Keywords: []string{
				"deprecated",
				"elb",
				"listener",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to create the elb listener. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new elb listener.`,
				},
				resource.Attribute{
					Name:        "loadbalancer_id",
					Description: `(Required) Specifies the ID of the load balancer to which the listener belongs.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the load balancer name. The name is a string of 1 to 64 characters that consist of letters, digits, underscores (_), and hyphens (-).`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Specifies the listening protocol used for layer 4 or 7. The value can be HTTP, TCP, HTTPS, or UDP.`,
				},
				resource.Attribute{
					Name:        "protocol_port",
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
					Name:        "description",
					Description: `(Optional) Provides supplementary information about the listener. The value is a string of 0 to 128 characters and cannot be <>.`,
				},
				resource.Attribute{
					Name:        "session_sticky",
					Description: `(Optional) Specifies whether to enable sticky session. The value can be true or false. The Sticky session is enabled when the value is true, and is disabled when the value is false. If the value of protocol is HTTP, HTTPS, or TCP, and the value of lb_algorithm is not roundrobin, the value of this parameter can only be false.`,
				},
				resource.Attribute{
					Name:        "session_sticky_type",
					Description: `(Optional) Specifies the cookie processing method. The value is insert. insert indicates that the cookie is inserted by the load balancer. This parameter is valid when protocol is set to HTTP, and session_sticky to true. The default value is insert. This parameter is invalid when protocol is set to TCP or UDP, which means the parameter is empty.`,
				},
				resource.Attribute{
					Name:        "cookie_timeout",
					Description: `(Optional) Specifies the cookie timeout period (minutes). This parameter is valid when protocol is set to HTTP, session_sticky to true, and session_sticky_type to insert. This parameter is invalid when protocol is set to TCP or UDP. The value ranges from 1 to 1440.`,
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
					Name:        "id",
					Description: `Specifies the listener ID.`,
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
					Name:        "loadbalancer_id",
					Description: `See Argument Reference above.`,
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
					Name:        "session_sticky_type",
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
					Name:        "admin_state_up",
					Description: `Specifies the status of the load balancer. Value range: false: The load balancer is disabled. true: The load balancer runs properly.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Specifies the listener ID.`,
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
					Name:        "loadbalancer_id",
					Description: `See Argument Reference above.`,
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
					Name:        "session_sticky_type",
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
					Name:        "admin_state_up",
					Description: `Specifies the status of the load balancer. Value range: false: The load balancer is disabled. true: The load balancer runs properly.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_elb_loadbalancer",
			Category:         "Deprecated",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "Network-ELB.svg",
			Keywords: []string{
				"deprecated",
				"elb",
				"loadbalancer",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to create the loadbalancer. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new loadbalancer.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the load balancer name. The name is a string of 1 to 64 characters that consist of letters, digits, underscores (_), and hyphens (-).`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Specifies the load balancer type. The value can be Internal or External.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) Specifies the VPC ID.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Provides supplementary information about the listener. The value is a string of 0 to 128 characters and cannot be <>.`,
				},
				resource.Attribute{
					Name:        "vip_address",
					Description: `(Optional) Specifies the IP address provided by ELB. When type is set to External, the value of this parameter is the elastic IP address. When type is set to Internal, the value of this parameter is the private network IP address. You can select an existing elastic IP address and create a public network load balancer. When this parameter is configured, parameter ` + "`" + `bandwidth` + "`" + ` is invalid.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Optional) Specifies the bandwidth (Mbit/s). This parameter is valid when type is set to External, and it is invalid when type is set to Internal. The value ranges from 1 to 300.`,
				},
				resource.Attribute{
					Name:        "vip_subnet_id",
					Description: `(Optional) Specifies the ID of the private network to be added. This parameter is mandatory when type is set to Internal, and it is invalid when type is set to External.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Optional) Specifies the security group ID. The value is a string of 1 to 200 characters that consists of uppercase and lowercase letters, digits, and hyphens (-). This parameter is mandatory when type is set to Internal, and it is invalid when type is set to External.`,
				},
				resource.Attribute{
					Name:        "az",
					Description: `(Optional) Specifies the ID of the availability zone (AZ). This parameter is mandatory when type is set to Internal, and it is invalid when type is set to External.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) Specifies the status of the load balancer. Defaults to true. + true: indicates that the load balancer is running. + false: indicates that the load balancer is stopped.`,
				},
				resource.Attribute{
					Name:        "tenantid",
					Description: `(Optional) Specifies the tenant ID. This parameter is mandatory only when type is set to Internal. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Specifies the load balancer ID.`,
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
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Specifies the load balancer ID.`,
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
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_enterprise_project",
			Category:         "Enterprise Project Management Service (EPS)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"enterprise",
				"project",
				"management",
				"service",
				"eps",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, String) Specifies the name of the enterprise project. This parameter can contain 1 to 64 characters. Only letters, digits, underscores (_), and hyphens (-) are allowed. The name must be unique in the domain and cannot include any form of the word "default" ("deFaulT", for instance).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, String) Specifies the description of the enterprise project.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, Bool) Specifies whether to enable the enterprise project. Default to`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates the status of an enterprise project. + 1 indicates Enabled. + 2 indicates Disabled.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Indicates the type of the enterprise project.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Indicates the UTC time when the enterprise project was created. Example: 2018-05-18T06:49:06Z`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Indicates the UTC time when the enterprise project was modified. Example: 2018-05-28T02:21:36Z ## Import Enterprise projects can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_enterprise_project.test 88f889c7-270e-4e77-8230-bf7db08d9b0e ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `Indicates the status of an enterprise project. + 1 indicates Enabled. + 2 indicates Disabled.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Indicates the type of the enterprise project.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Indicates the UTC time when the enterprise project was created. Example: 2018-05-18T06:49:06Z`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Indicates the UTC time when the enterprise project was modified. Example: 2018-05-28T02:21:36Z ## Import Enterprise projects can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_enterprise_project.test 88f889c7-270e-4e77-8230-bf7db08d9b0e ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_fgs_dependency",
			Category:         "FunctionGraph",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"functiongraph",
				"fgs",
				"dependency",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String, ForceNew) Specifies the region in which to create a custom dependency package. If omitted, the provider-level region will be used. Changing this will create a new resource.`,
				},
				resource.Attribute{
					Name:        "runtime",
					Description: `(Required, String) Specifies the dependency package runtime. The valid values are`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, String) Specifies the dependeny name. The name can contain a maximum of 96 characters and must start with a letter and end with a letter or digit. Only letters, digits, underscores (_), periods (.), and hyphens (-) are allowed.`,
				},
				resource.Attribute{
					Name:        "link",
					Description: `(Required, String) Specifies the OBS bucket path where the dependency package is located. The OBS object URL must be in zip format, such as 'https://obs-terraform.oss.eu-west-0.prod-cloud-ocb.orange-business.com/dependencies/sdkcore.zip'. -> A link can only be used to create at most one dependency package.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, String) Specifies the dependency description. The description can contain a maximum of 512 characters. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The dependency ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `The base64 encoded digest of the dependency after encryption by MD5.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `The unique ID of the dependency package.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The dependency package size in bytes. ## Import Dependencies can be imported using the ` + "`" + `id` + "`" + `, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_fgs_dependency.test 795e722f-0c23-41b6-a189-dcd56f889cf6 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The dependency ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `The base64 encoded digest of the dependency after encryption by MD5.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `The unique ID of the dependency package.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The dependency package size in bytes. ## Import Dependencies can be imported using the ` + "`" + `id` + "`" + `, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_fgs_dependency.test 795e722f-0c23-41b6-a189-dcd56f889cf6 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_fgs_function",
			Category:         "FunctionGraph",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"functiongraph",
				"fgs",
				"function",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String, ForceNew) Specifies the region in which to create the Function resource. If omitted, the provider-level region will be used. Changing this creates a new Function resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, String, ForceNew) Specifies the name of the function.`,
				},
				resource.Attribute{
					Name:        "app",
					Description: `(Required, String) Specifies the group to which the function belongs.`,
				},
				resource.Attribute{
					Name:        "handler",
					Description: `(Required, String) Specifies the entry point of the function.`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `(Required, Int) Specifies the memory size(MB) allocated to the function.`,
				},
				resource.Attribute{
					Name:        "runtime",
					Description: `(Required, String, ForceNew) Specifies the environment for executing the function. Changing this creates a new Function resource.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Required, Int) Specifies the timeout interval of the function, ranges from 3s to 900s.`,
				},
				resource.Attribute{
					Name:        "code_type",
					Description: `(Required, String) Specifies the function code type, which can be inline: inline code, zip: ZIP file, jar: JAR file or java functions, obs: function code stored in an OBS bucket.`,
				},
				resource.Attribute{
					Name:        "func_code",
					Description: `(Optional, String) Specifies the function code. When code_type is set to inline, zip, or jar, this parameter is mandatory, and the code can be encoded using Base64 or just with the text code.`,
				},
				resource.Attribute{
					Name:        "code_url",
					Description: `(Optional, String) Specifies the code url. This parameter is mandatory when code_type is set to obs.`,
				},
				resource.Attribute{
					Name:        "code_filename",
					Description: `(Optional, String) Specifies the name of a function file, This field is mandatory only when coe_type is set to jar or zip.`,
				},
				resource.Attribute{
					Name:        "depend_list",
					Description: `(Optional, String) Specifies the ID list of the dependencies.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional, String) Specifies the Key/Value information defined for the function. Key/value data might be parsed with [Terraform ` + "`" + `jsonencode()` + "`" + ` function]('https://www.terraform.io/docs/language/functions/jsonencode.html').`,
				},
				resource.Attribute{
					Name:        "encrypted_user_data",
					Description: `(Optional, String) Specifies the key/value information defined to be encrypted for the function. The format is the same as ` + "`" + `user_data` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "agency",
					Description: `(Optional, String) Specifies the agency. This parameter is mandatory if the function needs to access other cloud services.`,
				},
				resource.Attribute{
					Name:        "app_agency",
					Description: `(Optional, String) Specifies An execution agency enables you to obtain a token or an AK/SK for accessing other cloud services.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, String) Specifies the description of the function.`,
				},
				resource.Attribute{
					Name:        "initializer_handler",
					Description: `(Optional, String) Specifies the initializer of the function.`,
				},
				resource.Attribute{
					Name:        "initializer_timeout",
					Description: `(Optional, Int) Specifies the maximum duration the function can be initialized. Value range: 1s to 300s.`,
				},
				resource.Attribute{
					Name:        "enterprise_project_id",
					Description: `(Optional, String, ForceNew) Specifies the enterprise project id of the function. Changing this creates a new function.`,
				},
				resource.Attribute{
					Name:        "mount_user_id",
					Description: `(Optional, String) Specifies the user ID, a non-0 integer from –1 to 65534. Default to -1.`,
				},
				resource.Attribute{
					Name:        "mount_user_group_id",
					Description: `(Optional, String) Specifies the user group ID, a non-0 integer from –1 to 65534. Default to -1.`,
				},
				resource.Attribute{
					Name:        "func_mounts",
					Description: `(Optional, List) Specifies the file system list. The ` + "`" + `func_mounts` + "`" + ` object structure is documented below. The ` + "`" + `func_mounts` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "mount_type",
					Description: `(Required, String) Specifies the mount type. Options: sfs, sfsTurbo, and ecs.`,
				},
				resource.Attribute{
					Name:        "mount_resource",
					Description: `(Required, String) Specifies the ID of the mounted resource (corresponding cloud service).`,
				},
				resource.Attribute{
					Name:        "mount_share_path",
					Description: `(Required, String) Specifies the remote mount path. Example: 192.168.0.12:/data.`,
				},
				resource.Attribute{
					Name:        "local_mount_path",
					Description: `(Required, String) Specifies the function access path. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Specifies a resource ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "func_mounts/status",
					Description: `The status of file system.`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `Uniform Resource Name`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version of the function ## Timeouts This resource provides the following timeouts configuration options:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `Default is 10 minute.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `Default is 10 minute. ## Import Functions can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_fgs_function.my-func 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Specifies a resource ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "func_mounts/status",
					Description: `The status of file system.`,
				},
				resource.Attribute{
					Name:        "urn",
					Description: `Uniform Resource Name`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version of the function ## Timeouts This resource provides the following timeouts configuration options:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `Default is 10 minute.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `Default is 10 minute. ## Import Functions can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_fgs_function.my-func 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_fgs_trigger",
			Category:         "FunctionGraph",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"functiongraph",
				"fgs",
				"trigger",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String, ForceNew) Specifies the region in which to create the trigger resource. If omitted, the provider-level region will be used. Changing this will create a new trigger resource.`,
				},
				resource.Attribute{
					Name:        "function_urn",
					Description: `(Required, String, ForceNew) Specifies the Uniform Resource Name (URN) of the function. Changing this will create a new trigger resource.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required, String, ForceNew) Specifies the type of the function. The valid values currently only support`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional, String) Specifies whether trigger is enabled. The valid values are`,
				},
				resource.Attribute{
					Name:        "timer",
					Description: `(Optional, List, ForceNew) Specifies the configuration of the timing trigger. Changing this will create a new trigger resource. The [object](#fgs_trigger_timer) structure is documented below.`,
				},
				resource.Attribute{
					Name:        "obs",
					Description: `(Optional, List, ForceNew) Specifies the configuration of the OBS trigger. Changing this will create a new trigger resource. The [object](#fgs_trigger_obs) structure is documented below.`,
				},
				resource.Attribute{
					Name:        "smn",
					Description: `(Optional, List, ForceNew) Specifies the configuration of the SMN trigger. Changing this will create a new trigger resource. The [object](#fgs_trigger_smn) structure is documented below.`,
				},
				resource.Attribute{
					Name:        "dis",
					Description: `(Optional, List, ForceNew) Specifies the configuration of the DIS trigger. Changing this will create a new trigger resource. The [object](#fgs_trigger_dis) structure is documented below. ->`,
				},
				resource.Attribute{
					Name:        "apig",
					Description: `(Optional, List, ForceNew) Specifies the configuration of the shared APIG trigger. Changing this will create a new trigger resource. The [object](#fgs_trigger_apig) structure is documented below. <a name="fgs_trigger_timer"></a> The ` + "`" + `timer` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, String, ForceNew) Specifies the trigger name, which can contains of 1 to 64 characters. The name must start with a letter, only letters, digits, hyphens (-) and underscores (_) are allowed. Changing this will create a new trigger resource.`,
				},
				resource.Attribute{
					Name:        "schedule_type",
					Description: `(Required, String, ForceNew) Specifies the type of the time schedule. The valid values are`,
				},
				resource.Attribute{
					Name:        "schedule",
					Description: `(Required, String, ForceNew) Specifies the time schedule. For the rate type, schedule is composed of time and time unit. The time unit supports minutes (m), hours (h) and days (d). For the corn expression, please refer to the [User Guide](https://docs.prod-cloud-ocb.orange-business.com/usermanual/functiongraph/functiongraph_01_0908.html). Changing this will create a new trigger resource.`,
				},
				resource.Attribute{
					Name:        "additional_information",
					Description: `(Optional, String, ForceNew) Specifies the event used by the timer to trigger the function. Changing this will create a new trigger resource. <a name="fgs_trigger_obs"></a> The ` + "`" + `obs` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "bucket_name",
					Description: `(Required, String, ForceNew) Specifies the OBS bucket name. Changing this will create a new trigger resource.`,
				},
				resource.Attribute{
					Name:        "events",
					Description: `(Required, List, ForceNew) Specifies the events that can trigger functions. Changing this will create a new trigger resource. The valid values are as follows: +`,
				},
				resource.Attribute{
					Name:        "event_notification_name",
					Description: `(Required, String, ForceNew) Specifies the event notification name. Changing this will create a new trigger resource.`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Optional, String, ForceNew) Specifies the prefix to limit notifications to objects beginning with this keyword. Changing this will create a new trigger resource.`,
				},
				resource.Attribute{
					Name:        "suffix",
					Description: `(Optional, String, ForceNew) Specifies the suffix to limit notifications to objects ending with this keyword. Changing this will create a new trigger resource. <a name="fgs_trigger_smn"></a> The ` + "`" + `smn` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "topic_urn",
					Description: `(Required, String, ForceNew) Specifies the Uniform Resource Name (URN) for SMN topic. Changing this will create a new trigger resource. <a name="fgs_trigger_dis"></a> The ` + "`" + `dis` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "stream_name",
					Description: `(Required, String, ForceNew) Specifies the name of the DIS stream resource. Changing this will create a new trigger resource.`,
				},
				resource.Attribute{
					Name:        "starting_position",
					Description: `(Required, String, ForceNew) Specifies the type of starting position for DIS queue. The valid values are as follows: +`,
				},
				resource.Attribute{
					Name:        "max_fetch_bytes",
					Description: `(Required, Int, ForceNew) Specifies the maximum volume of data that can be obtained for a single request, in Byte. Only the records with a size smaller than this value can be obtained. The valid value is range from ` + "`" + `1,024` + "`" + ` to ` + "`" + `4,194,304` + "`" + `. Changing this will create a new trigger resource.`,
				},
				resource.Attribute{
					Name:        "pull_period",
					Description: `(Required, Int, ForceNew) Specifies the interval at which data is pulled from the specified stream. The valid value is range from ` + "`" + `2` + "`" + ` to ` + "`" + `60,000` + "`" + `. Changing this will create a new trigger resource.`,
				},
				resource.Attribute{
					Name:        "serial_enable",
					Description: `(Required, Bool, ForceNew) Specifies the determines whether to pull data only after the data pulled in the last period has been processed. Changing this will create a new trigger resource. <a name="fgs_trigger_apig"></a> The ` + "`" + `apig` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required, String, ForceNew) Specifies the ID of the APIG group to which the API belongs. Changing this will create a new trigger resource.`,
				},
				resource.Attribute{
					Name:        "env_name",
					Description: `(Required, String, ForceNew) Specifies the API environment name. Changing this will create a new trigger resource.`,
				},
				resource.Attribute{
					Name:        "api_name",
					Description: `(Required, String, ForceNew) Specifies the API name. Changing this will create a new trigger resource.`,
				},
				resource.Attribute{
					Name:        "security_authentication",
					Description: `(Optional, String, ForceNew) Specifies the security authentication mode. The valid values are`,
				},
				resource.Attribute{
					Name:        "request_protocol",
					Description: `(Optional, String, ForceNew) Specifies the request protocol of the API. The valid value are`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional, Int, ForceNew) Specifies the timeout for request sending. The valid value is range form ` + "`" + `1` + "`" + ` to ` + "`" + `60,000` + "`" + `, default to ` + "`" + `5,000` + "`" + `. Changing this will create a new trigger resource. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `resource ID in UUID format. ## Timeouts This resource provides the following timeouts configuration options:`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `Default is 2 minute.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `resource ID in UUID format. ## Timeouts This resource provides the following timeouts configuration options:`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `Default is 2 minute.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_fw_firewall_group_v2",
			Category:         "Network ACL",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "firewall.svg",
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
					Description: `See Argument Reference above. ## Import Firewall Groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_fw_firewall_group_v2.firewall_group_1 c9e39fb2-ce20-46c8-a964-25f3898c7a97 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `See Argument Reference above. ## Import Firewall Groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_fw_firewall_group_v2.firewall_group_1 c9e39fb2-ce20-46c8-a964-25f3898c7a97 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_fw_policy_v2",
			Category:         "Network ACL",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "firewall.svg",
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
					Description: `See Argument Reference above. ## Import Firewall Policies can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_fw_policy_v2.policy_1 07f422e6-c596-474b-8b94-fe2c12506ce0 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `See Argument Reference above. ## Import Firewall Policies can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_fw_policy_v2.policy_1 07f422e6-c596-474b-8b94-fe2c12506ce0 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_fw_rule_v2",
			Category:         "Network ACL",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "firewall.svg",
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
					Description: `(Optional) The region in which to obtain the v2 networking client. A Compute client is needed to create a firewall rule. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new firewall rule.`,
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
					Description: `See Argument Reference above. ## Import Firewall Rules can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_fw_rule_v2.rule_1 8dbc0c28-e49c-463f-b712-5c5d1bbac327 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `See Argument Reference above. ## Import Firewall Rules can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_fw_rule_v2.rule_1 8dbc0c28-e49c-463f-b712-5c5d1bbac327 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_identity_agency_v3",
			Category:         "Identity and Access Management (IAM)",
			ShortDescription: ``,
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
					Description: `(Required) Specifies the name of agency. The name is a string of 1 to 64 characters. Changing this will create a new agency.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Specifies the supplementary information about the agency. The value is a string of 0 to 255 characters.`,
				},
				resource.Attribute{
					Name:        "delegated_domain_name",
					Description: `(Optional) Specifies the name of delegated user domain. This parameter and ` + "`" + `delegated_service_name` + "`" + ` are alternative.`,
				},
				resource.Attribute{
					Name:        "delegated_service_name",
					Description: `(Optional) Specifies the name of delegated cloud service. This parameter and ` + "`" + `delegated_domain_name` + "`" + ` are alternative.`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `(Optional) Specifies the validity period of an agency. The valid value are`,
				},
				resource.Attribute{
					Name:        "project_role",
					Description: `(Optional) Specifies an array of one or more roles and projects which are used to grant permissions to agency on project. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "domain_roles",
					Description: `(optional) Specifies an array of one or more role names which stand for the permissionis to be granted to agency on domain. The ` + "`" + `project_role` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Required) Specifies the name of project.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `(Required) Specifies an array of role names. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The agency ID.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `The expiration time of agency.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time when the agency was created. ## Import Agencies can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_identity_agency_v3.agency 0b97661f9900f23f4fc2c00971ea4dc0 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The agency ID.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `The expiration time of agency.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time when the agency was created. ## Import Agencies can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_identity_agency_v3.agency 0b97661f9900f23f4fc2c00971ea4dc0 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_identity_group_membership_v3",
			Category:         "Identity and Access Management (IAM)",
			ShortDescription: ``,
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
			Type:             "flexibleengine_identity_group_v3",
			Category:         "Identity and Access Management (IAM)",
			ShortDescription: ``,
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
					Description: `(Optional) The domain this group belongs to. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `See Argument Reference above. ## Import Groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_identity_group_v3.group_1 89c60255-9bd6-460c-822a-e2b959ede9d2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "domain_id",
					Description: `See Argument Reference above. ## Import Groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_identity_group_v3.group_1 89c60255-9bd6-460c-822a-e2b959ede9d2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_identity_project_v3",
			Category:         "Identity and Access Management (IAM)",
			ShortDescription: ``,
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
					Description: `(Required) The name of the project. The length is less than or equal to 64 bytes. Name mut be prefixed with a valid region name (eg. eu-west-0_project_1).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the project. ## Atribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A resource ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `The parent of this project.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enabling status of this project. ## Import Projects can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_identity_project_v3.project_1 <ID> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_identity_provider",
			Category:         "Identity and Access Management (IAM)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"identity",
				"and",
				"access",
				"management",
				"iam",
				"provider",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required, String, ForceNew) Specifies the name of the identity provider to be registered. The maximum length is 64 characters. Only letters, digits, underscores (_), and hyphens (-) are allowed. The name is unique, it is recommended to include domain name information. Changing this creates a new resource.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required, String, ForceNew) Specifies the protocol of the identity provider. Valid values are`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional, Bool) Specifies the status for the identity provider. Defaults to true.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, String) Specifies the description of the identity provider.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional, String) Specifies the metadata of the IDP(Identity Provider) server. To obtain the metadata file of your enterprise IDP, contact the enterprise administrator. This field is used to import a metadata file to IAM to implement federated identity authentication. This field is required only if the protocol is set to`,
				},
				resource.Attribute{
					Name:        "openid_connect_config",
					Description: `(Optional, List) Specifies the description of the identity provider. This field is required only if the protocol is set to`,
				},
				resource.Attribute{
					Name:        "access_type",
					Description: `(Required, String) Specifies the access type of the identity provider. Available options are: + ` + "`" + `program` + "`" + `: programmatic access only. + ` + "`" + `program_console` + "`" + `: programmatic access and management console access.`,
				},
				resource.Attribute{
					Name:        "provider_url",
					Description: `(Required, String) Specifies the URL of the identity provider. This field corresponds to the iss field in the ID token.`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `(Required, String) Specifies the ID of a client registered with the OpenID Connect identity provider.`,
				},
				resource.Attribute{
					Name:        "signing_key",
					Description: `(Required, String) Public key used to sign the ID token of the OpenID Connect identity provider. This field is required only if the protocol is set to`,
				},
				resource.Attribute{
					Name:        "authorization_endpoint",
					Description: `(Optional, String) Specifies the authorization endpoint of the OpenID Connect identity provider. This field is required only if the access type is set to ` + "`" + `program_console` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scopes",
					Description: `(Optional, List) Specifies the scopes of authorization requests. It is an array of one or more scopes. Valid values are`,
				},
				resource.Attribute{
					Name:        "response_type",
					Description: `(Optional, String) Response type. Valid values is`,
				},
				resource.Attribute{
					Name:        "response_mode",
					Description: `(Optional, String) Response mode. Valid values is`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID which equals to the name.`,
				},
				resource.Attribute{
					Name:        "login_link",
					Description: `The login link of the identity provider.`,
				},
				resource.Attribute{
					Name:        "sso_type",
					Description: `The single sign-on type of the identity provider.`,
				},
				resource.Attribute{
					Name:        "conversion_rules",
					Description: `The identity conversion rules of the identity provider. The [object](#conversion_rules) structure is documented below <a name="conversion_rules"></a> The ` + "`" + `conversion_rules` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "local",
					Description: `The federated user information on the cloud platform.`,
				},
				resource.Attribute{
					Name:        "remote",
					Description: `The description of the identity provider. The ` + "`" + `local` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The name of a federated user on the cloud platform.`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `The user group to which the federated user belongs on the cloud platform. The ` + "`" + `remote` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "attribute",
					Description: `The attribute in the IDP assertion.`,
				},
				resource.Attribute{
					Name:        "condition",
					Description: `The condition of conversion rule.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The rule is matched only if the specified strings appear in the attribute type. ## Import Identity provider can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_identity_provider.provider_1 example_com_provider_saml ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID which equals to the name.`,
				},
				resource.Attribute{
					Name:        "login_link",
					Description: `The login link of the identity provider.`,
				},
				resource.Attribute{
					Name:        "sso_type",
					Description: `The single sign-on type of the identity provider.`,
				},
				resource.Attribute{
					Name:        "conversion_rules",
					Description: `The identity conversion rules of the identity provider. The [object](#conversion_rules) structure is documented below <a name="conversion_rules"></a> The ` + "`" + `conversion_rules` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "local",
					Description: `The federated user information on the cloud platform.`,
				},
				resource.Attribute{
					Name:        "remote",
					Description: `The description of the identity provider. The ` + "`" + `local` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The name of a federated user on the cloud platform.`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `The user group to which the federated user belongs on the cloud platform. The ` + "`" + `remote` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "attribute",
					Description: `The attribute in the IDP assertion.`,
				},
				resource.Attribute{
					Name:        "condition",
					Description: `The condition of conversion rule.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The rule is matched only if the specified strings appear in the attribute type. ## Import Identity provider can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_identity_provider.provider_1 example_com_provider_saml ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_identity_provider_conversion",
			Category:         "Identity and Access Management (IAM)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"identity",
				"and",
				"access",
				"management",
				"iam",
				"provider",
				"conversion",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "provider_id",
					Description: `(Required, String) The ID or name of the identity provider used to manage the conversion rules.`,
				},
				resource.Attribute{
					Name:        "conversion_rules",
					Description: `(Required, List) Specifies the identity conversion rules of the identity provider. You can use identity conversion rules to map the identities of existing users to FlexibleEngine and manage their access to cloud resources. The [object](#conversion_rules) structure is documented below. <a name="conversion_rules"></a> The ` + "`" + `conversion_rules` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "local",
					Description: `(Required, List) Specifies the federated user information on the cloud platform.`,
				},
				resource.Attribute{
					Name:        "remote",
					Description: `(Required, List) Specifies Federated user information in the IDP system. ->`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required, String) Specifies the name of a federated user on the cloud platform.`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `(Optional, String) Specifies the user group to which the federated user belongs on the cloud platform. The ` + "`" + `remote` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "attribute",
					Description: `(Required, String) Specifies the attribute in the IDP assertion.`,
				},
				resource.Attribute{
					Name:        "condition",
					Description: `(Optional, String) Specifies the condition of conversion rule. Available options are: + ` + "`" + `any_one_of` + "`" + `: The rule is matched only if the specified strings appear in the attribute type. + ` + "`" + `not_any_of` + "`" + `: The rule is matched only if the specified strings do not appear in the attribute type.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional, List) Specifies the rule is matched only if the specified strings appear in the attribute type. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of conversion rules. ## Import Identity provider conversion rules are imported using the ` + "`" + `provider_id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_identity_provider_conversion.conversion example_com_provider_oidc ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of conversion rules. ## Import Identity provider conversion rules are imported using the ` + "`" + `provider_id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_identity_provider_conversion.conversion example_com_provider_oidc ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_identity_role_assignment_v3",
			Category:         "Identity and Access Management (IAM)",
			ShortDescription: ``,
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
					Name:        "role_id",
					Description: `(Required) The role to assign.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required) The group to assign the role in.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Optional; Required if ` + "`" + `project_id` + "`" + ` is empty) The domain to assign the role in.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional; Required if ` + "`" + `domain_id` + "`" + ` is empty) The project to assign the role in. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "group_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_identity_role_v3",
			Category:         "Identity and Access Management (IAM)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"identity",
				"and",
				"access",
				"management",
				"iam",
				"role",
				"v3",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required, String) Name of the custom policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required, String) Description of the custom policy.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required, String) Display mode. Valid options are AX: Account level and XA: Project level.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Required, String) Document of the custom policy. - - - ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The role id.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `The account id.`,
				},
				resource.Attribute{
					Name:        "references",
					Description: `The number of references. ## Import Role can be imported using the following format: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_identity_role_v3.default {{ resource id}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The role id.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `The account id.`,
				},
				resource.Attribute{
					Name:        "references",
					Description: `The number of references. ## Import Role can be imported using the following format: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_identity_role_v3.default {{ resource id}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_identity_user_v3",
			Category:         "Identity and Access Management (IAM)",
			ShortDescription: ``,
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
					Description: `(Required, String) Specifies the name of the user. The user name consists of 5 to 32 characters. It can contain only uppercase letters, lowercase letters, digits, spaces, and special characters (-_) and cannot start with a digit.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, String) Specifies the description of the user.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Optional, String) Specifies the email address with a maximum of 255 characters.`,
				},
				resource.Attribute{
					Name:        "phone",
					Description: `(Optional, String) Specifies the mobile number with a maximum of 32 digits. This parameter must be used together with ` + "`" + `country_code` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "country_code",
					Description: `(Optional, String) Specifies the country code. The country code of the Chinese mainland is 0086. This parameter must be used together with ` + "`" + `phone` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional, String) Specifies the password for the user with 6 to 32 characters. It must contain at least two of the following character types: uppercase letters, lowercase letters, digits, and special characters.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional, Bool) Specifies whether the user is enabled or disabled. Valid values are ` + "`" + `true` + "`" + ` and ` + "`" + `false` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "password_strength",
					Description: `Indicates the password strength.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time when the IAM user was created.`,
				},
				resource.Attribute{
					Name:        "last_login",
					Description: `The tiem when the IAM user last login. ## Import Users can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_identity_user_v3.user_1 89c60255-9bd6-460c-822a-e2b959ede9d2 ` + "`" + `` + "`" + `` + "`" + ` But due to the security reason, ` + "`" + `password` + "`" + ` can not be imported, you can ignore it as below. ` + "`" + `` + "`" + `` + "`" + ` resource "flexibleengine_identity_user_v3" "user_1" { ... lifecycle { ignore_changes = [ "password", ] } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "password_strength",
					Description: `Indicates the password strength.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time when the IAM user was created.`,
				},
				resource.Attribute{
					Name:        "last_login",
					Description: `The tiem when the IAM user last login. ## Import Users can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_identity_user_v3.user_1 89c60255-9bd6-460c-822a-e2b959ede9d2 ` + "`" + `` + "`" + `` + "`" + ` But due to the security reason, ` + "`" + `password` + "`" + ` can not be imported, you can ignore it as below. ` + "`" + `` + "`" + `` + "`" + ` resource "flexibleengine_identity_user_v3" "user_1" { ... lifecycle { ignore_changes = [ "password", ] } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_images_image_v2",
			Category:         "Image Management Service (IMS)",
			ShortDescription: ``,
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
					Description: `(Required) The container format. Must be one of "ami", "ari", "aki", "bare", "ovf".`,
				},
				resource.Attribute{
					Name:        "disk_format",
					Description: `(Required) The disk format. Must be one of "ami", "ari", "aki", "vhd", "vmdk", "raw", "qcow2", "vdi", "iso".`,
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
					Description: `(Optional) This is the url of the raw image that will be downloaded in the ` + "`" + `image_cache_path` + "`" + ` before being uploaded to Glance. Glance is able to download image from internet but the ` + "`" + `gophercloud` + "`" + ` library does not yet provide a way to do so. Conflicts with ` + "`" + `local_file_path` + "`" + `.`,
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
					Description: `(Optional) The visibility of the image. Must be one of "public", "private", "community", or "shared". The ability to set the visibility depends upon the configuration of the FlexibleEngine cloud. Note: The ` + "`" + `properties` + "`" + ` attribute handling in the gophercloud library is currently buggy and needs to be fixed before being implemented in this resource. ## Attributes Reference The following attributes are exported:`,
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
					Description: `The id of the flexibleengine user who owns the image.`,
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
					Name:        "visibility",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date the image was created.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The date the image was last updated. ## Import Images can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_images_image_v2.rancheros 89c60255-9bd6-460c-822a-e2b959ede9d2 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `The id of the flexibleengine user who owns the image.`,
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
					Name:        "visibility",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date the image was created.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The date the image was last updated. ## Import Images can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_images_image_v2.rancheros 89c60255-9bd6-460c-822a-e2b959ede9d2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_kms_key_v1",
			Category:         "Key Management Service (KMS)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "Security-KMS.svg",
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
					Description: `(Required) Specifies the name of a KMS key.`,
				},
				resource.Attribute{
					Name:        "key_description",
					Description: `(Optional) Specifies the description of a KMS key.`,
				},
				resource.Attribute{
					Name:        "realm",
					Description: `(Optional) Region where a key resides. Changing this creates a new key.`,
				},
				resource.Attribute{
					Name:        "pending_days",
					Description: `(Optional) Specifies the duration in days after which the key is deleted after destruction of the resource, must be between 7 and 1096 days. Defaults to 7. It only be used when delete a key.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `(Optional) Specifies whether the key is enabled. Defaults to true.`,
				},
				resource.Attribute{
					Name:        "rotation_enabled",
					Description: `(Optional) Specifies whether the key rotation is enabled. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "rotation_interval",
					Description: `(Optional) Specifies the key rotation interval. The valid value is range from 30 to 365, defaults to 365. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The globally unique identifier for the key.`,
				},
				resource.Attribute{
					Name:        "default_key_flag",
					Description: `Identification of a Master Key. The value 1 indicates a Default Master Key, and the value 0 indicates a key.`,
				},
				resource.Attribute{
					Name:        "origin",
					Description: `Origin of a key. The default value is kms.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `ID of a user domain for the key.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `Creation time (time stamp) of a key.`,
				},
				resource.Attribute{
					Name:        "rotation_number",
					Description: `The total number of key rotations. ## Import KMS Keys can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_kms_key_v1.key_1 7056d636-ac60-4663-8a6c-82d3c32c1c64 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The globally unique identifier for the key.`,
				},
				resource.Attribute{
					Name:        "default_key_flag",
					Description: `Identification of a Master Key. The value 1 indicates a Default Master Key, and the value 0 indicates a key.`,
				},
				resource.Attribute{
					Name:        "origin",
					Description: `Origin of a key. The default value is kms.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `ID of a user domain for the key.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `Creation time (time stamp) of a key.`,
				},
				resource.Attribute{
					Name:        "rotation_number",
					Description: `The total number of key rotations. ## Import KMS Keys can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_kms_key_v1.key_1 7056d636-ac60-4663-8a6c-82d3c32c1c64 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_lb_certificate_v2",
			Category:         "Deprecated",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"deprecated",
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
					Name:        "private_key",
					Description: `(Required) The private encrypted key of the Certificate, PEM format.`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(Required) The public encrypted key of the Certificate, PEM format.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-readable description for the Certificate.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional) The domain of the Certificate. ## Attributes Reference The following attributes are exported:`,
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
			Type:             "flexibleengine_lb_l7policy_v2",
			Category:         "Elastic Load Balance (ELB)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"load",
				"balance",
				"elb",
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
					Description: `(Optional) The administrative state of the L7 Policy. This value can only be true (UP).`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The UUID of the tenant who owns the L7 Policy. Only administrative users can specify a tenant UUID other than their own. Changing this creates a new L7 Policy. ## Attributes Reference The following attributes are exported:`,
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
					Description: `See Argument Reference above. ## Import Load Balancer L7 Policy can be imported using the L7 Policy ID, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_lb_l7policy_v2.l7policy_1 8a7a79c2-cf17-4e65-b2ae-ddc8bfcf6c74 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `See Argument Reference above. ## Import Load Balancer L7 Policy can be imported using the L7 Policy ID, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_lb_l7policy_v2.l7policy_1 8a7a79c2-cf17-4e65-b2ae-ddc8bfcf6c74 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_lb_l7rule_v2",
			Category:         "Elastic Load Balance (ELB)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"load",
				"balance",
				"elb",
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
					Description: `(Optional) The administrative state of the L7 Rule. The value can only be true (UP).`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The UUID of the tenant who owns the L7 Rule. Only administrative users can specify a tenant UUID other than their own. Changing this creates a new L7 Rule. ## Attributes Reference The following attributes are exported:`,
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
					Description: `The ID of the Listener owning this resource. ## Import Load Balancer L7 Rule can be imported using the L7 Policy ID and L7 Rule ID separated by a slash, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_lb_l7rule_v2.l7rule_1 e0bd694a-abbe-450e-b329-0931fd1cc5eb/4086b0c9-b18c-4d1c-b6b8-4c56c3ad2a9e ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `The ID of the Listener owning this resource. ## Import Load Balancer L7 Rule can be imported using the L7 Policy ID and L7 Rule ID separated by a slash, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_lb_l7rule_v2.l7rule_1 e0bd694a-abbe-450e-b329-0931fd1cc5eb/4086b0c9-b18c-4d1c-b6b8-4c56c3ad2a9e ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_lb_listener_v2",
			Category:         "Elastic Load Balance (ELB)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"load",
				"balance",
				"elb",
				"lb",
				"listener",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to create the listener resource. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new listener.`,
				},
				resource.Attribute{
					Name:        "loadbalancer_id",
					Description: `(Required) The load balancer on which to provision this listener. Changing this creates a new listener.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The protocol - can either be TCP, UDP, HTTP or TERMINATED_HTTPS. Changing this creates a new listener.`,
				},
				resource.Attribute{
					Name:        "protocol_port",
					Description: `(Required) The port on which to listen for client traffic. Changing this creates a new listener.`,
				},
				resource.Attribute{
					Name:        "default_pool_id",
					Description: `(Optional) The ID of the default pool with which the listener is associated. Changing this creates a new listener.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Human-readable name for the listener. Does not have to be unique.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-readable description for the listener.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value pairs to associate with the listener.`,
				},
				resource.Attribute{
					Name:        "http2_enable",
					Description: `(Optional, Bool) Specifies whether to use HTTP/2. The default value is false. This parameter is valid only when the protocol is set to`,
				},
				resource.Attribute{
					Name:        "transparent_client_ip_enable",
					Description: `(Optional, Bool) Specifies whether to pass source IP addresses of the clients to backend servers. + For TCP and UDP listeners, the value can be true or false, and the default value is false. + For HTTP and HTTPS listeners, the value can only be true.`,
				},
				resource.Attribute{
					Name:        "idle_timeout",
					Description: `(Optional, Int) Specifies the idle timeout duration, in seconds. + For TCP listeners, the value ranges from 10 to 4000, and the default value is 300. + For HTTP and HTTPS listeners, the value ranges from 1 to 300, and the default value is 60. + For UDP listeners, this parameter does not take effect.`,
				},
				resource.Attribute{
					Name:        "request_timeout",
					Description: `(Optional, Int) Specifies the timeout duration for waiting for a request from a client, in seconds. This parameter is available only for HTTP and HTTPS listeners. The value ranges from 1 to 300, and the default value is 60.`,
				},
				resource.Attribute{
					Name:        "response_timeout",
					Description: `(Optional, Int) Specifies the timeout duration for waiting for a request from a backend server, in seconds. This parameter is available only for HTTP and HTTPS listeners. The value ranges from 1 to 300, and the default value is 60.`,
				},
				resource.Attribute{
					Name:        "default_tls_container_ref",
					Description: `(Optional) A reference to a Barbican Secrets container which stores TLS information. This is required if the protocol is ` + "`" + `TERMINATED_HTTPS` + "`" + `. See [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer) for more information.`,
				},
				resource.Attribute{
					Name:        "sni_container_refs",
					Description: `(Optional) A list of references to Barbican Secrets containers which store SNI information. See [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer) for more information.`,
				},
				resource.Attribute{
					Name:        "tls_ciphers_policy",
					Description: `(Optional) Specifies the security policy used by the listener. This parameter is valid only when the load balancer protocol is set to TERMINATED_HTTPS. The value can be tls-1-0, tls-1-1, tls-1-2, or tls-1-2-strict, and the default value is tls-1-0. For details of cipher suites for each security policy, see the table below. <table> <tr> <th>Security Policy</th> <th>TLS Version</th> <th>Cipher Suite</th> </tr > <tr > <td>tls-1-0</td> <td>TLSv1.2 TLSv1.1 TLSv1</td> <td rowspan="3">ECDHE-RSA-AES256-GCM-SHA384:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-ECDSA-AES128-GCM-SHA256:AES128-GCM-SHA256:AES256-GCM-SHA384:ECDHE-ECDSA-AES128-SHA256:ECDHE-RSA-AES128-SHA256:AES128-SHA256:AES256-SHA256:ECDHE-ECDSA-AES256-SHA384:ECDHE-RSA-AES256-SHA384:ECDHE-ECDSA-AES128-SHA:ECDHE-RSA-AES128-SHA:ECDHE-RSA-AES256-SHA:ECDHE-ECDSA-AES256-SHA:AES128-SHA:AES256-SHA</td> </tr> <tr> <td>tls-1-1</td> <td>TLSv1.2 TLSv1.1</td> </tr> <tr> <td>tls-1-2</td> <td>TLSv1.2</td> </tr> <tr> <td >tls-1-2-strict</td> <td >TLSv1.2</td> <td >ECDHE-RSA-AES256-GCM-SHA384:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-ECDSA-AES128-GCM-SHA256:AES128-GCM-SHA256:AES256-GCM-SHA384:ECDHE-ECDSA-AES128-SHA256:ECDHE-RSA-AES128-SHA256:AES128-SHA256:AES256-SHA256:ECDHE-ECDSA-AES256-SHA384:ECDHE-RSA-AES256-SHA384</td> </tr> </table> ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the listener.`,
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
					Name:        "tls_ciphers_policy",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the listener.`,
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
					Name:        "tls_ciphers_policy",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_lb_listener_v3",
			Category:         "Elastic Load Balance (ELB)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"load",
				"balance",
				"elb",
				"lb",
				"listener",
				"v3",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String, ForceNew) The region in which to create the listener resource. If omitted, the provider-level region will be used. Changing this creates a new listener.`,
				},
				resource.Attribute{
					Name:        "loadbalancer_id",
					Description: `(Required, String, ForceNew) The load balancer on which to provision this listener. Changing this creates a new listener.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required, String, ForceNew) The protocol can either be`,
				},
				resource.Attribute{
					Name:        "protocol_port",
					Description: `(Required, Int, ForceNew) The port on which to listen for client traffic. Changing this creates a new listener.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, String) Human-readable name for the listener.`,
				},
				resource.Attribute{
					Name:        "default_pool_id",
					Description: `(Optional, String) The ID of the default pool with which the listener is associated. Changing this creates a new listener.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, String) Human-readable description for the listener.`,
				},
				resource.Attribute{
					Name:        "http2_enable",
					Description: `(Optional, Bool) Specifies whether to use HTTP/2. The default value is false. This parameter is valid only when the protocol is set to`,
				},
				resource.Attribute{
					Name:        "forward_eip",
					Description: `(Optional, Bool) Specifies whether transfer the load balancer EIP in the X-Forward-EIP header to backend servers. The default value is false. This parameter is valid only when the protocol is set to`,
				},
				resource.Attribute{
					Name:        "access_policy",
					Description: `(Optional, String) Specifies the access policy for the listener. Valid options are`,
				},
				resource.Attribute{
					Name:        "ip_group",
					Description: `(Optional, String) Specifies the ip group id for the listener.`,
				},
				resource.Attribute{
					Name:        "server_certificate",
					Description: `(Optional, String) Specifies the ID of the server certificate used by the listener. This parameter is mandatory when protocol is set to`,
				},
				resource.Attribute{
					Name:        "sni_certificate",
					Description: `(Optional, List) Lists the IDs of SNI certificates (server certificates with a domain name) used by the listener. This parameter is valid when protocol is set to`,
				},
				resource.Attribute{
					Name:        "ca_certificate",
					Description: `(Optional, String) Specifies the ID of the CA certificate used by the listener. This parameter is valid when protocol is set to`,
				},
				resource.Attribute{
					Name:        "tls_ciphers_policy",
					Description: `(Optional, String) Specifies the TLS cipher policy for the listener. Valid options are: tls-1-0-inherit, tls-1-0, tls-1-1, tls-1-2, tls-1-2-strict, tls-1-2-fs, tls-1-0-with-1-3, and tls-1-2-fs-with-1-3. This parameter is valid when protocol is set to`,
				},
				resource.Attribute{
					Name:        "idle_timeout",
					Description: `(Optional, Int) Specifies the idle timeout for the listener. Value range: 0 to 4000.`,
				},
				resource.Attribute{
					Name:        "request_timeout",
					Description: `(Optional, Int) Specifies the request timeout for the listener. Value range: 1 to 300. This parameter is valid when protocol is set to`,
				},
				resource.Attribute{
					Name:        "response_timeout",
					Description: `(Optional, Int) Specifies the response timeout for the listener. Value range: 1 to 300. This parameter is valid when protocol is set to`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Map) The key/value pairs to associate with the listener. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the listener. ## Timeouts This resource provides the following timeouts configuration options:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `Default is 10 minute.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `Default is 10 minute.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `Default is 10 minute. ## Import ELB listener can be imported using the listener ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_lb_listener_v3.listener_1 5c20fdad-7288-11eb-b817-0255ac10158b ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the listener. ## Timeouts This resource provides the following timeouts configuration options:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `Default is 10 minute.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `Default is 10 minute.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `Default is 10 minute. ## Import ELB listener can be imported using the listener ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_lb_listener_v3.listener_1 5c20fdad-7288-11eb-b817-0255ac10158b ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_lb_loadbalancer_v2",
			Category:         "Elastic Load Balance (ELB)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "Network-ELB.svg",
			Keywords: []string{
				"elastic",
				"load",
				"balance",
				"elb",
				"lb",
				"loadbalancer",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to create the loadbalancer resource. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new loadbalancer.`,
				},
				resource.Attribute{
					Name:        "vip_subnet_id",
					Description: `(Required) The network on which to allocate the loadbalancer's address. A tenant can only create Loadbalancers on networks authorized by policy (e.g. networks that belong to them or networks that are shared). Changing this creates a new loadbalancer.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Human-readable name for the loadbalancer. Does not have to be unique.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-readable description for the loadbalancer.`,
				},
				resource.Attribute{
					Name:        "vip_address",
					Description: `(Optional) The ip address of the load balancer. Changing this creates a new loadbalancer.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value pairs to associate with the loadbalancer.`,
				},
				resource.Attribute{
					Name:        "loadbalancer_provider",
					Description: `(Optional) The name of the provider. Currently, only vlb is supported. Changing this creates a new loadbalancer.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Optional) A list of security group IDs to apply to the loadbalancer. The security groups must be specified by ID and not name (as opposed to how they are configured with the Compute Instance).`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) The administrative state of the loadbalancer. A valid value is true (UP) or false (DOWN).`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The UUID of the tenant who owns the loadbalancer. Only administrative users can specify a tenant UUID other than their own. Changing this creates a new loadbalancer. ## Attributes Reference The following attributes are exported:`,
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
					Name:        "tags",
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
					Description: `The Port ID of the Load Balancer IP. ## Import Loadbalancers can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_lb_loadbalancer_v2.loadbalancer_1 3e3632db-36c6-4b28-a92e-e72e6562daa6 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "tags",
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
					Description: `The Port ID of the Load Balancer IP. ## Import Loadbalancers can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_lb_loadbalancer_v2.loadbalancer_1 3e3632db-36c6-4b28-a92e-e72e6562daa6 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_lb_loadbalancer_v3",
			Category:         "Elastic Load Balance (ELB)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"load",
				"balance",
				"elb",
				"lb",
				"loadbalancer",
				"v3",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String, ForceNew) The region in which to create the loadbalancer resource. If omitted, the provider-level region will be used. Changing this creates a new loadbalancer.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required, List, ForceNew) Specifies the list of AZ names. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, String) Human-readable name for the loadbalancer.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, String) Human-readable description for the loadbalancer.`,
				},
				resource.Attribute{
					Name:        "cross_vpc_backend",
					Description: `(Optional, Bool) Enable this if you want to associate the IP addresses of backend servers with your load balancer. Can only be true when updating.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional, String, ForceNew) The vpc on which to create the loadbalancer. Changing this creates a new loadbalancer.`,
				},
				resource.Attribute{
					Name:        "ipv4_subnet_id",
					Description: `(Optional, String) The subnet on which to allocate the loadbalancer's ipv4 address.`,
				},
				resource.Attribute{
					Name:        "ipv6_network_id",
					Description: `(Optional, String) The network on which to allocate the loadbalancer's ipv6 address.`,
				},
				resource.Attribute{
					Name:        "ipv6_bandwidth_id",
					Description: `(Optional, String) The ipv6 bandwidth id. Only support shared bandwidth.`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `(Optional, String) The ipv4 address of the load balancer.`,
				},
				resource.Attribute{
					Name:        "ipv4_eip_id",
					Description: `(Optional, String, ForceNew) The ID of the EIP. Changing this parameter will create a new resource. ->`,
				},
				resource.Attribute{
					Name:        "iptype",
					Description: `(Optional, String, ForceNew) Elastic IP type. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "bandwidth_charge_mode",
					Description: `(Optional, String, ForceNew) Bandwidth billing type. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "sharetype",
					Description: `(Optional, String, ForceNew) Bandwidth sharing type. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "bandwidth_size",
					Description: `(Optional, Int, ForceNew) Bandwidth size. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "l4_flavor_id",
					Description: `(Optional, String) The L4 flavor id of the load balancer.`,
				},
				resource.Attribute{
					Name:        "l7_flavor_id",
					Description: `(Optional, String) The L7 flavor id of the load balancer.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Map) The key/value pairs to associate with the loadbalancer. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "ipv4_eip",
					Description: `The ipv4 eip address of the Load Balancer.`,
				},
				resource.Attribute{
					Name:        "ipv6_eip",
					Description: `The ipv6 eip address of the Load Balancer.`,
				},
				resource.Attribute{
					Name:        "ipv6_eip_id",
					Description: `The ipv6 eip id of the Load Balancer.`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `The ipv6 address of the Load Balancer. ## Timeouts This resource provides the following timeouts configuration options:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `Default is 10 minute.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `Default is 10 minute.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `Default is 5 minute. ## Import ELB loadbalancer can be imported using the loadbalancer ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_lb_loadbalancer_v3.loadbalancer_1 5c20fdad-7288-11eb-b817-0255ac10158b ` + "`" + `` + "`" + `` + "`" + ` Note that the imported state may not be identical to your resource definition, due to some attrubutes missing from the API response, security or some other reason. The missing attributes include: ` + "`" + `ipv6_bandwidth_id` + "`" + `, ` + "`" + `iptype` + "`" + `, ` + "`" + `bandwidth_charge_mode` + "`" + `, ` + "`" + `sharetype` + "`" + ` and ` + "`" + `bandwidth_size` + "`" + `. It is generally recommended running ` + "`" + `terraform plan` + "`" + ` after importing a loadbalancer. You can then decide if changes should be applied to the loadbalancer, or the resource definition should be updated to align with the loadbalancer. Also you can ignore changes as below. ` + "`" + `` + "`" + `` + "`" + ` resource "flexibleengine_lb_loadbalancer_v3" "loadbalancer_1" { ... lifecycle { ignore_changes = [ ipv6_bandwidth_id, iptype, bandwidth_charge_mode, sharetype, bandwidth_size, ] } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ipv4_eip",
					Description: `The ipv4 eip address of the Load Balancer.`,
				},
				resource.Attribute{
					Name:        "ipv6_eip",
					Description: `The ipv6 eip address of the Load Balancer.`,
				},
				resource.Attribute{
					Name:        "ipv6_eip_id",
					Description: `The ipv6 eip id of the Load Balancer.`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `The ipv6 address of the Load Balancer. ## Timeouts This resource provides the following timeouts configuration options:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `Default is 10 minute.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `Default is 10 minute.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `Default is 5 minute. ## Import ELB loadbalancer can be imported using the loadbalancer ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_lb_loadbalancer_v3.loadbalancer_1 5c20fdad-7288-11eb-b817-0255ac10158b ` + "`" + `` + "`" + `` + "`" + ` Note that the imported state may not be identical to your resource definition, due to some attrubutes missing from the API response, security or some other reason. The missing attributes include: ` + "`" + `ipv6_bandwidth_id` + "`" + `, ` + "`" + `iptype` + "`" + `, ` + "`" + `bandwidth_charge_mode` + "`" + `, ` + "`" + `sharetype` + "`" + ` and ` + "`" + `bandwidth_size` + "`" + `. It is generally recommended running ` + "`" + `terraform plan` + "`" + ` after importing a loadbalancer. You can then decide if changes should be applied to the loadbalancer, or the resource definition should be updated to align with the loadbalancer. Also you can ignore changes as below. ` + "`" + `` + "`" + `` + "`" + ` resource "flexibleengine_lb_loadbalancer_v3" "loadbalancer_1" { ... lifecycle { ignore_changes = [ ipv6_bandwidth_id, iptype, bandwidth_charge_mode, sharetype, bandwidth_size, ] } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_lb_member_v2",
			Category:         "Elastic Load Balance (ELB)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"load",
				"balance",
				"elb",
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
					Description: `(Optional) The administrative state of the member. A valid value is true (UP) or false (DOWN).`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The UUID of the tenant who owns the member. Only administrative users can specify a tenant UUID other than their own. Changing this creates a new member. ## Attributes Reference The following attributes are exported:`,
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
			Type:             "flexibleengine_lb_monitor_v2",
			Category:         "Elastic Load Balance (ELB)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"load",
				"balance",
				"elb",
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
					Description: `(Required) Number of permissible ping failures before changing the member's status to INACTIVE. Must be a number between 1 and 10..`,
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
					Name:        "port",
					Description: `(Optional) Specifies the health check port. The value ranges from 1 to 65536.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) The administrative state of the monitor. A valid value is true (UP) or false (DOWN).`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The UUID of the tenant who owns the monitor. Only administrative users can specify a tenant UUID other than their own. Changing this creates a new monitor. ## Attributes Reference The following attributes are exported:`,
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
				resource.Attribute{
					Name:        "port",
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
				resource.Attribute{
					Name:        "port",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_lb_pool_v2",
			Category:         "Elastic Load Balance (ELB)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"load",
				"balance",
				"elb",
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
					Name:        "tenant_id",
					Description: `(Optional) The UUID of the tenant who owns the pool. Only administrative users can specify a tenant UUID other than their own. Changing this creates a new pool.`,
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
			Type:             "flexibleengine_lb_whitelist_v2",
			Category:         "Elastic Load Balance (ELB)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"load",
				"balance",
				"elb",
				"lb",
				"whitelist",
				"v2",
			},
			Arguments: []resource.Attribute{
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
					Description: `(Optional) Specifies the IP addresses in the whitelist. Use commas(,) to separate the multiple IP addresses.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The UUID of the tenant who owns the whitelist. Only administrative users can specify a tenant UUID other than their own. Changing this creates a new whitelist. ## Attributes Reference The following attributes are exported:`,
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
			Type:             "flexibleengine_lts_group",
			Category:         "Log Tank Service (LTS)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"log",
				"tank",
				"service",
				"lts",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String, ForceNew) The region in which to create the log group resource. If omitted, the provider-level region will be used. Changing this creates a new log group resource.`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `(Required, String, ForceNew) Specifies the log group name. Changing this parameter will create a new resource. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The log group ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "ttl_in_days",
					Description: `Indicates the log expiration time. The value is fixed to 7 days. ## Import Log group can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `sh terraform import flexibleengine_lts_group.group_1 6e728c21-e3b6-11eb-b081-286ed488cb76 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The log group ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "ttl_in_days",
					Description: `Indicates the log expiration time. The value is fixed to 7 days. ## Import Log group can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `sh terraform import flexibleengine_lts_group.group_1 6e728c21-e3b6-11eb-b081-286ed488cb76 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_lts_topic",
			Category:         "Log Tank Service (LTS)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"log",
				"tank",
				"service",
				"lts",
				"topic",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String, ForceNew) The region in which to create the log topic resource. If omitted, the provider-level region will be used. Changing this creates a new log topic resource.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required, String, ForceNew) Specifies the ID of a created log group. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "topic_name",
					Description: `(Required, String, ForceNew) Specifies the log topic name. Changing this parameter will create a new resource. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The log topic ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "index_enabled",
					Description: `Indicates the search switch. When index is enabled, the topic allows you to search for logs by keyword. ## Import Log topic can be imported using the group ID and topic ID separated by a slash, e.g. ` + "`" + `` + "`" + `` + "`" + `sh terraform import flexibleengine_lts_topic.topic_1 393f2bfd-2244-11ea-adb7-286ed488c87f/137159d3-e3b7-11eb-b952-286ed488cb76 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The log topic ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "index_enabled",
					Description: `Indicates the search switch. When index is enabled, the topic allows you to search for logs by keyword. ## Import Log topic can be imported using the group ID and topic ID separated by a slash, e.g. ` + "`" + `` + "`" + `` + "`" + `sh terraform import flexibleengine_lts_topic.topic_1 393f2bfd-2244-11ea-adb7-286ed488c87f/137159d3-e3b7-11eb-b952-286ed488cb76 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_mls_instance_v1",
			Category:         "Machine Learning Service (MLS)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "Data Analysis&AI-MLS.svg",
			Keywords: []string{
				"machine",
				"learning",
				"service",
				"mls",
				"instance",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to create the MLS instance. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the MLS instance name. The DB instance name of the same type is unique in the same tenant. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) Specifies MLS Software version, only ` + "`" + `1.2.0` + "`" + ` is supported now. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Required) Specifies the instance network information. The structure is described below. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "agency",
					Description: `(Optional) Specifies the agency name. This parameter is mandatory only when you bind an instance to an elastic IP address (EIP). An instance must be bound to an EIP to grant MLS rights to abtain a tenant's token. NOTE: The tenant must create an agency on the Identity and Access Management (IAM) interface in advance. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `(Required) Specifies the instance flavor, only ` + "`" + `mls.c2.2xlarge.common` + "`" + ` is supported now. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "mrs_cluster",
					Description: `(Required) Specifies the MRS cluster information which the instance is associated. The structure is described below. NOTE: The current MRS instance requires an MRS cluster whose version is 1.3.0 and that is configured with the Spark component. MRS clusters whose version is not 1.3.0 or that are not configured with the Spark component cannot be selected. Changing this creates a new instance. The ` + "`" + `network` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) Specifies the ID of the virtual private cloud (VPC) where the instance resides. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) Specifies the ID of the subnet where the instance resides. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "security_group",
					Description: `(Optional) Specifies the ID of the security group of the instance. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "available_zone",
					Description: `(Required) Specifies the AZ of the instance. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `(Required) Specifies the IP address of the instance. The structure is described below. Changing this creates a new instance. The ` + "`" + `public_ip` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "bind_type",
					Description: `(Required) Specifies the bind type. Possible values: ` + "`" + `auto_assign` + "`" + ` and ` + "`" + `not_use` + "`" + `. Changing this creates a new instance. The ` + "`" + `mrs_cluster` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) Specifies the ID of the MRS cluster. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `(Optional) Specifies the MRS cluster username. This parameter is mandatory only when the MRS cluster is in the security mode. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "user_password",
					Description: `(Optional) Specifies the password of the MRS cluster user. The password and username work in a pair. Changing this creates a new instance. ## Attributes Reference The following attributes are exported:`,
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
					Name:        "version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "agency",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/vpc_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/security_group",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/available_zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/public_ip/bind_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/public_ip/eip_id",
					Description: `Indicates the EIP ID, This is returned only when bind_type is set to auto_assign.`,
				},
				resource.Attribute{
					Name:        "mrs_cluster",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates the MLS instance status.`,
				},
				resource.Attribute{
					Name:        "inner_endpoint",
					Description: `Indicates the URL for accessing the instance. Only machines in the same VPC and subnet as the instance can access the URL.`,
				},
				resource.Attribute{
					Name:        "public_endpoint",
					Description: `Indicates the URL for accessing the instance. The URL can be accessed from the Internet. The URL is created only after the instance is bound to an EIP.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `Indicates the creation time in the following format: yyyy-mm-dd Thh:mm:ssZ.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `Indicates the update time in the following format: yyyy-mm-dd Thh:mm:ssZ.`,
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
					Name:        "version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "agency",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/vpc_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/security_group",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/available_zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/public_ip/bind_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/public_ip/eip_id",
					Description: `Indicates the EIP ID, This is returned only when bind_type is set to auto_assign.`,
				},
				resource.Attribute{
					Name:        "mrs_cluster",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates the MLS instance status.`,
				},
				resource.Attribute{
					Name:        "inner_endpoint",
					Description: `Indicates the URL for accessing the instance. Only machines in the same VPC and subnet as the instance can access the URL.`,
				},
				resource.Attribute{
					Name:        "public_endpoint",
					Description: `Indicates the URL for accessing the instance. The URL can be accessed from the Internet. The URL is created only after the instance is bound to an EIP.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `Indicates the creation time in the following format: yyyy-mm-dd Thh:mm:ssZ.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `Indicates the update time in the following format: yyyy-mm-dd Thh:mm:ssZ.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_mrs_cluster_v1",
			Category:         "MapReduce Service (MRS)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "Data Analysis&AI-MRS.svg",
			Keywords: []string{
				"mapreduce",
				"service",
				"mrs",
				"cluster",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Cluster region information. Obtain the value from Regions and Endpoints.`,
				},
				resource.Attribute{
					Name:        "available_zone_id",
					Description: `(Required) ID or Name of an available zone. Obtain the value from Regions and Endpoints.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Required) Cluster name, which is globally unique and contains only 1 to 64 letters, digits, hyphens (-), and underscores (_).`,
				},
				resource.Attribute{
					Name:        "master_node_num",
					Description: `(Required) Number of Master nodes The value is 2.`,
				},
				resource.Attribute{
					Name:        "master_node_size",
					Description: `(Required) Best match based on several years of commissioning experience. MRS supports specifications of hosts, and host specifications are determined by CPUs, memory, and disks space. - Master nodes support s1.4xlarge and s1.8xlarge, c3.2xlarge.2, c3.xlarge.4, c3.2xlarge.4, c3.4xlarge.2, c3.4xlarge.4, c3.8xlarge.4, c3.15xlarge.4. - Core nodes of a streaming cluster support s1.xlarge, c2.2xlarge, s1.2xlarge, s1.4xlarge, s1.8xlarge, d1.8xlarge, , c3.2xlarge.2, c3.xlarge.4, c3.2xlarge.4, c3.4xlarge.2, c3.4xlarge.4, c3.8xlarge.4, c3.15xlarge.4. - Core nodes of an analysis cluster support all specifications c2.2xlarge, s1.xlarge, s1.4xlarge, s1.8xlarge, d1.xlarge, d1.2xlarge, d1.4xlarge, d1.8xlarge, , c3.2xlarge.2, c3.xlarge.4, c3.2xlarge.4, c3.4xlarge.2, c3.4xlarge.4, c3.8xlarge.4, c3.15xlarge.4, d2.xlarge.8, d2.2xlarge.8, d2.4xlarge.8, d2.8xlarge.8. The following provides specification details. node_size | CPU(core) | Memory(GB) | System Disk | Data Disk | --- | --- | --- | --- | --- | c2.2xlarge.linux.mrs | 8 | 16 | 40 | - cc3.xlarge.4.linux.mrs | 4 | 16 | 40 | - cc3.2xlarge.4.linux.mrs | 8 | 32 | 40 | - cc3.4xlarge.4.linux.mrs | 16 | 64 | 40 | - cc3.8xlarge.4.linux.mrs | 32 | 128 | 40 | - s1.xlarge.linux.mrs | 4 | 16 | 40 | - s1.4xlarge.linux.mrs | 16 | 64 | 40 | - s1.8xlarge.linux.mrs | 32 | 128 | 40 | - s3.xlarge.4.linux.mrs| 4 | 16 | 40 | - s3.2xlarge.4.linux.mrs| 8 | 32 | 40 | - s3.4xlarge.4.linux.mrs| 16 | 64 | 40 | - d1.xlarge.linux.mrs | 6 | 55 | 40 | 1.8 TB x 3 HDDs d1.2xlarge.linux.mrs | 12 | 110 | 40 | 1.8 TB x 6 HDDs d1.4xlarge.linux.mrs | 24 | 220 | 40 | 1.8 TB x 12 HDDs d1.8xlarge.linux.mrs | 48 | 440 | 40 | 1.8 TB x 24 HDDs d2.xlarge.linux.mrs | 4 | 32 | 40 | - d2.2xlarge.linux.mrs | 8 | 64 | 40 | - d2.4xlarge.linux.mrs | 16 | 128 | 40 | 1.8TB`,
				},
				resource.Attribute{
					Name:        "core_node_num",
					Description: `(Required) Number of Core nodes Value range: 3 to 500. A maximum of 500 Core nodes are supported by default. If more than 500 Core nodes are required, contact technical support engineers or invoke background APIs to modify the database.`,
				},
				resource.Attribute{
					Name:        "core_node_size",
					Description: `(Required) Instance specification of a Core node Configuration method of this parameter is identical to that of master_node_size.`,
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
					Description: `(Required) MRS cluster running mode. +`,
				},
				resource.Attribute{
					Name:        "cluster_admin_secret",
					Description: `(Required) Indicates the password of the MRS Manager administrator. + Must contain 8 to 32 characters. + Must contain at least three of the following: Lowercase letters, Uppercase letters, Digits and Special characters: ` + "`" + `~!@#$%^&`,
				},
				resource.Attribute{
					Name:        "cluster_version",
					Description: `(Optional) Version of the clusters. Possible values are as follows: MRS 1.8.9, MRS 2.0.1, MRS 2.1.0 and MRS 3.1.0-LTS.1. The latest version of MRS is used by default.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `(Optional) Type of clusters. 0: analysis cluster; 1: streaming cluster. The default value is 0.`,
				},
				resource.Attribute{
					Name:        "log_collection",
					Description: `(Optional) Indicates whether logs are collected when cluster installation fails. 0: not collected 1: collected The default value is 0. If log_collection is set to 1, OBS buckets will be created to collect the MRS logs. These buckets will be charged.`,
				},
				resource.Attribute{
					Name:        "component_list",
					Description: `(Required) Service component list. The object structure is documented below.`,
				},
				resource.Attribute{
					Name:        "add_jobs",
					Description: `(Optional) You can submit a job when you create a cluster to save time and use MRS easily. Only one job can be added. The object structure is documented below. The ` + "`" + `component_list` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "component_name",
					Description: `(Required) the Component name. + MRS 3.1.0-LTS.1 supports the following components: - The analysis cluster contains the following components: Hadoop, Spark2x, HBase, Hive, Hue, HetuEngine, Loader, Flink, Oozie, ZooKeeper, Ranger, and Tez. - The streaming cluster contains the following components: Kafka, Flume, ZooKeeper, and Ranger. + MRS 2.0.1 supports the following components: - The analysis cluster contains the following components: Presto, Hadoop, Spark, HBase, Hive, Hue, Loader, and Tez - The streaming cluster contains the following components: Kafka, Storm, and Flume. + MRS 1.8.9 supports the following components: - The analysis cluster contains the following components: Presto, Hadoop, Spark, HBase, Opentsdb, Hive, Hue, Loader, and Flink - The streaming cluster contains the following components: Kafka, KafkaManager, Storm, and Flume. The ` + "`" + `add_jobs` + "`" + ` block supports:`,
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
					Description: `(Optional) SQL program path This parameter is needed by Spark Script and Hive Script jobs only and must meet the following requirements: Contains a maximum of 1023 characters, excluding special characters such as ;|&><'$. The address cannot be empty or full of spaces. Starts with / or s3a://. Ends with .sql. sql is case-insensitive. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in UUID format.`,
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
					Description: `Cluster status. Valid values include: starting, running, terminated, failed, abnormal, terminating, frozen, scaling-out and scaling-in.`,
				},
				resource.Attribute{
					Name:        "order_id",
					Description: `Order ID for creating clusters.`,
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
					Name:        "duration",
					Description: `Cluster subscription duration.`,
				},
				resource.Attribute{
					Name:        "charging_start_time",
					Description: `Time when charging starts.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `Remarks of a cluster.`,
				},
				resource.Attribute{
					Name:        "error_info",
					Description: `Error information.`,
				},
				resource.Attribute{
					Name:        "component_list",
					Description: `See Argument Reference below. The component_list attributes:`,
				},
				resource.Attribute{
					Name:        "component_id",
					Description: `Component ID. For example, component_id of Hadoop is MRS 3.1.0-LTS.1_001, MRS 2.1.0_001, MRS 2.0.1_001, and MRS 1.8.9_001.`,
				},
				resource.Attribute{
					Name:        "component_name",
					Description: `Component name.`,
				},
				resource.Attribute{
					Name:        "component_version",
					Description: `Component version.`,
				},
				resource.Attribute{
					Name:        "component_desc",
					Description: `Component description.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in UUID format.`,
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
					Description: `Cluster status. Valid values include: starting, running, terminated, failed, abnormal, terminating, frozen, scaling-out and scaling-in.`,
				},
				resource.Attribute{
					Name:        "order_id",
					Description: `Order ID for creating clusters.`,
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
					Name:        "duration",
					Description: `Cluster subscription duration.`,
				},
				resource.Attribute{
					Name:        "charging_start_time",
					Description: `Time when charging starts.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `Remarks of a cluster.`,
				},
				resource.Attribute{
					Name:        "error_info",
					Description: `Error information.`,
				},
				resource.Attribute{
					Name:        "component_list",
					Description: `See Argument Reference below. The component_list attributes:`,
				},
				resource.Attribute{
					Name:        "component_id",
					Description: `Component ID. For example, component_id of Hadoop is MRS 3.1.0-LTS.1_001, MRS 2.1.0_001, MRS 2.0.1_001, and MRS 1.8.9_001.`,
				},
				resource.Attribute{
					Name:        "component_name",
					Description: `Component name.`,
				},
				resource.Attribute{
					Name:        "component_version",
					Description: `Component version.`,
				},
				resource.Attribute{
					Name:        "component_desc",
					Description: `Component description.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_mrs_cluster_v2",
			Category:         "MapReduce Service (MRS)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"mapreduce",
				"service",
				"mrs",
				"cluster",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String, ForceNew) The region in which to create the MRS cluster resource. If omitted, the provider-level region will be used. Changing this will create a new MRS cluster resource.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required, String, ForceNew) Specifies the availability zone in which to create the cluster. Changing this will create a new MRS cluster resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, String, ForceNew) Specifies the name of the MRS cluster. The name can contain 2 to 64 characters, which may consist of letters, digits, underscores (_) and hyphens (-). Changing this will create a new MRS cluster resource.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required, String, ForceNew) Specifies the MRS cluster version. Currently, ` + "`" + `MRS 1.8.9` + "`" + `, ` + "`" + `MRS 2.0.1` + "`" + `, and ` + "`" + `MRS 3.1.0-LTS.1` + "`" + ` are supported. Changing this will create a new MRS cluster resource.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional, String, ForceNew) Specifies the type of the MRS cluster. The valid values are`,
				},
				resource.Attribute{
					Name:        "component_list",
					Description: `(Required, List, ForceNew) Specifies the list of component names. Changing this will create a new MRS cluster resource. The supported components are as follows: <table border="2"> <tr> <th>Cluster Version</th> <th>Cluster Type</th> <th>Components</th> </tr> <tr> <td rowspan="4">MRS 3.1.0-LTS.1</td> <td>analysis</td> <td>Hadoop, Spark2x, HBase, Hive, Hue, HetuEngine, Loader, Flink, Oozie, ZooKeeper, Ranger, and Tez</td> </tr> <tr> <td>streaming</td> <td>Kafka, Flume, ZooKeeper, and Ranger</td> </tr> <tr> <td>hybrid</td> <td>Hadoop, Spark2x, HBase, Hive, Hue, HetuEngine, Loader, Flink, Oozie, ZooKeeper, Ranger, Tez, Kafka, and Flume</td> </tr> <tr> <td>custom</td> <td>Hadoop, Spark2x, HBase, Hive, Hue, HetuEngine, Loader, Kafka, Flume, Flink, Oozie, ZooKeeper, Ranger, Tez, and ClickHouse</td> </tr> <tr> <td rowspan="2">MRS 2.0.1</td> <td>analysis</td> <td>Presto, Hadoop, Spark, HBase, Hive, Hue, Loader, and Tez</td> </tr> <tr> <td>streaming</td> <td>Kafka, Storm, and Flume</td> </tr> <tr> <td rowspan="2">MRS 1.8.9</td> <td>analysis</td> <td>Presto, Hadoop, Spark, HBase, Opentsdb, Hive, Hue, Loader, and Flink</td> </tr> <tr> <td>streaming</td> <td>Kafka, KafkaManager, Storm, and Flume</td> </tr> </table>`,
				},
				resource.Attribute{
					Name:        "master_nodes",
					Description: `(Required, List, ForceNew) Specifies a list of the informations about the master nodes in the MRS cluster. The nodes object structure of the ` + "`" + `master_nodes` + "`" + ` is documented below. Changing this will create a new MRS cluster resource.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, String, ForceNew) Specifies the ID of the VPC which bound to the MRS cluster. Changing this will create a new MRS cluster resource.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required, String, ForceNew) Specifies the network ID of a subnet which bound to the MRS cluster. Changing this will create a new MRS cluster resource.`,
				},
				resource.Attribute{
					Name:        "manager_admin_pwd",
					Description: `(Required, String, ForceNew) Specifies the administrator password, which is used to login to the cluster management page. The password can contain 8 to 26 charactors and cannot be the username or the username spelled backwards. The password must contain lowercase letters, uppercase letters, digits, spaces and the special characters: ` + "`" + `!?,.:-_{}[]@$^+=/` + "`" + `. Changing this will create a new MRS cluster resource.`,
				},
				resource.Attribute{
					Name:        "node_key_pair",
					Description: `(Required, String, ForceNew) Specifies the name of a key pair, which is used to login to the each nodes(ECSs). Changing this will create a new MRS cluster resource.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `(Optional, String, ForceNew) Specifies the EIP address which bound to the MRS cluster. The EIP must have been created and must be in the same region as the cluster. Changing this will create a new MRS cluster resource.`,
				},
				resource.Attribute{
					Name:        "eip_id",
					Description: `(Optional, String, ForceNew) Specifies the EIP ID which bound to the MRS cluster. The EIP must have been created and must be in the same region as the cluster. Changing this will create a new MRS cluster resource.`,
				},
				resource.Attribute{
					Name:        "log_collection",
					Description: `(Optional, Bool, ForceNew) Specifies whether logs are collected when cluster installation fails. Default to true. If ` + "`" + `log_collection` + "`" + ` set true, the OBS buckets will be created and only used to collect logs that record MRS cluster creation failures. Changing this will create a new MRS cluster resource.`,
				},
				resource.Attribute{
					Name:        "safe_mode",
					Description: `(Optional, Bool, ForceNew) Specifies whether the running mode of the MRS cluster is secure, default to true. + true: enable Kerberos authentication. + false: disable Kerberos authentication. Changing this will create a new MRS cluster resource.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Optional, List, ForceNew) Specifies an array of one or more security group ID to attach to the MRS cluster. If using the specified security group, the group need to open the specified port (9022) rules.`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `(Optional, List, ForceNew) Specifies the template used for node deployment when the cluster type is CUSTOM. + mgmt_control_combined_v2: template for jointly deploying the management and control nodes. The management and control roles are co-deployed on the Master node, and data instances are deployed in the same node group. This deployment mode applies to scenarios where the number of control nodes is less than 100, reducing costs. + mgmt_control_separated_v2: The management and control roles are deployed on different master nodes, and data instances are deployed in the same node group. This deployment mode is applicable to a cluster with 100 to 500 nodes and delivers better performance in high-concurrency load scenarios. + mgmt_control_data_separated_v2: The management role and control role are deployed on different Master nodes, and data instances are deployed in different node groups. This deployment mode is applicable to a cluster with more than 500 nodes. Components can be deployed separately, which can be used for a larger cluster scale.`,
				},
				resource.Attribute{
					Name:        "analysis_core_nodes",
					Description: `(Optional, List) Specifies a list of the informations about the analysis core nodes in the MRS cluster. The nodes object structure of the ` + "`" + `analysis_core_nodes` + "`" + ` is documented below.`,
				},
				resource.Attribute{
					Name:        "streaming_core_nodes",
					Description: `(Optional, List) Specifies a list of the informations about the streaming core nodes in the MRS cluster. The nodes object structure of the ` + "`" + `streaming_core_nodes` + "`" + ` is documented below.`,
				},
				resource.Attribute{
					Name:        "analysis_task_nodes",
					Description: `(Optional, List) Specifies a list of the informations about the analysis task nodes in the MRS cluster. The nodes object structure of the ` + "`" + `analysis_task_nodes` + "`" + ` is documented below.`,
				},
				resource.Attribute{
					Name:        "streaming_task_nodes",
					Description: `(Optional, List) Specifies a list of the informations about the streaming task nodes in the MRS cluster. The nodes object structure of the ` + "`" + `streaming_task_nodes` + "`" + ` is documented below.`,
				},
				resource.Attribute{
					Name:        "custom_nodes",
					Description: `(Optional, List) Specifies a list of the informations about the custom nodes in the MRS cluster. The nodes object structure of the ` + "`" + `custom_nodes` + "`" + ` is documented below. Unlike other nodes, it needs to specify group_name.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Map, ForceNew) Specifies the key/value pairs to associate with the cluster. The ` + "`" + `nodes` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `(Required, String, ForceNew) Specifies the instance specifications for each nodes in node group. Changing this will create a new MRS cluster resource.`,
				},
				resource.Attribute{
					Name:        "node_number",
					Description: `(Required, Int) Specifies the number of nodes for the node group. Only the core group and task group updations are allowed. The number of nodes after scaling cannot be less than the number of nodes originally created.`,
				},
				resource.Attribute{
					Name:        "root_volume_type",
					Description: `(Required, String, ForceNew) Specifies the system disk flavor of the nodes. Changing this will create a new MRS cluster resource.`,
				},
				resource.Attribute{
					Name:        "root_volume_size",
					Description: `(Required, Int, ForceNew) Specifies the system disk size of the nodes. Changing this will create a new MRS cluster resource.`,
				},
				resource.Attribute{
					Name:        "data_volume_count",
					Description: `(Required, Int, ForceNew) Specifies the data disk number of the nodes. The number configuration of each node are as follows: + master_nodes: 1. + analysis_core_nodes: minimum is one and the maximum is subject to the configuration of the corresponding flavor. + streaming_core_nodes: minimum is one and the maximum is subject to the configuration of the corresponding flavor. + analysis_task_nodes: minimum is zero and the maximum is subject to the configuration of the corresponding flavor. + streaming_task_nodes: minimum is zero and the maximum is subject to the configuration of the corresponding flavor. Changing this will create a new MRS cluster resource.`,
				},
				resource.Attribute{
					Name:        "data_volume_type",
					Description: `(Optional, String, ForceNew) Specifies the data disk flavor of the nodes. Required if ` + "`" + `data_volume_count` + "`" + ` is greater than zero. Changing this will create a new MRS cluster resource. The following disk types are supported: + ` + "`" + `SATA` + "`" + `: common I/O disk + ` + "`" + `SAS` + "`" + `: high I/O disk + ` + "`" + `SSD` + "`" + `: ultra-high I/O disk`,
				},
				resource.Attribute{
					Name:        "data_volume_size",
					Description: `(Optional, Int, ForceNew) Specifies the data disk size of the nodes,in GB. The value range is 10 to 32768. Required if ` + "`" + `data_volume_count` + "`" + ` is greater than zero. Changing this will create a new MRS cluster resource.`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `(Optional, String, ForceNew) Specifies the name of nodes for the node group. This argument is mandatory when the cluster type is CUSTOM.`,
				},
				resource.Attribute{
					Name:        "assigned_roles",
					Description: `(Optional, List, ForceNew) Specifies the roles deployed in a node group. This argument is mandatory when the cluster type is CUSTOM. Each character string represents a role expression.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The cluster ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "total_node_number",
					Description: `The total number of nodes deployed in the cluster.`,
				},
				resource.Attribute{
					Name:        "master_node_ip",
					Description: `The IP address of the master node.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The preferred private IP address of the master node.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The cluster state, which include: running, frozen, abnormal and failed.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The cluster creation time, in RFC-3339 format.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `The cluster update time, in RFC-3339 format.`,
				},
				resource.Attribute{
					Name:        "charging_start_time",
					Description: `The charging start time which is the start time of billing, in RFC-3339 format.`,
				},
				resource.Attribute{
					Name:        "node",
					Description: `all the nodes attributes: master_nodes/analysis_core_nodes/streaming_core_nodes/analysis_task_nodes /streaming_task_nodes.`,
				},
				resource.Attribute{
					Name:        "host_ips",
					Description: `The host list of this nodes group in the cluster. The ` + "`" + `components` + "`" + ` attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Component ID. For example, component_id of Hadoop is MRS 3.1.0-LTS.1_001, MRS 2.0.1_001, and MRS 1.8.9_001.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Component name.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Component version.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Component description. ## Timeouts This resource provides the following timeouts configuration options:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `Default is 60 minute.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `Default is 180 minute.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `Default is 40 minute. ## Import Clusters can be imported by their ` + "`" + `id` + "`" + `. For example, ` + "`" + `` + "`" + `` + "`" + ` terraform import flexibleengine_mrs_cluster_v2.test b11b407c-e604-4e8d-8bc4-92398320b847 ` + "`" + `` + "`" + `` + "`" + ` Note that the imported state may not be identical to your resource definition, due to some attrubutes missing from the API response, security or some other reason. The missing attributes include: ` + "`" + `manager_admin_pwd` + "`" + `, ` + "`" + `template_id` + "`" + ` and ` + "`" + `assigned_roles` + "`" + `. It is generally recommended running ` + "`" + `terraform plan` + "`" + ` after importing a cluster. You can then decide if changes should be applied to the cluster, or the resource definition should be updated to align with the cluster. Also you can ignore changes as below. ` + "`" + `` + "`" + `` + "`" + ` resource "flexibleengine_mrs_cluster_v2" "test" { ... lifecycle { ignore_changes = [ manager_admin_pwd, ] } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The cluster ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "total_node_number",
					Description: `The total number of nodes deployed in the cluster.`,
				},
				resource.Attribute{
					Name:        "master_node_ip",
					Description: `The IP address of the master node.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The preferred private IP address of the master node.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The cluster state, which include: running, frozen, abnormal and failed.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The cluster creation time, in RFC-3339 format.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `The cluster update time, in RFC-3339 format.`,
				},
				resource.Attribute{
					Name:        "charging_start_time",
					Description: `The charging start time which is the start time of billing, in RFC-3339 format.`,
				},
				resource.Attribute{
					Name:        "node",
					Description: `all the nodes attributes: master_nodes/analysis_core_nodes/streaming_core_nodes/analysis_task_nodes /streaming_task_nodes.`,
				},
				resource.Attribute{
					Name:        "host_ips",
					Description: `The host list of this nodes group in the cluster. The ` + "`" + `components` + "`" + ` attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Component ID. For example, component_id of Hadoop is MRS 3.1.0-LTS.1_001, MRS 2.0.1_001, and MRS 1.8.9_001.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Component name.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Component version.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Component description. ## Timeouts This resource provides the following timeouts configuration options:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `Default is 60 minute.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `Default is 180 minute.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `Default is 40 minute. ## Import Clusters can be imported by their ` + "`" + `id` + "`" + `. For example, ` + "`" + `` + "`" + `` + "`" + ` terraform import flexibleengine_mrs_cluster_v2.test b11b407c-e604-4e8d-8bc4-92398320b847 ` + "`" + `` + "`" + `` + "`" + ` Note that the imported state may not be identical to your resource definition, due to some attrubutes missing from the API response, security or some other reason. The missing attributes include: ` + "`" + `manager_admin_pwd` + "`" + `, ` + "`" + `template_id` + "`" + ` and ` + "`" + `assigned_roles` + "`" + `. It is generally recommended running ` + "`" + `terraform plan` + "`" + ` after importing a cluster. You can then decide if changes should be applied to the cluster, or the resource definition should be updated to align with the cluster. Also you can ignore changes as below. ` + "`" + `` + "`" + `` + "`" + ` resource "flexibleengine_mrs_cluster_v2" "test" { ... lifecycle { ignore_changes = [ manager_admin_pwd, ] } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_mrs_hybrid_cluster_v1",
			Category:         "MapReduce Service (MRS)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"mapreduce",
				"service",
				"mrs",
				"hybrid",
				"cluster",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Cluster region information. Obtain the value from Regions and Endpoints.`,
				},
				resource.Attribute{
					Name:        "available_zone",
					Description: `(Required) ID or Name of an available zone. Obtain the value from Regions and Endpoints.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Required) Cluster name, which is globally unique and contains only 1 to 64 letters, digits, hyphens (-), and underscores (_).`,
				},
				resource.Attribute{
					Name:        "cluster_version",
					Description: `(Required) Version of the clusters. Possible values are as follows: MRS 1.8.9, MRS 2.0.1, MRS 2.1.0 and MRS 3.1.0-LTS.1.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) Specifies the id of the VPC.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) Specifies the id of the subnet.`,
				},
				resource.Attribute{
					Name:        "safe_mode",
					Description: `(Optional) MRS cluster running mode - 0: common mode The value indicates that the Kerberos authentication is disabled. Users can use all functions provided by the cluster. - 1: safe mode (by default) The value indicates that the Kerberos authentication is enabled. Common users cannot use the file management or job management functions of an MRS cluster and cannot view cluster resource usage or the job records of Hadoop and Spark. To use these functions, the users must obtain the relevant permissions from the MRS Manager administrator. The request has the cluster_admin_secret parameter only when safe_mode is set to 1.`,
				},
				resource.Attribute{
					Name:        "cluster_admin_secret",
					Description: `(Required) Indicates the password of the MRS Manager administrator. - Must contain 8 to 32 characters. - Must contain at least three types of the following: Lowercase letters, Uppercase letters, Digits, Special characters of ` + "`" + `~!@#$%^&`,
				},
				resource.Attribute{
					Name:        "master_node_key_pair",
					Description: `(Required) Name of a key pair You can use a key to log in to the Master node in the cluster.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Optional) Specifies the id of the security group which the cluster belongs to. If this parameter is empty, MRS automatically creates a security group, whose name starts with mrs_{cluster_name}.`,
				},
				resource.Attribute{
					Name:        "log_collection",
					Description: `(Optional) Indicates whether logs are collected when cluster installation fails. 0: not collected 1: collected The default value is 0. If log_collection is set to 1, OBS buckets will be created to collect the MRS logs. These buckets will be charged.`,
				},
				resource.Attribute{
					Name:        "component_list",
					Description: `(Required) Component name - Presto, Hadoop, Spark, HBase, Hive, Tez, Hue, Loader, Flume, Kafka and Storm are supported by MRS 2.0.1 or later. - Presto, Hadoop, Spark, HBase, Opentsdb, Hive, Hue, Loader, Flink, Flume, Kafka, KafkaManager and Storm are supported by MRS 1.8.9. - Hadoop, Spark, HBase, Hive, Hue, Loader, Flume, Kafka and Storm are supported by versions earlier than MRS 1.8.9.`,
				},
				resource.Attribute{
					Name:        "master_nodes",
					Description: `(Required) Specifies the master nodes information.`,
				},
				resource.Attribute{
					Name:        "analysis_core_nodes",
					Description: `(Required) Specifies the analysis core nodes information.`,
				},
				resource.Attribute{
					Name:        "streaming_core_nodes",
					Description: `(Required) Specifies the streaming core nodes information.`,
				},
				resource.Attribute{
					Name:        "analysis_task_nodes",
					Description: `(Optional) Specifies the analysis task nodes information.`,
				},
				resource.Attribute{
					Name:        "streaming_task_nodes",
					Description: `(Optional) Specifies the streaming task nodes information. The ` + "`" + `master_nodes` + "`" + `, ` + "`" + `analysis_core_nodes` + "`" + `, ` + "`" + `streaming_core_nodes` + "`" + `, ` + "`" + `analysis_task_nodes` + "`" + `, ` + "`" + `streaming_task_nodes` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `(Required) Best match based on several years of commissioning experience. MRS supports specifications of hosts, and host specifications are determined by CPUs, memory, and disks space. - Master nodes support s1.4xlarge and s1.8xlarge, c3.2xlarge.2, c3.xlarge.4, c3.2xlarge.4, c3.4xlarge.2, c3.4xlarge.4, c3.8xlarge.4, c3.15xlarge.4. - Core nodes of a streaming cluster support s1.xlarge, c2.2xlarge, s1.2xlarge, s1.4xlarge, s1.8xlarge, d1.8xlarge, , c3.2xlarge.2, c3.xlarge.4, c3.2xlarge.4, c3.4xlarge.2, c3.4xlarge.4, c3.8xlarge.4, c3.15xlarge.4. - Core nodes of an analysis cluster support all specifications c2.2xlarge, s1.xlarge, s1.4xlarge, s1.8xlarge, d1.xlarge, d1.2xlarge, d1.4xlarge, d1.8xlarge, , c3.2xlarge.2, c3.xlarge.4, c3.2xlarge.4, c3.4xlarge.2, c3.4xlarge.4, c3.8xlarge.4, c3.15xlarge.4, d2.xlarge.8, d2.2xlarge.8, d2.4xlarge.8, d2.8xlarge.8. The following provides specification details. node_size | CPU(core) | Memory(GB) | System Disk | Data Disk | --- | --- | --- | --- | --- | c2.2xlarge.linux.mrs | 8 | 16 | 40 | - cc3.xlarge.4.linux.mrs | 4 | 16 | 40 | - cc3.2xlarge.4.linux.mrs | 8 | 32 | 40 | - cc3.4xlarge.4.linux.mrs | 16 | 64 | 40 | - cc3.8xlarge.4.linux.mrs | 32 | 128 | 40 | - s1.xlarge.linux.mrs | 4 | 16 | 40 | - s1.4xlarge.linux.mrs | 16 | 64 | 40 | - s1.8xlarge.linux.mrs | 32 | 128 | 40 | - s3.xlarge.4.linux.mrs| 4 | 16 | 40 | - s3.2xlarge.4.linux.mrs| 8 | 32 | 40 | - s3.4xlarge.4.linux.mrs| 16 | 64 | 40 | - d1.xlarge.linux.mrs | 6 | 55 | 40 | 1.8 TB x 3 HDDs d1.2xlarge.linux.mrs | 12 | 110 | 40 | 1.8 TB x 6 HDDs d1.4xlarge.linux.mrs | 24 | 220 | 40 | 1.8 TB x 12 HDDs d1.8xlarge.linux.mrs | 48 | 440 | 40 | 1.8 TB x 24 HDDs d2.xlarge.linux.mrs | 4 | 32 | 40 | - d2.2xlarge.linux.mrs | 8 | 64 | 40 | - d2.4xlarge.linux.mrs | 16 | 128 | 40 | 1.8TB`,
				},
				resource.Attribute{
					Name:        "node_number",
					Description: `(Required) Number of nodes. The value ranges from 0 to 500 and the default value is 0. The total number of Core and Task nodes cannot exceed 500.`,
				},
				resource.Attribute{
					Name:        "data_volume_type",
					Description: `(Required) Data disk storage type of the node, supporting SATA and SSD currently - SATA: common I/O - SSD: Ultrahigh-speed I/O`,
				},
				resource.Attribute{
					Name:        "data_volume_size",
					Description: `(Required) Data disk size of the node Value range: 100 GB to 32000 GB`,
				},
				resource.Attribute{
					Name:        "data_volume_count",
					Description: `(Required) Number of data disks of the node Value range: 0 to 10 ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "billing_type",
					Description: `The value is Metered, indicating on-demand payment.`,
				},
				resource.Attribute{
					Name:        "total_node_number",
					Description: `Total node number.`,
				},
				resource.Attribute{
					Name:        "master_node_ip",
					Description: `IP address of a Master node.`,
				},
				resource.Attribute{
					Name:        "internal_ip",
					Description: `Iternal IP address.`,
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
					Name:        "external_alternate_ip",
					Description: `Backup external IP address.`,
				},
				resource.Attribute{
					Name:        "vnc",
					Description: `URI address for remote login of the elastic cloud server.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Cluster creation fee, which is automatically calculated.`,
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
					Name:        "charging_start_time",
					Description: `Time when charging starts. The ` + "`" + `components` + "`" + ` attributes:`,
				},
				resource.Attribute{
					Name:        "component_id",
					Description: `Component ID. For example, component_id of Hadoop is MRS 3.1.0-LTS.1_001, MRS 2.1.0_001, MRS 2.0.1_001, and MRS 1.8.9_001.`,
				},
				resource.Attribute{
					Name:        "component_name",
					Description: `Component name.`,
				},
				resource.Attribute{
					Name:        "component_version",
					Description: `Component version.`,
				},
				resource.Attribute{
					Name:        "component_desc",
					Description: `Component description.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "billing_type",
					Description: `The value is Metered, indicating on-demand payment.`,
				},
				resource.Attribute{
					Name:        "total_node_number",
					Description: `Total node number.`,
				},
				resource.Attribute{
					Name:        "master_node_ip",
					Description: `IP address of a Master node.`,
				},
				resource.Attribute{
					Name:        "internal_ip",
					Description: `Iternal IP address.`,
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
					Name:        "external_alternate_ip",
					Description: `Backup external IP address.`,
				},
				resource.Attribute{
					Name:        "vnc",
					Description: `URI address for remote login of the elastic cloud server.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Cluster creation fee, which is automatically calculated.`,
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
					Name:        "charging_start_time",
					Description: `Time when charging starts. The ` + "`" + `components` + "`" + ` attributes:`,
				},
				resource.Attribute{
					Name:        "component_id",
					Description: `Component ID. For example, component_id of Hadoop is MRS 3.1.0-LTS.1_001, MRS 2.1.0_001, MRS 2.0.1_001, and MRS 1.8.9_001.`,
				},
				resource.Attribute{
					Name:        "component_name",
					Description: `Component name.`,
				},
				resource.Attribute{
					Name:        "component_version",
					Description: `Component version.`,
				},
				resource.Attribute{
					Name:        "component_desc",
					Description: `Component description.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_mrs_job_v1",
			Category:         "MapReduce Service (MRS)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "Data Analysis&AI-MRS.svg",
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
			Type:             "flexibleengine_mrs_job_v2",
			Category:         "MapReduce Service (MRS)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"mapreduce",
				"service",
				"mrs",
				"job",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String, ForceNew) Specifies the region in which to create the MRS job resource. If omitted, the provider-level region will be used. Changing this will create a new MRS job resource.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required, String, ForceNew) Specifies an ID of the MRS cluster to which the job belongs to. Changing this will create a new MRS job resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, String, ForceNew) Specifies the name of the MRS job. The name can contain 1 to 64 characters, which may consist of letters, digits, underscores (_) and hyphens (-). Changing this will create a new MRS job resource.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required, String, ForceNew) Specifies the job type. The valid values are`,
				},
				resource.Attribute{
					Name:        "program_path",
					Description: `(Optional, String, ForceNew) Specifies the .jar package path or .py file path for program execution. The parameter must meet the following requirements: + Contains a maximum of 1023 characters, excluding special characters such as ` + "`" + `;|&><'$` + "`" + `. + The address cannot be empty or full of spaces. + The program support OBS or DHFS to storage program file or package. For OBS, starts with (OBS:)`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `(Optional, String, ForceNew) Specifies the parameters for the MRS job. Add an at sign (@) before each parameter can prevent the parameters being saved in plaintext format. Each parameters are separated with spaces. This parameter can be set when ` + "`" + `type` + "`" + ` is __Flink__, __MRS__ or __SparkSubmit__. Changing this will create a new MRS job resource.`,
				},
				resource.Attribute{
					Name:        "program_parameters",
					Description: `(Optional, Map, ForceNew) Specifies the the key/value pairs of the program parameters, such as thread, memory, and vCPUs, are used to optimize resource usage and improve job execution performance. This parameter can be set when ` + "`" + `type` + "`" + ` is __Flink__, __SparkSubmit__, __SparkSql__, __SparkScript__, __HiveSql__ or __HiveScript__. Changing this will create a new MRS job resource.`,
				},
				resource.Attribute{
					Name:        "service_parameters",
					Description: `(Optional, Map, ForceNew) Specifies the key/value pairs used to modify service configuration. Parameter configurations of services are available on the Service Configuration tab page of MRS Manager. Changing this will create a new MRS job resource.`,
				},
				resource.Attribute{
					Name:        "sql",
					Description: `(Optional, String, ForceNew) Specifies the SQL command or file path. Only required if ` + "`" + `type` + "`" + ` is __HiveSql__ or __SparkSql__. Changing this will create a new MRS job resource. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the MRS job in UUID format.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the MRS job.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `The creation time of the MRS job.`,
				},
				resource.Attribute{
					Name:        "submit_time",
					Description: `The submission time of the MRS job.`,
				},
				resource.Attribute{
					Name:        "finish_time",
					Description: `The completion time of the MRS job. ## Timeouts This resource provides the following timeouts configuration options:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `Default is 20 minute. ## Import MRS jobs can be imported using their ` + "`" + `id` + "`" + ` and the IDs of the MRS cluster to which the job belongs, separated by a slash, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_mrs_job_v2.test <cluster_id>/<id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the MRS job in UUID format.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the MRS job.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `The creation time of the MRS job.`,
				},
				resource.Attribute{
					Name:        "submit_time",
					Description: `The submission time of the MRS job.`,
				},
				resource.Attribute{
					Name:        "finish_time",
					Description: `The completion time of the MRS job. ## Timeouts This resource provides the following timeouts configuration options:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `Default is 20 minute. ## Import MRS jobs can be imported using their ` + "`" + `id` + "`" + ` and the IDs of the MRS cluster to which the job belongs, separated by a slash, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_mrs_job_v2.test <cluster_id>/<id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_nat_dnat_rule_v2",
			Category:         "NAT Gateway (NAT)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "nat-gateway.svg",
			Keywords: []string{
				"nat",
				"gateway",
				"dnat",
				"rule",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "nat_gateway_id",
					Description: `(Required) Specifies the ID of the nat gateway this dnat rule belongs to. Changing this creates a new dnat rule.`,
				},
				resource.Attribute{
					Name:        "floating_ip_id",
					Description: `(Required) Specifies the ID of the floating IP address. Changing this creates a new dnat rule.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Specifies the protocol type. Currently, TCP, UDP, and ANY are supported. Changing this creates a new dnat rule.`,
				},
				resource.Attribute{
					Name:        "internal_service_port",
					Description: `(Required) Specifies the port used by ECSs or BMSs to provide services that are accessible from external systems. Changing this creates a new dnat rule.`,
				},
				resource.Attribute{
					Name:        "external_service_port",
					Description: `(Required) Specifies the port for providing services that are accessible from external systems. Changing this creates a new dnat rule.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `(Optional) Specifies the port ID of an ECS or a BMS. This parameter is mandatory in VPC scenario. Changing this creates a new dnat rule.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `(Optional) Specifies the private IP address of a user, for example, the IP address of a VPC for dedicated connection. This parameter is mandatory in Direct Connect scenario. Changing this creates a new dnat rule.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Specifies the description of the dnat rule. The value is a string of no more than 255 characters, and angle brackets (<>) are not allowed. Changing this creates a new dnat rule. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "floating_ip_address",
					Description: `The actual floating IP address.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `DNAT rule creation time.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `DNAT rule status. ## Import DNAT can be imported using the following format: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_nat_dnat_rule_v2.dnat_1 f4f783a7-b908-4215-b018-724960e5df4a ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "floating_ip_address",
					Description: `The actual floating IP address.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `DNAT rule creation time.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `DNAT rule status. ## Import DNAT can be imported using the following format: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_nat_dnat_rule_v2.dnat_1 f4f783a7-b908-4215-b018-724960e5df4a ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_nat_gateway_v2",
			Category:         "NAT Gateway (NAT)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "nat-gateway.svg",
			Keywords: []string{
				"nat",
				"gateway",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String, ForceNew) Specifies the region in which to create the Nat gateway resource. If omitted, the provider-level region will be used. Changing this creates a new nat gateway.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, String) Specifies the nat gateway name. The name can contain only digits, letters, underscores (_), and hyphens(-).`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `(Required, String) Specifies the nat gateway type. The value can be: + ` + "`" + `1` + "`" + `: small type, which supports up to 10,000 SNAT connections. + ` + "`" + `2` + "`" + `: medium type, which supports up to 50,000 SNAT connections. + ` + "`" + `3` + "`" + `: large type, which supports up to 200,000 SNAT connections. + ` + "`" + `4` + "`" + `: extra-large type, which supports up to 1,000,000 SNAT connections.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, String, ForceNew) Specifies the ID of the VPC this nat gateway belongs to. Changing this creates a new nat gateway.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required, String, ForceNew) Specifies the subnet ID of the downstream interface (the next hop of the DVR) of the NAT gateway. Changing this creates a new nat gateway.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, String) Specifies the description of the nat gateway. The value contains 0 to 255 characters, and angle brackets (<) and (>) are not allowed. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the nat gateway. ## Timeouts This resource provides the following timeouts configuration options:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `Default is 10 minute.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `Default is 10 minute. ## Import Nat gateway can be imported using the following format: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_nat_gateway_v2.nat_1 d126fb87-43ce-4867-a2ff-cf34af3765d9 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the nat gateway. ## Timeouts This resource provides the following timeouts configuration options:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `Default is 10 minute.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `Default is 10 minute. ## Import Nat gateway can be imported using the following format: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_nat_gateway_v2.nat_1 d126fb87-43ce-4867-a2ff-cf34af3765d9 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_nat_snat_rule_v2",
			Category:         "NAT Gateway (NAT)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "nat-gateway.svg",
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
					Name:        "floating_ip_id",
					Description: `(Required) ID of the floating ip this snat rule connets to. Changing this creates a new snat rule.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) ID of the subnet this snat rule connects to. This parameter and ` + "`" + `cidr` + "`" + ` are alternative. Changing this creates a new snat rule.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Optional) Specifies CIDR, which can be in the format of a network segment or a host IP address. This parameter and ` + "`" + `network_id` + "`" + ` are alternative. Changing this creates a new snat rule.`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `(Optional) Specifies the scenario. The valid value is 0 (VPC scenario) and 1 (Direct Connect scenario). Only ` + "`" + `cidr` + "`" + ` can be specified over a Direct Connect connection. If no value is entered, the default value 0 (VPC scenario) is used. Changing this creates a new snat rule. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "floating_ip_address",
					Description: `The actual floating IP address.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the snat rule. ## Timeouts This resource provides the following timeouts configuration options:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `Default is 10 minute.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `Default is 10 minute. ## Import SNAT rules can be imported using the following format: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_nat_snat_rule_v2.snat_1 9e0713cb-0a2f-484e-8c7d-daecbb61dbe4 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "floating_ip_address",
					Description: `The actual floating IP address.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the snat rule. ## Timeouts This resource provides the following timeouts configuration options:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `Default is 10 minute.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `Default is 10 minute. ## Import SNAT rules can be imported using the following format: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_nat_snat_rule_v2.snat_1 9e0713cb-0a2f-484e-8c7d-daecbb61dbe4 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_network_acl",
			Category:         "Network ACL",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"acl",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the network ACL name. This parameter can contain a maximum of 64 characters, which may consist of letters, digits, underscores (_), and hyphens (-).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Specifies the supplementary information about the network ACL. This parameter can contain a maximum of 255 characters and cannot contain angle brackets (< or >).`,
				},
				resource.Attribute{
					Name:        "inbound_rules",
					Description: `(Optional) A list of the IDs of ingress rules associated with the network ACL.`,
				},
				resource.Attribute{
					Name:        "outbound_rules",
					Description: `(Optional) A list of the IDs of egress rules associated with the network ACL.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `(Optional) A list of the IDs of networks associated with the network ACL. ## Attributes Reference All of the argument attributes are also exported as result attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the network ACL.`,
				},
				resource.Attribute{
					Name:        "inbound_policy_id",
					Description: `The ID of the ingress firewall policy for the network ACL.`,
				},
				resource.Attribute{
					Name:        "outbound_policy_id",
					Description: `The ID of the egress firewall policy for the network ACL.`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `A list of the port IDs of the subnet gateway.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the network ACL.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the network ACL.`,
				},
				resource.Attribute{
					Name:        "inbound_policy_id",
					Description: `The ID of the ingress firewall policy for the network ACL.`,
				},
				resource.Attribute{
					Name:        "outbound_policy_id",
					Description: `The ID of the egress firewall policy for the network ACL.`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `A list of the port IDs of the subnet gateway.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the network ACL.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_network_acl_rule",
			Category:         "Network ACL",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"acl",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Specifies a unique name for the network ACL rule.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Specifies the description for the network ACL rule.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Specifies the protocol supported by the network ACL rule. Valid values are:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) Specifies the action in the network ACL rule. Currently, the value can be`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `(Optional) Specifies the IP version, either 4 (default) or 6. This parameter is available after the IPv6 function is enabled.`,
				},
				resource.Attribute{
					Name:        "source_ip_address",
					Description: `(Optional) Specifies the source IP address that the traffic is allowed from. The default value is`,
				},
				resource.Attribute{
					Name:        "destination_ip_address",
					Description: `(Optional) Specifies the destination IP address to which the traffic is allowed. The default value is`,
				},
				resource.Attribute{
					Name:        "source_port",
					Description: `(Optional) Specifies the source port number or port number range. The value ranges from 1 to 65535. For a port number range, enter two port numbers connected by a hyphen (-). For example, 1-100.`,
				},
				resource.Attribute{
					Name:        "destination_port",
					Description: `(Optional) Specifies the destination port number or port number range. The value ranges from 1 to 65535. For a port number range, enter two port numbers connected by a hyphen (-). For example, 1-100.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Enabled status for the network ACL rule. Defaults to true. ## Attributes Reference The following attributes are exported:`,
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
					Description: `See Argument Reference above. ## Import network ACL rules can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_network_acl_rule.rule_1 89a84b28-4cc2-4859-9885-c67e802a46a3 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `See Argument Reference above. ## Import network ACL rules can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_network_acl_rule.rule_1 89a84b28-4cc2-4859-9885-c67e802a46a3 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_networking_floatingip_associate_v2",
			Category:         "Virtual Private Cloud (VPC)",
			ShortDescription: ``,
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
					Description: `See Argument Reference above. ## Import Floating IP associations can be imported using the ` + "`" + `id` + "`" + ` of the floating IP, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_networking_floatingip_associate_v2.fip 2c7f39f3-702b-48d1-940c-b50384177ee1 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `See Argument Reference above. ## Import Floating IP associations can be imported using the ` + "`" + `id` + "`" + ` of the floating IP, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_networking_floatingip_associate_v2.fip 2c7f39f3-702b-48d1-940c-b50384177ee1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_networking_floatingip_v2",
			Category:         "Virtual Private Cloud (VPC)",
			ShortDescription: ``,
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
					Description: `(Optional) The name of the pool from which to obtain the floating IP. Default value is admin_external_net. Changing this creates a new floating IP.`,
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
					Description: `The fixed IP which the floating IP maps to. ## Import Floating IPs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_networking_floatingip_v2.floatip_1 2c7f39f3-702b-48d1-940c-b50384177ee1 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `The fixed IP which the floating IP maps to. ## Import Floating IPs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_networking_floatingip_v2.floatip_1 2c7f39f3-702b-48d1-940c-b50384177ee1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_networking_network_v2",
			Category:         "Virtual Private Cloud (VPC)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "Network.svg",
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
					Description: `The physical network where this network is implemented.`,
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
					Description: `See Argument Reference above. ## Import Networks can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_networking_network_v2.network_1 d90ce693-5ccf-4136-a0ed-152ce412b6b9 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `See Argument Reference above. ## Import Networks can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_networking_network_v2.network_1 d90ce693-5ccf-4136-a0ed-152ce412b6b9 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_networking_port_v2",
			Category:         "Virtual Private Cloud (VPC)",
			ShortDescription: ``,
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
					Description: `(Optional) A list of security group IDs to apply to the port. The security groups must be specified by ID and not name (as opposed to how they are configured with the Compute Instance).`,
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
					Description: `(Optional) An array of IP/MAC Address pairs of additional IP addresses that can be active on this port. The structure is described below.`,
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
					Description: `(Optional) IP address desired in the subnet for this port. If you don't specify ` + "`" + `ip_address` + "`" + `, an available IP address from the specified subnet will be allocated to this port. The ` + "`" + `allowed_address_pairs` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Required) The additional IP address. The value can be an IP Address or a CIDR, and can not be`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `(Optional) The additional MAC address. ## Attributes Reference The following attributes are exported:`,
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
					Name:        "all fixed_ips",
					Description: `The collection of Fixed IP addresses on the port in the order returned by the Network v2 API. ## Import Ports can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_networking_port_v2.port_1 eae26a3e-1c33-4cc1-9c31-0cd729c438a1 ` + "`" + `` + "`" + `` + "`" + ` ## Notes ### Ports and Instances There are some notes to consider when connecting Instances to networks using Ports. Please see the ` + "`" + `flexibleengine_compute_instance_v2` + "`" + ` documentation for further documentation.`,
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
					Name:        "all fixed_ips",
					Description: `The collection of Fixed IP addresses on the port in the order returned by the Network v2 API. ## Import Ports can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_networking_port_v2.port_1 eae26a3e-1c33-4cc1-9c31-0cd729c438a1 ` + "`" + `` + "`" + `` + "`" + ` ## Notes ### Ports and Instances There are some notes to consider when connecting Instances to networks using Ports. Please see the ` + "`" + `flexibleengine_compute_instance_v2` + "`" + ` documentation for further documentation.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_networking_router_interface_v2",
			Category:         "Virtual Private Cloud (VPC)",
			ShortDescription: ``,
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
					Description: `See Argument Reference above.`,
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
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_networking_router_route_v2",
			Category:         "Virtual Private Cloud (VPC)",
			ShortDescription: ``,
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
					Description: `See Argument Reference above. ## Notes The ` + "`" + `next_hop` + "`" + ` IP address must be directly reachable from the router at the ` + "`" + `` + "`" + `flexibleengine_networking_router_route_v2` + "`" + `` + "`" + ` resource creation time. You can ensure that by explicitly specifying a dependency on the ` + "`" + `` + "`" + `flexibleengine_networking_router_interface_v2` + "`" + `` + "`" + ` resource that connects the next hop to the router, as in the example above.`,
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
					Description: `See Argument Reference above. ## Notes The ` + "`" + `next_hop` + "`" + ` IP address must be directly reachable from the router at the ` + "`" + `` + "`" + `flexibleengine_networking_router_route_v2` + "`" + `` + "`" + ` resource creation time. You can ensure that by explicitly specifying a dependency on the ` + "`" + `` + "`" + `flexibleengine_networking_router_interface_v2` + "`" + `` + "`" + ` resource that connects the next hop to the router, as in the example above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_networking_router_v2",
			Category:         "Virtual Private Cloud (VPC)",
			ShortDescription: ``,
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
					Name:        "external_gateway",
					Description: `(Optional) The network UUID of an external gateway for the router. A router with an external gateway is required if any compute instances or load balancers will be using floating IPs. Changing this updates the ` + "`" + `external_gateway` + "`" + ` of an existing router.`,
				},
				resource.Attribute{
					Name:        "enable_snat",
					Description: `(Optional) Enable Source NAT for the router. Valid values are "true" or "false". An ` + "`" + `external_gateway` + "`" + ` has to be set in order to set this property. Changing this updates the ` + "`" + `enable_snat` + "`" + ` of the router.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the floating IP. Required if admin wants to create a router for another tenant. Changing this creates a new router.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional driver-specific options. ## Attributes Reference The following attributes are exported:`,
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
					Name:        "external_gateway",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "enable_snat",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `See Argument Reference above.`,
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
					Name:        "external_gateway",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "enable_snat",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_networking_secgroup_rule_v2",
			Category:         "Virtual Private Cloud (VPC)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "security-group.svg",
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
					Name:        "security_group_id",
					Description: `(Required) The security group ID the rule should belong to. Changing this creates a new security group rule.`,
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
					Description: `(Optional) The remote group id, the value needs to be an FlexibleEngine ID of a security group in the same tenant. Changing this creates a new security group rule.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Specifies the supplementary information about the security group rule. This parameter can contain a maximum of 255 characters and cannot contain angle brackets (< or >). Changing this creates a new security group rule. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in UUID format. ## Import Security Group Rules can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_networking_secgroup_rule_v2.secgroup_rule_1 aeb68ee3-6e9d-4256-955c-9584a6212745 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in UUID format. ## Import Security Group Rules can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_networking_secgroup_rule_v2.secgroup_rule_1 aeb68ee3-6e9d-4256-955c-9584a6212745 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_networking_secgroup_v2",
			Category:         "Virtual Private Cloud (VPC)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "security-group.svg",
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
					Name:        "delete_default_rules",
					Description: `(Optional) Whether or not to delete the default egress security rules. This is ` + "`" + `false` + "`" + ` by default. See the below note for more information. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in UUID format. ## Default Security Group Rules In most cases, FlexibleEngine will create some egress security group rules for each new security group. These security group rules will not be managed by Terraform, so if you prefer to have`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in UUID format. ## Default Security Group Rules In most cases, FlexibleEngine will create some egress security group rules for each new security group. These security group rules will not be managed by Terraform, so if you prefer to have`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_networking_subnet_v2",
			Category:         "Virtual Private Cloud (VPC)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "subnet.svg",
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
					Description: `(Optional) The administrative state of the network. Acceptable values are "true" and "false". Changing this value enables or disables the DHCP capabilities of the existing subnet. Defaults to true.`,
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
					Description: `(Required) The ending addresss. The ` + "`" + `host_routes` + "`" + ` block supports:`,
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
					Description: `See Argument Reference above. ## Import Subnets can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_networking_subnet_v2.subnet_1 da4faf16-5546-41e4-8330-4d0002b74048 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `See Argument Reference above. ## Import Subnets can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_networking_subnet_v2.subnet_1 da4faf16-5546-41e4-8330-4d0002b74048 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_networking_vip_associate_v2",
			Category:         "Virtual Private Cloud (VPC)",
			ShortDescription: ``,
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
			Type:             "flexibleengine_networking_vip_v2",
			Category:         "Virtual Private Cloud (VPC)",
			ShortDescription: ``,
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
			Type:             "flexibleengine_obs_bucket",
			Category:         "Object Storage Service (OSS)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"object",
				"storage",
				"service",
				"oss",
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
					Description: `(Optional) Specifies the storage class of the bucket. OBS provides three storage classes: "STANDARD", "STANDARD_IA" (Infrequent Access) and "GLACIER" (Archive). Defaults to ` + "`" + `STANDARD` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `(Optional) Specifies the ACL policy for a bucket. The predefined common policies are as follows: "private", "public-read", "public-read-write" and "log-delivery-write". Defaults to ` + "`" + `private` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "versioning",
					Description: `(Optional) Whether enable versioning. Once you version-enable a bucket, it can never return to an unversioned state. You can, however, suspend versioning on that bucket.`,
				},
				resource.Attribute{
					Name:        "encryption",
					Description: `(Optional) Whether enable default server-side encryption of the bucket in SSE-KMS mode.`,
				},
				resource.Attribute{
					Name:        "kms_key_id",
					Description: `(Optional) Specifies the ID of a kms key. If omitted, the default master key will be used.`,
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
					Name:        "multi_az",
					Description: `(Optional) Whether enable the multi-AZ mode for the bucket. When the multi-AZ mode is enabled, data in the bucket is duplicated and stored in multiple AZs. Changing this creates a new bucket.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) If specified, the region this bucket should reside in. Otherwise, the region used by the provider. Changing this creates a new bucket. The ` + "`" + `logging` + "`" + ` object supports the following:`,
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
					Description: `(Optional) A JSON or XML format containing routing rules describing redirect behavior and when redirects are applied. Each rule contains a ` + "`" + `Condition` + "`" + ` and a ` + "`" + `Redirect` + "`" + ` as shown in the following table: Parameter | Key --- | --- Condition | KeyPrefixEquals, HttpErrorCodeReturnedEquals Redirect | Protocol, HostName, ReplaceKeyPrefixWith, ReplaceKeyWith, HttpRedirectCode The ` + "`" + `cors_rule` + "`" + ` object supports the following:`,
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
					Description: `(Optional) Specifies a period when objects that have been last updated are automatically transitioned to ` + "`" + `STANDARD_IA` + "`" + ` or ` + "`" + `GLACIER` + "`" + ` storage class (documented below).`,
				},
				resource.Attribute{
					Name:        "noncurrent_version_expiration",
					Description: `(Optional) Specifies a period when noncurrent object versions are automatically deleted. (documented below).`,
				},
				resource.Attribute{
					Name:        "noncurrent_version_transition",
					Description: `(Optional) Specifies a period when noncurrent object versions are automatically transitioned to ` + "`" + `STANDARD_IA` + "`" + ` or ` + "`" + `GLACIER` + "`" + ` storage class (documented below). At least one of ` + "`" + `expiration` + "`" + `, ` + "`" + `transition` + "`" + `, ` + "`" + `noncurrent_version_expiration` + "`" + `, ` + "`" + `noncurrent_version_transition` + "`" + ` must be specified. The ` + "`" + `expiration` + "`" + ` object supports the following`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `(Required) The class of storage used to store the object. Only "STANDARD_IA" and "GLACIER" are supported. The ` + "`" + `noncurrent_version_expiration` + "`" + ` object supports the following`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `(Required) The class of storage used to store the object. Only "STANDARD_IA" and "GLACIER" are supported. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The name of the bucket.`,
				},
				resource.Attribute{
					Name:        "bucket_domain_name",
					Description: `The bucket domain name. Will be of format ` + "`" + `bucketname.oss.region.prod-cloud-ocb.orange-business.com` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region this bucket resides in. ## Import OBS bucket can be imported using the ` + "`" + `bucket` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_obs_bucket.bucket bucket-name ` + "`" + `` + "`" + `` + "`" + ` Note that the imported state may not be identical to your resource definition, due to some attrubutes missing from the API response. The missing attributes include ` + "`" + `acl` + "`" + ` and ` + "`" + `force_destroy` + "`" + `. It is generally recommended running ` + "`" + `terraform plan` + "`" + ` after importing an OBS bucket. Also you can ignore changes as below. ` + "`" + `` + "`" + `` + "`" + ` resource "flexibleengine_obs_bucket" "bucket" { ... lifecycle { ignore_changes = [ acl, force_destroy, ] } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The name of the bucket.`,
				},
				resource.Attribute{
					Name:        "bucket_domain_name",
					Description: `The bucket domain name. Will be of format ` + "`" + `bucketname.oss.region.prod-cloud-ocb.orange-business.com` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region this bucket resides in. ## Import OBS bucket can be imported using the ` + "`" + `bucket` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_obs_bucket.bucket bucket-name ` + "`" + `` + "`" + `` + "`" + ` Note that the imported state may not be identical to your resource definition, due to some attrubutes missing from the API response. The missing attributes include ` + "`" + `acl` + "`" + ` and ` + "`" + `force_destroy` + "`" + `. It is generally recommended running ` + "`" + `terraform plan` + "`" + ` after importing an OBS bucket. Also you can ignore changes as below. ` + "`" + `` + "`" + `` + "`" + ` resource "flexibleengine_obs_bucket" "bucket" { ... lifecycle { ignore_changes = [ acl, force_destroy, ] } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_obs_bucket_object",
			Category:         "Object Storage Service (OSS)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"object",
				"storage",
				"service",
				"oss",
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
			Type:             "flexibleengine_obs_bucket_replication",
			Category:         "Object Storage Service (OSS)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"object",
				"storage",
				"service",
				"oss",
				"obs",
				"bucket",
				"replication",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) Specifies the name of the source bucket. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "destination_bucket",
					Description: `(Required) Specifies the name of the destination bucket. -> The destination bucket cannot be in the region where the source bucket resides.`,
				},
				resource.Attribute{
					Name:        "agency",
					Description: `(Required) Specifies the IAM agency applied to the cross-region replication function. -> The IAM agency is a cloud service agency of OBS. The OBS project must have the`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Optional) A configuration of object cross-region replication management. The object supports the following: + ` + "`" + `enabled` + "`" + ` - (Optional) Specifies cross-region replication rule status. Defaults to ` + "`" + `true` + "`" + `. + ` + "`" + `prefix` + "`" + ` - (Optional) Specifies the object key prefix identifying one or more objects to which the rule applies and duplicated prefixes are not supported. If omitted, all objects in the bucket will be managed by the lifecycle rule. To copy a folder, end the prefix with a slash (/), for example, imgs/. + ` + "`" + `storage_class` + "`" + ` - (Optional) Specifies the storage class for replicated objects. Valid values are "STANDARD", "WARM" (Infrequent Access) and "COLD" (Archive). If omitted, the storage class of object copies is the same as that of objects in the source bucket. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The name of the bucket.`,
				},
				resource.Attribute{
					Name:        "rule/id",
					Description: `The ID of a rule in UUID format. ## Import OBS bucket cross-region replication can be imported using the`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The name of the bucket.`,
				},
				resource.Attribute{
					Name:        "rule/id",
					Description: `The ID of a rule in UUID format. ## Import OBS bucket cross-region replication can be imported using the`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_rds_account",
			Category:         "Relational Database Service (RDS)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"relational",
				"database",
				"service",
				"rds",
				"account",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String, ForceNew) The region in which to create the rds account resource. If omitted, the provider-level region will be used. Changing this creates a new resource.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, String, ForceNew) Specifies the rds instance id. Changing this will create a new resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, String, ForceNew) Specifies the username of the db account. Only lowercase letters, digits, hyphens (-), and userscores (_) are allowed. Changing this will create a new resource. + If the database version is MySQL 5.6, the username consists of 1 to 16 characters. + If the database version is MySQL 5.7 or 8.0, the username consists of 1 to 32 characters.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required, String) Specifies the password of the db account. The parameter must be 8 to 32 characters long and contain only letters(case-sensitive), digits, and special characters(~!@#$%^`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID of account which is formatted ` + "`" + `<instance_id>/<account_name>` + "`" + `. ## Import RDS account can be imported using the ` + "`" + `instance id` + "`" + ` and ` + "`" + `account name` + "`" + `, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_rds_account.user_1 instance_id/account_name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID of account which is formatted ` + "`" + `<instance_id>/<account_name>` + "`" + `. ## Import RDS account can be imported using the ` + "`" + `instance id` + "`" + ` and ` + "`" + `account name` + "`" + `, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_rds_account.user_1 instance_id/account_name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_rds_database",
			Category:         "Relational Database Service (RDS)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"relational",
				"database",
				"service",
				"rds",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String, ForceNew) The region in which to create the RDS database resource. If omitted, the provider-level region will be used. Changing this creates a new resource.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, String, ForceNew) Specifies the RDS instance ID. Changing this will create a new resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, String, ForceNew) Specifies the database name. The database name contains`,
				},
				resource.Attribute{
					Name:        "character_set",
					Description: `(Required, String, ForceNew) Specifies the character set used by the database, For example`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID of database which is formatted ` + "`" + `<instance_id>/<database_name>` + "`" + `. ## Import RDS database can be imported using the ` + "`" + `instance id` + "`" + ` and ` + "`" + `database name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_rds_database.database_1 instance_id/database_name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID of database which is formatted ` + "`" + `<instance_id>/<database_name>` + "`" + `. ## Import RDS database can be imported using the ` + "`" + `instance id` + "`" + ` and ` + "`" + `database name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_rds_database.database_1 instance_id/database_name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_rds_instance_v1",
			Category:         "Deprecated",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "Database-RDS .svg",
			Keywords: []string{
				"deprecated",
				"rds",
				"instance",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the DB instance name. The DB instance name of the same type is unique in the same tenant. The changes of the instance name will be suppressed in HA scenario.`,
				},
				resource.Attribute{
					Name:        "datastore",
					Description: `(Required) Specifies database information. The structure is described below.`,
				},
				resource.Attribute{
					Name:        "flavorref",
					Description: `(Required) Specifies the specification ID (flavors.id in the response message in Obtaining All DB Instance Specifications). If you want to enable ha for the rds instance, a flavor with ha speccode is required.`,
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
			Type:             "flexibleengine_rds_instance_v3",
			Category:         "Relational Database Service (RDS)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "Database-RDS .svg",
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
					Name:        "name",
					Description: `(Required, String) Specifies the DB instance name. The DB instance name of the same type must be unique for the same tenant. The value must be 4 to 64 characters in length and start with a letter. It is case-sensitive and can contain only letters, digits, hyphens (-), and underscores (_).`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `(Required, String) Specifies the specification code. ->`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required, List, ForceNew) Specifies the list of AZ name. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "db",
					Description: `(Required, String, ForceNew) Specifies the database information. Structure is documented below. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, String, ForceNew) Specifies the VPC ID. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required, String, ForceNew) Specifies the network ID of a subnet. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Required, String) Specifies the security group which the RDS DB instance belongs to.`,
				},
				resource.Attribute{
					Name:        "volume",
					Description: `(Required, List) Specifies the volume information. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "fixed_ip",
					Description: `(Optional, String, ForceNew) Specifies an intranet IP address of RDS DB instance. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "backup_strategy",
					Description: `(Optional, List) Specifies the advanced backup policy. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "ha_replication_mode",
					Description: `(Optional, String, ForceNew) Specifies the replication mode for the standby DB instance. Changing this parameter will create a new resource. + For MySQL, the value is`,
				},
				resource.Attribute{
					Name:        "param_group_id",
					Description: `(Optional, String, ForceNew) Specifies the parameter group ID. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "time_zone",
					Description: `(Optional, String, ForceNew) Specifies the UTC time zone. The value ranges from UTC-12:00 to UTC+12:00 at the full hour, and defaults to`,
				},
				resource.Attribute{
					Name:        "ssl_enable",
					Description: `(Optional, Bool) Specifies whether to enable the SSL for`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Map) A mapping of tags to assign to the RDS instance. Each tag is represented by one key-value pair. The ` + "`" + `db` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required, String, ForceNew) Specifies the DB engine. Available value are`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required, String, ForceNew) Specifies the database version. MySQL databases support MySQL 5.6 and 5.7, example values: "5.6", "5.7". PostgreSQL databases support PostgreSQL 9.5, 9.6, 10 and 11, example values: "9.5", "9.6", "10", "11". Microsoft SQL Server databases support 2014 SE and 2014 EE, example values: "2014_SE", "2014_EE". Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required, String, ForceNew) Specifies the database password. The value cannot be empty and should contain 8 to 32 characters, including uppercase and lowercase letters, digits, and the following special characters: ~!@#%^`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional, Int) Specifies the database port. + The MySQL database port ranges from 1024 to 65535 (excluding 12017 and 33071, which are occupied by the RDS system and cannot be used). The default value is 3306. + The PostgreSQL database port ranges from 2100 to 9500. The default value is 5432. + The Microsoft SQL Server database port can be 1433 or ranges from 2100 to 9500, excluding 5355 and 5985. The default value is 1433. The ` + "`" + `volume` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required, Int) Specifies the volume size. Its value range is from 40 GB to 4000 GB. The value must be a multiple of 10 and greater than the original size.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required, String, ForceNew) Specifies the volume type. Its value can be any of the following and is case-sensitive: +`,
				},
				resource.Attribute{
					Name:        "disk_encryption_id",
					Description: `(Optional, String) Specifies the key ID for disk encryption. Changing this parameter will create a new resource. The ` + "`" + `backup_strategy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "keep_days",
					Description: `(Optional, Int) Specifies the retention days for specific backup files. The value range is from 0 to 732. If this parameter is not specified or set to 0, the automated backup policy is disabled. ->`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `(Required, String) Specifies the backup time window. Automated backups will be triggered during the backup time window. It must be a valid value in the`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates the DB instance status.`,
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
					Description: `Indicates the public IP address list.`,
				},
				resource.Attribute{
					Name:        "db",
					Description: `See Argument Reference above. The ` + "`" + `db` + "`" + ` block also contains: + ` + "`" + `user_name` + "`" + ` - Indicates the default user name of database. The ` + "`" + `nodes` + "`" + ` block contains:`,
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
					Description: `Indicates the node status. ## Timeouts This resource provides the following timeouts configuration options:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `Default is 30 minute.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `Default is 30 minute. ## Import RDS instance can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_rds_instance_v3.instance_1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ` But due to some attrubutes missing from the API response, it's required to ignore changes as below. ` + "`" + `` + "`" + `` + "`" + ` resource "flexibleengine_rds_instance_v3" "instance_1" { ... lifecycle { ignore_changes = [ "db", ] } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates the DB instance status.`,
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
					Description: `Indicates the public IP address list.`,
				},
				resource.Attribute{
					Name:        "db",
					Description: `See Argument Reference above. The ` + "`" + `db` + "`" + ` block also contains: + ` + "`" + `user_name` + "`" + ` - Indicates the default user name of database. The ` + "`" + `nodes` + "`" + ` block contains:`,
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
					Description: `Indicates the node status. ## Timeouts This resource provides the following timeouts configuration options:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `Default is 30 minute.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `Default is 30 minute. ## Import RDS instance can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_rds_instance_v3.instance_1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ` But due to some attrubutes missing from the API response, it's required to ignore changes as below. ` + "`" + `` + "`" + `` + "`" + ` resource "flexibleengine_rds_instance_v3" "instance_1" { ... lifecycle { ignore_changes = [ "db", ] } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_rds_parametergroup_v3",
			Category:         "Relational Database Service (RDS)",
			ShortDescription: ``,
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
					Description: `Indicates the parameter description. ## Import Parameter groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_rds_parametergroup_v3.pg_1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `Indicates the parameter description. ## Import Parameter groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_rds_parametergroup_v3.pg_1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_rds_read_replica_v3",
			Category:         "Relational Database Service (RDS)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"relational",
				"database",
				"service",
				"rds",
				"read",
				"replica",
				"v3",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the DB instance name. The DB instance name of the same type must be unique for the same tenant. The value must be 4 to 64 characters in length and start with a letter. It is case-sensitive and can contain only letters, digits, hyphens (-), and underscores (_). Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `(Required) Specifies the specification code.`,
				},
				resource.Attribute{
					Name:        "replica_of_id",
					Description: `(Required) Specifies the DB instance ID, which is used to create a read replica. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "volume",
					Description: `(Required) Specifies the volume information. Structure is documented below. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required) Specifies the AZ name. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the RDS read replica instance. Each tag is represented by one key-value pair.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to create the instance. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Currently, read replicas can be created only in the same region as that of the promary DB instance. Changing this parameter will create a new resource. The ` + "`" + `volume` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Specifies the volume type. Its value can be any of the following and is case-sensitive: - ULTRAHIGH: indicates the SSD type. - ULTRAHIGHPRO: indicates the ultra-high I/O. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "disk_encryption_id",
					Description: `(Optional) Specifies the key ID for disk encryption. Changing this parameter will create a new resource. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Indicates the instance ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates the instance status.`,
				},
				resource.Attribute{
					Name:        "db",
					Description: `Indicates the database information. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `Indicates the private IP address list.`,
				},
				resource.Attribute{
					Name:        "public_ips",
					Description: `Indicates the public IP address list.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `Indicates the security group which the RDS DB instance belongs to.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `Indicates the subnet id.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `Indicates the VPC ID. The ` + "`" + `db` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Indicates the database port information.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Indicates the DB engine. Value: MySQL, PostgreSQL, SQLServer.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `Indicates the default user name of database.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Indicates the database version. ## Import RDS instance can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_rds_read_replica_v3.instance_1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Indicates the instance ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates the instance status.`,
				},
				resource.Attribute{
					Name:        "db",
					Description: `Indicates the database information. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `Indicates the private IP address list.`,
				},
				resource.Attribute{
					Name:        "public_ips",
					Description: `Indicates the public IP address list.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `Indicates the security group which the RDS DB instance belongs to.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `Indicates the subnet id.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `Indicates the VPC ID. The ` + "`" + `db` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Indicates the database port information.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Indicates the DB engine. Value: MySQL, PostgreSQL, SQLServer.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `Indicates the default user name of database.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Indicates the database version. ## Import RDS instance can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_rds_read_replica_v3.instance_1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_rts_software_config_v1",
			Category:         "Resource Template Service (RTS)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "Management&Deployment-RTS.svg",
			Keywords: []string{
				"resource",
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
			Type:             "flexibleengine_rts_stack_v1",
			Category:         "Resource Template Service (RTS)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"resource",
				"template",
				"service",
				"rts",
				"stack",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the stack. The value must meet the regular expression rule (` + "`" + `^[a-zA-Z][a-zA-Z0-9_.-]{0,254}$` + "`" + `). Changing this creates a new stack.`,
				},
				resource.Attribute{
					Name:        "template_body",
					Description: `(Optional; Required if ` + "`" + `template_url` + "`" + ` is empty) Structure containing the template body. The template content must use the yaml syntax.`,
				},
				resource.Attribute{
					Name:        "template_url",
					Description: `(Optional; Required if ` + "`" + `template_body` + "`" + ` is empty) Location of a file containing the template body.`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `(Optional) Tthe environment information about the stack.`,
				},
				resource.Attribute{
					Name:        "files",
					Description: `(Optional) Files used in the environment.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `(Optional) A list of Parameter structures that specify input parameters for the stack.`,
				},
				resource.Attribute{
					Name:        "disable_rollback",
					Description: `(Optional) Set to true to disable rollback of the stack if stack creation failed.`,
				},
				resource.Attribute{
					Name:        "timeout_mins",
					Description: `(Optional) Specifies the timeout duration. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
					Description: `Specifies the stack status. ## Import RTS Stacks can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_rts_stack_v1.mystack rts-stack ` + "`" + `` + "`" + `` + "`" + ` <a id="timeouts"></a> ## Timeouts ` + "`" + `flexibleengine_rts_stack_v1` + "`" + ` provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `30 minutes` + "`" + `) Used for Creating Stacks - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `30 minutes` + "`" + `) Used for Stack modifications - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `30 minutes` + "`" + `) Used for destroying stacks.`,
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
					Description: `Specifies the stack status. ## Import RTS Stacks can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_rts_stack_v1.mystack rts-stack ` + "`" + `` + "`" + `` + "`" + ` <a id="timeouts"></a> ## Timeouts ` + "`" + `flexibleengine_rts_stack_v1` + "`" + ` provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `30 minutes` + "`" + `) Used for Creating Stacks - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `30 minutes` + "`" + `) Used for Stack modifications - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `30 minutes` + "`" + `) Used for destroying stacks.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_s3_bucket",
			Category:         "Object Storage Service (OSS)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "bucket.svg",
			Keywords: []string{
				"object",
				"storage",
				"service",
				"oss",
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
					Name:        "policy",
					Description: `(Optional) A valid [bucket policy](https://docs.aws.amazon.com/AmazonS3/latest/dev/example-bucket-policies.html) JSON document. Note that if the policy document is not specific enough (but still valid), Terraform may view the policy as constantly changing in a ` + "`" + `terraform plan` + "`" + `. In this case, please make sure you use the verbose/specific version of the policy.`,
				},
				resource.Attribute{
					Name:        "force_destroy",
					Description: `(Optional, Default:false ) A boolean that indicates all objects should be deleted from the bucket so that the bucket can be destroyed without error. These objects are`,
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
					Description: `(Optional) A json array containing [routing rules](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules.html) describing redirect behavior and when redirects are applied. The ` + "`" + `cors_rule` + "`" + ` object supports the following:`,
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
					Description: `The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string. This is used to create Route 53 alias records. ## Import S3 bucket can be imported using the ` + "`" + `bucket` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_s3_bucket.bucket bucket-name ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string. This is used to create Route 53 alias records. ## Import S3 bucket can be imported using the ` + "`" + `bucket` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_s3_bucket.bucket bucket-name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_s3_bucket_object",
			Category:         "Object Storage Service (OSS)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "bucket.svg",
			Keywords: []string{
				"object",
				"storage",
				"service",
				"oss",
				"s3",
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
					Description: `(Optional) Specifies server-side encryption of the object in S3. Valid values are "` + "`" + `AES256` + "`" + `" and "` + "`" + `aws:kms` + "`" + `". Either ` + "`" + `source` + "`" + ` or ` + "`" + `content` + "`" + ` must be provided to specify the bucket content. These two arguments are mutually-exclusive. ## Attributes Reference The following attributes are exported`,
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
			Type:             "flexibleengine_s3_bucket_policy",
			Category:         "Object Storage Service (OSS)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "bucket.svg",
			Keywords: []string{
				"object",
				"storage",
				"service",
				"oss",
				"s3",
				"bucket",
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
			Type:             "flexibleengine_sdrs_drill_v1",
			Category:         "Storage Disaster Recovery Service (SDRS)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"storage",
				"disaster",
				"recovery",
				"service",
				"sdrs",
				"drill",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of a DR drill. The name can contain a maximum of 64 bytes. The value can contain only letters (a to z and A to Z), digits (0 to 9), decimal points (.), underscores (_), and hyphens (-).`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required) Specifies the ID of a protection group. Changing this creates a new drill.`,
				},
				resource.Attribute{
					Name:        "drill_vpc_id",
					Description: `(Required) Specifies the ID used for a DR drill. Changing this creates a new drill. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of a DR drill.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "drill_vpc_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of a DR drill. For details, see [DR Drill Status](https://docs.prod-cloud-ocb.orange-business.com/en-us/api/sdrs/en-us_topic_0126152933.html). ## Import DR drill can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_sdrs_drill_v1.drill_1 22fce838-4bfb-4a92-b9aa-fc80a583eb59 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of a DR drill.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "drill_vpc_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of a DR drill. For details, see [DR Drill Status](https://docs.prod-cloud-ocb.orange-business.com/en-us/api/sdrs/en-us_topic_0126152933.html). ## Import DR drill can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_sdrs_drill_v1.drill_1 22fce838-4bfb-4a92-b9aa-fc80a583eb59 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_sdrs_protectedinstance_v1",
			Category:         "Storage Disaster Recovery Service (SDRS)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"storage",
				"disaster",
				"recovery",
				"service",
				"sdrs",
				"protectedinstance",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of a protected instance.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of a protected instance. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required) Specifies the ID of the protection group where a protected instance is added. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "server_id",
					Description: `(Required) Specifies the ID of the source server. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Optional) Specifies the ID of a storage pool. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "primary_subnet_id",
					Description: `(Optional) Specifies the subnet ID of the primary NIC on the target server. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "primary_ip_address",
					Description: `(Optional) Specifies the IP address of the primary NIC on the target server. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "delete_target_server",
					Description: `(Optional) Specifies whether to delete the target server. The default value is false.. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "delete_target_eip",
					Description: `(Optional) Specifies whether to delete the EIP of the target server. The default value is false. Changing this creates a new instance. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the protected instance.`,
				},
				resource.Attribute{
					Name:        "target_server",
					Description: `ID of the target server.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the protected instance.`,
				},
				resource.Attribute{
					Name:        "target_server",
					Description: `ID of the target server.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_sdrs_protectiongroup_v1",
			Category:         "Storage Disaster Recovery Service (SDRS)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"storage",
				"disaster",
				"recovery",
				"service",
				"sdrs",
				"protectiongroup",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of a protection group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of a protection group. Changing this creates a new group.`,
				},
				resource.Attribute{
					Name:        "source_availability_zone",
					Description: `(Required) Specifies the source AZ of a protection group. Changing this creates a new group.`,
				},
				resource.Attribute{
					Name:        "target_availability_zone",
					Description: `(Required) Specifies the target AZ of a protection group. Changing this creates a new group.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Required) Specifies the ID of an active-active domain. Changing this creates a new group.`,
				},
				resource.Attribute{
					Name:        "source_vpc_id",
					Description: `(Required) Specifies the ID of the source VPC. Changing this creates a new group.`,
				},
				resource.Attribute{
					Name:        "dr_type",
					Description: `(Optional) Specifies the deployment model. The default value is migration indicating migration within a VPC. Changing this creates a new group.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional) Enable protection or not. It can only be set to true when there's replication pairs within the protection group. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the protection group. ## Import Protection groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_sdrs_protectiongroup_v1.group_1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the protection group. ## Import Protection groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_sdrs_protectiongroup_v1.group_1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_sdrs_replication_attach_v1",
			Category:         "Storage Disaster Recovery Service (SDRS)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"storage",
				"disaster",
				"recovery",
				"service",
				"sdrs",
				"replication",
				"attach",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) Specifies the ID of a protected instance. Changing this creates a new replication attach.`,
				},
				resource.Attribute{
					Name:        "replication_id",
					Description: `(Required) Specifies the ID of a replication pair. Changing this creates a new replication attach.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `(Required) Specifies the device name, eg. /dev/vdb. Changing this creates a new replication attach. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in format of <instance_id>:<replication_id>.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the SDRS replication attch resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in format of <instance_id>:<replication_id>.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the SDRS replication attch resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_sdrs_replication_pair_v1",
			Category:         "Storage Disaster Recovery Service (SDRS)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"storage",
				"disaster",
				"recovery",
				"service",
				"sdrs",
				"replication",
				"pair",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of a replication pair. The name can contain a maximum of 64 bytes. The value can contain only letters (a to z and A to Z), digits (0 to 9), decimal points (.), underscores (_), and hyphens (-).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of a replication pair. Changing this creates a new pair.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required) Specifies the ID of a protection group. Changing this creates a new pair.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `(Required) Specifies the ID of a source disk. Changing this creates a new pair.`,
				},
				resource.Attribute{
					Name:        "delete_target_volume",
					Description: `(Optional) Specifies whether to delete the target disk. The default value is ` + "`" + `false` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the replication pair.`,
				},
				resource.Attribute{
					Name:        "fault_level",
					Description: `Specifies the fault level of a replication pair.`,
				},
				resource.Attribute{
					Name:        "replication_model",
					Description: `Specifies the replication mode of a replication pair. The default value is ` + "`" + `hypermetro` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Specifies the status of a replication pair.`,
				},
				resource.Attribute{
					Name:        "target_volume_id",
					Description: `Specifies the ID of the disk in the protection availability zone. ## Import Replication pairs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_sdrs_replication_pair_v1.replication_1 43b28b66-770b-4e9e-b5c6-cfc43f0593d9 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the replication pair.`,
				},
				resource.Attribute{
					Name:        "fault_level",
					Description: `Specifies the fault level of a replication pair.`,
				},
				resource.Attribute{
					Name:        "replication_model",
					Description: `Specifies the replication mode of a replication pair. The default value is ` + "`" + `hypermetro` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Specifies the status of a replication pair.`,
				},
				resource.Attribute{
					Name:        "target_volume_id",
					Description: `Specifies the ID of the disk in the protection availability zone. ## Import Replication pairs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_sdrs_replication_pair_v1.replication_1 43b28b66-770b-4e9e-b5c6-cfc43f0593d9 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_sfs_access_rule_v2",
			Category:         "Scalable File Service (SFS)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"scalable",
				"file",
				"service",
				"sfs",
				"access",
				"rule",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "sfs_id",
					Description: `(Required) Specifies the UUID of the shared file system. Changing this will create a new access rule.`,
				},
				resource.Attribute{
					Name:        "access_level",
					Description: `(Optional) Specifies the access level of the shared file system. Possible values are`,
				},
				resource.Attribute{
					Name:        "access_type",
					Description: `(Optional) Specifies the type of the share access rule. The default value is`,
				},
				resource.Attribute{
					Name:        "access_to",
					Description: `(Required) Specifies the value that defines the access rule. The value contains 1 to 255 characters. Changing this will create a new access rule. The value varies according to the scenario: - Set the VPC ID in VPC authorization scenarios. - Set this parameter in IP address authorization scenario. For an NFS shared file system, the value in the format of`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the share access rule.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the share access rule. ## Import SFS access rule can be imported by specifying the SFS ID and access rule ID separated by a slash, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_sfs_access_rule_v2 <sfs_id>/<rule_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the share access rule.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the share access rule. ## Import SFS access rule can be imported by specifying the SFS ID and access rule ID separated by a slash, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_sfs_access_rule_v2 <sfs_id>/<rule_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_sfs_file_system_v2",
			Category:         "Scalable File Service (SFS)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "Storage-SFS.svg",
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
					Description: `(Optional) Metadata key and value pairs as a dictionary of strings. The supported metadata keys are "#sfs_crypt_key_id", "#sfs_crypt_domain_id" and "#sfs_crypt_alias", and the keys should be exist at the same time to enable the data encryption function. Changing this will create a new resource.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) The availability zone name. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "access_level",
					Description: `(Optional) Specifies the access level of the shared file system. Possible values are`,
				},
				resource.Attribute{
					Name:        "access_type",
					Description: `(Optional) Specifies the type of the share access rule. The default value is`,
				},
				resource.Attribute{
					Name:        "access_to",
					Description: `(Optional) Specifies the value that defines the access rule. The value contains 1 to 255 characters. Changing this will create a new access rule. The value varies according to the scenario: - Set the VPC ID in VPC authorization scenarios. - Set this parameter in IP address authorization scenario. - For an NFS shared file system, the value in the format of`,
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
					Name:        "volume_type",
					Description: `The volume type.`,
				},
				resource.Attribute{
					Name:        "export_location",
					Description: `The address for accessing the shared file system.`,
				},
				resource.Attribute{
					Name:        "share_access_id",
					Description: `The UUID of the share access rule.`,
				},
				resource.Attribute{
					Name:        "access_rules_status",
					Description: `The status of the share access rule.`,
				},
				resource.Attribute{
					Name:        "access_rules",
					Description: `All access rules of the shared file system. The object includes the following: - ` + "`" + `access_rule_id` + "`" + ` - The UUID of the share access rule. - ` + "`" + `access_level` + "`" + ` - The access level of the shared file system - ` + "`" + `access_type` + "`" + ` - The type of the share access rule. - ` + "`" + `access_to` + "`" + ` - The value that defines the access rule. - ` + "`" + `status` + "`" + ` - The status of the share access rule. ## Import SFS can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_sfs_file_system_v2 4779ab1c-7c1a-44b1-a02e-93dfc361b32d ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "volume_type",
					Description: `The volume type.`,
				},
				resource.Attribute{
					Name:        "export_location",
					Description: `The address for accessing the shared file system.`,
				},
				resource.Attribute{
					Name:        "share_access_id",
					Description: `The UUID of the share access rule.`,
				},
				resource.Attribute{
					Name:        "access_rules_status",
					Description: `The status of the share access rule.`,
				},
				resource.Attribute{
					Name:        "access_rules",
					Description: `All access rules of the shared file system. The object includes the following: - ` + "`" + `access_rule_id` + "`" + ` - The UUID of the share access rule. - ` + "`" + `access_level` + "`" + ` - The access level of the shared file system - ` + "`" + `access_type` + "`" + ` - The type of the share access rule. - ` + "`" + `access_to` + "`" + ` - The value that defines the access rule. - ` + "`" + `status` + "`" + ` - The status of the share access rule. ## Import SFS can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_sfs_file_system_v2 4779ab1c-7c1a-44b1-a02e-93dfc361b32d ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_sfs_turbo",
			Category:         "Scalable File Service (SFS)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"scalable",
				"file",
				"service",
				"sfs",
				"turbo",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of an SFS Turbo file system. The value contains 4 to 64 characters and must start with a letter. Changing this will create a new resource.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) Specifies the capacity of a common file system, in GB. The value ranges from 500 to 32768.`,
				},
				resource.Attribute{
					Name:        "share_proto",
					Description: `(Optional) Specifies the protocol for sharing file systems. The valid value is NFS. Changing this will create a new resource.`,
				},
				resource.Attribute{
					Name:        "share_type",
					Description: `(Optional) Specifies the file system type. The valid values are STANDARD and PERFORMANCE. Changing this will create a new resource.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required) Specifies the availability zone where the file system is located. Changing this will create a new resource.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) Specifies the VPC ID. Changing this will create a new resource.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) Specifies the network ID of the subnet. Changing this will create a new resource.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Required) Specifies the security group ID. Changing this will create a new resource.`,
				},
				resource.Attribute{
					Name:        "crypt_key_id",
					Description: `(Optional) Specifies the ID of a KMS key to encrypt the file system. Changing this will create a new resource. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the SFS Turbo file system.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region of the SFS Turbo file system.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the SFS Turbo file system.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version ID of the SFS Turbo file system.`,
				},
				resource.Attribute{
					Name:        "export_location",
					Description: `Tthe mount point of the SFS Turbo file system.`,
				},
				resource.Attribute{
					Name:        "available_capacity",
					Description: `The available capacity of the SFS Turbo file system in the unit of GB. ## Import SFS Turbo can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_sfs_turbo 1e3d5306-24c9-4316-9185-70e9787d71ab ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the SFS Turbo file system.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region of the SFS Turbo file system.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the SFS Turbo file system.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version ID of the SFS Turbo file system.`,
				},
				resource.Attribute{
					Name:        "export_location",
					Description: `Tthe mount point of the SFS Turbo file system.`,
				},
				resource.Attribute{
					Name:        "available_capacity",
					Description: `The available capacity of the SFS Turbo file system in the unit of GB. ## Import SFS Turbo can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_sfs_turbo 1e3d5306-24c9-4316-9185-70e9787d71ab ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_smn_subscription_v2",
			Category:         "Simple Message Notification (SMN)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "Application-SMN.svg",
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
			Type:             "flexibleengine_smn_topic_v2",
			Category:         "Simple Message Notification (SMN)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "Application-SMN.svg",
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
					Description: `(Optional) Topic display name, which is presented as the name of the email sender in an email message. ## Attributes Reference The following attributes are exported:`,
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
					Description: `Resource identifier of a topic, which is unique.`,
				},
				resource.Attribute{
					Name:        "push_policy",
					Description: `Message pushing policy. 0 indicates that the message sending fails and the message is cached in the queue. 1 indicates that the failed message is discarded.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Time when the topic was created.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Time when the topic was updated.`,
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
					Description: `Resource identifier of a topic, which is unique.`,
				},
				resource.Attribute{
					Name:        "push_policy",
					Description: `Message pushing policy. 0 indicates that the message sending fails and the message is cached in the queue. 1 indicates that the failed message is discarded.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Time when the topic was created.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Time when the topic was updated.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_swr_organization",
			Category:         "Software Repository for Container (SWR)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"software",
				"repository",
				"for",
				"container",
				"swr",
				"organization",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String, ForceNew) Specifies the region in which to create the resource. If omitted, the provider-level region will be used. Changing this creates a new resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, String, ForceNew) Specifies the name of the organization. The organization name must be globally unique. Changing this creates a new resource. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the organization.`,
				},
				resource.Attribute{
					Name:        "creator",
					Description: `The creator user name of the organization.`,
				},
				resource.Attribute{
					Name:        "permission",
					Description: `The permission of the organization, the value can be Manage, Write, and Read.`,
				},
				resource.Attribute{
					Name:        "login_server",
					Description: `The URL that can be used to log into the container registry. ## Import Organizations can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_swr_organization.test org-name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the organization.`,
				},
				resource.Attribute{
					Name:        "creator",
					Description: `The creator user name of the organization.`,
				},
				resource.Attribute{
					Name:        "permission",
					Description: `The permission of the organization, the value can be Manage, Write, and Read.`,
				},
				resource.Attribute{
					Name:        "login_server",
					Description: `The URL that can be used to log into the container registry. ## Import Organizations can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_swr_organization.test org-name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_swr_organization_users",
			Category:         "Software Repository for Container (SWR)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"software",
				"repository",
				"for",
				"container",
				"swr",
				"organization",
				"users",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String, ForceNew) Specifies the region in which to create the resource. If omitted, the provider-level region will be used. Changing this creates a new resource.`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `(Required, String, ForceNew) Specifies the name of the organization (namespace) to be accessed. Changing this creates a new resource.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `(Required, List) Specifies the users to access to the organization (namespace). Structure is documented below. The ` + "`" + `users` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "permission",
					Description: `(Required, String) Specifies the permission of the existing IAM user. The values can be`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `(Required, String) Specifies the ID of the existing IAM user.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `(Optional, String) Specifies the name of the existing IAM user. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. The value is the name of the organization.`,
				},
				resource.Attribute{
					Name:        "creator",
					Description: `The creator user name of the organization.`,
				},
				resource.Attribute{
					Name:        "self_permission",
					Description: `The permission informations of current user. The ` + "`" + `self_permission` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `The name of current user.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `The ID of current user.`,
				},
				resource.Attribute{
					Name:        "permission",
					Description: `The permission of current user. ## Import Organization Permissions can be imported using the ` + "`" + `id` + "`" + ` (organization name), e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_swr_organization_users.test org-test ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. The value is the name of the organization.`,
				},
				resource.Attribute{
					Name:        "creator",
					Description: `The creator user name of the organization.`,
				},
				resource.Attribute{
					Name:        "self_permission",
					Description: `The permission informations of current user. The ` + "`" + `self_permission` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `The name of current user.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `The ID of current user.`,
				},
				resource.Attribute{
					Name:        "permission",
					Description: `The permission of current user. ## Import Organization Permissions can be imported using the ` + "`" + `id` + "`" + ` (organization name), e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_swr_organization_users.test org-test ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_swr_repository",
			Category:         "Software Repository for Container (SWR)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"software",
				"repository",
				"for",
				"container",
				"swr",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String, ForceNew) Specifies the region in which to create the resource. If omitted, the provider-level region will be used. Changing this creates a new resource.`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `(Required, String, ForceNew) Specifies the name of the organization (namespace) the repository belongs. Changing this creates a new resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, String, ForceNew) Specifies the name of the repository. Changing this creates a new resource.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `(Optional, Bool) Specifies whether the repository is public. Default is ` + "`" + `false` + "`" + `. + ` + "`" + `true` + "`" + ` - Indicates the repository is`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, String) Specifies the description of the repository.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `(Optional, String) Specifies the category of the repository. The value can be ` + "`" + `app_server` + "`" + `, ` + "`" + `linux` + "`" + `, ` + "`" + `framework_app` + "`" + `, ` + "`" + `database` + "`" + `, ` + "`" + `lang` + "`" + `, ` + "`" + `other` + "`" + `, ` + "`" + `windows` + "`" + `, ` + "`" + `arm` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the repository. The value is the name of the repository.`,
				},
				resource.Attribute{
					Name:        "repository_id",
					Description: `Numeric ID of the repository`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `Image address for docker pull.`,
				},
				resource.Attribute{
					Name:        "internal_path",
					Description: `Intra-cluster image address for docker pull.`,
				},
				resource.Attribute{
					Name:        "num_images",
					Description: `Number of image tags in a repository.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Repository size. ## Import Repository can be imported using the organization name and repository name separated by a slash, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_swr_repository.test org-name/repo-name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the repository. The value is the name of the repository.`,
				},
				resource.Attribute{
					Name:        "repository_id",
					Description: `Numeric ID of the repository`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `Image address for docker pull.`,
				},
				resource.Attribute{
					Name:        "internal_path",
					Description: `Intra-cluster image address for docker pull.`,
				},
				resource.Attribute{
					Name:        "num_images",
					Description: `Number of image tags in a repository.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Repository size. ## Import Repository can be imported using the organization name and repository name separated by a slash, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_swr_repository.test org-name/repo-name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_swr_repository_sharing",
			Category:         "Software Repository for Container (SWR)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"software",
				"repository",
				"for",
				"container",
				"swr",
				"sharing",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String, ForceNew) Specifies the region in which to create the resource. If omitted, the provider-level region will be used. Changing this creates a new resource.`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `(Required, String, ForceNew) Specifies the name of the organization (namespace) the repository belongs. Changing this creates a new resource.`,
				},
				resource.Attribute{
					Name:        "repository",
					Description: `(Required, String, ForceNew) Specifies the name of the repository to be shared. Changing this creates a new resource.`,
				},
				resource.Attribute{
					Name:        "sharing_account",
					Description: `(Required, String, ForceNew) Specifies the name of the account for repository sharing. Changing this creates a new resource. ->`,
				},
				resource.Attribute{
					Name:        "deadline",
					Description: `(Required, String) Specifies the end date of image sharing (UTC time in YYYY-MM-DD format, for example ` + "`" + `2021-10-01` + "`" + `). When the value is set to`,
				},
				resource.Attribute{
					Name:        "permission",
					Description: `(Optional, String) Specifies the permission to be granted. Currently, only the`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, String) Specifies the description of the repository sharing. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the repository sharing. The value is the value of ` + "`" + `sharing_account` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates the repository sharing is valid (true) or expired (false). ## Import Repository sharing can be imported using the organization name, repository name and sharing account separated by a slash, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_swr_repository_sharing.test org-name/repo-name/sharing-account ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the repository sharing. The value is the value of ` + "`" + `sharing_account` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates the repository sharing is valid (true) or expired (false). ## Import Repository sharing can be imported using the organization name, repository name and sharing account separated by a slash, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_swr_repository_sharing.test org-name/repo-name/sharing-account ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_vbs_backup_policy_v2",
			Category:         "Volume Backup Service (VBS)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"volume",
				"backup",
				"service",
				"vbs",
				"policy",
				"v2",
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
					Description: `(Optional) Specifies one or more volumes associated with the backup policy. Any previously associated backup policy will no longer apply. ## Attributes Reference All of the argument attributes are also exported as result attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Specifies a backup policy ID.`,
				},
				resource.Attribute{
					Name:        "policy_resource_count",
					Description: `Specifies the number of volumes associated with the backup policy. ## Import Backup Policy can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_vbs_backup_policy_v2.vbs 4779ab1c-7c1a-44b1-a02e-93dfc361b32d ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Specifies a backup policy ID.`,
				},
				resource.Attribute{
					Name:        "policy_resource_count",
					Description: `Specifies the number of volumes associated with the backup policy. ## Import Backup Policy can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_vbs_backup_policy_v2.vbs 4779ab1c-7c1a-44b1-a02e-93dfc361b32d ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_vbs_backup_v2",
			Category:         "Volume Backup Service (VBS)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"volume",
				"backup",
				"service",
				"vbs",
				"v2",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_vpc_eip_v1",
			Category:         "Elastic IP (EIP)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "Network-EIP.svg",
			Keywords: []string{
				"elastic",
				"ip",
				"eip",
				"vpc",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to create the EIP. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new EIP.`,
				},
				resource.Attribute{
					Name:        "publicip",
					Description: `(Required) The elastic IP address object.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Required) The bandwidth object.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value pairs to associate with the EIP. The ` + "`" + `publicip` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The value must be a type supported by the system. Only ` + "`" + `5_bgp` + "`" + ` supported now. Changing this creates a new EIP.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) The value must be a valid IP address in the available IP address segment. Changing this creates a new EIP.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `(Optional) The port id which this EIP will associate with. If the value is not specified, the EIP will be in unbind state. The ` + "`" + `bandwidth` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The bandwidth name, which is a string of 1 to 64 characters that contain letters, digits, underscores (_), and hyphens (-).`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) The bandwidth size. The value ranges from 1 to 1000 Mbit/s.`,
				},
				resource.Attribute{
					Name:        "share_type",
					Description: `(Required) Specifies the bandwidth type. The value is`,
				},
				resource.Attribute{
					Name:        "charge_mode",
					Description: `(Optional) Specifies whether the bandwidth is billed by traffic or by bandwidth size. Only`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The IP address of the EIP.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of EIP. ## Import EIPs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_vpc_eip_v1.eip_1 2c7f39f3-702b-48d1-940c-b50384177ee1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The IP address of the EIP.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of EIP. ## Import EIPs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_vpc_eip_v1.eip_1 2c7f39f3-702b-48d1-940c-b50384177ee1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_vpc_flow_log_v1",
			Category:         "Virtual Private Cloud (VPC)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"flow",
				"log",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String, ForceNew) The region in which to create the VPC flow log resource. If omitted, the provider-level region will be used. Changing this creates a new VPC flow log.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, String) Specifies the VPC flow log name. The value is a string of 1 to 64 characters that can contain letters, digits, underscores (_), hyphens (-) and periods (.).`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required, String, ForceNew) Specifies the network port ID. Changing this creates a new VPC flow log.`,
				},
				resource.Attribute{
					Name:        "log_group_id",
					Description: `(Required, String, ForceNew) Specifies the LTS log group ID. Changing this creates a new VPC flow log.`,
				},
				resource.Attribute{
					Name:        "log_topic_id",
					Description: `(Required, String, ForceNew) Specifies the LTS log topic ID. Changing this creates a new VPC flow log.`,
				},
				resource.Attribute{
					Name:        "traffic_type",
					Description: `(Optinal, String, ForceNew) Specifies the type of traffic to log. The value can be: -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optinal, String) Specifies supplementary information about the VPC flow log. The value is a string of no more than 255 characters and cannot contain angle brackets (< or >). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The VPC flow log ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `The type of resource on which to create the VPC flow log. The value is fixed to`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the flow log. The value can be ` + "`" + `ACTIVE` + "`" + `, ` + "`" + `DOWN` + "`" + ` or ` + "`" + `ERROR` + "`" + `. ## Import VPC flow logs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `sh terraform import flexibleengine_vpc_flow_log_v1.flowlog1 41b9d73f-eb1c-4795-a100-59a99b062513 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The VPC flow log ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `The type of resource on which to create the VPC flow log. The value is fixed to`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the flow log. The value can be ` + "`" + `ACTIVE` + "`" + `, ` + "`" + `DOWN` + "`" + ` or ` + "`" + `ERROR` + "`" + `. ## Import VPC flow logs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `sh terraform import flexibleengine_vpc_flow_log_v1.flowlog1 41b9d73f-eb1c-4795-a100-59a99b062513 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_vpc_peering_accepter_v2",
			Category:         "Virtual Private Cloud (VPC)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"peering",
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
			Type:             "flexibleengine_vpc_peering_v2",
			Category:         "Virtual Private Cloud (VPC)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"peering",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The VPC peering connection ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The VPC peering connection status. The value can be PENDING_ACCEPTANCE, REJECTED, EXPIRED, DELETED, or ACTIVE. ## Notes If you create a VPC peering connection with another VPC of your own, the connection is created without the need for you to accept the connection. ## Import VPC Peering resources can be imported using the ` + "`" + `vpc peering id` + "`" + `, e.g. > $ terraform import flexibleengine_vpc_peering_connection_v2.test_connection 22b76469-08e3-4937-8c1d-7aad34892be1`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The VPC peering connection ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The VPC peering connection status. The value can be PENDING_ACCEPTANCE, REJECTED, EXPIRED, DELETED, or ACTIVE. ## Notes If you create a VPC peering connection with another VPC of your own, the connection is created without the need for you to accept the connection. ## Import VPC Peering resources can be imported using the ` + "`" + `vpc peering id` + "`" + `, e.g. > $ terraform import flexibleengine_vpc_peering_connection_v2.test_connection 22b76469-08e3-4937-8c1d-7aad34892be1`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_vpc_route_v2",
			Category:         "Virtual Private Cloud (VPC)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "Network-VPC.svg",
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
			Type:             "flexibleengine_vpc_subnet_v1",
			Category:         "Virtual Private Cloud (VPC)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "subnet.svg",
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"subnet",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String, ForceNew) Specifies the region in which to create the vpc subnet. If omitted, the provider-level region will be used. Changing this creates a new subnet.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Map) The key/value pairs to associate with the subnet. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the subnet. The value can be ACTIVE, DOWN, UNKNOWN, or ERROR.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The subnet (Native OpenStack API) ID.`,
				},
				resource.Attribute{
					Name:        "ipv6_subnet_id",
					Description: `The ID of the IPv6 subnet (Native OpenStack API).`,
				},
				resource.Attribute{
					Name:        "ipv6_cidr",
					Description: `The IPv6 subnet CIDR block.`,
				},
				resource.Attribute{
					Name:        "ipv6_gateway",
					Description: `The IPv6 subnet gateway. ## Import Subnets can be imported using the ` + "`" + `subnet id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_vpc_subnet_v1 4779ab1c-7c1a-44b1-a02e-93dfc361b32d ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the subnet. The value can be ACTIVE, DOWN, UNKNOWN, or ERROR.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The subnet (Native OpenStack API) ID.`,
				},
				resource.Attribute{
					Name:        "ipv6_subnet_id",
					Description: `The ID of the IPv6 subnet (Native OpenStack API).`,
				},
				resource.Attribute{
					Name:        "ipv6_cidr",
					Description: `The IPv6 subnet CIDR block.`,
				},
				resource.Attribute{
					Name:        "ipv6_gateway",
					Description: `The IPv6 subnet gateway. ## Import Subnets can be imported using the ` + "`" + `subnet id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_vpc_subnet_v1 4779ab1c-7c1a-44b1-a02e-93dfc361b32d ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_vpc_v1",
			Category:         "Virtual Private Cloud (VPC)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "Network-VPC.svg",
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String, ForceNew) Specifies the region in which to create the VPC. If omitted, the provider-level region will be used. Changing this creates a new VPC resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, String) Specifies the name of the VPC. The name must be unique for a tenant. The value is a string of no more than 64 characters and can contain digits, letters, underscores (_), and hyphens (-).`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Required, String) Specifies the range of available subnets in the VPC. The value ranges from 10.0.0.0/8 to 10.255.255.0/24, 172.16.0.0/12 to 172.31.255.0/24, or 192.168.0.0/16 to 192.168.255.0/24.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, String) Specifies supplementary information about the VPC. The value is a string of no more than 255 characters and cannot contain angle brackets (< or >).`,
				},
				resource.Attribute{
					Name:        "secondary_cidr",
					Description: `(Optional, String) Specifies the secondary CIDR block of the VPC. -> The argument cannot be imported into your Terraform state. And the following secondary CIDR blocks cannot be added to a VPC: 10.0.0.0/8, 172.16.0.0/12, and 192.168.0.0/16. [View the complete list of unsupported CIDR blocks](https://docs.prod-cloud-ocb.orange-business.com/usermanual/vpc/vpc_vpc_0007.html).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Map) Specifies the key/value pairs to associate with the VPC. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The VPC ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current status of the VPC. Possible values are as follows: CREATING, OK or ERROR. ## Import VPCs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_vpc_v1.vpc_v1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The VPC ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current status of the VPC. Possible values are as follows: CREATING, OK or ERROR. ## Import VPCs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_vpc_v1.vpc_v1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_vpcep_approval",
			Category:         "VPC Endpoint (VPCEP)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"endpoint",
				"vpcep",
				"approval",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID in UUID format which equals to the ID of the VPC endpoint service.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region in which to obtain the VPC endpoint service.`,
				},
				resource.Attribute{
					Name:        "connections",
					Description: `An array of VPC endpoints connect to the VPC endpoint service. Structure is documented below. - ` + "`" + `endpoint_id` + "`" + ` - The unique ID of the VPC endpoint. - ` + "`" + `packet_id` + "`" + ` - The packet ID of the VPC endpoint. - ` + "`" + `domain_id` + "`" + ` - The user's domain ID. - ` + "`" + `status` + "`" + ` - The connection status of the VPC endpoint. ## Timeouts This resource provides the following timeouts configuration options: - ` + "`" + `create` + "`" + ` - Default is 10 minute. - ` + "`" + `delete` + "`" + ` - Default is 3 minute.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID in UUID format which equals to the ID of the VPC endpoint service.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region in which to obtain the VPC endpoint service.`,
				},
				resource.Attribute{
					Name:        "connections",
					Description: `An array of VPC endpoints connect to the VPC endpoint service. Structure is documented below. - ` + "`" + `endpoint_id` + "`" + ` - The unique ID of the VPC endpoint. - ` + "`" + `packet_id` + "`" + ` - The packet ID of the VPC endpoint. - ` + "`" + `domain_id` + "`" + ` - The user's domain ID. - ` + "`" + `status` + "`" + ` - The connection status of the VPC endpoint. ## Timeouts This resource provides the following timeouts configuration options: - ` + "`" + `create` + "`" + ` - Default is 10 minute. - ` + "`" + `delete` + "`" + ` - Default is 3 minute.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_vpcep_endpoint",
			Category:         "VPC Endpoint (VPCEP)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"endpoint",
				"vpcep",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value pairs to associate with the VPC endpoint. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the VPC endpoint.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region in which to create the VPC endpoint.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the VPC endpoint. The value can be`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `The name of the VPC endpoint service.`,
				},
				resource.Attribute{
					Name:        "service_type",
					Description: `The type of the VPC endpoint service.`,
				},
				resource.Attribute{
					Name:        "packet_id",
					Description: `The packet ID of the VPC endpoint.`,
				},
				resource.Attribute{
					Name:        "private_domain_name",
					Description: `The domain name for accessing the associated VPC endpoint service. This parameter is only available when enable_dns is set to true. ## Import VPC endpoint can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_vpcep_endpoint.test 828907cc-40c9-42fe-8206-ecc1bdd30060 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the VPC endpoint.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region in which to create the VPC endpoint.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the VPC endpoint. The value can be`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `The name of the VPC endpoint service.`,
				},
				resource.Attribute{
					Name:        "service_type",
					Description: `The type of the VPC endpoint service.`,
				},
				resource.Attribute{
					Name:        "packet_id",
					Description: `The packet ID of the VPC endpoint.`,
				},
				resource.Attribute{
					Name:        "private_domain_name",
					Description: `The domain name for accessing the associated VPC endpoint service. This parameter is only available when enable_dns is set to true. ## Import VPC endpoint can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_vpcep_endpoint.test 828907cc-40c9-42fe-8206-ecc1bdd30060 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_vpcep_service",
			Category:         "VPC Endpoint (VPCEP)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"endpoint",
				"vpcep",
				"service",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value pairs to associate with the VPC endpoint service. The ` + "`" + `port_mapping` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) Specifies the protocol used in port mappings. The value can be _TCP_ or _UDP_. The default value is _TCP_.`,
				},
				resource.Attribute{
					Name:        "service_port",
					Description: `(Optional) Specifies the port for accessing the VPC endpoint service. This port is provided by the backend service to provide services. The value ranges from 1 to 65535.`,
				},
				resource.Attribute{
					Name:        "terminal_port",
					Description: `(Optional) Specifies the port for accessing the VPC endpoint. This port is provided by the VPC endpoint, allowing you to access the VPC endpoint service. The value ranges from 1 to 65535. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the VPC endpoint service.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region in which to create the VPC endpoint service.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the VPC endpoint service. The value can be`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `The full name of the VPC endpoint service in the format:`,
				},
				resource.Attribute{
					Name:        "service_type",
					Description: `The type of the VPC endpoint service. Only`,
				},
				resource.Attribute{
					Name:        "connections",
					Description: `An array of VPC endpoints connect to the VPC endpoint service. Structure is documented below. - ` + "`" + `endpoint_id` + "`" + ` - The unique ID of the VPC endpoint. - ` + "`" + `packet_id` + "`" + ` - The packet ID of the VPC endpoint. - ` + "`" + `domain_id` + "`" + ` - The user's domain ID. - ` + "`" + `status` + "`" + ` - The connection status of the VPC endpoint. ## Import VPC endpoint services can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_vpcep_service.test_service 950cd3ba-9d0e-4451-97c1-3e97dd515d46 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the VPC endpoint service.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region in which to create the VPC endpoint service.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the VPC endpoint service. The value can be`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `The full name of the VPC endpoint service in the format:`,
				},
				resource.Attribute{
					Name:        "service_type",
					Description: `The type of the VPC endpoint service. Only`,
				},
				resource.Attribute{
					Name:        "connections",
					Description: `An array of VPC endpoints connect to the VPC endpoint service. Structure is documented below. - ` + "`" + `endpoint_id` + "`" + ` - The unique ID of the VPC endpoint. - ` + "`" + `packet_id` + "`" + ` - The packet ID of the VPC endpoint. - ` + "`" + `domain_id` + "`" + ` - The user's domain ID. - ` + "`" + `status` + "`" + ` - The connection status of the VPC endpoint. ## Import VPC endpoint services can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_vpcep_service.test_service 950cd3ba-9d0e-4451-97c1-3e97dd515d46 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_waf_certificate",
			Category:         "Web Application Firewall (WAF)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"web",
				"application",
				"firewall",
				"waf",
				"certificate",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required, String) Specifies the certificate name. The maximum length is 256 characters. Only digits, letters, underscores(` + "`" + `_` + "`" + `), and hyphens(` + "`" + `-` + "`" + `) are allowed.`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(Required, String, ForceNew) Specifies the certificate content. Changing this creates a new certificate.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Required, String, ForceNew) Specifies the private key. Changing this creates a new certificate. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The certificate ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `Indicates the time when the certificate expires. ## Import Certificates can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `sh terraform import flexibleengine_waf_certificate.cert_1 9251a0ed5aa640b68a35cf2eb6a3b733 ` + "`" + `` + "`" + `` + "`" + ` Note that the imported state is not identical to your resource definition, due to security reason. The missing attributes include ` + "`" + `certificate` + "`" + `, and ` + "`" + `private_key` + "`" + `. You can ignore changes as below. ` + "`" + `` + "`" + `` + "`" + ` resource "flexibleengine_waf_certificate" "cert_1" { ... lifecycle { ignore_changes = [ certificate, private_key, ] } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The certificate ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `Indicates the time when the certificate expires. ## Import Certificates can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `sh terraform import flexibleengine_waf_certificate.cert_1 9251a0ed5aa640b68a35cf2eb6a3b733 ` + "`" + `` + "`" + `` + "`" + ` Note that the imported state is not identical to your resource definition, due to security reason. The missing attributes include ` + "`" + `certificate` + "`" + `, and ` + "`" + `private_key` + "`" + `. You can ignore changes as below. ` + "`" + `` + "`" + `` + "`" + ` resource "flexibleengine_waf_certificate" "cert_1" { ... lifecycle { ignore_changes = [ certificate, private_key, ] } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_waf_domain",
			Category:         "Web Application Firewall (WAF)",
			ShortDescription: ``,
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
					Name:        "domain",
					Description: `(Required, String, ForceNew) Specifies the domain name to be protected. For example, www.example.com or`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `(Required, List) Specifies an array of origin web servers. The object structure is documented below.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `(Optional, String) Specifies the certificate ID. This parameter is mandatory when ` + "`" + `client_protocol` + "`" + ` is set to HTTPS.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Optional, String, ForceNew) Specifies the policy ID associated with the domain. If not specified, a new policy will be created automatically. Changing this create a new domain.`,
				},
				resource.Attribute{
					Name:        "keep_proxy",
					Description: `(Optional, Bool) Specifies whether to retain the policy when deleting a domain name. Defaults to true.`,
				},
				resource.Attribute{
					Name:        "proxy",
					Description: `(Optional, Bool) Specifies whether a proxy is configured.`,
				},
				resource.Attribute{
					Name:        "sip_header_name",
					Description: `(Optional, String) Specifies the type of the source IP header. This parameter is required only when proxy is set to true. The options are as follows:`,
				},
				resource.Attribute{
					Name:        "sip_header_list",
					Description: `(Optional, List) Specifies an array of HTTP request header for identifying the real source IP address. This parameter is required only when proxy is set to true. - If ` + "`" + `sip_header_name` + "`" + ` is`,
				},
				resource.Attribute{
					Name:        "client_protocol",
					Description: `(Required, String) Protocol type of the client. The options are`,
				},
				resource.Attribute{
					Name:        "server_protocol",
					Description: `(Required, String) Protocol used by WAF to forward client requests to the server. The options are`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required, String) IP address or domain name of the web server that the client accesses. For example, 192.168.1.1 or www.a.com.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required, Int) Port number used by the web server. The value ranges from 0 to 65535, for example, 8080. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the domain.`,
				},
				resource.Attribute{
					Name:        "cname",
					Description: `The CNAME value.`,
				},
				resource.Attribute{
					Name:        "txt_code",
					Description: `The TXT record. This attribute is returned only when proxy is set to true.`,
				},
				resource.Attribute{
					Name:        "sub_domain",
					Description: `The subdomain name. This attribute is returned only when proxy is set to true.`,
				},
				resource.Attribute{
					Name:        "protect_status",
					Description: `The WAF mode. -1: bypassed, 0: disabled, 1: enabled.`,
				},
				resource.Attribute{
					Name:        "access_status",
					Description: `Whether a domain name is connected to WAF. 0: The domain name is not connected to WAF, 1: The domain name is connected to WAF.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol type of the client. The options are HTTP, HTTPS, and HTTP&HTTPS. ## Import Domains can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `sh terraform import flexibleengine_waf_domain.dom_1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the domain.`,
				},
				resource.Attribute{
					Name:        "cname",
					Description: `The CNAME value.`,
				},
				resource.Attribute{
					Name:        "txt_code",
					Description: `The TXT record. This attribute is returned only when proxy is set to true.`,
				},
				resource.Attribute{
					Name:        "sub_domain",
					Description: `The subdomain name. This attribute is returned only when proxy is set to true.`,
				},
				resource.Attribute{
					Name:        "protect_status",
					Description: `The WAF mode. -1: bypassed, 0: disabled, 1: enabled.`,
				},
				resource.Attribute{
					Name:        "access_status",
					Description: `Whether a domain name is connected to WAF. 0: The domain name is not connected to WAF, 1: The domain name is connected to WAF.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol type of the client. The options are HTTP, HTTPS, and HTTP&HTTPS. ## Import Domains can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `sh terraform import flexibleengine_waf_domain.dom_1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_waf_policy",
			Category:         "Web Application Firewall (WAF)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"web",
				"application",
				"firewall",
				"waf",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required, String) Specifies the policy name. The maximum length is 256 characters. Only digits, letters, underscores(_), and hyphens(-) are allowed.`,
				},
				resource.Attribute{
					Name:        "protection_mode",
					Description: `(Optional, String) Specifies the protective action after a rule is matched. Valid values are: -`,
				},
				resource.Attribute{
					Name:        "level",
					Description: `(Optional, Int) Specifies the protection level. Valid values are: -`,
				},
				resource.Attribute{
					Name:        "full_detection",
					Description: `(Optional, Bool) Specifies the detection mode in Precise Protection. Valid values are:`,
				},
				resource.Attribute{
					Name:        "domains",
					Description: `(Optional, List) An array of domain IDs.`,
				},
				resource.Attribute{
					Name:        "protection_status",
					Description: `(Optional, Object) Specifies the protection switches. The object structure is documented below. The ` + "`" + `protection_status` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "basic_web_protection",
					Description: `Specifies whether Basic Web Protection is enabled.`,
				},
				resource.Attribute{
					Name:        "general_check",
					Description: `Specifies whether General Check in Basic Web Protection is enabled.`,
				},
				resource.Attribute{
					Name:        "crawler_engine",
					Description: `Specifies whether the Search Engine switch in Basic Web Protection is enabled.`,
				},
				resource.Attribute{
					Name:        "crawler_scanner",
					Description: `Specifies whether the Scanner switch in Basic Web Protection is enabled.`,
				},
				resource.Attribute{
					Name:        "crawler_script",
					Description: `Specifies whether the Script Tool switch in Basic Web Protection is enabled.`,
				},
				resource.Attribute{
					Name:        "crawler_other",
					Description: `Specifies whether detection of other crawlers in Basic Web Protection is enabled.`,
				},
				resource.Attribute{
					Name:        "webshell",
					Description: `Specifies whether webshell detection in Basic Web Protection is enabled.`,
				},
				resource.Attribute{
					Name:        "cc_protection",
					Description: `Specifies whether CC Attack Protection is enabled.`,
				},
				resource.Attribute{
					Name:        "precise_protection",
					Description: `Specifies whether Precise Protection is enabled.`,
				},
				resource.Attribute{
					Name:        "blacklist",
					Description: `Specifies whether Blacklist and Whitelist is enabled.`,
				},
				resource.Attribute{
					Name:        "data_masking",
					Description: `Specifies whether Data Masking is enabled.`,
				},
				resource.Attribute{
					Name:        "false_alarm_masking",
					Description: `Specifies whether False Alarm Masking is enabled.`,
				},
				resource.Attribute{
					Name:        "web_tamper_protection",
					Description: `Specifies whether Web Tamper Protection is enabled. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The policy ID in UUID format. ## Import Policies can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `sh terraform import flexibleengine_waf_policy.policy_1 c5946141e52441d9b13c5e9d4e9560c7 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The policy ID in UUID format. ## Import Policies can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `sh terraform import flexibleengine_waf_policy.policy_1 c5946141e52441d9b13c5e9d4e9560c7 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_waf_rule_alarm_masking",
			Category:         "Web Application Firewall (WAF)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"web",
				"application",
				"firewall",
				"waf",
				"rule",
				"alarm",
				"masking",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Required, String, ForceNew) Specifies the WAF policy ID. Changing this creates a new rule.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required, String) Specifies a misreported URL excluding a domain name.`,
				},
				resource.Attribute{
					Name:        "event_id",
					Description: `(Required, String) Specifies the event ID. It is the ID of a misreported event in Events whose type is not`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The rule ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "event_type",
					Description: `The event type. ## Import Alarm Masking Rules can be imported using the policy ID and rule ID separated by a slash, e.g.: ` + "`" + `` + "`" + `` + "`" + `sh terraform import flexibleengine_waf_rule_alarm_masking.rule_1 44d887434169475794b2717438f7fa78/6cdc116040d444f6b3e1bf1dd629f1d0 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The rule ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "event_type",
					Description: `The event type. ## Import Alarm Masking Rules can be imported using the policy ID and rule ID separated by a slash, e.g.: ` + "`" + `` + "`" + `` + "`" + `sh terraform import flexibleengine_waf_rule_alarm_masking.rule_1 44d887434169475794b2717438f7fa78/6cdc116040d444f6b3e1bf1dd629f1d0 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_waf_rule_blacklist",
			Category:         "Web Application Firewall (WAF)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"web",
				"application",
				"firewall",
				"waf",
				"rule",
				"blacklist",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Required, String, ForceNew) Specifies the WAF policy ID. Changing this creates a new rule.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required, String) Specifies the IP address or range. For example, 192.168.0.125 or 192.168.0.0/24.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional, Int) Specifies the protective action. 1: Whitelist, 0: Blacklist. If you do not configure the parameter, the value is Blacklist by default. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The rule ID in UUID format. ## Import Blacklist Rules can be imported using the policy ID and rule ID separated by a slash, e.g.: ` + "`" + `` + "`" + `` + "`" + `sh terraform import flexibleengine_waf_rule_blacklist.rule_1 523083f4543c497faecd25fcfcc0b2a0/e7f49f736bc74b828ce45e0e5c49d156 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The rule ID in UUID format. ## Import Blacklist Rules can be imported using the policy ID and rule ID separated by a slash, e.g.: ` + "`" + `` + "`" + `` + "`" + `sh terraform import flexibleengine_waf_rule_blacklist.rule_1 523083f4543c497faecd25fcfcc0b2a0/e7f49f736bc74b828ce45e0e5c49d156 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_waf_rule_cc_protection",
			Category:         "Web Application Firewall (WAF)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"web",
				"application",
				"firewall",
				"waf",
				"rule",
				"cc",
				"protection",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Required, String, ForceNew) Specifies the WAF policy ID. Changing this creates a new rule.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required, String) Specifies the URL to which the rule applies. The path ending with`,
				},
				resource.Attribute{
					Name:        "limit_num",
					Description: `(Required, Int) Specifies the number of requests allowed from a web visitor in a rate limiting period. The value ranges from 0 to 2^32.`,
				},
				resource.Attribute{
					Name:        "limit_period",
					Description: `(Required, Int) Specifies the rate limiting period. The value ranges from 0 seconds to 2^32 seconds.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Required, String) Specifies the rate limit mode. Valid Options are:`,
				},
				resource.Attribute{
					Name:        "cookie",
					Description: `(Optional, String) Specifies the cookie name. This field is mandatory when ` + "`" + `mode` + "`" + ` is set to`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Optional, String) Specifies the category content. The format is as follows: http://www.example.com/path. This field is mandatory when ` + "`" + `mode` + "`" + ` is set to`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required, String) Specifies the action when the number of requests reaches the upper limit. Valid Options are:`,
				},
				resource.Attribute{
					Name:        "block_time",
					Description: `(Optional, Int) Specifies the lock duration. The value ranges from 0 seconds to 2^32 seconds.`,
				},
				resource.Attribute{
					Name:        "block_page_type",
					Description: `(Optional, String) Specifies the type of the returned page. The options are ` + "`" + `application/json` + "`" + `, ` + "`" + `text/html` + "`" + `, and ` + "`" + `text/xml` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "block_page_content",
					Description: `(Optional, String) Specifies the content of the returned page. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The rule ID in UUID format. ## Import CC Attack Protection Rules can be imported using the policy ID and rule ID separated by a slash, e.g.: ` + "`" + `` + "`" + `` + "`" + `sh terraform import flexibleengine_waf_rule_cc_protection.rule_1 523083f4543c497faecd25fcfcc0b2a0/dd3c14e91550453f81cff5fc3b7c3e89 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The rule ID in UUID format. ## Import CC Attack Protection Rules can be imported using the policy ID and rule ID separated by a slash, e.g.: ` + "`" + `` + "`" + `` + "`" + `sh terraform import flexibleengine_waf_rule_cc_protection.rule_1 523083f4543c497faecd25fcfcc0b2a0/dd3c14e91550453f81cff5fc3b7c3e89 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_waf_rule_data_masking",
			Category:         "Web Application Firewall (WAF)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"web",
				"application",
				"firewall",
				"waf",
				"rule",
				"data",
				"masking",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Required, String, ForceNew) Specifies the WAF policy ID. Changing this creates a new rule.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required, String) Specifies the URL to which the data masking rule applies (exact match by default).`,
				},
				resource.Attribute{
					Name:        "field",
					Description: `(Required, String) Specifies the masked field. The options are`,
				},
				resource.Attribute{
					Name:        "subfield",
					Description: `(Required, String) Specifies the masked subfield. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The rule ID in UUID format. ## Import Data Masking Rules can be imported using the policy ID and rule ID separated by a slash, e.g.: ` + "`" + `` + "`" + `` + "`" + `sh terraform import flexibleengine_waf_rule_data_masking.rule_1 523083f4543c497faecd25fcfcc0b2a0/c6482bd0059148559b625f78e8ce92be ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The rule ID in UUID format. ## Import Data Masking Rules can be imported using the policy ID and rule ID separated by a slash, e.g.: ` + "`" + `` + "`" + `` + "`" + `sh terraform import flexibleengine_waf_rule_data_masking.rule_1 523083f4543c497faecd25fcfcc0b2a0/c6482bd0059148559b625f78e8ce92be ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_waf_rule_precise_protection",
			Category:         "Web Application Firewall (WAF)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"web",
				"application",
				"firewall",
				"waf",
				"rule",
				"precise",
				"protection",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Required, String, ForceNew) Specifies the WAF policy ID. Changing this creates a new rule.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, String) Specifies the name of a precise protection rule. The maximum length is 256 characters. Only digits, letters, underscores (_), and hyphens (-) are allowed.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional, String) Specifies the protective action after the precise protection rule is matched. The value can be`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Required, Int) Specifies the priority of a rule being executed. Smaller values correspond to higher priorities. If two rules are assigned with the same priority, the rule added earlier has higher priority, the rule added earlier has higher priority. The value ranges from 0 to 65535.`,
				},
				resource.Attribute{
					Name:        "conditions",
					Description: `(Required, List) Specifies the condition parameters. The object structure is documented below.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `(Optional, String) Specifies the UTC time when the precise protection rule takes effect. The time must be in "yyyy-MM-dd HH:mm:ss" format. If not specified, the rule takes effect immediately.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `(Optional, String) Specifies the UTC time when the precise protection rule expires. The time must be in "yyyy-MM-dd HH:mm:ss" format. The ` + "`" + `conditions` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "field",
					Description: `(Required, String) Specifies the matched field. The value can be`,
				},
				resource.Attribute{
					Name:        "subfield",
					Description: `(Optional, String) Specifies the matched subfield. - If ` + "`" + `field` + "`" + ` is set to`,
				},
				resource.Attribute{
					Name:        "logic",
					Description: `(Required, String) Specifies the logic relationship. The value can be`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Required, String) Specifies the content matching the condition. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The rule ID in UUID format. ## Import Precise Protection Rules can be imported using the policy ID and rule ID separated by a slash, e.g.: ` + "`" + `` + "`" + `` + "`" + `sh terraform import flexibleengine_waf_rule_precise_protection.rule_1 523083f4543c497faecd25fcfcc0b2a0/620801321b254f8fbc7dafa6bbebe652 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The rule ID in UUID format. ## Import Precise Protection Rules can be imported using the policy ID and rule ID separated by a slash, e.g.: ` + "`" + `` + "`" + `` + "`" + `sh terraform import flexibleengine_waf_rule_precise_protection.rule_1 523083f4543c497faecd25fcfcc0b2a0/620801321b254f8fbc7dafa6bbebe652 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_waf_rule_web_tamper_protection",
			Category:         "Web Application Firewall (WAF)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"web",
				"application",
				"firewall",
				"waf",
				"rule",
				"tamper",
				"protection",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Required, String, ForceNew) Specifies the WAF policy ID. Changing this creates a new rule.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Required, String, ForceNew) Specifies the domain name. Changing this creates a new rule.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required, String, ForceNew) Specifies the URL protected by the web tamper protection rule, excluding a domain name. Changing this creates a new rule. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The rule ID in UUID format. ## Import Web Tamper Protection Rules can be imported using the policy ID and rule ID separated by a slash, e.g.: ` + "`" + `` + "`" + `` + "`" + `sh terraform import flexibleengine_waf_rule_web_tamper_protection.rule_1 523083f4543c497faecd25fcfcc0b2a0/5b3b07fedc3642d18e424b2e45aebc8a ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The rule ID in UUID format. ## Import Web Tamper Protection Rules can be imported using the policy ID and rule ID separated by a slash, e.g.: ` + "`" + `` + "`" + `` + "`" + `sh terraform import flexibleengine_waf_rule_web_tamper_protection.rule_1 523083f4543c497faecd25fcfcc0b2a0/5b3b07fedc3642d18e424b2e45aebc8a ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"flexibleengine_antiddos_v1":                        0,
		"flexibleengine_api_gateway_api":                    1,
		"flexibleengine_api_gateway_group":                  2,
		"flexibleengine_as_configuration_v1":                3,
		"flexibleengine_as_group_v1":                        4,
		"flexibleengine_as_lifecycle_hook_v1":               5,
		"flexibleengine_as_policy_v1":                       6,
		"flexibleengine_blockstorage_volume_v2":             7,
		"flexibleengine_cbr_policy":                         8,
		"flexibleengine_cbr_vault":                          9,
		"flexibleengine_cce_addon_v3":                       10,
		"flexibleengine_cce_cluster_v3":                     11,
		"flexibleengine_cce_namespace":                      12,
		"flexibleengine_cce_node_pool_v3":                   13,
		"flexibleengine_cce_node_v3":                        14,
		"flexibleengine_cce_pvc":                            15,
		"flexibleengine_ces_alarmrule":                      16,
		"flexibleengine_compute_bms_server_v2":              17,
		"flexibleengine_compute_floatingip_associate_v2":    18,
		"flexibleengine_compute_floatingip_v2":              19,
		"flexibleengine_compute_instance_v2":                20,
		"flexibleengine_compute_interface_attach_v2":        21,
		"flexibleengine_compute_keypair_v2":                 22,
		"flexibleengine_compute_servergroup_v2":             23,
		"flexibleengine_compute_volume_attach_v2":           24,
		"flexibleengine_csbs_backup_policy_v1":              25,
		"flexibleengine_csbs_backup_v1":                     26,
		"flexibleengine_css_cluster_v1":                     27,
		"flexibleengine_css_snapshot_v1":                    28,
		"flexibleengine_cts_tracker_v1":                     29,
		"flexibleengine_dcs_instance_v1":                    30,
		"flexibleengine_dds_instance_v3":                    31,
		"flexibleengine_dis_stream":                         32,
		"flexibleengine_dli_queue":                          33,
		"flexibleengine_dms_kafka_instance":                 34,
		"flexibleengine_dms_kafka_topic":                    35,
		"flexibleengine_dns_ptrrecord_v2":                   36,
		"flexibleengine_dns_recordset_v2":                   37,
		"flexibleengine_dns_zone_v2":                        38,
		"flexibleengine_drs_replication_v2":                 39,
		"flexibleengine_drs_replicationconsistencygroup_v2": 40,
		"flexibleengine_dws_cluster_v1":                     41,
		"flexibleengine_elb_backend":                        42,
		"flexibleengine_elb_certificate":                    43,
		"flexibleengine_elb_health":                         44,
		"flexibleengine_elb_ipgroup":                        45,
		"flexibleengine_elb_listener":                       46,
		"flexibleengine_elb_loadbalancer":                   47,
		"flexibleengine_enterprise_project":                 48,
		"flexibleengine_fgs_dependency":                     49,
		"flexibleengine_fgs_function":                       50,
		"flexibleengine_fgs_trigger":                        51,
		"flexibleengine_fw_firewall_group_v2":               52,
		"flexibleengine_fw_policy_v2":                       53,
		"flexibleengine_fw_rule_v2":                         54,
		"flexibleengine_identity_agency_v3":                 55,
		"flexibleengine_identity_group_membership_v3":       56,
		"flexibleengine_identity_group_v3":                  57,
		"flexibleengine_identity_project_v3":                58,
		"flexibleengine_identity_provider":                  59,
		"flexibleengine_identity_provider_conversion":       60,
		"flexibleengine_identity_role_assignment_v3":        61,
		"flexibleengine_identity_role_v3":                   62,
		"flexibleengine_identity_user_v3":                   63,
		"flexibleengine_images_image_v2":                    64,
		"flexibleengine_kms_key_v1":                         65,
		"flexibleengine_lb_certificate_v2":                  66,
		"flexibleengine_lb_l7policy_v2":                     67,
		"flexibleengine_lb_l7rule_v2":                       68,
		"flexibleengine_lb_listener_v2":                     69,
		"flexibleengine_lb_listener_v3":                     70,
		"flexibleengine_lb_loadbalancer_v2":                 71,
		"flexibleengine_lb_loadbalancer_v3":                 72,
		"flexibleengine_lb_member_v2":                       73,
		"flexibleengine_lb_monitor_v2":                      74,
		"flexibleengine_lb_pool_v2":                         75,
		"flexibleengine_lb_whitelist_v2":                    76,
		"flexibleengine_lts_group":                          77,
		"flexibleengine_lts_topic":                          78,
		"flexibleengine_mls_instance_v1":                    79,
		"flexibleengine_mrs_cluster_v1":                     80,
		"flexibleengine_mrs_cluster_v2":                     81,
		"flexibleengine_mrs_hybrid_cluster_v1":              82,
		"flexibleengine_mrs_job_v1":                         83,
		"flexibleengine_mrs_job_v2":                         84,
		"flexibleengine_nat_dnat_rule_v2":                   85,
		"flexibleengine_nat_gateway_v2":                     86,
		"flexibleengine_nat_snat_rule_v2":                   87,
		"flexibleengine_network_acl":                        88,
		"flexibleengine_network_acl_rule":                   89,
		"flexibleengine_networking_floatingip_associate_v2": 90,
		"flexibleengine_networking_floatingip_v2":           91,
		"flexibleengine_networking_network_v2":              92,
		"flexibleengine_networking_port_v2":                 93,
		"flexibleengine_networking_router_interface_v2":     94,
		"flexibleengine_networking_router_route_v2":         95,
		"flexibleengine_networking_router_v2":               96,
		"flexibleengine_networking_secgroup_rule_v2":        97,
		"flexibleengine_networking_secgroup_v2":             98,
		"flexibleengine_networking_subnet_v2":               99,
		"flexibleengine_networking_vip_associate_v2":        100,
		"flexibleengine_networking_vip_v2":                  101,
		"flexibleengine_obs_bucket":                         102,
		"flexibleengine_obs_bucket_object":                  103,
		"flexibleengine_obs_bucket_replication":             104,
		"flexibleengine_rds_account":                        105,
		"flexibleengine_rds_database":                       106,
		"flexibleengine_rds_instance_v1":                    107,
		"flexibleengine_rds_instance_v3":                    108,
		"flexibleengine_rds_parametergroup_v3":              109,
		"flexibleengine_rds_read_replica_v3":                110,
		"flexibleengine_rts_software_config_v1":             111,
		"flexibleengine_rts_stack_v1":                       112,
		"flexibleengine_s3_bucket":                          113,
		"flexibleengine_s3_bucket_object":                   114,
		"flexibleengine_s3_bucket_policy":                   115,
		"flexibleengine_sdrs_drill_v1":                      116,
		"flexibleengine_sdrs_protectedinstance_v1":          117,
		"flexibleengine_sdrs_protectiongroup_v1":            118,
		"flexibleengine_sdrs_replication_attach_v1":         119,
		"flexibleengine_sdrs_replication_pair_v1":           120,
		"flexibleengine_sfs_access_rule_v2":                 121,
		"flexibleengine_sfs_file_system_v2":                 122,
		"flexibleengine_sfs_turbo":                          123,
		"flexibleengine_smn_subscription_v2":                124,
		"flexibleengine_smn_topic_v2":                       125,
		"flexibleengine_swr_organization":                   126,
		"flexibleengine_swr_organization_users":             127,
		"flexibleengine_swr_repository":                     128,
		"flexibleengine_swr_repository_sharing":             129,
		"flexibleengine_vbs_backup_policy_v2":               130,
		"flexibleengine_vbs_backup_v2":                      131,
		"flexibleengine_vpc_eip_v1":                         132,
		"flexibleengine_vpc_flow_log_v1":                    133,
		"flexibleengine_vpc_peering_accepter_v2":            134,
		"flexibleengine_vpc_peering_v2":                     135,
		"flexibleengine_vpc_route_v2":                       136,
		"flexibleengine_vpc_subnet_v1":                      137,
		"flexibleengine_vpc_v1":                             138,
		"flexibleengine_vpcep_approval":                     139,
		"flexibleengine_vpcep_endpoint":                     140,
		"flexibleengine_vpcep_service":                      141,
		"flexibleengine_waf_certificate":                    142,
		"flexibleengine_waf_domain":                         143,
		"flexibleengine_waf_policy":                         144,
		"flexibleengine_waf_rule_alarm_masking":             145,
		"flexibleengine_waf_rule_blacklist":                 146,
		"flexibleengine_waf_rule_cc_protection":             147,
		"flexibleengine_waf_rule_data_masking":              148,
		"flexibleengine_waf_rule_precise_protection":        149,
		"flexibleengine_waf_rule_web_tamper_protection":     150,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
