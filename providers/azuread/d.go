package aws

import (
	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*Resource{

		&resource.Resource{
			Name:             "",
			Type:             "azuread_application",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing Application within Azure Active Directory.`,
			Description: `

Use this data source to access information about an existing Application within Azure Active Directory.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both ` + "`" + `Read and write all applications` + "`" + ` and ` + "`" + `Sign in and read user profile` + "`" + ` within the ` + "`" + `Windows Azure Active Directory` + "`" + ` API.

`,
			Keywords: []string{},
			Arguments: []resource.Argument{
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
				resource.Attribute{
					Name:        "group_membership_claims",
					Description: `The ` + "`" + `groups` + "`" + ` claim issued in a user or OAuth 2.0 access token that the app expects.`,
				},
				resource.Attribute{
					Name:        "required_resource_access",
					Description: `A collection of ` + "`" + `required_resource_access` + "`" + ` blocks as documented below.`,
				},
				resource.Attribute{
					Name:        "oauth2_permissions",
					Description: `A collection of OAuth 2.0 permission scopes that the web API (resource) app exposes to client apps. Each permission is covered by a ` + "`" + `oauth2_permission` + "`" + ` block as documented below. --- ` + "`" + `required_resource_access` + "`" + ` block exports the following:`,
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
					Description: `Specifies whether the id property references an ` + "`" + `OAuth2Permission` + "`" + ` or an ` + "`" + `AppRole` + "`" + `. --- ` + "`" + `oauth2_permission` + "`" + ` block exports the following:`,
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
					Description: `The name of this permission`,
				},
			},
			Attributes: []resource.Argument{
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
				resource.Attribute{
					Name:        "group_membership_claims",
					Description: `The ` + "`" + `groups` + "`" + ` claim issued in a user or OAuth 2.0 access token that the app expects.`,
				},
				resource.Attribute{
					Name:        "required_resource_access",
					Description: `A collection of ` + "`" + `required_resource_access` + "`" + ` blocks as documented below.`,
				},
				resource.Attribute{
					Name:        "oauth2_permissions",
					Description: `A collection of OAuth 2.0 permission scopes that the web API (resource) app exposes to client apps. Each permission is covered by a ` + "`" + `oauth2_permission` + "`" + ` block as documented below. --- ` + "`" + `required_resource_access` + "`" + ` block exports the following:`,
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
					Description: `Specifies whether the id property references an ` + "`" + `OAuth2Permission` + "`" + ` or an ` + "`" + `AppRole` + "`" + `. --- ` + "`" + `oauth2_permission` + "`" + ` block exports the following:`,
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
					Description: `The name of this permission`,
				},
			},
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
			Arguments: []resource.Argument{
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
			Attributes: []resource.Argument{
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
			Arguments: []resource.Argument{
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
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The Object ID of the Azure AD Group.`,
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
			Arguments: []resource.Argument{
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
					Name:        "id",
					Description: `The Object ID for the Service Principal.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The Object ID for the Service Principal.`,
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
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "user_principal_name",
					Description: `(Required) The User Principal Name of the Azure AD User.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `(Optional) Specifies the Object ID of the Application within Azure Active Directory. ->`,
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
			},
			Attributes: []resource.Argument{
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
			},
		},
	}

	dataSourcesMap = map[string]Resource{

		"azuread_application":       0,
		"azuread_domains":           1,
		"azuread_group":             2,
		"azuread_service_principal": 3,
		"azuread_user":              4,
	}
)

func GetDataSource(r string) (*resouce.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs]
}
