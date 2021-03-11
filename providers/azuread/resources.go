package azuread

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "azuread_application",
			Category:         "Applications",
			ShortDescription: ``,
			Description: `

Manages an Application within Azure Active Directory.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both ` + "`" + `Read and write owned by applications` + "`" + ` and ` + "`" + `Sign in and read user profile` + "`" + ` within the ` + "`" + `Windows Azure Active Directory` + "`" + ` API.

`,
			Keywords: []string{
				"applications",
				"application",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "app_role",
					Description: `(Optional) A collection of ` + "`" + `app_role` + "`" + ` blocks as documented below. For more information https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles`,
				},
				resource.Attribute{
					Name:        "available_to_other_tenants",
					Description: `(Optional) Is this Azure AD Application available to other tenants? Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) The display name for the application.`,
				},
				resource.Attribute{
					Name:        "group_membership_claims",
					Description: `(Optional) Configures the ` + "`" + `groups` + "`" + ` claim issued in a user or OAuth 2.0 access token that the app expects. Defaults to ` + "`" + `SecurityGroup` + "`" + `. Possible values are ` + "`" + `None` + "`" + `, ` + "`" + `SecurityGroup` + "`" + `, ` + "`" + `DirectoryRole` + "`" + `, ` + "`" + `ApplicationGroup` + "`" + ` or ` + "`" + `All` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "homepage",
					Description: `(optional) The URL to the application's home page.`,
				},
				resource.Attribute{
					Name:        "identifier_uris",
					Description: `(Optional) A list of user-defined URI(s) that uniquely identify a Web application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant.`,
				},
				resource.Attribute{
					Name:        "logout_url",
					Description: `(Optional) The URL of the logout page.`,
				},
				resource.Attribute{
					Name:        "oauth2_allow_implicit_flow",
					Description: `(Optional) Does this Azure AD Application allow OAuth2.0 implicit flow tokens? Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "oauth2_permissions",
					Description: `(Optional) A collection of OAuth 2.0 permission scopes that the web API (resource) app exposes to client apps. Each permission is covered by ` + "`" + `oauth2_permissions` + "`" + ` blocks as documented below. ->`,
				},
				resource.Attribute{
					Name:        "optional_claims",
					Description: `(Optional) A collection of ` + "`" + `access_token` + "`" + ` or ` + "`" + `id_token` + "`" + ` blocks as documented below which list the optional claims configured for each token type. For more information see https://docs.microsoft.com/en-us/azure/active-directory/develop/active-directory-optional-claims`,
				},
				resource.Attribute{
					Name:        "owners",
					Description: `(Optional) A list of Azure AD Object IDs that will be granted ownership of the application. Defaults to the Object ID of the caller creating the application. If a list is specified the caller Object ID will no longer be included unless explicitly added to the list.`,
				},
				resource.Attribute{
					Name:        "prevent_duplicate_names",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, will return an error when an existing Application is found with the same name. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "public_client",
					Description: `(Optional) Is this Azure AD Application a public client? Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "reply_urls",
					Description: `(Optional) A list of URLs that user tokens are sent to for sign in, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to.`,
				},
				resource.Attribute{
					Name:        "required_resource_access",
					Description: `(Optional) A collection of ` + "`" + `required_resource_access` + "`" + ` blocks as documented below.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type of an application: ` + "`" + `webapp/api` + "`" + ` or ` + "`" + `native` + "`" + `. Defaults to ` + "`" + `webapp/api` + "`" + `. For ` + "`" + `native` + "`" + ` apps type ` + "`" + `identifier_uris` + "`" + ` property can not not be set. ~>`,
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
					Description: `The source of the claim. If ` + "`" + `source` + "`" + ` is absent, the claim is a predefined optional claim. If ` + "`" + `source` + "`" + ` is ` + "`" + `user` + "`" + `, the value of ` + "`" + `name` + "`" + ` is the extension property from the user object. --- ` + "`" + `app_role` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "allowed_member_types",
					Description: `(Required) Specifies whether this app role definition can be assigned to users and groups by setting to ` + "`" + `User` + "`" + `, or to other applications (that are accessing this application in daemon service scenarios) by setting to ` + "`" + `Application` + "`" + `, or to both.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) Permission help text that appears in the admin app assignment and consent experiences.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) Display name for the permission that appears in the admin consent and app assignment experiences.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier of the app role. This attribute is computed and cannot be specified manually in this block. If you need to specify a custom ` + "`" + `id` + "`" + `, it's recommended to use the [azuread_application_app_role](application_app_role.html) resource.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `(Optional) Determines if the app role is enabled: Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Specifies the value of the roles claim that the application should expect in the authentication and access tokens. --- ` + "`" + `oauth2_permissions` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "admin_consent_description",
					Description: `(Required) Permission help text that appears in the admin consent and app assignment experiences.`,
				},
				resource.Attribute{
					Name:        "admin_consent_display_name",
					Description: `(Required) Display name for the permission that appears in the admin consent and app assignment experiences.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier of the permision. This attribute is computed and cannot be specified manually in this block. If you need to specify a custom ` + "`" + `id` + "`" + `, it's recommended to use the [azuread_application_oauth2_permission](application_oauth2_permission.html) resource.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `(Optional) Determines if the permission is enabled: defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Specifies whether this scope permission can be consented to by an end user, or whether it is a tenant-wide permission that must be consented to by a Company Administrator. Possible values are "User" or "Admin".`,
				},
				resource.Attribute{
					Name:        "user_consent_description",
					Description: `(Optional) Permission help text that appears in the end user consent experience.`,
				},
				resource.Attribute{
					Name:        "user_consent_display_name",
					Description: `(Optional) Display name for the permission that appears in the end user consent experience.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of the scope claim that the resource application should expect in the OAuth 2.0 access token. If you don't specify any ` + "`" + `oauth2_permissions` + "`" + ` blocks, your Application will be assigned the default ` + "`" + `user_impersonation` + "`" + ` scope by Azure Active Directory. However, due to the declarative nature of Terraform configuration, if you do specify any ` + "`" + `oauth2_permissions` + "`" + ` blocks, you will need to include a block for the ` + "`" + `user_impersonation` + "`" + ` scope if you need it, or it will be removed (see the example above). To ensure that no OAuth 2.0 permission scopes are present for your Application, specify ` + "`" + `oauth2_permissions = []` + "`" + ` in your Application resource. --- ` + "`" + `required_resource_access` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "resource_access",
					Description: `(Required) A collection of ` + "`" + `resource_access` + "`" + ` blocks as documented below.`,
				},
				resource.Attribute{
					Name:        "resource_app_id",
					Description: `(Required) The unique identifier for the resource that the application requires access to. This should be equal to the appId declared on the target resource application. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The unique identifier for one of the ` + "`" + `OAuth2Permission` + "`" + ` or ` + "`" + `AppRole` + "`" + ` instances that the resource application exposes.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Specifies whether the id property references an ` + "`" + `OAuth2Permission` + "`" + ` or an ` + "`" + `AppRole` + "`" + `. Possible values are ` + "`" + `Scope` + "`" + ` or ` + "`" + `Role` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "application_id",
					Description: `The Application ID (Client ID).`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The Application's Object ID. ## Import Azure Active Directory Applications can be imported using the ` + "`" + `object id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_application.test 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "application_id",
					Description: `The Application ID (Client ID).`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The Application's Object ID. ## Import Azure Active Directory Applications can be imported using the ` + "`" + `object id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_application.test 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_application_app_role",
			Category:         "Applications",
			ShortDescription: ``,
			Description: `

Manages an App Role associated with an Application within Azure Active Directory.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both ` + "`" + `Read and write all applications` + "`" + ` and ` + "`" + `Sign in and read user profile` + "`" + ` within the ` + "`" + `Windows Azure Active Directory` + "`" + ` API.

`,
			Keywords: []string{
				"applications",
				"application",
				"app",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "allowed_member_types",
					Description: `(Required) Specifies whether this app role definition can be assigned to users and groups by setting to ` + "`" + `User` + "`" + `, or to other applications (that are accessing this application in daemon service scenarios) by setting to ` + "`" + `Application` + "`" + `, or to both.`,
				},
				resource.Attribute{
					Name:        "application_object_id",
					Description: `(Required) The Object ID of the Application for which this App Role should be created. Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) Permission help text that appears in the admin app assignment and consent experiences.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) Display name for the permission that appears in the admin consent and app assignment experiences.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `(Optional) Determines if the app role is enabled. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `(Optional) Specifies a custom UUID for the app role. If omitted, a random UUID will be automatically generated. Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Specifies the value of the roles claim that the application should expect in the authentication and access tokens. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_application_certificate",
			Category:         "Applications",
			ShortDescription: ``,
			Description: `

Manages a Certificate associated with an Application within Azure Active Directory.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both ` + "`" + `Read and write all applications` + "`" + ` and ` + "`" + `Sign in and read user profile` + "`" + ` within the ` + "`" + `Windows Azure Active Directory` + "`" + ` API.

`,
			Keywords: []string{
				"applications",
				"application",
				"certificate",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "application_object_id",
					Description: `(Required) The Object ID of the Application for which this Certificate should be created. Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "encoding",
					Description: `(Optional) Specifies the encoding used for the supplied certificate data. Must be one of ` + "`" + `pem` + "`" + `, ` + "`" + `base64` + "`" + ` or ` + "`" + `hex` + "`" + `. Defaults to ` + "`" + `pem` + "`" + `. ->`,
				},
				resource.Attribute{
					Name:        "end_date",
					Description: `(Optional) The End Date which the Certificate is valid until, formatted as a RFC3339 date string (e.g. ` + "`" + `2018-01-01T01:02:03Z` + "`" + `). Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "end_date_relative",
					Description: `(Optional) A relative duration for which the Certificate is valid until, for example ` + "`" + `240h` + "`" + ` (10 days) or ` + "`" + `2400h30m` + "`" + `. Changing this field forces a new resource to be created. ~>`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `(Optional) A GUID used to uniquely identify this Certificate. If not specified a GUID will be created. Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "start_date",
					Description: `(Optional) The Start Date which the Certificate is valid from, formatted as a RFC3339 date string (e.g. ` + "`" + `2018-01-01T01:02:03Z` + "`" + `). If this isn't specified, the current date is used. Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of key/certificate. Must be one of ` + "`" + `AsymmetricX509Cert` + "`" + ` or ` + "`" + `Symmetric` + "`" + `. Changing this fields forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The certificate data, which can be PEM encoded, base64 encoded DER or hexadecimal encoded DER. See also the ` + "`" + `encoding` + "`" + ` argument. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_application_oauth2_permission",
			Category:         "Applications",
			ShortDescription: ``,
			Description: `

Manages an OAuth2 Permission (also known as a Scope) associated with an Application within Azure Active Directory.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both ` + "`" + `Read and write all applications` + "`" + ` and ` + "`" + `Sign in and read user profile` + "`" + ` within the ` + "`" + `Windows Azure Active Directory` + "`" + ` API.

`,
			Keywords: []string{
				"applications",
				"application",
				"oauth2",
				"permission",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "admin_consent_description",
					Description: `(Required) Permission help text that appears in the admin consent and app assignment experiences.`,
				},
				resource.Attribute{
					Name:        "admin_consent_display_name",
					Description: `(Required) Display name for the permission that appears in the admin consent and app assignment experiences.`,
				},
				resource.Attribute{
					Name:        "application_object_id",
					Description: `(Required) The Object ID of the Application for which this Permission should be created. Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `(Optional) Determines if the Permission is enabled. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "permission_id",
					Description: `(Optional) Specifies a custom UUID for the Permission. If omitted, a random UUID will be automatically generated. Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Specifies whether this scope permission can be consented to by an end user, or whether it is a tenant-wide permission that must be consented to by an Administrator. Possible values are "User" or "Admin".`,
				},
				resource.Attribute{
					Name:        "user_consent_description",
					Description: `(Optional) Permission help text that appears in the end user consent experience.`,
				},
				resource.Attribute{
					Name:        "user_consent_display_name",
					Description: `(Optional) Display name for the permission that appears in the end user consent experience.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of the scope claim that the resource application should expect in the OAuth 2.0 access token. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_application_password",
			Category:         "Applications",
			ShortDescription: ``,
			Description: `

Manages a Password associated with an Application within Azure Active Directory. Also can be referred to as Client secrets.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both ` + "`" + `Read and write all applications` + "`" + ` and ` + "`" + `Sign in and read user profile` + "`" + ` within the ` + "`" + `Windows Azure Active Directory` + "`" + ` API.

`,
			Keywords: []string{
				"applications",
				"application",
				"password",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "application_object_id",
					Description: `(Required) The Object ID of the Application for which this password should be created. Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description for the Password. ->`,
				},
				resource.Attribute{
					Name:        "end_date",
					Description: `(Optional) The End Date which the Password is valid until, formatted as a RFC3339 date string (e.g. ` + "`" + `2018-01-01T01:02:03Z` + "`" + `). Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "end_date_relative",
					Description: `(Optional) A relative duration for which the Password is valid until, for example ` + "`" + `240h` + "`" + ` (10 days) or ` + "`" + `2400h30m` + "`" + `. Changing this field forces a new resource to be created. ~>`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `(Optional) A GUID used to uniquely identify this Password. If not specified a GUID will be created. Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "start_date",
					Description: `(Optional) The Start Date which the Password is valid from, formatted as a RFC3339 date string (e.g. ` + "`" + `2018-01-01T01:02:03Z` + "`" + `). If this isn't specified, the current date is used. Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The Password for this Application. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_group",
			Category:         "Groups",
			ShortDescription: ``,
			Description: `

Manages a Group within Azure Active Directory.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to ` + "`" + `Read and write all groups` + "`" + ` within the ` + "`" + `Windows Azure Active Directory` + "`" + ` API. In addition it must also have either the ` + "`" + `Groups Administrator` + "`" + ` or ` + "`" + `User Administrator` + "`" + ` Azure Active Directory roles assigned in order to be able to delete groups. You can assign one of the required Azure Active Directory Roles with the **AzureAD PowerShell Module**, which is available for Windows PowerShell or in the Azure Cloud Shell. Please refer to [this documentation](https://docs.microsoft.com/en-us/powershell/module/azuread/add-azureaddirectoryrolemember) for more details.

`,
			Keywords: []string{
				"groups",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description for the Group. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) The display name for the Group. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `(Optional) A set of members who should be present in this Group. Supported Object types are Users, Groups or Service Principals.`,
				},
				resource.Attribute{
					Name:        "owners",
					Description: `(Optional) A set of owners who own this Group. Supported Object types are Users or Service Principals.`,
				},
				resource.Attribute{
					Name:        "prevent_duplicate_names",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, will return an error when an existing Group is found with the same name. Defaults to ` + "`" + `false` + "`" + `. ->`,
				},
				resource.Attribute{
					Name:        "mail_enabled",
					Description: `Whether the group is mail-enabled.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The Object ID of the Group.`,
				},
				resource.Attribute{
					Name:        "security_enabled",
					Description: `Whether the group is a security group. ~>`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "mail_enabled",
					Description: `Whether the group is mail-enabled.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The Object ID of the Group.`,
				},
				resource.Attribute{
					Name:        "security_enabled",
					Description: `Whether the group is a security group. ~>`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_group_member",
			Category:         "Groups",
			ShortDescription: ``,
			Description: `

Manages a single Group Membership within Azure Active Directory.

-> **NOTE:** Do not use this resource at the same time as ` + "`" + `azuread_group.members` + "`" + `.

`,
			Keywords: []string{
				"groups",
				"group",
				"member",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group_object_id",
					Description: `(Required) The Object ID of the Azure AD Group you want to add the Member to. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "member_object_id",
					Description: `(Required) The Object ID of the Azure AD Object you want to add as a Member to the Group. Supported Object types are Users, Groups or Service Principals. Changing this forces a new resource to be created. ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_service_principal",
			Category:         "Service Principals",
			ShortDescription: ``,
			Description: `

Manages a Service Principal associated with an Application within Azure Active Directory.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both ` + "`" + `Read and write all applications` + "`" + ` and ` + "`" + `Sign in and read user profile` + "`" + ` within the ` + "`" + `Windows Azure Active Directory` + "`" + ` API. Please see The [Granting a Service Principal permission to manage AAD](../guides/service_principal_configuration.html) for the required steps.

`,
			Keywords: []string{
				"service",
				"principals",
				"principal",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "app_role_assignment_required",
					Description: `(Optional) Whether this Service Principal requires an AppRoleAssignment to a user or group before Azure AD will issue a user or access token to the application. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "application_id",
					Description: `(Required) The App ID of the Application for which to create a Service Principal.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A list of tags to apply to the Service Principal. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The Display Name of the Application associated with this Service Principal.`,
				},
				resource.Attribute{
					Name:        "oauth2_permissions",
					Description: `A collection of OAuth 2.0 permissions exposed by the associated Application. Each permission is covered by an ` + "`" + `oauth2_permission` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The Object ID of the Service Principal. --- ` + "`" + `oauth2_permission` + "`" + ` block exports the following:`,
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
					Name:        "id",
					Description: `The unique identifier for one of the ` + "`" + `OAuth2Permission` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `Is this permission enabled?`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the permission.`,
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
					Description: `The name of this permission. ## Import Azure Active Directory Service Principals can be imported using the ` + "`" + `object id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_service_principal.test 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `The Display Name of the Application associated with this Service Principal.`,
				},
				resource.Attribute{
					Name:        "oauth2_permissions",
					Description: `A collection of OAuth 2.0 permissions exposed by the associated Application. Each permission is covered by an ` + "`" + `oauth2_permission` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The Object ID of the Service Principal. --- ` + "`" + `oauth2_permission` + "`" + ` block exports the following:`,
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
					Name:        "id",
					Description: `The unique identifier for one of the ` + "`" + `OAuth2Permission` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `Is this permission enabled?`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the permission.`,
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
					Description: `The name of this permission. ## Import Azure Active Directory Service Principals can be imported using the ` + "`" + `object id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_service_principal.test 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_service_principal_certificate",
			Category:         "Service Principals",
			ShortDescription: ``,
			Description: `

Manages a Certificate associated with a Service Principal within Azure Active Directory.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both ` + "`" + `Read and write all applications` + "`" + ` and ` + "`" + `Sign in and read user profile` + "`" + ` within the ` + "`" + `Windows Azure Active Directory` + "`" + ` API.

`,
			Keywords: []string{
				"service",
				"principals",
				"principal",
				"certificate",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "encoding",
					Description: `(Optional) Specifies the encoding used for the supplied certificate data. Must be one of ` + "`" + `pem` + "`" + `, ` + "`" + `base64` + "`" + ` or ` + "`" + `hex` + "`" + `. Defaults to ` + "`" + `pem` + "`" + `. ->`,
				},
				resource.Attribute{
					Name:        "end_date",
					Description: `(Optional) The End Date which the Certificate is valid until, formatted as a RFC3339 date string (e.g. ` + "`" + `2018-01-01T01:02:03Z` + "`" + `). Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "end_date_relative",
					Description: `(Optional) A relative duration for which the Certificate is valid until, for example ` + "`" + `240h` + "`" + ` (10 days) or ` + "`" + `2400h30m` + "`" + `. Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h". Changing this field forces a new resource to be created. ~>`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `(Optional) A GUID used to uniquely identify this Certificate. If not specified a GUID will be created. Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "service_principal_id",
					Description: `(Required) The ID of the Service Principal for which this certificate should be created. Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "start_date",
					Description: `(Optional) The Start Date which the Certificate is valid from, formatted as a RFC3339 date string (e.g. ` + "`" + `2018-01-01T01:02:03Z` + "`" + `). If this isn't specified, the current date is used. Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of key/certificate. Must be one of ` + "`" + `AsymmetricX509Cert` + "`" + ` or ` + "`" + `Symmetric` + "`" + `. Changing this fields forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The certificate data, which can be PEM encoded, base64 encoded DER or hexadecimal encoded DER. See also the ` + "`" + `encoding` + "`" + ` argument. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_service_principal_password",
			Category:         "Service Principals",
			ShortDescription: ``,
			Description: `

Manages a Password associated with a Service Principal within Azure Active Directory.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both ` + "`" + `Read and write all applications` + "`" + ` and ` + "`" + `Sign in and read user profile` + "`" + ` within the ` + "`" + `Windows Azure Active Directory` + "`" + ` API.

`,
			Keywords: []string{
				"service",
				"principals",
				"principal",
				"password",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description for the Password. ->`,
				},
				resource.Attribute{
					Name:        "end_date",
					Description: `(Optional) The End Date which the Password is valid until, formatted as a RFC3339 date string (e.g. ` + "`" + `2018-01-01T01:02:03Z` + "`" + `). Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "end_date_relative",
					Description: `(Optional) A relative duration for which the Password is valid until, for example ` + "`" + `240h` + "`" + ` (10 days) or ` + "`" + `2400h30m` + "`" + `. Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h". Changing this field forces a new resource to be created. ~>`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `(Optional) A GUID used to uniquely identify this Key. If not specified a GUID will be created. Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "service_principal_id",
					Description: `(Required) The ID of the Service Principal for which this password should be created. Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "start_date",
					Description: `(Optional) The Start Date which the Password is valid from, formatted as a RFC3339 date string (e.g. ` + "`" + `2018-01-01T01:02:03Z` + "`" + `). If this isn't specified, the current date is used. Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The Password for this Service Principal. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_user",
			Category:         "Users",
			ShortDescription: ``,
			Description: `

Manages a User within Azure Active Directory.

-> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to ` + "`" + `Directory.ReadWrite.All` + "`" + ` within the ` + "`" + `Windows Azure Active Directory` + "`" + ` API.

`,
			Keywords: []string{
				"users",
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_enabled",
					Description: `(Optional) ` + "`" + `true` + "`" + ` if the account should be enabled, otherwise ` + "`" + `false` + "`" + `. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "city",
					Description: `(Optional) The city in which the user is located.`,
				},
				resource.Attribute{
					Name:        "company_name",
					Description: `(Optional) The company name which the user is associated. This property can be useful for describing the company that an external user comes from.`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `(Optional) The country/region in which the user is located; for example, “US” or “UK”.`,
				},
				resource.Attribute{
					Name:        "department",
					Description: `(Optional) The name for the department in which the user works.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) The name to display in the address book for the user.`,
				},
				resource.Attribute{
					Name:        "force_password_change",
					Description: `(Optional) ` + "`" + `true` + "`" + ` if the User is forced to change the password during the next sign-in. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "given_name",
					Description: `(Optional) The given name (first name) of the user.`,
				},
				resource.Attribute{
					Name:        "immutable_id",
					Description: `(Optional) The value used to associate an on-premise Active Directory user account with their Azure AD user object. This must be specified if you are using a federated domain for the user's userPrincipalName (UPN) property when creating a new user account.`,
				},
				resource.Attribute{
					Name:        "job_title",
					Description: `(Optional) The user’s job title.`,
				},
				resource.Attribute{
					Name:        "mail_nickname",
					Description: `(Optional) The mail alias for the user. Defaults to the user name part of the User Principal Name.`,
				},
				resource.Attribute{
					Name:        "mobile",
					Description: `(Optional) The primary cellular telephone number for the user.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) The password for the User. The password must satisfy minimum requirements as specified by the password policy. The maximum length is 256 characters.`,
				},
				resource.Attribute{
					Name:        "physical_delivery_office_name",
					Description: `(Optional) The office location in the user's place of business.`,
				},
				resource.Attribute{
					Name:        "postal_code",
					Description: `(Optional) The postal code for the user's postal address. The postal code is specific to the user's country/region. In the United States of America, this attribute contains the ZIP code.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) The state or province in the user's address.`,
				},
				resource.Attribute{
					Name:        "street_address",
					Description: `(Optional) The street address of the user's place of business.`,
				},
				resource.Attribute{
					Name:        "surname",
					Description: `(Optional) The user's surname (family name or last name).`,
				},
				resource.Attribute{
					Name:        "usage_location",
					Description: `(Optional) The usage location of the User. Required for users that will be assigned licenses due to legal requirement to check for availability of services in countries. The usage location is a two letter country code (ISO standard 3166). Examples include: ` + "`" + `NO` + "`" + `, ` + "`" + `JP` + "`" + `, and ` + "`" + `GB` + "`" + `. Cannot be reset to null once set.`,
				},
				resource.Attribute{
					Name:        "user_principal_name",
					Description: `(Required) The User Principal Name of the User. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "mail",
					Description: `The primary email address of the User.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The Object ID of the User.`,
				},
				resource.Attribute{
					Name:        "onpremises_sam_account_name",
					Description: `The on-premise SAM account name of the User.`,
				},
				resource.Attribute{
					Name:        "onpremises_user_principal_name",
					Description: `The on-premise user principal name of the User.`,
				},
				resource.Attribute{
					Name:        "user_type",
					Description: `The user type in the directory. One of ` + "`" + `Guest` + "`" + ` or ` + "`" + `Member` + "`" + `. ## Import Azure Active Directory Users can be imported using the ` + "`" + `object id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_user.my_user 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "mail",
					Description: `The primary email address of the User.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The Object ID of the User.`,
				},
				resource.Attribute{
					Name:        "onpremises_sam_account_name",
					Description: `The on-premise SAM account name of the User.`,
				},
				resource.Attribute{
					Name:        "onpremises_user_principal_name",
					Description: `The on-premise user principal name of the User.`,
				},
				resource.Attribute{
					Name:        "user_type",
					Description: `The user type in the directory. One of ` + "`" + `Guest` + "`" + ` or ` + "`" + `Member` + "`" + `. ## Import Azure Active Directory Users can be imported using the ` + "`" + `object id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_user.my_user 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"azuread_application":                   0,
		"azuread_application_app_role":          1,
		"azuread_application_certificate":       2,
		"azuread_application_oauth2_permission": 3,
		"azuread_application_password":          4,
		"azuread_group":                         5,
		"azuread_group_member":                  6,
		"azuread_service_principal":             7,
		"azuread_service_principal_certificate": 8,
		"azuread_service_principal_password":    9,
		"azuread_user":                          10,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
