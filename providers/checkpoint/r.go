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
					Description: `(Optional) Active Directory name or UID or Identity Tag.`,
				},
				resource.Attribute{
					Name:        "selection",
					Description: `(Optional) Name or UID of an object selected from source.selection blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "base_dn",
					Description: `(Optional) When source is "Active Directory" use "base-dn" to refine the query in AD database. ` + "`" + `users` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) Active Directory name or UID or Identity Tag or Internal User Groups or LDAP groups or Guests.`,
				},
				resource.Attribute{
					Name:        "selection",
					Description: `(Optional) Name or UID of an object selected from source.selection blocks are documented below.`,
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
					Description: `(Required) Rule name.`,
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
					Description: `(Optional) N/A. ## Import ` + "`" + `checkpoint_management_access_rule` + "`" + ` can be imported by using the following format: LAYER_NAME;RULE_UID ` + "`" + `` + "`" + `` + "`" + ` $ terraform import checkpoint_management_access_rule.example Network;9423d36f-2d66-4754-b9e2-e9f4493751d3 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Required) Position in the rulebase.`,
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
					Description: `(Optional) Whether to match sub-domains in addition to the domain itself.`,
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
					Description: `(Optional) Add rule at the bottom of the rulebase. ## Import ` + "`" + `checkpoint_management_nat_rule` + "`" + ` can be imported by using the following format: PACKAGE_NAME;RULE_UID ` + "`" + `` + "`" + `` + "`" + ` $ terraform import checkpoint_management_nat_rule.example Standard;9423d36f-2d66-4754-b9e2-e9f4493751d3 ` + "`" + `` + "`" + `` + "`" + ``,
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
			Description: `

This resource allows you to execute Check Point Simple Cluster.

`,
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
					Name:        "ipv4_address",
					Description: `(Optional) IPv4 address.`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `(Optional) IPv6 address.`,
				},
				resource.Attribute{
					Name:        "cluster_mode",
					Description: `(Optional) Cluster mode.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `(Optional) Cluster interfaces. interfaces blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `(Optional) Cluster members. members blocks are documented below.`,
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
					Name:        "content_awareness",
					Description: `(Optional) Content Awareness blade enabled.`,
				},
				resource.Attribute{
					Name:        "data_awareness",
					Description: `(Optional) Data Awareness blade enabled.`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `(Optional) Intrusion Prevention System blade enabled.`,
				},
				resource.Attribute{
					Name:        "threat_emulation",
					Description: `(Optional) Threat Emulation blade enabled.`,
				},
				resource.Attribute{
					Name:        "url_filtering",
					Description: `(Optional) URL Filtering blade enabled.`,
				},
				resource.Attribute{
					Name:        "firewall",
					Description: `(Optional) Firewall blade enabled.`,
				},
				resource.Attribute{
					Name:        "firewall_settings",
					Description: `(Optional) Firewall settings. firewall_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "vpn",
					Description: `(Optional) VPN blade enabled.`,
				},
				resource.Attribute{
					Name:        "vpn_settings",
					Description: `(Optional) Cluster VPN settings. vpn_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "dynamic_ip",
					Description: `(Computed) Dynamic IP address.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) Cluster platform version.`,
				},
				resource.Attribute{
					Name:        "os_name",
					Description: `(Optional) Cluster Operating system name.`,
				},
				resource.Attribute{
					Name:        "hardware",
					Description: `(Optional) Cluster platform hardware name.`,
				},
				resource.Attribute{
					Name:        "one_time_password",
					Description: `(Optional) Secure Internal Communication one time password.`,
				},
				resource.Attribute{
					Name:        "sic_name",
					Description: `(Computed) Secure Internal Communication name.`,
				},
				resource.Attribute{
					Name:        "sic_state",
					Description: `(Computed) Secure Internal Communication state.`,
				},
				resource.Attribute{
					Name:        "save_logs_locally",
					Description: `(Optional) Enable save logs locally.`,
				},
				resource.Attribute{
					Name:        "send_alerts_to_server",
					Description: `(Optional) Collection of Server(s) to send alerts to identified by the name.`,
				},
				resource.Attribute{
					Name:        "send_logs_to_backup_server",
					Description: `(Optional) Collection of Backup server(s) to send logs to identified by the name.`,
				},
				resource.Attribute{
					Name:        "send_logs_to_server",
					Description: `(Optional) Collection of Server(s) to send logs to identified by the name.`,
				},
				resource.Attribute{
					Name:        "logs_settings",
					Description: `(Optional) Logs settings. logs_settings blocks are documented below.`,
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
					Name:        "tags",
					Description: `(Optional) Collection of tags identified by name.`,
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
					Description: `(Optional) IPv4 network mask length.`,
				},
				resource.Attribute{
					Name:        "anti_spoofing",
					Description: `(Optional) Anti spoofing.`,
				},
				resource.Attribute{
					Name:        "anti_spoofing_settings",
					Description: `(Optional) Anti spoofing settings. anti_spoofing_settings blocks are documented below.`,
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
					Description: `(Optional) Security zone.`,
				},
				resource.Attribute{
					Name:        "security_zone_settings",
					Description: `(Optional) Security zone settings. security_zone_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "topology",
					Description: `(Optional) Topology.`,
				},
				resource.Attribute{
					Name:        "topology_settings",
					Description: `(Optional) Topology settings. topology_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "topology_automatic_calculation",
					Description: `(Computed) Shows the automatic topology calculation..`,
				},
				resource.Attribute{
					Name:        "topology",
					Description: `(Optional) Topology.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string. ` + "`" + `anti_spoofing_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) If packets will be rejected (the Prevent option) or whether the packets will be monitored (the Detect option). ` + "`" + `security_zone_settings` + "`" + ` supports the following:`,
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
					Name:        "ip_address_behind_this_interface",
					Description: `(Optional) Ip address behind this interface.`,
				},
				resource.Attribute{
					Name:        "specific_network",
					Description: `(Optional) Network behind this interface. ` + "`" + `members` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name. Should be unique in the domain..`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) IPv4 or IPv6 address.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `(Optional) Cluster Member network interfaces. interfaces blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "one_time_password",
					Description: `(Optional) Secure Internal Communication one time password.`,
				},
				resource.Attribute{
					Name:        "sic_name",
					Description: `(Computed) Secure Internal Communication name.`,
				},
				resource.Attribute{
					Name:        "sic_message",
					Description: `(Computed) Secure Internal Communication state. ` + "`" + `interfaces` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Interface name.`,
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
					Description: `(Optional) IPv6 network mask length. ` + "`" + `firewall_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "auto_calculate_connections_hash_table_size_and_memory_pool",
					Description: `(Optional) Auto calculate connections hash table size and memory pool.`,
				},
				resource.Attribute{
					Name:        "auto_maximum_limit_for_concurrent_connections",
					Description: `(Optional) Auto maximum limit for concurrent connections.`,
				},
				resource.Attribute{
					Name:        "connections_hash_size",
					Description: `(Optional) Connections hash size.`,
				},
				resource.Attribute{
					Name:        "maximum_limit_for_concurrent_connections",
					Description: `(Optional) Maximum limit for concurrent connections.`,
				},
				resource.Attribute{
					Name:        "maximum_memory_pool_size",
					Description: `(Optional) Maximum memory pool size.`,
				},
				resource.Attribute{
					Name:        "memory_pool_size",
					Description: `(Optional) Memory pool size. ` + "`" + `vpn_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "authentication",
					Description: `(Optional) authentication blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "link_selection",
					Description: `(Optional) Link selection blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "maximum_concurrent_ike_negotiations",
					Description: `(Optional) Maximum concurrent ike negotiations.`,
				},
				resource.Attribute{
					Name:        "maximum_concurrent_tunnels",
					Description: `(Optional) Maximum concurrent tunnels.`,
				},
				resource.Attribute{
					Name:        "office_mode",
					Description: `(Optional) Office Mode. Notation Wide Impact - Office Mode apply IPSec VPN Software Blade clients and to the Mobile Access Software Blade clients. office_mode blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "remote_access",
					Description: `(Optional) remote_access blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "vpn_domain",
					Description: `(Optional) Gateway VPN domain identified by the name.`,
				},
				resource.Attribute{
					Name:        "vpn_domain_type",
					Description: `(Optional) Gateway VPN domain type. ` + "`" + `authentication` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "authentication_clients",
					Description: `(Optional) Collection of VPN Authentication clients identified by the name. ` + "`" + `link_selection` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "ip_selection",
					Description: `(Optional) IP selection.`,
				},
				resource.Attribute{
					Name:        "dns_resolving_hostname",
					Description: `(Optional) DNS Resolving Hostname. Must be set when "ip-selection" was selected to be "dns-resolving-from-hostname".`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) IP Address. Must be set when "ip-selection" was selected to be "use-selected-address-from-topology" or "use-statically-nated-ip". ` + "`" + `office_mode` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) Office Mode Permissions. When selected to be "off", all the other definitions are irrelevant.`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `(Optional) Group. Identified by name. Must be set when "office-mode-permissions" was selected to be "group".`,
				},
				resource.Attribute{
					Name:        "allocate_ip_address_from",
					Description: `(Optional) Allocate IP address Method. Allocate IP address by sequentially trying the given methods until success. allocate_ip_address_from blocks are documented below.`,
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
					Description: `(Optional) Additional IP Addresses for Anti-Spoofing. Identified by name. Must be set when "perform-anti-spoofings" is true. ` + "`" + `allocate_ip_address_from` + "`" + ` supports the following:`,
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
					Description: `(Optional) Manual Network. Identified by name. Must be set when "allocate-method" was selected to be "manual".`,
				},
				resource.Attribute{
					Name:        "dhcp_server",
					Description: `(Optional) DHCP Server. Identified by name. Must be set when "allocate-method" was selected to be "automatic".`,
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
					Description: `(Optional) This configuration applies to all Office Mode methods except Automatic (using DHCP) and ipassignment.conf entries which contain this data. optional_parameters blocks are documented below. ` + "`" + `optional_parameters` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "use_primary_dns_server",
					Description: `(Optional) Use Primary DNS Server.`,
				},
				resource.Attribute{
					Name:        "primary_dns_server",
					Description: `(Optional) Primary DNS Server. Identified by name. Must be set when "use-primary-dns-server" is true and can not be set when "use-primary-dns-server" is false.`,
				},
				resource.Attribute{
					Name:        "use_first_backup_dns_server",
					Description: `(Optional) Use First Backup DNS Server.`,
				},
				resource.Attribute{
					Name:        "first_backup_dns_server",
					Description: `(Optional) First Backup DNS Server. Identified by name. Must be set when "use-first-backup-dns-server" is true and can not be set when "use-first-backup-dns-server" is false.`,
				},
				resource.Attribute{
					Name:        "use_second_backup_dns_server",
					Description: `(Optional) Use Second Backup DNS Server.`,
				},
				resource.Attribute{
					Name:        "second_backup_dns_server",
					Description: `(Optional) Second Backup DNS Server. Identified by name. Must be set when "use-second-backup-dns-server" is true and can not be set when "use-second-backup-dns-server" is false.`,
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
					Description: `(Optional) Primary WINS Server. Identified by name. Must be set when "use-primary-wins-server" is true and can not be set when "use-primary-wins-server" is false.`,
				},
				resource.Attribute{
					Name:        "use_first_backup_wins_server",
					Description: `(Optional) Use First Backup WINS Server.`,
				},
				resource.Attribute{
					Name:        "first_backup_wins_server",
					Description: `(Optional) First Backup WINS Server. Identified by name. Must be set when "use-first-backup-wins-server" is true and can not be set when "use-first-backup-wins-server" is false.`,
				},
				resource.Attribute{
					Name:        "use_second_backup_wins_server",
					Description: `(Optional) Use Second Backup WINS Server.`,
				},
				resource.Attribute{
					Name:        "second_backup_wins_server",
					Description: `(Optional) Second Backup WINS Server. Identified by name. Must be set when "use-second-backup-wins-server" is true and can not be set when "use-second-backup-wins-server" is false.`,
				},
				resource.Attribute{
					Name:        "ip_lease_duration",
					Description: `(Optional) IP Lease Duration in Minutes. The value must be in the range 2-32767. ` + "`" + `remote_access` + "`" + ` supports the following:`,
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
					Description: `(Optional) Allocated NAT traversal UDP service. Identified by name. Must be set when "support-nat-traversal-mechanism" is true.`,
				},
				resource.Attribute{
					Name:        "support_visitor_mode",
					Description: `(Optional) Support Visitor Mode.`,
				},
				resource.Attribute{
					Name:        "visitor_mode_service",
					Description: `(Optional) TCP Service for Visitor Mode. Identified by name. Must be set when "support-visitor-mode" is true.`,
				},
				resource.Attribute{
					Name:        "visitor_mode_interface",
					Description: `(Optional) Interface for Visitor Mode. Must be set when "support-visitor-mode" is true. Insert IPV4 Address of existing interface or "All IPs" when you want all interfaces.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_simple_gateway",
			Category:         "Management Resources",
			ShortDescription: `This resource allows you to execute Check Point Simple Gateway.`,
			Description: `

This resource allows you to execute Check Point Simple Gateway.

`,
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
					Name:        "ipv4_address",
					Description: `(Optional) IPv4 address.`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `(Optional) IPv6 address.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `(Optional) Gateway interfaces. interfaces blocks are documented below.`,
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
					Name:        "content_awareness",
					Description: `(Optional) Content Awareness blade enabled.`,
				},
				resource.Attribute{
					Name:        "icap_server",
					Description: `(Optional) ICAP Server enabled.`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `(Optional) Intrusion Prevention System blade enabled.`,
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
					Name:        "url_filtering",
					Description: `(Optional) URL Filtering blade enabled.`,
				},
				resource.Attribute{
					Name:        "firewall",
					Description: `(Optional) Firewall blade enabled.`,
				},
				resource.Attribute{
					Name:        "firewall_settings",
					Description: `(Optional) Firewall settings. firewall_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "vpn",
					Description: `(Optional) VPN blade enabled.`,
				},
				resource.Attribute{
					Name:        "vpn_settings",
					Description: `(Optional) Gateway VPN settings. vpn_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "dynamic_ip",
					Description: `(Computed) Dynamic IP address.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) Gateway platform version.`,
				},
				resource.Attribute{
					Name:        "os_name",
					Description: `(Optional) Operating system name.`,
				},
				resource.Attribute{
					Name:        "hardware",
					Description: `(Computed) Gateway platform hardware name.`,
				},
				resource.Attribute{
					Name:        "one_time_password",
					Description: `(Optional) Secure internal connection one time password.`,
				},
				resource.Attribute{
					Name:        "sic_name",
					Description: `(Computed) Secure Internal Communication name.`,
				},
				resource.Attribute{
					Name:        "sic_state",
					Description: `(Computed) Secure Internal Communication state.`,
				},
				resource.Attribute{
					Name:        "save_logs_locally",
					Description: `(Optional) Enable save logs locally.`,
				},
				resource.Attribute{
					Name:        "send_alerts_to_server",
					Description: `(Optional) Collection of Server(s) to send alerts to identified by the name.`,
				},
				resource.Attribute{
					Name:        "send_logs_to_backup_server",
					Description: `(Optional) Collection of Backup server(s) to send logs to identified by the name.`,
				},
				resource.Attribute{
					Name:        "send_logs_to_server",
					Description: `(Optional) Collection of Server(s) to send logs to identified by the name.`,
				},
				resource.Attribute{
					Name:        "logs_settings",
					Description: `(Optional) Logs settings. logs_settings blocks are documented below.`,
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
					Name:        "tags",
					Description: `(Optional) Collection of tags identified by name.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `(Optional) Apply changes ignoring warnings. ` + "`" + `interfaces` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Interface name.`,
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
					Description: `(Optional) Anti spoofing.`,
				},
				resource.Attribute{
					Name:        "anti_spoofing_settings",
					Description: `(Optional) Anti spoofing settings. anti_spoofing_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "security_zone",
					Description: `(Optional) Security zone.`,
				},
				resource.Attribute{
					Name:        "security_zone_settings",
					Description: `(Optional) Security zone settings. security_zone_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "topology",
					Description: `(Optional) Topology.`,
				},
				resource.Attribute{
					Name:        "topology_settings",
					Description: `(Optional) Topology settings. topology_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "topology_automatic_calculation",
					Description: `(Computed) Shows the automatic topology calculation..`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments string. ` + "`" + `anti_spoofing_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) If packets will be rejected (the Prevent option) or whether the packets will be monitored (the Detect option). ` + "`" + `security_zone_settings` + "`" + ` supports the following:`,
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
					Name:        "ip_address_behind_this_interface",
					Description: `(Optional) Ip address behind this interface.`,
				},
				resource.Attribute{
					Name:        "specific_network",
					Description: `(Optional) Network behind this interface. ` + "`" + `firewall_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "auto_calculate_connections_hash_table_size_and_memory_pool",
					Description: `(Optional) Auto calculate connections hash table size and memory pool.`,
				},
				resource.Attribute{
					Name:        "auto_maximum_limit_for_concurrent_connections",
					Description: `(Optional) Auto maximum limit for concurrent connections.`,
				},
				resource.Attribute{
					Name:        "connections_hash_size",
					Description: `(Optional) Connections hash size.`,
				},
				resource.Attribute{
					Name:        "maximum_limit_for_concurrent_connections",
					Description: `(Optional) Maximum limit for concurrent connections.`,
				},
				resource.Attribute{
					Name:        "maximum_memory_pool_size",
					Description: `(Optional) Maximum memory pool size.`,
				},
				resource.Attribute{
					Name:        "memory_pool_size",
					Description: `(Optional) Memory pool size. ` + "`" + `vpn_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "authentication",
					Description: `(Optional) authentication blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "link_selection",
					Description: `(Optional) Link selection blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "maximum_concurrent_ike_negotiations",
					Description: `(Optional) Maximum concurrent ike negotiations.`,
				},
				resource.Attribute{
					Name:        "maximum_concurrent_tunnels",
					Description: `(Optional) Maximum concurrent tunnels.`,
				},
				resource.Attribute{
					Name:        "office_mode",
					Description: `(Optional) Office Mode. Notation Wide Impact - Office Mode apply IPSec VPN Software Blade clients and to the Mobile Access Software Blade clients. office_mode blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "remote_access",
					Description: `(Optional) remote_access blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "vpn_domain",
					Description: `(Optional) Gateway VPN domain identified by the name.`,
				},
				resource.Attribute{
					Name:        "vpn_domain_type",
					Description: `(Optional) Gateway VPN domain type. ` + "`" + `authentication` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "authentication_clients",
					Description: `(Optional) Collection of VPN Authentication clients identified by the name. ` + "`" + `link_selection` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "ip_selection",
					Description: `(Optional) IP selection.`,
				},
				resource.Attribute{
					Name:        "dns_resolving_hostname",
					Description: `(Optional) DNS Resolving Hostname. Must be set when "ip-selection" was selected to be "dns-resolving-from-hostname".`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) IP Address. Must be set when "ip-selection" was selected to be "use-selected-address-from-topology" or "use-statically-nated-ip". ` + "`" + `office_mode` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) Office Mode Permissions. When selected to be "off", all the other definitions are irrelevant.`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `(Optional) Group. Identified by name. Must be set when "office-mode-permissions" was selected to be "group".`,
				},
				resource.Attribute{
					Name:        "allocate_ip_address_from",
					Description: `(Optional) Allocate IP address Method. Allocate IP address by sequentially trying the given methods until success. allocate_ip_address_from blocks are documented below.`,
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
					Description: `(Optional) Additional IP Addresses for Anti-Spoofing. Identified by name. Must be set when "perform-anti-spoofings" is true. ` + "`" + `allocate_ip_address_from` + "`" + ` supports the following:`,
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
					Description: `(Optional) Manual Network. Identified by name. Must be set when "allocate-method" was selected to be "manual".`,
				},
				resource.Attribute{
					Name:        "dhcp_server",
					Description: `(Optional) DHCP Server. Identified by name. Must be set when "allocate-method" was selected to be "automatic".`,
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
					Description: `(Optional) This configuration applies to all Office Mode methods except Automatic (using DHCP) and ipassignment.conf entries which contain this data. optional_parameters blocks are documented below. ` + "`" + `optional_parameters` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "use_primary_dns_server",
					Description: `(Optional) Use Primary DNS Server.`,
				},
				resource.Attribute{
					Name:        "primary_dns_server",
					Description: `(Optional) Primary DNS Server. Identified by name. Must be set when "use-primary-dns-server" is true and can not be set when "use-primary-dns-server" is false.`,
				},
				resource.Attribute{
					Name:        "use_first_backup_dns_server",
					Description: `(Optional) Use First Backup DNS Server.`,
				},
				resource.Attribute{
					Name:        "first_backup_dns_server",
					Description: `(Optional) First Backup DNS Server. Identified by name. Must be set when "use-first-backup-dns-server" is true and can not be set when "use-first-backup-dns-server" is false.`,
				},
				resource.Attribute{
					Name:        "use_second_backup_dns_server",
					Description: `(Optional) Use Second Backup DNS Server.`,
				},
				resource.Attribute{
					Name:        "second_backup_dns_server",
					Description: `(Optional) Second Backup DNS Server. Identified by name. Must be set when "use-second-backup-dns-server" is true and can not be set when "use-second-backup-dns-server" is false.`,
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
					Description: `(Optional) Primary WINS Server. Identified by name. Must be set when "use-primary-wins-server" is true and can not be set when "use-primary-wins-server" is false.`,
				},
				resource.Attribute{
					Name:        "use_first_backup_wins_server",
					Description: `(Optional) Use First Backup WINS Server.`,
				},
				resource.Attribute{
					Name:        "first_backup_wins_server",
					Description: `(Optional) First Backup WINS Server. Identified by name. Must be set when "use-first-backup-wins-server" is true and can not be set when "use-first-backup-wins-server" is false.`,
				},
				resource.Attribute{
					Name:        "use_second_backup_wins_server",
					Description: `(Optional) Use Second Backup WINS Server.`,
				},
				resource.Attribute{
					Name:        "second_backup_wins_server",
					Description: `(Optional) Second Backup WINS Server. Identified by name. Must be set when "use-second-backup-wins-server" is true and can not be set when "use-second-backup-wins-server" is false.`,
				},
				resource.Attribute{
					Name:        "ip_lease_duration",
					Description: `(Optional) IP Lease Duration in Minutes. The value must be in the range 2-32767. ` + "`" + `remote_access` + "`" + ` supports the following:`,
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
					Description: `(Optional) Allocated NAT traversal UDP service. Identified by name. Must be set when "support-nat-traversal-mechanism" is true.`,
				},
				resource.Attribute{
					Name:        "support_visitor_mode",
					Description: `(Optional) Support Visitor Mode.`,
				},
				resource.Attribute{
					Name:        "visitor_mode_service",
					Description: `(Optional) TCP Service for Visitor Mode. Identified by name. Must be set when "support-visitor-mode" is true.`,
				},
				resource.Attribute{
					Name:        "visitor_mode_interface",
					Description: `(Optional) Interface for Visitor Mode. Must be set when "support-visitor-mode" is true. Insert IPV4 Address of existing interface or "All IPs" when you want all interfaces. ` + "`" + `logs_settings` + "`" + ` supports the following:`,
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
					Name:        "stop_logging_when_free_disk_space_below",
					Description: `(Optional) Enable stop logging when free disk space below.`,
				},
				resource.Attribute{
					Name:        "stop_logging_when_free_disk_space_below_threshold",
					Description: `(Optional) Stop logging when free disk space below threshold.`,
				},
				resource.Attribute{
					Name:        "turn_on_qos_logging",
					Description: `(Optional) Enable turn on qos loggig.`,
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
					Description: `(Optional) Packet capture. ## Import ` + "`" + `checkpoint_management_threat_rule` + "`" + ` can be imported by using the following format: LAYER_NAME;RULE_UID ` + "`" + `` + "`" + `` + "`" + ` $ terraform import checkpoint_management_threat_rule.example Layer_Name;9423d36f-2d66-4754-b9e2-e9f4493751d3 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Optional) Ike Phase 1 settings. Only applicable when the encryption-suite is set to [custom].ike_phase_1 blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "ike_phase_2",
					Description: `(Optional) Ike Phase 2 settings. Only applicable when the encryption-suite is set to [custom].ike_phase_2 blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "override_vpn_domains",
					Description: `(Optional) The Overrides VPN Domains of the participants GWs.override_vpn_domains blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "shared_secrets",
					Description: `(Optional) Shared secrets for external gateways.shared_secrets blocks are documented below.`,
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
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. ` + "`" + `ike_phase_1` + "`" + ` supports the following:`,
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
					Description: `(Optional) The encryption algorithm to be used. ` + "`" + `ike_phase_2` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "data_integrity",
					Description: `(Optional) The hash algorithm to be used.`,
				},
				resource.Attribute{
					Name:        "encryption_algorithm",
					Description: `(Optional) The encryption algorithm to be used. ` + "`" + `override_vpn_domains` + "`" + ` supports the following:`,
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
					Description: `(Optional) Shared secret.`,
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
					Description: `(Optional) Ike Phase 1 settings. Only applicable when the encryption-suite is set to [custom].ike_phase_1 blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "ike_phase_2",
					Description: `(Optional) Ike Phase 2 settings. Only applicable when the encryption-suite is set to [custom].ike_phase_2 blocks are documented below.`,
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
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. ` + "`" + `ike_phase_1` + "`" + ` supports the following:`,
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
					Description: `(Optional) The encryption algorithm to be used. ` + "`" + `ike_phase_2` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "data_integrity",
					Description: `(Optional) The hash algorithm to be used.`,
				},
				resource.Attribute{
					Name:        "encryption_algorithm",
					Description: `(Optional) The encryption algorithm to be used. ` + "`" + `override_vpn_domains` + "`" + ` supports the following:`,
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
					Description: `(Optional) Shared secret.`,
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
		"checkpoint_management_add_api_key":                                    6,
		"checkpoint_management_add_data_center_object":                         7,
		"checkpoint_management_add_threat_protections":                         8,
		"checkpoint_management_add_updatable_object":                           9,
		"checkpoint_management_address_range":                                  10,
		"checkpoint_management_application_site":                               11,
		"checkpoint_management_application_site_category":                      12,
		"checkpoint_management_application_site_group":                         13,
		"checkpoint_management_assign_global_assignment":                       14,
		"checkpoint_management_backup_domain":                                  15,
		"checkpoint_management_checkpoint_host":                                16,
		"checkpoint_management_delete_api_key":                                 17,
		"checkpoint_management_delete_data_center_object":                      18,
		"checkpoint_management_delete_threat_protections":                      19,
		"checkpoint_management_delete_updatable_object":                        20,
		"checkpoint_management_discard":                                        21,
		"checkpoint_management_disconnect":                                     22,
		"checkpoint_management_dns_domain":                                     23,
		"checkpoint_management_dynamic_object":                                 24,
		"checkpoint_management_exception_group":                                25,
		"checkpoint_management_export":                                         26,
		"checkpoint_management_get_attachment":                                 27,
		"checkpoint_management_group":                                          28,
		"checkpoint_management_group_with_exclusion":                           29,
		"checkpoint_management_gsn_handover_group":                             30,
		"checkpoint_management_ha_full_sync":                                   31,
		"checkpoint_management_host":                                           32,
		"checkpoint_management_https_layer":                                    33,
		"checkpoint_management_https_rule":                                     34,
		"checkpoint_management_https_section":                                  35,
		"checkpoint_management_identity_tag":                                   36,
		"checkpoint_management_install_database":                               37,
		"checkpoint_management_install_policy":                                 38,
		"checkpoint_management_install_software_package":                       39,
		"checkpoint_management_keepalive":                                      40,
		"checkpoint_management_login":                                          41,
		"checkpoint_management_logout":                                         42,
		"checkpoint_management_mds":                                            43,
		"checkpoint_management_migrate_export_domain":                          44,
		"checkpoint_management_migrate_import_domain":                          45,
		"checkpoint_management_multicast_address_range":                        46,
		"checkpoint_management_nat_rule":                                       47,
		"checkpoint_management_nat_section":                                    48,
		"checkpoint_management_network":                                        49,
		"checkpoint_management_opsec_application":                              50,
		"checkpoint_management_package":                                        51,
		"checkpoint_management_publish":                                        52,
		"checkpoint_management_put_file":                                       53,
		"checkpoint_management_restore_domain":                                 54,
		"checkpoint_management_revert_to_revision":                             55,
		"checkpoint_management_run_ips_update":                                 56,
		"checkpoint_management_run_script":                                     57,
		"checkpoint_management_run_threat_emulation_file_types_offline_update": 58,
		"checkpoint_management_security_zone":                                  59,
		"checkpoint_management_service_citrix_tcp":                             60,
		"checkpoint_management_service_compound_tcp":                           61,
		"checkpoint_management_service_dce_rpc":                                62,
		"checkpoint_management_service_group":                                  63,
		"checkpoint_management_service_icmp":                                   64,
		"checkpoint_management_service_icmp6":                                  65,
		"checkpoint_management_service_other":                                  66,
		"checkpoint_management_service_rpc":                                    67,
		"checkpoint_management_service_sctp":                                   68,
		"checkpoint_management_service_tcp":                                    69,
		"checkpoint_management_service_udp":                                    70,
		"checkpoint_management_set_api_settings":                               71,
		"checkpoint_management_set_automatic_purge":                            72,
		"checkpoint_management_set_global_domain":                              73,
		"checkpoint_management_set_ha_state":                                   74,
		"checkpoint_management_set_ips_update_schedule":                        75,
		"checkpoint_management_set_login_message":                              76,
		"checkpoint_management_set_threat_protection":                          77,
		"checkpoint_management_simple_cluster":                                 78,
		"checkpoint_management_simple_gateway":                                 79,
		"checkpoint_management_threat_exception":                               80,
		"checkpoint_management_threat_indicator":                               81,
		"checkpoint_management_threat_rule":                                    82,
		"checkpoint_management_time_group":                                     83,
		"checkpoint_management_uninstall_software_package":                     84,
		"checkpoint_management_unlock_administrator":                           85,
		"checkpoint_management_update_updatable_objects_repository_content":    86,
		"checkpoint_management_user":                                           87,
		"checkpoint_management_user_group":                                     88,
		"checkpoint_management_user_template":                                  89,
		"checkpoint_management_verify_policy":                                  90,
		"checkpoint_management_verify_revert":                                  91,
		"checkpoint_management_verify_software_package":                        92,
		"checkpoint_management_vpn_community_meshed":                           93,
		"checkpoint_management_vpn_community_remote_access":                    94,
		"checkpoint_management_vpn_community_star":                             95,
		"checkpoint_management_where_used":                                     96,
		"checkpoint_management_wildcard":                                       97,
		"checkpoint_physical_interface":                                        98,
		"checkpoint_put_file":                                                  99,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
