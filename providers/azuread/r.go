package aws

import (
	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*Resource{

		&resource.Resource{
			Name:             "",
			Type:             "azuread_application",
			Category:         "Azure Active Directory Resources",
			ShortDescription: `Manages an Application within Azure Active Directory.`,
			Description:      ``,
			Keywords: []string{
				"azure",
				"active",
				"directory",
				"application",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The display name for the application.`,
				},
				resource.Attribute{
					Name:        "homepage",
					Description: `(optional) The URL to the application's home page. If no homepage is specified this defaults to ` + "`" + `https://{name}` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "identifier_uris",
					Description: `(Optional) A list of user-defined URI(s) that uniquely identify a Web application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant.`,
				},
				resource.Attribute{
					Name:        "reply_urls",
					Description: `(Optional) A list of URLs that user tokens are sent to for sign in, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to.`,
				},
				resource.Attribute{
					Name:        "available_to_other_tenants",
					Description: `(Optional) Is this Azure AD Application available to other tenants? Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "oauth2_allow_implicit_flow",
					Description: `(Optional) Does this Azure AD Application allow OAuth2.0 implicit flow tokens? Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "group_membership_claims",
					Description: `(Optional) Configures the ` + "`" + `groups` + "`" + ` claim issued in a user or OAuth 2.0 access token that the app expects. Defaults to ` + "`" + `SecurityGroup` + "`" + `. Possible values are ` + "`" + `None` + "`" + `, ` + "`" + `SecurityGroup` + "`" + ` or ` + "`" + `All` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "required_resource_access",
					Description: `(Optional) A collection of ` + "`" + `required_resource_access` + "`" + ` blocks as documented below.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type of an application: ` + "`" + `webapp/api` + "`" + ` or ` + "`" + `native` + "`" + `. Defaults to ` + "`" + `webapp/api` + "`" + `. For ` + "`" + `native` + "`" + ` apps type ` + "`" + `identifier_uris` + "`" + ` property can not not be set. --- ` + "`" + `required_resource_access` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "resource_app_id",
					Description: `(Required) The unique identifier for the resource that the application requires access to. This should be equal to the appId declared on the target resource application.`,
				},
				resource.Attribute{
					Name:        "resource_access",
					Description: `(Required) A collection of ` + "`" + `resource_access` + "`" + ` blocks as documented below. --- ` + "`" + `resource_access` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The unique identifier for one of the ` + "`" + `OAuth2Permission` + "`" + ` or ` + "`" + `AppRole` + "`" + ` instances that the resource application exposes.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Specifies whether the id property references an ` + "`" + `OAuth2Permission` + "`" + ` or an ` + "`" + `AppRole` + "`" + `. Possible values are ` + "`" + `Scope` + "`" + ` or ` + "`" + `Role` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "application_id",
					Description: `The Application ID.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The Application's Object ID.`,
				},
				resource.Attribute{
					Name:        "oauth2_permissions",
					Description: `A collection of OAuth 2.0 permission scopes that the web API (resource) app exposes to client apps. Each permission is covered by a ` + "`" + `oauth2_permission` + "`" + ` block as documented below. --- ` + "`" + `oauth2_permission` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier for one of the ` + "`" + `OAuth2Permission` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the permission.`,
				},
				resource.Attribute{
					Name:        "admin_consent_description",
					Description: `The description of the admin consent.`,
				},
				resource.Attribute{
					Name:        "admin_consent_display_name",
					Description: `The display name of the admin consent.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `Is this permission enabled?`,
				},
				resource.Attribute{
					Name:        "user_consent_description",
					Description: `The description of the user consent.`,
				},
				resource.Attribute{
					Name:        "user_consent_display_name",
					Description: `The display name of the user consent.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The name of this permission. ## Import Azure Active Directory Applications can be imported using the ` + "`" + `object id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_application.test 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "application_id",
					Description: `The Application ID.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The Application's Object ID.`,
				},
				resource.Attribute{
					Name:        "oauth2_permissions",
					Description: `A collection of OAuth 2.0 permission scopes that the web API (resource) app exposes to client apps. Each permission is covered by a ` + "`" + `oauth2_permission` + "`" + ` block as documented below. --- ` + "`" + `oauth2_permission` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier for one of the ` + "`" + `OAuth2Permission` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the permission.`,
				},
				resource.Attribute{
					Name:        "admin_consent_description",
					Description: `The description of the admin consent.`,
				},
				resource.Attribute{
					Name:        "admin_consent_display_name",
					Description: `The display name of the admin consent.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `Is this permission enabled?`,
				},
				resource.Attribute{
					Name:        "user_consent_description",
					Description: `The description of the user consent.`,
				},
				resource.Attribute{
					Name:        "user_consent_display_name",
					Description: `The display name of the user consent.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The name of this permission. ## Import Azure Active Directory Applications can be imported using the ` + "`" + `object id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_application.test 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_application_password",
			Category:         "Azure Active Directory Resources",
			ShortDescription: `Manages a Password associated with an Application within Azure Active Directory.`,
			Description:      ``,
			Keywords: []string{
				"azure",
				"active",
				"directory",
				"application",
				"password",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "object_id",
					Description: `(Required) The Object ID of the Application for which this password should be created. Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The Password for this Application .`,
				},
				resource.Attribute{
					Name:        "end_date",
					Description: `(Optional) The End Date which the Password is valid until, formatted as a RFC3339 date string (e.g. ` + "`" + `2018-01-01T01:02:03Z` + "`" + `). Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "end_date_relative",
					Description: `(Optional) A relative duration for which the Password is valid until, for example ` + "`" + `240h` + "`" + ` (10 days) or ` + "`" + `2400h30m` + "`" + `. Changing this field forces a new resource to be created. ->`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `(Optional) A GUID used to uniquely identify this Password. If not specified a GUID will be created. Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "start_date",
					Description: `(Optional) The Start Date which the Password is valid from, formatted as a RFC3339 date string (e.g. ` + "`" + `2018-01-01T01:02:03Z` + "`" + `). If this isn't specified, the current date is used. Changing this field forces a new resource to be created. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Key ID for the Password. ## Import Passwords can be imported using the ` + "`" + `object id` + "`" + ` of an Application, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_application_password.test 00000000-0000-0000-0000-000000000000/11111111-1111-1111-1111-111111111111 ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The Key ID for the Password. ## Import Passwords can be imported using the ` + "`" + `object id` + "`" + ` of an Application, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_application_password.test 00000000-0000-0000-0000-000000000000/11111111-1111-1111-1111-111111111111 ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_group",
			Category:         "Azure Active Directory Resources",
			ShortDescription: `Manages a Group within Azure Active Directory.`,
			Description:      ``,
			Keywords: []string{
				"azure",
				"active",
				"directory",
				"group",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The display name for the Group. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Object ID of the Group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The Display Name of the Group. ## Import Azure Active Directory Groups can be imported using the ` + "`" + `object id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_group.my_group 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The Object ID of the Group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The Display Name of the Group. ## Import Azure Active Directory Groups can be imported using the ` + "`" + `object id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_group.my_group 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_service_principal",
			Category:         "Azure Active Directory Resources",
			ShortDescription: `Manages a Service Principal associated with an Application within Azure Active Directory.`,
			Description:      ``,
			Keywords: []string{
				"azure",
				"active",
				"directory",
				"service",
				"principal",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "application_id",
					Description: `(Required) The ID of the Azure AD Application for which to create a Service Principal.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A list of tags to apply to the Service Principal. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Object ID (internal ID) for the Service Principal.`,
				},
				resource.Attribute{
					Name:        "application_id",
					Description: `The Application ID (appId) for the Service Principal.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The Service Principal's Object ID.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The Display Name of the Azure Active Directory Application associated with this Service Principal. ## Import Azure Active Directory Service Principals can be imported using the ` + "`" + `object id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_service_principal.test 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The Object ID (internal ID) for the Service Principal.`,
				},
				resource.Attribute{
					Name:        "application_id",
					Description: `The Application ID (appId) for the Service Principal.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The Service Principal's Object ID.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The Display Name of the Azure Active Directory Application associated with this Service Principal. ## Import Azure Active Directory Service Principals can be imported using the ` + "`" + `object id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_service_principal.test 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_service_principal_password",
			Category:         "Azure Active Directory Resources",
			ShortDescription: `Manages a Password associated with a Service Principal within Azure Active Directory.`,
			Description:      ``,
			Keywords: []string{
				"azure",
				"active",
				"directory",
				"service",
				"principal",
				"password",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "service_principal_id",
					Description: `(Required) The ID of the Service Principal for which this password should be created. Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The Password for this Service Principal.`,
				},
				resource.Attribute{
					Name:        "end_date",
					Description: `(Optional) The End Date which the Password is valid until, formatted as a RFC3339 date string (e.g. ` + "`" + `2018-01-01T01:02:03Z` + "`" + `). Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "end_date_relative",
					Description: `(Optional) A relative duration for which the Password is valid until, for example ` + "`" + `240h` + "`" + ` (10 days) or ` + "`" + `2400h30m` + "`" + `. Changing this field forces a new resource to be created. ->`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `(Optional) A GUID used to uniquely identify this Key. If not specified a GUID will be created. Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "start_date",
					Description: `(Optional) The Start Date which the Password is valid from, formatted as a RFC3339 date string (e.g. ` + "`" + `2018-01-01T01:02:03Z` + "`" + `). If this isn't specified, the current date is used. Changing this field forces a new resource to be created. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Key ID for the Service Principal Password. ## Import Service Principal Passwords can be imported using the ` + "`" + `object id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_service_principal_password.test 00000000-0000-0000-0000-000000000000/11111111-1111-1111-1111-111111111111 ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The Key ID for the Service Principal Password. ## Import Service Principal Passwords can be imported using the ` + "`" + `object id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_service_principal_password.test 00000000-0000-0000-0000-000000000000/11111111-1111-1111-1111-111111111111 ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_user",
			Category:         "Azure Active Directory Resources",
			ShortDescription: `Manages a User within Azure Active Directory.`,
			Description:      ``,
			Keywords: []string{
				"azure",
				"active",
				"directory",
				"user",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "user_principal_name",
					Description: `(Required) The User Principal Name of the Azure AD User.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) The name to display in the address book for the user.`,
				},
				resource.Attribute{
					Name:        "account_enabled",
					Description: `(Optional) ` + "`" + `true` + "`" + ` if the account should be enabled, otherwise ` + "`" + `false` + "`" + `. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) The password for the User. The password must satisfy minimum requirements as specified by the password policy. The maximum length is 256 characters.`,
				},
				resource.Attribute{
					Name:        "force_password_change",
					Description: `(Optional) ` + "`" + `true` + "`" + ` if the User is forced to change the password during the next sign-in. Defaults to ` + "`" + `false` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The Object ID of the Azure AD User.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Object ID of the Azure AD User.`,
				},
				resource.Attribute{
					Name:        "mail",
					Description: `The primary email address of the Azure AD User.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "object_id",
					Description: `The Object ID of the Azure AD User.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Object ID of the Azure AD User.`,
				},
				resource.Attribute{
					Name:        "mail",
					Description: `The primary email address of the Azure AD User.`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"azuread_application":                0,
		"azuread_application_password":       1,
		"azuread_group":                      2,
		"azuread_service_principal":          3,
		"azuread_service_principal_password": 4,
		"azuread_user":                       5,
	}
)

func GetResource(r string) (*resouce.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs]
}
