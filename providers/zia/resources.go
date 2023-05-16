package zia

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_activation_status",
			Category:         "Activation",
			ShortDescription: `"Activates configuration changes".`,
			Description: `

The **zia_activation_status** resource allows the activation of ZIA pending configurations. This resource must always be executed after the resource creation for successfully policy/configuration activation to occur.

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
			Type:             "zia_zia_admin_users",
			Category:         "Admin & Role Management",
			ShortDescription: `Creates and manages ZIA administrator users.`,
			Description: `

The **zia_admin_users** resource allows the creation and management of ZIA admin user account created in the Zscaler Internet Access cloud or via the API.

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
					Name:        "password",
					Description: `(Required) The username of the admin user to be exported.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) Role of the admin. This is not required for an auditor.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) Identifier that uniquely identifies an entity !>`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) Admin or auditor's email address.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Additional information about the admin or auditor.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional) Indicates whether or not the admin account is disabled.`,
				},
				resource.Attribute{
					Name:        "is_auditor",
					Description: `(Optional) Indicates whether the user is an auditor. This attribute is subject to change.`,
				},
				resource.Attribute{
					Name:        "is_exec_mobile_app_enabled",
					Description: `(Optional) Indicates whether or not Executive Insights App access is enabled for the admin.`,
				},
				resource.Attribute{
					Name:        "is_non_editable",
					Description: `(Optional) Indicates whether or not the admin can be edited or deleted.`,
				},
				resource.Attribute{
					Name:        "is_password_expired",
					Description: `(Optional) Indicates whether or not an admin's password has expired.`,
				},
				resource.Attribute{
					Name:        "is_password_login_allowed",
					Description: `(Optional) The default is true when SAML Authentication is disabled. When SAML Authentication is enabled, this can be set to false in order to force the admin to login via SSO only.`,
				},
				resource.Attribute{
					Name:        "is_product_update_comm_enabled",
					Description: `(Optional) Communication setting for Product Update.`,
				},
				resource.Attribute{
					Name:        "is_security_report_comm_enabled",
					Description: `(Optional) Communication for Security Report is enabled.`,
				},
				resource.Attribute{
					Name:        "is_service_update_comm_enabled",
					Description: `(Optional) Communication setting for Service Update.`,
				},
				resource.Attribute{
					Name:        "admin_scope",
					Description: `(Optional) The admin's scope. A scope is required for admins, but not applicable to auditors. This attribute is subject to change.`,
				},
				resource.Attribute{
					Name:        "scope_group_member_entities",
					Description: `(Optional) Only applicable for the LOCATION_GROUP admin scope type, in which case this attribute gives the list of ID/name pairs of locations within the location group.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The configured name of the entity`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The admin scope type. The attribute name is subject to change.`,
				},
				resource.Attribute{
					Name:        "scope_entities",
					Description: `(Optional) Based on the admin scope type, the entities can be the ID/name pair of departments, locations, or location groups.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The configured name of the entity`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_auth_settings_urls",
			Category:         "User Authentication Settings",
			ShortDescription: `Adds a URL to or removes a URL from the cookie authentication exempt list`,
			Description: `

The **zia_auth_settings_urls** resource alows you to add or remove a URL from the cookie authentication exempt list in the Zscaler Internet Access cloud or via the API. To learn more see [URL Format Guidelines](https://help.zscaler.com/zia/url-format-guidelines)

`,
			Keywords: []string{
				"user",
				"authentication",
				"settings",
				"auth",
				"urls",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "urls",
					Description: `(Required) The email address of the admin user to be exported. ### Optional There are no optional parameters supported by this resource.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_dlp_dictionaries",
			Category:         "Data Loss Prevention",
			ShortDescription: `Creates and manages ZIA DLP dictionaries.`,
			Description: `

The **zia_dlp_dictionaries** resource allows the creation and management of ZIA DLP dictionaries in the Zscaler Internet Access cloud or via the API.

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
					Description: `(Required) The DLP dictionary's name`,
				},
				resource.Attribute{
					Name:        "dictionary_type",
					Description: `(Required) The DLP dictionary type. The following values are supported:`,
				},
				resource.Attribute{
					Name:        "custom_phrase_match_type",
					Description: `(Required) The DLP custom phrase match type. Supported values are:`,
				},
				resource.Attribute{
					Name:        "phrases",
					Description: `(Required) List containing the phrases used within a custom DLP dictionary. This attribute is not applicable to predefined DLP dictionaries. Required when ` + "`" + `dictionary_type` + "`" + ` is ` + "`" + `PATTERNS_AND_PHRASES` + "`" + ``,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) The action applied to a DLP dictionary using patterns. The following values are supported:`,
				},
				resource.Attribute{
					Name:        "phrase",
					Description: `(Required) DLP dictionary phrase`,
				},
				resource.Attribute{
					Name:        "patterns",
					Description: `(Required) List containing the patterns used within a custom DLP dictionary. This attribute is not applicable to predefined DLP dictionaries. Required when ` + "`" + `dictionary_type` + "`" + ` is ` + "`" + `PATTERNS_AND_PHRASES` + "`" + ``,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) The action applied to a DLP dictionary using patterns. The following values are supported:`,
				},
				resource.Attribute{
					Name:        "pattern",
					Description: `(Required) DLP dictionary pattern ### Optional`,
				},
				resource.Attribute{
					Name:        "confidence_threshold",
					Description: `(Optional) The DLP confidence threshold. The following values are supported:`,
				},
				resource.Attribute{
					Name:        "threshold_type",
					Description: `(Optional) The DLP threshold type. The following values are supported:`,
				},
				resource.Attribute{
					Name:        "threshold_type",
					Description: `(Optional) The DLP threshold type. The following values are supported:`,
				},
				resource.Attribute{
					Name:        "exact_data_match_details",
					Description: `(Optional) Exact Data Match (EDM) related information for custom DLP dictionaries.`,
				},
				resource.Attribute{
					Name:        "dictionary_edm_mapping_id",
					Description: `(Optional) The unique identifier for the EDM mapping.`,
				},
				resource.Attribute{
					Name:        "schema_id",
					Description: `(Optional) The unique identifier for the EDM template (or schema).`,
				},
				resource.Attribute{
					Name:        "primary_field",
					Description: `(Optional) The EDM template's primary field.`,
				},
				resource.Attribute{
					Name:        "secondary_fields",
					Description: `(Optional) The EDM template's secondary fields.`,
				},
				resource.Attribute{
					Name:        "secondary_field_match_on",
					Description: `(Optional) The EDM secondary field to match on. - ` + "`" + `"MATCHON_NONE"` + "`" + ` - ` + "`" + `"MATCHON_ANY_1"` + "`" + ` - ` + "`" + `"MATCHON_ANY_2"` + "`" + ` - ` + "`" + `"MATCHON_ANY_3"` + "`" + ` - ` + "`" + `"MATCHON_ANY_4"` + "`" + ` - ` + "`" + `"MATCHON_ANY_5"` + "`" + ` - ` + "`" + `"MATCHON_ANY_6"` + "`" + ` - ` + "`" + `"MATCHON_ANY_7"` + "`" + ` - ` + "`" + `"MATCHON_ANY_8"` + "`" + ` - ` + "`" + `"MATCHON_ANY_9"` + "`" + ` - ` + "`" + `"MATCHON_ANY_10"` + "`" + ` - ` + "`" + `"MATCHON_ANY_11"` + "`" + ` - ` + "`" + `"MATCHON_ANY_12"` + "`" + ` - ` + "`" + `"MATCHON_ANY_13"` + "`" + ` - ` + "`" + `"MATCHON_ANY_14"` + "`" + ` - ` + "`" + `"MATCHON_ANY_15"` + "`" + ` - ` + "`" + `"MATCHON_ALL"` + "`" + ``,
				},
				resource.Attribute{
					Name:        "idm_profile_match_accuracy",
					Description: `(Optional) List of Indexed Document Match (IDM) profiles and their corresponding match accuracy for custom DLP dictionaries.`,
				},
				resource.Attribute{
					Name:        "adp_idm_profile",
					Description: `(Optional) The IDM template reference.`,
				},
				resource.Attribute{
					Name:        "match_accuracy",
					Description: `(Optional) The IDM template match accuracy. - ` + "`" + `"LOW"` + "`" + ` - ` + "`" + `"MEDIUM"` + "`" + ` - ` + "`" + `"HEAVY"` + "`" + ``,
				},
				resource.Attribute{
					Name:        "proximity",
					Description: `(Optional) The DLP dictionary proximity length.`,
				},
				resource.Attribute{
					Name:        "ignore_exact_match_idm_dict",
					Description: `(Optional) Indicates whether to exclude documents that are a 100% match to already-indexed documents from triggering an Indexed Document Match (IDM) Dictionary.`,
				},
				resource.Attribute{
					Name:        "include_bin_numbers",
					Description: `(Optional) A true value denotes that the specified Bank Identification Number (BIN) values are included in the Credit Cards dictionary. A false value denotes that the specified BIN values are excluded from the Credit Cards dictionary. Note: This field is applicable only to the predefined Credit Cards dictionary and its clones.`,
				},
				resource.Attribute{
					Name:        "bin_numbers",
					Description: `(Optional) The list of Bank Identification Number (BIN) values that are included or excluded from the Credit Cards dictionary. BIN values can be specified only for Diners Club, Mastercard, RuPay, and Visa cards. Up to 512 BIN values can be configured in a dictionary. Note: This field is applicable only to the predefined Credit Cards dictionary and its clones.`,
				},
				resource.Attribute{
					Name:        "dict_template_id",
					Description: `(Optional) ID of the predefined dictionary (original source dictionary) that is used for cloning. This field is applicable only to cloned dictionaries. Only a limited set of identification-based predefined dictionaries (e.g., Credit Cards, Social Security Numbers, National Identification Numbers, etc.) can be cloned. Up to 4 clones can be created from a predefined dictionary.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_dlp_notification_templates",
			Category:         "Data Loss Prevention",
			ShortDescription: `Creates and manages ZIA DLP Notification Templates.`,
			Description: `

The **zia_dlp_notification_templates** resource allows the creation and management of ZIA DLP Notification Templates in the Zscaler Internet Access cloud or via the API.

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
					Description: `(Required) The DLP policy rule name.`,
				},
				resource.Attribute{
					Name:        "plain_text_message",
					Description: `(Required) The template for the plain text UTF-8 message body that must be displayed in the DLP notification email.`,
				},
				resource.Attribute{
					Name:        "html_message",
					Description: `(Required) The template for the HTML message body that must be displayed in the DLP notification email. ### Optional`,
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
			ShortDescription: `Creates and manages ZIA DLP Web Rules.`,
			Description: `

The **zia_dlp_web_rules** resource allows the creation and management of ZIA DLP Web Rules in the Zscaler Internet Access cloud or via the API.

⚠️ **WARNING:** Zscaler Internet Access DLP supports a maximum of 127 Web DLP Rules to be created via API.

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
					Description: `(Required) The DLP policy rule name.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Required) The rule order of execution for the DLP policy rule with respect to other rules. ### Optional`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the DLP policy rule.`,
				},
				resource.Attribute{
					Name:        "external_auditor_email",
					Description: `(Optional) The email address of an external auditor to whom DLP email notifications are sent.`,
				},
				resource.Attribute{
					Name:        "match_only",
					Description: `(Optional) The match only criteria for DLP engines.`,
				},
				resource.Attribute{
					Name:        "without_content_inspection",
					Description: `(Optional) Indicates a DLP policy rule without content inspection, when the value is set to true.`,
				},
				resource.Attribute{
					Name:        "ocr_enabled",
					Description: `(Optional) Enables or disables image file scanning. When OCR is enabled only the following ` + "`" + `` + "`" + `file_types` + "`" + `` + "`" + ` are supported: ` + "`" + `` + "`" + `WINDOWS_META_FORMAT` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `BITMAP` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `JPEG` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `PNG` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `TIFF` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "zscaler_incident_receiver",
					Description: `(Optional) Indicates whether a Zscaler Incident Receiver is associated to the DLP policy rule.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) The action taken when traffic matches the DLP policy rule criteria. The supported values are:`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) Enables or disables the DLP policy rule.. The supported values are:`,
				},
				resource.Attribute{
					Name:        "file_types",
					Description: `(Optional) The list of file types to which the DLP policy rule must be applied. For the complete list of supported file types refer to the [ZIA API documentation](https://help.zscaler.com/zia/data-loss-prevention#/webDlpRules-post)`,
				},
				resource.Attribute{
					Name:        "cloud_applications",
					Description: `(Optional) The list of cloud applications to which the DLP policy rule must be applied. For the complete list of supported cloud applications refer to the [ZIA API documentation](https://help.zscaler.com/zia/data-loss-prevention#/webDlpRules-post)`,
				},
				resource.Attribute{
					Name:        "notification_template",
					Description: `(Optional) The template used for DLP notification emails.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "auditor",
					Description: `(Optional) The auditor to which the DLP policy rule must be applied.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "url_categories",
					Description: `(Optional) The list of URL categories to which the DLP policy rule must be applied.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "dlp_engines",
					Description: `(Optional) The list of DLP engines to which the DLP policy rule must be applied.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Identifier that uniquely identifies an entity. Maximum of up to ` + "`" + `4` + "`" + ` dlp engines. When not used it implies ` + "`" + `Any` + "`" + ` to apply the rule to all locations.`,
				},
				resource.Attribute{
					Name:        "locations",
					Description: `(Optional) The Name-ID pairs of locations to which the DLP policy rule must be applied. Maximum of up to ` + "`" + `8` + "`" + ` locations. When not used it implies ` + "`" + `Any` + "`" + ` to apply the rule to all locations.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "location_groups",
					Description: `(Optional) The Name-ID pairs of locations groups to which the DLP policy rule must be applied. Maximum of up to ` + "`" + `32` + "`" + ` location groups. When not used it implies ` + "`" + `Any` + "`" + ` to apply the rule to all location groups.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `(Optional) The Name-ID pairs of users to which the DLP policy rule must be applied. Maximum of up to ` + "`" + `4` + "`" + ` users. When not used it implies ` + "`" + `Any` + "`" + ` to apply the rule to all users.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `(Optional) The Name-ID pairs of groups to which the DLP policy rule must be applied. Maximum of up to ` + "`" + `8` + "`" + ` groups. When not used it implies ` + "`" + `Any` + "`" + ` to apply the rule to all groups.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "departments",
					Description: `(Optional) The name-ID pairs of the departments that are excluded from the DLP policy rule.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "excluded_users",
					Description: `(Optional) The name-ID pairs of the users that are excluded from the DLP policy rule. Maximum of up to ` + "`" + `256` + "`" + ` users.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "excluded_departments",
					Description: `(Optional) The name-ID pairs of the groups that are excluded from the DLP policy rule. Maximum of up to ` + "`" + `256` + "`" + ` departments.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "excluded_groups",
					Description: `(Optional) The name-ID pairs of the groups that are excluded from the DLP policy rule. Maximum of up to ` + "`" + `256` + "`" + ` groups.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "time_windows",
					Description: `(Optional) The Name-ID pairs of time windows to which the DLP policy rule must be applied. Maximum of up to ` + "`" + `2` + "`" + ` time intervals. When not used it implies ` + "`" + `always` + "`" + ` to apply the rule to all time intervals.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Identifier that uniquely identifies an entity`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_firewall_filtering_ip_destination_groups",
			Category:         "Firewall Policies",
			ShortDescription: `Creates and manages ZIA Cloud firewall IP destination groups.`,
			Description: `

The **zia_firewall_filtering_destination_groups** resource allows the creation and management of ZIA Cloud Firewall IP destination groups in the Zscaler Internet Access. This resource can then be associated with a ZIA cloud firewall filtering rule.

`,
			Keywords: []string{
				"firewall",
				"policies",
				"filtering",
				"ip",
				"destination",
				"groups",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Destination IP group name`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Destination IP group type (i.e., the group can contain destination IP addresses or FQDNs). The supported values are:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_firewall_filtering_ip_source_groups",
			Category:         "Firewall Policies",
			ShortDescription: `Creates and manages ZIA Cloud firewall IP source groups.`,
			Description: `

The **zia_firewall_filtering_ip_source_groups** resource allows the creation and management of ZIA Cloud Firewall IP source groups in the Zscaler Internet Access. This resource can then be associated with a ZIA cloud firewall filtering rule.

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
					Description: `(Required) Source IP group name`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_firewall_filtering_network_application_groups",
			Category:         "Firewall Policies",
			ShortDescription: `Creates and manages ZIA Cloud firewall Network Application Groups.`,
			Description: `

The **zia_firewall_filtering_network_application_groups** resource allows the creation and management of ZIA Cloud Firewall IP source groups in the Zscaler Internet Access. This resource can then be associated with a ZIA cloud firewall filtering rule.

`,
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
					Description: `(Required) Network application group name`,
				},
				resource.Attribute{
					Name:        "network_applications",
					Description: `(Required) Any number of applications to be added to the group`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_firewall_filtering_network_service",
			Category:         "Firewall Policies",
			ShortDescription: `Creates and manages a ZIA Cloud firewall network service.`,
			Description: `

The **zia_firewall_filtering_network_service** resource allows the creation and management of ZIA Cloud Firewall IP network services in the Zscaler Internet Access. This resource can then be associated with a ZIA cloud firewall filtering rule and network service group resources.

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
					Description: `(Required) Name of the service`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Supported values: ` + "`" + `STANDARD` + "`" + `, ` + "`" + `PREDEFINED` + "`" + `, ` + "`" + `CUSTOM` + "`" + ``,
				},
				resource.Attribute{
					Name:        "src_tcp_ports",
					Description: `(Required) The TCP source port number (example: 50) or port number range (example: 1000-1050), if any, that is used by the network service`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `(Required)`,
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
					Description: `(Number) ->`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the service`,
				},
				resource.Attribute{
					Name:        "is_name_l10n_tag",
					Description: `(Optional`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) The following values are supported: ` + "`" + `"ICMP_ANY` + "`" + `, ` + "`" + `"UDP_ANY"` + "`" + `, ` + "`" + `"TCP_ANY"` + "`" + `, ` + "`" + `"OTHER_NETWORK_SERVICE"` + "`" + `, ` + "`" + `"DNS"` + "`" + `, ` + "`" + `"NETBIOS"` + "`" + `, ` + "`" + `"FTP"` + "`" + `, ` + "`" + `"GNUTELLA"` + "`" + `, ` + "`" + `"H_323"` + "`" + `, ` + "`" + `"HTTP"` + "`" + `, ` + "`" + `"HTTPS"` + "`" + `, ` + "`" + `"IKE"` + "`" + `, ` + "`" + `"IMAP"` + "`" + `, ` + "`" + `"ILS"` + "`" + `, ` + "`" + `"IKE_NAT"` + "`" + `, ` + "`" + `"IRC"` + "`" + `, ` + "`" + `"LDAP"` + "`" + `, ` + "`" + `"QUIC"` + "`" + `, ` + "`" + `"TDS"` + "`" + `, ` + "`" + `"NETMEETING"` + "`" + `, ` + "`" + `"NFS"` + "`" + `, ` + "`" + `"NTP"` + "`" + `, ` + "`" + `"SIP"` + "`" + `, ` + "`" + `"SNMP"` + "`" + `, ` + "`" + `"SMB"` + "`" + `, ` + "`" + `"SMTP"` + "`" + `, ` + "`" + `"SSH"` + "`" + `, ` + "`" + `"SYSLOG"` + "`" + `, ` + "`" + `"TELNET"` + "`" + `, ` + "`" + `"TRACEROUTE"` + "`" + `, ` + "`" + `"POP3"` + "`" + `, ` + "`" + `"PPTP"` + "`" + `, ` + "`" + `"RADIUS"` + "`" + `, ` + "`" + `"REAL_MEDIA"` + "`" + `, ` + "`" + `"RTSP"` + "`" + `, ` + "`" + `"VNC"` + "`" + `, ` + "`" + `"WHOIS"` + "`" + `, ` + "`" + `"KERBEROS_SEC"` + "`" + `, ` + "`" + `"TACACS"` + "`" + `, ` + "`" + `"SNMPTRAP"` + "`" + `, ` + "`" + `"NMAP"` + "`" + `, ` + "`" + `"RSYNC"` + "`" + `, ` + "`" + `"L2TP"` + "`" + `, ` + "`" + `"HTTP_PROXY"` + "`" + `, ` + "`" + `"PC_ANYWHERE"` + "`" + `, ` + "`" + `"MSN"` + "`" + `, ` + "`" + `"ECHO"` + "`" + `, ` + "`" + `"AIM"` + "`" + `, ` + "`" + `"IDENT"` + "`" + `, ` + "`" + `"YMSG"` + "`" + `, ` + "`" + `"SCCP"` + "`" + `, ` + "`" + `"MGCP_UA"` + "`" + `, ` + "`" + `"MGCP_CA"` + "`" + `, ` + "`" + `"VDO_LIVE"` + "`" + `, ` + "`" + `"OPENVPN"` + "`" + `, ` + "`" + `"TFTP"` + "`" + `, ` + "`" + `"FTPS_IMPLICIT"` + "`" + `, ` + "`" + `"ZSCALER_PROXY_NW_SERVICES"` + "`" + `, ` + "`" + `"GRE_PROTOCOL"` + "`" + `, ` + "`" + `"ESP_PROTOCOL"` + "`" + `, ` + "`" + `"DHCP"` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_firewall_filtering_network_service_groups",
			Category:         "Firewall Policies",
			ShortDescription: `Creates and manages ZIA Cloud firewall Network Service Groups.`,
			Description: `

