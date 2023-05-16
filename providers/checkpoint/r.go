package checkpoint

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_hostname",
			Category:         "GAIA Resources",
			ShortDescription: `This resource allows you to set the hostname of a Check Point machine.`,
			Description: `

This resource allows you to set the hostname of a Check Point machine.

`,
			Keywords: []string{
				"gaia",
				"hostname",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) New hostname to change.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_access_layer",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Access Layer.`,
			Description: `

This resource allows you to execute Check Point Access Layer.

`,
			Keywords: []string{
				"management",
				"access",
				"layer",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "add_default_rule",
					Description: `(Optional) Indicates whether to include a cleanup rule in the new layer.`,
				},
				resource.Attribute{
					Name:        "applications_and_url_filtering",
					Description: `(Optional) Whether to enable Applications & URL Filtering blade on the layer.`,
				},
				resource.Attribute{
					Name:        "content_awareness",
					Description: `(Optional) Whether to enable Content Awareness blade on the layer.`,
				},
				resource.Attribute{
					Name:        "detect_using_x_forward_for",
					Description: `(Optional) Whether to use X-Forward-For HTTP header, which is added by the proxy server to keep track of the original source IP.`,
				},
				resource.Attribute{
					Name:        "firewall",
					Description: `(Optional) Whether to enable Firewall blade on the layer.`,
				},
				resource.Attribute{
					Name:        "implicit_cleanup_action",
					Description: `(Optional) The default "catch-all" action for traffic that does not match any explicit or implied rules in the layer.`,
				},
				resource.Attribute{
					Name:        "mobile_access",
					Description: `(Optional) Whether to enable Mobile Access blade on the layer.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `(Optional) Whether this layer is shared.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_access_point_name",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Access Point Name.`,
			Description: `

This resource allows you to execute Check Point Access Point Name.

`,
			Keywords: []string{
				"management",
				"access",
				"point",
				"name",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "apn",
					Description: `(Optional) APN name.`,
				},
				resource.Attribute{
					Name:        "enforce_end_user_domain",
					Description: `(Optional) Enable enforce end user domain.`,
				},
				resource.Attribute{
					Name:        "block_traffic_other_end_user_domains",
					Description: `(Optional) Block MS to MS traffic between this and other APN end user domains.`,
				},
				resource.Attribute{
					Name:        "block_traffic_this_end_user_domain",
					Description: `(Optional) Block MS to MS traffic within this end user domain.`,
				},
				resource.Attribute{
					Name:        "end_user_domain",
					Description: `(Optional) End user domain name or UID.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_access_role",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Access Role.`,
			Description: `

This resource allows you to execute Check Point Access Role.

`,
			Keywords: []string{
				"management",
				"access",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "machines",
					Description: `(Optional) Machines that can access the system.machines blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "networks",
					Description: `(Optional) Collection of Network objects identified by the name or UID that can access the system.networks blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "remote_access_clients",
					Description: `(Optional) Remote access clients identified by name or UID.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `(Optional) Users that can access the system.users blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. ` + "`" + `machines` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) any, all identified, Active Directory name or UID or Identity Tag. default value = "any"`,
				},
				resource.Attribute{
					Name:        "selection",
					Description: `(Optional) Name or UID of an object selected from source. selection blocks are documented below. default value = ["any"]`,
				},
				resource.Attribute{
					Name:        "base_dn",
					Description: `(Optional) When source is "Active Directory" use "base-dn" to refine the query in AD database. ` + "`" + `users` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) any, all identified, UID or Identity Tag or Internal User Groups or LDAP groups or Guests. default value = "any", supports only one AD group named CpmiAdGroup due to API limitations.`,
				},
				resource.Attribute{
					Name:        "selection",
					Description: `(Optional) Name or UID of an object selected from source. selection blocks are documented below. default value = ["any"], on ad groups: Adds ad_group prefix to the selection, and removes spaces due to API limitations.`,
				},
				resource.Attribute{
					Name:        "base_dn",
					Description: `(Optional) When source is "Active Directory" use "base-dn" to refine the query in AD database.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_access_rule",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to add/update/delete Check Point Access Rule.`,
			Description: `

This resource allows you to add/update/delete Check Point Access Rule.

`,
			Keywords: []string{
				"management",
				"access",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "layer",
					Description: `(Required) Layer that the rule belongs to identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "position",
					Description: `(Required) Position in the rulebase. Position blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Rule name.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) \"Accept\", \"Drop\", \"Ask\", \"Inform\", \"Reject\", \"User Auth\", \"Client Auth\", \"Apply Layer\".`,
				},
				resource.Attribute{
					Name:        "action_settings",
					Description: `(Optional) Action settings. Action settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Optional) List of processed file types that this rule applies on.`,
				},
				resource.Attribute{
					Name:        "content_direction",
					Description: `(Optional) On which direction the file types processing is applied.`,
				},
				resource.Attribute{
					Name:        "content_negate",
					Description: `(Optional) True if negate is set for data.`,
				},
				resource.Attribute{
					Name:        "custom_fields",
					Description: `(Optional) Custom fields. Custom fields blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `(Optional) Collection of Network objects identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "destination_negate",
					Description: `(Optional) True if negate is set for destination.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Enable/Disable the rule.`,
				},
				resource.Attribute{
					Name:        "inline_layer",
					Description: `(Optional) Inline Layer identified by the name or UID. Relevant only if \"Action\" was set to \"Apply Layer\".`,
				},
				resource.Attribute{
					Name:        "install_on",
					Description: `(Optional) Which Gateways identified by the name or UID to install the policy on.`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `(Optional) Collection of Network objects identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "service_negate",
					Description: `(Optional) True if negate is set for service.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) Collection of Network objects identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "source_negate",
					Description: `(Optional) True if negate is set for source.`,
				},
				resource.Attribute{
					Name:        "time",
					Description: `(Optional) List of time objects. For example: \"Weekend\", \"Off-Work\", \"Every-Day\".`,
				},
				resource.Attribute{
					Name:        "track",
					Description: `(Optional) Track Settings. Track Settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "user_check",
					Description: `(Optional) User check settings. User check settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "vpn",
					Description: `(Optional) Communities or Directional.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string. ` + "`" + `position` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "top",
					Description: `(Optional) Add rule at the top of the rulebase.`,
				},
				resource.Attribute{
					Name:        "above",
					Description: `(Optional) Add rule above specific section/rule identified by uid or name.`,
				},
				resource.Attribute{
					Name:        "below",
					Description: `(Optional) Add rule below specific section/rule identified by uid or name.`,
				},
				resource.Attribute{
					Name:        "bottom",
					Description: `(Optional) Add rule at the bottom of the rulebase. ` + "`" + `action_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "enable_identity_captive_portal",
					Description: `(Optional) N/A.`,
				},
				resource.Attribute{
					Name:        "limit",
					Description: `(Optional) N/A. ` + "`" + `custom_fields` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "field_1",
					Description: `(Optional) First custom field.`,
				},
				resource.Attribute{
					Name:        "field_2",
					Description: `(Optional) Second custom field.`,
				},
				resource.Attribute{
					Name:        "field_3",
					Description: `(Optional) Third custom field. ` + "`" + `track` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "accounting",
					Description: `(Optional) Turns accounting for track on and off.`,
				},
				resource.Attribute{
					Name:        "alert",
					Description: `(Optional) Type of alert for the track.`,
				},
				resource.Attribute{
					Name:        "enable_firewall_session",
					Description: `(Optional) Determine whether to generate session log to firewall only connections.`,
				},
				resource.Attribute{
					Name:        "per_connection",
					Description: `(Optional) Determines whether to perform the log per connection.`,
				},
				resource.Attribute{
					Name:        "per_session",
					Description: `(Optional) Determines whether to perform the log per session.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) \"Log\", \"Extended Log\", \"Detailed Log\", \"None\". ` + "`" + `user_check` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "confirm",
					Description: `(Optional) N/A.`,
				},
				resource.Attribute{
					Name:        "custom_frequency",
					Description: `(Optional) N/A. Custom Frequency blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "frequency",
					Description: `(Optional) N/A.`,
				},
				resource.Attribute{
					Name:        "interaction",
					Description: `(Optional) N/A. ` + "`" + `custom_frequency` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "every",
					Description: `(Optional) N/A.`,
				},
				resource.Attribute{
					Name:        "unit",
					Description: `(Optional) N/A. ## Import ` + "`" + `checkpoint_management_access_rule` + "`" + ` can be imported by using the following format: LAYER_NAME;RULE_UID ` + "`" + `` + "`" + `` + "`" + ` $ terraform import checkpoint_management_access_rule.example "Network;9423d36f-2d66-4754-b9e2-e9f4493751d3" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_access_section",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Access Section.`,
			Description: `

This resource allows you to execute Check Point Access Section.

`,
			Keywords: []string{
				"management",
				"access",
				"section",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "layer",
					Description: `(Required) Layer that the rule belongs to identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
				resource.Attribute{
					Name:        "position",
					Description: `(Required) Position in the rulebase. ## Import ` + "`" + `checkpoint_management_access_section` + "`" + ` can be imported by using the following format: LAYER_NAME;SECTION_UID ` + "`" + `` + "`" + `` + "`" + ` $ terraform import checkpoint_management_access_section.example "Network;354e184c-2f42-485c-b62d-ff9b3d29ee3e" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_aci_data_center_server",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point aci data center server.`,
			Description: `

This resource allows you to execute Check Point Cisco APIC Data Center Server.

`,
			Keywords: []string{
				"management",
				"aci",
				"data",
				"center",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "urls",
					Description: `(Required) Address of APIC cluster members. Example: http(s)://<host1 ip/url>.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) User ID of the Cisco APIC server. When using Login Domains use the following syntax:apic:<domain>\<username>.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Password of the Cisco APIC server.`,
				},
				resource.Attribute{
					Name:        "password_base64",
					Description: `(Optional) Password of the Cisco APIC server encoded in Base64.`,
				},
				resource.Attribute{
					Name:        "certificate_fingerprint",
					Description: `(Optional) Specify the SHA-1 or SHA-256 fingerprint of the Data Center Server's certificate.`,
				},
				resource.Attribute{
					Name:        "unsafe_auto_accept",
					Description: `(Optional) When set to false, the current Data Center Server's certificate should be trusted, either by providing the certificate-fingerprint argument or by relying on a previously trusted certificate of this hostname. When set to true, trust the current Data Center Server's certificate as-is.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers. tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_add_api_key",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Add Api Key.`,
			Description: `

This command resource allows you to execute Check Point Add Api Key.

`,
			Keywords: []string{
				"management",
				"add",
				"api",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "admin_uid",
					Description: `(Required) Administrator uid to generate API key for.`,
				},
				resource.Attribute{
					Name:        "admin_name",
					Description: `(Required) Administrator name to generate API key for.`,
				},
				resource.Attribute{
					Name:        "api_key",
					Description: `(Computed) Represents the API Key to be used for Login. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_add_data_center_object",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Add Data Center Object.`,
			Description: `

This command resource allows you to execute Check Point Add Data Center Object.

`,
			Keywords: []string{
				"management",
				"add",
				"data",
				"center",
				"object",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "data_center_name",
					Description: `(Required) Name of the Data Center Server the object is in.`,
				},
				resource.Attribute{
					Name:        "data_center_uid",
					Description: `(Required) Unique identifier of the Data Center Server the object is in.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `(Required) URI of the object in the Data Center Server.`,
				},
				resource.Attribute{
					Name:        "uid_in_data_center",
					Description: `(Required) Unique identifier of the object in the Data Center Server.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Override default name on data-center.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `(Optional) Collection of group identifiers.groups blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_add_repository_package",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Add Repository Package.`,
			Description:      ``,
			Keywords: []string{
				"management",
				"add",
				"repository",
				"package",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the repository package.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The path of the repository package.<br><font color="red">Required only for</font> adding package from local.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Required) The source of the repository package. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_add_threat_protections",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Add Threat Protections.`,
			Description: `

This command resource allows you to execute Check Point Add Threat Protections.

`,
			Keywords: []string{
				"management",
				"add",
				"threat",
				"protections",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "package_format",
					Description: `(Optional) Protections package format.`,
				},
				resource.Attribute{
					Name:        "package_path",
					Description: `(Optional) Protections package path.`,
				},
				resource.Attribute{
					Name:        "task_id",
					Description: `(Computed) Asynchronous task unique identifier. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_add_updatable_object",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Add Updatable Object.`,
			Description: `

This command resource allows you to execute Check Point Add Updatable Object.

`,
			Keywords: []string{
				"management",
				"add",
				"updatable",
				"object",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "uri",
					Description: `(Required) URI of the updatable object in the Updatable Objects Repository.`,
				},
				resource.Attribute{
					Name:        "uid_in_updatable_objects_repository",
					Description: `(Required) Unique identifier of the updatable object in the Updatable Objects Repository.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_address_range",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to add/update/delete Check Point Address Range.`,
			Description: `

This resource allows you to add/update/delete Check Point Address Range.

`,
			Keywords: []string{
				"management",
				"address",
				"range",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name. Should be unique in the domain.`,
				},
				resource.Attribute{
					Name:        "ipv4_address_first",
					Description: `(Optional) First IPv4 address in the range.`,
				},
				resource.Attribute{
					Name:        "ipv6_address_first",
					Description: `(Optional) First IPv6 address in the range.`,
				},
				resource.Attribute{
					Name:        "ipv4_address_last",
					Description: `(Optional) Last IPv4 address in the range.`,
				},
				resource.Attribute{
					Name:        "ipv6_address_last",
					Description: `(Optional) Last IPv6 address in the range.`,
				},
				resource.Attribute{
					Name:        "nat_settings",
					Description: `(Optional) NAT settings. NAT settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers. ` + "`" + `nat_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "auto_rule",
					Description: `(Optional) Whether to add automatic address translation rules.`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `(Optional) IPv4 address.`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `(Optional) IPv6 address.`,
				},
				resource.Attribute{
					Name:        "hide_behind",
					Description: `(Optional) Hide behind method. This parameter is not required in case \"method\" parameter is \"static\".`,
				},
				resource.Attribute{
					Name:        "install_on",
					Description: `(Optional) Which gateway should apply the NAT translation.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Optional) NAT translation method.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_administrator",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to add/update/delete Check Point Administrator.`,
			Description: `

This resource allows you to add/update/delete Check Point Administrator.

`,
			Keywords: []string{
				"management",
				"administrator",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name. Should be unique in the domain.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Administrator password.`,
				},
				resource.Attribute{
					Name:        "authentication_method",
					Description: `(Required) Authentication method.`,
				},
				resource.Attribute{
					Name:        "permission_profile",
					Description: `(Required) Administrator permissions profile. Permissions profile should not be provided when multi-domain-profile is set to "Multi-Domain Super User" or "Domain Super User". In SMC, permissions_profile with single object, domain must be "SMC User".`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) Administrator email.`,
				},
				resource.Attribute{
					Name:        "expiration_date",
					Description: `(Optional) Format: YYYY-MM-DD, YYYY-mm-ddThh:mm:ss.`,
				},
				resource.Attribute{
					Name:        "multi_domain_profile",
					Description: `(Optional) Administrator multi-domain profile. Only in MDS.`,
				},
				resource.Attribute{
					Name:        "must_change_password",
					Description: `(Optional) True if administrator must change password on the next login.`,
				},
				resource.Attribute{
					Name:        "phone_number",
					Description: `(Optional) Administrator phone number.`,
				},
				resource.Attribute{
					Name:        "radius_server",
					Description: `(Optional) RADIUS server object identified by the name or UID. Must be set when "authentication-method" was selected to be "RADIUS".`,
				},
				resource.Attribute{
					Name:        "tacacs_server",
					Description: `(Optional) TACACS server object identified by the name or UID. Must be set when "authentication-method" was selected to be "TACACS".`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
				resource.Attribute{
					Name:        "sic_name",
					Description: `Name of the Secure Internal Connection Trust.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_application_site",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Application Site.`,
			Description: `

This resource allows you to execute Check Point Application Site.

`,
			Keywords: []string{
				"management",
				"application",
				"site",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "additional_categories",
					Description: `(Optional) Used to configure or edit the additional categories of a custom application / site used in the Application and URL Filtering or Threat Prevention.additional_categories blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description for the application.`,
				},
				resource.Attribute{
					Name:        "primary_category",
					Description: `(Optional) Each application is assigned to one primary category based on its most defining aspect.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "url_list",
					Description: `(Optional) URLs that determine this particular application.url_list blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "application_signature",
					Description: `(Optional) Application signature generated by <a href="https://supportcenter.checkpoint.com/supportcenter/portal?eventSubmit_doGoviewsolutiondetails=&solutionid=sk103051">Signature Tool</a>.`,
				},
				resource.Attribute{
					Name:        "urls_defined_as_regular_expression",
					Description: `(Optional) States whether the URL is defined as a Regular Expression or not.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_application_site_category",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Application Site Category.`,
			Description: `

This resource allows you to execute Check Point Application Site Category.

`,
			Keywords: []string{
				"management",
				"application",
				"site",
				"category",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) N/A`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_application_site_group",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Application Site Group.`,
			Description: `

This resource allows you to execute Check Point Application Site Group.

`,
			Keywords: []string{
				"management",
				"application",
				"site",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `(Optional) Collection of application and URL filtering objects identified by the name or UID.members blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_assign_global_assignment",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Assign Global Assignment.`,
			Description: `

This command resource allows you to execute Check Point Assign Global Assignment.

`,
			Keywords: []string{
				"management",
				"assign",
				"global",
				"assignment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dependent_domains",
					Description: `(Optional) N/Adependent_domains blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "global_domains",
					Description: `(Optional) N/Aglobal_domains blocks are documented below. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_aws_data_center_server",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point aws data center server.`,
			Description: `

This resource allows you to execute Check Point AWS Data Center Server.

`,
			Keywords: []string{
				"management",
				"aws",
				"data",
				"center",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "authentication_method",
					Description: `(Required) user-authentication Uses the Access keys to authenticate. role-authentication Uses the AWS IAM role to authenticate. This option requires the Security Management Server be deployed in AWS and has an IAM Role.`,
				},
				resource.Attribute{
					Name:        "access_key_id",
					Description: `(Required for authentication-method: user-authentication) Access key ID for the AWS account. Required for authentication-method:user-authentication.`,
				},
				resource.Attribute{
					Name:        "secret_access_key",
					Description: `(Required for authentication-method: user-authentication) Secret access key for the AWS account. Required for authentication-method:user-authentication.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Select the AWS region.`,
				},
				resource.Attribute{
					Name:        "enable_sts_assume_role",
					Description: `(Optional) Enables the STS Assume Role option. After it is enabled, the sts-role field is mandatory, whereas the sts-external-id is optional.`,
				},
				resource.Attribute{
					Name:        "sts_role",
					Description: `(Required for enable-sts-assume-role: true) Enables the STS Assume Role option. After it is enabled, the sts-role field is mandatory, whereas the sts-external-id is optional.`,
				},
				resource.Attribute{
					Name:        "sts_external_id",
					Description: `(Optional) An optional STS External-Id to use when assuming the role.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers. tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_azure_ad",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Azure Ad.`,
			Description: `

This resource allows you to execute Check Point Azure Ad.

`,
			Keywords: []string{
				"management",
				"azure",
				"ad",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "authentication_method",
					Description: `(Required) <b>user-authentication</b><br>Uses the Azure AD User to authenticate.<br><b>service-principal-authentication</b><br>Uses the Service Principal to authenticate.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Password of the Azure account.<br><p><font color="red">Required for authentication-method:</font></p>user-authentication.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) An Azure Active Directory user Format<br>&lt;username&gt;@&lt;domain&gt;.<br><p><font color="red">Required for authentication-method:</font></p>user-authentication.`,
				},
				resource.Attribute{
					Name:        "application_id",
					Description: `(Required) The Application ID of the Service Principal, in UUID format.<br><p><font color="red">Required for authentication-method:</font></p>service-principal-authentication.`,
				},
				resource.Attribute{
					Name:        "application_key",
					Description: `(Required) The key created for the Service Principal.<br><p><font color="red">Required for authentication-method:</font></p>service-principal-authentication.`,
				},
				resource.Attribute{
					Name:        "directory_id",
					Description: `(Required) The Directory ID of the Azure AD, in UUID format.<br><p><font color="red">Required for authentication-method:</font></p>service-principal-authentication.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
				resource.Attribute{
					Name:        "properties",
					Description: `Azure AD connection properties. properties blocks are documented below. ` + "`" + `properties` + "`" + ` supports the following:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_azure_data_center_server",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point azure data center server.`,
			Description: `

This resource allows you to execute Check Point Azure Data Center Server.

`,
			Keywords: []string{
				"management",
				"azure",
				"data",
				"center",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "authentication_method",
					Description: `(Required) user-authentication Uses the Azure AD User to authenticate. service-principal-authentication Uses the Service Principal to authenticate.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required for authentication-method: user-authentication) An Azure Active Directory user Format <username>@<domain>. Required for authentication-method: user-authentication.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Password of the Azure account. Required for authentication-method: user-authentication.`,
				},
				resource.Attribute{
					Name:        "password_base64",
					Description: `(Optional) Password of the Azure account encoded in Base64. Required for authentication-method: user-authentication.`,
				},
				resource.Attribute{
					Name:        "application_id",
					Description: `(Required for authentication-method: service-principal-authentication) The Application ID of the Service Principal, in UUID format. Required for authentication-method: service-principal-authentication.`,
				},
				resource.Attribute{
					Name:        "application_key",
					Description: `(Required for authentication-method: service-principal-authentication) The key created for the Service Principal. Required for authentication-method: service-principal-authentication.`,
				},
				resource.Attribute{
					Name:        "directory_id",
					Description: `(Required for authentication-method: service-principal-authentication) The Directory ID of the Azure AD, in UUID format. Required for authentication-method: service-principal-authentication.`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `(Optional) Select the Azure Cloud Environment.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers. tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_backup_domain",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Backup Domain.`,
			Description: `

This command resource allows you to execute Check Point Backup Domain.

`,
			Keywords: []string{
				"management",
				"backup",
				"domain",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) Domain can be identified by name or UID.`,
				},
				resource.Attribute{
					Name:        "file_path",
					Description: `(Optional) Path in which the backup domain data will be saved. <br>Should be the directory path or the full file path with ".tgz" <br>If no path was inserted the default will be: "/var/log/&lt;domain name&gt;_&lt;date&gt;.tgz". ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_check_network_feed",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Check Network Feed.`,
			Description:      ``,
			Keywords: []string{
				"management",
				"check",
				"network",
				"feed",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "network_feed",
					Description: `(Required) network feed parameters.network_feed blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "targets",
					Description: `(Required) On what targets to execute this command. Targets may be identified by their name, or object unique identifier.targets blocks are documented below. ` + "`" + `network_feed` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name.`,
				},
				resource.Attribute{
					Name:        "feed_url",
					Description: `(Optional) URL of the feed. URL should be written as http or https.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `(Optional) Certificate SHA-1 fingerprint to access the feed.`,
				},
				resource.Attribute{
					Name:        "feed_format",
					Description: `(Optional) Feed file format.`,
				},
				resource.Attribute{
					Name:        "feed_type",
					Description: `(Optional) Feed type to be enforced.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) password for authenticating with the URL.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) username for authenticating with the URL.`,
				},
				resource.Attribute{
					Name:        "custom_header",
					Description: `(Optional) Headers to allow different authentication methods with the URL.custom_header blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "update_interval",
					Description: `(Optional) Interval in minutes for updating the feed on the Security Gateway.`,
				},
				resource.Attribute{
					Name:        "data_column",
					Description: `(Optional) Number of the column that contains the feed's data.`,
				},
				resource.Attribute{
					Name:        "fields_delimiter",
					Description: `(Optional) The delimiter that separates between the columns in the feed.`,
				},
				resource.Attribute{
					Name:        "ignore_lines_that_start_with",
					Description: `(Optional) A prefix that will determine which lines to ignore.`,
				},
				resource.Attribute{
					Name:        "json_query",
					Description: `(Optional) JQ query to be parsed.`,
				},
				resource.Attribute{
					Name:        "use_gateway_proxy",
					Description: `(Optional) Use the gateway's proxy for retrieving the feed.`,
				},
				resource.Attribute{
					Name:        "domains_to_process",
					Description: `(Optional) Indicates which domains to process the commands on. It cannot be used with the details-level full, must be run from the System Domain only and with ignore-warnings true. Valid values are: CURRENT_DOMAIN, ALL_DOMAINS_ON_THIS_SERVER.domains_to_process blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. ` + "`" + `custom_header` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "header_name",
					Description: `(Optional) The name of the HTTP header we wish to add.`,
				},
				resource.Attribute{
					Name:        "header_value",
					Description: `(Optional) The name of the HTTP value we wish to add. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_check_threat_ioc_feed",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Check Threat Ioc Feed.`,
			Description:      ``,
			Keywords: []string{
				"management",
				"check",
				"threat",
				"ioc",
				"feed",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ioc_feed",
					Description: `(Required) threat ioc feed parameters.ioc_feed blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "targets",
					Description: `(Required) On what targets to execute this command. Targets may be identified by their name, or object unique identifier.targets blocks are documented below. ` + "`" + `ioc_feed` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name.`,
				},
				resource.Attribute{
					Name:        "feed_url",
					Description: `(Optional) URL of the feed. URL should be written as http or https.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) The feed indicator's action.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `(Optional) Certificate SHA-1 fingerprint to access the feed.`,
				},
				resource.Attribute{
					Name:        "custom_comment",
					Description: `(Optional) Custom IOC feed - the column number of comment.`,
				},
				resource.Attribute{
					Name:        "custom_confidence",
					Description: `(Optional) Custom IOC feed - the column number of confidence.`,
				},
				resource.Attribute{
					Name:        "custom_header",
					Description: `(Optional) Custom HTTP headers.custom_header blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "custom_name",
					Description: `(Optional) Custom IOC feed - the column number of name.`,
				},
				resource.Attribute{
					Name:        "custom_severity",
					Description: `(Optional) Custom IOC feed - the column number of severity.`,
				},
				resource.Attribute{
					Name:        "custom_type",
					Description: `(Optional) Custom IOC feed - the column number of type in case a specific type is not chosen.`,
				},
				resource.Attribute{
					Name:        "custom_value",
					Description: `(Optional) Custom IOC feed - the column number of value in case a specific type is chosen.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Sets whether this indicator feed is enabled.`,
				},
				resource.Attribute{
					Name:        "feed_type",
					Description: `(Optional) Feed type to be enforced.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) password for authenticating with the URL.`,
				},
				resource.Attribute{
					Name:        "use_custom_feed_settings",
					Description: `(Optional) Set in order to configure a custom indicator feed.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) username for authenticating with the URL.`,
				},
				resource.Attribute{
					Name:        "fields_delimiter",
					Description: `(Optional) The delimiter that separates between the columns in the feed.`,
				},
				resource.Attribute{
					Name:        "ignore_lines_that_start_with",
					Description: `(Optional) A prefix that will determine which lines to ignore.`,
				},
				resource.Attribute{
					Name:        "use_gateway_proxy",
					Description: `(Optional) Use the gateway's proxy for retrieving the feed.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. ` + "`" + `custom_header` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "header_name",
					Description: `(Optional) The name of the HTTP header we wish to add.`,
				},
				resource.Attribute{
					Name:        "header_value",
					Description: `(Optional) The name of the HTTP value we wish to add. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_checkpoint_host",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Checkpoint Host.`,
			Description: `

This resource allows you to execute Check Point Checkpoint Host.

`,
			Keywords: []string{
				"management",
				"host",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `(Optional) Checkpoint host interfaces. interfaces blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `(Optional) IPv4 address.`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `(Optional) IPv6 address.`,
				},
				resource.Attribute{
					Name:        "nat_settings",
					Description: `(Optional) NAT settings. nat_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "one_time_password",
					Description: `(Optional) Secure internal connection one time password.`,
				},
				resource.Attribute{
					Name:        "hardware",
					Description: `(Optional) Hardware name.`,
				},
				resource.Attribute{
					Name:        "os",
					Description: `(Optional) Operating system name.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) Checkpoint host platform version.`,
				},
				resource.Attribute{
					Name:        "management_blades",
					Description: `(Optional) Management blades. management_blades blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "logs_settings",
					Description: `(Optional) Logs settings. logs_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "save_logs_locally",
					Description: `(Optional) Enable save logs locally.`,
				},
				resource.Attribute{
					Name:        "send_alerts_to_server",
					Description: `(Optional) Collection of Server(s) to send alerts to identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "send_logs_to_backup_server",
					Description: `(Optional) Collection of Backup server(s) to send logs to identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "send_logs_to_server",
					Description: `(Optional) Collection of Server(s) to send logs to identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "sic_name",
					Description: `(Computed) Name of the Secure Internal Connection Trust.`,
				},
				resource.Attribute{
					Name:        "sic_state",
					Description: `(Computed) State the Secure Internal Connection Trust.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. ` + "`" + `interfaces` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Interface name.`,
				},
				resource.Attribute{
					Name:        "subnet4",
					Description: `(Optional) IPv4 network address.`,
				},
				resource.Attribute{
					Name:        "subnet6",
					Description: `(Optional) IPv6 network address.`,
				},
				resource.Attribute{
					Name:        "mask_length4",
					Description: `(Optional) IPv4 network mask length.`,
				},
				resource.Attribute{
					Name:        "mask_length6",
					Description: `(Optional) IPv6 network mask length.`,
				},
				resource.Attribute{
					Name:        "subnet_mask",
					Description: `(Optional) IPv4 network mask.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. ` + "`" + `nat_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "auto_rule",
					Description: `(Optional) Whether to add automatic address translation rules.`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `(Optional) IPv4 address.`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `(Optional) IPv6 address.`,
				},
				resource.Attribute{
					Name:        "hide_behind",
					Description: `(Optional) Hide behind method. This parameter is not required in case "method" parameter is "static".`,
				},
				resource.Attribute{
					Name:        "install_on",
					Description: `(Optional) Which gateway should apply the NAT translation.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Optional) NAT translation method. ` + "`" + `management_blades` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "network_policy_management",
					Description: `(Optional) Enable Network Policy Management.`,
				},
				resource.Attribute{
					Name:        "logging_and_status",
					Description: `(Optional) Enable Logging & Status.`,
				},
				resource.Attribute{
					Name:        "smart_event_server",
					Description: `(Optional) Enable SmartEvent server. When activating SmartEvent server, blades 'logging-and-status' and 'smart-event-correlation' should be set to True. To complete SmartEvent configuration, perform Install Database or Install Policy on your Security Management servers and Log servers. </br>Activating SmartEvent Server is not recommended in Management High Availability environment. For more information refer to sk25164.`,
				},
				resource.Attribute{
					Name:        "smart_event_correlation",
					Description: `(Optional) Enable SmartEvent Correlation Unit.`,
				},
				resource.Attribute{
					Name:        "endpoint_policy",
					Description: `(Optional) Enable Endpoint Policy. To complete Endpoint Security Management configuration, perform Install Database on your Endpoint Management Server. Field is not supported on Multi Domain Server environment.`,
				},
				resource.Attribute{
					Name:        "compliance",
					Description: `(Optional) Compliance blade. Can be set when 'network-policy-management' was selected to be True.`,
				},
				resource.Attribute{
					Name:        "user_directory",
					Description: `(Optional) Enable User Directory. Can be set when 'network-policy-management' was selected to be True.`,
				},
				resource.Attribute{
					Name:        "secondary",
					Description: `(Computed) Secondary Management enabled.`,
				},
				resource.Attribute{
					Name:        "identity_logging",
					Description: `(Computed) Identity Logging enabled. ` + "`" + `logs_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "free_disk_space_metrics",
					Description: `(Optional) Free disk space metrics.`,
				},
				resource.Attribute{
					Name:        "accept_syslog_messages",
					Description: `(Optional) Enable accept syslog messages.`,
				},
				resource.Attribute{
					Name:        "alert_when_free_disk_space_below",
					Description: `(Optional) Enable alert when free disk space is below threshold.`,
				},
				resource.Attribute{
					Name:        "alert_when_free_disk_space_below_threshold",
					Description: `(Optional) Alert when free disk space below threshold.`,
				},
				resource.Attribute{
					Name:        "alert_when_free_disk_space_below_type",
					Description: `(Optional) Alert when free disk space below type.`,
				},
				resource.Attribute{
					Name:        "before_delete_keep_logs_from_the_last_days",
					Description: `(Optional) Enable before delete keep logs from the last days.`,
				},
				resource.Attribute{
					Name:        "before_delete_keep_logs_from_the_last_days_threshold",
					Description: `(Optional) Before delete keep logs from the last days threshold.`,
				},
				resource.Attribute{
					Name:        "before_delete_run_script",
					Description: `(Optional) Enable Before delete run script.`,
				},
				resource.Attribute{
					Name:        "before_delete_run_script_command",
					Description: `(Optional) Before delete run script command.`,
				},
				resource.Attribute{
					Name:        "delete_index_files_older_than_days",
					Description: `(Optional) Enable delete index files older than days.`,
				},
				resource.Attribute{
					Name:        "delete_index_files_older_than_days_threshold",
					Description: `(Optional) Delete index files older than days threshold.`,
				},
				resource.Attribute{
					Name:        "delete_when_free_disk_space_below",
					Description: `(Optional) Enable delete when free disk space below.`,
				},
				resource.Attribute{
					Name:        "delete_when_free_disk_space_below_threshold",
					Description: `(Optional) Delete when free disk space below threshold.`,
				},
				resource.Attribute{
					Name:        "detect_new_citrix_ica_application_names",
					Description: `(Optional) Enable detect new citrix ica application names.`,
				},
				resource.Attribute{
					Name:        "enable_log_indexing",
					Description: `(Optional) Enable log indexing.`,
				},
				resource.Attribute{
					Name:        "forward_logs_to_log_server",
					Description: `(Optional) Enable forward logs to log server.`,
				},
				resource.Attribute{
					Name:        "forward_logs_to_log_server_name",
					Description: `(Optional) Forward logs to log server name.`,
				},
				resource.Attribute{
					Name:        "forward_logs_to_log_server_schedule_name",
					Description: `(Optional) Forward logs to log server schedule name.`,
				},
				resource.Attribute{
					Name:        "rotate_log_by_file_size",
					Description: `(Optional) Enable rotate log by file size.`,
				},
				resource.Attribute{
					Name:        "rotate_log_file_size_threshold",
					Description: `(Optional) Log file size threshold.`,
				},
				resource.Attribute{
					Name:        "rotate_log_on_schedule",
					Description: `(Optional) Enable rotate log on schedule.`,
				},
				resource.Attribute{
					Name:        "rotate_log_schedule_name",
					Description: `(Optional) Rotate log schedule name.`,
				},
				resource.Attribute{
					Name:        "smart_event_intro_correletion_unit",
					Description: `(Optional) Enable SmartEvent intro correletion unit.`,
				},
				resource.Attribute{
					Name:        "stop_logging_when_free_disk_space_below",
					Description: `(Optional) Enable stop logging when free disk space below.`,
				},
				resource.Attribute{
					Name:        "stop_logging_when_free_disk_space_below_threshold",
					Description: `(Optional) Stop logging when free disk space below threshold.`,
				},
				resource.Attribute{
					Name:        "turn_on_qos_logging",
					Description: `(Optional) Enable turn on qos logging.`,
				},
				resource.Attribute{
					Name:        "update_account_log_every",
					Description: `(Optional) Update account log in every amount of seconds.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_command_abort_get_interfaces",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Abort Get Interfaces.`,
			Description: `

This resource allows you to execute Check Point Abort Get Interfaces.

`,
			Keywords: []string{
				"management",
				"command",
				"abort",
				"get",
				"interfaces",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "task_id",
					Description: `(Required) get-interfaces task UID.`,
				},
				resource.Attribute{
					Name:        "force_cleanup",
					Description: `(Optional) Forcefully abort the "get-interfaces" task.`,
				},
				resource.Attribute{
					Name:        "message",
					Description: `Operation status.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_command_get_interfaces",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Get Interfaces.`,
			Description: `

This resource allows you to execute Check Point Get Interfaces.

` + "`" + `` + "`" + `` + "`" + `hcl
resource "checkpoint_management_command_get_interfaces" "get_interfaces" {
  target_uid = "2220d9ad-a251-5555-9a0a-4772a6511111"
}
` + "`" + `` + "`" + `` + "`" + `

`,
			Keywords: []string{
				"management",
				"command",
				"get",
				"interfaces",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "target_uid",
					Description: `(Required) Target unique identifier.`,
				},
				resource.Attribute{
					Name:        "target_name",
					Description: `(Required) Target name.`,
				},
				resource.Attribute{
					Name:        "group_interfaces_by_subnet",
					Description: `(Optional) Specify whether to group the cluster interfaces by a subnet. Otherwise, group the cluster interfaces by their names.`,
				},
				resource.Attribute{
					Name:        "use_defined_by_routes",
					Description: `(Optional) Specify whether to configure the topology "Defined by Routes" where applicable. Otherwise, configure the topology to "This Network" as default for internal interfaces.`,
				},
				resource.Attribute{
					Name:        "with_topology",
					Description: `(Optioanl) Specify whether to fetch the interfaces with their topology. Otherwise, the Management Server fetches the interfaces without their topology.`,
				},
				resource.Attribute{
					Name:        "task_id",
					Description: `The UID of the "get-interfaces" task. Use the "show-task" command to check the progress of the "get-interfaces" task.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_connect_cloud_services",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Connect Cloud Services.`,
			Description:      ``,
			Keywords: []string{
				"management",
				"connect",
				"cloud",
				"services",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "auth_token",
					Description: `(Required) Copy the authentication token from the Smart-1 cloud service hosted in the Infinity Portal.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the connection to the Infinity Portal.`,
				},
				resource.Attribute{
					Name:        "connected_at",
					Description: `The time of the connection between the Management Server and the Infinity Portal. connected_at is documented below.`,
				},
				resource.Attribute{
					Name:        "management_url",
					Description: `The Management Server's public URL. ` + "`" + `connected_at` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "iso_8601",
					Description: `Date and time represented in international ISO 8601 format.`,
				},
				resource.Attribute{
					Name:        "posix",
					Description: `Number of milliseconds that have elapsed since 00:00:00, 1 January 1970. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_data_center_query",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Data Center Query.`,
			Description:      ``,
			Keywords: []string{
				"management",
				"data",
				"center",
				"query",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "data_centers",
					Description: `(Optional) Collection of Data Center servers identified by the name or UID. Use "All" to select all data centers.data_centers blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "query_rules",
					Description: `(Optional) Data Center Query Rules.<br>There is an 'AND' operation between multiple Query Rules.query_rules blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. ` + "`" + `query_rules` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "key_type",
					Description: `(Optional) The type of the "key" parameter.<br>Use "predefined" for these keys: type-in-data-center, name-in-data-center, and ip-address.<br>Use "tag" to query the Data Center tag�s property.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) Defines in which Data Center property to query.<br>For key-type "predefined", use these keys:type-in-data-center, name-in-data-center, and ip-address.<br>For key-type "tag", use the Data Center tag key to query.<br>Keys are case-insensitive.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Optional) The value(s) of the Data Center property to match the Query Rule.<br>Values are case-insensitive.<br>There is an 'OR' operation between multiple values.<br>For key-type "predefined" and key 'ip-address', the values must be an IPv4 or IPv6 address.<br>For key-type "tag", the values must be the Data Center tag values.values blocks are documented below.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_delete_api_key",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Delete Api Key.`,
			Description: `

This command resource allows you to execute Check Point Delete Api Key.

`,
			Keywords: []string{
				"management",
				"delete",
				"api",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "api_key",
					Description: `(Required) API key to be deleted.`,
				},
				resource.Attribute{
					Name:        "admin_uid",
					Description: `(Required) Administrator uid to generate API key for.`,
				},
				resource.Attribute{
					Name:        "admin_name",
					Description: `(Required) Administrator name to generate API key for. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_delete_data_center_object",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Delete Data Center Object.`,
			Description: `

This command resource allows you to execute Check Point Delete Data Center Object.

`,
			Keywords: []string{
				"management",
				"delete",
				"data",
				"center",
				"object",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_delete_repository_package",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Delete Repository Package.`,
			Description:      ``,
			Keywords: []string{
				"management",
				"delete",
				"repository",
				"package",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the software package. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_delete_threat_protections",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Delete Threat Protections.`,
			Description: `

This command resource allows you to execute Check Point Delete Threat Protections.

`,
			Keywords: []string{
				"management",
				"delete",
				"threat",
				"protections",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "package_format",
					Description: `(Optional) Protections package format. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_delete_updatable_object",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Delete Updatable Object.`,
			Description: `

This command resource allows you to execute Check Point Delete Updatable Object.

`,
			Keywords: []string{
				"management",
				"delete",
				"updatable",
				"object",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_discard",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Discard.`,
			Description: `

This command resource allows you to execute Check Point Discard.

`,
			Keywords: []string{
				"management",
				"discard",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_disconnect",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Disconnect.`,
			Description: `

This command resource allows you to execute Check Point Disconnect.

`,
			Keywords: []string{
				"management",
				"disconnect",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "discard",
					Description: `(Optional) Discard all changes committed during the session. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_disconnect_cloud_services",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Disconnect Cloud Services.`,
			Description:      ``,
			Keywords: []string{
				"management",
				"disconnect",
				"cloud",
				"services",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "force",
					Description: `(Optional) Disconnect the Management Server from Check Point Infinity Portal, and reset the connection locally, regardless of the result in the Infinity Portal. This flag can be used if the disconnect-cloud-services command failed. Since with this flag this command affects only the local configuration, make sure to disconnect the Management Server in the Infinity Portal as well. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_dns_domain",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Dns Domain.`,
			Description: `

This resource allows you to execute Check Point Dns Domain.

`,
			Keywords: []string{
				"management",
				"dns",
				"domain",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "is_sub_domain",
					Description: `(Required) Whether to match sub-domains in addition to the domain itself.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_domain",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Domain.`,
			Description: `

This resource allows you to execute Check Point Domain.

`,
			Keywords: []string{
				"management",
				"domain",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name. Must be unique in the domain.`,
				},
				resource.Attribute{
					Name:        "servers",
					Description: `(Required) Domain servers.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. ` + "`" + `servers` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name. Must be unique in the domain.`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `(Optional) IPv4 address.`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `(Optional) IPv6 address.`,
				},
				resource.Attribute{
					Name:        "multi_domain_server",
					Description: `(Required) Multi Domain server name or UID.`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `(Optional) Activate domain server. Only one domain server is allowed to be active.`,
				},
				resource.Attribute{
					Name:        "skip_start_domain_server",
					Description: `(Optional) Set this value to be true to prevent starting the new created domain.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Domain server type.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_domain_permissions_profile",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Domain Permissions Profile.`,
			Description:      ``,
			Keywords: []string{
				"management",
				"domain",
				"permissions",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "permission_type",
					Description: `(Optional) The type of the Permissions Profile.`,
				},
				resource.Attribute{
					Name:        "edit_common_objects",
					Description: `(Optional) Define and manage objects in the Check Point database: Network Objects, Services, Custom Application Site, VPN Community, Users, Servers, Resources, Time, UserCheck, and Limit.<br>Only a 'Customized' permission-type profile can edit this permission.`,
				},
				resource.Attribute{
					Name:        "access_control",
					Description: `(Optional) Access Control permissions.<br>Only a 'Customized' permission-type profile can edit these permissions.access_control blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Optional) Endpoint permissions. Not supported for Multi-Domain Servers.<br>Only a 'Customized' permission-type profile can edit these permissions.endpoint blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "events_and_reports",
					Description: `(Optional) Events and Reports permissions.<br>Only a 'Customized' permission-type profile can edit these permissions.events_and_reports blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "gateways",
					Description: `(Optional) Gateways permissions. <br>Only a 'Customized' permission-type profile can edit these permissions.gateways blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "management",
					Description: `(Optional) Management permissions.management blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "monitoring_and_logging",
					Description: `(Optional) Monitoring and Logging permissions.<br>'Customized' permission-type profile can edit all these permissions. "Read Write All" permission-type can edit only dlp-logs-including-confidential-fields and manage-dlp-messages permissions.monitoring_and_logging blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "threat_prevention",
					Description: `(Optional) Threat Prevention permissions.<br>Only a 'Customized' permission-type profile can edit these permissions.threat_prevention blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "others",
					Description: `(Optional) Additional permissions.<br>Only a 'Customized' permission-type profile can edit these permissions.others blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. ` + "`" + `access_control` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "show_policy",
					Description: `(Optional) Select to let administrators work with Access Control rules and NAT rules. If not selected, administrators cannot see these rules.`,
				},
				resource.Attribute{
					Name:        "policy_layers",
					Description: `(Optional) Layer editing permissions.<br>Available only if show-policy is set to true.policy_layers blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "dlp_policy",
					Description: `(Optional) Configure DLP rules and Policies.`,
				},
				resource.Attribute{
					Name:        "geo_control_policy",
					Description: `(Optional) Work with Access Control rules that control traffic to and from specified countries.`,
				},
				resource.Attribute{
					Name:        "nat_policy",
					Description: `(Optional) Work with NAT in Access Control rules.`,
				},
				resource.Attribute{
					Name:        "qos_policy",
					Description: `(Optional) Work with QoS Policies and rules.`,
				},
				resource.Attribute{
					Name:        "access_control_objects_and_settings",
					Description: `(Optional) Allow editing of the following objet types: VPN Community, Access Role, Custom application group,Custom application, Custom category, Limit, Application - Match Settings, Application Category - Match Settings,Override Categorization, Application and URL filtering blade - Advanced Settings, Content Awareness blade - Advanced Settings.`,
				},
				resource.Attribute{
					Name:        "app_control_and_url_filtering_update",
					Description: `(Optional) Install Application and URL Filtering updates.`,
				},
				resource.Attribute{
					Name:        "install_policy",
					Description: `(Optional) Install Access Control Policies. ` + "`" + `endpoint` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "manage_policies_and_software_deployment",
					Description: `(Optional) The administrator can work with policies, rules and actions.`,
				},
				resource.Attribute{
					Name:        "edit_endpoint_policies",
					Description: `(Optional) Available only if manage-policies-and-software-deployment is set to true.`,
				},
				resource.Attribute{
					Name:        "policies_installation",
					Description: `(Optional) The administrator can install policies on endpoint computers.`,
				},
				resource.Attribute{
					Name:        "edit_software_deployment",
					Description: `(Optional) The administrator can define deployment rules, create packages for export, and configure advanced package settings.<br>Available only if manage-policies-and-software-deployment is set to true.`,
				},
				resource.Attribute{
					Name:        "software_deployment_installation",
					Description: `(Optional) The administrator can deploy packages and install endpoint clients.`,
				},
				resource.Attribute{
					Name:        "allow_executing_push_operations",
					Description: `(Optional) The administrator can start operations that the Security Management Server pushes directly to client computers with no policy installation required.`,
				},
				resource.Attribute{
					Name:        "authorize_preboot_users",
					Description: `(Optional) The administrator can add and remove the users who are permitted to log on to Endpoint Security client computers with Full Disk Encryption.`,
				},
				resource.Attribute{
					Name:        "recovery_media",
					Description: `(Optional) The administrator can create recovery media on endpoint computers and devices.`,
				},
				resource.Attribute{
					Name:        "remote_help",
					Description: `(Optional) The administrator can use the Remote Help feature to reset user passwords and give access to locked out users.`,
				},
				resource.Attribute{
					Name:        "reset_computer_data",
					Description: `(Optional) The administrator can reset a computer, which deletes all information about the computer from the Security Management Server. ` + "`" + `events_and_reports` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "smart_event",
					Description: `(Optional) 'Custom' - Configure SmartEvent permissions.`,
				},
				resource.Attribute{
					Name:        "events",
					Description: `(Optional) Work with event queries on the Events tab. Create custom event queries.<br>Available only if smart-event is set to 'Custom'.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Optional) Configure SmartEvent Policy rules and install SmartEvent Policies.<br>Available only if smart-event is set to 'Custom'.`,
				},
				resource.Attribute{
					Name:        "reports",
					Description: `(Optional) Create and run SmartEvent reports.<br>Available only if smart-event is set to 'Custom'. ` + "`" + `gateways` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "smart_update",
					Description: `(Optional) Install, update and delete Check Point licenses. This includes permissions to use SmartUpdate to manage licenses.`,
				},
				resource.Attribute{
					Name:        "lsm_gw_db",
					Description: `(Optional) Access to objects defined in LSM gateway tables. These objects are managed in the SmartProvisioning GUI or LSMcli command-line.<br>Note: 'Write' permission on lsm-gw-db allows administrator to run a script on SmartLSM gateway in Expert mode.`,
				},
				resource.Attribute{
					Name:        "manage_provisioning_profiles",
					Description: `(Optional) Administrator can add, edit, delete, and assign provisioning profiles to gateways (both LSM and non-LSM).<br>Available for edit only if lsm-gw-db is set with 'Write' permission.<br>Note: 'Read' permission on lsm-gw-db enables 'Read' permission for manage-provisioning-profiles.`,
				},
				resource.Attribute{
					Name:        "vsx_provisioning",
					Description: `(Optional) Create and configure Virtual Systems and other VSX virtual objects.`,
				},
				resource.Attribute{
					Name:        "system_backup",
					Description: `(Optional) Backup Security Gateways.`,
				},
				resource.Attribute{
					Name:        "system_restore",
					Description: `(Optional) Restore Security Gateways from saved backups.`,
				},
				resource.Attribute{
					Name:        "open_shell",
					Description: `(Optional) Use the SmartConsole CLI to run commands.`,
				},
				resource.Attribute{
					Name:        "run_one_time_script",
					Description: `(Optional) Run user scripts from the command line.`,
				},
				resource.Attribute{
					Name:        "run_repository_script",
					Description: `(Optional) Run scripts from the repository.`,
				},
				resource.Attribute{
					Name:        "manage_repository_scripts",
					Description: `(Optional) Add, change and remove scripts in the repository. ` + "`" + `management` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "cme_operations",
					Description: `(Optional) Permission to read / edit the Cloud Management Extension (CME) configuration.<br>Not supported for Multi-Domain Servers.`,
				},
				resource.Attribute{
					Name:        "manage_admins",
					Description: `(Optional) Controls the ability to manage Administrators, Permission Profiles, Trusted clients,API settings and Policy settings.<br>Only a "Read Write All" permission-type profile can edit this permission.<br>Not supported for Multi-Domain Servers.`,
				},
				resource.Attribute{
					Name:        "management_api_login",
					Description: `(Optional) Permission to log in to the Security Management Server and run API commands using thesetools: mgmt_cli (Linux and Windows binaries), Gaia CLI (clish) and Web Services (REST). Useful if you want to prevent administrators from running automatic scripts on the Management.<br>Note: This permission is not required to run commands from within the API terminal in SmartConsole.<br>Not supported for Multi-Domain Servers.`,
				},
				resource.Attribute{
					Name:        "manage_sessions",
					Description: `(Optional) Lets you disconnect, discard, publish, or take over other administrator sessions.<br>Only a "Read Write All" permission-type profile can edit this permission.`,
				},
				resource.Attribute{
					Name:        "high_availability_operations",
					Description: `(Optional) Configure and work with Domain High Availability.<br>Only a 'Customized' permission-type profile can edit this permission.`,
				},
				resource.Attribute{
					Name:        "approve_or_reject_sessions",
					Description: `(Optional) Approve / reject other sessions.`,
				},
				resource.Attribute{
					Name:        "publish_sessions",
					Description: `(Optional) Allow session publishing without an approval.`,
				},
				resource.Attribute{
					Name:        "manage_integration_with_cloud_services",
					Description: `(Optional) Manage integration with Cloud Services. ` + "`" + `monitoring_and_logging` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "monitoring",
					Description: `(Optional) See monitoring views and reports.`,
				},
				resource.Attribute{
					Name:        "management_logs",
					Description: `(Optional) See Multi-Domain Server audit logs.`,
				},
				resource.Attribute{
					Name:        "track_logs",
					Description: `(Optional) Use the log tracking features in SmartConsole.`,
				},
				resource.Attribute{
					Name:        "app_and_url_filtering_logs",
					Description: `(Optional) Work with Application and URL Filtering logs.`,
				},
				resource.Attribute{
					Name:        "https_inspection_logs",
					Description: `(Optional) See logs generated by HTTPS Inspection.`,
				},
				resource.Attribute{
					Name:        "packet_capture_and_forensics",
					Description: `(Optional) See logs generated by the IPS and Forensics features.`,
				},
				resource.Attribute{
					Name:        "show_packet_capture_by_default",
					Description: `(Optional) Enable packet capture by default.`,
				},
				resource.Attribute{
					Name:        "identities",
					Description: `(Optional) Show user and computer identity information in logs.`,
				},
				resource.Attribute{
					Name:        "show_identities_by_default",
					Description: `(Optional) Show user and computer identity information in logs by default.`,
				},
				resource.Attribute{
					Name:        "dlp_logs_including_confidential_fields",
					Description: `(Optional) Show DLP logs including confidential fields.`,
				},
				resource.Attribute{
					Name:        "manage_dlp_messages",
					Description: `(Optional) View/Release/Discard DLP messages.<br>Available only if dlp-logs-including-confidential-fields is set to true. ` + "`" + `threat_prevention` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "policy_layers",
					Description: `(Optional) Configure Threat Prevention Policy rules.<br>Note: To have policy-layers permissions you must set policy-exceptionsand profiles permissions. To have 'Write' permissions for policy-layers, policy-exceptions must be set with 'Write' permission as well.`,
				},
				resource.Attribute{
					Name:        "edit_layers",
					Description: `(Optional) 'ALL' - Gives permission to edit all layers.<br>"By Selected Profile In A Layer Editor" - Administrators can only edit the layer if the Threat Prevention layer editor gives editing permission to their profiles.<br>Available only if policy-layers is set to 'Write'.`,
				},
				resource.Attribute{
					Name:        "edit_settings",
					Description: `(Optional) Work with general Threat Prevention settings.`,
				},
				resource.Attribute{
					Name:        "policy_exceptions",
					Description: `(Optional) Configure exceptions to Threat Prevention rules.<br>Note: To have policy-exceptions you must set the protections permission.`,
				},
				resource.Attribute{
					Name:        "profiles",
					Description: `(Optional) Configure Threat Prevention profiles.`,
				},
				resource.Attribute{
					Name:        "protections",
					Description: `(Optional) Work with malware protections.`,
				},
				resource.Attribute{
					Name:        "install_policy",
					Description: `(Optional) Install Policies.`,
				},
				resource.Attribute{
					Name:        "ips_update",
					Description: `(Optional) Update IPS protections.<br>Note: You do not have to log into the User Center to receive IPS updates. ` + "`" + `others` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "client_certificates",
					Description: `(Optional) Create and manage client certificates for Mobile Access.`,
				},
				resource.Attribute{
					Name:        "edit_cp_users_db",
					Description: `(Optional) Work with user accounts and groups.`,
				},
				resource.Attribute{
					Name:        "https_inspection",
					Description: `(Optional) Enable and configure HTTPS Inspection rules.`,
				},
				resource.Attribute{
					Name:        "ldap_users_db",
					Description: `(Optional) Work with the LDAP database and user accounts, groups and OUs.`,
				},
				resource.Attribute{
					Name:        "user_authority_access",
					Description: `(Optional) Work with Check Point User Authority authentication.`,
				},
				resource.Attribute{
					Name:        "user_device_mgmt_conf",
					Description: `(Optional) Gives access to the UDM (User & Device Management) web-based application that handles security challenges in a "bring your own device" (BYOD) workspace. ` + "`" + `policy_layers` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "edit_layers",
					Description: `(Optional) "By Software Blades" - Edit Access Control layers that contain the blades enabled in the Permissions Profile.<br>"By Selected Profile In A Layer Editor" - Administrators can only edit the layer if the Access Control layer editor gives editing permission to their profiles.`,
				},
				resource.Attribute{
					Name:        "app_control_and_url_filtering",
					Description: `(Optional) Use Application and URL Filtering in Access Control rules.<br>Available only if edit-layers is set to "By Software Blades".`,
				},
				resource.Attribute{
					Name:        "content_awareness",
					Description: `(Optional) Use specified data types in Access Control rules.<br>Available only if edit-layers is set to "By Software Blades".`,
				},
				resource.Attribute{
					Name:        "firewall",
					Description: `(Optional) Work with Access Control and other Software Blades that do not have their own Policies.<br>Available only if edit-layers is set to "By Software Blades".`,
				},
				resource.Attribute{
					Name:        "mobile_access",
					Description: `(Optional) Work with Mobile Access rules.<br>Available only if edit-layers is set to "By Software Blades".`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_dynamic_global_network_object",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Dynamic Global Network Object.`,
			Description: `

This resource allows you to execute Check Point Dynamic Global Network Object.

`,
			Keywords: []string{
				"management",
				"dynamic",
				"global",
				"network",
				"object",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `(Optional) Collection of group identifiers.groups blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_dynamic_object",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Dynamic Object.`,
			Description: `

This resource allows you to execute Check Point Dynamic Object.

`,
			Keywords: []string{
				"management",
				"dynamic",
				"object",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_exception_group",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Exception Group.`,
			Description: `

This resource allows you to execute Check Point Exception Group.

`,
			Keywords: []string{
				"management",
				"exception",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "applied_profile",
					Description: `(Optional) The threat profile to apply this group to in the case of apply-on threat-rules-with-specific-profile.`,
				},
				resource.Attribute{
					Name:        "applied_threat_rules",
					Description: `(Optional) The threat rules to apply this group on in the case of apply-on manually-select-threat-rules.applied_threat_rules blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "apply_on",
					Description: `(Optional) An exception group can be set to apply on all threat rules, all threat rules which have a specific profile, or those rules manually chosen by the user.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. ` + "`" + `applied_threat_rules` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "layer",
					Description: `(Optional) The layer of the threat rule to which the group is to be attached.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the threat rule to which the group is to be attached.`,
				},
				resource.Attribute{
					Name:        "rule_number",
					Description: `(Optional) The rule-number of the threat rule to which the group is to be attached.`,
				},
				resource.Attribute{
					Name:        "position",
					Description: `(Optional) Position in the rulebase.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_export",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Export.`,
			Description: `

This command resource allows you to execute Check Point Export.

`,
			Keywords: []string{
				"management",
				"export",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "exclude_classes",
					Description: `(Optional) N/Aexclude_classes blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "exclude_topics",
					Description: `(Optional) N/Aexclude_topics blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "export_files_by_class",
					Description: `(Optional) N/A`,
				},
				resource.Attribute{
					Name:        "include_classes",
					Description: `(Optional) N/Ainclude_classes blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "include_topics",
					Description: `(Optional) N/Ainclude_topics blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "query_limit",
					Description: `(Optional) N/A`,
				},
				resource.Attribute{
					Name:        "task_id",
					Description: `(Computed) Asynchronous task unique identifier. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_gaia_best_practice",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Gaia Best Practice.`,
			Description:      ``,
			Keywords: []string{
				"management",
				"gaia",
				"best",
				"practice",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "best_practice_id",
					Description: `(Optional) Best Practice ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Best Practice Name.`,
				},
				resource.Attribute{
					Name:        "action_item",
					Description: `(Optional) To comply with Best Practice, do this action item.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Best Practice.`,
				},
				resource.Attribute{
					Name:        "expected_output_text",
					Description: `(Optional) The expected output of the script as plain text.`,
				},
				resource.Attribute{
					Name:        "expected_output_base64",
					Description: `(Optional) The expected output of the script as Base64.`,
				},
				resource.Attribute{
					Name:        "practice_script_path",
					Description: `(Optional) The absolute path of the script on the Management Server to run on Gaia Security Gateways during the Compliance scans.`,
				},
				resource.Attribute{
					Name:        "practice_script_base64",
					Description: `(Optional) The entire content of the script encoded in Base64 to run on Gaia Security Gateways during the Compliance scans.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
				resource.Attribute{
					Name:        "best_practice_id",
					Description: `ID of the Best Practice.`,
				},
				resource.Attribute{
					Name:        "regulations",
					Description: `The applicable regulations of the Gaia Best Practice. Appear only when the value of the 'details-level' parameter is set to 'full'. regulations blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "relevant_objects",
					Description: `The applicable objects of the Gaia Best Practice. Appear only when the value of the 'details-level' parameter is set to 'full'. relevant_objects blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current status of the Best Practice.`,
				},
				resource.Attribute{
					Name:        "user_defined",
					Description: `Determines if the Gaia Best Practice is a user-defined best practice. ` + "`" + `regulations` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "regulation_name",
					Description: `The name of the regulation.`,
				},
				resource.Attribute{
					Name:        "requirement_description",
					Description: `The description of the requirement.`,
				},
				resource.Attribute{
					Name:        "requirement_id",
					Description: `The id of the requirement.`,
				},
				resource.Attribute{
					Name:        "requirement_status",
					Description: `The status of the requirement. ` + "`" + `relevant_objects` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Determines if the relevant object is enabled or not.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the relevant object.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the relevant object.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `The uid of the relevant object.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_gcp_data_center_server",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point gcp data center server.`,
			Description: `

This resource allows you to execute Check Point Google Cloud Platform Data Center Server.

`,
			Keywords: []string{
				"management",
				"gcp",
				"data",
				"center",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "authentication_method",
					Description: `(Required) key-authentication Uses the Service Account private key file to authenticate. vm-instance-authentication Uses the Service Account VM Instance to authenticate. This option requires the Security Management Server deployed in a GCP, and runs as a Service Account with the required permissions.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Required for authentication-method: key-authentication) A Service Account Key JSON file, encoded in base64. Required for authentication-method:key-authentication.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers. tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_generic_data_center_server",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point generic data center server.`,
			Description: `

This resource allows you to execute Check Point Generic Data Center Server.

`,
			Keywords: []string{
				"management",
				"generic",
				"data",
				"center",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) URL of the JSON feed (e.g. https://example.com/file.json).`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Required) Update interval of the feed in seconds.`,
				},
				resource.Attribute{
					Name:        "custom_header",
					Description: `(Optional) When set to false, The admin is not using Key and Value for a Custom Header in order to connect to the feed server. When set to true, The admin is using Key and Value for a Custom Header in order to connect to the feed server.`,
				},
				resource.Attribute{
					Name:        "custom_key",
					Description: `(Optional) Key for the Custom Header, relevant and required only when custom_header set to true.`,
				},
				resource.Attribute{
					Name:        "custom_value",
					Description: `(Optional) Value for the Custom Header, relevant and required only when custom_header set to true.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers. tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_get_attachment",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Get Attachment.`,
			Description: `

This command resource allows you to execute Check Point Get Attachment.

`,
			Keywords: []string{
				"management",
				"get",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "attachment_id",
					Description: `(Optional) Attachment identifier from a log record.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Log id from a log record.`,
				},
				resource.Attribute{
					Name:        "task_id",
					Description: `(Computed) Asynchronous task unique identifier. ## How To Use Make sure this command will be executed in the right execution order.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_get_platform",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Get Platform.`,
			Description:      ``,
			Keywords: []string{
				"management",
				"get",
				"platform",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Gateway, cluster or Check Point host unique identifier.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Gateway, cluster or Check Point host name. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_global_assignment",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Global Assignment.`,
			Description: `

This resource allows you to execute Check Point Global Assignment.

`,
			Keywords: []string{
				"management",
				"global",
				"assignment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dependent_domain",
					Description: `(Optional) N/A`,
				},
				resource.Attribute{
					Name:        "global_access_policy",
					Description: `(Optional) Global domain access policy that is assigned to a dependent domain.`,
				},
				resource.Attribute{
					Name:        "global_domain",
					Description: `(Optional) N/A`,
				},
				resource.Attribute{
					Name:        "global_threat_prevention_policy",
					Description: `(Optional) Global domain threat prevention policy that is assigned to a dependent domain.`,
				},
				resource.Attribute{
					Name:        "manage_protection_actions",
					Description: `(Optional) N/A`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
				resource.Attribute{
					Name:        "assignment_up_to_date",
					Description: `The time when the assignment was assigned. assignment_up_to_date blocks are documented below. ` + "`" + `assignment_up_to_date` + "`" + ` supports the follwoing:`,
				},
				resource.Attribute{
					Name:        "iso_8601",
					Description: `Date and time represented in international ISO 8601 format.`,
				},
				resource.Attribute{
					Name:        "posix",
					Description: `Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_group",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to add/update/delete Check Point Group.`,
			Description: `

This resource allows you to add/update/delete Check Point Group.

`,
			Keywords: []string{
				"management",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name. Should be unique in the domain.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `(Optional) Collection of Network objects identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_group_with_exclusion",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Group With Exclusion.`,
			Description: `

This resource allows you to execute Check Point Group With Exclusion.

`,
			Keywords: []string{
				"management",
				"group",
				"with",
				"exclusion",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "except",
					Description: `(Optional) Name or UID of an object which the group excludes.`,
				},
				resource.Attribute{
					Name:        "include",
					Description: `(Optional) Name or UID of an object which the group includes.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_gsn_handover_group",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Gsn Handover Group.`,
			Description: `

This resource allows you to execute Check Point Gsn Handover Group.

`,
			Keywords: []string{
				"management",
				"gsn",
				"handover",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "enforce_gtp",
					Description: `(Optional) Enable enforce GTP signal packet rate limit from this group.`,
				},
				resource.Attribute{
					Name:        "gtp_rate",
					Description: `(Optional) Limit of the GTP rate in PDU/sec.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `(Optional) Collection of GSN handover group members identified by the name or UID.members blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_ha_full_sync",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point HA Full Sync.`,
			Description: `

This command resource allows you to execute Check Point HA Full Sync.

`,
			Keywords: []string{
				"management",
				"ha",
				"full",
				"sync",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Peer name (Multi Domain Server, Domain Server or Security Management Server).`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Peer unique identifier (Multi Domain Server, Domain Server or Security Management Server).`,
				},
				resource.Attribute{
					Name:        "task_id",
					Description: `(Computed) Asynchronous task unique identifier. ## How To Use Make sure this command will be executed in the right execution order.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_host",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to add/update/delete Check Point Host.`,
			Description: `

This resource allows you to add/update/delete Check Point Host.

`,
			Keywords: []string{
				"management",
				"host",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name. Should be unique in the domain.`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `(Optional) IPv4 address.`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `(Optional) IPv6 address.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `(Optional) Host interfaces. Host interfaces blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "nat_settings",
					Description: `(Optional) NAT settings. NAT settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "host_servers",
					Description: `(Optional) Servers Configuration. Servers Configuration blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers. ` + "`" + `interfaces` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name. Should be unique in the domain.`,
				},
				resource.Attribute{
					Name:        "subnet4",
					Description: `(Optional) IPv4 network address.`,
				},
				resource.Attribute{
					Name:        "subnet6",
					Description: `(Optional) IPv6 network address.`,
				},
				resource.Attribute{
					Name:        "mask_length4",
					Description: `(Optional) IPv4 network mask length.`,
				},
				resource.Attribute{
					Name:        "mask_length6",
					Description: `(Optional) IPv6 network mask length.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string. ` + "`" + `nat_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "auto_rule",
					Description: `(Optional) Whether to add automatic address translation rules.`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `(Optional) IPv4 address.`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `(Optional) IPv6 address.`,
				},
				resource.Attribute{
					Name:        "hide_behind",
					Description: `(Optional) Hide behind method. This parameter is not required in case \"method\" parameter is \"static\".`,
				},
				resource.Attribute{
					Name:        "install_on",
					Description: `(Optional) Which gateway should apply the NAT translation.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Optional) NAT translation method. ` + "`" + `host_servers` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "dns_server",
					Description: `(Optional) Gets True if this server is a DNS Server.`,
				},
				resource.Attribute{
					Name:        "mail_server",
					Description: `(Optional) Gets True if this server is a Mail Server.`,
				},
				resource.Attribute{
					Name:        "web_server",
					Description: `(Optional) Gets True if this server is a Web Server.`,
				},
				resource.Attribute{
					Name:        "web_server_config",
					Description: `(Optional) Web Server configuration. Web Server configuration blocks are documented below. ` + "`" + `web_server_config` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "additional_ports",
					Description: `(Optional) Server additional ports.`,
				},
				resource.Attribute{
					Name:        "application_engines",
					Description: `(Optional) Application engines of this web server.`,
				},
				resource.Attribute{
					Name:        "listen_standard_port",
					Description: `(Optional) "Whether server listens to standard port.`,
				},
				resource.Attribute{
					Name:        "operating_system",
					Description: `(Optional) Operating System.`,
				},
				resource.Attribute{
					Name:        "protected_by",
					Description: `(Optional) Network object which protects this server identified by the name or UID.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_https_layer",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Https Layer.`,
			Description: `

This resource allows you to execute Check Point Https Layer.

`,
			Keywords: []string{
				"management",
				"https",
				"layer",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `(Optional) Define the Layer as Shared (TRUE/FALSE).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_https_rule",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Https Rule.`,
			Description: `

This resource allows you to execute Check Point Https Rule.

`,
			Keywords: []string{
				"management",
				"https",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "rule_number",
					Description: `(Required) Rule number.`,
				},
				resource.Attribute{
					Name:        "layer",
					Description: `(Required) Layer that holds the Object. Identified by the Name or UID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) HTTPS rule name.`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `(Optional) Collection of Network objects identified by Name or UID that represents connection destination.destination blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `(Optional) Collection of Network objects identified by Name or UID that represents connection service.service blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) Collection of Network objects identified by Name or UID that represents connection source.source blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Rule inspect level. "Bypass" or "Inspect".`,
				},
				resource.Attribute{
					Name:        "blade",
					Description: `(Optional) Blades for HTTPS Inspection. Identified by Name or UID to enable the inspection for. "Anti Bot","Anti Virus","Application Control","Data Awareness","DLP","IPS","Threat Emulation","Url Filtering".blade blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(Optional) Internal Server Certificate identified by Name or UID, otherwise, "Outbound Certificate" is a default value.`,
				},
				resource.Attribute{
					Name:        "destination_negate",
					Description: `(Optional) TRUE if "negate" value is set for Destination.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Enable/Disable the rule.`,
				},
				resource.Attribute{
					Name:        "install_on",
					Description: `(Optional) Which Gateways identified by the name or UID to install the policy on.install_on blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "service_negate",
					Description: `(Optional) TRUE if "negate" value is set for Service.`,
				},
				resource.Attribute{
					Name:        "site_category",
					Description: `(Optional) Collection of Site Categories objects identified by the name or UID.site_category blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "site_category_negate",
					Description: `(Optional) TRUE if "negate" value is set for Site Category.`,
				},
				resource.Attribute{
					Name:        "source_negate",
					Description: `(Optional) TRUE if "negate" value is set for Source.`,
				},
				resource.Attribute{
					Name:        "track",
					Description: `(Optional) "None","Log","Alert","Mail","SNMP trap","Mail","User Alert", "User Alert 2", "User Alert 3".`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
				resource.Attribute{
					Name:        "position",
					Description: `(Required) Position in the rulebase.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_https_section",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Https Section.`,
			Description: `

This resource allows you to execute Check Point Https Section.

`,
			Keywords: []string{
				"management",
				"https",
				"section",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "layer",
					Description: `(Required) Layer that holds the Object. Identified by the Name or UID.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
				resource.Attribute{
					Name:        "position",
					Description: `(Required) Position in the rulebase.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_identity_tag",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Identity Tag.`,
			Description: `

This resource allows you to execute Check Point Identity Tag.

`,
			Keywords: []string{
				"management",
				"identity",
				"tag",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "external_identifier",
					Description: `(Optional) External identifier. For example: Cisco ISE security group tag.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_idp_administrator_group",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Idp Administrator Group.`,
			Description:      ``,
			Keywords: []string{
				"management",
				"idp",
				"administrator",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required) Group ID or Name should be set base on the source attribute of 'groups' in the Saml Assertion.`,
				},
				resource.Attribute{
					Name:        "multi_domain_profile",
					Description: `(Optional) Administrator multi-domain profile.`,
				},
				resource.Attribute{
					Name:        "permissions_profile",
					Description: `(Optional) Administrator permissions profile. Permissions profile should not be provided when multi-domain-profile is set to "Multi-Domain Super User" or "Domain Super User".permissions_profile blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. ` + "`" + `permissions_profile` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional) N/A`,
				},
				resource.Attribute{
					Name:        "profile",
					Description: `(Optional) Permission profile.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_install_database",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Install Database.`,
			Description: `

This command resource allows you to execute Check Point Install Database.

`,
			Keywords: []string{
				"management",
				"install",
				"database",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "targets",
					Description: `(Required) Check Point host(s) with one or more Management Software Blades enabled. The targets can be identified by their name or unique identifier.targets blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "tasks",
					Description: `(Computed) Collection of asynchronous task unique identifiers. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_install_lsm_policy",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Install Lsm Policy.`,
			Description:      ``,
			Keywords: []string{
				"management",
				"install",
				"lsm",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "targets",
					Description: `(Required) On what targets to execute this command. Targets may be identified by their name, or object unique identifier.targets blocks are documented below. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_install_lsm_settings",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Install Lsm Settings.`,
			Description:      ``,
			Keywords: []string{
				"management",
				"install",
				"lsm",
				"settings",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "targets",
					Description: `(Required) On what targets to execute this command. Targets may be identified by their name, or object unique identifier.targets blocks are documented below. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_install_policy",
			Category:         "Management Resources",
			ShortDescription: `Install the published policy.`,
			Description: `

This command resource allows you to install the published policy.

`,
			Keywords: []string{
				"management",
				"install",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_package",
					Description: `(Required) The name of the Policy Package to be installed.`,
				},
				resource.Attribute{
					Name:        "targets",
					Description: `(Required) On what targets to execute this command. Targets may be identified by their name, or object unique identifier.`,
				},
				resource.Attribute{
					Name:        "access",
					Description: `(Optional) Set to be true in order to install the Access Control policy. By default, the value is true if Access Control policy is enabled on the input policy package, otherwise false.`,
				},
				resource.Attribute{
					Name:        "desktop_security",
					Description: `(Optional) Set to be true in order to install the Desktop Security policy. By default, the value is true if desktop security policy is enabled on the input policy package, otherwise false.`,
				},
				resource.Attribute{
					Name:        "qos",
					Description: `(Optional) Set to be true in order to install the QoS policy. By default, the value is true if Quality-of-Service policy is enabled on the input policy package, otherwise false.`,
				},
				resource.Attribute{
					Name:        "threat_prevention",
					Description: `(Optional) Set to be true in order to install the Threat Prevention policy. By default, the value is true if Threat Prevention policy is enabled on the input policy package, otherwise false.`,
				},
				resource.Attribute{
					Name:        "install_on_all_cluster_members_or_fail",
					Description: `(Optional) Relevant for the gateway clusters. If true, the policy is installed on all the cluster members. If the installation on a cluster member fails, don't install on that cluster.`,
				},
				resource.Attribute{
					Name:        "prepare_only",
					Description: `(Optional) If true, prepares the policy for the installation, but doesn't install it on an installation target.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `(Optional) The UID of the revision of the policy to install.`,
				},
				resource.Attribute{
					Name:        "triggers",
					Description: `(Optional) Triggers a install-policy if there are any changes to objects in this list.`,
				},
				resource.Attribute{
					Name:        "task_id",
					Description: `(Computed) Asynchronous task unique identifier. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_install_software_package",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Install Software Package.`,
			Description: `

This command resource allows you to execute Check Point Install Software Package.

`,
			Keywords: []string{
				"management",
				"install",
				"software",
				"package",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the software package.`,
				},
				resource.Attribute{
					Name:        "targets",
					Description: `(Required) On what targets to execute this command. Targets may be identified by their name, or object unique identifier.targets blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "cluster_installation_settings",
					Description: `(Optional) Installation settings for cluster.cluster_installation_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "concurrency_limit",
					Description: `(Optional) The number of targets, on which the same package is installed at the same time.`,
				},
				resource.Attribute{
					Name:        "task_id",
					Description: `(Computed) Asynchronous task unique identifier. ` + "`" + `cluster_installation_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "cluster_delay",
					Description: `(Optional) The delay between end of installation on one cluster members and start of installation on the next cluster member.`,
				},
				resource.Attribute{
					Name:        "cluster_strategy",
					Description: `(Optional) The cluster installation strategy. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_interoperable_device",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Interoperable Device.`,
			Description:      ``,
			Keywords: []string{
				"management",
				"interoperable",
				"device",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `(Optional) IPv4 address of the Interoperable Device.`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `(Optional) IPv6 address of the Interoperable Device.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `(Optional) Network interfaces.interfaces blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "vpn_settings",
					Description: `(Optional) VPN domain properties for the Interoperable Device.vpn_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "domains_to_process",
					Description: `(Optional) Indicates which domains to process the commands on. It cannot be used with the details-level full, must be run from the System Domain only and with ignore-warnings true. Valid values are: CURRENT_DOMAIN, ALL_DOMAINS_ON_THIS_SERVER.domains_to_process blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `(Optional) Collection of group identifiers.groups blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. ` + "`" + `interfaces` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name. Must be unique in the domain.`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `(Optional) IPv4 address.`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `(Optional) IPv6 address.`,
				},
				resource.Attribute{
					Name:        "network_mask",
					Description: `(Optional) IPv4 or IPv6 network mask. If both masks are required use ipv4-network-mask and ipv6-network-mask fields explicitly. Instead of providing mask itself it is possible to specify IPv4 or IPv6 mask length in mask-length field. If both masks length are required use ipv4-mask-length and ipv6-mask-length fields explicitly.`,
				},
				resource.Attribute{
					Name:        "ipv4_network_mask",
					Description: `(Optional) IPv4 network address.`,
				},
				resource.Attribute{
					Name:        "ipv6_network_mask",
					Description: `(Optional) IPv6 network address.`,
				},
				resource.Attribute{
					Name:        "ipv4_mask_length",
					Description: `(Optional) IPv4 network mask length.`,
				},
				resource.Attribute{
					Name:        "ipv6_mask_length",
					Description: `(Optional) IPv6 network mask length.`,
				},
				resource.Attribute{
					Name:        "anti_spoofing",
					Description: `(Optional) Is anti spoofing enabled on the interface.`,
				},
				resource.Attribute{
					Name:        "anti_spoofing_settings",
					Description: `(Optional) Anti spoofing settings.anti_spoofing_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "topology",
					Description: `(Optional) Topology configuration.`,
				},
				resource.Attribute{
					Name:        "topology_settings",
					Description: `(Optional) Internal topology settings.topology_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "domains_to_process",
					Description: `(Optional) Indicates which domains to process the commands on. It cannot be used with the details-level full, must be run from the System Domain only and with ignore-warnings true. Valid values are: CURRENT_DOMAIN, ALL_DOMAINS_ON_THIS_SERVER.domains_to_process blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. ` + "`" + `vpn_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "vpn_domain",
					Description: `(Optional) Network group representing the customized encryption domain. Must be set when vpn-domain-type is set to 'manual' option.`,
				},
				resource.Attribute{
					Name:        "vpn_domain_exclude_external_ip_addresses",
					Description: `(Optional) Exclude the external IP addresses from the VPN domain of this Interoperable device.`,
				},
				resource.Attribute{
					Name:        "vpn_domain_type",
					Description: `(Optional) Indicates the encryption domain. ` + "`" + `anti_spoofing_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) If packets will be rejected (the Prevent option) or whether the packets will be monitored (the Detect option).`,
				},
				resource.Attribute{
					Name:        "exclude_packets",
					Description: `(Optional) Don't check packets from excluded network.`,
				},
				resource.Attribute{
					Name:        "excluded_network_name",
					Description: `(Optional) Excluded network name.`,
				},
				resource.Attribute{
					Name:        "excluded_network_uid",
					Description: `(Optional) Excluded network UID.`,
				},
				resource.Attribute{
					Name:        "spoof_tracking",
					Description: `(Optional) Spoof tracking. ` + "`" + `topology_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "interface_leads_to_dmz",
					Description: `(Optional) Whether this interface leads to demilitarized zone (perimeter network).`,
				},
				resource.Attribute{
					Name:        "specific_network",
					Description: `(Optional) Network behind this interface.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_ise_data_center_server",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Cisco ISE data center server.`,
			Description: `

This resource allows you to execute Check Point Cisco ISE Data Center Server.

`,
			Keywords: []string{
				"management",
				"ise",
				"data",
				"center",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "hostnames",
					Description: `(Required) IP address or hostname of the Cisco ISE administration Node(s).`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) Username of the ISE administrator.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Password of the ISE administrator.`,
				},
				resource.Attribute{
					Name:        "password_base64",
					Description: `(Optional) Password of the ISE administrator encoded in Base64.`,
				},
				resource.Attribute{
					Name:        "certificate_fingerprint",
					Description: `(Optional) Specify the SHA-1 or SHA-256 fingerprint of the Data Center Server's certificate.`,
				},
				resource.Attribute{
					Name:        "unsafe_auto_accept",
					Description: `(Optional) When set to false, the current Data Center Server's certificate should be trusted, either by providing the certificate-fingerprint argument or by relying on a previously trusted certificate of this hostname. When set to true, trust the current Data Center Server's certificate as-is.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers. tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_keepalive",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Keepalive.`,
			Description: `

This command resource allows you to execute Check Point Keepalive.

`,
			Keywords: []string{
				"management",
				"keepalive",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_kubernetes_data_center_server",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point kubernetes data center server.`,
			Description: `

This resource allows you to execute Check Point Kubernetes Data Center Server.

`,
			Keywords: []string{
				"management",
				"kubernetes",
				"data",
				"center",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required) IP address or hostname of the Kubernetes server.`,
				},
				resource.Attribute{
					Name:        "token_file",
					Description: `(Required) Kubernetes access token encoded in base64.`,
				},
				resource.Attribute{
					Name:        "ca_certificate",
					Description: `(Optional) The Kubernetes public certificate key encoded in base64.`,
				},
				resource.Attribute{
					Name:        "unsafe_auto_accept",
					Description: `(Optional) When set to false, the current Data Center Server's certificate should be trusted, either by providing the certificate-fingerprint argument or by relying on a previously trusted certificate of this hostname. When set to true, trust the current Data Center Server's certificate as-is.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers. tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_login",
			Category:         "Management Resources",
			ShortDescription: `Log in to the server with username and password.`,
			Description: `

This command resource allows you to log in to the server with username and password.

`,
			Keywords: []string{
				"management",
				"login",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "user",
					Description: `(Required) Session unique identifier. Specify it to publish a different session than the one you currently use.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Administrator password.`,
				},
				resource.Attribute{
					Name:        "continue_last_session",
					Description: `(Optional) When 'continue-last-session' is set to 'True', the new session would continue where the last session was stopped. This option is available when the administrator has only one session that can be continued. If there is more than one session, see 'switch-session' API.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional) Use domain to login to specific domain. Domain can be identified by name or UID.`,
				},
				resource.Attribute{
					Name:        "enter_last_published_session",
					Description: `(Optional) Login to the last published session. Such login is done with the Read Only permissions.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Login with Read Only permissions. This parameter is not considered in case continue-last-session is true.`,
				},
				resource.Attribute{
					Name:        "session_comments",
					Description: `(Optional) Session comments.`,
				},
				resource.Attribute{
					Name:        "session_description",
					Description: `(Optional) Session description.`,
				},
				resource.Attribute{
					Name:        "session_name",
					Description: `(Optional) Session unique name.`,
				},
				resource.Attribute{
					Name:        "session_timeout",
					Description: `(Optional) Session expiration timeout in seconds. Default 600 seconds. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_logout",
			Category:         "Management Resources",
			ShortDescription: `Log out from the current session. After logging out the session id is not valid any more.`,
			Description: `

This command resource allows you to log out from the current session. After logging out the session id is not valid any more.

`,
			Keywords: []string{
				"management",
				"logout",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "triggers",
					Description: `(Optional) Triggers a logout if there are any changes to objects in this list. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_lsm_run_script",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Lsm Run Script.`,
			Description:      ``,
			Keywords: []string{
				"management",
				"lsm",
				"run",
				"script",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "script_base64",
					Description: `(Optional) The entire content of the script encoded in Base64.`,
				},
				resource.Attribute{
					Name:        "script",
					Description: `(Optional) The entire content of the script.`,
				},
				resource.Attribute{
					Name:        "targets",
					Description: `(Required) On what targets to execute this command. Targets may be identified by their name, or object unique identifier.targets blocks are documented below. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_lsv_profile",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to add/update/delete Check Point Lsv Profile.`,
			Description: `

This resource allows you to add/update/delete Check Point Lsv Profile.

`,
			Keywords: []string{
				"management",
				"lsv",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name. Should be unique in the domain.`,
				},
				resource.Attribute{
					Name:        "certificate_authority",
					Description: `(Required) Trusted Certificate authority for establishing trust between VPN peers, identified by name or UID.`,
				},
				resource.Attribute{
					Name:        "allowed_ip_addresses",
					Description: `(Optional) Collection of network objects identified by name or UID that represent IP addresses allowed in profile's VPN domain.`,
				},
				resource.Attribute{
					Name:        "restrict_allowed_addresses",
					Description: `(Optional) Indicate whether the IP addresses allowed in the VPN Domain will be restricted or not, according to allowed-ip-addresses field.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "vpn_domain",
					Description: `(Optional) peers' VPN Domain properties. vpn_domain blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. ` + "`" + `vpn_domain` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "limit_peer_domain_size",
					Description: `(Optional) Use this parameter to limit the number of IP addresses in the VPN Domain of each peer according to the value in the max-allowed-addresses field.`,
				},
				resource.Attribute{
					Name:        "max_allowed_addresses",
					Description: `(Optional) Maximum number of IP addresses in the VPN Domain of each peer. This value will be enforced only when limit-peer-domain-size field is set to true. Select a value between 1 and 256. Default value is 256.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_md_permissions_profile",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Md Permissions Profile.`,
			Description:      ``,
			Keywords: []string{
				"management",
				"md",
				"permissions",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "mds_provisioning",
					Description: `(Optional) Create and manage Multi-Domain Servers and Multi-Domain Log Servers.<br>Only a "Super User" permission-level profile can select this option.`,
				},
				resource.Attribute{
					Name:        "manage_admins",
					Description: `(Optional) Create and manage Multi-Domain Security Management administrators with the same or lower permission level. For example, a Domain manager cannot create Superusers or global managers.<br>Only a 'Manager' permission-level profile can edit this permission.`,
				},
				resource.Attribute{
					Name:        "manage_sessions",
					Description: `(Optional) Connect/disconnect Domain sessions, publish changes, and delete other administrator sessions.<br>Only a 'Manager' permission-level profile can edit this permission.`,
				},
				resource.Attribute{
					Name:        "management_api_login",
					Description: `(Optional) Permission to log in to the Security Management Server and run API commands using these tools: mgmt_cli (Linux and Windows binaries), Gaia CLI (clish) and Web Services (REST). Useful if you want to prevent administrators from running automatic scripts on the Management.<br>Note: This permission is not required to run commands from within the API terminal in SmartConsole.`,
				},
				resource.Attribute{
					Name:        "cme_operations",
					Description: `(Optional) Permission to read / edit the Cloud Management Extension (CME) configuration.`,
				},
				resource.Attribute{
					Name:        "global_vpn_management",
					Description: `(Optional) Lets the administrator select Enable global use for a Security Gateway shown in the MDS Gateways & Servers view.<br>Only a 'Manager' permission-level profile can edit this permission.`,
				},
				resource.Attribute{
					Name:        "manage_global_assignments",
					Description: `(Optional) Controls the ability to create, edit and delete global assignment and not the ability to reassign, which is set according to the specific Domain's permission profile.`,
				},
				resource.Attribute{
					Name:        "enable_default_profile_for_global_domains",
					Description: `(Optional) Enable the option to specify a default profile for all global domains.`,
				},
				resource.Attribute{
					Name:        "default_profile_global_domains",
					Description: `(Optional) Name or UID of the required default profile for all global domains.`,
				},
				resource.Attribute{
					Name:        "view_global_objects_in_domain",
					Description: `(Optional) Lets an administrator with no global objects permissions view the global objects in the domain. This option is required for valid domain management.`,
				},
				resource.Attribute{
					Name:        "enable_default_profile_for_local_domains",
					Description: `(Optional) Enable the option to specify a default profile for all local domains.`,
				},
				resource.Attribute{
					Name:        "default_profile_local_domains",
					Description: `(Optional) Name or UID of the required default profile for all local domains.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "domains_to_process",
					Description: `(Optional) Indicates which domains to process the commands on. It cannot be used with the details-level full, must be run from the System Domain only and with ignore-warnings true. Valid values are: CURRENT_DOMAIN, ALL_DOMAINS_ON_THIS_SERVER.domains_to_process blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
				resource.Attribute{
					Name:        "permission_level",
					Description: `(Optional) The level of the Multi Domain Permissions Profile.<br>The level cannot be changed after creation.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_mds",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point MDS.`,
			Description: `

This resource allows you to execute Check Point MDS.

`,
			Keywords: []string{
				"management",
				"mds",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `(Optional) IPv4 address.`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `(Optional) IPv6 address.`,
				},
				resource.Attribute{
					Name:        "hardware",
					Description: `(Optional) Hardware name. For example: Open server, Smart-1, Other.`,
				},
				resource.Attribute{
					Name:        "os",
					Description: `(Optional) Operating system name. For example: Gaia, Linux, SecurePlatform.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) System version.`,
				},
				resource.Attribute{
					Name:        "one_time_password",
					Description: `(Optional) Secure internal connection one time password.`,
				},
				resource.Attribute{
					Name:        "ip_pool_first",
					Description: `(Optional) First IP address in the range.`,
				},
				resource.Attribute{
					Name:        "ip_pool_last",
					Description: `(Optional) Last IP address in the range.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
				resource.Attribute{
					Name:        "server_type",
					Description: `(Optional) Type of the management server.`,
				},
				resource.Attribute{
					Name:        "sic_name",
					Description: `(Computed) Name of the Secure Internal Connection Trust.`,
				},
				resource.Attribute{
					Name:        "sic_state",
					Description: `(Computed) State the Secure Internal Connection Trust.`,
				},
				resource.Attribute{
					Name:        "domains",
					Description: `(Computed) Collection of Domain objects identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "global_domains",
					Description: `(Computed) Collection of Global domain objects identified by the name or UID.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_migrate_export_domain",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Migrate Export Domain.`,
			Description: `

This command resource allows you to execute Check Point Migrate Export Domain.

`,
			Keywords: []string{
				"management",
				"migrate",
				"export",
				"domain",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) Domain can be identified by name or UID.<br><font color="red">Required only for</font> exporting domain from Multi-Domain Server.`,
				},
				resource.Attribute{
					Name:        "file_path",
					Description: `(Optional) Path in which the exported domain data will be saved. <br>Should be the directory path or the full file path with ".tgz" <br>If no path was inserted the default will be: "/var/log/&lt;domain name&gt;_&lt;date&gt;.tgz".`,
				},
				resource.Attribute{
					Name:        "include_logs",
					Description: `(Optional) Export logs.`,
				},
				resource.Attribute{
					Name:        "task_id",
					Description: `(Computed) Asynchronous task unique identifier. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_migrate_import_domain",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Migrate Import Domain.`,
			Description: `

This command resource allows you to execute Check Point Migrate Import Domain.

`,
			Keywords: []string{
				"management",
				"migrate",
				"import",
				"domain",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "file_path",
					Description: `(Required) Path to the exported file to be imported. <br>Should be the full file path (example, "/var/log/domain1_exported.tgz").`,
				},
				resource.Attribute{
					Name:        "domain_ip_address",
					Description: `(Required) IPv4 address.<br><font color="red">Required only for</font> importing Security Management Server into Multi-Domain Server.`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `(Required) Domain name. Should be unique in the MDS.<br><font color="red">Required only for</font> importing Security Management Server into Multi-Domain Server.`,
				},
				resource.Attribute{
					Name:        "domain_server_name",
					Description: `(Required) Multi Domain server name.<br><font color="red">Required only for</font> importing Security Management Server into Multi-Domain Server.`,
				},
				resource.Attribute{
					Name:        "include_logs",
					Description: `(Optional) Import logs from the input package.`,
				},
				resource.Attribute{
					Name:        "task_id",
					Description: `(Computed) Asynchronous task unique identifier. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_multicast_address_range",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Multicast Address Range.`,
			Description: `

This resource allows you to execute Check Point Multicast Address Range.

`,
			Keywords: []string{
				"management",
				"multicast",
				"address",
				"range",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `(Optional) IPv4 address.`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `(Optional) IPv6 address.`,
				},
				resource.Attribute{
					Name:        "ipv4_address_first",
					Description: `(Optional) First IPv4 address in the range.`,
				},
				resource.Attribute{
					Name:        "ipv6_address_first",
					Description: `(Optional) First IPv6 address in the range.`,
				},
				resource.Attribute{
					Name:        "ipv4_address_last",
					Description: `(Optional) Last IPv4 address in the range.`,
				},
				resource.Attribute{
					Name:        "ipv6_address_last",
					Description: `(Optional) Last IPv6 address in the range.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_nat_rule",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to add/update/delete Check Point NAT Rule.`,
			Description: `

This resource allows you to add/update/delete Check Point NAT Rule.

`,
			Keywords: []string{
				"management",
				"nat",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "package",
					Description: `(Required) Name of the package.`,
				},
				resource.Attribute{
					Name:        "position",
					Description: `(Required) Position in the rulebase. Position blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Rule name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Enable/Disable the rule.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Optional) Nat method.`,
				},
				resource.Attribute{
					Name:        "install_on",
					Description: `(Optional) Which Gateways identified by the name or UID to install the policy on.`,
				},
				resource.Attribute{
					Name:        "original_destination",
					Description: `(Optional) Original destination.`,
				},
				resource.Attribute{
					Name:        "original_service",
					Description: `(Optional) Original service.`,
				},
				resource.Attribute{
					Name:        "original_source",
					Description: `(Optional) Original source.`,
				},
				resource.Attribute{
					Name:        "translated_destination",
					Description: `(Optional) Translated destination.`,
				},
				resource.Attribute{
					Name:        "translated_service",
					Description: `(Optional) Translated service.`,
				},
				resource.Attribute{
					Name:        "translated_source",
					Description: `(Optional) Translated source.`,
				},
				resource.Attribute{
					Name:        "auto_generated",
					Description: `(Computed) Auto generated.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string. ` + "`" + `position` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "top",
					Description: `(Optional) Add rule at the top of the rulebase.`,
				},
				resource.Attribute{
					Name:        "above",
					Description: `(Optional) Add rule above specific section/rule identified by uid or name.`,
				},
				resource.Attribute{
					Name:        "below",
					Description: `(Optional) Add rule below specific section/rule identified by uid or name.`,
				},
				resource.Attribute{
					Name:        "bottom",
					Description: `(Optional) Add rule at the bottom of the rulebase. ## Import ` + "`" + `checkpoint_management_nat_rule` + "`" + ` can be imported by using the following format: PACKAGE_NAME;RULE_UID ` + "`" + `` + "`" + `` + "`" + ` $ terraform import checkpoint_management_nat_rule.example "Standard;9423d36f-2d66-4754-b9e2-e9f4493751d3" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_nat_section",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to add/update/delete Check Point NAT section.`,
			Description: `

This resource allows you to add/update/delete Check Point NAT section.

`,
			Keywords: []string{
				"management",
				"nat",
				"section",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "package",
					Description: `(Required) Name of the package.`,
				},
				resource.Attribute{
					Name:        "position",
					Description: `(Required) Position in the rulebase. Position blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. ` + "`" + `position` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "top",
					Description: `(Optional) Add rule at the top of the rulebase.`,
				},
				resource.Attribute{
					Name:        "above",
					Description: `(Optional) Add rule above specific section/rule identified by uid or name.`,
				},
				resource.Attribute{
					Name:        "below",
					Description: `(Optional) Add rule below specific section/rule identified by uid or name.`,
				},
				resource.Attribute{
					Name:        "bottom",
					Description: `(Optional) Add rule at the bottom of the rulebase.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_network",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to add/update/delete Check Point Network Object.`,
			Description: `

This resource allows you to add/update/delete Check Point Network Object.

`,
			Keywords: []string{
				"management",
				"network",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name. Should be unique in the domain.`,
				},
				resource.Attribute{
					Name:        "subnet4",
					Description: `(Optional) IPv4 network address.`,
				},
				resource.Attribute{
					Name:        "subnet6",
					Description: `(Optional) IPv6 network address..`,
				},
				resource.Attribute{
					Name:        "mask_length4",
					Description: `(Optional) IPv4 network mask length.`,
				},
				resource.Attribute{
					Name:        "mask_length6",
					Description: `(Optional) IPv6 network mask length.`,
				},
				resource.Attribute{
					Name:        "nat_settings",
					Description: `(Optional) NAT settings. NAT settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "broadcast",
					Description: `(Optional) "Allow broadcast address inclusion.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string. ` + "`" + `nat_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "auto_rule",
					Description: `(Optional) Whether to add automatic address translation rules.`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `(Optional) IPv4 address.`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `(Optional) IPv6 address.`,
				},
				resource.Attribute{
					Name:        "hide_behind",
					Description: `(Optional) Hide behind method. This parameter is not required in case \"method\" parameter is \"static\".`,
				},
				resource.Attribute{
					Name:        "install_on",
					Description: `(Optional) Which gateway should apply the NAT translation.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Optional) NAT translation method.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_network_feed",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Network Feed.`,
			Description:      ``,
			Keywords: []string{
				"management",
				"network",
				"feed",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "feed_url",
					Description: `(Optional) URL of the feed. URL should be written as http or https.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `(Optional) Certificate SHA-1 fingerprint to access the feed.`,
				},
				resource.Attribute{
					Name:        "feed_format",
					Description: `(Optional) Feed file format.`,
				},
				resource.Attribute{
					Name:        "feed_type",
					Description: `(Optional) Feed type to be enforced.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) password for authenticating with the URL.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) username for authenticating with the URL.`,
				},
				resource.Attribute{
					Name:        "custom_header",
					Description: `(Optional) Headers to allow different authentication methods with the URL.custom_header blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "update_interval",
					Description: `(Optional) Interval in minutes for updating the feed on the Security Gateway.`,
				},
				resource.Attribute{
					Name:        "data_column",
					Description: `(Optional) Number of the column that contains the feed's data.`,
				},
				resource.Attribute{
					Name:        "fields_delimiter",
					Description: `(Optional) The delimiter that separates between the columns in the feed. For feed format 'Flat List' default is '\n'(new line).`,
				},
				resource.Attribute{
					Name:        "ignore_lines_that_start_with",
					Description: `(Optional) A prefix that will determine which lines to ignore.`,
				},
				resource.Attribute{
					Name:        "json_query",
					Description: `(Optional) JQ query to be parsed.`,
				},
				resource.Attribute{
					Name:        "use_gateway_proxy",
					Description: `(Optional) Use the gateway's proxy for retrieving the feed.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "domains_to_process",
					Description: `(Optional) Indicates which domains to process the commands on. It cannot be used with the details-level full, must be run from the System Domain only and with ignore-warnings true. Valid values are: CURRENT_DOMAIN, ALL_DOMAINS_ON_THIS_SERVER.domains_to_process blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. ` + "`" + `custom_header` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "header_name",
					Description: `(Optional) The name of the HTTP header we wish to add.`,
				},
				resource.Attribute{
					Name:        "header_value",
					Description: `(Optional) The name of the HTTP value we wish to add.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_nuage_data_center_server",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point nuage data center server.`,
			Description: `

This resource allows you to execute Check Point Nuage Data Center Server.

`,
			Keywords: []string{
				"management",
				"nuage",
				"data",
				"center",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required) IP address or hostname of the Nuage server.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) Username of the Nuage administrator.`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `(Required) Organization name or enterprise.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Password of the Nuage administrator.`,
				},
				resource.Attribute{
					Name:        "password_base64",
					Description: `(Optional) Password of the Nuage administrator encoded in Base64.`,
				},
				resource.Attribute{
					Name:        "certificate_fingerprint",
					Description: `(Optional) Specify the SHA-1 or SHA-256 fingerprint of the Data Center Server's certificate.`,
				},
				resource.Attribute{
					Name:        "unsafe_auto_accept",
					Description: `(Optional) When set to false, the current Data Center Server's certificate should be trusted, either by providing the certificate-fingerprint argument or by relying on a previously trusted certificate of this hostname. When set to true, trust the current Data Center Server's certificate as-is.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers. tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_nutanix_data_center_server",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point nutanix data center server.`,
			Description: `

This resource allows you to execute Check Point Nutanix Data Center Server.

`,
			Keywords: []string{
				"management",
				"nutanix",
				"data",
				"center",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required) IP Address or hostname of the Nutanix Prism server.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) Username of the Nutanix Prism server.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Password of the Nutanix Prism server.`,
				},
				resource.Attribute{
					Name:        "certificate_fingerprint",
					Description: `(Optional) Specify the SHA-1 or SHA-256 fingerprint of the Data Center Server's certificate.`,
				},
				resource.Attribute{
					Name:        "unsafe_auto_accept",
					Description: `(Optional) When set to false, the current Data Center Server's certificate should be trusted, either by providing the certificate-fingerprint argument or by relying on a previously trusted certificate of this hostname. When set to true, trust the current Data Center Server's certificate as-is.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
				resource.Attribute{
					Name:        "data_center_type",
					Description: `Data center type.`,
				},
				resource.Attribute{
					Name:        "automatic_refresh",
					Description: `Indicates whether the data center server's content is automatically updated.`,
				},
				resource.Attribute{
					Name:        "properties",
					Description: `Data center properties. properties blocks are documented beloe. ` + "`" + `properties` + "`" + ` supports the following:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_openstack_data_center_server",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point openstack data center server.`,
			Description: `

This resource allows you to execute Check Point OpenStack Data Center Server.

`,
			Keywords: []string{
				"management",
				"openstack",
				"data",
				"center",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required) URL of the OpenStack server. http(s)://<host>:<port>/<version>Example: https://1.2.3.4:5000/v2.0`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) Username of the OpenStack server. To login to specific domain insert domain name before username. Example: <domain>/<username>`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Password of the OpenStack server.`,
				},
				resource.Attribute{
					Name:        "password_base64",
					Description: `(Optional) Password of the OpenStack server encoded in Base64.`,
				},
				resource.Attribute{
					Name:        "certificate_fingerprint",
					Description: `(Optional) Specify the SHA-1 or SHA-256 fingerprint of the Data Center Server's certificate.`,
				},
				resource.Attribute{
					Name:        "unsafe_auto_accept",
					Description: `(Optional) When set to false, the current Data Center Server's certificate should be trusted, either by providing the certificate-fingerprint argument or by relying on a previously trusted certificate of this hostname. When set to true, trust the current Data Center Server's certificate as-is.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers. tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_opsec_application",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Opsec Application.`,
			Description: `

This resource allows you to execute Check Point Opsec Application.

`,
			Keywords: []string{
				"management",
				"opsec",
				"application",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "cpmi",
					Description: `(Optional) Used to setup the CPMI client entity.cpmi blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Optional) The host where the server is running. Pre-define the host as a network object.`,
				},
				resource.Attribute{
					Name:        "lea",
					Description: `(Optional) Used to setup the LEA client entity.lea blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "one_time_password",
					Description: `(Optional) A password required for establishing a Secure Internal Communication (SIC).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. ` + "`" + `cpmi` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "administrator_profile",
					Description: `(Optional) A profile to set the log reading permissions by for the client entity.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Whether to enable this client entity on the Opsec Application.`,
				},
				resource.Attribute{
					Name:        "use_administrator_credentials",
					Description: `(Optional) Whether to use the Admin's credentials to login to the security management server. ` + "`" + `lea` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "access_permissions",
					Description: `(Optional) Log reading permissions for the LEA client entity.`,
				},
				resource.Attribute{
					Name:        "administrator_profile",
					Description: `(Optional) A profile to set the log reading permissions by for the client entity.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Whether to enable this client entity on the Opsec Application.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_oracle_cloud_data_center_server",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point oracle cloud data center server.`,
			Description: `

This resource allows you to execute Check Point Oracle Cloud Data Center Server.

`,
			Keywords: []string{
				"management",
				"oracle",
				"cloud",
				"data",
				"center",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name. Must be unique in the domain.`,
				},
				resource.Attribute{
					Name:        "authentication_method",
					Description: `(Required) key-authentication uses the Service Account private key file to authenticate. vm-instance-authentication uses VM Instance to authenticate. This option requires the Security Management Server deployed in Oracle Cloud, and running in a dynamic group with the required permissions.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Required) An Oracle Cloud API key PEM file, encoded in base64. Required for authentication-method: key-authentication.`,
				},
				resource.Attribute{
					Name:        "key_user",
					Description: `(Required) An Oracle Cloud user id associated with key. Required for authentication-method: key-authentication.`,
				},
				resource.Attribute{
					Name:        "key_tenant",
					Description: `(Required) An Oracle Cloud tenancy id where the key was created. Required for authentication-method: key-authentication.`,
				},
				resource.Attribute{
					Name:        "key_region",
					Description: `(Required) An Oracle Cloud region for where to create scanner. Required for authentication-method: key-authentication.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
				resource.Attribute{
					Name:        "automatic_refresh",
					Description: `Indicates whether the data center server's content is automatically updated.`,
				},
				resource.Attribute{
					Name:        "data_center_type",
					Description: `Data center type.`,
				},
				resource.Attribute{
					Name:        "properties",
					Description: `Data center properties. properties blocks are documented below. ` + "`" + `properties` + "`" + ` supports the following:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_package",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to add/update/delete Check Point Package Object.`,
			Description: `

This resource allows you to add/update/delete Check Point Package Object.

`,
			Keywords: []string{
				"management",
				"package",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name. Should be unique in the domain.`,
				},
				resource.Attribute{
					Name:        "access",
					Description: `(Optional) True - enables, False - disables access & NAT policies, empty - nothing is changed.`,
				},
				resource.Attribute{
					Name:        "desktop_security",
					Description: `(Optional) True - enables, False - disables Desktop security policy, empty - nothing is changed.`,
				},
				resource.Attribute{
					Name:        "qos",
					Description: `(Optional) True - enables, False - disables QoS policy, empty - nothing is changed.`,
				},
				resource.Attribute{
					Name:        "qos_policy_type",
					Description: `(Optional) QoS policy type.`,
				},
				resource.Attribute{
					Name:        "threat_prevention",
					Description: `(Optional) True - enables, False - disables Threat policy, empty - nothing is changed.`,
				},
				resource.Attribute{
					Name:        "vpn_traditional_mode",
					Description: `(Optional) True - enables, False - disables VPN traditional mode, empty - nothing is changed.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_publish",
			Category:         "Management Resources",
			ShortDescription: `Publish Changes.`,
			Description: `

This command resource allows you to Publish Changes.

`,
			Keywords: []string{
				"management",
				"publish",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Session unique identifier. Specify it to publish a different session than the one you currently use.`,
				},
				resource.Attribute{
					Name:        "triggers",
					Description: `(Optional) Triggers a publish if there are any changes to objects in this list.`,
				},
				resource.Attribute{
					Name:        "task_id",
					Description: `(Computed) Asynchronous task unique identifier. ## How To Use Make sure this command resource will be executed by terraform when you meant it will run.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_put_file",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Put File.`,
			Description: `

This command resource allows you to execute Check Point Put File.

`,
			Keywords: []string{
				"management",
				"put",
				"file",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "targets",
					Description: `(Required) On what targets to execute this command. Targets may be identified by their name, or object unique identifier.targets blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "file_content",
					Description: `(Optional) Text file content.`,
				},
				resource.Attribute{
					Name:        "file_name",
					Description: `(Optional) Text file name.`,
				},
				resource.Attribute{
					Name:        "file_path",
					Description: `(Optional) Text file target path.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "tasks",
					Description: `(Computed) Collection of asynchronous task unique identifiers. ## How To Use Make sure this command will be executed in the right execution order.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_radius_group",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to add/update/delete Check Point Radius Group.`,
			Description: `

This resource allows you to add/update/delete Check Point Radius Group.

`,
			Keywords: []string{
				"management",
				"radius",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name. Must be unique in the domain.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `(Optional) Collection of radius servers identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_radius_server",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to add/update/delete Check Point Radius Server.`,
			Description: `

This resource allows you to add/update/delete Check Point Radius Server.

`,
			Keywords: []string{
				"management",
				"radius",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name. Should be unique in the domain.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `(Required) The UID or Name of the host that is the RADIUS Server.`,
				},
				resource.Attribute{
					Name:        "shared_secret",
					Description: `(Required) The secret between the RADIUS server and the Security Gateway.`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `(Optional) The UID or Name of the Service to which the RADIUS server listens.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) The version can be either RADIUS Version 1.0, which is RFC 2138 compliant, and RADIUS Version 2.0 which is RFC 2865 compliant.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) The type of authentication protocol that will be used when authenticating the user to the RADIUS server.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) The priority of the RADIUS Server in case it is a member of a RADIUS Group.`,
				},
				resource.Attribute{
					Name:        "accounting",
					Description: `(Optional) Accounting settings. accounting blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. ` + "`" + `accounting` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "enable_ip_pool_management",
					Description: `(Optional) IP pool management, enables Accounting service.`,
				},
				resource.Attribute{
					Name:        "accounting_service",
					Description: `(Optional) The UID or Name of the the accounting interface to notify the server when users login and logout which will then lock and release the IP addresses that the server allocated to those users.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_repository_script",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Repository Script.`,
			Description:      ``,
			Keywords: []string{
				"management",
				"repository",
				"script",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "script_body",
					Description: `(Optional) The entire content of the script.`,
				},
				resource.Attribute{
					Name:        "script_body_base64",
					Description: `(Optional) The entire content of the script encoded in Base64.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_reset_sic",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Reset Sic.`,
			Description:      ``,
			Keywords: []string{
				"management",
				"reset",
				"sic",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Gateway, cluster member or Check Point host unique identifier.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Gateway, cluster member or Check Point host name. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_restore_domain",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Restore Domain.`,
			Description: `

This command resource allows you to execute Check Point Restore Domain.

`,
			Keywords: []string{
				"management",
				"restore",
				"domain",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "file_path",
					Description: `(Required) Path to the backup file to be restored. <br>Should be the full file path (example, "/var/log/domain1_backup.tgz").`,
				},
				resource.Attribute{
					Name:        "domain_ip_address",
					Description: `(Required) IPv4 address.<br><font color="red">Required only for</font> importing Security Management Server into Multi-Domain Server.`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `(Required) Domain name. Should be unique in the MDS.<br><font color="red">Required only for</font> importing Security Management Server into Multi-Domain Server.`,
				},
				resource.Attribute{
					Name:        "domain_server_name",
					Description: `(Required) Multi Domain server name.<br><font color="red">Required only for</font> importing Security Management Server into Multi-Domain Server.`,
				},
				resource.Attribute{
					Name:        "verify_only",
					Description: `(Optional) If true, verify that the import operation is valid for this input file and this environment <br>Note: Restore operation will not be executed.`,
				},
				resource.Attribute{
					Name:        "task_id",
					Description: `(Computed) Asynchronous task unique identifier. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_revert_to_revision",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Revert To Revision.`,
			Description: `

This command resource allows you to execute Check Point Revert To Revision.

`,
			Keywords: []string{
				"management",
				"revert",
				"to",
				"revision",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "to_session",
					Description: `(Optional) Session unique identifier. Specify the session id you would like to revert your database to.`,
				},
				resource.Attribute{
					Name:        "task_id",
					Description: `(Computed) Asynchronous task unique identifier. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_run_ips_update",
			Category:         "Management Resources",
			ShortDescription: `Runs IPS database update. If "package-path" is not provided server will try to get the latest package from the User Center.`,
			Description: `

This command resource allows you to Runs IPS database update. If "package-path" is not provided server will try to get the latest package from the User Center.

`,
			Keywords: []string{
				"management",
				"run",
				"ips",
				"update",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "package_path",
					Description: `(Optional) Offline update package path.`,
				},
				resource.Attribute{
					Name:        "task_id",
					Description: `(Computed) Asynchronous task unique identifier. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_run_script",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Run Script.`,
			Description: `

This command resource allows you to execute Check Point Run Script.

`,
			Keywords: []string{
				"management",
				"run",
				"script",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "script_name",
					Description: `(Required) Script name.`,
				},
				resource.Attribute{
					Name:        "script",
					Description: `(Required) Script body.`,
				},
				resource.Attribute{
					Name:        "targets",
					Description: `(Required) On what targets to execute this command. Targets may be identified by their name, or object unique identifier.targets blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "args",
					Description: `(Optional) Script arguments.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "tasks",
					Description: `(Computed) Collection of asynchronous task unique identifiers. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_run_threat_emulation_file_types_offline_update",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Run Threat Emulation File Types Offline Update.`,
			Description: `

This command resource allows you to execute Check Point Run Threat Emulation File Types Offline Update.

`,
			Keywords: []string{
				"management",
				"run",
				"threat",
				"emulation",
				"file",
				"types",
				"offline",
				"update",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "file_path",
					Description: `(Required) File path for offline update of Threat Emulation file types, the file path should be on the management machine.`,
				},
				resource.Attribute{
					Name:        "file_raw_data",
					Description: `(Required) The contents of a file containing the Threat Emulation file types. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_security_zone",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Security Zone.`,
			Description: `

This resource allows you to execute Check Point Security Zone.

`,
			Keywords: []string{
				"management",
				"security",
				"zone",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_service_citrix_tcp",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Service Citrix Tcp.`,
			Description: `

This resource allows you to execute Check Point Service Citrix Tcp.

`,
			Keywords: []string{
				"management",
				"service",
				"citrix",
				"tcp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "application",
					Description: `(Optional) Citrix application name.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_service_compound_tcp",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Service Compound Tcp.`,
			Description: `

This resource allows you to execute Check Point Service Compound Tcp.

`,
			Keywords: []string{
				"management",
				"service",
				"compound",
				"tcp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "compound_service",
					Description: `(Optional) Compound service type.`,
				},
				resource.Attribute{
					Name:        "keep_connections_open_after_policy_installation",
					Description: `(Optional) Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_service_dce_rpc",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Service Dce Rpc.`,
			Description: `

This resource allows you to execute Check Point Service Dce Rpc.

`,
			Keywords: []string{
				"management",
				"service",
				"dce",
				"rpc",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "interface_uuid",
					Description: `(Optional) Network interface UUID.`,
				},
				resource.Attribute{
					Name:        "keep_connections_open_after_policy_installation",
					Description: `(Optional) Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_service_group",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to add/update/delete Check Point Service Group.`,
			Description: `

This resource allows you to add/update/delete Check Point Service Group.

`,
			Keywords: []string{
				"management",
				"service",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name. Should be unique in the domain.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `(Optional) Collection of Network objects identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_service_icmp",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Service Icmp.`,
			Description: `

This resource allows you to execute Check Point Service Icmp.

`,
			Keywords: []string{
				"management",
				"service",
				"icmp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "icmp_code",
					Description: `(Optional) As listed in: <a href="http://www.iana.org/assignments/icmp-parameters" target="_blank">RFC 792</a>.`,
				},
				resource.Attribute{
					Name:        "icmp_type",
					Description: `(Optional) As listed in: <a href="http://www.iana.org/assignments/icmp-parameters" target="_blank">RFC 792</a>.`,
				},
				resource.Attribute{
					Name:        "keep_connections_open_after_policy_installation",
					Description: `(Optional) Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_service_icmp6",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Service Icmp6.`,
			Description: `

This resource allows you to execute Check Point Service Icmp6.

`,
			Keywords: []string{
				"management",
				"service",
				"icmp6",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "icmp_code",
					Description: `(Optional) As listed in: <a href="http://www.iana.org/assignments/icmp-parameters" target="_blank">RFC 792</a>.`,
				},
				resource.Attribute{
					Name:        "icmp_type",
					Description: `(Optional) As listed in: <a href="http://www.iana.org/assignments/icmp-parameters" target="_blank">RFC 792</a>.`,
				},
				resource.Attribute{
					Name:        "keep_connections_open_after_policy_installation",
					Description: `(Optional) Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_service_other",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Service Other.`,
			Description: `

This resource allows you to execute Check Point Service Other.

`,
			Keywords: []string{
				"management",
				"service",
				"other",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "accept_replies",
					Description: `(Optional) Specifies whether Other Service replies are to be accepted.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Contains an INSPECT expression that defines the action to take if a rule containing this service is matched. Example: set r_mhandler &open_ssl_handler sets a handler on the connection.`,
				},
				resource.Attribute{
					Name:        "aggressive_aging",
					Description: `(Optional) Sets short (aggressive) timeouts for idle connections.aggressive_aging blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `(Optional) IP protocol number.`,
				},
				resource.Attribute{
					Name:        "keep_connections_open_after_policy_installation",
					Description: `(Optional) Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections.`,
				},
				resource.Attribute{
					Name:        "match",
					Description: `(Optional) Contains an INSPECT expression that defines the matching criteria. The connection is examined against the expression during the first packet. Example: tcp, dport = 21, direction = 0 matches incoming FTP control connections.`,
				},
				resource.Attribute{
					Name:        "match_for_any",
					Description: `(Optional) Indicates whether this service is used when 'Any' is set as the rule's service and there are several service objects with the same source port and protocol.`,
				},
				resource.Attribute{
					Name:        "override_default_settings",
					Description: `(Optional) Indicates whether this service is a Data Domain service which has been overridden.`,
				},
				resource.Attribute{
					Name:        "session_timeout",
					Description: `(Optional) Time (in seconds) before the session times out.`,
				},
				resource.Attribute{
					Name:        "sync_connections_on_cluster",
					Description: `(Optional) Enables state-synchronized High Availability or Load Sharing on a ClusterXL or OPSEC-certified cluster.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "use_default_session_timeout",
					Description: `(Optional) Use default virtual session timeout.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. ` + "`" + `aggressive_aging` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "default_timeout",
					Description: `(Optional) Default aggressive aging timeout in seconds.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional) N/A`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) Aggressive aging timeout in seconds.`,
				},
				resource.Attribute{
					Name:        "use_default_timeout",
					Description: `(Optional) N/A`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_service_rpc",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Service Rpc.`,
			Description: `

This resource allows you to execute Check Point Service Rpc.

`,
			Keywords: []string{
				"management",
				"service",
				"rpc",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "keep_connections_open_after_policy_installation",
					Description: `(Optional) Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections.`,
				},
				resource.Attribute{
					Name:        "program_number",
					Description: `(Optional) N/A`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_service_sctp",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Service Sctp.`,
			Description: `

This resource allows you to execute Check Point Service Sctp.

`,
			Keywords: []string{
				"management",
				"service",
				"sctp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "aggressive_aging",
					Description: `(Optional) Sets short (aggressive) timeouts for idle connections.aggressive_aging blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "keep_connections_open_after_policy_installation",
					Description: `(Optional) Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections.`,
				},
				resource.Attribute{
					Name:        "match_for_any",
					Description: `(Optional) Indicates whether this service is used when 'Any' is set as the rule's service and there are several service objects with the same source port and protocol.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Port number. To specify a port range add a hyphen between the lowest and the highest port numbers, for example 44-45.`,
				},
				resource.Attribute{
					Name:        "session_timeout",
					Description: `(Optional) Time (in seconds) before the session times out.`,
				},
				resource.Attribute{
					Name:        "source_port",
					Description: `(Optional) Source port number. To specify a port range add a hyphen between the lowest and the highest port numbers, for example 44-45.`,
				},
				resource.Attribute{
					Name:        "sync_connections_on_cluster",
					Description: `(Optional) Enables state-synchronized High Availability or Load Sharing on a ClusterXL or OPSEC-certified cluster.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "use_default_session_timeout",
					Description: `(Optional) Use default virtual session timeout.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. ` + "`" + `aggressive_aging` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "default_timeout",
					Description: `(Optional) Default aggressive aging timeout in seconds.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional) N/A`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) Aggressive aging timeout in seconds.`,
				},
				resource.Attribute{
					Name:        "use_default_timeout",
					Description: `(Optional) N/A`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_service_tcp",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to add/update/delete Check Point Service Tcp.`,
			Description: `

This resource allows you to add/update/delete Check Point Service Tcp.

`,
			Keywords: []string{
				"management",
				"service",
				"tcp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name. Should be unique in the domain.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The number of the port used to provide this service. To specify a port range, place a hyphen between the lowest and highest port numbers, for example 44-55.`,
				},
				resource.Attribute{
					Name:        "aggressive_aging",
					Description: `(Optional) Sets short (aggressive) timeouts for idle connections. Aggressive Aging blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "keep_connections_open_after_policy_installation",
					Description: `(Optional) Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections.`,
				},
				resource.Attribute{
					Name:        "match_by_protocol_signature",
					Description: `(Optional) A value of true enables matching by the selected protocol's signature - the signature identifies the protocol as genuine. Select this option to limit the port to the specified protocol. If the selected protocol does not support matching by signature, this field cannot be set to true.`,
				},
				resource.Attribute{
					Name:        "match_for_any",
					Description: `(Optional) Indicates whether this service is used when 'Any' is set as the rule's service and there are several service objects with the same source port and protocol.`,
				},
				resource.Attribute{
					Name:        "override_default_settings",
					Description: `(Optional) Indicates whether this service is a Data Domain service which has been overridden.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) Select the protocol type associated with the service, and by implication, the management server (if any) that enforces Content Security and Authentication for the service. Selecting a Protocol Type invokes the specific protocol handlers for each protocol type, thus enabling higher level of security by parsing the protocol, and higher level of connectivity by tracking dynamic actions (such as opening of ports).`,
				},
				resource.Attribute{
					Name:        "session_timeout",
					Description: `(Optional) Time (in seconds) before the session times out.`,
				},
				resource.Attribute{
					Name:        "source_port",
					Description: `(Optional) Port number for the client side service. If specified, only those Source port Numbers will be Accepted, Dropped, or Rejected during packet inspection. Otherwise, the source port is not inspected.`,
				},
				resource.Attribute{
					Name:        "sync_connections_on_cluster",
					Description: `(Optional)Enables state-synchronized High Availability or Load Sharing on a ClusterXL or OPSEC-certified cluster.`,
				},
				resource.Attribute{
					Name:        "use_default_session_timeout",
					Description: `(Optional) Use default virtual session timeout.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers. ` + "`" + `aggressive_aging` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "default_timeout",
					Description: `(Optional) Default aggressive aging timeout in seconds.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional) N/A`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) Aggressive aging timeout in seconds.`,
				},
				resource.Attribute{
					Name:        "use_default_timeout",
					Description: `(Optional) N/A.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_service_udp",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to add/update/delete Check Point Service Udp.`,
			Description: `

This resource allows you to add/update/delete Check Point Service Udp.

`,
			Keywords: []string{
				"management",
				"service",
				"udp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name. Should be unique in the domain.`,
				},
				resource.Attribute{
					Name:        "accept_replies",
					Description: `(Optional) N/A.`,
				},
				resource.Attribute{
					Name:        "aggressive_aging",
					Description: `(Optional) Sets short (aggressive) timeouts for idle connections. Aggressive Aging blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "keep_connections_open_after_policy_installation",
					Description: `(Optional) Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections.`,
				},
				resource.Attribute{
					Name:        "match_by_protocol_signature",
					Description: `(Optional) A value of true enables matching by the selected protocol's signature - the signature identifies the protocol as genuine. Select this option to limit the port to the specified protocol. If the selected protocol does not support matching by signature, this field cannot be set to true.`,
				},
				resource.Attribute{
					Name:        "match_for_any",
					Description: `(Optional) Indicates whether this service is used when 'Any' is set as the rule's service and there are several service objects with the same source port and protocol.`,
				},
				resource.Attribute{
					Name:        "override_default_settings",
					Description: `(Optional) Indicates whether this service is a Data Domain service which has been overridden.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The number of the port used to provide this service. To specify a port range, place a hyphen between the lowest and highest port numbers, for example 44-55.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) Select the protocol type associated with the service, and by implication, the management server (if any) that enforces Content Security and Authentication for the service. Selecting a Protocol Type invokes the specific protocol handlers for each protocol type, thus enabling higher level of security by parsing the protocol, and higher level of connectivity by tracking dynamic actions (such as opening of ports).`,
				},
				resource.Attribute{
					Name:        "session_timeout",
					Description: `(Optional) Time (in seconds) before the session times out.`,
				},
				resource.Attribute{
					Name:        "source_port",
					Description: `(Optional) Port number for the client side service. If specified, only those Source port Numbers will be Accepted, Dropped, or Rejected during packet inspection. Otherwise, the source port is not inspected.`,
				},
				resource.Attribute{
					Name:        "sync_connections_on_cluster",
					Description: `(Optional)Enables state-synchronized High Availability or Load Sharing on a ClusterXL or OPSEC-certified cluster.`,
				},
				resource.Attribute{
					Name:        "use_default_session_timeout",
					Description: `(Optional) Use default virtual session timeout.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers. ` + "`" + `aggressive_aging` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "default_timeout",
					Description: `(Optional) Default aggressive aging timeout in seconds.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional) N/A`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) Aggressive aging timeout in seconds.`,
				},
				resource.Attribute{
					Name:        "use_default_timeout",
					Description: `(Optional) N/A.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_set_api_settings",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Set Api Settings.`,
			Description: `

This command resource allows you to execute Check Point Set Api Settings.

`,
			Keywords: []string{
				"management",
				"set",
				"api",
				"settings",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "accepted_api_calls_from",
					Description: `(Optional) Clients allowed to connect to the API Server.`,
				},
				resource.Attribute{
					Name:        "automatic_start",
					Description: `(Optional) MGMT API will start after server will start. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_set_automatic_purge",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Set Automatic Purge.`,
			Description: `

This command resource allows you to execute Check Point Set Automatic Purge.

`,
			Keywords: []string{
				"management",
				"set",
				"automatic",
				"purge",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Turn on/off the automatic-purge feature.`,
				},
				resource.Attribute{
					Name:        "keep_sessions_by_count",
					Description: `(Optional) Whether or not to keep the latest N sessions. Note: when the automatic purge feature is enabled, this field and/or the "keep-sessions-by-date" field must be set to 'true'.`,
				},
				resource.Attribute{
					Name:        "number_of_sessions_to_keep",
					Description: `(Optional) When "keep-sessions-by-count = true" this sets the number of newest sessions to preserve, by the sessions's publish date.`,
				},
				resource.Attribute{
					Name:        "keep_sessions_by_days",
					Description: `(Optional) Whether or not to keep the sessions for D days. Note: when the automatic purge feature is enabled, this field and/or the "keep-sessions-by-count" field must be set to 'true'.`,
				},
				resource.Attribute{
					Name:        "number_of_days_to_keep",
					Description: `(Optional) When "keep-sessions-by-days = true" this sets the number of days to keep the sessions.`,
				},
				resource.Attribute{
					Name:        "scheduling",
					Description: `(Optional) When to purge sessions that do not meet the "keep" criteria. Note: when the automatic purge feature is enabled, this field must be set.scheduling blocks are documented below. ` + "`" + `scheduling` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "start_date",
					Description: `(Optional) The first time to check whether or not there are sessions to purge. ISO 8601. If timezone isn't specified in the input, the Management server's timezone is used. Instead - If you want to start immediately, type: "now". Note: when the automatic purge feature is enabled, this field must be set.`,
				},
				resource.Attribute{
					Name:        "time_units",
					Description: `(Optional) Note: when the automatic purge feature is enabled, this field must be set.`,
				},
				resource.Attribute{
					Name:        "check_interval",
					Description: `(Optional) Number of time-units between two purge checks. Note: when the automatic purge feature is enabled, this field must be set. ## How To Use Make sure this command will be executed in the right execution order.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_set_global_domain",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Set Global Domain.`,
			Description: `

This command resource allows you to execute Check Point Set Global Domain.

`,
			Keywords: []string{
				"management",
				"set",
				"global",
				"domain",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "servers",
					Description: `(Optional) Multi Domain Servers. When the field is provided, 'set-global-domain' command is executed asynchronously.servers blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers. Note: The list of tags can not be modified in a singlecommand together with the domain servers. To modify tags, please use the separate 'set-global-domain' command, without providing the list of domain servers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
				resource.Attribute{
					Name:        "tasks",
					Description: `(Computed) Collection of asynchronous task unique identifiers. ` + "`" + `servers` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "add",
					Description: `(Optional) Adds to collection of valuesadd blocks are documented below. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_set_ha_state",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Set HA State.`,
			Description: `

This command resource allows you to execute Check Point Set HA State.

`,
			Keywords: []string{
				"management",
				"set",
				"ha",
				"state",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "new_state",
					Description: `(Required) Domain server new state.`,
				},
				resource.Attribute{
					Name:        "task_id",
					Description: `(Computed) Asynchronous task unique identifier. ## How To Use Make sure this command will be executed in the right execution order.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_set_idp_default_assignment",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Set Idp Default Assignment.`,
			Description:      ``,
			Keywords: []string{
				"management",
				"set",
				"idp",
				"default",
				"assignment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "identity_provider",
					Description: `(Required) Represents the Identity Provider to be used for Login by this assignment identified by the name or UID, to cancel existing assignment should set to 'none'.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_set_ips_update_schedule",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Set Ips Update Schedule.`,
			Description: `

This command resource allows you to execute Check Point Set Ips Update Schedule.

`,
			Keywords: []string{
				"management",
				"set",
				"ips",
				"update",
				"schedule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Enable/Disable IPS Update Schedule.`,
				},
				resource.Attribute{
					Name:        "time",
					Description: `(Optional) Time in format HH:mm.`,
				},
				resource.Attribute{
					Name:        "recurrence",
					Description: `(Optional) Days recurrence.recurrence blocks are documented below. ` + "`" + `recurrence` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "days",
					Description: `(Optional) Valid on specific days. Multiple options, support range of days in months. Example:["1","3","9-20"].days blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "minutes",
					Description: `(Optional) Valid on interval. The length of time in minutes between updates.`,
				},
				resource.Attribute{
					Name:        "pattern",
					Description: `(Optional) Valid on "Interval", "Daily", "Weekly", "Monthly" base.`,
				},
				resource.Attribute{
					Name:        "weekdays",
					Description: `(Optional) Valid on weekdays. Example: "Sun", "Mon"..."Sat".weekdays blocks are documented below. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_set_login_message",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Set Login Message.`,
			Description: `

This command resource allows you to execute Check Point Set Login Message.

`,
			Keywords: []string{
				"management",
				"set",
				"login",
				"message",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "header",
					Description: `(Optional) Login message header.`,
				},
				resource.Attribute{
					Name:        "message",
					Description: `(Optional) Login message body.`,
				},
				resource.Attribute{
					Name:        "show_message",
					Description: `(Optional) Whether to show login message.`,
				},
				resource.Attribute{
					Name:        "warning",
					Description: `(Optional) Add warning sign. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_set_threat_protection",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Set Threat Protection.`,
			Description: `

This command resource allows you to execute Check Point Set Threat Protection.

`,
			Keywords: []string{
				"management",
				"set",
				"threat",
				"protection",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Protection comments.`,
				},
				resource.Attribute{
					Name:        "follow_up",
					Description: `(Optional) Tag the protection with pre-defined follow-up flag.`,
				},
				resource.Attribute{
					Name:        "overrides",
					Description: `(Optional) Overrides per profile for this protection<br> Note: Remove override for Core protections removes only the actions override. Remove override for Threat Cloud protections removes the action, track and packet captures.overrides blocks are documented below. ` + "`" + `overrides` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Protection action.`,
				},
				resource.Attribute{
					Name:        "profile",
					Description: `(Optional) Profile name.`,
				},
				resource.Attribute{
					Name:        "capture_packets",
					Description: `(Optional) Capture packets.`,
				},
				resource.Attribute{
					Name:        "track",
					Description: `(Optional) Tracking method for protection. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_simple_cluster",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Simple Cluster.`,
			Description:      ``,
			Keywords: []string{
				"management",
				"simple",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "advanced_settings",
					Description: `(Optional) N/Aadvanced_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "anti_bot",
					Description: `(Optional) Anti-Bot blade enabled.`,
				},
				resource.Attribute{
					Name:        "anti_virus",
					Description: `(Optional) Anti-Virus blade enabled.`,
				},
				resource.Attribute{
					Name:        "application_control",
					Description: `(Optional) Application Control blade enabled.`,
				},
				resource.Attribute{
					Name:        "application_control_and_url_filtering_settings",
					Description: `(Optional) Gateway Application Control and URL filtering settings.application_control_and_url_filtering_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "cluster_mode",
					Description: `(Optional) Cluster mode.`,
				},
				resource.Attribute{
					Name:        "cluster_settings",
					Description: `(Optional) ClusterXL and VRRP Settings.cluster_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "content_awareness",
					Description: `(Optional) Content Awareness blade enabled.`,
				},
				resource.Attribute{
					Name:        "enable_https_inspection",
					Description: `(Optional) Enable HTTPS Inspection after defining an outbound inspection certificate. <br>To define the outbound certificate use outbound inspection certificate API.`,
				},
				resource.Attribute{
					Name:        "fetch_policy",
					Description: `(Optional) Security management server(s) to fetch the policy from.fetch_policy blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "firewall",
					Description: `(Optional) Firewall blade enabled.`,
				},
				resource.Attribute{
					Name:        "firewall_settings",
					Description: `(Optional) N/Afirewall_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "geo_mode",
					Description: `(Optional) Cluster High Availability Geo mode.<br>This setting applies only to a cluster deployed in a cloud. Available when the cluster mode equals "cluster-xl-ha".`,
				},
				resource.Attribute{
					Name:        "hardware",
					Description: `(Optional) Cluster platform hardware.`,
				},
				resource.Attribute{
					Name:        "hit_count",
					Description: `(Optional) Hit count tracks the number of connections each rule matches.`,
				},
				resource.Attribute{
					Name:        "https_inspection",
					Description: `(Optional) HTTPS inspection.https_inspection blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "identity_awareness",
					Description: `(Optional) Identity awareness blade enabled.`,
				},
				resource.Attribute{
					Name:        "identity_awareness_settings",
					Description: `(Optional) Gateway Identity Awareness settings.identity_awareness_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `(Optional) Cluster interfaces.interfaces blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `(Optional) IPv4 address.`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `(Optional) IPv6 address.`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `(Optional) Intrusion Prevention System blade enabled.`,
				},
				resource.Attribute{
					Name:        "ips_update_policy",
					Description: `(Optional) Specifies whether the IPS will be downloaded from the Management or directly to the Gateway.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `(Optional) Cluster members list. Only new cluster member can be added. Adding existing gateway is not supported.members blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "nat_hide_internal_interfaces",
					Description: `(Optional) Hide internal networks behind the Gateway's external IP.`,
				},
				resource.Attribute{
					Name:        "nat_settings",
					Description: `(Optional) NAT settings.nat_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "os_name",
					Description: `(Optional) Cluster platform operating system.`,
				},
				resource.Attribute{
					Name:        "platform_portal_settings",
					Description: `(Optional) Platform portal settings.platform_portal_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "proxy_settings",
					Description: `(Optional) Proxy Server for Gateway.proxy_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "qos",
					Description: `(Optional) QoS.`,
				},
				resource.Attribute{
					Name:        "send_alerts_to_server",
					Description: `(Optional) Server(s) to send alerts to.send_alerts_to_server blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "send_logs_to_backup_server",
					Description: `(Optional) Backup server(s) to send logs to.send_logs_to_backup_server blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "send_logs_to_server",
					Description: `(Optional) Server(s) to send logs to.send_logs_to_server blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "threat_emulation",
					Description: `(Optional) Threat Emulation blade enabled.`,
				},
				resource.Attribute{
					Name:        "threat_extraction",
					Description: `(Optional) Threat Extraction blade enabled.`,
				},
				resource.Attribute{
					Name:        "threat_prevention_mode",
					Description: `(Optional) The mode of Threat Prevention to use. When using Autonomous Threat Prevention, disabling the Threat Prevention blades is not allowed.`,
				},
				resource.Attribute{
					Name:        "url_filtering",
					Description: `(Optional) URL Filtering blade enabled.`,
				},
				resource.Attribute{
					Name:        "usercheck_portal_settings",
					Description: `(Optional) UserCheck portal settings.usercheck_portal_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) Cluster platform version.`,
				},
				resource.Attribute{
					Name:        "vpn",
					Description: `(Optional) VPN blade enabled.`,
				},
				resource.Attribute{
					Name:        "vpn_settings",
					Description: `(Optional) Gateway VPN settings.vpn_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "zero_phishing",
					Description: `(Optional) Zero Phishing blade enabled.`,
				},
				resource.Attribute{
					Name:        "zero_phishing_fqdn",
					Description: `(Optional) Zero Phishing gateway FQDN.`,
				},
				resource.Attribute{
					Name:        "show_portals_certificate",
					Description: `(Optional) Indicates whether to show the portals certificate value in the reply.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `(Optional) Collection of group identifiers.groups blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. ` + "`" + `advanced_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "connection_persistence",
					Description: `(Optional) Handling established connections when installing a new policy.`,
				},
				resource.Attribute{
					Name:        "sam",
					Description: `(Optional) SAM.sam blocks are documented below. ` + "`" + `application_control_and_url_filtering_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "global_settings_mode",
					Description: `(Optional) Whether to override global settings or not.`,
				},
				resource.Attribute{
					Name:        "override_global_settings",
					Description: `(Optional) override global settings object.override_global_settings blocks are documented below. ` + "`" + `cluster_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "member_recovery_mode",
					Description: `(Optional) In a High Availability cluster, each member is given a priority. The member with the highest priority serves as the gateway. If this gateway fails, control is passed to the member with the next highest priority. If that member fails, control is passed to the next, and so on. Upon gateway recovery, it is possible to: Maintain current active Cluster Member (maintain-current-active) or Switch to higher priority Cluster Member (according-to-priority).`,
				},
				resource.Attribute{
					Name:        "state_synchronization",
					Description: `(Optional) Cluster State Synchronization settings.state_synchronization blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "track_changes_of_cluster_members",
					Description: `(Optional) Track changes in the status of Cluster Members.`,
				},
				resource.Attribute{
					Name:        "use_virtual_mac",
					Description: `(Optional) Use Virtual MAC. By enabling Virtual MAC in ClusterXL High Availability New mode, or Load Sharing Unicast mode, all cluster members associate the same Virtual MAC address with All Cluster Virtual Interfaces and the Virtual IP address. ` + "`" + `firewall_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "auto_calculate_connections_hash_table_size_and_memory_pool",
					Description: `(Optional) N/A`,
				},
				resource.Attribute{
					Name:        "auto_maximum_limit_for_concurrent_connections",
					Description: `(Optional) N/A`,
				},
				resource.Attribute{
					Name:        "connections_hash_size",
					Description: `(Optional) N/A`,
				},
				resource.Attribute{
					Name:        "maximum_limit_for_concurrent_connections",
					Description: `(Optional) N/A`,
				},
				resource.Attribute{
					Name:        "maximum_memory_pool_size",
					Description: `(Optional) N/A`,
				},
				resource.Attribute{
					Name:        "memory_pool_size",
					Description: `(Optional) N/A ` + "`" + `https_inspection` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "bypass_on_failure",
					Description: `(Optional) Set to be true in order to bypass all requests (Fail-open) in case of internal system error.bypass_on_failure blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "site_categorization_allow_mode",
					Description: `(Optional) Set to 'background' in order to allowed requests until categorization is complete.site_categorization_allow_mode blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "deny_untrusted_server_cert",
					Description: `(Optional) Set to be true in order to drop traffic from servers with untrusted server certificate.deny_untrusted_server_cert blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "deny_revoked_server_cert",
					Description: `(Optional) Set to be true in order to drop traffic from servers with revoked server certificate (validate CRL).deny_revoked_server_cert blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "deny_expired_server_cert",
					Description: `(Optional) Set to be true in order to drop traffic from servers with expired server certificate.deny_expired_server_cert blocks are documented below. ` + "`" + `identity_awareness_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "browser_based_authentication",
					Description: `(Optional) Enable Browser Based Authentication source.`,
				},
				resource.Attribute{
					Name:        "browser_based_authentication_settings",
					Description: `(Optional) Browser Based Authentication settings.browser_based_authentication_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "identity_agent",
					Description: `(Optional) Enable Identity Agent source.`,
				},
				resource.Attribute{
					Name:        "identity_agent_settings",
					Description: `(Optional) Identity Agent settings.identity_agent_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "identity_collector",
					Description: `(Optional) Enable Identity Collector source.`,
				},
				resource.Attribute{
					Name:        "identity_collector_settings",
					Description: `(Optional) Identity Collector settings.identity_collector_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "identity_sharing_settings",
					Description: `(Optional) Identity sharing settings.identity_sharing_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "proxy_settings",
					Description: `(Optional) Identity-Awareness Proxy settings.proxy_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "remote_access",
					Description: `(Optional) Enable Remote Access Identity source. ` + "`" + `interfaces` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name. Must be unique in the domain.`,
				},
				resource.Attribute{
					Name:        "interface_type",
					Description: `(Optional) Cluster interface type.`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `(Optional) IPv4 address.`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `(Optional) IPv6 address.`,
				},
				resource.Attribute{
					Name:        "network_mask",
					Description: `(Optional) IPv4 or IPv6 network mask. If both masks are required use ipv4-network-mask and ipv6-network-mask fields explicitly. Instead of providing mask itself it is possible to specify IPv4 or IPv6 mask length in mask-length field. If both masks length are required use ipv4-mask-length and ipv6-mask-length fields explicitly.`,
				},
				resource.Attribute{
					Name:        "ipv4_network_mask",
					Description: `(Optional) IPv4 network address.`,
				},
				resource.Attribute{
					Name:        "ipv6_network_mask",
					Description: `(Optional) IPv6 network address.`,
				},
				resource.Attribute{
					Name:        "ipv4_mask_length",
					Description: `(Optional) IPv4 network mask length.`,
				},
				resource.Attribute{
					Name:        "ipv6_mask_length",
					Description: `(Optional) IPv6 network mask length.`,
				},
				resource.Attribute{
					Name:        "anti_spoofing",
					Description: `(Optional) N/A`,
				},
				resource.Attribute{
					Name:        "anti_spoofing_settings",
					Description: `(Optional) N/Aanti_spoofing_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "multicast_address",
					Description: `(Optional) Multicast IP Address.`,
				},
				resource.Attribute{
					Name:        "multicast_address_type",
					Description: `(Optional) Multicast Address Type.`,
				},
				resource.Attribute{
					Name:        "security_zone",
					Description: `(Optional) N/A`,
				},
				resource.Attribute{
					Name:        "security_zone_settings",
					Description: `(Optional) N/Asecurity_zone_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "topology",
					Description: `(Optional) N/A`,
				},
				resource.Attribute{
					Name:        "topology_settings",
					Description: `(Optional) N/Atopology_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. ` + "`" + `members` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `(Optional) Cluster Member network interfaces.interfaces blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `(Optional) IPv4 address.`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `(Optional) IPv6 address.`,
				},
				resource.Attribute{
					Name:        "one_time_password",
					Description: `(Optional) N/A`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) In a High Availability New mode cluster each machine is given a priority. The highest priority machine serves as the gateway in normal circumstances. If this machine fails, control is passed to the next highest priority machine. If that machine fails, control is passed to the next machine, and so on. In Load Sharing Unicast mode cluster, the highest priority is the pivot machine. The values must be in a range from 1 to N, where N is number of cluster members.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. ` + "`" + `nat_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "auto_rule",
					Description: `(Optional) Whether to add automatic address translation rules.`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `(Optional) IPv4 address.`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `(Optional) IPv6 address.`,
				},
				resource.Attribute{
					Name:        "hide_behind",
					Description: `(Optional) Hide behind method. This parameter is forbidden in case "method" parameter is "static".`,
				},
				resource.Attribute{
					Name:        "install_on",
					Description: `(Optional) Which gateway should apply the NAT translation.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Optional) NAT translation method. ` + "`" + `platform_portal_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "portal_web_settings",
					Description: `(Optional) Configuration of the portal web settings.portal_web_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "certificate_settings",
					Description: `(Optional) Configuration of the portal certificate settings.certificate_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "accessibility",
					Description: `(Optional) Configuration of the portal access settings.accessibility blocks are documented below. ` + "`" + `proxy_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "use_custom_proxy",
					Description: `(Optional) Use custom proxy settings for this network object.`,
				},
				resource.Attribute{
					Name:        "proxy_server",
					Description: `(Optional) N/A`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) N/A ` + "`" + `usercheck_portal_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) State of the web portal (enabled or disabled). The supported blades are: {'Application Control', 'URL Filtering', 'Data Loss Prevention', 'Anti Virus', 'Anti Bot', 'Threat Emulation', 'Threat Extraction', 'Data Awareness'}.`,
				},
				resource.Attribute{
					Name:        "portal_web_settings",
					Description: `(Optional) Configuration of the portal web settings.portal_web_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "certificate_settings",
					Description: `(Optional) Configuration of the portal certificate settings.certificate_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "accessibility",
					Description: `(Optional) Configuration of the portal access settings.accessibility blocks are documented below. ` + "`" + `vpn_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "authentication",
					Description: `(Optional) Authentication.authentication blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "link_selection",
					Description: `(Optional) Link Selection.link_selection blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "maximum_concurrent_ike_negotiations",
					Description: `(Optional) N/A`,
				},
				resource.Attribute{
					Name:        "maximum_concurrent_tunnels",
					Description: `(Optional) N/A`,
				},
				resource.Attribute{
					Name:        "office_mode",
					Description: `(Optional) Office Mode. Notation Wide Impact - Office Mode apply IPSec VPN Software Blade clients and to the Mobile Access Software Blade clients.office_mode blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "remote_access",
					Description: `(Optional) Remote Access.remote_access blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "vpn_domain",
					Description: `(Optional) Gateway VPN domain identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "vpn_domain_exclude_external_ip_addresses",
					Description: `(Optional) Exclude the external IP addresses from the VPN domain of this Security Gateway.`,
				},
				resource.Attribute{
					Name:        "vpn_domain_type",
					Description: `(Optional) Gateway VPN domain type. ` + "`" + `sam` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "forward_to_other_sam_servers",
					Description: `(Optional) Forward SAM clients' requests to other SAM servers.`,
				},
				resource.Attribute{
					Name:        "use_early_versions",
					Description: `(Optional) Use early versions compatibility mode.use_early_versions blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "purge_sam_file",
					Description: `(Optional) Purge SAM File.purge_sam_file blocks are documented below. ` + "`" + `override_global_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "fail_mode",
					Description: `(Optional) Fail mode - allow or block all requests.`,
				},
				resource.Attribute{
					Name:        "website_categorization",
					Description: `(Optional) Website categorization object.website_categorization blocks are documented below. ` + "`" + `state_synchronization` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "delayed",
					Description: `(Optional) Start synchronizing with delay of seconds, as defined by delayed-seconds, after connection initiation. Disabled when state-synchronization disabled.`,
				},
				resource.Attribute{
					Name:        "delayed_seconds",
					Description: `(Optional) Start synchronizing X seconds after connection initiation . The values must be in a range between 2 and 3600.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Use State Synchronization. ` + "`" + `bypass_on_failure` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "override_profile",
					Description: `(Optional) Override profile of global configuration.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Override value.<br><font color="red">Required only for</font> 'override-profile' is True. ` + "`" + `site_categorization_allow_mode` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "override_profile",
					Description: `(Optional) Override profile of global configuration.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Override value.<br><font color="red">Required only for</font> 'override-profile' is True. ` + "`" + `deny_untrusted_server_cert` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "override_profile",
					Description: `(Optional) Override profile of global configuration.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Override value.<br><font color="red">Required only for</font> 'override-profile' is True. ` + "`" + `deny_revoked_server_cert` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "override_profile",
					Description: `(Optional) Override profile of global configuration.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Override value.<br><font color="red">Required only for</font> 'override-profile' is True. ` + "`" + `deny_expired_server_cert` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "override_profile",
					Description: `(Optional) Override profile of global configuration.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Override value.<br><font color="red">Required only for</font> 'override-profile' is True. ` + "`" + `browser_based_authentication_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "authentication_settings",
					Description: `(Optional) Authentication Settings for Browser Based Authentication.authentication_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "browser_based_authentication_portal_settings",
					Description: `(Optional) Browser Based Authentication portal settings.browser_based_authentication_portal_settings blocks are documented below. ` + "`" + `identity_agent_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "agents_interval_keepalive",
					Description: `(Optional) Agents send keepalive period (minutes).`,
				},
				resource.Attribute{
					Name:        "user_reauthenticate_interval",
					Description: `(Optional) Agent reauthenticate time interval (minutes).`,
				},
				resource.Attribute{
					Name:        "authentication_settings",
					Description: `(Optional) Authentication Settings for Identity Agent.authentication_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "identity_agent_portal_settings",
					Description: `(Optional) Identity Agent accessibility settings.identity_agent_portal_settings blocks are documented below. ` + "`" + `identity_collector_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "authorized_clients",
					Description: `(Optional) Authorized Clients.authorized_clients blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "authentication_settings",
					Description: `(Optional) Authentication Settings for Identity Collector.authentication_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "client_access_permissions",
					Description: `(Optional) Identity Collector accessibility settings.client_access_permissions blocks are documented below. ` + "`" + `identity_sharing_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "share_with_other_gateways",
					Description: `(Optional) Enable identity sharing with other gateways.`,
				},
				resource.Attribute{
					Name:        "receive_from_other_gateways",
					Description: `(Optional) Enable receiving identity from other gateways.`,
				},
				resource.Attribute{
					Name:        "receive_from",
					Description: `(Optional) Gateway(s) to receive identity from.receive_from blocks are documented below. ` + "`" + `proxy_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "detect_using_x_forward_for",
					Description: `(Optional) Whether to use X-Forward-For HTTP header, which is added by the proxy server to keep track of the original source IP. ` + "`" + `anti_spoofing_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) If packets will be rejected (the Prevent option) or whether the packets will be monitored (the Detect option).`,
				},
				resource.Attribute{
					Name:        "exclude_packets",
					Description: `(Optional) Don't check packets from excluded network.`,
				},
				resource.Attribute{
					Name:        "excluded_network_name",
					Description: `(Optional) Excluded network name.`,
				},
				resource.Attribute{
					Name:        "excluded_network_uid",
					Description: `(Optional) Excluded network UID.`,
				},
				resource.Attribute{
					Name:        "spoof_tracking",
					Description: `(Optional) Spoof tracking. ` + "`" + `security_zone_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "auto_calculated",
					Description: `(Optional) Security Zone is calculated according to where the interface leads to.`,
				},
				resource.Attribute{
					Name:        "specific_zone",
					Description: `(Optional) Security Zone specified manually. ` + "`" + `topology_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "interface_leads_to_dmz",
					Description: `(Optional) Whether this interface leads to demilitarized zone (perimeter network).`,
				},
				resource.Attribute{
					Name:        "specific_network",
					Description: `(Optional) Network behind this interface. ` + "`" + `interfaces` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name.`,
				},
				resource.Attribute{
					Name:        "anti_spoofing",
					Description: `(Optional) N/A`,
				},
				resource.Attribute{
					Name:        "anti_spoofing_settings",
					Description: `(Optional) N/Aanti_spoofing_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `(Optional) IPv4 address.`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `(Optional) IPv6 address.`,
				},
				resource.Attribute{
					Name:        "network_mask",
					Description: `(Optional) IPv4 or IPv6 network mask. If both masks are required use ipv4-network-mask and ipv6-network-mask fields explicitly. Instead of providing mask itself it is possible to specify IPv4 or IPv6 mask length in mask-length field. If both masks length are required use ipv4-mask-length and ipv6-mask-length fields explicitly.`,
				},
				resource.Attribute{
					Name:        "ipv4_network_mask",
					Description: `(Optional) IPv4 network address.`,
				},
				resource.Attribute{
					Name:        "ipv6_network_mask",
					Description: `(Optional) IPv6 network address.`,
				},
				resource.Attribute{
					Name:        "ipv4_mask_length",
					Description: `(Optional) IPv4 network mask length.`,
				},
				resource.Attribute{
					Name:        "ipv6_mask_length",
					Description: `(Optional) IPv6 network mask length.`,
				},
				resource.Attribute{
					Name:        "security_zone",
					Description: `(Optional) N/A`,
				},
				resource.Attribute{
					Name:        "security_zone_settings",
					Description: `(Optional) N/Asecurity_zone_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "topology",
					Description: `(Optional) N/A`,
				},
				resource.Attribute{
					Name:        "topology_settings",
					Description: `(Optional) N/Atopology_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. ` + "`" + `portal_web_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "aliases",
					Description: `(Optional) List of URL aliases that are redirected to the main portal URL.aliases blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "main_url",
					Description: `(Optional) The main URL for the web portal. ` + "`" + `certificate_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "base64_certificate",
					Description: `(Optional) The certificate file encoded in Base64 with padding. This file must be in the`,
				},
				resource.Attribute{
					Name:        "base64_password",
					Description: `(Optional) Password (encoded in Base64 with padding) for the certificate file. ` + "`" + `accessibility` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "allow_access_from",
					Description: `(Optional) Allowed access to the web portal (based on interfaces, or security policy).`,
				},
				resource.Attribute{
					Name:        "internal_access_settings",
					Description: `(Optional) Configuration of the additional portal access settings for internal interfaces only.internal_access_settings blocks are documented below. ` + "`" + `portal_web_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "aliases",
					Description: `(Optional) List of URL aliases that are redirected to the main portal URL.aliases blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "main_url",
					Description: `(Optional) The main URL for the web portal. ` + "`" + `certificate_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "base64_certificate",
					Description: `(Optional) The certificate file encoded in Base64 with padding. This file must be in the`,
				},
				resource.Attribute{
					Name:        "base64_password",
					Description: `(Optional) Password (encoded in Base64 with padding) for the certificate file. ` + "`" + `accessibility` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "allow_access_from",
					Description: `(Optional) Allowed access to the web portal (based on interfaces, or security policy).`,
				},
				resource.Attribute{
					Name:        "internal_access_settings",
					Description: `(Optional) Configuration of the additional portal access settings for internal interfaces only.internal_access_settings blocks are documented below. ` + "`" + `authentication` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "authentication_clients",
					Description: `(Optional) Collection of VPN Authentication clients identified by the name or UID.authentication_clients blocks are documented below. ` + "`" + `link_selection` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "dns_resolving_hostname",
					Description: `(Optional) DNS Resolving Hostname. Must be set when "ip-selection" was selected to be "dns-resolving-from-hostname". ` + "`" + `office_mode` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) Office Mode Permissions. When selected to be "off", all the other definitions are irrelevant.`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `(Optional) Group. Identified by name or UID. Must be set when "office-mode-permissions" was selected to be "group".`,
				},
				resource.Attribute{
					Name:        "allocate_ip_address_from",
					Description: `(Optional) Allocate IP address Method. Allocate IP address by sequentially trying the given methods until success.allocate_ip_address_from blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "support_multiple_interfaces",
					Description: `(Optional) Support connectivity enhancement for gateways with multiple external interfaces.`,
				},
				resource.Attribute{
					Name:        "perform_anti_spoofing",
					Description: `(Optional) Perform Anti-Spoofing on Office Mode addresses.`,
				},
				resource.Attribute{
					Name:        "anti_spoofing_additional_addresses",
					Description: `(Optional) Additional IP Addresses for Anti-Spoofing. Identified by name or UID. Must be set when "perform-anti-spoofings" is true. ` + "`" + `remote_access` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "support_l2tp",
					Description: `(Optional) Support L2TP (relevant only when office mode is active).`,
				},
				resource.Attribute{
					Name:        "l2tp_auth_method",
					Description: `(Optional) L2TP Authentication Method. Must be set when "support-l2tp" is true.`,
				},
				resource.Attribute{
					Name:        "l2tp_certificate",
					Description: `(Optional) L2TP Certificate. Must be set when "l2tp-auth-method" was selected to be "certificate". Insert "defaultCert" when you want to use the default certificate.`,
				},
				resource.Attribute{
					Name:        "allow_vpn_clients_to_route_traffic",
					Description: `(Optional) Allow VPN clients to route traffic.`,
				},
				resource.Attribute{
					Name:        "support_nat_traversal_mechanism",
					Description: `(Optional) Support NAT traversal mechanism (UDP encapsulation).`,
				},
				resource.Attribute{
					Name:        "nat_traversal_service",
					Description: `(Optional) Allocated NAT traversal UDP service. Identified by name or UID. Must be set when "support-nat-traversal-mechanism" is true.`,
				},
				resource.Attribute{
					Name:        "support_visitor_mode",
					Description: `(Optional) Support Visitor Mode.`,
				},
				resource.Attribute{
					Name:        "visitor_mode_service",
					Description: `(Optional) TCP Service for Visitor Mode. Identified by name or UID. Must be set when "support-visitor-mode" is true.`,
				},
				resource.Attribute{
					Name:        "visitor_mode_interface",
					Description: `(Optional) Interface for Visitor Mode. Must be set when "support-visitor-mode" is true. Insert IPV4 Address of existing interface or "All IPs" when you want all interfaces. ` + "`" + `use_early_versions` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Use early versions compatibility mode.`,
				},
				resource.Attribute{
					Name:        "compatibility_mode",
					Description: `(Optional) Early versions compatibility mode. ` + "`" + `purge_sam_file` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Purge SAM File.`,
				},
				resource.Attribute{
					Name:        "purge_when_size_reaches_to",
					Description: `(Optional) Purge SAM File When it Reaches to. ` + "`" + `website_categorization` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) Website categorization mode.`,
				},
				resource.Attribute{
					Name:        "custom_mode",
					Description: `(Optional) Custom mode object.custom_mode blocks are documented below. ` + "`" + `authentication_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "authentication_method",
					Description: `(Optional) Authentication method.`,
				},
				resource.Attribute{
					Name:        "identity_provider",
					Description: `(Optional) Identity provider object identified by the name or UID. Must be set when "authentication-method" was selected to be "identity provider".identity_provider blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "radius",
					Description: `(Optional) Radius server object identified by the name or UID. Must be set when "authentication-method" was selected to be "radius".`,
				},
				resource.Attribute{
					Name:        "users_directories",
					Description: `(Optional) Users directories.users_directories blocks are documented below. ` + "`" + `browser_based_authentication_portal_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "portal_web_settings",
					Description: `(Optional) Configuration of the portal web settings.portal_web_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "certificate_settings",
					Description: `(Optional) Configuration of the portal certificate settings.certificate_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "accessibility",
					Description: `(Optional) Configuration of the portal access settings.accessibility blocks are documented below. ` + "`" + `authentication_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "authentication_method",
					Description: `(Optional) Authentication method.`,
				},
				resource.Attribute{
					Name:        "radius",
					Description: `(Optional) Radius server object identified by the name or UID. Must be set when "authentication-method" was selected to be "radius".`,
				},
				resource.Attribute{
					Name:        "users_directories",
					Description: `(Optional) Users directories.users_directories blocks are documented below. ` + "`" + `identity_agent_portal_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "accessibility",
					Description: `(Optional) Configuration of the portal access settings.accessibility blocks are documented below. ` + "`" + `authorized_clients` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "client",
					Description: `(Optional) Host / Network Group Name or UID.`,
				},
				resource.Attribute{
					Name:        "client_secret",
					Description: `(Optional) Client Secret. ` + "`" + `authentication_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "users_directories",
					Description: `(Optional) Users directories.users_directories blocks are documented below. ` + "`" + `client_access_permissions` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "accessibility",
					Description: `(Optional) Configuration of the portal access settings.accessibility blocks are documented below. ` + "`" + `anti_spoofing_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) If packets will be rejected (the Prevent option) or whether the packets will be monitored (the Detect option).`,
				},
				resource.Attribute{
					Name:        "exclude_packets",
					Description: `(Optional) Don't check packets from excluded network.`,
				},
				resource.Attribute{
					Name:        "excluded_network_name",
					Description: `(Optional) Excluded network name.`,
				},
				resource.Attribute{
					Name:        "excluded_network_uid",
					Description: `(Optional) Excluded network UID.`,
				},
				resource.Attribute{
					Name:        "spoof_tracking",
					Description: `(Optional) Spoof tracking. ` + "`" + `security_zone_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "auto_calculated",
					Description: `(Optional) Security Zone is calculated according to where the interface leads to.`,
				},
				resource.Attribute{
					Name:        "specific_zone",
					Description: `(Optional) Security Zone specified manually. ` + "`" + `topology_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "interface_leads_to_dmz",
					Description: `(Optional) Whether this interface leads to demilitarized zone (perimeter network).`,
				},
				resource.Attribute{
					Name:        "specific_network",
					Description: `(Optional) Network behind this interface. ` + "`" + `internal_access_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "undefined",
					Description: `(Optional) Controls portal access settings for internal interfaces, whose topology is set to 'Undefined'.`,
				},
				resource.Attribute{
					Name:        "dmz",
					Description: `(Optional) Controls portal access settings for internal interfaces, whose topology is set to 'DMZ'.`,
				},
				resource.Attribute{
					Name:        "vpn",
					Description: `(Optional) Controls portal access settings for interfaces that are part of a VPN Encryption Domain. ` + "`" + `internal_access_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "undefined",
					Description: `(Optional) Controls portal access settings for internal interfaces, whose topology is set to 'Undefined'.`,
				},
				resource.Attribute{
					Name:        "dmz",
					Description: `(Optional) Controls portal access settings for internal interfaces, whose topology is set to 'DMZ'.`,
				},
				resource.Attribute{
					Name:        "vpn",
					Description: `(Optional) Controls portal access settings for interfaces that are part of a VPN Encryption Domain. ` + "`" + `allocate_ip_address_from` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "radius_server",
					Description: `(Optional) Radius server used to authenticate the user.`,
				},
				resource.Attribute{
					Name:        "use_allocate_method",
					Description: `(Optional) Use Allocate Method.`,
				},
				resource.Attribute{
					Name:        "allocate_method",
					Description: `(Optional) Using either Manual (IP Pool) or Automatic (DHCP). Must be set when "use-allocate-method" is true.`,
				},
				resource.Attribute{
					Name:        "manual_network",
					Description: `(Optional) Manual Network. Identified by name or UID. Must be set when "allocate-method" was selected to be "manual".`,
				},
				resource.Attribute{
					Name:        "dhcp_server",
					Description: `(Optional) DHCP Server. Identified by name or UID. Must be set when "allocate-method" was selected to be "automatic".`,
				},
				resource.Attribute{
					Name:        "virtual_ip_address",
					Description: `(Optional) Virtual IPV4 address for DHCP server replies. Must be set when "allocate-method" was selected to be "automatic".`,
				},
				resource.Attribute{
					Name:        "dhcp_mac_address",
					Description: `(Optional) Calculated MAC address for DHCP allocation. Must be set when "allocate-method" was selected to be "automatic".`,
				},
				resource.Attribute{
					Name:        "optional_parameters",
					Description: `(Optional) This configuration applies to all Office Mode methods except Automatic (using DHCP) and ipassignment.conf entries which contain this data.optional_parameters blocks are documented below. ` + "`" + `custom_mode` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "social_networking_widgets",
					Description: `(Optional) Social networking widgets mode.`,
				},
				resource.Attribute{
					Name:        "url_filtering",
					Description: `(Optional) URL filtering mode. ` + "`" + `users_directories` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "external_user_profile",
					Description: `(Optional) External user profile.`,
				},
				resource.Attribute{
					Name:        "internal_users",
					Description: `(Optional) Internal users.`,
				},
				resource.Attribute{
					Name:        "users_from_external_directories",
					Description: `(Optional) Users from external directories.`,
				},
				resource.Attribute{
					Name:        "specific",
					Description: `(Optional) LDAP AU objects identified by the name or UID. Must be set when "users-from-external-directories" was selected to be "specific".specific blocks are documented below. ` + "`" + `portal_web_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "aliases",
					Description: `(Optional) List of URL aliases that are redirected to the main portal URL.aliases blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "main_url",
					Description: `(Optional) The main URL for the web portal. ` + "`" + `certificate_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "base64_certificate",
					Description: `(Optional) The certificate file encoded in Base64 with padding. This file must be in the`,
				},
				resource.Attribute{
					Name:        "base64_password",
					Description: `(Optional) Password (encoded in Base64 with padding) for the certificate file. ` + "`" + `accessibility` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "allow_access_from",
					Description: `(Optional) Allowed access to the web portal (based on interfaces, or security policy).`,
				},
				resource.Attribute{
					Name:        "internal_access_settings",
					Description: `(Optional) Configuration of the additional portal access settings for internal interfaces only.internal_access_settings blocks are documented below. ` + "`" + `users_directories` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "external_user_profile",
					Description: `(Optional) External user profile.`,
				},
				resource.Attribute{
					Name:        "internal_users",
					Description: `(Optional) Internal users.`,
				},
				resource.Attribute{
					Name:        "users_from_external_directories",
					Description: `(Optional) Users from external directories.`,
				},
				resource.Attribute{
					Name:        "specific",
					Description: `(Optional) LDAP AU objects identified by the name or UID. Must be set when "users-from-external-directories" was selected to be "specific".specific blocks are documented below. ` + "`" + `accessibility` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "allow_access_from",
					Description: `(Optional) Allowed access to the web portal (based on interfaces, or security policy).`,
				},
				resource.Attribute{
					Name:        "internal_access_settings",
					Description: `(Optional) Configuration of the additional portal access settings for internal interfaces only.internal_access_settings blocks are documented below. ` + "`" + `users_directories` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "external_user_profile",
					Description: `(Optional) External user profile.`,
				},
				resource.Attribute{
					Name:        "internal_users",
					Description: `(Optional) Internal users.`,
				},
				resource.Attribute{
					Name:        "users_from_external_directories",
					Description: `(Optional) Users from external directories.`,
				},
				resource.Attribute{
					Name:        "specific",
					Description: `(Optional) LDAP AU objects identified by the name or UID. Must be set when "users-from-external-directories" was selected to be "specific".specific blocks are documented below. ` + "`" + `accessibility` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "allow_access_from",
					Description: `(Optional) Allowed access to the web portal (based on interfaces, or security policy).`,
				},
				resource.Attribute{
					Name:        "internal_access_settings",
					Description: `(Optional) Configuration of the additional portal access settings for internal interfaces only.internal_access_settings blocks are documented below. ` + "`" + `optional_parameters` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "use_primary_dns_server",
					Description: `(Optional) Use Primary DNS Server.`,
				},
				resource.Attribute{
					Name:        "primary_dns_server",
					Description: `(Optional) Primary DNS Server. Identified by name or UID. Must be set when "use-primary-dns-server" is true and can not be set when "use-primary-dns-server" is false.`,
				},
				resource.Attribute{
					Name:        "use_first_backup_dns_server",
					Description: `(Optional) Use First Backup DNS Server.`,
				},
				resource.Attribute{
					Name:        "first_backup_dns_server",
					Description: `(Optional) First Backup DNS Server. Identified by name or UID. Must be set when "use-first-backup-dns-server" is true and can not be set when "use-first-backup-dns-server" is false.`,
				},
				resource.Attribute{
					Name:        "use_second_backup_dns_server",
					Description: `(Optional) Use Second Backup DNS Server.`,
				},
				resource.Attribute{
					Name:        "second_backup_dns_server",
					Description: `(Optional) Second Backup DNS Server. Identified by name or UID. Must be set when "use-second-backup-dns-server" is true and can not be set when "use-second-backup-dns-server" is false.`,
				},
				resource.Attribute{
					Name:        "dns_suffixes",
					Description: `(Optional) DNS Suffixes.`,
				},
				resource.Attribute{
					Name:        "use_primary_wins_server",
					Description: `(Optional) Use Primary WINS Server.`,
				},
				resource.Attribute{
					Name:        "primary_wins_server",
					Description: `(Optional) Primary WINS Server. Identified by name or UID. Must be set when "use-primary-wins-server" is true and can not be set when "use-primary-wins-server" is false.`,
				},
				resource.Attribute{
					Name:        "use_first_backup_wins_server",
					Description: `(Optional) Use First Backup WINS Server.`,
				},
				resource.Attribute{
					Name:        "first_backup_wins_server",
					Description: `(Optional) First Backup WINS Server. Identified by name or UID. Must be set when "use-first-backup-wins-server" is true and can not be set when "use-first-backup-wins-server" is false.`,
				},
				resource.Attribute{
					Name:        "use_second_backup_wins_server",
					Description: `(Optional) Use Second Backup WINS Server.`,
				},
				resource.Attribute{
					Name:        "second_backup_wins_server",
					Description: `(Optional) Second Backup WINS Server. Identified by name or UID. Must be set when "use-second-backup-wins-server" is true and can not be set when "use-second-backup-wins-server" is false. ` + "`" + `internal_access_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "undefined",
					Description: `(Optional) Controls portal access settings for internal interfaces, whose topology is set to 'Undefined'.`,
				},
				resource.Attribute{
					Name:        "dmz",
					Description: `(Optional) Controls portal access settings for internal interfaces, whose topology is set to 'DMZ'.`,
				},
				resource.Attribute{
					Name:        "vpn",
					Description: `(Optional) Controls portal access settings for interfaces that are part of a VPN Encryption Domain. ` + "`" + `internal_access_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "undefined",
					Description: `(Optional) Controls portal access settings for internal interfaces, whose topology is set to 'Undefined'.`,
				},
				resource.Attribute{
					Name:        "dmz",
					Description: `(Optional) Controls portal access settings for internal interfaces, whose topology is set to 'DMZ'.`,
				},
				resource.Attribute{
					Name:        "vpn",
					Description: `(Optional) Controls portal access settings for interfaces that are part of a VPN Encryption Domain. ` + "`" + `internal_access_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "undefined",
					Description: `(Optional) Controls portal access settings for internal interfaces, whose topology is set to 'Undefined'.`,
				},
				resource.Attribute{
					Name:        "dmz",
					Description: `(Optional) Controls portal access settings for internal interfaces, whose topology is set to 'DMZ'.`,
				},
				resource.Attribute{
					Name:        "vpn",
					Description: `(Optional) Controls portal access settings for interfaces that are part of a VPN Encryption Domain.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_simple_gateway",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Simple Gateway.`,
			Description:      ``,
			Keywords: []string{
				"management",
				"simple",
				"gateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "advanced_settings",
					Description: `(Optional) N/Aadvanced_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "anti_bot",
					Description: `(Optional) Anti-Bot blade enabled.`,
				},
				resource.Attribute{
					Name:        "anti_virus",
					Description: `(Optional) Anti-Virus blade enabled.`,
				},
				resource.Attribute{
					Name:        "application_control",
					Description: `(Optional) Application Control blade enabled.`,
				},
				resource.Attribute{
					Name:        "application_control_and_url_filtering_settings",
					Description: `(Optional) Gateway Application Control and URL filtering settings.application_control_and_url_filtering_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "content_awareness",
					Description: `(Optional) Content Awareness blade enabled.`,
				},
				resource.Attribute{
					Name:        "enable_https_inspection",
					Description: `(Optional) Enable HTTPS Inspection after defining an outbound inspection certificate. <br>To define the outbound certificate use outbound inspection certificate API.`,
				},
				resource.Attribute{
					Name:        "fetch_policy",
					Description: `(Optional) Security management server(s) to fetch the policy from.fetch_policy blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "firewall",
					Description: `(Optional) Firewall blade enabled.`,
				},
				resource.Attribute{
					Name:        "firewall_settings",
					Description: `(Optional) N/Afirewall_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "hit_count",
					Description: `(Optional) Hit count tracks the number of connections each rule matches.`,
				},
				resource.Attribute{
					Name:        "https_inspection",
					Description: `(Optional) HTTPS inspection.https_inspection blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "icap_server",
					Description: `(Optional) ICAP Server enabled.`,
				},
				resource.Attribute{
					Name:        "identity_awareness",
					Description: `(Optional) Identity awareness blade enabled.`,
				},
				resource.Attribute{
					Name:        "identity_awareness_settings",
					Description: `(Optional) Gateway Identity Awareness settings.identity_awareness_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `(Optional) Network interfaces.interfaces blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `(Optional) IPv4 address.`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `(Optional) IPv6 address.`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `(Optional) Intrusion Prevention System blade enabled.`,
				},
				resource.Attribute{
					Name:        "ips_update_policy",
					Description: `(Optional) Specifies whether the IPS will be downloaded from the Management or directly to the Gateway.`,
				},
				resource.Attribute{
					Name:        "nat_hide_internal_interfaces",
					Description: `(Optional) Hide internal networks behind the Gateway's external IP.`,
				},
				resource.Attribute{
					Name:        "nat_settings",
					Description: `(Optional) NAT settings.nat_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "one_time_password",
					Description: `(Optional) N/A`,
				},
				resource.Attribute{
					Name:        "os_name",
					Description: `(Optional) Gateway platform operating system.`,
				},
				resource.Attribute{
					Name:        "platform_portal_settings",
					Description: `(Optional) Platform portal settings.platform_portal_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "proxy_settings",
					Description: `(Optional) Proxy Server for Gateway.proxy_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "qos",
					Description: `(Optional) QoS.`,
				},
				resource.Attribute{
					Name:        "save_logs_locally",
					Description: `(Optional) Save logs locally on the gateway.`,
				},
				resource.Attribute{
					Name:        "send_alerts_to_server",
					Description: `(Optional) Server(s) to send alerts to.send_alerts_to_server blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "send_logs_to_backup_server",
					Description: `(Optional) Backup server(s) to send logs to.send_logs_to_backup_server blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "send_logs_to_server",
					Description: `(Optional) Server(s) to send logs to.send_logs_to_server blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "threat_emulation",
					Description: `(Optional) Threat Emulation blade enabled.`,
				},
				resource.Attribute{
					Name:        "threat_extraction",
					Description: `(Optional) Threat Extraction blade enabled.`,
				},
				resource.Attribute{
					Name:        "threat_prevention_mode",
					Description: `(Optional) The mode of Threat Prevention to use. When using Autonomous Threat Prevention, disabling the Threat Prevention blades is not allowed.`,
				},
				resource.Attribute{
					Name:        "url_filtering",
					Description: `(Optional) URL Filtering blade enabled.`,
				},
				resource.Attribute{
					Name:        "usercheck_portal_settings",
					Description: `(Optional) UserCheck portal settings.usercheck_portal_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) Gateway platform version.`,
				},
				resource.Attribute{
					Name:        "vpn",
					Description: `(Optional) VPN blade enabled.`,
				},
				resource.Attribute{
					Name:        "vpn_settings",
					Description: `(Optional) Gateway VPN settings.vpn_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "zero_phishing",
					Description: `(Optional) Zero Phishing blade enabled.`,
				},
				resource.Attribute{
					Name:        "zero_phishing_fqdn",
					Description: `(Optional) Zero Phishing gateway FQDN.`,
				},
				resource.Attribute{
					Name:        "logs_settings",
					Description: `(Optional) Logs settings that apply to Quantum Security Gateways that run Gaia OS.logs_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "show_portals_certificate",
					Description: `(Optional) Indicates whether to show the portals certificate value in the reply.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `(Optional) Collection of group identifiers.groups blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. ` + "`" + `advanced_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "connection_persistence",
					Description: `(Optional) Handling established connections when installing a new policy.`,
				},
				resource.Attribute{
					Name:        "sam",
					Description: `(Optional) SAM.sam blocks are documented below. ` + "`" + `application_control_and_url_filtering_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "global_settings_mode",
					Description: `(Optional) Whether to override global settings or not.`,
				},
				resource.Attribute{
					Name:        "override_global_settings",
					Description: `(Optional) override global settings object.override_global_settings blocks are documented below. ` + "`" + `firewall_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "auto_calculate_connections_hash_table_size_and_memory_pool",
					Description: `(Optional) N/A`,
				},
				resource.Attribute{
					Name:        "auto_maximum_limit_for_concurrent_connections",
					Description: `(Optional) N/A`,
				},
				resource.Attribute{
					Name:        "connections_hash_size",
					Description: `(Optional) N/A`,
				},
				resource.Attribute{
					Name:        "maximum_limit_for_concurrent_connections",
					Description: `(Optional) N/A`,
				},
				resource.Attribute{
					Name:        "maximum_memory_pool_size",
					Description: `(Optional) N/A`,
				},
				resource.Attribute{
					Name:        "memory_pool_size",
					Description: `(Optional) N/A ` + "`" + `https_inspection` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "bypass_on_failure",
					Description: `(Optional) Set to be true in order to bypass all requests (Fail-open) in case of internal system error.bypass_on_failure blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "site_categorization_allow_mode",
					Description: `(Optional) Set to 'background' in order to allowed requests until categorization is complete.site_categorization_allow_mode blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "deny_untrusted_server_cert",
					Description: `(Optional) Set to be true in order to drop traffic from servers with untrusted server certificate.deny_untrusted_server_cert blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "deny_revoked_server_cert",
					Description: `(Optional) Set to be true in order to drop traffic from servers with revoked server certificate (validate CRL).deny_revoked_server_cert blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "deny_expired_server_cert",
					Description: `(Optional) Set to be true in order to drop traffic from servers with expired server certificate.deny_expired_server_cert blocks are documented below. ` + "`" + `identity_awareness_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "browser_based_authentication",
					Description: `(Optional) Enable Browser Based Authentication source.`,
				},
				resource.Attribute{
					Name:        "browser_based_authentication_settings",
					Description: `(Optional) Browser Based Authentication settings.browser_based_authentication_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "identity_agent",
					Description: `(Optional) Enable Identity Agent source.`,
				},
				resource.Attribute{
					Name:        "identity_agent_settings",
					Description: `(Optional) Identity Agent settings.identity_agent_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "identity_collector",
					Description: `(Optional) Enable Identity Collector source.`,
				},
				resource.Attribute{
					Name:        "identity_collector_settings",
					Description: `(Optional) Identity Collector settings.identity_collector_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "identity_sharing_settings",
					Description: `(Optional) Identity sharing settings.identity_sharing_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "proxy_settings",
					Description: `(Optional) Identity-Awareness Proxy settings.proxy_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "remote_access",
					Description: `(Optional) Enable Remote Access Identity source. ` + "`" + `interfaces` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name. Must be unique in the domain.`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `(Optional) IPv4 address.`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `(Optional) IPv6 address.`,
				},
				resource.Attribute{
					Name:        "network_mask",
					Description: `(Optional) IPv4 or IPv6 network mask. If both masks are required use ipv4-network-mask and ipv6-network-mask fields explicitly. Instead of providing mask itself it is possible to specify IPv4 or IPv6 mask length in mask-length field. If both masks length are required use ipv4-mask-length and ipv6-mask-length fields explicitly.`,
				},
				resource.Attribute{
					Name:        "ipv4_network_mask",
					Description: `(Optional) IPv4 network address.`,
				},
				resource.Attribute{
					Name:        "ipv6_network_mask",
					Description: `(Optional) IPv6 network address.`,
				},
				resource.Attribute{
					Name:        "ipv4_mask_length",
					Description: `(Optional) IPv4 network mask length.`,
				},
				resource.Attribute{
					Name:        "ipv6_mask_length",
					Description: `(Optional) IPv6 network mask length.`,
				},
				resource.Attribute{
					Name:        "anti_spoofing",
					Description: `(Optional) N/A`,
				},
				resource.Attribute{
					Name:        "anti_spoofing_settings",
					Description: `(Optional) N/Aanti_spoofing_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "security_zone",
					Description: `(Optional) N/A`,
				},
				resource.Attribute{
					Name:        "security_zone_settings",
					Description: `(Optional) N/Asecurity_zone_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "topology",
					Description: `(Optional) N/A`,
				},
				resource.Attribute{
					Name:        "topology_settings",
					Description: `(Optional) N/Atopology_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. ` + "`" + `nat_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "auto_rule",
					Description: `(Optional) Whether to add automatic address translation rules.`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `(Optional) IPv4 address.`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `(Optional) IPv6 address.`,
				},
				resource.Attribute{
					Name:        "hide_behind",
					Description: `(Optional) Hide behind method. This parameter is forbidden in case "method" parameter is "static".`,
				},
				resource.Attribute{
					Name:        "install_on",
					Description: `(Optional) Which gateway should apply the NAT translation.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Optional) NAT translation method. ` + "`" + `platform_portal_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "portal_web_settings",
					Description: `(Optional) Configuration of the portal web settings.portal_web_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "certificate_settings",
					Description: `(Optional) Configuration of the portal certificate settings.certificate_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "accessibility",
					Description: `(Optional) Configuration of the portal access settings.accessibility blocks are documented below. ` + "`" + `proxy_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "use_custom_proxy",
					Description: `(Optional) Use custom proxy settings for this network object.`,
				},
				resource.Attribute{
					Name:        "proxy_server",
					Description: `(Optional) N/A`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) N/A ` + "`" + `usercheck_portal_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) State of the web portal (enabled or disabled). The supported blades are: {'Application Control', 'URL Filtering', 'Data Loss Prevention', 'Anti Virus', 'Anti Bot', 'Threat Emulation', 'Threat Extraction', 'Data Awareness'}.`,
				},
				resource.Attribute{
					Name:        "portal_web_settings",
					Description: `(Optional) Configuration of the portal web settings.portal_web_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "certificate_settings",
					Description: `(Optional) Configuration of the portal certificate settings.certificate_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "accessibility",
					Description: `(Optional) Configuration of the portal access settings.accessibility blocks are documented below. ` + "`" + `vpn_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "authentication",
					Description: `(Optional) Authentication.authentication blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "link_selection",
					Description: `(Optional) Link Selection.link_selection blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "maximum_concurrent_ike_negotiations",
					Description: `(Optional) N/A`,
				},
				resource.Attribute{
					Name:        "maximum_concurrent_tunnels",
					Description: `(Optional) N/A`,
				},
				resource.Attribute{
					Name:        "office_mode",
					Description: `(Optional) Office Mode. Notation Wide Impact - Office Mode apply IPSec VPN Software Blade clients and to the Mobile Access Software Blade clients.office_mode blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "remote_access",
					Description: `(Optional) Remote Access.remote_access blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "vpn_domain",
					Description: `(Optional) Gateway VPN domain identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "vpn_domain_exclude_external_ip_addresses",
					Description: `(Optional) Exclude the external IP addresses from the VPN domain of this Security Gateway.`,
				},
				resource.Attribute{
					Name:        "vpn_domain_type",
					Description: `(Optional) Gateway VPN domain type. ` + "`" + `logs_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "alert_when_free_disk_space_below",
					Description: `(Optional) Enable alert when free disk space is below threshold.`,
				},
				resource.Attribute{
					Name:        "alert_when_free_disk_space_below_threshold",
					Description: `(Optional) Alert when free disk space below threshold.`,
				},
				resource.Attribute{
					Name:        "alert_when_free_disk_space_below_type",
					Description: `(Optional) Alert when free disk space below type.`,
				},
				resource.Attribute{
					Name:        "before_delete_keep_logs_from_the_last_days",
					Description: `(Optional) Enable before delete keep logs from the last days.`,
				},
				resource.Attribute{
					Name:        "before_delete_keep_logs_from_the_last_days_threshold",
					Description: `(Optional) Before delete keep logs from the last days threshold.`,
				},
				resource.Attribute{
					Name:        "before_delete_run_script",
					Description: `(Optional) Enable Before delete run script.`,
				},
				resource.Attribute{
					Name:        "before_delete_run_script_command",
					Description: `(Optional) Before delete run script command.`,
				},
				resource.Attribute{
					Name:        "delete_index_files_older_than_days",
					Description: `(Optional) Enable delete index files older than days.`,
				},
				resource.Attribute{
					Name:        "delete_index_files_older_than_days_threshold",
					Description: `(Optional) Delete index files older than days threshold.`,
				},
				resource.Attribute{
					Name:        "delete_index_files_when_index_size_above",
					Description: `(Optional) Enable delete index files when index size above.`,
				},
				resource.Attribute{
					Name:        "delete_index_files_when_index_size_above_threshold",
					Description: `(Optional) Delete index files when index size above threshold.`,
				},
				resource.Attribute{
					Name:        "delete_when_free_disk_space_below",
					Description: `(Optional) Enable delete when free disk space below.`,
				},
				resource.Attribute{
					Name:        "delete_when_free_disk_space_below_threshold",
					Description: `(Optional) Delete when free disk space below threshold.`,
				},
				resource.Attribute{
					Name:        "detect_new_citrix_ica_application_names",
					Description: `(Optional) Enable detect new Citrix ICA application names.`,
				},
				resource.Attribute{
					Name:        "distribute_logs_between_all_active_servers",
					Description: `(Optional) Distribute logs between all active servers.`,
				},
				resource.Attribute{
					Name:        "forward_logs_to_log_server",
					Description: `(Optional) Enable forward logs to log server.`,
				},
				resource.Attribute{
					Name:        "forward_logs_to_log_server_name",
					Description: `(Optional) Forward logs to log server name.`,
				},
				resource.Attribute{
					Name:        "forward_logs_to_log_server_schedule_name",
					Description: `(Optional) Forward logs to log server schedule name.`,
				},
				resource.Attribute{
					Name:        "free_disk_space_metrics",
					Description: `(Optional) Free disk space metrics.`,
				},
				resource.Attribute{
					Name:        "perform_log_rotate_before_log_forwarding",
					Description: `(Optional) Enable perform log rotate before log forwarding.`,
				},
				resource.Attribute{
					Name:        "reject_connections_when_free_disk_space_below_threshold",
					Description: `(Optional) Enable reject connections when free disk space below threshold.`,
				},
				resource.Attribute{
					Name:        "reserve_for_packet_capture_metrics",
					Description: `(Optional) Reserve for packet capture metrics.`,
				},
				resource.Attribute{
					Name:        "reserve_for_packet_capture_threshold",
					Description: `(Optional) Reserve for packet capture threshold.`,
				},
				resource.Attribute{
					Name:        "rotate_log_by_file_size",
					Description: `(Optional) Enable rotate log by file size.`,
				},
				resource.Attribute{
					Name:        "rotate_log_file_size_threshold",
					Description: `(Optional) Log file size threshold.`,
				},
				resource.Attribute{
					Name:        "rotate_log_on_schedule",
					Description: `(Optional) Enable rotate log on schedule.`,
				},
				resource.Attribute{
					Name:        "rotate_log_schedule_name",
					Description: `(Optional) Rotate log schedule name.`,
				},
				resource.Attribute{
					Name:        "stop_logging_when_free_disk_space_below",
					Description: `(Optional) Enable stop logging when free disk space below.`,
				},
				resource.Attribute{
					Name:        "stop_logging_when_free_disk_space_below_threshold",
					Description: `(Optional) Stop logging when free disk space below threshold.`,
				},
				resource.Attribute{
					Name:        "turn_on_qos_logging",
					Description: `(Optional) Enable turn on QoS Logging.`,
				},
				resource.Attribute{
					Name:        "update_account_log_every",
					Description: `(Optional) Update account log in every amount of seconds. ` + "`" + `sam` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "forward_to_other_sam_servers",
					Description: `(Optional) Forward SAM clients' requests to other SAM servers.`,
				},
				resource.Attribute{
					Name:        "use_early_versions",
					Description: `(Optional) Use early versions compatibility mode.use_early_versions blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "purge_sam_file",
					Description: `(Optional) Purge SAM File.purge_sam_file blocks are documented below. ` + "`" + `override_global_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "fail_mode",
					Description: `(Optional) Fail mode - allow or block all requests.`,
				},
				resource.Attribute{
					Name:        "website_categorization",
					Description: `(Optional) Website categorization object.website_categorization blocks are documented below. ` + "`" + `bypass_on_failure` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "override_profile",
					Description: `(Optional) Override profile of global configuration.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Override value.<br><font color="red">Required only for</font> 'override-profile' is True. ` + "`" + `site_categorization_allow_mode` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "override_profile",
					Description: `(Optional) Override profile of global configuration.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Override value.<br><font color="red">Required only for</font> 'override-profile' is True. ` + "`" + `deny_untrusted_server_cert` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "override_profile",
					Description: `(Optional) Override profile of global configuration.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Override value.<br><font color="red">Required only for</font> 'override-profile' is True. ` + "`" + `deny_revoked_server_cert` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "override_profile",
					Description: `(Optional) Override profile of global configuration.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Override value.<br><font color="red">Required only for</font> 'override-profile' is True. ` + "`" + `deny_expired_server_cert` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "override_profile",
					Description: `(Optional) Override profile of global configuration.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Override value.<br><font color="red">Required only for</font> 'override-profile' is True. ` + "`" + `browser_based_authentication_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "authentication_settings",
					Description: `(Optional) Authentication Settings for Browser Based Authentication.authentication_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "browser_based_authentication_portal_settings",
					Description: `(Optional) Browser Based Authentication portal settings.browser_based_authentication_portal_settings blocks are documented below. ` + "`" + `identity_agent_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "agents_interval_keepalive",
					Description: `(Optional) Agents send keepalive period (minutes).`,
				},
				resource.Attribute{
					Name:        "user_reauthenticate_interval",
					Description: `(Optional) Agent reauthenticate time interval (minutes).`,
				},
				resource.Attribute{
					Name:        "authentication_settings",
					Description: `(Optional) Authentication Settings for Identity Agent.authentication_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "identity_agent_portal_settings",
					Description: `(Optional) Identity Agent accessibility settings.identity_agent_portal_settings blocks are documented below. ` + "`" + `identity_collector_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "authorized_clients",
					Description: `(Optional) Authorized Clients.authorized_clients blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "authentication_settings",
					Description: `(Optional) Authentication Settings for Identity Collector.authentication_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "client_access_permissions",
					Description: `(Optional) Identity Collector accessibility settings.client_access_permissions blocks are documented below. ` + "`" + `identity_sharing_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "share_with_other_gateways",
					Description: `(Optional) Enable identity sharing with other gateways.`,
				},
				resource.Attribute{
					Name:        "receive_from_other_gateways",
					Description: `(Optional) Enable receiving identity from other gateways.`,
				},
				resource.Attribute{
					Name:        "receive_from",
					Description: `(Optional) Gateway(s) to receive identity from.receive_from blocks are documented below. ` + "`" + `proxy_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "detect_using_x_forward_for",
					Description: `(Optional) Whether to use X-Forward-For HTTP header, which is added by the proxy server to keep track of the original source IP. ` + "`" + `anti_spoofing_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) If packets will be rejected (the Prevent option) or whether the packets will be monitored (the Detect option).`,
				},
				resource.Attribute{
					Name:        "exclude_packets",
					Description: `(Optional) Don't check packets from excluded network.`,
				},
				resource.Attribute{
					Name:        "excluded_network_name",
					Description: `(Optional) Excluded network name.`,
				},
				resource.Attribute{
					Name:        "excluded_network_uid",
					Description: `(Optional) Excluded network UID.`,
				},
				resource.Attribute{
					Name:        "spoof_tracking",
					Description: `(Optional) Spoof tracking. ` + "`" + `security_zone_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "auto_calculated",
					Description: `(Optional) Security Zone is calculated according to where the interface leads to.`,
				},
				resource.Attribute{
					Name:        "specific_zone",
					Description: `(Optional) Security Zone specified manually. ` + "`" + `topology_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "interface_leads_to_dmz",
					Description: `(Optional) Whether this interface leads to demilitarized zone (perimeter network).`,
				},
				resource.Attribute{
					Name:        "specific_network",
					Description: `(Optional) Network behind this interface. ` + "`" + `portal_web_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "aliases",
					Description: `(Optional) List of URL aliases that are redirected to the main portal URL.aliases blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "main_url",
					Description: `(Optional) The main URL for the web portal. ` + "`" + `certificate_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "base64_certificate",
					Description: `(Optional) The certificate file encoded in Base64 with padding. This file must be in the`,
				},
				resource.Attribute{
					Name:        "base64_password",
					Description: `(Optional) Password (encoded in Base64 with padding) for the certificate file. ` + "`" + `accessibility` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "allow_access_from",
					Description: `(Optional) Allowed access to the web portal (based on interfaces, or security policy).`,
				},
				resource.Attribute{
					Name:        "internal_access_settings",
					Description: `(Optional) Configuration of the additional portal access settings for internal interfaces only.internal_access_settings blocks are documented below. ` + "`" + `portal_web_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "aliases",
					Description: `(Optional) List of URL aliases that are redirected to the main portal URL.aliases blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "main_url",
					Description: `(Optional) The main URL for the web portal. ` + "`" + `certificate_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "base64_certificate",
					Description: `(Optional) The certificate file encoded in Base64 with padding. This file must be in the`,
				},
				resource.Attribute{
					Name:        "base64_password",
					Description: `(Optional) Password (encoded in Base64 with padding) for the certificate file. ` + "`" + `accessibility` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "allow_access_from",
					Description: `(Optional) Allowed access to the web portal (based on interfaces, or security policy).`,
				},
				resource.Attribute{
					Name:        "internal_access_settings",
					Description: `(Optional) Configuration of the additional portal access settings for internal interfaces only.internal_access_settings blocks are documented below. ` + "`" + `authentication` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "authentication_clients",
					Description: `(Optional) Collection of VPN Authentication clients identified by the name or UID.authentication_clients blocks are documented below. ` + "`" + `link_selection` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "dns_resolving_hostname",
					Description: `(Optional) DNS Resolving Hostname. Must be set when "ip-selection" was selected to be "dns-resolving-from-hostname". ` + "`" + `office_mode` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) Office Mode Permissions. When selected to be "off", all the other definitions are irrelevant.`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `(Optional) Group. Identified by name or UID. Must be set when "office-mode-permissions" was selected to be "group".`,
				},
				resource.Attribute{
					Name:        "allocate_ip_address_from",
					Description: `(Optional) Allocate IP address Method. Allocate IP address by sequentially trying the given methods until success.allocate_ip_address_from blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "support_multiple_interfaces",
					Description: `(Optional) Support connectivity enhancement for gateways with multiple external interfaces.`,
				},
				resource.Attribute{
					Name:        "perform_anti_spoofing",
					Description: `(Optional) Perform Anti-Spoofing on Office Mode addresses.`,
				},
				resource.Attribute{
					Name:        "anti_spoofing_additional_addresses",
					Description: `(Optional) Additional IP Addresses for Anti-Spoofing. Identified by name or UID. Must be set when "perform-anti-spoofings" is true. ` + "`" + `remote_access` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "support_l2tp",
					Description: `(Optional) Support L2TP (relevant only when office mode is active).`,
				},
				resource.Attribute{
					Name:        "l2tp_auth_method",
					Description: `(Optional) L2TP Authentication Method. Must be set when "support-l2tp" is true.`,
				},
				resource.Attribute{
					Name:        "l2tp_certificate",
					Description: `(Optional) L2TP Certificate. Must be set when "l2tp-auth-method" was selected to be "certificate". Insert "defaultCert" when you want to use the default certificate.`,
				},
				resource.Attribute{
					Name:        "allow_vpn_clients_to_route_traffic",
					Description: `(Optional) Allow VPN clients to route traffic.`,
				},
				resource.Attribute{
					Name:        "support_nat_traversal_mechanism",
					Description: `(Optional) Support NAT traversal mechanism (UDP encapsulation).`,
				},
				resource.Attribute{
					Name:        "nat_traversal_service",
					Description: `(Optional) Allocated NAT traversal UDP service. Identified by name or UID. Must be set when "support-nat-traversal-mechanism" is true.`,
				},
				resource.Attribute{
					Name:        "support_visitor_mode",
					Description: `(Optional) Support Visitor Mode.`,
				},
				resource.Attribute{
					Name:        "visitor_mode_service",
					Description: `(Optional) TCP Service for Visitor Mode. Identified by name or UID. Must be set when "support-visitor-mode" is true.`,
				},
				resource.Attribute{
					Name:        "visitor_mode_interface",
					Description: `(Optional) Interface for Visitor Mode. Must be set when "support-visitor-mode" is true. Insert IPV4 Address of existing interface or "All IPs" when you want all interfaces. ` + "`" + `use_early_versions` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Use early versions compatibility mode.`,
				},
				resource.Attribute{
					Name:        "compatibility_mode",
					Description: `(Optional) Early versions compatibility mode. ` + "`" + `purge_sam_file` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Purge SAM File.`,
				},
				resource.Attribute{
					Name:        "purge_when_size_reaches_to",
					Description: `(Optional) Purge SAM File When it Reaches to. ` + "`" + `website_categorization` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) Website categorization mode.`,
				},
				resource.Attribute{
					Name:        "custom_mode",
					Description: `(Optional) Custom mode object.custom_mode blocks are documented below. ` + "`" + `authentication_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "authentication_method",
					Description: `(Optional) Authentication method.`,
				},
				resource.Attribute{
					Name:        "identity_provider",
					Description: `(Optional) Identity provider object identified by the name or UID. Must be set when "authentication-method" was selected to be "identity provider".identity_provider blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "radius",
					Description: `(Optional) Radius server object identified by the name or UID. Must be set when "authentication-method" was selected to be "radius".`,
				},
				resource.Attribute{
					Name:        "users_directories",
					Description: `(Optional) Users directories.users_directories blocks are documented below. ` + "`" + `browser_based_authentication_portal_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "portal_web_settings",
					Description: `(Optional) Configuration of the portal web settings.portal_web_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "certificate_settings",
					Description: `(Optional) Configuration of the portal certificate settings.certificate_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "accessibility",
					Description: `(Optional) Configuration of the portal access settings.accessibility blocks are documented below. ` + "`" + `authentication_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "authentication_method",
					Description: `(Optional) Authentication method.`,
				},
				resource.Attribute{
					Name:        "radius",
					Description: `(Optional) Radius server object identified by the name or UID. Must be set when "authentication-method" was selected to be "radius".`,
				},
				resource.Attribute{
					Name:        "users_directories",
					Description: `(Optional) Users directories.users_directories blocks are documented below. ` + "`" + `identity_agent_portal_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "accessibility",
					Description: `(Optional) Configuration of the portal access settings.accessibility blocks are documented below. ` + "`" + `authorized_clients` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "client",
					Description: `(Optional) Host / Network Group Name or UID.`,
				},
				resource.Attribute{
					Name:        "client_secret",
					Description: `(Optional) Client Secret. ` + "`" + `authentication_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "users_directories",
					Description: `(Optional) Users directories.users_directories blocks are documented below. ` + "`" + `client_access_permissions` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "accessibility",
					Description: `(Optional) Configuration of the portal access settings.accessibility blocks are documented below. ` + "`" + `internal_access_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "undefined",
					Description: `(Optional) Controls portal access settings for internal interfaces, whose topology is set to 'Undefined'.`,
				},
				resource.Attribute{
					Name:        "dmz",
					Description: `(Optional) Controls portal access settings for internal interfaces, whose topology is set to 'DMZ'.`,
				},
				resource.Attribute{
					Name:        "vpn",
					Description: `(Optional) Controls portal access settings for interfaces that are part of a VPN Encryption Domain. ` + "`" + `internal_access_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "undefined",
					Description: `(Optional) Controls portal access settings for internal interfaces, whose topology is set to 'Undefined'.`,
				},
				resource.Attribute{
					Name:        "dmz",
					Description: `(Optional) Controls portal access settings for internal interfaces, whose topology is set to 'DMZ'.`,
				},
				resource.Attribute{
					Name:        "vpn",
					Description: `(Optional) Controls portal access settings for interfaces that are part of a VPN Encryption Domain. ` + "`" + `allocate_ip_address_from` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "radius_server",
					Description: `(Optional) Radius server used to authenticate the user.`,
				},
				resource.Attribute{
					Name:        "use_allocate_method",
					Description: `(Optional) Use Allocate Method.`,
				},
				resource.Attribute{
					Name:        "allocate_method",
					Description: `(Optional) Using either Manual (IP Pool) or Automatic (DHCP). Must be set when "use-allocate-method" is true.`,
				},
				resource.Attribute{
					Name:        "manual_network",
					Description: `(Optional) Manual Network. Identified by name or UID. Must be set when "allocate-method" was selected to be "manual".`,
				},
				resource.Attribute{
					Name:        "dhcp_server",
					Description: `(Optional) DHCP Server. Identified by name or UID. Must be set when "allocate-method" was selected to be "automatic".`,
				},
				resource.Attribute{
					Name:        "virtual_ip_address",
					Description: `(Optional) Virtual IPV4 address for DHCP server replies. Must be set when "allocate-method" was selected to be "automatic".`,
				},
				resource.Attribute{
					Name:        "dhcp_mac_address",
					Description: `(Optional) Calculated MAC address for DHCP allocation. Must be set when "allocate-method" was selected to be "automatic".`,
				},
				resource.Attribute{
					Name:        "optional_parameters",
					Description: `(Optional) This configuration applies to all Office Mode methods except Automatic (using DHCP) and ipassignment.conf entries which contain this data.optional_parameters blocks are documented below. ` + "`" + `custom_mode` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "social_networking_widgets",
					Description: `(Optional) Social networking widgets mode.`,
				},
				resource.Attribute{
					Name:        "url_filtering",
					Description: `(Optional) URL filtering mode. ` + "`" + `users_directories` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "external_user_profile",
					Description: `(Optional) External user profile.`,
				},
				resource.Attribute{
					Name:        "internal_users",
					Description: `(Optional) Internal users.`,
				},
				resource.Attribute{
					Name:        "users_from_external_directories",
					Description: `(Optional) Users from external directories.`,
				},
				resource.Attribute{
					Name:        "specific",
					Description: `(Optional) LDAP AU objects identified by the name or UID. Must be set when "users-from-external-directories" was selected to be "specific".specific blocks are documented below. ` + "`" + `portal_web_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "aliases",
					Description: `(Optional) List of URL aliases that are redirected to the main portal URL.aliases blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "main_url",
					Description: `(Optional) The main URL for the web portal. ` + "`" + `certificate_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "base64_certificate",
					Description: `(Optional) The certificate file encoded in Base64 with padding. This file must be in the`,
				},
				resource.Attribute{
					Name:        "base64_password",
					Description: `(Optional) Password (encoded in Base64 with padding) for the certificate file. ` + "`" + `accessibility` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "allow_access_from",
					Description: `(Optional) Allowed access to the web portal (based on interfaces, or security policy).`,
				},
				resource.Attribute{
					Name:        "internal_access_settings",
					Description: `(Optional) Configuration of the additional portal access settings for internal interfaces only.internal_access_settings blocks are documented below. ` + "`" + `users_directories` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "external_user_profile",
					Description: `(Optional) External user profile.`,
				},
				resource.Attribute{
					Name:        "internal_users",
					Description: `(Optional) Internal users.`,
				},
				resource.Attribute{
					Name:        "users_from_external_directories",
					Description: `(Optional) Users from external directories.`,
				},
				resource.Attribute{
					Name:        "specific",
					Description: `(Optional) LDAP AU objects identified by the name or UID. Must be set when "users-from-external-directories" was selected to be "specific".specific blocks are documented below. ` + "`" + `accessibility` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "allow_access_from",
					Description: `(Optional) Allowed access to the web portal (based on interfaces, or security policy).`,
				},
				resource.Attribute{
					Name:        "internal_access_settings",
					Description: `(Optional) Configuration of the additional portal access settings for internal interfaces only.internal_access_settings blocks are documented below. ` + "`" + `users_directories` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "external_user_profile",
					Description: `(Optional) External user profile.`,
				},
				resource.Attribute{
					Name:        "internal_users",
					Description: `(Optional) Internal users.`,
				},
				resource.Attribute{
					Name:        "users_from_external_directories",
					Description: `(Optional) Users from external directories.`,
				},
				resource.Attribute{
					Name:        "specific",
					Description: `(Optional) LDAP AU objects identified by the name or UID. Must be set when "users-from-external-directories" was selected to be "specific".specific blocks are documented below. ` + "`" + `accessibility` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "allow_access_from",
					Description: `(Optional) Allowed access to the web portal (based on interfaces, or security policy).`,
				},
				resource.Attribute{
					Name:        "internal_access_settings",
					Description: `(Optional) Configuration of the additional portal access settings for internal interfaces only.internal_access_settings blocks are documented below. ` + "`" + `optional_parameters` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "use_primary_dns_server",
					Description: `(Optional) Use Primary DNS Server.`,
				},
				resource.Attribute{
					Name:        "primary_dns_server",
					Description: `(Optional) Primary DNS Server. Identified by name or UID. Must be set when "use-primary-dns-server" is true and can not be set when "use-primary-dns-server" is false.`,
				},
				resource.Attribute{
					Name:        "use_first_backup_dns_server",
					Description: `(Optional) Use First Backup DNS Server.`,
				},
				resource.Attribute{
					Name:        "first_backup_dns_server",
					Description: `(Optional) First Backup DNS Server. Identified by name or UID. Must be set when "use-first-backup-dns-server" is true and can not be set when "use-first-backup-dns-server" is false.`,
				},
				resource.Attribute{
					Name:        "use_second_backup_dns_server",
					Description: `(Optional) Use Second Backup DNS Server.`,
				},
				resource.Attribute{
					Name:        "second_backup_dns_server",
					Description: `(Optional) Second Backup DNS Server. Identified by name or UID. Must be set when "use-second-backup-dns-server" is true and can not be set when "use-second-backup-dns-server" is false.`,
				},
				resource.Attribute{
					Name:        "dns_suffixes",
					Description: `(Optional) DNS Suffixes.`,
				},
				resource.Attribute{
					Name:        "use_primary_wins_server",
					Description: `(Optional) Use Primary WINS Server.`,
				},
				resource.Attribute{
					Name:        "primary_wins_server",
					Description: `(Optional) Primary WINS Server. Identified by name or UID. Must be set when "use-primary-wins-server" is true and can not be set when "use-primary-wins-server" is false.`,
				},
				resource.Attribute{
					Name:        "use_first_backup_wins_server",
					Description: `(Optional) Use First Backup WINS Server.`,
				},
				resource.Attribute{
					Name:        "first_backup_wins_server",
					Description: `(Optional) First Backup WINS Server. Identified by name or UID. Must be set when "use-first-backup-wins-server" is true and can not be set when "use-first-backup-wins-server" is false.`,
				},
				resource.Attribute{
					Name:        "use_second_backup_wins_server",
					Description: `(Optional) Use Second Backup WINS Server.`,
				},
				resource.Attribute{
					Name:        "second_backup_wins_server",
					Description: `(Optional) Second Backup WINS Server. Identified by name or UID. Must be set when "use-second-backup-wins-server" is true and can not be set when "use-second-backup-wins-server" is false. ` + "`" + `internal_access_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "undefined",
					Description: `(Optional) Controls portal access settings for internal interfaces, whose topology is set to 'Undefined'.`,
				},
				resource.Attribute{
					Name:        "dmz",
					Description: `(Optional) Controls portal access settings for internal interfaces, whose topology is set to 'DMZ'.`,
				},
				resource.Attribute{
					Name:        "vpn",
					Description: `(Optional) Controls portal access settings for interfaces that are part of a VPN Encryption Domain. ` + "`" + `internal_access_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "undefined",
					Description: `(Optional) Controls portal access settings for internal interfaces, whose topology is set to 'Undefined'.`,
				},
				resource.Attribute{
					Name:        "dmz",
					Description: `(Optional) Controls portal access settings for internal interfaces, whose topology is set to 'DMZ'.`,
				},
				resource.Attribute{
					Name:        "vpn",
					Description: `(Optional) Controls portal access settings for interfaces that are part of a VPN Encryption Domain. ` + "`" + `internal_access_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "undefined",
					Description: `(Optional) Controls portal access settings for internal interfaces, whose topology is set to 'Undefined'.`,
				},
				resource.Attribute{
					Name:        "dmz",
					Description: `(Optional) Controls portal access settings for internal interfaces, whose topology is set to 'DMZ'.`,
				},
				resource.Attribute{
					Name:        "vpn",
					Description: `(Optional) Controls portal access settings for interfaces that are part of a VPN Encryption Domain.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_smtp_server",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Smtp Server.`,
			Description:      ``,
			Keywords: []string{
				"management",
				"smtp",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) The SMTP port to use.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `(Required) The SMTP server address.`,
				},
				resource.Attribute{
					Name:        "authentication",
					Description: `(Optional) Does the mail server requires authentication.`,
				},
				resource.Attribute{
					Name:        "encryption",
					Description: `(Optional) Encryption type.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) A password for the SMTP server.Required only if authentication is set to true.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) A username for the SMTP server.Required only if authentication is set to true.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "domains_to_process",
					Description: `(Optional) Indicates which domains to process the commands on. It cannot be used with the details-level full, must be run from the System Domain only and with ignore-warnings true. Valid values are: CURRENT_DOMAIN, ALL_DOMAINS_ON_THIS_SERVER.domains_to_process blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_tacacs_group",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Tacacs Group.`,
			Description: `

This resource allows you to  add/update/delete Check Point Tacacs Group.

`,
			Keywords: []string{
				"management",
				"tacacs",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `(Optional) Collection of tacacs servers identified by the name or UID.members blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_tacacs_server",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Tacacs Server.`,
			Description: `

This resource allows you to execute Check Point Tacacs Server.

`,
			Keywords: []string{
				"management",
				"tacacs",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Optional) The server's secret key.<br><font color="red">Required only when</font> "server-type" was selected to be "TACACS+".`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `(Required) The UID or Name of the host that is the TACACS Server.`,
				},
				resource.Attribute{
					Name:        "encryption",
					Description: `(Optional) Is there a secret key defined on the server. Must be set true when "server-type" was selected to be "TACACS+".`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) The priority of the TACACS Server in case it is a member of a TACACS Group.`,
				},
				resource.Attribute{
					Name:        "server_type",
					Description: `(Optional) Server type, TACACS or TACACS+.`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `(Optional) Server service, only relevant when "server-type" is TACACS.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_tag",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Tag.`,
			Description: `

This resource allows you to add/update/delete Check Point Tag.

`,
			Keywords: []string{
				"management",
				"tag",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name Must be unique in the domain.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_test_sic_status",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Test Sic Status.`,
			Description:      ``,
			Keywords: []string{
				"management",
				"test",
				"sic",
				"status",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Gateway, cluster member or Check Point host name.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Gateway, cluster member or Check Point host name. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_threat_exception",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to add/update/delete Check Point Threat Exception.`,
			Description: `

This resource allows you to add/update/delete Check Point Threat Exception.

`,
			Keywords: []string{
				"management",
				"threat",
				"exception",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "layer",
					Description: `(Required) Layer that the rule belongs to identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "position",
					Description: `(Required) Position in the rulebase. Position blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the exception.`,
				},
				resource.Attribute{
					Name:        "exception_group_uid",
					Description: `(Optional) The UID of the exception-group.`,
				},
				resource.Attribute{
					Name:        "exception_group_name",
					Description: `(Optional) The name of the exception-group.`,
				},
				resource.Attribute{
					Name:        "rule_uid",
					Description: `(Optional) The UID of the parent rule.`,
				},
				resource.Attribute{
					Name:        "rule_name",
					Description: `(Optional) The name of the parent rule.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Action-the enforced profile.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Enable/Disable the rule.`,
				},
				resource.Attribute{
					Name:        "install_on",
					Description: `(Optional) Which Gateways identified by the name or UID to install the policy on.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) Collection of Network objects identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "source_negate",
					Description: `(Optional) True if negate is set for source.`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `(Optional) Collection of Network objects identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "destination_negate",
					Description: `(Optional) True if negate is set for destination.`,
				},
				resource.Attribute{
					Name:        "protected_scope",
					Description: `(Optional) Collection of objects defining Protected Scope identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "protected_scope_negate",
					Description: `(Optional) True if negate is set for Protected Scope.`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `(Optional) Collection of Network objects identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "service_negate",
					Description: `(Optional) True if negate is set for service.`,
				},
				resource.Attribute{
					Name:        "protection_or_site",
					Description: `(Optional) Collection of protection or site objects identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "track",
					Description: `(Optional) Packet tracking.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `(Computed) Owner UID. ` + "`" + `position` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "top",
					Description: `(Optional) Add rule at the top of the rulebase.`,
				},
				resource.Attribute{
					Name:        "above",
					Description: `(Optional) Add rule above specific section/rule identified by uid or name.`,
				},
				resource.Attribute{
					Name:        "below",
					Description: `(Optional) Add rule below specific section/rule identified by uid or name.`,
				},
				resource.Attribute{
					Name:        "bottom",
					Description: `(Optional) Add rule at the bottom of the rulebase. ## Import ` + "`" + `checkpoint_management_threat_exception` + "`" + ` can be imported by using the following format: LAYER_UID;exception_group_uid or rule_uid;EXCEPTION_GROUP_UID or PARENT_RULE_UID;EXCEPTION_GROUP_UID ` + "`" + `` + "`" + `` + "`" + ` $ terraform import checkpoint_management_threat_exception.example "Standard Threat Prevention;exception_group_uid;9423d36f-2d66-4754-b9e2-e9f4493751e5;9423d36f-2d66-4754-b9e2-e9f4493751d3" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_threat_indicator",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to add/update/delete Check Point Threat Indicator.`,
			Description: `

This resource allows you to add/update/delete Check Point Threat Indicator.

`,
			Keywords: []string{
				"management",
				"threat",
				"indicator",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name. Should be unique in the domain.`,
				},
				resource.Attribute{
					Name:        "observables",
					Description: `(Optional) The indicator's observables. Indicator's observables blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) The indicator's action.`,
				},
				resource.Attribute{
					Name:        "profile_overrides",
					Description: `(Optional) Profiles in which to override the indicator's default action. Profile Overrides blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers. ` + "`" + `observables` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name. Should be unique in the domain.`,
				},
				resource.Attribute{
					Name:        "md5",
					Description: `(Optional) A valid MD5 sequence.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Optional) A valid URL.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) A valid IP-Address.`,
				},
				resource.Attribute{
					Name:        "ip_address_first",
					Description: `(Optional) A valid IP-Address, the beginning of the range. If you configure this parameter with a value, you must also configure the value of the 'ip-address-last' parameter.`,
				},
				resource.Attribute{
					Name:        "ip_address_last",
					Description: `(Optional) A valid IP-Address, the end of the range. If you configure this parameter with a value, you must also configure the value of the 'ip-address-first' parameter.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional) The name of a domain.`,
				},
				resource.Attribute{
					Name:        "mail_to",
					Description: `(Optional) A valid E-Mail address, recipient filed.`,
				},
				resource.Attribute{
					Name:        "mail_from",
					Description: `(Optional) A valid E-Mail address, sender field.`,
				},
				resource.Attribute{
					Name:        "mail_cc",
					Description: `(Optional) A valid E-Mail address, cc field.`,
				},
				resource.Attribute{
					Name:        "mail_reply_to",
					Description: `(Optional) A valid E-Mail address, reply-to field.`,
				},
				resource.Attribute{
					Name:        "mail_subject",
					Description: `(Optional) Subject of E-Mail.`,
				},
				resource.Attribute{
					Name:        "confidence",
					Description: `(Optional) The confidence level the indicator has that a real threat has been uncovered.`,
				},
				resource.Attribute{
					Name:        "product",
					Description: `(Optional) The software blade that processes the observable: AV - AntiVirus, AB - AntiBot.`,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `(Optional) The severity level of the threat. ` + "`" + `profile_overrides` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) The indicator's action in this profile.`,
				},
				resource.Attribute{
					Name:        "profile",
					Description: `(Optional) The profile in which to override the indicator's action.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_threat_ioc_feed",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Threat Ioc Feed.`,
			Description:      ``,
			Keywords: []string{
				"management",
				"threat",
				"ioc",
				"feed",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "feed_url",
					Description: `(Optional) URL of the feed. URL should be written as http or https.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) The feed indicator's action.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `(Optional) Certificate SHA-1 fingerprint to access the feed.`,
				},
				resource.Attribute{
					Name:        "custom_comment",
					Description: `(Optional) Custom IOC feed - the column number of comment.`,
				},
				resource.Attribute{
					Name:        "custom_confidence",
					Description: `(Optional) Custom IOC feed - the column number of confidence.`,
				},
				resource.Attribute{
					Name:        "custom_header",
					Description: `(Optional) Custom HTTP headers.custom_header blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "custom_name",
					Description: `(Optional) Custom IOC feed - the column number of name.`,
				},
				resource.Attribute{
					Name:        "custom_severity",
					Description: `(Optional) Custom IOC feed - the column number of severity.`,
				},
				resource.Attribute{
					Name:        "custom_type",
					Description: `(Optional) Custom IOC feed - the column number of type in case a specific type is not chosen.`,
				},
				resource.Attribute{
					Name:        "custom_value",
					Description: `(Optional) Custom IOC feed - the column number of value in case a specific type is chosen.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Sets whether this indicator feed is enabled.`,
				},
				resource.Attribute{
					Name:        "feed_type",
					Description: `(Optional) Feed type to be enforced.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) password for authenticating with the URL.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "use_custom_feed_settings",
					Description: `(Optional) Set in order to configure a custom indicator feed.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) username for authenticating with the URL.`,
				},
				resource.Attribute{
					Name:        "fields_delimiter",
					Description: `(Optional) The delimiter that separates between the columns in the feed.`,
				},
				resource.Attribute{
					Name:        "ignore_lines_that_start_with",
					Description: `(Optional) A prefix that will determine which lines to ignore.`,
				},
				resource.Attribute{
					Name:        "use_gateway_proxy",
					Description: `(Optional) Use the gateway's proxy for retrieving the feed.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. ` + "`" + `custom_header` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "header_name",
					Description: `(Optional) The name of the HTTP header we wish to add.`,
				},
				resource.Attribute{
					Name:        "header_value",
					Description: `(Optional) The name of the HTTP value we wish to add.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_threat_layer",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to add/update/delete Check Point Threat Layer.`,
			Description: `

This resource allows you to add/update/delete Check Point Threat Layer.

`,
			Keywords: []string{
				"management",
				"threat",
				"layer",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
				resource.Attribute{
					Name:        "add_default_rule",
					Description: `(Optional) Indicates whether to include a default rule in the new layer.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_threat_profile",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to add/update/delete Check Point Threat Profile.`,
			Description: `

This resource allows you to add/update/delete Check Point Threat Profile.

`,
			Keywords: []string{
				"management",
				"threat",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name. Should be unique in the domain.`,
				},
				resource.Attribute{
					Name:        "active_protections_performance_impact",
					Description: `(Optional) Protections with this performance impact only will be activated in the profile.`,
				},
				resource.Attribute{
					Name:        "active_protections_severity",
					Description: `(Optional) Protections with this severity only will be activated in the profile.`,
				},
				resource.Attribute{
					Name:        "confidence_level_high",
					Description: `(Optional) Action for protections with high confidence level.`,
				},
				resource.Attribute{
					Name:        "confidence_level_medium",
					Description: `(Optional) Action for protections with medium confidence level.`,
				},
				resource.Attribute{
					Name:        "confidence_level_low",
					Description: `(Optional) Action for protections with low confidence level.`,
				},
				resource.Attribute{
					Name:        "indicator_overrides",
					Description: `(Optional) Indicators whose action will be overridden in this profile. indicator_overrides blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "ips_settings",
					Description: `(Optional) IPS blade settings. ips_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "malicious_mail_policy_settings",
					Description: `(Optional) Malicious Mail Policy for MTA Gateways. malicious_mail_policy_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "overrides",
					Description: `(Optional) Overrides per profile for this protection. overrides blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "scan_malicious_links",
					Description: `(Optional) Scans malicious links (URLs) inside email messages. scan_malicious_links blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "use_indicators",
					Description: `(Optional) Indicates whether the profile should make use of indicators.`,
				},
				resource.Attribute{
					Name:        "anti_virus",
					Description: `(Optional) Is Anti-Virus blade activated.`,
				},
				resource.Attribute{
					Name:        "anti_bot",
					Description: `(Optional) Is Anti-Bot blade activated.`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `(Optional) Is IPS blade activated.`,
				},
				resource.Attribute{
					Name:        "threat_emulation",
					Description: `(Optional) Is Threat Emulation blade activated.`,
				},
				resource.Attribute{
					Name:        "use_extended_attributes",
					Description: `(Optional) Whether to activate/deactivate IPS protections according to the extended attributes.`,
				},
				resource.Attribute{
					Name:        "activate_protections_by_extended_attributes",
					Description: `(Optional) Activate protections by these extended attributes. activate_protections_by_extended_attributes blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "deactivate_protections_by_extended_attributes",
					Description: `(Optional) Deactivate protections by these extended attributes. deactivate_protections_by_extended_attributes blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. ` + "`" + `indicator_overrides` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) The indicator's action in this profile.`,
				},
				resource.Attribute{
					Name:        "indicator",
					Description: `(Optional) The indicator whose action is to be overriden. ` + "`" + `ips_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "exclude_protection_with_performance_impact",
					Description: `(Optional) Whether to exclude protections depending on their level of performance impact.`,
				},
				resource.Attribute{
					Name:        "exclude_protection_with_performance_impact_mode",
					Description: `(Optional) Exclude protections with this level of performance impact.`,
				},
				resource.Attribute{
					Name:        "exclude_protection_with_severity",
					Description: `(Optional) Whether to exclude protections depending on their level of severity.`,
				},
				resource.Attribute{
					Name:        "exclude_protection_with_severity_mode",
					Description: `(Optional) Exclude protections with this level of severity.`,
				},
				resource.Attribute{
					Name:        "newly_updated_protections",
					Description: `(Optional) Activation of newly updated protections. ` + "`" + `malicious_mail_policy_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "add_customized_text_to_email_body",
					Description: `(Optional) Add customized text to the malicious email body.`,
				},
				resource.Attribute{
					Name:        "add_email_subject_prefix",
					Description: `(Optional) Add a prefix to the malicious email subject.`,
				},
				resource.Attribute{
					Name:        "add_x_header_to_email",
					Description: `(Optional) Add an X-Header to the malicious email.`,
				},
				resource.Attribute{
					Name:        "email_action",
					Description: `(Optional) Block - block the entire malicious email. Allow - pass the malicious email and apply email changes (like: remove attachments and links, add x-header, etc...).`,
				},
				resource.Attribute{
					Name:        "email_body_customized_text",
					Description: `(Optional) Customized text for the malicious email body. Available predefined fields: $verdicts$ - the malicious/error attachments/links verdict.`,
				},
				resource.Attribute{
					Name:        "email_subject_prefix_text",
					Description: `(Optional) Prefix for the malicious email subject.`,
				},
				resource.Attribute{
					Name:        "failed_to_scan_attachments_text",
					Description: `(Optional) Replace attachments that failed to be scanned with this text. Available predefined fields: $filename$ - the malicious file name. $md5$ - MD5 of the malicious file.`,
				},
				resource.Attribute{
					Name:        "malicious_attachments_text",
					Description: `(Optional) Replace malicious attachments with this text. Available predefined fields: $filename$ - the malicious file name. $md5$ - MD5 of the malicious file.`,
				},
				resource.Attribute{
					Name:        "malicious_links_text",
					Description: `(Optional) Replace malicious links with this text. Available predefined fields: $neutralized_url$ - neutralized malicious link.`,
				},
				resource.Attribute{
					Name:        "remove_attachments_and_links",
					Description: `(Optional) Remove attachments and links from the malicious email.`,
				},
				resource.Attribute{
					Name:        "send_copy",
					Description: `(Optional) Send a copy of the malicious email to the recipient list.`,
				},
				resource.Attribute{
					Name:        "send_copy_list",
					Description: `(Optional) Recipient list to send a copy of the malicious email. ` + "`" + `overrides` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "protection",
					Description: `(Required) IPS protection identified by name.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) Protection action.`,
				},
				resource.Attribute{
					Name:        "capture_packets",
					Description: `(Optional) Capture packets.`,
				},
				resource.Attribute{
					Name:        "track",
					Description: `(Optional) Tracking method for protection.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `Default settings. default blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "final",
					Description: `Final settings. final blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "protection_external_info",
					Description: `Collection of industry reference (CVE).`,
				},
				resource.Attribute{
					Name:        "protection_uid",
					Description: `IPS protection unique identifier. ` + "`" + `scan_malicious_links` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "max_bytes",
					Description: `(Optional) Scan links in the first bytes of the mail body.`,
				},
				resource.Attribute{
					Name:        "max_links",
					Description: `(Optional) Maximum links to scan in mail body. ` + "`" + `activate_protections_by_extended_attributes` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) IPS tag unique identifier.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) IPS tag name.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `(Optional) IPS tag category name.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `Collection of IPS protection extended attribute values (name and uid). ` + "`" + `deactivate_protections_by_extended_attributes` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) IPS tag unique identifier.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) IPS tag name.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `(Optional) IPS tag category name.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `Collection of IPS protection extended attribute values (name and uid). ` + "`" + `default` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Protection action.`,
				},
				resource.Attribute{
					Name:        "capture_packets",
					Description: `Capture packets.`,
				},
				resource.Attribute{
					Name:        "track",
					Description: `Tracking method for protection. ` + "`" + `final` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Protection action.`,
				},
				resource.Attribute{
					Name:        "capture_packets",
					Description: `Capture packets.`,
				},
				resource.Attribute{
					Name:        "track",
					Description: `Tracking method for protection.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_threat_rule",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to add/update/delete Check Point Threat Rule.`,
			Description: `

This resource allows you to add/update/delete Check Point Threat Rule.

`,
			Keywords: []string{
				"management",
				"threat",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "layer",
					Description: `(Required) Layer that the rule belongs to identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "position",
					Description: `(Required) Position in the rulebase. Position blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Rule name.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Action-the enforced profile.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Enable/Disable the rule.`,
				},
				resource.Attribute{
					Name:        "install_on",
					Description: `(Optional) Which Gateways identified by the name or UID to install the policy on.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) Collection of Network objects identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "source_negate",
					Description: `(Optional) True if negate is set for source.`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `(Optional) Collection of Network objects identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "destination_negate",
					Description: `(Optional) True if negate is set for destination.`,
				},
				resource.Attribute{
					Name:        "protected_scope",
					Description: `(Optional) Collection of objects defining Protected Scope identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "protected_scope_negate",
					Description: `(Optional) True if negate is set for Protected Scope.`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `(Optional) Collection of Network objects identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "service_negate",
					Description: `(Optional) True if negate is set for service.`,
				},
				resource.Attribute{
					Name:        "track",
					Description: `(Optional) Packet tracking.`,
				},
				resource.Attribute{
					Name:        "track_settings",
					Description: `(Optional) Threat rule track settings. track_settings block are documented below.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "exceptions",
					Description: `(Computed) Collection of the rule's exceptions identified by UID. ` + "`" + `position` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "top",
					Description: `(Optional) Add rule at the top of the rulebase.`,
				},
				resource.Attribute{
					Name:        "above",
					Description: `(Optional) Add rule above specific section/rule identified by uid or name.`,
				},
				resource.Attribute{
					Name:        "below",
					Description: `(Optional) Add rule below specific section/rule identified by uid or name.`,
				},
				resource.Attribute{
					Name:        "bottom",
					Description: `(Optional) Add rule at the bottom of the rulebase. ` + "`" + `track_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "packet_capture",
					Description: `(Optional) Packet capture. ## Import ` + "`" + `checkpoint_management_threat_rule` + "`" + ` can be imported by using the following format: LAYER_NAME;RULE_UID ` + "`" + `` + "`" + `` + "`" + ` $ terraform import checkpoint_management_threat_rule.example "Layer_Name;9423d36f-2d66-4754-b9e2-e9f4493751d3" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_time",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Time.`,
			Description: `

This resource allows you to execute Check Point Time.

`,
			Keywords: []string{
				"management",
				"time",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Time object name. Cannot be more than 11 characters. Should be unique in the domain.`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `(Optional) End time. Note: Each gateway may interpret this time differently according to its time zone.`,
				},
				resource.Attribute{
					Name:        "end_never",
					Description: `(Optional) End never.`,
				},
				resource.Attribute{
					Name:        "hours_ranges",
					Description: `(Optional) Hours recurrence. Note: Each gateway may interpret this time differently according to its time zone.`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `(Optional) Starting time. Note: Each gateway may interpret this time differently according to its time zone.`,
				},
				resource.Attribute{
					Name:        "start_now",
					Description: `(Optional) Start immediately.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "recurrence",
					Description: `(Optional) Days recurrence.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. ` + "`" + `end` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "date",
					Description: `(Optional) Date in format dd-MMM-yyyy.`,
				},
				resource.Attribute{
					Name:        "iso_8601",
					Description: `(Optional) Date and time represented in international ISO 8601 format. Time zone information is ignored.`,
				},
				resource.Attribute{
					Name:        "posix",
					Description: `(Optional) Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.`,
				},
				resource.Attribute{
					Name:        "time",
					Description: `(Optional) Time in format HH:mm. ` + "`" + `start` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "date",
					Description: `(Optional) Date in format dd-MMM-yyyy.`,
				},
				resource.Attribute{
					Name:        "iso_8601",
					Description: `(Optional) Date and time represented in international ISO 8601 format. Time zone information is ignored.`,
				},
				resource.Attribute{
					Name:        "posix",
					Description: `(Optional) Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.`,
				},
				resource.Attribute{
					Name:        "time",
					Description: `(Optional) Time in format HH:mm. ` + "`" + `recurrence` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "days",
					Description: `(Optional) Valid on specific days. Multiple options, support range of days in months. Example:["1","3","9-20"].`,
				},
				resource.Attribute{
					Name:        "month",
					Description: `(Optional) Valid on month. Example: "1", "2","12","Any".`,
				},
				resource.Attribute{
					Name:        "pattern",
					Description: `(Optional) Valid on "Daily", "Weekly", "Monthly" base.`,
				},
				resource.Attribute{
					Name:        "weekday",
					Description: `(Optional) Valid on weekdays. Example: "Sun", "Mon"..."Sat". ` + "`" + `hours_ranges` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Is hour range enabled.`,
				},
				resource.Attribute{
					Name:        "from",
					Description: `(Optional) Time in format HH:MM.`,
				},
				resource.Attribute{
					Name:        "index",
					Description: `(Optional) Hour range index.`,
				},
				resource.Attribute{
					Name:        "to",
					Description: `(Optional) Time in format HH:MM.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_time_group",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Time Group.`,
			Description: `

This resource allows you to execute Check Point Time Group.

`,
			Keywords: []string{
				"management",
				"time",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `(Optional) Collection of Time Group objects identified by the name or UID.members blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_trusted_client",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Trusted Client.`,
			Description: `

This resource allows you to execute Check Point Trusted Client.

`,
			Keywords: []string{
				"management",
				"trusted",
				"client",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `(Optional) IPv4 address.`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `(Optional) IPv6 address.`,
				},
				resource.Attribute{
					Name:        "ipv4_address_first",
					Description: `(Optional) First IPv4 address in the range.`,
				},
				resource.Attribute{
					Name:        "ipv6_address_first",
					Description: `(Optional) First IPv6 address in the range.`,
				},
				resource.Attribute{
					Name:        "ipv4_address_last",
					Description: `(Optional) Last IPv4 address in the range.`,
				},
				resource.Attribute{
					Name:        "ipv6_address_last",
					Description: `(Optional) Last IPv6 address in the range.`,
				},
				resource.Attribute{
					Name:        "domains_assignment",
					Description: `(Optional) Domains to be added to this profile. Use domain name only. See example below: "add-trusted-client (with domain)".`,
				},
				resource.Attribute{
					Name:        "mask_length4",
					Description: `(Optional) IPv4 mask length.`,
				},
				resource.Attribute{
					Name:        "mask_length6",
					Description: `(Optional) IPv6 mask length.`,
				},
				resource.Attribute{
					Name:        "multi_domain_server_trusted_client",
					Description: `(Optional) Let this trusted client connect to all Multi-Domain Servers in the deployment.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Trusted client type.`,
				},
				resource.Attribute{
					Name:        "wild_card",
					Description: `(Optional) IP wild card (e.g. 192.0.2.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_uninstall_software_package",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Uninstall Software Package.`,
			Description: `

This command resource allows you to execute Check Point Uninstall Software Package.

`,
			Keywords: []string{
				"management",
				"uninstall",
				"software",
				"package",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the software package.`,
				},
				resource.Attribute{
					Name:        "targets",
					Description: `(Required) On what targets to execute this command. Targets may be identified by their name, or object unique identifier.targets blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "cluster_installation_settings",
					Description: `(Optional) Installation settings for cluster.cluster_installation_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "concurrency_limit",
					Description: `(Optional) The number of targets, on which the same package is installed at the same time.`,
				},
				resource.Attribute{
					Name:        "task_id",
					Description: `(Computed) Asynchronous task unique identifier. ` + "`" + `cluster_installation_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "cluster_delay",
					Description: `(Optional) The delay between end of installation on one cluster members and start of installation on the next cluster member.`,
				},
				resource.Attribute{
					Name:        "cluster_strategy",
					Description: `(Optional) The cluster installation strategy. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_unlock_administrator",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Unlock Administrator.`,
			Description: `

This command resource allows you to execute Check Point Unlock Administrator.

`,
			Keywords: []string{
				"management",
				"unlock",
				"administrator",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_update_provisioned_satellites",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Update Provisioned Satellites.`,
			Description:      ``,
			Keywords: []string{
				"management",
				"update",
				"provisioned",
				"satellites",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vpn_center_gateways",
					Description: `(Required) On what targets to execute this command. Targets may be identified by their name, or object unique identifier. The targets should be a corporate gateways.vpn_center_gateways blocks are documented below. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_update_updatable_objects_repository_content",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Update Updatable Objects Repository Content.`,
			Description: `

This command resource allows you to execute Check Point Update Updatable Objects Repository Content.

`,
			Keywords: []string{
				"management",
				"update",
				"updatable",
				"objects",
				"repository",
				"content",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "task_id",
					Description: `(Computed) Asynchronous task unique identifier. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_user",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point User.`,
			Description: `

This resource allows you to execute Check Point User.

`,
			Keywords: []string{
				"management",
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) User email.`,
				},
				resource.Attribute{
					Name:        "expiration_date",
					Description: `(Optional) Expiration date in format: yyyy-MM-dd.`,
				},
				resource.Attribute{
					Name:        "phone_number",
					Description: `(Optional) User phone number.`,
				},
				resource.Attribute{
					Name:        "authentication_method",
					Description: `(Optional) Authentication method.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Checkpoint password authentication method identified by the name or UID. Must be set when "authentication-method" was selected to be "Check Point Password".`,
				},
				resource.Attribute{
					Name:        "radius_server",
					Description: `(Optional) RADIUS server object identified by the name or UID. Must be set when "authentication-method" was selected to be "RADIUS".`,
				},
				resource.Attribute{
					Name:        "tacacs_server",
					Description: `(Optional) TACACS server object identified by the name or UID. Must be set when "authentication-method" was selected to be "TACACS".`,
				},
				resource.Attribute{
					Name:        "connect_on_days",
					Description: `(Optional) Days users allow to connect.`,
				},
				resource.Attribute{
					Name:        "connect_daily",
					Description: `(Optional) Connect every day.`,
				},
				resource.Attribute{
					Name:        "from_hour",
					Description: `(Optional) Allow users connect from hour.`,
				},
				resource.Attribute{
					Name:        "to_hour",
					Description: `(Optional) Allow users connect until hour.`,
				},
				resource.Attribute{
					Name:        "allowed_locations",
					Description: `(Optional) User allowed locations. allowed_locations blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "encryption",
					Description: `(Optional) User encryption. encryption blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Optional) User template name or UID. ` + "`" + `allowed_locations` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "destinations",
					Description: `(Optional) Collection of allowed destination locations name or uid.`,
				},
				resource.Attribute{
					Name:        "sources",
					Description: `(Optional) Collection of allowed source locations name or uid. ` + "`" + `encryption` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "enable_ike",
					Description: `(Optional) Enable IKE encryption for users.`,
				},
				resource.Attribute{
					Name:        "enable_public_key",
					Description: `(Optional) Enable IKE public key.`,
				},
				resource.Attribute{
					Name:        "enable_shared_secret",
					Description: `(Optional) Enable IKE shared secret.`,
				},
				resource.Attribute{
					Name:        "shared_secret",
					Description: `(Optional) IKE shared secret.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_user_group",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point User Group.`,
			Description: `

This resource allows you to execute Check Point User Group.

`,
			Keywords: []string{
				"management",
				"user",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) Email Address.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `(Optional) Collection of User Group objects identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_user_template",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point User Template.`,
			Description: `

This resource allows you to execute Check Point User Template.

`,
			Keywords: []string{
				"management",
				"user",
				"template",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "expiration_by_global_properties",
					Description: `(Optional) Expiration date according to global properties.`,
				},
				resource.Attribute{
					Name:        "expiration_date",
					Description: `(Optional) Expiration date in format: yyyy-MM-dd.`,
				},
				resource.Attribute{
					Name:        "authentication_method",
					Description: `(Optional) Authentication method.`,
				},
				resource.Attribute{
					Name:        "radius_server",
					Description: `(Optional) RADIUS server object identified by the name or UID. Must be set when "authentication-method" was selected to be "RADIUS".`,
				},
				resource.Attribute{
					Name:        "tacacs_server",
					Description: `(Optional) TACACS server object identified by the name or UID. Must be set when "authentication-method" was selected to be "TACACS".`,
				},
				resource.Attribute{
					Name:        "connect_on_days",
					Description: `(Optional) Days users allow to connect.`,
				},
				resource.Attribute{
					Name:        "connect_daily",
					Description: `(Optional) Connect every day.`,
				},
				resource.Attribute{
					Name:        "from_hour",
					Description: `(Optional) Allow users connect from hour.`,
				},
				resource.Attribute{
					Name:        "to_hour",
					Description: `(Optional) Allow users connect until hour.`,
				},
				resource.Attribute{
					Name:        "allowed_locations",
					Description: `(Optional) User allowed locations. allowed_locations blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "encryption",
					Description: `(Optional) User encryption. encryption blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. ` + "`" + `allowed_locations` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "destinations",
					Description: `(Optional) Collection of allowed destination locations name or uid.destinations blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "sources",
					Description: `(Optional) Collection of allowed source locations name or uid.sources blocks are documented below. ` + "`" + `encryption` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "enable_ike",
					Description: `(Optional) Enable IKE encryption for users.`,
				},
				resource.Attribute{
					Name:        "enable_public_key",
					Description: `(Optional) Enable IKE public key.`,
				},
				resource.Attribute{
					Name:        "enable_shared_secret",
					Description: `(Optional) Enable IKE shared secret.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_verify_policy",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Verify Policy.`,
			Description: `

This command resource allows you to execute Check Point Verify Policy.

`,
			Keywords: []string{
				"management",
				"verify",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_package",
					Description: `(Required) Policy package identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "task_id",
					Description: `(Computed) Asynchronous task unique identifier. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_verify_revert",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Verify Revert.`,
			Description: `

This command resource allows you to execute Check Point Verify Revert.

`,
			Keywords: []string{
				"management",
				"verify",
				"revert",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "to_session",
					Description: `(Required) Session unique identifier. Specify the session you would like to verify a revert operation to.`,
				},
				resource.Attribute{
					Name:        "task_id",
					Description: `(Computed) Asynchronous task unique identifier. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_verify_software_package",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Verify Software Package.`,
			Description: `

This command resource allows you to execute Check Point Verify Software Package.

`,
			Keywords: []string{
				"management",
				"verify",
				"software",
				"package",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the software package.`,
				},
				resource.Attribute{
					Name:        "targets",
					Description: `(Required) On what targets to execute this command. Targets may be identified by their name, or object unique identifier.targets blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "concurrency_limit",
					Description: `(Optional) The number of targets, on which the same package is installed at the same time.`,
				},
				resource.Attribute{
					Name:        "task_id",
					Description: `(Computed) Asynchronous task unique identifier. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_vmware_data_center_server",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point vmware data center server.`,
			Description: `

This resource allows you to execute Check Point VMware Data Center Server.

`,
			Keywords: []string{
				"management",
				"vmware",
				"data",
				"center",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Object type. nsx or nsxt or vcenter.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required) IP Address or hostname of the VMware server.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) Username of the VMware server.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Password of the VMware server.`,
				},
				resource.Attribute{
					Name:        "password_base64",
					Description: `(Optional) Password of the VMware server encoded in Base64.`,
				},
				resource.Attribute{
					Name:        "certificate_fingerprint",
					Description: `(Optional) Specify the SHA-1 or SHA-256 fingerprint of the Data Center Server's certificate.`,
				},
				resource.Attribute{
					Name:        "unsafe_auto_accept",
					Description: `(Optional) When set to false, the current Data Center Server's certificate should be trusted, either by providing the certificate-fingerprint argument or by relying on a previously trusted certificate of this hostname. When set to true, trust the current Data Center Server's certificate as-is.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers. tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_vpn_community_meshed",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Vpn Community Meshed.`,
			Description: `

This resource allows you to execute Check Point Vpn Community Meshed.

`,
			Keywords: []string{
				"management",
				"vpn",
				"community",
				"meshed",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "encryption_method",
					Description: `(Optional) The encryption method to be used.`,
				},
				resource.Attribute{
					Name:        "encryption_suite",
					Description: `(Optional) The encryption suite to be used.`,
				},
				resource.Attribute{
					Name:        "gateways",
					Description: `(Optional) Collection of Gateway objects identified by the name or UID.gateways blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "ike_phase_1",
					Description: `(Optional) Ike Phase 1 settings. Only applicable when the encryption-suite is set to [custom]. ike_phase_1 blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "ike_phase_2",
					Description: `(Optional) Ike Phase 2 settings. Only applicable when the encryption-suite is set to [custom]. ike_phase_2 blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "override_vpn_domains",
					Description: `(Optional) The Overrides VPN Domains of the participants GWs. override_vpn_domains blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "shared_secrets",
					Description: `(Optional) Shared secrets for external gateways. shared_secrets blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "use_shared_secret",
					Description: `(Optional) Indicates whether the shared secret should be used for all external gateways.`,
				},
				resource.Attribute{
					Name:        "tunnel_granularity",
					Description: `(Optional) VPN tunnel sharing option to be used.`,
				},
				resource.Attribute{
					Name:        "granular_encryptions",
					Description: `(Optional) VPN granular encryption settings. granular_encryptions blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. ` + "`" + `override_vpn_domains` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `(Optional) Participant gateway in override VPN domain identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "vpn_domain",
					Description: `(Optional) VPN domain network identified by the name or UID. ` + "`" + `shared_secrets` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "external_gateway",
					Description: `(Optional) External gateway identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "shared_secret",
					Description: `(Optional) Shared secret. ` + "`" + `granular_encryptions` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "internal-gateway",
					Description: `(Required) Internally managed Check Point gateway identified by name or UID, or 'Any' for all internal-gateways participants in this community.`,
				},
				resource.Attribute{
					Name:        "external_gateway",
					Description: `(Required) Externally managed or 3rd party gateway identified by name or UID.`,
				},
				resource.Attribute{
					Name:        "encryption_method",
					Description: `(Optional) The encryption method to be used: prefer ikev2 but support ikev1, ikev2 only, ikev1 for ipv4 and ikev2 for ipv6 only.`,
				},
				resource.Attribute{
					Name:        "encryption_suite",
					Description: `(Optional) The encryption suite to be used: suite-b-gcm-256, custom, vpn b, vpn a, suite-b-gcm-128.`,
				},
				resource.Attribute{
					Name:        "ike_phase_1",
					Description: `(Optional) Ike Phase 1 settings. Only applicable when the encryption-suite is set to [custom]. ike_phase_1 blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "ike_phase_2",
					Description: `(Optional) Ike Phase 2 settings. Only applicable when the encryption-suite is set to [custom]. ike_phase_2 blocks are documented below. ` + "`" + `ike_phase_1` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "data_integrity",
					Description: `(Optional) The hash algorithm to be used.`,
				},
				resource.Attribute{
					Name:        "diffie_hellman_group",
					Description: `(Optional) The Diffie-Hellman group to be used.`,
				},
				resource.Attribute{
					Name:        "encryption_algorithm",
					Description: `(Optional) The encryption algorithm to be used.`,
				},
				resource.Attribute{
					Name:        "ike_p1_rekey_time",
					Description: `(Optional) Indicates the time interval for IKE phase 1 renegotiation. ` + "`" + `ike_phase_2` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "data_integrity",
					Description: `(Optional) The hash algorithm to be used.`,
				},
				resource.Attribute{
					Name:        "encryption_algorithm",
					Description: `(Optional) The encryption algorithm to be used.`,
				},
				resource.Attribute{
					Name:        "ike_p2_use_pfs",
					Description: `(Optional) Indicates whether Perfect Forward Secrecy (PFS) is being used for IKE phase 2.`,
				},
				resource.Attribute{
					Name:        "ike_p2_pfs_dh_grp",
					Description: `(Optional) The Diffie-Hellman group to be used: group-1, group-2, group-5, group-14, group-15, group-16, group-17, group-18, group-19, group-20, group-24.`,
				},
				resource.Attribute{
					Name:        "ike_p2_rekey_time",
					Description: `(Optional) Indicates the time interval for IKE phase 2 renegotiation.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_vpn_community_remote_access",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point VPN Community Remote Access.`,
			Description: `

This resource allows you to execute Check Point VPN Community Remote Access.

`,
			Keywords: []string{
				"management",
				"vpn",
				"community",
				"remote",
				"access",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name.`,
				},
				resource.Attribute{
					Name:        "gateways",
					Description: `(Optional) Collection of Gateway objects identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "user_groups",
					Description: `(Optional) Collection of User group objects identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_vpn_community_star",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Vpn Community Star.`,
			Description: `

This resource allows you to execute Check Point Vpn Community Star.

`,
			Keywords: []string{
				"management",
				"vpn",
				"community",
				"star",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "center_gateways",
					Description: `(Optional) Collection of Gateway objects representing center gateways identified by the name or UID.center_gateways blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "encryption_method",
					Description: `(Optional) The encryption method to be used.`,
				},
				resource.Attribute{
					Name:        "encryption_suite",
					Description: `(Optional) The encryption suite to be used.`,
				},
				resource.Attribute{
					Name:        "ike_phase_1",
					Description: `(Optional) Ike Phase 1 settings. Only applicable when the encryption-suite is set to [custom]. ike_phase_1 blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "ike_phase_2",
					Description: `(Optional) Ike Phase 2 settings. Only applicable when the encryption-suite is set to [custom]. ike_phase_2 blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "mesh_center_gateways",
					Description: `(Optional) Indicates whether the meshed community is in center.`,
				},
				resource.Attribute{
					Name:        "override_vpn_domains",
					Description: `(Optional) The Overrides VPN Domains of the participants GWs.override_vpn_domains blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "satellite_gateways",
					Description: `(Optional) Collection of Gateway objects representing satellite gateways identified by the name or UID.satellite_gateways blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "shared_secrets",
					Description: `(Optional) Shared secrets for external gateways.shared_secrets blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "tunnel_granularity",
					Description: `(Optional) VPN tunnel sharing option to be used.`,
				},
				resource.Attribute{
					Name:        "granular_encryptions",
					Description: `(Optional) VPN granular encryption settings. granular_encryptions blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "use_shared_secret",
					Description: `(Optional) Indicates whether the shared secret should be used for all external gateways.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. ` + "`" + `override_vpn_domains` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `(Optional) Participant gateway in override VPN domain identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "vpn_domain",
					Description: `(Optional) VPN domain network identified by the name or UID. ` + "`" + `shared_secrets` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "external_gateway",
					Description: `(Optional) External gateway identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "shared_secret",
					Description: `(Optional) Shared secret. granular_encryptions` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "internal-gateway",
					Description: `(Required) Internally managed Check Point gateway identified by name or UID, or 'Any' for all internal-gateways participants in this community.`,
				},
				resource.Attribute{
					Name:        "external_gateway",
					Description: `(Required) Externally managed or 3rd party gateway identified by name or UID.`,
				},
				resource.Attribute{
					Name:        "encryption_method",
					Description: `(Optional) The encryption method to be used: prefer ikev2 but support ikev1, ikev2 only, ikev1 for ipv4 and ikev2 for ipv6 only.`,
				},
				resource.Attribute{
					Name:        "encryption_suite",
					Description: `(Optional) The encryption suite to be used: suite-b-gcm-256, custom, vpn b, vpn a, suite-b-gcm-128.`,
				},
				resource.Attribute{
					Name:        "ike_phase_1",
					Description: `(Optional) Ike Phase 1 settings. Only applicable when the encryption-suite is set to [custom]. ike_phase_1 blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "ike_phase_2",
					Description: `(Optional) Ike Phase 2 settings. Only applicable when the encryption-suite is set to [custom]. ike_phase_2 blocks are documented below. ` + "`" + `ike_phase_1` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "data_integrity",
					Description: `(Optional) The hash algorithm to be used.`,
				},
				resource.Attribute{
					Name:        "diffie_hellman_group",
					Description: `(Optional) The Diffie-Hellman group to be used.`,
				},
				resource.Attribute{
					Name:        "encryption_algorithm",
					Description: `(Optional) The encryption algorithm to be used.`,
				},
				resource.Attribute{
					Name:        "ike_p1_rekey_time",
					Description: `(Optional) Indicates the time interval for IKE phase 1 renegotiation. ` + "`" + `ike_phase_2` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "data_integrity",
					Description: `(Optional) The hash algorithm to be used.`,
				},
				resource.Attribute{
					Name:        "encryption_algorithm",
					Description: `(Optional) The encryption algorithm to be used.`,
				},
				resource.Attribute{
					Name:        "ike_p2_use_pfs",
					Description: `(Optional) Indicates whether Perfect Forward Secrecy (PFS) is being used for IKE phase 2.`,
				},
				resource.Attribute{
					Name:        "ike_p2_pfs_dh_grp",
					Description: `(Optional) The Diffie-Hellman group to be used: group-1, group-2, group-5, group-14, group-15, group-16, group-17, group-18, group-19, group-20, group-24.`,
				},
				resource.Attribute{
					Name:        "ike_p2_rekey_time",
					Description: `(Optional) Indicates the time interval for IKE phase 2 renegotiation.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_where_used",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Where Used.`,
			Description: `

This command resource allows you to execute Check Point Where Used.

`,
			Keywords: []string{
				"management",
				"where",
				"used",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "dereference_group_members",
					Description: `(Optional) Indicates whether to dereference "members" field by details level for every object in reply.`,
				},
				resource.Attribute{
					Name:        "show_membership",
					Description: `(Optional) Indicates whether to calculate and show "groups" field for every object in reply.`,
				},
				resource.Attribute{
					Name:        "indirect",
					Description: `(Optional) Search for indirect usage.`,
				},
				resource.Attribute{
					Name:        "indirect_max_depth",
					Description: `(Optional) Maximum nesting level during indirect usage search. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_wildcard",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Wildcard.`,
			Description: `

This resource allows you to execute Check Point Wildcard.

`,
			Keywords: []string{
				"management",
				"wildcard",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `(Optional) IPv4 address.`,
				},
				resource.Attribute{
					Name:        "ipv4_mask_wildcard",
					Description: `(Optional) IPv4 mask wildcard.`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `(Optional) IPv6 address.`,
				},
				resource.Attribute{
					Name:        "ipv6_mask_wildcard",
					Description: `(Optional) IPv6 mask wildcard.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_physical_interface",
			Category:         "GAIA Resources",
			ShortDescription: `This resource allows you to set a Physical interface.`,
			Description: `

This resource allows you to set a Physical interface.

`,
			Keywords: []string{
				"gaia",
				"physical",
				"interface",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Interface name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Interface state.`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `(Optional) IPv4 address to set for the interface.`,
				},
				resource.Attribute{
					Name:        "ipv4_mask_length",
					Description: `(Optional) Interface IPv4 address mask length.`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `(Optional) IPv6 address to set for the interface.`,
				},
				resource.Attribute{
					Name:        "ipv6_mask_length",
					Description: `(Optional) Interface IPv6 address mask length.`,
				},
				resource.Attribute{
					Name:        "ipv6_autoconfig",
					Description: `(Optional) Configure IPv6 auto-configuration true/false.`,
				},
				resource.Attribute{
					Name:        "mac_addr",
					Description: `(Optional) Configure hardware address.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `(Optional) Interface Mtu.`,
				},
				resource.Attribute{
					Name:        "rx_ringsize",
					Description: `(Optional) Set receive buffer size for the interface.`,
				},
				resource.Attribute{
					Name:        "tx_ringsize",
					Description: `(Optional) Set transmit buffer size for the interface.`,
				},
				resource.Attribute{
					Name:        "monitor_mode",
					Description: `(Optional) Set monitor mode for the interface true/false.`,
				},
				resource.Attribute{
					Name:        "auto_negotiation",
					Description: `(Optional) Configure auto-negotiation. Activating Auto-Negotiation will skip the speed and duplex configuration.`,
				},
				resource.Attribute{
					Name:        "duplex",
					Description: `(Optional) duplex for the interface. Duplex is not relevant when 'auto_negotiation' is enabled.`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `(Optional) Interface link speed. Speed is not relevant when 'auto_negotiation' is enabled.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) interface Comments.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_put_file",
			Category:         "GAIA Resources",
			ShortDescription: `This resource allows you to add a new file to a Check Point machine.`,
			Description: `

This resource allows you to add a new file to a Check Point machine.

`,
			Keywords: []string{
				"gaia",
				"put",
				"file",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "file_name",
					Description: `(Required) Filename include the desired path. The file will be created in the user home directory if the full path wasn't provided.`,
				},
				resource.Attribute{
					Name:        "text_content",
					Description: `(Required) Content to add to the new file.`,
				},
				resource.Attribute{
					Name:        "override",
					Description: `(Optional) If the file already exists, indicates whether to overwrite it.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"checkpoint_hostname":                                                  0,
		"checkpoint_management_access_layer":                                   1,
		"checkpoint_management_access_point_name":                              2,
		"checkpoint_management_access_role":                                    3,
		"checkpoint_management_access_rule":                                    4,
		"checkpoint_management_access_section":                                 5,
		"checkpoint_management_aci_data_center_server":                         6,
		"checkpoint_management_add_api_key":                                    7,
		"checkpoint_management_add_data_center_object":                         8,
		"checkpoint_management_add_repository_package":                         9,
		"checkpoint_management_add_threat_protections":                         10,
		"checkpoint_management_add_updatable_object":                           11,
		"checkpoint_management_address_range":                                  12,
		"checkpoint_management_administrator":                                  13,
		"checkpoint_management_application_site":                               14,
		"checkpoint_management_application_site_category":                      15,
		"checkpoint_management_application_site_group":                         16,
		"checkpoint_management_assign_global_assignment":                       17,
		"checkpoint_management_aws_data_center_server":                         18,
		"checkpoint_management_azure_ad":                                       19,
		"checkpoint_management_azure_data_center_server":                       20,
		"checkpoint_management_backup_domain":                                  21,
		"checkpoint_management_check_network_feed":                             22,
		"checkpoint_management_check_threat_ioc_feed":                          23,
		"checkpoint_management_checkpoint_host":                                24,
		"checkpoint_management_command_abort_get_interfaces":                   25,
		"checkpoint_management_command_get_interfaces":                         26,
		"checkpoint_management_connect_cloud_services":                         27,
		"checkpoint_management_data_center_query":                              28,
		"checkpoint_management_delete_api_key":                                 29,
		"checkpoint_management_delete_data_center_object":                      30,
		"checkpoint_management_delete_repository_package":                      31,
		"checkpoint_management_delete_threat_protections":                      32,
		"checkpoint_management_delete_updatable_object":                        33,
		"checkpoint_management_discard":                                        34,
		"checkpoint_management_disconnect":                                     35,
		"checkpoint_management_disconnect_cloud_services":                      36,
		"checkpoint_management_dns_domain":                                     37,
		"checkpoint_management_domain":                                         38,
		"checkpoint_management_domain_permissions_profile":                     39,
		"checkpoint_management_dynamic_global_network_object":                  40,
		"checkpoint_management_dynamic_object":                                 41,
		"checkpoint_management_exception_group":                                42,
		"checkpoint_management_export":                                         43,
		"checkpoint_management_gaia_best_practice":                             44,
		"checkpoint_management_gcp_data_center_server":                         45,
		"checkpoint_management_generic_data_center_server":                     46,
		"checkpoint_management_get_attachment":                                 47,
		"checkpoint_management_get_platform":                                   48,
		"checkpoint_management_global_assignment":                              49,
		"checkpoint_management_group":                                          50,
		"checkpoint_management_group_with_exclusion":                           51,
		"checkpoint_management_gsn_handover_group":                             52,
		"checkpoint_management_ha_full_sync":                                   53,
		"checkpoint_management_host":                                           54,
		"checkpoint_management_https_layer":                                    55,
		"checkpoint_management_https_rule":                                     56,
		"checkpoint_management_https_section":                                  57,
		"checkpoint_management_identity_tag":                                   58,
		"checkpoint_management_idp_administrator_group":                        59,
		"checkpoint_management_install_database":                               60,
		"checkpoint_management_install_lsm_policy":                             61,
		"checkpoint_management_install_lsm_settings":                           62,
		"checkpoint_management_install_policy":                                 63,
		"checkpoint_management_install_software_package":                       64,
		"checkpoint_management_interoperable_device":                           65,
		"checkpoint_management_ise_data_center_server":                         66,
		"checkpoint_management_keepalive":                                      67,
		"checkpoint_management_kubernetes_data_center_server":                  68,
		"checkpoint_management_login":                                          69,
		"checkpoint_management_logout":                                         70,
		"checkpoint_management_lsm_run_script":                                 71,
		"checkpoint_management_lsv_profile":                                    72,
		"checkpoint_management_md_permissions_profile":                         73,
		"checkpoint_management_mds":                                            74,
		"checkpoint_management_migrate_export_domain":                          75,
		"checkpoint_management_migrate_import_domain":                          76,
		"checkpoint_management_multicast_address_range":                        77,
		"checkpoint_management_nat_rule":                                       78,
		"checkpoint_management_nat_section":                                    79,
		"checkpoint_management_network":                                        80,
		"checkpoint_management_network_feed":                                   81,
		"checkpoint_management_nuage_data_center_server":                       82,
		"checkpoint_management_nutanix_data_center_server":                     83,
		"checkpoint_management_openstack_data_center_server":                   84,
		"checkpoint_management_opsec_application":                              85,
		"checkpoint_management_oracle_cloud_data_center_server":                86,
		"checkpoint_management_package":                                        87,
		"checkpoint_management_publish":                                        88,
		"checkpoint_management_put_file":                                       89,
		"checkpoint_management_radius_group":                                   90,
		"checkpoint_management_radius_server":                                  91,
		"checkpoint_management_repository_script":                              92,
		"checkpoint_management_reset_sic":                                      93,
		"checkpoint_management_restore_domain":                                 94,
		"checkpoint_management_revert_to_revision":                             95,
		"checkpoint_management_run_ips_update":                                 96,
		"checkpoint_management_run_script":                                     97,
		"checkpoint_management_run_threat_emulation_file_types_offline_update": 98,
		"checkpoint_management_security_zone":                                  99,
		"checkpoint_management_service_citrix_tcp":                             100,
		"checkpoint_management_service_compound_tcp":                           101,
		"checkpoint_management_service_dce_rpc":                                102,
		"checkpoint_management_service_group":                                  103,
		"checkpoint_management_service_icmp":                                   104,
		"checkpoint_management_service_icmp6":                                  105,
		"checkpoint_management_service_other":                                  106,
		"checkpoint_management_service_rpc":                                    107,
		"checkpoint_management_service_sctp":                                   108,
		"checkpoint_management_service_tcp":                                    109,
		"checkpoint_management_service_udp":                                    110,
		"checkpoint_management_set_api_settings":                               111,
		"checkpoint_management_set_automatic_purge":                            112,
		"checkpoint_management_set_global_domain":                              113,
		"checkpoint_management_set_ha_state":                                   114,
		"checkpoint_management_set_idp_default_assignment":                     115,
		"checkpoint_management_set_ips_update_schedule":                        116,
		"checkpoint_management_set_login_message":                              117,
		"checkpoint_management_set_threat_protection":                          118,
		"checkpoint_management_simple_cluster":                                 119,
		"checkpoint_management_simple_gateway":                                 120,
		"checkpoint_management_smtp_server":                                    121,
		"checkpoint_management_tacacs_group":                                   122,
		"checkpoint_management_tacacs_server":                                  123,
		"checkpoint_management_tag":                                            124,
		"checkpoint_management_test_sic_status":                                125,
		"checkpoint_management_threat_exception":                               126,
		"checkpoint_management_threat_indicator":                               127,
		"checkpoint_management_threat_ioc_feed":                                128,
		"checkpoint_management_threat_layer":                                   129,
		"checkpoint_management_threat_profile":                                 130,
		"checkpoint_management_threat_rule":                                    131,
		"checkpoint_management_time":                                           132,
		"checkpoint_management_time_group":                                     133,
		"checkpoint_management_trusted_client":                                 134,
		"checkpoint_management_uninstall_software_package":                     135,
		"checkpoint_management_unlock_administrator":                           136,
		"checkpoint_management_update_provisioned_satellites":                  137,
		"checkpoint_management_update_updatable_objects_repository_content":    138,
		"checkpoint_management_user":                                           139,
		"checkpoint_management_user_group":                                     140,
		"checkpoint_management_user_template":                                  141,
		"checkpoint_management_verify_policy":                                  142,
		"checkpoint_management_verify_revert":                                  143,
		"checkpoint_management_verify_software_package":                        144,
		"checkpoint_management_vmware_data_center_server":                      145,
		"checkpoint_management_vpn_community_meshed":                           146,
		"checkpoint_management_vpn_community_remote_access":                    147,
		"checkpoint_management_vpn_community_star":                             148,
		"checkpoint_management_where_used":                                     149,
		"checkpoint_management_wildcard":                                       150,
		"checkpoint_physical_interface":                                        151,
		"checkpoint_put_file":                                                  152,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
