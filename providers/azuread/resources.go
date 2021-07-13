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

Manages an application registration within Azure Active Directory.

`,
			Keywords: []string{
				"applications",
				"application",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "api",
					Description: `(Optional) An ` + "`" + `api` + "`" + ` block as documented below, which configures API related settings for this Application.`,
				},
				resource.Attribute{
					Name:        "app_role",
					Description: `(Optional) A collection of ` + "`" + `app_role` + "`" + ` blocks as documented below. For more information see [official documentation on Application Roles](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) The display name for the application.`,
				},
				resource.Attribute{
					Name:        "fallback_public_client_enabled",
					Description: `(Optional) Specifies whether the application is a public client. Appropriate for apps using token grant flows that don't use a redirect URI. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "group_membership_claims",
					Description: `(Optional) Configures the ` + "`" + `groups` + "`" + ` claim issued in a user or OAuth 2.0 access token that the app expects. Possible values are ` + "`" + `None` + "`" + `, ` + "`" + `SecurityGroup` + "`" + `, ` + "`" + `DirectoryRole` + "`" + `, ` + "`" + `ApplicationGroup` + "`" + ` or ` + "`" + `All` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "identifier_uris",
					Description: `(Optional) The user-defined URI(s) that uniquely identify an application within its Azure AD tenant, or within a verified custom domain if the application is multi-tenant.`,
				},
				resource.Attribute{
					Name:        "optional_claims",
					Description: `(Optional) An ` + "`" + `optional_claims` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "owners",
					Description: `(Optional) A list of object IDs of principals that will be granted ownership of the application. It's recommended to specify the object ID of the authenticated principal running Terraform, to ensure sufficient permissions that the application can be subsequently updated.`,
				},
				resource.Attribute{
					Name:        "prevent_duplicate_names",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, will return an error if an existing application is found with the same name. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "required_resource_access",
					Description: `(Optional) A collection of ` + "`" + `required_resource_access` + "`" + ` blocks as documented below.`,
				},
				resource.Attribute{
					Name:        "sign_in_audience",
					Description: `(Optional) The Microsoft account types that are supported for the current application. Must be one of ` + "`" + `AzureADMyOrg` + "`" + `, ` + "`" + `AzureADMultipleOrgs` + "`" + `, ` + "`" + `AzureADandPersonalMicrosoftAccount` + "`" + ` or ` + "`" + `PersonalMicrosoftAccount` + "`" + `. Defaults to ` + "`" + `AzureADMyOrg` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "web",
					Description: `(Optional) A ` + "`" + `web` + "`" + ` block as documented below, which configures web related settings for this Application. ->`,
				},
				resource.Attribute{
					Name:        "oauth2_permission_scope",
					Description: `(Optional) One or more ` + "`" + `oauth2_permission_scope` + "`" + ` blocks as documented below, to describe delegated permissions exposed by the web API represented by this application. --- ` + "`" + `oauth2_permission_scope` + "`" + ` blocks support the following:`,
				},
				resource.Attribute{
					Name:        "admin_consent_description",
					Description: `(Required) Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.`,
				},
				resource.Attribute{
					Name:        "admin_consent_display_name",
					Description: `(Required) Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Determines if the permission scope is enabled. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The unique identifier of the delegated permission. Must be a valid UUID. ->`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Defaults to ` + "`" + `User` + "`" + `. Possible values are ` + "`" + `User` + "`" + ` or ` + "`" + `Admin` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user_consent_description",
					Description: `(Optional) Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.`,
				},
				resource.Attribute{
					Name:        "user_consent_display_name",
					Description: `(Optional) Display name for the delegated permission that appears in the end user consent experience.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) The value that is used for the ` + "`" + `scp` + "`" + ` claim in OAuth 2.0 access tokens. ~>`,
				},
				resource.Attribute{
					Name:        "allowed_member_types",
					Description: `(Required) Specifies whether this app role definition can be assigned to users and groups by setting to ` + "`" + `User` + "`" + `, or to other applications (that are accessing this application in a standalone scenario) by setting to ` + "`" + `Application` + "`" + `, or to both.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) Description of the app role that appears when the role is being assigned and, if the role functions as an application permissions, during the consent experiences.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) Display name for the app role that appears during app role assignment and in consent experiences.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Determines if the app role is enabled. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The unique identifier of the app role. Must be a valid UUID. ->`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) The value that is used for the ` + "`" + `roles` + "`" + ` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal. ->`,
				},
				resource.Attribute{
					Name:        "access_token",
					Description: `(Optional) One or more ` + "`" + `access_token` + "`" + ` blocks as documented below.`,
				},
				resource.Attribute{
					Name:        "id_token",
					Description: `(Optional) One or more ` + "`" + `id_token` + "`" + ` blocks as documented below.`,
				},
				resource.Attribute{
					Name:        "saml2_token",
					Description: `(Optional) One or more ` + "`" + `saml2_token` + "`" + ` blocks as documented below. --- ` + "`" + `access_token` + "`" + `, ` + "`" + `id_token` + "`" + ` and ` + "`" + `saml2_token` + "`" + ` blocks support the following:`,
				},
				resource.Attribute{
					Name:        "additional_properties",
					Description: `List of additional properties of the claim. If a property exists in this list, it modifies the behaviour of the optional claim.`,
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
					Description: `The source of the claim. If ` + "`" + `source` + "`" + ` is absent, the claim is a predefined optional claim. If ` + "`" + `source` + "`" + ` is ` + "`" + `user` + "`" + `, the value of ` + "`" + `name` + "`" + ` is the extension property from the user object. --- ` + "`" + `required_resource_access` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "resource_access",
					Description: `(Required) A collection of ` + "`" + `resource_access` + "`" + ` blocks as documented below, describing OAuth2.0 permission scopes and app roles that the application requires from the specified resource.`,
				},
				resource.Attribute{
					Name:        "resource_app_id",
					Description: `(Required) The unique identifier for the resource that the application requires access to. This should be the Application ID of the target application. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The unique identifier for an app role or OAuth2 permission scope published by the resource application.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Specifies whether the ` + "`" + `id` + "`" + ` property references an app role or an OAuth2 permission scope. Possible values are ` + "`" + `Role` + "`" + ` or ` + "`" + `Scope` + "`" + `. --- ` + "`" + `web` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "homepage_url",
					Description: `(Optional) Home page or landing page of the application.`,
				},
				resource.Attribute{
					Name:        "implicit_grant",
					Description: `(Optional) An ` + "`" + `implicit_grant` + "`" + ` block as documented above.`,
				},
				resource.Attribute{
					Name:        "logout_url",
					Description: `(Optional) The URL that will be used by Microsoft's authorization service to sign out a user using front-channel, back-channel or SAML logout protocols.`,
				},
				resource.Attribute{
					Name:        "redirect_uris",
					Description: `(Optional) A list of URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent. --- ` + "`" + `implicit_grant` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "access_token_issuance_enabled",
					Description: `(Optional) Whether this web application can request an access token using OAuth 2.0 implicit flow.`,
				},
				resource.Attribute{
					Name:        "id_token_issuance_enabled",
					Description: `(Optional) Whether this web application can request an ID token using OAuth 2.0 implicit flow. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "application_id",
					Description: `The Application ID (also called Client ID).`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The application's object ID. ## Import Applications can be imported using their object ID, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_application.test 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "application_id",
					Description: `The Application ID (also called Client ID).`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The application's object ID. ## Import Applications can be imported using their object ID, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_application.test 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_application_certificate",
			Category:         "Applications",
			ShortDescription: ``,
			Description: `

Manages a certificate associated with an application within Azure Active Directory. These are also referred to as client certificates during authentication.

`,
			Keywords: []string{
				"applications",
				"application",
				"certificate",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "application_object_id",
					Description: `(Required) The object ID of the application for which this certificate should be created. Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "encoding",
					Description: `(Optional) Specifies the encoding used for the supplied certificate data. Must be one of ` + "`" + `pem` + "`" + `, ` + "`" + `base64` + "`" + ` or ` + "`" + `hex` + "`" + `. Defaults to ` + "`" + `pem` + "`" + `. ->`,
				},
				resource.Attribute{
					Name:        "end_date",
					Description: `(Optional) The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. ` + "`" + `2018-01-01T01:02:03Z` + "`" + `). If omitted, the API will decide a suitable expiry date, which is typically around 2 years from the start date. Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "end_date_relative",
					Description: `(Optional) A relative duration for which the certificate is valid until, for example ` + "`" + `240h` + "`" + ` (10 days) or ` + "`" + `2400h30m` + "`" + `. Changing this field forces a new resource to be created. ~>`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `(Optional) A UUID used to uniquely identify this certificate. If omitted, a random UUID will be automatically generated. Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "start_date",
					Description: `(Optional) The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. ` + "`" + `2018-01-01T01:02:03Z` + "`" + `). If this isn't specified, the current date and time are used. Changing this field forces a new resource to be created.`,
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
			Type:             "azuread_application_password",
			Category:         "Applications",
			ShortDescription: ``,
			Description: `

Manages a password credential associated with an application within Azure Active Directory. These are also referred to as client secrets during authentication.

`,
			Keywords: []string{
				"applications",
				"application",
				"password",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "application_object_id",
					Description: `(Required) The object ID of the application for which this password should be created. Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) A display name for the password.`,
				},
				resource.Attribute{
					Name:        "end_date",
					Description: `(Optional) The end date until which the password is valid, formatted as an RFC3339 date string (e.g. ` + "`" + `2018-01-01T01:02:03Z` + "`" + `). Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "end_date_relative",
					Description: `(Optional) A relative duration for which the password is valid until, for example ` + "`" + `240h` + "`" + ` (10 days) or ` + "`" + `2400h30m` + "`" + `. Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "start_date",
					Description: `(Optional) The start date from which the password is valid, formatted as an RFC3339 date string (e.g. ` + "`" + `2018-01-01T01:02:03Z` + "`" + `). If this isn't specified, the current date is used. Changing this field forces a new resource to be created. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `A UUID used to uniquely identify this password credential.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The password for this application, which is generated by Azure Active Directory. ## Import This resource does not support importing.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "key_id",
					Description: `A UUID used to uniquely identify this password credential.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The password for this application, which is generated by Azure Active Directory. ## Import This resource does not support importing.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_application_pre_authorized",
			Category:         "Applications",
			ShortDescription: ``,
			Description: `

Manages client applications that are pre-authorized with the specified permissions to access an application's APIs without requiring user consent.

`,
			Keywords: []string{
				"applications",
				"application",
				"pre",
				"authorized",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "application_object_id",
					Description: `(Required) The object ID of the application for which permissions are being authorized. Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "authorized_application_id",
					Description: `(Optional) The application ID (client ID) of the application being authorized. Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "permission_ids",
					Description: `(Required) A set of permission scope IDs required by the authorized application. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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

Manages a group within Azure Active Directory.

`,
			Keywords: []string{
				"groups",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description for the group.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) The display name for the group.`,
				},
				resource.Attribute{
					Name:        "mail_enabled",
					Description: `(Optional) Whether the group is a mail enabled, with a shared group mailbox. At least one of ` + "`" + `mail_enabled` + "`" + ` or ` + "`" + `security_enabled` + "`" + ` must be specified. A group can be mail enabled _and_ security enabled.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `(Optional) A set of members who should be present in this group. Supported object types are Users, Groups or Service Principals.`,
				},
				resource.Attribute{
					Name:        "owners",
					Description: `(Optional) A set of owners who own this group. Supported object types are Users or Service Principals.`,
				},
				resource.Attribute{
					Name:        "prevent_duplicate_names",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, will return an error if an existing group is found with the same name. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "security_enabled",
					Description: `(Optional) Whether the group is a security group for controlling access to in-app resources. At least one of ` + "`" + `security_enabled` + "`" + ` or ` + "`" + `mail_enabled` + "`" + ` must be specified. A group can be security enabled _and_ mail enabled.`,
				},
				resource.Attribute{
					Name:        "types",
					Description: `(Optional) A set of group types to configure for the group. The only supported type is ` + "`" + `Unified` + "`" + `, which specifies a Microsoft 365 group. Required when ` + "`" + `mail_enabled` + "`" + ` is true. Changing this forces a new resource to be created. ->`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The object ID of the group. ## Import Groups can be imported using their object ID, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_group.my_group 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "object_id",
					Description: `The object ID of the group. ## Import Groups can be imported using their object ID, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_group.my_group 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_group_member",
			Category:         "Groups",
			ShortDescription: ``,
			Description: `

Manages a single group membership within Azure Active Directory.

-> **Warning** Do not use this resource at the same time as the ` + "`" + `members` + "`" + ` property of the ` + "`" + `azuread_group` + "`" + ` resource.

`,
			Keywords: []string{
				"groups",
				"group",
				"member",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group_object_id",
					Description: `(Required) The object ID of the group you want to add the member to. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "member_object_id",
					Description: `(Required) The object ID of the principal you want to add as a member to the group. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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

Manages a service principal associated with an application within Azure Active Directory.

`,
			Keywords: []string{
				"service",
				"principals",
				"principal",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "app_role_assignment_required",
					Description: `(Optional) Whether this service principal requires an app role assignment to a user or group before Azure AD will issue a user or access token to the application. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "application_id",
					Description: `(Required) The application ID (client ID) of the application for which to create a service principal.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A set of tags to apply to the service principal. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "app_roles",
					Description: `A list of app roles published b the associated application, as documented below. For more information [official documentation](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name of the application associated with this service principal.`,
				},
				resource.Attribute{
					Name:        "oauth2_permission_scopes",
					Description: `A list of OAuth 2.0 delegated permission scopes published by the associated application, as documented below.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The object ID of the service principal. --- ` + "`" + `app_roles` + "`" + ` is a list of objects with the following attributes:`,
				},
				resource.Attribute{
					Name:        "allowed_member_types",
					Description: `Specifies whether this app role definition can be assigned to users and groups, or to other applications (that are accessing this application in a standalone scenario). Possible values are: ` + "`" + `User` + "`" + ` and ` + "`" + `Application` + "`" + `, or both.`,
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
					Description: `The unique identifier of the ` + "`" + `app_role` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value that is used for the ` + "`" + `roles` + "`" + ` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal. --- ` + "`" + `oauth2_permission_scopes` + "`" + ` is a list of objects with the following attributes:`,
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
					Description: `Specifies whether the permission scope is enabled.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier of the delegated permission.`,
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
					Description: `The value that is used for the ` + "`" + `scp` + "`" + ` claim in OAuth 2.0 access tokens. ## Import Service principals can be imported using their object ID, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_service_principal.test 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "app_roles",
					Description: `A list of app roles published b the associated application, as documented below. For more information [official documentation](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name of the application associated with this service principal.`,
				},
				resource.Attribute{
					Name:        "oauth2_permission_scopes",
					Description: `A list of OAuth 2.0 delegated permission scopes published by the associated application, as documented below.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The object ID of the service principal. --- ` + "`" + `app_roles` + "`" + ` is a list of objects with the following attributes:`,
				},
				resource.Attribute{
					Name:        "allowed_member_types",
					Description: `Specifies whether this app role definition can be assigned to users and groups, or to other applications (that are accessing this application in a standalone scenario). Possible values are: ` + "`" + `User` + "`" + ` and ` + "`" + `Application` + "`" + `, or both.`,
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
					Description: `The unique identifier of the ` + "`" + `app_role` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value that is used for the ` + "`" + `roles` + "`" + ` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal. --- ` + "`" + `oauth2_permission_scopes` + "`" + ` is a list of objects with the following attributes:`,
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
					Description: `Specifies whether the permission scope is enabled.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier of the delegated permission.`,
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
					Description: `The value that is used for the ` + "`" + `scp` + "`" + ` claim in OAuth 2.0 access tokens. ## Import Service principals can be imported using their object ID, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_service_principal.test 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_service_principal_certificate",
			Category:         "Service Principals",
			ShortDescription: ``,
			Description: `

Manages a certificate associated with a service principal within Azure Active Directory.

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
					Description: `(Optional) The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. ` + "`" + `2018-01-01T01:02:03Z` + "`" + `). Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "end_date_relative",
					Description: `(Optional) A relative duration for which the certificate is valid until, for example ` + "`" + `240h` + "`" + ` (10 days) or ` + "`" + `2400h30m` + "`" + `. Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h". Changing this field forces a new resource to be created. ~>`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `(Optional) A UUID used to uniquely identify this certificate. If not specified a UUID will be automatically generated. Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "service_principal_id",
					Description: `(Required) The object ID of the service principal for which this certificate should be created. Changing this field forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "start_date",
					Description: `(Optional) The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. ` + "`" + `2018-01-01T01:02:03Z` + "`" + `). If this isn't specified, the current date is used. Changing this field forces a new resource to be created.`,
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

Manages a password credential associated with a service principal within Azure Active Directory. See also the [azuread_application_password resource](application_password.html).

`,
			Keywords: []string{
				"service",
				"principals",
				"principal",
				"password",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_principal_id",
					Description: `(Required) The object ID of the service principal for which this password should be created. Changing this field forces a new resource to be created. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name for the password.`,
				},
				resource.Attribute{
					Name:        "end_date",
					Description: `The end date until which the password is valid, formatted as an RFC3339 date string (e.g. ` + "`" + `2018-01-01T01:02:03Z` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `A UUID used to uniquely identify this password credential.`,
				},
				resource.Attribute{
					Name:        "start_date",
					Description: `The start date from which the password is valid, formatted as an RFC3339 date string (e.g. ` + "`" + `2018-01-01T01:02:03Z` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The password for this service principal, which is generated by Azure Active Directory. ## Import This resource does not support importing.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name for the password.`,
				},
				resource.Attribute{
					Name:        "end_date",
					Description: `The end date until which the password is valid, formatted as an RFC3339 date string (e.g. ` + "`" + `2018-01-01T01:02:03Z` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `A UUID used to uniquely identify this password credential.`,
				},
				resource.Attribute{
					Name:        "start_date",
					Description: `The start date from which the password is valid, formatted as an RFC3339 date string (e.g. ` + "`" + `2018-01-01T01:02:03Z` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The password for this service principal, which is generated by Azure Active Directory. ## Import This resource does not support importing.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuread_user",
			Category:         "Users",
			ShortDescription: ``,
			Description: `

Manages a user within Azure Active Directory.

`,
			Keywords: []string{
				"users",
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_enabled",
					Description: `(Optional) Whether or not the account should be enabled.`,
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
					Description: `(Optional) The country/region in which the user is located, e.g. ` + "`" + `US` + "`" + ` or ` + "`" + `UK` + "`" + `.`,
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
					Description: `(Optional) Whether the user is forced to change the password during the next sign-in. Only takes effect when also changing the password. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "given_name",
					Description: `(Optional) The given name (first name) of the user.`,
				},
				resource.Attribute{
					Name:        "job_title",
					Description: `(Optional) The user’s job title.`,
				},
				resource.Attribute{
					Name:        "mail_nickname",
					Description: `(Optional) The mail alias for the user. Defaults to the user name part of the user principal name (UPN).`,
				},
				resource.Attribute{
					Name:        "mobile_phone",
					Description: `(Optional) The primary cellular telephone number for the user.`,
				},
				resource.Attribute{
					Name:        "office_location",
					Description: `(Optional) The office location in the user's place of business.`,
				},
				resource.Attribute{
					Name:        "onpremises_immutable_id",
					Description: `(Optional) The value used to associate an on-premise Active Directory user account with their Azure AD user object. This must be specified if you are using a federated domain for the user's ` + "`" + `user_principal_name` + "`" + ` property when creating a new user account.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The password for the user. The password must satisfy minimum requirements as specified by the password policy. The maximum length is 256 characters. This property is required when creating a new user.`,
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
					Description: `(Optional) The usage location of the user. Required for users that will be assigned licenses due to legal requirement to check for availability of services in countries. The usage location is a two letter country code (ISO standard 3166). Examples include: ` + "`" + `NO` + "`" + `, ` + "`" + `JP` + "`" + `, and ` + "`" + `GB` + "`" + `. Cannot be reset to null once set.`,
				},
				resource.Attribute{
					Name:        "user_principal_name",
					Description: `(Required) The user principal name (UPN) of the user. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
					Name:        "onpremises_sam_account_name",
					Description: `The on-premise SAM account name of the user.`,
				},
				resource.Attribute{
					Name:        "onpremises_user_principal_name",
					Description: `The on-premise user principal name of the user.`,
				},
				resource.Attribute{
					Name:        "user_type",
					Description: `The user type in the directory. Possible values are ` + "`" + `Guest` + "`" + ` or ` + "`" + `Member` + "`" + `. ## Import Users can be imported using their object ID, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_user.my_user 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "mail",
					Description: `The primary email address of the user.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The object ID of the user.`,
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
					Name:        "user_type",
					Description: `The user type in the directory. Possible values are ` + "`" + `Guest` + "`" + ` or ` + "`" + `Member` + "`" + `. ## Import Users can be imported using their object ID, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_user.my_user 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"azuread_application":                   0,
		"azuread_application_certificate":       1,
		"azuread_application_password":          2,
		"azuread_application_pre_authorized":    3,
		"azuread_group":                         4,
		"azuread_group_member":                  5,
		"azuread_service_principal":             6,
		"azuread_service_principal_certificate": 7,
		"azuread_service_principal_password":    8,
		"azuread_user":                          9,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