The **zia_firewall_filtering_network_service_groups** resource allows the creation and management of ZIA Cloud Firewall IP network service groups in the Zscaler Internet Access. This resource can then be associated with a ZIA cloud firewall filtering rule.

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
					Description: `(Required) Name of the network service group`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `(Required) Any number of network services ID to be added to the group ### Optional`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_firewall_filtering_rule",
			Category:         "Firewall Policies",
			ShortDescription: `Creates and manages ZIA Cloud firewall filtering rule.`,
			Description: `

The **zia_firewall_filtering_rule** resource allows the creation and management of ZIA Cloud Firewall filtering rules in the Zscaler Internet Access.

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
					Name:        "order",
					Description: `(Required) Policy rules are evaluated in ascending numerical order (Rule 1 before Rule 2, and so on), and the Rule Order reflects this rule's place in the order. ### Optional`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Enter additional notes or information. The description cannot exceed 10,240 characters.`,
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
					Description: `(Optional) You can manually select up to ` + "`" + `8` + "`" + ` locations. When not used it implies ` + "`" + `Any` + "`" + ` to apply the rule to all groups. - ` + "`" + `id` + "`" + ` - (String) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "location_groups",
					Description: `(Optional) You can manually select up to ` + "`" + `32` + "`" + ` location groups. When not used it implies ` + "`" + `Any` + "`" + ` to apply the rule to all location groups. - ` + "`" + `id` + "`" + ` - (String) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `(Optional) You can manually select up to ` + "`" + `4` + "`" + ` general and/or special users. When not used it implies ` + "`" + `Any` + "`" + ` to apply the rule to all users. - ` + "`" + `id` + "`" + ` - (String) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `(Optional) You can manually select up to ` + "`" + `8` + "`" + ` groups. When not used it implies ` + "`" + `Any` + "`" + ` to apply the rule to all groups. - ` + "`" + `id` + "`" + ` - (String) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "departments",
					Description: `(Optional) Apply to any number of departments When not used it implies ` + "`" + `Any` + "`" + ` to apply the rule to all departments. - ` + "`" + `id` + "`" + ` - (String) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "time_windows",
					Description: `(Optional) You can manually select up to ` + "`" + `2` + "`" + ` time intervals. When not used it implies ` + "`" + `always` + "`" + ` to apply the rule to all time intervals. - ` + "`" + `id` + "`" + ` - (String) Identifier that uniquely identifies an entity ` + "`" + `network services` + "`" + ` supports the following attributes:`,
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
					Description: `(Optional) Any number of source IP address groups that you want to control with this rule. - ` + "`" + `id` + "`" + ` - (String) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "src_ips",
					Description: `(Optional) You can enter individual IP addresses, subnets, or address ranges. ` + "`" + `destinations` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "app_service_groups",
					Description: `Application service groups on which this rule is applied - ` + "`" + `id` + "`" + ` - (String) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "app_services",
					Description: `Application services on which this rule is applied - ` + "`" + `id` + "`" + ` - (String) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "predefined",
					Description: `(Boolean) If set to true, a predefined rule is applied`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_location_management",
			Category:         "Location Management",
			ShortDescription: `Creates and manages ZIA locations and sub-locations.`,
			Description:      ``,
			Keywords: []string{
				"location",
				"management",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) - Location Name.`,
				},
				resource.Attribute{
					Name:        "ip_addresses",
					Description: `(Required) For locations: IP addresses of the egress points that are provisioned in the Zscaler Cloud. Each entry is a single IP address (e.g., ` + "`" + `238.10.33.9` + "`" + `). For sub-locations: Egress, internal, or GRE tunnel IP addresses. Each entry is either a single IP address, CIDR (e.g., ` + "`" + `10.10.33.0/24` + "`" + `), or range (e.g., ` + "`" + `10.10.33.1-10.10.33.10` + "`" + `)). The value is required if ` + "`" + `vpn_credentials` + "`" + ` are not defined.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) VPN credential resource id. The value is required if ` + "`" + `ip_addresses` + "`" + ` are not defined. ### Optional`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(String) Additional notes or information regarding the location or sub-location. The description cannot exceed 1024 characters.`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `(Optional) Country`,
				},
				resource.Attribute{
					Name:        "tz",
					Description: `(Optional) Timezone of the location. If not specified, it defaults to GMT.`,
				},
				resource.Attribute{
					Name:        "profile",
					Description: `(Optional) Profile tag that specifies the location traffic type. If not specified, this tag defaults to ` + "`" + `Unassigned` + "`" + `. The supported options are: ` + "`" + `NONE` + "`" + `, ` + "`" + `CORPORATE` + "`" + `, ` + "`" + `SERVER` + "`" + `, ` + "`" + `GUESTWIFI` + "`" + `, ` + "`" + `IOT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "aup_block_internet_until_accepted",
					Description: `(Optional) For First Time AUP Behavior, Block Internet Access. When set, all internet access (including non-HTTP traffic) is disabled until the user accepts the AUP.`,
				},
				resource.Attribute{
					Name:        "aup_enabled",
					Description: `(Optional) Enable AUP. When set to true, AUP is enabled for the location.`,
				},
				resource.Attribute{
					Name:        "aup_force_ssl_inspection",
					Description: `(Optional) For First Time AUP Behavior, Force SSL Inspection. When set, Zscaler will force SSL Inspection in order to enforce AUP for HTTPS traffic.`,
				},
				resource.Attribute{
					Name:        "aup_timeout_in_days",
					Description: `(Optional) Custom AUP Frequency. Refresh time (in days) to re-validate the AUP.`,
				},
				resource.Attribute{
					Name:        "auth_required",
					Description: `(Optional) Enforce Authentication. Required when ports are enabled, IP Surrogate is enabled, or Kerberos Authentication is enabled.`,
				},
				resource.Attribute{
					Name:        "caution_enabled",
					Description: `(Optional) Enable Caution. When set to true, a caution notifcation is enabled for the location.`,
				},
				resource.Attribute{
					Name:        "display_time_unit",
					Description: `(Optional) Display Time Unit. The time unit to display for IP Surrogate idle time to disassociation.`,
				},
				resource.Attribute{
					Name:        "dn_bandwidth",
					Description: `(Optional) Download bandwidth in bytes. The value ` + "`" + `0` + "`" + ` implies no Bandwidth Control enforcement.`,
				},
				resource.Attribute{
					Name:        "up_bandwidth",
					Description: `(Optional) Upload bandwidth in bytes. The value ` + "`" + `0` + "`" + ` implies no Bandwidth Control enforcement.`,
				},
				resource.Attribute{
					Name:        "idle_time_in_minutes",
					Description: `(Optional) Idle Time to Disassociation. The user mapping idle time (in minutes) is required if a Surrogate IP is enabled.`,
				},
				resource.Attribute{
					Name:        "ips_control",
					Description: `(Optional) Enable IPS Control. When set to true, IPS Control is enabled for the location if Firewall is enabled.`,
				},
				resource.Attribute{
					Name:        "ofw_enabled",
					Description: `(Optional) Enable Firewall. When set to true, Firewall is enabled for the location.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `(Optional) - Parent Location ID. If this ID does not exist or is ` + "`" + `0` + "`" + `, it is implied that it is a parent location. Otherwise, it is a sub-location whose parent has this ID. x-applicableTo: ` + "`" + `SUB` + "`" + ``,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `(Optional) IP ports that are associated with the location.`,
				},
				resource.Attribute{
					Name:        "ssl_scan_enabled",
					Description: `(Optional) This parameter was deprecated and no longer has an effect on SSL policy. It remains supported in the API payload in order to maintain backwards compatibility with existing scripts, but it will be removed in future.`,
				},
				resource.Attribute{
					Name:        "surrogate_ip",
					Description: `(Optional) Enable Surrogate IP. When set to true, users are mapped to internal device IP addresses.`,
				},
				resource.Attribute{
					Name:        "surrogate_ip_enforced_for_known_browsers",
					Description: `(Optional) Enforce Surrogate IP for Known Browsers. When set to true, IP Surrogate is enforced for all known browsers.`,
				},
				resource.Attribute{
					Name:        "surrogate_refresh_time_in_minutes",
					Description: `(Optional) Refresh Time for re-validation of Surrogacy. The surrogate refresh time (in minutes) to re-validate the IP surrogates.`,
				},
				resource.Attribute{
					Name:        "surrogate_refresh_time_unit",
					Description: `(Optional) Display Refresh Time Unit. The time unit to display for refresh time for re-validation of surrogacy.`,
				},
				resource.Attribute{
					Name:        "xff_forward_enabled",
					Description: `(Optional) Enable XFF Forwarding. When set to true, traffic is passed to Zscaler Cloud via the X-Forwarded-For (XFF) header.`,
				},
				resource.Attribute{
					Name:        "zapp_ssl_scan_enabled",
					Description: `(Optional) This parameter was deprecated and no longer has an effect on SSL policy. It remains supported in the API payload in order to maintain backwards compatibility with existing scripts, but it will be removed in future.`,
				},
				resource.Attribute{
					Name:        "managed_by",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The configured name of the entity`,
				},
				resource.Attribute{
					Name:        "extensions",
					Description: `(Optional)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_rule_labels",
			Category:         "Rule Labels",
			ShortDescription: `Creates and manages ZIA rule labels.`,
			Description: `

The **zia_rule_labels** resource allows the creation and management of rule labels in the Zscaler Internet Access cloud or via the API. This resource can then be associated with resources such as: Firewall Rules and URL filtering rules

`,
			Keywords: []string{
				"rule",
				"labels",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the devices to be created. ### Optional`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(String) The rule label description.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_security_policy_settings",
			Category:         "Security Policy Settings",
			ShortDescription: `Add or Remove a URL to and from the allow and deny list in the advanced threat protection.`,
			Description: `

The **zia_security_policy_settings** resource alows you to add or remove a URL to the allow and denylist under the Advanced Threat Protection policy in the Zscaler Internet Access cloud or via the API.

`,
			Keywords: []string{
				"security",
				"policy",
				"settings",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "whitelist_urls",
					Description: `(Required) Allowlist URLs whose contents will not be scanned. Allows up to 255 URLs`,
				},
				resource.Attribute{
					Name:        "blacklist_urls",
					Description: `(Optional) URLs on the denylist for your organization. Allow up to 25000 URLs.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_traffic_forwarding_gre_tunnel",
			Category:         "Traffic Forwarding",
			ShortDescription: `Creates and manages GRE tunnel configuration.`,
			Description: `

The **zia_traffic_forwarding_gre_tunnel** resource allows the creation and management of GRE tunnel configuration in the Zscaler Internet Access (ZIA) portal.

-> **Note:** The provider automatically query the Zscaler cloud for the primary and secondary destination datacenter and virtual IP address (VIP) of the GRE tunnel. The parameter can be overriden if needed by setting the parameters: ` + "`" + `primary_dest_vip` + "`" + ` and ` + "`" + `secondary_dest_vip` + "`" + `.

`,
			Keywords: []string{
				"traffic",
				"forwarding",
				"gre",
				"tunnel",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Unique identifer of the GRE virtual IP address (VIP)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Unique identifer of the GRE virtual IP address (VIP)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_traffic_forwarding_static_ip",
			Category:         "Traffic Forwarding",
			ShortDescription: `Creates and manages static IP addresses.`,
			Description: `

The **zia_traffic_forwarding_static_ip** resource allows the creation and management of static ip addresses in the Zscaler Internet Access cloud. The resource, can then be associated with other resources such as:

* VPN Credentials of type ` + "`" + `IP` + "`" + `
* Location Management
* GRE Tunnel

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
					Description: `(Required) The static IP address ### Optional`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Additional information about this static IP address`,
				},
				resource.Attribute{
					Name:        "latitude",
					Description: `(Optional) Required only if the geoOverride attribute is set. Latitude with 7 digit precision after decimal point, ranges between -90 and 90 degrees.`,
				},
				resource.Attribute{
					Name:        "longitude",
					Description: `(Optional) Required only if the geoOverride attribute is set. Longitude with 7 digit precision after decimal point, ranges between -180 and 180 degrees.`,
				},
				resource.Attribute{
					Name:        "geo_override",
					Description: `(Optional) If not set, geographic coordinates and city are automatically determined from the IP address. Otherwise, the latitude and longitude coordinates must be provided.`,
				},
				resource.Attribute{
					Name:        "routable_ip",
					Description: `(Optional) Indicates whether a non-RFC 1918 IP address is publicly routable. This attribute is ignored if there is no ZIA Private Service Edge associated to the organization.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ## Import Static IP resources can be imported by using ` + "`" + `<STATIC IP ID>` + "`" + ` or ` + "`" + `<IP ADDRESS>` + "`" + `as the import ID. ` + "`" + `` + "`" + `` + "`" + `shell terraform import zia_traffic_forwarding_static_ip.example <static_ip_id> ` + "`" + `` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `` + "`" + `shell terraform import zpa_app_connector_group.example <ip_address> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_traffic_forwarding_vpn_credentials",
			Category:         "Traffic Forwarding",
			ShortDescription: `Creates and manages VPN credentials that can be associated to locations.`,
			Description: `

The **zia_traffic_forwarding_vpn_credentials** creates and manages VPN credentials that can be associated to locations. VPN is one way to route traffic from customer locations to the cloud. Site-to-site IPSec VPN credentials can be identified by the cloud through one of the following methods:

* Common Name (CN) of IPSec Certificate
* VPN User FQDN - requires VPN_SITE_TO_SITE subscription
* VPN IP Address - requires VPN_SITE_TO_SITE subscription
* Extended Authentication (XAUTH) or hosted mobile UserID - requires VPN_MOBILE subscription

`,
			Keywords: []string{
				"traffic",
				"forwarding",
				"vpn",
				"credentials",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `(Required) VPN authentication type (i.e., how the VPN credential is sent to the server). It is not modifiable after VpnCredential is created. The supported values are: ` + "`" + `UFQDN` + "`" + ` and ` + "`" + `IP` + "`" + ``,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `(Required) Fully Qualified Domain Name. Applicable only to ` + "`" + `UFQDN` + "`" + ` or ` + "`" + `XAUTH` + "`" + ` (or ` + "`" + `HOSTED_MOBILE_USERS` + "`" + `) auth type.`,
				},
				resource.Attribute{
					Name:        "pre_shared_key",
					Description: `(Required) Pre-shared key. This is a required field for UFQDN and IP auth type. ### Optional`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Additional information about this VPN credential.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) IP Address for the VON credentials. The parameter becomes required if ` + "`" + `type = IP` + "`" + ` !>`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_url_categories",
			Category:         "URL Categories",
			ShortDescription: `Creates and manages a new custom URL category. If keywords are included within the request, they will be added to the new category.`,
			Description: `

The **zia_url_categories** resource creates and manages a new custom URL category. If keywords are included within the request, they will be added to the new category.

`,
			Keywords: []string{
				"url",
				"categories",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "configured_name",
					Description: `(Required) Name of the URL category. This is only required for custom URL categories.`,
				},
				resource.Attribute{
					Name:        "super_category",
					Description: `(Required) ### Optional`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the category.`,
				},
				resource.Attribute{
					Name:        "keywords",
					Description: `(Optional) Custom keywords associated to a URL category. Up to 2048 custom keywords can be added per organization across all categories (including bandwidth classes).`,
				},
				resource.Attribute{
					Name:        "custom_category",
					Description: `(Boolean) Set to true for custom URL category. Up to 48 custom URL categories can be added per organization.`,
				},
				resource.Attribute{
					Name:        "custom_urls_count",
					Description: `(Optional) The number of custom URLs associated to the URL category.`,
				},
				resource.Attribute{
					Name:        "db_categorized_urls",
					Description: `(Optional) URLs added to a custom URL category are also retained under the original parent URL category (i.e., the predefined category the URL previously belonged to).`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `(Optional) Custom IP address ranges associated to a URL category. Up to 2000 custom IP address ranges and retaining parent custom IP address ranges can be added, per organization, across all categories.`,
				},
				resource.Attribute{
					Name:        "ip_ranges_retaining_parent_category",
					Description: `(Optional) The retaining parent custom IP address ranges associated to a URL category. Up to 2000 custom IP ranges and retaining parent custom IP address ranges can be added, per organization, across all categories.`,
				},
				resource.Attribute{
					Name:        "ip_ranges_retaining_parent_category_count",
					Description: `(Optional) The number of custom IP address ranges associated to the URL category, that also need to be retained under the original parent category.`,
				},
				resource.Attribute{
					Name:        "custom_ip_ranges_count",
					Description: `(Optional) The number of custom IP address ranges associated to the URL category.`,
				},
				resource.Attribute{
					Name:        "editable",
					Description: `(Boolean) Value is set to false for custom URL category when due to scope user does not have edit permission`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type of the custom categories. ` + "`" + `URL_CATEGORY` + "`" + `, ` + "`" + `TLD_CATEGORY` + "`" + `, ` + "`" + `ALL` + "`" + ``,
				},
				resource.Attribute{
					Name:        "urls",
					Description: `(Optional) Custom URLs to add to a URL category. Up to 25,000 custom URLs can be added per organization across all categories (including bandwidth classes).`,
				},
				resource.Attribute{
					Name:        "urls_retaining_parent_category_count",
					Description: `(Optional) The number of custom IP address ranges associated to the URL category, that also need to be retained under the original parent category.`,
				},
				resource.Attribute{
					Name:        "scopes",
					Description: `(Optional) Scope of the custom categories.`,
				},
				resource.Attribute{
					Name:        "scope_group_member_entities",
					Description: `(List of Object) Only applicable for the LOCATION_GROUP admin scope type, in which case this attribute gives the list of ID/name pairs of locations within the location group. The attribute name is subject to change.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The admin scope type. The attribute name is subject to change. ` + "`" + `ORGANIZATION` + "`" + `, ` + "`" + `DEPARTMENT` + "`" + `, ` + "`" + `LOCATION` + "`" + `, ` + "`" + `LOCATION_GROUP` + "`" + ``,
				},
				resource.Attribute{
					Name:        "scope_entities",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "url_keyword_counts",
					Description: `(Optional) URL and keyword counts for the category.`,
				},
				resource.Attribute{
					Name:        "total_url_count",
					Description: `(Optional) Custom URL count for the category.`,
				},
				resource.Attribute{
					Name:        "retain_parent_url_count",
					Description: `(Optional) Count of URLs with retain parent category.`,
				},
				resource.Attribute{
					Name:        "total_keyword_count",
					Description: `(Optional) Total keyword count for the category.`,
				},
				resource.Attribute{
					Name:        "retain_parent_keyword_count",
					Description: `(Optional) Count of total keywords with retain parent category.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_url_filtering_rules",
			Category:         "URL Filtering Rule",
			ShortDescription: `Creates and manages a URL Filtering Policy rule.`,
			Description: `

The **zia_url_filtering_rules** resource creates and manages a URL filtering rules within the Zscaler Internet Access cloud.

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
					Description: `(Required) Name of the Firewall Filtering policy rule`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Required) Order of execution of rule with respect to other URL Filtering rules`,
				},
				resource.Attribute{
					Name:        "protocols",
					Description: `(Required) Protocol criteria. Supported values: ` + "`" + `SMRULEF_ZPA_BROKERS_RULE` + "`" + `, ` + "`" + `ANY_RULE` + "`" + `, ` + "`" + `TCP_RULE` + "`" + `, ` + "`" + `UDP_RULE` + "`" + `, ` + "`" + `DOHTTPS_RULE` + "`" + `, ` + "`" + `TUNNELSSL_RULE` + "`" + `, ` + "`" + `HTTP_PROXY` + "`" + `, ` + "`" + `FOHTTP_RULE` + "`" + `, ` + "`" + `FTP_RULE` + "`" + `, ` + "`" + `HTTPS_RULE` + "`" + `, ` + "`" + `HTTP_RULE` + "`" + `, ` + "`" + `SSL_RULE` + "`" + `, ` + "`" + `TUNNEL_RULE` + "`" + `. ### Optional`,
				},
				resource.Attribute{
					Name:        "request_methods",
					Description: `(Optional) Request method for which the rule must be applied. If not set, rule will be applied to all methods`,
				},
				resource.Attribute{
					Name:        "rank",
					Description: `(Optional) Admin rank of the admin who creates this rule`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) Rule State`,
				},
				resource.Attribute{
					Name:        "end_user_notification_url",
					Description: `(Optional) URL of end user notification page to be displayed when the rule is matched. Not applicable if either 'overrideUsers' or 'overrideGroups' is specified.`,
				},
				resource.Attribute{
					Name:        "block_override",
					Description: `(Optional) When set to true, a ` + "`" + `BLOCK` + "`" + ` action triggered by the rule could be overridden. If true and both overrideGroup and overrideUsers are not set, the ` + "`" + `BLOCK` + "`" + ` triggered by this rule could be overridden for any users. If block)Override is not set, ` + "`" + `BLOCK` + "`" + ` action cannot be overridden.`,
				},
				resource.Attribute{
					Name:        "time_quota",
					Description: `(Optional) Time quota in minutes, after which the URL Filtering rule is applied. If not set, no quota is enforced. If a policy rule action is set to ` + "`" + `BLOCK` + "`" + `, this field is not applicable.`,
				},
				resource.Attribute{
					Name:        "size_quota",
					Description: `(Optional) Size quota in KB beyond which the URL Filtering rule is applied. If not set, no quota is enforced. If a policy rule action is set to ` + "`" + `BLOCK` + "`" + `, this field is not applicable.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Additional information about the rule`,
				},
				resource.Attribute{
					Name:        "validity_start_time",
					Description: `(Optional) If enforceTimeValidity is set to true, the URL Filtering rule will be valid starting on this date and time.`,
				},
				resource.Attribute{
					Name:        "validity_end_time",
					Description: `(Optional) If ` + "`" + `enforceTimeValidity` + "`" + ` is set to true, the URL Filtering rule will cease to be valid on this end date and time.`,
				},
				resource.Attribute{
					Name:        "validity_time_zone_id",
					Description: `(Optional) If ` + "`" + `enforceTimeValidity` + "`" + ` is set to true, the URL Filtering rule date and time will be valid based on this time zone ID.`,
				},
				resource.Attribute{
					Name:        "last_modified_time",
					Description: `(Optional) When the rule was last modified`,
				},
				resource.Attribute{
					Name:        "enforce_time_validity",
					Description: `(Optional) Enforce a set a validity time period for the URL Filtering rule.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Action taken when traffic matches rule criteria. Supported values: ` + "`" + `ANY` + "`" + `, ` + "`" + `NONE` + "`" + `, ` + "`" + `BLOCK` + "`" + `, ` + "`" + `CAUTION` + "`" + `, ` + "`" + `ALLOW` + "`" + `, ` + "`" + `ICAP_RESPONSE` + "`" + ``,
				},
				resource.Attribute{
					Name:        "device_trust_levels",
					Description: `(Optional) List of device trust levels for which the rule must be applied. This field is applicable for devices that are managed using Zscaler Client Connector. The trust levels are assigned to the devices based on your posture configurations in the Zscaler Client Connector Portal. If no value is set, this field is ignored during the policy evaluation. Supported values: ` + "`" + `ANY` + "`" + `, ` + "`" + `UNKNOWN_DEVICETRUSTLEVEL` + "`" + `, ` + "`" + `LOW_TRUST` + "`" + `, ` + "`" + `MEDIUM_TRUST` + "`" + `, ` + "`" + `HIGH_TRUST` + "`" + ``,
				},
				resource.Attribute{
					Name:        "cipa_rule",
					Description: `(Optional) If set to true, the CIPA Compliance rule is enabled`,
				},
				resource.Attribute{
					Name:        "url_categories",
					Description: `(Optional) List of URL categories for which rule must be applied`,
				},
				resource.Attribute{
					Name:        "locations",
					Description: `(List of Object) The locations to which the Firewall Filtering policy rule applies`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `(List of Object) The groups to which the Firewall Filtering policy rule applies`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "departments",
					Description: `(List of Object) The departments to which the Firewall Filtering policy rule applies`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Identifier that uniquely identifies an entity`,
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
					Name:        "time_windows",
					Description: `(List of Object) The time interval in which the Firewall Filtering policy rule applies`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "override_users",
					Description: `(List of Object) Name-ID pairs of users for which this rule can be overridden. Applicable only if blockOverride is set to ` + "`" + `true` + "`" + `, action is ` + "`" + `BLOCK` + "`" + ` and overrideGroups is not set.If this overrideUsers is not set, ` + "`" + `BLOCK` + "`" + ` action can be overridden for any user.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "override_groups",
					Description: `(List of Object) Name-ID pairs of users for which this rule can be overridden. Applicable only if blockOverride is set to ` + "`" + `true` + "`" + `, action is ` + "`" + `BLOCK` + "`" + ` and overrideGroups is not set.If this overrideUsers is not set, ` + "`" + `BLOCK` + "`" + ` action can be overridden for any group.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "location_groups",
					Description: `(List of Object) The location groups to which the Firewall Filtering policy rule applies`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Identifier that uniquely identifies an entity`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Identifier that uniquely identifies an entity`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zia_zia_user_management",
			Category:         "User Management",
			ShortDescription: `Creates and manages ZIA local user accounts.`,
			Description:      ``,
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
					Name:        "email",
					Description: `(Required) User email consists of a user name and domain name. It does not have to be a valid email address, but it must be unique and its domain must belong to the organization.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) User's password. Applicable only when authentication type is Hosted DB. Password strength must follow what is defined in the auth settings.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `(Required) List of Groups a user belongs to. Groups are used in policies.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) Unique identfier for the group`,
				},
				resource.Attribute{
					Name:        "department",
					Description: `(Required) Department a user belongs to`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) Department ID !>`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Additional information about this user.`,
				},
				resource.Attribute{
					Name:        "temp_auth_email",
					Description: `(Optional) Temporary Authentication Email. If you enabled one-time tokens or links, enter the email address to which the Zscaler service sends the tokens or links. If this is empty, the service will send the email to the User email.`,
				},
				resource.Attribute{
					Name:        "auth_methods",
					Description: `(Optional) Type of authentication method to be enabled. Supported values are: ` + "`" + `` + "`" + `BASIC` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `DIGEST` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"zia_zia_activation_status":                             0,
		"zia_zia_admin_users":                                   1,
		"zia_zia_auth_settings_urls":                            2,
		"zia_zia_dlp_dictionaries":                              3,
		"zia_zia_dlp_notification_templates":                    4,
		"zia_zia_dlp_web_rules":                                 5,
		"zia_zia_firewall_filtering_ip_destination_groups":      6,
		"zia_zia_firewall_filtering_ip_source_groups":           7,
		"zia_zia_firewall_filtering_network_application_groups": 8,
		"zia_zia_firewall_filtering_network_service":            9,
		"zia_zia_firewall_filtering_network_service_groups":     10,
		"zia_zia_firewall_filtering_rule":                       11,
		"zia_zia_location_management":                           12,
		"zia_zia_rule_labels":                                   13,
		"zia_zia_security_policy_settings":                      14,
		"zia_zia_traffic_forwarding_gre_tunnel":                 15,
		"zia_zia_traffic_forwarding_static_ip":                  16,
		"zia_zia_traffic_forwarding_vpn_credentials":            17,
		"zia_zia_url_categories":                                18,
		"zia_zia_url_filtering_rules":                           19,
		"zia_zia_user_management":                               20,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
