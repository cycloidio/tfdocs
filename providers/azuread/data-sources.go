package azuread

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "azuread_access_package",
			Category:         "Identity Governance",
			ShortDescription: ``,
			Description: `

Use this data source to retrieve information for an existing access package within Identity Governance in Azure Active Directory.

`,
			Keywords: []string{
				"identity",
				"governance",
				"access",
				"package",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "catalog_id",
					Description: `(Optional) The ID of the Catalog this access package is in.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of the access package.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `(Optional) The ID of this access package. ~> Either ` + "`" + `object_id` + "`" + `, or both ` + "`" + `catalog_id` + "`" + ` and ` + "`" + `display_name` + "`" + `, must be specified. ## Attributes Reference In addition to the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of this resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the access package.`,
				},
				resource.Attribute{
					Name:        "hidden",
					Description: `Whether the access package is hidden from the requestor.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of this resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the access package.`,
				},
				resource.Attribute{
					Name:        "hidden",
					Description: `Whether the access package is hidden from the requestor.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_access_package_catalog",
			Category:         "Identity Governance",
			ShortDescription: ``,
			Description: `
i
Use this resource to retrieve information for an existing access package catalog within Identity Governance in Azure Active Directory.

`,
			Keywords: []string{
				"identity",
				"governance",
				"access",
				"package",
				"catalog",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of the access package catalog.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `(Optional) The ID of this access package catalog. ~> One of ` + "`" + `display_name` + "`" + ` or ` + "`" + `object_id` + "`" + ` must be specified. ## Attributes Reference In additional to the arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of this resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the access package catalog.`,
				},
				resource.Attribute{
					Name:        "externally_visible",
					Description: `Whether the access packages in this catalog can be requested by users outside the tenant.`,
				},
				resource.Attribute{
					Name:        "published",
					Description: `Whether the access packages in this catalog are available for management.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of this resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the access package catalog.`,
				},
				resource.Attribute{
					Name:        "externally_visible",
					Description: `Whether the access packages in this catalog can be requested by users outside the tenant.`,
				},
				resource.Attribute{
					Name:        "published",
					Description: `Whether the access packages in this catalog are available for management.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_access_package_catalog_role",
			Category:         "Identity Governance",
			ShortDescription: ``,
			Description: `

Gets information about an access package catalog role.

`,
			Keywords: []string{
				"identity",
				"governance",
				"access",
				"package",
				"catalog",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Specifies the display name of the role.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `(Optional) Specifies the object ID of the role. ~> One of ` + "`" + `display_name` + "`" + ` or ` + "`" + `object_id` + "`" + ` must be specified. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the role.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name of the role.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The object ID of the role.`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `The object ID of the role.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the role.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name of the role.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The object ID of the role.`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `The object ID of the role.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_administrative_unit",
			Category:         "Administrative Units",
			ShortDescription: ``,
			Description: `

Gets information about an adminisrative unit in Azure Active Directory.

`,
			Keywords: []string{
				"administrative",
				"units",
				"unit",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Specifies the display name of the administrative unit.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `(Optional) Specifies the object ID of the administrative unit. ~> One of ` + "`" + `display_name` + "`" + ` or ` + "`" + `object_id` + "`" + ` must be specified. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the administrative unit.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name of the administrative unit.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `A list of object IDs of members who are present in this administrative unit.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The object ID of the administrative unit.`,
				},
				resource.Attribute{
					Name:        "visibility",
					Description: `Whether the administrative unit _and_ its members are hidden or publicly viewable in the directory. One of: ` + "`" + `Hiddenmembership` + "`" + ` or ` + "`" + `Public` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the administrative unit.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name of the administrative unit.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `A list of object IDs of members who are present in this administrative unit.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The object ID of the administrative unit.`,
				},
				resource.Attribute{
					Name:        "visibility",
					Description: `Whether the administrative unit _and_ its members are hidden or publicly viewable in the directory. One of: ` + "`" + `Hiddenmembership` + "`" + ` or ` + "`" + `Public` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_application",
			Category:         "Applications",
			ShortDescription: ``,
			Description: `

Use this data source to access information about an existing Application within Azure Active Directory.

`,
			Keywords: []string{
				"applications",
				"application",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "application_id",
					Description: `(Optional) Specifies the Application ID (also called Client ID).`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Specifies the display name of the application.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `(Optional) Specifies the Object ID of the application. ~> One of ` + "`" + `object_id` + "`" + `, ` + "`" + `application_id` + "`" + ` or ` + "`" + `display_name` + "`" + ` must be specified. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "api",
					Description: `An ` + "`" + `api` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "app_role_ids",
					Description: `A mapping of app role values to app role IDs, intended to be useful when referencing app roles in other resources in your configuration.`,
				},
				resource.Attribute{
					Name:        "app_roles",
					Description: `A collection of ` + "`" + `app_role` + "`" + ` blocks as documented below. For more information see [official documentation on Application Roles](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).`,
				},
				resource.Attribute{
					Name:        "application_id",
					Description: `The Application ID (also called Client ID).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the application, as shown to end users.`,
				},
				resource.Attribute{
					Name:        "device_only_auth_enabled",
					Description: `Specifies whether this application supports device authentication without a user.`,
				},
				resource.Attribute{
					Name:        "disabled_by_microsoft",
					Description: `Whether Microsoft has disabled the registered application. If the application is disabled, this will be a string indicating the status/reason, e.g. ` + "`" + `DisabledDueToViolationOfServicesAgreement` + "`" + ``,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name for the application.`,
				},
				resource.Attribute{
					Name:        "fallback_public_client_enabled",
					Description: `The fallback application type as public client, such as an installed application running on a mobile device.`,
				},
				resource.Attribute{
					Name:        "feature_tags",
					Description: `A ` + "`" + `features` + "`" + ` block as described below.`,
				},
				resource.Attribute{
					Name:        "group_membership_claims",
					Description: `The ` + "`" + `groups` + "`" + ` claim issued in a user or OAuth 2.0 access token that the app expects.`,
				},
				resource.Attribute{
					Name:        "identifier_uris",
					Description: `A list of user-defined URI(s) that uniquely identify a Web application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant.`,
				},
				resource.Attribute{
					Name:        "logo_url",
					Description: `CDN URL to the application's logo.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `User-specified notes relevant for the management of the application.`,
				},
				resource.Attribute{
					Name:        "marketing_url",
					Description: `URL of the application's marketing page.`,
				},
				resource.Attribute{
					Name:        "oauth2_permission_scope_ids",
					Description: `A mapping of OAuth2.0 permission scope values to scope IDs, intended to be useful when referencing permission scopes in other resources in your configuration.`,
				},
				resource.Attribute{
					Name:        "oauth2_post_response_required",
					Description: `Specifies whether, as part of OAuth 2.0 token requests, Azure AD allows POST requests, as opposed to GET requests. When ` + "`" + `false` + "`" + `, only GET requests are allowed.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The application's object ID.`,
				},
				resource.Attribute{
					Name:        "optional_claims",
					Description: `An ` + "`" + `optional_claims` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "owners",
					Description: `A list of object IDs of principals that are assigned ownership of the application.`,
				},
				resource.Attribute{
					Name:        "privacy_statement_url",
					Description: `URL of the application's privacy statement.`,
				},
				resource.Attribute{
					Name:        "public_client",
					Description: `A ` + "`" + `public_client` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "publisher_domain",
					Description: `The verified publisher domain for the application.`,
				},
				resource.Attribute{
					Name:        "required_resource_access",
					Description: `A collection of ` + "`" + `required_resource_access` + "`" + ` blocks as documented below.`,
				},
				resource.Attribute{
					Name:        "service_management_reference",
					Description: `References application context information from a Service or Asset Management database.`,
				},
				resource.Attribute{
					Name:        "sign_in_audience",
					Description: `The Microsoft account types that are supported for the current application. One of ` + "`" + `AzureADMyOrg` + "`" + `, ` + "`" + `AzureADMultipleOrgs` + "`" + `, ` + "`" + `AzureADandPersonalMicrosoftAccount` + "`" + ` or ` + "`" + `PersonalMicrosoftAccount` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "single_page_application",
					Description: `A ` + "`" + `single_page_application` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "support_url",
					Description: `URL of the application's support page.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of tags applied to the application.`,
				},
				resource.Attribute{
					Name:        "terms_of_service_url",
					Description: `URL of the application's terms of service statement.`,
				},
				resource.Attribute{
					Name:        "web",
					Description: `A ` + "`" + `web` + "`" + ` block as documented below. --- ` + "`" + `api` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "known_client_applications",
					Description: `A set of application IDs (client IDs), used for bundling consent if you have a solution that contains two parts: a client app and a custom web API app.`,
				},
				resource.Attribute{
					Name:        "mapped_claims_enabled",
					Description: `Allows an application to use claims mapping without specifying a custom signing key.`,
				},
				resource.Attribute{
					Name:        "oauth2_permission_scopes",
					Description: `One or more ` + "`" + `oauth2_permission_scope` + "`" + ` blocks as documented below, to describe delegated permissions exposed by the web API represented by this application.`,
				},
				resource.Attribute{
					Name:        "requested_access_token_version",
					Description: `The access token version expected by this resource. Possible values are ` + "`" + `1` + "`" + ` or ` + "`" + `2` + "`" + `. --- ` + "`" + `oauth2_permission_scope` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "admin_consent_description",
					Description: `Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.`,
				},
				resource.Attribute{
					Name:        "admin_consent_display_name",
					Description: `Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Determines if the permission scope is enabled.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier of the delegated permission. Must be a valid UUID.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Possible values are ` + "`" + `User` + "`" + ` or ` + "`" + `Admin` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user_consent_description",
					Description: `Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.`,
				},
				resource.Attribute{
					Name:        "user_consent_display_name",
					Description: `Display name for the delegated permission that appears in the end user consent experience.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value that is used for the ` + "`" + `scp` + "`" + ` claim in OAuth 2.0 access tokens. --- ` + "`" + `app_role` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "allowed_member_types",
					Description: `Specifies whether this app role definition can be assigned to users and groups, or to other applications (that are accessing this application in a standalone scenario). Possible values are ` + "`" + `User` + "`" + ` or ` + "`" + `Application` + "`" + `, or both.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the app role that appears when the role is being assigned and, if the role functions as an application permissions, during the consent experiences.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Display name for the app role that appears during app role assignment and in consent experiences.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Determines if the app role is enabled.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier of the app role.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value that is used for the ` + "`" + `roles` + "`" + ` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal. --- ` + "`" + `features` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "custom_single_sign_on",
					Description: `Whether this application represents a custom SAML application for linked service principals.`,
				},
				resource.Attribute{
					Name:        "enterprise",
					Description: `Whether this application represents an Enterprise Application for linked service principals.`,
				},
				resource.Attribute{
					Name:        "gallery",
					Description: `Whether this application represents a gallery application for linked service principals.`,
				},
				resource.Attribute{
					Name:        "hide",
					Description: `Whether this app is visible to users in My Apps and Office 365 Launcher. --- ` + "`" + `optional_claims` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "access_token",
					Description: `One or more ` + "`" + `access_token` + "`" + ` blocks as documented below.`,
				},
				resource.Attribute{
					Name:        "id_token",
					Description: `One or more ` + "`" + `id_token` + "`" + ` blocks as documented below.`,
				},
				resource.Attribute{
					Name:        "saml2_token",
					Description: `One or more ` + "`" + `saml2_token` + "`" + ` blocks as documented below. --- ` + "`" + `access_token` + "`" + `, ` + "`" + `id_token` + "`" + ` and ` + "`" + `saml2_token` + "`" + ` blocks export the following:`,
				},
				resource.Attribute{
					Name:        "additional_properties",
					Description: `List of Additional Properties of the claim. If a property exists in this list, it modifies the behaviour of the optional claim.`,
				},
				resource.Attribute{
					Name:        "essential",
					Description: `Whether the claim specified by the client is necessary to ensure a smooth authorization experience.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the optional claim.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `The source of the claim. If ` + "`" + `source` + "`" + ` is absent, the claim is a predefined optional claim. If ` + "`" + `source` + "`" + ` is ` + "`" + `user` + "`" + `, the value of ` + "`" + `name` + "`" + ` is the extension property from the user object. --- ` + "`" + `public_client` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "redirect_uris",
					Description: `A list of URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent. --- ` + "`" + `required_resource_access` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "resource_access",
					Description: `A collection of ` + "`" + `resource_access` + "`" + ` blocks as documented below, describing OAuth2.0 permission scopes and app roles that the application requires from the specified resource.`,
				},
				resource.Attribute{
					Name:        "resource_app_id",
					Description: `The unique identifier for the resource that the application requires access to. This is the Application ID of the target application. --- ` + "`" + `resource_access` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier for an app role or OAuth2 permission scope published by the resource application.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Specifies whether the ` + "`" + `id` + "`" + ` property references an app role or an OAuth2 permission scope. Possible values are ` + "`" + `Role` + "`" + ` or ` + "`" + `Scope` + "`" + `. --- ` + "`" + `single_page_application` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "redirect_uris",
					Description: `A list of URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent. --- ` + "`" + `web` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "homepage_url",
					Description: `Home page or landing page of the application.`,
				},
				resource.Attribute{
					Name:        "implicit_grant",
					Description: `An ` + "`" + `implicit_grant` + "`" + ` block as documented above.`,
				},
				resource.Attribute{
					Name:        "logout_url",
					Description: `The URL that will be used by Microsoft's authorization service to sign out a user using front-channel, back-channel or SAML logout protocols.`,
				},
				resource.Attribute{
					Name:        "redirect_uris",
					Description: `A list of URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent. --- ` + "`" + `implicit_grant` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "access_token_issuance_enabled",
					Description: `Whether this web application can request an access token using OAuth 2.0 implicit flow.`,
				},
				resource.Attribute{
					Name:        "id_token_issuance_enabled",
					Description: `Whether this web application can request an ID token using OAuth 2.0 implicit flow.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "api",
					Description: `An ` + "`" + `api` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "app_role_ids",
					Description: `A mapping of app role values to app role IDs, intended to be useful when referencing app roles in other resources in your configuration.`,
				},
				resource.Attribute{
					Name:        "app_roles",
					Description: `A collection of ` + "`" + `app_role` + "`" + ` blocks as documented below. For more information see [official documentation on Application Roles](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).`,
				},
				resource.Attribute{
					Name:        "application_id",
					Description: `The Application ID (also called Client ID).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the application, as shown to end users.`,
				},
				resource.Attribute{
					Name:        "device_only_auth_enabled",
					Description: `Specifies whether this application supports device authentication without a user.`,
				},
				resource.Attribute{
					Name:        "disabled_by_microsoft",
					Description: `Whether Microsoft has disabled the registered application. If the application is disabled, this will be a string indicating the status/reason, e.g. ` + "`" + `DisabledDueToViolationOfServicesAgreement` + "`" + ``,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name for the application.`,
				},
				resource.Attribute{
					Name:        "fallback_public_client_enabled",
					Description: `The fallback application type as public client, such as an installed application running on a mobile device.`,
				},
				resource.Attribute{
					Name:        "feature_tags",
					Description: `A ` + "`" + `features` + "`" + ` block as described below.`,
				},
				resource.Attribute{
					Name:        "group_membership_claims",
					Description: `The ` + "`" + `groups` + "`" + ` claim issued in a user or OAuth 2.0 access token that the app expects.`,
				},
				resource.Attribute{
					Name:        "identifier_uris",
					Description: `A list of user-defined URI(s) that uniquely identify a Web application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant.`,
				},
				resource.Attribute{
					Name:        "logo_url",
					Description: `CDN URL to the application's logo.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `User-specified notes relevant for the management of the application.`,
				},
				resource.Attribute{
					Name:        "marketing_url",
					Description: `URL of the application's marketing page.`,
				},
				resource.Attribute{
					Name:        "oauth2_permission_scope_ids",
					Description: `A mapping of OAuth2.0 permission scope values to scope IDs, intended to be useful when referencing permission scopes in other resources in your configuration.`,
				},
				resource.Attribute{
					Name:        "oauth2_post_response_required",
					Description: `Specifies whether, as part of OAuth 2.0 token requests, Azure AD allows POST requests, as opposed to GET requests. When ` + "`" + `false` + "`" + `, only GET requests are allowed.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The application's object ID.`,
				},
				resource.Attribute{
					Name:        "optional_claims",
					Description: `An ` + "`" + `optional_claims` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "owners",
					Description: `A list of object IDs of principals that are assigned ownership of the application.`,
				},
				resource.Attribute{
					Name:        "privacy_statement_url",
					Description: `URL of the application's privacy statement.`,
				},
				resource.Attribute{
					Name:        "public_client",
					Description: `A ` + "`" + `public_client` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "publisher_domain",
					Description: `The verified publisher domain for the application.`,
				},
				resource.Attribute{
					Name:        "required_resource_access",
					Description: `A collection of ` + "`" + `required_resource_access` + "`" + ` blocks as documented below.`,
				},
				resource.Attribute{
					Name:        "service_management_reference",
					Description: `References application context information from a Service or Asset Management database.`,
				},
				resource.Attribute{
					Name:        "sign_in_audience",
					Description: `The Microsoft account types that are supported for the current application. One of ` + "`" + `AzureADMyOrg` + "`" + `, ` + "`" + `AzureADMultipleOrgs` + "`" + `, ` + "`" + `AzureADandPersonalMicrosoftAccount` + "`" + ` or ` + "`" + `PersonalMicrosoftAccount` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "single_page_application",
					Description: `A ` + "`" + `single_page_application` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "support_url",
					Description: `URL of the application's support page.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of tags applied to the application.`,
				},
				resource.Attribute{
					Name:        "terms_of_service_url",
					Description: `URL of the application's terms of service statement.`,
				},
				resource.Attribute{
					Name:        "web",
					Description: `A ` + "`" + `web` + "`" + ` block as documented below. --- ` + "`" + `api` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "known_client_applications",
					Description: `A set of application IDs (client IDs), used for bundling consent if you have a solution that contains two parts: a client app and a custom web API app.`,
				},
				resource.Attribute{
					Name:        "mapped_claims_enabled",
					Description: `Allows an application to use claims mapping without specifying a custom signing key.`,
				},
				resource.Attribute{
					Name:        "oauth2_permission_scopes",
					Description: `One or more ` + "`" + `oauth2_permission_scope` + "`" + ` blocks as documented below, to describe delegated permissions exposed by the web API represented by this application.`,
				},
				resource.Attribute{
					Name:        "requested_access_token_version",
					Description: `The access token version expected by this resource. Possible values are ` + "`" + `1` + "`" + ` or ` + "`" + `2` + "`" + `. --- ` + "`" + `oauth2_permission_scope` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "admin_consent_description",
					Description: `Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.`,
				},
				resource.Attribute{
					Name:        "admin_consent_display_name",
					Description: `Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Determines if the permission scope is enabled.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier of the delegated permission. Must be a valid UUID.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Possible values are ` + "`" + `User` + "`" + ` or ` + "`" + `Admin` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user_consent_description",
					Description: `Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.`,
				},
				resource.Attribute{
					Name:        "user_consent_display_name",
					Description: `Display name for the delegated permission that appears in the end user consent experience.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value that is used for the ` + "`" + `scp` + "`" + ` claim in OAuth 2.0 access tokens. --- ` + "`" + `app_role` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "allowed_member_types",
					Description: `Specifies whether this app role definition can be assigned to users and groups, or to other applications (that are accessing this application in a standalone scenario). Possible values are ` + "`" + `User` + "`" + ` or ` + "`" + `Application` + "`" + `, or both.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the app role that appears when the role is being assigned and, if the role functions as an application permissions, during the consent experiences.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Display name for the app role that appears during app role assignment and in consent experiences.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Determines if the app role is enabled.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier of the app role.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value that is used for the ` + "`" + `roles` + "`" + ` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal. --- ` + "`" + `features` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "custom_single_sign_on",
					Description: `Whether this application represents a custom SAML application for linked service principals.`,
				},
				resource.Attribute{
					Name:        "enterprise",
					Description: `Whether this application represents an Enterprise Application for linked service principals.`,
				},
				resource.Attribute{
					Name:        "gallery",
					Description: `Whether this application represents a gallery application for linked service principals.`,
				},
				resource.Attribute{
					Name:        "hide",
					Description: `Whether this app is visible to users in My Apps and Office 365 Launcher. --- ` + "`" + `optional_claims` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "access_token",
					Description: `One or more ` + "`" + `access_token` + "`" + ` blocks as documented below.`,
				},
				resource.Attribute{
					Name:        "id_token",
					Description: `One or more ` + "`" + `id_token` + "`" + ` blocks as documented below.`,
				},
				resource.Attribute{
					Name:        "saml2_token",
					Description: `One or more ` + "`" + `saml2_token` + "`" + ` blocks as documented below. --- ` + "`" + `access_token` + "`" + `, ` + "`" + `id_token` + "`" + ` and ` + "`" + `saml2_token` + "`" + ` blocks export the following:`,
				},
				resource.Attribute{
					Name:        "additional_properties",
					Description: `List of Additional Properties of the claim. If a property exists in this list, it modifies the behaviour of the optional claim.`,
				},
				resource.Attribute{
					Name:        "essential",
					Description: `Whether the claim specified by the client is necessary to ensure a smooth authorization experience.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the optional claim.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `The source of the claim. If ` + "`" + `source` + "`" + ` is absent, the claim is a predefined optional claim. If ` + "`" + `source` + "`" + ` is ` + "`" + `user` + "`" + `, the value of ` + "`" + `name` + "`" + ` is the extension property from the user object. --- ` + "`" + `public_client` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "redirect_uris",
					Description: `A list of URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent. --- ` + "`" + `required_resource_access` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "resource_access",
					Description: `A collection of ` + "`" + `resource_access` + "`" + ` blocks as documented below, describing OAuth2.0 permission scopes and app roles that the application requires from the specified resource.`,
				},
				resource.Attribute{
					Name:        "resource_app_id",
					Description: `The unique identifier for the resource that the application requires access to. This is the Application ID of the target application. --- ` + "`" + `resource_access` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier for an app role or OAuth2 permission scope published by the resource application.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Specifies whether the ` + "`" + `id` + "`" + ` property references an app role or an OAuth2 permission scope. Possible values are ` + "`" + `Role` + "`" + ` or ` + "`" + `Scope` + "`" + `. --- ` + "`" + `single_page_application` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "redirect_uris",
					Description: `A list of URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent. --- ` + "`" + `web` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "homepage_url",
					Description: `Home page or landing page of the application.`,
				},
				resource.Attribute{
					Name:        "implicit_grant",
					Description: `An ` + "`" + `implicit_grant` + "`" + ` block as documented above.`,
				},
				resource.Attribute{
					Name:        "logout_url",
					Description: `The URL that will be used by Microsoft's authorization service to sign out a user using front-channel, back-channel or SAML logout protocols.`,
				},
				resource.Attribute{
					Name:        "redirect_uris",
					Description: `A list of URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent. --- ` + "`" + `implicit_grant` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "access_token_issuance_enabled",
					Description: `Whether this web application can request an access token using OAuth 2.0 implicit flow.`,
				},
				resource.Attribute{
					Name:        "id_token_issuance_enabled",
					Description: `Whether this web application can request an ID token using OAuth 2.0 implicit flow.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_application_published_app_ids",
			Category:         "Applications",
			ShortDescription: ``,
			Description: `

Use this data source to discover application IDs for APIs published by Microsoft.

This data source uses an [unofficial source of application IDs](https://github.com/hashicorp/go-azure-sdk/blob/main/sdk/environments/application_ids.go), as there is currently no available official indexed source for applications or APIs published by Microsoft.

The app IDs returned by this data source are sourced from the Azure Global (Public) Cloud, however some of them are known to work in government and national clouds.

`,
			Keywords: []string{
				"applications",
				"application",
				"published",
				"app",
				"ids",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "result",
					Description: `A map of application names to application IDs.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "result",
					Description: `A map of application names to application IDs.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_application_template",
			Category:         "Applications",
			ShortDescription: ``,
			Description: `

Use this data source to access information about an Application Template from the [Azure AD App Gallery](https://azuremarketplace.microsoft.com/en-US/marketplace/apps/category/azure-active-directory-apps).

`,
			Keywords: []string{
				"applications",
				"application",
				"template",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Specifies the display name of the templated application.`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `(Optional) Specifies the ID of the templated application. ~> One of ` + "`" + `template_id` + "`" + ` or ` + "`" + `display_name` + "`" + ` must be specified. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "categories",
					Description: `List of categories for this templated application.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name for the templated application.`,
				},
				resource.Attribute{
					Name:        "homepage_url",
					Description: `Home page URL of the templated application.`,
				},
				resource.Attribute{
					Name:        "logo_url",
					Description: `URL to retrieve the logo for this templated application.`,
				},
				resource.Attribute{
					Name:        "publisher",
					Description: `Name of the publisher for this templated application.`,
				},
				resource.Attribute{
					Name:        "supported_provisioning_types",
					Description: `List of provisioning modes supported by this templated application.`,
				},
				resource.Attribute{
					Name:        "supported_single_sign_on_modes",
					Description: `List of single sign on modes supported by this templated application.`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `The ID of the templated application.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "categories",
					Description: `List of categories for this templated application.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name for the templated application.`,
				},
				resource.Attribute{
					Name:        "homepage_url",
					Description: `Home page URL of the templated application.`,
				},
				resource.Attribute{
					Name:        "logo_url",
					Description: `URL to retrieve the logo for this templated application.`,
				},
				resource.Attribute{
					Name:        "publisher",
					Description: `Name of the publisher for this templated application.`,
				},
				resource.Attribute{
					Name:        "supported_provisioning_types",
					Description: `List of provisioning modes supported by this templated application.`,
				},
				resource.Attribute{
					Name:        "supported_single_sign_on_modes",
					Description: `List of single sign on modes supported by this templated application.`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `The ID of the templated application.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_client_config",
			Category:         "Base",
			ShortDescription: ``,
			Description: `

Use this data source to access the configuration of the AzureAD provider.

`,
			Keywords: []string{
				"base",
				"client",
				"config",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "client_id",
					Description: `The client ID (application ID) linked to the authenticated principal, or the application used for delegated authentication.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The object ID of the authenticated principal.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `The tenant ID of the authenticated principal.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "client_id",
					Description: `The client ID (application ID) linked to the authenticated principal, or the application used for delegated authentication.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The object ID of the authenticated principal.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `The tenant ID of the authenticated principal.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_directory_object",
			Category:         "Base",
			ShortDescription: ``,
			Description: `

Retrieves the OData type for a generic directory object having the provided object ID.

`,
			Keywords: []string{
				"base",
				"directory",
				"object",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "object_id",
					Description: `(Optional) Specifies the Object ID of the directory object to look up. ## Attributes Reference The following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_directory_roles",
			Category:         "Directory Roles",
			ShortDescription: ``,
			Description: `

Use this data source to access information about activated directory roles within Azure Active Directory.

`,
			Keywords: []string{
				"directory",
				"roles",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "object_ids",
					Description: `The object IDs of the roles.`,
				},
				resource.Attribute{
					Name:        "template_ids",
					Description: `The template IDs of the roles.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `A list of users. Each ` + "`" + `role` + "`" + ` object provides the attributes documented below. --- ` + "`" + `role` + "`" + ` object exports the following:`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name of the directory role.`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `The template ID of the directory role.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the directory role.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The object ID of the directory role.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "object_ids",
					Description: `The object IDs of the roles.`,
				},
				resource.Attribute{
					Name:        "template_ids",
					Description: `The template IDs of the roles.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `A list of users. Each ` + "`" + `role` + "`" + ` object provides the attributes documented below. --- ` + "`" + `role` + "`" + ` object exports the following:`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name of the directory role.`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `The template ID of the directory role.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the directory role.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The object ID of the directory role.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_domains",
			Category:         "Domains",
			ShortDescription: ``,
			Description: `

Use this data source to access information about existing Domains within Azure Active Directory.

`,
			Keywords: []string{
				"domains",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "admin_managed",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to only return domains whose DNS is managed by Microsoft 365. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "include_unverified",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` if unverified Azure AD domains should be included. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "only_default",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to only return the default domain.`,
				},
				resource.Attribute{
					Name:        "only_initial",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to only return the initial domain, which is your primary Azure Active Directory tenant domain. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "only_root",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to only return verified root domains. Excludes subdomains and unverified domains.`,
				},
				resource.Attribute{
					Name:        "supports_services",
					Description: `(Optional) A list of supported services that must be supported by a domain. Possible values include ` + "`" + `Email` + "`" + `, ` + "`" + `Sharepoint` + "`" + `, ` + "`" + `EmailInternalRelayOnly` + "`" + `, ` + "`" + `OfficeCommunicationsOnline` + "`" + `, ` + "`" + `SharePointDefaultDomain` + "`" + `, ` + "`" + `FullRedelegation` + "`" + `, ` + "`" + `SharePointPublic` + "`" + `, ` + "`" + `OrgIdAuthentication` + "`" + `, ` + "`" + `Yammer` + "`" + ` and ` + "`" + `Intune` + "`" + `. ->`,
				},
				resource.Attribute{
					Name:        "domains",
					Description: `A list of tenant domains. Each ` + "`" + `domain` + "`" + ` object provides the attributes documented below. --- ` + "`" + `domain` + "`" + ` object exports the following:`,
				},
				resource.Attribute{
					Name:        "admin_managed",
					Description: `Whether the DNS for the domain is managed by Microsoft 365.`,
				},
				resource.Attribute{
					Name:        "authentication_type",
					Description: `The authentication type of the domain. Possible values include ` + "`" + `Managed` + "`" + ` or ` + "`" + `Federated` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `The name of the domain.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `Whether this is the default domain that is used for user creation.`,
				},
				resource.Attribute{
					Name:        "initial",
					Description: `Whether this is the initial domain created by Azure Active Directory.`,
				},
				resource.Attribute{
					Name:        "root",
					Description: `Whether the domain is a verified root domain (not a subdomain).`,
				},
				resource.Attribute{
					Name:        "verified",
					Description: `Whether the domain has completed domain ownership verification.`,
				},
				resource.Attribute{
					Name:        "supported_services",
					Description: `A list of capabilities / services supported by the domain. Possible values include ` + "`" + `Email` + "`" + `, ` + "`" + `Sharepoint` + "`" + `, ` + "`" + `EmailInternalRelayOnly` + "`" + `, ` + "`" + `OfficeCommunicationsOnline` + "`" + `, ` + "`" + `SharePointDefaultDomain` + "`" + `, ` + "`" + `FullRedelegation` + "`" + `, ` + "`" + `SharePointPublic` + "`" + `, ` + "`" + `OrgIdAuthentication` + "`" + `, ` + "`" + `Yammer` + "`" + ` and ` + "`" + `Intune` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "domains",
					Description: `A list of tenant domains. Each ` + "`" + `domain` + "`" + ` object provides the attributes documented below. --- ` + "`" + `domain` + "`" + ` object exports the following:`,
				},
				resource.Attribute{
					Name:        "admin_managed",
					Description: `Whether the DNS for the domain is managed by Microsoft 365.`,
				},
				resource.Attribute{
					Name:        "authentication_type",
					Description: `The authentication type of the domain. Possible values include ` + "`" + `Managed` + "`" + ` or ` + "`" + `Federated` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `The name of the domain.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `Whether this is the default domain that is used for user creation.`,
				},
				resource.Attribute{
					Name:        "initial",
					Description: `Whether this is the initial domain created by Azure Active Directory.`,
				},
				resource.Attribute{
					Name:        "root",
					Description: `Whether the domain is a verified root domain (not a subdomain).`,
				},
				resource.Attribute{
					Name:        "verified",
					Description: `Whether the domain has completed domain ownership verification.`,
				},
				resource.Attribute{
					Name:        "supported_services",
					Description: `A list of capabilities / services supported by the domain. Possible values include ` + "`" + `Email` + "`" + `, ` + "`" + `Sharepoint` + "`" + `, ` + "`" + `EmailInternalRelayOnly` + "`" + `, ` + "`" + `OfficeCommunicationsOnline` + "`" + `, ` + "`" + `SharePointDefaultDomain` + "`" + `, ` + "`" + `FullRedelegation` + "`" + `, ` + "`" + `SharePointPublic` + "`" + `, ` + "`" + `OrgIdAuthentication` + "`" + `, ` + "`" + `Yammer` + "`" + ` and ` + "`" + `Intune` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_group",
			Category:         "Groups",
			ShortDescription: ``,
			Description: `

Gets information about an Azure Active Directory group.

`,
			Keywords: []string{
				"groups",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name for the group.`,
				},
				resource.Attribute{
					Name:        "mail_enabled",
					Description: `(Optional) Whether the group is mail-enabled.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `(Optional) Specifies the object ID of the group.`,
				},
				resource.Attribute{
					Name:        "security_enabled",
					Description: `(Optional) Whether the group is a security group. ~> One of ` + "`" + `display_name` + "`" + ` or ` + "`" + `object_id` + "`" + ` must be specified. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "assignable_to_role",
					Description: `Indicates whether this group can be assigned to an Azure Active Directory role.`,
				},
				resource.Attribute{
					Name:        "auto_subscribe_new_members",
					Description: `Indicates whether new members added to the group will be auto-subscribed to receive email notifications. Only set for Unified groups.`,
				},
				resource.Attribute{
					Name:        "behaviors",
					Description: `A list of behaviors for a Microsoft 365 group, such as ` + "`" + `AllowOnlyMembersToPost` + "`" + `, ` + "`" + `HideGroupInOutlook` + "`" + `, ` + "`" + `SubscribeNewGroupMembers` + "`" + ` and ` + "`" + `WelcomeEmailDisabled` + "`" + `. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for more details.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The optional description of the group.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name for the group.`,
				},
				resource.Attribute{
					Name:        "dynamic_membership",
					Description: `A ` + "`" + `dynamic_membership` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "external_senders_allowed",
					Description: `Indicates whether people external to the organization can send messages to the group. Only set for Unified groups.`,
				},
				resource.Attribute{
					Name:        "hide_from_address_lists",
					Description: `Indicates whether the group is displayed in certain parts of the Outlook user interface: in the Address Book, in address lists for selecting message recipients, and in the Browse Groups dialog for searching groups. Only set for Unified groups.`,
				},
				resource.Attribute{
					Name:        "hide_from_outlook_clients",
					Description: `Indicates whether the group is displayed in Outlook clients, such as Outlook for Windows and Outlook on the web. Only set for Unified groups.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The object ID of the group.`,
				},
				resource.Attribute{
					Name:        "mail",
					Description: `The SMTP address for the group.`,
				},
				resource.Attribute{
					Name:        "mail_enabled",
					Description: `Whether the group is mail-enabled.`,
				},
				resource.Attribute{
					Name:        "mail_nickname",
					Description: `The mail alias for the group, unique in the organisation.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `List of object IDs of the group members.`,
				},
				resource.Attribute{
					Name:        "onpremises_domain_name",
					Description: `The on-premises FQDN, also called dnsDomainName, synchronised from the on-premises directory when Azure AD Connect is used.`,
				},
				resource.Attribute{
					Name:        "onpremises_group_type",
					Description: `The on-premises group type that the AAD group will be written as, when writeback is enabled. Possible values are ` + "`" + `UniversalDistributionGroup` + "`" + `, ` + "`" + `UniversalMailEnabledSecurityGroup` + "`" + `, or ` + "`" + `UniversalSecurityGroup` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "onpremises_netbios_name",
					Description: `The on-premises NetBIOS name, synchronised from the on-premises directory when Azure AD Connect is used.`,
				},
				resource.Attribute{
					Name:        "onpremises_sam_account_name",
					Description: `The on-premises SAM account name, synchronised from the on-premises directory when Azure AD Connect is used.`,
				},
				resource.Attribute{
					Name:        "onpremises_security_identifier",
					Description: `The on-premises security identifier (SID), synchronised from the on-premises directory when Azure AD Connect is used.`,
				},
				resource.Attribute{
					Name:        "onpremises_sync_enabled",
					Description: `Whether this group is synchronised from an on-premises directory (` + "`" + `true` + "`" + `), no longer synchronised (` + "`" + `false` + "`" + `), or has never been synchronised (` + "`" + `null` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "owners",
					Description: `List of object IDs of the group owners.`,
				},
				resource.Attribute{
					Name:        "preferred_language",
					Description: `The preferred language for a Microsoft 365 group, in ISO 639-1 notation.`,
				},
				resource.Attribute{
					Name:        "provisioning_options",
					Description: `A list of provisioning options for a Microsoft 365 group, such as ` + "`" + `Team` + "`" + `. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for details.`,
				},
				resource.Attribute{
					Name:        "proxy_addresses",
					Description: `List of email addresses for the group that direct to the same group mailbox.`,
				},
				resource.Attribute{
					Name:        "security_enabled",
					Description: `Whether the group is a security group.`,
				},
				resource.Attribute{
					Name:        "theme",
					Description: `The colour theme for a Microsoft 365 group. Possible values are ` + "`" + `Blue` + "`" + `, ` + "`" + `Green` + "`" + `, ` + "`" + `Orange` + "`" + `, ` + "`" + `Pink` + "`" + `, ` + "`" + `Purple` + "`" + `, ` + "`" + `Red` + "`" + ` or ` + "`" + `Teal` + "`" + `. When no theme is set, the value is ` + "`" + `null` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "types",
					Description: `A list of group types configured for the group. Supported values are ` + "`" + `DynamicMembership` + "`" + `, which denotes a group with dynamic membership, and ` + "`" + `Unified` + "`" + `, which specifies a Microsoft 365 group.`,
				},
				resource.Attribute{
					Name:        "visibility",
					Description: `The group join policy and group content visibility. Possible values are ` + "`" + `Private` + "`" + `, ` + "`" + `Public` + "`" + `, or ` + "`" + `Hiddenmembership` + "`" + `. Only Microsoft 365 groups can have ` + "`" + `Hiddenmembership` + "`" + ` visibility.`,
				},
				resource.Attribute{
					Name:        "writeback_enabled",
					Description: `Whether the group will be written back to the configured on-premises Active Directory when Azure AD Connect is used. --- ` + "`" + `dynamic_membership` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Whether rule processing is "On" (true) or "Paused" (false).`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `The rule that determines membership of this group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "assignable_to_role",
					Description: `Indicates whether this group can be assigned to an Azure Active Directory role.`,
				},
				resource.Attribute{
					Name:        "auto_subscribe_new_members",
					Description: `Indicates whether new members added to the group will be auto-subscribed to receive email notifications. Only set for Unified groups.`,
				},
				resource.Attribute{
					Name:        "behaviors",
					Description: `A list of behaviors for a Microsoft 365 group, such as ` + "`" + `AllowOnlyMembersToPost` + "`" + `, ` + "`" + `HideGroupInOutlook` + "`" + `, ` + "`" + `SubscribeNewGroupMembers` + "`" + ` and ` + "`" + `WelcomeEmailDisabled` + "`" + `. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for more details.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The optional description of the group.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name for the group.`,
				},
				resource.Attribute{
					Name:        "dynamic_membership",
					Description: `A ` + "`" + `dynamic_membership` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "external_senders_allowed",
					Description: `Indicates whether people external to the organization can send messages to the group. Only set for Unified groups.`,
				},
				resource.Attribute{
					Name:        "hide_from_address_lists",
					Description: `Indicates whether the group is displayed in certain parts of the Outlook user interface: in the Address Book, in address lists for selecting message recipients, and in the Browse Groups dialog for searching groups. Only set for Unified groups.`,
				},
				resource.Attribute{
					Name:        "hide_from_outlook_clients",
					Description: `Indicates whether the group is displayed in Outlook clients, such as Outlook for Windows and Outlook on the web. Only set for Unified groups.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The object ID of the group.`,
				},
				resource.Attribute{
					Name:        "mail",
					Description: `The SMTP address for the group.`,
				},
				resource.Attribute{
					Name:        "mail_enabled",
					Description: `Whether the group is mail-enabled.`,
				},
				resource.Attribute{
					Name:        "mail_nickname",
					Description: `The mail alias for the group, unique in the organisation.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `List of object IDs of the group members.`,
				},
				resource.Attribute{
					Name:        "onpremises_domain_name",
					Description: `The on-premises FQDN, also called dnsDomainName, synchronised from the on-premises directory when Azure AD Connect is used.`,
				},
				resource.Attribute{
					Name:        "onpremises_group_type",
					Description: `The on-premises group type that the AAD group will be written as, when writeback is enabled. Possible values are ` + "`" + `UniversalDistributionGroup` + "`" + `, ` + "`" + `UniversalMailEnabledSecurityGroup` + "`" + `, or ` + "`" + `UniversalSecurityGroup` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "onpremises_netbios_name",
					Description: `The on-premises NetBIOS name, synchronised from the on-premises directory when Azure AD Connect is used.`,
				},
				resource.Attribute{
					Name:        "onpremises_sam_account_name",
					Description: `The on-premises SAM account name, synchronised from the on-premises directory when Azure AD Connect is used.`,
				},
				resource.Attribute{
					Name:        "onpremises_security_identifier",
					Description: `The on-premises security identifier (SID), synchronised from the on-premises directory when Azure AD Connect is used.`,
				},
				resource.Attribute{
					Name:        "onpremises_sync_enabled",
					Description: `Whether this group is synchronised from an on-premises directory (` + "`" + `true` + "`" + `), no longer synchronised (` + "`" + `false` + "`" + `), or has never been synchronised (` + "`" + `null` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "owners",
					Description: `List of object IDs of the group owners.`,
				},
				resource.Attribute{
					Name:        "preferred_language",
					Description: `The preferred language for a Microsoft 365 group, in ISO 639-1 notation.`,
				},
				resource.Attribute{
					Name:        "provisioning_options",
					Description: `A list of provisioning options for a Microsoft 365 group, such as ` + "`" + `Team` + "`" + `. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for details.`,
				},
				resource.Attribute{
					Name:        "proxy_addresses",
					Description: `List of email addresses for the group that direct to the same group mailbox.`,
				},
				resource.Attribute{
					Name:        "security_enabled",
					Description: `Whether the group is a security group.`,
				},
				resource.Attribute{
					Name:        "theme",
					Description: `The colour theme for a Microsoft 365 group. Possible values are ` + "`" + `Blue` + "`" + `, ` + "`" + `Green` + "`" + `, ` + "`" + `Orange` + "`" + `, ` + "`" + `Pink` + "`" + `, ` + "`" + `Purple` + "`" + `, ` + "`" + `Red` + "`" + ` or ` + "`" + `Teal` + "`" + `. When no theme is set, the value is ` + "`" + `null` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "types",
					Description: `A list of group types configured for the group. Supported values are ` + "`" + `DynamicMembership` + "`" + `, which denotes a group with dynamic membership, and ` + "`" + `Unified` + "`" + `, which specifies a Microsoft 365 group.`,
				},
				resource.Attribute{
					Name:        "visibility",
					Description: `The group join policy and group content visibility. Possible values are ` + "`" + `Private` + "`" + `, ` + "`" + `Public` + "`" + `, or ` + "`" + `Hiddenmembership` + "`" + `. Only Microsoft 365 groups can have ` + "`" + `Hiddenmembership` + "`" + ` visibility.`,
				},
				resource.Attribute{
					Name:        "writeback_enabled",
					Description: `Whether the group will be written back to the configured on-premises Active Directory when Azure AD Connect is used. --- ` + "`" + `dynamic_membership` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Whether rule processing is "On" (true) or "Paused" (false).`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `The rule that determines membership of this group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_groups",
			Category:         "Groups",
			ShortDescription: ``,
			Description: `

Gets Object IDs or Display Names for multiple Azure Active Directory groups.

`,
			Keywords: []string{
				"groups",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_names",
					Description: `(Optional) The display names of the groups.`,
				},
				resource.Attribute{
					Name:        "display_name_prefix",
					Description: `(Optional) A common display name prefix to match when returning groups.`,
				},
				resource.Attribute{
					Name:        "ignore_missing",
					Description: `(Optional) Ignore missing groups and return groups that were found. The data source will still fail if no groups are found. Cannot be specified with ` + "`" + `return_all` + "`" + `. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mail_enabled",
					Description: `(Optional) Whether the returned groups should be mail-enabled. By itself this does not exclude security-enabled groups. Setting this to ` + "`" + `true` + "`" + ` ensures all groups are mail-enabled, and setting to ` + "`" + `false` + "`" + ` ensures that all groups are _not_ mail-enabled. To ignore this filter, omit the property or set it to null. Cannot be specified together with ` + "`" + `object_ids` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "object_ids",
					Description: `(Optional) The object IDs of the groups.`,
				},
				resource.Attribute{
					Name:        "return_all",
					Description: `(Optional) A flag to denote if all groups should be fetched and returned. Cannot be specified wth ` + "`" + `ignore_missing` + "`" + `. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "security_enabled",
					Description: `(Optional) Whether the returned groups should be security-enabled. By itself this does not exclude mail-enabled groups. Setting this to ` + "`" + `true` + "`" + ` ensures all groups are security-enabled, and setting to ` + "`" + `false` + "`" + ` ensures that all groups are _not_ security-enabled. To ignore this filter, omit the property or set it to null. Cannot be specified together with ` + "`" + `object_ids` + "`" + `. ~> One of ` + "`" + `display_names` + "`" + `, ` + "`" + `display_name_prefix` + "`" + `, ` + "`" + `object_ids` + "`" + ` or ` + "`" + `return_all` + "`" + ` should be specified. Either ` + "`" + `display_name` + "`" + ` or ` + "`" + `object_ids` + "`" + ` _may_ be specified as an empty list, in which case no results will be returned. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "display_names",
					Description: `The display names of the groups.`,
				},
				resource.Attribute{
					Name:        "object_ids",
					Description: `The object IDs of the groups.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "display_names",
					Description: `The display names of the groups.`,
				},
				resource.Attribute{
					Name:        "object_ids",
					Description: `The object IDs of the groups.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_service_principal",
			Category:         "Service Principals",
			ShortDescription: ``,
			Description: `

Gets information about an existing service principal associated with an application within Azure Active Directory.

`,
			Keywords: []string{
				"service",
				"principals",
				"principal",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "application_id",
					Description: `(Optional) The application ID (client ID) of the application associated with this service principal.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of the application associated with this service principal.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `(Optional) The object ID of the service principal. ~> One of ` + "`" + `application_id` + "`" + `, ` + "`" + `display_name` + "`" + ` or ` + "`" + `object_id` + "`" + ` must be specified. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "account_enabled",
					Description: `Whether or not the service principal account is enabled.`,
				},
				resource.Attribute{
					Name:        "alternative_names",
					Description: `A list of alternative names, used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities.`,
				},
				resource.Attribute{
					Name:        "application_id",
					Description: `The application ID (client ID) of the application associated with this service principal.`,
				},
				resource.Attribute{
					Name:        "app_role_assignment_required",
					Description: `Whether this service principal requires an app role assignment to a user or group before Azure AD will issue a user or access token to the application.`,
				},
				resource.Attribute{
					Name:        "app_role_ids",
					Description: `A mapping of app role values to app role IDs, as published by the associated application, intended to be useful when referencing app roles in other resources in your configuration.`,
				},
				resource.Attribute{
					Name:        "app_roles",
					Description: `A list of app roles published by the associated application, as documented below. For more information [official documentation](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).`,
				},
				resource.Attribute{
					Name:        "application_tenant_id",
					Description: `The tenant ID where the associated application is registered.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the service principal provided for internal end-users.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name of the application associated with this service principal.`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `A ` + "`" + `features` + "`" + ` block as described below.`,
				},
				resource.Attribute{
					Name:        "homepage_url",
					Description: `Home page or landing page of the associated application.`,
				},
				resource.Attribute{
					Name:        "login_url",
					Description: `The URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps.`,
				},
				resource.Attribute{
					Name:        "logout_url",
					Description: `The URL that will be used by Microsoft's authorization service to logout an user using OpenId Connect front-channel, back-channel or SAML logout protocols, taken from the associated application.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `A free text field to capture information about the service principal, typically used for operational purposes.`,
				},
				resource.Attribute{
					Name:        "notification_email_addresses",
					Description: `A list of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The object ID of the service principal.`,
				},
				resource.Attribute{
					Name:        "oauth2_permission_scope_ids",
					Description: `A mapping of OAuth2.0 permission scope values to scope IDs, as exposed by the associated application, intended to be useful when referencing permission scopes in other resources in your configuration.`,
				},
				resource.Attribute{
					Name:        "oauth2_permission_scopes",
					Description: `A collection of OAuth 2.0 delegated permissions exposed by the associated application. Each permission is covered by an ` + "`" + `oauth2_permission_scopes` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "preferred_single_sign_on_mode",
					Description: `The single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps.`,
				},
				resource.Attribute{
					Name:        "redirect_uris",
					Description: `A list of URLs where user tokens are sent for sign-in with the associated application, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent for the associated application.`,
				},
				resource.Attribute{
					Name:        "saml_metadata_url",
					Description: `The URL where the service exposes SAML metadata for federation.`,
				},
				resource.Attribute{
					Name:        "saml_single_sign_on",
					Description: `A ` + "`" + `saml_single_sign_on` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "service_principal_names",
					Description: `A list of identifier URI(s), copied over from the associated application.`,
				},
				resource.Attribute{
					Name:        "sign_in_audience",
					Description: `The Microsoft account types that are supported for the associated application. Possible values include ` + "`" + `AzureADMyOrg` + "`" + `, ` + "`" + `AzureADMultipleOrgs` + "`" + `, ` + "`" + `AzureADandPersonalMicrosoftAccount` + "`" + ` or ` + "`" + `PersonalMicrosoftAccount` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of tags applied to the service principal.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Identifies whether the service principal represents an application or a managed identity. Possible values include ` + "`" + `Application` + "`" + ` or ` + "`" + `ManagedIdentity` + "`" + `. --- ` + "`" + `app_roles` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "allowed_member_types",
					Description: `Specifies whether this app role definition can be assigned to users and groups, or to other applications (that are accessing this application in daemon service scenarios). Possible values are: ` + "`" + `User` + "`" + ` and ` + "`" + `Application` + "`" + `, or both.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Permission help text that appears in the admin app assignment and consent experiences.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Display name for the permission that appears in the admin consent and app assignment experiences.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier of the app role.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `Determines if the app role is enabled.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Specifies the value of the roles claim that the application should expect in the authentication and access tokens. --- ` + "`" + `features` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "custom_single_sign_on_app",
					Description: `Whether this service principal represents a custom SAML application.`,
				},
				resource.Attribute{
					Name:        "enterprise_application",
					Description: `Whether this service principal represents an Enterprise Application.`,
				},
				resource.Attribute{
					Name:        "gallery_application",
					Description: `Whether this service principal represents a gallery application.`,
				},
				resource.Attribute{
					Name:        "visible_to_users",
					Description: `Whether this app is visible to users in My Apps and Office 365 Launcher. --- ` + "`" + `oauth2_permission_scopes` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "admin_consent_description",
					Description: `Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.`,
				},
				resource.Attribute{
					Name:        "admin_consent_display_name",
					Description: `Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Determines if the permission scope is enabled.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier of the delegated permission. Must be a valid UUID.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Possible values are ` + "`" + `User` + "`" + ` or ` + "`" + `Admin` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user_consent_description",
					Description: `Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.`,
				},
				resource.Attribute{
					Name:        "user_consent_display_name",
					Description: `Display name for the delegated permission that appears in the end user consent experience.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value that is used for the ` + "`" + `scp` + "`" + ` claim in OAuth 2.0 access tokens. --- ` + "`" + `saml_single_sign_on` + "`" + ` exports the following:`,
				},
				resource.Attribute{
					Name:        "relay_state",
					Description: `The relative URI the service provider would redirect to after completion of the single sign-on flow.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "account_enabled",
					Description: `Whether or not the service principal account is enabled.`,
				},
				resource.Attribute{
					Name:        "alternative_names",
					Description: `A list of alternative names, used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities.`,
				},
				resource.Attribute{
					Name:        "application_id",
					Description: `The application ID (client ID) of the application associated with this service principal.`,
				},
				resource.Attribute{
					Name:        "app_role_assignment_required",
					Description: `Whether this service principal requires an app role assignment to a user or group before Azure AD will issue a user or access token to the application.`,
				},
				resource.Attribute{
					Name:        "app_role_ids",
					Description: `A mapping of app role values to app role IDs, as published by the associated application, intended to be useful when referencing app roles in other resources in your configuration.`,
				},
				resource.Attribute{
					Name:        "app_roles",
					Description: `A list of app roles published by the associated application, as documented below. For more information [official documentation](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).`,
				},
				resource.Attribute{
					Name:        "application_tenant_id",
					Description: `The tenant ID where the associated application is registered.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the service principal provided for internal end-users.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name of the application associated with this service principal.`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `A ` + "`" + `features` + "`" + ` block as described below.`,
				},
				resource.Attribute{
					Name:        "homepage_url",
					Description: `Home page or landing page of the associated application.`,
				},
				resource.Attribute{
					Name:        "login_url",
					Description: `The URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps.`,
				},
				resource.Attribute{
					Name:        "logout_url",
					Description: `The URL that will be used by Microsoft's authorization service to logout an user using OpenId Connect front-channel, back-channel or SAML logout protocols, taken from the associated application.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `A free text field to capture information about the service principal, typically used for operational purposes.`,
				},
				resource.Attribute{
					Name:        "notification_email_addresses",
					Description: `A list of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The object ID of the service principal.`,
				},
				resource.Attribute{
					Name:        "oauth2_permission_scope_ids",
					Description: `A mapping of OAuth2.0 permission scope values to scope IDs, as exposed by the associated application, intended to be useful when referencing permission scopes in other resources in your configuration.`,
				},
				resource.Attribute{
					Name:        "oauth2_permission_scopes",
					Description: `A collection of OAuth 2.0 delegated permissions exposed by the associated application. Each permission is covered by an ` + "`" + `oauth2_permission_scopes` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "preferred_single_sign_on_mode",
					Description: `The single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps.`,
				},
				resource.Attribute{
					Name:        "redirect_uris",
					Description: `A list of URLs where user tokens are sent for sign-in with the associated application, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent for the associated application.`,
				},
				resource.Attribute{
					Name:        "saml_metadata_url",
					Description: `The URL where the service exposes SAML metadata for federation.`,
				},
				resource.Attribute{
					Name:        "saml_single_sign_on",
					Description: `A ` + "`" + `saml_single_sign_on` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "service_principal_names",
					Description: `A list of identifier URI(s), copied over from the associated application.`,
				},
				resource.Attribute{
					Name:        "sign_in_audience",
					Description: `The Microsoft account types that are supported for the associated application. Possible values include ` + "`" + `AzureADMyOrg` + "`" + `, ` + "`" + `AzureADMultipleOrgs` + "`" + `, ` + "`" + `AzureADandPersonalMicrosoftAccount` + "`" + ` or ` + "`" + `PersonalMicrosoftAccount` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of tags applied to the service principal.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Identifies whether the service principal represents an application or a managed identity. Possible values include ` + "`" + `Application` + "`" + ` or ` + "`" + `ManagedIdentity` + "`" + `. --- ` + "`" + `app_roles` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "allowed_member_types",
					Description: `Specifies whether this app role definition can be assigned to users and groups, or to other applications (that are accessing this application in daemon service scenarios). Possible values are: ` + "`" + `User` + "`" + ` and ` + "`" + `Application` + "`" + `, or both.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Permission help text that appears in the admin app assignment and consent experiences.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Display name for the permission that appears in the admin consent and app assignment experiences.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier of the app role.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `Determines if the app role is enabled.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Specifies the value of the roles claim that the application should expect in the authentication and access tokens. --- ` + "`" + `features` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "custom_single_sign_on_app",
					Description: `Whether this service principal represents a custom SAML application.`,
				},
				resource.Attribute{
					Name:        "enterprise_application",
					Description: `Whether this service principal represents an Enterprise Application.`,
				},
				resource.Attribute{
					Name:        "gallery_application",
					Description: `Whether this service principal represents a gallery application.`,
				},
				resource.Attribute{
					Name:        "visible_to_users",
					Description: `Whether this app is visible to users in My Apps and Office 365 Launcher. --- ` + "`" + `oauth2_permission_scopes` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "admin_consent_description",
					Description: `Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.`,
				},
				resource.Attribute{
					Name:        "admin_consent_display_name",
					Description: `Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Determines if the permission scope is enabled.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier of the delegated permission. Must be a valid UUID.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Possible values are ` + "`" + `User` + "`" + ` or ` + "`" + `Admin` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user_consent_description",
					Description: `Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.`,
				},
				resource.Attribute{
					Name:        "user_consent_display_name",
					Description: `Display name for the delegated permission that appears in the end user consent experience.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value that is used for the ` + "`" + `scp` + "`" + ` claim in OAuth 2.0 access tokens. --- ` + "`" + `saml_single_sign_on` + "`" + ` exports the following:`,
				},
				resource.Attribute{
					Name:        "relay_state",
					Description: `The relative URI the service provider would redirect to after completion of the single sign-on flow.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_service_principals",
			Category:         "Service Principals",
			ShortDescription: ``,
			Description: `

Gets basic information for multiple Azure Active Directory service principals.

`,
			Keywords: []string{
				"service",
				"principals",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "application_ids",
					Description: `(Optional) A list of application IDs (client IDs) of the applications associated with the service principals.`,
				},
				resource.Attribute{
					Name:        "display_names",
					Description: `(Optional) A list of display names of the applications associated with the service principals.`,
				},
				resource.Attribute{
					Name:        "ignore_missing",
					Description: `(Optional) Ignore missing service principals and return all service principals that are found. The data source will still fail if no service principals are found. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "object_ids",
					Description: `(Optional) The object IDs of the service principals.`,
				},
				resource.Attribute{
					Name:        "return_all",
					Description: `(Optional) When ` + "`" + `true` + "`" + `, the data source will return all service principals. Cannot be used with ` + "`" + `ignore_missing` + "`" + `. Defaults to false. ~> Either ` + "`" + `return_all` + "`" + `, or one of ` + "`" + `application_ids` + "`" + `, ` + "`" + `display_names` + "`" + ` or ` + "`" + `object_ids` + "`" + ` must be specified. These _may_ be specified as an empty list, in which case no results will be returned. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "application_ids",
					Description: `A list of application IDs (client IDs) of the applications associated with the service principals.`,
				},
				resource.Attribute{
					Name:        "display_names",
					Description: `A list of display names of the applications associated with the service principals.`,
				},
				resource.Attribute{
					Name:        "object_ids",
					Description: `The object IDs of the service principals.`,
				},
				resource.Attribute{
					Name:        "service_principals",
					Description: `A list of service principals. Each ` + "`" + `service_principal` + "`" + ` object provides the attributes documented below. --- ` + "`" + `service_principal` + "`" + ` object exports the following:`,
				},
				resource.Attribute{
					Name:        "account_enabled",
					Description: `Whether or not the service principal account is enabled.`,
				},
				resource.Attribute{
					Name:        "app_role_assignment_required",
					Description: `Whether this service principal requires an app role assignment to a user or group before Azure AD will issue a user or access token to the application.`,
				},
				resource.Attribute{
					Name:        "application_id",
					Description: `The application ID (client ID) of the application associated with this service principal.`,
				},
				resource.Attribute{
					Name:        "application_tenant_id",
					Description: `The tenant ID where the associated application is registered.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name of the application associated with this service principal.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The object ID of the service principal.`,
				},
				resource.Attribute{
					Name:        "preferred_single_sign_on_mode",
					Description: `The single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps.`,
				},
				resource.Attribute{
					Name:        "saml_metadata_url",
					Description: `The URL where the service exposes SAML metadata for federation.`,
				},
				resource.Attribute{
					Name:        "service_principal_names",
					Description: `A list of identifier URI(s), copied over from the associated application.`,
				},
				resource.Attribute{
					Name:        "sign_in_audience",
					Description: `The Microsoft account types that are supported for the associated application. Possible values include ` + "`" + `AzureADMyOrg` + "`" + `, ` + "`" + `AzureADMultipleOrgs` + "`" + `, ` + "`" + `AzureADandPersonalMicrosoftAccount` + "`" + ` or ` + "`" + `PersonalMicrosoftAccount` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of tags applied to the service principal.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Identifies whether the service principal represents an application or a managed identity. Possible values include ` + "`" + `Application` + "`" + ` or ` + "`" + `ManagedIdentity` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "application_ids",
					Description: `A list of application IDs (client IDs) of the applications associated with the service principals.`,
				},
				resource.Attribute{
					Name:        "display_names",
					Description: `A list of display names of the applications associated with the service principals.`,
				},
				resource.Attribute{
					Name:        "object_ids",
					Description: `The object IDs of the service principals.`,
				},
				resource.Attribute{
					Name:        "service_principals",
					Description: `A list of service principals. Each ` + "`" + `service_principal` + "`" + ` object provides the attributes documented below. --- ` + "`" + `service_principal` + "`" + ` object exports the following:`,
				},
				resource.Attribute{
					Name:        "account_enabled",
					Description: `Whether or not the service principal account is enabled.`,
				},
				resource.Attribute{
					Name:        "app_role_assignment_required",
					Description: `Whether this service principal requires an app role assignment to a user or group before Azure AD will issue a user or access token to the application.`,
				},
				resource.Attribute{
					Name:        "application_id",
					Description: `The application ID (client ID) of the application associated with this service principal.`,
				},
				resource.Attribute{
					Name:        "application_tenant_id",
					Description: `The tenant ID where the associated application is registered.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name of the application associated with this service principal.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The object ID of the service principal.`,
				},
				resource.Attribute{
					Name:        "preferred_single_sign_on_mode",
					Description: `The single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps.`,
				},
				resource.Attribute{
					Name:        "saml_metadata_url",
					Description: `The URL where the service exposes SAML metadata for federation.`,
				},
				resource.Attribute{
					Name:        "service_principal_names",
					Description: `A list of identifier URI(s), copied over from the associated application.`,
				},
				resource.Attribute{
					Name:        "sign_in_audience",
					Description: `The Microsoft account types that are supported for the associated application. Possible values include ` + "`" + `AzureADMyOrg` + "`" + `, ` + "`" + `AzureADMultipleOrgs` + "`" + `, ` + "`" + `AzureADandPersonalMicrosoftAccount` + "`" + ` or ` + "`" + `PersonalMicrosoftAccount` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of tags applied to the service principal.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Identifies whether the service principal represents an application or a managed identity. Possible values include ` + "`" + `Application` + "`" + ` or ` + "`" + `ManagedIdentity` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_user",
			Category:         "Users",
			ShortDescription: ``,
			Description: `

Gets information about an Azure Active Directory user.

`,
			Keywords: []string{
				"users",
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "mail",
					Description: `(Optional) The SMTP address for the user.`,
				},
				resource.Attribute{
					Name:        "mail_nickname",
					Description: `(Optional) The email alias of the user.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `(Optional) The object ID of the user.`,
				},
				resource.Attribute{
					Name:        "user_principal_name",
					Description: `(Optional) The user principal name (UPN) of the user. ~> One of ` + "`" + `user_principal_name` + "`" + `, ` + "`" + `object_id` + "`" + `, ` + "`" + `mail` + "`" + ` or ` + "`" + `mail_nickname` + "`" + ` must be specified. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "account_enabled",
					Description: `Whether or not the account is enabled.`,
				},
				resource.Attribute{
					Name:        "age_group",
					Description: `The age group of the user. Supported values are ` + "`" + `Adult` + "`" + `, ` + "`" + `NotAdult` + "`" + ` and ` + "`" + `Minor` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "business_phones",
					Description: `A list of telephone numbers for the user.`,
				},
				resource.Attribute{
					Name:        "city",
					Description: `The city in which the user is located.`,
				},
				resource.Attribute{
					Name:        "company_name",
					Description: `The company name which the user is associated. This property can be useful for describing the company that an external user comes from.`,
				},
				resource.Attribute{
					Name:        "consent_provided_for_minor",
					Description: `Whether consent has been obtained for minors. Supported values are ` + "`" + `Granted` + "`" + `, ` + "`" + `Denied` + "`" + ` and ` + "`" + `NotRequired` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `The country/region in which the user is located, e.g. ` + "`" + `US` + "`" + ` or ` + "`" + `UK` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cost_center",
					Description: `The cost center associated with the user.`,
				},
				resource.Attribute{
					Name:        "creation_type",
					Description: `Indicates whether the user account was created as a regular school or work account (` + "`" + `null` + "`" + `), an external account (` + "`" + `Invitation` + "`" + `), a local account for an Azure Active Directory B2C tenant (` + "`" + `LocalAccount` + "`" + `) or self-service sign-up using email verification (` + "`" + `EmailVerified` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "department",
					Description: `The name for the department in which the user works.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name of the user.`,
				},
				resource.Attribute{
					Name:        "division",
					Description: `The name of the division in which the user works.`,
				},
				resource.Attribute{
					Name:        "employee_id",
					Description: `The employee identifier assigned to the user by the organisation.`,
				},
				resource.Attribute{
					Name:        "employee_type",
					Description: `Captures enterprise worker type. For example, Employee, Contractor, Consultant, or Vendor.`,
				},
				resource.Attribute{
					Name:        "external_user_state",
					Description: `For an external user invited to the tenant, this property represents the invited user's invitation status. Possible values are ` + "`" + `PendingAcceptance` + "`" + ` or ` + "`" + `Accepted` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fax_number",
					Description: `The fax number of the user.`,
				},
				resource.Attribute{
					Name:        "given_name",
					Description: `The given name (first name) of the user.`,
				},
				resource.Attribute{
					Name:        "im_addresses",
					Description: `A list of instant message voice over IP (VOIP) session initiation protocol (SIP) addresses for the user.`,
				},
				resource.Attribute{
					Name:        "job_title",
					Description: `The user’s job title.`,
				},
				resource.Attribute{
					Name:        "mail",
					Description: `The SMTP address for the user.`,
				},
				resource.Attribute{
					Name:        "mail_nickname",
					Description: `The email alias of the user.`,
				},
				resource.Attribute{
					Name:        "manager_id",
					Description: `The object ID of the user's manager.`,
				},
				resource.Attribute{
					Name:        "mobile_phone",
					Description: `The primary cellular telephone number for the user.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The object ID of the user.`,
				},
				resource.Attribute{
					Name:        "office_location",
					Description: `The office location in the user's place of business.`,
				},
				resource.Attribute{
					Name:        "onpremises_distinguished_name",
					Description: `The on-premises distinguished name (DN) of the user, synchronised from the on-premises directory when Azure AD Connect is used.`,
				},
				resource.Attribute{
					Name:        "onpremises_domain_name",
					Description: `The on-premises FQDN, also called dnsDomainName, synchronised from the on-premises directory when Azure AD Connect is used.`,
				},
				resource.Attribute{
					Name:        "onpremises_immutable_id",
					Description: `The value used to associate an on-premise Active Directory user account with their Azure AD user object.`,
				},
				resource.Attribute{
					Name:        "onpremises_sam_account_name",
					Description: `The on-premise SAM account name of the user.`,
				},
				resource.Attribute{
					Name:        "onpremises_security_identifier",
					Description: `The on-premises security identifier (SID), synchronised from the on-premises directory when Azure AD Connect is used.`,
				},
				resource.Attribute{
					Name:        "onpremises_sync_enabled",
					Description: `Whether this user is synchronised from an on-premises directory (` + "`" + `true` + "`" + `), no longer synchronised (` + "`" + `false` + "`" + `), or has never been synchronised (` + "`" + `null` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "onpremises_user_principal_name",
					Description: `The on-premise user principal name of the user.`,
				},
				resource.Attribute{
					Name:        "other_mails",
					Description: `A list of additional email addresses for the user.`,
				},
				resource.Attribute{
					Name:        "postal_code",
					Description: `The postal code for the user's postal address. The postal code is specific to the user's country/region. In the United States of America, this attribute contains the ZIP code.`,
				},
				resource.Attribute{
					Name:        "preferred_language",
					Description: `The user's preferred language, in ISO 639-1 notation.`,
				},
				resource.Attribute{
					Name:        "proxy_addresses",
					Description: `List of email addresses for the user that direct to the same mailbox.`,
				},
				resource.Attribute{
					Name:        "show_in_address_list",
					Description: `Whether or not the Outlook global address list should include this user.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state or province in the user's address.`,
				},
				resource.Attribute{
					Name:        "street_address",
					Description: `The street address of the user's place of business.`,
				},
				resource.Attribute{
					Name:        "surname",
					Description: `The user's surname (family name or last name).`,
				},
				resource.Attribute{
					Name:        "usage_location",
					Description: `The usage location of the user.`,
				},
				resource.Attribute{
					Name:        "user_principal_name",
					Description: `The user principal name (UPN) of the user.`,
				},
				resource.Attribute{
					Name:        "user_type",
					Description: `The user type in the directory. Possible values are ` + "`" + `Guest` + "`" + ` or ` + "`" + `Member` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "account_enabled",
					Description: `Whether or not the account is enabled.`,
				},
				resource.Attribute{
					Name:        "age_group",
					Description: `The age group of the user. Supported values are ` + "`" + `Adult` + "`" + `, ` + "`" + `NotAdult` + "`" + ` and ` + "`" + `Minor` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "business_phones",
					Description: `A list of telephone numbers for the user.`,
				},
				resource.Attribute{
					Name:        "city",
					Description: `The city in which the user is located.`,
				},
				resource.Attribute{
					Name:        "company_name",
					Description: `The company name which the user is associated. This property can be useful for describing the company that an external user comes from.`,
				},
				resource.Attribute{
					Name:        "consent_provided_for_minor",
					Description: `Whether consent has been obtained for minors. Supported values are ` + "`" + `Granted` + "`" + `, ` + "`" + `Denied` + "`" + ` and ` + "`" + `NotRequired` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `The country/region in which the user is located, e.g. ` + "`" + `US` + "`" + ` or ` + "`" + `UK` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cost_center",
					Description: `The cost center associated with the user.`,
				},
				resource.Attribute{
					Name:        "creation_type",
					Description: `Indicates whether the user account was created as a regular school or work account (` + "`" + `null` + "`" + `), an external account (` + "`" + `Invitation` + "`" + `), a local account for an Azure Active Directory B2C tenant (` + "`" + `LocalAccount` + "`" + `) or self-service sign-up using email verification (` + "`" + `EmailVerified` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "department",
					Description: `The name for the department in which the user works.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name of the user.`,
				},
				resource.Attribute{
					Name:        "division",
					Description: `The name of the division in which the user works.`,
				},
				resource.Attribute{
					Name:        "employee_id",
					Description: `The employee identifier assigned to the user by the organisation.`,
				},
				resource.Attribute{
					Name:        "employee_type",
					Description: `Captures enterprise worker type. For example, Employee, Contractor, Consultant, or Vendor.`,
				},
				resource.Attribute{
					Name:        "external_user_state",
					Description: `For an external user invited to the tenant, this property represents the invited user's invitation status. Possible values are ` + "`" + `PendingAcceptance` + "`" + ` or ` + "`" + `Accepted` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fax_number",
					Description: `The fax number of the user.`,
				},
				resource.Attribute{
					Name:        "given_name",
					Description: `The given name (first name) of the user.`,
				},
				resource.Attribute{
					Name:        "im_addresses",
					Description: `A list of instant message voice over IP (VOIP) session initiation protocol (SIP) addresses for the user.`,
				},
				resource.Attribute{
					Name:        "job_title",
					Description: `The user’s job title.`,
				},
				resource.Attribute{
					Name:        "mail",
					Description: `The SMTP address for the user.`,
				},
				resource.Attribute{
					Name:        "mail_nickname",
					Description: `The email alias of the user.`,
				},
				resource.Attribute{
					Name:        "manager_id",
					Description: `The object ID of the user's manager.`,
				},
				resource.Attribute{
					Name:        "mobile_phone",
					Description: `The primary cellular telephone number for the user.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The object ID of the user.`,
				},
				resource.Attribute{
					Name:        "office_location",
					Description: `The office location in the user's place of business.`,
				},
				resource.Attribute{
					Name:        "onpremises_distinguished_name",
					Description: `The on-premises distinguished name (DN) of the user, synchronised from the on-premises directory when Azure AD Connect is used.`,
				},
				resource.Attribute{
					Name:        "onpremises_domain_name",
					Description: `The on-premises FQDN, also called dnsDomainName, synchronised from the on-premises directory when Azure AD Connect is used.`,
				},
				resource.Attribute{
					Name:        "onpremises_immutable_id",
					Description: `The value used to associate an on-premise Active Directory user account with their Azure AD user object.`,
				},
				resource.Attribute{
					Name:        "onpremises_sam_account_name",
					Description: `The on-premise SAM account name of the user.`,
				},
				resource.Attribute{
					Name:        "onpremises_security_identifier",
					Description: `The on-premises security identifier (SID), synchronised from the on-premises directory when Azure AD Connect is used.`,
				},
				resource.Attribute{
					Name:        "onpremises_sync_enabled",
					Description: `Whether this user is synchronised from an on-premises directory (` + "`" + `true` + "`" + `), no longer synchronised (` + "`" + `false` + "`" + `), or has never been synchronised (` + "`" + `null` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "onpremises_user_principal_name",
					Description: `The on-premise user principal name of the user.`,
				},
				resource.Attribute{
					Name:        "other_mails",
					Description: `A list of additional email addresses for the user.`,
				},
				resource.Attribute{
					Name:        "postal_code",
					Description: `The postal code for the user's postal address. The postal code is specific to the user's country/region. In the United States of America, this attribute contains the ZIP code.`,
				},
				resource.Attribute{
					Name:        "preferred_language",
					Description: `The user's preferred language, in ISO 639-1 notation.`,
				},
				resource.Attribute{
					Name:        "proxy_addresses",
					Description: `List of email addresses for the user that direct to the same mailbox.`,
				},
				resource.Attribute{
					Name:        "show_in_address_list",
					Description: `Whether or not the Outlook global address list should include this user.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state or province in the user's address.`,
				},
				resource.Attribute{
					Name:        "street_address",
					Description: `The street address of the user's place of business.`,
				},
				resource.Attribute{
					Name:        "surname",
					Description: `The user's surname (family name or last name).`,
				},
				resource.Attribute{
					Name:        "usage_location",
					Description: `The usage location of the user.`,
				},
				resource.Attribute{
					Name:        "user_principal_name",
					Description: `The user principal name (UPN) of the user.`,
				},
				resource.Attribute{
					Name:        "user_type",
					Description: `The user type in the directory. Possible values are ` + "`" + `Guest` + "`" + ` or ` + "`" + `Member` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_users",
			Category:         "Users",
			ShortDescription: ``,
			Description: `

Gets basic information for multiple Azure Active Directory users.

`,
			Keywords: []string{
				"users",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ignore_missing",
					Description: `(Optional) Ignore missing users and return users that were found. The data source will still fail if no users are found. Cannot be specified with ` + "`" + `return_all` + "`" + `. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mail_nicknames",
					Description: `(Optional) The email aliases of the users.`,
				},
				resource.Attribute{
					Name:        "object_ids",
					Description: `(Optional) The object IDs of the users.`,
				},
				resource.Attribute{
					Name:        "return_all",
					Description: `(Optional) When ` + "`" + `true` + "`" + `, the data source will return all users. Cannot be used with ` + "`" + `ignore_missing` + "`" + `. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user_principal_names",
					Description: `(Optional) The user principal names (UPNs) of the users. ~> Either ` + "`" + `return_all` + "`" + `, or one of ` + "`" + `user_principal_names` + "`" + `, ` + "`" + `object_ids` + "`" + ` or ` + "`" + `mail_nicknames` + "`" + ` must be specified. These _may_ be specified as an empty list, in which case no results will be returned. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "mail_nicknames",
					Description: `The email aliases of the users.`,
				},
				resource.Attribute{
					Name:        "object_ids",
					Description: `The object IDs of the users.`,
				},
				resource.Attribute{
					Name:        "user_principal_names",
					Description: `The user principal names (UPNs) of the users.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `A list of users. Each ` + "`" + `user` + "`" + ` object provides the attributes documented below. --- ` + "`" + `user` + "`" + ` object exports the following:`,
				},
				resource.Attribute{
					Name:        "account_enabled",
					Description: `Whether or not the account is enabled.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name of the user.`,
				},
				resource.Attribute{
					Name:        "mail_nickname",
					Description: `The email alias of the user.`,
				},
				resource.Attribute{
					Name:        "mail",
					Description: `The primary email address of the user.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The object ID of the user.`,
				},
				resource.Attribute{
					Name:        "onpremises_immutable_id",
					Description: `The value used to associate an on-premises Active Directory user account with their Azure AD user object.`,
				},
				resource.Attribute{
					Name:        "onpremises_sam_account_name",
					Description: `The on-premise SAM account name of the user.`,
				},
				resource.Attribute{
					Name:        "onpremises_user_principal_name",
					Description: `The on-premise user principal name of the user.`,
				},
				resource.Attribute{
					Name:        "usage_location",
					Description: `The usage location of the user.`,
				},
				resource.Attribute{
					Name:        "user_principal_name",
					Description: `The user principal name (UPN) of the user.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "mail_nicknames",
					Description: `The email aliases of the users.`,
				},
				resource.Attribute{
					Name:        "object_ids",
					Description: `The object IDs of the users.`,
				},
				resource.Attribute{
					Name:        "user_principal_names",
					Description: `The user principal names (UPNs) of the users.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `A list of users. Each ` + "`" + `user` + "`" + ` object provides the attributes documented below. --- ` + "`" + `user` + "`" + ` object exports the following:`,
				},
				resource.Attribute{
					Name:        "account_enabled",
					Description: `Whether or not the account is enabled.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name of the user.`,
				},
				resource.Attribute{
					Name:        "mail_nickname",
					Description: `The email alias of the user.`,
				},
				resource.Attribute{
					Name:        "mail",
					Description: `The primary email address of the user.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The object ID of the user.`,
				},
				resource.Attribute{
					Name:        "onpremises_immutable_id",
					Description: `The value used to associate an on-premises Active Directory user account with their Azure AD user object.`,
				},
				resource.Attribute{
					Name:        "onpremises_sam_account_name",
					Description: `The on-premise SAM account name of the user.`,
				},
				resource.Attribute{
					Name:        "onpremises_user_principal_name",
					Description: `The on-premise user principal name of the user.`,
				},
				resource.Attribute{
					Name:        "usage_location",
					Description: `The usage location of the user.`,
				},
				resource.Attribute{
					Name:        "user_principal_name",
					Description: `The user principal name (UPN) of the user.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"azuread_access_package":                0,
		"azuread_access_package_catalog":        1,
		"azuread_access_package_catalog_role":   2,
		"azuread_administrative_unit":           3,
		"azuread_application":                   4,
		"azuread_application_published_app_ids": 5,
		"azuread_application_template":          6,
		"azuread_client_config":                 7,
		"azuread_directory_object":              8,
		"azuread_directory_roles":               9,
		"azuread_domains":                       10,
		"azuread_group":                         11,
		"azuread_groups":                        12,
		"azuread_service_principal":             13,
		"azuread_service_principals":            14,
		"azuread_user":                          15,
		"azuread_users":                         16,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
