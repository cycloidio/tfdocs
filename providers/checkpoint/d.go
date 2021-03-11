package checkpoint

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_access_point_name",
			Category:         "Data Sources",
			ShortDescription: `This resource allows you to execute Check Point Access Point Name.`,
			Description: `

This resource allows you to execute Check Point Access Point Name.

`,
			Keywords: []string{},
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
					Name:        "apn",
					Description: `APN name.`,
				},
				resource.Attribute{
					Name:        "enforce_end_user_domain",
					Description: `Enable enforce end user domain.`,
				},
				resource.Attribute{
					Name:        "block_traffic_other_end_user_domains",
					Description: `Block MS to MS traffic between this and other APN end user domains.`,
				},
				resource.Attribute{
					Name:        "block_traffic_this_end_user_domain",
					Description: `Block MS to MS traffic within this end user domain.`,
				},
				resource.Attribute{
					Name:        "end_user_domain",
					Description: `End user domain name or UID.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object.`,
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
			Type:             "checkpoint_management_checkpoint_host",
			Category:         "Data Sources",
			ShortDescription: `This resource allows you to execute Check Point Checkpoint Host.`,
			Description: `

This resource allows you to execute Check Point Checkpoint Host.

`,
			Keywords: []string{},
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
					Name:        "interfaces",
					Description: `Checkpoint host interfaces. interfaces blocks are documented below.`,
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
					Name:        "nat_settings",
					Description: `NAT settings. nat_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "hardware",
					Description: `Hardware name.`,
				},
				resource.Attribute{
					Name:        "os",
					Description: `Operating system name.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Checkpoint host platform version.`,
				},
				resource.Attribute{
					Name:        "management_blades",
					Description: `Management blades. management_blades blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "logs_settings",
					Description: `Logs settings. logs_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "save_logs_locally",
					Description: `Enable save logs locally.`,
				},
				resource.Attribute{
					Name:        "send_alerts_to_server",
					Description: `Collection of Server(s) to send alerts to identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "send_logs_to_backup_server",
					Description: `Collection of Backup server(s) to send logs to identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "send_logs_to_server",
					Description: `Collection of Server(s) to send logs to identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string.`,
				},
				resource.Attribute{
					Name:        "sic_name",
					Description: `Name of the Secure Internal Connection Trust.`,
				},
				resource.Attribute{
					Name:        "sic_state",
					Description: `State the Secure Internal Connection Trust. ` + "`" + `interfaces` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Interface name.`,
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
					Name:        "subnet_mask",
					Description: `IPv4 network mask.`,
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
					Description: `Hide behind method. This parameter is not required in case "method" parameter is "static".`,
				},
				resource.Attribute{
					Name:        "install_on",
					Description: `Which gateway should apply the NAT translation.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `NAT translation method. ` + "`" + `management_blades` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "network_policy_management",
					Description: `Enable Network Policy Management.`,
				},
				resource.Attribute{
					Name:        "logging_and_status",
					Description: `Enable Logging & Status.`,
				},
				resource.Attribute{
					Name:        "smart_event_server",
					Description: `Enable SmartEvent server. When activating SmartEvent server, blades 'logging-and-status' and 'smart-event-correlation' should be set to True. To complete SmartEvent configuration, perform Install Database or Install Policy on your Security Management servers and Log servers. </br>Activating SmartEvent Server is not recommended in Management High Availability environment. For more information refer to sk25164.`,
				},
				resource.Attribute{
					Name:        "smart_event_correlation",
					Description: `Enable SmartEvent Correlation Unit.`,
				},
				resource.Attribute{
					Name:        "endpoint_policy",
					Description: `Enable Endpoint Policy. To complete Endpoint Security Management configuration, perform Install Database on your Endpoint Management Server. Field is not supported on Multi Domain Server environment.`,
				},
				resource.Attribute{
					Name:        "compliance",
					Description: `Compliance blade. Can be set when 'network-policy-management' was selected to be True.`,
				},
				resource.Attribute{
					Name:        "user_directory",
					Description: `Enable User Directory. Can be set when 'network-policy-management' was selected to be True.`,
				},
				resource.Attribute{
					Name:        "secondary",
					Description: `Secondary Management enabled.`,
				},
				resource.Attribute{
					Name:        "identity_logging",
					Description: `Identity Logging enabled. ` + "`" + `logs_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "free_disk_space_metrics",
					Description: `Free disk space metrics.`,
				},
				resource.Attribute{
					Name:        "accept_syslog_messages",
					Description: `Enable accept syslog messages.`,
				},
				resource.Attribute{
					Name:        "alert_when_free_disk_space_below",
					Description: `Enable alert when free disk space is below threshold.`,
				},
				resource.Attribute{
					Name:        "alert_when_free_disk_space_below_threshold",
					Description: `Alert when free disk space below threshold.`,
				},
				resource.Attribute{
					Name:        "alert_when_free_disk_space_below_type",
					Description: `Alert when free disk space below type.`,
				},
				resource.Attribute{
					Name:        "before_delete_keep_logs_from_the_last_days",
					Description: `Enable before delete keep logs from the last days.`,
				},
				resource.Attribute{
					Name:        "before_delete_keep_logs_from_the_last_days_threshold",
					Description: `Before delete keep logs from the last days threshold.`,
				},
				resource.Attribute{
					Name:        "before_delete_run_script",
					Description: `Enable Before delete run script.`,
				},
				resource.Attribute{
					Name:        "before_delete_run_script_command",
					Description: `Before delete run script command.`,
				},
				resource.Attribute{
					Name:        "delete_index_files_older_than_days",
					Description: `Enable delete index files older than days.`,
				},
				resource.Attribute{
					Name:        "delete_index_files_older_than_days_threshold",
					Description: `Delete index files older than days threshold.`,
				},
				resource.Attribute{
					Name:        "delete_when_free_disk_space_below",
					Description: `Enable delete when free disk space below.`,
				},
				resource.Attribute{
					Name:        "delete_when_free_disk_space_below_threshold",
					Description: `Delete when free disk space below threshold.`,
				},
				resource.Attribute{
					Name:        "detect_new_citrix_ica_application_names",
					Description: `Enable detect new citrix ica application names.`,
				},
				resource.Attribute{
					Name:        "enable_log_indexing",
					Description: `Enable log indexing.`,
				},
				resource.Attribute{
					Name:        "forward_logs_to_log_server",
					Description: `Enable forward logs to log server.`,
				},
				resource.Attribute{
					Name:        "forward_logs_to_log_server_name",
					Description: `Forward logs to log server name.`,
				},
				resource.Attribute{
					Name:        "forward_logs_to_log_server_schedule_name",
					Description: `Forward logs to log server schedule name.`,
				},
				resource.Attribute{
					Name:        "rotate_log_by_file_size",
					Description: `Enable rotate log by file size.`,
				},
				resource.Attribute{
					Name:        "rotate_log_file_size_threshold",
					Description: `Log file size threshold.`,
				},
				resource.Attribute{
					Name:        "rotate_log_on_schedule",
					Description: `Enable rotate log on schedule.`,
				},
				resource.Attribute{
					Name:        "rotate_log_schedule_name",
					Description: `Rotate log schedule name.`,
				},
				resource.Attribute{
					Name:        "smart_event_intro_correletion_unit",
					Description: `Enable SmartEvent intro correletion unit.`,
				},
				resource.Attribute{
					Name:        "stop_logging_when_free_disk_space_below",
					Description: `Enable stop logging when free disk space below.`,
				},
				resource.Attribute{
					Name:        "stop_logging_when_free_disk_space_below_threshold",
					Description: `Stop logging when free disk space below threshold.`,
				},
				resource.Attribute{
					Name:        "turn_on_qos_logging",
					Description: `Enable turn on qos logging.`,
				},
				resource.Attribute{
					Name:        "update_account_log_every",
					Description: `Update account log in every amount of seconds.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_data_access_layer",
			Category:         "Data Sources",
			ShortDescription: `This resource allows you to execute Check Point Access Layer.`,
			Description: `

Use this data source to get information on an existing Check Point Access Layer.

`,
			Keywords: []string{},
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
			Description: `

Use this data source to get information on an existing Check Point Access Role.

`,
			Keywords: []string{},
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
			Description: `

Use this data source to get information on an existing Check Point Access Rule.

`,
			Keywords: []string{},
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
			Description: `

Use this data source to get information on an existing Check Point Access Section.

`,
			Keywords: []string{},
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
			Description: `

Use this data source to get information on an existing Check Point Address Range.

`,
			Keywords: []string{},
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
			Description: `

Use this data source to get information on an existing Check Point Application Site.

`,
			Keywords: []string{},
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
			Description: `

Use this data source to get information on an existing Check Point Application Site Category.

`,
			Keywords: []string{},
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
			Description: `

  Use this data source to get information on an existing Check Point Application Site Group.

`,
			Keywords: []string{},
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
			Description: `

Use this data source to get information on an existing Check Point Dns Domain.

`,
			Keywords: []string{},
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
			Description: `

Use this data source to get information on an existing Check Point Dynamic Object.

`,
			Keywords: []string{},
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
			Description: `

Use this data source to get information on an existing Check Point Exception Group.

`,
			Keywords: []string{},
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
			Description: `

Use this data source to get information on an existing Check Point Group.

`,
			Keywords: []string{},
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
			Description: `

Use this data source to get information on an existing Check Point Group With Exclusion.

`,
			Keywords: []string{},
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
			Description: `

Use this data source to get information on an existing Check Point Host.

`,
			Keywords: []string{},
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
			Description: `

Use this data source to get information on an existing Check Point Https Layer.

`,
			Keywords: []string{},
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
			Description: `

Use this data source to get information on an existing Check Point Https Rule.

`,
			Keywords: []string{},
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
			Description: `

Use this data source to get information on an existing Check Point Https Section.

`,
			Keywords: []string{},
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
			Description: `

Use this data source to get information on an existing Check Point Multicast Address Range.

`,
			Keywords: []string{},
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
			Description: `

Use this data source to get information on an existing Check Point Network Object.

`,
			Keywords: []string{},
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
			Description: `

Use this data source to get information on an existing Check Point Opsec Application.

`,
			Keywords: []string{},
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
			Description: `

Use this data source to get information on an existing Check Point Package Object.

`,
			Keywords: []string{},
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
			Description: `

Use this data source to get information on an existing Check Point Security Zone.

`,
			Keywords: []string{},
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
			Description: `

Use this data source to get information on an existing Check Point Service Dce Rpc.

`,
			Keywords: []string{},
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
			Description: `

Use this data source to get information on an existing Check Point Service Group.

`,
			Keywords: []string{},
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
			Description: `

Use this data source to get information on an existing Check Point Service Icmp.

`,
			Keywords: []string{},
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
			Description: `

Use this data source to get information on an existing Check Point Service Icmp6.

`,
			Keywords: []string{},
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
			Description: `

Use this data source to get information on an existing Check Point Service Other.

`,
			Keywords: []string{},
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
			Description: `

Use this data source to get information on an existing Check Point Service Rpc.

`,
			Keywords: []string{},
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
			Description: `

Use this data source to get information on an existing Check Point Service Sctp.

`,
			Keywords: []string{},
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
			Description: `

Use this data source to get information on an existing Check Point Service Tcp.

`,
			Keywords: []string{},
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
			Description: `

Use this data source to get information on an existing Check Point Service Udp.

`,
			Keywords: []string{},
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
			Description: `

Use this data source to get information on an existing Check Point Threat Indicator.

`,
			Keywords: []string{},
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
			Description: `

Use this data source to get information on an existing Check Point Time Group.

`,
			Keywords: []string{},
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
			Description: `

Use this data source to get information on an existing Check Point Vpn Community Meshed.

`,
			Keywords: []string{},
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
			Description: `

Use this data source to get information on an existing Check Point Vpn Community Star.

`,
			Keywords: []string{},
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
			Description: `

Use this data source to get information on an existing Check Point Wildcard.

`,
			Keywords: []string{},
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
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_gsn_handover_group",
			Category:         "Data Sources",
			ShortDescription: `This resource allows you to execute Check Point Gsn Handover Group.`,
			Description: `

This resource allows you to execute Check Point Gsn Handover Group.

`,
			Keywords: []string{},
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
					Name:        "enforce_gtp",
					Description: `Enable enforce GTP signal packet rate limit from this group.`,
				},
				resource.Attribute{
					Name:        "gtp_rate",
					Description: `Limit of the GTP rate in PDU/sec.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `Collection of GSN handover group members identified by the name or UID.members blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object.`,
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
			Type:             "checkpoint_management_identity_tag",
			Category:         "Data Sources",
			ShortDescription: `This resource allows you to execute Check Point Identity Tag.`,
			Description: `

This resource allows you to execute Check Point Identity Tag.

`,
			Keywords: []string{},
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
					Name:        "external_identifier",
					Description: `External identifier. For example: Cisco ISE security group tag.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object.`,
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
			Type:             "checkpoint_management_mds",
			Category:         "Data Sources",
			ShortDescription: `This resource allows you to execute Check Point MDS.`,
			Description: `

This resource allows you to execute Check Point MDS.

`,
			Keywords: []string{},
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
					Name:        "hardware",
					Description: `Hardware name. For example: Open server, Smart-1, Other.`,
				},
				resource.Attribute{
					Name:        "os",
					Description: `Operating system name. For example: Gaia, Linux, SecurePlatform.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `System version.`,
				},
				resource.Attribute{
					Name:        "ip_pool_first",
					Description: `First IP address in the range.`,
				},
				resource.Attribute{
					Name:        "ip_pool_last",
					Description: `Last IP address in the range.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string.`,
				},
				resource.Attribute{
					Name:        "server_type",
					Description: `Type of the management server.`,
				},
				resource.Attribute{
					Name:        "sic_name",
					Description: `Name of the Secure Internal Connection Trust..`,
				},
				resource.Attribute{
					Name:        "sic_state",
					Description: `State the Secure Internal Connection Trust..`,
				},
				resource.Attribute{
					Name:        "domains",
					Description: `Collection of Domain objects identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "global_domains",
					Description: `Collection of Global domain objects identified by the name or UID.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_nat_rule",
			Category:         "Data Sources",
			ShortDescription: `This resource allows you to execute Check Point NAT Rule.`,
			Description: `

This resource allows you to execute Check Point NAT Rule.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "package",
					Description: `(Required) Name of the package.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Rule name.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enable/Disable the rule.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `Nat method.`,
				},
				resource.Attribute{
					Name:        "install_on",
					Description: `Which Gateways identified by the name or UID to install the policy on.`,
				},
				resource.Attribute{
					Name:        "original_destination",
					Description: `Original destination.`,
				},
				resource.Attribute{
					Name:        "original_service",
					Description: `Original service.`,
				},
				resource.Attribute{
					Name:        "original_source",
					Description: `Original source.`,
				},
				resource.Attribute{
					Name:        "translated_destination",
					Description: `Translated destination.`,
				},
				resource.Attribute{
					Name:        "translated_service",
					Description: `Translated service.`,
				},
				resource.Attribute{
					Name:        "translated_source",
					Description: `Translated source.`,
				},
				resource.Attribute{
					Name:        "auto_generated",
					Description: `Auto generated.`,
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
			Type:             "checkpoint_management_nat_section",
			Category:         "Data Sources",
			ShortDescription: `This resource allows you to execute Check Point NAT section.`,
			Description: `

This resource allows you to execute Check Point NAT section.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "package",
					Description: `(Required) Name of the package.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_service_citrix_tcp",
			Category:         "Data Sources",
			ShortDescription: `This resource allows you to execute Check Point Service Citrix Tcp.`,
			Description: `

This resource allows you to execute Check Point Service Citrix Tcp.

`,
			Keywords: []string{},
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
					Name:        "application",
					Description: `Citrix application name.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object.`,
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
			Type:             "checkpoint_management_service_compound_tcp",
			Category:         "Data Sources",
			ShortDescription: `This resource allows you to execute Check Point Service Compound Tcp.`,
			Description: `

This resource allows you to execute Check Point Service Compound Tcp.

`,
			Keywords: []string{},
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
					Name:        "compound_service",
					Description: `Compound service type.`,
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
					Description: `Color of the object.`,
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
			Type:             "checkpoint_management_show_objects",
			Category:         "Data Sources",
			ShortDescription: `This resource allows you to execute Check Point Show Objects.`,
			Description: `

This resource allows you to execute Check Point Show Objects.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Search expression to filter objects by. The provided text should be exactly the same as it would be given in Smart Console. The logical operators in the expression ('AND', 'OR') should be provided in capital letters. By default, the search involves both a textual search and a IP search. To use IP search only, set the "ip-only" parameter to true.`,
				},
				resource.Attribute{
					Name:        "ip_only",
					Description: `(Optional) If using "filter", use this field to search objects by their IP address only, without involving the textual search.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The objects' type, e.g.: host, service-tcp, network, address-range.`,
				},
				resource.Attribute{
					Name:        "limit",
					Description: `(Optional) The maximal number of returned results.`,
				},
				resource.Attribute{
					Name:        "offset",
					Description: `(Optional) Number of the results to initially skip.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Optional) Sorts the results by search criteria. Automatically sorts the results by Name, in the ascending order. order blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "from",
					Description: `From which element number the query was done.`,
				},
				resource.Attribute{
					Name:        "to",
					Description: `To which element number the query was done.`,
				},
				resource.Attribute{
					Name:        "total",
					Description: `Total number of elements returned by the query.`,
				},
				resource.Attribute{
					Name:        "objects",
					Description: `Collection of retrieved objects. objects blocks blocks are documented below. ` + "`" + `order` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "asc",
					Description: `(Optional) Sorts results by the given field in ascending order.`,
				},
				resource.Attribute{
					Name:        "desc",
					Description: `(Optional) Sorts results by the given field in descending order. ` + "`" + `objects` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Object name. Must be unique in the domain.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Object type.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Information about the domain that holds the Object. domain blocks are documented below. ` + "`" + `domain` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Object name. Must be unique in the domain.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "domain_type",
					Description: `Domain type.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_show_updatable_objects_repository_content",
			Category:         "Data Sources",
			ShortDescription: `This resource allows you to execute Check Point Show Updatable Objects Repository Content.`,
			Description: `

This resource allows you to execute Check Point Show Updatable Objects Repository Content.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "uid_in_updatable_objects_repository",
					Description: `(Optional) The object's unique identifier in the Updatable Objects repository.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Return results matching the specified filter. filter blocks blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "limit",
					Description: `(Optional) The maximal number of returned results.`,
				},
				resource.Attribute{
					Name:        "offset",
					Description: `(Optional) Number of the results to initially skip.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Optional) Sorts the results by search criteria. Automatically sorts the results by Name, in the ascending order. order blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "from",
					Description: `From which element number the query was done.`,
				},
				resource.Attribute{
					Name:        "to",
					Description: `To which element number the query was done.`,
				},
				resource.Attribute{
					Name:        "total",
					Description: `Total number of elements returned by the query.`,
				},
				resource.Attribute{
					Name:        "objects",
					Description: `Collection of retrieved Updatable Objects. objects blocks blocks are documented below. ` + "`" + `filter` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "text",
					Description: `(Optional) Return results containing the specified text value.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `(Optional) Return results under the specified uri value.`,
				},
				resource.Attribute{
					Name:        "parent_uid_in_updatable_objects_repository",
					Description: `(Optional) Return results under the specified Updatable Object. ` + "`" + `order` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "asc",
					Description: `(Optional) Sorts results by the given field in ascending order.`,
				},
				resource.Attribute{
					Name:        "desc",
					Description: `(Optional) Sorts results by the given field in descending order. ` + "`" + `objects` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name_in_updatable_objects_repository",
					Description: `Object name in the Updatable Objects Repository.`,
				},
				resource.Attribute{
					Name:        "uid_in_updatable_objects_repository",
					Description: `Unique identifier of the object in the Updatable Objects Repository.`,
				},
				resource.Attribute{
					Name:        "additional_properties",
					Description: `Additional properties on the object. additional_properties blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "updatable_object",
					Description: `The imported management object (if exists). updatable_object blocks are documented below. ` + "`" + `additional_properties` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of retrieved Updatable Object.`,
				},
				resource.Attribute{
					Name:        "info_text",
					Description: `Information about the Updatable Object IP ranges source.`,
				},
				resource.Attribute{
					Name:        "info_url",
					Description: `URL of the Updatable Object IP ranges source.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `URI of the Updatable Object under the Updatable Objects Repository. ` + "`" + `updatable_object` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Object name. Must be unique in the domain.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Object type.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Information about the domain that holds the Object. domain blocks are documented below. ` + "`" + `domain` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Object name. Must be unique in the domain.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "domain_type",
					Description: `Domain type.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_simple_cluster",
			Category:         "Data Sources",
			ShortDescription: `This resource allows you to execute Check Point Simple Cluster.`,
			Description: `

This resource allows you to execute Check Point Simple Cluster.

`,
			Keywords: []string{},
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
					Name:        "cluster_mode",
					Description: `Cluster mode.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `Cluster interfaces. interfaces blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `Cluster members. members blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "anti_bot",
					Description: `Anti-Bot blade enabled.`,
				},
				resource.Attribute{
					Name:        "anti_virus",
					Description: `Anti-Virus blade enabled.`,
				},
				resource.Attribute{
					Name:        "application_control",
					Description: `Application Control blade enabled.`,
				},
				resource.Attribute{
					Name:        "content_awareness",
					Description: `Content Awareness blade enabled.`,
				},
				resource.Attribute{
					Name:        "data_awareness",
					Description: `Data Awareness blade enabled.`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `Intrusion Prevention System blade enabled.`,
				},
				resource.Attribute{
					Name:        "threat_emulation",
					Description: `Threat Emulation blade enabled.`,
				},
				resource.Attribute{
					Name:        "url_filtering",
					Description: `URL Filtering blade enabled.`,
				},
				resource.Attribute{
					Name:        "firewall",
					Description: `Firewall blade enabled.`,
				},
				resource.Attribute{
					Name:        "firewall_settings",
					Description: `Firewall settings. firewall_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "vpn",
					Description: `VPN blade enabled.`,
				},
				resource.Attribute{
					Name:        "vpn_settings",
					Description: `Cluster VPN settings. vpn_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "dynamic_ip",
					Description: `Dynamic IP address.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Cluster platform version.`,
				},
				resource.Attribute{
					Name:        "os_name",
					Description: `Cluster Operating system name.`,
				},
				resource.Attribute{
					Name:        "hardware",
					Description: `Cluster platform hardware name.`,
				},
				resource.Attribute{
					Name:        "one_time_password",
					Description: `Secure Internal Communication one time password.`,
				},
				resource.Attribute{
					Name:        "sic_name",
					Description: `Secure Internal Communication name.`,
				},
				resource.Attribute{
					Name:        "sic_state",
					Description: `Secure Internal Communication state.`,
				},
				resource.Attribute{
					Name:        "save_logs_locally",
					Description: `Enable save logs locally.`,
				},
				resource.Attribute{
					Name:        "send_alerts_to_server",
					Description: `Collection of Server(s) to send alerts to identified by the name.`,
				},
				resource.Attribute{
					Name:        "send_logs_to_backup_server",
					Description: `Collection of Backup server(s) to send logs to identified by the name.`,
				},
				resource.Attribute{
					Name:        "send_logs_to_server",
					Description: `Collection of Server(s) to send logs to identified by the name.`,
				},
				resource.Attribute{
					Name:        "logs_settings",
					Description: `Logs settings. logs_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tags identified by name. ` + "`" + `interfaces` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Interface name.`,
				},
				resource.Attribute{
					Name:        "interface_type",
					Description: `Cluster interface type.`,
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
					Name:        "ipv4_network_mask",
					Description: `IPv4 network address.`,
				},
				resource.Attribute{
					Name:        "ipv6_network_mask",
					Description: `IPv6 network address.`,
				},
				resource.Attribute{
					Name:        "ipv4_mask_length",
					Description: `IPv4 network mask length.`,
				},
				resource.Attribute{
					Name:        "ipv6_mask_length",
					Description: `IPv6 network mask length.`,
				},
				resource.Attribute{
					Name:        "anti_spoofing",
					Description: `Anti spoofing.`,
				},
				resource.Attribute{
					Name:        "anti_spoofing_settings",
					Description: `Anti spoofing settings. anti_spoofing_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "multicast_address",
					Description: `Multicast IP Address.`,
				},
				resource.Attribute{
					Name:        "multicast_address_type",
					Description: `Multicast Address Type.`,
				},
				resource.Attribute{
					Name:        "security_zone",
					Description: `Security zone.`,
				},
				resource.Attribute{
					Name:        "security_zone_settings",
					Description: `Security zone settings. security_zone_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "topology",
					Description: `Topology.`,
				},
				resource.Attribute{
					Name:        "topology_settings",
					Description: `Topology settings. topology_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "topology_automatic_calculation",
					Description: `Shows the automatic topology calculation..`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string. ` + "`" + `anti_spoofing_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `If packets will be rejected (the Prevent option) or whether the packets will be monitored (the Detect option). ` + "`" + `security_zone_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "auto_calculated",
					Description: `Security Zone is calculated according to where the interface leads to.`,
				},
				resource.Attribute{
					Name:        "specific_zone",
					Description: `Security Zone specified manually. ` + "`" + `topology_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "interface_leads_to_dmz",
					Description: `Whether this interface leads to demilitarized zone (perimeter network).`,
				},
				resource.Attribute{
					Name:        "ip_address_behind_this_interface",
					Description: `Ip address behind this interface.`,
				},
				resource.Attribute{
					Name:        "specific_network",
					Description: `Network behind this interface. ` + "`" + `members` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Object name. Should be unique in the domain..`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `IPv4 or IPv6 address.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `Cluster Member network interfaces. interfaces blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "one_time_password",
					Description: `Secure Internal Communication one time password.`,
				},
				resource.Attribute{
					Name:        "sic_name",
					Description: `Secure Internal Communication name.`,
				},
				resource.Attribute{
					Name:        "sic_message",
					Description: `Secure Internal Communication state. ` + "`" + `interfaces` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Interface name.`,
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
					Name:        "ipv4_network_mask",
					Description: `IPv4 network address.`,
				},
				resource.Attribute{
					Name:        "ipv6_network_mask",
					Description: `IPv6 network address.`,
				},
				resource.Attribute{
					Name:        "ipv4_mask_length",
					Description: `IPv4 network mask length.`,
				},
				resource.Attribute{
					Name:        "ipv6_mask_length",
					Description: `IPv6 network mask length. ` + "`" + `firewall_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "auto_calculate_connections_hash_table_size_and_memory_pool",
					Description: `Auto calculate connections hash table size and memory pool.`,
				},
				resource.Attribute{
					Name:        "auto_maximum_limit_for_concurrent_connections",
					Description: `Auto maximum limit for concurrent connections.`,
				},
				resource.Attribute{
					Name:        "connections_hash_size",
					Description: `Connections hash size.`,
				},
				resource.Attribute{
					Name:        "maximum_limit_for_concurrent_connections",
					Description: `Maximum limit for concurrent connections.`,
				},
				resource.Attribute{
					Name:        "maximum_memory_pool_size",
					Description: `Maximum memory pool size.`,
				},
				resource.Attribute{
					Name:        "memory_pool_size",
					Description: `Memory pool size. ` + "`" + `vpn_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "authentication",
					Description: `authentication blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "link_selection",
					Description: `Link selection blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "maximum_concurrent_ike_negotiations",
					Description: `Maximum concurrent ike negotiations.`,
				},
				resource.Attribute{
					Name:        "maximum_concurrent_tunnels",
					Description: `Maximum concurrent tunnels.`,
				},
				resource.Attribute{
					Name:        "office_mode",
					Description: `Office Mode. Notation Wide Impact - Office Mode apply IPSec VPN Software Blade clients and to the Mobile Access Software Blade clients. office_mode blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "remote_access",
					Description: `remote_access blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "vpn_domain",
					Description: `Gateway VPN domain identified by the name.`,
				},
				resource.Attribute{
					Name:        "vpn_domain_type",
					Description: `Gateway VPN domain type. ` + "`" + `authentication` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "authentication_clients",
					Description: `Collection of VPN Authentication clients identified by the name. ` + "`" + `link_selection` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "ip_selection",
					Description: `IP selection.`,
				},
				resource.Attribute{
					Name:        "dns_resolving_hostname",
					Description: `DNS Resolving Hostname. Must be set when "ip-selection" was selected to be "dns-resolving-from-hostname".`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `IP Address. Must be set when "ip-selection" was selected to be "use-selected-address-from-topology" or "use-statically-nated-ip". ` + "`" + `office_mode` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Office Mode Permissions. When selected to be "off", all the other definitions are irrelevant.`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `Group. Identified by name. Must be set when "office-mode-permissions" was selected to be "group".`,
				},
				resource.Attribute{
					Name:        "allocate_ip_address_from",
					Description: `Allocate IP address Method. Allocate IP address by sequentially trying the given methods until success. allocate_ip_address_from blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "support_multiple_interfaces",
					Description: `Support connectivity enhancement for gateways with multiple external interfaces.`,
				},
				resource.Attribute{
					Name:        "perform_anti_spoofing",
					Description: `Perform Anti-Spoofing on Office Mode addresses.`,
				},
				resource.Attribute{
					Name:        "anti_spoofing_additional_addresses",
					Description: `Additional IP Addresses for Anti-Spoofing. Identified by name. Must be set when "perform-anti-spoofings" is true. ` + "`" + `allocate_ip_address_from` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "radius_server",
					Description: `Radius server used to authenticate the user.`,
				},
				resource.Attribute{
					Name:        "use_allocate_method",
					Description: `Use Allocate Method.`,
				},
				resource.Attribute{
					Name:        "allocate_method",
					Description: `Using either Manual (IP Pool) or Automatic (DHCP). Must be set when "use-allocate-method" is true.`,
				},
				resource.Attribute{
					Name:        "manual_network",
					Description: `Manual Network. Identified by name. Must be set when "allocate-method" was selected to be "manual".`,
				},
				resource.Attribute{
					Name:        "dhcp_server",
					Description: `DHCP Server. Identified by name. Must be set when "allocate-method" was selected to be "automatic".`,
				},
				resource.Attribute{
					Name:        "virtual_ip_address",
					Description: `Virtual IPV4 address for DHCP server replies. Must be set when "allocate-method" was selected to be "automatic".`,
				},
				resource.Attribute{
					Name:        "dhcp_mac_address",
					Description: `Calculated MAC address for DHCP allocation. Must be set when "allocate-method" was selected to be "automatic".`,
				},
				resource.Attribute{
					Name:        "optional_parameters",
					Description: `This configuration applies to all Office Mode methods except Automatic (using DHCP) and ipassignment.conf entries which contain this data. optional_parameters blocks are documented below. ` + "`" + `optional_parameters` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "use_primary_dns_server",
					Description: `Use Primary DNS Server.`,
				},
				resource.Attribute{
					Name:        "primary_dns_server",
					Description: `Primary DNS Server. Identified by name. Must be set when "use-primary-dns-server" is true and can not be set when "use-primary-dns-server" is false.`,
				},
				resource.Attribute{
					Name:        "use_first_backup_dns_server",
					Description: `Use First Backup DNS Server.`,
				},
				resource.Attribute{
					Name:        "first_backup_dns_server",
					Description: `First Backup DNS Server. Identified by name. Must be set when "use-first-backup-dns-server" is true and can not be set when "use-first-backup-dns-server" is false.`,
				},
				resource.Attribute{
					Name:        "use_second_backup_dns_server",
					Description: `Use Second Backup DNS Server.`,
				},
				resource.Attribute{
					Name:        "second_backup_dns_server",
					Description: `Second Backup DNS Server. Identified by name. Must be set when "use-second-backup-dns-server" is true and can not be set when "use-second-backup-dns-server" is false.`,
				},
				resource.Attribute{
					Name:        "dns_suffixes",
					Description: `DNS Suffixes.`,
				},
				resource.Attribute{
					Name:        "use_primary_wins_server",
					Description: `Use Primary WINS Server.`,
				},
				resource.Attribute{
					Name:        "primary_wins_server",
					Description: `Primary WINS Server. Identified by name. Must be set when "use-primary-wins-server" is true and can not be set when "use-primary-wins-server" is false.`,
				},
				resource.Attribute{
					Name:        "use_first_backup_wins_server",
					Description: `Use First Backup WINS Server.`,
				},
				resource.Attribute{
					Name:        "first_backup_wins_server",
					Description: `First Backup WINS Server. Identified by name. Must be set when "use-first-backup-wins-server" is true and can not be set when "use-first-backup-wins-server" is false.`,
				},
				resource.Attribute{
					Name:        "use_second_backup_wins_server",
					Description: `Use Second Backup WINS Server.`,
				},
				resource.Attribute{
					Name:        "second_backup_wins_server",
					Description: `Second Backup WINS Server. Identified by name. Must be set when "use-second-backup-wins-server" is true and can not be set when "use-second-backup-wins-server" is false.`,
				},
				resource.Attribute{
					Name:        "ip_lease_duration",
					Description: `IP Lease Duration in Minutes. The value must be in the range 2-32767. ` + "`" + `remote_access` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "support_l2tp",
					Description: `Support L2TP (relevant only when office mode is active).`,
				},
				resource.Attribute{
					Name:        "l2tp_auth_method",
					Description: `L2TP Authentication Method. Must be set when "support-l2tp" is true.`,
				},
				resource.Attribute{
					Name:        "l2tp_certificate",
					Description: `L2TP Certificate. Must be set when "l2tp-auth-method" was selected to be "certificate". Insert "defaultCert" when you want to use the default certificate.`,
				},
				resource.Attribute{
					Name:        "allow_vpn_clients_to_route_traffic",
					Description: `Allow VPN clients to route traffic.`,
				},
				resource.Attribute{
					Name:        "support_nat_traversal_mechanism",
					Description: `Support NAT traversal mechanism (UDP encapsulation).`,
				},
				resource.Attribute{
					Name:        "nat_traversal_service",
					Description: `Allocated NAT traversal UDP service. Identified by name. Must be set when "support-nat-traversal-mechanism" is true.`,
				},
				resource.Attribute{
					Name:        "support_visitor_mode",
					Description: `Support Visitor Mode.`,
				},
				resource.Attribute{
					Name:        "visitor_mode_service",
					Description: `TCP Service for Visitor Mode. Identified by name. Must be set when "support-visitor-mode" is true.`,
				},
				resource.Attribute{
					Name:        "visitor_mode_interface",
					Description: `Interface for Visitor Mode. Must be set when "support-visitor-mode" is true. Insert IPV4 Address of existing interface or "All IPs" when you want all interfaces.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_simple_gateway",
			Category:         "Data Sources",
			ShortDescription: `This resource allows you to execute Check Point Simple Gateway.`,
			Description: `

This resource allows you to execute Check Point Simple Gateway.

`,
			Keywords: []string{},
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
					Name:        "interfaces",
					Description: `Gateway interfaces. interfaces blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "anti_bot",
					Description: `Anti-Bot blade enabled.`,
				},
				resource.Attribute{
					Name:        "anti_virus",
					Description: `Anti-Virus blade enabled.`,
				},
				resource.Attribute{
					Name:        "application_control",
					Description: `Application Control blade enabled.`,
				},
				resource.Attribute{
					Name:        "content_awareness",
					Description: `Content Awareness blade enabled.`,
				},
				resource.Attribute{
					Name:        "icap_server",
					Description: `ICAP Server enabled.`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `Intrusion Prevention System blade enabled.`,
				},
				resource.Attribute{
					Name:        "threat_emulation",
					Description: `Threat Emulation blade enabled.`,
				},
				resource.Attribute{
					Name:        "threat_extraction",
					Description: `Threat Extraction blade enabled.`,
				},
				resource.Attribute{
					Name:        "url_filtering",
					Description: `URL Filtering blade enabled.`,
				},
				resource.Attribute{
					Name:        "firewall",
					Description: `Firewall blade enabled.`,
				},
				resource.Attribute{
					Name:        "firewall_settings",
					Description: `Firewall settings. firewall_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "vpn",
					Description: `VPN blade enabled.`,
				},
				resource.Attribute{
					Name:        "vpn_settings",
					Description: `Gateway VPN settings. vpn_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "dynamic_ip",
					Description: `Dynamic IP address.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Gateway platform version.`,
				},
				resource.Attribute{
					Name:        "os_name",
					Description: `Operating system name.`,
				},
				resource.Attribute{
					Name:        "hardware",
					Description: `Gateway platform hardware name.`,
				},
				resource.Attribute{
					Name:        "one_time_password",
					Description: `Secure internal connection one time password.`,
				},
				resource.Attribute{
					Name:        "sic_name",
					Description: `Secure Internal Communication name.`,
				},
				resource.Attribute{
					Name:        "sic_state",
					Description: `Secure Internal Communication state.`,
				},
				resource.Attribute{
					Name:        "save_logs_locally",
					Description: `Enable save logs locally.`,
				},
				resource.Attribute{
					Name:        "send_alerts_to_server",
					Description: `Collection of Server(s) to send alerts to identified by the name.`,
				},
				resource.Attribute{
					Name:        "send_logs_to_backup_server",
					Description: `Collection of Backup server(s) to send logs to identified by the name.`,
				},
				resource.Attribute{
					Name:        "send_logs_to_server",
					Description: `Collection of Server(s) to send logs to identified by the name.`,
				},
				resource.Attribute{
					Name:        "logs_settings",
					Description: `Logs settings. logs_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tags identified by name. ` + "`" + `interfaces` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Interface name.`,
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
					Name:        "ipv4_network_mask",
					Description: `IPv4 network address.`,
				},
				resource.Attribute{
					Name:        "ipv6_network_mask",
					Description: `IPv6 network address.`,
				},
				resource.Attribute{
					Name:        "ipv4_mask_length",
					Description: `IPv4 network mask length.`,
				},
				resource.Attribute{
					Name:        "ipv6_mask_length",
					Description: `IPv6 network mask length.`,
				},
				resource.Attribute{
					Name:        "anti_spoofing",
					Description: `Anti spoofing.`,
				},
				resource.Attribute{
					Name:        "anti_spoofing_settings",
					Description: `Anti spoofing settings. anti_spoofing_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "security_zone",
					Description: `Security zone.`,
				},
				resource.Attribute{
					Name:        "security_zone_settings",
					Description: `Security zone settings. security_zone_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "topology",
					Description: `Topology.`,
				},
				resource.Attribute{
					Name:        "topology_settings",
					Description: `Topology settings. topology_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "topology_automatic_calculation",
					Description: `Shows the automatic topology calculation..`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string. ` + "`" + `anti_spoofing_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `If packets will be rejected (the Prevent option) or whether the packets will be monitored (the Detect option). ` + "`" + `security_zone_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "auto_calculated",
					Description: `Security Zone is calculated according to where the interface leads to.`,
				},
				resource.Attribute{
					Name:        "specific_zone",
					Description: `Security Zone specified manually. ` + "`" + `topology_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "interface_leads_to_dmz",
					Description: `Whether this interface leads to demilitarized zone (perimeter network).`,
				},
				resource.Attribute{
					Name:        "ip_address_behind_this_interface",
					Description: `Ip address behind this interface.`,
				},
				resource.Attribute{
					Name:        "specific_network",
					Description: `Network behind this interface. ` + "`" + `firewall_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "auto_calculate_connections_hash_table_size_and_memory_pool",
					Description: `Auto calculate connections hash table size and memory pool.`,
				},
				resource.Attribute{
					Name:        "auto_maximum_limit_for_concurrent_connections",
					Description: `Auto maximum limit for concurrent connections.`,
				},
				resource.Attribute{
					Name:        "connections_hash_size",
					Description: `Connections hash size.`,
				},
				resource.Attribute{
					Name:        "maximum_limit_for_concurrent_connections",
					Description: `Maximum limit for concurrent connections.`,
				},
				resource.Attribute{
					Name:        "maximum_memory_pool_size",
					Description: `Maximum memory pool size.`,
				},
				resource.Attribute{
					Name:        "memory_pool_size",
					Description: `Memory pool size. ` + "`" + `vpn_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "authentication",
					Description: `authentication blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "link_selection",
					Description: `Link selection blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "maximum_concurrent_ike_negotiations",
					Description: `Maximum concurrent ike negotiations.`,
				},
				resource.Attribute{
					Name:        "maximum_concurrent_tunnels",
					Description: `Maximum concurrent tunnels.`,
				},
				resource.Attribute{
					Name:        "office_mode",
					Description: `Office Mode. Notation Wide Impact - Office Mode apply IPSec VPN Software Blade clients and to the Mobile Access Software Blade clients. office_mode blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "remote_access",
					Description: `remote_access blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "vpn_domain",
					Description: `Gateway VPN domain identified by the name.`,
				},
				resource.Attribute{
					Name:        "vpn_domain_type",
					Description: `Gateway VPN domain type. ` + "`" + `authentication` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "authentication_clients",
					Description: `Collection of VPN Authentication clients identified by the name. ` + "`" + `link_selection` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "ip_selection",
					Description: `IP selection.`,
				},
				resource.Attribute{
					Name:        "dns_resolving_hostname",
					Description: `DNS Resolving Hostname. Must be set when "ip-selection" was selected to be "dns-resolving-from-hostname".`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `IP Address. Must be set when "ip-selection" was selected to be "use-selected-address-from-topology" or "use-statically-nated-ip". ` + "`" + `office_mode` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Office Mode Permissions. When selected to be "off", all the other definitions are irrelevant.`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `Group. Identified by name. Must be set when "office-mode-permissions" was selected to be "group".`,
				},
				resource.Attribute{
					Name:        "allocate_ip_address_from",
					Description: `Allocate IP address Method. Allocate IP address by sequentially trying the given methods until success. allocate_ip_address_from blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "support_multiple_interfaces",
					Description: `Support connectivity enhancement for gateways with multiple external interfaces.`,
				},
				resource.Attribute{
					Name:        "perform_anti_spoofing",
					Description: `Perform Anti-Spoofing on Office Mode addresses.`,
				},
				resource.Attribute{
					Name:        "anti_spoofing_additional_addresses",
					Description: `Additional IP Addresses for Anti-Spoofing. Identified by name. Must be set when "perform-anti-spoofings" is true. ` + "`" + `allocate_ip_address_from` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "radius_server",
					Description: `Radius server used to authenticate the user.`,
				},
				resource.Attribute{
					Name:        "use_allocate_method",
					Description: `Use Allocate Method.`,
				},
				resource.Attribute{
					Name:        "allocate_method",
					Description: `Using either Manual (IP Pool) or Automatic (DHCP). Must be set when "use-allocate-method" is true.`,
				},
				resource.Attribute{
					Name:        "manual_network",
					Description: `Manual Network. Identified by name. Must be set when "allocate-method" was selected to be "manual".`,
				},
				resource.Attribute{
					Name:        "dhcp_server",
					Description: `DHCP Server. Identified by name. Must be set when "allocate-method" was selected to be "automatic".`,
				},
				resource.Attribute{
					Name:        "virtual_ip_address",
					Description: `Virtual IPV4 address for DHCP server replies. Must be set when "allocate-method" was selected to be "automatic".`,
				},
				resource.Attribute{
					Name:        "dhcp_mac_address",
					Description: `Calculated MAC address for DHCP allocation. Must be set when "allocate-method" was selected to be "automatic".`,
				},
				resource.Attribute{
					Name:        "optional_parameters",
					Description: `This configuration applies to all Office Mode methods except Automatic (using DHCP) and ipassignment.conf entries which contain this data. optional_parameters blocks are documented below. ` + "`" + `optional_parameters` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "use_primary_dns_server",
					Description: `Use Primary DNS Server.`,
				},
				resource.Attribute{
					Name:        "primary_dns_server",
					Description: `Primary DNS Server. Identified by name. Must be set when "use-primary-dns-server" is true and can not be set when "use-primary-dns-server" is false.`,
				},
				resource.Attribute{
					Name:        "use_first_backup_dns_server",
					Description: `Use First Backup DNS Server.`,
				},
				resource.Attribute{
					Name:        "first_backup_dns_server",
					Description: `First Backup DNS Server. Identified by name. Must be set when "use-first-backup-dns-server" is true and can not be set when "use-first-backup-dns-server" is false.`,
				},
				resource.Attribute{
					Name:        "use_second_backup_dns_server",
					Description: `Use Second Backup DNS Server.`,
				},
				resource.Attribute{
					Name:        "second_backup_dns_server",
					Description: `Second Backup DNS Server. Identified by name. Must be set when "use-second-backup-dns-server" is true and can not be set when "use-second-backup-dns-server" is false.`,
				},
				resource.Attribute{
					Name:        "dns_suffixes",
					Description: `DNS Suffixes.`,
				},
				resource.Attribute{
					Name:        "use_primary_wins_server",
					Description: `Use Primary WINS Server.`,
				},
				resource.Attribute{
					Name:        "primary_wins_server",
					Description: `Primary WINS Server. Identified by name. Must be set when "use-primary-wins-server" is true and can not be set when "use-primary-wins-server" is false.`,
				},
				resource.Attribute{
					Name:        "use_first_backup_wins_server",
					Description: `Use First Backup WINS Server.`,
				},
				resource.Attribute{
					Name:        "first_backup_wins_server",
					Description: `First Backup WINS Server. Identified by name. Must be set when "use-first-backup-wins-server" is true and can not be set when "use-first-backup-wins-server" is false.`,
				},
				resource.Attribute{
					Name:        "use_second_backup_wins_server",
					Description: `Use Second Backup WINS Server.`,
				},
				resource.Attribute{
					Name:        "second_backup_wins_server",
					Description: `Second Backup WINS Server. Identified by name. Must be set when "use-second-backup-wins-server" is true and can not be set when "use-second-backup-wins-server" is false.`,
				},
				resource.Attribute{
					Name:        "ip_lease_duration",
					Description: `IP Lease Duration in Minutes. The value must be in the range 2-32767. ` + "`" + `remote_access` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "support_l2tp",
					Description: `Support L2TP (relevant only when office mode is active).`,
				},
				resource.Attribute{
					Name:        "l2tp_auth_method",
					Description: `L2TP Authentication Method. Must be set when "support-l2tp" is true.`,
				},
				resource.Attribute{
					Name:        "l2tp_certificate",
					Description: `L2TP Certificate. Must be set when "l2tp-auth-method" was selected to be "certificate". Insert "defaultCert" when you want to use the default certificate.`,
				},
				resource.Attribute{
					Name:        "allow_vpn_clients_to_route_traffic",
					Description: `Allow VPN clients to route traffic.`,
				},
				resource.Attribute{
					Name:        "support_nat_traversal_mechanism",
					Description: `Support NAT traversal mechanism (UDP encapsulation).`,
				},
				resource.Attribute{
					Name:        "nat_traversal_service",
					Description: `Allocated NAT traversal UDP service. Identified by name. Must be set when "support-nat-traversal-mechanism" is true.`,
				},
				resource.Attribute{
					Name:        "support_visitor_mode",
					Description: `Support Visitor Mode.`,
				},
				resource.Attribute{
					Name:        "visitor_mode_service",
					Description: `TCP Service for Visitor Mode. Identified by name. Must be set when "support-visitor-mode" is true.`,
				},
				resource.Attribute{
					Name:        "visitor_mode_interface",
					Description: `Interface for Visitor Mode. Must be set when "support-visitor-mode" is true. Insert IPV4 Address of existing interface or "All IPs" when you want all interfaces. ` + "`" + `logs_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "free_disk_space_metrics",
					Description: `Free disk space metrics.`,
				},
				resource.Attribute{
					Name:        "accept_syslog_messages",
					Description: `Enable accept syslog messages.`,
				},
				resource.Attribute{
					Name:        "alert_when_free_disk_space_below",
					Description: `Enable alert when free disk space is below threshold.`,
				},
				resource.Attribute{
					Name:        "alert_when_free_disk_space_below_threshold",
					Description: `Alert when free disk space below threshold.`,
				},
				resource.Attribute{
					Name:        "alert_when_free_disk_space_below_type",
					Description: `Alert when free disk space below type.`,
				},
				resource.Attribute{
					Name:        "before_delete_keep_logs_from_the_last_days",
					Description: `Enable before delete keep logs from the last days.`,
				},
				resource.Attribute{
					Name:        "before_delete_keep_logs_from_the_last_days_threshold",
					Description: `Before delete keep logs from the last days threshold.`,
				},
				resource.Attribute{
					Name:        "before_delete_run_script",
					Description: `Enable Before delete run script.`,
				},
				resource.Attribute{
					Name:        "before_delete_run_script_command",
					Description: `Before delete run script command.`,
				},
				resource.Attribute{
					Name:        "delete_index_files_older_than_days",
					Description: `Enable delete index files older than days.`,
				},
				resource.Attribute{
					Name:        "delete_index_files_older_than_days_threshold",
					Description: `Delete index files older than days threshold.`,
				},
				resource.Attribute{
					Name:        "delete_when_free_disk_space_below",
					Description: `Enable delete when free disk space below.`,
				},
				resource.Attribute{
					Name:        "delete_when_free_disk_space_below_threshold",
					Description: `Delete when free disk space below threshold.`,
				},
				resource.Attribute{
					Name:        "detect_new_citrix_ica_application_names",
					Description: `Enable detect new citrix ica application names.`,
				},
				resource.Attribute{
					Name:        "enable_log_indexing",
					Description: `Enable log indexing.`,
				},
				resource.Attribute{
					Name:        "forward_logs_to_log_server",
					Description: `Enable forward logs to log server.`,
				},
				resource.Attribute{
					Name:        "forward_logs_to_log_server_name",
					Description: `Forward logs to log server name.`,
				},
				resource.Attribute{
					Name:        "forward_logs_to_log_server_schedule_name",
					Description: `Forward logs to log server schedule name.`,
				},
				resource.Attribute{
					Name:        "rotate_log_by_file_size",
					Description: `Enable rotate log by file size.`,
				},
				resource.Attribute{
					Name:        "rotate_log_file_size_threshold",
					Description: `Log file size threshold.`,
				},
				resource.Attribute{
					Name:        "rotate_log_on_schedule",
					Description: `Enable rotate log on schedule.`,
				},
				resource.Attribute{
					Name:        "rotate_log_schedule_name",
					Description: `Rotate log schedule name.`,
				},
				resource.Attribute{
					Name:        "stop_logging_when_free_disk_space_below",
					Description: `Enable stop logging when free disk space below.`,
				},
				resource.Attribute{
					Name:        "stop_logging_when_free_disk_space_below_threshold",
					Description: `Stop logging when free disk space below threshold.`,
				},
				resource.Attribute{
					Name:        "turn_on_qos_logging",
					Description: `Enable turn on qos logging.`,
				},
				resource.Attribute{
					Name:        "update_account_log_every",
					Description: `Update account log in every amount of seconds.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_threat_exception",
			Category:         "Data Sources",
			ShortDescription: `This resource allows you to execute Check Point Threat Exception.`,
			Description: `

This resource allows you to execute Check Point Threat Exception.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "layer",
					Description: `(Required) Layer that the rule belongs to identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the exception.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Object unique identifier.`,
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
					Description: `Action-the enforced profile.`,
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
					Name:        "source",
					Description: `Collection of Network objects identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "source_negate",
					Description: `True if negate is set for source.`,
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
					Name:        "protected_scope",
					Description: `Collection of objects defining Protected Scope identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "protected_scope_negate",
					Description: `True if negate is set for Protected Scope.`,
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
					Name:        "protection_or_site",
					Description: `Collection of protection or site objects identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "track",
					Description: `Packet tracking.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Owner UID.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_threat_rule",
			Category:         "Data Sources",
			ShortDescription: `This resource allows you to execute Check Point Threat Rule.`,
			Description: `

This resource allows you to execute Check Point Threat Rule.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "layer",
					Description: `(Required) Layer that the rule belongs to identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Rule name.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action-the enforced profile.`,
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
					Name:        "source",
					Description: `Collection of Network objects identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "source_negate",
					Description: `True if negate is set for source.`,
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
					Name:        "protected_scope",
					Description: `Collection of objects defining Protected Scope identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "protected_scope_negate",
					Description: `True if negate is set for Protected Scope.`,
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
					Name:        "track",
					Description: `Packet tracking.`,
				},
				resource.Attribute{
					Name:        "track_settings",
					Description: `Threat rule track settings. track_settings block are documented below.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string.`,
				},
				resource.Attribute{
					Name:        "exceptions",
					Description: `Collection of the rule's exceptions identified by UID. ` + "`" + `track_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "packet_capture",
					Description: `Packet capture.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_user",
			Category:         "Data Sources",
			ShortDescription: `This resource allows you to execute Check Point User.`,
			Description: `

This resource allows you to execute Check Point User.

`,
			Keywords: []string{},
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
					Name:        "email",
					Description: `User email.`,
				},
				resource.Attribute{
					Name:        "expiration_date",
					Description: `Expiration date in format: yyyy-MM-dd.`,
				},
				resource.Attribute{
					Name:        "phone_number",
					Description: `User phone number.`,
				},
				resource.Attribute{
					Name:        "authentication_method",
					Description: `Authentication method.`,
				},
				resource.Attribute{
					Name:        "radius_server",
					Description: `RADIUS server object identified by the name or UID. Must be set when "authentication-method" was selected to be "RADIUS".`,
				},
				resource.Attribute{
					Name:        "tacacs_server",
					Description: `TACACS server object identified by the name or UID. Must be set when "authentication-method" was selected to be "TACACS".`,
				},
				resource.Attribute{
					Name:        "connect_on_days",
					Description: `Days users allow to connect.`,
				},
				resource.Attribute{
					Name:        "connect_daily",
					Description: `Connect every day.`,
				},
				resource.Attribute{
					Name:        "from_hour",
					Description: `Allow users connect from hour.`,
				},
				resource.Attribute{
					Name:        "to_hour",
					Description: `Allow users connect until hour.`,
				},
				resource.Attribute{
					Name:        "allowed_locations",
					Description: `User allowed locations. allowed_locations blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "encryption",
					Description: `User encryption. encryption blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string. ` + "`" + `allowed_locations` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "destinations",
					Description: `Collection of allowed destination locations name or uid.`,
				},
				resource.Attribute{
					Name:        "sources",
					Description: `Collection of allowed source locations name or uid. ` + "`" + `encryption` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "enable_ike",
					Description: `Enable IKE encryption for users.`,
				},
				resource.Attribute{
					Name:        "enable_public_key",
					Description: `Enable IKE public key.`,
				},
				resource.Attribute{
					Name:        "enable_shared_secret",
					Description: `Enable IKE shared secret.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_user_group",
			Category:         "Data Sources",
			ShortDescription: `This resource allows you to execute Check Point User Group.`,
			Description: `

This resource allows you to execute Check Point User Group.

`,
			Keywords: []string{},
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
					Name:        "email",
					Description: `Email Address.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `Collection of User Group objects identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object.`,
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
			Type:             "checkpoint_management_user_template",
			Category:         "Data Sources",
			ShortDescription: `This resource allows you to execute Check Point User Template.`,
			Description: `

This resource allows you to execute Check Point User Template.

`,
			Keywords: []string{},
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
					Name:        "expiration_by_global_properties",
					Description: `Expiration date according to global properties.`,
				},
				resource.Attribute{
					Name:        "expiration_date",
					Description: `Expiration date in format: yyyy-MM-dd.`,
				},
				resource.Attribute{
					Name:        "authentication_method",
					Description: `Authentication method.`,
				},
				resource.Attribute{
					Name:        "radius_server",
					Description: `RADIUS server object identified by the name or UID. Must be set when "authentication-method" was selected to be "RADIUS".`,
				},
				resource.Attribute{
					Name:        "tacacs_server",
					Description: `TACACS server object identified by the name or UID. Must be set when "authentication-method" was selected to be "TACACS".`,
				},
				resource.Attribute{
					Name:        "connect_on_days",
					Description: `Days users allow to connect.`,
				},
				resource.Attribute{
					Name:        "connect_daily",
					Description: `Connect every day.`,
				},
				resource.Attribute{
					Name:        "from_hour",
					Description: `Allow users connect from hour.`,
				},
				resource.Attribute{
					Name:        "to_hour",
					Description: `Allow users connect until hour.`,
				},
				resource.Attribute{
					Name:        "allowed_locations",
					Description: `User allowed locations. allowed_locations blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "encryption",
					Description: `User encryption. encryption blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string. ` + "`" + `allowed_locations` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "destinations",
					Description: `Collection of allowed destination locations name or uid.`,
				},
				resource.Attribute{
					Name:        "sources",
					Description: `Collection of allowed source locations name or uid. ` + "`" + `encryption` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "enable_ike",
					Description: `Enable IKE encryption for users.`,
				},
				resource.Attribute{
					Name:        "enable_public_key",
					Description: `Enable IKE public key.`,
				},
				resource.Attribute{
					Name:        "enable_shared_secret",
					Description: `Enable IKE shared secret.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_vpn_community_remote_access",
			Category:         "Data Sources",
			ShortDescription: `This resource allows you to execute Check Point VPN Community Remote Access.`,
			Description: `

This resource allows you to execute Check Point VPN Community Remote Access.

`,
			Keywords: []string{},
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
					Name:        "gateways",
					Description: `Collection of Gateway objects identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "user_groups",
					Description: `Collection of User group objects identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag identifiers.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"checkpoint_management_access_point_name":                         0,
		"checkpoint_management_checkpoint_host":                           1,
		"checkpoint_management_data_access_layer":                         2,
		"checkpoint_management_data_access_role":                          3,
		"checkpoint_management_data_access_rule":                          4,
		"checkpoint_management_data_access_section":                       5,
		"checkpoint_management_data_address_range":                        6,
		"checkpoint_management_data_application_site":                     7,
		"checkpoint_management_data_application_site_category":            8,
		"checkpoint_management_data_application_site_group":               9,
		"checkpoint_management_data_dns_domain":                           10,
		"checkpoint_management_data_dynamic_object":                       11,
		"checkpoint_management_data_exception_group":                      12,
		"checkpoint_management_data_group":                                13,
		"checkpoint_management_data_group_with_exclusion":                 14,
		"checkpoint_management_data_host":                                 15,
		"checkpoint_management_data_https_layer":                          16,
		"checkpoint_management_data_https_rule":                           17,
		"checkpoint_management_data_https_section":                        18,
		"checkpoint_management_data_multicast_address_range":              19,
		"checkpoint_management_data_network":                              20,
		"checkpoint_management_data_opsec_application":                    21,
		"checkpoint_management_data_package":                              22,
		"checkpoint_management_data_security_zone":                        23,
		"checkpoint_management_data_service_dce_rpc":                      24,
		"checkpoint_management_data_service_group":                        25,
		"checkpoint_management_data_service_icmp":                         26,
		"checkpoint_management_data_service_icmp6":                        27,
		"checkpoint_management_data_service_other":                        28,
		"checkpoint_management_data_service_rpc":                          29,
		"checkpoint_management_data_service_sctp":                         30,
		"checkpoint_management_data_service_tcp":                          31,
		"checkpoint_management_data_service_udp":                          32,
		"checkpoint_management_data_threat_indicator":                     33,
		"checkpoint_management_data_time_group":                           34,
		"checkpoint_management_data_vpn_community_meshed":                 35,
		"checkpoint_management_data_vpn_community_star":                   36,
		"checkpoint_management_data_wildcard":                             37,
		"checkpoint_management_gsn_handover_group":                        38,
		"checkpoint_management_identity_tag":                              39,
		"checkpoint_management_mds":                                       40,
		"checkpoint_management_nat_rule":                                  41,
		"checkpoint_management_nat_section":                               42,
		"checkpoint_management_service_citrix_tcp":                        43,
		"checkpoint_management_service_compound_tcp":                      44,
		"checkpoint_management_show_objects":                              45,
		"checkpoint_management_show_updatable_objects_repository_content": 46,
		"checkpoint_management_simple_cluster":                            47,
		"checkpoint_management_simple_gateway":                            48,
		"checkpoint_management_threat_exception":                          49,
		"checkpoint_management_threat_rule":                               50,
		"checkpoint_management_user":                                      51,
		"checkpoint_management_user_group":                                52,
		"checkpoint_management_user_template":                             53,
		"checkpoint_management_vpn_community_remote_access":               54,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
