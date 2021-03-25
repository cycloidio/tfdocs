package onelogin

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "onelogin_onelogin_app",
			Category:         "Resources",
			ShortDescription: `Creates a Basic Application.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The app's name.`,
				},
				resource.Attribute{
					Name:        "connector_id",
					Description: `(Required) The ID for the app connector, dictates the type of app (e.g. AWS Multi-Role App).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) App description.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) Notes about the app.`,
				},
				resource.Attribute{
					Name:        "visible",
					Description: `(Optional) Determine if app should be visible in OneLogin portal. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "allow_assumed_signin",
					Description: `(Optional) Enable sign in when user has been assumed by the account owner. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "provisioning",
					Description: `(Optional) Settings regarding the app's provisioning ability.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Indicates if provisioning is enabled for this app.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `(Optional) a list of custom parameters for this app.`,
				},
				resource.Attribute{
					Name:        "param_key_name",
					Description: `(Required) Name to represent the parameter in OneLogin.`,
				},
				resource.Attribute{
					Name:        "safe_entitlements_enabled",
					Description: `(Optional) Indicates that the parameter is used to support creating entitlements using OneLogin Mappings. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user_attribute_mappings",
					Description: `(Optional) A user attribute to map values from. For custom attributes prefix the name of the attribute with ` + "`" + `custom_attribute_` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "provisioned_entitlements",
					Description: `(Optional) Provisioned access entitlements for the app. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "skip_if_blank",
					Description: `(Optional) Flag to let the SCIM provisioner know not include this value if it's blank. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user_attribute_macros",
					Description: `(Optional) When ` + "`" + `user_attribute_mappings` + "`" + ` is set to ` + "`" + `_macro_` + "`" + ` this macro will be used to assign the parameter value.`,
				},
				resource.Attribute{
					Name:        "attributes_transformations",
					Description: `(Optional) Describes how the app's attributes should be transformed.`,
				},
				resource.Attribute{
					Name:        "default_values",
					Description: `(Optional) Default Parameter values.`,
				},
				resource.Attribute{
					Name:        "include_in_saml_assertion",
					Description: `(Optional) When true, this parameter will be included in a SAML assertion payload.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Optional) The can only be set when creating a new parameter. It can not be updated.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Optional) Parameter values. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `App's unique ID in OneLogin.`,
				},
				resource.Attribute{
					Name:        "allow_assumed_signin",
					Description: `App sign in allowed when user assumed by account administrator.`,
				},
				resource.Attribute{
					Name:        "auth_method",
					Description: `The apps auth method. Refer to the [OneLogin Apps Documentation](https://developers.onelogin.com/api-docs/2/apps/app-resource) for a comprehensive list of available auth methods.`,
				},
				resource.Attribute{
					Name:        "connector_id",
					Description: `ID of the apps underlying connector. Dictates the type of app (e.g. AWS Multi-Role App).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `App description.`,
				},
				resource.Attribute{
					Name:        "icon_url",
					Description: `The url for the app's icon.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The app's name.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `Notes about the app.`,
				},
				resource.Attribute{
					Name:        "tab_id",
					Description: `The tab in which to display in OneLogin portal.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Timestamp for app's last update.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Timestamp for app's creation.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `The security policy assigned to the app.`,
				},
				resource.Attribute{
					Name:        "visible",
					Description: `Indicates if the app is visible in the OneLogin portal.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `The parameters section contains parameterized attributes that have defined at the connector level as well as custom attributes that have been defined specifically for this app. Regardless of how they are defined, all parameters have the following attributes.`,
				},
				resource.Attribute{
					Name:        "attributes_transformations",
					Description: `Describes how the app's attributes should be transformed.`,
				},
				resource.Attribute{
					Name:        "default_values",
					Description: `Default Parameter values.`,
				},
				resource.Attribute{
					Name:        "include_in_saml_assertion",
					Description: `Dictates if the parameter needs to be included in a SAML assertion`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The attribute label.`,
				},
				resource.Attribute{
					Name:        "param_id",
					Description: `The parameter ID.`,
				},
				resource.Attribute{
					Name:        "param_key_name",
					Description: `The name of the parameter stored in OneLogin.`,
				},
				resource.Attribute{
					Name:        "provisioned_entitlements",
					Description: `Provisioned access entitlements for the app.`,
				},
				resource.Attribute{
					Name:        "safe_entitlements_enabled",
					Description: `Indicates whether entitlements can be created.`,
				},
				resource.Attribute{
					Name:        "skip_if_blank",
					Description: `Flag to let the SCIM provisioner know not include this value if it's blank.`,
				},
				resource.Attribute{
					Name:        "user_attribute_macros",
					Description: `When ` + "`" + `user_attribute_mappings` + "`" + ` is set to ` + "`" + `_macro_` + "`" + ` this macro will be used to assign the parameter value.`,
				},
				resource.Attribute{
					Name:        "user_attribute_mappings",
					Description: `A user attribute to map values from. For custom attributes the name of the attribute is prefixed with ` + "`" + `custom_attribute_` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `Parameter values.`,
				},
				resource.Attribute{
					Name:        "provisioning",
					Description: `Settings regarding the app's provisioning ability.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Indicates if provisioning is enabled for this app. ## Import An App can be imported via the OneLogin App ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import onelogin_apps.my_app <app id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `App's unique ID in OneLogin.`,
				},
				resource.Attribute{
					Name:        "allow_assumed_signin",
					Description: `App sign in allowed when user assumed by account administrator.`,
				},
				resource.Attribute{
					Name:        "auth_method",
					Description: `The apps auth method. Refer to the [OneLogin Apps Documentation](https://developers.onelogin.com/api-docs/2/apps/app-resource) for a comprehensive list of available auth methods.`,
				},
				resource.Attribute{
					Name:        "connector_id",
					Description: `ID of the apps underlying connector. Dictates the type of app (e.g. AWS Multi-Role App).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `App description.`,
				},
				resource.Attribute{
					Name:        "icon_url",
					Description: `The url for the app's icon.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The app's name.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `Notes about the app.`,
				},
				resource.Attribute{
					Name:        "tab_id",
					Description: `The tab in which to display in OneLogin portal.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Timestamp for app's last update.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Timestamp for app's creation.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `The security policy assigned to the app.`,
				},
				resource.Attribute{
					Name:        "visible",
					Description: `Indicates if the app is visible in the OneLogin portal.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `The parameters section contains parameterized attributes that have defined at the connector level as well as custom attributes that have been defined specifically for this app. Regardless of how they are defined, all parameters have the following attributes.`,
				},
				resource.Attribute{
					Name:        "attributes_transformations",
					Description: `Describes how the app's attributes should be transformed.`,
				},
				resource.Attribute{
					Name:        "default_values",
					Description: `Default Parameter values.`,
				},
				resource.Attribute{
					Name:        "include_in_saml_assertion",
					Description: `Dictates if the parameter needs to be included in a SAML assertion`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The attribute label.`,
				},
				resource.Attribute{
					Name:        "param_id",
					Description: `The parameter ID.`,
				},
				resource.Attribute{
					Name:        "param_key_name",
					Description: `The name of the parameter stored in OneLogin.`,
				},
				resource.Attribute{
					Name:        "provisioned_entitlements",
					Description: `Provisioned access entitlements for the app.`,
				},
				resource.Attribute{
					Name:        "safe_entitlements_enabled",
					Description: `Indicates whether entitlements can be created.`,
				},
				resource.Attribute{
					Name:        "skip_if_blank",
					Description: `Flag to let the SCIM provisioner know not include this value if it's blank.`,
				},
				resource.Attribute{
					Name:        "user_attribute_macros",
					Description: `When ` + "`" + `user_attribute_mappings` + "`" + ` is set to ` + "`" + `_macro_` + "`" + ` this macro will be used to assign the parameter value.`,
				},
				resource.Attribute{
					Name:        "user_attribute_mappings",
					Description: `A user attribute to map values from. For custom attributes the name of the attribute is prefixed with ` + "`" + `custom_attribute_` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `Parameter values.`,
				},
				resource.Attribute{
					Name:        "provisioning",
					Description: `Settings regarding the app's provisioning ability.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Indicates if provisioning is enabled for this app. ## Import An App can be imported via the OneLogin App ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import onelogin_apps.my_app <app id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "onelogin_onelogin_app_role_attachment",
			Category:         "Resources",
			ShortDescription: `Manage App Role Attachment resources.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "app_id",
					Description: `(Required) The id of the App resource to which the role should belong.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `(Required) The id of the Role being attached to the App. ## Attributes Reference No further attributes are exported. ## Import An App Role Attachment cannot be imported at this time.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "onelogin_onelogin_app_rule",
			Category:         "Resources",
			ShortDescription: `Manage App Rule resources.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "app_id",
					Description: `(Required) The id of the App resource to which the rule should belong.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Indicate if the rule should go into effect.`,
				},
				resource.Attribute{
					Name:        "match",
					Description: `(Required) Indicates how conditions should be matched. Must be one of ` + "`" + `all` + "`" + ` or ` + "`" + `any` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Rule's name`,
				},
				resource.Attribute{
					Name:        "position",
					Description: `(Required) Indicates the order of the rule. When ` + "`" + `null` + "`" + ` this will default to last position.`,
				},
				resource.Attribute{
					Name:        "conditions",
					Description: `(Required) An array of conditions that the user must meet in order for the rule to be applied.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `The source field to check. See [List Conditions](https://developers.onelogin.com/api-docs/2/app-rules/list-conditions) for possible values.`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `A valid operator for the selected condition source. See [List Condition Operators](https://developers.onelogin.com/api-docs/2/app-rules/list-condition-operators) for possible values.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `A plain text string or valid value for the selected condition source. See [List Condition Values](https://developers.onelogin.com/api-docs/2/app-rules/list-condition-values) for possible values.`,
				},
				resource.Attribute{
					Name:        "actions",
					Description: `(Required) An array of actions that will be applied to the users that are matched by the conditions.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `The action to apply. See [List Actions](https://developers.onelogin.com/api-docs/2/app-rules/list-conditions) for possible values.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `An array of strings. Only applicable to provisioned and set_`,
				},
				resource.Attribute{
					Name:        "expression",
					Description: `A regular expression to extract a value. Applies to provisionable, multi-selects, and string actions.`,
				},
				resource.Attribute{
					Name:        "scriptlet",
					Description: `A hash containing scriptlet code that returns a value. Scriptlets can not be modified and the same hash should not be applied to other applications.`,
				},
				resource.Attribute{
					Name:        "macro",
					Description: `A template to construct a value. Applies to default, string, and list actions. ## Attributes Reference No further attributes are exported. ## Import An App Rule cannot be imported at this time.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "onelogin_onelogin_auth_server",
			Category:         "Resources",
			ShortDescription: `Creates an Authentication Server Resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The resource's name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) A brief description about the resource.`,
				},
				resource.Attribute{
					Name:        "configuration",
					Description: `(Required) Configuration parameters`,
				},
				resource.Attribute{
					Name:        "resource_identifier",
					Description: `(Required) Unique identifier for the API that the Authorization Server will issue Access Tokens for.`,
				},
				resource.Attribute{
					Name:        "audiences",
					Description: `(Required) List of API endpoints that will be returned in Access Tokens.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "onelogin_onelogin_oidc_app",
			Category:         "Resources",
			ShortDescription: `Creates a OIDC Application.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The app's name.`,
				},
				resource.Attribute{
					Name:        "connector_id",
					Description: `(Required) The ID for the app connector, dictates the type of app (e.g. AWS Multi-Role App).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) App description.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) Notes about the app.`,
				},
				resource.Attribute{
					Name:        "visible",
					Description: `(Optional) Determine if app should be visible in OneLogin portal. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "allow_assumed_signin",
					Description: `(Optional) Enable sign in when user has been assumed by the account owner. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "provisioning",
					Description: `(Optional) Settings regarding the app's provisioning ability.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Indicates if provisioning is enabled for this app.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `(Optional) a list of custom parameters for this app.`,
				},
				resource.Attribute{
					Name:        "param_key_name",
					Description: `(Required) Name to represent the parameter in OneLogin.`,
				},
				resource.Attribute{
					Name:        "safe_entitlements_enabled",
					Description: `(Optional) Indicates that the parameter is used to support creating entitlements using OneLogin Mappings. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user_attribute_mappings",
					Description: `(Optional) A user attribute to map values from. For custom attributes prefix the name of the attribute with ` + "`" + `custom_attribute_` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "provisioned_entitlements",
					Description: `(Optional) Provisioned access entitlements for the app. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "skip_if_blank",
					Description: `(Optional) Flag to let the SCIM provisioner know not include this value if it's blank. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user_attribute_macros",
					Description: `(Optional) When ` + "`" + `user_attribute_mappings` + "`" + ` is set to ` + "`" + `_macro_` + "`" + ` this macro will be used to assign the parameter value.`,
				},
				resource.Attribute{
					Name:        "attributes_transformations",
					Description: `(Optional) Describes how the app's attributes should be transformed.`,
				},
				resource.Attribute{
					Name:        "default_values",
					Description: `(Optional) Default parameter values.`,
				},
				resource.Attribute{
					Name:        "include_in_saml_assertion",
					Description: `(Optional) When true, this parameter will be included in a SAML assertion payload.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Optional) The can only be set when creating a new parameter. It can not be updated.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Optional) Parameter values.`,
				},
				resource.Attribute{
					Name:        "configuration",
					Description: `OIDC settings that control the authentication flow e.g. redirect urls and token settings.`,
				},
				resource.Attribute{
					Name:        "redirect_uri",
					Description: `(Optional) The redirect_uri for the OIDC flow. Will be computed by API if not given.`,
				},
				resource.Attribute{
					Name:        "refresh_token_expiration_minutes",
					Description: `(Optional) Number of minutes for the refresh token to be valid. Defaults to 1 minute.`,
				},
				resource.Attribute{
					Name:        "login_url",
					Description: `(Optional) The login_url for the OIDC flow. Will be computed by API if not given.`,
				},
				resource.Attribute{
					Name:        "oidc_application_type",
					Description: `(Optional) Must be one of one of ` + "`" + `0` + "`" + ` (Web) or ` + "`" + `1` + "`" + ` (Native/Mobile). Defaults to 0.`,
				},
				resource.Attribute{
					Name:        "token_endpoint_auth_method",
					Description: `(Optional) Must be one of one of ` + "`" + `0` + "`" + ` (Basic) ` + "`" + `1` + "`" + ` (POST) ` + "`" + `2` + "`" + ` (Nonce/PKCE). ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `App's unique ID in OneLogin.`,
				},
				resource.Attribute{
					Name:        "allow_assumed_signin",
					Description: `App sign in allowed when user assumed by account administrator.`,
				},
				resource.Attribute{
					Name:        "auth_method",
					Description: `The apps auth method. Refer to the [OneLogin Apps Documentation](https://developers.onelogin.com/api-docs/2/apps/app-resource) for a comprehensive list of available auth methods.`,
				},
				resource.Attribute{
					Name:        "connector_id",
					Description: `ID of the apps underlying connector. Dictates the type of app (e.g. AWS Multi-Role App).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `App description.`,
				},
				resource.Attribute{
					Name:        "icon_url",
					Description: `The url for the app's icon.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The app's name.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `Notes about the app.`,
				},
				resource.Attribute{
					Name:        "tab_id",
					Description: `The tab in which to display in OneLogin portal.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Timestamp for app's last update.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Timestamp for app's creation.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `The security policy assigned to the app.`,
				},
				resource.Attribute{
					Name:        "visible",
					Description: `Indicates if the app is visible in the OneLogin portal.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `The parameters section contains parameterized attributes that have defined at the connector level as well as custom attributes that have been defined specifically for this app. Regardless of how they are defined, all parameters have the following attributes.`,
				},
				resource.Attribute{
					Name:        "attributes_transformations",
					Description: `Describes how the app's attributes should be transformed.`,
				},
				resource.Attribute{
					Name:        "default_values",
					Description: `Default parameter values.`,
				},
				resource.Attribute{
					Name:        "include_in_saml_assertion",
					Description: `Dictates if the parameter needs to be included in a SAML assertion`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The attribute label`,
				},
				resource.Attribute{
					Name:        "param_id",
					Description: `The parameter ID.`,
				},
				resource.Attribute{
					Name:        "param_key_name",
					Description: `The name of the parameter stored in OneLogin.`,
				},
				resource.Attribute{
					Name:        "provisioned_entitlements",
					Description: `Provisioned access entitlements for the app.`,
				},
				resource.Attribute{
					Name:        "safe_entitlements_enabled",
					Description: `Indicates whether entitlements can be created.`,
				},
				resource.Attribute{
					Name:        "skip_if_blank",
					Description: `Flag to let the SCIM provisioner know not include this value if it's blank.`,
				},
				resource.Attribute{
					Name:        "user_attribute_macros",
					Description: `When ` + "`" + `user_attribute_mappings` + "`" + ` is set to ` + "`" + `_macro_` + "`" + ` this macro will be used to assign the parameter value.`,
				},
				resource.Attribute{
					Name:        "user_attribute_mappings",
					Description: `A user attribute to map values from. For custom attributes the name of the attribute is prefixed with ` + "`" + `custom_attribute_` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `Parameter values.`,
				},
				resource.Attribute{
					Name:        "provisioning",
					Description: `Settings regarding the app's provisioning ability.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Indicates if provisioning is enabled for this app.`,
				},
				resource.Attribute{
					Name:        "redirect_uri",
					Description: `The redirect_uri for the OIDC flow.`,
				},
				resource.Attribute{
					Name:        "refresh_token_expiration_minutes",
					Description: `Number of minutes for the refresh token to be valid.`,
				},
				resource.Attribute{
					Name:        "login_url",
					Description: `The login_url for the OIDC flow.`,
				},
				resource.Attribute{
					Name:        "oidc_application_type",
					Description: `Indicates OIDC app type.`,
				},
				resource.Attribute{
					Name:        "token_endpoint_auth_method",
					Description: `Indicates the token endpoint authentication method. ## Import A OIDC App can be imported via the OneLogin App ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import onelogin_oidc_apps.my_oidc_app <app id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `App's unique ID in OneLogin.`,
				},
				resource.Attribute{
					Name:        "allow_assumed_signin",
					Description: `App sign in allowed when user assumed by account administrator.`,
				},
				resource.Attribute{
					Name:        "auth_method",
					Description: `The apps auth method. Refer to the [OneLogin Apps Documentation](https://developers.onelogin.com/api-docs/2/apps/app-resource) for a comprehensive list of available auth methods.`,
				},
				resource.Attribute{
					Name:        "connector_id",
					Description: `ID of the apps underlying connector. Dictates the type of app (e.g. AWS Multi-Role App).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `App description.`,
				},
				resource.Attribute{
					Name:        "icon_url",
					Description: `The url for the app's icon.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The app's name.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `Notes about the app.`,
				},
				resource.Attribute{
					Name:        "tab_id",
					Description: `The tab in which to display in OneLogin portal.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Timestamp for app's last update.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Timestamp for app's creation.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `The security policy assigned to the app.`,
				},
				resource.Attribute{
					Name:        "visible",
					Description: `Indicates if the app is visible in the OneLogin portal.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `The parameters section contains parameterized attributes that have defined at the connector level as well as custom attributes that have been defined specifically for this app. Regardless of how they are defined, all parameters have the following attributes.`,
				},
				resource.Attribute{
					Name:        "attributes_transformations",
					Description: `Describes how the app's attributes should be transformed.`,
				},
				resource.Attribute{
					Name:        "default_values",
					Description: `Default parameter values.`,
				},
				resource.Attribute{
					Name:        "include_in_saml_assertion",
					Description: `Dictates if the parameter needs to be included in a SAML assertion`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The attribute label`,
				},
				resource.Attribute{
					Name:        "param_id",
					Description: `The parameter ID.`,
				},
				resource.Attribute{
					Name:        "param_key_name",
					Description: `The name of the parameter stored in OneLogin.`,
				},
				resource.Attribute{
					Name:        "provisioned_entitlements",
					Description: `Provisioned access entitlements for the app.`,
				},
				resource.Attribute{
					Name:        "safe_entitlements_enabled",
					Description: `Indicates whether entitlements can be created.`,
				},
				resource.Attribute{
					Name:        "skip_if_blank",
					Description: `Flag to let the SCIM provisioner know not include this value if it's blank.`,
				},
				resource.Attribute{
					Name:        "user_attribute_macros",
					Description: `When ` + "`" + `user_attribute_mappings` + "`" + ` is set to ` + "`" + `_macro_` + "`" + ` this macro will be used to assign the parameter value.`,
				},
				resource.Attribute{
					Name:        "user_attribute_mappings",
					Description: `A user attribute to map values from. For custom attributes the name of the attribute is prefixed with ` + "`" + `custom_attribute_` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `Parameter values.`,
				},
				resource.Attribute{
					Name:        "provisioning",
					Description: `Settings regarding the app's provisioning ability.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Indicates if provisioning is enabled for this app.`,
				},
				resource.Attribute{
					Name:        "redirect_uri",
					Description: `The redirect_uri for the OIDC flow.`,
				},
				resource.Attribute{
					Name:        "refresh_token_expiration_minutes",
					Description: `Number of minutes for the refresh token to be valid.`,
				},
				resource.Attribute{
					Name:        "login_url",
					Description: `The login_url for the OIDC flow.`,
				},
				resource.Attribute{
					Name:        "oidc_application_type",
					Description: `Indicates OIDC app type.`,
				},
				resource.Attribute{
					Name:        "token_endpoint_auth_method",
					Description: `Indicates the token endpoint authentication method. ## Import A OIDC App can be imported via the OneLogin App ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import onelogin_oidc_apps.my_oidc_app <app id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "onelogin_onelogin_role",
			Category:         "Resources",
			ShortDescription: `Manage App Rule resources.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the role.`,
				},
				resource.Attribute{
					Name:        "apps",
					Description: `(Required) A list of app IDs for which the role applies.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `(Required) A list of user IDs for whom the role applies.`,
				},
				resource.Attribute{
					Name:        "admins",
					Description: `(Required) A list of IDs of users who administer the role. ## Attributes Reference No further attributes are exported. ## Import A role can be imported using the OneLogin Role ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import onelogin_roles.executive_admin <role id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "onelogin_onelogin_saml_app",
			Category:         "Resources",
			ShortDescription: `Creates a SAML Application.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The app's name.`,
				},
				resource.Attribute{
					Name:        "connector_id",
					Description: `(Required) The ID for the app connector, dictates the type of app (e.g. AWS Multi-Role App).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) App description.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) Notes about the app.`,
				},
				resource.Attribute{
					Name:        "visible",
					Description: `(Optional) Determine if app should be visible in OneLogin portal. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "allow_assumed_signin",
					Description: `(Optional) Enable sign in when user has been assumed by the account owner. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "provisioning",
					Description: `(Optional) Settings regarding the app's provisioning ability.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Indicates if provisioning is enabled for this app.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `(Optional) a list of custom parameters for this app.`,
				},
				resource.Attribute{
					Name:        "param_key_name",
					Description: `(Required) Name to represent the parameter in OneLogin.`,
				},
				resource.Attribute{
					Name:        "safe_entitlements_enabled",
					Description: `(Optional) Indicates that the parameter is used to support creating entitlements using OneLogin Mappings. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user_attribute_mappings",
					Description: `(Optional) A user attribute to map values from. For custom attributes prefix the name of the attribute with ` + "`" + `custom_attribute_` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "provisioned_entitlements",
					Description: `(Optional) Provisioned access entitlements for the app. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "skip_if_blank",
					Description: `(Optional) Flag to let the SCIM provisioner know not include this value if it's blank. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user_attribute_macros",
					Description: `(Optional) When ` + "`" + `user_attribute_mappings` + "`" + ` is set to ` + "`" + `_macro_` + "`" + ` this macro will be used to assign the parameter value.`,
				},
				resource.Attribute{
					Name:        "attributes_transformations",
					Description: `(Optional) Describes how the app's attributes should be transformed.`,
				},
				resource.Attribute{
					Name:        "default_values",
					Description: `(Optional) Default parameter values.`,
				},
				resource.Attribute{
					Name:        "include_in_saml_assertion",
					Description: `(Optional) When true, this parameter will be included in a SAML assertion payload.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Optional) The can only be set when creating a new parameter. It can not be updated.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Optional) Parameter values.`,
				},
				resource.Attribute{
					Name:        "signature_algorithm",
					Description: `(Required) Hashing algorithm to use for signing. Must be one of the following: ` + "`" + `SHA-1` + "`" + ` ` + "`" + `SHA-256` + "`" + ` ` + "`" + `SHA-348` + "`" + ` ` + "`" + `SHA-512` + "`" + ``,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `(Optional) The ID of the certificate.`,
				},
				resource.Attribute{
					Name:        "provider_arn",
					Description: `(Optional) The resource identifier of the provider. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `App's unique ID in OneLogin.`,
				},
				resource.Attribute{
					Name:        "allow_assumed_signin",
					Description: `App sign in allowed when user assumed by account administrator.`,
				},
				resource.Attribute{
					Name:        "auth_method",
					Description: `The apps auth method. Refer to the [OneLogin Apps Documentation](https://developers.onelogin.com/api-docs/2/apps/app-resource) for a comprehensive list of available auth methods.`,
				},
				resource.Attribute{
					Name:        "connector_id",
					Description: `ID of the apps underlying connector. Dictates the type of app (e.g. AWS Multi-Role App).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `App description.`,
				},
				resource.Attribute{
					Name:        "icon_url",
					Description: `The url for the app's icon.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The app's name.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `Notes about the app.`,
				},
				resource.Attribute{
					Name:        "tab_id",
					Description: `The tab in which to display in OneLogin portal.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Timestamp for app's last update.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Timestamp for app's creation.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `The security policy assigned to the app.`,
				},
				resource.Attribute{
					Name:        "visible",
					Description: `Indicates if the app is visible in the OneLogin portal.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `The parameters section contains parameterized attributes that have defined at the connector level as well as custom attributes that have been defined specifically for this app. Regardless of how they are defined, all parameters have the following attributes.`,
				},
				resource.Attribute{
					Name:        "attributes_transformations",
					Description: `Describes how the app's attributes should be transformed.`,
				},
				resource.Attribute{
					Name:        "default_values",
					Description: `Default parameter values.`,
				},
				resource.Attribute{
					Name:        "include_in_saml_assertion",
					Description: `Dictates if the parameter needs to be included in a SAML assertion`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The attribute label`,
				},
				resource.Attribute{
					Name:        "param_id",
					Description: `The parameter ID.`,
				},
				resource.Attribute{
					Name:        "param_key_name",
					Description: `The name of the parameter stored in OneLogin.`,
				},
				resource.Attribute{
					Name:        "provisioned_entitlements",
					Description: `Provisioned access entitlements for the app.`,
				},
				resource.Attribute{
					Name:        "safe_entitlements_enabled",
					Description: `Indicates whether entitlements can be created.`,
				},
				resource.Attribute{
					Name:        "skip_if_blank",
					Description: `Flag to let the SCIM provisioner know not include this value if it's blank.`,
				},
				resource.Attribute{
					Name:        "user_attribute_macros",
					Description: `When ` + "`" + `user_attribute_mappings` + "`" + ` is set to ` + "`" + `_macro_` + "`" + ` this macro will be used to assign the parameter value.`,
				},
				resource.Attribute{
					Name:        "user_attribute_mappings",
					Description: `A user attribute to map values from. For custom attributes the name of the attribute is prefixed with ` + "`" + `custom_attribute_` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `Parameter values.`,
				},
				resource.Attribute{
					Name:        "provisioning",
					Description: `Settings regarding the app's provisioning ability.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Indicates if provisioning is enabled for this app.`,
				},
				resource.Attribute{
					Name:        "configuration",
					Description: `SAML settings that control the authentication e.g. signature hashing algorithm or provider`,
				},
				resource.Attribute{
					Name:        "signature_algorithm",
					Description: `Hashing algorithm to use for signing.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `The ID of the certificate.`,
				},
				resource.Attribute{
					Name:        "provider_arn",
					Description: `The resource identifier of the provider.`,
				},
				resource.Attribute{
					Name:        "sso",
					Description: `The attributes included in the sso section are determined by the type of app. ` + "`" + `sso` + "`" + ` attributes are read only.`,
				},
				resource.Attribute{
					Name:        "metadata_url",
					Description: `The URL for the sso metadata.`,
				},
				resource.Attribute{
					Name:        "acs_url",
					Description: `The sso ACS URL.`,
				},
				resource.Attribute{
					Name:        "sls_url",
					Description: `The sso SLS URL.`,
				},
				resource.Attribute{
					Name:        "issuer",
					Description: `The certificate isssuer.`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `The SSO certificate generated by OneLogin.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Certificate value.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Certificate ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Certificate Name. ## Import A SAML App can be imported via the OneLogin App ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import onelogin_saml_apps.example_saml_app <app id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `App's unique ID in OneLogin.`,
				},
				resource.Attribute{
					Name:        "allow_assumed_signin",
					Description: `App sign in allowed when user assumed by account administrator.`,
				},
				resource.Attribute{
					Name:        "auth_method",
					Description: `The apps auth method. Refer to the [OneLogin Apps Documentation](https://developers.onelogin.com/api-docs/2/apps/app-resource) for a comprehensive list of available auth methods.`,
				},
				resource.Attribute{
					Name:        "connector_id",
					Description: `ID of the apps underlying connector. Dictates the type of app (e.g. AWS Multi-Role App).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `App description.`,
				},
				resource.Attribute{
					Name:        "icon_url",
					Description: `The url for the app's icon.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The app's name.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `Notes about the app.`,
				},
				resource.Attribute{
					Name:        "tab_id",
					Description: `The tab in which to display in OneLogin portal.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Timestamp for app's last update.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Timestamp for app's creation.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `The security policy assigned to the app.`,
				},
				resource.Attribute{
					Name:        "visible",
					Description: `Indicates if the app is visible in the OneLogin portal.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `The parameters section contains parameterized attributes that have defined at the connector level as well as custom attributes that have been defined specifically for this app. Regardless of how they are defined, all parameters have the following attributes.`,
				},
				resource.Attribute{
					Name:        "attributes_transformations",
					Description: `Describes how the app's attributes should be transformed.`,
				},
				resource.Attribute{
					Name:        "default_values",
					Description: `Default parameter values.`,
				},
				resource.Attribute{
					Name:        "include_in_saml_assertion",
					Description: `Dictates if the parameter needs to be included in a SAML assertion`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The attribute label`,
				},
				resource.Attribute{
					Name:        "param_id",
					Description: `The parameter ID.`,
				},
				resource.Attribute{
					Name:        "param_key_name",
					Description: `The name of the parameter stored in OneLogin.`,
				},
				resource.Attribute{
					Name:        "provisioned_entitlements",
					Description: `Provisioned access entitlements for the app.`,
				},
				resource.Attribute{
					Name:        "safe_entitlements_enabled",
					Description: `Indicates whether entitlements can be created.`,
				},
				resource.Attribute{
					Name:        "skip_if_blank",
					Description: `Flag to let the SCIM provisioner know not include this value if it's blank.`,
				},
				resource.Attribute{
					Name:        "user_attribute_macros",
					Description: `When ` + "`" + `user_attribute_mappings` + "`" + ` is set to ` + "`" + `_macro_` + "`" + ` this macro will be used to assign the parameter value.`,
				},
				resource.Attribute{
					Name:        "user_attribute_mappings",
					Description: `A user attribute to map values from. For custom attributes the name of the attribute is prefixed with ` + "`" + `custom_attribute_` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `Parameter values.`,
				},
				resource.Attribute{
					Name:        "provisioning",
					Description: `Settings regarding the app's provisioning ability.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Indicates if provisioning is enabled for this app.`,
				},
				resource.Attribute{
					Name:        "configuration",
					Description: `SAML settings that control the authentication e.g. signature hashing algorithm or provider`,
				},
				resource.Attribute{
					Name:        "signature_algorithm",
					Description: `Hashing algorithm to use for signing.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `The ID of the certificate.`,
				},
				resource.Attribute{
					Name:        "provider_arn",
					Description: `The resource identifier of the provider.`,
				},
				resource.Attribute{
					Name:        "sso",
					Description: `The attributes included in the sso section are determined by the type of app. ` + "`" + `sso` + "`" + ` attributes are read only.`,
				},
				resource.Attribute{
					Name:        "metadata_url",
					Description: `The URL for the sso metadata.`,
				},
				resource.Attribute{
					Name:        "acs_url",
					Description: `The sso ACS URL.`,
				},
				resource.Attribute{
					Name:        "sls_url",
					Description: `The sso SLS URL.`,
				},
				resource.Attribute{
					Name:        "issuer",
					Description: `The certificate isssuer.`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `The SSO certificate generated by OneLogin.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Certificate value.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Certificate ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Certificate Name. ## Import A SAML App can be imported via the OneLogin App ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import onelogin_saml_apps.example_saml_app <app id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "onelogin_onelogin_smarthook",
			Category:         "Resources",
			ShortDescription: `Manage SmartHook resources.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The name of the hook. Must be one of: ` + "`" + `user-migration` + "`" + ` ` + "`" + `pre-authentication` + "`" + ` ` + "`" + `pre-user-create` + "`" + ` ` + "`" + `post-user-create` + "`" + ` ` + "`" + `pre-user-update` + "`" + ` ` + "`" + `post-user-update` + "`" + ``,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Computed) The smarthook's status.`,
				},
				resource.Attribute{
					Name:        "packages",
					Description: `(Required) A list of public npm packages than will be installed as part of the function build process. These packages names must be on our allowlist. See Node Modules section of this doc. Packages can be any version and support the semantic versioning syntax used by NPM.`,
				},
				resource.Attribute{
					Name:        "function",
					Description: `(Required) A base64 encoded blob, or Heredoc string containing the javascript function code.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Required) Indicates if function is available for execution or not. Default true`,
				},
				resource.Attribute{
					Name:        "options",
					Description: `(Required if type = pre-authentication) A list of options for the hook`,
				},
				resource.Attribute{
					Name:        "risk_enabled",
					Description: `(Required) When true a risk score and risk reasons will be passed in the context. Only applies authentication time hooks. E.g. pre-authentication, user-migration. Default false`,
				},
				resource.Attribute{
					Name:        "location_enabled",
					Description: `(Required) When true an ip to location lookup is done and the location info is passed in the context. Only applies authentication time hooks. E.g. pre-authentication, user-migration. Default false`,
				},
				resource.Attribute{
					Name:        "retries",
					Description: `(Required) Number of retries if execution fails. Default 0, Max 4`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Required) The number of milliseconds to allow before timeout. Default 1000, Max 10000`,
				},
				resource.Attribute{
					Name:        "env_vars",
					Description: `(Required) An array of predefined environment variables to be supplied to the function at runtime.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `(Computed) Timestamp for smarthook's last update`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `(Computed) Timestamp for smarthook's last update ## Attributes Reference No further attributes are exported ## Import A SmartHook can be imported via the OneLogin SmartHook. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import onelogin_smarthooks.example <smarthook_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "onelogin_onelogin_user",
			Category:         "Resources",
			ShortDescription: `Manage User resources.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The user's username.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Required) The user's email.`,
				},
				resource.Attribute{
					Name:        "firstname",
					Description: `The user's first name`,
				},
				resource.Attribute{
					Name:        "lastname",
					Description: `The user's last name`,
				},
				resource.Attribute{
					Name:        "distinguished_name",
					Description: `The user's distinguished name`,
				},
				resource.Attribute{
					Name:        "samaccountname",
					Description: `The user's samaccount name`,
				},
				resource.Attribute{
					Name:        "userprincipalname",
					Description: `The user's user principal name`,
				},
				resource.Attribute{
					Name:        "member_of",
					Description: `The user's member_of`,
				},
				resource.Attribute{
					Name:        "phone",
					Description: `The user's phone number`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `The user's title`,
				},
				resource.Attribute{
					Name:        "company",
					Description: `The user's company`,
				},
				resource.Attribute{
					Name:        "department",
					Description: `The user's department`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `A comment about the user`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The user's state. Must be one of ` + "`" + `0: Unapproved` + "`" + ` ` + "`" + `1: Approved` + "`" + ` ` + "`" + `2: Rejected` + "`" + ` ` + "`" + `3: Unlicensed` + "`" + ``,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The user's status. Must be one of ` + "`" + `0: Unactivated` + "`" + ` ` + "`" + `1: Active` + "`" + ` ` + "`" + `2: Suspended` + "`" + ` ` + "`" + `3: Locked` + "`" + ` ` + "`" + `4: Password expired` + "`" + ` ` + "`" + `5: Awaiting password reset` + "`" + ` ` + "`" + `7: Password Pending` + "`" + ` ` + "`" + `8: Security questions required` + "`" + ``,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `The user's group_id`,
				},
				resource.Attribute{
					Name:        "directory_id",
					Description: `The user's directory_id`,
				},
				resource.Attribute{
					Name:        "trusted_idp_id",
					Description: `The user's trusted_idp_id`,
				},
				resource.Attribute{
					Name:        "manager_ad_id",
					Description: `The user's manager_ad_id`,
				},
				resource.Attribute{
					Name:        "manager_user_id",
					Description: `The user's manager_user_id`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `The user's external_id ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The user's id ## Import A User can be imported via the OneLogin User ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import onelogin_users.example 12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The user's id ## Import A User can be imported via the OneLogin User ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import onelogin_users.example 12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "onelogin_onelogin_user_mapping",
			Category:         "Resources",
			ShortDescription: `Manage User Mappings resources.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The resource's name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Indicates if a mapping is enabled.`,
				},
				resource.Attribute{
					Name:        "match",
					Description: `(Required) Indicates how conditions should be matched. Must be one of ` + "`" + `all` + "`" + ` or ` + "`" + `any` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "position",
					Description: `(Required) Indicates the ordering of the mapping. When ` + "`" + `null` + "`" + ` this will be placed last.`,
				},
				resource.Attribute{
					Name:        "conditions",
					Description: `(Required) An array of conditions that the user must meet in order for the mapping to be applied.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Required) The source field to check. See [List Conditions](https://developers.onelogin.com/api-docs/2/user-mappings/list-conditions) for possible values.`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `(Required) A valid operator for the selected condition source. See [List Condition Operators](https://developers.onelogin.com/api-docs/2/user-mappings/list-condition-operators) for possible values.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) A plain text string or valid value for the selected condition source. See [List Condition Values](https://developers.onelogin.com/api-docs/2/user-mappings/list-condition-values) for possible values.`,
				},
				resource.Attribute{
					Name:        "actions",
					Description: `(Required) The number of minutes until the token expires`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) The action to apply. See [List Actions](https://developers.onelogin.com/api-docs/2/user-mappings/list-conditions) for possible values.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) An array of strings. Items in the array will be a plain text string or valid value for the selected action. See [List Action Values](https://developers.onelogin.com/api-docs/2/user-mappings/list-action-values) for possible values. In most cases only a single item will be accepted in the array. ## Attributes Reference No further attributes are exported ## Import A User Mapping can be imported via the OneLogin User Mapping. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import onelogin_user_mappings.example <user_mapping_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"onelogin_onelogin_app":                 0,
		"onelogin_onelogin_app_role_attachment": 1,
		"onelogin_onelogin_app_rule":            2,
		"onelogin_onelogin_auth_server":         3,
		"onelogin_onelogin_oidc_app":            4,
		"onelogin_onelogin_role":                5,
		"onelogin_onelogin_saml_app":            6,
		"onelogin_onelogin_smarthook":           7,
		"onelogin_onelogin_user":                8,
		"onelogin_onelogin_user_mapping":        9,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
