package azurerm

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "azurerm_api_management",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing API Management Service.`,
			Description: `

Use this data source to access information about an existing API Management Service.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the API Management service.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The Name of the Resource Group in which the API Management Service exists. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the API Management Service.`,
				},
				resource.Attribute{
					Name:        "additional_location",
					Description: `One or more ` + "`" + `additional_location` + "`" + ` blocks as defined below`,
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
					Name:        "public_ip_addresses",
					Description: `The Public IP addresses of the API Management Service.`,
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
					Description: `Public Static Load Balanced IP addresses of the API Management service in the additional location. Available only for Basic, Standard and Premium SKU. --- A ` + "`" + `hostname_configuration` + "`" + ` block exports the following:`,
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
					Description: `Specifies the number of units associated with this API Management service.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the API Management Service.`,
				},
				resource.Attribute{
					Name:        "additional_location",
					Description: `One or more ` + "`" + `additional_location` + "`" + ` blocks as defined below`,
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
					Name:        "public_ip_addresses",
					Description: `The Public IP addresses of the API Management Service.`,
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
					Description: `Public Static Load Balanced IP addresses of the API Management service in the additional location. Available only for Basic, Standard and Premium SKU. --- A ` + "`" + `hostname_configuration` + "`" + ` block exports the following:`,
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
					Description: `Specifies the number of units associated with this API Management service.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_api_management_api",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing API Management API.`,
			Description: `

Use this data source to access information about an existing API Management API.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the API Management API.`,
				},
				resource.Attribute{
					Name:        "api_management_name",
					Description: `(Required) The name of the API Management Service in which the API Management API exists.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The Name of the Resource Group in which the API Management Service exists.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `(Required) The Revision of the API Management API. ## Attributes Reference`,
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
					Description: `The name of endpoint (port) to import from WSDL.`,
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
					Description: `The name of endpoint (port) to import from WSDL.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_api_management_group",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing API Management Group.`,
			Description: `

Use this data source to access information about an existing API Management Group.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "api_management_name",
					Description: `(Required) The Name of the API Management Service in which this Group exists.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Name of the API Management Group.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The Name of the Resource Group in which the API Management Service exists. ## Attributes Reference`,
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
					Description: `The type of this API Management Group, such as ` + "`" + `custom` + "`" + ` or ` + "`" + `external` + "`" + `.`,
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
					Description: `The type of this API Management Group, such as ` + "`" + `custom` + "`" + ` or ` + "`" + `external` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_api_management_product",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing API Management Product.`,
			Description: `

Use this data source to access information about an existing API Management Product.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "api_management_name",
					Description: `(Required) The Name of the API Management Service in which this Product exists.`,
				},
				resource.Attribute{
					Name:        "product_id",
					Description: `(Required) The Identifier for the API Management Product.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The Name of the Resource Group in which the API Management Service exists. ## Attributes Reference`,
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
					Description: `Any Terms and Conditions for this Product, which must be accepted by Developers before they can begin the Subscription process.`,
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
					Description: `Any Terms and Conditions for this Product, which must be accepted by Developers before they can begin the Subscription process.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_api_management_user",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing API Management User.`,
			Description: `

Use this data source to access information about an existing API Management User.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "api_management_name",
					Description: `(Required) The Name of the API Management Service in which this User exists.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The Name of the Resource Group in which the API Management Service exists.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `(Required) The Identifier for the User. ## Attributes Reference`,
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
					Description: `The current state of this User, for example ` + "`" + `active` + "`" + `, ` + "`" + `blocked` + "`" + ` or ` + "`" + `pending` + "`" + `.`,
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
					Description: `The current state of this User, for example ` + "`" + `active` + "`" + `, ` + "`" + `blocked` + "`" + ` or ` + "`" + `pending` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_app_service",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing App Service.`,
			Description: `

Use this data source to access information about an existing App Service.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the App Service.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The Name of the Resource Group where the App Service exists. ## Attributes Reference`,
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
					Name:        "outbound_ip_addresses",
					Description: `A comma separated list of outbound IP addresses - such as ` + "`" + `52.23.25.3,52.143.43.12` + "`" + ``,
				},
				resource.Attribute{
					Name:        "possible_outbound_ip_addresses",
					Description: `A comma separated list of outbound IP addresses - such as ` + "`" + `52.23.25.3,52.143.43.12,52.143.43.17` + "`" + ` - not all of which are necessarily in use. Superset of ` + "`" + `outbound_ip_addresses` + "`" + `. --- ` + "`" + `connection_string` + "`" + ` supports the following:`,
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
					Description: `Are credentials supported? --- A ` + "`" + `ip_restriction` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The IP Address used for this IP Restriction.`,
				},
				resource.Attribute{
					Name:        "subnet_mask",
					Description: `The Subnet mask used for this IP Restriction. --- ` + "`" + `site_config` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "always_on",
					Description: `Is the app be loaded at all times?`,
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
					Name:        "ip_restriction",
					Description: `One or more ` + "`" + `ip_restriction` + "`" + ` blocks as defined above.`,
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
					Description: `Are WebSockets enabled for this App Service?`,
				},
				resource.Attribute{
					Name:        "virtual_network_name",
					Description: `The name of the Virtual Network which this App Service is attached to.`,
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
					Name:        "outbound_ip_addresses",
					Description: `A comma separated list of outbound IP addresses - such as ` + "`" + `52.23.25.3,52.143.43.12` + "`" + ``,
				},
				resource.Attribute{
					Name:        "possible_outbound_ip_addresses",
					Description: `A comma separated list of outbound IP addresses - such as ` + "`" + `52.23.25.3,52.143.43.12,52.143.43.17` + "`" + ` - not all of which are necessarily in use. Superset of ` + "`" + `outbound_ip_addresses` + "`" + `. --- ` + "`" + `connection_string` + "`" + ` supports the following:`,
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
					Description: `Are credentials supported? --- A ` + "`" + `ip_restriction` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The IP Address used for this IP Restriction.`,
				},
				resource.Attribute{
					Name:        "subnet_mask",
					Description: `The Subnet mask used for this IP Restriction. --- ` + "`" + `site_config` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "always_on",
					Description: `Is the app be loaded at all times?`,
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
					Name:        "ip_restriction",
					Description: `One or more ` + "`" + `ip_restriction` + "`" + ` blocks as defined above.`,
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
					Description: `Are WebSockets enabled for this App Service?`,
				},
				resource.Attribute{
					Name:        "virtual_network_name",
					Description: `The name of the Virtual Network which this App Service is attached to.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_app_service_plan",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing App Service Plan.`,
			Description: `

Use this data source to access information about an existing App Service Plan (formerly known as a ` + "`" + `Server Farm` + "`" + `).

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the App Service Plan.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The Name of the Resource Group where the App Service Plan exists. ## Attributes Reference`,
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
					Name:        "properties",
					Description: `A ` + "`" + `properties` + "`" + ` block as documented below.`,
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
					Description: `Specifies the number of workers associated with this App Service Plan. A ` + "`" + `properties` + "`" + ` block supports the following:`,
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
					Name:        "properties",
					Description: `A ` + "`" + `properties` + "`" + ` block as documented below.`,
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
					Description: `Specifies the number of workers associated with this App Service Plan. A ` + "`" + `properties` + "`" + ` block supports the following:`,
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
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_application_insights",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing Application Insights component.`,
			Description: `

Use this data source to access information about an existing Application Insights component.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the Application Insights component.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) Specifies the name of the resource group the Application Insights component is located in. ## Attributes Reference`,
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
					Name:        "location",
					Description: `The Azure location where the component exists.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags applied to the component.`,
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
					Name:        "location",
					Description: `The Azure location where the component exists.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags applied to the component.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_application_security_group",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing Application Security Group.`,
			Description: `

Use this data source to access information about an existing Application Security Group.

`,
			Keywords: []string{},
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
					Description: `A mapping of tags assigned to the resource.`,
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
					Description: `A mapping of tags assigned to the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_automation_variable_bool",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing Automation Bool Variable`,
			Description: `

Use this data source to access information about an existing Automation Bool Variable.


`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Automation Variable.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The Name of the Resource Group where the automation account exists.`,
				},
				resource.Attribute{
					Name:        "automation_account_name",
					Description: `(Required) The name of the automation account in which the Automation Variable exists. ## Attributes Reference The following attributes are exported:`,
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
					Description: `The value of the Automation Variable as a ` + "`" + `boolean` + "`" + `.`,
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
					Description: `The value of the Automation Variable as a ` + "`" + `boolean` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_automation_variable_datetime",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing Automation Datetime Variable`,
			Description: `

Use this data source to access information about an existing Automation Datetime Variable.


`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Automation Variable.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The Name of the Resource Group where the automation account exists.`,
				},
				resource.Attribute{
					Name:        "automation_account_name",
					Description: `(Required) The name of the automation account in which the Automation Variable exists. ## Attributes Reference The following attributes are exported:`,
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
					Description: `The value of the Automation Variable in the [RFC3339 Section 5.6 Internet Date/Time Format](https://tools.ietf.org/html/rfc3339#section-5.6).`,
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
					Description: `The value of the Automation Variable in the [RFC3339 Section 5.6 Internet Date/Time Format](https://tools.ietf.org/html/rfc3339#section-5.6).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_automation_variable_int",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing Automation Int Variable`,
			Description: `

Use this data source to access information about an existing Automation Int Variable.


`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Automation Variable.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The Name of the Resource Group where the automation account exists.`,
				},
				resource.Attribute{
					Name:        "automation_account_name",
					Description: `(Required) The name of the automation account in which the Automation Variable exists. ## Attributes Reference The following attributes are exported:`,
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
					Description: `The value of the Automation Variable as a ` + "`" + `integer` + "`" + `.`,
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
					Description: `The value of the Automation Variable as a ` + "`" + `integer` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_automation_variable_string",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing Automation String Variable`,
			Description: `

Use this data source to access information about an existing Automation String Variable.


`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Automation Variable.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The Name of the Resource Group where the automation account exists.`,
				},
				resource.Attribute{
					Name:        "automation_account_name",
					Description: `(Required) The name of the automation account in which the Automation Variable exists. ## Attributes Reference The following attributes are exported:`,
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
					Description: `The value of the Automation Variable as a ` + "`" + `string` + "`" + `.`,
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
					Description: `The value of the Automation Variable as a ` + "`" + `string` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_availability_set",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing Availability Set.`,
			Description: `

Use this data source to access information about an existing Availability Set.

`,
			Keywords: []string{},
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
					Description: `A mapping of tags assigned to the resource.`,
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
					Description: `A mapping of tags assigned to the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_azuread_application",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing Application within Azure Active Directory.`,
			Description: `

Use this data source to access information about an existing Application within Azure Active Directory.

~> **NOTE:** The Azure Active Directory resources have been split out into [a new AzureAD Provider](http://terraform.io/docs/providers/azuread/index.html) - as such the AzureAD resources within the AzureRM Provider are deprecated and will be removed in the next major version (2.0). Information on how to migrate from the existing resources to the new AzureAD Provider [can be found here](../guides/migrating-to-azuread.html).

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both ` + "`" + `Read and write all applications` + "`" + ` and ` + "`" + `Sign in and read user profile` + "`" + ` within the ` + "`" + `Windows Azure Active Directory` + "`" + ` API.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "object_id",
					Description: `(Optional) Specifies the Object ID of the Application within Azure Active Directory.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Specifies the name of the Application within Azure Active Directory. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `the Object ID of the Azure Active Directory Application.`,
				},
				resource.Attribute{
					Name:        "application_id",
					Description: `the Application ID of the Azure Active Directory Application.`,
				},
				resource.Attribute{
					Name:        "available_to_other_tenants",
					Description: `Is this Azure AD Application available to other tenants?`,
				},
				resource.Attribute{
					Name:        "identifier_uris",
					Description: `A list of user-defined URI(s) that uniquely identify a Web application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant.`,
				},
				resource.Attribute{
					Name:        "oauth2_allow_implicit_flow",
					Description: `Does this Azure AD Application allow OAuth2.0 implicit flow tokens?`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `the Object ID of the Azure Active Directory Application.`,
				},
				resource.Attribute{
					Name:        "reply_urls",
					Description: `A list of URLs that user tokens are sent to for sign in, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `the Object ID of the Azure Active Directory Application.`,
				},
				resource.Attribute{
					Name:        "application_id",
					Description: `the Application ID of the Azure Active Directory Application.`,
				},
				resource.Attribute{
					Name:        "available_to_other_tenants",
					Description: `Is this Azure AD Application available to other tenants?`,
				},
				resource.Attribute{
					Name:        "identifier_uris",
					Description: `A list of user-defined URI(s) that uniquely identify a Web application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant.`,
				},
				resource.Attribute{
					Name:        "oauth2_allow_implicit_flow",
					Description: `Does this Azure AD Application allow OAuth2.0 implicit flow tokens?`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `the Object ID of the Azure Active Directory Application.`,
				},
				resource.Attribute{
					Name:        "reply_urls",
					Description: `A list of URLs that user tokens are sent to for sign in, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_azuread_service_principal",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing Service Principal associated with an Application within Azure Active Directory.`,
			Description: `

Gets information about an existing Service Principal associated with an Application within Azure Active Directory.

~> **NOTE:** The Azure Active Directory resources have been split out into [a new AzureAD Provider](http://terraform.io/docs/providers/azuread/index.html) - as such the AzureAD resources within the AzureRM Provider are deprecated and will be removed in the next major version (2.0). Information on how to migrate from the existing resources to the new AzureAD Provider [can be found here](../guides/migrating-to-azuread.html).

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both ` + "`" + `Read and write all applications` + "`" + ` and ` + "`" + `Sign in and read user profile` + "`" + ` within the ` + "`" + `Windows Azure Active Directory` + "`" + ` API.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "application_id",
					Description: `(Optional) The ID of the Azure AD Application for which to create a Service Principal.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `(Optional) The ID of the Azure AD Service Principal.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name of the Azure AD Application associated with this Service Principal. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Object ID for the Service Principal.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Object ID for the Service Principal.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_batch_account",
			Category:         "Data Sources",
			ShortDescription: `Get information about an existing Batch Account`,
			Description: `

Use this data source to access information about an existing Batch Account.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Batch account.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The Name of the Resource Group where this Batch account exists. ## Attributes Reference The following attributes are exported:`,
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
					Description: `The HTTPS URL of the Azure KeyVault reference. ---`,
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
					Description: `The HTTPS URL of the Azure KeyVault reference. ---`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_batch_certificate",
			Category:         "Data Sources",
			ShortDescription: `Get information about an existing certificate in a Batch Account`,
			Description: `

Use this data source to access information about an existing certificate in a Batch Account.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Batch certificate.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `(Required) The name of the Batch account.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The Name of the Resource Group where this Batch account exists. ## Attributes Reference The following attributes are exported:`,
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
					Description: `The algorithm of the certificate thumbprint.`,
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
					Description: `The algorithm of the certificate thumbprint.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_batch_pool",
			Category:         "Data Sources",
			ShortDescription: `Get information about an existing Azure Batch pool.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
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
					Description: `(Optional) One or more ` + "`" + `resource_file` + "`" + ` blocks that describe the files to be downloaded to a compute node. --- A ` + "`" + `user_identity` + "`" + ` block exports the following:`,
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
					Description: `The password to log into the registry server.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_builtin_role_definition",
			Category:         "Data Sources",
			ShortDescription: `Get information about an existing built-in Role Definition.`,
			Description: `

Use this data source to access information about a built-in Role Definition. To access information about a custom Role Definition, [please see the ` + "`" + `azurerm_role_definition` + "`" + ` data source](role_definition.html) instead.

~> **NOTE:** The this datasource has been deprecated in favour of ` + "`" + `azurerm_role_definition` + "`" + ` that now can look up role definitions by name. As such this data source will be removed in version 2.0 of the AzureRM Provider.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the built-in Role Definition. Possible values are: ` + "`" + `Contributor` + "`" + `, ` + "`" + `Owner` + "`" + `, ` + "`" + `Reader` + "`" + ` and ` + "`" + `VirtualMachineContributor` + "`" + `. ## Attributes Reference`,
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
					Name:        "data_actions",
					Description: `a list of data actions supported by this role`,
				},
				resource.Attribute{
					Name:        "not_actions",
					Description: `a list of actions which are denied by this role`,
				},
				resource.Attribute{
					Name:        "not_data_actions",
					Description: `a list of data actions which are denied by this role`,
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
					Name:        "data_actions",
					Description: `a list of data actions supported by this role`,
				},
				resource.Attribute{
					Name:        "not_actions",
					Description: `a list of actions which are denied by this role`,
				},
				resource.Attribute{
					Name:        "not_data_actions",
					Description: `a list of data actions which are denied by this role`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_cdn_profile",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing CDN Profile`,
			Description: `

Use this data source to access information about an existing CDN Profile.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the CDN Profile.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The name of the resource group in which the CDN Profile exists. ## Attributes Reference`,
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
					Description: `A mapping of tags assigned to the resource.`,
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
					Description: `A mapping of tags assigned to the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_client_config",
			Category:         "Data Sources",
			ShortDescription: `Gets information about the configuration of the azurerm provider.`,
			Description: `

Use this data source to access the configuration of the AzureRM provider.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_container_registry",
			Category:         "Data Sources",
			ShortDescription: `Get information about an existing Container Registry`,
			Description: `

Use this data source to access information about an existing Container Registry.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Container Registry.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The Name of the Resource Group where this Container Registry exists. ## Attributes Reference The following attributes are exported:`,
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
					Description: `A map of tags assigned to the Container Registry.`,
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
					Description: `A map of tags assigned to the Container Registry.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_cosmosdb_account",
			Category:         "Data Sources",
			ShortDescription: `Gets information about the specified CosmosDB (formally DocumentDB) Account.`,
			Description: `

Use this data source to access information about an existing CosmosDB (formally DocumentDB) Account.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the CosmosDB Account.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) Specifies the name of the resource group in which the CosmosDB Account resides. ## Attributes Reference The following attributes are exported:`,
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
					Name:        "ip_range_filter",
					Description: `The current IP Filter for this CosmosDB account`,
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
					Name:        "primary_master_key",
					Description: `The Primary master key for the CosmosDB Account.`,
				},
				resource.Attribute{
					Name:        "secondary_master_key",
					Description: `The Secondary master key for the CosmosDB Account.`,
				},
				resource.Attribute{
					Name:        "primary_readonly_master_key",
					Description: `The Primary read-only master Key for the CosmosDB Account.`,
				},
				resource.Attribute{
					Name:        "secondary_readonly_master_key",
					Description: `The Secondary read-only master key for the CosmosDB Account.`,
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
					Name:        "ip_range_filter",
					Description: `The current IP Filter for this CosmosDB account`,
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
					Name:        "primary_master_key",
					Description: `The Primary master key for the CosmosDB Account.`,
				},
				resource.Attribute{
					Name:        "secondary_master_key",
					Description: `The Secondary master key for the CosmosDB Account.`,
				},
				resource.Attribute{
					Name:        "primary_readonly_master_key",
					Description: `The Primary read-only master Key for the CosmosDB Account.`,
				},
				resource.Attribute{
					Name:        "secondary_readonly_master_key",
					Description: `The Secondary read-only master key for the CosmosDB Account.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_data_lake_store",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing Data Lake Store`,
			Description: `

Use this data source to access information about an existing Data Lake Store.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Data Lake Store.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The Name of the Resource Group where the Data Lake Store exists. ## Attributes Reference`,
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
					Description: `A mapping of tags to assign to the Data Lake Store.`,
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
					Description: `A mapping of tags to assign to the Data Lake Store.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_dev_test_lab",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing Dev Test Lab.`,
			Description: `

Use this data source to access information about an existing Dev Test Lab.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Dev Test Lab.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The Name of the Resource Group where the Dev Test Lab exists. ## Attributes Reference`,
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
					Description: `The unique immutable identifier of the Dev Test Lab.`,
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
					Description: `The unique immutable identifier of the Dev Test Lab.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_dns_zone",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing DNS Zone.`,
			Description: `

Use this data source to access information about an existing DNS Zone.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the DNS Zone.`,
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
					Name:        "registration_virtual_network_ids",
					Description: `A list of Virtual Network ID's that register hostnames in this DNS zone.`,
				},
				resource.Attribute{
					Name:        "resolution_virtual_network_ids",
					Description: `A list of Virtual Network ID's that resolve records in this DNS zone.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the EventHub Namespace.`,
				},
				resource.Attribute{
					Name:        "zone_type",
					Description: `(`,
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
					Name:        "registration_virtual_network_ids",
					Description: `A list of Virtual Network ID's that register hostnames in this DNS zone.`,
				},
				resource.Attribute{
					Name:        "resolution_virtual_network_ids",
					Description: `A list of Virtual Network ID's that resolve records in this DNS zone.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the EventHub Namespace.`,
				},
				resource.Attribute{
					Name:        "zone_type",
					Description: `(`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_eventhub_namespace",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing EventHub Namespace.`,
			Description: `

Use this data source to access information about an existing EventHub Namespace.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the EventHub Namespace.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The Name of the Resource Group where the EventHub Namespace exists. ## Attributes Reference`,
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
					Name:        "tags",
					Description: `A mapping of tags to assign to the EventHub Namespace. The following attributes are exported only if there is an authorization rule named ` + "`" + `RootManageSharedAccessKey` + "`" + ` which is created automatically by Azure.`,
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
					Description: `The secondary access key for the authorization rule ` + "`" + `RootManageSharedAccessKey` + "`" + `.`,
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
					Name:        "tags",
					Description: `A mapping of tags to assign to the EventHub Namespace. The following attributes are exported only if there is an authorization rule named ` + "`" + `RootManageSharedAccessKey` + "`" + ` which is created automatically by Azure.`,
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
					Description: `The secondary access key for the authorization rule ` + "`" + `RootManageSharedAccessKey` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_express_route_circuit",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing ExpressRoute circuit.`,
			Description: `

Use this data source to access information about an existing ExpressRoute circuit.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the ExpressRoute circuit.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The Name of the Resource Group where the ExpressRoute circuit exists. ## Attributes Reference`,
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
					Description: `The service tier. Possible values are ` + "`" + `Standard` + "`" + ` or ` + "`" + `Premium` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "family",
					Description: `The billing mode for bandwidth. Possible values are ` + "`" + `MeteredData` + "`" + ` or ` + "`" + `UnlimitedData` + "`" + `.`,
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
					Description: `The service tier. Possible values are ` + "`" + `Standard` + "`" + ` or ` + "`" + `Premium` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "family",
					Description: `The billing mode for bandwidth. Possible values are ` + "`" + `MeteredData` + "`" + ` or ` + "`" + `UnlimitedData` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_firewall",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing Azure Firewall.`,
			Description: `

Use this data source to access information about an existing Azure Firewall.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Azure Firewall.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The name of the Resource Group in which the Azure Firewall exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Resource ID of the Azure Firewall.`,
				},
				resource.Attribute{
					Name:        "ip_configuration",
					Description: `A ` + "`" + `ip_configuration` + "`" + ` block as defined below. --- A ` + "`" + `ip_configuration` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The Resource ID of the subnet where the Azure Firewall is deployed.`,
				},
				resource.Attribute{
					Name:        "private_ip_address",
					Description: `The private IP address of the Azure Firewall.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Resource ID of the Azure Firewall.`,
				},
				resource.Attribute{
					Name:        "ip_configuration",
					Description: `A ` + "`" + `ip_configuration` + "`" + ` block as defined below. --- A ` + "`" + `ip_configuration` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The Resource ID of the subnet where the Azure Firewall is deployed.`,
				},
				resource.Attribute{
					Name:        "private_ip_address",
					Description: `The private IP address of the Azure Firewall.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_hdinsight_cluster",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing HDInsight Cluster.`,
			Description: `

Use this data source to access information about an existing HDInsight Cluster.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of this HDInsight Cluster.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) Specifies the name of the Resource Group in which this HDInsight Cluster exists. ## Attributes Reference`,
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
					Description: `The password used for the Ambari Portal.`,
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
					Description: `The password used for the Ambari Portal.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_image",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing Image`,
			Description: `

Use this data source to access information about an existing Image.

`,
			Keywords: []string{},
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
					Description: `(Required) The Name of the Resource Group where this Image exists. ## Attributes Reference`,
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
					Description: `the size of this Data Disk in GB.`,
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
					Description: `the size of this Data Disk in GB.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_key_vault",
			Category:         "Data Sources",
			ShortDescription: `Gets information about a Key Vault.`,
			Description: `

Use this data source to access information about an existing Key Vault.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the Key Vault.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The name of the Resource Group in which the Key Vault exists. ## Attributes Reference The following attributes are exported:`,
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
					Name:        "sku",
					Description: `A ` + "`" + `sku` + "`" + ` block as described below.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `The Azure Active Directory Tenant ID used for authenticating requests to the Key Vault.`,
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
					Name:        "tags",
					Description: `A mapping of tags assigned to the Key Vault. A ` + "`" + `sku` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the SKU used for this Key Vault. ` + "`" + `access_policy` + "`" + ` supports the following:`,
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
					Description: `A list of storage permissions applicable to this Access Policy.`,
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
					Name:        "sku",
					Description: `A ` + "`" + `sku` + "`" + ` block as described below.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `The Azure Active Directory Tenant ID used for authenticating requests to the Key Vault.`,
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
					Name:        "tags",
					Description: `A mapping of tags assigned to the Key Vault. A ` + "`" + `sku` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the SKU used for this Key Vault. ` + "`" + `access_policy` + "`" + ` supports the following:`,
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
					Description: `A list of storage permissions applicable to this Access Policy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_key_vault_access_policy",
			Category:         "Data Sources",
			ShortDescription: `Get information about the templated Access Policies for Key Vault.`,
			Description: `

Use this data source to access information about the permissions from the Management Key Vault Templates.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the Management Template. Possible values are: ` + "`" + `Key Management` + "`" + `, ` + "`" + `Secret Management` + "`" + `, ` + "`" + `Certificate Management` + "`" + `, ` + "`" + `Key & Secret Management` + "`" + `, ` + "`" + `Key & Certificate Management` + "`" + `, ` + "`" + `Secret & Certificate Management` + "`" + `, ` + "`" + `Key, Secret, & Certificate Management` + "`" + ` ## Attributes Reference`,
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
					Description: `the certificate permissions for the access policy`,
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
					Description: `the certificate permissions for the access policy`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_key_vault_key",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing Key Vault Key.`,
			Description: `

Use this data source to access information about an existing Key Vault Key.

~> **Note:** All arguments including the secret value will be stored in the raw state as plain-text.
[Read more about sensitive data in state](/docs/state/sensitive-data.html).

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the Key Vault Key.`,
				},
				resource.Attribute{
					Name:        "vault_uri",
					Description: `(Required) Specifies the ID of the Key Vault Key Vault instance where the Key resides, available on the ` + "`" + `azurerm_key_vault` + "`" + ` Data Source / Resource. ## Attributes Reference The following attributes are exported:`,
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
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_key_vault_secret",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing Key Vault Secret.`,
			Description: `

Use this data source to access information about an existing Key Vault Secret.

~> **Note:** All arguments including the secret value will be stored in the raw state as plain-text.
[Read more about sensitive data in state](/docs/state/sensitive-data.html).

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the Key Vault Secret.`,
				},
				resource.Attribute{
					Name:        "key_vault_id",
					Description: `(Required) Specifies the ID of the Key Vault Key Vault instance where the Secret resides, available on the ` + "`" + `azurerm_key_vault` + "`" + ` Data Source / Resource. ## Attributes Reference The following attributes are exported:`,
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
					Description: `Any tags assigned to this resource.`,
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
					Description: `Any tags assigned to this resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_kubernetes_cluster",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing Managed Kubernetes Cluster (AKS)`,
			Description: `

Use this data source to access information about an existing Managed Kubernetes Cluster (AKS).

~> **Note:** All arguments including the client secret will be stored in the raw state as plain-text.
[Read more about sensitive data in state](/docs/state/sensitive-data.html).

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the managed Kubernetes Cluster.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The name of the Resource Group in which the managed Kubernetes Cluster exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Kubernetes Managed Cluster.`,
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
					Name:        "location",
					Description: `The Azure Region in which the managed Kubernetes Cluster exists.`,
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
					Name:        "tags",
					Description: `A mapping of tags assigned to this resource. --- A ` + "`" + `addon_profile` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "http_application_routing",
					Description: `A ` + "`" + `http_application_routing` + "`" + ` block.`,
				},
				resource.Attribute{
					Name:        "oms_agent",
					Description: `A ` + "`" + `oms_agent` + "`" + ` block. --- A ` + "`" + `agent_pool_profile` + "`" + ` block exports the following:`,
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
					Name:        "vm_size",
					Description: `The size of each VM in the Agent Pool (e.g. ` + "`" + `Standard_F1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "vnet_subnet_id",
					Description: `The ID of the Subnet where the Agents in the Pool are provisioned.`,
				},
				resource.Attribute{
					Name:        "node_taints",
					Description: `The list of Kubernetes taints which are applied to nodes in the agent pool --- A ` + "`" + `azure_active_directory` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "client_app_id",
					Description: `The Client ID of an Azure Active Directory Application.`,
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
					Description: `The ID of the Log Analytics Workspace which the OMS Agent should send data to. --- A ` + "`" + `role_based_access_control` + "`" + ` block exports the following:`,
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
					Description: `The Client ID of the Service Principal used by this Managed Kubernetes Cluster. --- A ` + "`" + `ssh_key` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "key_data",
					Description: `The Public SSH Key used to access the cluster.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Kubernetes Managed Cluster.`,
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
					Name:        "location",
					Description: `The Azure Region in which the managed Kubernetes Cluster exists.`,
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
					Name:        "tags",
					Description: `A mapping of tags assigned to this resource. --- A ` + "`" + `addon_profile` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "http_application_routing",
					Description: `A ` + "`" + `http_application_routing` + "`" + ` block.`,
				},
				resource.Attribute{
					Name:        "oms_agent",
					Description: `A ` + "`" + `oms_agent` + "`" + ` block. --- A ` + "`" + `agent_pool_profile` + "`" + ` block exports the following:`,
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
					Name:        "vm_size",
					Description: `The size of each VM in the Agent Pool (e.g. ` + "`" + `Standard_F1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "vnet_subnet_id",
					Description: `The ID of the Subnet where the Agents in the Pool are provisioned.`,
				},
				resource.Attribute{
					Name:        "node_taints",
					Description: `The list of Kubernetes taints which are applied to nodes in the agent pool --- A ` + "`" + `azure_active_directory` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "client_app_id",
					Description: `The Client ID of an Azure Active Directory Application.`,
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
					Description: `The ID of the Log Analytics Workspace which the OMS Agent should send data to. --- A ` + "`" + `role_based_access_control` + "`" + ` block exports the following:`,
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
					Description: `The Client ID of the Service Principal used by this Managed Kubernetes Cluster. --- A ` + "`" + `ssh_key` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "key_data",
					Description: `The Public SSH Key used to access the cluster.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_kubernetes_service_versions",
			Category:         "Data Sources",
			ShortDescription: `Gets the available versions of Kubernetes supported by Azure Kubernetes Service.`,
			Description: `

Use this data source to retrieve the version of Kubernetes supported by Azure Kubernetes Service.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "location",
					Description: `(Required) Specifies the location in which to query for versions.`,
				},
				resource.Attribute{
					Name:        "version_prefix",
					Description: `(Optional) A prefix filter for the versions of Kubernetes which should be returned; for example ` + "`" + `1.` + "`" + ` will return ` + "`" + `1.9` + "`" + ` to ` + "`" + `1.14` + "`" + `, whereas ` + "`" + `1.12` + "`" + ` will return ` + "`" + `1.12.2` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "versions",
					Description: `The list of all supported versions.`,
				},
				resource.Attribute{
					Name:        "latest_version",
					Description: `The most recent version available.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "versions",
					Description: `The list of all supported versions.`,
				},
				resource.Attribute{
					Name:        "latest_version",
					Description: `The most recent version available.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_lb",
			Category:         "Data Sources",
			ShortDescription: `Get information about an existing Load Balancer`,
			Description: `

Use this data source to access information about an existing Load Balancer

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the Load Balancer.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The name of the Resource Group in which the Load Balancer exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Load Balancer.`,
				},
				resource.Attribute{
					Name:        "frontend_ip_configuration",
					Description: `(Optional) A ` + "`" + `frontend_ip_configuration` + "`" + ` block as documented below.`,
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
					Name:        "public_ip_address_id",
					Description: `The ID of a Public IP Address which is associated with this Load Balancer.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of Availability Zones which the Load Balancer's IP Addresses should be created in.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Load Balancer.`,
				},
				resource.Attribute{
					Name:        "frontend_ip_configuration",
					Description: `(Optional) A ` + "`" + `frontend_ip_configuration` + "`" + ` block as documented below.`,
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
					Name:        "public_ip_address_id",
					Description: `The ID of a Public IP Address which is associated with this Load Balancer.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of Availability Zones which the Load Balancer's IP Addresses should be created in.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_lb_backend_address_pool",
			Category:         "Data Sources",
			ShortDescription: `Get information about an existing Load Balancer Backend Address Pool`,
			Description: `

Use this data source to access information about an existing Load Balancer Backend Address Pool.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the Backend Address Pool.`,
				},
				resource.Attribute{
					Name:        "loadbalancer_id",
					Description: `(Required) The ID of the Load Balancer in which the Backend Address Pool exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Backend Address Pool.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Backend Address Pool.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_log_analytics_workspace",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing Log Analytics (formally Operational Insights) Workspace.`,
			Description: `

Use this data source to access information about an existing Log Analytics (formally Operational Insights) Workspace.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the Log Analytics Workspace.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The name of the resource group in which the Log Analytics workspace is located in. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Azure Resource ID of the Log Analytics Workspace.`,
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
					Name:        "portal_url",
					Description: `The Portal URL for the Log Analytics Workspace.`,
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
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Azure Resource ID of the Log Analytics Workspace.`,
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
					Name:        "portal_url",
					Description: `The Portal URL for the Log Analytics Workspace.`,
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
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_logic_app_workflow",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing Logic App Workflow.`,
			Description: `

Use this data source to access information about an existing Logic App Workflow.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Logic App Workflow.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The name of the Resource Group in which the Logic App Workflow exists. ## Attributes Reference The following attributes are exported:`,
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
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_managed_disk",
			Category:         "Data Sources",
			ShortDescription: `Get information about an existing Managed Disk.`,
			Description: `

Use this data source to access information about an existing Managed Disk.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the Managed Disk.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) Specifies the name of the resource group. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "storage_account_type",
					Description: `The storage account type for the managed disk.`,
				},
				resource.Attribute{
					Name:        "source_uri",
					Description: `The source URI for the managed disk`,
				},
				resource.Attribute{
					Name:        "source_resource_id",
					Description: `ID of an existing managed disk that the current resource was created from.`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `The operating system for managed disk. Valid values are ` + "`" + `Linux` + "`" + ` or ` + "`" + `Windows` + "`" + ``,
				},
				resource.Attribute{
					Name:        "disk_size_gb",
					Description: `The size of the managed disk in gigabytes.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A collection containing the availability zone the managed disk is allocated in.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "storage_account_type",
					Description: `The storage account type for the managed disk.`,
				},
				resource.Attribute{
					Name:        "source_uri",
					Description: `The source URI for the managed disk`,
				},
				resource.Attribute{
					Name:        "source_resource_id",
					Description: `ID of an existing managed disk that the current resource was created from.`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `The operating system for managed disk. Valid values are ` + "`" + `Linux` + "`" + ` or ` + "`" + `Windows` + "`" + ``,
				},
				resource.Attribute{
					Name:        "disk_size_gb",
					Description: `The size of the managed disk in gigabytes.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A collection containing the availability zone the managed disk is allocated in.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_management_group",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing Management Group.`,
			Description: `

Use this data source to access information about an existing Management Group.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required) Specifies the UUID of this Management Group. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Management Group.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `A friendly name for the Management Group.`,
				},
				resource.Attribute{
					Name:        "parent_management_group_id",
					Description: `The ID of any Parent Management Group.`,
				},
				resource.Attribute{
					Name:        "subscription_ids",
					Description: `A list of Subscription ID's which are assigned to the Management Group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Management Group.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `A friendly name for the Management Group.`,
				},
				resource.Attribute{
					Name:        "parent_management_group_id",
					Description: `The ID of any Parent Management Group.`,
				},
				resource.Attribute{
					Name:        "subscription_ids",
					Description: `A list of Subscription ID's which are assigned to the Management Group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_maps_account",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing Azure Maps Account.`,
			Description: `

Use this data source to access information about an existing Azure Maps Account.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the Maps Account.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) Specifies the name of the Resource Group in which the Maps Account is located. ## Attributes Reference`,
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
					Description: `A unique identifier for the Maps Account.`,
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
					Description: `A unique identifier for the Maps Account.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_monitor_action_group",
			Category:         "Data Sources",
			ShortDescription: `Get information about the specified Action Group.`,
			Description: `

Use this data source to access the properties of an Action Group.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the Action Group.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) Specifies the name of the resource group the Action Group is located in. ## Attributes Reference`,
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
					Name:        "email_receiver",
					Description: `One or more ` + "`" + `email_receiver` + "`" + ` blocks as defined below.`,
				},
				resource.Attribute{
					Name:        "sms_receiver",
					Description: `One or more ` + "`" + `sms_receiver ` + "`" + ` blocks as defined below.`,
				},
				resource.Attribute{
					Name:        "webhook_receiver",
					Description: `One or more ` + "`" + `webhook_receiver ` + "`" + ` blocks as defined below. --- ` + "`" + `email_receiver` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the email receiver.`,
				},
				resource.Attribute{
					Name:        "email_address",
					Description: `The email address of this receiver. --- ` + "`" + `sms_receiver` + "`" + ` supports the following:`,
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
					Description: `The phone number of the SMS receiver. --- ` + "`" + `webhook_receiver` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the webhook receiver.`,
				},
				resource.Attribute{
					Name:        "service_uri",
					Description: `The URI where webhooks should be sent.`,
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
					Name:        "email_receiver",
					Description: `One or more ` + "`" + `email_receiver` + "`" + ` blocks as defined below.`,
				},
				resource.Attribute{
					Name:        "sms_receiver",
					Description: `One or more ` + "`" + `sms_receiver ` + "`" + ` blocks as defined below.`,
				},
				resource.Attribute{
					Name:        "webhook_receiver",
					Description: `One or more ` + "`" + `webhook_receiver ` + "`" + ` blocks as defined below. --- ` + "`" + `email_receiver` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the email receiver.`,
				},
				resource.Attribute{
					Name:        "email_address",
					Description: `The email address of this receiver. --- ` + "`" + `sms_receiver` + "`" + ` supports the following:`,
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
					Description: `The phone number of the SMS receiver. --- ` + "`" + `webhook_receiver` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the webhook receiver.`,
				},
				resource.Attribute{
					Name:        "service_uri",
					Description: `The URI where webhooks should be sent.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_monitor_diagnostic_categories",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an the Monitor Diagnostics Categories supported by an existing Resource.`,
			Description: `

Use this data source to access information about the Monitor Diagnostics Categories supported by an existing Resource.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required) The ID of an existing Resource which Monitor Diagnostics Categories should be retrieved for. ## Attributes Reference`,
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
					Description: `A list of the Metric Categories supported for this Resource.`,
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
					Description: `A list of the Metric Categories supported for this Resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_monitor_log_profile",
			Category:         "Data Sources",
			ShortDescription: `Get information about the specified Log Profile.`,
			Description: `

Use this data source to access the properties of a Log Profile.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the Name of the Log Profile. ## Attributes Reference`,
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
					Description: `The number of days for the retention policy.`,
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
					Description: `The number of days for the retention policy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_mssql_elasticpool",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing SQL elastic pool.`,
			Description: `

Use this data source to access information about an existing SQL elastic pool.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the elastic pool.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The name of the resource group which contains the elastic pool.`,
				},
				resource.Attribute{
					Name:        "server_name",
					Description: `(Required) The name of the SQL Server which contains the elastic pool. ## Attributes Reference`,
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
					Description: `Whether or not this elastic pool is zone redundant.`,
				},
			},
			Attributes: []resource.Attribute{
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
					Description: `Whether or not this elastic pool is zone redundant.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_network_interface",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing Network Interface.`,
			Description: `

Use this data source to access information about an existing Network Interface.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the Network Interface.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) Specifies the name of the resource group the Network Interface is located in. ## Attributes Reference`,
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
					Description: `is this the Primary IP Configuration for this Network Interface?`,
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
					Description: `is this the Primary IP Configuration for this Network Interface?`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_network_security_group",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing Network Security Group.`,
			Description: `

Use this data source to access information about an existing Network Security Group.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the Name of the Network Security Group.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) Specifies the Name of the Resource Group within which the Network Security Group exists ## Attributes Reference`,
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
					Description: `The direction specifies if rule will be evaluated on incoming or outgoing traffic.`,
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
					Description: `The direction specifies if rule will be evaluated on incoming or outgoing traffic.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_network_watcher",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing Network Watcher.`,
			Description: `

Use this data source to access information about an existing Network Watcher.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the Name of the Network Watcher.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) Specifies the Name of the Resource Group within which the Network Watcher exists. ## Attributes Reference`,
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
					Description: `A mapping of tags assigned to the resource.`,
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
					Description: `A mapping of tags assigned to the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_notification_hub_namespace",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing Notification Hub Namespace.`,
			Description: `

Use this data source to access information about an existing Notification Hub Namespace.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the Name of the Notification Hub Namespace.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) Specifies the Name of the Resource Group within which the Notification Hub exists. ## Attributes Reference`,
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
					Description: `Is this Notification Hub Namespace enabled? --- A ` + "`" + `sku` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the SKU to use for this Notification Hub Namespace. Possible values are ` + "`" + `Free` + "`" + `, ` + "`" + `Basic` + "`" + ` or ` + "`" + `Standard.` + "`" + ``,
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
					Description: `Is this Notification Hub Namespace enabled? --- A ` + "`" + `sku` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the SKU to use for this Notification Hub Namespace. Possible values are ` + "`" + `Free` + "`" + `, ` + "`" + `Basic` + "`" + ` or ` + "`" + `Standard.` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_platform_image",
			Category:         "Data Sources",
			ShortDescription: `Gets information about a Platform Image.`,
			Description: `

Use this data source to access information about a Platform Image.

`,
			Keywords: []string{},
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
					Description: `(Required) Specifies the SKU of the Platform Image. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Platform Image.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The latest version of the Platform Image.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Platform Image.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The latest version of the Platform Image.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_policy_definition",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Policy Definition.`,
			Description: `

Use this data source to access information about a Policy Definition, both custom and built in. Retrieves Policy Definitions from your current subscription by default.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) Specifies the name of the Policy Definition.`,
				},
				resource.Attribute{
					Name:        "management_group_id",
					Description: `(Optional) Only retrieve Policy Definitions from this Management Group. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Policy Definition.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The Name of the Policy Definition.`,
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
					Description: `The Type of the Policy, such as ` + "`" + `Microsoft.Authorization/policyDefinitions` + "`" + `.`,
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
					Description: `Any Metadata defined in the Policy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Policy Definition.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The Name of the Policy Definition.`,
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
					Description: `The Type of the Policy, such as ` + "`" + `Microsoft.Authorization/policyDefinitions` + "`" + `.`,
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
					Description: `Any Metadata defined in the Policy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_public_ip",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing Public IP Address.`,
			Description: `

Use this data source to access information about an existing Public IP Address.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the public IP address.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) Specifies the name of the resource group. ## Attributes Reference`,
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
					Name:        "tags",
					Description: `A mapping of tags to assigned to the resource.`,
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
					Name:        "tags",
					Description: `A mapping of tags to assigned to the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_public_ips",
			Category:         "Data Sources",
			ShortDescription: `Gets information about a set of existing Public IP Addresses.`,
			Description: `

Use this data source to access information about a set of existing Public IP Addresses.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) Specifies the name of the resource group.`,
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
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_recovery_services_protection_policy_vm",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing Recovery Services VM Protection Policy.`,
			Description: `

Use this data source to access information about an existing Recovery Services VM Protection Policy.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the Recovery Services VM Protection Policy.`,
				},
				resource.Attribute{
					Name:        "recovery_vault_name",
					Description: `(Required) Specifies the name of the Recovery Services Vault.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The name of the resource group in which the Recovery Services VM Protection Policy resides. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Recovery Services VM Protection Policy.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Recovery Services VM Protection Policy.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_recovery_services_vault",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing Recovery Services Vault.`,
			Description: `

Use this data source to access information about an existing Recovery Services Vault.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the Recovery Services Vault.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The name of the resource group in which the Recovery Services Vault resides. ## Attributes Reference The following attributes are exported:`,
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
					Description: `The vault's current SKU.`,
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
					Description: `The vault's current SKU.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_resource_group",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing Resource Group.`,
			Description: `

Use this data source to access information about an existing Resource Group.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the resource group. ~>`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location of the resource group.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "location",
					Description: `The location of the resource group.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_role_definition",
			Category:         "Data Sources",
			ShortDescription: `Get information about an existing Role Definition.`,
			Description: `

Use this data source to access information about an existing Role Definition.

`,
			Keywords: []string{},
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
					Description: `a list of actions which are denied by this role`,
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
					Description: `a list of actions which are denied by this role`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_route_table",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing Route Table`,
			Description: `

Use this data source to access information about an existing Route Table.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Route Table.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The name of the Resource Group in which the Route Table exists. ## Attributes Reference The following attributes are exported:`,
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
					Description: `Contains the IP address packets should be forwarded to.`,
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
					Description: `Contains the IP address packets should be forwarded to.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_scheduler_job_collection",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing Scheduler Job Collection.`,
			Description: `

Use this data source to access information about an existing Scheduler Job Collection.

~> **NOTE:** Support for Scheduler Job Collections has been deprecated by Microsoft in favour of Logic Apps ([more information can be found at this link](https://docs.microsoft.com/en-us/azure/scheduler/migrate-from-scheduler-to-logic-apps)) - as such we plan to remove support for this data source as a part of version 2.0 of the AzureRM Provider.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the Scheduler Job Collection.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) Specifies the name of the resource group in which the Scheduler Job Collection resides. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Scheduler Job Collection.`,
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
					Name:        "sku",
					Description: `The Job Collection's pricing level's SKU.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The Job Collection's state.`,
				},
				resource.Attribute{
					Name:        "quota",
					Description: `The Job collection quotas as documented in the ` + "`" + `quota` + "`" + ` block below. The ` + "`" + `quota` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "max_job_count",
					Description: `Sets the maximum number of jobs in the collection.`,
				},
				resource.Attribute{
					Name:        "max_recurrence_frequency",
					Description: `The maximum frequency of recurrence.`,
				},
				resource.Attribute{
					Name:        "max_retry_interval",
					Description: `The maximum interval between retries.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Scheduler Job Collection.`,
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
					Name:        "sku",
					Description: `The Job Collection's pricing level's SKU.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The Job Collection's state.`,
				},
				resource.Attribute{
					Name:        "quota",
					Description: `The Job collection quotas as documented in the ` + "`" + `quota` + "`" + ` block below. The ` + "`" + `quota` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "max_job_count",
					Description: `Sets the maximum number of jobs in the collection.`,
				},
				resource.Attribute{
					Name:        "max_recurrence_frequency",
					Description: `The maximum frequency of recurrence.`,
				},
				resource.Attribute{
					Name:        "max_retry_interval",
					Description: `The maximum interval between retries.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_servicebus_namespace",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing ServiceBus Namespace.`,
			Description: `

Use this data source to access information about an existing ServiceBus Namespace.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the ServiceBus Namespace.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) Specifies the name of the Resource Group where the ServiceBus Namespace exists. ## Attributes Reference`,
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
					Description: `The secondary access key for the authorization rule ` + "`" + `RootManageSharedAccessKey` + "`" + `.`,
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
					Description: `The secondary access key for the authorization rule ` + "`" + `RootManageSharedAccessKey` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_shared_image",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing Shared Image a Shared Image Gallery.`,
			Description: `

Use this data source to access information about an existing Shared Image within a Shared Image Gallery.

-> **NOTE** Shared Image Galleries are currently in Public Preview. You can find more information, including [how to register for the Public Preview here](https://azure.microsoft.com/en-gb/blog/announcing-the-public-preview-of-shared-image-gallery/).

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Shared Image.`,
				},
				resource.Attribute{
					Name:        "gallery_name",
					Description: `(Required) The name of the Shared Image Gallery in which the Shared Image exists.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The name of the Resource Group in which the Shared Image Gallery exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Resource ID of the Shared Image.`,
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
					Name:        "identifier",
					Description: `An ` + "`" + `identifier` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `The type of Operating System present in this Shared Image.`,
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
					Description: `The Name of the SKU for this Gallery Image.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Resource ID of the Shared Image.`,
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
					Name:        "identifier",
					Description: `An ` + "`" + `identifier` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `The type of Operating System present in this Shared Image.`,
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
					Description: `The Name of the SKU for this Gallery Image.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_shared_image_gallery",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing Shared Image Gallery.`,
			Description: `

Use this data source to access information about an existing Shared Image Gallery.

-> **NOTE** Shared Image Galleries are currently in Public Preview. You can find more information, including [how to register for the Public Preview here](https://azure.microsoft.com/en-gb/blog/announcing-the-public-preview-of-shared-image-gallery/).

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Shared Image Gallery.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The name of the Resource Group in which the Shared Image Gallery exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Resource ID of the Shared Image Gallery.`,
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
					Description: `A mapping of tags which are assigned to the Shared Image Gallery.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Resource ID of the Shared Image Gallery.`,
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
					Description: `A mapping of tags which are assigned to the Shared Image Gallery.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_shared_image_version",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing Version of a Shared Image within a Shared Image Gallery.`,
			Description: `

Use this data source to access information about an existing Version of a Shared Image within a Shared Image Gallery.

-> **NOTE** Shared Image Galleries are currently in Public Preview. You can find more information, including [how to register for the Public Preview here](https://azure.microsoft.com/en-gb/blog/announcing-the-public-preview-of-shared-image-gallery/).

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Image Version.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `(Required) The name of the Shared Image in which this Version exists.`,
				},
				resource.Attribute{
					Name:        "gallery_name",
					Description: `(Required) The name of the Shared Image in which the Shared Image exists.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The name of the Resource Group in which the Shared Image Gallery exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Resource ID of the Shared Image.`,
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
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Resource ID of the Shared Image.`,
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
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_sql_server",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing SQL Azure Database Server.`,
			Description: `

Use this data source to access information about an existing SQL Azure Database Server.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the SQL Server.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) Specifies the name of the Resource Group where the SQL Server exists. ## Attributes Reference`,
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
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource.`,
				},
			},
			Attributes: []resource.Attribute{
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
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_storage_account",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing Storage Account.`,
			Description: `

Use this data source to access information about an existing Storage Account.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the Storage Account`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) Specifies the name of the resource group the Storage Account is located in. ## Attributes Reference`,
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
					Name:        "enable_blob_encryption",
					Description: `Are Encryption Services are enabled for Blob storage? See [here](https://azure.microsoft.com/en-us/documentation/articles/storage-service-encryption/) for more information.`,
				},
				resource.Attribute{
					Name:        "enable_file_encryption",
					Description: `Are Encryption Services are enabled for File storage? See [here](https://azure.microsoft.com/en-us/documentation/articles/storage-service-encryption/) for more information.`,
				},
				resource.Attribute{
					Name:        "enable_https_traffic_only",
					Description: `Is traffic only allowed via HTTPS? See [here](https://docs.microsoft.com/en-us/azure/storage/storage-require-secure-transfer/) for more information.`,
				},
				resource.Attribute{
					Name:        "is_hns_enabled",
					Description: `Is Hierarchical Namespace enabled?`,
				},
				resource.Attribute{
					Name:        "account_encryption_source",
					Description: `The Encryption Source for this Storage Account.`,
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
					Description: `The connection string associated with the secondary blob location ---`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The Custom Domain Name used for the Storage Account.`,
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
					Name:        "enable_blob_encryption",
					Description: `Are Encryption Services are enabled for Blob storage? See [here](https://azure.microsoft.com/en-us/documentation/articles/storage-service-encryption/) for more information.`,
				},
				resource.Attribute{
					Name:        "enable_file_encryption",
					Description: `Are Encryption Services are enabled for File storage? See [here](https://azure.microsoft.com/en-us/documentation/articles/storage-service-encryption/) for more information.`,
				},
				resource.Attribute{
					Name:        "enable_https_traffic_only",
					Description: `Is traffic only allowed via HTTPS? See [here](https://docs.microsoft.com/en-us/azure/storage/storage-require-secure-transfer/) for more information.`,
				},
				resource.Attribute{
					Name:        "is_hns_enabled",
					Description: `Is Hierarchical Namespace enabled?`,
				},
				resource.Attribute{
					Name:        "account_encryption_source",
					Description: `The Encryption Source for this Storage Account.`,
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
					Description: `The connection string associated with the secondary blob location ---`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The Custom Domain Name used for the Storage Account.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_storage_account_sas",
			Category:         "Data Sources",
			ShortDescription: `Gets a Shared Access Signature (SAS Token) for an existing Storage Account.`,
			Description: `

Use this data source to obtain a Shared Access Signature (SAS Token) for an existing Storage Account.

Shared access signatures allow fine-grained, ephemeral access control to various aspects of an Azure Storage Account.

Note that this is an [Account SAS](https://docs.microsoft.com/en-us/rest/api/storageservices/constructing-an-account-sas)
and *not* a [Service SAS](https://docs.microsoft.com/en-us/rest/api/storageservices/constructing-a-service-sas).

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "connection_string",
					Description: `(Required) The connection string for the storage account to which this SAS applies. Typically directly from the ` + "`" + `primary_connection_string` + "`" + ` attribute of a terraform created ` + "`" + `azurerm_storage_account` + "`" + ` resource.`,
				},
				resource.Attribute{
					Name:        "https_only",
					Description: `(Optional) Only permit ` + "`" + `https` + "`" + ` access. If ` + "`" + `false` + "`" + `, both ` + "`" + `http` + "`" + ` and ` + "`" + `https` + "`" + ` are permitted. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "resource_types",
					Description: `(Required) A ` + "`" + `resource_types` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `(Required) A ` + "`" + `services` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `(Required) The starting time and date of validity of this SAS. Must be a valid ISO-8601 format time/date string.`,
				},
				resource.Attribute{
					Name:        "expiry",
					Description: `(Required) The expiration time and date of this SAS. Must be a valid ISO-8601 format time/date string.`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `(Required) A ` + "`" + `permissions` + "`" + ` block as defined below. --- ` + "`" + `resource_types` + "`" + ` is a set of ` + "`" + `true` + "`" + `/` + "`" + `false` + "`" + ` flags which define the storage account resource types that are granted access by this SAS. This can be thought of as the scope over which the permissions apply. A ` + "`" + `service` + "`" + ` will have larger scope (affecting all sub-resources) than ` + "`" + `object` + "`" + `. A ` + "`" + `resource_types` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `(Required) Should permission be granted to the entire service?`,
				},
				resource.Attribute{
					Name:        "container",
					Description: `(Required) Should permission be granted to the container?`,
				},
				resource.Attribute{
					Name:        "object",
					Description: `(Required) Should permission be granted only to a specific object? --- ` + "`" + `services` + "`" + ` is a set of ` + "`" + `true` + "`" + `/` + "`" + `false` + "`" + ` flags which define the storage account services that are granted access by this SAS. A ` + "`" + `services` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "blob",
					Description: `(Required) Should permission be granted to ` + "`" + `blob` + "`" + ` services within this storage account?`,
				},
				resource.Attribute{
					Name:        "queue",
					Description: `(Required) Should permission be granted to ` + "`" + `queue` + "`" + ` services within this storage account?`,
				},
				resource.Attribute{
					Name:        "table",
					Description: `(Required) Should permission be granted to ` + "`" + `table` + "`" + ` services within this storage account?`,
				},
				resource.Attribute{
					Name:        "file",
					Description: `(Required) Should permission be granted to ` + "`" + `file` + "`" + ` services within this storage account? --- A ` + "`" + `permissions` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Required) Should Read permissions be enabled for this SAS?`,
				},
				resource.Attribute{
					Name:        "write",
					Description: `(Required) Should Write permissions be enabled for this SAS?`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Required) Should Delete permissions be enabled for this SAS?`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `(Required) Should List permissions be enabled for this SAS?`,
				},
				resource.Attribute{
					Name:        "add",
					Description: `(Required) Should Add permissions be enabled for this SAS?`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Required) Should Create permissions be enabled for this SAS?`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Required) Should Update permissions be enabled for this SAS?`,
				},
				resource.Attribute{
					Name:        "process",
					Description: `(Required) Should Process permissions be enabled for this SAS? Refer to the [SAS creation reference from Azure](https://docs.microsoft.com/en-us/rest/api/storageservices/constructing-an-account-sas) for additional details on the fields above. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "sas",
					Description: `The computed Account Shared Access Signature (SAS).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "sas",
					Description: `The computed Account Shared Access Signature (SAS).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_stream_analytics_job",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing Stream Analytics Job.`,
			Description: `

Use this data source to access information about an existing Stream Analytics Job.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the Stream Analytics Job.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) Specifies the name of the resource group the Stream Analytics Job is located in. ## Attributes Reference`,
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
					Description: `The query that will be run in the streaming job, [written in Stream Analytics Query Language (SAQL)](https://msdn.microsoft.com/library/azure/dn834998).`,
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
					Description: `The query that will be run in the streaming job, [written in Stream Analytics Query Language (SAQL)](https://msdn.microsoft.com/library/azure/dn834998).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_subnet",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing Subnet located within a Virtual Network.`,
			Description: `

Use this data source to access information about an existing Subnet within a Virtual Network.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the Subnet.`,
				},
				resource.Attribute{
					Name:        "virtual_network_name",
					Description: `(Required) Specifies the name of the Virtual Network this Subnet is located within.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) Specifies the name of the resource group the Virtual Network is located in. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Subnet.`,
				},
				resource.Attribute{
					Name:        "address_prefix",
					Description: `The address prefix used for the subnet.`,
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
					Name:        "ip_configurations",
					Description: `The collection of IP Configurations with IPs within this subnet.`,
				},
				resource.Attribute{
					Name:        "service_endpoints",
					Description: `A list of Service Endpoints within this subnet.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Subnet.`,
				},
				resource.Attribute{
					Name:        "address_prefix",
					Description: `The address prefix used for the subnet.`,
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
					Name:        "ip_configurations",
					Description: `The collection of IP Configurations with IPs within this subnet.`,
				},
				resource.Attribute{
					Name:        "service_endpoints",
					Description: `A list of Service Endpoints within this subnet.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_subscription",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing Subscription.`,
			Description: `

Use this data source to access information about an existing Subscription.

`,
			Keywords: []string{},
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
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_subscriptions",
			Category:         "Data Sources",
			ShortDescription: `Get information about the available subscriptions.`,
			Description: `

Use this data source to access information about all the Subscriptions currently available.

`,
			Keywords: []string{},
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
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "subscriptions",
					Description: `One or more ` + "`" + `subscription` + "`" + ` blocks as defined below. The ` + "`" + `subscription` + "`" + ` block contains:`,
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
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_traffic_manager_geographical_location",
			Category:         "Data Sources",
			ShortDescription: `Gets information about a specified Traffic Manager Geographical Location within the Geographical Hierarchy.`,
			Description: `

Use this data source to access the ID of a specified Traffic Manager Geographical Location within the Geographical Hierarchy.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the Location, for example ` + "`" + `World` + "`" + `, ` + "`" + `Europe` + "`" + ` or ` + "`" + `Germany` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of this Location, also known as the ` + "`" + `Code` + "`" + ` of this Location.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of this Location, also known as the ` + "`" + `Code` + "`" + ` of this Location.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_user_assigned_identity",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing User Assigned Identity.`,
			Description: `

Use this data source to access information about an existing User Assigned Identity.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the User Assigned Identity.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The name of the Resource Group in which the User Assigned Identity exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Resource ID of the User Assigned Identity.`,
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
					Description: `A mapping of tags assigned to the User Assigned Identity.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Resource ID of the User Assigned Identity.`,
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
					Description: `A mapping of tags assigned to the User Assigned Identity.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_virtual_machine",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing Virtual Machine.`,
			Description: `

Use this data source to access information about an existing Virtual Machine.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the Virtual Machine.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) Specifies the name of the resource group the Virtual Machine is located in. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Virtual Machine.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Virtual Machine.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_virtual_network",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing Virtual Network.`,
			Description: `

Use this data source to access information about an existing Virtual Network.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the Virtual Network.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) Specifies the name of the resource group the Virtual Network is located in. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the virtual network.`,
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
					Name:        "subnets",
					Description: `The list of name of the subnets that are attached to this virtual network.`,
				},
				resource.Attribute{
					Name:        "vnet_peerings",
					Description: `A mapping of name - virtual network id of the virtual network peerings.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the virtual network.`,
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
					Name:        "subnets",
					Description: `The list of name of the subnets that are attached to this virtual network.`,
				},
				resource.Attribute{
					Name:        "vnet_peerings",
					Description: `A mapping of name - virtual network id of the virtual network peerings.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_virtual_network_gateway",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing Virtual Network Gateway.`,
			Description: `

Use this data source to access information about an existing Virtual Network Gateway.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the Virtual Network Gateway.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) Specifies the name of the resource group the Virtual Network Gateway is located in. ## Attributes Reference`,
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
					Description: `(Optional) Is this an Active-Active Gateway?`,
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
					Name:        "root_certificate",
					Description: `One or more ` + "`" + `root_certificate` + "`" + ` blocks which are defined below. These root certificates are used to sign the client certificate used by the VPN clients to connect to the gateway.`,
				},
				resource.Attribute{
					Name:        "revoked_certificate",
					Description: `One or more ` + "`" + `revoked_certificate` + "`" + ` blocks which are defined below.`,
				},
				resource.Attribute{
					Name:        "radius_server_address",
					Description: `(Optional) The address of the Radius server. This setting is incompatible with the use of ` + "`" + `root_certificate` + "`" + ` and ` + "`" + `revoked_certificate` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "radius_server_secret",
					Description: `(Optional) The secret used by the Radius server. This setting is incompatible with the use of ` + "`" + `root_certificate` + "`" + ` and ` + "`" + `revoked_certificate` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vpn_client_protocols",
					Description: `(Optional) List of the protocols supported by the vpn client. The supported values are ` + "`" + `SSTP` + "`" + `, ` + "`" + `IkeV2` + "`" + ` and ` + "`" + `OpenVPN` + "`" + `. The ` + "`" + `bgp_settings` + "`" + ` block supports:`,
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
					Description: `The weight added to routes which have been learned through BGP peering. The ` + "`" + `root_certificate` + "`" + ` block supports:`,
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
					Description: `The SHA1 thumbprint of the certificate to be revoked.`,
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
					Description: `(Optional) Is this an Active-Active Gateway?`,
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
					Name:        "root_certificate",
					Description: `One or more ` + "`" + `root_certificate` + "`" + ` blocks which are defined below. These root certificates are used to sign the client certificate used by the VPN clients to connect to the gateway.`,
				},
				resource.Attribute{
					Name:        "revoked_certificate",
					Description: `One or more ` + "`" + `revoked_certificate` + "`" + ` blocks which are defined below.`,
				},
				resource.Attribute{
					Name:        "radius_server_address",
					Description: `(Optional) The address of the Radius server. This setting is incompatible with the use of ` + "`" + `root_certificate` + "`" + ` and ` + "`" + `revoked_certificate` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "radius_server_secret",
					Description: `(Optional) The secret used by the Radius server. This setting is incompatible with the use of ` + "`" + `root_certificate` + "`" + ` and ` + "`" + `revoked_certificate` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vpn_client_protocols",
					Description: `(Optional) List of the protocols supported by the vpn client. The supported values are ` + "`" + `SSTP` + "`" + `, ` + "`" + `IkeV2` + "`" + ` and ` + "`" + `OpenVPN` + "`" + `. The ` + "`" + `bgp_settings` + "`" + ` block supports:`,
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
					Description: `The weight added to routes which have been learned through BGP peering. The ` + "`" + `root_certificate` + "`" + ` block supports:`,
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
					Description: `The SHA1 thumbprint of the certificate to be revoked.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurerm_virtual_network_gateway_connection",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing Virtual Network Gateway Connection.`,
			Description: `

Use this data source to access information about an existing Virtual Network Gateway Connection.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the Virtual Network Gateway Connection.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) Specifies the name of the resource group the Virtual Network Gateway Connection is located in. ## Attributes Reference`,
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
					Name:        "express_route_circuit_id",
					Description: `The ID of the Express Route Circuit (i.e. when ` + "`" + `type` + "`" + ` is ` + "`" + `ExpressRoute` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "peer_virtual_network_gateway_id",
					Description: `The ID of the peer virtual network gateway when a VNet-to-VNet connection (i.e. when ` + "`" + `type` + "`" + ` is ` + "`" + `Vnet2Vnet` + "`" + `).`,
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
					Description: `(Optional) A mapping of tags to assign to the resource. The ` + "`" + `ipsec_policy` + "`" + ` block supports:`,
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
					Description: `The IPSec SA lifetime in seconds. Must be at least ` + "`" + `300` + "`" + ` seconds.`,
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
					Name:        "express_route_circuit_id",
					Description: `The ID of the Express Route Circuit (i.e. when ` + "`" + `type` + "`" + ` is ` + "`" + `ExpressRoute` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "peer_virtual_network_gateway_id",
					Description: `The ID of the peer virtual network gateway when a VNet-to-VNet connection (i.e. when ` + "`" + `type` + "`" + ` is ` + "`" + `Vnet2Vnet` + "`" + `).`,
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
					Description: `(Optional) A mapping of tags to assign to the resource. The ` + "`" + `ipsec_policy` + "`" + ` block supports:`,
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
					Description: `The IPSec SA lifetime in seconds. Must be at least ` + "`" + `300` + "`" + ` seconds.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"azurerm_api_management":                         0,
		"azurerm_api_management_api":                     1,
		"azurerm_api_management_group":                   2,
		"azurerm_api_management_product":                 3,
		"azurerm_api_management_user":                    4,
		"azurerm_app_service":                            5,
		"azurerm_app_service_plan":                       6,
		"azurerm_application_insights":                   7,
		"azurerm_application_security_group":             8,
		"azurerm_automation_variable_bool":               9,
		"azurerm_automation_variable_datetime":           10,
		"azurerm_automation_variable_int":                11,
		"azurerm_automation_variable_string":             12,
		"azurerm_availability_set":                       13,
		"azurerm_azuread_application":                    14,
		"azurerm_azuread_service_principal":              15,
		"azurerm_batch_account":                          16,
		"azurerm_batch_certificate":                      17,
		"azurerm_batch_pool":                             18,
		"azurerm_builtin_role_definition":                19,
		"azurerm_cdn_profile":                            20,
		"azurerm_client_config":                          21,
		"azurerm_container_registry":                     22,
		"azurerm_cosmosdb_account":                       23,
		"azurerm_data_lake_store":                        24,
		"azurerm_dev_test_lab":                           25,
		"azurerm_dns_zone":                               26,
		"azurerm_eventhub_namespace":                     27,
		"azurerm_express_route_circuit":                  28,
		"azurerm_firewall":                               29,
		"azurerm_hdinsight_cluster":                      30,
		"azurerm_image":                                  31,
		"azurerm_key_vault":                              32,
		"azurerm_key_vault_access_policy":                33,
		"azurerm_key_vault_key":                          34,
		"azurerm_key_vault_secret":                       35,
		"azurerm_kubernetes_cluster":                     36,
		"azurerm_kubernetes_service_versions":            37,
		"azurerm_lb":                                     38,
		"azurerm_lb_backend_address_pool":                39,
		"azurerm_log_analytics_workspace":                40,
		"azurerm_logic_app_workflow":                     41,
		"azurerm_managed_disk":                           42,
		"azurerm_management_group":                       43,
		"azurerm_maps_account":                           44,
		"azurerm_monitor_action_group":                   45,
		"azurerm_monitor_diagnostic_categories":          46,
		"azurerm_monitor_log_profile":                    47,
		"azurerm_mssql_elasticpool":                      48,
		"azurerm_network_interface":                      49,
		"azurerm_network_security_group":                 50,
		"azurerm_network_watcher":                        51,
		"azurerm_notification_hub_namespace":             52,
		"azurerm_platform_image":                         53,
		"azurerm_policy_definition":                      54,
		"azurerm_public_ip":                              55,
		"azurerm_public_ips":                             56,
		"azurerm_recovery_services_protection_policy_vm": 57,
		"azurerm_recovery_services_vault":                58,
		"azurerm_resource_group":                         59,
		"azurerm_role_definition":                        60,
		"azurerm_route_table":                            61,
		"azurerm_scheduler_job_collection":               62,
		"azurerm_servicebus_namespace":                   63,
		"azurerm_shared_image":                           64,
		"azurerm_shared_image_gallery":                   65,
		"azurerm_shared_image_version":                   66,
		"azurerm_sql_server":                             67,
		"azurerm_storage_account":                        68,
		"azurerm_storage_account_sas":                    69,
		"azurerm_stream_analytics_job":                   70,
		"azurerm_subnet":                                 71,
		"azurerm_subscription":                           72,
		"azurerm_subscriptions":                          73,
		"azurerm_traffic_manager_geographical_location":  74,
		"azurerm_user_assigned_identity":                 75,
		"azurerm_virtual_machine":                        76,
		"azurerm_virtual_network":                        77,
		"azurerm_virtual_network_gateway":                78,
		"azurerm_virtual_network_gateway_connection":     79,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
