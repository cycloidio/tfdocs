package checkpoint

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_data_access_layer",
			Category:         "Data Sources",
			ShortDescription: `This resource allows you to execute Check Point Access Layer.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "applications_and_url_filtering",
					Description: `Whether to enable Applications & URL Filtering blade on the layer.`,
				},
				resource.Attribute{
					Name:        "content_awareness",
					Description: `Whether to enable Content Awareness blade on the layer.`,
				},
				resource.Attribute{
					Name:        "detect_using_x_forward_for",
					Description: `Whether to use X-Forward-For HTTP header, which is added by the proxy server to keep track of the original source IP.`,
				},
				resource.Attribute{
					Name:        "firewall",
					Description: `Whether to enable Firewall blade on the layer.`,
				},
				resource.Attribute{
					Name:        "implicit_cleanup_action",
					Description: `The default "catch-all" action for traffic that does not match any explicit or implied rules in the layer.`,
				},
				resource.Attribute{
					Name:        "mobile_access",
					Description: `Whether to enable Mobile Access blade on the layer.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `Whether this layer is shared.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_data_access_role",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Access Role.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "machines",
					Description: `Machines that can access the system. machines blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "networks",
					Description: `Collection of Network objects identified by the name or UID that can access the system.`,
				},
				resource.Attribute{
					Name:        "remote_access_clients",
					Description: `Remote access clients identified by name or UID.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `Users that can access the system. users blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string. ` + "`" + `machines` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `Active Directory name or UID or Identity Tag.`,
				},
				resource.Attribute{
					Name:        "selection",
					Description: `Name or UID of an object selected from source.selection blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "base_dn",
					Description: `When source is "Active Directory" use "base-dn" to refine the query in AD database. ` + "`" + `users` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `Active Directory name or UID or Identity Tag or Internal User Groups or LDAP groups or Guests.`,
				},
				resource.Attribute{
					Name:        "selection",
					Description: `Name or UID of an object selected from source.selection blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "base_dn",
					Description: `When source is "Active Directory" use "base-dn" to refine the query in AD database.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_data_access_rule",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Access Rule.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "layer",
					Description: `(Required) Layer that the rule belongs to identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Rule name.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `\"Accept\", \"Drop\", \"Ask\", \"Inform\", \"Reject\", \"User Auth\", \"Client Auth\", \"Apply Layer\".`,
				},
				resource.Attribute{
					Name:        "action_settings",
					Description: `Action settings. Action settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `List of processed file types that this rule applies on.`,
				},
				resource.Attribute{
					Name:        "content_direction",
					Description: `On which direction the file types processing is applied.`,
				},
				resource.Attribute{
					Name:        "content_negate",
					Description: `True if negate is set for data.`,
				},
				resource.Attribute{
					Name:        "custom_fields",
					Description: `Custom fields. Custom fields blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `Collection of Network objects identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "destination_negate",
					Description: `True if negate is set for destination.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enable/Disable the rule.`,
				},
				resource.Attribute{
					Name:        "inline_layer",
					Description: `Inline Layer identified by the name or UID. Relevant only if \"Action\" was set to \"Apply Layer\".`,
				},
				resource.Attribute{
					Name:        "install_on",
					Description: `Which Gateways identified by the name or UID to install the policy on.`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `Collection of Network objects identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "service_negate",
					Description: `True if negate is set for service.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `Collection of Network objects identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "source_negate",
					Description: `True if negate is set for source.`,
				},
				resource.Attribute{
					Name:        "time",
					Description: `List of time objects. For example: \"Weekend\", \"Off-Work\", \"Every-Day\".`,
				},
				resource.Attribute{
					Name:        "track",
					Description: `Track Settings. Track Settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "user_check",
					Description: `User check settings. User check settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "vpn",
					Description: `Communities or Directional.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string. ` + "`" + `action_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "field_1",
					Description: `First custom field.`,
				},
				resource.Attribute{
					Name:        "field_2",
					Description: `Second custom field.`,
				},
				resource.Attribute{
					Name:        "field_3",
					Description: `Third custom field. ` + "`" + `track` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "accounting",
					Description: `Turns accounting for track on and off.`,
				},
				resource.Attribute{
					Name:        "alert",
					Description: `Type of alert for the track.`,
				},
				resource.Attribute{
					Name:        "enable_firewall_session",
					Description: `Determine whether to generate session log to firewall only connections.`,
				},
				resource.Attribute{
					Name:        "per_connection",
					Description: `Determines whether to perform the log per connection.`,
				},
				resource.Attribute{
					Name:        "per_session",
					Description: `Determines whether to perform the log per session.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `\"Log\", \"Extended Log\", \"Detailed Log\", \"None\". ` + "`" + `user_check` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "custom_frequency",
					Description: `Custom Frequency blocks are documented below.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_data_access_section",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Access Section.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "layer",
					Description: `(Required) Layer that the rule belongs to identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Object unique identifier.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_data_address_range",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Address Range.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name. Should be unique in the domain.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "ipv4_address_first",
					Description: `First IPv4 address in the range.`,
				},
				resource.Attribute{
					Name:        "ipv6_address_first",
					Description: `First IPv6 address in the range.`,
				},
				resource.Attribute{
					Name:        "ipv4_address_last",
					Description: `Last IPv4 address in the range.`,
				},
				resource.Attribute{
					Name:        "ipv6_address_last",
					Description: `Last IPv6 address in the range.`,
				},
				resource.Attribute{
					Name:        "nat_settings",
					Description: `NAT settings. NAT settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `Collection of group identifiers.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag identifiers. ` + "`" + `nat_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "auto_rule",
					Description: `Whether to add automatic address translation rules.`,
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
			Type:             "checkpoint_management_data_application_site",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Application Site.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "additional_categories",
					Description: `Used to configure or edit the additional categories of a custom application / site used in the Application and URL Filtering or Threat Prevention.additional_categories blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description for the application.`,
				},
				resource.Attribute{
					Name:        "primary_category",
					Description: `Each application is assigned to one primary category based on its most defining aspect.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "url_list",
					Description: `URLs that determine this particular application.`,
				},
				resource.Attribute{
					Name:        "application_signature",
					Description: `Application signature generated by <a href="https://supportcenter.checkpoint.com/supportcenter/portal?eventSubmit_doGoviewsolutiondetails=&solutionid=sk103051">Signature Tool</a>.`,
				},
				resource.Attribute{
					Name:        "urls_defined_as_regular_expression",
					Description: `States whether the URL is defined as a Regular Expression or not.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `Collection of group identifiers.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_data_application_site_category",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Application Site Category.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description string`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `Collection of group identifiers.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_data_application_site_group",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Application Site Group.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `Collection of application and URL filtering objects identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `Collection of group identifiers.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_data_dns_domain",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Dns Domain.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "is_sub_domain",
					Description: `Whether to match sub-domains in addition to the domain itself.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_data_dynamic_object",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Dynamic Object.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_data_exception_group",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Exception Group.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "applied_profile",
					Description: `The threat profile to apply this group to in the case of apply-on threat-rules-with-specific-profile.`,
				},
				resource.Attribute{
					Name:        "apply_on",
					Description: `An exception group can be set to apply on all threat rules, all threat rules which have a specific profile, or those rules manually chosen by the user.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_data_group",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Group.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name. Should be unique in the domain.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `Collection of Network objects identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `Collection of group identifiers.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag identifiers.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_data_group_with_exclusion",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Group With Exclusion.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "except",
					Description: `Name or UID of an object which the group excludes.`,
				},
				resource.Attribute{
					Name:        "include",
					Description: `Name or UID of an object which the group includes.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `Collection of group identifiers.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_data_host",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Host.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name. Should be unique in the domain.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `IPv4 address.`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `IPv6 address.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `Host interfaces. interfaces blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "nat_settings",
					Description: `NAT settings. nat_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "host_servers",
					Description: `Servers Configuration. host_servers blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `Collection of group identifiers.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag identifiers. ` + "`" + `interfaces` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Object name. Should be unique in the domain.`,
				},
				resource.Attribute{
					Name:        "subnet4",
					Description: `IPv4 network address.`,
				},
				resource.Attribute{
					Name:        "subnet6",
					Description: `IPv6 network address.`,
				},
				resource.Attribute{
					Name:        "mask_length4",
					Description: `IPv4 network mask length.`,
				},
				resource.Attribute{
					Name:        "mask_length6",
					Description: `IPv6 network mask length.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string. ` + "`" + `nat_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "auto_rule",
					Description: `Whether to add automatic address translation rules.`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `IPv4 address.`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `IPv6 address.`,
				},
				resource.Attribute{
					Name:        "hide_behind",
					Description: `Hide behind method. This parameter is not required in case \"method\" parameter is \"static\".`,
				},
				resource.Attribute{
					Name:        "install_on",
					Description: `Which gateway should apply the NAT translation.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `NAT translation method. ` + "`" + `host_servers` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "dns_server",
					Description: `Gets True if this server is a DNS Server.`,
				},
				resource.Attribute{
					Name:        "mail_server",
					Description: `Gets True if this server is a Mail Server.`,
				},
				resource.Attribute{
					Name:        "web_server",
					Description: `Gets True if this server is a Web Server.`,
				},
				resource.Attribute{
					Name:        "web_server_config",
					Description: `Web Server configuration. Web Server configuration blocks are documented below. ` + "`" + `web_server_config` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "additional_ports",
					Description: `Server additional ports.`,
				},
				resource.Attribute{
					Name:        "application_engines",
					Description: `Application engines of this web server.`,
				},
				resource.Attribute{
					Name:        "listen_standard_port",
					Description: `"Whether server listens to standard port.`,
				},
				resource.Attribute{
					Name:        "operating_system",
					Description: `Operating System.`,
				},
				resource.Attribute{
					Name:        "protected_by",
					Description: `Network object which protects this server identified by the name or UID.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_data_https_layer",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Https Layer.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `Define the Layer as Shared (TRUE/FALSE).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_data_https_rule",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Https Rule.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "layer",
					Description: `(Required) Layer that holds the Object. Identified by the Name or UID.`,
				},
				resource.Attribute{
					Name:        "rule_number",
					Description: `(Optional) Rule number.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `HTTPS rule name.`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `Collection of Network objects identified by Name or UID that represents connection destination.`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `Collection of Network objects identified by Name or UID that represents connection service.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `Collection of Network objects identified by Name or UID that represents connection source.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Rule inspect level. "Bypass" or "Inspect".`,
				},
				resource.Attribute{
					Name:        "blade",
					Description: `Blades for HTTPS Inspection. Identified by Name or UID to enable the inspection for. "Anti Bot","Anti Virus","Application Control","Data Awareness","DLP","IPS","Threat Emulation","Url Filtering".`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `Internal Server Certificate identified by Name or UID. otherwise, "Outbound Certificate" is a default value.`,
				},
				resource.Attribute{
					Name:        "destination_negate",
					Description: `TRUE if "negate" value is set for Destination.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enable/Disable the rule.`,
				},
				resource.Attribute{
					Name:        "install_on",
					Description: `Which Gateways identified by the name or UID to install the policy on.`,
				},
				resource.Attribute{
					Name:        "service_negate",
					Description: `TRUE if "negate" value is set for Service.`,
				},
				resource.Attribute{
					Name:        "site_category",
					Description: `Collection of Site Categories objects identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "site_category_negate",
					Description: `TRUE if "negate" value is set for Site Category.`,
				},
				resource.Attribute{
					Name:        "source_negate",
					Description: `TRUE if "negate" value is set for Source.`,
				},
				resource.Attribute{
					Name:        "track",
					Description: `"None","Log","Alert","Mail","SNMP trap","Mail","User Alert", "User Alert 2", "User Alert 3".`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_data_https_section",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Https Section.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "layer",
					Description: `(Required) Layer that holds the Object. Identified by the Name or UID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Object unique identifier.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_data_multicast_address_range",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Multicast Address Range.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `IPv4 address.`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `IPv6 address.`,
				},
				resource.Attribute{
					Name:        "ipv4_address_first",
					Description: `First IPv4 address in the range.`,
				},
				resource.Attribute{
					Name:        "ipv6_address_first",
					Description: `First IPv6 address in the range.`,
				},
				resource.Attribute{
					Name:        "ipv4_address_last",
					Description: `Last IPv4 address in the range.`,
				},
				resource.Attribute{
					Name:        "ipv6_address_last",
					Description: `Last IPv6 address in the range.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `Collection of group identifiers.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_data_network",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Network Object.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name. Should be unique in the domain.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "subnet4",
					Description: `IPv4 network address.`,
				},
				resource.Attribute{
					Name:        "subnet6",
					Description: `IPv6 network address.`,
				},
				resource.Attribute{
					Name:        "mask_length4",
					Description: `IPv4 network mask length.`,
				},
				resource.Attribute{
					Name:        "mask_length6",
					Description: `IPv6 network mask length.`,
				},
				resource.Attribute{
					Name:        "nat_settings",
					Description: `NAT settings. nat_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `Collection of group identifiers.`,
				},
				resource.Attribute{
					Name:        "broadcast",
					Description: `Allow broadcast address inclusion.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string. ` + "`" + `nat_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "auto_rule",
					Description: `Whether to add automatic address translation rules.`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `IPv4 address.`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `IPv6 address.`,
				},
				resource.Attribute{
					Name:        "hide_behind",
					Description: `Hide behind method. This parameter is not required in case \"method\" parameter is \"static\".`,
				},
				resource.Attribute{
					Name:        "install_on",
					Description: `Which gateway should apply the NAT translation.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `NAT translation method.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_data_opsec_application",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Opsec Application.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "cpmi",
					Description: `Used to setup the CPMI client entity. cpmi blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `The host where the server is running. Pre-define the host as a network object.`,
				},
				resource.Attribute{
					Name:        "lea",
					Description: `Used to setup the LEA client entity. lea blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string. ` + "`" + `cpmi` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "administrator_profile",
					Description: `A profile to set the log reading permissions by for the client entity.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Whether to enable this client entity on the Opsec Application.`,
				},
				resource.Attribute{
					Name:        "use_administrator_credentials",
					Description: `Whether to use the Admin's credentials to login to the security management server. ` + "`" + `lea` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "access_permissions",
					Description: `Log reading permissions for the LEA client entity.`,
				},
				resource.Attribute{
					Name:        "administrator_profile",
					Description: `A profile to set the log reading permissions by for the client entity.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Whether to enable this client entity on the Opsec Application.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_data_package",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Package Object.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name. Should be unique in the domain.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "access",
					Description: `True - enables, False - disables access & NAT policies, empty - nothing is changed.`,
				},
				resource.Attribute{
					Name:        "desktop_security",
					Description: `True - enables, False - disables Desktop security policy, empty - nothing is changed.`,
				},
				resource.Attribute{
					Name:        "qos",
					Description: `True - enables, False - disables QoS policy, empty - nothing is changed.`,
				},
				resource.Attribute{
					Name:        "qos_policy_type",
					Description: `QoS policy type.`,
				},
				resource.Attribute{
					Name:        "threat_prevention",
					Description: `True - enables, False - disables Threat policy, empty - nothing is changed.`,
				},
				resource.Attribute{
					Name:        "vpn_traditional_mode",
					Description: `True - enables, False - disables VPN traditional mode, empty - nothing is changed.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag identifiers.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_data_security_zone",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Security Zone.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_data_service_dce_rpc",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Service Dce Rpc.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "interface_uuid",
					Description: `Network interface UUID.`,
				},
				resource.Attribute{
					Name:        "keep_connections_open_after_policy_installation",
					Description: `Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `Collection of group identifiers.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_data_service_group",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Service Group.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name. Should be unique in the domain.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `Collection of Network objects identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `Collection of group identifiers.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag identifiers.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_data_service_icmp",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Service Icmp.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "icmp_code",
					Description: `As listed in: <a href="http://www.iana.org/assignments/icmp-parameters" target="_blank">RFC 792</a>.`,
				},
				resource.Attribute{
					Name:        "icmp_type",
					Description: `As listed in: <a href="http://www.iana.org/assignments/icmp-parameters" target="_blank">RFC 792</a>.`,
				},
				resource.Attribute{
					Name:        "keep_connections_open_after_policy_installation",
					Description: `Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `Collection of group identifiers.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_data_service_icmp6",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Service Icmp6.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "icmp_code",
					Description: `As listed in: <a href="http://www.iana.org/assignments/icmp-parameters" target="_blank">RFC 792</a>.`,
				},
				resource.Attribute{
					Name:        "icmp_type",
					Description: `As listed in: <a href="http://www.iana.org/assignments/icmp-parameters" target="_blank">RFC 792</a>.`,
				},
				resource.Attribute{
					Name:        "keep_connections_open_after_policy_installation",
					Description: `(Optional) Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `Collection of group identifiers.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_data_service_other",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Service Other.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "accept_replies",
					Description: `Specifies whether Other Service replies are to be accepted.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Contains an INSPECT expression that defines the action to take if a rule containing this service is matched. Example: set r_mhandler &open_ssl_handler sets a handler on the connection.`,
				},
				resource.Attribute{
					Name:        "aggressive_aging",
					Description: `Sets short (aggressive) timeouts for idle connections. aggressive_aging blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `IP protocol number.`,
				},
				resource.Attribute{
					Name:        "keep_connections_open_after_policy_installation",
					Description: `Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections.`,
				},
				resource.Attribute{
					Name:        "match",
					Description: `Contains an INSPECT expression that defines the matching criteria. The connection is examined against the expression during the first packet. Example: tcp, dport = 21, direction = 0 matches incoming FTP control connections.`,
				},
				resource.Attribute{
					Name:        "match_for_any",
					Description: `Indicates whether this service is used when 'Any' is set as the rule's service and there are several service objects with the same source port and protocol.`,
				},
				resource.Attribute{
					Name:        "override_default_settings",
					Description: `Indicates whether this service is a Data Domain service which has been overridden.`,
				},
				resource.Attribute{
					Name:        "session_timeout",
					Description: `Time (in seconds) before the session times out.`,
				},
				resource.Attribute{
					Name:        "sync_connections_on_cluster",
					Description: `Enables state-synchronized High Availability or Load Sharing on a ClusterXL or OPSEC-certified cluster.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "use_default_session_timeout",
					Description: `Use default virtual session timeout.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `Collection of group identifiers. ` + "`" + `aggressive_aging` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "default_timeout",
					Description: `Default aggressive aging timeout in seconds.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `Aggressive aging timeout in seconds.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_data_service_rpc",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Service Rpc.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "keep_connections_open_after_policy_installation",
					Description: `Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `Collection of group identifiers.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_data_service_sctp",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Service Sctp.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "aggressive_aging",
					Description: `Sets short (aggressive) timeouts for idle connections. aggressive_aging blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "keep_connections_open_after_policy_installation",
					Description: `Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections.`,
				},
				resource.Attribute{
					Name:        "match_for_any",
					Description: `Indicates whether this service is used when 'Any' is set as the rule's service and there are several service objects with the same source port and protocol.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port number. To specify a port range add a hyphen between the lowest and the highest port numbers, for example 44-45.`,
				},
				resource.Attribute{
					Name:        "session_timeout",
					Description: `Time (in seconds) before the session times out.`,
				},
				resource.Attribute{
					Name:        "source_port",
					Description: `Source port number. To specify a port range add a hyphen between the lowest and the highest port numbers, for example 44-45.`,
				},
				resource.Attribute{
					Name:        "sync_connections_on_cluster",
					Description: `Enables state-synchronized High Availability or Load Sharing on a ClusterXL or OPSEC-certified cluster.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "use_default_session_timeout",
					Description: `Use default virtual session timeout.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `Collection of group identifiers. ` + "`" + `aggressive_aging` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "default_timeout",
					Description: `Default aggressive aging timeout in seconds.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `Aggressive aging timeout in seconds.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_data_service_tcp",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Service Tcp.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name. Should be unique in the domain.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The number of the port used to provide this service. To specify a port range, place a hyphen between the lowest and highest port numbers, for example 44-55.`,
				},
				resource.Attribute{
					Name:        "aggressive_aging",
					Description: `Sets short (aggressive) timeouts for idle connections. Aggressive Aging blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "keep_connections_open_after_policy_installation",
					Description: `Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections.`,
				},
				resource.Attribute{
					Name:        "match_by_protocol_signature",
					Description: `A value of true enables matching by the selected protocol's signature - the signature identifies the protocol as genuine. Select this option to limit the port to the specified protocol. If the selected protocol does not support matching by signature, this field cannot be set to true.`,
				},
				resource.Attribute{
					Name:        "match_for_any",
					Description: `Indicates whether this service is used when 'Any' is set as the rule's service and there are several service objects with the same source port and protocol.`,
				},
				resource.Attribute{
					Name:        "override_default_settings",
					Description: `Indicates whether this service is a Data Domain service which has been overridden.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Select the protocol type associated with the service, and by implication, the management server (if any) that enforces Content Security and Authentication for the service. Selecting a Protocol Type invokes the specific protocol handlers for each protocol type, thus enabling higher level of security by parsing the protocol, and higher level of connectivity by tracking dynamic actions (such as opening of ports).`,
				},
				resource.Attribute{
					Name:        "session_timeout",
					Description: `Time (in seconds) before the session times out.`,
				},
				resource.Attribute{
					Name:        "source_port",
					Description: `Port number for the client side service. If specified, only those Source port Numbers will be Accepted, Dropped, or Rejected during packet inspection. Otherwise, the source port is not inspected.`,
				},
				resource.Attribute{
					Name:        "sync_connections_on_cluster",
					Description: `Enables state-synchronized High Availability or Load Sharing on a ClusterXL or OPSEC-certified cluster.`,
				},
				resource.Attribute{
					Name:        "use_default_session_timeout",
					Description: `Use default virtual session timeout.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `Collection of group identifiers.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag identifiers. ` + "`" + `aggressive_aging` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "default_timeout",
					Description: `(Optional) Default aggressive aging timeout in seconds.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) Aggressive aging timeout in seconds.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_data_service_udp",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Service Udp.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name. Should be unique in the domain.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "aggressive_aging",
					Description: `Sets short (aggressive) timeouts for idle connections. Aggressive Aging blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "keep_connections_open_after_policy_installation",
					Description: `Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections.`,
				},
				resource.Attribute{
					Name:        "match_by_protocol_signature",
					Description: `A value of true enables matching by the selected protocol's signature - the signature identifies the protocol as genuine. Select this option to limit the port to the specified protocol. If the selected protocol does not support matching by signature, this field cannot be set to true.`,
				},
				resource.Attribute{
					Name:        "match_for_any",
					Description: `Indicates whether this service is used when 'Any' is set as the rule's service and there are several service objects with the same source port and protocol.`,
				},
				resource.Attribute{
					Name:        "override_default_settings",
					Description: `Indicates whether this service is a Data Domain service which has been overridden.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The number of the port used to provide this service. To specify a port range, place a hyphen between the lowest and highest port numbers, for example 44-55.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Select the protocol type associated with the service, and by implication, the management server (if any) that enforces Content Security and Authentication for the service. Selecting a Protocol Type invokes the specific protocol handlers for each protocol type, thus enabling higher level of security by parsing the protocol, and higher level of connectivity by tracking dynamic actions (such as opening of ports).`,
				},
				resource.Attribute{
					Name:        "session_timeout",
					Description: `Time (in seconds) before the session times out.`,
				},
				resource.Attribute{
					Name:        "source_port",
					Description: `Port number for the client side service. If specified, only those Source port Numbers will be Accepted, Dropped, or Rejected during packet inspection. Otherwise, the source port is not inspected.`,
				},
				resource.Attribute{
					Name:        "sync_connections_on_cluster",
					Description: `Enables state-synchronized High Availability or Load Sharing on a ClusterXL or OPSEC-certified cluster.`,
				},
				resource.Attribute{
					Name:        "use_default_session_timeout",
					Description: `Use default virtual session timeout.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `Collection of group identifiers.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag identifiers. ` + "`" + `aggressive_aging` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "default_timeout",
					Description: `(Optional) Default aggressive aging timeout in seconds.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) Aggressive aging timeout in seconds.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_data_threat_indicator",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Threat Indicator.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name. Should be unique in the domain.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `The indicator's action.`,
				},
				resource.Attribute{
					Name:        "profile_overrides",
					Description: `Profiles in which to override the indicator's default action. Profile Overrides blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag identifiers. ` + "`" + `profile_overrides` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `The indicator's action in this profile.`,
				},
				resource.Attribute{
					Name:        "profile",
					Description: `The profile in which to override the indicator's action.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_data_time_group",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Time Group.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `Collection of Time Group objects identified by the name or UID.members blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `Collection of group identifiers.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_data_vpn_community_meshed",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Vpn Community Meshed.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "encryption_method",
					Description: `The encryption method to be used.`,
				},
				resource.Attribute{
					Name:        "encryption_suite",
					Description: `The encryption suite to be used.`,
				},
				resource.Attribute{
					Name:        "gateways",
					Description: `Collection of Gateway objects identified by the name or UID. gateways blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "ike_phase_1",
					Description: `Ike Phase 1 settings. Only applicable when the encryption-suite is set to [custom]. ike_phase_1 blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "ike_phase_2",
					Description: `Ike Phase 2 settings. Only applicable when the encryption-suite is set to [custom]. ike_phase_2 blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "override_vpn_domains",
					Description: `The Overrides VPN Domains of the participants GWs. override_vpn_domains blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "shared_secrets",
					Description: `Shared secrets for external gateways. shared_secrets blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "use_shared_secret",
					Description: `Indicates whether the shared secret should be used for all external gateways.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string. ` + "`" + `ike_phase_1` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "data_integrity",
					Description: `The hash algorithm to be used.`,
				},
				resource.Attribute{
					Name:        "diffie_hellman_group",
					Description: `The Diffie-Hellman group to be used.`,
				},
				resource.Attribute{
					Name:        "encryption_algorithm",
					Description: `The encryption algorithm to be used. ` + "`" + `ike_phase_2` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "data_integrity",
					Description: `The hash algorithm to be used.`,
				},
				resource.Attribute{
					Name:        "encryption_algorithm",
					Description: `The encryption algorithm to be used. ` + "`" + `override_vpn_domains` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `Participant gateway in override VPN domain identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "vpn_domain",
					Description: `VPN domain network identified by the name or UID. ` + "`" + `shared_secrets` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "external_gateway",
					Description: `External gateway identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "shared_secret",
					Description: `Shared secret.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_data_vpn_community_star",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Vpn Community Star.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "center_gateways",
					Description: `Collection of Gateway objects representing center gateways identified by the name or UID. center_gateways blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "encryption_method",
					Description: `The encryption method to be used.`,
				},
				resource.Attribute{
					Name:        "encryption_suite",
					Description: `The encryption suite to be used.`,
				},
				resource.Attribute{
					Name:        "ike_phase_1",
					Description: `Ike Phase 1 settings. Only applicable when the encryption-suite is set to [custom]. ike_phase_1 blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "ike_phase_2",
					Description: `Ike Phase 2 settings. Only applicable when the encryption-suite is set to [custom]. ike_phase_2 blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "mesh_center_gateways",
					Description: `Indicates whether the meshed community is in center.`,
				},
				resource.Attribute{
					Name:        "override_vpn_domains",
					Description: `The Overrides VPN Domains of the participants GWs. override_vpn_domains blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "satellite_gateways",
					Description: `Collection of Gateway objects representing satellite gateways identified by the name or UID. satellite_gateways blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "shared_secrets",
					Description: `Shared secrets for external gateways. shared_secrets blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "use_shared_secret",
					Description: `Indicates whether the shared secret should be used for all external gateways.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. ` + "`" + `ike_phase_1` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "data_integrity",
					Description: `The hash algorithm to be used.`,
				},
				resource.Attribute{
					Name:        "diffie_hellman_group",
					Description: `The Diffie-Hellman group to be used.`,
				},
				resource.Attribute{
					Name:        "encryption_algorithm",
					Description: `The encryption algorithm to be used. ` + "`" + `ike_phase_2` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "data_integrity",
					Description: `The hash algorithm to be used.`,
				},
				resource.Attribute{
					Name:        "encryption_algorithm",
					Description: `The encryption algorithm to be used. ` + "`" + `override_vpn_domains` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `Participant gateway in override VPN domain identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "vpn_domain",
					Description: `VPN domain network identified by the name or UID. ` + "`" + `shared_secrets` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "external_gateway",
					Description: `External gateway identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "shared_secret",
					Description: `Shared secret.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_data_wildcard",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Wildcard.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `IPv4 address.`,
				},
				resource.Attribute{
					Name:        "ipv4_mask_wildcard",
					Description: `IPv4 mask wildcard.`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `IPv6 address.`,
				},
				resource.Attribute{
					Name:        "ipv6_mask_wildcard",
					Description: `IPv6 mask wildcard.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `Collection of group identifiers.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"checkpoint_management_data_access_layer":              0,
		"checkpoint_management_data_access_role":               1,
		"checkpoint_management_data_access_rule":               2,
		"checkpoint_management_data_access_section":            3,
		"checkpoint_management_data_address_range":             4,
		"checkpoint_management_data_application_site":          5,
		"checkpoint_management_data_application_site_category": 6,
		"checkpoint_management_data_application_site_group":    7,
		"checkpoint_management_data_dns_domain":                8,
		"checkpoint_management_data_dynamic_object":            9,
		"checkpoint_management_data_exception_group":           10,
		"checkpoint_management_data_group":                     11,
		"checkpoint_management_data_group_with_exclusion":      12,
		"checkpoint_management_data_host":                      13,
		"checkpoint_management_data_https_layer":               14,
		"checkpoint_management_data_https_rule":                15,
		"checkpoint_management_data_https_section":             16,
		"checkpoint_management_data_multicast_address_range":   17,
		"checkpoint_management_data_network":                   18,
		"checkpoint_management_data_opsec_application":         19,
		"checkpoint_management_data_package":                   20,
		"checkpoint_management_data_security_zone":             21,
		"checkpoint_management_data_service_dce_rpc":           22,
		"checkpoint_management_data_service_group":             23,
		"checkpoint_management_data_service_icmp":              24,
		"checkpoint_management_data_service_icmp6":             25,
		"checkpoint_management_data_service_other":             26,
		"checkpoint_management_data_service_rpc":               27,
		"checkpoint_management_data_service_sctp":              28,
		"checkpoint_management_data_service_tcp":               29,
		"checkpoint_management_data_service_udp":               30,
		"checkpoint_management_data_threat_indicator":          31,
		"checkpoint_management_data_time_group":                32,
		"checkpoint_management_data_vpn_community_meshed":      33,
		"checkpoint_management_data_vpn_community_star":        34,
		"checkpoint_management_data_wildcard":                  35,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
