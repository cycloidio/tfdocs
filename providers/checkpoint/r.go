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
			Category:         "Resources",
			ShortDescription: `This resource allows you to set the hostname of a Check Point machine.`,
			Description:      ``,
			Keywords: []string{
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
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Access Layer.`,
			Description:      ``,
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
			Type:             "checkpoint_management_access_role",
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Access Role.`,
			Description:      ``,
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
			Category:         "Resources",
			ShortDescription: `This resource allows you to add/update/delete Check Point Access Rule.`,
			Description:      ``,
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
					Description: `(Optional) N/A.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_access_section",
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Access Section.`,
			Description:      ``,
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
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Add Api Key.`,
			Description:      ``,
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
					Description: `(Required) Administrator name to generate API key for. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_add_data_center_object",
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Add Data Center Object.`,
			Description:      ``,
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
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Add Threat Protections.`,
			Description:      ``,
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
					Description: `(Optional) Protections package path. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_add_updatable_object",
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Add Updatable Object.`,
			Description:      ``,
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
			Category:         "Resources",
			ShortDescription: `This resource allows you to add/update/delete Check Point Address Range.`,
			Description:      ``,
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
					Name:        "groups",
					Description: `(Optional) Collection of group identifiers.`,
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
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Application Site.`,
			Description:      ``,
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
			Type:             "checkpoint_management_application_site_category",
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Application Site Category.`,
			Description:      ``,
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
			Type:             "checkpoint_management_application_site_group",
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Application Site Group.`,
			Description:      ``,
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
			Type:             "checkpoint_management_assign_global_assignment",
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Assign Global Assignment.`,
			Description:      ``,
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
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Backup Domain.`,
			Description:      ``,
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
			Type:             "checkpoint_management_delete_api_key",
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Delete Api Key.`,
			Description:      ``,
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
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Delete Data Center Object.`,
			Description:      ``,
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
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Delete Threat Protections.`,
			Description:      ``,
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
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Delete Updatable Object.`,
			Description:      ``,
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
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Discard.`,
			Description:      ``,
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
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Disconnect.`,
			Description:      ``,
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
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Dns Domain.`,
			Description:      ``,
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
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Dynamic Object.`,
			Description:      ``,
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
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Exception Group.`,
			Description:      ``,
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
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Export.`,
			Description:      ``,
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
					Description: `(Optional) N/A ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_group",
			Category:         "Resources",
			ShortDescription: `This resource allows you to add/update/delete Check Point Group.`,
			Description:      ``,
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
					Name:        "groups",
					Description: `(Optional) Collection of group identifiers.`,
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
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Group With Exclusion.`,
			Description:      ``,
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
			Type:             "checkpoint_management_host",
			Category:         "Resources",
			ShortDescription: `This resource allows you to add/update/delete Check Point Host.`,
			Description:      ``,
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
					Name:        "groups",
					Description: `(Optional) Collection of group identifiers.`,
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
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Https Layer.`,
			Description:      ``,
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
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Https Rule.`,
			Description:      ``,
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
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Https Section.`,
			Description:      ``,
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
			Type:             "checkpoint_management_install_database",
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Install Database.`,
			Description:      ``,
			Keywords: []string{
				"management",
				"install",
				"database",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "targets",
					Description: `(Required) Check Point host(s) with one or more Management Software Blades enabled. The targets can be identified by their name or unique identifier.targets blocks are documented below. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_install_policy",
			Category:         "Resources",
			ShortDescription: `Install the published policy.`,
			Description:      ``,
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
					Description: `(Optional) The UID of the revision of the policy to install. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_install_software_package",
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Install Software Package.`,
			Description:      ``,
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
					Description: `(Optional) The number of targets, on which the same package is installed at the same time. ` + "`" + `cluster_installation_settings` + "`" + ` supports the following:`,
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
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Keepalive.`,
			Description:      ``,
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
			Category:         "Resources",
			ShortDescription: `Log in to the server with username and password.`,
			Description:      ``,
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
			Category:         "Resources",
			ShortDescription: `Log out from the current session. After logging out the session id is not valid any more.`,
			Description:      ``,
			Keywords: []string{
				"management",
				"logout",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_migrate_export_domain",
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Migrate Export Domain.`,
			Description:      ``,
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
					Description: `(Optional) Export logs. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_migrate_import_domain",
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Migrate Import Domain.`,
			Description:      ``,
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
					Description: `(Optional) Import logs from the input package. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_multicast_address_range",
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Multicast Address Range.`,
			Description:      ``,
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
			Type:             "checkpoint_management_network",
			Category:         "Resources",
			ShortDescription: `This resource allows you to add/update/delete Check Point Network Object.`,
			Description:      ``,
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
					Name:        "groups",
					Description: `(Optional) Collection of group identifiers.`,
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
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Opsec Application.`,
			Description:      ``,
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
			Category:         "Resources",
			ShortDescription: `This resource allows you to add/update/delete Check Point Package Object.`,
			Description:      ``,
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
			Category:         "Resources",
			ShortDescription: `Publish Changes.`,
			Description:      ``,
			Keywords: []string{
				"management",
				"publish",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Session unique identifier. Specify it to publish a different session than the one you currently use. ## How To Use Make sure this command resource will be executed by terraform when you meant it will run.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_restore_domain",
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Restore Domain.`,
			Description:      ``,
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
					Description: `(Optional) If true, verify that the import operation is valid for this input file and this environment <br>Note: Restore operation will not be executed. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_revert_to_revision",
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Revert To Revision.`,
			Description:      ``,
			Keywords: []string{
				"management",
				"revert",
				"to",
				"revision",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "to_session",
					Description: `(Optional) Session unique identifier. Specify the session id you would like to revert your database to. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_run_ips_update",
			Category:         "Resources",
			ShortDescription: `Runs IPS database update. If "package-path" is not provided server will try to get the latest package from the User Center.`,
			Description:      ``,
			Keywords: []string{
				"management",
				"run",
				"ips",
				"update",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "package_path",
					Description: `(Optional) Offline update package path. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_run_script",
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Run Script.`,
			Description:      ``,
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
					Description: `(Optional) Comments string. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_run_threat_emulation_file_types_offline_update",
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Run Threat Emulation File Types Offline Update.`,
			Description:      ``,
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
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Security Zone.`,
			Description:      ``,
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
			Type:             "checkpoint_management_service_dce_rpc",
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Service Dce Rpc.`,
			Description:      ``,
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
			Type:             "checkpoint_management_service_group",
			Category:         "Resources",
			ShortDescription: `This resource allows you to add/update/delete Check Point Service Group.`,
			Description:      ``,
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
					Name:        "groups",
					Description: `(Optional) Collection of group identifiers.`,
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
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Service Icmp.`,
			Description:      ``,
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
			Type:             "checkpoint_management_service_icmp6",
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Service Icmp6.`,
			Description:      ``,
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
			Type:             "checkpoint_management_service_other",
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Service Other.`,
			Description:      ``,
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
					Name:        "groups",
					Description: `(Optional) Collection of group identifiers.groups blocks are documented below.`,
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
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Service Rpc.`,
			Description:      ``,
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
			Type:             "checkpoint_management_service_sctp",
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Service Sctp.`,
			Description:      ``,
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
					Name:        "groups",
					Description: `(Optional) Collection of group identifiers.groups blocks are documented below.`,
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
			Category:         "Resources",
			ShortDescription: `This resource allows you to add/update/delete Check Point Service Tcp.`,
			Description:      ``,
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
					Name:        "groups",
					Description: `(Optional) Collection of group identifiers.`,
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
			Category:         "Resources",
			ShortDescription: `This resource allows you to add/update/delete Check Point Service Udp.`,
			Description:      ``,
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
					Name:        "groups",
					Description: `(Optional) Collection of group identifiers.`,
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
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Set Api Settings.`,
			Description:      ``,
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
			Type:             "checkpoint_management_set_global_domain",
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Set Global Domain.`,
			Description:      ``,
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
					Description: `(Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. ` + "`" + `servers` + "`" + ` supports the following:`,
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
			Type:             "checkpoint_management_set_ips_update_schedule",
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Set Ips Update Schedule.`,
			Description:      ``,
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
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Set Login Message.`,
			Description:      ``,
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
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Set Threat Protection.`,
			Description:      ``,
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
			Type:             "checkpoint_management_threat_indicator",
			Category:         "Resources",
			ShortDescription: `This resource allows you to add/update/delete Check Point Threat Indicator.`,
			Description:      ``,
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
			Type:             "checkpoint_management_time_group",
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Time Group.`,
			Description:      ``,
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
			Type:             "checkpoint_management_uninstall_software_package",
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Uninstall Software Package.`,
			Description:      ``,
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
					Description: `(Optional) The number of targets, on which the same package is installed at the same time. ` + "`" + `cluster_installation_settings` + "`" + ` supports the following:`,
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
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Unlock Administrator.`,
			Description:      ``,
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
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Update Updatable Objects Repository Content.`,
			Description:      ``,
			Keywords: []string{
				"management",
				"update",
				"updatable",
				"objects",
				"repository",
				"content",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_verify_policy",
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Verify Policy.`,
			Description:      ``,
			Keywords: []string{
				"management",
				"verify",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_package",
					Description: `(Required) Policy package identified by the name or UID. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_verify_revert",
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Verify Revert.`,
			Description:      ``,
			Keywords: []string{
				"management",
				"verify",
				"revert",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "to_session",
					Description: `(Required) Session unique identifier. Specify the session you would like to verify a revert operation to. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_verify_software_package",
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Verify Software Package.`,
			Description:      ``,
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
					Description: `(Optional) The number of targets, on which the same package is installed at the same time. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_vpn_community_meshed",
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Vpn Community Meshed.`,
			Description:      ``,
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
			Type:             "checkpoint_management_vpn_community_star",
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Vpn Community Star.`,
			Description:      ``,
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
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Where Used.`,
			Description:      ``,
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
			Category:         "Resources",
			ShortDescription: `This resource allows you to execute Check Point Wildcard.`,
			Description:      ``,
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
			Type:             "checkpoint_physical_interface",
			Category:         "Resources",
			ShortDescription: `This resource allows you to set a Physical interface.`,
			Description:      ``,
			Keywords: []string{
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
			Category:         "Resources",
			ShortDescription: `This resource allows you to add a new file to a Check Point machine.`,
			Description:      ``,
			Keywords: []string{
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
		"checkpoint_management_access_role":                                    2,
		"checkpoint_management_access_rule":                                    3,
		"checkpoint_management_access_section":                                 4,
		"checkpoint_management_add_api_key":                                    5,
		"checkpoint_management_add_data_center_object":                         6,
		"checkpoint_management_add_threat_protections":                         7,
		"checkpoint_management_add_updatable_object":                           8,
		"checkpoint_management_address_range":                                  9,
		"checkpoint_management_application_site":                               10,
		"checkpoint_management_application_site_category":                      11,
		"checkpoint_management_application_site_group":                         12,
		"checkpoint_management_assign_global_assignment":                       13,
		"checkpoint_management_backup_domain":                                  14,
		"checkpoint_management_delete_api_key":                                 15,
		"checkpoint_management_delete_data_center_object":                      16,
		"checkpoint_management_delete_threat_protections":                      17,
		"checkpoint_management_delete_updatable_object":                        18,
		"checkpoint_management_discard":                                        19,
		"checkpoint_management_disconnect":                                     20,
		"checkpoint_management_dns_domain":                                     21,
		"checkpoint_management_dynamic_object":                                 22,
		"checkpoint_management_exception_group":                                23,
		"checkpoint_management_export":                                         24,
		"checkpoint_management_group":                                          25,
		"checkpoint_management_group_with_exclusion":                           26,
		"checkpoint_management_host":                                           27,
		"checkpoint_management_https_layer":                                    28,
		"checkpoint_management_https_rule":                                     29,
		"checkpoint_management_https_section":                                  30,
		"checkpoint_management_install_database":                               31,
		"checkpoint_management_install_policy":                                 32,
		"checkpoint_management_install_software_package":                       33,
		"checkpoint_management_keepalive":                                      34,
		"checkpoint_management_login":                                          35,
		"checkpoint_management_logout":                                         36,
		"checkpoint_management_migrate_export_domain":                          37,
		"checkpoint_management_migrate_import_domain":                          38,
		"checkpoint_management_multicast_address_range":                        39,
		"checkpoint_management_network":                                        40,
		"checkpoint_management_opsec_application":                              41,
		"checkpoint_management_package":                                        42,
		"checkpoint_management_publish":                                        43,
		"checkpoint_management_restore_domain":                                 44,
		"checkpoint_management_revert_to_revision":                             45,
		"checkpoint_management_run_ips_update":                                 46,
		"checkpoint_management_run_script":                                     47,
		"checkpoint_management_run_threat_emulation_file_types_offline_update": 48,
		"checkpoint_management_security_zone":                                  49,
		"checkpoint_management_service_dce_rpc":                                50,
		"checkpoint_management_service_group":                                  51,
		"checkpoint_management_service_icmp":                                   52,
		"checkpoint_management_service_icmp6":                                  53,
		"checkpoint_management_service_other":                                  54,
		"checkpoint_management_service_rpc":                                    55,
		"checkpoint_management_service_sctp":                                   56,
		"checkpoint_management_service_tcp":                                    57,
		"checkpoint_management_service_udp":                                    58,
		"checkpoint_management_set_api_settings":                               59,
		"checkpoint_management_set_global_domain":                              60,
		"checkpoint_management_set_ips_update_schedule":                        61,
		"checkpoint_management_set_login_message":                              62,
		"checkpoint_management_set_threat_protection":                          63,
		"checkpoint_management_threat_indicator":                               64,
		"checkpoint_management_time_group":                                     65,
		"checkpoint_management_uninstall_software_package":                     66,
		"checkpoint_management_unlock_administrator":                           67,
		"checkpoint_management_update_updatable_objects_repository_content":    68,
		"checkpoint_management_verify_policy":                                  69,
		"checkpoint_management_verify_revert":                                  70,
		"checkpoint_management_verify_software_package":                        71,
		"checkpoint_management_vpn_community_meshed":                           72,
		"checkpoint_management_vpn_community_star":                             73,
		"checkpoint_management_where_used":                                     74,
		"checkpoint_management_wildcard":                                       75,
		"checkpoint_physical_interface":                                        76,
		"checkpoint_put_file":                                                  77,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
