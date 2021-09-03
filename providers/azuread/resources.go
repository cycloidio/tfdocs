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
					Description: `(Optional) An ` + "`" + `api` + "`" + ` block as documented below, which configures API related settings for this application.`,
				},
				resource.Attribute{
					Name:        "app_role",
					Description: `(Optional) A collection of ` + "`" + `app_role` + "`" + ` blocks as documented below. For more information see [official documentation on Application Roles](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).`,
				},
				resource.Attribute{
					Name:        "device_only_auth_enabled",
					Description: `(Optional) Specifies whether this application supports device authentication without a user. Defaults to ` + "`" + `false` + "`" + `.`,
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
					Description: `(Optional) A set of user-defined URI(s) that uniquely identify an application within its Azure AD tenant, or within a verified custom domain if the application is multi-tenant.`,
				},
				resource.Attribute{
					Name:        "marketing_url",
					Description: `(Optional) URL of the application's marketing page.`,
				},
				resource.Attribute{
					Name:        "oauth2_post_response_required",
					Description: `(Optional) Specifies whether, as part of OAuth 2.0 token requests, Azure AD allows POST requests, as opposed to GET requests. Defaults to ` + "`" + `false` + "`" + `, which specifies that only GET requests are allowed.`,
				},
				resource.Attribute{
					Name:        "optional_claims",
					Description: `(Optional) An ` + "`" + `optional_claims` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "owners",
					Description: `(Optional) A set of object IDs of principals that will be granted ownership of the application. Supported object types are users or service principals. By default, no owners are assigned. ->`,
				},
				resource.Attribute{
					Name:        "prevent_duplicate_names",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, will return an error if an existing application is found with the same name. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "privacy_statement_url",
					Description: `(Optional) URL of the application's privacy statement.`,
				},
				resource.Attribute{
					Name:        "public_client",
					Description: `(Optional) A ` + "`" + `public_client` + "`" + ` block as documented below, which configures non-web app or non-web API application settings, for example mobile or other public clients such as an installed application running on a desktop device.`,
				},
				resource.Attribute{
					Name:        "required_resource_access",
					Description: `(Optional) A collection of ` + "`" + `required_resource_access` + "`" + ` blocks as documented below.`,
				},
				resource.Attribute{
					Name:        "sign_in_audience",
					Description: `(Optional) The Microsoft account types that are supported for the current application. Must be one of ` + "`" + `AzureADMyOrg` + "`" + `, ` + "`" + `AzureADMultipleOrgs` + "`" + `, ` + "`" + `AzureADandPersonalMicrosoftAccount` + "`" + ` or ` + "`" + `PersonalMicrosoftAccount` + "`" + `. Defaults to ` + "`" + `AzureADMyOrg` + "`" + `. ~>`,
				},
				resource.Attribute{
					Name:        "single_page_application",
					Description: `(Optional) A ` + "`" + `single_page_application` + "`" + ` block as documented below, which configures single-page application (SPA) related settings for this application.`,
				},
				resource.Attribute{
					Name:        "support_url",
					Description: `(Optional) URL of the application's support page.`,
				},
				resource.Attribute{
					Name:        "terms_of_service_url",
					Description: `(Optional) URL of the application's terms of service statement.`,
				},
				resource.Attribute{
					Name:        "web",
					Description: `(Optional) A ` + "`" + `web` + "`" + ` block as documented below, which configures web related settings for this application. ->`,
				},
				resource.Attribute{
					Name:        "known_client_applications",
					Description: `(Optional) A set of application IDs (client IDs), used for bundling consent if you have a solution that contains two parts: a client app and a custom web API app.`,
				},
				resource.Attribute{
					Name:        "mapped_claims_enabled",
					Description: `(Optional) Allows an application to use claims mapping without specifying a custom signing key. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "oauth2_permission_scope",
					Description: `(Optional) One or more ` + "`" + `oauth2_permission_scope` + "`" + ` blocks as documented below, to describe delegated permissions exposed by the web API represented by this application.`,
				},
				resource.Attribute{
					Name:        "requested_access_token_version",
					Description: `(Optional) The access token version expected by this resource. Must be one of ` + "`" + `1` + "`" + ` or ` + "`" + `2` + "`" + `, and must be ` + "`" + `2` + "`" + ` when ` + "`" + `sign_in_audience` + "`" + ` is either ` + "`" + `AzureADandPersonalMicrosoftAccount` + "`" + ` or ` + "`" + `PersonalMicrosoftAccount` + "`" + ` Defaults to ` + "`" + `1` + "`" + `. --- ` + "`" + `oauth2_permission_scope` + "`" + ` blocks support the following:`,
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
					Description: `The source of the claim. If ` + "`" + `source` + "`" + ` is absent, the claim is a predefined optional claim. If ` + "`" + `source` + "`" + ` is ` + "`" + `user` + "`" + `, the value of ` + "`" + `name` + "`" + ` is the extension property from the user object. --- ` + "`" + `public_client` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "redirect_uris",
					Description: `(Optional) A set of URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent. --- ` + "`" + `required_resource_access` + "`" + ` block supports the following:`,
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
					Description: `(Required) Specifies whether the ` + "`" + `id` + "`" + ` property references an app role or an OAuth2 permission scope. Possible values are ` + "`" + `Role` + "`" + ` or ` + "`" + `Scope` + "`" + `. --- ` + "`" + `single_page_application` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "redirect_uris",
					Description: `(Optional) A set of URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent. --- ` + "`" + `web` + "`" + ` block supports the following:`,
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
					Description: `(Optional) A set of URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent. --- ` + "`" + `implicit_grant` + "`" + ` block supports the following:`,
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
					Name:        "app_role_ids",
					Description: `A mapping of app role values to app role IDs, intended to be useful when referencing app roles in other resources in your configuration.`,
				},
				resource.Attribute{
					Name:        "application_id",
					Description: `The Application ID (also called Client ID).`,
				},
				resource.Attribute{
					Name:        "disabled_by_microsoft",
					Description: `Whether Microsoft has disabled the registered application. If the application is disabled, this will be a string indicating the status/reason, e.g. ` + "`" + `DisabledDueToViolationOfServicesAgreement` + "`" + ``,
				},
				resource.Attribute{
					Name:        "logo_url",
					Description: `CDN URL to the application's logo.`,
				},
				resource.Attribute{
					Name:        "oauth2_permission_scope_ids",
					Description: `A mapping of OAuth2.0 permission scope values to scope IDs, intended to be useful when referencing permission scopes in other resources in your configuration.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The application's object ID.`,
				},
				resource.Attribute{
					Name:        "publisher_domain",
					Description: `The verified publisher domain for the application. ## Import Applications can be imported using their object ID, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_application.test 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "app_role_ids",
					Description: `A mapping of app role values to app role IDs, intended to be useful when referencing app roles in other resources in your configuration.`,
				},
				resource.Attribute{
					Name:        "application_id",
					Description: `The Application ID (also called Client ID).`,
				},
				resource.Attribute{
					Name:        "disabled_by_microsoft",
					Description: `Whether Microsoft has disabled the registered application. If the application is disabled, this will be a string indicating the status/reason, e.g. ` + "`" + `DisabledDueToViolationOfServicesAgreement` + "`" + ``,
				},
				resource.Attribute{
					Name:        "logo_url",
					Description: `CDN URL to the application's logo.`,
				},
				resource.Attribute{
					Name:        "oauth2_permission_scope_ids",
					Description: `A mapping of OAuth2.0 permission scope values to scope IDs, intended to be useful when referencing permission scopes in other resources in your configuration.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The application's object ID.`,
				},
				resource.Attribute{
					Name:        "publisher_domain",
					Description: `The verified publisher domain for the application. ## Import Applications can be imported using their object ID, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_application.test 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Optional) A relative duration for which the certificate is valid until, for example ` + "`" + `240h` + "`" + ` (10 days) or ` + "`" + `2400h30m` + "`" + `. Changing this field forces a new resource to be created. ~> One of ` + "`" + `end_date` + "`" + ` or ` + "`" + `end_date_relative` + "`" + ` must be set. The maximum allowed duration is determined by Azure AD.`,
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
					Description: `(Required) The certificate data, which can be PEM encoded, base64 encoded DER or hexadecimal encoded DER. See also the ` + "`" + `encoding` + "`" + ` argument. ## Attributes Reference No additional attributes are exported. ## Import Certificates can be imported using the object ID of the associated application and the key ID of the certificate credential, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_application_certificate.test 00000000-0000-0000-0000-000000000000/certificate/11111111-1111-1111-1111-111111111111 ` + "`" + `` + "`" + `` + "`" + ` -> This ID format is unique to Terraform and is composed of the application's object ID, the string "certificate" and the certificate's key ID in the format ` + "`" + `{ObjectId}/certificate/{CertificateKeyId}` + "`" + `.`,
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
					Description: `(Required) A set of permission scope IDs required by the authorized application. ## Attributes Reference No additional attributes are exported. ## Import Pre-authorized applications can be imported using the object ID of the authorizing application and the application ID of the application being authorized, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_application_pre_authorized.example 00000000-0000-0000-0000-000000000000/preAuthorizedApplication/11111111-1111-1111-1111-111111111111 ` + "`" + `` + "`" + `` + "`" + ` -> This ID format is unique to Terraform and is composed of the authorizing application's object ID, the string "preAuthorizedApplication" and the authorized application's application ID (client ID) in the format ` + "`" + `{ObjectId}/preAuthorizedApplication/{ApplicationId}` + "`" + `.`,
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
					Name:        "assignable_to_role",
					Description: `(Optional) Indicates whether this group can be assigned to an Azure Active Directory role. Can only be ` + "`" + `true` + "`" + ` for security-enabled groups. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "behaviors",
					Description: `(Optional) A set of behaviors for a Microsoft 365 group. Possible values are ` + "`" + `AllowOnlyMembersToPost` + "`" + `, ` + "`" + `HideGroupInOutlook` + "`" + `, ` + "`" + `SubscribeNewGroupMembers` + "`" + ` and ` + "`" + `WelcomeEmailDisabled` + "`" + `. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for more details. Changing this forces a new resource to be created.`,
				},
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
					Description: `(Optional) Whether the group is a mail enabled, with a shared group mailbox. At least one of ` + "`" + `mail_enabled` + "`" + ` or ` + "`" + `security_enabled` + "`" + ` must be specified. Only Microsoft 365 groups can be mail enabled (see the ` + "`" + `types` + "`" + ` property).`,
				},
				resource.Attribute{
					Name:        "mail_nickname",
					Description: `(Optional) The mail alias for the group, unique in the organisation. Required for mail-enabled groups. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `(Optional) A set of members who should be present in this group. Supported object types are Users, Groups or Service Principals. !>`,
				},
				resource.Attribute{
					Name:        "owners",
					Description: `(Optional) A set of object IDs of principals that will be granted ownership of the group. Supported object types are users or service principals. By default, the principal being used to execute Terraform is assigned as the sole owner. Groups cannot be created with no owners or have all their owners removed. ->`,
				},
				resource.Attribute{
					Name:        "prevent_duplicate_names",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, will return an error if an existing group is found with the same name. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "provisioning_options",
					Description: `(Optional) A set of provisioning options for a Microsoft 365 group. The only supported value is ` + "`" + `Team` + "`" + `. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for details. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "security_enabled",
					Description: `(Optional) Whether the group is a security group for controlling access to in-app resources. At least one of ` + "`" + `security_enabled` + "`" + ` or ` + "`" + `mail_enabled` + "`" + ` must be specified. A Microsoft 365 group can be security enabled _and_ mail enabled (see the ` + "`" + `types` + "`" + ` property).`,
				},
				resource.Attribute{
					Name:        "theme",
					Description: `(Optional) The colour theme for a Microsoft 365 group. Possible values are ` + "`" + `Blue` + "`" + `, ` + "`" + `Green` + "`" + `, ` + "`" + `Orange` + "`" + `, ` + "`" + `Pink` + "`" + `, ` + "`" + `Purple` + "`" + `, ` + "`" + `Red` + "`" + ` or ` + "`" + `Teal` + "`" + `. By default, no theme is set.`,
				},
				resource.Attribute{
					Name:        "types",
					Description: `(Optional) A set of group types to configure for the group. The only supported type is ` + "`" + `Unified` + "`" + `, which specifies a Microsoft 365 group. Required when ` + "`" + `mail_enabled` + "`" + ` is true. Changing this forces a new resource to be created. ->`,
				},
				resource.Attribute{
					Name:        "visibility",
					Description: `(Optional) The group join policy and group content visibility. Possible values are ` + "`" + `Private` + "`" + `, ` + "`" + `Public` + "`" + `, or ` + "`" + `Hiddenmembership` + "`" + `. Only Microsoft 365 groups can have ` + "`" + `Hiddenmembership` + "`" + ` visibility and this value must be set when the group is created. By default, security groups will receive ` + "`" + `Private` + "`" + ` visibility and Microsoft 365 groups will receive ` + "`" + `Public` + "`" + ` visibility. ->`,
				},
				resource.Attribute{
					Name:        "mail",
					Description: `The SMTP address for the group.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The object ID of the group.`,
				},
				resource.Attribute{
					Name:        "onpremises_domain_name",
					Description: `The on-premises FQDN, also called dnsDomainName, synchronised from the on-premises directory when Azure AD Connect is used.`,
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
					Name:        "preferred_language",
					Description: `The preferred language for a Microsoft 365 group, in ISO 639-1 notation.`,
				},
				resource.Attribute{
					Name:        "proxy_addresses",
					Description: `List of email addresses for the group that direct to the same group mailbox. ## Import Groups can be imported using their object ID, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_group.my_group 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "mail",
					Description: `The SMTP address for the group.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The object ID of the group.`,
				},
				resource.Attribute{
					Name:        "onpremises_domain_name",
					Description: `The on-premises FQDN, also called dnsDomainName, synchronised from the on-premises directory when Azure AD Connect is used.`,
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
					Name:        "preferred_language",
					Description: `The preferred language for a Microsoft 365 group, in ISO 639-1 notation.`,
				},
				resource.Attribute{
					Name:        "proxy_addresses",
					Description: `List of email addresses for the group that direct to the same group mailbox. ## Import Groups can be imported using their object ID, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_group.my_group 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
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

~> **Warning** Do not use this resource at the same time as the ` + "`" + `members` + "`" + ` property of the ` + "`" + `azuread_group` + "`" + ` resource for the same group. Doing so will cause a conflict and group members will be removed.

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
					Name:        "account_enabled",
					Description: `(Optional) Whether or not the service principal account is enabled. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "alternative_names",
					Description: `(Optional) A set of alternative names, used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities.`,
				},
				resource.Attribute{
					Name:        "app_role_assignment_required",
					Description: `(Optional) Whether this service principal requires an app role assignment to a user or group before Azure AD will issue a user or access token to the application. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "application_id",
					Description: `(Required) The application ID (client ID) of the application for which to create a service principal.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the service principal provided for internal end-users.`,
				},
				resource.Attribute{
					Name:        "login_url",
					Description: `(Optional) The URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps. When blank, Azure AD performs IdP-initiated sign-on for applications configured with SAML-based single sign-on.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) A free text field to capture information about the service principal, typically used for operational purposes.`,
				},
				resource.Attribute{
					Name:        "notification_email_addresses",
					Description: `(Optional) A set of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications.`,
				},
				resource.Attribute{
					Name:        "owners",
					Description: `(Optional) A set of object IDs of principals that will be granted ownership of the service principal. Supported object types are users or service principals. By default, no owners are assigned. ->`,
				},
				resource.Attribute{
					Name:        "preferred_single_sign_on_mode",
					Description: `(Optional) The single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps. Supported values are ` + "`" + `oidc` + "`" + `, ` + "`" + `password` + "`" + `, ` + "`" + `saml` + "`" + ` or ` + "`" + `notSupported` + "`" + `. Omit this property or specify a blank string to unset.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A set of tags to apply to the service principal.`,
				},
				resource.Attribute{
					Name:        "use_existing",
					Description: `(Optional) When true, any existing service principal linked to the same application will be automatically imported. When false, an import error will be raised for any pre-existing service principal. ->`,
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
					Name:        "display_name",
					Description: `The display name of the application associated with this service principal.`,
				},
				resource.Attribute{
					Name:        "homepage_url",
					Description: `Home page or landing page of the associated application.`,
				},
				resource.Attribute{
					Name:        "logout_url",
					Description: `The URL that will be used by Microsoft's authorization service to logout an user using OpenId Connect front-channel, back-channel or SAML logout protocols, taken from the associated application.`,
				},
				resource.Attribute{
					Name:        "oauth2_permission_scope_ids",
					Description: `A mapping of OAuth2.0 permission scope values to scope IDs, as exposed by the associated application, intended to be useful when referencing permission scopes in other resources in your configuration.`,
				},
				resource.Attribute{
					Name:        "oauth2_permission_scopes",
					Description: `A list of OAuth 2.0 delegated permission scopes exposed by the associated application, as documented below.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The object ID of the service principal.`,
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
					Name:        "service_principal_names",
					Description: `A list of identifier URI(s), copied over from the associated application.`,
				},
				resource.Attribute{
					Name:        "sign_in_audience",
					Description: `The Microsoft account types that are supported for the associated application. Possible values include ` + "`" + `AzureADMyOrg` + "`" + `, ` + "`" + `AzureADMultipleOrgs` + "`" + `, ` + "`" + `AzureADandPersonalMicrosoftAccount` + "`" + ` or ` + "`" + `PersonalMicrosoftAccount` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Identifies whether the service principal represents an application or a managed identity. Possible values include ` + "`" + `Application` + "`" + ` or ` + "`" + `ManagedIdentity` + "`" + `. --- ` + "`" + `app_roles` + "`" + ` is a list of objects with the following attributes:`,
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
					Name:        "display_name",
					Description: `The display name of the application associated with this service principal.`,
				},
				resource.Attribute{
					Name:        "homepage_url",
					Description: `Home page or landing page of the associated application.`,
				},
				resource.Attribute{
					Name:        "logout_url",
					Description: `The URL that will be used by Microsoft's authorization service to logout an user using OpenId Connect front-channel, back-channel or SAML logout protocols, taken from the associated application.`,
				},
				resource.Attribute{
					Name:        "oauth2_permission_scope_ids",
					Description: `A mapping of OAuth2.0 permission scope values to scope IDs, as exposed by the associated application, intended to be useful when referencing permission scopes in other resources in your configuration.`,
				},
				resource.Attribute{
					Name:        "oauth2_permission_scopes",
					Description: `A list of OAuth 2.0 delegated permission scopes exposed by the associated application, as documented below.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The object ID of the service principal.`,
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
					Name:        "service_principal_names",
					Description: `A list of identifier URI(s), copied over from the associated application.`,
				},
				resource.Attribute{
					Name:        "sign_in_audience",
					Description: `The Microsoft account types that are supported for the associated application. Possible values include ` + "`" + `AzureADMyOrg` + "`" + `, ` + "`" + `AzureADMultipleOrgs` + "`" + `, ` + "`" + `AzureADandPersonalMicrosoftAccount` + "`" + ` or ` + "`" + `PersonalMicrosoftAccount` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Identifies whether the service principal represents an application or a managed identity. Possible values include ` + "`" + `Application` + "`" + ` or ` + "`" + `ManagedIdentity` + "`" + `. --- ` + "`" + `app_roles` + "`" + ` is a list of objects with the following attributes:`,
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
					Description: `(Optional) A relative duration for which the certificate is valid until, for example ` + "`" + `240h` + "`" + ` (10 days) or ` + "`" + `2400h30m` + "`" + `. Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h". Changing this field forces a new resource to be created. ~> One of ` + "`" + `end_date` + "`" + ` or ` + "`" + `end_date_relative` + "`" + ` must be set. The maximum duration is determined by Azure AD.`,
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
					Description: `(Required) The certificate data, which can be PEM encoded, base64 encoded DER or hexadecimal encoded DER. See also the ` + "`" + `encoding` + "`" + ` argument. ## Attributes Reference No additional attributes are exported. ## Import Certificates can be imported using the object ID of the associated service principal and the key ID of the certificate credential, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_service_principal_certificate.test 00000000-0000-0000-0000-000000000000/certificate/11111111-1111-1111-1111-111111111111 ` + "`" + `` + "`" + `` + "`" + ` -> This ID format is unique to Terraform and is composed of the service principal's object ID, the string "certificate" and the certificate's key ID in the format ` + "`" + `{ServicePrincipalObjectId}/certificate/{CertificateKeyId}` + "`" + `.`,
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
					Name:        "age_group",
					Description: `(Optional) The age group of the user. Supported values are ` + "`" + `Adult` + "`" + `, ` + "`" + `NotAdult` + "`" + ` and ` + "`" + `Minor` + "`" + `. Omit this property or specify a blank string to unset.`,
				},
				resource.Attribute{
					Name:        "business_phones",
					Description: `(Optional) A list of telephone numbers for the user. Only one number can be set for this property. Read-only for users synced with Azure AD Connect.`,
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
					Name:        "consent_provided_for_minor",
					Description: `(Optional) Whether consent has been obtained for minors. Supported values are ` + "`" + `Granted` + "`" + `, ` + "`" + `Denied` + "`" + ` and ` + "`" + `NotRequired` + "`" + `. Omit this property or specify a blank string to unset.`,
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
					Name:        "employee_id",
					Description: `(Optional) The employee identifier assigned to the user by the organisation.`,
				},
				resource.Attribute{
					Name:        "fax_number",
					Description: `(Optional) The fax number of the user.`,
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
					Name:        "mail",
					Description: `(Optional) The SMTP address for the user. This property cannot be unset once specified.`,
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
					Name:        "other_mails",
					Description: `(Optional) A list of additional email addresses for the user.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The password for the user. The password must satisfy minimum requirements as specified by the password policy. The maximum length is 256 characters. This property is required when creating a new user. ->`,
				},
				resource.Attribute{
					Name:        "postal_code",
					Description: `(Optional) The postal code for the user's postal address. The postal code is specific to the user's country/region. In the United States of America, this attribute contains the ZIP code.`,
				},
				resource.Attribute{
					Name:        "preferred_language",
					Description: `(Optional) The user's preferred language, in ISO 639-1 notation.`,
				},
				resource.Attribute{
					Name:        "show_in_address_list",
					Description: `(Optional) Whether or not the Outlook global address list should include this user. Defaults to ` + "`" + `true` + "`" + `.`,
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
					Name:        "creation_type",
					Description: `Indicates whether the user account was created as a regular school or work account (` + "`" + `null` + "`" + `), an external account (` + "`" + `Invitation` + "`" + `), a local account for an Azure Active Directory B2C tenant (` + "`" + `LocalAccount` + "`" + `) or self-service sign-up using email verification (` + "`" + `EmailVerified` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "external_user_state",
					Description: `For an external user invited to the tenant, this property represents the invited user's invitation status. Possible values are ` + "`" + `PendingAcceptance` + "`" + ` or ` + "`" + `Accepted` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "im_addresses",
					Description: `A list of instant message voice over IP (VOIP) session initiation protocol (SIP) addresses for the user.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The object ID of the user.`,
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
					Name:        "proxy_addresses",
					Description: `List of email addresses for the user that direct to the same mailbox.`,
				},
				resource.Attribute{
					Name:        "user_type",
					Description: `The user type in the directory. Possible values are ` + "`" + `Guest` + "`" + ` or ` + "`" + `Member` + "`" + `. ## Import Users can be imported using their object ID, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azuread_user.my_user 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "creation_type",
					Description: `Indicates whether the user account was created as a regular school or work account (` + "`" + `null` + "`" + `), an external account (` + "`" + `Invitation` + "`" + `), a local account for an Azure Active Directory B2C tenant (` + "`" + `LocalAccount` + "`" + `) or self-service sign-up using email verification (` + "`" + `EmailVerified` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "external_user_state",
					Description: `For an external user invited to the tenant, this property represents the invited user's invitation status. Possible values are ` + "`" + `PendingAcceptance` + "`" + ` or ` + "`" + `Accepted` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "im_addresses",
					Description: `A list of instant message voice over IP (VOIP) session initiation protocol (SIP) addresses for the user.`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `The object ID of the user.`,
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
					Name:        "proxy_addresses",
					Description: `List of email addresses for the user that direct to the same mailbox.`,
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
