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
					Description: `(Optional) Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. Valid value ranges: (30~7200). and default value is ` + "`" + `300` + "`" + `.`,
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
					Name:        "cam_role_name",
					Description: `(Optional) CAM role name authorized to access.`,
				},
				resource.Attribute{
					Name:        "data_disk",
					Description: `(Optional) Configurations of data disk.`,
				},
				resource.Attribute{
					Name:        "disk_type_policy",
					Description: `(Optional) Policy of cloud disk type. Valid values: ` + "`" + `ORIGINAL` + "`" + ` and ` + "`" + `AUTOMATIC` + "`" + `. Default is ` + "`" + `ORIGINAL` + "`" + `.`,
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
					Name:        "instance_charge_type_prepaid_period",
					Description: `(Optional) The tenancy (in month) of the prepaid instance, NOTE: it only works when instance_charge_type is set to ` + "`" + `PREPAID` + "`" + `. Valid values are ` + "`" + `1` + "`" + `, ` + "`" + `2` + "`" + `, ` + "`" + `3` + "`" + `, ` + "`" + `4` + "`" + `, ` + "`" + `5` + "`" + `, ` + "`" + `6` + "`" + `, ` + "`" + `7` + "`" + `, ` + "`" + `8` + "`" + `, ` + "`" + `9` + "`" + `, ` + "`" + `10` + "`" + `, ` + "`" + `11` + "`" + `, ` + "`" + `12` + "`" + `, ` + "`" + `24` + "`" + `, ` + "`" + `36` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type_prepaid_renew_flag",
					Description: `(Optional) Auto renewal flag. Valid values: ` + "`" + `NOTIFY_AND_AUTO_RENEW` + "`" + `: notify upon expiration and renew automatically, ` + "`" + `NOTIFY_AND_MANUAL_RENEW` + "`" + `: notify upon expiration but do not renew automatically, ` + "`" + `DISABLE_NOTIFY_AND_MANUAL_RENEW` + "`" + `: neither notify upon expiration nor renew automatically. Default value: ` + "`" + `NOTIFY_AND_MANUAL_RENEW` + "`" + `. If this parameter is specified as ` + "`" + `NOTIFY_AND_AUTO_RENEW` + "`" + `, the instance will be automatically renewed on a monthly basis if the account balance is sufficient. NOTE: it only works when instance_charge_type is set to ` + "`" + `PREPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `(Optional) Charge type of instance. Valid values are ` + "`" + `PREPAID` + "`" + `, ` + "`" + `POSTPAID_BY_HOUR` + "`" + `, ` + "`" + `SPOTPAID` + "`" + `. The default is ` + "`" + `POSTPAID_BY_HOUR` + "`" + `. NOTE: ` + "`" + `SPOTPAID` + "`" + ` instance must set ` + "`" + `spot_instance_type` + "`" + ` and ` + "`" + `spot_max_price` + "`" + ` at the same time.`,
				},
				resource.Attribute{
					Name:        "instance_name_settings",
					Description: `(Optional) Settings of CVM instance names.`,
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
					Name:        "spot_instance_type",
					Description: `(Optional) Type of spot instance, only support ` + "`" + `one-time` + "`" + ` now. Note: it only works when instance_charge_type is set to ` + "`" + `SPOTPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "spot_max_price",
					Description: `(Optional) Max price of a spot instance, is the format of decimal string, for example "0.50". Note: it only works when instance_charge_type is set to ` + "`" + `SPOTPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "system_disk_size",
					Description: `(Optional) Volume of system disk in GB. Default is ` + "`" + `50` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "system_disk_type",
					Description: `(Optional) Type of a CVM disk. Valid values: ` + "`" + `CLOUD_PREMIUM` + "`" + ` and ` + "`" + `CLOUD_SSD` + "`" + `. Default is ` + "`" + `CLOUD_PREMIUM` + "`" + `. valid when disk_type_policy is ORIGINAL.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional) ase64-encoded User Data text, the length limit is 16KB. The ` + "`" + `data_disk` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "delete_with_instance",
					Description: `(Optional) Indicates whether the disk remove after instance terminated.`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `(Optional) Volume of disk in GB. Default is ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `(Optional) Types of disk. Valid values: ` + "`" + `CLOUD_PREMIUM` + "`" + ` and ` + "`" + `CLOUD_SSD` + "`" + `. valid when disk_type_policy is ORIGINAL.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) Data disk snapshot ID. The ` + "`" + `instance_name_settings` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Required) CVM instance name.`,
				},
				resource.Attribute{
					Name:        "instance_name_style",
					Description: `(Optional) Type of CVM instance name. Valid values: ` + "`" + `ORIGINAL` + "`" + ` and ` + "`" + `UNIQUE` + "`" + `. Default is ` + "`" + `ORIGINAL` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
					Name:        "multi_zone_subnet_policy",
					Description: `(Optional) Multi zone or subnet strategy, Valid values: PRIORITY and EQUALITY.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) Specifies to which project the scaling group belongs.`,
				},
				resource.Attribute{
					Name:        "replace_load_balancer_unhealthy",
					Description: `(Optional) Enable unhealthy instance replacement. If set to ` + "`" + `true` + "`" + `, AS will replace instances that are found unhealthy in the CLB health check.`,
				},
				resource.Attribute{
					Name:        "replace_monitor_unhealthy",
					Description: `(Optional) Enables unhealthy instance replacement. If set to ` + "`" + `true` + "`" + `, AS will replace instances that are flagged as unhealthy by Cloud Monitor.`,
				},
				resource.Attribute{
					Name:        "retry_policy",
					Description: `(Optional) Available values for retry policies. Valid values: IMMEDIATE_RETRY and INCREMENTAL_INTERVALS.`,
				},
				resource.Attribute{
					Name:        "scaling_mode",
					Description: `(Optional) Indicates scaling mode which creates and terminates instances (classic method), or method first tries to start stopped instances (wake up stopped) to perform scaling operations. Available values: ` + "`" + `CLASSIC_SCALING` + "`" + `, ` + "`" + `WAKE_UP_STOPPED_SCALING` + "`" + `. Default: ` + "`" + `CLASSIC_SCALING` + "`" + `.`,
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
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "user_names",
					Description: `(Optional) User name set as ID of the CAM group members. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
			Type:             "tencentcloud_cam_oidc_sso",
			Category:         "Cloud Access Management(CAM)",
			ShortDescription: `Provides a resource to create a CAM-OIDC-SSO.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"access",
				"management",
				"cam",
				"oidc",
				"sso",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "authorization_endpoint",
					Description: `(Required) Authorization request Endpoint, OpenID Connect identity provider authorization address. Corresponds to the value of the ` + "`" + `authorization_endpoint` + "`" + ` field in the Openid-configuration provided by the Enterprise IdP.`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `(Required) Client ID, the client ID registered with the OpenID Connect identity provider.`,
				},
				resource.Attribute{
					Name:        "identity_key",
					Description: `(Required) The signature public key requires base64_encode. Verify the public key signed by the OpenID Connect identity provider ID Token. For the security of your account, we recommend that you rotate the signed public key regularly.`,
				},
				resource.Attribute{
					Name:        "identity_url",
					Description: `(Required) Identity provider URL. OpenID Connect identity provider identity.Corresponds to the value of the ` + "`" + `issuer` + "`" + ` field in the Openid-configuration provided by the Enterprise IdP.`,
				},
				resource.Attribute{
					Name:        "mapping_filed",
					Description: `(Required) Map field names. Which field in the IdP's id_token maps to the user name of the subuser, usually the sub or name field.`,
				},
				resource.Attribute{
					Name:        "response_mode",
					Description: `(Required) Authorize the request Forsonse mode. Authorization request return mode, form_post and frogment two optional modes, recommended to select form_post mode.`,
				},
				resource.Attribute{
					Name:        "response_type",
					Description: `(Required) Authorization requests The Response type, with a fixed value id_token.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) Authorize the request Scope. openid; email; profile; Authorization request information scope. The default is required openid. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import CAM-OIDC-SSO can be imported using the client_id or any string which can identifier resource, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cam_oidc_sso.foo xxxxxxxxxxx ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import CAM-OIDC-SSO can be imported using the client_id or any string which can identifier resource, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cam_oidc_sso.foo xxxxxxxxxxx ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Optional, ForceNew) Indicates whether the CAM role can login or not.`,
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
					Description: `(Optional, ForceNew,`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `(Optional, ForceNew) Name of the attached CAM user as uniq key. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
					Description: `(Required, ForceNew) Type of CBS medium. Valid values: CLOUD_PREMIUM, CLOUD_SSD, CLOUD_TSSD and CLOUD_HSSD.`,
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
			Type:             "tencentcloud_cdh_instance",
			Category:         "CVM Dedicated Host(CDH)",
			ShortDescription: `Provides a resource to manage CDH instance.`,
			Description:      ``,
			Keywords: []string{
				"cvm",
				"dedicated",
				"host",
				"cdh",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required, ForceNew) The available zone for the CDH instance.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `(Optional) The charge type of instance. Valid values are ` + "`" + `PREPAID` + "`" + `. The default is ` + "`" + `PREPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "host_name",
					Description: `(Optional) The name of the CDH instance. The max length of host_name is 60.`,
				},
				resource.Attribute{
					Name:        "host_type",
					Description: `(Optional, ForceNew) The type of the CDH instance.`,
				},
				resource.Attribute{
					Name:        "prepaid_period",
					Description: `(Optional) The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when charge_type is set to ` + "`" + `PREPAID` + "`" + `. Valid values are ` + "`" + `1` + "`" + `, ` + "`" + `2` + "`" + `, ` + "`" + `3` + "`" + `, ` + "`" + `4` + "`" + `, ` + "`" + `5` + "`" + `, ` + "`" + `6` + "`" + `, ` + "`" + `7` + "`" + `, ` + "`" + `8` + "`" + `, ` + "`" + `9` + "`" + `, ` + "`" + `10` + "`" + `, ` + "`" + `11` + "`" + `, ` + "`" + `12` + "`" + `, ` + "`" + `24` + "`" + `, ` + "`" + `36` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "prepaid_renew_flag",
					Description: `(Optional) Auto renewal flag. Valid values: ` + "`" + `NOTIFY_AND_AUTO_RENEW` + "`" + `: notify upon expiration and renew automatically, ` + "`" + `NOTIFY_AND_MANUAL_RENEW` + "`" + `: notify upon expiration but do not renew automatically, ` + "`" + `DISABLE_NOTIFY_AND_MANUAL_RENEW` + "`" + `: neither notify upon expiration nor renew automatically. Default value: ` + "`" + `NOTIFY_AND_MANUAL_RENEW` + "`" + `. If this parameter is specified as ` + "`" + `NOTIFY_AND_AUTO_RENEW` + "`" + `, the instance will be automatically renewed on a monthly basis if the account balance is sufficient. NOTE: it only works when charge_type is set to ` + "`" + `PREPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) The project the instance belongs to, default to 0. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
					Name:        "cvm_instance_ids",
					Description: `Id of CVM instances that have been created on the CDH instance.`,
				},
				resource.Attribute{
					Name:        "expired_time",
					Description: `Expired time of the instance.`,
				},
				resource.Attribute{
					Name:        "host_resource",
					Description: `An information list of host resource. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "cpu_available_num",
					Description: `The number of available CPU cores of the instance.`,
				},
				resource.Attribute{
					Name:        "cpu_total_num",
					Description: `The number of total CPU cores of the instance.`,
				},
				resource.Attribute{
					Name:        "disk_available_size",
					Description: `Instance disk available capacity, unit in GB.`,
				},
				resource.Attribute{
					Name:        "disk_total_size",
					Description: `Instance disk total capacity, unit in GB.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `Type of the disk.`,
				},
				resource.Attribute{
					Name:        "memory_available_size",
					Description: `Instance memory available capacity, unit in GB.`,
				},
				resource.Attribute{
					Name:        "memory_total_size",
					Description: `Instance memory total capacity, unit in GB.`,
				},
				resource.Attribute{
					Name:        "host_state",
					Description: `State of the CDH instance. ## Import CDH instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cdh_instance.foo host-d6s7i5q4 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "cvm_instance_ids",
					Description: `Id of CVM instances that have been created on the CDH instance.`,
				},
				resource.Attribute{
					Name:        "expired_time",
					Description: `Expired time of the instance.`,
				},
				resource.Attribute{
					Name:        "host_resource",
					Description: `An information list of host resource. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "cpu_available_num",
					Description: `The number of available CPU cores of the instance.`,
				},
				resource.Attribute{
					Name:        "cpu_total_num",
					Description: `The number of total CPU cores of the instance.`,
				},
				resource.Attribute{
					Name:        "disk_available_size",
					Description: `Instance disk available capacity, unit in GB.`,
				},
				resource.Attribute{
					Name:        "disk_total_size",
					Description: `Instance disk total capacity, unit in GB.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `Type of the disk.`,
				},
				resource.Attribute{
					Name:        "memory_available_size",
					Description: `Instance memory available capacity, unit in GB.`,
				},
				resource.Attribute{
					Name:        "memory_total_size",
					Description: `Instance memory total capacity, unit in GB.`,
				},
				resource.Attribute{
					Name:        "host_state",
					Description: `State of the CDH instance. ## Import CDH instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cdh_instance.foo host-d6s7i5q4 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "ipv6_access_switch",
					Description: `(Optional) ipv6 access configuration switch. Only available when area set to ` + "`" + `mainland` + "`" + `. Valid values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. Default value is ` + "`" + `off` + "`" + `.`,
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
			Type:             "tencentcloud_ckafka_instance",
			Category:         "Ckafka",
			ShortDescription: `Use this resource to create ckafka instance.`,
			Description:      ``,
			Keywords: []string{
				"ckafka",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Required) Instance name.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required, ForceNew) Subnet id.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, ForceNew) Vpc id.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Required, ForceNew) Available zone id.`,
				},
				resource.Attribute{
					Name:        "band_width",
					Description: `(Optional, ForceNew) Instance bandwidth in MBps.`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `(Optional) Instance configuration.`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `(Optional, ForceNew) Disk Size. Its interval varies with bandwidth, and the input must be within the interval, which can be viewed through the control. If it is not within the interval, the plan will cause a change when first created.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `(Optional, ForceNew) Type of disk.`,
				},
				resource.Attribute{
					Name:        "dynamic_retention_config",
					Description: `(Optional) Dynamic message retention policy configuration.`,
				},
				resource.Attribute{
					Name:        "kafka_version",
					Description: `(Optional, ForceNew) Kafka version (0.10.2/1.1.1/2.4.1).`,
				},
				resource.Attribute{
					Name:        "msg_retention_time",
					Description: `(Optional) The maximum retention time of instance logs, in minutes. the default is 10080 (7 days), the maximum is 30 days, and the default 0 is not filled, which means that the log retention time recovery policy is not enabled.`,
				},
				resource.Attribute{
					Name:        "multi_zone_flag",
					Description: `(Optional, ForceNew) Indicates whether the instance is multi zones. NOTE: if set to ` + "`" + `true` + "`" + `, ` + "`" + `zone_ids` + "`" + ` must set together.`,
				},
				resource.Attribute{
					Name:        "partition",
					Description: `(Optional, ForceNew) Partition Size. Its interval varies with bandwidth, and the input must be within the interval, which can be viewed through the control. If it is not within the interval, the plan will cause a change when first created.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional, ForceNew) Prepaid purchase time, such as 1, is one month.`,
				},
				resource.Attribute{
					Name:        "public_network",
					Description: `(Optional) Timestamp.`,
				},
				resource.Attribute{
					Name:        "rebalance_time",
					Description: `(Optional) Modification of the rebalancing time after upgrade.`,
				},
				resource.Attribute{
					Name:        "renew_flag",
					Description: `(Optional, ForceNew) Prepaid automatic renewal mark, 0 means the default state, the initial state, 1 means automatic renewal, 2 means clear no automatic renewal (user setting).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, ForceNew) Partition size, the professional version does not need tag.`,
				},
				resource.Attribute{
					Name:        "zone_ids",
					Description: `(Optional, ForceNew) List of available zone id. NOTE: this argument must set together with ` + "`" + `multi_zone_flag` + "`" + `. The ` + "`" + `config` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "auto_create_topic_enable",
					Description: `(Required) Automatic creation. true: enabled, false: not enabled.`,
				},
				resource.Attribute{
					Name:        "default_num_partitions",
					Description: `(Required) If auto.create.topic.enable is set to true and this value is not set, 3 will be used by default.`,
				},
				resource.Attribute{
					Name:        "default_replication_factor",
					Description: `(Required) If auto.create.topic.enable is set to true but this value is not set, 2 will be used by default. The ` + "`" + `dynamic_retention_config` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "bottom_retention",
					Description: `(Optional) Minimum retention time, in minutes.`,
				},
				resource.Attribute{
					Name:        "disk_quota_percentage",
					Description: `(Optional) Disk quota threshold (in percentage) for triggering the message retention time change event.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional) Whether the dynamic message retention time configuration is enabled. 0: disabled; 1: enabled.`,
				},
				resource.Attribute{
					Name:        "step_forward_percentage",
					Description: `(Optional) Percentage by which the message retention time is shortened each time. The ` + "`" + `tags` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Tag key.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Tag value. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import ckafka instance can be imported using the instance_id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_ckafka_instance.foo ckafka-f9ife4zz ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import ckafka instance can be imported using the instance_id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_ckafka_instance.foo ckafka-f9ife4zz ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Required) The number of replica.`,
				},
				resource.Attribute{
					Name:        "topic_name",
					Description: `(Required, ForceNew) Name of the CKafka topic. It must start with a letter, the rest can contain letters, numbers and dashes(-).`,
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
					Description: `(Optional) The subject note. It must start with a letter, and the remaining part can contain letters, numbers and dashes (-).`,
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
			Type:             "tencentcloud_clb_customized_config",
			Category:         "Cloud Load Balancer(CLB)",
			ShortDescription: `Provides a resource to create a CLB customized config.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"load",
				"balancer",
				"clb",
				"customized",
				"config",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config_content",
					Description: `(Required) Content of Customized Config.`,
				},
				resource.Attribute{
					Name:        "config_name",
					Description: `(Required) Name of Customized Config.`,
				},
				resource.Attribute{
					Name:        "load_balancer_ids",
					Description: `(Optional) List of LoadBalancer Ids. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of Customized Config.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Update time of Customized Config. ## Import CLB customized config can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_clb_customized_config.foo pz-diowqstq ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of Customized Config.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Update time of Customized Config. ## Import CLB customized config can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_clb_customized_config.foo pz-diowqstq ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "bandwidth_package_id",
					Description: `(Optional) Bandwidth package id. If set, the ` + "`" + `internet_charge_type` + "`" + ` must be ` + "`" + `BANDWIDTH_PACKAGE` + "`" + `.`,
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
					Name:        "load_balancer_pass_to_target",
					Description: `(Optional) Whether the target allow flow come from clb. If value is true, only check security group of clb, or check both clb and backend instance security group.`,
				},
				resource.Attribute{
					Name:        "log_set_id",
					Description: `(Optional) The id of log set.`,
				},
				resource.Attribute{
					Name:        "log_topic_id",
					Description: `(Optional) The id of log topic.`,
				},
				resource.Attribute{
					Name:        "master_zone_id",
					Description: `(Optional) Setting master zone id of cross available zone disaster recovery, only applicable to open CLB.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional, ForceNew) ID of the project within the CLB instance, ` + "`" + `0` + "`" + ` - Default Project.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Optional) Security groups of the CLB instance. Supports both ` + "`" + `OPEN` + "`" + ` and ` + "`" + `INTERNAL` + "`" + ` CLBs.`,
				},
				resource.Attribute{
					Name:        "slave_zone_id",
					Description: `(Optional) Setting slave zone id of cross available zone disaster recovery, only applicable to open CLB. this zone will undertake traffic when the master is down.`,
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
					Description: `(Optional, ForceNew) VPC ID of the CLB.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Optional) Available zone id, only applicable to open CLB. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
					Description: `(Optional) Backend target type. Valid values: ` + "`" + `NODE` + "`" + `, ` + "`" + `TARGETGROUP` + "`" + `. ` + "`" + `NODE` + "`" + ` means to bind ordinary nodes, ` + "`" + `TARGETGROUP` + "`" + ` means to bind target group. NOTES: TCP/UDP/TCP_SSL listener must configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
			Type:             "tencentcloud_clb_log_set",
			Category:         "Cloud Load Balancer(CLB)",
			ShortDescription: `Provides a resource to create an exclusive CLB Logset.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"load",
				"balancer",
				"clb",
				"log",
				"set",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "period",
					Description: `(Optional, ForceNew) Logset retention period in days. Maximun value is ` + "`" + `90` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Logset creation time.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Logset name, which unique and fixed ` + "`" + `clb_logset` + "`" + ` among all CLS logsets.`,
				},
				resource.Attribute{
					Name:        "topic_count",
					Description: `Number of log topics in logset. ## Import CLB log set can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_clb_logset.foo 4eb9e3a8-9c42-4b32-9ddf-e215e9c92764 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Logset creation time.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Logset name, which unique and fixed ` + "`" + `clb_logset` + "`" + ` among all CLS logsets.`,
				},
				resource.Attribute{
					Name:        "topic_count",
					Description: `Number of log topics in logset. ## Import CLB log set can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_clb_logset.foo 4eb9e3a8-9c42-4b32-9ddf-e215e9c92764 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_clb_log_topic",
			Category:         "Cloud Load Balancer(CLB)",
			ShortDescription: `Provides a resource to create a CLB instance topic.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"load",
				"balancer",
				"clb",
				"log",
				"topic",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "log_set_id",
					Description: `(Required, ForceNew) Log topic of CLB instance.`,
				},
				resource.Attribute{
					Name:        "topic_name",
					Description: `(Required, ForceNew) Log topic of CLB instance. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Log topic creation time.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of log topic. ## Import CLB log topic can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_clb_log_topic.topic lb-7a0t6zqb`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Log topic creation time.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of log topic. ## Import CLB log topic can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_clb_log_topic.topic lb-7a0t6zqb`,
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
					Name:        "port",
					Description: `(Optional) The default port of target group, add server after can use it.`,
				},
				resource.Attribute{
					Name:        "target_group_instances",
					Description: `(Optional) The backend server of target group bind.`,
				},
				resource.Attribute{
					Name:        "target_group_name",
					Description: `(Optional) Target group name.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional, ForceNew) VPC ID, default is based on the network. The ` + "`" + `target_group_instances` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "bind_ip",
					Description: `(Required) The internal ip of target group instance.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) The port of target group instance.`,
				},
				resource.Attribute{
					Name:        "new_port",
					Description: `(Optional) The new port of target group instance.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) The weight of target group instance. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
			Type:             "tencentcloud_cls_config",
			Category:         "CLS",
			ShortDescription: `Provides a resource to create a cls config`,
			Description:      ``,
			Keywords: []string{
				"cls",
				"config",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "extract_rule",
					Description: `(Required) Extraction rule. If ExtractRule is set, LogType must be set.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Collection configuration name.`,
				},
				resource.Attribute{
					Name:        "exclude_paths",
					Description: `(Optional) Collection path blocklist.`,
				},
				resource.Attribute{
					Name:        "log_type",
					Description: `(Optional) Type of the log to be collected. Valid values: json_log: log in JSON format; delimiter_log: log in delimited format; minimalist_log: minimalist log; multiline_log: log in multi-line format; fullregex_log: log in full regex format. Default value: minimalist_log.`,
				},
				resource.Attribute{
					Name:        "output",
					Description: `(Optional) Log topic ID (TopicId) of collection configuration.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) Log collection path containing the filename.`,
				},
				resource.Attribute{
					Name:        "user_define_rule",
					Description: `(Optional) Custom collection rule, which is a serialized JSON string. The ` + "`" + `exclude_paths` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type. Valid values: File, Path.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Specific content corresponding to Type. The ` + "`" + `extract_rule` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "backtracking",
					Description: `(Optional) Size of the data to be rewound in incremental collection mode. Default value: -1 (full collection).`,
				},
				resource.Attribute{
					Name:        "begin_regex",
					Description: `(Optional) First-Line matching rule, which is valid only if log_type is multiline_log or fullregex_log.`,
				},
				resource.Attribute{
					Name:        "delimiter",
					Description: `(Optional) Delimiter for delimited log, which is valid only if log_type is delimiter_log.`,
				},
				resource.Attribute{
					Name:        "filter_key_regex",
					Description: `(Optional) Log keys to be filtered and the corresponding regex.`,
				},
				resource.Attribute{
					Name:        "keys",
					Description: `(Optional) Key name of each extracted field. An empty key indicates to discard the field. This parameter is valid only if log_type is delimiter_log. json_log logs use the key of JSON itself.`,
				},
				resource.Attribute{
					Name:        "log_regex",
					Description: `(Optional) Full log matching rule, which is valid only if log_type is fullregex_log.`,
				},
				resource.Attribute{
					Name:        "time_format",
					Description: `(Optional) Time field format. For more information, please see the output parameters of the time format description of the strftime function in C language.`,
				},
				resource.Attribute{
					Name:        "time_key",
					Description: `(Optional) Time field key name. time_key and time_format must appear in pair.`,
				},
				resource.Attribute{
					Name:        "un_match_log_key",
					Description: `(Optional) Unmatched log key.`,
				},
				resource.Attribute{
					Name:        "un_match_up_load_switch",
					Description: `(Optional) Whether to upload the logs that failed to be parsed. Valid values: true: yes; false: no. The ` + "`" + `filter_key_regex` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) Log key to be filtered.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) Filter rule regex corresponding to key. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
			Type:             "tencentcloud_cls_config_attachment",
			Category:         "CLS",
			ShortDescription: `Provides a resource to create a cls config attachment`,
			Description:      ``,
			Keywords: []string{
				"cls",
				"config",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required, ForceNew) Collection configuration id.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required, ForceNew) Machine group id. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
			Type:             "tencentcloud_cls_config_extra",
			Category:         "CLS",
			ShortDescription: `Provides a resource to create a cls config extra`,
			Description:      ``,
			Keywords: []string{
				"cls",
				"config",
				"extra",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config_flag",
					Description: `(Required) Collection configuration flag.`,
				},
				resource.Attribute{
					Name:        "log_type",
					Description: `(Required) Type of the log to be collected. Valid values: json_log: log in JSON format; delimiter_log: log in delimited format; minimalist_log: minimalist log; multiline_log: log in multi-line format; fullregex_log: log in full regex format. Default value: minimalist_log.`,
				},
				resource.Attribute{
					Name:        "logset_id",
					Description: `(Required) Logset Id.`,
				},
				resource.Attribute{
					Name:        "logset_name",
					Description: `(Required) Logset Name.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Collection configuration name.`,
				},
				resource.Attribute{
					Name:        "topic_id",
					Description: `(Required) Log topic ID (TopicId) of collection configuration.`,
				},
				resource.Attribute{
					Name:        "topic_name",
					Description: `(Required) Topic Name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type. Valid values: container_stdout; container_file; host_file.`,
				},
				resource.Attribute{
					Name:        "container_file",
					Description: `(Optional) Container file path info.`,
				},
				resource.Attribute{
					Name:        "container_stdout",
					Description: `(Optional) Container stdout info.`,
				},
				resource.Attribute{
					Name:        "exclude_paths",
					Description: `(Optional) Collection path blocklist.`,
				},
				resource.Attribute{
					Name:        "extract_rule",
					Description: `(Optional) Extraction rule. If ExtractRule is set, LogType must be set.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Optional) Binding group id.`,
				},
				resource.Attribute{
					Name:        "group_ids",
					Description: `(Optional, ForceNew) Binding group ids.`,
				},
				resource.Attribute{
					Name:        "host_file",
					Description: `(Optional) Node file config info.`,
				},
				resource.Attribute{
					Name:        "user_define_rule",
					Description: `(Optional) Custom collection rule, which is a serialized JSON string. The ` + "`" + `container_file` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "container",
					Description: `(Required) Container name.`,
				},
				resource.Attribute{
					Name:        "file_pattern",
					Description: `(Required) log name.`,
				},
				resource.Attribute{
					Name:        "log_path",
					Description: `(Required) Log Path.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Required) Namespace. There can be multiple namespaces, separated by separators, such as A, B.`,
				},
				resource.Attribute{
					Name:        "exclude_labels",
					Description: `(Optional) Pod label to be excluded.`,
				},
				resource.Attribute{
					Name:        "exclude_namespace",
					Description: `(Optional) Namespaces to be excluded, separated by separators, such as A, B.`,
				},
				resource.Attribute{
					Name:        "include_labels",
					Description: `(Optional) Pod label info.`,
				},
				resource.Attribute{
					Name:        "workload",
					Description: `(Optional) Workload info. The ` + "`" + `container_stdout` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "all_containers",
					Description: `(Required) Is all containers.`,
				},
				resource.Attribute{
					Name:        "exclude_labels",
					Description: `(Optional) Pod label to be excluded.`,
				},
				resource.Attribute{
					Name:        "exclude_namespace",
					Description: `(Optional) Namespaces to be excluded, separated by separators, such as A, B.`,
				},
				resource.Attribute{
					Name:        "include_labels",
					Description: `(Optional) Pod label info.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) Namespace. There can be multiple namespaces, separated by separators, such as A, B.`,
				},
				resource.Attribute{
					Name:        "workloads",
					Description: `(Optional) Workload info. The ` + "`" + `exclude_paths` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type. Valid values: File, Path.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Specific content corresponding to Type. The ` + "`" + `extract_rule` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "backtracking",
					Description: `(Optional) Size of the data to be rewound in incremental collection mode. Default value: -1 (full collection).`,
				},
				resource.Attribute{
					Name:        "begin_regex",
					Description: `(Optional) First-Line matching rule, which is valid only if log_type is multiline_log or fullregex_log.`,
				},
				resource.Attribute{
					Name:        "delimiter",
					Description: `(Optional) Delimiter for delimited log, which is valid only if log_type is delimiter_log.`,
				},
				resource.Attribute{
					Name:        "filter_key_regex",
					Description: `(Optional) Log keys to be filtered and the corresponding regex.`,
				},
				resource.Attribute{
					Name:        "keys",
					Description: `(Optional) Key name of each extracted field. An empty key indicates to discard the field. This parameter is valid only if log_type is delimiter_log. json_log logs use the key of JSON itself.`,
				},
				resource.Attribute{
					Name:        "log_regex",
					Description: `(Optional) Full log matching rule, which is valid only if log_type is fullregex_log.`,
				},
				resource.Attribute{
					Name:        "time_format",
					Description: `(Optional) Time field format. For more information, please see the output parameters of the time format description of the strftime function in C language.`,
				},
				resource.Attribute{
					Name:        "time_key",
					Description: `(Optional) Time field key name. time_key and time_format must appear in pair.`,
				},
				resource.Attribute{
					Name:        "un_match_log_key",
					Description: `(Optional) Unmatched log key.`,
				},
				resource.Attribute{
					Name:        "un_match_up_load_switch",
					Description: `(Optional) Whether to upload the logs that failed to be parsed. Valid values: true: yes; false: no. The ` + "`" + `filter_key_regex` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) Log key to be filtered.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) Filter rule regex corresponding to key. The ` + "`" + `host_file` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "file_pattern",
					Description: `(Required) Log file name.`,
				},
				resource.Attribute{
					Name:        "log_path",
					Description: `(Required) Log file dir.`,
				},
				resource.Attribute{
					Name:        "custom_labels",
					Description: `(Optional) Metadata info. The ` + "`" + `workload` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `(Required) workload type.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) workload name.`,
				},
				resource.Attribute{
					Name:        "container",
					Description: `(Optional) container name.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) namespace. The ` + "`" + `workloads` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `(Required) workload type.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) workload name.`,
				},
				resource.Attribute{
					Name:        "container",
					Description: `(Optional) container name.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) namespace. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
			Type:             "tencentcloud_cls_cos_shipper",
			Category:         "CLS",
			ShortDescription: `Provides a resource to create a cls cos shipper.`,
			Description:      ``,
			Keywords: []string{
				"cls",
				"cos",
				"shipper",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) Destination bucket in the shipping rule to be created.`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Required) Prefix of the shipping directory in the shipping rule to be created.`,
				},
				resource.Attribute{
					Name:        "shipper_name",
					Description: `(Required) Shipping rule name.`,
				},
				resource.Attribute{
					Name:        "topic_id",
					Description: `(Required) ID of the log topic to which the shipping rule to be created belongs.`,
				},
				resource.Attribute{
					Name:        "compress",
					Description: `(Optional) Compression configuration of shipped log.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Optional) Format configuration of shipped log content.`,
				},
				resource.Attribute{
					Name:        "filter_rules",
					Description: `(Optional) Filter rules for shipped logs. Only logs matching the rules can be shipped. All rules are in the AND relationship, and up to five rules can be added. If the array is empty, no filtering will be performed, and all logs will be shipped.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional) Shipping time interval in seconds. Default value: 300. Value range: 300~900.`,
				},
				resource.Attribute{
					Name:        "max_size",
					Description: `(Optional) Maximum size of a file to be shipped, in MB. Default value: 256. Value range: 100~256.`,
				},
				resource.Attribute{
					Name:        "partition",
					Description: `(Optional) Partition rule of shipped log, which can be represented in strftime time format. The ` + "`" + `compress` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Required) Compression format. Valid values: gzip, lzop, none (no compression). The ` + "`" + `content` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Required) Content format. Valid values: json, csv.`,
				},
				resource.Attribute{
					Name:        "csv",
					Description: `(Optional) CSV format content description.Note: this field may return null, indicating that no valid values can be obtained.`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `(Optional) JSON format content description.Note: this field may return null, indicating that no valid values can be obtained. The ` + "`" + `csv` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "delimiter",
					Description: `(Required) Field delimiter.`,
				},
				resource.Attribute{
					Name:        "escape_char",
					Description: `(Required) Field delimiter.`,
				},
				resource.Attribute{
					Name:        "keys",
					Description: `(Required) Names of keys.Note: this field may return null, indicating that no valid values can be obtained.`,
				},
				resource.Attribute{
					Name:        "non_existing_field",
					Description: `(Required) Content used to populate non-existing fields.`,
				},
				resource.Attribute{
					Name:        "print_key",
					Description: `(Required) Whether to print key on the first row of the CSV file. The ` + "`" + `filter_rules` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Filter rule key.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Required) Filter rule.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Filter rule value. The ` + "`" + `json` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "enable_tag",
					Description: `(Required) Enablement flag.`,
				},
				resource.Attribute{
					Name:        "meta_fields",
					Description: `(Required) Metadata information list Note: this field may return null, indicating that no valid values can be obtained.. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import cls cos shipper can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cls_cos_shipper.shipper 5d1b7b2a-c163-4c48-bb01-9ee00584d761 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import cls cos shipper can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cls_cos_shipper.shipper 5d1b7b2a-c163-4c48-bb01-9ee00584d761 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cls_logset",
			Category:         "CLS",
			ShortDescription: `Provides a resource to create a cls logset`,
			Description:      ``,
			Keywords: []string{
				"cls",
				"logset",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "logset_name",
					Description: `(Required) Logset name, which must be unique.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tag description list. Up to 10 tag key-value pairs are supported and must be unique. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time.`,
				},
				resource.Attribute{
					Name:        "role_name",
					Description: `If AssumerUin is not empty, it indicates the service provider who creates the logset.`,
				},
				resource.Attribute{
					Name:        "topic_count",
					Description: `Number of log topics in logset. ## Import cls logset can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cls_logset.logset 5cd3a17e-fb0b-418c-afd7-77b365397426 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time.`,
				},
				resource.Attribute{
					Name:        "role_name",
					Description: `If AssumerUin is not empty, it indicates the service provider who creates the logset.`,
				},
				resource.Attribute{
					Name:        "topic_count",
					Description: `Number of log topics in logset. ## Import cls logset can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cls_logset.logset 5cd3a17e-fb0b-418c-afd7-77b365397426 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cls_machine_group",
			Category:         "CLS",
			ShortDescription: `Provides a resource to create a cls machine group.`,
			Description:      ``,
			Keywords: []string{
				"cls",
				"machine",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group_name",
					Description: `(Required) Machine group name, which must be unique.`,
				},
				resource.Attribute{
					Name:        "machine_group_type",
					Description: `(Required) Type of the machine group to be created.`,
				},
				resource.Attribute{
					Name:        "auto_update",
					Description: `(Optional, ForceNew) Whether to enable automatic update for the machine group.`,
				},
				resource.Attribute{
					Name:        "service_logging",
					Description: `(Optional) Whether to enable the service log to record the logs generated by the LogListener service itself. After it is enabled, the internal logset cls_service_logging and the loglistener_status, loglistener_alarm, and loglistener_business log topics will be created, which will not incur fees.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tag description list. Up to 10 tag key-value pairs are supported and must be unique.`,
				},
				resource.Attribute{
					Name:        "update_end_time",
					Description: `(Optional, ForceNew) Update end time. We recommend you update LogListener during off-peak hours.`,
				},
				resource.Attribute{
					Name:        "update_start_time",
					Description: `(Optional, ForceNew) pdate start time. We recommend you update LogListener during off-peak hours. The ` + "`" + `machine_group_type` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Machine group type. Valid values: ip: the IP addresses of collection machines are stored in Values of the machine group; label: the tags of the machines are stored in Values of the machine group.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Machine description list. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import cls machine group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cls_machine_group.group caf168e7-32cd-4ac6-bf89-1950a760e09c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import cls machine group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cls_machine_group.group caf168e7-32cd-4ac6-bf89-1950a760e09c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cls_topic",
			Category:         "CLS",
			ShortDescription: `Provides a resource to create a cls topic.`,
			Description:      ``,
			Keywords: []string{
				"cls",
				"topic",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "logset_id",
					Description: `(Required) Logset ID.`,
				},
				resource.Attribute{
					Name:        "topic_name",
					Description: `(Required) Log topic name.`,
				},
				resource.Attribute{
					Name:        "auto_split",
					Description: `(Optional) Whether to enable automatic split. Default value: true.`,
				},
				resource.Attribute{
					Name:        "max_split_partitions",
					Description: `(Optional) Maximum number of partitions to split into for this topic if automatic split is enabled. Default value: 50.`,
				},
				resource.Attribute{
					Name:        "partition_count",
					Description: `(Optional, ForceNew) Number of log topic partitions. Default value: 1. Maximum value: 10.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) Lifecycle in days. Value range: 1~366. Default value: 30.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `(Optional, ForceNew) Log topic storage class. Valid values: hot: real-time storage; cold: offline storage. Default value: hot. If cold is passed in, please contact the customer service to add the log topic to the allowlist first..`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tag description list. Up to 10 tag key-value pairs are supported and must be unique. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import cls topic can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cls_topic.topic 2f5764c1-c833-44c5-84c7-950979b2a278 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import cls topic can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cls_topic.topic 2f5764c1-c833-44c5-84c7-950979b2a278 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "acl_body",
					Description: `(Optional) ACL XML body for multiple grant info. NOTE: this argument will overwrite ` + "`" + `acl` + "`" + `. Check https://intl.cloud.tencent.com/document/product/436/7737 for more detail.`,
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
					Name:        "multi_az",
					Description: `(Optional, ForceNew) Indicates whether to create a bucket of multi available zone. NOTE: If set to true, the versioning must enable.`,
				},
				resource.Attribute{
					Name:        "origin_domain_rules",
					Description: `(Optional) Bucket Origin Domain settings.`,
				},
				resource.Attribute{
					Name:        "origin_pull_rules",
					Description: `(Optional) Bucket Origin-Pull settings.`,
				},
				resource.Attribute{
					Name:        "replica_role",
					Description: `(Optional) Request initiator identifier, format: ` + "`" + `qcs::cam::uin/<owneruin>:uin/<subuin>` + "`" + `. NOTE: only ` + "`" + `versioning_enable` + "`" + ` is true can configure this argument.`,
				},
				resource.Attribute{
					Name:        "replica_rules",
					Description: `(Optional) List of replica rule. NOTE: only ` + "`" + `versioning_enable` + "`" + ` is true and ` + "`" + `replica_role` + "`" + ` set can configure this argument.`,
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
					Description: `(Optional) Specifies the number of days after object creation when the specific rule action takes effect.`,
				},
				resource.Attribute{
					Name:        "delete_marker",
					Description: `(Optional) Indicates whether the delete marker of an expired object will be removed. The ` + "`" + `lifecycle_rules` + "`" + ` object supports the following:`,
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
					Name:        "id",
					Description: `(Optional) A unique identifier for the rule. It can be up to 255 characters.`,
				},
				resource.Attribute{
					Name:        "non_current_expiration",
					Description: `(Optional) Specifies when non current object versions shall expire.`,
				},
				resource.Attribute{
					Name:        "non_current_transition",
					Description: `(Optional) Specifies a period in the non current object's transitions.`,
				},
				resource.Attribute{
					Name:        "transition",
					Description: `(Optional) Specifies a period in the object's transitions (documented below). The ` + "`" + `non_current_expiration` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "non_current_days",
					Description: `(Optional) Number of days after non current object creation when the specific rule action takes effect. The maximum value is 3650. The ` + "`" + `non_current_transition` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `(Required) Specifies the storage class to which you want the non current object to transition. Available values include ` + "`" + `STANDARD` + "`" + `, ` + "`" + `STANDARD_IA` + "`" + ` and ` + "`" + `ARCHIVE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "non_current_days",
					Description: `(Optional) Number of days after non current object creation when the specific rule action takes effect. The ` + "`" + `origin_domain_rules` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) Specify domain host.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Domain status, default: ` + "`" + `ENABLED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Specify origin domain type, available values: ` + "`" + `REST` + "`" + `, ` + "`" + `WEBSITE` + "`" + `, ` + "`" + `ACCELERATE` + "`" + `, default: ` + "`" + `REST` + "`" + `. The ` + "`" + `origin_pull_rules` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Required) Allows only a domain name or IP address. You can optionally append a port number to the address.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Required) Priority of origin-pull rules, do not set the same value for multiple rules.`,
				},
				resource.Attribute{
					Name:        "custom_http_headers",
					Description: `(Optional) Specifies the custom headers that you can add for COS to access your origin server.`,
				},
				resource.Attribute{
					Name:        "follow_http_headers",
					Description: `(Optional) Specifies the pass through headers when accessing the origin server.`,
				},
				resource.Attribute{
					Name:        "follow_query_string",
					Description: `(Optional) Specifies whether to pass through COS request query string when accessing the origin server.`,
				},
				resource.Attribute{
					Name:        "follow_redirection",
					Description: `(Optional) Specifies whether to follow 3XX redirect to another origin server to pull data from.`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Optional) Triggers the origin-pull rule when the requested file name matches this prefix.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) the protocol used for COS to access the specified origin server. The available value include ` + "`" + `HTTP` + "`" + `, ` + "`" + `HTTPS` + "`" + ` and ` + "`" + `FOLLOW` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sync_back_to_source",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, COS will not return 3XX status code when pulling data from an origin server. Current available zone: ap-beijing, ap-shanghai, ap-singapore, ap-mumbai. The ` + "`" + `replica_rules` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "destination_bucket",
					Description: `(Required) Destination bucket identifier, format: ` + "`" + `qcs::cos:<region>::<bucketname-appid>` + "`" + `. NOTE: destination bucket must enable versioning.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Required) Status identifier, available values: ` + "`" + `Enabled` + "`" + `, ` + "`" + `Disabled` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "destination_storage_class",
					Description: `(Optional) Storage class of destination, available values: ` + "`" + `STANDARD` + "`" + `, ` + "`" + `INTELLIGENT_TIERING` + "`" + `, ` + "`" + `STANDARD_IA` + "`" + `. default is following current class of destination.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Name of a specific rule.`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Optional) Prefix matching policy. Policies cannot overlap; otherwise, an error will be returned. To match the root directory, leave this parameter empty. The ` + "`" + `transition` + "`" + ` object supports the following:`,
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
					Description: `(Optional) Object storage type, Available values include ` + "`" + `STANDARD` + "`" + `, ` + "`" + `STANDARD_IA` + "`" + ` and ` + "`" + `ARCHIVE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tag of the object. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
					Name:        "param_items",
					Description: `(Optional) Specify parameter list of database. Use ` + "`" + `data.tencentcloud_mysql_default_params` + "`" + ` to query available parameter details.`,
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
					Name:        "storage_limit",
					Description: `(Optional, ForceNew) Storage limit of CynosDB cluster instance, unit in GB. The maximum storage of a non-serverless instance in GB. NOTE: If db_type is ` + "`" + `MYSQL` + "`" + ` and charge_type is ` + "`" + `PREPAID` + "`" + `, the value cannot exceed the maximum storage corresponding to the CPU and memory specifications, when charge_type is ` + "`" + `POSTPAID_BY_HOUR` + "`" + `, this argument is unnecessary.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The tags of the CynosDB cluster. The ` + "`" + `param_items` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "current_value",
					Description: `(Required) Param expected value to set.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of param, e.g. ` + "`" + `character_set_server` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "old_value",
					Description: `(Optional) Param old value, indicates the value which already set, this value is required when modifying current_value. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
			Type:             "tencentcloud_dayu_cc_policy_v2",
			Category:         "Anti-DDoS(DayuV2)",
			ShortDescription: `Use this resource to create a dayu CC policy`,
			Description:      ``,
			Keywords: []string{
				"anti",
				"ddos",
				"dayuv2",
				"dayu",
				"cc",
				"policy",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "business",
					Description: `(Required) Bussiness of resource instance. bgpip indicates anti-anti-ip ip; bgp means exclusive package; bgp-multip means shared packet; net indicates anti-anti-ip pro version.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required) The ID of the resource instance.`,
				},
				resource.Attribute{
					Name:        "cc_black_white_ips",
					Description: `(Optional) Blacklist and whitelist.`,
				},
				resource.Attribute{
					Name:        "cc_geo_ip_policys",
					Description: `(Optional) Details of the CC region blocking policy list.`,
				},
				resource.Attribute{
					Name:        "cc_precision_policys",
					Description: `(Optional) CC Precision Protection List.`,
				},
				resource.Attribute{
					Name:        "cc_precision_req_limits",
					Description: `(Optional) CC frequency throttling policy.`,
				},
				resource.Attribute{
					Name:        "thresholds",
					Description: `(Optional) List of protection threshold configurations. The ` + "`" + `cc_black_white_ips` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "black_white_ip",
					Description: `(Required) Blacklist and whitelist IP addresses.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) Domain.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Protocol.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) IP type, value [black(blacklist IP), white (whitelist IP)].`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `(Optional) Create time.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `(Optional) Modify time. The ` + "`" + `cc_geo_ip_policys` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) User action, drop or arg.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) domain.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Protocol, preferably HTTP, HTTPS.`,
				},
				resource.Attribute{
					Name:        "region_type",
					Description: `(Required) Regional types, divided into china, oversea and customized.`,
				},
				resource.Attribute{
					Name:        "area_list",
					Description: `(Optional) The list of region IDs that the user selects to block.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `(Optional) Create time.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `(Optional) Modify time. The ` + "`" + `cc_precision_policys` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) Domain.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) Ip address.`,
				},
				resource.Attribute{
					Name:        "policy_action",
					Description: `(Required) Policy mode (discard or captcha).`,
				},
				resource.Attribute{
					Name:        "policys",
					Description: `(Required) A list of policies.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Protocol. The ` + "`" + `cc_precision_req_limits` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) Domain.`,
				},
				resource.Attribute{
					Name:        "level",
					Description: `(Required) Protection rating, the optional value of default means default policy, loose means loose, and strict means strict.`,
				},
				resource.Attribute{
					Name:        "policys",
					Description: `(Required) The CC Frequency Limit Policy Item field.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Protocol, preferably HTTP, HTTPS. The ` + "`" + `policys` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) The frequency limit policy mode, the optional value of arg indicates the verification code, and drop indicates the discard.`,
				},
				resource.Attribute{
					Name:        "execute_duration",
					Description: `(Required) The duration of the frequency limit policy can be taken from 1 to 86400 per second.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Required) The policy item is compared, and the optional value include indicates inclusion, and equal means equal.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Required) Statistical period, take values 1, 10, 30, 60, in seconds.`,
				},
				resource.Attribute{
					Name:        "request_num",
					Description: `(Required) The number of requests, the value is 1 to 20000.`,
				},
				resource.Attribute{
					Name:        "cookie",
					Description: `(Optional) Cookies, one of the three policy entries can only be filled in.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `(Optional) Uri, one of the three policy entries can only be filled in.`,
				},
				resource.Attribute{
					Name:        "user_agent",
					Description: `(Optional) User-Agent, only one of the three policy entries can be filled in. The ` + "`" + `policys` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "field_name",
					Description: `(Required) Configuration item types, currently only support value.`,
				},
				resource.Attribute{
					Name:        "field_type",
					Description: `(Required) Configuration fields with the desirable values cgi, ua, cookie, referer, accept, srcip.`,
				},
				resource.Attribute{
					Name:        "value_operator",
					Description: `(Required) Configure the item-value comparison mode, which can be taken as the value of evaluate, not_equal, include.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Configure the value. The ` + "`" + `thresholds` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) domain.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `(Required) Cleaning threshold, -1 indicates that the ` + "`" + `default` + "`" + ` mode is turned on. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
			Type:             "tencentcloud_dayu_ddos_policy_v2",
			Category:         "Anti-DDoS(DayuV2)",
			ShortDescription: `Use this resource to create dayu DDoS policy v2`,
			Description:      ``,
			Keywords: []string{
				"anti",
				"ddos",
				"dayuv2",
				"dayu",
				"policy",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required, ForceNew) The ID of the resource instance.`,
				},
				resource.Attribute{
					Name:        "acls",
					Description: `(Optional) Port ACL policy for DDoS protection.`,
				},
				resource.Attribute{
					Name:        "black_white_ips",
					Description: `(Optional) DDoS-protected IP blacklist and whitelist.`,
				},
				resource.Attribute{
					Name:        "business",
					Description: `(Optional) Bussiness of resource instance. bgpip indicates anti-anti-ip ip; bgp means exclusive package; bgp-multip means shared packet; net indicates anti-anti-ip pro version.`,
				},
				resource.Attribute{
					Name:        "ddos_ai",
					Description: `(Optional) AI protection switch, take the value [` + "`" + `on` + "`" + `, ` + "`" + `off` + "`" + `].`,
				},
				resource.Attribute{
					Name:        "ddos_connect_limit",
					Description: `(Optional) DDoS connection suppression options.`,
				},
				resource.Attribute{
					Name:        "ddos_geo_ip_block_config",
					Description: `(Optional) DDoS-protected area block configuration.`,
				},
				resource.Attribute{
					Name:        "ddos_level",
					Description: `(Optional) Protection class, value [` + "`" + `low` + "`" + `, ` + "`" + `middle` + "`" + `, ` + "`" + `high` + "`" + `].`,
				},
				resource.Attribute{
					Name:        "ddos_speed_limit_config",
					Description: `(Optional) Access speed limit configuration for DDoS protection.`,
				},
				resource.Attribute{
					Name:        "ddos_threshold",
					Description: `(Optional) DDoS cleaning threshold, value[0, 60, 80, 100, 150, 200, 250, 300, 400, 500, 700, 1000]; When the value is set to 0, it means that the default value is adopted.`,
				},
				resource.Attribute{
					Name:        "packet_filters",
					Description: `(Optional) Feature filtering rules for DDoS protection.`,
				},
				resource.Attribute{
					Name:        "protocol_block_config",
					Description: `(Optional) Protocol block configuration for DDoS protection. The ` + "`" + `acls` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) Action, optional values: drop, transmit, forward.`,
				},
				resource.Attribute{
					Name:        "d_port_end",
					Description: `(Required) The destination port ends, and the value range is 0~65535.`,
				},
				resource.Attribute{
					Name:        "d_port_start",
					Description: `(Required) The destination port starts, and the value range is 0~65535.`,
				},
				resource.Attribute{
					Name:        "forward_protocol",
					Description: `(Required) Protocol type, desirable values tcp, udp, all.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Required) Policy priority, the lower the number, the higher the level, the higher the rule matches, taking a value of 1-1000.Note: This field may return null, indicating that a valid value could not be retrieved.`,
				},
				resource.Attribute{
					Name:        "s_port_end",
					Description: `(Required) The source port ends, and the acceptable value ranges from 0 to 65535.`,
				},
				resource.Attribute{
					Name:        "s_port_start",
					Description: `(Required) The source port starts, and the value range is 0~65535. The ` + "`" + `black_white_ips` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "ip_type",
					Description: `(Required) IP type, value [` + "`" + `black` + "`" + `(blacklist IP), ` + "`" + `white` + "`" + ` (whitelist IP)].`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) Ip of resource instance. The ` + "`" + `ddos_connect_limit` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "bad_conn_threshold",
					Description: `(Required) Based on connection suppression trigger threshold, value range [0,4294967295].`,
				},
				resource.Attribute{
					Name:        "conn_timeout",
					Description: `(Required) Abnormal connection detection condition, connection timeout, value range [0,65535].`,
				},
				resource.Attribute{
					Name:        "dst_conn_limit",
					Description: `(Required) Concurrent connection control based on destination IP+ destination port.`,
				},
				resource.Attribute{
					Name:        "dst_new_limit",
					Description: `(Required) Limit on the number of news per second based on the destination IP.`,
				},
				resource.Attribute{
					Name:        "null_conn_enable",
					Description: `(Required) Abnormal connection detection conditions, empty connection guard switch, value range[0,1].`,
				},
				resource.Attribute{
					Name:        "sd_conn_limit",
					Description: `(Required) Concurrent connection control based on source IP + destination IP.`,
				},
				resource.Attribute{
					Name:        "sd_new_limit",
					Description: `(Required) The limit on the number of news per second based on source IP + destination IP.`,
				},
				resource.Attribute{
					Name:        "syn_limit",
					Description: `(Required) Anomaly connection detection condition, syn threshold, value range [0,100].`,
				},
				resource.Attribute{
					Name:        "syn_rate",
					Description: `(Required) Anomalous connection detection condition, percentage of syn ack, value range [0,100]. The ` + "`" + `ddos_geo_ip_block_config` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) Block action, take the value [` + "`" + `drop` + "`" + `, ` + "`" + `trans` + "`" + `].`,
				},
				resource.Attribute{
					Name:        "area_list",
					Description: `(Required) When the RegionType is customized, the AreaList must be filled in, and a maximum of 128 must be filled in.`,
				},
				resource.Attribute{
					Name:        "region_type",
					Description: `(Required) Zone type, value [oversea (overseas),china (domestic),customized (custom region)]. The ` + "`" + `ddos_speed_limit_config` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Required) Bandwidth bps.`,
				},
				resource.Attribute{
					Name:        "dst_port_list",
					Description: `(Required) List of port ranges, up to 8, multiple; Separated, the range is represented with -; this port range must be filled in; fill in the style 1:0-65535, style 2:80; 443; 1000-2000.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Required) Speed limit mode, take the value [1 (speed limit based on source IP),2 (speed limit based on destination port)].`,
				},
				resource.Attribute{
					Name:        "packet_rate",
					Description: `(Required) Packet rate pps.`,
				},
				resource.Attribute{
					Name:        "protocol_list",
					Description: `(Required) IP protocol numbers, take the value[ ALL (all protocols),TCP (tcp protocol),UDP (udp protocol),SMP (smp protocol),1; 2-100 (custom protocol number range, up to 8)]. The ` + "`" + `packet_filters` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) Action, take the value [drop,transmit,drop_black (discard and black out),drop_rst (Interception),drop_black_rst (intercept and block),forward].`,
				},
				resource.Attribute{
					Name:        "d_port_end",
					Description: `(Required) The end destination port, take the value 1~65535, which must be greater than or equal to the starting destination port.`,
				},
				resource.Attribute{
					Name:        "d_port_start",
					Description: `(Required) From the destination port, take the value 0~65535.`,
				},
				resource.Attribute{
					Name:        "depth2",
					Description: `(Required) Second detection depth starting from the second detection position, value [0,1500].`,
				},
				resource.Attribute{
					Name:        "depth",
					Description: `(Required) Detection depth from the detection position, value [0,1500].`,
				},
				resource.Attribute{
					Name:        "is_not2",
					Description: `(Required) Whether the second detection contains the detected value, the value [0 (included),1 (not included)].`,
				},
				resource.Attribute{
					Name:        "is_not",
					Description: `(Required) Whether to include the detected value, take the value [0 (included),1 (not included)].`,
				},
				resource.Attribute{
					Name:        "match_begin2",
					Description: `(Required) The second detection position. take the value [begin_l3 (IP header),begin_l4 (TCP/UDP header),begin_l5 (T load), no_match (mismatch)].`,
				},
				resource.Attribute{
					Name:        "match_begin",
					Description: `(Required) Detect position, take the value [begin_l3 (IP header),begin_l4 (TCP/UDP header),begin_l5 (T load), no_match (mismatch)].`,
				},
				resource.Attribute{
					Name:        "match_logic",
					Description: `(Required) When there is a second detection condition, the and/or relationship with the first detection condition, takes the value [And (and relationship),none (fill in this value when there is no second detection condition)].`,
				},
				resource.Attribute{
					Name:        "match_type2",
					Description: `(Required) The second type of detection, takes the value [sunday (keyword),pcre (regular expression)].`,
				},
				resource.Attribute{
					Name:        "match_type",
					Description: `(Required) Detection type, value [sunday (keyword),pcre (regular expression)].`,
				},
				resource.Attribute{
					Name:        "offset2",
					Description: `(Required) Offset from the second detection position, value range [0,Depth2].`,
				},
				resource.Attribute{
					Name:        "offset",
					Description: `(Required) Offset from detection position, value range [0, Depth].`,
				},
				resource.Attribute{
					Name:        "pktlen_max",
					Description: `(Required) The maximum message length, taken from 1 to 1500, must be greater than or equal to the minimum message length.`,
				},
				resource.Attribute{
					Name:        "pktlen_min",
					Description: `(Required) Minimum message length, 1-1500.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Protocol, value [tcp udp icmp all].`,
				},
				resource.Attribute{
					Name:        "s_port_end",
					Description: `(Required) End source port, take the value 1~65535, must be greater than or equal to the starting source port.`,
				},
				resource.Attribute{
					Name:        "s_port_start",
					Description: `(Required) Start the source port, take the value 0~65535.`,
				},
				resource.Attribute{
					Name:        "str2",
					Description: `(Required) The second detection value, the key string or regular expression, takes the value [When the detection type is sunday, please fill in the string or hexadecimal bytecode, for example 13233 corresponds to the hexadecimal bytecode of the string ` + "`" + `123` + "`" + `;When the detection type is pcre, please fill in the regular expression string;].`,
				},
				resource.Attribute{
					Name:        "str",
					Description: `(Required) Detect values, key strings or regular expressions, take the value [When the detection type is sunday, please fill in the string or hexadecimal bytecode, for example 13233 corresponds to the hexadecimal bytecode of the string ` + "`" + `123` + "`" + `;When the detection type is pcre, please fill in the regular expression string;]. The ` + "`" + `protocol_block_config` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "drop_icmp",
					Description: `(Required) ICMP block, value [0 (block off), 1 (block on)].`,
				},
				resource.Attribute{
					Name:        "drop_other",
					Description: `(Required) Other block, value [0 (block off), 1 (block on)].`,
				},
				resource.Attribute{
					Name:        "drop_tcp",
					Description: `(Required) TCP block, value [0 (block off), 1 (block on)].`,
				},
				resource.Attribute{
					Name:        "drop_udp",
					Description: `(Required) UDP block, value [0 (block off), 1 (block on)]. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
			Type:             "tencentcloud_dayu_eip",
			Category:         "Anti-DDoS(DayuV2)",
			ShortDescription: `Use this resource to create dayu eip rule`,
			Description:      ``,
			Keywords: []string{
				"anti",
				"ddos",
				"dayuv2",
				"dayu",
				"eip",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bind_resource_id",
					Description: `(Required, ForceNew) Resource id to bind.`,
				},
				resource.Attribute{
					Name:        "bind_resource_region",
					Description: `(Required, ForceNew) Resource region to bind.`,
				},
				resource.Attribute{
					Name:        "bind_resource_type",
					Description: `(Required, ForceNew) Resource type to bind, value range [` + "`" + `clb` + "`" + `, ` + "`" + `cvm` + "`" + `].`,
				},
				resource.Attribute{
					Name:        "eip",
					Description: `(Required, ForceNew) Eip of the resource.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required, ForceNew) ID of the resource. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "created_time",
					Description: `Created time of the resource instance.`,
				},
				resource.Attribute{
					Name:        "eip_address_status",
					Description: `Eip address status of the resource instance.`,
				},
				resource.Attribute{
					Name:        "eip_bound_rsc_eni",
					Description: `Eip bound rsc eni of the resource instance.`,
				},
				resource.Attribute{
					Name:        "eip_bound_rsc_ins",
					Description: `Eip bound rsc ins of the resource instance.`,
				},
				resource.Attribute{
					Name:        "eip_bound_rsc_vip",
					Description: `Eip bound rsc vip of the resource instance.`,
				},
				resource.Attribute{
					Name:        "expired_time",
					Description: `Expired time of the resource instance.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `Modify time of the resource instance.`,
				},
				resource.Attribute{
					Name:        "protection_status",
					Description: `Protection status of the resource instance.`,
				},
				resource.Attribute{
					Name:        "resource_region",
					Description: `Region of the resource instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "created_time",
					Description: `Created time of the resource instance.`,
				},
				resource.Attribute{
					Name:        "eip_address_status",
					Description: `Eip address status of the resource instance.`,
				},
				resource.Attribute{
					Name:        "eip_bound_rsc_eni",
					Description: `Eip bound rsc eni of the resource instance.`,
				},
				resource.Attribute{
					Name:        "eip_bound_rsc_ins",
					Description: `Eip bound rsc ins of the resource instance.`,
				},
				resource.Attribute{
					Name:        "eip_bound_rsc_vip",
					Description: `Eip bound rsc vip of the resource instance.`,
				},
				resource.Attribute{
					Name:        "expired_time",
					Description: `Expired time of the resource instance.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `Modify time of the resource instance.`,
				},
				resource.Attribute{
					Name:        "protection_status",
					Description: `Protection status of the resource instance.`,
				},
				resource.Attribute{
					Name:        "resource_region",
					Description: `Region of the resource instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dayu_l4_rule",
			Category:         "Anti-DDoS(DayuV2) Anti-DDoS(Dayu)",
			ShortDescription: `Use this resource to create dayu layer 4 rule`,
			Description:      ``,
			Keywords: []string{
				"anti",
				"ddos",
				"dayuv2",
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
			Type:             "tencentcloud_dayu_l7_rule_v2",
			Category:         "Anti-DDoS(DayuV2)",
			ShortDescription: `Use this resource to create dayu new layer 7 rule`,
			Description:      ``,
			Keywords: []string{
				"anti",
				"ddos",
				"dayuv2",
				"dayu",
				"l7",
				"rule",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required, ForceNew) ID of the resource that the layer 7 rule works for.`,
				},
				resource.Attribute{
					Name:        "resource_ip",
					Description: `(Required, ForceNew) Ip of the resource that the layer 7 rule works for.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `(Required, ForceNew) Type of the resource that the layer 7 rule works for, valid value is ` + "`" + `bgpip` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Required) A list of layer 7 rules. Each element contains the following attributes: The ` + "`" + `rule` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) Domain of the rule.`,
				},
				resource.Attribute{
					Name:        "keep_enable",
					Description: `(Required) session hold switch.`,
				},
				resource.Attribute{
					Name:        "keeptime",
					Description: `(Required) The keeptime of the layer 4 rule.`,
				},
				resource.Attribute{
					Name:        "lb_type",
					Description: `(Required) LB type of the rule, ` + "`" + `1` + "`" + ` for weight cycling and ` + "`" + `2` + "`" + ` for IP hash.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Protocol of the rule.`,
				},
				resource.Attribute{
					Name:        "source_list",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) Source type, ` + "`" + `1` + "`" + ` for source of host, ` + "`" + `2` + "`" + ` for source of IP.`,
				},
				resource.Attribute{
					Name:        "cc_enable",
					Description: `(Optional) HTTPS protocol CC protection status, value [0 (off), 1 (on)], defaule is 0.`,
				},
				resource.Attribute{
					Name:        "cert_type",
					Description: `(Optional) The source of the certificate must be filled in when the forwarding protocol is https, the value [2 (Tencent Cloud Hosting Certificate)], and 0 when the forwarding protocol is http.`,
				},
				resource.Attribute{
					Name:        "https_to_http_enable",
					Description: `(Optional) Whether to enable the Https protocol to use Http back-to-source, take the value [0 (off), 1 (on)], do not fill in the default is off, defaule is 0.`,
				},
				resource.Attribute{
					Name:        "ssl_id",
					Description: `(Optional) When the certificate source is a Tencent Cloud managed certificate, this field must be filled in with the managed certificate ID. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
			Type:             "tencentcloud_dnspod_domain_instance",
			Category:         "DNSPOD",
			ShortDescription: `Provide a resource to create a DnsPod Domain instance.`,
			Description:      ``,
			Keywords: []string{
				"dnspod",
				"domain",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) The Domain.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Optional, ForceNew) The Group Id of Domain.`,
				},
				resource.Attribute{
					Name:        "is_mark",
					Description: `(Optional, ForceNew) Whether to Mark the Domain.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `(Optional) The remark of Domain.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The status of Domain. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the domain. ## Import DnsPod Domain instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_dnspod_domain_instance.foo domain ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the domain. ## Import DnsPod Domain instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_dnspod_domain_instance.foo domain ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dnspod_record",
			Category:         "DNSPOD",
			ShortDescription: `Provide a resource to create a DnsPod record.`,
			Description:      ``,
			Keywords: []string{
				"dnspod",
				"record",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) The Domain.`,
				},
				resource.Attribute{
					Name:        "record_line",
					Description: `(Required) The record line.`,
				},
				resource.Attribute{
					Name:        "record_type",
					Description: `(Required) The record type.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The record value.`,
				},
				resource.Attribute{
					Name:        "mx",
					Description: `(Optional) MX priority, valid when the record type is MX, range 1-20. Note: must set when record type equal MX.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Records the initial state, with values ranging from ENABLE and DISABLE. The default is ENABLE, and if DISABLE is passed in, resolution will not take effect and the limits of load balancing will not be verified.`,
				},
				resource.Attribute{
					Name:        "sub_domain",
					Description: `(Optional) The host records, default value is ` + "`" + `@` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) TTL, the range is 1-604800, and the minimum value of different levels of domain names is different. Default is 600.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) Weight information. An integer from 0 to 100. Only enterprise VIP domain names are available, 0 means off, does not pass this parameter, means that the weight information is not set. Default is 0. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "monitor_status",
					Description: `The D monitoring status of the record.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "monitor_status",
					Description: `The D monitoring status of the record.`,
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
					Description: `(Optional, ForceNew) The type of eip. Valid value: ` + "`" + `EIP` + "`" + ` and ` + "`" + `AnycastEIP` + "`" + ` and ` + "`" + `HighQualityEIP` + "`" + `. Default is ` + "`" + `EIP` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
			Type:             "tencentcloud_eks_cluster",
			Category:         "Tencent Kubernetes Engine(TKE)",
			ShortDescription: `Provides an elastic kubernetes cluster resource.`,
			Description:      ``,
			Keywords: []string{
				"tencent",
				"kubernetes",
				"engine",
				"tke",
				"eks",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Required) Name of EKS cluster.`,
				},
				resource.Attribute{
					Name:        "k8s_version",
					Description: `(Required, ForceNew) Kubernetes version of EKS cluster.`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `(Required) Subnet Ids for EKS cluster.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, ForceNew) Vpc Id of EKS cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_desc",
					Description: `(Optional) Description of EKS cluster.`,
				},
				resource.Attribute{
					Name:        "dns_servers",
					Description: `(Optional) List of cluster custom DNS Server info.`,
				},
				resource.Attribute{
					Name:        "enable_vpc_core_dns",
					Description: `(Optional, ForceNew) Indicates whether to enable dns in user cluster, default value is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "extra_param",
					Description: `(Optional, ForceNew) Extend parameters.`,
				},
				resource.Attribute{
					Name:        "internal_lb",
					Description: `(Optional) Cluster internal access LoadBalancer info.`,
				},
				resource.Attribute{
					Name:        "need_delete_cbs",
					Description: `(Optional) Delete CBS after EKS cluster remove.`,
				},
				resource.Attribute{
					Name:        "public_lb",
					Description: `(Optional) Cluster public access LoadBalancer info.`,
				},
				resource.Attribute{
					Name:        "service_subnet_id",
					Description: `(Optional) Subnet id of service.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of EKS cluster. The ` + "`" + `dns_servers` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional) DNS Server domain. Empty indicates all domain.`,
				},
				resource.Attribute{
					Name:        "servers",
					Description: `(Optional) List of DNS Server IP address, pattern: "ip[:port]". The ` + "`" + `internal_lb` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Indicates weather the internal access LB enabled.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) ID of subnet which related to Internal LB. The ` + "`" + `public_lb` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Indicates weather the public access LB enabled.`,
				},
				resource.Attribute{
					Name:        "allow_from_cidrs",
					Description: `(Optional) List of CIDRs which allowed to access.`,
				},
				resource.Attribute{
					Name:        "extra_param",
					Description: `(Optional) Extra param text json.`,
				},
				resource.Attribute{
					Name:        "security_policies",
					Description: `(Optional) List of security allow IP or CIDRs, default deny all. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "kube_config",
					Description: `EKS cluster kubeconfig. ## Import ` + "`" + `` + "`" + `` + "`" + ` terraform import tencentcloud_eks_cluster.foo cluster-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "kube_config",
					Description: `EKS cluster kubeconfig. ## Import ` + "`" + `` + "`" + `` + "`" + ` terraform import tencentcloud_eks_cluster.foo cluster-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_eks_container_instance",
			Category:         "Tencent Kubernetes Engine(TKE)",
			ShortDescription: `Provides an elastic kubernetes service container instance.`,
			Description:      ``,
			Keywords: []string{
				"tencent",
				"kubernetes",
				"engine",
				"tke",
				"eks",
				"container",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "container",
					Description: `(Required) List of container.`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `(Required) The number of CPU cores. Check https://intl.cloud.tencent.com/document/product/457/34057 for specification references.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `(Required) Memory size. Check https://intl.cloud.tencent.com/document/product/457/34057 for specification references.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of EKS container instance.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Required) List of security group id.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) Subnet ID of container instance.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) VPC ID.`,
				},
				resource.Attribute{
					Name:        "auto_create_eip",
					Description: `(Optional) Indicates whether to create EIP instead of specify existing EIPs. Conflict with ` + "`" + `existed_eip_ids` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cam_role_name",
					Description: `(Optional) CAM role name authorized to access.`,
				},
				resource.Attribute{
					Name:        "cbs_volume",
					Description: `(Optional) List of CBS volume.`,
				},
				resource.Attribute{
					Name:        "cpu_type",
					Description: `(Optional) Type of cpu, which can set to ` + "`" + `intel` + "`" + ` or ` + "`" + `amd` + "`" + `. It also support backup list like ` + "`" + `amd,intel` + "`" + ` which indicates using ` + "`" + `intel` + "`" + ` when ` + "`" + `amd` + "`" + ` sold out.`,
				},
				resource.Attribute{
					Name:        "dns_config_options",
					Description: `(Optional, ForceNew) Map of DNS config options.`,
				},
				resource.Attribute{
					Name:        "dns_names_servers",
					Description: `(Optional, ForceNew) IP Addresses of DNS Servers.`,
				},
				resource.Attribute{
					Name:        "dns_searches",
					Description: `(Optional, ForceNew) List of DNS Search Domain.`,
				},
				resource.Attribute{
					Name:        "eip_delete_policy",
					Description: `(Optional) Indicates weather the EIP release or not after instance deleted. Conflict with ` + "`" + `existed_eip_ids` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "eip_max_bandwidth_out",
					Description: `(Optional) Maximum outgoing bandwidth to the public network, measured in Mbps (Mega bits per second). Conflict with ` + "`" + `existed_eip_ids` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "eip_service_provider",
					Description: `(Optional) EIP service provider. Default is ` + "`" + `BGP` + "`" + `, values ` + "`" + `CMCC` + "`" + `,` + "`" + `CTCC` + "`" + `,` + "`" + `CUCC` + "`" + ` are available for whitelist customer. Conflict with ` + "`" + `existed_eip_ids` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "existed_eip_ids",
					Description: `(Optional) Existed EIP ID List which used to bind container instance. Conflict with ` + "`" + `auto_create_eip` + "`" + ` and auto create EIP options.`,
				},
				resource.Attribute{
					Name:        "gpu_count",
					Description: `(Optional) Count of GPU. Check https://intl.cloud.tencent.com/document/product/457/34057 for specification references.`,
				},
				resource.Attribute{
					Name:        "gpu_type",
					Description: `(Optional) Type of GPU. Check https://intl.cloud.tencent.com/document/product/457/34057 for specification references.`,
				},
				resource.Attribute{
					Name:        "image_registry_credential",
					Description: `(Optional) List of credentials which pull from image registry.`,
				},
				resource.Attribute{
					Name:        "init_container",
					Description: `(Optional) List of initialized container.`,
				},
				resource.Attribute{
					Name:        "nfs_volume",
					Description: `(Optional) List of NFS volume.`,
				},
				resource.Attribute{
					Name:        "restart_policy",
					Description: `(Optional) Container instance restart policy. Available values: ` + "`" + `Always` + "`" + `, ` + "`" + `Never` + "`" + `, ` + "`" + `OnFailure` + "`" + `. The ` + "`" + `cbs_volume` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "disk_id",
					Description: `(Required) ID of CBS.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of CBS volume. The ` + "`" + `container` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Required) Image of Container.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Container.`,
				},
				resource.Attribute{
					Name:        "args",
					Description: `(Optional) Container launch argument list.`,
				},
				resource.Attribute{
					Name:        "commands",
					Description: `(Optional) Container launch command list.`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `(Optional) Number of cpu core of container.`,
				},
				resource.Attribute{
					Name:        "env_vars",
					Description: `(Optional) Map of environment variables of container OS.`,
				},
				resource.Attribute{
					Name:        "liveness_probe",
					Description: `(Optional) Configuration block of LivenessProbe.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `(Optional) Memory size of container.`,
				},
				resource.Attribute{
					Name:        "readiness_probe",
					Description: `(Optional) Configuration block of ReadinessProbe.`,
				},
				resource.Attribute{
					Name:        "volume_mount",
					Description: `(Optional) List of volume mount informations.`,
				},
				resource.Attribute{
					Name:        "working_dir",
					Description: `(Optional) Container working directory. The ` + "`" + `image_registry_credential` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of credential.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Password.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `(Optional) Address of image registry.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) Username. The ` + "`" + `init_container` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Required) Image of Container.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Container.`,
				},
				resource.Attribute{
					Name:        "args",
					Description: `(Optional) Container launch argument list.`,
				},
				resource.Attribute{
					Name:        "commands",
					Description: `(Optional) Container launch command list.`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `(Optional) Number of cpu core of container.`,
				},
				resource.Attribute{
					Name:        "env_vars",
					Description: `(Optional) Map of environment variables of container OS.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `(Optional) Memory size of container.`,
				},
				resource.Attribute{
					Name:        "volume_mount",
					Description: `(Optional) List of volume mount informations.`,
				},
				resource.Attribute{
					Name:        "working_dir",
					Description: `(Optional) Container working directory. The ` + "`" + `liveness_probe` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "exec_commands",
					Description: `(Optional) List of execution commands.`,
				},
				resource.Attribute{
					Name:        "failure_threshold",
					Description: `(Optional) Minimum consecutive failures for the probe to be considered failed after having succeeded.Default: ` + "`" + `3` + "`" + `. Minimum value is ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "http_get_path",
					Description: `(Optional) HttpGet detection path.`,
				},
				resource.Attribute{
					Name:        "http_get_port",
					Description: `(Optional) HttpGet detection port.`,
				},
				resource.Attribute{
					Name:        "http_get_scheme",
					Description: `(Optional) HttpGet detection scheme. Available values: ` + "`" + `HTTP` + "`" + `, ` + "`" + `HTTPS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "init_delay_seconds",
					Description: `(Optional) Number of seconds after the container has started before probes are initiated.`,
				},
				resource.Attribute{
					Name:        "period_seconds",
					Description: `(Optional) How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "success_threshold",
					Description: `(Optional) Minimum consecutive successes for the probe to be considered successful after having failed. Default: ` + "`" + `1` + "`" + `. Must be 1 for liveness. Minimum value is ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tcp_socket_port",
					Description: `(Optional) TCP Socket detection port.`,
				},
				resource.Attribute{
					Name:        "timeout_seconds",
					Description: `(Optional) Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is ` + "`" + `1` + "`" + `. The ` + "`" + `nfs_volume` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of NFS volume.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) NFS volume path.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `(Required) NFS server address.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Indicates whether the volume is read only. Default is ` + "`" + `false` + "`" + `. The ` + "`" + `readiness_probe` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "exec_commands",
					Description: `(Optional) List of execution commands.`,
				},
				resource.Attribute{
					Name:        "failure_threshold",
					Description: `(Optional) Minimum consecutive failures for the probe to be considered failed after having succeeded.Default: ` + "`" + `3` + "`" + `. Minimum value is ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "http_get_path",
					Description: `(Optional) HttpGet detection path.`,
				},
				resource.Attribute{
					Name:        "http_get_port",
					Description: `(Optional) HttpGet detection port.`,
				},
				resource.Attribute{
					Name:        "http_get_scheme",
					Description: `(Optional) HttpGet detection scheme. Available values: ` + "`" + `HTTP` + "`" + `, ` + "`" + `HTTPS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "init_delay_seconds",
					Description: `(Optional) Number of seconds after the container has started before probes are initiated.`,
				},
				resource.Attribute{
					Name:        "period_seconds",
					Description: `(Optional) How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "success_threshold",
					Description: `(Optional) Minimum consecutive successes for the probe to be considered successful after having failed. Default: ` + "`" + `1` + "`" + `. Must be 1 for liveness. Minimum value is ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tcp_socket_port",
					Description: `(Optional) TCP Socket detection port.`,
				},
				resource.Attribute{
					Name:        "timeout_seconds",
					Description: `(Optional) Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is ` + "`" + `1` + "`" + `. The ` + "`" + `volume_mount` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Volume name.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) Volume mount path.`,
				},
				resource.Attribute{
					Name:        "mount_propagation",
					Description: `(Optional) Volume mount propagation.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether the volume is read-only.`,
				},
				resource.Attribute{
					Name:        "sub_path_expr",
					Description: `(Optional) Volume mount sub-path expression.`,
				},
				resource.Attribute{
					Name:        "sub_path",
					Description: `(Optional) Volume mount sub-path. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "auto_create_eip_id",
					Description: `ID of EIP which create automatically.`,
				},
				resource.Attribute{
					Name:        "created_time",
					Description: `Container instance creation time.`,
				},
				resource.Attribute{
					Name:        "eip_address",
					Description: `EIP address.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Private IP address.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Container instance status. ## Import ` + "`" + `` + "`" + `` + "`" + ` terraform import tencentcloud_eks_container_instance.foo container-instance-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "auto_create_eip_id",
					Description: `ID of EIP which create automatically.`,
				},
				resource.Attribute{
					Name:        "created_time",
					Description: `Container instance creation time.`,
				},
				resource.Attribute{
					Name:        "eip_address",
					Description: `EIP address.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Private IP address.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Container instance status. ## Import ` + "`" + `` + "`" + `` + "`" + ` terraform import tencentcloud_eks_container_instance.foo container-instance-id ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "node_info_list",
					Description: `(Required) Node information list, which is used to describe the specification information of various types of nodes in the cluster, such as node type, node quantity, node specification, disk type, and disk size.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Password to an instance.`,
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
					Name:        "availability_zone",
					Description: `(Optional, ForceNew) Availability zone. When create multi-az es, this parameter must be omitted.`,
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
					Name:        "subnet_id",
					Description: `(Optional, ForceNew) The ID of a VPC subnetwork. When create multi-az es, this parameter must be omitted.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the instance. For tag limits, please refer to [Use Limits](https://intl.cloud.tencent.com/document/product/651/13354).`,
				},
				resource.Attribute{
					Name:        "web_node_type_info",
					Description: `(Optional) Visual node configuration. The ` + "`" + `multi_zone_infos` + "`" + ` object supports the following:`,
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
					Description: `(Optional) Node type. Valid values are ` + "`" + `hotData` + "`" + `, ` + "`" + `warmData` + "`" + ` and ` + "`" + `dedicatedMaster` + "`" + `. The default value is 'hotData` + "`" + `. The ` + "`" + `web_node_type_info` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "node_num",
					Description: `(Required) Visual node number.`,
				},
				resource.Attribute{
					Name:        "node_type",
					Description: `(Required) Visual node specifications. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
			Type:             "tencentcloud_emr_cluster",
			Category:         "EMR",
			ShortDescription: `Provide a resource to create a emr cluster.`,
			Description:      ``,
			Keywords: []string{
				"emr",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_strategy",
					Description: `(Required, ForceNew) Display strategy of EMR instance.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Required, ForceNew) Name of the instance, which can contain 6 to 36 English letters, Chinese characters, digits, dashes(-), or underscores(_).`,
				},
				resource.Attribute{
					Name:        "login_settings",
					Description: `(Required, ForceNew) Instance login settings.`,
				},
				resource.Attribute{
					Name:        "pay_mode",
					Description: `(Required) The pay mode of instance. 0 is pay on an annual basis, 1 is pay on a measure basis.`,
				},
				resource.Attribute{
					Name:        "placement",
					Description: `(Required, ForceNew) The location of the instance.`,
				},
				resource.Attribute{
					Name:        "product_id",
					Description: `(Required, ForceNew) The product id of EMR instance.`,
				},
				resource.Attribute{
					Name:        "softwares",
					Description: `(Required, ForceNew) The softwares of a EMR instance.`,
				},
				resource.Attribute{
					Name:        "support_ha",
					Description: `(Required, ForceNew) The flag whether the instance support high availability.(0=>not support, 1=>support).`,
				},
				resource.Attribute{
					Name:        "time_span",
					Description: `(Required) The length of time the instance was purchased. Use with TimeUnit.When TimeUnit is s, the parameter can only be filled in at 3600, representing a metered instance. When TimeUnit is m, the number filled in by this parameter indicates the length of purchase of the monthly instance of the package year, such as 1 for one month of purchase.`,
				},
				resource.Attribute{
					Name:        "time_unit",
					Description: `(Required) The unit of time in which the instance was purchased. When PayMode is 0, TimeUnit can only take values of s(second). When PayMode is 1, TimeUnit can only take the value m(month).`,
				},
				resource.Attribute{
					Name:        "vpc_settings",
					Description: `(Required, ForceNew) The private net config of EMR instance.`,
				},
				resource.Attribute{
					Name:        "extend_fs_field",
					Description: `(Optional) Access the external file system.`,
				},
				resource.Attribute{
					Name:        "need_master_wan",
					Description: `(Optional, ForceNew) Whether to enable the cluster Master node public network. Value range: - NEED_MASTER_WAN: Indicates that the cluster Master node public network is enabled. - NOT_NEED_MASTER_WAN: Indicates that it is not turned on. By default, the cluster Master node internet is enabled.`,
				},
				resource.Attribute{
					Name:        "resource_spec",
					Description: `(Optional) Resource specification of EMR instance.`,
				},
				resource.Attribute{
					Name:        "sg_id",
					Description: `(Optional, ForceNew) The ID of the security group to which the instance belongs, in the form of sg-xxxxxxxx. The ` + "`" + `resource_spec` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "common_count",
					Description: `(Optional, ForceNew) The number of common node.`,
				},
				resource.Attribute{
					Name:        "common_resource_spec",
					Description: `(Optional, ForceNew)`,
				},
				resource.Attribute{
					Name:        "core_count",
					Description: `(Optional) The number of core node.`,
				},
				resource.Attribute{
					Name:        "core_resource_spec",
					Description: `(Optional, ForceNew)`,
				},
				resource.Attribute{
					Name:        "master_count",
					Description: `(Optional) The number of master node.`,
				},
				resource.Attribute{
					Name:        "master_resource_spec",
					Description: `(Optional, ForceNew)`,
				},
				resource.Attribute{
					Name:        "task_count",
					Description: `(Optional) The number of core node.`,
				},
				resource.Attribute{
					Name:        "task_resource_spec",
					Description: `(Optional, ForceNew) ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `Created EMR instance id.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `Created EMR instance id.`,
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
					Description: `(Optional) CA certificate domain of the realserver. It has been deprecated.`,
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
					Name:        "client_ip_method",
					Description: `(Optional, ForceNew) The way the listener gets the client IP, 0 for TOA, 1 for Proxy Protocol, default value is 0. NOTES: Only supports listeners of ` + "`" + `TCP` + "`" + ` protocol.`,
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
					Description: `(Required) The image to use for the instance. Changing ` + "`" + `image_id` + "`" + ` will cause the instance reset.`,
				},
				resource.Attribute{
					Name:        "allocate_public_ip",
					Description: `(Optional, ForceNew) Associate a public IP address with an instance in a VPC or Classic. Boolean value, Default is false.`,
				},
				resource.Attribute{
					Name:        "bandwidth_package_id",
					Description: `(Optional) bandwidth package id. if user is standard user, then the bandwidth_package_id is needed, or default has bandwidth_package_id.`,
				},
				resource.Attribute{
					Name:        "cam_role_name",
					Description: `(Optional, ForceNew) CAM role name authorized to access.`,
				},
				resource.Attribute{
					Name:        "cdh_host_id",
					Description: `(Optional, ForceNew) Id of cdh instance. Note: it only works when instance_charge_type is set to ` + "`" + `CDHPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cdh_instance_type",
					Description: `(Optional) Type of instance created on cdh, the value of this parameter is in the format of CDH_XCXG based on the number of CPU cores and memory capacity. Note: it only works when instance_charge_type is set to ` + "`" + `CDHPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "data_disks",
					Description: `(Optional, ForceNew) Settings for data disks.`,
				},
				resource.Attribute{
					Name:        "disable_monitor_service",
					Description: `(Optional) Disable enhance service for monitor, it is enabled by default. When this options is set, monitor agent won't be installed. Modifying will cause the instance reset.`,
				},
				resource.Attribute{
					Name:        "disable_security_service",
					Description: `(Optional) Disable enhance service for security, it is enabled by default. When this options is set, security agent won't be installed. Modifying will cause the instance reset.`,
				},
				resource.Attribute{
					Name:        "force_delete",
					Description: `(Optional) Indicate whether to force delete the instance. Default is ` + "`" + `false` + "`" + `. If set true, the instance will be permanently deleted instead of being moved into the recycle bin. Note: only works for ` + "`" + `PREPAID` + "`" + ` instance.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional) The hostname of the instance. Windows instance: The name should be a combination of 2 to 15 characters comprised of letters (case insensitive), numbers, and hyphens (-). Period (.) is not supported, and the name cannot be a string of pure numbers. Other types (such as Linux) of instances: The name should be a combination of 2 to 60 characters, supporting multiple periods (.). The piece between two periods is composed of letters (case insensitive), numbers, and hyphens (-). Modifying will cause the instance reset.`,
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
					Description: `(Optional) The charge type of instance. Valid values are ` + "`" + `PREPAID` + "`" + `, ` + "`" + `POSTPAID_BY_HOUR` + "`" + `, ` + "`" + `SPOTPAID` + "`" + ` and ` + "`" + `CDHPAID` + "`" + `. The default is ` + "`" + `POSTPAID_BY_HOUR` + "`" + `. Note: TencentCloud International only supports ` + "`" + `POSTPAID_BY_HOUR` + "`" + ` and ` + "`" + `CDHPAID` + "`" + `. ` + "`" + `PREPAID` + "`" + ` instance may not allow to delete before expired. ` + "`" + `SPOTPAID` + "`" + ` instance must set ` + "`" + `spot_instance_type` + "`" + ` and ` + "`" + `spot_max_price` + "`" + ` at the same time. ` + "`" + `CDHPAID` + "`" + ` instance must set ` + "`" + `cdh_instance_type` + "`" + ` and ` + "`" + `cdh_host_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_count",
					Description: `(Optional,`,
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
					Description: `(Optional) Whether to keep image login or not, default is ` + "`" + `false` + "`" + `. When the image type is private or shared or imported, this parameter can be set ` + "`" + `true` + "`" + `. Modifying will cause the instance reset.`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `(Optional) The key pair to use for the instance, it looks like ` + "`" + `skey-16jig7tx` + "`" + `. Modifying will cause the instance reset.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Password for the instance. In order for the new password to take effect, the instance will be restarted after the password change. Modifying will cause the instance reset.`,
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
					Name:        "stopped_mode",
					Description: `(Optional) Billing method of a pay-as-you-go instance after shutdown. Available values: ` + "`" + `KEEP_CHARGING` + "`" + `,` + "`" + `STOP_CHARGING` + "`" + `. Default ` + "`" + `KEEP_CHARGING` + "`" + `.`,
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
					Description: `(Optional) Size of the system disk. Valid value ranges: (50~1000). and unit is GB. Default is 50GB. If modified, the instance may force stop.`,
				},
				resource.Attribute{
					Name:        "system_disk_type",
					Description: `(Optional) System disk type. For more information on limits of system disk types, see [Storage Overview](https://intl.cloud.tencent.com/document/product/213/4952). Valid values: ` + "`" + `LOCAL_BASIC` + "`" + `: local disk, ` + "`" + `LOCAL_SSD` + "`" + `: local SSD disk, ` + "`" + `CLOUD_SSD` + "`" + `: SSD, ` + "`" + `CLOUD_PREMIUM` + "`" + `: Premium Cloud Storage. NOTE: 1. ` + "`" + `CLOUD_BASIC` + "`" + `, ` + "`" + `LOCAL_BASIC` + "`" + ` and ` + "`" + `LOCAL_SSD` + "`" + ` are deprecated; 2. If modified, the instance may force stop.`,
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
					Description: `(Required) Size of the data disk, and unit is GB. If disk type is ` + "`" + `CLOUD_SSD` + "`" + `, the size range is [100, 16000], and the others are [10-16000].`,
				},
				resource.Attribute{
					Name:        "data_disk_type",
					Description: `(Required, ForceNew) Data disk type. For more information about limits on different data disk types, see [Storage Overview](https://intl.cloud.tencent.com/document/product/213/4952). Valid values: ` + "`" + `LOCAL_BASIC` + "`" + `: local disk, ` + "`" + `LOCAL_SSD` + "`" + `: local SSD disk, ` + "`" + `CLOUD_PREMIUM` + "`" + `: Premium Cloud Storage, ` + "`" + `CLOUD_SSD` + "`" + `: SSD, ` + "`" + `CLOUD_HSSD` + "`" + `: Enhanced SSD. NOTE: ` + "`" + `CLOUD_BASIC` + "`" + `, ` + "`" + `LOCAL_BASIC` + "`" + ` and ` + "`" + `LOCAL_SSD` + "`" + ` are deprecated.`,
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
			Type:             "tencentcloud_kubernetes_addon_attachment",
			Category:         "Tencent Kubernetes Engine(TKE)",
			ShortDescription: `Provide a resource to configure kubernetes cluster app addons.`,
			Description:      ``,
			Keywords: []string{
				"tencent",
				"kubernetes",
				"engine",
				"tke",
				"addon",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required, ForceNew) ID of cluster.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Name of addon.`,
				},
				resource.Attribute{
					Name:        "request_body",
					Description: `(Optional) Serialized json string as request body of addon spec. If set, will ignore ` + "`" + `version` + "`" + ` and ` + "`" + `values` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Optional) Values the addon passthroughs. Conflict with ` + "`" + `request_body` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) Addon version, default latest version. Conflict with ` + "`" + `request_body` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "response_body",
					Description: `Addon response body.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Addon current status. ## Import Addon can be imported by using cluster_id#addon_name ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_kubernetes_addon_attachment.addon_cos cls-xxxxxxxx#cos ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "response_body",
					Description: `Addon response body.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Addon current status. ## Import Addon can be imported by using cluster_id#addon_name ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_kubernetes_addon_attachment.addon_cos cls-xxxxxxxx#cos ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Optional) Max bandwidth of Internet access in Mbps. Default is ` + "`" + `0` + "`" + `.`,
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
			Type:             "tencentcloud_kubernetes_auth_attachment",
			Category:         "Tencent Kubernetes Engine(TKE)",
			ShortDescription: `Provide a resource to configure kubernetes cluster authentication info.`,
			Description:      ``,
			Keywords: []string{
				"tencent",
				"kubernetes",
				"engine",
				"tke",
				"auth",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) ID of clusters.`,
				},
				resource.Attribute{
					Name:        "issuer",
					Description: `(Required) Specify service-account-issuer.`,
				},
				resource.Attribute{
					Name:        "auto_create_discovery_anonymous_auth",
					Description: `(Optional) If set to ` + "`" + `true` + "`" + `, the rbac rule will be created automatically which allow anonymous user to access '/.well-known/openid-configuration' and '/openid/v1/jwks'.`,
				},
				resource.Attribute{
					Name:        "jwks_uri",
					Description: `(Optional) Specify service-account-jwks-uri. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
					Name:        "acquire_cluster_admin_role",
					Description: `(Optional) If set to true, it will acquire the ClusterRole tke:admin. NOTE: this arguments cannot revoke to ` + "`" + `false` + "`" + ` after acquired.`,
				},
				resource.Attribute{
					Name:        "auth_options",
					Description: `(Optional) Specify cluster authentication configuration. Only available for managed cluster and ` + "`" + `cluster_version` + "`" + ` >= 1.20.`,
				},
				resource.Attribute{
					Name:        "auto_upgrade_cluster_level",
					Description: `(Optional) Whether the cluster level auto upgraded, valid for managed cluster.`,
				},
				resource.Attribute{
					Name:        "base_pod_num",
					Description: `(Optional, ForceNew) The number of basic pods. valid when enable_customized_pod_cidr=true.`,
				},
				resource.Attribute{
					Name:        "claim_expired_seconds",
					Description: `(Optional) Claim expired seconds to recycle ENI. This field can only set when field ` + "`" + `network_type` + "`" + ` is 'VPC-CNI'. ` + "`" + `claim_expired_seconds` + "`" + ` must greater or equal than 300 and less than 15768000.`,
				},
				resource.Attribute{
					Name:        "cluster_as_enabled",
					Description: `(Optional, ForceNew) Indicates whether to enable cluster node auto scaling. Default is false.`,
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
					Description: `(Optional) Open internet access or not. If this field is set 'true', the field below ` + "`" + `worker_config` + "`" + ` must be set. Because only cluster with node is allowed enable access endpoint.`,
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
					Name:        "cluster_level",
					Description: `(Optional) Specify cluster level, valid for managed cluster, use data source ` + "`" + `tencentcloud_kubernetes_cluster_levels` + "`" + ` to query available levels. Available value examples ` + "`" + `L5` + "`" + `, ` + "`" + `LL20` + "`" + `, ` + "`" + `L50` + "`" + `, ` + "`" + `L100` + "`" + `, etc.`,
				},
				resource.Attribute{
					Name:        "cluster_max_pod_num",
					Description: `(Optional, ForceNew) The maximum number of Pods per node in the cluster. Default is 256. The minimum value is 4. When its power unequal to 2, it will round upward to the closest power of 2.`,
				},
				resource.Attribute{
					Name:        "cluster_max_service_num",
					Description: `(Optional, ForceNew) The maximum number of services in the cluster. Default is 256. The range is from 32 to 32768. When its power unequal to 2, it will round upward to the closest power of 2.`,
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
					Description: `(Optional, ForceNew) Operating system of the cluster, the available values include: 'centos7.6.0_x64','ubuntu18.04.1x86_64','tlinux2.4x86_64'. Default is 'tlinux2.4x86_64'.`,
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
					Name:        "enable_customized_pod_cidr",
					Description: `(Optional) Whether to enable the custom mode of node podCIDR size. Default is false.`,
				},
				resource.Attribute{
					Name:        "eni_subnet_ids",
					Description: `(Optional) Subnet Ids for cluster with VPC-CNI network mode. This field can only set when field ` + "`" + `network_type` + "`" + ` is 'VPC-CNI'. ` + "`" + `eni_subnet_ids` + "`" + ` can not empty once be set.`,
				},
				resource.Attribute{
					Name:        "exist_instance",
					Description: `(Optional, ForceNew) create tke cluster by existed instances.`,
				},
				resource.Attribute{
					Name:        "extension_addon",
					Description: `(Optional, ForceNew) Information of the add-on to be installed.`,
				},
				resource.Attribute{
					Name:        "extra_args",
					Description: `(Optional, ForceNew) Custom parameter information related to the node.`,
				},
				resource.Attribute{
					Name:        "globe_desired_pod_num",
					Description: `(Optional, ForceNew) Indicate to set desired pod number in node. valid when enable_customized_pod_cidr=true, and it takes effect for all nodes.`,
				},
				resource.Attribute{
					Name:        "ignore_cluster_cidr_conflict",
					Description: `(Optional, ForceNew) Indicates whether to ignore the cluster cidr conflict error. Default is false.`,
				},
				resource.Attribute{
					Name:        "is_non_static_ip_mode",
					Description: `(Optional, ForceNew) Indicates whether non-static ip mode is enabled. Default is false.`,
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
					Name:        "node_pool_global_config",
					Description: `(Optional) Global config effective for all node pools.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) Project ID, default value is 0.`,
				},
				resource.Attribute{
					Name:        "runtime_version",
					Description: `(Optional) Container Runtime version.`,
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
					Name:        "upgrade_instances_follow_cluster",
					Description: `(Optional) Indicates whether upgrade all instances when cluster_version change. Default is false.`,
				},
				resource.Attribute{
					Name:        "worker_config",
					Description: `(Optional, ForceNew) Deploy the machine configuration information of the 'WORKER' service, and create <=20 units for common users. The other 'WORK' service are added by 'tencentcloud_kubernetes_worker'. The ` + "`" + `auth_options` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "auto_create_discovery_anonymous_auth",
					Description: `(Optional) If set to ` + "`" + `true` + "`" + `, the rbac rule will be created automatically which allow anonymous user to access '/.well-known/openid-configuration' and '/openid/v1/jwks'.`,
				},
				resource.Attribute{
					Name:        "issuer",
					Description: `(Optional) Specify service-account-issuer.`,
				},
				resource.Attribute{
					Name:        "jwks_uri",
					Description: `(Optional) Specify service-account-jwks-uri. The ` + "`" + `cluster_extra_args` + "`" + ` object supports the following:`,
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
					Name:        "auto_format_and_mount",
					Description: `(Optional, ForceNew) Indicate whether to auto format and mount or not. Default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disk_partition",
					Description: `(Optional, ForceNew) The name of the device or partition to mount.`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `(Optional, ForceNew) Volume of disk in GB. Default is ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `(Optional, ForceNew) Types of disk, available values: ` + "`" + `CLOUD_PREMIUM` + "`" + ` and ` + "`" + `CLOUD_SSD` + "`" + ` and ` + "`" + `CLOUD_HSSD` + "`" + ` and ` + "`" + `CLOUD_TSSD` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "file_system",
					Description: `(Optional, ForceNew) File system, e.g. ` + "`" + `ext3/ext4/xfs` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mount_target",
					Description: `(Optional, ForceNew) Mount target.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional, ForceNew) Data disk snapshot ID. The ` + "`" + `exist_instance` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "desired_pod_numbers",
					Description: `(Optional, ForceNew) Custom mode cluster, you can specify the number of pods for each node. corresponding to the existed_instances_para.instance_ids parameter.`,
				},
				resource.Attribute{
					Name:        "instances_para",
					Description: `(Optional, ForceNew) Reinstallation parameters of an existing instance.`,
				},
				resource.Attribute{
					Name:        "node_role",
					Description: `(Optional, ForceNew) Role of existed node. value:MASTER_ETCD or WORKER. The ` + "`" + `extension_addon` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Add-on name.`,
				},
				resource.Attribute{
					Name:        "param",
					Description: `(Required) Description of the add-on resource object in JSON string format. The ` + "`" + `instances_para` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "instance_ids",
					Description: `(Required, ForceNew) Cluster IDs. The ` + "`" + `master_config` + "`" + ` object supports the following:`,
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
					Name:        "bandwidth_package_id",
					Description: `(Optional) bandwidth package id. if user is standard user, then the bandwidth_package_id is needed, or default has bandwidth_package_id.`,
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
					Name:        "desired_pod_num",
					Description: `(Optional, ForceNew) Indicate to set desired pod number in node. valid when enable_customized_pod_cidr=true, and it override ` + "`" + `[globe_]desired_pod_num` + "`" + ` for current node. Either all the fields ` + "`" + `desired_pod_num` + "`" + ` or none.`,
				},
				resource.Attribute{
					Name:        "disaster_recover_group_ids",
					Description: `(Optional, ForceNew) Disaster recover groups to which a CVM instance belongs. Only support maximum 1.`,
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
					Name:        "img_id",
					Description: `(Optional) The valid image id, format of img-xxx.`,
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
					Description: `(Optional) Max bandwidth of Internet access in Mbps. Default is 0.`,
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
					Description: `(Optional, ForceNew) System disk type. For more information on limits of system disk types, see [Storage Overview](https://intl.cloud.tencent.com/document/product/213/4952). Valid values: ` + "`" + `LOCAL_BASIC` + "`" + `: local disk, ` + "`" + `LOCAL_SSD` + "`" + `: local SSD disk, ` + "`" + `CLOUD_SSD` + "`" + `: SSD, ` + "`" + `CLOUD_PREMIUM` + "`" + `: Premium Cloud Storage. NOTE: ` + "`" + `CLOUD_BASIC` + "`" + `, ` + "`" + `LOCAL_BASIC` + "`" + ` and ` + "`" + `LOCAL_SSD` + "`" + ` are deprecated.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional, ForceNew) ase64-encoded User Data text, the length limit is 16KB. The ` + "`" + `node_pool_global_config` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "expander",
					Description: `(Optional) Indicates which scale-out method will be used when there are multiple scaling groups. Valid values: ` + "`" + `random` + "`" + ` - select a random scaling group, ` + "`" + `most-pods` + "`" + ` - select the scaling group that can schedule the most pods, ` + "`" + `least-waste` + "`" + ` - select the scaling group that can ensure the fewest remaining resources after Pod scheduling.`,
				},
				resource.Attribute{
					Name:        "ignore_daemon_sets_utilization",
					Description: `(Optional) Whether to ignore DaemonSet pods by default when calculating resource usage.`,
				},
				resource.Attribute{
					Name:        "is_scale_in_enabled",
					Description: `(Optional) Indicates whether to enable scale-in.`,
				},
				resource.Attribute{
					Name:        "max_concurrent_scale_in",
					Description: `(Optional) Max concurrent scale-in volume.`,
				},
				resource.Attribute{
					Name:        "scale_in_delay",
					Description: `(Optional) Number of minutes after cluster scale-out when the system starts judging whether to perform scale-in.`,
				},
				resource.Attribute{
					Name:        "scale_in_unneeded_time",
					Description: `(Optional) Number of consecutive minutes of idleness after which the node is subject to scale-in.`,
				},
				resource.Attribute{
					Name:        "scale_in_utilization_threshold",
					Description: `(Optional) Percentage of node resource usage below which the node is considered to be idle.`,
				},
				resource.Attribute{
					Name:        "skip_nodes_with_local_storage",
					Description: `(Optional) During scale-in, ignore nodes with local storage pods.`,
				},
				resource.Attribute{
					Name:        "skip_nodes_with_system_pods",
					Description: `(Optional) During scale-in, ignore nodes with pods in the kube-system namespace that are not managed by DaemonSet. The ` + "`" + `worker_config` + "`" + ` object supports the following:`,
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
					Name:        "bandwidth_package_id",
					Description: `(Optional) bandwidth package id. if user is standard user, then the bandwidth_package_id is needed, or default has bandwidth_package_id.`,
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
					Name:        "desired_pod_num",
					Description: `(Optional, ForceNew) Indicate to set desired pod number in node. valid when enable_customized_pod_cidr=true, and it override ` + "`" + `[globe_]desired_pod_num` + "`" + ` for current node. Either all the fields ` + "`" + `desired_pod_num` + "`" + ` or none.`,
				},
				resource.Attribute{
					Name:        "disaster_recover_group_ids",
					Description: `(Optional, ForceNew) Disaster recover groups to which a CVM instance belongs. Only support maximum 1.`,
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
					Name:        "img_id",
					Description: `(Optional) The valid image id, format of img-xxx.`,
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
					Description: `(Optional) Max bandwidth of Internet access in Mbps. Default is 0.`,
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
					Description: `(Optional, ForceNew) System disk type. For more information on limits of system disk types, see [Storage Overview](https://intl.cloud.tencent.com/document/product/213/4952). Valid values: ` + "`" + `LOCAL_BASIC` + "`" + `: local disk, ` + "`" + `LOCAL_SSD` + "`" + `: local SSD disk, ` + "`" + `CLOUD_SSD` + "`" + `: SSD, ` + "`" + `CLOUD_PREMIUM` + "`" + `: Premium Cloud Storage. NOTE: ` + "`" + `CLOUD_BASIC` + "`" + `, ` + "`" + `LOCAL_BASIC` + "`" + ` and ` + "`" + `LOCAL_SSD` + "`" + ` are deprecated.`,
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
					Name:        "worker_config_overrides",
					Description: `(Optional, ForceNew) Override variable worker_config, commonly used to attach existing instances.`,
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
					Name:        "disk_partition",
					Description: `(Optional, ForceNew) The name of the device or partition to mount. NOTE: this argument doesn't support setting in node pool, or will leads to mount error.`,
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
					Description: `(Optional, ForceNew) Mount target. The ` + "`" + `worker_config_overrides` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "data_disk",
					Description: `(Optional, ForceNew) Configurations of data disk.`,
				},
				resource.Attribute{
					Name:        "desired_pod_num",
					Description: `(Optional, ForceNew) Indicate to set desired pod number in node. valid when the cluster is podCIDR.`,
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
					Description: `(Optional, ForceNew) Base64-encoded User Data text, the length limit is 16KB. The ` + "`" + `worker_config` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "data_disk",
					Description: `(Optional, ForceNew) Configurations of data disk.`,
				},
				resource.Attribute{
					Name:        "desired_pod_num",
					Description: `(Optional, ForceNew) Indicate to set desired pod number in node. valid when the cluster is podCIDR.`,
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
					Description: `(Required) Auto scaling config parameters.`,
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
					Description: `(Required) Name of the node pool. The name does not exceed 25 characters, and only supports Chinese, English, numbers, underscores, separators (` + "`" + `-` + "`" + `) and decimal points.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, ForceNew) ID of VPC network.`,
				},
				resource.Attribute{
					Name:        "default_cooldown",
					Description: `(Optional) Seconds of scaling group cool down. Default value is ` + "`" + `300` + "`" + `.`,
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
					Name:        "multi_zone_subnet_policy",
					Description: `(Optional, ForceNew) Multi-availability zone/subnet policy. Valid values: PRIORITY and EQUALITY. Default value: PRIORITY.`,
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
					Name:        "scaling_group_name",
					Description: `(Optional) Name of relative scaling group.`,
				},
				resource.Attribute{
					Name:        "scaling_group_project_id",
					Description: `(Optional) Project ID the scaling group belongs to.`,
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
					Name:        "termination_policies",
					Description: `(Optional) Policy of scaling group termination. Available values: ` + "`" + `["OLDEST_INSTANCE"]` + "`" + `, ` + "`" + `["NEWEST_INSTANCE"]` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "unschedulable",
					Description: `(Optional, ForceNew) Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `(Optional) List of auto scaling group available zones, for Basic network it is required. The ` + "`" + `auto_scaling_config` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Required, ForceNew) Specified types of CVM instance.`,
				},
				resource.Attribute{
					Name:        "backup_instance_types",
					Description: `(Optional) Backup CVM instance types if specified instance type sold out or mismatch.`,
				},
				resource.Attribute{
					Name:        "bandwidth_package_id",
					Description: `(Optional) bandwidth package id. if user is standard user, then the bandwidth_package_id is needed, or default has bandwidth_package_id.`,
				},
				resource.Attribute{
					Name:        "cam_role_name",
					Description: `(Optional, ForceNew) Name of cam role.`,
				},
				resource.Attribute{
					Name:        "data_disk",
					Description: `(Optional) Configurations of data disk.`,
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
					Name:        "instance_charge_type_prepaid_period",
					Description: `(Optional) The tenancy (in month) of the prepaid instance, NOTE: it only works when instance_charge_type is set to ` + "`" + `PREPAID` + "`" + `. Valid values are ` + "`" + `1` + "`" + `, ` + "`" + `2` + "`" + `, ` + "`" + `3` + "`" + `, ` + "`" + `4` + "`" + `, ` + "`" + `5` + "`" + `, ` + "`" + `6` + "`" + `, ` + "`" + `7` + "`" + `, ` + "`" + `8` + "`" + `, ` + "`" + `9` + "`" + `, ` + "`" + `10` + "`" + `, ` + "`" + `11` + "`" + `, ` + "`" + `12` + "`" + `, ` + "`" + `24` + "`" + `, ` + "`" + `36` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type_prepaid_renew_flag",
					Description: `(Optional) Auto renewal flag. Valid values: ` + "`" + `NOTIFY_AND_AUTO_RENEW` + "`" + `: notify upon expiration and renew automatically, ` + "`" + `NOTIFY_AND_MANUAL_RENEW` + "`" + `: notify upon expiration but do not renew automatically, ` + "`" + `DISABLE_NOTIFY_AND_MANUAL_RENEW` + "`" + `: neither notify upon expiration nor renew automatically. Default value: ` + "`" + `NOTIFY_AND_MANUAL_RENEW` + "`" + `. If this parameter is specified as ` + "`" + `NOTIFY_AND_AUTO_RENEW` + "`" + `, the instance will be automatically renewed on a monthly basis if the account balance is sufficient. NOTE: it only works when instance_charge_type is set to ` + "`" + `PREPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `(Optional) Charge type of instance. Valid values are ` + "`" + `PREPAID` + "`" + `, ` + "`" + `POSTPAID_BY_HOUR` + "`" + `, ` + "`" + `SPOTPAID` + "`" + `. The default is ` + "`" + `POSTPAID_BY_HOUR` + "`" + `. NOTE: ` + "`" + `SPOTPAID` + "`" + ` instance must set ` + "`" + `spot_instance_type` + "`" + ` and ` + "`" + `spot_max_price` + "`" + ` at the same time.`,
				},
				resource.Attribute{
					Name:        "internet_charge_type",
					Description: `(Optional) Charge types for network traffic. Valid value: ` + "`" + `BANDWIDTH_PREPAID` + "`" + `, ` + "`" + `TRAFFIC_POSTPAID_BY_HOUR` + "`" + `, ` + "`" + `TRAFFIC_POSTPAID_BY_HOUR` + "`" + ` and ` + "`" + `BANDWIDTH_PACKAGE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "internet_max_bandwidth_out",
					Description: `(Optional) Max bandwidth of Internet access in Mbps. Default is ` + "`" + `0` + "`" + `.`,
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
					Description: `(Optional) Specify whether to assign an Internet IP address.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Optional) Security groups to which a CVM instance belongs.`,
				},
				resource.Attribute{
					Name:        "spot_instance_type",
					Description: `(Optional) Type of spot instance, only support ` + "`" + `one-time` + "`" + ` now. Note: it only works when instance_charge_type is set to ` + "`" + `SPOTPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "spot_max_price",
					Description: `(Optional) Max price of a spot instance, is the format of decimal string, for example "0.50". Note: it only works when instance_charge_type is set to ` + "`" + `SPOTPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "system_disk_size",
					Description: `(Optional) Volume of system disk in GB. Default is ` + "`" + `50` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "system_disk_type",
					Description: `(Optional) Type of a CVM disk. Valid value: ` + "`" + `CLOUD_PREMIUM` + "`" + ` and ` + "`" + `CLOUD_SSD` + "`" + `. Default is ` + "`" + `CLOUD_PREMIUM` + "`" + `. The ` + "`" + `data_disk` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "auto_format_and_mount",
					Description: `(Optional, ForceNew) Indicate whether to auto format and mount or not. Default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disk_partition",
					Description: `(Optional, ForceNew) The name of the device or partition to mount. NOTE: this argument doesn't support setting in node pool, or will leads to mount error.`,
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
					Name:        "delete_with_instance",
					Description: `(Optional) Indicates whether the disk remove after instance terminated.`,
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
					Name:        "desired_pod_num",
					Description: `(Optional, ForceNew) Indicate to set desired pod number in node. valid when the cluster is podCIDR.`,
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
					Name:        "desired_pod_num",
					Description: `(Optional, ForceNew) Indicate to set desired pod number in current node. Valid when the cluster enable customized pod cidr.`,
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
					Name:        "disk_partition",
					Description: `(Optional, ForceNew) The name of the device or partition to mount.`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `(Optional, ForceNew) Volume of disk in GB. Default is ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `(Optional, ForceNew) Types of disk, available values: ` + "`" + `CLOUD_PREMIUM` + "`" + ` and ` + "`" + `CLOUD_SSD` + "`" + ` and ` + "`" + `CLOUD_HSSD` + "`" + ` and ` + "`" + `CLOUD_TSSD` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "file_system",
					Description: `(Optional, ForceNew) File system, e.g. ` + "`" + `ext3/ext4/xfs` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mount_target",
					Description: `(Optional, ForceNew) Mount target.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional, ForceNew) Data disk snapshot ID. The ` + "`" + `data_disk` + "`" + ` object supports the following:`,
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
					Description: `(Optional, ForceNew) Types of disk, available values: ` + "`" + `CLOUD_PREMIUM` + "`" + ` and ` + "`" + `CLOUD_SSD` + "`" + ` and ` + "`" + `CLOUD_HSSD` + "`" + ` and ` + "`" + `CLOUD_TSSD` + "`" + `.`,
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
					Name:        "bandwidth_package_id",
					Description: `(Optional) bandwidth package id. if user is standard user, then the bandwidth_package_id is needed, or default has bandwidth_package_id.`,
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
					Name:        "desired_pod_num",
					Description: `(Optional, ForceNew) Indicate to set desired pod number in node. valid when enable_customized_pod_cidr=true, and it override ` + "`" + `[globe_]desired_pod_num` + "`" + ` for current node. Either all the fields ` + "`" + `desired_pod_num` + "`" + ` or none.`,
				},
				resource.Attribute{
					Name:        "disaster_recover_group_ids",
					Description: `(Optional, ForceNew) Disaster recover groups to which a CVM instance belongs. Only support maximum 1.`,
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
					Name:        "img_id",
					Description: `(Optional) The valid image id, format of img-xxx.`,
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
					Description: `(Optional) Max bandwidth of Internet access in Mbps. Default is 0.`,
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
					Description: `(Optional, ForceNew) System disk type. For more information on limits of system disk types, see [Storage Overview](https://intl.cloud.tencent.com/document/product/213/4952). Valid values: ` + "`" + `LOCAL_BASIC` + "`" + `: local disk, ` + "`" + `LOCAL_SSD` + "`" + `: local SSD disk, ` + "`" + `CLOUD_SSD` + "`" + `: SSD, ` + "`" + `CLOUD_PREMIUM` + "`" + `: Premium Cloud Storage. NOTE: ` + "`" + `CLOUD_BASIC` + "`" + `, ` + "`" + `LOCAL_BASIC` + "`" + ` and ` + "`" + `LOCAL_SSD` + "`" + ` are deprecated.`,
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
					Description: `(Required, ForceNew) Version of the Mongodb, and available values include ` + "`" + `MONGO_36_WT` + "`" + ` (MongoDB 3.6 WiredTiger Edition), ` + "`" + `MONGO_40_WT` + "`" + ` (MongoDB 4.0 WiredTiger Edition) and ` + "`" + `MONGO_42_WT` + "`" + ` (MongoDB 4.2 WiredTiger Edition). NOTE: ` + "`" + `MONGO_3_WT` + "`" + ` (MongoDB 3.2 WiredTiger Edition) and ` + "`" + `MONGO_3_ROCKS` + "`" + ` (MongoDB 3.2 RocksDB Edition) will deprecated.`,
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
					Description: `(Required, ForceNew) Version of the Mongodb, and available values include ` + "`" + `MONGO_36_WT` + "`" + ` (MongoDB 3.6 WiredTiger Edition), ` + "`" + `MONGO_40_WT` + "`" + ` (MongoDB 4.0 WiredTiger Edition) and ` + "`" + `MONGO_42_WT` + "`" + ` (MongoDB 4.2 WiredTiger Edition). NOTE: ` + "`" + `MONGO_3_WT` + "`" + ` (MongoDB 3.2 WiredTiger Edition) and ` + "`" + `MONGO_3_ROCKS` + "`" + ` (MongoDB 3.2 RocksDB Edition) will deprecated.`,
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
			Type:             "tencentcloud_monitor_alarm_policy",
			Category:         "Monitor",
			ShortDescription: `Provides a alarm policy resource for monitor.`,
			Description:      ``,
			Keywords: []string{
				"monitor",
				"alarm",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "monitor_type",
					Description: `(Required, ForceNew) The type of monitor.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Required, ForceNew) The type of alarm.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `(Required) The name of policy.`,
				},
				resource.Attribute{
					Name:        "conditions",
					Description: `(Optional) A list of metric trigger condition.`,
				},
				resource.Attribute{
					Name:        "conditon_template_id",
					Description: `(Optional, ForceNew) ID of trigger condition template.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional) Whether to enable, default is ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "event_conditions",
					Description: `(Optional) A list of event trigger condition.`,
				},
				resource.Attribute{
					Name:        "notice_ids",
					Description: `(Optional) List of notification rule IDs.`,
				},
				resource.Attribute{
					Name:        "policy_tag",
					Description: `(Optional, ForceNew) Policy tag to bind object.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional, ForceNew) Project ID. For products with different projects, a value other than -1 must be passed in.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `(Optional) The remark of policy group.`,
				},
				resource.Attribute{
					Name:        "trigger_tasks",
					Description: `(Optional) Triggered task list. The ` + "`" + `conditions` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "is_union_rule",
					Description: `(Optional) The and or relation of indicator alarm rule.`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `(Optional) A list of metric trigger condition. The ` + "`" + `event_conditions` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "continue_period",
					Description: `(Optional) Number of periods.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Metric display name, which is used in the output parameter.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Filter condition for one single trigger rule. Must set it when create tke-xxx rules.`,
				},
				resource.Attribute{
					Name:        "is_power_notice",
					Description: `(Optional) Whether the alarm frequency increases exponentially.`,
				},
				resource.Attribute{
					Name:        "metric_name",
					Description: `(Optional) Metric name or event name.`,
				},
				resource.Attribute{
					Name:        "notice_frequency",
					Description: `(Optional) Alarm interval in seconds.`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `(Optional) Operator.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) Statistical period in seconds.`,
				},
				resource.Attribute{
					Name:        "rule_type",
					Description: `(Optional) Trigger condition type.`,
				},
				resource.Attribute{
					Name:        "unit",
					Description: `(Optional) Unit, which is used in the output parameter.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Threshold. The ` + "`" + `filter` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "dimensions",
					Description: `(Optional) JSON string generated by serializing the AlarmPolicyDimension two-dimensional array.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Filter condition type. Valid values: DIMENSION (uses dimensions for filtering). The ` + "`" + `policy_tag` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Tag key.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Tag value. The ` + "`" + `rules` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "continue_period",
					Description: `(Optional) Number of periods.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Metric display name, which is used in the output parameter.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Filter condition for one single trigger rule. Must set it when create tke-xxx rules.`,
				},
				resource.Attribute{
					Name:        "is_power_notice",
					Description: `(Optional) Whether the alarm frequency increases exponentially.`,
				},
				resource.Attribute{
					Name:        "metric_name",
					Description: `(Optional) Metric name or event name.`,
				},
				resource.Attribute{
					Name:        "notice_frequency",
					Description: `(Optional) Alarm interval in seconds.`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `(Optional) Operator.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) Statistical period in seconds.`,
				},
				resource.Attribute{
					Name:        "rule_type",
					Description: `(Optional) Trigger condition type.`,
				},
				resource.Attribute{
					Name:        "unit",
					Description: `(Optional) Unit, which is used in the output parameter.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Threshold. The ` + "`" + `trigger_tasks` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "task_config",
					Description: `(Required) Configuration information in JSON format.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Triggered task type. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The alarm policy create time.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `The alarm policy update time. ## Import Alarm policy instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_monitor_alarm_policy.policy policy-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The alarm policy create time.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `The alarm policy update time. ## Import Alarm policy instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_monitor_alarm_policy.policy policy-id ` + "`" + `` + "`" + `` + "`" + ``,
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
			Type:             "tencentcloud_monitor_policy_binding_object",
			Category:         "Monitor",
			ShortDescription: `Provides a resource for bind objects to a alarm policy resource.`,
			Description:      ``,
			Keywords: []string{
				"monitor",
				"policy",
				"binding",
				"object",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dimensions",
					Description: `(Required, ForceNew) A list objects. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Required, ForceNew) Alarm policy ID for binding objects. The ` + "`" + `dimensions` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "dimensions_json",
					Description: `(Required, ForceNew) Represents a collection of dimensions of an object instance, json format.eg:'{"unInstanceId":"ins-ot3cq4bi"}'. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import Monitor Policy Binding Object can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_monitor_policy_binding_object.binding policyId ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import Monitor Policy Binding Object can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_monitor_policy_binding_object.binding policyId ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "cpu",
					Description: `(Optional) CPU cores.`,
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
					Description: `(Required) Global privileges. available values for Privileges:ALTER,ALTER ROUTINE,CREATE,CREATE ROUTINE,CREATE TEMPORARY TABLES,CREATE USER,CREATE VIEW,DELETE,DROP,EVENT,EXECUTE,INDEX,INSERT,LOCK TABLES,PROCESS,REFERENCES,RELOAD,REPLICATION CLIENT,REPLICATION SLAVE,SELECT,SHOW DATABASES,SHOW VIEW,TRIGGER,UPDATE.`,
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
					Name:        "cpu",
					Description: `(Optional) CPU cores.`,
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
					Name:        "master_region",
					Description: `(Optional, ForceNew) The zone information of the primary instance is required when you purchase a disaster recovery instance.`,
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
					Description: `(Optional) ID of VPC, which can be modified once every 24 hours and can't be removed.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional, ForceNew) Zone information, this parameter defaults to, the system automatically selects an Availability Zone. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
			Type:             "tencentcloud_nat_gateway_snat",
			Category:         "Virtual Private Cloud(VPC)",
			ShortDescription: `Provides a resource to create a NAT Gateway SNat rule.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"nat",
				"gateway",
				"snat",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Required) Description.`,
				},
				resource.Attribute{
					Name:        "nat_gateway_id",
					Description: `(Required, ForceNew) NAT gateway ID.`,
				},
				resource.Attribute{
					Name:        "public_ip_addr",
					Description: `(Required) Elastic IP address pool.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `(Required, ForceNew) Resource type. Valid values: SUBNET, NETWORKINTERFACE.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Optional, ForceNew) Instance ID, required when ` + "`" + `resource_type` + "`" + ` is NETWORKINTERFACE.`,
				},
				resource.Attribute{
					Name:        "instance_private_ip_addr",
					Description: `(Optional, ForceNew) Private IPs of the instance's primary ENI, required when ` + "`" + `resource_type` + "`" + ` is NETWORKINTERFACE.`,
				},
				resource.Attribute{
					Name:        "subnet_cidr_block",
					Description: `(Optional, ForceNew) The IPv4 CIDR of the subnet, required when ` + "`" + `resource_type` + "`" + ` is SUBNET.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional, ForceNew) Subnet instance ID, required when ` + "`" + `resource_type` + "`" + ` is SUBNET. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
					Name:        "snat_id",
					Description: `SNAT rule ID. ## Import VPN gateway route can be imported using the id, the id format must be '{nat_gateway_id}#{resource_id}', resource_id range ` + "`" + `subnet_id` + "`" + `, ` + "`" + `instance_id` + "`" + `, e.g. SUBNET SNat ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_nat_gateway_snat.my_snat nat-r4ip1cwt#subnet-2ap74y35 ` + "`" + `` + "`" + `` + "`" + ` NETWORKINTERFACT SNat ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_nat_gateway_snat.my_snat nat-r4ip1cwt#ins-da412f5a ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "snat_id",
					Description: `SNAT rule ID. ## Import VPN gateway route can be imported using the id, the id format must be '{nat_gateway_id}#{resource_id}', resource_id range ` + "`" + `subnet_id` + "`" + `, ` + "`" + `instance_id` + "`" + `, e.g. SUBNET SNat ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_nat_gateway_snat.my_snat nat-r4ip1cwt#subnet-2ap74y35 ` + "`" + `` + "`" + `` + "`" + ` NETWORKINTERFACT SNat ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_nat_gateway_snat.my_snat nat-r4ip1cwt#ins-da412f5a ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "availability_zone",
					Description: `(Required, ForceNew) Availability zone. NOTE: If value modified but included in ` + "`" + `db_node_set` + "`" + `, the diff will be suppressed.`,
				},
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
					Name:        "backup_plan",
					Description: `(Optional) Specify DB backup plan.`,
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
					Name:        "db_kernel_version",
					Description: `(Optional) PostgreSQL kernel version number. If it is specified, an instance running kernel DBKernelVersion will be created.`,
				},
				resource.Attribute{
					Name:        "db_major_vesion",
					Description: `(Optional) PostgreSQL major version number. Valid values: 10, 11, 12, 13. If it is specified, an instance running the latest kernel of PostgreSQL DBMajorVersion will be created.`,
				},
				resource.Attribute{
					Name:        "db_node_set",
					Description: `(Optional) Specify instance node info for disaster migration.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `(Optional, ForceNew) Version of the postgresql database engine. Valid values: ` + "`" + `10.4` + "`" + `, ` + "`" + `11.8` + "`" + `, ` + "`" + `12.4` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "kms_key_id",
					Description: `(Optional) KeyId of the custom key.`,
				},
				resource.Attribute{
					Name:        "kms_region",
					Description: `(Optional) Region of the custom key.`,
				},
				resource.Attribute{
					Name:        "max_standby_archive_delay",
					Description: `(Optional) max_standby_archive_delay applies when WAL data is being read from WAL archive (and is therefore not current). Units are milliseconds if not specified.`,
				},
				resource.Attribute{
					Name:        "max_standby_streaming_delay",
					Description: `(Optional) max_standby_streaming_delay applies when WAL data is being received via streaming replication. Units are milliseconds if not specified.`,
				},
				resource.Attribute{
					Name:        "need_support_tde",
					Description: `(Optional) Whether to support data transparent encryption, 1: yes, 0: no (default).`,
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
					Name:        "root_user",
					Description: `(Optional, ForceNew) Instance root account name. This parameter is optional, Default value is ` + "`" + `root` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Optional) ID of security group. If both vpc_id and subnet_id are not set, this argument should not be set either.`,
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
					Description: `(Optional, ForceNew) ID of VPC. The ` + "`" + `backup_plan` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "backup_period",
					Description: `(Optional) List of backup period per week, available values: ` + "`" + `monday` + "`" + `, ` + "`" + `tuesday` + "`" + `, ` + "`" + `wednesday` + "`" + `, ` + "`" + `thursday` + "`" + `, ` + "`" + `friday` + "`" + `, ` + "`" + `saturday` + "`" + `, ` + "`" + `sunday` + "`" + `. NOTE: At least specify two days.`,
				},
				resource.Attribute{
					Name:        "base_backup_retention_period",
					Description: `(Optional) Specify days of the retention.`,
				},
				resource.Attribute{
					Name:        "max_backup_start_time",
					Description: `(Optional) Specify latest backup start time, format ` + "`" + `hh:mm:ss` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "min_backup_start_time",
					Description: `(Optional) Specify earliest backup start time, format ` + "`" + `hh:mm:ss` + "`" + `. The ` + "`" + `db_node_set` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) Indicates the node available zone.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Optional) Indicates node type, available values:` + "`" + `Primary` + "`" + `, ` + "`" + `Standby` + "`" + `. Default: ` + "`" + `Standby` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
					Description: `Port for public access.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `Uid of the postgresql instance. ## Import postgresql instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_postgresql_instance.foo postgres-cda1iex1 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `Port for public access.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `Uid of the postgresql instance. ## Import postgresql instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_postgresql_instance.foo postgres-cda1iex1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_postgresql_readonly_attachment",
			Category:         "PostgreSQL",
			ShortDescription: `Use this resource to create postgresql readonly attachment.`,
			Description:      ``,
			Keywords: []string{
				"postgresql",
				"readonly",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "db_instance_id",
					Description: `(Required, ForceNew) Read only instance ID.`,
				},
				resource.Attribute{
					Name:        "read_only_group_id",
					Description: `(Required, ForceNew) Read only group ID. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
			Type:             "tencentcloud_postgresql_readonly_group",
			Category:         "PostgreSQL",
			ShortDescription: `Use this resource to create postgresql readonly group.`,
			Description:      ``,
			Keywords: []string{
				"postgresql",
				"readonly",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "master_db_instance_id",
					Description: `(Required, ForceNew) Primary instance ID.`,
				},
				resource.Attribute{
					Name:        "max_replay_lag",
					Description: `(Required) Delay threshold in ms.`,
				},
				resource.Attribute{
					Name:        "max_replay_latency",
					Description: `(Required) Delayed log size threshold in MB.`,
				},
				resource.Attribute{
					Name:        "min_delay_eliminate_reserve",
					Description: `(Required) The minimum number of read-only replicas that must be retained in an RO group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) RO group name.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) Project ID.`,
				},
				resource.Attribute{
					Name:        "replay_lag_eliminate",
					Description: `(Required) Whether to remove a read-only replica from an RO group if the delay between the read-only replica and the primary instance exceeds the threshold. Valid values: 0 (no), 1 (yes).`,
				},
				resource.Attribute{
					Name:        "replay_latency_eliminate",
					Description: `(Required) Whether to remove a read-only replica from an RO group if the sync log size difference between the read-only replica and the primary instance exceeds the threshold. Valid values: 0 (no), 1 (yes).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) VPC subnet ID.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, ForceNew) VPC ID.`,
				},
				resource.Attribute{
					Name:        "security_groups_ids",
					Description: `(Optional) ID of security group. If both vpc_id and subnet_id are not set, this argument should not be set either. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the postgresql instance.`,
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
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_postgresql_readonly_instance",
			Category:         "PostgreSQL",
			ShortDescription: `Use this resource to create postgresql readonly instance.`,
			Description:      ``,
			Keywords: []string{
				"postgresql",
				"readonly",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "db_version",
					Description: `(Required, ForceNew) PostgreSQL kernel version, which must be the same as that of the primary instance.`,
				},
				resource.Attribute{
					Name:        "master_db_instance_id",
					Description: `(Required, ForceNew) ID of the primary instance to which the read-only replica belongs.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `(Required) Memory size(in GB). Allowed value must be larger than ` + "`" + `memory` + "`" + ` that data source ` + "`" + `tencentcloud_postgresql_specinfos` + "`" + ` provides.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Instance name.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) Project ID.`,
				},
				resource.Attribute{
					Name:        "security_groups_ids",
					Description: `(Required) ID of security group.`,
				},
				resource.Attribute{
					Name:        "storage",
					Description: `(Required) Instance storage capacity in GB.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) VPC subnet ID.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, ForceNew) VPC ID.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Required, ForceNew) Availability zone ID, which can be obtained through the Zone field in the returned value of the DescribeZones API.`,
				},
				resource.Attribute{
					Name:        "auto_renew_flag",
					Description: `(Optional, ForceNew) Renewal flag. Valid values: 0 (manual renewal), 1 (auto-renewal). Default value: 0.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `(Optional, ForceNew) instance billing mode. Valid values: PREPAID (monthly subscription), POSTPAID_BY_HOUR (pay-as-you-go).`,
				},
				resource.Attribute{
					Name:        "need_support_ipv6",
					Description: `(Optional, ForceNew) Whether to support IPv6 address access. Valid values: 1 (yes), 0 (no). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the postgresql instance. ## Import postgresql readonly instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_postgresql_readonly_instance.foo pgro-bcqx8b9a ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the postgresql instance. ## Import postgresql readonly instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_postgresql_readonly_instance.foo pgro-bcqx8b9a ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_private_dns_record",
			Category:         "PrivateDNS",
			ShortDescription: `Provide a resource to create a Private Dns Record.`,
			Description:      ``,
			Keywords: []string{
				"privatedns",
				"private",
				"dns",
				"record",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "record_type",
					Description: `(Required) Record type. Valid values: "A", "AAAA", "CNAME", "MX", "TXT", "PTR".`,
				},
				resource.Attribute{
					Name:        "record_value",
					Description: `(Required) Record value, such as IP: 192.168.10.2, CNAME: cname.qcloud.com, and MX: mail.qcloud.com..`,
				},
				resource.Attribute{
					Name:        "sub_domain",
					Description: `(Required) Subdomain, such as "www", "m", and "@".`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Required) Private domain ID.`,
				},
				resource.Attribute{
					Name:        "mx",
					Description: `(Optional) MX priority, which is required when the record type is MX. Valid values: 5, 10, 15, 20, 30, 40, 50.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) Record cache time. The smaller the value, the faster the record will take effect. Value range: 1~86400s.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) Record weight. Value range: 1~100. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import Private Dns Record can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_private_dns_zone.foo zone_id#record_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import Private Dns Record can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_private_dns_zone.foo zone_id#record_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_private_dns_zone",
			Category:         "PrivateDNS",
			ShortDescription: `Provide a resource to create a Private Dns Zone.`,
			Description:      ``,
			Keywords: []string{
				"privatedns",
				"private",
				"dns",
				"zone",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) Domain name, which must be in the format of standard TLD.`,
				},
				resource.Attribute{
					Name:        "account_vpc_set",
					Description: `(Optional) List of authorized accounts' VPCs to associate with the private domain.`,
				},
				resource.Attribute{
					Name:        "dns_forward_status",
					Description: `(Optional) Whether to enable subdomain recursive DNS. Valid values: ENABLED, DISABLED. Default value: DISABLED.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `(Optional) Remarks.`,
				},
				resource.Attribute{
					Name:        "tag_set",
					Description: `(Optional) Tags the private domain when it is created.`,
				},
				resource.Attribute{
					Name:        "vpc_set",
					Description: `(Optional) Associates the private domain to a VPC when it is created. The ` + "`" + `account_vpc_set` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Region.`,
				},
				resource.Attribute{
					Name:        "uin",
					Description: `(Required) UIN of the VPC account.`,
				},
				resource.Attribute{
					Name:        "uniq_vpc_id",
					Description: `(Required) VPC ID.`,
				},
				resource.Attribute{
					Name:        "vpc_name",
					Description: `(Required) VPC NAME. The ` + "`" + `tag_set` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "tag_key",
					Description: `(Required) Key of Tag.`,
				},
				resource.Attribute{
					Name:        "tag_value",
					Description: `(Required) Value of Tag. The ` + "`" + `vpc_set` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) VPC REGION.`,
				},
				resource.Attribute{
					Name:        "uniq_vpc_id",
					Description: `(Required) VPC ID. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import Private Dns Zone can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_private_dns_zone.foo zone_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import Private Dns Zone can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_private_dns_zone.foo zone_id ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "backup_time",
					Description: `(Required) Specifys what time the backup action should take place. And the time interval should be one hour.`,
				},
				resource.Attribute{
					Name:        "redis_id",
					Description: `(Required, ForceNew) ID of a redis instance to which the policy will be applied.`,
				},
				resource.Attribute{
					Name:        "backup_period",
					Description: `(Optional,`,
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
					Description: `(Required) The memory volume of an available instance(in MB), please refer to ` + "`" + `tencentcloud_redis_zone_config.list[zone].shard_memories` + "`" + `. When redis is standard type, it represents total memory size of the instance; when Redis is cluster type, it represents memory size of per sharding.`,
				},
				resource.Attribute{
					Name:        "auto_renew_flag",
					Description: `(Optional, ForceNew) Auto-renew flag. 0 - default state (manual renewal); 1 - automatic renewal; 2 - explicit no automatic renewal.`,
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
					Name:        "no_auth",
					Description: `(Optional, ForceNew) Indicates whether the redis instance support no-auth access. NOTE: Only available in private cloud environment.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Password for a Redis user, which should be 8 to 16 characters. NOTE: Only ` + "`" + `no_auth=true` + "`" + ` specified can make password empty.`,
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
					Description: `(Optional) The number of instance copies. This is not required for standalone and master slave versions.`,
				},
				resource.Attribute{
					Name:        "redis_shard_num",
					Description: `(Optional, ForceNew) The number of instance shard. This is not required for standalone and master slave versions.`,
				},
				resource.Attribute{
					Name:        "replica_zone_ids",
					Description: `(Optional) ID of replica nodes available zone. This is not required for standalone and master slave versions.`,
				},
				resource.Attribute{
					Name:        "replicas_read_only",
					Description: `(Optional, ForceNew) Whether copy read-only is supported, Redis 2.8 Standard Edition and CKV Standard Edition do not support replica read-only, turn on replica read-only, the instance will automatically read and write separate, write requests are routed to the primary node, read requests are routed to the replica node, if you need to open replica read-only, the recommended number of replicas >=2.`,
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
					Description: `(Required) Number of reserved instances to be purchased.`,
				},
				resource.Attribute{
					Name:        "reserved_instance_name",
					Description: `(Optional) Reserved Instance display name. - If you do not specify an instance display name, 'Unnamed' is displayed by default. - Up to 60 characters (including pattern strings) are supported. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
					Description: `(Required, ForceNew) The next hop type. Valid values: ` + "`" + `public_gateway` + "`" + `,` + "`" + `vpn_gateway` + "`" + `,` + "`" + `sslvpn_gateway` + "`" + `,` + "`" + `dc_gateway` + "`" + `,` + "`" + `peering_connection` + "`" + `,` + "`" + `nat_gateway` + "`" + `,` + "`" + `havip` + "`" + `,` + "`" + `local_gateway` + "`" + ` and ` + "`" + `instance` + "`" + `. ` + "`" + `instance` + "`" + ` points to CVM Instance.`,
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
					Name:        "cfs_config",
					Description: `(Optional) List of CFS configurations.`,
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
					Name:        "enable_eip_config",
					Description: `(Optional) Indicates whether EIP config set to ` + "`" + `ENABLE` + "`" + ` when ` + "`" + `enable_public_net` + "`" + ` was true.`,
				},
				resource.Attribute{
					Name:        "enable_public_net",
					Description: `(Optional) Indicates whether public net config enabled. NOTE: only ` + "`" + `vpc_id` + "`" + ` specified can disable public net config.`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `(Optional) Environment of the SCF function.`,
				},
				resource.Attribute{
					Name:        "image_config",
					Description: `(Optional) Image of the SCF function, conflict with ` + "`" + `` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "l5_enable",
					Description: `(Optional) Enable L5 for SCF function, default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "layers",
					Description: `(Optional) The list of association layers.`,
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
					Description: `(Optional) Zip file of the SCF function, conflict with ` + "`" + `cos_bucket_name` + "`" + `, ` + "`" + `cos_object_name` + "`" + `, ` + "`" + `cos_bucket_region` + "`" + `. The ` + "`" + `cfs_config` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "cfs_id",
					Description: `(Required) File system instance ID.`,
				},
				resource.Attribute{
					Name:        "local_mount_dir",
					Description: `(Required) Local mount directory.`,
				},
				resource.Attribute{
					Name:        "mount_ins_id",
					Description: `(Required) File system mount instance ID.`,
				},
				resource.Attribute{
					Name:        "remote_mount_dir",
					Description: `(Required) Remote mount directory.`,
				},
				resource.Attribute{
					Name:        "user_group_id",
					Description: `(Required) ID of user group.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `(Required) ID of user. The ` + "`" + `image_config` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "image_type",
					Description: `(Required) The image type. personal or enterprise.`,
				},
				resource.Attribute{
					Name:        "image_uri",
					Description: `(Required) The uri of image.`,
				},
				resource.Attribute{
					Name:        "args",
					Description: `(Optional) the parameters of command.`,
				},
				resource.Attribute{
					Name:        "command",
					Description: `(Optional) The command of entrypoint.`,
				},
				resource.Attribute{
					Name:        "entry_point",
					Description: `(Optional) The entrypoint of app.`,
				},
				resource.Attribute{
					Name:        "registry_id",
					Description: `(Optional) The registry id of TCR. When image type is enterprise, it must be set. The ` + "`" + `layers` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "layer_name",
					Description: `(Required) The name of Layer.`,
				},
				resource.Attribute{
					Name:        "layer_version",
					Description: `(Required) The version of layer. The ` + "`" + `triggers` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the SCF function trigger, if ` + "`" + `type` + "`" + ` is ` + "`" + `ckafka` + "`" + `, the format of name must be ` + "`" + `<ckafkaInstanceId>-<topicId>` + "`" + `; if ` + "`" + `type` + "`" + ` is ` + "`" + `cos` + "`" + `, the name is cos bucket id, other In any case, it can be combined arbitrarily. It can only contain English letters, numbers, connectors and underscores. The maximum length is 100.`,
				},
				resource.Attribute{
					Name:        "trigger_desc",
					Description: `(Required) TriggerDesc of the SCF function trigger, parameter format of ` + "`" + `timer` + "`" + ` is linux cron expression; parameter of ` + "`" + `cos` + "`" + ` type is json string ` + "`" + `{"bucketUrl":"<name-appid>.cos.<region>.myqcloud.com","event":"cos:ObjectCreated:`,
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
			Type:             "tencentcloud_scf_layer",
			Category:         "Serverless Cloud Function(SCF)",
			ShortDescription: `Provide a resource to create a SCF layer.`,
			Description:      ``,
			Keywords: []string{
				"serverless",
				"cloud",
				"function",
				"scf",
				"layer",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "compatible_runtimes",
					Description: `(Required) The compatible runtimes of layer.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Required) The source code of layer.`,
				},
				resource.Attribute{
					Name:        "layer_name",
					Description: `(Required) The name of layer.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of layer.`,
				},
				resource.Attribute{
					Name:        "license_info",
					Description: `(Optional) The license info of layer. The ` + "`" + `content` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "cos_bucket_name",
					Description: `(Optional) Cos bucket name of the SCF layer, such as ` + "`" + `cos-1234567890` + "`" + `, conflict with ` + "`" + `zip_file` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cos_bucket_region",
					Description: `(Optional) Cos bucket region of the SCF layer, conflict with ` + "`" + `zip_file` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cos_object_name",
					Description: `(Optional) Cos object name of the SCF layer, should have suffix ` + "`" + `.zip` + "`" + ` or ` + "`" + `.jar` + "`" + `, conflict with ` + "`" + `zip_file` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "zip_file",
					Description: `(Optional) Zip file of the SCF layer, conflict with ` + "`" + `cos_bucket_name` + "`" + `, ` + "`" + `cos_object_name` + "`" + `, ` + "`" + `cos_bucket_region` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "code_sha_256",
					Description: `The code type of layer.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The create time of layer.`,
				},
				resource.Attribute{
					Name:        "layer_version",
					Description: `The version of layer.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The download location url of layer.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current status of layer. ## Import Scf layer can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_scf_layer.layer layerId#layerVersion ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "code_sha_256",
					Description: `The code type of layer.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The create time of layer.`,
				},
				resource.Attribute{
					Name:        "layer_version",
					Description: `The version of layer.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The download location url of layer.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current status of layer. ## Import Scf layer can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_scf_layer.layer layerId#layerVersion ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Optional) Egress rules set. A rule must match the following format: [action]#[source]#[port]#[protocol]. The available value of 'action' is ` + "`" + `ACCEPT` + "`" + ` and ` + "`" + `DROP` + "`" + `. The 'source' can be an IP address network, segment, security group ID and Address Template ID. The 'port' valid format is ` + "`" + `80` + "`" + `, ` + "`" + `80,443` + "`" + `, ` + "`" + `80-90` + "`" + ` or ` + "`" + `ALL` + "`" + `. The available value of 'protocol' is ` + "`" + `TCP` + "`" + `, ` + "`" + `UDP` + "`" + `, ` + "`" + `ICMP` + "`" + ` and ` + "`" + `ALL` + "`" + `. When 'protocol' is ` + "`" + `ICMP` + "`" + ` or ` + "`" + `ALL` + "`" + `, the 'port' must be ` + "`" + `ALL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ingress",
					Description: `(Optional) Ingress rules set. A rule must match the following format: [action]#[source]#[port]#[protocol]. The available value of 'action' is ` + "`" + `ACCEPT` + "`" + ` and ` + "`" + `DROP` + "`" + `. The 'source' can be an IP address network, segment, security group ID and Address Template ID. The 'port' valid format is ` + "`" + `80` + "`" + `, ` + "`" + `80,443` + "`" + `, ` + "`" + `80-90` + "`" + ` or ` + "`" + `ALL` + "`" + `. The available value of 'protocol' is ` + "`" + `TCP` + "`" + `, ` + "`" + `UDP` + "`" + `, ` + "`" + `ICMP` + "`" + ` and ` + "`" + `ALL` + "`" + `. When 'protocol' is ` + "`" + `ICMP` + "`" + ` or ` + "`" + `ALL` + "`" + `, the 'port' must be ` + "`" + `ALL` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
					Name:        "auto_renew",
					Description: `(Optional) Automatic renewal sign. 0 for normal renewal, 1 for automatic renewal (Default). Only valid when purchasing a prepaid instance.`,
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
					Description: `(Optional, ForceNew) Pay type of the SQL Server instance. Available values ` + "`" + `PREPAID` + "`" + `, ` + "`" + `POSTPAID_BY_HOUR` + "`" + `.`,
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
					Name:        "period",
					Description: `(Optional) Purchase instance period in month. The value does not exceed 48.`,
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
					Name:        "auto_voucher",
					Description: `(Optional) Whether to use the voucher automatically; 1 for yes, 0 for no, the default is 0.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional, ForceNew) Availability zone.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `(Optional, ForceNew) Pay type of the SQL Server instance. Available values ` + "`" + `PREPAID` + "`" + `, ` + "`" + `POSTPAID_BY_HOUR` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "force_upgrade",
					Description: `(Optional, ForceNew) Indicate that the master instance upgrade or not. ` + "`" + `true` + "`" + ` for upgrading the master SQL Server instance to cluster type by force. Default is false. Note: this is not supported with ` + "`" + `DUAL` + "`" + `(ha_type), ` + "`" + `2017` + "`" + `(engine_version) master SQL Server instance, for it will cause ha_type of the master SQL Server instance change.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) Purchase instance period in month. The value does not exceed 48.`,
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
					Description: `(Optional) Name of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) Project ID of the SSL certificate. Default is ` + "`" + `0` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
					Name:        "security_policy",
					Description: `(Optional) Public network access allowlist policies of the TCR instance. Only available when ` + "`" + `open_public_operation` + "`" + ` is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The available tags within this TCR instance. The ` + "`" + `security_policy` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Optional) The public network IP address of the access source.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Remarks of policy. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
					Name:        "instance_id",
					Description: `(Required, ForceNew) ID of the TCR instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Name of the TCR repository. Valid length is [2~200]. It can only contain lowercase letters, numbers and separators (` + "`" + `.` + "`" + `, ` + "`" + `_` + "`" + `, ` + "`" + `-` + "`" + `, ` + "`" + `/` + "`" + `), and cannot start, end or continue with separators. Support the use of multi-level address formats, such as ` + "`" + `sub1/sub2/repo` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "namespace_name",
					Description: `(Required, ForceNew) Name of the TCR namespace.`,
				},
				resource.Attribute{
					Name:        "brief_desc",
					Description: `(Optional) Brief description of the repository. Valid length is [1~100].`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the repository. Valid length is [1~1000]. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
					Description: `(Optional) Whether to enable vpc domain dns. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `(Optional) ID of region. Conflict with region_name, can not be set at the same time.`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `(Optional) Name of region. Conflict with region_id, can not be set at the same time. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
					Description: `Status of the internal access. ## Import tcr vpc attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_tcr_vpc_attachment.foo tcrId#vpcId#subnetId ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `Status of the internal access. ## Import tcr vpc attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_tcr_vpc_attachment.foo tcrId#vpcId#subnetId ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_tdmq_instance",
			Category:         "TDMQ",
			ShortDescription: `Provide a resource to create a TDMQ instance.`,
			Description:      ``,
			Keywords: []string{
				"tdmq",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Required) The name of tdmq cluster to be created.`,
				},
				resource.Attribute{
					Name:        "bind_cluster_id",
					Description: `(Optional) The Dedicated Cluster Id.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `(Optional) Description of the tdmq cluster. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import Tdmq instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_tdmq_instance.test tdmq_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import Tdmq instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_tdmq_instance.test tdmq_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_tdmq_namespace",
			Category:         "TDMQ",
			ShortDescription: `Provide a resource to create a tdmq namespace.`,
			Description:      ``,
			Keywords: []string{
				"tdmq",
				"namespace",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) The Dedicated Cluster Id.`,
				},
				resource.Attribute{
					Name:        "environ_name",
					Description: `(Required) The name of namespace to be created.`,
				},
				resource.Attribute{
					Name:        "msg_ttl",
					Description: `(Required) The expiration time of unconsumed message.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `(Optional) Description of the namespace.`,
				},
				resource.Attribute{
					Name:        "retention_policy",
					Description: `(Optional) The Policy of message to retain. The ` + "`" + `retention_policy` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "size_in_mb",
					Description: `(Optional) the size of message to retain.`,
				},
				resource.Attribute{
					Name:        "time_in_minutes",
					Description: `(Optional) the time of message to retain. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import Tdmq namespace can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_tdmq_instance.test namespace_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import Tdmq namespace can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_tdmq_instance.test namespace_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_tdmq_namespace_role_attachment",
			Category:         "TDMQ",
			ShortDescription: `Provide a resource to create a TDMQ role.`,
			Description:      ``,
			Keywords: []string{
				"tdmq",
				"namespace",
				"role",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) The id of tdmq cluster.`,
				},
				resource.Attribute{
					Name:        "environ_id",
					Description: `(Required) The name of tdmq namespace.`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `(Required) The permissions of tdmq role.`,
				},
				resource.Attribute{
					Name:        "role_name",
					Description: `(Required) The name of tdmq role. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of resource.`,
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
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_tdmq_role",
			Category:         "TDMQ",
			ShortDescription: `Provide a resource to create a TDMQ role.`,
			Description:      ``,
			Keywords: []string{
				"tdmq",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) The id of tdmq cluster.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `(Required) The description of tdmq role.`,
				},
				resource.Attribute{
					Name:        "role_name",
					Description: `(Required) The name of tdmq role. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import Tdmq instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_tdmq_instance.test tdmq_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import Tdmq instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_tdmq_instance.test tdmq_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_tdmq_topic",
			Category:         "TDMQ",
			ShortDescription: `Provide a resource to create a TDMQ topic.`,
			Description:      ``,
			Keywords: []string{
				"tdmq",
				"topic",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) The Dedicated Cluster Id.`,
				},
				resource.Attribute{
					Name:        "environ_id",
					Description: `(Required, ForceNew) The name of tdmq namespace.`,
				},
				resource.Attribute{
					Name:        "partitions",
					Description: `(Required) The partitions of topic.`,
				},
				resource.Attribute{
					Name:        "topic_name",
					Description: `(Required, ForceNew) The name of topic to be created.`,
				},
				resource.Attribute{
					Name:        "topic_type",
					Description: `(Required, ForceNew) The type of topic.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `(Optional) Description of the namespace. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of resource. ## Import Tdmq Topic can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_tdmq_topic.test topic_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of resource. ## Import Tdmq Topic can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_tdmq_topic.test topic_id ` + "`" + `` + "`" + `` + "`" + ``,
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
			Type:             "tencentcloud_vod_sub_application",
			Category:         "Video on Demand(VOD)",
			ShortDescription: `Provide a resource to create a VOD sub application.`,
			Description:      ``,
			Keywords: []string{
				"video",
				"on",
				"demand",
				"vod",
				"sub",
				"application",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Sub application name, which can contain up to 64 letters, digits, underscores, and hyphens (such as test_ABC-123) and must be unique under a user.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Required) Sub appliaction status.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Sub application description. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time when the sub application was created. ## Import VOD super player config can be imported using the name+, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_vod_sub_application.foo name+"#"+id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time when the sub application was created. ## Import VOD super player config can be imported using the name+, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_vod_sub_application.foo name+"#"+id ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "assistant_cidrs",
					Description: `(Optional) List of Assistant CIDR.`,
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
					Name:        "default_route_table_id",
					Description: `Default route table id, which created automatically after VPC create.`,
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
					Name:        "default_route_table_id",
					Description: `Default route table id, which created automatically after VPC create.`,
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
					Description: `(Optional) Proto authenticate algorithm of the IKE operation specification. Valid values: ` + "`" + `MD5` + "`" + `, ` + "`" + `SHA` + "`" + `, ` + "`" + `SHA-256` + "`" + `. Default Value is ` + "`" + `MD5` + "`" + `.`,
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
					Description: `(Optional) Integrity algorithm of the IPSEC operation specification. Valid values: ` + "`" + `SHA1` + "`" + `, ` + "`" + `MD5` + "`" + `, ` + "`" + `SHA-256` + "`" + `. Default value is ` + "`" + `MD5` + "`" + `.`,
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
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_vpn_gateway_route",
			Category:         "VPN",
			ShortDescription: `Provides a resource to create a VPN gateway route.`,
			Description:      ``,
			Keywords: []string{
				"vpn",
				"gateway",
				"route",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "destination_cidr_block",
					Description: `(Required, ForceNew) Destination IDC IP range.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, ForceNew) Instance ID of the next hop.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Required, ForceNew) Next hop type (type of the associated instance). Valid values: VPNCONN (VPN tunnel) and CCN (CCN instance).`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Required, ForceNew) Priority. Valid values: 0 and 100.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Required) Status. Valid values: ENABLE and DISABLE.`,
				},
				resource.Attribute{
					Name:        "vpn_gateway_id",
					Description: `(Required, ForceNew) VPN gateway ID. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
					Name:        "route_id",
					Description: `Route ID.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Route type. Default value: Static.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Update time. ## Import VPN gateway route can be imported using the id, the id format must be '{vpn_gateway_id}#{route_id}', e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_vpn_gateway_route.route1 vpngw-ak9sjem2#vpngw-8ccsnclt ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "route_id",
					Description: `Route ID.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Route type. Default value: Static.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Update time. ## Import VPN gateway route can be imported using the id, the id format must be '{vpn_gateway_id}#{route_id}', e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_vpn_gateway_route.route1 vpngw-ak9sjem2#vpngw-8ccsnclt ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_vpn_ssl_client",
			Category:         "VPN",
			ShortDescription: `Provide a resource to create a VPN SSL Client.`,
			Description:      ``,
			Keywords: []string{
				"vpn",
				"ssl",
				"client",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ssl_vpn_client_name",
					Description: `(Required, ForceNew) The name of ssl vpn client to be created.`,
				},
				resource.Attribute{
					Name:        "ssl_vpn_server_id",
					Description: `(Required, ForceNew) VPN ssl server id. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import VPN SSL Client can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_vpn_ssl_client.client vpn-client-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import VPN SSL Client can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_vpn_ssl_client.client vpn-client-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_vpn_ssl_server",
			Category:         "VPN",
			ShortDescription: `Provide a resource to create a VPN SSL Server.`,
			Description:      ``,
			Keywords: []string{
				"vpn",
				"ssl",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "local_address",
					Description: `(Required, ForceNew) List of local CIDR.`,
				},
				resource.Attribute{
					Name:        "remote_address",
					Description: `(Required, ForceNew) Remote CIDR for client.`,
				},
				resource.Attribute{
					Name:        "ssl_vpn_server_name",
					Description: `(Required, ForceNew) The name of ssl vpn server to be created.`,
				},
				resource.Attribute{
					Name:        "vpn_gateway_id",
					Description: `(Required, ForceNew) VPN gateway ID.`,
				},
				resource.Attribute{
					Name:        "compress",
					Description: `(Optional, ForceNew) need compressed. Default value: False.`,
				},
				resource.Attribute{
					Name:        "encrypt_algorithm",
					Description: `(Optional, ForceNew) The encrypt algorithm. Valid values: AES-128-CBC, AES-192-CBC, AES-256-CBC, NONE.Default value: NONE.`,
				},
				resource.Attribute{
					Name:        "integrity_algorithm",
					Description: `(Optional, ForceNew) The integrity algorithm. Valid values: SHA1, MD5 and NONE. Default value: NONE.`,
				},
				resource.Attribute{
					Name:        "ssl_vpn_port",
					Description: `(Optional, ForceNew) The port of ssl vpn. Default value: 1194.`,
				},
				resource.Attribute{
					Name:        "ssl_vpn_protocol",
					Description: `(Optional, ForceNew) The protocol of ssl vpn. Default value: UDP. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import VPN SSL Server can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_vpn_ssl_server.server vpn-server-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import VPN SSL Server can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_vpn_ssl_server.server vpn-server-id ` + "`" + `` + "`" + `` + "`" + ``,
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
		"tencentcloud_cam_oidc_sso":                            24,
		"tencentcloud_cam_policy":                              25,
		"tencentcloud_cam_role":                                26,
		"tencentcloud_cam_role_policy_attachment":              27,
		"tencentcloud_cam_saml_provider":                       28,
		"tencentcloud_cam_user":                                29,
		"tencentcloud_cam_user_policy_attachment":              30,
		"tencentcloud_cbs_snapshot":                            31,
		"tencentcloud_cbs_snapshot_policy":                     32,
		"tencentcloud_cbs_snapshot_policy_attachment":          33,
		"tencentcloud_cbs_storage":                             34,
		"tencentcloud_cbs_storage_attachment":                  35,
		"tencentcloud_ccn":                                     36,
		"tencentcloud_ccn_attachment":                          37,
		"tencentcloud_ccn_bandwidth_limit":                     38,
		"tencentcloud_cdh_instance":                            39,
		"tencentcloud_cdn_domain":                              40,
		"tencentcloud_cfs_access_group":                        41,
		"tencentcloud_cfs_access_rule":                         42,
		"tencentcloud_cfs_file_system":                         43,
		"tencentcloud_ckafka_acl":                              44,
		"tencentcloud_ckafka_instance":                         45,
		"tencentcloud_ckafka_topic":                            46,
		"tencentcloud_ckafka_user":                             47,
		"tencentcloud_clb_attachment":                          48,
		"tencentcloud_clb_customized_config":                   49,
		"tencentcloud_clb_instance":                            50,
		"tencentcloud_clb_listener":                            51,
		"tencentcloud_clb_listener_rule":                       52,
		"tencentcloud_clb_log_set":                             53,
		"tencentcloud_clb_log_topic":                           54,
		"tencentcloud_clb_redirection":                         55,
		"tencentcloud_clb_target_group":                        56,
		"tencentcloud_clb_target_group_attachment":             57,
		"tencentcloud_clb_target_group_instance_attachment":    58,
		"tencentcloud_cls_config":                              59,
		"tencentcloud_cls_config_attachment":                   60,
		"tencentcloud_cls_config_extra":                        61,
		"tencentcloud_cls_cos_shipper":                         62,
		"tencentcloud_cls_logset":                              63,
		"tencentcloud_cls_machine_group":                       64,
		"tencentcloud_cls_topic":                               65,
		"tencentcloud_container_cluster":                       66,
		"tencentcloud_container_cluster_instance":              67,
		"tencentcloud_cos_bucket":                              68,
		"tencentcloud_cos_bucket_object":                       69,
		"tencentcloud_cos_bucket_policy":                       70,
		"tencentcloud_cynosdb_cluster":                         71,
		"tencentcloud_cynosdb_readonly_instance":               72,
		"tencentcloud_dayu_cc_http_policy":                     73,
		"tencentcloud_dayu_cc_https_policy":                    74,
		"tencentcloud_dayu_cc_policy_v2":                       75,
		"tencentcloud_dayu_ddos_policy":                        76,
		"tencentcloud_dayu_ddos_policy_attachment":             77,
		"tencentcloud_dayu_ddos_policy_case":                   78,
		"tencentcloud_dayu_ddos_policy_v2":                     79,
		"tencentcloud_dayu_eip":                                80,
		"tencentcloud_dayu_l4_rule":                            81,
		"tencentcloud_dayu_l7_rule":                            82,
		"tencentcloud_dayu_l7_rule_v2":                         83,
		"tencentcloud_dc_gateway":                              84,
		"tencentcloud_dc_gateway_ccn_route":                    85,
		"tencentcloud_dcx":                                     86,
		"tencentcloud_dnat":                                    87,
		"tencentcloud_dnspod_domain_instance":                  88,
		"tencentcloud_dnspod_record":                           89,
		"tencentcloud_eip":                                     90,
		"tencentcloud_eip_association":                         91,
		"tencentcloud_eks_cluster":                             92,
		"tencentcloud_eks_container_instance":                  93,
		"tencentcloud_elasticsearch_instance":                  94,
		"tencentcloud_emr_cluster":                             95,
		"tencentcloud_eni":                                     96,
		"tencentcloud_eni_attachment":                          97,
		"tencentcloud_gaap_certificate":                        98,
		"tencentcloud_gaap_domain_error_page":                  99,
		"tencentcloud_gaap_http_domain":                        100,
		"tencentcloud_gaap_http_rule":                          101,
		"tencentcloud_gaap_layer4_listener":                    102,
		"tencentcloud_gaap_layer7_listener":                    103,
		"tencentcloud_gaap_proxy":                              104,
		"tencentcloud_gaap_realserver":                         105,
		"tencentcloud_gaap_security_policy":                    106,
		"tencentcloud_gaap_security_rule":                      107,
		"tencentcloud_ha_vip":                                  108,
		"tencentcloud_ha_vip_eip_attachment":                   109,
		"tencentcloud_image":                                   110,
		"tencentcloud_instance":                                111,
		"tencentcloud_key_pair":                                112,
		"tencentcloud_kms_external_key":                        113,
		"tencentcloud_kms_key":                                 114,
		"tencentcloud_kubernetes_addon_attachment":             115,
		"tencentcloud_kubernetes_as_scaling_group":             116,
		"tencentcloud_kubernetes_auth_attachment":              117,
		"tencentcloud_kubernetes_cluster":                      118,
		"tencentcloud_kubernetes_cluster_attachment":           119,
		"tencentcloud_kubernetes_node_pool":                    120,
		"tencentcloud_kubernetes_scale_worker":                 121,
		"tencentcloud_lb":                                      122,
		"tencentcloud_mongodb_instance":                        123,
		"tencentcloud_mongodb_sharding_instance":               124,
		"tencentcloud_mongodb_standby_instance":                125,
		"tencentcloud_monitor_alarm_policy":                    126,
		"tencentcloud_monitor_binding_object":                  127,
		"tencentcloud_monitor_binding_receiver":                128,
		"tencentcloud_monitor_policy_binding_object":           129,
		"tencentcloud_monitor_policy_group":                    130,
		"tencentcloud_mysql_account":                           131,
		"tencentcloud_mysql_account_privilege":                 132,
		"tencentcloud_mysql_backup_policy":                     133,
		"tencentcloud_mysql_instance":                          134,
		"tencentcloud_mysql_privilege":                         135,
		"tencentcloud_mysql_readonly_instance":                 136,
		"tencentcloud_nat_gateway":                             137,
		"tencentcloud_nat_gateway_snat":                        138,
		"tencentcloud_placement_group":                         139,
		"tencentcloud_postgresql_instance":                     140,
		"tencentcloud_postgresql_readonly_attachment":          141,
		"tencentcloud_postgresql_readonly_group":               142,
		"tencentcloud_postgresql_readonly_instance":            143,
		"tencentcloud_private_dns_record":                      144,
		"tencentcloud_private_dns_zone":                        145,
		"tencentcloud_protocol_template":                       146,
		"tencentcloud_protocol_template_group":                 147,
		"tencentcloud_redis_backup_config":                     148,
		"tencentcloud_redis_instance":                          149,
		"tencentcloud_reserved_instance":                       150,
		"tencentcloud_route_entry":                             151,
		"tencentcloud_route_table":                             152,
		"tencentcloud_route_table_entry":                       153,
		"tencentcloud_scf_function":                            154,
		"tencentcloud_scf_layer":                               155,
		"tencentcloud_scf_namespace":                           156,
		"tencentcloud_security_group":                          157,
		"tencentcloud_security_group_lite_rule":                158,
		"tencentcloud_security_group_rule":                     159,
		"tencentcloud_sqlserver_account":                       160,
		"tencentcloud_sqlserver_account_db_attachment":         161,
		"tencentcloud_sqlserver_basic_instance":                162,
		"tencentcloud_sqlserver_db":                            163,
		"tencentcloud_sqlserver_instance":                      164,
		"tencentcloud_sqlserver_publish_subscribe":             165,
		"tencentcloud_sqlserver_readonly_instance":             166,
		"tencentcloud_ssl_certificate":                         167,
		"tencentcloud_ssl_pay_certificate":                     168,
		"tencentcloud_ssm_secret":                              169,
		"tencentcloud_ssm_secret_version":                      170,
		"tencentcloud_subnet":                                  171,
		"tencentcloud_tcaplus_cluster":                         172,
		"tencentcloud_tcaplus_idl":                             173,
		"tencentcloud_tcaplus_table":                           174,
		"tencentcloud_tcaplus_tablegroup":                      175,
		"tencentcloud_tcr_instance":                            176,
		"tencentcloud_tcr_namespace":                           177,
		"tencentcloud_tcr_repository":                          178,
		"tencentcloud_tcr_token":                               179,
		"tencentcloud_tcr_vpc_attachment":                      180,
		"tencentcloud_tdmq_instance":                           181,
		"tencentcloud_tdmq_namespace":                          182,
		"tencentcloud_tdmq_namespace_role_attachment":          183,
		"tencentcloud_tdmq_role":                               184,
		"tencentcloud_tdmq_topic":                              185,
		"tencentcloud_vod_adaptive_dynamic_streaming_template": 186,
		"tencentcloud_vod_image_sprite_template":               187,
		"tencentcloud_vod_procedure_template":                  188,
		"tencentcloud_vod_snapshot_by_time_offset_template":    189,
		"tencentcloud_vod_sub_application":                     190,
		"tencentcloud_vod_super_player_config":                 191,
		"tencentcloud_vpc":                                     192,
		"tencentcloud_vpc_acl":                                 193,
		"tencentcloud_vpc_acl_attachment":                      194,
		"tencentcloud_vpn_connection":                          195,
		"tencentcloud_vpn_customer_gateway":                    196,
		"tencentcloud_vpn_gateway":                             197,
		"tencentcloud_vpn_gateway_route":                       198,
		"tencentcloud_vpn_ssl_client":                          199,
		"tencentcloud_vpn_ssl_server":                          200,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
