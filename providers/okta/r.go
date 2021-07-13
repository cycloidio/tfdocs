package okta

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "okta_admin_role_targets",
			Category:         "Resources",
			ShortDescription: `Manages targets for administrator roles.`,
			Description:      ``,
			Keywords: []string{
				"admin",
				"role",
				"targets",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
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
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
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
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
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
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
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
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
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
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_app_oauth_api_scope",
			Category:         "Resources",
			ShortDescription: `Manages API scopes for OAuth applications.`,
			Description:      ``,
			Keywords: []string{
				"app",
				"oauth",
				"api",
				"scope",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
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
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
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
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_app_shared_credentials",
			Category:         "Resources",
			ShortDescription: `Creates a SWA shared credentials app.`,
			Description:      ``,
			Keywords: []string{
				"app",
				"shared",
				"credentials",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
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
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_app_three_field",
			Category:         "Resources",
			ShortDescription: `Creates a Three Field Application.`,
			Description:      ``,
			Keywords: []string{
				"app",
				"three",
				"field",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
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
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
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
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
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
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
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
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
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
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_auth_server_claim_default",
			Category:         "Resources",
			ShortDescription: `Configures Default Authorization Server Claim`,
			Description:      ``,
			Keywords: []string{
				"auth",
				"server",
				"claim",
				"default",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
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
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
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
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
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
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_event_hook",
			Category:         "Resources",
			ShortDescription: `Creates an event hook.`,
			Description:      ``,
			Keywords: []string{
				"event",
				"hook",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
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
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_factor_totp",
			Category:         "Resources",
			ShortDescription: `Allows you to manage the time-based one-time password (TOTP) factors.`,
			Description:      ``,
			Keywords: []string{
				"factor",
				"totp",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
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
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_group_membership",
			Category:         "Resources",
			ShortDescription: `Manage an individual instance of group membership.`,
			Description:      ``,
			Keywords: []string{
				"group",
				"membership",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_group_role",
			Category:         "Resources",
			ShortDescription: `Assigns Admin roles to Okta Groups.`,
			Description:      ``,
			Keywords: []string{
				"group",
				"role",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
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
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
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
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
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
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
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
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
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
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
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
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
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
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
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
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
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
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_policy_mfa_default",
			Category:         "Resources",
			ShortDescription: `Configures default MFA Policy.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"mfa",
				"default",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
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
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_policy_password_default",
			Category:         "Resources",
			ShortDescription: `Configures default password policy.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"password",
				"default",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
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
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
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
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
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
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
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
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
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
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
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
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
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
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_template_sms",
			Category:         "Resources",
			ShortDescription: `Creates an Okta SMS Template.`,
			Description:      ``,
			Keywords: []string{
				"template",
				"sms",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
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
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
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
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
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
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_user_schema",
			Category:         "Resources",
			ShortDescription: `Creates a User Schema property.`,
			Description:      ``,
			Keywords: []string{
				"user",
				"schema",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_user_type",
			Category:         "Resources",
			ShortDescription: `Creates a User Type.`,
			Description:      ``,
			Keywords: []string{
				"user",
				"type",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"okta_admin_role_targets":        0,
		"okta_app_auto_login":            1,
		"okta_app_basic_auth":            2,
		"okta_app_bookmark":              3,
		"okta_app_group_assignment":      4,
		"okta_app_oauth":                 5,
		"okta_app_oauth_api_scope":       6,
		"okta_app_saml":                  7,
		"okta_app_secure_password_store": 8,
		"okta_app_shared_credentials":    9,
		"okta_app_swa":                   10,
		"okta_app_three_field":           11,
		"okta_app_user":                  12,
		"okta_app_user_base_schema":      13,
		"okta_app_user_schema":           14,
		"okta_auth_server":               15,
		"okta_auth_server_claim":         16,
		"okta_auth_server_claim_default": 17,
		"okta_auth_server_policy":        18,
		"okta_auth_server_policy_rule":   19,
		"okta_auth_server_scope":         20,
		"okta_event_hook":                21,
		"okta_factor":                    22,
		"okta_factor_totp":               23,
		"okta_group":                     24,
		"okta_group_membership":          25,
		"okta_group_role":                26,
		"okta_group_roles":               27,
		"okta_group_rule":                28,
		"okta_idp_oidc":                  29,
		"okta_idp_saml":                  30,
		"okta_idp_saml_signing_key":      31,
		"okta_idp_social":                32,
		"okta_inline_hook":               33,
		"okta_network_zone":              34,
		"okta_policy_mfa":                35,
		"okta_policy_mfa_default":        36,
		"okta_policy_password":           37,
		"okta_policy_password_default":   38,
		"okta_policy_rule_idp_discovery": 39,
		"okta_policy_rule_mfa":           40,
		"okta_policy_rule_password":      41,
		"okta_policy_rule_signon":        42,
		"okta_policy_signon":             43,
		"okta_profile_mapping":           44,
		"okta_template_email":            45,
		"okta_template_sms":              46,
		"okta_trusted_origin":            47,
		"okta_user":                      48,
		"okta_user_base_schema":          49,
		"okta_user_schema":               50,
		"okta_user_type":                 51,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
