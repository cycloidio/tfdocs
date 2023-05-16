package zia

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_activation_status",
			Category:         "Activation",
			ShortDescription: `"Get Activation Status changes".`,
			Description: `

The **zia_activation_status** data source allows to get information about the activation status of ZIA configurations.

~> As of right now, Terraform does not provide native support for commits or post-activation configuration, so configuration and policy activations are handled out-of-band. In order to handle the activation as part of the provider, a separate source code have been developed to generate a CLI binary.

`,
			Keywords: []string{
				"activation",
				"status",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `(Required) Activates configuration changes.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_admin_roles",
			Category:         "Admin & Role Management",
			ShortDescription: `Get information about ZIA administrator roles.`,
			Description: `

Use the **zia_admin_roles** data source to get information about an admin role created in the Zscaler Internet Access cloud or via the API. This data source can then be associated with a ZIA administrator account.

`,
			Keywords: []string{
				"admin",
				"role",
				"management",
				"roles",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Admin role to be exported.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the admin role to be exported. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_admin_users",
			Category:         "Admin & Role Management",
			ShortDescription: `Get information about ZIA administrator users.`,
			Description: `

Use the **zia_admin_users** data source to get information about an admin user account created in the Zscaler Internet Access cloud or via the API. This data source can then be associated with a ZIA administrator role.

`,
			Keywords: []string{
				"admin",
				"role",
				"management",
				"users",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "login_name",
					Description: `(Required) The email address of the admin user to be exported.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The username of the admin user to be exported.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the admin user to be exported. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(String) Admin or auditor's email address.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(String) Additional information about the admin or auditor.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Boolean) Indicates whether or not the admin account is disabled.`,
				},
				resource.Attribute{
					Name:        "is_auditor",
					Description: `(Boolean) Indicates whether the user is an auditor. This attribute is subject to change.`,
				},
				resource.Attribute{
					Name:        "is_exec_mobile_app_enabled",
					Description: `(Boolean) Indicates whether or not Executive Insights App access is enabled for the admin.`,
				},
				resource.Attribute{
					Name:        "is_non_editable",
					Description: `(Boolean) Indicates whether or not the admin can be edited or deleted.`,
				},
				resource.Attribute{
					Name:        "is_password_expired",
					Description: `(Boolean) Indicates whether or not an admin's password has expired.`,
				},
				resource.Attribute{
					Name:        "is_password_login_allowed",
					Description: `(Boolean) The default is true when SAML Authentication is disabled. When SAML Authentication is enabled, this can be set to false in order to force the admin to login via SSO only.`,
				},
				resource.Attribute{
					Name:        "is_product_update_comm_enabled",
					Description: `(Boolean) Communication setting for Product Update.`,
				},
				resource.Attribute{
					Name:        "is_security_report_comm_enabled",
					Description: `(Boolean) Communication for Security Report is enabled.`,
				},
				resource.Attribute{
					Name:        "is_service_update_comm_enabled",
					Description: `(Boolean) Communication setting for Service Update.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Set of Object) Role of the admin. This is not required for an auditor.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(String) The configured name of the entity`,
				},
				resource.Attribute{
					Name:        "admin_scope",
					Description: `(Set of Object) The admin's scope. Only applicable for the LOCATION_GROUP admin scope type, in which case this attribute gives the list of ID/name pairs of locations within the location group.`,
				},
				resource.Attribute{
					Name:        "scope_group_member_entities",
					Description: `(Number) Only applicable for the LOCATION_GROUP admin scope type, in which case this attribute gives the list of ID/name pairs of locations within the location group.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(String) The configured name of the entity`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(String) The admin scope type. The attribute name is subject to change.`,
				},
				resource.Attribute{
					Name:        "scope_entities",
					Description: `(String) Based on the admin scope type, the entities can be the ID/name pair of departments, locations, or location groups.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(String) The configured name of the entity`,
				},
				resource.Attribute{
					Name:        "exec_mobile_app_tokens",
					Description: `(List of Object)`,
				},
				resource.Attribute{
					Name:        "cloud",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `(Number)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "token_id",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "token_expiry",
					Description: `(Number)`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `(Number)`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `(String)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "email",
					Description: `(String) Admin or auditor's email address.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(String) Additional information about the admin or auditor.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Boolean) Indicates whether or not the admin account is disabled.`,
				},
				resource.Attribute{
					Name:        "is_auditor",
					Description: `(Boolean) Indicates whether the user is an auditor. This attribute is subject to change.`,
				},
				resource.Attribute{
					Name:        "is_exec_mobile_app_enabled",
					Description: `(Boolean) Indicates whether or not Executive Insights App access is enabled for the admin.`,
				},
				resource.Attribute{
					Name:        "is_non_editable",
					Description: `(Boolean) Indicates whether or not the admin can be edited or deleted.`,
				},
				resource.Attribute{
					Name:        "is_password_expired",
					Description: `(Boolean) Indicates whether or not an admin's password has expired.`,
				},
				resource.Attribute{
					Name:        "is_password_login_allowed",
					Description: `(Boolean) The default is true when SAML Authentication is disabled. When SAML Authentication is enabled, this can be set to false in order to force the admin to login via SSO only.`,
				},
				resource.Attribute{
					Name:        "is_product_update_comm_enabled",
					Description: `(Boolean) Communication setting for Product Update.`,
				},
				resource.Attribute{
					Name:        "is_security_report_comm_enabled",
					Description: `(Boolean) Communication for Security Report is enabled.`,
				},
				resource.Attribute{
					Name:        "is_service_update_comm_enabled",
					Description: `(Boolean) Communication setting for Service Update.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Set of Object) Role of the admin. This is not required for an auditor.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(String) The configured name of the entity`,
				},
				resource.Attribute{
					Name:        "admin_scope",
					Description: `(Set of Object) The admin's scope. Only applicable for the LOCATION_GROUP admin scope type, in which case this attribute gives the list of ID/name pairs of locations within the location group.`,
				},
				resource.Attribute{
					Name:        "scope_group_member_entities",
					Description: `(Number) Only applicable for the LOCATION_GROUP admin scope type, in which case this attribute gives the list of ID/name pairs of locations within the location group.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(String) The configured name of the entity`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(String) The admin scope type. The attribute name is subject to change.`,
				},
				resource.Attribute{
					Name:        "scope_entities",
					Description: `(String) Based on the admin scope type, the entities can be the ID/name pair of departments, locations, or location groups.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(String) The configured name of the entity`,
				},
				resource.Attribute{
					Name:        "exec_mobile_app_tokens",
					Description: `(List of Object)`,
				},
				resource.Attribute{
					Name:        "cloud",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `(Number)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "token_id",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "token_expiry",
					Description: `(Number)`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `(Number)`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `(String)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_auth_settings_urls",
			Category:         "User Authentication Settings",
			ShortDescription: `Gets a list of URLs that were exempted from cookie authentication and SSL Inspection.`,
			Description: `

Use the **zia_auth_settings_urls** data source to get a list of URLs that were exempted from cookie authentiation and SSL Inspection in the Zscaler Internet Access cloud or via the API. To learn more see [URL Format Guidelines](https://help.zscaler.com/zia/url-format-guidelines)

`,
			Keywords: []string{
				"user",
				"authentication",
				"settings",
				"auth",
				"urls",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_device_groups",
			Category:         "Device Groups",
			ShortDescription: `Get information about ZIA device groups.`,
			Description: `

Use the **zia_device_groups** data source to get information about a device group in the Zscaler Internet Access cloud or via the API. This data source can then be associated with resources such as: URL Filtering Rules

`,
			Keywords: []string{
				"device",
				"groups",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the device group to be exported.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The unique identifer for the device group. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(String) The unique identifer for the device group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(String) The device group name.`,
				},
				resource.Attribute{
					Name:        "group_type",
					Description: `(String) The device group type. i.e ` + "`" + `` + "`" + `ZCC_OS` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `NON_ZCC` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `CBI` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(String) The device group's description.`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `(String) The operating system (OS).`,
				},
				resource.Attribute{
					Name:        "predefined",
					Description: `(Boolean) Indicates whether this is a predefined device group. If this value is set to true, the group is predefined.`,
				},
				resource.Attribute{
					Name:        "device_names",
					Description: `(String) The names of devices that belong to the device group. The device names are comma-separated.`,
				},
				resource.Attribute{
					Name:        "device_count",
					Description: `(int) The number of devices within the group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(String) The unique identifer for the device group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(String) The device group name.`,
				},
				resource.Attribute{
					Name:        "group_type",
					Description: `(String) The device group type. i.e ` + "`" + `` + "`" + `ZCC_OS` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `NON_ZCC` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `CBI` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(String) The device group's description.`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `(String) The operating system (OS).`,
				},
				resource.Attribute{
					Name:        "predefined",
					Description: `(Boolean) Indicates whether this is a predefined device group. If this value is set to true, the group is predefined.`,
				},
				resource.Attribute{
					Name:        "device_names",
					Description: `(String) The names of devices that belong to the device group. The device names are comma-separated.`,
				},
				resource.Attribute{
					Name:        "device_count",
					Description: `(int) The number of devices within the group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_devices",
			Category:         "Device Groups",
			ShortDescription: `Get information about ZIA devices.`,
			Description: `

Use the **zia_devices** data source to get information about a device in the Zscaler Internet Access cloud or via the API. This data source can then be associated with resources such as: URL Filtering Rules

`,
			Keywords: []string{
				"device",
				"groups",
				"devices",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the devices to be exported.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The unique identifer for the devices. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(String) The unique identifer for the device group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(String) The device name.`,
				},
				resource.Attribute{
					Name:        "device_group_type",
					Description: `(String) The device group type. i.e ` + "`" + `` + "`" + `ZCC_OS` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `NON_ZCC` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `CBI` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "device_model",
					Description: `(String) The device model.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(String) The device group's description.`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `(String) The operating system (OS). ` + "`" + `` + "`" + `ANY` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `OTHER_OS` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `IOS` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `ANDROID_OS` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `WINDOWS_OS` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `MAC_OS` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `LINUX` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "os_version",
					Description: `(String) The operating system version.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(String) The device's description.`,
				},
				resource.Attribute{
					Name:        "owner_user_id",
					Description: `(int) The unique identifier of the device owner (i.e., user).`,
				},
				resource.Attribute{
					Name:        "owner_name",
					Description: `(String) The device owner's user name.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(String) The unique identifer for the device group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(String) The device name.`,
				},
				resource.Attribute{
					Name:        "device_group_type",
					Description: `(String) The device group type. i.e ` + "`" + `` + "`" + `ZCC_OS` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `NON_ZCC` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `CBI` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "device_model",
					Description: `(String) The device model.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(String) The device group's description.`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `(String) The operating system (OS). ` + "`" + `` + "`" + `ANY` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `OTHER_OS` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `IOS` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `ANDROID_OS` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `WINDOWS_OS` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `MAC_OS` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `LINUX` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "os_version",
					Description: `(String) The operating system version.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(String) The device's description.`,
				},
				resource.Attribute{
					Name:        "owner_user_id",
					Description: `(int) The unique identifier of the device owner (i.e., user).`,
				},
				resource.Attribute{
					Name:        "owner_name",
					Description: `(String) The device owner's user name.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_dlp_dictionaries",
			Category:         "Data Loss Prevention",
			ShortDescription: `Get information about ZIA DLP Dictionaries.`,
			Description: `

Use the **zia_dlp_dictionaries** data source to get information about a DLP dictionary option available in the Zscaler Internet Access.

` + "`" + `` + "`" + `` + "`" + `hcl
# Retrieve a DLP Dictionary by name
data "zia_dlp_dictionaries" "example"{
    name = "SALESFORCE_REPORT_LEAKAGE"
}
` + "`" + `` + "`" + `` + "`" + `

`,
			Keywords: []string{
				"data",
				"loss",
				"prevention",
				"dlp",
				"dictionaries",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) DLP dictionary name`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Unique identifier for the DLP dictionary ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "confidence_threshold",
					Description: `(String) he DLP confidence threshold. [` + "`" + `CONFIDENCE_LEVEL_LOW` + "`" + `, ` + "`" + `CONFIDENCE_LEVEL_MEDIUM` + "`" + ` ` + "`" + `CONFIDENCE_LEVEL_HIGH` + "`" + ` ]`,
				},
				resource.Attribute{
					Name:        "custom_phrase_match_type",
					Description: `(String) The DLP custom phrase match type. [ ` + "`" + `MATCH_ALL_CUSTOM_PHRASE_PATTERN_DICTIONARY` + "`" + `, ` + "`" + `MATCH_ANY_CUSTOM_PHRASE_PATTERN_DICTIONARY` + "`" + ` ]`,
				},
				resource.Attribute{
					Name:        "dictionary_type",
					Description: `(String) The DLP dictionary type. The cloud service API only supports custom DLP dictionaries that are using the ` + "`" + `PATTERNS_AND_PHRASES` + "`" + ` type.`,
				},
				resource.Attribute{
					Name:        "name_l10n_tag",
					Description: `(Boolean) Indicates whether the name is localized or not. This is always set to True for predefined DLP dictionaries. ` + "`" + `phrases` + "`" + ` - (List of Object) List containing the patterns used within a custom DLP dictionary. This attribute is not applicable to predefined DLP dictionaries`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(String) The action applied to a DLP dictionary using patterns`,
				},
				resource.Attribute{
					Name:        "pattern",
					Description: `(String) DLP dictionary pattern ` + "`" + `patterns` + "`" + ` - (List of Object) List containing the patterns used within a custom DLP dictionary. This attribute is not applicable to predefined DLP dictionaries`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(String) The action applied to a DLP dictionary using patterns`,
				},
				resource.Attribute{
					Name:        "pattern",
					Description: `(String) DLP dictionary pattern`,
				},
				resource.Attribute{
					Name:        "ignore_exact_match_idm_dict",
					Description: `(Boolean) Indicates whether to exclude documents that are a 100% match to already-indexed documents from triggering an Indexed Document Match (IDM) Dictionary.`,
				},
				resource.Attribute{
					Name:        "include_bin_numbers",
					Description: `(Boolean) A true value denotes that the specified Bank Identification Number (BIN) values are included in the Credit Cards dictionary. A false value denotes that the specified BIN values are excluded from the Credit Cards dictionary. Note: This field is applicable only to the predefined Credit Cards dictionary and its clones.`,
				},
				resource.Attribute{
					Name:        "bin_numbers",
					Description: `(Boolean) The list of Bank Identification Number (BIN) values that are included or excluded from the Credit Cards dictionary. BIN values can be specified only for Diners Club, Mastercard, RuPay, and Visa cards. Up to 512 BIN values can be configured in a dictionary. Note: This field is applicable only to the predefined Credit Cards dictionary and its clones.`,
				},
				resource.Attribute{
					Name:        "dict_template_id",
					Description: `(Number) ID of the predefined dictionary (original source dictionary) that is used for cloning. This field is applicable only to cloned dictionaries. Only a limited set of identification-based predefined dictionaries (e.g., Credit Cards, Social Security Numbers, National Identification Numbers, etc.) can be cloned. Up to 4 clones can be created from a predefined dictionary.`,
				},
				resource.Attribute{
					Name:        "predefined_clone",
					Description: `(Boolean) This field is set to true if the dictionary is cloned from a predefined dictionary. Otherwise, it is set to false.`,
				},
				resource.Attribute{
					Name:        "custom",
					Description: `(Boolean) This value is set to true for custom DLP dictionaries.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "confidence_threshold",
					Description: `(String) he DLP confidence threshold. [` + "`" + `CONFIDENCE_LEVEL_LOW` + "`" + `, ` + "`" + `CONFIDENCE_LEVEL_MEDIUM` + "`" + ` ` + "`" + `CONFIDENCE_LEVEL_HIGH` + "`" + ` ]`,
				},
				resource.Attribute{
					Name:        "custom_phrase_match_type",
					Description: `(String) The DLP custom phrase match type. [ ` + "`" + `MATCH_ALL_CUSTOM_PHRASE_PATTERN_DICTIONARY` + "`" + `, ` + "`" + `MATCH_ANY_CUSTOM_PHRASE_PATTERN_DICTIONARY` + "`" + ` ]`,
				},
				resource.Attribute{
					Name:        "dictionary_type",
					Description: `(String) The DLP dictionary type. The cloud service API only supports custom DLP dictionaries that are using the ` + "`" + `PATTERNS_AND_PHRASES` + "`" + ` type.`,
				},
				resource.Attribute{
					Name:        "name_l10n_tag",
					Description: `(Boolean) Indicates whether the name is localized or not. This is always set to True for predefined DLP dictionaries. ` + "`" + `phrases` + "`" + ` - (List of Object) List containing the patterns used within a custom DLP dictionary. This attribute is not applicable to predefined DLP dictionaries`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(String) The action applied to a DLP dictionary using patterns`,
				},
				resource.Attribute{
					Name:        "pattern",
					Description: `(String) DLP dictionary pattern ` + "`" + `patterns` + "`" + ` - (List of Object) List containing the patterns used within a custom DLP dictionary. This attribute is not applicable to predefined DLP dictionaries`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(String) The action applied to a DLP dictionary using patterns`,
				},
				resource.Attribute{
					Name:        "pattern",
					Description: `(String) DLP dictionary pattern`,
				},
				resource.Attribute{
					Name:        "ignore_exact_match_idm_dict",
					Description: `(Boolean) Indicates whether to exclude documents that are a 100% match to already-indexed documents from triggering an Indexed Document Match (IDM) Dictionary.`,
				},
				resource.Attribute{
					Name:        "include_bin_numbers",
					Description: `(Boolean) A true value denotes that the specified Bank Identification Number (BIN) values are included in the Credit Cards dictionary. A false value denotes that the specified BIN values are excluded from the Credit Cards dictionary. Note: This field is applicable only to the predefined Credit Cards dictionary and its clones.`,
				},
				resource.Attribute{
					Name:        "bin_numbers",
					Description: `(Boolean) The list of Bank Identification Number (BIN) values that are included or excluded from the Credit Cards dictionary. BIN values can be specified only for Diners Club, Mastercard, RuPay, and Visa cards. Up to 512 BIN values can be configured in a dictionary. Note: This field is applicable only to the predefined Credit Cards dictionary and its clones.`,
				},
				resource.Attribute{
					Name:        "dict_template_id",
					Description: `(Number) ID of the predefined dictionary (original source dictionary) that is used for cloning. This field is applicable only to cloned dictionaries. Only a limited set of identification-based predefined dictionaries (e.g., Credit Cards, Social Security Numbers, National Identification Numbers, etc.) can be cloned. Up to 4 clones can be created from a predefined dictionary.`,
				},
				resource.Attribute{
					Name:        "predefined_clone",
					Description: `(Boolean) This field is set to true if the dictionary is cloned from a predefined dictionary. Otherwise, it is set to false.`,
				},
				resource.Attribute{
					Name:        "custom",
					Description: `(Boolean) This value is set to true for custom DLP dictionaries.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_dlp_engines",
			Category:         "Data Loss Prevention",
			ShortDescription: `Get information about ZIA DLP Engines.`,
			Description: `

Use the **zia_dlp_engines** data source to get information about a ZIA DLP Engines in the Zscaler Internet Access cloud or via the API.

`,
			Keywords: []string{
				"data",
				"loss",
				"prevention",
				"dlp",
				"engines",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The DLP engine name as configured by the admin. This attribute is required in POST and PUT requests for custom DLP engines. ### Optional`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) The unique identifier for the DLP engine.`,
				},
				resource.Attribute{
					Name:        "predefined_engine_name",
					Description: `(String) The name of the predefined DLP engine.`,
				},
				resource.Attribute{
					Name:        "engine_expression",
					Description: `(String) The boolean logical operator in which various DLP dictionaries are combined within a DLP engine's expression.`,
				},
				resource.Attribute{
					Name:        "custom_dlp_engine",
					Description: `(Bool) Indicates whether this is a custom DLP engine. If this value is set to true, the engine is custom.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(String) The DLP engine's description.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_dlp_icap_servers",
			Category:         "Data Loss Prevention",
			ShortDescription: `Gets a the list of DLP servers using ICAP`,
			Description: `

Use the **zia_dlp_engines** data source to get information about a the list of DLP servers using ICAP in the Zscaler Internet Access cloud or via the API.

`,
			Keywords: []string{
				"data",
				"loss",
				"prevention",
				"dlp",
				"icap",
				"servers",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The DLP server name as configured by the admin. ### Optional`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) The unique identifier for a DLP server.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(String) The DLP server name.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(String) The DLP server URL.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Bool) The DLP server status`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_dlp_idm_profiles",
			Category:         "Data Loss Prevention",
			ShortDescription: `Get information about ZIA DLP IDM Profile.`,
			Description: `

Use the **zia_dlp_idm_profile** data source to get information about a ZIA DLP IDM Profile in the Zscaler Internet Access cloud or via the API.

`,
			Keywords: []string{
				"data",
				"loss",
				"prevention",
				"dlp",
				"idm",
				"profiles",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "profile_name",
					Description: `(Required) The IDM template name, which is unique per Index Tool. ### Optional`,
				},
				resource.Attribute{
					Name:        "profile_id",
					Description: `(String) The identifier (1-64) for the IDM template (i.e., IDM profile) that is unique within the organization.`,
				},
				resource.Attribute{
					Name:        "profile_desc",
					Description: `(String) The IDM template's description.`,
				},
				resource.Attribute{
					Name:        "profile_type",
					Description: `(String) The IDM template's type. The returned values are:`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(String) The fully qualified domain name (FQDN) of the IDM template's host machine.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Number) The port number of the IDM template's host machine.`,
				},
				resource.Attribute{
					Name:        "profile_dir_path",
					Description: `(String) The IDM template's directory file path, where all files are present.`,
				},
				resource.Attribute{
					Name:        "schedule_type",
					Description: `(String) The schedule type for the IDM template's schedule (i.e., Monthly, Weekly, Daily, or None). This attribute is required by PUT and POST requests.`,
				},
				resource.Attribute{
					Name:        "schedule_day",
					Description: `(Number) The day the IDM template is scheduled for. This attribute is required by PUT and POST requests.`,
				},
				resource.Attribute{
					Name:        "schedule_day_of_month",
					Description: `(Number) The day of the month that the IDM template is scheduled for. This attribute is required by PUT and POST requests, and when scheduleType is set to MONTHLY.`,
				},
				resource.Attribute{
					Name:        "schedule_time",
					Description: `(Number) The time of the day (in minutes) that the IDM template is scheduled for. For example: at 3am= 180 mins. This attribute is required by PUT and POST requests.`,
				},
				resource.Attribute{
					Name:        "schedule_disabled",
					Description: `(Bool) If set to true, the schedule for the IDM template is temporarily in a disabled state. This attribute is required by PUT requests in order to disable or enable a schedule.`,
				},
				resource.Attribute{
					Name:        "upload_status",
					Description: `(Bool) The status of the file uploaded to the Index Tool for the IDM template.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(String) The username to be used on the IDM template's host machine.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Number) The version number for the IDM template.`,
				},
				resource.Attribute{
					Name:        "idm_client",
					Description: `(String ) The unique identifer for the Index Tool that was used to create the IDM template. This attribute is required by POST requests, but ignored if provided in PUT requests.`,
				},
				resource.Attribute{
					Name:        "volume_of_documents",
					Description: `(Number) The total volume of all the documents associated to the IDM template.`,
				},
				resource.Attribute{
					Name:        "num_documents",
					Description: `(Number) The number of documents associated to the IDM template.`,
				},
				resource.Attribute{
					Name:        "last_modified_time",
					Description: `(Number) The date and time the IDM template was last modified.`,
				},
				resource.Attribute{
					Name:        "modified_by",
					Description: `(Number) The admin that modified the IDM template last.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(String) The configured name of the entity`,
				},
				resource.Attribute{
					Name:        "extensions",
					Description: `(Map) The configured name of the entity`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_dlp_idm_profiles_lite",
			Category:         "Data Loss Prevention",
			ShortDescription: `Get information about ZIA DLP IDM Profile Lite.`,
			Description: `

Use the **zia_dlp_idm_profile_lite** data source to get summarized information about a ZIA DLP IDM Profile Lite in the Zscaler Internet Access cloud or via the API.

`,
			Keywords: []string{
				"data",
				"loss",
				"prevention",
				"dlp",
				"idm",
				"profiles",
				"lite",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) The IDM template name. ### Optional`,
				},
				resource.Attribute{
					Name:        "profile_id",
					Description: `(Number) The unique identifier for the IDM template (i.e., IDM profile).`,
				},
				resource.Attribute{
					Name:        "num_documents",
					Description: `(Number) The number of documents associated to the IDM template.`,
				},
				resource.Attribute{
					Name:        "client_vm",
					Description: `(Number) This is an immutable reference to an entity. which mainly consists of id and name.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(String) The configured name of the entity`,
				},
				resource.Attribute{
					Name:        "extensions",
					Description: `(Map) The configured name of the entity`,
				},
				resource.Attribute{
					Name:        "last_modified_time",
					Description: `(Number) The date and time the IDM template was last modified.`,
				},
				resource.Attribute{
					Name:        "modified_by",
					Description: `(Number) The admin that modified the IDM template last.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(String) The configured name of the entity`,
				},
				resource.Attribute{
					Name:        "extensions",
					Description: `(Map) The configured name of the entity`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_dlp_incident_receiver_servers",
			Category:         "Data Loss Prevention",
			ShortDescription: `Get information about ZIA DLP Incident Receiver Servers.`,
			Description: `

Use the **zia_dlp_incident_receiver_servers** data source to get information about a ZIA DLP Incident Receiver Server in the Zscaler Internet Access cloud or via the API.

`,
			Keywords: []string{
				"data",
				"loss",
				"prevention",
				"dlp",
				"incident",
				"receiver",
				"servers",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The DLP Incident Receiver Server name as configured by the admin. ### Optional`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) The unique identifier for the DLP engine.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(String) The Incident Receiver server URL.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(String) The status of the Incident Receiver. The returned values are:`,
				},
				resource.Attribute{
					Name:        "flags",
					Description: `(Number) The Incident Receiver server flag.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_dlp_notification_templates",
			Category:         "Data Loss Prevention",
			ShortDescription: `Get information about DLP Notification Templates.`,
			Description: `

Use the **zia_dlp_notification_templates** data source to get information about a ZIA DLP Notification Templates in the Zscaler Internet Access cloud or via the API.

`,
			Keywords: []string{
				"data",
				"loss",
				"prevention",
				"dlp",
				"notification",
				"templates",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The DLP policy rule name. ### Optional`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The unique identifier for a DLP notification template.`,
				},
				resource.Attribute{
					Name:        "plain_text_message",
					Description: `(Optional) The template for the plain text UTF-8 message body that must be displayed in the DLP notification email.`,
				},
				resource.Attribute{
					Name:        "html_message",
					Description: `(Optional) The template for the HTML message body that must be displayed in the DLP notification email.`,
				},
				resource.Attribute{
					Name:        "subject",
					Description: `(Optional) The Subject line that is displayed within the DLP notification email.`,
				},
				resource.Attribute{
					Name:        "attach_content",
					Description: `(Optional) If set to true, the content that is violation is attached to the DLP notification email.`,
				},
				resource.Attribute{
					Name:        "tls_enabled",
					Description: `(Optional) If set to true, the content that is violation is attached to the DLP notification email.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_dlp_web_rules",
			Category:         "Data Loss Prevention",
			ShortDescription: `Get information about ZIA DLP Web Rules.`,
			Description: `

Use the **zia_dlp_web_rules** data source to get information about a ZIA DLP Web Rules in the Zscaler Internet Access cloud or via the API.

`,
			Keywords: []string{
				"data",
				"loss",
				"prevention",
				"dlp",
				"web",
				"rules",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The DLP policy rule name. rules. ### Optional`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(String) The description of the DLP policy rule.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Number) The rule order of execution for the DLP policy rule with respect to other`,
				},
				resource.Attribute{
					Name:        "external_auditor_email",
					Description: `(String) The email address of an external auditor to whom DLP email notifications are sent.`,
				},
				resource.Attribute{
					Name:        "match_only",
					Description: `(Bool) The match only criteria for DLP engines.`,
				},
				resource.Attribute{
					Name:        "without_content_inspection",
					Description: `(Bool) Indicates a DLP policy rule without content inspection, when the value is set to true.`,
				},
				resource.Attribute{
					Name:        "ocr_enabled",
					Description: `(Bool) Enables or disables image file scanning.`,
				},
				resource.Attribute{
					Name:        "zscaler_incident_receiver",
					Description: `(Bool) Indicates whether a Zscaler Incident Receiver is associated to the DLP policy rule.`,
				},
				resource.Attribute{
					Name:        "last_modified_time",
					Description: `(Number) Timestamp when the DLP policy rule was last modified.`,
				},
				resource.Attribute{
					Name:        "access_control",
					Description: `(String) The access privilege for this DLP policy rule based on the admin's state. The supported values are:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(String) The action taken when traffic matches the DLP policy rule criteria. The supported values are:`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(String) Enables or disables the DLP policy rule.. The supported values are:`,
				},
				resource.Attribute{
					Name:        "file_types",
					Description: `(String) The list of file types to which the DLP policy rule must be applied. For the complete list of supported file types refer to the [ZIA API documentation](https://help.zscaler.com/zia/data-loss-prevention#/webDlpRules-post)`,
				},
				resource.Attribute{
					Name:        "cloud_applications",
					Description: `(Optional) The list of cloud applications to which the DLP policy rule must be applied. For the complete list of supported cloud applications refer to the [ZIA API documentation](https://help.zscaler.com/zia/data-loss-prevention#/webDlpRules-post)`,
				},
				resource.Attribute{
					Name:        "last_modified_by",
					Description: `(Number) The admin that modified the DLP policy rule last.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "notification_template",
					Description: `(Optional) The template used for DLP notification emails.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "auditor",
					Description: `(Optional) The auditor to which the DLP policy rule must be applied.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "url_categories",
					Description: `(Optional) The list of URL categories to which the DLP policy rule must be applied.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "dlp_engines",
					Description: `(Optional) The list of DLP engines to which the DLP policy rule must be applied.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "locations",
					Description: `(Optional) The Name-ID pairs of locations to which the DLP policy rule must be applied. Maximum of up to ` + "`" + `8` + "`" + ` locations. When not used it implies ` + "`" + `Any` + "`" + ` to apply the rule to all locations.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "location_groups",
					Description: `(Optional) The Name-ID pairs of locations groups to which the DLP policy rule must be applied. Maximum of up to ` + "`" + `32` + "`" + ` location groups. When not used it implies ` + "`" + `Any` + "`" + ` to apply the rule to all location groups.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `(Optional) The Name-ID pairs of users to which the DLP policy rule must be applied. Maximum of up to ` + "`" + `4` + "`" + ` users. When not used it implies ` + "`" + `Any` + "`" + ` to apply the rule to all users.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "excluded_users",
					Description: `(Optional) The name-ID pairs of the users that are excluded from the DLP policy rule. Maximum of up to ` + "`" + `256` + "`" + ` users.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `(Optional) The Name-ID pairs of groups to which the DLP policy rule must be applied. Maximum of up to ` + "`" + `8` + "`" + ` groups. When not used it implies ` + "`" + `Any` + "`" + ` to apply the rule to all groups.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "excluded_groups",
					Description: `(Optional) The name-ID pairs of the groups that are excluded from the DLP policy rule. Maximum of up to ` + "`" + `256` + "`" + ` groups.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "departments",
					Description: `(Optional) The name-ID pairs of the departments that are excluded from the DLP policy rule.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "excluded_departments",
					Description: `(Optional) The name-ID pairs of the groups that are excluded from the DLP policy rule. Maximum of up to ` + "`" + `256` + "`" + ` departments.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "time_windows",
					Description: `(Optional) The Name-ID pairs of time windows to which the DLP policy rule must be applied. Maximum of up to ` + "`" + `2` + "`" + ` time intervals. When not used it implies ` + "`" + `always` + "`" + ` to apply the rule to all time intervals.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) The Name-ID pairs of rule labels associated to the DLP policy rule.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "icap_server",
					Description: `(Optional) The DLP server, using ICAP, to which the transaction content is forwarded.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_firewall_filtering_destination_groups",
			Category:         "Firewall Policies",
			ShortDescription: `Get information about IP destination groups.`,
			Description: `

Use the **zia_firewall_filtering_destination_groups** data source to get information about IP destination groups option available in the Zscaler Internet Access cloud firewall. This data source can then be associated with a ZIA firewall filtering rule.

`,
			Keywords: []string{
				"firewall",
				"policies",
				"filtering",
				"destination",
				"groups",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the destination group to be exported.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the destination group resource. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(String) Additional information about the destination IP group`,
				},
				resource.Attribute{
					Name:        "addresses",
					Description: `(List of String) Destination IP addresses within the group`,
				},
				resource.Attribute{
					Name:        "countries",
					Description: `(List of String) Destination IP address counties. You can identify destinations based on the location of a server.`,
				},
				resource.Attribute{
					Name:        "ip_categories",
					Description: `(List of String) Destination IP address URL categories. You can identify destinations based on the URL category of the domain.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(String) Destination IP group type (i.e., the group can contain destination IP addresses or FQDNs)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(String) Additional information about the destination IP group`,
				},
				resource.Attribute{
					Name:        "addresses",
					Description: `(List of String) Destination IP addresses within the group`,
				},
				resource.Attribute{
					Name:        "countries",
					Description: `(List of String) Destination IP address counties. You can identify destinations based on the location of a server.`,
				},
				resource.Attribute{
					Name:        "ip_categories",
					Description: `(List of String) Destination IP address URL categories. You can identify destinations based on the URL category of the domain.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(String) Destination IP group type (i.e., the group can contain destination IP addresses or FQDNs)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_firewall_filtering_ip_source_groups",
			Category:         "Firewall Policies",
			ShortDescription: `Get information about IP Source groups.`,
			Description: `

Use the **zia_firewall_filtering_ip_source_groups** data source to get information about ip source groups available in the Zscaler Internet Access cloud firewall. This data source can then be associated with a ZIA firewall filtering rule.

`,
			Keywords: []string{
				"firewall",
				"policies",
				"filtering",
				"ip",
				"source",
				"groups",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the ip source group to be exported.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the ip source group resource. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of this resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "ip_addresses",
					Description: `(List of String)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of this resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "ip_addresses",
					Description: `(List of String)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_firewall_filtering_network_application",
			Category:         "Firewall Policies",
			ShortDescription: `Get information about ZIA firewall rule network application.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"policies",
				"filtering",
				"network",
				"application",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The name of the ip source group to be exported.`,
				},
				resource.Attribute{
					Name:        "locale",
					Description: `(Optional) ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "deprecated",
					Description: `(Boolean)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "parent_category",
					Description: `(String)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "deprecated",
					Description: `(Boolean)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "parent_category",
					Description: `(String)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_firewall_filtering_network_application_groups",
			Category:         "Firewall Policies",
			ShortDescription: `Get information about firewall rule network application groups.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"policies",
				"filtering",
				"network",
				"application",
				"groups",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the ip source group to be exported.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the ip source group resource. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "network_applications",
					Description: `(List of String)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of this resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "network_applications",
					Description: `(List of String)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of this resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_firewall_filtering_network_service",
			Category:         "Firewall Policies",
			ShortDescription: `Get information about firewall rule network services.`,
			Description: `

The **zia_firewall_filtering_network_service** data source to get information about a network service available in the Zscaler Internet Access cloud firewall. This data source can then be associated with a ZIA firewall filtering network service rule.

`,
			Keywords: []string{
				"firewall",
				"policies",
				"filtering",
				"network",
				"service",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the application layer service that you want to control. It can include any character and spaces.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the application layer service to be exported. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(String) (Optional) Enter additional notes or information. The description cannot exceed 10240 characters.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(String) - Supported values are: ` + "`" + `STANDARD` + "`" + `, ` + "`" + `PREDEFINED` + "`" + ` and ` + "`" + `CUSTOM` + "`" + ``,
				},
				resource.Attribute{
					Name:        "is_name_l10n_tag",
					Description: `(Bool) - Default: false`,
				},
				resource.Attribute{
					Name:        "src_tcp_ports",
					Description: `(Optional) The TCP source port number (example: 50) or port number range (example: 1000-1050), if any, that is used by the network service`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `(Number)`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `(Number)`,
				},
				resource.Attribute{
					Name:        "dest_tcp_ports",
					Description: `(Required) The TCP destination port number (example: 50) or port number range (example: 1000-1050), if any, that is used by the network service.`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `(Number)`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `(Number)`,
				},
				resource.Attribute{
					Name:        "src_udp_ports",
					Description: `The UDP source port number (example: 50) or port number range (example: 1000-1050), if any, that is used by the network service.`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `(Number)`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `(Number)`,
				},
				resource.Attribute{
					Name:        "dest_udp_ports",
					Description: `The UDP source port number (example: 50) or port number range (example: 1000-1050), if any, that is used by the network service.`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `(Number)`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `(Number)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(String) (Optional) Enter additional notes or information. The description cannot exceed 10240 characters.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(String) - Supported values are: ` + "`" + `STANDARD` + "`" + `, ` + "`" + `PREDEFINED` + "`" + ` and ` + "`" + `CUSTOM` + "`" + ``,
				},
				resource.Attribute{
					Name:        "is_name_l10n_tag",
					Description: `(Bool) - Default: false`,
				},
				resource.Attribute{
					Name:        "src_tcp_ports",
					Description: `(Optional) The TCP source port number (example: 50) or port number range (example: 1000-1050), if any, that is used by the network service`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `(Number)`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `(Number)`,
				},
				resource.Attribute{
					Name:        "dest_tcp_ports",
					Description: `(Required) The TCP destination port number (example: 50) or port number range (example: 1000-1050), if any, that is used by the network service.`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `(Number)`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `(Number)`,
				},
				resource.Attribute{
					Name:        "src_udp_ports",
					Description: `The UDP source port number (example: 50) or port number range (example: 1000-1050), if any, that is used by the network service.`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `(Number)`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `(Number)`,
				},
				resource.Attribute{
					Name:        "dest_udp_ports",
					Description: `The UDP source port number (example: 50) or port number range (example: 1000-1050), if any, that is used by the network service.`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `(Number)`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `(Number)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_firewall_filtering_network_service_groups",
			Category:         "Firewall Policies",
			ShortDescription: `Get information about firewall rule network service groups.`,
			Description: `

Use the **zia_firewall_filtering_network_service_groups** data source to get information about a network service groups available in the Zscaler Internet Access cloud firewall. This data source can then be associated with a ZIA firewall filtering network service rule.

`,
			Keywords: []string{
				"firewall",
				"policies",
				"filtering",
				"network",
				"service",
				"groups",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the ip source group to be exported.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the ip source group to be exported. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `(Number) The ID of this resource.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "dest_tcp_ports",
					Description: `(Required) The TCP destination port number (example: 50) or port number range (example: 1000-1050), if any, that is used by the network service. - ` + "`" + `start` + "`" + ` - (Number) - ` + "`" + `end` + "`" + ` - (Number)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(String) - Supported values are: ` + "`" + `STANDARD` + "`" + `, ` + "`" + `PREDEFINED` + "`" + ` and ` + "`" + `CUSTOM` + "`" + ``,
				},
				resource.Attribute{
					Name:        "is_name_l10n_tag",
					Description: `(Bool) - Default: false`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `(Number) The ID of this resource.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "dest_tcp_ports",
					Description: `(Required) The TCP destination port number (example: 50) or port number range (example: 1000-1050), if any, that is used by the network service. - ` + "`" + `start` + "`" + ` - (Number) - ` + "`" + `end` + "`" + ` - (Number)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(String) - Supported values are: ` + "`" + `STANDARD` + "`" + `, ` + "`" + `PREDEFINED` + "`" + ` and ` + "`" + `CUSTOM` + "`" + ``,
				},
				resource.Attribute{
					Name:        "is_name_l10n_tag",
					Description: `(Bool) - Default: false`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_firewall_filtering_rule",
			Category:         "Firewall Policies",
			ShortDescription: `Get information about firewall filtering rule.`,
			Description: `

Use the **zia_firewall_filtering_rule** data source to get information about a cloud firewall rule available in the Zscaler Internet Access cloud firewall.

`,
			Keywords: []string{
				"firewall",
				"policies",
				"filtering",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Firewall Filtering policy rule`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Unique identifier for the Firewall Filtering policy rule ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Enter additional notes or information. The description cannot exceed 10,240 characters.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Required) Policy rules are evaluated in ascending numerical order (Rule 1 before Rule 2, and so on), and the Rule Order reflects this rule's place in the order.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) An enabled rule is actively enforced. A disabled rule is not actively enforced but does not lose its place in the Rule Order. The service skips it and moves to the next rule.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Choose the action of the service when packets match the rule. The following actions are accepted: ` + "`" + `ALLOW` + "`" + `, ` + "`" + `BLOCK_DROP` + "`" + `, ` + "`" + `BLOCK_RESET` + "`" + `, ` + "`" + `BLOCK_ICMP` + "`" + `, ` + "`" + `EVAL_NWAPP` + "`" + ``,
				},
				resource.Attribute{
					Name:        "rank",
					Description: `(Optional) By default, the admin ranking is disabled. To use this feature, you must enable admin rank. The default value is ` + "`" + `7` + "`" + `. ` + "`" + `Who, Where and When` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "locations",
					Description: `(Optional) You can manually select up to ` + "`" + `8` + "`" + ` locations. When not used it implies ` + "`" + `Any` + "`" + ` to apply the rule to all groups. - ` + "`" + `id` + "`" + ` - (String) Identifier that uniquely identifies an entity - ` + "`" + `name` + "`" + ` - (String) The configured name of the entity - ` + "`" + `extensions` + "`" + ` - (Map of String)`,
				},
				resource.Attribute{
					Name:        "location_groups",
					Description: `(Optional) You can manually select up to ` + "`" + `32` + "`" + ` location groups. When not used it implies ` + "`" + `Any` + "`" + ` to apply the rule to all location groups. - ` + "`" + `id` + "`" + ` - (String) Identifier that uniquely identifies an entity - ` + "`" + `name` + "`" + ` - (String) The configured name of the entity - ` + "`" + `extensions` + "`" + ` - (Map of String)`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `(Optional) You can manually select up to ` + "`" + `4` + "`" + ` general and/or special users. When not used it implies ` + "`" + `Any` + "`" + ` to apply the rule to all users. - ` + "`" + `id` + "`" + ` - (String) Identifier that uniquely identifies an entity - ` + "`" + `name` + "`" + ` - (String) The configured name of the entity - ` + "`" + `extensions` + "`" + ` - (Map of String)`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `(Optional) You can manually select up to ` + "`" + `8` + "`" + ` groups. When not used it implies ` + "`" + `Any` + "`" + ` to apply the rule to all groups. - ` + "`" + `id` + "`" + ` - (String) Identifier that uniquely identifies an entity - ` + "`" + `name` + "`" + ` - (String) The configured name of the entity - ` + "`" + `extensions` + "`" + ` - (Map of String)`,
				},
				resource.Attribute{
					Name:        "departments",
					Description: `(Optional) Apply to any number of departments When not used it implies ` + "`" + `Any` + "`" + ` to apply the rule to all departments. - ` + "`" + `id` + "`" + ` - (String) Identifier that uniquely identifies an entity - ` + "`" + `name` + "`" + ` - (String) The configured name of the entity - ` + "`" + `extensions` + "`" + ` - (Map of String)`,
				},
				resource.Attribute{
					Name:        "time_windows",
					Description: `(Optional) You can manually select up to ` + "`" + `2` + "`" + ` time intervals. When not used it implies ` + "`" + `always` + "`" + ` to apply the rule to all time intervals. - ` + "`" + `id` + "`" + ` - (String) Identifier that uniquely identifies an entity - ` + "`" + `name` + "`" + ` - (String) The configured name of the entity - ` + "`" + `extensions` + "`" + ` - (Map of String) ` + "`" + `network services` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "nw_service_groups",
					Description: `(Optional) Any number of predefined or custom network service groups to which the rule applies.`,
				},
				resource.Attribute{
					Name:        "nw_application_groups",
					Description: `(Optional) Any number of application groups that you want to control with this rule. The service provides predefined applications that you can group, but not modify`,
				},
				resource.Attribute{
					Name:        "nw_applications",
					Description: `(Optional) When not used it applies the rule to all applications. The service provides predefined applications, which you can group, but not modify. ` + "`" + `source ip addresses` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "src_ip_groups",
					Description: `(Optional) Any number of source IP address groups that you want to control with this rule. - ` + "`" + `id` + "`" + ` - (String) Identifier that uniquely identifies an entity - ` + "`" + `name` + "`" + ` - (String) The configured name of the entity - ` + "`" + `extensions` + "`" + ` - (Map of String)`,
				},
				resource.Attribute{
					Name:        "src_ips",
					Description: `(Optional) You can enter individual IP addresses, subnets, or address ranges. ` + "`" + `destinations` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "app_service_groups",
					Description: `Application service groups on which this rule is applied - ` + "`" + `id` + "`" + ` - (String) Identifier that uniquely identifies an entity - ` + "`" + `name` + "`" + ` - (String) The configured name of the entity - ` + "`" + `extensions` + "`" + ` - (Map of String)`,
				},
				resource.Attribute{
					Name:        "app_services",
					Description: `Application services on which this rule is applied - ` + "`" + `id` + "`" + ` - (String) Identifier that uniquely identifies an entity - ` + "`" + `name` + "`" + ` - (String) The configured name of the entity - ` + "`" + `extensions` + "`" + ` - (Map of String)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) The ID of this resource.`,
				},
				resource.Attribute{
					Name:        "last_modified_time",
					Description: `(Number)`,
				},
				resource.Attribute{
					Name:        "access_control",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "default_rule",
					Description: `(Boolean)`,
				},
				resource.Attribute{
					Name:        "predefined",
					Description: `(Boolean)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Enter additional notes or information. The description cannot exceed 10,240 characters.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Required) Policy rules are evaluated in ascending numerical order (Rule 1 before Rule 2, and so on), and the Rule Order reflects this rule's place in the order.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) An enabled rule is actively enforced. A disabled rule is not actively enforced but does not lose its place in the Rule Order. The service skips it and moves to the next rule.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Choose the action of the service when packets match the rule. The following actions are accepted: ` + "`" + `ALLOW` + "`" + `, ` + "`" + `BLOCK_DROP` + "`" + `, ` + "`" + `BLOCK_RESET` + "`" + `, ` + "`" + `BLOCK_ICMP` + "`" + `, ` + "`" + `EVAL_NWAPP` + "`" + ``,
				},
				resource.Attribute{
					Name:        "rank",
					Description: `(Optional) By default, the admin ranking is disabled. To use this feature, you must enable admin rank. The default value is ` + "`" + `7` + "`" + `. ` + "`" + `Who, Where and When` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "locations",
					Description: `(Optional) You can manually select up to ` + "`" + `8` + "`" + ` locations. When not used it implies ` + "`" + `Any` + "`" + ` to apply the rule to all groups. - ` + "`" + `id` + "`" + ` - (String) Identifier that uniquely identifies an entity - ` + "`" + `name` + "`" + ` - (String) The configured name of the entity - ` + "`" + `extensions` + "`" + ` - (Map of String)`,
				},
				resource.Attribute{
					Name:        "location_groups",
					Description: `(Optional) You can manually select up to ` + "`" + `32` + "`" + ` location groups. When not used it implies ` + "`" + `Any` + "`" + ` to apply the rule to all location groups. - ` + "`" + `id` + "`" + ` - (String) Identifier that uniquely identifies an entity - ` + "`" + `name` + "`" + ` - (String) The configured name of the entity - ` + "`" + `extensions` + "`" + ` - (Map of String)`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `(Optional) You can manually select up to ` + "`" + `4` + "`" + ` general and/or special users. When not used it implies ` + "`" + `Any` + "`" + ` to apply the rule to all users. - ` + "`" + `id` + "`" + ` - (String) Identifier that uniquely identifies an entity - ` + "`" + `name` + "`" + ` - (String) The configured name of the entity - ` + "`" + `extensions` + "`" + ` - (Map of String)`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `(Optional) You can manually select up to ` + "`" + `8` + "`" + ` groups. When not used it implies ` + "`" + `Any` + "`" + ` to apply the rule to all groups. - ` + "`" + `id` + "`" + ` - (String) Identifier that uniquely identifies an entity - ` + "`" + `name` + "`" + ` - (String) The configured name of the entity - ` + "`" + `extensions` + "`" + ` - (Map of String)`,
				},
				resource.Attribute{
					Name:        "departments",
					Description: `(Optional) Apply to any number of departments When not used it implies ` + "`" + `Any` + "`" + ` to apply the rule to all departments. - ` + "`" + `id` + "`" + ` - (String) Identifier that uniquely identifies an entity - ` + "`" + `name` + "`" + ` - (String) The configured name of the entity - ` + "`" + `extensions` + "`" + ` - (Map of String)`,
				},
				resource.Attribute{
					Name:        "time_windows",
					Description: `(Optional) You can manually select up to ` + "`" + `2` + "`" + ` time intervals. When not used it implies ` + "`" + `always` + "`" + ` to apply the rule to all time intervals. - ` + "`" + `id` + "`" + ` - (String) Identifier that uniquely identifies an entity - ` + "`" + `name` + "`" + ` - (String) The configured name of the entity - ` + "`" + `extensions` + "`" + ` - (Map of String) ` + "`" + `network services` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "nw_service_groups",
					Description: `(Optional) Any number of predefined or custom network service groups to which the rule applies.`,
				},
				resource.Attribute{
					Name:        "nw_application_groups",
					Description: `(Optional) Any number of application groups that you want to control with this rule. The service provides predefined applications that you can group, but not modify`,
				},
				resource.Attribute{
					Name:        "nw_applications",
					Description: `(Optional) When not used it applies the rule to all applications. The service provides predefined applications, which you can group, but not modify. ` + "`" + `source ip addresses` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "src_ip_groups",
					Description: `(Optional) Any number of source IP address groups that you want to control with this rule. - ` + "`" + `id` + "`" + ` - (String) Identifier that uniquely identifies an entity - ` + "`" + `name` + "`" + ` - (String) The configured name of the entity - ` + "`" + `extensions` + "`" + ` - (Map of String)`,
				},
				resource.Attribute{
					Name:        "src_ips",
					Description: `(Optional) You can enter individual IP addresses, subnets, or address ranges. ` + "`" + `destinations` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "app_service_groups",
					Description: `Application service groups on which this rule is applied - ` + "`" + `id` + "`" + ` - (String) Identifier that uniquely identifies an entity - ` + "`" + `name` + "`" + ` - (String) The configured name of the entity - ` + "`" + `extensions` + "`" + ` - (Map of String)`,
				},
				resource.Attribute{
					Name:        "app_services",
					Description: `Application services on which this rule is applied - ` + "`" + `id` + "`" + ` - (String) Identifier that uniquely identifies an entity - ` + "`" + `name` + "`" + ` - (String) The configured name of the entity - ` + "`" + `extensions` + "`" + ` - (Map of String)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) The ID of this resource.`,
				},
				resource.Attribute{
					Name:        "last_modified_time",
					Description: `(Number)`,
				},
				resource.Attribute{
					Name:        "access_control",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "default_rule",
					Description: `(Boolean)`,
				},
				resource.Attribute{
					Name:        "predefined",
					Description: `(Boolean)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_firewall_filtering_time_window",
			Category:         "Firewall Policies",
			ShortDescription: `Get information about firewall rule time window.`,
			Description: `

Use the **zia_firewall_filtering_time_window** data source to get information about a time window option available in the Zscaler Internet Access cloud firewall. This data source can then be associated with a ZIA firewall filtering rule.

`,
			Keywords: []string{
				"firewall",
				"policies",
				"filtering",
				"time",
				"window",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the time window to be exported.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the time window resource. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "day_of_week",
					Description: `(String). The supported values are:`,
				},
				resource.Attribute{
					Name:        "ANY",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "NONE",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "EVERYDAY",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "SUN",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "MON",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "TUE",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "WED",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "THU",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "FRI",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "SAT",
					Description: `(String)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "start_time",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "day_of_week",
					Description: `(String). The supported values are:`,
				},
				resource.Attribute{
					Name:        "ANY",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "NONE",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "EVERYDAY",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "SUN",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "MON",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "TUE",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "WED",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "THU",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "FRI",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "SAT",
					Description: `(String)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_location_groups",
			Category:         "Location Management",
			ShortDescription: `Get information about Location Groups.`,
			Description: `

Use the **zia_location_groups** data source to get information about a location group option available in the Zscaler Internet Access.

` + "`" + `` + "`" + `` + "`" + `hcl
# Retrieve ZIA Location Group
data "zia_location_groups" "example"{
    name = "Corporate User Traffic Group"
}
` + "`" + `` + "`" + `` + "`" + `

` + "`" + `` + "`" + `` + "`" + `hcl
# Retrieve ZIA Location Group
data "zia_location_groups" "example"{
    name = "Guest Wifi Group"
}
` + "`" + `` + "`" + `` + "`" + `

` + "`" + `` + "`" + `` + "`" + `hcl
# Retrieve ZIA Location Group
data "zia_location_groups" "example"{
    name = "IoT Traffic Group"
}
` + "`" + `` + "`" + `` + "`" + `

` + "`" + `` + "`" + `` + "`" + `hcl
# Retrieve ZIA Location Group
data "zia_location_groups" "example"{
    name = "Server Traffic Group"
}
` + "`" + `` + "`" + `` + "`" + `

` + "`" + `` + "`" + `` + "`" + `hcl
# Retrieve ZIA Location Group
data "zia_location_groups" "example"{
    name = "Server Traffic Group"
}
` + "`" + `` + "`" + `` + "`" + `

`,
			Keywords: []string{
				"location",
				"management",
				"groups",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Location group name`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Unique identifier for the location group ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(String) Additional information about the location group`,
				},
				resource.Attribute{
					Name:        "deleted",
					Description: `(Boolean) Indicates the location group was deleted`,
				},
				resource.Attribute{
					Name:        "group_type",
					Description: `(String) The location group's type (i.e., Static or Dynamic)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) The ID of this resource.`,
				},
				resource.Attribute{
					Name:        "last_mod_time",
					Description: `(Number) Automatically populated with the current time, after a successful ` + "`" + `POST` + "`" + ` or ` + "`" + `PUT` + "`" + ` request.`,
				},
				resource.Attribute{
					Name:        "predefined",
					Description: `(Boolean)`,
				},
				resource.Attribute{
					Name:        "dynamic_location_group_criteria",
					Description: `(Block Set) Dynamic location group information.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Block List)`,
				},
				resource.Attribute{
					Name:        "match_string",
					Description: `(String) String value to be matched or partially matched`,
				},
				resource.Attribute{
					Name:        "match_type",
					Description: `(String) Operator that performs match action`,
				},
				resource.Attribute{
					Name:        "countries",
					Description: `(List of String) One or more countries from a predefined set`,
				},
				resource.Attribute{
					Name:        "city",
					Description: `(Block List)`,
				},
				resource.Attribute{
					Name:        "match_string",
					Description: `(String) String value to be matched or partially matched`,
				},
				resource.Attribute{
					Name:        "match_type",
					Description: `(String) Operator that performs match action`,
				},
				resource.Attribute{
					Name:        "managed_by",
					Description: `(Block List)`,
				},
				resource.Attribute{
					Name:        "enforce_authentication",
					Description: `(Boolean) Enforce Authentication. Required when ports are enabled, IP Surrogate is enabled, or Kerberos Authentication is enabled.`,
				},
				resource.Attribute{
					Name:        "enforce_aup",
					Description: `(Boolean) Enable AUP. When set to true, AUP is enabled for the location.`,
				},
				resource.Attribute{
					Name:        "enforce_firewall_control",
					Description: `(Boolean) Enable Firewall. When set to true, Firewall is enabled for the location.`,
				},
				resource.Attribute{
					Name:        "enable_xff_forwarding",
					Description: `(Boolean) Enable ` + "`" + `XFF` + "`" + ` Forwarding. When set to true, traffic is passed to Zscaler Cloud via the X-Forwarded-For (XFF) header.`,
				},
				resource.Attribute{
					Name:        "enable_caution",
					Description: `(Boolean) Enable Caution. When set to true, a caution notifcation is enabled for the location.`,
				},
				resource.Attribute{
					Name:        "enable_bandwidth_control",
					Description: `(Boolean) Enable Bandwidth Control. When set to true, Bandwidth Control is enabled for the location.`,
				},
				resource.Attribute{
					Name:        "profiles",
					Description: `(List of String) One or more location profiles from a predefined set`,
				},
				resource.Attribute{
					Name:        "locations",
					Description: `(List of Object) The Name-ID pairs of the locations that are assigned to the static location group. This is ignored if the groupType is Dynamic.`,
				},
				resource.Attribute{
					Name:        "extensions",
					Description: `(Map of String)`,
				},
				resource.Attribute{
					Name:        "last_mod_user",
					Description: `(List of Object)`,
				},
				resource.Attribute{
					Name:        "extensions",
					Description: `(Map of String)`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(List of Object)`,
				},
				resource.Attribute{
					Name:        "last_mod_time",
					Description: `(List of Object) Automatically populated with the current time, after a successful POST or PUT request.`,
				},
				resource.Attribute{
					Name:        "predefined",
					Description: `(Boolean)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "comments",
					Description: `(String) Additional information about the location group`,
				},
				resource.Attribute{
					Name:        "deleted",
					Description: `(Boolean) Indicates the location group was deleted`,
				},
				resource.Attribute{
					Name:        "group_type",
					Description: `(String) The location group's type (i.e., Static or Dynamic)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) The ID of this resource.`,
				},
				resource.Attribute{
					Name:        "last_mod_time",
					Description: `(Number) Automatically populated with the current time, after a successful ` + "`" + `POST` + "`" + ` or ` + "`" + `PUT` + "`" + ` request.`,
				},
				resource.Attribute{
					Name:        "predefined",
					Description: `(Boolean)`,
				},
				resource.Attribute{
					Name:        "dynamic_location_group_criteria",
					Description: `(Block Set) Dynamic location group information.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Block List)`,
				},
				resource.Attribute{
					Name:        "match_string",
					Description: `(String) String value to be matched or partially matched`,
				},
				resource.Attribute{
					Name:        "match_type",
					Description: `(String) Operator that performs match action`,
				},
				resource.Attribute{
					Name:        "countries",
					Description: `(List of String) One or more countries from a predefined set`,
				},
				resource.Attribute{
					Name:        "city",
					Description: `(Block List)`,
				},
				resource.Attribute{
					Name:        "match_string",
					Description: `(String) String value to be matched or partially matched`,
				},
				resource.Attribute{
					Name:        "match_type",
					Description: `(String) Operator that performs match action`,
				},
				resource.Attribute{
					Name:        "managed_by",
					Description: `(Block List)`,
				},
				resource.Attribute{
					Name:        "enforce_authentication",
					Description: `(Boolean) Enforce Authentication. Required when ports are enabled, IP Surrogate is enabled, or Kerberos Authentication is enabled.`,
				},
				resource.Attribute{
					Name:        "enforce_aup",
					Description: `(Boolean) Enable AUP. When set to true, AUP is enabled for the location.`,
				},
				resource.Attribute{
					Name:        "enforce_firewall_control",
					Description: `(Boolean) Enable Firewall. When set to true, Firewall is enabled for the location.`,
				},
				resource.Attribute{
					Name:        "enable_xff_forwarding",
					Description: `(Boolean) Enable ` + "`" + `XFF` + "`" + ` Forwarding. When set to true, traffic is passed to Zscaler Cloud via the X-Forwarded-For (XFF) header.`,
				},
				resource.Attribute{
					Name:        "enable_caution",
					Description: `(Boolean) Enable Caution. When set to true, a caution notifcation is enabled for the location.`,
				},
				resource.Attribute{
					Name:        "enable_bandwidth_control",
					Description: `(Boolean) Enable Bandwidth Control. When set to true, Bandwidth Control is enabled for the location.`,
				},
				resource.Attribute{
					Name:        "profiles",
					Description: `(List of String) One or more location profiles from a predefined set`,
				},
				resource.Attribute{
					Name:        "locations",
					Description: `(List of Object) The Name-ID pairs of the locations that are assigned to the static location group. This is ignored if the groupType is Dynamic.`,
				},
				resource.Attribute{
					Name:        "extensions",
					Description: `(Map of String)`,
				},
				resource.Attribute{
					Name:        "last_mod_user",
					Description: `(List of Object)`,
				},
				resource.Attribute{
					Name:        "extensions",
					Description: `(Map of String)`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(List of Object)`,
				},
				resource.Attribute{
					Name:        "last_mod_time",
					Description: `(List of Object) Automatically populated with the current time, after a successful POST or PUT request.`,
				},
				resource.Attribute{
					Name:        "predefined",
					Description: `(Boolean)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_location_lite",
			Category:         "Location Management",
			ShortDescription: `Get information about Location Lite.`,
			Description: `

Use the **zia_location_lite** data source to get information about a location in lite mode option available in the Zscaler Internet Access. This data source can be used to retrieve the Road Warrior location to then associated with one of the following resources: ` + "`" + `` + "`" + `zia_url_filtering_rules` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `zia_firewall_filtering_rule` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `zia_dlp_web_rules` + "`" + `

` + "`" + `` + "`" + `` + "`" + `hcl
# Retrieve ZIA Location Lite
data "zia_location_lite" "this" {
 name = "Road Warrior"
}
` + "`" + `` + "`" + `` + "`" + `

`,
			Keywords: []string{
				"location",
				"management",
				"lite",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Location group name`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Unique identifier for the location group ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "kerberos_auth",
					Description: `(Boolean)`,
				},
				resource.Attribute{
					Name:        "digest_auth_enabled",
					Description: `(Boolean)`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `(Number) - Parent Location ID. If this ID does not exist or is ` + "`" + `0` + "`" + `, it is implied that it is a parent location. Otherwise, it is a sub-location whose parent has this ID. x-applicableTo: ` + "`" + `SUB` + "`" + ``,
				},
				resource.Attribute{
					Name:        "tz",
					Description: `(String) Timezone of the location. If not specified, it defaults to GMT.`,
				},
				resource.Attribute{
					Name:        "zapp_ssl_scan_enabled",
					Description: `(Boolean) This parameter was deprecated and no longer has an effect on SSL policy. It remains supported in the API payload in order to maintain backwards compatibility with existing scripts, but it will be removed in future.`,
				},
				resource.Attribute{
					Name:        "xff_forward_enabled",
					Description: `(Boolean) Enable XFF Forwarding. When set to true, traffic is passed to Zscaler Cloud via the X-Forwarded-For (XFF) header.`,
				},
				resource.Attribute{
					Name:        "surrogate_ip",
					Description: `(Boolean) Enable Surrogate IP. When set to true, users are mapped to internal device IP addresses.`,
				},
				resource.Attribute{
					Name:        "surrogate_ip_enforced_for_known_browsers",
					Description: `(Boolean) Enforce Surrogate IP for Known Browsers. When set to true, IP Surrogate is enforced for all known browsers.`,
				},
				resource.Attribute{
					Name:        "ofw_enabled",
					Description: `(Boolean) Enable Firewall. When set to true, Firewall is enabled for the location.`,
				},
				resource.Attribute{
					Name:        "ips_control",
					Description: `(Boolean) Enable IPS Control. When set to true, IPS Control is enabled for the location if Firewall is enabled.`,
				},
				resource.Attribute{
					Name:        "aup_enabled",
					Description: `(Boolean) Enable AUP. When set to true, AUP is enabled for the location.`,
				},
				resource.Attribute{
					Name:        "caution_enabled",
					Description: `(Boolean) Enable Caution. When set to true, a caution notifcation is enabled for the location.`,
				},
				resource.Attribute{
					Name:        "aup_block_internet_until_accepted",
					Description: `(Boolean) For First Time AUP Behavior, Block Internet Access. When set, all internet access (including non-HTTP traffic) is disabled until the user accepts the AUP.`,
				},
				resource.Attribute{
					Name:        "aup_force_ssl_inspection",
					Description: `(Boolean) For First Time AUP Behavior, Force SSL Inspection. When set, Zscaler will force SSL Inspection in order to enforce AUP for HTTPS traffic.`,
				},
				resource.Attribute{
					Name:        "ec_location",
					Description: `(Boolean)`,
				},
				resource.Attribute{
					Name:        "other_sub_location",
					Description: `(Boolean) If set to true, indicates that this is a default sub-location created by the Zscaler service to accommodate IPv4 addresses that are not part of any user-defined sub-locations. The default sub-location is created with the name Other and it can be renamed, if required.`,
				},
				resource.Attribute{
					Name:        "other6_sub_location",
					Description: `(Boolean) If set to true, indicates that this is a default sub-location created by the Zscaler service to accommodate IPv6 addresses that are not part of any user-defined sub-locations. The default sub-location is created with the name Other6 and it can be renamed, if required. This field is applicable only if ipv6Enabled is set is true`,
				},
				resource.Attribute{
					Name:        "ipv6_enabled",
					Description: `(Number) If set to true, IPv6 is enabled for the location and IPv6 traffic from the location can be forwarded to the Zscaler service to enforce security policies.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "kerberos_auth",
					Description: `(Boolean)`,
				},
				resource.Attribute{
					Name:        "digest_auth_enabled",
					Description: `(Boolean)`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `(Number) - Parent Location ID. If this ID does not exist or is ` + "`" + `0` + "`" + `, it is implied that it is a parent location. Otherwise, it is a sub-location whose parent has this ID. x-applicableTo: ` + "`" + `SUB` + "`" + ``,
				},
				resource.Attribute{
					Name:        "tz",
					Description: `(String) Timezone of the location. If not specified, it defaults to GMT.`,
				},
				resource.Attribute{
					Name:        "zapp_ssl_scan_enabled",
					Description: `(Boolean) This parameter was deprecated and no longer has an effect on SSL policy. It remains supported in the API payload in order to maintain backwards compatibility with existing scripts, but it will be removed in future.`,
				},
				resource.Attribute{
					Name:        "xff_forward_enabled",
					Description: `(Boolean) Enable XFF Forwarding. When set to true, traffic is passed to Zscaler Cloud via the X-Forwarded-For (XFF) header.`,
				},
				resource.Attribute{
					Name:        "surrogate_ip",
					Description: `(Boolean) Enable Surrogate IP. When set to true, users are mapped to internal device IP addresses.`,
				},
				resource.Attribute{
					Name:        "surrogate_ip_enforced_for_known_browsers",
					Description: `(Boolean) Enforce Surrogate IP for Known Browsers. When set to true, IP Surrogate is enforced for all known browsers.`,
				},
				resource.Attribute{
					Name:        "ofw_enabled",
					Description: `(Boolean) Enable Firewall. When set to true, Firewall is enabled for the location.`,
				},
				resource.Attribute{
					Name:        "ips_control",
					Description: `(Boolean) Enable IPS Control. When set to true, IPS Control is enabled for the location if Firewall is enabled.`,
				},
				resource.Attribute{
					Name:        "aup_enabled",
					Description: `(Boolean) Enable AUP. When set to true, AUP is enabled for the location.`,
				},
				resource.Attribute{
					Name:        "caution_enabled",
					Description: `(Boolean) Enable Caution. When set to true, a caution notifcation is enabled for the location.`,
				},
				resource.Attribute{
					Name:        "aup_block_internet_until_accepted",
					Description: `(Boolean) For First Time AUP Behavior, Block Internet Access. When set, all internet access (including non-HTTP traffic) is disabled until the user accepts the AUP.`,
				},
				resource.Attribute{
					Name:        "aup_force_ssl_inspection",
					Description: `(Boolean) For First Time AUP Behavior, Force SSL Inspection. When set, Zscaler will force SSL Inspection in order to enforce AUP for HTTPS traffic.`,
				},
				resource.Attribute{
					Name:        "ec_location",
					Description: `(Boolean)`,
				},
				resource.Attribute{
					Name:        "other_sub_location",
					Description: `(Boolean) If set to true, indicates that this is a default sub-location created by the Zscaler service to accommodate IPv4 addresses that are not part of any user-defined sub-locations. The default sub-location is created with the name Other and it can be renamed, if required.`,
				},
				resource.Attribute{
					Name:        "other6_sub_location",
					Description: `(Boolean) If set to true, indicates that this is a default sub-location created by the Zscaler service to accommodate IPv6 addresses that are not part of any user-defined sub-locations. The default sub-location is created with the name Other6 and it can be renamed, if required. This field is applicable only if ipv6Enabled is set is true`,
				},
				resource.Attribute{
					Name:        "ipv6_enabled",
					Description: `(Number) If set to true, IPv6 is enabled for the location and IPv6 traffic from the location can be forwarded to the Zscaler service to enforce security policies.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_location_management",
			Category:         "Location Management",
			ShortDescription: `Get information about Location.`,
			Description: `

Use the **zia_location_management** data source to get information about a location resource available in the Zscaler Internet Access Location Management. This resource can then be referenced in multiple other resources, such as URL Filtering Rules, Firewall rules etc.

`,
			Keywords: []string{
				"location",
				"management",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) - The name of the location to be exported.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) - The ID of the location to be exported. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "aup_block_internet_until_accepted",
					Description: `(Boolean) For First Time AUP Behavior, Block Internet Access. When set, all internet access (including non-HTTP traffic) is disabled until the user accepts the AUP.`,
				},
				resource.Attribute{
					Name:        "aup_enabled",
					Description: `(Boolean) Enable AUP. When set to true, AUP is enabled for the location.`,
				},
				resource.Attribute{
					Name:        "aup_force_ssl_inspection",
					Description: `(Boolean) For First Time AUP Behavior, Force SSL Inspection. When set, Zscaler will force SSL Inspection in order to enforce AUP for HTTPS traffic.`,
				},
				resource.Attribute{
					Name:        "aup_timeout_in_days",
					Description: `(Number) Custom AUP Frequency. Refresh time (in days) to re-validate the AUP.`,
				},
				resource.Attribute{
					Name:        "auth_required",
					Description: `(Boolean) Enforce Authentication. Required when ports are enabled, IP Surrogate is enabled, or Kerberos Authentication is enabled.`,
				},
				resource.Attribute{
					Name:        "caution_enabled",
					Description: `(Boolean) Enable Caution. When set to true, a caution notifcation is enabled for the location.`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `(String) Country`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(String) Additional notes or information regarding the location or sub-location. The description cannot exceed 1024 characters.`,
				},
				resource.Attribute{
					Name:        "display_time_unit",
					Description: `(String) Display Time Unit. The time unit to display for IP Surrogate idle time to disassociation.`,
				},
				resource.Attribute{
					Name:        "dn_bandwidth",
					Description: `(Number) Download bandwidth in bytes. The value ` + "`" + `0` + "`" + ` implies no Bandwidth Control enforcement.`,
				},
				resource.Attribute{
					Name:        "up_bandwidth",
					Description: `(Number) Upload bandwidth in bytes. The value ` + "`" + `0` + "`" + ` implies no Bandwidth Control enforcement.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Location ID.`,
				},
				resource.Attribute{
					Name:        "idle_time_in_minutes",
					Description: `(Number) Idle Time to Disassociation. The user mapping idle time (in minutes) is required if a Surrogate IP is enabled.`,
				},
				resource.Attribute{
					Name:        "ip_addresses",
					Description: `(List of String) For locations: IP addresses of the egress points that are provisioned in the Zscaler Cloud. Each entry is a single IP address (e.g., ` + "`" + `238.10.33.9` + "`" + `). For sub-locations: Egress, internal, or GRE tunnel IP addresses. Each entry is either a single IP address, CIDR (e.g., ` + "`" + `10.10.33.0/24` + "`" + `), or range (e.g., ` + "`" + `10.10.33.1-10.10.33.10` + "`" + `)).`,
				},
				resource.Attribute{
					Name:        "ips_control",
					Description: `(Boolean) Enable IPS Control. When set to true, IPS Control is enabled for the location if Firewall is enabled.`,
				},
				resource.Attribute{
					Name:        "ofw_enabled",
					Description: `(Boolean) Enable Firewall. When set to true, Firewall is enabled for the location.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `(Number) - Parent Location ID. If this ID does not exist or is ` + "`" + `0` + "`" + `, it is implied that it is a parent location. Otherwise, it is a sub-location whose parent has this ID. x-applicableTo: ` + "`" + `SUB` + "`" + ``,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `(String) IP ports that are associated with the location.`,
				},
				resource.Attribute{
					Name:        "profile",
					Description: `(String) Profile tag that specifies the location traffic type. If not specified, this tag defaults to ` + "`" + `Unassigned` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ssl_scan_enabled",
					Description: `(Boolean) This parameter was deprecated and no longer has an effect on SSL policy. It remains supported in the API payload in order to maintain backwards compatibility with existing scripts, but it will be removed in future.`,
				},
				resource.Attribute{
					Name:        "surrogate_ip",
					Description: `(Boolean) Enable Surrogate IP. When set to true, users are mapped to internal device IP addresses.`,
				},
				resource.Attribute{
					Name:        "surrogate_ip_enforced_for_known_browsers",
					Description: `(Boolean) Enforce Surrogate IP for Known Browsers. When set to true, IP Surrogate is enforced for all known browsers.`,
				},
				resource.Attribute{
					Name:        "surrogate_refresh_time_in_minutes",
					Description: `(Number) Refresh Time for re-validation of Surrogacy. The surrogate refresh time (in minutes) to re-validate the IP surrogates.`,
				},
				resource.Attribute{
					Name:        "surrogate_refresh_time_unit",
					Description: `(String) Display Refresh Time Unit. The time unit to display for refresh time for re-validation of surrogacy.`,
				},
				resource.Attribute{
					Name:        "tz",
					Description: `(String) Timezone of the location. If not specified, it defaults to GMT.`,
				},
				resource.Attribute{
					Name:        "xff_forward_enabled",
					Description: `(Boolean) Enable XFF Forwarding. When set to true, traffic is passed to Zscaler Cloud via the X-Forwarded-For (XFF) header.`,
				},
				resource.Attribute{
					Name:        "zapp_ssl_scan_enabled",
					Description: `(Boolean) This parameter was deprecated and no longer has an effect on SSL policy. It remains supported in the API payload in order to maintain backwards compatibility with existing scripts, but it will be removed in future.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(String) Additional information about this VPN credential. Additional information about this VPN credential.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `(String) Fully Qualified Domain Name. Applicable only to ` + "`" + `UFQDN` + "`" + ` or ` + "`" + `XAUTH` + "`" + ` (or ` + "`" + `HOSTED_MOBILE_USERS` + "`" + `) auth type.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) VPN credential id`,
				},
				resource.Attribute{
					Name:        "pre_shared_key",
					Description: `(String) Pre-shared key. This is a required field for ` + "`" + `UFQDN` + "`" + ` and IP auth type.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(String) VPN authentication type (i.e., how the VPN credential is sent to the server). It is not modifiable after VpnCredential is created.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(List of Object)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(String) The configured name of the entity`,
				},
				resource.Attribute{
					Name:        "extensions",
					Description: `(Map of String)`,
				},
				resource.Attribute{
					Name:        "managed_by",
					Description: `(List of Object)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(String) The configured name of the entity`,
				},
				resource.Attribute{
					Name:        "extensions",
					Description: `(Map of String)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "aup_block_internet_until_accepted",
					Description: `(Boolean) For First Time AUP Behavior, Block Internet Access. When set, all internet access (including non-HTTP traffic) is disabled until the user accepts the AUP.`,
				},
				resource.Attribute{
					Name:        "aup_enabled",
					Description: `(Boolean) Enable AUP. When set to true, AUP is enabled for the location.`,
				},
				resource.Attribute{
					Name:        "aup_force_ssl_inspection",
					Description: `(Boolean) For First Time AUP Behavior, Force SSL Inspection. When set, Zscaler will force SSL Inspection in order to enforce AUP for HTTPS traffic.`,
				},
				resource.Attribute{
					Name:        "aup_timeout_in_days",
					Description: `(Number) Custom AUP Frequency. Refresh time (in days) to re-validate the AUP.`,
				},
				resource.Attribute{
					Name:        "auth_required",
					Description: `(Boolean) Enforce Authentication. Required when ports are enabled, IP Surrogate is enabled, or Kerberos Authentication is enabled.`,
				},
				resource.Attribute{
					Name:        "caution_enabled",
					Description: `(Boolean) Enable Caution. When set to true, a caution notifcation is enabled for the location.`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `(String) Country`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(String) Additional notes or information regarding the location or sub-location. The description cannot exceed 1024 characters.`,
				},
				resource.Attribute{
					Name:        "display_time_unit",
					Description: `(String) Display Time Unit. The time unit to display for IP Surrogate idle time to disassociation.`,
				},
				resource.Attribute{
					Name:        "dn_bandwidth",
					Description: `(Number) Download bandwidth in bytes. The value ` + "`" + `0` + "`" + ` implies no Bandwidth Control enforcement.`,
				},
				resource.Attribute{
					Name:        "up_bandwidth",
					Description: `(Number) Upload bandwidth in bytes. The value ` + "`" + `0` + "`" + ` implies no Bandwidth Control enforcement.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Location ID.`,
				},
				resource.Attribute{
					Name:        "idle_time_in_minutes",
					Description: `(Number) Idle Time to Disassociation. The user mapping idle time (in minutes) is required if a Surrogate IP is enabled.`,
				},
				resource.Attribute{
					Name:        "ip_addresses",
					Description: `(List of String) For locations: IP addresses of the egress points that are provisioned in the Zscaler Cloud. Each entry is a single IP address (e.g., ` + "`" + `238.10.33.9` + "`" + `). For sub-locations: Egress, internal, or GRE tunnel IP addresses. Each entry is either a single IP address, CIDR (e.g., ` + "`" + `10.10.33.0/24` + "`" + `), or range (e.g., ` + "`" + `10.10.33.1-10.10.33.10` + "`" + `)).`,
				},
				resource.Attribute{
					Name:        "ips_control",
					Description: `(Boolean) Enable IPS Control. When set to true, IPS Control is enabled for the location if Firewall is enabled.`,
				},
				resource.Attribute{
					Name:        "ofw_enabled",
					Description: `(Boolean) Enable Firewall. When set to true, Firewall is enabled for the location.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `(Number) - Parent Location ID. If this ID does not exist or is ` + "`" + `0` + "`" + `, it is implied that it is a parent location. Otherwise, it is a sub-location whose parent has this ID. x-applicableTo: ` + "`" + `SUB` + "`" + ``,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `(String) IP ports that are associated with the location.`,
				},
				resource.Attribute{
					Name:        "profile",
					Description: `(String) Profile tag that specifies the location traffic type. If not specified, this tag defaults to ` + "`" + `Unassigned` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ssl_scan_enabled",
					Description: `(Boolean) This parameter was deprecated and no longer has an effect on SSL policy. It remains supported in the API payload in order to maintain backwards compatibility with existing scripts, but it will be removed in future.`,
				},
				resource.Attribute{
					Name:        "surrogate_ip",
					Description: `(Boolean) Enable Surrogate IP. When set to true, users are mapped to internal device IP addresses.`,
				},
				resource.Attribute{
					Name:        "surrogate_ip_enforced_for_known_browsers",
					Description: `(Boolean) Enforce Surrogate IP for Known Browsers. When set to true, IP Surrogate is enforced for all known browsers.`,
				},
				resource.Attribute{
					Name:        "surrogate_refresh_time_in_minutes",
					Description: `(Number) Refresh Time for re-validation of Surrogacy. The surrogate refresh time (in minutes) to re-validate the IP surrogates.`,
				},
				resource.Attribute{
					Name:        "surrogate_refresh_time_unit",
					Description: `(String) Display Refresh Time Unit. The time unit to display for refresh time for re-validation of surrogacy.`,
				},
				resource.Attribute{
					Name:        "tz",
					Description: `(String) Timezone of the location. If not specified, it defaults to GMT.`,
				},
				resource.Attribute{
					Name:        "xff_forward_enabled",
					Description: `(Boolean) Enable XFF Forwarding. When set to true, traffic is passed to Zscaler Cloud via the X-Forwarded-For (XFF) header.`,
				},
				resource.Attribute{
					Name:        "zapp_ssl_scan_enabled",
					Description: `(Boolean) This parameter was deprecated and no longer has an effect on SSL policy. It remains supported in the API payload in order to maintain backwards compatibility with existing scripts, but it will be removed in future.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(String) Additional information about this VPN credential. Additional information about this VPN credential.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `(String) Fully Qualified Domain Name. Applicable only to ` + "`" + `UFQDN` + "`" + ` or ` + "`" + `XAUTH` + "`" + ` (or ` + "`" + `HOSTED_MOBILE_USERS` + "`" + `) auth type.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) VPN credential id`,
				},
				resource.Attribute{
					Name:        "pre_shared_key",
					Description: `(String) Pre-shared key. This is a required field for ` + "`" + `UFQDN` + "`" + ` and IP auth type.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(String) VPN authentication type (i.e., how the VPN credential is sent to the server). It is not modifiable after VpnCredential is created.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(List of Object)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(String) The configured name of the entity`,
				},
				resource.Attribute{
					Name:        "extensions",
					Description: `(Map of String)`,
				},
				resource.Attribute{
					Name:        "managed_by",
					Description: `(List of Object)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(String) The configured name of the entity`,
				},
				resource.Attribute{
					Name:        "extensions",
					Description: `(Map of String)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_rule_labels",
			Category:         "Rule Labels",
			ShortDescription: `Get information about rule labels details.`,
			Description: `

Use the **zia_rule_labels** data source to get information about a rule label resource in the Zscaler Internet Access cloud or via the API. This data source can then be associated with resources such as: Firewall Rules and URL filtering rules

`,
			Keywords: []string{
				"rule",
				"labels",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the rule label to be exported.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(String) The unique identifer for the device group. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(String) The rule label description.`,
				},
				resource.Attribute{
					Name:        "last_modified_time",
					Description: `(String) Timestamp when the rule lable was last modified. This is a read-only field. Ignored by PUT and DELETE requests.`,
				},
				resource.Attribute{
					Name:        "last_modified_by",
					Description: `(String) The admin that modified the rule label last. This is a read-only field. Ignored by PUT requests.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `(String) The admin that created the rule label. This is a read-only field. Ignored by PUT requests.`,
				},
				resource.Attribute{
					Name:        "referenced_rule_count",
					Description: `(int) The number of rules that reference the label.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(String) The rule label description.`,
				},
				resource.Attribute{
					Name:        "last_modified_time",
					Description: `(String) Timestamp when the rule lable was last modified. This is a read-only field. Ignored by PUT and DELETE requests.`,
				},
				resource.Attribute{
					Name:        "last_modified_by",
					Description: `(String) The admin that modified the rule label last. This is a read-only field. Ignored by PUT requests.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `(String) The admin that created the rule label. This is a read-only field. Ignored by PUT requests.`,
				},
				resource.Attribute{
					Name:        "referenced_rule_count",
					Description: `(int) The number of rules that reference the label.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_security_policy_settings",
			Category:         "Security Policy Settings",
			ShortDescription: `Gets a list of URLs that are on the allow and denylist.`,
			Description: `

Use the **zia_security_policy_settings** data source to get a list of URLs that were added to the allow and denylist under the Advanced Threat Protection policy in the Zscaler Internet Access cloud or via the API.

`,
			Keywords: []string{
				"security",
				"policy",
				"settings",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_sub_location_management",
			Category:         "Location Management",
			ShortDescription: `Get information about Sublocation.`,
			Description: `

Use the **zia_location_management** data source to get information about a sublocation resource available in the Zscaler Internet Access Location Management. This resource can then be referenced in multiple other resources, such as URL Filtering Rules, Firewall rules etc.

`,
			Keywords: []string{
				"location",
				"management",
				"sub",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) - The name of the location to be exported.`,
				},
				resource.Attribute{
					Name:        "parent_name",
					Description: `(Required) - The name of the parent location Note: If the ` + "`" + `` + "`" + `parent_name` + "`" + `` + "`" + ` attribute is not provided, the provider will return the first location which matches the name being queried. This attribute is specially required in case of overlapping sublocation names exist across multiple parent locations. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "aup_block_internet_until_accepted",
					Description: `(Boolean) For First Time AUP Behavior, Block Internet Access. When set, all internet access (including non-HTTP traffic) is disabled until the user accepts the AUP.`,
				},
				resource.Attribute{
					Name:        "aup_enabled",
					Description: `(Boolean) Enable AUP. When set to true, AUP is enabled for the location.`,
				},
				resource.Attribute{
					Name:        "aup_force_ssl_inspection",
					Description: `(Boolean) For First Time AUP Behavior, Force SSL Inspection. When set, Zscaler will force SSL Inspection in order to enforce AUP for HTTPS traffic.`,
				},
				resource.Attribute{
					Name:        "aup_timeout_in_days",
					Description: `(Number) Custom AUP Frequency. Refresh time (in days) to re-validate the AUP.`,
				},
				resource.Attribute{
					Name:        "auth_required",
					Description: `(Boolean) Enforce Authentication. Required when ports are enabled, IP Surrogate is enabled, or Kerberos Authentication is enabled.`,
				},
				resource.Attribute{
					Name:        "caution_enabled",
					Description: `(Boolean) Enable Caution. When set to true, a caution notifcation is enabled for the location.`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `(String) Country`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(String) Additional notes or information regarding the location or sub-location. The description cannot exceed 1024 characters.`,
				},
				resource.Attribute{
					Name:        "display_time_unit",
					Description: `(String) Display Time Unit. The time unit to display for IP Surrogate idle time to disassociation.`,
				},
				resource.Attribute{
					Name:        "dn_bandwidth",
					Description: `(Number) Download bandwidth in bytes. The value ` + "`" + `0` + "`" + ` implies no Bandwidth Control enforcement.`,
				},
				resource.Attribute{
					Name:        "up_bandwidth",
					Description: `(Number) Upload bandwidth in bytes. The value ` + "`" + `0` + "`" + ` implies no Bandwidth Control enforcement.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Location ID.`,
				},
				resource.Attribute{
					Name:        "idle_time_in_minutes",
					Description: `(Number) Idle Time to Disassociation. The user mapping idle time (in minutes) is required if a Surrogate IP is enabled.`,
				},
				resource.Attribute{
					Name:        "ip_addresses",
					Description: `(List of String) For locations: IP addresses of the egress points that are provisioned in the Zscaler Cloud. Each entry is a single IP address (e.g., ` + "`" + `238.10.33.9` + "`" + `). For sub-locations: Egress, internal, or GRE tunnel IP addresses. Each entry is either a single IP address, CIDR (e.g., ` + "`" + `10.10.33.0/24` + "`" + `), or range (e.g., ` + "`" + `10.10.33.1-10.10.33.10` + "`" + `)).`,
				},
				resource.Attribute{
					Name:        "ips_control",
					Description: `(Boolean) Enable IPS Control. When set to true, IPS Control is enabled for the location if Firewall is enabled.`,
				},
				resource.Attribute{
					Name:        "ofw_enabled",
					Description: `(Boolean) Enable Firewall. When set to true, Firewall is enabled for the location.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `(Number) - Parent Location ID. If this ID does not exist or is ` + "`" + `0` + "`" + `, it is implied that it is a parent location. Otherwise, it is a sub-location whose parent has this ID. x-applicableTo: ` + "`" + `SUB` + "`" + ``,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `(String) IP ports that are associated with the location.`,
				},
				resource.Attribute{
					Name:        "profile",
					Description: `(String) Profile tag that specifies the location traffic type. If not specified, this tag defaults to ` + "`" + `Unassigned` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ssl_scan_enabled",
					Description: `(Boolean) This parameter was deprecated and no longer has an effect on SSL policy. It remains supported in the API payload in order to maintain backwards compatibility with existing scripts, but it will be removed in future.`,
				},
				resource.Attribute{
					Name:        "surrogate_ip",
					Description: `(Boolean) Enable Surrogate IP. When set to true, users are mapped to internal device IP addresses.`,
				},
				resource.Attribute{
					Name:        "surrogate_ip_enforced_for_known_browsers",
					Description: `(Boolean) Enforce Surrogate IP for Known Browsers. When set to true, IP Surrogate is enforced for all known browsers.`,
				},
				resource.Attribute{
					Name:        "surrogate_refresh_time_in_minutes",
					Description: `(Number) Refresh Time for re-validation of Surrogacy. The surrogate refresh time (in minutes) to re-validate the IP surrogates.`,
				},
				resource.Attribute{
					Name:        "surrogate_refresh_time_unit",
					Description: `(String) Display Refresh Time Unit. The time unit to display for refresh time for re-validation of surrogacy.`,
				},
				resource.Attribute{
					Name:        "tz",
					Description: `(String) Timezone of the location. If not specified, it defaults to GMT.`,
				},
				resource.Attribute{
					Name:        "xff_forward_enabled",
					Description: `(Boolean) Enable XFF Forwarding. When set to true, traffic is passed to Zscaler Cloud via the X-Forwarded-For (XFF) header.`,
				},
				resource.Attribute{
					Name:        "zapp_ssl_scan_enabled",
					Description: `(Boolean) This parameter was deprecated and no longer has an effect on SSL policy. It remains supported in the API payload in order to maintain backwards compatibility with existing scripts, but it will be removed in future.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(String) Additional information about this VPN credential. Additional information about this VPN credential.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `(String) Fully Qualified Domain Name. Applicable only to ` + "`" + `UFQDN` + "`" + ` or ` + "`" + `XAUTH` + "`" + ` (or ` + "`" + `HOSTED_MOBILE_USERS` + "`" + `) auth type.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) VPN credential id`,
				},
				resource.Attribute{
					Name:        "pre_shared_key",
					Description: `(String) Pre-shared key. This is a required field for ` + "`" + `UFQDN` + "`" + ` and IP auth type.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(String) VPN authentication type (i.e., how the VPN credential is sent to the server). It is not modifiable after VpnCredential is created.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(List of Object)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(String) The configured name of the entity`,
				},
				resource.Attribute{
					Name:        "extensions",
					Description: `(Map of String)`,
				},
				resource.Attribute{
					Name:        "managed_by",
					Description: `(List of Object)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(String) The configured name of the entity`,
				},
				resource.Attribute{
					Name:        "extensions",
					Description: `(Map of String)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "aup_block_internet_until_accepted",
					Description: `(Boolean) For First Time AUP Behavior, Block Internet Access. When set, all internet access (including non-HTTP traffic) is disabled until the user accepts the AUP.`,
				},
				resource.Attribute{
					Name:        "aup_enabled",
					Description: `(Boolean) Enable AUP. When set to true, AUP is enabled for the location.`,
				},
				resource.Attribute{
					Name:        "aup_force_ssl_inspection",
					Description: `(Boolean) For First Time AUP Behavior, Force SSL Inspection. When set, Zscaler will force SSL Inspection in order to enforce AUP for HTTPS traffic.`,
				},
				resource.Attribute{
					Name:        "aup_timeout_in_days",
					Description: `(Number) Custom AUP Frequency. Refresh time (in days) to re-validate the AUP.`,
				},
				resource.Attribute{
					Name:        "auth_required",
					Description: `(Boolean) Enforce Authentication. Required when ports are enabled, IP Surrogate is enabled, or Kerberos Authentication is enabled.`,
				},
				resource.Attribute{
					Name:        "caution_enabled",
					Description: `(Boolean) Enable Caution. When set to true, a caution notifcation is enabled for the location.`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `(String) Country`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(String) Additional notes or information regarding the location or sub-location. The description cannot exceed 1024 characters.`,
				},
				resource.Attribute{
					Name:        "display_time_unit",
					Description: `(String) Display Time Unit. The time unit to display for IP Surrogate idle time to disassociation.`,
				},
				resource.Attribute{
					Name:        "dn_bandwidth",
					Description: `(Number) Download bandwidth in bytes. The value ` + "`" + `0` + "`" + ` implies no Bandwidth Control enforcement.`,
				},
				resource.Attribute{
					Name:        "up_bandwidth",
					Description: `(Number) Upload bandwidth in bytes. The value ` + "`" + `0` + "`" + ` implies no Bandwidth Control enforcement.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Location ID.`,
				},
				resource.Attribute{
					Name:        "idle_time_in_minutes",
					Description: `(Number) Idle Time to Disassociation. The user mapping idle time (in minutes) is required if a Surrogate IP is enabled.`,
				},
				resource.Attribute{
					Name:        "ip_addresses",
					Description: `(List of String) For locations: IP addresses of the egress points that are provisioned in the Zscaler Cloud. Each entry is a single IP address (e.g., ` + "`" + `238.10.33.9` + "`" + `). For sub-locations: Egress, internal, or GRE tunnel IP addresses. Each entry is either a single IP address, CIDR (e.g., ` + "`" + `10.10.33.0/24` + "`" + `), or range (e.g., ` + "`" + `10.10.33.1-10.10.33.10` + "`" + `)).`,
				},
				resource.Attribute{
					Name:        "ips_control",
					Description: `(Boolean) Enable IPS Control. When set to true, IPS Control is enabled for the location if Firewall is enabled.`,
				},
				resource.Attribute{
					Name:        "ofw_enabled",
					Description: `(Boolean) Enable Firewall. When set to true, Firewall is enabled for the location.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `(Number) - Parent Location ID. If this ID does not exist or is ` + "`" + `0` + "`" + `, it is implied that it is a parent location. Otherwise, it is a sub-location whose parent has this ID. x-applicableTo: ` + "`" + `SUB` + "`" + ``,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `(String) IP ports that are associated with the location.`,
				},
				resource.Attribute{
					Name:        "profile",
					Description: `(String) Profile tag that specifies the location traffic type. If not specified, this tag defaults to ` + "`" + `Unassigned` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ssl_scan_enabled",
					Description: `(Boolean) This parameter was deprecated and no longer has an effect on SSL policy. It remains supported in the API payload in order to maintain backwards compatibility with existing scripts, but it will be removed in future.`,
				},
				resource.Attribute{
					Name:        "surrogate_ip",
					Description: `(Boolean) Enable Surrogate IP. When set to true, users are mapped to internal device IP addresses.`,
				},
				resource.Attribute{
					Name:        "surrogate_ip_enforced_for_known_browsers",
					Description: `(Boolean) Enforce Surrogate IP for Known Browsers. When set to true, IP Surrogate is enforced for all known browsers.`,
				},
				resource.Attribute{
					Name:        "surrogate_refresh_time_in_minutes",
					Description: `(Number) Refresh Time for re-validation of Surrogacy. The surrogate refresh time (in minutes) to re-validate the IP surrogates.`,
				},
				resource.Attribute{
					Name:        "surrogate_refresh_time_unit",
					Description: `(String) Display Refresh Time Unit. The time unit to display for refresh time for re-validation of surrogacy.`,
				},
				resource.Attribute{
					Name:        "tz",
					Description: `(String) Timezone of the location. If not specified, it defaults to GMT.`,
				},
				resource.Attribute{
					Name:        "xff_forward_enabled",
					Description: `(Boolean) Enable XFF Forwarding. When set to true, traffic is passed to Zscaler Cloud via the X-Forwarded-For (XFF) header.`,
				},
				resource.Attribute{
					Name:        "zapp_ssl_scan_enabled",
					Description: `(Boolean) This parameter was deprecated and no longer has an effect on SSL policy. It remains supported in the API payload in order to maintain backwards compatibility with existing scripts, but it will be removed in future.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(String) Additional information about this VPN credential. Additional information about this VPN credential.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `(String) Fully Qualified Domain Name. Applicable only to ` + "`" + `UFQDN` + "`" + ` or ` + "`" + `XAUTH` + "`" + ` (or ` + "`" + `HOSTED_MOBILE_USERS` + "`" + `) auth type.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) VPN credential id`,
				},
				resource.Attribute{
					Name:        "pre_shared_key",
					Description: `(String) Pre-shared key. This is a required field for ` + "`" + `UFQDN` + "`" + ` and IP auth type.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(String) VPN authentication type (i.e., how the VPN credential is sent to the server). It is not modifiable after VpnCredential is created.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(List of Object)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(String) The configured name of the entity`,
				},
				resource.Attribute{
					Name:        "extensions",
					Description: `(Map of String)`,
				},
				resource.Attribute{
					Name:        "managed_by",
					Description: `(List of Object)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(String) The configured name of the entity`,
				},
				resource.Attribute{
					Name:        "extensions",
					Description: `(Map of String)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_traffic_forwarding_gre_internal_ip_range_list",
			Category:         "Traffic Forwarding",
			ShortDescription: `Gets the next available GRE tunnel internal IP address ranges.`,
			Description: `

Use the **zia_gre_internal_ip_range_list** data source to get information about the next available GRE tunnel internal ip ranges for the purposes of GRE tunnel creation in the Zscaler Internet Access when the ` + "`" + `ip_unnumbered` + "`" + ` parameter is set to ` + "`" + `false` + "`" + `

`,
			Keywords: []string{
				"traffic",
				"forwarding",
				"gre",
				"internal",
				"ip",
				"range",
				"list",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "required_count",
					Description: `(Required) ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "end_ip_address",
					Description: `(String) Starting IP address in the range`,
				},
				resource.Attribute{
					Name:        "start_ip_address",
					Description: `(String) Ending IP address in the range`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "end_ip_address",
					Description: `(String) Starting IP address in the range`,
				},
				resource.Attribute{
					Name:        "start_ip_address",
					Description: `(String) Ending IP address in the range`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_traffic_forwarding_gre_tunnel_info",
			Category:         "Traffic Forwarding",
			ShortDescription: `Gets a list of IP addresses with GRE tunnel details.`,
			Description: `

The **zia_traffic_forwarding_gre_tunnel_info** data source to get information about provisioned GRE tunnel information created in the Zscaler Internet Access portal.

`,
			Keywords: []string{
				"traffic",
				"forwarding",
				"gre",
				"tunnel",
				"info",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Required) - Filter based on an IP address range.`,
				},
				resource.Attribute{
					Name:        "gre_enabled",
					Description: `(Optional) - Displays only ip addresses with GRE tunnel enabled ->`,
				},
				resource.Attribute{
					Name:        "gre_tunnel_ip",
					Description: `(String) The start of the internal IP address in /29 CIDR range`,
				},
				resource.Attribute{
					Name:        "primary_gw",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "secondary_gw",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "tun_id",
					Description: `(Number)`,
				},
				resource.Attribute{
					Name:        "gre_range_primary",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "gre_range_secondary",
					Description: `(String)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "gre_tunnel_ip",
					Description: `(String) The start of the internal IP address in /29 CIDR range`,
				},
				resource.Attribute{
					Name:        "primary_gw",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "secondary_gw",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "tun_id",
					Description: `(Number)`,
				},
				resource.Attribute{
					Name:        "gre_range_primary",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "gre_range_secondary",
					Description: `(String)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_traffic_forwarding_gre_vip_recommended_list",
			Category:         "Traffic Forwarding",
			ShortDescription: `Gets a list of recommended GRE tunnel virtual IP addresses (VIPs), based on source IP address or latitude/longitude coordinates.`,
			Description: `

Use the **zia_gre_vip_recommended_list** data source to get information about a list of recommended GRE tunnel virtual IP addresses (VIPs), based on source IP address or latitude/longitude coordinates.

`,
			Keywords: []string{
				"traffic",
				"forwarding",
				"gre",
				"vip",
				"recommended",
				"list",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "source_ip",
					Description: `(Required) - Filter based on an IP address range.`,
				},
				resource.Attribute{
					Name:        "required_count",
					Description: `(Optional) - Number of IP address to be exported.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Unique identifer of the GRE virtual IP address (VIP) ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "source_ip",
					Description: `(String) The public source IP address.`,
				},
				resource.Attribute{
					Name:        "virtual_ip",
					Description: `(String) GRE cluster virtual IP address (VIP)`,
				},
				resource.Attribute{
					Name:        "private_service_edge",
					Description: `(Boolean) Set to true if the virtual IP address (VIP) is a ZIA Private Service Edge`,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `(String) Data center information`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "source_ip",
					Description: `(String) The public source IP address.`,
				},
				resource.Attribute{
					Name:        "virtual_ip",
					Description: `(String) GRE cluster virtual IP address (VIP)`,
				},
				resource.Attribute{
					Name:        "private_service_edge",
					Description: `(Boolean) Set to true if the virtual IP address (VIP) is a ZIA Private Service Edge`,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `(String) Data center information`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_traffic_forwarding_public_node_vips",
			Category:         "Traffic Forwarding",
			ShortDescription: `Gets a paginated list of the virtual IP addresses (VIPs) available in the Zscaler cloud`,
			Description: `

Use the **zia_traffic_forwarding_public_node_vips** data source to retrieve a paginated list of virtual IP addresses (VIPs) available in the Zscaler cloud.

`,
			Keywords: []string{
				"traffic",
				"forwarding",
				"public",
				"node",
				"vips",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_name",
					Description: `(String) Cloud Name`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(String) Region`,
				},
				resource.Attribute{
					Name:        "city",
					Description: `(String) City`,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `(String) Data-center Name`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(String) Location Coordinates`,
				},
				resource.Attribute{
					Name:        "vpn_ips",
					Description: `(List of String) VPN IPs`,
				},
				resource.Attribute{
					Name:        "vpn_domain_name",
					Description: `(String) VPN Server DN`,
				},
				resource.Attribute{
					Name:        "gre_ips",
					Description: `(List of String) GRE IPs`,
				},
				resource.Attribute{
					Name:        "gre_domain_name",
					Description: `(String) Proxy Host Name`,
				},
				resource.Attribute{
					Name:        "pac_ips",
					Description: `(List of String) Pac IPs`,
				},
				resource.Attribute{
					Name:        "pac_domain_name",
					Description: `(String) Pac Server DN`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_traffic_forwarding_static_ip",
			Category:         "Traffic Forwarding",
			ShortDescription: `Gets static IP address for the specified ID`,
			Description: `

Use the **zia_traffic_forwarding_static_ip** data source to get information about all provisioned static IP addresses. This resource can then be utilized when creating a GRE Tunnel or VPN Credential resource of Type ` + "`" + `IP` + "`" + `

`,
			Keywords: []string{
				"traffic",
				"forwarding",
				"static",
				"ip",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Required) The static IP address`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The unique identifier for the static IP address ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) The unique identifier for the static IP address`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(String) The static IP address`,
				},
				resource.Attribute{
					Name:        "geo_override",
					Description: `(Boolean) If not set, geographic coordinates and city are automatically determined from the IP address. Otherwise, the latitude and longitude coordinates must be provided.`,
				},
				resource.Attribute{
					Name:        "latitude",
					Description: `(Number) Required only if the geoOverride attribute is set. Latitude with 7 digit precision after decimal point, ranges between ` + "`" + `-90` + "`" + ` and ` + "`" + `90` + "`" + ` degrees.`,
				},
				resource.Attribute{
					Name:        "longitude",
					Description: `(Number) Required only if the geoOverride attribute is set. Longitude with 7 digit precision after decimal point, ranges between ` + "`" + `-180` + "`" + ` and ` + "`" + `180` + "`" + ` degrees.`,
				},
				resource.Attribute{
					Name:        "routable_ip",
					Description: `(Boolean) Indicates whether a non-RFC 1918 IP address is publicly routable. This attribute is ignored if there is no ZIA Private Service Edge associated to the organization.`,
				},
				resource.Attribute{
					Name:        "last_modification_time",
					Description: `(Number) When the static IP address was last modified`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(String) Additional information about this static IP address`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Number) The unique identifier for the static IP address`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(String) The static IP address`,
				},
				resource.Attribute{
					Name:        "geo_override",
					Description: `(Boolean) If not set, geographic coordinates and city are automatically determined from the IP address. Otherwise, the latitude and longitude coordinates must be provided.`,
				},
				resource.Attribute{
					Name:        "latitude",
					Description: `(Number) Required only if the geoOverride attribute is set. Latitude with 7 digit precision after decimal point, ranges between ` + "`" + `-90` + "`" + ` and ` + "`" + `90` + "`" + ` degrees.`,
				},
				resource.Attribute{
					Name:        "longitude",
					Description: `(Number) Required only if the geoOverride attribute is set. Longitude with 7 digit precision after decimal point, ranges between ` + "`" + `-180` + "`" + ` and ` + "`" + `180` + "`" + ` degrees.`,
				},
				resource.Attribute{
					Name:        "routable_ip",
					Description: `(Boolean) Indicates whether a non-RFC 1918 IP address is publicly routable. This attribute is ignored if there is no ZIA Private Service Edge associated to the organization.`,
				},
				resource.Attribute{
					Name:        "last_modification_time",
					Description: `(Number) When the static IP address was last modified`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(String) Additional information about this static IP address`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_traffic_forwarding_vpn_credentials",
			Category:         "Traffic Forwarding",
			ShortDescription: `Gets VPN credentials that can be associated to locations.`,
			Description:      ``,
			Keywords: []string{
				"traffic",
				"forwarding",
				"vpn",
				"credentials",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Required) - Filter based on an IP address range.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Unique identifer of the GRE virtual IP address (VIP) ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) VPN credential id`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(String) VPN authentication type (i.e., how the VPN credential is sent to the server). It is not modifiable after VpnCredential is created.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `(String) Fully Qualified Domain Name. Applicable only to ` + "`" + `UFQDN` + "`" + ` or ` + "`" + `XAUTH` + "`" + ` (or ` + "`" + `HOSTED_MOBILE_USERS` + "`" + `) auth type.`,
				},
				resource.Attribute{
					Name:        "pre_shared_key",
					Description: `(String) Pre-shared key. This is a required field for ` + "`" + `UFQDN` + "`" + ` and ` + "`" + `IP` + "`" + ` auth type.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(String) Additional information about this VPN credential.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Set of Object) Location that is associated to this VPN credential. Non-existence means not associated to any location.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(String) The configured name of the entity`,
				},
				resource.Attribute{
					Name:        "extensions",
					Description: `(Map of String)`,
				},
				resource.Attribute{
					Name:        "managed_by",
					Description: `(Set of Object) SD-WAN Partner that manages the location. If a partner does not manage the locaton, this is set to Self.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(String) The configured name of the entity`,
				},
				resource.Attribute{
					Name:        "extensions",
					Description: `(Map of String) ->`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Number) VPN credential id`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(String) VPN authentication type (i.e., how the VPN credential is sent to the server). It is not modifiable after VpnCredential is created.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `(String) Fully Qualified Domain Name. Applicable only to ` + "`" + `UFQDN` + "`" + ` or ` + "`" + `XAUTH` + "`" + ` (or ` + "`" + `HOSTED_MOBILE_USERS` + "`" + `) auth type.`,
				},
				resource.Attribute{
					Name:        "pre_shared_key",
					Description: `(String) Pre-shared key. This is a required field for ` + "`" + `UFQDN` + "`" + ` and ` + "`" + `IP` + "`" + ` auth type.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(String) Additional information about this VPN credential.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Set of Object) Location that is associated to this VPN credential. Non-existence means not associated to any location.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(String) The configured name of the entity`,
				},
				resource.Attribute{
					Name:        "extensions",
					Description: `(Map of String)`,
				},
				resource.Attribute{
					Name:        "managed_by",
					Description: `(Set of Object) SD-WAN Partner that manages the location. If a partner does not manage the locaton, this is set to Self.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(String) The configured name of the entity`,
				},
				resource.Attribute{
					Name:        "extensions",
					Description: `(Map of String) ->`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_url_categories",
			Category:         "URL Categories",
			ShortDescription: `Gets information about all or custom URL categories. By default, the response includes keywords.`,
			Description: `

Use the **zia_url_categories** data source to get information about all or custom URL categories. By default, the response includes keywords.

` + "`" + `` + "`" + `` + "`" + `hcl
data "zia_url_categories" "example"{
    id = "CUSTOM_08"
}
` + "`" + `` + "`" + `` + "`" + `

`,
			Keywords: []string{
				"url",
				"categories",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(String) URL category ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "configured_name",
					Description: `(String) Name of the URL category. This is only required for custom URL categories.`,
				},
				resource.Attribute{
					Name:        "keywords",
					Description: `(List of String) Custom keywords associated to a URL category. Up to 2048 custom keywords can be added per organization across all categories (including bandwidth classes).`,
				},
				resource.Attribute{
					Name:        "super_category",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "custom_category",
					Description: `(Boolean) Set to true for custom URL category. Up to 48 custom URL categories can be added per organization.`,
				},
				resource.Attribute{
					Name:        "custom_urls_count",
					Description: `(Number) The number of custom URLs associated to the URL category.`,
				},
				resource.Attribute{
					Name:        "db_categorized_urls",
					Description: `(List of String) URLs added to a custom URL category are also retained under the original parent URL category (i.e., the predefined category the URL previously belonged to).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(String) Description of the category.`,
				},
				resource.Attribute{
					Name:        "editable",
					Description: `(Boolean) Value is set to false for custom URL category when due to scope user does not have edit permission`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(String) Type of the custom categories. ` + "`" + `URL_CATEGORY` + "`" + `, ` + "`" + `TLD_CATEGORY` + "`" + `, ` + "`" + `ALL` + "`" + ``,
				},
				resource.Attribute{
					Name:        "urls",
					Description: `(List of String) Custom URLs to add to a URL category. Up to 25,000 custom URLs can be added per organization across all categories (including bandwidth classes).`,
				},
				resource.Attribute{
					Name:        "urls_retaining_parent_category_count",
					Description: `(Number) The number of custom URLs associated to the URL category, that also need to be retained under the original parent category.`,
				},
				resource.Attribute{
					Name:        "val",
					Description: `(Number)`,
				},
				resource.Attribute{
					Name:        "scopes",
					Description: `(List of Object) Scope of the custom categories.`,
				},
				resource.Attribute{
					Name:        "scope_group_member_entities",
					Description: `(List of Object) Only applicable for the LOCATION_GROUP admin scope type, in which case this attribute gives the list of ID/name pairs of locations within the location group. The attribute name is subject to change. - ` + "`" + `id` + "`" + ` - (String) Identifier that uniquely identifies an entity - ` + "`" + `name` + "`" + ` - (String) The configured name of the entity - ` + "`" + `extensions` + "`" + ` - (Map of String)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(String) The admin scope type. The attribute name is subject to change. ` + "`" + `ORGANIZATION` + "`" + `, ` + "`" + `DEPARTMENT` + "`" + `, ` + "`" + `LOCATION` + "`" + `, ` + "`" + `LOCATION_GROUP` + "`" + ``,
				},
				resource.Attribute{
					Name:        "scope_entities",
					Description: `(List of Object) - ` + "`" + `id` + "`" + ` - (String) Identifier that uniquely identifies an entity - ` + "`" + `name` + "`" + ` - (String) The configured name of the entity - ` + "`" + `extensions` + "`" + ` - (Map of String)`,
				},
				resource.Attribute{
					Name:        "url_keyword_counts",
					Description: `(List of Object) URL and keyword counts for the category.`,
				},
				resource.Attribute{
					Name:        "total_url_count",
					Description: `(Number) Custom URL count for the category.`,
				},
				resource.Attribute{
					Name:        "retain_parent_url_count",
					Description: `(Number) Count of URLs with retain parent category.`,
				},
				resource.Attribute{
					Name:        "total_keyword_count",
					Description: `(Number) Total keyword count for the category.`,
				},
				resource.Attribute{
					Name:        "retain_parent_keyword_count",
					Description: `(Number) Count of total keywords with retain parent category.`,
				},
				resource.Attribute{
					Name:        "custom_ip_ranges_count",
					Description: `(Number) The number of custom IP address ranges associated to the URL category.`,
				},
				resource.Attribute{
					Name:        "ip_ranges_retaining_parent_category_count",
					Description: `(Number) The number of custom IP address ranges associated to the URL category, that also need to be retained under the original parent category.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "configured_name",
					Description: `(String) Name of the URL category. This is only required for custom URL categories.`,
				},
				resource.Attribute{
					Name:        "keywords",
					Description: `(List of String) Custom keywords associated to a URL category. Up to 2048 custom keywords can be added per organization across all categories (including bandwidth classes).`,
				},
				resource.Attribute{
					Name:        "super_category",
					Description: `(String)`,
				},
				resource.Attribute{
					Name:        "custom_category",
					Description: `(Boolean) Set to true for custom URL category. Up to 48 custom URL categories can be added per organization.`,
				},
				resource.Attribute{
					Name:        "custom_urls_count",
					Description: `(Number) The number of custom URLs associated to the URL category.`,
				},
				resource.Attribute{
					Name:        "db_categorized_urls",
					Description: `(List of String) URLs added to a custom URL category are also retained under the original parent URL category (i.e., the predefined category the URL previously belonged to).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(String) Description of the category.`,
				},
				resource.Attribute{
					Name:        "editable",
					Description: `(Boolean) Value is set to false for custom URL category when due to scope user does not have edit permission`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(String) Type of the custom categories. ` + "`" + `URL_CATEGORY` + "`" + `, ` + "`" + `TLD_CATEGORY` + "`" + `, ` + "`" + `ALL` + "`" + ``,
				},
				resource.Attribute{
					Name:        "urls",
					Description: `(List of String) Custom URLs to add to a URL category. Up to 25,000 custom URLs can be added per organization across all categories (including bandwidth classes).`,
				},
				resource.Attribute{
					Name:        "urls_retaining_parent_category_count",
					Description: `(Number) The number of custom URLs associated to the URL category, that also need to be retained under the original parent category.`,
				},
				resource.Attribute{
					Name:        "val",
					Description: `(Number)`,
				},
				resource.Attribute{
					Name:        "scopes",
					Description: `(List of Object) Scope of the custom categories.`,
				},
				resource.Attribute{
					Name:        "scope_group_member_entities",
					Description: `(List of Object) Only applicable for the LOCATION_GROUP admin scope type, in which case this attribute gives the list of ID/name pairs of locations within the location group. The attribute name is subject to change. - ` + "`" + `id` + "`" + ` - (String) Identifier that uniquely identifies an entity - ` + "`" + `name` + "`" + ` - (String) The configured name of the entity - ` + "`" + `extensions` + "`" + ` - (Map of String)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(String) The admin scope type. The attribute name is subject to change. ` + "`" + `ORGANIZATION` + "`" + `, ` + "`" + `DEPARTMENT` + "`" + `, ` + "`" + `LOCATION` + "`" + `, ` + "`" + `LOCATION_GROUP` + "`" + ``,
				},
				resource.Attribute{
					Name:        "scope_entities",
					Description: `(List of Object) - ` + "`" + `id` + "`" + ` - (String) Identifier that uniquely identifies an entity - ` + "`" + `name` + "`" + ` - (String) The configured name of the entity - ` + "`" + `extensions` + "`" + ` - (Map of String)`,
				},
				resource.Attribute{
					Name:        "url_keyword_counts",
					Description: `(List of Object) URL and keyword counts for the category.`,
				},
				resource.Attribute{
					Name:        "total_url_count",
					Description: `(Number) Custom URL count for the category.`,
				},
				resource.Attribute{
					Name:        "retain_parent_url_count",
					Description: `(Number) Count of URLs with retain parent category.`,
				},
				resource.Attribute{
					Name:        "total_keyword_count",
					Description: `(Number) Total keyword count for the category.`,
				},
				resource.Attribute{
					Name:        "retain_parent_keyword_count",
					Description: `(Number) Count of total keywords with retain parent category.`,
				},
				resource.Attribute{
					Name:        "custom_ip_ranges_count",
					Description: `(Number) The number of custom IP address ranges associated to the URL category.`,
				},
				resource.Attribute{
					Name:        "ip_ranges_retaining_parent_category_count",
					Description: `(Number) The number of custom IP address ranges associated to the URL category, that also need to be retained under the original parent category.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_url_filtering_rules",
			Category:         "URL Filtering Rule",
			ShortDescription: `Gets a list of all of URL Filtering Policy rules.`,
			Description: `

Use the **zia_url_filtering_rules** data source to get information about a URL filtering rule information for the specified ` + "`" + `Name` + "`" + `.

` + "`" + `` + "`" + `` + "`" + `hcl
# URL filtering rule
data "zia_url_filtering_rules" "example"{
    name = "Example"
}
` + "`" + `` + "`" + `` + "`" + `

`,
			Keywords: []string{
				"url",
				"filtering",
				"rule",
				"rules",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(String) Name of the URL Filtering policy rule`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(String) URL Filtering Rule ID ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Number) Order of execution of rule with respect to other URL Filtering rules`,
				},
				resource.Attribute{
					Name:        "protocols",
					Description: `(List of Object) Protocol criteria. Supported values: ` + "`" + `SMRULEF_ZPA_BROKERS_RULE` + "`" + `, ` + "`" + `ANY_RULE` + "`" + `, ` + "`" + `TCP_RULE` + "`" + `, ` + "`" + `UDP_RULE` + "`" + `, ` + "`" + `DOHTTPS_RULE` + "`" + `, ` + "`" + `TUNNELSSL_RULE` + "`" + `, ` + "`" + `HTTP_PROXY` + "`" + `, ` + "`" + `FOHTTP_RULE` + "`" + `, ` + "`" + `FTP_RULE` + "`" + `, ` + "`" + `HTTPS_RULE` + "`" + `, ` + "`" + `HTTP_RULE` + "`" + `, ` + "`" + `SSL_RULE` + "`" + `, ` + "`" + `TUNNEL_RULE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(String) Rule State`,
				},
				resource.Attribute{
					Name:        "rank",
					Description: `(String) Admin rank of the admin who creates this rule`,
				},
				resource.Attribute{
					Name:        "request_methods",
					Description: `(String) Request method for which the rule must be applied. If not set, rule will be applied to all methods`,
				},
				resource.Attribute{
					Name:        "end_user_notification_url",
					Description: `(String) URL of end user notification page to be displayed when the rule is matched. Not applicable if either 'overrideUsers' or 'overrideGroups' is specified.`,
				},
				resource.Attribute{
					Name:        "block_override",
					Description: `(String) When set to true, a ` + "`" + `BLOCK` + "`" + ` action triggered by the rule could be overridden. If true and both overrideGroup and overrideUsers are not set, the ` + "`" + `BLOCK` + "`" + ` triggered by this rule could be overridden for any users. If block)Override is not set, ` + "`" + `BLOCK` + "`" + ` action cannot be overridden.`,
				},
				resource.Attribute{
					Name:        "time_quota",
					Description: `(String) Time quota in minutes, after which the URL Filtering rule is applied. If not set, no quota is enforced. If a policy rule action is set to ` + "`" + `BLOCK` + "`" + `, this field is not applicable.`,
				},
				resource.Attribute{
					Name:        "size_quota",
					Description: `(String) Size quota in KB beyond which the URL Filtering rule is applied. If not set, no quota is enforced. If a policy rule action is set to ` + "`" + `BLOCK` + "`" + `, this field is not applicable.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(String) Additional information about the rule`,
				},
				resource.Attribute{
					Name:        "validity_start_time",
					Description: `(Number) If enforceTimeValidity is set to true, the URL Filtering rule will be valid starting on this date and time.`,
				},
				resource.Attribute{
					Name:        "validity_end_time",
					Description: `(Number) If enforceTimeValidity is set to true, the URL Filtering rule will cease to be valid on this end date and time.`,
				},
				resource.Attribute{
					Name:        "validity_time_zone_id",
					Description: `(Number) If enforceTimeValidity is set to true, the URL Filtering rule date and time will be valid based on this time zone ID.`,
				},
				resource.Attribute{
					Name:        "last_modified_time",
					Description: `(Number) When the rule was last modified`,
				},
				resource.Attribute{
					Name:        "enforce_time_validity",
					Description: `(String) Enforce a set a validity time period for the URL Filtering rule.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(String) Action taken when traffic matches rule criteria. Supported values: ` + "`" + `ANY` + "`" + `, ` + "`" + `NONE` + "`" + `, ` + "`" + `BLOCK` + "`" + `, ` + "`" + `CAUTION` + "`" + `, ` + "`" + `ALLOW` + "`" + `, ` + "`" + `ICAP_RESPONSE` + "`" + ``,
				},
				resource.Attribute{
					Name:        "cipa_rule",
					Description: `(String) If set to true, the CIPA Compliance rule is enabled`,
				},
				resource.Attribute{
					Name:        "url_categories",
					Description: `(String) List of URL categories for which rule must be applied`,
				},
				resource.Attribute{
					Name:        "locations",
					Description: `(List of Object) The locations to which the Firewall Filtering policy rule applies`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "extensions",
					Description: `(Map of String)`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `(List of Object) The groups to which the Firewall Filtering policy rule applies`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "extensions",
					Description: `(Map of String)`,
				},
				resource.Attribute{
					Name:        "departments",
					Description: `(List of Object) The departments to which the Firewall Filtering policy rule applies`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "extensions",
					Description: `(Map of String)`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `(List of Object) The users to which the Firewall Filtering policy rule applies`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "extensions",
					Description: `(Map of String)`,
				},
				resource.Attribute{
					Name:        "time_windows",
					Description: `(List of Object) The time interval in which the Firewall Filtering policy rule applies`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "extensions",
					Description: `(Map of String)`,
				},
				resource.Attribute{
					Name:        "override_users",
					Description: `(List of Object) Name-ID pairs of users for which this rule can be overridden. Applicable only if blockOverride is set to ` + "`" + `true` + "`" + `, action is ` + "`" + `BLOCK` + "`" + ` and overrideGroups is not set.If this overrideUsers is not set, ` + "`" + `BLOCK` + "`" + ` action can be overridden for any user.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "extensions",
					Description: `(Map of String)`,
				},
				resource.Attribute{
					Name:        "override_groups",
					Description: `(List of Object) Name-ID pairs of users for which this rule can be overridden. Applicable only if blockOverride is set to ` + "`" + `true` + "`" + `, action is ` + "`" + `BLOCK` + "`" + ` and overrideGroups is not set.If this overrideUsers is not set, ` + "`" + `BLOCK` + "`" + ` action can be overridden for any group.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "extensions",
					Description: `(Map of String)`,
				},
				resource.Attribute{
					Name:        "location_groups",
					Description: `(List of Object) The location groups to which the Firewall Filtering policy rule applies`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "extensions",
					Description: `(Map of String)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "extensions",
					Description: `(Map of String)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "extensions",
					Description: `(Map of String)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "order",
					Description: `(Number) Order of execution of rule with respect to other URL Filtering rules`,
				},
				resource.Attribute{
					Name:        "protocols",
					Description: `(List of Object) Protocol criteria. Supported values: ` + "`" + `SMRULEF_ZPA_BROKERS_RULE` + "`" + `, ` + "`" + `ANY_RULE` + "`" + `, ` + "`" + `TCP_RULE` + "`" + `, ` + "`" + `UDP_RULE` + "`" + `, ` + "`" + `DOHTTPS_RULE` + "`" + `, ` + "`" + `TUNNELSSL_RULE` + "`" + `, ` + "`" + `HTTP_PROXY` + "`" + `, ` + "`" + `FOHTTP_RULE` + "`" + `, ` + "`" + `FTP_RULE` + "`" + `, ` + "`" + `HTTPS_RULE` + "`" + `, ` + "`" + `HTTP_RULE` + "`" + `, ` + "`" + `SSL_RULE` + "`" + `, ` + "`" + `TUNNEL_RULE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(String) Rule State`,
				},
				resource.Attribute{
					Name:        "rank",
					Description: `(String) Admin rank of the admin who creates this rule`,
				},
				resource.Attribute{
					Name:        "request_methods",
					Description: `(String) Request method for which the rule must be applied. If not set, rule will be applied to all methods`,
				},
				resource.Attribute{
					Name:        "end_user_notification_url",
					Description: `(String) URL of end user notification page to be displayed when the rule is matched. Not applicable if either 'overrideUsers' or 'overrideGroups' is specified.`,
				},
				resource.Attribute{
					Name:        "block_override",
					Description: `(String) When set to true, a ` + "`" + `BLOCK` + "`" + ` action triggered by the rule could be overridden. If true and both overrideGroup and overrideUsers are not set, the ` + "`" + `BLOCK` + "`" + ` triggered by this rule could be overridden for any users. If block)Override is not set, ` + "`" + `BLOCK` + "`" + ` action cannot be overridden.`,
				},
				resource.Attribute{
					Name:        "time_quota",
					Description: `(String) Time quota in minutes, after which the URL Filtering rule is applied. If not set, no quota is enforced. If a policy rule action is set to ` + "`" + `BLOCK` + "`" + `, this field is not applicable.`,
				},
				resource.Attribute{
					Name:        "size_quota",
					Description: `(String) Size quota in KB beyond which the URL Filtering rule is applied. If not set, no quota is enforced. If a policy rule action is set to ` + "`" + `BLOCK` + "`" + `, this field is not applicable.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(String) Additional information about the rule`,
				},
				resource.Attribute{
					Name:        "validity_start_time",
					Description: `(Number) If enforceTimeValidity is set to true, the URL Filtering rule will be valid starting on this date and time.`,
				},
				resource.Attribute{
					Name:        "validity_end_time",
					Description: `(Number) If enforceTimeValidity is set to true, the URL Filtering rule will cease to be valid on this end date and time.`,
				},
				resource.Attribute{
					Name:        "validity_time_zone_id",
					Description: `(Number) If enforceTimeValidity is set to true, the URL Filtering rule date and time will be valid based on this time zone ID.`,
				},
				resource.Attribute{
					Name:        "last_modified_time",
					Description: `(Number) When the rule was last modified`,
				},
				resource.Attribute{
					Name:        "enforce_time_validity",
					Description: `(String) Enforce a set a validity time period for the URL Filtering rule.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(String) Action taken when traffic matches rule criteria. Supported values: ` + "`" + `ANY` + "`" + `, ` + "`" + `NONE` + "`" + `, ` + "`" + `BLOCK` + "`" + `, ` + "`" + `CAUTION` + "`" + `, ` + "`" + `ALLOW` + "`" + `, ` + "`" + `ICAP_RESPONSE` + "`" + ``,
				},
				resource.Attribute{
					Name:        "cipa_rule",
					Description: `(String) If set to true, the CIPA Compliance rule is enabled`,
				},
				resource.Attribute{
					Name:        "url_categories",
					Description: `(String) List of URL categories for which rule must be applied`,
				},
				resource.Attribute{
					Name:        "locations",
					Description: `(List of Object) The locations to which the Firewall Filtering policy rule applies`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "extensions",
					Description: `(Map of String)`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `(List of Object) The groups to which the Firewall Filtering policy rule applies`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "extensions",
					Description: `(Map of String)`,
				},
				resource.Attribute{
					Name:        "departments",
					Description: `(List of Object) The departments to which the Firewall Filtering policy rule applies`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "extensions",
					Description: `(Map of String)`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `(List of Object) The users to which the Firewall Filtering policy rule applies`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "extensions",
					Description: `(Map of String)`,
				},
				resource.Attribute{
					Name:        "time_windows",
					Description: `(List of Object) The time interval in which the Firewall Filtering policy rule applies`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "extensions",
					Description: `(Map of String)`,
				},
				resource.Attribute{
					Name:        "override_users",
					Description: `(List of Object) Name-ID pairs of users for which this rule can be overridden. Applicable only if blockOverride is set to ` + "`" + `true` + "`" + `, action is ` + "`" + `BLOCK` + "`" + ` and overrideGroups is not set.If this overrideUsers is not set, ` + "`" + `BLOCK` + "`" + ` action can be overridden for any user.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "extensions",
					Description: `(Map of String)`,
				},
				resource.Attribute{
					Name:        "override_groups",
					Description: `(List of Object) Name-ID pairs of users for which this rule can be overridden. Applicable only if blockOverride is set to ` + "`" + `true` + "`" + `, action is ` + "`" + `BLOCK` + "`" + ` and overrideGroups is not set.If this overrideUsers is not set, ` + "`" + `BLOCK` + "`" + ` action can be overridden for any group.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "extensions",
					Description: `(Map of String)`,
				},
				resource.Attribute{
					Name:        "location_groups",
					Description: `(List of Object) The location groups to which the Firewall Filtering policy rule applies`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "extensions",
					Description: `(Map of String)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "extensions",
					Description: `(Map of String)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "extensions",
					Description: `(Map of String)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_user_management",
			Category:         "User Management",
			ShortDescription: `Gets a list of all users and allows user filtering by name, department, or group`,
			Description: `

Use the **zia_user_management** data source to get information about a user account that may have been created in the Zscaler Internet Access portal or via API. This data source can then be associated with a ZIA cloud firewall filtering rule, and URL filtering rules.

`,
			Keywords: []string{
				"user",
				"management",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) User name. This appears when choosing users for policies.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the time window resource. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Required) User email consists of a user name and domain name. It does not have to be a valid email address, but it must be unique and its domain must belong to the organization`,
				},
				resource.Attribute{
					Name:        "admin_user",
					Description: `(String) True if this user is an Admin user. readOnly: ` + "`" + `true` + "`" + ` default: ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(String) Additional information about this user.`,
				},
				resource.Attribute{
					Name:        "temp_auth_email",
					Description: `(String) Temporary Authentication Email. If you enabled one-time tokens or links, enter the email address to which the Zscaler service sends the tokens or links. If this is empty, the service will send the email to the User email.`,
				},
				resource.Attribute{
					Name:        "auth_methods",
					Description: `(String) Type of authentication method to be enabled. Supported values are: ` + "`" + `` + "`" + `BASIC` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `DIGEST` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(String) User type. Provided only if this user is not an end user. The supported types are:`,
				},
				resource.Attribute{
					Name:        "department",
					Description: `(String) Department a user belongs to`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Department ID`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(String) Department name`,
				},
				resource.Attribute{
					Name:        "idp_id",
					Description: `(Number) Identity provider (IdP) ID`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(String) Additional information about this department`,
				},
				resource.Attribute{
					Name:        "deleted",
					Description: `(Boolean) default: ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `(String) List of Groups a user belongs to. Groups are used in policies.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Unique identfier for the group`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(String) Group name`,
				},
				resource.Attribute{
					Name:        "idp_id",
					Description: `(Number) Unique identfier for the identity provider (IdP)`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(String) Additional information about the group`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "email",
					Description: `(Required) User email consists of a user name and domain name. It does not have to be a valid email address, but it must be unique and its domain must belong to the organization`,
				},
				resource.Attribute{
					Name:        "admin_user",
					Description: `(String) True if this user is an Admin user. readOnly: ` + "`" + `true` + "`" + ` default: ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(String) Additional information about this user.`,
				},
				resource.Attribute{
					Name:        "temp_auth_email",
					Description: `(String) Temporary Authentication Email. If you enabled one-time tokens or links, enter the email address to which the Zscaler service sends the tokens or links. If this is empty, the service will send the email to the User email.`,
				},
				resource.Attribute{
					Name:        "auth_methods",
					Description: `(String) Type of authentication method to be enabled. Supported values are: ` + "`" + `` + "`" + `BASIC` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `DIGEST` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(String) User type. Provided only if this user is not an end user. The supported types are:`,
				},
				resource.Attribute{
					Name:        "department",
					Description: `(String) Department a user belongs to`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Department ID`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(String) Department name`,
				},
				resource.Attribute{
					Name:        "idp_id",
					Description: `(Number) Identity provider (IdP) ID`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(String) Additional information about this department`,
				},
				resource.Attribute{
					Name:        "deleted",
					Description: `(Boolean) default: ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `(String) List of Groups a user belongs to. Groups are used in policies.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Number) Unique identfier for the group`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(String) Group name`,
				},
				resource.Attribute{
					Name:        "idp_id",
					Description: `(Number) Unique identfier for the identity provider (IdP)`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(String) Additional information about the group`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_user_management_department",
			Category:         "User Management",
			ShortDescription: `Gets a list of user departments details.`,
			Description: `

Use the **zia_department_management** data source to get information about user department created in the Zscaler Internet Access cloud or via the API. This data source can then be associated with several ZIA resources such as: URL filtering rules, Cloud Firewall rules, and locations.

`,
			Keywords: []string{
				"user",
				"management",
				"department",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(String) Name of the user department`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(String) ID of the user department ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "idp_id",
					Description: `(Optional) Unique identfier for the identity provider (IdP)`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Additional information about this department`,
				},
				resource.Attribute{
					Name:        "deleted",
					Description: `(Boolean) default: false`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "idp_id",
					Description: `(Optional) Unique identfier for the identity provider (IdP)`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Additional information about this department`,
				},
				resource.Attribute{
					Name:        "deleted",
					Description: `(Boolean) default: false`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_user_management_group",
			Category:         "User Management",
			ShortDescription: `Gets a list of user user group.`,
			Description:      ``,
			Keywords: []string{
				"user",
				"management",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the user group`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Unique identfier for the group. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "idp_id",
					Description: `(Optional) Unique identfier for the identity provider (IdP)`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Additional information about the group`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "idp_id",
					Description: `(Optional) Unique identfier for the identity provider (IdP)`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Additional information about the group`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"zia_zia_activation_status":                             0,
		"zia_zia_admin_roles":                                   1,
		"zia_zia_admin_users":                                   2,
		"zia_zia_auth_settings_urls":                            3,
		"zia_zia_device_groups":                                 4,
		"zia_zia_devices":                                       5,
		"zia_zia_dlp_dictionaries":                              6,
		"zia_zia_dlp_engines":                                   7,
		"zia_zia_dlp_icap_servers":                              8,
		"zia_zia_dlp_idm_profiles":                              9,
		"zia_zia_dlp_idm_profiles_lite":                         10,
		"zia_zia_dlp_incident_receiver_servers":                 11,
		"zia_zia_dlp_notification_templates":                    12,
		"zia_zia_dlp_web_rules":                                 13,
		"zia_zia_firewall_filtering_destination_groups":         14,
		"zia_zia_firewall_filtering_ip_source_groups":           15,
		"zia_zia_firewall_filtering_network_application":        16,
		"zia_zia_firewall_filtering_network_application_groups": 17,
		"zia_zia_firewall_filtering_network_service":            18,
		"zia_zia_firewall_filtering_network_service_groups":     19,
		"zia_zia_firewall_filtering_rule":                       20,
		"zia_zia_firewall_filtering_time_window":                21,
		"zia_zia_location_groups":                               22,
		"zia_zia_location_lite":                                 23,
		"zia_zia_location_management":                           24,
		"zia_zia_rule_labels":                                   25,
		"zia_zia_security_policy_settings":                      26,
		"zia_zia_sub_location_management":                       27,
		"zia_zia_traffic_forwarding_gre_internal_ip_range_list": 28,
		"zia_zia_traffic_forwarding_gre_tunnel_info":            29,
		"zia_zia_traffic_forwarding_gre_vip_recommended_list":   30,
		"zia_zia_traffic_forwarding_public_node_vips":           31,
		"zia_zia_traffic_forwarding_static_ip":                  32,
		"zia_zia_traffic_forwarding_vpn_credentials":            33,
		"zia_zia_url_categories":                                34,
		"zia_zia_url_filtering_rules":                           35,
		"zia_zia_user_management":                               36,
		"zia_zia_user_management_department":                    37,
		"zia_zia_user_management_group":                         38,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
