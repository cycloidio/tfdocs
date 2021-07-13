package azuread

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

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
					Description: `(Optional) Specifies the Object ID of the application. ~>`,
				},
				resource.Attribute{
					Name:        "api",
					Description: `An ` + "`" + `api` + "`" + ` block as documented below.`,
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
					Name:        "display_name",
					Description: `The display name for the application.`,
				},
				resource.Attribute{
					Name:        "fallback_public_client_enabled",
					Description: `The fallback application type as public client, such as an installed application running on a mobile device.`,
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
					Name:        "required_resource_access",
					Description: `A collection of ` + "`" + `required_resource_access` + "`" + ` blocks as documented below.`,
				},
				resource.Attribute{
					Name:        "sign_in_audience",
					Description: `The Microsoft account types that are supported for the current application. One of ` + "`" + `AzureADMyOrg` + "`" + `, ` + "`" + `AzureADMultipleOrgs` + "`" + `, ` + "`" + `AzureADandPersonalMicrosoftAccount` + "`" + ` or ` + "`" + `PersonalMicrosoftAccount` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "web",
					Description: `A ` + "`" + `web` + "`" + ` block as documented below. --- ` + "`" + `api` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "oauth2_permission_scope",
					Description: `One or more ` + "`" + `oauth2_permission_scope` + "`" + ` blocks as documented below, to describe delegated permissions exposed by the web API represented by this application. --- ` + "`" + `oauth2_permission_scope` + "`" + ` block exports the following:`,
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
					Description: `The value that is used for the ` + "`" + `roles` + "`" + ` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal. --- ` + "`" + `optional_claims` + "`" + ` block exports the following:`,
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
					Description: `The source of the claim. If ` + "`" + `source` + "`" + ` is absent, the claim is a predefined optional claim. If ` + "`" + `source` + "`" + ` is ` + "`" + `user` + "`" + `, the value of ` + "`" + `name` + "`" + ` is the extension property from the user object. --- ` + "`" + `required_resource_access` + "`" + ` block exports the following:`,
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
					Description: `Specifies whether the ` + "`" + `id` + "`" + ` property references an app role or an OAuth2 permission scope. Possible values are ` + "`" + `Role` + "`" + ` or ` + "`" + `Scope` + "`" + `. --- ` + "`" + `web` + "`" + ` block exports the following:`,
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
					Name:        "app_roles",
					Description: `A collection of ` + "`" + `app_role` + "`" + ` blocks as documented below. For more information see [official documentation on Application Roles](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).`,
				},
				resource.Attribute{
					Name:        "application_id",
					Description: `The Application ID (also called Client ID).`,
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
					Name:        "group_membership_claims",
					Description: `The ` + "`" + `groups` + "`" + ` claim issued in a user or OAuth 2.0 access token that the app expects.`,
				},
				resource.Attribute{
					Name:        "identifier_uris",
					Description: `A list of user-defined URI(s) that uniquely identify a Web application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant.`,
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
					Name:        "required_resource_access",
					Description: `A collection of ` + "`" + `required_resource_access` + "`" + ` blocks as documented below.`,
				},
				resource.Attribute{
					Name:        "sign_in_audience",
					Description: `The Microsoft account types that are supported for the current application. One of ` + "`" + `AzureADMyOrg` + "`" + `, ` + "`" + `AzureADMultipleOrgs` + "`" + `, ` + "`" + `AzureADandPersonalMicrosoftAccount` + "`" + ` or ` + "`" + `PersonalMicrosoftAccount` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "web",
					Description: `A ` + "`" + `web` + "`" + ` block as documented below. --- ` + "`" + `api` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "oauth2_permission_scope",
					Description: `One or more ` + "`" + `oauth2_permission_scope` + "`" + ` blocks as documented below, to describe delegated permissions exposed by the web API represented by this application. --- ` + "`" + `oauth2_permission_scope` + "`" + ` block exports the following:`,
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
					Description: `The value that is used for the ` + "`" + `roles` + "`" + ` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal. --- ` + "`" + `optional_claims` + "`" + ` block exports the following:`,
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
					Description: `The source of the claim. If ` + "`" + `source` + "`" + ` is absent, the claim is a predefined optional claim. If ` + "`" + `source` + "`" + ` is ` + "`" + `user` + "`" + `, the value of ` + "`" + `name` + "`" + ` is the extension property from the user object. --- ` + "`" + `required_resource_access` + "`" + ` block exports the following:`,
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
					Description: `Specifies whether the ` + "`" + `id` + "`" + ` property references an app role or an OAuth2 permission scope. Possible values are ` + "`" + `Role` + "`" + ` or ` + "`" + `Scope` + "`" + `. --- ` + "`" + `web` + "`" + ` block exports the following:`,
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
			Type:             "azuread_domains",
			Category:         "Domains",
			ShortDescription: ``,
			Description: `

Use this data source to access information about existing Domains within Azure Active Directory.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to ` + "`" + `Directory.Read.All` + "`" + ` within the ` + "`" + `Windows Azure Active Directory` + "`" + ` API.

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
					Description: `(Optional) A list of supported services that must be supported by a domain. Possible values include ` + "`" + `Email` + "`" + `, ` + "`" + `Sharepoint` + "`" + `, ` + "`" + `EmailInternalRelayOnly` + "`" + `, ` + "`" + `OfficeCommunicationsOnline` + "`" + `, ` + "`" + `SharePointDefaultDomain` + "`" + `, ` + "`" + `FullRedelegation` + "`" + `, ` + "`" + `SharePointPublic` + "`" + `, ` + "`" + `OrgIdAuthentication` + "`" + `, ` + "`" + `Yammer` + "`" + ` and ` + "`" + `Intune` + "`" + `. ~>`,
				},
				resource.Attribute{
					Name:        "domains",
					Description: `A list of tenant domains. Each ` + "`" + `domain` + "`" + ` object provides the attributes documented below. ` + "`" + `domain` + "`" + ` object exports the following:`,
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
					Description: `A list of tenant domains. Each ` + "`" + `domain` + "`" + ` object provides the attributes documented below. ` + "`" + `domain` + "`" + ` object exports the following:`,
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
					Description: `(Optional) Whether the group is a security group. ~>`,
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
					Name:        "object_id",
					Description: `The object ID of the group.`,
				},
				resource.Attribute{
					Name:        "mail_enabled",
					Description: `Whether the group is mail-enabled.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `The object IDs of the group members.`,
				},
				resource.Attribute{
					Name:        "owners",
					Description: `The object IDs of the group owners.`,
				},
				resource.Attribute{
					Name:        "security_enabled",
					Description: `Whether the group is a security group.`,
				},
				resource.Attribute{
					Name:        "types",
					Description: `A list of group types configured for the group. The only supported type is ` + "`" + `Unified` + "`" + `, which specifies a Microsoft 365 group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The optional description of the group.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name for the group.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The object ID of the group.`,
				},
				resource.Attribute{
					Name:        "mail_enabled",
					Description: `Whether the group is mail-enabled.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `The object IDs of the group members.`,
				},
				resource.Attribute{
					Name:        "owners",
					Description: `The object IDs of the group owners.`,
				},
				resource.Attribute{
					Name:        "security_enabled",
					Description: `Whether the group is a security group.`,
				},
				resource.Attribute{
					Name:        "types",
					Description: `A list of group types configured for the group. The only supported type is ` + "`" + `Unified` + "`" + `, which specifies a Microsoft 365 group.`,
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
					Name:        "object_ids",
					Description: `(Optional) The object IDs of the groups. ~>`,
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
					Description: `(Optional) The object ID of the service principal. ~>`,
				},
				resource.Attribute{
					Name:        "app_roles",
					Description: `A collection of ` + "`" + `app_roles` + "`" + ` blocks as documented below. For more information [official documentation](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The object ID for the service principal.`,
				},
				resource.Attribute{
					Name:        "oauth2_permission_scopes",
					Description: `A collection of OAuth 2.0 delegated permissions exposed by the associated application. Each permission is covered by an ` + "`" + `oauth2_permission_scopes` + "`" + ` block as documented below. --- ` + "`" + `app_roles` + "`" + ` block exports the following:`,
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
					Description: `Specifies the value of the roles claim that the application should expect in the authentication and access tokens. --- ` + "`" + `oauth2_permission_scopes` + "`" + ` block exports the following:`,
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
					Description: `The value that is used for the ` + "`" + `scp` + "`" + ` claim in OAuth 2.0 access tokens.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "app_roles",
					Description: `A collection of ` + "`" + `app_roles` + "`" + ` blocks as documented below. For more information [official documentation](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The object ID for the service principal.`,
				},
				resource.Attribute{
					Name:        "oauth2_permission_scopes",
					Description: `A collection of OAuth 2.0 delegated permissions exposed by the associated application. Each permission is covered by an ` + "`" + `oauth2_permission_scopes` + "`" + ` block as documented below. --- ` + "`" + `app_roles` + "`" + ` block exports the following:`,
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
					Description: `Specifies the value of the roles claim that the application should expect in the authentication and access tokens. --- ` + "`" + `oauth2_permission_scopes` + "`" + ` block exports the following:`,
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
					Description: `The value that is used for the ` + "`" + `scp` + "`" + ` claim in OAuth 2.0 access tokens.`,
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
					Name:        "mail_nickname",
					Description: `(Optional) The email alias of the user.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `(Optional) The object ID of the user.`,
				},
				resource.Attribute{
					Name:        "user_principal_name",
					Description: `(Optional) The user principal name (UPN) of the user. ~>`,
				},
				resource.Attribute{
					Name:        "account_enabled",
					Description: `Whether or not the account is enabled.`,
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
					Name:        "country",
					Description: `The country/region in which the user is located, e.g. ` + "`" + `US` + "`" + ` or ` + "`" + `UK` + "`" + `.`,
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
					Name:        "given_name",
					Description: `The given name (first name) of the user.`,
				},
				resource.Attribute{
					Name:        "job_title",
					Description: `The user’s job title.`,
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
					Name:        "mobile_phone",
					Description: `The primary cellular telephone number for the user.`,
				},
				resource.Attribute{
					Name:        "office_location",
					Description: `The office location in the user's place of business.`,
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
					Name:        "onpremises_user_principal_name",
					Description: `The on-premise user principal name of the user.`,
				},
				resource.Attribute{
					Name:        "postal_code",
					Description: `The postal code for the user's postal address. The postal code is specific to the user's country/region. In the United States of America, this attribute contains the ZIP code.`,
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
					Name:        "city",
					Description: `The city in which the user is located.`,
				},
				resource.Attribute{
					Name:        "company_name",
					Description: `The company name which the user is associated. This property can be useful for describing the company that an external user comes from.`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `The country/region in which the user is located, e.g. ` + "`" + `US` + "`" + ` or ` + "`" + `UK` + "`" + `.`,
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
					Name:        "given_name",
					Description: `The given name (first name) of the user.`,
				},
				resource.Attribute{
					Name:        "job_title",
					Description: `The user’s job title.`,
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
					Name:        "mobile_phone",
					Description: `The primary cellular telephone number for the user.`,
				},
				resource.Attribute{
					Name:        "office_location",
					Description: `The office location in the user's place of business.`,
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
					Name:        "onpremises_user_principal_name",
					Description: `The on-premise user principal name of the user.`,
				},
				resource.Attribute{
					Name:        "postal_code",
					Description: `The postal code for the user's postal address. The postal code is specific to the user's country/region. In the United States of America, this attribute contains the ZIP code.`,
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

Gets object IDs or user principal names for multiple Azure Active Directory users.

`,
			Keywords: []string{
				"users",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "mail_nicknames",
					Description: `(Optional) The email aliases of the users.`,
				},
				resource.Attribute{
					Name:        "ignore_missing",
					Description: `(Optional) Ignore missing users and return users that were found. The data source will still fail if no users are found. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "object_ids",
					Description: `(Optional) The object IDs of the users.`,
				},
				resource.Attribute{
					Name:        "user_principal_names",
					Description: `(Optional) The user principal names (UPNs) of the users. ~>`,
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
					Description: `A list of users. Each ` + "`" + `user` + "`" + ` object provides the attributes documented below. ___ ` + "`" + `user` + "`" + ` object exports the following:`,
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
					Description: `A list of users. Each ` + "`" + `user` + "`" + ` object provides the attributes documented below. ___ ` + "`" + `user` + "`" + ` object exports the following:`,
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

		"azuread_application":       0,
		"azuread_client_config":     1,
		"azuread_domains":           2,
		"azuread_group":             3,
		"azuread_groups":            4,
		"azuread_service_principal": 5,
		"azuread_user":              6,
		"azuread_users":             7,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
