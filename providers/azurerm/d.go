package azurerm

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "azurerm_advisor_recommendations",
			Category:         "Advisor",
			ShortDescription: `Gets information about an existing Advisor Recommendations.`,
			Description: `

Use this data source to access information about an existing Advisor Recommendations.

`,
			Keywords: []string{
				"advisor",
				"recommendations",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Advisor Recommendations.`,
				},
				resource.Attribute{
					Name:        "recommendations",
					Description: `One or more ` + "`" + `recommendations` + "`" + ` blocks as defined below. --- A ` + "`" + `recommendations` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `The category of the recommendation.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the issue or the opportunity identified by the recommendation.`,
				},
				resource.Attribute{
					Name:        "impact",
					Description: `The business impact of the recommendation.`,
				},
				resource.Attribute{
					Name:        "recommendation_name",
					Description: `The name of the Advisor Recommendation.`,
				},
				resource.Attribute{
					Name:        "recommendation_type_id",
					Description: `The recommendation type id of the Advisor Recommendation.`,
				},
				resource.Attribute{
					Name:        "resource_name",
					Description: `The name of the identified resource of the Advisor Recommendation.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `The type of the identified resource of the Advisor Recommendation.`,
				},
				resource.Attribute{
					Name:        "suppression_names",
					Description: `A list of Advisor Suppression names of the Advisor Recommendation.`,
				},
				resource.Attribute{
					Name:        "updated_time",
					Description: `The most recent time that Advisor checked the validity of the recommendation.. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 10 minutes) Used when retrieving the Advisor Recommendations.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_api_management",
			Category:         "API Management",
			ShortDescription: `Gets information about an existing API Management Service.`,
			Description: `

Use this data source to access information about an existing API Management Service.

`,
			Keywords: []string{
				"api",
				"management",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the API Management service.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The Name of the Resource Group in which the API Management Service exists. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the API Management Service.`,
				},
				resource.Attribute{
					Name:        "additional_location",
					Description: `Zero or more ` + "`" + `additional_location` + "`" + ` blocks as defined below`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the API Management Service exists.`,
				},
				resource.Attribute{
					Name:        "gateway_url",
					Description: `The URL for the API Management Service's Gateway.`,
				},
				resource.Attribute{
					Name:        "gateway_regional_url",
					Description: `The URL for the Gateway in the Default Region.`,
				},
				resource.Attribute{
					Name:        "identity",
					Description: `(Optional) An ` + "`" + `identity` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "hostname_configuration",
					Description: `A ` + "`" + `hostname_configuration` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "management_api_url",
					Description: `The URL for the Management API.`,
				},
				resource.Attribute{
					Name:        "notification_sender_email",
					Description: `The email address from which the notification will be sent.`,
				},
				resource.Attribute{
					Name:        "portal_url",
					Description: `The URL of the Publisher Portal.`,
				},
				resource.Attribute{
					Name:        "developer_portal_url",
					Description: `The URL for the Developer Portal associated with this API Management service.`,
				},
				resource.Attribute{
					Name:        "public_ip_addresses",
					Description: `The Public IP addresses of the API Management Service.`,
				},
				resource.Attribute{
					Name:        "private_ip_addresses",
					Description: `The Private IP addresses of the API Management Service.`,
				},
				resource.Attribute{
					Name:        "publisher_name",
					Description: `The name of the Publisher/Company of the API Management Service.`,
				},
				resource.Attribute{
					Name:        "publisher_email",
					Description: `The email of Publisher/Company of the API Management Service.`,
				},
				resource.Attribute{
					Name:        "scm_url",
					Description: `The SCM (Source Code Management) endpoint.`,
				},
				resource.Attribute{
					Name:        "sku",
					Description: `A ` + "`" + `sku` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource. --- A ` + "`" + `additional_location` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location name of the additional region among Azure Data center regions.`,
				},
				resource.Attribute{
					Name:        "gateway_regional_url",
					Description: `Gateway URL of the API Management service in the Region.`,
				},
				resource.Attribute{
					Name:        "public_ip_addresses",
					Description: `Public Static Load Balanced IP addresses of the API Management service in the additional location. Available only for Basic, Standard and Premium SKU.`,
				},
				resource.Attribute{
					Name:        "private_ip_addresses",
					Description: `Private IP addresses of the API Management service in the additional location, for instances using virtual network mode. --- A ` + "`" + `identity` + "`" + ` block exports the following: ~>`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Specifies the type of Managed Service Identity that is configured on this API Management Service.`,
				},
				resource.Attribute{
					Name:        "principal_id",
					Description: `Specifies the Principal ID of the System Assigned Managed Service Identity that is configured on this API Management Service.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `Specifies the Tenant ID of the System Assigned Managed Service Identity that is configured on this API Management Service.`,
				},
				resource.Attribute{
					Name:        "identity_ids",
					Description: `A list of IDs for User Assigned Managed Identity resources to be assigned. --- A ` + "`" + `hostname_configuration` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "management",
					Description: `One or more ` + "`" + `management` + "`" + ` blocks as documented below.`,
				},
				resource.Attribute{
					Name:        "portal",
					Description: `One or more ` + "`" + `portal` + "`" + ` blocks as documented below.`,
				},
				resource.Attribute{
					Name:        "developer_portal",
					Description: `One or more ` + "`" + `developer_portal` + "`" + ` blocks as documented below.`,
				},
				resource.Attribute{
					Name:        "proxy",
					Description: `One or more ` + "`" + `proxy` + "`" + ` blocks as documented below.`,
				},
				resource.Attribute{
					Name:        "scm",
					Description: `One or more ` + "`" + `scm` + "`" + ` blocks as documented below. --- A ` + "`" + `management` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "host_name",
					Description: `The Hostname used for the Management API.`,
				},
				resource.Attribute{
					Name:        "key_vault_id",
					Description: `The ID of the Key Vault Secret which contains the SSL Certificate.`,
				},
				resource.Attribute{
					Name:        "negotiate_client_certificate",
					Description: `Is Client Certificate Negotiation enabled? --- A ` + "`" + `portal` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "host_name",
					Description: `The Hostname used for the Portal.`,
				},
				resource.Attribute{
					Name:        "key_vault_id",
					Description: `The ID of the Key Vault Secret which contains the SSL Certificate.`,
				},
				resource.Attribute{
					Name:        "negotiate_client_certificate",
					Description: `Is Client Certificate Negotiation enabled? --- A ` + "`" + `developer_portal` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "host_name",
					Description: `The Hostname used for the Portal.`,
				},
				resource.Attribute{
					Name:        "key_vault_id",
					Description: `The ID of the Key Vault Secret which contains the SSL Certificate.`,
				},
				resource.Attribute{
					Name:        "negotiate_client_certificate",
					Description: `Is Client Certificate Negotiation enabled? --- A ` + "`" + `proxy` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "default_ssl_binding",
					Description: `Is this the default SSL Binding?`,
				},
				resource.Attribute{
					Name:        "host_name",
					Description: `The Hostname used for the Proxy.`,
				},
				resource.Attribute{
					Name:        "key_vault_id",
					Description: `The ID of the Key Vault Secret which contains the SSL Certificate.`,
				},
				resource.Attribute{
					Name:        "negotiate_client_certificate",
					Description: `Is Client Certificate Negotiation enabled? --- A ` + "`" + `scm` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "host_name",
					Description: `The Hostname used for the SCM URL.`,
				},
				resource.Attribute{
					Name:        "key_vault_id",
					Description: `The ID of the Key Vault Secret which contains the SSL Certificate.`,
				},
				resource.Attribute{
					Name:        "negotiate_client_certificate",
					Description: `Is Client Certificate Negotiation enabled? --- A ` + "`" + `sku` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the plan's pricing tier.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `Specifies the number of units associated with this API Management service. --- ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the API Management Service.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the API Management Service.`,
				},
				resource.Attribute{
					Name:        "additional_location",
					Description: `Zero or more ` + "`" + `additional_location` + "`" + ` blocks as defined below`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the API Management Service exists.`,
				},
				resource.Attribute{
					Name:        "gateway_url",
					Description: `The URL for the API Management Service's Gateway.`,
				},
				resource.Attribute{
					Name:        "gateway_regional_url",
					Description: `The URL for the Gateway in the Default Region.`,
				},
				resource.Attribute{
					Name:        "identity",
					Description: `(Optional) An ` + "`" + `identity` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "hostname_configuration",
					Description: `A ` + "`" + `hostname_configuration` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "management_api_url",
					Description: `The URL for the Management API.`,
				},
				resource.Attribute{
					Name:        "notification_sender_email",
					Description: `The email address from which the notification will be sent.`,
				},
				resource.Attribute{
					Name:        "portal_url",
					Description: `The URL of the Publisher Portal.`,
				},
				resource.Attribute{
					Name:        "developer_portal_url",
					Description: `The URL for the Developer Portal associated with this API Management service.`,
				},
				resource.Attribute{
					Name:        "public_ip_addresses",
					Description: `The Public IP addresses of the API Management Service.`,
				},
				resource.Attribute{
					Name:        "private_ip_addresses",
					Description: `The Private IP addresses of the API Management Service.`,
				},
				resource.Attribute{
					Name:        "publisher_name",
					Description: `The name of the Publisher/Company of the API Management Service.`,
				},
				resource.Attribute{
					Name:        "publisher_email",
					Description: `The email of Publisher/Company of the API Management Service.`,
				},
				resource.Attribute{
					Name:        "scm_url",
					Description: `The SCM (Source Code Management) endpoint.`,
				},
				resource.Attribute{
					Name:        "sku",
					Description: `A ` + "`" + `sku` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource. --- A ` + "`" + `additional_location` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location name of the additional region among Azure Data center regions.`,
				},
				resource.Attribute{
					Name:        "gateway_regional_url",
					Description: `Gateway URL of the API Management service in the Region.`,
				},
				resource.Attribute{
					Name:        "public_ip_addresses",
					Description: `Public Static Load Balanced IP addresses of the API Management service in the additional location. Available only for Basic, Standard and Premium SKU.`,
				},
				resource.Attribute{
					Name:        "private_ip_addresses",
					Description: `Private IP addresses of the API Management service in the additional location, for instances using virtual network mode. --- A ` + "`" + `identity` + "`" + ` block exports the following: ~>`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Specifies the type of Managed Service Identity that is configured on this API Management Service.`,
				},
				resource.Attribute{
					Name:        "principal_id",
					Description: `Specifies the Principal ID of the System Assigned Managed Service Identity that is configured on this API Management Service.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `Specifies the Tenant ID of the System Assigned Managed Service Identity that is configured on this API Management Service.`,
				},
				resource.Attribute{
					Name:        "identity_ids",
					Description: `A list of IDs for User Assigned Managed Identity resources to be assigned. --- A ` + "`" + `hostname_configuration` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "management",
					Description: `One or more ` + "`" + `management` + "`" + ` blocks as documented below.`,
				},
				resource.Attribute{
					Name:        "portal",
					Description: `One or more ` + "`" + `portal` + "`" + ` blocks as documented below.`,
				},
				resource.Attribute{
					Name:        "developer_portal",
					Description: `One or more ` + "`" + `developer_portal` + "`" + ` blocks as documented below.`,
				},
				resource.Attribute{
					Name:        "proxy",
					Description: `One or more ` + "`" + `proxy` + "`" + ` blocks as documented below.`,
				},
				resource.Attribute{
					Name:        "scm",
					Description: `One or more ` + "`" + `scm` + "`" + ` blocks as documented below. --- A ` + "`" + `management` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "host_name",
					Description: `The Hostname used for the Management API.`,
				},
				resource.Attribute{
					Name:        "key_vault_id",
					Description: `The ID of the Key Vault Secret which contains the SSL Certificate.`,
				},
				resource.Attribute{
					Name:        "negotiate_client_certificate",
					Description: `Is Client Certificate Negotiation enabled? --- A ` + "`" + `portal` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "host_name",
					Description: `The Hostname used for the Portal.`,
				},
				resource.Attribute{
					Name:        "key_vault_id",
					Description: `The ID of the Key Vault Secret which contains the SSL Certificate.`,
				},
				resource.Attribute{
					Name:        "negotiate_client_certificate",
					Description: `Is Client Certificate Negotiation enabled? --- A ` + "`" + `developer_portal` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "host_name",
					Description: `The Hostname used for the Portal.`,
				},
				resource.Attribute{
					Name:        "key_vault_id",
					Description: `The ID of the Key Vault Secret which contains the SSL Certificate.`,
				},
				resource.Attribute{
					Name:        "negotiate_client_certificate",
					Description: `Is Client Certificate Negotiation enabled? --- A ` + "`" + `proxy` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "default_ssl_binding",
					Description: `Is this the default SSL Binding?`,
				},
				resource.Attribute{
					Name:        "host_name",
					Description: `The Hostname used for the Proxy.`,
				},
				resource.Attribute{
					Name:        "key_vault_id",
					Description: `The ID of the Key Vault Secret which contains the SSL Certificate.`,
				},
				resource.Attribute{
					Name:        "negotiate_client_certificate",
					Description: `Is Client Certificate Negotiation enabled? --- A ` + "`" + `scm` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "host_name",
					Description: `The Hostname used for the SCM URL.`,
				},
				resource.Attribute{
					Name:        "key_vault_id",
					Description: `The ID of the Key Vault Secret which contains the SSL Certificate.`,
				},
				resource.Attribute{
					Name:        "negotiate_client_certificate",
					Description: `Is Client Certificate Negotiation enabled? --- A ` + "`" + `sku` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the plan's pricing tier.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `Specifies the number of units associated with this API Management service. --- ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the API Management Service.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_api_management_api",
			Category:         "API Management",
			ShortDescription: `Gets information about an existing API Management API.`,
			Description: `

Use this data source to access information about an existing API Management API.

`,
			Keywords: []string{
				"api",
				"management",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the API Management API.`,
				},
				resource.Attribute{
					Name:        "api_management_name",
					Description: `The name of the API Management Service in which the API Management API exists.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The Name of the Resource Group in which the API Management Service exists.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `The Revision of the API Management API. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the API Management API.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the API Management API, which may include HTML formatting tags.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name of the API.`,
				},
				resource.Attribute{
					Name:        "is_current",
					Description: `Is this the current API Revision?`,
				},
				resource.Attribute{
					Name:        "is_online",
					Description: `Is this API Revision online/accessible via the Gateway?`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The Path for this API Management API.`,
				},
				resource.Attribute{
					Name:        "protocols",
					Description: `A list of protocols the operations in this API can be invoked.`,
				},
				resource.Attribute{
					Name:        "service_url",
					Description: `Absolute URL of the backend service implementing this API.`,
				},
				resource.Attribute{
					Name:        "soap_pass_through",
					Description: `Should this API expose a SOAP frontend, rather than a HTTP frontend?`,
				},
				resource.Attribute{
					Name:        "subscription_key_parameter_names",
					Description: `A ` + "`" + `subscription_key_parameter_names` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "subscription_required",
					Description: `Should this API require a subscription key?`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The Version number of this API, if this API is versioned.`,
				},
				resource.Attribute{
					Name:        "version_set_id",
					Description: `The ID of the Version Set which this API is associated with. --- A ` + "`" + `subscription_key_parameter_names` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "header",
					Description: `The name of the HTTP Header which should be used for the Subscription Key.`,
				},
				resource.Attribute{
					Name:        "query",
					Description: `The name of the QueryString parameter which should be used for the Subscription Key. --- A ` + "`" + `wsdl_selector` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `The name of service to import from WSDL.`,
				},
				resource.Attribute{
					Name:        "endpoint_name",
					Description: `The name of endpoint (port) to import from WSDL. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the API Management API.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the API Management API.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the API Management API, which may include HTML formatting tags.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name of the API.`,
				},
				resource.Attribute{
					Name:        "is_current",
					Description: `Is this the current API Revision?`,
				},
				resource.Attribute{
					Name:        "is_online",
					Description: `Is this API Revision online/accessible via the Gateway?`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The Path for this API Management API.`,
				},
				resource.Attribute{
					Name:        "protocols",
					Description: `A list of protocols the operations in this API can be invoked.`,
				},
				resource.Attribute{
					Name:        "service_url",
					Description: `Absolute URL of the backend service implementing this API.`,
				},
				resource.Attribute{
					Name:        "soap_pass_through",
					Description: `Should this API expose a SOAP frontend, rather than a HTTP frontend?`,
				},
				resource.Attribute{
					Name:        "subscription_key_parameter_names",
					Description: `A ` + "`" + `subscription_key_parameter_names` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "subscription_required",
					Description: `Should this API require a subscription key?`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The Version number of this API, if this API is versioned.`,
				},
				resource.Attribute{
					Name:        "version_set_id",
					Description: `The ID of the Version Set which this API is associated with. --- A ` + "`" + `subscription_key_parameter_names` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "header",
					Description: `The name of the HTTP Header which should be used for the Subscription Key.`,
				},
				resource.Attribute{
					Name:        "query",
					Description: `The name of the QueryString parameter which should be used for the Subscription Key. --- A ` + "`" + `wsdl_selector` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `The name of service to import from WSDL.`,
				},
				resource.Attribute{
					Name:        "endpoint_name",
					Description: `The name of endpoint (port) to import from WSDL. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the API Management API.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_api_management_api_version_set",
			Category:         "API Management",
			ShortDescription: `Gets information about an existing API Version Set within an existing API Management Service.`,
			Description: `

Uses this data source to access information about an API Version Set within an API Management Service.

`,
			Keywords: []string{
				"api",
				"management",
				"version",
				"set",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the API Version Set.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of the Resource Group in which the parent API Management Service exists.`,
				},
				resource.Attribute{
					Name:        "api_management_name",
					Description: `The name of the [API Management Service](api_management.html) where the API Version Set exists. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the API Version Set.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of API Version Set.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name of this API Version Set.`,
				},
				resource.Attribute{
					Name:        "versioning_schema",
					Description: `The value that determines where the API Version identifer will be located in a HTTP request.`,
				},
				resource.Attribute{
					Name:        "version_header_name",
					Description: `The name of the Header which should be read from Inbound Requests which defines the API Version.`,
				},
				resource.Attribute{
					Name:        "version_query_name",
					Description: `The name of the Query String which should be read from Inbound Requests which defines the API Version. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the API Version Set.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the API Version Set.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of API Version Set.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name of this API Version Set.`,
				},
				resource.Attribute{
					Name:        "versioning_schema",
					Description: `The value that determines where the API Version identifer will be located in a HTTP request.`,
				},
				resource.Attribute{
					Name:        "version_header_name",
					Description: `The name of the Header which should be read from Inbound Requests which defines the API Version.`,
				},
				resource.Attribute{
					Name:        "version_query_name",
					Description: `The name of the Query String which should be read from Inbound Requests which defines the API Version. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the API Version Set.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_api_management_group",
			Category:         "API Management",
			ShortDescription: `Gets information about an existing API Management Group.`,
			Description: `

Use this data source to access information about an existing API Management Group.

`,
			Keywords: []string{
				"api",
				"management",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "api_management_name",
					Description: `The Name of the API Management Service in which this Group exists.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The Name of the API Management Group.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The Name of the Resource Group in which the API Management Service exists. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the API Management Group.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name of this API Management Group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of this API Management Group.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `The identifier of the external Group.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of this API Management Group, such as ` + "`" + `custom` + "`" + ` or ` + "`" + `external` + "`" + `. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the API Management Group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the API Management Group.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name of this API Management Group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of this API Management Group.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `The identifier of the external Group.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of this API Management Group, such as ` + "`" + `custom` + "`" + ` or ` + "`" + `external` + "`" + `. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the API Management Group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_api_management_product",
			Category:         "API Management",
			ShortDescription: `Gets information about an existing API Management Product.`,
			Description: `

Use this data source to access information about an existing API Management Product.

`,
			Keywords: []string{
				"api",
				"management",
				"product",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "api_management_name",
					Description: `The Name of the API Management Service in which this Product exists.`,
				},
				resource.Attribute{
					Name:        "product_id",
					Description: `The Identifier for the API Management Product.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The Name of the Resource Group in which the API Management Service exists. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the API Management Product.`,
				},
				resource.Attribute{
					Name:        "approval_required",
					Description: `Do subscribers need to be approved prior to being able to use the Product?`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The Display Name for this API Management Product.`,
				},
				resource.Attribute{
					Name:        "published",
					Description: `Is this Product Published?`,
				},
				resource.Attribute{
					Name:        "subscription_required",
					Description: `Is a Subscription required to access API's included in this Product?`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of this Product, which may include HTML formatting tags.`,
				},
				resource.Attribute{
					Name:        "subscriptions_limit",
					Description: `The number of subscriptions a user can have to this Product at the same time.`,
				},
				resource.Attribute{
					Name:        "terms",
					Description: `Any Terms and Conditions for this Product, which must be accepted by Developers before they can begin the Subscription process. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the API Management Product.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the API Management Product.`,
				},
				resource.Attribute{
					Name:        "approval_required",
					Description: `Do subscribers need to be approved prior to being able to use the Product?`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The Display Name for this API Management Product.`,
				},
				resource.Attribute{
					Name:        "published",
					Description: `Is this Product Published?`,
				},
				resource.Attribute{
					Name:        "subscription_required",
					Description: `Is a Subscription required to access API's included in this Product?`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of this Product, which may include HTML formatting tags.`,
				},
				resource.Attribute{
					Name:        "subscriptions_limit",
					Description: `The number of subscriptions a user can have to this Product at the same time.`,
				},
				resource.Attribute{
					Name:        "terms",
					Description: `Any Terms and Conditions for this Product, which must be accepted by Developers before they can begin the Subscription process. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the API Management Product.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_api_management_user",
			Category:         "API Management",
			ShortDescription: `Gets information about an existing API Management User.`,
			Description: `

Use this data source to access information about an existing API Management User.

`,
			Keywords: []string{
				"api",
				"management",
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "api_management_name",
					Description: `The Name of the API Management Service in which this User exists.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The Name of the Resource Group in which the API Management Service exists.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `The Identifier for the User. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the API Management User.`,
				},
				resource.Attribute{
					Name:        "first_name",
					Description: `The First Name for the User.`,
				},
				resource.Attribute{
					Name:        "last_name",
					Description: `The Last Name for the User.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `The Email Address used for this User.`,
				},
				resource.Attribute{
					Name:        "note",
					Description: `Any notes about this User.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The current state of this User, for example ` + "`" + `active` + "`" + `, ` + "`" + `blocked` + "`" + ` or ` + "`" + `pending` + "`" + `. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the API Management User.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the API Management User.`,
				},
				resource.Attribute{
					Name:        "first_name",
					Description: `The First Name for the User.`,
				},
				resource.Attribute{
					Name:        "last_name",
					Description: `The Last Name for the User.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `The Email Address used for this User.`,
				},
				resource.Attribute{
					Name:        "note",
					Description: `Any notes about this User.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The current state of this User, for example ` + "`" + `active` + "`" + `, ` + "`" + `blocked` + "`" + ` or ` + "`" + `pending` + "`" + `. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the API Management User.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_app_configuration",
			Category:         "App Configuration",
			ShortDescription: `Gets information about an existing App Configuration.`,
			Description: `

Use this data source to access information about an existing App Configuration.

`,
			Keywords: []string{
				"app",
				"configuration",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the App Configuration.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `The Endpoint used to access this App Configuration.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure Region where the App Configuration exists.`,
				},
				resource.Attribute{
					Name:        "primary_read_key",
					Description: `A ` + "`" + `primary_read_key` + "`" + ` block as defined below containing the primary read access key.`,
				},
				resource.Attribute{
					Name:        "primary_write_key",
					Description: `A ` + "`" + `primary_write_key` + "`" + ` block as defined below containing the primary write access key.`,
				},
				resource.Attribute{
					Name:        "secondary_read_key",
					Description: `A ` + "`" + `secondary_read_key` + "`" + ` block as defined below containing the secondary read access key.`,
				},
				resource.Attribute{
					Name:        "secondary_write_key",
					Description: `A ` + "`" + `secondary_write_key` + "`" + ` block as defined below containing the secondary write access key.`,
				},
				resource.Attribute{
					Name:        "sku",
					Description: `The name of the SKU used for this App Configuration.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the App Configuration. --- A ` + "`" + `primary_read_key` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "connection_string",
					Description: `The Connection String for this Access Key - comprising of the Endpoint, ID and Secret.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Access Key.`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `The Secret of the Access Key. --- A ` + "`" + `primary_write_key` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "connection_string",
					Description: `The Connection String for this Access Key - comprising of the Endpoint, ID and Secret.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Access Key.`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `The Secret of the Access Key. --- A ` + "`" + `secondary_read_key` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "connection_string",
					Description: `The Connection String for this Access Key - comprising of the Endpoint, ID and Secret.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Access Key.`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `The Secret of the Access Key. --- A ` + "`" + `secondary_write_key` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "connection_string",
					Description: `The Connection String for this Access Key - comprising of the Endpoint, ID and Secret.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Access Key.`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `The Secret of the Access Key. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the App Configuration.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_app_service",
			Category:         "App Service (Web Apps)",
			ShortDescription: `Gets information about an existing App Service.`,
			Description: `

Use this data source to access information about an existing App Service.

`,
			Icon: "Containers/10035-icon-service-App-Services.svg",
			Keywords: []string{
				"app",
				"service",
				"web",
				"apps",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the App Service.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The Name of the Resource Group where the App Service exists. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the App Service.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the App Service exists.`,
				},
				resource.Attribute{
					Name:        "app_service_plan_id",
					Description: `The ID of the App Service Plan within which the App Service exists.`,
				},
				resource.Attribute{
					Name:        "app_settings",
					Description: `A key-value pair of App Settings for the App Service.`,
				},
				resource.Attribute{
					Name:        "connection_string",
					Description: `An ` + "`" + `connection_string` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "client_affinity_enabled",
					Description: `Does the App Service send session affinity cookies, which route client requests in the same session to the same instance?`,
				},
				resource.Attribute{
					Name:        "custom_domain_verification_id",
					Description: `An identifier used by App Service to perform domain ownership verification via DNS TXT record.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Is the App Service Enabled?`,
				},
				resource.Attribute{
					Name:        "https_only",
					Description: `Can the App Service only be accessed via HTTPS?`,
				},
				resource.Attribute{
					Name:        "client_cert_enabled",
					Description: `Does the App Service require client certificates for incoming requests?`,
				},
				resource.Attribute{
					Name:        "site_config",
					Description: `A ` + "`" + `site_config` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "default_site_hostname",
					Description: `The Default Hostname associated with the App Service - such as ` + "`" + `mysite.azurewebsites.net` + "`" + ``,
				},
				resource.Attribute{
					Name:        "outbound_ip_addresses",
					Description: `A comma separated list of outbound IP addresses - such as ` + "`" + `52.23.25.3,52.143.43.12` + "`" + ``,
				},
				resource.Attribute{
					Name:        "outbound_ip_address_list",
					Description: `A list of outbound IP addresses - such as ` + "`" + `["52.23.25.3", "52.143.43.12"]` + "`" + ``,
				},
				resource.Attribute{
					Name:        "possible_outbound_ip_addresses",
					Description: `A comma separated list of outbound IP addresses - such as ` + "`" + `52.23.25.3,52.143.43.12,52.143.43.17` + "`" + ` - not all of which are necessarily in use. Superset of ` + "`" + `outbound_ip_addresses` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "possible_outbound_ip_address_list",
					Description: `A list of outbound IP addresses - such as ` + "`" + `["52.23.25.3", "52.143.43.12", "52.143.43.17"]` + "`" + ` - not all of which are necessarily in use. Superset of ` + "`" + `outbound_ip_address_list` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "source_control",
					Description: `A ` + "`" + `source_control` + "`" + ` block as defined below. --- A ` + "`" + `connection_string` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Connection String.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the Connection String.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value for the Connection String. --- A ` + "`" + `cors` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "allowed_origins",
					Description: `A list of origins which are able to make cross-origin calls.`,
				},
				resource.Attribute{
					Name:        "support_credentials",
					Description: `Are credentials supported? --- An ` + "`" + `ip_restriction` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The IP Address used for this IP Restriction.`,
				},
				resource.Attribute{
					Name:        "service_tag",
					Description: `The Service Tag used for this IP Restriction.`,
				},
				resource.Attribute{
					Name:        "subnet_mask",
					Description: `The Subnet mask used for this IP Restriction.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name for this IP Restriction.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority for this IP Restriction.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Does this restriction ` + "`" + `Allow` + "`" + ` or ` + "`" + `Deny` + "`" + ` access for this IP range? --- An ` + "`" + `scm_ip_restriction` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The IP Address used for this IP Restriction in CIDR notation.`,
				},
				resource.Attribute{
					Name:        "service_tag",
					Description: `The Service Tag used for this IP Restriction.`,
				},
				resource.Attribute{
					Name:        "virtual_network_subnet_id",
					Description: `The Virtual Network Subnet ID used for this IP Restriction.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name for this IP Restriction.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority for this IP Restriction.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Allow or Deny access for this IP range. Defaults to Allow. --- A ` + "`" + `site_config` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "always_on",
					Description: `Is the app loaded at all times?`,
				},
				resource.Attribute{
					Name:        "app_command_line",
					Description: `App command line to launch.`,
				},
				resource.Attribute{
					Name:        "cors",
					Description: `A ` + "`" + `cors` + "`" + ` block as defined above.`,
				},
				resource.Attribute{
					Name:        "default_documents",
					Description: `The ordering of default documents to load, if an address isn't specified.`,
				},
				resource.Attribute{
					Name:        "dotnet_framework_version",
					Description: `The version of the .net framework's CLR used in this App Service.`,
				},
				resource.Attribute{
					Name:        "http2_enabled",
					Description: `Is HTTP2 Enabled on this App Service?`,
				},
				resource.Attribute{
					Name:        "ftps_state",
					Description: `State of FTP / FTPS service for this AppService.`,
				},
				resource.Attribute{
					Name:        "health_check_path",
					Description: `The health check path to be pinged by App Service.`,
				},
				resource.Attribute{
					Name:        "number_of_workers",
					Description: `The scaled number of workers (for per site scaling) of this App Service.`,
				},
				resource.Attribute{
					Name:        "ip_restriction",
					Description: `One or more ` + "`" + `ip_restriction` + "`" + ` blocks as defined above.`,
				},
				resource.Attribute{
					Name:        "scm_use_main_ip_restriction",
					Description: `IP security restrictions for scm to use main.`,
				},
				resource.Attribute{
					Name:        "scm_ip_restriction",
					Description: `One or more ` + "`" + `scm_ip_restriction` + "`" + ` blocks as defined above.`,
				},
				resource.Attribute{
					Name:        "java_version",
					Description: `The version of Java in use.`,
				},
				resource.Attribute{
					Name:        "java_container",
					Description: `The Java Container in use.`,
				},
				resource.Attribute{
					Name:        "java_container_version",
					Description: `The version of the Java Container in use.`,
				},
				resource.Attribute{
					Name:        "linux_fx_version",
					Description: `Linux App Framework and version for the AppService.`,
				},
				resource.Attribute{
					Name:        "windows_fx_version",
					Description: `Windows Container Docker Image for the AppService.`,
				},
				resource.Attribute{
					Name:        "local_mysql_enabled",
					Description: `Is "MySQL In App" Enabled? This runs a local MySQL instance with your app and shares resources from the App Service plan.`,
				},
				resource.Attribute{
					Name:        "managed_pipeline_mode",
					Description: `The Managed Pipeline Mode used in this App Service.`,
				},
				resource.Attribute{
					Name:        "min_tls_version",
					Description: `The minimum supported TLS version for this App Service.`,
				},
				resource.Attribute{
					Name:        "php_version",
					Description: `The version of PHP used in this App Service.`,
				},
				resource.Attribute{
					Name:        "python_version",
					Description: `The version of Python used in this App Service.`,
				},
				resource.Attribute{
					Name:        "remote_debugging_enabled",
					Description: `Is Remote Debugging Enabled in this App Service?`,
				},
				resource.Attribute{
					Name:        "remote_debugging_version",
					Description: `Which version of Visual Studio is the Remote Debugger compatible with?`,
				},
				resource.Attribute{
					Name:        "scm_type",
					Description: `The type of Source Control enabled for this App Service.`,
				},
				resource.Attribute{
					Name:        "use_32_bit_worker_process",
					Description: `Does the App Service run in 32 bit mode, rather than 64 bit mode?`,
				},
				resource.Attribute{
					Name:        "websockets_enabled",
					Description: `Are WebSockets enabled for this App Service? --- A ` + "`" + `source_control` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "repo_url",
					Description: `The URL of the source code repository.`,
				},
				resource.Attribute{
					Name:        "branch",
					Description: `The branch of the remote repository in use.`,
				},
				resource.Attribute{
					Name:        "manual_integration",
					Description: `Limits to manual integration.`,
				},
				resource.Attribute{
					Name:        "rollback_enabled",
					Description: `Is roll-back enabled for the repository.`,
				},
				resource.Attribute{
					Name:        "use_mercurial",
					Description: `Uses Mercurial if ` + "`" + `true` + "`" + `, otherwise uses Git. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the App Service.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the App Service.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the App Service exists.`,
				},
				resource.Attribute{
					Name:        "app_service_plan_id",
					Description: `The ID of the App Service Plan within which the App Service exists.`,
				},
				resource.Attribute{
					Name:        "app_settings",
					Description: `A key-value pair of App Settings for the App Service.`,
				},
				resource.Attribute{
					Name:        "connection_string",
					Description: `An ` + "`" + `connection_string` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "client_affinity_enabled",
					Description: `Does the App Service send session affinity cookies, which route client requests in the same session to the same instance?`,
				},
				resource.Attribute{
					Name:        "custom_domain_verification_id",
					Description: `An identifier used by App Service to perform domain ownership verification via DNS TXT record.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Is the App Service Enabled?`,
				},
				resource.Attribute{
					Name:        "https_only",
					Description: `Can the App Service only be accessed via HTTPS?`,
				},
				resource.Attribute{
					Name:        "client_cert_enabled",
					Description: `Does the App Service require client certificates for incoming requests?`,
				},
				resource.Attribute{
					Name:        "site_config",
					Description: `A ` + "`" + `site_config` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "default_site_hostname",
					Description: `The Default Hostname associated with the App Service - such as ` + "`" + `mysite.azurewebsites.net` + "`" + ``,
				},
				resource.Attribute{
					Name:        "outbound_ip_addresses",
					Description: `A comma separated list of outbound IP addresses - such as ` + "`" + `52.23.25.3,52.143.43.12` + "`" + ``,
				},
				resource.Attribute{
					Name:        "outbound_ip_address_list",
					Description: `A list of outbound IP addresses - such as ` + "`" + `["52.23.25.3", "52.143.43.12"]` + "`" + ``,
				},
				resource.Attribute{
					Name:        "possible_outbound_ip_addresses",
					Description: `A comma separated list of outbound IP addresses - such as ` + "`" + `52.23.25.3,52.143.43.12,52.143.43.17` + "`" + ` - not all of which are necessarily in use. Superset of ` + "`" + `outbound_ip_addresses` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "possible_outbound_ip_address_list",
					Description: `A list of outbound IP addresses - such as ` + "`" + `["52.23.25.3", "52.143.43.12", "52.143.43.17"]` + "`" + ` - not all of which are necessarily in use. Superset of ` + "`" + `outbound_ip_address_list` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "source_control",
					Description: `A ` + "`" + `source_control` + "`" + ` block as defined below. --- A ` + "`" + `connection_string` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Connection String.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the Connection String.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value for the Connection String. --- A ` + "`" + `cors` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "allowed_origins",
					Description: `A list of origins which are able to make cross-origin calls.`,
				},
				resource.Attribute{
					Name:        "support_credentials",
					Description: `Are credentials supported? --- An ` + "`" + `ip_restriction` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The IP Address used for this IP Restriction.`,
				},
				resource.Attribute{
					Name:        "service_tag",
					Description: `The Service Tag used for this IP Restriction.`,
				},
				resource.Attribute{
					Name:        "subnet_mask",
					Description: `The Subnet mask used for this IP Restriction.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name for this IP Restriction.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority for this IP Restriction.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Does this restriction ` + "`" + `Allow` + "`" + ` or ` + "`" + `Deny` + "`" + ` access for this IP range? --- An ` + "`" + `scm_ip_restriction` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The IP Address used for this IP Restriction in CIDR notation.`,
				},
				resource.Attribute{
					Name:        "service_tag",
					Description: `The Service Tag used for this IP Restriction.`,
				},
				resource.Attribute{
					Name:        "virtual_network_subnet_id",
					Description: `The Virtual Network Subnet ID used for this IP Restriction.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name for this IP Restriction.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority for this IP Restriction.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Allow or Deny access for this IP range. Defaults to Allow. --- A ` + "`" + `site_config` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "always_on",
					Description: `Is the app loaded at all times?`,
				},
				resource.Attribute{
					Name:        "app_command_line",
					Description: `App command line to launch.`,
				},
				resource.Attribute{
					Name:        "cors",
					Description: `A ` + "`" + `cors` + "`" + ` block as defined above.`,
				},
				resource.Attribute{
					Name:        "default_documents",
					Description: `The ordering of default documents to load, if an address isn't specified.`,
				},
				resource.Attribute{
					Name:        "dotnet_framework_version",
					Description: `The version of the .net framework's CLR used in this App Service.`,
				},
				resource.Attribute{
					Name:        "http2_enabled",
					Description: `Is HTTP2 Enabled on this App Service?`,
				},
				resource.Attribute{
					Name:        "ftps_state",
					Description: `State of FTP / FTPS service for this AppService.`,
				},
				resource.Attribute{
					Name:        "health_check_path",
					Description: `The health check path to be pinged by App Service.`,
				},
				resource.Attribute{
					Name:        "number_of_workers",
					Description: `The scaled number of workers (for per site scaling) of this App Service.`,
				},
				resource.Attribute{
					Name:        "ip_restriction",
					Description: `One or more ` + "`" + `ip_restriction` + "`" + ` blocks as defined above.`,
				},
				resource.Attribute{
					Name:        "scm_use_main_ip_restriction",
					Description: `IP security restrictions for scm to use main.`,
				},
				resource.Attribute{
					Name:        "scm_ip_restriction",
					Description: `One or more ` + "`" + `scm_ip_restriction` + "`" + ` blocks as defined above.`,
				},
				resource.Attribute{
					Name:        "java_version",
					Description: `The version of Java in use.`,
				},
				resource.Attribute{
					Name:        "java_container",
					Description: `The Java Container in use.`,
				},
				resource.Attribute{
					Name:        "java_container_version",
					Description: `The version of the Java Container in use.`,
				},
				resource.Attribute{
					Name:        "linux_fx_version",
					Description: `Linux App Framework and version for the AppService.`,
				},
				resource.Attribute{
					Name:        "windows_fx_version",
					Description: `Windows Container Docker Image for the AppService.`,
				},
				resource.Attribute{
					Name:        "local_mysql_enabled",
					Description: `Is "MySQL In App" Enabled? This runs a local MySQL instance with your app and shares resources from the App Service plan.`,
				},
				resource.Attribute{
					Name:        "managed_pipeline_mode",
					Description: `The Managed Pipeline Mode used in this App Service.`,
				},
				resource.Attribute{
					Name:        "min_tls_version",
					Description: `The minimum supported TLS version for this App Service.`,
				},
				resource.Attribute{
					Name:        "php_version",
					Description: `The version of PHP used in this App Service.`,
				},
				resource.Attribute{
					Name:        "python_version",
					Description: `The version of Python used in this App Service.`,
				},
				resource.Attribute{
					Name:        "remote_debugging_enabled",
					Description: `Is Remote Debugging Enabled in this App Service?`,
				},
				resource.Attribute{
					Name:        "remote_debugging_version",
					Description: `Which version of Visual Studio is the Remote Debugger compatible with?`,
				},
				resource.Attribute{
					Name:        "scm_type",
					Description: `The type of Source Control enabled for this App Service.`,
				},
				resource.Attribute{
					Name:        "use_32_bit_worker_process",
					Description: `Does the App Service run in 32 bit mode, rather than 64 bit mode?`,
				},
				resource.Attribute{
					Name:        "websockets_enabled",
					Description: `Are WebSockets enabled for this App Service? --- A ` + "`" + `source_control` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "repo_url",
					Description: `The URL of the source code repository.`,
				},
				resource.Attribute{
					Name:        "branch",
					Description: `The branch of the remote repository in use.`,
				},
				resource.Attribute{
					Name:        "manual_integration",
					Description: `Limits to manual integration.`,
				},
				resource.Attribute{
					Name:        "rollback_enabled",
					Description: `Is roll-back enabled for the repository.`,
				},
				resource.Attribute{
					Name:        "use_mercurial",
					Description: `Uses Mercurial if ` + "`" + `true` + "`" + `, otherwise uses Git. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the App Service.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_app_service_certificate",
			Category:         "App Service (Web Apps)",
			ShortDescription: `Gets information about an existing App Service Certificate.`,
			Description: `

Use this data source to access information about an App Service Certificate.

`,
			Keywords: []string{
				"app",
				"service",
				"web",
				"apps",
				"certificate",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name of the certificate.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of the resource group in which to create the certificate. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The App Service certificate ID.`,
				},
				resource.Attribute{
					Name:        "friendly_name",
					Description: `The friendly name of the certificate.`,
				},
				resource.Attribute{
					Name:        "subject_name",
					Description: `The subject name of the certificate.`,
				},
				resource.Attribute{
					Name:        "host_names",
					Description: `List of host names the certificate applies to.`,
				},
				resource.Attribute{
					Name:        "issuer",
					Description: `The name of the certificate issuer.`,
				},
				resource.Attribute{
					Name:        "issue_date",
					Description: `The issue date for the certificate.`,
				},
				resource.Attribute{
					Name:        "expiration_date",
					Description: `The expiration date for the certificate.`,
				},
				resource.Attribute{
					Name:        "thumbprint",
					Description: `The thumbprint for the certificate. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the App Service Certificate.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The App Service certificate ID.`,
				},
				resource.Attribute{
					Name:        "friendly_name",
					Description: `The friendly name of the certificate.`,
				},
				resource.Attribute{
					Name:        "subject_name",
					Description: `The subject name of the certificate.`,
				},
				resource.Attribute{
					Name:        "host_names",
					Description: `List of host names the certificate applies to.`,
				},
				resource.Attribute{
					Name:        "issuer",
					Description: `The name of the certificate issuer.`,
				},
				resource.Attribute{
					Name:        "issue_date",
					Description: `The issue date for the certificate.`,
				},
				resource.Attribute{
					Name:        "expiration_date",
					Description: `The expiration date for the certificate.`,
				},
				resource.Attribute{
					Name:        "thumbprint",
					Description: `The thumbprint for the certificate. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the App Service Certificate.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_app_service_certificate_order",
			Category:         "App Service (Web Apps)",
			ShortDescription: `Gets information about an existing App Service Certificate Order.`,
			Description: `

Use this data source to access information about an existing App Service Certificate Order.

`,
			Keywords: []string{
				"app",
				"service",
				"web",
				"apps",
				"certificate",
				"order",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the App Service.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The Name of the Resource Group where the App Service exists. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the App Service.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the App Service exists.`,
				},
				resource.Attribute{
					Name:        "auto_renew",
					Description: `true if the certificate should be automatically renewed when it expires; otherwise, false.`,
				},
				resource.Attribute{
					Name:        "certificates",
					Description: `State of the Key Vault secret. A ` + "`" + `certificates` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "csr",
					Description: `Last CSR that was created for this order.`,
				},
				resource.Attribute{
					Name:        "distinguished_name",
					Description: `The Distinguished Name for the App Service Certificate Order.`,
				},
				resource.Attribute{
					Name:        "key_size",
					Description: `Certificate key size.`,
				},
				resource.Attribute{
					Name:        "product_type",
					Description: `Certificate product type, such as ` + "`" + `Standard` + "`" + ` or ` + "`" + `WildCard` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "validity_in_years",
					Description: `Duration in years (must be between 1 and 3).`,
				},
				resource.Attribute{
					Name:        "domain_verification_token",
					Description: `Domain verification token.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current order status.`,
				},
				resource.Attribute{
					Name:        "expiration_time",
					Description: `Certificate expiration time.`,
				},
				resource.Attribute{
					Name:        "is_private_key_external",
					Description: `Whether the private key is external or not.`,
				},
				resource.Attribute{
					Name:        "app_service_certificate_not_renewable_reasons",
					Description: `Reasons why App Service Certificate is not renewable at the current moment.`,
				},
				resource.Attribute{
					Name:        "signed_certificate_thumbprint",
					Description: `Certificate thumbprint for signed certificate.`,
				},
				resource.Attribute{
					Name:        "root_thumbprint",
					Description: `Certificate thumbprint for root certificate.`,
				},
				resource.Attribute{
					Name:        "intermediate_thumbprint",
					Description: `Certificate thumbprint intermediate certificate.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the resource. ---`,
				},
				resource.Attribute{
					Name:        "certificate_name",
					Description: `The name of the App Service Certificate.`,
				},
				resource.Attribute{
					Name:        "key_vault_id",
					Description: `Key Vault resource Id.`,
				},
				resource.Attribute{
					Name:        "key_vault_secret_name",
					Description: `Key Vault secret name.`,
				},
				resource.Attribute{
					Name:        "provisioning_state",
					Description: `Status of the Key Vault secret. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the App Service Certificate Order.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the App Service.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the App Service exists.`,
				},
				resource.Attribute{
					Name:        "auto_renew",
					Description: `true if the certificate should be automatically renewed when it expires; otherwise, false.`,
				},
				resource.Attribute{
					Name:        "certificates",
					Description: `State of the Key Vault secret. A ` + "`" + `certificates` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "csr",
					Description: `Last CSR that was created for this order.`,
				},
				resource.Attribute{
					Name:        "distinguished_name",
					Description: `The Distinguished Name for the App Service Certificate Order.`,
				},
				resource.Attribute{
					Name:        "key_size",
					Description: `Certificate key size.`,
				},
				resource.Attribute{
					Name:        "product_type",
					Description: `Certificate product type, such as ` + "`" + `Standard` + "`" + ` or ` + "`" + `WildCard` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "validity_in_years",
					Description: `Duration in years (must be between 1 and 3).`,
				},
				resource.Attribute{
					Name:        "domain_verification_token",
					Description: `Domain verification token.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current order status.`,
				},
				resource.Attribute{
					Name:        "expiration_time",
					Description: `Certificate expiration time.`,
				},
				resource.Attribute{
					Name:        "is_private_key_external",
					Description: `Whether the private key is external or not.`,
				},
				resource.Attribute{
					Name:        "app_service_certificate_not_renewable_reasons",
					Description: `Reasons why App Service Certificate is not renewable at the current moment.`,
				},
				resource.Attribute{
					Name:        "signed_certificate_thumbprint",
					Description: `Certificate thumbprint for signed certificate.`,
				},
				resource.Attribute{
					Name:        "root_thumbprint",
					Description: `Certificate thumbprint for root certificate.`,
				},
				resource.Attribute{
					Name:        "intermediate_thumbprint",
					Description: `Certificate thumbprint intermediate certificate.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the resource. ---`,
				},
				resource.Attribute{
					Name:        "certificate_name",
					Description: `The name of the App Service Certificate.`,
				},
				resource.Attribute{
					Name:        "key_vault_id",
					Description: `Key Vault resource Id.`,
				},
				resource.Attribute{
					Name:        "key_vault_secret_name",
					Description: `Key Vault secret name.`,
				},
				resource.Attribute{
					Name:        "provisioning_state",
					Description: `Status of the Key Vault secret. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the App Service Certificate Order.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_app_service_environment",
			Category:         "App Service (Web Apps)",
			ShortDescription: `Gets information about an existing App Service Environment.`,
			Description: `

Use this data source to access information about an existing App Service Environment.

`,
			Keywords: []string{
				"app",
				"service",
				"web",
				"apps",
				"environment",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the App Service Environment.`,
				},
				resource.Attribute{
					Name:        "cluster_setting",
					Description: `Zero or more ` + "`" + `cluster_setting` + "`" + ` blocks as defined below.`,
				},
				resource.Attribute{
					Name:        "front_end_scale_factor",
					Description: `The number of app instances per App Service Environment Front End.`,
				},
				resource.Attribute{
					Name:        "internal_ip_address",
					Description: `IP address of internal load balancer of the App Service Environment.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure Region where the App Service Environment exists.`,
				},
				resource.Attribute{
					Name:        "outbound_ip_addresses",
					Description: `List of outbound IP addresses of the App Service Environment.`,
				},
				resource.Attribute{
					Name:        "pricing_tier",
					Description: `The Pricing Tier (Isolated SKU) of the App Service Environment.`,
				},
				resource.Attribute{
					Name:        "service_ip_address",
					Description: `IP address of service endpoint of the App Service Environment.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the App Service Environment. --- A ` + "`" + `cluster_setting` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Cluster Setting.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value for the Cluster Setting. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions: - ` + "`" + `read` + "`" + ` - (Defaults to 5 minutes) Used when retrieving the App Service Environment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_app_service_plan",
			Category:         "App Service (Web Apps)",
			ShortDescription: `Gets information about an existing App Service Plan.`,
			Description: `

Use this data source to access information about an existing App Service Plan (formerly known as a ` + "`" + `Server Farm` + "`" + `).

`,
			Keywords: []string{
				"app",
				"service",
				"web",
				"apps",
				"plan",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the App Service Plan.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The Name of the Resource Group where the App Service Plan exists. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the App Service Plan.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the App Service Plan exists`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `The Operating System type of the App Service Plan`,
				},
				resource.Attribute{
					Name:        "sku",
					Description: `A ` + "`" + `sku` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "app_service_environment_id",
					Description: `The ID of the App Service Environment where the App Service Plan is located.`,
				},
				resource.Attribute{
					Name:        "maximum_number_of_workers",
					Description: `Maximum number of instances that can be assigned to this App Service plan.`,
				},
				resource.Attribute{
					Name:        "reserved",
					Description: `Is this App Service Plan ` + "`" + `Reserved` + "`" + `?`,
				},
				resource.Attribute{
					Name:        "per_site_scaling",
					Description: `Can Apps assigned to this App Service Plan be scaled independently?`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource.`,
				},
				resource.Attribute{
					Name:        "maximum_elastic_worker_count",
					Description: `The maximum number of total workers allowed for this ElasticScaleEnabled App Service Plan.`,
				},
				resource.Attribute{
					Name:        "is_xenon",
					Description: `A flag that indicates if it's a xenon plan (support for Windows Container)`,
				},
				resource.Attribute{
					Name:        "maximum_number_of_workers",
					Description: `The maximum number of workers supported with the App Service Plan's sku. --- A ` + "`" + `sku` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "tier",
					Description: `Specifies the plan's pricing tier.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Specifies the plan's instance size.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `Specifies the number of workers associated with this App Service Plan. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the App Service Plan.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the App Service Plan.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the App Service Plan exists`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `The Operating System type of the App Service Plan`,
				},
				resource.Attribute{
					Name:        "sku",
					Description: `A ` + "`" + `sku` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "app_service_environment_id",
					Description: `The ID of the App Service Environment where the App Service Plan is located.`,
				},
				resource.Attribute{
					Name:        "maximum_number_of_workers",
					Description: `Maximum number of instances that can be assigned to this App Service plan.`,
				},
				resource.Attribute{
					Name:        "reserved",
					Description: `Is this App Service Plan ` + "`" + `Reserved` + "`" + `?`,
				},
				resource.Attribute{
					Name:        "per_site_scaling",
					Description: `Can Apps assigned to this App Service Plan be scaled independently?`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource.`,
				},
				resource.Attribute{
					Name:        "maximum_elastic_worker_count",
					Description: `The maximum number of total workers allowed for this ElasticScaleEnabled App Service Plan.`,
				},
				resource.Attribute{
					Name:        "is_xenon",
					Description: `A flag that indicates if it's a xenon plan (support for Windows Container)`,
				},
				resource.Attribute{
					Name:        "maximum_number_of_workers",
					Description: `The maximum number of workers supported with the App Service Plan's sku. --- A ` + "`" + `sku` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "tier",
					Description: `Specifies the plan's pricing tier.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Specifies the plan's instance size.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `Specifies the number of workers associated with this App Service Plan. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the App Service Plan.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_application_gateway",
			Category:         "Network",
			ShortDescription: `Gets information about an existing Application Gateway.`,
			Description: `

Use this data source to access information about an existing Application Gateway.

`,
			Icon: "Other/10076-icon-service-Application-Gateways.svg",
			Keywords: []string{
				"network",
				"application",
				"gateway",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Application Gateway.`,
				},
				resource.Attribute{
					Name:        "identity",
					Description: `A ` + "`" + `identity` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure Region where the Application Gateway exists.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the Application Gateway. --- A ` + "`" + `identity` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "identity_ids",
					Description: `A list of Managed Identity ID's assigned to this Application Gateway.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of Managed Identity assigned to this Application Gateway. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Application Gateway.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_application_insights",
			Category:         "Application Insights",
			ShortDescription: `Gets information about an existing Application Insights component.`,
			Description: `

Use this data source to access information about an existing Application Insights component.

`,
			Icon: "Management_and_Governance/00012-icon-service-Application-Insights.svg",
			Keywords: []string{
				"application",
				"insights",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name of the Application Insights component.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `Specifies the name of the resource group the Application Insights component is located in. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Virtual Machine.`,
				},
				resource.Attribute{
					Name:        "app_id",
					Description: `The App ID associated with this Application Insights component.`,
				},
				resource.Attribute{
					Name:        "application_type",
					Description: `The type of the component.`,
				},
				resource.Attribute{
					Name:        "instrumentation_key",
					Description: `The instrumentation key of the Application Insights component.`,
				},
				resource.Attribute{
					Name:        "connection_string",
					Description: `The connection string of the Application Insights component. (Sensitive)`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the component exists.`,
				},
				resource.Attribute{
					Name:        "retention_in_days",
					Description: `The retention period in days.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags applied to the component. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Application Insights component.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Virtual Machine.`,
				},
				resource.Attribute{
					Name:        "app_id",
					Description: `The App ID associated with this Application Insights component.`,
				},
				resource.Attribute{
					Name:        "application_type",
					Description: `The type of the component.`,
				},
				resource.Attribute{
					Name:        "instrumentation_key",
					Description: `The instrumentation key of the Application Insights component.`,
				},
				resource.Attribute{
					Name:        "connection_string",
					Description: `The connection string of the Application Insights component. (Sensitive)`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the component exists.`,
				},
				resource.Attribute{
					Name:        "retention_in_days",
					Description: `The retention period in days.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags applied to the component. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Application Insights component.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_application_security_group",
			Category:         "Network",
			ShortDescription: `Gets information about an existing Application Security Group.`,
			Description: `

Use this data source to access information about an existing Application Security Group.

`,
			Icon: "Networking/10083-icon-service-Application-Security-Groups.svg",
			Keywords: []string{
				"network",
				"application",
				"security",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Application Security Group.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of the resource group in which the Application Security Group exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Application Security Group.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The supported Azure location where the Application Security Group exists.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Application Security Group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Application Security Group.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The supported Azure location where the Application Security Group exists.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Application Security Group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_attestation",
			Category:         "Attestation",
			ShortDescription: `Gets information about an existing Attestation Provider.`,
			Description: `

Use this data source to access information about an existing Attestation Provider.

`,
			Keywords: []string{
				"attestation",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Attestation Provider.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure Region where the Attestation Provider exists.`,
				},
				resource.Attribute{
					Name:        "attestation_uri",
					Description: `The (Endpoint|URI) of the Attestation Service.`,
				},
				resource.Attribute{
					Name:        "trust_model",
					Description: `Trust model used for the Attestation Service.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the Attestation Provider. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Attestation Provider.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_automation_account",
			Category:         "Automation",
			ShortDescription: `Gets information about an existing Automation Account.`,
			Description: `

Use this data source to access information about an existing Automation Account.

`,
			Icon: "Management_and_Governance/00022-icon-service-Automation-Accounts.svg",
			Keywords: []string{
				"automation",
				"account",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Automation Account.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `Specifies the name of the Resource Group where the Automation Account exists. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Automation Account`,
				},
				resource.Attribute{
					Name:        "primary_key",
					Description: `The Primary Access Key for the Automation Account.`,
				},
				resource.Attribute{
					Name:        "secondary_key",
					Description: `The Secondary Access Key for the Automation Account.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `The Endpoint for this Automation Account. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Automation Account.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Automation Account`,
				},
				resource.Attribute{
					Name:        "primary_key",
					Description: `The Primary Access Key for the Automation Account.`,
				},
				resource.Attribute{
					Name:        "secondary_key",
					Description: `The Secondary Access Key for the Automation Account.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `The Endpoint for this Automation Account. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Automation Account.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_automation_variable_bool",
			Category:         "Automation",
			ShortDescription: `Gets information about an existing Automation Bool Variable`,
			Description: `

Use this data source to access information about an existing Automation Bool Variable.


`,
			Keywords: []string{
				"automation",
				"variable",
				"bool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Automation Variable.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The Name of the Resource Group where the automation account exists.`,
				},
				resource.Attribute{
					Name:        "automation_account_name",
					Description: `The name of the automation account in which the Automation Variable exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Automation Variable.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the Automation Variable.`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `Specifies if the Automation Variable is encrypted. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the Automation Variable as a ` + "`" + `boolean` + "`" + `. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Automation Bool Variable.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Automation Variable.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the Automation Variable.`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `Specifies if the Automation Variable is encrypted. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the Automation Variable as a ` + "`" + `boolean` + "`" + `. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Automation Bool Variable.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_automation_variable_datetime",
			Category:         "Automation",
			ShortDescription: `Gets information about an existing Automation Datetime Variable`,
			Description: `

Use this data source to access information about an existing Automation Datetime Variable.


`,
			Keywords: []string{
				"automation",
				"variable",
				"datetime",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Automation Variable.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The Name of the Resource Group where the automation account exists.`,
				},
				resource.Attribute{
					Name:        "automation_account_name",
					Description: `The name of the automation account in which the Automation Variable exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Automation Variable.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the Automation Variable.`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `Specifies if the Automation Variable is encrypted. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the Automation Variable in the [RFC3339 Section 5.6 Internet Date/Time Format](https://tools.ietf.org/html/rfc3339#section-5.6). ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Automation Datetime Variable.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Automation Variable.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the Automation Variable.`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `Specifies if the Automation Variable is encrypted. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the Automation Variable in the [RFC3339 Section 5.6 Internet Date/Time Format](https://tools.ietf.org/html/rfc3339#section-5.6). ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Automation Datetime Variable.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_automation_variable_int",
			Category:         "Automation",
			ShortDescription: `Gets information about an existing Automation Int Variable`,
			Description: `

Use this data source to access information about an existing Automation Int Variable.


`,
			Keywords: []string{
				"automation",
				"variable",
				"int",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Automation Variable.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The Name of the Resource Group where the automation account exists.`,
				},
				resource.Attribute{
					Name:        "automation_account_name",
					Description: `The name of the automation account in which the Automation Variable exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Automation Variable.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the Automation Variable.`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `Specifies if the Automation Variable is encrypted. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the Automation Variable as a ` + "`" + `integer` + "`" + `. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Automation Int Variable.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Automation Variable.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the Automation Variable.`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `Specifies if the Automation Variable is encrypted. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the Automation Variable as a ` + "`" + `integer` + "`" + `. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Automation Int Variable.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_automation_variable_string",
			Category:         "Automation",
			ShortDescription: `Gets information about an existing Automation String Variable`,
			Description: `

Use this data source to access information about an existing Automation String Variable.


`,
			Keywords: []string{
				"automation",
				"variable",
				"string",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Automation Variable.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The Name of the Resource Group where the automation account exists.`,
				},
				resource.Attribute{
					Name:        "automation_account_name",
					Description: `The name of the automation account in which the Automation Variable exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Automation Variable.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the Automation Variable.`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `Specifies if the Automation Variable is encrypted. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the Automation Variable as a ` + "`" + `string` + "`" + `. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Automation String Variable.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Automation Variable.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the Automation Variable.`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `Specifies if the Automation Variable is encrypted. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the Automation Variable as a ` + "`" + `string` + "`" + `. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Automation String Variable.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_availability_set",
			Category:         "Compute",
			ShortDescription: `Gets information about an existing Availability Set.`,
			Description: `

Use this data source to access information about an existing Availability Set.

`,
			Icon: "Compute/10025-icon-service-Availability-Sets.svg",
			Keywords: []string{
				"compute",
				"availability",
				"set",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Availability Set.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of the resource group in which the Availability Set exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Availability Set.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The supported Azure location where the Availability Set exists.`,
				},
				resource.Attribute{
					Name:        "managed",
					Description: `Whether the availability set is managed or not.`,
				},
				resource.Attribute{
					Name:        "platform_fault_domain_count",
					Description: `The number of fault domains that are used.`,
				},
				resource.Attribute{
					Name:        "platform_update_domain_count",
					Description: `The number of update domains that are used.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Availability Set.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Availability Set.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The supported Azure location where the Availability Set exists.`,
				},
				resource.Attribute{
					Name:        "managed",
					Description: `Whether the availability set is managed or not.`,
				},
				resource.Attribute{
					Name:        "platform_fault_domain_count",
					Description: `The number of fault domains that are used.`,
				},
				resource.Attribute{
					Name:        "platform_update_domain_count",
					Description: `The number of update domains that are used.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Availability Set.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_backup_policy_vm",
			Category:         "Recovery Services",
			ShortDescription: `Gets information about an existing existing VM Backup Policy.`,
			Description: `

Use this data source to access information about an existing VM Backup Policy.

`,
			Keywords: []string{
				"recovery",
				"services",
				"backup",
				"policy",
				"vm",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name of the VM Backup Policy.`,
				},
				resource.Attribute{
					Name:        "recovery_vault_name",
					Description: `Specifies the name of the Recovery Services Vault.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of the resource group in which the VM Backup Policy resides. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Backup VM Protection Policy.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Recovery Services VM Protection Policy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Backup VM Protection Policy.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Recovery Services VM Protection Policy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_batch_account",
			Category:         "Batch",
			ShortDescription: `Get information about an existing Batch Account`,
			Description: `

Use this data source to access information about an existing Batch Account.

`,
			Icon: "Containers/10031-icon-service-Batch-Accounts.svg",
			Keywords: []string{
				"batch",
				"account",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Batch account.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The Name of the Resource Group where this Batch account exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Batch account ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The Batch account name.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure Region in which this Batch account exists.`,
				},
				resource.Attribute{
					Name:        "pool_allocation_mode",
					Description: `The pool allocation mode configured for this Batch account.`,
				},
				resource.Attribute{
					Name:        "storage_account_id",
					Description: `The ID of the Storage Account used for this Batch account.`,
				},
				resource.Attribute{
					Name:        "primary_access_key",
					Description: `The Batch account primary access key.`,
				},
				resource.Attribute{
					Name:        "secondary_access_key",
					Description: `The Batch account secondary access key.`,
				},
				resource.Attribute{
					Name:        "account_endpoint",
					Description: `The account endpoint used to interact with the Batch service.`,
				},
				resource.Attribute{
					Name:        "key_vault_reference",
					Description: `The ` + "`" + `key_vault_reference` + "`" + ` block that describes the Azure KeyVault reference to use when deploying the Azure Batch account using the ` + "`" + `UserSubscription` + "`" + ` pool allocation mode.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A map of tags assigned to the Batch account. ~>`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Azure identifier of the Azure KeyVault reference.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `The HTTPS URL of the Azure KeyVault reference. --- ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Batch Account.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Batch account ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The Batch account name.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure Region in which this Batch account exists.`,
				},
				resource.Attribute{
					Name:        "pool_allocation_mode",
					Description: `The pool allocation mode configured for this Batch account.`,
				},
				resource.Attribute{
					Name:        "storage_account_id",
					Description: `The ID of the Storage Account used for this Batch account.`,
				},
				resource.Attribute{
					Name:        "primary_access_key",
					Description: `The Batch account primary access key.`,
				},
				resource.Attribute{
					Name:        "secondary_access_key",
					Description: `The Batch account secondary access key.`,
				},
				resource.Attribute{
					Name:        "account_endpoint",
					Description: `The account endpoint used to interact with the Batch service.`,
				},
				resource.Attribute{
					Name:        "key_vault_reference",
					Description: `The ` + "`" + `key_vault_reference` + "`" + ` block that describes the Azure KeyVault reference to use when deploying the Azure Batch account using the ` + "`" + `UserSubscription` + "`" + ` pool allocation mode.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A map of tags assigned to the Batch account. ~>`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Azure identifier of the Azure KeyVault reference.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `The HTTPS URL of the Azure KeyVault reference. --- ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Batch Account.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_batch_certificate",
			Category:         "Batch",
			ShortDescription: `Get information about an existing certificate in a Batch Account`,
			Description: `

Use this data source to access information about an existing certificate in a Batch Account.

`,
			Keywords: []string{
				"batch",
				"certificate",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Batch certificate.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `The name of the Batch account.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The Name of the Resource Group where this Batch account exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Batch certificate ID.`,
				},
				resource.Attribute{
					Name:        "public_data",
					Description: `The public key of the certificate.`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `The format of the certificate, such as ` + "`" + `Cer` + "`" + ` or ` + "`" + `Pfx` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "thumbprint",
					Description: `The thumbprint of the certificate.`,
				},
				resource.Attribute{
					Name:        "thumbprint_algorithm",
					Description: `The algorithm of the certificate thumbprint. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the certificate.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Batch certificate ID.`,
				},
				resource.Attribute{
					Name:        "public_data",
					Description: `The public key of the certificate.`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `The format of the certificate, such as ` + "`" + `Cer` + "`" + ` or ` + "`" + `Pfx` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "thumbprint",
					Description: `The thumbprint of the certificate.`,
				},
				resource.Attribute{
					Name:        "thumbprint_algorithm",
					Description: `The algorithm of the certificate thumbprint. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the certificate.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_batch_pool",
			Category:         "Batch",
			ShortDescription: `Get information about an existing Azure Batch pool.`,
			Description:      ``,
			Keywords: []string{
				"batch",
				"pool",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Batch pool ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Batch pool.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `The name of the Batch account.`,
				},
				resource.Attribute{
					Name:        "node_agent_sku_id",
					Description: `The Sku of the node agents in the Batch pool.`,
				},
				resource.Attribute{
					Name:        "vm_size",
					Description: `The size of the VM created in the Batch pool.`,
				},
				resource.Attribute{
					Name:        "fixed_scale",
					Description: `A ` + "`" + `fixed_scale` + "`" + ` block that describes the scale settings when using fixed scale.`,
				},
				resource.Attribute{
					Name:        "auto_scale",
					Description: `A ` + "`" + `auto_scale` + "`" + ` block that describes the scale settings when using auto scale.`,
				},
				resource.Attribute{
					Name:        "storage_image_reference",
					Description: `The reference of the storage image used by the nodes in the Batch pool.`,
				},
				resource.Attribute{
					Name:        "start_task",
					Description: `A ` + "`" + `start_task` + "`" + ` block that describes the start task settings for the Batch pool.`,
				},
				resource.Attribute{
					Name:        "max_tasks_per_node",
					Description: `The maximum number of tasks that can run concurrently on a single compute node in the pool.`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `One or more ` + "`" + `certificate` + "`" + ` blocks that describe the certificates installed on each compute node in the pool.`,
				},
				resource.Attribute{
					Name:        "container_configuration",
					Description: `The container configuration used in the pool's VMs. --- A ` + "`" + `fixed_scale` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "target_dedicated_nodes",
					Description: `The number of nodes in the Batch pool.`,
				},
				resource.Attribute{
					Name:        "target_low_priority_nodes",
					Description: `The number of low priority nodes in the Batch pool.`,
				},
				resource.Attribute{
					Name:        "resize_timeout",
					Description: `The timeout for resize operations. --- A ` + "`" + `auto_scale` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "evaluation_interval",
					Description: `The interval to wait before evaluating if the pool needs to be scaled.`,
				},
				resource.Attribute{
					Name:        "formula",
					Description: `The autoscale formula that needs to be used for scaling the Batch pool. --- A ` + "`" + `start_task` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "command_line",
					Description: `The command line executed by the start task.`,
				},
				resource.Attribute{
					Name:        "max_task_retry_count",
					Description: `The number of retry count.`,
				},
				resource.Attribute{
					Name:        "wait_for_success",
					Description: `A flag that indicates if the Batch pool should wait for the start task to be completed.`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `A map of strings (key,value) that represents the environment variables to set in the start task.`,
				},
				resource.Attribute{
					Name:        "user_identity",
					Description: `A ` + "`" + `user_identity` + "`" + ` block that describes the user identity under which the start task runs.`,
				},
				resource.Attribute{
					Name:        "resource_file",
					Description: `One or more ` + "`" + `resource_file` + "`" + ` blocks that describe the files to be downloaded to a compute node. --- A ` + "`" + `user_identity` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `The username to be used by the Batch pool start task.`,
				},
				resource.Attribute{
					Name:        "auto_user",
					Description: `A ` + "`" + `auto_user` + "`" + ` block that describes the user identity under which the start task runs. --- A ` + "`" + `auto_user` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "elevation_level",
					Description: `The elevation level of the user identity under which the start task runs.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `The scope of the user identity under which the start task runs. --- A ` + "`" + `certificate` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The fully qualified ID of the certificate installed on the pool.`,
				},
				resource.Attribute{
					Name:        "store_location",
					Description: `The location of the certificate store on the compute node into which the certificate is installed, either ` + "`" + `CurrentUser` + "`" + ` or ` + "`" + `LocalMachine` + "`" + `. ->`,
				},
				resource.Attribute{
					Name:        "store_name",
					Description: `The name of the certificate store on the compute node into which the certificate is installed. ->`,
				},
				resource.Attribute{
					Name:        "visibility",
					Description: `Which user accounts on the compute node have access to the private data of the certificate. --- A ` + "`" + `resource_file` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "auto_storage_container_name",
					Description: `The storage container name in the auto storage account.`,
				},
				resource.Attribute{
					Name:        "blob_prefix",
					Description: `The blob prefix used when downloading blobs from an Azure Storage container.`,
				},
				resource.Attribute{
					Name:        "file_mode",
					Description: `The file permission mode attribute represented as a string in octal format (e.g. ` + "`" + `"0644"` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "file_path",
					Description: `The location on the compute node to which to download the file, relative to the task's working directory. If the ` + "`" + `http_url` + "`" + ` property is specified, the ` + "`" + `file_path` + "`" + ` is required and describes the path which the file will be downloaded to, including the filename. Otherwise, if the ` + "`" + `auto_storage_container_name` + "`" + ` or ` + "`" + `storage_container_url` + "`" + ` property is specified.`,
				},
				resource.Attribute{
					Name:        "http_url",
					Description: `The URL of the file to download. If the URL is Azure Blob Storage, it must be readable using anonymous access.`,
				},
				resource.Attribute{
					Name:        "storage_container_url",
					Description: `The URL of the blob container within Azure Blob Storage. --- A ` + "`" + `container_configuration` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of container configuration.`,
				},
				resource.Attribute{
					Name:        "container_image_names",
					Description: `A list of container image names to use, as would be specified by ` + "`" + `docker pull` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "container_registries",
					Description: `Additional container registries from which container images can be pulled by the pool's VMs. --- A ` + "`" + `container_registries` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "registry_server",
					Description: `The container registry URL. The default is "docker.io".`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `The user name to log into the registry server.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `The password to log into the registry server. --- A ` + "`" + `network_configuration` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ARM resource identifier of the virtual network subnet which the compute nodes of the pool are joined too.`,
				},
				resource.Attribute{
					Name:        "endpoint_configuration",
					Description: `The inbound NAT pools that are used to address specific ports on the individual compute node externally. --- A ` + "`" + `endpoint_configuration` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the endpoint.`,
				},
				resource.Attribute{
					Name:        "backend_port",
					Description: `The port number on the compute node.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol of the endpoint.`,
				},
				resource.Attribute{
					Name:        "frontend_port_range",
					Description: `The range of external ports that are used to provide inbound access to the backendPort on the individual compute nodes in the format of ` + "`" + `1000-1100` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "network_security_group_rules",
					Description: `The list of network security group rules that are applied to the endpoint. --- A ` + "`" + `network_security_group_rules` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "access",
					Description: `The action that should be taken for a specified IP address, subnet range or tag.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority for this rule.`,
				},
				resource.Attribute{
					Name:        "source_address_prefix",
					Description: `The source address prefix or tag to match for the rule. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Batch Pool.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_billing_enrollment_account_scope",
			Category:         "Billing",
			ShortDescription: `This is a helper Data Source to provide a correctly formatted Billing Scope ID for an Enterprise Account Enrollment.`,
			Description: `

Use this data source to access information about an existing Enrollment Account Billing Scope.

`,
			Keywords: []string{
				"billing",
				"enrollment",
				"account",
				"scope",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Enrollment Account Billing Scope. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Enrollment Account Billing Scope.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_billing_mca_account_scope",
			Category:         "Billing",
			ShortDescription: `This is a helper Data Source to provide a correctly formatted Billing Scope ID for a Microsoft Customer Agreement Account.`,
			Description: `

Use this data source to access an ID for your MCA Account billing scope.

`,
			Keywords: []string{
				"billing",
				"mca",
				"account",
				"scope",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Billing Scope. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Billing Scope.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_blueprint_definition",
			Category:         "Blueprints",
			ShortDescription: `Gets information about an existing Blueprint Definition`,
			Description: `

Use this data source to access information about an existing Azure Blueprint Definition

~> **NOTE:** Azure Blueprints are in Preview and potentially subject to breaking change without notice.

`,
			Keywords: []string{
				"blueprints",
				"blueprint",
				"definition",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Blueprint.`,
				},
				resource.Attribute{
					Name:        "scope_id",
					Description: `(Required) The ID of the Subscription or Management Group, as the scope at which the blueprint definition is stored. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Blueprint Definition.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the Blueprint Definition.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name of the Blueprint Definition.`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `The timestamp of when this last modification was saved to the Blueprint Definition.`,
				},
				resource.Attribute{
					Name:        "target_scope",
					Description: `The target scope.`,
				},
				resource.Attribute{
					Name:        "time_created",
					Description: `The timestamp of when this Blueprint Definition was created.`,
				},
				resource.Attribute{
					Name:        "versions",
					Description: `A list of versions published for this Blueprint Definition. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Blueprint Published Version.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Blueprint Definition.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the Blueprint Definition.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name of the Blueprint Definition.`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `The timestamp of when this last modification was saved to the Blueprint Definition.`,
				},
				resource.Attribute{
					Name:        "target_scope",
					Description: `The target scope.`,
				},
				resource.Attribute{
					Name:        "time_created",
					Description: `The timestamp of when this Blueprint Definition was created.`,
				},
				resource.Attribute{
					Name:        "versions",
					Description: `A list of versions published for this Blueprint Definition. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Blueprint Published Version.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_blueprint_published_version",
			Category:         "Blueprints",
			ShortDescription: `Gets information about an existing Blueprint Published Version`,
			Description: `

Use this data source to access information about an existing Blueprint Published Version

~> **NOTE:** Azure Blueprints are in Preview and potentially subject to breaking change without notice.

`,
			Keywords: []string{
				"blueprints",
				"blueprint",
				"published",
				"version",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "blueprint_name",
					Description: `(Required) The name of the Blueprint Definition`,
				},
				resource.Attribute{
					Name:        "scope_id",
					Description: `(Required) The ID of the Management Group / Subscription where this Blueprint Definition is stored.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The Version name of the Published Version of the Blueprint Definition ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Published Version`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the Blueprint`,
				},
				resource.Attribute{
					Name:        "target_scope",
					Description: `The target scope`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name of the Blueprint Published Version`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the Blueprint Published Version ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Blueprint Published Version.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Published Version`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the Blueprint`,
				},
				resource.Attribute{
					Name:        "target_scope",
					Description: `The target scope`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name of the Blueprint Published Version`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the Blueprint Published Version ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Blueprint Published Version.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_cdn_profile",
			Category:         "CDN",
			ShortDescription: `Gets information about an existing CDN Profile`,
			Description: `

Use this data source to access information about an existing CDN Profile.

`,
			Icon: "Networking/00056-icon-service-CDN-Profiles.svg",
			Keywords: []string{
				"cdn",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the CDN Profile.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of the resource group in which the CDN Profile exists. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure Region where the resource exists.`,
				},
				resource.Attribute{
					Name:        "sku",
					Description: `The pricing related information of current CDN profile.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the CDN Profile.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "location",
					Description: `The Azure Region where the resource exists.`,
				},
				resource.Attribute{
					Name:        "sku",
					Description: `The pricing related information of current CDN profile.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the CDN Profile.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_client_config",
			Category:         "Base",
			ShortDescription: `Gets information about the configuration of the azurerm provider.`,
			Description: `

Use this data source to access the configuration of the AzureRM provider.

`,
			Keywords: []string{
				"base",
				"client",
				"config",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the client config.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the client config.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_cognitive_account",
			Category:         "Cognitive Services",
			ShortDescription: `Gets information about an existing Cognitive Services Account.`,
			Description: `

Use this data source to access information about an existing Cognitive Services Account.

`,
			Keywords: []string{
				"cognitive",
				"services",
				"account",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the Cognitive Services Account.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) Specifies the name of the resource group where the Cognitive Services Account resides. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the Cognitive Services Account exists`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `The kind of the Cognitive Services Account`,
				},
				resource.Attribute{
					Name:        "sku_name",
					Description: `The sku name of the Cognitive Services Account`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `The endpoint of the Cognitive Services Account`,
				},
				resource.Attribute{
					Name:        "qna_runtime_endpoint",
					Description: `If ` + "`" + `kind` + "`" + ` is ` + "`" + `QnAMaker` + "`" + ` the link to the QNA runtime.`,
				},
				resource.Attribute{
					Name:        "primary_access_key",
					Description: `The primary access key of the Cognitive Services Account`,
				},
				resource.Attribute{
					Name:        "secondary_access_key",
					Description: `The secondary access key of the Cognitive Services Account`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assigned to the resource. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Cognitive Services Account.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the Cognitive Services Account exists`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `The kind of the Cognitive Services Account`,
				},
				resource.Attribute{
					Name:        "sku_name",
					Description: `The sku name of the Cognitive Services Account`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `The endpoint of the Cognitive Services Account`,
				},
				resource.Attribute{
					Name:        "qna_runtime_endpoint",
					Description: `If ` + "`" + `kind` + "`" + ` is ` + "`" + `QnAMaker` + "`" + ` the link to the QNA runtime.`,
				},
				resource.Attribute{
					Name:        "primary_access_key",
					Description: `The primary access key of the Cognitive Services Account`,
				},
				resource.Attribute{
					Name:        "secondary_access_key",
					Description: `The secondary access key of the Cognitive Services Account`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assigned to the resource. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Cognitive Services Account.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_container_registry",
			Category:         "Container",
			ShortDescription: `Get information about an existing Container Registry`,
			Description: `

Use this data source to access information about an existing Container Registry.

`,
			Icon: "Containers/10105-icon-service-Container-Registries.svg",
			Keywords: []string{
				"container",
				"registry",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Container Registry.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The Name of the Resource Group where this Container Registry exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Container Registry ID.`,
				},
				resource.Attribute{
					Name:        "login_server",
					Description: `The URL that can be used to log into the container registry.`,
				},
				resource.Attribute{
					Name:        "admin_username",
					Description: `The Username associated with the Container Registry Admin account - if the admin account is enabled.`,
				},
				resource.Attribute{
					Name:        "admin_password",
					Description: `The Password associated with the Container Registry Admin account - if the admin account is enabled.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure Region in which this Container Registry exists.`,
				},
				resource.Attribute{
					Name:        "admin_enabled",
					Description: `Is the Administrator account enabled for this Container Registry.`,
				},
				resource.Attribute{
					Name:        "sku",
					Description: `The SKU of this Container Registry, such as ` + "`" + `Basic` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "storage_account_id",
					Description: `The ID of the Storage Account used for this Container Registry. This is only returned for ` + "`" + `Classic` + "`" + ` SKU's.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A map of tags assigned to the Container Registry. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Container Registry.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Container Registry ID.`,
				},
				resource.Attribute{
					Name:        "login_server",
					Description: `The URL that can be used to log into the container registry.`,
				},
				resource.Attribute{
					Name:        "admin_username",
					Description: `The Username associated with the Container Registry Admin account - if the admin account is enabled.`,
				},
				resource.Attribute{
					Name:        "admin_password",
					Description: `The Password associated with the Container Registry Admin account - if the admin account is enabled.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure Region in which this Container Registry exists.`,
				},
				resource.Attribute{
					Name:        "admin_enabled",
					Description: `Is the Administrator account enabled for this Container Registry.`,
				},
				resource.Attribute{
					Name:        "sku",
					Description: `The SKU of this Container Registry, such as ` + "`" + `Basic` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "storage_account_id",
					Description: `The ID of the Storage Account used for this Container Registry. This is only returned for ` + "`" + `Classic` + "`" + ` SKU's.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A map of tags assigned to the Container Registry. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Container Registry.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_cosmosdb_account",
			Category:         "CosmosDB (DocumentDB)",
			ShortDescription: `Gets information about an existing CosmosDB (formally DocumentDB) Account.`,
			Description: `

Use this data source to access information about an existing CosmosDB (formally DocumentDB) Account.

`,
			Keywords: []string{
				"cosmosdb",
				"documentdb",
				"account",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name of the CosmosDB Account.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `Specifies the name of the resource group in which the CosmosDB Account resides. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the CosmosDB Account.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the resource exists.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource.`,
				},
				resource.Attribute{
					Name:        "offer_type",
					Description: `The Offer Type to used by this CosmosDB Account.`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `The Kind of the CosmosDB account.`,
				},
				resource.Attribute{
					Name:        "key_vault_key_id",
					Description: `The Key Vault key URI for CMK encryption. ~>`,
				},
				resource.Attribute{
					Name:        "ip_range_filter",
					Description: `The current IP Filter for this CosmosDB account`,
				},
				resource.Attribute{
					Name:        "enable_free_tier",
					Description: `If Free Tier pricing option is enabled for this CosmosDB Account.`,
				},
				resource.Attribute{
					Name:        "enable_automatic_failover",
					Description: `If automatic failover is enabled for this CosmosDB Account.`,
				},
				resource.Attribute{
					Name:        "capabilities",
					Description: `Capabilities enabled on this Cosmos DB account.`,
				},
				resource.Attribute{
					Name:        "is_virtual_network_filter_enabled",
					Description: `If virtual network filtering is enabled for this Cosmos DB account.`,
				},
				resource.Attribute{
					Name:        "virtual_network_rule",
					Description: `Subnets that are allowed to access this CosmosDB account.`,
				},
				resource.Attribute{
					Name:        "enable_multiple_write_locations",
					Description: `If multi-master is enabled for this Cosmos DB account. ` + "`" + `consistency_policy` + "`" + ` The current consistency Settings for this CosmosDB account with the following properties:`,
				},
				resource.Attribute{
					Name:        "consistency_level",
					Description: `The Consistency Level used by this CosmosDB Account.`,
				},
				resource.Attribute{
					Name:        "max_interval_in_seconds",
					Description: `The amount of staleness (in seconds) tolerated when the consistency level is Bounded Staleness.`,
				},
				resource.Attribute{
					Name:        "max_staleness_prefix",
					Description: `The number of stale requests tolerated when the consistency level is Bounded Staleness. ` + "`" + `geo_location` + "`" + ` The geographic locations data is replicated to with the following properties:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the location.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The name of the Azure region hosting replicated data.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The locations fail over priority. ` + "`" + `virtual_network_rule` + "`" + ` The virtual network subnets allowed to access this Cosmos DB account with the following properties:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the virtual network subnet.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `The endpoint used to connect to the CosmosDB account.`,
				},
				resource.Attribute{
					Name:        "read_endpoints",
					Description: `A list of read endpoints available for this CosmosDB account.`,
				},
				resource.Attribute{
					Name:        "write_endpoints",
					Description: `A list of write endpoints available for this CosmosDB account.`,
				},
				resource.Attribute{
					Name:        "primary_key",
					Description: `The Primary master key for the CosmosDB Account.`,
				},
				resource.Attribute{
					Name:        "secondary_key",
					Description: `The Secondary master key for the CosmosDB Account.`,
				},
				resource.Attribute{
					Name:        "primary_readonly_key",
					Description: `The Primary read-only master Key for the CosmosDB Account.`,
				},
				resource.Attribute{
					Name:        "secondary_readonly_key",
					Description: `The Secondary read-only master key for the CosmosDB Account. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the CosmosDB Account.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the CosmosDB Account.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the resource exists.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource.`,
				},
				resource.Attribute{
					Name:        "offer_type",
					Description: `The Offer Type to used by this CosmosDB Account.`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `The Kind of the CosmosDB account.`,
				},
				resource.Attribute{
					Name:        "key_vault_key_id",
					Description: `The Key Vault key URI for CMK encryption. ~>`,
				},
				resource.Attribute{
					Name:        "ip_range_filter",
					Description: `The current IP Filter for this CosmosDB account`,
				},
				resource.Attribute{
					Name:        "enable_free_tier",
					Description: `If Free Tier pricing option is enabled for this CosmosDB Account.`,
				},
				resource.Attribute{
					Name:        "enable_automatic_failover",
					Description: `If automatic failover is enabled for this CosmosDB Account.`,
				},
				resource.Attribute{
					Name:        "capabilities",
					Description: `Capabilities enabled on this Cosmos DB account.`,
				},
				resource.Attribute{
					Name:        "is_virtual_network_filter_enabled",
					Description: `If virtual network filtering is enabled for this Cosmos DB account.`,
				},
				resource.Attribute{
					Name:        "virtual_network_rule",
					Description: `Subnets that are allowed to access this CosmosDB account.`,
				},
				resource.Attribute{
					Name:        "enable_multiple_write_locations",
					Description: `If multi-master is enabled for this Cosmos DB account. ` + "`" + `consistency_policy` + "`" + ` The current consistency Settings for this CosmosDB account with the following properties:`,
				},
				resource.Attribute{
					Name:        "consistency_level",
					Description: `The Consistency Level used by this CosmosDB Account.`,
				},
				resource.Attribute{
					Name:        "max_interval_in_seconds",
					Description: `The amount of staleness (in seconds) tolerated when the consistency level is Bounded Staleness.`,
				},
				resource.Attribute{
					Name:        "max_staleness_prefix",
					Description: `The number of stale requests tolerated when the consistency level is Bounded Staleness. ` + "`" + `geo_location` + "`" + ` The geographic locations data is replicated to with the following properties:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the location.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The name of the Azure region hosting replicated data.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The locations fail over priority. ` + "`" + `virtual_network_rule` + "`" + ` The virtual network subnets allowed to access this Cosmos DB account with the following properties:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the virtual network subnet.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `The endpoint used to connect to the CosmosDB account.`,
				},
				resource.Attribute{
					Name:        "read_endpoints",
					Description: `A list of read endpoints available for this CosmosDB account.`,
				},
				resource.Attribute{
					Name:        "write_endpoints",
					Description: `A list of write endpoints available for this CosmosDB account.`,
				},
				resource.Attribute{
					Name:        "primary_key",
					Description: `The Primary master key for the CosmosDB Account.`,
				},
				resource.Attribute{
					Name:        "secondary_key",
					Description: `The Secondary master key for the CosmosDB Account.`,
				},
				resource.Attribute{
					Name:        "primary_readonly_key",
					Description: `The Primary read-only master Key for the CosmosDB Account.`,
				},
				resource.Attribute{
					Name:        "secondary_readonly_key",
					Description: `The Secondary read-only master key for the CosmosDB Account. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the CosmosDB Account.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_data_factory",
			Category:         "Data Factory",
			ShortDescription: `Gets information about an existing Azure Data Factory (Version 2).`,
			Description: `

Use this data source to access information about an existing Azure Data Factory (Version 2).

`,
			Icon: "Integration/00041-icon-service-Data-Factory.svg",
			Keywords: []string{
				"data",
				"factory",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_data_lake_store",
			Category:         "Data Lake",
			ShortDescription: `Gets information about an existing Data Lake Store`,
			Description: `

Use this data source to access information about an existing Data Lake Store.

`,
			Keywords: []string{
				"data",
				"lake",
				"store",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Data Lake Store.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The Name of the Resource Group where the Data Lake Store exists. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Data Lake Store.`,
				},
				resource.Attribute{
					Name:        "encryption_state",
					Description: `the Encryption State of this Data Lake Store Account, such as ` + "`" + `Enabled` + "`" + ` or ` + "`" + `Disabled` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "encryption_type",
					Description: `the Encryption Type used for this Data Lake Store Account.`,
				},
				resource.Attribute{
					Name:        "firewall_allow_azure_ips",
					Description: `are Azure Service IP's allowed through the firewall?`,
				},
				resource.Attribute{
					Name:        "firewall_state",
					Description: `the state of the firewall, such as ` + "`" + `Enabled` + "`" + ` or ` + "`" + `Disabled` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tier",
					Description: `Current monthly commitment tier for the account.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the Data Lake Store. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Data Lake Store.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Data Lake Store.`,
				},
				resource.Attribute{
					Name:        "encryption_state",
					Description: `the Encryption State of this Data Lake Store Account, such as ` + "`" + `Enabled` + "`" + ` or ` + "`" + `Disabled` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "encryption_type",
					Description: `the Encryption Type used for this Data Lake Store Account.`,
				},
				resource.Attribute{
					Name:        "firewall_allow_azure_ips",
					Description: `are Azure Service IP's allowed through the firewall?`,
				},
				resource.Attribute{
					Name:        "firewall_state",
					Description: `the state of the firewall, such as ` + "`" + `Enabled` + "`" + ` or ` + "`" + `Disabled` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tier",
					Description: `Current monthly commitment tier for the account.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the Data Lake Store. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Data Lake Store.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_data_share",
			Category:         "Data Share",
			ShortDescription: `Gets information about an existing Data Share.`,
			Description: `

Use this data source to access information about an existing Data Share.

`,
			Icon: "Storage/10098-icon-service-Data-Shares.svg",
			Keywords: []string{
				"data",
				"share",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Data Share.`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `The kind of the Data Share.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the Data Share.`,
				},
				resource.Attribute{
					Name:        "snapshot_schedule",
					Description: `A ` + "`" + `snapshot_schedule` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "terms",
					Description: `The terms of the Data Share. --- A ` + "`" + `snapshot_schedule` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the snapshot schedule.`,
				},
				resource.Attribute{
					Name:        "recurrence",
					Description: `The interval of the synchronization with the source data.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `The synchronization with the source data's start time. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Data Share.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_data_share_account",
			Category:         "Data Share",
			ShortDescription: `Gets information about an existing Data Share Account.`,
			Description: `

Use this data source to access information about an existing Data Share Account.

`,
			Keywords: []string{
				"data",
				"share",
				"account",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Data Share Account.`,
				},
				resource.Attribute{
					Name:        "identity",
					Description: `An ` + "`" + `identity` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the Data Share Account. --- An ` + "`" + `identity` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "principal_id",
					Description: `The ID of the Principal (Client) in Azure Active Directory.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `The ID of the Azure Active Directory Tenant.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The identity type of the Data Share Account. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Data Share Account.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_data_share_dataset_blob_storage",
			Category:         "Data Share",
			ShortDescription: `Gets information about an existing Data Share Blob Storage Dataset.`,
			Description: `

Use this data source to access information about an existing Data Share Blob Storage Dataset.

`,
			Keywords: []string{
				"data",
				"share",
				"dataset",
				"blob",
				"storage",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Data Share Blob Storage Dataset.`,
				},
				resource.Attribute{
					Name:        "container_name",
					Description: `The name of the storage account container to be shared with the receiver.`,
				},
				resource.Attribute{
					Name:        "storage_account",
					Description: `A ` + "`" + `storage_account` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "file_path",
					Description: `The path of the file in the storage container to be shared with the receiver.`,
				},
				resource.Attribute{
					Name:        "folder_path",
					Description: `The folder path of the file in the storage container to be shared with the receiver.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The name of the Data Share Dataset. --- A ` + "`" + `storage_account` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the storage account to be shared with the receiver.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The resource group name of the storage account to be shared with the receiver.`,
				},
				resource.Attribute{
					Name:        "subscription_id",
					Description: `The subscription id of the storage account to be shared with the receiver. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Data Share Blob Storage Dataset.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_data_share_dataset_data_lake_gen1",
			Category:         "Data Share",
			ShortDescription: `Gets information about an existing DataShareDataLakeGen1Dataset.`,
			Description: `

Use this data source to access information about an existing DataShareDataLakeGen1Dataset.

`,
			Keywords: []string{
				"data",
				"share",
				"dataset",
				"lake",
				"gen1",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID of the Data Share Data Lake Gen1 Dataset.`,
				},
				resource.Attribute{
					Name:        "data_lake_store_id",
					Description: `The resource ID of the Data Lake Store to be shared with the receiver.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The displayed name of the Data Share Dataset.`,
				},
				resource.Attribute{
					Name:        "file_name",
					Description: `The file name of the data lake store to be shared with the receiver.`,
				},
				resource.Attribute{
					Name:        "folder_path",
					Description: `The folder path of the data lake store to be shared with the receiver. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the DataShareDataLakeGen1Dataset.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_data_share_dataset_data_lake_gen2",
			Category:         "Data Share",
			ShortDescription: `Gets information about an existing Data Share Data Lake Gen2 Dataset.`,
			Description: `

Use this data source to access information about an existing Data Share Data Lake Gen2 Dataset.

`,
			Keywords: []string{
				"data",
				"share",
				"dataset",
				"lake",
				"gen2",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID of the Data Share Data Lake Gen2 Dataset.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The name of the Data Share Dataset.`,
				},
				resource.Attribute{
					Name:        "file_path",
					Description: `The path of the file in the data lake file system to be shared with the receiver.`,
				},
				resource.Attribute{
					Name:        "file_system_name",
					Description: `The name of the data lake file system to be shared with the receiver.`,
				},
				resource.Attribute{
					Name:        "folder_path",
					Description: `The folder path in the data lake file system to be shared with the receiver.`,
				},
				resource.Attribute{
					Name:        "storage_account_id",
					Description: `The resource ID of the storage account of the data lake file system to be shared with the receiver. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Data Share Data Lake Gen2 Dataset.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_data_share_dataset_kusto_cluster",
			Category:         "Data Share",
			ShortDescription: `Gets information about an existing Data Share Kusto Cluster Dataset.`,
			Description: `

Use this data source to access information about an existing Data Share Kusto Cluster Dataset.

`,
			Keywords: []string{
				"data",
				"share",
				"dataset",
				"kusto",
				"cluster",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID of the Data Share Kusto Cluster Dataset.`,
				},
				resource.Attribute{
					Name:        "kusto_cluster_id",
					Description: `The resource ID of the Kusto Cluster to be shared with the receiver.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The name of the Data Share Dataset.`,
				},
				resource.Attribute{
					Name:        "kusto_cluster_location",
					Description: `The location of the Kusto Cluster. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Data Share Kusto Cluster Dataset.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_data_share_dataset_kusto_database",
			Category:         "Data Share",
			ShortDescription: `Gets information about an existing Data Share Kusto Database Dataset.`,
			Description: `

Use this data source to access information about an existing Data Share Kusto Database Dataset.

`,
			Keywords: []string{
				"data",
				"share",
				"dataset",
				"kusto",
				"database",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID of the Data Share Kusto Database Dataset.`,
				},
				resource.Attribute{
					Name:        "kusto_database_id",
					Description: `The resource ID of the Kusto Cluster Database to be shared with the receiver.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The name of the Data Share Dataset.`,
				},
				resource.Attribute{
					Name:        "kusto_cluster_location",
					Description: `The location of the Kusto Cluster. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Data Share Kusto Database Dataset.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_database_migration_project",
			Category:         "Database Migration",
			ShortDescription: `Gets information about an existing Database Migration Project`,
			Description: `

Use this data source to access information about an existing Database Migration Project.

`,
			Keywords: []string{
				"database",
				"migration",
				"project",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the database migration project.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) Name of the resource group where resource belongs to.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) Name of the database migration service where resource belongs to. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Database Migration Project.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Azure location where the resource exists.`,
				},
				resource.Attribute{
					Name:        "source_platform",
					Description: `The platform type of the migration source.`,
				},
				resource.Attribute{
					Name:        "target_platform",
					Description: `The platform type of the migration target.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assigned to the resource. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the API Management API.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Database Migration Project.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Azure location where the resource exists.`,
				},
				resource.Attribute{
					Name:        "source_platform",
					Description: `The platform type of the migration source.`,
				},
				resource.Attribute{
					Name:        "target_platform",
					Description: `The platform type of the migration target.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assigned to the resource. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the API Management API.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_database_migration_service",
			Category:         "Database Migration",
			ShortDescription: `Gets information about an existing Database Migration Service`,
			Description: `

Use this data source to access information about an existing Database Migration Service.


`,
			Keywords: []string{
				"database",
				"migration",
				"service",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specify the name of the database migration service.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) Specifies the Name of the Resource Group within which the database migration service exists ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Database Migration Service.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Azure location where the resource exists.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the virtual subnet resource to which the database migration service exists.`,
				},
				resource.Attribute{
					Name:        "sku_name",
					Description: `The sku name of database migration service.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assigned to the resource. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the API Management API.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Database Migration Service.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Azure location where the resource exists.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the virtual subnet resource to which the database migration service exists.`,
				},
				resource.Attribute{
					Name:        "sku_name",
					Description: `The sku name of database migration service.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assigned to the resource. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the API Management API.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_databricks_workspace",
			Category:         "Databricks",
			ShortDescription: `Gets information on an existing Databricks Workspace`,
			Description: `

Use this data source to access information about an existing Databricks workspace.

`,
			Keywords: []string{
				"databricks",
				"workspace",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Databricks Workspace.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The Name of the Resource Group where the Databricks Workspace exists. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Databricks Workspace.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the Databricks Workspace exists.`,
				},
				resource.Attribute{
					Name:        "sku",
					Description: `SKU of this Databricks Workspace.`,
				},
				resource.Attribute{
					Name:        "workspace_id",
					Description: `Unique ID of this Databricks Workspace in Databricks management plane.`,
				},
				resource.Attribute{
					Name:        "workspace_url",
					Description: `URL this Databricks Workspace is accessible on.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the Databricks Workspace. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Databricks Workspace.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Databricks Workspace.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the Databricks Workspace exists.`,
				},
				resource.Attribute{
					Name:        "sku",
					Description: `SKU of this Databricks Workspace.`,
				},
				resource.Attribute{
					Name:        "workspace_id",
					Description: `Unique ID of this Databricks Workspace in Databricks management plane.`,
				},
				resource.Attribute{
					Name:        "workspace_url",
					Description: `URL this Databricks Workspace is accessible on.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the Databricks Workspace. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Databricks Workspace.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_dedicated_host",
			Category:         "Compute",
			ShortDescription: `Gets information about an existing Dedicated Host`,
			Description: `

Use this data source to access information about an existing Dedicated Host.

`,
			Keywords: []string{
				"compute",
				"dedicated",
				"host",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name of the Dedicated Host.`,
				},
				resource.Attribute{
					Name:        "dedicated_host_group_name",
					Description: `Specifies the name of the Dedicated Host Group the Dedicated Host is located in.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `Specifies the name of the resource group the Dedicated Host is located in. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Dedicated Host.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location where the Dedicated Host exists.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the Dedicated Host. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Dedicated Host.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Dedicated Host.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location where the Dedicated Host exists.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the Dedicated Host. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Dedicated Host.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_dedicated_host_group",
			Category:         "Compute",
			ShortDescription: `Gets information about an existing Dedicated Host Group`,
			Description: `

Use this data source to access information about an existing Dedicated Host Group.

`,
			Keywords: []string{
				"compute",
				"dedicated",
				"host",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name of the Dedicated Host Group.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `Specifies the name of the resource group the Dedicated Host Group is located in. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Dedicated Host Group.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the Dedicated Host Group exists.`,
				},
				resource.Attribute{
					Name:        "platform_fault_domain_count",
					Description: `The number of fault domains that the Dedicated Host Group spans.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `The Availability Zones in which this Dedicated Host Group is located.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Dedicated Host Group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Dedicated Host Group.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the Dedicated Host Group exists.`,
				},
				resource.Attribute{
					Name:        "platform_fault_domain_count",
					Description: `The number of fault domains that the Dedicated Host Group spans.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `The Availability Zones in which this Dedicated Host Group is located.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Dedicated Host Group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_dev_test_lab",
			Category:         "Dev Test",
			ShortDescription: `Gets information about an existing Dev Test Lab.`,
			Description: `

Use this data source to access information about an existing Dev Test Lab.

`,
			Icon: "DevOps/10264-icon-service-DevTest-Labs.svg",
			Keywords: []string{
				"dev",
				"test",
				"lab",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Dev Test Lab.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The Name of the Resource Group where the Dev Test Lab exists. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Dev Test Lab.`,
				},
				resource.Attribute{
					Name:        "artifacts_storage_account_id",
					Description: `The ID of the Storage Account used for Artifact Storage.`,
				},
				resource.Attribute{
					Name:        "default_storage_account_id",
					Description: `The ID of the Default Storage Account for this Dev Test Lab.`,
				},
				resource.Attribute{
					Name:        "default_premium_storage_account_id",
					Description: `The ID of the Default Premium Storage Account for this Dev Test Lab.`,
				},
				resource.Attribute{
					Name:        "key_vault_id",
					Description: `The ID of the Key used for this Dev Test Lab.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the Dev Test Lab exists.`,
				},
				resource.Attribute{
					Name:        "premium_data_disk_storage_account_id",
					Description: `The ID of the Storage Account used for Storage of Premium Data Disk.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `The type of storage used by the Dev Test Lab.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "unique_identifier",
					Description: `The unique immutable identifier of the Dev Test Lab. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Dev Test Lab.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Dev Test Lab.`,
				},
				resource.Attribute{
					Name:        "artifacts_storage_account_id",
					Description: `The ID of the Storage Account used for Artifact Storage.`,
				},
				resource.Attribute{
					Name:        "default_storage_account_id",
					Description: `The ID of the Default Storage Account for this Dev Test Lab.`,
				},
				resource.Attribute{
					Name:        "default_premium_storage_account_id",
					Description: `The ID of the Default Premium Storage Account for this Dev Test Lab.`,
				},
				resource.Attribute{
					Name:        "key_vault_id",
					Description: `The ID of the Key used for this Dev Test Lab.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the Dev Test Lab exists.`,
				},
				resource.Attribute{
					Name:        "premium_data_disk_storage_account_id",
					Description: `The ID of the Storage Account used for Storage of Premium Data Disk.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `The type of storage used by the Dev Test Lab.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "unique_identifier",
					Description: `The unique immutable identifier of the Dev Test Lab. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Dev Test Lab.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_dev_test_virtual_network",
			Category:         "Dev Test",
			ShortDescription: `Gets information about an existing Dev Test Lab Virtual Network.`,
			Description: `

Use this data source to access information about an existing Dev Test Lab Virtual Network.

`,
			Keywords: []string{
				"dev",
				"test",
				"virtual",
				"network",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name of the Virtual Network.`,
				},
				resource.Attribute{
					Name:        "lab_name",
					Description: `Specifies the name of the Dev Test Lab.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `Specifies the name of the resource group that contains the Virtual Network. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "allowed_subnets",
					Description: `The list of subnets enabled for the virtual network as defined below.`,
				},
				resource.Attribute{
					Name:        "subnet_overrides",
					Description: `The list of permission overrides for the subnets as defined below.`,
				},
				resource.Attribute{
					Name:        "unique_identifier",
					Description: `The unique immutable identifier of the virtual network. --- An ` + "`" + `allowed_subnets` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "allow_public_ip",
					Description: `Indicates if this subnet allows public IP addresses. Possible values are ` + "`" + `Allow` + "`" + `, ` + "`" + `Default` + "`" + ` and ` + "`" + `Deny` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lab_subnet_name",
					Description: `The name of the subnet.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `The resource identifier for the subnet. --- An ` + "`" + `subnets_override` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "lab_subnet_name",
					Description: `The name of the subnet.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `The resource identifier for the subnet.`,
				},
				resource.Attribute{
					Name:        "use_in_vm_creation_permission",
					Description: `Indicates if the subnet can be used for VM creation. Possible values are ` + "`" + `Allow` + "`" + `, ` + "`" + `Default` + "`" + ` and ` + "`" + `Deny` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "use_public_ip_permission",
					Description: `Indicates if the subnet can be assigned public IP addresses. Possible values are ` + "`" + `Allow` + "`" + `, ` + "`" + `Default` + "`" + ` and ` + "`" + `Deny` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "virtual_network_pool_name",
					Description: `The virtual network pool associated with this subnet. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Dev Test Lab Virtual Network.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "allowed_subnets",
					Description: `The list of subnets enabled for the virtual network as defined below.`,
				},
				resource.Attribute{
					Name:        "subnet_overrides",
					Description: `The list of permission overrides for the subnets as defined below.`,
				},
				resource.Attribute{
					Name:        "unique_identifier",
					Description: `The unique immutable identifier of the virtual network. --- An ` + "`" + `allowed_subnets` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "allow_public_ip",
					Description: `Indicates if this subnet allows public IP addresses. Possible values are ` + "`" + `Allow` + "`" + `, ` + "`" + `Default` + "`" + ` and ` + "`" + `Deny` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lab_subnet_name",
					Description: `The name of the subnet.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `The resource identifier for the subnet. --- An ` + "`" + `subnets_override` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "lab_subnet_name",
					Description: `The name of the subnet.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `The resource identifier for the subnet.`,
				},
				resource.Attribute{
					Name:        "use_in_vm_creation_permission",
					Description: `Indicates if the subnet can be used for VM creation. Possible values are ` + "`" + `Allow` + "`" + `, ` + "`" + `Default` + "`" + ` and ` + "`" + `Deny` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "use_public_ip_permission",
					Description: `Indicates if the subnet can be assigned public IP addresses. Possible values are ` + "`" + `Allow` + "`" + `, ` + "`" + `Default` + "`" + ` and ` + "`" + `Deny` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "virtual_network_pool_name",
					Description: `The virtual network pool associated with this subnet. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Dev Test Lab Virtual Network.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_digital_twins_instance",
			Category:         "Digital Twins",
			ShortDescription: `Gets information about an existing Digital Twins instance.`,
			Description: `

Use this data source to access information about an existing Digital Twins instance.

`,
			Keywords: []string{
				"digital",
				"twins",
				"instance",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Digital Twins instance.`,
				},
				resource.Attribute{
					Name:        "host_name",
					Description: `The Api endpoint to work with this Digital Twins instance.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure Region where the Digital Twins instance exists.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the Digital Twins instance. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Digital Twins instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_disk_access",
			Category:         "Compute",
			ShortDescription: `Gets information about an existing Disk Access.`,
			Description: `

Use this data source to access information about an existing Disk Access.

`,
			Keywords: []string{
				"compute",
				"disk",
				"access",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Disk Access. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Disk.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_disk_encryption_set",
			Category:         "Compute",
			ShortDescription: `Gets information about an existing Disk Encryption Set`,
			Description: `

Use this data source to access information about an existing Disk Encryption Set.

`,
			Icon: "Other/00398-icon-service-Disk-Encryption-Sets.svg",
			Keywords: []string{
				"compute",
				"disk",
				"encryption",
				"set",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Disk Encryption Set exists.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of the Resource Group where the Disk Encryption Set exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Disk Encryption Set.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location where the Disk Encryption Set exists.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the Disk Encryption Set. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Disk Encryption Set.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Disk Encryption Set.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location where the Disk Encryption Set exists.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the Disk Encryption Set. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Disk Encryption Set.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_dns_zone",
			Category:         "DNS",
			ShortDescription: `Gets information about an existing DNS Zone.`,
			Description: `

Use this data source to access information about an existing DNS Zone.

`,
			Icon: "Networking/10064-icon-service-DNS-Zones.svg",
			Keywords: []string{
				"dns",
				"zone",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the DNS Zone.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Optional) The Name of the Resource Group where the DNS Zone exists. If the Name of the Resource Group is not provided, the first DNS Zone from the list of DNS Zones in your subscription that matches ` + "`" + `name` + "`" + ` will be returned. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the DNS Zone.`,
				},
				resource.Attribute{
					Name:        "max_number_of_record_sets",
					Description: `Maximum number of Records in the zone.`,
				},
				resource.Attribute{
					Name:        "number_of_record_sets",
					Description: `The number of records already in the zone.`,
				},
				resource.Attribute{
					Name:        "name_servers",
					Description: `A list of values that make up the NS record for the zone.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the EventHub Namespace. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the DNS Zone.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the DNS Zone.`,
				},
				resource.Attribute{
					Name:        "max_number_of_record_sets",
					Description: `Maximum number of Records in the zone.`,
				},
				resource.Attribute{
					Name:        "number_of_record_sets",
					Description: `The number of records already in the zone.`,
				},
				resource.Attribute{
					Name:        "name_servers",
					Description: `A list of values that make up the NS record for the zone.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the EventHub Namespace. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the DNS Zone.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_eventgrid_domain_topic",
			Category:         "Messaging",
			ShortDescription: `Gets information about an existing EventGrid Domain Topic`,
			Description: `

Use this data source to access information about an existing EventGrid Domain Topic

`,
			Keywords: []string{
				"messaging",
				"eventgrid",
				"domain",
				"topic",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the EventGrid Domain Topic resource.`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `The name of the EventGrid Domain Topic domain.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of the resource group in which the EventGrid Domain Topic exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The EventGrid Domain Topic ID.`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `The EventGrid Domain Topic Domain name. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the EventGrid Domain Topic.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The EventGrid Domain Topic ID.`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `The EventGrid Domain Topic Domain name. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the EventGrid Domain Topic.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_eventgrid_topic",
			Category:         "Messaging",
			ShortDescription: `Gets information about an existing EventGrid Topic`,
			Description: `

Use this data source to access information about an existing EventGrid Topic

`,
			Icon: "Integration/10206-icon-service-Event-Grid-Topics.svg",
			Keywords: []string{
				"messaging",
				"eventgrid",
				"topic",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the EventGrid Topic resource.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of the resource group in which the EventGrid Topic exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The EventGrid Topic ID.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `The Endpoint associated with the EventGrid Topic.`,
				},
				resource.Attribute{
					Name:        "primary_access_key",
					Description: `The Primary Shared Access Key associated with the EventGrid Topic.`,
				},
				resource.Attribute{
					Name:        "secondary_access_key",
					Description: `The Secondary Shared Access Key associated with the EventGrid Topic. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the EventGrid Topic.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The EventGrid Topic ID.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `The Endpoint associated with the EventGrid Topic.`,
				},
				resource.Attribute{
					Name:        "primary_access_key",
					Description: `The Primary Shared Access Key associated with the EventGrid Topic.`,
				},
				resource.Attribute{
					Name:        "secondary_access_key",
					Description: `The Secondary Shared Access Key associated with the EventGrid Topic. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the EventGrid Topic.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_eventhub",
			Category:         "Messaging",
			ShortDescription: `Gets information about an existing EventHub.`,
			Description: `

Use this data source to access information about an existing EventHub.

`,
			Icon: "Analytics/00039-icon-service-Event-Hubs.svg",
			Keywords: []string{
				"messaging",
				"eventhub",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the EventHub.`,
				},
				resource.Attribute{
					Name:        "partition_count",
					Description: `The number of partitions in the EventHub.`,
				},
				resource.Attribute{
					Name:        "partition_ids",
					Description: `The identifiers for the partitions of this EventHub. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the EventHub.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_eventhub_authorization_rule",
			Category:         "Messaging",
			ShortDescription: `Gets information about an Event Hubs Authorization Rule within an Event Hub.`,
			Description:      ``,
			Keywords: []string{
				"messaging",
				"eventhub",
				"authorization",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name of the EventHub Authorization Rule resource. be created.`,
				},
				resource.Attribute{
					Name:        "namespace_name",
					Description: `Specifies the name of the grandparent EventHub Namespace.`,
				},
				resource.Attribute{
					Name:        "eventhub_name",
					Description: `Specifies the name of the EventHub.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of the resource group in which the EventHub Authorization Rule's grandparent Namespace exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The EventHub ID.`,
				},
				resource.Attribute{
					Name:        "primary_connection_string_alias",
					Description: `The alias of the Primary Connection String for the Event Hubs Authorization Rule.`,
				},
				resource.Attribute{
					Name:        "secondary_connection_string_alias",
					Description: `The alias of the Secondary Connection String for the Event Hubs Authorization Rule.`,
				},
				resource.Attribute{
					Name:        "primary_connection_string",
					Description: `The Primary Connection String for the Event Hubs Authorization Rule.`,
				},
				resource.Attribute{
					Name:        "primary_key",
					Description: `The Primary Key for the Event Hubs Authorization Rule.`,
				},
				resource.Attribute{
					Name:        "secondary_connection_string",
					Description: `The Secondary Connection String for the Event Hubs Authorization Rule.`,
				},
				resource.Attribute{
					Name:        "secondary_key",
					Description: `The Secondary Key for the Event Hubs Authorization Rule. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the EventHub Authorization Rule.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The EventHub ID.`,
				},
				resource.Attribute{
					Name:        "primary_connection_string_alias",
					Description: `The alias of the Primary Connection String for the Event Hubs Authorization Rule.`,
				},
				resource.Attribute{
					Name:        "secondary_connection_string_alias",
					Description: `The alias of the Secondary Connection String for the Event Hubs Authorization Rule.`,
				},
				resource.Attribute{
					Name:        "primary_connection_string",
					Description: `The Primary Connection String for the Event Hubs Authorization Rule.`,
				},
				resource.Attribute{
					Name:        "primary_key",
					Description: `The Primary Key for the Event Hubs Authorization Rule.`,
				},
				resource.Attribute{
					Name:        "secondary_connection_string",
					Description: `The Secondary Connection String for the Event Hubs Authorization Rule.`,
				},
				resource.Attribute{
					Name:        "secondary_key",
					Description: `The Secondary Key for the Event Hubs Authorization Rule. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the EventHub Authorization Rule.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_eventhub_consumer_group",
			Category:         "Messaging",
			ShortDescription: `Gets information about an Event Hubs Consumer Group within an Event Hub.`,
			Description: `

Use this data source to access information about an existing Event Hubs Consumer Group within an Event Hub.

`,
			Keywords: []string{
				"messaging",
				"eventhub",
				"consumer",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name of the EventHub Consumer Group resource.`,
				},
				resource.Attribute{
					Name:        "namespace_name",
					Description: `Specifies the name of the grandparent EventHub Namespace.`,
				},
				resource.Attribute{
					Name:        "eventhub_name",
					Description: `Specifies the name of the EventHub.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of the resource group in which the EventHub Consumer Group's grandparent Namespace exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The EventHub Consumer Group ID.`,
				},
				resource.Attribute{
					Name:        "user_metadata",
					Description: `Specifies the user metadata. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the EventHub Consumer Group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The EventHub Consumer Group ID.`,
				},
				resource.Attribute{
					Name:        "user_metadata",
					Description: `Specifies the user metadata. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the EventHub Consumer Group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_eventhub_namespace",
			Category:         "Messaging",
			ShortDescription: `Gets information about an existing EventHub Namespace.`,
			Description: `

Use this data source to access information about an existing EventHub Namespace.

`,
			Keywords: []string{
				"messaging",
				"eventhub",
				"namespace",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the EventHub Namespace.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The Name of the Resource Group where the EventHub Namespace exists. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the EventHub Namespace.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the EventHub Namespace exists`,
				},
				resource.Attribute{
					Name:        "sku",
					Description: `Defines which tier to use.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `The Capacity / Throughput Units for a ` + "`" + `Standard` + "`" + ` SKU namespace.`,
				},
				resource.Attribute{
					Name:        "auto_inflate_enabled",
					Description: `Is Auto Inflate enabled for the EventHub Namespace?`,
				},
				resource.Attribute{
					Name:        "maximum_throughput_units",
					Description: `Specifies the maximum number of throughput units when Auto Inflate is Enabled.`,
				},
				resource.Attribute{
					Name:        "zone_redundant",
					Description: `Is this EventHub Namespace deployed across Availability Zones?`,
				},
				resource.Attribute{
					Name:        "dedicated_cluster_id",
					Description: `The ID of the EventHub Dedicated Cluster where this Namespace exists.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the EventHub Namespace. The following attributes are exported only if there is an authorization rule named ` + "`" + `RootManageSharedAccessKey` + "`" + ` which is created automatically by Azure.`,
				},
				resource.Attribute{
					Name:        "default_primary_connection_string",
					Description: `The primary connection string for the authorization rule ` + "`" + `RootManageSharedAccessKey` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "default_primary_connection_string_alias",
					Description: `The alias of the primary connection string for the authorization rule ` + "`" + `RootManageSharedAccessKey` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "default_primary_key",
					Description: `The primary access key for the authorization rule ` + "`" + `RootManageSharedAccessKey` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "default_secondary_connection_string",
					Description: `The secondary connection string for the authorization rule ` + "`" + `RootManageSharedAccessKey` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "default_secondary_connection_string_alias",
					Description: `The alias of the secondary connection string for the authorization rule ` + "`" + `RootManageSharedAccessKey` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "default_secondary_key",
					Description: `The secondary access key for the authorization rule ` + "`" + `RootManageSharedAccessKey` + "`" + `. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the EventHub Namespace.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the EventHub Namespace.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the EventHub Namespace exists`,
				},
				resource.Attribute{
					Name:        "sku",
					Description: `Defines which tier to use.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `The Capacity / Throughput Units for a ` + "`" + `Standard` + "`" + ` SKU namespace.`,
				},
				resource.Attribute{
					Name:        "auto_inflate_enabled",
					Description: `Is Auto Inflate enabled for the EventHub Namespace?`,
				},
				resource.Attribute{
					Name:        "maximum_throughput_units",
					Description: `Specifies the maximum number of throughput units when Auto Inflate is Enabled.`,
				},
				resource.Attribute{
					Name:        "zone_redundant",
					Description: `Is this EventHub Namespace deployed across Availability Zones?`,
				},
				resource.Attribute{
					Name:        "dedicated_cluster_id",
					Description: `The ID of the EventHub Dedicated Cluster where this Namespace exists.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the EventHub Namespace. The following attributes are exported only if there is an authorization rule named ` + "`" + `RootManageSharedAccessKey` + "`" + ` which is created automatically by Azure.`,
				},
				resource.Attribute{
					Name:        "default_primary_connection_string",
					Description: `The primary connection string for the authorization rule ` + "`" + `RootManageSharedAccessKey` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "default_primary_connection_string_alias",
					Description: `The alias of the primary connection string for the authorization rule ` + "`" + `RootManageSharedAccessKey` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "default_primary_key",
					Description: `The primary access key for the authorization rule ` + "`" + `RootManageSharedAccessKey` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "default_secondary_connection_string",
					Description: `The secondary connection string for the authorization rule ` + "`" + `RootManageSharedAccessKey` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "default_secondary_connection_string_alias",
					Description: `The alias of the secondary connection string for the authorization rule ` + "`" + `RootManageSharedAccessKey` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "default_secondary_key",
					Description: `The secondary access key for the authorization rule ` + "`" + `RootManageSharedAccessKey` + "`" + `. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the EventHub Namespace.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_eventhub_namespace_authorization_rule",
			Category:         "Messaging",
			ShortDescription: `Gets information about an existing Authorization Rule for an Event Hub Namespace.`,
			Description: `

Use this data source to access information about an Authorization Rule for an Event Hub Namespace.

`,
			Keywords: []string{
				"messaging",
				"eventhub",
				"namespace",
				"authorization",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the EventHub Authorization Rule resource.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of the resource group in which the EventHub Namespace exists.`,
				},
				resource.Attribute{
					Name:        "namespace_name",
					Description: `Specifies the name of the EventHub Namespace. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The EventHub ID.`,
				},
				resource.Attribute{
					Name:        "primary_connection_string_alias",
					Description: `The alias of the Primary Connection String for the Event Hubs authorization Rule.`,
				},
				resource.Attribute{
					Name:        "secondary_connection_string_alias",
					Description: `The alias of the Secondary Connection String for the Event Hubs authorization Rule.`,
				},
				resource.Attribute{
					Name:        "listen",
					Description: `Does this Authorization Rule have permissions to Listen to the Event Hub?`,
				},
				resource.Attribute{
					Name:        "manage",
					Description: `Does this Authorization Rule have permissions to Manage to the Event Hub?`,
				},
				resource.Attribute{
					Name:        "primary_connection_string",
					Description: `The Primary Connection String for the Event Hubs authorization Rule.`,
				},
				resource.Attribute{
					Name:        "primary_key",
					Description: `The Primary Key for the Event Hubs authorization Rule.`,
				},
				resource.Attribute{
					Name:        "secondary_connection_string",
					Description: `The Secondary Connection String for the Event Hubs authorization Rule.`,
				},
				resource.Attribute{
					Name:        "secondary_key",
					Description: `The Secondary Key for the Event Hubs authorization Rule.`,
				},
				resource.Attribute{
					Name:        "send",
					Description: `Does this Authorization Rule have permissions to Send to the Event Hub? ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Authorization Rule.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The EventHub ID.`,
				},
				resource.Attribute{
					Name:        "primary_connection_string_alias",
					Description: `The alias of the Primary Connection String for the Event Hubs authorization Rule.`,
				},
				resource.Attribute{
					Name:        "secondary_connection_string_alias",
					Description: `The alias of the Secondary Connection String for the Event Hubs authorization Rule.`,
				},
				resource.Attribute{
					Name:        "listen",
					Description: `Does this Authorization Rule have permissions to Listen to the Event Hub?`,
				},
				resource.Attribute{
					Name:        "manage",
					Description: `Does this Authorization Rule have permissions to Manage to the Event Hub?`,
				},
				resource.Attribute{
					Name:        "primary_connection_string",
					Description: `The Primary Connection String for the Event Hubs authorization Rule.`,
				},
				resource.Attribute{
					Name:        "primary_key",
					Description: `The Primary Key for the Event Hubs authorization Rule.`,
				},
				resource.Attribute{
					Name:        "secondary_connection_string",
					Description: `The Secondary Connection String for the Event Hubs authorization Rule.`,
				},
				resource.Attribute{
					Name:        "secondary_key",
					Description: `The Secondary Key for the Event Hubs authorization Rule.`,
				},
				resource.Attribute{
					Name:        "send",
					Description: `Does this Authorization Rule have permissions to Send to the Event Hub? ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Authorization Rule.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_express_route_circuit",
			Category:         "Network",
			ShortDescription: `Gets information about an existing ExpressRoute circuit.`,
			Description: `

Use this data source to access information about an existing ExpressRoute circuit.

`,
			Icon: "Networking/10079-icon-service-ExpressRoute-Circuits.svg",
			Keywords: []string{
				"network",
				"express",
				"route",
				"circuit",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the ExpressRoute circuit.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The Name of the Resource Group where the ExpressRoute circuit exists. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the ExpressRoute circuit.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the ExpressRoute circuit exists`,
				},
				resource.Attribute{
					Name:        "peerings",
					Description: `A ` + "`" + `peerings` + "`" + ` block for the ExpressRoute circuit as documented below`,
				},
				resource.Attribute{
					Name:        "service_provider_provisioning_state",
					Description: `The ExpressRoute circuit provisioning state from your chosen service provider. Possible values are "NotProvisioned", "Provisioning", "Provisioned", and "Deprovisioning".`,
				},
				resource.Attribute{
					Name:        "service_key",
					Description: `The string needed by the service provider to provision the ExpressRoute circuit.`,
				},
				resource.Attribute{
					Name:        "service_provider_properties",
					Description: `A ` + "`" + `service_provider_properties` + "`" + ` block for the ExpressRoute circuit as documented below`,
				},
				resource.Attribute{
					Name:        "sku",
					Description: `A ` + "`" + `sku` + "`" + ` block for the ExpressRoute circuit as documented below. --- ` + "`" + `service_provider_properties` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "service_provider_name",
					Description: `The name of the ExpressRoute Service Provider.`,
				},
				resource.Attribute{
					Name:        "peering_location",
					Description: `The name of the peering location and`,
				},
				resource.Attribute{
					Name:        "bandwidth_in_mbps",
					Description: `The bandwidth in Mbps of the ExpressRoute circuit. ` + "`" + `peerings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "peering_type",
					Description: `The type of the ExpressRoute Circuit Peering. Acceptable values include ` + "`" + `AzurePrivatePeering` + "`" + `, ` + "`" + `AzurePublicPeering` + "`" + ` and ` + "`" + `MicrosoftPeering` + "`" + `. Changing this forces a new resource to be created. ~>`,
				},
				resource.Attribute{
					Name:        "primary_peer_address_prefix",
					Description: `A ` + "`" + `/30` + "`" + ` subnet for the primary link.`,
				},
				resource.Attribute{
					Name:        "secondary_peer_address_prefix",
					Description: `A ` + "`" + `/30` + "`" + ` subnet for the secondary link.`,
				},
				resource.Attribute{
					Name:        "vlan_id",
					Description: `A valid VLAN ID to establish this peering on.`,
				},
				resource.Attribute{
					Name:        "shared_key",
					Description: `The shared key. Can be a maximum of 25 characters.`,
				},
				resource.Attribute{
					Name:        "azure_asn",
					Description: `The Either a 16-bit or a 32-bit ASN for Azure.`,
				},
				resource.Attribute{
					Name:        "peer_asn",
					Description: `The Either a 16-bit or a 32-bit ASN. Can either be public or private. ` + "`" + `sku` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "tier",
					Description: `The service tier. Possible values are ` + "`" + `Basic` + "`" + `, ` + "`" + `Local` + "`" + `, ` + "`" + `Standard` + "`" + ` or ` + "`" + `Premium` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "family",
					Description: `The billing mode for bandwidth. Possible values are ` + "`" + `MeteredData` + "`" + ` or ` + "`" + `UnlimitedData` + "`" + `. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the ExpressRoute circuit.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the ExpressRoute circuit.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the ExpressRoute circuit exists`,
				},
				resource.Attribute{
					Name:        "peerings",
					Description: `A ` + "`" + `peerings` + "`" + ` block for the ExpressRoute circuit as documented below`,
				},
				resource.Attribute{
					Name:        "service_provider_provisioning_state",
					Description: `The ExpressRoute circuit provisioning state from your chosen service provider. Possible values are "NotProvisioned", "Provisioning", "Provisioned", and "Deprovisioning".`,
				},
				resource.Attribute{
					Name:        "service_key",
					Description: `The string needed by the service provider to provision the ExpressRoute circuit.`,
				},
				resource.Attribute{
					Name:        "service_provider_properties",
					Description: `A ` + "`" + `service_provider_properties` + "`" + ` block for the ExpressRoute circuit as documented below`,
				},
				resource.Attribute{
					Name:        "sku",
					Description: `A ` + "`" + `sku` + "`" + ` block for the ExpressRoute circuit as documented below. --- ` + "`" + `service_provider_properties` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "service_provider_name",
					Description: `The name of the ExpressRoute Service Provider.`,
				},
				resource.Attribute{
					Name:        "peering_location",
					Description: `The name of the peering location and`,
				},
				resource.Attribute{
					Name:        "bandwidth_in_mbps",
					Description: `The bandwidth in Mbps of the ExpressRoute circuit. ` + "`" + `peerings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "peering_type",
					Description: `The type of the ExpressRoute Circuit Peering. Acceptable values include ` + "`" + `AzurePrivatePeering` + "`" + `, ` + "`" + `AzurePublicPeering` + "`" + ` and ` + "`" + `MicrosoftPeering` + "`" + `. Changing this forces a new resource to be created. ~>`,
				},
				resource.Attribute{
					Name:        "primary_peer_address_prefix",
					Description: `A ` + "`" + `/30` + "`" + ` subnet for the primary link.`,
				},
				resource.Attribute{
					Name:        "secondary_peer_address_prefix",
					Description: `A ` + "`" + `/30` + "`" + ` subnet for the secondary link.`,
				},
				resource.Attribute{
					Name:        "vlan_id",
					Description: `A valid VLAN ID to establish this peering on.`,
				},
				resource.Attribute{
					Name:        "shared_key",
					Description: `The shared key. Can be a maximum of 25 characters.`,
				},
				resource.Attribute{
					Name:        "azure_asn",
					Description: `The Either a 16-bit or a 32-bit ASN for Azure.`,
				},
				resource.Attribute{
					Name:        "peer_asn",
					Description: `The Either a 16-bit or a 32-bit ASN. Can either be public or private. ` + "`" + `sku` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "tier",
					Description: `The service tier. Possible values are ` + "`" + `Basic` + "`" + `, ` + "`" + `Local` + "`" + `, ` + "`" + `Standard` + "`" + ` or ` + "`" + `Premium` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "family",
					Description: `The billing mode for bandwidth. Possible values are ` + "`" + `MeteredData` + "`" + ` or ` + "`" + `UnlimitedData` + "`" + `. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the ExpressRoute circuit.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_firewall",
			Category:         "Network",
			ShortDescription: `Gets information about an existing Azure Firewall.`,
			Description: `

Use this data source to access information about an existing Azure Firewall.

`,
			Keywords: []string{
				"network",
				"firewall",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Azure Firewall.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of the Resource Group in which the Azure Firewall exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Azure Firewall.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the Azure Firewall exists.`,
				},
				resource.Attribute{
					Name:        "sku_name",
					Description: `The sku name of the Azure Firewall.`,
				},
				resource.Attribute{
					Name:        "sku_tier",
					Description: `The sku tier of the Azure Firewall.`,
				},
				resource.Attribute{
					Name:        "firewall_policy_id",
					Description: `The ID of the Firewall Policy applied to the Azure Firewall.`,
				},
				resource.Attribute{
					Name:        "ip_configuration",
					Description: `A ` + "`" + `ip_configuration` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "dns_servers",
					Description: `The list of DNS servers that the Azure Firewall will direct DNS traffic to the for name resolution.`,
				},
				resource.Attribute{
					Name:        "management_ip_configuration",
					Description: `A ` + "`" + `management_ip_configuration` + "`" + ` block as defined below, which allows force-tunnelling of traffic to be performed by the firewall.`,
				},
				resource.Attribute{
					Name:        "threat_intel_mode",
					Description: `The operation mode for threat intelligence-based filtering.`,
				},
				resource.Attribute{
					Name:        "virtual_hub",
					Description: `A ` + "`" + `virtual_hub` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `The availability zones in which the Azure Firewall is created.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the Azure Firewall. --- A ` + "`" + `ip_configuration` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet where the Azure Firewall is deployed.`,
				},
				resource.Attribute{
					Name:        "private_ip_address",
					Description: `The Private IP Address of the Azure Firewall.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet where the Azure Firewall is deployed.`,
				},
				resource.Attribute{
					Name:        "private_ip_address",
					Description: `The Private IP Address of the Azure Firewall.`,
				},
				resource.Attribute{
					Name:        "virtual_hub_id",
					Description: `The ID of the Virtual Hub where the Azure Firewall resides in.`,
				},
				resource.Attribute{
					Name:        "public_ip_count",
					Description: `The number of public IPs assigned to the Azure Firewall.`,
				},
				resource.Attribute{
					Name:        "public_ip_addresses",
					Description: `The list of public IP addresses associated with the Azure Firewall.`,
				},
				resource.Attribute{
					Name:        "private_ip_address",
					Description: `The private IP address associated with the Azure Firewall. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Firewall.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Azure Firewall.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the Azure Firewall exists.`,
				},
				resource.Attribute{
					Name:        "sku_name",
					Description: `The sku name of the Azure Firewall.`,
				},
				resource.Attribute{
					Name:        "sku_tier",
					Description: `The sku tier of the Azure Firewall.`,
				},
				resource.Attribute{
					Name:        "firewall_policy_id",
					Description: `The ID of the Firewall Policy applied to the Azure Firewall.`,
				},
				resource.Attribute{
					Name:        "ip_configuration",
					Description: `A ` + "`" + `ip_configuration` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "dns_servers",
					Description: `The list of DNS servers that the Azure Firewall will direct DNS traffic to the for name resolution.`,
				},
				resource.Attribute{
					Name:        "management_ip_configuration",
					Description: `A ` + "`" + `management_ip_configuration` + "`" + ` block as defined below, which allows force-tunnelling of traffic to be performed by the firewall.`,
				},
				resource.Attribute{
					Name:        "threat_intel_mode",
					Description: `The operation mode for threat intelligence-based filtering.`,
				},
				resource.Attribute{
					Name:        "virtual_hub",
					Description: `A ` + "`" + `virtual_hub` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `The availability zones in which the Azure Firewall is created.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the Azure Firewall. --- A ` + "`" + `ip_configuration` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet where the Azure Firewall is deployed.`,
				},
				resource.Attribute{
					Name:        "private_ip_address",
					Description: `The Private IP Address of the Azure Firewall.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet where the Azure Firewall is deployed.`,
				},
				resource.Attribute{
					Name:        "private_ip_address",
					Description: `The Private IP Address of the Azure Firewall.`,
				},
				resource.Attribute{
					Name:        "virtual_hub_id",
					Description: `The ID of the Virtual Hub where the Azure Firewall resides in.`,
				},
				resource.Attribute{
					Name:        "public_ip_count",
					Description: `The number of public IPs assigned to the Azure Firewall.`,
				},
				resource.Attribute{
					Name:        "public_ip_addresses",
					Description: `The list of public IP addresses associated with the Azure Firewall.`,
				},
				resource.Attribute{
					Name:        "private_ip_address",
					Description: `The private IP address associated with the Azure Firewall. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Firewall.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_firewall_policy",
			Category:         "Network",
			ShortDescription: `Gets information about an existing Firewall Policy.`,
			Description: `

Use this data source to access information about an existing Firewall Policy.

`,
			Keywords: []string{
				"network",
				"firewall",
				"policy",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Firewall Policy.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the Firewall Policy. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Firewall Policy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_function_app",
			Category:         "App Service (Web Apps)",
			ShortDescription: `Gets information about an existing Function App.`,
			Description: `

Use this data source to access information about a Function App.

`,
			Icon: "IoT/10029-icon-service-Function-Apps.svg",
			Keywords: []string{
				"app",
				"service",
				"web",
				"apps",
				"function",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Function App resource.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of the Resource Group where the Function App exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Function App`,
				},
				resource.Attribute{
					Name:        "app_service_plan_id",
					Description: `The ID of the App Service Plan within which to create this Function App.`,
				},
				resource.Attribute{
					Name:        "app_settings",
					Description: `A key-value pair of App Settings.`,
				},
				resource.Attribute{
					Name:        "connection_string",
					Description: `An ` + "`" + `connection_string` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "custom_domain_verification_id",
					Description: `An identifier used by App Service to perform domain ownership verification via DNS TXT record.`,
				},
				resource.Attribute{
					Name:        "default_hostname",
					Description: `The default hostname associated with the Function App.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Is the Function App enabled?`,
				},
				resource.Attribute{
					Name:        "identity",
					Description: `A ` + "`" + `identity` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "site_credential",
					Description: `A ` + "`" + `site_credential` + "`" + ` block as defined below, which contains the site-level credentials used to publish to this App Service.`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `A string indicating the Operating System type for this function app. ~>`,
				},
				resource.Attribute{
					Name:        "outbound_ip_addresses",
					Description: `A comma separated list of outbound IP addresses.`,
				},
				resource.Attribute{
					Name:        "possible_outbound_ip_addresses",
					Description: `A comma separated list of outbound IP addresses, not all of which are necessarily in use. Superset of ` + "`" + `outbound_ip_addresses` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "source_control",
					Description: `A ` + "`" + `source_control` + "`" + ` block as defined below. --- The ` + "`" + `connection_string` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Connection String.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the Connection String.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value for the Connection String. --- The ` + "`" + `site_credential` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The username which can be used to publish to this App Service`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `The password associated with the username, which can be used to publish to this App Service. --- An ` + "`" + `ip_restriction` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The IP Address used for this IP Restriction.`,
				},
				resource.Attribute{
					Name:        "service_tag",
					Description: `The Service Tag used for this IP Restriction.`,
				},
				resource.Attribute{
					Name:        "subnet_mask",
					Description: `The Subnet mask used for this IP Restriction.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name for this IP Restriction.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority for this IP Restriction.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Does this restriction ` + "`" + `Allow` + "`" + ` or ` + "`" + `Deny` + "`" + ` access for this IP range? --- An ` + "`" + `scm_ip_restriction` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The IP Address used for this IP Restriction in CIDR notation.`,
				},
				resource.Attribute{
					Name:        "service_tag",
					Description: `The Service Tag used for this IP Restriction.`,
				},
				resource.Attribute{
					Name:        "virtual_network_subnet_id",
					Description: `The Virtual Network Subnet ID used for this IP Restriction.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name for this IP Restriction.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority for this IP Restriction.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Allow or Deny access for this IP range. Defaults to Allow. --- A ` + "`" + `site_config` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "always_on",
					Description: `Is the app loaded at all times?`,
				},
				resource.Attribute{
					Name:        "cors",
					Description: `A ` + "`" + `cors` + "`" + ` block as defined above.`,
				},
				resource.Attribute{
					Name:        "http2_enabled",
					Description: `Is HTTP2 Enabled on this App Service?`,
				},
				resource.Attribute{
					Name:        "ftps_state",
					Description: `State of FTP / FTPS service for this AppService.`,
				},
				resource.Attribute{
					Name:        "ip_restriction",
					Description: `One or more ` + "`" + `ip_restriction` + "`" + ` blocks as defined above.`,
				},
				resource.Attribute{
					Name:        "pre_warmed_instance_count",
					Description: `The number of pre-warmed instances for this function app. Only applicable to apps on the Premium plan.`,
				},
				resource.Attribute{
					Name:        "scm_use_main_ip_restriction",
					Description: `IP security restrictions for scm to use main.`,
				},
				resource.Attribute{
					Name:        "scm_ip_restriction",
					Description: `One or more ` + "`" + `scm_ip_restriction` + "`" + ` blocks as defined above.`,
				},
				resource.Attribute{
					Name:        "linux_fx_version",
					Description: `Linux App Framework and version for the AppService.`,
				},
				resource.Attribute{
					Name:        "min_tls_version",
					Description: `The minimum supported TLS version for this App Service.`,
				},
				resource.Attribute{
					Name:        "scm_type",
					Description: `The type of Source Control enabled for this App Service.`,
				},
				resource.Attribute{
					Name:        "use_32_bit_worker_process",
					Description: `Does the App Service run in 32 bit mode, rather than 64 bit mode?`,
				},
				resource.Attribute{
					Name:        "websockets_enabled",
					Description: `Are WebSockets enabled for this App Service? --- A ` + "`" + `source_control` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "repo_url",
					Description: `The URL of the source code repository.`,
				},
				resource.Attribute{
					Name:        "branch",
					Description: `The branch of the remote repository in use.`,
				},
				resource.Attribute{
					Name:        "manual_integration",
					Description: `Limits to manual integration.`,
				},
				resource.Attribute{
					Name:        "rollback_enabled",
					Description: `Is roll-back enabled for the repository.`,
				},
				resource.Attribute{
					Name:        "use_mercurial",
					Description: `Uses Mercurial if ` + "`" + `true` + "`" + `, otherwise uses Git. --- An ` + "`" + `identity` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "principal_id",
					Description: `The ID of the System Managed Service Principal assigned to the function app.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `The ID of the Tenant of the System Managed Service Principal assigned to the function app.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The identity type of the Managed Identity assigned to the function app. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Function App.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Function App`,
				},
				resource.Attribute{
					Name:        "app_service_plan_id",
					Description: `The ID of the App Service Plan within which to create this Function App.`,
				},
				resource.Attribute{
					Name:        "app_settings",
					Description: `A key-value pair of App Settings.`,
				},
				resource.Attribute{
					Name:        "connection_string",
					Description: `An ` + "`" + `connection_string` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "custom_domain_verification_id",
					Description: `An identifier used by App Service to perform domain ownership verification via DNS TXT record.`,
				},
				resource.Attribute{
					Name:        "default_hostname",
					Description: `The default hostname associated with the Function App.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Is the Function App enabled?`,
				},
				resource.Attribute{
					Name:        "identity",
					Description: `A ` + "`" + `identity` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "site_credential",
					Description: `A ` + "`" + `site_credential` + "`" + ` block as defined below, which contains the site-level credentials used to publish to this App Service.`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `A string indicating the Operating System type for this function app. ~>`,
				},
				resource.Attribute{
					Name:        "outbound_ip_addresses",
					Description: `A comma separated list of outbound IP addresses.`,
				},
				resource.Attribute{
					Name:        "possible_outbound_ip_addresses",
					Description: `A comma separated list of outbound IP addresses, not all of which are necessarily in use. Superset of ` + "`" + `outbound_ip_addresses` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "source_control",
					Description: `A ` + "`" + `source_control` + "`" + ` block as defined below. --- The ` + "`" + `connection_string` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Connection String.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the Connection String.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value for the Connection String. --- The ` + "`" + `site_credential` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The username which can be used to publish to this App Service`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `The password associated with the username, which can be used to publish to this App Service. --- An ` + "`" + `ip_restriction` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The IP Address used for this IP Restriction.`,
				},
				resource.Attribute{
					Name:        "service_tag",
					Description: `The Service Tag used for this IP Restriction.`,
				},
				resource.Attribute{
					Name:        "subnet_mask",
					Description: `The Subnet mask used for this IP Restriction.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name for this IP Restriction.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority for this IP Restriction.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Does this restriction ` + "`" + `Allow` + "`" + ` or ` + "`" + `Deny` + "`" + ` access for this IP range? --- An ` + "`" + `scm_ip_restriction` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The IP Address used for this IP Restriction in CIDR notation.`,
				},
				resource.Attribute{
					Name:        "service_tag",
					Description: `The Service Tag used for this IP Restriction.`,
				},
				resource.Attribute{
					Name:        "virtual_network_subnet_id",
					Description: `The Virtual Network Subnet ID used for this IP Restriction.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name for this IP Restriction.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority for this IP Restriction.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Allow or Deny access for this IP range. Defaults to Allow. --- A ` + "`" + `site_config` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "always_on",
					Description: `Is the app loaded at all times?`,
				},
				resource.Attribute{
					Name:        "cors",
					Description: `A ` + "`" + `cors` + "`" + ` block as defined above.`,
				},
				resource.Attribute{
					Name:        "http2_enabled",
					Description: `Is HTTP2 Enabled on this App Service?`,
				},
				resource.Attribute{
					Name:        "ftps_state",
					Description: `State of FTP / FTPS service for this AppService.`,
				},
				resource.Attribute{
					Name:        "ip_restriction",
					Description: `One or more ` + "`" + `ip_restriction` + "`" + ` blocks as defined above.`,
				},
				resource.Attribute{
					Name:        "pre_warmed_instance_count",
					Description: `The number of pre-warmed instances for this function app. Only applicable to apps on the Premium plan.`,
				},
				resource.Attribute{
					Name:        "scm_use_main_ip_restriction",
					Description: `IP security restrictions for scm to use main.`,
				},
				resource.Attribute{
					Name:        "scm_ip_restriction",
					Description: `One or more ` + "`" + `scm_ip_restriction` + "`" + ` blocks as defined above.`,
				},
				resource.Attribute{
					Name:        "linux_fx_version",
					Description: `Linux App Framework and version for the AppService.`,
				},
				resource.Attribute{
					Name:        "min_tls_version",
					Description: `The minimum supported TLS version for this App Service.`,
				},
				resource.Attribute{
					Name:        "scm_type",
					Description: `The type of Source Control enabled for this App Service.`,
				},
				resource.Attribute{
					Name:        "use_32_bit_worker_process",
					Description: `Does the App Service run in 32 bit mode, rather than 64 bit mode?`,
				},
				resource.Attribute{
					Name:        "websockets_enabled",
					Description: `Are WebSockets enabled for this App Service? --- A ` + "`" + `source_control` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "repo_url",
					Description: `The URL of the source code repository.`,
				},
				resource.Attribute{
					Name:        "branch",
					Description: `The branch of the remote repository in use.`,
				},
				resource.Attribute{
					Name:        "manual_integration",
					Description: `Limits to manual integration.`,
				},
				resource.Attribute{
					Name:        "rollback_enabled",
					Description: `Is roll-back enabled for the repository.`,
				},
				resource.Attribute{
					Name:        "use_mercurial",
					Description: `Uses Mercurial if ` + "`" + `true` + "`" + `, otherwise uses Git. --- An ` + "`" + `identity` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "principal_id",
					Description: `The ID of the System Managed Service Principal assigned to the function app.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `The ID of the Tenant of the System Managed Service Principal assigned to the function app.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The identity type of the Managed Identity assigned to the function app. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Function App.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_function_app_host_keys",
			Category:         "Resources",
			ShortDescription: `Gets the Host Keys of an existing Function App.`,
			Description: `

Use this data source to fetch the Host Keys of an existing Function App

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_hdinsight_cluster",
			Category:         "HDInsight",
			ShortDescription: `Gets information about an existing HDInsight Cluster.`,
			Description: `

Use this data source to access information about an existing HDInsight Cluster.

`,
			Icon: "Analytics/00037-icon-service-HD-Insight-Clusters.svg",
			Keywords: []string{
				"hdinsight",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name of this HDInsight Cluster.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `Specifies the name of the Resource Group in which this HDInsight Cluster exists. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure Region in which this HDInsight Cluster exists.`,
				},
				resource.Attribute{
					Name:        "cluster_version",
					Description: `The version of HDInsights which is used on this HDInsight Cluster.`,
				},
				resource.Attribute{
					Name:        "component_versions",
					Description: `A map of versions of software used on this HDInsights Cluster.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `A ` + "`" + `gateway` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "edge_ssh_endpoint",
					Description: `The SSH Endpoint of the Edge Node for this HDInsight Cluster, if an Edge Node exists.`,
				},
				resource.Attribute{
					Name:        "https_endpoint",
					Description: `The HTTPS Endpoint for this HDInsight Cluster.`,
				},
				resource.Attribute{
					Name:        "kafka_rest_proxy_endpoint",
					Description: `The Kafka Rest Proxy Endpoint for this HDInsight Cluster.`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `The kind of HDInsight Cluster this is, such as a Spark or Storm cluster.`,
				},
				resource.Attribute{
					Name:        "tier",
					Description: `The SKU / Tier of this HDInsight Cluster.`,
				},
				resource.Attribute{
					Name:        "ssh_endpoint",
					Description: `The SSH Endpoint for this HDInsight Cluster.`,
				},
				resource.Attribute{
					Name:        "tls_min_version",
					Description: `The minimal supported tls version.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A map of tags assigned to the HDInsight Cluster. --- A ` + "`" + `gateway` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Is the Ambari Portal enabled?`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The username used for the Ambari Portal.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `The password used for the Ambari Portal. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the HDInsight Cluster.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "location",
					Description: `The Azure Region in which this HDInsight Cluster exists.`,
				},
				resource.Attribute{
					Name:        "cluster_version",
					Description: `The version of HDInsights which is used on this HDInsight Cluster.`,
				},
				resource.Attribute{
					Name:        "component_versions",
					Description: `A map of versions of software used on this HDInsights Cluster.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `A ` + "`" + `gateway` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "edge_ssh_endpoint",
					Description: `The SSH Endpoint of the Edge Node for this HDInsight Cluster, if an Edge Node exists.`,
				},
				resource.Attribute{
					Name:        "https_endpoint",
					Description: `The HTTPS Endpoint for this HDInsight Cluster.`,
				},
				resource.Attribute{
					Name:        "kafka_rest_proxy_endpoint",
					Description: `The Kafka Rest Proxy Endpoint for this HDInsight Cluster.`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `The kind of HDInsight Cluster this is, such as a Spark or Storm cluster.`,
				},
				resource.Attribute{
					Name:        "tier",
					Description: `The SKU / Tier of this HDInsight Cluster.`,
				},
				resource.Attribute{
					Name:        "ssh_endpoint",
					Description: `The SSH Endpoint for this HDInsight Cluster.`,
				},
				resource.Attribute{
					Name:        "tls_min_version",
					Description: `The minimal supported tls version.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A map of tags assigned to the HDInsight Cluster. --- A ` + "`" + `gateway` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Is the Ambari Portal enabled?`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The username used for the Ambari Portal.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `The password used for the Ambari Portal. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the HDInsight Cluster.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_healthcare_service",
			Category:         "Healthcare",
			ShortDescription: `Get information about an existing Healthcare Service`,
			Description: `

Use this data source to access information about an existing Healthcare Service

`,
			Keywords: []string{
				"healthcare",
				"service",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name of the Healthcare Service.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of the Resource Group in which the Healthcare Service exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure Region where the Service is located. ~>`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `The type of the service.`,
				},
				resource.Attribute{
					Name:        "authentication_configuration",
					Description: `An ` + "`" + `authentication_configuration` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "cosmosdb_offer_throughput",
					Description: `The provisioned throughput for the backing database.`,
				},
				resource.Attribute{
					Name:        "cors_configuration",
					Description: `A ` + "`" + `cors_configuration` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the resource. --- An ` + "`" + `authentication_configuration` + "`" + ` exports the following:`,
				},
				resource.Attribute{
					Name:        "authority",
					Description: `The Azure Active Directory (tenant) that serves as the authentication authority to access the service.`,
				},
				resource.Attribute{
					Name:        "audience",
					Description: `The intended audience to receive authentication tokens for the service.`,
				},
				resource.Attribute{
					Name:        "smart_proxy_enabled",
					Description: `Is the 'SMART on FHIR' option for mobile and web implementations enabled? --- A ` + "`" + `cors_configuration` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "allowed_origins",
					Description: `The set of origins to be allowed via CORS.`,
				},
				resource.Attribute{
					Name:        "allowed_headers",
					Description: `The set of headers to be allowed via CORS.`,
				},
				resource.Attribute{
					Name:        "allowed_methods",
					Description: `The methods to be allowed via CORS.`,
				},
				resource.Attribute{
					Name:        "max_age_in_seconds",
					Description: `The max age to be allowed via CORS.`,
				},
				resource.Attribute{
					Name:        "allow_credentials",
					Description: `Are credentials are allowed via CORS? ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Healthcare Service.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "location",
					Description: `The Azure Region where the Service is located. ~>`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `The type of the service.`,
				},
				resource.Attribute{
					Name:        "authentication_configuration",
					Description: `An ` + "`" + `authentication_configuration` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "cosmosdb_offer_throughput",
					Description: `The provisioned throughput for the backing database.`,
				},
				resource.Attribute{
					Name:        "cors_configuration",
					Description: `A ` + "`" + `cors_configuration` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the resource. --- An ` + "`" + `authentication_configuration` + "`" + ` exports the following:`,
				},
				resource.Attribute{
					Name:        "authority",
					Description: `The Azure Active Directory (tenant) that serves as the authentication authority to access the service.`,
				},
				resource.Attribute{
					Name:        "audience",
					Description: `The intended audience to receive authentication tokens for the service.`,
				},
				resource.Attribute{
					Name:        "smart_proxy_enabled",
					Description: `Is the 'SMART on FHIR' option for mobile and web implementations enabled? --- A ` + "`" + `cors_configuration` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "allowed_origins",
					Description: `The set of origins to be allowed via CORS.`,
				},
				resource.Attribute{
					Name:        "allowed_headers",
					Description: `The set of headers to be allowed via CORS.`,
				},
				resource.Attribute{
					Name:        "allowed_methods",
					Description: `The methods to be allowed via CORS.`,
				},
				resource.Attribute{
					Name:        "max_age_in_seconds",
					Description: `The max age to be allowed via CORS.`,
				},
				resource.Attribute{
					Name:        "allow_credentials",
					Description: `Are credentials are allowed via CORS? ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Healthcare Service.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_image",
			Category:         "Compute",
			ShortDescription: `Gets information about an existing Image`,
			Description: `

Use this data source to access information about an existing Image.

`,
			Icon: "Compute/10033-icon-service-Images.svg",
			Keywords: []string{
				"compute",
				"image",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the Image.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) Regex pattern of the image to match.`,
				},
				resource.Attribute{
					Name:        "sort_descending",
					Description: `(Optional) By default when matching by regex, images are sorted by name in ascending order and the first match is chosen, to sort descending, set this flag.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The Name of the Resource Group where this Image exists. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "data_disk",
					Description: `a collection of ` + "`" + `data_disk` + "`" + ` blocks as defined below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `the name of the Image.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `the Azure Location where this Image exists.`,
				},
				resource.Attribute{
					Name:        "os_disk",
					Description: `a ` + "`" + `os_disk` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `a mapping of tags to assigned to the resource.`,
				},
				resource.Attribute{
					Name:        "zone_resilient",
					Description: `is zone resiliency enabled? ` + "`" + `os_disk` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "blob_uri",
					Description: `the URI in Azure storage of the blob used to create the image.`,
				},
				resource.Attribute{
					Name:        "caching",
					Description: `the caching mode for the OS Disk, such as ` + "`" + `ReadWrite` + "`" + `, ` + "`" + `ReadOnly` + "`" + `, or ` + "`" + `None` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "managed_disk_id",
					Description: `the ID of the Managed Disk used as the OS Disk Image.`,
				},
				resource.Attribute{
					Name:        "os_state",
					Description: `the State of the OS used in the Image, such as ` + "`" + `Generalized` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `the type of Operating System used on the OS Disk. such as ` + "`" + `Linux` + "`" + ` or ` + "`" + `Windows` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "size_gb",
					Description: `the size of the OS Disk in GB. ` + "`" + `data_disk` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "blob_uri",
					Description: `the URI in Azure storage of the blob used to create the image.`,
				},
				resource.Attribute{
					Name:        "caching",
					Description: `the caching mode for the Data Disk, such as ` + "`" + `ReadWrite` + "`" + `, ` + "`" + `ReadOnly` + "`" + `, or ` + "`" + `None` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lun",
					Description: `the logical unit number of the data disk.`,
				},
				resource.Attribute{
					Name:        "managed_disk_id",
					Description: `the ID of the Managed Disk used as the Data Disk Image.`,
				},
				resource.Attribute{
					Name:        "size_gb",
					Description: `the size of this Data Disk in GB. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Image.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "data_disk",
					Description: `a collection of ` + "`" + `data_disk` + "`" + ` blocks as defined below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `the name of the Image.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `the Azure Location where this Image exists.`,
				},
				resource.Attribute{
					Name:        "os_disk",
					Description: `a ` + "`" + `os_disk` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `a mapping of tags to assigned to the resource.`,
				},
				resource.Attribute{
					Name:        "zone_resilient",
					Description: `is zone resiliency enabled? ` + "`" + `os_disk` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "blob_uri",
					Description: `the URI in Azure storage of the blob used to create the image.`,
				},
				resource.Attribute{
					Name:        "caching",
					Description: `the caching mode for the OS Disk, such as ` + "`" + `ReadWrite` + "`" + `, ` + "`" + `ReadOnly` + "`" + `, or ` + "`" + `None` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "managed_disk_id",
					Description: `the ID of the Managed Disk used as the OS Disk Image.`,
				},
				resource.Attribute{
					Name:        "os_state",
					Description: `the State of the OS used in the Image, such as ` + "`" + `Generalized` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `the type of Operating System used on the OS Disk. such as ` + "`" + `Linux` + "`" + ` or ` + "`" + `Windows` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "size_gb",
					Description: `the size of the OS Disk in GB. ` + "`" + `data_disk` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "blob_uri",
					Description: `the URI in Azure storage of the blob used to create the image.`,
				},
				resource.Attribute{
					Name:        "caching",
					Description: `the caching mode for the Data Disk, such as ` + "`" + `ReadWrite` + "`" + `, ` + "`" + `ReadOnly` + "`" + `, or ` + "`" + `None` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lun",
					Description: `the logical unit number of the data disk.`,
				},
				resource.Attribute{
					Name:        "managed_disk_id",
					Description: `the ID of the Managed Disk used as the Data Disk Image.`,
				},
				resource.Attribute{
					Name:        "size_gb",
					Description: `the size of this Data Disk in GB. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Image.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_images",
			Category:         "Compute",
			ShortDescription: `Gets information about existing Images within a Resource Group.`,
			Description: `

Use this data source to access information about existing Images within a Resource Group.

`,
			Keywords: []string{
				"compute",
				"images",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of the Resource Group in which the Image exists.`,
				},
				resource.Attribute{
					Name:        "tags_filter",
					Description: `A mapping of tags to filter the list of images against. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "images",
					Description: `One or more ` + "`" + `images` + "`" + ` blocks as defined below: --- A ` + "`" + `images` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Image.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The supported Azure location where the Image exists.`,
				},
				resource.Attribute{
					Name:        "zone_resilient",
					Description: `Is zone resiliency enabled?`,
				},
				resource.Attribute{
					Name:        "os_disk",
					Description: `An ` + "`" + `os_disk` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "data_disk",
					Description: `One or more ` + "`" + `data_disk` + "`" + ` blocks as defined below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the Image. --- The ` + "`" + `os_disk` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "blob_uri",
					Description: `the URI in Azure storage of the blob used to create the image.`,
				},
				resource.Attribute{
					Name:        "caching",
					Description: `the caching mode for the OS Disk.`,
				},
				resource.Attribute{
					Name:        "managed_disk_id",
					Description: `the ID of the Managed Disk used as the OS Disk Image.`,
				},
				resource.Attribute{
					Name:        "os_state",
					Description: `the State of the OS used in the Image.`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `the type of Operating System used on the OS Disk.`,
				},
				resource.Attribute{
					Name:        "size_gb",
					Description: `the size of the OS Disk in GB. --- The ` + "`" + `data_disk` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "blob_uri",
					Description: `the URI in Azure storage of the blob used to create the image.`,
				},
				resource.Attribute{
					Name:        "caching",
					Description: `the caching mode for the Data Disk.`,
				},
				resource.Attribute{
					Name:        "lun",
					Description: `the logical unit number of the data disk.`,
				},
				resource.Attribute{
					Name:        "managed_disk_id",
					Description: `the ID of the Managed Disk used as the Data Disk Image.`,
				},
				resource.Attribute{
					Name:        "size_gb",
					Description: `the size of this Data Disk in GB. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Images within a Resource Group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "images",
					Description: `One or more ` + "`" + `images` + "`" + ` blocks as defined below: --- A ` + "`" + `images` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Image.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The supported Azure location where the Image exists.`,
				},
				resource.Attribute{
					Name:        "zone_resilient",
					Description: `Is zone resiliency enabled?`,
				},
				resource.Attribute{
					Name:        "os_disk",
					Description: `An ` + "`" + `os_disk` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "data_disk",
					Description: `One or more ` + "`" + `data_disk` + "`" + ` blocks as defined below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the Image. --- The ` + "`" + `os_disk` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "blob_uri",
					Description: `the URI in Azure storage of the blob used to create the image.`,
				},
				resource.Attribute{
					Name:        "caching",
					Description: `the caching mode for the OS Disk.`,
				},
				resource.Attribute{
					Name:        "managed_disk_id",
					Description: `the ID of the Managed Disk used as the OS Disk Image.`,
				},
				resource.Attribute{
					Name:        "os_state",
					Description: `the State of the OS used in the Image.`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `the type of Operating System used on the OS Disk.`,
				},
				resource.Attribute{
					Name:        "size_gb",
					Description: `the size of the OS Disk in GB. --- The ` + "`" + `data_disk` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "blob_uri",
					Description: `the URI in Azure storage of the blob used to create the image.`,
				},
				resource.Attribute{
					Name:        "caching",
					Description: `the caching mode for the Data Disk.`,
				},
				resource.Attribute{
					Name:        "lun",
					Description: `the logical unit number of the data disk.`,
				},
				resource.Attribute{
					Name:        "managed_disk_id",
					Description: `the ID of the Managed Disk used as the Data Disk Image.`,
				},
				resource.Attribute{
					Name:        "size_gb",
					Description: `the size of this Data Disk in GB. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Images within a Resource Group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_iothub",
			Category:         "IoT Hub",
			ShortDescription: `Gets information about an existing IoTHub.`,
			Description: `

Use this data source to access information about an existing IoTHub.

`,
			Keywords: []string{
				"iot",
				"hub",
				"iothub",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the IoTHub. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the IoTHub.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_iothub_dps",
			Category:         "IoT Hub",
			ShortDescription: `Gets information about an existing IoT Hub Device Provisioning Service.`,
			Description: `

Use this data source to access information about an existing IotHub Device Provisioning Service.

`,
			Keywords: []string{
				"iot",
				"hub",
				"iothub",
				"dps",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name of the Iot Device Provisioning Service resource.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of the resource group under which the Iot Device Provisioning Service is located in. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the IoT Device Provisioning Service.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Specifies the supported Azure location where the IoT Device Provisioning Service exists.`,
				},
				resource.Attribute{
					Name:        "allocation_policy",
					Description: `The allocation policy of the IoT Device Provisioning Service.`,
				},
				resource.Attribute{
					Name:        "device_provisioning_host_name",
					Description: `The device endpoint of the IoT Device Provisioning Service.`,
				},
				resource.Attribute{
					Name:        "id_scope",
					Description: `The unique identifier of the IoT Device Provisioning Service.`,
				},
				resource.Attribute{
					Name:        "service_operations_host_name",
					Description: `The service endpoint of the IoT Device Provisioning Service. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the IoT Hub Device Provisioning Service.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the IoT Device Provisioning Service.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Specifies the supported Azure location where the IoT Device Provisioning Service exists.`,
				},
				resource.Attribute{
					Name:        "allocation_policy",
					Description: `The allocation policy of the IoT Device Provisioning Service.`,
				},
				resource.Attribute{
					Name:        "device_provisioning_host_name",
					Description: `The device endpoint of the IoT Device Provisioning Service.`,
				},
				resource.Attribute{
					Name:        "id_scope",
					Description: `The unique identifier of the IoT Device Provisioning Service.`,
				},
				resource.Attribute{
					Name:        "service_operations_host_name",
					Description: `The service endpoint of the IoT Device Provisioning Service. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the IoT Hub Device Provisioning Service.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_iothub_dps_shared_access_policy",
			Category:         "IoT Hub",
			ShortDescription: `Gets information about an existing IotHub Device Provisioning Service Shared Access Policy`,
			Description: `

Use this data source to access information about an existing IotHub Device Provisioning Service Shared Access Policy

`,
			Keywords: []string{
				"iot",
				"hub",
				"iothub",
				"dps",
				"shared",
				"access",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name of the IotHub Shared Access Policy.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `Specifies the name of the resource group under which the IotHub Shared Access Policy resource exists.`,
				},
				resource.Attribute{
					Name:        "iothub_dps_name",
					Description: `Specifies the name of the IoT Hub Device Provisioning service to which the Shared Access Policy belongs. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the IoT Hub Device Provisioning Service Shared Access Policy.`,
				},
				resource.Attribute{
					Name:        "primary_key",
					Description: `The primary key used to create the authentication token.`,
				},
				resource.Attribute{
					Name:        "primary_connection_string",
					Description: `The primary connection string of the Shared Access Policy.`,
				},
				resource.Attribute{
					Name:        "secondary_key",
					Description: `The secondary key used to create the authentication token.`,
				},
				resource.Attribute{
					Name:        "secondary_connection_string",
					Description: `The secondary connection string of the Shared Access Policy. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the IotHub Device Provisioning Service Shared Access Policy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the IoT Hub Device Provisioning Service Shared Access Policy.`,
				},
				resource.Attribute{
					Name:        "primary_key",
					Description: `The primary key used to create the authentication token.`,
				},
				resource.Attribute{
					Name:        "primary_connection_string",
					Description: `The primary connection string of the Shared Access Policy.`,
				},
				resource.Attribute{
					Name:        "secondary_key",
					Description: `The secondary key used to create the authentication token.`,
				},
				resource.Attribute{
					Name:        "secondary_connection_string",
					Description: `The secondary connection string of the Shared Access Policy. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the IotHub Device Provisioning Service Shared Access Policy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_iothub_shared_access_policy",
			Category:         "IoT Hub",
			ShortDescription: `Gets information about an existing IotHub Shared Access Policy`,
			Description: `

Use this data source to access information about an existing IotHub Shared Access Policy

`,
			Keywords: []string{
				"iot",
				"hub",
				"iothub",
				"shared",
				"access",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name of the IotHub Shared Access Policy resource.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of the resource group under which the IotHub Shared Access Policy resource has to be created.`,
				},
				resource.Attribute{
					Name:        "iothub_name",
					Description: `The name of the IoTHub to which this Shared Access Policy belongs. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the IoTHub Shared Access Policy.`,
				},
				resource.Attribute{
					Name:        "primary_key",
					Description: `The primary key used to create the authentication token.`,
				},
				resource.Attribute{
					Name:        "primary_connection_string",
					Description: `The primary connection string of the Shared Access Policy.`,
				},
				resource.Attribute{
					Name:        "secondary_key",
					Description: `The secondary key used to create the authentication token.`,
				},
				resource.Attribute{
					Name:        "secondary_connection_string",
					Description: `The secondary connection string of the Shared Access Policy. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the IotHub Shared Access Policy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the IoTHub Shared Access Policy.`,
				},
				resource.Attribute{
					Name:        "primary_key",
					Description: `The primary key used to create the authentication token.`,
				},
				resource.Attribute{
					Name:        "primary_connection_string",
					Description: `The primary connection string of the Shared Access Policy.`,
				},
				resource.Attribute{
					Name:        "secondary_key",
					Description: `The secondary key used to create the authentication token.`,
				},
				resource.Attribute{
					Name:        "secondary_connection_string",
					Description: `The secondary connection string of the Shared Access Policy. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the IotHub Shared Access Policy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_ip_group",
			Category:         "Network",
			ShortDescription: `Gets information about an existing IP Group.`,
			Description: `

Use this data source to access information about an existing IP Group.

`,
			Keywords: []string{
				"network",
				"ip",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the Name of the IP Group.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `Specifies the Name of the Resource Group within which the IP Group exists ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the IP Group.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The supported Azure location where the resource exists.`,
				},
				resource.Attribute{
					Name:        "cidrs",
					Description: `A list of CIDRs or IP addresses.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the IP Group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the IP Group.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The supported Azure location where the resource exists.`,
				},
				resource.Attribute{
					Name:        "cidrs",
					Description: `A list of CIDRs or IP addresses.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the IP Group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_key_vault",
			Category:         "Key Vault",
			ShortDescription: `Gets information about an existing Key Vault.`,
			Description: `

Use this data source to access information about an existing Key Vault.

`,
			Keywords: []string{
				"key",
				"vault",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name of the Key Vault.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of the Resource Group in which the Key Vault exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Vault ID.`,
				},
				resource.Attribute{
					Name:        "vault_uri",
					Description: `The URI of the vault for performing operations on keys and secrets.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure Region in which the Key Vault exists.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `The Azure Active Directory Tenant ID used for authenticating requests to the Key Vault.`,
				},
				resource.Attribute{
					Name:        "sku_name",
					Description: `The Name of the SKU used for this Key Vault.`,
				},
				resource.Attribute{
					Name:        "access_policy",
					Description: `One or more ` + "`" + `access_policy` + "`" + ` blocks as defined below.`,
				},
				resource.Attribute{
					Name:        "enabled_for_deployment",
					Description: `Can Azure Virtual Machines retrieve certificates stored as secrets from the Key Vault?`,
				},
				resource.Attribute{
					Name:        "enabled_for_disk_encryption",
					Description: `Can Azure Disk Encryption retrieve secrets from the Key Vault?`,
				},
				resource.Attribute{
					Name:        "enabled_for_template_deployment",
					Description: `Can Azure Resource Manager retrieve secrets from the Key Vault?`,
				},
				resource.Attribute{
					Name:        "purge_protection_enabled",
					Description: `Is purge protection enabled on this Key Vault?`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the Key Vault. A ` + "`" + `access_policy` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `The Azure Active Directory Tenant ID used to authenticate requests for this Key Vault.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `An Object ID of a User, Service Principal or Security Group.`,
				},
				resource.Attribute{
					Name:        "application_id",
					Description: `The Object ID of a Azure Active Directory Application.`,
				},
				resource.Attribute{
					Name:        "certificate_permissions",
					Description: `A list of certificate permissions applicable to this Access Policy.`,
				},
				resource.Attribute{
					Name:        "key_permissions",
					Description: `A list of key permissions applicable to this Access Policy.`,
				},
				resource.Attribute{
					Name:        "secret_permissions",
					Description: `A list of secret permissions applicable to this Access Policy.`,
				},
				resource.Attribute{
					Name:        "storage_permissions",
					Description: `A list of storage permissions applicable to this Access Policy. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Key Vault.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Vault ID.`,
				},
				resource.Attribute{
					Name:        "vault_uri",
					Description: `The URI of the vault for performing operations on keys and secrets.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure Region in which the Key Vault exists.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `The Azure Active Directory Tenant ID used for authenticating requests to the Key Vault.`,
				},
				resource.Attribute{
					Name:        "sku_name",
					Description: `The Name of the SKU used for this Key Vault.`,
				},
				resource.Attribute{
					Name:        "access_policy",
					Description: `One or more ` + "`" + `access_policy` + "`" + ` blocks as defined below.`,
				},
				resource.Attribute{
					Name:        "enabled_for_deployment",
					Description: `Can Azure Virtual Machines retrieve certificates stored as secrets from the Key Vault?`,
				},
				resource.Attribute{
					Name:        "enabled_for_disk_encryption",
					Description: `Can Azure Disk Encryption retrieve secrets from the Key Vault?`,
				},
				resource.Attribute{
					Name:        "enabled_for_template_deployment",
					Description: `Can Azure Resource Manager retrieve secrets from the Key Vault?`,
				},
				resource.Attribute{
					Name:        "purge_protection_enabled",
					Description: `Is purge protection enabled on this Key Vault?`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the Key Vault. A ` + "`" + `access_policy` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `The Azure Active Directory Tenant ID used to authenticate requests for this Key Vault.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `An Object ID of a User, Service Principal or Security Group.`,
				},
				resource.Attribute{
					Name:        "application_id",
					Description: `The Object ID of a Azure Active Directory Application.`,
				},
				resource.Attribute{
					Name:        "certificate_permissions",
					Description: `A list of certificate permissions applicable to this Access Policy.`,
				},
				resource.Attribute{
					Name:        "key_permissions",
					Description: `A list of key permissions applicable to this Access Policy.`,
				},
				resource.Attribute{
					Name:        "secret_permissions",
					Description: `A list of secret permissions applicable to this Access Policy.`,
				},
				resource.Attribute{
					Name:        "storage_permissions",
					Description: `A list of storage permissions applicable to this Access Policy. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Key Vault.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_key_vault_access_policy",
			Category:         "Key Vault",
			ShortDescription: `Get information about the templated Access Policies for Key Vault.`,
			Description: `

Use this data source to access information about the permissions from the Management Key Vault Templates.

`,
			Keywords: []string{
				"key",
				"vault",
				"access",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name of the Management Template. Possible values are: ` + "`" + `Key Management` + "`" + `, ` + "`" + `Secret Management` + "`" + `, ` + "`" + `Certificate Management` + "`" + `, ` + "`" + `Key & Secret Management` + "`" + `, ` + "`" + `Key & Certificate Management` + "`" + `, ` + "`" + `Secret & Certificate Management` + "`" + `, ` + "`" + `Key, Secret, & Certificate Management` + "`" + ` ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `the ID of the Key Vault Access Policy`,
				},
				resource.Attribute{
					Name:        "key_permissions",
					Description: `the key permissions for the access policy`,
				},
				resource.Attribute{
					Name:        "secret_permissions",
					Description: `the secret permissions for the access policy`,
				},
				resource.Attribute{
					Name:        "certificate_permissions",
					Description: `the certificate permissions for the access policy ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Access Policy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `the ID of the Key Vault Access Policy`,
				},
				resource.Attribute{
					Name:        "key_permissions",
					Description: `the key permissions for the access policy`,
				},
				resource.Attribute{
					Name:        "secret_permissions",
					Description: `the secret permissions for the access policy`,
				},
				resource.Attribute{
					Name:        "certificate_permissions",
					Description: `the certificate permissions for the access policy ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Access Policy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_key_vault_certificate",
			Category:         "Key Vault",
			ShortDescription: `Gets information about an existing Key Vault Certificate.`,
			Description: `

Use this data source to access information about an existing Key Vault Certificate.

~> **Note:** All arguments including the secret value will be stored in the raw state as plain-text.
[Read more about sensitive data in state](/docs/state/sensitive-data.html).

`,
			Keywords: []string{
				"key",
				"vault",
				"certificate",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name of the Key Vault Certificate.`,
				},
				resource.Attribute{
					Name:        "key_vault_id",
					Description: `Specifies the ID of the Key Vault instance where the Secret resides, available on the ` + "`" + `azurerm_key_vault` + "`" + ` Data Source / Resource.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) Specifies the version of the certificate to look up. (Defaults to latest)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Key Vault Certificate ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name of the Key Vault Certificate.`,
				},
				resource.Attribute{
					Name:        "secret_id",
					Description: `The ID of the associated Key Vault Secret.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The current version of the Key Vault Certificate.`,
				},
				resource.Attribute{
					Name:        "certificate_data",
					Description: `The raw Key Vault Certificate data represented as a hexadecimal string.`,
				},
				resource.Attribute{
					Name:        "certificate_data_base64",
					Description: `The raw Key Vault Certificate data represented as a base64 string.`,
				},
				resource.Attribute{
					Name:        "thumbprint",
					Description: `The X509 Thumbprint of the Key Vault Certificate represented as a hexadecimal string.`,
				},
				resource.Attribute{
					Name:        "certificate_policy",
					Description: `A ` + "`" + `certificate_policy` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the resource. --- ` + "`" + `certificate_policy` + "`" + ` exports the following:`,
				},
				resource.Attribute{
					Name:        "issuer_parameters",
					Description: `A ` + "`" + `issuer_parameters` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "key_properties",
					Description: `A ` + "`" + `key_properties` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "lifetime_action",
					Description: `A ` + "`" + `lifetime_action` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "secret_properties",
					Description: `A ` + "`" + `secret_properties` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "x509_certificate_properties",
					Description: `An ` + "`" + `x509_certificate_properties` + "`" + ` block as defined below. --- ` + "`" + `issuer_parameters` + "`" + ` exports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Certificate Issuer. --- ` + "`" + `key_properties` + "`" + ` exports the following:`,
				},
				resource.Attribute{
					Name:        "exportable",
					Description: `Is this Certificate Exportable?`,
				},
				resource.Attribute{
					Name:        "key_size",
					Description: `The size of the Key used in the Certificate.`,
				},
				resource.Attribute{
					Name:        "key_type",
					Description: `Specifies the Type of Key, for example ` + "`" + `RSA` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "reuse_key",
					Description: `Is the key reusable? --- ` + "`" + `lifetime_action` + "`" + ` exports the following:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `A ` + "`" + `action` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "trigger",
					Description: `A ` + "`" + `trigger` + "`" + ` block as defined below. --- ` + "`" + `action` + "`" + ` exports the following:`,
				},
				resource.Attribute{
					Name:        "action_type",
					Description: `The Type of action to be performed when the lifetime trigger is triggerec. --- ` + "`" + `trigger` + "`" + ` exports the following:`,
				},
				resource.Attribute{
					Name:        "days_before_expiry",
					Description: `The number of days before the Certificate expires that the action associated with this Trigger should run.`,
				},
				resource.Attribute{
					Name:        "lifetime_percentage",
					Description: `The percentage at which during the Certificates Lifetime the action associated with this Trigger should run. --- ` + "`" + `secret_properties` + "`" + ` exports the following:`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `The Content-Type of the Certificate, for example ` + "`" + `application/x-pkcs12` + "`" + ` for a PFX or ` + "`" + `application/x-pem-file` + "`" + ` for a PEM. --- ` + "`" + `x509_certificate_properties` + "`" + ` exports the following:`,
				},
				resource.Attribute{
					Name:        "extended_key_usage",
					Description: `A list of Extended/Enhanced Key Usages.`,
				},
				resource.Attribute{
					Name:        "key_usage",
					Description: `A list of uses associated with this Key.`,
				},
				resource.Attribute{
					Name:        "subject",
					Description: `The Certificate's Subject.`,
				},
				resource.Attribute{
					Name:        "subject_alternative_names",
					Description: `A ` + "`" + `subject_alternative_names` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "validity_in_months",
					Description: `The Certificates Validity Period in Months. --- ` + "`" + `subject_alternative_names` + "`" + ` exports the following:`,
				},
				resource.Attribute{
					Name:        "dns_names",
					Description: `A list of alternative DNS names (FQDNs) identified by the Certificate.`,
				},
				resource.Attribute{
					Name:        "emails",
					Description: `A list of email addresses identified by this Certificate.`,
				},
				resource.Attribute{
					Name:        "upns",
					Description: `A list of User Principal Names identified by the Certificate. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Key Vault Certificate.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Key Vault Certificate ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name of the Key Vault Certificate.`,
				},
				resource.Attribute{
					Name:        "secret_id",
					Description: `The ID of the associated Key Vault Secret.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The current version of the Key Vault Certificate.`,
				},
				resource.Attribute{
					Name:        "certificate_data",
					Description: `The raw Key Vault Certificate data represented as a hexadecimal string.`,
				},
				resource.Attribute{
					Name:        "certificate_data_base64",
					Description: `The raw Key Vault Certificate data represented as a base64 string.`,
				},
				resource.Attribute{
					Name:        "thumbprint",
					Description: `The X509 Thumbprint of the Key Vault Certificate represented as a hexadecimal string.`,
				},
				resource.Attribute{
					Name:        "certificate_policy",
					Description: `A ` + "`" + `certificate_policy` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the resource. --- ` + "`" + `certificate_policy` + "`" + ` exports the following:`,
				},
				resource.Attribute{
					Name:        "issuer_parameters",
					Description: `A ` + "`" + `issuer_parameters` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "key_properties",
					Description: `A ` + "`" + `key_properties` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "lifetime_action",
					Description: `A ` + "`" + `lifetime_action` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "secret_properties",
					Description: `A ` + "`" + `secret_properties` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "x509_certificate_properties",
					Description: `An ` + "`" + `x509_certificate_properties` + "`" + ` block as defined below. --- ` + "`" + `issuer_parameters` + "`" + ` exports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Certificate Issuer. --- ` + "`" + `key_properties` + "`" + ` exports the following:`,
				},
				resource.Attribute{
					Name:        "exportable",
					Description: `Is this Certificate Exportable?`,
				},
				resource.Attribute{
					Name:        "key_size",
					Description: `The size of the Key used in the Certificate.`,
				},
				resource.Attribute{
					Name:        "key_type",
					Description: `Specifies the Type of Key, for example ` + "`" + `RSA` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "reuse_key",
					Description: `Is the key reusable? --- ` + "`" + `lifetime_action` + "`" + ` exports the following:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `A ` + "`" + `action` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "trigger",
					Description: `A ` + "`" + `trigger` + "`" + ` block as defined below. --- ` + "`" + `action` + "`" + ` exports the following:`,
				},
				resource.Attribute{
					Name:        "action_type",
					Description: `The Type of action to be performed when the lifetime trigger is triggerec. --- ` + "`" + `trigger` + "`" + ` exports the following:`,
				},
				resource.Attribute{
					Name:        "days_before_expiry",
					Description: `The number of days before the Certificate expires that the action associated with this Trigger should run.`,
				},
				resource.Attribute{
					Name:        "lifetime_percentage",
					Description: `The percentage at which during the Certificates Lifetime the action associated with this Trigger should run. --- ` + "`" + `secret_properties` + "`" + ` exports the following:`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `The Content-Type of the Certificate, for example ` + "`" + `application/x-pkcs12` + "`" + ` for a PFX or ` + "`" + `application/x-pem-file` + "`" + ` for a PEM. --- ` + "`" + `x509_certificate_properties` + "`" + ` exports the following:`,
				},
				resource.Attribute{
					Name:        "extended_key_usage",
					Description: `A list of Extended/Enhanced Key Usages.`,
				},
				resource.Attribute{
					Name:        "key_usage",
					Description: `A list of uses associated with this Key.`,
				},
				resource.Attribute{
					Name:        "subject",
					Description: `The Certificate's Subject.`,
				},
				resource.Attribute{
					Name:        "subject_alternative_names",
					Description: `A ` + "`" + `subject_alternative_names` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "validity_in_months",
					Description: `The Certificates Validity Period in Months. --- ` + "`" + `subject_alternative_names` + "`" + ` exports the following:`,
				},
				resource.Attribute{
					Name:        "dns_names",
					Description: `A list of alternative DNS names (FQDNs) identified by the Certificate.`,
				},
				resource.Attribute{
					Name:        "emails",
					Description: `A list of email addresses identified by this Certificate.`,
				},
				resource.Attribute{
					Name:        "upns",
					Description: `A list of User Principal Names identified by the Certificate. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Key Vault Certificate.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_key_vault_certificate_data",
			Category:         "Key Vault",
			ShortDescription: `Gets data contained in an existing Key Vault Certificate.`,
			Description: `

Use this data source to access data stored in an existing Key Vault Certificate.

~> **Note:** All arguments including the secret value will be stored in the raw state as plain-text.
[Read more about sensitive data in state](/docs/state/sensitive-data.html).

~> **Note:** This data source uses the ` + "`" + `GetSecret` + "`" + ` function of the Azure API, to get the key of the certificate. Therefore you need secret/get permission

`,
			Keywords: []string{
				"key",
				"vault",
				"certificate",
				"data",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the Key Vault Secret.`,
				},
				resource.Attribute{
					Name:        "key_vault_id",
					Description: `(Required) Specifies the ID of the Key Vault instance where the Secret resides, available on the ` + "`" + `azurerm_key_vault` + "`" + ` Data Source / Resource.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) Specifies the version of the certificate to look up. (Defaults to latest) ~>`,
				},
				resource.Attribute{
					Name:        "hex",
					Description: `The raw Key Vault Certificate data represented as a hexadecimal string.`,
				},
				resource.Attribute{
					Name:        "pem",
					Description: `The Key Vault Certificate in PEM format.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The Key Vault Certificate Key.`,
				},
				resource.Attribute{
					Name:        "expires",
					Description: `Expiry date of certificate in RFC3339 format.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the resource. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Key Vault Certificate.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "hex",
					Description: `The raw Key Vault Certificate data represented as a hexadecimal string.`,
				},
				resource.Attribute{
					Name:        "pem",
					Description: `The Key Vault Certificate in PEM format.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The Key Vault Certificate Key.`,
				},
				resource.Attribute{
					Name:        "expires",
					Description: `Expiry date of certificate in RFC3339 format.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the resource. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Key Vault Certificate.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_key_vault_certificate_issuer",
			Category:         "Key Vault",
			ShortDescription: `Gets information about an existing Key Vault Certificate Issuer.`,
			Description: `

Use this data source to access information about an existing Key Vault Certificate Issuer.

`,
			Keywords: []string{
				"key",
				"vault",
				"certificate",
				"issuer",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Key Vault Certificate Issuer.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account number with the third-party Certificate Issuer.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `The organization ID with the third-party Certificate Issuer.`,
				},
				resource.Attribute{
					Name:        "admin",
					Description: `A list of ` + "`" + `admin` + "`" + ` blocks as defined below.`,
				},
				resource.Attribute{
					Name:        "provider_name",
					Description: `The name of the third-party Certificate Issuer. --- An ` + "`" + `admin` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "email_address",
					Description: `E-mail address of the admin.`,
				},
				resource.Attribute{
					Name:        "first_name",
					Description: `First name of the admin.`,
				},
				resource.Attribute{
					Name:        "last_name",
					Description: `Last name of the admin.`,
				},
				resource.Attribute{
					Name:        "phone",
					Description: `Phone number of the admin. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Key Vault Certificate Issuer.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_key_vault_key",
			Category:         "Key Vault",
			ShortDescription: `Gets information about an existing Key Vault Key.`,
			Description: `

Use this data source to access information about an existing Key Vault Key.

~> **Note:** All arguments including the secret value will be stored in the raw state as plain-text.
[Read more about sensitive data in state](/docs/state/sensitive-data.html).

`,
			Keywords: []string{
				"key",
				"vault",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name of the Key Vault Key.`,
				},
				resource.Attribute{
					Name:        "key_vault_id",
					Description: `Specifies the ID of the Key Vault instance where the Secret resides, available on the ` + "`" + `azurerm_key_vault` + "`" + ` Data Source / Resource.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Key Vault Key.`,
				},
				resource.Attribute{
					Name:        "e",
					Description: `The RSA public exponent of this Key Vault Key.`,
				},
				resource.Attribute{
					Name:        "key_type",
					Description: `Specifies the Key Type of this Key Vault Key`,
				},
				resource.Attribute{
					Name:        "key_size",
					Description: `Specifies the Size of this Key Vault Key.`,
				},
				resource.Attribute{
					Name:        "key_opts",
					Description: `A list of JSON web key operations assigned to this Key Vault Key`,
				},
				resource.Attribute{
					Name:        "n",
					Description: `The RSA modulus of this Key Vault Key.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to this Key Vault Key.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The current version of the Key Vault Key.`,
				},
				resource.Attribute{
					Name:        "versionless_id",
					Description: `The Base ID of the Key Vault Key. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Key Vault Key.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Key Vault Key.`,
				},
				resource.Attribute{
					Name:        "e",
					Description: `The RSA public exponent of this Key Vault Key.`,
				},
				resource.Attribute{
					Name:        "key_type",
					Description: `Specifies the Key Type of this Key Vault Key`,
				},
				resource.Attribute{
					Name:        "key_size",
					Description: `Specifies the Size of this Key Vault Key.`,
				},
				resource.Attribute{
					Name:        "key_opts",
					Description: `A list of JSON web key operations assigned to this Key Vault Key`,
				},
				resource.Attribute{
					Name:        "n",
					Description: `The RSA modulus of this Key Vault Key.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to this Key Vault Key.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The current version of the Key Vault Key.`,
				},
				resource.Attribute{
					Name:        "versionless_id",
					Description: `The Base ID of the Key Vault Key. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Key Vault Key.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_key_vault_secret",
			Category:         "Key Vault",
			ShortDescription: `Gets information about an existing Key Vault Secret.`,
			Description: `

Use this data source to access information about an existing Key Vault Secret.

~> **Note:** All arguments including the secret value will be stored in the raw state as plain-text.
[Read more about sensitive data in state](/docs/state/sensitive-data.html).

`,
			Keywords: []string{
				"key",
				"vault",
				"secret",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name of the Key Vault Secret.`,
				},
				resource.Attribute{
					Name:        "key_vault_id",
					Description: `Specifies the ID of the Key Vault instance where the Secret resides, available on the ` + "`" + `azurerm_key_vault` + "`" + ` Data Source / Resource.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Key Vault Secret ID.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the Key Vault Secret.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The current version of the Key Vault Secret.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `The content type for the Key Vault Secret.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Any tags assigned to this resource. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Key Vault Secret.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Key Vault Secret ID.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the Key Vault Secret.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The current version of the Key Vault Secret.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `The content type for the Key Vault Secret.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Any tags assigned to this resource. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Key Vault Secret.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_kubernetes_cluster",
			Category:         "Container",
			ShortDescription: `Gets information about an existing Managed Kubernetes Cluster (AKS)`,
			Description: `

Use this data source to access information about an existing Managed Kubernetes Cluster (AKS).

~> **Note:** All arguments including the client secret will be stored in the raw state as plain-text.
[Read more about sensitive data in state](/docs/state/sensitive-data.html).

`,
			Keywords: []string{
				"container",
				"kubernetes",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the managed Kubernetes Cluster.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of the Resource Group in which the managed Kubernetes Cluster exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Kubernetes Managed Cluster.`,
				},
				resource.Attribute{
					Name:        "api_server_authorized_ip_ranges",
					Description: `The IP ranges to whitelist for incoming traffic to the primaries. ->`,
				},
				resource.Attribute{
					Name:        "addon_profile",
					Description: `A ` + "`" + `addon_profile` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "agent_pool_profile",
					Description: `An ` + "`" + `agent_pool_profile` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "dns_prefix",
					Description: `The DNS Prefix of the managed Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `The FQDN of the Azure Kubernetes Managed Cluster.`,
				},
				resource.Attribute{
					Name:        "private_fqdn",
					Description: `The FQDN of this Kubernetes Cluster when private link has been enabled. This name is only resolvable inside the Virtual Network where the Azure Kubernetes Service is located ->`,
				},
				resource.Attribute{
					Name:        "kube_admin_config",
					Description: `A ` + "`" + `kube_admin_config` + "`" + ` block as defined below. This is only available when Role Based Access Control with Azure Active Directory is enabled.`,
				},
				resource.Attribute{
					Name:        "kube_admin_config_raw",
					Description: `Raw Kubernetes config for the admin account to be used by [kubectl](https://kubernetes.io/docs/reference/kubectl/overview/) and other compatible tools. This is only available when Role Based Access Control with Azure Active Directory is enabled.`,
				},
				resource.Attribute{
					Name:        "kube_config",
					Description: `A ` + "`" + `kube_config` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "kube_config_raw",
					Description: `Base64 encoded Kubernetes configuration.`,
				},
				resource.Attribute{
					Name:        "kubernetes_version",
					Description: `The version of Kubernetes used on the managed Kubernetes Cluster.`,
				},
				resource.Attribute{
					Name:        "private_cluster_enabled",
					Description: `If the cluster has the Kubernetes API only exposed on internal IP addresses.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure Region in which the managed Kubernetes Cluster exists.`,
				},
				resource.Attribute{
					Name:        "disk_encryption_set_id",
					Description: `The ID of the Disk Encryption Set used for the Nodes and Volumes.`,
				},
				resource.Attribute{
					Name:        "linux_profile",
					Description: `A ` + "`" + `linux_profile` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "windows_profile",
					Description: `A ` + "`" + `windows_profile` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "network_profile",
					Description: `A ` + "`" + `network_profile` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "node_resource_group",
					Description: `Auto-generated Resource Group containing AKS Cluster resources.`,
				},
				resource.Attribute{
					Name:        "role_based_access_control",
					Description: `A ` + "`" + `role_based_access_control` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "service_principal",
					Description: `A ` + "`" + `service_principal` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "identity",
					Description: `A ` + "`" + `identity` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "kubelet_identity",
					Description: `A ` + "`" + `kubelet_identity` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to this resource. --- A ` + "`" + `addon_profile` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "http_application_routing",
					Description: `A ` + "`" + `http_application_routing` + "`" + ` block.`,
				},
				resource.Attribute{
					Name:        "oms_agent",
					Description: `A ` + "`" + `oms_agent` + "`" + ` block.`,
				},
				resource.Attribute{
					Name:        "kube_dashboard",
					Description: `A ` + "`" + `kube_dashboard` + "`" + ` block.`,
				},
				resource.Attribute{
					Name:        "azure_policy",
					Description: `A ` + "`" + `azure_policy` + "`" + ` block. --- A ` + "`" + `agent_pool_profile` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the Agent Pool.`,
				},
				resource.Attribute{
					Name:        "count",
					Description: `The number of Agents (VM's) in the Pool.`,
				},
				resource.Attribute{
					Name:        "max_pods",
					Description: `The maximum number of pods that can run on each agent.`,
				},
				resource.Attribute{
					Name:        "availability_zones",
					Description: `The availability zones used for the nodes.`,
				},
				resource.Attribute{
					Name:        "enable_auto_scaling",
					Description: `If the auto-scaler is enabled.`,
				},
				resource.Attribute{
					Name:        "min_count",
					Description: `Minimum number of nodes for auto-scaling`,
				},
				resource.Attribute{
					Name:        "max_count",
					Description: `Maximum number of nodes for auto-scaling`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name assigned to this pool of agents.`,
				},
				resource.Attribute{
					Name:        "os_disk_size_gb",
					Description: `The size of the Agent VM's Operating System Disk in GB.`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `The Operating System used for the Agents.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "orchestrator_version",
					Description: `Kubernetes version used for the Agents.`,
				},
				resource.Attribute{
					Name:        "upgrade_settings",
					Description: `A ` + "`" + `upgrade_settings` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "vm_size",
					Description: `The size of each VM in the Agent Pool (e.g. ` + "`" + `Standard_F1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "vnet_subnet_id",
					Description: `The ID of the Subnet where the Agents in the Pool are provisioned. --- A ` + "`" + `upgrade_settings` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "max_surge",
					Description: `The maximum number or percentage of nodes which will be added to the Node Pool size during an upgrade. --- A ` + "`" + `azure_active_directory` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "admin_group_object_ids",
					Description: `The list of Object IDs of Azure Active Directory Groups which have Admin Role on the Cluster (when using a Managed integration).`,
				},
				resource.Attribute{
					Name:        "client_app_id",
					Description: `The Client ID of an Azure Active Directory Application.`,
				},
				resource.Attribute{
					Name:        "managed",
					Description: `Is the Azure Active Directory Integration managed (also known as AAD Integration V2)?`,
				},
				resource.Attribute{
					Name:        "server_app_id",
					Description: `The Server ID of an Azure Active Directory Application.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `The Tenant ID used for Azure Active Directory Application. --- A ` + "`" + `http_application_routing` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Is HTTP Application Routing Enabled?`,
				},
				resource.Attribute{
					Name:        "http_application_routing_zone_name",
					Description: `The Zone Name of the HTTP Application Routing. --- The ` + "`" + `kube_admin_config` + "`" + ` and ` + "`" + `kube_config` + "`" + ` blocks exports the following:`,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: `Base64 encoded private key used by clients to authenticate to the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "client_certificate",
					Description: `Base64 encoded public certificate used by clients to authenticate to the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_ca_certificate",
					Description: `Base64 encoded public CA certificate used as the root of trust for the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `The Kubernetes cluster server host.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `A username used to authenticate to the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `A password or token used to authenticate to the Kubernetes cluster. ->`,
				},
				resource.Attribute{
					Name:        "admin_username",
					Description: `The username associated with the administrator account of the managed Kubernetes Cluster.`,
				},
				resource.Attribute{
					Name:        "ssh_key",
					Description: `An ` + "`" + `ssh_key` + "`" + ` block as defined below. --- A ` + "`" + `windows_profile` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "admin_username",
					Description: `The username associated with the administrator account of the Windows VMs. --- A ` + "`" + `network_profile` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "docker_bridge_cidr",
					Description: `IP address (in CIDR notation) used as the Docker bridge IP address on nodes.`,
				},
				resource.Attribute{
					Name:        "dns_service_ip",
					Description: `IP address within the Kubernetes service address range used by cluster service discovery (kube-dns).`,
				},
				resource.Attribute{
					Name:        "network_plugin",
					Description: `Network plugin used such as ` + "`" + `azure` + "`" + ` or ` + "`" + `kubenet` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "network_policy",
					Description: `Network policy to be used with Azure CNI. Eg: ` + "`" + `calico` + "`" + ` or ` + "`" + `azure` + "`" + ``,
				},
				resource.Attribute{
					Name:        "network_mode",
					Description: `Network mode to be used with Azure CNI. Eg: ` + "`" + `bridge` + "`" + ` or ` + "`" + `transparent` + "`" + ` ->`,
				},
				resource.Attribute{
					Name:        "pod_cidr",
					Description: `The CIDR used for pod IP addresses.`,
				},
				resource.Attribute{
					Name:        "service_cidr",
					Description: `Network range used by the Kubernetes service. --- A ` + "`" + `oms_agent` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Is the OMS Agent Enabled?`,
				},
				resource.Attribute{
					Name:        "log_analytics_workspace_id",
					Description: `The ID of the Log Analytics Workspace which the OMS Agent should send data to.`,
				},
				resource.Attribute{
					Name:        "oms_agent_identity",
					Description: `An ` + "`" + `oms_agent_identity` + "`" + ` block as defined below. --- The ` + "`" + `oms_agent_identity` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `The Client ID of the user-defined Managed Identity used by the OMS Agents.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The Object ID of the user-defined Managed Identity used by the OMS Agents.`,
				},
				resource.Attribute{
					Name:        "user_assigned_identity_id",
					Description: `The ID of the User Assigned Identity used by the OMS Agents. --- A ` + "`" + `kube_dashboard` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Is the Kubernetes Dashboard enabled? --- A ` + "`" + `azure_policy` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Is Azure Policy for Kubernetes enabled? --- A ` + "`" + `role_based_access_control` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "azure_active_directory",
					Description: `A ` + "`" + `azure_active_directory` + "`" + ` block as documented above.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Is Role Based Access Control enabled? --- A ` + "`" + `service_principal` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `The Client ID of the Service Principal used by this Managed Kubernetes Cluster. --- The ` + "`" + `identity` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of identity used for the managed cluster.`,
				},
				resource.Attribute{
					Name:        "principal_id",
					Description: `The principal id of the system assigned identity which is used by primary components.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `The tenant id of the system assigned identity which is used by primary components. --- The ` + "`" + `kubelet_identity` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `The Client ID of the user-defined Managed Identity assigned to the Kubelets.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The Object ID of the user-defined Managed Identity assigned to the Kubelets.`,
				},
				resource.Attribute{
					Name:        "user_assigned_identity_id",
					Description: `The ID of the User Assigned Identity assigned to the Kubelets. --- A ` + "`" + `ssh_key` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "key_data",
					Description: `The Public SSH Key used to access the cluster. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Managed Kubernetes Cluster (AKS).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Kubernetes Managed Cluster.`,
				},
				resource.Attribute{
					Name:        "api_server_authorized_ip_ranges",
					Description: `The IP ranges to whitelist for incoming traffic to the primaries. ->`,
				},
				resource.Attribute{
					Name:        "addon_profile",
					Description: `A ` + "`" + `addon_profile` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "agent_pool_profile",
					Description: `An ` + "`" + `agent_pool_profile` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "dns_prefix",
					Description: `The DNS Prefix of the managed Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `The FQDN of the Azure Kubernetes Managed Cluster.`,
				},
				resource.Attribute{
					Name:        "private_fqdn",
					Description: `The FQDN of this Kubernetes Cluster when private link has been enabled. This name is only resolvable inside the Virtual Network where the Azure Kubernetes Service is located ->`,
				},
				resource.Attribute{
					Name:        "kube_admin_config",
					Description: `A ` + "`" + `kube_admin_config` + "`" + ` block as defined below. This is only available when Role Based Access Control with Azure Active Directory is enabled.`,
				},
				resource.Attribute{
					Name:        "kube_admin_config_raw",
					Description: `Raw Kubernetes config for the admin account to be used by [kubectl](https://kubernetes.io/docs/reference/kubectl/overview/) and other compatible tools. This is only available when Role Based Access Control with Azure Active Directory is enabled.`,
				},
				resource.Attribute{
					Name:        "kube_config",
					Description: `A ` + "`" + `kube_config` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "kube_config_raw",
					Description: `Base64 encoded Kubernetes configuration.`,
				},
				resource.Attribute{
					Name:        "kubernetes_version",
					Description: `The version of Kubernetes used on the managed Kubernetes Cluster.`,
				},
				resource.Attribute{
					Name:        "private_cluster_enabled",
					Description: `If the cluster has the Kubernetes API only exposed on internal IP addresses.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure Region in which the managed Kubernetes Cluster exists.`,
				},
				resource.Attribute{
					Name:        "disk_encryption_set_id",
					Description: `The ID of the Disk Encryption Set used for the Nodes and Volumes.`,
				},
				resource.Attribute{
					Name:        "linux_profile",
					Description: `A ` + "`" + `linux_profile` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "windows_profile",
					Description: `A ` + "`" + `windows_profile` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "network_profile",
					Description: `A ` + "`" + `network_profile` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "node_resource_group",
					Description: `Auto-generated Resource Group containing AKS Cluster resources.`,
				},
				resource.Attribute{
					Name:        "role_based_access_control",
					Description: `A ` + "`" + `role_based_access_control` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "service_principal",
					Description: `A ` + "`" + `service_principal` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "identity",
					Description: `A ` + "`" + `identity` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "kubelet_identity",
					Description: `A ` + "`" + `kubelet_identity` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to this resource. --- A ` + "`" + `addon_profile` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "http_application_routing",
					Description: `A ` + "`" + `http_application_routing` + "`" + ` block.`,
				},
				resource.Attribute{
					Name:        "oms_agent",
					Description: `A ` + "`" + `oms_agent` + "`" + ` block.`,
				},
				resource.Attribute{
					Name:        "kube_dashboard",
					Description: `A ` + "`" + `kube_dashboard` + "`" + ` block.`,
				},
				resource.Attribute{
					Name:        "azure_policy",
					Description: `A ` + "`" + `azure_policy` + "`" + ` block. --- A ` + "`" + `agent_pool_profile` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the Agent Pool.`,
				},
				resource.Attribute{
					Name:        "count",
					Description: `The number of Agents (VM's) in the Pool.`,
				},
				resource.Attribute{
					Name:        "max_pods",
					Description: `The maximum number of pods that can run on each agent.`,
				},
				resource.Attribute{
					Name:        "availability_zones",
					Description: `The availability zones used for the nodes.`,
				},
				resource.Attribute{
					Name:        "enable_auto_scaling",
					Description: `If the auto-scaler is enabled.`,
				},
				resource.Attribute{
					Name:        "min_count",
					Description: `Minimum number of nodes for auto-scaling`,
				},
				resource.Attribute{
					Name:        "max_count",
					Description: `Maximum number of nodes for auto-scaling`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name assigned to this pool of agents.`,
				},
				resource.Attribute{
					Name:        "os_disk_size_gb",
					Description: `The size of the Agent VM's Operating System Disk in GB.`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `The Operating System used for the Agents.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "orchestrator_version",
					Description: `Kubernetes version used for the Agents.`,
				},
				resource.Attribute{
					Name:        "upgrade_settings",
					Description: `A ` + "`" + `upgrade_settings` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "vm_size",
					Description: `The size of each VM in the Agent Pool (e.g. ` + "`" + `Standard_F1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "vnet_subnet_id",
					Description: `The ID of the Subnet where the Agents in the Pool are provisioned. --- A ` + "`" + `upgrade_settings` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "max_surge",
					Description: `The maximum number or percentage of nodes which will be added to the Node Pool size during an upgrade. --- A ` + "`" + `azure_active_directory` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "admin_group_object_ids",
					Description: `The list of Object IDs of Azure Active Directory Groups which have Admin Role on the Cluster (when using a Managed integration).`,
				},
				resource.Attribute{
					Name:        "client_app_id",
					Description: `The Client ID of an Azure Active Directory Application.`,
				},
				resource.Attribute{
					Name:        "managed",
					Description: `Is the Azure Active Directory Integration managed (also known as AAD Integration V2)?`,
				},
				resource.Attribute{
					Name:        "server_app_id",
					Description: `The Server ID of an Azure Active Directory Application.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `The Tenant ID used for Azure Active Directory Application. --- A ` + "`" + `http_application_routing` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Is HTTP Application Routing Enabled?`,
				},
				resource.Attribute{
					Name:        "http_application_routing_zone_name",
					Description: `The Zone Name of the HTTP Application Routing. --- The ` + "`" + `kube_admin_config` + "`" + ` and ` + "`" + `kube_config` + "`" + ` blocks exports the following:`,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: `Base64 encoded private key used by clients to authenticate to the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "client_certificate",
					Description: `Base64 encoded public certificate used by clients to authenticate to the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_ca_certificate",
					Description: `Base64 encoded public CA certificate used as the root of trust for the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `The Kubernetes cluster server host.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `A username used to authenticate to the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `A password or token used to authenticate to the Kubernetes cluster. ->`,
				},
				resource.Attribute{
					Name:        "admin_username",
					Description: `The username associated with the administrator account of the managed Kubernetes Cluster.`,
				},
				resource.Attribute{
					Name:        "ssh_key",
					Description: `An ` + "`" + `ssh_key` + "`" + ` block as defined below. --- A ` + "`" + `windows_profile` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "admin_username",
					Description: `The username associated with the administrator account of the Windows VMs. --- A ` + "`" + `network_profile` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "docker_bridge_cidr",
					Description: `IP address (in CIDR notation) used as the Docker bridge IP address on nodes.`,
				},
				resource.Attribute{
					Name:        "dns_service_ip",
					Description: `IP address within the Kubernetes service address range used by cluster service discovery (kube-dns).`,
				},
				resource.Attribute{
					Name:        "network_plugin",
					Description: `Network plugin used such as ` + "`" + `azure` + "`" + ` or ` + "`" + `kubenet` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "network_policy",
					Description: `Network policy to be used with Azure CNI. Eg: ` + "`" + `calico` + "`" + ` or ` + "`" + `azure` + "`" + ``,
				},
				resource.Attribute{
					Name:        "network_mode",
					Description: `Network mode to be used with Azure CNI. Eg: ` + "`" + `bridge` + "`" + ` or ` + "`" + `transparent` + "`" + ` ->`,
				},
				resource.Attribute{
					Name:        "pod_cidr",
					Description: `The CIDR used for pod IP addresses.`,
				},
				resource.Attribute{
					Name:        "service_cidr",
					Description: `Network range used by the Kubernetes service. --- A ` + "`" + `oms_agent` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Is the OMS Agent Enabled?`,
				},
				resource.Attribute{
					Name:        "log_analytics_workspace_id",
					Description: `The ID of the Log Analytics Workspace which the OMS Agent should send data to.`,
				},
				resource.Attribute{
					Name:        "oms_agent_identity",
					Description: `An ` + "`" + `oms_agent_identity` + "`" + ` block as defined below. --- The ` + "`" + `oms_agent_identity` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `The Client ID of the user-defined Managed Identity used by the OMS Agents.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The Object ID of the user-defined Managed Identity used by the OMS Agents.`,
				},
				resource.Attribute{
					Name:        "user_assigned_identity_id",
					Description: `The ID of the User Assigned Identity used by the OMS Agents. --- A ` + "`" + `kube_dashboard` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Is the Kubernetes Dashboard enabled? --- A ` + "`" + `azure_policy` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Is Azure Policy for Kubernetes enabled? --- A ` + "`" + `role_based_access_control` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "azure_active_directory",
					Description: `A ` + "`" + `azure_active_directory` + "`" + ` block as documented above.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Is Role Based Access Control enabled? --- A ` + "`" + `service_principal` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `The Client ID of the Service Principal used by this Managed Kubernetes Cluster. --- The ` + "`" + `identity` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of identity used for the managed cluster.`,
				},
				resource.Attribute{
					Name:        "principal_id",
					Description: `The principal id of the system assigned identity which is used by primary components.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `The tenant id of the system assigned identity which is used by primary components. --- The ` + "`" + `kubelet_identity` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `The Client ID of the user-defined Managed Identity assigned to the Kubelets.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The Object ID of the user-defined Managed Identity assigned to the Kubelets.`,
				},
				resource.Attribute{
					Name:        "user_assigned_identity_id",
					Description: `The ID of the User Assigned Identity assigned to the Kubelets. --- A ` + "`" + `ssh_key` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "key_data",
					Description: `The Public SSH Key used to access the cluster. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Managed Kubernetes Cluster (AKS).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_kubernetes_cluster_node_pool",
			Category:         "Container",
			ShortDescription: `Gets information about an existing Kubernetes Cluster Node Pool.`,
			Description: `

Use this data source to access information about an existing Kubernetes Cluster Node Pool.

`,
			Keywords: []string{
				"container",
				"kubernetes",
				"cluster",
				"node",
				"pool",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Kubernetes Cluster Node Pool.`,
				},
				resource.Attribute{
					Name:        "availability_zones",
					Description: `A list of Availability Zones in which the Nodes in this Node Pool exists.`,
				},
				resource.Attribute{
					Name:        "enable_auto_scaling",
					Description: `Does this Node Pool have Auto-Scaling enabled?`,
				},
				resource.Attribute{
					Name:        "enable_node_public_ip",
					Description: `Do nodes in this Node Pool have a Public IP Address?`,
				},
				resource.Attribute{
					Name:        "eviction_policy",
					Description: `The eviction policy used for Virtual Machines in the Virtual Machine Scale Set, when ` + "`" + `priority` + "`" + ` is set to ` + "`" + `Spot` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max_count",
					Description: `The maximum number of Nodes allowed when auto-scaling is enabled.`,
				},
				resource.Attribute{
					Name:        "max_pods",
					Description: `The maximum number of Pods allowed on each Node in this Node Pool.`,
				},
				resource.Attribute{
					Name:        "min_count",
					Description: `The minimum number of Nodes allowed when auto-scaling is enabled.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `The Mode for this Node Pool, specifying how these Nodes should be used (for either System or User resources).`,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `The current number of Nodes in the Node Pool.`,
				},
				resource.Attribute{
					Name:        "node_labels",
					Description: `A map of Kubernetes Labels applied to each Node in this Node Pool.`,
				},
				resource.Attribute{
					Name:        "node_taints",
					Description: `A map of Kubernetes Taints applied to each Node in this Node Pool.`,
				},
				resource.Attribute{
					Name:        "orchestrator_version",
					Description: `The version of Kubernetes configured on each Node in this Node Pool.`,
				},
				resource.Attribute{
					Name:        "os_disk_size_gb",
					Description: `The size of the OS Disk on each Node in this Node Pool.`,
				},
				resource.Attribute{
					Name:        "os_disk_type",
					Description: `The type of the OS Disk on each Node in this Node Pool.`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `The operating system used on each Node in this Node Pool.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority of the Virtual Machines in the Virtual Machine Scale Set backing this Node Pool.`,
				},
				resource.Attribute{
					Name:        "proximity_placement_group_id",
					Description: `The ID of the Proximity Placement Group where the Virtual Machine Scale Set backing this Node Pool will be placed.`,
				},
				resource.Attribute{
					Name:        "spot_max_price",
					Description: `The maximum price being paid for Virtual Machines in this Scale Set. ` + "`" + `-1` + "`" + ` means the current on-demand price for a Virtual Machine.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the Kubernetes Cluster Node Pool.`,
				},
				resource.Attribute{
					Name:        "upgrade_settings",
					Description: `A ` + "`" + `upgrade_settings` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "vm_size",
					Description: `The size of the Virtual Machines used in the Virtual Machine Scale Set backing this Node Pool.`,
				},
				resource.Attribute{
					Name:        "vnet_subnet_id",
					Description: `The ID of the Subnet in which this Node Pool exists. --- A ` + "`" + `upgrade_settings` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "max_surge",
					Description: `The maximum number or percentage of nodes which will be added to the Node Pool size during an upgrade. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Kubernetes Cluster Node Pool.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_kubernetes_service_versions",
			Category:         "Container",
			ShortDescription: `Gets the available versions of Kubernetes supported by Azure Kubernetes Service.`,
			Description: `

Use this data source to retrieve the version of Kubernetes supported by Azure Kubernetes Service.

`,
			Keywords: []string{
				"container",
				"kubernetes",
				"service",
				"versions",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "location",
					Description: `Specifies the location in which to query for versions.`,
				},
				resource.Attribute{
					Name:        "version_prefix",
					Description: `(Optional) A prefix filter for the versions of Kubernetes which should be returned; for example ` + "`" + `1.` + "`" + ` will return ` + "`" + `1.9` + "`" + ` to ` + "`" + `1.14` + "`" + `, whereas ` + "`" + `1.12` + "`" + ` will return ` + "`" + `1.12.2` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "include_preview",
					Description: `(Optional) Should Preview versions of Kubernetes in AKS be included? Defaults to ` + "`" + `true` + "`" + ` ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "versions",
					Description: `The list of all supported versions.`,
				},
				resource.Attribute{
					Name:        "latest_version",
					Description: `The most recent version available. If ` + "`" + `include_preview == false` + "`" + `, this is the most recent non-preview version available. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the versions.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "versions",
					Description: `The list of all supported versions.`,
				},
				resource.Attribute{
					Name:        "latest_version",
					Description: `The most recent version available. If ` + "`" + `include_preview == false` + "`" + `, this is the most recent non-preview version available. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the versions.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_kusto_cluster",
			Category:         "Data Explorer",
			ShortDescription: `Get information about an existing Kusto (also known as Azure Data Explorer) Cluster`,
			Description: `

Use this data source to access information about an existing Kusto (also known as Azure Data Explorer) Cluster

`,
			Keywords: []string{
				"data",
				"explorer",
				"kusto",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name of the Kusto Cluster.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of the Resource Group where the Kusto Cluster exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Kusto Cluster ID.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `The FQDN of the Azure Kusto Cluster.`,
				},
				resource.Attribute{
					Name:        "data_ingestion_uri",
					Description: `The Kusto Cluster URI to be used for data ingestion. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Kusto Cluster.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Kusto Cluster ID.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `The FQDN of the Azure Kusto Cluster.`,
				},
				resource.Attribute{
					Name:        "data_ingestion_uri",
					Description: `The Kusto Cluster URI to be used for data ingestion. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Kusto Cluster.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_lb",
			Category:         "Network",
			ShortDescription: `Get information about an existing Load Balancer`,
			Description: `

Use this data source to access information about an existing Load Balancer

`,
			Icon: "Networking/10062-icon-service-Load-Balancers.svg",
			Keywords: []string{
				"network",
				"lb",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name of the Load Balancer.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of the Resource Group in which the Load Balancer exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Load Balancer.`,
				},
				resource.Attribute{
					Name:        "frontend_ip_configuration",
					Description: `A ` + "`" + `frontend_ip_configuration` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the Load Balancer exists.`,
				},
				resource.Attribute{
					Name:        "private_ip_address",
					Description: `The first private IP address assigned to the load balancer in ` + "`" + `frontend_ip_configuration` + "`" + ` blocks, if any.`,
				},
				resource.Attribute{
					Name:        "private_ip_addresses",
					Description: `The list of private IP address assigned to the load balancer in ` + "`" + `frontend_ip_configuration` + "`" + ` blocks, if any.`,
				},
				resource.Attribute{
					Name:        "sku",
					Description: `The SKU of the Load Balancer.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource. --- A ` + "`" + `frontend_ip_configuration` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Frontend IP Configuration.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the Frontend IP Configuration.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet which is associated with the IP Configuration.`,
				},
				resource.Attribute{
					Name:        "private_ip_address",
					Description: `Private IP Address to assign to the Load Balancer.`,
				},
				resource.Attribute{
					Name:        "private_ip_address_allocation",
					Description: `The allocation method for the Private IP Address used by this Load Balancer.`,
				},
				resource.Attribute{
					Name:        "private_ip_address_version",
					Description: `The Private IP Address Version, either ` + "`" + `IPv4` + "`" + ` or ` + "`" + `IPv6` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "public_ip_address_id",
					Description: `The ID of a Public IP Address which is associated with this Load Balancer.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of Availability Zones which the Load Balancer's IP Addresses should be created in. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Load Balancer.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Load Balancer.`,
				},
				resource.Attribute{
					Name:        "frontend_ip_configuration",
					Description: `A ` + "`" + `frontend_ip_configuration` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the Load Balancer exists.`,
				},
				resource.Attribute{
					Name:        "private_ip_address",
					Description: `The first private IP address assigned to the load balancer in ` + "`" + `frontend_ip_configuration` + "`" + ` blocks, if any.`,
				},
				resource.Attribute{
					Name:        "private_ip_addresses",
					Description: `The list of private IP address assigned to the load balancer in ` + "`" + `frontend_ip_configuration` + "`" + ` blocks, if any.`,
				},
				resource.Attribute{
					Name:        "sku",
					Description: `The SKU of the Load Balancer.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource. --- A ` + "`" + `frontend_ip_configuration` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Frontend IP Configuration.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the Frontend IP Configuration.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet which is associated with the IP Configuration.`,
				},
				resource.Attribute{
					Name:        "private_ip_address",
					Description: `Private IP Address to assign to the Load Balancer.`,
				},
				resource.Attribute{
					Name:        "private_ip_address_allocation",
					Description: `The allocation method for the Private IP Address used by this Load Balancer.`,
				},
				resource.Attribute{
					Name:        "private_ip_address_version",
					Description: `The Private IP Address Version, either ` + "`" + `IPv4` + "`" + ` or ` + "`" + `IPv6` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "public_ip_address_id",
					Description: `The ID of a Public IP Address which is associated with this Load Balancer.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of Availability Zones which the Load Balancer's IP Addresses should be created in. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Load Balancer.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_lb_backend_address_pool",
			Category:         "Network",
			ShortDescription: `Get information about an existing Load Balancer Backend Address Pool`,
			Description: `

Use this data source to access information about an existing Load Balancer's Backend Address Pool.

`,
			Keywords: []string{
				"network",
				"lb",
				"backend",
				"address",
				"pool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name of the Backend Address Pool.`,
				},
				resource.Attribute{
					Name:        "loadbalancer_id",
					Description: `The ID of the Load Balancer in which the Backend Address Pool exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Backend Address Pool.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Backend Address Pool.`,
				},
				resource.Attribute{
					Name:        "backend_address",
					Description: `A list of ` + "`" + `backend_address` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "backend_ip_configurations",
					Description: `A list of references to IP addresses defined in network interfaces.`,
				},
				resource.Attribute{
					Name:        "load_balancing_rules",
					Description: `A list of the Load Balancing Rules associated with this Backend Address Pool.`,
				},
				resource.Attribute{
					Name:        "outbound_rules",
					Description: `A list of the Load Balancing Outbound Rules associated with this Backend Address Pool. --- A ` + "`" + `backend_address` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Backend Address.`,
				},
				resource.Attribute{
					Name:        "virtual_network_id",
					Description: `The ID of the Virtual Network where the Backend Address of the Load Balancer exists.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The Static IP address for this Load Balancer within the Virtual Network. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Backend Address Pool.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Backend Address Pool.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Backend Address Pool.`,
				},
				resource.Attribute{
					Name:        "backend_address",
					Description: `A list of ` + "`" + `backend_address` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "backend_ip_configurations",
					Description: `A list of references to IP addresses defined in network interfaces.`,
				},
				resource.Attribute{
					Name:        "load_balancing_rules",
					Description: `A list of the Load Balancing Rules associated with this Backend Address Pool.`,
				},
				resource.Attribute{
					Name:        "outbound_rules",
					Description: `A list of the Load Balancing Outbound Rules associated with this Backend Address Pool. --- A ` + "`" + `backend_address` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Backend Address.`,
				},
				resource.Attribute{
					Name:        "virtual_network_id",
					Description: `The ID of the Virtual Network where the Backend Address of the Load Balancer exists.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The Static IP address for this Load Balancer within the Virtual Network. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Backend Address Pool.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_lb_rule",
			Category:         "Load Balancer",
			ShortDescription: `Gets information about an existing Load Balancer Rule.`,
			Description: `

Use this data source to access information about an existing Load Balancer Rule.

`,
			Keywords: []string{
				"load",
				"balancer",
				"lb",
				"rule",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Load Balancer Rule.`,
				},
				resource.Attribute{
					Name:        "backend_address_pool_id",
					Description: `A reference to a Backend Address Pool over which this Load Balancing Rule operates.`,
				},
				resource.Attribute{
					Name:        "probe_id",
					Description: `A reference to a Probe used by this Load Balancing Rule.`,
				},
				resource.Attribute{
					Name:        "frontend_ip_configuration_name",
					Description: `The name of the frontend IP configuration to which the rule is associated.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The transport protocol for the external endpoint.`,
				},
				resource.Attribute{
					Name:        "frontend_port",
					Description: `The port for the external endpoint.`,
				},
				resource.Attribute{
					Name:        "backend_port",
					Description: `The port used for internal connections on the endpoint.`,
				},
				resource.Attribute{
					Name:        "enable_floating_ip",
					Description: `If Floating IPs are enabled for this Load Balancer Rule`,
				},
				resource.Attribute{
					Name:        "idle_timeout_in_minutes",
					Description: `Specifies the idle timeout in minutes for TCP connections.`,
				},
				resource.Attribute{
					Name:        "load_distribution",
					Description: `Specifies the load balancing distribution type used by the Load Balancer.`,
				},
				resource.Attribute{
					Name:        "disable_outbound_snat",
					Description: `If outbound SNAT is enabled for this Load Balancer Rule.`,
				},
				resource.Attribute{
					Name:        "enable_tcp_reset",
					Description: `If TCP Reset is enabled for this Load Balancer Rule. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Load Balancer Rule.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_log_analytics_workspace",
			Category:         "Log Analytics",
			ShortDescription: `Gets information about an existing Log Analytics (formally Operational Insights) Workspace.`,
			Description: `

Use this data source to access information about an existing Log Analytics (formally Operational Insights) Workspace.

`,
			Icon: "Analytics/00009-icon-service-Log-Analytics-Workspaces.svg",
			Keywords: []string{
				"log",
				"analytics",
				"workspace",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name of the Log Analytics Workspace.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of the resource group in which the Log Analytics workspace is located in. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Log Analytics Workspace.`,
				},
				resource.Attribute{
					Name:        "primary_shared_key",
					Description: `The Primary shared key for the Log Analytics Workspace.`,
				},
				resource.Attribute{
					Name:        "secondary_shared_key",
					Description: `The Secondary shared key for the Log Analytics Workspace.`,
				},
				resource.Attribute{
					Name:        "workspace_id",
					Description: `The Workspace (or Customer) ID for the Log Analytics Workspace.`,
				},
				resource.Attribute{
					Name:        "sku",
					Description: `The Sku of the Log Analytics Workspace.`,
				},
				resource.Attribute{
					Name:        "retention_in_days",
					Description: `The workspace data retention in days.`,
				},
				resource.Attribute{
					Name:        "daily_quota_gb",
					Description: `The workspace daily quota for ingestion in GB.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Workspace.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Log Analytics Workspace.`,
				},
				resource.Attribute{
					Name:        "primary_shared_key",
					Description: `The Primary shared key for the Log Analytics Workspace.`,
				},
				resource.Attribute{
					Name:        "secondary_shared_key",
					Description: `The Secondary shared key for the Log Analytics Workspace.`,
				},
				resource.Attribute{
					Name:        "workspace_id",
					Description: `The Workspace (or Customer) ID for the Log Analytics Workspace.`,
				},
				resource.Attribute{
					Name:        "sku",
					Description: `The Sku of the Log Analytics Workspace.`,
				},
				resource.Attribute{
					Name:        "retention_in_days",
					Description: `The workspace data retention in days.`,
				},
				resource.Attribute{
					Name:        "daily_quota_gb",
					Description: `The workspace daily quota for ingestion in GB.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Workspace.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_logic_app_integration_account",
			Category:         "Logic App",
			ShortDescription: `Gets information about an existing Logic App Integration Account.`,
			Description: `

Use this data source to access information about an existing Logic App Integration Account.

`,
			Keywords: []string{
				"logic",
				"app",
				"integration",
				"account",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Logic App Integration Account.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure Region where the Logic App Integration Account exists.`,
				},
				resource.Attribute{
					Name:        "sku_name",
					Description: `The sku name of the Logic App Integration Account.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the Logic App Integration Account. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Logic App Integration Account.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_logic_app_workflow",
			Category:         "Logic App",
			ShortDescription: `Gets information about an existing Logic App Workflow.`,
			Description: `

Use this data source to access information about an existing Logic App Workflow.

`,
			Keywords: []string{
				"logic",
				"app",
				"workflow",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Logic App Workflow.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of the Resource Group in which the Logic App Workflow exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Logic App Workflow ID.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the Logic App Workflow exists.`,
				},
				resource.Attribute{
					Name:        "logic_app_integration_account_id",
					Description: `The ID of the integration account linked by this Logic App Workflow.`,
				},
				resource.Attribute{
					Name:        "workflow_schema",
					Description: `The Schema used for this Logic App Workflow.`,
				},
				resource.Attribute{
					Name:        "workflow_version",
					Description: `The version of the Schema used for this Logic App Workflow. Defaults to ` + "`" + `1.0.0.0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `A map of Key-Value pairs.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource.`,
				},
				resource.Attribute{
					Name:        "access_endpoint",
					Description: `The Access Endpoint for the Logic App Workflow`,
				},
				resource.Attribute{
					Name:        "connector_endpoint_ip_addresses",
					Description: `The list of access endpoint ip addresses of connector.`,
				},
				resource.Attribute{
					Name:        "connector_outbound_ip_addresses",
					Description: `The list of outgoing ip addresses of connector.`,
				},
				resource.Attribute{
					Name:        "workflow_endpoint_ip_addresses",
					Description: `The list of access endpoint ip addresses of workflow.`,
				},
				resource.Attribute{
					Name:        "workflow_outbound_ip_addresses",
					Description: `The list of outgoing ip addresses of workflow. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Logic App Workflow.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Logic App Workflow ID.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the Logic App Workflow exists.`,
				},
				resource.Attribute{
					Name:        "logic_app_integration_account_id",
					Description: `The ID of the integration account linked by this Logic App Workflow.`,
				},
				resource.Attribute{
					Name:        "workflow_schema",
					Description: `The Schema used for this Logic App Workflow.`,
				},
				resource.Attribute{
					Name:        "workflow_version",
					Description: `The version of the Schema used for this Logic App Workflow. Defaults to ` + "`" + `1.0.0.0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `A map of Key-Value pairs.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource.`,
				},
				resource.Attribute{
					Name:        "access_endpoint",
					Description: `The Access Endpoint for the Logic App Workflow`,
				},
				resource.Attribute{
					Name:        "connector_endpoint_ip_addresses",
					Description: `The list of access endpoint ip addresses of connector.`,
				},
				resource.Attribute{
					Name:        "connector_outbound_ip_addresses",
					Description: `The list of outgoing ip addresses of connector.`,
				},
				resource.Attribute{
					Name:        "workflow_endpoint_ip_addresses",
					Description: `The list of access endpoint ip addresses of workflow.`,
				},
				resource.Attribute{
					Name:        "workflow_outbound_ip_addresses",
					Description: `The list of outgoing ip addresses of workflow. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Logic App Workflow.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_machine_learning_workspace",
			Category:         "Machine Learning",
			ShortDescription: `Gets information about an existing Machine Learning Workspace`,
			Description: `

Use this data source to access information about an existing Machine Learning Workspace.

`,
			Keywords: []string{
				"machine",
				"learning",
				"workspace",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Machine Learning Workspace exists.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The name of the Resource Group where the Machine Learning Workspace exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Machine Learning Workspace.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location where the Machine Learning Workspace exists.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the Machine Learning Workspace. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Machine Learning Workspace.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Machine Learning Workspace.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location where the Machine Learning Workspace exists.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the Machine Learning Workspace. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Machine Learning Workspace.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_maintenance_configuration",
			Category:         "Maintenance",
			ShortDescription: `Get information about an existing Maintenance Configuration.`,
			Description: `

Use this data source to access information about an existing Maintenance Configuration.

`,
			Keywords: []string{
				"maintenance",
				"configuration",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name of the Maintenance Configuration.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `Specifies the name of the Resource Group where this Maintenance Configuration exists. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the resource exists.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `The scope of the Maintenance Configuration.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Maintenance Configuration.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the resource exists.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `The scope of the Maintenance Configuration.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Maintenance Configuration.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_managed_application_definition",
			Category:         "Managed Applications",
			ShortDescription: `Gets information about an existing Managed Application Definition`,
			Description: `

Uses this data source to access information about an existing Managed Application Definition.

`,
			Keywords: []string{
				"managed",
				"applications",
				"application",
				"definition",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the Managed Application Definition.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) Specifies the name of the Resource Group where this Managed Application Definition exists.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the resource exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Managed Application Definition. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Managed Application Definition.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Managed Application Definition. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Managed Application Definition.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_managed_disk",
			Category:         "Compute",
			ShortDescription: `Get information about an existing Managed Disk.`,
			Description: `

Use this data source to access information about an existing Managed Disk.

`,
			Keywords: []string{
				"compute",
				"managed",
				"disk",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name of the Managed Disk.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `Specifies the name of the Resource Group where this Managed Disk exists. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "disk_encryption_set_id",
					Description: `The ID of the Disk Encryption Set used to encrypt this Managed Disk.`,
				},
				resource.Attribute{
					Name:        "disk_iops_read_write",
					Description: `The number of IOPS allowed for this disk, where one operation can transfer between 4k and 256k bytes.`,
				},
				resource.Attribute{
					Name:        "disk_mbps_read_write",
					Description: `The bandwidth allowed for this disk.`,
				},
				resource.Attribute{
					Name:        "disk_size_gb",
					Description: `The size of the Managed Disk in gigabytes.`,
				},
				resource.Attribute{
					Name:        "image_reference_id",
					Description: `The ID of the source image used for creating this Managed Disk.`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `The operating system used for this Managed Disk.`,
				},
				resource.Attribute{
					Name:        "storage_account_type",
					Description: `The storage account type for the Managed Disk.`,
				},
				resource.Attribute{
					Name:        "source_uri",
					Description: `The Source URI for this Managed Disk.`,
				},
				resource.Attribute{
					Name:        "source_resource_id",
					Description: `The ID of an existing Managed Disk which this Disk was created from.`,
				},
				resource.Attribute{
					Name:        "storage_account_id",
					Description: `The ID of the Storage Account where the ` + "`" + `source_uri` + "`" + ` is located.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of Availability Zones where the Managed Disk exists. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Managed Disk.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "disk_encryption_set_id",
					Description: `The ID of the Disk Encryption Set used to encrypt this Managed Disk.`,
				},
				resource.Attribute{
					Name:        "disk_iops_read_write",
					Description: `The number of IOPS allowed for this disk, where one operation can transfer between 4k and 256k bytes.`,
				},
				resource.Attribute{
					Name:        "disk_mbps_read_write",
					Description: `The bandwidth allowed for this disk.`,
				},
				resource.Attribute{
					Name:        "disk_size_gb",
					Description: `The size of the Managed Disk in gigabytes.`,
				},
				resource.Attribute{
					Name:        "image_reference_id",
					Description: `The ID of the source image used for creating this Managed Disk.`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `The operating system used for this Managed Disk.`,
				},
				resource.Attribute{
					Name:        "storage_account_type",
					Description: `The storage account type for the Managed Disk.`,
				},
				resource.Attribute{
					Name:        "source_uri",
					Description: `The Source URI for this Managed Disk.`,
				},
				resource.Attribute{
					Name:        "source_resource_id",
					Description: `The ID of an existing Managed Disk which this Disk was created from.`,
				},
				resource.Attribute{
					Name:        "storage_account_id",
					Description: `The ID of the Storage Account where the ` + "`" + `source_uri` + "`" + ` is located.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of Availability Zones where the Managed Disk exists. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Managed Disk.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_management_group",
			Category:         "Management",
			ShortDescription: `Gets information about an existing Management Group.`,
			Description: `

Use this data source to access information about an existing Management Group.

`,
			Keywords: []string{
				"management",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name or UUID of this Management Group.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `Specifies the name or UUID of this Management Group. ~>`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Specifies the display name of this Management Group. ~>`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Management Group.`,
				},
				resource.Attribute{
					Name:        "parent_management_group_id",
					Description: `The ID of any Parent Management Group.`,
				},
				resource.Attribute{
					Name:        "subscription_ids",
					Description: `A list of Subscription IDs which are assigned to the Management Group. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Management Group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Management Group.`,
				},
				resource.Attribute{
					Name:        "parent_management_group_id",
					Description: `The ID of any Parent Management Group.`,
				},
				resource.Attribute{
					Name:        "subscription_ids",
					Description: `A list of Subscription IDs which are assigned to the Management Group. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Management Group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_maps_account",
			Category:         "Maps",
			ShortDescription: `Gets information about an existing Azure Maps Account.`,
			Description: `

Use this data source to access information about an existing Azure Maps Account.

`,
			Keywords: []string{
				"maps",
				"account",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name of the Maps Account.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `Specifies the name of the Resource Group in which the Maps Account is located. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Maps Account.`,
				},
				resource.Attribute{
					Name:        "sku_name",
					Description: `The sku of the Azure Maps Account.`,
				},
				resource.Attribute{
					Name:        "primary_access_key",
					Description: `The primary key used to authenticate and authorize access to the Maps REST APIs.`,
				},
				resource.Attribute{
					Name:        "secondary_access_key",
					Description: `The primary key used to authenticate and authorize access to the Maps REST APIs. The second key is given to provide seamless key regeneration.`,
				},
				resource.Attribute{
					Name:        "x_ms_client_id",
					Description: `A unique identifier for the Maps Account. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Maps Account.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Maps Account.`,
				},
				resource.Attribute{
					Name:        "sku_name",
					Description: `The sku of the Azure Maps Account.`,
				},
				resource.Attribute{
					Name:        "primary_access_key",
					Description: `The primary key used to authenticate and authorize access to the Maps REST APIs.`,
				},
				resource.Attribute{
					Name:        "secondary_access_key",
					Description: `The primary key used to authenticate and authorize access to the Maps REST APIs. The second key is given to provide seamless key regeneration.`,
				},
				resource.Attribute{
					Name:        "x_ms_client_id",
					Description: `A unique identifier for the Maps Account. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Maps Account.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_mariadb_server",
			Category:         "Database",
			ShortDescription: `Gets information about an existing MariaDB Server.`,
			Description: `

Use this data source to access information about an existing MariaDB Server.

`,
			Icon: "Databases/10123-icon-service-Azure-Database-MariaDB-Server.svg",
			Keywords: []string{
				"database",
				"mariadb",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the MariaDB Server to retrieve information about.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of the resource group where the MariaDB Server exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the MariaDB Server.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `The FQDN of the MariaDB Server.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the resource exists.`,
				},
				resource.Attribute{
					Name:        "sku_name",
					Description: `The SKU Name for this MariaDB Server.`,
				},
				resource.Attribute{
					Name:        "storage_profile",
					Description: `A ` + "`" + `storage_profile` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "administrator_login",
					Description: `The Administrator Login for the MariaDB Server.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version of MariaDB being used.`,
				},
				resource.Attribute{
					Name:        "ssl_enforcement",
					Description: `The SSL being enforced on connections.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource. --- A ` + "`" + `storage_profile` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "storage_mb",
					Description: `The max storage allowed for a server.`,
				},
				resource.Attribute{
					Name:        "backup_retention_days",
					Description: `Backup retention days for the server.`,
				},
				resource.Attribute{
					Name:        "geo_redundant_backup",
					Description: `Whether Geo-redundant is enabled or not for server backup.`,
				},
				resource.Attribute{
					Name:        "auto_grow",
					Description: `Whether autogrow is enabled or disabled for the storage. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the MariaDB Server.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the MariaDB Server.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `The FQDN of the MariaDB Server.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the resource exists.`,
				},
				resource.Attribute{
					Name:        "sku_name",
					Description: `The SKU Name for this MariaDB Server.`,
				},
				resource.Attribute{
					Name:        "storage_profile",
					Description: `A ` + "`" + `storage_profile` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "administrator_login",
					Description: `The Administrator Login for the MariaDB Server.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version of MariaDB being used.`,
				},
				resource.Attribute{
					Name:        "ssl_enforcement",
					Description: `The SSL being enforced on connections.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource. --- A ` + "`" + `storage_profile` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "storage_mb",
					Description: `The max storage allowed for a server.`,
				},
				resource.Attribute{
					Name:        "backup_retention_days",
					Description: `Backup retention days for the server.`,
				},
				resource.Attribute{
					Name:        "geo_redundant_backup",
					Description: `Whether Geo-redundant is enabled or not for server backup.`,
				},
				resource.Attribute{
					Name:        "auto_grow",
					Description: `Whether autogrow is enabled or disabled for the storage. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the MariaDB Server.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_monitor_action_group",
			Category:         "Monitor",
			ShortDescription: `Get information about the specified Action Group.`,
			Description: `

Use this data source to access the properties of an Action Group.

`,
			Keywords: []string{
				"monitor",
				"action",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name of the Action Group.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `Specifies the name of the resource group the Action Group is located in. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Action Group.`,
				},
				resource.Attribute{
					Name:        "short_name",
					Description: `The short name of the action group.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Whether this action group is enabled.`,
				},
				resource.Attribute{
					Name:        "arm_role_receiver",
					Description: `One or more ` + "`" + `arm_role_receiver` + "`" + ` blocks as defined below.`,
				},
				resource.Attribute{
					Name:        "automation_runbook_receiver",
					Description: `One or more ` + "`" + `automation_runbook_receiver` + "`" + ` blocks as defined below.`,
				},
				resource.Attribute{
					Name:        "azure_app_push_receiver",
					Description: `One or more ` + "`" + `azure_app_push_receiver` + "`" + ` blocks as defined below.`,
				},
				resource.Attribute{
					Name:        "azure_function_receiver",
					Description: `One or more ` + "`" + `azure_function_receiver` + "`" + ` blocks as defined below.`,
				},
				resource.Attribute{
					Name:        "email_receiver",
					Description: `One or more ` + "`" + `email_receiver` + "`" + ` blocks as defined below.`,
				},
				resource.Attribute{
					Name:        "itsm_receiver",
					Description: `One or more ` + "`" + `itsm_receiver` + "`" + ` blocks as defined below.`,
				},
				resource.Attribute{
					Name:        "logic_app_receiver",
					Description: `One or more ` + "`" + `logic_app_receiver` + "`" + ` blocks as defined below.`,
				},
				resource.Attribute{
					Name:        "sms_receiver",
					Description: `One or more ` + "`" + `sms_receiver` + "`" + ` blocks as defined below.`,
				},
				resource.Attribute{
					Name:        "webhook_receiver",
					Description: `One or more ` + "`" + `webhook_receiver` + "`" + ` blocks as defined below.`,
				},
				resource.Attribute{
					Name:        "voice_receiver",
					Description: `One or more ` + "`" + `voice_receiver` + "`" + ` blocks as defined below. --- ` + "`" + `arm_role_receiver` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the ARM role receiver.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `The arm role id.`,
				},
				resource.Attribute{
					Name:        "use_common_alert_schema",
					Description: `Indicates whether to use common alert schema. --- ` + "`" + `automation_runbook_receiver` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the automation runbook receiver.`,
				},
				resource.Attribute{
					Name:        "automation_account_id",
					Description: `The automation account ID which holds this runbook and authenticates to Azure resources.`,
				},
				resource.Attribute{
					Name:        "runbook_name",
					Description: `The name for this runbook.`,
				},
				resource.Attribute{
					Name:        "webhook_resource_id",
					Description: `The resource id for webhook linked to this runbook.`,
				},
				resource.Attribute{
					Name:        "is_global_runbook",
					Description: `Indicates whether this instance is global runbook.`,
				},
				resource.Attribute{
					Name:        "service_uri",
					Description: `The URI where webhooks should be sent.`,
				},
				resource.Attribute{
					Name:        "use_common_alert_schema",
					Description: `Indicates whether to use common alert schema. --- ` + "`" + `azure_app_push_receiver` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Azure app push receiver.`,
				},
				resource.Attribute{
					Name:        "email_address",
					Description: `The email address of the user signed into the mobile app who will receive push notifications from this receiver. --- ` + "`" + `azure_function_receiver` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Azure Function receiver.`,
				},
				resource.Attribute{
					Name:        "function_app_resource_id",
					Description: `The Azure resource ID of the function app.`,
				},
				resource.Attribute{
					Name:        "function_name",
					Description: `The function name in the function app.`,
				},
				resource.Attribute{
					Name:        "http_trigger_url",
					Description: `The http trigger url where http request sent to.`,
				},
				resource.Attribute{
					Name:        "use_common_alert_schema",
					Description: `Indicates whether to use common alert schema. --- ` + "`" + `email_receiver` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the email receiver.`,
				},
				resource.Attribute{
					Name:        "email_address",
					Description: `The email address of this receiver.`,
				},
				resource.Attribute{
					Name:        "use_common_alert_schema",
					Description: `Indicates whether to use common alert schema. --- ` + "`" + `itsm_receiver` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the ITSM receiver.`,
				},
				resource.Attribute{
					Name:        "workspace_id",
					Description: `The Azure Log Analytics workspace ID where this connection is defined.`,
				},
				resource.Attribute{
					Name:        "connection_id",
					Description: `The unique connection identifier of the ITSM connection.`,
				},
				resource.Attribute{
					Name:        "ticket_configuration",
					Description: `A JSON blob for the configurations of the ITSM action. CreateMultipleWorkItems option will be part of this blob as well.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region of the workspace. --- ` + "`" + `logic_app_receiver` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the logic app receiver.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `The Azure resource ID of the logic app.`,
				},
				resource.Attribute{
					Name:        "callback_url",
					Description: `The callback url where http request sent to.`,
				},
				resource.Attribute{
					Name:        "use_common_alert_schema",
					Description: `Indicates whether to use common alert schema. --- ` + "`" + `sms_receiver` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the SMS receiver.`,
				},
				resource.Attribute{
					Name:        "country_code",
					Description: `The country code of the SMS receiver.`,
				},
				resource.Attribute{
					Name:        "phone_number",
					Description: `The phone number of the SMS receiver. --- ` + "`" + `voice_receiver` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the voice receiver.`,
				},
				resource.Attribute{
					Name:        "country_code",
					Description: `The country code of the voice receiver.`,
				},
				resource.Attribute{
					Name:        "phone_number",
					Description: `The phone number of the voice receiver. --- ` + "`" + `webhook_receiver` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the webhook receiver.`,
				},
				resource.Attribute{
					Name:        "service_uri",
					Description: `The URI where webhooks should be sent.`,
				},
				resource.Attribute{
					Name:        "use_common_alert_schema",
					Description: `Indicates whether to use common alert schema. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Action Group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Action Group.`,
				},
				resource.Attribute{
					Name:        "short_name",
					Description: `The short name of the action group.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Whether this action group is enabled.`,
				},
				resource.Attribute{
					Name:        "arm_role_receiver",
					Description: `One or more ` + "`" + `arm_role_receiver` + "`" + ` blocks as defined below.`,
				},
				resource.Attribute{
					Name:        "automation_runbook_receiver",
					Description: `One or more ` + "`" + `automation_runbook_receiver` + "`" + ` blocks as defined below.`,
				},
				resource.Attribute{
					Name:        "azure_app_push_receiver",
					Description: `One or more ` + "`" + `azure_app_push_receiver` + "`" + ` blocks as defined below.`,
				},
				resource.Attribute{
					Name:        "azure_function_receiver",
					Description: `One or more ` + "`" + `azure_function_receiver` + "`" + ` blocks as defined below.`,
				},
				resource.Attribute{
					Name:        "email_receiver",
					Description: `One or more ` + "`" + `email_receiver` + "`" + ` blocks as defined below.`,
				},
				resource.Attribute{
					Name:        "itsm_receiver",
					Description: `One or more ` + "`" + `itsm_receiver` + "`" + ` blocks as defined below.`,
				},
				resource.Attribute{
					Name:        "logic_app_receiver",
					Description: `One or more ` + "`" + `logic_app_receiver` + "`" + ` blocks as defined below.`,
				},
				resource.Attribute{
					Name:        "sms_receiver",
					Description: `One or more ` + "`" + `sms_receiver` + "`" + ` blocks as defined below.`,
				},
				resource.Attribute{
					Name:        "webhook_receiver",
					Description: `One or more ` + "`" + `webhook_receiver` + "`" + ` blocks as defined below.`,
				},
				resource.Attribute{
					Name:        "voice_receiver",
					Description: `One or more ` + "`" + `voice_receiver` + "`" + ` blocks as defined below. --- ` + "`" + `arm_role_receiver` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the ARM role receiver.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `The arm role id.`,
				},
				resource.Attribute{
					Name:        "use_common_alert_schema",
					Description: `Indicates whether to use common alert schema. --- ` + "`" + `automation_runbook_receiver` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the automation runbook receiver.`,
				},
				resource.Attribute{
					Name:        "automation_account_id",
					Description: `The automation account ID which holds this runbook and authenticates to Azure resources.`,
				},
				resource.Attribute{
					Name:        "runbook_name",
					Description: `The name for this runbook.`,
				},
				resource.Attribute{
					Name:        "webhook_resource_id",
					Description: `The resource id for webhook linked to this runbook.`,
				},
				resource.Attribute{
					Name:        "is_global_runbook",
					Description: `Indicates whether this instance is global runbook.`,
				},
				resource.Attribute{
					Name:        "service_uri",
					Description: `The URI where webhooks should be sent.`,
				},
				resource.Attribute{
					Name:        "use_common_alert_schema",
					Description: `Indicates whether to use common alert schema. --- ` + "`" + `azure_app_push_receiver` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Azure app push receiver.`,
				},
				resource.Attribute{
					Name:        "email_address",
					Description: `The email address of the user signed into the mobile app who will receive push notifications from this receiver. --- ` + "`" + `azure_function_receiver` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Azure Function receiver.`,
				},
				resource.Attribute{
					Name:        "function_app_resource_id",
					Description: `The Azure resource ID of the function app.`,
				},
				resource.Attribute{
					Name:        "function_name",
					Description: `The function name in the function app.`,
				},
				resource.Attribute{
					Name:        "http_trigger_url",
					Description: `The http trigger url where http request sent to.`,
				},
				resource.Attribute{
					Name:        "use_common_alert_schema",
					Description: `Indicates whether to use common alert schema. --- ` + "`" + `email_receiver` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the email receiver.`,
				},
				resource.Attribute{
					Name:        "email_address",
					Description: `The email address of this receiver.`,
				},
				resource.Attribute{
					Name:        "use_common_alert_schema",
					Description: `Indicates whether to use common alert schema. --- ` + "`" + `itsm_receiver` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the ITSM receiver.`,
				},
				resource.Attribute{
					Name:        "workspace_id",
					Description: `The Azure Log Analytics workspace ID where this connection is defined.`,
				},
				resource.Attribute{
					Name:        "connection_id",
					Description: `The unique connection identifier of the ITSM connection.`,
				},
				resource.Attribute{
					Name:        "ticket_configuration",
					Description: `A JSON blob for the configurations of the ITSM action. CreateMultipleWorkItems option will be part of this blob as well.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region of the workspace. --- ` + "`" + `logic_app_receiver` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the logic app receiver.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `The Azure resource ID of the logic app.`,
				},
				resource.Attribute{
					Name:        "callback_url",
					Description: `The callback url where http request sent to.`,
				},
				resource.Attribute{
					Name:        "use_common_alert_schema",
					Description: `Indicates whether to use common alert schema. --- ` + "`" + `sms_receiver` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the SMS receiver.`,
				},
				resource.Attribute{
					Name:        "country_code",
					Description: `The country code of the SMS receiver.`,
				},
				resource.Attribute{
					Name:        "phone_number",
					Description: `The phone number of the SMS receiver. --- ` + "`" + `voice_receiver` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the voice receiver.`,
				},
				resource.Attribute{
					Name:        "country_code",
					Description: `The country code of the voice receiver.`,
				},
				resource.Attribute{
					Name:        "phone_number",
					Description: `The phone number of the voice receiver. --- ` + "`" + `webhook_receiver` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the webhook receiver.`,
				},
				resource.Attribute{
					Name:        "service_uri",
					Description: `The URI where webhooks should be sent.`,
				},
				resource.Attribute{
					Name:        "use_common_alert_schema",
					Description: `Indicates whether to use common alert schema. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Action Group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_monitor_diagnostic_categories",
			Category:         "Monitor",
			ShortDescription: `Gets information about an the Monitor Diagnostics Categories supported by an existing Resource.`,
			Description: `

Use this data source to access information about the Monitor Diagnostics Categories supported by an existing Resource.

`,
			Keywords: []string{
				"monitor",
				"diagnostic",
				"categories",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_id",
					Description: `The ID of an existing Resource which Monitor Diagnostics Categories should be retrieved for. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Resource.`,
				},
				resource.Attribute{
					Name:        "logs",
					Description: `A list of the Log Categories supported for this Resource.`,
				},
				resource.Attribute{
					Name:        "metrics",
					Description: `A list of the Metric Categories supported for this Resource. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Monitor Diagnostics Categories.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Resource.`,
				},
				resource.Attribute{
					Name:        "logs",
					Description: `A list of the Log Categories supported for this Resource.`,
				},
				resource.Attribute{
					Name:        "metrics",
					Description: `A list of the Metric Categories supported for this Resource. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Monitor Diagnostics Categories.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_monitor_log_profile",
			Category:         "Monitor",
			ShortDescription: `Get information about the specified Log Profile.`,
			Description: `

Use this data source to access the properties of a Log Profile.

`,
			Keywords: []string{
				"monitor",
				"log",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the Name of the Log Profile. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Log Profile.`,
				},
				resource.Attribute{
					Name:        "storage_account_id",
					Description: `The resource id of the storage account in which the Activity Log is stored.`,
				},
				resource.Attribute{
					Name:        "servicebus_rule_id",
					Description: `The service bus (or event hub) rule ID of the service bus (or event hub) namespace in which the Activity Log is streamed to.`,
				},
				resource.Attribute{
					Name:        "locations",
					Description: `List of regions for which Activity Log events are stored or streamed.`,
				},
				resource.Attribute{
					Name:        "categories",
					Description: `List of categories of the logs.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `A boolean value indicating whether the retention policy is enabled.`,
				},
				resource.Attribute{
					Name:        "days",
					Description: `The number of days for the retention policy. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Log Profile.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Log Profile.`,
				},
				resource.Attribute{
					Name:        "storage_account_id",
					Description: `The resource id of the storage account in which the Activity Log is stored.`,
				},
				resource.Attribute{
					Name:        "servicebus_rule_id",
					Description: `The service bus (or event hub) rule ID of the service bus (or event hub) namespace in which the Activity Log is streamed to.`,
				},
				resource.Attribute{
					Name:        "locations",
					Description: `List of regions for which Activity Log events are stored or streamed.`,
				},
				resource.Attribute{
					Name:        "categories",
					Description: `List of categories of the logs.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `A boolean value indicating whether the retention policy is enabled.`,
				},
				resource.Attribute{
					Name:        "days",
					Description: `The number of days for the retention policy. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Log Profile.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_monitor_scheduled_query_rules_alert",
			Category:         "Monitor",
			ShortDescription: `Get information about the specified AlertingAction Scheduled Query Rules resource.`,
			Description: `

Use this data source to access the properties of an AlertingAction scheduled query rule.

`,
			Keywords: []string{
				"monitor",
				"scheduled",
				"query",
				"rules",
				"alert",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the scheduled query rule.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) Specifies the name of the resource group where the scheduled query rule is located. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the scheduled query rule.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `An ` + "`" + `action` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "authorized_resource_ids",
					Description: `The list of Resource IDs referred into query.`,
				},
				resource.Attribute{
					Name:        "data_source_id",
					Description: `The resource URI over which log search query is to be run.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the scheduled query rule.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Whether this scheduled query rule is enabled.`,
				},
				resource.Attribute{
					Name:        "frequency",
					Description: `Frequency at which rule condition should be evaluated.`,
				},
				resource.Attribute{
					Name:        "query",
					Description: `Log search query.`,
				},
				resource.Attribute{
					Name:        "time_window",
					Description: `Time window for which data needs to be fetched for query.`,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `Severity of the alert.`,
				},
				resource.Attribute{
					Name:        "throttling",
					Description: `Time for which alerts should be throttled or suppressed.`,
				},
				resource.Attribute{
					Name:        "trigger",
					Description: `A ` + "`" + `trigger` + "`" + ` block as defined below. ---`,
				},
				resource.Attribute{
					Name:        "action_group",
					Description: `List of action group reference resource IDs.`,
				},
				resource.Attribute{
					Name:        "custom_webhook_payload",
					Description: `Custom payload to be sent for all webhook URI in Azure action group.`,
				},
				resource.Attribute{
					Name:        "email_subject",
					Description: `Custom subject override for all email IDs in Azure action group. --- ` + "`" + `metricTrigger` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "metricColumn",
					Description: `Evaluation of metric on a particular column.`,
				},
				resource.Attribute{
					Name:        "metricTriggerType",
					Description: `The metric trigger type.`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `Evaluation operation for rule.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `The threshold of the metric trigger. --- ` + "`" + `trigger` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "metricTrigger",
					Description: `A ` + "`" + `metricTrigger` + "`" + ` block as defined above.`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `Evaluation operation for rule.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `Result or count threshold based on which rule should be triggered. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the App Service Environment.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the scheduled query rule.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `An ` + "`" + `action` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "authorized_resource_ids",
					Description: `The list of Resource IDs referred into query.`,
				},
				resource.Attribute{
					Name:        "data_source_id",
					Description: `The resource URI over which log search query is to be run.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the scheduled query rule.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Whether this scheduled query rule is enabled.`,
				},
				resource.Attribute{
					Name:        "frequency",
					Description: `Frequency at which rule condition should be evaluated.`,
				},
				resource.Attribute{
					Name:        "query",
					Description: `Log search query.`,
				},
				resource.Attribute{
					Name:        "time_window",
					Description: `Time window for which data needs to be fetched for query.`,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `Severity of the alert.`,
				},
				resource.Attribute{
					Name:        "throttling",
					Description: `Time for which alerts should be throttled or suppressed.`,
				},
				resource.Attribute{
					Name:        "trigger",
					Description: `A ` + "`" + `trigger` + "`" + ` block as defined below. ---`,
				},
				resource.Attribute{
					Name:        "action_group",
					Description: `List of action group reference resource IDs.`,
				},
				resource.Attribute{
					Name:        "custom_webhook_payload",
					Description: `Custom payload to be sent for all webhook URI in Azure action group.`,
				},
				resource.Attribute{
					Name:        "email_subject",
					Description: `Custom subject override for all email IDs in Azure action group. --- ` + "`" + `metricTrigger` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "metricColumn",
					Description: `Evaluation of metric on a particular column.`,
				},
				resource.Attribute{
					Name:        "metricTriggerType",
					Description: `The metric trigger type.`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `Evaluation operation for rule.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `The threshold of the metric trigger. --- ` + "`" + `trigger` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "metricTrigger",
					Description: `A ` + "`" + `metricTrigger` + "`" + ` block as defined above.`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `Evaluation operation for rule.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `Result or count threshold based on which rule should be triggered. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the App Service Environment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_monitor_scheduled_query_rules_log",
			Category:         "Monitor",
			ShortDescription: `Get information about the specified LogToMetricAction Scheduled Query Rules resource.`,
			Description: `

Use this data source to access the properties of a LogToMetricAction scheduled query rule.

`,
			Keywords: []string{
				"monitor",
				"scheduled",
				"query",
				"rules",
				"log",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the scheduled query rule.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) Specifies the name of the resource group where the scheduled query rule is located. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the scheduled query rule.`,
				},
				resource.Attribute{
					Name:        "criteria",
					Description: `A ` + "`" + `criteria` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "data_source_id",
					Description: `The resource URI over which log search query is to be run.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the scheduled query rule.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Whether this scheduled query rule is enabled. --- ` + "`" + `criteria` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "dimension",
					Description: `A ` + "`" + `dimension` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "metric_name",
					Description: `Name of the metric. --- ` + "`" + `dimension` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the dimension.`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `Operator for dimension values.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `List of dimension values. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the App Service Environment.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the scheduled query rule.`,
				},
				resource.Attribute{
					Name:        "criteria",
					Description: `A ` + "`" + `criteria` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "data_source_id",
					Description: `The resource URI over which log search query is to be run.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the scheduled query rule.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Whether this scheduled query rule is enabled. --- ` + "`" + `criteria` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "dimension",
					Description: `A ` + "`" + `dimension` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "metric_name",
					Description: `Name of the metric. --- ` + "`" + `dimension` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the dimension.`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `Operator for dimension values.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `List of dimension values. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the App Service Environment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_mssql_database",
			Category:         "Database",
			ShortDescription: `Gets information about an existing SQL database.`,
			Description: `

Use this data source to access information about an existing SQL database.

`,
			Keywords: []string{
				"database",
				"mssql",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Ms SQL Database.`,
				},
				resource.Attribute{
					Name:        "server_id",
					Description: `The id of the Ms SQL Server on which to create the database. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "collation",
					Description: `The collation of the database.`,
				},
				resource.Attribute{
					Name:        "elastic_pool_id",
					Description: `The id of the elastic pool containing this database.`,
				},
				resource.Attribute{
					Name:        "license_type",
					Description: `The license type to apply for this database.`,
				},
				resource.Attribute{
					Name:        "max_size_gb",
					Description: `The max size of the database in gigabytes.`,
				},
				resource.Attribute{
					Name:        "read_replica_count",
					Description: `The number of readonly secondary replicas associated with the database to which readonly application intent connections may be routed.`,
				},
				resource.Attribute{
					Name:        "read_scale",
					Description: `If enabled, connections that have application intent set to readonly in their connection string may be routed to a readonly secondary replica.`,
				},
				resource.Attribute{
					Name:        "sku_name",
					Description: `The name of the sku of the database.`,
				},
				resource.Attribute{
					Name:        "storage_account_type",
					Description: `The storage account type used to store backups for this database.`,
				},
				resource.Attribute{
					Name:        "zone_redundant",
					Description: `Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the resource. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the SQL database.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "collation",
					Description: `The collation of the database.`,
				},
				resource.Attribute{
					Name:        "elastic_pool_id",
					Description: `The id of the elastic pool containing this database.`,
				},
				resource.Attribute{
					Name:        "license_type",
					Description: `The license type to apply for this database.`,
				},
				resource.Attribute{
					Name:        "max_size_gb",
					Description: `The max size of the database in gigabytes.`,
				},
				resource.Attribute{
					Name:        "read_replica_count",
					Description: `The number of readonly secondary replicas associated with the database to which readonly application intent connections may be routed.`,
				},
				resource.Attribute{
					Name:        "read_scale",
					Description: `If enabled, connections that have application intent set to readonly in their connection string may be routed to a readonly secondary replica.`,
				},
				resource.Attribute{
					Name:        "sku_name",
					Description: `The name of the sku of the database.`,
				},
				resource.Attribute{
					Name:        "storage_account_type",
					Description: `The storage account type used to store backups for this database.`,
				},
				resource.Attribute{
					Name:        "zone_redundant",
					Description: `Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the resource. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the SQL database.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_mssql_elasticpool",
			Category:         "Database",
			ShortDescription: `Gets information about an existing SQL elastic pool.`,
			Description: `

Use this data source to access information about an existing SQL elastic pool.

`,
			Icon: "Databases/10134-icon-service-SQL-Elastic-Pools.svg",
			Keywords: []string{
				"database",
				"mssql",
				"elasticpool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the elastic pool.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of the resource group which contains the elastic pool.`,
				},
				resource.Attribute{
					Name:        "server_name",
					Description: `The name of the SQL Server which contains the elastic pool. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "license_type",
					Description: `The license type to apply for this database.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Specifies the supported Azure location where the resource exists.`,
				},
				resource.Attribute{
					Name:        "max_size_gb",
					Description: `The max data size of the elastic pool in gigabytes.`,
				},
				resource.Attribute{
					Name:        "max_size_bytes",
					Description: `The max data size of the elastic pool in bytes.`,
				},
				resource.Attribute{
					Name:        "per_db_min_capacity",
					Description: `The minimum capacity all databases are guaranteed.`,
				},
				resource.Attribute{
					Name:        "per_db_max_capacity",
					Description: `The maximum capacity any one database can consume.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "zone_redundant",
					Description: `Whether or not this elastic pool is zone redundant. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the SQL elastic pool.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "license_type",
					Description: `The license type to apply for this database.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Specifies the supported Azure location where the resource exists.`,
				},
				resource.Attribute{
					Name:        "max_size_gb",
					Description: `The max data size of the elastic pool in gigabytes.`,
				},
				resource.Attribute{
					Name:        "max_size_bytes",
					Description: `The max data size of the elastic pool in bytes.`,
				},
				resource.Attribute{
					Name:        "per_db_min_capacity",
					Description: `The minimum capacity all databases are guaranteed.`,
				},
				resource.Attribute{
					Name:        "per_db_max_capacity",
					Description: `The maximum capacity any one database can consume.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "zone_redundant",
					Description: `Whether or not this elastic pool is zone redundant. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the SQL elastic pool.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_mssql_server",
			Category:         "Database",
			ShortDescription: `Gets information about an existing Microsoft SQL Server.`,
			Description: `

Use this data source to access information about an existing Microsoft SQL Server.

`,
			Keywords: []string{
				"database",
				"mssql",
				"server",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Microsoft SQL Server.`,
				},
				resource.Attribute{
					Name:        "administrator_login",
					Description: `The server's administrator login name.`,
				},
				resource.Attribute{
					Name:        "fully_qualified_domain_name",
					Description: `The fully qualified domain name of the Azure SQL Server.`,
				},
				resource.Attribute{
					Name:        "identity",
					Description: `A ` + "`" + `identity` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure Region where the Microsoft SQL Server exists.`,
				},
				resource.Attribute{
					Name:        "restorable_dropped_database_ids",
					Description: `A list of dropped restorable database IDs on the server.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to this Microsoft SQL Server.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `This servers MS SQL version. --- A ` + "`" + `identity` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "principal_id",
					Description: `The Principal ID for the Service Principal associated with the Identity of this SQL Server.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `The Tenant ID for the Service Principal associated with the Identity of this SQL Server.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The identity type of the Microsoft SQL Server. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Microsoft SQL Server.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_mysql_server",
			Category:         "Database",
			ShortDescription: `Gets information about an existing MySQL Server.`,
			Description:      ``,
			Icon:             "Databases/10122-icon-service-Azure-Database-MySQL-Server.svg",
			Keywords: []string{
				"database",
				"mysql",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the MySQL Server.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The name of the resource group for the MySQL Server. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the MySQL Server.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `The FQDN of the MySQL Server.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the resource exists.`,
				},
				resource.Attribute{
					Name:        "sku_name",
					Description: `The SKU Name for this MySQL Server.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version of this MySQL Server.`,
				},
				resource.Attribute{
					Name:        "administrator_login",
					Description: `The Administrator Login for the MySQL Server.`,
				},
				resource.Attribute{
					Name:        "auto_grow_enabled",
					Description: `The auto grow setting for this MySQL Server.`,
				},
				resource.Attribute{
					Name:        "backup_retention_days",
					Description: `The backup retention days for this MySQL server.`,
				},
				resource.Attribute{
					Name:        "geo_redundant_backup_enabled",
					Description: `The geo redundant backup setting for this MySQL Server.`,
				},
				resource.Attribute{
					Name:        "identity",
					Description: `An ` + "`" + `identity` + "`" + ` block for this MySQL server as defined below.`,
				},
				resource.Attribute{
					Name:        "infrastructure_encryption_enabled",
					Description: `Whether or not infrastructure is encrypted for this MySQL Server.`,
				},
				resource.Attribute{
					Name:        "public_network_access_enabled",
					Description: `Whether or not public network access is allowed for this MySQL Server.`,
				},
				resource.Attribute{
					Name:        "ssl_enforcement_enabled",
					Description: `Specifies if SSL should be enforced on connections for this MySQL Server.`,
				},
				resource.Attribute{
					Name:        "ssl_minimal_tls_version_enforced",
					Description: `The minimum TLS version to support for this MySQL Server.`,
				},
				resource.Attribute{
					Name:        "storage_mb",
					Description: `Max storage allowed for this MySQL Server.`,
				},
				resource.Attribute{
					Name:        "threat_detection_policy",
					Description: `Threat detection policy configuration, known in the API as Server Security Alerts Policy. The ` + "`" + `threat_detection_policy` + "`" + ` block exports fields documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the resource. --- An ` + "`" + `identity` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The Type of Identity used for this MySQL Server.`,
				},
				resource.Attribute{
					Name:        "principal_id",
					Description: `The Client ID of the Service Principal assigned to this MySQL Server.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `The ID of the Tenant the Service Principal is assigned in. --- A ` + "`" + `threat_detection_policy` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Is the policy enabled?`,
				},
				resource.Attribute{
					Name:        "disabled_alerts",
					Description: `Specifies a list of alerts which should be disabled. Possible values include ` + "`" + `Access_Anomaly` + "`" + `, ` + "`" + `Sql_Injection` + "`" + ` and ` + "`" + `Sql_Injection_Vulnerability` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "email_account_admins",
					Description: `Should the account administrators be emailed when this alert is triggered?`,
				},
				resource.Attribute{
					Name:        "email_addresses",
					Description: `A list of email addresses which alerts should be sent to.`,
				},
				resource.Attribute{
					Name:        "retention_days",
					Description: `Specifies the number of days to keep in the Threat Detection audit logs.`,
				},
				resource.Attribute{
					Name:        "storage_account_access_key",
					Description: `Specifies the identifier key of the Threat Detection audit storage account.`,
				},
				resource.Attribute{
					Name:        "storage_endpoint",
					Description: `Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the MySQL Server.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the MySQL Server.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `The FQDN of the MySQL Server.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the resource exists.`,
				},
				resource.Attribute{
					Name:        "sku_name",
					Description: `The SKU Name for this MySQL Server.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version of this MySQL Server.`,
				},
				resource.Attribute{
					Name:        "administrator_login",
					Description: `The Administrator Login for the MySQL Server.`,
				},
				resource.Attribute{
					Name:        "auto_grow_enabled",
					Description: `The auto grow setting for this MySQL Server.`,
				},
				resource.Attribute{
					Name:        "backup_retention_days",
					Description: `The backup retention days for this MySQL server.`,
				},
				resource.Attribute{
					Name:        "geo_redundant_backup_enabled",
					Description: `The geo redundant backup setting for this MySQL Server.`,
				},
				resource.Attribute{
					Name:        "identity",
					Description: `An ` + "`" + `identity` + "`" + ` block for this MySQL server as defined below.`,
				},
				resource.Attribute{
					Name:        "infrastructure_encryption_enabled",
					Description: `Whether or not infrastructure is encrypted for this MySQL Server.`,
				},
				resource.Attribute{
					Name:        "public_network_access_enabled",
					Description: `Whether or not public network access is allowed for this MySQL Server.`,
				},
				resource.Attribute{
					Name:        "ssl_enforcement_enabled",
					Description: `Specifies if SSL should be enforced on connections for this MySQL Server.`,
				},
				resource.Attribute{
					Name:        "ssl_minimal_tls_version_enforced",
					Description: `The minimum TLS version to support for this MySQL Server.`,
				},
				resource.Attribute{
					Name:        "storage_mb",
					Description: `Max storage allowed for this MySQL Server.`,
				},
				resource.Attribute{
					Name:        "threat_detection_policy",
					Description: `Threat detection policy configuration, known in the API as Server Security Alerts Policy. The ` + "`" + `threat_detection_policy` + "`" + ` block exports fields documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the resource. --- An ` + "`" + `identity` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The Type of Identity used for this MySQL Server.`,
				},
				resource.Attribute{
					Name:        "principal_id",
					Description: `The Client ID of the Service Principal assigned to this MySQL Server.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `The ID of the Tenant the Service Principal is assigned in. --- A ` + "`" + `threat_detection_policy` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Is the policy enabled?`,
				},
				resource.Attribute{
					Name:        "disabled_alerts",
					Description: `Specifies a list of alerts which should be disabled. Possible values include ` + "`" + `Access_Anomaly` + "`" + `, ` + "`" + `Sql_Injection` + "`" + ` and ` + "`" + `Sql_Injection_Vulnerability` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "email_account_admins",
					Description: `Should the account administrators be emailed when this alert is triggered?`,
				},
				resource.Attribute{
					Name:        "email_addresses",
					Description: `A list of email addresses which alerts should be sent to.`,
				},
				resource.Attribute{
					Name:        "retention_days",
					Description: `Specifies the number of days to keep in the Threat Detection audit logs.`,
				},
				resource.Attribute{
					Name:        "storage_account_access_key",
					Description: `Specifies the identifier key of the Threat Detection audit storage account.`,
				},
				resource.Attribute{
					Name:        "storage_endpoint",
					Description: `Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the MySQL Server.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_nat_gateway",
			Category:         "Network",
			ShortDescription: `Gets information about an existing NAT Gateway`,
			Description: `

Use this data source to access information about an existing NAT Gateway.

`,
			Keywords: []string{
				"network",
				"nat",
				"gateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the Name of the NAT Gateway.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `Specifies the name of the Resource Group where the NAT Gateway exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location where the NAT Gateway exists.`,
				},
				resource.Attribute{
					Name:        "idle_timeout_in_minutes",
					Description: `The idle timeout in minutes which is used for the NAT Gateway.`,
				},
				resource.Attribute{
					Name:        "public_ip_address_ids",
					Description: `A list of existing Public IP Address resource IDs which the NAT Gateway is using.`,
				},
				resource.Attribute{
					Name:        "public_ip_prefix_ids",
					Description: `A list of existing Public IP Prefix resource IDs which the NAT Gateway is using.`,
				},
				resource.Attribute{
					Name:        "resource_guid",
					Description: `The Resource GUID of the NAT Gateway.`,
				},
				resource.Attribute{
					Name:        "sku_name",
					Description: `The SKU used by the NAT Gateway.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of Availability Zones which the NAT Gateway exists in. ~>`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the NAT Gateway.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "location",
					Description: `The location where the NAT Gateway exists.`,
				},
				resource.Attribute{
					Name:        "idle_timeout_in_minutes",
					Description: `The idle timeout in minutes which is used for the NAT Gateway.`,
				},
				resource.Attribute{
					Name:        "public_ip_address_ids",
					Description: `A list of existing Public IP Address resource IDs which the NAT Gateway is using.`,
				},
				resource.Attribute{
					Name:        "public_ip_prefix_ids",
					Description: `A list of existing Public IP Prefix resource IDs which the NAT Gateway is using.`,
				},
				resource.Attribute{
					Name:        "resource_guid",
					Description: `The Resource GUID of the NAT Gateway.`,
				},
				resource.Attribute{
					Name:        "sku_name",
					Description: `The SKU used by the NAT Gateway.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of Availability Zones which the NAT Gateway exists in. ~>`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the NAT Gateway.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_netapp_account",
			Category:         "NetApp",
			ShortDescription: `Gets information about an existing NetApp Account`,
			Description: `

Uses this data source to access information about an existing NetApp Account.

`,
			Keywords: []string{
				"netapp",
				"account",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the NetApp Account.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The Name of the Resource Group where the NetApp Account exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure Region where the NetApp Account exists. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the NetApp Account.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "location",
					Description: `The Azure Region where the NetApp Account exists. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the NetApp Account.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_netapp_pool",
			Category:         "NetApp",
			ShortDescription: `Gets information about an existing NetApp Pool`,
			Description: `

Uses this data source to access information about an existing NetApp Pool.


`,
			Keywords: []string{
				"netapp",
				"pool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the NetApp Pool.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `The name of the NetApp account where the NetApp pool exists.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The Name of the Resource Group where the NetApp Pool exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure Region where the NetApp Pool exists.`,
				},
				resource.Attribute{
					Name:        "service_level",
					Description: `The service level of the file system.`,
				},
				resource.Attribute{
					Name:        "size_in_tb",
					Description: `Provisioned size of the pool in TB. --- ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the NetApp Pool.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "location",
					Description: `The Azure Region where the NetApp Pool exists.`,
				},
				resource.Attribute{
					Name:        "service_level",
					Description: `The service level of the file system.`,
				},
				resource.Attribute{
					Name:        "size_in_tb",
					Description: `Provisioned size of the pool in TB. --- ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the NetApp Pool.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_netapp_snapshot",
			Category:         "NetApp",
			ShortDescription: `Gets information about an existing NetApp Snapshot`,
			Description: `

Uses this data source to access information about an existing NetApp Snapshot.

`,
			Keywords: []string{
				"netapp",
				"snapshot",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the NetApp Snapshot.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `The name of the NetApp Account where the NetApp Pool exists.`,
				},
				resource.Attribute{
					Name:        "pool_name",
					Description: `The name of the NetApp Pool where the NetApp Volume exists.`,
				},
				resource.Attribute{
					Name:        "volume_name",
					Description: `The name of the NetApp Volume where the NetApp Snapshot exists.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The Name of the Resource Group where the NetApp Snapshot exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure Region where the NetApp Snapshot exists. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the NetApp Snapshot.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "location",
					Description: `The Azure Region where the NetApp Snapshot exists. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the NetApp Snapshot.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_netapp_volume",
			Category:         "NetApp",
			ShortDescription: `Gets information about an existing NetApp Volume`,
			Description: `

Uses this data source to access information about an existing NetApp Volume.

`,
			Keywords: []string{
				"netapp",
				"volume",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the NetApp Volume.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The Name of the Resource Group where the NetApp Volume exists.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `The name of the NetApp account where the NetApp pool exists.`,
				},
				resource.Attribute{
					Name:        "pool_name",
					Description: `The name of the NetApp pool where the NetApp volume exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure Region where the NetApp Volume exists.`,
				},
				resource.Attribute{
					Name:        "mount_ip_addresses",
					Description: `A list of IPv4 Addresses which should be used to mount the volume.`,
				},
				resource.Attribute{
					Name:        "volume_path",
					Description: `The unique file path of the volume.`,
				},
				resource.Attribute{
					Name:        "service_level",
					Description: `The service level of the file system.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of a Subnet in which the NetApp Volume resides.`,
				},
				resource.Attribute{
					Name:        "storage_quota_in_gb",
					Description: `The maximum Storage Quota in Gigabytes allowed for a file system. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the NetApp Volume.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "location",
					Description: `The Azure Region where the NetApp Volume exists.`,
				},
				resource.Attribute{
					Name:        "mount_ip_addresses",
					Description: `A list of IPv4 Addresses which should be used to mount the volume.`,
				},
				resource.Attribute{
					Name:        "volume_path",
					Description: `The unique file path of the volume.`,
				},
				resource.Attribute{
					Name:        "service_level",
					Description: `The service level of the file system.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of a Subnet in which the NetApp Volume resides.`,
				},
				resource.Attribute{
					Name:        "storage_quota_in_gb",
					Description: `The maximum Storage Quota in Gigabytes allowed for a file system. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the NetApp Volume.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_network_ddos_protection_plan",
			Category:         "Network",
			ShortDescription: `Use this data source to access information about an existing Azure Network DDoS Protection Plan.`,
			Description: `

Use this data source to access information about an existing Azure Network DDoS Protection Plan.

`,
			Keywords: []string{
				"network",
				"ddos",
				"protection",
				"plan",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Network DDoS Protection Plan.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of the resource group where the Network DDoS Protection Plan exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the DDoS Protection Plan`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Specifies the supported Azure location where the resource exists.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource.`,
				},
				resource.Attribute{
					Name:        "virtual_network_ids",
					Description: `A list of ID's of the Virtual Networks associated with this DDoS Protection Plan. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Protection Plan.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the DDoS Protection Plan`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Specifies the supported Azure location where the resource exists.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource.`,
				},
				resource.Attribute{
					Name:        "virtual_network_ids",
					Description: `A list of ID's of the Virtual Networks associated with this DDoS Protection Plan. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Protection Plan.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_network_interface",
			Category:         "Network",
			ShortDescription: `Gets information about an existing Network Interface.`,
			Description: `

Use this data source to access information about an existing Network Interface.

`,
			Icon: "Networking/10080-icon-service-Network-Interfaces.svg",
			Keywords: []string{
				"network",
				"interface",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name of the Network Interface.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `Specifies the name of the resource group the Network Interface is located in. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Network Interface.`,
				},
				resource.Attribute{
					Name:        "applied_dns_servers",
					Description: `List of DNS servers applied to the specified Network Interface.`,
				},
				resource.Attribute{
					Name:        "enable_accelerated_networking",
					Description: `Indicates if accelerated networking is set on the specified Network Interface.`,
				},
				resource.Attribute{
					Name:        "enable_ip_forwarding",
					Description: `Indicate if IP forwarding is set on the specified Network Interface.`,
				},
				resource.Attribute{
					Name:        "dns_servers",
					Description: `The list of DNS servers used by the specified Network Interface.`,
				},
				resource.Attribute{
					Name:        "internal_dns_name_label",
					Description: `The internal dns name label of the specified Network Interface.`,
				},
				resource.Attribute{
					Name:        "ip_configuration",
					Description: `One or more ` + "`" + `ip_configuration` + "`" + ` blocks as defined below.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location of the specified Network Interface.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `The MAC address used by the specified Network Interface.`,
				},
				resource.Attribute{
					Name:        "network_security_group_id",
					Description: `The ID of the network security group associated to the specified Network Interface.`,
				},
				resource.Attribute{
					Name:        "private_ip_address",
					Description: `The primary private ip address associated to the specified Network Interface.`,
				},
				resource.Attribute{
					Name:        "private_ip_addresses",
					Description: `The list of private ip addresses associates to the specified Network Interface.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `List the tags associated to the specified Network Interface.`,
				},
				resource.Attribute{
					Name:        "virtual_machine_id",
					Description: `The ID of the virtual machine that the specified Network Interface is attached to. --- A ` + "`" + `ip_configuration` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the IP Configuration.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet which the Network Interface is connected to.`,
				},
				resource.Attribute{
					Name:        "private_ip_address",
					Description: `The Private IP Address assigned to this Network Interface.`,
				},
				resource.Attribute{
					Name:        "private_ip_address_allocation",
					Description: `The IP Address allocation type for the Private address, such as ` + "`" + `Dynamic` + "`" + ` or ` + "`" + `Static` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "public_ip_address_id",
					Description: `The ID of the Public IP Address which is connected to this Network Interface.`,
				},
				resource.Attribute{
					Name:        "application_gateway_backend_address_pools_ids",
					Description: `A list of Backend Address Pool ID's within a Application Gateway that this Network Interface is connected to.`,
				},
				resource.Attribute{
					Name:        "load_balancer_backend_address_pools_ids",
					Description: `A list of Backend Address Pool ID's within a Load Balancer that this Network Interface is connected to.`,
				},
				resource.Attribute{
					Name:        "load_balancer_inbound_nat_rules_ids",
					Description: `A list of Inbound NAT Rule ID's within a Load Balancer that this Network Interface is connected to.`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `is this the Primary IP Configuration for this Network Interface? ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Network Interface.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Network Interface.`,
				},
				resource.Attribute{
					Name:        "applied_dns_servers",
					Description: `List of DNS servers applied to the specified Network Interface.`,
				},
				resource.Attribute{
					Name:        "enable_accelerated_networking",
					Description: `Indicates if accelerated networking is set on the specified Network Interface.`,
				},
				resource.Attribute{
					Name:        "enable_ip_forwarding",
					Description: `Indicate if IP forwarding is set on the specified Network Interface.`,
				},
				resource.Attribute{
					Name:        "dns_servers",
					Description: `The list of DNS servers used by the specified Network Interface.`,
				},
				resource.Attribute{
					Name:        "internal_dns_name_label",
					Description: `The internal dns name label of the specified Network Interface.`,
				},
				resource.Attribute{
					Name:        "ip_configuration",
					Description: `One or more ` + "`" + `ip_configuration` + "`" + ` blocks as defined below.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location of the specified Network Interface.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `The MAC address used by the specified Network Interface.`,
				},
				resource.Attribute{
					Name:        "network_security_group_id",
					Description: `The ID of the network security group associated to the specified Network Interface.`,
				},
				resource.Attribute{
					Name:        "private_ip_address",
					Description: `The primary private ip address associated to the specified Network Interface.`,
				},
				resource.Attribute{
					Name:        "private_ip_addresses",
					Description: `The list of private ip addresses associates to the specified Network Interface.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `List the tags associated to the specified Network Interface.`,
				},
				resource.Attribute{
					Name:        "virtual_machine_id",
					Description: `The ID of the virtual machine that the specified Network Interface is attached to. --- A ` + "`" + `ip_configuration` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the IP Configuration.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet which the Network Interface is connected to.`,
				},
				resource.Attribute{
					Name:        "private_ip_address",
					Description: `The Private IP Address assigned to this Network Interface.`,
				},
				resource.Attribute{
					Name:        "private_ip_address_allocation",
					Description: `The IP Address allocation type for the Private address, such as ` + "`" + `Dynamic` + "`" + ` or ` + "`" + `Static` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "public_ip_address_id",
					Description: `The ID of the Public IP Address which is connected to this Network Interface.`,
				},
				resource.Attribute{
					Name:        "application_gateway_backend_address_pools_ids",
					Description: `A list of Backend Address Pool ID's within a Application Gateway that this Network Interface is connected to.`,
				},
				resource.Attribute{
					Name:        "load_balancer_backend_address_pools_ids",
					Description: `A list of Backend Address Pool ID's within a Load Balancer that this Network Interface is connected to.`,
				},
				resource.Attribute{
					Name:        "load_balancer_inbound_nat_rules_ids",
					Description: `A list of Inbound NAT Rule ID's within a Load Balancer that this Network Interface is connected to.`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `is this the Primary IP Configuration for this Network Interface? ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Network Interface.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_network_security_group",
			Category:         "Network",
			ShortDescription: `Gets information about an existing Network Security Group.`,
			Description: `

Use this data source to access information about an existing Network Security Group.

`,
			Icon: "Networking/10067-icon-service-Network-Security-Groups.svg",
			Keywords: []string{
				"network",
				"security",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the Name of the Network Security Group.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `Specifies the Name of the Resource Group within which the Network Security Group exists ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Network Security Group.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The supported Azure location where the resource exists.`,
				},
				resource.Attribute{
					Name:        "security_rule",
					Description: `One or more ` + "`" + `security_rule` + "`" + ` blocks as defined below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource. The ` + "`" + `security_rule` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the security rule.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description for this rule.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The network protocol this rule applies to.`,
				},
				resource.Attribute{
					Name:        "source_port_range",
					Description: `The Source Port or Range.`,
				},
				resource.Attribute{
					Name:        "destination_port_range",
					Description: `The Destination Port or Range.`,
				},
				resource.Attribute{
					Name:        "source_address_prefix",
					Description: `CIDR or source IP range or`,
				},
				resource.Attribute{
					Name:        "source_address_prefixes",
					Description: `A list of CIDRs or source IP ranges.`,
				},
				resource.Attribute{
					Name:        "destination_address_prefix",
					Description: `CIDR or destination IP range or`,
				},
				resource.Attribute{
					Name:        "destination_address_prefixes",
					Description: `A list of CIDRs or destination IP ranges.`,
				},
				resource.Attribute{
					Name:        "source_application_security_group_ids",
					Description: `A List of source Application Security Group ID's`,
				},
				resource.Attribute{
					Name:        "destination_application_security_group_ids",
					Description: `A List of destination Application Security Group ID's`,
				},
				resource.Attribute{
					Name:        "access",
					Description: `Is network traffic is allowed or denied?`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority of the rule`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `The direction specifies if rule will be evaluated on incoming or outgoing traffic. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Network Security Group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Network Security Group.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The supported Azure location where the resource exists.`,
				},
				resource.Attribute{
					Name:        "security_rule",
					Description: `One or more ` + "`" + `security_rule` + "`" + ` blocks as defined below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource. The ` + "`" + `security_rule` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the security rule.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description for this rule.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The network protocol this rule applies to.`,
				},
				resource.Attribute{
					Name:        "source_port_range",
					Description: `The Source Port or Range.`,
				},
				resource.Attribute{
					Name:        "destination_port_range",
					Description: `The Destination Port or Range.`,
				},
				resource.Attribute{
					Name:        "source_address_prefix",
					Description: `CIDR or source IP range or`,
				},
				resource.Attribute{
					Name:        "source_address_prefixes",
					Description: `A list of CIDRs or source IP ranges.`,
				},
				resource.Attribute{
					Name:        "destination_address_prefix",
					Description: `CIDR or destination IP range or`,
				},
				resource.Attribute{
					Name:        "destination_address_prefixes",
					Description: `A list of CIDRs or destination IP ranges.`,
				},
				resource.Attribute{
					Name:        "source_application_security_group_ids",
					Description: `A List of source Application Security Group ID's`,
				},
				resource.Attribute{
					Name:        "destination_application_security_group_ids",
					Description: `A List of destination Application Security Group ID's`,
				},
				resource.Attribute{
					Name:        "access",
					Description: `Is network traffic is allowed or denied?`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority of the rule`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `The direction specifies if rule will be evaluated on incoming or outgoing traffic. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Network Security Group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_network_service_tags",
			Category:         "Network",
			ShortDescription: `Gets information about Service Tags for a specific service type.`,
			Description: `

Use this data source to access information about Service Tags.

`,
			Keywords: []string{
				"network",
				"service",
				"tags",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of this Service Tags block.`,
				},
				resource.Attribute{
					Name:        "address_prefixes",
					Description: `List of address prefixes for the service type (and optionally a specific region). ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Service Tags.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_network_watcher",
			Category:         "Network",
			ShortDescription: `Gets information about an existing Network Watcher.`,
			Description: `

Use this data source to access information about an existing Network Watcher.

`,
			Icon: "Networking/00023-icon-service-Network-Watcher.svg",
			Keywords: []string{
				"network",
				"watcher",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the Name of the Network Watcher.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `Specifies the Name of the Resource Group within which the Network Watcher exists. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Network Watcher.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The supported Azure location where the resource exists.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Network Watcher.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Network Watcher.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The supported Azure location where the resource exists.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Network Watcher.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_notification_hub",
			Category:         "Messaging",
			ShortDescription: `Gets information about an existing Notification Hub within a Notification Hub Namespace.`,
			Description: `

Use this data source to access information about an existing Notification Hub within a Notification Hub Namespace.

`,
			Icon: "IoT/10045-icon-service-Notification-Hubs.svg",
			Keywords: []string{
				"messaging",
				"notification",
				"hub",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the Name of the Notification Hub.`,
				},
				resource.Attribute{
					Name:        "namespace_name",
					Description: `Specifies the Name of the Notification Hub Namespace which contains the Notification Hub.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `Specifies the Name of the Resource Group within which the Notification Hub exists. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Notification Hub.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure Region in which this Notification Hub exists.`,
				},
				resource.Attribute{
					Name:        "apns_credential",
					Description: `A ` + "`" + `apns_credential` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "gcm_credential",
					Description: `A ` + "`" + `gcm_credential` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the resource. --- A ` + "`" + `apns_credential` + "`" + ` block exports:`,
				},
				resource.Attribute{
					Name:        "application_mode",
					Description: `The Application Mode which defines which server the APNS Messages should be sent to. Possible values are ` + "`" + `Production` + "`" + ` and ` + "`" + `Sandbox` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "bundle_id",
					Description: `The Bundle ID of the iOS/macOS application to send push notifications for, such as ` + "`" + `com.hashicorp.example` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `The Apple Push Notifications Service (APNS) Key.`,
				},
				resource.Attribute{
					Name:        "team_id",
					Description: `The ID of the team the Token.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `The Push Token associated with the Apple Developer Account. --- A ` + "`" + `gcm_credential` + "`" + ` block exports:`,
				},
				resource.Attribute{
					Name:        "api_key",
					Description: `The API Key associated with the Google Cloud Messaging service. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Notification Hub within a Notification Hub Namespace.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Notification Hub.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure Region in which this Notification Hub exists.`,
				},
				resource.Attribute{
					Name:        "apns_credential",
					Description: `A ` + "`" + `apns_credential` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "gcm_credential",
					Description: `A ` + "`" + `gcm_credential` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the resource. --- A ` + "`" + `apns_credential` + "`" + ` block exports:`,
				},
				resource.Attribute{
					Name:        "application_mode",
					Description: `The Application Mode which defines which server the APNS Messages should be sent to. Possible values are ` + "`" + `Production` + "`" + ` and ` + "`" + `Sandbox` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "bundle_id",
					Description: `The Bundle ID of the iOS/macOS application to send push notifications for, such as ` + "`" + `com.hashicorp.example` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `The Apple Push Notifications Service (APNS) Key.`,
				},
				resource.Attribute{
					Name:        "team_id",
					Description: `The ID of the team the Token.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `The Push Token associated with the Apple Developer Account. --- A ` + "`" + `gcm_credential` + "`" + ` block exports:`,
				},
				resource.Attribute{
					Name:        "api_key",
					Description: `The API Key associated with the Google Cloud Messaging service. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Notification Hub within a Notification Hub Namespace.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_notification_hub_namespace",
			Category:         "Messaging",
			ShortDescription: `Gets information about an existing Notification Hub Namespace.`,
			Description: `

Use this data source to access information about an existing Notification Hub Namespace.

`,
			Icon: "Web/10053-icon-service-Notification-Hub-Namespaces.svg",
			Keywords: []string{
				"messaging",
				"notification",
				"hub",
				"namespace",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the Name of the Notification Hub Namespace.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `Specifies the Name of the Resource Group within which the Notification Hub exists. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Notification Hub Namespace.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure Region in which this Notification Hub Namespace exists.`,
				},
				resource.Attribute{
					Name:        "namespace_type",
					Description: `The Type of Namespace, such as ` + "`" + `Messaging` + "`" + ` or ` + "`" + `NotificationHub` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sku",
					Description: `A ` + "`" + `sku` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Is this Notification Hub Namespace enabled?`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the resource. --- A ` + "`" + `sku` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the SKU to use for this Notification Hub Namespace. Possible values are ` + "`" + `Free` + "`" + `, ` + "`" + `Basic` + "`" + ` or ` + "`" + `Standard.` + "`" + ` ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Notification Hub Namespace.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Notification Hub Namespace.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure Region in which this Notification Hub Namespace exists.`,
				},
				resource.Attribute{
					Name:        "namespace_type",
					Description: `The Type of Namespace, such as ` + "`" + `Messaging` + "`" + ` or ` + "`" + `NotificationHub` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sku",
					Description: `A ` + "`" + `sku` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Is this Notification Hub Namespace enabled?`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the resource. --- A ` + "`" + `sku` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the SKU to use for this Notification Hub Namespace. Possible values are ` + "`" + `Free` + "`" + `, ` + "`" + `Basic` + "`" + ` or ` + "`" + `Standard.` + "`" + ` ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Notification Hub Namespace.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_platform_image",
			Category:         "Compute",
			ShortDescription: `Gets information about a Platform Image.`,
			Description: `

Use this data source to access information about a Platform Image.

`,
			Keywords: []string{
				"compute",
				"platform",
				"image",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "location",
					Description: `(Required) Specifies the Location to pull information about this Platform Image from.`,
				},
				resource.Attribute{
					Name:        "publisher",
					Description: `(Required) Specifies the Publisher associated with the Platform Image.`,
				},
				resource.Attribute{
					Name:        "offer",
					Description: `(Required) Specifies the Offer associated with the Platform Image.`,
				},
				resource.Attribute{
					Name:        "sku",
					Description: `(Required) Specifies the SKU of the Platform Image.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) The version of the Platform Image. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Platform Image. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Platform Image.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Platform Image. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Platform Image.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_policy_definition",
			Category:         "Policy",
			ShortDescription: `Get information about a Policy Definition.`,
			Description: `

Use this data source to access information about a Policy Definition, both custom and built in. Retrieves Policy Definitions from your current subscription by default.

`,
			Keywords: []string{
				"policy",
				"definition",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name of the Policy Definition. Conflicts with ` + "`" + `display_name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Specifies the display name of the Policy Definition. Conflicts with ` + "`" + `name` + "`" + `. ~>`,
				},
				resource.Attribute{
					Name:        "management_group_name",
					Description: `(Optional) Only retrieve Policy Definitions from this Management Group. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Policy Definition.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The Type of Policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The Description of the Policy.`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `The Type of the Policy. Possible values are "BuiltIn", "Custom" and "NotSpecified".`,
				},
				resource.Attribute{
					Name:        "policy_rule",
					Description: `The Rule as defined (in JSON) in the Policy.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `Any Parameters defined in the Policy.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `Any Metadata defined in the Policy. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Policy Definition.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Policy Definition.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The Type of Policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The Description of the Policy.`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `The Type of the Policy. Possible values are "BuiltIn", "Custom" and "NotSpecified".`,
				},
				resource.Attribute{
					Name:        "policy_rule",
					Description: `The Rule as defined (in JSON) in the Policy.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `Any Parameters defined in the Policy.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `Any Metadata defined in the Policy. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Policy Definition.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_policy_set_definition",
			Category:         "Policy",
			ShortDescription: `Get information about a Policy Set Definition.`,
			Description: `

Use this data source to access information about an existing Policy Set Definition.

`,
			Keywords: []string{
				"policy",
				"set",
				"definition",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name of the Policy Set Definition. Conflicts with ` + "`" + `display_name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Specifies the display name of the Policy Set Definition. Conflicts with ` + "`" + `name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "management_group_name",
					Description: `(Optional) Only retrieve Policy Set Definitions from this Management Group. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Policy Set Definition.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The Description of the Policy Set Definition.`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `The Type of the Policy Set Definition.`,
				},
				resource.Attribute{
					Name:        "policy_definitions",
					Description: `The policy definitions contained within the policy set definition.`,
				},
				resource.Attribute{
					Name:        "policy_definition_reference",
					Description: `One or more ` + "`" + `policy_definition_reference` + "`" + ` blocks as defined below.`,
				},
				resource.Attribute{
					Name:        "policy_definition_group",
					Description: `One or more ` + "`" + `policy_definition_group` + "`" + ` blocks as defined below.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `Any Parameters defined in the Policy Set Definition.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `Any Metadata defined in the Policy Set Definition. --- An ` + "`" + `policy_definition_reference` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "policy_definition_id",
					Description: `The ID of the policy definition or policy set definition that is included in this policy set definition.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `The mapping of the parameter values for the referenced policy rule. The keys are the parameter names. ->`,
				},
				resource.Attribute{
					Name:        "parameter_values",
					Description: `The parameter values for the referenced policy rule. This field is a json object.`,
				},
				resource.Attribute{
					Name:        "reference_id",
					Description: `The unique ID within this policy set definition for this policy definition reference.`,
				},
				resource.Attribute{
					Name:        "group_names",
					Description: `The list of names of the policy definition groups that this policy definition reference belongs to. --- An ` + "`" + `policy_definition_group` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of this policy definition group.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name of this policy definition group.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `The category of this policy definition group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of this policy definition group.`,
				},
				resource.Attribute{
					Name:        "additional_metadata_id",
					Description: `The ID of a resource that contains additional metadata about this policy definition group. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Policy Set Definition.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Policy Set Definition.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The Description of the Policy Set Definition.`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `The Type of the Policy Set Definition.`,
				},
				resource.Attribute{
					Name:        "policy_definitions",
					Description: `The policy definitions contained within the policy set definition.`,
				},
				resource.Attribute{
					Name:        "policy_definition_reference",
					Description: `One or more ` + "`" + `policy_definition_reference` + "`" + ` blocks as defined below.`,
				},
				resource.Attribute{
					Name:        "policy_definition_group",
					Description: `One or more ` + "`" + `policy_definition_group` + "`" + ` blocks as defined below.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `Any Parameters defined in the Policy Set Definition.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `Any Metadata defined in the Policy Set Definition. --- An ` + "`" + `policy_definition_reference` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "policy_definition_id",
					Description: `The ID of the policy definition or policy set definition that is included in this policy set definition.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `The mapping of the parameter values for the referenced policy rule. The keys are the parameter names. ->`,
				},
				resource.Attribute{
					Name:        "parameter_values",
					Description: `The parameter values for the referenced policy rule. This field is a json object.`,
				},
				resource.Attribute{
					Name:        "reference_id",
					Description: `The unique ID within this policy set definition for this policy definition reference.`,
				},
				resource.Attribute{
					Name:        "group_names",
					Description: `The list of names of the policy definition groups that this policy definition reference belongs to. --- An ` + "`" + `policy_definition_group` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of this policy definition group.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name of this policy definition group.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `The category of this policy definition group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of this policy definition group.`,
				},
				resource.Attribute{
					Name:        "additional_metadata_id",
					Description: `The ID of a resource that contains additional metadata about this policy definition group. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Policy Set Definition.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_postgresql_server",
			Category:         "Database",
			ShortDescription: `Gets information about an existing PostgreSQL Azure Database Server.`,
			Description: `

Use this data source to access information about an existing PostgreSQL Azure Database Server.

`,
			Icon: "Databases/10131-icon-service-Azure-Database-PostgreSQL-Server.svg",
			Keywords: []string{
				"database",
				"postgresql",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the PostgreSQL Server.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `Specifies the name of the Resource Group where the PostgreSQL Server exists. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location of the Resource Group in which the PostgreSQL Server exists.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `The fully qualified domain name of the PostgreSQL Server.`,
				},
				resource.Attribute{
					Name:        "identity",
					Description: `An ` + "`" + `identity` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version of the PostgreSQL Server.`,
				},
				resource.Attribute{
					Name:        "administrator_login",
					Description: `The administrator username of the PostgreSQL Server.`,
				},
				resource.Attribute{
					Name:        "sku_name",
					Description: `The SKU name of the PostgreSQL Server.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource. --- An ` + "`" + `identity` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "principal_id",
					Description: `The ID of the System Managed Service Principal assigned to the PostgreSQL Server.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `The ID of the Tenant of the System Managed Service Principal assigned to the PostgreSQL Server.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The identity type of the Managed Identity assigned to the PostgreSQL Server. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the PostgreSQL Azure Database Server.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "location",
					Description: `The location of the Resource Group in which the PostgreSQL Server exists.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `The fully qualified domain name of the PostgreSQL Server.`,
				},
				resource.Attribute{
					Name:        "identity",
					Description: `An ` + "`" + `identity` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version of the PostgreSQL Server.`,
				},
				resource.Attribute{
					Name:        "administrator_login",
					Description: `The administrator username of the PostgreSQL Server.`,
				},
				resource.Attribute{
					Name:        "sku_name",
					Description: `The SKU name of the PostgreSQL Server.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource. --- An ` + "`" + `identity` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "principal_id",
					Description: `The ID of the System Managed Service Principal assigned to the PostgreSQL Server.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `The ID of the Tenant of the System Managed Service Principal assigned to the PostgreSQL Server.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The identity type of the Managed Identity assigned to the PostgreSQL Server. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the PostgreSQL Azure Database Server.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_private_dns_zone",
			Category:         "Private DNS",
			ShortDescription: `Gets information about an existing Private DNS Zone.`,
			Description: `

Use this data source to access information about an existing Private DNS Zone.

`,
			Keywords: []string{
				"private",
				"dns",
				"zone",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Private DNS Zone.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Optional) The Name of the Resource Group where the Private DNS Zone exists. If the Name of the Resource Group is not provided, the first Private DNS Zone from the list of Private DNS Zones in your subscription that matches ` + "`" + `name` + "`" + ` will be returned. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Private DNS Zone.`,
				},
				resource.Attribute{
					Name:        "max_number_of_record_sets",
					Description: `Maximum number of recordsets that can be created in this Private Zone.`,
				},
				resource.Attribute{
					Name:        "max_number_of_virtual_network_links",
					Description: `Maximum number of Virtual Networks that can be linked to this Private Zone.`,
				},
				resource.Attribute{
					Name:        "max_number_of_virtual_network_links_with_registration",
					Description: `Maximum number of Virtual Networks that can be linked to this Private Zone with registration enabled.`,
				},
				resource.Attribute{
					Name:        "number_of_record_sets",
					Description: `The number of recordsets currently in the zone.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags for the zone. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Private DNS Zone.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Private DNS Zone.`,
				},
				resource.Attribute{
					Name:        "max_number_of_record_sets",
					Description: `Maximum number of recordsets that can be created in this Private Zone.`,
				},
				resource.Attribute{
					Name:        "max_number_of_virtual_network_links",
					Description: `Maximum number of Virtual Networks that can be linked to this Private Zone.`,
				},
				resource.Attribute{
					Name:        "max_number_of_virtual_network_links_with_registration",
					Description: `Maximum number of Virtual Networks that can be linked to this Private Zone with registration enabled.`,
				},
				resource.Attribute{
					Name:        "number_of_record_sets",
					Description: `The number of recordsets currently in the zone.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags for the zone. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Private DNS Zone.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_private_endpoint_connection",
			Category:         "Network",
			ShortDescription: `Gets the connection status information about an existing Private Endpoint`,
			Description: `

Use this data source to access the connection status information about an existing Private Endpoint Connection.

`,
			Keywords: []string{
				"network",
				"private",
				"endpoint",
				"connection",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the Name of the private endpoint.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `Specifies the Name of the Resource Group within which the private endpoint exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Private Endpoint.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The supported Azure location where the resource exists. A ` + "`" + `private_service_connection` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the private endpoint.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current status of the private endpoint request, possible values will be ` + "`" + `Pending` + "`" + `, ` + "`" + `Approved` + "`" + `, ` + "`" + `Rejected` + "`" + `, or ` + "`" + `Disconnected` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "private_ip_address",
					Description: `The private IP address associated with the private endpoint, note that you will have a private IP address assigned to the private endpoint even if the connection request was ` + "`" + `Rejected` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "request_response",
					Description: `Possible values are as follows: Value | Meaning -- | -- ` + "`" + `Auto-Approved` + "`" + ` | The remote resource owner has added you to the ` + "`" + `Auto-Approved` + "`" + ` RBAC permission list for the remote resource, all private endpoint connection requests will be automatically ` + "`" + `Approved` + "`" + `. ` + "`" + `Deleted state` + "`" + ` | The resource owner has ` + "`" + `Rejected` + "`" + ` the private endpoint connection request and has removed your private endpoint request from the remote resource. ` + "`" + `request/response message` + "`" + ` | If you submitted a manual private endpoint connection request, while in the ` + "`" + `Pending` + "`" + ` status the ` + "`" + `request_response` + "`" + ` will display the same text from your ` + "`" + `request_message` + "`" + ` in the ` + "`" + `private_service_connection` + "`" + ` block above. If the private endpoint connection request was ` + "`" + `Rejected` + "`" + ` by the owner of the remote resource, the text for the rejection will be displayed as the ` + "`" + `request_response` + "`" + ` text, if the private endpoint connection request was ` + "`" + `Approved` + "`" + ` by the owner of the remote resource, the text for the approval will be displayed as the ` + "`" + `request_response` + "`" + ` text ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Private Endpoint.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Private Endpoint.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The supported Azure location where the resource exists. A ` + "`" + `private_service_connection` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the private endpoint.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current status of the private endpoint request, possible values will be ` + "`" + `Pending` + "`" + `, ` + "`" + `Approved` + "`" + `, ` + "`" + `Rejected` + "`" + `, or ` + "`" + `Disconnected` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "private_ip_address",
					Description: `The private IP address associated with the private endpoint, note that you will have a private IP address assigned to the private endpoint even if the connection request was ` + "`" + `Rejected` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "request_response",
					Description: `Possible values are as follows: Value | Meaning -- | -- ` + "`" + `Auto-Approved` + "`" + ` | The remote resource owner has added you to the ` + "`" + `Auto-Approved` + "`" + ` RBAC permission list for the remote resource, all private endpoint connection requests will be automatically ` + "`" + `Approved` + "`" + `. ` + "`" + `Deleted state` + "`" + ` | The resource owner has ` + "`" + `Rejected` + "`" + ` the private endpoint connection request and has removed your private endpoint request from the remote resource. ` + "`" + `request/response message` + "`" + ` | If you submitted a manual private endpoint connection request, while in the ` + "`" + `Pending` + "`" + ` status the ` + "`" + `request_response` + "`" + ` will display the same text from your ` + "`" + `request_message` + "`" + ` in the ` + "`" + `private_service_connection` + "`" + ` block above. If the private endpoint connection request was ` + "`" + `Rejected` + "`" + ` by the owner of the remote resource, the text for the rejection will be displayed as the ` + "`" + `request_response` + "`" + ` text, if the private endpoint connection request was ` + "`" + `Approved` + "`" + ` by the owner of the remote resource, the text for the approval will be displayed as the ` + "`" + `request_response` + "`" + ` text ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Private Endpoint.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_private_link_service",
			Category:         "Network",
			ShortDescription: `Use this data source to access information about an existing Private Link Service.`,
			Description: `

Use this data source to access information about an existing Private Link Service.

`,
			Icon: "Networking/01105-icon-service-Private-Link-Service.svg",
			Keywords: []string{
				"network",
				"private",
				"link",
				"service",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the private link service.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of the resource group in which the private link service resides. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Azure resource ID of the Private Link Service.`,
				},
				resource.Attribute{
					Name:        "alias",
					Description: `The alias is a globally unique name for your private link service which Azure generates for you. Your can use this alias to request a connection to your private link service.`,
				},
				resource.Attribute{
					Name:        "auto_approval_subscription_ids",
					Description: `The list of subscription(s) globally unique identifiers that will be auto approved to use the private link service.`,
				},
				resource.Attribute{
					Name:        "enable_proxy_protocol",
					Description: `Does the Private Link Service support the Proxy Protocol?`,
				},
				resource.Attribute{
					Name:        "load_balancer_frontend_ip_configuration_ids",
					Description: `The list of Standard Load Balancer(SLB) resource IDs. The Private Link service is tied to the frontend IP address of a SLB. All traffic destined for the private link service will reach the frontend of the SLB. You can configure SLB rules to direct this traffic to appropriate backend pools where your applications are running.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The supported Azure location where the resource exists.`,
				},
				resource.Attribute{
					Name:        "nat_ip_configuration",
					Description: `The ` + "`" + `nat_ip_configuration` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "visibility_subscription_ids",
					Description: `The list of subscription(s) globally unique identifiers(GUID) that will be able to see the private link service. --- The ` + "`" + `nat_ip_configuration` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of private link service NAT IP configuration.`,
				},
				resource.Attribute{
					Name:        "private_ip_address",
					Description: `The private IP address of the NAT IP configuration.`,
				},
				resource.Attribute{
					Name:        "private_ip_address_version",
					Description: `The version of the IP Protocol.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the subnet to be used by the service.`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `Value that indicates if the IP configuration is the primary configuration or not. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Private Link Service.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Azure resource ID of the Private Link Service.`,
				},
				resource.Attribute{
					Name:        "alias",
					Description: `The alias is a globally unique name for your private link service which Azure generates for you. Your can use this alias to request a connection to your private link service.`,
				},
				resource.Attribute{
					Name:        "auto_approval_subscription_ids",
					Description: `The list of subscription(s) globally unique identifiers that will be auto approved to use the private link service.`,
				},
				resource.Attribute{
					Name:        "enable_proxy_protocol",
					Description: `Does the Private Link Service support the Proxy Protocol?`,
				},
				resource.Attribute{
					Name:        "load_balancer_frontend_ip_configuration_ids",
					Description: `The list of Standard Load Balancer(SLB) resource IDs. The Private Link service is tied to the frontend IP address of a SLB. All traffic destined for the private link service will reach the frontend of the SLB. You can configure SLB rules to direct this traffic to appropriate backend pools where your applications are running.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The supported Azure location where the resource exists.`,
				},
				resource.Attribute{
					Name:        "nat_ip_configuration",
					Description: `The ` + "`" + `nat_ip_configuration` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "visibility_subscription_ids",
					Description: `The list of subscription(s) globally unique identifiers(GUID) that will be able to see the private link service. --- The ` + "`" + `nat_ip_configuration` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of private link service NAT IP configuration.`,
				},
				resource.Attribute{
					Name:        "private_ip_address",
					Description: `The private IP address of the NAT IP configuration.`,
				},
				resource.Attribute{
					Name:        "private_ip_address_version",
					Description: `The version of the IP Protocol.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the subnet to be used by the service.`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `Value that indicates if the IP configuration is the primary configuration or not. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Private Link Service.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_private_link_service_endpoint_connections",
			Category:         "Network",
			ShortDescription: `Use this data source to access endpoint connection information about an existing Private Link Service.`,
			Description: `

Use this data source to access endpoint connection information about an existing Private Link Service.

`,
			Keywords: []string{
				"network",
				"private",
				"link",
				"service",
				"endpoint",
				"connections",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_id",
					Description: `The resource ID of the private link service.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of the resource group in which the private link service resides. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `The name of the private link service. The ` + "`" + `private_endpoint_connections` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "connection_id",
					Description: `The resource id of the private link service connection between the private link service and the private link endpoint.`,
				},
				resource.Attribute{
					Name:        "connection_name",
					Description: `The name of the connection between the private link service and the private link endpoint.`,
				},
				resource.Attribute{
					Name:        "private_endpoint_id",
					Description: `The resource id of the private link endpoint.`,
				},
				resource.Attribute{
					Name:        "private_endpoint_name",
					Description: `The name of the private link endpoint.`,
				},
				resource.Attribute{
					Name:        "action_required",
					Description: `A message indicating if changes on the service provider require any updates or not.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The request for approval message or the reason for rejection message.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates the state of the connection between the private link service and the private link endpoint, possible values are ` + "`" + `Pending` + "`" + `, ` + "`" + `Approved` + "`" + ` or ` + "`" + `Rejected` + "`" + `. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Private Link Service.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `The name of the private link service. The ` + "`" + `private_endpoint_connections` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "connection_id",
					Description: `The resource id of the private link service connection between the private link service and the private link endpoint.`,
				},
				resource.Attribute{
					Name:        "connection_name",
					Description: `The name of the connection between the private link service and the private link endpoint.`,
				},
				resource.Attribute{
					Name:        "private_endpoint_id",
					Description: `The resource id of the private link endpoint.`,
				},
				resource.Attribute{
					Name:        "private_endpoint_name",
					Description: `The name of the private link endpoint.`,
				},
				resource.Attribute{
					Name:        "action_required",
					Description: `A message indicating if changes on the service provider require any updates or not.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The request for approval message or the reason for rejection message.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates the state of the connection between the private link service and the private link endpoint, possible values are ` + "`" + `Pending` + "`" + `, ` + "`" + `Approved` + "`" + ` or ` + "`" + `Rejected` + "`" + `. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Private Link Service.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_proximity_placement_group",
			Category:         "Compute",
			ShortDescription: `Gets information about an existing Proximity Placement Group.`,
			Description: `

Use this data source to access information about an existing Proximity Placement Group.

`,
			Icon: "Networking/10365-icon-service-Proximity-Placement-Groups.svg",
			Keywords: []string{
				"compute",
				"proximity",
				"placement",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Proximity Placement Group.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of the resource group in which the Proximity Placement Group exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Proximity Placement Group. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Proximity Placement Group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Proximity Placement Group. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Proximity Placement Group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_public_ip",
			Category:         "Network",
			ShortDescription: `Gets information about an existing Public IP Address.`,
			Description: `

Use this data source to access information about an existing Public IP Address.

`,
			Keywords: []string{
				"network",
				"public",
				"ip",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name of the public IP address.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `Specifies the name of the resource group. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "domain_name_label",
					Description: `The label for the Domain Name.`,
				},
				resource.Attribute{
					Name:        "idle_timeout_in_minutes",
					Description: `Specifies the timeout for the TCP idle connection.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `Fully qualified domain name of the A DNS record associated with the public IP. This is the concatenation of the domainNameLabel and the regionalized DNS zone.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The IP address value that was allocated.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `The IP version being used, for example ` + "`" + `IPv4` + "`" + ` or ` + "`" + `IPv6` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sku",
					Description: `The SKU of the Public IP.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assigned to the resource. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Public IP Address.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "domain_name_label",
					Description: `The label for the Domain Name.`,
				},
				resource.Attribute{
					Name:        "idle_timeout_in_minutes",
					Description: `Specifies the timeout for the TCP idle connection.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `Fully qualified domain name of the A DNS record associated with the public IP. This is the concatenation of the domainNameLabel and the regionalized DNS zone.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The IP address value that was allocated.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `The IP version being used, for example ` + "`" + `IPv4` + "`" + ` or ` + "`" + `IPv6` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sku",
					Description: `The SKU of the Public IP.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assigned to the resource. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Public IP Address.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_public_ip_prefix",
			Category:         "Network",
			ShortDescription: `Gets information about an existing Public IP Prefix.`,
			Description: `

Use this data source to access information about an existing Public IP Prefix.

`,
			Icon: "Networking/10372-icon-service-Public-IP-Prefixes.svg",
			Keywords: []string{
				"network",
				"public",
				"ip",
				"prefix",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name of the public IP prefix.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `Specifies the name of the resource group. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Public IP prefix resource.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of the resource group in which to create the public IP.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The supported Azure location where the resource exists.`,
				},
				resource.Attribute{
					Name:        "sku",
					Description: `The SKU of the Public IP Prefix.`,
				},
				resource.Attribute{
					Name:        "prefix_length",
					Description: `The number of bits of the prefix.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assigned to the resource. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Public IP Prefix.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Public IP prefix resource.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of the resource group in which to create the public IP.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The supported Azure location where the resource exists.`,
				},
				resource.Attribute{
					Name:        "sku",
					Description: `The SKU of the Public IP Prefix.`,
				},
				resource.Attribute{
					Name:        "prefix_length",
					Description: `The number of bits of the prefix.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assigned to the resource. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Public IP Prefix.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_public_ips",
			Category:         "Network",
			ShortDescription: `Gets information about a set of existing Public IP Addresses.`,
			Description: `

Use this data source to access information about a set of existing Public IP Addresses.

`,
			Keywords: []string{
				"network",
				"public",
				"ips",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `Specifies the name of the resource group.`,
				},
				resource.Attribute{
					Name:        "attached",
					Description: `(Optional) Filter to include IP Addresses which are attached to a device, such as a VM/LB (` + "`" + `true` + "`" + `) or unattached (` + "`" + `false` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "name_prefix",
					Description: `(Optional) A prefix match used for the IP Addresses ` + "`" + `name` + "`" + ` field, case sensitive.`,
				},
				resource.Attribute{
					Name:        "allocation_type",
					Description: `(Optional) The Allocation Type for the Public IP Address. Possible values include ` + "`" + `Static` + "`" + ` or ` + "`" + `Dynamic` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "public_ips",
					Description: `A List of ` + "`" + `public_ips` + "`" + ` blocks as defined below filtered by the criteria above. A ` + "`" + `public_ips` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Public IP Address`,
				},
				resource.Attribute{
					Name:        "domain_name_label",
					Description: `The Domain Name Label of the Public IP Address`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `The FQDN of the Public IP Address`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The Name of the Public IP Address`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The IP address of the Public IP Address ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Public IP Addresses.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "public_ips",
					Description: `A List of ` + "`" + `public_ips` + "`" + ` blocks as defined below filtered by the criteria above. A ` + "`" + `public_ips` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Public IP Address`,
				},
				resource.Attribute{
					Name:        "domain_name_label",
					Description: `The Domain Name Label of the Public IP Address`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `The FQDN of the Public IP Address`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The Name of the Public IP Address`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The IP address of the Public IP Address ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Public IP Addresses.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_recovery_services_vault",
			Category:         "Recovery Services",
			ShortDescription: `Gets information about an existing Recovery Services Vault.`,
			Description: `

Use this data source to access information about an existing Recovery Services Vault.

`,
			Icon: "Storage/00017-icon-service-Recovery-Services-Vaults.svg",
			Keywords: []string{
				"recovery",
				"services",
				"vault",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name of the Recovery Services Vault.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of the resource group in which the Recovery Services Vault resides. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Recovery Services Vault.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the resource resides.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource.`,
				},
				resource.Attribute{
					Name:        "sku",
					Description: `The vault's current SKU. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Recovery Services Vault.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Recovery Services Vault.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the resource resides.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource.`,
				},
				resource.Attribute{
					Name:        "sku",
					Description: `The vault's current SKU. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Recovery Services Vault.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_redis_cache",
			Category:         "Redis",
			ShortDescription: `Gets information about an existing Azure Redis Cache.`,
			Description: `

Use this data source to access information about an existing Redis Cache

`,
			Icon: "Databases/10137-icon-service-Cache-Redis.svg",
			Keywords: []string{
				"redis",
				"cache",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Redis cache`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of the resource group the Redis cache instance is located in. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Cache ID.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location of the Redis Cache.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `The size of the Redis Cache deployed.`,
				},
				resource.Attribute{
					Name:        "family",
					Description: `The SKU family/pricing group used. Possible values are ` + "`" + `C` + "`" + ` (for Basic/Standard SKU family) and ` + "`" + `P` + "`" + ` (for ` + "`" + `Premium` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "sku_name",
					Description: `The SKU of Redis used. Possible values are ` + "`" + `Basic` + "`" + `, ` + "`" + `Standard` + "`" + ` and ` + "`" + `Premium` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enable_non_ssl_port",
					Description: `Whether the SSL port is enabled.`,
				},
				resource.Attribute{
					Name:        "minimum_tls_version",
					Description: `The minimum TLS version.`,
				},
				resource.Attribute{
					Name:        "patch_schedule",
					Description: `A list of ` + "`" + `patch_schedule` + "`" + ` blocks as defined below.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `The Hostname of the Redis Instance`,
				},
				resource.Attribute{
					Name:        "ssl_port",
					Description: `The SSL Port of the Redis Instance`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The non-SSL Port of the Redis Instance`,
				},
				resource.Attribute{
					Name:        "primary_access_key",
					Description: `The Primary Access Key for the Redis Instance`,
				},
				resource.Attribute{
					Name:        "secondary_access_key",
					Description: `The Secondary Access Key for the Redis Instance`,
				},
				resource.Attribute{
					Name:        "primary_connection_string",
					Description: `The primary connection string of the Redis Instance.`,
				},
				resource.Attribute{
					Name:        "secondary_connection_string",
					Description: `The secondary connection string of the Redis Instance.`,
				},
				resource.Attribute{
					Name:        "redis_configuration",
					Description: `A ` + "`" + `redis_configuration` + "`" + ` block as defined below. --- A ` + "`" + `patch_schedule` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "day_of_week",
					Description: `the Weekday name for the patch item`,
				},
				resource.Attribute{
					Name:        "start_hour_utc",
					Description: `The Start Hour for maintenance in UTC ~>`,
				},
				resource.Attribute{
					Name:        "enable_authentication",
					Description: `Specifies if authentication is enabled`,
				},
				resource.Attribute{
					Name:        "maxmemory_reserved",
					Description: `The value in megabytes reserved for non-cache usage e.g. failover`,
				},
				resource.Attribute{
					Name:        "maxmemory_delta",
					Description: `The max-memory delta for this Redis instance.`,
				},
				resource.Attribute{
					Name:        "maxmemory_policy",
					Description: `How Redis will select what to remove when ` + "`" + `maxmemory` + "`" + ` is reached.`,
				},
				resource.Attribute{
					Name:        "maxfragmentationmemory_reserved",
					Description: `Value in megabytes reserved to accommodate for memory fragmentation.`,
				},
				resource.Attribute{
					Name:        "rdb_backup_enabled",
					Description: `Is Backup Enabled? Only supported on Premium SKU's.`,
				},
				resource.Attribute{
					Name:        "rdb_backup_frequency",
					Description: `The Backup Frequency in Minutes. Only supported on Premium SKU's.`,
				},
				resource.Attribute{
					Name:        "rdb_backup_max_snapshot_count",
					Description: `The maximum number of snapshots that can be created as a backup.`,
				},
				resource.Attribute{
					Name:        "rdb_storage_connection_string",
					Description: `The Connection String to the Storage Account. Only supported for Premium SKU's. ~>`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Redis Cache.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Cache ID.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location of the Redis Cache.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `The size of the Redis Cache deployed.`,
				},
				resource.Attribute{
					Name:        "family",
					Description: `The SKU family/pricing group used. Possible values are ` + "`" + `C` + "`" + ` (for Basic/Standard SKU family) and ` + "`" + `P` + "`" + ` (for ` + "`" + `Premium` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "sku_name",
					Description: `The SKU of Redis used. Possible values are ` + "`" + `Basic` + "`" + `, ` + "`" + `Standard` + "`" + ` and ` + "`" + `Premium` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enable_non_ssl_port",
					Description: `Whether the SSL port is enabled.`,
				},
				resource.Attribute{
					Name:        "minimum_tls_version",
					Description: `The minimum TLS version.`,
				},
				resource.Attribute{
					Name:        "patch_schedule",
					Description: `A list of ` + "`" + `patch_schedule` + "`" + ` blocks as defined below.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `The Hostname of the Redis Instance`,
				},
				resource.Attribute{
					Name:        "ssl_port",
					Description: `The SSL Port of the Redis Instance`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The non-SSL Port of the Redis Instance`,
				},
				resource.Attribute{
					Name:        "primary_access_key",
					Description: `The Primary Access Key for the Redis Instance`,
				},
				resource.Attribute{
					Name:        "secondary_access_key",
					Description: `The Secondary Access Key for the Redis Instance`,
				},
				resource.Attribute{
					Name:        "primary_connection_string",
					Description: `The primary connection string of the Redis Instance.`,
				},
				resource.Attribute{
					Name:        "secondary_connection_string",
					Description: `The secondary connection string of the Redis Instance.`,
				},
				resource.Attribute{
					Name:        "redis_configuration",
					Description: `A ` + "`" + `redis_configuration` + "`" + ` block as defined below. --- A ` + "`" + `patch_schedule` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "day_of_week",
					Description: `the Weekday name for the patch item`,
				},
				resource.Attribute{
					Name:        "start_hour_utc",
					Description: `The Start Hour for maintenance in UTC ~>`,
				},
				resource.Attribute{
					Name:        "enable_authentication",
					Description: `Specifies if authentication is enabled`,
				},
				resource.Attribute{
					Name:        "maxmemory_reserved",
					Description: `The value in megabytes reserved for non-cache usage e.g. failover`,
				},
				resource.Attribute{
					Name:        "maxmemory_delta",
					Description: `The max-memory delta for this Redis instance.`,
				},
				resource.Attribute{
					Name:        "maxmemory_policy",
					Description: `How Redis will select what to remove when ` + "`" + `maxmemory` + "`" + ` is reached.`,
				},
				resource.Attribute{
					Name:        "maxfragmentationmemory_reserved",
					Description: `Value in megabytes reserved to accommodate for memory fragmentation.`,
				},
				resource.Attribute{
					Name:        "rdb_backup_enabled",
					Description: `Is Backup Enabled? Only supported on Premium SKU's.`,
				},
				resource.Attribute{
					Name:        "rdb_backup_frequency",
					Description: `The Backup Frequency in Minutes. Only supported on Premium SKU's.`,
				},
				resource.Attribute{
					Name:        "rdb_backup_max_snapshot_count",
					Description: `The maximum number of snapshots that can be created as a backup.`,
				},
				resource.Attribute{
					Name:        "rdb_storage_connection_string",
					Description: `The Connection String to the Storage Account. Only supported for Premium SKU's. ~>`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Redis Cache.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_resource_group",
			Category:         "Base",
			ShortDescription: `Gets information about an existing Resource Group.`,
			Description: `

Use this data source to access information about an existing Resource Group.

`,
			Keywords: []string{
				"base",
				"resource",
				"group",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Resource Group.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure Region where the Resource Group exists.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the Resource Group. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Resource Group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_resources",
			Category:         "Base",
			ShortDescription: `Gets information about an existing Resources.`,
			Description: `

Use this data source to access information about existing resources.

`,
			Keywords: []string{
				"base",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the Resource.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Optional) The name of the Resource group where the Resources are located.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The Resource Type of the Resources you want to list (e.g. ` + "`" + `Microsoft.Network/virtualNetworks` + "`" + `). A full list of available Resource Types can be found [here](https://docs.microsoft.com/en-us/azure/azure-resource-manager/azure-services-resource-providers).`,
				},
				resource.Attribute{
					Name:        "required_tags",
					Description: `(Optional) A mapping of tags which the resource has to have in order to be included in the result. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `One or more ` + "`" + `resource` + "`" + ` blocks as defined below. --- The ` + "`" + `resource` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of this Resource.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of this Resource.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of this Resource. (e.g. ` + "`" + `Microsoft.Network/virtualNetworks` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure Region in which this Resource exists.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A map of tags assigned to this Resource. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Resources.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "resources",
					Description: `One or more ` + "`" + `resource` + "`" + ` blocks as defined below. --- The ` + "`" + `resource` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of this Resource.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of this Resource.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of this Resource. (e.g. ` + "`" + `Microsoft.Network/virtualNetworks` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure Region in which this Resource exists.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A map of tags assigned to this Resource. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Resources.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_role_definition",
			Category:         "Authorization",
			ShortDescription: `Get information about an existing Role Definition.`,
			Description: `

Use this data source to access information about an existing Role Definition.

`,
			Keywords: []string{
				"authorization",
				"role",
				"definition",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Specifies the Name of either a built-in or custom Role Definition. -> You can also use this for built-in roles such as ` + "`" + `Contributor` + "`" + `, ` + "`" + `Owner` + "`" + `, ` + "`" + `Reader` + "`" + ` and ` + "`" + `Virtual Machine Contributor` + "`" + ``,
				},
				resource.Attribute{
					Name:        "role_definition_id",
					Description: `(Optional) Specifies the ID of the Role Definition as a UUID/GUID.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) Specifies the Scope at which the Custom Role Definition exists. ~>`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `the ID of the built-in Role Definition.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `the Description of the built-in Role.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `the Type of the Role.`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `a ` + "`" + `permissions` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "assignable_scopes",
					Description: `One or more assignable scopes for this Role Definition, such as ` + "`" + `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333` + "`" + `, ` + "`" + `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup` + "`" + `, or ` + "`" + `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM` + "`" + `. A ` + "`" + `permissions` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "actions",
					Description: `a list of actions supported by this role`,
				},
				resource.Attribute{
					Name:        "not_actions",
					Description: `a list of actions which are denied by this role ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Role Definition.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `the ID of the built-in Role Definition.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `the Description of the built-in Role.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `the Type of the Role.`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `a ` + "`" + `permissions` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "assignable_scopes",
					Description: `One or more assignable scopes for this Role Definition, such as ` + "`" + `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333` + "`" + `, ` + "`" + `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup` + "`" + `, or ` + "`" + `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM` + "`" + `. A ` + "`" + `permissions` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "actions",
					Description: `a list of actions supported by this role`,
				},
				resource.Attribute{
					Name:        "not_actions",
					Description: `a list of actions which are denied by this role ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Role Definition.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_route_filter",
			Category:         "Network",
			ShortDescription: `Gets information about an existing Route Filter.`,
			Description: `

Use this data source to access information about an existing Route Filter.

`,
			Icon: "Networking/10071-icon-service-Route-Filters.svg",
			Keywords: []string{
				"network",
				"route",
				"filter",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Route Filter.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure Region where the Route Filter exists.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `A ` + "`" + `rule` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the Route Filter. --- A ` + "`" + `rule` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "access",
					Description: `The access type of the rule`,
				},
				resource.Attribute{
					Name:        "communities",
					Description: `The collection for bgp community values.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The Name of Route Filter Rule`,
				},
				resource.Attribute{
					Name:        "rule_type",
					Description: `The Route Filter Rule Type. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Route Filter.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_route_table",
			Category:         "Network",
			ShortDescription: `Gets information about an existing Route Table`,
			Description: `

Use this data source to access information about an existing Route Table.

`,
			Keywords: []string{
				"network",
				"route",
				"table",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Route Table.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of the Resource Group in which the Route Table exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Route Table ID.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure Region in which the Route Table exists.`,
				},
				resource.Attribute{
					Name:        "route",
					Description: `One or more ` + "`" + `route` + "`" + ` blocks as documented below.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `The collection of Subnets associated with this route table.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the Route Table. The ` + "`" + `route` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Route.`,
				},
				resource.Attribute{
					Name:        "address_prefix",
					Description: `The destination CIDR to which the route applies.`,
				},
				resource.Attribute{
					Name:        "next_hop_type",
					Description: `The type of Azure hop the packet should be sent to.`,
				},
				resource.Attribute{
					Name:        "next_hop_in_ip_address",
					Description: `Contains the IP address packets should be forwarded to. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Route Table.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Route Table ID.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure Region in which the Route Table exists.`,
				},
				resource.Attribute{
					Name:        "route",
					Description: `One or more ` + "`" + `route` + "`" + ` blocks as documented below.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `The collection of Subnets associated with this route table.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the Route Table. The ` + "`" + `route` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Route.`,
				},
				resource.Attribute{
					Name:        "address_prefix",
					Description: `The destination CIDR to which the route applies.`,
				},
				resource.Attribute{
					Name:        "next_hop_type",
					Description: `The type of Azure hop the packet should be sent to.`,
				},
				resource.Attribute{
					Name:        "next_hop_in_ip_address",
					Description: `Contains the IP address packets should be forwarded to. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Route Table.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_search_service",
			Category:         "Search",
			ShortDescription: `Manages a Search Service.`,
			Description:      ``,
			Icon:             "App_Services/10044-icon-service-Search-Services.svg",
			Keywords: []string{
				"search",
				"service",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Search Service.`,
				},
				resource.Attribute{
					Name:        "primary_key",
					Description: `The Primary Key used for Search Service Administration.`,
				},
				resource.Attribute{
					Name:        "secondary_key",
					Description: `The Secondary Key used for Search Service Administration.`,
				},
				resource.Attribute{
					Name:        "query_keys",
					Description: `A ` + "`" + `query_keys` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "public_network_access_enabled",
					Description: `Whether or not public network access is enabled for this resource.`,
				},
				resource.Attribute{
					Name:        "partition_count",
					Description: `The number of partitions which have been created.`,
				},
				resource.Attribute{
					Name:        "replica_count",
					Description: `The number of replica's which have been created.`,
				},
				resource.Attribute{
					Name:        "identity",
					Description: `An ` + "`" + `identity` + "`" + ` block as defined below. --- A ` + "`" + `identity` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The Type of Managed Identity which is used for the Search Service. --- A ` + "`" + `query_keys` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The value of this Query Key.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of this Query Key. --- A ` + "`" + `identity` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "principal_id",
					Description: `The (Client) ID of the Service Principal.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `The ID of the Tenant the Service Principal is assigned in. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Search Service.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_sentinel_alert_rule",
			Category:         "Sentinel",
			ShortDescription: `Gets information about an existing Sentinel Alert Rule.`,
			Description: `

Use this data source to access information about an existing Sentinel Alert Rule.

`,
			Keywords: []string{
				"sentinel",
				"alert",
				"rule",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Sentinel Alert Rule. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Sentinel Alert Rule.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_sentinel_alert_rule_template",
			Category:         "Sentinel",
			ShortDescription: `Gets information about an existing Sentinel Alert Rule Template.`,
			Description: `

Use this data source to access information about an existing Sentinel Alert Rule Template.

`,
			Keywords: []string{
				"sentinel",
				"alert",
				"rule",
				"template",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Sentinel.`,
				},
				resource.Attribute{
					Name:        "security_incident_template",
					Description: `A ` + "`" + `security_incident_template` + "`" + ` block as defined below. This only applies to Sentinel MS Security Incident Alert Rule Template.`,
				},
				resource.Attribute{
					Name:        "scheduled_template",
					Description: `A ` + "`" + `scheduled_template` + "`" + ` block as defined below. This only applies to Sentinel Scheduled Alert Rule Template. --- A ` + "`" + `security_incident_template` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of this Sentinel MS Security Incident Alert Rule Template.`,
				},
				resource.Attribute{
					Name:        "product_filter",
					Description: `The Microsoft Security Service from where the alert will be generated. --- A ` + "`" + `scheduled_template` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of this Sentinel Scheduled Alert Rule Template.`,
				},
				resource.Attribute{
					Name:        "query",
					Description: `The query of this Sentinel Scheduled Alert Rule Template.`,
				},
				resource.Attribute{
					Name:        "query_frequency",
					Description: `The ISO 8601 timespan duration between two consecutive queries.`,
				},
				resource.Attribute{
					Name:        "query_period",
					Description: `The ISO 8601 timespan duration, which determine the time period of the data covered by the query.`,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `The alert severity of this Sentinel Scheduled Alert Rule Template.`,
				},
				resource.Attribute{
					Name:        "tactics",
					Description: `A list of categories of attacks by which to classify the rule.`,
				},
				resource.Attribute{
					Name:        "trigger_operator",
					Description: `The alert trigger operator, combined with ` + "`" + `trigger_threshold` + "`" + `, setting alert threshold of this Sentinel Scheduled Alert Rule Template.`,
				},
				resource.Attribute{
					Name:        "trigger_threshold",
					Description: `The baseline number of query results generated, combined with ` + "`" + `trigger_operator` + "`" + `, setting alert threshold of this Sentinel Scheduled Alert Rule Template. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Sentinel.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_servicebus_namespace",
			Category:         "Messaging",
			ShortDescription: `Gets information about an existing ServiceBus Namespace.`,
			Description: `

Use this data source to access information about an existing ServiceBus Namespace.

`,
			Keywords: []string{
				"messaging",
				"servicebus",
				"namespace",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name of the ServiceBus Namespace.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `Specifies the name of the Resource Group where the ServiceBus Namespace exists. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location of the Resource Group in which the ServiceBus Namespace exists.`,
				},
				resource.Attribute{
					Name:        "sku",
					Description: `The Tier used for the ServiceBus Namespace.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `The capacity of the ServiceBus Namespace.`,
				},
				resource.Attribute{
					Name:        "zone_redundant",
					Description: `Whether or not this ServiceBus Namespace is zone redundant.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource. The following attributes are exported only if there is an authorization rule named ` + "`" + `RootManageSharedAccessKey` + "`" + ` which is created automatically by Azure.`,
				},
				resource.Attribute{
					Name:        "default_primary_connection_string",
					Description: `The primary connection string for the authorization rule ` + "`" + `RootManageSharedAccessKey` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "default_secondary_connection_string",
					Description: `The secondary connection string for the authorization rule ` + "`" + `RootManageSharedAccessKey` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "default_primary_key",
					Description: `The primary access key for the authorization rule ` + "`" + `RootManageSharedAccessKey` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "default_secondary_key",
					Description: `The secondary access key for the authorization rule ` + "`" + `RootManageSharedAccessKey` + "`" + `. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the ServiceBus Namespace.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "location",
					Description: `The location of the Resource Group in which the ServiceBus Namespace exists.`,
				},
				resource.Attribute{
					Name:        "sku",
					Description: `The Tier used for the ServiceBus Namespace.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `The capacity of the ServiceBus Namespace.`,
				},
				resource.Attribute{
					Name:        "zone_redundant",
					Description: `Whether or not this ServiceBus Namespace is zone redundant.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource. The following attributes are exported only if there is an authorization rule named ` + "`" + `RootManageSharedAccessKey` + "`" + ` which is created automatically by Azure.`,
				},
				resource.Attribute{
					Name:        "default_primary_connection_string",
					Description: `The primary connection string for the authorization rule ` + "`" + `RootManageSharedAccessKey` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "default_secondary_connection_string",
					Description: `The secondary connection string for the authorization rule ` + "`" + `RootManageSharedAccessKey` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "default_primary_key",
					Description: `The primary access key for the authorization rule ` + "`" + `RootManageSharedAccessKey` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "default_secondary_key",
					Description: `The secondary access key for the authorization rule ` + "`" + `RootManageSharedAccessKey` + "`" + `. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the ServiceBus Namespace.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_servicebus_namespace_authorization_rule",
			Category:         "Messaging",
			ShortDescription: `Gets information about an existing ServiceBus Namespace Authorization Rule.`,
			Description: `

Use this data source to access information about an existing ServiceBus Namespace Authorization Rule.

`,
			Keywords: []string{
				"messaging",
				"servicebus",
				"namespace",
				"authorization",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name of the ServiceBus Namespace Authorization Rule.`,
				},
				resource.Attribute{
					Name:        "namespace_name",
					Description: `Specifies the name of the ServiceBus Namespace.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `Specifies the name of the Resource Group where the ServiceBus Namespace exists. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the ServiceBus Namespace Authorization Rule.`,
				},
				resource.Attribute{
					Name:        "primary_connection_string",
					Description: `The primary connection string for the authorization rule.`,
				},
				resource.Attribute{
					Name:        "primary_key",
					Description: `The primary access key for the authorization rule.`,
				},
				resource.Attribute{
					Name:        "secondary_connection_string",
					Description: `The secondary connection string for the authorization rule.`,
				},
				resource.Attribute{
					Name:        "secondary_key",
					Description: `The secondary access key for the authorization rule. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the ServiceBus Namespace Authorization Rule.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the ServiceBus Namespace Authorization Rule.`,
				},
				resource.Attribute{
					Name:        "primary_connection_string",
					Description: `The primary connection string for the authorization rule.`,
				},
				resource.Attribute{
					Name:        "primary_key",
					Description: `The primary access key for the authorization rule.`,
				},
				resource.Attribute{
					Name:        "secondary_connection_string",
					Description: `The secondary connection string for the authorization rule.`,
				},
				resource.Attribute{
					Name:        "secondary_key",
					Description: `The secondary access key for the authorization rule. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the ServiceBus Namespace Authorization Rule.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_servicebus_queue",
			Category:         "Messaging",
			ShortDescription: `Gets information about an existing Service Bus Queue.`,
			Description: `

Use this data source to access information about an existing Service Bus Queue.

`,
			Keywords: []string{
				"messaging",
				"servicebus",
				"queue",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Service Bus Queue.`,
				},
				resource.Attribute{
					Name:        "auto_delete_on_idle",
					Description: `The ISO 8601 timespan duration of the idle interval after which the Queue is automatically deleted, minimum of 5 minutes.`,
				},
				resource.Attribute{
					Name:        "dead_lettering_on_message_expiration",
					Description: `Boolean flag which controls whether the Queue has dead letter support when a message expires.`,
				},
				resource.Attribute{
					Name:        "default_message_ttl",
					Description: `The ISO 8601 timespan duration of the TTL of messages sent to this queue. This is the default value used when TTL is not set on a message itself.`,
				},
				resource.Attribute{
					Name:        "duplicate_detection_history_time_window",
					Description: `The ISO 8601 timespan duration during which duplicates can be detected.`,
				},
				resource.Attribute{
					Name:        "enable_batched_operations",
					Description: `Boolean flag which controls whether server-side batched operations are enabled.`,
				},
				resource.Attribute{
					Name:        "enable_express",
					Description: `Boolean flag which controls whether Express Entities are enabled. An express queue holds a message in memory temporarily before writing it to persistent storage.`,
				},
				resource.Attribute{
					Name:        "enable_partitioning",
					Description: `Boolean flag which controls whether to enable the queue to be partitioned across multiple message brokers.`,
				},
				resource.Attribute{
					Name:        "forward_dead_lettered_messages_to",
					Description: `The name of a Queue or Topic to automatically forward dead lettered messages to.`,
				},
				resource.Attribute{
					Name:        "forward_to",
					Description: `The name of a Queue or Topic to automatically forward messages to. Please [see the documentation](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-auto-forwarding) for more information.`,
				},
				resource.Attribute{
					Name:        "lock_duration",
					Description: `The ISO 8601 timespan duration of a peek-lock; that is, the amount of time that the message is locked for other receivers.`,
				},
				resource.Attribute{
					Name:        "max_delivery_count",
					Description: `Integer value which controls when a message is automatically dead lettered.`,
				},
				resource.Attribute{
					Name:        "max_size_in_megabytes",
					Description: `Integer value which controls the size of memory allocated for the queue. For supported values see the "Queue or topic size" section of [Service Bus Quotas](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas).`,
				},
				resource.Attribute{
					Name:        "requires_duplicate_detection",
					Description: `Boolean flag which controls whether the Queue requires duplicate detection.`,
				},
				resource.Attribute{
					Name:        "requires_session",
					Description: `Boolean flag which controls whether the Queue requires sessions. This will allow ordered handling of unbounded sequences of related messages. With sessions enabled a queue can guarantee first-in-first-out delivery of messages.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the Queue. Possible values are ` + "`" + `Active` + "`" + `, ` + "`" + `Creating` + "`" + `, ` + "`" + `Deleting` + "`" + `, ` + "`" + `Disabled` + "`" + `, ` + "`" + `ReceiveDisabled` + "`" + `, ` + "`" + `Renaming` + "`" + `, ` + "`" + `SendDisabled` + "`" + `, ` + "`" + `Unknown` + "`" + `. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Service Bus Queue.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_servicebus_queue_authorization_rule",
			Category:         "Messaging",
			ShortDescription: `Gets information about an existing ServiceBus Queue Authorisation Rule within a ServiceBus Queue.`,
			Description: `

Use this data source to access information about an existing ServiceBus Queue Authorisation Rule within a ServiceBus Queue.

`,
			Keywords: []string{
				"messaging",
				"servicebus",
				"queue",
				"authorization",
				"rule",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the ServiceBus Queue Authorisation Rule.`,
				},
				resource.Attribute{
					Name:        "primary_key",
					Description: `The Primary Key for the ServiceBus Queue authorization Rule.`,
				},
				resource.Attribute{
					Name:        "primary_connection_string",
					Description: `The Primary Connection String for the ServiceBus Queue authorization Rule.`,
				},
				resource.Attribute{
					Name:        "secondary_key",
					Description: `The Secondary Key for the ServiceBus Queue authorization Rule.`,
				},
				resource.Attribute{
					Name:        "secondary_connection_string",
					Description: `The Secondary Connection String for the ServiceBus Queue authorization Rule. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the ServiceBus Queue Authorisation Rule.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_servicebus_subscription",
			Category:         "Messaging",
			ShortDescription: `Gets information about an existing ServiceBus Subscription.`,
			Description: `

Use this data source to access information about an existing ServiceBus Subscription.

`,
			Keywords: []string{
				"messaging",
				"servicebus",
				"subscription",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the ServiceBus Subscription.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) Specifies the name of the Resource Group where the ServiceBus Namespace exists.`,
				},
				resource.Attribute{
					Name:        "namespace_name",
					Description: `(Required) The name of the ServiceBus Namespace.`,
				},
				resource.Attribute{
					Name:        "topic_name",
					Description: `(Required) The name of the ServiceBus Topic. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "max_delivery_count",
					Description: `The maximum number of deliveries.`,
				},
				resource.Attribute{
					Name:        "auto_delete_on_idle",
					Description: `The idle interval after which the topic is automatically deleted.`,
				},
				resource.Attribute{
					Name:        "default_message_ttl",
					Description: `The Default message timespan to live. This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.`,
				},
				resource.Attribute{
					Name:        "lock_duration",
					Description: `The lock duration for the subscription.`,
				},
				resource.Attribute{
					Name:        "dead_lettering_on_message_expiration",
					Description: `Does the Service Bus Subscription have dead letter support when a message expires?`,
				},
				resource.Attribute{
					Name:        "dead_lettering_on_filter_evaluation_error",
					Description: `Does the ServiceBus Subscription have dead letter support on filter evaluation exceptions?`,
				},
				resource.Attribute{
					Name:        "enable_batched_operations",
					Description: `Are batched operations enabled on this ServiceBus Subscription?`,
				},
				resource.Attribute{
					Name:        "requires_session",
					Description: `Whether or not this ServiceBus Subscription supports session.`,
				},
				resource.Attribute{
					Name:        "forward_to",
					Description: `The name of a ServiceBus Queue or ServiceBus Topic where messages are automatically forwarded.`,
				},
				resource.Attribute{
					Name:        "forward_dead_lettered_messages_to",
					Description: `The name of a Queue or Topic to automatically forward Dead Letter messages to. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the ServiceBus Subscription.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "max_delivery_count",
					Description: `The maximum number of deliveries.`,
				},
				resource.Attribute{
					Name:        "auto_delete_on_idle",
					Description: `The idle interval after which the topic is automatically deleted.`,
				},
				resource.Attribute{
					Name:        "default_message_ttl",
					Description: `The Default message timespan to live. This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.`,
				},
				resource.Attribute{
					Name:        "lock_duration",
					Description: `The lock duration for the subscription.`,
				},
				resource.Attribute{
					Name:        "dead_lettering_on_message_expiration",
					Description: `Does the Service Bus Subscription have dead letter support when a message expires?`,
				},
				resource.Attribute{
					Name:        "dead_lettering_on_filter_evaluation_error",
					Description: `Does the ServiceBus Subscription have dead letter support on filter evaluation exceptions?`,
				},
				resource.Attribute{
					Name:        "enable_batched_operations",
					Description: `Are batched operations enabled on this ServiceBus Subscription?`,
				},
				resource.Attribute{
					Name:        "requires_session",
					Description: `Whether or not this ServiceBus Subscription supports session.`,
				},
				resource.Attribute{
					Name:        "forward_to",
					Description: `The name of a ServiceBus Queue or ServiceBus Topic where messages are automatically forwarded.`,
				},
				resource.Attribute{
					Name:        "forward_dead_lettered_messages_to",
					Description: `The name of a Queue or Topic to automatically forward Dead Letter messages to. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the ServiceBus Subscription.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_servicebus_topic",
			Category:         "Messaging",
			ShortDescription: `Gets information about an existing Service Bus Topic.`,
			Description: `

Use this data source to access information about an existing Service Bus Topic.

`,
			Keywords: []string{
				"messaging",
				"servicebus",
				"topic",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Service Bus Topic.`,
				},
				resource.Attribute{
					Name:        "auto_delete_on_idle",
					Description: `The ISO 8601 timespan duration of the idle interval after which the Topic is automatically deleted, minimum of 5 minutes.`,
				},
				resource.Attribute{
					Name:        "default_message_ttl",
					Description: `The ISO 8601 timespan duration of TTL of messages sent to this topic if no TTL value is set on the message itself.`,
				},
				resource.Attribute{
					Name:        "duplicate_detection_history_time_window",
					Description: `The ISO 8601 timespan duration during which duplicates can be detected.`,
				},
				resource.Attribute{
					Name:        "enable_batched_operations",
					Description: `Boolean flag which controls if server-side batched operations are enabled.`,
				},
				resource.Attribute{
					Name:        "enable_express",
					Description: `Boolean flag which controls whether Express Entities are enabled. An express topic holds a message in memory temporarily before writing it to persistent storage.`,
				},
				resource.Attribute{
					Name:        "enable_partitioning",
					Description: `Boolean flag which controls whether to enable the topic to be partitioned across multiple message brokers.`,
				},
				resource.Attribute{
					Name:        "max_size_in_megabytes",
					Description: `Integer value which controls the size of memory allocated for the topic. For supported values see the "Queue/topic size" section of [this document](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas).`,
				},
				resource.Attribute{
					Name:        "requires_duplicate_detection",
					Description: `Boolean flag which controls whether the Topic requires duplicate detection.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The Status of the Service Bus Topic. Acceptable values are Active or Disabled.`,
				},
				resource.Attribute{
					Name:        "support_ordering",
					Description: `Boolean flag which controls whether the Topic supports ordering. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Service Bus Topic.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_servicebus_topic_authorization_rule",
			Category:         "Messaging",
			ShortDescription: `Gets information about a ServiceBus Topic authorization Rule within a ServiceBus Topic.`,
			Description: `

Use this data source to access information about a ServiceBus Topic Authorization Rule within a ServiceBus Topic.

`,
			Keywords: []string{
				"messaging",
				"servicebus",
				"topic",
				"authorization",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the ServiceBus Topic Authorization Rule resource.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of the resource group in which the ServiceBus Namespace exists.`,
				},
				resource.Attribute{
					Name:        "namespace_name",
					Description: `The name of the ServiceBus Namespace.`,
				},
				resource.Attribute{
					Name:        "topic_name",
					Description: `The name of the ServiceBus Topic. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ServiceBus Topic ID.`,
				},
				resource.Attribute{
					Name:        "primary_key",
					Description: `The Primary Key for the ServiceBus Topic authorization Rule.`,
				},
				resource.Attribute{
					Name:        "primary_connection_string",
					Description: `The Primary Connection String for the ServiceBus Topic authorization Rule.`,
				},
				resource.Attribute{
					Name:        "secondary_key",
					Description: `The Secondary Key for the ServiceBus Topic authorization Rule.`,
				},
				resource.Attribute{
					Name:        "secondary_connection_string",
					Description: `The Secondary Connection String for the ServiceBus Topic authorization Rule. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the ServiceBus Topic Authorization Rule.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ServiceBus Topic ID.`,
				},
				resource.Attribute{
					Name:        "primary_key",
					Description: `The Primary Key for the ServiceBus Topic authorization Rule.`,
				},
				resource.Attribute{
					Name:        "primary_connection_string",
					Description: `The Primary Connection String for the ServiceBus Topic authorization Rule.`,
				},
				resource.Attribute{
					Name:        "secondary_key",
					Description: `The Secondary Key for the ServiceBus Topic authorization Rule.`,
				},
				resource.Attribute{
					Name:        "secondary_connection_string",
					Description: `The Secondary Connection String for the ServiceBus Topic authorization Rule. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the ServiceBus Topic Authorization Rule.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_shared_image",
			Category:         "Compute",
			ShortDescription: `Gets information about an existing Shared Image a Shared Image Gallery.`,
			Description: `

Use this data source to access information about an existing Shared Image within a Shared Image Gallery.

`,
			Keywords: []string{
				"compute",
				"shared",
				"image",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Shared Image.`,
				},
				resource.Attribute{
					Name:        "gallery_name",
					Description: `The name of the Shared Image Gallery in which the Shared Image exists.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of the Resource Group in which the Shared Image Gallery exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Shared Image.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of this Shared Image.`,
				},
				resource.Attribute{
					Name:        "eula",
					Description: `The End User Licence Agreement for the Shared Image.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The supported Azure location where the Shared Image Gallery exists.`,
				},
				resource.Attribute{
					Name:        "specialized",
					Description: `Specifies that the Operating System used inside this Image has not been Generalized (for example, ` + "`" + `sysprep` + "`" + ` on Windows has not been run).`,
				},
				resource.Attribute{
					Name:        "identifier",
					Description: `An ` + "`" + `identifier` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `The type of Operating System present in this Shared Image.`,
				},
				resource.Attribute{
					Name:        "hyper_v_generation",
					Description: `The generation of HyperV that the Virtual Machine used to create the Shared Image is based on.`,
				},
				resource.Attribute{
					Name:        "privacy_statement_uri",
					Description: `The URI containing the Privacy Statement for this Shared Image.`,
				},
				resource.Attribute{
					Name:        "release_note_uri",
					Description: `The URI containing the Release Notes for this Shared Image.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the Shared Image. --- A ` + "`" + `identifier` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "offer",
					Description: `The Offer Name for this Shared Image.`,
				},
				resource.Attribute{
					Name:        "publisher",
					Description: `The Publisher Name for this Gallery Image.`,
				},
				resource.Attribute{
					Name:        "sku",
					Description: `The Name of the SKU for this Gallery Image. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Shared Image a Shared Image Gallery.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Shared Image.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of this Shared Image.`,
				},
				resource.Attribute{
					Name:        "eula",
					Description: `The End User Licence Agreement for the Shared Image.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The supported Azure location where the Shared Image Gallery exists.`,
				},
				resource.Attribute{
					Name:        "specialized",
					Description: `Specifies that the Operating System used inside this Image has not been Generalized (for example, ` + "`" + `sysprep` + "`" + ` on Windows has not been run).`,
				},
				resource.Attribute{
					Name:        "identifier",
					Description: `An ` + "`" + `identifier` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `The type of Operating System present in this Shared Image.`,
				},
				resource.Attribute{
					Name:        "hyper_v_generation",
					Description: `The generation of HyperV that the Virtual Machine used to create the Shared Image is based on.`,
				},
				resource.Attribute{
					Name:        "privacy_statement_uri",
					Description: `The URI containing the Privacy Statement for this Shared Image.`,
				},
				resource.Attribute{
					Name:        "release_note_uri",
					Description: `The URI containing the Release Notes for this Shared Image.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the Shared Image. --- A ` + "`" + `identifier` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "offer",
					Description: `The Offer Name for this Shared Image.`,
				},
				resource.Attribute{
					Name:        "publisher",
					Description: `The Publisher Name for this Gallery Image.`,
				},
				resource.Attribute{
					Name:        "sku",
					Description: `The Name of the SKU for this Gallery Image. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Shared Image a Shared Image Gallery.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_shared_image_gallery",
			Category:         "Compute",
			ShortDescription: `Gets information about an existing Shared Image Gallery.`,
			Description: `

Use this data source to access information about an existing Shared Image Gallery.

`,
			Icon: "Compute/10039-icon-service-Shared-Image-Galleries.svg",
			Keywords: []string{
				"compute",
				"shared",
				"image",
				"gallery",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Shared Image Gallery.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of the Resource Group in which the Shared Image Gallery exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Shared Image Gallery.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description for the Shared Image Gallery.`,
				},
				resource.Attribute{
					Name:        "unique_name",
					Description: `The unique name assigned to the Shared Image Gallery.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags which are assigned to the Shared Image Gallery. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Shared Image Gallery.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Shared Image Gallery.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description for the Shared Image Gallery.`,
				},
				resource.Attribute{
					Name:        "unique_name",
					Description: `The unique name assigned to the Shared Image Gallery.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags which are assigned to the Shared Image Gallery. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Shared Image Gallery.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_shared_image_version",
			Category:         "Compute",
			ShortDescription: `Gets information about an existing Version of a Shared Image within a Shared Image Gallery.`,
			Description: `

Use this data source to access information about an existing Version of a Shared Image within a Shared Image Gallery.

`,
			Keywords: []string{
				"compute",
				"shared",
				"image",
				"version",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Image Version. ~>`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `The name of the Shared Image in which this Version exists.`,
				},
				resource.Attribute{
					Name:        "gallery_name",
					Description: `The name of the Shared Image in which the Shared Image exists.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of the Resource Group in which the Shared Image Gallery exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Shared Image.`,
				},
				resource.Attribute{
					Name:        "exclude_from_latest",
					Description: `Is this Image Version excluded from the ` + "`" + `latest` + "`" + ` filter?`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The supported Azure location where the Shared Image Gallery exists.`,
				},
				resource.Attribute{
					Name:        "managed_image_id",
					Description: `The ID of the Managed Image which was the source of this Shared Image Version.`,
				},
				resource.Attribute{
					Name:        "target_region",
					Description: `One or more ` + "`" + `target_region` + "`" + ` blocks as documented below.`,
				},
				resource.Attribute{
					Name:        "os_disk_snapshot_id",
					Description: `The ID of the OS disk snapshot which was the source of this Shared Image Version.`,
				},
				resource.Attribute{
					Name:        "os_disk_image_size_gb",
					Description: `The size of the OS disk snapshot (in Gigabytes) which was the source of this Shared Image Version.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the Shared Image. --- The ` + "`" + `target_region` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The Azure Region in which this Image Version exists.`,
				},
				resource.Attribute{
					Name:        "regional_replica_count",
					Description: `The number of replicas of the Image Version to be created per region.`,
				},
				resource.Attribute{
					Name:        "storage_account_type",
					Description: `The storage account type for the image version. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Version of a Shared Image within a Shared Image Gallery.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Shared Image.`,
				},
				resource.Attribute{
					Name:        "exclude_from_latest",
					Description: `Is this Image Version excluded from the ` + "`" + `latest` + "`" + ` filter?`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The supported Azure location where the Shared Image Gallery exists.`,
				},
				resource.Attribute{
					Name:        "managed_image_id",
					Description: `The ID of the Managed Image which was the source of this Shared Image Version.`,
				},
				resource.Attribute{
					Name:        "target_region",
					Description: `One or more ` + "`" + `target_region` + "`" + ` blocks as documented below.`,
				},
				resource.Attribute{
					Name:        "os_disk_snapshot_id",
					Description: `The ID of the OS disk snapshot which was the source of this Shared Image Version.`,
				},
				resource.Attribute{
					Name:        "os_disk_image_size_gb",
					Description: `The size of the OS disk snapshot (in Gigabytes) which was the source of this Shared Image Version.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the Shared Image. --- The ` + "`" + `target_region` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The Azure Region in which this Image Version exists.`,
				},
				resource.Attribute{
					Name:        "regional_replica_count",
					Description: `The number of replicas of the Image Version to be created per region.`,
				},
				resource.Attribute{
					Name:        "storage_account_type",
					Description: `The storage account type for the image version. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Version of a Shared Image within a Shared Image Gallery.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_shared_image_versions",
			Category:         "Compute",
			ShortDescription: `Gets information about existing Versions of a Shared Image within a Shared Image Gallery.`,
			Description: `

Use this data source to access information about existing Versions of a Shared Image within a Shared Image Gallery.

`,
			Icon: "Compute/10038-icon-service-Image-Versions.svg",
			Keywords: []string{
				"compute",
				"shared",
				"image",
				"versions",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "image_name",
					Description: `The name of the Shared Image in which this Version exists.`,
				},
				resource.Attribute{
					Name:        "gallery_name",
					Description: `The name of the Shared Image in which the Shared Image exists.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of the Resource Group in which the Shared Image Gallery exists.`,
				},
				resource.Attribute{
					Name:        "tags_filter",
					Description: `A mapping of tags to filter the list of images against. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "images",
					Description: `An ` + "`" + `images` + "`" + ` block as defined below: --- A ` + "`" + `images` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "exclude_from_latest",
					Description: `Is this Image Version excluded from the ` + "`" + `latest` + "`" + ` filter?`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The supported Azure location where the Shared Image Gallery exists.`,
				},
				resource.Attribute{
					Name:        "managed_image_id",
					Description: `The ID of the Managed Image which was the source of this Shared Image Version.`,
				},
				resource.Attribute{
					Name:        "target_region",
					Description: `One or more ` + "`" + `target_region` + "`" + ` blocks as documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the Shared Image. --- The ` + "`" + `target_region` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The Azure Region in which this Image Version exists.`,
				},
				resource.Attribute{
					Name:        "regional_replica_count",
					Description: `The number of replicas of the Image Version to be created per region.`,
				},
				resource.Attribute{
					Name:        "storage_account_type",
					Description: `The storage account type for the image version. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Versions of a Shared Image within a Shared Image Gallery.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "images",
					Description: `An ` + "`" + `images` + "`" + ` block as defined below: --- A ` + "`" + `images` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "exclude_from_latest",
					Description: `Is this Image Version excluded from the ` + "`" + `latest` + "`" + ` filter?`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The supported Azure location where the Shared Image Gallery exists.`,
				},
				resource.Attribute{
					Name:        "managed_image_id",
					Description: `The ID of the Managed Image which was the source of this Shared Image Version.`,
				},
				resource.Attribute{
					Name:        "target_region",
					Description: `One or more ` + "`" + `target_region` + "`" + ` blocks as documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the Shared Image. --- The ` + "`" + `target_region` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The Azure Region in which this Image Version exists.`,
				},
				resource.Attribute{
					Name:        "regional_replica_count",
					Description: `The number of replicas of the Image Version to be created per region.`,
				},
				resource.Attribute{
					Name:        "storage_account_type",
					Description: `The storage account type for the image version. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Versions of a Shared Image within a Shared Image Gallery.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_signalr_service",
			Category:         "Messaging",
			ShortDescription: `Gets information about an existing Azure SignalR service.`,
			Description: `

Use this data source to access information about an existing Azure SignalR service.

`,
			Keywords: []string{
				"messaging",
				"signalr",
				"service",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name of the SignalR service.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `Specifies the name of the resource group the SignalR service is located in. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the SignalR service.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `The FQDN of the SignalR service.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The publicly accessible IP of the SignalR service.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Specifies the supported Azure location where the SignalR service exists.`,
				},
				resource.Attribute{
					Name:        "public_port",
					Description: `The publicly accessible port of the SignalR service which is designed for browser/client use.`,
				},
				resource.Attribute{
					Name:        "server_port",
					Description: `The publicly accessible port of the SignalR service which is designed for customer server side use.`,
				},
				resource.Attribute{
					Name:        "primary_access_key",
					Description: `The primary access key of the SignalR service.`,
				},
				resource.Attribute{
					Name:        "primary_connection_string",
					Description: `The primary connection string of the SignalR service.`,
				},
				resource.Attribute{
					Name:        "secondary_access_key",
					Description: `The secondary access key of the SignalR service.`,
				},
				resource.Attribute{
					Name:        "secondary_connection_string",
					Description: `The secondary connection string of the SignalR service. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the SignalR service.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the SignalR service.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `The FQDN of the SignalR service.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The publicly accessible IP of the SignalR service.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Specifies the supported Azure location where the SignalR service exists.`,
				},
				resource.Attribute{
					Name:        "public_port",
					Description: `The publicly accessible port of the SignalR service which is designed for browser/client use.`,
				},
				resource.Attribute{
					Name:        "server_port",
					Description: `The publicly accessible port of the SignalR service which is designed for customer server side use.`,
				},
				resource.Attribute{
					Name:        "primary_access_key",
					Description: `The primary access key of the SignalR service.`,
				},
				resource.Attribute{
					Name:        "primary_connection_string",
					Description: `The primary connection string of the SignalR service.`,
				},
				resource.Attribute{
					Name:        "secondary_access_key",
					Description: `The secondary access key of the SignalR service.`,
				},
				resource.Attribute{
					Name:        "secondary_connection_string",
					Description: `The secondary connection string of the SignalR service. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the SignalR service.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_snapshot",
			Category:         "Compute",
			ShortDescription: `Get information about an existing Snapshot`,
			Description: `

Use this data source to access information about an existing Snapshot.

`,
			Keywords: []string{
				"compute",
				"snapshot",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name of the Snapshot.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `Specifies the name of the resource group the Snapshot is located in. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Snapshot.`,
				},
				resource.Attribute{
					Name:        "create_option",
					Description: `How the snapshot was created.`,
				},
				resource.Attribute{
					Name:        "source_uri",
					Description: `The URI to a Managed or Unmanaged Disk.`,
				},
				resource.Attribute{
					Name:        "source_resource_id",
					Description: `The reference to an existing snapshot.`,
				},
				resource.Attribute{
					Name:        "storage_account_id",
					Description: `The ID of an storage account.`,
				},
				resource.Attribute{
					Name:        "disk_size_gb",
					Description: `The size of the Snapshotted Disk in GB. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Snapshot.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Snapshot.`,
				},
				resource.Attribute{
					Name:        "create_option",
					Description: `How the snapshot was created.`,
				},
				resource.Attribute{
					Name:        "source_uri",
					Description: `The URI to a Managed or Unmanaged Disk.`,
				},
				resource.Attribute{
					Name:        "source_resource_id",
					Description: `The reference to an existing snapshot.`,
				},
				resource.Attribute{
					Name:        "storage_account_id",
					Description: `The ID of an storage account.`,
				},
				resource.Attribute{
					Name:        "disk_size_gb",
					Description: `The size of the Snapshotted Disk in GB. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Snapshot.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_spring_cloud_app",
			Category:         "Spring Cloud",
			ShortDescription: `Gets information about an existing Spring Cloud Application`,
			Description: `

Use this data source to access information about an existing Spring Cloud Application.

`,
			Keywords: []string{
				"spring",
				"cloud",
				"app",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Spring Cloud Application.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The name of the Resource Group where the Spring Cloud Application exists.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The name of the Spring Cloud Service. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Spring Cloud Application.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `The Fully Qualified DNS Name.`,
				},
				resource.Attribute{
					Name:        "https_only",
					Description: `Is only https allowed?`,
				},
				resource.Attribute{
					Name:        "identity",
					Description: `An ` + "`" + `identity` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `Does the Spring Cloud Application have public endpoint?`,
				},
				resource.Attribute{
					Name:        "persistent_disk",
					Description: `A ` + "`" + `persistent_disk` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `The public endpoint of the Spring Cloud Application.`,
				},
				resource.Attribute{
					Name:        "tls_enabled",
					Description: `Is End to End TLS Enabled? --- The ` + "`" + `identity` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "principal_id",
					Description: `The Principal ID for the Service Principal associated with the Managed Service Identity of this Spring Cloud Application.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `The Tenant ID for the Service Principal associated with the Managed Service Identity of this Spring Cloud Application.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The Type of Managed Identity assigned to the Spring Cloud Application. --- The ` + "`" + `persistent_disk` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "mount_path",
					Description: `The mount path of the persistent disk.`,
				},
				resource.Attribute{
					Name:        "size_in_gb",
					Description: `The size of the persistent disk in GB. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Spring Cloud Application.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Spring Cloud Application.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `The Fully Qualified DNS Name.`,
				},
				resource.Attribute{
					Name:        "https_only",
					Description: `Is only https allowed?`,
				},
				resource.Attribute{
					Name:        "identity",
					Description: `An ` + "`" + `identity` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `Does the Spring Cloud Application have public endpoint?`,
				},
				resource.Attribute{
					Name:        "persistent_disk",
					Description: `A ` + "`" + `persistent_disk` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `The public endpoint of the Spring Cloud Application.`,
				},
				resource.Attribute{
					Name:        "tls_enabled",
					Description: `Is End to End TLS Enabled? --- The ` + "`" + `identity` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "principal_id",
					Description: `The Principal ID for the Service Principal associated with the Managed Service Identity of this Spring Cloud Application.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `The Tenant ID for the Service Principal associated with the Managed Service Identity of this Spring Cloud Application.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The Type of Managed Identity assigned to the Spring Cloud Application. --- The ` + "`" + `persistent_disk` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "mount_path",
					Description: `The mount path of the persistent disk.`,
				},
				resource.Attribute{
					Name:        "size_in_gb",
					Description: `The size of the persistent disk in GB. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Spring Cloud Application.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_spring_cloud_service",
			Category:         "Spring Cloud",
			ShortDescription: `Gets information about an existing Spring Cloud Service`,
			Description: `

Use this data source to access information about an existing Spring Cloud Service.

`,
			Keywords: []string{
				"spring",
				"cloud",
				"service",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies The name of the Spring Cloud Service resource.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) Specifies the name of the Resource Group where the Spring Cloud Service exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Spring Cloud Service.`,
				},
				resource.Attribute{
					Name:        "config_server_git_setting",
					Description: `A ` + "`" + `config_server_git_setting` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location of Spring Cloud Service.`,
				},
				resource.Attribute{
					Name:        "outbound_public_ip_addresses",
					Description: `A list of the outbound Public IP Addresses used by this Spring Cloud Service.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to Spring Cloud Service. --- The ` + "`" + `config_server_git_setting` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `The URI of the Git repository`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The default label of the Git repository, which is a branch name, tag name, or commit-id of the repository`,
				},
				resource.Attribute{
					Name:        "search_paths",
					Description: `An array of strings used to search subdirectories of the Git repository.`,
				},
				resource.Attribute{
					Name:        "http_basic_auth",
					Description: `A ` + "`" + `http_basic_auth` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "ssh_auth",
					Description: `A ` + "`" + `ssh_auth` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "repository",
					Description: `One or more ` + "`" + `repository` + "`" + ` blocks as defined below. --- The ` + "`" + `repository` + "`" + ` block contains the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name to identify on the Git repository.`,
				},
				resource.Attribute{
					Name:        "pattern",
					Description: `An array of strings used to match an application name. For each pattern, use the ` + "`" + `{application}/{profile}` + "`" + ` format with wildcards.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `The URI of the Git repository`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The default label of the Git repository, which is a branch name, tag name, or commit-id of the repository`,
				},
				resource.Attribute{
					Name:        "search_paths",
					Description: `An array of strings used to search subdirectories of the Git repository.`,
				},
				resource.Attribute{
					Name:        "http_basic_auth",
					Description: `A ` + "`" + `http_basic_auth` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "ssh_auth",
					Description: `A ` + "`" + `ssh_auth` + "`" + ` block as defined below. --- The ` + "`" + `http_basic_auth` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The username used to access the Http Basic Authentication Git repository server.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `The password used to access the Http Basic Authentication Git repository server. --- The ` + "`" + `ssh_auth` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `The SSH private key to access the Git repository, needed when the URI starts with ` + "`" + `git@` + "`" + ` or ` + "`" + `ssh://` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "host_key",
					Description: `The host key of the Git repository server.`,
				},
				resource.Attribute{
					Name:        "host_key_algorithm",
					Description: `The host key algorithm.`,
				},
				resource.Attribute{
					Name:        "strict_host_key_checking_enabled",
					Description: `Indicates whether the Config Server instance will fail to start if the host_key does not match. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Spring Cloud Service.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Spring Cloud Service.`,
				},
				resource.Attribute{
					Name:        "config_server_git_setting",
					Description: `A ` + "`" + `config_server_git_setting` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location of Spring Cloud Service.`,
				},
				resource.Attribute{
					Name:        "outbound_public_ip_addresses",
					Description: `A list of the outbound Public IP Addresses used by this Spring Cloud Service.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to Spring Cloud Service. --- The ` + "`" + `config_server_git_setting` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `The URI of the Git repository`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The default label of the Git repository, which is a branch name, tag name, or commit-id of the repository`,
				},
				resource.Attribute{
					Name:        "search_paths",
					Description: `An array of strings used to search subdirectories of the Git repository.`,
				},
				resource.Attribute{
					Name:        "http_basic_auth",
					Description: `A ` + "`" + `http_basic_auth` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "ssh_auth",
					Description: `A ` + "`" + `ssh_auth` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "repository",
					Description: `One or more ` + "`" + `repository` + "`" + ` blocks as defined below. --- The ` + "`" + `repository` + "`" + ` block contains the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name to identify on the Git repository.`,
				},
				resource.Attribute{
					Name:        "pattern",
					Description: `An array of strings used to match an application name. For each pattern, use the ` + "`" + `{application}/{profile}` + "`" + ` format with wildcards.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `The URI of the Git repository`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The default label of the Git repository, which is a branch name, tag name, or commit-id of the repository`,
				},
				resource.Attribute{
					Name:        "search_paths",
					Description: `An array of strings used to search subdirectories of the Git repository.`,
				},
				resource.Attribute{
					Name:        "http_basic_auth",
					Description: `A ` + "`" + `http_basic_auth` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "ssh_auth",
					Description: `A ` + "`" + `ssh_auth` + "`" + ` block as defined below. --- The ` + "`" + `http_basic_auth` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The username used to access the Http Basic Authentication Git repository server.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `The password used to access the Http Basic Authentication Git repository server. --- The ` + "`" + `ssh_auth` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `The SSH private key to access the Git repository, needed when the URI starts with ` + "`" + `git@` + "`" + ` or ` + "`" + `ssh://` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "host_key",
					Description: `The host key of the Git repository server.`,
				},
				resource.Attribute{
					Name:        "host_key_algorithm",
					Description: `The host key algorithm.`,
				},
				resource.Attribute{
					Name:        "strict_host_key_checking_enabled",
					Description: `Indicates whether the Config Server instance will fail to start if the host_key does not match. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Spring Cloud Service.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_sql_database",
			Category:         "Database",
			ShortDescription: `Gets information about an existing SQL Azure Database.`,
			Description: `

Use this data source to access information about an existing SQL Azure Database.

`,
			Icon: "Databases/10130-icon-service-SQL-Database.svg",
			Keywords: []string{
				"database",
				"sql",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the SQL Database.`,
				},
				resource.Attribute{
					Name:        "server_name",
					Description: `The name of the SQL Server.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `Specifies the name of the Resource Group where the Azure SQL Database exists. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The SQL Database ID.`,
				},
				resource.Attribute{
					Name:        "collation",
					Description: `The name of the collation.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `The creation date of the SQL Database.`,
				},
				resource.Attribute{
					Name:        "default_secondary_location",
					Description: `The default secondary location of the SQL Database.`,
				},
				resource.Attribute{
					Name:        "edition",
					Description: `The edition of the database.`,
				},
				resource.Attribute{
					Name:        "elastic_pool_name",
					Description: `The name of the elastic database pool the database belongs to.`,
				},
				resource.Attribute{
					Name:        "failover_group_id",
					Description: `The ID of the failover group the database belongs to.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location of the Resource Group in which the SQL Server exists.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the database.`,
				},
				resource.Attribute{
					Name:        "read_scale",
					Description: `Indicate if read-only connections will be redirected to a high-available replica.`,
				},
				resource.Attribute{
					Name:        "requested_service_objective_id",
					Description: `The ID pertaining to the performance level of the database.`,
				},
				resource.Attribute{
					Name:        "requested_service_objective_name",
					Description: `The name pertaining to the performance level of the database.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of the resource group in which the database resides. This will always be the same resource group as the Database Server.`,
				},
				resource.Attribute{
					Name:        "server_name",
					Description: `The name of the SQL Server on which to create the database.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the SQL Azure Database.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The SQL Database ID.`,
				},
				resource.Attribute{
					Name:        "collation",
					Description: `The name of the collation.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `The creation date of the SQL Database.`,
				},
				resource.Attribute{
					Name:        "default_secondary_location",
					Description: `The default secondary location of the SQL Database.`,
				},
				resource.Attribute{
					Name:        "edition",
					Description: `The edition of the database.`,
				},
				resource.Attribute{
					Name:        "elastic_pool_name",
					Description: `The name of the elastic database pool the database belongs to.`,
				},
				resource.Attribute{
					Name:        "failover_group_id",
					Description: `The ID of the failover group the database belongs to.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location of the Resource Group in which the SQL Server exists.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the database.`,
				},
				resource.Attribute{
					Name:        "read_scale",
					Description: `Indicate if read-only connections will be redirected to a high-available replica.`,
				},
				resource.Attribute{
					Name:        "requested_service_objective_id",
					Description: `The ID pertaining to the performance level of the database.`,
				},
				resource.Attribute{
					Name:        "requested_service_objective_name",
					Description: `The name pertaining to the performance level of the database.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of the resource group in which the database resides. This will always be the same resource group as the Database Server.`,
				},
				resource.Attribute{
					Name:        "server_name",
					Description: `The name of the SQL Server on which to create the database.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the SQL Azure Database.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_sql_server",
			Category:         "Database",
			ShortDescription: `Gets information about an existing SQL Azure Database Server.`,
			Description: `

Use this data source to access information about an existing SQL Azure Database Server.

`,
			Icon: "Databases/10132-icon-service-SQL-Server.svg",
			Keywords: []string{
				"database",
				"sql",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the SQL Server.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `Specifies the name of the Resource Group where the SQL Server exists. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the SQL Server resource.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location of the Resource Group in which the SQL Server exists.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `The fully qualified domain name of the SQL Server.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version of the SQL Server.`,
				},
				resource.Attribute{
					Name:        "administrator_login",
					Description: `The administrator username of the SQL Server.`,
				},
				resource.Attribute{
					Name:        "identity",
					Description: `An ` + "`" + `identity` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource. --- An ` + "`" + `identity` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "principal_id",
					Description: `The ID of the Principal (Client) in Azure Active Directory.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `The ID of the Azure Active Directory Tenant.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The identity type of the SQL Server. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the SQL Azure Database Server.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the SQL Server resource.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location of the Resource Group in which the SQL Server exists.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `The fully qualified domain name of the SQL Server.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version of the SQL Server.`,
				},
				resource.Attribute{
					Name:        "administrator_login",
					Description: `The administrator username of the SQL Server.`,
				},
				resource.Attribute{
					Name:        "identity",
					Description: `An ` + "`" + `identity` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource. --- An ` + "`" + `identity` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "principal_id",
					Description: `The ID of the Principal (Client) in Azure Active Directory.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `The ID of the Azure Active Directory Tenant.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The identity type of the SQL Server. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the SQL Azure Database Server.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_ssh_public_key",
			Category:         "Compute",
			ShortDescription: `Gets information about an existing SSH Public Key.`,
			Description: `

Use this data source to access information about an existing SSH Public Key.

`,
			Keywords: []string{
				"compute",
				"ssh",
				"public",
				"key",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the SSH Public Key.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `The SSH public key used to authenticate to a virtual machine through ssh. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the SSH Public Key.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_storage_account",
			Category:         "Storage",
			ShortDescription: `Gets information about an existing Storage Account.`,
			Description: `

Use this data source to access information about an existing Storage Account.

`,
			Icon: "Storage/10087-icon-service-Storage-Accounts-(Classic).svg",
			Keywords: []string{
				"storage",
				"account",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name of the Storage Account`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `Specifies the name of the resource group the Storage Account is located in. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Storage Account.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the Storage Account exists`,
				},
				resource.Attribute{
					Name:        "account_kind",
					Description: `The Kind of account.`,
				},
				resource.Attribute{
					Name:        "account_tier",
					Description: `The Tier of this storage account.`,
				},
				resource.Attribute{
					Name:        "account_replication_type",
					Description: `The type of replication used for this storage account.`,
				},
				resource.Attribute{
					Name:        "access_tier",
					Description: `The access tier for ` + "`" + `BlobStorage` + "`" + ` accounts.`,
				},
				resource.Attribute{
					Name:        "enable_https_traffic_only",
					Description: `Is traffic only allowed via HTTPS? See [here](https://docs.microsoft.com/en-us/azure/storage/storage-require-secure-transfer/) for more information.`,
				},
				resource.Attribute{
					Name:        "min_tls_version",
					Description: `The minimum supported TLS version for this storage account.`,
				},
				resource.Attribute{
					Name:        "allow_blob_public_access",
					Description: `Is public access allowed to all blobs or containers in the storage account?`,
				},
				resource.Attribute{
					Name:        "is_hns_enabled",
					Description: `Is Hierarchical Namespace enabled?`,
				},
				resource.Attribute{
					Name:        "custom_domain",
					Description: `A ` + "`" + `custom_domain` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assigned to the resource.`,
				},
				resource.Attribute{
					Name:        "primary_location",
					Description: `The primary location of the Storage Account.`,
				},
				resource.Attribute{
					Name:        "secondary_location",
					Description: `The secondary location of the Storage Account.`,
				},
				resource.Attribute{
					Name:        "primary_blob_endpoint",
					Description: `The endpoint URL for blob storage in the primary location.`,
				},
				resource.Attribute{
					Name:        "primary_blob_host",
					Description: `The hostname with port if applicable for blob storage in the primary location.`,
				},
				resource.Attribute{
					Name:        "secondary_blob_endpoint",
					Description: `The endpoint URL for blob storage in the secondary location.`,
				},
				resource.Attribute{
					Name:        "secondary_blob_host",
					Description: `The hostname with port if applicable for blob storage in the secondary location.`,
				},
				resource.Attribute{
					Name:        "primary_queue_endpoint",
					Description: `The endpoint URL for queue storage in the primary location.`,
				},
				resource.Attribute{
					Name:        "primary_queue_host",
					Description: `The hostname with port if applicable for queue storage in the primary location.`,
				},
				resource.Attribute{
					Name:        "secondary_queue_endpoint",
					Description: `The endpoint URL for queue storage in the secondary location.`,
				},
				resource.Attribute{
					Name:        "secondary_queue_host",
					Description: `The hostname with port if applicable for queue storage in the secondary location.`,
				},
				resource.Attribute{
					Name:        "primary_table_endpoint",
					Description: `The endpoint URL for table storage in the primary location.`,
				},
				resource.Attribute{
					Name:        "primary_table_host",
					Description: `The hostname with port if applicable for table storage in the primary location.`,
				},
				resource.Attribute{
					Name:        "secondary_table_endpoint",
					Description: `The endpoint URL for table storage in the secondary location.`,
				},
				resource.Attribute{
					Name:        "secondary_table_host",
					Description: `The hostname with port if applicable for table storage in the secondary location.`,
				},
				resource.Attribute{
					Name:        "primary_file_endpoint",
					Description: `The endpoint URL for file storage in the primary location.`,
				},
				resource.Attribute{
					Name:        "primary_file_host",
					Description: `The hostname with port if applicable for file storage in the primary location.`,
				},
				resource.Attribute{
					Name:        "secondary_file_endpoint",
					Description: `The endpoint URL for file storage in the secondary location.`,
				},
				resource.Attribute{
					Name:        "secondary_file_host",
					Description: `The hostname with port if applicable for file storage in the secondary location.`,
				},
				resource.Attribute{
					Name:        "primary_dfs_endpoint",
					Description: `The endpoint URL for DFS storage in the primary location.`,
				},
				resource.Attribute{
					Name:        "primary_dfs_host",
					Description: `The hostname with port if applicable for DFS storage in the primary location.`,
				},
				resource.Attribute{
					Name:        "secondary_dfs_endpoint",
					Description: `The endpoint URL for DFS storage in the secondary location.`,
				},
				resource.Attribute{
					Name:        "secondary_dfs_host",
					Description: `The hostname with port if applicable for DFS storage in the secondary location.`,
				},
				resource.Attribute{
					Name:        "primary_web_endpoint",
					Description: `The endpoint URL for web storage in the primary location.`,
				},
				resource.Attribute{
					Name:        "primary_web_host",
					Description: `The hostname with port if applicable for web storage in the primary location.`,
				},
				resource.Attribute{
					Name:        "secondary_web_endpoint",
					Description: `The endpoint URL for web storage in the secondary location.`,
				},
				resource.Attribute{
					Name:        "secondary_web_host",
					Description: `The hostname with port if applicable for web storage in the secondary location.`,
				},
				resource.Attribute{
					Name:        "primary_access_key",
					Description: `The primary access key for the Storage Account.`,
				},
				resource.Attribute{
					Name:        "secondary_access_key",
					Description: `The secondary access key for the Storage Account.`,
				},
				resource.Attribute{
					Name:        "primary_connection_string",
					Description: `The connection string associated with the primary location`,
				},
				resource.Attribute{
					Name:        "secondary_connection_string",
					Description: `The connection string associated with the secondary location`,
				},
				resource.Attribute{
					Name:        "primary_blob_connection_string",
					Description: `The connection string associated with the primary blob location`,
				},
				resource.Attribute{
					Name:        "secondary_blob_connection_string",
					Description: `The connection string associated with the secondary blob location ~>`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The Custom Domain Name used for the Storage Account. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Storage Account.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Storage Account.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the Storage Account exists`,
				},
				resource.Attribute{
					Name:        "account_kind",
					Description: `The Kind of account.`,
				},
				resource.Attribute{
					Name:        "account_tier",
					Description: `The Tier of this storage account.`,
				},
				resource.Attribute{
					Name:        "account_replication_type",
					Description: `The type of replication used for this storage account.`,
				},
				resource.Attribute{
					Name:        "access_tier",
					Description: `The access tier for ` + "`" + `BlobStorage` + "`" + ` accounts.`,
				},
				resource.Attribute{
					Name:        "enable_https_traffic_only",
					Description: `Is traffic only allowed via HTTPS? See [here](https://docs.microsoft.com/en-us/azure/storage/storage-require-secure-transfer/) for more information.`,
				},
				resource.Attribute{
					Name:        "min_tls_version",
					Description: `The minimum supported TLS version for this storage account.`,
				},
				resource.Attribute{
					Name:        "allow_blob_public_access",
					Description: `Is public access allowed to all blobs or containers in the storage account?`,
				},
				resource.Attribute{
					Name:        "is_hns_enabled",
					Description: `Is Hierarchical Namespace enabled?`,
				},
				resource.Attribute{
					Name:        "custom_domain",
					Description: `A ` + "`" + `custom_domain` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assigned to the resource.`,
				},
				resource.Attribute{
					Name:        "primary_location",
					Description: `The primary location of the Storage Account.`,
				},
				resource.Attribute{
					Name:        "secondary_location",
					Description: `The secondary location of the Storage Account.`,
				},
				resource.Attribute{
					Name:        "primary_blob_endpoint",
					Description: `The endpoint URL for blob storage in the primary location.`,
				},
				resource.Attribute{
					Name:        "primary_blob_host",
					Description: `The hostname with port if applicable for blob storage in the primary location.`,
				},
				resource.Attribute{
					Name:        "secondary_blob_endpoint",
					Description: `The endpoint URL for blob storage in the secondary location.`,
				},
				resource.Attribute{
					Name:        "secondary_blob_host",
					Description: `The hostname with port if applicable for blob storage in the secondary location.`,
				},
				resource.Attribute{
					Name:        "primary_queue_endpoint",
					Description: `The endpoint URL for queue storage in the primary location.`,
				},
				resource.Attribute{
					Name:        "primary_queue_host",
					Description: `The hostname with port if applicable for queue storage in the primary location.`,
				},
				resource.Attribute{
					Name:        "secondary_queue_endpoint",
					Description: `The endpoint URL for queue storage in the secondary location.`,
				},
				resource.Attribute{
					Name:        "secondary_queue_host",
					Description: `The hostname with port if applicable for queue storage in the secondary location.`,
				},
				resource.Attribute{
					Name:        "primary_table_endpoint",
					Description: `The endpoint URL for table storage in the primary location.`,
				},
				resource.Attribute{
					Name:        "primary_table_host",
					Description: `The hostname with port if applicable for table storage in the primary location.`,
				},
				resource.Attribute{
					Name:        "secondary_table_endpoint",
					Description: `The endpoint URL for table storage in the secondary location.`,
				},
				resource.Attribute{
					Name:        "secondary_table_host",
					Description: `The hostname with port if applicable for table storage in the secondary location.`,
				},
				resource.Attribute{
					Name:        "primary_file_endpoint",
					Description: `The endpoint URL for file storage in the primary location.`,
				},
				resource.Attribute{
					Name:        "primary_file_host",
					Description: `The hostname with port if applicable for file storage in the primary location.`,
				},
				resource.Attribute{
					Name:        "secondary_file_endpoint",
					Description: `The endpoint URL for file storage in the secondary location.`,
				},
				resource.Attribute{
					Name:        "secondary_file_host",
					Description: `The hostname with port if applicable for file storage in the secondary location.`,
				},
				resource.Attribute{
					Name:        "primary_dfs_endpoint",
					Description: `The endpoint URL for DFS storage in the primary location.`,
				},
				resource.Attribute{
					Name:        "primary_dfs_host",
					Description: `The hostname with port if applicable for DFS storage in the primary location.`,
				},
				resource.Attribute{
					Name:        "secondary_dfs_endpoint",
					Description: `The endpoint URL for DFS storage in the secondary location.`,
				},
				resource.Attribute{
					Name:        "secondary_dfs_host",
					Description: `The hostname with port if applicable for DFS storage in the secondary location.`,
				},
				resource.Attribute{
					Name:        "primary_web_endpoint",
					Description: `The endpoint URL for web storage in the primary location.`,
				},
				resource.Attribute{
					Name:        "primary_web_host",
					Description: `The hostname with port if applicable for web storage in the primary location.`,
				},
				resource.Attribute{
					Name:        "secondary_web_endpoint",
					Description: `The endpoint URL for web storage in the secondary location.`,
				},
				resource.Attribute{
					Name:        "secondary_web_host",
					Description: `The hostname with port if applicable for web storage in the secondary location.`,
				},
				resource.Attribute{
					Name:        "primary_access_key",
					Description: `The primary access key for the Storage Account.`,
				},
				resource.Attribute{
					Name:        "secondary_access_key",
					Description: `The secondary access key for the Storage Account.`,
				},
				resource.Attribute{
					Name:        "primary_connection_string",
					Description: `The connection string associated with the primary location`,
				},
				resource.Attribute{
					Name:        "secondary_connection_string",
					Description: `The connection string associated with the secondary location`,
				},
				resource.Attribute{
					Name:        "primary_blob_connection_string",
					Description: `The connection string associated with the primary blob location`,
				},
				resource.Attribute{
					Name:        "secondary_blob_connection_string",
					Description: `The connection string associated with the secondary blob location ~>`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The Custom Domain Name used for the Storage Account. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Storage Account.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_storage_account_blob_container_sas",
			Category:         "Storage",
			ShortDescription: `Gets a Shared Access Signature (SAS Token) for an existing Storage Account Blob Container.`,
			Description: `

Use this data source to obtain a Shared Access Signature (SAS Token) for an existing Storage Account Blob Container.

Shared access signatures allow fine-grained, ephemeral access control to various aspects of an Azure Storage Account Blob Container.

`,
			Keywords: []string{
				"storage",
				"account",
				"blob",
				"container",
				"sas",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "connection_string",
					Description: `The connection string for the storage account to which this SAS applies. Typically directly from the ` + "`" + `primary_connection_string` + "`" + ` attribute of a terraform created ` + "`" + `azurerm_storage_account` + "`" + ` resource.`,
				},
				resource.Attribute{
					Name:        "container_name",
					Description: `Name of the container.`,
				},
				resource.Attribute{
					Name:        "https_only",
					Description: `(Optional) Only permit ` + "`" + `https` + "`" + ` access. If ` + "`" + `false` + "`" + `, both ` + "`" + `http` + "`" + ` and ` + "`" + `https` + "`" + ` are permitted. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) Single ipv4 address or range (connected with a dash) of ipv4 addresses.`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `The starting time and date of validity of this SAS. Must be a valid ISO-8601 format time/date string.`,
				},
				resource.Attribute{
					Name:        "expiry",
					Description: `The expiration time and date of this SAS. Must be a valid ISO-8601 format time/date string.`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `A ` + "`" + `permissions` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "cache_control",
					Description: `(Optional) The ` + "`" + `Cache-Control` + "`" + ` response header that is sent when this SAS token is used.`,
				},
				resource.Attribute{
					Name:        "content_disposition",
					Description: `(Optional) The ` + "`" + `Content-Disposition` + "`" + ` response header that is sent when this SAS token is used.`,
				},
				resource.Attribute{
					Name:        "content_encoding",
					Description: `(Optional) The ` + "`" + `Content-Encoding` + "`" + ` response header that is sent when this SAS token is used.`,
				},
				resource.Attribute{
					Name:        "content_language",
					Description: `(Optional) The ` + "`" + `Content-Language` + "`" + ` response header that is sent when this SAS token is used.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `(Optional) The ` + "`" + `Content-Type` + "`" + ` response header that is sent when this SAS token is used. --- A ` + "`" + `permissions` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `Should Read permissions be enabled for this SAS?`,
				},
				resource.Attribute{
					Name:        "add",
					Description: `Should Add permissions be enabled for this SAS?`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `Should Create permissions be enabled for this SAS?`,
				},
				resource.Attribute{
					Name:        "write",
					Description: `Should Write permissions be enabled for this SAS?`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `Should Delete permissions be enabled for this SAS?`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `Should List permissions be enabled for this SAS? Refer to the [SAS creation reference from Azure](https://docs.microsoft.com/en-us/rest/api/storageservices/create-service-sas) for additional details on the fields above. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "sas",
					Description: `The computed Blob Container Shared Access Signature (SAS). ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Blob Container.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "sas",
					Description: `The computed Blob Container Shared Access Signature (SAS). ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Blob Container.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_storage_account_sas",
			Category:         "Storage",
			ShortDescription: `Gets a Shared Access Signature (SAS Token) for an existing Storage Account.`,
			Description: `

Use this data source to obtain a Shared Access Signature (SAS Token) for an existing Storage Account.

Shared access signatures allow fine-grained, ephemeral access control to various aspects of an Azure Storage Account.

Note that this is an [Account SAS](https://docs.microsoft.com/en-us/rest/api/storageservices/constructing-an-account-sas)
and *not* a [Service SAS](https://docs.microsoft.com/en-us/rest/api/storageservices/constructing-a-service-sas).

`,
			Keywords: []string{
				"storage",
				"account",
				"sas",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "connection_string",
					Description: `The connection string for the storage account to which this SAS applies. Typically directly from the ` + "`" + `primary_connection_string` + "`" + ` attribute of a terraform created ` + "`" + `azurerm_storage_account` + "`" + ` resource.`,
				},
				resource.Attribute{
					Name:        "https_only",
					Description: `(Optional) Only permit ` + "`" + `https` + "`" + ` access. If ` + "`" + `false` + "`" + `, both ` + "`" + `http` + "`" + ` and ` + "`" + `https` + "`" + ` are permitted. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "signed_version",
					Description: `(Optional) Specifies the signed storage service version to use to authorize requests made with this account SAS. Defaults to ` + "`" + `2017-07-29` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "resource_types",
					Description: `A ` + "`" + `resource_types` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `A ` + "`" + `services` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `The starting time and date of validity of this SAS. Must be a valid ISO-8601 format time/date string.`,
				},
				resource.Attribute{
					Name:        "expiry",
					Description: `The expiration time and date of this SAS. Must be a valid ISO-8601 format time/date string.`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `A ` + "`" + `permissions` + "`" + ` block as defined below. --- ` + "`" + `resource_types` + "`" + ` is a set of ` + "`" + `true` + "`" + `/` + "`" + `false` + "`" + ` flags which define the storage account resource types that are granted access by this SAS. This can be thought of as the scope over which the permissions apply. A ` + "`" + `service` + "`" + ` will have larger scope (affecting all sub-resources) than ` + "`" + `object` + "`" + `. A ` + "`" + `resource_types` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `Should permission be granted to the entire service?`,
				},
				resource.Attribute{
					Name:        "container",
					Description: `Should permission be granted to the container?`,
				},
				resource.Attribute{
					Name:        "object",
					Description: `Should permission be granted only to a specific object? --- ` + "`" + `services` + "`" + ` is a set of ` + "`" + `true` + "`" + `/` + "`" + `false` + "`" + ` flags which define the storage account services that are granted access by this SAS. A ` + "`" + `services` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "blob",
					Description: `Should permission be granted to ` + "`" + `blob` + "`" + ` services within this storage account?`,
				},
				resource.Attribute{
					Name:        "queue",
					Description: `Should permission be granted to ` + "`" + `queue` + "`" + ` services within this storage account?`,
				},
				resource.Attribute{
					Name:        "table",
					Description: `Should permission be granted to ` + "`" + `table` + "`" + ` services within this storage account?`,
				},
				resource.Attribute{
					Name:        "file",
					Description: `Should permission be granted to ` + "`" + `file` + "`" + ` services within this storage account? --- A ` + "`" + `permissions` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `Should Read permissions be enabled for this SAS?`,
				},
				resource.Attribute{
					Name:        "write",
					Description: `Should Write permissions be enabled for this SAS?`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `Should Delete permissions be enabled for this SAS?`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `Should List permissions be enabled for this SAS?`,
				},
				resource.Attribute{
					Name:        "add",
					Description: `Should Add permissions be enabled for this SAS?`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `Should Create permissions be enabled for this SAS?`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `Should Update permissions be enabled for this SAS?`,
				},
				resource.Attribute{
					Name:        "process",
					Description: `Should Process permissions be enabled for this SAS? Refer to the [SAS creation reference from Azure](https://docs.microsoft.com/en-us/rest/api/storageservices/constructing-an-account-sas) for additional details on the fields above. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "sas",
					Description: `The computed Account Shared Access Signature (SAS). ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the SAS Token.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "sas",
					Description: `The computed Account Shared Access Signature (SAS). ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the SAS Token.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_storage_container",
			Category:         "Storage",
			ShortDescription: `Gets information about an existing Storage Container.`,
			Description: `

Use this data source to access information about an existing Storage Container.

`,
			Icon: "General/10839-icon-service-Storage-Container.svg",
			Keywords: []string{
				"storage",
				"container",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Container.`,
				},
				resource.Attribute{
					Name:        "storage_account_name",
					Description: `The name of the Storage Account where the Container exists. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "container_access_type",
					Description: `The Access Level configured for this Container.`,
				},
				resource.Attribute{
					Name:        "has_immutability_policy",
					Description: `Is there an Immutability Policy configured on this Storage Container?`,
				},
				resource.Attribute{
					Name:        "has_legal_hold",
					Description: `Is there a Legal Hold configured on this Storage Container?`,
				},
				resource.Attribute{
					Name:        "resource_manager_id",
					Description: `The Resource Manager ID of this Storage Container. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Storage Container.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "container_access_type",
					Description: `The Access Level configured for this Container.`,
				},
				resource.Attribute{
					Name:        "has_immutability_policy",
					Description: `Is there an Immutability Policy configured on this Storage Container?`,
				},
				resource.Attribute{
					Name:        "has_legal_hold",
					Description: `Is there a Legal Hold configured on this Storage Container?`,
				},
				resource.Attribute{
					Name:        "resource_manager_id",
					Description: `The Resource Manager ID of this Storage Container. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Storage Container.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_storage_encryption_scope",
			Category:         "Storage",
			ShortDescription: `Gets information about an existing Storage Encryption Scope.`,
			Description: `

Use this data source to access information about an existing Storage Encryption Scope.

`,
			Keywords: []string{
				"storage",
				"encryption",
				"scope",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Storage Encryption Scope.`,
				},
				resource.Attribute{
					Name:        "key_vault_key_id",
					Description: `The ID of the Key Vault Key.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `The source of the Storage Encryption Scope. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Storage Encryption Scope.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_storage_management_policy",
			Category:         "Storage",
			ShortDescription: `Gets information about an existing Storage Management Policy.`,
			Description: `

Use this data source to access information about an existing Storage Management Policy.

`,
			Keywords: []string{
				"storage",
				"management",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "storage_account_id",
					Description: `Specifies the id of the storage account to retrieve the management policy for. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Management Policy.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `A ` + "`" + `rule` + "`" + ` block as documented below. ---`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A rule name can contain any combination of alpha numeric characters. Rule name is case-sensitive. It must be unique within a policy.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Boolean to specify whether the rule is enabled.`,
				},
				resource.Attribute{
					Name:        "filters",
					Description: `A ` + "`" + `filter` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "actions",
					Description: `An ` + "`" + `actions` + "`" + ` block as documented below. --- ` + "`" + `filters` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "prefix_match",
					Description: `An array of strings for prefixes to be matched.`,
				},
				resource.Attribute{
					Name:        "blob_types",
					Description: `An array of predefined values. Valid options are ` + "`" + `blockBlob` + "`" + ` and ` + "`" + `appendBlob` + "`" + `. --- ` + "`" + `actions` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "base_blob",
					Description: `A ` + "`" + `base_blob` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "snapshot",
					Description: `A ` + "`" + `snapshot` + "`" + ` block as documented below. --- ` + "`" + `base_blob` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "tier_to_cool_after_days_since_modification_greater_than",
					Description: `The age in days after last modification to tier blobs to cool storage. Supports blob currently at Hot tier.`,
				},
				resource.Attribute{
					Name:        "tier_to_archive_after_days_since_modification_greater_than",
					Description: `The age in days after last modification to tier blobs to archive storage. Supports blob currently at Hot or Cool tier.`,
				},
				resource.Attribute{
					Name:        "delete_after_days_since_modification_greater_than",
					Description: `The age in days after last modification to delete the blob. --- ` + "`" + `snapshot` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "delete_after_days_since_creation_greater_than",
					Description: `The age in days after create to delete the snapshot. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Storage Management Policy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Management Policy.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `A ` + "`" + `rule` + "`" + ` block as documented below. ---`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A rule name can contain any combination of alpha numeric characters. Rule name is case-sensitive. It must be unique within a policy.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Boolean to specify whether the rule is enabled.`,
				},
				resource.Attribute{
					Name:        "filters",
					Description: `A ` + "`" + `filter` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "actions",
					Description: `An ` + "`" + `actions` + "`" + ` block as documented below. --- ` + "`" + `filters` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "prefix_match",
					Description: `An array of strings for prefixes to be matched.`,
				},
				resource.Attribute{
					Name:        "blob_types",
					Description: `An array of predefined values. Valid options are ` + "`" + `blockBlob` + "`" + ` and ` + "`" + `appendBlob` + "`" + `. --- ` + "`" + `actions` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "base_blob",
					Description: `A ` + "`" + `base_blob` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "snapshot",
					Description: `A ` + "`" + `snapshot` + "`" + ` block as documented below. --- ` + "`" + `base_blob` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "tier_to_cool_after_days_since_modification_greater_than",
					Description: `The age in days after last modification to tier blobs to cool storage. Supports blob currently at Hot tier.`,
				},
				resource.Attribute{
					Name:        "tier_to_archive_after_days_since_modification_greater_than",
					Description: `The age in days after last modification to tier blobs to archive storage. Supports blob currently at Hot or Cool tier.`,
				},
				resource.Attribute{
					Name:        "delete_after_days_since_modification_greater_than",
					Description: `The age in days after last modification to delete the blob. --- ` + "`" + `snapshot` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "delete_after_days_since_creation_greater_than",
					Description: `The age in days after create to delete the snapshot. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Storage Management Policy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_storage_sync",
			Category:         "Storage",
			ShortDescription: `Gets information about an existing Storage Sync.`,
			Description: `

Use this data source to access information about an existing Storage Sync.

`,
			Keywords: []string{
				"storage",
				"sync",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Storage Sync.`,
				},
				resource.Attribute{
					Name:        "incoming_traffic_policy",
					Description: `Incoming traffic policy.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure Region where the Storage Sync exists.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the Storage Sync. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Storage Sync.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_storage_sync_group",
			Category:         "Storage",
			ShortDescription: `Gets information about an existing Storage Sync Group.`,
			Description: `

Use this data source to access information about an existing Storage Sync Group.

`,
			Keywords: []string{
				"storage",
				"sync",
				"group",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Storage Sync Group. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Storage Sync Group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_stream_analytics_job",
			Category:         "Stream Analytics",
			ShortDescription: `Gets information about an existing Stream Analytics Job.`,
			Description: `

Use this data source to access information about an existing Stream Analytics Job.

`,
			Icon: "IoT/00042-icon-service-Stream-Analytics-Jobs.svg",
			Keywords: []string{
				"stream",
				"analytics",
				"job",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name of the Stream Analytics Job.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `Specifies the name of the resource group the Stream Analytics Job is located in. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Stream Analytics Job.`,
				},
				resource.Attribute{
					Name:        "compatibility_level",
					Description: `The compatibility level for this job.`,
				},
				resource.Attribute{
					Name:        "data_locale",
					Description: `The Data Locale of the Job.`,
				},
				resource.Attribute{
					Name:        "events_late_arrival_max_delay_in_seconds",
					Description: `The maximum tolerable delay in seconds where events arriving late could be included.`,
				},
				resource.Attribute{
					Name:        "events_out_of_order_max_delay_in_seconds",
					Description: `The maximum tolerable delay in seconds where out-of-order events can be adjusted to be back in order.`,
				},
				resource.Attribute{
					Name:        "events_out_of_order_policy",
					Description: `The policy which should be applied to events which arrive out of order in the input event stream.`,
				},
				resource.Attribute{
					Name:        "job_id",
					Description: `The Job ID assigned by the Stream Analytics Job.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the Stream Analytics Job exists.`,
				},
				resource.Attribute{
					Name:        "output_error_policy",
					Description: `The policy which should be applied to events which arrive at the output and cannot be written to the external storage due to being malformed (such as missing column values, column values of wrong type or size).`,
				},
				resource.Attribute{
					Name:        "streaming_units",
					Description: `The number of streaming units that the streaming job uses.`,
				},
				resource.Attribute{
					Name:        "transformation_query",
					Description: `The query that will be run in the streaming job, [written in Stream Analytics Query Language (SAQL)](https://msdn.microsoft.com/library/azure/dn834998). ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Stream Analytics Job.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Stream Analytics Job.`,
				},
				resource.Attribute{
					Name:        "compatibility_level",
					Description: `The compatibility level for this job.`,
				},
				resource.Attribute{
					Name:        "data_locale",
					Description: `The Data Locale of the Job.`,
				},
				resource.Attribute{
					Name:        "events_late_arrival_max_delay_in_seconds",
					Description: `The maximum tolerable delay in seconds where events arriving late could be included.`,
				},
				resource.Attribute{
					Name:        "events_out_of_order_max_delay_in_seconds",
					Description: `The maximum tolerable delay in seconds where out-of-order events can be adjusted to be back in order.`,
				},
				resource.Attribute{
					Name:        "events_out_of_order_policy",
					Description: `The policy which should be applied to events which arrive out of order in the input event stream.`,
				},
				resource.Attribute{
					Name:        "job_id",
					Description: `The Job ID assigned by the Stream Analytics Job.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the Stream Analytics Job exists.`,
				},
				resource.Attribute{
					Name:        "output_error_policy",
					Description: `The policy which should be applied to events which arrive at the output and cannot be written to the external storage due to being malformed (such as missing column values, column values of wrong type or size).`,
				},
				resource.Attribute{
					Name:        "streaming_units",
					Description: `The number of streaming units that the streaming job uses.`,
				},
				resource.Attribute{
					Name:        "transformation_query",
					Description: `The query that will be run in the streaming job, [written in Stream Analytics Query Language (SAQL)](https://msdn.microsoft.com/library/azure/dn834998). ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Stream Analytics Job.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_subnet",
			Category:         "Network",
			ShortDescription: `Gets information about an existing Subnet located within a Virtual Network.`,
			Description: `

Use this data source to access information about an existing Subnet within a Virtual Network.

`,
			Keywords: []string{
				"network",
				"subnet",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name of the Subnet.`,
				},
				resource.Attribute{
					Name:        "virtual_network_name",
					Description: `Specifies the name of the Virtual Network this Subnet is located within.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `Specifies the name of the resource group the Virtual Network is located in. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Subnet.`,
				},
				resource.Attribute{
					Name:        "address_prefix",
					Description: `(Deprecated) The address prefix used for the subnet.`,
				},
				resource.Attribute{
					Name:        "address_prefixes",
					Description: `The address prefixes for the subnet.`,
				},
				resource.Attribute{
					Name:        "enforce_private_link_service_network_policies",
					Description: `Enable or Disable network policies on private link service in the subnet.`,
				},
				resource.Attribute{
					Name:        "network_security_group_id",
					Description: `The ID of the Network Security Group associated with the subnet.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The ID of the Route Table associated with this subnet.`,
				},
				resource.Attribute{
					Name:        "service_endpoints",
					Description: `A list of Service Endpoints within this subnet.`,
				},
				resource.Attribute{
					Name:        "enforce_private_link_endpoint_network_policies",
					Description: `Enable or Disable network policies for the private link endpoint on the subnet.`,
				},
				resource.Attribute{
					Name:        "enforce_private_link_service_network_policies",
					Description: `Enable or Disable network policies for the private link service on the subnet. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Subnet located within a Virtual Network.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Subnet.`,
				},
				resource.Attribute{
					Name:        "address_prefix",
					Description: `(Deprecated) The address prefix used for the subnet.`,
				},
				resource.Attribute{
					Name:        "address_prefixes",
					Description: `The address prefixes for the subnet.`,
				},
				resource.Attribute{
					Name:        "enforce_private_link_service_network_policies",
					Description: `Enable or Disable network policies on private link service in the subnet.`,
				},
				resource.Attribute{
					Name:        "network_security_group_id",
					Description: `The ID of the Network Security Group associated with the subnet.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The ID of the Route Table associated with this subnet.`,
				},
				resource.Attribute{
					Name:        "service_endpoints",
					Description: `A list of Service Endpoints within this subnet.`,
				},
				resource.Attribute{
					Name:        "enforce_private_link_endpoint_network_policies",
					Description: `Enable or Disable network policies for the private link endpoint on the subnet.`,
				},
				resource.Attribute{
					Name:        "enforce_private_link_service_network_policies",
					Description: `Enable or Disable network policies for the private link service on the subnet. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Subnet located within a Virtual Network.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_subscription",
			Category:         "Base",
			ShortDescription: `Gets information about an existing Subscription.`,
			Description: `

Use this data source to access information about an existing Subscription.

`,
			Keywords: []string{
				"base",
				"subscription",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "subscription_id",
					Description: `(Optional) Specifies the ID of the subscription. If this argument is omitted, the subscription ID of the current Azure Resource Manager provider is used. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the subscription.`,
				},
				resource.Attribute{
					Name:        "subscription_id",
					Description: `The subscription GUID.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The subscription display name.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `The subscription tenant ID.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The subscription state. Possible values are Enabled, Warned, PastDue, Disabled, and Deleted.`,
				},
				resource.Attribute{
					Name:        "location_placement_id",
					Description: `The subscription location placement ID.`,
				},
				resource.Attribute{
					Name:        "quota_id",
					Description: `The subscription quota ID.`,
				},
				resource.Attribute{
					Name:        "spending_limit",
					Description: `The subscription spending limit.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the Subscription. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Subscription.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the subscription.`,
				},
				resource.Attribute{
					Name:        "subscription_id",
					Description: `The subscription GUID.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The subscription display name.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `The subscription tenant ID.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The subscription state. Possible values are Enabled, Warned, PastDue, Disabled, and Deleted.`,
				},
				resource.Attribute{
					Name:        "location_placement_id",
					Description: `The subscription location placement ID.`,
				},
				resource.Attribute{
					Name:        "quota_id",
					Description: `The subscription quota ID.`,
				},
				resource.Attribute{
					Name:        "spending_limit",
					Description: `The subscription spending limit.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the Subscription. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Subscription.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_subscriptions",
			Category:         "Base",
			ShortDescription: `Get information about the available subscriptions.`,
			Description: `

Use this data source to access information about all the Subscriptions currently available.

`,
			Keywords: []string{
				"base",
				"subscriptions",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name_prefix",
					Description: `(Optional) A case-insensitive prefix which can be used to filter on the ` + "`" + `display_name` + "`" + ` field`,
				},
				resource.Attribute{
					Name:        "display_name_contains",
					Description: `(Optional) A case-insensitive value which must be contained within the ` + "`" + `display_name` + "`" + ` field, used to filter the results ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "subscriptions",
					Description: `One or more ` + "`" + `subscription` + "`" + ` blocks as defined below. The ` + "`" + `subscription` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of this subscription.`,
				},
				resource.Attribute{
					Name:        "subscription_id",
					Description: `The subscription GUID.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The subscription display name.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `The subscription tenant ID.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The subscription state. Possible values are Enabled, Warned, PastDue, Disabled, and Deleted.`,
				},
				resource.Attribute{
					Name:        "location_placement_id",
					Description: `The subscription location placement ID.`,
				},
				resource.Attribute{
					Name:        "quota_id",
					Description: `The subscription quota ID.`,
				},
				resource.Attribute{
					Name:        "spending_limit",
					Description: `The subscription spending limit.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the subscriptions.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "subscriptions",
					Description: `One or more ` + "`" + `subscription` + "`" + ` blocks as defined below. The ` + "`" + `subscription` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of this subscription.`,
				},
				resource.Attribute{
					Name:        "subscription_id",
					Description: `The subscription GUID.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The subscription display name.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `The subscription tenant ID.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The subscription state. Possible values are Enabled, Warned, PastDue, Disabled, and Deleted.`,
				},
				resource.Attribute{
					Name:        "location_placement_id",
					Description: `The subscription location placement ID.`,
				},
				resource.Attribute{
					Name:        "quota_id",
					Description: `The subscription quota ID.`,
				},
				resource.Attribute{
					Name:        "spending_limit",
					Description: `The subscription spending limit.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the subscriptions.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_synapse_workspace",
			Category:         "Synapse",
			ShortDescription: `Gets information about an existing Synapse Workspace.`,
			Description: `

Use this data source to access information about an existing Synapse Workspace.

`,
			Keywords: []string{
				"synapse",
				"workspace",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the synapse Workspace.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the Synapse Workspace exists.`,
				},
				resource.Attribute{
					Name:        "connectivity_endpoints",
					Description: `A list of Connectivity endpoints for this Synapse Workspace.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Synapse Workspace.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_template_spec_version",
			Category:         "Template",
			ShortDescription: `Gets information about an existing Template Spec Version.`,
			Description: `

Use this data source to access information about an existing Template Spec Version.

`,
			Keywords: []string{
				"template",
				"spec",
				"version",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Template Spec version.`,
				},
				resource.Attribute{
					Name:        "template_body",
					Description: `The ARM Template body of the Template Spec Version.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the Template. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Template.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_traffic_manager_geographical_location",
			Category:         "Network",
			ShortDescription: `Gets information about a specified Traffic Manager Geographical Location within the Geographical Hierarchy.`,
			Description: `

Use this data source to access the ID of a specified Traffic Manager Geographical Location within the Geographical Hierarchy.

`,
			Keywords: []string{
				"network",
				"traffic",
				"manager",
				"geographical",
				"location",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name of the Location, for example ` + "`" + `World` + "`" + `, ` + "`" + `Europe` + "`" + ` or ` + "`" + `Germany` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of this Location, also known as the ` + "`" + `Code` + "`" + ` of this Location. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Location.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of this Location, also known as the ` + "`" + `Code` + "`" + ` of this Location. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Location.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_traffic_manager_profile",
			Category:         "Network",
			ShortDescription: `Gets information about a specified Traffic Manager Profile.`,
			Description: `

Use this data source to access information about an existing Traffic Manager Profile.

`,
			Icon: "Networking/10065-icon-service-Traffic-Manager-Profiles.svg",
			Keywords: []string{
				"network",
				"traffic",
				"manager",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name of the Traffic Manager Profile.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `Specifies the name of the resource group the Traffic Manager Profile is located in. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Traffic Manager Profile.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the Traffic Manager Profile exists.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `The FQDN of the created Profile.`,
				},
				resource.Attribute{
					Name:        "profile_status",
					Description: `The status of the profile.`,
				},
				resource.Attribute{
					Name:        "traffic_routing_method",
					Description: `Specifies the algorithm used to route traffic.`,
				},
				resource.Attribute{
					Name:        "traffic_view_enabled",
					Description: `Indicates whether Traffic View is enabled for the Traffic Manager profile.`,
				},
				resource.Attribute{
					Name:        "dns_config",
					Description: `This block specifies the DNS configuration of the Profile.`,
				},
				resource.Attribute{
					Name:        "monitor_config",
					Description: `This block specifies the Endpoint monitoring configuration for the Profile.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the resource. The ` + "`" + `dns_config` + "`" + ` block provides:`,
				},
				resource.Attribute{
					Name:        "relative_name",
					Description: `The relative domain name, this is combined with the domain name used by Traffic Manager to form the FQDN which is exported as documented below.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `The TTL value of the Profile used by Local DNS resolvers and clients. The ` + "`" + `monitor_config` + "`" + ` block provides:`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol used by the monitoring checks.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port number used by the monitoring checks.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The path used by the monitoring checks.`,
				},
				resource.Attribute{
					Name:        "expected_status_code_ranges",
					Description: `A list of status code ranges.`,
				},
				resource.Attribute{
					Name:        "custom_header",
					Description: `One or more ` + "`" + `custom_header` + "`" + ` blocks as defined below.`,
				},
				resource.Attribute{
					Name:        "interval_in_seconds",
					Description: `The interval used to check the endpoint health from a Traffic Manager probing agent.`,
				},
				resource.Attribute{
					Name:        "timeout_in_seconds",
					Description: `The amount of time the Traffic Manager probing agent should wait before considering that check a failure when a health check probe is sent to the endpoint.`,
				},
				resource.Attribute{
					Name:        "tolerated_number_of_failures",
					Description: `The number of failures a Traffic Manager probing agent tolerates before marking that endpoint as unhealthy. A ` + "`" + `custom_header` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the custom header.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of custom header. Applicable for Http and Https protocol. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Location.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Traffic Manager Profile.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the Traffic Manager Profile exists.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `The FQDN of the created Profile.`,
				},
				resource.Attribute{
					Name:        "profile_status",
					Description: `The status of the profile.`,
				},
				resource.Attribute{
					Name:        "traffic_routing_method",
					Description: `Specifies the algorithm used to route traffic.`,
				},
				resource.Attribute{
					Name:        "traffic_view_enabled",
					Description: `Indicates whether Traffic View is enabled for the Traffic Manager profile.`,
				},
				resource.Attribute{
					Name:        "dns_config",
					Description: `This block specifies the DNS configuration of the Profile.`,
				},
				resource.Attribute{
					Name:        "monitor_config",
					Description: `This block specifies the Endpoint monitoring configuration for the Profile.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the resource. The ` + "`" + `dns_config` + "`" + ` block provides:`,
				},
				resource.Attribute{
					Name:        "relative_name",
					Description: `The relative domain name, this is combined with the domain name used by Traffic Manager to form the FQDN which is exported as documented below.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `The TTL value of the Profile used by Local DNS resolvers and clients. The ` + "`" + `monitor_config` + "`" + ` block provides:`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol used by the monitoring checks.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port number used by the monitoring checks.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The path used by the monitoring checks.`,
				},
				resource.Attribute{
					Name:        "expected_status_code_ranges",
					Description: `A list of status code ranges.`,
				},
				resource.Attribute{
					Name:        "custom_header",
					Description: `One or more ` + "`" + `custom_header` + "`" + ` blocks as defined below.`,
				},
				resource.Attribute{
					Name:        "interval_in_seconds",
					Description: `The interval used to check the endpoint health from a Traffic Manager probing agent.`,
				},
				resource.Attribute{
					Name:        "timeout_in_seconds",
					Description: `The amount of time the Traffic Manager probing agent should wait before considering that check a failure when a health check probe is sent to the endpoint.`,
				},
				resource.Attribute{
					Name:        "tolerated_number_of_failures",
					Description: `The number of failures a Traffic Manager probing agent tolerates before marking that endpoint as unhealthy. A ` + "`" + `custom_header` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the custom header.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of custom header. Applicable for Http and Https protocol. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Location.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_user_assigned_identity",
			Category:         "Authorization",
			ShortDescription: `Gets information about an existing User Assigned Identity.`,
			Description: `

Use this data source to access information about an existing User Assigned Identity.

`,
			Keywords: []string{
				"authorization",
				"user",
				"assigned",
				"identity",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the User Assigned Identity.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of the Resource Group in which the User Assigned Identity exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the User Assigned Identity.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the User Assigned Identity exists.`,
				},
				resource.Attribute{
					Name:        "principal_id",
					Description: `The Service Principal ID of the User Assigned Identity.`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `The Client ID of the User Assigned Identity.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the User Assigned Identity. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the User Assigned Identity.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the User Assigned Identity.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the User Assigned Identity exists.`,
				},
				resource.Attribute{
					Name:        "principal_id",
					Description: `The Service Principal ID of the User Assigned Identity.`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `The Client ID of the User Assigned Identity.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the User Assigned Identity. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the User Assigned Identity.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_virtual_hub",
			Category:         "Network",
			ShortDescription: `Gets information about an existing Virtual Hub`,
			Description: `

Uses this data source to access information about an existing Virtual Hub.

`,
			Keywords: []string{
				"network",
				"virtual",
				"hub",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Virtual Hub.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The Name of the Resource Group where the Virtual Hub exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure Region where the Virtual Hub exists.`,
				},
				resource.Attribute{
					Name:        "address_prefix",
					Description: `The Address Prefix used for this Virtual Hub.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the Virtual Hub.`,
				},
				resource.Attribute{
					Name:        "virtual_wan_id",
					Description: `The ID of the Virtual WAN within which the Virtual Hub exists. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Virtual Hub.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "location",
					Description: `The Azure Region where the Virtual Hub exists.`,
				},
				resource.Attribute{
					Name:        "address_prefix",
					Description: `The Address Prefix used for this Virtual Hub.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the Virtual Hub.`,
				},
				resource.Attribute{
					Name:        "virtual_wan_id",
					Description: `The ID of the Virtual WAN within which the Virtual Hub exists. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Virtual Hub.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_virtual_machine",
			Category:         "Compute",
			ShortDescription: `Gets information about an existing Virtual Machine.`,
			Description: `

Use this data source to access information about an existing Virtual Machine.

`,
			Icon: "Compute/10040-icon-service-VM-Images-(Classic).svg",
			Keywords: []string{
				"compute",
				"virtual",
				"machine",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name of the Virtual Machine.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `Specifies the name of the resource group the Virtual Machine is located in. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Virtual Machine.`,
				},
				resource.Attribute{
					Name:        "identity",
					Description: `A ` + "`" + `identity` + "`" + ` block as defined below. --- An ` + "`" + `identity` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "identity_ids",
					Description: `The list of User Managed Identity ID's which are assigned to the Virtual Machine.`,
				},
				resource.Attribute{
					Name:        "principal_id",
					Description: `The ID of the System Managed Service Principal assigned to the Virtual Machine.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `The ID of the Tenant of the System Managed Service Principal assigned to the Virtual Machine.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The identity type of the Managed Identity assigned to the Virtual Machine. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Virtual Machine.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Virtual Machine.`,
				},
				resource.Attribute{
					Name:        "identity",
					Description: `A ` + "`" + `identity` + "`" + ` block as defined below. --- An ` + "`" + `identity` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "identity_ids",
					Description: `The list of User Managed Identity ID's which are assigned to the Virtual Machine.`,
				},
				resource.Attribute{
					Name:        "principal_id",
					Description: `The ID of the System Managed Service Principal assigned to the Virtual Machine.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `The ID of the Tenant of the System Managed Service Principal assigned to the Virtual Machine.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The identity type of the Managed Identity assigned to the Virtual Machine. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Virtual Machine.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_virtual_machine_scale_set",
			Category:         "Compute",
			ShortDescription: `Gets information about an existing Virtual Machine Scale Set.`,
			Description: `

Use this data source to access information about an existing Virtual Machine Scale Set.

`,
			Icon: "Compute/10034-icon-service-VM-Scale-Sets.svg",
			Keywords: []string{
				"compute",
				"virtual",
				"machine",
				"scale",
				"set",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Virtual Machine Scale Set.`,
				},
				resource.Attribute{
					Name:        "identity",
					Description: `A ` + "`" + `identity` + "`" + ` block as defined below. --- A ` + "`" + `identity` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "identity_ids",
					Description: `The list of User Managed Identity ID's which are assigned to the Virtual Machine Scale Set.`,
				},
				resource.Attribute{
					Name:        "principal_id",
					Description: `The ID of the System Managed Service Principal assigned to the Virtual Machine Scale Set.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The identity type of the Managed Identity assigned to the Virtual Machine Scale Set. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Virtual Machine Scale Set.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_virtual_network",
			Category:         "Network",
			ShortDescription: `Gets information about an existing Virtual Network.`,
			Description: `

Use this data source to access information about an existing Virtual Network.

`,
			Icon: "Networking/10075-icon-service-Virtual-Networks-(Classic).svg",
			Keywords: []string{
				"network",
				"virtual",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name of the Virtual Network.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `Specifies the name of the resource group the Virtual Network is located in. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the virtual network.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Location of the virtual network.`,
				},
				resource.Attribute{
					Name:        "address_space",
					Description: `The list of address spaces used by the virtual network.`,
				},
				resource.Attribute{
					Name:        "dns_servers",
					Description: `The list of DNS servers used by the virtual network.`,
				},
				resource.Attribute{
					Name:        "guid",
					Description: `The GUID of the virtual network.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `The list of name of the subnets that are attached to this virtual network.`,
				},
				resource.Attribute{
					Name:        "vnet_peerings",
					Description: `A mapping of name - virtual network id of the virtual network peerings. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Virtual Network.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the virtual network.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Location of the virtual network.`,
				},
				resource.Attribute{
					Name:        "address_space",
					Description: `The list of address spaces used by the virtual network.`,
				},
				resource.Attribute{
					Name:        "dns_servers",
					Description: `The list of DNS servers used by the virtual network.`,
				},
				resource.Attribute{
					Name:        "guid",
					Description: `The GUID of the virtual network.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `The list of name of the subnets that are attached to this virtual network.`,
				},
				resource.Attribute{
					Name:        "vnet_peerings",
					Description: `A mapping of name - virtual network id of the virtual network peerings. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Virtual Network.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_virtual_network_gateway",
			Category:         "Network",
			ShortDescription: `Gets information about an existing Virtual Network Gateway.`,
			Description: `

Use this data source to access information about an existing Virtual Network Gateway.

`,
			Icon: "Networking/10063-icon-service-Virtual-Network-Gateways.svg",
			Keywords: []string{
				"network",
				"virtual",
				"gateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name of the Virtual Network Gateway.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `Specifies the name of the resource group the Virtual Network Gateway is located in. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Virtual Network Gateway.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location/region where the Virtual Network Gateway is located.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the Virtual Network Gateway.`,
				},
				resource.Attribute{
					Name:        "vpn_type",
					Description: `The routing type of the Virtual Network Gateway.`,
				},
				resource.Attribute{
					Name:        "enable_bgp",
					Description: `Will BGP (Border Gateway Protocol) will be enabled for this Virtual Network Gateway.`,
				},
				resource.Attribute{
					Name:        "active_active",
					Description: `Is this an Active-Active Gateway?`,
				},
				resource.Attribute{
					Name:        "private_ip_address_enabled",
					Description: `Whether a private IP will be used for this gateway for connections.`,
				},
				resource.Attribute{
					Name:        "default_local_network_gateway_id",
					Description: `The ID of the local network gateway through which outbound Internet traffic from the virtual network in which the gateway is created will be routed (`,
				},
				resource.Attribute{
					Name:        "sku",
					Description: `Configuration of the size and capacity of the Virtual Network Gateway.`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `The Generation of the Virtual Network Gateway.`,
				},
				resource.Attribute{
					Name:        "ip_configuration",
					Description: `One or two ` + "`" + `ip_configuration` + "`" + ` blocks documented below.`,
				},
				resource.Attribute{
					Name:        "vpn_client_configuration",
					Description: `A ` + "`" + `vpn_client_configuration` + "`" + ` block which is documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource. The ` + "`" + `ip_configuration` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A user-defined name of the IP configuration.`,
				},
				resource.Attribute{
					Name:        "private_ip_address_allocation",
					Description: `Defines how the private IP address of the gateways virtual interface is assigned.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the gateway subnet of a virtual network in which the virtual network gateway will be created. It is mandatory that the associated subnet is named ` + "`" + `GatewaySubnet` + "`" + `. Therefore, each virtual network can contain at most a single Virtual Network Gateway.`,
				},
				resource.Attribute{
					Name:        "public_ip_address_id",
					Description: `The ID of the Public IP Address associated with the Virtual Network Gateway. The ` + "`" + `vpn_client_configuration` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "address_space",
					Description: `The address space out of which ip addresses for vpn clients will be taken. You can provide more than one address space, e.g. in CIDR notation.`,
				},
				resource.Attribute{
					Name:        "aad_tenant",
					Description: `AzureAD Tenant URL This setting is incompatible with the use of ` + "`" + `root_certificate` + "`" + ` and ` + "`" + `revoked_certificate` + "`" + `, ` + "`" + `radius_server_address` + "`" + `, and ` + "`" + `radius_server_secret` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "aad_audience",
					Description: `The client id of the Azure VPN application. See [Create an Active Directory (AD) tenant for P2S OpenVPN protocol connections](https://docs.microsoft.com/en-gb/azure/vpn-gateway/openvpn-azure-ad-tenant-multi-app) for values This setting is incompatible with the use of ` + "`" + `root_certificate` + "`" + ` and ` + "`" + `revoked_certificate` + "`" + `, ` + "`" + `radius_server_address` + "`" + `, and ` + "`" + `radius_server_secret` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "aad_issuer",
					Description: `The STS url for your tenant This setting is incompatible with the use of ` + "`" + `root_certificate` + "`" + ` and ` + "`" + `revoked_certificate` + "`" + `, ` + "`" + `radius_server_address` + "`" + `, and ` + "`" + `radius_server_secret` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "root_certificate",
					Description: `One or more ` + "`" + `root_certificate` + "`" + ` blocks which are defined below. These root certificates are used to sign the client certificate used by the VPN clients to connect to the gateway. This setting is incompatible with the use of ` + "`" + `aad_tenant` + "`" + `, ` + "`" + `aad_audience` + "`" + `, ` + "`" + `aad_issuer` + "`" + `, ` + "`" + `radius_server_address` + "`" + `, and ` + "`" + `radius_server_secret` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "revoked_certificate",
					Description: `One or more ` + "`" + `revoked_certificate` + "`" + ` blocks which are defined below. This setting is incompatible with the use of ` + "`" + `aad_tenant` + "`" + `, ` + "`" + `aad_audience` + "`" + `, ` + "`" + `aad_issuer` + "`" + `, ` + "`" + `radius_server_address` + "`" + `, and ` + "`" + `radius_server_secret` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "radius_server_address",
					Description: `The address of the Radius server. This setting is incompatible with the use of ` + "`" + `aad_tenant` + "`" + `, ` + "`" + `aad_audience` + "`" + `, ` + "`" + `aad_issuer` + "`" + `, ` + "`" + `root_certificate` + "`" + ` and ` + "`" + `revoked_certificate` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "radius_server_secret",
					Description: `The secret used by the Radius server. This setting is incompatible with the use of ` + "`" + `aad_tenant` + "`" + `, ` + "`" + `aad_audience` + "`" + `, ` + "`" + `aad_issuer` + "`" + `, ` + "`" + `root_certificate` + "`" + ` and ` + "`" + `revoked_certificate` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vpn_client_protocols",
					Description: `List of the protocols supported by the vpn client. The supported values are ` + "`" + `SSTP` + "`" + `, ` + "`" + `IkeV2` + "`" + ` and ` + "`" + `OpenVPN` + "`" + `. The ` + "`" + `bgp_settings` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "asn",
					Description: `The Autonomous System Number (ASN) to use as part of the BGP.`,
				},
				resource.Attribute{
					Name:        "peering_address",
					Description: `The BGP peer IP address of the virtual network gateway. This address is needed to configure the created gateway as a BGP Peer on the on-premises VPN devices.`,
				},
				resource.Attribute{
					Name:        "peer_weight",
					Description: `The weight added to routes which have been learned through BGP peering. A ` + "`" + `custom_route` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "address_prefixes",
					Description: `A list of address blocks reserved for this virtual network in CIDR notation. The ` + "`" + `root_certificate` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The user-defined name of the root certificate.`,
				},
				resource.Attribute{
					Name:        "public_cert_data",
					Description: `The public certificate of the root certificate authority. The certificate must be provided in Base-64 encoded X.509 format (PEM). The ` + "`" + `root_revoked_certificate` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The user-defined name of the revoked certificate.`,
				},
				resource.Attribute{
					Name:        "public_cert_data",
					Description: `The SHA1 thumbprint of the certificate to be revoked. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Virtual Network Gateway.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Virtual Network Gateway.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location/region where the Virtual Network Gateway is located.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the Virtual Network Gateway.`,
				},
				resource.Attribute{
					Name:        "vpn_type",
					Description: `The routing type of the Virtual Network Gateway.`,
				},
				resource.Attribute{
					Name:        "enable_bgp",
					Description: `Will BGP (Border Gateway Protocol) will be enabled for this Virtual Network Gateway.`,
				},
				resource.Attribute{
					Name:        "active_active",
					Description: `Is this an Active-Active Gateway?`,
				},
				resource.Attribute{
					Name:        "private_ip_address_enabled",
					Description: `Whether a private IP will be used for this gateway for connections.`,
				},
				resource.Attribute{
					Name:        "default_local_network_gateway_id",
					Description: `The ID of the local network gateway through which outbound Internet traffic from the virtual network in which the gateway is created will be routed (`,
				},
				resource.Attribute{
					Name:        "sku",
					Description: `Configuration of the size and capacity of the Virtual Network Gateway.`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `The Generation of the Virtual Network Gateway.`,
				},
				resource.Attribute{
					Name:        "ip_configuration",
					Description: `One or two ` + "`" + `ip_configuration` + "`" + ` blocks documented below.`,
				},
				resource.Attribute{
					Name:        "vpn_client_configuration",
					Description: `A ` + "`" + `vpn_client_configuration` + "`" + ` block which is documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource. The ` + "`" + `ip_configuration` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A user-defined name of the IP configuration.`,
				},
				resource.Attribute{
					Name:        "private_ip_address_allocation",
					Description: `Defines how the private IP address of the gateways virtual interface is assigned.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the gateway subnet of a virtual network in which the virtual network gateway will be created. It is mandatory that the associated subnet is named ` + "`" + `GatewaySubnet` + "`" + `. Therefore, each virtual network can contain at most a single Virtual Network Gateway.`,
				},
				resource.Attribute{
					Name:        "public_ip_address_id",
					Description: `The ID of the Public IP Address associated with the Virtual Network Gateway. The ` + "`" + `vpn_client_configuration` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "address_space",
					Description: `The address space out of which ip addresses for vpn clients will be taken. You can provide more than one address space, e.g. in CIDR notation.`,
				},
				resource.Attribute{
					Name:        "aad_tenant",
					Description: `AzureAD Tenant URL This setting is incompatible with the use of ` + "`" + `root_certificate` + "`" + ` and ` + "`" + `revoked_certificate` + "`" + `, ` + "`" + `radius_server_address` + "`" + `, and ` + "`" + `radius_server_secret` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "aad_audience",
					Description: `The client id of the Azure VPN application. See [Create an Active Directory (AD) tenant for P2S OpenVPN protocol connections](https://docs.microsoft.com/en-gb/azure/vpn-gateway/openvpn-azure-ad-tenant-multi-app) for values This setting is incompatible with the use of ` + "`" + `root_certificate` + "`" + ` and ` + "`" + `revoked_certificate` + "`" + `, ` + "`" + `radius_server_address` + "`" + `, and ` + "`" + `radius_server_secret` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "aad_issuer",
					Description: `The STS url for your tenant This setting is incompatible with the use of ` + "`" + `root_certificate` + "`" + ` and ` + "`" + `revoked_certificate` + "`" + `, ` + "`" + `radius_server_address` + "`" + `, and ` + "`" + `radius_server_secret` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "root_certificate",
					Description: `One or more ` + "`" + `root_certificate` + "`" + ` blocks which are defined below. These root certificates are used to sign the client certificate used by the VPN clients to connect to the gateway. This setting is incompatible with the use of ` + "`" + `aad_tenant` + "`" + `, ` + "`" + `aad_audience` + "`" + `, ` + "`" + `aad_issuer` + "`" + `, ` + "`" + `radius_server_address` + "`" + `, and ` + "`" + `radius_server_secret` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "revoked_certificate",
					Description: `One or more ` + "`" + `revoked_certificate` + "`" + ` blocks which are defined below. This setting is incompatible with the use of ` + "`" + `aad_tenant` + "`" + `, ` + "`" + `aad_audience` + "`" + `, ` + "`" + `aad_issuer` + "`" + `, ` + "`" + `radius_server_address` + "`" + `, and ` + "`" + `radius_server_secret` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "radius_server_address",
					Description: `The address of the Radius server. This setting is incompatible with the use of ` + "`" + `aad_tenant` + "`" + `, ` + "`" + `aad_audience` + "`" + `, ` + "`" + `aad_issuer` + "`" + `, ` + "`" + `root_certificate` + "`" + ` and ` + "`" + `revoked_certificate` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "radius_server_secret",
					Description: `The secret used by the Radius server. This setting is incompatible with the use of ` + "`" + `aad_tenant` + "`" + `, ` + "`" + `aad_audience` + "`" + `, ` + "`" + `aad_issuer` + "`" + `, ` + "`" + `root_certificate` + "`" + ` and ` + "`" + `revoked_certificate` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vpn_client_protocols",
					Description: `List of the protocols supported by the vpn client. The supported values are ` + "`" + `SSTP` + "`" + `, ` + "`" + `IkeV2` + "`" + ` and ` + "`" + `OpenVPN` + "`" + `. The ` + "`" + `bgp_settings` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "asn",
					Description: `The Autonomous System Number (ASN) to use as part of the BGP.`,
				},
				resource.Attribute{
					Name:        "peering_address",
					Description: `The BGP peer IP address of the virtual network gateway. This address is needed to configure the created gateway as a BGP Peer on the on-premises VPN devices.`,
				},
				resource.Attribute{
					Name:        "peer_weight",
					Description: `The weight added to routes which have been learned through BGP peering. A ` + "`" + `custom_route` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "address_prefixes",
					Description: `A list of address blocks reserved for this virtual network in CIDR notation. The ` + "`" + `root_certificate` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The user-defined name of the root certificate.`,
				},
				resource.Attribute{
					Name:        "public_cert_data",
					Description: `The public certificate of the root certificate authority. The certificate must be provided in Base-64 encoded X.509 format (PEM). The ` + "`" + `root_revoked_certificate` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The user-defined name of the revoked certificate.`,
				},
				resource.Attribute{
					Name:        "public_cert_data",
					Description: `The SHA1 thumbprint of the certificate to be revoked. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Virtual Network Gateway.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_virtual_network_gateway_connection",
			Category:         "Network",
			ShortDescription: `Gets information about an existing Virtual Network Gateway Connection.`,
			Description: `

Use this data source to access information about an existing Virtual Network Gateway Connection.

`,
			Keywords: []string{
				"network",
				"virtual",
				"gateway",
				"connection",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name of the Virtual Network Gateway Connection.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `Specifies the name of the resource group the Virtual Network Gateway Connection is located in. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Virtual Network Gateway Connection.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location/region where the connection is located.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of connection. Valid options are ` + "`" + `IPsec` + "`" + ` (Site-to-Site), ` + "`" + `ExpressRoute` + "`" + ` (ExpressRoute), and ` + "`" + `Vnet2Vnet` + "`" + ` (VNet-to-VNet).`,
				},
				resource.Attribute{
					Name:        "virtual_network_gateway_id",
					Description: `The ID of the Virtual Network Gateway in which the connection is created.`,
				},
				resource.Attribute{
					Name:        "authorization_key",
					Description: `The authorization key associated with the Express Route Circuit. This field is present only if the type is an ExpressRoute connection.`,
				},
				resource.Attribute{
					Name:        "dpd_timeout_seconds",
					Description: `The dead peer detection timeout of this connection in seconds.`,
				},
				resource.Attribute{
					Name:        "express_route_circuit_id",
					Description: `The ID of the Express Route Circuit (i.e. when ` + "`" + `type` + "`" + ` is ` + "`" + `ExpressRoute` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "peer_virtual_network_gateway_id",
					Description: `The ID of the peer virtual network gateway when a VNet-to-VNet connection (i.e. when ` + "`" + `type` + "`" + ` is ` + "`" + `Vnet2Vnet` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "local_azure_ip_address_enabled",
					Description: `Use private local Azure IP for the connection.`,
				},
				resource.Attribute{
					Name:        "local_network_gateway_id",
					Description: `The ID of the local network gateway when a Site-to-Site connection (i.e. when ` + "`" + `type` + "`" + ` is ` + "`" + `IPsec` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "routing_weight",
					Description: `The routing weight.`,
				},
				resource.Attribute{
					Name:        "shared_key",
					Description: `The shared IPSec key.`,
				},
				resource.Attribute{
					Name:        "enable_bgp",
					Description: `If ` + "`" + `true` + "`" + `, BGP (Border Gateway Protocol) is enabled for this connection.`,
				},
				resource.Attribute{
					Name:        "express_route_gateway_bypass",
					Description: `If ` + "`" + `true` + "`" + `, data packets will bypass ExpressRoute Gateway for data forwarding. This is only valid for ExpressRoute connections.`,
				},
				resource.Attribute{
					Name:        "use_policy_based_traffic_selectors",
					Description: `If ` + "`" + `true` + "`" + `, policy-based traffic selectors are enabled for this connection. Enabling policy-based traffic selectors requires an ` + "`" + `ipsec_policy` + "`" + ` block.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the resource. The ` + "`" + `ipsec_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "dh_group",
					Description: `The DH group used in IKE phase 1 for initial SA. Valid options are ` + "`" + `DHGroup1` + "`" + `, ` + "`" + `DHGroup14` + "`" + `, ` + "`" + `DHGroup2` + "`" + `, ` + "`" + `DHGroup2048` + "`" + `, ` + "`" + `DHGroup24` + "`" + `, ` + "`" + `ECP256` + "`" + `, ` + "`" + `ECP384` + "`" + `, or ` + "`" + `None` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ike_encryption",
					Description: `The IKE encryption algorithm. Valid options are ` + "`" + `AES128` + "`" + `, ` + "`" + `AES192` + "`" + `, ` + "`" + `AES256` + "`" + `, ` + "`" + `DES` + "`" + `, or ` + "`" + `DES3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ike_integrity",
					Description: `The IKE integrity algorithm. Valid options are ` + "`" + `MD5` + "`" + `, ` + "`" + `SHA1` + "`" + `, ` + "`" + `SHA256` + "`" + `, or ` + "`" + `SHA384` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ipsec_encryption",
					Description: `The IPSec encryption algorithm. Valid options are ` + "`" + `AES128` + "`" + `, ` + "`" + `AES192` + "`" + `, ` + "`" + `AES256` + "`" + `, ` + "`" + `DES` + "`" + `, ` + "`" + `DES3` + "`" + `, ` + "`" + `GCMAES128` + "`" + `, ` + "`" + `GCMAES192` + "`" + `, ` + "`" + `GCMAES256` + "`" + `, or ` + "`" + `None` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ipsec_integrity",
					Description: `The IPSec integrity algorithm. Valid options are ` + "`" + `GCMAES128` + "`" + `, ` + "`" + `GCMAES192` + "`" + `, ` + "`" + `GCMAES256` + "`" + `, ` + "`" + `MD5` + "`" + `, ` + "`" + `SHA1` + "`" + `, or ` + "`" + `SHA256` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "pfs_group",
					Description: `The DH group used in IKE phase 2 for new child SA. Valid options are ` + "`" + `ECP256` + "`" + `, ` + "`" + `ECP384` + "`" + `, ` + "`" + `PFS1` + "`" + `, ` + "`" + `PFS2` + "`" + `, ` + "`" + `PFS2048` + "`" + `, ` + "`" + `PFS24` + "`" + `, or ` + "`" + `None` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sa_datasize",
					Description: `The IPSec SA payload size in KB. Must be at least ` + "`" + `1024` + "`" + ` KB.`,
				},
				resource.Attribute{
					Name:        "sa_lifetime",
					Description: `The IPSec SA lifetime in seconds. Must be at least ` + "`" + `300` + "`" + ` seconds. The ` + "`" + `traffic_selector_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "local_address_cidrs",
					Description: `List of local CIDRs.`,
				},
				resource.Attribute{
					Name:        "remote_address_cidrs",
					Description: `List of remote CIDRs. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Virtual Network Gateway Connection.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Virtual Network Gateway Connection.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location/region where the connection is located.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of connection. Valid options are ` + "`" + `IPsec` + "`" + ` (Site-to-Site), ` + "`" + `ExpressRoute` + "`" + ` (ExpressRoute), and ` + "`" + `Vnet2Vnet` + "`" + ` (VNet-to-VNet).`,
				},
				resource.Attribute{
					Name:        "virtual_network_gateway_id",
					Description: `The ID of the Virtual Network Gateway in which the connection is created.`,
				},
				resource.Attribute{
					Name:        "authorization_key",
					Description: `The authorization key associated with the Express Route Circuit. This field is present only if the type is an ExpressRoute connection.`,
				},
				resource.Attribute{
					Name:        "dpd_timeout_seconds",
					Description: `The dead peer detection timeout of this connection in seconds.`,
				},
				resource.Attribute{
					Name:        "express_route_circuit_id",
					Description: `The ID of the Express Route Circuit (i.e. when ` + "`" + `type` + "`" + ` is ` + "`" + `ExpressRoute` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "peer_virtual_network_gateway_id",
					Description: `The ID of the peer virtual network gateway when a VNet-to-VNet connection (i.e. when ` + "`" + `type` + "`" + ` is ` + "`" + `Vnet2Vnet` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "local_azure_ip_address_enabled",
					Description: `Use private local Azure IP for the connection.`,
				},
				resource.Attribute{
					Name:        "local_network_gateway_id",
					Description: `The ID of the local network gateway when a Site-to-Site connection (i.e. when ` + "`" + `type` + "`" + ` is ` + "`" + `IPsec` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "routing_weight",
					Description: `The routing weight.`,
				},
				resource.Attribute{
					Name:        "shared_key",
					Description: `The shared IPSec key.`,
				},
				resource.Attribute{
					Name:        "enable_bgp",
					Description: `If ` + "`" + `true` + "`" + `, BGP (Border Gateway Protocol) is enabled for this connection.`,
				},
				resource.Attribute{
					Name:        "express_route_gateway_bypass",
					Description: `If ` + "`" + `true` + "`" + `, data packets will bypass ExpressRoute Gateway for data forwarding. This is only valid for ExpressRoute connections.`,
				},
				resource.Attribute{
					Name:        "use_policy_based_traffic_selectors",
					Description: `If ` + "`" + `true` + "`" + `, policy-based traffic selectors are enabled for this connection. Enabling policy-based traffic selectors requires an ` + "`" + `ipsec_policy` + "`" + ` block.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the resource. The ` + "`" + `ipsec_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "dh_group",
					Description: `The DH group used in IKE phase 1 for initial SA. Valid options are ` + "`" + `DHGroup1` + "`" + `, ` + "`" + `DHGroup14` + "`" + `, ` + "`" + `DHGroup2` + "`" + `, ` + "`" + `DHGroup2048` + "`" + `, ` + "`" + `DHGroup24` + "`" + `, ` + "`" + `ECP256` + "`" + `, ` + "`" + `ECP384` + "`" + `, or ` + "`" + `None` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ike_encryption",
					Description: `The IKE encryption algorithm. Valid options are ` + "`" + `AES128` + "`" + `, ` + "`" + `AES192` + "`" + `, ` + "`" + `AES256` + "`" + `, ` + "`" + `DES` + "`" + `, or ` + "`" + `DES3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ike_integrity",
					Description: `The IKE integrity algorithm. Valid options are ` + "`" + `MD5` + "`" + `, ` + "`" + `SHA1` + "`" + `, ` + "`" + `SHA256` + "`" + `, or ` + "`" + `SHA384` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ipsec_encryption",
					Description: `The IPSec encryption algorithm. Valid options are ` + "`" + `AES128` + "`" + `, ` + "`" + `AES192` + "`" + `, ` + "`" + `AES256` + "`" + `, ` + "`" + `DES` + "`" + `, ` + "`" + `DES3` + "`" + `, ` + "`" + `GCMAES128` + "`" + `, ` + "`" + `GCMAES192` + "`" + `, ` + "`" + `GCMAES256` + "`" + `, or ` + "`" + `None` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ipsec_integrity",
					Description: `The IPSec integrity algorithm. Valid options are ` + "`" + `GCMAES128` + "`" + `, ` + "`" + `GCMAES192` + "`" + `, ` + "`" + `GCMAES256` + "`" + `, ` + "`" + `MD5` + "`" + `, ` + "`" + `SHA1` + "`" + `, or ` + "`" + `SHA256` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "pfs_group",
					Description: `The DH group used in IKE phase 2 for new child SA. Valid options are ` + "`" + `ECP256` + "`" + `, ` + "`" + `ECP384` + "`" + `, ` + "`" + `PFS1` + "`" + `, ` + "`" + `PFS2` + "`" + `, ` + "`" + `PFS2048` + "`" + `, ` + "`" + `PFS24` + "`" + `, or ` + "`" + `None` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sa_datasize",
					Description: `The IPSec SA payload size in KB. Must be at least ` + "`" + `1024` + "`" + ` KB.`,
				},
				resource.Attribute{
					Name:        "sa_lifetime",
					Description: `The IPSec SA lifetime in seconds. Must be at least ` + "`" + `300` + "`" + ` seconds. The ` + "`" + `traffic_selector_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "local_address_cidrs",
					Description: `List of local CIDRs.`,
				},
				resource.Attribute{
					Name:        "remote_address_cidrs",
					Description: `List of remote CIDRs. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Virtual Network Gateway Connection.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_virtual_wan",
			Category:         "Network",
			ShortDescription: `Gets information about an existing Virtual Wan.`,
			Description: `

Use this data source to access information about an existing Virtual Wan.

`,
			Keywords: []string{
				"network",
				"virtual",
				"wan",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_vmware_private_cloud",
			Category:         "VMware (AVS)",
			ShortDescription: `Gets information about an existing Vmware Private Cloud.`,
			Description: `

Use this data source to access information about an existing Vmware Private Cloud.

`,
			Keywords: []string{
				"vmware",
				"avs",
				"private",
				"cloud",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Vmware Private Cloud.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure Region where the Vmware Private Cloud exists.`,
				},
				resource.Attribute{
					Name:        "circuit",
					Description: `A ` + "`" + `circuit` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "internet_connection_enabled",
					Description: `Is the Vmware Private Cluster connected to the internet?`,
				},
				resource.Attribute{
					Name:        "management_cluster",
					Description: `A ` + "`" + `management_cluster` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "network_subnet_cidr",
					Description: `The subnet cidr of the Vmware Private Cloud.`,
				},
				resource.Attribute{
					Name:        "hcx_cloud_manager_endpoint",
					Description: `The endpoint for the HCX Cloud Manager.`,
				},
				resource.Attribute{
					Name:        "nsxt_manager_endpoint",
					Description: `The endpoint for the NSX-T Data Center manager.`,
				},
				resource.Attribute{
					Name:        "vcsa_endpoint",
					Description: `The endpoint for Virtual Center Server Appliance.`,
				},
				resource.Attribute{
					Name:        "sku_name",
					Description: `The Name of the SKU used for this Private Cloud.`,
				},
				resource.Attribute{
					Name:        "nsxt_certificate_thumbprint",
					Description: `The thumbprint of the NSX-T Manager SSL certificate.`,
				},
				resource.Attribute{
					Name:        "vcenter_certificate_thumbprint",
					Description: `The thumbprint of the vCenter Server SSL certificate.`,
				},
				resource.Attribute{
					Name:        "management_subnet_cidr",
					Description: `The network used to access vCenter Server and NSX-T Manager.`,
				},
				resource.Attribute{
					Name:        "provisioning_subnet_cidr",
					Description: `The network which isused for virtual machine cold migration, cloning, and snapshot migration.`,
				},
				resource.Attribute{
					Name:        "vmotion_subnet_cidr",
					Description: `The network which is used for live migration of virtual machines.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the Vmware Private Cloud. --- A ` + "`" + `circuit` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "express_route_id",
					Description: `The ID of the ExpressRoute Circuit.`,
				},
				resource.Attribute{
					Name:        "express_route_private_peering_id",
					Description: `The ID of the ExpressRoute Circuit private peering.`,
				},
				resource.Attribute{
					Name:        "primary_subnet_cidr",
					Description: `The CIDR of the primary subnet.`,
				},
				resource.Attribute{
					Name:        "secondary_subnet_cidr",
					Description: `The CIDR of the secondary subnet. --- A ` + "`" + `management_cluster` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the management cluster.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the management cluster.`,
				},
				resource.Attribute{
					Name:        "hosts",
					Description: `The list of the hosts in the management cluster. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Vmware Private Cloud.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_web_application_firewall_policy",
			Category:         "Network",
			ShortDescription: `Gets information about an existing Web Application Firewall Policy.`,
			Description: `

Use this data source to access information about an existing Web Application Firewall Policy.

`,
			Icon: "Networking/10362-icon-service-Web-Application-Firewall-Policies(WAF).svg",
			Keywords: []string{
				"network",
				"web",
				"application",
				"firewall",
				"policy",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Web Application Firewall Policy. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the Web Application Firewall Policy.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"azurerm_advisor_recommendations":                   0,
		"azurerm_api_management":                            1,
		"azurerm_api_management_api":                        2,
		"azurerm_api_management_api_version_set":            3,
		"azurerm_api_management_group":                      4,
		"azurerm_api_management_product":                    5,
		"azurerm_api_management_user":                       6,
		"azurerm_app_configuration":                         7,
		"azurerm_app_service":                               8,
		"azurerm_app_service_certificate":                   9,
		"azurerm_app_service_certificate_order":             10,
		"azurerm_app_service_environment":                   11,
		"azurerm_app_service_plan":                          12,
		"azurerm_application_gateway":                       13,
		"azurerm_application_insights":                      14,
		"azurerm_application_security_group":                15,
		"azurerm_attestation":                               16,
		"azurerm_automation_account":                        17,
		"azurerm_automation_variable_bool":                  18,
		"azurerm_automation_variable_datetime":              19,
		"azurerm_automation_variable_int":                   20,
		"azurerm_automation_variable_string":                21,
		"azurerm_availability_set":                          22,
		"azurerm_backup_policy_vm":                          23,
		"azurerm_batch_account":                             24,
		"azurerm_batch_certificate":                         25,
		"azurerm_batch_pool":                                26,
		"azurerm_billing_enrollment_account_scope":          27,
		"azurerm_billing_mca_account_scope":                 28,
		"azurerm_blueprint_definition":                      29,
		"azurerm_blueprint_published_version":               30,
		"azurerm_cdn_profile":                               31,
		"azurerm_client_config":                             32,
		"azurerm_cognitive_account":                         33,
		"azurerm_container_registry":                        34,
		"azurerm_cosmosdb_account":                          35,
		"azurerm_data_factory":                              36,
		"azurerm_data_lake_store":                           37,
		"azurerm_data_share":                                38,
		"azurerm_data_share_account":                        39,
		"azurerm_data_share_dataset_blob_storage":           40,
		"azurerm_data_share_dataset_data_lake_gen1":         41,
		"azurerm_data_share_dataset_data_lake_gen2":         42,
		"azurerm_data_share_dataset_kusto_cluster":          43,
		"azurerm_data_share_dataset_kusto_database":         44,
		"azurerm_database_migration_project":                45,
		"azurerm_database_migration_service":                46,
		"azurerm_databricks_workspace":                      47,
		"azurerm_dedicated_host":                            48,
		"azurerm_dedicated_host_group":                      49,
		"azurerm_dev_test_lab":                              50,
		"azurerm_dev_test_virtual_network":                  51,
		"azurerm_digital_twins_instance":                    52,
		"azurerm_disk_access":                               53,
		"azurerm_disk_encryption_set":                       54,
		"azurerm_dns_zone":                                  55,
		"azurerm_eventgrid_domain_topic":                    56,
		"azurerm_eventgrid_topic":                           57,
		"azurerm_eventhub":                                  58,
		"azurerm_eventhub_authorization_rule":               59,
		"azurerm_eventhub_consumer_group":                   60,
		"azurerm_eventhub_namespace":                        61,
		"azurerm_eventhub_namespace_authorization_rule":     62,
		"azurerm_express_route_circuit":                     63,
		"azurerm_firewall":                                  64,
		"azurerm_firewall_policy":                           65,
		"azurerm_function_app":                              66,
		"azurerm_function_app_host_keys":                    67,
		"azurerm_hdinsight_cluster":                         68,
		"azurerm_healthcare_service":                        69,
		"azurerm_image":                                     70,
		"azurerm_images":                                    71,
		"azurerm_iothub":                                    72,
		"azurerm_iothub_dps":                                73,
		"azurerm_iothub_dps_shared_access_policy":           74,
		"azurerm_iothub_shared_access_policy":               75,
		"azurerm_ip_group":                                  76,
		"azurerm_key_vault":                                 77,
		"azurerm_key_vault_access_policy":                   78,
		"azurerm_key_vault_certificate":                     79,
		"azurerm_key_vault_certificate_data":                80,
		"azurerm_key_vault_certificate_issuer":              81,
		"azurerm_key_vault_key":                             82,
		"azurerm_key_vault_secret":                          83,
		"azurerm_kubernetes_cluster":                        84,
		"azurerm_kubernetes_cluster_node_pool":              85,
		"azurerm_kubernetes_service_versions":               86,
		"azurerm_kusto_cluster":                             87,
		"azurerm_lb":                                        88,
		"azurerm_lb_backend_address_pool":                   89,
		"azurerm_lb_rule":                                   90,
		"azurerm_log_analytics_workspace":                   91,
		"azurerm_logic_app_integration_account":             92,
		"azurerm_logic_app_workflow":                        93,
		"azurerm_machine_learning_workspace":                94,
		"azurerm_maintenance_configuration":                 95,
		"azurerm_managed_application_definition":            96,
		"azurerm_managed_disk":                              97,
		"azurerm_management_group":                          98,
		"azurerm_maps_account":                              99,
		"azurerm_mariadb_server":                            100,
		"azurerm_monitor_action_group":                      101,
		"azurerm_monitor_diagnostic_categories":             102,
		"azurerm_monitor_log_profile":                       103,
		"azurerm_monitor_scheduled_query_rules_alert":       104,
		"azurerm_monitor_scheduled_query_rules_log":         105,
		"azurerm_mssql_database":                            106,
		"azurerm_mssql_elasticpool":                         107,
		"azurerm_mssql_server":                              108,
		"azurerm_mysql_server":                              109,
		"azurerm_nat_gateway":                               110,
		"azurerm_netapp_account":                            111,
		"azurerm_netapp_pool":                               112,
		"azurerm_netapp_snapshot":                           113,
		"azurerm_netapp_volume":                             114,
		"azurerm_network_ddos_protection_plan":              115,
		"azurerm_network_interface":                         116,
		"azurerm_network_security_group":                    117,
		"azurerm_network_service_tags":                      118,
		"azurerm_network_watcher":                           119,
		"azurerm_notification_hub":                          120,
		"azurerm_notification_hub_namespace":                121,
		"azurerm_platform_image":                            122,
		"azurerm_policy_definition":                         123,
		"azurerm_policy_set_definition":                     124,
		"azurerm_postgresql_server":                         125,
		"azurerm_private_dns_zone":                          126,
		"azurerm_private_endpoint_connection":               127,
		"azurerm_private_link_service":                      128,
		"azurerm_private_link_service_endpoint_connections": 129,
		"azurerm_proximity_placement_group":                 130,
		"azurerm_public_ip":                                 131,
		"azurerm_public_ip_prefix":                          132,
		"azurerm_public_ips":                                133,
		"azurerm_recovery_services_vault":                   134,
		"azurerm_redis_cache":                               135,
		"azurerm_resource_group":                            136,
		"azurerm_resources":                                 137,
		"azurerm_role_definition":                           138,
		"azurerm_route_filter":                              139,
		"azurerm_route_table":                               140,
		"azurerm_search_service":                            141,
		"azurerm_sentinel_alert_rule":                       142,
		"azurerm_sentinel_alert_rule_template":              143,
		"azurerm_servicebus_namespace":                      144,
		"azurerm_servicebus_namespace_authorization_rule":   145,
		"azurerm_servicebus_queue":                          146,
		"azurerm_servicebus_queue_authorization_rule":       147,
		"azurerm_servicebus_subscription":                   148,
		"azurerm_servicebus_topic":                          149,
		"azurerm_servicebus_topic_authorization_rule":       150,
		"azurerm_shared_image":                              151,
		"azurerm_shared_image_gallery":                      152,
		"azurerm_shared_image_version":                      153,
		"azurerm_shared_image_versions":                     154,
		"azurerm_signalr_service":                           155,
		"azurerm_snapshot":                                  156,
		"azurerm_spring_cloud_app":                          157,
		"azurerm_spring_cloud_service":                      158,
		"azurerm_sql_database":                              159,
		"azurerm_sql_server":                                160,
		"azurerm_ssh_public_key":                            161,
		"azurerm_storage_account":                           162,
		"azurerm_storage_account_blob_container_sas":        163,
		"azurerm_storage_account_sas":                       164,
		"azurerm_storage_container":                         165,
		"azurerm_storage_encryption_scope":                  166,
		"azurerm_storage_management_policy":                 167,
		"azurerm_storage_sync":                              168,
		"azurerm_storage_sync_group":                        169,
		"azurerm_stream_analytics_job":                      170,
		"azurerm_subnet":                                    171,
		"azurerm_subscription":                              172,
		"azurerm_subscriptions":                             173,
		"azurerm_synapse_workspace":                         174,
		"azurerm_template_spec_version":                     175,
		"azurerm_traffic_manager_geographical_location":     176,
		"azurerm_traffic_manager_profile":                   177,
		"azurerm_user_assigned_identity":                    178,
		"azurerm_virtual_hub":                               179,
		"azurerm_virtual_machine":                           180,
		"azurerm_virtual_machine_scale_set":                 181,
		"azurerm_virtual_network":                           182,
		"azurerm_virtual_network_gateway":                   183,
		"azurerm_virtual_network_gateway_connection":        184,
		"azurerm_virtual_wan":                               185,
		"azurerm_vmware_private_cloud":                      186,
		"azurerm_web_application_firewall_policy":           187,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
