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
			Category:         "Actiontrail Resources",
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
			Type:             "alicloud_api_gateway_api",
			Category:         "Api Gateway Resources",
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
			Category:         "Api Gateway Resources",
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
					Description: `(Required) The name of the app. Defaults to null.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the app. Defaults to null. ## Attributes Reference The following attributes are exported:`,
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
			Category:         "Api Gateway Resources",
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
			Category:         "Api Gateway Resources",
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
					Description: `The ID of the api group of api gateway. ## Import Api gateway group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_api_gateway_group.example "ab2351f2ce904edaa8d92a0510832b91" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the api group of api gateway. ## Import Api gateway group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_api_gateway_group.example "ab2351f2ce904edaa8d92a0510832b91" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_api_gateway_vpc_access",
			Category:         "Api Gateway Resources",
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
			Type:             "alicloud_cas_certificate",
			Category:         "Cas Certificates",
			ShortDescription: `Provides a CAS Certificate resource.`,
			Description:      ``,
			Keywords: []string{
				"cas",
				"certificates",
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
			Type:             "alicloud_cdn_domain",
			Category:         "CDN Resources",
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
			Category:         "CDN Resources",
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
			Category:         "CDN Resources",
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
					Description: `(Optional, Type: list, Available in 1.52.0+) Certificate config of the accelerated domain. It's a list and consist of at most 1 item. ### Block sources The ` + "`" + `sources` + "`" + ` block supports the following:`,
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
					Description: `(Optional) The SSL certificate type, can be "upload", "cas" and "free". ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The cdn domain id. The value is same as the domain name. ## Import CDN domain can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import alicloud_cdn_domain_new.example cdn-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The cdn domain id. The value is same as the domain name. ## Import CDN domain can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import alicloud_cdn_domain_new.example cdn-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cen_bandwidth_limit",
			Category:         "CEN Resources",
			ShortDescription: `Provides a Alicloud CEN cross-regional interconnection bandwidth configuration resource.`,
			Description:      ``,
			Keywords: []string{
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
			Category:         "CEN Resources",
			ShortDescription: `Provides a Alicloud CEN bandwidth package resource.`,
			Description:      ``,
			Keywords: []string{
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
					Description: `(Required) List of the two areas to connect. Valid value: China | North-America | Asia-Pacific | Europe | Middle-East.`,
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
			Category:         "CEN Resources",
			ShortDescription: `Provides a Alicloud CEN bandwidth package attachment resource.`,
			Description:      ``,
			Keywords: []string{
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
			Type:             "alicloud_cen_instance",
			Category:         "CEN Resources",
			ShortDescription: `Provides a Alicloud CEN instance resource.`,
			Description:      ``,
			Keywords: []string{
				"cen",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the CEN instance. Defaults to null.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the CEN instance. Defaults to null. ### Timeouts ->`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 3 mins) Used when creating the cen instance (until it reaches the initial ` + "`" + `Active` + "`" + ` status).`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 3 mins) Used when terminating the cen instance. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the CEN instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the CEN instance.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the CEN instance. ## Import CEN instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cen_instance.example cen-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the CEN instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the CEN instance.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the CEN instance. ## Import CEN instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cen_instance.example cen-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cen_instance_attachment",
			Category:         "CEN Resources",
			ShortDescription: `Provides a Alicloud CEN child instance attachment resource.`,
			Description:      ``,
			Keywords: []string{
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
			Category:         "CEN Resources",
			ShortDescription: `Provides a Alicloud CEN child instance grant resource.`,
			Description:      ``,
			Keywords: []string{
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
			Type:             "alicloud_cen_route_entry",
			Category:         "CEN Resources",
			ShortDescription: `Provides a Alicloud CEN manage route entried resource.`,
			Description:      ``,
			Keywords: []string{
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
			Type:             "alicloud_cms_alarm",
			Category:         "Cloud Monitor Resources",
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
					Name:        "start_time",
					Description: `Start time of the alarm effective period. Default to 0 and it indicates the time 00:00. Valid value range: [0, 24].`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `End time of the alarm effective period. Default value 24 and it indicates the time 24:00. Valid value range: [0, 24].`,
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
					Name:        "name",
					Description: `The alarm name.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `Monitor project name.`,
				},
				resource.Attribute{
					Name:        "metric",
					Description: `Name of the monitoring metrics.`,
				},
				resource.Attribute{
					Name:        "dimensions",
					Description: `Map of the resources associated with the alarm rule.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `Index query cycle.`,
				},
				resource.Attribute{
					Name:        "statistics",
					Description: `Statistical method.`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `Alarm comparison operator.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `Alarm threshold value.`,
				},
				resource.Attribute{
					Name:        "triggered_count",
					Description: `Number of trigger alarm.`,
				},
				resource.Attribute{
					Name:        "contact_groups",
					Description: `List contact groups of the alarm rule.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `Start time of the alarm effective period.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `End time of the alarm effective period.`,
				},
				resource.Attribute{
					Name:        "silence_time",
					Description: `Notification silence period in the alarm state.`,
				},
				resource.Attribute{
					Name:        "notify_type",
					Description: `Notification type.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Whether to enable alarm rule.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current alarm rule status.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the alarm rule.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The alarm name.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `Monitor project name.`,
				},
				resource.Attribute{
					Name:        "metric",
					Description: `Name of the monitoring metrics.`,
				},
				resource.Attribute{
					Name:        "dimensions",
					Description: `Map of the resources associated with the alarm rule.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `Index query cycle.`,
				},
				resource.Attribute{
					Name:        "statistics",
					Description: `Statistical method.`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `Alarm comparison operator.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `Alarm threshold value.`,
				},
				resource.Attribute{
					Name:        "triggered_count",
					Description: `Number of trigger alarm.`,
				},
				resource.Attribute{
					Name:        "contact_groups",
					Description: `List contact groups of the alarm rule.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `Start time of the alarm effective period.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `End time of the alarm effective period.`,
				},
				resource.Attribute{
					Name:        "silence_time",
					Description: `Notification silence period in the alarm state.`,
				},
				resource.Attribute{
					Name:        "notify_type",
					Description: `Notification type.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Whether to enable alarm rule.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current alarm rule status.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_common_bandwidth_package",
			Category:         "VPC Resources",
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
					Description: `(Optional) Ratio of the common bandwidth package. It is valid when ` + "`" + `internet_charge_type` + "`" + ` is ` + "`" + `PayBy95` + "`" + `. Default to 100. Valid values: [10-100].`,
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
			Category:         "VPC Resources",
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
			Category:         "Container Resources",
			ShortDescription: `Provides a Alicloud container cluster resource.`,
			Description:      ``,
			Keywords: []string{
				"container",
				"cluster",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cr_namespace",
			Category:         "Container Registry Resources",
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
			Category:         "Container Registry Resources",
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
			Category:         "Container Resources",
			ShortDescription: `Provides a resource to deploy application in one container cluster.`,
			Description:      ``,
			Keywords: []string{
				"container",
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
			Category:         "Container Resources",
			ShortDescription: `Provides a Alicloud resource to manage container kubernetes cluster.`,
			Description:      ``,
			Keywords: []string{
				"container",
				"cs",
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
					Name:        "force_update",
					Description: `(Optional, Available in 1.50.0+) Whether to force the update of kubernetes cluster arguments. Default to false.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional, ForceNew) The Zone where new kubernetes cluster will be located. If it is not be specified, the ` + "`" + `vswitch_ids` + "`" + ` should be set, its value will be vswitch's zone.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Deprecated from version 1.16.0)(Force new resource) The vswitch where new kubernetes cluster will be located. If it is not specified, a new VPC and VSwicth will be built. It must be in the zone which ` + "`" + `availability_zone` + "`" + ` specified.`,
				},
				resource.Attribute{
					Name:        "vswitch_ids",
					Description: `(Optional, ForceNew) The vswitch where new kubernetes cluster will be located. For SingleAZ Cluster, if it is not specified, a new VPC and VSwicth will be built. It must be in the zone which ` + "`" + `availability_zone` + "`" + ` specified. For MultiAZ Cluster, you must create three vswitches firstly, specify them here.`,
				},
				resource.Attribute{
					Name:        "new_nat_gateway",
					Description: `(Optional, ForceNew) Whether to create a new nat gateway while creating kubernetes cluster. Default to true.`,
				},
				resource.Attribute{
					Name:        "master_instance_type",
					Description: `(Deprecated from version 1.16.0)(Required, Force new resource) The instance type of master node.`,
				},
				resource.Attribute{
					Name:        "master_instance_types",
					Description: `(Required, ForceNew) The instance type of master node. Specify one type for single AZ Cluster, three types for MultiAZ Cluster. You can get the available kubetnetes master node instance types by [datasource instance_types](https://www.terraform.io/docs/providers/alicloud/d/instance_types.html#kubernetes_node_role)`,
				},
				resource.Attribute{
					Name:        "worker_instance_type",
					Description: `(Deprecated from version 1.16.0)(Required, Force new resource) The instance type of worker node.`,
				},
				resource.Attribute{
					Name:        "worker_instance_types",
					Description: `(Required, ForceNew) The instance type of worker node. Specify one type for single AZ Cluster, three types for MultiAZ Cluster. You can get the available kubetnetes master node instance types by [datasource instance_types](https://www.terraform.io/docs/providers/alicloud/d/instance_types.html#kubernetes_node_role)`,
				},
				resource.Attribute{
					Name:        "worker_number",
					Description: `(Required) The worker node number of the kubernetes cluster. Default to 3. It is limited up to 50 and if you want to enlarge it, please apply white list or contact with us.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional, ForceNew) The password of ssh login cluster node. You have to specify one of ` + "`" + `password` + "`" + ` and ` + "`" + `key_name` + "`" + ` fields.`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `(Optional, ForceNew) The keypair of ssh login cluster node, you have to create it first.`,
				},
				resource.Attribute{
					Name:        "user_ca",
					Description: `(Optional, ForceNew) The path of customized CA cert, you can use this CA to sign client certs to connect your cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_network_type",
					Description: `(Required, ForceNew) The network that cluster uses, use ` + "`" + `flannel` + "`" + ` or ` + "`" + `terway` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "pod_cidr",
					Description: `(Optional, ForceNew) The CIDR block for the pod network. It will be allocated automatically when ` + "`" + `vswitch_ids` + "`" + ` is not specified. It cannot be duplicated with the VPC CIDR and CIDR used by Kubernetes cluster in VPC, cannot be modified after creation. Maximum number of hosts allowed in the cluster: 256. Refer to [Plan Kubernetes CIDR blocks under VPC](https://www.alibabacloud.com/help/doc-detail/64530.htm).`,
				},
				resource.Attribute{
					Name:        "service_cidr",
					Description: `(Optional, ForceNew) The CIDR block for the service network. It will be allocated automatically when ` + "`" + `vswitch_id` + "`" + ` is not specified. It cannot be duplicated with the VPC CIDR and CIDR used by Kubernetes cluster in VPC, cannot be modified after creation.`,
				},
				resource.Attribute{
					Name:        "master_instance_charge_type",
					Description: `(Optional, ForceNew) Master payment type. ` + "`" + `PrePaid` + "`" + ` or ` + "`" + `PostPaid` + "`" + `, defaults to ` + "`" + `PostPaid` + "`" + `.`,
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
					Name:        "node_cidr_mask",
					Description: `(Optional, Force new resource) The network mask used on pods for each node, ranging from ` + "`" + `24` + "`" + ` to ` + "`" + `28` + "`" + `. Larger this number is, less pods can be allocated on each node. Default value is ` + "`" + `24` + "`" + `, means you can allocate 256 pods on each node.`,
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
					Name:        "enable_ssh",
					Description: `(Optional, ForceNew) Whether to allow to SSH login kubernetes. Default to false.`,
				},
				resource.Attribute{
					Name:        "slb_internet_enabled",
					Description: `(Optional, ForceNew) Whether to create internet load balancer for API Server. Default to true.`,
				},
				resource.Attribute{
					Name:        "master_disk_category",
					Description: `(Optional, ForceNew) The system disk category of master node. Its valid value are ` + "`" + `cloud_ssd` + "`" + ` and ` + "`" + `cloud_efficiency` + "`" + `. Default to ` + "`" + `cloud_efficiency` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "master_disk_size",
					Description: `(Optional, ForceNew) The system disk size of master node. Its valid value range [20~32768] in GB. Default to 20.`,
				},
				resource.Attribute{
					Name:        "worker_disk_category",
					Description: `(Optional, ForceNew) The system disk category of worker node. Its valid value are ` + "`" + `cloud_ssd` + "`" + ` and ` + "`" + `cloud_efficiency` + "`" + `. Default to ` + "`" + `cloud_efficiency` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "worker_disk_size",
					Description: `(Optional, ForceNew) The system disk size of worker node. Its valid value range [20~32768] in GB. Default to 20.`,
				},
				resource.Attribute{
					Name:        "worker_data_disk_size",
					Description: `(Optional, ForceNew) The data disk size of worker node. Its valid value range [20~32768] in GB. When ` + "`" + `worker_data_disk_category` + "`" + ` is presented, it defaults to 40.`,
				},
				resource.Attribute{
					Name:        "worker_data_disk_category",
					Description: `(Optional, ForceNew) The data disk category of worker node. Its valid value are ` + "`" + `cloud_ssd` + "`" + ` and ` + "`" + `cloud_efficiency` + "`" + `, if not set, data disk will not be created.`,
				},
				resource.Attribute{
					Name:        "install_cloud_monitor",
					Description: `(Optional, ForceNew) Whether to install cloud monitor for the kubernetes' node.`,
				},
				resource.Attribute{
					Name:        "is_outdated",
					Description: `(Optional) Whether to use outdated instance type. Default to false.`,
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
					Description: `(Optional) The path of cluster ca certificate, like ` + "`" + `~/.kube/cluster-ca-cert.pem` + "`" + ` ## Attributes Reference The following attributes are exported:`,
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
					Description: `Map of kubernetes cluster connection information. It contains several attributes to ` + "`" + `Block Connections` + "`" + `. ### Block Nodes`,
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
					Description: `Service Access Domain. ## Import Kubernetes cluster can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cs_kubernetes.main ce4273f9156874b46bb ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `Map of kubernetes cluster connection information. It contains several attributes to ` + "`" + `Block Connections` + "`" + `. ### Block Nodes`,
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
					Description: `Service Access Domain. ## Import Kubernetes cluster can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cs_kubernetes.main ce4273f9156874b46bb ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cs_managed_kubernetes",
			Category:         "Container Resources",
			ShortDescription: `Provides a Alicloud resource to manage container managed kubernetes cluster.`,
			Description:      ``,
			Keywords: []string{
				"container",
				"cs",
				"managed",
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
					Name:        "availability_zone",
					Description: `(Optional, ForceNew) The Zone where new kubernetes cluster will be located. If it is not be specified, the ` + "`" + `vswitch_ids` + "`" + ` should be set, the value will be vswitch's zone.`,
				},
				resource.Attribute{
					Name:        "vswitch_ids",
					Description: `(Optional, ForceNew) The vswitch where new kubernetes cluster will be located. Specify one vswitch's id, if it is not specified, a new VPC and VSwicth will be built. It must be in the zone which ` + "`" + `availability_zone` + "`" + ` specified.`,
				},
				resource.Attribute{
					Name:        "new_nat_gateway",
					Description: `(Optional) Whether to create a new nat gateway while creating kubernetes cluster. Default to true.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required, ForceNew) The password of ssh login cluster node. You have to specify one of ` + "`" + `password` + "`" + ` and ` + "`" + `key_name` + "`" + ` fields.`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `(Required, ForceNew) The keypair of ssh login cluster node, you have to create it first.`,
				},
				resource.Attribute{
					Name:        "pod_cidr",
					Description: `(Optional, ForceNew) The CIDR block for the pod network. It will be allocated automatically when ` + "`" + `vswitch_ids` + "`" + ` is not specified. It cannot be duplicated with the VPC CIDR and CIDR used by Kubernetes cluster in VPC, cannot be modified after creation. Maximum number of hosts allowed in the cluster: 256. Refer to [Plan Kubernetes CIDR blocks under VPC](https://www.alibabacloud.com/help/doc-detail/64530.htm).`,
				},
				resource.Attribute{
					Name:        "service_cidr",
					Description: `(Optional, ForceNew) The CIDR block for the service network. It will be allocated automatically when ` + "`" + `vswitch_id` + "`" + ` is not specified. It cannot be duplicated with the VPC CIDR and CIDR used by Kubernetes cluster in VPC, cannot be modified after creation.`,
				},
				resource.Attribute{
					Name:        "slb_internet_enabled",
					Description: `(Optional, ForceNew) Whether to create internet load balancer for API Server. Default to false.`,
				},
				resource.Attribute{
					Name:        "install_cloud_monitor",
					Description: `(Optional, ForceNew) Whether to install cloud monitor for the kubernetes' node.`,
				},
				resource.Attribute{
					Name:        "worker_disk_size",
					Description: `(Optional, ForceNew) The system disk size of worker node. Its valid value range [20~32768] in GB. Default to 20.`,
				},
				resource.Attribute{
					Name:        "worker_disk_category",
					Description: `(Optional, ForceNew) The system disk category of worker node. Its valid value are ` + "`" + `cloud_ssd` + "`" + ` and ` + "`" + `cloud_efficiency` + "`" + `. Default to ` + "`" + `cloud_efficiency` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "worker_data_disk_size",
					Description: `(Optional, ForceNew) The data disk size of worker node. Its valid value range [20~32768] in GB. When ` + "`" + `worker_data_disk_category` + "`" + ` is presented, it defaults to 40.`,
				},
				resource.Attribute{
					Name:        "worker_data_disk_category",
					Description: `(Optional, ForceNew) The data disk category of worker node. Its valid value are ` + "`" + `cloud_ssd` + "`" + ` and ` + "`" + `cloud_efficiency` + "`" + `, if not set, data disk will not be created.`,
				},
				resource.Attribute{
					Name:        "worker_numbers",
					Description: `(Required) The worker node number of the kubernetes cluster. Default to [3]. It is limited up to 50 and if you want to enlarge it, please apply white list or contact with us.`,
				},
				resource.Attribute{
					Name:        "worker_instance_types",
					Description: `(Required, ForceNew) The instance type of worker node. Specify one type for single AZ Cluster, three types for MultiAZ Cluster. You can get the available kubetnetes master node instance types by [datasource instance_types](https://www.terraform.io/docs/providers/alicloud/d/instance_types.html#kubernetes_node_role)`,
				},
				resource.Attribute{
					Name:        "worker_instance_charge_type",
					Description: `(Optional, ForceNew) Worker payment type. ` + "`" + `PrePaid` + "`" + ` or ` + "`" + `PostPaid` + "`" + `, defaults to ` + "`" + `PostPaid` + "`" + `.`,
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
					Name:        "cluster_network_type",
					Description: `(Optional, ForceNew) The network that cluster uses, use ` + "`" + `flannel` + "`" + ` or ` + "`" + `terway` + "`" + `.`,
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
					Description: `(Optional) The path of cluster ca certificate, like ` + "`" + `~/.kube/cluster-ca-cert.pem` + "`" + ` ## Attributes Reference The following attributes are exported:`,
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
					Name:        "key_name",
					Description: `The keypair of ssh login cluster node, you have to create it first.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of VPC where the current cluster is located.`,
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
					Name:        "worker_disk_size",
					Description: `The system disk size of worker node.`,
				},
				resource.Attribute{
					Name:        "worker_disk_category",
					Description: `The system disk category of worker node.`,
				},
				resource.Attribute{
					Name:        "worker_data_disk_size",
					Description: `The data disk category of worker node.`,
				},
				resource.Attribute{
					Name:        "worker_data_disk_category",
					Description: `The data disk size of worker node.`,
				},
				resource.Attribute{
					Name:        "worker_nodes",
					Description: `List of cluster worker nodes. It contains several attributes to ` + "`" + `Block Nodes` + "`" + `. ### Block Nodes`,
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
					Description: `The private IP address of node. ## Import Managed Kubernetes cluster can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cs_managed_kubernetes.main ce4273f9156874b46bb ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "key_name",
					Description: `The keypair of ssh login cluster node, you have to create it first.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of VPC where the current cluster is located.`,
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
					Name:        "worker_disk_size",
					Description: `The system disk size of worker node.`,
				},
				resource.Attribute{
					Name:        "worker_disk_category",
					Description: `The system disk category of worker node.`,
				},
				resource.Attribute{
					Name:        "worker_data_disk_size",
					Description: `The data disk category of worker node.`,
				},
				resource.Attribute{
					Name:        "worker_data_disk_category",
					Description: `The data disk size of worker node.`,
				},
				resource.Attribute{
					Name:        "worker_nodes",
					Description: `List of cluster worker nodes. It contains several attributes to ` + "`" + `Block Nodes` + "`" + `. ### Block Nodes`,
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
					Description: `The private IP address of node. ## Import Managed Kubernetes cluster can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_cs_managed_kubernetes.main ce4273f9156874b46bb ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cs_swarm",
			Category:         "Container Resources",
			ShortDescription: `Provides a Alicloud container swarm cluster resource.`,
			Description:      ``,
			Keywords: []string{
				"container",
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
					Description: `(Required, ForceNew) The password of ECS instance node.`,
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
			Category:         "Datahub Service Resources",
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
			Category:         "Datahub Service Resources",
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
			Category:         "Datahub Service Resources",
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
					Description: `(Optional) The number of shards this topic contains. The permitted range of values is [1, 10]. The default value is 1.`,
				},
				resource.Attribute{
					Name:        "life_cycle",
					Description: `(Optional) How many days this topic lives. The permitted range of values is [1, 7]. The default value is 3.`,
				},
				resource.Attribute{
					Name:        "record_type",
					Description: `(Optional) The type of this topic. Its value must be one of {BLOB, TUPLE}. For BLOB topic, data will be organized as binary and encoded by BASE64. For TUPLE topic, data has fixed schema. The default value is "TUPLE" with a schema {STRING}.`,
				},
				resource.Attribute{
					Name:        "record_schema",
					Description: `(Optional) Schema of this topic, required only for TUPLE topic. Supported data types (case-insensitive) are: - BIGINT - STRING - BOOLEAN - DOUBLE - TIMESTAMP`,
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
			Category:         "RDS Resources",
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
					Description: `(Required) Operation password. It may consist of letters, digits, or underlines, with a length of 6 to 32 characters.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional, ForceNew)Privilege type of account. - Normal: Common privilege. - Super: High privilege. - defalut Normal to MySQL、MariaDB、SQL Server(exclude SQL Server 2017 clustered edition). - defalut Super to PostgreSQL, PPAS, SQL Server 2017 clustered edition. Currently, MySQL 5.7, SQL Server 2012/2016, PostgreSQL, and PPAS each can have only one initial account. Other accounts are created by the initial account that has logged on to the database. [Refer to details](https://www.alibabacloud.com/help/doc-detail/26263.htm). ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The current account resource ID. Composed of instance ID and account name with format ` + "`" + `<instance_id>:<name>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The Id of DB instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of DB account.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The account description.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Privilege type of account. ## Import RDS account can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_db_account.example "rm-12345:tf_account" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The current account resource ID. Composed of instance ID and account name with format ` + "`" + `<instance_id>:<name>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The Id of DB instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of DB account.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The account description.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Privilege type of account. ## Import RDS account can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_db_account.example "rm-12345:tf_account" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_db_account_privilege",
			Category:         "RDS Resources",
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
					Description: `The privilege of one account access database. Valid values: ["ReadOnly", "ReadWrite"]. Default to "ReadOnly".`,
				},
				resource.Attribute{
					Name:        "db_names",
					Description: `(Required) List of specified database name. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The current account resource ID. Composed of instance ID, account name and privilege with format ` + "`" + `<instance_id>:<name>:<privilege>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The Id of DB instance.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `The name of DB account.`,
				},
				resource.Attribute{
					Name:        "privilege",
					Description: `The specified account privilege.`,
				},
				resource.Attribute{
					Name:        "db_names",
					Description: `List of granted privilege database names. ## Import RDS account privilege can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_db_account_privilege.example "rm-12345:tf_account:ReadOnly" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The current account resource ID. Composed of instance ID, account name and privilege with format ` + "`" + `<instance_id>:<name>:<privilege>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The Id of DB instance.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `The name of DB account.`,
				},
				resource.Attribute{
					Name:        "privilege",
					Description: `The specified account privilege.`,
				},
				resource.Attribute{
					Name:        "db_names",
					Description: `List of granted privilege database names. ## Import RDS account privilege can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_db_account_privilege.example "rm-12345:tf_account:ReadOnly" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_db_backup_policy",
			Category:         "RDS Resources",
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
					Description: `(Optional) DB Instance backup period. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]. Default to ["Tuesday", "Thursday", "Saturday"].`,
				},
				resource.Attribute{
					Name:        "backup_time",
					Description: `(Optional) DB instance backup time, in the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. Default to "02:00Z-03:00Z". China time is 8 hours behind it.`,
				},
				resource.Attribute{
					Name:        "retention_period",
					Description: `(Optional) Instance backup retention days. Valid values: [7-730]. Default to 7.`,
				},
				resource.Attribute{
					Name:        "log_backup",
					Description: `(Optional) Whether to backup instance log. Note: The 'Basic Edition' category Rds instance does not support setting log backup. [What is Basic Edition](https://www.alibabacloud.com/help/doc-detail/48980.htm).`,
				},
				resource.Attribute{
					Name:        "log_retention_period",
					Description: `(Optional) Instance log backup retention days. Valid when the ` + "`" + `log_backup` + "`" + ` is ` + "`" + `true` + "`" + `. Valid values: [7-730]. Default to 7. It cannot be larger than ` + "`" + `retention_period` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The current backup policy resource ID. It is same as 'instance_id'.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The Id of DB instance.`,
				},
				resource.Attribute{
					Name:        "backup_period",
					Description: `DB Instance backup period.`,
				},
				resource.Attribute{
					Name:        "backup_time",
					Description: `DB instance backup time.`,
				},
				resource.Attribute{
					Name:        "retention_period",
					Description: `Instance backup retention days.`,
				},
				resource.Attribute{
					Name:        "log_backup",
					Description: `Whether to backup instance log.`,
				},
				resource.Attribute{
					Name:        "log_retention_period",
					Description: `Instance log backup retention days. ## Import RDS backup policy can be imported using the id or instance id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_db_backup_policy.example "rm-12345678" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The current backup policy resource ID. It is same as 'instance_id'.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The Id of DB instance.`,
				},
				resource.Attribute{
					Name:        "backup_period",
					Description: `DB Instance backup period.`,
				},
				resource.Attribute{
					Name:        "backup_time",
					Description: `DB instance backup time.`,
				},
				resource.Attribute{
					Name:        "retention_period",
					Description: `Instance backup retention days.`,
				},
				resource.Attribute{
					Name:        "log_backup",
					Description: `Whether to backup instance log.`,
				},
				resource.Attribute{
					Name:        "log_retention_period",
					Description: `Instance log backup retention days. ## Import RDS backup policy can be imported using the id or instance id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_db_backup_policy.example "rm-12345678" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_db_connection",
			Category:         "RDS Resources",
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
			Category:         "RDS Resources",
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
					Description: `(Required) Character set. The value range is limited to the following: - MySQL: [ utf8, gbk, latin1, utf8mb4 ] \(` + "`" + `utf8mb4` + "`" + ` only supports versions 5.5 and 5.6\). - SQLServer: [ Chinese_PRC_CI_AS, Chinese_PRC_CS_AS, SQL_Latin1_General_CP1_CI_AS, SQL_Latin1_General_CP1_CS_AS, Chinese_PRC_BIN ]`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(ForceNew) Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The current database resource ID. Composed of instance ID and database name with format ` + "`" + `<instance_id>:<name>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The Id of DB instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of DB database.`,
				},
				resource.Attribute{
					Name:        "character_set",
					Description: `Character set that database used.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The database description. ## Import RDS database can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_db_database.example "rm-12345:tf_database" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The current database resource ID. Composed of instance ID and database name with format ` + "`" + `<instance_id>:<name>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The Id of DB instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of DB database.`,
				},
				resource.Attribute{
					Name:        "character_set",
					Description: `Character set that database used.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The database description. ## Import RDS database can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_db_database.example "rm-12345:tf_database" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_db_instance",
			Category:         "RDS Resources",
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
					Name:        "db_instance_class",
					Description: `(Deprecated) It has been deprecated from version 1.5.0 and use 'instance_type' to replace.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Required) DB Instance type. For details, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/26312.htm).`,
				},
				resource.Attribute{
					Name:        "db_instance_storage",
					Description: `(Deprecated) It has been deprecated from version 1.5.0 and use 'instance_storage' to replace.`,
				},
				resource.Attribute{
					Name:        "instance_storage",
					Description: `(Required) User-defined DB instance storage space. Value range: - [5, 2000] for MySQL/PostgreSQL/PPAS HA dual node edition; - [20,1000] for MySQL 5.7 basic single node edition; - [10, 2000] for SQL Server 2008R2; - [20,2000] for SQL Server 2012 basic single node edition Increase progressively at a rate of 5 GB. For details, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/26312.htm). Note: There is extra 5 GB storage for SQL Server Instance and it is not in specified ` + "`" + `instance_storage` + "`" + `.`,
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
					Name:        "multi_az",
					Description: `(Optional) It has been deprecated from version 1.8.1, and ` + "`" + `zone_id` + "`" + ` can support multiple zone.`,
				},
				resource.Attribute{
					Name:        "db_instance_net_type",
					Description: `(Deprecated) It has been deprecated from version 1.5.0. If you want to set public connection, please use new resource ` + "`" + `alicloud_db_connection` + "`" + `. Default to Intranet.`,
				},
				resource.Attribute{
					Name:        "allocate_public_connection",
					Description: `(Deprecated) It has been deprecated from version 1.5.0. If you want to allocate public connection string, please use new resource ` + "`" + `alicloud_db_connection` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_network_type",
					Description: `(Deprecated) It has been deprecated from version 1.5.0. If you want to create instances in VPC network, this parameter must be set.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(ForceNew) The virtual switch ID to launch DB instances in one VPC.`,
				},
				resource.Attribute{
					Name:        "master_user_name",
					Description: `(Deprecated) It has been deprecated from version 1.5.0. New resource ` + "`" + `alicloud_db_account` + "`" + ` field 'name' replaces it.`,
				},
				resource.Attribute{
					Name:        "preferred_backup_time",
					Description: `(Deprecated) It has been deprecated from version 1.5.0. New resource ` + "`" + `alicloud_db_backup_policy` + "`" + ` field 'backup_time' replaces it.`,
				},
				resource.Attribute{
					Name:        "backup_retention_period",
					Description: `(Deprecated) It has been deprecated from version 1.5.0. New resource ` + "`" + `alicloud_db_backup_policy` + "`" + ` field 'retention_period' replaces it.`,
				},
				resource.Attribute{
					Name:        "security_ips",
					Description: `(Optional) List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).`,
				},
				resource.Attribute{
					Name:        "db_mappings",
					Description: `(Deprecated) It has been deprecated from version 1.5.0. New resource ` + "`" + `alicloud_db_database` + "`" + ` replaces it.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `(Optional) Set of parameters needs to be set after DB instance was launched. Available parameters can refer to the latest docs [View database parameter templates](https://www.alibabacloud.com/help/doc-detail/26284.htm) .`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) the instance bound to the tag. The format of the incoming value is ` + "`" + `json` + "`" + ` string, including ` + "`" + `TagKey` + "`" + ` and ` + "`" + `TagValue` + "`" + `. ` + "`" + `TagKey` + "`" + ` cannot be null, and ` + "`" + `TagValue` + "`" + ` can be empty, and both cannot begin with ` + "`" + `aliyun` + "`" + `. Format example ` + "`" + `{"key1":"value1"}` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Optional) Input the ECS Security Group ID to join ECS Security Group. Only support mysql 5.5, mysql 5.6 ->`,
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
					Description: `(Defaults to 20 mins) Used when terminating the db instance. ## Import RDS instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_db_instance.example rm-abc12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_db_read_write_splitting_connection",
			Category:         "RDS Resources",
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
			Category:         "RDS Resources",
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
					Description: `(Optional, ForceNew) The virtual switch ID to launch DB instances in one VPC. ->`,
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
			Type:             "alicloud_ddoscoo_instance",
			Category:         "BGP-Line Anti-DDoS Pro Resources",
			ShortDescription: `Provides a Alicloud BGP-line Anti-DDoS Pro(Ddoscoo) Instance Resource.`,
			Description:      ``,
			Keywords: []string{
				"bgp",
				"line",
				"anti",
				"ddos",
				"pro",
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
					Description: `(Required) Base defend bandwidth of the instance. Valid values: 30, 60, 100, 300, 400, 500, 600. The unit is Gbps.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Required) Elastic defend bandwidth of the instance. This value must be larger than the base defend bandwidth. Valid values: 30, 60, 100, 300, 400, 500, 600. The unit is Gbps.`,
				},
				resource.Attribute{
					Name:        "service_bandwidth",
					Description: `(Required) Business bandwidth of the instance. At leaset 100. Increased 100 per step, such as 100, 200, 300. The unit is Mbps.`,
				},
				resource.Attribute{
					Name:        "port_count",
					Description: `(Required) Port retransmission rule count of the instance. At least 50. Increase 5 per step, such as 55, 60, 65.`,
				},
				resource.Attribute{
					Name:        "domain_count",
					Description: `(Required) Domain retransmission rule count of the instance. At least 50. Increase 5 per step, such as 55, 60, 65.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional, ForceNew) The duration that you will buy Ddoscoo instance (in month). Valid values: [1~9], 12, 24, 36. Default to 1. At present, the provider does not support modify "period". ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the instance resource of Ddoscoo.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the instance resource of Ddoscoo.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_disk",
			Category:         "ECS Resources",
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
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `(Optional) If true, the disk will be encrypted, conflict with ` + "`" + `snapshot_id` + "`" + `. ->`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required,ForceNew) The Zone to create the disk in.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The disk name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The disk description.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The disk status.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `(ForceNew) The disk category.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) The disk size.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The disk snapshot ID.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The disk tags.`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `(ForceNew) Whether the disk is encrypted. ## Import Cloud disk can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_disk.example d-abc12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required,ForceNew) The Zone to create the disk in.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The disk name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The disk description.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The disk status.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `(ForceNew) The disk category.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) The disk size.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The disk snapshot ID.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The disk tags.`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `(ForceNew) Whether the disk is encrypted. ## Import Cloud disk can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_disk.example d-abc12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_disk_attachment",
			Category:         "ECS Resources",
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
			Type:             "alicloud_dns",
			Category:         "DNS Resources",
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
					Description: `(Optional) Id of the group in which the domain will add. If not supplied, then use default group. ## Attributes Reference The following attributes are exported:`,
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
			Type:             "alicloud_dns_group",
			Category:         "DNS Resources",
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
			Type:             "alicloud_dns_record",
			Category:         "DNS Resources",
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
					Description: `(Required) The value of domain record.`,
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
					Description: `(Optional) The parsing line of domain record. Valid values are ` + "`" + `default` + "`" + `, ` + "`" + `telecom` + "`" + `, ` + "`" + `unicom` + "`" + `, ` + "`" + `mobile` + "`" + `, ` + "`" + `oversea` + "`" + ` and ` + "`" + `edu` + "`" + `. When the ` + "`" + `type` + "`" + ` is ` + "`" + `FORWORD_URL` + "`" + `, this parameter must be ` + "`" + `default` + "`" + `. Default value is ` + "`" + `default` + "`" + `. ## Attributes Reference The following attributes are exported:`,
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
					Description: `The record parsing line.`,
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
					Description: `The record parsing line.`,
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
			Category:         "DRDS Resources",
			ShortDescription: `Provides an DRDS instance resource.`,
			Description:      ``,
			Keywords: []string{
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
			Type:             "alicloud_eip",
			Category:         "VPC Resources",
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
					Description: `(Optional, ForceNew) Internet charge type of the EIP, Valid values are ` + "`" + `PayByBandwidth` + "`" + `, ` + "`" + `PayByTraffic` + "`" + `. Default to ` + "`" + `PayByBandwidth` + "`" + `. From version ` + "`" + `1.7.1` + "`" + `, default to ` + "`" + `PayByTraffic` + "`" + `.`,
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
					Description: `(Optional, ForceNew, Available in 1.47.0+) The line type of the Elastic IP instance. Default to ` + "`" + `BGP` + "`" + `. Other type of the isp need to open a whitelist. ## Attributes Reference The following attributes are exported:`,
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
			Category:         "VPC Resources",
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
			Category:         "Elasticsearch Resources",
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
					Description: `(Optional, ForceNew) Valid values are ` + "`" + `PrePaid` + "`" + `, ` + "`" + `PostPaid` + "`" + `, Default to ` + "`" + `PostPaid` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) The duration that you will buy Elasticsearch instance (in month). It is valid when instance_charge_type is ` + "`" + `PrePaid` + "`" + `. Valid values: [1~9], 12, 24, 36. Default to 1.`,
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
					Name:        "vswitch_id",
					Description: `(Required, ForceNew) The ID of VSwitch.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) The password of the instance. The password can be 8 to 32 characters in length and must contain three of the following conditions: uppercase letters, lowercase letters, numbers, and special characters (!@#$%^&`,
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
					Name:        "kibana_whitelist",
					Description: `(Optional) Set the Kibana's IP whitelist in internet network.`,
				},
				resource.Attribute{
					Name:        "master_node_spec",
					Description: `(Optional) The dedicated master node spec. If specified, dedicated master node will be created.`,
				},
				resource.Attribute{
					Name:        "zone_count",
					Description: `(Optional, Available in 1.44.0+) The Multi-AZ supported for Elasticsearch, between 1 and 3. The ` + "`" + `data_node_amount` + "`" + ` value must be an integral multiple of the ` + "`" + `zone_count` + "`" + ` value. ### Timeouts ->`,
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
					Description: `The Elasticsearch instance status. Includes ` + "`" + `active` + "`" + `, ` + "`" + `activating` + "`" + `, ` + "`" + `inactive` + "`" + `. Some operations are denied when status is not ` + "`" + `active` + "`" + `. ## Import Elasticsearch can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_elasticsearch.example es-cn-abcde123456 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `The Elasticsearch instance status. Includes ` + "`" + `active` + "`" + `, ` + "`" + `activating` + "`" + `, ` + "`" + `inactive` + "`" + `. Some operations are denied when status is not ` + "`" + `active` + "`" + `. ## Import Elasticsearch can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_elasticsearch.example es-cn-abcde123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ess_alarm",
			Category:         "ESS Resources",
			ShortDescription: `Provides a ESS alarm task resource.`,
			Description:      ``,
			Keywords: []string{
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
					Description: `(Required) The scaling group associated with this alarm.`,
				},
				resource.Attribute{
					Name:        "metric_type",
					Description: `(Optional, ForceNew) The type for the alarm's associated metric. Supported value: system, custom. "system" means the metric data is collected by Aliyun Cloud Monitor Service(CMS), "custom" means the metric data is upload to CMS by users. Defaults to system.`,
				},
				resource.Attribute{
					Name:        "metric_name",
					Description: `(Required) The name for the alarm's associated metric.`,
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
					Description: `(Optional) The dimension map for the alarm's associated metric (documented below). For all metrics, you can not set the dimension key as "scaling_group" or "userId", which is set by default, the second dimension for metric, such as "device" for "PackagesNetIn", need to be set by users. ## Attribute Reference The following attributes are exported:`,
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
			Category:         "ESS Resources",
			ShortDescription: `Provides a ESS Attachment resource to attach or remove ECS instances.`,
			Description:      ``,
			Keywords: []string{
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
			Type:             "alicloud_ess_scaling_configuration",
			Category:         "ESS Resources",
			ShortDescription: `Provides a ESS scaling configuration resource.`,
			Description:      ``,
			Keywords: []string{
				"ess",
				"scaling",
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
					Description: `(Optional) ID of the security group to which a newly created instance belongs.`,
				},
				resource.Attribute{
					Name:        "scaling_configuration_name",
					Description: `(Optional) Name shown for the scheduled task. If this parameter value is not specified, the default value is ScalingConfigurationId.`,
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
					Description: `(Optional) Category of the system disk. The parameter value options are ` + "`" + `ephemeral_ssd` + "`" + `, ` + "`" + `cloud_efficiency` + "`" + `, ` + "`" + `cloud_ssd` + "`" + ` and ` + "`" + `cloud` + "`" + `. ` + "`" + `cloud` + "`" + ` only is used to some no I/O optimized instance. Default to ` + "`" + `cloud_efficiency` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "system_disk_size",
					Description: `(Optional) Size of system disk, in GiB. Optional values: cloud: 40-500, cloud_efficiency: 40-500, cloud_ssd: 40-500, ephemeral_ssd: 40-500 The default value is {40, ImageSize}. If this parameter is set, the system disk size must be greater than or equal to max{40, ImageSize}.`,
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
					Description: `(Optional, Available in 1.46.0+) Indicates whether to overwrite the existing data. Default to false. ->`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) Size of data disk, in GB. The value ranges from 5 to 2,000 for a cloud disk and from 5 to 1,024 for an ephemeral disk. A maximum of four values can be entered.`,
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
			Category:         "ESS Resources",
			ShortDescription: `Provides a ESS scaling group resource.`,
			Description:      ``,
			Keywords: []string{
				"ess",
				"scaling",
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
					Name:        "scaling_group_name",
					Description: `(Optional) Name shown for the scaling group, which must contain 2-40 characters (English or Chinese). If this parameter is not specified, the default value is ScalingGroupId.`,
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
					Description: `(Optional) RemovalPolicy is used to select the ECS instances you want to remove from the scaling group when multiple candidates for removal exist. Optional values: - OldestInstance: removes the first ECS instance attached to the scaling group. - NewestInstance: removes the first ECS instance attached to the scaling group. - OldestScalingConfiguration: removes the ECS instance with the oldest scaling configuration. - Default values: OldestScalingConfiguration and OldestInstance. You can enter up to two removal policies.`,
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
					Description: `(Optional, ForceNew) Multi-AZ scaling group ECS instance expansion and contraction strategy. PRIORITY or BALANCE. ->`,
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
			Category:         "ESS Resources",
			ShortDescription: `Provides a ESS lifecycle hook resource.`,
			Description:      ``,
			Keywords: []string{
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
					Description: `(Optional, ForceNew) The name of the lifecycle hook, which is a string containing 2 to 40 English or Chinese characters. If this parameter value is not specified, the default value is lifecycle hook id.`,
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
			Category:         "ESS Resources",
			ShortDescription: `Provides a ESS scaling rule resource.`,
			Description:      ``,
			Keywords: []string{
				"ess",
				"scaling",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `(Required) ID of the scaling group of a scaling rule.`,
				},
				resource.Attribute{
					Name:        "adjustment_type",
					Description: `(Required) Adjustment mode of a scaling rule. Optional values: - QuantityChangeInCapacity: It is used to increase or decrease a specified number of ECS instances. - PercentChangeInCapacity: It is used to increase or decrease a specified proportion of ECS instances. - TotalCapacity: It is used to adjust the quantity of ECS instances in the current scaling group to a specified value.`,
				},
				resource.Attribute{
					Name:        "adjustment_value",
					Description: `(Required) Adjusted value of a scaling rule. Value range: - QuantityChangeInCapacity：(0, 100] U (-100, 0] - PercentChangeInCapacity：[0, 10000] U [-10000, 0] - TotalCapacity：[0, 100]`,
				},
				resource.Attribute{
					Name:        "scaling_rule_name",
					Description: `(Optional) Name shown for the scaling rule, which is a string containing 2 to 40 English or Chinese characters.`,
				},
				resource.Attribute{
					Name:        "cooldown",
					Description: `(Optional) Cool-down time of a scaling rule. Value range: [0, 86,400], in seconds. The default value is empty. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The scaling rule ID. ## Import ESS scaling rule can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ess_scalingrule.example abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The scaling rule ID. ## Import ESS scaling rule can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ess_scalingrule.example abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ess_schedule",
			Category:         "ESS Resources",
			ShortDescription: `Provides a ESS schedule resource.`,
			Description:      ``,
			Keywords: []string{
				"ess",
				"schedule",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ess_scheduled_task",
			Category:         "ESS Resources",
			ShortDescription: `Provides a ESS schedule resource.`,
			Description:      ``,
			Keywords: []string{
				"ess",
				"scheduled",
				"task",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "scheduled_action",
					Description: `(Required) Operations performed when the scheduled task is triggered. Fill in the unique identifier of the scaling rule.`,
				},
				resource.Attribute{
					Name:        "launch_time",
					Description: `(Required) Operations performed when the scheduled task is triggered. Fill in the unique identifier of the scaling rule.`,
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
					Name:        "launch_expiration_time",
					Description: `(Optional) Time period within which the failed scheduled task is retried. The default value is 600s. Value range: [0, 21600]`,
				},
				resource.Attribute{
					Name:        "recurrence_type",
					Description: `(Optional) Type of the scheduled task to be repeated. RecurrenceType, RecurrenceValue and RecurrenceEndTime must be specified. Optional values: - Daily: Recurrence interval by day for a scheduled task. - Weekly: Recurrence interval by week for a scheduled task. - Monthly: Recurrence interval by month for a scheduled task.`,
				},
				resource.Attribute{
					Name:        "recurrence_value",
					Description: `(Optional) Value of the scheduled task to be repeated. RecurrenceType, RecurrenceValue and RecurrenceEndTime must be specified. - Daily: Only one value in the range [1,31] can be filled. - Weekly: Multiple values can be filled. The values of Sunday to Saturday are 0 to 6 in sequence. Multiple values shall be separated by a comma “,”. - Monthly: In the format of A-B. The value range of A and B is 1 to 31, and the B value must be greater than the A value.`,
				},
				resource.Attribute{
					Name:        "recurrence_end_time",
					Description: `(Optional) End time of the scheduled task to be repeated. The date format follows the ISO8601 standard and uses UTC time. It is in the format of YYYY-MM-DDThh:mmZ. A time point 90 days after creation or modification cannot be entered. RecurrenceType, RecurrenceValue and RecurrenceEndTime must be specified.`,
				},
				resource.Attribute{
					Name:        "task_enabled",
					Description: `(Optional) Whether to enable the scheduled task. The default value is true. ## Attributes Reference The following attributes are exported:`,
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
			Category:         "Function Compute Service Resources",
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
					Description: `(Optional, Available in 1.36.0+) A map that defines environment variables for the function. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the function. The value is formate as ` + "`" + `<service>:<name>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `The date this resource was last modified. ## Import Function Compute function can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_fc_service.foo my-fc-service:hello-world ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the function. The value is formate as ` + "`" + `<service>:<name>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `The date this resource was last modified. ## Import Function Compute function can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_fc_service.foo my-fc-service:hello-world ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_fc_service",
			Category:         "Function Compute Service Resources",
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
					Name:        "last_modified",
					Description: `The date this resource was last modified. ## Import Function Compute Service can be imported using the id or name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_fc_service.foo my-fc-service ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_fc_trigger",
			Category:         "Function Compute Service Resources",
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
					Description: `The date this resource was last modified. ## Import Function Compute trigger can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_fc_service.foo my-fc-service:hello-world:hello-trigger ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the function. The value is formate as ` + "`" + `<service>:<function>:<name>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `The date this resource was last modified. ## Import Function Compute trigger can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_fc_service.foo my-fc-service:hello-world:hello-trigger ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_forward_entry",
			Category:         "VPC Resources",
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
			Category:         "GPDB Resources",
			ShortDescription: `Provides an AnalyticDB for PostgreSQL instance connection resource.`,
			Description:      ``,
			Keywords: []string{
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
					Description: `(Optional) Internet connection port. Valid value: [3200-3999]. Default to 3306. ## Attributes Reference The following attributes are exported:`,
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
			Category:         "GPDB Resources",
			ShortDescription: `Provides a AnalyticDB for PostgreSQL instance resource.`,
			Description:      ``,
			Keywords: []string{
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
					Description: `(Optional) List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]). ## Attributes Reference The following attributes are exported:`,
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
			Type:             "alicloud_instance",
			Category:         "ECS Resources",
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
					Description: `(Required) The type of instance to start.`,
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
					Description: `(Optional) The name of the ECS. This instance_name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://. If not specified, Terraform will autogenerate a default name is ` + "`" + `ECS-Instance` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "allocate_public_ip",
					Description: `(Deprecated) It has been deprecated from version "1.7.0". Setting "internet_max_bandwidth_out" larger than 0 can allocate a public ip address for an instance.`,
				},
				resource.Attribute{
					Name:        "system_disk_category",
					Description: `(Optional) Valid values are ` + "`" + `ephemeral_ssd` + "`" + `, ` + "`" + `cloud_efficiency` + "`" + `, ` + "`" + `cloud_ssd` + "`" + `, ` + "`" + `cloud_essd` + "`" + `, ` + "`" + `cloud` + "`" + `. ` + "`" + `cloud` + "`" + ` only is used to some none I/O optimized instance. Default to ` + "`" + `cloud_efficiency` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "system_disk_size",
					Description: `(Optional) Size of the system disk, measured in GiB. Value range: [20, 500]. The specified value must be equal to or greater than max{20, Imagesize}. Default value: max{40, ImageSize}. ECS instance's system disk can be reset when replacing system disk.`,
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
					Description: `(Optional) Host name of the ECS, which is a string of at least two characters. “hostname” cannot start or end with “.” or “-“. In addition, two or more consecutive “.” or “-“ symbols are not allowed. On Windows, the host name can contain a maximum of 15 characters, which can be a combination of uppercase/lowercase letters, numerals, and “-“. The host name cannot contain dots (“.”) or contain only numeric characters. On other OSs such as Linux, the host name can contain a maximum of 30 characters, which can be segments separated by dots (“.”), where each segment can contain uppercase/lowercase letters, numerals, or “_“. When it is changed, the instance will reboot to make the change take effect.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Password to an instance is a string of 8 to 30 characters. It must contain uppercase/lowercase letters and numerals, but cannot contain special symbols. When it is changed, the instance will reboot to make the change take effect.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Optional) The virtual switch ID to launch in VPC. This parameter must be set unless you can create classic network instances.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `(Optional) Valid values are ` + "`" + `PrePaid` + "`" + `, ` + "`" + `PostPaid` + "`" + `, The default is ` + "`" + `PostPaid` + "`" + `.`,
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
					Description: `(Optional) User-defined data to customize the startup behaviors of an ECS instance and to pass data into an ECS instance.`,
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
					Description: `(Optional) Whether to pre-detection. When it is true, only pre-detection and not actually modify the payment type operation. It is valid when ` + "`" + `instance_charge_type` + "`" + ` is 'PrePaid'. Default to false.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `(Optional) Instance private IP address can be specified when you creating new instance. It is valid when ` + "`" + `vswitch_id` + "`" + ` is specified.`,
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
					Description: `(Optional, true) Whether enable the deletion protection or not. - true: Enable deletion protection. - false: Disable deletion protection. Default to false.`,
				},
				resource.Attribute{
					Name:        "force_delete",
					Description: `(Optional, Available 1.18.0+) If it is true, the "PrePaid" instance will be change to "PostPaid" and then deleted forcibly. However, because of changing instance charge type has CPU core count quota limitation, so strongly recommand that "Don't modify instance charge type frequentlly in one month".`,
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
					Name:        "snapshot_id",
					Description: `(Optional, ForceNew) The snapshot ID used to initialize the data disk. If the size specified by snapshot is greater that the size of the disk, use the size specified by snapshot as the size of the data disk.`,
				},
				resource.Attribute{
					Name:        "delete_with_instance",
					Description: `(Optional, ForceNew) Delete this data disk when the instance is destroyed. It only works on cloud, cloud_efficiency, cloud_essd, cloud_ssd disk. If the category of this data disk was ephemeral_ssd, please don't set this param. Default to true`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, ForceNew) The description of the data disk. ->`,
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
					Name:        "status",
					Description: `The instance status.`,
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
			Category:         "ECS Resources",
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
					Description: `(ForceNew) The name of file to save your new key pair's private key. Strongly suggest you to specified it when you creating key pair, otherwise, you wouldn't get its private key ever. ->`,
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
			Category:         "ECS Resources",
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
			Type:             "alicloud_kms_key",
			Category:         "KMS Resources",
			ShortDescription: `Provides a Alicloud kms key resource.`,
			Description:      ``,
			Keywords: []string{
				"kms",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, ForceNew) The description of the key as viewed in Alicloud console. Default to "From Terraform".`,
				},
				resource.Attribute{
					Name:        "key_usage",
					Description: `(Optional) Specifies the usage of CMK. Currently, default to 'ENCRYPT/DECRYPT', indicating that CMK is used for encryption and decryption.`,
				},
				resource.Attribute{
					Name:        "deletion_window_in_days",
					Description: `(Optional) Duration in days after which the key is deleted after destruction of the resource, must be between 7 and 30 days. Defaults to 30 days.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `(Optional) Specifies whether the key is enabled. Defaults to true. ->`,
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
					Name:        "description",
					Description: `The description of the key.`,
				},
				resource.Attribute{
					Name:        "key_usage",
					Description: `(ForceNew) Specifies the usage of CMK.`,
				},
				resource.Attribute{
					Name:        "deletion_window_in_days",
					Description: `During pre-deletion days.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `Whether the key is enabled. ## Import KMS key can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_kms_key.example abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "description",
					Description: `The description of the key.`,
				},
				resource.Attribute{
					Name:        "key_usage",
					Description: `(ForceNew) Specifies the usage of CMK.`,
				},
				resource.Attribute{
					Name:        "deletion_window_in_days",
					Description: `During pre-deletion days.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `Whether the key is enabled. ## Import KMS key can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_kms_key.example abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_kvstore_backup_policy",
			Category:         "KVStore Resources",
			ShortDescription: `Provides a backup policy for ApsaraDB Redis / Memcache instance resource.`,
			Description:      ``,
			Keywords: []string{
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
			Category:         "KVStore Resources",
			ShortDescription: `Provides an ApsaraDB Redis / Memcache instance resource.`,
			Description:      ``,
			Keywords: []string{
				"kvstore",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Optional) The name of DB instance. It a string of 2 to 256 characters.`,
				},
				resource.Attribute{
					Name:        "instance_class",
					Description: `(Required) Type of the applied ApsaraDB for Redis instance. For more information, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/61135.htm).`,
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
					Name:        "parameters",
					Description: `(Optional) Set of parameters needs to be set after instance was launched. Available parameters can refer to the latest docs [Instance configurations table](https://www.alibabacloud.com/help/doc-detail/61209.htm) . ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The KVStore instance ID.`,
				},
				resource.Attribute{
					Name:        "connection_domain",
					Description: `Instance connection domain (only Intranet access supported). ## Import KVStore instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_kvstore_instance.example r-abc12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The KVStore instance ID.`,
				},
				resource.Attribute{
					Name:        "connection_domain",
					Description: `Instance connection domain (only Intranet access supported). ## Import KVStore instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_kvstore_instance.example r-abc12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_launch_template",
			Category:         "ECS Resources",
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
					Description: `(Optional) Internet bandwidth billing method. Optional values: PayByTraffic.`,
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
			Type:             "alicloud_log_machine_group",
			Category:         "Log Service Resources",
			ShortDescription: `Provides a Alicloud log tail machine group resource.`,
			Description:      ``,
			Keywords: []string{
				"log",
				"service",
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
					Description: `The machine identification type, including IP and user-defined identity. Valid values are "ip" and "userdefined". Default to "ip".`,
				},
				resource.Attribute{
					Name:        "topic",
					Description: `The topic of a machine group. ## Attributes Reference The following attributes are exported:`,
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
			Category:         "Log Service Resources",
			ShortDescription: `Provides a Alicloud log project resource.`,
			Description:      ``,
			Keywords: []string{
				"log",
				"service",
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
			Category:         "Log Service Resources",
			ShortDescription: `Provides a Alicloud log store resource.`,
			Description:      ``,
			Keywords: []string{
				"log",
				"service",
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
			Category:         "Log Service Resources",
			ShortDescription: `Provides a Alicloud log store index resource.`,
			Description:      ``,
			Keywords: []string{
				"log",
				"service",
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
					Description: `Whether the case sensitive. Default to false.`,
				},
				resource.Attribute{
					Name:        "include_chinese",
					Description: `Whether includes the chinese. Default to false.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `The string of several split words, like "\r", "#"`,
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
					Description: `The type of one field. Valid values: ["long", "text", "double", "json"]. Default to "long".`,
				},
				resource.Attribute{
					Name:        "alias",
					Description: `The alias of one field`,
				},
				resource.Attribute{
					Name:        "case_sensitive",
					Description: `Whether the case sensitive for the field. Default to false. It is valid when "type" is "text" or "json".`,
				},
				resource.Attribute{
					Name:        "include_chinese",
					Description: `Whether includes the chinese for the field. Default to false. It is valid when "type" is "text" or "json".`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `The string of several split words, like "\r", "#". It is valid when "type" is "text" or "json".`,
				},
				resource.Attribute{
					Name:        "enable_analytics",
					Description: `Whether to enable field analytics. Default to true. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the log store index. It formats of ` + "`" + `<project>:<logstore>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `The project name.`,
				},
				resource.Attribute{
					Name:        "logstore",
					Description: `Log store name.`,
				},
				resource.Attribute{
					Name:        "full_text",
					Description: `The full text index config.`,
				},
				resource.Attribute{
					Name:        "field_search",
					Description: `The field search index config. ## Import Log store index can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_log_store_index.example tf-log:tf-log-store ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the log store index. It formats of ` + "`" + `<project>:<logstore>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `The project name.`,
				},
				resource.Attribute{
					Name:        "logstore",
					Description: `Log store name.`,
				},
				resource.Attribute{
					Name:        "full_text",
					Description: `The full text index config.`,
				},
				resource.Attribute{
					Name:        "field_search",
					Description: `The field search index config. ## Import Log store index can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_log_store_index.example tf-log:tf-log-store ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_logtail_attachment",
			Category:         "Log Service Resources",
			ShortDescription: `Provides a Alicloud logtail attachment resource.`,
			Description:      ``,
			Keywords: []string{
				"log",
				"service",
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
			Category:         "Log Service Resources",
			ShortDescription: `Provides a Alicloud logtail config resource.`,
			Description:      ``,
			Keywords: []string{
				"log",
				"service",
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
					Description: `(Required) The logtail configure the required JSON files.([Refer to details](https://www.alibabacloud.com/help/doc-detail/29058.htm)) ## Attributes Reference The following attributes are exported:`,
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
			Type:             "alicloud_mns_queue",
			Category:         "MNS  Resources",
			ShortDescription: `Provides a Alicloud MNS Queue resource.`,
			Description:      ``,
			Keywords: []string{
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
			Category:         "MNS  Resources",
			ShortDescription: `Provides a Alicloud MNS Topic resource.`,
			Description:      ``,
			Keywords: []string{
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
			Category:         "MNS  Resources",
			ShortDescription: `Provides a Alicloud MNS Topic Subscription resource.`,
			Description:      ``,
			Keywords: []string{
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
					Description: `(Optional, ForceNew) The NotifyContentFormat attribute of Subscription. This attribute specifies the content format of the messages pushed to users. the attribute has two value SIMPLIFIED or XML.Default value to SIMPLIFIED .`,
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
			Category:         "MongoDB Resources",
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
					Description: `(Optional) Number of replica set nodes. Valid values: [3,5,7]`,
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
					Description: `(Optional, ForceNew) The Zone to launch the DB instance. it supports multiple zone. If it is a multi-zone and ` + "`" + `vswitch_id` + "`" + ` is specified, the vswitch must in one of them. The multiple zone ID can be retrieved by setting ` + "`" + `multi` + "`" + ` to "true" in the data source ` + "`" + `alicloud_zones` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Optional, ForceNew) The virtual switch ID to launch DB instances in one VPC.`,
				},
				resource.Attribute{
					Name:        "account_password",
					Description: `(Optional) Password of the root account. It is a string of 6 to 32 characters and is composed of letters, numbers, and underlines.`,
				},
				resource.Attribute{
					Name:        "security_ip_list",
					Description: `(Optional) List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).`,
				},
				resource.Attribute{
					Name:        "backup_period",
					Description: `(Optional, Available in 1.42.0+) MongoDB Instance backup period. It is required when ` + "`" + `backup_time` + "`" + ` was existed. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]. Default to [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]`,
				},
				resource.Attribute{
					Name:        "backup_time",
					Description: `(Optional, Available in 1.42.0+) MongoDB instance backup time. It is required when ` + "`" + `backup_period` + "`" + ` was existed. In the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. Default to a random time, like "23:00Z-24:00Z". ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the MongoDB.`,
				},
				resource.Attribute{
					Name:        "retention_period",
					Description: `Instance log backup retention days. Available in 1.42.0+. ## Import MongoDB can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_mongodb_instance.example dds-bp1291daeda44194 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the MongoDB.`,
				},
				resource.Attribute{
					Name:        "retention_period",
					Description: `Instance log backup retention days. Available in 1.42.0+. ## Import MongoDB can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_mongodb_instance.example dds-bp1291daeda44194 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodb_sharding_instance",
			Category:         "MongoDB Resources",
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
					Description: `(Optional) Password of the root account. It is a string of 6 to 32 characters and is composed of letters, numbers, and underlines.`,
				},
				resource.Attribute{
					Name:        "security_ip_list",
					Description: `(Optional) List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]). System default to ` + "`" + `["127.0.0.1"]` + "`" + `.`,
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
					Description: `(Optional, Available in 1.42.0+) MongoDB instance backup time. It is required when ` + "`" + `backup_period` + "`" + ` was existed. In the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. Default to a random time, like "23:00Z-24:00Z". ## Attributes Reference The following attributes are exported:`,
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
			Category:         "NAS Resources",
			ShortDescription: `Provides a Alicloud Nas Access Group resource.`,
			Description:      ``,
			Keywords: []string{
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
			Category:         "NAS Resources",
			ShortDescription: `Provides a Alicloud Nas Access Rule resource.`,
			Description:      ``,
			Keywords: []string{
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
					Description: `This ID of this resource. The value is formate as ` + "`" + `<access_group_name>:<access rule id>` + "`" + `. ## Import Nas Access Rule can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_nas_access_rule.foo tf-testAccNasConfigName:1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `This ID of this resource. The value is formate as ` + "`" + `<access_group_name>:<access rule id>` + "`" + `. ## Import Nas Access Rule can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_nas_access_rule.foo tf-testAccNasConfigName:1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_nas_file_system",
			Category:         "NAS Resources",
			ShortDescription: `Provides a Alicloud NAS File System resource.`,
			Description:      ``,
			Keywords: []string{
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
			Category:         "NAS Resources",
			ShortDescription: `Provides a Alicloud Nas MountTarget resource.`,
			Description:      ``,
			Keywords: []string{
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
			Category:         "VPC Resources",
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
			Category:         "VPC Resources",
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
			Category:         "VPC Resources",
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
			Category:         "VPC Resources",
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
			Category:         "ECS Resources",
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
					Description: `(Optional) A mapping of tags to assign to the resource. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ENI ID. ## Import ENI can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_network_interface.eni eni-abc1234567890000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ENI ID. ## Import ENI can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_network_interface.eni eni-abc1234567890000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_network_interface_attachment",
			Category:         "ECS Resources",
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
					Description: `The ID of the resource, formatted as ` + "`" + `<network_interface_id>:<instance_id>` + "`" + `. ## Import Network Interfaces Attachment resource can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_network_interface.eni eni-abc123456789000:i-abc123456789000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the resource, formatted as ` + "`" + `<network_interface_id>:<instance_id>` + "`" + `. ## Import Network Interfaces Attachment resource can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_network_interface.eni eni-abc123456789000:i-abc123456789000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ons_instance",
			Category:         "RocketMQ Resources",
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
			Type:             "alicloud_oss_bucket",
			Category:         "OSS Resources",
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
					Description: `(Optional) The [canned ACL](https://www.alibabacloud.com/help/doc-detail/31898.htm) to apply. Defaults to "private".`,
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
					Description: `(Optional, Available in 1.41.0) Json format text of bucket policy [bucket policy management](https://www.alibabacloud.com/help/doc-detail/100680.htm) (documented below).`,
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
					Description: `(Optional, Type: bool) Allows referer to be empty. Defaults true.`,
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
					Description: `(Required) Object key prefix identifying one or more objects to which the rule applies.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required, Type: bool) Specifies lifecycle rule status.`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `(Optional, Required, Type: set) Specifies a period in the object's expire (documented below). #### Block expiration The lifecycle_rule expiration object supports the following:`,
				},
				resource.Attribute{
					Name:        "date",
					Description: `(Optional) Specifies the date after which you want the corresponding action to take effect. The value obeys ISO8601 format like ` + "`" + `2017-03-09` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "days",
					Description: `(Optional, Type: int) Specifies the number of days after object creation when the specific rule action takes effect. ` + "`" + `NOTE` + "`" + `: One and only one of "date" and "days" can be specified in one expiration configuration. #### Block server-side encryption rule The server-side encryption rule supports the following:`,
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
			Category:         "OSS Resources",
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
					Description: `(Optional) Specifies server-side encryption of the object in OSS. At present, it valid value is "` + "`" + `AES256` + "`" + `". Either ` + "`" + `source` + "`" + ` or ` + "`" + `content` + "`" + ` must be provided to specify the bucket content. These two arguments are mutually-exclusive. ## Attributes Reference The following attributes are exported`,
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
			Category:         "Table Store Resources",
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
			Category:         "Table Store Resources",
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
			Category:         "Table Store Resources",
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
					Description: `(Required, Type: List) The property of ` + "`" + `TableMeta` + "`" + ` which indicates the structure information of a table. It describes the attribute value of primary key. The number of ` + "`" + `primary_key` + "`" + ` should not be less than one and not be more than four.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name for primary key.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required, Type: list) Type for primary key. Only ` + "`" + `Integer` + "`" + `, ` + "`" + `String` + "`" + ` or ` + "`" + `Binary` + "`" + ` is allowed.`,
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
			Type:             "alicloud_pvtz_zone",
			Category:         "Private Zone Resources",
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
					Description: `(Optional) The name of the Private Zone.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `(Optional) The remark of the Private Zone. ## Attributes Reference The following attributes are exported:`,
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
			Category:         "Private Zone Resources",
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
					Description: `(Required) The id List of the VPC, for example:["vpc-1","vpc-2"]. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Private Zone VPC Attachment.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Private Zone VPC Attachment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_pvtz_zone_record",
			Category:         "Private Zone Resources",
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
					Description: `The ID of the Private Zone Record. ## Import Private Zone Record can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_pvtz_zone_record.example abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Private Zone Record. ## Import Private Zone Record can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_pvtz_zone_record.example abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ram_access_key",
			Category:         "RAM Resources",
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
			Category:         "RAM Resources",
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
					Name:        "account_alias",
					Description: `The account alias.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "account_alias",
					Description: `The account alias.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ram_account_password_policy",
			Category:         "RAM Resources",
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
					Name:        "default_minimum_password_length",
					Description: `(Optional, Type: int) Minimal required length of password for a user. Valid value range: [8-32]. Default to 12.`,
				},
				resource.Attribute{
					Name:        "default_require_lowercase_characters",
					Description: `(Optional, Type: bool) Specifies if the occurrence of a lowercase character in the password is mandatory. Default to true.`,
				},
				resource.Attribute{
					Name:        "default_require_uppercase_characters",
					Description: `(Optional, Type: bool) Specifies if the occurrence of an uppercase character in the password is mandatory. Default to true.`,
				},
				resource.Attribute{
					Name:        "default_require_numbers",
					Description: `(Optional, Type: bool) Specifies if the occurrence of a number in the password is mandatory. Default to true.`,
				},
				resource.Attribute{
					Name:        "default_require_symbols",
					Description: `(Optional, Type: bool) Specifies if the occurrence of a special character in the password is mandatory. Default to true.`,
				},
				resource.Attribute{
					Name:        "default_hard_expiry",
					Description: `(Optional, Type: bool) Specifies if a password can expire in a hard way. Default to false.`,
				},
				resource.Attribute{
					Name:        "default_max_password_age",
					Description: `(Optional, Type: int) The number of days after which password expires. A value of 0 indicates that the password never expires. Valid value range: [0-1095]. Default to 0.`,
				},
				resource.Attribute{
					Name:        "default_password_reuse_prevention",
					Description: `(Optional, Type: int) User is not allowed to use the latest number of passwords specified in this parameter. A value of 0 indicates the password history check policy is disabled. Valid value range: [0-24]. Default to 0.`,
				},
				resource.Attribute{
					Name:        "default_max_login_attempts",
					Description: `(Optional, Type: int) Maximum logon attempts with an incorrect password within an hour. Valid value range: [0-32]. Default to 5. ## Import RAM account password policy can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import alicloud_ram_account_password_policy.example ram-account-password-policy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ram_alias",
			Category:         "RAM Resources",
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
			Category:         "RAM Resources",
			ShortDescription: `Provides a RAM Group resource.`,
			Description:      ``,
			Keywords: []string{
				"ram",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Name of the RAM group. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.`,
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
			Category:         "RAM Resources",
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
					Description: `The membership ID.`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `The group name.`,
				},
				resource.Attribute{
					Name:        "user_names",
					Description: `The list of names of users which in the group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The membership ID.`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `The group name.`,
				},
				resource.Attribute{
					Name:        "user_names",
					Description: `The list of names of users which in the group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ram_group_policy_attachment",
			Category:         "RAM Resources",
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
					Description: `The attachment ID. Composed of policy name, policy type and group name with format ` + "`" + `group:<policy_name>:<policy_type>:<group_name>` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The attachment ID. Composed of policy name, policy type and group name with format ` + "`" + `group:<policy_name>:<policy_type>:<group_name>` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ram_login_profile",
			Category:         "RAM Resources",
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
					Description: `(Required) Password of the RAM user.`,
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
			Category:         "RAM Resources",
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
			Category:         "RAM Resources",
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
			Category:         "RAM Resources",
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
			Category:         "RAM Resources",
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
					Description: `The attachment ID. Composed of policy name, policy type and role name with format ` + "`" + `role:<policy_name>:<policy_type>:<role_name>` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The attachment ID. Composed of policy name, policy type and role name with format ` + "`" + `role:<policy_name>:<policy_type>:<role_name>` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ram_user",
			Category:         "RAM Resources",
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
					Description: `(Optional) Name of the RAM user which for display. This name can have a string of 1 to 12 characters or Chinese characters, must contain only alphanumeric characters or Chinese characters or hyphens, such as "-",".", and must not end with a hyphen.`,
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
					Description: `The original id is user name, but it is user id in 1.37.0+.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The user name.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The user display name.`,
				},
				resource.Attribute{
					Name:        "mobile",
					Description: `The user phone number.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `The user email.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `The user comments. ## Import RAM user can be imported using the id or name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ram_user.example user ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The original id is user name, but it is user id in 1.37.0+.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The user name.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The user display name.`,
				},
				resource.Attribute{
					Name:        "mobile",
					Description: `The user phone number.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `The user email.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `The user comments. ## Import RAM user can be imported using the id or name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ram_user.example user ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ram_user_policy_attachment",
			Category:         "RAM Resources",
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
					Description: `The attachment ID. Composed of policy name, policy type and user name with format ` + "`" + `user:<policy_name>:<policy_type>:<user_name>` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The attachment ID. Composed of policy name, policy type and user name with format ` + "`" + `user:<policy_name>:<policy_type>:<user_name>` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_route_entry",
			Category:         "VPC Resources",
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
					Description: `(ForceNew) The route entry's next hop. ECS instance ID or VPC router interface ID. ## Attributes Reference The following attributes are exported:`,
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
			Category:         "VPC Resources",
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
					Description: `(Optional) The description of the route table instance. ## Attributes Reference The following attributes are exported:`,
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
			Category:         "VPC Resources",
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
			Category:         "VPC Resources",
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
			Category:         "VPC Resources",
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
			Type:             "alicloud_security_group",
			Category:         "ECS Resources",
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
					Name:        "inner_access",
					Description: `(Optional) Whether to allow both machines to access each other on all ports in the same security group.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. Combining security group rules, the policy can define multiple application scenario. Default to true. It is valid from verison ` + "`" + `1.7.2` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the security group`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The VPC ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the security group`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the security group`,
				},
				resource.Attribute{
					Name:        "inner_access",
					Description: `Whether to allow inner network access.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The instance tags, use jsonencode(item) to display the value. ## Import Security Group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_security_group.example sg-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the security group`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The VPC ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the security group`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the security group`,
				},
				resource.Attribute{
					Name:        "inner_access",
					Description: `Whether to allow inner network access.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The instance tags, use jsonencode(item) to display the value. ## Import Security Group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_security_group.example sg-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_security_group_rule",
			Category:         "ECS Resources",
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
			Type:             "alicloud_slb",
			Category:         "SLB Resources",
			ShortDescription: `Provides an Application Load Banlancer resource.`,
			Description:      ``,
			Keywords: []string{
				"slb",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the SLB. This name must be unique within your AliCloud account, can have a maximum of 80 characters, must contain only alphanumeric characters or hyphens, such as "-","/",".","_", and must not begin or end with a hyphen. If not specified, Terraform will autogenerate a name beginning with ` + "`" + `tf-lb` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "internet",
					Description: `(Optional, ForceNew) If true, the SLB addressType will be internet, false will be intranet, Default is false. If load balancer launched in VPC, this value must be "false".`,
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
					Name:        "listener",
					Description: `(Deprecated) The field has been deprecated from terraform-alicloud-provider [version 1.3.0](https://github.com/alibaba/terraform-provider/releases/tag/V1.3.0), and use resource ` + "`" + `alicloud_slb_listener` + "`" + ` to replace.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Required for a VPC SLB, Forces New Resource) The VSwitch ID to launch in.`,
				},
				resource.Attribute{
					Name:        "specification",
					Description: `(Optional) The specification of the Server Load Balancer instance. Default to empty string indicating it is "Shared-Performance" instance. Launching "[Performance-guaranteed](https://www.alibabacloud.com/help/doc-detail/27657.htm)" instance, it is must be specified and it valid values are: "slb.s1.small", "slb.s2.small", "slb.s2.medium", "slb.s3.small", "slb.s3.medium" and "slb.s3.large".`,
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
					Description: `(Optional, Available in v1.51.0+) Whether enable the deletion protection or not. on: Enable deletion protection. off: Disable deletion protection. Default to off. Only postpaid instance support this function. ->`,
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
			Category:         "SLB Resources",
			ShortDescription: `Provides a Load Banlancer Access Control List resource.`,
			Description:      ``,
			Keywords: []string{
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
					Description: `(Optional) A list of entry (IP addresses or CIDR blocks) to be added. At most 50 etnry can be supported in one resource. It contains two sub-fields as ` + "`" + `Entry Block` + "`" + ` follows. ## Entry Block The entry mapping supports the following:`,
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
			Category:         "SLB Resources",
			ShortDescription: `Provides an Application Load Banlancer Attachment resource.`,
			Description:      ``,
			Keywords: []string{
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
					Name:        "slb_id",
					Description: `(Deprecated) It has been deprecated from provider version 1.6.0. New field 'load_balancer_id' replaces it.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `(Deprecated) It has been deprecated from provider version 1.6.0. New field 'instance_ids' replaces it. ## Attributes Reference The following attributes are exported:`,
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
					Description: `(Optional) Weight of the instances.`,
				},
				resource.Attribute{
					Name:        "backend_servers",
					Description: `The backend servers of the load balancer. ## Import Load balancer attachment can be imported using the id or load balancer id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_slb_attachment.example lb-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Optional) Weight of the instances.`,
				},
				resource.Attribute{
					Name:        "backend_servers",
					Description: `The backend servers of the load balancer. ## Import Load balancer attachment can be imported using the id or load balancer id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_slb_attachment.example lb-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_slb_ca_certificate",
			Category:         "SLB Resources",
			ShortDescription: `Provides a Load Banlancer CA Certificate resource.`,
			Description:      ``,
			Keywords: []string{
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
			Type:             "alicloud_slb_listener",
			Category:         "SLB Resources",
			ShortDescription: `Provides an Application Load Banlancer resource.`,
			Description:      ``,
			Keywords: []string{
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
					Name:        "ssl_certificate_id",
					Description: `(Optional) Security certificate ID. It is required when ` + "`" + `protocol` + "`" + ` is ` + "`" + `https` + "`" + `.`,
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
					Description: `(Optional, ForceNew, Available in 1.40.0+) The port that http redirect to https. ->`,
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
					Description: `(Optional) Whether to use the XForwardedFor_proto header to obtain the protocol used by the listener. Default to false. ## Listener fields and protocol mapping load balance support 4 protocal to listen on, they are ` + "`" + `http` + "`" + `,` + "`" + `https` + "`" + `,` + "`" + `tcp` + "`" + `,` + "`" + `udp` + "`" + `, the every listener support which portocal following: listener parameter | support protocol | value range | ------------- | ------------- | ------------- | backend_port | http & https & tcp & udp | 1-65535 | frontend_port | http & https & tcp & udp | 1-65535 | protocol | http & https & tcp & udp | bandwidth | http & https & tcp & udp | -1 / 1-1000 | scheduler | http & https & tcp & udp | wrr rr or wlc | sticky_session | http & https | on or off | sticky_session_type | http & https | insert or server | cookie_timeout | http & https | 1-86400 | cookie | http & https | | persistence_timeout | tcp & udp | 0-3600 | health_check | http & https | on or off | health_check_type | tcp | tcp or http | health_check_domain | http & https & tcp | health_check_uri | http & https & tcp | | health_check_connect_port | http & https & tcp & udp | 1-65535 or -520 | healthy_threshold | http & https & tcp & udp | 1-10 | unhealthy_threshold | http & https & tcp & udp | 1-10 | health_check_timeout | http & https & tcp & udp | 1-300 | health_check_interval | http & https & tcp & udp | 1-50 | health_check_http_code | http & https & tcp | http_2xx,http_3xx,http_4xx,http_5xx | ssl_certificate_id | https | | gzip | http & https | true or false | x_forwarded_for | http & https | | acl_status | http & https & tcp & udp | on or off | acl_type | http & https & tcp & udp | white or black | acl_id | http & https & tcp & udp | the id of resource alicloud_slb_acl| established_timeout | tcp | 10-900| idle_timeout |http & https | 1-60 | request_timeout |http & https | 1-180 | enable_http2 |https | on or off | tls_cipher_policy |https | tls_cipher_policy_1_0, tls_cipher_policy_1_1, tls_cipher_policy_1_2, tls_cipher_policy_1_2_strict | server_group_id | http & https & tcp & udp | the id of resource alicloud_slb_server_group | The listener mapping supports the following: ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the load balancer listener. It is consist of ` + "`" + `load_balancer_id` + "`" + ` and ` + "`" + `frontend_port` + "`" + `: ` + "`" + `<load_balancer_id>:<frontend_port>` + "`" + `.`,
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
					Name:        "ssl_certificate_id",
					Description: `(Optional) Security certificate ID. ## Import Load balancer listener can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_slb_listener.example "lb-abc123456:22" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the load balancer listener. It is consist of ` + "`" + `load_balancer_id` + "`" + ` and ` + "`" + `frontend_port` + "`" + `: ` + "`" + `<load_balancer_id>:<frontend_port>` + "`" + `.`,
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
					Name:        "ssl_certificate_id",
					Description: `(Optional) Security certificate ID. ## Import Load balancer listener can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_slb_listener.example "lb-abc123456:22" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_slb_rule",
			Category:         "SLB Resources",
			ShortDescription: `Provides a Load Banlancer Forwarding Rule Resource and add it to one Listener.`,
			Description:      ``,
			Keywords: []string{
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
					Description: `(Optional, ForceNew) Name of the forwarding rule. Our plugin provides a default name: "tf-slb-rule".`,
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
					Description: `(Optional, Available in v1.51.0+) Indicates whether a forwarding rule inherits the settings of a health check , session persistence, and scheduling algorithm from a listener. Default to on. ## Attributes Reference The following attributes are exported:`,
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
			Category:         "SLB Resources",
			ShortDescription: `Provides a Load Banlancer Server Certificate resource.`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"server",
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
			Category:         "SLB Resources",
			ShortDescription: `Provides a Load Banlancer Virtual Backend Server Group resource.`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"server",
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
					Description: `A list of ECS instances to be added. At most 20 ECS instances can be supported in one resource. It contains three sub-fields as ` + "`" + `Block server` + "`" + ` follows. ## Block servers The servers mapping supports the following:`,
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
			Category:         "ECS Resources",
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
			Category:         "ECS Resources",
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
			Category:         "VPC Resources",
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
					Description: `(Required, ForceNew) The vswitch ID.`,
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
			Type:             "alicloud_ssl_vpn_cert_client",
			Category:         "VPN Gateway Resources",
			ShortDescription: `Provides a Alicloud SSL VPN Client Cert resource.`,
			Description:      ``,
			Keywords: []string{
				"vpn",
				"gateway",
				"ssl",
				"cert",
				"client",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the client certificate.`,
				},
				resource.Attribute{
					Name:        "ssl_vpn_server_id",
					Description: `(Required, ForceNew) The ID of the SSL-VPN server. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the SSL-VPN client certificate.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the client certificate. ## Import SSL-VPN client certificates can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ssl_vpn_client_cert.example vsc-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the SSL-VPN client certificate.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the client certificate. ## Import SSL-VPN client certificates can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import alicloud_ssl_vpn_client_cert.example vsc-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ssl_vpn_server",
			Category:         "VPN Gateway Resources",
			ShortDescription: `Provides a Alicloud SSL VPN server resource.`,
			Description:      ``,
			Keywords: []string{
				"vpn",
				"gateway",
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
					Description: `(Required) The CIDR block to be accessed by the client through the SSL-VPN connection.`,
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
			Category:         "VPC Resources",
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
					Description: `(Optional, Available in 1.40.0+) The Id of resource group which the VPC belongs. ## Attributes Reference The following attributes are exported:`,
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
			Category:         "VPN Gateway Resources",
			ShortDescription: `Provides a Alicloud VPN connection resource.`,
			Description:      ``,
			Keywords: []string{
				"vpn",
				"gateway",
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
					Description: `(Optional) The Diffie-Hellman key exchange algorithm used by phase-two negotiation. Valid value: group1 | group2 | group5 | group14 | group24. Default value: group2`,
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
			Category:         "VPN Gateway Resources",
			ShortDescription: `Provides a Alicloud VPN customer gateway resource.`,
			Description:      ``,
			Keywords: []string{
				"vpn",
				"gateway",
				"customer",
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
			Category:         "VPN Gateway Resources",
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
					Description: `(ForceNew) The charge type for instance. Valid value: PostPaid, PrePaid. Default to PostPaid.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) The filed is only required while the InstanceChargeType is prepaid.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Required) The value should be 10, 100, 200, 500, 1000 if the user is postpaid, otherwise it can be 5, 10, 20, 50, 100, 200, 500, 1000. It can't be changed by terraform.`,
				},
				resource.Attribute{
					Name:        "enable_ipsec",
					Description: `(Optional) Enable or Disable IPSec VPN. At least one type of VPN should be enabled.`,
				},
				resource.Attribute{
					Name:        "ssl_connections",
					Description: `(Optional) The max connections of SSL VPN. Default to 5. This field is ignored when enable_ssl is false.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the VPN instance. ## Attributes Reference The following attributes are exported:`,
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
			Type:             "alicloud_vswitch",
			Category:         "VPC Resources",
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
					Description: `(Optional) The switch description. Defaults to null. ## Attributes Reference The following attributes are exported:`,
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
	}

	resourcesMap = map[string]int{

		"alicloud_actiontrail":                         0,
		"alicloud_api_gateway_api":                     1,
		"alicloud_api_gateway_app":                     2,
		"alicloud_api_gateway_app_attachment":          3,
		"alicloud_api_gateway_group":                   4,
		"alicloud_api_gateway_vpc_access":              5,
		"alicloud_cas_certificate":                     6,
		"alicloud_cdn_domain":                          7,
		"alicloud_cdn_domain_config":                   8,
		"alicloud_cdn_domain_new":                      9,
		"alicloud_cen_bandwidth_limit":                 10,
		"alicloud_cen_bandwidth_package":               11,
		"alicloud_cen_bandwidth_package_attachment":    12,
		"alicloud_cen_instance":                        13,
		"alicloud_cen_instance_attachment":             14,
		"alicloud_cen_instance_grant":                  15,
		"alicloud_cen_route_entry":                     16,
		"alicloud_cms_alarm":                           17,
		"alicloud_common_bandwidth_package":            18,
		"alicloud_common_bandwidth_package_attachment": 19,
		"alicloud_container_cluster":                   20,
		"alicloud_cr_namespace":                        21,
		"alicloud_cr_repo":                             22,
		"alicloud_cs_application":                      23,
		"alicloud_cs_kubernetes":                       24,
		"alicloud_cs_managed_kubernetes":               25,
		"alicloud_cs_swarm":                            26,
		"alicloud_datahub_project":                     27,
		"alicloud_datahub_subscription":                28,
		"alicloud_datahub_topic":                       29,
		"alicloud_db_account":                          30,
		"alicloud_db_account_privilege":                31,
		"alicloud_db_backup_policy":                    32,
		"alicloud_db_connection":                       33,
		"alicloud_db_database":                         34,
		"alicloud_db_instance":                         35,
		"alicloud_db_read_write_splitting_connection":  36,
		"alicloud_db_readonly_instance":                37,
		"alicloud_ddoscoo_instance":                    38,
		"alicloud_disk":                                39,
		"alicloud_disk_attachment":                     40,
		"alicloud_dns":                                 41,
		"alicloud_dns_group":                           42,
		"alicloud_dns_record":                          43,
		"alicloud_drds_instance":                       44,
		"alicloud_eip":                                 45,
		"alicloud_eip_association":                     46,
		"alicloud_elasticsearch_instance":              47,
		"alicloud_ess_alarm":                           48,
		"alicloud_ess_attachment":                      49,
		"alicloud_ess_scaling_configuration":           50,
		"alicloud_ess_scaling_group":                   51,
		"alicloud_ess_lifecycle_hook":                  52,
		"alicloud_ess_scaling_rule":                    53,
		"alicloud_ess_schedule":                        54,
		"alicloud_ess_scheduled_task":                  55,
		"alicloud_fc_function":                         56,
		"alicloud_fc_service":                          57,
		"alicloud_fc_trigger":                          58,
		"alicloud_forward_entry":                       59,
		"alicloud_gpdb_connection":                     60,
		"alicloud_gpdb_instance":                       61,
		"alicloud_instance":                            62,
		"alicloud_key_pair":                            63,
		"alicloud_key_pair_attachment":                 64,
		"alicloud_kms_key":                             65,
		"alicloud_kvstore_backup_policy":               66,
		"alicloud_kvstore_instance":                    67,
		"alicloud_launch_template":                     68,
		"alicloud_log_machine_group":                   69,
		"alicloud_log_project":                         70,
		"alicloud_log_store":                           71,
		"alicloud_log_store_index":                     72,
		"alicloud_logtail_attachment":                  73,
		"alicloud_logtail_config":                      74,
		"alicloud_mns_queue":                           75,
		"alicloud_mns_topic":                           76,
		"alicloud_mns_topic_subscription":              77,
		"alicloud_mongodb_instance":                    78,
		"mongodb_sharding_instance":                    79,
		"alicloud_nas_access_group":                    80,
		"alicloud_nas_access_rule":                     81,
		"alicloud_nas_file_system":                     82,
		"alicloud_nas_mount_target":                    83,
		"alicloud_nat_gateway":                         84,
		"alicloud_network_acl":                         85,
		"alicloud_network_acl_attachment":              86,
		"alicloud_network_acl_entries":                 87,
		"alicloud_network_interface":                   88,
		"alicloud_network_interface_attachment":        89,
		"alicloud_ons_instance":                        90,
		"alicloud_oss_bucket":                          91,
		"alicloud_oss_bucket_object":                   92,
		"alicloud_ots_instance":                        93,
		"alicloud_ots_instance_attachment":             94,
		"alicloud_ots_table":                           95,
		"alicloud_pvtz_zone":                           96,
		"alicloud_pvtz_zone_attachment":                97,
		"alicloud_pvtz_zone_record":                    98,
		"alicloud_ram_access_key":                      99,
		"alicloud_ram_account_alias":                   100,
		"alicloud_ram_account_password_policy":         101,
		"alicloud_ram_alias":                           102,
		"alicloud_ram_group":                           103,
		"alicloud_ram_group_membership":                104,
		"alicloud_ram_group_policy_attachment":         105,
		"alicloud_ram_login_profile":                   106,
		"alicloud_ram_policy":                          107,
		"alicloud_ram_role":                            108,
		"alicloud_ram_role_attachment":                 109,
		"alicloud_ram_role_policy_attachment":          110,
		"alicloud_ram_user":                            111,
		"alicloud_ram_user_policy_attachment":          112,
		"alicloud_route_entry":                         113,
		"alicloud_route_table":                         114,
		"alicloud_route_table_attachment":              115,
		"alicloud_router_interface":                    116,
		"alicloud_router_interface_connection":         117,
		"alicloud_security_group":                      118,
		"alicloud_security_group_rule":                 119,
		"alicloud_slb":                                 120,
		"alicloud_slb_acl":                             121,
		"alicloud_slb_attachment":                      122,
		"alicloud_slb_ca_certificate":                  123,
		"alicloud_slb_listener":                        124,
		"alicloud_slb_rule":                            125,
		"alicloud_slb_server_certificate":              126,
		"alicloud_slb_server_group":                    127,
		"alicloud_snapshot":                            128,
		"alicloud_snapshot_policy":                     129,
		"alicloud_snat_entry":                          130,
		"alicloud_ssl_vpn_cert_client":                 131,
		"alicloud_ssl_vpn_server":                      132,
		"alicloud_vpc":                                 133,
		"alicloud_vpn_connection":                      134,
		"alicloud_vpn_customer_gateway":                135,
		"alicloud_vpn_gateway":                         136,
		"alicloud_vswitch":                             137,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
