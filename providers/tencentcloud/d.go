package tencentcloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_address_template_groups",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of address template groups.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Id of the address template group to query.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the address template group to query.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "group_list",
					Description: `Information list of the dedicated address template groups.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of the address template group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of address template group.`,
				},
				resource.Attribute{
					Name:        "template_ids",
					Description: `ID set of the address template.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "group_list",
					Description: `Information list of the dedicated address template groups.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of the address template group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of address template group.`,
				},
				resource.Attribute{
					Name:        "template_ids",
					Description: `ID set of the address template.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_address_templates",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of address templates.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the address template to query.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the address template to query.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "template_list",
					Description: `Information list of the dedicated address templates.`,
				},
				resource.Attribute{
					Name:        "addresses",
					Description: `Set of the addresses.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the address template.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of address template.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "template_list",
					Description: `Information list of the dedicated address templates.`,
				},
				resource.Attribute{
					Name:        "addresses",
					Description: `Set of the addresses.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the address template.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of address template.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_api_gateway_api_keys",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query API gateway access keys.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "api_key_id",
					Description: `(Optional) Created API key ID, this field is exactly the same as ID.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "secret_name",
					Description: `(Optional) Custom key name. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list of API keys.`,
				},
				resource.Attribute{
					Name:        "access_key_secret",
					Description: `Created API key.`,
				},
				resource.Attribute{
					Name:        "api_key_id",
					Description: `API key ID.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time in the format of ` + "`" + `YYYY-MM-DDThh:mm:ssZ` + "`" + ` according to ISO 8601 standard. UTC time is used.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `Last modified time in the format of ` + "`" + `YYYY-MM-DDThh:mm:ssZ` + "`" + ` according to ISO 8601 standard. UTC time is used.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Key status. Values: ` + "`" + `on` + "`" + `, ` + "`" + `off` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list of API keys.`,
				},
				resource.Attribute{
					Name:        "access_key_secret",
					Description: `Created API key.`,
				},
				resource.Attribute{
					Name:        "api_key_id",
					Description: `API key ID.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time in the format of ` + "`" + `YYYY-MM-DDThh:mm:ssZ` + "`" + ` according to ISO 8601 standard. UTC time is used.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `Last modified time in the format of ` + "`" + `YYYY-MM-DDThh:mm:ssZ` + "`" + ` according to ISO 8601 standard. UTC time is used.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Key status. Values: ` + "`" + `on` + "`" + `, ` + "`" + `off` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_api_gateway_apis",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query API gateway APIs.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_id",
					Description: `(Required) Service ID for query.`,
				},
				resource.Attribute{
					Name:        "api_id",
					Description: `(Optional) Created API ID.`,
				},
				resource.Attribute{
					Name:        "api_name",
					Description: `(Optional) Custom API name.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list of APIs.`,
				},
				resource.Attribute{
					Name:        "api_desc",
					Description: `Custom API description.`,
				},
				resource.Attribute{
					Name:        "api_name",
					Description: `Custom API name.`,
				},
				resource.Attribute{
					Name:        "auth_type",
					Description: `API authentication type. Valid values: ` + "`" + `SECRET` + "`" + `, ` + "`" + `NONE` + "`" + `. ` + "`" + `SECRET` + "`" + ` means key pair authentication, ` + "`" + `NONE` + "`" + ` means no authentication.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.`,
				},
				resource.Attribute{
					Name:        "enable_cors",
					Description: `Whether to enable CORS.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `API frontend request type, such as ` + "`" + `HTTP` + "`" + `,` + "`" + `WEBSOCKET` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "request_config_method",
					Description: `Request frontend method configuration. Like ` + "`" + `GET` + "`" + `,` + "`" + `POST` + "`" + `,` + "`" + `PUT` + "`" + `,` + "`" + `DELETE` + "`" + `,` + "`" + `HEAD` + "`" + `,` + "`" + `ANY` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "request_config_path",
					Description: `Request frontend path configuration. Like ` + "`" + `/user/getinfo` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "request_parameters",
					Description: `Frontend request parameters.`,
				},
				resource.Attribute{
					Name:        "default_value",
					Description: `Parameter default value.`,
				},
				resource.Attribute{
					Name:        "desc",
					Description: `Parameter description.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Parameter name.`,
				},
				resource.Attribute{
					Name:        "position",
					Description: `Parameter location.`,
				},
				resource.Attribute{
					Name:        "required",
					Description: `If this parameter required.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Parameter type.`,
				},
				resource.Attribute{
					Name:        "response_error_codes",
					Description: `Custom error code configuration. Must keep at least one after set.`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `Custom response configuration error code.`,
				},
				resource.Attribute{
					Name:        "converted_code",
					Description: `Custom error code conversion.`,
				},
				resource.Attribute{
					Name:        "desc",
					Description: `Parameter description.`,
				},
				resource.Attribute{
					Name:        "msg",
					Description: `Custom response configuration error message.`,
				},
				resource.Attribute{
					Name:        "need_convert",
					Description: `Whether to enable error code conversion. Default value: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "response_fail_example",
					Description: `Response failure sample of custom response configuration.`,
				},
				resource.Attribute{
					Name:        "response_success_example",
					Description: `Successful response sample of custom response configuration.`,
				},
				resource.Attribute{
					Name:        "response_type",
					Description: `Return type.`,
				},
				resource.Attribute{
					Name:        "service_config_method",
					Description: `API backend service request method, such as ` + "`" + `GET` + "`" + `. If ` + "`" + `service_config_type` + "`" + ` is ` + "`" + `HTTP` + "`" + `, this parameter will be required. The frontend ` + "`" + `request_config_method` + "`" + ` and backend method ` + "`" + `service_config_method` + "`" + ` can be different.`,
				},
				resource.Attribute{
					Name:        "service_config_mock_return_message",
					Description: `Returned information of API backend mocking.`,
				},
				resource.Attribute{
					Name:        "service_config_path",
					Description: `API backend service path, such as /path. If ` + "`" + `service_config_type` + "`" + ` is ` + "`" + `HTTP` + "`" + `, this parameter will be required. The frontend ` + "`" + `request_config_path` + "`" + ` and backend path ` + "`" + `service_config_path` + "`" + ` can be different.`,
				},
				resource.Attribute{
					Name:        "service_config_product",
					Description: `Backend type. This parameter takes effect when VPC is enabled. Currently, only ` + "`" + `clb` + "`" + ` is supported.`,
				},
				resource.Attribute{
					Name:        "service_config_scf_function_name",
					Description: `SCF function name. This parameter takes effect when ` + "`" + `service_config_type` + "`" + ` is ` + "`" + `SCF` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "service_config_scf_function_namespace",
					Description: `SCF function namespace. This parameter takes effect when ` + "`" + `service_config_type` + "`" + ` is ` + "`" + `SCF` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "service_config_scf_function_qualifier",
					Description: `SCF function version. This parameter takes effect when ` + "`" + `service_config_type` + "`" + ` is ` + "`" + `SCF` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "service_config_timeout",
					Description: `API backend service timeout period in seconds.`,
				},
				resource.Attribute{
					Name:        "service_config_type",
					Description: `API backend service type.`,
				},
				resource.Attribute{
					Name:        "service_config_url",
					Description: `API backend service url. This parameter is required when ` + "`" + `service_config_type` + "`" + ` is ` + "`" + `HTTP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "service_config_vpc_id",
					Description: `Unique VPC ID.`,
				},
				resource.Attribute{
					Name:        "service_id",
					Description: `Which service this API belongs. Refer to resource ` + "`" + `tencentcloud_api_gateway_service` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list of APIs.`,
				},
				resource.Attribute{
					Name:        "api_desc",
					Description: `Custom API description.`,
				},
				resource.Attribute{
					Name:        "api_name",
					Description: `Custom API name.`,
				},
				resource.Attribute{
					Name:        "auth_type",
					Description: `API authentication type. Valid values: ` + "`" + `SECRET` + "`" + `, ` + "`" + `NONE` + "`" + `. ` + "`" + `SECRET` + "`" + ` means key pair authentication, ` + "`" + `NONE` + "`" + ` means no authentication.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.`,
				},
				resource.Attribute{
					Name:        "enable_cors",
					Description: `Whether to enable CORS.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `API frontend request type, such as ` + "`" + `HTTP` + "`" + `,` + "`" + `WEBSOCKET` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "request_config_method",
					Description: `Request frontend method configuration. Like ` + "`" + `GET` + "`" + `,` + "`" + `POST` + "`" + `,` + "`" + `PUT` + "`" + `,` + "`" + `DELETE` + "`" + `,` + "`" + `HEAD` + "`" + `,` + "`" + `ANY` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "request_config_path",
					Description: `Request frontend path configuration. Like ` + "`" + `/user/getinfo` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "request_parameters",
					Description: `Frontend request parameters.`,
				},
				resource.Attribute{
					Name:        "default_value",
					Description: `Parameter default value.`,
				},
				resource.Attribute{
					Name:        "desc",
					Description: `Parameter description.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Parameter name.`,
				},
				resource.Attribute{
					Name:        "position",
					Description: `Parameter location.`,
				},
				resource.Attribute{
					Name:        "required",
					Description: `If this parameter required.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Parameter type.`,
				},
				resource.Attribute{
					Name:        "response_error_codes",
					Description: `Custom error code configuration. Must keep at least one after set.`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `Custom response configuration error code.`,
				},
				resource.Attribute{
					Name:        "converted_code",
					Description: `Custom error code conversion.`,
				},
				resource.Attribute{
					Name:        "desc",
					Description: `Parameter description.`,
				},
				resource.Attribute{
					Name:        "msg",
					Description: `Custom response configuration error message.`,
				},
				resource.Attribute{
					Name:        "need_convert",
					Description: `Whether to enable error code conversion. Default value: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "response_fail_example",
					Description: `Response failure sample of custom response configuration.`,
				},
				resource.Attribute{
					Name:        "response_success_example",
					Description: `Successful response sample of custom response configuration.`,
				},
				resource.Attribute{
					Name:        "response_type",
					Description: `Return type.`,
				},
				resource.Attribute{
					Name:        "service_config_method",
					Description: `API backend service request method, such as ` + "`" + `GET` + "`" + `. If ` + "`" + `service_config_type` + "`" + ` is ` + "`" + `HTTP` + "`" + `, this parameter will be required. The frontend ` + "`" + `request_config_method` + "`" + ` and backend method ` + "`" + `service_config_method` + "`" + ` can be different.`,
				},
				resource.Attribute{
					Name:        "service_config_mock_return_message",
					Description: `Returned information of API backend mocking.`,
				},
				resource.Attribute{
					Name:        "service_config_path",
					Description: `API backend service path, such as /path. If ` + "`" + `service_config_type` + "`" + ` is ` + "`" + `HTTP` + "`" + `, this parameter will be required. The frontend ` + "`" + `request_config_path` + "`" + ` and backend path ` + "`" + `service_config_path` + "`" + ` can be different.`,
				},
				resource.Attribute{
					Name:        "service_config_product",
					Description: `Backend type. This parameter takes effect when VPC is enabled. Currently, only ` + "`" + `clb` + "`" + ` is supported.`,
				},
				resource.Attribute{
					Name:        "service_config_scf_function_name",
					Description: `SCF function name. This parameter takes effect when ` + "`" + `service_config_type` + "`" + ` is ` + "`" + `SCF` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "service_config_scf_function_namespace",
					Description: `SCF function namespace. This parameter takes effect when ` + "`" + `service_config_type` + "`" + ` is ` + "`" + `SCF` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "service_config_scf_function_qualifier",
					Description: `SCF function version. This parameter takes effect when ` + "`" + `service_config_type` + "`" + ` is ` + "`" + `SCF` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "service_config_timeout",
					Description: `API backend service timeout period in seconds.`,
				},
				resource.Attribute{
					Name:        "service_config_type",
					Description: `API backend service type.`,
				},
				resource.Attribute{
					Name:        "service_config_url",
					Description: `API backend service url. This parameter is required when ` + "`" + `service_config_type` + "`" + ` is ` + "`" + `HTTP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "service_config_vpc_id",
					Description: `Unique VPC ID.`,
				},
				resource.Attribute{
					Name:        "service_id",
					Description: `Which service this API belongs. Refer to resource ` + "`" + `tencentcloud_api_gateway_service` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_api_gateway_customer_domains",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query API gateway domain list.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_id",
					Description: `(Required) The service ID.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `Service custom domain name list.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `The certificate ID.`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `Domain name.`,
				},
				resource.Attribute{
					Name:        "is_default_mapping",
					Description: `Whether to use default path mapping. Valid values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + `. ` + "`" + `true` + "`" + ` means to use default path mapping, ` + "`" + `false` + "`" + ` means to use custom path mapping.`,
				},
				resource.Attribute{
					Name:        "is_status_on",
					Description: `Domain name resolution status. Valid values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + `. ` + "`" + `true` + "`" + ` means normal parsing, ` + "`" + `false` + "`" + ` means parsing failed.`,
				},
				resource.Attribute{
					Name:        "net_type",
					Description: `Network type.`,
				},
				resource.Attribute{
					Name:        "path_mappings",
					Description: `Domain name mapping path and environment list.`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `Release environment.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The domain mapping path.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Custom domain name agreement type.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `Service custom domain name list.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `The certificate ID.`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `Domain name.`,
				},
				resource.Attribute{
					Name:        "is_default_mapping",
					Description: `Whether to use default path mapping. Valid values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + `. ` + "`" + `true` + "`" + ` means to use default path mapping, ` + "`" + `false` + "`" + ` means to use custom path mapping.`,
				},
				resource.Attribute{
					Name:        "is_status_on",
					Description: `Domain name resolution status. Valid values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + `. ` + "`" + `true` + "`" + ` means normal parsing, ` + "`" + `false` + "`" + ` means parsing failed.`,
				},
				resource.Attribute{
					Name:        "net_type",
					Description: `Network type.`,
				},
				resource.Attribute{
					Name:        "path_mappings",
					Description: `Domain name mapping path and environment list.`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `Release environment.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The domain mapping path.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Custom domain name agreement type.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_api_gateway_ip_strategies",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query API gateway IP strategy.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_id",
					Description: `(Required) The service ID to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "strategy_name",
					Description: `(Optional) Name of IP policy. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list of strategy.`,
				},
				resource.Attribute{
					Name:        "attach_list",
					Description: `List of bound API details.`,
				},
				resource.Attribute{
					Name:        "api_business_type",
					Description: `The type of oauth API. This field is valid when the ` + "`" + `auth_type` + "`" + ` is ` + "`" + `OAUTH` + "`" + `, and the values are ` + "`" + `NORMAL` + "`" + ` (business API) and ` + "`" + `OAUTH` + "`" + ` (authorization API).`,
				},
				resource.Attribute{
					Name:        "api_desc",
					Description: `API interface description.`,
				},
				resource.Attribute{
					Name:        "api_id",
					Description: `The API ID.`,
				},
				resource.Attribute{
					Name:        "api_name",
					Description: `API name.`,
				},
				resource.Attribute{
					Name:        "api_type",
					Description: `API type. Valid values: ` + "`" + `NORMAL` + "`" + `, ` + "`" + `TSF` + "`" + `. ` + "`" + `NORMAL` + "`" + ` means common API, ` + "`" + `TSF` + "`" + ` means microservice API.`,
				},
				resource.Attribute{
					Name:        "auth_relation_api_id",
					Description: `The unique ID of the associated authorization API, which takes effect when the authType is ` + "`" + `OAUTH` + "`" + ` and ` + "`" + `ApiBusinessType` + "`" + ` is normal. Identifies the unique ID of the oauth2.0 authorization API bound to the business API.`,
				},
				resource.Attribute{
					Name:        "auth_type",
					Description: `API authentication type. Valid values: ` + "`" + `SECRET` + "`" + `, ` + "`" + `NONE` + "`" + `, ` + "`" + `OAUTH` + "`" + `. ` + "`" + `SECRET` + "`" + ` means key pair authentication, ` + "`" + `NONE` + "`" + ` means no authentication.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time in the format of ` + "`" + `YYYY-MM-DDThh:mm:ssZ` + "`" + ` according to ISO 8601 standard. UTC time is used.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `API request method.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `Last modified time in the format of ` + "`" + `YYYY-MM-DDThh:mm:ssZ` + "`" + ` according to ISO 8601 standard. UTC time is used.`,
				},
				resource.Attribute{
					Name:        "oauth_config",
					Description: `OAUTH configuration information. It takes effect when authType is ` + "`" + `OAUTH` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `API path.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `API protocol.`,
				},
				resource.Attribute{
					Name:        "relation_business_api_ids",
					Description: `List of business API associated with authorized API.`,
				},
				resource.Attribute{
					Name:        "service_id",
					Description: `The service ID.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The label information associated with the API.`,
				},
				resource.Attribute{
					Name:        "uniq_vpc_id",
					Description: `VPC unique ID.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC ID.`,
				},
				resource.Attribute{
					Name:        "bind_api_total_count",
					Description: `The number of API bound to the strategy.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.`,
				},
				resource.Attribute{
					Name:        "ip_list",
					Description: `The list of IP.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.`,
				},
				resource.Attribute{
					Name:        "service_id",
					Description: `The service ID.`,
				},
				resource.Attribute{
					Name:        "strategy_id",
					Description: `The strategy ID.`,
				},
				resource.Attribute{
					Name:        "strategy_name",
					Description: `Name of the strategy.`,
				},
				resource.Attribute{
					Name:        "strategy_type",
					Description: `Type of the strategy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list of strategy.`,
				},
				resource.Attribute{
					Name:        "attach_list",
					Description: `List of bound API details.`,
				},
				resource.Attribute{
					Name:        "api_business_type",
					Description: `The type of oauth API. This field is valid when the ` + "`" + `auth_type` + "`" + ` is ` + "`" + `OAUTH` + "`" + `, and the values are ` + "`" + `NORMAL` + "`" + ` (business API) and ` + "`" + `OAUTH` + "`" + ` (authorization API).`,
				},
				resource.Attribute{
					Name:        "api_desc",
					Description: `API interface description.`,
				},
				resource.Attribute{
					Name:        "api_id",
					Description: `The API ID.`,
				},
				resource.Attribute{
					Name:        "api_name",
					Description: `API name.`,
				},
				resource.Attribute{
					Name:        "api_type",
					Description: `API type. Valid values: ` + "`" + `NORMAL` + "`" + `, ` + "`" + `TSF` + "`" + `. ` + "`" + `NORMAL` + "`" + ` means common API, ` + "`" + `TSF` + "`" + ` means microservice API.`,
				},
				resource.Attribute{
					Name:        "auth_relation_api_id",
					Description: `The unique ID of the associated authorization API, which takes effect when the authType is ` + "`" + `OAUTH` + "`" + ` and ` + "`" + `ApiBusinessType` + "`" + ` is normal. Identifies the unique ID of the oauth2.0 authorization API bound to the business API.`,
				},
				resource.Attribute{
					Name:        "auth_type",
					Description: `API authentication type. Valid values: ` + "`" + `SECRET` + "`" + `, ` + "`" + `NONE` + "`" + `, ` + "`" + `OAUTH` + "`" + `. ` + "`" + `SECRET` + "`" + ` means key pair authentication, ` + "`" + `NONE` + "`" + ` means no authentication.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time in the format of ` + "`" + `YYYY-MM-DDThh:mm:ssZ` + "`" + ` according to ISO 8601 standard. UTC time is used.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `API request method.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `Last modified time in the format of ` + "`" + `YYYY-MM-DDThh:mm:ssZ` + "`" + ` according to ISO 8601 standard. UTC time is used.`,
				},
				resource.Attribute{
					Name:        "oauth_config",
					Description: `OAUTH configuration information. It takes effect when authType is ` + "`" + `OAUTH` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `API path.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `API protocol.`,
				},
				resource.Attribute{
					Name:        "relation_business_api_ids",
					Description: `List of business API associated with authorized API.`,
				},
				resource.Attribute{
					Name:        "service_id",
					Description: `The service ID.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The label information associated with the API.`,
				},
				resource.Attribute{
					Name:        "uniq_vpc_id",
					Description: `VPC unique ID.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC ID.`,
				},
				resource.Attribute{
					Name:        "bind_api_total_count",
					Description: `The number of API bound to the strategy.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.`,
				},
				resource.Attribute{
					Name:        "ip_list",
					Description: `The list of IP.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.`,
				},
				resource.Attribute{
					Name:        "service_id",
					Description: `The service ID.`,
				},
				resource.Attribute{
					Name:        "strategy_id",
					Description: `The strategy ID.`,
				},
				resource.Attribute{
					Name:        "strategy_name",
					Description: `Name of the strategy.`,
				},
				resource.Attribute{
					Name:        "strategy_type",
					Description: `Type of the strategy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_api_gateway_services",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query API gateway services.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "service_id",
					Description: `(Optional) Service ID for query.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `(Optional) Service name for query. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list of services.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.`,
				},
				resource.Attribute{
					Name:        "exclusive_set_name",
					Description: `Self-deployed cluster name, which is used to specify the self-deployed cluster where the service is to be created.`,
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
					Description: `Private network access sub-domain name.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `IP version number.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.`,
				},
				resource.Attribute{
					Name:        "net_type",
					Description: `Network type list, which is used to specify the supported network types. Valid values: ` + "`" + `INNER` + "`" + `, ` + "`" + `OUTER` + "`" + `. ` + "`" + `INNER` + "`" + ` indicates access over private network, and ` + "`" + `OUTER` + "`" + ` indicates access over public network.`,
				},
				resource.Attribute{
					Name:        "outer_sub_domain",
					Description: `Public network access subdomain name.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Service frontend request type. Valid values: ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + `, ` + "`" + `http&https` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "service_desc",
					Description: `Custom service description.`,
				},
				resource.Attribute{
					Name:        "service_id",
					Description: `Custom service ID.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `Custom service name.`,
				},
				resource.Attribute{
					Name:        "usage_plan_list",
					Description: `A list of attach usage plans. Each element contains the following attributes:`,
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
					Description: `Name of the usage plan.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list of services.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.`,
				},
				resource.Attribute{
					Name:        "exclusive_set_name",
					Description: `Self-deployed cluster name, which is used to specify the self-deployed cluster where the service is to be created.`,
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
					Description: `Private network access sub-domain name.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `IP version number.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.`,
				},
				resource.Attribute{
					Name:        "net_type",
					Description: `Network type list, which is used to specify the supported network types. Valid values: ` + "`" + `INNER` + "`" + `, ` + "`" + `OUTER` + "`" + `. ` + "`" + `INNER` + "`" + ` indicates access over private network, and ` + "`" + `OUTER` + "`" + ` indicates access over public network.`,
				},
				resource.Attribute{
					Name:        "outer_sub_domain",
					Description: `Public network access subdomain name.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Service frontend request type. Valid values: ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + `, ` + "`" + `http&https` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "service_desc",
					Description: `Custom service description.`,
				},
				resource.Attribute{
					Name:        "service_id",
					Description: `Custom service ID.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `Custom service name.`,
				},
				resource.Attribute{
					Name:        "usage_plan_list",
					Description: `A list of attach usage plans. Each element contains the following attributes:`,
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
					Description: `Name of the usage plan.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_api_gateway_throttling_apis",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query API gateway throttling APIs.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "environment_names",
					Description: `(Optional) Environment list.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "service_id",
					Description: `(Optional) Unique service ID of API. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list of policies bound to API.`,
				},
				resource.Attribute{
					Name:        "api_environment_strategies",
					Description: `List of throttling policies bound to API.`,
				},
				resource.Attribute{
					Name:        "api_id",
					Description: `Unique API ID.`,
				},
				resource.Attribute{
					Name:        "api_name",
					Description: `Custom API name.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `API method.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `API path.`,
				},
				resource.Attribute{
					Name:        "strategy_list",
					Description: `Environment throttling information.`,
				},
				resource.Attribute{
					Name:        "environment_name",
					Description: `Environment name.`,
				},
				resource.Attribute{
					Name:        "quota",
					Description: `Throttling value.`,
				},
				resource.Attribute{
					Name:        "service_id",
					Description: `Unique service ID of API.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list of policies bound to API.`,
				},
				resource.Attribute{
					Name:        "api_environment_strategies",
					Description: `List of throttling policies bound to API.`,
				},
				resource.Attribute{
					Name:        "api_id",
					Description: `Unique API ID.`,
				},
				resource.Attribute{
					Name:        "api_name",
					Description: `Custom API name.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `API method.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `API path.`,
				},
				resource.Attribute{
					Name:        "strategy_list",
					Description: `Environment throttling information.`,
				},
				resource.Attribute{
					Name:        "environment_name",
					Description: `Environment name.`,
				},
				resource.Attribute{
					Name:        "quota",
					Description: `Throttling value.`,
				},
				resource.Attribute{
					Name:        "service_id",
					Description: `Unique service ID of API.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_api_gateway_throttling_services",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query API gateway throttling services.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "service_id",
					Description: `(Optional) Service ID for query. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list of Throttling policy.`,
				},
				resource.Attribute{
					Name:        "environments",
					Description: `A list of Throttling policy.`,
				},
				resource.Attribute{
					Name:        "environment_name",
					Description: `Environment name.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Release status.`,
				},
				resource.Attribute{
					Name:        "strategy",
					Description: `Throttling value.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Access service environment URL.`,
				},
				resource.Attribute{
					Name:        "version_name",
					Description: `Published version number.`,
				},
				resource.Attribute{
					Name:        "service_id",
					Description: `Service ID for query.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list of Throttling policy.`,
				},
				resource.Attribute{
					Name:        "environments",
					Description: `A list of Throttling policy.`,
				},
				resource.Attribute{
					Name:        "environment_name",
					Description: `Environment name.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Release status.`,
				},
				resource.Attribute{
					Name:        "strategy",
					Description: `Throttling value.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Access service environment URL.`,
				},
				resource.Attribute{
					Name:        "version_name",
					Description: `Published version number.`,
				},
				resource.Attribute{
					Name:        "service_id",
					Description: `Service ID for query.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_api_gateway_usage_plan_environments",
			Category:         "Data Sources",
			ShortDescription: `Used to query the environment list bound by the plan.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "usage_plan_id",
					Description: `(Required) ID of the usage plan to be queried.`,
				},
				resource.Attribute{
					Name:        "bind_type",
					Description: `(Optional) Binding type. Valid values: ` + "`" + `API` + "`" + `, ` + "`" + `SERVICE` + "`" + `. Default value: ` + "`" + `SERVICE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list of usage plan binding details.`,
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
					Description: `Creation time in the format of ` + "`" + `YYYY-MM-DDThh:mm:ssZ` + "`" + ` according to ISO 8601 standard. UTC time is used.`,
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
					Description: `Last modified time in the format of ` + "`" + `YYYY-MM-DDThh:mm:ssZ` + "`" + ` according to ISO 8601 standard. UTC time is used.`,
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
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list of usage plan binding details.`,
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
					Description: `Creation time in the format of ` + "`" + `YYYY-MM-DDThh:mm:ssZ` + "`" + ` according to ISO 8601 standard. UTC time is used.`,
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
					Description: `Last modified time in the format of ` + "`" + `YYYY-MM-DDThh:mm:ssZ` + "`" + ` according to ISO 8601 standard. UTC time is used.`,
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
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_api_gateway_usage_plans",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query API gateway usage plans.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "usage_plan_id",
					Description: `(Optional) ID of the usage plan.`,
				},
				resource.Attribute{
					Name:        "usage_plan_name",
					Description: `(Optional) Name of the usage plan. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list of usage plans.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time in the format of ` + "`" + `YYYY-MM-DDThh:mm:ssZ` + "`" + ` according to ISO 8601 standard. UTC time is used.`,
				},
				resource.Attribute{
					Name:        "max_request_num_pre_sec",
					Description: `Limit of requests per second. Valid values formats: ` + "`" + `-1` + "`" + `, ` + "`" + `[1,2000]` + "`" + `. The default value is -1, which indicates no limit.`,
				},
				resource.Attribute{
					Name:        "max_request_num",
					Description: `Total number of requests allowed. Valid value formats: ` + "`" + `-1` + "`" + `, ` + "`" + `[1,99999999]` + "`" + `. The default value is -1, which indicates no limit.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `Last modified time in the format of ` + "`" + `YYYY-MM-DDThh:mm:ssZ` + "`" + ` according to ISO 8601 standard. UTC time is used.`,
				},
				resource.Attribute{
					Name:        "usage_plan_desc",
					Description: `Custom usage plan description.`,
				},
				resource.Attribute{
					Name:        "usage_plan_id",
					Description: `ID of the usage plan.`,
				},
				resource.Attribute{
					Name:        "usage_plan_name",
					Description: `Name of the usage plan.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list of usage plans.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time in the format of ` + "`" + `YYYY-MM-DDThh:mm:ssZ` + "`" + ` according to ISO 8601 standard. UTC time is used.`,
				},
				resource.Attribute{
					Name:        "max_request_num_pre_sec",
					Description: `Limit of requests per second. Valid values formats: ` + "`" + `-1` + "`" + `, ` + "`" + `[1,2000]` + "`" + `. The default value is -1, which indicates no limit.`,
				},
				resource.Attribute{
					Name:        "max_request_num",
					Description: `Total number of requests allowed. Valid value formats: ` + "`" + `-1` + "`" + `, ` + "`" + `[1,99999999]` + "`" + `. The default value is -1, which indicates no limit.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `Last modified time in the format of ` + "`" + `YYYY-MM-DDThh:mm:ssZ` + "`" + ` according to ISO 8601 standard. UTC time is used.`,
				},
				resource.Attribute{
					Name:        "usage_plan_desc",
					Description: `Custom usage plan description.`,
				},
				resource.Attribute{
					Name:        "usage_plan_id",
					Description: `ID of the usage plan.`,
				},
				resource.Attribute{
					Name:        "usage_plan_name",
					Description: `Name of the usage plan.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_as_scaling_configs",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query scaling configuration information.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "configuration_id",
					Description: `(Optional) Launch configuration ID.`,
				},
				resource.Attribute{
					Name:        "configuration_name",
					Description: `(Optional) Launch configuration name.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "configuration_list",
					Description: `A list of configuration. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "configuration_id",
					Description: `Launch configuration ID.`,
				},
				resource.Attribute{
					Name:        "configuration_name",
					Description: `Launch configuration name.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time when the launch configuration was created.`,
				},
				resource.Attribute{
					Name:        "data_disk",
					Description: `Configurations of data disk.`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `Volume of disk in GB. Default is ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `Type of disk.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `Data disk snapshot ID.`,
				},
				resource.Attribute{
					Name:        "disk_type_policy",
					Description: `Policy of cloud disk type.`,
				},
				resource.Attribute{
					Name:        "enhanced_monitor_service",
					Description: `Whether to activate cloud monitor service.`,
				},
				resource.Attribute{
					Name:        "enhanced_security_service",
					Description: `Whether to activate cloud security service.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `ID of available image, for example ` + "`" + `img-8toqc6s3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_tags",
					Description: `A tag list associates with an instance.`,
				},
				resource.Attribute{
					Name:        "instance_types",
					Description: `Instance type list of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "internet_charge_type",
					Description: `Charge types for network traffic.`,
				},
				resource.Attribute{
					Name:        "internet_max_bandwidth_out",
					Description: `Max bandwidth of Internet access in Mbps.`,
				},
				resource.Attribute{
					Name:        "key_ids",
					Description: `ID list of login keys.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of the project to which the configuration belongs. Default value is 0.`,
				},
				resource.Attribute{
					Name:        "public_ip_assigned",
					Description: `Specify whether to assign an Internet IP address.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `Security groups to which the instance belongs.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of a launch configuration.`,
				},
				resource.Attribute{
					Name:        "system_disk_size",
					Description: `System disk size of the scaling configuration in GB.`,
				},
				resource.Attribute{
					Name:        "system_disk_type",
					Description: `System disk category of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `Base64-encoded User Data text.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "configuration_list",
					Description: `A list of configuration. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "configuration_id",
					Description: `Launch configuration ID.`,
				},
				resource.Attribute{
					Name:        "configuration_name",
					Description: `Launch configuration name.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time when the launch configuration was created.`,
				},
				resource.Attribute{
					Name:        "data_disk",
					Description: `Configurations of data disk.`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `Volume of disk in GB. Default is ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `Type of disk.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `Data disk snapshot ID.`,
				},
				resource.Attribute{
					Name:        "disk_type_policy",
					Description: `Policy of cloud disk type.`,
				},
				resource.Attribute{
					Name:        "enhanced_monitor_service",
					Description: `Whether to activate cloud monitor service.`,
				},
				resource.Attribute{
					Name:        "enhanced_security_service",
					Description: `Whether to activate cloud security service.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `ID of available image, for example ` + "`" + `img-8toqc6s3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_tags",
					Description: `A tag list associates with an instance.`,
				},
				resource.Attribute{
					Name:        "instance_types",
					Description: `Instance type list of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "internet_charge_type",
					Description: `Charge types for network traffic.`,
				},
				resource.Attribute{
					Name:        "internet_max_bandwidth_out",
					Description: `Max bandwidth of Internet access in Mbps.`,
				},
				resource.Attribute{
					Name:        "key_ids",
					Description: `ID list of login keys.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of the project to which the configuration belongs. Default value is 0.`,
				},
				resource.Attribute{
					Name:        "public_ip_assigned",
					Description: `Specify whether to assign an Internet IP address.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `Security groups to which the instance belongs.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of a launch configuration.`,
				},
				resource.Attribute{
					Name:        "system_disk_size",
					Description: `System disk size of the scaling configuration in GB.`,
				},
				resource.Attribute{
					Name:        "system_disk_type",
					Description: `System disk category of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `Base64-encoded User Data text.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_as_scaling_groups",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query the detail information of an existing autoscaling group.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "configuration_id",
					Description: `(Optional) Filter results by launch configuration ID.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `(Optional) A specified scaling group ID used to query.`,
				},
				resource.Attribute{
					Name:        "scaling_group_name",
					Description: `(Optional) A scaling group name used to query.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags used to query. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "scaling_group_list",
					Description: `A list of scaling group. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "configuration_id",
					Description: `Launch configuration ID.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time when the AS group was created.`,
				},
				resource.Attribute{
					Name:        "default_cooldown",
					Description: `Default cooldown time of scaling group.`,
				},
				resource.Attribute{
					Name:        "desired_capacity",
					Description: `The desired number of CVM instances.`,
				},
				resource.Attribute{
					Name:        "forward_balancer_ids",
					Description: `A list of application clb ids.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `Listener ID for application load balancers.`,
				},
				resource.Attribute{
					Name:        "load_balancer_id",
					Description: `ID of available load balancers.`,
				},
				resource.Attribute{
					Name:        "location_id",
					Description: `ID of forwarding rules.`,
				},
				resource.Attribute{
					Name:        "target_attribute",
					Description: `Attribute list of target rules.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port number.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Weight.`,
				},
				resource.Attribute{
					Name:        "instance_count",
					Description: `Number of instance.`,
				},
				resource.Attribute{
					Name:        "load_balancer_ids",
					Description: `A list of traditional clb ids which the CVM instances attached to.`,
				},
				resource.Attribute{
					Name:        "max_size",
					Description: `The maximum number of CVM instances.`,
				},
				resource.Attribute{
					Name:        "min_size",
					Description: `The minimum number of CVM instances.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of the project to which the scaling group belongs. Default value is 0.`,
				},
				resource.Attribute{
					Name:        "retry_policy",
					Description: `A retry policy can be used when a creation fails.`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `Auto scaling group ID.`,
				},
				resource.Attribute{
					Name:        "scaling_group_name",
					Description: `Auto scaling group name.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of a scaling group.`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `A list of subnet IDs.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the scaling group.`,
				},
				resource.Attribute{
					Name:        "termination_policies",
					Description: `A policy used to select a CVM instance to be terminated from the scaling group.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the vpc with which the instance is associated.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of available zones.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "scaling_group_list",
					Description: `A list of scaling group. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "configuration_id",
					Description: `Launch configuration ID.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time when the AS group was created.`,
				},
				resource.Attribute{
					Name:        "default_cooldown",
					Description: `Default cooldown time of scaling group.`,
				},
				resource.Attribute{
					Name:        "desired_capacity",
					Description: `The desired number of CVM instances.`,
				},
				resource.Attribute{
					Name:        "forward_balancer_ids",
					Description: `A list of application clb ids.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `Listener ID for application load balancers.`,
				},
				resource.Attribute{
					Name:        "load_balancer_id",
					Description: `ID of available load balancers.`,
				},
				resource.Attribute{
					Name:        "location_id",
					Description: `ID of forwarding rules.`,
				},
				resource.Attribute{
					Name:        "target_attribute",
					Description: `Attribute list of target rules.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port number.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Weight.`,
				},
				resource.Attribute{
					Name:        "instance_count",
					Description: `Number of instance.`,
				},
				resource.Attribute{
					Name:        "load_balancer_ids",
					Description: `A list of traditional clb ids which the CVM instances attached to.`,
				},
				resource.Attribute{
					Name:        "max_size",
					Description: `The maximum number of CVM instances.`,
				},
				resource.Attribute{
					Name:        "min_size",
					Description: `The minimum number of CVM instances.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of the project to which the scaling group belongs. Default value is 0.`,
				},
				resource.Attribute{
					Name:        "retry_policy",
					Description: `A retry policy can be used when a creation fails.`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `Auto scaling group ID.`,
				},
				resource.Attribute{
					Name:        "scaling_group_name",
					Description: `Auto scaling group name.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of a scaling group.`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `A list of subnet IDs.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the scaling group.`,
				},
				resource.Attribute{
					Name:        "termination_policies",
					Description: `A policy used to select a CVM instance to be terminated from the scaling group.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the vpc with which the instance is associated.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of available zones.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_as_scaling_policies",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of scaling policy.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_name",
					Description: `(Optional) Scaling policy name.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `(Optional) Scaling group ID.`,
				},
				resource.Attribute{
					Name:        "scaling_policy_id",
					Description: `(Optional) Scaling policy ID. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "scaling_policy_list",
					Description: `A list of scaling policy. Each element contains the following attributes:`,
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
					Name:        "comparison_operator",
					Description: `Comparison operator.`,
				},
				resource.Attribute{
					Name:        "continuous_time",
					Description: `Retry times.`,
				},
				resource.Attribute{
					Name:        "cooldown",
					Description: `Cool down time of the scaling rule.`,
				},
				resource.Attribute{
					Name:        "metric_name",
					Description: `Name of an indicator.`,
				},
				resource.Attribute{
					Name:        "notification_user_group_ids",
					Description: `Users need to be notified when an alarm is triggered.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `Time period in second.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `Scaling policy name.`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `Scaling policy ID.`,
				},
				resource.Attribute{
					Name:        "statistic",
					Description: `Statistic types.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `Alarm threshold.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "scaling_policy_list",
					Description: `A list of scaling policy. Each element contains the following attributes:`,
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
					Name:        "comparison_operator",
					Description: `Comparison operator.`,
				},
				resource.Attribute{
					Name:        "continuous_time",
					Description: `Retry times.`,
				},
				resource.Attribute{
					Name:        "cooldown",
					Description: `Cool down time of the scaling rule.`,
				},
				resource.Attribute{
					Name:        "metric_name",
					Description: `Name of an indicator.`,
				},
				resource.Attribute{
					Name:        "notification_user_group_ids",
					Description: `Users need to be notified when an alarm is triggered.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `Time period in second.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `Scaling policy name.`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `Scaling policy ID.`,
				},
				resource.Attribute{
					Name:        "statistic",
					Description: `Statistic types.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `Alarm threshold.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_audit_cos_regions",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query the cos region list supported by the audit.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "audit_cos_region_list",
					Description: `List of available regions supported by audit cos.`,
				},
				resource.Attribute{
					Name:        "cos_region_name",
					Description: `Cos region chinese name.`,
				},
				resource.Attribute{
					Name:        "cos_region",
					Description: `Cos region.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "audit_cos_region_list",
					Description: `List of available regions supported by audit cos.`,
				},
				resource.Attribute{
					Name:        "cos_region_name",
					Description: `Cos region chinese name.`,
				},
				resource.Attribute{
					Name:        "cos_region",
					Description: `Cos region.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_audit_key_alias",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query the key alias list specified with region supported by the audit.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Region.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "audit_key_alias_list",
					Description: `List of available key alias supported by audit.`,
				},
				resource.Attribute{
					Name:        "key_alias",
					Description: `Key alias.`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `Key ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "audit_key_alias_list",
					Description: `List of available key alias supported by audit.`,
				},
				resource.Attribute{
					Name:        "key_alias",
					Description: `Key alias.`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `Key ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_audits",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of audits.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the audits.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "audit_list",
					Description: `Information list of the dedicated audits.`,
				},
				resource.Attribute{
					Name:        "audit_switch",
					Description: `Indicate whether audit start logging or not.`,
				},
				resource.Attribute{
					Name:        "cos_bucket",
					Description: `Cos bucket name where audit save logs.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the audit.`,
				},
				resource.Attribute{
					Name:        "log_file_prefix",
					Description: `Prefix of the log file of the audit.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the audit.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "audit_list",
					Description: `Information list of the dedicated audits.`,
				},
				resource.Attribute{
					Name:        "audit_switch",
					Description: `Indicate whether audit start logging or not.`,
				},
				resource.Attribute{
					Name:        "cos_bucket",
					Description: `Cos bucket name where audit save logs.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the audit.`,
				},
				resource.Attribute{
					Name:        "log_file_prefix",
					Description: `Prefix of the log file of the audit.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the audit.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_availability_regions",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get the available regions. By default only ` + "`" + `AVAILABLE` + "`" + ` regions will be returned, but ` + "`" + `UNAVAILABLE` + "`" + ` regions can also be fetched when ` + "`" + `include_unavailable` + "`" + ` is specified.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "include_unavailable",
					Description: `(Optional) A bool variable indicates that the query will include ` + "`" + `UNAVAILABLE` + "`" + ` regions.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) When specified, only the region with the exactly name match will be returned. ` + "`" + `default` + "`" + ` value means it consistent with the provider region.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `A list of regions will be exported and its every element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the region, like ` + "`" + `Guangzhou Region` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the region, like ` + "`" + `ap-guangzhou` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the region, indicate availability using ` + "`" + `AVAILABLE` + "`" + ` and ` + "`" + `UNAVAILABLE` + "`" + ` values.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "regions",
					Description: `A list of regions will be exported and its every element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the region, like ` + "`" + `Guangzhou Region` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the region, like ` + "`" + `ap-guangzhou` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the region, indicate availability using ` + "`" + `AVAILABLE` + "`" + ` and ` + "`" + `UNAVAILABLE` + "`" + ` values.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_availability_zones",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get the available zones in current region. By default only ` + "`" + `AVAILABLE` + "`" + ` zones will be returned, but ` + "`" + `UNAVAILABLE` + "`" + ` zones can also be fetched when ` + "`" + `include_unavailable` + "`" + ` is specified.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "include_unavailable",
					Description: `(Optional) A bool variable indicates that the query will include ` + "`" + `UNAVAILABLE` + "`" + ` zones.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) When specified, only the zone with the exactly name match will be returned.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of zones will be exported and its every element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the zone, like ` + "`" + `Guangzhou Zone 3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `An internal id for the zone, like ` + "`" + `200003` + "`" + `, usually not so useful.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the zone, like ` + "`" + `ap-guangzhou-3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the zone, indicate availability using ` + "`" + `AVAILABLE` + "`" + ` and ` + "`" + `UNAVAILABLE` + "`" + ` values.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "zones",
					Description: `A list of zones will be exported and its every element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the zone, like ` + "`" + `Guangzhou Zone 3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `An internal id for the zone, like ` + "`" + `200003` + "`" + `, usually not so useful.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the zone, like ` + "`" + `ap-guangzhou-3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the zone, indicate availability using ` + "`" + `AVAILABLE` + "`" + ` and ` + "`" + `UNAVAILABLE` + "`" + ` values.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cam_group_memberships",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of CAM group memberships`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group_id",
					Description: `(Optional) ID of CAM group to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "membership_list",
					Description: `A list of CAM group membership. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `ID of CAM group.`,
				},
				resource.Attribute{
					Name:        "user_ids",
					Description: `ID set of the CAM group members.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "membership_list",
					Description: `A list of CAM group membership. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `ID of CAM group.`,
				},
				resource.Attribute{
					Name:        "user_ids",
					Description: `ID set of the CAM group members.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cam_group_policy_attachments",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of CAM group policy attachments`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required) ID of the attached CAM group to be queried.`,
				},
				resource.Attribute{
					Name:        "create_mode",
					Description: `(Optional) Mode of creation of the CAM user policy attachment. 1 means the cam policy attachment is created by production, and the others indicate syntax strategy ways.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Optional) ID of CAM policy to be queried.`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `(Optional) Type of the policy strategy. 'User' means customer strategy and 'QCS' means preset strategy.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "group_policy_attachment_list",
					Description: `A list of CAM group policy attachments. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_mode",
					Description: `Mode of Creation of the CAM group policy attachment. 1 means the cam policy attachment is created by production, and the others indicate syntax strategy ways.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CAM group policy attachment.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `ID of CAM group.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Name of CAM group.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `Name of the policy.`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `Type of the policy strategy. 'User' means customer strategy and 'QCS' means preset strategy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "group_policy_attachment_list",
					Description: `A list of CAM group policy attachments. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_mode",
					Description: `Mode of Creation of the CAM group policy attachment. 1 means the cam policy attachment is created by production, and the others indicate syntax strategy ways.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CAM group policy attachment.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `ID of CAM group.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Name of CAM group.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `Name of the policy.`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `Type of the policy strategy. 'User' means customer strategy and 'QCS' means preset strategy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cam_groups",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of CAM groups`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group_id",
					Description: `(Optional) ID of CAM group to be queried.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the CAM group to be queried.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `(Optional) Description of the cam group to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "group_list",
					Description: `A list of CAM groups. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CAM group.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `ID of the CAM group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of CAM group.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `Description of CAM group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "group_list",
					Description: `A list of CAM groups. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CAM group.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `ID of the CAM group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of CAM group.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `Description of CAM group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cam_policies",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of CAM policies`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "create_mode",
					Description: `(Optional) Mode of creation of policy strategy. Valid values: ` + "`" + `1` + "`" + `, ` + "`" + `2` + "`" + `. ` + "`" + `1` + "`" + ` means policy was created with console, and ` + "`" + `2` + "`" + ` means it was created by strategies.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the CAM policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the CAM policy to be queried.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Optional) ID of CAM policy to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type of the policy strategy. Valid values: ` + "`" + `1` + "`" + `, ` + "`" + `2` + "`" + `. ` + "`" + `1` + "`" + ` means customer strategy and ` + "`" + `2` + "`" + ` means preset strategy. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "policy_list",
					Description: `A list of CAM policies. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "attachments",
					Description: `Number of attached users.`,
				},
				resource.Attribute{
					Name:        "create_mode",
					Description: `Mode of creation of policy strategy. ` + "`" + `1` + "`" + ` means policy was created with console, and ` + "`" + `2` + "`" + ` means it was created by strategies.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CAM policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of CAM policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of CAM policy.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `ID of the policy strategy.`,
				},
				resource.Attribute{
					Name:        "service_type",
					Description: `Name of attached products.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the policy strategy. ` + "`" + `1` + "`" + ` means customer strategy and ` + "`" + `2` + "`" + ` means preset strategy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_list",
					Description: `A list of CAM policies. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "attachments",
					Description: `Number of attached users.`,
				},
				resource.Attribute{
					Name:        "create_mode",
					Description: `Mode of creation of policy strategy. ` + "`" + `1` + "`" + ` means policy was created with console, and ` + "`" + `2` + "`" + ` means it was created by strategies.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CAM policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of CAM policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of CAM policy.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `ID of the policy strategy.`,
				},
				resource.Attribute{
					Name:        "service_type",
					Description: `Name of attached products.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the policy strategy. ` + "`" + `1` + "`" + ` means customer strategy and ` + "`" + `2` + "`" + ` means preset strategy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cam_role_policy_attachments",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of CAM role policy attachments`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "role_id",
					Description: `(Required) ID of the attached CAM role to be queried.`,
				},
				resource.Attribute{
					Name:        "create_mode",
					Description: `(Optional) Mode of Creation of the CAM user policy attachment. ` + "`" + `1` + "`" + ` means the cam policy attachment is created by production, and the others indicate syntax strategy ways.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Optional) ID of CAM policy to be queried.`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `(Optional) Type of the policy strategy. Valid values are 'User', 'QCS'. 'User' means customer strategy and 'QCS' means preset strategy.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "role_policy_attachment_list",
					Description: `A list of CAM role policy attachments. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_mode",
					Description: `Mode of Creation of the CAM role policy attachment. ` + "`" + `1` + "`" + ` means the cam policy attachment is created by production, and the others indicate syntax strategy ways.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CAM role policy attachment.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Name of CAM role.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `Name of the policy.`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `Type of the policy strategy. 'User' means customer strategy and 'QCS' means preset strategy.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `ID of CAM role.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "role_policy_attachment_list",
					Description: `A list of CAM role policy attachments. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_mode",
					Description: `Mode of Creation of the CAM role policy attachment. ` + "`" + `1` + "`" + ` means the cam policy attachment is created by production, and the others indicate syntax strategy ways.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CAM role policy attachment.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Name of CAM role.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `Name of the policy.`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `Type of the policy strategy. 'User' means customer strategy and 'QCS' means preset strategy.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `ID of CAM role.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cam_roles",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of CAM roles`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the CAM role to be queried.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the CAM policy to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `(Optional) ID of the CAM role to be queried. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "role_list",
					Description: `A list of CAM roles. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "console_login",
					Description: `Indicate whether the CAM role can be login or not.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The create time of the CAM role.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of CAM role.`,
				},
				resource.Attribute{
					Name:        "document",
					Description: `Policy document of CAM role.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of CAM role.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `Id of CAM role.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `The last update time of the CAM role.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "role_list",
					Description: `A list of CAM roles. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "console_login",
					Description: `Indicate whether the CAM role can be login or not.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The create time of the CAM role.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of CAM role.`,
				},
				resource.Attribute{
					Name:        "document",
					Description: `Policy document of CAM role.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of CAM role.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `Id of CAM role.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `The last update time of the CAM role.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cam_saml_providers",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of CAM SAML providers`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the CAM SAML provider.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the CAM SAML provider to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "provider_list",
					Description: `A list of CAM SAML providers. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CAM SAML provider.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of CAM SAML provider.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `The last modify time of the CAM SAML provider.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of CAM SAML provider.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "provider_list",
					Description: `A list of CAM SAML providers. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CAM SAML provider.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of CAM SAML provider.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `The last modify time of the CAM SAML provider.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of CAM SAML provider.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cam_user_policy_attachments",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of CAM user policy attachments`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "user_id",
					Description: `(Required) ID of the attached CAM user to be queried.`,
				},
				resource.Attribute{
					Name:        "create_mode",
					Description: `(Optional) Mode of Creation of the CAM user policy attachment. ` + "`" + `1` + "`" + ` means the CAM policy attachment is created by production, and the others indicate syntax strategy ways.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Optional) ID of CAM policy to be queried.`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `(Optional) Type of the policy strategy. 'User' means customer strategy and 'QCS' means preset strategy.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "user_policy_attachment_list",
					Description: `A list of CAM user policy attachments. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_mode",
					Description: `Mode of Creation of the CAM user policy attachment. ` + "`" + `1` + "`" + ` means the cam policy attachment is created by production, and the others indicate syntax strategy ways.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The create time of the CAM user policy attachment.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Name of CAM user.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `The name of the policy.`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `Type of the policy strategy. 'User' means customer strategy and 'QCS' means preset strategy.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `ID of CAM user.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "user_policy_attachment_list",
					Description: `A list of CAM user policy attachments. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_mode",
					Description: `Mode of Creation of the CAM user policy attachment. ` + "`" + `1` + "`" + ` means the cam policy attachment is created by production, and the others indicate syntax strategy ways.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The create time of the CAM user policy attachment.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Name of CAM user.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `The name of the policy.`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `Type of the policy strategy. 'User' means customer strategy and 'QCS' means preset strategy.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `ID of CAM user.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cam_users",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of CAM users`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "console_login",
					Description: `(Optional) Indicate whether the user can login in.`,
				},
				resource.Attribute{
					Name:        "country_code",
					Description: `(Optional) Country code of the CAM user to be queried.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) Email of the CAM user to be queried.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of CAM user to be queried.`,
				},
				resource.Attribute{
					Name:        "phone_num",
					Description: `(Optional) Phone num of the CAM user to be queried.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `(Optional) Remark of the CAM user to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Uid of the CAM user to be queried.`,
				},
				resource.Attribute{
					Name:        "uin",
					Description: `(Optional) Uin of the CAM user to be queried. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "user_list",
					Description: `A list of CAM users. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "country_code",
					Description: `Country code of the CAM user.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `Email of the CAM user.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of CAM user.`,
				},
				resource.Attribute{
					Name:        "phone_num",
					Description: `Phone num of the CAM user.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `Remark of the CAM user.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `Uid of the CAM user.`,
				},
				resource.Attribute{
					Name:        "uin",
					Description: `Uin of the CAM user.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `ID of CAM user. Its value equals to ` + "`" + `name` + "`" + ` argument.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "user_list",
					Description: `A list of CAM users. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "country_code",
					Description: `Country code of the CAM user.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `Email of the CAM user.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of CAM user.`,
				},
				resource.Attribute{
					Name:        "phone_num",
					Description: `Phone num of the CAM user.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `Remark of the CAM user.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `Uid of the CAM user.`,
				},
				resource.Attribute{
					Name:        "uin",
					Description: `Uin of the CAM user.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `ID of CAM user. Its value equals to ` + "`" + `name` + "`" + ` argument.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cbs_snapshot_policies",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of CBS snapshot policies.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "snapshot_policy_id",
					Description: `(Optional) ID of the snapshot policy to be queried.`,
				},
				resource.Attribute{
					Name:        "snapshot_policy_name",
					Description: `(Optional) Name of the snapshot policy to be queried. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "snapshot_policy_list",
					Description: `A list of snapshot policy. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "attached_storage_ids",
					Description: `Storage IDs that the snapshot policy attached.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the snapshot policy.`,
				},
				resource.Attribute{
					Name:        "repeat_hours",
					Description: `Trigger hours of periodic snapshot.`,
				},
				resource.Attribute{
					Name:        "repeat_weekdays",
					Description: `Trigger days of periodic snapshot.`,
				},
				resource.Attribute{
					Name:        "retention_days",
					Description: `Retention days of the snapshot.`,
				},
				resource.Attribute{
					Name:        "snapshot_policy_id",
					Description: `ID of the snapshot policy.`,
				},
				resource.Attribute{
					Name:        "snapshot_policy_name",
					Description: `Name of the snapshot policy.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the snapshot policy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "snapshot_policy_list",
					Description: `A list of snapshot policy. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "attached_storage_ids",
					Description: `Storage IDs that the snapshot policy attached.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the snapshot policy.`,
				},
				resource.Attribute{
					Name:        "repeat_hours",
					Description: `Trigger hours of periodic snapshot.`,
				},
				resource.Attribute{
					Name:        "repeat_weekdays",
					Description: `Trigger days of periodic snapshot.`,
				},
				resource.Attribute{
					Name:        "retention_days",
					Description: `Retention days of the snapshot.`,
				},
				resource.Attribute{
					Name:        "snapshot_policy_id",
					Description: `ID of the snapshot policy.`,
				},
				resource.Attribute{
					Name:        "snapshot_policy_name",
					Description: `Name of the snapshot policy.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the snapshot policy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cbs_snapshots",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of CBS snapshots.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) The available zone that the CBS instance locates at.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) ID of the project within the snapshot.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) ID of the snapshot to be queried.`,
				},
				resource.Attribute{
					Name:        "snapshot_name",
					Description: `(Optional) Name of the snapshot to be queried.`,
				},
				resource.Attribute{
					Name:        "storage_id",
					Description: `(Optional) ID of the the CBS which this snapshot created from.`,
				},
				resource.Attribute{
					Name:        "storage_usage",
					Description: `(Optional) Types of CBS which this snapshot created from, and available values include SYSTEM_DISK and DATA_DISK. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "snapshot_list",
					Description: `A list of snapshot. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The available zone that the CBS instance locates at.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of snapshot.`,
				},
				resource.Attribute{
					Name:        "encrypt",
					Description: `Indicates whether the snapshot is encrypted.`,
				},
				resource.Attribute{
					Name:        "percent",
					Description: `Snapshot creation progress percentage.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of the project within the snapshot.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `ID of the snapshot.`,
				},
				resource.Attribute{
					Name:        "snapshot_name",
					Description: `Name of the snapshot.`,
				},
				resource.Attribute{
					Name:        "storage_id",
					Description: `ID of the the CBS which this snapshot created from.`,
				},
				resource.Attribute{
					Name:        "storage_size",
					Description: `Volume of storage which this snapshot created from.`,
				},
				resource.Attribute{
					Name:        "storage_usage",
					Description: `Types of CBS which this snapshot created from.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "snapshot_list",
					Description: `A list of snapshot. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The available zone that the CBS instance locates at.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of snapshot.`,
				},
				resource.Attribute{
					Name:        "encrypt",
					Description: `Indicates whether the snapshot is encrypted.`,
				},
				resource.Attribute{
					Name:        "percent",
					Description: `Snapshot creation progress percentage.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of the project within the snapshot.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `ID of the snapshot.`,
				},
				resource.Attribute{
					Name:        "snapshot_name",
					Description: `Name of the snapshot.`,
				},
				resource.Attribute{
					Name:        "storage_id",
					Description: `ID of the the CBS which this snapshot created from.`,
				},
				resource.Attribute{
					Name:        "storage_size",
					Description: `Volume of storage which this snapshot created from.`,
				},
				resource.Attribute{
					Name:        "storage_usage",
					Description: `Types of CBS which this snapshot created from.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cbs_storages",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of CBS storages.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) The available zone that the CBS instance locates at.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) ID of the project with which the CBS is associated.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "storage_id",
					Description: `(Optional) ID of the CBS to be queried.`,
				},
				resource.Attribute{
					Name:        "storage_name",
					Description: `(Optional) Name of the CBS to be queried.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `(Optional) Filter by cloud disk media type (` + "`" + `CLOUD_BASIC` + "`" + `: HDD cloud disk | ` + "`" + `CLOUD_PREMIUM` + "`" + `: Premium Cloud Storage | ` + "`" + `CLOUD_SSD` + "`" + `: SSD cloud disk).`,
				},
				resource.Attribute{
					Name:        "storage_usage",
					Description: `(Optional) Filter by cloud disk type (` + "`" + `SYSTEM_DISK` + "`" + `: system disk | ` + "`" + `DATA_DISK` + "`" + `: data disk). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "storage_list",
					Description: `A list of storage. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "attached",
					Description: `Indicates whether the CBS is mounted the CVM.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The zone of CBS.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `Pay type of the CBS instance.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of CBS.`,
				},
				resource.Attribute{
					Name:        "encrypt",
					Description: `Indicates whether CBS is encrypted.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the CVM instance that be mounted by this CBS.`,
				},
				resource.Attribute{
					Name:        "prepaid_renew_flag",
					Description: `The way that CBS instance will be renew automatically or not when it reach the end of the prepaid tenancy.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of the project.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of CBS.`,
				},
				resource.Attribute{
					Name:        "storage_id",
					Description: `ID of CBS.`,
				},
				resource.Attribute{
					Name:        "storage_name",
					Description: `Name of CBS.`,
				},
				resource.Attribute{
					Name:        "storage_size",
					Description: `Volume of CBS.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `Types of storage medium.`,
				},
				resource.Attribute{
					Name:        "storage_usage",
					Description: `Types of CBS.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The available tags within this CBS.`,
				},
				resource.Attribute{
					Name:        "throughput_performance",
					Description: `Add extra performance to the data disk. Only works when disk type is ` + "`" + `CLOUD_TSSD` + "`" + ` or ` + "`" + `CLOUD_HSSD` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "storage_list",
					Description: `A list of storage. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "attached",
					Description: `Indicates whether the CBS is mounted the CVM.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The zone of CBS.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `Pay type of the CBS instance.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of CBS.`,
				},
				resource.Attribute{
					Name:        "encrypt",
					Description: `Indicates whether CBS is encrypted.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the CVM instance that be mounted by this CBS.`,
				},
				resource.Attribute{
					Name:        "prepaid_renew_flag",
					Description: `The way that CBS instance will be renew automatically or not when it reach the end of the prepaid tenancy.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of the project.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of CBS.`,
				},
				resource.Attribute{
					Name:        "storage_id",
					Description: `ID of CBS.`,
				},
				resource.Attribute{
					Name:        "storage_name",
					Description: `Name of CBS.`,
				},
				resource.Attribute{
					Name:        "storage_size",
					Description: `Volume of CBS.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `Types of storage medium.`,
				},
				resource.Attribute{
					Name:        "storage_usage",
					Description: `Types of CBS.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The available tags within this CBS.`,
				},
				resource.Attribute{
					Name:        "throughput_performance",
					Description: `Add extra performance to the data disk. Only works when disk type is ` + "`" + `CLOUD_TSSD` + "`" + ` or ` + "`" + `CLOUD_HSSD` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_ccn_bandwidth_limits",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of CCN bandwidth limits.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ccn_id",
					Description: `(Required) ID of the CCN to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "limits",
					Description: `The bandwidth limits of regions:`,
				},
				resource.Attribute{
					Name:        "bandwidth_limit",
					Description: `Limitation of bandwidth.`,
				},
				resource.Attribute{
					Name:        "dst_region",
					Description: `Destination area restriction.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Limitation of region.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "limits",
					Description: `The bandwidth limits of regions:`,
				},
				resource.Attribute{
					Name:        "bandwidth_limit",
					Description: `Limitation of bandwidth.`,
				},
				resource.Attribute{
					Name:        "dst_region",
					Description: `Destination area restriction.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Limitation of region.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_ccn_instances",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of CCN instances.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ccn_id",
					Description: `(Optional) ID of the CCN to be queried.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the CCN to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "instance_list",
					Description: `Information list of CCN.`,
				},
				resource.Attribute{
					Name:        "attachment_list",
					Description: `Information list of instance is attached.`,
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
					Name:        "instance_id",
					Description: `ID of instance is attached.`,
				},
				resource.Attribute{
					Name:        "instance_region",
					Description: `The region that the instance locates at.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `Type of attached instance network, and available values include VPC, DIRECTCONNECT, BMVPC and VPNGW.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `States of instance is attached, and available values include PENDING, ACTIVE, EXPIRED, REJECTED, DELETED, FAILED(asynchronous forced disassociation after 2 hours), ATTACHING, DETACHING and DETACHFAILED(asynchronous forced disassociation after 2 hours).`,
				},
				resource.Attribute{
					Name:        "bandwidth_limit_type",
					Description: `The speed limit type.`,
				},
				resource.Attribute{
					Name:        "ccn_id",
					Description: `ID of the CCN.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `Billing mode.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the CCN.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the CCN.`,
				},
				resource.Attribute{
					Name:        "qos",
					Description: `Service quality of CCN, and the available value include 'PT', 'AU', 'AG'. The default is 'AU'.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `States of instance. The available value include 'ISOLATED'(arrears) and 'AVAILABLE'.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_list",
					Description: `Information list of CCN.`,
				},
				resource.Attribute{
					Name:        "attachment_list",
					Description: `Information list of instance is attached.`,
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
					Name:        "instance_id",
					Description: `ID of instance is attached.`,
				},
				resource.Attribute{
					Name:        "instance_region",
					Description: `The region that the instance locates at.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `Type of attached instance network, and available values include VPC, DIRECTCONNECT, BMVPC and VPNGW.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `States of instance is attached, and available values include PENDING, ACTIVE, EXPIRED, REJECTED, DELETED, FAILED(asynchronous forced disassociation after 2 hours), ATTACHING, DETACHING and DETACHFAILED(asynchronous forced disassociation after 2 hours).`,
				},
				resource.Attribute{
					Name:        "bandwidth_limit_type",
					Description: `The speed limit type.`,
				},
				resource.Attribute{
					Name:        "ccn_id",
					Description: `ID of the CCN.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `Billing mode.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the CCN.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the CCN.`,
				},
				resource.Attribute{
					Name:        "qos",
					Description: `Service quality of CCN, and the available value include 'PT', 'AU', 'AG'. The default is 'AU'.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `States of instance. The available value include 'ISOLATED'(arrears) and 'AVAILABLE'.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cdh_instances",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query CDH instances.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) The available zone that the CDH instance locates at.`,
				},
				resource.Attribute{
					Name:        "host_id",
					Description: `(Optional) ID of the CDH instances to be queried.`,
				},
				resource.Attribute{
					Name:        "host_name",
					Description: `(Optional) Name of the CDH instances to be queried.`,
				},
				resource.Attribute{
					Name:        "host_state",
					Description: `(Optional) State of the CDH instances to be queried. Valid values: ` + "`" + `PENDING` + "`" + `, ` + "`" + `LAUNCH_FAILURE` + "`" + `, ` + "`" + `RUNNING` + "`" + `, ` + "`" + `EXPIRED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) The project CDH belongs to.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cdh_instance_list",
					Description: `An information list of cdh instance. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The available zone that the CDH instance locates at.`,
				},
				resource.Attribute{
					Name:        "cage_id",
					Description: `Cage ID of the CDH instance. This parameter is only valid for CDH instances in the cages of finance availability zones.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `The charge type of the CDH instance.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the CDH instance.`,
				},
				resource.Attribute{
					Name:        "cvm_instance_ids",
					Description: `Id of CVM instances that have been created on the CDH instance.`,
				},
				resource.Attribute{
					Name:        "expired_time",
					Description: `Expired time of the CDH instance.`,
				},
				resource.Attribute{
					Name:        "host_id",
					Description: `ID of the CDH instance.`,
				},
				resource.Attribute{
					Name:        "host_name",
					Description: `Name of the CDH instance.`,
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
					Description: `State of the CDH instance.`,
				},
				resource.Attribute{
					Name:        "host_type",
					Description: `Type of the CDH instance.`,
				},
				resource.Attribute{
					Name:        "prepaid_renew_flag",
					Description: `Auto renewal flag.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The project CDH belongs to.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cdh_instance_list",
					Description: `An information list of cdh instance. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The available zone that the CDH instance locates at.`,
				},
				resource.Attribute{
					Name:        "cage_id",
					Description: `Cage ID of the CDH instance. This parameter is only valid for CDH instances in the cages of finance availability zones.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `The charge type of the CDH instance.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the CDH instance.`,
				},
				resource.Attribute{
					Name:        "cvm_instance_ids",
					Description: `Id of CVM instances that have been created on the CDH instance.`,
				},
				resource.Attribute{
					Name:        "expired_time",
					Description: `Expired time of the CDH instance.`,
				},
				resource.Attribute{
					Name:        "host_id",
					Description: `ID of the CDH instance.`,
				},
				resource.Attribute{
					Name:        "host_name",
					Description: `Name of the CDH instance.`,
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
					Description: `State of the CDH instance.`,
				},
				resource.Attribute{
					Name:        "host_type",
					Description: `Type of the CDH instance.`,
				},
				resource.Attribute{
					Name:        "prepaid_renew_flag",
					Description: `Auto renewal flag.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The project CDH belongs to.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cdn_domains",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query the detail information of CDN domain.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional) Acceleration domain name.`,
				},
				resource.Attribute{
					Name:        "full_url_cache",
					Description: `(Optional) Whether to enable full-path cache.`,
				},
				resource.Attribute{
					Name:        "https_switch",
					Description: `(Optional) HTTPS configuration. Valid values: ` + "`" + `on` + "`" + `, ` + "`" + `off` + "`" + ` and ` + "`" + `processing` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "origin_pull_protocol",
					Description: `(Optional) Origin-pull protocol configuration. Valid values: ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + ` and ` + "`" + `follow` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "service_type",
					Description: `(Optional) Service type of acceleration domain name. The available value include ` + "`" + `web` + "`" + `, ` + "`" + `download` + "`" + ` and ` + "`" + `media` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "domain_list",
					Description: `An information list of cdn domain. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "area",
					Description: `Acceleration region.`,
				},
				resource.Attribute{
					Name:        "cname",
					Description: `CNAME address of domain name.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Domain name creation time.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Acceleration domain name.`,
				},
				resource.Attribute{
					Name:        "full_url_cache",
					Description: `Whether to enable full-path cache.`,
				},
				resource.Attribute{
					Name:        "https_config",
					Description: `HTTPS acceleration configuration. It's a list and consist of at most one item.`,
				},
				resource.Attribute{
					Name:        "http2_switch",
					Description: `HTTP2 configuration switch.`,
				},
				resource.Attribute{
					Name:        "https_switch",
					Description: `HTTPS configuration switch.`,
				},
				resource.Attribute{
					Name:        "ocsp_stapling_switch",
					Description: `OCSP configuration switch.`,
				},
				resource.Attribute{
					Name:        "spdy_switch",
					Description: `Spdy configuration switch.`,
				},
				resource.Attribute{
					Name:        "verify_client",
					Description: `Client certificate authentication feature.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Domain name ID.`,
				},
				resource.Attribute{
					Name:        "origin",
					Description: `Origin server configuration.`,
				},
				resource.Attribute{
					Name:        "backup_origin_list",
					Description: `Backup origin server list.`,
				},
				resource.Attribute{
					Name:        "backup_origin_type",
					Description: `Backup origin server type.`,
				},
				resource.Attribute{
					Name:        "backup_server_name",
					Description: `Host header used when accessing the backup origin server. If left empty, the ServerName of master origin server will be used by default.`,
				},
				resource.Attribute{
					Name:        "cos_private_access",
					Description: `When OriginType is COS, you can specify if access to private buckets is allowed.`,
				},
				resource.Attribute{
					Name:        "origin_list",
					Description: `Master origin server list.`,
				},
				resource.Attribute{
					Name:        "origin_pull_protocol",
					Description: `Origin-pull protocol configuration.`,
				},
				resource.Attribute{
					Name:        "origin_type",
					Description: `Master origin server type.`,
				},
				resource.Attribute{
					Name:        "server_name",
					Description: `Host header used when accessing the master origin server. If left empty, the acceleration domain name will be used by default.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The project CDN belongs to.`,
				},
				resource.Attribute{
					Name:        "range_origin_switch",
					Description: `Sharding back to source configuration switch.`,
				},
				resource.Attribute{
					Name:        "request_header",
					Description: `Request header configuration.`,
				},
				resource.Attribute{
					Name:        "header_rules",
					Description: `Custom request header configuration rules.`,
				},
				resource.Attribute{
					Name:        "header_mode",
					Description: `Http header setting method.`,
				},
				resource.Attribute{
					Name:        "header_name",
					Description: `Http header name.`,
				},
				resource.Attribute{
					Name:        "header_value",
					Description: `Http header value.`,
				},
				resource.Attribute{
					Name:        "rule_paths",
					Description: `Rule paths.`,
				},
				resource.Attribute{
					Name:        "rule_type",
					Description: `Rule type.`,
				},
				resource.Attribute{
					Name:        "switch",
					Description: `Custom request header configuration switch.`,
				},
				resource.Attribute{
					Name:        "rule_cache",
					Description: `Advanced path cache configuration.`,
				},
				resource.Attribute{
					Name:        "follow_origin_switch",
					Description: `Follow the source station configuration switch.`,
				},
				resource.Attribute{
					Name:        "ignore_set_cookie",
					Description: `Ignore the Set-Cookie header of the origin site.`,
				},
				resource.Attribute{
					Name:        "no_cache_switch",
					Description: `Cache configuration switch.`,
				},
				resource.Attribute{
					Name:        "re_validate",
					Description: `Always check back to origin.`,
				},
				resource.Attribute{
					Name:        "rule_paths",
					Description: `Rule paths.`,
				},
				resource.Attribute{
					Name:        "rule_type",
					Description: `Rule type.`,
				},
				resource.Attribute{
					Name:        "switch",
					Description: `Cache configuration switch.`,
				},
				resource.Attribute{
					Name:        "service_type",
					Description: `Service type of acceleration domain name.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Acceleration service status.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of cdn domain.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Last modified time of domain name.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "domain_list",
					Description: `An information list of cdn domain. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "area",
					Description: `Acceleration region.`,
				},
				resource.Attribute{
					Name:        "cname",
					Description: `CNAME address of domain name.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Domain name creation time.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Acceleration domain name.`,
				},
				resource.Attribute{
					Name:        "full_url_cache",
					Description: `Whether to enable full-path cache.`,
				},
				resource.Attribute{
					Name:        "https_config",
					Description: `HTTPS acceleration configuration. It's a list and consist of at most one item.`,
				},
				resource.Attribute{
					Name:        "http2_switch",
					Description: `HTTP2 configuration switch.`,
				},
				resource.Attribute{
					Name:        "https_switch",
					Description: `HTTPS configuration switch.`,
				},
				resource.Attribute{
					Name:        "ocsp_stapling_switch",
					Description: `OCSP configuration switch.`,
				},
				resource.Attribute{
					Name:        "spdy_switch",
					Description: `Spdy configuration switch.`,
				},
				resource.Attribute{
					Name:        "verify_client",
					Description: `Client certificate authentication feature.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Domain name ID.`,
				},
				resource.Attribute{
					Name:        "origin",
					Description: `Origin server configuration.`,
				},
				resource.Attribute{
					Name:        "backup_origin_list",
					Description: `Backup origin server list.`,
				},
				resource.Attribute{
					Name:        "backup_origin_type",
					Description: `Backup origin server type.`,
				},
				resource.Attribute{
					Name:        "backup_server_name",
					Description: `Host header used when accessing the backup origin server. If left empty, the ServerName of master origin server will be used by default.`,
				},
				resource.Attribute{
					Name:        "cos_private_access",
					Description: `When OriginType is COS, you can specify if access to private buckets is allowed.`,
				},
				resource.Attribute{
					Name:        "origin_list",
					Description: `Master origin server list.`,
				},
				resource.Attribute{
					Name:        "origin_pull_protocol",
					Description: `Origin-pull protocol configuration.`,
				},
				resource.Attribute{
					Name:        "origin_type",
					Description: `Master origin server type.`,
				},
				resource.Attribute{
					Name:        "server_name",
					Description: `Host header used when accessing the master origin server. If left empty, the acceleration domain name will be used by default.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The project CDN belongs to.`,
				},
				resource.Attribute{
					Name:        "range_origin_switch",
					Description: `Sharding back to source configuration switch.`,
				},
				resource.Attribute{
					Name:        "request_header",
					Description: `Request header configuration.`,
				},
				resource.Attribute{
					Name:        "header_rules",
					Description: `Custom request header configuration rules.`,
				},
				resource.Attribute{
					Name:        "header_mode",
					Description: `Http header setting method.`,
				},
				resource.Attribute{
					Name:        "header_name",
					Description: `Http header name.`,
				},
				resource.Attribute{
					Name:        "header_value",
					Description: `Http header value.`,
				},
				resource.Attribute{
					Name:        "rule_paths",
					Description: `Rule paths.`,
				},
				resource.Attribute{
					Name:        "rule_type",
					Description: `Rule type.`,
				},
				resource.Attribute{
					Name:        "switch",
					Description: `Custom request header configuration switch.`,
				},
				resource.Attribute{
					Name:        "rule_cache",
					Description: `Advanced path cache configuration.`,
				},
				resource.Attribute{
					Name:        "follow_origin_switch",
					Description: `Follow the source station configuration switch.`,
				},
				resource.Attribute{
					Name:        "ignore_set_cookie",
					Description: `Ignore the Set-Cookie header of the origin site.`,
				},
				resource.Attribute{
					Name:        "no_cache_switch",
					Description: `Cache configuration switch.`,
				},
				resource.Attribute{
					Name:        "re_validate",
					Description: `Always check back to origin.`,
				},
				resource.Attribute{
					Name:        "rule_paths",
					Description: `Rule paths.`,
				},
				resource.Attribute{
					Name:        "rule_type",
					Description: `Rule type.`,
				},
				resource.Attribute{
					Name:        "switch",
					Description: `Cache configuration switch.`,
				},
				resource.Attribute{
					Name:        "service_type",
					Description: `Service type of acceleration domain name.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Acceleration service status.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of cdn domain.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Last modified time of domain name.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cfs_access_groups",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query the detail information of CFS access group.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "access_group_id",
					Description: `(Optional) A specified access group ID used to query.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A access group Name used to query.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "access_group_list",
					Description: `An information list of CFS access group. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "access_group_id",
					Description: `ID of the access group.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the access group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the access group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the access group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "access_group_list",
					Description: `An information list of CFS access group. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "access_group_id",
					Description: `ID of the access group.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the access group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the access group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the access group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cfs_access_rules",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query the detail information of CFS access rule.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "access_group_id",
					Description: `(Required) A specified access group ID used to query.`,
				},
				resource.Attribute{
					Name:        "access_rule_id",
					Description: `(Optional) A specified access rule ID used to query.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "access_rule_list",
					Description: `An information list of CFS access rule. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "access_rule_id",
					Description: `ID of the access rule.`,
				},
				resource.Attribute{
					Name:        "auth_client_ip",
					Description: `Allowed IP of the access rule.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority level of access rule.`,
				},
				resource.Attribute{
					Name:        "rw_permission",
					Description: `Read and write permissions.`,
				},
				resource.Attribute{
					Name:        "user_permission",
					Description: `The permissions of accessing users.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "access_rule_list",
					Description: `An information list of CFS access rule. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "access_rule_id",
					Description: `ID of the access rule.`,
				},
				resource.Attribute{
					Name:        "auth_client_ip",
					Description: `Allowed IP of the access rule.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority level of access rule.`,
				},
				resource.Attribute{
					Name:        "rw_permission",
					Description: `Read and write permissions.`,
				},
				resource.Attribute{
					Name:        "user_permission",
					Description: `The permissions of accessing users.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cfs_file_systems",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query the detail information of cloud file systems(CFS).`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) The available zone that the file system locates at.`,
				},
				resource.Attribute{
					Name:        "file_system_id",
					Description: `(Optional) A specified file system ID used to query.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A file system name used to query.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) ID of a vpc subnet.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) ID of the vpc to be queried. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "file_system_list",
					Description: `An information list of cloud file system. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "access_group_id",
					Description: `ID of the access group.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The available zone that the file system locates at.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the file system.`,
				},
				resource.Attribute{
					Name:        "file_system_id",
					Description: `ID of the file system.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the file system.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol of the file system.`,
				},
				resource.Attribute{
					Name:        "size_limit",
					Description: `Size limit of the file system.`,
				},
				resource.Attribute{
					Name:        "size_used",
					Description: `Size used of the file system.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the file system.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `Storage type of the file system.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "file_system_list",
					Description: `An information list of cloud file system. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "access_group_id",
					Description: `ID of the access group.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The available zone that the file system locates at.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the file system.`,
				},
				resource.Attribute{
					Name:        "file_system_id",
					Description: `ID of the file system.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the file system.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol of the file system.`,
				},
				resource.Attribute{
					Name:        "size_limit",
					Description: `Size limit of the file system.`,
				},
				resource.Attribute{
					Name:        "size_used",
					Description: `Size used of the file system.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the file system.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `Storage type of the file system.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_ckafka_acls",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed acl information of Ckafka`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) Id of the ckafka instance.`,
				},
				resource.Attribute{
					Name:        "resource_name",
					Description: `(Required) ACL resource name, which is related to ` + "`" + `resource_type` + "`" + `. For example, if ` + "`" + `resource_type` + "`" + ` is ` + "`" + `TOPIC` + "`" + `, this field indicates the topic name; if ` + "`" + `resource_type` + "`" + ` is ` + "`" + `GROUP` + "`" + `, this field indicates the group name.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `(Required) ACL resource type. Valid values are ` + "`" + `UNKNOWN` + "`" + `, ` + "`" + `ANY` + "`" + `, ` + "`" + `TOPIC` + "`" + `, ` + "`" + `GROUP` + "`" + `, ` + "`" + `CLUSTER` + "`" + `, ` + "`" + `TRANSACTIONAL_ID` + "`" + `. Currently, only ` + "`" + `TOPIC` + "`" + ` is available, and other fields will be used for future ACLs compatible with open-source Kafka.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Optional) Host substr used for querying.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "acl_list",
					Description: `A list of ckafka acls. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `IP address allowed to access.`,
				},
				resource.Attribute{
					Name:        "operation_type",
					Description: `ACL operation mode.`,
				},
				resource.Attribute{
					Name:        "permission_type",
					Description: `ACL permission type, valid values are ` + "`" + `UNKNOWN` + "`" + `, ` + "`" + `ANY` + "`" + `, ` + "`" + `DENY` + "`" + `, ` + "`" + `ALLOW` + "`" + `, and ` + "`" + `ALLOW` + "`" + ` by default. Currently, CKafka supports ` + "`" + `ALLOW` + "`" + ` (equivalent to allow list), and other fields will be used for future ACLs compatible with open-source Kafka.`,
				},
				resource.Attribute{
					Name:        "principal",
					Description: `User which can access. ` + "`" + ``,
				},
				resource.Attribute{
					Name:        "resource_name",
					Description: `ACL resource name, which is related to ` + "`" + `resource_type` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `ACL resource type.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "acl_list",
					Description: `A list of ckafka acls. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `IP address allowed to access.`,
				},
				resource.Attribute{
					Name:        "operation_type",
					Description: `ACL operation mode.`,
				},
				resource.Attribute{
					Name:        "permission_type",
					Description: `ACL permission type, valid values are ` + "`" + `UNKNOWN` + "`" + `, ` + "`" + `ANY` + "`" + `, ` + "`" + `DENY` + "`" + `, ` + "`" + `ALLOW` + "`" + `, and ` + "`" + `ALLOW` + "`" + ` by default. Currently, CKafka supports ` + "`" + `ALLOW` + "`" + ` (equivalent to allow list), and other fields will be used for future ACLs compatible with open-source Kafka.`,
				},
				resource.Attribute{
					Name:        "principal",
					Description: `User which can access. ` + "`" + ``,
				},
				resource.Attribute{
					Name:        "resource_name",
					Description: `ACL resource name, which is related to ` + "`" + `resource_type` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `ACL resource type.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_ckafka_topics",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of ckafka topic.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) Ckafka instance ID.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to store results.`,
				},
				resource.Attribute{
					Name:        "topic_name",
					Description: `(Optional) Name of the CKafka topic. It must start with a letter, the rest can contain letters, numbers and dashes(-). The length range is from 1 to 64. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "instance_list",
					Description: `A list of instances. Each element contains the following attributes.`,
				},
				resource.Attribute{
					Name:        "clean_up_policy",
					Description: `Clear log policy, log clear mode. ` + "`" + `delete` + "`" + `: logs are deleted according to the storage time, ` + "`" + `compact` + "`" + `: logs are compressed according to the key, ` + "`" + `compact, delete` + "`" + `: logs are compressed according to the key and will be deleted according to the storage time.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CKafka topic.`,
				},
				resource.Attribute{
					Name:        "enable_white_list",
					Description: `Whether to open the IP Whitelist. ` + "`" + `true` + "`" + `: open, ` + "`" + `false` + "`" + `: close.`,
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
					Description: `Data backup cos status. ` + "`" + `1` + "`" + `: do not open data backup, ` + "`" + `0` + "`" + `: open data backup.`,
				},
				resource.Attribute{
					Name:        "ip_white_list_count",
					Description: `IP Whitelist count.`,
				},
				resource.Attribute{
					Name:        "max_message_bytes",
					Description: `Max message bytes.`,
				},
				resource.Attribute{
					Name:        "note",
					Description: `CKafka topic note description.`,
				},
				resource.Attribute{
					Name:        "partition_num",
					Description: `The number of partition.`,
				},
				resource.Attribute{
					Name:        "replica_num",
					Description: `The number of replica.`,
				},
				resource.Attribute{
					Name:        "retention",
					Description: `Message can be selected. Retention time(unit ms).`,
				},
				resource.Attribute{
					Name:        "segment_bytes",
					Description: `Number of bytes rolled by shard.`,
				},
				resource.Attribute{
					Name:        "segment",
					Description: `Segment scrolling time, in ms.`,
				},
				resource.Attribute{
					Name:        "sync_replica_min_num",
					Description: `Min number of sync replicas.`,
				},
				resource.Attribute{
					Name:        "topic_id",
					Description: `ID of the CKafka topic.`,
				},
				resource.Attribute{
					Name:        "topic_name",
					Description: `Name of the CKafka topic.`,
				},
				resource.Attribute{
					Name:        "unclean_leader_election_enable",
					Description: `Whether to allow unsynchronized replicas to be selected as leader, default is ` + "`" + `false` + "`" + `, ` + "`" + `true: ` + "`" + `allowed, ` + "`" + `false` + "`" + `: not allowed.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_list",
					Description: `A list of instances. Each element contains the following attributes.`,
				},
				resource.Attribute{
					Name:        "clean_up_policy",
					Description: `Clear log policy, log clear mode. ` + "`" + `delete` + "`" + `: logs are deleted according to the storage time, ` + "`" + `compact` + "`" + `: logs are compressed according to the key, ` + "`" + `compact, delete` + "`" + `: logs are compressed according to the key and will be deleted according to the storage time.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CKafka topic.`,
				},
				resource.Attribute{
					Name:        "enable_white_list",
					Description: `Whether to open the IP Whitelist. ` + "`" + `true` + "`" + `: open, ` + "`" + `false` + "`" + `: close.`,
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
					Description: `Data backup cos status. ` + "`" + `1` + "`" + `: do not open data backup, ` + "`" + `0` + "`" + `: open data backup.`,
				},
				resource.Attribute{
					Name:        "ip_white_list_count",
					Description: `IP Whitelist count.`,
				},
				resource.Attribute{
					Name:        "max_message_bytes",
					Description: `Max message bytes.`,
				},
				resource.Attribute{
					Name:        "note",
					Description: `CKafka topic note description.`,
				},
				resource.Attribute{
					Name:        "partition_num",
					Description: `The number of partition.`,
				},
				resource.Attribute{
					Name:        "replica_num",
					Description: `The number of replica.`,
				},
				resource.Attribute{
					Name:        "retention",
					Description: `Message can be selected. Retention time(unit ms).`,
				},
				resource.Attribute{
					Name:        "segment_bytes",
					Description: `Number of bytes rolled by shard.`,
				},
				resource.Attribute{
					Name:        "segment",
					Description: `Segment scrolling time, in ms.`,
				},
				resource.Attribute{
					Name:        "sync_replica_min_num",
					Description: `Min number of sync replicas.`,
				},
				resource.Attribute{
					Name:        "topic_id",
					Description: `ID of the CKafka topic.`,
				},
				resource.Attribute{
					Name:        "topic_name",
					Description: `Name of the CKafka topic.`,
				},
				resource.Attribute{
					Name:        "unclean_leader_election_enable",
					Description: `Whether to allow unsynchronized replicas to be selected as leader, default is ` + "`" + `false` + "`" + `, ` + "`" + `true: ` + "`" + `allowed, ` + "`" + `false` + "`" + `: not allowed.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_ckafka_users",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed user information of Ckafka`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) Id of the ckafka instance.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `(Optional) Account name used when query ckafka users' infos. Could be a substr of user name.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "user_list",
					Description: `A list of ckafka users. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `Account name of user.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the account.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `The last update time of the account.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "user_list",
					Description: `A list of ckafka users. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `Account name of user.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the account.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `The last update time of the account.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_clb_attachments",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of CLB attachments`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "clb_id",
					Description: `(Required) ID of the CLB to be queried.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Required) ID of the CLB listener to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `(Optional) ID of the CLB listener rule. If the protocol of listener is ` + "`" + `HTTP` + "`" + `/` + "`" + `HTTPS` + "`" + `, this para is required. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "attachment_list",
					Description: `A list of cloud load balancer attachment configurations. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "clb_id",
					Description: `ID of the CLB.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `ID of the CLB listener.`,
				},
				resource.Attribute{
					Name:        "protocol_type",
					Description: `Type of protocol within the listener, and available values include ` + "`" + `TCP` + "`" + `, ` + "`" + `UDP` + "`" + `, ` + "`" + `HTTP` + "`" + `, ` + "`" + `HTTPS` + "`" + ` and ` + "`" + `TCP_SSL` + "`" + `. NOTES: ` + "`" + `TCP_SSL` + "`" + ` is testing internally, please apply if you need to use.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `ID of the CLB listener rule.`,
				},
				resource.Attribute{
					Name:        "targets",
					Description: `Information of the backends to be attached.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `Id of the backend server.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port of the backend server.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Forwarding weight of the backend service, the range of [0, 100], defaults to ` + "`" + `10` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "attachment_list",
					Description: `A list of cloud load balancer attachment configurations. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "clb_id",
					Description: `ID of the CLB.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `ID of the CLB listener.`,
				},
				resource.Attribute{
					Name:        "protocol_type",
					Description: `Type of protocol within the listener, and available values include ` + "`" + `TCP` + "`" + `, ` + "`" + `UDP` + "`" + `, ` + "`" + `HTTP` + "`" + `, ` + "`" + `HTTPS` + "`" + ` and ` + "`" + `TCP_SSL` + "`" + `. NOTES: ` + "`" + `TCP_SSL` + "`" + ` is testing internally, please apply if you need to use.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `ID of the CLB listener rule.`,
				},
				resource.Attribute{
					Name:        "targets",
					Description: `Information of the backends to be attached.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `Id of the backend server.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port of the backend server.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Forwarding weight of the backend service, the range of [0, 100], defaults to ` + "`" + `10` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_clb_instances",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of CLB`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "clb_id",
					Description: `(Optional) ID of the CLB to be queried.`,
				},
				resource.Attribute{
					Name:        "clb_name",
					Description: `(Optional) Name of the CLB to be queried.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `(Optional) Type of CLB instance, and available values include ` + "`" + `OPEN` + "`" + ` and ` + "`" + `INTERNAL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) Project ID of the CLB.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "clb_list",
					Description: `A list of cloud load balancers. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "address_ip_version",
					Description: `IP version, only applicable to open CLB. Valid values are ` + "`" + `IPV4` + "`" + `, ` + "`" + `IPV6` + "`" + ` and ` + "`" + `IPv6FullChain` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "clb_id",
					Description: `ID of CLB.`,
				},
				resource.Attribute{
					Name:        "clb_name",
					Description: `Name of CLB.`,
				},
				resource.Attribute{
					Name:        "clb_vips",
					Description: `The virtual service address table of the CLB.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CLB.`,
				},
				resource.Attribute{
					Name:        "internet_bandwidth_max_out",
					Description: `Max bandwidth out, only applicable to open CLB. Valid value ranges is [1, 2048]. Unit is MB.`,
				},
				resource.Attribute{
					Name:        "internet_charge_type",
					Description: `Internet charge type, only applicable to open CLB. Valid values are ` + "`" + `TRAFFIC_POSTPAID_BY_HOUR` + "`" + `, ` + "`" + `BANDWIDTH_POSTPAID_BY_HOUR` + "`" + ` and ` + "`" + `BANDWIDTH_PACKAGE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `Types of CLB.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of the project.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `ID set of the security groups.`,
				},
				resource.Attribute{
					Name:        "status_time",
					Description: `Latest state transition time of CLB.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of CLB.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of the subnet.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The available tags within this CLB.`,
				},
				resource.Attribute{
					Name:        "target_region_info_region",
					Description: `Region information of backend service are attached the CLB.`,
				},
				resource.Attribute{
					Name:        "target_region_info_vpc_id",
					Description: `VpcId information of backend service are attached the CLB.`,
				},
				resource.Attribute{
					Name:        "vip_isp",
					Description: `Network operator, only applicable to open CLB. Valid values are ` + "`" + `CMCC` + "`" + `(China Mobile), ` + "`" + `CTCC` + "`" + `(Telecom), ` + "`" + `CUCC` + "`" + `(China Unicom) and ` + "`" + `BGP` + "`" + `. If this ISP is specified, network billing method can only use the bandwidth package billing (BANDWIDTH_PACKAGE).`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "clb_list",
					Description: `A list of cloud load balancers. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "address_ip_version",
					Description: `IP version, only applicable to open CLB. Valid values are ` + "`" + `IPV4` + "`" + `, ` + "`" + `IPV6` + "`" + ` and ` + "`" + `IPv6FullChain` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "clb_id",
					Description: `ID of CLB.`,
				},
				resource.Attribute{
					Name:        "clb_name",
					Description: `Name of CLB.`,
				},
				resource.Attribute{
					Name:        "clb_vips",
					Description: `The virtual service address table of the CLB.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CLB.`,
				},
				resource.Attribute{
					Name:        "internet_bandwidth_max_out",
					Description: `Max bandwidth out, only applicable to open CLB. Valid value ranges is [1, 2048]. Unit is MB.`,
				},
				resource.Attribute{
					Name:        "internet_charge_type",
					Description: `Internet charge type, only applicable to open CLB. Valid values are ` + "`" + `TRAFFIC_POSTPAID_BY_HOUR` + "`" + `, ` + "`" + `BANDWIDTH_POSTPAID_BY_HOUR` + "`" + ` and ` + "`" + `BANDWIDTH_PACKAGE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `Types of CLB.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of the project.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `ID set of the security groups.`,
				},
				resource.Attribute{
					Name:        "status_time",
					Description: `Latest state transition time of CLB.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of CLB.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of the subnet.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The available tags within this CLB.`,
				},
				resource.Attribute{
					Name:        "target_region_info_region",
					Description: `Region information of backend service are attached the CLB.`,
				},
				resource.Attribute{
					Name:        "target_region_info_vpc_id",
					Description: `VpcId information of backend service are attached the CLB.`,
				},
				resource.Attribute{
					Name:        "vip_isp",
					Description: `Network operator, only applicable to open CLB. Valid values are ` + "`" + `CMCC` + "`" + `(China Mobile), ` + "`" + `CTCC` + "`" + `(Telecom), ` + "`" + `CUCC` + "`" + `(China Unicom) and ` + "`" + `BGP` + "`" + `. If this ISP is specified, network billing method can only use the bandwidth package billing (BANDWIDTH_PACKAGE).`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_clb_listener_rules",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of CLB listener rule`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "clb_id",
					Description: `(Required) ID of the CLB to be queried.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Required) ID of the CLB listener to be queried.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional) Domain name of the forwarding rule to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `(Optional) ID of the forwarding rule to be queried.`,
				},
				resource.Attribute{
					Name:        "scheduler",
					Description: `(Optional) Scheduling method of the forwarding rule of thr CLB listener, and available values include ` + "`" + `WRR` + "`" + `, ` + "`" + `IP HASH` + "`" + ` and ` + "`" + `LEAST_CONN` + "`" + `. The default is ` + "`" + `WRR` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Optional) Url of the forwarding rule to be queried. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "rule_list",
					Description: `A list of forward rules of listeners. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "certificate_ca_id",
					Description: `ID of the client certificate. NOTES: Only supports listeners of 'HTTPS' and 'TCP_SSL' protocol.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `ID of the server certificate. NOTES: Only supports listeners of 'HTTPS' and 'TCP_SSL' protocol.`,
				},
				resource.Attribute{
					Name:        "certificate_ssl_mode",
					Description: `Type of SSL Mode, and available values inclue 'UNIDIRECTIONAL', 'MUTUAL'.NOTES: Only supports listeners of 'HTTPS' and 'TCP_SSL' protocol.`,
				},
				resource.Attribute{
					Name:        "clb_id",
					Description: `ID of the CLB.`,
				},
				resource.Attribute{
					Name:        "health_check_health_num",
					Description: `Health threshold of health check, and the default is ` + "`" + `3` + "`" + `. If a success result is returned for the health check three consecutive times, the CVM is identified as healthy. The value range is 2-10. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "health_check_http_code",
					Description: `HTTP Status Code. The default is 31 and value range is 1-31. 1 means the return value '1xx' is health. 2 means the return value '2xx' is health. 4 means the return value '3xx' is health. 8 means the return value 4xx is health. 16 means the return value '5xx' is health. If you want multiple return codes to indicate health, need to add the corresponding values. NOTES: The 'HTTP' health check of the 'TCP' listener only supports specifying one health check status code. NOTES: Only supports listeners of 'HTTP' and 'HTTPS' protocol.`,
				},
				resource.Attribute{
					Name:        "health_check_http_domain",
					Description: `Domain name of health check. NOTES: Only supports listeners of 'HTTPS' and 'HTTP' protocol.`,
				},
				resource.Attribute{
					Name:        "health_check_http_method",
					Description: `Methods of health check. NOTES: Only supports listeners of 'HTTPS' and 'HTTP' protocol. The default is 'HEAD', the available value include 'HEAD' and 'GET'.`,
				},
				resource.Attribute{
					Name:        "health_check_http_path",
					Description: `Path of health check. NOTES: Only supports listeners of 'HTTPS' and 'HTTP' protocol.`,
				},
				resource.Attribute{
					Name:        "health_check_interval_time",
					Description: `Interval time of health check. The value range is 5-300 sec, and the default is ` + "`" + `5` + "`" + ` sec. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "health_check_switch",
					Description: `Indicates whether health check is enabled.`,
				},
				resource.Attribute{
					Name:        "health_check_unhealth_num",
					Description: `Unhealth threshold of health check, and the default is ` + "`" + `3` + "`" + `. If a success result is returned for the health check three consecutive times, the CVM is identified as unhealthy. The value range is 2-10. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "http2_switch",
					Description: `Indicate to set HTTP2 protocol or not.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `ID of the listener.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `ID of the rule.`,
				},
				resource.Attribute{
					Name:        "scheduler",
					Description: `Scheduling method of the CLB listener, and available values include 'WRR', 'IP_HASH' and 'LEAST_CONN'. The default is 'WRR'. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "session_expire_time",
					Description: `Time of session persistence within the CLB listener. NOTES: Available when scheduler is specified as 'WRR'. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "rule_list",
					Description: `A list of forward rules of listeners. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "certificate_ca_id",
					Description: `ID of the client certificate. NOTES: Only supports listeners of 'HTTPS' and 'TCP_SSL' protocol.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `ID of the server certificate. NOTES: Only supports listeners of 'HTTPS' and 'TCP_SSL' protocol.`,
				},
				resource.Attribute{
					Name:        "certificate_ssl_mode",
					Description: `Type of SSL Mode, and available values inclue 'UNIDIRECTIONAL', 'MUTUAL'.NOTES: Only supports listeners of 'HTTPS' and 'TCP_SSL' protocol.`,
				},
				resource.Attribute{
					Name:        "clb_id",
					Description: `ID of the CLB.`,
				},
				resource.Attribute{
					Name:        "health_check_health_num",
					Description: `Health threshold of health check, and the default is ` + "`" + `3` + "`" + `. If a success result is returned for the health check three consecutive times, the CVM is identified as healthy. The value range is 2-10. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "health_check_http_code",
					Description: `HTTP Status Code. The default is 31 and value range is 1-31. 1 means the return value '1xx' is health. 2 means the return value '2xx' is health. 4 means the return value '3xx' is health. 8 means the return value 4xx is health. 16 means the return value '5xx' is health. If you want multiple return codes to indicate health, need to add the corresponding values. NOTES: The 'HTTP' health check of the 'TCP' listener only supports specifying one health check status code. NOTES: Only supports listeners of 'HTTP' and 'HTTPS' protocol.`,
				},
				resource.Attribute{
					Name:        "health_check_http_domain",
					Description: `Domain name of health check. NOTES: Only supports listeners of 'HTTPS' and 'HTTP' protocol.`,
				},
				resource.Attribute{
					Name:        "health_check_http_method",
					Description: `Methods of health check. NOTES: Only supports listeners of 'HTTPS' and 'HTTP' protocol. The default is 'HEAD', the available value include 'HEAD' and 'GET'.`,
				},
				resource.Attribute{
					Name:        "health_check_http_path",
					Description: `Path of health check. NOTES: Only supports listeners of 'HTTPS' and 'HTTP' protocol.`,
				},
				resource.Attribute{
					Name:        "health_check_interval_time",
					Description: `Interval time of health check. The value range is 5-300 sec, and the default is ` + "`" + `5` + "`" + ` sec. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "health_check_switch",
					Description: `Indicates whether health check is enabled.`,
				},
				resource.Attribute{
					Name:        "health_check_unhealth_num",
					Description: `Unhealth threshold of health check, and the default is ` + "`" + `3` + "`" + `. If a success result is returned for the health check three consecutive times, the CVM is identified as unhealthy. The value range is 2-10. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "http2_switch",
					Description: `Indicate to set HTTP2 protocol or not.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `ID of the listener.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `ID of the rule.`,
				},
				resource.Attribute{
					Name:        "scheduler",
					Description: `Scheduling method of the CLB listener, and available values include 'WRR', 'IP_HASH' and 'LEAST_CONN'. The default is 'WRR'. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "session_expire_time",
					Description: `Time of session persistence within the CLB listener. NOTES: Available when scheduler is specified as 'WRR'. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_clb_listeners",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of CLB listener`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "clb_id",
					Description: `(Required) Id of the CLB to be queried.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Optional) Id of the listener to be queried.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Port of the CLB listener.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) Type of protocol within the listener, and available values are ` + "`" + `TCP` + "`" + `, ` + "`" + `UDP` + "`" + `, ` + "`" + `HTTP` + "`" + `, ` + "`" + `HTTPS` + "`" + ` and ` + "`" + `TCP_SSL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "listener_list",
					Description: `A list of listeners of cloud load balancers. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "certificate_ca_id",
					Description: `Id of the client certificate. It must be set when SSLMode is ` + "`" + `mutual` + "`" + `. NOTES: only supported by listeners of ` + "`" + `HTTPS` + "`" + ` and ` + "`" + `TCP_SSL` + "`" + ` protocol.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `Id of the server certificate. It must be set when protocol is ` + "`" + `HTTPS` + "`" + ` or ` + "`" + `TCP_SSL` + "`" + `. NOTES: only supported by listeners of ` + "`" + `HTTPS` + "`" + ` and ` + "`" + `TCP_SSL` + "`" + ` protocol and must be set when it is available.`,
				},
				resource.Attribute{
					Name:        "certificate_ssl_mode",
					Description: `Type of certificate, and available values inclue ` + "`" + `UNIDIRECTIONAL` + "`" + `, ` + "`" + `MUTUAL` + "`" + `. NOTES: Only supports listeners of ` + "`" + `HTTPS` + "`" + ` and ` + "`" + `TCP_SSL` + "`" + ` protocol and must be set when it is available.`,
				},
				resource.Attribute{
					Name:        "clb_id",
					Description: `ID of the CLB.`,
				},
				resource.Attribute{
					Name:        "health_check_context_type",
					Description: `Health check protocol.`,
				},
				resource.Attribute{
					Name:        "health_check_health_num",
					Description: `Health threshold of health check, and the default is ` + "`" + `3` + "`" + `. If a success result is returned for the health check three consecutive times, the CVM is identified as healthy. The value range is 2-10. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "health_check_http_code",
					Description: `HTTP health check code of TCP listener.`,
				},
				resource.Attribute{
					Name:        "health_check_http_domain",
					Description: `HTTP health check domain of TCP listener.`,
				},
				resource.Attribute{
					Name:        "health_check_http_method",
					Description: `HTTP health check method of TCP listener.`,
				},
				resource.Attribute{
					Name:        "health_check_http_path",
					Description: `HTTP health check path of TCP listener.`,
				},
				resource.Attribute{
					Name:        "health_check_http_version",
					Description: `The HTTP version of the backend service.`,
				},
				resource.Attribute{
					Name:        "health_check_interval_time",
					Description: `Interval time of health check. The value range is 5-300 sec, and the default is ` + "`" + `5` + "`" + ` sec. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "health_check_port",
					Description: `The health check port is the port of the backend service.`,
				},
				resource.Attribute{
					Name:        "health_check_recv_context",
					Description: `It represents the result returned by the health check.`,
				},
				resource.Attribute{
					Name:        "health_check_send_context",
					Description: `It represents the content of the request sent by the health check.`,
				},
				resource.Attribute{
					Name:        "health_check_switch",
					Description: `Indicates whether health check is enabled.`,
				},
				resource.Attribute{
					Name:        "health_check_time_out",
					Description: `Response timeout of health check. The value range is 2-60 sec, and the default is ` + "`" + `2` + "`" + ` sec. Response timeout needs to be less than check interval. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration.`,
				},
				resource.Attribute{
					Name:        "health_check_type",
					Description: `Protocol used for health check.`,
				},
				resource.Attribute{
					Name:        "health_check_unhealth_num",
					Description: `Unhealthy threshold of health check, and the default is ` + "`" + `3` + "`" + `. If a success result is returned for the health check three consecutive times, the CVM is identified as unhealthy. The value range is 2-10. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `ID of the listener.`,
				},
				resource.Attribute{
					Name:        "listener_name",
					Description: `Name of the CLB listener.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port of the CLB listener.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol of the listener. Available values are ` + "`" + `HTTP` + "`" + `, ` + "`" + `HTTPS` + "`" + `, ` + "`" + `TCP` + "`" + `, ` + "`" + `UDP` + "`" + `, ` + "`" + `TCP_SSL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scheduler",
					Description: `Scheduling method of the CLB listener, and available values are ` + "`" + `WRR` + "`" + ` and ` + "`" + `LEAST_CONN` + "`" + `. The default is ` + "`" + `WRR` + "`" + `. NOTES: The listener of 'HTTP' and ` + "`" + `HTTPS` + "`" + ` protocol additionally supports the ` + "`" + `IP HASH` + "`" + ` method. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "session_expire_time",
					Description: `Time of session persistence within the CLB listener. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "sni_switch",
					Description: `Indicates whether SNI is enabled. NOTES: Only supported by ` + "`" + `HTTPS` + "`" + ` protocol.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "listener_list",
					Description: `A list of listeners of cloud load balancers. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "certificate_ca_id",
					Description: `Id of the client certificate. It must be set when SSLMode is ` + "`" + `mutual` + "`" + `. NOTES: only supported by listeners of ` + "`" + `HTTPS` + "`" + ` and ` + "`" + `TCP_SSL` + "`" + ` protocol.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `Id of the server certificate. It must be set when protocol is ` + "`" + `HTTPS` + "`" + ` or ` + "`" + `TCP_SSL` + "`" + `. NOTES: only supported by listeners of ` + "`" + `HTTPS` + "`" + ` and ` + "`" + `TCP_SSL` + "`" + ` protocol and must be set when it is available.`,
				},
				resource.Attribute{
					Name:        "certificate_ssl_mode",
					Description: `Type of certificate, and available values inclue ` + "`" + `UNIDIRECTIONAL` + "`" + `, ` + "`" + `MUTUAL` + "`" + `. NOTES: Only supports listeners of ` + "`" + `HTTPS` + "`" + ` and ` + "`" + `TCP_SSL` + "`" + ` protocol and must be set when it is available.`,
				},
				resource.Attribute{
					Name:        "clb_id",
					Description: `ID of the CLB.`,
				},
				resource.Attribute{
					Name:        "health_check_context_type",
					Description: `Health check protocol.`,
				},
				resource.Attribute{
					Name:        "health_check_health_num",
					Description: `Health threshold of health check, and the default is ` + "`" + `3` + "`" + `. If a success result is returned for the health check three consecutive times, the CVM is identified as healthy. The value range is 2-10. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "health_check_http_code",
					Description: `HTTP health check code of TCP listener.`,
				},
				resource.Attribute{
					Name:        "health_check_http_domain",
					Description: `HTTP health check domain of TCP listener.`,
				},
				resource.Attribute{
					Name:        "health_check_http_method",
					Description: `HTTP health check method of TCP listener.`,
				},
				resource.Attribute{
					Name:        "health_check_http_path",
					Description: `HTTP health check path of TCP listener.`,
				},
				resource.Attribute{
					Name:        "health_check_http_version",
					Description: `The HTTP version of the backend service.`,
				},
				resource.Attribute{
					Name:        "health_check_interval_time",
					Description: `Interval time of health check. The value range is 5-300 sec, and the default is ` + "`" + `5` + "`" + ` sec. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "health_check_port",
					Description: `The health check port is the port of the backend service.`,
				},
				resource.Attribute{
					Name:        "health_check_recv_context",
					Description: `It represents the result returned by the health check.`,
				},
				resource.Attribute{
					Name:        "health_check_send_context",
					Description: `It represents the content of the request sent by the health check.`,
				},
				resource.Attribute{
					Name:        "health_check_switch",
					Description: `Indicates whether health check is enabled.`,
				},
				resource.Attribute{
					Name:        "health_check_time_out",
					Description: `Response timeout of health check. The value range is 2-60 sec, and the default is ` + "`" + `2` + "`" + ` sec. Response timeout needs to be less than check interval. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration.`,
				},
				resource.Attribute{
					Name:        "health_check_type",
					Description: `Protocol used for health check.`,
				},
				resource.Attribute{
					Name:        "health_check_unhealth_num",
					Description: `Unhealthy threshold of health check, and the default is ` + "`" + `3` + "`" + `. If a success result is returned for the health check three consecutive times, the CVM is identified as unhealthy. The value range is 2-10. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `ID of the listener.`,
				},
				resource.Attribute{
					Name:        "listener_name",
					Description: `Name of the CLB listener.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port of the CLB listener.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol of the listener. Available values are ` + "`" + `HTTP` + "`" + `, ` + "`" + `HTTPS` + "`" + `, ` + "`" + `TCP` + "`" + `, ` + "`" + `UDP` + "`" + `, ` + "`" + `TCP_SSL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scheduler",
					Description: `Scheduling method of the CLB listener, and available values are ` + "`" + `WRR` + "`" + ` and ` + "`" + `LEAST_CONN` + "`" + `. The default is ` + "`" + `WRR` + "`" + `. NOTES: The listener of 'HTTP' and ` + "`" + `HTTPS` + "`" + ` protocol additionally supports the ` + "`" + `IP HASH` + "`" + ` method. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "session_expire_time",
					Description: `Time of session persistence within the CLB listener. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "sni_switch",
					Description: `Indicates whether SNI is enabled. NOTES: Only supported by ` + "`" + `HTTPS` + "`" + ` protocol.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_clb_redirections",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of CLB redirections`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "clb_id",
					Description: `(Required) ID of the CLB to be queried.`,
				},
				resource.Attribute{
					Name:        "source_listener_id",
					Description: `(Required) ID of source listener to be queried.`,
				},
				resource.Attribute{
					Name:        "source_rule_id",
					Description: `(Required) Rule ID of source listener to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "target_listener_id",
					Description: `(Optional) ID of target listener to be queried.`,
				},
				resource.Attribute{
					Name:        "target_rule_id",
					Description: `(Optional) Rule ID of target listener to be queried. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "redirection_list",
					Description: `A list of cloud load balancer redirection configurations. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "clb_id",
					Description: `ID of the CLB.`,
				},
				resource.Attribute{
					Name:        "source_listener_id",
					Description: `ID of source listener.`,
				},
				resource.Attribute{
					Name:        "source_rule_id",
					Description: `Rule ID of source listener.`,
				},
				resource.Attribute{
					Name:        "target_listener_id",
					Description: `ID of target listener.`,
				},
				resource.Attribute{
					Name:        "target_rule_id",
					Description: `Rule ID of target listener.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "redirection_list",
					Description: `A list of cloud load balancer redirection configurations. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "clb_id",
					Description: `ID of the CLB.`,
				},
				resource.Attribute{
					Name:        "source_listener_id",
					Description: `ID of source listener.`,
				},
				resource.Attribute{
					Name:        "source_rule_id",
					Description: `Rule ID of source listener.`,
				},
				resource.Attribute{
					Name:        "target_listener_id",
					Description: `ID of target listener.`,
				},
				resource.Attribute{
					Name:        "target_rule_id",
					Description: `Rule ID of target listener.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_clb_target_groups",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query target group information.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "target_group_id",
					Description: `(Optional) ID of Target group. Mutually exclusive with ` + "`" + `vpc_id` + "`" + ` and ` + "`" + `target_group_name` + "`" + `. ` + "`" + `target_group_id` + "`" + ` is preferred.`,
				},
				resource.Attribute{
					Name:        "target_group_name",
					Description: `(Optional) Name of target group. Mutually exclusive with ` + "`" + `target_group_id` + "`" + `. ` + "`" + `target_group_id` + "`" + ` is preferred.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) Target group VPC ID. Mutually exclusive with ` + "`" + `target_group_id` + "`" + `. ` + "`" + `target_group_id` + "`" + ` is preferred. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `Target group info list.`,
				},
				resource.Attribute{
					Name:        "associated_rule_list",
					Description: `List of associated rules.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Forwarding rule domain.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `Listener ID.`,
				},
				resource.Attribute{
					Name:        "listener_name",
					Description: `Listener name.`,
				},
				resource.Attribute{
					Name:        "listener_port",
					Description: `Listener port.`,
				},
				resource.Attribute{
					Name:        "load_balancer_id",
					Description: `Load balance ID.`,
				},
				resource.Attribute{
					Name:        "load_balancer_name",
					Description: `Load balance name.`,
				},
				resource.Attribute{
					Name:        "location_id",
					Description: `Forwarding rule ID.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Listener protocol type.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Forwarding rule URL.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the target group.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port of target group.`,
				},
				resource.Attribute{
					Name:        "target_group_id",
					Description: `ID of Target group.`,
				},
				resource.Attribute{
					Name:        "target_group_instance_list",
					Description: `List of backend servers bound to the target group.`,
				},
				resource.Attribute{
					Name:        "eni_id",
					Description: `ID of Elastic Network Interface.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of backend service.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `The instance name of the backend service.`,
				},
				resource.Attribute{
					Name:        "private_ip_addresses",
					Description: `Intranet IP list of back-end services.`,
				},
				resource.Attribute{
					Name:        "public_ip_addresses",
					Description: `List of external network IP of back-end services.`,
				},
				resource.Attribute{
					Name:        "registered_time",
					Description: `The time the backend service was bound.`,
				},
				resource.Attribute{
					Name:        "server_port",
					Description: `Port of backend service.`,
				},
				resource.Attribute{
					Name:        "server_type",
					Description: `Type of backend service.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Forwarding weight of back-end services.`,
				},
				resource.Attribute{
					Name:        "target_group_name",
					Description: `Target group VPC ID.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Modification time of the target group.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `Name of target group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `Target group info list.`,
				},
				resource.Attribute{
					Name:        "associated_rule_list",
					Description: `List of associated rules.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Forwarding rule domain.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `Listener ID.`,
				},
				resource.Attribute{
					Name:        "listener_name",
					Description: `Listener name.`,
				},
				resource.Attribute{
					Name:        "listener_port",
					Description: `Listener port.`,
				},
				resource.Attribute{
					Name:        "load_balancer_id",
					Description: `Load balance ID.`,
				},
				resource.Attribute{
					Name:        "load_balancer_name",
					Description: `Load balance name.`,
				},
				resource.Attribute{
					Name:        "location_id",
					Description: `Forwarding rule ID.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Listener protocol type.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Forwarding rule URL.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the target group.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port of target group.`,
				},
				resource.Attribute{
					Name:        "target_group_id",
					Description: `ID of Target group.`,
				},
				resource.Attribute{
					Name:        "target_group_instance_list",
					Description: `List of backend servers bound to the target group.`,
				},
				resource.Attribute{
					Name:        "eni_id",
					Description: `ID of Elastic Network Interface.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of backend service.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `The instance name of the backend service.`,
				},
				resource.Attribute{
					Name:        "private_ip_addresses",
					Description: `Intranet IP list of back-end services.`,
				},
				resource.Attribute{
					Name:        "public_ip_addresses",
					Description: `List of external network IP of back-end services.`,
				},
				resource.Attribute{
					Name:        "registered_time",
					Description: `The time the backend service was bound.`,
				},
				resource.Attribute{
					Name:        "server_port",
					Description: `Port of backend service.`,
				},
				resource.Attribute{
					Name:        "server_type",
					Description: `Type of backend service.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Forwarding weight of back-end services.`,
				},
				resource.Attribute{
					Name:        "target_group_name",
					Description: `Target group VPC ID.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Modification time of the target group.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `Name of target group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_container_cluster_instances",
			Category:         "Data Sources",
			ShortDescription: `Get all instances of the specific cluster.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) An ID identify the cluster, like cls-xxxxxx.`,
				},
				resource.Attribute{
					Name:        "limit",
					Description: `(Optional) An int variable describe how many instances in return at most. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `An information list of kubernetes instances.`,
				},
				resource.Attribute{
					Name:        "abnormal_reason",
					Description: `Describe the reason when node is in abnormal state(if it was).`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `Describe the cpu of the node.`,
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
					Description: `Describe the LAN IP of the node.`,
				},
				resource.Attribute{
					Name:        "mem",
					Description: `Describe the memory of the node.`,
				},
				resource.Attribute{
					Name:        "wan_ip",
					Description: `Describe the WAN IP of the node.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Number of instances.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "nodes",
					Description: `An information list of kubernetes instances.`,
				},
				resource.Attribute{
					Name:        "abnormal_reason",
					Description: `Describe the reason when node is in abnormal state(if it was).`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `Describe the cpu of the node.`,
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
					Description: `Describe the LAN IP of the node.`,
				},
				resource.Attribute{
					Name:        "mem",
					Description: `Describe the memory of the node.`,
				},
				resource.Attribute{
					Name:        "wan_ip",
					Description: `Describe the WAN IP of the node.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Number of instances.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_container_clusters",
			Category:         "Data Sources",
			ShortDescription: `Get container clusters in the current region.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Optional) An id identify the cluster, like ` + "`" + `cls-xxxxxx` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "limit",
					Description: `(Optional) An int variable describe how many cluster in return at most. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "clusters",
					Description: `An information list of kubernetes clusters.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `An id identify the cluster, like ` + "`" + `cls-xxxxxx` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `Name the cluster.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the cluster.`,
				},
				resource.Attribute{
					Name:        "kubernetes_version",
					Description: `Describe the running kubernetes version on the cluster.`,
				},
				resource.Attribute{
					Name:        "nodes_num",
					Description: `Describe how many cluster instances in the cluster.`,
				},
				resource.Attribute{
					Name:        "nodes_status",
					Description: `Describe the current status of the instances in the cluster.`,
				},
				resource.Attribute{
					Name:        "security_certification_authority",
					Description: `Describe the certificate string needed for using kubectl to access to kubernetes.`,
				},
				resource.Attribute{
					Name:        "security_cluster_external_endpoint",
					Description: `Describe the address needed for using kubectl to access to kubernetes.`,
				},
				resource.Attribute{
					Name:        "security_password",
					Description: `Describe the password needed for using kubectl to access to kubernetes.`,
				},
				resource.Attribute{
					Name:        "security_username",
					Description: `Describe the username needed for using kubectl to access to kubernetes.`,
				},
				resource.Attribute{
					Name:        "total_cpu",
					Description: `Describe the total cpu of each instance in the cluster.`,
				},
				resource.Attribute{
					Name:        "total_mem",
					Description: `Describe the total memory of each instance in the cluster.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Number of clusters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "clusters",
					Description: `An information list of kubernetes clusters.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `An id identify the cluster, like ` + "`" + `cls-xxxxxx` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `Name the cluster.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the cluster.`,
				},
				resource.Attribute{
					Name:        "kubernetes_version",
					Description: `Describe the running kubernetes version on the cluster.`,
				},
				resource.Attribute{
					Name:        "nodes_num",
					Description: `Describe how many cluster instances in the cluster.`,
				},
				resource.Attribute{
					Name:        "nodes_status",
					Description: `Describe the current status of the instances in the cluster.`,
				},
				resource.Attribute{
					Name:        "security_certification_authority",
					Description: `Describe the certificate string needed for using kubectl to access to kubernetes.`,
				},
				resource.Attribute{
					Name:        "security_cluster_external_endpoint",
					Description: `Describe the address needed for using kubectl to access to kubernetes.`,
				},
				resource.Attribute{
					Name:        "security_password",
					Description: `Describe the password needed for using kubectl to access to kubernetes.`,
				},
				resource.Attribute{
					Name:        "security_username",
					Description: `Describe the username needed for using kubectl to access to kubernetes.`,
				},
				resource.Attribute{
					Name:        "total_cpu",
					Description: `Describe the total cpu of each instance in the cluster.`,
				},
				resource.Attribute{
					Name:        "total_mem",
					Description: `Describe the total memory of each instance in the cluster.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Number of clusters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cos_bucket_object",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query the metadata of an object stored inside a bucket.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) Name of the bucket that contains the objects to query.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The full path to the object inside the bucket.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cache_control",
					Description: `Specifies caching behavior along the request/reply chain.`,
				},
				resource.Attribute{
					Name:        "content_disposition",
					Description: `Specifies presentational information for the object.`,
				},
				resource.Attribute{
					Name:        "content_encoding",
					Description: `Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `A standard MIME type describing the format of the object data.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `ETag generated for the object, which is may not equal to MD5 value.`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `Last modified date of the object.`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `Object storage type such as STANDARD.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cache_control",
					Description: `Specifies caching behavior along the request/reply chain.`,
				},
				resource.Attribute{
					Name:        "content_disposition",
					Description: `Specifies presentational information for the object.`,
				},
				resource.Attribute{
					Name:        "content_encoding",
					Description: `Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `A standard MIME type describing the format of the object data.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `ETag generated for the object, which is may not equal to MD5 value.`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `Last modified date of the object.`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `Object storage type such as STANDARD.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cos_buckets",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query the COS buckets of the current Tencent Cloud user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket_prefix",
					Description: `(Optional) A prefix string to filter results by bucket name.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags to filter bucket. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "bucket_list",
					Description: `A list of bucket. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "bucket",
					Description: `Bucket name, the format likes ` + "`" + `<bucket>-<appid>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cors_rules",
					Description: `A list of CORS rule configurations.`,
				},
				resource.Attribute{
					Name:        "allowed_headers",
					Description: `Specifies which headers are allowed.`,
				},
				resource.Attribute{
					Name:        "allowed_methods",
					Description: `Specifies which methods are allowed. Can be GET, PUT, POST, DELETE or HEAD.`,
				},
				resource.Attribute{
					Name:        "allowed_origins",
					Description: `Specifies which origins are allowed.`,
				},
				resource.Attribute{
					Name:        "expose_headers",
					Description: `Specifies expose header in the response.`,
				},
				resource.Attribute{
					Name:        "max_age_seconds",
					Description: `Specifies time in seconds that browser can cache the response for a preflight request.`,
				},
				resource.Attribute{
					Name:        "cos_bucket_url",
					Description: `The URL of this cos bucket.`,
				},
				resource.Attribute{
					Name:        "lifecycle_rules",
					Description: `The lifecycle configuration of a bucket.`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `Specifies a period in the object's expire.`,
				},
				resource.Attribute{
					Name:        "date",
					Description: `Specifies the date after which you want the corresponding action to take effect.`,
				},
				resource.Attribute{
					Name:        "days",
					Description: `Specifies the number of days after object creation when the specific rule action takes effect.`,
				},
				resource.Attribute{
					Name:        "filter_prefix",
					Description: `Object key prefix identifying one or more objects to which the rule applies.`,
				},
				resource.Attribute{
					Name:        "transition",
					Description: `Specifies a period in the object's transitions.`,
				},
				resource.Attribute{
					Name:        "date",
					Description: `Specifies the date after which you want the corresponding action to take effect.`,
				},
				resource.Attribute{
					Name:        "days",
					Description: `Specifies the number of days after object creation when the specific rule action takes effect.`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `Specifies the storage class to which you want the object to transition. Available values include STANDARD, STANDARD_IA and ARCHIVE.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The tags of a bucket.`,
				},
				resource.Attribute{
					Name:        "website",
					Description: `A list of one element containing configuration parameters used when the bucket is used as a website.`,
				},
				resource.Attribute{
					Name:        "error_document",
					Description: `An absolute path to the document to return in case of a 4XX error.`,
				},
				resource.Attribute{
					Name:        "index_document",
					Description: `COS returns this index document when requests are made to the root domain or any of the subfolders.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket_list",
					Description: `A list of bucket. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "bucket",
					Description: `Bucket name, the format likes ` + "`" + `<bucket>-<appid>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cors_rules",
					Description: `A list of CORS rule configurations.`,
				},
				resource.Attribute{
					Name:        "allowed_headers",
					Description: `Specifies which headers are allowed.`,
				},
				resource.Attribute{
					Name:        "allowed_methods",
					Description: `Specifies which methods are allowed. Can be GET, PUT, POST, DELETE or HEAD.`,
				},
				resource.Attribute{
					Name:        "allowed_origins",
					Description: `Specifies which origins are allowed.`,
				},
				resource.Attribute{
					Name:        "expose_headers",
					Description: `Specifies expose header in the response.`,
				},
				resource.Attribute{
					Name:        "max_age_seconds",
					Description: `Specifies time in seconds that browser can cache the response for a preflight request.`,
				},
				resource.Attribute{
					Name:        "cos_bucket_url",
					Description: `The URL of this cos bucket.`,
				},
				resource.Attribute{
					Name:        "lifecycle_rules",
					Description: `The lifecycle configuration of a bucket.`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `Specifies a period in the object's expire.`,
				},
				resource.Attribute{
					Name:        "date",
					Description: `Specifies the date after which you want the corresponding action to take effect.`,
				},
				resource.Attribute{
					Name:        "days",
					Description: `Specifies the number of days after object creation when the specific rule action takes effect.`,
				},
				resource.Attribute{
					Name:        "filter_prefix",
					Description: `Object key prefix identifying one or more objects to which the rule applies.`,
				},
				resource.Attribute{
					Name:        "transition",
					Description: `Specifies a period in the object's transitions.`,
				},
				resource.Attribute{
					Name:        "date",
					Description: `Specifies the date after which you want the corresponding action to take effect.`,
				},
				resource.Attribute{
					Name:        "days",
					Description: `Specifies the number of days after object creation when the specific rule action takes effect.`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `Specifies the storage class to which you want the object to transition. Available values include STANDARD, STANDARD_IA and ARCHIVE.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The tags of a bucket.`,
				},
				resource.Attribute{
					Name:        "website",
					Description: `A list of one element containing configuration parameters used when the bucket is used as a website.`,
				},
				resource.Attribute{
					Name:        "error_document",
					Description: `An absolute path to the document to return in case of a 4XX error.`,
				},
				resource.Attribute{
					Name:        "index_document",
					Description: `COS returns this index document when requests are made to the root domain or any of the subfolders.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cynosdb_clusters",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of Cynosdb clusters.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Optional) ID of the cluster to be queried.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Optional) Name of the cluster to be queried.`,
				},
				resource.Attribute{
					Name:        "db_type",
					Description: `(Optional) Type of CynosDB, and available values include ` + "`" + `MYSQL` + "`" + `, ` + "`" + `POSTGRESQL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) ID of the project to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cluster_list",
					Description: `A list of clusters. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "auto_renew_flag",
					Description: `Auto renew flag. Valid values are ` + "`" + `0` + "`" + `(MANUAL_RENEW), ` + "`" + `1` + "`" + `(AUTO_RENEW). Only works for PREPAID cluster.`,
				},
				resource.Attribute{
					Name:        "available_zone",
					Description: `The available zone of the CynosDB Cluster.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `The charge type of instance. Valid values are ` + "`" + `PREPAID` + "`" + ` and ` + "`" + `POSTPAID_BY_HOUR` + "`" + `. Default value is ` + "`" + `POSTPAID_BY_HOUR` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `ID of CynosDB cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_limit",
					Description: `Storage limit of CynosDB cluster instance, unit in GB.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `Name of CynosDB cluster.`,
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
					Name:        "db_type",
					Description: `Type of CynosDB, and available values include ` + "`" + `MYSQL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "db_version",
					Description: `Version of CynosDB, which is related to ` + "`" + `db_type` + "`" + `. For ` + "`" + `MYSQL` + "`" + `, available value is ` + "`" + `5.7` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port of CynosDB cluster.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of the project.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of the subnet within this VPC.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_list",
					Description: `A list of clusters. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "auto_renew_flag",
					Description: `Auto renew flag. Valid values are ` + "`" + `0` + "`" + `(MANUAL_RENEW), ` + "`" + `1` + "`" + `(AUTO_RENEW). Only works for PREPAID cluster.`,
				},
				resource.Attribute{
					Name:        "available_zone",
					Description: `The available zone of the CynosDB Cluster.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `The charge type of instance. Valid values are ` + "`" + `PREPAID` + "`" + ` and ` + "`" + `POSTPAID_BY_HOUR` + "`" + `. Default value is ` + "`" + `POSTPAID_BY_HOUR` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `ID of CynosDB cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_limit",
					Description: `Storage limit of CynosDB cluster instance, unit in GB.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `Name of CynosDB cluster.`,
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
					Name:        "db_type",
					Description: `Type of CynosDB, and available values include ` + "`" + `MYSQL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "db_version",
					Description: `Version of CynosDB, which is related to ` + "`" + `db_type` + "`" + `. For ` + "`" + `MYSQL` + "`" + `, available value is ` + "`" + `5.7` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port of CynosDB cluster.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of the project.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of the subnet within this VPC.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cynosdb_instances",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of Cynosdb instances.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Optional) ID of the cluster.`,
				},
				resource.Attribute{
					Name:        "db_type",
					Description: `(Optional) Type of CynosDB, and available values include ` + "`" + `MYSQL` + "`" + `, ` + "`" + `POSTGRESQL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Optional) ID of the Cynosdb instance to be queried.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Optional) Name of the Cynosdb instance to be queried.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) ID of the project to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "instance_list",
					Description: `A list of instances. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the CynosDB instance.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `Name of CynosDB instance.`,
				},
				resource.Attribute{
					Name:        "instance_status",
					Description: `Status of the Cynosdb instance.`,
				},
				resource.Attribute{
					Name:        "instance_storage_size",
					Description: `Storage size of the Cynosdb instance, unit in GB.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `Instance type. ` + "`" + `ro` + "`" + ` for readonly instance, ` + "`" + `rw` + "`" + ` for read and write instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_list",
					Description: `A list of instances. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the CynosDB instance.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `Name of CynosDB instance.`,
				},
				resource.Attribute{
					Name:        "instance_status",
					Description: `Status of the Cynosdb instance.`,
				},
				resource.Attribute{
					Name:        "instance_storage_size",
					Description: `Storage size of the Cynosdb instance, unit in GB.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `Instance type. ` + "`" + `ro` + "`" + ` for readonly instance, ` + "`" + `rw` + "`" + ` for read and write instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dayu_cc_http_policies",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query dayu CC http policies`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required) ID of the resource that the CC http policy works for.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `(Required) Type of the resource that the CC http policy works for, valid values are ` + "`" + `bgpip` + "`" + `, ` + "`" + `bgp` + "`" + `, ` + "`" + `bgp-multip` + "`" + ` and ` + "`" + `net` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the CC http policy to be queried.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Optional) Id of the CC http policy to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list of CC http policies. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action mode.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CC self-define http policy.`,
				},
				resource.Attribute{
					Name:        "frequency",
					Description: `Max frequency per minute.`,
				},
				resource.Attribute{
					Name:        "ip_list",
					Description: `IP of the CC self-define http policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the CC self-define http policy.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `ID of the CC self-define http policy.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `ID of the resource that the CC self-define http policy works for.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `Type of the resource that the CC self-define http policy works for.`,
				},
				resource.Attribute{
					Name:        "smode",
					Description: `Match mode.`,
				},
				resource.Attribute{
					Name:        "switch",
					Description: `Indicate the CC self-define http policy takes effect or not.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list of CC http policies. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action mode.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CC self-define http policy.`,
				},
				resource.Attribute{
					Name:        "frequency",
					Description: `Max frequency per minute.`,
				},
				resource.Attribute{
					Name:        "ip_list",
					Description: `IP of the CC self-define http policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the CC self-define http policy.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `ID of the CC self-define http policy.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `ID of the resource that the CC self-define http policy works for.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `Type of the resource that the CC self-define http policy works for.`,
				},
				resource.Attribute{
					Name:        "smode",
					Description: `Match mode.`,
				},
				resource.Attribute{
					Name:        "switch",
					Description: `Indicate the CC self-define http policy takes effect or not.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dayu_cc_https_policies",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query dayu CC https policies`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required) Id of the resource that the CC https policy works for.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `(Required) Type of the resource that the CC https policy works for, valid value is ` + "`" + `bgpip` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the CC https policy to be queried.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Optional) Id of the CC https policy to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list of CC https policies. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action mode.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CC self-define https policy.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Domain that the CC self-define https policy works for.`,
				},
				resource.Attribute{
					Name:        "ip_list",
					Description: `Ip of the CC self-define https policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the CC self-define https policy.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Id of the CC self-define https policy.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `ID of the resource that the CC self-define https policy works for.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `Type of the resource that the CC self-define https policy works for.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `Rule id of the domain that the CC self-define https policy works for.`,
				},
				resource.Attribute{
					Name:        "switch",
					Description: `Indicate the CC self-define https policy takes effect or not.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list of CC https policies. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action mode.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CC self-define https policy.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Domain that the CC self-define https policy works for.`,
				},
				resource.Attribute{
					Name:        "ip_list",
					Description: `Ip of the CC self-define https policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the CC self-define https policy.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Id of the CC self-define https policy.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `ID of the resource that the CC self-define https policy works for.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `Type of the resource that the CC self-define https policy works for.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `Rule id of the domain that the CC self-define https policy works for.`,
				},
				resource.Attribute{
					Name:        "switch",
					Description: `Indicate the CC self-define https policy takes effect or not.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dayu_ddos_policies",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query dayu DDoS policies`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_type",
					Description: `(Required) Type of the resource that the DDoS policy works for, valid values are ` + "`" + `bgpip` + "`" + `, ` + "`" + `bgp` + "`" + `, ` + "`" + `bgp-multip` + "`" + ` and ` + "`" + `net` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Optional) ID of the DDoS policy to be query.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list of DDoS policies. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the DDoS policy.`,
				},
				resource.Attribute{
					Name:        "drop_options",
					Description: `Option list of abnormal check of the DDoS policy.`,
				},
				resource.Attribute{
					Name:        "bad_conn_threshold",
					Description: `The number of new connections based on destination IP that trigger suppression of connections.`,
				},
				resource.Attribute{
					Name:        "check_sync_conn",
					Description: `Indicate whether to check null connection or not.`,
				},
				resource.Attribute{
					Name:        "conn_timeout",
					Description: `Connection timeout of abnormal connection check.`,
				},
				resource.Attribute{
					Name:        "d_conn_limit",
					Description: `The limit of concurrent connections based on destination IP.`,
				},
				resource.Attribute{
					Name:        "d_new_limit",
					Description: `The limit of new connections based on destination IP.`,
				},
				resource.Attribute{
					Name:        "drop_icmp",
					Description: `Indicate whether to drop ICMP protocol or not.`,
				},
				resource.Attribute{
					Name:        "drop_other",
					Description: `Indicate whether to drop other protocols(exclude TCP/UDP/ICMP) or not.`,
				},
				resource.Attribute{
					Name:        "drop_tcp",
					Description: `Indicate whether to drop TCP protocol or not.`,
				},
				resource.Attribute{
					Name:        "drop_udp",
					Description: `Indicate to drop UDP protocol or not.`,
				},
				resource.Attribute{
					Name:        "icmp_mbps_limit",
					Description: `The limit of ICMP traffic rate.`,
				},
				resource.Attribute{
					Name:        "null_conn_enable",
					Description: `Indicate to enable null connection or not.`,
				},
				resource.Attribute{
					Name:        "other_mbps_limit",
					Description: `The limit of other protocols(exclude TCP/UDP/ICMP) traffic rate.`,
				},
				resource.Attribute{
					Name:        "s_conn_limit",
					Description: `The limit of concurrent connections based on source IP.`,
				},
				resource.Attribute{
					Name:        "s_new_limit",
					Description: `The limit of new connections based on source IP.`,
				},
				resource.Attribute{
					Name:        "syn_limit",
					Description: `The limit of syn of abnormal connection check.`,
				},
				resource.Attribute{
					Name:        "syn_rate",
					Description: `The percentage of syn in ack of abnormal connection check.`,
				},
				resource.Attribute{
					Name:        "tcp_mbps_limit",
					Description: `The limit of TCP traffic.`,
				},
				resource.Attribute{
					Name:        "udp_mbps_limit",
					Description: `The limit of UDP traffic rate.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the DDoS policy.`,
				},
				resource.Attribute{
					Name:        "packet_filters",
					Description: `Message filter options list.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action of port to take.`,
				},
				resource.Attribute{
					Name:        "d_end_port",
					Description: `End port of the destination.`,
				},
				resource.Attribute{
					Name:        "d_start_port",
					Description: `Start port of the destination.`,
				},
				resource.Attribute{
					Name:        "depth",
					Description: `The depth of match.`,
				},
				resource.Attribute{
					Name:        "is_include",
					Description: `Indicate whether to include the key word/regular expression or not.`,
				},
				resource.Attribute{
					Name:        "match_begin",
					Description: `Indicate whether to check load or not.`,
				},
				resource.Attribute{
					Name:        "match_str",
					Description: `The key word or regular expression.`,
				},
				resource.Attribute{
					Name:        "match_type",
					Description: `Match type.`,
				},
				resource.Attribute{
					Name:        "offset",
					Description: `The offset of match.`,
				},
				resource.Attribute{
					Name:        "pkt_length_max",
					Description: `The max length of the packet.`,
				},
				resource.Attribute{
					Name:        "pkt_length_min",
					Description: `The minimum length of the packet.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol.`,
				},
				resource.Attribute{
					Name:        "s_end_port",
					Description: `End port of the source.`,
				},
				resource.Attribute{
					Name:        "s_start_port",
					Description: `Start port of the source.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Id of policy.`,
				},
				resource.Attribute{
					Name:        "port_filters",
					Description: `Port limits of abnormal check of the DDoS policy.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action of port to take.`,
				},
				resource.Attribute{
					Name:        "end_port",
					Description: `End port.`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `The type of forbidden port, and valid values are 0, 1, 2. 0 for destination port, 1 for source port and 2 for both destination and source posts.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol.`,
				},
				resource.Attribute{
					Name:        "start_port",
					Description: `Start port.`,
				},
				resource.Attribute{
					Name:        "scene_id",
					Description: `Id of policy case that the DDoS policy works for.`,
				},
				resource.Attribute{
					Name:        "watermark_filters",
					Description: `Watermark policy options, and only support one watermark policy at most.`,
				},
				resource.Attribute{
					Name:        "auto_remove",
					Description: `Indicate whether to auto-remove the watermark or not.`,
				},
				resource.Attribute{
					Name:        "offset",
					Description: `The offset of watermark.`,
				},
				resource.Attribute{
					Name:        "open_switch",
					Description: `Indicate whether to open watermark or not.`,
				},
				resource.Attribute{
					Name:        "tcp_port_list",
					Description: `Port range of TCP.`,
				},
				resource.Attribute{
					Name:        "udp_port_list",
					Description: `Port range of TCP.`,
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
					Name:        "list",
					Description: `A list of DDoS policies. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the DDoS policy.`,
				},
				resource.Attribute{
					Name:        "drop_options",
					Description: `Option list of abnormal check of the DDoS policy.`,
				},
				resource.Attribute{
					Name:        "bad_conn_threshold",
					Description: `The number of new connections based on destination IP that trigger suppression of connections.`,
				},
				resource.Attribute{
					Name:        "check_sync_conn",
					Description: `Indicate whether to check null connection or not.`,
				},
				resource.Attribute{
					Name:        "conn_timeout",
					Description: `Connection timeout of abnormal connection check.`,
				},
				resource.Attribute{
					Name:        "d_conn_limit",
					Description: `The limit of concurrent connections based on destination IP.`,
				},
				resource.Attribute{
					Name:        "d_new_limit",
					Description: `The limit of new connections based on destination IP.`,
				},
				resource.Attribute{
					Name:        "drop_icmp",
					Description: `Indicate whether to drop ICMP protocol or not.`,
				},
				resource.Attribute{
					Name:        "drop_other",
					Description: `Indicate whether to drop other protocols(exclude TCP/UDP/ICMP) or not.`,
				},
				resource.Attribute{
					Name:        "drop_tcp",
					Description: `Indicate whether to drop TCP protocol or not.`,
				},
				resource.Attribute{
					Name:        "drop_udp",
					Description: `Indicate to drop UDP protocol or not.`,
				},
				resource.Attribute{
					Name:        "icmp_mbps_limit",
					Description: `The limit of ICMP traffic rate.`,
				},
				resource.Attribute{
					Name:        "null_conn_enable",
					Description: `Indicate to enable null connection or not.`,
				},
				resource.Attribute{
					Name:        "other_mbps_limit",
					Description: `The limit of other protocols(exclude TCP/UDP/ICMP) traffic rate.`,
				},
				resource.Attribute{
					Name:        "s_conn_limit",
					Description: `The limit of concurrent connections based on source IP.`,
				},
				resource.Attribute{
					Name:        "s_new_limit",
					Description: `The limit of new connections based on source IP.`,
				},
				resource.Attribute{
					Name:        "syn_limit",
					Description: `The limit of syn of abnormal connection check.`,
				},
				resource.Attribute{
					Name:        "syn_rate",
					Description: `The percentage of syn in ack of abnormal connection check.`,
				},
				resource.Attribute{
					Name:        "tcp_mbps_limit",
					Description: `The limit of TCP traffic.`,
				},
				resource.Attribute{
					Name:        "udp_mbps_limit",
					Description: `The limit of UDP traffic rate.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the DDoS policy.`,
				},
				resource.Attribute{
					Name:        "packet_filters",
					Description: `Message filter options list.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action of port to take.`,
				},
				resource.Attribute{
					Name:        "d_end_port",
					Description: `End port of the destination.`,
				},
				resource.Attribute{
					Name:        "d_start_port",
					Description: `Start port of the destination.`,
				},
				resource.Attribute{
					Name:        "depth",
					Description: `The depth of match.`,
				},
				resource.Attribute{
					Name:        "is_include",
					Description: `Indicate whether to include the key word/regular expression or not.`,
				},
				resource.Attribute{
					Name:        "match_begin",
					Description: `Indicate whether to check load or not.`,
				},
				resource.Attribute{
					Name:        "match_str",
					Description: `The key word or regular expression.`,
				},
				resource.Attribute{
					Name:        "match_type",
					Description: `Match type.`,
				},
				resource.Attribute{
					Name:        "offset",
					Description: `The offset of match.`,
				},
				resource.Attribute{
					Name:        "pkt_length_max",
					Description: `The max length of the packet.`,
				},
				resource.Attribute{
					Name:        "pkt_length_min",
					Description: `The minimum length of the packet.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol.`,
				},
				resource.Attribute{
					Name:        "s_end_port",
					Description: `End port of the source.`,
				},
				resource.Attribute{
					Name:        "s_start_port",
					Description: `Start port of the source.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Id of policy.`,
				},
				resource.Attribute{
					Name:        "port_filters",
					Description: `Port limits of abnormal check of the DDoS policy.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action of port to take.`,
				},
				resource.Attribute{
					Name:        "end_port",
					Description: `End port.`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `The type of forbidden port, and valid values are 0, 1, 2. 0 for destination port, 1 for source port and 2 for both destination and source posts.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol.`,
				},
				resource.Attribute{
					Name:        "start_port",
					Description: `Start port.`,
				},
				resource.Attribute{
					Name:        "scene_id",
					Description: `Id of policy case that the DDoS policy works for.`,
				},
				resource.Attribute{
					Name:        "watermark_filters",
					Description: `Watermark policy options, and only support one watermark policy at most.`,
				},
				resource.Attribute{
					Name:        "auto_remove",
					Description: `Indicate whether to auto-remove the watermark or not.`,
				},
				resource.Attribute{
					Name:        "offset",
					Description: `The offset of watermark.`,
				},
				resource.Attribute{
					Name:        "open_switch",
					Description: `Indicate whether to open watermark or not.`,
				},
				resource.Attribute{
					Name:        "tcp_port_list",
					Description: `Port range of TCP.`,
				},
				resource.Attribute{
					Name:        "udp_port_list",
					Description: `Port range of TCP.`,
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
			Type:             "tencentcloud_dayu_ddos_policy_attachments",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of dayu DDoS policy attachments`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_type",
					Description: `(Required) Type of the resource that the DDoS policy works for, valid values are ` + "`" + `bgpip` + "`" + `, ` + "`" + `bgp` + "`" + `, ` + "`" + `bgp-multip` + "`" + ` and ` + "`" + `net` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Optional) Id of the policy to be queried.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Optional) ID of the attached resource to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "dayu_ddos_policy_attachment_list",
					Description: `A list of dayu DDoS policy attachments. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `ID of the policy.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `ID of the attached resource.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `Type of the resource that the DDoS policy works for.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "dayu_ddos_policy_attachment_list",
					Description: `A list of dayu DDoS policy attachments. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `ID of the policy.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `ID of the attached resource.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `Type of the resource that the DDoS policy works for.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dayu_ddos_policy_cases",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query dayu DDoS policy cases`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_type",
					Description: `(Required) Type of the resource that the DDoS policy case works for, valid values are ` + "`" + `bgpip` + "`" + `, ` + "`" + `bgp` + "`" + `, ` + "`" + `bgp-multip` + "`" + ` and ` + "`" + `net` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scene_id",
					Description: `(Required) ID of the DDoS policy case to be query.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list of DDoS policy cases. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "app_protocols",
					Description: `App protocol set of the DDoS policy case.`,
				},
				resource.Attribute{
					Name:        "app_type",
					Description: `App type of the DDoS policy case.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the DDoS policy case.`,
				},
				resource.Attribute{
					Name:        "has_abroad",
					Description: `Indicate whether the service involves overseas or not.`,
				},
				resource.Attribute{
					Name:        "has_initiate_tcp",
					Description: `Indicate whether the service actively initiates TCP requests or not.`,
				},
				resource.Attribute{
					Name:        "has_initiate_udp",
					Description: `Indicate whether the actively initiate UDP requests or not.`,
				},
				resource.Attribute{
					Name:        "has_vpn",
					Description: `Indicate whether the service involves VPN service or not.`,
				},
				resource.Attribute{
					Name:        "max_tcp_package_len",
					Description: `The max length of TCP message package.`,
				},
				resource.Attribute{
					Name:        "max_udp_package_len",
					Description: `The max length of UDP message package.`,
				},
				resource.Attribute{
					Name:        "min_tcp_package_len",
					Description: `The minimum length of TCP message package.`,
				},
				resource.Attribute{
					Name:        "min_udp_package_len",
					Description: `The minimum length of UDP message package.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the DDoS policy case.`,
				},
				resource.Attribute{
					Name:        "peer_tcp_port",
					Description: `The port that actively initiates TCP requests.`,
				},
				resource.Attribute{
					Name:        "peer_udp_port",
					Description: `The port that actively initiates UDP requests.`,
				},
				resource.Attribute{
					Name:        "platform_types",
					Description: `Platform set of the DDoS policy case.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `Type of the resource that the DDoS policy case works for.`,
				},
				resource.Attribute{
					Name:        "scene_id",
					Description: `ID of the DDoS policy case.`,
				},
				resource.Attribute{
					Name:        "tcp_end_port",
					Description: `End port of the TCP service.`,
				},
				resource.Attribute{
					Name:        "tcp_footprint",
					Description: `The fixed signature of TCP protocol load.`,
				},
				resource.Attribute{
					Name:        "tcp_start_port",
					Description: `Start port of the TCP service.`,
				},
				resource.Attribute{
					Name:        "udp_end_port",
					Description: `End port of the UDP service.`,
				},
				resource.Attribute{
					Name:        "udp_footprint",
					Description: `The fixed signature of TCP protocol load.`,
				},
				resource.Attribute{
					Name:        "udp_start_port",
					Description: `Start port of the UDP service.`,
				},
				resource.Attribute{
					Name:        "web_api_urls",
					Description: `Web API url set.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list of DDoS policy cases. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "app_protocols",
					Description: `App protocol set of the DDoS policy case.`,
				},
				resource.Attribute{
					Name:        "app_type",
					Description: `App type of the DDoS policy case.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the DDoS policy case.`,
				},
				resource.Attribute{
					Name:        "has_abroad",
					Description: `Indicate whether the service involves overseas or not.`,
				},
				resource.Attribute{
					Name:        "has_initiate_tcp",
					Description: `Indicate whether the service actively initiates TCP requests or not.`,
				},
				resource.Attribute{
					Name:        "has_initiate_udp",
					Description: `Indicate whether the actively initiate UDP requests or not.`,
				},
				resource.Attribute{
					Name:        "has_vpn",
					Description: `Indicate whether the service involves VPN service or not.`,
				},
				resource.Attribute{
					Name:        "max_tcp_package_len",
					Description: `The max length of TCP message package.`,
				},
				resource.Attribute{
					Name:        "max_udp_package_len",
					Description: `The max length of UDP message package.`,
				},
				resource.Attribute{
					Name:        "min_tcp_package_len",
					Description: `The minimum length of TCP message package.`,
				},
				resource.Attribute{
					Name:        "min_udp_package_len",
					Description: `The minimum length of UDP message package.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the DDoS policy case.`,
				},
				resource.Attribute{
					Name:        "peer_tcp_port",
					Description: `The port that actively initiates TCP requests.`,
				},
				resource.Attribute{
					Name:        "peer_udp_port",
					Description: `The port that actively initiates UDP requests.`,
				},
				resource.Attribute{
					Name:        "platform_types",
					Description: `Platform set of the DDoS policy case.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `Type of the resource that the DDoS policy case works for.`,
				},
				resource.Attribute{
					Name:        "scene_id",
					Description: `ID of the DDoS policy case.`,
				},
				resource.Attribute{
					Name:        "tcp_end_port",
					Description: `End port of the TCP service.`,
				},
				resource.Attribute{
					Name:        "tcp_footprint",
					Description: `The fixed signature of TCP protocol load.`,
				},
				resource.Attribute{
					Name:        "tcp_start_port",
					Description: `Start port of the TCP service.`,
				},
				resource.Attribute{
					Name:        "udp_end_port",
					Description: `End port of the UDP service.`,
				},
				resource.Attribute{
					Name:        "udp_footprint",
					Description: `The fixed signature of TCP protocol load.`,
				},
				resource.Attribute{
					Name:        "udp_start_port",
					Description: `Start port of the UDP service.`,
				},
				resource.Attribute{
					Name:        "web_api_urls",
					Description: `Web API url set.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dayu_l4_rules",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query dayu layer 4 rules`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required) Id of the resource that the layer 4 rule works for.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `(Required) Type of the resource that the layer 4 rule works for, valid values are ` + "`" + `bgpip` + "`" + `, ` + "`" + `bgp` + "`" + `, ` + "`" + `bgp-multip` + "`" + ` and ` + "`" + `net` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the layer 4 rule to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `(Optional) Id of the layer 4 rule to be queried. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list of layer 4 rules. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "d_port",
					Description: `The destination port of the layer 4 rule.`,
				},
				resource.Attribute{
					Name:        "health_check_health_num",
					Description: `Health threshold of health check.`,
				},
				resource.Attribute{
					Name:        "health_check_interval",
					Description: `Interval time of health check.`,
				},
				resource.Attribute{
					Name:        "health_check_switch",
					Description: `Indicates whether health check is enabled.`,
				},
				resource.Attribute{
					Name:        "health_check_timeout",
					Description: `HTTP Status Code. ` + "`" + `1` + "`" + ` means the return value ` + "`" + `1xx` + "`" + ` is health. ` + "`" + `2` + "`" + ` means the return value ` + "`" + `2xx` + "`" + ` is health. ` + "`" + `4` + "`" + ` means the return value ` + "`" + `3xx` + "`" + ` is health. ` + "`" + `8` + "`" + ` means the return value ` + "`" + `4xx` + "`" + ` is health. ` + "`" + `16` + "`" + ` means the return value ` + "`" + `5xx` + "`" + ` is health. If you want multiple return codes to indicate health, need to add the corresponding values.`,
				},
				resource.Attribute{
					Name:        "health_check_unhealth_num",
					Description: `Unhealthy threshold of health check.`,
				},
				resource.Attribute{
					Name:        "lb_type",
					Description: `LB type of the rule, ` + "`" + `1` + "`" + ` for weight cycling and ` + "`" + `2` + "`" + ` for IP hash.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the rule.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol of the rule.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `ID of the 4 layer rule.`,
				},
				resource.Attribute{
					Name:        "s_port",
					Description: `The source port of the layer 4 rule.`,
				},
				resource.Attribute{
					Name:        "session_switch",
					Description: `Indicate that the session will keep or not.`,
				},
				resource.Attribute{
					Name:        "session_time",
					Description: `Session keep time, only valid when ` + "`" + `session_switch` + "`" + ` is true, the available value ranges from 1 to 300 and unit is second.`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `Source type, ` + "`" + `1` + "`" + ` for source of host, ` + "`" + `2` + "`" + ` for source of IP.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list of layer 4 rules. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "d_port",
					Description: `The destination port of the layer 4 rule.`,
				},
				resource.Attribute{
					Name:        "health_check_health_num",
					Description: `Health threshold of health check.`,
				},
				resource.Attribute{
					Name:        "health_check_interval",
					Description: `Interval time of health check.`,
				},
				resource.Attribute{
					Name:        "health_check_switch",
					Description: `Indicates whether health check is enabled.`,
				},
				resource.Attribute{
					Name:        "health_check_timeout",
					Description: `HTTP Status Code. ` + "`" + `1` + "`" + ` means the return value ` + "`" + `1xx` + "`" + ` is health. ` + "`" + `2` + "`" + ` means the return value ` + "`" + `2xx` + "`" + ` is health. ` + "`" + `4` + "`" + ` means the return value ` + "`" + `3xx` + "`" + ` is health. ` + "`" + `8` + "`" + ` means the return value ` + "`" + `4xx` + "`" + ` is health. ` + "`" + `16` + "`" + ` means the return value ` + "`" + `5xx` + "`" + ` is health. If you want multiple return codes to indicate health, need to add the corresponding values.`,
				},
				resource.Attribute{
					Name:        "health_check_unhealth_num",
					Description: `Unhealthy threshold of health check.`,
				},
				resource.Attribute{
					Name:        "lb_type",
					Description: `LB type of the rule, ` + "`" + `1` + "`" + ` for weight cycling and ` + "`" + `2` + "`" + ` for IP hash.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the rule.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol of the rule.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `ID of the 4 layer rule.`,
				},
				resource.Attribute{
					Name:        "s_port",
					Description: `The source port of the layer 4 rule.`,
				},
				resource.Attribute{
					Name:        "session_switch",
					Description: `Indicate that the session will keep or not.`,
				},
				resource.Attribute{
					Name:        "session_time",
					Description: `Session keep time, only valid when ` + "`" + `session_switch` + "`" + ` is true, the available value ranges from 1 to 300 and unit is second.`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `Source type, ` + "`" + `1` + "`" + ` for source of host, ` + "`" + `2` + "`" + ` for source of IP.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dayu_l7_rules",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query dayu layer 7 rules`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required) Id of the resource that the layer 7 rule works for.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `(Required) Type of the resource that the layer 7 rule works for, valid value is ` + "`" + `bgpip` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional) Domain of the layer 7 rule to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `(Optional) Id of the layer 7 rule to be queried. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list of layer 7 rules. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Domain that the 7 layer rule works for.`,
				},
				resource.Attribute{
					Name:        "health_check_code",
					Description: `HTTP Status Code. ` + "`" + `1` + "`" + ` means the return value ` + "`" + `1xx` + "`" + ` is health. ` + "`" + `2` + "`" + ` means the return value ` + "`" + `2xx` + "`" + ` is health. ` + "`" + `4` + "`" + ` means the return value ` + "`" + `3xx` + "`" + ` is health. ` + "`" + `8` + "`" + ` means the return value ` + "`" + `4xx` + "`" + ` is health. ` + "`" + `16` + "`" + ` means the return value ` + "`" + `5xx` + "`" + ` is health. If you want multiple return codes to indicate health, need to add the corresponding values.`,
				},
				resource.Attribute{
					Name:        "health_check_health_num",
					Description: `Health threshold of health check.`,
				},
				resource.Attribute{
					Name:        "health_check_interval",
					Description: `Interval time of health check.`,
				},
				resource.Attribute{
					Name:        "health_check_method",
					Description: `Methods of health check.`,
				},
				resource.Attribute{
					Name:        "health_check_path",
					Description: `Path of health check.`,
				},
				resource.Attribute{
					Name:        "health_check_switch",
					Description: `Indicates whether health check is enabled.`,
				},
				resource.Attribute{
					Name:        "health_check_unhealth_num",
					Description: `Unhealthy threshold of health check.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the rule.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol of the rule.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `Id of the 7 layer rule.`,
				},
				resource.Attribute{
					Name:        "source_list",
					Description: `Source list of the rule.`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `Source type, 1 for source of host, 2 for source of ip.`,
				},
				resource.Attribute{
					Name:        "ssl_id",
					Description: `SSL id.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the rule. ` + "`" + `0` + "`" + ` for create/modify success, ` + "`" + `2` + "`" + ` for create/modify fail, ` + "`" + `3` + "`" + ` for delete success, ` + "`" + `5` + "`" + ` for waiting to be created/modified, ` + "`" + `7` + "`" + ` for waiting to be deleted and ` + "`" + `8` + "`" + ` for waiting to get SSL id.`,
				},
				resource.Attribute{
					Name:        "switch",
					Description: `Indicate the rule will take effect or not.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `Threshold of the rule.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list of layer 7 rules. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Domain that the 7 layer rule works for.`,
				},
				resource.Attribute{
					Name:        "health_check_code",
					Description: `HTTP Status Code. ` + "`" + `1` + "`" + ` means the return value ` + "`" + `1xx` + "`" + ` is health. ` + "`" + `2` + "`" + ` means the return value ` + "`" + `2xx` + "`" + ` is health. ` + "`" + `4` + "`" + ` means the return value ` + "`" + `3xx` + "`" + ` is health. ` + "`" + `8` + "`" + ` means the return value ` + "`" + `4xx` + "`" + ` is health. ` + "`" + `16` + "`" + ` means the return value ` + "`" + `5xx` + "`" + ` is health. If you want multiple return codes to indicate health, need to add the corresponding values.`,
				},
				resource.Attribute{
					Name:        "health_check_health_num",
					Description: `Health threshold of health check.`,
				},
				resource.Attribute{
					Name:        "health_check_interval",
					Description: `Interval time of health check.`,
				},
				resource.Attribute{
					Name:        "health_check_method",
					Description: `Methods of health check.`,
				},
				resource.Attribute{
					Name:        "health_check_path",
					Description: `Path of health check.`,
				},
				resource.Attribute{
					Name:        "health_check_switch",
					Description: `Indicates whether health check is enabled.`,
				},
				resource.Attribute{
					Name:        "health_check_unhealth_num",
					Description: `Unhealthy threshold of health check.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the rule.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol of the rule.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `Id of the 7 layer rule.`,
				},
				resource.Attribute{
					Name:        "source_list",
					Description: `Source list of the rule.`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `Source type, 1 for source of host, 2 for source of ip.`,
				},
				resource.Attribute{
					Name:        "ssl_id",
					Description: `SSL id.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the rule. ` + "`" + `0` + "`" + ` for create/modify success, ` + "`" + `2` + "`" + ` for create/modify fail, ` + "`" + `3` + "`" + ` for delete success, ` + "`" + `5` + "`" + ` for waiting to be created/modified, ` + "`" + `7` + "`" + ` for waiting to be deleted and ` + "`" + `8` + "`" + ` for waiting to get SSL id.`,
				},
				resource.Attribute{
					Name:        "switch",
					Description: `Indicate the rule will take effect or not.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `Threshold of the rule.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dc_gateway_ccn_routes",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of direct connect gateway route entries.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dcg_id",
					Description: `(Required) ID of the DCG to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "instance_list",
					Description: `Information list of the DCG route entries.`,
				},
				resource.Attribute{
					Name:        "as_path",
					Description: `As path list of the BGP.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `A network address segment of IDC.`,
				},
				resource.Attribute{
					Name:        "dcg_id",
					Description: `ID of the DCG.`,
				},
				resource.Attribute{
					Name:        "route_id",
					Description: `ID of the DCG route.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_list",
					Description: `Information list of the DCG route entries.`,
				},
				resource.Attribute{
					Name:        "as_path",
					Description: `As path list of the BGP.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `A network address segment of IDC.`,
				},
				resource.Attribute{
					Name:        "dcg_id",
					Description: `ID of the DCG.`,
				},
				resource.Attribute{
					Name:        "route_id",
					Description: `ID of the DCG route.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dc_gateway_instances",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of direct connect gateway instances.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dcg_id",
					Description: `(Optional) ID of the DCG to be queried.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the DCG to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "instance_list",
					Description: `Information list of the DCG.`,
				},
				resource.Attribute{
					Name:        "cnn_route_type",
					Description: `Type of CCN route. Valid values: ` + "`" + `BGP` + "`" + ` and ` + "`" + `STATIC` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of resource.`,
				},
				resource.Attribute{
					Name:        "dcg_id",
					Description: `ID of the DCG.`,
				},
				resource.Attribute{
					Name:        "dcg_ip",
					Description: `IP of the DCG.`,
				},
				resource.Attribute{
					Name:        "enable_bgp",
					Description: `Indicates whether the BGP is enabled.`,
				},
				resource.Attribute{
					Name:        "gateway_type",
					Description: `Type of the gateway. Valid values: ` + "`" + `NORMAL` + "`" + ` and ` + "`" + `NAT` + "`" + `. Default is ` + "`" + `NORMAL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the DCG.`,
				},
				resource.Attribute{
					Name:        "network_instance_id",
					Description: `Type of associated network. Valid values: ` + "`" + `VPC` + "`" + ` and ` + "`" + `CCN` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `IP of the DCG.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_list",
					Description: `Information list of the DCG.`,
				},
				resource.Attribute{
					Name:        "cnn_route_type",
					Description: `Type of CCN route. Valid values: ` + "`" + `BGP` + "`" + ` and ` + "`" + `STATIC` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of resource.`,
				},
				resource.Attribute{
					Name:        "dcg_id",
					Description: `ID of the DCG.`,
				},
				resource.Attribute{
					Name:        "dcg_ip",
					Description: `IP of the DCG.`,
				},
				resource.Attribute{
					Name:        "enable_bgp",
					Description: `Indicates whether the BGP is enabled.`,
				},
				resource.Attribute{
					Name:        "gateway_type",
					Description: `Type of the gateway. Valid values: ` + "`" + `NORMAL` + "`" + ` and ` + "`" + `NAT` + "`" + `. Default is ` + "`" + `NORMAL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the DCG.`,
				},
				resource.Attribute{
					Name:        "network_instance_id",
					Description: `Type of associated network. Valid values: ` + "`" + `VPC` + "`" + ` and ` + "`" + `CCN` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `IP of the DCG.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dc_instances",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of DC instances.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dc_id",
					Description: `(Optional) ID of the DC to be queried.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the DC to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "instance_list",
					Description: `Information list of the DC.`,
				},
				resource.Attribute{
					Name:        "access_point_id",
					Description: `Access point ID of tne DC.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `Bandwidth of the DC.`,
				},
				resource.Attribute{
					Name:        "circuit_code",
					Description: `The circuit code provided by the operator for the DC.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of resource.`,
				},
				resource.Attribute{
					Name:        "customer_address",
					Description: `Interconnect IP of the DC within client. Note: This field may return null, indicating that no valid values are taken.`,
				},
				resource.Attribute{
					Name:        "customer_email",
					Description: `Applicant email of the DC, the default is obtained from the account. Note: This field may return null, indicating that no valid values are taken.`,
				},
				resource.Attribute{
					Name:        "customer_name",
					Description: `Applicant name of the DC, the default is obtained from the account. Note: This field may return null, indicating that no valid values are taken.`,
				},
				resource.Attribute{
					Name:        "customer_phone",
					Description: `Applicant phone number of the DC, the default is obtained from the account. Note: This field may return null, indicating that no valid values are taken.`,
				},
				resource.Attribute{
					Name:        "dc_id",
					Description: `ID of the DC.`,
				},
				resource.Attribute{
					Name:        "enabled_time",
					Description: `Enable time of resource.`,
				},
				resource.Attribute{
					Name:        "expired_time",
					Description: `Expire date of resource.`,
				},
				resource.Attribute{
					Name:        "fault_report_contact_person",
					Description: `Contact of reporting a faulty. Note: This field may return null, indicating that no valid values are taken.`,
				},
				resource.Attribute{
					Name:        "fault_report_contact_phone",
					Description: `Phone number of reporting a faulty. Note: This field may return null, indicating that no valid values are taken.`,
				},
				resource.Attribute{
					Name:        "line_operator",
					Description: `Operator of the DC, and available values include ` + "`" + `ChinaTelecom` + "`" + `, ` + "`" + `ChinaMobile` + "`" + `, ` + "`" + `ChinaUnicom` + "`" + `, ` + "`" + `In-houseWiring` + "`" + `, ` + "`" + `ChinaOther` + "`" + ` and ` + "`" + `InternationalOperator` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The DC location where the connection is located.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the DC.`,
				},
				resource.Attribute{
					Name:        "port_type",
					Description: `Port type of the DC in client, and available values include ` + "`" + `100Base-T` + "`" + `, ` + "`" + `1000Base-T` + "`" + `, ` + "`" + `1000Base-LX` + "`" + `, ` + "`" + `10GBase-T` + "`" + ` and ` + "`" + `10GBase-LR` + "`" + `. The default value is ` + "`" + `1000Base-LX` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "redundant_dc_id",
					Description: `ID of the redundant DC.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the DC, and available values include ` + "`" + `REJECTED` + "`" + `, ` + "`" + `TOPAY` + "`" + `, ` + "`" + `PAID` + "`" + `, ` + "`" + `ALLOCATED` + "`" + `, ` + "`" + `AVAILABLE` + "`" + `, ` + "`" + `DELETING` + "`" + ` and ` + "`" + `DELETED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tencent_address",
					Description: `Interconnect IP of the DC within Tencent. Note: This field may return null, indicating that no valid values are taken.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_list",
					Description: `Information list of the DC.`,
				},
				resource.Attribute{
					Name:        "access_point_id",
					Description: `Access point ID of tne DC.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `Bandwidth of the DC.`,
				},
				resource.Attribute{
					Name:        "circuit_code",
					Description: `The circuit code provided by the operator for the DC.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of resource.`,
				},
				resource.Attribute{
					Name:        "customer_address",
					Description: `Interconnect IP of the DC within client. Note: This field may return null, indicating that no valid values are taken.`,
				},
				resource.Attribute{
					Name:        "customer_email",
					Description: `Applicant email of the DC, the default is obtained from the account. Note: This field may return null, indicating that no valid values are taken.`,
				},
				resource.Attribute{
					Name:        "customer_name",
					Description: `Applicant name of the DC, the default is obtained from the account. Note: This field may return null, indicating that no valid values are taken.`,
				},
				resource.Attribute{
					Name:        "customer_phone",
					Description: `Applicant phone number of the DC, the default is obtained from the account. Note: This field may return null, indicating that no valid values are taken.`,
				},
				resource.Attribute{
					Name:        "dc_id",
					Description: `ID of the DC.`,
				},
				resource.Attribute{
					Name:        "enabled_time",
					Description: `Enable time of resource.`,
				},
				resource.Attribute{
					Name:        "expired_time",
					Description: `Expire date of resource.`,
				},
				resource.Attribute{
					Name:        "fault_report_contact_person",
					Description: `Contact of reporting a faulty. Note: This field may return null, indicating that no valid values are taken.`,
				},
				resource.Attribute{
					Name:        "fault_report_contact_phone",
					Description: `Phone number of reporting a faulty. Note: This field may return null, indicating that no valid values are taken.`,
				},
				resource.Attribute{
					Name:        "line_operator",
					Description: `Operator of the DC, and available values include ` + "`" + `ChinaTelecom` + "`" + `, ` + "`" + `ChinaMobile` + "`" + `, ` + "`" + `ChinaUnicom` + "`" + `, ` + "`" + `In-houseWiring` + "`" + `, ` + "`" + `ChinaOther` + "`" + ` and ` + "`" + `InternationalOperator` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The DC location where the connection is located.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the DC.`,
				},
				resource.Attribute{
					Name:        "port_type",
					Description: `Port type of the DC in client, and available values include ` + "`" + `100Base-T` + "`" + `, ` + "`" + `1000Base-T` + "`" + `, ` + "`" + `1000Base-LX` + "`" + `, ` + "`" + `10GBase-T` + "`" + ` and ` + "`" + `10GBase-LR` + "`" + `. The default value is ` + "`" + `1000Base-LX` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "redundant_dc_id",
					Description: `ID of the redundant DC.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the DC, and available values include ` + "`" + `REJECTED` + "`" + `, ` + "`" + `TOPAY` + "`" + `, ` + "`" + `PAID` + "`" + `, ` + "`" + `ALLOCATED` + "`" + `, ` + "`" + `AVAILABLE` + "`" + `, ` + "`" + `DELETING` + "`" + ` and ` + "`" + `DELETED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tencent_address",
					Description: `Interconnect IP of the DC within Tencent. Note: This field may return null, indicating that no valid values are taken.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dcx_instances",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of dedicated tunnels instances.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dcx_id",
					Description: `(Optional) ID of the dedicated tunnels to be queried.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the dedicated tunnels to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "instance_list",
					Description: `Information list of the dedicated tunnels.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `Bandwidth of the DC.`,
				},
				resource.Attribute{
					Name:        "bgp_asn",
					Description: `BGP ASN of the user.`,
				},
				resource.Attribute{
					Name:        "bgp_auth_key",
					Description: `BGP key of the user.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of resource.`,
				},
				resource.Attribute{
					Name:        "customer_address",
					Description: `Interconnect IP of the DC within client.`,
				},
				resource.Attribute{
					Name:        "dc_id",
					Description: `ID of the DC.`,
				},
				resource.Attribute{
					Name:        "dcg_id",
					Description: `ID of the DC Gateway. Currently only new in the console.`,
				},
				resource.Attribute{
					Name:        "dcx_id",
					Description: `ID of the dedicated tunnel.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the dedicated tunnel.`,
				},
				resource.Attribute{
					Name:        "network_region",
					Description: `The region of the dedicated tunnel.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `Type of the network. Valid values: ` + "`" + `VPC` + "`" + `, ` + "`" + `BMVPC` + "`" + ` and ` + "`" + `CCN` + "`" + `. The default value is ` + "`" + `VPC` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "route_filter_prefixes",
					Description: `Static route, the network address of the user IDC.`,
				},
				resource.Attribute{
					Name:        "route_type",
					Description: `Type of the route. Valid values: ` + "`" + `BGP` + "`" + ` and ` + "`" + `STATIC` + "`" + `. The default value is ` + "`" + `BGP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the dedicated tunnels. Valid values: ` + "`" + `PENDING` + "`" + `, ` + "`" + `ALLOCATING` + "`" + `, ` + "`" + `ALLOCATED` + "`" + `, ` + "`" + `ALTERING` + "`" + `, ` + "`" + `DELETING` + "`" + `, ` + "`" + `DELETED` + "`" + `, ` + "`" + `COMFIRMING` + "`" + ` and ` + "`" + `REJECTED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tencent_address",
					Description: `Interconnect IP of the DC within Tencent.`,
				},
				resource.Attribute{
					Name:        "vlan",
					Description: `Vlan of the dedicated tunnels. Valid value ranges: [0-3000]. ` + "`" + `0` + "`" + ` means that only one tunnel can be created for the physical connect.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC or BMVPC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_list",
					Description: `Information list of the dedicated tunnels.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `Bandwidth of the DC.`,
				},
				resource.Attribute{
					Name:        "bgp_asn",
					Description: `BGP ASN of the user.`,
				},
				resource.Attribute{
					Name:        "bgp_auth_key",
					Description: `BGP key of the user.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of resource.`,
				},
				resource.Attribute{
					Name:        "customer_address",
					Description: `Interconnect IP of the DC within client.`,
				},
				resource.Attribute{
					Name:        "dc_id",
					Description: `ID of the DC.`,
				},
				resource.Attribute{
					Name:        "dcg_id",
					Description: `ID of the DC Gateway. Currently only new in the console.`,
				},
				resource.Attribute{
					Name:        "dcx_id",
					Description: `ID of the dedicated tunnel.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the dedicated tunnel.`,
				},
				resource.Attribute{
					Name:        "network_region",
					Description: `The region of the dedicated tunnel.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `Type of the network. Valid values: ` + "`" + `VPC` + "`" + `, ` + "`" + `BMVPC` + "`" + ` and ` + "`" + `CCN` + "`" + `. The default value is ` + "`" + `VPC` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "route_filter_prefixes",
					Description: `Static route, the network address of the user IDC.`,
				},
				resource.Attribute{
					Name:        "route_type",
					Description: `Type of the route. Valid values: ` + "`" + `BGP` + "`" + ` and ` + "`" + `STATIC` + "`" + `. The default value is ` + "`" + `BGP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the dedicated tunnels. Valid values: ` + "`" + `PENDING` + "`" + `, ` + "`" + `ALLOCATING` + "`" + `, ` + "`" + `ALLOCATED` + "`" + `, ` + "`" + `ALTERING` + "`" + `, ` + "`" + `DELETING` + "`" + `, ` + "`" + `DELETED` + "`" + `, ` + "`" + `COMFIRMING` + "`" + ` and ` + "`" + `REJECTED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tencent_address",
					Description: `Interconnect IP of the DC within Tencent.`,
				},
				resource.Attribute{
					Name:        "vlan",
					Description: `Vlan of the dedicated tunnels. Valid value ranges: [0-3000]. ` + "`" + `0` + "`" + ` means that only one tunnel can be created for the physical connect.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC or BMVPC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dnats",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of DNATs.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the NAT forward.`,
				},
				resource.Attribute{
					Name:        "elastic_ip",
					Description: `(Optional) Network address of the EIP.`,
				},
				resource.Attribute{
					Name:        "elastic_port",
					Description: `(Optional) Port of the EIP.`,
				},
				resource.Attribute{
					Name:        "nat_id",
					Description: `(Optional) ID of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `(Optional) Network address of the backend service.`,
				},
				resource.Attribute{
					Name:        "private_port",
					Description: `(Optional) Port of intranet.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) ID of the VPC. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "dnat_list",
					Description: `Information list of the DNATs.`,
				},
				resource.Attribute{
					Name:        "elastic_ip",
					Description: `Network address of the EIP.`,
				},
				resource.Attribute{
					Name:        "elastic_port",
					Description: `Port of the EIP.`,
				},
				resource.Attribute{
					Name:        "nat_id",
					Description: `ID of the NAT.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Network address of the backend service.`,
				},
				resource.Attribute{
					Name:        "private_port",
					Description: `Port of intranet.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Type of the network protocol. Valid values: ` + "`" + `TCP` + "`" + ` and ` + "`" + `UDP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "dnat_list",
					Description: `Information list of the DNATs.`,
				},
				resource.Attribute{
					Name:        "elastic_ip",
					Description: `Network address of the EIP.`,
				},
				resource.Attribute{
					Name:        "elastic_port",
					Description: `Port of the EIP.`,
				},
				resource.Attribute{
					Name:        "nat_id",
					Description: `ID of the NAT.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Network address of the backend service.`,
				},
				resource.Attribute{
					Name:        "private_port",
					Description: `Port of intranet.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Type of the network protocol. Valid values: ` + "`" + `TCP` + "`" + ` and ` + "`" + `UDP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_eip",
			Category:         "Data Sources",
			ShortDescription: `Provides an available EIP for the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) One or more name/value pairs to filter.`,
				},
				resource.Attribute{
					Name:        "include_arrears",
					Description: `(Optional) Whether the IP is arrears.`,
				},
				resource.Attribute{
					Name:        "include_blocked",
					Description: `(Optional) Whether the IP is blocked. The ` + "`" + `filter` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Key of the filter, valid keys: ` + "`" + `address-id` + "`" + `,` + "`" + `address-name` + "`" + `,` + "`" + `address-ip` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Value of the filter. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `An EIP id indicate the uniqueness of a certain EIP, which can be used for instance binding or network interface binding.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `An public IP address for the EIP.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the EIP, there are several status like ` + "`" + `BIND` + "`" + `, ` + "`" + `UNBIND` + "`" + `, and ` + "`" + `BIND_ENI` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `An EIP id indicate the uniqueness of a certain EIP, which can be used for instance binding or network interface binding.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `An public IP address for the EIP.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the EIP, there are several status like ` + "`" + `BIND` + "`" + `, ` + "`" + `UNBIND` + "`" + `, and ` + "`" + `BIND_ENI` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_eips",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query eip instances.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "eip_id",
					Description: `(Optional) ID of the EIP to be queried.`,
				},
				resource.Attribute{
					Name:        "eip_name",
					Description: `(Optional) Name of the EIP to be queried.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `(Optional) The elastic ip address.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The tags of EIP. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "eip_list",
					Description: `An information list of EIP. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the EIP.`,
				},
				resource.Attribute{
					Name:        "eip_id",
					Description: `ID of the EIP.`,
				},
				resource.Attribute{
					Name:        "eip_name",
					Description: `Name of the EIP.`,
				},
				resource.Attribute{
					Name:        "eip_type",
					Description: `Type of the EIP.`,
				},
				resource.Attribute{
					Name:        "eni_id",
					Description: `The eni id to bind with the EIP.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The instance id to bind with the EIP.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The elastic ip address.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The EIP current status.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the EIP.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "eip_list",
					Description: `An information list of EIP. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the EIP.`,
				},
				resource.Attribute{
					Name:        "eip_id",
					Description: `ID of the EIP.`,
				},
				resource.Attribute{
					Name:        "eip_name",
					Description: `Name of the EIP.`,
				},
				resource.Attribute{
					Name:        "eip_type",
					Description: `Type of the EIP.`,
				},
				resource.Attribute{
					Name:        "eni_id",
					Description: `The eni id to bind with the EIP.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The instance id to bind with the EIP.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The elastic ip address.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The EIP current status.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the EIP.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_elasticsearch_instances",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query elasticsearch instances.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Optional) ID of the instance to be queried.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Optional) Name of the instance to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tag of the instance to be queried. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "instance_list",
					Description: `An information list of elasticsearch instance. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Availability zone.`,
				},
				resource.Attribute{
					Name:        "basic_security_type",
					Description: `Whether to enable X-Pack security authentication in Basic Edition 6.8 and above.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `The charge type of instance.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Instance creation time.`,
				},
				resource.Attribute{
					Name:        "deploy_mode",
					Description: `Cluster deployment mode.`,
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
					Name:        "instance_id",
					Description: `ID of the instance.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `Name of the instance.`,
				},
				resource.Attribute{
					Name:        "kibana_url",
					Description: `Kibana access URL.`,
				},
				resource.Attribute{
					Name:        "license_type",
					Description: `License type.`,
				},
				resource.Attribute{
					Name:        "multi_zone_infos",
					Description: `Details of AZs in multi-AZ deployment mode.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Availability zone.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The id of a VPC subnet.`,
				},
				resource.Attribute{
					Name:        "node_info_list",
					Description: `Node information list, which describe the specification information of various types of nodes in the cluster.`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `Node disk size.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `Node disk type.`,
				},
				resource.Attribute{
					Name:        "encrypt",
					Description: `Decides this disk encrypted or not.`,
				},
				resource.Attribute{
					Name:        "node_num",
					Description: `Number of nodes.`,
				},
				resource.Attribute{
					Name:        "node_type",
					Description: `Node specification.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Node type.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of a VPC subnet.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the instance.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Version of the instance.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of a VPC network.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_list",
					Description: `An information list of elasticsearch instance. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Availability zone.`,
				},
				resource.Attribute{
					Name:        "basic_security_type",
					Description: `Whether to enable X-Pack security authentication in Basic Edition 6.8 and above.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `The charge type of instance.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Instance creation time.`,
				},
				resource.Attribute{
					Name:        "deploy_mode",
					Description: `Cluster deployment mode.`,
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
					Name:        "instance_id",
					Description: `ID of the instance.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `Name of the instance.`,
				},
				resource.Attribute{
					Name:        "kibana_url",
					Description: `Kibana access URL.`,
				},
				resource.Attribute{
					Name:        "license_type",
					Description: `License type.`,
				},
				resource.Attribute{
					Name:        "multi_zone_infos",
					Description: `Details of AZs in multi-AZ deployment mode.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Availability zone.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The id of a VPC subnet.`,
				},
				resource.Attribute{
					Name:        "node_info_list",
					Description: `Node information list, which describe the specification information of various types of nodes in the cluster.`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `Node disk size.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `Node disk type.`,
				},
				resource.Attribute{
					Name:        "encrypt",
					Description: `Decides this disk encrypted or not.`,
				},
				resource.Attribute{
					Name:        "node_num",
					Description: `Number of nodes.`,
				},
				resource.Attribute{
					Name:        "node_type",
					Description: `Node specification.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Node type.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of a VPC subnet.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the instance.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Version of the instance.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of a VPC network.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_enis",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query query ENIs.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the ENI. Conflict with ` + "`" + `ids` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) ID of the ENIs to be queried. Conflict with ` + "`" + `vpc_id` + "`" + `,` + "`" + `subnet_id` + "`" + `,` + "`" + `instance_id` + "`" + `,` + "`" + `security_group` + "`" + `,` + "`" + `name` + "`" + `,` + "`" + `ipv4` + "`" + ` and ` + "`" + `tags` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Optional) ID of the instance which bind the ENI. Conflict with ` + "`" + `ids` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `(Optional) Intranet IP of the ENI. Conflict with ` + "`" + `ids` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the ENI to be queried. Conflict with ` + "`" + `ids` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "security_group",
					Description: `(Optional) A set of security group IDs which bind the ENI. Conflict with ` + "`" + `ids` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) ID of the subnet within this vpc to be queried. Conflict with ` + "`" + `ids` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of the ENI. Conflict with ` + "`" + `ids` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) ID of the vpc to be queried. Conflict with ` + "`" + `ids` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "enis",
					Description: `An information list of ENIs. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the ENI.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the ENI.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the ENI.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the instance which bind the ENI.`,
				},
				resource.Attribute{
					Name:        "ipv4s",
					Description: `A set of intranet IPv4s.`,
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
					Name:        "name",
					Description: `Name of the ENI.`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `Indicates whether the IP is primary.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `A set of security group IDs which bind the ENI.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `States of the ENI.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of the subnet within this vpc.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the ENI.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the vpc.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "enis",
					Description: `An information list of ENIs. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the ENI.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the ENI.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the ENI.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the instance which bind the ENI.`,
				},
				resource.Attribute{
					Name:        "ipv4s",
					Description: `A set of intranet IPv4s.`,
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
					Name:        "name",
					Description: `Name of the ENI.`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `Indicates whether the IP is primary.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `A set of security group IDs which bind the ENI.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `States of the ENI.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of the subnet within this vpc.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the ENI.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the vpc.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_gaap_certificates",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query GAAP certificate.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the certificate to be queried.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the certificate to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type of the certificate to be queried. Valid values: ` + "`" + `BASIC` + "`" + `, ` + "`" + `CLIENT` + "`" + `, ` + "`" + `SERVER` + "`" + `, ` + "`" + `REALSERVER` + "`" + ` and ` + "`" + `PROXY` + "`" + `. ` + "`" + `BASIC` + "`" + ` means basic certificate; ` + "`" + `CLIENT` + "`" + ` means client CA certificate; ` + "`" + `SERVER` + "`" + ` means server SSL certificate; ` + "`" + `REALSERVER` + "`" + ` means realserver CA certificate; ` + "`" + `PROXY` + "`" + ` means proxy SSL certificate. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "certificates",
					Description: `An information list of certificate. Each element contains the following attributes:`,
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
					Name:        "id",
					Description: `ID of the certificate.`,
				},
				resource.Attribute{
					Name:        "issuer_cn",
					Description: `Issuer name of the certificate.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the certificate.`,
				},
				resource.Attribute{
					Name:        "subject_cn",
					Description: `Subject name of the certificate.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the certificate.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "certificates",
					Description: `An information list of certificate. Each element contains the following attributes:`,
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
					Name:        "id",
					Description: `ID of the certificate.`,
				},
				resource.Attribute{
					Name:        "issuer_cn",
					Description: `Issuer name of the certificate.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the certificate.`,
				},
				resource.Attribute{
					Name:        "subject_cn",
					Description: `Subject name of the certificate.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the certificate.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_gaap_domain_error_pages",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query custom GAAP HTTP domain error page info list.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) HTTP domain to be queried.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Required) ID of the layer7 listener to be queried.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) List of the error page info ID to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "error_page_info_list",
					Description: `An information list of error page info detail. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "body",
					Description: `New response body.`,
				},
				resource.Attribute{
					Name:        "clear_headers",
					Description: `Response headers to be removed.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `HTTP domain.`,
				},
				resource.Attribute{
					Name:        "error_codes",
					Description: `Original error codes.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the error page info.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `ID of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "new_error_codes",
					Description: `New error code.`,
				},
				resource.Attribute{
					Name:        "set_headers",
					Description: `Response headers to be set.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "error_page_info_list",
					Description: `An information list of error page info detail. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "body",
					Description: `New response body.`,
				},
				resource.Attribute{
					Name:        "clear_headers",
					Description: `Response headers to be removed.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `HTTP domain.`,
				},
				resource.Attribute{
					Name:        "error_codes",
					Description: `Original error codes.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the error page info.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `ID of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "new_error_codes",
					Description: `New error code.`,
				},
				resource.Attribute{
					Name:        "set_headers",
					Description: `Response headers to be set.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_gaap_http_domains",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query forward domain of layer7 listeners.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) Forward domain of the layer7 listener to be queried.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Required) ID of the layer7 listener to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "domains",
					Description: `An information list of forward domain of the layer7 listeners. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "basic_auth_id",
					Description: `ID of the basic authentication.`,
				},
				resource.Attribute{
					Name:        "basic_auth",
					Description: `Indicates whether basic authentication is enable.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `ID of the server certificate.`,
				},
				resource.Attribute{
					Name:        "client_certificate_id",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "client_certificate_ids",
					Description: `ID list of the client certificate.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Forward domain of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "gaap_auth_id",
					Description: `ID of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "gaap_auth",
					Description: `Indicates whether SSL certificate authentication is enable.`,
				},
				resource.Attribute{
					Name:        "realserver_auth",
					Description: `Indicates whether realserver authentication is enable.`,
				},
				resource.Attribute{
					Name:        "realserver_certificate_domain",
					Description: `CA certificate domain of the realserver.`,
				},
				resource.Attribute{
					Name:        "realserver_certificate_id",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "realserver_certificate_ids",
					Description: `CA certificate ID list of the realserver.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "domains",
					Description: `An information list of forward domain of the layer7 listeners. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "basic_auth_id",
					Description: `ID of the basic authentication.`,
				},
				resource.Attribute{
					Name:        "basic_auth",
					Description: `Indicates whether basic authentication is enable.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `ID of the server certificate.`,
				},
				resource.Attribute{
					Name:        "client_certificate_id",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "client_certificate_ids",
					Description: `ID list of the client certificate.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Forward domain of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "gaap_auth_id",
					Description: `ID of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "gaap_auth",
					Description: `Indicates whether SSL certificate authentication is enable.`,
				},
				resource.Attribute{
					Name:        "realserver_auth",
					Description: `Indicates whether realserver authentication is enable.`,
				},
				resource.Attribute{
					Name:        "realserver_certificate_domain",
					Description: `CA certificate domain of the realserver.`,
				},
				resource.Attribute{
					Name:        "realserver_certificate_id",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "realserver_certificate_ids",
					Description: `CA certificate ID list of the realserver.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_gaap_http_rules",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query forward rule of layer7 listeners.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Required) ID of the layer7 listener to be queried.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional) Forward domain of the layer7 listener to be queried.`,
				},
				resource.Attribute{
					Name:        "forward_host",
					Description: `(Optional) Requested host which is forwarded to the realserver by the listener to be queried.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) Path of the forward rule to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `An information list of forward rule of the layer7 listeners. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "connect_timeout",
					Description: `Timeout of the health check response.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Forward domain of the forward rule.`,
				},
				resource.Attribute{
					Name:        "forward_host",
					Description: `Requested host which is forwarded to the realserver by the listener.`,
				},
				resource.Attribute{
					Name:        "health_check_method",
					Description: `Method of the health check.`,
				},
				resource.Attribute{
					Name:        "health_check_path",
					Description: `Path of health check.`,
				},
				resource.Attribute{
					Name:        "health_check_status_codes",
					Description: `Return code of confirmed normal.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `Indicates whether health check is enable.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the forward rule.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `Interval of the health check.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `ID of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `Path of the forward rule.`,
				},
				resource.Attribute{
					Name:        "realserver_type",
					Description: `Type of the realserver.`,
				},
				resource.Attribute{
					Name:        "realservers",
					Description: `An information list of GAAP realserver. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Domain of the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `IP of the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port of the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Scheduling weight.`,
				},
				resource.Attribute{
					Name:        "scheduler",
					Description: `Scheduling policy of the forward rule.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "rules",
					Description: `An information list of forward rule of the layer7 listeners. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "connect_timeout",
					Description: `Timeout of the health check response.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Forward domain of the forward rule.`,
				},
				resource.Attribute{
					Name:        "forward_host",
					Description: `Requested host which is forwarded to the realserver by the listener.`,
				},
				resource.Attribute{
					Name:        "health_check_method",
					Description: `Method of the health check.`,
				},
				resource.Attribute{
					Name:        "health_check_path",
					Description: `Path of health check.`,
				},
				resource.Attribute{
					Name:        "health_check_status_codes",
					Description: `Return code of confirmed normal.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `Indicates whether health check is enable.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the forward rule.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `Interval of the health check.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `ID of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `Path of the forward rule.`,
				},
				resource.Attribute{
					Name:        "realserver_type",
					Description: `Type of the realserver.`,
				},
				resource.Attribute{
					Name:        "realservers",
					Description: `An information list of GAAP realserver. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Domain of the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `IP of the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port of the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Scheduling weight.`,
				},
				resource.Attribute{
					Name:        "scheduler",
					Description: `Scheduling policy of the forward rule.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_gaap_layer4_listeners",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query gaap layer4 listeners.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Protocol of the layer4 listener to be queried. Valid values: ` + "`" + `TCP` + "`" + ` and ` + "`" + `UDP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Optional) ID of the layer4 listener to be queried.`,
				},
				resource.Attribute{
					Name:        "listener_name",
					Description: `(Optional) Name of the layer4 listener to be queried.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Port of the layer4 listener to be queried.`,
				},
				resource.Attribute{
					Name:        "proxy_id",
					Description: `(Optional) ID of the GAAP proxy to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "listeners",
					Description: `An information list of layer4 listeners. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "connect_timeout",
					Description: `Timeout of the health check response.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the layer4 listener.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `Indicates whether health check is enable.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the layer4 listener.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `Interval of the health check.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the layer4 listener.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port of the layer4 listener.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol of the layer4 listener.`,
				},
				resource.Attribute{
					Name:        "realserver_type",
					Description: `Type of the realserver.`,
				},
				resource.Attribute{
					Name:        "scheduler",
					Description: `Scheduling policy of the layer4 listener.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the layer4 listener.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "listeners",
					Description: `An information list of layer4 listeners. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "connect_timeout",
					Description: `Timeout of the health check response.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the layer4 listener.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `Indicates whether health check is enable.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the layer4 listener.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `Interval of the health check.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the layer4 listener.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port of the layer4 listener.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol of the layer4 listener.`,
				},
				resource.Attribute{
					Name:        "realserver_type",
					Description: `Type of the realserver.`,
				},
				resource.Attribute{
					Name:        "scheduler",
					Description: `Scheduling policy of the layer4 listener.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the layer4 listener.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_gaap_layer7_listeners",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query gaap layer7 listeners.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Protocol of the layer7 listener to be queried. Valid values: ` + "`" + `HTTP` + "`" + ` and ` + "`" + `HTTPS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Optional) ID of the layer7 listener to be queried.`,
				},
				resource.Attribute{
					Name:        "listener_name",
					Description: `(Optional) Name of the layer7 listener to be queried.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Port of the layer7 listener to be queried.`,
				},
				resource.Attribute{
					Name:        "proxy_id",
					Description: `(Optional) ID of the GAAP proxy to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "listeners",
					Description: `An information list of layer7 listeners. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "auth_type",
					Description: `Authentication type of the layer7 listener. ` + "`" + `0` + "`" + ` is one-way authentication and ` + "`" + `1` + "`" + ` is mutual authentication.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `Certificate ID of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "client_certificate_id",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "client_certificate_ids",
					Description: `ID list of the client certificate.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "forward_protocol",
					Description: `Protocol type of the forwarding.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the layer7 listener.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "listeners",
					Description: `An information list of layer7 listeners. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "auth_type",
					Description: `Authentication type of the layer7 listener. ` + "`" + `0` + "`" + ` is one-way authentication and ` + "`" + `1` + "`" + ` is mutual authentication.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `Certificate ID of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "client_certificate_id",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "client_certificate_ids",
					Description: `ID list of the client certificate.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "forward_protocol",
					Description: `Protocol type of the forwarding.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the layer7 listener.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_gaap_proxies",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query gaap proxies.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "access_region",
					Description: `(Optional) Access region of the GAAP proxy to be queried. Conflict with ` + "`" + `ids` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) ID of the GAAP proxy to be queried. Conflict with ` + "`" + `project_id` + "`" + `, ` + "`" + `access_region` + "`" + ` and ` + "`" + `realserver_region` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) Project ID of the GAAP proxy to be queried. Conflict with ` + "`" + `ids` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "realserver_region",
					Description: `(Optional) Region of the GAAP realserver to be queried. Conflict with ` + "`" + `ids` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of the GAAP proxy to be queried. Support up to 5, display the information as long as it matches one. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "proxies",
					Description: `An information list of GAAP proxy. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "access_region",
					Description: `Access region of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `Maximum bandwidth of the GAAP proxy, unit is Mbps.`,
				},
				resource.Attribute{
					Name:        "concurrent",
					Description: `Maximum concurrency of the GAAP proxy, unit is 10k.`,
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
					Name:        "id",
					Description: `ID of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Access domain of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Security policy ID of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of the project within the GAAP proxy, '0' means is default project.`,
				},
				resource.Attribute{
					Name:        "realserver_region",
					Description: `Region of the GAAP realserver.`,
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
					Description: `Supported protocols of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Version of the GAAP proxy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "proxies",
					Description: `An information list of GAAP proxy. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "access_region",
					Description: `Access region of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `Maximum bandwidth of the GAAP proxy, unit is Mbps.`,
				},
				resource.Attribute{
					Name:        "concurrent",
					Description: `Maximum concurrency of the GAAP proxy, unit is 10k.`,
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
					Name:        "id",
					Description: `ID of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Access domain of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Security policy ID of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of the project within the GAAP proxy, '0' means is default project.`,
				},
				resource.Attribute{
					Name:        "realserver_region",
					Description: `Region of the GAAP realserver.`,
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
					Description: `Supported protocols of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Version of the GAAP proxy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_gaap_realservers",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query gaap realservers.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional) Domain of the GAAP realserver to be queried, conflict with ` + "`" + `ip` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Optional) IP of the GAAP realserver to be queried, conflict with ` + "`" + `domain` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the GAAP realserver to be queried, the maximum length is 30.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) ID of the project within the GAAP realserver to be queried, default value is ` + "`" + `-1` + "`" + `, no set means all projects.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of the GAAP proxy to be queried. Support up to 5, display the information as long as it matches one. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "realservers",
					Description: `An information list of GAAP realserver. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Domain of the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `IP of the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of the project within the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the GAAP realserver.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "realservers",
					Description: `An information list of GAAP realserver. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Domain of the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `IP of the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of the project within the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the GAAP realserver.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_gaap_security_policies",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query security policies of GAAP proxy.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) ID of the security policy to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Default policy.`,
				},
				resource.Attribute{
					Name:        "proxy_id",
					Description: `ID of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the security policy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "action",
					Description: `Default policy.`,
				},
				resource.Attribute{
					Name:        "proxy_id",
					Description: `ID of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the security policy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_gaap_security_rules",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query security policy rule.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Required) ID of the security policy to be queried.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Policy of the rule to be queried.`,
				},
				resource.Attribute{
					Name:        "cidr_ip",
					Description: `(Optional) A network address block of the request source to be queried.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the security policy rule to be queried.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Port of the security policy rule to be queried.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) Protocol of the security policy rule to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `(Optional) ID of the security policy rules to be queried. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `An information list of security policy rule. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Policy of the rule.`,
				},
				resource.Attribute{
					Name:        "cidr_ip",
					Description: `A network address block of the request source.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the security policy rule.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the security policy rule.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port of the security policy rule.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol of the security policy rule.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "rules",
					Description: `An information list of security policy rule. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Policy of the rule.`,
				},
				resource.Attribute{
					Name:        "cidr_ip",
					Description: `A network address block of the request source.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the security policy rule.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the security policy rule.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port of the security policy rule.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol of the security policy rule.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_ha_vip_eip_attachments",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of HA VIP EIP attachments`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "havip_id",
					Description: `(Required) ID of the attached HA VIP to be queried.`,
				},
				resource.Attribute{
					Name:        "address_ip",
					Description: `(Optional) Public IP address of EIP to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "ha_vip_eip_attachment_list",
					Description: `A list of HA VIP EIP attachments. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "address_ip",
					Description: `Public IP address of EIP.`,
				},
				resource.Attribute{
					Name:        "havip_id",
					Description: `ID of the attached HA VIP.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ha_vip_eip_attachment_list",
					Description: `A list of HA VIP EIP attachments. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "address_ip",
					Description: `Public IP address of EIP.`,
				},
				resource.Attribute{
					Name:        "havip_id",
					Description: `ID of the attached HA VIP.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_ha_vips",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of HA VIPs.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "address_ip",
					Description: `(Optional) EIP of the HA VIP to be queried.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the HA VIP to be queried.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the HA VIP. The length of character is limited to 1-60.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) Subnet id of the HA VIP to be queried.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) VPC id of the HA VIP to be queried. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "ha_vip_list",
					Description: `Information list of the dedicated HA VIPs.`,
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
					Name:        "id",
					Description: `ID of the HA VIP.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `Instance id that is associated.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the HA VIP.`,
				},
				resource.Attribute{
					Name:        "network_interface_id",
					Description: `Network interface id that is associated.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the HA VIP. Valid values: ` + "`" + `AVAILABLE` + "`" + `, ` + "`" + `UNBIND` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `Subnet id.`,
				},
				resource.Attribute{
					Name:        "vip",
					Description: `Virtual IP address, it must not be occupied and in this VPC network segment. If not set, it will be assigned after resource created automatically.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC id.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ha_vip_list",
					Description: `Information list of the dedicated HA VIPs.`,
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
					Name:        "id",
					Description: `ID of the HA VIP.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `Instance id that is associated.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the HA VIP.`,
				},
				resource.Attribute{
					Name:        "network_interface_id",
					Description: `Network interface id that is associated.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the HA VIP. Valid values: ` + "`" + `AVAILABLE` + "`" + `, ` + "`" + `UNBIND` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `Subnet id.`,
				},
				resource.Attribute{
					Name:        "vip",
					Description: `Virtual IP address, it must not be occupied and in this VPC network segment. If not set, it will be assigned after resource created automatically.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC id.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_image",
			Category:         "Data Sources",
			ShortDescription: `Provides an available image for the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) One or more name/value pairs to filter.`,
				},
				resource.Attribute{
					Name:        "image_name_regex",
					Description: `(Optional) A regex string to apply to the image list returned by TencentCloud.`,
				},
				resource.Attribute{
					Name:        "os_name",
					Description: `(Optional) A string to apply with fuzzy match to the os_name attribute on the image list returned by TencentCloud.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. The ` + "`" + `filter` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Key of the filter, valid keys: ` + "`" + `image-id` + "`" + `, ` + "`" + `image-type` + "`" + `, ` + "`" + `image-name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Values of the filter. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `An image id indicate the uniqueness of a certain image, which can be used for instance creation or resetting.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `Name of this image.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "image_id",
					Description: `An image id indicate the uniqueness of a certain image, which can be used for instance creation or resetting.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `Name of this image.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_images",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query images.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "image_id",
					Description: `(Optional) ID of the image to be queried.`,
				},
				resource.Attribute{
					Name:        "image_name_regex",
					Description: `(Optional) A regex string to apply to the image list returned by TencentCloud, conflict with 'os_name'.`,
				},
				resource.Attribute{
					Name:        "image_type",
					Description: `(Optional) A list of the image type to be queried. Valid values: 'PUBLIC_IMAGE', 'PRIVATE_IMAGE', 'SHARED_IMAGE', 'MARKET_IMAGE'.`,
				},
				resource.Attribute{
					Name:        "os_name",
					Description: `(Optional) A string to apply with fuzzy match to the os_name attribute on the image list returned by TencentCloud, conflict with 'image_name_regex'.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "images",
					Description: `An information list of image. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "architecture",
					Description: `Architecture of the image.`,
				},
				resource.Attribute{
					Name:        "created_time",
					Description: `Created time of the image.`,
				},
				resource.Attribute{
					Name:        "image_creator",
					Description: `Image creator of the image.`,
				},
				resource.Attribute{
					Name:        "image_description",
					Description: `Description of the image.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `ID of the image.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `Name of the image.`,
				},
				resource.Attribute{
					Name:        "image_size",
					Description: `Size of the image.`,
				},
				resource.Attribute{
					Name:        "image_source",
					Description: `Image source of the image.`,
				},
				resource.Attribute{
					Name:        "image_state",
					Description: `State of the image.`,
				},
				resource.Attribute{
					Name:        "image_type",
					Description: `Type of the image.`,
				},
				resource.Attribute{
					Name:        "os_name",
					Description: `OS name of the image.`,
				},
				resource.Attribute{
					Name:        "platform",
					Description: `Platform of the image.`,
				},
				resource.Attribute{
					Name:        "snapshots",
					Description: `List of snapshot details.`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `Size of the cloud disk used to create the snapshot; unit: GB.`,
				},
				resource.Attribute{
					Name:        "disk_usage",
					Description: `Type of the cloud disk used to create the snapshot.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `Snapshot ID.`,
				},
				resource.Attribute{
					Name:        "snapshot_name",
					Description: `Snapshot name, the user-defined snapshot alias.`,
				},
				resource.Attribute{
					Name:        "support_cloud_init",
					Description: `Whether support cloud-init.`,
				},
				resource.Attribute{
					Name:        "sync_percent",
					Description: `Sync percent of the image.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "images",
					Description: `An information list of image. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "architecture",
					Description: `Architecture of the image.`,
				},
				resource.Attribute{
					Name:        "created_time",
					Description: `Created time of the image.`,
				},
				resource.Attribute{
					Name:        "image_creator",
					Description: `Image creator of the image.`,
				},
				resource.Attribute{
					Name:        "image_description",
					Description: `Description of the image.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `ID of the image.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `Name of the image.`,
				},
				resource.Attribute{
					Name:        "image_size",
					Description: `Size of the image.`,
				},
				resource.Attribute{
					Name:        "image_source",
					Description: `Image source of the image.`,
				},
				resource.Attribute{
					Name:        "image_state",
					Description: `State of the image.`,
				},
				resource.Attribute{
					Name:        "image_type",
					Description: `Type of the image.`,
				},
				resource.Attribute{
					Name:        "os_name",
					Description: `OS name of the image.`,
				},
				resource.Attribute{
					Name:        "platform",
					Description: `Platform of the image.`,
				},
				resource.Attribute{
					Name:        "snapshots",
					Description: `List of snapshot details.`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `Size of the cloud disk used to create the snapshot; unit: GB.`,
				},
				resource.Attribute{
					Name:        "disk_usage",
					Description: `Type of the cloud disk used to create the snapshot.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `Snapshot ID.`,
				},
				resource.Attribute{
					Name:        "snapshot_name",
					Description: `Snapshot name, the user-defined snapshot alias.`,
				},
				resource.Attribute{
					Name:        "support_cloud_init",
					Description: `Whether support cloud-init.`,
				},
				resource.Attribute{
					Name:        "sync_percent",
					Description: `Sync percent of the image.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_instance_types",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query instances type.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) The available zone that the CVM instance locates at. This field is conflict with ` + "`" + `filter` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cpu_core_count",
					Description: `(Optional) The number of CPU cores of the instance.`,
				},
				resource.Attribute{
					Name:        "exclude_sold_out",
					Description: `(Optional) Indicate to filter instances types that is sold out or not, default is false.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) One or more name/value pairs to filter. This field is conflict with ` + "`" + `availability_zone` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "gpu_core_count",
					Description: `(Optional) The number of GPU cores of the instance.`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `(Optional) Instance memory capacity, unit in GB.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. The ` + "`" + `filter` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The filter name. Valid values: ` + "`" + `zone` + "`" + ` and ` + "`" + `instance-family` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) The filter values. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "instance_types",
					Description: `An information list of cvm instance. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The available zone that the CVM instance locates at.`,
				},
				resource.Attribute{
					Name:        "cpu_core_count",
					Description: `The number of CPU cores of the instance.`,
				},
				resource.Attribute{
					Name:        "family",
					Description: `Type series of the instance.`,
				},
				resource.Attribute{
					Name:        "gpu_core_count",
					Description: `The number of GPU cores of the instance.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `Type of the instance.`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `Instance memory capacity, unit in GB.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_types",
					Description: `An information list of cvm instance. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The available zone that the CVM instance locates at.`,
				},
				resource.Attribute{
					Name:        "cpu_core_count",
					Description: `The number of CPU cores of the instance.`,
				},
				resource.Attribute{
					Name:        "family",
					Description: `Type series of the instance.`,
				},
				resource.Attribute{
					Name:        "gpu_core_count",
					Description: `The number of GPU cores of the instance.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `Type of the instance.`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `Instance memory capacity, unit in GB.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_instances",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query cvm instances.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) The available zone that the CVM instance locates at.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Optional) ID of the instances to be queried.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Optional) Name of the instances to be queried.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) The project CVM belongs to.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) ID of a vpc subnetwork.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of the instance.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) ID of the vpc to be queried. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "instance_list",
					Description: `An information list of cvm instance. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "allocate_public_ip",
					Description: `Indicates whether public ip is assigned.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The available zone that the CVM instance locates at.`,
				},
				resource.Attribute{
					Name:        "cam_role_name",
					Description: `CAM role name authorized to access.`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `The number of CPU cores of the instance.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the instance.`,
				},
				resource.Attribute{
					Name:        "data_disks",
					Description: `An information list of data disk. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "data_disk_id",
					Description: `Image ID of the data disk.`,
				},
				resource.Attribute{
					Name:        "data_disk_size",
					Description: `Size of the data disk.`,
				},
				resource.Attribute{
					Name:        "data_disk_type",
					Description: `Type of the data disk.`,
				},
				resource.Attribute{
					Name:        "delete_with_instance",
					Description: `Indicates whether the data disk is destroyed with the instance.`,
				},
				resource.Attribute{
					Name:        "expired_time",
					Description: `Expired time of the instance.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `ID of the image.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type_prepaid_renew_flag",
					Description: `The way that CVM instance will be renew automatically or not when it reach the end of the prepaid tenancy.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `The charge type of the instance.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the instances.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `Name of the instances.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `Type of the instance.`,
				},
				resource.Attribute{
					Name:        "internet_charge_type",
					Description: `The charge type of the instance.`,
				},
				resource.Attribute{
					Name:        "internet_max_bandwidth_out",
					Description: `Public network maximum output bandwidth of the instance.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `Instance memory capacity, unit in GB.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Private IP of the instance.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The project CVM belongs to.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP of the instance.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `Security groups of the instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the instance.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of a vpc subnetwork.`,
				},
				resource.Attribute{
					Name:        "system_disk_id",
					Description: `Image ID of the system disk.`,
				},
				resource.Attribute{
					Name:        "system_disk_size",
					Description: `Size of the system disk.`,
				},
				resource.Attribute{
					Name:        "system_disk_type",
					Description: `Type of the system disk.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the instance.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the vpc.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_list",
					Description: `An information list of cvm instance. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "allocate_public_ip",
					Description: `Indicates whether public ip is assigned.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The available zone that the CVM instance locates at.`,
				},
				resource.Attribute{
					Name:        "cam_role_name",
					Description: `CAM role name authorized to access.`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `The number of CPU cores of the instance.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the instance.`,
				},
				resource.Attribute{
					Name:        "data_disks",
					Description: `An information list of data disk. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "data_disk_id",
					Description: `Image ID of the data disk.`,
				},
				resource.Attribute{
					Name:        "data_disk_size",
					Description: `Size of the data disk.`,
				},
				resource.Attribute{
					Name:        "data_disk_type",
					Description: `Type of the data disk.`,
				},
				resource.Attribute{
					Name:        "delete_with_instance",
					Description: `Indicates whether the data disk is destroyed with the instance.`,
				},
				resource.Attribute{
					Name:        "expired_time",
					Description: `Expired time of the instance.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `ID of the image.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type_prepaid_renew_flag",
					Description: `The way that CVM instance will be renew automatically or not when it reach the end of the prepaid tenancy.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `The charge type of the instance.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the instances.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `Name of the instances.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `Type of the instance.`,
				},
				resource.Attribute{
					Name:        "internet_charge_type",
					Description: `The charge type of the instance.`,
				},
				resource.Attribute{
					Name:        "internet_max_bandwidth_out",
					Description: `Public network maximum output bandwidth of the instance.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `Instance memory capacity, unit in GB.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Private IP of the instance.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The project CVM belongs to.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP of the instance.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `Security groups of the instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the instance.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of a vpc subnetwork.`,
				},
				resource.Attribute{
					Name:        "system_disk_id",
					Description: `Image ID of the system disk.`,
				},
				resource.Attribute{
					Name:        "system_disk_size",
					Description: `Size of the system disk.`,
				},
				resource.Attribute{
					Name:        "system_disk_type",
					Description: `Type of the system disk.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the instance.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the vpc.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_key_pairs",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query key pairs.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key_id",
					Description: `(Optional) ID of the key pair to be queried.`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `(Optional) Name of the key pair to be queried. Support regular expression search, only ` + "`" + `^` + "`" + ` and ` + "`" + `$` + "`" + ` are supported.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) Project ID of the key pair to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "key_pair_list",
					Description: `An information list of key pair. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the key pair.`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `ID of the key pair.`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `Name of the key pair.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Project ID of the key pair.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `public key of the key pair.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "key_pair_list",
					Description: `An information list of key pair. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the key pair.`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `ID of the key pair.`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `Name of the key pair.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Project ID of the key pair.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `public key of the key pair.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_kms_keys",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of KMS key`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key_state",
					Description: `(Optional) Filter by state of CMK. ` + "`" + `0` + "`" + ` - all CMKs are queried, ` + "`" + `1` + "`" + ` - only Enabled CMKs are queried, ` + "`" + `2` + "`" + ` - only Disabled CMKs are queried, ` + "`" + `3` + "`" + ` - only PendingDelete CMKs are queried, ` + "`" + `4` + "`" + ` - only PendingImport CMKs are queried, ` + "`" + `5` + "`" + ` - only Archived CMKs are queried.`,
				},
				resource.Attribute{
					Name:        "key_usage",
					Description: `(Optional) Filter by usage of CMK. Available values include ` + "`" + `ALL` + "`" + `, ` + "`" + `ENCRYPT_DECRYPT` + "`" + `, ` + "`" + `ASYMMETRIC_DECRYPT_RSA_2048` + "`" + `, ` + "`" + `ASYMMETRIC_DECRYPT_SM2` + "`" + `, ` + "`" + `ASYMMETRIC_SIGN_VERIFY_SM2` + "`" + `, ` + "`" + `ASYMMETRIC_SIGN_VERIFY_RSA_2048` + "`" + `, ` + "`" + `ASYMMETRIC_SIGN_VERIFY_ECC` + "`" + `. Default value is ` + "`" + `ENCRYPT_DECRYPT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "order_type",
					Description: `(Optional) Order to sort the CMK create time. ` + "`" + `0` + "`" + ` - desc, ` + "`" + `1` + "`" + ` - asc. Default value is ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "origin",
					Description: `(Optional) Filter by origin of CMK. ` + "`" + `TENCENT_KMS` + "`" + ` - CMK created by KMS, ` + "`" + `EXTERNAL` + "`" + ` - CMK imported by user, ` + "`" + `ALL` + "`" + ` - all CMKs. Default value is ` + "`" + `ALL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Optional) Filter by role of the CMK creator. ` + "`" + `0` + "`" + ` - created by user, ` + "`" + `1` + "`" + ` - created by cloud product. Default value is ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "search_key_alias",
					Description: `(Optional) Words used to match the results, and the words can be: key_id and alias.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags to filter CMK. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "key_list",
					Description: `A list of KMS keys.`,
				},
				resource.Attribute{
					Name:        "alias",
					Description: `Name of CMK.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of CMK.`,
				},
				resource.Attribute{
					Name:        "creator_uin",
					Description: `Uin of CMK Creator.`,
				},
				resource.Attribute{
					Name:        "deletion_date",
					Description: `Delete time of CMK.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of CMK.`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `ID of CMK.`,
				},
				resource.Attribute{
					Name:        "key_rotation_enabled",
					Description: `Specify whether to enable key rotation.`,
				},
				resource.Attribute{
					Name:        "key_state",
					Description: `State of CMK.`,
				},
				resource.Attribute{
					Name:        "key_usage",
					Description: `Usage of CMK.`,
				},
				resource.Attribute{
					Name:        "next_rotate_time",
					Description: `Next rotate time of CMK when key_rotation_enabled is true.`,
				},
				resource.Attribute{
					Name:        "origin",
					Description: `Origin of CMK. ` + "`" + `TENCENT_KMS` + "`" + ` - CMK created by KMS, ` + "`" + `EXTERNAL` + "`" + ` - CMK imported by user.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Creator of CMK.`,
				},
				resource.Attribute{
					Name:        "valid_to",
					Description: `Valid when origin is ` + "`" + `EXTERNAL` + "`" + `, it means the effective date of the key material.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "key_list",
					Description: `A list of KMS keys.`,
				},
				resource.Attribute{
					Name:        "alias",
					Description: `Name of CMK.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of CMK.`,
				},
				resource.Attribute{
					Name:        "creator_uin",
					Description: `Uin of CMK Creator.`,
				},
				resource.Attribute{
					Name:        "deletion_date",
					Description: `Delete time of CMK.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of CMK.`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `ID of CMK.`,
				},
				resource.Attribute{
					Name:        "key_rotation_enabled",
					Description: `Specify whether to enable key rotation.`,
				},
				resource.Attribute{
					Name:        "key_state",
					Description: `State of CMK.`,
				},
				resource.Attribute{
					Name:        "key_usage",
					Description: `Usage of CMK.`,
				},
				resource.Attribute{
					Name:        "next_rotate_time",
					Description: `Next rotate time of CMK when key_rotation_enabled is true.`,
				},
				resource.Attribute{
					Name:        "origin",
					Description: `Origin of CMK. ` + "`" + `TENCENT_KMS` + "`" + ` - CMK created by KMS, ` + "`" + `EXTERNAL` + "`" + ` - CMK imported by user.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Creator of CMK.`,
				},
				resource.Attribute{
					Name:        "valid_to",
					Description: `Valid when origin is ` + "`" + `EXTERNAL` + "`" + `, it means the effective date of the key material.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_kubernetes_clusters",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of kubernetes clusters.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Optional) ID of the cluster. Conflict with cluster_name, can not be set at the same time.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Optional) Name of the cluster. Conflict with cluster_id, can not be set at the same time.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of the cluster. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `An information list of kubernetes clusters. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "certification_authority",
					Description: `The certificate used for access.`,
				},
				resource.Attribute{
					Name:        "claim_expired_seconds",
					Description: `The expired seconds to recycle ENI.`,
				},
				resource.Attribute{
					Name:        "cluster_as_enabled",
					Description: `Indicates whether to enable cluster node auto scaler.`,
				},
				resource.Attribute{
					Name:        "cluster_cidr",
					Description: `A network address block of the cluster. Different from vpc cidr and cidr of other clusters within this VPC.`,
				},
				resource.Attribute{
					Name:        "cluster_deploy_type",
					Description: `Deployment type of the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_desc",
					Description: `Description of the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_external_endpoint",
					Description: `External network address to access.`,
				},
				resource.Attribute{
					Name:        "cluster_extra_args",
					Description: `Customized parameters for master component.`,
				},
				resource.Attribute{
					Name:        "kube_apiserver",
					Description: `The customized parameters for kube-apiserver.`,
				},
				resource.Attribute{
					Name:        "kube_controller_manager",
					Description: `The customized parameters for kube-controller-manager.`,
				},
				resource.Attribute{
					Name:        "kube_scheduler",
					Description: `The customized parameters for kube-scheduler.`,
				},
				resource.Attribute{
					Name:        "cluster_ipvs",
					Description: `Indicates whether ipvs is enabled.`,
				},
				resource.Attribute{
					Name:        "cluster_max_pod_num",
					Description: `The maximum number of Pods per node in the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_max_service_num",
					Description: `The maximum number of services in the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `Name of the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_node_num",
					Description: `Number of nodes in the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_os",
					Description: `Operating system of the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_version",
					Description: `Version of the cluster.`,
				},
				resource.Attribute{
					Name:        "container_runtime",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "deletion_protection",
					Description: `Indicates whether cluster deletion protection is enabled.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Domain name for access.`,
				},
				resource.Attribute{
					Name:        "eni_subnet_ids",
					Description: `Subnet IDs for cluster with VPC-CNI network mode.`,
				},
				resource.Attribute{
					Name:        "ignore_cluster_cidr_conflict",
					Description: `Indicates whether to ignore the cluster cidr conflict error.`,
				},
				resource.Attribute{
					Name:        "is_non_static_ip_mode",
					Description: `Indicates whether static ip mode is enabled.`,
				},
				resource.Attribute{
					Name:        "kube_config",
					Description: `kubernetes config.`,
				},
				resource.Attribute{
					Name:        "kube_proxy_mode",
					Description: `Cluster kube-proxy mode.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `Cluster network type.`,
				},
				resource.Attribute{
					Name:        "node_name_type",
					Description: `Node name type of cluster.`,
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
					Name:        "project_id",
					Description: `Project ID of the cluster.`,
				},
				resource.Attribute{
					Name:        "security_policy",
					Description: `Access policy.`,
				},
				resource.Attribute{
					Name:        "service_cidr",
					Description: `The network address block of the cluster.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the cluster.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `User name of account.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `Vpc ID of the cluster.`,
				},
				resource.Attribute{
					Name:        "worker_instances_list",
					Description: `An information list of cvm within the WORKER clusters. Each element contains the following attributes.`,
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
					Name:        "list",
					Description: `An information list of kubernetes clusters. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "certification_authority",
					Description: `The certificate used for access.`,
				},
				resource.Attribute{
					Name:        "claim_expired_seconds",
					Description: `The expired seconds to recycle ENI.`,
				},
				resource.Attribute{
					Name:        "cluster_as_enabled",
					Description: `Indicates whether to enable cluster node auto scaler.`,
				},
				resource.Attribute{
					Name:        "cluster_cidr",
					Description: `A network address block of the cluster. Different from vpc cidr and cidr of other clusters within this VPC.`,
				},
				resource.Attribute{
					Name:        "cluster_deploy_type",
					Description: `Deployment type of the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_desc",
					Description: `Description of the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_external_endpoint",
					Description: `External network address to access.`,
				},
				resource.Attribute{
					Name:        "cluster_extra_args",
					Description: `Customized parameters for master component.`,
				},
				resource.Attribute{
					Name:        "kube_apiserver",
					Description: `The customized parameters for kube-apiserver.`,
				},
				resource.Attribute{
					Name:        "kube_controller_manager",
					Description: `The customized parameters for kube-controller-manager.`,
				},
				resource.Attribute{
					Name:        "kube_scheduler",
					Description: `The customized parameters for kube-scheduler.`,
				},
				resource.Attribute{
					Name:        "cluster_ipvs",
					Description: `Indicates whether ipvs is enabled.`,
				},
				resource.Attribute{
					Name:        "cluster_max_pod_num",
					Description: `The maximum number of Pods per node in the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_max_service_num",
					Description: `The maximum number of services in the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `Name of the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_node_num",
					Description: `Number of nodes in the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_os",
					Description: `Operating system of the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_version",
					Description: `Version of the cluster.`,
				},
				resource.Attribute{
					Name:        "container_runtime",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "deletion_protection",
					Description: `Indicates whether cluster deletion protection is enabled.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Domain name for access.`,
				},
				resource.Attribute{
					Name:        "eni_subnet_ids",
					Description: `Subnet IDs for cluster with VPC-CNI network mode.`,
				},
				resource.Attribute{
					Name:        "ignore_cluster_cidr_conflict",
					Description: `Indicates whether to ignore the cluster cidr conflict error.`,
				},
				resource.Attribute{
					Name:        "is_non_static_ip_mode",
					Description: `Indicates whether static ip mode is enabled.`,
				},
				resource.Attribute{
					Name:        "kube_config",
					Description: `kubernetes config.`,
				},
				resource.Attribute{
					Name:        "kube_proxy_mode",
					Description: `Cluster kube-proxy mode.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `Cluster network type.`,
				},
				resource.Attribute{
					Name:        "node_name_type",
					Description: `Node name type of cluster.`,
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
					Name:        "project_id",
					Description: `Project ID of the cluster.`,
				},
				resource.Attribute{
					Name:        "security_policy",
					Description: `Access policy.`,
				},
				resource.Attribute{
					Name:        "service_cidr",
					Description: `The network address block of the cluster.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the cluster.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `User name of account.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `Vpc ID of the cluster.`,
				},
				resource.Attribute{
					Name:        "worker_instances_list",
					Description: `An information list of cvm within the WORKER clusters. Each element contains the following attributes.`,
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
			Type:             "tencentcloud_mongodb_instances",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of Mongodb instances.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_type",
					Description: `(Optional) Type of Mongodb cluster, and available values include replica set cluster(expressed with ` + "`" + `REPLSET` + "`" + `), sharding cluster(expressed with ` + "`" + `SHARD` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Optional) ID of the Mongodb instance to be queried.`,
				},
				resource.Attribute{
					Name:        "instance_name_prefix",
					Description: `(Optional) Name prefix of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to store results.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of the Mongodb instance to be queried. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "instance_list",
					Description: `A list of instances. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "auto_renew_flag",
					Description: `Auto renew flag.`,
				},
				resource.Attribute{
					Name:        "available_zone",
					Description: `The available zone of the Mongodb.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `The charge type of instance.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `Type of Mongodb cluster.`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `Number of cpu's core.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `Version of the Mongodb engine.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `Name of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "machine_type",
					Description: `Type of Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `Memory size.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of the project which the instance belongs.`,
				},
				resource.Attribute{
					Name:        "shard_quantity",
					Description: `Number of sharding.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the Mongodb, and available values include pending initialization(expressed with 0), processing(expressed with 1), running(expressed with 2) and expired(expressed with -2).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of the subnet.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "vip",
					Description: `IP of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "volume",
					Description: `Disk size.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC.`,
				},
				resource.Attribute{
					Name:        "vport",
					Description: `IP port of the Mongodb instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_list",
					Description: `A list of instances. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "auto_renew_flag",
					Description: `Auto renew flag.`,
				},
				resource.Attribute{
					Name:        "available_zone",
					Description: `The available zone of the Mongodb.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `The charge type of instance.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `Type of Mongodb cluster.`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `Number of cpu's core.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `Version of the Mongodb engine.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `Name of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "machine_type",
					Description: `Type of Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `Memory size.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of the project which the instance belongs.`,
				},
				resource.Attribute{
					Name:        "shard_quantity",
					Description: `Number of sharding.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the Mongodb, and available values include pending initialization(expressed with 0), processing(expressed with 1), running(expressed with 2) and expired(expressed with -2).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of the subnet.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "vip",
					Description: `IP of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "volume",
					Description: `Disk size.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC.`,
				},
				resource.Attribute{
					Name:        "vport",
					Description: `IP port of the Mongodb instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_mongodb_zone_config",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query the available mongodb specifications for different zone.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "available_zone",
					Description: `(Optional) The available zone of the Mongodb.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to store results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list of zone config. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "available_zone",
					Description: `The available zone of the Mongodb.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `Type of Mongodb cluster.`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `Number of cpu's core.`,
				},
				resource.Attribute{
					Name:        "default_storage",
					Description: `Default disk size.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `Version of the Mongodb version.`,
				},
				resource.Attribute{
					Name:        "machine_type",
					Description: `Type of Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "max_storage",
					Description: `Maximum size of the disk.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `Memory size.`,
				},
				resource.Attribute{
					Name:        "min_storage",
					Description: `Minimum sie of the disk.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list of zone config. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "available_zone",
					Description: `The available zone of the Mongodb.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `Type of Mongodb cluster.`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `Number of cpu's core.`,
				},
				resource.Attribute{
					Name:        "default_storage",
					Description: `Default disk size.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `Version of the Mongodb version.`,
				},
				resource.Attribute{
					Name:        "machine_type",
					Description: `Type of Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "max_storage",
					Description: `Maximum size of the disk.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `Memory size.`,
				},
				resource.Attribute{
					Name:        "min_storage",
					Description: `Minimum sie of the disk.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_monitor_binding_objects",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query policy group binding objects.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required) Policy group ID for query.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to store results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list objects. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "dimensions_json",
					Description: `Represents a collection of dimensions of an object instance, json format.`,
				},
				resource.Attribute{
					Name:        "is_shielded",
					Description: `Whether the object is shielded or not, ` + "`" + `0` + "`" + ` means unshielded and ` + "`" + `1` + "`" + ` means shielded.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region where the object is located.`,
				},
				resource.Attribute{
					Name:        "unique_id",
					Description: `Object unique ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list objects. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "dimensions_json",
					Description: `Represents a collection of dimensions of an object instance, json format.`,
				},
				resource.Attribute{
					Name:        "is_shielded",
					Description: `Whether the object is shielded or not, ` + "`" + `0` + "`" + ` means unshielded and ` + "`" + `1` + "`" + ` means shielded.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region where the object is located.`,
				},
				resource.Attribute{
					Name:        "unique_id",
					Description: `Object unique ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_monitor_data",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query monitor data. for complex queries, use (https://github.com/tencentyun/tencentcloud-exporter)`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dimensions",
					Description: `(Required) Dimensional composition of instance objects.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `(Required) End time for this query, eg:` + "`" + `2018-09-22T20:00:00+08:00` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "metric_name",
					Description: `(Required) Metric name, please refer to the documentation of monitor interface of each product.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Required) Namespace of each cloud product in monitor system, refer to ` + "`" + `data.tencentcloud_monitor_product_namespace` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `(Required) Start time for this query, eg:` + "`" + `2018-09-22T19:51:23+08:00` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) Statistical period.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to store results. The ` + "`" + `dimensions` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Instance dimension name, eg: ` + "`" + `InstanceId` + "`" + ` for cvm.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Instance dimension value, eg: ` + "`" + `ins-j0hk02zo` + "`" + ` for cvm. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list data point. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "timestamp",
					Description: `Statistical timestamp.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Statistical value.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list data point. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "timestamp",
					Description: `Statistical timestamp.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Statistical value.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_monitor_policy_conditions",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query monitor policy conditions(There is a lot of data and it is recommended to output to a file)`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the policy name, support partial matching, eg:` + "`" + `Cloud Virtual Machine` + "`" + `,` + "`" + `Virtual` + "`" + `,` + "`" + `Cloud Load Banlancer-Private CLB Listener` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to store results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list policy condition. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "event_metrics",
					Description: `A list of event condition metrics. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "event_id",
					Description: `The ID of this event metric.`,
				},
				resource.Attribute{
					Name:        "event_show_name",
					Description: `The name of this event metric.`,
				},
				resource.Attribute{
					Name:        "need_recovered",
					Description: `Whether to recover.`,
				},
				resource.Attribute{
					Name:        "is_support_multi_region",
					Description: `Whether to support multi region.`,
				},
				resource.Attribute{
					Name:        "metrics",
					Description: `A list of event condition metrics. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "calc_type_keys",
					Description: `Calculate type of this metric.`,
				},
				resource.Attribute{
					Name:        "calc_type_need",
					Description: `Whether ` + "`" + `calc_type` + "`" + ` required in the configuration.`,
				},
				resource.Attribute{
					Name:        "calc_value_default",
					Description: `The default calculate value of this metric.`,
				},
				resource.Attribute{
					Name:        "calc_value_fixed",
					Description: `The fixed calculate value of this metric.`,
				},
				resource.Attribute{
					Name:        "calc_value_max",
					Description: `The max calculate value of this metric.`,
				},
				resource.Attribute{
					Name:        "calc_value_min",
					Description: `The min calculate value of this metric.`,
				},
				resource.Attribute{
					Name:        "calc_value_need",
					Description: `Whether ` + "`" + `calc_value` + "`" + ` required in the configuration.`,
				},
				resource.Attribute{
					Name:        "continue_time_default",
					Description: `The default continue time(seconds) config for this metric.`,
				},
				resource.Attribute{
					Name:        "continue_time_keys",
					Description: `The continue time(seconds) keys for this metric.`,
				},
				resource.Attribute{
					Name:        "continue_time_need",
					Description: `Whether ` + "`" + `continue_time` + "`" + ` required in the configuration.`,
				},
				resource.Attribute{
					Name:        "metric_id",
					Description: `The ID of this metric.`,
				},
				resource.Attribute{
					Name:        "metric_show_name",
					Description: `The name of this metric.`,
				},
				resource.Attribute{
					Name:        "metric_unit",
					Description: `The unit of this metric.`,
				},
				resource.Attribute{
					Name:        "period_default",
					Description: `The default data time(seconds) config for this metric.`,
				},
				resource.Attribute{
					Name:        "period_keys",
					Description: `The data time(seconds) keys for this metric.`,
				},
				resource.Attribute{
					Name:        "period_need",
					Description: `Whether ` + "`" + `period` + "`" + ` required in the configuration.`,
				},
				resource.Attribute{
					Name:        "period_num_default",
					Description: `The default period number config for this metric.`,
				},
				resource.Attribute{
					Name:        "period_num_keys",
					Description: `The period number keys for this metric.`,
				},
				resource.Attribute{
					Name:        "period_num_need",
					Description: `Whether ` + "`" + `period_num` + "`" + ` required in the configuration.`,
				},
				resource.Attribute{
					Name:        "stat_type_p10",
					Description: `Data aggregation mode, cycle of 10 seconds.`,
				},
				resource.Attribute{
					Name:        "stat_type_p1800",
					Description: `Data aggregation mode, cycle of 1800 seconds.`,
				},
				resource.Attribute{
					Name:        "stat_type_p300",
					Description: `Data aggregation mode, cycle of 300 seconds.`,
				},
				resource.Attribute{
					Name:        "stat_type_p3600",
					Description: `Data aggregation mode, cycle of 3600 seconds.`,
				},
				resource.Attribute{
					Name:        "stat_type_p5",
					Description: `Data aggregation mode, cycle of 5 seconds.`,
				},
				resource.Attribute{
					Name:        "stat_type_p600",
					Description: `Data aggregation mode, cycle of 600 seconds.`,
				},
				resource.Attribute{
					Name:        "stat_type_p60",
					Description: `Data aggregation mode, cycle of 60 seconds.`,
				},
				resource.Attribute{
					Name:        "stat_type_p86400",
					Description: `Data aggregation mode, cycle of 86400 seconds.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of this policy name.`,
				},
				resource.Attribute{
					Name:        "policy_view_name",
					Description: `Policy view name, eg:` + "`" + `cvm_device` + "`" + `,` + "`" + `BANDWIDTHPACKAGE` + "`" + `, refer to ` + "`" + `data.tencentcloud_monitor_policy_conditions(policy_view_name)` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "support_regions",
					Description: `Support regions of this policy view.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list policy condition. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "event_metrics",
					Description: `A list of event condition metrics. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "event_id",
					Description: `The ID of this event metric.`,
				},
				resource.Attribute{
					Name:        "event_show_name",
					Description: `The name of this event metric.`,
				},
				resource.Attribute{
					Name:        "need_recovered",
					Description: `Whether to recover.`,
				},
				resource.Attribute{
					Name:        "is_support_multi_region",
					Description: `Whether to support multi region.`,
				},
				resource.Attribute{
					Name:        "metrics",
					Description: `A list of event condition metrics. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "calc_type_keys",
					Description: `Calculate type of this metric.`,
				},
				resource.Attribute{
					Name:        "calc_type_need",
					Description: `Whether ` + "`" + `calc_type` + "`" + ` required in the configuration.`,
				},
				resource.Attribute{
					Name:        "calc_value_default",
					Description: `The default calculate value of this metric.`,
				},
				resource.Attribute{
					Name:        "calc_value_fixed",
					Description: `The fixed calculate value of this metric.`,
				},
				resource.Attribute{
					Name:        "calc_value_max",
					Description: `The max calculate value of this metric.`,
				},
				resource.Attribute{
					Name:        "calc_value_min",
					Description: `The min calculate value of this metric.`,
				},
				resource.Attribute{
					Name:        "calc_value_need",
					Description: `Whether ` + "`" + `calc_value` + "`" + ` required in the configuration.`,
				},
				resource.Attribute{
					Name:        "continue_time_default",
					Description: `The default continue time(seconds) config for this metric.`,
				},
				resource.Attribute{
					Name:        "continue_time_keys",
					Description: `The continue time(seconds) keys for this metric.`,
				},
				resource.Attribute{
					Name:        "continue_time_need",
					Description: `Whether ` + "`" + `continue_time` + "`" + ` required in the configuration.`,
				},
				resource.Attribute{
					Name:        "metric_id",
					Description: `The ID of this metric.`,
				},
				resource.Attribute{
					Name:        "metric_show_name",
					Description: `The name of this metric.`,
				},
				resource.Attribute{
					Name:        "metric_unit",
					Description: `The unit of this metric.`,
				},
				resource.Attribute{
					Name:        "period_default",
					Description: `The default data time(seconds) config for this metric.`,
				},
				resource.Attribute{
					Name:        "period_keys",
					Description: `The data time(seconds) keys for this metric.`,
				},
				resource.Attribute{
					Name:        "period_need",
					Description: `Whether ` + "`" + `period` + "`" + ` required in the configuration.`,
				},
				resource.Attribute{
					Name:        "period_num_default",
					Description: `The default period number config for this metric.`,
				},
				resource.Attribute{
					Name:        "period_num_keys",
					Description: `The period number keys for this metric.`,
				},
				resource.Attribute{
					Name:        "period_num_need",
					Description: `Whether ` + "`" + `period_num` + "`" + ` required in the configuration.`,
				},
				resource.Attribute{
					Name:        "stat_type_p10",
					Description: `Data aggregation mode, cycle of 10 seconds.`,
				},
				resource.Attribute{
					Name:        "stat_type_p1800",
					Description: `Data aggregation mode, cycle of 1800 seconds.`,
				},
				resource.Attribute{
					Name:        "stat_type_p300",
					Description: `Data aggregation mode, cycle of 300 seconds.`,
				},
				resource.Attribute{
					Name:        "stat_type_p3600",
					Description: `Data aggregation mode, cycle of 3600 seconds.`,
				},
				resource.Attribute{
					Name:        "stat_type_p5",
					Description: `Data aggregation mode, cycle of 5 seconds.`,
				},
				resource.Attribute{
					Name:        "stat_type_p600",
					Description: `Data aggregation mode, cycle of 600 seconds.`,
				},
				resource.Attribute{
					Name:        "stat_type_p60",
					Description: `Data aggregation mode, cycle of 60 seconds.`,
				},
				resource.Attribute{
					Name:        "stat_type_p86400",
					Description: `Data aggregation mode, cycle of 86400 seconds.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of this policy name.`,
				},
				resource.Attribute{
					Name:        "policy_view_name",
					Description: `Policy view name, eg:` + "`" + `cvm_device` + "`" + `,` + "`" + `BANDWIDTHPACKAGE` + "`" + `, refer to ` + "`" + `data.tencentcloud_monitor_policy_conditions(policy_view_name)` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "support_regions",
					Description: `Support regions of this policy view.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_monitor_policy_groups",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query monitor policy groups (There is a lot of data and it is recommended to output to a file)`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Policy group name for query.`,
				},
				resource.Attribute{
					Name:        "policy_view_names",
					Description: `(Optional) The policy view for query.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to store results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list policy groups. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "can_set_default",
					Description: `Whether it can be set as the default policy.`,
				},
				resource.Attribute{
					Name:        "conditions",
					Description: `A list of threshold rules. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "alarm_notify_period",
					Description: `Alarm sending cycle per second. ` + "`" + `<0` + "`" + ` does not fire, ` + "`" + `0` + "`" + ` only fires once, and ` + "`" + `>0` + "`" + ` fires every triggerTime second.`,
				},
				resource.Attribute{
					Name:        "alarm_notify_type",
					Description: `Alarm sending convergence type. ` + "`" + `0` + "`" + ` continuous alarm, ` + "`" + `1` + "`" + ` index alarm.`,
				},
				resource.Attribute{
					Name:        "calc_type",
					Description: `Compare type, ` + "`" + `1` + "`" + ` means more than, ` + "`" + `2` + "`" + ` means greater than or equal, ` + "`" + `3` + "`" + ` means less than, ` + "`" + `4` + "`" + ` means less than or equal to, ` + "`" + `5` + "`" + ` means equal, ` + "`" + `6` + "`" + ` means not equal, ` + "`" + `7` + "`" + ` means days rose, ` + "`" + `8` + "`" + ` means days fell, ` + "`" + `9` + "`" + ` means weeks rose, ` + "`" + `10` + "`" + ` means weeks fell, ` + "`" + `11` + "`" + ` means period rise, ` + "`" + `12` + "`" + ` means period fell.`,
				},
				resource.Attribute{
					Name:        "calc_value",
					Description: `Threshold value.`,
				},
				resource.Attribute{
					Name:        "continue_time",
					Description: `How long does the triggering rule last (per second).`,
				},
				resource.Attribute{
					Name:        "metric_id",
					Description: `The ID of this metric.`,
				},
				resource.Attribute{
					Name:        "metric_show_name",
					Description: `The name of this metric.`,
				},
				resource.Attribute{
					Name:        "metric_unit",
					Description: `The unit of this metric.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `Data aggregation cycle (unit second).`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `Threshold rule ID.`,
				},
				resource.Attribute{
					Name:        "event_conditions",
					Description: `A list of event rules. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "alarm_notify_period",
					Description: `Alarm sending cycle per second. ` + "`" + `<0` + "`" + ` does not fire, ` + "`" + `0` + "`" + ` only fires once, and ` + "`" + `>0` + "`" + ` fires every triggerTime second.`,
				},
				resource.Attribute{
					Name:        "alarm_notify_type",
					Description: `Alarm sending convergence type. ` + "`" + `0` + "`" + ` continuous alarm, ` + "`" + `1` + "`" + ` index alarm.`,
				},
				resource.Attribute{
					Name:        "event_id",
					Description: `The ID of this event metric.`,
				},
				resource.Attribute{
					Name:        "event_show_name",
					Description: `The name of this event metric.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `Threshold rule ID.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `The policy group id.`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `The policy group name.`,
				},
				resource.Attribute{
					Name:        "insert_time",
					Description: `The policy group create timestamp.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `If is default policy group or not, ` + "`" + `0` + "`" + ` represents the non-default policy, and ` + "`" + `1` + "`" + ` represents the default policy.`,
				},
				resource.Attribute{
					Name:        "is_open",
					Description: `Whether open or not.`,
				},
				resource.Attribute{
					Name:        "last_edit_uin",
					Description: `Recently edited user uin.`,
				},
				resource.Attribute{
					Name:        "no_shielded_sum",
					Description: `Number of unmasked instances of policy group bindings.`,
				},
				resource.Attribute{
					Name:        "parent_group_id",
					Description: `Parent policy group ID.`,
				},
				resource.Attribute{
					Name:        "policy_view_name",
					Description: `The policy group view name.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The project ID to which the policy group belongs.`,
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
					Description: `Do need a telephone alarm contact prompt.You don't need 0, you need 1.`,
				},
				resource.Attribute{
					Name:        "notify_way",
					Description: `Method of warning notification.Optional ` + "`" + `CALL` + "`" + `,` + "`" + `EMAIL` + "`" + `,` + "`" + `SITE` + "`" + `,` + "`" + `SMS` + "`" + `,` + "`" + `WECHAT` + "`" + `.`,
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
					Description: `Receive type. Optional 'group' or 'user'.`,
				},
				resource.Attribute{
					Name:        "receiver_user_list",
					Description: `Alarm receiver ID list.`,
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
					Description: `Telephone warning time.Option "OCCUR", "RECOVER".`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `Alarm period start time.Range [0,86399], which removes the date after it is converted to Beijing time as a Unix timestamp, for example 7200 means '10:0:0'.`,
				},
				resource.Attribute{
					Name:        "uid_list",
					Description: `The phone alerts the receiver uid.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `Policy group remarks.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `The policy group update timestamp.`,
				},
				resource.Attribute{
					Name:        "use_sum",
					Description: `Number of instances of policy group bindings.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list policy groups. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "can_set_default",
					Description: `Whether it can be set as the default policy.`,
				},
				resource.Attribute{
					Name:        "conditions",
					Description: `A list of threshold rules. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "alarm_notify_period",
					Description: `Alarm sending cycle per second. ` + "`" + `<0` + "`" + ` does not fire, ` + "`" + `0` + "`" + ` only fires once, and ` + "`" + `>0` + "`" + ` fires every triggerTime second.`,
				},
				resource.Attribute{
					Name:        "alarm_notify_type",
					Description: `Alarm sending convergence type. ` + "`" + `0` + "`" + ` continuous alarm, ` + "`" + `1` + "`" + ` index alarm.`,
				},
				resource.Attribute{
					Name:        "calc_type",
					Description: `Compare type, ` + "`" + `1` + "`" + ` means more than, ` + "`" + `2` + "`" + ` means greater than or equal, ` + "`" + `3` + "`" + ` means less than, ` + "`" + `4` + "`" + ` means less than or equal to, ` + "`" + `5` + "`" + ` means equal, ` + "`" + `6` + "`" + ` means not equal, ` + "`" + `7` + "`" + ` means days rose, ` + "`" + `8` + "`" + ` means days fell, ` + "`" + `9` + "`" + ` means weeks rose, ` + "`" + `10` + "`" + ` means weeks fell, ` + "`" + `11` + "`" + ` means period rise, ` + "`" + `12` + "`" + ` means period fell.`,
				},
				resource.Attribute{
					Name:        "calc_value",
					Description: `Threshold value.`,
				},
				resource.Attribute{
					Name:        "continue_time",
					Description: `How long does the triggering rule last (per second).`,
				},
				resource.Attribute{
					Name:        "metric_id",
					Description: `The ID of this metric.`,
				},
				resource.Attribute{
					Name:        "metric_show_name",
					Description: `The name of this metric.`,
				},
				resource.Attribute{
					Name:        "metric_unit",
					Description: `The unit of this metric.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `Data aggregation cycle (unit second).`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `Threshold rule ID.`,
				},
				resource.Attribute{
					Name:        "event_conditions",
					Description: `A list of event rules. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "alarm_notify_period",
					Description: `Alarm sending cycle per second. ` + "`" + `<0` + "`" + ` does not fire, ` + "`" + `0` + "`" + ` only fires once, and ` + "`" + `>0` + "`" + ` fires every triggerTime second.`,
				},
				resource.Attribute{
					Name:        "alarm_notify_type",
					Description: `Alarm sending convergence type. ` + "`" + `0` + "`" + ` continuous alarm, ` + "`" + `1` + "`" + ` index alarm.`,
				},
				resource.Attribute{
					Name:        "event_id",
					Description: `The ID of this event metric.`,
				},
				resource.Attribute{
					Name:        "event_show_name",
					Description: `The name of this event metric.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `Threshold rule ID.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `The policy group id.`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `The policy group name.`,
				},
				resource.Attribute{
					Name:        "insert_time",
					Description: `The policy group create timestamp.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `If is default policy group or not, ` + "`" + `0` + "`" + ` represents the non-default policy, and ` + "`" + `1` + "`" + ` represents the default policy.`,
				},
				resource.Attribute{
					Name:        "is_open",
					Description: `Whether open or not.`,
				},
				resource.Attribute{
					Name:        "last_edit_uin",
					Description: `Recently edited user uin.`,
				},
				resource.Attribute{
					Name:        "no_shielded_sum",
					Description: `Number of unmasked instances of policy group bindings.`,
				},
				resource.Attribute{
					Name:        "parent_group_id",
					Description: `Parent policy group ID.`,
				},
				resource.Attribute{
					Name:        "policy_view_name",
					Description: `The policy group view name.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The project ID to which the policy group belongs.`,
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
					Description: `Do need a telephone alarm contact prompt.You don't need 0, you need 1.`,
				},
				resource.Attribute{
					Name:        "notify_way",
					Description: `Method of warning notification.Optional ` + "`" + `CALL` + "`" + `,` + "`" + `EMAIL` + "`" + `,` + "`" + `SITE` + "`" + `,` + "`" + `SMS` + "`" + `,` + "`" + `WECHAT` + "`" + `.`,
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
					Description: `Receive type. Optional 'group' or 'user'.`,
				},
				resource.Attribute{
					Name:        "receiver_user_list",
					Description: `Alarm receiver ID list.`,
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
					Description: `Telephone warning time.Option "OCCUR", "RECOVER".`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `Alarm period start time.Range [0,86399], which removes the date after it is converted to Beijing time as a Unix timestamp, for example 7200 means '10:0:0'.`,
				},
				resource.Attribute{
					Name:        "uid_list",
					Description: `The phone alerts the receiver uid.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `Policy group remarks.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `The policy group update timestamp.`,
				},
				resource.Attribute{
					Name:        "use_sum",
					Description: `Number of instances of policy group bindings.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_monitor_product_event",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query monitor events(There is a lot of data and it is recommended to output to a file)`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dimensions",
					Description: `(Optional) Dimensional composition of instance objects.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `(Optional) End timestamp for this query, eg:` + "`" + `1588232111` + "`" + `. Default start time is ` + "`" + `now-3000` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "event_name",
					Description: `(Optional) Event name filtering, such as ` + "`" + `guest_reboot` + "`" + ` indicates that the machine restart.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Optional) Affect objects, such as ` + "`" + `ins-19708ino` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "is_alarm_config",
					Description: `(Optional) Alarm status configuration filter, 1means configured, 0(default) means not configured.`,
				},
				resource.Attribute{
					Name:        "product_name",
					Description: `(Optional) Product type filtering, such as ` + "`" + `cvm` + "`" + ` for cloud server.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) Project ID filter.`,
				},
				resource.Attribute{
					Name:        "region_list",
					Description: `(Optional) Region filter, such as ` + "`" + `gz` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to store results.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `(Optional) Start timestamp for this query, eg:` + "`" + `1588230000` + "`" + `. Default start time is ` + "`" + `now-3600` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Event status filter, value range ` + "`" + `-` + "`" + `,` + "`" + `alarm` + "`" + `,` + "`" + `recover` + "`" + `, indicating recovered, unrecovered and stateless.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Event type filtering, with value range ` + "`" + `abnormal` + "`" + `,` + "`" + `status_change` + "`" + `, indicating state change and abnormal events. The ` + "`" + `dimensions` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Instance dimension name, eg: ` + "`" + `deviceWanIp` + "`" + ` for internet ip.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Instance dimension value, eg: ` + "`" + `119.119.119.119` + "`" + ` for internet ip. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list events. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "addition_msg",
					Description: `A list of addition message. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of this addition message.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of this addition message.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of this addition message.`,
				},
				resource.Attribute{
					Name:        "dimensions",
					Description: `A list of event dimensions. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of this dimension.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of this dimension.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of this dimension.`,
				},
				resource.Attribute{
					Name:        "event_cname",
					Description: `Event chinese name.`,
				},
				resource.Attribute{
					Name:        "event_ename",
					Description: `Event english name.`,
				},
				resource.Attribute{
					Name:        "event_id",
					Description: `Event ID.`,
				},
				resource.Attribute{
					Name:        "event_name",
					Description: `Event short name.`,
				},
				resource.Attribute{
					Name:        "group_info",
					Description: `A list of group info. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `Policy group ID.`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `Policy group name.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The instance ID of this event.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `The name of this instance.`,
				},
				resource.Attribute{
					Name:        "is_alarm_config",
					Description: `Whether to configure alarm.`,
				},
				resource.Attribute{
					Name:        "product_cname",
					Description: `Product chinese name.`,
				},
				resource.Attribute{
					Name:        "product_ename",
					Description: `Product english name.`,
				},
				resource.Attribute{
					Name:        "product_name",
					Description: `Product short name.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Project ID of this instance.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region of this instance.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `The start timestamp of this event.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of this event.`,
				},
				resource.Attribute{
					Name:        "support_alarm",
					Description: `Whether to support alarm.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of this event.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `The update timestamp of this event.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list events. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "addition_msg",
					Description: `A list of addition message. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of this addition message.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of this addition message.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of this addition message.`,
				},
				resource.Attribute{
					Name:        "dimensions",
					Description: `A list of event dimensions. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of this dimension.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of this dimension.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of this dimension.`,
				},
				resource.Attribute{
					Name:        "event_cname",
					Description: `Event chinese name.`,
				},
				resource.Attribute{
					Name:        "event_ename",
					Description: `Event english name.`,
				},
				resource.Attribute{
					Name:        "event_id",
					Description: `Event ID.`,
				},
				resource.Attribute{
					Name:        "event_name",
					Description: `Event short name.`,
				},
				resource.Attribute{
					Name:        "group_info",
					Description: `A list of group info. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `Policy group ID.`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `Policy group name.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The instance ID of this event.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `The name of this instance.`,
				},
				resource.Attribute{
					Name:        "is_alarm_config",
					Description: `Whether to configure alarm.`,
				},
				resource.Attribute{
					Name:        "product_cname",
					Description: `Product chinese name.`,
				},
				resource.Attribute{
					Name:        "product_ename",
					Description: `Product english name.`,
				},
				resource.Attribute{
					Name:        "product_name",
					Description: `Product short name.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Project ID of this instance.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region of this instance.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `The start timestamp of this event.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of this event.`,
				},
				resource.Attribute{
					Name:        "support_alarm",
					Description: `Whether to support alarm.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of this event.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `The update timestamp of this event.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_monitor_product_namespace",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query product namespace in monitor)`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name for filter, eg:` + "`" + `Load Banlancer` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to store results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list product namespaces. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `Namespace of each cloud product in monitor system.`,
				},
				resource.Attribute{
					Name:        "product_chinese_name",
					Description: `Chinese name of this product.`,
				},
				resource.Attribute{
					Name:        "product_name",
					Description: `English name of this product.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list product namespaces. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `Namespace of each cloud product in monitor system.`,
				},
				resource.Attribute{
					Name:        "product_chinese_name",
					Description: `Chinese name of this product.`,
				},
				resource.Attribute{
					Name:        "product_name",
					Description: `English name of this product.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_mysql_backup_list",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query the list of backup databases.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "mysql_id",
					Description: `(Required) Instance ID, such as ` + "`" + `cdb-c1nl9rpv` + "`" + `. It is identical to the instance ID displayed in the database console page.`,
				},
				resource.Attribute{
					Name:        "max_number",
					Description: `(Optional) The latest files to list, rang from 1 to 10000. And the default value is ` + "`" + `10` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to store results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list of MySQL backup. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "backup_id",
					Description: `ID of Backup task.`,
				},
				resource.Attribute{
					Name:        "backup_model",
					Description: `Backup method. Supported values include: ` + "`" + `physical` + "`" + ` - physical backup, and ` + "`" + `logical` + "`" + ` - logical backup.`,
				},
				resource.Attribute{
					Name:        "creator",
					Description: `The owner of the backup files.`,
				},
				resource.Attribute{
					Name:        "finish_time",
					Description: `The time at which the backup finishes.`,
				},
				resource.Attribute{
					Name:        "internet_url",
					Description: `URL for downloads externally.`,
				},
				resource.Attribute{
					Name:        "intranet_url",
					Description: `URL for downloads internally.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `the size of backup file.`,
				},
				resource.Attribute{
					Name:        "time",
					Description: `The earliest time at which the backup starts. For example, ` + "`" + `2` + "`" + ` indicates 2:00 am.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list of MySQL backup. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "backup_id",
					Description: `ID of Backup task.`,
				},
				resource.Attribute{
					Name:        "backup_model",
					Description: `Backup method. Supported values include: ` + "`" + `physical` + "`" + ` - physical backup, and ` + "`" + `logical` + "`" + ` - logical backup.`,
				},
				resource.Attribute{
					Name:        "creator",
					Description: `The owner of the backup files.`,
				},
				resource.Attribute{
					Name:        "finish_time",
					Description: `The time at which the backup finishes.`,
				},
				resource.Attribute{
					Name:        "internet_url",
					Description: `URL for downloads externally.`,
				},
				resource.Attribute{
					Name:        "intranet_url",
					Description: `URL for downloads internally.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `the size of backup file.`,
				},
				resource.Attribute{
					Name:        "time",
					Description: `The earliest time at which the backup starts. For example, ` + "`" + `2` + "`" + ` indicates 2:00 am.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_mysql_instance",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information about a MySQL instance.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "charge_type",
					Description: `(Optional) Pay type of instance, valid values are ` + "`" + `PREPAID` + "`" + ` and ` + "`" + `POSTPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `(Optional) The version number of the database engine to use. Supported versions include 5.5/5.6/5.7/8.0.`,
				},
				resource.Attribute{
					Name:        "init_flag",
					Description: `(Optional) Initialization mark. Available values: ` + "`" + `0` + "`" + ` - Uninitialized; ` + "`" + `1` + "`" + ` - Initialized.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Optional) Name of mysql instance.`,
				},
				resource.Attribute{
					Name:        "instance_role",
					Description: `(Optional) Instance type. Supported values include: ` + "`" + `master` + "`" + ` - master instance, ` + "`" + `dr` + "`" + ` - disaster recovery instance, and ` + "`" + `ro` + "`" + ` - read-only instance.`,
				},
				resource.Attribute{
					Name:        "limit",
					Description: `(Optional) Number of results returned for a single request. Default is ` + "`" + `20` + "`" + `, and maximum is 2000.`,
				},
				resource.Attribute{
					Name:        "mysql_id",
					Description: `(Optional) Instance ID, such as ` + "`" + `cdb-c1nl9rpv` + "`" + `. It is identical to the instance ID displayed in the database console page.`,
				},
				resource.Attribute{
					Name:        "offset",
					Description: `(Optional) Record offset. Default is 0.`,
				},
				resource.Attribute{
					Name:        "pay_type",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to store results.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Optional) Security groups ID of instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Instance status. Available values: ` + "`" + `0` + "`" + ` - Creating; ` + "`" + `1` + "`" + ` - Running; ` + "`" + `4` + "`" + ` - Isolating; ` + "`" + `5` + "`" + ` - Isolated.`,
				},
				resource.Attribute{
					Name:        "with_dr",
					Description: `(Optional) Indicates whether to query disaster recovery instances.`,
				},
				resource.Attribute{
					Name:        "with_master",
					Description: `(Optional) Indicates whether to query master instances.`,
				},
				resource.Attribute{
					Name:        "with_ro",
					Description: `(Optional) Indicates whether to query read-only instances. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "instance_list",
					Description: `A list of instances. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "auto_renew_flag",
					Description: `Auto renew flag. NOTES: Only supported prepay instance.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `Pay type of instance.`,
				},
				resource.Attribute{
					Name:        "cpu_core_count",
					Description: `CPU count.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time at which a instance is created.`,
				},
				resource.Attribute{
					Name:        "dead_line_time",
					Description: `Expire date of instance. NOTES: Only supported prepay instance.`,
				},
				resource.Attribute{
					Name:        "device_type",
					Description: `Supported instance model. ` + "`" + `HA` + "`" + ` - high available version; ` + "`" + `Basic` + "`" + ` - basic version.`,
				},
				resource.Attribute{
					Name:        "dr_instance_ids",
					Description: `ID list of disaster-recovery type associated with the current instance.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `The version number of the database engine to use. Supported versions include ` + "`" + `5.5` + "`" + `/` + "`" + `5.6` + "`" + `/` + "`" + `5.7` + "`" + `/` + "`" + `8.0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "init_flag",
					Description: `Initialization mark. Available values: ` + "`" + `0` + "`" + ` - Uninitialized; ` + "`" + `1` + "`" + ` - Initialized.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `Name of mysql instance.`,
				},
				resource.Attribute{
					Name:        "instance_role",
					Description: `Instance type. Supported values include: ` + "`" + `master` + "`" + ` - master instance, ` + "`" + `dr` + "`" + ` - disaster recovery instance, and ` + "`" + `ro` + "`" + ` - read-only instance.`,
				},
				resource.Attribute{
					Name:        "internet_host",
					Description: `Public network domain name.`,
				},
				resource.Attribute{
					Name:        "internet_port",
					Description: `Public network port.`,
				},
				resource.Attribute{
					Name:        "internet_status",
					Description: `Status of public network.`,
				},
				resource.Attribute{
					Name:        "intranet_ip",
					Description: `Instance IP for internal access.`,
				},
				resource.Attribute{
					Name:        "intranet_port",
					Description: `Transport layer port number for internal purpose.`,
				},
				resource.Attribute{
					Name:        "master_instance_id",
					Description: `Indicates the master instance ID of recovery instances.`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `Memory size (in MB).`,
				},
				resource.Attribute{
					Name:        "mysql_id",
					Description: `Instance ID, such as ` + "`" + `cdb-c1nl9rpv` + "`" + `. It is identical to the instance ID displayed in the database console page.`,
				},
				resource.Attribute{
					Name:        "pay_type",
					Description: `Pay type of instance, ` + "`" + `0` + "`" + `: prepaid, ` + "`" + `1` + "`" + `: postpaid.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Project ID to which the current instance belongs.`,
				},
				resource.Attribute{
					Name:        "ro_instance_ids",
					Description: `ID list of read-only type associated with the current instance.`,
				},
				resource.Attribute{
					Name:        "slave_sync_mode",
					Description: `Data replication mode. ` + "`" + `0` + "`" + ` - Async replication; ` + "`" + `1` + "`" + ` - Semisync replication; ` + "`" + `2` + "`" + ` - Strongsync replication.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Instance status. Available values: ` + "`" + `0` + "`" + ` - Creating; ` + "`" + `1` + "`" + ` - Running; ` + "`" + `4` + "`" + ` - Isolating; ` + "`" + `5` + "`" + ` - Isolated.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of subnet to which the current instance belongs.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `Disk capacity (in GB).`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of Virtual Private Cloud.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `Information of available zone.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_list",
					Description: `A list of instances. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "auto_renew_flag",
					Description: `Auto renew flag. NOTES: Only supported prepay instance.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `Pay type of instance.`,
				},
				resource.Attribute{
					Name:        "cpu_core_count",
					Description: `CPU count.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time at which a instance is created.`,
				},
				resource.Attribute{
					Name:        "dead_line_time",
					Description: `Expire date of instance. NOTES: Only supported prepay instance.`,
				},
				resource.Attribute{
					Name:        "device_type",
					Description: `Supported instance model. ` + "`" + `HA` + "`" + ` - high available version; ` + "`" + `Basic` + "`" + ` - basic version.`,
				},
				resource.Attribute{
					Name:        "dr_instance_ids",
					Description: `ID list of disaster-recovery type associated with the current instance.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `The version number of the database engine to use. Supported versions include ` + "`" + `5.5` + "`" + `/` + "`" + `5.6` + "`" + `/` + "`" + `5.7` + "`" + `/` + "`" + `8.0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "init_flag",
					Description: `Initialization mark. Available values: ` + "`" + `0` + "`" + ` - Uninitialized; ` + "`" + `1` + "`" + ` - Initialized.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `Name of mysql instance.`,
				},
				resource.Attribute{
					Name:        "instance_role",
					Description: `Instance type. Supported values include: ` + "`" + `master` + "`" + ` - master instance, ` + "`" + `dr` + "`" + ` - disaster recovery instance, and ` + "`" + `ro` + "`" + ` - read-only instance.`,
				},
				resource.Attribute{
					Name:        "internet_host",
					Description: `Public network domain name.`,
				},
				resource.Attribute{
					Name:        "internet_port",
					Description: `Public network port.`,
				},
				resource.Attribute{
					Name:        "internet_status",
					Description: `Status of public network.`,
				},
				resource.Attribute{
					Name:        "intranet_ip",
					Description: `Instance IP for internal access.`,
				},
				resource.Attribute{
					Name:        "intranet_port",
					Description: `Transport layer port number for internal purpose.`,
				},
				resource.Attribute{
					Name:        "master_instance_id",
					Description: `Indicates the master instance ID of recovery instances.`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `Memory size (in MB).`,
				},
				resource.Attribute{
					Name:        "mysql_id",
					Description: `Instance ID, such as ` + "`" + `cdb-c1nl9rpv` + "`" + `. It is identical to the instance ID displayed in the database console page.`,
				},
				resource.Attribute{
					Name:        "pay_type",
					Description: `Pay type of instance, ` + "`" + `0` + "`" + `: prepaid, ` + "`" + `1` + "`" + `: postpaid.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Project ID to which the current instance belongs.`,
				},
				resource.Attribute{
					Name:        "ro_instance_ids",
					Description: `ID list of read-only type associated with the current instance.`,
				},
				resource.Attribute{
					Name:        "slave_sync_mode",
					Description: `Data replication mode. ` + "`" + `0` + "`" + ` - Async replication; ` + "`" + `1` + "`" + ` - Semisync replication; ` + "`" + `2` + "`" + ` - Strongsync replication.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Instance status. Available values: ` + "`" + `0` + "`" + ` - Creating; ` + "`" + `1` + "`" + ` - Running; ` + "`" + `4` + "`" + ` - Isolating; ` + "`" + `5` + "`" + ` - Isolated.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of subnet to which the current instance belongs.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `Disk capacity (in GB).`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of Virtual Private Cloud.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `Information of available zone.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_mysql_parameter_list",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information about a parameter group of a database instance.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "engine_version",
					Description: `(Optional) The version number of the database engine to use. Supported versions include 5.5/5.6/5.7.`,
				},
				resource.Attribute{
					Name:        "mysql_id",
					Description: `(Optional) Instance ID.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to store results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "parameter_list",
					Description: `A list of parameters. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "current_value",
					Description: `Current value.`,
				},
				resource.Attribute{
					Name:        "default_value",
					Description: `Default value.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Parameter specification description.`,
				},
				resource.Attribute{
					Name:        "enum_value",
					Description: `Enumerated value.`,
				},
				resource.Attribute{
					Name:        "max",
					Description: `Maximum value for the parameter.`,
				},
				resource.Attribute{
					Name:        "min",
					Description: `Minimum value for the parameter.`,
				},
				resource.Attribute{
					Name:        "need_reboot",
					Description: `Indicates whether reboot is needed to enable the new parameters.`,
				},
				resource.Attribute{
					Name:        "parameter_name",
					Description: `Parameter name.`,
				},
				resource.Attribute{
					Name:        "parameter_type",
					Description: `Parameter type.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "parameter_list",
					Description: `A list of parameters. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "current_value",
					Description: `Current value.`,
				},
				resource.Attribute{
					Name:        "default_value",
					Description: `Default value.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Parameter specification description.`,
				},
				resource.Attribute{
					Name:        "enum_value",
					Description: `Enumerated value.`,
				},
				resource.Attribute{
					Name:        "max",
					Description: `Maximum value for the parameter.`,
				},
				resource.Attribute{
					Name:        "min",
					Description: `Minimum value for the parameter.`,
				},
				resource.Attribute{
					Name:        "need_reboot",
					Description: `Indicates whether reboot is needed to enable the new parameters.`,
				},
				resource.Attribute{
					Name:        "parameter_name",
					Description: `Parameter name.`,
				},
				resource.Attribute{
					Name:        "parameter_type",
					Description: `Parameter type.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_mysql_zone_config",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query the available database specifications for different regions. And a maximum of 20 requests can be initiated per second for this query.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Region parameter, which is used to identify the region to which the data you want to work with belongs.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to store results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list of zone config. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "disaster_recovery_zones",
					Description: `Information about available zones of recovery.`,
				},
				resource.Attribute{
					Name:        "engine_versions",
					Description: `The version number of the database engine to use. Supported versions include ` + "`" + `5.5` + "`" + `/` + "`" + `5.6` + "`" + `/` + "`" + `5.7` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "first_slave_zones",
					Description: `Zone information about first slave instance.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Indicates whether the current DC is the default DC for the region. Possible returned values: ` + "`" + `0` + "`" + ` - no; ` + "`" + `1` + "`" + ` - yes.`,
				},
				resource.Attribute{
					Name:        "is_support_disaster_recovery",
					Description: `Indicates whether recovery is supported: ` + "`" + `0` + "`" + ` - No; ` + "`" + `1` + "`" + ` - Yes.`,
				},
				resource.Attribute{
					Name:        "is_support_vpc",
					Description: `Indicates whether VPC is supported: ` + "`" + `0` + "`" + ` - No; ` + "`" + `1` + "`" + ` - Yes.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of available zone which is equal to a specific datacenter.`,
				},
				resource.Attribute{
					Name:        "second_slave_zones",
					Description: `Zone information about second slave instance.`,
				},
				resource.Attribute{
					Name:        "sells",
					Description: `A list of supported instance types for sell:`,
				},
				resource.Attribute{
					Name:        "max_volume_size",
					Description: `Maximum disk size (in GB).`,
				},
				resource.Attribute{
					Name:        "mem_size",
					Description: `Memory size (in MB).`,
				},
				resource.Attribute{
					Name:        "min_volume_size",
					Description: `Minimum disk size (in GB).`,
				},
				resource.Attribute{
					Name:        "qps",
					Description: `Queries per second.`,
				},
				resource.Attribute{
					Name:        "volume_step",
					Description: `Disk increment (in GB).`,
				},
				resource.Attribute{
					Name:        "slave_deploy_modes",
					Description: `Availability zone deployment method. Available values: ` + "`" + `0` + "`" + ` - Single availability zone; ` + "`" + `1` + "`" + ` - Multiple availability zones.`,
				},
				resource.Attribute{
					Name:        "support_slave_sync_modes",
					Description: `Data replication mode. ` + "`" + `0` + "`" + ` - Async replication; ` + "`" + `1` + "`" + ` - Semisync replication; ` + "`" + `2` + "`" + ` - Strongsync replication.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list of zone config. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "disaster_recovery_zones",
					Description: `Information about available zones of recovery.`,
				},
				resource.Attribute{
					Name:        "engine_versions",
					Description: `The version number of the database engine to use. Supported versions include ` + "`" + `5.5` + "`" + `/` + "`" + `5.6` + "`" + `/` + "`" + `5.7` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "first_slave_zones",
					Description: `Zone information about first slave instance.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Indicates whether the current DC is the default DC for the region. Possible returned values: ` + "`" + `0` + "`" + ` - no; ` + "`" + `1` + "`" + ` - yes.`,
				},
				resource.Attribute{
					Name:        "is_support_disaster_recovery",
					Description: `Indicates whether recovery is supported: ` + "`" + `0` + "`" + ` - No; ` + "`" + `1` + "`" + ` - Yes.`,
				},
				resource.Attribute{
					Name:        "is_support_vpc",
					Description: `Indicates whether VPC is supported: ` + "`" + `0` + "`" + ` - No; ` + "`" + `1` + "`" + ` - Yes.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of available zone which is equal to a specific datacenter.`,
				},
				resource.Attribute{
					Name:        "second_slave_zones",
					Description: `Zone information about second slave instance.`,
				},
				resource.Attribute{
					Name:        "sells",
					Description: `A list of supported instance types for sell:`,
				},
				resource.Attribute{
					Name:        "max_volume_size",
					Description: `Maximum disk size (in GB).`,
				},
				resource.Attribute{
					Name:        "mem_size",
					Description: `Memory size (in MB).`,
				},
				resource.Attribute{
					Name:        "min_volume_size",
					Description: `Minimum disk size (in GB).`,
				},
				resource.Attribute{
					Name:        "qps",
					Description: `Queries per second.`,
				},
				resource.Attribute{
					Name:        "volume_step",
					Description: `Disk increment (in GB).`,
				},
				resource.Attribute{
					Name:        "slave_deploy_modes",
					Description: `Availability zone deployment method. Available values: ` + "`" + `0` + "`" + ` - Single availability zone; ` + "`" + `1` + "`" + ` - Multiple availability zones.`,
				},
				resource.Attribute{
					Name:        "support_slave_sync_modes",
					Description: `Data replication mode. ` + "`" + `0` + "`" + ` - Async replication; ` + "`" + `1` + "`" + ` - Semisync replication; ` + "`" + `2` + "`" + ` - Strongsync replication.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_nat_gateways",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of NAT gateways.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) ID of the VPC. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "nats",
					Description: `Information list of the dedicated NATs.`,
				},
				resource.Attribute{
					Name:        "assigned_eip_set",
					Description: `EIP IP address set bound to the gateway. The value of at least 1.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `The maximum public network output bandwidth of NAT gateway (unit: Mbps), the available values include: 20,50,100,200,500,1000,2000,5000. Default is 100.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "max_concurrent",
					Description: `The upper limit of concurrent connection of NAT gateway, the available values include: 1000000,3000000,10000000. Default is 1000000.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The available tags within this NAT gateway.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "nats",
					Description: `Information list of the dedicated NATs.`,
				},
				resource.Attribute{
					Name:        "assigned_eip_set",
					Description: `EIP IP address set bound to the gateway. The value of at least 1.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `The maximum public network output bandwidth of NAT gateway (unit: Mbps), the available values include: 20,50,100,200,500,1000,2000,5000. Default is 100.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "max_concurrent",
					Description: `The upper limit of concurrent connection of NAT gateway, the available values include: 1000000,3000000,10000000. Default is 1000000.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The available tags within this NAT gateway.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_nats",
			Category:         "Data Sources",
			ShortDescription: `The NATs data source lists a number of NATs resource information owned by an TencentCloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Optional) The maximum public network output bandwidth of the gateway (unit: Mbps), for example: ` + "`" + `10` + "`" + `, ` + "`" + `20` + "`" + `, ` + "`" + `50` + "`" + `, ` + "`" + `100` + "`" + `, ` + "`" + `200` + "`" + `, ` + "`" + `500` + "`" + `, ` + "`" + `1000` + "`" + `, ` + "`" + `2000` + "`" + `, ` + "`" + `5000` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID for NAT Gateway.`,
				},
				resource.Attribute{
					Name:        "max_concurrent",
					Description: `(Optional) The upper limit of concurrent connection of NAT gateway, for example: ` + "`" + `1000000` + "`" + `, ` + "`" + `3000000` + "`" + `, ` + "`" + `10000000` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name for NAT Gateway.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) NAT gateway status. Valid values: 0, 1, 2. 0: Running, 1: Unavailable, 2: Be in arrears and out of service.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) The VPC ID for NAT Gateway. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "nats",
					Description: `Information list of the dedicated tunnels.`,
				},
				resource.Attribute{
					Name:        "assigned_eip_set",
					Description: `Elastic IP arrays bound to the gateway.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `The maximum public network output bandwidth of the gateway (unit: Mbps), for example: ` + "`" + `10` + "`" + `, ` + "`" + `20` + "`" + `, ` + "`" + `50` + "`" + `, ` + "`" + `100` + "`" + `, ` + "`" + `200` + "`" + `, ` + "`" + `500` + "`" + `, ` + "`" + `1000` + "`" + `, ` + "`" + `2000` + "`" + `, ` + "`" + `5000` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The create time of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID for NAT Gateway.`,
				},
				resource.Attribute{
					Name:        "max_concurrent",
					Description: `The upper limit of concurrent connection of NAT gateway, for example: ` + "`" + `1000000` + "`" + `, ` + "`" + `3000000` + "`" + `, ` + "`" + `10000000` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name for NAT Gateway.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `NAT gateway status, ` + "`" + `0` + "`" + `: Running, ` + "`" + `1` + "`" + `: Unavailable, ` + "`" + `2` + "`" + `: Be in arrears and out of service.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The VPC ID for NAT Gateway.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "nats",
					Description: `Information list of the dedicated tunnels.`,
				},
				resource.Attribute{
					Name:        "assigned_eip_set",
					Description: `Elastic IP arrays bound to the gateway.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `The maximum public network output bandwidth of the gateway (unit: Mbps), for example: ` + "`" + `10` + "`" + `, ` + "`" + `20` + "`" + `, ` + "`" + `50` + "`" + `, ` + "`" + `100` + "`" + `, ` + "`" + `200` + "`" + `, ` + "`" + `500` + "`" + `, ` + "`" + `1000` + "`" + `, ` + "`" + `2000` + "`" + `, ` + "`" + `5000` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The create time of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID for NAT Gateway.`,
				},
				resource.Attribute{
					Name:        "max_concurrent",
					Description: `The upper limit of concurrent connection of NAT gateway, for example: ` + "`" + `1000000` + "`" + `, ` + "`" + `3000000` + "`" + `, ` + "`" + `10000000` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name for NAT Gateway.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `NAT gateway status, ` + "`" + `0` + "`" + `: Running, ` + "`" + `1` + "`" + `: Unavailable, ` + "`" + `2` + "`" + `: Be in arrears and out of service.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The VPC ID for NAT Gateway.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_placement_groups",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query placement groups.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the placement group to be queried.`,
				},
				resource.Attribute{
					Name:        "placement_group_id",
					Description: `(Optional) ID of the placement group to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "placement_group_list",
					Description: `An information list of placement group. Each element contains the following attributes:`,
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
					Description: `Maximum number of hosts in the placement group.`,
				},
				resource.Attribute{
					Name:        "instance_ids",
					Description: `Host IDs in the placement group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the placement group.`,
				},
				resource.Attribute{
					Name:        "placement_group_id",
					Description: `ID of the placement group.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the placement group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "placement_group_list",
					Description: `An information list of placement group. Each element contains the following attributes:`,
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
					Description: `Maximum number of hosts in the placement group.`,
				},
				resource.Attribute{
					Name:        "instance_ids",
					Description: `Host IDs in the placement group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the placement group.`,
				},
				resource.Attribute{
					Name:        "placement_group_id",
					Description: `ID of the placement group.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the placement group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_postgresql_instances",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query postgresql instances`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the postgresql instance to be query.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the postgresql instance to be query.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) Project ID of the postgresql instance to be query.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "instance_list",
					Description: `A list of postgresql instances. Each element contains the following attributes.`,
				},
				resource.Attribute{
					Name:        "auto_renew_flag",
					Description: `Auto renew flag.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Availability zone.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `Pay type of the postgresql instance.`,
				},
				resource.Attribute{
					Name:        "charset",
					Description: `Charset of the postgresql instance.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the postgresql instance.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `Version of the postgresql database engine.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the postgresql instance.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `Memory size(in GB).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the postgresql instance.`,
				},
				resource.Attribute{
					Name:        "private_access_ip",
					Description: `IP address for private access.`,
				},
				resource.Attribute{
					Name:        "private_access_port",
					Description: `Port for private access.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Project id, default value is 0.`,
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
					Name:        "public_access_switch",
					Description: `Indicates whether to enable the access to an instance from public network or not.`,
				},
				resource.Attribute{
					Name:        "storage",
					Description: `Volume size(in GB).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of subnet.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The available tags within this postgresql.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of VPC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_list",
					Description: `A list of postgresql instances. Each element contains the following attributes.`,
				},
				resource.Attribute{
					Name:        "auto_renew_flag",
					Description: `Auto renew flag.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Availability zone.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `Pay type of the postgresql instance.`,
				},
				resource.Attribute{
					Name:        "charset",
					Description: `Charset of the postgresql instance.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the postgresql instance.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `Version of the postgresql database engine.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the postgresql instance.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `Memory size(in GB).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the postgresql instance.`,
				},
				resource.Attribute{
					Name:        "private_access_ip",
					Description: `IP address for private access.`,
				},
				resource.Attribute{
					Name:        "private_access_port",
					Description: `Port for private access.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Project id, default value is 0.`,
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
					Name:        "public_access_switch",
					Description: `Indicates whether to enable the access to an instance from public network or not.`,
				},
				resource.Attribute{
					Name:        "storage",
					Description: `Volume size(in GB).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of subnet.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The available tags within this postgresql.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of VPC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_postgresql_specinfos",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get the available product configs of the postgresql instance.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required) The zone of the postgresql instance to query.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list of zones will be exported and its every element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `The CPU number of the postgresql instance.`,
				},
				resource.Attribute{
					Name:        "engine_version_name",
					Description: `Version name of the postgresql database engine.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `Version of the postgresql database engine.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the postgresql instance speccode.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `Memory size(in GB).`,
				},
				resource.Attribute{
					Name:        "qps",
					Description: `The QPS of the postgresql instance.`,
				},
				resource.Attribute{
					Name:        "storage_max",
					Description: `The maximum volume size(in GB).`,
				},
				resource.Attribute{
					Name:        "storage_min",
					Description: `The minimum volume size(in GB).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list of zones will be exported and its every element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `The CPU number of the postgresql instance.`,
				},
				resource.Attribute{
					Name:        "engine_version_name",
					Description: `Version name of the postgresql database engine.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `Version of the postgresql database engine.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the postgresql instance speccode.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `Memory size(in GB).`,
				},
				resource.Attribute{
					Name:        "qps",
					Description: `The QPS of the postgresql instance.`,
				},
				resource.Attribute{
					Name:        "storage_max",
					Description: `The maximum volume size(in GB).`,
				},
				resource.Attribute{
					Name:        "storage_min",
					Description: `The minimum volume size(in GB).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_protocol_template_groups",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of protocol template groups.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the protocol template group to query.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the protocol template group to query.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "group_list",
					Description: `Information list of the dedicated protocol template groups.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of the protocol template group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of protocol template group.`,
				},
				resource.Attribute{
					Name:        "template_ids",
					Description: `ID set of the protocol template.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "group_list",
					Description: `Information list of the dedicated protocol template groups.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of the protocol template group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of protocol template group.`,
				},
				resource.Attribute{
					Name:        "template_ids",
					Description: `ID set of the protocol template.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_protocol_templates",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of protocol templates.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the protocol template to query.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the protocol template to query.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "template_list",
					Description: `Information list of the dedicated protocol templates.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the protocol template.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of protocol template.`,
				},
				resource.Attribute{
					Name:        "protocols",
					Description: `Set of the protocols.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "template_list",
					Description: `Information list of the dedicated protocol templates.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the protocol template.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of protocol template.`,
				},
				resource.Attribute{
					Name:        "protocols",
					Description: `Set of the protocols.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_redis_instances",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query the detail information of redis instance.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "limit",
					Description: `(Optional) The number limitation of results for a query.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) ID of the project to which redis instance belongs.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "search_key",
					Description: `(Optional) Key words used to match the results, and the key words can be: instance ID, instance name and IP address.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of redis instance.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) ID of an available zone. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "instance_list",
					Description: `A list of redis instance. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `The charge type of instance. Valid values are ` + "`" + `POSTPAID` + "`" + ` and ` + "`" + `PREPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time when the instance is created.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `IP address of an instance.`,
				},
				resource.Attribute{
					Name:        "mem_size",
					Description: `Memory size in MB.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of a redis instance.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port used to access a redis instance.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of the project to which a redis instance belongs.`,
				},
				resource.Attribute{
					Name:        "redis_id",
					Description: `ID of a redis instance.`,
				},
				resource.Attribute{
					Name:        "redis_replicas_num",
					Description: `The number of instance copies.`,
				},
				resource.Attribute{
					Name:        "redis_shard_num",
					Description: `The number of instance shard.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of an instance, maybe: ` + "`" + `init` + "`" + `, ` + "`" + `processing` + "`" + `, ` + "`" + `online` + "`" + `, ` + "`" + `isolate` + "`" + ` and ` + "`" + `todelete` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of the vpc subnet.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of an instance.`,
				},
				resource.Attribute{
					Name:        "type_id",
					Description: `Instance type. Refer to ` + "`" + `data.tencentcloud_redis_zone_config.list.type_id` + "`" + ` get available values.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the vpc with which the instance is associated.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `Available zone to which a redis instance belongs.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_list",
					Description: `A list of redis instance. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `The charge type of instance. Valid values are ` + "`" + `POSTPAID` + "`" + ` and ` + "`" + `PREPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time when the instance is created.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `IP address of an instance.`,
				},
				resource.Attribute{
					Name:        "mem_size",
					Description: `Memory size in MB.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of a redis instance.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port used to access a redis instance.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of the project to which a redis instance belongs.`,
				},
				resource.Attribute{
					Name:        "redis_id",
					Description: `ID of a redis instance.`,
				},
				resource.Attribute{
					Name:        "redis_replicas_num",
					Description: `The number of instance copies.`,
				},
				resource.Attribute{
					Name:        "redis_shard_num",
					Description: `The number of instance shard.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of an instance, maybe: ` + "`" + `init` + "`" + `, ` + "`" + `processing` + "`" + `, ` + "`" + `online` + "`" + `, ` + "`" + `isolate` + "`" + ` and ` + "`" + `todelete` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of the vpc subnet.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of an instance.`,
				},
				resource.Attribute{
					Name:        "type_id",
					Description: `Instance type. Refer to ` + "`" + `data.tencentcloud_redis_zone_config.list.type_id` + "`" + ` get available values.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the vpc with which the instance is associated.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `Available zone to which a redis instance belongs.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_redis_zone_config",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query which instance types of Redis are available in a specific region.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Name of a region. If this value is not set, the current region getting from provider's configuration will be used.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "type_id",
					Description: `(Optional) Instance type ID. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list of zone. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "mem_sizes",
					Description: `The memory volume of an available instance(in MB).`,
				},
				resource.Attribute{
					Name:        "redis_replicas_nums",
					Description: `The support numbers of instance copies.`,
				},
				resource.Attribute{
					Name:        "redis_shard_nums",
					Description: `The support numbers of instance shard.`,
				},
				resource.Attribute{
					Name:        "type_id",
					Description: `Instance type. Which redis type supports in this zone.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Version description of an available instance. Possible values: ` + "`" + `Redis 3.2` + "`" + `, ` + "`" + `Redis 4.0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `ID of available zone.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list of zone. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "mem_sizes",
					Description: `The memory volume of an available instance(in MB).`,
				},
				resource.Attribute{
					Name:        "redis_replicas_nums",
					Description: `The support numbers of instance copies.`,
				},
				resource.Attribute{
					Name:        "redis_shard_nums",
					Description: `The support numbers of instance shard.`,
				},
				resource.Attribute{
					Name:        "type_id",
					Description: `Instance type. Which redis type supports in this zone.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Version description of an available instance. Possible values: ` + "`" + `Redis 3.2` + "`" + `, ` + "`" + `Redis 4.0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `ID of available zone.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_reserved_instance_configs",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query reserved instances configuration.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) The available zone that the reserved instance locates at.`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `(Optional) Validity period of the reserved instance. Valid values are ` + "`" + `31536000` + "`" + `(1 year) and ` + "`" + `94608000` + "`" + `(3 years).`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Optional) The type of reserved instance.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "config_list",
					Description: `An information list of reserved instance configuration. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Availability zone of the purchasable reserved instance.`,
				},
				resource.Attribute{
					Name:        "config_id",
					Description: `Configuration ID of the purchasable reserved instance.`,
				},
				resource.Attribute{
					Name:        "currency_code",
					Description: `Settlement currency of the reserved instance, which is a standard currency code as listed in ISO 4217.`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `Validity period of the reserved instance.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `Instance type of the reserved instance.`,
				},
				resource.Attribute{
					Name:        "platform",
					Description: `Platform of the reserved instance.`,
				},
				resource.Attribute{
					Name:        "price",
					Description: `Purchase price of the reserved instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "config_list",
					Description: `An information list of reserved instance configuration. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Availability zone of the purchasable reserved instance.`,
				},
				resource.Attribute{
					Name:        "config_id",
					Description: `Configuration ID of the purchasable reserved instance.`,
				},
				resource.Attribute{
					Name:        "currency_code",
					Description: `Settlement currency of the reserved instance, which is a standard currency code as listed in ISO 4217.`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `Validity period of the reserved instance.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `Instance type of the reserved instance.`,
				},
				resource.Attribute{
					Name:        "platform",
					Description: `Platform of the reserved instance.`,
				},
				resource.Attribute{
					Name:        "price",
					Description: `Purchase price of the reserved instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_reserved_instances",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query reserved instances.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) The available zone that the reserved instance locates at.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Optional) The type of reserved instance.`,
				},
				resource.Attribute{
					Name:        "reserved_instance_id",
					Description: `(Optional) ID of the reserved instance to be query.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "reserved_instance_list",
					Description: `An information list of reserved instance. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Availability zone of the reserved instance.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `Expiry time of the reserved instance.`,
				},
				resource.Attribute{
					Name:        "instance_count",
					Description: `Number of reserved instance.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `The type of reserved instance.`,
				},
				resource.Attribute{
					Name:        "reserved_instance_id",
					Description: `ID of the reserved instance.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `Start time of the reserved instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the reserved instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "reserved_instance_list",
					Description: `An information list of reserved instance. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Availability zone of the reserved instance.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `Expiry time of the reserved instance.`,
				},
				resource.Attribute{
					Name:        "instance_count",
					Description: `Number of reserved instance.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `The type of reserved instance.`,
				},
				resource.Attribute{
					Name:        "reserved_instance_id",
					Description: `ID of the reserved instance.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `Start time of the reserved instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the reserved instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_route_table",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Route Table.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "route_table_id",
					Description: `(Required) The Route Table ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The Route Table name. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of routing table.`,
				},
				resource.Attribute{
					Name:        "routes",
					Description: `The information list of the VPC route table.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The RouteEntry's target network segment.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The RouteEntry's description.`,
				},
				resource.Attribute{
					Name:        "next_hub",
					Description: `The RouteEntry's next hub.`,
				},
				resource.Attribute{
					Name:        "next_type",
					Description: `The ` + "`" + `next_hub` + "`" + ` type.`,
				},
				resource.Attribute{
					Name:        "subnet_num",
					Description: `Number of associated subnets.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The VPC ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of routing table.`,
				},
				resource.Attribute{
					Name:        "routes",
					Description: `The information list of the VPC route table.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The RouteEntry's target network segment.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The RouteEntry's description.`,
				},
				resource.Attribute{
					Name:        "next_hub",
					Description: `The RouteEntry's next hub.`,
				},
				resource.Attribute{
					Name:        "next_type",
					Description: `The ` + "`" + `next_hub` + "`" + ` type.`,
				},
				resource.Attribute{
					Name:        "subnet_num",
					Description: `Number of associated subnets.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The VPC ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_scf_functions",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query SCF functions.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the SCF function to be queried.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the SCF function to be queried.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) Namespace of the SCF function to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of the SCF function to be queried, can use up to 10 tags. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "functions",
					Description: `An information list of functions. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "cls_logset_id",
					Description: `CLS logset ID of the SCF function.`,
				},
				resource.Attribute{
					Name:        "cls_topic_id",
					Description: `CLS topic ID of the SCF function.`,
				},
				resource.Attribute{
					Name:        "code_error",
					Description: `Code error of the SCF function.`,
				},
				resource.Attribute{
					Name:        "code_result",
					Description: `Code result of the SCF function.`,
				},
				resource.Attribute{
					Name:        "code_size",
					Description: `Code size of the SCF function.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the SCF function.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the SCF function.`,
				},
				resource.Attribute{
					Name:        "eip_fixed",
					Description: `Whether EIP is a fixed IP.`,
				},
				resource.Attribute{
					Name:        "eips",
					Description: `EIP list of the SCF function.`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `Environment variable of the SCF function.`,
				},
				resource.Attribute{
					Name:        "err_no",
					Description: `Errno of the SCF function.`,
				},
				resource.Attribute{
					Name:        "handler",
					Description: `Handler of the SCF function.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `Host of the SCF function.`,
				},
				resource.Attribute{
					Name:        "install_dependency",
					Description: `Whether to automatically install dependencies.`,
				},
				resource.Attribute{
					Name:        "l5_enable",
					Description: `Whether to enable L5.`,
				},
				resource.Attribute{
					Name:        "mem_size",
					Description: `Memory size of the SCF function runtime, unit is M.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `Modify time of the SCF function.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the SCF function.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `Namespace of the SCF function.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `CAM role of the SCF function.`,
				},
				resource.Attribute{
					Name:        "runtime",
					Description: `Runtime of the SCF function.`,
				},
				resource.Attribute{
					Name:        "status_desc",
					Description: `Status description of the SCF function.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the SCF function.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `Subnet ID of the SCF function.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the SCF function.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `Timeout of the SCF function maximum execution time, unit is second.`,
				},
				resource.Attribute{
					Name:        "trigger_info",
					Description: `Trigger details list the SCF function. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "custom_argument",
					Description: `user-defined parameter of the SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `Whether to enable SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `Modify time of the SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "trigger_desc",
					Description: `TriggerDesc of the SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "vip",
					Description: `Vip of the SCF function.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC ID of the SCF function.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "functions",
					Description: `An information list of functions. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "cls_logset_id",
					Description: `CLS logset ID of the SCF function.`,
				},
				resource.Attribute{
					Name:        "cls_topic_id",
					Description: `CLS topic ID of the SCF function.`,
				},
				resource.Attribute{
					Name:        "code_error",
					Description: `Code error of the SCF function.`,
				},
				resource.Attribute{
					Name:        "code_result",
					Description: `Code result of the SCF function.`,
				},
				resource.Attribute{
					Name:        "code_size",
					Description: `Code size of the SCF function.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the SCF function.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the SCF function.`,
				},
				resource.Attribute{
					Name:        "eip_fixed",
					Description: `Whether EIP is a fixed IP.`,
				},
				resource.Attribute{
					Name:        "eips",
					Description: `EIP list of the SCF function.`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `Environment variable of the SCF function.`,
				},
				resource.Attribute{
					Name:        "err_no",
					Description: `Errno of the SCF function.`,
				},
				resource.Attribute{
					Name:        "handler",
					Description: `Handler of the SCF function.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `Host of the SCF function.`,
				},
				resource.Attribute{
					Name:        "install_dependency",
					Description: `Whether to automatically install dependencies.`,
				},
				resource.Attribute{
					Name:        "l5_enable",
					Description: `Whether to enable L5.`,
				},
				resource.Attribute{
					Name:        "mem_size",
					Description: `Memory size of the SCF function runtime, unit is M.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `Modify time of the SCF function.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the SCF function.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `Namespace of the SCF function.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `CAM role of the SCF function.`,
				},
				resource.Attribute{
					Name:        "runtime",
					Description: `Runtime of the SCF function.`,
				},
				resource.Attribute{
					Name:        "status_desc",
					Description: `Status description of the SCF function.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the SCF function.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `Subnet ID of the SCF function.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the SCF function.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `Timeout of the SCF function maximum execution time, unit is second.`,
				},
				resource.Attribute{
					Name:        "trigger_info",
					Description: `Trigger details list the SCF function. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "custom_argument",
					Description: `user-defined parameter of the SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `Whether to enable SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `Modify time of the SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "trigger_desc",
					Description: `TriggerDesc of the SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "vip",
					Description: `Vip of the SCF function.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC ID of the SCF function.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_scf_logs",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query SCF function logs.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "function_name",
					Description: `(Required) Name of the SCF function to be queried.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `(Optional) The end time of the query, the format is ` + "`" + `2017-05-16 20:00:00` + "`" + `, which can only be within one day from ` + "`" + `start_time` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "invoke_request_id",
					Description: `(Optional) Corresponding requestId when executing function.`,
				},
				resource.Attribute{
					Name:        "limit",
					Description: `(Optional) Number of logs, the default is ` + "`" + `10000` + "`" + `, offset+limit cannot be greater than 10000.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) Namespace of the SCF function to be queried.`,
				},
				resource.Attribute{
					Name:        "offset",
					Description: `(Optional) Log offset, default is ` + "`" + `0` + "`" + `, offset+limit cannot be greater than 10000.`,
				},
				resource.Attribute{
					Name:        "order_by",
					Description: `(Optional) Sort the logs according to the following fields: ` + "`" + `function_name` + "`" + `, ` + "`" + `duration` + "`" + `, ` + "`" + `mem_usage` + "`" + `, ` + "`" + `start_time` + "`" + `, default ` + "`" + `start_time` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Optional) Order to sort the log, optional values ` + "`" + `desc` + "`" + ` and ` + "`" + `asc` + "`" + `, default ` + "`" + `desc` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "ret_code",
					Description: `(Optional) Use to filter log, optional value: ` + "`" + `not0` + "`" + ` only returns the error log. ` + "`" + `is0` + "`" + ` only returns the correct log. ` + "`" + `TimeLimitExceeded` + "`" + ` returns the log of the function call timeout. ` + "`" + `ResourceLimitExceeded` + "`" + ` returns the function call generation resource overrun log. ` + "`" + `UserCodeException` + "`" + ` returns logs of the user code error that occurred in the function call. Not passing the parameter means returning all logs.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `(Optional) The start time of the query, the format is ` + "`" + `2017-05-16 20:00:00` + "`" + `, which can only be within one day from ` + "`" + `end_time` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "logs",
					Description: `An information list of logs. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "bill_duration",
					Description: `Function billing time, according to duration up to the last 100ms, unit is ms.`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `Function execution time-consuming, unit is ms.`,
				},
				resource.Attribute{
					Name:        "function_name",
					Description: `Name of the SCF function.`,
				},
				resource.Attribute{
					Name:        "invoke_finished",
					Description: `Whether the function call ends, ` + "`" + `1` + "`" + ` means the execution ends, other values indicate the call exception.`,
				},
				resource.Attribute{
					Name:        "level",
					Description: `Log level.`,
				},
				resource.Attribute{
					Name:        "log",
					Description: `Log output during function execution.`,
				},
				resource.Attribute{
					Name:        "mem_usage",
					Description: `The actual memory size consumed in the execution of the function, unit is Byte.`,
				},
				resource.Attribute{
					Name:        "request_id",
					Description: `Execute the requestId corresponding to the function.`,
				},
				resource.Attribute{
					Name:        "ret_code",
					Description: `Execution result of function, ` + "`" + `0` + "`" + ` means the execution is successful, other values indicate failure.`,
				},
				resource.Attribute{
					Name:        "ret_msg",
					Description: `Return value after function execution is completed.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `Log source.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `Point in time at which the function begins execution.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "logs",
					Description: `An information list of logs. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "bill_duration",
					Description: `Function billing time, according to duration up to the last 100ms, unit is ms.`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `Function execution time-consuming, unit is ms.`,
				},
				resource.Attribute{
					Name:        "function_name",
					Description: `Name of the SCF function.`,
				},
				resource.Attribute{
					Name:        "invoke_finished",
					Description: `Whether the function call ends, ` + "`" + `1` + "`" + ` means the execution ends, other values indicate the call exception.`,
				},
				resource.Attribute{
					Name:        "level",
					Description: `Log level.`,
				},
				resource.Attribute{
					Name:        "log",
					Description: `Log output during function execution.`,
				},
				resource.Attribute{
					Name:        "mem_usage",
					Description: `The actual memory size consumed in the execution of the function, unit is Byte.`,
				},
				resource.Attribute{
					Name:        "request_id",
					Description: `Execute the requestId corresponding to the function.`,
				},
				resource.Attribute{
					Name:        "ret_code",
					Description: `Execution result of function, ` + "`" + `0` + "`" + ` means the execution is successful, other values indicate failure.`,
				},
				resource.Attribute{
					Name:        "ret_msg",
					Description: `Return value after function execution is completed.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `Log source.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `Point in time at which the function begins execution.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_scf_namespaces",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query SCF namespaces.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the SCF namespace to be queried.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) Name of the SCF namespace to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "namespaces",
					Description: `An information list of namespace. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the SCF namespace.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the SCF namespace.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `Modify time of the SCF namespace.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `Name of the SCF namespace.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the SCF namespace.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "namespaces",
					Description: `An information list of namespace. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the SCF namespace.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the SCF namespace.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `Modify time of the SCF namespace.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `Name of the SCF namespace.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the SCF namespace.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_security_group",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of security group.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the security group to be queried. Conflict with ` + "`" + `security_group_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Optional) ID of the security group to be queried. Conflict with ` + "`" + `name` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "be_associate_count",
					Description: `Number of security group binding resources.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of security group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the security group.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Project ID of the security group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "be_associate_count",
					Description: `Number of security group binding resources.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of security group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the security group.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Project ID of the security group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_security_groups",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of security groups.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the security group to be queried. Conflict with ` + "`" + `security_group_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) Project ID of the security group to be queried. Conflict with ` + "`" + `security_group_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Optional) ID of the security group to be queried. Conflict with ` + "`" + `name` + "`" + ` and ` + "`" + `project_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of the security group to be queried. Conflict with ` + "`" + `security_group_id` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `Information list of security group.`,
				},
				resource.Attribute{
					Name:        "be_associate_count",
					Description: `Number of security group binding resources.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of security group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the security group.`,
				},
				resource.Attribute{
					Name:        "egress",
					Description: `Egress rules set. For items like ` + "`" + `[action]#[cidr_ip]#[port]#[protocol]` + "`" + `, it means a regular rule; for items like ` + "`" + `sg-XXXX` + "`" + `, it means a nested security group.`,
				},
				resource.Attribute{
					Name:        "ingress",
					Description: `Ingress rules set. For items like ` + "`" + `[action]#[cidr_ip]#[port]#[protocol]` + "`" + `, it means a regular rule; for items like ` + "`" + `sg-XXXX` + "`" + `, it means a nested security group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the security group.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Project ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `ID of the security group.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the security group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "security_groups",
					Description: `Information list of security group.`,
				},
				resource.Attribute{
					Name:        "be_associate_count",
					Description: `Number of security group binding resources.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of security group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the security group.`,
				},
				resource.Attribute{
					Name:        "egress",
					Description: `Egress rules set. For items like ` + "`" + `[action]#[cidr_ip]#[port]#[protocol]` + "`" + `, it means a regular rule; for items like ` + "`" + `sg-XXXX` + "`" + `, it means a nested security group.`,
				},
				resource.Attribute{
					Name:        "ingress",
					Description: `Ingress rules set. For items like ` + "`" + `[action]#[cidr_ip]#[port]#[protocol]` + "`" + `, it means a regular rule; for items like ` + "`" + `sg-XXXX` + "`" + `, it means a nested security group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the security group.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Project ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `ID of the security group.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the security group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_sqlserver_account_db_attachments",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query the list of SQL Server account DB privileges.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) SQL Server instance ID that the account belongs to.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `(Optional) Name of the SQL Server account to be queried.`,
				},
				resource.Attribute{
					Name:        "db_name",
					Description: `(Optional) Name of the DB to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to store results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list of SQL Server account. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `SQL Server account name.`,
				},
				resource.Attribute{
					Name:        "db_name",
					Description: `SQL Server DB name.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `SQL Server instance ID that the account belongs to.`,
				},
				resource.Attribute{
					Name:        "privilege",
					Description: `Privilege of the account on DB. Valid value are ` + "`" + `ReadOnly` + "`" + `, ` + "`" + `ReadWrite` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list of SQL Server account. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `SQL Server account name.`,
				},
				resource.Attribute{
					Name:        "db_name",
					Description: `SQL Server DB name.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `SQL Server instance ID that the account belongs to.`,
				},
				resource.Attribute{
					Name:        "privilege",
					Description: `Privilege of the account on DB. Valid value are ` + "`" + `ReadOnly` + "`" + `, ` + "`" + `ReadWrite` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_sqlserver_accounts",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query the list of SQL Server accounts.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) SQL server instance ID that the account belongs to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the SQL server account to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to store results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list of SQL Server account. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the SQL Server account.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `SQL server instance ID that the account belongs to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the SQL server account.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `Remark of the SQL Server account.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the SQL Server account. ` + "`" + `1` + "`" + ` for creating, ` + "`" + `2` + "`" + ` for running, ` + "`" + `3` + "`" + ` for modifying, 4 for resetting password, -1 for deleting.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Last updated time of the SQL Server account.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list of SQL Server account. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the SQL Server account.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `SQL server instance ID that the account belongs to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the SQL server account.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `Remark of the SQL Server account.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the SQL Server account. ` + "`" + `1` + "`" + ` for creating, ` + "`" + `2` + "`" + ` for running, ` + "`" + `3` + "`" + ` for modifying, 4 for resetting password, -1 for deleting.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Last updated time of the SQL Server account.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_sqlserver_backups",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query the list of SQL Server backups.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "end_time",
					Description: `(Required) End time of the instance list, like yyyy-MM-dd HH:mm:ss.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) Instance ID.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `(Required) Start time of the instance list, like yyyy-MM-dd HH:mm:ss.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to store results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list of SQL Server backup. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "db_list",
					Description: `Database name list of the backup.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `End time of the backup.`,
				},
				resource.Attribute{
					Name:        "file_name",
					Description: `File name of the backup.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the backup.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `Instance ID.`,
				},
				resource.Attribute{
					Name:        "internet_url",
					Description: `URL for downloads externally.`,
				},
				resource.Attribute{
					Name:        "intranet_url",
					Description: `URL for downloads internally.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of backup file. Unit is KB.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `Start time of the backup.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the backup. ` + "`" + `1` + "`" + ` for creating, ` + "`" + `2` + "`" + ` for successfully created, 3 for failed.`,
				},
				resource.Attribute{
					Name:        "strategy",
					Description: `Strategy of the backup. ` + "`" + `0` + "`" + ` for instance backup, ` + "`" + `1` + "`" + ` for multi-databases backup.`,
				},
				resource.Attribute{
					Name:        "trigger_model",
					Description: `The way to trigger backup. ` + "`" + `0` + "`" + ` for timed trigger, ` + "`" + `1` + "`" + ` for manual trigger.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list of SQL Server backup. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "db_list",
					Description: `Database name list of the backup.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `End time of the backup.`,
				},
				resource.Attribute{
					Name:        "file_name",
					Description: `File name of the backup.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the backup.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `Instance ID.`,
				},
				resource.Attribute{
					Name:        "internet_url",
					Description: `URL for downloads externally.`,
				},
				resource.Attribute{
					Name:        "intranet_url",
					Description: `URL for downloads internally.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of backup file. Unit is KB.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `Start time of the backup.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the backup. ` + "`" + `1` + "`" + ` for creating, ` + "`" + `2` + "`" + ` for successfully created, 3 for failed.`,
				},
				resource.Attribute{
					Name:        "strategy",
					Description: `Strategy of the backup. ` + "`" + `0` + "`" + ` for instance backup, ` + "`" + `1` + "`" + ` for multi-databases backup.`,
				},
				resource.Attribute{
					Name:        "trigger_model",
					Description: `The way to trigger backup. ` + "`" + `0` + "`" + ` for timed trigger, ` + "`" + `1` + "`" + ` for manual trigger.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_sqlserver_basic_instances",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query SQL Server basic instances`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the SQL Server basic instance to be query.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) Project ID of the SQL Server basic instance to be query.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) Subnet ID of the SQL Server basic instance to be query.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) Vpc ID of the SQL Server basic instance to be query. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "instance_list",
					Description: `A list of SQL Server basic instances. Each element contains the following attributes.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Availability zone.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `Pay type of the SQL Server basic instance. For now, only ` + "`" + `POSTPAID_BY_HOUR` + "`" + ` is valid.`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `The CPU number of the SQL Server basic instance.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the SQL Server basic instance.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `Version of the SQL Server basic database engine. Allowed values are ` + "`" + `2008R2` + "`" + `(SQL Server 2008 Enterprise), ` + "`" + `2012SP3` + "`" + `(SQL Server 2012 Enterprise), ` + "`" + `2016SP1` + "`" + ` (SQL Server 2016 Enterprise), ` + "`" + `201602` + "`" + `(SQL Server 2016 Standard) and ` + "`" + `2017` + "`" + `(SQL Server 2017 Enterprise). Default is ` + "`" + `2008R2` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the SQL Server basic instance.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `Memory size (in GB). Allowed value must be larger than ` + "`" + `memory` + "`" + ` that data source ` + "`" + `tencentcloud_sqlserver_specinfos` + "`" + ` provides.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the SQL Server basic instance.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Project ID, default value is ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the SQL Server basic instance. ` + "`" + `1` + "`" + ` for applying, ` + "`" + `2` + "`" + ` for running, ` + "`" + `3` + "`" + ` for running with limit, ` + "`" + `4` + "`" + ` for isolated, ` + "`" + `5` + "`" + ` for recycling, ` + "`" + `6` + "`" + ` for recycled, ` + "`" + `7` + "`" + ` for running with task, ` + "`" + `8` + "`" + ` for off-line, ` + "`" + `9` + "`" + ` for expanding, ` + "`" + `10` + "`" + ` for migrating, ` + "`" + `11` + "`" + ` for readonly, ` + "`" + `12` + "`" + ` for rebooting.`,
				},
				resource.Attribute{
					Name:        "storage",
					Description: `Disk size (in GB). Allowed value must be a multiple of 10. The storage must be set with the limit of ` + "`" + `storage_min` + "`" + ` and ` + "`" + `storage_max` + "`" + ` which data source ` + "`" + `tencentcloud_sqlserver_specinfos` + "`" + ` provides.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of subnet.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the SQL Server basic instance.`,
				},
				resource.Attribute{
					Name:        "used_storage",
					Description: `Used storage.`,
				},
				resource.Attribute{
					Name:        "vip",
					Description: `IP for private access.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of VPC.`,
				},
				resource.Attribute{
					Name:        "vport",
					Description: `Port for private access.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_list",
					Description: `A list of SQL Server basic instances. Each element contains the following attributes.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Availability zone.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `Pay type of the SQL Server basic instance. For now, only ` + "`" + `POSTPAID_BY_HOUR` + "`" + ` is valid.`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `The CPU number of the SQL Server basic instance.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the SQL Server basic instance.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `Version of the SQL Server basic database engine. Allowed values are ` + "`" + `2008R2` + "`" + `(SQL Server 2008 Enterprise), ` + "`" + `2012SP3` + "`" + `(SQL Server 2012 Enterprise), ` + "`" + `2016SP1` + "`" + ` (SQL Server 2016 Enterprise), ` + "`" + `201602` + "`" + `(SQL Server 2016 Standard) and ` + "`" + `2017` + "`" + `(SQL Server 2017 Enterprise). Default is ` + "`" + `2008R2` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the SQL Server basic instance.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `Memory size (in GB). Allowed value must be larger than ` + "`" + `memory` + "`" + ` that data source ` + "`" + `tencentcloud_sqlserver_specinfos` + "`" + ` provides.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the SQL Server basic instance.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Project ID, default value is ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the SQL Server basic instance. ` + "`" + `1` + "`" + ` for applying, ` + "`" + `2` + "`" + ` for running, ` + "`" + `3` + "`" + ` for running with limit, ` + "`" + `4` + "`" + ` for isolated, ` + "`" + `5` + "`" + ` for recycling, ` + "`" + `6` + "`" + ` for recycled, ` + "`" + `7` + "`" + ` for running with task, ` + "`" + `8` + "`" + ` for off-line, ` + "`" + `9` + "`" + ` for expanding, ` + "`" + `10` + "`" + ` for migrating, ` + "`" + `11` + "`" + ` for readonly, ` + "`" + `12` + "`" + ` for rebooting.`,
				},
				resource.Attribute{
					Name:        "storage",
					Description: `Disk size (in GB). Allowed value must be a multiple of 10. The storage must be set with the limit of ` + "`" + `storage_min` + "`" + ` and ` + "`" + `storage_max` + "`" + ` which data source ` + "`" + `tencentcloud_sqlserver_specinfos` + "`" + ` provides.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of subnet.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the SQL Server basic instance.`,
				},
				resource.Attribute{
					Name:        "used_storage",
					Description: `Used storage.`,
				},
				resource.Attribute{
					Name:        "vip",
					Description: `IP for private access.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of VPC.`,
				},
				resource.Attribute{
					Name:        "vport",
					Description: `Port for private access.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_sqlserver_dbs",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query DB resources for the specific SQL Server instance.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) SQL Server instance ID which DB belongs to.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to store results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "db_list",
					Description: `A list of dbs belong to the specific instance. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "charset",
					Description: `Character set DB uses, could be ` + "`" + `Chinese_PRC_CI_AS` + "`" + `, ` + "`" + `Chinese_PRC_CS_AS` + "`" + `, ` + "`" + `Chinese_PRC_BIN` + "`" + `, ` + "`" + `Chinese_Taiwan_Stroke_CI_AS` + "`" + `, ` + "`" + `SQL_Latin1_General_CP1_CI_AS` + "`" + `, and ` + "`" + `SQL_Latin1_General_CP1_CS_AS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Database creation time.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `SQL Server instance ID which DB belongs to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of DB.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `Remark of the DB.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Database status. Valid values are ` + "`" + `creating` + "`" + `, ` + "`" + `running` + "`" + `, ` + "`" + `modifying` + "`" + `, ` + "`" + `dropping` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "db_list",
					Description: `A list of dbs belong to the specific instance. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "charset",
					Description: `Character set DB uses, could be ` + "`" + `Chinese_PRC_CI_AS` + "`" + `, ` + "`" + `Chinese_PRC_CS_AS` + "`" + `, ` + "`" + `Chinese_PRC_BIN` + "`" + `, ` + "`" + `Chinese_Taiwan_Stroke_CI_AS` + "`" + `, ` + "`" + `SQL_Latin1_General_CP1_CI_AS` + "`" + `, and ` + "`" + `SQL_Latin1_General_CP1_CS_AS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Database creation time.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `SQL Server instance ID which DB belongs to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of DB.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `Remark of the DB.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Database status. Valid values are ` + "`" + `creating` + "`" + `, ` + "`" + `running` + "`" + `, ` + "`" + `modifying` + "`" + `, ` + "`" + `dropping` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_sqlserver_instances",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query SQL Server instances`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the SQL Server instance to be query.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) Project ID of the SQL Server instance to be query.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) Subnet ID of the SQL Server instance to be query.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) Vpc ID of the SQL Server instance to be query. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "instance_list",
					Description: `A list of SQL Server instances. Each element contains the following attributes.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Availability zone.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `Pay type of the SQL Server instance. For now, only ` + "`" + `POSTPAID_BY_HOUR` + "`" + ` is valid.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the SQL Server instance.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `Version of the SQL Server database engine. Allowed values are ` + "`" + `2008R2` + "`" + `(SQL Server 2008 Enterprise), ` + "`" + `2012SP3` + "`" + `(SQL Server 2012 Enterprise), ` + "`" + `2016SP1` + "`" + ` (SQL Server 2016 Enterprise), ` + "`" + `201602` + "`" + `(SQL Server 2016 Standard) and ` + "`" + `2017` + "`" + `(SQL Server 2017 Enterprise). Default is ` + "`" + `2008R2` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ha_type",
					Description: `Instance type. ` + "`" + `DUAL` + "`" + ` (dual-server high availability), ` + "`" + `CLUSTER` + "`" + ` (cluster).`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the SQL Server instance.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `Memory size (in GB). Allowed value must be larger than ` + "`" + `memory` + "`" + ` that data source ` + "`" + `tencentcloud_sqlserver_specinfos` + "`" + ` provides.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the SQL Server instance.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Project ID, default value is 0.`,
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
					Name:        "storage",
					Description: `Disk size (in GB). Allowed value must be a multiple of 10. The storage must be set with the limit of ` + "`" + `storage_min` + "`" + ` and ` + "`" + `storage_max` + "`" + ` which data source ` + "`" + `tencentcloud_sqlserver_specinfos` + "`" + ` provides.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of subnet.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the SQL Server instance.`,
				},
				resource.Attribute{
					Name:        "used_storage",
					Description: `Used storage.`,
				},
				resource.Attribute{
					Name:        "vip",
					Description: `IP for private access.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of VPC.`,
				},
				resource.Attribute{
					Name:        "vport",
					Description: `Port for private access.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_list",
					Description: `A list of SQL Server instances. Each element contains the following attributes.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Availability zone.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `Pay type of the SQL Server instance. For now, only ` + "`" + `POSTPAID_BY_HOUR` + "`" + ` is valid.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the SQL Server instance.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `Version of the SQL Server database engine. Allowed values are ` + "`" + `2008R2` + "`" + `(SQL Server 2008 Enterprise), ` + "`" + `2012SP3` + "`" + `(SQL Server 2012 Enterprise), ` + "`" + `2016SP1` + "`" + ` (SQL Server 2016 Enterprise), ` + "`" + `201602` + "`" + `(SQL Server 2016 Standard) and ` + "`" + `2017` + "`" + `(SQL Server 2017 Enterprise). Default is ` + "`" + `2008R2` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ha_type",
					Description: `Instance type. ` + "`" + `DUAL` + "`" + ` (dual-server high availability), ` + "`" + `CLUSTER` + "`" + ` (cluster).`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the SQL Server instance.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `Memory size (in GB). Allowed value must be larger than ` + "`" + `memory` + "`" + ` that data source ` + "`" + `tencentcloud_sqlserver_specinfos` + "`" + ` provides.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the SQL Server instance.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Project ID, default value is 0.`,
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
					Name:        "storage",
					Description: `Disk size (in GB). Allowed value must be a multiple of 10. The storage must be set with the limit of ` + "`" + `storage_min` + "`" + ` and ` + "`" + `storage_max` + "`" + ` which data source ` + "`" + `tencentcloud_sqlserver_specinfos` + "`" + ` provides.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of subnet.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the SQL Server instance.`,
				},
				resource.Attribute{
					Name:        "used_storage",
					Description: `Used storage.`,
				},
				resource.Attribute{
					Name:        "vip",
					Description: `IP for private access.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of VPC.`,
				},
				resource.Attribute{
					Name:        "vport",
					Description: `Port for private access.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_sqlserver_publish_subscribes",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query Publish Subscribe resources for the specific SQL Server instance.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) ID of the SQL Server instance.`,
				},
				resource.Attribute{
					Name:        "pub_or_sub_instance_id",
					Description: `(Optional) The subscribe/publish instance ID. It is related to whether the ` + "`" + `instance_id` + "`" + ` is a publish instance or a subscribe instance. when ` + "`" + `instance_id` + "`" + ` is a publish instance, this field is filtered according to the subscribe instance ID; when ` + "`" + `instance_id` + "`" + ` is a subscribe instance, this field is filtering according to the publish instance ID.`,
				},
				resource.Attribute{
					Name:        "pub_or_sub_instance_ip",
					Description: `(Optional) The intranet IP of the subscribe/publish instance. It is related to whether the ` + "`" + `instance_id` + "`" + ` is a publish instance or a subscribe instance. when ` + "`" + `instance_id` + "`" + ` is a publish instance, this field is filtered according to the intranet IP of the subscribe instance; when ` + "`" + `instance_id` + "`" + ` is a subscribe instance, this field is based on the publish instance intranet IP filter.`,
				},
				resource.Attribute{
					Name:        "publish_database",
					Description: `(Optional) Name of publish database.`,
				},
				resource.Attribute{
					Name:        "publish_subscribe_id",
					Description: `(Optional) The id of the Publish and Subscribe.`,
				},
				resource.Attribute{
					Name:        "publish_subscribe_name",
					Description: `(Optional) The name of the Publish and Subscribe.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to store results.`,
				},
				resource.Attribute{
					Name:        "subscribe_database",
					Description: `(Optional) Name of subscribe database. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "publish_subscribe_list",
					Description: `Publish and subscribe list. Each element contains the following attributes.`,
				},
				resource.Attribute{
					Name:        "database_tuples",
					Description: `Database Publish and Publish relationship list.`,
				},
				resource.Attribute{
					Name:        "last_sync_time",
					Description: `Last sync time.`,
				},
				resource.Attribute{
					Name:        "publish_database",
					Description: `Name of the publish SQL Server instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Publish and subscribe status between databases, valid values are ` + "`" + `running` + "`" + `, ` + "`" + `success` + "`" + `, ` + "`" + `fail` + "`" + `, ` + "`" + `unknow` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "subscribe_database",
					Description: `Name of the subscribe SQL Server instance.`,
				},
				resource.Attribute{
					Name:        "publish_instance_id",
					Description: `ID of the SQL Server instance which publish.`,
				},
				resource.Attribute{
					Name:        "publish_instance_ip",
					Description: `IP of the the SQL Server instance which publish.`,
				},
				resource.Attribute{
					Name:        "publish_instance_name",
					Description: `Name of the SQL Server instance which publish.`,
				},
				resource.Attribute{
					Name:        "publish_subscribe_id",
					Description: `The id of the Publish and Subscribe.`,
				},
				resource.Attribute{
					Name:        "publish_subscribe_name",
					Description: `The name of the Publish and Subscribe.`,
				},
				resource.Attribute{
					Name:        "subscribe_instance_id",
					Description: `ID of the SQL Server instance which subscribe.`,
				},
				resource.Attribute{
					Name:        "subscribe_instance_ip",
					Description: `IP of the SQL Server instance which subscribe.`,
				},
				resource.Attribute{
					Name:        "subscribe_instance_name",
					Description: `Name of the SQL Server instance which subscribe.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "publish_subscribe_list",
					Description: `Publish and subscribe list. Each element contains the following attributes.`,
				},
				resource.Attribute{
					Name:        "database_tuples",
					Description: `Database Publish and Publish relationship list.`,
				},
				resource.Attribute{
					Name:        "last_sync_time",
					Description: `Last sync time.`,
				},
				resource.Attribute{
					Name:        "publish_database",
					Description: `Name of the publish SQL Server instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Publish and subscribe status between databases, valid values are ` + "`" + `running` + "`" + `, ` + "`" + `success` + "`" + `, ` + "`" + `fail` + "`" + `, ` + "`" + `unknow` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "subscribe_database",
					Description: `Name of the subscribe SQL Server instance.`,
				},
				resource.Attribute{
					Name:        "publish_instance_id",
					Description: `ID of the SQL Server instance which publish.`,
				},
				resource.Attribute{
					Name:        "publish_instance_ip",
					Description: `IP of the the SQL Server instance which publish.`,
				},
				resource.Attribute{
					Name:        "publish_instance_name",
					Description: `Name of the SQL Server instance which publish.`,
				},
				resource.Attribute{
					Name:        "publish_subscribe_id",
					Description: `The id of the Publish and Subscribe.`,
				},
				resource.Attribute{
					Name:        "publish_subscribe_name",
					Description: `The name of the Publish and Subscribe.`,
				},
				resource.Attribute{
					Name:        "subscribe_instance_id",
					Description: `ID of the SQL Server instance which subscribe.`,
				},
				resource.Attribute{
					Name:        "subscribe_instance_ip",
					Description: `IP of the SQL Server instance which subscribe.`,
				},
				resource.Attribute{
					Name:        "subscribe_instance_name",
					Description: `Name of the SQL Server instance which subscribe.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_sqlserver_readonly_groups",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query the list of SQL Server readonly groups.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "master_instance_id",
					Description: `(Optional) Master SQL Server instance ID.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to store results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list of SQL Server readonly group. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the readonly group.`,
				},
				resource.Attribute{
					Name:        "is_offline_delay",
					Description: `Indicate whether to offline delayed readonly instances.`,
				},
				resource.Attribute{
					Name:        "master_instance_id",
					Description: `Master instance id.`,
				},
				resource.Attribute{
					Name:        "max_delay_time",
					Description: `Maximum delay time of the readonly instances.`,
				},
				resource.Attribute{
					Name:        "min_instances",
					Description: `Minimum readonly instances that stays in the group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the readonly group.`,
				},
				resource.Attribute{
					Name:        "readonly_instance_set",
					Description: `Readonly instance ID set of the readonly group.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the readonly group. ` + "`" + `1` + "`" + ` for running, ` + "`" + `5` + "`" + ` for applying.`,
				},
				resource.Attribute{
					Name:        "vip",
					Description: `Virtual IP address of the readonly group.`,
				},
				resource.Attribute{
					Name:        "vport",
					Description: `Virtual port of the readonly group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list of SQL Server readonly group. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the readonly group.`,
				},
				resource.Attribute{
					Name:        "is_offline_delay",
					Description: `Indicate whether to offline delayed readonly instances.`,
				},
				resource.Attribute{
					Name:        "master_instance_id",
					Description: `Master instance id.`,
				},
				resource.Attribute{
					Name:        "max_delay_time",
					Description: `Maximum delay time of the readonly instances.`,
				},
				resource.Attribute{
					Name:        "min_instances",
					Description: `Minimum readonly instances that stays in the group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the readonly group.`,
				},
				resource.Attribute{
					Name:        "readonly_instance_set",
					Description: `Readonly instance ID set of the readonly group.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the readonly group. ` + "`" + `1` + "`" + ` for running, ` + "`" + `5` + "`" + ` for applying.`,
				},
				resource.Attribute{
					Name:        "vip",
					Description: `Virtual IP address of the readonly group.`,
				},
				resource.Attribute{
					Name:        "vport",
					Description: `Virtual port of the readonly group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_sqlserver_zone_config",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query purchasable specification configuration for each availability zone in this specific region.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to store results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "zone_list",
					Description: `A list of availability zones. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Alphabet ID of availability zone.`,
				},
				resource.Attribute{
					Name:        "specinfo_list",
					Description: `A list of specinfo configurations for the specific availability zone. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `Billing mode under this specification. Valid values are ` + "`" + `POSTPAID_BY_HOUR` + "`" + `, ` + "`" + `PREPAID` + "`" + ` and ` + "`" + `ALL` + "`" + `. ` + "`" + `ALL` + "`" + ` means both POSTPAID_BY_HOUR and PREPAID.`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `Number of CPU cores.`,
				},
				resource.Attribute{
					Name:        "db_version_name",
					Description: `Version name corresponding to the ` + "`" + `db_version` + "`" + ` field.`,
				},
				resource.Attribute{
					Name:        "db_version",
					Description: `Database version information. Valid values: ` + "`" + `2008R2 (SQL Server 2008 Enterprise)` + "`" + `, ` + "`" + `2012SP3 (SQL Server 2012 Enterprise)` + "`" + `, ` + "`" + `2016SP1 (SQL Server 2016 Enterprise)` + "`" + `, ` + "`" + `201602 (SQL Server 2016 Standard)` + "`" + `, ` + "`" + `2017 (SQL Server 2017 Enterprise)` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "machine_type",
					Description: `Model ID.`,
				},
				resource.Attribute{
					Name:        "max_storage_size",
					Description: `Maximum disk size under this specification in GB.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `Memory size in GB.`,
				},
				resource.Attribute{
					Name:        "min_storage_size",
					Description: `Minimum disk size under this specification in GB.`,
				},
				resource.Attribute{
					Name:        "qps",
					Description: `QPS of this specification.`,
				},
				resource.Attribute{
					Name:        "spec_id",
					Description: `Instance specification ID.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `Number ID of availability zone.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_list",
					Description: `A list of availability zones. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Alphabet ID of availability zone.`,
				},
				resource.Attribute{
					Name:        "specinfo_list",
					Description: `A list of specinfo configurations for the specific availability zone. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `Billing mode under this specification. Valid values are ` + "`" + `POSTPAID_BY_HOUR` + "`" + `, ` + "`" + `PREPAID` + "`" + ` and ` + "`" + `ALL` + "`" + `. ` + "`" + `ALL` + "`" + ` means both POSTPAID_BY_HOUR and PREPAID.`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `Number of CPU cores.`,
				},
				resource.Attribute{
					Name:        "db_version_name",
					Description: `Version name corresponding to the ` + "`" + `db_version` + "`" + ` field.`,
				},
				resource.Attribute{
					Name:        "db_version",
					Description: `Database version information. Valid values: ` + "`" + `2008R2 (SQL Server 2008 Enterprise)` + "`" + `, ` + "`" + `2012SP3 (SQL Server 2012 Enterprise)` + "`" + `, ` + "`" + `2016SP1 (SQL Server 2016 Enterprise)` + "`" + `, ` + "`" + `201602 (SQL Server 2016 Standard)` + "`" + `, ` + "`" + `2017 (SQL Server 2017 Enterprise)` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "machine_type",
					Description: `Model ID.`,
				},
				resource.Attribute{
					Name:        "max_storage_size",
					Description: `Maximum disk size under this specification in GB.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `Memory size in GB.`,
				},
				resource.Attribute{
					Name:        "min_storage_size",
					Description: `Minimum disk size under this specification in GB.`,
				},
				resource.Attribute{
					Name:        "qps",
					Description: `QPS of this specification.`,
				},
				resource.Attribute{
					Name:        "spec_id",
					Description: `Instance specification ID.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `Number ID of availability zone.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_ssl_certificates",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query SSL certificate.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the SSL certificate to be queried.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the SSL certificate to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type of the SSL certificate to be queried. Available values includes: ` + "`" + `CA` + "`" + ` and ` + "`" + `SVR` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "certificates",
					Description: `An information list of certificate. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "begin_time",
					Description: `Beginning time of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "cert",
					Description: `Content of the SSL certificate.`,
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
					Name:        "id",
					Description: `ID of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "product_zh_name",
					Description: `Certificate authority.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Project ID of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "subject_names",
					Description: `ALL domains included in the SSL certificate. Including the primary domain name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the SSL certificate.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "certificates",
					Description: `An information list of certificate. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "begin_time",
					Description: `Beginning time of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "cert",
					Description: `Content of the SSL certificate.`,
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
					Name:        "id",
					Description: `ID of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "product_zh_name",
					Description: `Certificate authority.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Project ID of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "subject_names",
					Description: `ALL domains included in the SSL certificate. Including the primary domain name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the SSL certificate.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_ssm_secret_versions",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of SSM secret version`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "secret_name",
					Description: `(Required) Secret name used to filter result.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `(Optional) VersionId used to filter result. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "secret_version_list",
					Description: `A list of SSM secret versions. When secret status is ` + "`" + `Disabled` + "`" + `, this field will not update anymore.`,
				},
				resource.Attribute{
					Name:        "secret_binary",
					Description: `The base64-encoded binary secret.`,
				},
				resource.Attribute{
					Name:        "secret_string",
					Description: `The string text of secret.`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `Version of secret.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "secret_version_list",
					Description: `A list of SSM secret versions. When secret status is ` + "`" + `Disabled` + "`" + `, this field will not update anymore.`,
				},
				resource.Attribute{
					Name:        "secret_binary",
					Description: `The base64-encoded binary secret.`,
				},
				resource.Attribute{
					Name:        "secret_string",
					Description: `The string text of secret.`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `Version of secret.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_ssm_secrets",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of SSM secret`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "order_type",
					Description: `(Optional) The order to sort the create time of secret. ` + "`" + `0` + "`" + ` - desc, ` + "`" + `1` + "`" + ` - asc. Default value is ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "secret_name",
					Description: `(Optional) Secret name used to filter result.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) Filter by state of secret. ` + "`" + `0` + "`" + ` - all secrets are queried, ` + "`" + `1` + "`" + ` - only Enabled secrets are queried, ` + "`" + `2` + "`" + ` - only Disabled secrets are queried, ` + "`" + `3` + "`" + ` - only PendingDelete secrets are queried.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags to filter secret. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "secret_list",
					Description: `A list of SSM secrets.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of secret.`,
				},
				resource.Attribute{
					Name:        "create_uin",
					Description: `Uin of Creator.`,
				},
				resource.Attribute{
					Name:        "delete_time",
					Description: `Delete time of CMK.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of secret.`,
				},
				resource.Attribute{
					Name:        "kms_key_id",
					Description: `KMS keyId used to encrypt secret.`,
				},
				resource.Attribute{
					Name:        "secret_name",
					Description: `Name of secret.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of secret.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "secret_list",
					Description: `A list of SSM secrets.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of secret.`,
				},
				resource.Attribute{
					Name:        "create_uin",
					Description: `Uin of Creator.`,
				},
				resource.Attribute{
					Name:        "delete_time",
					Description: `Delete time of CMK.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of secret.`,
				},
				resource.Attribute{
					Name:        "kms_key_id",
					Description: `KMS keyId used to encrypt secret.`,
				},
				resource.Attribute{
					Name:        "secret_name",
					Description: `Name of secret.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of secret.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_subnet",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific VPC subnet.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) The ID of the Subnet.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) The VPC ID. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The AZ for the subnet.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The CIDR block of the Subnet.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name for the Subnet.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The Route Table ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The AZ for the subnet.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The CIDR block of the Subnet.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name for the Subnet.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The Route Table ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_tcaplus_clusters",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query TcaplusDB clusters.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Optional) ID of the TcaplusDB cluster to be query.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Optional) Name of the TcaplusDB cluster to be query.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) File for saving results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list of TcaplusDB cluster. Each element contains the following attributes.`,
				},
				resource.Attribute{
					Name:        "api_access_id",
					Description: `Access id of the TcaplusDB cluster.For TcaplusDB SDK connect.`,
				},
				resource.Attribute{
					Name:        "api_access_ip",
					Description: `Access ip of the TcaplusDB cluster.For TcaplusDB SDK connect.`,
				},
				resource.Attribute{
					Name:        "api_access_port",
					Description: `Access port of the TcaplusDB cluster.For TcaplusDB SDK connect.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `ID of the TcaplusDB cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `Name of the TcaplusDB cluster.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the TcaplusDB cluster.`,
				},
				resource.Attribute{
					Name:        "idl_type",
					Description: `IDL type of the TcaplusDB cluster.`,
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
					Description: `Password status of the TcaplusDB cluster. Valid values: ` + "`" + `unmodifiable` + "`" + `, ` + "`" + `modifiable` + "`" + `. ` + "`" + `unmodifiable` + "`" + ` means the password can not be changed in this moment; ` + "`" + `modifiable` + "`" + ` means the password can be changed in this moment.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Access password of the TcaplusDB cluster.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `Subnet ID of the TcaplusDB cluster.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC ID of the TcaplusDB cluster.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list of TcaplusDB cluster. Each element contains the following attributes.`,
				},
				resource.Attribute{
					Name:        "api_access_id",
					Description: `Access id of the TcaplusDB cluster.For TcaplusDB SDK connect.`,
				},
				resource.Attribute{
					Name:        "api_access_ip",
					Description: `Access ip of the TcaplusDB cluster.For TcaplusDB SDK connect.`,
				},
				resource.Attribute{
					Name:        "api_access_port",
					Description: `Access port of the TcaplusDB cluster.For TcaplusDB SDK connect.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `ID of the TcaplusDB cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `Name of the TcaplusDB cluster.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the TcaplusDB cluster.`,
				},
				resource.Attribute{
					Name:        "idl_type",
					Description: `IDL type of the TcaplusDB cluster.`,
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
					Description: `Password status of the TcaplusDB cluster. Valid values: ` + "`" + `unmodifiable` + "`" + `, ` + "`" + `modifiable` + "`" + `. ` + "`" + `unmodifiable` + "`" + ` means the password can not be changed in this moment; ` + "`" + `modifiable` + "`" + ` means the password can be changed in this moment.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Access password of the TcaplusDB cluster.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `Subnet ID of the TcaplusDB cluster.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC ID of the TcaplusDB cluster.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_tcaplus_idls",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query IDL information of the TcaplusDB table.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) ID of the TcaplusDB cluster to be query.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) File for saving results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list of TcaplusDB table IDL. Each element contains the following attributes.`,
				},
				resource.Attribute{
					Name:        "idl_id",
					Description: `ID of the IDL.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list of TcaplusDB table IDL. Each element contains the following attributes.`,
				},
				resource.Attribute{
					Name:        "idl_id",
					Description: `ID of the IDL.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_tcaplus_tablegroups",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query table groups of the TcaplusDB cluster.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) Id of the TcaplusDB cluster to be query.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) File for saving results.`,
				},
				resource.Attribute{
					Name:        "tablegroup_id",
					Description: `(Optional) Id of the table group to be query.`,
				},
				resource.Attribute{
					Name:        "tablegroup_name",
					Description: `(Optional) Name of the table group to be query. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list of table group. Each element contains the following attributes.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the table group..`,
				},
				resource.Attribute{
					Name:        "table_count",
					Description: `Number of tables.`,
				},
				resource.Attribute{
					Name:        "tablegroup_id",
					Description: `Id of the table group.`,
				},
				resource.Attribute{
					Name:        "tablegroup_name",
					Description: `Name of the table group.`,
				},
				resource.Attribute{
					Name:        "total_size",
					Description: `Total storage size (MB).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list of table group. Each element contains the following attributes.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the table group..`,
				},
				resource.Attribute{
					Name:        "table_count",
					Description: `Number of tables.`,
				},
				resource.Attribute{
					Name:        "tablegroup_id",
					Description: `Id of the table group.`,
				},
				resource.Attribute{
					Name:        "tablegroup_name",
					Description: `Name of the table group.`,
				},
				resource.Attribute{
					Name:        "total_size",
					Description: `Total storage size (MB).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_tcaplus_tables",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query TcaplusDB tables.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) ID of the TcaplusDB cluster to be query.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) File for saving results.`,
				},
				resource.Attribute{
					Name:        "table_id",
					Description: `(Optional) Table ID to be query.`,
				},
				resource.Attribute{
					Name:        "table_name",
					Description: `(Optional) Table name to be query.`,
				},
				resource.Attribute{
					Name:        "tablegroup_id",
					Description: `(Optional) ID of the table group to be query. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list of TcaplusDB tables. Each element contains the following attributes.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "error",
					Description: `Error message for creating TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "idl_id",
					Description: `IDL file id of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "reserved_read_cu",
					Description: `Reserved read capacity units of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "reserved_volume",
					Description: `Reserved storage capacity of the TcaplusDB table (unit:GB).`,
				},
				resource.Attribute{
					Name:        "reserved_write_cu",
					Description: `Reserved write capacity units of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "table_id",
					Description: `ID of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "table_idl_type",
					Description: `IDL type of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "table_name",
					Description: `Name of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "table_size",
					Description: `Size of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "table_type",
					Description: `Type of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "tablegroup_id",
					Description: `Table group id of the TcaplusDB table.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list of TcaplusDB tables. Each element contains the following attributes.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "error",
					Description: `Error message for creating TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "idl_id",
					Description: `IDL file id of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "reserved_read_cu",
					Description: `Reserved read capacity units of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "reserved_volume",
					Description: `Reserved storage capacity of the TcaplusDB table (unit:GB).`,
				},
				resource.Attribute{
					Name:        "reserved_write_cu",
					Description: `Reserved write capacity units of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "table_id",
					Description: `ID of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "table_idl_type",
					Description: `IDL type of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "table_name",
					Description: `Name of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "table_size",
					Description: `Size of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "table_type",
					Description: `Type of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "tablegroup_id",
					Description: `Table group id of the TcaplusDB table.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_tcr_instances",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of TCR instances.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Optional) ID of the TCR instance to query.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the TCR instance to query.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "instance_list",
					Description: `Information list of the dedicated TCR instances.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the TCR instance.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `Instance type.`,
				},
				resource.Attribute{
					Name:        "internal_end_point",
					Description: `Internal address for access of the TCR instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of TCR instance.`,
				},
				resource.Attribute{
					Name:        "public_domain",
					Description: `Public address for access of the TCR instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the TCR instance.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the TCR instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_list",
					Description: `Information list of the dedicated TCR instances.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the TCR instance.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `Instance type.`,
				},
				resource.Attribute{
					Name:        "internal_end_point",
					Description: `Internal address for access of the TCR instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of TCR instance.`,
				},
				resource.Attribute{
					Name:        "public_domain",
					Description: `Public address for access of the TCR instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the TCR instance.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the TCR instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_tcr_namespaces",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of TCR namespaces.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) ID of the instance that the namespace belongs to.`,
				},
				resource.Attribute{
					Name:        "namespace_name",
					Description: `(Optional) ID of the TCR namespace to query.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "namespace_list",
					Description: `Information list of the dedicated TCR namespaces.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `Indicate that the namespace is public or not.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of TCR namespace.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace_list",
					Description: `Information list of the dedicated TCR namespaces.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `Indicate that the namespace is public or not.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of TCR namespace.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_tcr_repositories",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of TCR repositories.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) ID of the TCR instance that the repository belongs to.`,
				},
				resource.Attribute{
					Name:        "namespace_name",
					Description: `(Required) Name of the namespace that the repository belongs to.`,
				},
				resource.Attribute{
					Name:        "repository_name",
					Description: `(Optional) ID of the TCR repositories to query.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "repository_list",
					Description: `Information list of the dedicated TCR repositories.`,
				},
				resource.Attribute{
					Name:        "brief_desc",
					Description: `Brief description of the repository.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the repository.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `Indicate that the repository is public or not.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of repository.`,
				},
				resource.Attribute{
					Name:        "namespace_name",
					Description: `Name of the namespace that the repository belongs to.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Last update time.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `URL of the repository.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "repository_list",
					Description: `Information list of the dedicated TCR repositories.`,
				},
				resource.Attribute{
					Name:        "brief_desc",
					Description: `Brief description of the repository.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the repository.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `Indicate that the repository is public or not.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of repository.`,
				},
				resource.Attribute{
					Name:        "namespace_name",
					Description: `Name of the namespace that the repository belongs to.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Last update time.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `URL of the repository.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_tcr_tokens",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of TCR tokens.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) ID of the instance that the token belongs to.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "token_id",
					Description: `(Optional) ID of the TCR token to query. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "token_list",
					Description: `Information list of the dedicated TCR tokens.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the token.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `Indicate that the token is enabled or not.`,
				},
				resource.Attribute{
					Name:        "token_id",
					Description: `Id of TCR token.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "token_list",
					Description: `Information list of the dedicated TCR tokens.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the token.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `Indicate that the token is enabled or not.`,
				},
				resource.Attribute{
					Name:        "token_id",
					Description: `Id of TCR token.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_tcr_vpc_attachments",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of TCR VPC attachment.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) ID of the instance to query.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) ID of subnet to query.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) ID of VPC to query. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "vpc_attachment_list",
					Description: `Information list of the dedicated TCR namespaces.`,
				},
				resource.Attribute{
					Name:        "access_ip",
					Description: `IP address of this VPC access.`,
				},
				resource.Attribute{
					Name:        "enable_public_domain_dns",
					Description: `Whether to enable public domain dns.`,
				},
				resource.Attribute{
					Name:        "enable_vpc_domain_dns",
					Description: `Whether to enable vpc domain dns.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of this VPC access.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of subnet.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of VPC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vpc_attachment_list",
					Description: `Information list of the dedicated TCR namespaces.`,
				},
				resource.Attribute{
					Name:        "access_ip",
					Description: `IP address of this VPC access.`,
				},
				resource.Attribute{
					Name:        "enable_public_domain_dns",
					Description: `Whether to enable public domain dns.`,
				},
				resource.Attribute{
					Name:        "enable_vpc_domain_dns",
					Description: `Whether to enable vpc domain dns.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of this VPC access.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of subnet.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of VPC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_vod_adaptive_dynamic_streaming_templates",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of VOD adaptive dynamic streaming templates.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "definition",
					Description: `(Optional) Unique ID filter of adaptive dynamic streaming template.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "sub_app_id",
					Description: `(Optional) Subapplication ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Template type filter. Valid values: ` + "`" + `Preset` + "`" + `, ` + "`" + `Custom` + "`" + `. ` + "`" + `Preset` + "`" + `: preset template; ` + "`" + `Custom` + "`" + `: custom template. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "template_list",
					Description: `A list of adaptive dynamic streaming templates. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Template description.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of template in ISO date format.`,
				},
				resource.Attribute{
					Name:        "definition",
					Description: `Unique ID of adaptive dynamic streaming template.`,
				},
				resource.Attribute{
					Name:        "disable_higher_video_bitrate",
					Description: `Whether to prohibit transcoding video from low bitrate to high bitrate. ` + "`" + `false` + "`" + `: no, ` + "`" + `true` + "`" + `: yes.`,
				},
				resource.Attribute{
					Name:        "disable_higher_video_resolution",
					Description: `Whether to prohibit transcoding from low resolution to high resolution. ` + "`" + `false` + "`" + `: no, ` + "`" + `true` + "`" + `: yes.`,
				},
				resource.Attribute{
					Name:        "drm_type",
					Description: `DRM scheme type.`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `Adaptive bitstream format.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Template name.`,
				},
				resource.Attribute{
					Name:        "stream_info",
					Description: `List of AdaptiveStreamTemplate parameter information of output substream for adaptive bitrate streaming.`,
				},
				resource.Attribute{
					Name:        "audio",
					Description: `Audio parameter information.`,
				},
				resource.Attribute{
					Name:        "audio_channel",
					Description: `Audio channel system. Valid values: mono, dual, stereo.`,
				},
				resource.Attribute{
					Name:        "bitrate",
					Description: `Audio stream bitrate in Kbps. Value range: ` + "`" + `0` + "`" + ` and ` + "`" + `[26, 256]` + "`" + `. If the value is ` + "`" + `0` + "`" + `, the bitrate of the audio stream will be the same as that of the original audio.`,
				},
				resource.Attribute{
					Name:        "codec",
					Description: `Audio stream encoder. Valid value are: ` + "`" + `libfdk_aac` + "`" + ` and ` + "`" + `libmp3lame` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sample_rate",
					Description: `Audio stream sample rate. Valid values: ` + "`" + `32000` + "`" + `, ` + "`" + `44100` + "`" + `, ` + "`" + `48000` + "`" + `. Unit is HZ.`,
				},
				resource.Attribute{
					Name:        "remove_audio",
					Description: `Whether to remove audio stream. ` + "`" + `false` + "`" + `: no, ` + "`" + `true` + "`" + `: yes.`,
				},
				resource.Attribute{
					Name:        "video",
					Description: `Video parameter information.`,
				},
				resource.Attribute{
					Name:        "bitrate",
					Description: `Bitrate of video stream in Kbps. Value range: ` + "`" + `0` + "`" + ` and ` + "`" + `[128, 35000]` + "`" + `. If the value is ` + "`" + `0` + "`" + `, the bitrate of the video will be the same as that of the source video.`,
				},
				resource.Attribute{
					Name:        "codec",
					Description: `Video stream encoder. Valid values: ` + "`" + `libx264` + "`" + `, ` + "`" + `libx265` + "`" + `, ` + "`" + `av1` + "`" + `.` + "`" + `libx264` + "`" + `: H.264, ` + "`" + `libx265` + "`" + `: H.265, ` + "`" + `av1` + "`" + `: AOMedia Video 1. Currently, a resolution within 640x480 must be specified for ` + "`" + `H.265` + "`" + `. and the ` + "`" + `av1` + "`" + ` container only supports mp4.`,
				},
				resource.Attribute{
					Name:        "fill_type",
					Description: `Fill type. Fill refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported: ` + "`" + `stretch` + "`" + `: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot shorter or longer; ` + "`" + `black` + "`" + `: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks. Note: this field may return null, indicating that no valid values can be obtained.`,
				},
				resource.Attribute{
					Name:        "fps",
					Description: `Video frame rate in Hz. Value range: ` + "`" + `[0, 60]` + "`" + `. If the value is ` + "`" + `0` + "`" + `, the frame rate will be the same as that of the source video.`,
				},
				resource.Attribute{
					Name:        "height",
					Description: `Maximum value of the height (or short side) of a video stream in px. Value range: ` + "`" + `0` + "`" + ` and ` + "`" + `[128, 4096]` + "`" + `. If both ` + "`" + `width` + "`" + ` and ` + "`" + `height` + "`" + ` are ` + "`" + `0` + "`" + `, the resolution will be the same as that of the source video; If ` + "`" + `width` + "`" + ` is ` + "`" + `0` + "`" + `, but ` + "`" + `height` + "`" + ` is not ` + "`" + `0` + "`" + `, ` + "`" + `width` + "`" + ` will be proportionally scaled; If ` + "`" + `width` + "`" + ` is not ` + "`" + `0` + "`" + `, but ` + "`" + `height` + "`" + ` is ` + "`" + `0` + "`" + `, ` + "`" + `height` + "`" + ` will be proportionally scaled; If both ` + "`" + `width` + "`" + ` and ` + "`" + `height` + "`" + ` are not ` + "`" + `0` + "`" + `, the custom resolution will be used. Note: this field may return null, indicating that no valid values can be obtained.`,
				},
				resource.Attribute{
					Name:        "resolution_adaptive",
					Description: `Resolution adaption. Valid values: ` + "`" + `true` + "`" + `,` + "`" + `false` + "`" + `. ` + "`" + `true` + "`" + `: enabled. In this case, ` + "`" + `width` + "`" + ` represents the long side of a video, while ` + "`" + `height` + "`" + ` the short side; ` + "`" + `false` + "`" + `: disabled. In this case, ` + "`" + `width` + "`" + ` represents the width of a video, while ` + "`" + `height` + "`" + ` the height. Note: this field may return null, indicating that no valid values can be obtained.`,
				},
				resource.Attribute{
					Name:        "width",
					Description: `Maximum value of the width (or long side) of a video stream in px. Value range: ` + "`" + `0` + "`" + ` and ` + "`" + `[128, 4096]` + "`" + `. If both ` + "`" + `width` + "`" + ` and ` + "`" + `height` + "`" + ` are ` + "`" + `0` + "`" + `, the resolution will be the same as that of the source video; If ` + "`" + `width` + "`" + ` is ` + "`" + `0` + "`" + `, but ` + "`" + `height` + "`" + ` is not ` + "`" + `0` + "`" + `, ` + "`" + `width` + "`" + ` will be proportionally scaled; If ` + "`" + `width` + "`" + ` is not ` + "`" + `0` + "`" + `, but ` + "`" + `height` + "`" + ` is ` + "`" + `0` + "`" + `, ` + "`" + `height` + "`" + ` will be proportionally scaled; If both ` + "`" + `width` + "`" + ` and ` + "`" + `height` + "`" + ` are not ` + "`" + `0` + "`" + `, the custom resolution will be used. Note: this field may return null, indicating that no valid values can be obtained.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Template type filter. Valid values: ` + "`" + `Preset` + "`" + `,` + "`" + `Custom` + "`" + `. ` + "`" + `Preset` + "`" + `: preset template; ` + "`" + `Custom` + "`" + `: custom template.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Last modified time of template in ISO date format.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "template_list",
					Description: `A list of adaptive dynamic streaming templates. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Template description.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of template in ISO date format.`,
				},
				resource.Attribute{
					Name:        "definition",
					Description: `Unique ID of adaptive dynamic streaming template.`,
				},
				resource.Attribute{
					Name:        "disable_higher_video_bitrate",
					Description: `Whether to prohibit transcoding video from low bitrate to high bitrate. ` + "`" + `false` + "`" + `: no, ` + "`" + `true` + "`" + `: yes.`,
				},
				resource.Attribute{
					Name:        "disable_higher_video_resolution",
					Description: `Whether to prohibit transcoding from low resolution to high resolution. ` + "`" + `false` + "`" + `: no, ` + "`" + `true` + "`" + `: yes.`,
				},
				resource.Attribute{
					Name:        "drm_type",
					Description: `DRM scheme type.`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `Adaptive bitstream format.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Template name.`,
				},
				resource.Attribute{
					Name:        "stream_info",
					Description: `List of AdaptiveStreamTemplate parameter information of output substream for adaptive bitrate streaming.`,
				},
				resource.Attribute{
					Name:        "audio",
					Description: `Audio parameter information.`,
				},
				resource.Attribute{
					Name:        "audio_channel",
					Description: `Audio channel system. Valid values: mono, dual, stereo.`,
				},
				resource.Attribute{
					Name:        "bitrate",
					Description: `Audio stream bitrate in Kbps. Value range: ` + "`" + `0` + "`" + ` and ` + "`" + `[26, 256]` + "`" + `. If the value is ` + "`" + `0` + "`" + `, the bitrate of the audio stream will be the same as that of the original audio.`,
				},
				resource.Attribute{
					Name:        "codec",
					Description: `Audio stream encoder. Valid value are: ` + "`" + `libfdk_aac` + "`" + ` and ` + "`" + `libmp3lame` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sample_rate",
					Description: `Audio stream sample rate. Valid values: ` + "`" + `32000` + "`" + `, ` + "`" + `44100` + "`" + `, ` + "`" + `48000` + "`" + `. Unit is HZ.`,
				},
				resource.Attribute{
					Name:        "remove_audio",
					Description: `Whether to remove audio stream. ` + "`" + `false` + "`" + `: no, ` + "`" + `true` + "`" + `: yes.`,
				},
				resource.Attribute{
					Name:        "video",
					Description: `Video parameter information.`,
				},
				resource.Attribute{
					Name:        "bitrate",
					Description: `Bitrate of video stream in Kbps. Value range: ` + "`" + `0` + "`" + ` and ` + "`" + `[128, 35000]` + "`" + `. If the value is ` + "`" + `0` + "`" + `, the bitrate of the video will be the same as that of the source video.`,
				},
				resource.Attribute{
					Name:        "codec",
					Description: `Video stream encoder. Valid values: ` + "`" + `libx264` + "`" + `, ` + "`" + `libx265` + "`" + `, ` + "`" + `av1` + "`" + `.` + "`" + `libx264` + "`" + `: H.264, ` + "`" + `libx265` + "`" + `: H.265, ` + "`" + `av1` + "`" + `: AOMedia Video 1. Currently, a resolution within 640x480 must be specified for ` + "`" + `H.265` + "`" + `. and the ` + "`" + `av1` + "`" + ` container only supports mp4.`,
				},
				resource.Attribute{
					Name:        "fill_type",
					Description: `Fill type. Fill refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported: ` + "`" + `stretch` + "`" + `: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot shorter or longer; ` + "`" + `black` + "`" + `: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks. Note: this field may return null, indicating that no valid values can be obtained.`,
				},
				resource.Attribute{
					Name:        "fps",
					Description: `Video frame rate in Hz. Value range: ` + "`" + `[0, 60]` + "`" + `. If the value is ` + "`" + `0` + "`" + `, the frame rate will be the same as that of the source video.`,
				},
				resource.Attribute{
					Name:        "height",
					Description: `Maximum value of the height (or short side) of a video stream in px. Value range: ` + "`" + `0` + "`" + ` and ` + "`" + `[128, 4096]` + "`" + `. If both ` + "`" + `width` + "`" + ` and ` + "`" + `height` + "`" + ` are ` + "`" + `0` + "`" + `, the resolution will be the same as that of the source video; If ` + "`" + `width` + "`" + ` is ` + "`" + `0` + "`" + `, but ` + "`" + `height` + "`" + ` is not ` + "`" + `0` + "`" + `, ` + "`" + `width` + "`" + ` will be proportionally scaled; If ` + "`" + `width` + "`" + ` is not ` + "`" + `0` + "`" + `, but ` + "`" + `height` + "`" + ` is ` + "`" + `0` + "`" + `, ` + "`" + `height` + "`" + ` will be proportionally scaled; If both ` + "`" + `width` + "`" + ` and ` + "`" + `height` + "`" + ` are not ` + "`" + `0` + "`" + `, the custom resolution will be used. Note: this field may return null, indicating that no valid values can be obtained.`,
				},
				resource.Attribute{
					Name:        "resolution_adaptive",
					Description: `Resolution adaption. Valid values: ` + "`" + `true` + "`" + `,` + "`" + `false` + "`" + `. ` + "`" + `true` + "`" + `: enabled. In this case, ` + "`" + `width` + "`" + ` represents the long side of a video, while ` + "`" + `height` + "`" + ` the short side; ` + "`" + `false` + "`" + `: disabled. In this case, ` + "`" + `width` + "`" + ` represents the width of a video, while ` + "`" + `height` + "`" + ` the height. Note: this field may return null, indicating that no valid values can be obtained.`,
				},
				resource.Attribute{
					Name:        "width",
					Description: `Maximum value of the width (or long side) of a video stream in px. Value range: ` + "`" + `0` + "`" + ` and ` + "`" + `[128, 4096]` + "`" + `. If both ` + "`" + `width` + "`" + ` and ` + "`" + `height` + "`" + ` are ` + "`" + `0` + "`" + `, the resolution will be the same as that of the source video; If ` + "`" + `width` + "`" + ` is ` + "`" + `0` + "`" + `, but ` + "`" + `height` + "`" + ` is not ` + "`" + `0` + "`" + `, ` + "`" + `width` + "`" + ` will be proportionally scaled; If ` + "`" + `width` + "`" + ` is not ` + "`" + `0` + "`" + `, but ` + "`" + `height` + "`" + ` is ` + "`" + `0` + "`" + `, ` + "`" + `height` + "`" + ` will be proportionally scaled; If both ` + "`" + `width` + "`" + ` and ` + "`" + `height` + "`" + ` are not ` + "`" + `0` + "`" + `, the custom resolution will be used. Note: this field may return null, indicating that no valid values can be obtained.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Template type filter. Valid values: ` + "`" + `Preset` + "`" + `,` + "`" + `Custom` + "`" + `. ` + "`" + `Preset` + "`" + `: preset template; ` + "`" + `Custom` + "`" + `: custom template.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Last modified time of template in ISO date format.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_vod_image_sprite_templates",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of VOD image sprite templates.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "definition",
					Description: `(Optional) Unique ID filter of image sprite template.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "sub_app_id",
					Description: `(Optional) Subapplication ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Template type filter. Valid values: ` + "`" + `Preset` + "`" + `, ` + "`" + `Custom` + "`" + `. ` + "`" + `Preset` + "`" + `: preset template; ` + "`" + `Custom` + "`" + `: custom template. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "template_list",
					Description: `A list of image sprite templates. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "column_count",
					Description: `Subimage column count of an image sprite.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Template description.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of template in ISO date format.`,
				},
				resource.Attribute{
					Name:        "definition",
					Description: `Unique ID of image sprite template.`,
				},
				resource.Attribute{
					Name:        "fill_type",
					Description: `Fill refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported: ` + "`" + `stretch` + "`" + `: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot shorter or longer; ` + "`" + `black` + "`" + `: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks.`,
				},
				resource.Attribute{
					Name:        "height",
					Description: `Maximum value of the ` + "`" + `height` + "`" + ` (or short side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both ` + "`" + `width` + "`" + ` and ` + "`" + `height` + "`" + ` are ` + "`" + `0` + "`" + `, the resolution will be the same as that of the source video; If ` + "`" + `width` + "`" + ` is ` + "`" + `0` + "`" + `, but ` + "`" + `height` + "`" + ` is not ` + "`" + `0` + "`" + `, ` + "`" + `width` + "`" + ` will be proportionally scaled; If ` + "`" + `width` + "`" + ` is not ` + "`" + `0` + "`" + `, but ` + "`" + `height` + "`" + ` is ` + "`" + `0` + "`" + `, ` + "`" + `height` + "`" + ` will be proportionally scaled; If both ` + "`" + `width` + "`" + ` and ` + "`" + `height` + "`" + ` are not ` + "`" + `0` + "`" + `, the custom resolution will be used.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of a time point screen capturing template.`,
				},
				resource.Attribute{
					Name:        "resolution_adaptive",
					Description: `Resolution adaption. Valid values: ` + "`" + `true` + "`" + `: enabled. In this case, ` + "`" + `width` + "`" + ` represents the long side of a video, while ` + "`" + `height` + "`" + ` the short side; ` + "`" + `false` + "`" + `: disabled. In this case, ` + "`" + `width` + "`" + ` represents the width of a video, while ` + "`" + `height` + "`" + ` the height.`,
				},
				resource.Attribute{
					Name:        "row_count",
					Description: `Subimage row count of an image sprite.`,
				},
				resource.Attribute{
					Name:        "sample_interval",
					Description: `Sampling interval. If ` + "`" + `sample_type` + "`" + ` is ` + "`" + `Percent` + "`" + `, sampling will be performed at an interval of the specified percentage. If ` + "`" + `sample_type` + "`" + ` is ` + "`" + `Time` + "`" + `, sampling will be performed at the specified time interval in seconds.`,
				},
				resource.Attribute{
					Name:        "sample_type",
					Description: `Sampling type. Valid values: ` + "`" + `Percent` + "`" + `, ` + "`" + `Time` + "`" + `. ` + "`" + `Percent` + "`" + `: by percent. ` + "`" + `Time` + "`" + `: by time interval.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Template type filter. Valid values: ` + "`" + `Preset` + "`" + `, ` + "`" + `Custom` + "`" + `. ` + "`" + `Preset` + "`" + `: preset template; ` + "`" + `Custom` + "`" + `: custom template.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Last modified time of template in ISO date format.`,
				},
				resource.Attribute{
					Name:        "width",
					Description: `Maximum value of the ` + "`" + `width` + "`" + ` (or long side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both ` + "`" + `width` + "`" + ` and ` + "`" + `height` + "`" + ` are ` + "`" + `0` + "`" + `, the resolution will be the same as that of the source video; If ` + "`" + `width` + "`" + ` is ` + "`" + `0` + "`" + `, but ` + "`" + `height` + "`" + ` is not ` + "`" + `0` + "`" + `, width will be proportionally scaled; If ` + "`" + `width` + "`" + ` is not ` + "`" + `0` + "`" + `, but ` + "`" + `height` + "`" + ` is ` + "`" + `0` + "`" + `, ` + "`" + `height` + "`" + ` will be proportionally scaled; If both ` + "`" + `width` + "`" + ` and ` + "`" + `height` + "`" + ` are not ` + "`" + `0` + "`" + `, the custom resolution will be used.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "template_list",
					Description: `A list of image sprite templates. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "column_count",
					Description: `Subimage column count of an image sprite.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Template description.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of template in ISO date format.`,
				},
				resource.Attribute{
					Name:        "definition",
					Description: `Unique ID of image sprite template.`,
				},
				resource.Attribute{
					Name:        "fill_type",
					Description: `Fill refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported: ` + "`" + `stretch` + "`" + `: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot shorter or longer; ` + "`" + `black` + "`" + `: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks.`,
				},
				resource.Attribute{
					Name:        "height",
					Description: `Maximum value of the ` + "`" + `height` + "`" + ` (or short side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both ` + "`" + `width` + "`" + ` and ` + "`" + `height` + "`" + ` are ` + "`" + `0` + "`" + `, the resolution will be the same as that of the source video; If ` + "`" + `width` + "`" + ` is ` + "`" + `0` + "`" + `, but ` + "`" + `height` + "`" + ` is not ` + "`" + `0` + "`" + `, ` + "`" + `width` + "`" + ` will be proportionally scaled; If ` + "`" + `width` + "`" + ` is not ` + "`" + `0` + "`" + `, but ` + "`" + `height` + "`" + ` is ` + "`" + `0` + "`" + `, ` + "`" + `height` + "`" + ` will be proportionally scaled; If both ` + "`" + `width` + "`" + ` and ` + "`" + `height` + "`" + ` are not ` + "`" + `0` + "`" + `, the custom resolution will be used.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of a time point screen capturing template.`,
				},
				resource.Attribute{
					Name:        "resolution_adaptive",
					Description: `Resolution adaption. Valid values: ` + "`" + `true` + "`" + `: enabled. In this case, ` + "`" + `width` + "`" + ` represents the long side of a video, while ` + "`" + `height` + "`" + ` the short side; ` + "`" + `false` + "`" + `: disabled. In this case, ` + "`" + `width` + "`" + ` represents the width of a video, while ` + "`" + `height` + "`" + ` the height.`,
				},
				resource.Attribute{
					Name:        "row_count",
					Description: `Subimage row count of an image sprite.`,
				},
				resource.Attribute{
					Name:        "sample_interval",
					Description: `Sampling interval. If ` + "`" + `sample_type` + "`" + ` is ` + "`" + `Percent` + "`" + `, sampling will be performed at an interval of the specified percentage. If ` + "`" + `sample_type` + "`" + ` is ` + "`" + `Time` + "`" + `, sampling will be performed at the specified time interval in seconds.`,
				},
				resource.Attribute{
					Name:        "sample_type",
					Description: `Sampling type. Valid values: ` + "`" + `Percent` + "`" + `, ` + "`" + `Time` + "`" + `. ` + "`" + `Percent` + "`" + `: by percent. ` + "`" + `Time` + "`" + `: by time interval.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Template type filter. Valid values: ` + "`" + `Preset` + "`" + `, ` + "`" + `Custom` + "`" + `. ` + "`" + `Preset` + "`" + `: preset template; ` + "`" + `Custom` + "`" + `: custom template.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Last modified time of template in ISO date format.`,
				},
				resource.Attribute{
					Name:        "width",
					Description: `Maximum value of the ` + "`" + `width` + "`" + ` (or long side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both ` + "`" + `width` + "`" + ` and ` + "`" + `height` + "`" + ` are ` + "`" + `0` + "`" + `, the resolution will be the same as that of the source video; If ` + "`" + `width` + "`" + ` is ` + "`" + `0` + "`" + `, but ` + "`" + `height` + "`" + ` is not ` + "`" + `0` + "`" + `, width will be proportionally scaled; If ` + "`" + `width` + "`" + ` is not ` + "`" + `0` + "`" + `, but ` + "`" + `height` + "`" + ` is ` + "`" + `0` + "`" + `, ` + "`" + `height` + "`" + ` will be proportionally scaled; If both ` + "`" + `width` + "`" + ` and ` + "`" + `height` + "`" + ` are not ` + "`" + `0` + "`" + `, the custom resolution will be used.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_vod_procedure_templates",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of VOD procedure templates.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of procedure template.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "sub_app_id",
					Description: `(Optional) Subapplication ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Template type filter. Valid values: ` + "`" + `Preset` + "`" + `, ` + "`" + `Custom` + "`" + `. ` + "`" + `Preset` + "`" + `: preset template; ` + "`" + `Custom` + "`" + `: custom template. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "template_list",
					Description: `A list of adaptive dynamic streaming templates. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Template description.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of template in ISO date format.`,
				},
				resource.Attribute{
					Name:        "media_process_task",
					Description: `Parameter of video processing task.`,
				},
				resource.Attribute{
					Name:        "adaptive_dynamic_streaming_task_list",
					Description: `List of adaptive bitrate streaming tasks. Note: this field may return null, indicating that no valid values can be obtained.`,
				},
				resource.Attribute{
					Name:        "definition",
					Description: `Adaptive bitrate streaming template ID.`,
				},
				resource.Attribute{
					Name:        "watermark_list",
					Description: `List of up to ` + "`" + `10` + "`" + ` image or text watermarks. Note: this field may return null, indicating that no valid values can be obtained.`,
				},
				resource.Attribute{
					Name:        "animated_graphic_task_list",
					Description: `List of animated image generating tasks. Note: this field may return null, indicating that no valid values can be obtained.`,
				},
				resource.Attribute{
					Name:        "definition",
					Description: `Animated image generating template ID.`,
				},
				resource.Attribute{
					Name:        "end_time_offset",
					Description: `End time of animated image in video in seconds.`,
				},
				resource.Attribute{
					Name:        "start_time_offset",
					Description: `Start time of animated image in video in seconds.`,
				},
				resource.Attribute{
					Name:        "cover_by_snapshot_task_list",
					Description: `List of cover generating tasks. Note: this field may return null, indicating that no valid values can be obtained.`,
				},
				resource.Attribute{
					Name:        "definition",
					Description: `Time point screen capturing template ID.`,
				},
				resource.Attribute{
					Name:        "position_type",
					Description: `Screen capturing mode. Valid values: ` + "`" + `Time` + "`" + `, ` + "`" + `Percent` + "`" + `. ` + "`" + `Time` + "`" + `: screen captures by time point, ` + "`" + `Percent` + "`" + `: screen captures by percentage.`,
				},
				resource.Attribute{
					Name:        "position_value",
					Description: `Screenshot position: For time point screen capturing, this means to take a screenshot at a specified time point (in seconds) and use it as the cover. For percentage screen capturing, this value means to take a screenshot at a specified percentage of the video duration and use it as the cover.`,
				},
				resource.Attribute{
					Name:        "watermark_list",
					Description: `List of up to ` + "`" + `10` + "`" + ` image or text watermarks. Note: this field may return null, indicating that no valid values can be obtained.`,
				},
				resource.Attribute{
					Name:        "image_sprite_task_list",
					Description: `List of image sprite generating tasks. Note: this field may return null, indicating that no valid values can be obtained.`,
				},
				resource.Attribute{
					Name:        "definition",
					Description: `Image sprite generating template ID.`,
				},
				resource.Attribute{
					Name:        "sample_snapshot_task_list",
					Description: `List of sampled screen capturing tasks. Note: this field may return null, indicating that no valid values can be obtained.`,
				},
				resource.Attribute{
					Name:        "definition",
					Description: `Sampled screen capturing template ID.`,
				},
				resource.Attribute{
					Name:        "watermark_list",
					Description: `List of up to ` + "`" + `10` + "`" + ` image or text watermarks. Note: this field may return null, indicating that no valid values can be obtained.`,
				},
				resource.Attribute{
					Name:        "snapshot_by_time_offset_task_list",
					Description: `List of time point screen capturing tasks. Note: this field may return null, indicating that no valid values can be obtained.`,
				},
				resource.Attribute{
					Name:        "definition",
					Description: `Time point screen capturing template ID.`,
				},
				resource.Attribute{
					Name:        "ext_time_offset_list",
					Description: `The list of screenshot time points. ` + "`" + `s` + "`" + ` and ` + "`" + `%` + "`" + ` formats are supported: When a time point string ends with ` + "`" + `s` + "`" + `, its unit is second. For example, ` + "`" + `3.5s` + "`" + ` means the 3.5th second of the video; When a time point string ends with ` + "`" + `%` + "`" + `, it is marked with corresponding percentage of the video duration. For example, ` + "`" + `10%` + "`" + ` means that the time point is at the 10% of the video entire duration.`,
				},
				resource.Attribute{
					Name:        "watermark_list",
					Description: `List of up to ` + "`" + `10` + "`" + ` image or text watermarks. Note: this field may return null, indicating that no valid values can be obtained.`,
				},
				resource.Attribute{
					Name:        "transcode_task_list",
					Description: `List of transcoding tasks. Note: this field may return null, indicating that no valid values can be obtained.`,
				},
				resource.Attribute{
					Name:        "definition",
					Description: `Video transcoding template ID.`,
				},
				resource.Attribute{
					Name:        "mosaic_list",
					Description: `List of blurs. Up to 10 ones can be supported.`,
				},
				resource.Attribute{
					Name:        "coordinate_origin",
					Description: `Origin position, which currently can only be: ` + "`" + `TopLeft` + "`" + `: the origin of coordinates is in the top-left corner of the video, and the origin of the blur is in the top-left corner of the image or text.`,
				},
				resource.Attribute{
					Name:        "end_time_offset",
					Description: `End time offset of blur in seconds. If this parameter is left empty or ` + "`" + `0` + "`" + ` is entered, the blur will exist till the last video frame; If this value is greater than ` + "`" + `0` + "`" + ` (e.g., n), the blur will exist till second n; If this value is smaller than ` + "`" + `0` + "`" + ` (e.g., -n), the blur will exist till second n before the last video frame.`,
				},
				resource.Attribute{
					Name:        "height",
					Description: `Blur height. ` + "`" + `%` + "`" + ` and ` + "`" + `px` + "`" + ` formats are supported: If the string ends in ` + "`" + `%` + "`" + `, the ` + "`" + `height` + "`" + ` of the blur will be the specified percentage of the video height; for example, 10% means that Height is 10% of the video height; If the string ends in ` + "`" + `px` + "`" + `, the ` + "`" + `height` + "`" + ` of the blur will be in px; for example, 100px means that Height is 100 px.`,
				},
				resource.Attribute{
					Name:        "start_time_offset",
					Description: `Start time offset of blur in seconds. If this parameter is left empty or ` + "`" + `0` + "`" + ` is entered, the blur will appear upon the first video frame. If this parameter is left empty or ` + "`" + `0` + "`" + ` is entered, the blur will appear upon the first video frame; If this value is greater than ` + "`" + `0` + "`" + ` (e.g., n), the blur will appear at second n after the first video frame; If this value is smaller than ` + "`" + `0` + "`" + ` (e.g., -n), the blur will appear at second n before the last video frame.`,
				},
				resource.Attribute{
					Name:        "width",
					Description: `Blur width. ` + "`" + `%` + "`" + ` and ` + "`" + `px` + "`" + ` formats are supported: If the string ends in ` + "`" + `%` + "`" + `, the ` + "`" + `width` + "`" + ` of the blur will be the specified percentage of the video width; for example, 10% means that ` + "`" + `width` + "`" + ` is 10% of the video width; If the string ends in ` + "`" + `px` + "`" + `, the ` + "`" + `width` + "`" + ` of the blur will be in px; for example, 100px means that Width is 100 px.`,
				},
				resource.Attribute{
					Name:        "x_pos",
					Description: `The horizontal position of the origin of the blur relative to the origin of coordinates of the video. ` + "`" + `%` + "`" + ` and ` + "`" + `px` + "`" + ` formats are supported: If the string ends in ` + "`" + `%` + "`" + `, the XPos of the blur will be the specified percentage of the video width; for example, 10% means that XPos is 10% of the video width; If the string ends in ` + "`" + `px` + "`" + `, the XPos of the blur will be the specified px; for example, 100px means that XPos is 100 px.`,
				},
				resource.Attribute{
					Name:        "y_pos",
					Description: `Vertical position of the origin of blur relative to the origin of coordinates of video. ` + "`" + `%` + "`" + ` and ` + "`" + `px` + "`" + ` formats are supported: If the string ends in ` + "`" + `%` + "`" + `, the YPos of the blur will be the specified percentage of the video height; for example, 10% means that YPos is 10% of the video height; If the string ends in ` + "`" + `px` + "`" + `, the YPos of the blur will be the specified px; for example, 100px means that YPos is 100 px.`,
				},
				resource.Attribute{
					Name:        "watermark_list",
					Description: `List of up to ` + "`" + `10` + "`" + ` image or text watermarks. Note: this field may return null, indicating that no valid values can be obtained.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Task flow name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Template type filter. Valid values: ` + "`" + `Preset` + "`" + `, ` + "`" + `Custom` + "`" + `. ` + "`" + `Preset` + "`" + `: preset template; ` + "`" + `Custom` + "`" + `: custom template.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Last modified time of template in ISO date format.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "template_list",
					Description: `A list of adaptive dynamic streaming templates. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Template description.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of template in ISO date format.`,
				},
				resource.Attribute{
					Name:        "media_process_task",
					Description: `Parameter of video processing task.`,
				},
				resource.Attribute{
					Name:        "adaptive_dynamic_streaming_task_list",
					Description: `List of adaptive bitrate streaming tasks. Note: this field may return null, indicating that no valid values can be obtained.`,
				},
				resource.Attribute{
					Name:        "definition",
					Description: `Adaptive bitrate streaming template ID.`,
				},
				resource.Attribute{
					Name:        "watermark_list",
					Description: `List of up to ` + "`" + `10` + "`" + ` image or text watermarks. Note: this field may return null, indicating that no valid values can be obtained.`,
				},
				resource.Attribute{
					Name:        "animated_graphic_task_list",
					Description: `List of animated image generating tasks. Note: this field may return null, indicating that no valid values can be obtained.`,
				},
				resource.Attribute{
					Name:        "definition",
					Description: `Animated image generating template ID.`,
				},
				resource.Attribute{
					Name:        "end_time_offset",
					Description: `End time of animated image in video in seconds.`,
				},
				resource.Attribute{
					Name:        "start_time_offset",
					Description: `Start time of animated image in video in seconds.`,
				},
				resource.Attribute{
					Name:        "cover_by_snapshot_task_list",
					Description: `List of cover generating tasks. Note: this field may return null, indicating that no valid values can be obtained.`,
				},
				resource.Attribute{
					Name:        "definition",
					Description: `Time point screen capturing template ID.`,
				},
				resource.Attribute{
					Name:        "position_type",
					Description: `Screen capturing mode. Valid values: ` + "`" + `Time` + "`" + `, ` + "`" + `Percent` + "`" + `. ` + "`" + `Time` + "`" + `: screen captures by time point, ` + "`" + `Percent` + "`" + `: screen captures by percentage.`,
				},
				resource.Attribute{
					Name:        "position_value",
					Description: `Screenshot position: For time point screen capturing, this means to take a screenshot at a specified time point (in seconds) and use it as the cover. For percentage screen capturing, this value means to take a screenshot at a specified percentage of the video duration and use it as the cover.`,
				},
				resource.Attribute{
					Name:        "watermark_list",
					Description: `List of up to ` + "`" + `10` + "`" + ` image or text watermarks. Note: this field may return null, indicating that no valid values can be obtained.`,
				},
				resource.Attribute{
					Name:        "image_sprite_task_list",
					Description: `List of image sprite generating tasks. Note: this field may return null, indicating that no valid values can be obtained.`,
				},
				resource.Attribute{
					Name:        "definition",
					Description: `Image sprite generating template ID.`,
				},
				resource.Attribute{
					Name:        "sample_snapshot_task_list",
					Description: `List of sampled screen capturing tasks. Note: this field may return null, indicating that no valid values can be obtained.`,
				},
				resource.Attribute{
					Name:        "definition",
					Description: `Sampled screen capturing template ID.`,
				},
				resource.Attribute{
					Name:        "watermark_list",
					Description: `List of up to ` + "`" + `10` + "`" + ` image or text watermarks. Note: this field may return null, indicating that no valid values can be obtained.`,
				},
				resource.Attribute{
					Name:        "snapshot_by_time_offset_task_list",
					Description: `List of time point screen capturing tasks. Note: this field may return null, indicating that no valid values can be obtained.`,
				},
				resource.Attribute{
					Name:        "definition",
					Description: `Time point screen capturing template ID.`,
				},
				resource.Attribute{
					Name:        "ext_time_offset_list",
					Description: `The list of screenshot time points. ` + "`" + `s` + "`" + ` and ` + "`" + `%` + "`" + ` formats are supported: When a time point string ends with ` + "`" + `s` + "`" + `, its unit is second. For example, ` + "`" + `3.5s` + "`" + ` means the 3.5th second of the video; When a time point string ends with ` + "`" + `%` + "`" + `, it is marked with corresponding percentage of the video duration. For example, ` + "`" + `10%` + "`" + ` means that the time point is at the 10% of the video entire duration.`,
				},
				resource.Attribute{
					Name:        "watermark_list",
					Description: `List of up to ` + "`" + `10` + "`" + ` image or text watermarks. Note: this field may return null, indicating that no valid values can be obtained.`,
				},
				resource.Attribute{
					Name:        "transcode_task_list",
					Description: `List of transcoding tasks. Note: this field may return null, indicating that no valid values can be obtained.`,
				},
				resource.Attribute{
					Name:        "definition",
					Description: `Video transcoding template ID.`,
				},
				resource.Attribute{
					Name:        "mosaic_list",
					Description: `List of blurs. Up to 10 ones can be supported.`,
				},
				resource.Attribute{
					Name:        "coordinate_origin",
					Description: `Origin position, which currently can only be: ` + "`" + `TopLeft` + "`" + `: the origin of coordinates is in the top-left corner of the video, and the origin of the blur is in the top-left corner of the image or text.`,
				},
				resource.Attribute{
					Name:        "end_time_offset",
					Description: `End time offset of blur in seconds. If this parameter is left empty or ` + "`" + `0` + "`" + ` is entered, the blur will exist till the last video frame; If this value is greater than ` + "`" + `0` + "`" + ` (e.g., n), the blur will exist till second n; If this value is smaller than ` + "`" + `0` + "`" + ` (e.g., -n), the blur will exist till second n before the last video frame.`,
				},
				resource.Attribute{
					Name:        "height",
					Description: `Blur height. ` + "`" + `%` + "`" + ` and ` + "`" + `px` + "`" + ` formats are supported: If the string ends in ` + "`" + `%` + "`" + `, the ` + "`" + `height` + "`" + ` of the blur will be the specified percentage of the video height; for example, 10% means that Height is 10% of the video height; If the string ends in ` + "`" + `px` + "`" + `, the ` + "`" + `height` + "`" + ` of the blur will be in px; for example, 100px means that Height is 100 px.`,
				},
				resource.Attribute{
					Name:        "start_time_offset",
					Description: `Start time offset of blur in seconds. If this parameter is left empty or ` + "`" + `0` + "`" + ` is entered, the blur will appear upon the first video frame. If this parameter is left empty or ` + "`" + `0` + "`" + ` is entered, the blur will appear upon the first video frame; If this value is greater than ` + "`" + `0` + "`" + ` (e.g., n), the blur will appear at second n after the first video frame; If this value is smaller than ` + "`" + `0` + "`" + ` (e.g., -n), the blur will appear at second n before the last video frame.`,
				},
				resource.Attribute{
					Name:        "width",
					Description: `Blur width. ` + "`" + `%` + "`" + ` and ` + "`" + `px` + "`" + ` formats are supported: If the string ends in ` + "`" + `%` + "`" + `, the ` + "`" + `width` + "`" + ` of the blur will be the specified percentage of the video width; for example, 10% means that ` + "`" + `width` + "`" + ` is 10% of the video width; If the string ends in ` + "`" + `px` + "`" + `, the ` + "`" + `width` + "`" + ` of the blur will be in px; for example, 100px means that Width is 100 px.`,
				},
				resource.Attribute{
					Name:        "x_pos",
					Description: `The horizontal position of the origin of the blur relative to the origin of coordinates of the video. ` + "`" + `%` + "`" + ` and ` + "`" + `px` + "`" + ` formats are supported: If the string ends in ` + "`" + `%` + "`" + `, the XPos of the blur will be the specified percentage of the video width; for example, 10% means that XPos is 10% of the video width; If the string ends in ` + "`" + `px` + "`" + `, the XPos of the blur will be the specified px; for example, 100px means that XPos is 100 px.`,
				},
				resource.Attribute{
					Name:        "y_pos",
					Description: `Vertical position of the origin of blur relative to the origin of coordinates of video. ` + "`" + `%` + "`" + ` and ` + "`" + `px` + "`" + ` formats are supported: If the string ends in ` + "`" + `%` + "`" + `, the YPos of the blur will be the specified percentage of the video height; for example, 10% means that YPos is 10% of the video height; If the string ends in ` + "`" + `px` + "`" + `, the YPos of the blur will be the specified px; for example, 100px means that YPos is 100 px.`,
				},
				resource.Attribute{
					Name:        "watermark_list",
					Description: `List of up to ` + "`" + `10` + "`" + ` image or text watermarks. Note: this field may return null, indicating that no valid values can be obtained.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Task flow name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Template type filter. Valid values: ` + "`" + `Preset` + "`" + `, ` + "`" + `Custom` + "`" + `. ` + "`" + `Preset` + "`" + `: preset template; ` + "`" + `Custom` + "`" + `: custom template.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Last modified time of template in ISO date format.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_vod_snapshot_by_time_offset_templates",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of VOD snapshot by time offset templates.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "definition",
					Description: `(Optional) Unique ID filter of snapshot by time offset template.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "sub_app_id",
					Description: `(Optional) Subapplication ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Template type filter. Valid values: ` + "`" + `Preset` + "`" + `, ` + "`" + `Custom` + "`" + `. ` + "`" + `Preset` + "`" + `: preset template; ` + "`" + `Custom` + "`" + `: custom template. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "template_list",
					Description: `A list of snapshot by time offset templates. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Template description.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of template in ISO date format.`,
				},
				resource.Attribute{
					Name:        "definition",
					Description: `Unique ID of snapshot by time offset template.`,
				},
				resource.Attribute{
					Name:        "fill_type",
					Description: `Fill refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported: ` + "`" + `stretch` + "`" + `: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot ` + "`" + `shorter` + "`" + ` or ` + "`" + `longer` + "`" + `; ` + "`" + `black` + "`" + `: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks. ` + "`" + `white` + "`" + `: fill with white. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with white color blocks. ` + "`" + `gauss` + "`" + `: fill with Gaussian blur. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with Gaussian blur.`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `Image format. Valid values: ` + "`" + `jpg` + "`" + `, ` + "`" + `png` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "height",
					Description: `Maximum value of the ` + "`" + `height` + "`" + ` (or short side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both ` + "`" + `width` + "`" + ` and ` + "`" + `height` + "`" + ` are ` + "`" + `0` + "`" + `, the resolution will be the same as that of the source video; If ` + "`" + `width` + "`" + ` is ` + "`" + `0` + "`" + `, but ` + "`" + `height` + "`" + ` is not ` + "`" + `0` + "`" + `, ` + "`" + `width` + "`" + ` will be proportionally scaled; If ` + "`" + `width` + "`" + ` is not ` + "`" + `0` + "`" + `, but ` + "`" + `height` + "`" + ` is ` + "`" + `0` + "`" + `, ` + "`" + `height` + "`" + ` will be proportionally scaled; If both ` + "`" + `width` + "`" + ` and ` + "`" + `height` + "`" + ` are not ` + "`" + `0` + "`" + `, the custom resolution will be used.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of a time point screen capturing template.`,
				},
				resource.Attribute{
					Name:        "resolution_adaptive",
					Description: `Resolution adaption. Valid values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + `. ` + "`" + `true` + "`" + `: enabled. In this case, ` + "`" + `width` + "`" + ` represents the long side of a video, while ` + "`" + `height` + "`" + ` the short side; ` + "`" + `false` + "`" + `: disabled. In this case, ` + "`" + `width` + "`" + ` represents the width of a video, while ` + "`" + `height` + "`" + ` the height.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Template type filter. Valid values: ` + "`" + `Preset` + "`" + `, ` + "`" + `Custom` + "`" + `. ` + "`" + `Preset` + "`" + `: preset template; ` + "`" + `Custom` + "`" + `: custom template.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Last modified time of template in ISO date format.`,
				},
				resource.Attribute{
					Name:        "width",
					Description: `Maximum value of the ` + "`" + `width` + "`" + ` (or long side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both ` + "`" + `width` + "`" + ` and ` + "`" + `height` + "`" + ` are ` + "`" + `0` + "`" + `, the resolution will be the same as that of the source video; If ` + "`" + `width` + "`" + ` is ` + "`" + `0` + "`" + `, but ` + "`" + `height` + "`" + ` is not ` + "`" + `0` + "`" + `, width will be proportionally scaled; If ` + "`" + `width` + "`" + ` is not ` + "`" + `0` + "`" + `, but ` + "`" + `height` + "`" + ` is ` + "`" + `0` + "`" + `, ` + "`" + `height` + "`" + ` will be proportionally scaled; If both ` + "`" + `width` + "`" + ` and ` + "`" + `height` + "`" + ` are not ` + "`" + `0` + "`" + `, the custom resolution will be used.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "template_list",
					Description: `A list of snapshot by time offset templates. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Template description.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of template in ISO date format.`,
				},
				resource.Attribute{
					Name:        "definition",
					Description: `Unique ID of snapshot by time offset template.`,
				},
				resource.Attribute{
					Name:        "fill_type",
					Description: `Fill refers to the way of processing a screenshot when its aspect ratio is different from that of the source video. The following fill types are supported: ` + "`" + `stretch` + "`" + `: stretch. The screenshot will be stretched frame by frame to match the aspect ratio of the source video, which may make the screenshot ` + "`" + `shorter` + "`" + ` or ` + "`" + `longer` + "`" + `; ` + "`" + `black` + "`" + `: fill with black. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black color blocks. ` + "`" + `white` + "`" + `: fill with white. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with white color blocks. ` + "`" + `gauss` + "`" + `: fill with Gaussian blur. This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with Gaussian blur.`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `Image format. Valid values: ` + "`" + `jpg` + "`" + `, ` + "`" + `png` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "height",
					Description: `Maximum value of the ` + "`" + `height` + "`" + ` (or short side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both ` + "`" + `width` + "`" + ` and ` + "`" + `height` + "`" + ` are ` + "`" + `0` + "`" + `, the resolution will be the same as that of the source video; If ` + "`" + `width` + "`" + ` is ` + "`" + `0` + "`" + `, but ` + "`" + `height` + "`" + ` is not ` + "`" + `0` + "`" + `, ` + "`" + `width` + "`" + ` will be proportionally scaled; If ` + "`" + `width` + "`" + ` is not ` + "`" + `0` + "`" + `, but ` + "`" + `height` + "`" + ` is ` + "`" + `0` + "`" + `, ` + "`" + `height` + "`" + ` will be proportionally scaled; If both ` + "`" + `width` + "`" + ` and ` + "`" + `height` + "`" + ` are not ` + "`" + `0` + "`" + `, the custom resolution will be used.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of a time point screen capturing template.`,
				},
				resource.Attribute{
					Name:        "resolution_adaptive",
					Description: `Resolution adaption. Valid values: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + `. ` + "`" + `true` + "`" + `: enabled. In this case, ` + "`" + `width` + "`" + ` represents the long side of a video, while ` + "`" + `height` + "`" + ` the short side; ` + "`" + `false` + "`" + `: disabled. In this case, ` + "`" + `width` + "`" + ` represents the width of a video, while ` + "`" + `height` + "`" + ` the height.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Template type filter. Valid values: ` + "`" + `Preset` + "`" + `, ` + "`" + `Custom` + "`" + `. ` + "`" + `Preset` + "`" + `: preset template; ` + "`" + `Custom` + "`" + `: custom template.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Last modified time of template in ISO date format.`,
				},
				resource.Attribute{
					Name:        "width",
					Description: `Maximum value of the ` + "`" + `width` + "`" + ` (or long side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both ` + "`" + `width` + "`" + ` and ` + "`" + `height` + "`" + ` are ` + "`" + `0` + "`" + `, the resolution will be the same as that of the source video; If ` + "`" + `width` + "`" + ` is ` + "`" + `0` + "`" + `, but ` + "`" + `height` + "`" + ` is not ` + "`" + `0` + "`" + `, width will be proportionally scaled; If ` + "`" + `width` + "`" + ` is not ` + "`" + `0` + "`" + `, but ` + "`" + `height` + "`" + ` is ` + "`" + `0` + "`" + `, ` + "`" + `height` + "`" + ` will be proportionally scaled; If both ` + "`" + `width` + "`" + ` and ` + "`" + `height` + "`" + ` are not ` + "`" + `0` + "`" + `, the custom resolution will be used.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_vod_super_player_configs",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of VOD super player configs.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of super player config.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "sub_app_id",
					Description: `(Optional) Subapplication ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Config type filter. Valid values: ` + "`" + `Preset` + "`" + `, ` + "`" + `Custom` + "`" + `. ` + "`" + `Preset` + "`" + `: preset template; ` + "`" + `Custom` + "`" + `: custom template. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "config_list",
					Description: `A list of super player configs. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "adaptive_dynamic_streaming_definition",
					Description: `ID of the unencrypted adaptive bitrate streaming template that allows output, which is required if ` + "`" + `drm_switch` + "`" + ` is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Template description.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of template in ISO date format.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Domain name used for playback. If it is left empty or set to ` + "`" + `Default` + "`" + `, the domain name configured in [Default Distribution Configuration](https://cloud.tencent.com/document/product/266/33373) will be used.`,
				},
				resource.Attribute{
					Name:        "drm_streaming_info",
					Description: `Content of the DRM-protected adaptive bitrate streaming template that allows output, which is required if ` + "`" + `drm_switch` + "`" + ` is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "simple_aes_definition",
					Description: `ID of the adaptive dynamic streaming template whose protection type is ` + "`" + `SimpleAES` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "drm_switch",
					Description: `Switch of DRM-protected adaptive bitstream playback: ` + "`" + `true` + "`" + `: enabled, indicating to play back only output adaptive bitstreams protected by DRM; ` + "`" + `false` + "`" + `: disabled, indicating to play back unencrypted output adaptive bitstreams.`,
				},
				resource.Attribute{
					Name:        "image_sprite_definition",
					Description: `ID of the image sprite template that allows output.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Player configuration name, which can contain up to 64 letters, digits, underscores, and hyphens (such as test_ABC-123) and must be unique under a user.`,
				},
				resource.Attribute{
					Name:        "resolution_names",
					Description: `Display name of player for substreams with different resolutions. If this parameter is left empty or an empty array, the default configuration will be used: ` + "`" + `min_edge_length: 240, name: LD` + "`" + `; ` + "`" + `min_edge_length: 480, name: SD` + "`" + `; ` + "`" + `min_edge_length: 720, name: HD` + "`" + `; ` + "`" + `min_edge_length: 1080, name: FHD` + "`" + `; ` + "`" + `min_edge_length: 1440, name: 2K` + "`" + `; ` + "`" + `min_edge_length: 2160, name: 4K` + "`" + `; ` + "`" + `min_edge_length: 4320, name: 8K` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "min_edge_length",
					Description: `Length of video short side in px.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Display name.`,
				},
				resource.Attribute{
					Name:        "scheme",
					Description: `Scheme used for playback. If it is left empty or set to ` + "`" + `Default` + "`" + `, the scheme configured in [Default Distribution Configuration](https://cloud.tencent.com/document/product/266/33373) will be used. Other valid values: ` + "`" + `HTTP` + "`" + `; ` + "`" + `HTTPS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Template type filter. Valid values: ` + "`" + `Preset` + "`" + `, ` + "`" + `Custom` + "`" + `. ` + "`" + `Preset` + "`" + `: preset template; ` + "`" + `Custom` + "`" + `: custom template.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Last modified time of template in ISO date format.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "config_list",
					Description: `A list of super player configs. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "adaptive_dynamic_streaming_definition",
					Description: `ID of the unencrypted adaptive bitrate streaming template that allows output, which is required if ` + "`" + `drm_switch` + "`" + ` is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Template description.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of template in ISO date format.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Domain name used for playback. If it is left empty or set to ` + "`" + `Default` + "`" + `, the domain name configured in [Default Distribution Configuration](https://cloud.tencent.com/document/product/266/33373) will be used.`,
				},
				resource.Attribute{
					Name:        "drm_streaming_info",
					Description: `Content of the DRM-protected adaptive bitrate streaming template that allows output, which is required if ` + "`" + `drm_switch` + "`" + ` is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "simple_aes_definition",
					Description: `ID of the adaptive dynamic streaming template whose protection type is ` + "`" + `SimpleAES` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "drm_switch",
					Description: `Switch of DRM-protected adaptive bitstream playback: ` + "`" + `true` + "`" + `: enabled, indicating to play back only output adaptive bitstreams protected by DRM; ` + "`" + `false` + "`" + `: disabled, indicating to play back unencrypted output adaptive bitstreams.`,
				},
				resource.Attribute{
					Name:        "image_sprite_definition",
					Description: `ID of the image sprite template that allows output.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Player configuration name, which can contain up to 64 letters, digits, underscores, and hyphens (such as test_ABC-123) and must be unique under a user.`,
				},
				resource.Attribute{
					Name:        "resolution_names",
					Description: `Display name of player for substreams with different resolutions. If this parameter is left empty or an empty array, the default configuration will be used: ` + "`" + `min_edge_length: 240, name: LD` + "`" + `; ` + "`" + `min_edge_length: 480, name: SD` + "`" + `; ` + "`" + `min_edge_length: 720, name: HD` + "`" + `; ` + "`" + `min_edge_length: 1080, name: FHD` + "`" + `; ` + "`" + `min_edge_length: 1440, name: 2K` + "`" + `; ` + "`" + `min_edge_length: 2160, name: 4K` + "`" + `; ` + "`" + `min_edge_length: 4320, name: 8K` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "min_edge_length",
					Description: `Length of video short side in px.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Display name.`,
				},
				resource.Attribute{
					Name:        "scheme",
					Description: `Scheme used for playback. If it is left empty or set to ` + "`" + `Default` + "`" + `, the scheme configured in [Default Distribution Configuration](https://cloud.tencent.com/document/product/266/33373) will be used. Other valid values: ` + "`" + `HTTP` + "`" + `; ` + "`" + `HTTPS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Template type filter. Valid values: ` + "`" + `Preset` + "`" + `, ` + "`" + `Custom` + "`" + `. ` + "`" + `Preset` + "`" + `: preset template; ` + "`" + `Custom` + "`" + `: custom template.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Last modified time of template in ISO date format.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_vpc",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific VPC.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the specific VPC to retrieve.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the specific VPC to retrieve. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The CIDR block of the VPC.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Whether or not the default VPC.`,
				},
				resource.Attribute{
					Name:        "is_multicast",
					Description: `Whether or not the VPC has Multicast support.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The CIDR block of the VPC.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Whether or not the default VPC.`,
				},
				resource.Attribute{
					Name:        "is_multicast",
					Description: `Whether or not the VPC has Multicast support.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_vpc_acls",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query VPC Network ACL information.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the network ACL instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the network ACL.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) ID of the VPC instance. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "acl_list",
					Description: `The information list of the VPC. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time.`,
				},
				resource.Attribute{
					Name:        "egress",
					Description: `Outbound rules of the network ACL.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `An IP address network or segment.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Rule description.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `Rule policy of Network ACL.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Range of the port.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Type of IP protocol.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the network ACL instance.`,
				},
				resource.Attribute{
					Name:        "ingress",
					Description: `Inbound rules of the network ACL.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `An IP address network or segment.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Rule description.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `Rule policy of Network ACL.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Range of the port.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Type of IP protocol.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the network ACL.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `Subnets associated with the network ACL.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The IPv4 CIDR of the subnet.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `Subnet instance ID.`,
				},
				resource.Attribute{
					Name:        "subnet_name",
					Description: `Subnet name.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the subnet.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC instance.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "acl_list",
					Description: `The information list of the VPC. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time.`,
				},
				resource.Attribute{
					Name:        "egress",
					Description: `Outbound rules of the network ACL.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `An IP address network or segment.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Rule description.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `Rule policy of Network ACL.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Range of the port.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Type of IP protocol.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the network ACL instance.`,
				},
				resource.Attribute{
					Name:        "ingress",
					Description: `Inbound rules of the network ACL.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `An IP address network or segment.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Rule description.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `Rule policy of Network ACL.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Range of the port.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Type of IP protocol.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the network ACL.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `Subnets associated with the network ACL.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The IPv4 CIDR of the subnet.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `Subnet instance ID.`,
				},
				resource.Attribute{
					Name:        "subnet_name",
					Description: `Subnet name.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the subnet.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC instance.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_vpc_instances",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query vpc instances' information.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Optional) Filter VPC with this CIDR.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `(Optional) Filter default or no default VPC.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the VPC to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "tag_key",
					Description: `(Optional) Filter if VPC has this tag.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of the VPC to be queried.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) ID of the VPC to be queried. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "instance_list",
					Description: `The information list of the VPC.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `A network address block of a VPC CIDR.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of VPC.`,
				},
				resource.Attribute{
					Name:        "dns_servers",
					Description: `A list of DNS servers which can be used within the VPC.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Indicates whether it is the default VPC for this region.`,
				},
				resource.Attribute{
					Name:        "is_multicast",
					Description: `Indicates whether VPC multicast is enabled.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the VPC.`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `A ID list of subnets within this VPC.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the VPC.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_list",
					Description: `The information list of the VPC.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `A network address block of a VPC CIDR.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of VPC.`,
				},
				resource.Attribute{
					Name:        "dns_servers",
					Description: `A list of DNS servers which can be used within the VPC.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Indicates whether it is the default VPC for this region.`,
				},
				resource.Attribute{
					Name:        "is_multicast",
					Description: `Indicates whether VPC multicast is enabled.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the VPC.`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `A ID list of subnets within this VPC.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the VPC.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_vpc_route_tables",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query vpc route tables information.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "association_main",
					Description: `(Optional) Filter the main routing table.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the routing table to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `(Optional) ID of the routing table to be queried.`,
				},
				resource.Attribute{
					Name:        "tag_key",
					Description: `(Optional) Filter if routing table has this tag.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of the routing table to be queried.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) ID of the VPC to be queried. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "instance_list",
					Description: `The information list of the VPC route table.`,
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
					Name:        "name",
					Description: `Name of the routing table.`,
				},
				resource.Attribute{
					Name:        "route_entry_infos",
					Description: `Detailed information of each entry of the route table.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description information user defined for a route table rule.`,
				},
				resource.Attribute{
					Name:        "destination_cidr_block",
					Description: `The destination address block.`,
				},
				resource.Attribute{
					Name:        "next_hub",
					Description: `ID of next-hop gateway. Note: when 'next_type' is EIP, GatewayId will fix the value ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "next_type",
					Description: `Type of next-hop, and available values include ` + "`" + `CVM` + "`" + `, ` + "`" + `VPN` + "`" + `, ` + "`" + `DIRECTCONNECT` + "`" + `, ` + "`" + `PEERCONNECTION` + "`" + `, ` + "`" + `SSLVPN` + "`" + `, ` + "`" + `NAT` + "`" + `, ` + "`" + `NORMAL_CVM` + "`" + `, ` + "`" + `EIP` + "`" + ` and ` + "`" + `CCN` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "route_entry_id",
					Description: `ID of a route table entry.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `ID of the routing table.`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `List of subnet IDs bound to the route table.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the routing table.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_list",
					Description: `The information list of the VPC route table.`,
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
					Name:        "name",
					Description: `Name of the routing table.`,
				},
				resource.Attribute{
					Name:        "route_entry_infos",
					Description: `Detailed information of each entry of the route table.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description information user defined for a route table rule.`,
				},
				resource.Attribute{
					Name:        "destination_cidr_block",
					Description: `The destination address block.`,
				},
				resource.Attribute{
					Name:        "next_hub",
					Description: `ID of next-hop gateway. Note: when 'next_type' is EIP, GatewayId will fix the value ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "next_type",
					Description: `Type of next-hop, and available values include ` + "`" + `CVM` + "`" + `, ` + "`" + `VPN` + "`" + `, ` + "`" + `DIRECTCONNECT` + "`" + `, ` + "`" + `PEERCONNECTION` + "`" + `, ` + "`" + `SSLVPN` + "`" + `, ` + "`" + `NAT` + "`" + `, ` + "`" + `NORMAL_CVM` + "`" + `, ` + "`" + `EIP` + "`" + ` and ` + "`" + `CCN` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "route_entry_id",
					Description: `ID of a route table entry.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `ID of the routing table.`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `List of subnet IDs bound to the route table.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the routing table.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_vpc_subnets",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query vpc subnets information.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) Zone of the subnet to be queried.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Optional) Filter subnet with this CIDR.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `(Optional) Filter default or no default subnets.`,
				},
				resource.Attribute{
					Name:        "is_remote_vpc_snat",
					Description: `(Optional) Filter the VPC SNAT address pool subnet.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the subnet to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) ID of the subnet to be queried.`,
				},
				resource.Attribute{
					Name:        "tag_key",
					Description: `(Optional) Filter if subnet has this tag.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of the subnet to be queried.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) ID of the VPC to be queried. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "instance_list",
					Description: `List of subnets.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The availability zone of the subnet.`,
				},
				resource.Attribute{
					Name:        "available_ip_count",
					Description: `The number of available IPs.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `A network address block of the subnet.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the subnet resource.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Indicates whether it is the default subnet of the VPC for this region.`,
				},
				resource.Attribute{
					Name:        "is_multicast",
					Description: `Indicates whether multicast is enabled.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the subnet.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `ID of the routing table.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of the subnet.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the subnet resource.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_list",
					Description: `List of subnets.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The availability zone of the subnet.`,
				},
				resource.Attribute{
					Name:        "available_ip_count",
					Description: `The number of available IPs.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `A network address block of the subnet.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the subnet resource.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Indicates whether it is the default subnet of the VPC for this region.`,
				},
				resource.Attribute{
					Name:        "is_multicast",
					Description: `Indicates whether multicast is enabled.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the subnet.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `ID of the routing table.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of the subnet.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the subnet resource.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_vpn_connections",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of VPN connections.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "customer_gateway_id",
					Description: `(Optional) Customer gateway ID of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the VPN connection. The length of character is limited to 1-60.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of the VPN connection to be queried.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) ID of the VPC.`,
				},
				resource.Attribute{
					Name:        "vpn_gateway_id",
					Description: `(Optional) VPN gateway ID of the VPN connection. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "connection_list",
					Description: `Information list of the dedicated connections.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "customer_gateway_id",
					Description: `ID of the customer gateway.`,
				},
				resource.Attribute{
					Name:        "encrypt_proto",
					Description: `Encrypt proto of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "ike_dh_group_name",
					Description: `DH group name of the IKE operation specification.`,
				},
				resource.Attribute{
					Name:        "ike_exchange_mode",
					Description: `Exchange mode of the IKE operation specification.`,
				},
				resource.Attribute{
					Name:        "ike_local_address",
					Description: `Local address of the IKE operation specification.`,
				},
				resource.Attribute{
					Name:        "ike_local_fqdn_name",
					Description: `Local FQDN name of the IKE operation specification.`,
				},
				resource.Attribute{
					Name:        "ike_local_identity",
					Description: `Local identity of the IKE operation specification.`,
				},
				resource.Attribute{
					Name:        "ike_proto_authen_algorithm",
					Description: `Proto authenticate algorithm of the IKE operation specification.`,
				},
				resource.Attribute{
					Name:        "ike_proto_encry_algorithm",
					Description: `Proto encrypt algorithm of the IKE operation specification.`,
				},
				resource.Attribute{
					Name:        "ike_remote_address",
					Description: `Remote address of the IKE operation specification.`,
				},
				resource.Attribute{
					Name:        "ike_remote_fqdn_name",
					Description: `Remote FQDN name of the IKE operation specification.`,
				},
				resource.Attribute{
					Name:        "ike_remote_identity",
					Description: `Remote identity of the IKE operation specification.`,
				},
				resource.Attribute{
					Name:        "ike_sa_lifetime_seconds",
					Description: `SA lifetime of the IKE operation specification, unit is ` + "`" + `second` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ike_version",
					Description: `Version of the IKE operation specification.`,
				},
				resource.Attribute{
					Name:        "ipsec_encrypt_algorithm",
					Description: `Encrypt algorithm of the IPSEC operation specification.`,
				},
				resource.Attribute{
					Name:        "ipsec_integrity_algorithm",
					Description: `Integrity algorithm of the IPSEC operation specification.`,
				},
				resource.Attribute{
					Name:        "ipsec_pfs_dh_group",
					Description: `PFS DH group name of the IPSEC operation specification.`,
				},
				resource.Attribute{
					Name:        "ipsec_sa_lifetime_seconds",
					Description: `SA lifetime of the IPSEC operation specification, unit is ` + "`" + `second` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ipsec_sa_lifetime_traffic",
					Description: `SA lifetime traffic of the IPSEC operation specification, unit is ` + "`" + `KB` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "net_status",
					Description: `Net status of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "pre_share_key",
					Description: `Pre-shared key of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "route_type",
					Description: `Route type of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "security_group_policy",
					Description: `Security group policy of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "local_cidr_block",
					Description: `Local cidr block.`,
				},
				resource.Attribute{
					Name:        "remote_cidr_block",
					Description: `Remote cidr block list.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of tags used to associate different resources.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC.`,
				},
				resource.Attribute{
					Name:        "vpn_gateway_id",
					Description: `ID of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "vpn_proto",
					Description: `Vpn proto of the VPN connection.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "connection_list",
					Description: `Information list of the dedicated connections.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "customer_gateway_id",
					Description: `ID of the customer gateway.`,
				},
				resource.Attribute{
					Name:        "encrypt_proto",
					Description: `Encrypt proto of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "ike_dh_group_name",
					Description: `DH group name of the IKE operation specification.`,
				},
				resource.Attribute{
					Name:        "ike_exchange_mode",
					Description: `Exchange mode of the IKE operation specification.`,
				},
				resource.Attribute{
					Name:        "ike_local_address",
					Description: `Local address of the IKE operation specification.`,
				},
				resource.Attribute{
					Name:        "ike_local_fqdn_name",
					Description: `Local FQDN name of the IKE operation specification.`,
				},
				resource.Attribute{
					Name:        "ike_local_identity",
					Description: `Local identity of the IKE operation specification.`,
				},
				resource.Attribute{
					Name:        "ike_proto_authen_algorithm",
					Description: `Proto authenticate algorithm of the IKE operation specification.`,
				},
				resource.Attribute{
					Name:        "ike_proto_encry_algorithm",
					Description: `Proto encrypt algorithm of the IKE operation specification.`,
				},
				resource.Attribute{
					Name:        "ike_remote_address",
					Description: `Remote address of the IKE operation specification.`,
				},
				resource.Attribute{
					Name:        "ike_remote_fqdn_name",
					Description: `Remote FQDN name of the IKE operation specification.`,
				},
				resource.Attribute{
					Name:        "ike_remote_identity",
					Description: `Remote identity of the IKE operation specification.`,
				},
				resource.Attribute{
					Name:        "ike_sa_lifetime_seconds",
					Description: `SA lifetime of the IKE operation specification, unit is ` + "`" + `second` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ike_version",
					Description: `Version of the IKE operation specification.`,
				},
				resource.Attribute{
					Name:        "ipsec_encrypt_algorithm",
					Description: `Encrypt algorithm of the IPSEC operation specification.`,
				},
				resource.Attribute{
					Name:        "ipsec_integrity_algorithm",
					Description: `Integrity algorithm of the IPSEC operation specification.`,
				},
				resource.Attribute{
					Name:        "ipsec_pfs_dh_group",
					Description: `PFS DH group name of the IPSEC operation specification.`,
				},
				resource.Attribute{
					Name:        "ipsec_sa_lifetime_seconds",
					Description: `SA lifetime of the IPSEC operation specification, unit is ` + "`" + `second` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ipsec_sa_lifetime_traffic",
					Description: `SA lifetime traffic of the IPSEC operation specification, unit is ` + "`" + `KB` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "net_status",
					Description: `Net status of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "pre_share_key",
					Description: `Pre-shared key of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "route_type",
					Description: `Route type of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "security_group_policy",
					Description: `Security group policy of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "local_cidr_block",
					Description: `Local cidr block.`,
				},
				resource.Attribute{
					Name:        "remote_cidr_block",
					Description: `Remote cidr block list.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of tags used to associate different resources.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC.`,
				},
				resource.Attribute{
					Name:        "vpn_gateway_id",
					Description: `ID of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "vpn_proto",
					Description: `Vpn proto of the VPN connection.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_vpn_customer_gateways",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of VPN customer gateways.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the VPN customer gateway.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the customer gateway. The length of character is limited to 1-60.`,
				},
				resource.Attribute{
					Name:        "public_ip_address",
					Description: `(Optional) Public ip address of the VPN customer gateway.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of the VPN customer gateway to be queried. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "gateway_list",
					Description: `Information list of the dedicated gateways.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the VPN customer gateway.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the VPN customer gateway.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the VPN customer gateway.`,
				},
				resource.Attribute{
					Name:        "public_ip_address",
					Description: `Public ip address of the VPN customer gateway.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the VPN customer gateway.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "gateway_list",
					Description: `Information list of the dedicated gateways.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the VPN customer gateway.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the VPN customer gateway.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the VPN customer gateway.`,
				},
				resource.Attribute{
					Name:        "public_ip_address",
					Description: `Public ip address of the VPN customer gateway.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the VPN customer gateway.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_vpn_gateways",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of VPN gateways.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the VPN gateway. The length of character is limited to 1-60.`,
				},
				resource.Attribute{
					Name:        "public_ip_address",
					Description: `(Optional) Public ip address of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of the VPN gateway to be queried.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) ID of the VPC.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) Zone of the VPN gateway. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "gateway_list",
					Description: `Information list of the dedicated gateways.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `The maximum public network output bandwidth of VPN gateway (unit: Mbps).`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `Charge Type of the VPN gateway.`,
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
					Name:        "id",
					Description: `ID of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "is_address_blocked",
					Description: `Indicates whether ip address is blocked.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "new_purchase_plan",
					Description: `The plan of new purchase.`,
				},
				resource.Attribute{
					Name:        "prepaid_renew_flag",
					Description: `Flag indicates whether to renew or not.`,
				},
				resource.Attribute{
					Name:        "public_ip_address",
					Description: `Public ip of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "restrict_state",
					Description: `Restrict state of VPN gateway.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of tags used to associate different resources.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of gateway instance.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `Zone of the VPN gateway.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "gateway_list",
					Description: `Information list of the dedicated gateways.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `The maximum public network output bandwidth of VPN gateway (unit: Mbps).`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `Charge Type of the VPN gateway.`,
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
					Name:        "id",
					Description: `ID of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "is_address_blocked",
					Description: `Indicates whether ip address is blocked.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "new_purchase_plan",
					Description: `The plan of new purchase.`,
				},
				resource.Attribute{
					Name:        "prepaid_renew_flag",
					Description: `Flag indicates whether to renew or not.`,
				},
				resource.Attribute{
					Name:        "public_ip_address",
					Description: `Public ip of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "restrict_state",
					Description: `Restrict state of VPN gateway.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of tags used to associate different resources.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of gateway instance.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `Zone of the VPN gateway.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"tencentcloud_address_template_groups":                  0,
		"tencentcloud_address_templates":                        1,
		"tencentcloud_api_gateway_api_keys":                     2,
		"tencentcloud_api_gateway_apis":                         3,
		"tencentcloud_api_gateway_customer_domains":             4,
		"tencentcloud_api_gateway_ip_strategies":                5,
		"tencentcloud_api_gateway_services":                     6,
		"tencentcloud_api_gateway_throttling_apis":              7,
		"tencentcloud_api_gateway_throttling_services":          8,
		"tencentcloud_api_gateway_usage_plan_environments":      9,
		"tencentcloud_api_gateway_usage_plans":                  10,
		"tencentcloud_as_scaling_configs":                       11,
		"tencentcloud_as_scaling_groups":                        12,
		"tencentcloud_as_scaling_policies":                      13,
		"tencentcloud_audit_cos_regions":                        14,
		"tencentcloud_audit_key_alias":                          15,
		"tencentcloud_audits":                                   16,
		"tencentcloud_availability_regions":                     17,
		"tencentcloud_availability_zones":                       18,
		"tencentcloud_cam_group_memberships":                    19,
		"tencentcloud_cam_group_policy_attachments":             20,
		"tencentcloud_cam_groups":                               21,
		"tencentcloud_cam_policies":                             22,
		"tencentcloud_cam_role_policy_attachments":              23,
		"tencentcloud_cam_roles":                                24,
		"tencentcloud_cam_saml_providers":                       25,
		"tencentcloud_cam_user_policy_attachments":              26,
		"tencentcloud_cam_users":                                27,
		"tencentcloud_cbs_snapshot_policies":                    28,
		"tencentcloud_cbs_snapshots":                            29,
		"tencentcloud_cbs_storages":                             30,
		"tencentcloud_ccn_bandwidth_limits":                     31,
		"tencentcloud_ccn_instances":                            32,
		"tencentcloud_cdh_instances":                            33,
		"tencentcloud_cdn_domains":                              34,
		"tencentcloud_cfs_access_groups":                        35,
		"tencentcloud_cfs_access_rules":                         36,
		"tencentcloud_cfs_file_systems":                         37,
		"tencentcloud_ckafka_acls":                              38,
		"tencentcloud_ckafka_topics":                            39,
		"tencentcloud_ckafka_users":                             40,
		"tencentcloud_clb_attachments":                          41,
		"tencentcloud_clb_instances":                            42,
		"tencentcloud_clb_listener_rules":                       43,
		"tencentcloud_clb_listeners":                            44,
		"tencentcloud_clb_redirections":                         45,
		"tencentcloud_clb_target_groups":                        46,
		"tencentcloud_container_cluster_instances":              47,
		"tencentcloud_container_clusters":                       48,
		"tencentcloud_cos_bucket_object":                        49,
		"tencentcloud_cos_buckets":                              50,
		"tencentcloud_cynosdb_clusters":                         51,
		"tencentcloud_cynosdb_instances":                        52,
		"tencentcloud_dayu_cc_http_policies":                    53,
		"tencentcloud_dayu_cc_https_policies":                   54,
		"tencentcloud_dayu_ddos_policies":                       55,
		"tencentcloud_dayu_ddos_policy_attachments":             56,
		"tencentcloud_dayu_ddos_policy_cases":                   57,
		"tencentcloud_dayu_l4_rules":                            58,
		"tencentcloud_dayu_l7_rules":                            59,
		"tencentcloud_dc_gateway_ccn_routes":                    60,
		"tencentcloud_dc_gateway_instances":                     61,
		"tencentcloud_dc_instances":                             62,
		"tencentcloud_dcx_instances":                            63,
		"tencentcloud_dnats":                                    64,
		"tencentcloud_eip":                                      65,
		"tencentcloud_eips":                                     66,
		"tencentcloud_elasticsearch_instances":                  67,
		"tencentcloud_enis":                                     68,
		"tencentcloud_gaap_certificates":                        69,
		"tencentcloud_gaap_domain_error_pages":                  70,
		"tencentcloud_gaap_http_domains":                        71,
		"tencentcloud_gaap_http_rules":                          72,
		"tencentcloud_gaap_layer4_listeners":                    73,
		"tencentcloud_gaap_layer7_listeners":                    74,
		"tencentcloud_gaap_proxies":                             75,
		"tencentcloud_gaap_realservers":                         76,
		"tencentcloud_gaap_security_policies":                   77,
		"tencentcloud_gaap_security_rules":                      78,
		"tencentcloud_ha_vip_eip_attachments":                   79,
		"tencentcloud_ha_vips":                                  80,
		"tencentcloud_image":                                    81,
		"tencentcloud_images":                                   82,
		"tencentcloud_instance_types":                           83,
		"tencentcloud_instances":                                84,
		"tencentcloud_key_pairs":                                85,
		"tencentcloud_kms_keys":                                 86,
		"tencentcloud_kubernetes_clusters":                      87,
		"tencentcloud_mongodb_instances":                        88,
		"tencentcloud_mongodb_zone_config":                      89,
		"tencentcloud_monitor_binding_objects":                  90,
		"tencentcloud_monitor_data":                             91,
		"tencentcloud_monitor_policy_conditions":                92,
		"tencentcloud_monitor_policy_groups":                    93,
		"tencentcloud_monitor_product_event":                    94,
		"tencentcloud_monitor_product_namespace":                95,
		"tencentcloud_mysql_backup_list":                        96,
		"tencentcloud_mysql_instance":                           97,
		"tencentcloud_mysql_parameter_list":                     98,
		"tencentcloud_mysql_zone_config":                        99,
		"tencentcloud_nat_gateways":                             100,
		"tencentcloud_nats":                                     101,
		"tencentcloud_placement_groups":                         102,
		"tencentcloud_postgresql_instances":                     103,
		"tencentcloud_postgresql_specinfos":                     104,
		"tencentcloud_protocol_template_groups":                 105,
		"tencentcloud_protocol_templates":                       106,
		"tencentcloud_redis_instances":                          107,
		"tencentcloud_redis_zone_config":                        108,
		"tencentcloud_reserved_instance_configs":                109,
		"tencentcloud_reserved_instances":                       110,
		"tencentcloud_route_table":                              111,
		"tencentcloud_scf_functions":                            112,
		"tencentcloud_scf_logs":                                 113,
		"tencentcloud_scf_namespaces":                           114,
		"tencentcloud_security_group":                           115,
		"tencentcloud_security_groups":                          116,
		"tencentcloud_sqlserver_account_db_attachments":         117,
		"tencentcloud_sqlserver_accounts":                       118,
		"tencentcloud_sqlserver_backups":                        119,
		"tencentcloud_sqlserver_basic_instances":                120,
		"tencentcloud_sqlserver_dbs":                            121,
		"tencentcloud_sqlserver_instances":                      122,
		"tencentcloud_sqlserver_publish_subscribes":             123,
		"tencentcloud_sqlserver_readonly_groups":                124,
		"tencentcloud_sqlserver_zone_config":                    125,
		"tencentcloud_ssl_certificates":                         126,
		"tencentcloud_ssm_secret_versions":                      127,
		"tencentcloud_ssm_secrets":                              128,
		"tencentcloud_subnet":                                   129,
		"tencentcloud_tcaplus_clusters":                         130,
		"tencentcloud_tcaplus_idls":                             131,
		"tencentcloud_tcaplus_tablegroups":                      132,
		"tencentcloud_tcaplus_tables":                           133,
		"tencentcloud_tcr_instances":                            134,
		"tencentcloud_tcr_namespaces":                           135,
		"tencentcloud_tcr_repositories":                         136,
		"tencentcloud_tcr_tokens":                               137,
		"tencentcloud_tcr_vpc_attachments":                      138,
		"tencentcloud_vod_adaptive_dynamic_streaming_templates": 139,
		"tencentcloud_vod_image_sprite_templates":               140,
		"tencentcloud_vod_procedure_templates":                  141,
		"tencentcloud_vod_snapshot_by_time_offset_templates":    142,
		"tencentcloud_vod_super_player_configs":                 143,
		"tencentcloud_vpc":                                      144,
		"tencentcloud_vpc_acls":                                 145,
		"tencentcloud_vpc_instances":                            146,
		"tencentcloud_vpc_route_tables":                         147,
		"tencentcloud_vpc_subnets":                              148,
		"tencentcloud_vpn_connections":                          149,
		"tencentcloud_vpn_customer_gateways":                    150,
		"tencentcloud_vpn_gateways":                             151,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
