package tencentcloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_address_template",
			Category:         "Virtual Private Cloud(VPC)",
			ShortDescription: `Provides a resource to manage address template.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"address",
				"template",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "addresses",
					Description: `(Required) Address list. IP(` + "`" + `10.0.0.1` + "`" + `), CIDR(` + "`" + `10.0.1.0/24` + "`" + `), IP range(` + "`" + `10.0.0.1-10.0.0.100` + "`" + `) format are supported.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Name of the address template. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import Address template can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_address_template.foo ipm-makf7k9e" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import Address template can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_address_template.foo ipm-makf7k9e" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_address_template_group",
			Category:         "Virtual Private Cloud(VPC)",
			ShortDescription: `Provides a resource to manage address template group.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"address",
				"template",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Name of the address template group.`,
				},
				resource.Attribute{
					Name:        "template_ids",
					Description: `(Required) Template ID list. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import Address template group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_address_template_group.foo ipmg-0np3u974 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import Address template group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_address_template_group.foo ipmg-0np3u974 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_alb_server_attachment",
			Category:         "Cloud Load Balancer(CLB)",
			ShortDescription: `Provides an tencentcloud application load balancer servers attachment as a resource, to attach and detach instances from load balancer.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"load",
				"balancer",
				"clb",
				"alb",
				"server",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "backends",
					Description: `(Required) list of backend server.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Required, ForceNew) listener ID.`,
				},
				resource.Attribute{
					Name:        "loadbalancer_id",
					Description: `(Required, ForceNew) loadbalancer ID.`,
				},
				resource.Attribute{
					Name:        "location_id",
					Description: `(Optional, ForceNew) location ID, only support for layer 7 loadbalancer. The ` + "`" + `backends` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) A list backend instance ID (CVM instance ID).`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) The port used by the backend server. Valid value range: [1-65535].`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) Weight of the backend server. Valid value range: [0-100]. Default to 10. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "protocol_type",
					Description: `The protocol type, http or tcp.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "protocol_type",
					Description: `The protocol type, http or tcp.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_api_gateway_api",
			Category:         "API GateWay",
			ShortDescription: `Use this resource to create API of API gateway.`,
			Description:      ``,
			Keywords: []string{
				"api",
				"gateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "api_name",
					Description: `(Required) Custom API name.`,
				},
				resource.Attribute{
					Name:        "request_config_path",
					Description: `(Required) Request frontend path configuration. Like ` + "`" + `/user/getinfo` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "service_id",
					Description: `(Required, ForceNew) Which service this API belongs. Refer to resource ` + "`" + `tencentcloud_api_gateway_service` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "api_desc",
					Description: `(Optional) Custom API description.`,
				},
				resource.Attribute{
					Name:        "auth_type",
					Description: `(Optional) API authentication type. Valid values: ` + "`" + `SECRET` + "`" + ` (key pair authentication),` + "`" + `NONE` + "`" + ` (no authentication). Default value: ` + "`" + `NONE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enable_cors",
					Description: `(Optional) Whether to enable CORS. Default value: ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "pre_limit",
					Description: `(Optional) API QPS value. Enter a positive number to limit the API query rate per second ` + "`" + `QPS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional, ForceNew) API frontend request type. Valid values: ` + "`" + `HTTP` + "`" + `, ` + "`" + `WEBSOCKET` + "`" + `. Default value: ` + "`" + `HTTP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "release_limit",
					Description: `(Optional) API QPS value. Enter a positive number to limit the API query rate per second ` + "`" + `QPS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "request_config_method",
					Description: `(Optional) Request frontend method configuration. Valid values: ` + "`" + `GET` + "`" + `,` + "`" + `POST` + "`" + `,` + "`" + `PUT` + "`" + `,` + "`" + `DELETE` + "`" + `,` + "`" + `HEAD` + "`" + `,` + "`" + `ANY` + "`" + `. Default value: ` + "`" + `GET` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "request_parameters",
					Description: `(Optional) Frontend request parameters.`,
				},
				resource.Attribute{
					Name:        "response_error_codes",
					Description: `(Optional) Custom error code configuration. Must keep at least one after set.`,
				},
				resource.Attribute{
					Name:        "response_fail_example",
					Description: `(Optional) Response failure sample of custom response configuration.`,
				},
				resource.Attribute{
					Name:        "response_success_example",
					Description: `(Optional) Successful response sample of custom response configuration.`,
				},
				resource.Attribute{
					Name:        "response_type",
					Description: `(Optional) Return type. Valid values: ` + "`" + `HTML` + "`" + `, ` + "`" + `JSON` + "`" + `, ` + "`" + `TEXT` + "`" + `, ` + "`" + `BINARY` + "`" + `, ` + "`" + `XML` + "`" + `. Default value: ` + "`" + `HTML` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "service_config_method",
					Description: `(Optional) API backend service request method, such as ` + "`" + `GET` + "`" + `. If ` + "`" + `service_config_type` + "`" + ` is ` + "`" + `HTTP` + "`" + `, this parameter will be required. The frontend ` + "`" + `request_config_method` + "`" + ` and backend method ` + "`" + `service_config_method` + "`" + ` can be different.`,
				},
				resource.Attribute{
					Name:        "service_config_mock_return_message",
					Description: `(Optional) Returned information of API backend mocking. This parameter is required when ` + "`" + `service_config_type` + "`" + ` is ` + "`" + `MOCK` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "service_config_path",
					Description: `(Optional) API backend service path, such as /path. If ` + "`" + `service_config_type` + "`" + ` is ` + "`" + `HTTP` + "`" + `, this parameter will be required. The frontend ` + "`" + `request_config_path` + "`" + ` and backend path ` + "`" + `service_config_path` + "`" + ` can be different.`,
				},
				resource.Attribute{
					Name:        "service_config_product",
					Description: `(Optional) Backend type. This parameter takes effect when VPC is enabled. Currently, only ` + "`" + `clb` + "`" + ` is supported.`,
				},
				resource.Attribute{
					Name:        "service_config_scf_function_name",
					Description: `(Optional) SCF function name. This parameter takes effect when ` + "`" + `service_config_type` + "`" + ` is ` + "`" + `SCF` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "service_config_scf_function_namespace",
					Description: `(Optional) SCF function namespace. This parameter takes effect when ` + "`" + `service_config_type` + "`" + ` is ` + "`" + `SCF` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "service_config_scf_function_qualifier",
					Description: `(Optional) SCF function version. This parameter takes effect when ` + "`" + `service_config_type` + "`" + ` is ` + "`" + `SCF` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "service_config_timeout",
					Description: `(Optional) API backend service timeout period in seconds. Default value: ` + "`" + `5` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "service_config_type",
					Description: `(Optional) API backend service type. Valid values: ` + "`" + `WEBSOCKET` + "`" + `, ` + "`" + `HTTP` + "`" + `, ` + "`" + `SCF` + "`" + `, ` + "`" + `MOCK` + "`" + `. Default value: ` + "`" + `HTTP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "service_config_url",
					Description: `(Optional) API backend service url. This parameter is required when ` + "`" + `service_config_type` + "`" + ` is ` + "`" + `HTTP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "service_config_vpc_id",
					Description: `(Optional) Unique VPC ID.`,
				},
				resource.Attribute{
					Name:        "test_limit",
					Description: `(Optional) API QPS value. Enter a positive number to limit the API query rate per second ` + "`" + `QPS` + "`" + `. The ` + "`" + `request_parameters` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Parameter name.`,
				},
				resource.Attribute{
					Name:        "position",
					Description: `(Required) Parameter location.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Parameter type.`,
				},
				resource.Attribute{
					Name:        "default_value",
					Description: `(Optional) Parameter default value.`,
				},
				resource.Attribute{
					Name:        "desc",
					Description: `(Optional) Parameter description.`,
				},
				resource.Attribute{
					Name:        "required",
					Description: `(Optional) If this parameter required. Default value: ` + "`" + `false` + "`" + `. The ` + "`" + `response_error_codes` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `(Required) Custom response configuration error code.`,
				},
				resource.Attribute{
					Name:        "msg",
					Description: `(Required) Custom response configuration error message.`,
				},
				resource.Attribute{
					Name:        "converted_code",
					Description: `(Optional) Custom error code conversion.`,
				},
				resource.Attribute{
					Name:        "desc",
					Description: `(Optional) Parameter description.`,
				},
				resource.Attribute{
					Name:        "need_convert",
					Description: `(Optional) Whether to enable error code conversion. Default value: ` + "`" + `false` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_api_gateway_api_key",
			Category:         "API GateWay",
			ShortDescription: `Use this resource to create API gateway access key.`,
			Description:      ``,
			Keywords: []string{
				"api",
				"gateway",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "secret_name",
					Description: `(Required, ForceNew) Custom key name.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Key status. Valid values: ` + "`" + `on` + "`" + `, ` + "`" + `off` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "access_key_secret",
					Description: `Created API key.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used. ## Import API gateway access key can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_api_gateway_api_key.test AKIDMZwceezso9ps5p8jkro8a9fwe1e7nzF2k50B ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "access_key_secret",
					Description: `Created API key.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used. ## Import API gateway access key can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_api_gateway_api_key.test AKIDMZwceezso9ps5p8jkro8a9fwe1e7nzF2k50B ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_api_gateway_api_key_attachment",
			Category:         "API GateWay",
			ShortDescription: `Use this resource to API gateway attach access key to usage plan.`,
			Description:      ``,
			Keywords: []string{
				"api",
				"gateway",
				"key",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "api_key_id",
					Description: `(Required, ForceNew) ID of API key.`,
				},
				resource.Attribute{
					Name:        "usage_plan_id",
					Description: `(Required, ForceNew) ID of the usage plan. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import API gateway attach access key can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_api_gateway_api_key_attachment.attach AKID110b8Rmuw7t0fP1N8bi809n327023Is7xN8f#usagePlan-gyeafpab ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import API gateway attach access key can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_api_gateway_api_key_attachment.attach AKID110b8Rmuw7t0fP1N8bi809n327023Is7xN8f#usagePlan-gyeafpab ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_api_gateway_custom_domain",
			Category:         "API GateWay",
			ShortDescription: `Use this resource to create custom domain of API gateway.`,
			Description:      ``,
			Keywords: []string{
				"api",
				"gateway",
				"custom",
				"domain",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "default_domain",
					Description: `(Required) Default domain name.`,
				},
				resource.Attribute{
					Name:        "net_type",
					Description: `(Required) Network type. Valid values: ` + "`" + `OUTER` + "`" + `, ` + "`" + `INNER` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Protocol supported by service. Valid values: ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + `, ` + "`" + `http&https` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "service_id",
					Description: `(Required, ForceNew) Unique service ID.`,
				},
				resource.Attribute{
					Name:        "sub_domain",
					Description: `(Required) Custom domain name to be bound.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `(Optional) Unique certificate ID of the custom domain name to be bound. You can choose to upload for the ` + "`" + `protocol` + "`" + ` attribute value ` + "`" + `https` + "`" + ` or ` + "`" + `http&https` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "is_default_mapping",
					Description: `(Optional) Whether the default path mapping is used. The default value is ` + "`" + `true` + "`" + `. When it is ` + "`" + `false` + "`" + `, it means custom path mapping. In this case, the ` + "`" + `path_mappings` + "`" + ` attribute is required.`,
				},
				resource.Attribute{
					Name:        "path_mappings",
					Description: `(Optional) Custom domain name path mapping. The data format is: ` + "`" + `path#environment` + "`" + `. Optional values for the environment are ` + "`" + `test` + "`" + `, ` + "`" + `prepub` + "`" + `, and ` + "`" + `release` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Domain name resolution status. ` + "`" + `1` + "`" + ` means normal analysis, ` + "`" + `0` + "`" + ` means parsing failed.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Domain name resolution status. ` + "`" + `1` + "`" + ` means normal analysis, ` + "`" + `0` + "`" + ` means parsing failed.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_api_gateway_ip_strategy",
			Category:         "API GateWay",
			ShortDescription: `Use this resource to create IP strategy of API gateway.`,
			Description:      ``,
			Keywords: []string{
				"api",
				"gateway",
				"ip",
				"strategy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_id",
					Description: `(Required, ForceNew) The ID of the API gateway service.`,
				},
				resource.Attribute{
					Name:        "strategy_data",
					Description: `(Required) IP address data.`,
				},
				resource.Attribute{
					Name:        "strategy_name",
					Description: `(Required, ForceNew) User defined strategy name.`,
				},
				resource.Attribute{
					Name:        "strategy_type",
					Description: `(Required, ForceNew) Blacklist or whitelist. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.`,
				},
				resource.Attribute{
					Name:        "strategy_id",
					Description: `IP policy ID. ## Import IP strategy of API gateway can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_api_gateway_ip_strategy.test service-ohxqslqe#IPStrategy-q1lk8ud2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.`,
				},
				resource.Attribute{
					Name:        "strategy_id",
					Description: `IP policy ID. ## Import IP strategy of API gateway can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_api_gateway_ip_strategy.test service-ohxqslqe#IPStrategy-q1lk8ud2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_api_gateway_service",
			Category:         "API GateWay",
			ShortDescription: `Use this resource to create API gateway service.`,
			Description:      ``,
			Keywords: []string{
				"api",
				"gateway",
				"service",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "net_type",
					Description: `(Required) Network type list, which is used to specify the supported network types. Valid values: ` + "`" + `INNER` + "`" + `, ` + "`" + `OUTER` + "`" + `. ` + "`" + `INNER` + "`" + ` indicates access over private network, and ` + "`" + `OUTER` + "`" + ` indicates access over public network.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Service frontend request type. Valid values: ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + `, ` + "`" + `http&https` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) Custom service name.`,
				},
				resource.Attribute{
					Name:        "exclusive_set_name",
					Description: `(Optional, ForceNew) Self-deployed cluster name, which is used to specify the self-deployed cluster where the service is to be created.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `(Optional, ForceNew) IP version number. Valid values: ` + "`" + `IPv4` + "`" + `, ` + "`" + `IPv6` + "`" + `. Default value: ` + "`" + `IPv4` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "pre_limit",
					Description: `(Optional) API QPS value. Enter a positive number to limit the API query rate per second ` + "`" + `QPS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "release_limit",
					Description: `(Optional) API QPS value. Enter a positive number to limit the API query rate per second ` + "`" + `QPS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "service_desc",
					Description: `(Optional) Custom service description.`,
				},
				resource.Attribute{
					Name:        "test_limit",
					Description: `(Optional) API QPS value. Enter a positive number to limit the API query rate per second ` + "`" + `QPS` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "api_list",
					Description: `A list of APIs.`,
				},
				resource.Attribute{
					Name:        "api_desc",
					Description: `Description of the API.`,
				},
				resource.Attribute{
					Name:        "api_id",
					Description: `ID of the API.`,
				},
				resource.Attribute{
					Name:        "api_name",
					Description: `Name of the API.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `Method of the API.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `Path of the API.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.`,
				},
				resource.Attribute{
					Name:        "inner_http_port",
					Description: `Port number for http access over private network.`,
				},
				resource.Attribute{
					Name:        "inner_https_port",
					Description: `Port number for https access over private network.`,
				},
				resource.Attribute{
					Name:        "internal_sub_domain",
					Description: `Private network access subdomain name.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.`,
				},
				resource.Attribute{
					Name:        "outer_sub_domain",
					Description: `Public network access subdomain name.`,
				},
				resource.Attribute{
					Name:        "usage_plan_list",
					Description: `A list of attach usage plans.`,
				},
				resource.Attribute{
					Name:        "api_id",
					Description: `ID of the API.`,
				},
				resource.Attribute{
					Name:        "bind_type",
					Description: `Binding type.`,
				},
				resource.Attribute{
					Name:        "usage_plan_id",
					Description: `ID of the usage plan.`,
				},
				resource.Attribute{
					Name:        "usage_plan_name",
					Description: `Name of the usage plan. ## Import API gateway service can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_api_gateway_service.service service-pg6ud8pa ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "api_list",
					Description: `A list of APIs.`,
				},
				resource.Attribute{
					Name:        "api_desc",
					Description: `Description of the API.`,
				},
				resource.Attribute{
					Name:        "api_id",
					Description: `ID of the API.`,
				},
				resource.Attribute{
					Name:        "api_name",
					Description: `Name of the API.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `Method of the API.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `Path of the API.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.`,
				},
				resource.Attribute{
					Name:        "inner_http_port",
					Description: `Port number for http access over private network.`,
				},
				resource.Attribute{
					Name:        "inner_https_port",
					Description: `Port number for https access over private network.`,
				},
				resource.Attribute{
					Name:        "internal_sub_domain",
					Description: `Private network access subdomain name.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.`,
				},
				resource.Attribute{
					Name:        "outer_sub_domain",
					Description: `Public network access subdomain name.`,
				},
				resource.Attribute{
					Name:        "usage_plan_list",
					Description: `A list of attach usage plans.`,
				},
				resource.Attribute{
					Name:        "api_id",
					Description: `ID of the API.`,
				},
				resource.Attribute{
					Name:        "bind_type",
					Description: `Binding type.`,
				},
				resource.Attribute{
					Name:        "usage_plan_id",
					Description: `ID of the usage plan.`,
				},
				resource.Attribute{
					Name:        "usage_plan_name",
					Description: `Name of the usage plan. ## Import API gateway service can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_api_gateway_service.service service-pg6ud8pa ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_api_gateway_service_release",
			Category:         "API GateWay",
			ShortDescription: `Use this resource to create API gateway service release.`,
			Description:      ``,
			Keywords: []string{
				"api",
				"gateway",
				"service",
				"release",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "environment_name",
					Description: `(Required, ForceNew) API gateway service environment name to be released. Valid values: ` + "`" + `test` + "`" + `, ` + "`" + `prepub` + "`" + `, ` + "`" + `release` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "release_desc",
					Description: `(Required, ForceNew) This release description of the API gateway service.`,
				},
				resource.Attribute{
					Name:        "service_id",
					Description: `(Required, ForceNew) ID of API gateway service.`,
				},
				resource.Attribute{
					Name:        "release_version",
					Description: `(Optional) The release version. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import API gateway service release can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_api_gateway_service_release.service service-jjt3fs3s#release#20201015121916d85fb161-eaec-4dda-a7e0-659aa5f401be ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import API gateway service release can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_api_gateway_service_release.service service-jjt3fs3s#release#20201015121916d85fb161-eaec-4dda-a7e0-659aa5f401be ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_api_gateway_strategy_attachment",
			Category:         "API GateWay",
			ShortDescription: `Use this resource to create IP strategy attachment of API gateway.`,
			Description:      ``,
			Keywords: []string{
				"api",
				"gateway",
				"strategy",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bind_api_id",
					Description: `(Required, ForceNew) The API that needs to be bound.`,
				},
				resource.Attribute{
					Name:        "environment_name",
					Description: `(Required, ForceNew) The environment of the strategy association. Valid values: ` + "`" + `test` + "`" + `, ` + "`" + `release` + "`" + `, ` + "`" + `prepub` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "service_id",
					Description: `(Required, ForceNew) The ID of the API gateway service.`,
				},
				resource.Attribute{
					Name:        "strategy_id",
					Description: `(Required, ForceNew) The ID of the API gateway strategy. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import IP strategy attachment of API gateway can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_api_gateway_strategy_attachment.test service-pk2r6bcc#IPStrategy-4kz2ljfi#api-h3wc5r0s#release ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import IP strategy attachment of API gateway can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_api_gateway_strategy_attachment.test service-pk2r6bcc#IPStrategy-4kz2ljfi#api-h3wc5r0s#release ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_api_gateway_usage_plan",
			Category:         "API GateWay",
			ShortDescription: `Use this resource to create API gateway usage plan.`,
			Description:      ``,
			Keywords: []string{
				"api",
				"gateway",
				"usage",
				"plan",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "usage_plan_name",
					Description: `(Required) Custom usage plan name.`,
				},
				resource.Attribute{
					Name:        "max_request_num_pre_sec",
					Description: `(Optional) Limit of requests per second. Valid values: -1, [1,2000]. The default value is -1, which indicates no limit.`,
				},
				resource.Attribute{
					Name:        "max_request_num",
					Description: `(Optional) Total number of requests allowed. Valid values: -1, [1,99999999]. The default value is -1, which indicates no limit.`,
				},
				resource.Attribute{
					Name:        "usage_plan_desc",
					Description: `(Optional) Custom usage plan description. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "attach_api_keys",
					Description: `Attach API keys list.`,
				},
				resource.Attribute{
					Name:        "attach_list",
					Description: `Attach service and API list.`,
				},
				resource.Attribute{
					Name:        "api_id",
					Description: `The API ID, this value is empty if attach service.`,
				},
				resource.Attribute{
					Name:        "api_name",
					Description: `The API name, this value is empty if attach service.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `The environment name.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `The API method, this value is empty if attach service.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The API path, this value is empty if attach service.`,
				},
				resource.Attribute{
					Name:        "service_id",
					Description: `The service ID.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `The service name.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used. ## Import API gateway usage plan can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_api_gateway_usage_plan.plan usagePlan-gyeafpab ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "attach_api_keys",
					Description: `Attach API keys list.`,
				},
				resource.Attribute{
					Name:        "attach_list",
					Description: `Attach service and API list.`,
				},
				resource.Attribute{
					Name:        "api_id",
					Description: `The API ID, this value is empty if attach service.`,
				},
				resource.Attribute{
					Name:        "api_name",
					Description: `The API name, this value is empty if attach service.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `The environment name.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `The API method, this value is empty if attach service.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The API path, this value is empty if attach service.`,
				},
				resource.Attribute{
					Name:        "service_id",
					Description: `The service ID.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `The service name.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used. ## Import API gateway usage plan can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_api_gateway_usage_plan.plan usagePlan-gyeafpab ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_api_gateway_usage_plan_attachment",
			Category:         "API GateWay",
			ShortDescription: `Use this resource to attach API gateway usage plan to service.`,
			Description:      ``,
			Keywords: []string{
				"api",
				"gateway",
				"usage",
				"plan",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "environment",
					Description: `(Required, ForceNew) The environment to be bound. Valid values: ` + "`" + `test` + "`" + `, ` + "`" + `prepub` + "`" + `, ` + "`" + `release` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "service_id",
					Description: `(Required, ForceNew) ID of the service.`,
				},
				resource.Attribute{
					Name:        "usage_plan_id",
					Description: `(Required, ForceNew) ID of the usage plan.`,
				},
				resource.Attribute{
					Name:        "api_id",
					Description: `(Optional, ForceNew) ID of the API. This parameter will be required when ` + "`" + `bind_type` + "`" + ` is ` + "`" + `API` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "bind_type",
					Description: `(Optional, ForceNew) Binding type. Valid values: ` + "`" + `API` + "`" + `, ` + "`" + `SERVICE` + "`" + `. Default value is ` + "`" + `SERVICE` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import API gateway usage plan attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_api_gateway_usage_plan_attachment.attach_service usagePlan-pe7fbdgn#service-kuqd6xqk#release#API#api-p8gtanvy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import API gateway usage plan attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_api_gateway_usage_plan_attachment.attach_service usagePlan-pe7fbdgn#service-kuqd6xqk#release#API#api-p8gtanvy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_as_attachment",
			Category:         "Auto Scaling(AS)",
			ShortDescription: `Provides a resource to attach or detach CVM instances to a specified scaling group.`,
			Description:      ``,
			Keywords: []string{
				"auto",
				"scaling",
				"as",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_ids",
					Description: `(Required) ID list of CVM instances to be attached to the scaling group.`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `(Required, ForceNew) ID of a scaling group. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_as_lifecycle_hook",
			Category:         "Auto Scaling(AS)",
			ShortDescription: `Provides a resource for an AS (Auto scaling) lifecycle hook.`,
			Description:      ``,
			Keywords: []string{
				"auto",
				"scaling",
				"as",
				"lifecycle",
				"hook",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "lifecycle_hook_name",
					Description: `(Required) The name of the lifecycle hook.`,
				},
				resource.Attribute{
					Name:        "lifecycle_transition",
					Description: `(Required) The instance state to which you want to attach the lifecycle hook. Valid values: ` + "`" + `INSTANCE_LAUNCHING` + "`" + ` and ` + "`" + `INSTANCE_TERMINATING` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `(Required, ForceNew) ID of a scaling group.`,
				},
				resource.Attribute{
					Name:        "default_result",
					Description: `(Optional) Defines the action the AS group should take when the lifecycle hook timeout elapses or if an unexpected failure occurs. Valid values: ` + "`" + `CONTINUE` + "`" + ` and ` + "`" + `ABANDON` + "`" + `. The default value is ` + "`" + `CONTINUE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "heartbeat_timeout",
					Description: `(Optional) Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. Valid value ranges: (30~3600). and default value is ` + "`" + `300` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "notification_metadata",
					Description: `(Optional) Contains additional information that you want to include any time AS sends a message to the notification target.`,
				},
				resource.Attribute{
					Name:        "notification_queue_name",
					Description: `(Optional) For CMQ_QUEUE type, a name of queue must be set.`,
				},
				resource.Attribute{
					Name:        "notification_target_type",
					Description: `(Optional) Target type. Valid values: ` + "`" + `CMQ_QUEUE` + "`" + `, ` + "`" + `CMQ_TOPIC` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "notification_topic_name",
					Description: `(Optional) For CMQ_TOPIC type, a name of topic must be set. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_as_notification",
			Category:         "Auto Scaling(AS)",
			ShortDescription: `Provides a resource for an AS (Auto scaling) notification.`,
			Description:      ``,
			Keywords: []string{
				"auto",
				"scaling",
				"as",
				"notification",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "notification_types",
					Description: `(Required) A list of Notification Types that trigger notifications. Acceptable values are ` + "`" + `SCALE_OUT_FAILED` + "`" + `, ` + "`" + `SCALE_IN_SUCCESSFUL` + "`" + `, ` + "`" + `SCALE_IN_FAILED` + "`" + `, ` + "`" + `REPLACE_UNHEALTHY_INSTANCE_SUCCESSFUL` + "`" + ` and ` + "`" + `REPLACE_UNHEALTHY_INSTANCE_FAILED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "notification_user_group_ids",
					Description: `(Required) A group of user IDs to be notified.`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `(Required, ForceNew) ID of a scaling group. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_as_scaling_config",
			Category:         "Auto Scaling(AS)",
			ShortDescription: `Provides a resource to create a configuration for an AS (Auto scaling) instance.`,
			Description:      ``,
			Keywords: []string{
				"auto",
				"scaling",
				"as",
				"config",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "configuration_name",
					Description: `(Required) Name of a launch configuration.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Required) An available image ID for a cvm instance.`,
				},
				resource.Attribute{
					Name:        "instance_types",
					Description: `(Required) Specified types of CVM instances.`,
				},
				resource.Attribute{
					Name:        "data_disk",
					Description: `(Optional) Configurations of data disk.`,
				},
				resource.Attribute{
					Name:        "enhanced_monitor_service",
					Description: `(Optional) To specify whether to enable cloud monitor service. Default is ` + "`" + `TRUE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enhanced_security_service",
					Description: `(Optional) To specify whether to enable cloud security service. Default is ` + "`" + `TRUE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_tags",
					Description: `(Optional) A list of tags used to associate different resources.`,
				},
				resource.Attribute{
					Name:        "internet_charge_type",
					Description: `(Optional) Charge types for network traffic. Valid values: ` + "`" + `BANDWIDTH_PREPAID` + "`" + `, ` + "`" + `TRAFFIC_POSTPAID_BY_HOUR` + "`" + `, ` + "`" + `TRAFFIC_POSTPAID_BY_HOUR` + "`" + ` and ` + "`" + `BANDWIDTH_PACKAGE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "internet_max_bandwidth_out",
					Description: `(Optional) Max bandwidth of Internet access in Mbps. Default is ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "keep_image_login",
					Description: `(Optional) Specify whether to keep original settings of a CVM image. And it can't be used with password or key_ids together.`,
				},
				resource.Attribute{
					Name:        "key_ids",
					Description: `(Optional) ID list of keys.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Password to access.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) Specifys to which project the configuration belongs.`,
				},
				resource.Attribute{
					Name:        "public_ip_assigned",
					Description: `(Optional) Specify whether to assign an Internet IP address.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Optional) Security groups to which a CVM instance belongs.`,
				},
				resource.Attribute{
					Name:        "system_disk_size",
					Description: `(Optional) Volume of system disk in GB. Default is ` + "`" + `50` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "system_disk_type",
					Description: `(Optional) Type of a CVM disk. Valid values: ` + "`" + `CLOUD_PREMIUM` + "`" + ` and ` + "`" + `CLOUD_SSD` + "`" + `. Default is ` + "`" + `CLOUD_PREMIUM` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional) ase64-encoded User Data text, the length limit is 16KB. The ` + "`" + `data_disk` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `(Optional) Volume of disk in GB. Default is ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `(Optional) Types of disk. Valid values: ` + "`" + `CLOUD_PREMIUM` + "`" + ` and ` + "`" + `CLOUD_SSD` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) Data disk snapshot ID. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time when the launch configuration was created.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current statues of a launch configuration. ## Import AutoScaling Configuration can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_as_scaling_config.scaling_config asc-n32ymck2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time when the launch configuration was created.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current statues of a launch configuration. ## Import AutoScaling Configuration can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_as_scaling_config.scaling_config asc-n32ymck2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_as_scaling_group",
			Category:         "Auto Scaling(AS)",
			ShortDescription: `Provides a resource to create a group of AS (Auto scaling) instances.`,
			Description:      ``,
			Keywords: []string{
				"auto",
				"scaling",
				"as",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "configuration_id",
					Description: `(Required) An available ID for a launch configuration.`,
				},
				resource.Attribute{
					Name:        "max_size",
					Description: `(Required) Maximum number of CVM instances. Valid value ranges: (0~2000).`,
				},
				resource.Attribute{
					Name:        "min_size",
					Description: `(Required) Minimum number of CVM instances. Valid value ranges: (0~2000).`,
				},
				resource.Attribute{
					Name:        "scaling_group_name",
					Description: `(Required) Name of a scaling group.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) ID of VPC network.`,
				},
				resource.Attribute{
					Name:        "default_cooldown",
					Description: `(Optional) Default cooldown time in second, and default value is ` + "`" + `300` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "desired_capacity",
					Description: `(Optional) Desired volume of CVM instances, which is between ` + "`" + `max_size` + "`" + ` and ` + "`" + `min_size` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "forward_balancer_ids",
					Description: `(Optional) List of application load balancers, which can't be specified with ` + "`" + `load_balancer_ids` + "`" + ` together.`,
				},
				resource.Attribute{
					Name:        "load_balancer_ids",
					Description: `(Optional) ID list of traditional load balancers.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) Specifies to which project the scaling group belongs.`,
				},
				resource.Attribute{
					Name:        "retry_policy",
					Description: `(Optional) Available values for retry policies. Valid values: IMMEDIATE_RETRY and INCREMENTAL_INTERVALS.`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `(Optional) ID list of subnet, and for VPC it is required.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of a scaling group.`,
				},
				resource.Attribute{
					Name:        "termination_policies",
					Description: `(Optional) Available values for termination policies. Valid values: OLDEST_INSTANCE and NEWEST_INSTANCE.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `(Optional) List of available zones, for Basic network it is required. The ` + "`" + `forward_balancer_ids` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Required) Listener ID for application load balancers.`,
				},
				resource.Attribute{
					Name:        "load_balancer_id",
					Description: `(Required) ID of available load balancers.`,
				},
				resource.Attribute{
					Name:        "target_attribute",
					Description: `(Required) Attribute list of target rules.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `(Optional) ID of forwarding rules. The ` + "`" + `target_attribute` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) Port number.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Required) Weight. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time when the AS group was created.`,
				},
				resource.Attribute{
					Name:        "instance_count",
					Description: `Instance number of a scaling group.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of a scaling group. ## Import AutoScaling Groups can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_as_scaling_group.scaling_group asg-n32ymck2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time when the AS group was created.`,
				},
				resource.Attribute{
					Name:        "instance_count",
					Description: `Instance number of a scaling group.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of a scaling group. ## Import AutoScaling Groups can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_as_scaling_group.scaling_group asg-n32ymck2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_as_scaling_policy",
			Category:         "Auto Scaling(AS)",
			ShortDescription: `Provides a resource for an AS (Auto scaling) policy.`,
			Description:      ``,
			Keywords: []string{
				"auto",
				"scaling",
				"as",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "adjustment_type",
					Description: `(Required) Specifies whether the adjustment is an absolute number or a percentage of the current capacity. Valid values: ` + "`" + `CHANGE_IN_CAPACITY` + "`" + `, ` + "`" + `EXACT_CAPACITY` + "`" + ` and ` + "`" + `PERCENT_CHANGE_IN_CAPACITY` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "adjustment_value",
					Description: `(Required) Define the number of instances by which to scale.For ` + "`" + `CHANGE_IN_CAPACITY` + "`" + ` type or PERCENT_CHANGE_IN_CAPACITY, a positive increment adds to the current capacity and a negative value removes from the current capacity. For ` + "`" + `EXACT_CAPACITY` + "`" + ` type, it defines an absolute number of the existing Auto Scaling group size.`,
				},
				resource.Attribute{
					Name:        "comparison_operator",
					Description: `(Required) Comparison operator. Valid values: ` + "`" + `GREATER_THAN` + "`" + `, ` + "`" + `GREATER_THAN_OR_EQUAL_TO` + "`" + `, ` + "`" + `LESS_THAN` + "`" + `, ` + "`" + `LESS_THAN_OR_EQUAL_TO` + "`" + `, ` + "`" + `EQUAL_TO` + "`" + ` and ` + "`" + `NOT_EQUAL_TO` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "continuous_time",
					Description: `(Required) Retry times. Valid value ranges: (1~10).`,
				},
				resource.Attribute{
					Name:        "metric_name",
					Description: `(Required) Name of an indicator. Valid values: ` + "`" + `CPU_UTILIZATION` + "`" + `, ` + "`" + `MEM_UTILIZATION` + "`" + `, ` + "`" + `LAN_TRAFFIC_OUT` + "`" + `, ` + "`" + `LAN_TRAFFIC_IN` + "`" + `, ` + "`" + `WAN_TRAFFIC_OUT` + "`" + ` and ` + "`" + `WAN_TRAFFIC_IN` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Required) Time period in second. Valid values: ` + "`" + `60` + "`" + ` and ` + "`" + `300` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `(Required) Name of a policy used to define a reaction when an alarm is triggered.`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `(Required, ForceNew) ID of a scaling group.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `(Required) Alarm threshold.`,
				},
				resource.Attribute{
					Name:        "cooldown",
					Description: `(Optional) Cooldwon time in second. Default is ` + "`" + `30` + "`" + `0.`,
				},
				resource.Attribute{
					Name:        "notification_user_group_ids",
					Description: `(Optional) An ID group of users to be notified when an alarm is triggered.`,
				},
				resource.Attribute{
					Name:        "statistic",
					Description: `(Optional) Statistic types. Valid values: ` + "`" + `AVERAGE` + "`" + `, ` + "`" + `MAXIMUM` + "`" + ` and ` + "`" + `MINIMUM` + "`" + `. Default is ` + "`" + `AVERAGE` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_as_schedule",
			Category:         "Auto Scaling(AS)",
			ShortDescription: `Provides a resource for an AS (Auto scaling) schedule.`,
			Description:      ``,
			Keywords: []string{
				"auto",
				"scaling",
				"as",
				"schedule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "desired_capacity",
					Description: `(Required) The desired number of CVM instances that should be running in the group.`,
				},
				resource.Attribute{
					Name:        "max_size",
					Description: `(Required) The maximum size for the Auto Scaling group.`,
				},
				resource.Attribute{
					Name:        "min_size",
					Description: `(Required) The minimum size for the Auto Scaling group.`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `(Required, ForceNew) ID of a scaling group.`,
				},
				resource.Attribute{
					Name:        "schedule_action_name",
					Description: `(Required) The name of this scaling action.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `(Required) The time for this action to start, in "YYYY-MM-DDThh:mm:ss+08:00" format (UTC+8).`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `(Optional) The time for this action to end, in "YYYY-MM-DDThh:mm:ss+08:00" format (UTC+8).`,
				},
				resource.Attribute{
					Name:        "recurrence",
					Description: `(Optional) The time when recurring future actions will start. Start time is specified by the user following the Unix cron syntax format. And this argument should be set with end_time together. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_audit",
			Category:         "Audit",
			ShortDescription: `Provides a resource to create an audit.`,
			Description:      ``,
			Keywords: []string{
				"audit",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "audit_switch",
					Description: `(Required) Indicate whether to turn on audit logging or not.`,
				},
				resource.Attribute{
					Name:        "cos_bucket",
					Description: `(Required) Name of the cos bucket to save audit log. Caution: the validation of existing cos bucket will not be checked by terraform.`,
				},
				resource.Attribute{
					Name:        "cos_region",
					Description: `(Required) Region of the cos bucket.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Name of audit. Valid length ranges from 3 to 128. Only alpha character or numbers or '_' supported.`,
				},
				resource.Attribute{
					Name:        "read_write_attribute",
					Description: `(Required) Event attribute filter. Valid values: ` + "`" + `1` + "`" + `, ` + "`" + `2` + "`" + `, ` + "`" + `3` + "`" + `. ` + "`" + `1` + "`" + ` for readonly, ` + "`" + `2` + "`" + ` for write-only, ` + "`" + `3` + "`" + ` for all.`,
				},
				resource.Attribute{
					Name:        "enable_kms_encry",
					Description: `(Optional) Indicate whether the log is encrypt with KMS algorithm or not.`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `(Optional) Existing CMK unique key. This field can be get by data source ` + "`" + `tencentcloud_audit_key_alias` + "`" + `. Caution: the region of the KMS must be as same as the ` + "`" + `cos_region` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "log_file_prefix",
					Description: `(Optional) The log file name prefix. The length ranges from 3 to 40. If not set, the account ID will be the log file prefix. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import Audit can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_audit.foo audit-test ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import Audit can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_audit.foo audit-test ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cam_group",
			Category:         "Cloud Access Management(CAM)",
			ShortDescription: `Provides a resource to create a CAM group.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"access",
				"management",
				"cam",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of CAM group.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `(Optional) Description of the CAM group. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CAM group. ## Import CAM group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cam_group.foo 90496 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CAM group. ## Import CAM group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cam_group.foo 90496 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cam_group_membership",
			Category:         "Cloud Access Management(CAM)",
			ShortDescription: `Provides a resource to create a CAM group membership.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"access",
				"management",
				"cam",
				"group",
				"membership",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required) ID of CAM group.`,
				},
				resource.Attribute{
					Name:        "user_ids",
					Description: `(Required) ID set of the CAM group members. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import CAM group membership can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cam_group_membership.foo 12515263 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import CAM group membership can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cam_group_membership.foo 12515263 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cam_group_policy_attachment",
			Category:         "Cloud Access Management(CAM)",
			ShortDescription: `Provides a resource to create a CAM group policy attachment.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"access",
				"management",
				"cam",
				"group",
				"policy",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required, ForceNew) ID of the attached CAM group.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Required, ForceNew) ID of the policy. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_mode",
					Description: `Mode of Creation of the CAM group policy attachment. ` + "`" + `1` + "`" + ` means the cam policy attachment is created by production, and the others indicate syntax strategy ways.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CAM group policy attachment.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `Name of the policy.`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `Type of the policy strategy. 'Group' means customer strategy and 'QCS' means preset strategy. ## Import CAM group policy attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cam_group_policy_attachment.foo 12515263#26800353 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_mode",
					Description: `Mode of Creation of the CAM group policy attachment. ` + "`" + `1` + "`" + ` means the cam policy attachment is created by production, and the others indicate syntax strategy ways.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CAM group policy attachment.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `Name of the policy.`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `Type of the policy strategy. 'Group' means customer strategy and 'QCS' means preset strategy. ## Import CAM group policy attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cam_group_policy_attachment.foo 12515263#26800353 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cam_policy",
			Category:         "Cloud Access Management(CAM)",
			ShortDescription: `Provides a resource to create a CAM policy.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"access",
				"management",
				"cam",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "document",
					Description: `(Required) Document of the CAM policy. The syntax refers to [CAM POLICY](https://intl.cloud.tencent.com/document/product/598/10604). There are some notes when using this para in terraform: 1. The elements in JSON claimed supporting two types as ` + "`" + `string` + "`" + ` and ` + "`" + `array` + "`" + ` only support type ` + "`" + `array` + "`" + `; 2. Terraform does not support the ` + "`" + `root` + "`" + ` syntax, when it appears, it must be replaced with the uin it stands for.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Name of CAM policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the CAM policy. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CAM policy.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the policy strategy. Valid values: ` + "`" + `1` + "`" + `, ` + "`" + `2` + "`" + `. ` + "`" + `1` + "`" + ` means customer strategy and ` + "`" + `2` + "`" + ` means preset strategy.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `The last update time of the CAM policy. ## Import CAM policy can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cam_policy.foo 26655801 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CAM policy.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the policy strategy. Valid values: ` + "`" + `1` + "`" + `, ` + "`" + `2` + "`" + `. ` + "`" + `1` + "`" + ` means customer strategy and ` + "`" + `2` + "`" + ` means preset strategy.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `The last update time of the CAM policy. ## Import CAM policy can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cam_policy.foo 26655801 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cam_role",
			Category:         "Cloud Access Management(CAM)",
			ShortDescription: `Provides a resource to create a CAM role.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"access",
				"management",
				"cam",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "document",
					Description: `(Required) Document of the CAM role. The syntax refers to [CAM POLICY](https://intl.cloud.tencent.com/document/product/598/10604). There are some notes when using this para in terraform: 1. The elements in json claimed supporting two types as ` + "`" + `string` + "`" + ` and ` + "`" + `array` + "`" + ` only support type ` + "`" + `array` + "`" + `; 2. Terraform does not support the ` + "`" + `root` + "`" + ` syntax, when appears, it must be replaced with the uin it stands for.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Name of CAM role.`,
				},
				resource.Attribute{
					Name:        "console_login",
					Description: `(Optional, ForceNew) Indicade whether the CAM role can login or not.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the CAM role. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CAM role.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `The last update time of the CAM role. ## Import CAM role can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cam_role.foo 4611686018427733635 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CAM role.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `The last update time of the CAM role. ## Import CAM role can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cam_role.foo 4611686018427733635 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cam_role_policy_attachment",
			Category:         "Cloud Access Management(CAM)",
			ShortDescription: `Provides a resource to create a CAM role policy attachment.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"access",
				"management",
				"cam",
				"role",
				"policy",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Required, ForceNew) ID of the policy.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `(Required, ForceNew) ID of the attached CAM role. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_mode",
					Description: `Mode of Creation of the CAM role policy attachment. ` + "`" + `1` + "`" + ` means the CAM policy attachment is created by production, and the others indicate syntax strategy ways.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The create time of the CAM role policy attachment.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `The name of the policy.`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `Type of the policy strategy. ` + "`" + `User` + "`" + ` means customer strategy and ` + "`" + `QCS` + "`" + ` means preset strategy. ## Import CAM role policy attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cam_role_policy_attachment.foo 4611686018427922725#26800353 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_mode",
					Description: `Mode of Creation of the CAM role policy attachment. ` + "`" + `1` + "`" + ` means the CAM policy attachment is created by production, and the others indicate syntax strategy ways.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The create time of the CAM role policy attachment.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `The name of the policy.`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `Type of the policy strategy. ` + "`" + `User` + "`" + ` means customer strategy and ` + "`" + `QCS` + "`" + ` means preset strategy. ## Import CAM role policy attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cam_role_policy_attachment.foo 4611686018427922725#26800353 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cam_saml_provider",
			Category:         "Cloud Access Management(CAM)",
			ShortDescription: `Provides a resource to create a CAM SAML provider.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"access",
				"management",
				"cam",
				"saml",
				"provider",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Required) The description of the CAM SAML provider.`,
				},
				resource.Attribute{
					Name:        "meta_data",
					Description: `(Required) The meta data document of the CAM SAML provider.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Name of CAM SAML provider. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The create time of the CAM SAML provider.`,
				},
				resource.Attribute{
					Name:        "provider_arn",
					Description: `The ARN of the CAM SAML provider.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `The last update time of the CAM SAML provider. ## Import CAM SAML provider can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cam_saml_provider.foo cam-SAML-provider-test ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The create time of the CAM SAML provider.`,
				},
				resource.Attribute{
					Name:        "provider_arn",
					Description: `The ARN of the CAM SAML provider.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `The last update time of the CAM SAML provider. ## Import CAM SAML provider can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cam_saml_provider.foo cam-SAML-provider-test ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cam_user",
			Category:         "Cloud Access Management(CAM)",
			ShortDescription: `Provides a resource to manage CAM user.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"access",
				"management",
				"cam",
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Name of the CAM user.`,
				},
				resource.Attribute{
					Name:        "console_login",
					Description: `(Optional) Indicate whether the CAM user can login to the web console or not.`,
				},
				resource.Attribute{
					Name:        "country_code",
					Description: `(Optional) Country code of the phone number, for example: '86'.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) Email of the CAM user.`,
				},
				resource.Attribute{
					Name:        "force_delete",
					Description: `(Optional) Indicate whether to force deletes the CAM user. If set false, the API secret key will be checked and failed when exists; otherwise the user will be deleted directly. Default is false.`,
				},
				resource.Attribute{
					Name:        "need_reset_password",
					Description: `(Optional) Indicate whether the CAM user need to reset the password when first logins.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The password of the CAM user. Password should be at least 8 characters and no more than 32 characters, includes uppercase letters, lowercase letters, numbers and special characters. Only required when ` + "`" + `console_login` + "`" + ` is true. If not set, a random password will be automatically generated.`,
				},
				resource.Attribute{
					Name:        "phone_num",
					Description: `(Optional) Phone number of the CAM user.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `(Optional) Remark of the CAM user.`,
				},
				resource.Attribute{
					Name:        "use_api",
					Description: `(Optional) Indicate whether to generate the API secret key or not. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "secret_id",
					Description: `Secret ID of the CAM user.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `Secret key of the CAM user.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `ID of the CAM user.`,
				},
				resource.Attribute{
					Name:        "uin",
					Description: `Uin of the CAM User. ## Import CAM user can be imported using the user name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cam_user.foo cam-user-test ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "secret_id",
					Description: `Secret ID of the CAM user.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `Secret key of the CAM user.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `ID of the CAM user.`,
				},
				resource.Attribute{
					Name:        "uin",
					Description: `Uin of the CAM User. ## Import CAM user can be imported using the user name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cam_user.foo cam-user-test ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cam_user_policy_attachment",
			Category:         "Cloud Access Management(CAM)",
			ShortDescription: `Provides a resource to create a CAM user policy attachment.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"access",
				"management",
				"cam",
				"user",
				"policy",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Required, ForceNew) ID of the policy.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `(Required, ForceNew) ID of the attached CAM user. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_mode",
					Description: `Mode of Creation of the CAM user policy attachment. ` + "`" + `1` + "`" + ` means the CAM policy attachment is created by production, and the others indicate syntax strategy ways.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CAM user policy attachment.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `Name of the policy.`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `Type of the policy strategy. ` + "`" + `User` + "`" + ` means customer strategy and ` + "`" + `QCS` + "`" + ` means preset strategy. ## Import CAM user policy attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cam_user_policy_attachment.foo cam-test#26800353 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_mode",
					Description: `Mode of Creation of the CAM user policy attachment. ` + "`" + `1` + "`" + ` means the CAM policy attachment is created by production, and the others indicate syntax strategy ways.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CAM user policy attachment.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `Name of the policy.`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `Type of the policy strategy. ` + "`" + `User` + "`" + ` means customer strategy and ` + "`" + `QCS` + "`" + ` means preset strategy. ## Import CAM user policy attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cam_user_policy_attachment.foo cam-test#26800353 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cbs_snapshot",
			Category:         "Cloud Block Storage(CBS)",
			ShortDescription: `Provides a resource to create a CBS snapshot.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"block",
				"storage",
				"cbs",
				"snapshot",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "snapshot_name",
					Description: `(Required) Name of the snapshot.`,
				},
				resource.Attribute{
					Name:        "storage_id",
					Description: `(Required, ForceNew) ID of the the CBS which this snapshot created from. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of snapshot.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `Types of CBS which this snapshot created from.`,
				},
				resource.Attribute{
					Name:        "percent",
					Description: `Snapshot creation progress percentage. If the snapshot has created successfully, the constant value is 100.`,
				},
				resource.Attribute{
					Name:        "snapshot_status",
					Description: `Status of the snapshot.`,
				},
				resource.Attribute{
					Name:        "storage_size",
					Description: `Volume of storage which this snapshot created from. ## Import CBS snapshot can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cbs_snapshot.snapshot snap-3sa3f39b ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of snapshot.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `Types of CBS which this snapshot created from.`,
				},
				resource.Attribute{
					Name:        "percent",
					Description: `Snapshot creation progress percentage. If the snapshot has created successfully, the constant value is 100.`,
				},
				resource.Attribute{
					Name:        "snapshot_status",
					Description: `Status of the snapshot.`,
				},
				resource.Attribute{
					Name:        "storage_size",
					Description: `Volume of storage which this snapshot created from. ## Import CBS snapshot can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cbs_snapshot.snapshot snap-3sa3f39b ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cbs_snapshot_policy",
			Category:         "Cloud Block Storage(CBS)",
			ShortDescription: `Provides a snapshot policy resource.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"block",
				"storage",
				"cbs",
				"snapshot",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "repeat_hours",
					Description: `(Required) Trigger times of periodic snapshot. Valid value ranges: (0~23). The 0 means 00:00, and so on.`,
				},
				resource.Attribute{
					Name:        "repeat_weekdays",
					Description: `(Required) Periodic snapshot is enabled. Valid values: [0, 1, 2, 3, 4, 5, 6]. 0 means Sunday, 1-6 means Monday to Saturday.`,
				},
				resource.Attribute{
					Name:        "snapshot_policy_name",
					Description: `(Required) Name of snapshot policy. The maximum length can not exceed 60 bytes.`,
				},
				resource.Attribute{
					Name:        "retention_days",
					Description: `(Optional) Retention days of the snapshot, and the default value is 7. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import CBS snapshot policy can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cbs_snapshot_policy.snapshot_policy asp-jliex1tn ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import CBS snapshot policy can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cbs_snapshot_policy.snapshot_policy asp-jliex1tn ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cbs_snapshot_policy_attachment",
			Category:         "Cloud Block Storage(CBS)",
			ShortDescription: `Provides a CBS snapshot policy attachment resource.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"block",
				"storage",
				"cbs",
				"snapshot",
				"policy",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "snapshot_policy_id",
					Description: `(Required, ForceNew) ID of CBS snapshot policy.`,
				},
				resource.Attribute{
					Name:        "storage_id",
					Description: `(Required, ForceNew) ID of CBS. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cbs_storage",
			Category:         "Cloud Block Storage(CBS)",
			ShortDescription: `Provides a resource to create a CBS.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"block",
				"storage",
				"cbs",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required, ForceNew) The available zone that the CBS instance locates at.`,
				},
				resource.Attribute{
					Name:        "storage_name",
					Description: `(Required) Name of CBS. The maximum length can not exceed 60 bytes.`,
				},
				resource.Attribute{
					Name:        "storage_size",
					Description: `(Required) Volume of CBS, and unit is GB. If storage type is ` + "`" + `CLOUD_SSD` + "`" + `, the size range is [100, 16000], and the others are [10-16000].`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `(Required, ForceNew) Type of CBS medium. Valid values: CLOUD_BASIC, CLOUD_PREMIUM, CLOUD_SSD, CLOUD_TSSD and CLOUD_HSSD.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `(Optional) The charge type of CBS instance. Valid values are ` + "`" + `PREPAID` + "`" + ` and ` + "`" + `POSTPAID_BY_HOUR` + "`" + `. The default is ` + "`" + `POSTPAID_BY_HOUR` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "encrypt",
					Description: `(Optional, ForceNew) Indicates whether CBS is encrypted.`,
				},
				resource.Attribute{
					Name:        "force_delete",
					Description: `(Optional) Indicate whether to delete CBS instance directly or not. Default is false. If set true, the instance will be deleted instead of staying recycle bin.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "prepaid_period",
					Description: `(Optional) The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when charge_type is set to ` + "`" + `PREPAID` + "`" + `. Valid values are 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36.`,
				},
				resource.Attribute{
					Name:        "prepaid_renew_flag",
					Description: `(Optional) Auto Renewal flag. Value range: ` + "`" + `NOTIFY_AND_AUTO_RENEW` + "`" + `: Notify expiry and renew automatically, ` + "`" + `NOTIFY_AND_MANUAL_RENEW` + "`" + `: Notify expiry but do not renew automatically, ` + "`" + `DISABLE_NOTIFY_AND_MANUAL_RENEW` + "`" + `: Neither notify expiry nor renew automatically. Default value range: ` + "`" + `NOTIFY_AND_MANUAL_RENEW` + "`" + `: Notify expiry but do not renew automatically. NOTE: it only works when charge_type is set to ` + "`" + `PREPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) ID of the project to which the instance belongs.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) ID of the snapshot. If specified, created the CBS by this snapshot.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The available tags within this CBS.`,
				},
				resource.Attribute{
					Name:        "throughput_performance",
					Description: `(Optional) Add extra performance to the data disk. Only works when disk type is ` + "`" + `CLOUD_TSSD` + "`" + ` or ` + "`" + `CLOUD_HSSD` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "attached",
					Description: `Indicates whether the CBS is mounted the CVM.`,
				},
				resource.Attribute{
					Name:        "storage_status",
					Description: `Status of CBS. Valid values: UNATTACHED, ATTACHING, ATTACHED, DETACHING, EXPANDING, ROLLBACKING, TORECYCLE and DUMPING. ## Import CBS storage can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cbs_storage.storage disk-41s6jwy4 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "attached",
					Description: `Indicates whether the CBS is mounted the CVM.`,
				},
				resource.Attribute{
					Name:        "storage_status",
					Description: `Status of CBS. Valid values: UNATTACHED, ATTACHING, ATTACHED, DETACHING, EXPANDING, ROLLBACKING, TORECYCLE and DUMPING. ## Import CBS storage can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cbs_storage.storage disk-41s6jwy4 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cbs_storage_attachment",
			Category:         "Cloud Block Storage(CBS)",
			ShortDescription: `Provides a CBS storage attachment resource.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"block",
				"storage",
				"cbs",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, ForceNew) ID of the CVM instance.`,
				},
				resource.Attribute{
					Name:        "storage_id",
					Description: `(Required, ForceNew) ID of the mounted CBS. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import CBS storage attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cbs_storage_attachment.attachment disk-41s6jwy4 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import CBS storage attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cbs_storage_attachment.attachment disk-41s6jwy4 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_ccn",
			Category:         "Cloud Connect Network(CCN)",
			ShortDescription: `Provides a resource to create a CCN instance.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"connect",
				"network",
				"ccn",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the CCN to be queried, and maximum length does not exceed 60 bytes.`,
				},
				resource.Attribute{
					Name:        "bandwidth_limit_type",
					Description: `(Optional) The speed limit type. Valid values: ` + "`" + `INTER_REGION_LIMIT` + "`" + `, ` + "`" + `OUTER_REGION_LIMIT` + "`" + `. ` + "`" + `OUTER_REGION_LIMIT` + "`" + ` represents the regional export speed limit, ` + "`" + `INTER_REGION_LIMIT` + "`" + ` is the inter-regional speed limit. The default is ` + "`" + `OUTER_REGION_LIMIT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `(Optional, ForceNew) Billing mode. Valid values: ` + "`" + `PREPAID` + "`" + `, ` + "`" + `POSTPAID` + "`" + `. ` + "`" + `PREPAID` + "`" + ` means prepaid, which means annual and monthly subscription, ` + "`" + `POSTPAID` + "`" + ` means post-payment, which means billing by volume. The default is ` + "`" + `POSTPAID` + "`" + `. The prepaid model only supports inter-regional speed limit, and the post-paid model supports inter-regional speed limit and regional export speed limit.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of CCN, and maximum length does not exceed 100 bytes.`,
				},
				resource.Attribute{
					Name:        "qos",
					Description: `(Optional, ForceNew) Service quality of CCN. Valid values: ` + "`" + `PT` + "`" + `, ` + "`" + `AU` + "`" + `, ` + "`" + `AG` + "`" + `. The default is ` + "`" + `AU` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Instance tag. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of resource.`,
				},
				resource.Attribute{
					Name:        "instance_count",
					Description: `Number of attached instances.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `States of instance. Valid values: ` + "`" + `ISOLATED` + "`" + `(arrears) and ` + "`" + `AVAILABLE` + "`" + `. ## Import Ccn instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_ccn.test ccn-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of resource.`,
				},
				resource.Attribute{
					Name:        "instance_count",
					Description: `Number of attached instances.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `States of instance. Valid values: ` + "`" + `ISOLATED` + "`" + `(arrears) and ` + "`" + `AVAILABLE` + "`" + `. ## Import Ccn instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_ccn.test ccn-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_ccn_attachment",
			Category:         "Cloud Connect Network(CCN)",
			ShortDescription: `Provides a CCN attaching resource.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"connect",
				"network",
				"ccn",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ccn_id",
					Description: `(Required, ForceNew) ID of the CCN.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, ForceNew) ID of instance is attached.`,
				},
				resource.Attribute{
					Name:        "instance_region",
					Description: `(Required, ForceNew) The region that the instance locates at.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Required, ForceNew) Type of attached instance network, and available values include ` + "`" + `VPC` + "`" + `, ` + "`" + `DIRECTCONNECT` + "`" + `, ` + "`" + `BMVPC` + "`" + ` and ` + "`" + `VPNGW` + "`" + `. Note: ` + "`" + `VPNGW` + "`" + ` type is only for whitelist customer now.`,
				},
				resource.Attribute{
					Name:        "ccn_uin",
					Description: `(Optional, ForceNew) Uin of the ccn attached. Default is ` + "`" + `` + "`" + `, which means the uin of this account. This parameter is used with case when attaching ccn of other account to the instance of this account. For now only support instance type ` + "`" + `VPC` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "attached_time",
					Description: `Time of attaching.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `A network address block of the instance that is attached.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `States of instance is attached. Valid values: ` + "`" + `PENDING` + "`" + `, ` + "`" + `ACTIVE` + "`" + `, ` + "`" + `EXPIRED` + "`" + `, ` + "`" + `REJECTED` + "`" + `, ` + "`" + `DELETED` + "`" + `, ` + "`" + `FAILED` + "`" + `, ` + "`" + `ATTACHING` + "`" + `, ` + "`" + `DETACHING` + "`" + ` and ` + "`" + `DETACHFAILED` + "`" + `. ` + "`" + `FAILED` + "`" + ` means asynchronous forced disassociation after 2 hours. ` + "`" + `DETACHFAILED` + "`" + ` means asynchronous forced disassociation after 2 hours.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "attached_time",
					Description: `Time of attaching.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `A network address block of the instance that is attached.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `States of instance is attached. Valid values: ` + "`" + `PENDING` + "`" + `, ` + "`" + `ACTIVE` + "`" + `, ` + "`" + `EXPIRED` + "`" + `, ` + "`" + `REJECTED` + "`" + `, ` + "`" + `DELETED` + "`" + `, ` + "`" + `FAILED` + "`" + `, ` + "`" + `ATTACHING` + "`" + `, ` + "`" + `DETACHING` + "`" + ` and ` + "`" + `DETACHFAILED` + "`" + `. ` + "`" + `FAILED` + "`" + ` means asynchronous forced disassociation after 2 hours. ` + "`" + `DETACHFAILED` + "`" + ` means asynchronous forced disassociation after 2 hours.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_ccn_bandwidth_limit",
			Category:         "Cloud Connect Network(CCN)",
			ShortDescription: `Provides a resource to limit CCN bandwidth.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"connect",
				"network",
				"ccn",
				"bandwidth",
				"limit",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ccn_id",
					Description: `(Required, ForceNew) ID of the CCN.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required, ForceNew) Limitation of region.`,
				},
				resource.Attribute{
					Name:        "bandwidth_limit",
					Description: `(Optional) Limitation of bandwidth.`,
				},
				resource.Attribute{
					Name:        "dst_region",
					Description: `(Optional, ForceNew) Destination area restriction. If the ` + "`" + `CCN` + "`" + ` rate limit type is ` + "`" + `OUTER_REGION_LIMIT` + "`" + `, this value does not need to be set. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cdn_domain",
			Category:         "Content Delivery Network(CDN)",
			ShortDescription: `Provides a resource to create a CDN domain.`,
			Description:      ``,
			Keywords: []string{
				"content",
				"delivery",
				"network",
				"cdn",
				"domain",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `(Required, ForceNew) Name of the acceleration domain.`,
				},
				resource.Attribute{
					Name:        "origin",
					Description: `(Required) Origin server configuration. It's a list and consist of at most one item.`,
				},
				resource.Attribute{
					Name:        "service_type",
					Description: `(Required, ForceNew) Acceleration domain name service type. ` + "`" + `web` + "`" + `: static acceleration, ` + "`" + `download` + "`" + `: download acceleration, ` + "`" + `media` + "`" + `: streaming media VOD acceleration.`,
				},
				resource.Attribute{
					Name:        "area",
					Description: `(Optional) Domain name acceleration region. ` + "`" + `mainland` + "`" + `: acceleration inside mainland China, ` + "`" + `overseas` + "`" + `: acceleration outside mainland China, ` + "`" + `global` + "`" + `: global acceleration. Overseas acceleration service must be enabled to use overseas acceleration and global acceleration.`,
				},
				resource.Attribute{
					Name:        "full_url_cache",
					Description: `(Optional) Whether to enable full-path cache. Default value is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "https_config",
					Description: `(Optional) HTTPS acceleration configuration. It's a list and consist of at most one item.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) The project CDN belongs to, default to 0.`,
				},
				resource.Attribute{
					Name:        "range_origin_switch",
					Description: `(Optional) Sharding back to source configuration switch. Valid values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. Default value is ` + "`" + `on` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "request_header",
					Description: `(Optional) Request header configuration. It's a list and consist of at most one item.`,
				},
				resource.Attribute{
					Name:        "rule_cache",
					Description: `(Optional) Advanced path cache configuration.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of cdn domain. The ` + "`" + `client_certificate_config` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "certificate_content",
					Description: `(Required) Client Certificate PEM format, requires Base64 encoding. The ` + "`" + `force_redirect` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "redirect_status_code",
					Description: `(Optional) Forced redirect status code. Valid values are ` + "`" + `301` + "`" + ` and ` + "`" + `302` + "`" + `. When ` + "`" + `switch` + "`" + ` setting ` + "`" + `off` + "`" + `, this property does not need to be set or set to ` + "`" + `302` + "`" + `. Default value is ` + "`" + `302` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "redirect_type",
					Description: `(Optional) Forced redirect type. Valid values are ` + "`" + `http` + "`" + ` and ` + "`" + `https` + "`" + `. ` + "`" + `http` + "`" + ` means a forced redirect from HTTPS to HTTP, ` + "`" + `https` + "`" + ` means a forced redirect from HTTP to HTTPS. When ` + "`" + `switch` + "`" + ` setting ` + "`" + `off` + "`" + `, this property does not need to be set or set to ` + "`" + `http` + "`" + `. Default value is ` + "`" + `http` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "switch",
					Description: `(Optional) Forced redirect configuration switch. Valid values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. Default value is ` + "`" + `off` + "`" + `. The ` + "`" + `header_rules` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "header_mode",
					Description: `(Required) Http header setting method. The following types are supported: ` + "`" + `add` + "`" + `: add a head, if a head already exists, there will be a duplicate head, ` + "`" + `del` + "`" + `: delete the head.`,
				},
				resource.Attribute{
					Name:        "header_name",
					Description: `(Required) Http header name.`,
				},
				resource.Attribute{
					Name:        "header_value",
					Description: `(Required) Http header value, optional when Mode is ` + "`" + `del` + "`" + `, Required when Mode is ` + "`" + `add` + "`" + `/` + "`" + `set` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rule_paths",
					Description: `(Required) Matching content under the corresponding type of CacheType: ` + "`" + `all` + "`" + `: fill`,
				},
				resource.Attribute{
					Name:        "rule_type",
					Description: `(Required) Rule type. The following types are supported: ` + "`" + `all` + "`" + `: all documents take effect, ` + "`" + `file` + "`" + `: the specified file suffix takes effect, ` + "`" + `directory` + "`" + `: the specified path takes effect, ` + "`" + `path` + "`" + `: specify the absolute path to take effect. The ` + "`" + `https_config` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "https_switch",
					Description: `(Required) HTTPS configuration switch. Valid values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "client_certificate_config",
					Description: `(Optional) Client certificate configuration information.`,
				},
				resource.Attribute{
					Name:        "force_redirect",
					Description: `(Optional) Configuration of forced HTTP or HTTPS redirects.`,
				},
				resource.Attribute{
					Name:        "http2_switch",
					Description: `(Optional) HTTP2 configuration switch. Valid values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. and default value is ` + "`" + `off` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ocsp_stapling_switch",
					Description: `(Optional) OCSP configuration switch. Valid values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. and default value is ` + "`" + `off` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "server_certificate_config",
					Description: `(Optional) Server certificate configuration information.`,
				},
				resource.Attribute{
					Name:        "spdy_switch",
					Description: `(Optional) Spdy configuration switch. Valid values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. and default value is ` + "`" + `off` + "`" + `. This parameter is for white-list customer.`,
				},
				resource.Attribute{
					Name:        "verify_client",
					Description: `(Optional) Client certificate authentication feature. Valid values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. and default value is ` + "`" + `off` + "`" + `. The ` + "`" + `origin` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "origin_list",
					Description: `(Required) Master origin server list. Valid values can be ip or domain name. When modifying the origin server, you need to enter the corresponding ` + "`" + `origin_type` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "origin_type",
					Description: `(Required) Master origin server type. The following types are supported: ` + "`" + `domain` + "`" + `: domain name type, ` + "`" + `cos` + "`" + `: COS origin, ` + "`" + `ip` + "`" + `: IP list used as origin server, ` + "`" + `ipv6` + "`" + `: origin server list is a single IPv6 address, ` + "`" + `ip_ipv6` + "`" + `: origin server list is multiple IPv4 addresses and an IPv6 address.`,
				},
				resource.Attribute{
					Name:        "backup_origin_list",
					Description: `(Optional) Backup origin server list. Valid values can be ip or domain name. When modifying the backup origin server, you need to enter the corresponding ` + "`" + `backup_origin_type` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "backup_origin_type",
					Description: `(Optional) Backup origin server type, which supports the following types: ` + "`" + `domain` + "`" + `: domain name type, ` + "`" + `ip` + "`" + `: IP list used as origin server.`,
				},
				resource.Attribute{
					Name:        "backup_server_name",
					Description: `(Optional) Host header used when accessing the backup origin server. If left empty, the ServerName of master origin server will be used by default.`,
				},
				resource.Attribute{
					Name:        "cos_private_access",
					Description: `(Optional) When OriginType is COS, you can specify if access to private buckets is allowed. Valid values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. and default value is ` + "`" + `off` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "origin_pull_protocol",
					Description: `(Optional) Origin-pull protocol configuration. ` + "`" + `http` + "`" + `: forced HTTP origin-pull, ` + "`" + `follow` + "`" + `: protocol follow origin-pull, ` + "`" + `https` + "`" + `: forced HTTPS origin-pull. This only supports origin server port 443 for origin-pull.`,
				},
				resource.Attribute{
					Name:        "server_name",
					Description: `(Optional) Host header used when accessing the master origin server. If left empty, the acceleration domain name will be used by default. The ` + "`" + `request_header` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "header_rules",
					Description: `(Optional) Custom request header configuration rules.`,
				},
				resource.Attribute{
					Name:        "switch",
					Description: `(Optional) Custom request header configuration switch. Valid values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. and default value is ` + "`" + `off` + "`" + `. The ` + "`" + `rule_cache` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "cache_time",
					Description: `(Required) Cache expiration time setting, the unit is second, the maximum can be set to 365 days.`,
				},
				resource.Attribute{
					Name:        "compare_max_age",
					Description: `(Optional) Advanced cache expiration configuration. When it is turned on, it will compare the max-age value returned by the origin site with the cache expiration time set in CacheRules, and take the minimum value to cache at the node. Valid values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. Default value is ` + "`" + `off` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "follow_origin_switch",
					Description: `(Optional) Follow the source station configuration switch. Valid values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ignore_cache_control",
					Description: `(Optional) Force caching. After opening, the no-store and no-cache resources returned by the origin site will also be cached in accordance with the CacheRules rules. Valid values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. Default value is ` + "`" + `off` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ignore_set_cookie",
					Description: `(Optional) Ignore the Set-Cookie header of the origin site. Valid values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. Default value is ` + "`" + `off` + "`" + `. This parameter is for white-list customer.`,
				},
				resource.Attribute{
					Name:        "no_cache_switch",
					Description: `(Optional) Cache configuration switch. Valid values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "re_validate",
					Description: `(Optional) Always check back to origin. Valid values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. Default value is ` + "`" + `off` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rule_paths",
					Description: `(Optional) Matching content under the corresponding type of CacheType: ` + "`" + `all` + "`" + `: fill`,
				},
				resource.Attribute{
					Name:        "rule_type",
					Description: `(Optional) Rule type. The following types are supported: ` + "`" + `all` + "`" + `: all documents take effect, ` + "`" + `file` + "`" + `: the specified file suffix takes effect, ` + "`" + `directory` + "`" + `: the specified path takes effect, ` + "`" + `path` + "`" + `: specify the absolute path to take effect, ` + "`" + `index` + "`" + `: home page, ` + "`" + `default` + "`" + `: effective when the source site has no max-age.`,
				},
				resource.Attribute{
					Name:        "switch",
					Description: `(Optional) Cache configuration switch. Valid values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. The ` + "`" + `server_certificate_config` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "certificate_content",
					Description: `(Optional) Server certificate information. This is required when uploading an external certificate, which should contain the complete certificate chain.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `(Optional) Server certificate ID.`,
				},
				resource.Attribute{
					Name:        "message",
					Description: `(Optional) Certificate remarks.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Optional) Server key information. This is required when uploading an external certificate. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "cname",
					Description: `CNAME address of domain name.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of domain name.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Acceleration service status. ## Import CDN domain can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cdn_domain.foo xxxx.com ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "cname",
					Description: `CNAME address of domain name.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of domain name.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Acceleration service status. ## Import CDN domain can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cdn_domain.foo xxxx.com ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cfs_access_group",
			Category:         "Cloud File Storage(CFS)",
			ShortDescription: `Provides a resource to create a CFS access group.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"file",
				"storage",
				"cfs",
				"access",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the access group, and max length is 64.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the access group, and max length is 255. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the access group. ## Import CFS access group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cfs_access_group.foo pgroup-7nx89k7l ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the access group. ## Import CFS access group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cfs_access_group.foo pgroup-7nx89k7l ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cfs_access_rule",
			Category:         "Cloud File Storage(CFS)",
			ShortDescription: `Provides a resource to create a CFS access rule.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"file",
				"storage",
				"cfs",
				"access",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "access_group_id",
					Description: `(Required, ForceNew) ID of a access group.`,
				},
				resource.Attribute{
					Name:        "auth_client_ip",
					Description: `(Required) A single IP or a single IP address range such as 10.1.10.11 or 10.10.1.0/24 indicates that all IPs are allowed. Please note that the IP entered should be CVM's private IP.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Required) The priority level of rule. Valid value ranges: (1~100). ` + "`" + `1` + "`" + ` indicates the highest priority.`,
				},
				resource.Attribute{
					Name:        "rw_permission",
					Description: `(Optional) Read and write permissions. Valid values are ` + "`" + `RO` + "`" + ` and ` + "`" + `RW` + "`" + `. and default is ` + "`" + `RO` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user_permission",
					Description: `(Optional) The permissions of accessing users. Valid values are ` + "`" + `all_squash` + "`" + `, ` + "`" + `no_all_squash` + "`" + `, ` + "`" + `root_squash` + "`" + ` and ` + "`" + `no_root_squash` + "`" + `. and default is ` + "`" + `root_squash` + "`" + `. ` + "`" + `all_squash` + "`" + ` indicates that all access users are mapped as anonymous users or user groups; ` + "`" + `no_all_squash` + "`" + ` indicates that access users will match local users first and be mapped to anonymous users or user groups after matching failed; ` + "`" + `root_squash` + "`" + ` indicates that map access root users to anonymous users or user groups; ` + "`" + `no_root_squash` + "`" + ` indicates that access root users keep root account permission. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cfs_file_system",
			Category:         "Cloud File Storage(CFS)",
			ShortDescription: `Provides a resource to create a cloud file system(CFS).`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"file",
				"storage",
				"cfs",
				"system",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "access_group_id",
					Description: `(Required) ID of a access group.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required, ForceNew) The available zone that the file system locates at.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required, ForceNew) ID of a subnet.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, ForceNew) ID of a VPC network.`,
				},
				resource.Attribute{
					Name:        "mount_ip",
					Description: `(Optional, ForceNew) IP of mount point.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of a file system.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional, ForceNew) File service protocol. Valid values are ` + "`" + `NFS` + "`" + ` and ` + "`" + `CIFS` + "`" + `. and the default is ` + "`" + `NFS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `(Optional, ForceNew) File service StorageType. Valid values are ` + "`" + `SD` + "`" + ` and ` + "`" + `HP` + "`" + `. and the default is ` + "`" + `SD` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Instance tags. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the file system. ## Import Cloud file system can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cfs_file_system.foo cfs-6hgquxmj ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the file system. ## Import Cloud file system can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cfs_file_system.foo cfs-6hgquxmj ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_ckafka_acl",
			Category:         "Ckafka",
			ShortDescription: `Provides a resource to create a Ckafka Acl.`,
			Description:      ``,
			Keywords: []string{
				"ckafka",
				"acl",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, ForceNew) ID of the ckafka instance.`,
				},
				resource.Attribute{
					Name:        "operation_type",
					Description: `(Required, ForceNew) ACL operation mode. Valid values: ` + "`" + `UNKNOWN` + "`" + `, ` + "`" + `ANY` + "`" + `, ` + "`" + `ALL` + "`" + `, ` + "`" + `READ` + "`" + `, ` + "`" + `WRITE` + "`" + `, ` + "`" + `CREATE` + "`" + `, ` + "`" + `DELETE` + "`" + `, ` + "`" + `ALTER` + "`" + `, ` + "`" + `DESCRIBE` + "`" + `, ` + "`" + `CLUSTER_ACTION` + "`" + `, ` + "`" + `DESCRIBE_CONFIGS` + "`" + ` and ` + "`" + `ALTER_CONFIGS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "resource_name",
					Description: `(Required, ForceNew) ACL resource name, which is related to ` + "`" + `resource_type` + "`" + `. For example, if ` + "`" + `resource_type` + "`" + ` is ` + "`" + `TOPIC` + "`" + `, this field indicates the topic name; if ` + "`" + `resource_type` + "`" + ` is ` + "`" + `GROUP` + "`" + `, this field indicates the group name.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Optional, ForceNew) IP address allowed to access. The default value is ` + "`" + ``,
				},
				resource.Attribute{
					Name:        "permission_type",
					Description: `(Optional, ForceNew) ACL permission type. Valid values: ` + "`" + `UNKNOWN` + "`" + `, ` + "`" + `ANY` + "`" + `, ` + "`" + `DENY` + "`" + `, ` + "`" + `ALLOW` + "`" + `. and ` + "`" + `ALLOW` + "`" + ` by default. Currently, CKafka supports ` + "`" + `ALLOW` + "`" + ` (equivalent to allow list), and other fields will be used for future ACLs compatible with open-source Kafka.`,
				},
				resource.Attribute{
					Name:        "principal",
					Description: `(Optional, ForceNew) User list. The default value is ` + "`" + ``,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `(Optional, ForceNew) ACL resource type. Valid values are ` + "`" + `UNKNOWN` + "`" + `, ` + "`" + `ANY` + "`" + `, ` + "`" + `TOPIC` + "`" + `, ` + "`" + `GROUP` + "`" + `, ` + "`" + `CLUSTER` + "`" + `, ` + "`" + `TRANSACTIONAL_ID` + "`" + `. and ` + "`" + `TOPIC` + "`" + ` by default. Currently, only ` + "`" + `TOPIC` + "`" + ` is available, and other fields will be used for future ACLs compatible with open-source Kafka. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import Ckafka acl can be imported using the instance_id#permission_type#principal#host#operation_type#resource_type#resource_name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_ckafka_acl.foo ckafka-f9ife4zz#ALLOW#test#`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import Ckafka acl can be imported using the instance_id#permission_type#principal#host#operation_type#resource_type#resource_name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_ckafka_acl.foo ckafka-f9ife4zz#ALLOW#test#`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_ckafka_topic",
			Category:         "Ckafka",
			ShortDescription: `Use this resource to create ckafka topic.`,
			Description:      ``,
			Keywords: []string{
				"ckafka",
				"topic",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, ForceNew) Ckafka instance ID.`,
				},
				resource.Attribute{
					Name:        "partition_num",
					Description: `(Required) The number of partition.`,
				},
				resource.Attribute{
					Name:        "replica_num",
					Description: `(Required) The number of replica, the maximum is 3.`,
				},
				resource.Attribute{
					Name:        "topic_name",
					Description: `(Required, ForceNew) Name of the CKafka topic. It must start with a letter, the rest can contain letters, numbers and dashes(-). The length range is from 1 to 64.`,
				},
				resource.Attribute{
					Name:        "clean_up_policy",
					Description: `(Optional) Clear log policy, log clear mode, default is ` + "`" + `delete` + "`" + `. ` + "`" + `delete` + "`" + `: logs are deleted according to the storage time. ` + "`" + `compact` + "`" + `: logs are compressed according to the key. ` + "`" + `compact, delete` + "`" + `: logs are compressed according to the key and will be deleted according to the storage time.`,
				},
				resource.Attribute{
					Name:        "enable_white_list",
					Description: `(Optional) Whether to open the ip whitelist, ` + "`" + `true` + "`" + `: open, ` + "`" + `false` + "`" + `: close.`,
				},
				resource.Attribute{
					Name:        "ip_white_list",
					Description: `(Optional) Ip whitelist, quota limit, required when enableWhileList=true.`,
				},
				resource.Attribute{
					Name:        "max_message_bytes",
					Description: `(Optional) Max message bytes.`,
				},
				resource.Attribute{
					Name:        "note",
					Description: `(Optional) The subject note is a string of no more than 64 characters. It must start with a letter, and the remaining part can contain letters, numbers and dashes (-).`,
				},
				resource.Attribute{
					Name:        "retention",
					Description: `(Optional) Message can be selected. Retention time, unit is ms, the current minimum value is 60000ms.`,
				},
				resource.Attribute{
					Name:        "segment",
					Description: `(Optional) Segment scrolling time, in ms, the current minimum is 3600000ms.`,
				},
				resource.Attribute{
					Name:        "sync_replica_min_num",
					Description: `(Optional) Min number of sync replicas, Default is ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "unclean_leader_election_enable",
					Description: `(Optional) Whether to allow unsynchronized replicas to be selected as leader, default is ` + "`" + `false` + "`" + `, ` + "`" + `true: ` + "`" + `allowed, ` + "`" + `false` + "`" + `: not allowed. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CKafka topic.`,
				},
				resource.Attribute{
					Name:        "forward_cos_bucket",
					Description: `Data backup cos bucket: the bucket address that is dumped to cos.`,
				},
				resource.Attribute{
					Name:        "forward_interval",
					Description: `Periodic frequency of data backup to cos.`,
				},
				resource.Attribute{
					Name:        "forward_status",
					Description: `Data backup cos status. Valid values: ` + "`" + `0` + "`" + `, ` + "`" + `1` + "`" + `. ` + "`" + `1` + "`" + `: do not open data backup, ` + "`" + `0` + "`" + `: open data backup.`,
				},
				resource.Attribute{
					Name:        "message_storage_location",
					Description: `Message storage location.`,
				},
				resource.Attribute{
					Name:        "segment_bytes",
					Description: `Number of bytes rolled by shard. ## Import ckafka topic can be imported using the instance_id#topic_name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_ckafka_topic.foo ckafka-f9ife4zz#example ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CKafka topic.`,
				},
				resource.Attribute{
					Name:        "forward_cos_bucket",
					Description: `Data backup cos bucket: the bucket address that is dumped to cos.`,
				},
				resource.Attribute{
					Name:        "forward_interval",
					Description: `Periodic frequency of data backup to cos.`,
				},
				resource.Attribute{
					Name:        "forward_status",
					Description: `Data backup cos status. Valid values: ` + "`" + `0` + "`" + `, ` + "`" + `1` + "`" + `. ` + "`" + `1` + "`" + `: do not open data backup, ` + "`" + `0` + "`" + `: open data backup.`,
				},
				resource.Attribute{
					Name:        "message_storage_location",
					Description: `Message storage location.`,
				},
				resource.Attribute{
					Name:        "segment_bytes",
					Description: `Number of bytes rolled by shard. ## Import ckafka topic can be imported using the instance_id#topic_name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_ckafka_topic.foo ckafka-f9ife4zz#example ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_ckafka_user",
			Category:         "Ckafka",
			ShortDescription: `Provides a resource to create a Ckafka user.`,
			Description:      ``,
			Keywords: []string{
				"ckafka",
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_name",
					Description: `(Required, ForceNew) Account name used to access to ckafka instance.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, ForceNew) ID of the ckafka instance.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Password of the account. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the account.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `The last update time of the account. ## Import Ckafka user can be imported using the instance_id#account_name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_ckafka_user.foo ckafka-f9ife4zz#tf-test ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the account.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `The last update time of the account. ## Import Ckafka user can be imported using the instance_id#account_name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_ckafka_user.foo ckafka-f9ife4zz#tf-test ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_clb_attachment",
			Category:         "Cloud Load Balancer(CLB)",
			ShortDescription: `Provides a resource to create a CLB attachment.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"load",
				"balancer",
				"clb",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "clb_id",
					Description: `(Required, ForceNew) ID of the CLB.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Required, ForceNew) ID of the CLB listener.`,
				},
				resource.Attribute{
					Name:        "targets",
					Description: `(Required) Information of the backends to be attached.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `(Optional, ForceNew) ID of the CLB listener rule. Only supports listeners of ` + "`" + `HTTPS` + "`" + ` and ` + "`" + `HTTP` + "`" + ` protocol. The ` + "`" + `targets` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) Id of the backend server.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) Port of the backend server. Valid value ranges: (0~65535).`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) Forwarding weight of the backend service. Valid value ranges: (0~100). defaults to ` + "`" + `10` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "protocol_type",
					Description: `Type of protocol within the listener. ## Import CLB attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_clb_attachment.foo loc-4xxr2cy7#lbl-hh141sn9#lb-7a0t6zqb ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "protocol_type",
					Description: `Type of protocol within the listener. ## Import CLB attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_clb_attachment.foo loc-4xxr2cy7#lbl-hh141sn9#lb-7a0t6zqb ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_clb_instance",
			Category:         "Cloud Load Balancer(CLB)",
			ShortDescription: `Provides a resource to create a CLB instance.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"load",
				"balancer",
				"clb",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "clb_name",
					Description: `(Required) Name of the CLB. The name can only contain Chinese characters, English letters, numbers, underscore and hyphen '-'.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `(Required, ForceNew) Type of CLB instance. Valid values: ` + "`" + `OPEN` + "`" + ` and ` + "`" + `INTERNAL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "address_ip_version",
					Description: `(Optional) IP version, only applicable to open CLB. Valid values are ` + "`" + `ipv4` + "`" + `, ` + "`" + `ipv6` + "`" + ` and ` + "`" + `IPv6FullChain` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "internet_bandwidth_max_out",
					Description: `(Optional) Max bandwidth out, only applicable to open CLB. Valid value ranges is [1, 2048]. Unit is MB.`,
				},
				resource.Attribute{
					Name:        "internet_charge_type",
					Description: `(Optional) Internet charge type, only applicable to open CLB. Valid values are ` + "`" + `TRAFFIC_POSTPAID_BY_HOUR` + "`" + `, ` + "`" + `BANDWIDTH_POSTPAID_BY_HOUR` + "`" + ` and ` + "`" + `BANDWIDTH_PACKAGE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional, ForceNew) ID of the project within the CLB instance, ` + "`" + `0` + "`" + ` - Default Project.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Optional) Security groups of the CLB instance. Only supports ` + "`" + `OPEN` + "`" + ` CLBs.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional, ForceNew) Subnet ID of the CLB. Effective only for CLB within the VPC. Only supports ` + "`" + `INTERNAL` + "`" + ` CLBs. Default is ` + "`" + `ipv4` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The available tags within this CLB.`,
				},
				resource.Attribute{
					Name:        "target_region_info_region",
					Description: `(Optional) Region information of backend services are attached the CLB instance. Only supports ` + "`" + `OPEN` + "`" + ` CLBs.`,
				},
				resource.Attribute{
					Name:        "target_region_info_vpc_id",
					Description: `(Optional) Vpc information of backend services are attached the CLB instance. Only supports ` + "`" + `OPEN` + "`" + ` CLBs.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional, ForceNew) VPC ID of the CLB. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "clb_vips",
					Description: `The virtual service address table of the CLB.`,
				},
				resource.Attribute{
					Name:        "vip_isp",
					Description: `Network operator, only applicable to open CLB. Valid values are ` + "`" + `CMCC` + "`" + `(China Mobile), ` + "`" + `CTCC` + "`" + `(Telecom), ` + "`" + `CUCC` + "`" + `(China Unicom) and ` + "`" + `BGP` + "`" + `. If this ISP is specified, network billing method can only use the bandwidth package billing (BANDWIDTH_PACKAGE). ## Import CLB instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_clb_instance.foo lb-7a0t6zqb ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "clb_vips",
					Description: `The virtual service address table of the CLB.`,
				},
				resource.Attribute{
					Name:        "vip_isp",
					Description: `Network operator, only applicable to open CLB. Valid values are ` + "`" + `CMCC` + "`" + `(China Mobile), ` + "`" + `CTCC` + "`" + `(Telecom), ` + "`" + `CUCC` + "`" + `(China Unicom) and ` + "`" + `BGP` + "`" + `. If this ISP is specified, network billing method can only use the bandwidth package billing (BANDWIDTH_PACKAGE). ## Import CLB instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_clb_instance.foo lb-7a0t6zqb ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_clb_listener",
			Category:         "Cloud Load Balancer(CLB)",
			ShortDescription: `Provides a resource to create a CLB listener.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"load",
				"balancer",
				"clb",
				"listener",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "clb_id",
					Description: `(Required, ForceNew) ID of the CLB.`,
				},
				resource.Attribute{
					Name:        "listener_name",
					Description: `(Required) Name of the CLB listener, and available values can only be Chinese characters, English letters, numbers, underscore and hyphen '-'.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required, ForceNew) Type of protocol within the listener. Valid values: ` + "`" + `TCP` + "`" + `, ` + "`" + `UDP` + "`" + `, ` + "`" + `HTTP` + "`" + `, ` + "`" + `HTTPS` + "`" + ` and ` + "`" + `TCP_SSL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "certificate_ca_id",
					Description: `(Optional) ID of the client certificate. NOTES: Only supports listeners of ` + "`" + `HTTPS` + "`" + ` and ` + "`" + `TCP_SSL` + "`" + ` protocol and must be set when the ssl mode is ` + "`" + `MUTUAL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `(Optional) ID of the server certificate. NOTES: Only supports listeners of ` + "`" + `HTTPS` + "`" + ` and ` + "`" + `TCP_SSL` + "`" + ` protocol and must be set when it is available.`,
				},
				resource.Attribute{
					Name:        "certificate_ssl_mode",
					Description: `(Optional) Type of certificate. Valid values: ` + "`" + `UNIDIRECTIONAL` + "`" + `, ` + "`" + `MUTUAL` + "`" + `. NOTES: Only supports listeners of ` + "`" + `HTTPS` + "`" + ` and ` + "`" + `TCP_SSL` + "`" + ` protocol and must be set when it is available.`,
				},
				resource.Attribute{
					Name:        "health_check_context_type",
					Description: `(Optional) Health check protocol. When the value of ` + "`" + `health_check_type` + "`" + ` of the health check protocol is ` + "`" + `CUSTOM` + "`" + `, this field is required, which represents the input format of the health check. Valid values: ` + "`" + `HEX` + "`" + `, ` + "`" + `TEXT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_check_health_num",
					Description: `(Optional) Health threshold of health check, and the default is ` + "`" + `3` + "`" + `. If a success result is returned for the health check for 3 consecutive times, the backend CVM is identified as healthy. The value range is 2-10. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "health_check_http_code",
					Description: `(Optional) HTTP health check code of TCP listener. When the value of ` + "`" + `health_check_type` + "`" + ` of the health check protocol is ` + "`" + `HTTP` + "`" + `, this field is required. Valid values: ` + "`" + `1` + "`" + `, ` + "`" + `2` + "`" + `, ` + "`" + `4` + "`" + `, ` + "`" + `8` + "`" + `, ` + "`" + `16` + "`" + `. ` + "`" + `1` + "`" + ` means http_1xx, ` + "`" + `2` + "`" + ` means http_2xx, ` + "`" + `4` + "`" + ` means http_3xx, ` + "`" + `8` + "`" + ` means http_4xx, ` + "`" + `16` + "`" + ` means http_5xx.`,
				},
				resource.Attribute{
					Name:        "health_check_http_domain",
					Description: `(Optional) HTTP health check domain of TCP listener.`,
				},
				resource.Attribute{
					Name:        "health_check_http_method",
					Description: `(Optional) HTTP health check method of TCP listener. Valid values: ` + "`" + `HEAD` + "`" + `, ` + "`" + `GET` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_check_http_path",
					Description: `(Optional) HTTP health check path of TCP listener.`,
				},
				resource.Attribute{
					Name:        "health_check_http_version",
					Description: `(Optional) The HTTP version of the backend service. When the value of ` + "`" + `health_check_type` + "`" + ` of the health check protocol is ` + "`" + `HTTP` + "`" + `, this field is required. Valid values: ` + "`" + `HTTP/1.0` + "`" + `, ` + "`" + `HTTP/1.1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_check_interval_time",
					Description: `(Optional) Interval time of health check. Valid value ranges: [5~300] sec. and the default is 5 sec. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in ` + "`" + `tencentcloud_clb_listener_rule` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_check_port",
					Description: `(Optional) The health check port is the port of the backend service by default. Unless you want to specify a specific port, it is recommended to leave it blank. Only applicable to TCP/UDP listener.`,
				},
				resource.Attribute{
					Name:        "health_check_recv_context",
					Description: `(Optional) It represents the result returned by the health check. When the value of ` + "`" + `health_check_type` + "`" + ` of the health check protocol is ` + "`" + `CUSTOM` + "`" + `, this field is required. Only ASCII visible characters are allowed and the maximum length is 500. When ` + "`" + `health_check_context_type` + "`" + ` value is ` + "`" + `HEX` + "`" + `, the characters of SendContext and RecvContext can only be selected in ` + "`" + `0123456789ABCDEF` + "`" + ` and the length must be even digits.`,
				},
				resource.Attribute{
					Name:        "health_check_send_context",
					Description: `(Optional) It represents the content of the request sent by the health check. When the value of ` + "`" + `health_check_type` + "`" + ` of the health check protocol is ` + "`" + `CUSTOM` + "`" + `, this field is required. Only visible ASCII characters are allowed and the maximum length is 500. When ` + "`" + `health_check_context_type` + "`" + ` value is ` + "`" + `HEX` + "`" + `, the characters of SendContext and RecvContext can only be selected in ` + "`" + `0123456789ABCDEF` + "`" + ` and the length must be even digits.`,
				},
				resource.Attribute{
					Name:        "health_check_switch",
					Description: `(Optional) Indicates whether health check is enabled.`,
				},
				resource.Attribute{
					Name:        "health_check_time_out",
					Description: `(Optional) Response timeout of health check. Valid value ranges: [2~60] sec. Default is 2 sec. Response timeout needs to be less than check interval. NOTES: Only supports listeners of ` + "`" + `TCP` + "`" + `,` + "`" + `UDP` + "`" + `,` + "`" + `TCP_SSL` + "`" + ` protocol.`,
				},
				resource.Attribute{
					Name:        "health_check_type",
					Description: `(Optional) Protocol used for health check. Valid values: ` + "`" + `CUSTOM` + "`" + `, ` + "`" + `TCP` + "`" + `, ` + "`" + `HTTP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_check_unhealth_num",
					Description: `(Optional) Unhealthy threshold of health check, and the default is ` + "`" + `3` + "`" + `. If a success result is returned for the health check 3 consecutive times, the CVM is identified as unhealthy. The value range is [2-10]. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in ` + "`" + `tencentcloud_clb_listener_rule` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional, ForceNew) Port of the CLB listener.`,
				},
				resource.Attribute{
					Name:        "scheduler",
					Description: `(Optional) Scheduling method of the CLB listener, and available values are 'WRR' and 'LEAST_CONN'. The default is 'WRR'. NOTES: The listener of ` + "`" + `HTTP` + "`" + ` and ` + "`" + `HTTPS` + "`" + ` protocol additionally supports the ` + "`" + `IP Hash` + "`" + ` method. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in ` + "`" + `tencentcloud_clb_listener_rule` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "session_expire_time",
					Description: `(Optional) Time of session persistence within the CLB listener. NOTES: Available when scheduler is specified as ` + "`" + `WRR` + "`" + `, and not available when listener protocol is ` + "`" + `TCP_SSL` + "`" + `. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in ` + "`" + `tencentcloud_clb_listener_rule` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sni_switch",
					Description: `(Optional, ForceNew) Indicates whether SNI is enabled, and only supported with protocol ` + "`" + `HTTPS` + "`" + `. If enabled, you can set a certificate for each rule in ` + "`" + `tencentcloud_clb_listener_rule` + "`" + `, otherwise all rules have a certificate.`,
				},
				resource.Attribute{
					Name:        "target_type",
					Description: `(Optional, ForceNew) Backend target type. Valid values: ` + "`" + `NODE` + "`" + `, ` + "`" + `TARGETGROUP` + "`" + `. ` + "`" + `NODE` + "`" + ` means to bind ordinary nodes, ` + "`" + `TARGETGROUP` + "`" + ` means to bind target group. NOTES: TCP/UDP/TCP_SSL listener must configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `ID of this CLB listener. ## Import CLB listener can be imported using the id (version >= 1.47.0), e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_clb_listener.foo lb-7a0t6zqb#lbl-hh141sn9 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `ID of this CLB listener. ## Import CLB listener can be imported using the id (version >= 1.47.0), e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_clb_listener.foo lb-7a0t6zqb#lbl-hh141sn9 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_clb_listener_rule",
			Category:         "Cloud Load Balancer(CLB)",
			ShortDescription: `Provides a resource to create a CLB listener rule.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"load",
				"balancer",
				"clb",
				"listener",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "clb_id",
					Description: `(Required) ID of CLB instance.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) Domain name of the listener rule.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Required, ForceNew) ID of CLB listener.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Url of the listener rule.`,
				},
				resource.Attribute{
					Name:        "certificate_ca_id",
					Description: `(Optional) ID of the client certificate. NOTES: Only supports listeners of HTTPS protocol.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `(Optional) ID of the server certificate. NOTES: Only supports listeners of HTTPS protocol.`,
				},
				resource.Attribute{
					Name:        "certificate_ssl_mode",
					Description: `(Optional, ForceNew) Type of certificate. Valid values: ` + "`" + `UNIDIRECTIONAL` + "`" + `, ` + "`" + `MUTUAL` + "`" + `. NOTES: Only supports listeners of HTTPS protocol.`,
				},
				resource.Attribute{
					Name:        "forward_type",
					Description: `(Optional) Forwarding protocol between the CLB instance and real server. Valid values: ` + "`" + `HTTP` + "`" + `, ` + "`" + `HTTPS` + "`" + `, ` + "`" + `TRPC` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_check_health_num",
					Description: `(Optional) Health threshold of health check, and the default is ` + "`" + `3` + "`" + `. If a success result is returned for the health check 3 consecutive times, indicates that the forwarding is normal. The value range is [2-10]. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in ` + "`" + `tencentcloud_clb_listener_rule` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_check_http_code",
					Description: `(Optional) HTTP Status Code. The default is 31. Valid value ranges: [1~31]. ` + "`" + `1 means the return value '1xx' is health. ` + "`" + `2` + "`" + ` means the return value '2xx' is health. ` + "`" + `4` + "`" + ` means the return value '3xx' is health. ` + "`" + `8` + "`" + ` means the return value '4xx' is health. 16 means the return value '5xx' is health. If you want multiple return codes to indicate health, need to add the corresponding values. NOTES: The 'HTTP' health check of the 'TCP' listener only supports specifying one health check status code. NOTES: Only supports listeners of 'HTTP' and 'HTTPS' protocol.`,
				},
				resource.Attribute{
					Name:        "health_check_http_domain",
					Description: `(Optional) Domain name of health check. NOTES: Only supports listeners of ` + "`" + `HTTP` + "`" + ` and ` + "`" + `HTTPS` + "`" + ` protocol.`,
				},
				resource.Attribute{
					Name:        "health_check_http_method",
					Description: `(Optional) Methods of health check. NOTES: Only supports listeners of ` + "`" + `HTTP` + "`" + ` and ` + "`" + `HTTPS` + "`" + ` protocol. The default is ` + "`" + `HEAD` + "`" + `, the available value are ` + "`" + `HEAD` + "`" + ` and ` + "`" + `GET` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_check_http_path",
					Description: `(Optional) Path of health check. NOTES: Only supports listeners of ` + "`" + `HTTP` + "`" + ` and ` + "`" + `HTTPS` + "`" + ` protocol.`,
				},
				resource.Attribute{
					Name:        "health_check_interval_time",
					Description: `(Optional) Interval time of health check. Valid value ranges: (5~300) sec. and the default is ` + "`" + `5` + "`" + ` sec. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in ` + "`" + `tencentcloud_clb_listener_rule` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_check_switch",
					Description: `(Optional) Indicates whether health check is enabled.`,
				},
				resource.Attribute{
					Name:        "health_check_unhealth_num",
					Description: `(Optional) Unhealthy threshold of health check, and the default is ` + "`" + `3` + "`" + `. If the unhealthy result is returned 3 consecutive times, indicates that the forwarding is abnormal. The value range is [2-10]. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in ` + "`" + `tencentcloud_clb_listener_rule` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "http2_switch",
					Description: `(Optional) Indicate to apply HTTP2.0 protocol or not.`,
				},
				resource.Attribute{
					Name:        "scheduler",
					Description: `(Optional) Scheduling method of the CLB listener rules. Valid values: ` + "`" + `WRR` + "`" + `, ` + "`" + `IP HASH` + "`" + `, ` + "`" + `LEAST_CONN` + "`" + `. The default is ` + "`" + `WRR` + "`" + `. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in ` + "`" + `tencentcloud_clb_listener_rule` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "session_expire_time",
					Description: `(Optional) Time of session persistence within the CLB listener. NOTES: Available when scheduler is specified as ` + "`" + `WRR` + "`" + `, and not available when listener protocol is ` + "`" + `TCP_SSL` + "`" + `. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in ` + "`" + `tencentcloud_clb_listener_rule` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "target_type",
					Description: `(Optional, ForceNew) Backend target type. Valid values: ` + "`" + `NODE` + "`" + `, ` + "`" + `TARGETGROUP` + "`" + `. ` + "`" + `NODE` + "`" + ` means to bind ordinary nodes, ` + "`" + `TARGETGROUP` + "`" + ` means to bind target group. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `ID of this CLB listener rule. ## Import CLB listener rule can be imported using the id (version >= 1.47.0), e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_clb_listener_rule.foo lb-7a0t6zqb#lbl-hh141sn9#loc-agg236ys ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `ID of this CLB listener rule. ## Import CLB listener rule can be imported using the id (version >= 1.47.0), e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_clb_listener_rule.foo lb-7a0t6zqb#lbl-hh141sn9#loc-agg236ys ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_clb_redirection",
			Category:         "Cloud Load Balancer(CLB)",
			ShortDescription: `Provides a resource to create a CLB redirection.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"load",
				"balancer",
				"clb",
				"redirection",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "clb_id",
					Description: `(Required, ForceNew) ID of CLB instance.`,
				},
				resource.Attribute{
					Name:        "target_listener_id",
					Description: `(Required, ForceNew) ID of source listener.`,
				},
				resource.Attribute{
					Name:        "target_rule_id",
					Description: `(Required, ForceNew) Rule ID of target listener.`,
				},
				resource.Attribute{
					Name:        "delete_all_auto_rewrite",
					Description: `(Optional) Indicates whether delete all auto redirection. Default is ` + "`" + `false` + "`" + `. It will take effect only when this redirection is auto-rewrite and this auto-rewrite auto redirected more than one rules. All the auto-rewrite relations will be deleted when this parameter set true.`,
				},
				resource.Attribute{
					Name:        "is_auto_rewrite",
					Description: `(Optional, ForceNew) Indicates whether automatic forwarding is enable, default is ` + "`" + `false` + "`" + `. If enabled, the source listener and location should be empty, the target listener must be https protocol and port is 443.`,
				},
				resource.Attribute{
					Name:        "source_listener_id",
					Description: `(Optional, ForceNew) ID of source listener.`,
				},
				resource.Attribute{
					Name:        "source_rule_id",
					Description: `(Optional, ForceNew) Rule ID of source listener. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import CLB redirection can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_clb_redirection.foo loc-ft8fmngv#loc-4xxr2cy7#lbl-jc1dx6ju#lbl-asj1hzuo#lb-p7olt9e5 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import CLB redirection can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_clb_redirection.foo loc-ft8fmngv#loc-4xxr2cy7#lbl-jc1dx6ju#lbl-asj1hzuo#lb-p7olt9e5 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_clb_target_group",
			Category:         "Cloud Load Balancer(CLB)",
			ShortDescription: `Provides a resource to create a CLB target group.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"load",
				"balancer",
				"clb",
				"target",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "target_group_name",
					Description: `(Optional) Target group name.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional, ForceNew) VPC ID, default is based on the network. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import CLB target group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_clb_target_group.test lbtg-3k3io0i0 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import CLB target group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_clb_target_group.test lbtg-3k3io0i0 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_clb_target_group_attachment",
			Category:         "Cloud Load Balancer(CLB)",
			ShortDescription: `Provides a resource to create a CLB target group attachment is bound to the load balancing listener or forwarding rule.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"load",
				"balancer",
				"clb",
				"target",
				"group",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "clb_id",
					Description: `(Required, ForceNew) ID of the CLB.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Required, ForceNew) ID of the CLB listener.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `(Optional, ForceNew) ID of the CLB listener rule.`,
				},
				resource.Attribute{
					Name:        "target_group_id",
					Description: `(Optional, ForceNew) ID of the CLB target group.`,
				},
				resource.Attribute{
					Name:        "targrt_group_id",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import CLB target group attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_clb_target_group_attachment.group lbtg-odareyb2#lbl-bicjmx3i#lb-cv0iz74c#loc-ac6uk7b6 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import CLB target group attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_clb_target_group_attachment.group lbtg-odareyb2#lbl-bicjmx3i#lb-cv0iz74c#loc-ac6uk7b6 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_clb_target_group_instance_attachment",
			Category:         "Cloud Load Balancer(CLB)",
			ShortDescription: `Provides a resource to create a CLB target group instance attachment.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"load",
				"balancer",
				"clb",
				"target",
				"group",
				"instance",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bind_ip",
					Description: `(Required, ForceNew) The Intranet IP of the target group instance.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required, ForceNew) Port of the target group instance.`,
				},
				resource.Attribute{
					Name:        "target_group_id",
					Description: `(Required, ForceNew) Target group ID.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Required) The weight of the target group instance. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import CLB target group instance attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_clb_target_group_instance_attachment.test lbtg-3k3io0i0#172.16.48.18#222 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import CLB target group instance attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_clb_target_group_instance_attachment.test lbtg-3k3io0i0#172.16.48.18#222 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_container_cluster",
			Category:         "Container Cluster",
			ShortDescription: `Provides a TencentCloud Container Cluster resource.`,
			Description:      ``,
			Keywords: []string{
				"container",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bandwidth_type",
					Description: `(Required) The network type of the node.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Required) The network bandwidth of the node.`,
				},
				resource.Attribute{
					Name:        "cluster_cidr",
					Description: `(Required) The CIDR which the cluster is going to use.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Required) The name of the cluster.`,
				},
				resource.Attribute{
					Name:        "goods_num",
					Description: `(Required) The node number is going to create in the cluster.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Required) The instance type of the node needed by cvm.`,
				},
				resource.Attribute{
					Name:        "is_vpc_gateway",
					Description: `(Required) Describe whether the node enable the gateway capability.`,
				},
				resource.Attribute{
					Name:        "os_name",
					Description: `(Required) The system os name of the node.`,
				},
				resource.Attribute{
					Name:        "root_size",
					Description: `(Required) The size of the root volume.`,
				},
				resource.Attribute{
					Name:        "storage_size",
					Description: `(Required) The size of the data volume.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) The subnet id which the node stays in.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) Specify vpc which the node(s) stay in.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Required) The zone which the node stays in.`,
				},
				resource.Attribute{
					Name:        "cluster_desc",
					Description: `(Optional) The description of the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_version",
					Description: `(Optional) The kubernetes version of the cluster.`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "cvm_type",
					Description: `(Optional) The type of node needed by cvm.`,
				},
				resource.Attribute{
					Name:        "docker_graph_path",
					Description: `(Optional) The docker graph path is going to mounted.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Optional) The name ot node.`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `(Optional) The key_id of each node(if using key pair to access).`,
				},
				resource.Attribute{
					Name:        "mem",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "mount_target",
					Description: `(Optional) The path which volume is going to be mounted.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The password of each node.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) The puchase duration of the node needed by cvm.`,
				},
				resource.Attribute{
					Name:        "require_wan_ip",
					Description: `(Optional) Indicate whether wan ip is needed.`,
				},
				resource.Attribute{
					Name:        "root_type",
					Description: `(Optional) The type of the root volume. see more from CVM.`,
				},
				resource.Attribute{
					Name:        "sg_id",
					Description: `(Optional) The security group id.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `(Optional) The type of the data volume. see more from CVM.`,
				},
				resource.Attribute{
					Name:        "unschedulable",
					Description: `(Optional) Determine whether the node will be schedulable. 0 is the default meaning node will be schedulable. 1 for unschedulable.`,
				},
				resource.Attribute{
					Name:        "user_script",
					Description: `(Optional) User defined script in a base64-format. The script runs after the kubernetes component is ready on node. see more from CCS api documents. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "kubernetes_version",
					Description: `The kubernetes version of the cluster.`,
				},
				resource.Attribute{
					Name:        "nodes_num",
					Description: `The node number of the cluster.`,
				},
				resource.Attribute{
					Name:        "nodes_status",
					Description: `The node status of the cluster.`,
				},
				resource.Attribute{
					Name:        "total_cpu",
					Description: `The total cpu of the cluster.`,
				},
				resource.Attribute{
					Name:        "total_mem",
					Description: `The total memory of the cluster.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "kubernetes_version",
					Description: `The kubernetes version of the cluster.`,
				},
				resource.Attribute{
					Name:        "nodes_num",
					Description: `The node number of the cluster.`,
				},
				resource.Attribute{
					Name:        "nodes_status",
					Description: `The node status of the cluster.`,
				},
				resource.Attribute{
					Name:        "total_cpu",
					Description: `The total cpu of the cluster.`,
				},
				resource.Attribute{
					Name:        "total_mem",
					Description: `The total memory of the cluster.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_container_cluster_instance",
			Category:         "Container Cluster",
			ShortDescription: `Provides a TencentCloud Container Cluster Instance resource.`,
			Description:      ``,
			Keywords: []string{
				"container",
				"cluster",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bandwidth_type",
					Description: `(Required) The network type of the node.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Required) The network bandwidth of the node.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) The id of the cluster.`,
				},
				resource.Attribute{
					Name:        "is_vpc_gateway",
					Description: `(Required) Describe whether the node enable the gateway capability.`,
				},
				resource.Attribute{
					Name:        "root_size",
					Description: `(Required) The size of the root volume.`,
				},
				resource.Attribute{
					Name:        "storage_size",
					Description: `(Required) The size of the data volume.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) The subnet id which the node stays in.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Required) The zone which the node stays in.`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "cvm_type",
					Description: `(Optional) The type of node needed by cvm.`,
				},
				resource.Attribute{
					Name:        "docker_graph_path",
					Description: `(Optional) The docker graph path is going to mounted.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Optional) The name ot node.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Optional) The instance type of the node needed by cvm.`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `(Optional) The key_id of each node(if using key pair to access).`,
				},
				resource.Attribute{
					Name:        "mem",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "mount_target",
					Description: `(Optional) The path which volume is going to be mounted.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The password of each node.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) The puchase duration of the node needed by cvm.`,
				},
				resource.Attribute{
					Name:        "require_wan_ip",
					Description: `(Optional) Indicate whether wan ip is needed.`,
				},
				resource.Attribute{
					Name:        "root_type",
					Description: `(Optional) The type of the root volume. see more from CVM.`,
				},
				resource.Attribute{
					Name:        "sg_id",
					Description: `(Optional) The security group id.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `(Optional) The type of the data volume. see more from CVM.`,
				},
				resource.Attribute{
					Name:        "unschedulable",
					Description: `(Optional) Determine whether the node will be schedulable. 0 is the default meaning node will be schedulable. 1 for unschedulable.`,
				},
				resource.Attribute{
					Name:        "user_script",
					Description: `(Optional) User defined script in a base64-format. The script runs after the kubernetes component is ready on node. see more from CCS api documents. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "abnormal_reason",
					Description: `Describe the reason when node is in abnormal state(if it was).`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `An id identify the node, provided by cvm.`,
				},
				resource.Attribute{
					Name:        "is_normal",
					Description: `Describe whether the node is normal.`,
				},
				resource.Attribute{
					Name:        "lan_ip",
					Description: `Describe the lan ip of the node.`,
				},
				resource.Attribute{
					Name:        "wan_ip",
					Description: `Describe the wan ip of the node.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "abnormal_reason",
					Description: `Describe the reason when node is in abnormal state(if it was).`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `An id identify the node, provided by cvm.`,
				},
				resource.Attribute{
					Name:        "is_normal",
					Description: `Describe whether the node is normal.`,
				},
				resource.Attribute{
					Name:        "lan_ip",
					Description: `Describe the lan ip of the node.`,
				},
				resource.Attribute{
					Name:        "wan_ip",
					Description: `Describe the wan ip of the node.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cos_bucket",
			Category:         "Cloud Object Storage(COS)",
			ShortDescription: `Provides a COS resource to create a COS bucket and set its attributes.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"object",
				"storage",
				"cos",
				"bucket",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required, ForceNew) The name of a bucket to be created. Bucket format should be [custom name]-[appid], for example ` + "`" + `mycos-1258798060` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `(Optional) The canned ACL to apply. Valid values: private, public-read, and public-read-write. Defaults to private.`,
				},
				resource.Attribute{
					Name:        "cors_rules",
					Description: `(Optional) A rule of Cross-Origin Resource Sharing (documented below).`,
				},
				resource.Attribute{
					Name:        "encryption_algorithm",
					Description: `(Optional) The server-side encryption algorithm to use. Valid value is ` + "`" + `AES256` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lifecycle_rules",
					Description: `(Optional) A configuration of object lifecycle management (documented below).`,
				},
				resource.Attribute{
					Name:        "log_enable",
					Description: `(Optional) Indicate the access log of this bucket to be saved or not. Default is ` + "`" + `false` + "`" + `. If set ` + "`" + `true` + "`" + `, the access log will be saved with ` + "`" + `log_target_bucket` + "`" + `. To enable log, the full access of log service must be granted. [Full Access Role Policy](https://intl.cloud.tencent.com/document/product/436/16920).`,
				},
				resource.Attribute{
					Name:        "log_prefix",
					Description: `(Optional) The prefix log name which saves the access log of this bucket per 5 minutes. Eg. ` + "`" + `MyLogPrefix/` + "`" + `. The log access file format is ` + "`" + `log_target_bucket` + "`" + `/` + "`" + `log_prefix` + "`" + `{YYYY}/{MM}/{DD}/{time}_{random}_{index}.gz. Only valid when ` + "`" + `log_enable` + "`" + ` is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "log_target_bucket",
					Description: `(Optional) The target bucket name which saves the access log of this bucket per 5 minutes. The log access file format is ` + "`" + `log_target_bucket` + "`" + `/` + "`" + `log_prefix` + "`" + `{YYYY}/{MM}/{DD}/{time}_{random}_{index}.gz. Only valid when ` + "`" + `log_enable` + "`" + ` is ` + "`" + `true` + "`" + `. User must have full access on this bucket.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The tags of a bucket.`,
				},
				resource.Attribute{
					Name:        "versioning_enable",
					Description: `(Optional) Enable bucket versioning.`,
				},
				resource.Attribute{
					Name:        "website",
					Description: `(Optional) A website object(documented below). The ` + "`" + `cors_rules` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "allowed_headers",
					Description: `(Required) Specifies which headers are allowed.`,
				},
				resource.Attribute{
					Name:        "allowed_methods",
					Description: `(Required) Specifies which methods are allowed. Can be ` + "`" + `GET` + "`" + `, ` + "`" + `PUT` + "`" + `, ` + "`" + `POST` + "`" + `, ` + "`" + `DELETE` + "`" + ` or ` + "`" + `HEAD` + "`" + `.`,
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
					Description: `(Optional) Specifies time in seconds that browser can cache the response for a preflight request. The ` + "`" + `expiration` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "date",
					Description: `(Optional) Specifies the date after which you want the corresponding action to take effect.`,
				},
				resource.Attribute{
					Name:        "days",
					Description: `(Optional) Specifies the number of days after object creation when the specific rule action takes effect. The ` + "`" + `lifecycle_rules` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "filter_prefix",
					Description: `(Required) Object key prefix identifying one or more objects to which the rule applies.`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `(Optional) Specifies a period in the object's expire (documented below).`,
				},
				resource.Attribute{
					Name:        "transition",
					Description: `(Optional) Specifies a period in the object's transitions (documented below). The ` + "`" + `transition` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `(Required) Specifies the storage class to which you want the object to transition. Available values include ` + "`" + `STANDARD` + "`" + `, ` + "`" + `STANDARD_IA` + "`" + ` and ` + "`" + `ARCHIVE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "date",
					Description: `(Optional) Specifies the date after which you want the corresponding action to take effect.`,
				},
				resource.Attribute{
					Name:        "days",
					Description: `(Optional) Specifies the number of days after object creation when the specific rule action takes effect. The ` + "`" + `website` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "error_document",
					Description: `(Optional) An absolute path to the document to return in case of a 4XX error.`,
				},
				resource.Attribute{
					Name:        "index_document",
					Description: `(Optional) COS returns this index document when requests are made to the root domain or any of the subfolders. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "cos_bucket_url",
					Description: `The URL of this cos bucket. ## Import COS bucket can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cos_bucket.bucket bucket-name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "cos_bucket_url",
					Description: `The URL of this cos bucket. ## Import COS bucket can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cos_bucket.bucket bucket-name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cos_bucket_object",
			Category:         "Cloud Object Storage(COS)",
			ShortDescription: `Provides a COS object resource to put an object(content or file) to the bucket.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"object",
				"storage",
				"cos",
				"bucket",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required, ForceNew) The name of a bucket to use. Bucket format should be [custom name]-[appid], for example ` + "`" + `mycos-1258798060` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required, ForceNew) The name of the object once it is in the bucket.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `(Optional) The canned ACL to apply. Available values include ` + "`" + `private` + "`" + `, ` + "`" + `public-read` + "`" + `, and ` + "`" + `public-read-write` + "`" + `. Defaults to ` + "`" + `private` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cache_control",
					Description: `(Optional) Specifies caching behavior along the request/reply chain. For further details, RFC2616 can be referred.`,
				},
				resource.Attribute{
					Name:        "content_disposition",
					Description: `(Optional) Specifies presentational information for the object.`,
				},
				resource.Attribute{
					Name:        "content_encoding",
					Description: `(Optional) Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `(Optional) A standard MIME type describing the format of the object data.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Optional) Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `(Optional) The ETag generated for the object (an MD5 sum of the object content).`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) The path to the source file being uploaded to the bucket.`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `(Optional) Object storage type, Available values include ` + "`" + `STANDARD` + "`" + `, ` + "`" + `STANDARD_IA` + "`" + ` and ` + "`" + `ARCHIVE` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cos_bucket_policy",
			Category:         "Cloud Object Storage(COS)",
			ShortDescription: `Provides a COS resource to create a COS bucket policy and set its attributes.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"object",
				"storage",
				"cos",
				"bucket",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required, ForceNew) The name of a bucket to be created. Bucket format should be [custom name]-[appid], for example ` + "`" + `mycos-1258798060` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Required) The text of the policy. For more info please refer to [Tencent official doc](https://intl.cloud.tencent.com/document/product/436/18023). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import COS bucket policy can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cos_bucket_policy.bucket bucket-name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import COS bucket policy can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cos_bucket_policy.bucket bucket-name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cynosdb_cluster",
			Category:         "CynosDB",
			ShortDescription: `Provide a resource to create a CynosDB cluster.`,
			Description:      ``,
			Keywords: []string{
				"cynosdb",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "available_zone",
					Description: `(Required, ForceNew) The available zone of the CynosDB Cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Required, ForceNew) Name of CynosDB cluster.`,
				},
				resource.Attribute{
					Name:        "db_type",
					Description: `(Required, ForceNew) Type of CynosDB, and available values include ` + "`" + `MYSQL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "db_version",
					Description: `(Required, ForceNew) Version of CynosDB, which is related to ` + "`" + `db_type` + "`" + `. For ` + "`" + `MYSQL` + "`" + `, available value is ` + "`" + `5.7` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_cpu_core",
					Description: `(Required) The number of CPU cores of read-write type instance in the CynosDB cluster. Note: modification of this field will take effect immediately, if want to upgrade on maintenance window, please upgrade from console.`,
				},
				resource.Attribute{
					Name:        "instance_memory_size",
					Description: `(Required) Memory capacity of read-write type instance, unit in GB. Note: modification of this field will take effect immediately, if want to upgrade on maintenance window, please upgrade from console.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required, ForceNew) Password of ` + "`" + `root` + "`" + ` account.`,
				},
				resource.Attribute{
					Name:        "storage_limit",
					Description: `(Required, ForceNew) Storage limit of CynosDB cluster instance, unit in GB.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required, ForceNew) ID of the subnet within this VPC.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, ForceNew) ID of the VPC.`,
				},
				resource.Attribute{
					Name:        "auto_renew_flag",
					Description: `(Optional, ForceNew) Auto renew flag. Valid values are ` + "`" + `0` + "`" + `(MANUAL_RENEW), ` + "`" + `1` + "`" + `(AUTO_RENEW). Default value is ` + "`" + `0` + "`" + `. Only works for PREPAID cluster.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `(Optional, ForceNew) The charge type of instance. Valid values are ` + "`" + `PREPAID` + "`" + ` and ` + "`" + `POSTPAID_BY_HOUR` + "`" + `. Default value is ` + "`" + `POSTPAID_BY_HOUR` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "force_delete",
					Description: `(Optional) Indicate whether to delete cluster instance directly or not. Default is false. If set true, the cluster and its ` + "`" + `All RELATED INSTANCES` + "`" + ` will be deleted instead of staying recycle bin. Note: works for both ` + "`" + `PREPAID` + "`" + ` and ` + "`" + `POSTPAID_BY_HOUR` + "`" + ` cluster.`,
				},
				resource.Attribute{
					Name:        "instance_maintain_duration",
					Description: `(Optional) Duration time for maintenance, unit in second. ` + "`" + `3600` + "`" + ` by default.`,
				},
				resource.Attribute{
					Name:        "instance_maintain_start_time",
					Description: `(Optional) Offset time from 00:00, unit in second. For example, 03:00am should be ` + "`" + `10800` + "`" + `. ` + "`" + `10800` + "`" + ` by default.`,
				},
				resource.Attribute{
					Name:        "instance_maintain_weekdays",
					Description: `(Optional) Weekdays for maintenance. ` + "`" + `["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"]` + "`" + ` by default.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional, ForceNew) Port of CynosDB cluster.`,
				},
				resource.Attribute{
					Name:        "prepaid_period",
					Description: `(Optional, ForceNew) The tenancy (time unit is month) of the prepaid instance. Valid values are ` + "`" + `1` + "`" + `, ` + "`" + `2` + "`" + `, ` + "`" + `3` + "`" + `, ` + "`" + `4` + "`" + `, ` + "`" + `5` + "`" + `, ` + "`" + `6` + "`" + `, ` + "`" + `7` + "`" + `, ` + "`" + `8` + "`" + `, ` + "`" + `9` + "`" + `, ` + "`" + `10` + "`" + `, ` + "`" + `11` + "`" + `, ` + "`" + `12` + "`" + `, ` + "`" + `24` + "`" + `, ` + "`" + `36` + "`" + `. NOTE: it only works when charge_type is set to ` + "`" + `PREPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional, ForceNew) ID of the project. ` + "`" + `0` + "`" + ` by default.`,
				},
				resource.Attribute{
					Name:        "ro_group_sg",
					Description: `(Optional) IDs of security group for ` + "`" + `ro_group` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rw_group_sg",
					Description: `(Optional) IDs of security group for ` + "`" + `rw_group` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The tags of the CynosDB cluster. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "charset",
					Description: `Charset used by CynosDB cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_status",
					Description: `Status of the Cynosdb cluster.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the CynosDB cluster.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of instance.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `Name of instance.`,
				},
				resource.Attribute{
					Name:        "instance_status",
					Description: `Status of the instance.`,
				},
				resource.Attribute{
					Name:        "instance_storage_size",
					Description: `Storage size of the instance, unit in GB.`,
				},
				resource.Attribute{
					Name:        "ro_group_addr",
					Description: `Readonly addresses. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `IP address for readonly connection.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port number for readonly connection.`,
				},
				resource.Attribute{
					Name:        "ro_group_id",
					Description: `ID of read-only instance group.`,
				},
				resource.Attribute{
					Name:        "ro_group_instances",
					Description: `List of instances in the read-only instance group.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of instance.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `Name of instance.`,
				},
				resource.Attribute{
					Name:        "rw_group_addr",
					Description: `Read-write addresses. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `IP address for read-write connection.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port number for read-write connection.`,
				},
				resource.Attribute{
					Name:        "rw_group_id",
					Description: `ID of read-write instance group.`,
				},
				resource.Attribute{
					Name:        "rw_group_instances",
					Description: `List of instances in the read-write instance group.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of instance.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `Name of instance.`,
				},
				resource.Attribute{
					Name:        "storage_used",
					Description: `Used storage of CynosDB cluster, unit in MB. ## Import CynosDB cluster can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cynosdb_cluster.foo cynosdbmysql-dzj5l8gz ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "charset",
					Description: `Charset used by CynosDB cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_status",
					Description: `Status of the Cynosdb cluster.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the CynosDB cluster.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of instance.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `Name of instance.`,
				},
				resource.Attribute{
					Name:        "instance_status",
					Description: `Status of the instance.`,
				},
				resource.Attribute{
					Name:        "instance_storage_size",
					Description: `Storage size of the instance, unit in GB.`,
				},
				resource.Attribute{
					Name:        "ro_group_addr",
					Description: `Readonly addresses. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `IP address for readonly connection.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port number for readonly connection.`,
				},
				resource.Attribute{
					Name:        "ro_group_id",
					Description: `ID of read-only instance group.`,
				},
				resource.Attribute{
					Name:        "ro_group_instances",
					Description: `List of instances in the read-only instance group.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of instance.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `Name of instance.`,
				},
				resource.Attribute{
					Name:        "rw_group_addr",
					Description: `Read-write addresses. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `IP address for read-write connection.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port number for read-write connection.`,
				},
				resource.Attribute{
					Name:        "rw_group_id",
					Description: `ID of read-write instance group.`,
				},
				resource.Attribute{
					Name:        "rw_group_instances",
					Description: `List of instances in the read-write instance group.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of instance.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `Name of instance.`,
				},
				resource.Attribute{
					Name:        "storage_used",
					Description: `Used storage of CynosDB cluster, unit in MB. ## Import CynosDB cluster can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cynosdb_cluster.foo cynosdbmysql-dzj5l8gz ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cynosdb_readonly_instance",
			Category:         "CynosDB",
			ShortDescription: `Provide a resource to create a CynosDB readonly instance.`,
			Description:      ``,
			Keywords: []string{
				"cynosdb",
				"readonly",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required, ForceNew) Cluster ID which the readonly instance belongs to.`,
				},
				resource.Attribute{
					Name:        "instance_cpu_core",
					Description: `(Required) The number of CPU cores of read-write type instance in the CynosDB cluster. Note: modification of this field will take effect immediately, if want to upgrade on maintenance window, please upgrade from console.`,
				},
				resource.Attribute{
					Name:        "instance_memory_size",
					Description: `(Required) Memory capacity of read-write type instance, unit in GB. Note: modification of this field will take effect immediately, if want to upgrade on maintenance window, please upgrade from console.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Required, ForceNew) Name of instance.`,
				},
				resource.Attribute{
					Name:        "force_delete",
					Description: `(Optional) Indicate whether to delete readonly instance directly or not. Default is false. If set true, instance will be deleted instead of staying recycle bin. Note: works for both ` + "`" + `PREPAID` + "`" + ` and ` + "`" + `POSTPAID_BY_HOUR` + "`" + ` cluster.`,
				},
				resource.Attribute{
					Name:        "instance_maintain_duration",
					Description: `(Optional) Duration time for maintenance, unit in second. ` + "`" + `3600` + "`" + ` by default.`,
				},
				resource.Attribute{
					Name:        "instance_maintain_start_time",
					Description: `(Optional) Offset time from 00:00, unit in second. For example, 03:00am should be ` + "`" + `10800` + "`" + `. ` + "`" + `10800` + "`" + ` by default.`,
				},
				resource.Attribute{
					Name:        "instance_maintain_weekdays",
					Description: `(Optional) Weekdays for maintenance. ` + "`" + `["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"]` + "`" + ` by default. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "instance_status",
					Description: `Status of the instance.`,
				},
				resource.Attribute{
					Name:        "instance_storage_size",
					Description: `Storage size of the instance, unit in GB. ## Import CynosDB readonly instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cynosdb_readonly_instance.foo cynosdbmysql-ins-dhwynib6 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "instance_status",
					Description: `Status of the instance.`,
				},
				resource.Attribute{
					Name:        "instance_storage_size",
					Description: `Storage size of the instance, unit in GB. ## Import CynosDB readonly instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cynosdb_readonly_instance.foo cynosdbmysql-ins-dhwynib6 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dayu_cc_http_policy",
			Category:         "Anti-DDoS(Dayu)",
			ShortDescription: `Use this resource to create a dayu CC self-define http policy`,
			Description:      ``,
			Keywords: []string{
				"anti",
				"ddos",
				"dayu",
				"cc",
				"http",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Name of the CC self-define http policy. Length should between 1 and 20.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required, ForceNew) ID of the resource that the CC self-define http policy works for.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `(Required, ForceNew) Type of the resource that the CC self-define http policy works for, valid values are ` + "`" + `bgpip` + "`" + `, ` + "`" + `bgp` + "`" + `, ` + "`" + `bgp-multip` + "`" + ` and ` + "`" + `net` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Action mode, only valid when ` + "`" + `smode` + "`" + ` is ` + "`" + `matching` + "`" + `. Valid values are ` + "`" + `alg` + "`" + ` and ` + "`" + `drop` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "frequency",
					Description: `(Optional) Max frequency per minute, only valid when ` + "`" + `smode` + "`" + ` is ` + "`" + `speedlimit` + "`" + `, the valid value ranges from 1 to 10000.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Optional) Ip of the CC self-define http policy, only valid when ` + "`" + `resource_type` + "`" + ` is ` + "`" + `bgp-multip` + "`" + `. The num of list items can only be set one.`,
				},
				resource.Attribute{
					Name:        "rule_list",
					Description: `(Optional) Rule list of the CC self-define http policy, only valid when ` + "`" + `smode` + "`" + ` is ` + "`" + `matching` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "smode",
					Description: `(Optional) Match mode, and valid values are ` + "`" + `matching` + "`" + `, ` + "`" + `speedlimit` + "`" + `. Note: the speed limit type CC self-define policy can only set one.`,
				},
				resource.Attribute{
					Name:        "switch",
					Description: `(Optional) Indicate the CC self-define http policy takes effect or not. The ` + "`" + `rule_list` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `(Optional) Operator of the rule. Valid values: ` + "`" + `include` + "`" + `, ` + "`" + `not_include` + "`" + `, ` + "`" + `equal` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "skey",
					Description: `(Optional) Key of the rule. Valid values: ` + "`" + `host` + "`" + `, ` + "`" + `cgi` + "`" + `, ` + "`" + `ua` + "`" + `, ` + "`" + `referer` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Rule value, then length should be less than 31 bytes. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CC self-define http policy.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Id of the CC self-define http policy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CC self-define http policy.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Id of the CC self-define http policy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dayu_cc_https_policy",
			Category:         "Anti-DDoS(Dayu)",
			ShortDescription: `Use this resource to create a dayu CC self-define https policy`,
			Description:      ``,
			Keywords: []string{
				"anti",
				"ddos",
				"dayu",
				"cc",
				"https",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `(Required, ForceNew) Domain that the CC self-define https policy works for, only valid when ` + "`" + `protocol` + "`" + ` is ` + "`" + `https` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Name of the CC self-define https policy. Length should between 1 and 20.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required, ForceNew) ID of the resource that the CC self-define https policy works for.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `(Required, ForceNew) Type of the resource that the CC self-define https policy works for, valid value is ` + "`" + `bgpip` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `(Required, ForceNew) Rule id of the domain that the CC self-define https policy works for, only valid when ` + "`" + `protocol` + "`" + ` is ` + "`" + `https` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rule_list",
					Description: `(Required) Rule list of the CC self-define https policy.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Action mode. Valid values are ` + "`" + `alg` + "`" + ` and ` + "`" + `drop` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "switch",
					Description: `(Optional) Indicate the CC self-define https policy takes effect or not. The ` + "`" + `rule_list` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `(Required) Operator of the rule. Valid values are ` + "`" + `include` + "`" + ` and ` + "`" + `equal` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "skey",
					Description: `(Required) Key of the rule. Valid values are ` + "`" + `cgi` + "`" + `, ` + "`" + `ua` + "`" + ` and ` + "`" + `referer` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Rule value, then length should be less than 31 bytes. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CC self-define https policy.`,
				},
				resource.Attribute{
					Name:        "ip_list",
					Description: `Ip of the CC self-define https policy.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Id of the CC self-define https policy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CC self-define https policy.`,
				},
				resource.Attribute{
					Name:        "ip_list",
					Description: `Ip of the CC self-define https policy.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Id of the CC self-define https policy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dayu_ddos_policy",
			Category:         "Anti-DDoS(Dayu)",
			ShortDescription: `Use this resource to create dayu DDoS policy`,
			Description:      ``,
			Keywords: []string{
				"anti",
				"ddos",
				"dayu",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "drop_options",
					Description: `(Required) Option list of abnormal check of the DDos policy, should set at least one policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the DDoS policy. Length should between 1 and 32.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `(Required, ForceNew) Type of the resource that the DDoS policy works for. Valid values: ` + "`" + `bgpip` + "`" + `, ` + "`" + `bgp` + "`" + `, ` + "`" + `bgp-multip` + "`" + ` and ` + "`" + `net` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "black_ips",
					Description: `(Optional) Black IP list.`,
				},
				resource.Attribute{
					Name:        "packet_filters",
					Description: `(Optional) Message filter options list.`,
				},
				resource.Attribute{
					Name:        "port_filters",
					Description: `(Optional) Port limits of abnormal check of the DDos policy.`,
				},
				resource.Attribute{
					Name:        "watermark_filters",
					Description: `(Optional) Watermark policy options, and only support one watermark policy at most.`,
				},
				resource.Attribute{
					Name:        "white_ips",
					Description: `(Optional) White IP list. The ` + "`" + `drop_options` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "bad_conn_threshold",
					Description: `(Required) The number of new connections based on destination IP that trigger suppression of connections. Valid value ranges: (0~4294967295).`,
				},
				resource.Attribute{
					Name:        "check_sync_conn",
					Description: `(Required) Indicate whether to check null connection or not.`,
				},
				resource.Attribute{
					Name:        "conn_timeout",
					Description: `(Required) Connection timeout of abnormal connection check. Valid value ranges: (0~65535).`,
				},
				resource.Attribute{
					Name:        "d_conn_limit",
					Description: `(Required) The limit of concurrent connections based on destination IP. Valid value ranges: (0~4294967295).`,
				},
				resource.Attribute{
					Name:        "d_new_limit",
					Description: `(Required) The limit of new connections based on destination IP. Valid value ranges: (0~4294967295).`,
				},
				resource.Attribute{
					Name:        "drop_abroad",
					Description: `(Required) Indicate whether to drop abroad traffic or not.`,
				},
				resource.Attribute{
					Name:        "drop_icmp",
					Description: `(Required) Indicate whether to drop ICMP protocol or not.`,
				},
				resource.Attribute{
					Name:        "drop_other",
					Description: `(Required) Indicate whether to drop other protocols(exclude TCP/UDP/ICMP) or not.`,
				},
				resource.Attribute{
					Name:        "drop_tcp",
					Description: `(Required) Indicate whether to drop TCP protocol or not.`,
				},
				resource.Attribute{
					Name:        "drop_udp",
					Description: `(Required) Indicate to drop UDP protocol or not.`,
				},
				resource.Attribute{
					Name:        "icmp_mbps_limit",
					Description: `(Required) The limit of ICMP traffic rate. Valid value ranges: (0~4294967295)(Mbps).`,
				},
				resource.Attribute{
					Name:        "null_conn_enable",
					Description: `(Required) Indicate to enable null connection or not.`,
				},
				resource.Attribute{
					Name:        "other_mbps_limit",
					Description: `(Required) The limit of other protocols(exclude TCP/UDP/ICMP) traffic rate. Valid value ranges: (0~4294967295)(Mbps).`,
				},
				resource.Attribute{
					Name:        "s_conn_limit",
					Description: `(Required) The limit of concurrent connections based on source IP. Valid value ranges: (0~4294967295).`,
				},
				resource.Attribute{
					Name:        "s_new_limit",
					Description: `(Required) The limit of new connections based on source IP. Valid value ranges: (0~4294967295).`,
				},
				resource.Attribute{
					Name:        "syn_limit",
					Description: `(Required) The limit of syn of abnormal connection check. Valid value ranges: (0~100).`,
				},
				resource.Attribute{
					Name:        "tcp_mbps_limit",
					Description: `(Required) The limit of TCP traffic. Valid value ranges: (0~4294967295)(Mbps).`,
				},
				resource.Attribute{
					Name:        "udp_mbps_limit",
					Description: `(Required) The limit of UDP traffic rate. Valid value ranges: (0~4294967295)(Mbps).`,
				},
				resource.Attribute{
					Name:        "syn_rate",
					Description: `(Optional) The percentage of syn in ack of abnormal connection check. Valid value ranges: (0~100). The ` + "`" + `packet_filters` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Action of port to take. Valid values: ` + "`" + `drop` + "`" + `, ` + "`" + `drop_black` + "`" + `,` + "`" + `drop_rst` + "`" + `,` + "`" + `drop_black_rst` + "`" + `,` + "`" + `transmit` + "`" + `.` + "`" + `drop` + "`" + `(drop the packet), ` + "`" + `drop_black` + "`" + `(drop the packet and black the ip),` + "`" + `drop_rst` + "`" + `(drop the packet and disconnect),` + "`" + `drop_black_rst` + "`" + `(drop the packet, black the ip and disconnect),` + "`" + `transmit` + "`" + `(transmit the packet).`,
				},
				resource.Attribute{
					Name:        "d_end_port",
					Description: `(Optional) End port of the destination. Valid value ranges: (0~65535). It must be greater than ` + "`" + `d_start_port` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "d_start_port",
					Description: `(Optional) Start port of the destination. Valid value ranges: (0~65535).`,
				},
				resource.Attribute{
					Name:        "depth",
					Description: `(Optional) The depth of match. Valid value ranges: (0~1500).`,
				},
				resource.Attribute{
					Name:        "is_include",
					Description: `(Optional) Indicate whether to include the key word/regular expression or not.`,
				},
				resource.Attribute{
					Name:        "match_begin",
					Description: `(Optional) Indicate whether to check load or not, ` + "`" + `begin_l5` + "`" + ` means to match and ` + "`" + `no_match` + "`" + ` means not.`,
				},
				resource.Attribute{
					Name:        "match_str",
					Description: `(Optional) The key word or regular expression.`,
				},
				resource.Attribute{
					Name:        "match_type",
					Description: `(Optional) Match type. Valid values: ` + "`" + `sunday` + "`" + ` and ` + "`" + `pcre` + "`" + `. ` + "`" + `sunday` + "`" + ` means key word match while ` + "`" + `pcre` + "`" + ` means regular match.`,
				},
				resource.Attribute{
					Name:        "offset",
					Description: `(Optional) The offset of match. Valid value ranges: (0~1500).`,
				},
				resource.Attribute{
					Name:        "pkt_length_max",
					Description: `(Optional) The max length of the packet. Valid value ranges: (0~1500)(Mbps). It must be greater than ` + "`" + `pkt_length_min` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "pkt_length_min",
					Description: `(Optional) The minimum length of the packet. Valid value ranges: (0~1500)(Mbps).`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) Protocol. Valid values: ` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `, ` + "`" + `all` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "s_end_port",
					Description: `(Optional) End port of the source. Valid value ranges: (0~65535). It must be greater than ` + "`" + `s_start_port` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "s_start_port",
					Description: `(Optional) Start port of the source. Valid value ranges: (0~65535). The ` + "`" + `port_filters` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Action of port to take. Valid values: ` + "`" + `drop` + "`" + `, ` + "`" + `transmit` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "end_port",
					Description: `(Optional) End port. Valid value ranges: (0~65535). It must be greater than ` + "`" + `start_port` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `(Optional) The type of forbidden port. Valid values: ` + "`" + `0` + "`" + `, ` + "`" + `1` + "`" + `, ` + "`" + `2` + "`" + `. ` + "`" + `0` + "`" + ` for destination ports make effect, ` + "`" + `1` + "`" + ` for source ports make effect. ` + "`" + `2` + "`" + ` for both destination and source ports.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) Protocol. Valid values are ` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `, ` + "`" + `all` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "start_port",
					Description: `(Optional) Start port. Valid value ranges: (0~65535). The ` + "`" + `watermark_filters` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "auto_remove",
					Description: `(Optional) Indicate whether to auto-remove the watermark or not.`,
				},
				resource.Attribute{
					Name:        "offset",
					Description: `(Optional) The offset of watermark. Valid value ranges: (0~1500).`,
				},
				resource.Attribute{
					Name:        "open_switch",
					Description: `(Optional) Indicate whether to open watermark or not. It muse be set ` + "`" + `true` + "`" + ` when any field of watermark was set.`,
				},
				resource.Attribute{
					Name:        "tcp_port_list",
					Description: `(Optional) Port range of TCP, the format is like ` + "`" + `2000-3000` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "udp_port_list",
					Description: `(Optional) Port range of TCP, the format is like ` + "`" + `2000-3000` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the DDoS policy.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Id of policy.`,
				},
				resource.Attribute{
					Name:        "scene_id",
					Description: `Id of policy case that the DDoS policy works for.`,
				},
				resource.Attribute{
					Name:        "watermark_key",
					Description: `Watermark content.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `Content of the watermark.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of the watermark.`,
				},
				resource.Attribute{
					Name:        "open_switch",
					Description: `Indicate whether to auto-remove the watermark or not.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the DDoS policy.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Id of policy.`,
				},
				resource.Attribute{
					Name:        "scene_id",
					Description: `Id of policy case that the DDoS policy works for.`,
				},
				resource.Attribute{
					Name:        "watermark_key",
					Description: `Watermark content.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `Content of the watermark.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of the watermark.`,
				},
				resource.Attribute{
					Name:        "open_switch",
					Description: `Indicate whether to auto-remove the watermark or not.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dayu_ddos_policy_attachment",
			Category:         "Anti-DDoS(Dayu)",
			ShortDescription: `Provides a resource to create a dayu DDoS policy attachment.`,
			Description:      ``,
			Keywords: []string{
				"anti",
				"ddos",
				"dayu",
				"policy",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Required, ForceNew) ID of the policy.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required, ForceNew) ID of the attached resource.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `(Required, ForceNew) Type of the resource that the DDoS policy works for. Valid values are ` + "`" + `bgpip` + "`" + `, ` + "`" + `bgp` + "`" + `, ` + "`" + `bgp-multip` + "`" + `, ` + "`" + `net` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dayu_ddos_policy_case",
			Category:         "Anti-DDoS(Dayu)",
			ShortDescription: `Use this resource to create dayu DDoS policy case`,
			Description:      ``,
			Keywords: []string{
				"anti",
				"ddos",
				"dayu",
				"policy",
				"case",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "app_protocols",
					Description: `(Required) App protocol set of the DDoS policy case.`,
				},
				resource.Attribute{
					Name:        "app_type",
					Description: `(Required) App type of the DDoS policy case. Valid values: ` + "`" + `WEB` + "`" + `, ` + "`" + `GAME` + "`" + `, ` + "`" + `APP` + "`" + ` and ` + "`" + `OTHER` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "has_abroad",
					Description: `(Required) Indicate whether the service involves overseas or not. Valid values: ` + "`" + `no` + "`" + ` and ` + "`" + `yes` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "has_initiate_tcp",
					Description: `(Required) Indicate whether the service actively initiates TCP requests or not. Valid values: ` + "`" + `no` + "`" + ` and ` + "`" + `yes` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Name of the DDoS policy case. Length should between 1 and 64.`,
				},
				resource.Attribute{
					Name:        "platform_types",
					Description: `(Required) Platform set of the DDoS policy case.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `(Required, ForceNew) Type of the resource that the DDoS policy case works for. Valid values: ` + "`" + `bgpip` + "`" + `, ` + "`" + `bgp` + "`" + ` and ` + "`" + `bgp-multip` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tcp_end_port",
					Description: `(Required) End port of the TCP service. Valid value ranges: (0~65535). It must be greater than ` + "`" + `tcp_start_port` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tcp_start_port",
					Description: `(Required) Start port of the TCP service. Valid value ranges: (0~65535).`,
				},
				resource.Attribute{
					Name:        "udp_end_port",
					Description: `(Required) End port of the UDP service. Valid value ranges: (0~65535). It must be greater than ` + "`" + `udp_start_port` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "udp_start_port",
					Description: `(Required) Start port of the UDP service. Valid value ranges: (0~65535).`,
				},
				resource.Attribute{
					Name:        "web_api_urls",
					Description: `(Required) Web API url set.`,
				},
				resource.Attribute{
					Name:        "has_initiate_udp",
					Description: `(Optional) Indicate whether the actively initiate UDP requests or not. Valid values: ` + "`" + `no` + "`" + ` and ` + "`" + `yes` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "has_vpn",
					Description: `(Optional) Indicate whether the service involves VPN service or not. Valid values: ` + "`" + `no` + "`" + ` and ` + "`" + `yes` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max_tcp_package_len",
					Description: `(Optional) The max length of TCP message package, valid value length should be greater than 0 and less than 1500. It should be greater than ` + "`" + `min_tcp_package_len` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max_udp_package_len",
					Description: `(Optional) The max length of UDP message package, valid value length should be greater than 0 and less than 1500. It should be greater than ` + "`" + `min_udp_package_len` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "min_tcp_package_len",
					Description: `(Optional) The minimum length of TCP message package, valid value length should be greater than 0 and less than 1500.`,
				},
				resource.Attribute{
					Name:        "min_udp_package_len",
					Description: `(Optional) The minimum length of UDP message package, valid value length should be greater than 0 and less than 1500.`,
				},
				resource.Attribute{
					Name:        "peer_tcp_port",
					Description: `(Optional) The port that actively initiates TCP requests. Valid value ranges: (1~65535).`,
				},
				resource.Attribute{
					Name:        "peer_udp_port",
					Description: `(Optional) The port that actively initiates UDP requests. Valid value ranges: (1~65535).`,
				},
				resource.Attribute{
					Name:        "tcp_footprint",
					Description: `(Optional) The fixed signature of TCP protocol load, valid value length is range from 1 to 512.`,
				},
				resource.Attribute{
					Name:        "udp_footprint",
					Description: `(Optional) The fixed signature of TCP protocol load, valid value length is range from 1 to 512. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the DDoS policy case.`,
				},
				resource.Attribute{
					Name:        "scene_id",
					Description: `ID of the DDoS policy case.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the DDoS policy case.`,
				},
				resource.Attribute{
					Name:        "scene_id",
					Description: `ID of the DDoS policy case.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dayu_l4_rule",
			Category:         "Anti-DDoS(Dayu)",
			ShortDescription: `Use this resource to create dayu layer 4 rule`,
			Description:      ``,
			Keywords: []string{
				"anti",
				"ddos",
				"dayu",
				"l4",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "d_port",
					Description: `(Required) The destination port of the L4 rule.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Name of the rule. When the ` + "`" + `resource_type` + "`" + ` is ` + "`" + `net` + "`" + `, this field should be set with valid domain.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Protocol of the rule. Valid values: ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + `. When ` + "`" + `source_type` + "`" + ` is 1(host source), the value of this field can only set with ` + "`" + `tcp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required, ForceNew) ID of the resource that the layer 4 rule works for.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `(Required, ForceNew) Type of the resource that the layer 4 rule works for. Valid values: ` + "`" + `bgpip` + "`" + ` and ` + "`" + `net` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "s_port",
					Description: `(Required) The source port of the L4 rule.`,
				},
				resource.Attribute{
					Name:        "source_list",
					Description: `(Required) Source list of the rule, it can be a set of ip sources or a set of domain sources. The number of items ranges from 1 to 20.`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required, ForceNew) Source type, ` + "`" + `1` + "`" + ` for source of host, ` + "`" + `2` + "`" + ` for source of IP.`,
				},
				resource.Attribute{
					Name:        "health_check_health_num",
					Description: `(Optional) Health threshold of health check, and the default is 3. If a success result is returned for the health check 3 consecutive times, indicates that the forwarding is normal. The value range is 2-10.`,
				},
				resource.Attribute{
					Name:        "health_check_interval",
					Description: `(Optional) Interval time of health check. The value range is 10-60 sec, and the default is 15 sec.`,
				},
				resource.Attribute{
					Name:        "health_check_switch",
					Description: `(Optional) Indicates whether health check is enabled. The default is ` + "`" + `false` + "`" + `. Only valid when source list has more than one source item.`,
				},
				resource.Attribute{
					Name:        "health_check_timeout",
					Description: `(Optional) HTTP Status Code. The default is 26 and value range is 2-60.`,
				},
				resource.Attribute{
					Name:        "health_check_unhealth_num",
					Description: `(Optional) Unhealthy threshold of health check, and the default is 3. If the unhealthy result is returned 3 consecutive times, indicates that the forwarding is abnormal. The value range is 2-10.`,
				},
				resource.Attribute{
					Name:        "session_switch",
					Description: `(Optional) Indicate that the session will keep or not, and default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "session_time",
					Description: `(Optional) Session keep time, only valid when ` + "`" + `session_switch` + "`" + ` is true, the available value ranges from 1 to 300 and unit is second. The ` + "`" + `source_list` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Required) Source IP or domain, valid format of ip is like ` + "`" + `1.1.1.1` + "`" + ` and valid format of host source is like ` + "`" + `abc.com` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Required) Weight of the source, the valid value ranges from 0 to 100. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "lb_type",
					Description: `LB type of the rule. Valid values: ` + "`" + `1` + "`" + `, ` + "`" + `2` + "`" + `. ` + "`" + `1` + "`" + ` for weight cycling and ` + "`" + `2` + "`" + ` for IP hash.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `ID of the layer 4 rule.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "lb_type",
					Description: `LB type of the rule. Valid values: ` + "`" + `1` + "`" + `, ` + "`" + `2` + "`" + `. ` + "`" + `1` + "`" + ` for weight cycling and ` + "`" + `2` + "`" + ` for IP hash.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `ID of the layer 4 rule.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dayu_l7_rule",
			Category:         "Anti-DDoS(Dayu)",
			ShortDescription: `Use this resource to create dayu layer 7 rule`,
			Description:      ``,
			Keywords: []string{
				"anti",
				"ddos",
				"dayu",
				"l7",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `(Required, ForceNew) Domain that the layer 7 rule works for. Valid string length ranges from 0 to 80.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Name of the rule.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Protocol of the rule. Valid values: ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required, ForceNew) ID of the resource that the layer 7 rule works for.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `(Required, ForceNew) Type of the resource that the layer 7 rule works for, valid value is ` + "`" + `bgpip` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "source_list",
					Description: `(Required) Source list of the rule, it can be a set of ip sources or a set of domain sources. The number of items ranges from 1 to 16.`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) Source type, ` + "`" + `1` + "`" + ` for source of host, ` + "`" + `2` + "`" + ` for source of IP.`,
				},
				resource.Attribute{
					Name:        "switch",
					Description: `(Required) Indicate the rule will take effect or not.`,
				},
				resource.Attribute{
					Name:        "health_check_code",
					Description: `(Optional) HTTP Status Code. The default is ` + "`" + `26` + "`" + `. Valid value ranges: [1~31]. ` + "`" + `1` + "`" + ` means the return value '1xx' is health. ` + "`" + `2` + "`" + ` means the return value '2xx' is health. ` + "`" + `4` + "`" + ` means the return value '3xx' is health. ` + "`" + `8` + "`" + ` means the return value '4xx' is health. ` + "`" + `16` + "`" + ` means the return value '5xx' is health. If you want multiple return codes to indicate health, need to add the corresponding values.`,
				},
				resource.Attribute{
					Name:        "health_check_health_num",
					Description: `(Optional) Health threshold of health check, and the default is ` + "`" + `3` + "`" + `. If a success result is returned for the health check 3 consecutive times, indicates that the forwarding is normal. The value range is [2-10].`,
				},
				resource.Attribute{
					Name:        "health_check_interval",
					Description: `(Optional) Interval time of health check. Valid value ranges: [10~60]sec. The default is 15 sec.`,
				},
				resource.Attribute{
					Name:        "health_check_method",
					Description: `(Optional) Methods of health check. The default is 'HEAD', the available value are 'HEAD' and 'GET'.`,
				},
				resource.Attribute{
					Name:        "health_check_path",
					Description: `(Optional) Path of health check. The default is ` + "`" + `/` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_check_switch",
					Description: `(Optional) Indicates whether health check is enabled. The default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_check_unhealth_num",
					Description: `(Optional) Unhealthy threshold of health check, and the default is ` + "`" + `3` + "`" + `. If the unhealthy result is returned 3 consecutive times, indicates that the forwarding is abnormal. The value range is [2-10].`,
				},
				resource.Attribute{
					Name:        "ssl_id",
					Description: `(Optional) SSL ID, when the ` + "`" + `protocol` + "`" + ` is ` + "`" + `https` + "`" + `, the field should be set with valid SSL id. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `ID of the layer 7 rule.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the rule. ` + "`" + `0` + "`" + ` for create/modify success, ` + "`" + `2` + "`" + ` for create/modify fail, ` + "`" + `3` + "`" + ` for delete success, ` + "`" + `5` + "`" + ` for delete failed, ` + "`" + `6` + "`" + ` for waiting to be created/modified, ` + "`" + `7` + "`" + ` for waiting to be deleted and 8 for waiting to get SSL ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `ID of the layer 7 rule.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the rule. ` + "`" + `0` + "`" + ` for create/modify success, ` + "`" + `2` + "`" + ` for create/modify fail, ` + "`" + `3` + "`" + ` for delete success, ` + "`" + `5` + "`" + ` for delete failed, ` + "`" + `6` + "`" + ` for waiting to be created/modified, ` + "`" + `7` + "`" + ` for waiting to be deleted and 8 for waiting to get SSL ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dc_gateway",
			Category:         "Direct Connect Gateway(DCG)",
			ShortDescription: `Provides a resource to creating direct connect gateway instance.`,
			Description:      ``,
			Keywords: []string{
				"direct",
				"connect",
				"gateway",
				"dcg",
				"dc",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the DCG.`,
				},
				resource.Attribute{
					Name:        "network_instance_id",
					Description: `(Required, ForceNew) If the ` + "`" + `network_type` + "`" + ` value is ` + "`" + `VPC` + "`" + `, the available value is VPC ID. But when the ` + "`" + `network_type` + "`" + ` value is ` + "`" + `CCN` + "`" + `, the available value is CCN instance ID.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `(Required, ForceNew) Type of associated network. Valid value: ` + "`" + `VPC` + "`" + ` and ` + "`" + `CCN` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "gateway_type",
					Description: `(Optional, ForceNew) Type of the gateway. Valid value: ` + "`" + `NORMAL` + "`" + ` and ` + "`" + `NAT` + "`" + `. Default is ` + "`" + `NORMAL` + "`" + `. NOTES: CCN only supports ` + "`" + `NORMAL` + "`" + ` and a VPC can create two DCGs, the one is NAT type and the other is non-NAT type. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "cnn_route_type",
					Description: `Type of CCN route. Valid value: ` + "`" + `BGP` + "`" + ` and ` + "`" + `STATIC` + "`" + `. The property is available when the DCG type is CCN gateway and BGP enabled.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of resource.`,
				},
				resource.Attribute{
					Name:        "enable_bgp",
					Description: `Indicates whether the BGP is enabled. ## Import Direct connect gateway instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_dc_gateway.instance dcg-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "cnn_route_type",
					Description: `Type of CCN route. Valid value: ` + "`" + `BGP` + "`" + ` and ` + "`" + `STATIC` + "`" + `. The property is available when the DCG type is CCN gateway and BGP enabled.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of resource.`,
				},
				resource.Attribute{
					Name:        "enable_bgp",
					Description: `Indicates whether the BGP is enabled. ## Import Direct connect gateway instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_dc_gateway.instance dcg-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dc_gateway_ccn_route",
			Category:         "Direct Connect Gateway(DCG)",
			ShortDescription: `Provides a resource to creating direct connect gateway route entry.`,
			Description:      ``,
			Keywords: []string{
				"direct",
				"connect",
				"gateway",
				"dcg",
				"dc",
				"ccn",
				"route",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Required, ForceNew) A network address segment of IDC.`,
				},
				resource.Attribute{
					Name:        "dcg_id",
					Description: `(Required, ForceNew) ID of the DCG. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "as_path",
					Description: `As path list of the BGP.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "as_path",
					Description: `As path list of the BGP.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dcx",
			Category:         "Direct Connect(DC)",
			ShortDescription: `Provides a resource to creating dedicated tunnels instances.`,
			Description:      ``,
			Keywords: []string{
				"direct",
				"connect",
				"dc",
				"dcx",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dc_id",
					Description: `(Required, ForceNew) ID of the DC to be queried, application deployment offline.`,
				},
				resource.Attribute{
					Name:        "dcg_id",
					Description: `(Required, ForceNew) ID of the DC Gateway. Currently only new in the console.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the dedicated tunnel.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, ForceNew) ID of the VPC or BMVPC.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Optional, ForceNew) Bandwidth of the DC.`,
				},
				resource.Attribute{
					Name:        "bgp_asn",
					Description: `(Optional, ForceNew) BGP ASN of the user. A required field within BGP.`,
				},
				resource.Attribute{
					Name:        "bgp_auth_key",
					Description: `(Optional, ForceNew) BGP key of the user.`,
				},
				resource.Attribute{
					Name:        "customer_address",
					Description: `(Optional, ForceNew) Interconnect IP of the DC within client.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `(Optional, ForceNew) Type of the network. Valid value: ` + "`" + `VPC` + "`" + `, ` + "`" + `BMVPC` + "`" + ` and ` + "`" + `CCN` + "`" + `. The default value is ` + "`" + `VPC` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "route_filter_prefixes",
					Description: `(Optional, ForceNew) Static route, the network address of the user IDC. It can be modified after setting but cannot be deleted. AN unable field within BGP.`,
				},
				resource.Attribute{
					Name:        "route_type",
					Description: `(Optional, ForceNew) Type of the route, and available values include BGP and STATIC. The default value is ` + "`" + `BGP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tencent_address",
					Description: `(Optional, ForceNew) Interconnect IP of the DC within Tencent.`,
				},
				resource.Attribute{
					Name:        "vlan",
					Description: `(Optional, ForceNew) Vlan of the dedicated tunnels. Valid value ranges: (0~3000). ` + "`" + `0` + "`" + ` means that only one tunnel can be created for the physical connect. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of resource.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the dedicated tunnels. Valid value: ` + "`" + `PENDING` + "`" + `, ` + "`" + `ALLOCATING` + "`" + `, ` + "`" + `ALLOCATED` + "`" + `, ` + "`" + `ALTERING` + "`" + `, ` + "`" + `DELETING` + "`" + `, ` + "`" + `DELETED` + "`" + `, ` + "`" + `COMFIRMING` + "`" + ` and ` + "`" + `REJECTED` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of resource.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the dedicated tunnels. Valid value: ` + "`" + `PENDING` + "`" + `, ` + "`" + `ALLOCATING` + "`" + `, ` + "`" + `ALLOCATED` + "`" + `, ` + "`" + `ALTERING` + "`" + `, ` + "`" + `DELETING` + "`" + `, ` + "`" + `DELETED` + "`" + `, ` + "`" + `COMFIRMING` + "`" + ` and ` + "`" + `REJECTED` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dnat",
			Category:         "Virtual Private Cloud(VPC)",
			ShortDescription: `Provides a resource to create a NAT forwarding.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"dnat",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "elastic_ip",
					Description: `(Required, ForceNew) Network address of the EIP.`,
				},
				resource.Attribute{
					Name:        "elastic_port",
					Description: `(Required, ForceNew) Port of the EIP.`,
				},
				resource.Attribute{
					Name:        "nat_id",
					Description: `(Required, ForceNew) ID of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `(Required, ForceNew) Network address of the backend service.`,
				},
				resource.Attribute{
					Name:        "private_port",
					Description: `(Required, ForceNew) Port of intranet.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required, ForceNew) Type of the network protocol. Valid value: ` + "`" + `TCP` + "`" + ` and ` + "`" + `UDP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, ForceNew) ID of the VPC.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the NAT forward. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import NAT forwarding can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_dnat.foo tcp://vpc-asg3sfa3:nat-1asg3t63@127.15.2.3:8080 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import NAT forwarding can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_dnat.foo tcp://vpc-asg3sfa3:nat-1asg3t63@127.15.2.3:8080 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_eip",
			Category:         "Cloud Virtual Machine(CVM)",
			ShortDescription: `Provides an EIP resource.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"virtual",
				"machine",
				"cvm",
				"eip",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "anycast_zone",
					Description: `(Optional, ForceNew) The zone of anycast. Valid value: ` + "`" + `ANYCAST_ZONE_GLOBAL` + "`" + ` and ` + "`" + `ANYCAST_ZONE_OVERSEAS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "applicable_for_clb",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "internet_charge_type",
					Description: `(Optional, ForceNew) The charge type of eip. Valid value: ` + "`" + `BANDWIDTH_PACKAGE` + "`" + `, ` + "`" + `BANDWIDTH_POSTPAID_BY_HOUR` + "`" + ` and ` + "`" + `TRAFFIC_POSTPAID_BY_HOUR` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "internet_max_bandwidth_out",
					Description: `(Optional) The bandwidth limit of EIP, unit is Mbps.`,
				},
				resource.Attribute{
					Name:        "internet_service_provider",
					Description: `(Optional, ForceNew) Internet service provider of eip. Valid value: ` + "`" + `BGP` + "`" + `, ` + "`" + `CMCC` + "`" + `, ` + "`" + `CTCC` + "`" + ` and ` + "`" + `CUCC` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of eip.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The tags of eip.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional, ForceNew) The type of eip. Valid value: ` + "`" + `EIP` + "`" + ` and ` + "`" + `AnycastEIP` + "`" + `. Default is ` + "`" + `EIP` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The elastic IP address.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The EIP current status. ## Import EIP can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_eip.foo eip-nyvf60va ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The elastic IP address.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The EIP current status. ## Import EIP can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_eip.foo eip-nyvf60va ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_eip_association",
			Category:         "Cloud Virtual Machine(CVM)",
			ShortDescription: `Provides an eip resource associated with other resource like CVM, ENI and CLB.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"virtual",
				"machine",
				"cvm",
				"eip",
				"association",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "eip_id",
					Description: `(Required, ForceNew) The ID of EIP.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Optional, ForceNew) The CVM or CLB instance id going to bind with the EIP. This field is conflict with ` + "`" + `network_interface_id` + "`" + ` and ` + "`" + `private_ip fields` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "network_interface_id",
					Description: `(Optional, ForceNew) Indicates the network interface id like ` + "`" + `eni-xxxxxx` + "`" + `. This field is conflict with ` + "`" + `instance_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `(Optional, ForceNew) Indicates an IP belongs to the ` + "`" + `network_interface_id` + "`" + `. This field is conflict with ` + "`" + `instance_id` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import Eip association can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_eip_association.bar eip-41s6jwy4::ins-34jwj3 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import Eip association can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_eip_association.bar eip-41s6jwy4::ins-34jwj3 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_elasticsearch_instance",
			Category:         "Elasticsearch",
			ShortDescription: `Provides an elasticsearch instance resource.`,
			Description:      ``,
			Keywords: []string{
				"elasticsearch",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required, ForceNew) Availability zone.`,
				},
				resource.Attribute{
					Name:        "node_info_list",
					Description: `(Required) Node information list, which is used to describe the specification information of various types of nodes in the cluster, such as node type, node quantity, node specification, disk type, and disk size.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Password to an instance.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required, ForceNew) The ID of a VPC subnetwork.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) Version of the instance. Valid values are ` + "`" + `5.6.4` + "`" + `, ` + "`" + `6.4.3` + "`" + `, ` + "`" + `6.8.2` + "`" + ` and ` + "`" + `7.5.1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, ForceNew) The ID of a VPC network.`,
				},
				resource.Attribute{
					Name:        "basic_security_type",
					Description: `(Optional) Whether to enable X-Pack security authentication in Basic Edition 6.8 and above. Valid values are ` + "`" + `1` + "`" + ` and ` + "`" + `2` + "`" + `. ` + "`" + `1` + "`" + ` is disabled, ` + "`" + `2` + "`" + ` is enabled, and default value is ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "charge_period",
					Description: `(Optional, ForceNew) The tenancy of the prepaid instance, and uint is month. NOTE: it only works when charge_type is set to ` + "`" + `PREPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `(Optional, ForceNew) The charge type of instance. Valid values are ` + "`" + `PREPAID` + "`" + ` and ` + "`" + `POSTPAID_BY_HOUR` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "deploy_mode",
					Description: `(Optional, ForceNew) Cluster deployment mode. Valid values are ` + "`" + `0` + "`" + ` and ` + "`" + `1` + "`" + `. ` + "`" + `0` + "`" + ` is single-AZ deployment, and ` + "`" + `1` + "`" + ` is multi-AZ deployment. Default value is ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Optional) Name of the instance, which can contain 1 to 50 English letters, Chinese characters, digits, dashes(-), or underscores(_).`,
				},
				resource.Attribute{
					Name:        "license_type",
					Description: `(Optional) License type. Valid values are ` + "`" + `oss` + "`" + `, ` + "`" + `basic` + "`" + ` and ` + "`" + `platinum` + "`" + `. The default value is ` + "`" + `platinum` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "multi_zone_infos",
					Description: `(Optional, ForceNew) Details of AZs in multi-AZ deployment mode (which is required when deploy_mode is ` + "`" + `1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "renew_flag",
					Description: `(Optional, ForceNew) When enabled, the instance will be renew automatically when it reach the end of the prepaid tenancy. Valid values are ` + "`" + `RENEW_FLAG_AUTO` + "`" + ` and ` + "`" + `RENEW_FLAG_MANUAL` + "`" + `. NOTE: it only works when charge_type is set to ` + "`" + `PREPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the instance. For tag limits, please refer to [Use Limits](https://intl.cloud.tencent.com/document/product/651/13354). The ` + "`" + `multi_zone_infos` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required) Availability zone.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) The ID of a VPC subnetwork. The ` + "`" + `node_info_list` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "node_num",
					Description: `(Required) Number of nodes.`,
				},
				resource.Attribute{
					Name:        "node_type",
					Description: `(Required) Node specification, and valid values refer to [document of tencentcloud](https://intl.cloud.tencent.com/document/product/845/18376).`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `(Optional) Node disk size. Unit is GB, and default value is ` + "`" + `100` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `(Optional) Node disk type. Valid values are ` + "`" + `CLOUD_SSD` + "`" + ` and ` + "`" + `CLOUD_PREMIUM` + "`" + `. The default value is ` + "`" + `CLOUD_SSD` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "encrypt",
					Description: `(Optional) Decides to encrypt this disk or not.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Node type. Valid values are ` + "`" + `hotData` + "`" + `, ` + "`" + `warmData` + "`" + ` and ` + "`" + `dedicatedMaster` + "`" + `. The default value is 'hotData` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Instance creation time.`,
				},
				resource.Attribute{
					Name:        "elasticsearch_domain",
					Description: `Elasticsearch domain name.`,
				},
				resource.Attribute{
					Name:        "elasticsearch_port",
					Description: `Elasticsearch port.`,
				},
				resource.Attribute{
					Name:        "elasticsearch_vip",
					Description: `Elasticsearch VIP.`,
				},
				resource.Attribute{
					Name:        "kibana_url",
					Description: `Kibana access URL. ## Import Elasticsearch instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_elasticsearch_instance.foo es-17634f05 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Instance creation time.`,
				},
				resource.Attribute{
					Name:        "elasticsearch_domain",
					Description: `Elasticsearch domain name.`,
				},
				resource.Attribute{
					Name:        "elasticsearch_port",
					Description: `Elasticsearch port.`,
				},
				resource.Attribute{
					Name:        "elasticsearch_vip",
					Description: `Elasticsearch VIP.`,
				},
				resource.Attribute{
					Name:        "kibana_url",
					Description: `Kibana access URL. ## Import Elasticsearch instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_elasticsearch_instance.foo es-17634f05 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_eni",
			Category:         "Virtual Private Cloud(VPC)",
			ShortDescription: `Provides a resource to create an ENI.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"eni",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the ENI, maximum length 60.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required, ForceNew) ID of the subnet within this vpc.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, ForceNew) ID of the vpc.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the ENI, maximum length 60.`,
				},
				resource.Attribute{
					Name:        "ipv4_count",
					Description: `(Optional) The number of intranet IPv4s. When it is greater than 1, there is only one primary intranet IP. The others are auxiliary intranet IPs, which conflict with ` + "`" + `ipv4s` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ipv4s",
					Description: `(Optional) Applying for intranet IPv4s collection, conflict with ` + "`" + `ipv4_count` + "`" + `. When there are multiple ipv4s, can only be one primary IP, and the maximum length of the array is 30. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Optional) A set of security group IDs.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of the ENI. The ` + "`" + `ipv4s` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) Intranet IP.`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `(Required) Indicates whether the IP is primary.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the IP, maximum length 25. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the ENI.`,
				},
				resource.Attribute{
					Name:        "ipv4_info",
					Description: `An information list of IPv4s. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the IP.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Intranet IP.`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `Indicates whether the IP is primary.`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `MAC address.`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `Indicates whether the IP is primary.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the ENI. ## Import ENI can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_eni.foo eni-qka182br ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the ENI.`,
				},
				resource.Attribute{
					Name:        "ipv4_info",
					Description: `An information list of IPv4s. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the IP.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Intranet IP.`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `Indicates whether the IP is primary.`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `MAC address.`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `Indicates whether the IP is primary.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the ENI. ## Import ENI can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_eni.foo eni-qka182br ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_eni_attachment",
			Category:         "Virtual Private Cloud(VPC)",
			ShortDescription: `Provides a resource to detailed information of attached backend server to an ENI.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"eni",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "eni_id",
					Description: `(Required, ForceNew) ID of the ENI.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, ForceNew) ID of the instance which bind the ENI. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import ENI attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_eni_attachment.foo eni-gtlvkjvz+ins-0h3a5new ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import ENI attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_eni_attachment.foo eni-gtlvkjvz+ins-0h3a5new ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_gaap_certificate",
			Category:         "Global Application Acceleration(GAAP)",
			ShortDescription: `Provides a resource to create a certificate of GAAP.`,
			Description:      ``,
			Keywords: []string{
				"global",
				"application",
				"acceleration",
				"gaap",
				"certificate",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "content",
					Description: `(Required, ForceNew) Content of the certificate, and URL encoding. When the certificate is basic authentication, use the ` + "`" + `user:xxx password:xxx` + "`" + ` format, where the password is encrypted with ` + "`" + `htpasswd` + "`" + ` or ` + "`" + `openssl` + "`" + `; When the certificate is ` + "`" + `CA` + "`" + ` or ` + "`" + `SSL` + "`" + `, the format is ` + "`" + `pem` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required, ForceNew) Type of the certificate. Valid value: ` + "`" + `BASIC` + "`" + `, ` + "`" + `CLIENT` + "`" + `, ` + "`" + `SERVER` + "`" + `, ` + "`" + `REALSERVER` + "`" + ` and ` + "`" + `PROXY` + "`" + `. ` + "`" + `BASIC` + "`" + ` means basic certificate; ` + "`" + `CLIENT` + "`" + ` means client CA certificate; ` + "`" + `SERVER` + "`" + ` means server SSL certificate; ` + "`" + `REALSERVER` + "`" + ` means realserver CA certificate; ` + "`" + `PROXY` + "`" + ` means proxy SSL certificate.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional, ForceNew) Key of the ` + "`" + `SSL` + "`" + ` certificate.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the certificate. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "begin_time",
					Description: `Beginning time of the certificate.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the certificate.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `Ending time of the certificate.`,
				},
				resource.Attribute{
					Name:        "issuer_cn",
					Description: `Issuer name of the certificate.`,
				},
				resource.Attribute{
					Name:        "subject_cn",
					Description: `Subject name of the certificate. ## Import GAAP certificate can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_gaap_certificate.foo cert-d5y6ei3b ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "begin_time",
					Description: `Beginning time of the certificate.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the certificate.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `Ending time of the certificate.`,
				},
				resource.Attribute{
					Name:        "issuer_cn",
					Description: `Issuer name of the certificate.`,
				},
				resource.Attribute{
					Name:        "subject_cn",
					Description: `Subject name of the certificate. ## Import GAAP certificate can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_gaap_certificate.foo cert-d5y6ei3b ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_gaap_domain_error_page",
			Category:         "Global Application Acceleration(GAAP)",
			ShortDescription: `Provide a resource to custom error page info for a GAAP HTTP domain.`,
			Description:      ``,
			Keywords: []string{
				"global",
				"application",
				"acceleration",
				"gaap",
				"domain",
				"error",
				"page",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "body",
					Description: `(Required, ForceNew) New response body.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Required, ForceNew) HTTP domain.`,
				},
				resource.Attribute{
					Name:        "error_codes",
					Description: `(Required, ForceNew) Original error codes.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Required, ForceNew) ID of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "clear_headers",
					Description: `(Optional, ForceNew) Response headers to be removed.`,
				},
				resource.Attribute{
					Name:        "new_error_code",
					Description: `(Optional, ForceNew) New error code.`,
				},
				resource.Attribute{
					Name:        "set_headers",
					Description: `(Optional, ForceNew) Response headers to be set. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_gaap_http_domain",
			Category:         "Global Application Acceleration(GAAP)",
			ShortDescription: `Provides a resource to create a forward domain of layer7 listener.`,
			Description:      ``,
			Keywords: []string{
				"global",
				"application",
				"acceleration",
				"gaap",
				"http",
				"domain",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `(Required, ForceNew) Forward domain of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Required, ForceNew) ID of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "basic_auth_id",
					Description: `(Optional) ID of the basic authentication.`,
				},
				resource.Attribute{
					Name:        "basic_auth",
					Description: `(Optional) Indicates whether basic authentication is enable, default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `(Optional) ID of the server certificate, default value is ` + "`" + `default` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "client_certificate_id",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "client_certificate_ids",
					Description: `(Optional) ID list of the poly client certificate.`,
				},
				resource.Attribute{
					Name:        "gaap_auth_id",
					Description: `(Optional) ID of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "gaap_auth",
					Description: `(Optional) Indicates whether SSL certificate authentication is enable, default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "realserver_auth",
					Description: `(Optional) Indicates whether realserver authentication is enable, default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "realserver_certificate_domain",
					Description: `(Optional) CA certificate domain of the realserver.`,
				},
				resource.Attribute{
					Name:        "realserver_certificate_id",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "realserver_certificate_ids",
					Description: `(Optional) CA certificate ID list of the realserver. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import GAAP http domain can be imported using the id, e.g. ->`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import GAAP http domain can be imported using the id, e.g. ->`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_gaap_http_rule",
			Category:         "Global Application Acceleration(GAAP)",
			ShortDescription: `Provides a resource to create a forward rule of layer7 listener.`,
			Description:      ``,
			Keywords: []string{
				"global",
				"application",
				"acceleration",
				"gaap",
				"http",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `(Required, ForceNew) Forward domain of the forward rule.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `(Required) Indicates whether health check is enable.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Required, ForceNew) ID of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) Path of the forward rule. Maximum length is 80.`,
				},
				resource.Attribute{
					Name:        "realserver_type",
					Description: `(Required, ForceNew) Type of the realserver. Valid value: ` + "`" + `IP` + "`" + ` and ` + "`" + `DOMAIN` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "connect_timeout",
					Description: `(Optional) Timeout of the health check response, default value is 2s.`,
				},
				resource.Attribute{
					Name:        "forward_host",
					Description: `(Optional) The default value of requested host which is forwarded to the realserver by the listener is ` + "`" + `default` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_check_method",
					Description: `(Optional) Method of the health check. Valid value: ` + "`" + `GET` + "`" + ` and ` + "`" + `HEAD` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_check_path",
					Description: `(Optional) Path of health check. Maximum length is 80.`,
				},
				resource.Attribute{
					Name:        "health_check_status_codes",
					Description: `(Optional) Return code of confirmed normal. Valid value: ` + "`" + `100` + "`" + `, ` + "`" + `200` + "`" + `, ` + "`" + `300` + "`" + `, ` + "`" + `400` + "`" + ` and ` + "`" + `500` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional) Interval of the health check, default value is 5s.`,
				},
				resource.Attribute{
					Name:        "realservers",
					Description: `(Optional) An information list of GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "scheduler",
					Description: `(Optional) Scheduling policy of the forward rule, default value is ` + "`" + `rr` + "`" + `. Valid value: ` + "`" + `rr` + "`" + `, ` + "`" + `wrr` + "`" + ` and ` + "`" + `lc` + "`" + `. The ` + "`" + `realservers` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) ID of the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) IP of the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) Port of the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) Scheduling weight, default value is ` + "`" + `1` + "`" + `. Valid value ranges: (1~100). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import GAAP http rule can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_gaap_http_rule.foo rule-3bsuu01r ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import GAAP http rule can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_gaap_http_rule.foo rule-3bsuu01r ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_gaap_layer4_listener",
			Category:         "Global Application Acceleration(GAAP)",
			ShortDescription: `Provides a resource to create a layer4 listener of GAAP.`,
			Description:      ``,
			Keywords: []string{
				"global",
				"application",
				"acceleration",
				"gaap",
				"layer4",
				"listener",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the layer4 listener, the maximum length is 30.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required, ForceNew) Port of the layer4 listener.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required, ForceNew) Protocol of the layer4 listener. Valid value: ` + "`" + `TCP` + "`" + ` and ` + "`" + `UDP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "proxy_id",
					Description: `(Required, ForceNew) ID of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "realserver_type",
					Description: `(Required, ForceNew) Type of the realserver. Valid value: ` + "`" + `IP` + "`" + ` and ` + "`" + `DOMAIN` + "`" + `. NOTES: when the ` + "`" + `protocol` + "`" + ` is specified as ` + "`" + `TCP` + "`" + ` and the ` + "`" + `scheduler` + "`" + ` is specified as ` + "`" + `wrr` + "`" + `, the item can only be set to ` + "`" + `IP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "connect_timeout",
					Description: `(Optional) Timeout of the health check response, should less than interval, default value is 2s. NOTES: Only supports listeners of ` + "`" + `TCP` + "`" + ` protocol and require less than ` + "`" + `interval` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `(Optional) Indicates whether health check is enable, default value is ` + "`" + `false` + "`" + `. NOTES: Only supports listeners of ` + "`" + `TCP` + "`" + ` protocol.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional) Interval of the health check, default value is 5s. NOTES: Only supports listeners of ` + "`" + `TCP` + "`" + ` protocol.`,
				},
				resource.Attribute{
					Name:        "realserver_bind_set",
					Description: `(Optional) An information list of GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "scheduler",
					Description: `(Optional) Scheduling policy of the layer4 listener, default value is ` + "`" + `rr` + "`" + `. Valid value: ` + "`" + `rr` + "`" + `, ` + "`" + `wrr` + "`" + ` and ` + "`" + `lc` + "`" + `. The ` + "`" + `realserver_bind_set` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) ID of the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) IP of the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) Port of the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) Scheduling weight, default value is ` + "`" + `1` + "`" + `. The range of values is [1,100]. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the layer4 listener.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the layer4 listener. ## Import GAAP layer4 listener can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_gaap_layer4_listener.foo listener-11112222 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the layer4 listener.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the layer4 listener. ## Import GAAP layer4 listener can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_gaap_layer4_listener.foo listener-11112222 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_gaap_layer7_listener",
			Category:         "Global Application Acceleration(GAAP)",
			ShortDescription: `Provides a resource to create a layer7 listener of GAAP.`,
			Description:      ``,
			Keywords: []string{
				"global",
				"application",
				"acceleration",
				"gaap",
				"layer7",
				"listener",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the layer7 listener, the maximum length is 30.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required, ForceNew) Port of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required, ForceNew) Protocol of the layer7 listener. Valid value: ` + "`" + `HTTP` + "`" + ` and ` + "`" + `HTTPS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "proxy_id",
					Description: `(Required, ForceNew) ID of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "auth_type",
					Description: `(Optional, ForceNew) Authentication type of the layer7 listener. ` + "`" + `0` + "`" + ` is one-way authentication and ` + "`" + `1` + "`" + ` is mutual authentication. NOTES: Only supports listeners of ` + "`" + `HTTPS` + "`" + ` protocol.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `(Optional) Certificate ID of the layer7 listener. NOTES: Only supports listeners of ` + "`" + `HTTPS` + "`" + ` protocol.`,
				},
				resource.Attribute{
					Name:        "client_certificate_id",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "client_certificate_ids",
					Description: `(Optional) ID list of the client certificate. Set only when ` + "`" + `auth_type` + "`" + ` is specified as mutual authentication. NOTES: Only supports listeners of ` + "`" + `HTTPS` + "`" + ` protocol.`,
				},
				resource.Attribute{
					Name:        "forward_protocol",
					Description: `(Optional, ForceNew) Protocol type of the forwarding. Valid value: ` + "`" + `HTTP` + "`" + ` and ` + "`" + `HTTPS` + "`" + `. NOTES: Only supports listeners of ` + "`" + `HTTPS` + "`" + ` protocol. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the layer7 listener. ## Import GAAP layer7 listener can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_gaap_layer7_listener.foo listener-11112222 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the layer7 listener. ## Import GAAP layer7 listener can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_gaap_layer7_listener.foo listener-11112222 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_gaap_proxy",
			Category:         "Global Application Acceleration(GAAP)",
			ShortDescription: `Provides a resource to create a GAAP proxy.`,
			Description:      ``,
			Keywords: []string{
				"global",
				"application",
				"acceleration",
				"gaap",
				"proxy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "access_region",
					Description: `(Required, ForceNew) Access region of the GAAP proxy. Valid value: ` + "`" + `NorthChina` + "`" + `, ` + "`" + `EastChina` + "`" + `, ` + "`" + `SouthChina` + "`" + `, ` + "`" + `SouthwestChina` + "`" + `, ` + "`" + `Hongkong` + "`" + `, ` + "`" + `SL_TAIWAN` + "`" + `, ` + "`" + `SoutheastAsia` + "`" + `, ` + "`" + `Korea` + "`" + `, ` + "`" + `SL_India` + "`" + `, ` + "`" + `SL_Australia` + "`" + `, ` + "`" + `Europe` + "`" + `, ` + "`" + `SL_UK` + "`" + `, ` + "`" + `SL_SouthAmerica` + "`" + `, ` + "`" + `NorthAmerica` + "`" + `, ` + "`" + `SL_MiddleUSA` + "`" + `, ` + "`" + `Canada` + "`" + `, ` + "`" + `SL_VIET` + "`" + `, ` + "`" + `WestIndia` + "`" + `, ` + "`" + `Thailand` + "`" + `, ` + "`" + `Virginia` + "`" + `, ` + "`" + `Russia` + "`" + `, ` + "`" + `Japan` + "`" + ` and ` + "`" + `SL_Indonesia` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Required) Maximum bandwidth of the GAAP proxy, unit is Mbps. Valid value: ` + "`" + `10` + "`" + `, ` + "`" + `20` + "`" + `, ` + "`" + `50` + "`" + `, ` + "`" + `100` + "`" + `, ` + "`" + `200` + "`" + `, ` + "`" + `500` + "`" + ` and ` + "`" + `1000` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "concurrent",
					Description: `(Required) Maximum concurrency of the GAAP proxy, unit is 10k. Valid value: ` + "`" + `2` + "`" + `, ` + "`" + `5` + "`" + `, ` + "`" + `10` + "`" + `, ` + "`" + `20` + "`" + `, ` + "`" + `30` + "`" + `, ` + "`" + `40` + "`" + `, ` + "`" + `50` + "`" + `, ` + "`" + `60` + "`" + `, ` + "`" + `70` + "`" + `, ` + "`" + `80` + "`" + `, ` + "`" + `90` + "`" + ` and ` + "`" + `100` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the GAAP proxy, the maximum length is 30.`,
				},
				resource.Attribute{
					Name:        "realserver_region",
					Description: `(Required, ForceNew) Region of the GAAP realserver. Valid value: ` + "`" + `NorthChina` + "`" + `, ` + "`" + `EastChina` + "`" + `, ` + "`" + `SouthChina` + "`" + `, ` + "`" + `SouthwestChina` + "`" + `, ` + "`" + `Hongkong` + "`" + `, ` + "`" + `SL_TAIWAN` + "`" + `, ` + "`" + `SoutheastAsia` + "`" + `, ` + "`" + `Korea` + "`" + `, ` + "`" + `SL_India` + "`" + `, ` + "`" + `SL_Australia` + "`" + `, ` + "`" + `Europe` + "`" + `, ` + "`" + `SL_UK` + "`" + `, ` + "`" + `SL_SouthAmerica` + "`" + `, ` + "`" + `NorthAmerica` + "`" + `, ` + "`" + `SL_MiddleUSA` + "`" + `, ` + "`" + `Canada` + "`" + `, ` + "`" + `SL_VIET` + "`" + `, ` + "`" + `WestIndia` + "`" + `, ` + "`" + `Thailand` + "`" + `, ` + "`" + `Virginia` + "`" + `, ` + "`" + `Russia` + "`" + `, ` + "`" + `Japan` + "`" + ` and ` + "`" + `SL_Indonesia` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional) Indicates whether GAAP proxy is enabled, default value is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) ID of the project within the GAAP proxy, ` + "`" + `0` + "`" + ` means is default project.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of the GAAP proxy. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Access domain of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "forward_ip",
					Description: `Forwarding IP of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Access IP of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "scalable",
					Description: `Indicates whether GAAP proxy can scalable.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "support_protocols",
					Description: `Supported protocols of the GAAP proxy. ## Import GAAP proxy can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_gaap_proxy.foo link-11112222 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Access domain of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "forward_ip",
					Description: `Forwarding IP of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Access IP of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "scalable",
					Description: `Indicates whether GAAP proxy can scalable.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "support_protocols",
					Description: `Supported protocols of the GAAP proxy. ## Import GAAP proxy can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_gaap_proxy.foo link-11112222 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_gaap_realserver",
			Category:         "Global Application Acceleration(GAAP)",
			ShortDescription: `Provides a resource to create a GAAP realserver.`,
			Description:      ``,
			Keywords: []string{
				"global",
				"application",
				"acceleration",
				"gaap",
				"realserver",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the GAAP realserver, the maximum length is 30.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional, ForceNew) Domain of the GAAP realserver, conflict with ` + "`" + `ip` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Optional, ForceNew) IP of the GAAP realserver, conflict with ` + "`" + `domain` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional, ForceNew) ID of the project within the GAAP realserver, '0' means is default project.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of the GAAP realserver. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import GAAP realserver can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_gaap_realserver.foo rs-4ftghy6 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import GAAP realserver can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_gaap_realserver.foo rs-4ftghy6 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_gaap_security_policy",
			Category:         "Global Application Acceleration(GAAP)",
			ShortDescription: `Provides a resource to create a security policy of GAAP proxy.`,
			Description:      ``,
			Keywords: []string{
				"global",
				"application",
				"acceleration",
				"gaap",
				"security",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "action",
					Description: `(Required, ForceNew) Default policy. Valid value: ` + "`" + `ACCEPT` + "`" + ` and ` + "`" + `DROP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "proxy_id",
					Description: `(Required, ForceNew) ID of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional) Indicates whether policy is enable, default value is ` + "`" + `true` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import GAAP security policy can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_gaap_security_policy.foo pl-xxxx ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import GAAP security policy can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_gaap_security_policy.foo pl-xxxx ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_gaap_security_rule",
			Category:         "Global Application Acceleration(GAAP)",
			ShortDescription: `Provides a resource to create a security policy rule.`,
			Description:      ``,
			Keywords: []string{
				"global",
				"application",
				"acceleration",
				"gaap",
				"security",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "action",
					Description: `(Required, ForceNew) Policy of the rule. Valid value: ` + "`" + `ACCEPT` + "`" + ` and ` + "`" + `DROP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cidr_ip",
					Description: `(Required, ForceNew) A network address block of the request source.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Required, ForceNew) ID of the security policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the security policy rule. Maximum length is 30.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional, ForceNew) Target port. Default value is ` + "`" + `ALL` + "`" + `. Valid examples: ` + "`" + `80` + "`" + `, ` + "`" + `80,443` + "`" + ` and ` + "`" + `3306-20000` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional, ForceNew) Protocol of the security policy rule. Default value is ` + "`" + `ALL` + "`" + `. Valid value: ` + "`" + `TCP` + "`" + `, ` + "`" + `UDP` + "`" + ` and ` + "`" + `ALL` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import GAAP security rule can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_gaap_security_rule.foo sr-xxxxxxxx ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import GAAP security rule can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_gaap_security_rule.foo sr-xxxxxxxx ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_ha_vip",
			Category:         "Virtual Private Cloud(VPC)",
			ShortDescription: `Provides a resource to create a HA VIP.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"ha",
				"vip",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the HA VIP. The length of character is limited to 1-60.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required, ForceNew) Subnet ID.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, ForceNew) VPC ID.`,
				},
				resource.Attribute{
					Name:        "vip",
					Description: `(Optional, ForceNew) Virtual IP address, it must not be occupied and in this VPC network segment. If not set, it will be assigned after resource created automatically. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "address_ip",
					Description: `EIP that is associated.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the HA VIP.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `Instance ID that is associated.`,
				},
				resource.Attribute{
					Name:        "network_interface_id",
					Description: `Network interface ID that is associated.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the HA VIP. Valid value: ` + "`" + `AVAILABLE` + "`" + `, ` + "`" + `UNBIND` + "`" + `. ## Import HA VIP can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_ha_vip.foo havip-kjqwe4ba ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "address_ip",
					Description: `EIP that is associated.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the HA VIP.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `Instance ID that is associated.`,
				},
				resource.Attribute{
					Name:        "network_interface_id",
					Description: `Network interface ID that is associated.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the HA VIP. Valid value: ` + "`" + `AVAILABLE` + "`" + `, ` + "`" + `UNBIND` + "`" + `. ## Import HA VIP can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_ha_vip.foo havip-kjqwe4ba ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_ha_vip_eip_attachment",
			Category:         "Virtual Private Cloud(VPC)",
			ShortDescription: `Provides a resource to create a HA VIP EIP attachment.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"ha",
				"vip",
				"eip",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "address_ip",
					Description: `(Required, ForceNew) Public address of the EIP.`,
				},
				resource.Attribute{
					Name:        "havip_id",
					Description: `(Required, ForceNew) ID of the attached HA VIP. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import HA VIP EIP attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_ha_vip_eip_attachment.foo havip-kjqwe4ba#1.1.1.1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import HA VIP EIP attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_ha_vip_eip_attachment.foo havip-kjqwe4ba#1.1.1.1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_image",
			Category:         "Cloud Virtual Machine(CVM)",
			ShortDescription: `Provide a resource to manage image.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"virtual",
				"machine",
				"cvm",
				"image",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "image_name",
					Description: `(Required) Image name.`,
				},
				resource.Attribute{
					Name:        "data_disk_ids",
					Description: `(Optional, ForceNew) Cloud disk ID list, When creating a whole machine image based on an instance, specify the data disk ID contained in the image.`,
				},
				resource.Attribute{
					Name:        "force_poweroff",
					Description: `(Optional) Set whether to force shutdown during mirroring. The default value is ` + "`" + `false` + "`" + `, when set to true, it means that the mirror will be made after shutdown.`,
				},
				resource.Attribute{
					Name:        "image_description",
					Description: `(Optional) Image Description.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Optional, ForceNew) Cloud server instance ID.`,
				},
				resource.Attribute{
					Name:        "snapshot_ids",
					Description: `(Optional, ForceNew) Cloud disk snapshot ID list; creating a mirror based on a snapshot must include a system disk snapshot. It cannot be passed in simultaneously with InstanceId.`,
				},
				resource.Attribute{
					Name:        "sysprep",
					Description: `(Optional) Sysprep function under Windows. When creating a Windows image, you can select true or false to enable or disable the Syspre function. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import image instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_image.image_snap img-gf7jspk6 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import image instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_image.image_snap img-gf7jspk6 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_instance",
			Category:         "Cloud Virtual Machine(CVM)",
			ShortDescription: `Provides a CVM instance resource.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"virtual",
				"machine",
				"cvm",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required, ForceNew) The available zone for the CVM instance.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Required, ForceNew) The image to use for the instance. Changing ` + "`" + `image_id` + "`" + ` will cause the instance to be destroyed and re-created.`,
				},
				resource.Attribute{
					Name:        "allocate_public_ip",
					Description: `(Optional, ForceNew) Associate a public IP address with an instance in a VPC or Classic. Boolean value, Default is false.`,
				},
				resource.Attribute{
					Name:        "cam_role_name",
					Description: `(Optional, ForceNew) CAM role name authorized to access.`,
				},
				resource.Attribute{
					Name:        "data_disks",
					Description: `(Optional, ForceNew) Settings for data disks.`,
				},
				resource.Attribute{
					Name:        "disable_monitor_service",
					Description: `(Optional) Disable enhance service for monitor, it is enabled by default. When this options is set, monitor agent won't be installed.`,
				},
				resource.Attribute{
					Name:        "disable_security_service",
					Description: `(Optional) Disable enhance service for security, it is enabled by default. When this options is set, security agent won't be installed.`,
				},
				resource.Attribute{
					Name:        "force_delete",
					Description: `(Optional) Indicate whether to force delete the instance. Default is ` + "`" + `false` + "`" + `. If set true, the instance will be permanently deleted instead of being moved into the recycle bin. Note: only works for ` + "`" + `PREPAID` + "`" + ` instance.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional, ForceNew) The hostname of the instance. Windows instance: The name should be a combination of 2 to 15 characters comprised of letters (case insensitive), numbers, and hyphens (-). Period (.) is not supported, and the name cannot be a string of pure numbers. Other types (such as Linux) of instances: The name should be a combination of 2 to 60 characters, supporting multiple periods (.). The piece between two periods is composed of letters (case insensitive), numbers, and hyphens (-).`,
				},
				resource.Attribute{
					Name:        "instance_charge_type_prepaid_period",
					Description: `(Optional) The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when instance_charge_type is set to ` + "`" + `PREPAID` + "`" + `. Valid values are ` + "`" + `1` + "`" + `, ` + "`" + `2` + "`" + `, ` + "`" + `3` + "`" + `, ` + "`" + `4` + "`" + `, ` + "`" + `5` + "`" + `, ` + "`" + `6` + "`" + `, ` + "`" + `7` + "`" + `, ` + "`" + `8` + "`" + `, ` + "`" + `9` + "`" + `, ` + "`" + `10` + "`" + `, ` + "`" + `11` + "`" + `, ` + "`" + `12` + "`" + `, ` + "`" + `24` + "`" + `, ` + "`" + `36` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type_prepaid_renew_flag",
					Description: `(Optional) Auto renewal flag. Valid values: ` + "`" + `NOTIFY_AND_AUTO_RENEW` + "`" + `: notify upon expiration and renew automatically, ` + "`" + `NOTIFY_AND_MANUAL_RENEW` + "`" + `: notify upon expiration but do not renew automatically, ` + "`" + `DISABLE_NOTIFY_AND_MANUAL_RENEW` + "`" + `: neither notify upon expiration nor renew automatically. Default value: ` + "`" + `NOTIFY_AND_MANUAL_RENEW` + "`" + `. If this parameter is specified as ` + "`" + `NOTIFY_AND_AUTO_RENEW` + "`" + `, the instance will be automatically renewed on a monthly basis if the account balance is sufficient. NOTE: it only works when instance_charge_type is set to ` + "`" + `PREPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `(Optional, ForceNew) The charge type of instance. Valid values are ` + "`" + `PREPAID` + "`" + `, ` + "`" + `POSTPAID_BY_HOUR` + "`" + ` and ` + "`" + `SPOTPAID` + "`" + `. The default is ` + "`" + `POSTPAID_BY_HOUR` + "`" + `. Note: TencentCloud International only supports ` + "`" + `POSTPAID_BY_HOUR` + "`" + `. ` + "`" + `PREPAID` + "`" + ` instance may not allow to delete before expired. ` + "`" + `SPOTPAID` + "`" + ` instance must set ` + "`" + `spot_instance_type` + "`" + ` and ` + "`" + `spot_max_price` + "`" + ` at the same time.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Optional) The name of the instance. The max length of instance_name is 60, and default value is ` + "`" + `Terraform-CVM-Instance` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Optional) The type of the instance.`,
				},
				resource.Attribute{
					Name:        "internet_charge_type",
					Description: `(Optional, ForceNew) Internet charge type of the instance, Valid values are ` + "`" + `BANDWIDTH_PREPAID` + "`" + `, ` + "`" + `TRAFFIC_POSTPAID_BY_HOUR` + "`" + `, ` + "`" + `BANDWIDTH_POSTPAID_BY_HOUR` + "`" + ` and ` + "`" + `BANDWIDTH_PACKAGE` + "`" + `. This value does not need to be set when ` + "`" + `allocate_public_ip` + "`" + ` is false.`,
				},
				resource.Attribute{
					Name:        "internet_max_bandwidth_out",
					Description: `(Optional) Maximum outgoing bandwidth to the public network, measured in Mbps (Mega bits per second). This value does not need to be set when ` + "`" + `allocate_public_ip` + "`" + ` is false.`,
				},
				resource.Attribute{
					Name:        "keep_image_login",
					Description: `(Optional, ForceNew) Whether to keep image login or not, default is ` + "`" + `false` + "`" + `. When the image type is private or shared or imported, this parameter can be set ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `(Optional) The key pair to use for the instance, it looks like ` + "`" + `skey-16jig7tx` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Password for the instance. In order for the new password to take effect, the instance will be restarted after the password change.`,
				},
				resource.Attribute{
					Name:        "placement_group_id",
					Description: `(Optional, ForceNew) The ID of a placement group.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `(Optional) The private IP to be assigned to this instance, must be in the provided subnet and available.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) The project the instance belongs to, default to 0.`,
				},
				resource.Attribute{
					Name:        "running_flag",
					Description: `(Optional) Set instance to running or stop. Default value is true, the instance will shutdown when this flag is false.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Optional) A list of security group IDs to associate with.`,
				},
				resource.Attribute{
					Name:        "spot_instance_type",
					Description: `(Optional) Type of spot instance, only support ` + "`" + `ONE-TIME` + "`" + ` now. Note: it only works when instance_charge_type is set to ` + "`" + `SPOTPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "spot_max_price",
					Description: `(Optional, ForceNew) Max price of a spot instance, is the format of decimal string, for example "0.50". Note: it only works when instance_charge_type is set to ` + "`" + `SPOTPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) The ID of a VPC subnet. If you want to create instances in a VPC network, this parameter must be set.`,
				},
				resource.Attribute{
					Name:        "system_disk_id",
					Description: `(Optional) System disk snapshot ID used to initialize the system disk. When system disk type is ` + "`" + `LOCAL_BASIC` + "`" + ` and ` + "`" + `LOCAL_SSD` + "`" + `, disk id is not supported.`,
				},
				resource.Attribute{
					Name:        "system_disk_size",
					Description: `(Optional, ForceNew) Size of the system disk. Valid value ranges: (50~1000). and unit is GB. Default is 50GB.`,
				},
				resource.Attribute{
					Name:        "system_disk_type",
					Description: `(Optional, ForceNew) System disk type. For more information on limits of system disk types, see [Storage Overview](https://intl.cloud.tencent.com/document/product/213/4952). Valid values: ` + "`" + `LOCAL_BASIC` + "`" + `: local disk, ` + "`" + `LOCAL_SSD` + "`" + `: local SSD disk, ` + "`" + `CLOUD_BASIC` + "`" + `: HDD cloud disk, ` + "`" + `CLOUD_SSD` + "`" + `: SSD, ` + "`" + `CLOUD_PREMIUM` + "`" + `: Premium Cloud Storage. NOTE: ` + "`" + `LOCAL_BASIC` + "`" + ` and ` + "`" + `LOCAL_SSD` + "`" + ` are deprecated.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. For tag limits, please refer to [Use Limits](https://intl.cloud.tencent.com/document/product/651/13354).`,
				},
				resource.Attribute{
					Name:        "user_data_raw",
					Description: `(Optional, ForceNew) The user data to be injected into this instance, in plain text. Conflicts with ` + "`" + `user_data` + "`" + `. Up to 16 KB after base64 encoded.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional, ForceNew) The user data to be injected into this instance. Must be base64 encoded and up to 16 KB.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) The ID of a VPC network. If you want to create instances in a VPC network, this parameter must be set. The ` + "`" + `data_disks` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "data_disk_size",
					Description: `(Required, ForceNew) Size of the data disk, and unit is GB. If disk type is ` + "`" + `CLOUD_SSD` + "`" + `, the size range is [100, 16000], and the others are [10-16000].`,
				},
				resource.Attribute{
					Name:        "data_disk_type",
					Description: `(Required, ForceNew) Data disk type. For more information about limits on different data disk types, see [Storage Overview](https://intl.cloud.tencent.com/document/product/213/4952). Valid values: ` + "`" + `LOCAL_BASIC` + "`" + `: local disk, ` + "`" + `LOCAL_SSD` + "`" + `: local SSD disk, ` + "`" + `CLOUD_BASIC` + "`" + `: HDD cloud disk, ` + "`" + `CLOUD_PREMIUM` + "`" + `: Premium Cloud Storage, ` + "`" + `CLOUD_SSD` + "`" + `: SSD, ` + "`" + `CLOUD_HSSD` + "`" + `: Enhanced SSD. NOTE: ` + "`" + `LOCAL_BASIC` + "`" + ` and ` + "`" + `LOCAL_SSD` + "`" + ` are deprecated.`,
				},
				resource.Attribute{
					Name:        "data_disk_id",
					Description: `(Optional) Data disk ID used to initialize the data disk. When data disk type is ` + "`" + `LOCAL_BASIC` + "`" + ` and ` + "`" + `LOCAL_SSD` + "`" + `, disk id is not supported.`,
				},
				resource.Attribute{
					Name:        "data_disk_snapshot_id",
					Description: `(Optional, ForceNew) Snapshot ID of the data disk. The selected data disk snapshot size must be smaller than the data disk size.`,
				},
				resource.Attribute{
					Name:        "delete_with_instance",
					Description: `(Optional, ForceNew) Decides whether the disk is deleted with instance(only applied to ` + "`" + `CLOUD_BASIC` + "`" + `, ` + "`" + `CLOUD_SSD` + "`" + ` and ` + "`" + `CLOUD_PREMIUM` + "`" + ` disk with ` + "`" + `POSTPAID_BY_HOUR` + "`" + ` instance), default is true.`,
				},
				resource.Attribute{
					Name:        "encrypt",
					Description: `(Optional, ForceNew) Decides whether the disk is encrypted. Default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "throughput_performance",
					Description: `(Optional, ForceNew) Add extra performance to the data disk. Only works when disk type is ` + "`" + `CLOUD_TSSD` + "`" + ` or ` + "`" + `CLOUD_HSSD` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the instance.`,
				},
				resource.Attribute{
					Name:        "expired_time",
					Description: `Expired time of the instance.`,
				},
				resource.Attribute{
					Name:        "instance_status",
					Description: `Current status of the instance.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP of the instance. ## Import CVM instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import tencentcloud_instance.foo ins-2qol3a80 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the instance.`,
				},
				resource.Attribute{
					Name:        "expired_time",
					Description: `Expired time of the instance.`,
				},
				resource.Attribute{
					Name:        "instance_status",
					Description: `Current status of the instance.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP of the instance. ## Import CVM instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import tencentcloud_instance.foo ins-2qol3a80 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_key_pair",
			Category:         "Cloud Virtual Machine(CVM)",
			ShortDescription: `Provides a key pair resource.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"virtual",
				"machine",
				"cvm",
				"key",
				"pair",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key_name",
					Description: `(Required) The key pair's name. It is the only in one TencentCloud account.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(Required, ForceNew) You can import an existing public key and using TencentCloud key pair to manage it.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional, ForceNew) Specifys to which project the key pair belongs. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import Key pair can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_key_pair.foo skey-17634f05 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import Key pair can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_key_pair.foo skey-17634f05 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_kms_external_key",
			Category:         "KMS",
			ShortDescription: `Provide a resource to create a KMS external key.`,
			Description:      ``,
			Keywords: []string{
				"kms",
				"external",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "alias",
					Description: `(Required) Name of CMK. The name can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of CMK. The maximum is 1024 bytes.`,
				},
				resource.Attribute{
					Name:        "is_archived",
					Description: `(Optional) Specify whether to archive key. Default value is ` + "`" + `false` + "`" + `. This field is conflict with ` + "`" + `is_enabled` + "`" + `, valid when key_state is ` + "`" + `Enabled` + "`" + `, ` + "`" + `Disabled` + "`" + `, ` + "`" + `Archived` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `(Optional) Specify whether to enable key. Default value is ` + "`" + `false` + "`" + `. This field is conflict with ` + "`" + `is_archived` + "`" + `, valid when key_state is ` + "`" + `Enabled` + "`" + `, ` + "`" + `Disabled` + "`" + `, ` + "`" + `Archived` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "key_material_base64",
					Description: `(Optional) The base64-encoded key material encrypted with the public_key. For regions using the national secret version, the length of the imported key material is required to be 128 bits, and for regions using the FIPS version, the length of the imported key material is required to be 256 bits.`,
				},
				resource.Attribute{
					Name:        "pending_delete_window_in_days",
					Description: `(Optional) Duration in days after which the key is deleted after destruction of the resource, must be between 7 and 30 days. Defaults to 7 days.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of CMK.`,
				},
				resource.Attribute{
					Name:        "valid_to",
					Description: `(Optional) This value means the effective timestamp of the key material, 0 means it does not expire. Need to be greater than the current timestamp, the maximum support is 2147443200.`,
				},
				resource.Attribute{
					Name:        "wrapping_algorithm",
					Description: `(Optional) The algorithm for encrypting key material. Available values include ` + "`" + `RSAES_PKCS1_V1_5` + "`" + `, ` + "`" + `RSAES_OAEP_SHA_1` + "`" + ` and ` + "`" + `RSAES_OAEP_SHA_256` + "`" + `. Default value is ` + "`" + `RSAES_PKCS1_V1_5` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "key_state",
					Description: `State of CMK. ## Import KMS external keys can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_kms_external_key.foo 287e8f40-7cbb-11eb-9a3a-5254004f7f94 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "key_state",
					Description: `State of CMK. ## Import KMS external keys can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_kms_external_key.foo 287e8f40-7cbb-11eb-9a3a-5254004f7f94 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_kms_key",
			Category:         "KMS",
			ShortDescription: `Provide a resource to create a KMS key.`,
			Description:      ``,
			Keywords: []string{
				"kms",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "alias",
					Description: `(Required) Name of CMK. The name can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of CMK. The maximum is 1024 bytes.`,
				},
				resource.Attribute{
					Name:        "is_archived",
					Description: `(Optional) Specify whether to archive key. Default value is ` + "`" + `false` + "`" + `. This field is conflict with ` + "`" + `is_enabled` + "`" + `, valid when key_state is ` + "`" + `Enabled` + "`" + `, ` + "`" + `Disabled` + "`" + `, ` + "`" + `Archived` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `(Optional) Specify whether to enable key. Default value is ` + "`" + `false` + "`" + `. This field is conflict with ` + "`" + `is_archived` + "`" + `, valid when key_state is ` + "`" + `Enabled` + "`" + `, ` + "`" + `Disabled` + "`" + `, ` + "`" + `Archived` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "key_rotation_enabled",
					Description: `(Optional) Specify whether to enable key rotation, valid when key_usage is ` + "`" + `ENCRYPT_DECRYPT` + "`" + `. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "key_usage",
					Description: `(Optional, ForceNew) Usage of CMK. Available values include ` + "`" + `ENCRYPT_DECRYPT` + "`" + `, ` + "`" + `ASYMMETRIC_DECRYPT_RSA_2048` + "`" + `, ` + "`" + `ASYMMETRIC_DECRYPT_SM2` + "`" + `, ` + "`" + `ASYMMETRIC_SIGN_VERIFY_SM2` + "`" + `, ` + "`" + `ASYMMETRIC_SIGN_VERIFY_RSA_2048` + "`" + `, ` + "`" + `ASYMMETRIC_SIGN_VERIFY_ECC` + "`" + `. Default value is ` + "`" + `ENCRYPT_DECRYPT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "pending_delete_window_in_days",
					Description: `(Optional) Duration in days after which the key is deleted after destruction of the resource, must be between 7 and 30 days. Defaults to 7 days.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of CMK. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "key_state",
					Description: `State of CMK. ## Import KMS keys can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_kms_key.foo 287e8f40-7cbb-11eb-9a3a-5254004f7f94 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "key_state",
					Description: `State of CMK. ## Import KMS keys can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_kms_key.foo 287e8f40-7cbb-11eb-9a3a-5254004f7f94 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_kubernetes_as_scaling_group",
			Category:         "Tencent Kubernetes Engine(TKE)",
			ShortDescription: `Provide a resource to create an auto scaling group for kubernetes cluster.`,
			Description:      ``,
			Keywords: []string{
				"tencent",
				"kubernetes",
				"engine",
				"tke",
				"as",
				"scaling",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "auto_scaling_config",
					Description: `(Required, ForceNew) Auto scaling config parameters.`,
				},
				resource.Attribute{
					Name:        "auto_scaling_group",
					Description: `(Required) Auto scaling group parameters.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required, ForceNew) ID of the cluster.`,
				},
				resource.Attribute{
					Name:        "extra_args",
					Description: `(Optional, ForceNew) Custom parameter information related to the node.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional, ForceNew) Labels of kubernetes AS Group created nodes.`,
				},
				resource.Attribute{
					Name:        "unschedulable",
					Description: `(Optional, ForceNew) Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling. The ` + "`" + `auto_scaling_config` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "configuration_name",
					Description: `(Required, ForceNew) Name of a launch configuration.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Required, ForceNew) Specified types of CVM instance.`,
				},
				resource.Attribute{
					Name:        "data_disk",
					Description: `(Optional, ForceNew) Configurations of data disk.`,
				},
				resource.Attribute{
					Name:        "enhanced_monitor_service",
					Description: `(Optional, ForceNew) To specify whether to enable cloud monitor service. Default is TRUE.`,
				},
				resource.Attribute{
					Name:        "enhanced_security_service",
					Description: `(Optional, ForceNew) To specify whether to enable cloud security service. Default is TRUE.`,
				},
				resource.Attribute{
					Name:        "instance_tags",
					Description: `(Optional, ForceNew) A list of tags used to associate different resources.`,
				},
				resource.Attribute{
					Name:        "internet_charge_type",
					Description: `(Optional, ForceNew) Charge types for network traffic. Valid value: ` + "`" + `BANDWIDTH_PREPAID` + "`" + `, ` + "`" + `TRAFFIC_POSTPAID_BY_HOUR` + "`" + `, ` + "`" + `TRAFFIC_POSTPAID_BY_HOUR` + "`" + ` and ` + "`" + `BANDWIDTH_PACKAGE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "internet_max_bandwidth_out",
					Description: `(Optional, ForceNew) Max bandwidth of Internet access in Mbps. Default is ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "key_ids",
					Description: `(Optional, ForceNew) ID list of keys.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional, ForceNew) Password to access.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional, ForceNew) Specifys to which project the configuration belongs.`,
				},
				resource.Attribute{
					Name:        "public_ip_assigned",
					Description: `(Optional, ForceNew) Specify whether to assign an Internet IP address.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Optional, ForceNew) Security groups to which a CVM instance belongs.`,
				},
				resource.Attribute{
					Name:        "system_disk_size",
					Description: `(Optional, ForceNew) Volume of system disk in GB. Default is ` + "`" + `50` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "system_disk_type",
					Description: `(Optional, ForceNew) Type of a CVM disk. Valid value: ` + "`" + `CLOUD_PREMIUM` + "`" + ` and ` + "`" + `CLOUD_SSD` + "`" + `. Default is ` + "`" + `CLOUD_PREMIUM` + "`" + `. The ` + "`" + `auto_scaling_group` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "max_size",
					Description: `(Required) Maximum number of CVM instances (0~2000).`,
				},
				resource.Attribute{
					Name:        "min_size",
					Description: `(Required) Minimum number of CVM instances (0~2000).`,
				},
				resource.Attribute{
					Name:        "scaling_group_name",
					Description: `(Required, ForceNew) Name of a scaling group.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, ForceNew) ID of VPC network.`,
				},
				resource.Attribute{
					Name:        "default_cooldown",
					Description: `(Optional, ForceNew) Default cooldown time in second, and default value is 300.`,
				},
				resource.Attribute{
					Name:        "desired_capacity",
					Description: `(Optional, ForceNew) Desired volume of CVM instances, which is between max_size and min_size.`,
				},
				resource.Attribute{
					Name:        "forward_balancer_ids",
					Description: `(Optional, ForceNew) List of application load balancers, which can't be specified with load_balancer_ids together.`,
				},
				resource.Attribute{
					Name:        "load_balancer_ids",
					Description: `(Optional, ForceNew) ID list of traditional load balancers.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional, ForceNew) Specifys to which project the scaling group belongs.`,
				},
				resource.Attribute{
					Name:        "retry_policy",
					Description: `(Optional, ForceNew) Available values for retry policies include ` + "`" + `IMMEDIATE_RETRY` + "`" + ` and ` + "`" + `INCREMENTAL_INTERVALS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `(Optional, ForceNew) ID list of subnet, and for VPC it is required.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, ForceNew) Tags of a scaling group.`,
				},
				resource.Attribute{
					Name:        "termination_policies",
					Description: `(Optional, ForceNew) Available values for termination policies include ` + "`" + `OLDEST_INSTANCE` + "`" + ` and ` + "`" + `NEWEST_INSTANCE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `(Optional, ForceNew) List of available zones, for Basic network it is required. The ` + "`" + `data_disk` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `(Optional, ForceNew) Volume of disk in GB. Default is ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `(Optional, ForceNew) Types of disk. Valid value: ` + "`" + `CLOUD_PREMIUM` + "`" + ` and ` + "`" + `CLOUD_SSD` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional, ForceNew) Data disk snapshot ID. The ` + "`" + `forward_balancer_ids` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Required, ForceNew) Listener ID for application load balancers.`,
				},
				resource.Attribute{
					Name:        "load_balancer_id",
					Description: `(Required, ForceNew) ID of available load balancers.`,
				},
				resource.Attribute{
					Name:        "target_attribute",
					Description: `(Required, ForceNew) Attribute list of target rules.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `(Optional, ForceNew) ID of forwarding rules. The ` + "`" + `target_attribute` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required, ForceNew) Port number.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Required, ForceNew) Weight. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_kubernetes_cluster",
			Category:         "Tencent Kubernetes Engine(TKE)",
			ShortDescription: `Provide a resource to create a kubernetes cluster.`,
			Description:      ``,
			Keywords: []string{
				"tencent",
				"kubernetes",
				"engine",
				"tke",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, ForceNew) Vpc Id of the cluster.`,
				},
				resource.Attribute{
					Name:        "claim_expired_seconds",
					Description: `(Optional) Claim expired seconds to recycle ENI. This field can only set when field ` + "`" + `network_type` + "`" + ` is 'VPC-CNI'. ` + "`" + `claim_expired_seconds` + "`" + ` must greater or equal than 300 and less than 15768000.`,
				},
				resource.Attribute{
					Name:        "cluster_as_enabled",
					Description: `(Optional, ForceNew) Indicates whether to enable cluster node auto scaler.`,
				},
				resource.Attribute{
					Name:        "cluster_cidr",
					Description: `(Optional, ForceNew) A network address block of the cluster. Different from vpc cidr and cidr of other clusters within this vpc. Must be in 10./192.168/172.[16-31] segments.`,
				},
				resource.Attribute{
					Name:        "cluster_deploy_type",
					Description: `(Optional, ForceNew) Deployment type of the cluster, the available values include: 'MANAGED_CLUSTER' and 'INDEPENDENT_CLUSTER'. Default is 'MANAGED_CLUSTER'.`,
				},
				resource.Attribute{
					Name:        "cluster_desc",
					Description: `(Optional) Description of the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_extra_args",
					Description: `(Optional, ForceNew) Customized parameters for master component,such as kube-apiserver, kube-controller-manager, kube-scheduler.`,
				},
				resource.Attribute{
					Name:        "cluster_internet",
					Description: `(Optional) Open internet access or not.`,
				},
				resource.Attribute{
					Name:        "cluster_intranet_subnet_id",
					Description: `(Optional) Subnet id who can access this independent cluster, this field must and can only set when ` + "`" + `cluster_intranet` + "`" + ` is true. ` + "`" + `cluster_intranet_subnet_id` + "`" + ` can not modify once be set.`,
				},
				resource.Attribute{
					Name:        "cluster_intranet",
					Description: `(Optional) Open intranet access or not.`,
				},
				resource.Attribute{
					Name:        "cluster_ipvs",
					Description: `(Optional, ForceNew) Indicates whether ` + "`" + `ipvs` + "`" + ` is enabled. Default is true. False means ` + "`" + `iptables` + "`" + ` is enabled.`,
				},
				resource.Attribute{
					Name:        "cluster_max_pod_num",
					Description: `(Optional, ForceNew) The maximum number of Pods per node in the cluster. Default is 256. Must be a multiple of 16 and large than 32.`,
				},
				resource.Attribute{
					Name:        "cluster_max_service_num",
					Description: `(Optional, ForceNew) The maximum number of services in the cluster. Default is 256. Must be a multiple of 16.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Optional) Name of the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_os_type",
					Description: `(Optional, ForceNew) Image type of the cluster os, the available values include: 'GENERAL'. Default is 'GENERAL'.`,
				},
				resource.Attribute{
					Name:        "cluster_os",
					Description: `(Optional, ForceNew) Operating system of the cluster, the available values include: 'centos7.2x86_64','centos7.6x86_64','ubuntu16.04.1 LTSx86_64','ubuntu18.04.1 LTSx86_64','tlinux2.4x86_64'. Default is 'ubuntu16.04.1 LTSx86_64'.`,
				},
				resource.Attribute{
					Name:        "cluster_version",
					Description: `(Optional) Version of the cluster, Default is '1.10.5'.`,
				},
				resource.Attribute{
					Name:        "container_runtime",
					Description: `(Optional, ForceNew) Runtime type of the cluster, the available values include: 'docker' and 'containerd'. Default is 'docker'.`,
				},
				resource.Attribute{
					Name:        "deletion_protection",
					Description: `(Optional) Indicates whether cluster deletion protection is enabled. Default is false.`,
				},
				resource.Attribute{
					Name:        "docker_graph_path",
					Description: `(Optional, ForceNew) Docker graph path. Default is ` + "`" + `/var/lib/docker` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "eni_subnet_ids",
					Description: `(Optional) Subnet Ids for cluster with VPC-CNI network mode. This field can only set when field ` + "`" + `network_type` + "`" + ` is 'VPC-CNI'. ` + "`" + `eni_subnet_ids` + "`" + ` can not empty once be set.`,
				},
				resource.Attribute{
					Name:        "extra_args",
					Description: `(Optional, ForceNew) Custom parameter information related to the node.`,
				},
				resource.Attribute{
					Name:        "ignore_cluster_cidr_conflict",
					Description: `(Optional, ForceNew) Indicates whether to ignore the cluster cidr conflict error. Default is false.`,
				},
				resource.Attribute{
					Name:        "is_non_static_ip_mode",
					Description: `(Optional, ForceNew) Indicates whether static ip mode is enabled. Default is false.`,
				},
				resource.Attribute{
					Name:        "kube_proxy_mode",
					Description: `(Optional) Cluster kube-proxy mode, the available values include: 'kube-proxy-bpf'. Default is not set.When set to kube-proxy-bpf, cluster version greater than 1.14 and with Tencent Linux 2.4 is required.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional, ForceNew) Labels of tke cluster nodes.`,
				},
				resource.Attribute{
					Name:        "managed_cluster_internet_security_policies",
					Description: `(Optional) Security policies for managed cluster internet, like:'192.168.1.0/24' or '113.116.51.27', '0.0.0.0/0' means all. This field can only set when field ` + "`" + `cluster_deploy_type` + "`" + ` is 'MANAGED_CLUSTER' and ` + "`" + `cluster_internet` + "`" + ` is true. ` + "`" + `managed_cluster_internet_security_policies` + "`" + ` can not delete or empty once be set.`,
				},
				resource.Attribute{
					Name:        "master_config",
					Description: `(Optional, ForceNew) Deploy the machine configuration information of the 'MASTER_ETCD' service, and create <=7 units for common users.`,
				},
				resource.Attribute{
					Name:        "mount_target",
					Description: `(Optional, ForceNew) Mount target. Default is not mounting.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `(Optional, ForceNew) Cluster network type, GR or VPC-CNI. Default is GR.`,
				},
				resource.Attribute{
					Name:        "node_name_type",
					Description: `(Optional, ForceNew) Node name type of Cluster, the available values include: 'lan-ip' and 'hostname', Default is 'lan-ip'.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) Project ID, default value is 0.`,
				},
				resource.Attribute{
					Name:        "service_cidr",
					Description: `(Optional, ForceNew) A network address block of the service. Different from vpc cidr and cidr of other clusters within this vpc. Must be in 10./192.168/172.[16-31] segments.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The tags of the cluster.`,
				},
				resource.Attribute{
					Name:        "unschedulable",
					Description: `(Optional, ForceNew) Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.`,
				},
				resource.Attribute{
					Name:        "worker_config",
					Description: `(Optional, ForceNew) Deploy the machine configuration information of the 'WORKER' service, and create <=20 units for common users. The other 'WORK' service are added by 'tencentcloud_kubernetes_worker'. The ` + "`" + `cluster_extra_args` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "kube_apiserver",
					Description: `(Optional, ForceNew) The customized parameters for kube-apiserver.`,
				},
				resource.Attribute{
					Name:        "kube_controller_manager",
					Description: `(Optional, ForceNew) The customized parameters for kube-controller-manager.`,
				},
				resource.Attribute{
					Name:        "kube_scheduler",
					Description: `(Optional, ForceNew) The customized parameters for kube-scheduler. The ` + "`" + `data_disk` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `(Optional, ForceNew) Volume of disk in GB. Default is ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `(Optional, ForceNew) Types of disk, available values: ` + "`" + `CLOUD_PREMIUM` + "`" + ` and ` + "`" + `CLOUD_SSD` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional, ForceNew) Data disk snapshot ID. The ` + "`" + `master_config` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Required, ForceNew) Specified types of CVM instance.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required, ForceNew) Private network ID.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional, ForceNew) Indicates which availability zone will be used.`,
				},
				resource.Attribute{
					Name:        "cam_role_name",
					Description: `(Optional, ForceNew) CAM role name authorized to access.`,
				},
				resource.Attribute{
					Name:        "count",
					Description: `(Optional, ForceNew) Number of cvm.`,
				},
				resource.Attribute{
					Name:        "data_disk",
					Description: `(Optional, ForceNew) Configurations of data disk.`,
				},
				resource.Attribute{
					Name:        "enhanced_monitor_service",
					Description: `(Optional, ForceNew) To specify whether to enable cloud monitor service. Default is TRUE.`,
				},
				resource.Attribute{
					Name:        "enhanced_security_service",
					Description: `(Optional, ForceNew) To specify whether to enable cloud security service. Default is TRUE.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional, ForceNew) The host name of the attached instance. Dot (.) and dash (-) cannot be used as the first and last characters of HostName and cannot be used consecutively. Windows example: The length of the name character is [2, 15], letters (capitalization is not restricted), numbers and dashes (-) are allowed, dots (.) are not supported, and not all numbers are allowed. Examples of other types (Linux, etc.): The character length is [2, 60], and multiple dots are allowed. There is a segment between the dots. Each segment allows letters (with no limitation on capitalization), numbers and dashes (-).`,
				},
				resource.Attribute{
					Name:        "instance_charge_type_prepaid_period",
					Description: `(Optional, ForceNew) The tenancy (time unit is month) of the prepaid instance. NOTE: it only works when instance_charge_type is set to ` + "`" + `PREPAID` + "`" + `. Valid values are ` + "`" + `1` + "`" + `, ` + "`" + `2` + "`" + `, ` + "`" + `3` + "`" + `, ` + "`" + `4` + "`" + `, ` + "`" + `5` + "`" + `, ` + "`" + `6` + "`" + `, ` + "`" + `7` + "`" + `, ` + "`" + `8` + "`" + `, ` + "`" + `9` + "`" + `, ` + "`" + `10` + "`" + `, ` + "`" + `11` + "`" + `, ` + "`" + `12` + "`" + `, ` + "`" + `24` + "`" + `, ` + "`" + `36` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type_prepaid_renew_flag",
					Description: `(Optional, ForceNew) Auto renewal flag. Valid values: ` + "`" + `NOTIFY_AND_AUTO_RENEW` + "`" + `: notify upon expiration and renew automatically, ` + "`" + `NOTIFY_AND_MANUAL_RENEW` + "`" + `: notify upon expiration but do not renew automatically, ` + "`" + `DISABLE_NOTIFY_AND_MANUAL_RENEW` + "`" + `: neither notify upon expiration nor renew automatically. Default value: ` + "`" + `NOTIFY_AND_MANUAL_RENEW` + "`" + `. If this parameter is specified as ` + "`" + `NOTIFY_AND_AUTO_RENEW` + "`" + `, the instance will be automatically renewed on a monthly basis if the account balance is sufficient. NOTE: it only works when instance_charge_type is set to ` + "`" + `PREPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `(Optional, ForceNew) The charge type of instance. Valid values are ` + "`" + `PREPAID` + "`" + ` and ` + "`" + `POSTPAID_BY_HOUR` + "`" + `. The default is ` + "`" + `POSTPAID_BY_HOUR` + "`" + `. Note: TencentCloud International only supports ` + "`" + `POSTPAID_BY_HOUR` + "`" + `, ` + "`" + `PREPAID` + "`" + ` instance will not terminated after cluster deleted, and may not allow to delete before expired.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Optional, ForceNew) Name of the CVMs.`,
				},
				resource.Attribute{
					Name:        "internet_charge_type",
					Description: `(Optional, ForceNew) Charge types for network traffic. Available values include ` + "`" + `TRAFFIC_POSTPAID_BY_HOUR` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "internet_max_bandwidth_out",
					Description: `(Optional, ForceNew) Max bandwidth of Internet access in Mbps. Default is 0.`,
				},
				resource.Attribute{
					Name:        "key_ids",
					Description: `(Optional, ForceNew) ID list of keys, should be set if ` + "`" + `password` + "`" + ` not set.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional, ForceNew) Password to access, should be set if ` + "`" + `key_ids` + "`" + ` not set.`,
				},
				resource.Attribute{
					Name:        "public_ip_assigned",
					Description: `(Optional, ForceNew) Specify whether to assign an Internet IP address.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Optional, ForceNew) Security groups to which a CVM instance belongs.`,
				},
				resource.Attribute{
					Name:        "system_disk_size",
					Description: `(Optional, ForceNew) Volume of system disk in GB. Default is ` + "`" + `50` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "system_disk_type",
					Description: `(Optional, ForceNew) System disk type. For more information on limits of system disk types, see [Storage Overview](https://intl.cloud.tencent.com/document/product/213/4952). Valid values: ` + "`" + `LOCAL_BASIC` + "`" + `: local disk, ` + "`" + `LOCAL_SSD` + "`" + `: local SSD disk, ` + "`" + `CLOUD_BASIC` + "`" + `: HDD cloud disk, ` + "`" + `CLOUD_SSD` + "`" + `: SSD, ` + "`" + `CLOUD_PREMIUM` + "`" + `: Premium Cloud Storage. NOTE: ` + "`" + `LOCAL_BASIC` + "`" + ` and ` + "`" + `LOCAL_SSD` + "`" + ` are deprecated.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional, ForceNew) ase64-encoded User Data text, the length limit is 16KB. The ` + "`" + `worker_config` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Required, ForceNew) Specified types of CVM instance.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required, ForceNew) Private network ID.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional, ForceNew) Indicates which availability zone will be used.`,
				},
				resource.Attribute{
					Name:        "cam_role_name",
					Description: `(Optional, ForceNew) CAM role name authorized to access.`,
				},
				resource.Attribute{
					Name:        "count",
					Description: `(Optional, ForceNew) Number of cvm.`,
				},
				resource.Attribute{
					Name:        "data_disk",
					Description: `(Optional, ForceNew) Configurations of data disk.`,
				},
				resource.Attribute{
					Name:        "enhanced_monitor_service",
					Description: `(Optional, ForceNew) To specify whether to enable cloud monitor service. Default is TRUE.`,
				},
				resource.Attribute{
					Name:        "enhanced_security_service",
					Description: `(Optional, ForceNew) To specify whether to enable cloud security service. Default is TRUE.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional, ForceNew) The host name of the attached instance. Dot (.) and dash (-) cannot be used as the first and last characters of HostName and cannot be used consecutively. Windows example: The length of the name character is [2, 15], letters (capitalization is not restricted), numbers and dashes (-) are allowed, dots (.) are not supported, and not all numbers are allowed. Examples of other types (Linux, etc.): The character length is [2, 60], and multiple dots are allowed. There is a segment between the dots. Each segment allows letters (with no limitation on capitalization), numbers and dashes (-).`,
				},
				resource.Attribute{
					Name:        "instance_charge_type_prepaid_period",
					Description: `(Optional, ForceNew) The tenancy (time unit is month) of the prepaid instance. NOTE: it only works when instance_charge_type is set to ` + "`" + `PREPAID` + "`" + `. Valid values are ` + "`" + `1` + "`" + `, ` + "`" + `2` + "`" + `, ` + "`" + `3` + "`" + `, ` + "`" + `4` + "`" + `, ` + "`" + `5` + "`" + `, ` + "`" + `6` + "`" + `, ` + "`" + `7` + "`" + `, ` + "`" + `8` + "`" + `, ` + "`" + `9` + "`" + `, ` + "`" + `10` + "`" + `, ` + "`" + `11` + "`" + `, ` + "`" + `12` + "`" + `, ` + "`" + `24` + "`" + `, ` + "`" + `36` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type_prepaid_renew_flag",
					Description: `(Optional, ForceNew) Auto renewal flag. Valid values: ` + "`" + `NOTIFY_AND_AUTO_RENEW` + "`" + `: notify upon expiration and renew automatically, ` + "`" + `NOTIFY_AND_MANUAL_RENEW` + "`" + `: notify upon expiration but do not renew automatically, ` + "`" + `DISABLE_NOTIFY_AND_MANUAL_RENEW` + "`" + `: neither notify upon expiration nor renew automatically. Default value: ` + "`" + `NOTIFY_AND_MANUAL_RENEW` + "`" + `. If this parameter is specified as ` + "`" + `NOTIFY_AND_AUTO_RENEW` + "`" + `, the instance will be automatically renewed on a monthly basis if the account balance is sufficient. NOTE: it only works when instance_charge_type is set to ` + "`" + `PREPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `(Optional, ForceNew) The charge type of instance. Valid values are ` + "`" + `PREPAID` + "`" + ` and ` + "`" + `POSTPAID_BY_HOUR` + "`" + `. The default is ` + "`" + `POSTPAID_BY_HOUR` + "`" + `. Note: TencentCloud International only supports ` + "`" + `POSTPAID_BY_HOUR` + "`" + `, ` + "`" + `PREPAID` + "`" + ` instance will not terminated after cluster deleted, and may not allow to delete before expired.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Optional, ForceNew) Name of the CVMs.`,
				},
				resource.Attribute{
					Name:        "internet_charge_type",
					Description: `(Optional, ForceNew) Charge types for network traffic. Available values include ` + "`" + `TRAFFIC_POSTPAID_BY_HOUR` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "internet_max_bandwidth_out",
					Description: `(Optional, ForceNew) Max bandwidth of Internet access in Mbps. Default is 0.`,
				},
				resource.Attribute{
					Name:        "key_ids",
					Description: `(Optional, ForceNew) ID list of keys, should be set if ` + "`" + `password` + "`" + ` not set.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional, ForceNew) Password to access, should be set if ` + "`" + `key_ids` + "`" + ` not set.`,
				},
				resource.Attribute{
					Name:        "public_ip_assigned",
					Description: `(Optional, ForceNew) Specify whether to assign an Internet IP address.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Optional, ForceNew) Security groups to which a CVM instance belongs.`,
				},
				resource.Attribute{
					Name:        "system_disk_size",
					Description: `(Optional, ForceNew) Volume of system disk in GB. Default is ` + "`" + `50` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "system_disk_type",
					Description: `(Optional, ForceNew) System disk type. For more information on limits of system disk types, see [Storage Overview](https://intl.cloud.tencent.com/document/product/213/4952). Valid values: ` + "`" + `LOCAL_BASIC` + "`" + `: local disk, ` + "`" + `LOCAL_SSD` + "`" + `: local SSD disk, ` + "`" + `CLOUD_BASIC` + "`" + `: HDD cloud disk, ` + "`" + `CLOUD_SSD` + "`" + `: SSD, ` + "`" + `CLOUD_PREMIUM` + "`" + `: Premium Cloud Storage. NOTE: ` + "`" + `LOCAL_BASIC` + "`" + ` and ` + "`" + `LOCAL_SSD` + "`" + ` are deprecated.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional, ForceNew) ase64-encoded User Data text, the length limit is 16KB. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "certification_authority",
					Description: `The certificate used for access.`,
				},
				resource.Attribute{
					Name:        "cluster_external_endpoint",
					Description: `External network address to access.`,
				},
				resource.Attribute{
					Name:        "cluster_node_num",
					Description: `Number of nodes in the cluster.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Domain name for access.`,
				},
				resource.Attribute{
					Name:        "kube_config",
					Description: `kubernetes config.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Password of account.`,
				},
				resource.Attribute{
					Name:        "pgw_endpoint",
					Description: `The Intranet address used for access.`,
				},
				resource.Attribute{
					Name:        "security_policy",
					Description: `Access policy.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `User name of account.`,
				},
				resource.Attribute{
					Name:        "worker_instances_list",
					Description: `An information list of cvm within the 'WORKER' clusters. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "failed_reason",
					Description: `Information of the cvm when it is failed.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the cvm.`,
				},
				resource.Attribute{
					Name:        "instance_role",
					Description: `Role of the cvm.`,
				},
				resource.Attribute{
					Name:        "instance_state",
					Description: `State of the cvm.`,
				},
				resource.Attribute{
					Name:        "lan_ip",
					Description: `LAN IP of the cvm.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "certification_authority",
					Description: `The certificate used for access.`,
				},
				resource.Attribute{
					Name:        "cluster_external_endpoint",
					Description: `External network address to access.`,
				},
				resource.Attribute{
					Name:        "cluster_node_num",
					Description: `Number of nodes in the cluster.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Domain name for access.`,
				},
				resource.Attribute{
					Name:        "kube_config",
					Description: `kubernetes config.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Password of account.`,
				},
				resource.Attribute{
					Name:        "pgw_endpoint",
					Description: `The Intranet address used for access.`,
				},
				resource.Attribute{
					Name:        "security_policy",
					Description: `Access policy.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `User name of account.`,
				},
				resource.Attribute{
					Name:        "worker_instances_list",
					Description: `An information list of cvm within the 'WORKER' clusters. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "failed_reason",
					Description: `Information of the cvm when it is failed.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the cvm.`,
				},
				resource.Attribute{
					Name:        "instance_role",
					Description: `Role of the cvm.`,
				},
				resource.Attribute{
					Name:        "instance_state",
					Description: `State of the cvm.`,
				},
				resource.Attribute{
					Name:        "lan_ip",
					Description: `LAN IP of the cvm.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_kubernetes_cluster_attachment",
			Category:         "Tencent Kubernetes Engine(TKE)",
			ShortDescription: `Provide a resource to attach an existing cvm to kubernetes cluster.`,
			Description:      ``,
			Keywords: []string{
				"tencent",
				"kubernetes",
				"engine",
				"tke",
				"cluster",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required, ForceNew) ID of the cluster.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, ForceNew) ID of the CVM instance, this cvm will reinstall the system.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional, ForceNew) The host name of the attached instance. Dot (.) and dash (-) cannot be used as the first and last characters of HostName and cannot be used consecutively. Windows example: The length of the name character is [2, 15], letters (capitalization is not restricted), numbers and dashes (-) are allowed, dots (.) are not supported, and not all numbers are allowed. Examples of other types (Linux, etc.): The character length is [2, 60], and multiple dots are allowed. There is a segment between the dots. Each segment allows letters (with no limitation on capitalization), numbers and dashes (-).`,
				},
				resource.Attribute{
					Name:        "key_ids",
					Description: `(Optional, ForceNew) The key pair to use for the instance, it looks like skey-16jig7tx, it should be set if ` + "`" + `password` + "`" + ` not set.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional, ForceNew) Labels of tke attachment exits CVM.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional, ForceNew) Password to access, should be set if ` + "`" + `key_ids` + "`" + ` not set.`,
				},
				resource.Attribute{
					Name:        "unschedulable",
					Description: `(Optional, ForceNew) Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.`,
				},
				resource.Attribute{
					Name:        "worker_config",
					Description: `(Optional, ForceNew) Deploy the machine configuration information of the 'WORKER', commonly used to attach existing instances. The ` + "`" + `data_disk` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "auto_format_and_mount",
					Description: `(Optional, ForceNew) Indicate whether to auto format and mount or not. Default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `(Optional, ForceNew) Volume of disk in GB. Default is ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `(Optional, ForceNew) Types of disk, available values: ` + "`" + `CLOUD_PREMIUM` + "`" + ` and ` + "`" + `CLOUD_SSD` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "file_system",
					Description: `(Optional, ForceNew) File system, e.g. ` + "`" + `ext3/ext4/xfs` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mount_target",
					Description: `(Optional, ForceNew) Mount target. The ` + "`" + `worker_config` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "data_disk",
					Description: `(Optional, ForceNew) Configurations of data disk.`,
				},
				resource.Attribute{
					Name:        "docker_graph_path",
					Description: `(Optional, ForceNew) Docker graph path. Default is ` + "`" + `/var/lib/docker` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "extra_args",
					Description: `(Optional, ForceNew) Custom parameter information related to the node. This is a white-list parameter.`,
				},
				resource.Attribute{
					Name:        "is_schedule",
					Description: `(Optional, ForceNew) Indicate to schedule the adding node or not. Default is true.`,
				},
				resource.Attribute{
					Name:        "mount_target",
					Description: `(Optional, ForceNew) Mount target. Default is not mounting.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional, ForceNew) Base64-encoded User Data text, the length limit is 16KB. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `A list of security group IDs after attach to cluster.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the node.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `A list of security group IDs after attach to cluster.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the node.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_kubernetes_node_pool",
			Category:         "Tencent Kubernetes Engine(TKE)",
			ShortDescription: `Provide a resource to create an auto scaling group for kubernetes cluster.`,
			Description:      ``,
			Keywords: []string{
				"tencent",
				"kubernetes",
				"engine",
				"tke",
				"node",
				"pool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "auto_scaling_config",
					Description: `(Required, ForceNew) Auto scaling config parameters.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required, ForceNew) ID of the cluster.`,
				},
				resource.Attribute{
					Name:        "max_size",
					Description: `(Required) Maximum number of node.`,
				},
				resource.Attribute{
					Name:        "min_size",
					Description: `(Required) Minimum number of node.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Name of the node pool. The name does not exceed 25 characters, and only supports Chinese, English, numbers, underscores, separators (` + "`" + `-` + "`" + `) and decimal points.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, ForceNew) ID of VPC network.`,
				},
				resource.Attribute{
					Name:        "delete_keep_instance",
					Description: `(Optional) Indicate to keep the CVM instance when delete the node pool. Default is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "desired_capacity",
					Description: `(Optional) Desired capacity ot the node. If ` + "`" + `enable_auto_scale` + "`" + ` is set ` + "`" + `true` + "`" + `, this will be a computed parameter.`,
				},
				resource.Attribute{
					Name:        "enable_auto_scale",
					Description: `(Optional) Indicate whether to enable auto scaling or not.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Labels of kubernetes node pool created nodes. The label key name does not exceed 63 characters, only supports English, numbers,'/','-', and does not allow beginning with ('/').`,
				},
				resource.Attribute{
					Name:        "node_config",
					Description: `(Optional) Node config.`,
				},
				resource.Attribute{
					Name:        "node_os_type",
					Description: `(Optional) The image version of the node. Valida values are ` + "`" + `DOCKER_CUSTOMIZE` + "`" + ` and ` + "`" + `GENERAL` + "`" + `. Default is ` + "`" + `GENERAL` + "`" + `. This parameter will only affect new nodes, not including the existing nodes.`,
				},
				resource.Attribute{
					Name:        "node_os",
					Description: `(Optional) Operating system of the cluster, the available values include: ` + "`" + `tlinux2.4x86_64` + "`" + `, ` + "`" + `ubuntu18.04.1x86_64` + "`" + `, ` + "`" + `ubuntu16.04.1 LTSx86_64` + "`" + `, ` + "`" + `centos7.6.0_x64` + "`" + ` and ` + "`" + `centos7.2x86_64` + "`" + `. Default is 'tlinux2.4x86_64'. This parameter will only affect new nodes, not including the existing nodes.`,
				},
				resource.Attribute{
					Name:        "retry_policy",
					Description: `(Optional, ForceNew) Available values for retry policies include ` + "`" + `IMMEDIATE_RETRY` + "`" + ` and ` + "`" + `INCREMENTAL_INTERVALS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scaling_mode",
					Description: `(Optional, ForceNew) Auto scaling mode. Valid values are ` + "`" + `CLASSIC_SCALING` + "`" + `(scaling by create/destroy instances), ` + "`" + `WAKE_UP_STOPPED_SCALING` + "`" + `(Boot priority for expansion. When expanding the capacity, the shutdown operation is given priority to the shutdown of the instance. If the number of instances is still lower than the expected number of instances after the startup, the instance will be created, and the method of destroying the instance will still be used for shrinking).`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `(Optional, ForceNew) ID list of subnet, and for VPC it is required.`,
				},
				resource.Attribute{
					Name:        "taints",
					Description: `(Optional) Taints of kubernetes node pool created nodes.`,
				},
				resource.Attribute{
					Name:        "unschedulable",
					Description: `(Optional, ForceNew) Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling. The ` + "`" + `auto_scaling_config` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Required, ForceNew) Specified types of CVM instance.`,
				},
				resource.Attribute{
					Name:        "data_disk",
					Description: `(Optional, ForceNew) Configurations of data disk.`,
				},
				resource.Attribute{
					Name:        "enhanced_monitor_service",
					Description: `(Optional, ForceNew) To specify whether to enable cloud monitor service. Default is TRUE.`,
				},
				resource.Attribute{
					Name:        "enhanced_security_service",
					Description: `(Optional, ForceNew) To specify whether to enable cloud security service. Default is TRUE.`,
				},
				resource.Attribute{
					Name:        "internet_charge_type",
					Description: `(Optional, ForceNew) Charge types for network traffic. Valid value: ` + "`" + `BANDWIDTH_PREPAID` + "`" + `, ` + "`" + `TRAFFIC_POSTPAID_BY_HOUR` + "`" + `, ` + "`" + `TRAFFIC_POSTPAID_BY_HOUR` + "`" + ` and ` + "`" + `BANDWIDTH_PACKAGE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "internet_max_bandwidth_out",
					Description: `(Optional, ForceNew) Max bandwidth of Internet access in Mbps. Default is ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "key_ids",
					Description: `(Optional, ForceNew) ID list of keys.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional, ForceNew) Password to access.`,
				},
				resource.Attribute{
					Name:        "public_ip_assigned",
					Description: `(Optional, ForceNew) Specify whether to assign an Internet IP address.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Optional, ForceNew) Security groups to which a CVM instance belongs.`,
				},
				resource.Attribute{
					Name:        "system_disk_size",
					Description: `(Optional, ForceNew) Volume of system disk in GB. Default is ` + "`" + `50` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "system_disk_type",
					Description: `(Optional, ForceNew) Type of a CVM disk. Valid value: ` + "`" + `CLOUD_PREMIUM` + "`" + ` and ` + "`" + `CLOUD_SSD` + "`" + `. Default is ` + "`" + `CLOUD_PREMIUM` + "`" + `. The ` + "`" + `data_disk` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "auto_format_and_mount",
					Description: `(Optional, ForceNew) Indicate whether to auto format and mount or not. Default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `(Optional, ForceNew) Volume of disk in GB. Default is ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `(Optional, ForceNew) Types of disk, available values: ` + "`" + `CLOUD_PREMIUM` + "`" + ` and ` + "`" + `CLOUD_SSD` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "file_system",
					Description: `(Optional, ForceNew) File system, e.g. ` + "`" + `ext3/ext4/xfs` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mount_target",
					Description: `(Optional, ForceNew) Mount target. The ` + "`" + `data_disk` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `(Optional, ForceNew) Volume of disk in GB. Default is ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `(Optional, ForceNew) Types of disk. Valid value: ` + "`" + `CLOUD_PREMIUM` + "`" + ` and ` + "`" + `CLOUD_SSD` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional, ForceNew) Data disk snapshot ID. The ` + "`" + `node_config` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "data_disk",
					Description: `(Optional, ForceNew) Configurations of data disk.`,
				},
				resource.Attribute{
					Name:        "docker_graph_path",
					Description: `(Optional, ForceNew) Docker graph path. Default is ` + "`" + `/var/lib/docker` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "extra_args",
					Description: `(Optional, ForceNew) Custom parameter information related to the node. This is a white-list parameter.`,
				},
				resource.Attribute{
					Name:        "is_schedule",
					Description: `(Optional, ForceNew) Indicate to schedule the adding node or not. Default is true.`,
				},
				resource.Attribute{
					Name:        "mount_target",
					Description: `(Optional, ForceNew) Mount target. Default is not mounting.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional, ForceNew) Base64-encoded User Data text, the length limit is 16KB. The ` + "`" + `taints` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "effect",
					Description: `(Required) Effect of the taint. Valid values are: ` + "`" + `NoSchedule` + "`" + `, ` + "`" + `PreferNoSchedule` + "`" + `, ` + "`" + `NoExecute` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Key of the taint. The taint key name does not exceed 63 characters, only supports English, numbers,'/','-', and does not allow beginning with ('/').`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Value of the taint. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "auto_scaling_group_id",
					Description: `The auto scaling group ID.`,
				},
				resource.Attribute{
					Name:        "launch_config_id",
					Description: `The launch config ID.`,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `The total node count.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the node pool.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "auto_scaling_group_id",
					Description: `The auto scaling group ID.`,
				},
				resource.Attribute{
					Name:        "launch_config_id",
					Description: `The launch config ID.`,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `The total node count.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the node pool.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_kubernetes_scale_worker",
			Category:         "Tencent Kubernetes Engine(TKE)",
			ShortDescription: `Provide a resource to increase instance to cluster`,
			Description:      ``,
			Keywords: []string{
				"tencent",
				"kubernetes",
				"engine",
				"tke",
				"scale",
				"worker",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required, ForceNew) ID of the cluster.`,
				},
				resource.Attribute{
					Name:        "worker_config",
					Description: `(Required, ForceNew) Deploy the machine configuration information of the 'WORK' service, and create <=20 units for common users.`,
				},
				resource.Attribute{
					Name:        "data_disk",
					Description: `(Optional, ForceNew) Configurations of data disk.`,
				},
				resource.Attribute{
					Name:        "docker_graph_path",
					Description: `(Optional, ForceNew) Docker graph path. Default is ` + "`" + `/var/lib/docker` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "extra_args",
					Description: `(Optional, ForceNew) Custom parameter information related to the node.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional, ForceNew) Labels of kubernetes scale worker created nodes.`,
				},
				resource.Attribute{
					Name:        "mount_target",
					Description: `(Optional, ForceNew) Mount target. Default is not mounting.`,
				},
				resource.Attribute{
					Name:        "unschedulable",
					Description: `(Optional, ForceNew) Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling. The ` + "`" + `data_disk` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "auto_format_and_mount",
					Description: `(Optional, ForceNew) Indicate whether to auto format and mount or not. Default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `(Optional, ForceNew) Volume of disk in GB. Default is ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `(Optional, ForceNew) Types of disk, available values: ` + "`" + `CLOUD_PREMIUM` + "`" + ` and ` + "`" + `CLOUD_SSD` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "file_system",
					Description: `(Optional, ForceNew) File system, e.g. ` + "`" + `ext3/ext4/xfs` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mount_target",
					Description: `(Optional, ForceNew) Mount target. The ` + "`" + `data_disk` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `(Optional, ForceNew) Volume of disk in GB. Default is ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `(Optional, ForceNew) Types of disk, available values: ` + "`" + `CLOUD_PREMIUM` + "`" + ` and ` + "`" + `CLOUD_SSD` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional, ForceNew) Data disk snapshot ID. The ` + "`" + `worker_config` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Required, ForceNew) Specified types of CVM instance.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required, ForceNew) Private network ID.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional, ForceNew) Indicates which availability zone will be used.`,
				},
				resource.Attribute{
					Name:        "cam_role_name",
					Description: `(Optional, ForceNew) CAM role name authorized to access.`,
				},
				resource.Attribute{
					Name:        "count",
					Description: `(Optional, ForceNew) Number of cvm.`,
				},
				resource.Attribute{
					Name:        "data_disk",
					Description: `(Optional, ForceNew) Configurations of data disk.`,
				},
				resource.Attribute{
					Name:        "enhanced_monitor_service",
					Description: `(Optional, ForceNew) To specify whether to enable cloud monitor service. Default is TRUE.`,
				},
				resource.Attribute{
					Name:        "enhanced_security_service",
					Description: `(Optional, ForceNew) To specify whether to enable cloud security service. Default is TRUE.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional, ForceNew) The host name of the attached instance. Dot (.) and dash (-) cannot be used as the first and last characters of HostName and cannot be used consecutively. Windows example: The length of the name character is [2, 15], letters (capitalization is not restricted), numbers and dashes (-) are allowed, dots (.) are not supported, and not all numbers are allowed. Examples of other types (Linux, etc.): The character length is [2, 60], and multiple dots are allowed. There is a segment between the dots. Each segment allows letters (with no limitation on capitalization), numbers and dashes (-).`,
				},
				resource.Attribute{
					Name:        "instance_charge_type_prepaid_period",
					Description: `(Optional, ForceNew) The tenancy (time unit is month) of the prepaid instance. NOTE: it only works when instance_charge_type is set to ` + "`" + `PREPAID` + "`" + `. Valid values are ` + "`" + `1` + "`" + `, ` + "`" + `2` + "`" + `, ` + "`" + `3` + "`" + `, ` + "`" + `4` + "`" + `, ` + "`" + `5` + "`" + `, ` + "`" + `6` + "`" + `, ` + "`" + `7` + "`" + `, ` + "`" + `8` + "`" + `, ` + "`" + `9` + "`" + `, ` + "`" + `10` + "`" + `, ` + "`" + `11` + "`" + `, ` + "`" + `12` + "`" + `, ` + "`" + `24` + "`" + `, ` + "`" + `36` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type_prepaid_renew_flag",
					Description: `(Optional, ForceNew) Auto renewal flag. Valid values: ` + "`" + `NOTIFY_AND_AUTO_RENEW` + "`" + `: notify upon expiration and renew automatically, ` + "`" + `NOTIFY_AND_MANUAL_RENEW` + "`" + `: notify upon expiration but do not renew automatically, ` + "`" + `DISABLE_NOTIFY_AND_MANUAL_RENEW` + "`" + `: neither notify upon expiration nor renew automatically. Default value: ` + "`" + `NOTIFY_AND_MANUAL_RENEW` + "`" + `. If this parameter is specified as ` + "`" + `NOTIFY_AND_AUTO_RENEW` + "`" + `, the instance will be automatically renewed on a monthly basis if the account balance is sufficient. NOTE: it only works when instance_charge_type is set to ` + "`" + `PREPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `(Optional, ForceNew) The charge type of instance. Valid values are ` + "`" + `PREPAID` + "`" + ` and ` + "`" + `POSTPAID_BY_HOUR` + "`" + `. The default is ` + "`" + `POSTPAID_BY_HOUR` + "`" + `. Note: TencentCloud International only supports ` + "`" + `POSTPAID_BY_HOUR` + "`" + `, ` + "`" + `PREPAID` + "`" + ` instance will not terminated after cluster deleted, and may not allow to delete before expired.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Optional, ForceNew) Name of the CVMs.`,
				},
				resource.Attribute{
					Name:        "internet_charge_type",
					Description: `(Optional, ForceNew) Charge types for network traffic. Available values include ` + "`" + `TRAFFIC_POSTPAID_BY_HOUR` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "internet_max_bandwidth_out",
					Description: `(Optional, ForceNew) Max bandwidth of Internet access in Mbps. Default is 0.`,
				},
				resource.Attribute{
					Name:        "key_ids",
					Description: `(Optional, ForceNew) ID list of keys, should be set if ` + "`" + `password` + "`" + ` not set.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional, ForceNew) Password to access, should be set if ` + "`" + `key_ids` + "`" + ` not set.`,
				},
				resource.Attribute{
					Name:        "public_ip_assigned",
					Description: `(Optional, ForceNew) Specify whether to assign an Internet IP address.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Optional, ForceNew) Security groups to which a CVM instance belongs.`,
				},
				resource.Attribute{
					Name:        "system_disk_size",
					Description: `(Optional, ForceNew) Volume of system disk in GB. Default is ` + "`" + `50` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "system_disk_type",
					Description: `(Optional, ForceNew) System disk type. For more information on limits of system disk types, see [Storage Overview](https://intl.cloud.tencent.com/document/product/213/4952). Valid values: ` + "`" + `LOCAL_BASIC` + "`" + `: local disk, ` + "`" + `LOCAL_SSD` + "`" + `: local SSD disk, ` + "`" + `CLOUD_BASIC` + "`" + `: HDD cloud disk, ` + "`" + `CLOUD_SSD` + "`" + `: SSD, ` + "`" + `CLOUD_PREMIUM` + "`" + `: Premium Cloud Storage. NOTE: ` + "`" + `LOCAL_BASIC` + "`" + ` and ` + "`" + `LOCAL_SSD` + "`" + ` are deprecated.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional, ForceNew) ase64-encoded User Data text, the length limit is 16KB. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "worker_instances_list",
					Description: `An information list of kubernetes cluster 'WORKER'. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "failed_reason",
					Description: `Information of the cvm when it is failed.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the cvm.`,
				},
				resource.Attribute{
					Name:        "instance_role",
					Description: `Role of the cvm.`,
				},
				resource.Attribute{
					Name:        "instance_state",
					Description: `State of the cvm.`,
				},
				resource.Attribute{
					Name:        "lan_ip",
					Description: `LAN IP of the cvm.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "worker_instances_list",
					Description: `An information list of kubernetes cluster 'WORKER'. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "failed_reason",
					Description: `Information of the cvm when it is failed.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the cvm.`,
				},
				resource.Attribute{
					Name:        "instance_role",
					Description: `Role of the cvm.`,
				},
				resource.Attribute{
					Name:        "instance_state",
					Description: `State of the cvm.`,
				},
				resource.Attribute{
					Name:        "lan_ip",
					Description: `LAN IP of the cvm.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_lb",
			Category:         "Cloud Load Balancer(CLB)",
			ShortDescription: `Provides a Load Balancer resource.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"load",
				"balancer",
				"clb",
				"lb",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `(Required, ForceNew) The network type of the LB. Valid value: 'OPEN', 'INTERNAL'.`,
				},
				resource.Attribute{
					Name:        "forward",
					Description: `(Optional, ForceNew) The type of the LB. Valid value: 'CLASSIC', 'APPLICATION'.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the LB.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional, ForceNew) The project id of the LB, unspecified or 0 stands for default project.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional, ForceNew) The VPC ID of the LB, unspecified or 0 stands for CVM basic network. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the LB.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the LB.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_mongodb_instance",
			Category:         "MongoDB",
			ShortDescription: `Provide a resource to create a Mongodb instance.`,
			Description:      ``,
			Keywords: []string{
				"mongodb",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "available_zone",
					Description: `(Required, ForceNew) The available zone of the Mongodb.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `(Required, ForceNew) Version of the Mongodb, and available values include ` + "`" + `MONGO_3_WT` + "`" + ` (represents MongoDB 3.2 WiredTiger Edition), ` + "`" + `MONGO_3_ROCKS` + "`" + ` (represents MongoDB 3.2 RocksDB Edition), ` + "`" + `MONGO_36_WT` + "`" + ` (represents MongoDB 3.6 WiredTiger Edition) and ` + "`" + `MONGO_40_WT` + "`" + ` (represents MongoDB 4.0 WiredTiger Edition).`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Required) Name of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "machine_type",
					Description: `(Required, ForceNew) Type of Mongodb instance, and available values include ` + "`" + `HIO` + "`" + `(or ` + "`" + `GIO` + "`" + ` which will be deprecated, represents high IO) and ` + "`" + `HIO10G` + "`" + `(or ` + "`" + `TGIO` + "`" + ` which will be deprecated, represents 10-gigabit high IO).`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `(Required) Memory size. The minimum value is 2, and unit is GB. Memory and volume must be upgraded or degraded simultaneously.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Password of this Mongodb account.`,
				},
				resource.Attribute{
					Name:        "volume",
					Description: `(Required) Disk size. The minimum value is 25, and unit is GB. Memory and volume must be upgraded or degraded simultaneously.`,
				},
				resource.Attribute{
					Name:        "auto_renew_flag",
					Description: `(Optional) Auto renew flag. Valid values are ` + "`" + `0` + "`" + `(NOTIFY_AND_MANUAL_RENEW), ` + "`" + `1` + "`" + `(NOTIFY_AND_AUTO_RENEW) and ` + "`" + `2` + "`" + `(DISABLE_NOTIFY_AND_MANUAL_RENEW). Default value is ` + "`" + `0` + "`" + `. Note: only works for PREPAID instance. Only supports` + "`" + `0` + "`" + ` and ` + "`" + `1` + "`" + ` for creation.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `(Optional, ForceNew) The charge type of instance. Valid values are ` + "`" + `PREPAID` + "`" + ` and ` + "`" + `POSTPAID_BY_HOUR` + "`" + `. Default value is ` + "`" + `POSTPAID_BY_HOUR` + "`" + `. Note: TencentCloud International only supports ` + "`" + `POSTPAID_BY_HOUR` + "`" + `. Caution that update operation on this field will delete old instances and create new one with new charge type.`,
				},
				resource.Attribute{
					Name:        "prepaid_period",
					Description: `(Optional) The tenancy (time unit is month) of the prepaid instance. Valid values are 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36. NOTE: it only works when charge_type is set to ` + "`" + `PREPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) ID of the project which the instance belongs.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Optional, ForceNew) ID of the security group. NOTE: for instance which ` + "`" + `engine_version` + "`" + ` is ` + "`" + `MONGO_40_WT` + "`" + `, ` + "`" + `security_groups` + "`" + ` is not supported.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional, ForceNew) ID of the subnet within this VPC. The value is required if ` + "`" + `vpc_id` + "`" + ` is set.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The tags of the Mongodb. Key name ` + "`" + `project` + "`" + ` is system reserved and can't be used.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional, ForceNew) ID of the VPC. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "standby_instance_list",
					Description: `List of standby instances' info.`,
				},
				resource.Attribute{
					Name:        "standby_instance_id",
					Description: `Indicates the ID of standby instance.`,
				},
				resource.Attribute{
					Name:        "standby_instance_region",
					Description: `Indicates the region of standby instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the Mongodb instance, and available values include pending initialization(expressed with 0), processing(expressed with 1), running(expressed with 2) and expired(expressed with -2).`,
				},
				resource.Attribute{
					Name:        "vip",
					Description: `IP of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "vport",
					Description: `IP port of the Mongodb instance. ## Import Mongodb instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_mongodb_instance.mongodb cmgo-41s6jwy4 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "standby_instance_list",
					Description: `List of standby instances' info.`,
				},
				resource.Attribute{
					Name:        "standby_instance_id",
					Description: `Indicates the ID of standby instance.`,
				},
				resource.Attribute{
					Name:        "standby_instance_region",
					Description: `Indicates the region of standby instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the Mongodb instance, and available values include pending initialization(expressed with 0), processing(expressed with 1), running(expressed with 2) and expired(expressed with -2).`,
				},
				resource.Attribute{
					Name:        "vip",
					Description: `IP of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "vport",
					Description: `IP port of the Mongodb instance. ## Import Mongodb instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_mongodb_instance.mongodb cmgo-41s6jwy4 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_mongodb_sharding_instance",
			Category:         "MongoDB",
			ShortDescription: `Provide a resource to create a Mongodb sharding instance.`,
			Description:      ``,
			Keywords: []string{
				"mongodb",
				"sharding",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "available_zone",
					Description: `(Required, ForceNew) The available zone of the Mongodb.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `(Required, ForceNew) Version of the Mongodb, and available values include ` + "`" + `MONGO_3_WT` + "`" + ` (represents MongoDB 3.2 WiredTiger Edition), ` + "`" + `MONGO_3_ROCKS` + "`" + ` (represents MongoDB 3.2 RocksDB Edition), ` + "`" + `MONGO_36_WT` + "`" + ` (represents MongoDB 3.6 WiredTiger Edition) and ` + "`" + `MONGO_40_WT` + "`" + ` (represents MongoDB 4.0 WiredTiger Edition).`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Required) Name of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "machine_type",
					Description: `(Required, ForceNew) Type of Mongodb instance, and available values include ` + "`" + `HIO` + "`" + `(or ` + "`" + `GIO` + "`" + ` which will be deprecated, represents high IO) and ` + "`" + `HIO10G` + "`" + `(or ` + "`" + `TGIO` + "`" + ` which will be deprecated, represents 10-gigabit high IO).`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `(Required) Memory size. The minimum value is 2, and unit is GB. Memory and volume must be upgraded or degraded simultaneously.`,
				},
				resource.Attribute{
					Name:        "nodes_per_shard",
					Description: `(Required, ForceNew) Number of nodes per shard, at least 3(one master and two slaves).`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Password of this Mongodb account.`,
				},
				resource.Attribute{
					Name:        "shard_quantity",
					Description: `(Required, ForceNew) Number of sharding.`,
				},
				resource.Attribute{
					Name:        "volume",
					Description: `(Required) Disk size. The minimum value is 25, and unit is GB. Memory and volume must be upgraded or degraded simultaneously.`,
				},
				resource.Attribute{
					Name:        "auto_renew_flag",
					Description: `(Optional) Auto renew flag. Valid values are ` + "`" + `0` + "`" + `(NOTIFY_AND_MANUAL_RENEW), ` + "`" + `1` + "`" + `(NOTIFY_AND_AUTO_RENEW) and ` + "`" + `2` + "`" + `(DISABLE_NOTIFY_AND_MANUAL_RENEW). Default value is ` + "`" + `0` + "`" + `. Note: only works for PREPAID instance. Only supports` + "`" + `0` + "`" + ` and ` + "`" + `1` + "`" + ` for creation.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `(Optional, ForceNew) The charge type of instance. Valid values are ` + "`" + `PREPAID` + "`" + ` and ` + "`" + `POSTPAID_BY_HOUR` + "`" + `. Default value is ` + "`" + `POSTPAID_BY_HOUR` + "`" + `. Note: TencentCloud International only supports ` + "`" + `POSTPAID_BY_HOUR` + "`" + `. Caution that update operation on this field will delete old instances and create new one with new charge type.`,
				},
				resource.Attribute{
					Name:        "prepaid_period",
					Description: `(Optional) The tenancy (time unit is month) of the prepaid instance. Valid values are 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36. NOTE: it only works when charge_type is set to ` + "`" + `PREPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) ID of the project which the instance belongs.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Optional, ForceNew) ID of the security group. NOTE: for instance which ` + "`" + `engine_version` + "`" + ` is ` + "`" + `MONGO_40_WT` + "`" + `, ` + "`" + `security_groups` + "`" + ` is not supported.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional, ForceNew) ID of the subnet within this VPC. The value is required if ` + "`" + `vpc_id` + "`" + ` is set.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The tags of the Mongodb. Key name ` + "`" + `project` + "`" + ` is system reserved and can't be used.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional, ForceNew) ID of the VPC. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the Mongodb instance, and available values include pending initialization(expressed with 0), processing(expressed with 1), running(expressed with 2) and expired(expressed with -2).`,
				},
				resource.Attribute{
					Name:        "vip",
					Description: `IP of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "vport",
					Description: `IP port of the Mongodb instance. ## Import Mongodb sharding instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_mongodb_sharding_instance.mongodb cmgo-41s6jwy4 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the Mongodb instance, and available values include pending initialization(expressed with 0), processing(expressed with 1), running(expressed with 2) and expired(expressed with -2).`,
				},
				resource.Attribute{
					Name:        "vip",
					Description: `IP of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "vport",
					Description: `IP port of the Mongodb instance. ## Import Mongodb sharding instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_mongodb_sharding_instance.mongodb cmgo-41s6jwy4 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_mongodb_standby_instance",
			Category:         "MongoDB",
			ShortDescription: `Provide a resource to create a Mongodb standby instance.`,
			Description:      ``,
			Keywords: []string{
				"mongodb",
				"standby",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "available_zone",
					Description: `(Required, ForceNew) The available zone of the Mongodb standby instance. NOTE: must not be same with father instance's.`,
				},
				resource.Attribute{
					Name:        "father_instance_id",
					Description: `(Required, ForceNew) Indicates the main instance ID of standby instances.`,
				},
				resource.Attribute{
					Name:        "father_instance_region",
					Description: `(Required, ForceNew) Indicates the region of main instance.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Required) Name of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `(Required) Memory size. The minimum value is 2, and unit is GB. Memory and volume must be upgraded or degraded simultaneously.`,
				},
				resource.Attribute{
					Name:        "volume",
					Description: `(Required) Disk size. The minimum value is 25, and unit is GB. Memory and volume must be upgraded or degraded simultaneously.`,
				},
				resource.Attribute{
					Name:        "auto_renew_flag",
					Description: `(Optional) Auto renew flag. Valid values are ` + "`" + `0` + "`" + `(NOTIFY_AND_MANUAL_RENEW), ` + "`" + `1` + "`" + `(NOTIFY_AND_AUTO_RENEW) and ` + "`" + `2` + "`" + `(DISABLE_NOTIFY_AND_MANUAL_RENEW). Default value is ` + "`" + `0` + "`" + `. Note: only works for PREPAID instance. Only supports` + "`" + `0` + "`" + ` and ` + "`" + `1` + "`" + ` for creation.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `(Optional, ForceNew) The charge type of instance. Valid values are ` + "`" + `PREPAID` + "`" + ` and ` + "`" + `POSTPAID_BY_HOUR` + "`" + `. Default value is ` + "`" + `POSTPAID_BY_HOUR` + "`" + `. Note: TencentCloud International only supports ` + "`" + `POSTPAID_BY_HOUR` + "`" + `. Caution that update operation on this field will delete old instances and create new one with new charge type.`,
				},
				resource.Attribute{
					Name:        "prepaid_period",
					Description: `(Optional) The tenancy (time unit is month) of the prepaid instance. Valid values are 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36. NOTE: it only works when charge_type is set to ` + "`" + `PREPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) ID of the project which the instance belongs.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Optional, ForceNew) ID of the security group. NOTE: for instance which ` + "`" + `engine_version` + "`" + ` is ` + "`" + `MONGO_40_WT` + "`" + `, ` + "`" + `security_groups` + "`" + ` is not supported.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional, ForceNew) ID of the subnet within this VPC. The value is required if ` + "`" + `vpc_id` + "`" + ` is set.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The tags of the Mongodb. Key name ` + "`" + `project` + "`" + ` is system reserved and can't be used.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional, ForceNew) ID of the VPC. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `Version of the standby Mongodb instance and must be same as the version of main instance.`,
				},
				resource.Attribute{
					Name:        "machine_type",
					Description: `Type of standby Mongodb instance and must be same as the type of main instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the Mongodb instance, and available values include pending initialization(expressed with 0), processing(expressed with 1), running(expressed with 2) and expired(expressed with -2).`,
				},
				resource.Attribute{
					Name:        "vip",
					Description: `IP of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "vport",
					Description: `IP port of the Mongodb instance. ## Import Mongodb instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_mongodb_standby_instance.mongodb cmgo-41s6jwy4 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `Version of the standby Mongodb instance and must be same as the version of main instance.`,
				},
				resource.Attribute{
					Name:        "machine_type",
					Description: `Type of standby Mongodb instance and must be same as the type of main instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the Mongodb instance, and available values include pending initialization(expressed with 0), processing(expressed with 1), running(expressed with 2) and expired(expressed with -2).`,
				},
				resource.Attribute{
					Name:        "vip",
					Description: `IP of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "vport",
					Description: `IP port of the Mongodb instance. ## Import Mongodb instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_mongodb_standby_instance.mongodb cmgo-41s6jwy4 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_monitor_binding_object",
			Category:         "Monitor",
			ShortDescription: `Provides a resource for bind objects to a policy group resource.`,
			Description:      ``,
			Keywords: []string{
				"monitor",
				"binding",
				"object",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dimensions",
					Description: `(Required, ForceNew) A list objects. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required, ForceNew) Policy group ID for binding objects. The ` + "`" + `dimensions` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "dimensions_json",
					Description: `(Required, ForceNew) Represents a collection of dimensions of an object instance, json format.eg:'{"unInstanceId":"ins-ot3cq4bi"}'. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_monitor_binding_receiver",
			Category:         "Monitor",
			ShortDescription: `Provides a resource for bind receivers to a policy group resource.`,
			Description:      ``,
			Keywords: []string{
				"monitor",
				"binding",
				"receiver",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required, ForceNew) Policy group ID for binding receivers.`,
				},
				resource.Attribute{
					Name:        "receivers",
					Description: `(Optional) A list of receivers(will overwrite the configuration of the server or other resources). Each element contains the following attributes: The ` + "`" + `receivers` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "notify_way",
					Description: `(Required) Method of warning notification.Optional ` + "`" + `CALL` + "`" + `,` + "`" + `EMAIL` + "`" + `,` + "`" + `SITE` + "`" + `,` + "`" + `SMS` + "`" + `,` + "`" + `WECHAT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "receiver_type",
					Description: `(Required) Receive type. Optional ` + "`" + `group` + "`" + `,` + "`" + `user` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `(Optional) End of alarm period. Meaning with ` + "`" + `start_time` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "receive_language",
					Description: `(Optional) Alert sending language. Optional ` + "`" + `en-US` + "`" + `,` + "`" + `zh-CN` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "receiver_group_list",
					Description: `(Optional) Alarm receive group ID list.`,
				},
				resource.Attribute{
					Name:        "receiver_user_list",
					Description: `(Optional) Alarm receiver ID list.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `(Optional) Alarm period start time. Valid value ranges: (0~86399). which removes the date after it is converted to Beijing time as a Unix timestamp, for example 7200 means '10:0:0'. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_monitor_policy_group",
			Category:         "Monitor",
			ShortDescription: `Provides a policy group resource for monitor.`,
			Description:      ``,
			Keywords: []string{
				"monitor",
				"policy",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group_name",
					Description: `(Required) Policy group name, length should between 1 and 20.`,
				},
				resource.Attribute{
					Name:        "policy_view_name",
					Description: `(Required, ForceNew) Policy view name, eg:` + "`" + `cvm_device` + "`" + `,` + "`" + `BANDWIDTHPACKAGE` + "`" + `, refer to ` + "`" + `data.tencentcloud_monitor_policy_conditions(policy_view_name)` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `(Required, ForceNew) Policy group's remark information.`,
				},
				resource.Attribute{
					Name:        "conditions",
					Description: `(Optional) A list of threshold rules. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "event_conditions",
					Description: `(Optional) A list of event rules. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "is_union_rule",
					Description: `(Optional) The and or relation of indicator alarm rule. Valid values: ` + "`" + `0` + "`" + `, ` + "`" + `1` + "`" + `. ` + "`" + `0` + "`" + ` represents or rule (if any rule is met, the alarm will be raised), ` + "`" + `1` + "`" + ` represents and rule (if all rules are met, the alarm will be raised).The default is 0.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional, ForceNew) The project id to which the policy group belongs, default is ` + "`" + `0` + "`" + `. The ` + "`" + `conditions` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "alarm_notify_period",
					Description: `(Required) Alarm sending cycle per second. <0 does not fire, ` + "`" + `0` + "`" + ` only fires once, and >0 fires every triggerTime second.`,
				},
				resource.Attribute{
					Name:        "alarm_notify_type",
					Description: `(Required) Alarm sending convergence type. ` + "`" + `0` + "`" + ` continuous alarm, ` + "`" + `1` + "`" + ` index alarm.`,
				},
				resource.Attribute{
					Name:        "metric_id",
					Description: `(Required) Id of the metric, refer to ` + "`" + `data.tencentcloud_monitor_policy_conditions(metric_id)` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "calc_period",
					Description: `(Optional) Data aggregation cycle (unit of second), if the metric has a default value can not be filled, refer to ` + "`" + `data.tencentcloud_monitor_policy_conditions(period_keys)` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "calc_type",
					Description: `(Optional) Compare type. Valid value ranges: [1~12]. ` + "`" + `1` + "`" + ` means more than, ` + "`" + `2` + "`" + ` means greater than or equal, ` + "`" + `3` + "`" + ` means less than, ` + "`" + `4` + "`" + ` means less than or equal to, ` + "`" + `5` + "`" + ` means equal, ` + "`" + `6` + "`" + ` means not equal, ` + "`" + `7` + "`" + ` means days rose, ` + "`" + `8` + "`" + ` means days fell, ` + "`" + `9` + "`" + ` means weeks rose, ` + "`" + `10` + "`" + ` means weeks fell, ` + "`" + `11` + "`" + ` means period rise, ` + "`" + `12` + "`" + ` means period fell, refer to ` + "`" + `data.tencentcloud_monitor_policy_conditions(calc_type_keys)` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "calc_value",
					Description: `(Optional) Threshold value, refer to ` + "`" + `data.tencentcloud_monitor_policy_conditions(calc_value_`,
				},
				resource.Attribute{
					Name:        "continue_period",
					Description: `(Optional) The rule triggers an alert that lasts for several detection cycles, refer to ` + "`" + `data.tencentcloud_monitor_policy_conditions(period_num_keys)` + "`" + `. The ` + "`" + `event_conditions` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "alarm_notify_period",
					Description: `(Required) Alarm sending cycle per second. <0 does not fire, ` + "`" + `0` + "`" + ` only fires once, and >0 fires every triggerTime second.`,
				},
				resource.Attribute{
					Name:        "alarm_notify_type",
					Description: `(Required) Alarm sending convergence type. ` + "`" + `0` + "`" + ` continuous alarm, ` + "`" + `1` + "`" + ` index alarm.`,
				},
				resource.Attribute{
					Name:        "event_id",
					Description: `(Required) The ID of this event metric, refer to ` + "`" + `data.tencentcloud_monitor_policy_conditions(event_id). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "binding_objects",
					Description: `A list binding objects(list only those in the ` + "`" + `provider.region` + "`" + `). Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "dimensions_json",
					Description: `Represents a collection of dimensions of an object instance, json format.`,
				},
				resource.Attribute{
					Name:        "is_shielded",
					Description: `Whether the object is shielded or not, 0 means unshielded and 1 means shielded.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region where the object is located.`,
				},
				resource.Attribute{
					Name:        "unique_id",
					Description: `Object unique id.`,
				},
				resource.Attribute{
					Name:        "dimension_group",
					Description: `A list of dimensions for this policy group.`,
				},
				resource.Attribute{
					Name:        "last_edit_uin",
					Description: `Recently edited user uin.`,
				},
				resource.Attribute{
					Name:        "receivers",
					Description: `A list of receivers. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `End of alarm period. Meaning with ` + "`" + `start_time` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "need_send_notice",
					Description: `Do need a telephone alarm contact prompt. You don't need ` + "`" + `0` + "`" + `, you need ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "notify_way",
					Description: `Method of warning notification. Valid values: "SMS", "SITE", "EMAIL", "CALL", "WECHAT".`,
				},
				resource.Attribute{
					Name:        "person_interval",
					Description: `Telephone warning to individual interval (seconds).`,
				},
				resource.Attribute{
					Name:        "receive_language",
					Description: `Alert sending language.`,
				},
				resource.Attribute{
					Name:        "receiver_group_list",
					Description: `Alarm receive group ID list.`,
				},
				resource.Attribute{
					Name:        "receiver_type",
					Description: `Receive type. Valid values: group, user. 'group' (receiving group) or 'user' (receiver).`,
				},
				resource.Attribute{
					Name:        "receiver_user_list",
					Description: `Alarm receiver id list.`,
				},
				resource.Attribute{
					Name:        "recover_notify",
					Description: `Restore notification mode. Optional "SMS".`,
				},
				resource.Attribute{
					Name:        "round_interval",
					Description: `Telephone alarm interval per round (seconds).`,
				},
				resource.Attribute{
					Name:        "round_number",
					Description: `Telephone alarm number.`,
				},
				resource.Attribute{
					Name:        "send_for",
					Description: `Telephone warning time. Valid values: "OCCUR","RECOVER".`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `Alarm period start time.Range [0,86400], which removes the date after it is converted to Beijing time as a Unix timestamp, for example 7200 means '10:0:0'.`,
				},
				resource.Attribute{
					Name:        "uid_list",
					Description: `The phone alerts the receiver uid.`,
				},
				resource.Attribute{
					Name:        "support_regions",
					Description: `Support regions this policy group.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `The policy group update time. ## Import Policy group instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_monitor_policy_group.group group-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "binding_objects",
					Description: `A list binding objects(list only those in the ` + "`" + `provider.region` + "`" + `). Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "dimensions_json",
					Description: `Represents a collection of dimensions of an object instance, json format.`,
				},
				resource.Attribute{
					Name:        "is_shielded",
					Description: `Whether the object is shielded or not, 0 means unshielded and 1 means shielded.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region where the object is located.`,
				},
				resource.Attribute{
					Name:        "unique_id",
					Description: `Object unique id.`,
				},
				resource.Attribute{
					Name:        "dimension_group",
					Description: `A list of dimensions for this policy group.`,
				},
				resource.Attribute{
					Name:        "last_edit_uin",
					Description: `Recently edited user uin.`,
				},
				resource.Attribute{
					Name:        "receivers",
					Description: `A list of receivers. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `End of alarm period. Meaning with ` + "`" + `start_time` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "need_send_notice",
					Description: `Do need a telephone alarm contact prompt. You don't need ` + "`" + `0` + "`" + `, you need ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "notify_way",
					Description: `Method of warning notification. Valid values: "SMS", "SITE", "EMAIL", "CALL", "WECHAT".`,
				},
				resource.Attribute{
					Name:        "person_interval",
					Description: `Telephone warning to individual interval (seconds).`,
				},
				resource.Attribute{
					Name:        "receive_language",
					Description: `Alert sending language.`,
				},
				resource.Attribute{
					Name:        "receiver_group_list",
					Description: `Alarm receive group ID list.`,
				},
				resource.Attribute{
					Name:        "receiver_type",
					Description: `Receive type. Valid values: group, user. 'group' (receiving group) or 'user' (receiver).`,
				},
				resource.Attribute{
					Name:        "receiver_user_list",
					Description: `Alarm receiver id list.`,
				},
				resource.Attribute{
					Name:        "recover_notify",
					Description: `Restore notification mode. Optional "SMS".`,
				},
				resource.Attribute{
					Name:        "round_interval",
					Description: `Telephone alarm interval per round (seconds).`,
				},
				resource.Attribute{
					Name:        "round_number",
					Description: `Telephone alarm number.`,
				},
				resource.Attribute{
					Name:        "send_for",
					Description: `Telephone warning time. Valid values: "OCCUR","RECOVER".`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `Alarm period start time.Range [0,86400], which removes the date after it is converted to Beijing time as a Unix timestamp, for example 7200 means '10:0:0'.`,
				},
				resource.Attribute{
					Name:        "uid_list",
					Description: `The phone alerts the receiver uid.`,
				},
				resource.Attribute{
					Name:        "support_regions",
					Description: `Support regions this policy group.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `The policy group update time. ## Import Policy group instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_monitor_policy_group.group group-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_mysql_account",
			Category:         "MySQL",
			ShortDescription: `Provides a MySQL account resource for database management. A MySQL instance supports multiple database account.`,
			Description:      ``,
			Keywords: []string{
				"mysql",
				"account",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "mysql_id",
					Description: `(Required, ForceNew) Instance ID to which the account belongs.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Account name.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Operation password.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Database description.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Optional, ForceNew) Account host, default is ` + "`" + `%` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_mysql_account_privilege",
			Category:         "MySQL",
			ShortDescription: `Provides a mysql account privilege resource to grant different access privilege to different database. A database can be granted by multiple account.`,
			Description:      ``,
			Keywords: []string{
				"mysql",
				"account",
				"privilege",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_name",
					Description: `(Required, ForceNew) Account name.`,
				},
				resource.Attribute{
					Name:        "database_names",
					Description: `(Required) List of specified database name.`,
				},
				resource.Attribute{
					Name:        "mysql_id",
					Description: `(Required, ForceNew) Instance ID.`,
				},
				resource.Attribute{
					Name:        "account_host",
					Description: `(Optional, ForceNew) Account host, default is ` + "`" + `%` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "privileges",
					Description: `(Optional) Database permissions. Valid values: ` + "`" + `SELECT` + "`" + `, ` + "`" + `INSERT` + "`" + `, ` + "`" + `UPDATE` + "`" + `, ` + "`" + `DELETE` + "`" + `, ` + "`" + `CREATE` + "`" + `, ` + "`" + `DROP` + "`" + `, ` + "`" + `REFERENCES` + "`" + `, ` + "`" + `INDEX` + "`" + `, ` + "`" + `ALTER` + "`" + `, ` + "`" + `CREATE TEMPORARY TABLES` + "`" + `, ` + "`" + `LOCK TABLES` + "`" + `, ` + "`" + `EXECUTE` + "`" + `, ` + "`" + `CREATE VIEW` + "`" + `, ` + "`" + `SHOW VIEW` + "`" + `, ` + "`" + `CREATE ROUTINE` + "`" + `, ` + "`" + `ALTER ROUTINE` + "`" + `, ` + "`" + `EVENT` + "`" + ` and ` + "`" + `TRIGGER` + "`" + `` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_mysql_backup_policy",
			Category:         "MySQL",
			ShortDescription: `Provides a mysql policy resource to create a backup policy.`,
			Description:      ``,
			Keywords: []string{
				"mysql",
				"backup",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "mysql_id",
					Description: `(Required, ForceNew) Instance ID to which policies will be applied.`,
				},
				resource.Attribute{
					Name:        "backup_model",
					Description: `(Optional) Backup method. Supported values include: ` + "`" + `physical` + "`" + ` - physical backup.`,
				},
				resource.Attribute{
					Name:        "backup_time",
					Description: `(Optional) Instance backup time, in the format of 'HH:mm-HH:mm'. Time setting interval is four hours. Default to ` + "`" + `02:00-06:00` + "`" + `. The following value can be supported: ` + "`" + `02:00-06:00` + "`" + `, ` + "`" + `06:00-10:00` + "`" + `, ` + "`" + `10:00-14:00` + "`" + `, ` + "`" + `14:00-18:00` + "`" + `, ` + "`" + `18:00-22:00` + "`" + `, and ` + "`" + `22:00-02:00` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "retention_period",
					Description: `(Optional) Instance backup retention days. Valid value ranges: [7~730]. And default value is ` + "`" + `7` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "binlog_period",
					Description: `Retention period for binlog in days.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "binlog_period",
					Description: `Retention period for binlog in days.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_mysql_instance",
			Category:         "MySQL",
			ShortDescription: `Provides a mysql instance resource to create master database instances.`,
			Description:      ``,
			Keywords: []string{
				"mysql",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Required) The name of a mysql instance.`,
				},
				resource.Attribute{
					Name:        "mem_size",
					Description: `(Required) Memory size (in MB).`,
				},
				resource.Attribute{
					Name:        "root_password",
					Description: `(Required) Password of root account. This parameter can be specified when you purchase master instances, but it should be ignored when you purchase read-only instances or disaster recovery instances.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `(Required) Disk size (in GB).`,
				},
				resource.Attribute{
					Name:        "auto_renew_flag",
					Description: `(Optional) Auto renew flag. NOTES: Only supported prepaid instance.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional, ForceNew) Indicates which availability zone will be used.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `(Optional, ForceNew) Pay type of instance. Valid values:` + "`" + `PREPAID` + "`" + `, ` + "`" + `POSTPAID` + "`" + `. Default is ` + "`" + `POSTPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `(Optional, ForceNew) The version number of the database engine to use. Supported versions include 5.5/5.6/5.7/8.0, and default is 5.7.`,
				},
				resource.Attribute{
					Name:        "first_slave_zone",
					Description: `(Optional, ForceNew) Zone information about first slave instance.`,
				},
				resource.Attribute{
					Name:        "force_delete",
					Description: `(Optional) Indicate whether to delete instance directly or not. Default is ` + "`" + `false` + "`" + `. If set true, the instance will be deleted instead of staying recycle bin. Note: only works for ` + "`" + `PREPAID` + "`" + ` instance. When the main mysql instance set true, this para of the readonly mysql instance will not take effect.`,
				},
				resource.Attribute{
					Name:        "internet_service",
					Description: `(Optional) Indicates whether to enable the access to an instance from public network: 0 - No, 1 - Yes.`,
				},
				resource.Attribute{
					Name:        "intranet_port",
					Description: `(Optional) Public access port. Valid value ranges: [1024~65535]. The default value is ` + "`" + `3306` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `(Optional) List of parameters to use.`,
				},
				resource.Attribute{
					Name:        "pay_type",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "prepaid_period",
					Description: `(Optional) Period of instance. NOTES: Only supported prepaid instance.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) Project ID, default value is 0.`,
				},
				resource.Attribute{
					Name:        "second_slave_zone",
					Description: `(Optional, ForceNew) Zone information about second slave instance.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Optional) Security groups to use.`,
				},
				resource.Attribute{
					Name:        "slave_deploy_mode",
					Description: `(Optional, ForceNew) Availability zone deployment method. Available values: 0 - Single availability zone; 1 - Multiple availability zones.`,
				},
				resource.Attribute{
					Name:        "slave_sync_mode",
					Description: `(Optional, ForceNew) Data replication mode. 0 - Async replication; 1 - Semisync replication; 2 - Strongsync replication.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) Private network ID. If ` + "`" + `vpc_id` + "`" + ` is set, this value is required.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Instance tags.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) ID of VPC, which can be modified once every 24 hours and can't be removed. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "gtid",
					Description: `Indicates whether GTID is enable. ` + "`" + `0` + "`" + ` - Not enabled; ` + "`" + `1` + "`" + ` - Enabled.`,
				},
				resource.Attribute{
					Name:        "internet_host",
					Description: `host for public access.`,
				},
				resource.Attribute{
					Name:        "internet_port",
					Description: `Access port for public access.`,
				},
				resource.Attribute{
					Name:        "intranet_ip",
					Description: `instance intranet IP.`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `Indicates whether the instance is locked. Valid values: ` + "`" + `0` + "`" + `, ` + "`" + `1` + "`" + `. ` + "`" + `0` + "`" + ` - No; ` + "`" + `1` + "`" + ` - Yes.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Instance status. Valid values: ` + "`" + `0` + "`" + `, ` + "`" + `1` + "`" + `, ` + "`" + `4` + "`" + `, ` + "`" + `5` + "`" + `. ` + "`" + `0` + "`" + ` - Creating; ` + "`" + `1` + "`" + ` - Running; ` + "`" + `4` + "`" + ` - Isolating; ` + "`" + `5` + "`" + ` - Isolated.`,
				},
				resource.Attribute{
					Name:        "task_status",
					Description: `Indicates which kind of operations is being executed.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "gtid",
					Description: `Indicates whether GTID is enable. ` + "`" + `0` + "`" + ` - Not enabled; ` + "`" + `1` + "`" + ` - Enabled.`,
				},
				resource.Attribute{
					Name:        "internet_host",
					Description: `host for public access.`,
				},
				resource.Attribute{
					Name:        "internet_port",
					Description: `Access port for public access.`,
				},
				resource.Attribute{
					Name:        "intranet_ip",
					Description: `instance intranet IP.`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `Indicates whether the instance is locked. Valid values: ` + "`" + `0` + "`" + `, ` + "`" + `1` + "`" + `. ` + "`" + `0` + "`" + ` - No; ` + "`" + `1` + "`" + ` - Yes.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Instance status. Valid values: ` + "`" + `0` + "`" + `, ` + "`" + `1` + "`" + `, ` + "`" + `4` + "`" + `, ` + "`" + `5` + "`" + `. ` + "`" + `0` + "`" + ` - Creating; ` + "`" + `1` + "`" + ` - Running; ` + "`" + `4` + "`" + ` - Isolating; ` + "`" + `5` + "`" + ` - Isolated.`,
				},
				resource.Attribute{
					Name:        "task_status",
					Description: `Indicates which kind of operations is being executed.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_mysql_privilege",
			Category:         "MySQL",
			ShortDescription: `Provides a mysql account privilege resource to grant different access privilege to different database. A database can be granted by multiple account.`,
			Description:      ``,
			Keywords: []string{
				"mysql",
				"privilege",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_name",
					Description: `(Required, ForceNew) Account name.the forbidden value is:root,mysql.sys,tencentroot.`,
				},
				resource.Attribute{
					Name:        "global",
					Description: `(Required) Global privileges. available values for Privileges:SELECT,INSERT,UPDATE,DELETE,CREATE,PROCESS,DROP,REFERENCES,INDEX,ALTER,SHOW DATABASES,CREATE TEMPORARY TABLES,LOCK TABLES,EXECUTE,CREATE VIEW,SHOW VIEW,CREATE ROUTINE,ALTER ROUTINE,EVENT,TRIGGER.`,
				},
				resource.Attribute{
					Name:        "mysql_id",
					Description: `(Required, ForceNew) Instance ID.`,
				},
				resource.Attribute{
					Name:        "account_host",
					Description: `(Optional, ForceNew) Account host, default is ` + "`" + `%` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "column",
					Description: `(Optional) Column privileges list.`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Optional) Database privileges list.`,
				},
				resource.Attribute{
					Name:        "table",
					Description: `(Optional) Table privileges list. The ` + "`" + `column` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "column_name",
					Description: `(Required) Column name.`,
				},
				resource.Attribute{
					Name:        "database_name",
					Description: `(Required) Database name.`,
				},
				resource.Attribute{
					Name:        "privileges",
					Description: `(Required) Column privilege.available values for Privileges:SELECT,INSERT,UPDATE,REFERENCES.`,
				},
				resource.Attribute{
					Name:        "table_name",
					Description: `(Required) Table name. The ` + "`" + `database` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "database_name",
					Description: `(Required) Database name.`,
				},
				resource.Attribute{
					Name:        "privileges",
					Description: `(Required) Database privilege.available values for Privileges:SELECT,INSERT,UPDATE,DELETE,CREATE,DROP,REFERENCES,INDEX,ALTER,CREATE TEMPORARY TABLES,LOCK TABLES,EXECUTE,CREATE VIEW,SHOW VIEW,CREATE ROUTINE,ALTER ROUTINE,EVENT,TRIGGER. The ` + "`" + `table` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "database_name",
					Description: `(Required) Database name.`,
				},
				resource.Attribute{
					Name:        "privileges",
					Description: `(Required) Table privilege.available values for Privileges:SELECT,INSERT,UPDATE,DELETE,CREATE,DROP,REFERENCES,INDEX,ALTER,CREATE VIEW,SHOW VIEW,TRIGGER.`,
				},
				resource.Attribute{
					Name:        "table_name",
					Description: `(Required) Table name. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_mysql_readonly_instance",
			Category:         "MySQL",
			ShortDescription: `Provides a mysql instance resource to create read-only database instances.`,
			Description:      ``,
			Keywords: []string{
				"mysql",
				"readonly",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Required) The name of a mysql instance.`,
				},
				resource.Attribute{
					Name:        "master_instance_id",
					Description: `(Required, ForceNew) Indicates the master instance ID of recovery instances.`,
				},
				resource.Attribute{
					Name:        "mem_size",
					Description: `(Required) Memory size (in MB).`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `(Required) Disk size (in GB).`,
				},
				resource.Attribute{
					Name:        "auto_renew_flag",
					Description: `(Optional) Auto renew flag. NOTES: Only supported prepaid instance.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `(Optional, ForceNew) Pay type of instance. Valid values:` + "`" + `PREPAID` + "`" + `, ` + "`" + `POSTPAID` + "`" + `. Default is ` + "`" + `POSTPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "force_delete",
					Description: `(Optional) Indicate whether to delete instance directly or not. Default is ` + "`" + `false` + "`" + `. If set true, the instance will be deleted instead of staying recycle bin. Note: only works for ` + "`" + `PREPAID` + "`" + ` instance. When the main mysql instance set true, this para of the readonly mysql instance will not take effect.`,
				},
				resource.Attribute{
					Name:        "intranet_port",
					Description: `(Optional) Public access port. Valid value ranges: [1024~65535]. The default value is ` + "`" + `3306` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "pay_type",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "prepaid_period",
					Description: `(Optional) Period of instance. NOTES: Only supported prepaid instance.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Optional) Security groups to use.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) Private network ID. If ` + "`" + `vpc_id` + "`" + ` is set, this value is required.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Instance tags.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) ID of VPC, which can be modified once every 24 hours and can't be removed. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "intranet_ip",
					Description: `instance intranet IP.`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `Indicates whether the instance is locked. Valid values: ` + "`" + `0` + "`" + `, ` + "`" + `1` + "`" + `. ` + "`" + `0` + "`" + ` - No; ` + "`" + `1` + "`" + ` - Yes.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Instance status. Valid values: ` + "`" + `0` + "`" + `, ` + "`" + `1` + "`" + `, ` + "`" + `4` + "`" + `, ` + "`" + `5` + "`" + `. ` + "`" + `0` + "`" + ` - Creating; ` + "`" + `1` + "`" + ` - Running; ` + "`" + `4` + "`" + ` - Isolating; ` + "`" + `5` + "`" + ` - Isolated.`,
				},
				resource.Attribute{
					Name:        "task_status",
					Description: `Indicates which kind of operations is being executed.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "intranet_ip",
					Description: `instance intranet IP.`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `Indicates whether the instance is locked. Valid values: ` + "`" + `0` + "`" + `, ` + "`" + `1` + "`" + `. ` + "`" + `0` + "`" + ` - No; ` + "`" + `1` + "`" + ` - Yes.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Instance status. Valid values: ` + "`" + `0` + "`" + `, ` + "`" + `1` + "`" + `, ` + "`" + `4` + "`" + `, ` + "`" + `5` + "`" + `. ` + "`" + `0` + "`" + ` - Creating; ` + "`" + `1` + "`" + ` - Running; ` + "`" + `4` + "`" + ` - Isolating; ` + "`" + `5` + "`" + ` - Isolated.`,
				},
				resource.Attribute{
					Name:        "task_status",
					Description: `Indicates which kind of operations is being executed.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_nat_gateway",
			Category:         "Virtual Private Cloud(VPC)",
			ShortDescription: `Provides a resource to create a NAT gateway.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"nat",
				"gateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "assigned_eip_set",
					Description: `(Required) EIP IP address set bound to the gateway. The value of at least 1 and at most 10.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, ForceNew) ID of the vpc.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Optional) The maximum public network output bandwidth of NAT gateway (unit: Mbps). Valid values: ` + "`" + `20` + "`" + `, ` + "`" + `50` + "`" + `, ` + "`" + `100` + "`" + `, ` + "`" + `200` + "`" + `, ` + "`" + `500` + "`" + `, ` + "`" + `1000` + "`" + `, ` + "`" + `2000` + "`" + `, ` + "`" + `5000` + "`" + `. Default is 100.`,
				},
				resource.Attribute{
					Name:        "max_concurrent",
					Description: `(Optional) The upper limit of concurrent connection of NAT gateway. Valid values: ` + "`" + `1000000` + "`" + `, ` + "`" + `3000000` + "`" + `, ` + "`" + `10000000` + "`" + `. Default is ` + "`" + `1000000` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The available tags within this NAT gateway. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "created_time",
					Description: `Create time of the NAT gateway. ## Import NAT gateway can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_nat_gateway.foo nat-1asg3t63 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "created_time",
					Description: `Create time of the NAT gateway. ## Import NAT gateway can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_nat_gateway.foo nat-1asg3t63 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_placement_group",
			Category:         "Cloud Virtual Machine(CVM)",
			ShortDescription: `Provide a resource to create a placement group.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"virtual",
				"machine",
				"cvm",
				"placement",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the placement group, 1-60 characters in length.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required, ForceNew) Type of the placement group. Valid values: ` + "`" + `HOST` + "`" + `, ` + "`" + `SW` + "`" + ` and ` + "`" + `RACK` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the placement group.`,
				},
				resource.Attribute{
					Name:        "current_num",
					Description: `Number of hosts in the placement group.`,
				},
				resource.Attribute{
					Name:        "cvm_quota_total",
					Description: `Maximum number of hosts in the placement group. ## Import Placement group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_placement_group.foo ps-ilan8vjf ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the placement group.`,
				},
				resource.Attribute{
					Name:        "current_num",
					Description: `Number of hosts in the placement group.`,
				},
				resource.Attribute{
					Name:        "cvm_quota_total",
					Description: `Maximum number of hosts in the placement group. ## Import Placement group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_placement_group.foo ps-ilan8vjf ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_postgresql_instance",
			Category:         "PostgreSQL",
			ShortDescription: `Use this resource to create postgresql instance.`,
			Description:      ``,
			Keywords: []string{
				"postgresql",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "memory",
					Description: `(Required) Memory size(in GB). Allowed value must be larger than ` + "`" + `memory` + "`" + ` that data source ` + "`" + `tencentcloud_postgresql_specinfos` + "`" + ` provides.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the postgresql instance.`,
				},
				resource.Attribute{
					Name:        "root_password",
					Description: `(Required) Password of root account. This parameter can be specified when you purchase master instances, but it should be ignored when you purchase read-only instances or disaster recovery instances.`,
				},
				resource.Attribute{
					Name:        "storage",
					Description: `(Required) Volume size(in GB). Allowed value must be a multiple of 10. The storage must be set with the limit of ` + "`" + `storage_min` + "`" + ` and ` + "`" + `storage_max` + "`" + ` which data source ` + "`" + `tencentcloud_postgresql_specinfos` + "`" + ` provides.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional, ForceNew) Availability zone.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `(Optional, ForceNew) Pay type of the postgresql instance. For now, only ` + "`" + `POSTPAID_BY_HOUR` + "`" + ` is valid.`,
				},
				resource.Attribute{
					Name:        "charset",
					Description: `(Optional, ForceNew) Charset of the root account. Valid values are ` + "`" + `UTF8` + "`" + `,` + "`" + `LATIN1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `(Optional, ForceNew) Version of the postgresql database engine. Valid values: ` + "`" + `9.3.5` + "`" + `, ` + "`" + `9.5.4` + "`" + `, ` + "`" + `10.4` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) Project id, default value is ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "public_access_switch",
					Description: `(Optional) Indicates whether to enable the access to an instance from public network or not.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional, ForceNew) ID of subnet.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The available tags within this postgresql.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional, ForceNew) ID of VPC. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the postgresql instance.`,
				},
				resource.Attribute{
					Name:        "private_access_ip",
					Description: `IP for private access.`,
				},
				resource.Attribute{
					Name:        "private_access_port",
					Description: `Port for private access.`,
				},
				resource.Attribute{
					Name:        "public_access_host",
					Description: `Host for public access.`,
				},
				resource.Attribute{
					Name:        "public_access_port",
					Description: `Port for public access. ## Import postgresql instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_postgresql_instance.foo postgres-cda1iex1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the postgresql instance.`,
				},
				resource.Attribute{
					Name:        "private_access_ip",
					Description: `IP for private access.`,
				},
				resource.Attribute{
					Name:        "private_access_port",
					Description: `Port for private access.`,
				},
				resource.Attribute{
					Name:        "public_access_host",
					Description: `Host for public access.`,
				},
				resource.Attribute{
					Name:        "public_access_port",
					Description: `Port for public access. ## Import postgresql instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_postgresql_instance.foo postgres-cda1iex1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_protocol_template",
			Category:         "Virtual Private Cloud(VPC)",
			ShortDescription: `Provides a resource to manage protocol template.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"protocol",
				"template",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Name of the protocol template.`,
				},
				resource.Attribute{
					Name:        "protocols",
					Description: `(Required) Protocol list. Valid protocols are ` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `, ` + "`" + `gre` + "`" + `. Single port(tcp:80), multi-port(tcp:80,443), port range(tcp:3306-20000), all(tcp:all) format are support. Protocol ` + "`" + `icmp` + "`" + ` and ` + "`" + `gre` + "`" + ` cannot specify port. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import Protocol template can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_protocol_template.foo ppm-nwrggd14 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import Protocol template can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_protocol_template.foo ppm-nwrggd14 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_protocol_template_group",
			Category:         "Virtual Private Cloud(VPC)",
			ShortDescription: `Provides a resource to manage protocol template group.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"protocol",
				"template",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Name of the protocol template group.`,
				},
				resource.Attribute{
					Name:        "template_ids",
					Description: `(Required) Service template ID list. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import Protocol template group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_protocol_template_group.foo ppmg-0np3u974 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import Protocol template group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_protocol_template_group.foo ppmg-0np3u974 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_redis_backup_config",
			Category:         "Redis",
			ShortDescription: `Use this resource to create a backup config of redis.`,
			Description:      ``,
			Keywords: []string{
				"redis",
				"backup",
				"config",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "backup_period",
					Description: `(Required) Specifys which day the backup action should take place. Valid values: ` + "`" + `Monday` + "`" + `, ` + "`" + `Tuesday` + "`" + `, ` + "`" + `Wednesday` + "`" + `, ` + "`" + `Thursday` + "`" + `, ` + "`" + `Friday` + "`" + `, ` + "`" + `Saturday` + "`" + ` and ` + "`" + `Sunday` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "backup_time",
					Description: `(Required) Specifys what time the backup action should take place. And the time interval should be one hour.`,
				},
				resource.Attribute{
					Name:        "redis_id",
					Description: `(Required, ForceNew) ID of a redis instance to which the policy will be applied. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import Redis backup config can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_redis_backup_config.redisconfig redis-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import Redis backup config can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_redis_backup_config.redisconfig redis-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_redis_instance",
			Category:         "Redis",
			ShortDescription: `Provides a resource to create a Redis instance and set its attributes.`,
			Description:      ``,
			Keywords: []string{
				"redis",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required, ForceNew) The available zone ID of an instance to be created, please refer to ` + "`" + `tencentcloud_redis_zone_config.list` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mem_size",
					Description: `(Required) The memory volume of an available instance(in MB), please refer to ` + "`" + `tencentcloud_redis_zone_config.list[zone].mem_sizes` + "`" + `. When redis is standard type, it represents total memory size of the instance; when Redis is cluster type, it represents memory size of per sharding.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Password for a Redis user, which should be 8 to 16 characters.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `(Optional, ForceNew) The charge type of instance. Valid values: ` + "`" + `PREPAID` + "`" + ` and ` + "`" + `POSTPAID` + "`" + `. Default value is ` + "`" + `POSTPAID` + "`" + `. Note: TencentCloud International only supports ` + "`" + `POSTPAID` + "`" + `. Caution that update operation on this field will delete old instances and create new with new charge type.`,
				},
				resource.Attribute{
					Name:        "force_delete",
					Description: `(Optional) Indicate whether to delete Redis instance directly or not. Default is false. If set true, the instance will be deleted instead of staying recycle bin. Note: only works for ` + "`" + `PREPAID` + "`" + ` instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Instance name.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional, ForceNew) The port used to access a redis instance. The default value is 6379. And this value can't be changed after creation, or the Redis instance will be recreated.`,
				},
				resource.Attribute{
					Name:        "prepaid_period",
					Description: `(Optional) The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when charge_type is set to ` + "`" + `PREPAID` + "`" + `. Valid values are ` + "`" + `1` + "`" + `, ` + "`" + `2` + "`" + `, ` + "`" + `3` + "`" + `, ` + "`" + `4` + "`" + `, ` + "`" + `5` + "`" + `, ` + "`" + `6` + "`" + `, ` + "`" + `7` + "`" + `, ` + "`" + `8` + "`" + `, ` + "`" + `9` + "`" + `, ` + "`" + `10` + "`" + `, ` + "`" + `11` + "`" + `, ` + "`" + `12` + "`" + `, ` + "`" + `24` + "`" + `, ` + "`" + `36` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) Specifies which project the instance should belong to.`,
				},
				resource.Attribute{
					Name:        "redis_replicas_num",
					Description: `(Optional, ForceNew) The number of instance copies. This is not required for standalone and master slave versions.`,
				},
				resource.Attribute{
					Name:        "redis_shard_num",
					Description: `(Optional, ForceNew) The number of instance shard. This is not required for standalone and master slave versions.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Optional, ForceNew) ID of security group. If both vpc_id and subnet_id are not set, this argument should not be set either.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional, ForceNew) Specifies which subnet the instance should belong to.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Instance tags.`,
				},
				resource.Attribute{
					Name:        "type_id",
					Description: `(Optional, ForceNew) Instance type. Available values reference data source ` + "`" + `tencentcloud_redis_zone_config` + "`" + ` or [document](https://intl.cloud.tencent.com/document/product/239/32069).`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional, ForceNew,`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional, ForceNew) ID of the vpc with which the instance is to be associated. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time when the instance was created.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `IP address of an instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of an instance, maybe: init, processing, online, isolate and todelete. ## Import Redis instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_redis_instance.redislab redis-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time when the instance was created.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `IP address of an instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of an instance, maybe: init, processing, online, isolate and todelete. ## Import Redis instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_redis_instance.redislab redis-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_reserved_instance",
			Category:         "Cloud Virtual Machine(CVM)",
			ShortDescription: `Provides a reserved instance resource.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"virtual",
				"machine",
				"cvm",
				"reserved",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required) Configuration ID of the reserved instance.`,
				},
				resource.Attribute{
					Name:        "instance_count",
					Description: `(Required) Number of reserved instances to be purchased. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `Expiry time of the RI.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `Start time of the RI.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the RI at the time of purchase. ## Import Reserved instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_reserved_instance.foo 6cc16e7c-47d7-4fae-9b44-ce5c0f59a920 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `Expiry time of the RI.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `Start time of the RI.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the RI at the time of purchase. ## Import Reserved instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_reserved_instance.foo 6cc16e7c-47d7-4fae-9b44-ce5c0f59a920 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_route_entry",
			Category:         "Virtual Private Cloud(VPC)",
			ShortDescription: `Provides a resource to create a routing entry in a VPC routing table.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"route",
				"entry",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Required, ForceNew) The RouteEntry's target network segment.`,
				},
				resource.Attribute{
					Name:        "next_hub",
					Description: `(Required, ForceNew) The route entry's next hub. CVM instance ID or VPC router interface ID.`,
				},
				resource.Attribute{
					Name:        "next_type",
					Description: `(Required, ForceNew) The next hop type. Valid values: ` + "`" + `public_gateway` + "`" + `,` + "`" + `vpn_gateway` + "`" + `,` + "`" + `sslvpn_gateway` + "`" + `,` + "`" + `dc_gateway` + "`" + `,` + "`" + `peering_connection` + "`" + `,` + "`" + `nat_gateway` + "`" + ` and ` + "`" + `instance` + "`" + `. ` + "`" + `instance` + "`" + ` points to CVM Instance.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `(Required, ForceNew) The ID of the route table.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, ForceNew) The VPC ID. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_route_table",
			Category:         "Virtual Private Cloud(VPC)",
			ShortDescription: `Provides a resource to create a VPC routing table.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"route",
				"table",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of routing table.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, ForceNew) ID of VPC to which the route table should be associated.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The tags of routing table. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the routing table.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Indicates whether it is the default routing table.`,
				},
				resource.Attribute{
					Name:        "route_entry_ids",
					Description: `ID list of the routing entries.`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `ID list of the subnets associated with this route table. ## Import Vpc routetable instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_route_table.test route_table_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the routing table.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Indicates whether it is the default routing table.`,
				},
				resource.Attribute{
					Name:        "route_entry_ids",
					Description: `ID list of the routing entries.`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `ID list of the subnets associated with this route table. ## Import Vpc routetable instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_route_table.test route_table_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_route_table_entry",
			Category:         "Virtual Private Cloud(VPC)",
			ShortDescription: `Provides a resource to create an entry of a routing table.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"route",
				"table",
				"entry",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "destination_cidr_block",
					Description: `(Required, ForceNew) Destination address block.`,
				},
				resource.Attribute{
					Name:        "next_hub",
					Description: `(Required, ForceNew) ID of next-hop gateway. Note: when ` + "`" + `next_type` + "`" + ` is EIP, GatewayId should be ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "next_type",
					Description: `(Required, ForceNew) Type of next-hop. Valid values: ` + "`" + `CVM` + "`" + `, ` + "`" + `VPN` + "`" + `, ` + "`" + `DIRECTCONNECT` + "`" + `, ` + "`" + `PEERCONNECTION` + "`" + `, ` + "`" + `SSLVPN` + "`" + `, ` + "`" + `NAT` + "`" + `, ` + "`" + `NORMAL_CVM` + "`" + `, ` + "`" + `EIP` + "`" + ` and ` + "`" + `CCN` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `(Required, ForceNew) ID of routing table to which this entry belongs.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, ForceNew) Description of the routing table entry. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import Route table entry can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_route_table_entry.foo 83517.rtb-mlhpg09u ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import Route table entry can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_route_table_entry.foo 83517.rtb-mlhpg09u ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_scf_function",
			Category:         "Serverless Cloud Function(SCF)",
			ShortDescription: `Provide a resource to create a SCF function.`,
			Description:      ``,
			Keywords: []string{
				"serverless",
				"cloud",
				"function",
				"scf",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "handler",
					Description: `(Required) Handler of the SCF function. The format of name is ` + "`" + `<filename>.<method_name>` + "`" + `, and it supports 26 English letters, numbers, connectors, and underscores, it should start with a letter. The last character cannot be ` + "`" + `-` + "`" + ` or ` + "`" + `_` + "`" + `. Available length is 2-60.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Name of the SCF function. Name supports 26 English letters, numbers, connectors, and underscores, it should start with a letter. The last character cannot be ` + "`" + `-` + "`" + ` or ` + "`" + `_` + "`" + `. Available length is 2-60.`,
				},
				resource.Attribute{
					Name:        "runtime",
					Description: `(Required) Runtime of the SCF function, only supports ` + "`" + `Python2.7` + "`" + `, ` + "`" + `Python3.6` + "`" + `, ` + "`" + `Nodejs6.10` + "`" + `, ` + "`" + `Nodejs8.9` + "`" + `, ` + "`" + `Nodejs10.15` + "`" + `, ` + "`" + `PHP5` + "`" + `, ` + "`" + `PHP7` + "`" + `, ` + "`" + `Golang1` + "`" + `, and ` + "`" + `Java8` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cls_logset_id",
					Description: `(Optional) cls logset id of the SCF function.`,
				},
				resource.Attribute{
					Name:        "cls_topic_id",
					Description: `(Optional) cls topic id of the SCF function.`,
				},
				resource.Attribute{
					Name:        "cos_bucket_name",
					Description: `(Optional) Cos bucket name of the SCF function, such as ` + "`" + `cos-1234567890` + "`" + `, conflict with ` + "`" + `zip_file` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cos_bucket_region",
					Description: `(Optional) Cos bucket region of the SCF function, conflict with ` + "`" + `zip_file` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cos_object_name",
					Description: `(Optional) Cos object name of the SCF function, should have suffix ` + "`" + `.zip` + "`" + ` or ` + "`" + `.jar` + "`" + `, conflict with ` + "`" + `zip_file` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the SCF function. Description supports English letters, numbers, spaces, commas, newlines, periods and Chinese, the maximum length is 1000.`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `(Optional) Environment of the SCF function.`,
				},
				resource.Attribute{
					Name:        "l5_enable",
					Description: `(Optional) Enable L5 for SCF function, default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mem_size",
					Description: `(Optional) Memory size of the SCF function, unit is MB. The default is ` + "`" + `128` + "`" + `MB. The range is 128M-1536M, and the ladder is 128M.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional, ForceNew) Namespace of the SCF function, default is ` + "`" + `default` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Optional) Role of the SCF function.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) Subnet ID of the SCF function.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of the SCF function.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) Timeout of the SCF function, unit is second. Default ` + "`" + `3` + "`" + `. Available value is 1-900.`,
				},
				resource.Attribute{
					Name:        "triggers",
					Description: `(Optional) Trigger list of the SCF function, note that if you modify the trigger list, all existing triggers will be deleted, and then create triggers in the new list. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) VPC ID of the SCF function.`,
				},
				resource.Attribute{
					Name:        "zip_file",
					Description: `(Optional) Zip file of the SCF function, conflict with ` + "`" + `cos_bucket_name` + "`" + `, ` + "`" + `cos_object_name` + "`" + `, ` + "`" + `cos_bucket_region` + "`" + `. The ` + "`" + `triggers` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the SCF function trigger, if ` + "`" + `type` + "`" + ` is ` + "`" + `ckafka` + "`" + `, the format of name must be ` + "`" + `<ckafkaInstanceId>-<topicId>` + "`" + `; if ` + "`" + `type` + "`" + ` is ` + "`" + `cos` + "`" + `, the name is cos bucket id, other In any case, it can be combined arbitrarily. It can only contain English letters, numbers, connectors and underscores. The maximum length is 100.`,
				},
				resource.Attribute{
					Name:        "trigger_desc",
					Description: `(Required) TriggerDesc of the SCF function trigger, parameter format of ` + "`" + `timer` + "`" + ` is linux cron expression; parameter of ` + "`" + `cos` + "`" + ` type is json string ` + "`" + `{"event":"cos:ObjectCreated:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type of the SCF function trigger, support ` + "`" + `cos` + "`" + `, ` + "`" + `cmq` + "`" + `, ` + "`" + `timer` + "`" + `, ` + "`" + `ckafka` + "`" + `, ` + "`" + `apigw` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cos_region",
					Description: `(Optional) Region of cos bucket. if ` + "`" + `type` + "`" + ` is ` + "`" + `cos` + "`" + `, ` + "`" + `cos_region` + "`" + ` is required. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "code_error",
					Description: `SCF function code error message.`,
				},
				resource.Attribute{
					Name:        "code_result",
					Description: `SCF function code is correct.`,
				},
				resource.Attribute{
					Name:        "code_size",
					Description: `SCF function code size, unit is M.`,
				},
				resource.Attribute{
					Name:        "eip_fixed",
					Description: `Whether EIP is a fixed IP.`,
				},
				resource.Attribute{
					Name:        "eips",
					Description: `SCF function EIP list.`,
				},
				resource.Attribute{
					Name:        "err_no",
					Description: `SCF function code error code.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `SCF function domain name.`,
				},
				resource.Attribute{
					Name:        "install_dependency",
					Description: `Whether to automatically install dependencies.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `SCF function last modified time.`,
				},
				resource.Attribute{
					Name:        "status_desc",
					Description: `SCF status description.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `SCF function status.`,
				},
				resource.Attribute{
					Name:        "trigger_info",
					Description: `SCF trigger details list. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "custom_argument",
					Description: `User-defined parameters of SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `Whether SCF function trigger is enable.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `Modify time of SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "trigger_desc",
					Description: `TriggerDesc of SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "vip",
					Description: `SCF function vip. ## Import SCF function can be imported, e.g. ->`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "code_error",
					Description: `SCF function code error message.`,
				},
				resource.Attribute{
					Name:        "code_result",
					Description: `SCF function code is correct.`,
				},
				resource.Attribute{
					Name:        "code_size",
					Description: `SCF function code size, unit is M.`,
				},
				resource.Attribute{
					Name:        "eip_fixed",
					Description: `Whether EIP is a fixed IP.`,
				},
				resource.Attribute{
					Name:        "eips",
					Description: `SCF function EIP list.`,
				},
				resource.Attribute{
					Name:        "err_no",
					Description: `SCF function code error code.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `SCF function domain name.`,
				},
				resource.Attribute{
					Name:        "install_dependency",
					Description: `Whether to automatically install dependencies.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `SCF function last modified time.`,
				},
				resource.Attribute{
					Name:        "status_desc",
					Description: `SCF status description.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `SCF function status.`,
				},
				resource.Attribute{
					Name:        "trigger_info",
					Description: `SCF trigger details list. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "custom_argument",
					Description: `User-defined parameters of SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `Whether SCF function trigger is enable.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `Modify time of SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "trigger_desc",
					Description: `TriggerDesc of SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "vip",
					Description: `SCF function vip. ## Import SCF function can be imported, e.g. ->`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_scf_namespace",
			Category:         "Serverless Cloud Function(SCF)",
			ShortDescription: `Provide a resource to create a SCF namespace.`,
			Description:      ``,
			Keywords: []string{
				"serverless",
				"cloud",
				"function",
				"scf",
				"namespace",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Required, ForceNew) Name of the SCF namespace.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the SCF namespace. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `SCF namespace creation time.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `SCF namespace last modified time.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `SCF namespace type. ## Import SCF namespace can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_scf_function.test default ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `SCF namespace creation time.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `SCF namespace last modified time.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `SCF namespace type. ## Import SCF namespace can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_scf_function.test default ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_security_group",
			Category:         "Virtual Private Cloud(VPC)",
			ShortDescription: `Provides a resource to create security group.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"security",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the security group to be queried.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the security group.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional, ForceNew) Project ID of the security group.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of the security group. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import Security group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_security_group.sglab sg-ey3wmiz1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import Security group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_security_group.sglab sg-ey3wmiz1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_security_group_lite_rule",
			Category:         "Virtual Private Cloud(VPC)",
			ShortDescription: `Provide a resource to create security group some lite rules quickly.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"security",
				"group",
				"lite",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Required, ForceNew) ID of the security group.`,
				},
				resource.Attribute{
					Name:        "egress",
					Description: `(Optional) Egress rules set. A rule must match the following format: [action]#[cidr_ip]#[port]#[protocol]. The available value of 'action' is ` + "`" + `ACCEPT` + "`" + ` and ` + "`" + `DROP` + "`" + `. The 'cidr_ip' must be an IP address network or segment. The 'port' valid format is ` + "`" + `80` + "`" + `, ` + "`" + `80,443` + "`" + `, ` + "`" + `80-90` + "`" + ` or ` + "`" + `ALL` + "`" + `. The available value of 'protocol' is ` + "`" + `TCP` + "`" + `, ` + "`" + `UDP` + "`" + `, ` + "`" + `ICMP` + "`" + ` and ` + "`" + `ALL` + "`" + `. When 'protocol' is ` + "`" + `ICMP` + "`" + ` or ` + "`" + `ALL` + "`" + `, the 'port' must be ` + "`" + `ALL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ingress",
					Description: `(Optional) Ingress rules set. A rule must match the following format: [action]#[cidr_ip]#[port]#[protocol]. The available value of 'action' is ` + "`" + `ACCEPT` + "`" + ` and ` + "`" + `DROP` + "`" + `. The 'cidr_ip' must be an IP address network or segment. The 'port' valid format is ` + "`" + `80` + "`" + `, ` + "`" + `80,443` + "`" + `, ` + "`" + `80-90` + "`" + ` or ` + "`" + `ALL` + "`" + `. The available value of 'protocol' is ` + "`" + `TCP` + "`" + `, ` + "`" + `UDP` + "`" + `, ` + "`" + `ICMP` + "`" + ` and ` + "`" + `ALL` + "`" + `. When 'protocol' is ` + "`" + `ICMP` + "`" + ` or ` + "`" + `ALL` + "`" + `, the 'port' must be ` + "`" + `ALL` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import Security group lite rule can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_security_group_lite_rule.foo sg-ey3wmiz1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import Security group lite rule can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_security_group_lite_rule.foo sg-ey3wmiz1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_security_group_rule",
			Category:         "Virtual Private Cloud(VPC)",
			ShortDescription: `Provides a resource to create security group rule.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"security",
				"group",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policy",
					Description: `(Required, ForceNew) Rule policy of security group. Valid values: ` + "`" + `ACCEPT` + "`" + ` and ` + "`" + `DROP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Required, ForceNew) ID of the security group to be queried.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required, ForceNew) Type of the security group rule. Valid values: ` + "`" + `ingress` + "`" + ` and ` + "`" + `egress` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "address_template",
					Description: `(Optional, ForceNew) ID of the address template, and confilicts with ` + "`" + `source_sgid` + "`" + ` and ` + "`" + `cidr_ip` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cidr_ip",
					Description: `(Optional, ForceNew) An IP address network or segment, and conflict with ` + "`" + `source_sgid` + "`" + ` and ` + "`" + `address_template` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, ForceNew) Description of the security group rule.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `(Optional, ForceNew) Type of IP protocol. Valid values: ` + "`" + `TCP` + "`" + `, ` + "`" + `UDP` + "`" + ` and ` + "`" + `ICMP` + "`" + `. Default to all types protocol, and conflicts with ` + "`" + `protocol_template` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "port_range",
					Description: `(Optional, ForceNew) Range of the port. The available value can be one, multiple or one segment. E.g. ` + "`" + `80` + "`" + `, ` + "`" + `80,90` + "`" + ` and ` + "`" + `80-90` + "`" + `. Default to all ports, and confilicts with ` + "`" + `protocol_template` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "protocol_template",
					Description: `(Optional, ForceNew) ID of the address template, and conflict with ` + "`" + `ip_protocol` + "`" + `, ` + "`" + `port_range` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "source_sgid",
					Description: `(Optional, ForceNew) ID of the nested security group, and conflicts with ` + "`" + `cidr_ip` + "`" + ` and ` + "`" + `address_template` + "`" + `. The ` + "`" + `address_template` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Optional, ForceNew) Address template group ID, conflicts with ` + "`" + `template_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `(Optional, ForceNew) Address template ID, conflicts with ` + "`" + `group_id` + "`" + `. The ` + "`" + `protocol_template` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Optional, ForceNew) Address template group ID, conflicts with ` + "`" + `template_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `(Optional, ForceNew) Address template ID, conflicts with ` + "`" + `group_id` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_sqlserver_account",
			Category:         "SQLServer",
			ShortDescription: `Use this resource to create SQL Server account`,
			Description:      ``,
			Keywords: []string{
				"sqlserver",
				"account",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, ForceNew) Instance ID that the account belongs to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the SQL Server account.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Password of the SQL Server account.`,
				},
				resource.Attribute{
					Name:        "is_admin",
					Description: `(Optional) Indicate that the account is root account or not.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `(Optional) Remark of the SQL Server account. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the SQL Server account.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the SQL Server account. Valid values: 1, 2, 3, 4. 1 for creating, 2 for running, 3 for modifying, 4 for resetting password, -1 for deleting.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Last updated time of the SQL Server account. ## Import SQL Server account can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_sqlserver_account.foo mssql-3cdq7kx5#tf_sqlserver_account ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the SQL Server account.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the SQL Server account. Valid values: 1, 2, 3, 4. 1 for creating, 2 for running, 3 for modifying, 4 for resetting password, -1 for deleting.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Last updated time of the SQL Server account. ## Import SQL Server account can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_sqlserver_account.foo mssql-3cdq7kx5#tf_sqlserver_account ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_sqlserver_account_db_attachment",
			Category:         "SQLServer",
			ShortDescription: `Use this resource to create SQL Server account DB attachment`,
			Description:      ``,
			Keywords: []string{
				"sqlserver",
				"account",
				"db",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_name",
					Description: `(Required, ForceNew) SQL Server account name.`,
				},
				resource.Attribute{
					Name:        "db_name",
					Description: `(Required, ForceNew) SQL Server DB name.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, ForceNew) SQL Server instance ID that the account belongs to.`,
				},
				resource.Attribute{
					Name:        "privilege",
					Description: `(Required) Privilege of the account on DB. Valid values: ` + "`" + `ReadOnly` + "`" + `, ` + "`" + `ReadWrite` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import SQL Server account DB attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_sqlserver_account_db_attachment.foo mssql-3cdq7kx5#tf_sqlserver_account#test111 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import SQL Server account DB attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_sqlserver_account_db_attachment.foo mssql-3cdq7kx5#tf_sqlserver_account#test111 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_sqlserver_basic_instance",
			Category:         "SQLServer",
			ShortDescription: `Provides a SQL Server instance resource to create basic database instances.`,
			Description:      ``,
			Keywords: []string{
				"sqlserver",
				"basic",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cpu",
					Description: `(Required) The CPU number of the SQL Server basic instance.`,
				},
				resource.Attribute{
					Name:        "machine_type",
					Description: `(Required) The host type of the purchased instance, ` + "`" + `CLOUD_PREMIUM` + "`" + ` for virtual machine high-performance cloud disk, ` + "`" + `CLOUD_SSD` + "`" + ` for virtual machine SSD cloud disk.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `(Required) Memory size (in GB). Allowed value must be larger than ` + "`" + `memory` + "`" + ` that data source ` + "`" + `tencentcloud_sqlserver_specinfos` + "`" + ` provides.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the SQL Server basic instance.`,
				},
				resource.Attribute{
					Name:        "storage",
					Description: `(Required) Disk size (in GB). Allowed value must be a multiple of 10. The storage must be set with the limit of ` + "`" + `storage_min` + "`" + ` and ` + "`" + `storage_max` + "`" + ` which data source ` + "`" + `tencentcloud_sqlserver_specinfos` + "`" + ` provides.`,
				},
				resource.Attribute{
					Name:        "auto_renew",
					Description: `(Optional) Automatic renewal sign. 0 for normal renewal, 1 for automatic renewal, the default is 1 automatic renewal. Only valid when purchasing a prepaid instance.`,
				},
				resource.Attribute{
					Name:        "auto_voucher",
					Description: `(Optional) Whether to use the voucher automatically; 1 for yes, 0 for no, the default is 0.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional, ForceNew) Availability zone.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `(Optional, ForceNew) Pay type of the SQL Server basic instance. For now, only ` + "`" + `POSTPAID_BY_HOUR` + "`" + ` is valid.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `(Optional, ForceNew) Version of the SQL Server basic database engine. Allowed values are ` + "`" + `2008R2` + "`" + `(SQL Server 2008 Enterprise), ` + "`" + `2012SP3` + "`" + `(SQL Server 2012 Enterprise), ` + "`" + `2016SP1` + "`" + ` (SQL Server 2016 Enterprise), ` + "`" + `201602` + "`" + `(SQL Server 2016 Standard) and ` + "`" + `2017` + "`" + `(SQL Server 2017 Enterprise). Default is ` + "`" + `2008R2` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "maintenance_start_time",
					Description: `(Optional) Start time of the maintenance in one day, format like ` + "`" + `HH:mm` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "maintenance_time_span",
					Description: `(Optional) The timespan of maintenance in one day, unit is hour.`,
				},
				resource.Attribute{
					Name:        "maintenance_week_set",
					Description: `(Optional) A list of integer indicates weekly maintenance. For example, [1,7] presents do weekly maintenance on every Monday and Sunday.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) Purchase instance period, the default value is 1, which means one month. The value does not exceed 48.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) Project ID, default value is 0.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Optional) Security group bound to the instance.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional, ForceNew) ID of subnet.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The tags of the SQL Server basic instance.`,
				},
				resource.Attribute{
					Name:        "voucher_ids",
					Description: `(Optional) An array of voucher IDs, currently only one can be used for a single order.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional, ForceNew) ID of VPC. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the SQL Server basic instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the SQL Server basic instance. 1 for applying, 2 for running, 3 for running with limit, 4 for isolated, 5 for recycling, 6 for recycled, 7 for running with task, 8 for off-line, 9 for expanding, 10 for migrating, 11 for readonly, 12 for rebooting.`,
				},
				resource.Attribute{
					Name:        "vip",
					Description: `IP for private access.`,
				},
				resource.Attribute{
					Name:        "vport",
					Description: `Port for private access. ## Import SQL Server basic instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_sqlserver_basic_instance.foo mssql-3cdq7kx5 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the SQL Server basic instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the SQL Server basic instance. 1 for applying, 2 for running, 3 for running with limit, 4 for isolated, 5 for recycling, 6 for recycled, 7 for running with task, 8 for off-line, 9 for expanding, 10 for migrating, 11 for readonly, 12 for rebooting.`,
				},
				resource.Attribute{
					Name:        "vip",
					Description: `IP for private access.`,
				},
				resource.Attribute{
					Name:        "vport",
					Description: `Port for private access. ## Import SQL Server basic instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_sqlserver_basic_instance.foo mssql-3cdq7kx5 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_sqlserver_db",
			Category:         "SQLServer",
			ShortDescription: `Provides a SQL Server DB resource belongs to SQL Server instance.`,
			Description:      ``,
			Keywords: []string{
				"sqlserver",
				"db",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, ForceNew) SQL Server instance ID which DB belongs to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Name of SQL Server DB. The database name must be unique and must be composed of numbers, letters and underlines, and the first one can not be underline.`,
				},
				resource.Attribute{
					Name:        "charset",
					Description: `(Optional, ForceNew) Character set DB uses. Valid values: ` + "`" + `Chinese_PRC_CI_AS` + "`" + `, ` + "`" + `Chinese_PRC_CS_AS` + "`" + `, ` + "`" + `Chinese_PRC_BIN` + "`" + `, ` + "`" + `Chinese_Taiwan_Stroke_CI_AS` + "`" + `, ` + "`" + `SQL_Latin1_General_CP1_CI_AS` + "`" + `, and ` + "`" + `SQL_Latin1_General_CP1_CS_AS` + "`" + `. Default value is ` + "`" + `Chinese_PRC_CI_AS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `(Optional) Remark of the DB. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Database creation time.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Database status, could be ` + "`" + `creating` + "`" + `, ` + "`" + `running` + "`" + `, ` + "`" + `modifying` + "`" + ` which means changing the remark, and ` + "`" + `deleting` + "`" + `. ## Import SQL Server DB can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_sqlserver_db.foo mssql-3cdq7kx5#db_name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Database creation time.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Database status, could be ` + "`" + `creating` + "`" + `, ` + "`" + `running` + "`" + `, ` + "`" + `modifying` + "`" + ` which means changing the remark, and ` + "`" + `deleting` + "`" + `. ## Import SQL Server DB can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_sqlserver_db.foo mssql-3cdq7kx5#db_name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_sqlserver_instance",
			Category:         "SQLServer",
			ShortDescription: `Use this resource to create SQL Server instance`,
			Description:      ``,
			Keywords: []string{
				"sqlserver",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "memory",
					Description: `(Required) Memory size (in GB). Allowed value must be larger than ` + "`" + `memory` + "`" + ` that data source ` + "`" + `tencentcloud_sqlserver_specinfos` + "`" + ` provides.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the SQL Server instance.`,
				},
				resource.Attribute{
					Name:        "storage",
					Description: `(Required) Disk size (in GB). Allowed value must be a multiple of 10. The storage must be set with the limit of ` + "`" + `storage_min` + "`" + ` and ` + "`" + `storage_max` + "`" + ` which data source ` + "`" + `tencentcloud_sqlserver_specinfos` + "`" + ` provides.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional, ForceNew) Availability zone.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `(Optional, ForceNew) Pay type of the SQL Server instance. For now, only ` + "`" + `POSTPAID_BY_HOUR` + "`" + ` is valid.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `(Optional, ForceNew) Version of the SQL Server database engine. Allowed values are ` + "`" + `2008R2` + "`" + `(SQL Server 2008 Enterprise), ` + "`" + `2012SP3` + "`" + `(SQL Server 2012 Enterprise), ` + "`" + `2016SP1` + "`" + ` (SQL Server 2016 Enterprise), ` + "`" + `201602` + "`" + `(SQL Server 2016 Standard) and ` + "`" + `2017` + "`" + `(SQL Server 2017 Enterprise). Default is ` + "`" + `2008R2` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ha_type",
					Description: `(Optional, ForceNew) Instance type. ` + "`" + `DUAL` + "`" + ` (dual-server high availability), ` + "`" + `CLUSTER` + "`" + ` (cluster). Default is ` + "`" + `DUAL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "maintenance_start_time",
					Description: `(Optional) Start time of the maintenance in one day, format like ` + "`" + `HH:mm` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "maintenance_time_span",
					Description: `(Optional) The timespan of maintenance in one day, unit is hour.`,
				},
				resource.Attribute{
					Name:        "maintenance_week_set",
					Description: `(Optional) A list of integer indicates weekly maintenance. For example, [2,7] presents do weekly maintenance on every Tuesday and Sunday.`,
				},
				resource.Attribute{
					Name:        "multi_zones",
					Description: `(Optional, ForceNew) Indicate whether to deploy across availability zones.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) Project ID, default value is 0.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Optional) Security group bound to the instance.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional, ForceNew) ID of subnet.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The tags of the SQL Server.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional, ForceNew) ID of VPC. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the SQL Server instance.`,
				},
				resource.Attribute{
					Name:        "ro_flag",
					Description: `Readonly flag. ` + "`" + `RO` + "`" + ` (read-only instance), ` + "`" + `MASTER` + "`" + ` (primary instance with read-only instances). If it is left empty, it refers to an instance which is not read-only and has no RO group.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the SQL Server instance. 1 for applying, 2 for running, 3 for running with limit, 4 for isolated, 5 for recycling, 6 for recycled, 7 for running with task, 8 for off-line, 9 for expanding, 10 for migrating, 11 for readonly, 12 for rebooting.`,
				},
				resource.Attribute{
					Name:        "vip",
					Description: `IP for private access.`,
				},
				resource.Attribute{
					Name:        "vport",
					Description: `Port for private access. ## Import SQL Server instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_sqlserver_instance.foo mssql-3cdq7kx5 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the SQL Server instance.`,
				},
				resource.Attribute{
					Name:        "ro_flag",
					Description: `Readonly flag. ` + "`" + `RO` + "`" + ` (read-only instance), ` + "`" + `MASTER` + "`" + ` (primary instance with read-only instances). If it is left empty, it refers to an instance which is not read-only and has no RO group.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the SQL Server instance. 1 for applying, 2 for running, 3 for running with limit, 4 for isolated, 5 for recycling, 6 for recycled, 7 for running with task, 8 for off-line, 9 for expanding, 10 for migrating, 11 for readonly, 12 for rebooting.`,
				},
				resource.Attribute{
					Name:        "vip",
					Description: `IP for private access.`,
				},
				resource.Attribute{
					Name:        "vport",
					Description: `Port for private access. ## Import SQL Server instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_sqlserver_instance.foo mssql-3cdq7kx5 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_sqlserver_publish_subscribe",
			Category:         "SQLServer",
			ShortDescription: `Provides a SQL Server PublishSubscribe resource belongs to SQL Server instance.`,
			Description:      ``,
			Keywords: []string{
				"sqlserver",
				"publish",
				"subscribe",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "database_tuples",
					Description: `(Required) Database Publish and Publish relationship list. The elements inside can be deleted and added individually, but modification is not allowed.`,
				},
				resource.Attribute{
					Name:        "publish_instance_id",
					Description: `(Required, ForceNew) ID of the SQL Server instance which publish.`,
				},
				resource.Attribute{
					Name:        "subscribe_instance_id",
					Description: `(Required, ForceNew) ID of the SQL Server instance which subscribe.`,
				},
				resource.Attribute{
					Name:        "delete_subscribe_db",
					Description: `(Optional) Whether to delete the subscriber database when deleting the Publish and Subscribe. ` + "`" + `true` + "`" + ` for deletes the subscribe database, ` + "`" + `false` + "`" + ` for does not delete the subscribe database. default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "publish_subscribe_name",
					Description: `(Optional) The name of the Publish and Subscribe. Default is ` + "`" + `default_name` + "`" + `. The ` + "`" + `database_tuples` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "publish_database",
					Description: `(Required) Publish the database. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import SQL Server PublishSubscribe can be imported using the publish_sqlserver_id#subscribe_sqlserver_id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_sqlserver_publish_subscribe.foo publish_sqlserver_id#subscribe_sqlserver_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import SQL Server PublishSubscribe can be imported using the publish_sqlserver_id#subscribe_sqlserver_id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_sqlserver_publish_subscribe.foo publish_sqlserver_id#subscribe_sqlserver_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_sqlserver_readonly_instance",
			Category:         "SQLServer",
			ShortDescription: `Provides a SQL Server instance resource to create read-only database instances.`,
			Description:      ``,
			Keywords: []string{
				"sqlserver",
				"readonly",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "master_instance_id",
					Description: `(Required, ForceNew) Indicates the master instance ID of recovery instances.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `(Required) Memory size (in GB). Allowed value must be larger than ` + "`" + `memory` + "`" + ` that data source ` + "`" + `tencentcloud_sqlserver_specinfos` + "`" + ` provides.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the SQL Server instance.`,
				},
				resource.Attribute{
					Name:        "readonly_group_type",
					Description: `(Required, ForceNew) Type of readonly group. Valid values: ` + "`" + `1` + "`" + `, ` + "`" + `3` + "`" + `. ` + "`" + `1` + "`" + ` for one auto-assigned readonly instance per one readonly group, ` + "`" + `2` + "`" + ` for creating new readonly group, ` + "`" + `3` + "`" + ` for all exist readonly instances stay in the exist readonly group. For now, only ` + "`" + `1` + "`" + ` and ` + "`" + `3` + "`" + ` are supported.`,
				},
				resource.Attribute{
					Name:        "storage",
					Description: `(Required) Disk size (in GB). Allowed value must be a multiple of 10. The storage must be set with the limit of ` + "`" + `storage_min` + "`" + ` and ` + "`" + `storage_max` + "`" + ` which data source ` + "`" + `tencentcloud_sqlserver_specinfos` + "`" + ` provides.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional, ForceNew) Availability zone.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `(Optional, ForceNew) Pay type of the SQL Server instance. For now, only ` + "`" + `POSTPAID_BY_HOUR` + "`" + ` is valid.`,
				},
				resource.Attribute{
					Name:        "force_upgrade",
					Description: `(Optional, ForceNew) Indicate that the master instance upgrade or not. ` + "`" + `true` + "`" + ` for upgrading the master SQL Server instance to cluster type by force. Default is false. Note: this is not supported with ` + "`" + `DUAL` + "`" + `(ha_type), ` + "`" + `2017` + "`" + `(engine_version) master SQL Server instance, for it will cause ha_type of the master SQL Server instance change.`,
				},
				resource.Attribute{
					Name:        "readonly_group_id",
					Description: `(Optional) ID of the readonly group that this instance belongs to. When ` + "`" + `readonly_group_type` + "`" + ` set value ` + "`" + `3` + "`" + `, it must be set with valid value.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Optional) Security group bound to the instance.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional, ForceNew) ID of subnet.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The tags of the SQL Server.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional, ForceNew) ID of VPC. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the SQL Server instance.`,
				},
				resource.Attribute{
					Name:        "ro_flag",
					Description: `Readonly flag. ` + "`" + `RO` + "`" + ` (read-only instance), ` + "`" + `MASTER` + "`" + ` (primary instance with read-only instances). If it is left empty, it refers to an instance which is not read-only and has no RO group.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the SQL Server instance. 1 for applying, 2 for running, 3 for running with limit, 4 for isolated, 5 for recycling, 6 for recycled, 7 for running with task, 8 for off-line, 9 for expanding, 10 for migrating, 11 for readonly, 12 for rebooting.`,
				},
				resource.Attribute{
					Name:        "vip",
					Description: `IP for private access.`,
				},
				resource.Attribute{
					Name:        "vport",
					Description: `Port for private access. ## Import SQL Server readonly instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_sqlserver_readonly_instance.foo mssqlro-3cdq7kx5 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the SQL Server instance.`,
				},
				resource.Attribute{
					Name:        "ro_flag",
					Description: `Readonly flag. ` + "`" + `RO` + "`" + ` (read-only instance), ` + "`" + `MASTER` + "`" + ` (primary instance with read-only instances). If it is left empty, it refers to an instance which is not read-only and has no RO group.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the SQL Server instance. 1 for applying, 2 for running, 3 for running with limit, 4 for isolated, 5 for recycling, 6 for recycled, 7 for running with task, 8 for off-line, 9 for expanding, 10 for migrating, 11 for readonly, 12 for rebooting.`,
				},
				resource.Attribute{
					Name:        "vip",
					Description: `IP for private access.`,
				},
				resource.Attribute{
					Name:        "vport",
					Description: `Port for private access. ## Import SQL Server readonly instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_sqlserver_readonly_instance.foo mssqlro-3cdq7kx5 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_ssl_certificate",
			Category:         "SSL Certificates",
			ShortDescription: `Provides a resource to create a SSL certificate.`,
			Description:      ``,
			Keywords: []string{
				"ssl",
				"certificates",
				"certificate",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cert",
					Description: `(Required, ForceNew) Content of the SSL certificate. Not allowed newline at the start and end.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required, ForceNew) Type of the SSL certificate. Valid values: ` + "`" + `CA` + "`" + ` and ` + "`" + `SVR` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional, ForceNew) Key of the SSL certificate and required when certificate type is ` + "`" + `SVR` + "`" + `. Not allowed newline at the start and end.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, ForceNew) Name of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional, ForceNew) Project ID of the SSL certificate. Default is ` + "`" + `0` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "begin_time",
					Description: `Beginning time of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Primary domain of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `Ending time of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "product_zh_name",
					Description: `Certificate authority.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "subject_names",
					Description: `ALL domains included in the SSL certificate. Including the primary domain name. ## Import ssl certificate can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_ssl_certificate.cert GjTNRoK7 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "begin_time",
					Description: `Beginning time of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Primary domain of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `Ending time of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "product_zh_name",
					Description: `Certificate authority.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "subject_names",
					Description: `ALL domains included in the SSL certificate. Including the primary domain name. ## Import ssl certificate can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_ssl_certificate.cert GjTNRoK7 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_ssl_pay_certificate",
			Category:         "SSL Certificates",
			ShortDescription: `Provide a resource to create a payment SSL.`,
			Description:      ``,
			Keywords: []string{
				"ssl",
				"certificates",
				"pay",
				"certificate",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain_num",
					Description: `(Required, ForceNew) Number of domain names included in the certificate.`,
				},
				resource.Attribute{
					Name:        "information",
					Description: `(Required, ForceNew) Certificate information.`,
				},
				resource.Attribute{
					Name:        "product_id",
					Description: `(Required, ForceNew) Certificate commodity ID. Valid value ranges: (3~42). ` + "`" + `3` + "`" + ` means SecureSite Enhanced Enterprise Edition (EV Pro), ` + "`" + `4` + "`" + ` means SecureSite Enhanced (EV), ` + "`" + `5` + "`" + ` means SecureSite Enterprise Professional Edition (OV Pro), ` + "`" + `6` + "`" + ` means SecureSite Enterprise (OV), ` + "`" + `7` + "`" + ` means SecureSite Enterprise Type (OV) wildcard, ` + "`" + `8` + "`" + ` means Geotrust enhanced (EV), ` + "`" + `9` + "`" + ` means Geotrust enterprise (OV), ` + "`" + `10` + "`" + ` means Geotrust enterprise (OV) wildcard, ` + "`" + `11` + "`" + ` means TrustAsia domain type multi-domain SSL certificate, ` + "`" + `12` + "`" + ` means TrustAsia domain type ( DV) wildcard, ` + "`" + `13` + "`" + ` means TrustAsia enterprise wildcard (OV) SSL certificate (D3), ` + "`" + `14` + "`" + ` means TrustAsia enterprise (OV) SSL certificate (D3), ` + "`" + `15` + "`" + ` means TrustAsia enterprise multi-domain (OV) SSL certificate (D3), ` + "`" + `16` + "`" + ` means TrustAsia Enhanced (EV) SSL Certificate (D3), ` + "`" + `17` + "`" + ` means TrustAsia Enhanced Multiple Domain (EV) SSL Certificate (D3), ` + "`" + `18` + "`" + ` means GlobalSign Enterprise (OV) SSL Certificate, ` + "`" + `19` + "`" + ` means GlobalSign Enterprise Wildcard (OV) SSL Certificate, ` + "`" + `20` + "`" + ` means GlobalSign Enhanced (EV) SSL Certificate, ` + "`" + `21` + "`" + ` means TrustAsia Enterprise Wildcard Multiple Domain (OV) SSL Certificate (D3), ` + "`" + `22` + "`" + ` means GlobalSign Enterprise Multiple Domain (OV) SSL Certificate, ` + "`" + `23` + "`" + ` means GlobalSign Enterprise Multiple Wildcard Domain name (OV) SSL certificate, ` + "`" + `24` + "`" + ` means GlobalSign enhanced multi-domain (EV) SSL certificate, ` + "`" + `25` + "`" + ` means Wotrus domain type certificate, ` + "`" + `26` + "`" + ` means Wotrus domain type multi-domain certificate, ` + "`" + `27` + "`" + ` means Wotrus domain type wildcard certificate, ` + "`" + `28` + "`" + ` means Wotrus enterprise type certificate, ` + "`" + `29` + "`" + ` means Wotrus enterprise multi-domain certificate, ` + "`" + `30` + "`" + ` means Wotrus enterprise wildcard certificate, ` + "`" + `31` + "`" + ` means Wotrus enhanced certificate, ` + "`" + `32` + "`" + ` means Wotrus enhanced multi-domain certificate, ` + "`" + `33` + "`" + ` means DNSPod national secret domain name certificate, ` + "`" + `34` + "`" + ` means DNSPod national secret domain name certificate Multi-domain certificate, ` + "`" + `35` + "`" + ` means DNSPod national secret domain name wildcard certificate, ` + "`" + `37` + "`" + ` means DNSPod national secret enterprise certificate, ` + "`" + `38` + "`" + ` means DNSPod national secret enterprise multi-domain certificate, ` + "`" + `39` + "`" + ` means DNSPod national secret enterprise wildcard certificate, ` + "`" + `40` + "`" + ` means DNSPod national secret increase Strong certificate, ` + "`" + `41` + "`" + ` means DNSPod national secret enhanced multi-domain certificate, ` + "`" + `42` + "`" + ` means TrustAsia domain-type wildcard multi-domain certificate.`,
				},
				resource.Attribute{
					Name:        "alias",
					Description: `(Optional) Remark name.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) The ID of project.`,
				},
				resource.Attribute{
					Name:        "time_span",
					Description: `(Optional) Certificate period, currently only supports 1 year certificate purchase. The ` + "`" + `information` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "admin_email",
					Description: `(Required, ForceNew) The administrator's email address.`,
				},
				resource.Attribute{
					Name:        "admin_first_name",
					Description: `(Required, ForceNew) The first name of the administrator.`,
				},
				resource.Attribute{
					Name:        "admin_last_name",
					Description: `(Required, ForceNew) The last name of the administrator.`,
				},
				resource.Attribute{
					Name:        "admin_phone_num",
					Description: `(Required, ForceNew) Manager mobile phone number.`,
				},
				resource.Attribute{
					Name:        "admin_position",
					Description: `(Required, ForceNew) Manager position.`,
				},
				resource.Attribute{
					Name:        "certificate_domain",
					Description: `(Required, ForceNew) Domain name for binding certificate.`,
				},
				resource.Attribute{
					Name:        "contact_email",
					Description: `(Required, ForceNew) Contact email address.`,
				},
				resource.Attribute{
					Name:        "contact_first_name",
					Description: `(Required, ForceNew) Contact first name.`,
				},
				resource.Attribute{
					Name:        "contact_last_name",
					Description: `(Required, ForceNew) Contact last name.`,
				},
				resource.Attribute{
					Name:        "contact_number",
					Description: `(Required, ForceNew) Contact phone number.`,
				},
				resource.Attribute{
					Name:        "contact_position",
					Description: `(Required, ForceNew) Contact position.`,
				},
				resource.Attribute{
					Name:        "organization_address",
					Description: `(Required, ForceNew) Company address.`,
				},
				resource.Attribute{
					Name:        "organization_city",
					Description: `(Required, ForceNew) Company city.`,
				},
				resource.Attribute{
					Name:        "organization_country",
					Description: `(Required, ForceNew) Country name, such as China: CN.`,
				},
				resource.Attribute{
					Name:        "organization_division",
					Description: `(Required, ForceNew) Department name.`,
				},
				resource.Attribute{
					Name:        "organization_name",
					Description: `(Required, ForceNew) Company name.`,
				},
				resource.Attribute{
					Name:        "organization_region",
					Description: `(Required, ForceNew) The province where the company is located.`,
				},
				resource.Attribute{
					Name:        "phone_area_code",
					Description: `(Required, ForceNew) Company landline area code.`,
				},
				resource.Attribute{
					Name:        "phone_number",
					Description: `(Required, ForceNew) Company landline number.`,
				},
				resource.Attribute{
					Name:        "postal_code",
					Description: `(Required, ForceNew) Company postal code.`,
				},
				resource.Attribute{
					Name:        "verify_type",
					Description: `(Required, ForceNew) Certificate verification method. Valid values: ` + "`" + `DNS_AUTO` + "`" + `, ` + "`" + `DNS` + "`" + `, ` + "`" + `FILE` + "`" + `. ` + "`" + `DNS_AUTO` + "`" + ` means automatic DNS verification, this verification type is only supported for domain names resolved by Tencent Cloud and the resolution status is normal, ` + "`" + `DNS` + "`" + ` means manual DNS verification, ` + "`" + `FILE` + "`" + ` means file verification.`,
				},
				resource.Attribute{
					Name:        "csr_content",
					Description: `(Optional, ForceNew) CSR content uploaded.`,
				},
				resource.Attribute{
					Name:        "csr_type",
					Description: `(Optional, ForceNew) CSR generation method. Valid values: ` + "`" + `online` + "`" + `, ` + "`" + `parse` + "`" + `. ` + "`" + `online` + "`" + ` means online generation, ` + "`" + `parse` + "`" + ` means manual upload.`,
				},
				resource.Attribute{
					Name:        "domain_list",
					Description: `(Optional, ForceNew) Array of uploaded domain names, multi-domain certificates can be uploaded.`,
				},
				resource.Attribute{
					Name:        "key_password",
					Description: `(Optional, ForceNew) Private key password. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `Returned certificate ID.`,
				},
				resource.Attribute{
					Name:        "order_id",
					Description: `Order ID returned.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `SSL certificate status. ## Import payment SSL instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_ssl_pay_certificate.ssl iPQNn61x#33#1#1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `Returned certificate ID.`,
				},
				resource.Attribute{
					Name:        "order_id",
					Description: `Order ID returned.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `SSL certificate status. ## Import payment SSL instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_ssl_pay_certificate.ssl iPQNn61x#33#1#1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_ssm_secret",
			Category:         "SSM",
			ShortDescription: `Provide a resource to create a SSM secret.`,
			Description:      ``,
			Keywords: []string{
				"ssm",
				"secret",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "secret_name",
					Description: `(Required, ForceNew) Name of secret which cannot be repeated in the same region. The maximum length is 128 bytes. The name can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of secret. The maximum is 2048 bytes.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `(Optional) Specify whether to enable secret. Default value is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "kms_key_id",
					Description: `(Optional, ForceNew) KMS keyId used to encrypt secret. If it is empty, it means that the CMK created by SSM for you by default is used for encryption. You can also specify the KMS CMK created by yourself in the same region for encryption.`,
				},
				resource.Attribute{
					Name:        "recovery_window_in_days",
					Description: `(Optional) Specify the scheduled deletion date. Default value is ` + "`" + `0` + "`" + ` that means to delete immediately. 1-30 means the number of days reserved, completely deleted after this date.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of secret. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of secret. ## Import SSM secret can be imported using the secretName, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_ssm_secret.foo test ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of secret. ## Import SSM secret can be imported using the secretName, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_ssm_secret.foo test ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_ssm_secret_version",
			Category:         "SSM",
			ShortDescription: `Provide a resource to create a SSM secret version.`,
			Description:      ``,
			Keywords: []string{
				"ssm",
				"secret",
				"version",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "secret_name",
					Description: `(Required, ForceNew) Name of secret which cannot be repeated in the same region. The maximum length is 128 bytes. The name can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `(Required, ForceNew) Version of secret. The maximum length is 64 bytes. The version_id can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.`,
				},
				resource.Attribute{
					Name:        "secret_binary",
					Description: `(Optional) The base64-encoded binary secret. secret_binary and secret_string must be set only one, and the maximum support is 4096 bytes. When secret status is ` + "`" + `Disabled` + "`" + `, this field will not update anymore.`,
				},
				resource.Attribute{
					Name:        "secret_string",
					Description: `(Optional) The string text of secret. secret_binary and secret_string must be set only one, and the maximum support is 4096 bytes. When secret status is ` + "`" + `Disabled` + "`" + `, this field will not update anymore. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import SSM secret version can be imported using the secretName#versionId, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_ssm_secret_version.v1 test#v1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import SSM secret version can be imported using the secretName#versionId, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_ssm_secret_version.v1 test#v1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_subnet",
			Category:         "Virtual Private Cloud(VPC)",
			ShortDescription: `Provide a resource to create a VPC subnet.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"subnet",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required, ForceNew) The availability zone within which the subnet should be created.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Required, ForceNew) A network address block of the subnet.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of subnet to be created.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, ForceNew) ID of the VPC to be associated.`,
				},
				resource.Attribute{
					Name:        "is_multicast",
					Description: `(Optional) Indicates whether multicast is enabled. The default value is 'true'.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `(Optional) ID of a routing table to which the subnet should be associated.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of the subnet. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "available_ip_count",
					Description: `The number of available IPs.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of subnet resource.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Indicates whether it is the default VPC for this region. ## Import Vpc subnet instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_subnet.test subnet_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "available_ip_count",
					Description: `The number of available IPs.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of subnet resource.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Indicates whether it is the default VPC for this region. ## Import Vpc subnet instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_subnet.test subnet_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_tcaplus_cluster",
			Category:         "TcaplusDB",
			ShortDescription: `Use this resource to create TcaplusDB cluster.`,
			Description:      ``,
			Keywords: []string{
				"tcaplusdb",
				"tcaplus",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Required) Name of the TcaplusDB cluster. Name length should be between 1 and 30.`,
				},
				resource.Attribute{
					Name:        "idl_type",
					Description: `(Required, ForceNew) IDL type of the TcaplusDB cluster. Valid values: ` + "`" + `PROTO` + "`" + ` and ` + "`" + `TDR` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Password of the TcaplusDB cluster. Password length should be between 12 and 16. The password must be a`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required, ForceNew) Subnet id of the TcaplusDB cluster.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, ForceNew) VPC id of the TcaplusDB cluster.`,
				},
				resource.Attribute{
					Name:        "old_password_expire_last",
					Description: `(Optional) Expiration time of old password after password update, unit: second. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "api_access_id",
					Description: `Access ID of the TcaplusDB cluster.For TcaplusDB SDK connect.`,
				},
				resource.Attribute{
					Name:        "api_access_ip",
					Description: `Access IP of the TcaplusDB cluster.For TcaplusDB SDK connect.`,
				},
				resource.Attribute{
					Name:        "api_access_port",
					Description: `Access port of the TcaplusDB cluster.For TcaplusDB SDK connect.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the TcaplusDB cluster.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `Network type of the TcaplusDB cluster.`,
				},
				resource.Attribute{
					Name:        "old_password_expire_time",
					Description: `Expiration time of the old password. If ` + "`" + `password_status` + "`" + ` is ` + "`" + `unmodifiable` + "`" + `, it means the old password has not yet expired.`,
				},
				resource.Attribute{
					Name:        "password_status",
					Description: `Password status of the TcaplusDB cluster. Valid values: ` + "`" + `unmodifiable` + "`" + `, ` + "`" + `modifiable` + "`" + `. ` + "`" + `unmodifiable` + "`" + `. which means the password can not be changed in this moment; ` + "`" + `modifiable` + "`" + `, which means the password can be changed in this moment. ## Import tcaplus cluster can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_tcaplus_cluster.test 26655801 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "api_access_id",
					Description: `Access ID of the TcaplusDB cluster.For TcaplusDB SDK connect.`,
				},
				resource.Attribute{
					Name:        "api_access_ip",
					Description: `Access IP of the TcaplusDB cluster.For TcaplusDB SDK connect.`,
				},
				resource.Attribute{
					Name:        "api_access_port",
					Description: `Access port of the TcaplusDB cluster.For TcaplusDB SDK connect.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the TcaplusDB cluster.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `Network type of the TcaplusDB cluster.`,
				},
				resource.Attribute{
					Name:        "old_password_expire_time",
					Description: `Expiration time of the old password. If ` + "`" + `password_status` + "`" + ` is ` + "`" + `unmodifiable` + "`" + `, it means the old password has not yet expired.`,
				},
				resource.Attribute{
					Name:        "password_status",
					Description: `Password status of the TcaplusDB cluster. Valid values: ` + "`" + `unmodifiable` + "`" + `, ` + "`" + `modifiable` + "`" + `. ` + "`" + `unmodifiable` + "`" + `. which means the password can not be changed in this moment; ` + "`" + `modifiable` + "`" + `, which means the password can be changed in this moment. ## Import tcaplus cluster can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_tcaplus_cluster.test 26655801 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_tcaplus_idl",
			Category:         "TcaplusDB",
			ShortDescription: `Use this resource to create TcaplusDB IDL file.`,
			Description:      ``,
			Keywords: []string{
				"tcaplusdb",
				"tcaplus",
				"idl",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required, ForceNew) ID of the TcaplusDB cluster to which the table group belongs.`,
				},
				resource.Attribute{
					Name:        "file_content",
					Description: `(Required, ForceNew) IDL file content of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "file_ext_type",
					Description: `(Required, ForceNew) File ext type of the IDL file. If ` + "`" + `file_type` + "`" + ` is ` + "`" + `PROTO` + "`" + `, ` + "`" + `file_ext_type` + "`" + ` must be 'proto'; If ` + "`" + `file_type` + "`" + ` is ` + "`" + `TDR` + "`" + `, ` + "`" + `file_ext_type` + "`" + ` must be 'xml'.`,
				},
				resource.Attribute{
					Name:        "file_name",
					Description: `(Required, ForceNew) Name of the IDL file.`,
				},
				resource.Attribute{
					Name:        "file_type",
					Description: `(Required, ForceNew) Type of the IDL file. Valid values are PROTO and TDR.`,
				},
				resource.Attribute{
					Name:        "tablegroup_id",
					Description: `(Required, ForceNew) ID of the table group to which the IDL file belongs. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "table_infos",
					Description: `Table info of the IDL.`,
				},
				resource.Attribute{
					Name:        "error",
					Description: `Error messages for creating IDL file.`,
				},
				resource.Attribute{
					Name:        "index_key_set",
					Description: `Index key set of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "key_fields",
					Description: `Primary key fields of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "sum_key_field_size",
					Description: `Total size of primary key field of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "sum_value_field_size",
					Description: `Total size of non-primary key fields of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "table_name",
					Description: `Name of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "value_fields",
					Description: `Non-primary key fields of the TcaplusDB table.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "table_infos",
					Description: `Table info of the IDL.`,
				},
				resource.Attribute{
					Name:        "error",
					Description: `Error messages for creating IDL file.`,
				},
				resource.Attribute{
					Name:        "index_key_set",
					Description: `Index key set of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "key_fields",
					Description: `Primary key fields of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "sum_key_field_size",
					Description: `Total size of primary key field of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "sum_value_field_size",
					Description: `Total size of non-primary key fields of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "table_name",
					Description: `Name of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "value_fields",
					Description: `Non-primary key fields of the TcaplusDB table.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_tcaplus_table",
			Category:         "TcaplusDB",
			ShortDescription: `Use this resource to create TcaplusDB table.`,
			Description:      ``,
			Keywords: []string{
				"tcaplusdb",
				"tcaplus",
				"table",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required, ForceNew) ID of the TcaplusDB cluster to which the table belongs.`,
				},
				resource.Attribute{
					Name:        "idl_id",
					Description: `(Required) ID of the IDL File.`,
				},
				resource.Attribute{
					Name:        "reserved_read_cu",
					Description: `(Required, ForceNew) Reserved read capacity units of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "reserved_volume",
					Description: `(Required, ForceNew) Reserved storage capacity of the TcaplusDB table (unit: GB).`,
				},
				resource.Attribute{
					Name:        "reserved_write_cu",
					Description: `(Required, ForceNew) Reserved write capacity units of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "table_idl_type",
					Description: `(Required) IDL type of the TcaplusDB table. Valid values: ` + "`" + `PROTO` + "`" + ` and ` + "`" + `TDR` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "table_name",
					Description: `(Required, ForceNew) Name of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "table_type",
					Description: `(Required, ForceNew) Type of the TcaplusDB table. Valid values are ` + "`" + `GENERIC` + "`" + ` and ` + "`" + `LIST` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tablegroup_id",
					Description: `(Required, ForceNew) ID of the table group to which the table belongs.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the TcaplusDB table. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "error",
					Description: `Error messages for creating TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "table_size",
					Description: `Size of the TcaplusDB table.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "error",
					Description: `Error messages for creating TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "table_size",
					Description: `Size of the TcaplusDB table.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_tcaplus_tablegroup",
			Category:         "TcaplusDB",
			ShortDescription: `Use this resource to create TcaplusDB table group.`,
			Description:      ``,
			Keywords: []string{
				"tcaplusdb",
				"tcaplus",
				"tablegroup",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required, ForceNew) ID of the TcaplusDB cluster to which the table group belongs.`,
				},
				resource.Attribute{
					Name:        "tablegroup_name",
					Description: `(Required) Name of the TcaplusDB table group. Name length should be between 1 and 30. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the TcaplusDB table group.`,
				},
				resource.Attribute{
					Name:        "table_count",
					Description: `Number of tables.`,
				},
				resource.Attribute{
					Name:        "total_size",
					Description: `Total storage size (MB).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the TcaplusDB table group.`,
				},
				resource.Attribute{
					Name:        "table_count",
					Description: `Number of tables.`,
				},
				resource.Attribute{
					Name:        "total_size",
					Description: `Total storage size (MB).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_tcr_instance",
			Category:         "Tencent Container Registry(TCR)",
			ShortDescription: `Use this resource to create tcr instance.`,
			Description:      ``,
			Keywords: []string{
				"tencent",
				"container",
				"registry",
				"tcr",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Required, ForceNew) TCR types. Valid values are: ` + "`" + `standard` + "`" + `, ` + "`" + `basic` + "`" + `, ` + "`" + `premium` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Name of the TCR instance.`,
				},
				resource.Attribute{
					Name:        "delete_bucket",
					Description: `(Optional) Indicate to delete the COS bucket which is auto-created with the instance or not.`,
				},
				resource.Attribute{
					Name:        "open_public_operation",
					Description: `(Optional) Control public network access.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, ForceNew) The available tags within this TCR instance. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "internal_end_point",
					Description: `Internal address for access of the TCR instance.`,
				},
				resource.Attribute{
					Name:        "public_domain",
					Description: `Public address for access of the TCR instance.`,
				},
				resource.Attribute{
					Name:        "public_status",
					Description: `Status of the TCR instance public network access.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the TCR instance. ## Import tcr instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_tcr_instance.foo cls-cda1iex1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "internal_end_point",
					Description: `Internal address for access of the TCR instance.`,
				},
				resource.Attribute{
					Name:        "public_domain",
					Description: `Public address for access of the TCR instance.`,
				},
				resource.Attribute{
					Name:        "public_status",
					Description: `Status of the TCR instance public network access.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the TCR instance. ## Import tcr instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_tcr_instance.foo cls-cda1iex1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_tcr_namespace",
			Category:         "Tencent Container Registry(TCR)",
			ShortDescription: `Use this resource to create tcr namespace.`,
			Description:      ``,
			Keywords: []string{
				"tencent",
				"container",
				"registry",
				"tcr",
				"namespace",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, ForceNew) ID of the TCR instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Name of the TCR namespace. Valid length is [2~30]. It can only contain lowercase letters, numbers and separators (` + "`" + `.` + "`" + `, ` + "`" + `_` + "`" + `, ` + "`" + `-` + "`" + `), and cannot start, end or continue with separators.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `(Optional) Indicate that the namespace is public or not. Default is ` + "`" + `false` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import tcr namespace can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_tcr_namespace.foo cls-cda1iex1#namespace ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import tcr namespace can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_tcr_namespace.foo cls-cda1iex1#namespace ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_tcr_repository",
			Category:         "Tencent Container Registry(TCR)",
			ShortDescription: `Use this resource to create tcr repository.`,
			Description:      ``,
			Keywords: []string{
				"tencent",
				"container",
				"registry",
				"tcr",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "brief_desc",
					Description: `(Required) Brief description of the repository. Valid length is [1~100].`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) Description of the repository. Valid length is [1~1000].`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, ForceNew) ID of the TCR instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Name of the TCR repository. Valid length is [2~200]. It can only contain lowercase letters, numbers and separators (` + "`" + `.` + "`" + `, ` + "`" + `_` + "`" + `, ` + "`" + `-` + "`" + `, ` + "`" + `/` + "`" + `), and cannot start, end or continue with separators. Support the use of multi-level address formats, such as ` + "`" + `sub1/sub2/repo` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "namespace_name",
					Description: `(Required, ForceNew) Name of the TCR namespace. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `Indicate the repository is public or not.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Last updated time.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `URL of the repository. ## Import tcr repository can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_tcr_repository.foo cls-cda1iex1#namespace#repository ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `Indicate the repository is public or not.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Last updated time.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `URL of the repository. ## Import tcr repository can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_tcr_repository.foo cls-cda1iex1#namespace#repository ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_tcr_token",
			Category:         "Tencent Container Registry(TCR)",
			ShortDescription: `Use this resource to create tcr long term token.`,
			Description:      ``,
			Keywords: []string{
				"tencent",
				"container",
				"registry",
				"tcr",
				"token",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, ForceNew) ID of the TCR instance.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, ForceNew) Description of the token. Valid length is [0~255].`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional) Indicate to enable this token or not. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time.`,
				},
				resource.Attribute{
					Name:        "token_id",
					Description: `Sub ID of the TCR token. The full ID of token format like ` + "`" + `instance_id#token_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `The content of the token.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `User name of the token. ## Import tcr token can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_tcr_token.foo cls-cda1iex1#namespace#buv3h3j96j2d1rk1cllg ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time.`,
				},
				resource.Attribute{
					Name:        "token_id",
					Description: `Sub ID of the TCR token. The full ID of token format like ` + "`" + `instance_id#token_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `The content of the token.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `User name of the token. ## Import tcr token can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_tcr_token.foo cls-cda1iex1#namespace#buv3h3j96j2d1rk1cllg ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_tcr_vpc_attachment",
			Category:         "Tencent Container Registry(TCR)",
			ShortDescription: `Use this resource to create tcr vpc attachment to manage access of internal endpoint.`,
			Description:      ``,
			Keywords: []string{
				"tencent",
				"container",
				"registry",
				"tcr",
				"vpc",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, ForceNew) ID of the TCR instance.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required, ForceNew) ID of subnet.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, ForceNew) ID of VPC.`,
				},
				resource.Attribute{
					Name:        "enable_public_domain_dns",
					Description: `(Optional) Whether to enable public domain dns. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enable_vpc_domain_dns",
					Description: `(Optional) Whether to enable vpc domain dns. Default value is ` + "`" + `false` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "access_ip",
					Description: `IP address of the internal access.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the internal access. ## Import tcr vpc attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_tcr_vpc_attachment.foo cls-cda1iex1#vpcAccess ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "access_ip",
					Description: `IP address of the internal access.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the internal access. ## Import tcr vpc attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_tcr_vpc_attachment.foo cls-cda1iex1#vpcAccess ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_vod_adaptive_dynamic_streaming_template",
			Category:         "Video on Demand(VOD)",
			ShortDescription: `Provide a resource to create a VOD adaptive dynamic streaming template.`,
			Description:      ``,
			Keywords: []string{
				"video",
				"on",
				"demand",
				"vod",
				"adaptive",
				"dynamic",
				"streaming",
				"template",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "format",
					Description: `(Required) Adaptive bitstream format. Valid values: ` + "`" + `HLS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Template name. Length limit: 64 characters.`,
				},
				resource.Attribute{
					Name:        "stream_info",
					Description: `(Required) List of AdaptiveStreamTemplate parameter information of output substream for adaptive bitrate streaming. Up to 10 substreams can be output. Note: the frame rate of all substreams must be the same; otherwise, the frame rate of the first substream will be used as the output frame rate.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Template description. Length limit: 256 characters.`,
				},
				resource.Attribute{
					Name:        "disable_higher_video_bitrate",
					Description: `(Optional) Whether to prohibit transcoding video from low bitrate to high bitrate. Valid values: ` + "`" + `false` + "`" + `,` + "`" + `true` + "`" + `. ` + "`" + `false` + "`" + `: no, ` + "`" + `true` + "`" + `: yes. Default value: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disable_higher_video_resolution",
					Description: `(Optional) Whether to prohibit transcoding from low resolution to high resolution. Valid values: ` + "`" + `false` + "`" + `,` + "`" + `true` + "`" + `. ` + "`" + `false` + "`" + `: no, ` + "`" + `true` + "`" + `: yes. Default value: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "drm_type",
					Description: `(Optional, ForceNew) DRM scheme type. Valid values: ` + "`" + `SimpleAES` + "`" + `. If this field is an empty string, DRM will not be performed on the video.`,
				},
				resource.Attribute{
					Name:        "sub_app_id",
					Description: `(Optional) Subapplication ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty. The ` + "`" + `audio` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "bitrate",
					Description: `(Required) Audio stream bitrate in Kbps. Value range: ` + "`" + `0` + "`" + ` and ` + "`" + `[26, 256]` + "`" + `. If the value is ` + "`" + `0` + "`" + `, the bitrate of the audio stream will be the same as that of the original audio.`,
				},
				resource.Attribute{
					Name:        "codec",
					Description: `(Required) Audio stream encoder. Valid value are: ` + "`" + `libfdk_aac` + "`" + ` and ` + "`" + `libmp3lame` + "`" + `. while ` + "`" + `libfdk_aac` + "`" + ` is recommended.`,
				},
				resource.Attribute{
					Name:        "sample_rate",
					Description: `(Required) Audio stream sample rate. Valid values: ` + "`" + `32000` + "`" + `, ` + "`" + `44100` + "`" + `, ` + "`" + `48000` + "`" + `Hz.`,
				},
				resource.Attribute{
					Name:        "audio_channel",
					Description: `(Optional) Audio channel system. Valid values: mono, dual, stereo. Default value: dual. The ` + "`" + `stream_info` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "audio",
					Description: `(Required) Audio parameter information.`,
				},
				resource.Attribute{
					Name:        "video",
					Description: `(Required) Video parameter information.`,
				},
				resource.Attribute{
					Name:        "remove_audio",
					Description: `(Optional) Whether to remove audio stream. Valid values: ` + "`" + `false` + "`" + `: no, ` + "`" + `true` + "`" + `: yes. ` + "`" + `false` + "`" + ` by default. The ` + "`" + `video` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "bitrate",
					Description: `(Required) Bitrate of video stream in Kbps. Value range: ` + "`" + `0` + "`" + ` and ` + "`" + `[128, 35000]` + "`" + `. If the value is ` + "`" + `0` + "`" + `, the bitrate of the video will be the same as that of the source video.`,
				},
				resource.Attribute{
					Name:        "codec",
					Description: `(Required) Video stream encoder. Valid values: ` + "`" + `libx264` + "`" + `,` + "`" + `libx265` + "`" + `,` + "`" + `av1` + "`" + `. ` + "`" + `libx264` + "`" + `: H.264, ` + "`" + `libx265` + "`" + `: H.265, ` + "`" + `av1` + "`" + `: AOMedia Video 1. Currently, a resolution within 640x480 must be specified for ` + "`" + `H.265` + "`" + `. and the ` + "`" + `av1` + "`" + ` container only supports mp4.`,
				},
				resource.Attribute{
					Name:        "fps",
					Description: `(Required) Video frame rate in Hz. Value range: ` + "`" + `[0, 60]` + "`" + `. If the value is ` + "`" + `0` + "`" + `, the frame rate will be the same as that of the source video.`,
				},
				resource.Attribute{
					Name:        "fill_type",
					Description: `(Optional) Fill type. Fill refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported: ` + "`" + `stretch` + "`" + `: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot shorter or longer; ` + "`" + `black` + "`" + `: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks. Default value: black. Note: this field may return null, indicating that no valid values can be obtained.`,
				},
				resource.Attribute{
					Name:        "height",
					Description: `(Optional) Maximum value of the height (or short side) of a video stream in px. Value range: ` + "`" + `0` + "`" + ` and ` + "`" + `[128, 4096]` + "`" + `. If both ` + "`" + `width` + "`" + ` and ` + "`" + `height` + "`" + ` are ` + "`" + `0` + "`" + `, the resolution will be the same as that of the source video; If ` + "`" + `width` + "`" + ` is ` + "`" + `0` + "`" + `, but ` + "`" + `height` + "`" + ` is not ` + "`" + `0` + "`" + `, ` + "`" + `width` + "`" + ` will be proportionally scaled; If ` + "`" + `width` + "`" + ` is not ` + "`" + `0` + "`" + `, but ` + "`" + `height` + "`" + ` is ` + "`" + `0` + "`" + `, ` + "`" + `height` + "`" + ` will be proportionally scaled; If both ` + "`" + `width` + "`" + ` and ` + "`" + `height` + "`" + ` are not ` + "`" + `0` + "`" + `, the custom resolution will be used. Default value: ` + "`" + `0` + "`" + `. Note: this field may return null, indicating that no valid values can be obtained.`,
				},
				resource.Attribute{
					Name:        "resolution_adaptive",
					Description: `(Optional) Resolution adaption. Valid values: ` + "`" + `true` + "`" + `,` + "`" + `false` + "`" + `. ` + "`" + `true` + "`" + `: enabled. In this case, ` + "`" + `width` + "`" + ` represents the long side of a video, while ` + "`" + `height` + "`" + ` the short side; ` + "`" + `false` + "`" + `: disabled. In this case, ` + "`" + `width` + "`" + ` represents the width of a video, while ` + "`" + `height` + "`" + ` the height. Default value: ` + "`" + `true` + "`" + `. Note: this field may return null, indicating that no valid values can be obtained.`,
				},
				resource.Attribute{
					Name:        "width",
					Description: `(Optional) Maximum value of the width (or long side) of a video stream in px. Value range: ` + "`" + `0` + "`" + ` and ` + "`" + `[128, 4096]` + "`" + `. If both ` + "`" + `width` + "`" + ` and ` + "`" + `height` + "`" + ` are ` + "`" + `0` + "`" + `, the resolution will be the same as that of the source video; If ` + "`" + `width` + "`" + ` is ` + "`" + `0` + "`" + `, but ` + "`" + `height` + "`" + ` is not ` + "`" + `0` + "`" + `, ` + "`" + `width` + "`" + ` will be proportionally scaled; If ` + "`" + `width` + "`" + ` is not ` + "`" + `0` + "`" + `, but ` + "`" + `height` + "`" + ` is ` + "`" + `0` + "`" + `, ` + "`" + `height` + "`" + ` will be proportionally scaled; If both ` + "`" + `width` + "`" + ` and ` + "`" + `height` + "`" + ` are not ` + "`" + `0` + "`" + `, the custom resolution will be used. Default value: ` + "`" + `0` + "`" + `. Note: this field may return null, indicating that no valid values can be obtained. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of template in ISO date format.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Last modified time of template in ISO date format. ## Import VOD adaptive dynamic streaming template can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_vod_adaptive_dynamic_streaming_template.foo 169141 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of template in ISO date format.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Last modified time of template in ISO date format. ## Import VOD adaptive dynamic streaming template can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_vod_adaptive_dynamic_streaming_template.foo 169141 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_vod_image_sprite_template",
			Category:         "Video on Demand(VOD)",
			ShortDescription: `Provide a resource to create a VOD image sprite template.`,
			Description:      ``,
			Keywords: []string{
				"video",
				"on",
				"demand",
				"vod",
				"image",
				"sprite",
				"template",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "column_count",
					Description: `(Required) Subimage column count of an image sprite.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of a time point screen capturing template. Length limit: 64 characters.`,
				},
				resource.Attribute{
					Name:        "row_count",
					Description: `(Required) Subimage row count of an image sprite.`,
				},
				resource.Attribute{
					Name:        "sample_interval",
					Description: `(Required) Sampling interval. If ` + "`" + `sample_type` + "`" + ` is ` + "`" + `Percent` + "`" + `, sampling will be performed at an interval of the specified percentage. If ` + "`" + `sample_type` + "`" + ` is ` + "`" + `Time` + "`" + `, sampling will be performed at the specified time interval in seconds.`,
				},
				resource.Attribute{
					Name:        "sample_type",
					Description: `(Required) Sampling type. Valid values: ` + "`" + `Percent` + "`" + `, ` + "`" + `Time` + "`" + `. ` + "`" + `Percent` + "`" + `: by percent. ` + "`" + `Time` + "`" + `: by time interval.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Template description. Length limit: 256 characters.`,
				},
				resource.Attribute{
					Name:        "fill_type",
					Description: `(Optional) Fill refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported: ` + "`" + `stretch` + "`" + `: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot shorter or longer; ` + "`" + `black` + "`" + `: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks. Default value: ` + "`" + `black` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "height",
					Description: `(Optional) Maximum value of the ` + "`" + `height` + "`" + ` (or short side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both ` + "`" + `width` + "`" + ` and ` + "`" + `height` + "`" + ` are ` + "`" + `0` + "`" + `, the resolution will be the same as that of the source video; If ` + "`" + `width` + "`" + ` is ` + "`" + `0` + "`" + `, but ` + "`" + `height` + "`" + ` is not ` + "`" + `0` + "`" + `, ` + "`" + `width` + "`" + ` will be proportionally scaled; If ` + "`" + `width` + "`" + ` is not ` + "`" + `0` + "`" + `, but ` + "`" + `height` + "`" + ` is ` + "`" + `0` + "`" + `, ` + "`" + `height` + "`" + ` will be proportionally scaled; If both ` + "`" + `width` + "`" + ` and ` + "`" + `height` + "`" + ` are not ` + "`" + `0` + "`" + `, the custom resolution will be used. Default value: ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "resolution_adaptive",
					Description: `(Optional) Resolution adaption. Valid values: ` + "`" + `true` + "`" + `,` + "`" + `false` + "`" + `. ` + "`" + `true` + "`" + `: enabled. In this case, ` + "`" + `width` + "`" + ` represents the long side of a video, while ` + "`" + `height` + "`" + ` the short side; ` + "`" + `false` + "`" + `: disabled. In this case, ` + "`" + `width` + "`" + ` represents the width of a video, while ` + "`" + `height` + "`" + ` the height. Default value: ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sub_app_id",
					Description: `(Optional) Subapplication ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.`,
				},
				resource.Attribute{
					Name:        "width",
					Description: `(Optional) Maximum value of the ` + "`" + `width` + "`" + ` (or long side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both ` + "`" + `width` + "`" + ` and ` + "`" + `height` + "`" + ` are ` + "`" + `0` + "`" + `, the resolution will be the same as that of the source video; If ` + "`" + `width` + "`" + ` is ` + "`" + `0` + "`" + `, but ` + "`" + `height` + "`" + ` is not ` + "`" + `0` + "`" + `, width will be proportionally scaled; If ` + "`" + `width` + "`" + ` is not ` + "`" + `0` + "`" + `, but ` + "`" + `height` + "`" + ` is ` + "`" + `0` + "`" + `, ` + "`" + `height` + "`" + ` will be proportionally scaled; If both ` + "`" + `width` + "`" + ` and ` + "`" + `height` + "`" + ` are not ` + "`" + `0` + "`" + `, the custom resolution will be used. Default value: ` + "`" + `0` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of template in ISO date format.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Last modified time of template in ISO date format. ## Import VOD image sprite template can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_vod_image_sprite_template.foo 51156 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of template in ISO date format.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Last modified time of template in ISO date format. ## Import VOD image sprite template can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_vod_image_sprite_template.foo 51156 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_vod_procedure_template",
			Category:         "Video on Demand(VOD)",
			ShortDescription: `Provide a resource to create a VOD procedure template.`,
			Description:      ``,
			Keywords: []string{
				"video",
				"on",
				"demand",
				"vod",
				"procedure",
				"template",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Task flow name (up to 20 characters).`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Template description. Length limit: 256 characters.`,
				},
				resource.Attribute{
					Name:        "media_process_task",
					Description: `(Optional) Parameter of video processing task.`,
				},
				resource.Attribute{
					Name:        "sub_app_id",
					Description: `(Optional) Subapplication ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty. The ` + "`" + `adaptive_dynamic_streaming_task_list` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "definition",
					Description: `(Required) Adaptive bitrate streaming template ID.`,
				},
				resource.Attribute{
					Name:        "watermark_list",
					Description: `(Optional) List of up to ` + "`" + `10` + "`" + ` image or text watermarks. Note: this field may return null, indicating that no valid values can be obtained. The ` + "`" + `animated_graphic_task_list` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "definition",
					Description: `(Required) Animated image generating template ID.`,
				},
				resource.Attribute{
					Name:        "end_time_offset",
					Description: `(Required) End time of animated image in video in seconds.`,
				},
				resource.Attribute{
					Name:        "start_time_offset",
					Description: `(Required) Start time of animated image in video in seconds. The ` + "`" + `cover_by_snapshot_task_list` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "definition",
					Description: `(Required) Time point screen capturing template ID.`,
				},
				resource.Attribute{
					Name:        "position_type",
					Description: `(Required) Screen capturing mode. Valid values: ` + "`" + `Time` + "`" + `, ` + "`" + `Percent` + "`" + `. ` + "`" + `Time` + "`" + `: screen captures by time point, ` + "`" + `Percent` + "`" + `: screen captures by percentage.`,
				},
				resource.Attribute{
					Name:        "position_value",
					Description: `(Required) Screenshot position: For time point screen capturing, this means to take a screenshot at a specified time point (in seconds) and use it as the cover. For percentage screen capturing, this value means to take a screenshot at a specified percentage of the video duration and use it as the cover.`,
				},
				resource.Attribute{
					Name:        "watermark_list",
					Description: `(Optional) List of up to ` + "`" + `10` + "`" + ` image or text watermarks. Note: this field may return null, indicating that no valid values can be obtained. The ` + "`" + `image_sprite_task_list` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "definition",
					Description: `(Required) Image sprite generating template ID. The ` + "`" + `media_process_task` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "adaptive_dynamic_streaming_task_list",
					Description: `(Optional) List of adaptive bitrate streaming tasks. Note: this field may return null, indicating that no valid values can be obtained.`,
				},
				resource.Attribute{
					Name:        "animated_graphic_task_list",
					Description: `(Optional) List of animated image generating tasks. Note: this field may return null, indicating that no valid values can be obtained.`,
				},
				resource.Attribute{
					Name:        "cover_by_snapshot_task_list",
					Description: `(Optional) List of cover generating tasks. Note: this field may return null, indicating that no valid values can be obtained.`,
				},
				resource.Attribute{
					Name:        "image_sprite_task_list",
					Description: `(Optional) List of image sprite generating tasks. Note: this field may return null, indicating that no valid values can be obtained.`,
				},
				resource.Attribute{
					Name:        "sample_snapshot_task_list",
					Description: `(Optional) List of sampled screen capturing tasks. Note: this field may return null, indicating that no valid values can be obtained.`,
				},
				resource.Attribute{
					Name:        "snapshot_by_time_offset_task_list",
					Description: `(Optional) List of time point screen capturing tasks. Note: this field may return null, indicating that no valid values can be obtained.`,
				},
				resource.Attribute{
					Name:        "transcode_task_list",
					Description: `(Optional) List of transcoding tasks. Note: this field may return null, indicating that no valid values can be obtained. The ` + "`" + `mosaic_list` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "coordinate_origin",
					Description: `(Optional) Origin position, which currently can only be: ` + "`" + `TopLeft` + "`" + `: the origin of coordinates is in the top-left corner of the video, and the origin of the blur is in the top-left corner of the image or text. Default value: TopLeft.`,
				},
				resource.Attribute{
					Name:        "end_time_offset",
					Description: `(Optional) End time offset of blur in seconds. If this parameter is left empty or ` + "`" + `0` + "`" + ` is entered, the blur will exist till the last video frame; If this value is greater than ` + "`" + `0` + "`" + ` (e.g., n), the blur will exist till second n; If this value is smaller than ` + "`" + `0` + "`" + ` (e.g., -n), the blur will exist till second n before the last video frame.`,
				},
				resource.Attribute{
					Name:        "height",
					Description: `(Optional) Blur height. ` + "`" + `%` + "`" + ` and ` + "`" + `px` + "`" + ` formats are supported: If the string ends in ` + "`" + `%` + "`" + `, the ` + "`" + `height` + "`" + ` of the blur will be the specified percentage of the video height; for example, 10% means that Height is 10% of the video height; If the string ends in ` + "`" + `px` + "`" + `, the ` + "`" + `height` + "`" + ` of the blur will be in px; for example, 100px means that Height is 100 px. Default value: ` + "`" + `10%` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "start_time_offset",
					Description: `(Optional) Start time offset of blur in seconds. If this parameter is left empty or ` + "`" + `0` + "`" + ` is entered, the blur will appear upon the first video frame. If this parameter is left empty or ` + "`" + `0` + "`" + ` is entered, the blur will appear upon the first video frame; If this value is greater than ` + "`" + `0` + "`" + ` (e.g., n), the blur will appear at second n after the first video frame; If this value is smaller than ` + "`" + `0` + "`" + ` (e.g., -n), the blur will appear at second n before the last video frame.`,
				},
				resource.Attribute{
					Name:        "width",
					Description: `(Optional) Blur width. ` + "`" + `%` + "`" + ` and ` + "`" + `px` + "`" + ` formats are supported: If the string ends in ` + "`" + `%` + "`" + `, the ` + "`" + `width` + "`" + ` of the blur will be the specified percentage of the video width; for example, 10% means that ` + "`" + `width` + "`" + ` is 10% of the video width; If the string ends in ` + "`" + `px` + "`" + `, the ` + "`" + `width` + "`" + ` of the blur will be in px; for example, 100px means that Width is 100 px. Default value: ` + "`" + `10%` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "x_pos",
					Description: `(Optional) The horizontal position of the origin of the blur relative to the origin of coordinates of the video. ` + "`" + `%` + "`" + ` and ` + "`" + `px` + "`" + ` formats are supported: If the string ends in ` + "`" + `%` + "`" + `, the XPos of the blur will be the specified percentage of the video width; for example, 10% means that XPos is 10% of the video width; If the string ends in ` + "`" + `px` + "`" + `, the XPos of the blur will be the specified px; for example, 100px means that XPos is 100 px. Default value: ` + "`" + `0px` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "y_pos",
					Description: `(Optional) Vertical position of the origin of blur relative to the origin of coordinates of video. ` + "`" + `%` + "`" + ` and ` + "`" + `px` + "`" + ` formats are supported: If the string ends in ` + "`" + `%` + "`" + `, the YPos of the blur will be the specified percentage of the video height; for example, 10% means that YPos is 10% of the video height; If the string ends in ` + "`" + `px` + "`" + `, the YPos of the blur will be the specified px; for example, 100px means that YPos is 100 px. Default value: ` + "`" + `0px` + "`" + `. The ` + "`" + `sample_snapshot_task_list` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "definition",
					Description: `(Required) Sampled screen capturing template ID.`,
				},
				resource.Attribute{
					Name:        "watermark_list",
					Description: `(Optional) List of up to ` + "`" + `10` + "`" + ` image or text watermarks. Note: this field may return null, indicating that no valid values can be obtained. The ` + "`" + `snapshot_by_time_offset_task_list` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "definition",
					Description: `(Required) Time point screen capturing template ID.`,
				},
				resource.Attribute{
					Name:        "ext_time_offset_list",
					Description: `(Optional) The list of screenshot time points. ` + "`" + `s` + "`" + ` and ` + "`" + `%` + "`" + ` formats are supported: When a time point string ends with ` + "`" + `s` + "`" + `, its unit is second. For example, ` + "`" + `3.5s` + "`" + ` means the 3.5th second of the video; When a time point string ends with ` + "`" + `%` + "`" + `, it is marked with corresponding percentage of the video duration. For example, ` + "`" + `10%` + "`" + ` means that the time point is at the 10% of the video entire duration.`,
				},
				resource.Attribute{
					Name:        "watermark_list",
					Description: `(Optional) List of up to ` + "`" + `10` + "`" + ` image or text watermarks. Note: this field may return null, indicating that no valid values can be obtained. The ` + "`" + `transcode_task_list` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "definition",
					Description: `(Required) Video transcoding template ID.`,
				},
				resource.Attribute{
					Name:        "mosaic_list",
					Description: `(Optional) List of blurs. Up to 10 ones can be supported.`,
				},
				resource.Attribute{
					Name:        "watermark_list",
					Description: `(Optional) List of up to ` + "`" + `10` + "`" + ` image or text watermarks. Note: this field may return null, indicating that no valid values can be obtained. The ` + "`" + `watermark_list` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "definition",
					Description: `(Required) Watermarking template ID.`,
				},
				resource.Attribute{
					Name:        "end_time_offset",
					Description: `(Optional) End time offset of a watermark in seconds. If this parameter is left blank or ` + "`" + `0` + "`" + ` is entered, the watermark will exist till the last video frame; If this value is greater than ` + "`" + `0` + "`" + ` (e.g., n), the watermark will exist till second n; If this value is smaller than ` + "`" + `0` + "`" + ` (e.g., -n), the watermark will exist till second n before the last video frame.`,
				},
				resource.Attribute{
					Name:        "start_time_offset",
					Description: `(Optional) Start time offset of a watermark in seconds. If this parameter is left blank or ` + "`" + `0` + "`" + ` is entered, the watermark will appear upon the first video frame. If this parameter is left blank or ` + "`" + `0` + "`" + ` is entered, the watermark will appear upon the first video frame; If this value is greater than ` + "`" + `0` + "`" + ` (e.g., n), the watermark will appear at second n after the first video frame; If this value is smaller than ` + "`" + `0` + "`" + ` (e.g., -n), the watermark will appear at second n before the last video frame.`,
				},
				resource.Attribute{
					Name:        "svg_content",
					Description: `(Optional) SVG content of up to ` + "`" + `2000000` + "`" + ` characters. This needs to be entered only when the watermark type is ` + "`" + `SVG` + "`" + `. Note: this field may return null, indicating that no valid values can be obtained.`,
				},
				resource.Attribute{
					Name:        "text_content",
					Description: `(Optional) Text content of up to ` + "`" + `100` + "`" + ` characters. This needs to be entered only when the watermark type is text. Note: this field may return null, indicating that no valid values can be obtained. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of template in ISO date format.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Last modified time of template in ISO date format. ## Import VOD procedure template can be imported using the name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_vod_procedure_template.foo tf-procedure ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of template in ISO date format.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Last modified time of template in ISO date format. ## Import VOD procedure template can be imported using the name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_vod_procedure_template.foo tf-procedure ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_vod_snapshot_by_time_offset_template",
			Category:         "Video on Demand(VOD)",
			ShortDescription: `Provide a resource to create a VOD snapshot by time offset template.`,
			Description:      ``,
			Keywords: []string{
				"video",
				"on",
				"demand",
				"vod",
				"snapshot",
				"by",
				"time",
				"offset",
				"template",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of a time point screen capturing template. Length limit: 64 characters.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Template description. Length limit: 256 characters.`,
				},
				resource.Attribute{
					Name:        "fill_type",
					Description: `(Optional) Fill refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported: ` + "`" + `stretch` + "`" + `: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot ` + "`" + `shorter` + "`" + ` or ` + "`" + `longer` + "`" + `; ` + "`" + `black` + "`" + `: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks. ` + "`" + `white` + "`" + `: fill with white. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with white color blocks. ` + "`" + `gauss` + "`" + `: fill with Gaussian blur. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with Gaussian blur. Default value: ` + "`" + `black` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Optional) Image format. Valid values: ` + "`" + `jpg` + "`" + `, ` + "`" + `png` + "`" + `. Default value: ` + "`" + `jpg` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "height",
					Description: `(Optional) Maximum value of the ` + "`" + `height` + "`" + ` (or short side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both ` + "`" + `width` + "`" + ` and ` + "`" + `height` + "`" + ` are ` + "`" + `0` + "`" + `, the resolution will be the same as that of the source video; If ` + "`" + `width` + "`" + ` is ` + "`" + `0` + "`" + `, but ` + "`" + `height` + "`" + ` is not ` + "`" + `0` + "`" + `, ` + "`" + `width` + "`" + ` will be proportionally scaled; If ` + "`" + `width` + "`" + ` is not ` + "`" + `0` + "`" + `, but ` + "`" + `height` + "`" + ` is ` + "`" + `0` + "`" + `, ` + "`" + `height` + "`" + ` will be proportionally scaled; If both ` + "`" + `width` + "`" + ` and ` + "`" + `height` + "`" + ` are not ` + "`" + `0` + "`" + `, the custom resolution will be used. Default value: ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "resolution_adaptive",
					Description: `(Optional) Resolution adaption. Valid values: ` + "`" + `true` + "`" + `,` + "`" + `false` + "`" + `. ` + "`" + `true` + "`" + `: enabled. In this case, ` + "`" + `width` + "`" + ` represents the long side of a video, while ` + "`" + `height` + "`" + ` the short side; ` + "`" + `false` + "`" + `: disabled. In this case, ` + "`" + `width` + "`" + ` represents the width of a video, while ` + "`" + `height` + "`" + ` the height. Default value: ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sub_app_id",
					Description: `(Optional) Subapplication ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.`,
				},
				resource.Attribute{
					Name:        "width",
					Description: `(Optional) Maximum value of the ` + "`" + `width` + "`" + ` (or long side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both ` + "`" + `width` + "`" + ` and ` + "`" + `height` + "`" + ` are ` + "`" + `0` + "`" + `, the resolution will be the same as that of the source video; If ` + "`" + `width` + "`" + ` is ` + "`" + `0` + "`" + `, but ` + "`" + `height` + "`" + ` is not ` + "`" + `0` + "`" + `, width will be proportionally scaled; If ` + "`" + `width` + "`" + ` is not ` + "`" + `0` + "`" + `, but ` + "`" + `height` + "`" + ` is ` + "`" + `0` + "`" + `, ` + "`" + `height` + "`" + ` will be proportionally scaled; If both ` + "`" + `width` + "`" + ` and ` + "`" + `height` + "`" + ` are not ` + "`" + `0` + "`" + `, the custom resolution will be used. Default value: ` + "`" + `0` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of template in ISO date format.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Last modified time of template in ISO date format. ## Import VOD snapshot by time offset template can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_vod_snapshot_by_time_offset_template.foo 46906 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of template in ISO date format.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Last modified time of template in ISO date format. ## Import VOD snapshot by time offset template can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_vod_snapshot_by_time_offset_template.foo 46906 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_vod_super_player_config",
			Category:         "Video on Demand(VOD)",
			ShortDescription: `Provide a resource to create a VOD super player config.`,
			Description:      ``,
			Keywords: []string{
				"video",
				"on",
				"demand",
				"vod",
				"super",
				"player",
				"config",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Player configuration name, which can contain up to 64 letters, digits, underscores, and hyphens (such as test_ABC-123) and must be unique under a user.`,
				},
				resource.Attribute{
					Name:        "adaptive_dynamic_streaming_definition",
					Description: `(Optional) ID of the unencrypted adaptive bitrate streaming template that allows output, which is required if ` + "`" + `drm_switch` + "`" + ` is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Template description. Length limit: 256 characters.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional) Domain name used for playback. If it is left empty or set to ` + "`" + `Default` + "`" + `, the domain name configured in [Default Distribution Configuration](https://cloud.tencent.com/document/product/266/33373) will be used. ` + "`" + `Default` + "`" + ` by default.`,
				},
				resource.Attribute{
					Name:        "drm_streaming_info",
					Description: `(Optional) Content of the DRM-protected adaptive bitrate streaming template that allows output, which is required if ` + "`" + `drm_switch` + "`" + ` is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "drm_switch",
					Description: `(Optional) Switch of DRM-protected adaptive bitstream playback: ` + "`" + `true` + "`" + `: enabled, indicating to play back only output adaptive bitstreams protected by DRM; ` + "`" + `false` + "`" + `: disabled, indicating to play back unencrypted output adaptive bitstreams. Default value: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "image_sprite_definition",
					Description: `(Optional) ID of the image sprite template that allows output.`,
				},
				resource.Attribute{
					Name:        "resolution_names",
					Description: `(Optional) Display name of player for substreams with different resolutions. If this parameter is left empty or an empty array, the default configuration will be used: ` + "`" + `min_edge_length: 240, name: LD` + "`" + `; ` + "`" + `min_edge_length: 480, name: SD` + "`" + `; ` + "`" + `min_edge_length: 720, name: HD` + "`" + `; ` + "`" + `min_edge_length: 1080, name: FHD` + "`" + `; ` + "`" + `min_edge_length: 1440, name: 2K` + "`" + `; ` + "`" + `min_edge_length: 2160, name: 4K` + "`" + `; ` + "`" + `min_edge_length: 4320, name: 8K` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scheme",
					Description: `(Optional) Scheme used for playback. If it is left empty or set to ` + "`" + `Default` + "`" + `, the scheme configured in [Default Distribution Configuration](https://cloud.tencent.com/document/product/266/33373) will be used. Other valid values: ` + "`" + `HTTP` + "`" + `; ` + "`" + `HTTPS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sub_app_id",
					Description: `(Optional) Subapplication ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty. The ` + "`" + `drm_streaming_info` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "simple_aes_definition",
					Description: `(Optional) ID of the adaptive dynamic streaming template whose protection type is ` + "`" + `SimpleAES` + "`" + `. The ` + "`" + `resolution_names` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "min_edge_length",
					Description: `(Required) Length of video short side in px.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Display name. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of template in ISO date format.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Last modified time of template in ISO date format. ## Import VOD super player config can be imported using the name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_vod_super_player_config.foo tf-super-player ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of template in ISO date format.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Last modified time of template in ISO date format. ## Import VOD super player config can be imported using the name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_vod_super_player_config.foo tf-super-player ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_vpc",
			Category:         "Virtual Private Cloud(VPC)",
			ShortDescription: `Provide a resource to create a VPC.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Required, ForceNew) A network address block which should be a subnet of the three internal network segments (10.0.0.0/16, 172.16.0.0/12 and 192.168.0.0/16).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the VPC.`,
				},
				resource.Attribute{
					Name:        "dns_servers",
					Description: `(Optional) The DNS server list of the VPC. And you can specify 0 to 5 servers to this list.`,
				},
				resource.Attribute{
					Name:        "is_multicast",
					Description: `(Optional) Indicates whether VPC multicast is enabled. The default value is 'true'.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of the VPC. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of VPC.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Indicates whether it is the default VPC for this region. ## Import Vpc instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_vpc.test vpc-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of VPC.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Indicates whether it is the default VPC for this region. ## Import Vpc instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_vpc.test vpc-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_vpc_acl",
			Category:         "Virtual Private Cloud(VPC)",
			ShortDescription: `Provide a resource to create a VPC ACL instance.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"acl",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the network ACL.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) ID of the VPC instance.`,
				},
				resource.Attribute{
					Name:        "egress",
					Description: `(Optional) Egress rules. A rule must match the following format: [action]#[cidr_ip]#[port]#[protocol]. The available value of 'action' is ` + "`" + `ACCEPT` + "`" + ` and ` + "`" + `DROP` + "`" + `. The 'cidr_ip' must be an IP address network or segment. The 'port' valid format is ` + "`" + `80` + "`" + `, ` + "`" + `80,443` + "`" + `, ` + "`" + `80-90` + "`" + ` or ` + "`" + `ALL` + "`" + `. The available value of 'protocol' is ` + "`" + `TCP` + "`" + `, ` + "`" + `UDP` + "`" + `, ` + "`" + `ICMP` + "`" + ` and ` + "`" + `ALL` + "`" + `. When 'protocol' is ` + "`" + `ICMP` + "`" + ` or ` + "`" + `ALL` + "`" + `, the 'port' must be ` + "`" + `ALL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ingress",
					Description: `(Optional) Ingress rules. A rule must match the following format: [action]#[cidr_ip]#[port]#[protocol]. The available value of 'action' is ` + "`" + `ACCEPT` + "`" + ` and ` + "`" + `DROP` + "`" + `. The 'cidr_ip' must be an IP address network or segment. The 'port' valid format is ` + "`" + `80` + "`" + `, ` + "`" + `80,443` + "`" + `, ` + "`" + `80-90` + "`" + ` or ` + "`" + `ALL` + "`" + `. The available value of 'protocol' is ` + "`" + `TCP` + "`" + `, ` + "`" + `UDP` + "`" + `, ` + "`" + `ICMP` + "`" + ` and ` + "`" + `ALL` + "`" + `. When 'protocol' is ` + "`" + `ICMP` + "`" + ` or ` + "`" + `ALL` + "`" + `, the 'port' must be ` + "`" + `ALL` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of ACL. ## Import Vpc ACL can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_vpc_acl.default acl-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of ACL. ## Import Vpc ACL can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_vpc_acl.default acl-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_vpc_acl_attachment",
			Category:         "Virtual Private Cloud(VPC)",
			ShortDescription: `Provide a resource to attach an existing subnet to Network ACL.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"acl",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "acl_id",
					Description: `(Required, ForceNew) ID of the attached ACL.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required, ForceNew) The Subnet instance ID. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import Acl attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_vpc_acl_attachment.attachment acl-eotx5qsg#subnet-91x0geu6 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import Acl attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_vpc_acl_attachment.attachment acl-eotx5qsg#subnet-91x0geu6 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_vpn_connection",
			Category:         "VPN",
			ShortDescription: `Provides a resource to create a VPN connection.`,
			Description:      ``,
			Keywords: []string{
				"vpn",
				"connection",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "customer_gateway_id",
					Description: `(Required, ForceNew) ID of the customer gateway.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the VPN connection. The length of character is limited to 1-60.`,
				},
				resource.Attribute{
					Name:        "pre_share_key",
					Description: `(Required) Pre-shared key of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "security_group_policy",
					Description: `(Required) Security group policy of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "vpn_gateway_id",
					Description: `(Required, ForceNew) ID of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "ike_dh_group_name",
					Description: `(Optional) DH group name of the IKE operation specification. Valid values: ` + "`" + `GROUP1` + "`" + `, ` + "`" + `GROUP2` + "`" + `, ` + "`" + `GROUP5` + "`" + `, ` + "`" + `GROUP14` + "`" + `, ` + "`" + `GROUP24` + "`" + `. Default value is ` + "`" + `GROUP1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ike_exchange_mode",
					Description: `(Optional) Exchange mode of the IKE operation specification. Valid values: ` + "`" + `AGGRESSIVE` + "`" + `, ` + "`" + `MAIN` + "`" + `. Default value is ` + "`" + `MAIN` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ike_local_address",
					Description: `(Optional) Local address of IKE operation specification, valid when ike_local_identity is ` + "`" + `ADDRESS` + "`" + `, generally the value is ` + "`" + `public_ip_address` + "`" + ` of the related VPN gateway.`,
				},
				resource.Attribute{
					Name:        "ike_local_fqdn_name",
					Description: `(Optional) Local FQDN name of the IKE operation specification.`,
				},
				resource.Attribute{
					Name:        "ike_local_identity",
					Description: `(Optional) Local identity way of IKE operation specification. Valid values: ` + "`" + `ADDRESS` + "`" + `, ` + "`" + `FQDN` + "`" + `. Default value is ` + "`" + `ADDRESS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ike_proto_authen_algorithm",
					Description: `(Optional) Proto authenticate algorithm of the IKE operation specification. Valid values: ` + "`" + `MD5` + "`" + `, ` + "`" + `SHA` + "`" + `. Default Value is ` + "`" + `MD5` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ike_proto_encry_algorithm",
					Description: `(Optional) Proto encrypt algorithm of the IKE operation specification. Valid values: ` + "`" + `3DES-CBC` + "`" + `, ` + "`" + `AES-CBC-128` + "`" + `, ` + "`" + `AES-CBC-128` + "`" + `, ` + "`" + `AES-CBC-256` + "`" + `, ` + "`" + `DES-CBC` + "`" + `. Default value is ` + "`" + `3DES-CBC` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ike_remote_address",
					Description: `(Optional) Remote address of IKE operation specification, valid when ike_remote_identity is ` + "`" + `ADDRESS` + "`" + `, generally the value is ` + "`" + `public_ip_address` + "`" + ` of the related customer gateway.`,
				},
				resource.Attribute{
					Name:        "ike_remote_fqdn_name",
					Description: `(Optional) Remote FQDN name of the IKE operation specification.`,
				},
				resource.Attribute{
					Name:        "ike_remote_identity",
					Description: `(Optional) Remote identity way of IKE operation specification. Valid values: ` + "`" + `ADDRESS` + "`" + `, ` + "`" + `FQDN` + "`" + `. Default value is ` + "`" + `ADDRESS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ike_sa_lifetime_seconds",
					Description: `(Optional) SA lifetime of the IKE operation specification, unit is ` + "`" + `second` + "`" + `. The value ranges from 60 to 604800. Default value is 86400 seconds.`,
				},
				resource.Attribute{
					Name:        "ike_version",
					Description: `(Optional) Version of the IKE operation specification. Default value is ` + "`" + `IKEV1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ipsec_encrypt_algorithm",
					Description: `(Optional) Encrypt algorithm of the IPSEC operation specification. Valid values: ` + "`" + `3DES-CBC` + "`" + `, ` + "`" + `AES-CBC-128` + "`" + `, ` + "`" + `AES-CBC-128` + "`" + `, ` + "`" + `AES-CBC-256` + "`" + `, ` + "`" + `DES-CBC` + "`" + `. Default value is ` + "`" + `3DES-CBC` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ipsec_integrity_algorithm",
					Description: `(Optional) Integrity algorithm of the IPSEC operation specification. Valid values: ` + "`" + `SHA1` + "`" + `, ` + "`" + `MD5` + "`" + `. Default value is ` + "`" + `MD5` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ipsec_pfs_dh_group",
					Description: `(Optional) PFS DH group. Valid value: ` + "`" + `GROUP1` + "`" + `, ` + "`" + `GROUP2` + "`" + `, ` + "`" + `GROUP5` + "`" + `, ` + "`" + `GROUP14` + "`" + `, ` + "`" + `GROUP24` + "`" + `, ` + "`" + `NULL` + "`" + `. Default value is ` + "`" + `NULL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ipsec_sa_lifetime_seconds",
					Description: `(Optional) SA lifetime of the IPSEC operation specification, unit is second. Valid value ranges: [180~604800]. Default value is 3600 seconds.`,
				},
				resource.Attribute{
					Name:        "ipsec_sa_lifetime_traffic",
					Description: `(Optional) SA lifetime of the IPSEC operation specification, unit is KB. The value should not be less then 2560. Default value is 1843200.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A list of tags used to associate different resources.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional, ForceNew) ID of the VPC. Required if vpn gateway is not in ` + "`" + `CCN` + "`" + ` type, and doesn't make sense for ` + "`" + `CCN` + "`" + ` vpn gateway. The ` + "`" + `security_group_policy` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "local_cidr_block",
					Description: `(Required) Local cidr block.`,
				},
				resource.Attribute{
					Name:        "remote_cidr_block",
					Description: `(Required) Remote cidr block list. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "encrypt_proto",
					Description: `Encrypt proto of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "is_ccn_type",
					Description: `Indicate whether is ccn type. Modification of this field only impacts force new logic of ` + "`" + `vpc_id` + "`" + `. If ` + "`" + `is_ccn_type` + "`" + ` is true, modification of ` + "`" + `vpc_id` + "`" + ` will be ignored.`,
				},
				resource.Attribute{
					Name:        "net_status",
					Description: `Net status of the VPN connection. Valid value: ` + "`" + `AVAILABLE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "route_type",
					Description: `Route type of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the connection. Valid value: ` + "`" + `PENDING` + "`" + `, ` + "`" + `AVAILABLE` + "`" + `, ` + "`" + `DELETING` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vpn_proto",
					Description: `Vpn proto of the VPN connection. ## Import VPN connection can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_vpn_connection.foo vpnx-nadifg3s ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "encrypt_proto",
					Description: `Encrypt proto of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "is_ccn_type",
					Description: `Indicate whether is ccn type. Modification of this field only impacts force new logic of ` + "`" + `vpc_id` + "`" + `. If ` + "`" + `is_ccn_type` + "`" + ` is true, modification of ` + "`" + `vpc_id` + "`" + ` will be ignored.`,
				},
				resource.Attribute{
					Name:        "net_status",
					Description: `Net status of the VPN connection. Valid value: ` + "`" + `AVAILABLE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "route_type",
					Description: `Route type of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the connection. Valid value: ` + "`" + `PENDING` + "`" + `, ` + "`" + `AVAILABLE` + "`" + `, ` + "`" + `DELETING` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vpn_proto",
					Description: `Vpn proto of the VPN connection. ## Import VPN connection can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_vpn_connection.foo vpnx-nadifg3s ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_vpn_customer_gateway",
			Category:         "VPN",
			ShortDescription: `Provides a resource to create a VPN customer gateway.`,
			Description:      ``,
			Keywords: []string{
				"vpn",
				"customer",
				"gateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the customer gateway. The length of character is limited to 1-60.`,
				},
				resource.Attribute{
					Name:        "public_ip_address",
					Description: `(Required, ForceNew) Public IP of the customer gateway.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A list of tags used to associate different resources. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the customer gateway. ## Import VPN customer gateway can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_vpn_customer_gateway.foo cgw-xfqag ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the customer gateway. ## Import VPN customer gateway can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_vpn_customer_gateway.foo cgw-xfqag ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_vpn_gateway",
			Category:         "VPN",
			ShortDescription: `Provides a resource to create a VPN gateway.`,
			Description:      ``,
			Keywords: []string{
				"vpn",
				"gateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the VPN gateway. The length of character is limited to 1-60.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Required, ForceNew) Zone of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Optional) The maximum public network output bandwidth of VPN gateway (unit: Mbps), the available values include: 5,10,20,50,100,200,500,1000. Default is 5. When charge type is ` + "`" + `PREPAID` + "`" + `, bandwidth degradation operation is unsupported.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `(Optional) Charge Type of the VPN gateway. Valid value: ` + "`" + `PREPAID` + "`" + `, ` + "`" + `POSTPAID_BY_HOUR` + "`" + `. The default is ` + "`" + `POSTPAID_BY_HOUR` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "prepaid_period",
					Description: `(Optional) Period of instance to be prepaid. Valid value: ` + "`" + `1` + "`" + `, ` + "`" + `2` + "`" + `, ` + "`" + `3` + "`" + `, ` + "`" + `4` + "`" + `, ` + "`" + `6` + "`" + `, ` + "`" + `7` + "`" + `, ` + "`" + `8` + "`" + `, ` + "`" + `9` + "`" + `, ` + "`" + `12` + "`" + `, ` + "`" + `24` + "`" + `, ` + "`" + `36` + "`" + `. The unit is month. Caution: when this para and renew_flag para are valid, the request means to renew several months more pre-paid period. This para can only be set to take effect in create operation.`,
				},
				resource.Attribute{
					Name:        "prepaid_renew_flag",
					Description: `(Optional) Flag indicates whether to renew or not. Valid value: ` + "`" + `NOTIFY_AND_RENEW` + "`" + `, ` + "`" + `NOTIFY_AND_AUTO_RENEW` + "`" + `, ` + "`" + `NOT_NOTIFY_AND_NOT_RENEW` + "`" + `. This para can only be set to take effect in create operation.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A list of tags used to associate different resources.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type of gateway instance. Valid value: ` + "`" + `IPSEC` + "`" + `, ` + "`" + `SSL` + "`" + ` and ` + "`" + `CCN` + "`" + `. Note: CCN type is only for whitelist customer now.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional, ForceNew) ID of the VPC. Required if vpn gateway is not in ` + "`" + `CCN` + "`" + ` type, and doesn't make sense for ` + "`" + `CCN` + "`" + ` vpn gateway. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "expired_time",
					Description: `Expired time of the VPN gateway when charge type is ` + "`" + `PREPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "is_address_blocked",
					Description: `Indicates whether ip address is blocked.`,
				},
				resource.Attribute{
					Name:        "new_purchase_plan",
					Description: `The plan of new purchase. Valid value: ` + "`" + `PREPAID_TO_POSTPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "public_ip_address",
					Description: `Public IP of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "restrict_state",
					Description: `Restrict state of gateway. Valid value: ` + "`" + `PRETECIVELY_ISOLATED` + "`" + `, ` + "`" + `NORMAL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the VPN gateway. Valid value: ` + "`" + `PENDING` + "`" + `, ` + "`" + `DELETING` + "`" + `, ` + "`" + `AVAILABLE` + "`" + `. ## Import VPN gateway can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_vpn_gateway.foo vpngw-8ccsnclt ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "expired_time",
					Description: `Expired time of the VPN gateway when charge type is ` + "`" + `PREPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "is_address_blocked",
					Description: `Indicates whether ip address is blocked.`,
				},
				resource.Attribute{
					Name:        "new_purchase_plan",
					Description: `The plan of new purchase. Valid value: ` + "`" + `PREPAID_TO_POSTPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "public_ip_address",
					Description: `Public IP of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "restrict_state",
					Description: `Restrict state of gateway. Valid value: ` + "`" + `PRETECIVELY_ISOLATED` + "`" + `, ` + "`" + `NORMAL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the VPN gateway. Valid value: ` + "`" + `PENDING` + "`" + `, ` + "`" + `DELETING` + "`" + `, ` + "`" + `AVAILABLE` + "`" + `. ## Import VPN gateway can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_vpn_gateway.foo vpngw-8ccsnclt ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"tencentcloud_address_template":                        0,
		"tencentcloud_address_template_group":                  1,
		"tencentcloud_alb_server_attachment":                   2,
		"tencentcloud_api_gateway_api":                         3,
		"tencentcloud_api_gateway_api_key":                     4,
		"tencentcloud_api_gateway_api_key_attachment":          5,
		"tencentcloud_api_gateway_custom_domain":               6,
		"tencentcloud_api_gateway_ip_strategy":                 7,
		"tencentcloud_api_gateway_service":                     8,
		"tencentcloud_api_gateway_service_release":             9,
		"tencentcloud_api_gateway_strategy_attachment":         10,
		"tencentcloud_api_gateway_usage_plan":                  11,
		"tencentcloud_api_gateway_usage_plan_attachment":       12,
		"tencentcloud_as_attachment":                           13,
		"tencentcloud_as_lifecycle_hook":                       14,
		"tencentcloud_as_notification":                         15,
		"tencentcloud_as_scaling_config":                       16,
		"tencentcloud_as_scaling_group":                        17,
		"tencentcloud_as_scaling_policy":                       18,
		"tencentcloud_as_schedule":                             19,
		"tencentcloud_audit":                                   20,
		"tencentcloud_cam_group":                               21,
		"tencentcloud_cam_group_membership":                    22,
		"tencentcloud_cam_group_policy_attachment":             23,
		"tencentcloud_cam_policy":                              24,
		"tencentcloud_cam_role":                                25,
		"tencentcloud_cam_role_policy_attachment":              26,
		"tencentcloud_cam_saml_provider":                       27,
		"tencentcloud_cam_user":                                28,
		"tencentcloud_cam_user_policy_attachment":              29,
		"tencentcloud_cbs_snapshot":                            30,
		"tencentcloud_cbs_snapshot_policy":                     31,
		"tencentcloud_cbs_snapshot_policy_attachment":          32,
		"tencentcloud_cbs_storage":                             33,
		"tencentcloud_cbs_storage_attachment":                  34,
		"tencentcloud_ccn":                                     35,
		"tencentcloud_ccn_attachment":                          36,
		"tencentcloud_ccn_bandwidth_limit":                     37,
		"tencentcloud_cdn_domain":                              38,
		"tencentcloud_cfs_access_group":                        39,
		"tencentcloud_cfs_access_rule":                         40,
		"tencentcloud_cfs_file_system":                         41,
		"tencentcloud_ckafka_acl":                              42,
		"tencentcloud_ckafka_topic":                            43,
		"tencentcloud_ckafka_user":                             44,
		"tencentcloud_clb_attachment":                          45,
		"tencentcloud_clb_instance":                            46,
		"tencentcloud_clb_listener":                            47,
		"tencentcloud_clb_listener_rule":                       48,
		"tencentcloud_clb_redirection":                         49,
		"tencentcloud_clb_target_group":                        50,
		"tencentcloud_clb_target_group_attachment":             51,
		"tencentcloud_clb_target_group_instance_attachment":    52,
		"tencentcloud_container_cluster":                       53,
		"tencentcloud_container_cluster_instance":              54,
		"tencentcloud_cos_bucket":                              55,
		"tencentcloud_cos_bucket_object":                       56,
		"tencentcloud_cos_bucket_policy":                       57,
		"tencentcloud_cynosdb_cluster":                         58,
		"tencentcloud_cynosdb_readonly_instance":               59,
		"tencentcloud_dayu_cc_http_policy":                     60,
		"tencentcloud_dayu_cc_https_policy":                    61,
		"tencentcloud_dayu_ddos_policy":                        62,
		"tencentcloud_dayu_ddos_policy_attachment":             63,
		"tencentcloud_dayu_ddos_policy_case":                   64,
		"tencentcloud_dayu_l4_rule":                            65,
		"tencentcloud_dayu_l7_rule":                            66,
		"tencentcloud_dc_gateway":                              67,
		"tencentcloud_dc_gateway_ccn_route":                    68,
		"tencentcloud_dcx":                                     69,
		"tencentcloud_dnat":                                    70,
		"tencentcloud_eip":                                     71,
		"tencentcloud_eip_association":                         72,
		"tencentcloud_elasticsearch_instance":                  73,
		"tencentcloud_eni":                                     74,
		"tencentcloud_eni_attachment":                          75,
		"tencentcloud_gaap_certificate":                        76,
		"tencentcloud_gaap_domain_error_page":                  77,
		"tencentcloud_gaap_http_domain":                        78,
		"tencentcloud_gaap_http_rule":                          79,
		"tencentcloud_gaap_layer4_listener":                    80,
		"tencentcloud_gaap_layer7_listener":                    81,
		"tencentcloud_gaap_proxy":                              82,
		"tencentcloud_gaap_realserver":                         83,
		"tencentcloud_gaap_security_policy":                    84,
		"tencentcloud_gaap_security_rule":                      85,
		"tencentcloud_ha_vip":                                  86,
		"tencentcloud_ha_vip_eip_attachment":                   87,
		"tencentcloud_image":                                   88,
		"tencentcloud_instance":                                89,
		"tencentcloud_key_pair":                                90,
		"tencentcloud_kms_external_key":                        91,
		"tencentcloud_kms_key":                                 92,
		"tencentcloud_kubernetes_as_scaling_group":             93,
		"tencentcloud_kubernetes_cluster":                      94,
		"tencentcloud_kubernetes_cluster_attachment":           95,
		"tencentcloud_kubernetes_node_pool":                    96,
		"tencentcloud_kubernetes_scale_worker":                 97,
		"tencentcloud_lb":                                      98,
		"tencentcloud_mongodb_instance":                        99,
		"tencentcloud_mongodb_sharding_instance":               100,
		"tencentcloud_mongodb_standby_instance":                101,
		"tencentcloud_monitor_binding_object":                  102,
		"tencentcloud_monitor_binding_receiver":                103,
		"tencentcloud_monitor_policy_group":                    104,
		"tencentcloud_mysql_account":                           105,
		"tencentcloud_mysql_account_privilege":                 106,
		"tencentcloud_mysql_backup_policy":                     107,
		"tencentcloud_mysql_instance":                          108,
		"tencentcloud_mysql_privilege":                         109,
		"tencentcloud_mysql_readonly_instance":                 110,
		"tencentcloud_nat_gateway":                             111,
		"tencentcloud_placement_group":                         112,
		"tencentcloud_postgresql_instance":                     113,
		"tencentcloud_protocol_template":                       114,
		"tencentcloud_protocol_template_group":                 115,
		"tencentcloud_redis_backup_config":                     116,
		"tencentcloud_redis_instance":                          117,
		"tencentcloud_reserved_instance":                       118,
		"tencentcloud_route_entry":                             119,
		"tencentcloud_route_table":                             120,
		"tencentcloud_route_table_entry":                       121,
		"tencentcloud_scf_function":                            122,
		"tencentcloud_scf_namespace":                           123,
		"tencentcloud_security_group":                          124,
		"tencentcloud_security_group_lite_rule":                125,
		"tencentcloud_security_group_rule":                     126,
		"tencentcloud_sqlserver_account":                       127,
		"tencentcloud_sqlserver_account_db_attachment":         128,
		"tencentcloud_sqlserver_basic_instance":                129,
		"tencentcloud_sqlserver_db":                            130,
		"tencentcloud_sqlserver_instance":                      131,
		"tencentcloud_sqlserver_publish_subscribe":             132,
		"tencentcloud_sqlserver_readonly_instance":             133,
		"tencentcloud_ssl_certificate":                         134,
		"tencentcloud_ssl_pay_certificate":                     135,
		"tencentcloud_ssm_secret":                              136,
		"tencentcloud_ssm_secret_version":                      137,
		"tencentcloud_subnet":                                  138,
		"tencentcloud_tcaplus_cluster":                         139,
		"tencentcloud_tcaplus_idl":                             140,
		"tencentcloud_tcaplus_table":                           141,
		"tencentcloud_tcaplus_tablegroup":                      142,
		"tencentcloud_tcr_instance":                            143,
		"tencentcloud_tcr_namespace":                           144,
		"tencentcloud_tcr_repository":                          145,
		"tencentcloud_tcr_token":                               146,
		"tencentcloud_tcr_vpc_attachment":                      147,
		"tencentcloud_vod_adaptive_dynamic_streaming_template": 148,
		"tencentcloud_vod_image_sprite_template":               149,
		"tencentcloud_vod_procedure_template":                  150,
		"tencentcloud_vod_snapshot_by_time_offset_template":    151,
		"tencentcloud_vod_super_player_config":                 152,
		"tencentcloud_vpc":                                     153,
		"tencentcloud_vpc_acl":                                 154,
		"tencentcloud_vpc_acl_attachment":                      155,
		"tencentcloud_vpn_connection":                          156,
		"tencentcloud_vpn_customer_gateway":                    157,
		"tencentcloud_vpn_gateway":                             158,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
