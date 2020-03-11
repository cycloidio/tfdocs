package okta

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "okta_app_auto_login",
			Category:         "Resources",
			ShortDescription: `Creates an Auto Login Okta Application.`,
			Description:      ``,
			Keywords: []string{
				"app",
				"auto",
				"login",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "label",
					Description: `(Required) The Application's display name.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The status of the application, by default it is ` + "`" + `"ACTIVE"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "preconfigured_app",
					Description: `(Optional) Tells Okta to use an existing application in their application catalog, as opposed to a custom application. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name assigned to the application by Okta.`,
				},
				resource.Attribute{
					Name:        "sign_on_mode",
					Description: `Sign on mode of application. ## Import Okta Auto Login App can be imported via the Okta ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_app_auto_login.example <app id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name assigned to the application by Okta.`,
				},
				resource.Attribute{
					Name:        "sign_on_mode",
					Description: `Sign on mode of application. ## Import Okta Auto Login App can be imported via the Okta ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_app_auto_login.example <app id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_app_basic_auth",
			Category:         "Resources",
			ShortDescription: `Creates a Basic Auth Application.`,
			Description:      ``,
			Keywords: []string{
				"app",
				"basic",
				"auth",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "label",
					Description: `(Required) The Application's display name.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Optional) The URL of the sign-in page for this app.`,
				},
				resource.Attribute{
					Name:        "auth_url",
					Description: `(Optional) The URL of the authenticating site for this app. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Application.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The Application's display name.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `The URL of the sign-in page for basic auth app.`,
				},
				resource.Attribute{
					Name:        "auth_url",
					Description: `The URL of the authenticating site for basic auth app. ## Import A Basic Auth App can be imported via the Okta ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_app_basic_auth.example <app id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Application.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The Application's display name.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `The URL of the sign-in page for basic auth app.`,
				},
				resource.Attribute{
					Name:        "auth_url",
					Description: `The URL of the authenticating site for basic auth app. ## Import A Basic Auth App can be imported via the Okta ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_app_basic_auth.example <app id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_app_bookmark",
			Category:         "Resources",
			ShortDescription: `Creates a Bookmark Application.`,
			Description:      ``,
			Keywords: []string{
				"app",
				"bookmark",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "label",
					Description: `(Required) The Application's display name.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Optional) The URL of the bookmark.`,
				},
				resource.Attribute{
					Name:        "request_integration",
					Description: `(Optional) Would you like Okta to add an integration for this app? ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Application.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The Application's display name.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `The URL of the bookmark. ## Import A Bookmark App can be imported via the Okta ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_app_bookmark.example <app id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Application.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The Application's display name.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `The URL of the bookmark. ## Import A Bookmark App can be imported via the Okta ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_app_bookmark.example <app id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_app_group_assignment",
			Category:         "Resources",
			ShortDescription: `Assigns a group to an application.`,
			Description:      ``,
			Keywords: []string{
				"app",
				"group",
				"assignment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "app_id",
					Description: `(Required) The ID of the application to assign a group to.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required) The ID of the group to assign the app to.`,
				},
				resource.Attribute{
					Name:        "profile",
					Description: `(Optional) JSON document containing [application profile](https://developer.okta.com/docs/reference/api/apps/#profile-object) ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the group assignment. ## Import An application group assignment can be imported via assignment ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_app_group_assignment.example <id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the group assignment. ## Import An application group assignment can be imported via assignment ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_app_group_assignment.example <id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_app_oauth",
			Category:         "Resources",
			ShortDescription: `Creates an OIDC Application.`,
			Description:      ``,
			Keywords: []string{
				"app",
				"oauth",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "label",
					Description: `(Required) The Application's display name.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The status of the application, by default it is ` + "`" + `"ACTIVE"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of OAuth application.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `(Optional) The users assigned to the application. It is recommended not to use this and instead use ` + "`" + `okta_app_user` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `(Optional) The groups assigned to the application. It is recommended not to use this and instead use ` + "`" + `okta_app_group_assignment` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `(Optional) OAuth client ID. If set during creation, app is created with this id.`,
				},
				resource.Attribute{
					Name:        "omit_secret",
					Description: `(Optional) This tells the provider not to persist the application's secret to state. If this is ever changes from true => false your app will be recreated.`,
				},
				resource.Attribute{
					Name:        "client_basic_secret",
					Description: `(Optional) OAuth client secret key, this can be set when token_endpoint_auth_method is client_secret_basic.`,
				},
				resource.Attribute{
					Name:        "token_endpoint_auth_method",
					Description: `(Optional) Requested authentication method for the token endpoint. It can be set to ` + "`" + `"none"` + "`" + `, ` + "`" + `"client_secret_post"` + "`" + `, ` + "`" + `"client_secret_basic"` + "`" + `, ` + "`" + `"client_secret_jwt"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "auto_key_rotation",
					Description: `(Optional) Requested key rotation mode.`,
				},
				resource.Attribute{
					Name:        "client_uri",
					Description: `(Optional) URI to a web page providing information about the client.`,
				},
				resource.Attribute{
					Name:        "logo_uri",
					Description: `(Optional) URI that references a logo for the client.`,
				},
				resource.Attribute{
					Name:        "login_uri",
					Description: `(Optional) URI that initiates login.`,
				},
				resource.Attribute{
					Name:        "redirect_uris",
					Description: `(Optional) List of URIs for use in the redirect-based flow. This is required for all application types except service.`,
				},
				resource.Attribute{
					Name:        "post_logout_redirect_uris",
					Description: `(Optional) List of URIs for redirection after logout.`,
				},
				resource.Attribute{
					Name:        "response_types",
					Description: `(Optional) List of OAuth 2.0 response type strings.`,
				},
				resource.Attribute{
					Name:        "grant_types",
					Description: `(Optional) List of OAuth 2.0 grant types. Conditional validation params found here https://developer.okta.com/docs/api/resources/apps#credentials-settings-details. Defaults to minimum requirements per app type.`,
				},
				resource.Attribute{
					Name:        "tos_uri",
					Description: `(Optional) URI to web page providing client tos (terms of service).`,
				},
				resource.Attribute{
					Name:        "policy_uri",
					Description: `(Optional) URI to web page providing client policy document.`,
				},
				resource.Attribute{
					Name:        "consent_method",
					Description: `(Optional) Indicates whether user consent is required or implicit. Valid values: REQUIRED, TRUSTED. Default value is TRUSTED.`,
				},
				resource.Attribute{
					Name:        "issuer_mode",
					Description: `(Optional) Indicates whether the Okta Authorization Server uses the original Okta org domain URL or a custom domain URL as the issuer of ID token for this client.`,
				},
				resource.Attribute{
					Name:        "auto_submit_toolbar",
					Description: `(Optional) Display auto submit toolbar.`,
				},
				resource.Attribute{
					Name:        "hide_ios",
					Description: `(Optional) Do not display application icon on mobile app.`,
				},
				resource.Attribute{
					Name:        "hide_web",
					Description: `(Optional) Do not display application icon to users.`,
				},
				resource.Attribute{
					Name:        "profile",
					Description: `(Optional) Custom JSON that represents an OAuth application's profile. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name assigned to the application by Okta.`,
				},
				resource.Attribute{
					Name:        "sign_on_mode",
					Description: `Sign on mode of application.`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `The client ID of the application.`,
				},
				resource.Attribute{
					Name:        "client_secret",
					Description: `The client secret of the application. ## Import An OIDC Application can be imported via the Okta ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_app_oauth.example <app id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name assigned to the application by Okta.`,
				},
				resource.Attribute{
					Name:        "sign_on_mode",
					Description: `Sign on mode of application.`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `The client ID of the application.`,
				},
				resource.Attribute{
					Name:        "client_secret",
					Description: `The client secret of the application. ## Import An OIDC Application can be imported via the Okta ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_app_oauth.example <app id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_app_saml",
			Category:         "Resources",
			ShortDescription: `Creates an SAML Application.`,
			Description:      ``,
			Keywords: []string{
				"app",
				"saml",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "label",
					Description: `(Required) label of application.`,
				},
				resource.Attribute{
					Name:        "preconfigured_app",
					Description: `(Optional) name of application from the Okta Integration Network, if not included a custom app will be created.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) description of application.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) status of application.`,
				},
				resource.Attribute{
					Name:        "auto_submit_toolbar",
					Description: `(Optional) Display auto submit toolbar.`,
				},
				resource.Attribute{
					Name:        "hide_ios",
					Description: `(Optional) Do not display application icon on mobile app.`,
				},
				resource.Attribute{
					Name:        "hide_web",
					Description: `(Optional) Do not display application icon to users`,
				},
				resource.Attribute{
					Name:        "default_relay_state",
					Description: `(Optional) Identifies a specific application resource in an IDP initiated SSO scenario.`,
				},
				resource.Attribute{
					Name:        "sso_url",
					Description: `(Optional) Single Sign on Url.`,
				},
				resource.Attribute{
					Name:        "recipient",
					Description: `(Optional) The location where the app may present the SAML assertion.`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `(Optional) Identifies the location where the SAML response is intended to be sent inside of the SAML assertion.`,
				},
				resource.Attribute{
					Name:        "audience",
					Description: `(Optional) Audience restriction.`,
				},
				resource.Attribute{
					Name:        "idp_issuer",
					Description: `(Optional) SAML issuer ID.`,
				},
				resource.Attribute{
					Name:        "sp_issuer",
					Description: `(Optional) SAML service provider issuer.`,
				},
				resource.Attribute{
					Name:        "subject_name_id_template",
					Description: `(Optional) Template for app user's username when a user is assigned to the app.`,
				},
				resource.Attribute{
					Name:        "subject_name_id_format",
					Description: `(Optional) Identifies the SAML processing rules.`,
				},
				resource.Attribute{
					Name:        "response_signed",
					Description: `(Optional) Determines whether the SAML auth response message is digitally signed.`,
				},
				resource.Attribute{
					Name:        "request_compressed",
					Description: `(Optional) Denotes whether the request is compressed or not.`,
				},
				resource.Attribute{
					Name:        "assertion_signed",
					Description: `(Optional) Determines whether the SAML assertion is digitally signed.`,
				},
				resource.Attribute{
					Name:        "signature_algorithm",
					Description: `(Optional) Signature algorithm used ot digitally sign the assertion and response.`,
				},
				resource.Attribute{
					Name:        "digest_algorithm",
					Description: `(Optional) Determines the digest algorithm used to digitally sign the SAML assertion and response.`,
				},
				resource.Attribute{
					Name:        "honor_force_authn",
					Description: `(Optional) Prompt user to re-authenticate if SP asks for it.`,
				},
				resource.Attribute{
					Name:        "authn_context_class_ref",
					Description: `(Optional) Identifies the SAML authentication context class for the assertion’s authentication statement.`,
				},
				resource.Attribute{
					Name:        "accessibility_self_service",
					Description: `(Optional) Enable self service.`,
				},
				resource.Attribute{
					Name:        "accessibility_error_redirect_url",
					Description: `(Optional) Custom error page URL.`,
				},
				resource.Attribute{
					Name:        "accessibility_login_redirect_url",
					Description: `(Optional) Custom login page URL.`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `(Optional) features enabled.`,
				},
				resource.Attribute{
					Name:        "user_name_template",
					Description: `(Optional) Username template.`,
				},
				resource.Attribute{
					Name:        "user_name_template_suffix",
					Description: `(Optional) Username template suffix.`,
				},
				resource.Attribute{
					Name:        "user_name_template_type",
					Description: `(Optional) Username template type.`,
				},
				resource.Attribute{
					Name:        "app_settings_json",
					Description: `(Optional) Application settings in JSON format.`,
				},
				resource.Attribute{
					Name:        "attribute_statements",
					Description: `(Optional) List of SAML Attribute statements.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the attribute statement.`,
				},
				resource.Attribute{
					Name:        "filter_type",
					Description: `(Optional) Type of group attribute filter.`,
				},
				resource.Attribute{
					Name:        "filter_value",
					Description: `(Optional) Filter value to use.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The attribute namespace. It can be set to ` + "`" + `"urn:oasis:names:tc:SAML:2.0:attrname-format:unspecified"` + "`" + `, ` + "`" + `"urn:oasis:names:tc:SAML:2.0:attrname-format:uri"` + "`" + `, or ` + "`" + `"urn:oasis:names:tc:SAML:2.0:attrname-format:basic"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type of attribute statement value. Can be ` + "`" + `"EXPRESSION"` + "`" + ` or ` + "`" + `"GROUP"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Optional) Array of values to use. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `id of application.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name assigned to the application by Okta.`,
				},
				resource.Attribute{
					Name:        "sign_on_mode",
					Description: `Sign on mode of application.`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `Certificate key ID.`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `The raw signing certificate.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `The raw SAML metadata in XML.`,
				},
				resource.Attribute{
					Name:        "http_post_binding",
					Description: `` + "`" + `urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Post` + "`" + ` location from the SAML metadata.`,
				},
				resource.Attribute{
					Name:        "http_redirect_binding",
					Description: `` + "`" + `urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect` + "`" + ` location from the SAML metadata.`,
				},
				resource.Attribute{
					Name:        "entity_key",
					Description: `Entity ID, the ID portion of the ` + "`" + `entity_url` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "entity_url",
					Description: `Entity URL for instance http://www.okta.com/exk1fcia6d6EMsf331d8. ## Import A SAML App can be imported via the Okta ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_app_saml.example <app id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `id of application.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name assigned to the application by Okta.`,
				},
				resource.Attribute{
					Name:        "sign_on_mode",
					Description: `Sign on mode of application.`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `Certificate key ID.`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `The raw signing certificate.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `The raw SAML metadata in XML.`,
				},
				resource.Attribute{
					Name:        "http_post_binding",
					Description: `` + "`" + `urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Post` + "`" + ` location from the SAML metadata.`,
				},
				resource.Attribute{
					Name:        "http_redirect_binding",
					Description: `` + "`" + `urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect` + "`" + ` location from the SAML metadata.`,
				},
				resource.Attribute{
					Name:        "entity_key",
					Description: `Entity ID, the ID portion of the ` + "`" + `entity_url` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "entity_url",
					Description: `Entity URL for instance http://www.okta.com/exk1fcia6d6EMsf331d8. ## Import A SAML App can be imported via the Okta ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_app_saml.example <app id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_app_secure_password_store",
			Category:         "Resources",
			ShortDescription: `Creates a Secure Password Store Application.`,
			Description:      ``,
			Keywords: []string{
				"app",
				"secure",
				"password",
				"store",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "label",
					Description: `(Required) The display name of the Application.`,
				},
				resource.Attribute{
					Name:        "password_field",
					Description: `(Required) Login password field.`,
				},
				resource.Attribute{
					Name:        "username_field",
					Description: `(Required) Login username field.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Login URL.`,
				},
				resource.Attribute{
					Name:        "optional_field1",
					Description: `(Optional) Name of optional param in the login form.`,
				},
				resource.Attribute{
					Name:        "optional_field1_value",
					Description: `(Optional) Name of optional value in the login form.`,
				},
				resource.Attribute{
					Name:        "optional_field2",
					Description: `(Optional) Name of optional param in the login form.`,
				},
				resource.Attribute{
					Name:        "optional_field2_value",
					Description: `(Optional) Name of optional value in the login form.`,
				},
				resource.Attribute{
					Name:        "optional_field3",
					Description: `(Optional) Name of optional param in the login form.`,
				},
				resource.Attribute{
					Name:        "optional_field3_value",
					Description: `(Optional) Name of optional value in the login form.`,
				},
				resource.Attribute{
					Name:        "credentials_scheme",
					Description: `(Optional) Application credentials scheme. Can be set to ` + "`" + `"EDIT_USERNAME_AND_PASSWORD"` + "`" + `, ` + "`" + `"ADMIN_SETS_CREDENTIALS"` + "`" + `, ` + "`" + `"EDIT_PASSWORD_ONLY"` + "`" + `, ` + "`" + `"EXTERNAL_PASSWORD_SYNC"` + "`" + `, or ` + "`" + `"SHARED_USERNAME_AND_PASSWORD"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "reveal_password",
					Description: `(Optional) Allow user to reveal password.`,
				},
				resource.Attribute{
					Name:        "shared_username",
					Description: `(Optional) Shared username, required for certain schemes.`,
				},
				resource.Attribute{
					Name:        "shared_password",
					Description: `(Optional) Shared password, required for certain schemes.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `(Optional) The users assigned to the application. See ` + "`" + `okta_app_user` + "`" + ` for a more flexible approach.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `(Optional) Groups associated with the application. See ` + "`" + `okta_app_group_assignment` + "`" + ` for a more flexible approach.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Status of application. By default it is ` + "`" + `"ACTIVE"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "accessibility_self_service",
					Description: `(Optional) Enable self service. By default it is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "accessibility_error_redirect_url",
					Description: `(Optional) Custom error page URL.`,
				},
				resource.Attribute{
					Name:        "auto_submit_toolbar",
					Description: `(Optional) Display auto submit toolbar.`,
				},
				resource.Attribute{
					Name:        "hide_ios",
					Description: `(Optional) Do not display application icon on mobile app.`,
				},
				resource.Attribute{
					Name:        "hide_web",
					Description: `(Optional) Do not display application icon to users. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name assigned to the application by Okta.`,
				},
				resource.Attribute{
					Name:        "sign_on_mode",
					Description: `Sign on mode of application.`,
				},
				resource.Attribute{
					Name:        "user_name_template",
					Description: `The default username assigned to each user.`,
				},
				resource.Attribute{
					Name:        "user_name_template_type",
					Description: `The Username template type. ## Import Secure Password Store Application can be imported via the Okta ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_app_secure_password_store.example <app id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name assigned to the application by Okta.`,
				},
				resource.Attribute{
					Name:        "sign_on_mode",
					Description: `Sign on mode of application.`,
				},
				resource.Attribute{
					Name:        "user_name_template",
					Description: `The default username assigned to each user.`,
				},
				resource.Attribute{
					Name:        "user_name_template_type",
					Description: `The Username template type. ## Import Secure Password Store Application can be imported via the Okta ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_app_secure_password_store.example <app id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_app_swa",
			Category:         "Resources",
			ShortDescription: `Creates an SWA Application.`,
			Description:      ``,
			Keywords: []string{
				"app",
				"swa",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "label",
					Description: `(Required) The display name of the Application.`,
				},
				resource.Attribute{
					Name:        "button_field",
					Description: `(Required) Login button field.`,
				},
				resource.Attribute{
					Name:        "preconfigured_app",
					Description: `(Optional) name of application from the Okta Integration Network, if not included a custom app will be created.`,
				},
				resource.Attribute{
					Name:        "password_field",
					Description: `(Optional) Login password field.`,
				},
				resource.Attribute{
					Name:        "username_field",
					Description: `(Optional) Login username field.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Optional) Login URL.`,
				},
				resource.Attribute{
					Name:        "url_regex",
					Description: `(Optional) A regex that further restricts URL to the specified regex.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `(Optional) The users assigned to the application. See ` + "`" + `okta_app_user` + "`" + ` for a more flexible approach.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `(Optional) Groups associated with the application. See ` + "`" + `okta_app_group_assignment` + "`" + ` for a more flexible approach.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Status of application. By default it is ` + "`" + `"ACTIVE"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "accessibility_self_service",
					Description: `(Optional) Enable self service. By default it is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "accessibility_error_redirect_url",
					Description: `(Optional) Custom error page URL.`,
				},
				resource.Attribute{
					Name:        "auto_submit_toolbar",
					Description: `(Optional) Display auto submit toolbar.`,
				},
				resource.Attribute{
					Name:        "hide_ios",
					Description: `(Optional) Do not display application icon on mobile app.`,
				},
				resource.Attribute{
					Name:        "hide_web",
					Description: `(Optional) Do not display application icon to users. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name assigned to the application by Okta.`,
				},
				resource.Attribute{
					Name:        "sign_on_mode",
					Description: `Sign on mode of application.`,
				},
				resource.Attribute{
					Name:        "user_name_template",
					Description: `The default username assigned to each user.`,
				},
				resource.Attribute{
					Name:        "user_name_template_type",
					Description: `The Username template type. ## Import Okta SWA App can be imported via the Okta ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_app_swa.example <app id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name assigned to the application by Okta.`,
				},
				resource.Attribute{
					Name:        "sign_on_mode",
					Description: `Sign on mode of application.`,
				},
				resource.Attribute{
					Name:        "user_name_template",
					Description: `The default username assigned to each user.`,
				},
				resource.Attribute{
					Name:        "user_name_template_type",
					Description: `The Username template type. ## Import Okta SWA App can be imported via the Okta ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_app_swa.example <app id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_app_three_field",
			Category:         "Resources",
			ShortDescription: `Creates an Three Field Application.`,
			Description:      ``,
			Keywords: []string{
				"app",
				"three",
				"field",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "label",
					Description: `(Required) The display name of the Application.`,
				},
				resource.Attribute{
					Name:        "button_selector",
					Description: `(Required) Login button field CSS selector.`,
				},
				resource.Attribute{
					Name:        "password_selector",
					Description: `(Required) Login password field CSS selector.`,
				},
				resource.Attribute{
					Name:        "username_selector",
					Description: `(Required) Login username field CSS selector.`,
				},
				resource.Attribute{
					Name:        "extra_field_selector",
					Description: `(Required) Extra field CSS selector.`,
				},
				resource.Attribute{
					Name:        "extra_field_value",
					Description: `(Required) Value for extra form field.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Login URL.`,
				},
				resource.Attribute{
					Name:        "url_regex",
					Description: `(Optional) A regex that further restricts URL to the specified regex.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `(Optional) The users assigned to the application. See ` + "`" + `okta_app_user` + "`" + ` for a more flexible approach.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `(Optional) Groups associated with the application. See ` + "`" + `okta_app_group_assignment` + "`" + ` for a more flexible approach.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Status of application. By default it is ` + "`" + `"ACTIVE"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "accessibility_self_service",
					Description: `(Optional) Enable self service. By default it is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "accessibility_error_redirect_url",
					Description: `(Optional) Custom error page URL.`,
				},
				resource.Attribute{
					Name:        "auto_submit_toolbar",
					Description: `(Optional) Display auto submit toolbar.`,
				},
				resource.Attribute{
					Name:        "hide_ios",
					Description: `(Optional) Do not display application icon on mobile app.`,
				},
				resource.Attribute{
					Name:        "hide_web",
					Description: `(Optional) Do not display application icon to users. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name assigned to the application by Okta.`,
				},
				resource.Attribute{
					Name:        "sign_on_mode",
					Description: `Sign on mode of application.`,
				},
				resource.Attribute{
					Name:        "user_name_template",
					Description: `The default username assigned to each user.`,
				},
				resource.Attribute{
					Name:        "user_name_template_type",
					Description: `The Username template type. ## Import A Three Field App can be imported via the Okta ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_app_three_field.example <app id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name assigned to the application by Okta.`,
				},
				resource.Attribute{
					Name:        "sign_on_mode",
					Description: `Sign on mode of application.`,
				},
				resource.Attribute{
					Name:        "user_name_template",
					Description: `The default username assigned to each user.`,
				},
				resource.Attribute{
					Name:        "user_name_template_type",
					Description: `The Username template type. ## Import A Three Field App can be imported via the Okta ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_app_three_field.example <app id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_app_user",
			Category:         "Resources",
			ShortDescription: `Creates an Application User.`,
			Description:      ``,
			Keywords: []string{
				"app",
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "app_id",
					Description: `(Required) App to associate user with.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `(Required) User to associate the application with.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The username to use for the app user.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The password to use.`,
				},
				resource.Attribute{
					Name:        "profile",
					Description: `(Optional) The JSON profile of the App User. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the app user. ## Import An Application User can be imported via the Okta ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_app_user.example <app id>/<user id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the app user. ## Import An Application User can be imported via the Okta ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_app_user.example <app id>/<user id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_app_user_base_schema",
			Category:         "Resources",
			ShortDescription: `Manages an Application User Base Schema property.`,
			Description:      ``,
			Keywords: []string{
				"app",
				"user",
				"base",
				"schema",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "app_id",
					Description: `(Required) The Application's ID the user schema property should be assigned to.`,
				},
				resource.Attribute{
					Name:        "index",
					Description: `(Required) The property name.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `(Required) The property display name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of the schema property. It can be ` + "`" + `"string"` + "`" + `, ` + "`" + `"boolean"` + "`" + `, ` + "`" + `"number"` + "`" + `, ` + "`" + `"integer"` + "`" + `, ` + "`" + `"array"` + "`" + `, or ` + "`" + `"object"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "required",
					Description: `(Optional) Whether the property is required for this application's users.`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `(Optional) Access control permissions for the property. It can be set to ` + "`" + `"READ_WRITE"` + "`" + `, ` + "`" + `"READ_ONLY"` + "`" + `, ` + "`" + `"HIDE"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "master",
					Description: `(Optional) Master priority for the user schema property. It can be set to ` + "`" + `"PROFILE_MASTER"` + "`" + ` or ` + "`" + `"OKTA"` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "app_id",
					Description: `ID of the application the user property is associated with.`,
				},
				resource.Attribute{
					Name:        "index",
					Description: `ID of the user schema property. ## Import App user base schema property can be imported via the property index and app id. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_app_user_base_schema.example <app id>/<property name> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "app_id",
					Description: `ID of the application the user property is associated with.`,
				},
				resource.Attribute{
					Name:        "index",
					Description: `ID of the user schema property. ## Import App user base schema property can be imported via the property index and app id. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_app_user_base_schema.example <app id>/<property name> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_app_user_schema",
			Category:         "Resources",
			ShortDescription: `Creates an Application User Schema property.`,
			Description:      ``,
			Keywords: []string{
				"app",
				"user",
				"schema",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "app_id",
					Description: `(Required) The Application's ID the user custom schema property should be assigned to.`,
				},
				resource.Attribute{
					Name:        "index",
					Description: `(Required) The property name.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `(Required) The display name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of the schema property. It can be ` + "`" + `"string"` + "`" + `, ` + "`" + `"boolean"` + "`" + `, ` + "`" + `"number"` + "`" + `, ` + "`" + `"integer"` + "`" + `, ` + "`" + `"array"` + "`" + `, or ` + "`" + `"object"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enum",
					Description: `(Optional) Array of values a primitive property can be set to. See ` + "`" + `array_enum` + "`" + ` for arrays.`,
				},
				resource.Attribute{
					Name:        "one_of",
					Description: `(Optional) Array of maps containing a mapping for display name to enum value.`,
				},
				resource.Attribute{
					Name:        "const",
					Description: `(Required) value mapping to member of ` + "`" + `enum` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `(Required) display name for the enum value.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the user schema property.`,
				},
				resource.Attribute{
					Name:        "required",
					Description: `(Optional) Whether the property is required for this application's users.`,
				},
				resource.Attribute{
					Name:        "min_length",
					Description: `(Optional) The minimum length of the user property value. Only applies to type ` + "`" + `"string"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max_length",
					Description: `(Optional) The maximum length of the user property value. Only applies to type ` + "`" + `"string"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) determines whether an app user attribute can be set at the Individual or Group Level.`,
				},
				resource.Attribute{
					Name:        "array_type",
					Description: `(Optional) The type of the array elements if ` + "`" + `type` + "`" + ` is set to ` + "`" + `"array"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "array_enum",
					Description: `(Optional) Array of values that an array property's items can be set to.`,
				},
				resource.Attribute{
					Name:        "array_one_of",
					Description: `(Optional) Display name and value an enum array can be set to.`,
				},
				resource.Attribute{
					Name:        "const",
					Description: `(Required) value mapping to member of ` + "`" + `enum` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `(Required) display name for the enum value.`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `(Optional) Access control permissions for the property. It can be set to ` + "`" + `"READ_WRITE"` + "`" + `, ` + "`" + `"READ_ONLY"` + "`" + `, ` + "`" + `"HIDE"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "master",
					Description: `(Optional) Master priority for the user schema property. It can be set to ` + "`" + `"PROFILE_MASTER"` + "`" + ` or ` + "`" + `"OKTA"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "external_name",
					Description: `(Optional) External name of the user schema property. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "app_id",
					Description: `ID of the application the user property is associated with.`,
				},
				resource.Attribute{
					Name:        "index",
					Description: `ID of the user schema property. ## Import App user schema property can be imported via the property index and app id. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_app_user_schema.example <app id>/<property name> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "app_id",
					Description: `ID of the application the user property is associated with.`,
				},
				resource.Attribute{
					Name:        "index",
					Description: `ID of the user schema property. ## Import App user schema property can be imported via the property index and app id. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_app_user_schema.example <app id>/<property name> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_auth_server",
			Category:         "Resources",
			ShortDescription: `Creates an Authorization Server.`,
			Description:      ``,
			Keywords: []string{
				"auth",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "audiences",
					Description: `(Required) The recipients that the tokens are intended for. This becomes the ` + "`" + `aud` + "`" + ` claim in an access token.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The status of the auth server. It defaults to ` + "`" + `"ACTIVE"` + "`" + ``,
				},
				resource.Attribute{
					Name:        "credentials_rotation_mode",
					Description: `(Optional) The key rotation mode for the authorization server. Can be ` + "`" + `"AUTO"` + "`" + ` or ` + "`" + `"MANUAL"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the authorization server.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the authorization server.`,
				},
				resource.Attribute{
					Name:        "issuer_mode",
					Description: `(Optional) Allows you to use a custom issuer URL. It can be set to ` + "`" + `"CUSTOM_URL"` + "`" + ` or ` + "`" + `"ORG_URL"` + "`" + ` ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the authorization server.`,
				},
				resource.Attribute{
					Name:        "kid",
					Description: `The ID of the JSON Web Key used for signing tokens issued by the authorization server.`,
				},
				resource.Attribute{
					Name:        "issuer",
					Description: `The complete URL for a Custom Authorization Server. This becomes the ` + "`" + `iss` + "`" + ` claim in an access token.`,
				},
				resource.Attribute{
					Name:        "credentials_last_rotated",
					Description: `The timestamp when the authorization server started to use the ` + "`" + `kid` + "`" + ` for signing tokens.`,
				},
				resource.Attribute{
					Name:        "credentials_next_rotation",
					Description: `The timestamp when the authorization server changes the key for signing tokens. Only returned when ` + "`" + `credentials_rotation_mode` + "`" + ` is ` + "`" + `"AUTO"` + "`" + `. ## Import Authorization Server can be imported via the Okta ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_auth_server.example <auth server id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the authorization server.`,
				},
				resource.Attribute{
					Name:        "kid",
					Description: `The ID of the JSON Web Key used for signing tokens issued by the authorization server.`,
				},
				resource.Attribute{
					Name:        "issuer",
					Description: `The complete URL for a Custom Authorization Server. This becomes the ` + "`" + `iss` + "`" + ` claim in an access token.`,
				},
				resource.Attribute{
					Name:        "credentials_last_rotated",
					Description: `The timestamp when the authorization server started to use the ` + "`" + `kid` + "`" + ` for signing tokens.`,
				},
				resource.Attribute{
					Name:        "credentials_next_rotation",
					Description: `The timestamp when the authorization server changes the key for signing tokens. Only returned when ` + "`" + `credentials_rotation_mode` + "`" + ` is ` + "`" + `"AUTO"` + "`" + `. ## Import Authorization Server can be imported via the Okta ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_auth_server.example <auth server id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_auth_server_claim",
			Category:         "Resources",
			ShortDescription: `Creates an Authorization Server Claim.`,
			Description:      ``,
			Keywords: []string{
				"auth",
				"server",
				"claim",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "auth_server_id",
					Description: `(Required) The Application's display name.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the claim.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of the claim.`,
				},
				resource.Attribute{
					Name:        "scopes",
					Description: `(Optional) The list of scopes the auth server claim is tied to.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The status of the application. It defaults to ` + "`" + `"ACTIVE"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "value_type",
					Description: `(Optional) The type of value of the claim. It can be set to ` + "`" + `"EXPRESSION"` + "`" + ` or ` + "`" + `"GROUPS"` + "`" + `. It defaults to ` + "`" + `"EXPRESSION"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "claim_type",
					Description: `(Required) Specifies whether the claim is for an access token ` + "`" + `"RESOURCE"` + "`" + ` or ID token ` + "`" + `"IDENTITY"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "always_include_in_token",
					Description: `(Optional) Specifies whether to include claims in token, by default is is set to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "group_filter_type",
					Description: `(Optional) Specifies the type of group filter if ` + "`" + `value_type` + "`" + ` is ` + "`" + `"GROUPS"` + "`" + `. Can be set to one of the following ` + "`" + `"STARTS_WITH"` + "`" + `, ` + "`" + `"EQUALS"` + "`" + `, ` + "`" + `"CONTAINS"` + "`" + `, ` + "`" + `"REGEX"` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID for the auth server claim.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the claim. ## Import Authorization Server Claim can be imported via the Auth Server ID and Claim ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_auth_server_claim.example <auth server id>/<claim id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID for the auth server claim.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the claim. ## Import Authorization Server Claim can be imported via the Auth Server ID and Claim ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_auth_server_claim.example <auth server id>/<claim id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_auth_server_policy",
			Category:         "Resources",
			ShortDescription: `Creates an Authorization Server Policy.`,
			Description:      ``,
			Keywords: []string{
				"auth",
				"server",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Auth Server Policy.`,
				},
				resource.Attribute{
					Name:        "auth_server_id",
					Description: `(Required) The ID of the Auth Server.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The status of the Auth Server Policy.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Required) The priority of the Auth Server Policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the Auth Server Policy.`,
				},
				resource.Attribute{
					Name:        "client_whitelist",
					Description: `(Required) The clients to whitelist the policy for. ` + "`" + `["ALL_CLIENTS"]` + "`" + ` is a special value that can be used to whitelist for all clients. Otherwise it is a list of client ids. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The ID of the authorization server policy.`,
				},
				resource.Attribute{
					Name:        "auth_server_id",
					Description: `(Required) The ID of the Auth Server.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the Auth Server Policy. ## Import Authorization Server Policy can be imported via the Auth Server ID and Policy ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_auth_server_policy.example <auth server id>/<policy id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The ID of the authorization server policy.`,
				},
				resource.Attribute{
					Name:        "auth_server_id",
					Description: `(Required) The ID of the Auth Server.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the Auth Server Policy. ## Import Authorization Server Policy can be imported via the Auth Server ID and Policy ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_auth_server_policy.example <auth server id>/<policy id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_auth_server_policy_rule",
			Category:         "Resources",
			ShortDescription: `Creates an Authorization Server Policy Rule.`,
			Description:      ``,
			Keywords: []string{
				"auth",
				"server",
				"policy",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Auth Server Policy Rule name.`,
				},
				resource.Attribute{
					Name:        "auth_server_id",
					Description: `(Required) Auth Server ID.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Required) Auth Server Policy ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The status of the Auth Server Policy Rule.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Required) Priority of the auth server policy rule.`,
				},
				resource.Attribute{
					Name:        "grant_type_whitelist",
					Description: `(Required) Accepted grant type values, ` + "`" + `"authorization_code"` + "`" + `, ` + "`" + `"implicit"` + "`" + `, ` + "`" + `"password"` + "`" + ``,
				},
				resource.Attribute{
					Name:        "scope_whitelist",
					Description: `(Required) Scopes allowed for this policy rule. They can be whitelisted by name or all can be whitelisted with ` + "`" + `"`,
				},
				resource.Attribute{
					Name:        "access_token_lifetime_minutes",
					Description: `(Optional) Lifetime of access token. Can be set to a value between 5 and 1440.`,
				},
				resource.Attribute{
					Name:        "refresh_token_lifetime_minutes",
					Description: `(Optional) Lifetime of refresh token.`,
				},
				resource.Attribute{
					Name:        "refresh_token_window_minutes",
					Description: `(Optional) Window in which a refresh token can be used. It can be a value between 10 and 2628000 (5 years).`,
				},
				resource.Attribute{
					Name:        "inline_hook_id",
					Description: `(Optional) The ID of the inline token to trigger. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The ID of the Auth Server Policy Rule.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Required) The ID of the Auth Server Policy.`,
				},
				resource.Attribute{
					Name:        "auth_server_id",
					Description: `(Required) The ID of the Auth Server.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the Auth Server Policy Rule. ## Import Authorization Server Policy Rule can be imported via the Auth Server ID, Policy ID, and Policy Rule ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_auth_server_policy_rule.example <auth server id>/<policy id>/<policy rule id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The ID of the Auth Server Policy Rule.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Required) The ID of the Auth Server Policy.`,
				},
				resource.Attribute{
					Name:        "auth_server_id",
					Description: `(Required) The ID of the Auth Server.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the Auth Server Policy Rule. ## Import Authorization Server Policy Rule can be imported via the Auth Server ID, Policy ID, and Policy Rule ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_auth_server_policy_rule.example <auth server id>/<policy id>/<policy rule id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_auth_server_scope",
			Category:         "Resources",
			ShortDescription: `Creates an Authorization Server Scope.`,
			Description:      ``,
			Keywords: []string{
				"auth",
				"server",
				"scope",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Auth Server scope name.`,
				},
				resource.Attribute{
					Name:        "auth_server_id",
					Description: `(Required) Auth Server ID.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Auth Server Scope.`,
				},
				resource.Attribute{
					Name:        "consent",
					Description: `(Optional) Indicates whether a consent dialog is needed for the scope. It can be set to ` + "`" + `"REQUIRED"` + "`" + ` or ` + "`" + `"IMPLICIT"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "metadata_publish",
					Description: `(Optional) Whether to publish metadata or not. It can be set to ` + "`" + `"ALL_CLIENTS"` + "`" + ` or ` + "`" + `"NO_CLIENTS"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `(Optional) A default scope will be returned in an access token when the client omits the scope parameter in a token request, provided this scope is allowed as part of the access policy rule. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Auth Server Scope.`,
				},
				resource.Attribute{
					Name:        "auth_server_id",
					Description: `The ID of the Auth Server. ## Import Okta Auth Server Scope can be imported via the Auth Server ID and Scope ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_auth_server_scope.example <auth server id>/<scope id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Auth Server Scope.`,
				},
				resource.Attribute{
					Name:        "auth_server_id",
					Description: `The ID of the Auth Server. ## Import Okta Auth Server Scope can be imported via the Auth Server ID and Scope ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_auth_server_scope.example <auth server id>/<scope id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_factor",
			Category:         "Resources",
			ShortDescription: `Allows you to manage the activation of Okta MFA methods.`,
			Description:      ``,
			Keywords: []string{
				"factor",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "provider",
					Description: `(Required) The MFA provider name.`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `(Optional) Whether or not to activate the provider, by default it is set to ` + "`" + `true` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "provider",
					Description: `MFA provider name.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "provider",
					Description: `MFA provider name.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_group",
			Category:         "Resources",
			ShortDescription: `Creates an Okta Group.`,
			Description:      ``,
			Keywords: []string{
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Okta Group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the Okta Group.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `(Optional) The users associated with the group. This can also be done per user. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Okta Group. ## Import An Okta Group can be imported via the Okta ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_group.example <group id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Okta Group. ## Import An Okta Group can be imported via the Okta ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_group.example <group id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_group_roles",
			Category:         "Resources",
			ShortDescription: `Creates Group level Admin Role Assignments.`,
			Description:      ``,
			Keywords: []string{
				"group",
				"roles",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required) The ID of group to attach admin roles to.`,
				},
				resource.Attribute{
					Name:        "admin_roles",
					Description: `(Required) Admin roles associated with the group. It can be any of the following values ` + "`" + `"SUPER_ADMIN"` + "`" + `, ` + "`" + `"ORG_ADMIN"` + "`" + `, ` + "`" + `"APP_ADMIN"` + "`" + `, ` + "`" + `"USER_ADMIN"` + "`" + `, ` + "`" + `"HELP_DESK_ADMIN"` + "`" + `, ` + "`" + `"READ_ONLY_ADMIN"` + "`" + `, ` + "`" + `"MOBILE_ADMIN"` + "`" + `, ` + "`" + `"API_ACCESS_MANAGEMENT_ADMIN"` + "`" + `, ` + "`" + `"REPORT_ADMIN"` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Group Role Assignment. ## Import Group Role Assignment can be imported via the Okta Group ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_group_roles.example <group id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Group Role Assignment. ## Import Group Role Assignment can be imported via the Okta Group ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_group_roles.example <group id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_group_rule",
			Category:         "Resources",
			ShortDescription: `Creates an Okta Group Rule.`,
			Description:      ``,
			Keywords: []string{
				"group",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Group Rule.`,
				},
				resource.Attribute{
					Name:        "group_assignments",
					Description: `(Required) The list of group ids to assign the users to.`,
				},
				resource.Attribute{
					Name:        "expression_type",
					Description: `(Optional) The expression type to use to invoke the rule. The default is ` + "`" + `"urn:okta:expression:1.0"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "expression_value",
					Description: `(Required) The expression value.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The status of the group rule. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Group Rule. ## Import An Okta Group Rule can be imported via the Okta ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_group_rule.example <group rule id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Group Rule. ## Import An Okta Group Rule can be imported via the Okta ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_group_rule.example <group rule id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_idp_oidc",
			Category:         "Resources",
			ShortDescription: `Creates an OIDC Identity Provider.`,
			Description:      ``,
			Keywords: []string{
				"idp",
				"oidc",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Application's display name.`,
				},
				resource.Attribute{
					Name:        "scopes",
					Description: `(Required) The scopes of the IdP.`,
				},
				resource.Attribute{
					Name:        "authorization_url",
					Description: `(Required) IdP Authorization Server (AS) endpoint to request consent from the user and obtain an authorization code grant.`,
				},
				resource.Attribute{
					Name:        "authorization_binding",
					Description: `(Required) The method of making an authorization request. It can be set to ` + "`" + `"HTTP-POST"` + "`" + ` or ` + "`" + `"HTTP-REDIRECT"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "token_url",
					Description: `(Required) IdP Authorization Server (AS) endpoint to exchange the authorization code grant for an access token.`,
				},
				resource.Attribute{
					Name:        "token_binding",
					Description: `(Required) The method of making a token request. It can be set to ` + "`" + `"HTTP-POST"` + "`" + ` or ` + "`" + `"HTTP-REDIRECT"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "jwks_url",
					Description: `(Required) Endpoint where the signer of the keys publishes its keys in a JWK Set.`,
				},
				resource.Attribute{
					Name:        "jwks_binding",
					Description: `(Required) The method of making a request for the OIDC JWKS. It can be set to ` + "`" + `"HTTP-POST"` + "`" + ` or ` + "`" + `"HTTP-REDIRECT"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "acs_binding",
					Description: `(Required) The method of making an ACS request. It can be set to ` + "`" + `"HTTP-POST"` + "`" + ` or ` + "`" + `"HTTP-REDIRECT"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `(Required) Unique identifier issued by AS for the Okta IdP instance.`,
				},
				resource.Attribute{
					Name:        "client_secret",
					Description: `(Required) Client secret issued by AS for the Okta IdP instance.`,
				},
				resource.Attribute{
					Name:        "issuer_url",
					Description: `(Required) URI that identifies the issuer.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Status of the IdP.`,
				},
				resource.Attribute{
					Name:        "user_info_url",
					Description: `(Optional) Protected resource endpoint that returns claims about the authenticated user.`,
				},
				resource.Attribute{
					Name:        "user_info_binding",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "acs_type",
					Description: `(Optional) The type of ACS. Default is ` + "`" + `"INSTANCE"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "protocol_type",
					Description: `(Optional) The type of protocol to use. It can be ` + "`" + `"OIDC"` + "`" + ` or ` + "`" + `"OAUTH2"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "issuer_mode",
					Description: `(Optional) Indicates whether Okta uses the original Okta org domain URL, or a custom domain URL. It can be ` + "`" + `"ORG_URL"` + "`" + ` or ` + "`" + `"CUSTOM_URL"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max_clock_skew",
					Description: `(Optional) Maximum allowable clock-skew when processing messages from the IdP.`,
				},
				resource.Attribute{
					Name:        "account_link_action",
					Description: `(Optional) Specifies the account linking action for an IdP user.`,
				},
				resource.Attribute{
					Name:        "account_link_group_include",
					Description: `(Optional) Group memberships to determine link candidates.`,
				},
				resource.Attribute{
					Name:        "provisioning_action",
					Description: `(Optional) Provisioning action for an IdP user during authentication.`,
				},
				resource.Attribute{
					Name:        "deprovisioned_action",
					Description: `(Optional) Action for a previously deprovisioned IdP user during authentication. Can be ` + "`" + `"NONE"` + "`" + ` or ` + "`" + `"REACTIVATE"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "suspended_action",
					Description: `(Optional) Action for a previously suspended IdP user during authentication. Can be set to ` + "`" + `"NONE"` + "`" + ` or ` + "`" + `"UNSUSPEND"` + "`" + ``,
				},
				resource.Attribute{
					Name:        "groups_action",
					Description: `(Optional) Provisioning action for IdP user's group memberships. It can be ` + "`" + `"NONE"` + "`" + `, ` + "`" + `"SYNC"` + "`" + `, ` + "`" + `"APPEND"` + "`" + `, or ` + "`" + `"ASSIGN"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "groups_attribute",
					Description: `(Optional) IdP user profile attribute name (case-insensitive) for an array value that contains group memberships.`,
				},
				resource.Attribute{
					Name:        "groups_assignment",
					Description: `(Optional) List of Okta Group IDs to add an IdP user as a member with the ` + "`" + `"ASSIGN"` + "`" + ` ` + "`" + `groups_action` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "groups_filter",
					Description: `(Optional) Whitelist of Okta Group identifiers that are allowed for the ` + "`" + `"APPEND"` + "`" + ` or ` + "`" + `"SYNC"` + "`" + ` ` + "`" + `groups_action` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "username_template",
					Description: `(Optional) Okta EL Expression to generate or transform a unique username for the IdP user.`,
				},
				resource.Attribute{
					Name:        "subject_match_type",
					Description: `(Optional) Determines the Okta user profile attribute match conditions for account linking and authentication of the transformed IdP username. By default it is set to ` + "`" + `"USERNAME"` + "`" + `. It can be set to ` + "`" + `"USERNAME"` + "`" + `, ` + "`" + `"EMAIL"` + "`" + `, ` + "`" + `"USERNAME_OR_EMAIL"` + "`" + ` or ` + "`" + `"CUSTOM_ATTRIBUTE"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "subject_match_attribute",
					Description: `(Optional) Okta user profile attribute for matching transformed IdP username. Only for matchType ` + "`" + `"CUSTOM_ATTRIBUTE"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "profile_master",
					Description: `(Optional) Determines if the IdP should act as a source of truth for user profile attributes. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the IdP.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of OIDC IdP. ## Import An OIDC IdP can be imported via the Okta ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_idp_oidc.example <idp id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the IdP.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of OIDC IdP. ## Import An OIDC IdP can be imported via the Okta ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_idp_oidc.example <idp id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_idp_saml",
			Category:         "Resources",
			ShortDescription: `Creates a SAML Identity Provider.`,
			Description:      ``,
			Keywords: []string{
				"idp",
				"saml",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Application's display name.`,
				},
				resource.Attribute{
					Name:        "kid",
					Description: `(Required) The ID of the signing key.`,
				},
				resource.Attribute{
					Name:        "acs_binding",
					Description: `(Required) The method of making an ACS request. It can be set to ` + "`" + `"HTTP-POST"` + "`" + ` or ` + "`" + `"HTTP-REDIRECT"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sso_url",
					Description: `(Required) URL of binding-specific endpoint to send an AuthnRequest message to IdP.`,
				},
				resource.Attribute{
					Name:        "issuer",
					Description: `(Required) URI that identifies the issuer.`,
				},
				resource.Attribute{
					Name:        "acs_type",
					Description: `(Optional) The type of ACS. It can be ` + "`" + `"INSTANCE"` + "`" + ` or ` + "`" + `"ORG"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sso_binding",
					Description: `(Optional) The method of making an SSO request. It can be set to ` + "`" + `"HTTP-POST"` + "`" + ` or ` + "`" + `"HTTP-REDIRECT"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sso_destination",
					Description: `(Optional) URI reference indicating the address to which the AuthnRequest message is sent.`,
				},
				resource.Attribute{
					Name:        "name_format",
					Description: `(Optional) The name identifier format to use. By default ` + "`" + `"urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "subject_format",
					Description: `(Optional) The name formate. By default ` + "`" + `"urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "subject_filter",
					Description: `(Optional) Optional regular expression pattern used to filter untrusted IdP usernames.`,
				},
				resource.Attribute{
					Name:        "issuer_mode",
					Description: `(Optional) Indicates whether Okta uses the original Okta org domain URL, or a custom domain URL. It can be ` + "`" + `"ORG_URL"` + "`" + ` or ` + "`" + `"CUSTOM_URL"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Status of the IdP.`,
				},
				resource.Attribute{
					Name:        "account_link_action",
					Description: `(Optional) Specifies the account linking action for an IdP user.`,
				},
				resource.Attribute{
					Name:        "account_link_group_include",
					Description: `(Optional) Group memberships to determine link candidates.`,
				},
				resource.Attribute{
					Name:        "provisioning_action",
					Description: `(Optional) Provisioning action for an IdP user during authentication.`,
				},
				resource.Attribute{
					Name:        "deprovisioned_action",
					Description: `(Optional) Action for a previously deprovisioned IdP user during authentication. Can be ` + "`" + `"NONE"` + "`" + ` or ` + "`" + `"REACTIVATE"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "suspended_action",
					Description: `(Optional) Action for a previously suspended IdP user during authentication. Can be set to ` + "`" + `"NONE"` + "`" + ` or ` + "`" + `"UNSUSPEND"` + "`" + ``,
				},
				resource.Attribute{
					Name:        "groups_action",
					Description: `(Optional) Provisioning action for IdP user's group memberships. It can be ` + "`" + `"NONE"` + "`" + `, ` + "`" + `"SYNC"` + "`" + `, ` + "`" + `"APPEND"` + "`" + `, or ` + "`" + `"ASSIGN"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "groups_attribute",
					Description: `(Optional) IdP user profile attribute name (case-insensitive) for an array value that contains group memberships.`,
				},
				resource.Attribute{
					Name:        "groups_assignment",
					Description: `(Optional) List of Okta Group IDs to add an IdP user as a member with the ` + "`" + `"ASSIGN"` + "`" + ` ` + "`" + `groups_action` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "groups_filter",
					Description: `(Optional) Whitelist of Okta Group identifiers that are allowed for the ` + "`" + `"APPEND"` + "`" + ` or ` + "`" + `"SYNC"` + "`" + ` ` + "`" + `groups_action` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "username_template",
					Description: `(Optional) Okta EL Expression to generate or transform a unique username for the IdP user.`,
				},
				resource.Attribute{
					Name:        "subject_match_type",
					Description: `(Optional) Determines the Okta user profile attribute match conditions for account linking and authentication of the transformed IdP username. By default it is set to ` + "`" + `"USERNAME"` + "`" + `. It can be set to ` + "`" + `"USERNAME"` + "`" + `, ` + "`" + `"EMAIL"` + "`" + `, ` + "`" + `"USERNAME_OR_EMAIL"` + "`" + ` or ` + "`" + `"CUSTOM_ATTRIBUTE"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "subject_match_attribute",
					Description: `(Optional) Okta user profile attribute for matching transformed IdP username. Only for matchType ` + "`" + `"CUSTOM_ATTRIBUTE"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "profile_master",
					Description: `(Optional) Determines if the IdP should act as a source of truth for user profile attributes.`,
				},
				resource.Attribute{
					Name:        "request_signature_algorithm",
					Description: `(Optional) The XML digital signature algorithm used when signing an AuthnRequest message.`,
				},
				resource.Attribute{
					Name:        "request_signature_scope",
					Description: `(Optional) Specifies whether or not to digitally sign an AuthnRequest messages to the IdP. It can be ` + "`" + `"REQUEST"` + "`" + ` or ` + "`" + `"NONE"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "response_signature_algorithm",
					Description: `(Optional) The minimum XML digital signature algorithm allowed when verifying a SAMLResponse message or Assertion element.`,
				},
				resource.Attribute{
					Name:        "response_signature_scope",
					Description: `(Optional) Specifies whether to verify a SAMLResponse message or Assertion element XML digital signature. It can be ` + "`" + `"RESPONSE"` + "`" + `, ` + "`" + `"ASSERTION"` + "`" + `, or ` + "`" + `"ANY"` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the IdP.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the IdP.`,
				},
				resource.Attribute{
					Name:        "audience",
					Description: `The audience restriction for the IdP. ## Import An SAML IdP can be imported via the Okta ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_idp_saml.example <idp id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the IdP.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the IdP.`,
				},
				resource.Attribute{
					Name:        "audience",
					Description: `The audience restriction for the IdP. ## Import An SAML IdP can be imported via the Okta ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_idp_saml.example <idp id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_idp_saml_signing_key",
			Category:         "Resources",
			ShortDescription: `Creates a SAML Identity Provider Signing Key.`,
			Description:      ``,
			Keywords: []string{
				"idp",
				"saml",
				"signing",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "x5c",
					Description: `(Required) base64-encoded X.509 certificate chain with DER encoding. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Key ID.`,
				},
				resource.Attribute{
					Name:        "kid",
					Description: `Key ID.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `Date created.`,
				},
				resource.Attribute{
					Name:        "expires_at",
					Description: `Date the cert expires.`,
				},
				resource.Attribute{
					Name:        "kty",
					Description: `Identifies the cryptographic algorithm family used with the key.`,
				},
				resource.Attribute{
					Name:        "use",
					Description: `Intended use of the public key.`,
				},
				resource.Attribute{
					Name:        "x5t_s256",
					Description: `base64url-encoded SHA-256 thumbprint of the DER encoding of an X.509 certificate. ## Import A SAML IdP Signing Key can be imported via the key id. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_idp_saml_signing_key.example <key id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Key ID.`,
				},
				resource.Attribute{
					Name:        "kid",
					Description: `Key ID.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `Date created.`,
				},
				resource.Attribute{
					Name:        "expires_at",
					Description: `Date the cert expires.`,
				},
				resource.Attribute{
					Name:        "kty",
					Description: `Identifies the cryptographic algorithm family used with the key.`,
				},
				resource.Attribute{
					Name:        "use",
					Description: `Intended use of the public key.`,
				},
				resource.Attribute{
					Name:        "x5t_s256",
					Description: `base64url-encoded SHA-256 thumbprint of the DER encoding of an X.509 certificate. ## Import A SAML IdP Signing Key can be imported via the key id. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_idp_saml_signing_key.example <key id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_idp_social",
			Category:         "Resources",
			ShortDescription: `Creates an Social Identity Provider.`,
			Description:      ``,
			Keywords: []string{
				"idp",
				"social",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Application's display name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of Social IdP. It can be ` + "`" + `"FACEBOOK"` + "`" + `, ` + "`" + `"LINKEDIN"` + "`" + `, ` + "`" + `"MICROSOFT"` + "`" + `, or ` + "`" + `"GOOGLE"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scopes",
					Description: `(Required) The scopes of the IdP.`,
				},
				resource.Attribute{
					Name:        "authorization_url",
					Description: `(Optional) IdP Authorization Server (AS) endpoint to request consent from the user and obtain an authorization code grant.`,
				},
				resource.Attribute{
					Name:        "authorization_binding",
					Description: `(Optional) The method of making an authorization request. It can be set to ` + "`" + `"HTTP-POST"` + "`" + ` or ` + "`" + `"HTTP-REDIRECT"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "token_url",
					Description: `(Optional) IdP Authorization Server (AS) endpoint to exchange the authorization code grant for an access token.`,
				},
				resource.Attribute{
					Name:        "token_binding",
					Description: `(Optional) The method of making a token request. It can be set to ` + "`" + `"HTTP-POST"` + "`" + ` or ` + "`" + `"HTTP-REDIRECT"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Status of the IdP.`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `(Optional) Unique identifier issued by AS for the Okta IdP instance.`,
				},
				resource.Attribute{
					Name:        "client_secret",
					Description: `(Optional) Client secret issued by AS for the Okta IdP instance.`,
				},
				resource.Attribute{
					Name:        "protocol_type",
					Description: `(Optional) The type of protocol to use. It can be ` + "`" + `"OIDC"` + "`" + ` or ` + "`" + `"OAUTH2"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "issuer_mode",
					Description: `(Optional) Indicates whether Okta uses the original Okta org domain URL, or a custom domain URL. It can be ` + "`" + `"ORG_URL"` + "`" + ` or ` + "`" + `"CUSTOM_URL"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max_clock_skew",
					Description: `(Optional) Maximum allowable clock-skew when processing messages from the IdP.`,
				},
				resource.Attribute{
					Name:        "account_link_action",
					Description: `(Optional) Specifies the account linking action for an IdP user.`,
				},
				resource.Attribute{
					Name:        "account_link_group_include",
					Description: `(Optional) Group memberships to determine link candidates.`,
				},
				resource.Attribute{
					Name:        "provisioning_action",
					Description: `(Optional) Provisioning action for an IdP user during authentication.`,
				},
				resource.Attribute{
					Name:        "deprovisioned_action",
					Description: `(Optional) Action for a previously deprovisioned IdP user during authentication. Can be ` + "`" + `"NONE"` + "`" + ` or ` + "`" + `"REACTIVATE"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "suspended_action",
					Description: `(Optional) Action for a previously suspended IdP user during authentication. Can be set to ` + "`" + `"NONE"` + "`" + ` or ` + "`" + `"UNSUSPEND"` + "`" + ``,
				},
				resource.Attribute{
					Name:        "groups_action",
					Description: `(Optional) Provisioning action for IdP user's group memberships. It can be ` + "`" + `"NONE"` + "`" + `, ` + "`" + `"SYNC"` + "`" + `, ` + "`" + `"APPEND"` + "`" + `, or ` + "`" + `"ASSIGN"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "groups_attribute",
					Description: `(Optional) IdP user profile attribute name (case-insensitive) for an array value that contains group memberships.`,
				},
				resource.Attribute{
					Name:        "groups_assignment",
					Description: `(Optional) List of Okta Group IDs to add an IdP user as a member with the ` + "`" + `"ASSIGN"` + "`" + ` ` + "`" + `groups_action` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "groups_filter",
					Description: `(Optional) Whitelist of Okta Group identifiers that are allowed for the ` + "`" + `"APPEND"` + "`" + ` or ` + "`" + `"SYNC"` + "`" + ` ` + "`" + `groups_action` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "username_template",
					Description: `(Optional) Okta EL Expression to generate or transform a unique username for the IdP user.`,
				},
				resource.Attribute{
					Name:        "subject_match_type",
					Description: `(Optional) Determines the Okta user profile attribute match conditions for account linking and authentication of the transformed IdP username. By default it is set to ` + "`" + `"USERNAME"` + "`" + `. It can be set to ` + "`" + `"USERNAME"` + "`" + `, ` + "`" + `"EMAIL"` + "`" + `, ` + "`" + `"USERNAME_OR_EMAIL"` + "`" + ` or ` + "`" + `"CUSTOM_ATTRIBUTE"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "subject_match_attribute",
					Description: `(Optional) Okta user profile attribute for matching transformed IdP username. Only for matchType ` + "`" + `"CUSTOM_ATTRIBUTE"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "profile_master",
					Description: `(Optional) Determines if the IdP should act as a source of truth for user profile attributes.`,
				},
				resource.Attribute{
					Name:        "request_signature_algorithm",
					Description: `(Optional) The XML digital signature algorithm used when signing an AuthnRequest message.`,
				},
				resource.Attribute{
					Name:        "request_signature_scope",
					Description: `(Optional) Specifies whether or not to digitally sign an AuthnRequest messages to the IdP. It can be ` + "`" + `"REQUEST"` + "`" + ` or ` + "`" + `"NONE"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "response_signature_algorithm",
					Description: `(Optional) The minimum XML digital signature algorithm allowed when verifying a SAMLResponse message or Assertion element.`,
				},
				resource.Attribute{
					Name:        "response_signature_scope",
					Description: `(Optional) Specifies whether to verify a SAMLResponse message or Assertion element XML digital signature. It can be ` + "`" + `"RESPONSE"` + "`" + `, ` + "`" + `"ASSERTION"` + "`" + `, or ` + "`" + `"ANY"` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the IdP. ## Import A Social IdP can be imported via the Okta ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_idp_social.example <idp id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the IdP. ## Import A Social IdP can be imported via the Okta ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_idp_social.example <idp id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_inline_hook",
			Category:         "Resources",
			ShortDescription: `Creates an inline hook.`,
			Description:      ``,
			Keywords: []string{
				"inline",
				"hook",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The inline hook display name.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The version of the hook.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of hook to create. [See here for supported types](https://developer.okta.com/docs/reference/api/inline-hooks/#supported-inline-hook-types).`,
				},
				resource.Attribute{
					Name:        "headers",
					Description: `(Optional) Map of headers to send along in inline hook request.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Header name.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Header value.`,
				},
				resource.Attribute{
					Name:        "auth",
					Description: `(Optional) Authentication required for inline hook request.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Key to use for authentication, usually the header name, for example ` + "`" + `"Authorization"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Authentication secret.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Auth type, default is ` + "`" + `"HEADER"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "channel",
					Description: `(Optional) Details of the endpoint the inline hook will hit.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The version of the endpoint.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `(Required) The URI the hook will hit.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type of hook to trigger. Currently only ` + "`" + `"HTTP"` + "`" + ` is supported.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Optional) The request method to use. Default is ` + "`" + `"POST"` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the inline hooks. ## Import An inline hook can be imported via the Okta ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_inline_hook.example <hook id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the inline hooks. ## Import An inline hook can be imported via the Okta ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_inline_hook.example <hook id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_network_zone",
			Category:         "Resources",
			ShortDescription: `Creates an Okta Network Zone.`,
			Description:      ``,
			Keywords: []string{
				"network",
				"zone",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Network Zone Resource.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type of the Network Zone - can either be IP or DYNAMIC only.`,
				},
				resource.Attribute{
					Name:        "dynamic_locations",
					Description: `(Optional) Array of locations ISO-3166-1(2). Format code: countryCode OR countryCode-regionCode.`,
				},
				resource.Attribute{
					Name:        "gateways",
					Description: `(Optional) Array of values in CIDR/range form.`,
				},
				resource.Attribute{
					Name:        "proxies",
					Description: `(Optional) Array of values in CIDR/range form. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Network Zone ID. ## Import Okta Network Zone can be imported via the Okta ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_network_zone.example <zone id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Network Zone ID. ## Import Okta Network Zone can be imported via the Okta ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_network_zone.example <zone id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_policy_mfa",
			Category:         "Resources",
			ShortDescription: `Creates an MFA Policy.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"mfa",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Policy Name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Policy Description.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Priority of the policy.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Policy Status: ` + "`" + `"ACTIVE"` + "`" + ` or ` + "`" + `"INACTIVE"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "groups_included",
					Description: `(Optional) List of Group IDs to Include.`,
				},
				resource.Attribute{
					Name:        "duo",
					Description: `(Optional) DUO [MFA policy settings](#mfa-settings).`,
				},
				resource.Attribute{
					Name:        "fido_u2f",
					Description: `(Optional) Fido U2F [MFA policy settings](#mfa-settings).`,
				},
				resource.Attribute{
					Name:        "fido_webauthn",
					Description: `(Optional) Fido Web Authn [MFA policy settings](#mfa-settings).`,
				},
				resource.Attribute{
					Name:        "google_otp",
					Description: `(Optional) Google OTP [MFA policy settings](#mfa-settings).`,
				},
				resource.Attribute{
					Name:        "okta_call",
					Description: `(Optional) Okta Call [MFA policy settings](#mfa-settings).`,
				},
				resource.Attribute{
					Name:        "okta_otp",
					Description: `(Optional) Okta OTP [MFA policy settings](#mfa-settings).`,
				},
				resource.Attribute{
					Name:        "okta_password",
					Description: `(Optional) Okta Password [MFA policy settings](#mfa-settings).`,
				},
				resource.Attribute{
					Name:        "okta_push",
					Description: `(Optional) Okta Push [MFA policy settings](#mfa-settings).`,
				},
				resource.Attribute{
					Name:        "okta_question",
					Description: `(Optional) Okta Question [MFA policy settings](#mfa-settings).`,
				},
				resource.Attribute{
					Name:        "okta_sms",
					Description: `(Optional) Okta SMS [MFA policy settings](#mfa-settings).`,
				},
				resource.Attribute{
					Name:        "rsa_token",
					Description: `(Optional) RSA Token [MFA policy settings](#mfa-settings).`,
				},
				resource.Attribute{
					Name:        "symantec_vip",
					Description: `(Optional) Symantec VIP [MFA policy settings](#mfa-settings).`,
				},
				resource.Attribute{
					Name:        "yubikey_token",
					Description: `(Optional) Yubikey Token [MFA policy settings](#mfa-settings). ### MFA Settings All MFA settings above have the following structure.`,
				},
				resource.Attribute{
					Name:        "enroll",
					Description: `(Optional) Requirements for user initiated enrollment. Can be ` + "`" + `"NOT_ALLOWED"` + "`" + `, ` + "`" + `"OPTIONAL"` + "`" + `, or ` + "`" + `"REQUIRED"` + "`" + `. By default it is ` + "`" + `"OPTIONAL"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "consent_type",
					Description: `(Optional) User consent type required before enrolling in the factor: ` + "`" + `"NONE"` + "`" + ` or ` + "`" + `"TERMS_OF_SERVICE"` + "`" + `. By default it is ` + "`" + `"NONE"` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Policy. ## Import An MFA Policy can be imported via the Okta ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_policy_mfa.example <app id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Policy. ## Import An MFA Policy can be imported via the Okta ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_policy_mfa.example <app id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_policy_password",
			Category:         "Resources",
			ShortDescription: `Creates a Password Policy.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"password",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Policy Name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Policy Description.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Priority of the policy.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Policy Status: ` + "`" + `"ACTIVE"` + "`" + ` or ` + "`" + `"INACTIVE"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "groups_included",
					Description: `(Optional) List of Group IDs to Include.`,
				},
				resource.Attribute{
					Name:        "auth_provider",
					Description: `(Optional) Authentication Provider: ` + "`" + `"OKTA"` + "`" + ` or ` + "`" + `"ACTIVE_DIRECTORY"` + "`" + `. Default is ` + "`" + `"OKTA"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "password_min_length",
					Description: `(Optional) Minimum password length. Default is 8.`,
				},
				resource.Attribute{
					Name:        "password_min_lowercase",
					Description: `(Optional) Minimum number of lower case characters in password.`,
				},
				resource.Attribute{
					Name:        "password_min_uppercase",
					Description: `(Optional) Minimum number of upper case characters in password.`,
				},
				resource.Attribute{
					Name:        "password_min_number",
					Description: `(Optional) Minimum number of numbers in password.`,
				},
				resource.Attribute{
					Name:        "password_min_symbol",
					Description: `(Optional) Minimum number of symbols in password.`,
				},
				resource.Attribute{
					Name:        "password_exclude_username",
					Description: `(Optional) If the user name must be excluded from the password.`,
				},
				resource.Attribute{
					Name:        "password_exclude_first_name",
					Description: `(Optional) User firstName attribute must be excluded from the password.`,
				},
				resource.Attribute{
					Name:        "password_exclude_last_name",
					Description: `(Optional) User lastName attribute must be excluded from the password.`,
				},
				resource.Attribute{
					Name:        "password_dictionary_lookup",
					Description: `(Optional) Check Passwords Against Common Password Dictionary.`,
				},
				resource.Attribute{
					Name:        "password_max_age_days",
					Description: `(Optional) Length in days a password is valid before expiry: 0 = no limit.",`,
				},
				resource.Attribute{
					Name:        "password_expire_warn_days",
					Description: `(Optional) Length in days a user will be warned before password expiry: 0 = no warning.`,
				},
				resource.Attribute{
					Name:        "password_min_age_minutes",
					Description: `(Optional) Minimum time interval in minutes between password changes: 0 = no limit.`,
				},
				resource.Attribute{
					Name:        "password_history_count",
					Description: `(Optional) Number of distinct passwords that must be created before they can be reused: 0 = none.`,
				},
				resource.Attribute{
					Name:        "password_max_lockout_attempts",
					Description: `(Optional) Number of unsuccessful login attempts allowed before lockout: 0 = no limit.`,
				},
				resource.Attribute{
					Name:        "password_auto_unlock_minutes",
					Description: `(Optional) Number of minutes before a locked account is unlocked: 0 = no limit.`,
				},
				resource.Attribute{
					Name:        "password_show_lockout_failures",
					Description: `(Optional) If a user should be informed when their account is locked.`,
				},
				resource.Attribute{
					Name:        "question_min_length",
					Description: `(Optional) Min length of the password recovery question answer.`,
				},
				resource.Attribute{
					Name:        "email_recovery",
					Description: `(Optional) Enable or disable email password recovery: ACTIVE or INACTIVE.`,
				},
				resource.Attribute{
					Name:        "recovery_email_token",
					Description: `(Optional) Lifetime in minutes of the recovery email token.`,
				},
				resource.Attribute{
					Name:        "sms_recovery",
					Description: `(Optional) Enable or disable SMS password recovery: ACTIVE or INACTIVE.`,
				},
				resource.Attribute{
					Name:        "question_recovery",
					Description: `(Optional) Enable or disable security question password recovery: ACTIVE or INACTIVE.`,
				},
				resource.Attribute{
					Name:        "skip_unlock",
					Description: `(Optional) When an Active Directory user is locked out of Okta, the Okta unlock operation should also attempt to unlock the user's Windows account. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Policy. ## Import A Password Policy can be imported via the Okta ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_policy_password.example <policy id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Policy. ## Import A Password Policy can be imported via the Okta ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_policy_password.example <policy id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_policy_rule_idp_discovery",
			Category:         "Resources",
			ShortDescription: `Creates an IdP Discovery Policy Rule.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"rule",
				"idp",
				"discovery",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policyid",
					Description: `(Required) Policy ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Policy Rule Name.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Policy Rule Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid priority is provided. API defaults it to the last/lowest if not there.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Policy Rule Status: ` + "`" + `"ACTIVE"` + "`" + ` or ` + "`" + `"INACTIVE"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "network_connection",
					Description: `(Optional) Network selection mode: ` + "`" + `"ANYWHERE"` + "`" + `, ` + "`" + `"ZONE"` + "`" + `, ` + "`" + `"ON_NETWORK"` + "`" + `, or ` + "`" + `"OFF_NETWORK"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "network_includes",
					Description: `(Optional) The network zones to include. Conflicts with ` + "`" + `network_excludes` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "network_excludes",
					Description: `(Optional) The network zones to exclude. Conflicts with ` + "`" + `network_includes` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Rule.`,
				},
				resource.Attribute{
					Name:        "policyid",
					Description: `Policy ID. ## Import A Policy Rule can be imported via the Policy and Rule ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_policy_rule_idp_discovery.example <policy id>/<rule id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Rule.`,
				},
				resource.Attribute{
					Name:        "policyid",
					Description: `Policy ID. ## Import A Policy Rule can be imported via the Policy and Rule ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_policy_rule_idp_discovery.example <policy id>/<rule id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_policy_rule_mfa",
			Category:         "Resources",
			ShortDescription: `Creates an MFA Policy Rule.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"rule",
				"mfa",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policyid",
					Description: `(Required) Policy ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Policy Rule Name.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Policy Rule Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid priority is provided. API defaults it to the last/lowest if not there.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Policy Rule Status: ` + "`" + `"ACTIVE"` + "`" + ` or ` + "`" + `"INACTIVE"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enroll",
					Description: `(Optional) When a user should be prompted for MFA. It can be ` + "`" + `"CHALLENGE"` + "`" + `, ` + "`" + `"LOGIN"` + "`" + `, or ` + "`" + `"NEVER"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "network_connection",
					Description: `(Optional) Network selection mode: ` + "`" + `"ANYWHERE"` + "`" + `, ` + "`" + `"ZONE"` + "`" + `, ` + "`" + `"ON_NETWORK"` + "`" + `, or ` + "`" + `"OFF_NETWORK"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "network_includes",
					Description: `(Optional) The network zones to include. Conflicts with ` + "`" + `network_excludes` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "network_excludes",
					Description: `(Optional) The network zones to exclude. Conflicts with ` + "`" + `network_includes` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Rule.`,
				},
				resource.Attribute{
					Name:        "policyid",
					Description: `Policy ID. ## Import A Policy Rule can be imported via the Policy and Rule ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_policy_rule_mfa.example <policy id>/<rule id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Rule.`,
				},
				resource.Attribute{
					Name:        "policyid",
					Description: `Policy ID. ## Import A Policy Rule can be imported via the Policy and Rule ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_policy_rule_mfa.example <policy id>/<rule id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_policy_rule_password",
			Category:         "Resources",
			ShortDescription: `Creates a Password Policy Rule.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"rule",
				"password",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policyid",
					Description: `(Required) Policy ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Policy Rule Name.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Policy Rule Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid priority is provided. API defaults it to the last/lowest if not there.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Policy Rule Status: ` + "`" + `"ACTIVE"` + "`" + ` or ` + "`" + `"INACTIVE"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "password_change",
					Description: `(Optional) Allow or deny a user to change their password: ` + "`" + `"ALLOW"` + "`" + ` or ` + "`" + `"DENY"` + "`" + `. By default it is ` + "`" + `"ALLOW"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "password_reset",
					Description: `(Optional) Allow or deny a user to reset their password: ` + "`" + `"ALLOW"` + "`" + ` or ` + "`" + `"DENY"` + "`" + `. By default it is ` + "`" + `"ALLOW"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "password_unlock",
					Description: `(Optional) Allow or deny a user to unlock: ` + "`" + `"ALLOW"` + "`" + ` or ` + "`" + `"DENY"` + "`" + `. By default it is ` + "`" + `"DENY"` + "`" + `,`,
				},
				resource.Attribute{
					Name:        "network_connection",
					Description: `(Optional) Network selection mode: ` + "`" + `"ANYWHERE"` + "`" + `, ` + "`" + `"ZONE"` + "`" + `, ` + "`" + `"ON_NETWORK"` + "`" + `, or ` + "`" + `"OFF_NETWORK"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "network_includes",
					Description: `(Optional) The network zones to include. Conflicts with ` + "`" + `network_excludes` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "network_excludes",
					Description: `(Optional) The network zones to exclude. Conflicts with ` + "`" + `network_includes` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Rule.`,
				},
				resource.Attribute{
					Name:        "policyid",
					Description: `Policy ID. ## Import A Policy Rule can be imported via the Policy and Rule ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_policy_rule_password.example <policy id>/<rule id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Rule.`,
				},
				resource.Attribute{
					Name:        "policyid",
					Description: `Policy ID. ## Import A Policy Rule can be imported via the Policy and Rule ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_policy_rule_password.example <policy id>/<rule id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_policy_rule_signon",
			Category:         "Resources",
			ShortDescription: `Creates a Sign On Policy Rule.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"rule",
				"signon",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policyid",
					Description: `(Required) Policy ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Policy Rule Name.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Policy Rule Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid priority is provided. API defaults it to the last/lowest if not there.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Policy Rule Status: ` + "`" + `"ACTIVE"` + "`" + ` or ` + "`" + `"INACTIVE"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "authtype",
					Description: `(Optional) Authentication entrypoint: ` + "`" + `"ANY"` + "`" + ` or ` + "`" + `"RADIUS"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "access",
					Description: `(Optional) Allow or deny access based on the rule conditions: ` + "`" + `"ALLOW"` + "`" + ` or ` + "`" + `"DENY"` + "`" + `. The default is ` + "`" + `"ALLOW"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mfa_required",
					Description: `(Optional) Require MFA. By default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mfa_prompt",
					Description: `(Optional) Prompt for MFA based on the device used, a factor session lifetime, or every sign on attempt: ` + "`" + `"DEVICE"` + "`" + `, ` + "`" + `"SESSION"` + "`" + ` or ` + "`" + `"ALWAYS"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mfa_remember_device",
					Description: `(Optional) Remember MFA device. The default ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mfa_lifetime",
					Description: `(Optional) Elapsed time before the next MFA challenge.`,
				},
				resource.Attribute{
					Name:        "session_idle",
					Description: `(Optional) Max minutes a session can be idle.",`,
				},
				resource.Attribute{
					Name:        "session_lifetime",
					Description: `(Optional) Max minutes a session is active: Disable = 0.`,
				},
				resource.Attribute{
					Name:        "session_persistent",
					Description: `(Optional) Whether session cookies will last across browser sessions. Okta Administrators can never have persistent session cookies.`,
				},
				resource.Attribute{
					Name:        "network_connection",
					Description: `(Optional) Network selection mode: ` + "`" + `"ANYWHERE"` + "`" + `, ` + "`" + `"ZONE"` + "`" + `, ` + "`" + `"ON_NETWORK"` + "`" + `, or ` + "`" + `"OFF_NETWORK"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "network_includes",
					Description: `(Optional) The network zones to include. Conflicts with ` + "`" + `network_excludes` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "network_excludes",
					Description: `(Optional) The network zones to exclude. Conflicts with ` + "`" + `network_includes` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Rule.`,
				},
				resource.Attribute{
					Name:        "policyid",
					Description: `Policy ID. ## Import A Policy Rule can be imported via the Policy and Rule ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_policy_rule_signon.example <policy id>/<rule id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Rule.`,
				},
				resource.Attribute{
					Name:        "policyid",
					Description: `Policy ID. ## Import A Policy Rule can be imported via the Policy and Rule ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_policy_rule_signon.example <policy id>/<rule id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_policy_signon",
			Category:         "Resources",
			ShortDescription: `Creates a Sign On Policy.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"signon",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Policy Name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Policy Description.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Priority of the policy.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Policy Status: ` + "`" + `"ACTIVE"` + "`" + ` or ` + "`" + `"INACTIVE"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "groups_included",
					Description: `List of Group IDs to Include. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Policy. ## Import A Sign On Policy can be imported via the Okta ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_policy_signon.example <policy id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Policy. ## Import A Sign On Policy can be imported via the Okta ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_policy_signon.example <policy id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_profile_mapping",
			Category:         "Resources",
			ShortDescription: `Manages a profile mapping.`,
			Description:      ``,
			Keywords: []string{
				"profile",
				"mapping",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "source_id",
					Description: `(Required) Source id of the profile mapping.`,
				},
				resource.Attribute{
					Name:        "delete_when_absent",
					Description: `(Optional) Tells the provider whether to attempt to delete missing mappings under profile mapping.`,
				},
				resource.Attribute{
					Name:        "mappings",
					Description: `(Optional) Priority of the policy.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) Key of mapping.`,
				},
				resource.Attribute{
					Name:        "expression",
					Description: `(Required) Combination or single source properties that will be mapped to the target property.`,
				},
				resource.Attribute{
					Name:        "push_status",
					Description: `(Optional) Whether to update target properties on user create & update or just on create. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the mappings.`,
				},
				resource.Attribute{
					Name:        "target_id",
					Description: `ID of the mapping target.`,
				},
				resource.Attribute{
					Name:        "target_name",
					Description: `Name of the mapping target.`,
				},
				resource.Attribute{
					Name:        "target_type",
					Description: `ID of the mapping target.`,
				},
				resource.Attribute{
					Name:        "source_id",
					Description: `ID of the mapping source.`,
				},
				resource.Attribute{
					Name:        "source_name",
					Description: `Name of the mapping source.`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `ID of the mapping source. ## Import There is no reason to import this resource. You can simply create the resource config and point it to a source ID. Once the source is deleted this resources will no longer exist.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the mappings.`,
				},
				resource.Attribute{
					Name:        "target_id",
					Description: `ID of the mapping target.`,
				},
				resource.Attribute{
					Name:        "target_name",
					Description: `Name of the mapping target.`,
				},
				resource.Attribute{
					Name:        "target_type",
					Description: `ID of the mapping target.`,
				},
				resource.Attribute{
					Name:        "source_id",
					Description: `ID of the mapping source.`,
				},
				resource.Attribute{
					Name:        "source_name",
					Description: `Name of the mapping source.`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `ID of the mapping source. ## Import There is no reason to import this resource. You can simply create the resource config and point it to a source ID. Once the source is deleted this resources will no longer exist.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_template_email",
			Category:         "Resources",
			ShortDescription: `Creates an Okta Email Template.`,
			Description:      ``,
			Keywords: []string{
				"template",
				"email",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Email template type`,
				},
				resource.Attribute{
					Name:        "translations",
					Description: `(Required) Set of translations for particular template.`,
				},
				resource.Attribute{
					Name:        "language",
					Description: `(Required) The language to map tthe template to.`,
				},
				resource.Attribute{
					Name:        "subject",
					Description: `(Required) The email subject line.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Required) The email body.`,
				},
				resource.Attribute{
					Name:        "default_language",
					Description: `(Optional) The default language, by default is set to ` + "`" + `"en"` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Email Template. ## Import An Okta Email Template can be imported via the template type. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_template_email.example <template type> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Email Template. ## Import An Okta Email Template can be imported via the template type. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_template_email.example <template type> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_trusted_origin",
			Category:         "Resources",
			ShortDescription: `Creates a Trusted Origin.`,
			Description:      ``,
			Keywords: []string{
				"trusted",
				"origin",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "active",
					Description: `(Optional) Whether the Trusted Origin is active or not - can only be issued post-creation.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Trusted Origin Resource.`,
				},
				resource.Attribute{
					Name:        "origin",
					Description: `(Required) The origin to trust.`,
				},
				resource.Attribute{
					Name:        "scopes",
					Description: `(Required) Scopes of the Trusted Origin - can be ` + "`" + `"CORS"` + "`" + ` and/or ` + "`" + `"REDIRECT"` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Trusted Origin. ## Import A Trusted Origin can be imported via the Okta ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_trusted_origin.example <trusted origin id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Trusted Origin. ## Import A Trusted Origin can be imported via the Okta ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_trusted_origin.example <trusted origin id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_user",
			Category:         "Resources",
			ShortDescription: `Creates an Okta User.`,
			Description:      ``,
			Keywords: []string{
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "email",
					Description: `(Required) User profile property.`,
				},
				resource.Attribute{
					Name:        "login",
					Description: `(Required) User profile property.`,
				},
				resource.Attribute{
					Name:        "first_name",
					Description: `(Required) User's First Name, required by default.`,
				},
				resource.Attribute{
					Name:        "last_name",
					Description: `(Required) User's Last Name, required by default.`,
				},
				resource.Attribute{
					Name:        "custom_profile_attributes",
					Description: `(Optional) raw JSON containing all custom profile attributes.`,
				},
				resource.Attribute{
					Name:        "admin_roles",
					Description: `(Optional) Administrator roles assigned to User.`,
				},
				resource.Attribute{
					Name:        "city",
					Description: `(Optional) User profile property.`,
				},
				resource.Attribute{
					Name:        "cost_center",
					Description: `(Optional) User profile property.`,
				},
				resource.Attribute{
					Name:        "country_code",
					Description: `(Optional) User profile property.`,
				},
				resource.Attribute{
					Name:        "department",
					Description: `(Optional) User profile property.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) User profile property.`,
				},
				resource.Attribute{
					Name:        "division",
					Description: `(Optional) User profile property.`,
				},
				resource.Attribute{
					Name:        "employee_number",
					Description: `(Optional) User profile property.`,
				},
				resource.Attribute{
					Name:        "group_memberships",
					Description: `(Optional) User profile property.`,
				},
				resource.Attribute{
					Name:        "honorific_prefix",
					Description: `(Optional) User profile property.`,
				},
				resource.Attribute{
					Name:        "honorific_suffix",
					Description: `(Optional) User profile property.`,
				},
				resource.Attribute{
					Name:        "locale",
					Description: `(Optional) User profile property.`,
				},
				resource.Attribute{
					Name:        "manager",
					Description: `(Optional) User profile property.`,
				},
				resource.Attribute{
					Name:        "manager_id",
					Description: `(Optional) User profile property.`,
				},
				resource.Attribute{
					Name:        "middle_name",
					Description: `(Optional) User profile property.`,
				},
				resource.Attribute{
					Name:        "mobile_phone",
					Description: `(Optional) User profile property.`,
				},
				resource.Attribute{
					Name:        "nick_name",
					Description: `(Optional) User profile property.`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `(Optional) User profile property.`,
				},
				resource.Attribute{
					Name:        "postal_address",
					Description: `(Optional) User profile property.`,
				},
				resource.Attribute{
					Name:        "preferred_language",
					Description: `(Optional) User profile property.`,
				},
				resource.Attribute{
					Name:        "primary_phone",
					Description: `(Optional) User profile property.`,
				},
				resource.Attribute{
					Name:        "profile_url",
					Description: `(Optional) User profile property.`,
				},
				resource.Attribute{
					Name:        "second_email",
					Description: `(Optional) User profile property.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) User profile property.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) User profile property.`,
				},
				resource.Attribute{
					Name:        "street_address",
					Description: `(Optional) User profile property.`,
				},
				resource.Attribute{
					Name:        "timezone",
					Description: `(Optional) User profile property.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `(Optional) User profile property.`,
				},
				resource.Attribute{
					Name:        "user_type",
					Description: `(Optional) User profile property.`,
				},
				resource.Attribute{
					Name:        "zip_code",
					Description: `(Optional) User profile property.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) User password.`,
				},
				resource.Attribute{
					Name:        "recovery_question",
					Description: `(Optional) User password recovery question.`,
				},
				resource.Attribute{
					Name:        "recovery_answer",
					Description: `(Optional) User password recovery answer. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "index",
					Description: `(Optional) ID of the User schema property. ## Import An Okta User can be imported via the ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_user.example <user id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "index",
					Description: `(Optional) ID of the User schema property. ## Import An Okta User can be imported via the ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_user.example <user id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_user_base_schema",
			Category:         "Resources",
			ShortDescription: `Manages a User Base Schema property.`,
			Description:      ``,
			Keywords: []string{
				"user",
				"base",
				"schema",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "index",
					Description: `(Required) The property name.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `(Required) The property display name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of the schema property. It can be ` + "`" + `"string"` + "`" + `, ` + "`" + `"boolean"` + "`" + `, ` + "`" + `"number"` + "`" + `, ` + "`" + `"integer"` + "`" + `, ` + "`" + `"array"` + "`" + `, or ` + "`" + `"object"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "required",
					Description: `(Optional) Whether the property is required for this application's users.`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `(Optional) Access control permissions for the property. It can be set to ` + "`" + `"READ_WRITE"` + "`" + `, ` + "`" + `"READ_ONLY"` + "`" + `, ` + "`" + `"HIDE"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "master",
					Description: `(Optional) Master priority for the user schema property. It can be set to ` + "`" + `"PROFILE_MASTER"` + "`" + ` or ` + "`" + `"OKTA"` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "index",
					Description: `ID of the user schema property. ## Import User base schema property can be imported via the property index. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_user_base_schema.example <property name> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "index",
					Description: `ID of the user schema property. ## Import User base schema property can be imported via the property index. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_user_base_schema.example <property name> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_user_schema",
			Category:         "Resources",
			ShortDescription: `Creates an User Schema property.`,
			Description:      ``,
			Keywords: []string{
				"user",
				"schema",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "index",
					Description: `(Required) The property name.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `(Required) The display name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of the schema property. It can be ` + "`" + `"string"` + "`" + `, ` + "`" + `"boolean"` + "`" + `, ` + "`" + `"number"` + "`" + `, ` + "`" + `"integer"` + "`" + `, ` + "`" + `"array"` + "`" + `, or ` + "`" + `"object"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enum",
					Description: `(Optional) Array of values a primitive property can be set to. See ` + "`" + `array_enum` + "`" + ` for arrays.`,
				},
				resource.Attribute{
					Name:        "one_of",
					Description: `(Optional) Array of maps containing a mapping for display name to enum value.`,
				},
				resource.Attribute{
					Name:        "const",
					Description: `(Required) value mapping to member of ` + "`" + `enum` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `(Required) display name for the enum value.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the user schema property.`,
				},
				resource.Attribute{
					Name:        "required",
					Description: `(Optional) Whether the property is required for this application's users.`,
				},
				resource.Attribute{
					Name:        "min_length",
					Description: `(Optional) The minimum length of the user property value. Only applies to type ` + "`" + `"string"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max_length",
					Description: `(Optional) The maximum length of the user property value. Only applies to type ` + "`" + `"string"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) determines whether an app user attribute can be set at the Individual or Group Level.`,
				},
				resource.Attribute{
					Name:        "array_type",
					Description: `(Optional) The type of the array elements if ` + "`" + `type` + "`" + ` is set to ` + "`" + `"array"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "array_enum",
					Description: `(Optional) Array of values that an array property's items can be set to.`,
				},
				resource.Attribute{
					Name:        "array_one_of",
					Description: `(Optional) Display name and value an enum array can be set to.`,
				},
				resource.Attribute{
					Name:        "const",
					Description: `(Required) value mapping to member of ` + "`" + `enum` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `(Required) display name for the enum value.`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `(Optional) Access control permissions for the property. It can be set to ` + "`" + `"READ_WRITE"` + "`" + `, ` + "`" + `"READ_ONLY"` + "`" + `, ` + "`" + `"HIDE"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "master",
					Description: `(Optional) Master priority for the user schema property. It can be set to ` + "`" + `"PROFILE_MASTER"` + "`" + ` or ` + "`" + `"OKTA"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "external_name",
					Description: `(Optional) External name of the user schema property. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "index",
					Description: `ID of the user schema property. ## Import User schema property can be imported via the property index. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_user_schema.example <index> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "index",
					Description: `ID of the user schema property. ## Import User schema property can be imported via the property index. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import okta_user_schema.example <index> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"okta_app_auto_login":            0,
		"okta_app_basic_auth":            1,
		"okta_app_bookmark":              2,
		"okta_app_group_assignment":      3,
		"okta_app_oauth":                 4,
		"okta_app_saml":                  5,
		"okta_app_secure_password_store": 6,
		"okta_app_swa":                   7,
		"okta_app_three_field":           8,
		"okta_app_user":                  9,
		"okta_app_user_base_schema":      10,
		"okta_app_user_schema":           11,
		"okta_auth_server":               12,
		"okta_auth_server_claim":         13,
		"okta_auth_server_policy":        14,
		"okta_auth_server_policy_rule":   15,
		"okta_auth_server_scope":         16,
		"okta_factor":                    17,
		"okta_group":                     18,
		"okta_group_roles":               19,
		"okta_group_rule":                20,
		"okta_idp_oidc":                  21,
		"okta_idp_saml":                  22,
		"okta_idp_saml_signing_key":      23,
		"okta_idp_social":                24,
		"okta_inline_hook":               25,
		"okta_network_zone":              26,
		"okta_policy_mfa":                27,
		"okta_policy_password":           28,
		"okta_policy_rule_idp_discovery": 29,
		"okta_policy_rule_mfa":           30,
		"okta_policy_rule_password":      31,
		"okta_policy_rule_signon":        32,
		"okta_policy_signon":             33,
		"okta_profile_mapping":           34,
		"okta_template_email":            35,
		"okta_trusted_origin":            36,
		"okta_user":                      37,
		"okta_user_base_schema":          38,
		"okta_user_schema":               39,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
