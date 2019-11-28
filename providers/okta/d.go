package okta

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "okta_app",
			Category:         "Data Sources",
			ShortDescription: `Get an application of any kind from Okta.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `` + "`" + `id` + "`" + ` of application.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `` + "`" + `label` + "`" + ` of application.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `` + "`" + `description` + "`" + ` of application.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `` + "`" + `name` + "`" + ` of application.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `` + "`" + `status` + "`" + ` of application.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_app_saml",
			Category:         "Data Sources",
			ShortDescription: `Get a SAML application from Okta.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `id of application.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `label of application.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `description of application.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `name of application.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `status of application.`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `Certificate key ID.`,
				},
				resource.Attribute{
					Name:        "auto_submit_toolbar",
					Description: `Display auto submit toolbar.`,
				},
				resource.Attribute{
					Name:        "hide_ios",
					Description: `Do not display application icon on mobile app.`,
				},
				resource.Attribute{
					Name:        "hide_web",
					Description: `Do not display application icon to users`,
				},
				resource.Attribute{
					Name:        "default_relay_state",
					Description: `Identifies a specific application resource in an IDP initiated SSO scenario.`,
				},
				resource.Attribute{
					Name:        "sso_url",
					Description: `Single Sign on Url.`,
				},
				resource.Attribute{
					Name:        "recipient",
					Description: `The location where the app may present the SAML assertion.`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `Identifies the location where the SAML response is intended to be sent inside of the SAML assertion.`,
				},
				resource.Attribute{
					Name:        "audience",
					Description: `Audience restriction.`,
				},
				resource.Attribute{
					Name:        "idp_issuer",
					Description: `SAML issuer ID.`,
				},
				resource.Attribute{
					Name:        "sp_issuer",
					Description: `SAML service provider issuer.`,
				},
				resource.Attribute{
					Name:        "subject_name_id_template",
					Description: `Template for app user's username when a user is assigned to the app.`,
				},
				resource.Attribute{
					Name:        "subject_name_id_format",
					Description: `Identifies the SAML processing rules.`,
				},
				resource.Attribute{
					Name:        "response_signed",
					Description: `Determines whether the SAML auth response message is digitally signed.`,
				},
				resource.Attribute{
					Name:        "request_compressed",
					Description: `Denotes whether the request is compressed or not.`,
				},
				resource.Attribute{
					Name:        "assertion_signed",
					Description: `Determines whether the SAML assertion is digitally signed.`,
				},
				resource.Attribute{
					Name:        "signature_algorithm",
					Description: `Signature algorithm used ot digitally sign the assertion and response.`,
				},
				resource.Attribute{
					Name:        "digest_algorithm",
					Description: `Determines the digest algorithm used to digitally sign the SAML assertion and response.`,
				},
				resource.Attribute{
					Name:        "honor_force_authn",
					Description: `Prompt user to re-authenticate if SP asks for it.`,
				},
				resource.Attribute{
					Name:        "authn_context_class_ref",
					Description: `Identifies the SAML authentication context class for the assertion’s authentication statement.`,
				},
				resource.Attribute{
					Name:        "accessibility_self_service",
					Description: `Enable self service.`,
				},
				resource.Attribute{
					Name:        "accessibility_error_redirect_url",
					Description: `Custom error page URL.`,
				},
				resource.Attribute{
					Name:        "accessibility_login_redirect_url",
					Description: `Custom login page URL.`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `features enabled.`,
				},
				resource.Attribute{
					Name:        "user_name_template",
					Description: `Username template.`,
				},
				resource.Attribute{
					Name:        "user_name_template_suffix",
					Description: `Username template suffix.`,
				},
				resource.Attribute{
					Name:        "user_name_template_type",
					Description: `Username template type.`,
				},
				resource.Attribute{
					Name:        "app_settings_json",
					Description: `Application settings in JSON format.`,
				},
				resource.Attribute{
					Name:        "attribute_statements",
					Description: `SAML Attribute statements.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_app_saml_metadata",
			Category:         "Data Sources",
			ShortDescription: `Get a SAML application's metadata from Okta.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "metadata",
					Description: `raw metadata of application.`,
				},
				resource.Attribute{
					Name:        "http_redirect_binding",
					Description: `urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect location from the SAML metadata.`,
				},
				resource.Attribute{
					Name:        "http_post_binding",
					Description: `urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Post location from the SAML metadata.`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `public certificate from application metadata.`,
				},
				resource.Attribute{
					Name:        "want_authn_requests_signed",
					Description: `Whether authn requests are signed.`,
				},
				resource.Attribute{
					Name:        "entity_id",
					Description: `Entity URL for instance ` + "`" + `https://www.okta.com/saml2/service-provider/sposcfdmlybtwkdcgtuf` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_auth_server",
			Category:         "Data Sources",
			ShortDescription: `Get an auth server from Okta.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Authorization server id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the auth server.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `description of Authorization server.`,
				},
				resource.Attribute{
					Name:        "audiences",
					Description: `array of audiences,`,
				},
				resource.Attribute{
					Name:        "kid",
					Description: `auth server key id.`,
				},
				resource.Attribute{
					Name:        "credentials_last_rotated",
					Description: `last time credentials were rotated.`,
				},
				resource.Attribute{
					Name:        "credentials_next_rotation",
					Description: `next time credentials will be rotated`,
				},
				resource.Attribute{
					Name:        "credentials_rotation_mode",
					Description: `mode of credential rotation, auto or manual.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `the activation status of the authorization server.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_default_policy",
			Category:         "Data Sources",
			ShortDescription: `Get a Default policy from Okta.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `id of policy.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `type of policy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_everyone_group",
			Category:         "Data Sources",
			ShortDescription: `Get the everyone group from Okta.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `the id of the group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_group",
			Category:         "Data Sources",
			ShortDescription: `Get a group from Okta.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `id of group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `name of group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `description of group.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `user ids that are members of this group, only included if ` + "`" + `include_users` + "`" + ` is set to ` + "`" + `true` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_idp_saml",
			Category:         "Data Sources",
			ShortDescription: `Get a SAML IdP from Okta.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `id of idp.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `name of the idp.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `type of idp.`,
				},
				resource.Attribute{
					Name:        "acs_binding",
					Description: `HTTP binding used to receive a SAMLResponse message from the IdP.`,
				},
				resource.Attribute{
					Name:        "acs_type",
					Description: `Determines whether to publish an instance-specific (trust) or organization (shared) ACS endpoint in the SAML metadata.`,
				},
				resource.Attribute{
					Name:        "sso_url",
					Description: `single sign on url.`,
				},
				resource.Attribute{
					Name:        "sso_binding",
					Description: `single sign on binding.`,
				},
				resource.Attribute{
					Name:        "sso_destination",
					Description: `SSO request binding, HTTP-POST or HTTP-REDIRECT.`,
				},
				resource.Attribute{
					Name:        "subject_format",
					Description: `Expression to generate or transform a unique username for the IdP user.`,
				},
				resource.Attribute{
					Name:        "subject_filter",
					Description: `regular expression pattern used to filter untrusted IdP usernames.`,
				},
				resource.Attribute{
					Name:        "issuer",
					Description: `URI that identifies the issuer (IdP).`,
				},
				resource.Attribute{
					Name:        "issuer_mode",
					Description: `indicates whether Okta uses the original Okta org domain URL, or a custom domain URL in the request to the IdP.`,
				},
				resource.Attribute{
					Name:        "audience",
					Description: `URI that identifies the target Okta IdP instance (SP)`,
				},
				resource.Attribute{
					Name:        "kid",
					Description: `Key ID reference to the IdP's X.509 signature certificate.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_idp_saml_metadata",
			Category:         "Data Sources",
			ShortDescription: `Get SAML IdP metadata from Okta.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "assertions_signed",
					Description: `whether assertions are signed.`,
				},
				resource.Attribute{
					Name:        "authn_request_signed",
					Description: `whether authn requests are signed.`,
				},
				resource.Attribute{
					Name:        "encryption_certificate",
					Description: `SAML request encryption certificate.`,
				},
				resource.Attribute{
					Name:        "entity_id",
					Description: `Entity URL for instance ` + "`" + `https://www.okta.com/saml2/service-provider/sposcfdmlybtwkdcgtuf` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "http_post_binding",
					Description: `urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Post location from the SAML metadata.`,
				},
				resource.Attribute{
					Name:        "http_redirect_binding",
					Description: `urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect location from the SAML metadata.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `raw IdP metadata.`,
				},
				resource.Attribute{
					Name:        "signing_certificate",
					Description: `SAML request signing certificate.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_policy",
			Category:         "Data Sources",
			ShortDescription: `Get a policy from Okta.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `id of policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `name of policy.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `type of policy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_user",
			Category:         "Data Sources",
			ShortDescription: `Get a single users from Okta.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "admin_roles",
					Description: `Administrator roles assigned to user.`,
				},
				resource.Attribute{
					Name:        "city",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "cost_center",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "country_code",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "custom_profile_attributes",
					Description: `raw JSON containing all custom profile attributes.`,
				},
				resource.Attribute{
					Name:        "department",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "division",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "employee_number",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "first_name",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "group_memberships",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "honorific_prefix",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "honorific_suffix",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "last_name",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "locale",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "login",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "manager",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "manager_id",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "middle_name",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "mobile_phone",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "nick_name",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "postal_address",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "preferred_language",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "primary_phone",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "profile_url",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "second_email",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "street_address",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "timezone",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "user_type",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "zip_code",
					Description: `user profile property.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_users",
			Category:         "Data Sources",
			ShortDescription: `Get a list of users from Okta.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "users",
					Description: `collection of users retrieved from Okta with the following properties.`,
				},
				resource.Attribute{
					Name:        "admin_roles",
					Description: `Administrator roles assigned to user.`,
				},
				resource.Attribute{
					Name:        "city",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "cost_center",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "country_code",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "custom_profile_attributes",
					Description: `raw JSON containing all custom profile attributes.`,
				},
				resource.Attribute{
					Name:        "department",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "division",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "employee_number",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "first_name",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "group_memberships",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "honorific_prefix",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "honorific_suffix",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "last_name",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "locale",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "login",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "manager",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "manager_id",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "middle_name",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "mobile_phone",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "nick_name",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "postal_address",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "preferred_language",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "primary_phone",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "profile_url",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "second_email",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "street_address",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "timezone",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "user_type",
					Description: `user profile property.`,
				},
				resource.Attribute{
					Name:        "zip_code",
					Description: `user profile property.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"okta_app":               0,
		"okta_app_saml":          1,
		"okta_app_saml_metadata": 2,
		"okta_auth_server":       3,
		"okta_default_policy":    4,
		"okta_everyone_group":    5,
		"okta_group":             6,
		"okta_idp_saml":          7,
		"okta_idp_saml_metadata": 8,
		"okta_policy":            9,
		"okta_user":              10,
		"okta_users":             11,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
