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
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing Application within Azure Active Directory.`,
			Description: `

Use this data source to access information about an existing Application within Azure Active Directory.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both ` + "`" + `Read and write all (or owned by) applications` + "`" + ` and ` + "`" + `Sign in and read user profile` + "`" + ` within the ` + "`" + `Windows Azure Active Directory` + "`" + ` API.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "object_id",
					Description: `(Optional) Specifies the Object ID of the Application within Azure Active Directory.`,
				},
				resource.Attribute{
					Name:        "application_id",
					Description: `(Optional) Specifies the Application ID of the Azure Active Directory Application.`,
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
					Name:        "logout_url",
					Description: `The URL of the logout page.`,
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
				resource.Attribute{
					Name:        "group_membership_claims",
					Description: `The ` + "`" + `groups` + "`" + ` claim issued in a user or OAuth 2.0 access token that the app expects.`,
				},
				resource.Attribute{
					Name:        "optional_claims",
					Description: `A collection of ` + "`" + `access_token` + "`" + ` or ` + "`" + `id_token` + "`" + ` blocks as documented below which list the optional claims configured for each token type. For more information see https://docs.microsoft.com/en-us/azure/active-directory/develop/active-directory-optional-claims`,
				},
				resource.Attribute{
					Name:        "owners",
					Description: `A list of User Object IDs that are assigned ownership of the application registration.`,
				},
				resource.Attribute{
					Name:        "required_resource_access",
					Description: `A collection of ` + "`" + `required_resource_access` + "`" + ` blocks as documented below.`,
				},
				resource.Attribute{
					Name:        "oauth2_permissions",
					Description: `A collection of OAuth 2.0 permission scopes that the web API (resource) app exposes to client apps. Each permission is covered by a ` + "`" + `oauth2_permission` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "app_roles",
					Description: `A collection of ` + "`" + `app_role` + "`" + ` blocks as documented below. For more information https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles`,
				},
				resource.Attribute{
					Name:        "public_client",
					Description: `Is this Azure AD Application available publically? --- ` + "`" + `required_resource_access` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "resource_app_id",
					Description: `The unique identifier for the resource that the application requires access to.`,
				},
				resource.Attribute{
					Name:        "resource_access",
					Description: `A collection of ` + "`" + `resource_access` + "`" + ` blocks as documented below --- ` + "`" + `resource_access` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier for one of the ` + "`" + `OAuth2Permission` + "`" + ` or ` + "`" + `AppRole` + "`" + ` instances that the resource application exposes.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Specifies whether the id property references an ` + "`" + `OAuth2Permission` + "`" + ` or an ` + "`" + `AppRole` + "`" + `. --- ` + "`" + `access_token` + "`" + ` and/or ` + "`" + `id_token` + "`" + ` blocks export the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the optional claim.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `The source of the claim. If ` + "`" + `source` + "`" + ` is absent, the claim is a predefined optional claim. If ` + "`" + `source` + "`" + ` is ` + "`" + `user` + "`" + `, the value of ` + "`" + `name` + "`" + ` is the extension property from the user object.`,
				},
				resource.Attribute{
					Name:        "essential",
					Description: `Whether the claim specified by the client is necessary to ensure a smooth authorization experience.`,
				},
				resource.Attribute{
					Name:        "additional_properties",
					Description: `List of Additional Properties of the claim. If a property exists in this list, it modifies the behaviour of the optional claim. --- ` + "`" + `oauth2_permission` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier for one of the ` + "`" + `OAuth2Permission` + "`" + ``,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the permission`,
				},
				resource.Attribute{
					Name:        "admin_consent_description",
					Description: `The description of the admin consent`,
				},
				resource.Attribute{
					Name:        "admin_consent_display_name",
					Description: `The display name of the admin consent`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `Is this permission enabled?`,
				},
				resource.Attribute{
					Name:        "user_consent_description",
					Description: `The description of the user consent`,
				},
				resource.Attribute{
					Name:        "user_consent_display_name",
					Description: `The display name of the user consent`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The name of this permission --- ` + "`" + `app_role` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier of the ` + "`" + `app_role` + "`" + `.`,
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
					Name:        "is_enabled",
					Description: `Determines if the app role is enabled.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Specifies the value of the roles claim that the application should expect in the authentication and access tokens.`,
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
					Name:        "logout_url",
					Description: `The URL of the logout page.`,
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
				resource.Attribute{
					Name:        "group_membership_claims",
					Description: `The ` + "`" + `groups` + "`" + ` claim issued in a user or OAuth 2.0 access token that the app expects.`,
				},
				resource.Attribute{
					Name:        "optional_claims",
					Description: `A collection of ` + "`" + `access_token` + "`" + ` or ` + "`" + `id_token` + "`" + ` blocks as documented below which list the optional claims configured for each token type. For more information see https://docs.microsoft.com/en-us/azure/active-directory/develop/active-directory-optional-claims`,
				},
				resource.Attribute{
					Name:        "owners",
					Description: `A list of User Object IDs that are assigned ownership of the application registration.`,
				},
				resource.Attribute{
					Name:        "required_resource_access",
					Description: `A collection of ` + "`" + `required_resource_access` + "`" + ` blocks as documented below.`,
				},
				resource.Attribute{
					Name:        "oauth2_permissions",
					Description: `A collection of OAuth 2.0 permission scopes that the web API (resource) app exposes to client apps. Each permission is covered by a ` + "`" + `oauth2_permission` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "app_roles",
					Description: `A collection of ` + "`" + `app_role` + "`" + ` blocks as documented below. For more information https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles`,
				},
				resource.Attribute{
					Name:        "public_client",
					Description: `Is this Azure AD Application available publically? --- ` + "`" + `required_resource_access` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "resource_app_id",
					Description: `The unique identifier for the resource that the application requires access to.`,
				},
				resource.Attribute{
					Name:        "resource_access",
					Description: `A collection of ` + "`" + `resource_access` + "`" + ` blocks as documented below --- ` + "`" + `resource_access` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier for one of the ` + "`" + `OAuth2Permission` + "`" + ` or ` + "`" + `AppRole` + "`" + ` instances that the resource application exposes.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Specifies whether the id property references an ` + "`" + `OAuth2Permission` + "`" + ` or an ` + "`" + `AppRole` + "`" + `. --- ` + "`" + `access_token` + "`" + ` and/or ` + "`" + `id_token` + "`" + ` blocks export the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the optional claim.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `The source of the claim. If ` + "`" + `source` + "`" + ` is absent, the claim is a predefined optional claim. If ` + "`" + `source` + "`" + ` is ` + "`" + `user` + "`" + `, the value of ` + "`" + `name` + "`" + ` is the extension property from the user object.`,
				},
				resource.Attribute{
					Name:        "essential",
					Description: `Whether the claim specified by the client is necessary to ensure a smooth authorization experience.`,
				},
				resource.Attribute{
					Name:        "additional_properties",
					Description: `List of Additional Properties of the claim. If a property exists in this list, it modifies the behaviour of the optional claim. --- ` + "`" + `oauth2_permission` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier for one of the ` + "`" + `OAuth2Permission` + "`" + ``,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the permission`,
				},
				resource.Attribute{
					Name:        "admin_consent_description",
					Description: `The description of the admin consent`,
				},
				resource.Attribute{
					Name:        "admin_consent_display_name",
					Description: `The display name of the admin consent`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `Is this permission enabled?`,
				},
				resource.Attribute{
					Name:        "user_consent_description",
					Description: `The description of the user consent`,
				},
				resource.Attribute{
					Name:        "user_consent_display_name",
					Description: `The display name of the user consent`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The name of this permission --- ` + "`" + `app_role` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier of the ` + "`" + `app_role` + "`" + `.`,
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
					Name:        "is_enabled",
					Description: `Determines if the app role is enabled.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Specifies the value of the roles claim that the application should expect in the authentication and access tokens.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_client_config",
			Category:         "Data Sources",
			ShortDescription: `Gets information about the configuration of the azuread provider.`,
			Description: `

Use this data source to access the configuration of the AzureRM provider.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_domains",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing Domains within Azure Active Directory.`,
			Description: `

Use this data source to access information about an existing Domains within Azure Active Directory.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to ` + "`" + `Directory.Read.All` + "`" + ` within the ` + "`" + `Windows Azure Active Directory` + "`" + ` API.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "include_unverified",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` if unverified Azure AD Domains should be included. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "only_default",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to only return the default domain.`,
				},
				resource.Attribute{
					Name:        "only_initial",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to only return the initial domain, which is your primary Azure Active Directory tenant domain. Defaults to ` + "`" + `false` + "`" + `. ->`,
				},
				resource.Attribute{
					Name:        "domains",
					Description: `One or more ` + "`" + `domain` + "`" + ` blocks as defined below. The ` + "`" + `domain` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `The name of the domain.`,
				},
				resource.Attribute{
					Name:        "authentication_type",
					Description: `The authentication type of the domain (Managed or Federated).`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `` + "`" + `True` + "`" + ` if this is the default domain that is used for user creation.`,
				},
				resource.Attribute{
					Name:        "is_initial",
					Description: `` + "`" + `True` + "`" + ` if this is the initial domain created by Azure Activie Directory.`,
				},
				resource.Attribute{
					Name:        "is_verified",
					Description: `` + "`" + `True` + "`" + ` if the domain has completed domain ownership verification.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "domains",
					Description: `One or more ` + "`" + `domain` + "`" + ` blocks as defined below. The ` + "`" + `domain` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `The name of the domain.`,
				},
				resource.Attribute{
					Name:        "authentication_type",
					Description: `The authentication type of the domain (Managed or Federated).`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `` + "`" + `True` + "`" + ` if this is the default domain that is used for user creation.`,
				},
				resource.Attribute{
					Name:        "is_initial",
					Description: `` + "`" + `True` + "`" + ` if this is the initial domain created by Azure Activie Directory.`,
				},
				resource.Attribute{
					Name:        "is_verified",
					Description: `` + "`" + `True` + "`" + ` if the domain has completed domain ownership verification.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_group",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an Azure Active Directory group.`,
			Description: `

Gets information about an Azure Active Directory group.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to ` + "`" + `Read directory data` + "`" + ` within the ` + "`" + `Windows Azure Active Directory` + "`" + ` API.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The Name of the AD Group we want to lookup.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `(Optional) Specifies the Object ID of the AD Group within Azure Active Directory. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Object ID of the Azure AD Group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the AD Group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Azure AD Group.`,
				},
				resource.Attribute{
					Name:        "owners",
					Description: `The Object IDs of the Azure AD Group owners.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `The Object IDs of the Azure AD Group members.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Object ID of the Azure AD Group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the AD Group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Azure AD Group.`,
				},
				resource.Attribute{
					Name:        "owners",
					Description: `The Object IDs of the Azure AD Group owners.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `The Object IDs of the Azure AD Group members.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_groups",
			Category:         "Data Sources",
			ShortDescription: `Gets information about Azure Active Directory groups.`,
			Description: `

Gets Object IDs or Display Names for multiple Azure Active Directory groups.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to ` + "`" + `Read directory data` + "`" + ` within the ` + "`" + `Windows Azure Active Directory` + "`" + ` API.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "names",
					Description: `(Optional) The Display Names of the Azure AD Groups.`,
				},
				resource.Attribute{
					Name:        "object_ids",
					Description: `(Optional) The Object IDs of the Azure AD Groups. ->`,
				},
				resource.Attribute{
					Name:        "object_ids",
					Description: `The Object IDs of the Azure AD Groups.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `The Display Names of the Azure AD Groups.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "object_ids",
					Description: `The Object IDs of the Azure AD Groups.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `The Display Names of the Azure AD Groups.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_service_principal",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing Service Principal associated with an Application within Azure Active Directory.`,
			Description: `

Gets information about an existing Service Principal associated with an Application within Azure Active Directory.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both ` + "`" + `Read and write all applications` + "`" + ` and ` + "`" + `Sign in and read user profile` + "`" + ` within the ` + "`" + `Windows Azure Active Directory` + "`" + ` API.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "application_id",
					Description: `(Optional) The ID of the Azure AD Application.`,
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
					Name:        "app_roles",
					Description: `A collection of ` + "`" + `app_role` + "`" + ` blocks as documented below. For more information https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles`,
				},
				resource.Attribute{
					Name:        "oauth2_permissions",
					Description: `A collection of OAuth 2.0 permissions exposed by the associated application. Each permission is covered by a ` + "`" + `oauth2_permission` + "`" + ` block as documented below. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Object ID for the Service Principal. --- ` + "`" + `oauth2_permission` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier for one of the ` + "`" + `OAuth2Permission` + "`" + ``,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the permission`,
				},
				resource.Attribute{
					Name:        "admin_consent_description",
					Description: `The description of the admin consent`,
				},
				resource.Attribute{
					Name:        "admin_consent_display_name",
					Description: `The display name of the admin consent`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `Is this permission enabled?`,
				},
				resource.Attribute{
					Name:        "user_consent_description",
					Description: `The description of the user consent`,
				},
				resource.Attribute{
					Name:        "user_consent_display_name",
					Description: `The display name of the user consent`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The name of this permission --- ` + "`" + `app_role` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier of the ` + "`" + `app_role` + "`" + `.`,
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
					Name:        "is_enabled",
					Description: `Determines if the app role is enabled.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Specifies the value of the roles claim that the application should expect in the authentication and access tokens.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Object ID for the Service Principal. --- ` + "`" + `oauth2_permission` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier for one of the ` + "`" + `OAuth2Permission` + "`" + ``,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the permission`,
				},
				resource.Attribute{
					Name:        "admin_consent_description",
					Description: `The description of the admin consent`,
				},
				resource.Attribute{
					Name:        "admin_consent_display_name",
					Description: `The display name of the admin consent`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `Is this permission enabled?`,
				},
				resource.Attribute{
					Name:        "user_consent_description",
					Description: `The description of the user consent`,
				},
				resource.Attribute{
					Name:        "user_consent_display_name",
					Description: `The display name of the user consent`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The name of this permission --- ` + "`" + `app_role` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier of the ` + "`" + `app_role` + "`" + `.`,
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
					Name:        "is_enabled",
					Description: `Determines if the app role is enabled.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Specifies the value of the roles claim that the application should expect in the authentication and access tokens.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_user",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an Azure Active Directory user.`,
			Description: `

Gets information about an Azure Active Directory user.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to ` + "`" + `Read directory data` + "`" + ` within the ` + "`" + `Windows Azure Active Directory` + "`" + ` API.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "user_principal_name",
					Description: `(Required) The User Principal Name of the Azure AD User.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `(Optional) Specifies the Object ID of the Application within Azure Active Directory.`,
				},
				resource.Attribute{
					Name:        "mail_nickname",
					Description: `(Optional) The email alias of the Azure AD User. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Object ID of the Azure AD User.`,
				},
				resource.Attribute{
					Name:        "user_principal_name",
					Description: `The User Principal Name of the Azure AD User.`,
				},
				resource.Attribute{
					Name:        "account_enabled",
					Description: `` + "`" + `True` + "`" + ` if the account is enabled; otherwise ` + "`" + `False` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The Display Name of the Azure AD User.`,
				},
				resource.Attribute{
					Name:        "mail",
					Description: `The primary email address of the Azure AD User.`,
				},
				resource.Attribute{
					Name:        "mail_nickname",
					Description: `The email alias of the Azure AD User.`,
				},
				resource.Attribute{
					Name:        "mail_nickname",
					Description: `The email alias of the Azure AD User.`,
				},
				resource.Attribute{
					Name:        "onpremises_sam_account_name",
					Description: `The on premise sam account name of the Azure AD User.`,
				},
				resource.Attribute{
					Name:        "onpremises_user_principal_name",
					Description: `The on premise user principal name of the Azure AD User.`,
				},
				resource.Attribute{
					Name:        "usage_location",
					Description: `The usage location of the Azure AD User.`,
				},
				resource.Attribute{
					Name:        "immutable_id",
					Description: `The value used to associate an on-premises Active Directory user account with their Azure AD user object.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Object ID of the Azure AD User.`,
				},
				resource.Attribute{
					Name:        "user_principal_name",
					Description: `The User Principal Name of the Azure AD User.`,
				},
				resource.Attribute{
					Name:        "account_enabled",
					Description: `` + "`" + `True` + "`" + ` if the account is enabled; otherwise ` + "`" + `False` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The Display Name of the Azure AD User.`,
				},
				resource.Attribute{
					Name:        "mail",
					Description: `The primary email address of the Azure AD User.`,
				},
				resource.Attribute{
					Name:        "mail_nickname",
					Description: `The email alias of the Azure AD User.`,
				},
				resource.Attribute{
					Name:        "mail_nickname",
					Description: `The email alias of the Azure AD User.`,
				},
				resource.Attribute{
					Name:        "onpremises_sam_account_name",
					Description: `The on premise sam account name of the Azure AD User.`,
				},
				resource.Attribute{
					Name:        "onpremises_user_principal_name",
					Description: `The on premise user principal name of the Azure AD User.`,
				},
				resource.Attribute{
					Name:        "usage_location",
					Description: `The usage location of the Azure AD User.`,
				},
				resource.Attribute{
					Name:        "immutable_id",
					Description: `The value used to associate an on-premises Active Directory user account with their Azure AD user object.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_users",
			Category:         "Data Sources",
			ShortDescription: `Gets information about Azure Active Directory users.`,
			Description: `

Gets Object IDs or UPNs for multiple Azure Active Directory users.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to ` + "`" + `Read directory data` + "`" + ` within the ` + "`" + `Windows Azure Active Directory` + "`" + ` API.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "user_principal_names",
					Description: `(Optional) The User Principal Names of the Azure AD Users.`,
				},
				resource.Attribute{
					Name:        "object_ids",
					Description: `(Optional) The Object IDs of the Azure AD Users.`,
				},
				resource.Attribute{
					Name:        "mail_nicknames",
					Description: `(Optional) The email aliases of the Azure AD Users. ->`,
				},
				resource.Attribute{
					Name:        "ignore_missing",
					Description: `(Optional) Ignore missing users and return users that were found. The data source will still fail if no users are found. Defaults to false. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "object_ids",
					Description: `The Object IDs of the Azure AD Users.`,
				},
				resource.Attribute{
					Name:        "user_principal_names",
					Description: `The User Principal Names of the Azure AD Users.`,
				},
				resource.Attribute{
					Name:        "mail_nicknames",
					Description: `The email aliases of the Azure AD Users.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `An Array of Azure AD Users. Each ` + "`" + `user` + "`" + ` object consists of the fields documented below. ___ The ` + "`" + `user` + "`" + ` object consists of:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Object ID of the Azure AD User.`,
				},
				resource.Attribute{
					Name:        "user_principal_name",
					Description: `The User Principal Name of the Azure AD User.`,
				},
				resource.Attribute{
					Name:        "account_enabled",
					Description: `` + "`" + `True` + "`" + ` if the account is enabled; otherwise ` + "`" + `False` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The Display Name of the Azure AD User.`,
				},
				resource.Attribute{
					Name:        "mail",
					Description: `The primary email address of the Azure AD User.`,
				},
				resource.Attribute{
					Name:        "mail_nickname",
					Description: `The email alias of the Azure AD User.`,
				},
				resource.Attribute{
					Name:        "onpremises_sam_account_name",
					Description: `The on premise sam account name of the Azure AD User.`,
				},
				resource.Attribute{
					Name:        "onpremises_user_principal_name",
					Description: `The on premise user principal name of the Azure AD User.`,
				},
				resource.Attribute{
					Name:        "usage_location",
					Description: `The usage location of the Azure AD User.`,
				},
				resource.Attribute{
					Name:        "immutable_id",
					Description: `The value used to associate an on-premises Active Directory user account with their Azure AD user object.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "object_ids",
					Description: `The Object IDs of the Azure AD Users.`,
				},
				resource.Attribute{
					Name:        "user_principal_names",
					Description: `The User Principal Names of the Azure AD Users.`,
				},
				resource.Attribute{
					Name:        "mail_nicknames",
					Description: `The email aliases of the Azure AD Users.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `An Array of Azure AD Users. Each ` + "`" + `user` + "`" + ` object consists of the fields documented below. ___ The ` + "`" + `user` + "`" + ` object consists of:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Object ID of the Azure AD User.`,
				},
				resource.Attribute{
					Name:        "user_principal_name",
					Description: `The User Principal Name of the Azure AD User.`,
				},
				resource.Attribute{
					Name:        "account_enabled",
					Description: `` + "`" + `True` + "`" + ` if the account is enabled; otherwise ` + "`" + `False` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The Display Name of the Azure AD User.`,
				},
				resource.Attribute{
					Name:        "mail",
					Description: `The primary email address of the Azure AD User.`,
				},
				resource.Attribute{
					Name:        "mail_nickname",
					Description: `The email alias of the Azure AD User.`,
				},
				resource.Attribute{
					Name:        "onpremises_sam_account_name",
					Description: `The on premise sam account name of the Azure AD User.`,
				},
				resource.Attribute{
					Name:        "onpremises_user_principal_name",
					Description: `The on premise user principal name of the Azure AD User.`,
				},
				resource.Attribute{
					Name:        "usage_location",
					Description: `The usage location of the Azure AD User.`,
				},
				resource.Attribute{
					Name:        "immutable_id",
					Description: `The value used to associate an on-premises Active Directory user account with their Azure AD user object.`,
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
