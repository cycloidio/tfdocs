package auth0

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "auth0_client",
			Category:         "Resources",
			ShortDescription: `With this resource, you can create and configure applications that use Auth0 for authentication.`,
			Description:      ``,
			Keywords: []string{
				"client",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) String. Name of the client.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) String, (Max length = 140 characters). Description of the purpose of the client.`,
				},
				resource.Attribute{
					Name:        "client_secret_rotation_trigger",
					Description: `(Optional) Map.`,
				},
				resource.Attribute{
					Name:        "app_type",
					Description: `(Optional) String. Type of application the client represents. Options include ` + "`" + `native` + "`" + `, ` + "`" + `spa` + "`" + `, ` + "`" + `regular_web` + "`" + `, ` + "`" + `non_interactive` + "`" + `, ` + "`" + `rms` + "`" + `, ` + "`" + `box` + "`" + `, ` + "`" + `cloudbees` + "`" + `, ` + "`" + `concur` + "`" + `, ` + "`" + `dropbox` + "`" + `, ` + "`" + `mscrm` + "`" + `, ` + "`" + `echosign` + "`" + `, ` + "`" + `egnyte` + "`" + `, ` + "`" + `newrelic` + "`" + `, ` + "`" + `office365` + "`" + `, ` + "`" + `salesforce` + "`" + `, ` + "`" + `sentry` + "`" + `, ` + "`" + `sharepoint` + "`" + `, ` + "`" + `slack` + "`" + `, ` + "`" + `springcm` + "`" + `, ` + "`" + `zendesk` + "`" + `, ` + "`" + `zoom` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logo_uri",
					Description: `(Optional) String. URL of the logo for the client. Recommended size is 150px x 150px. If none is set, the default badge for the application type will be shown.`,
				},
				resource.Attribute{
					Name:        "is_first_party",
					Description: `(Optional) Boolean. Indicates whether or not this client is a first-party client.`,
				},
				resource.Attribute{
					Name:        "is_token_endpoint_ip_header_trusted",
					Description: `(Optional) Boolean. Indicates whether or not the token endpoint IP header is trusted.`,
				},
				resource.Attribute{
					Name:        "oidc_conformant",
					Description: `(Optional) Boolean. Indicates whether or not this client will conform to strict OIDC specifications.`,
				},
				resource.Attribute{
					Name:        "callbacks",
					Description: `(Optional) List(String). URLs that Auth0 may call back to after a user authenticates for the client. Make sure to specify the protocol (https://) otherwise the callback may fail in some cases. With the exception of custom URI schemes for native clients, all callbacks should use protocol https://.`,
				},
				resource.Attribute{
					Name:        "allowed_logout_urls",
					Description: `(Optional) List(String). URLs that Auth0 may redirect to after logout.`,
				},
				resource.Attribute{
					Name:        "grant_types",
					Description: `(Optional) List(String). Types of grants that this client is authorized to use.`,
				},
				resource.Attribute{
					Name:        "allowed_origins",
					Description: `(Optional) List(String). URLs that represent valid origins for cross-origin resource sharing. By default, all your callback URLs will be allowed.`,
				},
				resource.Attribute{
					Name:        "web_origins",
					Description: `(Optional) List(String). URLs that represent valid web origins for use with web message response mode.`,
				},
				resource.Attribute{
					Name:        "jwt_configuration",
					Description: `(Optional) List(Resource). Configuration settings for the JWTs issued for this client. For details, see [JWT Configuration](#jwt-configuration).`,
				},
				resource.Attribute{
					Name:        "encryption_key",
					Description: `(Optional) Map(String).`,
				},
				resource.Attribute{
					Name:        "sso",
					Description: `(Optional) Boolean. Indicates whether or not the client should use Auth0 rather than the IdP to perform Single Sign-On (SSO). True = Use Auth0.`,
				},
				resource.Attribute{
					Name:        "sso_disabled",
					Description: `(Optional) Boolean. Indicates whether or not SSO is disabled.`,
				},
				resource.Attribute{
					Name:        "cross_origin_auth",
					Description: `(Optional) Boolean. Indicates whether or not the client can be used to make cross-origin authentication requests.`,
				},
				resource.Attribute{
					Name:        "cross_origin_loc",
					Description: `(Optional) String. URL for the location on your site where the cross-origin verification takes place for the cross-origin auth flow. Used when performing auth in your own domain instead of through the Auth0-hosted login page.`,
				},
				resource.Attribute{
					Name:        "custom_login_page_on",
					Description: `(Optional) Boolean. Indicates whether or not a custom login page is to be used.`,
				},
				resource.Attribute{
					Name:        "custom_login_page",
					Description: `(Optional) String. Content of the custom login page.`,
				},
				resource.Attribute{
					Name:        "custom_login_page_preview",
					Description: `(Optional) String.`,
				},
				resource.Attribute{
					Name:        "form_template",
					Description: `(Optional) String. Form template for WS-Federation protocol.`,
				},
				resource.Attribute{
					Name:        "addons",
					Description: `(Optional) List(Resource). Configuration settings for add-ons for this client. For details, see [Add-ons](#add-ons).`,
				},
				resource.Attribute{
					Name:        "token_endpoint_auth_method",
					Description: `(Optional) String. Defines the requested authentication method for the token endpoint. Options include ` + "`" + `none` + "`" + ` (public client without a client secret), ` + "`" + `client_secret_post` + "`" + ` (client uses HTTP POST parameters), ` + "`" + `client_secret_basic` + "`" + ` (client uses HTTP Basic).`,
				},
				resource.Attribute{
					Name:        "client_metadata",
					Description: `(Optional) Map(String)`,
				},
				resource.Attribute{
					Name:        "mobile",
					Description: `(Optional) List(Resource). Configuration settings for mobile native applications. For details, see [Mobile](#mobile). ### JWT Configuration ` + "`" + `jwt_configuration` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "lifetime_in_seconds",
					Description: `(Optional) Integer. Number of seconds during which the JWT will be valid.`,
				},
				resource.Attribute{
					Name:        "secret_encoded",
					Description: `(Optional) Boolean. Indicates whether or not the client secret is base64 encoded.`,
				},
				resource.Attribute{
					Name:        "scopes",
					Description: `(Optional) Map(String). Permissions (scopes) included in JWTs.`,
				},
				resource.Attribute{
					Name:        "alg",
					Description: `(Optional) String. Algorithm used to sign JWTs. ### Add-ons ` + "`" + `addons` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "aws",
					Description: `(Optional) String`,
				},
				resource.Attribute{
					Name:        "azure_blob",
					Description: `(Optional) String`,
				},
				resource.Attribute{
					Name:        "azure_sb",
					Description: `(Optional) String`,
				},
				resource.Attribute{
					Name:        "box",
					Description: `(Optional) String`,
				},
				resource.Attribute{
					Name:        "cloudbees",
					Description: `(Optional) String`,
				},
				resource.Attribute{
					Name:        "concur",
					Description: `(Optional) String`,
				},
				resource.Attribute{
					Name:        "mscrm",
					Description: `(Optional) String`,
				},
				resource.Attribute{
					Name:        "rms",
					Description: `(Optional) String`,
				},
				resource.Attribute{
					Name:        "samlp",
					Description: `(Optional) List(Resource). Configuration settings for a SAML add-on. For details, see [SAML](#saml).`,
				},
				resource.Attribute{
					Name:        "sentry",
					Description: `(Optional) String`,
				},
				resource.Attribute{
					Name:        "slack",
					Description: `(Optional) String`,
				},
				resource.Attribute{
					Name:        "audience",
					Description: `(Optional) String. Audience of the SAML Assertion. Default will be the Issuer on SAMLRequest.`,
				},
				resource.Attribute{
					Name:        "recipient",
					Description: `(Optional) String. Recipient of the SAML Assertion (SubjectConfirmationData). Default is AssertionConsumerUrl on SAMLRequest or Callback URL if no SAMLRequest was sent.`,
				},
				resource.Attribute{
					Name:        "create_upn_claim",
					Description: `(Optional) Boolean, (Default=true) Indicates whether or not a UPN claim should be created.`,
				},
				resource.Attribute{
					Name:        "passthrough_claims_with_no_mapping",
					Description: `(Optional) Boolean, (Default=true). Indicates whether or not to passthrough claims that are not mapped to the common profile in the output assertion.`,
				},
				resource.Attribute{
					Name:        "map_unknown_claims_as_is",
					Description: `(Optional) Boolean, (Default=false). Indicates whether or not to add a prefix of ` + "`" + `http://schema.auth0.com` + "`" + ` to any claims that are not mapped to the common profile when passed through in the output assertion.`,
				},
				resource.Attribute{
					Name:        "map_identities",
					Description: `(Optional) Boolean, (Default=true). Indicates whether or not to add additional identity information in the token, such as the provider used and the access_token, if available.`,
				},
				resource.Attribute{
					Name:        "signature_algorithm",
					Description: `(Optional) String, (Default=` + "`" + `rsa-sha1` + "`" + `). Algorithm used to sign the SAML Assertion or response. Options include ` + "`" + `rsa-sha1` + "`" + ` and ` + "`" + `rsa-sha256` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "digest_algorithm",
					Description: `(Optional) String, (Default=` + "`" + `sha1` + "`" + `). Algorithm used to calculate the digest of the SAML Assertion or response. Options include ` + "`" + `defaultsha1` + "`" + ` and ` + "`" + `sha256` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `(Optional) String. Destination of the SAML Response. If not specified, it will be AssertionConsumerUrlof SAMLRequest or Callback URL if there was no SAMLRequest.`,
				},
				resource.Attribute{
					Name:        "lifetime_in_seconds",
					Description: `(Optional) Integer, (Default=3600). Number of seconds during which the token is valid.`,
				},
				resource.Attribute{
					Name:        "sign_response",
					Description: `(Optional) Boolean. Indicates whether or not the SAML Response should be signed instead of the SAML Assertion.`,
				},
				resource.Attribute{
					Name:        "typed_attributes",
					Description: `(Optional) Boolean, (Default=true). Indicates whether or not we should infer the ` + "`" + `xs:type` + "`" + ` of the element. Types include ` + "`" + `xs:string` + "`" + `, ` + "`" + `xs:boolean` + "`" + `, ` + "`" + `xs:double` + "`" + `, and ` + "`" + `xs:anyType` + "`" + `. When set to false, all ` + "`" + `xs:type` + "`" + ` are ` + "`" + `xs:anyType` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "include_attribute_name_format",
					Description: `(Optional) Boolean,(Default=true). Indicates whether or not we should infer the NameFormat based on the attribute name. If set to false, the attribute NameFormat is not set in the assertion.`,
				},
				resource.Attribute{
					Name:        "name_identifier_format",
					Description: `(Optional) String, (Default=` + "`" + `urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified` + "`" + `). Format of the name identifier.`,
				},
				resource.Attribute{
					Name:        "authn_context_class_ref",
					Description: `(Optional) String. Class reference of the authentication context.`,
				},
				resource.Attribute{
					Name:        "binding",
					Description: `(Optional) String. Protocol binding used for SAML logout responses.`,
				},
				resource.Attribute{
					Name:        "mappings",
					Description: `(Optional) Map(String). Mappings between the Auth0 user profile property name (` + "`" + `name` + "`" + `) and the output attributes on the SAML attribute in the assertion (` + "`" + `value` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "logout",
					Description: `(Optional) Map(Resource). Configuration settings for logout. For details, see [Logout](#logout).`,
				},
				resource.Attribute{
					Name:        "name_identifier_probes",
					Description: `(Optional) List(String). Attributes that can be used for Subject/NameID. Auth0 will try each of the attributes of this array in order and use the first value it finds. #### Logout ` + "`" + `logout` + "`" + ` supports the following options:`,
				},
				resource.Attribute{
					Name:        "callback",
					Description: `(Optional) String. Service provider's Single Logout Service URL, to which Auth0 will send logout requests and responses.`,
				},
				resource.Attribute{
					Name:        "slo_enabled",
					Description: `(Optional) Boolean. Indicates whether or not Auth0 should notify service providers of session termination. ### Mobile ` + "`" + `mobile` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "team_id",
					Description: `(Optional) String`,
				},
				resource.Attribute{
					Name:        "app_bundle_identifier",
					Description: `(Optional) String ## Attribute Reference Attributes exported by this resource include:`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `String. ID of the client.`,
				},
				resource.Attribute{
					Name:        "client_secret",
					Description: `String. Secret for the client; keep this private.`,
				},
				resource.Attribute{
					Name:        "is_first_party",
					Description: `Boolean. Indicates whether or not this client is a first-party client.`,
				},
				resource.Attribute{
					Name:        "is_token_endpoint_ip_header_trusted",
					Description: `Boolean`,
				},
				resource.Attribute{
					Name:        "oidc_conformant",
					Description: `Boolean. Indicates whether or not this client will conform to strict OIDC specifications.`,
				},
				resource.Attribute{
					Name:        "grant_types",
					Description: `List(String). Types of grants that this client is authorized to use.`,
				},
				resource.Attribute{
					Name:        "custom_login_page_on",
					Description: `Boolean. Indicates whether or not a custom login page is to be used.`,
				},
				resource.Attribute{
					Name:        "token_endpoint_auth_method",
					Description: `String. Defines the requested authentication method for the token endpoint. Options include ` + "`" + `none` + "`" + ` (public client without a client secret), ` + "`" + `client_secret_post` + "`" + ` (client uses HTTP POST parameters), ` + "`" + `client_secret_basic` + "`" + ` (client uses HTTP Basic).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "client_id",
					Description: `String. ID of the client.`,
				},
				resource.Attribute{
					Name:        "client_secret",
					Description: `String. Secret for the client; keep this private.`,
				},
				resource.Attribute{
					Name:        "is_first_party",
					Description: `Boolean. Indicates whether or not this client is a first-party client.`,
				},
				resource.Attribute{
					Name:        "is_token_endpoint_ip_header_trusted",
					Description: `Boolean`,
				},
				resource.Attribute{
					Name:        "oidc_conformant",
					Description: `Boolean. Indicates whether or not this client will conform to strict OIDC specifications.`,
				},
				resource.Attribute{
					Name:        "grant_types",
					Description: `List(String). Types of grants that this client is authorized to use.`,
				},
				resource.Attribute{
					Name:        "custom_login_page_on",
					Description: `Boolean. Indicates whether or not a custom login page is to be used.`,
				},
				resource.Attribute{
					Name:        "token_endpoint_auth_method",
					Description: `String. Defines the requested authentication method for the token endpoint. Options include ` + "`" + `none` + "`" + ` (public client without a client secret), ` + "`" + `client_secret_post` + "`" + ` (client uses HTTP POST parameters), ` + "`" + `client_secret_basic` + "`" + ` (client uses HTTP Basic).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "auth0_client_grant",
			Category:         "Resources",
			ShortDescription: `With this resource, you can create and manage client grants used with configured Auth0 clients.`,
			Description:      ``,
			Keywords: []string{
				"client",
				"grant",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "client_id",
					Description: `(Required) String. ID of the client for this grant.`,
				},
				resource.Attribute{
					Name:        "audience",
					Description: `(Required) String. Audience or API Identifier for this grant.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Required) List(String). Permissions (scopes) included in this grant.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "auth0_connection",
			Category:         "Resources",
			ShortDescription: `With this resource, you can configure and manage sources of users, which may include identity providers, databases, or passwordless authentication methods.`,
			Description:      ``,
			Keywords: []string{
				"connection",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the connection.`,
				},
				resource.Attribute{
					Name:        "is_domain_connection",
					Description: `(Optional) Indicates whether or not the connection is domain level.`,
				},
				resource.Attribute{
					Name:        "strategy",
					Description: `(Required) Type of the connection, which indicates the identity provider. Options include ` + "`" + `ad` + "`" + `, ` + "`" + `adfs` + "`" + `, ` + "`" + `amazon` + "`" + `, ` + "`" + `aol` + "`" + `, ` + "`" + `apple` + "`" + `, ` + "`" + `auth0` + "`" + `, ` + "`" + `auth0-adldap` + "`" + `, ` + "`" + `auth0-oidc` + "`" + `, ` + "`" + `baidu` + "`" + `, ` + "`" + `bitbucket` + "`" + `, ` + "`" + `bitly` + "`" + `, ` + "`" + `box` + "`" + `, ` + "`" + `custom` + "`" + `, ` + "`" + `daccount` + "`" + `, ` + "`" + `dropbox` + "`" + `, ` + "`" + `dwolla` + "`" + `, ` + "`" + `email` + "`" + `, ` + "`" + `evernote` + "`" + `, ` + "`" + `evernote-sandbox` + "`" + `, ` + "`" + `exact` + "`" + `, ` + "`" + `facebook` + "`" + `, ` + "`" + `fitbit` + "`" + `, ` + "`" + `flickr` + "`" + `, ` + "`" + `github` + "`" + `, ` + "`" + `google-apps` + "`" + `, ` + "`" + `google-oauth2` + "`" + `, ` + "`" + `guardian` + "`" + `, ` + "`" + `instagram` + "`" + `, ` + "`" + `ip` + "`" + `, ` + "`" + `line` + "`" + `, ` + "`" + `linkedin` + "`" + `, ` + "`" + `miicard` + "`" + `, ` + "`" + `oauth1` + "`" + `, ` + "`" + `oauth2` + "`" + `, ` + "`" + `office365` + "`" + `, ` + "`" + `oidc` + "`" + `, ` + "`" + `paypal` + "`" + `, ` + "`" + `paypal-sandbox` + "`" + `, ` + "`" + `pingfederate` + "`" + `, ` + "`" + `planningcenter` + "`" + `, ` + "`" + `renren` + "`" + `, ` + "`" + `salesforce` + "`" + `, ` + "`" + `salesforce-community` + "`" + `, ` + "`" + `salesforce-sandbox` + "`" + ` ` + "`" + `samlp` + "`" + `, ` + "`" + `sharepoint` + "`" + `, ` + "`" + `shopify` + "`" + `, ` + "`" + `sms` + "`" + `, ` + "`" + `soundcloud` + "`" + `, ` + "`" + `thecity` + "`" + `, ` + "`" + `thecity-sandbox` + "`" + `, ` + "`" + `thirtysevensignals` + "`" + `, ` + "`" + `twitter` + "`" + `, ` + "`" + `untappd` + "`" + `, ` + "`" + `vkontakte` + "`" + `, ` + "`" + `waad` + "`" + `, ` + "`" + `weibo` + "`" + `, ` + "`" + `windowslive` + "`" + `, ` + "`" + `wordpress` + "`" + `, ` + "`" + `yahoo` + "`" + `, ` + "`" + `yammer` + "`" + `, ` + "`" + `yandex` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "options",
					Description: `(Optional) Configuration settings for connection options. For details, see [Options](#options).`,
				},
				resource.Attribute{
					Name:        "enabled_clients",
					Description: `(Optional) IDs of the clients for which the connection is enabled. If not specified, no clients are enabled.`,
				},
				resource.Attribute{
					Name:        "realms",
					Description: `(Optional) Defines the realms for which the connection will be used (i.e., email domains). If not specified, the connection name is added as the realm. ### Options ` + "`" + `options` + "`" + ` supports different arguments depending on the connection ` + "`" + `strategy` + "`" + ` defined in [Argument Reference](#argument-reference). ### Auth0 With the ` + "`" + `auth0` + "`" + ` connection strategy, ` + "`" + `options` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "validation",
					Description: `(Optional) A map defining the validation options.`,
				},
				resource.Attribute{
					Name:        "password_policy",
					Description: `(Optional) Indicates level of password strength to enforce during authentication. A strong password policy will make it difficult, if not improbable, for someone to guess a password through either manual or automated means. Options include ` + "`" + `none` + "`" + `, ` + "`" + `low` + "`" + `, ` + "`" + `fair` + "`" + `, ` + "`" + `good` + "`" + `, ` + "`" + `excellent` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "password_history",
					Description: `(Optional) Configuration settings for the password history that is maintained for each user to prevent the reuse of passwords. For details, see [Password History](#password-history).`,
				},
				resource.Attribute{
					Name:        "password_no_personal_info",
					Description: `(Optional) Configuration settings for the password personal info check, which does not allow passwords that contain any part of the user's personal data, including user's name, username, nickname, user_metadata.name, user_metadata.first, user_metadata.last, user's email, or first part of the user's email. For details, see [Password No Personal Info](#password-no-personal-info).`,
				},
				resource.Attribute{
					Name:        "password_dictionary",
					Description: `(Optional) Configuration settings for the password dictionary check, which does not allow passwords that are part of the password dictionary. For details, see [Password Dictionary](#password-dictionary).`,
				},
				resource.Attribute{
					Name:        "password_complexity_options",
					Description: `(Optional) Configuration settings for password complexity. For details, see [Password Complexity Options](#password-complexity-options).`,
				},
				resource.Attribute{
					Name:        "api_enable_users",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "enabled_database_customization",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "brute_force_protection",
					Description: `(Optional) Indicates whether or not to enable brute force protection, which will limit the number of signups and failed logins from a suspicious IP address.`,
				},
				resource.Attribute{
					Name:        "import_mode",
					Description: `(Optional) Indicates whether or not you have a legacy user store and want to gradually migrate those users to the Auth0 user store. [Learn more](https://auth0.com/docs/users/guides/configure-automatic-migration).`,
				},
				resource.Attribute{
					Name:        "disable_signup",
					Description: `(Optional) Boolean. Indicates whether or not to allow user sign-ups to your application.`,
				},
				resource.Attribute{
					Name:        "requires_username",
					Description: `(Optional) Indicates whether or not the user is required to provide a username in addition to an email address.`,
				},
				resource.Attribute{
					Name:        "custom_scripts",
					Description: `(Optional) Custom database action scripts. For more information, read [Custom Database Action Script Templates](https://auth0.com/docs/connections/database/custom-db/templates).`,
				},
				resource.Attribute{
					Name:        "configuration",
					Description: `(Optional) A case-sensitive map of key value pairs used as configuration variables for the ` + "`" + `custom_script` + "`" + `. #### Password History ` + "`" + `password_history` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional) Indicates whether password history is enabled for the connection. When enabled, any existing users in this connection will be unaffected; the system will maintain their password history going forward.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) Indicates the number of passwords to keep in history with a maximum of 24. #### Password No Personal Info ` + "`" + `password_no_personal_info` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional) Indicates whether the password personal info check is enabled for this connection. #### Password Dictionary ` + "`" + `passsword_dictionary` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional) Indicates whether the password dictionary check is enabled for this connection.`,
				},
				resource.Attribute{
					Name:        "dictionary",
					Description: `(Optional) Customized contents of the password dictionary. By default, the password dictionary contains a list of the [10,000 most common passwords](https://github.com/danielmiessler/SecLists/blob/master/Passwords/Common-Credentials/10k-most-common.txt); your customized content is used in addition to the default password dictionary. Matching is not case-sensitive. #### Password Complexity Options ` + "`" + `password_complexity_options` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `(Optional) Facebook client ID.`,
				},
				resource.Attribute{
					Name:        "client_secret",
					Description: `(Optional) Facebook client secret.`,
				},
				resource.Attribute{
					Name:        "allowed_audiences",
					Description: `(Optional) List of allowed audiences.`,
				},
				resource.Attribute{
					Name:        "scopes",
					Description: `(Optional) Scopes.`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `(Optional) Facebook client ID.`,
				},
				resource.Attribute{
					Name:        "client_secret",
					Description: `(Optional) Facebook client secret.`,
				},
				resource.Attribute{
					Name:        "scopes",
					Description: `(Optional) Scopes.`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `(Optional) Apple client ID.`,
				},
				resource.Attribute{
					Name:        "client_secret",
					Description: `(Optional) App secret.`,
				},
				resource.Attribute{
					Name:        "team_id",
					Description: `(Optional) Team ID.`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `(Optional) Key ID.`,
				},
				resource.Attribute{
					Name:        "scopes",
					Description: `(Optional) Scopes.`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `(Optional) Linkedin API key.`,
				},
				resource.Attribute{
					Name:        "client_secret",
					Description: `(Optional) Linkedin secret key.`,
				},
				resource.Attribute{
					Name:        "strategy_version",
					Description: `(Optional) Version 1 is deprecated, use version 2.`,
				},
				resource.Attribute{
					Name:        "scopes",
					Description: `(Optional) Scopes.`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `(Optional) GitHub client ID.`,
				},
				resource.Attribute{
					Name:        "client_secret",
					Description: `(Optional) GitHub client secret.`,
				},
				resource.Attribute{
					Name:        "set_user_root_attributes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `(Optional) The Salesforce client ID.`,
				},
				resource.Attribute{
					Name:        "client_secret",
					Description: `(Optional) The Salesforce client secret.`,
				},
				resource.Attribute{
					Name:        "community_base_url",
					Description: `(Optional) String.`,
				},
				resource.Attribute{
					Name:        "scopes",
					Description: `(Optional) Scopes.`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `(Optional) OIDC provider client ID.`,
				},
				resource.Attribute{
					Name:        "client_secret",
					Description: `(Optional) OIDC provider client secret.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Value can be ` + "`" + `back_channel` + "`" + ` or ` + "`" + `front_channel` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scopes",
					Description: `(Optional) Scopes required by the connection. The value must be a list, for example ` + "`" + `["openid", "profile", "email"]` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "issuer",
					Description: `(Optional) Issuer URL. E.g. ` + "`" + `https://auth.example.com` + "`" + ``,
				},
				resource.Attribute{
					Name:        "discovery_url",
					Description: `(Optional) OpenID discovery URL. E.g. ` + "`" + `https://auth.example.com/.well-known/openid-configuration` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "jwks_uri",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "token_endpoint",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "userinfo_endpoint",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "authorization_endpoint",
					Description: `(Optional) ### Azure AD With the ` + "`" + `waad` + "`" + ` connection strategy, ` + "`" + `options` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "app_id",
					Description: `(Optional) Azure AD app ID.`,
				},
				resource.Attribute{
					Name:        "app_domain",
					Description: `(Optional) Azure AD domain name.`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `(Optional) Client ID for your Azure AD application.`,
				},
				resource.Attribute{
					Name:        "client_secret",
					Description: `(Optional) Client secret for your Azure AD application.`,
				},
				resource.Attribute{
					Name:        "domain_aliases",
					Description: `(Optional) List of the domains that can be authenticated using the Identity Provider. Only needed for Identifier First authentication flows.`,
				},
				resource.Attribute{
					Name:        "max_groups_to_retrieve",
					Description: `(Optional) Maximum number of groups to retrieve.`,
				},
				resource.Attribute{
					Name:        "tenant_domain",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "use_wsfed",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "waad_protocol",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "waad_common_endpoint",
					Description: `(Optional) Indicates whether or not to use the common endpoint rather than the default endpoint. Typically enabled if you're using this for a multi-tenant application in Azure AD. ### Twilio / SMS With the ` + "`" + `sms` + "`" + ` connection strategy, ` + "`" + `options` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "twilio_sid",
					Description: `(Optional) SID for your Twilio account.`,
				},
				resource.Attribute{
					Name:        "twilio_token",
					Description: `(Optional) AuthToken for your Twilio account.`,
				},
				resource.Attribute{
					Name:        "from",
					Description: `(Optional) SMS number for the sender. Used when SMS Source is From.`,
				},
				resource.Attribute{
					Name:        "syntax",
					Description: `(Optional) Syntax of the SMS. Options include ` + "`" + `markdown` + "`" + ` and ` + "`" + `liquid` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Optional) Template for the SMS. You can use ` + "`" + `@@password@@` + "`" + ` as a placeholder for the password value.`,
				},
				resource.Attribute{
					Name:        "totp",
					Description: `(Optional) Configuration options for one-time passwords. For details, see [TOTP](#totp).`,
				},
				resource.Attribute{
					Name:        "messaging_service_sid",
					Description: `(Optional) SID for Copilot. Used when SMS Source is Copilot. #### TOTP ` + "`" + `totp` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "time_step",
					Description: `(Optional) Integer. Seconds between allowed generation of new passwords.`,
				},
				resource.Attribute{
					Name:        "length",
					Description: `(Optional) Integer. Length of the one-time password.`,
				},
				resource.Attribute{
					Name:        "adfs_server",
					Description: `(Optional) ADFS Metadata source. ### SAML With the ` + "`" + `samlp` + "`" + ` connection strategy, ` + "`" + `options` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "debug",
					Description: `(Optional) (Boolean) When enabled additional debugging information will be generated.`,
				},
				resource.Attribute{
					Name:        "signing_cert",
					Description: `The X.509 signing certificate (encoded in PEM or CER) you retrieved from the IdP, Base64-encoded`,
				},
				resource.Attribute{
					Name:        "protocol_binding",
					Description: `(Optional) The SAML Response Binding - how the SAML token is received by Auth0 from IdP. Two possible values are ` + "`" + `urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect` + "`" + ` (default) and ` + "`" + `urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST` + "`" + ``,
				},
				resource.Attribute{
					Name:        "idpinitiated",
					Description: `(Optional) Configuration Options for IDP Initiated Authentication. This is an object with the properties: ` + "`" + `client_id` + "`" + `, ` + "`" + `client_protocol` + "`" + `, and ` + "`" + `client_authorizequery` + "`" + ``,
				},
				resource.Attribute{
					Name:        "tenant_domain",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "domain_aliases",
					Description: `(Optional) List of the domains that can be authenticated using the Identity Provider. Only needed for Identifier First authentication flows.`,
				},
				resource.Attribute{
					Name:        "sign_in_endpoint",
					Description: `SAML single login URL for the connection.`,
				},
				resource.Attribute{
					Name:        "sign_out_endpoint",
					Description: `(Optional) SAML single logout URL for the connection.`,
				},
				resource.Attribute{
					Name:        "fields_map",
					Description: `(Optional) SAML Attributes mapping. If you're configuring a SAML enterprise connection for a non-standard PingFederate Server, you must update the attribute mappings.`,
				},
				resource.Attribute{
					Name:        "sign_saml_request",
					Description: `(Optional) (Boolean) When enabled, the SAML authentication request will be signed.`,
				},
				resource.Attribute{
					Name:        "signature_algorithm",
					Description: `(Optional) Sign Request Algorithm`,
				},
				resource.Attribute{
					Name:        "digest_algorithm",
					Description: `(Optional) Sign Request Algorithm Digest`,
				},
				resource.Attribute{
					Name:        "is_domain_connection",
					Description: `Boolean. Indicates whether or not the connection is domain level.`,
				},
				resource.Attribute{
					Name:        "options",
					Description: `List(Resource). Configuration settings for connection options. For details, see [Options Attributes](#options-attributes).`,
				},
				resource.Attribute{
					Name:        "realms",
					Description: `List(String). Defines the realms for which the connection will be used (i.e., email domains). If the array is empty or the property is not specified, the connection name is added as the realm. ### Options Attributes ` + "`" + `options` + "`" + ` exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "password_history",
					Description: `List(Resource). Configuration settings for the password history that is maintained for each user to prevent the reuse of passwords. For details, see [Password History Attributes](#password-attributes-history). #### Password History Attributes ` + "`" + `password_history` + "`" + ` exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `Boolean. Indicates whether password history is enabled for the connection. When enabled, any existing users in this connection will be unaffected; the system will maintain their password history going forward.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Integer. Indicates the number of passwords to keep in history. ### Import Connections can be imported using their id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import auth0_connection.google con_a17f21fdb24d48a0 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "is_domain_connection",
					Description: `Boolean. Indicates whether or not the connection is domain level.`,
				},
				resource.Attribute{
					Name:        "options",
					Description: `List(Resource). Configuration settings for connection options. For details, see [Options Attributes](#options-attributes).`,
				},
				resource.Attribute{
					Name:        "realms",
					Description: `List(String). Defines the realms for which the connection will be used (i.e., email domains). If the array is empty or the property is not specified, the connection name is added as the realm. ### Options Attributes ` + "`" + `options` + "`" + ` exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "password_history",
					Description: `List(Resource). Configuration settings for the password history that is maintained for each user to prevent the reuse of passwords. For details, see [Password History Attributes](#password-attributes-history). #### Password History Attributes ` + "`" + `password_history` + "`" + ` exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `Boolean. Indicates whether password history is enabled for the connection. When enabled, any existing users in this connection will be unaffected; the system will maintain their password history going forward.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Integer. Indicates the number of passwords to keep in history. ### Import Connections can be imported using their id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import auth0_connection.google con_a17f21fdb24d48a0 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "auth0_custom_domain",
			Category:         "Resources",
			ShortDescription: `With this resource, you can create and manage a custom domain within your Auth0 tenant in order to maintain a consistent user experience.`,
			Description:      ``,
			Keywords: []string{
				"custom",
				"domain",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) String. Name of the custom domain.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) String. Provisioning type for the custom domain. Options include ` + "`" + `auth0_managed_certs` + "`" + ` and ` + "`" + `self_managed_certs` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "verification_method",
					Description: `(Required) String. Domain verification method. Options include ` + "`" + `txt` + "`" + `. ## Attribute Reference Attributes exported by this resource include:`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `Boolean. Indicates whether or not this is a primary domain.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `String. Configuration status for the custom domain. Options include ` + "`" + `disabled` + "`" + `, ` + "`" + `pending` + "`" + `, ` + "`" + `pending_verification` + "`" + `, and ` + "`" + `ready` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "verification",
					Description: `List(Resource). Configuration settings for verification. For details, see [Verification](#verification). ### Verification ` + "`" + `verification` + "`" + ` exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "methods",
					Description: `List(Map). Verification methods for the domain.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "primary",
					Description: `Boolean. Indicates whether or not this is a primary domain.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `String. Configuration status for the custom domain. Options include ` + "`" + `disabled` + "`" + `, ` + "`" + `pending` + "`" + `, ` + "`" + `pending_verification` + "`" + `, and ` + "`" + `ready` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "verification",
					Description: `List(Resource). Configuration settings for verification. For details, see [Verification](#verification). ### Verification ` + "`" + `verification` + "`" + ` exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "methods",
					Description: `List(Map). Verification methods for the domain.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "auth0_email",
			Category:         "Resources",
			ShortDescription: `With this resource, you can configure email providers so you can route all emails that are part of Auth0's authentication workflows through the supported high-volume email service of your choice.`,
			Description:      ``,
			Keywords: []string{
				"email",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) String. Name of the email provider. Options include ` + "`" + `mailgun` + "`" + `, ` + "`" + `mandrill` + "`" + `, ` + "`" + `sendgrid` + "`" + `, ` + "`" + `ses` + "`" + `, ` + "`" + `smtp` + "`" + `, and ` + "`" + `sparkpost` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Boolean. Indicates whether or not the email provider is enabled.`,
				},
				resource.Attribute{
					Name:        "default_from_address",
					Description: `(Required) String. Email address to use as the sender when no other "from" address is specified.`,
				},
				resource.Attribute{
					Name:        "credentials",
					Description: `(Required) List(Resource). Configuration settings for the credentials for the email provider. For details, see [Credentials](#credentials). ### Credentials ` + "`" + `credentials` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "api_user",
					Description: `(Optional) String. API User for your email service.`,
				},
				resource.Attribute{
					Name:        "api_key",
					Description: `(Optional) String, Case-sensitive. API Key for your email service. Will always be encrypted in our database.`,
				},
				resource.Attribute{
					Name:        "access_key_id",
					Description: `(Optional) String, Case-sensitive. AWS Access Key ID. Used only for AWS.`,
				},
				resource.Attribute{
					Name:        "secret_access_key",
					Description: `(Optional) String, Case-sensitive. AWS Secret Key. Will always be encrypted in our database. Used only for AWS.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) String. Default region. Used only for AWS, Mailgun, and SparkPost.`,
				},
				resource.Attribute{
					Name:        "smtp_host",
					Description: `(Optional) String. Hostname or IP address of your SMTP server. Used only for SMTP.`,
				},
				resource.Attribute{
					Name:        "smtp_port",
					Description: `(Optional) Integer. Port used by your SMTP server. Please avoid using port 25 if possible because many providers have limitations on this port. Used only for SMTP.`,
				},
				resource.Attribute{
					Name:        "smtp_user",
					Description: `(Optional) String. SMTP username. Used only for SMTP.`,
				},
				resource.Attribute{
					Name:        "smtp_pass",
					Description: `(Optional) String, Case-sensitive. SMTP password. Used only for SMTP.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "auth0_email_template",
			Category:         "Resources",
			ShortDescription: `With this resource, you can configure email templates to customize the look, feel, and sender identities of emails sent by Auth0 using configured email providers.`,
			Description:      ``,
			Keywords: []string{
				"email",
				"template",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `(Required) String. Template name. Options include ` + "`" + `verify_email` + "`" + `, ` + "`" + `reset_email` + "`" + `, ` + "`" + `welcome_email` + "`" + `, ` + "`" + `blocked_account` + "`" + `, ` + "`" + `stolen_credentials` + "`" + `, ` + "`" + `enrollment_email` + "`" + `, ` + "`" + `mfa_oob_code` + "`" + `, ` + "`" + `change_password` + "`" + ` (legacy), and ` + "`" + `password_reset` + "`" + ` (legacy).`,
				},
				resource.Attribute{
					Name:        "body",
					Description: `(Required) String. Body of the email template. You can include [common variables](https://auth0.com/docs/email/templates#common-variables).`,
				},
				resource.Attribute{
					Name:        "from",
					Description: `(Required) String. Email address to use as the sender. You can include [common variables](https://auth0.com/docs/email/templates#common-variables).`,
				},
				resource.Attribute{
					Name:        "result_url",
					Description: `(Required) String. URL to redirect the user to after a successful action. [Learn more](https://auth0.com/docs/email/templates#configuring-the-redirect-to-url).`,
				},
				resource.Attribute{
					Name:        "subject",
					Description: `(Required) String. Subject line of the email. You can include [common variables](https://auth0.com/docs/email/templates#common-variables).`,
				},
				resource.Attribute{
					Name:        "syntax",
					Description: `(Required) String. Syntax of the template body. You can use either text or HTML + Liquid syntax.`,
				},
				resource.Attribute{
					Name:        "url_lifetime_in_seconds",
					Description: `(Optional) Integer. Number of seconds during which the link within the email will be valid.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Boolean. Indicates whether or not the template is enabled.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "auth0_hook",
			Category:         "Resources",
			ShortDescription: `Hooks are secure, self-contained functions that allow you to customize the behavior of Auth0 when executed for selected extensibility points of the Auth0 platform. Auth0 invokes Hooks during runtime to execute your custom Node.js code. Depending on the extensibility point, you can use Hooks with Database Connections and/or Passwordless Connections.`,
			Description:      ``,
			Keywords: []string{
				"hook",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Whether the hook is enabled, or disabled`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of this hook`,
				},
				resource.Attribute{
					Name:        "script",
					Description: `(Required) Code to be executed when this hook runs`,
				},
				resource.Attribute{
					Name:        "trigger_id",
					Description: `(Required) Execution stage of this rule. Can be credentials-exchange, pre-user-registration, post-user-registration, post-change-password, or send-phone-message`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "auth0_prompt",
			Category:         "Resources",
			ShortDescription: `With this resource, you can manage your Auth0 prompts, including choosing the login experience version.`,
			Description:      ``,
			Keywords: []string{
				"prompt",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "auth0_resource_server",
			Category:         "Resources",
			ShortDescription: `With this resource, you can set up APIs that can be consumed from your authorized applications.`,
			Description:      ``,
			Keywords: []string{
				"resource",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) String. Friendly name for the resource server. Cannot include ` + "`" + `<` + "`" + ` or ` + "`" + `>` + "`" + ` characters.`,
				},
				resource.Attribute{
					Name:        "identifier",
					Description: `(Optional) String. Unique identifier for the resource server. Used as the audience parameter for authorization calls. Can not be changed once set.`,
				},
				resource.Attribute{
					Name:        "scopes",
					Description: `(Optional) Set(Resource). List of permissions (scopes) used by this resource server. For details, see [Scopes](#scopes).`,
				},
				resource.Attribute{
					Name:        "signing_alg",
					Description: `(Optional) String. Algorithm used to sign JWTs. Options include ` + "`" + `HS256` + "`" + ` and ` + "`" + `RS256` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "signing_secret",
					Description: `(Optional) String. Secret used to sign tokens when using symmetric algorithms (HS256).`,
				},
				resource.Attribute{
					Name:        "allow_offline_access",
					Description: `(Optional) Boolean. Indicates whether or not refresh tokens can be issued for this resource server.`,
				},
				resource.Attribute{
					Name:        "token_lifetime",
					Description: `(Optional) Integer. Number of seconds during which access tokens issued for this resource server from the token endpoint remain valid.`,
				},
				resource.Attribute{
					Name:        "token_lifetime_for_web",
					Description: `(Optional) Integer. Number of seconds during which access tokens issued for this resource server via implicit or hybrid flows remain valid. Cannot be greater than the ` + "`" + `token_lifetime` + "`" + ` value.`,
				},
				resource.Attribute{
					Name:        "skip_consent_for_verifiable_first_party_clients",
					Description: `(Optional) Boolean. Indicates whether or not to skip user consent for applications flagged as first party.`,
				},
				resource.Attribute{
					Name:        "verification_location",
					Description: `(Optional) String`,
				},
				resource.Attribute{
					Name:        "options",
					Description: `(Optional) Map(String). Used to store additional metadata`,
				},
				resource.Attribute{
					Name:        "enforce_policies",
					Description: `(Optional) Boolean. Indicates whether or not authorization polices are enforced.`,
				},
				resource.Attribute{
					Name:        "token_dialect",
					Description: `(Optional) String. Dialect of access tokens that should be issued for this resource server. Options include ` + "`" + `access_token` + "`" + ` or ` + "`" + `access_token_authz` + "`" + ` (includes permissions). ### Scopes ` + "`" + `scopes` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) String. Name of the permission (scope). Examples include ` + "`" + `read:appointments` + "`" + ` or ` + "`" + `delete:appointments` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) String. Description of the permission (scope). ## Attribute Reference Attributes exported by this resource include:`,
				},
				resource.Attribute{
					Name:        "signing_alg",
					Description: `String. Algorithm used to sign JWTs. Options include ` + "`" + `HS256` + "`" + ` and ` + "`" + `RS256` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "signing_secret",
					Description: `String. Secret used to sign tokens when using symmetric algorithms (HS256).`,
				},
				resource.Attribute{
					Name:        "token_lifetime",
					Description: `Integer. Number of seconds during which access tokens issued for this resource server from the token endpoint remain valid.`,
				},
				resource.Attribute{
					Name:        "token_lifetime_for_web",
					Description: `Integer. Number of seconds during which access tokens issued for this resource server via implicit or hybrid flows remain valid. Cannot be greater than the ` + "`" + `token_lifetime` + "`" + ` value.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "signing_alg",
					Description: `String. Algorithm used to sign JWTs. Options include ` + "`" + `HS256` + "`" + ` and ` + "`" + `RS256` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "signing_secret",
					Description: `String. Secret used to sign tokens when using symmetric algorithms (HS256).`,
				},
				resource.Attribute{
					Name:        "token_lifetime",
					Description: `Integer. Number of seconds during which access tokens issued for this resource server from the token endpoint remain valid.`,
				},
				resource.Attribute{
					Name:        "token_lifetime_for_web",
					Description: `Integer. Number of seconds during which access tokens issued for this resource server via implicit or hybrid flows remain valid. Cannot be greater than the ` + "`" + `token_lifetime` + "`" + ` value.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "auth0_role",
			Category:         "Resources",
			ShortDescription: `With this resource, you can created and manage collections of permissions that can be assigned to users, which are otherwise known as roles.`,
			Description:      ``,
			Keywords: []string{
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "role_id",
					Description: `(Optional) String. ID for this role.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) String. Name for this role.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) String. Description of the role.`,
				},
				resource.Attribute{
					Name:        "user_ids",
					Description: `(Optional) List(String). IDs of the users to which the role is assigned.`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `(Optional) Set(Resource). Configuration settings for permissions (scopes) attached to the role. For details, see [Permissions](#permissions). ### Permissions ` + "`" + `permissions` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) String. Name of the permission (scope).`,
				},
				resource.Attribute{
					Name:        "resource_server_identifier",
					Description: `(Required) String. Unique identifier for the resource server. ## Attribute Reference Attributes exported by this resource include:`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `String. ID for the role.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "role_id",
					Description: `String. ID for the role.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "auth0_rule",
			Category:         "Resources",
			ShortDescription: `With this resource, you can create and manage rules, which are custom Javascript snippets that run in a secure, isolate sandbox as part of your authentication pipeline.`,
			Description:      ``,
			Keywords: []string{
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) String. Name of the rule. May only contain alphanumeric characters, spaces, and hyphens. May neither start nor end with hyphens or spaces.`,
				},
				resource.Attribute{
					Name:        "script",
					Description: `(Required) String. Code to be executed when the rule runs.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Optional) Integer. Order in which the rule executes relative to other rules. Lower-valued rules execute first.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Boolean. Indicates whether the rule is enabled. ## Attribute Reference Attributes exported by this resource include:`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `Integer. Order in which the rule executes relative to other rules. Lower-valued rules execute first.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "order",
					Description: `Integer. Order in which the rule executes relative to other rules. Lower-valued rules execute first.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "auth0_rule_config",
			Category:         "Resources",
			ShortDescription: `With this resource, you can create and manage variables for rules, which are custom Javascript snippets that run in a secure, isolate sandbox as part of your authentication pipeline.`,
			Description:      ``,
			Keywords: []string{
				"rule",
				"config",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) String. Key for a rules configuration variable.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) String, Case-sensitive. Value for a rules configuration variable.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "auth0_tenant",
			Category:         "Resources",
			ShortDescription: `With this resource, you can manage Auth0 tenants, including setting logos and support contact information, setting error pages, and configuring default tenant behaviors.`,
			Description:      ``,
			Keywords: []string{
				"tenant",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "change_password",
					Description: `(Optional) List(Resource). Configuration settings for change passsword page. For details, see [Change Password Page](#change-password-page).`,
				},
				resource.Attribute{
					Name:        "guardian_mfa_page",
					Description: `(Optional) List(Resource). Configuration settings for the Guardian MFA page. For details, see [Guardian MFA Page](#guardian-mfa-page).`,
				},
				resource.Attribute{
					Name:        "default_audience",
					Description: `(Optional) String. API Audience to use by default for API Authorization flows. This setting is equivalent to appending the audience to every authorization request made to the tenant for every application.`,
				},
				resource.Attribute{
					Name:        "default_directory",
					Description: `(Optional) String. Name of the connection to be used for Password Grant exchanges. Options include ` + "`" + `auth0-adldap` + "`" + `, ` + "`" + `ad` + "`" + `, ` + "`" + `auth0` + "`" + `, ` + "`" + `email` + "`" + `, ` + "`" + `sms` + "`" + `, ` + "`" + `waad` + "`" + `, and ` + "`" + `adfs` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "default_redirection_uri",
					Description: `(Optional) String. The default absolute redirection uri, must be https and cannot contain a fragment.`,
				},
				resource.Attribute{
					Name:        "error_page",
					Description: `(Optional) List(Resource). Configuration settings for error pages. For details, see [Error Page](#error-page).`,
				},
				resource.Attribute{
					Name:        "friendly_name",
					Description: `(Optional) String. Friendly name for the tenant.`,
				},
				resource.Attribute{
					Name:        "picture_url",
					Description: `(Optional). String URL of logo to be shown for the tenant. Recommended size is 150px x 150px. If no URL is provided, the Auth0 logo will be used.`,
				},
				resource.Attribute{
					Name:        "support_email",
					Description: `(Optional) String. Support email address for authenticating users.`,
				},
				resource.Attribute{
					Name:        "support_url",
					Description: `(Optional) String. Support URL for authenticating users.`,
				},
				resource.Attribute{
					Name:        "allowed_logout_urls",
					Description: `(Optional) List(String). URLs that Auth0 may redirect to after logout.`,
				},
				resource.Attribute{
					Name:        "session_lifetime",
					Description: `(Optional) Integer. Number of hours during which a session will stay valid.`,
				},
				resource.Attribute{
					Name:        "sandbox_version",
					Description: `(Optional) String. Selected sandbox version for the extensibility environment, which allows you to use custom scripts to extend parts of Auth0's functionality.`,
				},
				resource.Attribute{
					Name:        "idle_session_lifetime",
					Description: `(Optional) Integer. Number of hours during which a session can be inactive before the user must log in again.`,
				},
				resource.Attribute{
					Name:        "flags",
					Description: `(Optional) List(Resource). Configuration settings for tenant flags. For details, see [Flags](#flags).`,
				},
				resource.Attribute{
					Name:        "universal_login",
					Description: `(Optional) List(Resource). Configuration settings for Universal Login. For details, see [Universal Login](#universal-login). ### Change Password Page ` + "`" + `change_password_page` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Boolean. Indicates whether or not to use the custom change password page.`,
				},
				resource.Attribute{
					Name:        "html",
					Description: `(Required) String, HTML format with supported Liquid syntax. Customized content of the change password page. ### Guardian MFA Page ` + "`" + `guardian_mfa_page` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Boolean. Indicates whether or not to use the custom Guardian page.`,
				},
				resource.Attribute{
					Name:        "html",
					Description: `(Required) String, HTML format with supported Liquid syntax. Customized content of the Guardian page. ### Error Page ` + "`" + `error_page` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "html",
					Description: `(Required) String, HTML format with supported Liquid syntax. Customized content of the error page.`,
				},
				resource.Attribute{
					Name:        "show_log_link",
					Description: `(Required) Boolean. Indicates whether or not to show the link to logs as part of the default error page.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) String. URL to redirect to when an error occurs rather than showing the default error page. ### Flags ` + "`" + `flags` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "change_pwd_flow_v1",
					Description: `(Optional) Boolean. Indicates whether or not to use the older v1 change password flow. Not recommended except for backward compatibility.`,
				},
				resource.Attribute{
					Name:        "enable_client_connections",
					Description: `(Optional) Boolean. Indicates whether or not all current connections should be enabled when a new client is created.`,
				},
				resource.Attribute{
					Name:        "enable_apis_section",
					Description: `(Optional) Boolean. Indicates whether or not the APIs section is enabled for the tenant.`,
				},
				resource.Attribute{
					Name:        "enable_pipeline2",
					Description: `(Optional) Boolean. Indicates whether or not advanced API Authorization scenarios are enabled.`,
				},
				resource.Attribute{
					Name:        "enable_dynamic_client_registration",
					Description: `(Optional) Boolean. Indicates whether or not the tenant allows dynamic client registration.`,
				},
				resource.Attribute{
					Name:        "enable_custom_domain_in_emails",
					Description: `(Optional) Boolean. Indicates whether or not the tenant allows custom domains in emails.`,
				},
				resource.Attribute{
					Name:        "universal_login",
					Description: `(Optional) Boolean. Indicates whether or not the tenant uses universal login.`,
				},
				resource.Attribute{
					Name:        "enable_legacy_logs_search_v2",
					Description: `(Optional) Boolean. Indicates whether or not to use the older v2 legacy logs search.`,
				},
				resource.Attribute{
					Name:        "disable_clickjack_protection_headers",
					Description: `(Optional) Boolean. Indicated whether or not classic Universal Login prompts include additional security headers to prevent clickjacking.`,
				},
				resource.Attribute{
					Name:        "enable_public_signup_user_exists_error",
					Description: `(Optional) Boolean. Indicates whether or not the public sign up process shows a user_exists error if the user already exists. ### Universal Login ` + "`" + `universal_login` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "colors",
					Description: `(Optional) List(Resource). Configuration settings for Universal Login colors. See [Universal Login - Colors](#colors). #### Colors ` + "`" + `colors` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `(Optional) String, Hexadecimal. Primary button background color.`,
				},
				resource.Attribute{
					Name:        "page_background",
					Description: `(Optional) String, Hexadecimal. Background color of login pages. ## Attribute Reference Attributes exported by this resource include:`,
				},
				resource.Attribute{
					Name:        "sandbox_version",
					Description: `String. Selected sandbox version for the extensibility environment, which allows you to use custom scripts to extend parts of Auth0's functionality.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "sandbox_version",
					Description: `String. Selected sandbox version for the extensibility environment, which allows you to use custom scripts to extend parts of Auth0's functionality.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "auth0_user",
			Category:         "Resources",
			ShortDescription: `With this resource, you can manage user identities, including resetting passwords, and creating, provisioning, blocking, and deleting users.`,
			Description:      ``,
			Keywords: []string{
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "user_id",
					Description: `(Optional) String. ID of the user.`,
				},
				resource.Attribute{
					Name:        "connection_name",
					Description: `(Required) String. Name of the connection from which the user information was sourced.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) String. Username of the user. Only valid if the connection requires a username.`,
				},
				resource.Attribute{
					Name:        "nickname",
					Description: `(Optional) String. Preferred nickname or alias of the user.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) String, Case-sensitive. Initial password for this user. Used for non-SMS connections.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) String. Email address of the user.`,
				},
				resource.Attribute{
					Name:        "email_verified",
					Description: `(Optional) Boolean. Indicates whether or not the email address has been verified.`,
				},
				resource.Attribute{
					Name:        "verify_email",
					Description: `(Optional) Boolean. Indicates whether or not the user will receive a verification email after creation. Overrides behavior of ` + "`" + `email_verified` + "`" + ` parameter.`,
				},
				resource.Attribute{
					Name:        "phone_number",
					Description: `(Optional) String. Phone number for the user; follows the E.164 recommendation. Used for SMS connections.`,
				},
				resource.Attribute{
					Name:        "phone_verified",
					Description: `(Optional) Boolean. Indicates whether or not the phone number has been verified.`,
				},
				resource.Attribute{
					Name:        "user_metadata",
					Description: `(Optional) String, JSON format. Custom fields that store info about the user that does not impact a user's core functionality. Examples include work address, home address, and user preferences.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `(Optional) Set(String). Set of IDs of roles assigned to the user.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"auth0_client":          0,
		"auth0_client_grant":    1,
		"auth0_connection":      2,
		"auth0_custom_domain":   3,
		"auth0_email":           4,
		"auth0_email_template":  5,
		"auth0_hook":            6,
		"auth0_prompt":          7,
		"auth0_resource_server": 8,
		"auth0_role":            9,
		"auth0_rule":            10,
		"auth0_rule_config":     11,
		"auth0_tenant":          12,
		"auth0_user":            13,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
