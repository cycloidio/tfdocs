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
			Type:             "checkpoint_management_access_rulebase",
			Category:         "Data Sources",
			ShortDescription: `This resource allows you to execute Check Point Access Rule Base.`,
			Description: `

Use this data source to get information on an existing access RuleBase.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name. Must be unique in the domain.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `APN name.`,
				},
				resource.Attribute{
					Name:        "filter_settings",
					Description: `Enable enforce end user domain. filter_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "limit",
					Description: `The maximal number of returned results.`,
				},
				resource.Attribute{
					Name:        "offset",
					Description: `Number of the results to initially skip.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `Sorts the results by search criteria. Automatically sorts the results by Name, in the ascending order. orders blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "package",
					Description: `Name of the package.`,
				},
				resource.Attribute{
					Name:        "show_as_ranges",
					Description: `When true, the source, destination and services & applications parameters are displayed as ranges of IP addresses and port numbers rather than network objects. Objects that are not represented using IP addresses or port numbers are presented as objects. In addition, the response of each rule does not contain the parameters: source, source-negate, destination, destination-negate, service and service-negate, but instead it contains the parameters:source-ranges, destination-ranges and service-ranges. Note: Requesting to show rules as ranges is limited up to 20 rules per request, otherwise an error is returned. If you wish to request more rules, use the offset and limit parameters to limit your request.`,
				},
				resource.Attribute{
					Name:        "show_hits",
					Description: `show hits.`,
				},
				resource.Attribute{
					Name:        "hits_settings",
					Description: `hits_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "dereference_group_members",
					Description: `Indicates whether to dereference "members" field by details level for every object in reply.`,
				},
				resource.Attribute{
					Name:        "show_membership",
					Description: `Indicates whether to calculate and show "groups" field for every object in reply. ` + "`" + `filter_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "search_mode",
					Description: `When set to 'general', both the Full Text Search and Packet Search are enabled. In this mode, Packet Search will not match on 'Any' object, a negated cell or a group-with-exclusion. When the search-mode is set to 'packet', by default, the match on 'Any' object, a negated cell or a group-with-exclusion are enabled. packet-search-settings may be provided to change the default behavior.`,
				},
				resource.Attribute{
					Name:        "expand_group_members",
					Description: `(Optional, can only be used when search_mode is set to "packet") When true, if the search expression contains a UID or a name of a group object, results will include rules that match on at least one member of the group.`,
				},
				resource.Attribute{
					Name:        "expand_group_with_exclusion_members",
					Description: `(Optional, can only be used when search_mode is set to "packet") When true, if the search expression contains a UID or a name of a group-with-exclusion object, results will include rules that match at least one member of the "include" part and is not a member of the "except" part.`,
				},
				resource.Attribute{
					Name:        "match_on_any",
					Description: `(Optional, can only be used when search_mode is set to "packet") Whether to match on 'Any' object.`,
				},
				resource.Attribute{
					Name:        "match_on_group_with_exclusion",
					Description: `(Optional, can only be used when search_mode is set to "packet") Whether to match on a group-with-exclusion.`,
				},
				resource.Attribute{
					Name:        "match_on_negate",
					Description: `(Optional, can only be used when search_mode is set to "packet") Whether to match on a negated cell. ` + "`" + `order` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "asc",
					Description: `(Optional) Sorts results by the given field in ascending order.`,
				},
				resource.Attribute{
					Name:        "desc",
					Description: `(Optional) Sorts results by the given field in descending order. ` + "`" + `hits_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "from-date",
					Description: `Format: YYYY-MM-DD, YYYY-mm-ddThh:mm:ss.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `Target gateway name or UID.`,
				},
				resource.Attribute{
					Name:        "to-date",
					Description: `Format: YYYY-MM-DD, YYYY-mm-ddThh:mm:ss.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_aci_data_center_server",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Cisco APIC data center server.`,
			Description: `

Use this data source to get information on an existing Cisco APIC Data Center Server.

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
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_administrator",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Administrator.`,
			Description: `

Use this data source to get information on an existing Check Point Administrator.

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
					Name:        "authentication_method",
					Description: `Authentication method.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `Administrator email.`,
				},
				resource.Attribute{
					Name:        "multi_domain_profile",
					Description: `Administrator multi-domain profile. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.`,
				},
				resource.Attribute{
					Name:        "must_change_password",
					Description: `True if administrator must change password on the next login.`,
				},
				resource.Attribute{
					Name:        "permissions_profile",
					Description: `Administrator permissions profile. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level. permissions_profile blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "phone_number",
					Description: `Administrator phone number.`,
				},
				resource.Attribute{
					Name:        "radius_server",
					Description: `RADIUS server object identified by the name or UID. Must be set when "authentication-method" was selected to be "RADIUS". Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.`,
				},
				resource.Attribute{
					Name:        "sic_name",
					Description: `Name of the Secure Internal Connection Trust.`,
				},
				resource.Attribute{
					Name:        "tacacs_server",
					Description: `TACACS server object identified by the name or UID . Must be set when "authentication-method" was selected to be "TACACS". Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag objects identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string. ` + "`" + `permissions_profile` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `The domain's profile.`,
				},
				resource.Attribute{
					Name:        "profile",
					Description: `Permission profile.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_api_settings",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Api Settings.`,
			Description: `

Use this data source to get information on an existing Check Point Api Settings.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Object name.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "accepted_api_calls_from",
					Description: `Clients allowed to connect to the API Server.`,
				},
				resource.Attribute{
					Name:        "automatic_start",
					Description: `MGMT API will start after server will start.`,
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
			Type:             "checkpoint_management_automatic_purge",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Automatic Purge.`,
			Description: `

Use this data source to get information on an existing Check Point Automatic Purge.


`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "enabled",
					Description: `Turn on/off the automatic-purge feature.`,
				},
				resource.Attribute{
					Name:        "keep_sessions_by_count",
					Description: `Whether or not to keep the latest N sessions.`,
				},
				resource.Attribute{
					Name:        "number_of_sessions_to_keep",
					Description: `The number of newest sessions to preserve, by the sessions's publish date.`,
				},
				resource.Attribute{
					Name:        "keep_sessions_by_days",
					Description: `Whether or not to keep the sessions for D days.`,
				},
				resource.Attribute{
					Name:        "number_of_days_to_keep",
					Description: `When "keep-sessions-by-days = true" this sets the number of days to keep the sessions.`,
				},
				resource.Attribute{
					Name:        "scheduling",
					Description: `When to purge sessions that do not meet the "keep" criteria. scheduling is type List. scheduling blocks are documented below. ` + "`" + `scheduling` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "start_date",
					Description: `The first time to check whether or not there are sessions to purge.`,
				},
				resource.Attribute{
					Name:        "time_units",
					Description: `The time units.`,
				},
				resource.Attribute{
					Name:        "check_interval",
					Description: `Number of time-units between two purge checks.`,
				},
				resource.Attribute{
					Name:        "last_check",
					Description: `Last time purge check was executed.`,
				},
				resource.Attribute{
					Name:        "next_check",
					Description: `Next time purge check will be executed.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_aws_data_center_server",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing AWS data center server.`,
			Description: `

Use this data source to get information on an existing AWS Data Center Server.

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
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_azure_ad",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Azure Ad.`,
			Description: `

Use this data source to get information on an existing Check Point Azure Ad.

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
					Name:        "properties",
					Description: `Azure AD connection properties. properties blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag objects identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.`,
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
			Type:             "checkpoint_management_azure_ad_content",
			Category:         "Data Sources",
			ShortDescription: `This resource allows you to execute Check Point Azure Ad Content.`,
			Description: `

This resource allows you to execute Check Point Azure Ad Content.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "azure_ad_name",
					Description: `(Optional) Name of the Azure AD Server where to search for objects.`,
				},
				resource.Attribute{
					Name:        "azure_ad_uid",
					Description: `(Optional) Unique identifier of the Azure AD Server where to search for objects.`,
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
					Name:        "uid_in_azure_ad",
					Description: `(Optional) Return result matching the unique identifier of the object on the Azure AD Server.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Return results matching the specified filter. filter blocks are documented below. ` + "`" + `order` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "asc",
					Description: `(Optional) Sorts results by the given field in ascending order.`,
				},
				resource.Attribute{
					Name:        "desc",
					Description: `(Optional) Sorts results by the given field in descending order. ` + "`" + `filter` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "text",
					Description: `(Optional) Return results containing the specified text value.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `(Optional) Return results under the specified Data Center Object (identified by URI).`,
				},
				resource.Attribute{
					Name:        "parent_uid_in_data_center",
					Description: `(Optional) Return results under the specified Data Center Object (identified by UID). Output:`,
				},
				resource.Attribute{
					Name:        "from",
					Description: `From which element number the query was done.`,
				},
				resource.Attribute{
					Name:        "objects",
					Description: `Remote objects views. objects blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "to",
					Description: `To which element number the query was done.`,
				},
				resource.Attribute{
					Name:        "total",
					Description: `Total number of elements returned by the query. ` + "`" + `objects` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name_in_azure_ad",
					Description: `Object name in the Azure AD.`,
				},
				resource.Attribute{
					Name:        "uid_in_azure_ad",
					Description: `Unique identifier of the object in the Azure AD.`,
				},
				resource.Attribute{
					Name:        "azure_ad_object",
					Description: `The imported management object (if exists). Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Object management name.`,
				},
				resource.Attribute{
					Name:        "type_in_azure_ad",
					Description: `Object type in Azure AD.`,
				},
				resource.Attribute{
					Name:        "additional_properties",
					Description: `Additional properties on the object. additional_properties blocks are documented below. ` + "`" + `additional_properties` + "`" + ` supports the following:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_azure_data_center_server",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing azure data center server.`,
			Description: `

Use this data source to get information on an existing Microsoft Azure Data Center Server.

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
			Type:             "checkpoint_management_cloud_services",
			Category:         "Data Sources",
			ShortDescription: `This resource allows you to execute Check Point Show Cloud Services.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `Status of the connection to the Infinity Portal.`,
				},
				resource.Attribute{
					Name:        "connected_at",
					Description: `The time of the connection between the Management Server and the Infinity Portal. ` + "`" + `connected_at` + "`" + ` is documented below.`,
				},
				resource.Attribute{
					Name:        "management_url",
					Description: `The Management Server's public URL.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `Tenant ID of Infinity Portal.`,
				},
				resource.Attribute{
					Name:        "gateways_onboarding_settings",
					Description: `Gateways on-boarding to Infinity Portal settings. ` + "`" + `gateways_onboarding_settings` + "`" + ` is documented below. ` + "`" + `connected_at` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "iso_8601",
					Description: `Date and time represented in international ISO 8601 format.`,
				},
				resource.Attribute{
					Name:        "posix",
					Description: `Number of milliseconds that have elapsed since 00:00:00, 1 January 1970. ` + "`" + `gateways_onboarding_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "connection_method",
					Description: `Indicate whether Gateways will be connected to Infinity Portal automatically or only after policy installation.`,
				},
				resource.Attribute{
					Name:        "participant_gateways",
					Description: `Which Gateways will be connected to Infinity Portal.`,
				},
				resource.Attribute{
					Name:        "specific_gateways",
					Description: `Collection of targets identified by Name or UID. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_cluster_member",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Cluster Member.`,
			Description: `

Use this data source to get information on an existing Check Point Cluster Member.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "uid",
					Description: `(Required) Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "limit_interfaces",
					Description: `(Optional) Limit number of cluster member interfaces to show.`,
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
			Type:             "checkpoint_management_data_center_content",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point data center content.`,
			Description: `

Use this data source to get information on an existing Check Point data center content.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "data_center_name",
					Description: `(Optional) Name of the Data Center Server where to search for objects.`,
				},
				resource.Attribute{
					Name:        "data_center_uid",
					Description: `(Optional) Unique identifier of the Data Center Server where to search for objects.`,
				},
				resource.Attribute{
					Name:        "limit",
					Description: `The maximal number of returned results.`,
				},
				resource.Attribute{
					Name:        "offset",
					Description: `Number of the results to initially skip.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `Sorts the results by search criteria. Automatically sorts the results by Name, in the ascending order. orders blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "uid_in_data_center",
					Description: `Return result matching the unique identifier of the object on the Data Center Server.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `Return results matching the specified filter. ` + "`" + `filter` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "text",
					Description: `(Optional) Return results containing the specified text value.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `(Optional) Return results under the specified Data Center Object (identified by URI).`,
				},
				resource.Attribute{
					Name:        "parent_uid_in_data_center",
					Description: `(Optional) Return results under the specified Data Center Object (identified by UID). ` + "`" + `order` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "asc",
					Description: `(Optional) Sorts results by the given field in ascending order.`,
				},
				resource.Attribute{
					Name:        "desc",
					Description: `(Optional) Sorts results by the given field in descending order.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_data_center_query",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Data Center Query.`,
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
					Name:        "tunnel_granularity",
					Description: `VPN tunnel sharing option to be used.`,
				},
				resource.Attribute{
					Name:        "granular_encryptions",
					Description: `VPN granular encryption settings. granular_encryptions blocks are documented below.`,
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
					Description: `Comments string. ` + "`" + `override_vpn_domains` + "`" + ` supports the following:`,
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
					Description: `Shared secret. granular_encryptions` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "internal-gateway",
					Description: `Internally managed Check Point gateway identified by name or UID, or 'Any' for all internal-gateways participants in this community.`,
				},
				resource.Attribute{
					Name:        "external_gateway",
					Description: `Externally managed or 3rd party gateway identified by name or UID.`,
				},
				resource.Attribute{
					Name:        "encryption_method",
					Description: `The encryption method to be used: prefer ikev2 but support ikev1, ikev2 only, ikev1 for ipv4 and ikev2 for ipv6 only.`,
				},
				resource.Attribute{
					Name:        "encryption_suite",
					Description: `The encryption suite to be used: suite-b-gcm-256, custom, vpn b, vpn a, suite-b-gcm-128.`,
				},
				resource.Attribute{
					Name:        "ike_phase_1",
					Description: `Ike Phase 1 settings. Only applicable when the encryption-suite is set to [custom]. ike_phase_1 blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "ike_phase_2",
					Description: `Ike Phase 2 settings. Only applicable when the encryption-suite is set to [custom]. ike_phase_2 blocks are documented below. ` + "`" + `ike_phase_1` + "`" + ` supports the following:`,
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
					Description: `The encryption algorithm to be used.`,
				},
				resource.Attribute{
					Name:        "ike_p1_rekey_time",
					Description: `Indicates the time interval for IKE phase 1 renegotiation. ` + "`" + `ike_phase_2` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "data_integrity",
					Description: `The hash algorithm to be used.`,
				},
				resource.Attribute{
					Name:        "encryption_algorithm",
					Description: `The encryption algorithm to be used.`,
				},
				resource.Attribute{
					Name:        "ike_p2_use_pfs",
					Description: `Indicates whether Perfect Forward Secrecy (PFS) is being used for IKE phase 2.`,
				},
				resource.Attribute{
					Name:        "ike_p2_pfs_dh_grp",
					Description: `The Diffie-Hellman group to be used: group-1, group-2, group-5, group-14, group-15, group-16, group-17, group-18, group-19, group-20, group-24.`,
				},
				resource.Attribute{
					Name:        "ike_p2_rekey_time",
					Description: `Indicates the time interval for IKE phase 2 renegotiation.`,
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
					Name:        "tunnel_granularity",
					Description: `VPN tunnel sharing option to be used.`,
				},
				resource.Attribute{
					Name:        "granular_encryptions",
					Description: `VPN granular encryption settings. granular_encryptions blocks are documented below.`,
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
					Description: `Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. ` + "`" + `override_vpn_domains` + "`" + ` supports the following:`,
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
					Description: `Shared secret. granular_encryptions` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "internal-gateway",
					Description: `Internally managed Check Point gateway identified by name or UID, or 'Any' for all internal-gateways participants in this community.`,
				},
				resource.Attribute{
					Name:        "external_gateway",
					Description: `Externally managed or 3rd party gateway identified by name or UID.`,
				},
				resource.Attribute{
					Name:        "encryption_method",
					Description: `The encryption method to be used: prefer ikev2 but support ikev1, ikev2 only, ikev1 for ipv4 and ikev2 for ipv6 only.`,
				},
				resource.Attribute{
					Name:        "encryption_suite",
					Description: `The encryption suite to be used: suite-b-gcm-256, custom, vpn b, vpn a, suite-b-gcm-128.`,
				},
				resource.Attribute{
					Name:        "ike_phase_1",
					Description: `Ike Phase 1 settings. Only applicable when the encryption-suite is set to [custom]. ike_phase_1 blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "ike_phase_2",
					Description: `Ike Phase 2 settings. Only applicable when the encryption-suite is set to [custom]. ike_phase_2 blocks are documented below. ` + "`" + `ike_phase_1` + "`" + ` supports the following:`,
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
					Description: `The encryption algorithm to be used.`,
				},
				resource.Attribute{
					Name:        "ike_p1_rekey_time",
					Description: `Indicates the time interval for IKE phase 1 renegotiation. ` + "`" + `ike_phase_2` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "data_integrity",
					Description: `The hash algorithm to be used.`,
				},
				resource.Attribute{
					Name:        "encryption_algorithm",
					Description: `The encryption algorithm to be used.`,
				},
				resource.Attribute{
					Name:        "ike_p2_use_pfs",
					Description: `Indicates whether Perfect Forward Secrecy (PFS) is being used for IKE phase 2.`,
				},
				resource.Attribute{
					Name:        "ike_p2_pfs_dh_grp",
					Description: `The Diffie-Hellman group to be used: group-1, group-2, group-5, group-14, group-15, group-16, group-17, group-18, group-19, group-20, group-24.`,
				},
				resource.Attribute{
					Name:        "ike_p2_rekey_time",
					Description: `Indicates the time interval for IKE phase 2 renegotiation.`,
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
			Type:             "checkpoint_management_domain",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Domain.`,
			Description: `

Use this data source to get information on an existing Check Point Domain.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name. Must be unique in the domain.`,
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
			Type:             "checkpoint_management_domain_permissions_profile",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Domain Permissions Profile.`,
			Description: `

Use this data source to get information on an existing Check Point Domain Permissions Profile.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
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
			Type:             "checkpoint_management_dynamic_global_network_object",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Dynamic Global Network Object.`,
			Description: `

Use this data source to get information on an existing Check Point Dynamic Global Network Object.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag objects identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.`,
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
			Type:             "checkpoint_management_gaia_best_practice",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Host.`,
			Description: `

Use this data source to get information on an existing Check Point Host.


`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "best_practice_id",
					Description: `(Optional) Best Practice ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Best Practice Name.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Best Practice UID.`,
				},
				resource.Attribute{
					Name:        "action_item",
					Description: `Action item to comply with the Best Practice.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the Best Practice.`,
				},
				resource.Attribute{
					Name:        "expected_output_base64",
					Description: `The expected output of the script in Base64. Available only for user-defined best practices.`,
				},
				resource.Attribute{
					Name:        "practice_script_base64",
					Description: `The script to run on Gaia Security Gateways during the Compliance scans in Base64. Available only for user-defined best practices.`,
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
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Google Cloud Platform Data Center Server.`,
			Description: `

Use this data source to get information on an existing Google Cloud Platform Data Center Server.

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
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_generic_data_center_server",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing generic data center server.`,
			Description: `

Use this data source to get information on an existing Generic Data Center Server.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required if uid is not given) Object name.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Required if name is not given) Object unique identifier.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_global_assignment",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Global Assignment.`,
			Description: `

Use this data source to get information on an existing Check Point Global Assignment.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "dependent_domain",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "global_domain",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "assignment_up_to_date",
					Description: `The time when the assignment was assigned. assignment_up_to_date blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "global_access_policy",
					Description: `Global domain access policy that is assigned to a dependent domain.`,
				},
				resource.Attribute{
					Name:        "global_threat_prevention_policy",
					Description: `Global domain threat prevention policy that is assigned to a dependent domain.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag objects identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.`,
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
			Type:             "checkpoint_management_global_domain",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Global Domain.`,
			Description: `

Use this data source to get information on an existing Check Point Global Domain.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optioanl) Object name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Object type.`,
				},
				resource.Attribute{
					Name:        "domain_type",
					Description: `The domain type.`,
				},
				resource.Attribute{
					Name:        "global_domain_assignments",
					Description: `The assignments. global_domain_assignments blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "servers",
					Description: `Domain Servers. servers blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag objects identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string. ` + "`" + `global_domain_assignments` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Object name.`,
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
					Name:        "assignment_status",
					Description: `The status of the assignment.`,
				},
				resource.Attribute{
					Name:        "assignment_up_to_date",
					Description: `The time when the assignment was assigned. assignment_up_to_date blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "dependent_domain",
					Description: `Dependent domain. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level. dependent_domain blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "global_access_policy",
					Description: `Global domain access policy that is assigned to a dependent domain.`,
				},
				resource.Attribute{
					Name:        "global_threat_prevention_policy",
					Description: `Global domain threat prevention policy that is assigned to a dependent domain.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag objects identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Object color.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Coemmnet string. ` + "`" + `assignment_up_to_date` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "iso_9601",
					Description: `Date and time represented in international ISO 8601 format.`,
				},
				resource.Attribute{
					Name:        "posix",
					Description: `Number of milliseconds that have elapsed since 00:00:00, 1 January 1970. ` + "`" + `dependent_domain` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Object name.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `Object unique identifier. ` + "`" + `servers` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Object name.`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `Domain server status.`,
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
					Name:        "multi_domain_server",
					Description: `Multi Domain server name or UID.`,
				},
				resource.Attribute{
					Name:        "skip_start_domain_server",
					Description: `Set this value to be true to prevent starting the new created domain.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Domain server type.`,
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
			Type:             "checkpoint_management_https_rulebase",
			Category:         "Data Sources",
			ShortDescription: `This resource allows you to execute Check Point Https Rule Base.`,
			Description: `

Use this data source to get information on an existing https RuleBase.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Object name. Must be unique in the domain.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `Search expression to filter the rulebase. The provided text should be exactly the same as it would be given in Smart Console. The logical operators in the expression ('AND', 'OR') should be provided in capital letters. If an operator is not used, the default OR operator applies.`,
				},
				resource.Attribute{
					Name:        "filter_settings",
					Description: `Enable enforce end user domain. filter_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "limit",
					Description: `The maximal number of returned results.`,
				},
				resource.Attribute{
					Name:        "offset",
					Description: `Number of the results to initially skip.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `Sorts the results by search criteria. Automatically sorts the results by Name, in the ascending order. orders blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "package",
					Description: `Name of the package.`,
				},
				resource.Attribute{
					Name:        "dereference_group_members",
					Description: `Indicates whether to dereference "members" field by details level for every object in reply.`,
				},
				resource.Attribute{
					Name:        "show_membership",
					Description: `Indicates whether to calculate and show "groups" field for every object in reply. ` + "`" + `filter_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "search_mode",
					Description: `When set to 'general', both the Full Text Search and Packet Search are enabled. In this mode, Packet Search will not match on 'Any' object, a negated cell or a group-with-exclusion. When the search-mode is set to 'packet', by default, the match on 'Any' object, a negated cell or a group-with-exclusion are enabled. packet-search-settings may be provided to change the default behavior.`,
				},
				resource.Attribute{
					Name:        "expand_group_members",
					Description: `(Optional, can only be used when search_mode is set to "packet") When true, if the search expression contains a UID or a name of a group object, results will include rules that match on at least one member of the group.`,
				},
				resource.Attribute{
					Name:        "expand_group_with_exclusion_members",
					Description: `(Optional, can only be used when search_mode is set to "packet") When true, if the search expression contains a UID or a name of a group-with-exclusion object, results will include rules that match at least one member of the "include" part and is not a member of the "except" part.`,
				},
				resource.Attribute{
					Name:        "match_on_any",
					Description: `(Optional, can only be used when search_mode is set to "packet") Whether to match on 'Any' object.`,
				},
				resource.Attribute{
					Name:        "match_on_group_with_exclusion",
					Description: `(Optional, can only be used when search_mode is set to "packet") Whether to match on a group-with-exclusion.`,
				},
				resource.Attribute{
					Name:        "match_on_negate",
					Description: `(Optional, can only be used when search_mode is set to "packet") Whether to match on a negated cell. ` + "`" + `order` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "asc",
					Description: `(Optional) Sorts results by the given field in ascending order.`,
				},
				resource.Attribute{
					Name:        "desc",
					Description: `(Optional) Sorts results by the given field in descending order.`,
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
			Type:             "checkpoint_management_idp_administrator_group",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Idp Administrator Group.`,
			Description: `

Use this data source to get information on an existing Check Point Idp Administrator Group.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
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
			Type:             "checkpoint_management_idp_default_assignment",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Idp Default Assignment.`,
			Description: `

Use this data source to get information on an existing Check Point Idp Default Assignment.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_idp_to_domain_assignment",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Idp To Domain Assignment.`,
			Description: `

Use this data source to get information on an existing Check Point Idp To Domain Assignment.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "assigned_domain",
					Description: `(Optional) Represents the Domain assigned by 'idp-to-domain-assignment', need to be domain name or UID. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_interoperable_device",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Interoperable Device.`,
			Description: `

Use this data source to get information on an existing Check Point Interoperable Device.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
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
			Type:             "checkpoint_management_ips_update_schedule",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Ips Update Schedule.`,
			Description: `

Use this data source to get information on an existing Check Point Ips Update Schedule.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "enabled",
					Description: `IPS Update Schedule status.`,
				},
				resource.Attribute{
					Name:        "time",
					Description: `Time in format HH:mm.`,
				},
				resource.Attribute{
					Name:        "recurrence",
					Description: `Days recurrence. recurrence blocks are documented below. ` + "`" + `recurrence` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "days",
					Description: `Valid on specific days. Multiple options, support range of days in months. Example:["1","3","9-20"].`,
				},
				resource.Attribute{
					Name:        "minutes",
					Description: `Valid on interval. The length of time in minutes between updates.`,
				},
				resource.Attribute{
					Name:        "pattern",
					Description: `Valid on "Interval", "Daily", "Weekly", "Monthly" base.`,
				},
				resource.Attribute{
					Name:        "weekdays",
					Description: `Valid on weekdays. Example: "Sun", "Mon"..."Sat".`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_ise_data_center_server",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Cisco ISE data center server.`,
			Description: `

Use this data source to get information on an existing Cisco ISE Data Center Server.

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
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_kubernetes_data_center_server",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Kubernetes Data Center Server.`,
			Description: `

Use this data source to get information on an existing Kubernetes Data Center Server.

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
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_login_message",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Login Message.`,
			Description: `

Use this data source to get information on an existing Check Point Login Message.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "header",
					Description: `Login message header.`,
				},
				resource.Attribute{
					Name:        "message",
					Description: `Login message body.`,
				},
				resource.Attribute{
					Name:        "show_message",
					Description: `Whether to show login message.`,
				},
				resource.Attribute{
					Name:        "warning",
					Description: `Add warning sign.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_lsm_cluster_profile",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Lsm Cluster Profile.`,
			Description: `

Use this data source to get information on an existing Check Point Lsm Cluster Profile.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_lsv_profile",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Lsv Profile.`,
			Description: `

Use this data source to get information on an existing Check Point Lsv Profile.

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
					Name:        "allowed_ip_addresses",
					Description: `Collection of network objects identified by name or UID that represent IP addresses allowed in profile's VPN domain.`,
				},
				resource.Attribute{
					Name:        "certificate_authority",
					Description: `Trusted Certificate authority for establishing trust between VPN peers, identified by name or UID.`,
				},
				resource.Attribute{
					Name:        "restrict_allowed_addresses",
					Description: `Indicate whether the IP addresses allowed in the VPN Domain will be restricted or not, according to allowed-ip-addresses field.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag objects identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.`,
				},
				resource.Attribute{
					Name:        "vpn_domain",
					Description: `peers' VPN Domain properties. vpn_domain blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string. ` + "`" + `vpn_domain` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "limit_peer_domain_size",
					Description: `Use this parameter to limit the number of IP addresses in the VPN Domain of each peer according to the value in the max-allowed-addresses field.`,
				},
				resource.Attribute{
					Name:        "max_allowed_addresses",
					Description: `Maximum number of IP addresses in the VPN Domain of each peer. This value will be enforced only when limit-peer-domain-size field is set to true. Select a value between 1 and 256. Default value is 256.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_md_permissions_profile",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Md Permissions Profile.`,
			Description: `

Use this data source to get information on an existing Check Point Md Permissions Profile.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
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
			Type:             "checkpoint_management_nat_rulebase",
			Category:         "Data Sources",
			ShortDescription: `This resource allows you to execute Check Point Nat Rule Base.`,
			Description: `

Use this data source to get information on an existing nat RuleBase.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "package",
					Description: `(Required) Name of the package.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `APN name.`,
				},
				resource.Attribute{
					Name:        "filter_settings",
					Description: `Enable enforce end user domain. filter_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "limit",
					Description: `The maximal number of returned results.`,
				},
				resource.Attribute{
					Name:        "offset",
					Description: `Number of the results to initially skip.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `Sorts the results by search criteria. Automatically sorts the results by Name, in the ascending order. orders blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "dereference_group_members",
					Description: `Indicates whether to dereference "members" field by details level for every object in reply.`,
				},
				resource.Attribute{
					Name:        "show_membership",
					Description: `Indicates whether to calculate and show "groups" field for every object in reply. ` + "`" + `filter_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "search_mode",
					Description: `When set to 'general', both the Full Text Search and Packet Search are enabled. In this mode, Packet Search will not match on 'Any' object, a negated cell or a group-with-exclusion. When the search-mode is set to ' packet', by default, the match on 'Any' object, a negated cell or a group-with-exclusion are enabled. packet-search-settings may be provided to change the default behavior.`,
				},
				resource.Attribute{
					Name:        "expand_group_members",
					Description: `(Optional, can only be used when search_mode is set to "packet") When true, if the search expression contains a UID or a name of a group object, results will include rules that match on at least one member of the group.`,
				},
				resource.Attribute{
					Name:        "expand_group_with_exclusion_members",
					Description: `(Optional, can only be used when search_mode is set to "packet") When true, if the search expression contains a UID or a name of a group-with-exclusion object, results will include rules that match at least one member of the "include" part and is not a member of the "except" part.`,
				},
				resource.Attribute{
					Name:        "match_on_any",
					Description: `(Optional, can only be used when search_mode is set to "packet") Whether to match on 'Any' object.`,
				},
				resource.Attribute{
					Name:        "match_on_group_with_exclusion",
					Description: `(Optional, can only be used when search_mode is set to "packet") Whether to match on a group-with-exclusion.`,
				},
				resource.Attribute{
					Name:        "match_on_negate",
					Description: `(Optional, can only be used when search_mode is set to "packet") Whether to match on a negated cell. ` + "`" + `order` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "asc",
					Description: `(Optional) Sorts results by the given field in ascending order.`,
				},
				resource.Attribute{
					Name:        "desc",
					Description: `(Optional) Sorts results by the given field in descending order.`,
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
			Type:             "checkpoint_management_network_feed",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Network Feed.`,
			Description: `

Use this data source to get information on an existing Check Point Network Feed.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
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
			Type:             "checkpoint_management_nuage_data_center_server",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Nuage Data Center Server.`,
			Description: `

Use this data source to get information on an existing Nuage Data Center Server.

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
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_nutanix_data_center_server",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Nutanix data center server.`,
			Description: `

Use this data source to get information on an existing Nutanix Data Center Server.

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
					Name:        "automatic_refresh",
					Description: `Indicates whether the data center server's content is automatically updated.`,
				},
				resource.Attribute{
					Name:        "data_center_type",
					Description: `Data center type.`,
				},
				resource.Attribute{
					Name:        "properties",
					Description: `Data center properties. properties blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag objects identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level. ` + "`" + `properties` + "`" + ` supports the following:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_objects",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Objects.`,
			Description: `

Use this data source to get information on an existing Check Point Objects.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "uids",
					Description: `(Optional) List of UIDs of the objects to retrieve.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Search expression to filter objects by. The provided text should be exactly the same as it would be given in Smart Console. The logical operators in the expression ('AND', 'OR') should be provided in capital letters. By default, the search involves both a textual search and a IP search. To use IP search only, set the "ip-only" parameter to true.`,
				},
				resource.Attribute{
					Name:        "ip_only",
					Description: `(Optional) If using "filter", use this field to search objects by their IP address only, without involving the textual search.`,
				},
				resource.Attribute{
					Name:        "limit",
					Description: `(Optional) The maximal number of returned results.`,
				},
				resource.Attribute{
					Name:        "offset",
					Description: `(Optional) The maximal number of returned results.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Optional) Sorts the results by search criteria. Automatically sorts the results by Name, in the ascending order. order blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The objects' type, e.g.: host, service-tcp, network, address-range...`,
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
					Name:        "domains_to_process",
					Description: `(Optional) Indicates which domains to process the commands on. It cannot be used with the details-level full, must be run from the System Domain only and with ignore-warnings true. Valid values are: CURRENT_DOMAIN, ALL_DOMAINS_ON_THIS_SERVER. ` + "`" + `order` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "asc",
					Description: `(Optional) Sorts results by the given field in ascending order.`,
				},
				resource.Attribute{
					Name:        "desc",
					Description: `(Optional) Sorts results by the given field in descending order.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_openstack_data_center_server",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing OpenStack Data Center Server.`,
			Description: `

Use this data source to get information on an existing OpenStack Data Center Server.

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
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_oracle_cloud_data_center_server",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Oracle Cloud data center server.`,
			Description: `

Use this data source to get information on an existing Oracle Cloud Data Center Server.

` + "`" + `` + "`" + `` + "`" + `hcl
resource "checkpoint_management_oracle_cloud_data_center_server" "testOracleCloud" {
  name = "MY-ORACLE-CLOUD"
  authentication_method = "key-authentication"
  private_key = "0SLtLS1CRUdJTiBQUklWQVRFIEtS0FWtLS0tDQpNSUlFdkFJQkFEQU5CZ2txaGtpRzl3AAAAUUVGQUFTQ0JLWXdnZ1NpQWdFQUFvSUJBUURUdmVrK1laMmNSekVmDQp1QkNoMkFxS2hzUFcrQUhUajY4dE5VbVl4OUFTRXBsREhnMkF0bCtMRWRRWUFRSUtLMUZ5L1JHRitkK3RkWjUrDQpabmprN0hESTQ5V3Rib0xodWN3YjBpNU4xbEVKWHVhOHhEN0FROTJXQy9PdzhzVktPRlJGNVJhMmxSa0svRS8xDQpxeDhKYnRoMGdXdHg0NHBQaWJwU3crMTB0QUhHR2FTLzVwN3hNUXhzajZTOThwL1hnalg5NzN4VStZZ2dLNUx3DQp6WlkzSDQ3UVREcmpyZzhOVmpDSFU3b3IrcEpCbjdldGF0V3psK3BQcVd4ODZub2tjdG5abUQxcHNnWnkwTEdDDQpRYys5ejdURGhEOFhuVERwckxiRGZXRnZqOTVKSmc3Q1krd29zN05vSENEOG5RWjFZZURVQkJjUkVlZXJVRlhBDQpaZ1I3UGNCN0FnTUJBQUVDZ2dFQUdkUWxCZVFER3ROMXJaTXNERGRUU2RudDU3b2NWdXlac2grNW1PRVVMakF3DQptOXhTOUt3ZnlYeTZRaXlxb3JKQ3REa2trR2c0OHhWOFBrVDN1RTBUTzU0aDF1UmxJMjNMbjcvVmFzOUZnVlFmDQpQS1dLVmdwYjdFMWhtT2gwVFNmRDRwRnpETlh4SzhMaXYycWVxdTJTTlRGWVR1M2RBRWpNL3EyWERmdXJQN2tiDQprZ3FKRFBwd2g4RWRXMVg1VVAyVE9CVWxwQllDTndxUkFJQ1E3eWlzbW5xeFlZS3RKc21MK21IQ3JYM3hNRHVTDQp4NHJCVDUvcXVrdVc4MmwrbGZmU3ZTNGpsb0VhajJ2QmozSk1udy9lYlNucFplU3FENTFjOUZlOCtocU4rU3NoDQozTnc0QXVybE1RRG5vZy9STUF3QUR3KzBRUlIwNVdaWDhMVXllVTBVVVFLQmdRRHd6R2I0b25BWHlNeUNkbHB3DQpRRnFCR0pwQnlsWEJjSkZnMGpCd1JMV2tQL3VjWnNoZlFlbkFWbkRZZS9yQ0FnWWxSdFFOVFRvb3BFSjlGcGgyDQp6TkVzd1EwcnV4WjFrVm41U1hwS2dF4668KalUxT3dGa3R1WFlJcEtBNGk5dFoxT04zb1lqdVRtMUlzb2xWZXVTDQpqK3Mwd1o3ZDAyYTNXcDN1UXJ3TFUwVjdpUUtCZ1FEaEcrc2xsNDYveGxONldWWEs3RVpkOGFEcTlTNEU0aEQvDQpvTmUwS0dVcHhZYngyTnFWN1VLSEwzOE41eG5qNGllWGt2U1BnL0twVUpqUmtLN0xJMnZsNmlndUJkdW01VUR1DQp5dW4rL1dNcVdnb2p4anZBbmxsS2lIa0JRMTJ2UFRqcE9HSGIrY0RqVWxROGVnOThFOEJ0ZktUQjFkRlcxUnBlDQorMXY0aXR3RzR3S0JnQzJLeXpMZExnd2hpeVJsbEFkRTlKa1QrU0RXVHMvT0pZREZZQ25ycE5zU3l0aXl5OVRRDQpWNUJzQ04yNDNSMVNXcTAwTHlqdzRUNE1peEt6Y2xTTnVrWVhvUkVUU2xVa0QzdEpmVnFYMVUrTE1XY0c2T1dPDQpmZndaMWRHUWRkM2dPL3BLQ3Q2NHlvUkt0eWJHa0U1ZzcrQkRlbk9ENXhwb2hoUXBCUDJ6V3lIWkFvR0FURndqDQpGUHBuUXVoc3Nza1JFQ2U3NnV3bkVPeWdjcW1ZNkkzUC9kM2lDeHhsSFM3Wlh4Zy9oQW41aUdiSFlvVDV0ekh6DQpZYWQ1cmpPWDB5YklGRUpzdkc0RXVTL2xoYVNvdFJnQjdpeFg4aXJlMjZuSDVSd1IzL1dSVG50aWtTb3NYdmh3DQpRYVZqNS9pcWVHVlRVVnlGM3QzMEtZaDFYWVltVHVmbkY5VktzODhDZ1lCTTNVN2QwOU9MemhMZTh3cnp1dEpuDQpGdmRGWlhCRnhXRGwyNXdteElWTFdpM0kvaWg2QXN5YlRNNWMzbFpTTUVvcjJYeXJqNnFUNzZ6amQ2eGE2NlN3DQpXMEVyL2lEY3dWK244MHpuU3lPSW5lRThIVkh1SGtNYVpPeHkvVzdVWDFqL0RmUnJPZG1iS1NWN2NBV2dVTlBrDQpnd1V5RkM2OTRKTR41Vko0WXZEZU13PT0NCi0tLS0tRU5EIFBSSVZBVEUgS0VZLS0tLS0="
  key_user = "ocid1.user.oc1..aaaaaaaad6n7rniiwgxehy6coo4ax2ti7pr5yr53cbdxdyp6sx6dhrttcz4a"
  key_tenant = "ocid1.tenancy.oc1..aaaaaaaaft6hqvl367uh4e3pmdxnzmca6cxamwjfaag5lm7bnhuwu6ypajca"
  key_region = "eu-frankfurt-1"
}

data "checkpoint_management_oracle_cloud_data_center_server" "data_oracle_cloud_data_center_server" {
  name = "${checkpoint_management_oracle_cloud_data_center_server.testOracleCloud.name}"
}
` + "`" + `` + "`" + `` + "`" + `

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
					Name:        "automatic_refresh",
					Description: `Indicates whether the data center server's content is automatically updated.`,
				},
				resource.Attribute{
					Name:        "data_center_type",
					Description: `Data center type.`,
				},
				resource.Attribute{
					Name:        "properties",
					Description: `Data center properties. properties blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag objects identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level. ` + "`" + `properties` + "`" + ` supports the following:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_policy_settings",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Policy Settings.`,
			Description: `

Use this data source to get information on an existing Check Point Policy Settings.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "last_in_cell",
					Description: `Added object after removing the last object in cell.`,
				},
				resource.Attribute{
					Name:        "none_object_behavior",
					Description: `'None' object behavior. Rules with object 'None' will never be matched.`,
				},
				resource.Attribute{
					Name:        "security_access_defaults",
					Description: `Access Policy default values. security_accesses_defaults blocks are documented below. ` + "`" + `security_access_defaults` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `Destination default value identified by name.`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `Service and Applications default value identified by name.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `Source default value identified by name.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_provisioning_profile",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Provisioning Profile.`,
			Description: `

Use this data source to get information on an existing Check Point Provisioning Profile.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_radius_group",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Radius Group.`,
			Description: `

Use this data source to get information on an existing Check Point Radius Group.

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
					Description: `Collection of radius servers identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag objects identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_radius_server",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Radius Server.`,
			Description: `

Use this data source to get information on an existing Check Point Radius Server.

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
					Name:        "server",
					Description: `The UID or Name of the host that is the RADIUS Server.`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `The UID or Name of the Service to which the RADIUS server listens.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version can be either RADIUS Version 1.0, which is RFC 2138 compliant, and RADIUS Version 2.0 which is RFC 2865 compliant.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The type of authentication protocol that will be used when authenticating the user to the RADIUS server.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority of the RADIUS Server in case it is a member of a RADIUS Group.`,
				},
				resource.Attribute{
					Name:        "accounting",
					Description: `Accounting settings. accounting blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag objects identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Color of the object. Should be one of existing colors.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments string. ` + "`" + `accounting` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "enable_ip_pool_management",
					Description: `IP pool management, enables Accounting service.`,
				},
				resource.Attribute{
					Name:        "accounting_service",
					Description: `The UID or Name of the the accounting interface to notify the server when users login and logout which will then lock and release the IP addresses that the server allocated to those users.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_repository_package",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Repository Package.`,
			Description: `

Use this data source to get information on an existing Check Point Repository Package.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the repository package. ## How To Use Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_repository_script",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Repository Script.`,
			Description: `

Use this data source to get information on an existing Check Point Repository Script.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
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
					Name:        "advanced_settings",
					Description: `N/Aadvanced_settings blocks are documented below.`,
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
					Name:        "application_control_and_url_filtering_settings",
					Description: `Gateway Application Control and URL filtering settings.application_control_and_url_filtering_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "cluster_mode",
					Description: `Cluster mode.`,
				},
				resource.Attribute{
					Name:        "cluster_settings",
					Description: `ClusterXL and VRRP Settings.cluster_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "content_awareness",
					Description: `Content Awareness blade enabled.`,
				},
				resource.Attribute{
					Name:        "enable_https_inspection",
					Description: `Enable HTTPS Inspection after defining an outbound inspection certificate. <br>To define the outbound certificate use outbound inspection certificate API.`,
				},
				resource.Attribute{
					Name:        "fetch_policy",
					Description: `Security management server(s) to fetch the policy from.fetch_policy blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "firewall",
					Description: `Firewall blade enabled.`,
				},
				resource.Attribute{
					Name:        "firewall_settings",
					Description: `N/Afirewall_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "geo_mode",
					Description: `Cluster High Availability Geo mode.<br>This setting applies only to a cluster deployed in a cloud. Available when the cluster mode equals "cluster-xl-ha".`,
				},
				resource.Attribute{
					Name:        "hardware",
					Description: `Cluster platform hardware.`,
				},
				resource.Attribute{
					Name:        "hit_count",
					Description: `Hit count tracks the number of connections each rule matches.`,
				},
				resource.Attribute{
					Name:        "https_inspection",
					Description: `HTTPS inspection.https_inspection blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "identity_awareness",
					Description: `Identity awareness blade enabled.`,
				},
				resource.Attribute{
					Name:        "identity_awareness_settings",
					Description: `Gateway Identity Awareness settings.identity_awareness_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `Cluster interfaces.interfaces blocks are documented below.`,
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
					Name:        "ips",
					Description: `Intrusion Prevention System blade enabled.`,
				},
				resource.Attribute{
					Name:        "ips_update_policy",
					Description: `Specifies whether the IPS will be downloaded from the Management or directly to the Gateway.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `Cluster members list. Only new cluster member can be added. Adding existing gateway is not supported.members blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "nat_hide_internal_interfaces",
					Description: `Hide internal networks behind the Gateway's external IP.`,
				},
				resource.Attribute{
					Name:        "nat_settings",
					Description: `NAT settings.nat_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "os_name",
					Description: `Cluster platform operating system.`,
				},
				resource.Attribute{
					Name:        "platform_portal_settings",
					Description: `Platform portal settings.platform_portal_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "proxy_settings",
					Description: `Proxy Server for Gateway.proxy_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "qos",
					Description: `QoS.`,
				},
				resource.Attribute{
					Name:        "send_alerts_to_server",
					Description: `Server(s) to send alerts to.send_alerts_to_server blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "send_logs_to_backup_server",
					Description: `Backup server(s) to send logs to.send_logs_to_backup_server blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "send_logs_to_server",
					Description: `Server(s) to send logs to.send_logs_to_server blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag identifiers.tags blocks are documented below.`,
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
					Name:        "threat_prevention_mode",
					Description: `The mode of Threat Prevention to use. When using Autonomous Threat Prevention, disabling the Threat Prevention blades is not allowed.`,
				},
				resource.Attribute{
					Name:        "url_filtering",
					Description: `URL Filtering blade enabled.`,
				},
				resource.Attribute{
					Name:        "usercheck_portal_settings",
					Description: `UserCheck portal settings.usercheck_portal_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Cluster platform version.`,
				},
				resource.Attribute{
					Name:        "vpn",
					Description: `VPN blade enabled.`,
				},
				resource.Attribute{
					Name:        "vpn_settings",
					Description: `Gateway VPN settings.vpn_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "zero_phishing",
					Description: `Zero Phishing blade enabled.`,
				},
				resource.Attribute{
					Name:        "zero_phishing_fqdn",
					Description: `Zero Phishing gateway FQDN.`,
				},
				resource.Attribute{
					Name:        "show_portals_certificate",
					Description: `Indicates whether to show the portals certificate value in the reply.`,
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
					Description: `Collection of group identifiers.groups blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "ignore_warnings",
					Description: `Apply changes ignoring warnings.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. ` + "`" + `advanced_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "connection_persistence",
					Description: `Handling established connections when installing a new policy.`,
				},
				resource.Attribute{
					Name:        "sam",
					Description: `SAM.sam blocks are documented below. ` + "`" + `application_control_and_url_filtering_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "global_settings_mode",
					Description: `Whether to override global settings or not.`,
				},
				resource.Attribute{
					Name:        "override_global_settings",
					Description: `override global settings object.override_global_settings blocks are documented below. ` + "`" + `cluster_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "member_recovery_mode",
					Description: `In a High Availability cluster, each member is given a priority. The member with the highest priority serves as the gateway. If this gateway fails, control is passed to the member with the next highest priority. If that member fails, control is passed to the next, and so on. Upon gateway recovery, it is possible to: Maintain current active Cluster Member (maintain-current-active) or Switch to higher priority Cluster Member (according-to-priority).`,
				},
				resource.Attribute{
					Name:        "state_synchronization",
					Description: `Cluster State Synchronization settings.state_synchronization blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "track_changes_of_cluster_members",
					Description: `Track changes in the status of Cluster Members.`,
				},
				resource.Attribute{
					Name:        "use_virtual_mac",
					Description: `Use Virtual MAC. By enabling Virtual MAC in ClusterXL High Availability New mode, or Load Sharing Unicast mode, all cluster members associate the same Virtual MAC address with All Cluster Virtual Interfaces and the Virtual IP address. ` + "`" + `firewall_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "auto_calculate_connections_hash_table_size_and_memory_pool",
					Description: `N/A`,
				},
				resource.Attribute{
					Name:        "auto_maximum_limit_for_concurrent_connections",
					Description: `N/A`,
				},
				resource.Attribute{
					Name:        "connections_hash_size",
					Description: `N/A`,
				},
				resource.Attribute{
					Name:        "maximum_limit_for_concurrent_connections",
					Description: `N/A`,
				},
				resource.Attribute{
					Name:        "maximum_memory_pool_size",
					Description: `N/A`,
				},
				resource.Attribute{
					Name:        "memory_pool_size",
					Description: `N/A ` + "`" + `https_inspection` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "bypass_on_failure",
					Description: `Set to be true in order to bypass all requests (Fail-open) in case of internal system error.bypass_on_failure blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "site_categorization_allow_mode",
					Description: `Set to 'background' in order to allowed requests until categorization is complete.site_categorization_allow_mode blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "deny_untrusted_server_cert",
					Description: `Set to be true in order to drop traffic from servers with untrusted server certificate.deny_untrusted_server_cert blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "deny_revoked_server_cert",
					Description: `Set to be true in order to drop traffic from servers with revoked server certificate (validate CRL).deny_revoked_server_cert blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "deny_expired_server_cert",
					Description: `Set to be true in order to drop traffic from servers with expired server certificate.deny_expired_server_cert blocks are documented below. ` + "`" + `identity_awareness_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "browser_based_authentication",
					Description: `Enable Browser Based Authentication source.`,
				},
				resource.Attribute{
					Name:        "browser_based_authentication_settings",
					Description: `Browser Based Authentication settings.browser_based_authentication_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "identity_agent",
					Description: `Enable Identity Agent source.`,
				},
				resource.Attribute{
					Name:        "identity_agent_settings",
					Description: `Identity Agent settings.identity_agent_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "identity_collector",
					Description: `Enable Identity Collector source.`,
				},
				resource.Attribute{
					Name:        "identity_collector_settings",
					Description: `Identity Collector settings.identity_collector_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "identity_sharing_settings",
					Description: `Identity sharing settings.identity_sharing_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "proxy_settings",
					Description: `Identity-Awareness Proxy settings.proxy_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "remote_access",
					Description: `Enable Remote Access Identity source. ` + "`" + `interfaces` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Object name. Must be unique in the domain.`,
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
					Name:        "network_mask",
					Description: `IPv4 or IPv6 network mask. If both masks are required use ipv4-network-mask and ipv6-network-mask fields explicitly. Instead of providing mask itself it is possible to specify IPv4 or IPv6 mask length in mask-length field. If both masks length are required use ipv4-mask-length and ipv6-mask-length fields explicitly.`,
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
					Description: `N/A`,
				},
				resource.Attribute{
					Name:        "anti_spoofing_settings",
					Description: `N/Aanti_spoofing_settings blocks are documented below.`,
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
					Description: `N/A`,
				},
				resource.Attribute{
					Name:        "security_zone_settings",
					Description: `N/Asecurity_zone_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "topology",
					Description: `N/A`,
				},
				resource.Attribute{
					Name:        "topology_settings",
					Description: `N/Atopology_settings blocks are documented below.`,
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
					Description: `Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. ` + "`" + `members` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Object name.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `Cluster Member network interfaces.interfaces blocks are documented below.`,
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
					Name:        "one_time_password",
					Description: `N/A`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `In a High Availability New mode cluster each machine is given a priority. The highest priority machine serves as the gateway in normal circumstances. If this machine fails, control is passed to the next highest priority machine. If that machine fails, control is passed to the next machine, and so on. In Load Sharing Unicast mode cluster, the highest priority is the pivot machine. The values must be in a range from 1 to N, where N is number of cluster members.`,
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
					Description: `Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. ` + "`" + `nat_settings` + "`" + ` supports the following:`,
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
					Description: `Hide behind method. This parameter is forbidden in case "method" parameter is "static".`,
				},
				resource.Attribute{
					Name:        "install_on",
					Description: `Which gateway should apply the NAT translation.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `NAT translation method. ` + "`" + `platform_portal_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "portal_web_settings",
					Description: `Configuration of the portal web settings.portal_web_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "certificate_settings",
					Description: `Configuration of the portal certificate settings.certificate_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "accessibility",
					Description: `Configuration of the portal access settings.accessibility blocks are documented below. ` + "`" + `proxy_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "use_custom_proxy",
					Description: `Use custom proxy settings for this network object.`,
				},
				resource.Attribute{
					Name:        "proxy_server",
					Description: `N/A`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `N/A ` + "`" + `usercheck_portal_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `State of the web portal (enabled or disabled). The supported blades are: {'Application Control', 'URL Filtering', 'Data Loss Prevention', 'Anti Virus', 'Anti Bot', 'Threat Emulation', 'Threat Extraction', 'Data Awareness'}.`,
				},
				resource.Attribute{
					Name:        "portal_web_settings",
					Description: `Configuration of the portal web settings.portal_web_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "certificate_settings",
					Description: `Configuration of the portal certificate settings.certificate_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "accessibility",
					Description: `Configuration of the portal access settings.accessibility blocks are documented below. ` + "`" + `vpn_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "authentication",
					Description: `Authentication.authentication blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "link_selection",
					Description: `Link Selection.link_selection blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "maximum_concurrent_ike_negotiations",
					Description: `N/A`,
				},
				resource.Attribute{
					Name:        "maximum_concurrent_tunnels",
					Description: `N/A`,
				},
				resource.Attribute{
					Name:        "office_mode",
					Description: `Office Mode. Notation Wide Impact - Office Mode apply IPSec VPN Software Blade clients and to the Mobile Access Software Blade clients.office_mode blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "remote_access",
					Description: `Remote Access.remote_access blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "vpn_domain",
					Description: `Gateway VPN domain identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "vpn_domain_exclude_external_ip_addresses",
					Description: `Exclude the external IP addresses from the VPN domain of this Security Gateway.`,
				},
				resource.Attribute{
					Name:        "vpn_domain_type",
					Description: `Gateway VPN domain type. ` + "`" + `sam` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "forward_to_other_sam_servers",
					Description: `Forward SAM clients' requests to other SAM servers.`,
				},
				resource.Attribute{
					Name:        "use_early_versions",
					Description: `Use early versions compatibility mode.use_early_versions blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "purge_sam_file",
					Description: `Purge SAM File.purge_sam_file blocks are documented below. ` + "`" + `override_global_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "fail_mode",
					Description: `Fail mode - allow or block all requests.`,
				},
				resource.Attribute{
					Name:        "website_categorization",
					Description: `Website categorization object.website_categorization blocks are documented below. ` + "`" + `state_synchronization` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "delayed",
					Description: `Start synchronizing with delay of seconds, as defined by delayed-seconds, after connection initiation. Disabled when state-synchronization disabled.`,
				},
				resource.Attribute{
					Name:        "delayed_seconds",
					Description: `Start synchronizing X seconds after connection initiation . The values must be in a range between 2 and 3600.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Use State Synchronization. ` + "`" + `bypass_on_failure` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "override_profile",
					Description: `Override profile of global configuration.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Override value.<br><font color="red">Required only for</font> 'override-profile' is True. ` + "`" + `site_categorization_allow_mode` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "override_profile",
					Description: `Override profile of global configuration.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Override value.<br><font color="red">Required only for</font> 'override-profile' is True. ` + "`" + `deny_untrusted_server_cert` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "override_profile",
					Description: `Override profile of global configuration.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Override value.<br><font color="red">Required only for</font> 'override-profile' is True. ` + "`" + `deny_revoked_server_cert` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "override_profile",
					Description: `Override profile of global configuration.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Override value.<br><font color="red">Required only for</font> 'override-profile' is True. ` + "`" + `deny_expired_server_cert` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "override_profile",
					Description: `Override profile of global configuration.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Override value.<br><font color="red">Required only for</font> 'override-profile' is True. ` + "`" + `browser_based_authentication_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "authentication_settings",
					Description: `Authentication Settings for Browser Based Authentication.authentication_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "browser_based_authentication_portal_settings",
					Description: `Browser Based Authentication portal settings.browser_based_authentication_portal_settings blocks are documented below. ` + "`" + `identity_agent_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "agents_interval_keepalive",
					Description: `Agents send keepalive period (minutes).`,
				},
				resource.Attribute{
					Name:        "user_reauthenticate_interval",
					Description: `Agent reauthenticate time interval (minutes).`,
				},
				resource.Attribute{
					Name:        "authentication_settings",
					Description: `Authentication Settings for Identity Agent.authentication_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "identity_agent_portal_settings",
					Description: `Identity Agent accessibility settings.identity_agent_portal_settings blocks are documented below. ` + "`" + `identity_collector_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "authorized_clients",
					Description: `Authorized Clients.authorized_clients blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "authentication_settings",
					Description: `Authentication Settings for Identity Collector.authentication_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "client_access_permissions",
					Description: `Identity Collector accessibility settings.client_access_permissions blocks are documented below. ` + "`" + `identity_sharing_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "share_with_other_gateways",
					Description: `Enable identity sharing with other gateways.`,
				},
				resource.Attribute{
					Name:        "receive_from_other_gateways",
					Description: `Enable receiving identity from other gateways.`,
				},
				resource.Attribute{
					Name:        "receive_from",
					Description: `Gateway(s) to receive identity from.receive_from blocks are documented below. ` + "`" + `proxy_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "detect_using_x_forward_for",
					Description: `Whether to use X-Forward-For HTTP header, which is added by the proxy server to keep track of the original source IP. ` + "`" + `anti_spoofing_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `If packets will be rejected (the Prevent option) or whether the packets will be monitored (the Detect option).`,
				},
				resource.Attribute{
					Name:        "exclude_packets",
					Description: `Don't check packets from excluded network.`,
				},
				resource.Attribute{
					Name:        "excluded_network_name",
					Description: `Excluded network name.`,
				},
				resource.Attribute{
					Name:        "excluded_network_uid",
					Description: `Excluded network UID.`,
				},
				resource.Attribute{
					Name:        "spoof_tracking",
					Description: `Spoof tracking. ` + "`" + `security_zone_settings` + "`" + ` supports the following:`,
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
					Name:        "specific_network",
					Description: `Network behind this interface. ` + "`" + `interfaces` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Object name.`,
				},
				resource.Attribute{
					Name:        "anti_spoofing",
					Description: `N/A`,
				},
				resource.Attribute{
					Name:        "anti_spoofing_settings",
					Description: `N/Aanti_spoofing_settings blocks are documented below.`,
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
					Name:        "network_mask",
					Description: `IPv4 or IPv6 network mask. If both masks are required use ipv4-network-mask and ipv6-network-mask fields explicitly. Instead of providing mask itself it is possible to specify IPv4 or IPv6 mask length in mask-length field. If both masks length are required use ipv4-mask-length and ipv6-mask-length fields explicitly.`,
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
					Name:        "security_zone",
					Description: `N/A`,
				},
				resource.Attribute{
					Name:        "security_zone_settings",
					Description: `N/Asecurity_zone_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "topology",
					Description: `N/A`,
				},
				resource.Attribute{
					Name:        "topology_settings",
					Description: `N/Atopology_settings blocks are documented below.`,
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
					Description: `Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. ` + "`" + `portal_web_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "aliases",
					Description: `List of URL aliases that are redirected to the main portal URL.aliases blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "main_url",
					Description: `The main URL for the web portal. ` + "`" + `certificate_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "base64_certificate",
					Description: `The certificate file encoded in Base64 with padding. This file must be in the`,
				},
				resource.Attribute{
					Name:        "base64_password",
					Description: `Password (encoded in Base64 with padding) for the certificate file. ` + "`" + `accessibility` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "allow_access_from",
					Description: `Allowed access to the web portal (based on interfaces, or security policy).`,
				},
				resource.Attribute{
					Name:        "internal_access_settings",
					Description: `Configuration of the additional portal access settings for internal interfaces only.internal_access_settings blocks are documented below. ` + "`" + `portal_web_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "aliases",
					Description: `List of URL aliases that are redirected to the main portal URL.aliases blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "main_url",
					Description: `The main URL for the web portal. ` + "`" + `certificate_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "base64_certificate",
					Description: `The certificate file encoded in Base64 with padding. This file must be in the`,
				},
				resource.Attribute{
					Name:        "base64_password",
					Description: `Password (encoded in Base64 with padding) for the certificate file. ` + "`" + `accessibility` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "allow_access_from",
					Description: `Allowed access to the web portal (based on interfaces, or security policy).`,
				},
				resource.Attribute{
					Name:        "internal_access_settings",
					Description: `Configuration of the additional portal access settings for internal interfaces only.internal_access_settings blocks are documented below. ` + "`" + `authentication` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "authentication_clients",
					Description: `Collection of VPN Authentication clients identified by the name or UID.authentication_clients blocks are documented below. ` + "`" + `link_selection` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "dns_resolving_hostname",
					Description: `DNS Resolving Hostname. Must be set when "ip-selection" was selected to be "dns-resolving-from-hostname". ` + "`" + `office_mode` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Office Mode Permissions. When selected to be "off", all the other definitions are irrelevant.`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `Group. Identified by name or UID. Must be set when "office-mode-permissions" was selected to be "group".`,
				},
				resource.Attribute{
					Name:        "allocate_ip_address_from",
					Description: `Allocate IP address Method. Allocate IP address by sequentially trying the given methods until success.allocate_ip_address_from blocks are documented below.`,
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
					Description: `Additional IP Addresses for Anti-Spoofing. Identified by name or UID. Must be set when "perform-anti-spoofings" is true. ` + "`" + `remote_access` + "`" + ` supports the following:`,
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
					Description: `Allocated NAT traversal UDP service. Identified by name or UID. Must be set when "support-nat-traversal-mechanism" is true.`,
				},
				resource.Attribute{
					Name:        "support_visitor_mode",
					Description: `Support Visitor Mode.`,
				},
				resource.Attribute{
					Name:        "visitor_mode_service",
					Description: `TCP Service for Visitor Mode. Identified by name or UID. Must be set when "support-visitor-mode" is true.`,
				},
				resource.Attribute{
					Name:        "visitor_mode_interface",
					Description: `Interface for Visitor Mode. Must be set when "support-visitor-mode" is true. Insert IPV4 Address of existing interface or "All IPs" when you want all interfaces. ` + "`" + `use_early_versions` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Use early versions compatibility mode.`,
				},
				resource.Attribute{
					Name:        "compatibility_mode",
					Description: `Early versions compatibility mode. ` + "`" + `purge_sam_file` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Purge SAM File.`,
				},
				resource.Attribute{
					Name:        "purge_when_size_reaches_to",
					Description: `Purge SAM File When it Reaches to. ` + "`" + `website_categorization` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Website categorization mode.`,
				},
				resource.Attribute{
					Name:        "custom_mode",
					Description: `Custom mode object.custom_mode blocks are documented below. ` + "`" + `authentication_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "authentication_method",
					Description: `Authentication method.`,
				},
				resource.Attribute{
					Name:        "identity_provider",
					Description: `Identity provider object identified by the name or UID. Must be set when "authentication-method" was selected to be "identity provider".identity_provider blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "radius",
					Description: `Radius server object identified by the name or UID. Must be set when "authentication-method" was selected to be "radius".`,
				},
				resource.Attribute{
					Name:        "users_directories",
					Description: `Users directories.users_directories blocks are documented below. ` + "`" + `browser_based_authentication_portal_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "portal_web_settings",
					Description: `Configuration of the portal web settings.portal_web_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "certificate_settings",
					Description: `Configuration of the portal certificate settings.certificate_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "accessibility",
					Description: `Configuration of the portal access settings.accessibility blocks are documented below. ` + "`" + `authentication_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "authentication_method",
					Description: `Authentication method.`,
				},
				resource.Attribute{
					Name:        "radius",
					Description: `Radius server object identified by the name or UID. Must be set when "authentication-method" was selected to be "radius".`,
				},
				resource.Attribute{
					Name:        "users_directories",
					Description: `Users directories.users_directories blocks are documented below. ` + "`" + `identity_agent_portal_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "accessibility",
					Description: `Configuration of the portal access settings.accessibility blocks are documented below. ` + "`" + `authorized_clients` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "client",
					Description: `Host / Network Group Name or UID.`,
				},
				resource.Attribute{
					Name:        "client_secret",
					Description: `Client Secret. ` + "`" + `authentication_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "users_directories",
					Description: `Users directories.users_directories blocks are documented below. ` + "`" + `client_access_permissions` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "accessibility",
					Description: `Configuration of the portal access settings.accessibility blocks are documented below. ` + "`" + `anti_spoofing_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `If packets will be rejected (the Prevent option) or whether the packets will be monitored (the Detect option).`,
				},
				resource.Attribute{
					Name:        "exclude_packets",
					Description: `Don't check packets from excluded network.`,
				},
				resource.Attribute{
					Name:        "excluded_network_name",
					Description: `Excluded network name.`,
				},
				resource.Attribute{
					Name:        "excluded_network_uid",
					Description: `Excluded network UID.`,
				},
				resource.Attribute{
					Name:        "spoof_tracking",
					Description: `Spoof tracking. ` + "`" + `security_zone_settings` + "`" + ` supports the following:`,
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
					Name:        "specific_network",
					Description: `Network behind this interface. ` + "`" + `internal_access_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "undefined",
					Description: `Controls portal access settings for internal interfaces, whose topology is set to 'Undefined'.`,
				},
				resource.Attribute{
					Name:        "dmz",
					Description: `Controls portal access settings for internal interfaces, whose topology is set to 'DMZ'.`,
				},
				resource.Attribute{
					Name:        "vpn",
					Description: `Controls portal access settings for interfaces that are part of a VPN Encryption Domain. ` + "`" + `internal_access_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "undefined",
					Description: `Controls portal access settings for internal interfaces, whose topology is set to 'Undefined'.`,
				},
				resource.Attribute{
					Name:        "dmz",
					Description: `Controls portal access settings for internal interfaces, whose topology is set to 'DMZ'.`,
				},
				resource.Attribute{
					Name:        "vpn",
					Description: `Controls portal access settings for interfaces that are part of a VPN Encryption Domain. ` + "`" + `allocate_ip_address_from` + "`" + ` supports the following:`,
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
					Description: `Manual Network. Identified by name or UID. Must be set when "allocate-method" was selected to be "manual".`,
				},
				resource.Attribute{
					Name:        "dhcp_server",
					Description: `DHCP Server. Identified by name or UID. Must be set when "allocate-method" was selected to be "automatic".`,
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
					Description: `This configuration applies to all Office Mode methods except Automatic (using DHCP) and ipassignment.conf entries which contain this data.optional_parameters blocks are documented below. ` + "`" + `custom_mode` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "social_networking_widgets",
					Description: `Social networking widgets mode.`,
				},
				resource.Attribute{
					Name:        "url_filtering",
					Description: `URL filtering mode. ` + "`" + `users_directories` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "external_user_profile",
					Description: `External user profile.`,
				},
				resource.Attribute{
					Name:        "internal_users",
					Description: `Internal users.`,
				},
				resource.Attribute{
					Name:        "users_from_external_directories",
					Description: `Users from external directories.`,
				},
				resource.Attribute{
					Name:        "specific",
					Description: `LDAP AU objects identified by the name or UID. Must be set when "users-from-external-directories" was selected to be "specific".specific blocks are documented below. ` + "`" + `portal_web_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "aliases",
					Description: `List of URL aliases that are redirected to the main portal URL.aliases blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "main_url",
					Description: `The main URL for the web portal. ` + "`" + `certificate_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "base64_certificate",
					Description: `The certificate file encoded in Base64 with padding. This file must be in the`,
				},
				resource.Attribute{
					Name:        "base64_password",
					Description: `Password (encoded in Base64 with padding) for the certificate file. ` + "`" + `accessibility` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "allow_access_from",
					Description: `Allowed access to the web portal (based on interfaces, or security policy).`,
				},
				resource.Attribute{
					Name:        "internal_access_settings",
					Description: `Configuration of the additional portal access settings for internal interfaces only.internal_access_settings blocks are documented below. ` + "`" + `users_directories` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "external_user_profile",
					Description: `External user profile.`,
				},
				resource.Attribute{
					Name:        "internal_users",
					Description: `Internal users.`,
				},
				resource.Attribute{
					Name:        "users_from_external_directories",
					Description: `Users from external directories.`,
				},
				resource.Attribute{
					Name:        "specific",
					Description: `LDAP AU objects identified by the name or UID. Must be set when "users-from-external-directories" was selected to be "specific".specific blocks are documented below. ` + "`" + `accessibility` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "allow_access_from",
					Description: `Allowed access to the web portal (based on interfaces, or security policy).`,
				},
				resource.Attribute{
					Name:        "internal_access_settings",
					Description: `Configuration of the additional portal access settings for internal interfaces only.internal_access_settings blocks are documented below. ` + "`" + `users_directories` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "external_user_profile",
					Description: `External user profile.`,
				},
				resource.Attribute{
					Name:        "internal_users",
					Description: `Internal users.`,
				},
				resource.Attribute{
					Name:        "users_from_external_directories",
					Description: `Users from external directories.`,
				},
				resource.Attribute{
					Name:        "specific",
					Description: `LDAP AU objects identified by the name or UID. Must be set when "users-from-external-directories" was selected to be "specific".specific blocks are documented below. ` + "`" + `accessibility` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "allow_access_from",
					Description: `Allowed access to the web portal (based on interfaces, or security policy).`,
				},
				resource.Attribute{
					Name:        "internal_access_settings",
					Description: `Configuration of the additional portal access settings for internal interfaces only.internal_access_settings blocks are documented below. ` + "`" + `optional_parameters` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "use_primary_dns_server",
					Description: `Use Primary DNS Server.`,
				},
				resource.Attribute{
					Name:        "primary_dns_server",
					Description: `Primary DNS Server. Identified by name or UID. Must be set when "use-primary-dns-server" is true and can not be set when "use-primary-dns-server" is false.`,
				},
				resource.Attribute{
					Name:        "use_first_backup_dns_server",
					Description: `Use First Backup DNS Server.`,
				},
				resource.Attribute{
					Name:        "first_backup_dns_server",
					Description: `First Backup DNS Server. Identified by name or UID. Must be set when "use-first-backup-dns-server" is true and can not be set when "use-first-backup-dns-server" is false.`,
				},
				resource.Attribute{
					Name:        "use_second_backup_dns_server",
					Description: `Use Second Backup DNS Server.`,
				},
				resource.Attribute{
					Name:        "second_backup_dns_server",
					Description: `Second Backup DNS Server. Identified by name or UID. Must be set when "use-second-backup-dns-server" is true and can not be set when "use-second-backup-dns-server" is false.`,
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
					Description: `Primary WINS Server. Identified by name or UID. Must be set when "use-primary-wins-server" is true and can not be set when "use-primary-wins-server" is false.`,
				},
				resource.Attribute{
					Name:        "use_first_backup_wins_server",
					Description: `Use First Backup WINS Server.`,
				},
				resource.Attribute{
					Name:        "first_backup_wins_server",
					Description: `First Backup WINS Server. Identified by name or UID. Must be set when "use-first-backup-wins-server" is true and can not be set when "use-first-backup-wins-server" is false.`,
				},
				resource.Attribute{
					Name:        "use_second_backup_wins_server",
					Description: `Use Second Backup WINS Server.`,
				},
				resource.Attribute{
					Name:        "second_backup_wins_server",
					Description: `Second Backup WINS Server. Identified by name or UID. Must be set when "use-second-backup-wins-server" is true and can not be set when "use-second-backup-wins-server" is false. ` + "`" + `internal_access_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "undefined",
					Description: `Controls portal access settings for internal interfaces, whose topology is set to 'Undefined'.`,
				},
				resource.Attribute{
					Name:        "dmz",
					Description: `Controls portal access settings for internal interfaces, whose topology is set to 'DMZ'.`,
				},
				resource.Attribute{
					Name:        "vpn",
					Description: `Controls portal access settings for interfaces that are part of a VPN Encryption Domain. ` + "`" + `internal_access_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "undefined",
					Description: `Controls portal access settings for internal interfaces, whose topology is set to 'Undefined'.`,
				},
				resource.Attribute{
					Name:        "dmz",
					Description: `Controls portal access settings for internal interfaces, whose topology is set to 'DMZ'.`,
				},
				resource.Attribute{
					Name:        "vpn",
					Description: `Controls portal access settings for interfaces that are part of a VPN Encryption Domain. ` + "`" + `internal_access_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "undefined",
					Description: `Controls portal access settings for internal interfaces, whose topology is set to 'Undefined'.`,
				},
				resource.Attribute{
					Name:        "dmz",
					Description: `Controls portal access settings for internal interfaces, whose topology is set to 'DMZ'.`,
				},
				resource.Attribute{
					Name:        "vpn",
					Description: `Controls portal access settings for interfaces that are part of a VPN Encryption Domain.`,
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
					Name:        "advanced_settings",
					Description: `N/Aadvanced_settings blocks are documented below.`,
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
					Name:        "application_control_and_url_filtering_settings",
					Description: `Gateway Application Control and URL filtering settings.application_control_and_url_filtering_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "content_awareness",
					Description: `Content Awareness blade enabled.`,
				},
				resource.Attribute{
					Name:        "enable_https_inspection",
					Description: `Enable HTTPS Inspection after defining an outbound inspection certificate. <br>To define the outbound certificate use outbound inspection certificate API.`,
				},
				resource.Attribute{
					Name:        "fetch_policy",
					Description: `Security management server(s) to fetch the policy from.fetch_policy blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "firewall",
					Description: `Firewall blade enabled.`,
				},
				resource.Attribute{
					Name:        "firewall_settings",
					Description: `N/Afirewall_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "hit_count",
					Description: `Hit count tracks the number of connections each rule matches.`,
				},
				resource.Attribute{
					Name:        "https_inspection",
					Description: `HTTPS inspection.https_inspection blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "icap_server",
					Description: `ICAP Server enabled.`,
				},
				resource.Attribute{
					Name:        "identity_awareness",
					Description: `Identity awareness blade enabled.`,
				},
				resource.Attribute{
					Name:        "identity_awareness_settings",
					Description: `Gateway Identity Awareness settings.identity_awareness_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `Network interfaces.interfaces blocks are documented below.`,
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
					Name:        "ips",
					Description: `Intrusion Prevention System blade enabled.`,
				},
				resource.Attribute{
					Name:        "ips_update_policy",
					Description: `Specifies whether the IPS will be downloaded from the Management or directly to the Gateway.`,
				},
				resource.Attribute{
					Name:        "nat_hide_internal_interfaces",
					Description: `Hide internal networks behind the Gateway's external IP.`,
				},
				resource.Attribute{
					Name:        "nat_settings",
					Description: `NAT settings.nat_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "one_time_password",
					Description: `N/A`,
				},
				resource.Attribute{
					Name:        "os_name",
					Description: `Gateway platform operating system.`,
				},
				resource.Attribute{
					Name:        "platform_portal_settings",
					Description: `Platform portal settings.platform_portal_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "proxy_settings",
					Description: `Proxy Server for Gateway.proxy_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "qos",
					Description: `QoS.`,
				},
				resource.Attribute{
					Name:        "save_logs_locally",
					Description: `Save logs locally on the gateway.`,
				},
				resource.Attribute{
					Name:        "send_alerts_to_server",
					Description: `Server(s) to send alerts to.send_alerts_to_server blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "send_logs_to_backup_server",
					Description: `Backup server(s) to send logs to.send_logs_to_backup_server blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "send_logs_to_server",
					Description: `Server(s) to send logs to.send_logs_to_server blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag identifiers.tags blocks are documented below.`,
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
					Name:        "threat_prevention_mode",
					Description: `The mode of Threat Prevention to use. When using Autonomous Threat Prevention, disabling the Threat Prevention blades is not allowed.`,
				},
				resource.Attribute{
					Name:        "url_filtering",
					Description: `URL Filtering blade enabled.`,
				},
				resource.Attribute{
					Name:        "usercheck_portal_settings",
					Description: `UserCheck portal settings.usercheck_portal_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Gateway platform version.`,
				},
				resource.Attribute{
					Name:        "vpn",
					Description: `VPN blade enabled.`,
				},
				resource.Attribute{
					Name:        "vpn_settings",
					Description: `Gateway VPN settings.vpn_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "zero_phishing",
					Description: `Zero Phishing blade enabled.`,
				},
				resource.Attribute{
					Name:        "zero_phishing_fqdn",
					Description: `Zero Phishing gateway FQDN.`,
				},
				resource.Attribute{
					Name:        "logs_settings",
					Description: `Logs settings that apply to Quantum Security Gateways that run Gaia OS.logs_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "show_portals_certificate",
					Description: `Indicates whether to show the portals certificate value in the reply.`,
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
					Description: `Collection of group identifiers.groups blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "ignore_errors",
					Description: `Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. ` + "`" + `advanced_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "connection_persistence",
					Description: `Handling established connections when installing a new policy.`,
				},
				resource.Attribute{
					Name:        "sam",
					Description: `SAM.sam blocks are documented below. ` + "`" + `application_control_and_url_filtering_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "global_settings_mode",
					Description: `Whether to override global settings or not.`,
				},
				resource.Attribute{
					Name:        "override_global_settings",
					Description: `override global settings object.override_global_settings blocks are documented below. ` + "`" + `firewall_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "auto_calculate_connections_hash_table_size_and_memory_pool",
					Description: `N/A`,
				},
				resource.Attribute{
					Name:        "auto_maximum_limit_for_concurrent_connections",
					Description: `N/A`,
				},
				resource.Attribute{
					Name:        "connections_hash_size",
					Description: `N/A`,
				},
				resource.Attribute{
					Name:        "maximum_limit_for_concurrent_connections",
					Description: `N/A`,
				},
				resource.Attribute{
					Name:        "maximum_memory_pool_size",
					Description: `N/A`,
				},
				resource.Attribute{
					Name:        "memory_pool_size",
					Description: `N/A ` + "`" + `https_inspection` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "bypass_on_failure",
					Description: `Set to be true in order to bypass all requests (Fail-open) in case of internal system error.bypass_on_failure blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "site_categorization_allow_mode",
					Description: `Set to 'background' in order to allowed requests until categorization is complete.site_categorization_allow_mode blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "deny_untrusted_server_cert",
					Description: `Set to be true in order to drop traffic from servers with untrusted server certificate.deny_untrusted_server_cert blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "deny_revoked_server_cert",
					Description: `Set to be true in order to drop traffic from servers with revoked server certificate (validate CRL).deny_revoked_server_cert blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "deny_expired_server_cert",
					Description: `Set to be true in order to drop traffic from servers with expired server certificate.deny_expired_server_cert blocks are documented below. ` + "`" + `identity_awareness_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "browser_based_authentication",
					Description: `Enable Browser Based Authentication source.`,
				},
				resource.Attribute{
					Name:        "browser_based_authentication_settings",
					Description: `Browser Based Authentication settings.browser_based_authentication_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "identity_agent",
					Description: `Enable Identity Agent source.`,
				},
				resource.Attribute{
					Name:        "identity_agent_settings",
					Description: `Identity Agent settings.identity_agent_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "identity_collector",
					Description: `Enable Identity Collector source.`,
				},
				resource.Attribute{
					Name:        "identity_collector_settings",
					Description: `Identity Collector settings.identity_collector_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "identity_sharing_settings",
					Description: `Identity sharing settings.identity_sharing_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "proxy_settings",
					Description: `Identity-Awareness Proxy settings.proxy_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "remote_access",
					Description: `Enable Remote Access Identity source. ` + "`" + `interfaces` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Object name. Must be unique in the domain.`,
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
					Name:        "network_mask",
					Description: `IPv4 or IPv6 network mask. If both masks are required use ipv4-network-mask and ipv6-network-mask fields explicitly. Instead of providing mask itself it is possible to specify IPv4 or IPv6 mask length in mask-length field. If both masks length are required use ipv4-mask-length and ipv6-mask-length fields explicitly.`,
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
					Description: `N/A`,
				},
				resource.Attribute{
					Name:        "anti_spoofing_settings",
					Description: `N/Aanti_spoofing_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "security_zone",
					Description: `N/A`,
				},
				resource.Attribute{
					Name:        "security_zone_settings",
					Description: `N/Asecurity_zone_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Collection of tag identifiers.tags blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "topology",
					Description: `N/A`,
				},
				resource.Attribute{
					Name:        "topology_settings",
					Description: `N/Atopology_settings blocks are documented below.`,
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
					Name:        "ignore_errors",
					Description: `Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. ` + "`" + `nat_settings` + "`" + ` supports the following:`,
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
					Description: `Hide behind method. This parameter is forbidden in case "method" parameter is "static".`,
				},
				resource.Attribute{
					Name:        "install_on",
					Description: `Which gateway should apply the NAT translation.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `NAT translation method. ` + "`" + `platform_portal_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "portal_web_settings",
					Description: `Configuration of the portal web settings.portal_web_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "certificate_settings",
					Description: `Configuration of the portal certificate settings.certificate_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "accessibility",
					Description: `Configuration of the portal access settings.accessibility blocks are documented below. ` + "`" + `proxy_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "use_custom_proxy",
					Description: `Use custom proxy settings for this network object.`,
				},
				resource.Attribute{
					Name:        "proxy_server",
					Description: `N/A`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `N/A ` + "`" + `usercheck_portal_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `State of the web portal (enabled or disabled). The supported blades are: {'Application Control', 'URL Filtering', 'Data Loss Prevention', 'Anti Virus', 'Anti Bot', 'Threat Emulation', 'Threat Extraction', 'Data Awareness'}.`,
				},
				resource.Attribute{
					Name:        "portal_web_settings",
					Description: `Configuration of the portal web settings.portal_web_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "certificate_settings",
					Description: `Configuration of the portal certificate settings.certificate_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "accessibility",
					Description: `Configuration of the portal access settings.accessibility blocks are documented below. ` + "`" + `vpn_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "authentication",
					Description: `Authentication.authentication blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "link_selection",
					Description: `Link Selection.link_selection blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "maximum_concurrent_ike_negotiations",
					Description: `N/A`,
				},
				resource.Attribute{
					Name:        "maximum_concurrent_tunnels",
					Description: `N/A`,
				},
				resource.Attribute{
					Name:        "office_mode",
					Description: `Office Mode. Notation Wide Impact - Office Mode apply IPSec VPN Software Blade clients and to the Mobile Access Software Blade clients.office_mode blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "remote_access",
					Description: `Remote Access.remote_access blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "vpn_domain",
					Description: `Gateway VPN domain identified by the name or UID.`,
				},
				resource.Attribute{
					Name:        "vpn_domain_exclude_external_ip_addresses",
					Description: `Exclude the external IP addresses from the VPN domain of this Security Gateway.`,
				},
				resource.Attribute{
					Name:        "vpn_domain_type",
					Description: `Gateway VPN domain type. ` + "`" + `logs_settings` + "`" + ` supports the following:`,
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
					Name:        "delete_index_files_when_index_size_above",
					Description: `Enable delete index files when index size above.`,
				},
				resource.Attribute{
					Name:        "delete_index_files_when_index_size_above_threshold",
					Description: `Delete index files when index size above threshold.`,
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
					Description: `Enable detect new Citrix ICA application names.`,
				},
				resource.Attribute{
					Name:        "distribute_logs_between_all_active_servers",
					Description: `Distribute logs between all active servers.`,
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
					Name:        "free_disk_space_metrics",
					Description: `Free disk space metrics.`,
				},
				resource.Attribute{
					Name:        "perform_log_rotate_before_log_forwarding",
					Description: `Enable perform log rotate before log forwarding.`,
				},
				resource.Attribute{
					Name:        "reject_connections_when_free_disk_space_below_threshold",
					Description: `Enable reject connections when free disk space below threshold.`,
				},
				resource.Attribute{
					Name:        "reserve_for_packet_capture_metrics",
					Description: `Reserve for packet capture metrics.`,
				},
				resource.Attribute{
					Name:        "reserve_for_packet_capture_threshold",
					Description: `Reserve for packet capture threshold.`,
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
					Description: `Enable turn on QoS Logging.`,
				},
				resource.Attribute{
					Name:        "update_account_log_every",
					Description: `Update account log in every amount of seconds. ` + "`" + `sam` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "forward_to_other_sam_servers",
					Description: `Forward SAM clients' requests to other SAM servers.`,
				},
				resource.Attribute{
					Name:        "use_early_versions",
					Description: `Use early versions compatibility mode.use_early_versions blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "purge_sam_file",
					Description: `Purge SAM File.purge_sam_file blocks are documented below. ` + "`" + `override_global_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "fail_mode",
					Description: `Fail mode - allow or block all requests.`,
				},
				resource.Attribute{
					Name:        "website_categorization",
					Description: `Website categorization object.website_categorization blocks are documented below. ` + "`" + `bypass_on_failure` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "override_profile",
					Description: `Override profile of global configuration.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Override value.<br><font color="red">Required only for</font> 'override-profile' is True. ` + "`" + `site_categorization_allow_mode` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "override_profile",
					Description: `Override profile of global configuration.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Override value.<br><font color="red">Required only for</font> 'override-profile' is True. ` + "`" + `deny_untrusted_server_cert` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "override_profile",
					Description: `Override profile of global configuration.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Override value.<br><font color="red">Required only for</font> 'override-profile' is True. ` + "`" + `deny_revoked_server_cert` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "override_profile",
					Description: `Override profile of global configuration.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Override value.<br><font color="red">Required only for</font> 'override-profile' is True. ` + "`" + `deny_expired_server_cert` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "override_profile",
					Description: `Override profile of global configuration.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Override value.<br><font color="red">Required only for</font> 'override-profile' is True. ` + "`" + `browser_based_authentication_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "authentication_settings",
					Description: `Authentication Settings for Browser Based Authentication.authentication_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "browser_based_authentication_portal_settings",
					Description: `Browser Based Authentication portal settings.browser_based_authentication_portal_settings blocks are documented below. ` + "`" + `identity_agent_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "agents_interval_keepalive",
					Description: `Agents send keepalive period (minutes).`,
				},
				resource.Attribute{
					Name:        "user_reauthenticate_interval",
					Description: `Agent reauthenticate time interval (minutes).`,
				},
				resource.Attribute{
					Name:        "authentication_settings",
					Description: `Authentication Settings for Identity Agent.authentication_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "identity_agent_portal_settings",
					Description: `Identity Agent accessibility settings.identity_agent_portal_settings blocks are documented below. ` + "`" + `identity_collector_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "authorized_clients",
					Description: `Authorized Clients.authorized_clients blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "authentication_settings",
					Description: `Authentication Settings for Identity Collector.authentication_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "client_access_permissions",
					Description: `Identity Collector accessibility settings.client_access_permissions blocks are documented below. ` + "`" + `identity_sharing_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "share_with_other_gateways",
					Description: `Enable identity sharing with other gateways.`,
				},
				resource.Attribute{
					Name:        "receive_from_other_gateways",
					Description: `Enable receiving identity from other gateways.`,
				},
				resource.Attribute{
					Name:        "receive_from",
					Description: `Gateway(s) to receive identity from.receive_from blocks are documented below. ` + "`" + `proxy_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "detect_using_x_forward_for",
					Description: `Whether to use X-Forward-For HTTP header, which is added by the proxy server to keep track of the original source IP. ` + "`" + `anti_spoofing_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `If packets will be rejected (the Prevent option) or whether the packets will be monitored (the Detect option).`,
				},
				resource.Attribute{
					Name:        "exclude_packets",
					Description: `Don't check packets from excluded network.`,
				},
				resource.Attribute{
					Name:        "excluded_network_name",
					Description: `Excluded network name.`,
				},
				resource.Attribute{
					Name:        "excluded_network_uid",
					Description: `Excluded network UID.`,
				},
				resource.Attribute{
					Name:        "spoof_tracking",
					Description: `Spoof tracking. ` + "`" + `security_zone_settings` + "`" + ` supports the following:`,
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
					Name:        "specific_network",
					Description: `Network behind this interface. ` + "`" + `portal_web_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "aliases",
					Description: `List of URL aliases that are redirected to the main portal URL.aliases blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "main_url",
					Description: `The main URL for the web portal. ` + "`" + `certificate_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "base64_certificate",
					Description: `The certificate file encoded in Base64 with padding. This file must be in the`,
				},
				resource.Attribute{
					Name:        "base64_password",
					Description: `Password (encoded in Base64 with padding) for the certificate file. ` + "`" + `accessibility` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "allow_access_from",
					Description: `Allowed access to the web portal (based on interfaces, or security policy).`,
				},
				resource.Attribute{
					Name:        "internal_access_settings",
					Description: `Configuration of the additional portal access settings for internal interfaces only.internal_access_settings blocks are documented below. ` + "`" + `portal_web_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "aliases",
					Description: `List of URL aliases that are redirected to the main portal URL.aliases blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "main_url",
					Description: `The main URL for the web portal. ` + "`" + `certificate_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "base64_certificate",
					Description: `The certificate file encoded in Base64 with padding. This file must be in the`,
				},
				resource.Attribute{
					Name:        "base64_password",
					Description: `Password (encoded in Base64 with padding) for the certificate file. ` + "`" + `accessibility` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "allow_access_from",
					Description: `Allowed access to the web portal (based on interfaces, or security policy).`,
				},
				resource.Attribute{
					Name:        "internal_access_settings",
					Description: `Configuration of the additional portal access settings for internal interfaces only.internal_access_settings blocks are documented below. ` + "`" + `authentication` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "authentication_clients",
					Description: `Collection of VPN Authentication clients identified by the name or UID.authentication_clients blocks are documented below. ` + "`" + `link_selection` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "dns_resolving_hostname",
					Description: `DNS Resolving Hostname. Must be set when "ip-selection" was selected to be "dns-resolving-from-hostname". ` + "`" + `office_mode` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Office Mode Permissions. When selected to be "off", all the other definitions are irrelevant.`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `Group. Identified by name or UID. Must be set when "office-mode-permissions" was selected to be "group".`,
				},
				resource.Attribute{
					Name:        "allocate_ip_address_from",
					Description: `Allocate IP address Method. Allocate IP address by sequentially trying the given methods until success.allocate_ip_address_from blocks are documented below.`,
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
					Description: `Additional IP Addresses for Anti-Spoofing. Identified by name or UID. Must be set when "perform-anti-spoofings" is true. ` + "`" + `remote_access` + "`" + ` supports the following:`,
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
					Description: `Allocated NAT traversal UDP service. Identified by name or UID. Must be set when "support-nat-traversal-mechanism" is true.`,
				},
				resource.Attribute{
					Name:        "support_visitor_mode",
					Description: `Support Visitor Mode.`,
				},
				resource.Attribute{
					Name:        "visitor_mode_service",
					Description: `TCP Service for Visitor Mode. Identified by name or UID. Must be set when "support-visitor-mode" is true.`,
				},
				resource.Attribute{
					Name:        "visitor_mode_interface",
					Description: `Interface for Visitor Mode. Must be set when "support-visitor-mode" is true. Insert IPV4 Address of existing interface or "All IPs" when you want all interfaces. ` + "`" + `use_early_versions` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Use early versions compatibility mode.`,
				},
				resource.Attribute{
					Name:        "compatibility_mode",
					Description: `Early versions compatibility mode. ` + "`" + `purge_sam_file` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Purge SAM File.`,
				},
				resource.Attribute{
					Name:        "purge_when_size_reaches_to",
					Description: `Purge SAM File When it Reaches to. ` + "`" + `website_categorization` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Website categorization mode.`,
				},
				resource.Attribute{
					Name:        "custom_mode",
					Description: `Custom mode object.custom_mode blocks are documented below. ` + "`" + `authentication_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "authentication_method",
					Description: `Authentication method.`,
				},
				resource.Attribute{
					Name:        "identity_provider",
					Description: `Identity provider object identified by the name or UID. Must be set when "authentication-method" was selected to be "identity provider".identity_provider blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "radius",
					Description: `Radius server object identified by the name or UID. Must be set when "authentication-method" was selected to be "radius".`,
				},
				resource.Attribute{
					Name:        "users_directories",
					Description: `Users directories.users_directories blocks are documented below. ` + "`" + `browser_based_authentication_portal_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "portal_web_settings",
					Description: `Configuration of the portal web settings.portal_web_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "certificate_settings",
					Description: `Configuration of the portal certificate settings.certificate_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "accessibility",
					Description: `Configuration of the portal access settings.accessibility blocks are documented below. ` + "`" + `authentication_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "authentication_method",
					Description: `Authentication method.`,
				},
				resource.Attribute{
					Name:        "radius",
					Description: `Radius server object identified by the name or UID. Must be set when "authentication-method" was selected to be "radius".`,
				},
				resource.Attribute{
					Name:        "users_directories",
					Description: `Users directories.users_directories blocks are documented below. ` + "`" + `identity_agent_portal_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "accessibility",
					Description: `Configuration of the portal access settings.accessibility blocks are documented below. ` + "`" + `authorized_clients` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "client",
					Description: `Host / Network Group Name or UID.`,
				},
				resource.Attribute{
					Name:        "client_secret",
					Description: `Client Secret. ` + "`" + `authentication_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "users_directories",
					Description: `Users directories.users_directories blocks are documented below. ` + "`" + `client_access_permissions` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "accessibility",
					Description: `Configuration of the portal access settings.accessibility blocks are documented below. ` + "`" + `internal_access_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "undefined",
					Description: `Controls portal access settings for internal interfaces, whose topology is set to 'Undefined'.`,
				},
				resource.Attribute{
					Name:        "dmz",
					Description: `Controls portal access settings for internal interfaces, whose topology is set to 'DMZ'.`,
				},
				resource.Attribute{
					Name:        "vpn",
					Description: `Controls portal access settings for interfaces that are part of a VPN Encryption Domain. ` + "`" + `internal_access_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "undefined",
					Description: `Controls portal access settings for internal interfaces, whose topology is set to 'Undefined'.`,
				},
				resource.Attribute{
					Name:        "dmz",
					Description: `Controls portal access settings for internal interfaces, whose topology is set to 'DMZ'.`,
				},
				resource.Attribute{
					Name:        "vpn",
					Description: `Controls portal access settings for interfaces that are part of a VPN Encryption Domain. ` + "`" + `allocate_ip_address_from` + "`" + ` supports the following:`,
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
					Description: `Manual Network. Identified by name or UID. Must be set when "allocate-method" was selected to be "manual".`,
				},
				resource.Attribute{
					Name:        "dhcp_server",
					Description: `DHCP Server. Identified by name or UID. Must be set when "allocate-method" was selected to be "automatic".`,
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
					Description: `This configuration applies to all Office Mode methods except Automatic (using DHCP) and ipassignment.conf entries which contain this data.optional_parameters blocks are documented below. ` + "`" + `custom_mode` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "social_networking_widgets",
					Description: `Social networking widgets mode.`,
				},
				resource.Attribute{
					Name:        "url_filtering",
					Description: `URL filtering mode. ` + "`" + `users_directories` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "external_user_profile",
					Description: `External user profile.`,
				},
				resource.Attribute{
					Name:        "internal_users",
					Description: `Internal users.`,
				},
				resource.Attribute{
					Name:        "users_from_external_directories",
					Description: `Users from external directories.`,
				},
				resource.Attribute{
					Name:        "specific",
					Description: `LDAP AU objects identified by the name or UID. Must be set when "users-from-external-directories" was selected to be "specific".specific blocks are documented below. ` + "`" + `portal_web_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "aliases",
					Description: `List of URL aliases that are redirected to the main portal URL.aliases blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "main_url",
					Description: `The main URL for the web portal. ` + "`" + `certificate_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "base64_certificate",
					Description: `The certificate file encoded in Base64 with padding. This file must be in the`,
				},
				resource.Attribute{
					Name:        "base64_password",
					Description: `Password (encoded in Base64 with padding) for the certificate file. ` + "`" + `accessibility` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "allow_access_from",
					Description: `Allowed access to the web portal (based on interfaces, or security policy).`,
				},
				resource.Attribute{
					Name:        "internal_access_settings",
					Description: `Configuration of the additional portal access settings for internal interfaces only.internal_access_settings blocks are documented below. ` + "`" + `users_directories` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "external_user_profile",
					Description: `External user profile.`,
				},
				resource.Attribute{
					Name:        "internal_users",
					Description: `Internal users.`,
				},
				resource.Attribute{
					Name:        "users_from_external_directories",
					Description: `Users from external directories.`,
				},
				resource.Attribute{
					Name:        "specific",
					Description: `LDAP AU objects identified by the name or UID. Must be set when "users-from-external-directories" was selected to be "specific".specific blocks are documented below. ` + "`" + `accessibility` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "allow_access_from",
					Description: `Allowed access to the web portal (based on interfaces, or security policy).`,
				},
				resource.Attribute{
					Name:        "internal_access_settings",
					Description: `Configuration of the additional portal access settings for internal interfaces only.internal_access_settings blocks are documented below. ` + "`" + `users_directories` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "external_user_profile",
					Description: `External user profile.`,
				},
				resource.Attribute{
					Name:        "internal_users",
					Description: `Internal users.`,
				},
				resource.Attribute{
					Name:        "users_from_external_directories",
					Description: `Users from external directories.`,
				},
				resource.Attribute{
					Name:        "specific",
					Description: `LDAP AU objects identified by the name or UID. Must be set when "users-from-external-directories" was selected to be "specific".specific blocks are documented below. ` + "`" + `accessibility` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "allow_access_from",
					Description: `Allowed access to the web portal (based on interfaces, or security policy).`,
				},
				resource.Attribute{
					Name:        "internal_access_settings",
					Description: `Configuration of the additional portal access settings for internal interfaces only.internal_access_settings blocks are documented below. ` + "`" + `optional_parameters` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "use_primary_dns_server",
					Description: `Use Primary DNS Server.`,
				},
				resource.Attribute{
					Name:        "primary_dns_server",
					Description: `Primary DNS Server. Identified by name or UID. Must be set when "use-primary-dns-server" is true and can not be set when "use-primary-dns-server" is false.`,
				},
				resource.Attribute{
					Name:        "use_first_backup_dns_server",
					Description: `Use First Backup DNS Server.`,
				},
				resource.Attribute{
					Name:        "first_backup_dns_server",
					Description: `First Backup DNS Server. Identified by name or UID. Must be set when "use-first-backup-dns-server" is true and can not be set when "use-first-backup-dns-server" is false.`,
				},
				resource.Attribute{
					Name:        "use_second_backup_dns_server",
					Description: `Use Second Backup DNS Server.`,
				},
				resource.Attribute{
					Name:        "second_backup_dns_server",
					Description: `Second Backup DNS Server. Identified by name or UID. Must be set when "use-second-backup-dns-server" is true and can not be set when "use-second-backup-dns-server" is false.`,
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
					Description: `Primary WINS Server. Identified by name or UID. Must be set when "use-primary-wins-server" is true and can not be set when "use-primary-wins-server" is false.`,
				},
				resource.Attribute{
					Name:        "use_first_backup_wins_server",
					Description: `Use First Backup WINS Server.`,
				},
				resource.Attribute{
					Name:        "first_backup_wins_server",
					Description: `First Backup WINS Server. Identified by name or UID. Must be set when "use-first-backup-wins-server" is true and can not be set when "use-first-backup-wins-server" is false.`,
				},
				resource.Attribute{
					Name:        "use_second_backup_wins_server",
					Description: `Use Second Backup WINS Server.`,
				},
				resource.Attribute{
					Name:        "second_backup_wins_server",
					Description: `Second Backup WINS Server. Identified by name or UID. Must be set when "use-second-backup-wins-server" is true and can not be set when "use-second-backup-wins-server" is false. ` + "`" + `internal_access_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "undefined",
					Description: `Controls portal access settings for internal interfaces, whose topology is set to 'Undefined'.`,
				},
				resource.Attribute{
					Name:        "dmz",
					Description: `Controls portal access settings for internal interfaces, whose topology is set to 'DMZ'.`,
				},
				resource.Attribute{
					Name:        "vpn",
					Description: `Controls portal access settings for interfaces that are part of a VPN Encryption Domain. ` + "`" + `internal_access_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "undefined",
					Description: `Controls portal access settings for internal interfaces, whose topology is set to 'Undefined'.`,
				},
				resource.Attribute{
					Name:        "dmz",
					Description: `Controls portal access settings for internal interfaces, whose topology is set to 'DMZ'.`,
				},
				resource.Attribute{
					Name:        "vpn",
					Description: `Controls portal access settings for interfaces that are part of a VPN Encryption Domain. ` + "`" + `internal_access_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "undefined",
					Description: `Controls portal access settings for internal interfaces, whose topology is set to 'Undefined'.`,
				},
				resource.Attribute{
					Name:        "dmz",
					Description: `Controls portal access settings for internal interfaces, whose topology is set to 'DMZ'.`,
				},
				resource.Attribute{
					Name:        "vpn",
					Description: `Controls portal access settings for interfaces that are part of a VPN Encryption Domain.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_smart_task_trigger",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Smart Task Trigger.`,
			Description: `

Use this data source to get information on an existing Check Point Smart Task Trigger.

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
					Name:        "before_operation",
					Description: `Whether or not this trigger is fired before an operation.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_smtp_server",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Smtp Server.`,
			Description: `

Use this data source to get information on an existing Check Point Smtp Server.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
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
			Type:             "checkpoint_management_tacacs_server",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Tacacs Server.`,
			Description: `

Use this data source to get information on an existing Check Point Tacacs Server.

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
					Name:        "encryption",
					Description: `Is there a secret key defined on the server. Must be set true when "server-type" was selected to be "TACACS+".`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority of the TACACS Server in case it is a member of a TACACS Group.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `The UID or Name of the host that is the TACACS Server. server blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "server_type",
					Description: `Server type, TACACS or TACACS+.`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `Server service, only relevant when "server-type" is TACACS. service blocks are documented below. ` + "`" + `server` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Object name. Must be unique in the domain.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `Object unique identifier. ` + "`" + `service` + "`" + ` supports the following:`,
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
					Name:        "aggressive_aging",
					Description: `Sets short (aggressive) timeouts for idle connections. aggressive_aging blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.`,
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
					Description: `The number of the port used to provide this service.`,
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
					Description: `Use default virtual session timeout. ` + "`" + `aggressive_aging` + "`" + ` supports the following:`,
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
			Type:             "checkpoint_management_task",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Task.`,
			Description: `

Use this data source to get information on an existing Check Point Task.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "task_id",
					Description: `(Required) Unique identifier of one or more tasks.`,
				},
				resource.Attribute{
					Name:        "tasks",
					Description: `The tasks. tasks blocks are documented below. ` + "`" + `tasks` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "task_id",
					Description: `Asynchronous task unique identifier. Use show-task command to check the progress of the task.`,
				},
				resource.Attribute{
					Name:        "task_name",
					Description: `The task name.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Task status.`,
				},
				resource.Attribute{
					Name:        "progress_percentage",
					Description: `The progress percentage of the task.`,
				},
				resource.Attribute{
					Name:        "suppressed",
					Description: `Is the task suppressed.`,
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
			Type:             "checkpoint_management_threat_advanced_settings",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Threat Advanced Settings.`,
			Description: `

Use this data source to get information on an existing Check Point Threat Advanced Settings.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "uid",
					Description: `Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "feed_retrieving_interval",
					Description: `Feed retrieving intervals of External Feed, in the form of HH:MM.`,
				},
				resource.Attribute{
					Name:        "httpi_non_standard_ports",
					Description: `Enable HTTP Inspection on non standard ports for Threat Prevention blades.`,
				},
				resource.Attribute{
					Name:        "internal_error_fail_mode",
					Description: `In case of internal system error, allow or block all connections.`,
				},
				resource.Attribute{
					Name:        "log_unification_timeout",
					Description: `Session unification timeout for logs (minutes).`,
				},
				resource.Attribute{
					Name:        "resource_classification",
					Description: `Allow (Background) or Block (Hold) requests until categorization is complete. resource_classification blocks are documented below. ` + "`" + `resource_classification` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "custom_settings",
					Description: `Custom resources classification per service. custom_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Set all services to the same mode or choose a custom mode.`,
				},
				resource.Attribute{
					Name:        "web_service_fail_mode",
					Description: `Block connections when the web service is unavailable. ` + "`" + `custom_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "anti_bot",
					Description: `Custom Settings for Anti Bot Blade.`,
				},
				resource.Attribute{
					Name:        "anti_virus",
					Description: `Custom Settings for Anti Virus Blade.`,
				},
				resource.Attribute{
					Name:        "zero_phishing",
					Description: `Custom Settings for Zero Phishing Blade.`,
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
			Type:             "checkpoint_management_threat_ioc_feed",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Threat Ioc Feed.`,
			Description: `

Use this data source to get information on an existing Check Point Threat Ioc Feed.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
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
			Type:             "checkpoint_management_threat_layer",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Threat Layer.`,
			Description: `

Use this data source to get information on an existing Check Point Threat Layer.

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
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_threat_profile",
			Category:         "Data Sources",
			ShortDescription: `This resource allows you to execute Check Point Threat Profile.`,
			Description: `

This resource allows you to execute Check Point Threat Profile.

`,
			Keywords: []string{},
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
					Description: `Add customized text to the malicious email body.`,
				},
				resource.Attribute{
					Name:        "add_email_subject_prefix",
					Description: `Add a prefix to the malicious email subject.`,
				},
				resource.Attribute{
					Name:        "add_x_header_to_email",
					Description: `Add an X-Header to the malicious email.`,
				},
				resource.Attribute{
					Name:        "email_action",
					Description: `Block - block the entire malicious email. Allow - pass the malicious email and apply email changes (like: remove attachments and links, add x-header, etc...).`,
				},
				resource.Attribute{
					Name:        "email_body_customized_text",
					Description: `Customized text for the malicious email body. Available predefined fields: $verdicts$ - the malicious/error attachments/links verdict.`,
				},
				resource.Attribute{
					Name:        "email_subject_prefix_text",
					Description: `Prefix for the malicious email subject.`,
				},
				resource.Attribute{
					Name:        "failed_to_scan_attachments_text",
					Description: `Replace attachments that failed to be scanned with this text. Available predefined fields: $filename$ - the malicious file name. $md5$ - MD5 of the malicious file.`,
				},
				resource.Attribute{
					Name:        "malicious_attachments_text",
					Description: `Replace malicious attachments with this text. Available predefined fields: $filename$ - the malicious file name. $md5$ - MD5 of the malicious file.`,
				},
				resource.Attribute{
					Name:        "malicious_links_text",
					Description: `Replace malicious links with this text. Available predefined fields: $neutralized_url$ - neutralized malicious link.`,
				},
				resource.Attribute{
					Name:        "remove_attachments_and_links",
					Description: `Remove attachments and links from the malicious email.`,
				},
				resource.Attribute{
					Name:        "send_copy",
					Description: `Send a copy of the malicious email to the recipient list.`,
				},
				resource.Attribute{
					Name:        "send_copy_list",
					Description: `Recipient list to send a copy of the malicious email. ` + "`" + `overrides` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "protection",
					Description: `IPS protection identified by name.`,
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
					Description: `Scan links in the first bytes of the mail body.`,
				},
				resource.Attribute{
					Name:        "max_links",
					Description: `Maximum links to scan in mail body. ` + "`" + `activate_protections_by_extended_attributes` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `IPS tag unique identifier.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `IPS tag name.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `IPS tag category name.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `Collection of IPS protection extended attribute values (name and uid). ` + "`" + `deactivate_protections_by_extended_attributes` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `IPS tag unique identifier.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `IPS tag name.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `IPS tag category name.`,
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
			Type:             "checkpoint_management_threat_rulebase",
			Category:         "Data Sources",
			ShortDescription: `This resource allows you to execute Check Point Threat Rule Base.`,
			Description: `

Use this data source to get information on an existing threat RuleBase.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Object name. Must be unique in the domain.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Object unique identifier.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `Search expression to filter the rulebase. The provided text should be exactly the same as it would be given in Smart Console. The logical operators in the expression ('AND', 'OR') should be provided in capital letters. If an operator is not used, the default OR operator applies.`,
				},
				resource.Attribute{
					Name:        "filter_settings",
					Description: `Enable enforce end user domain. filter_settings blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "limit",
					Description: `The maximal number of returned results.`,
				},
				resource.Attribute{
					Name:        "offset",
					Description: `Number of the results to initially skip.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `Sorts the results by search criteria. Automatically sorts the results by Name, in the ascending order. orders blocks are documented below.`,
				},
				resource.Attribute{
					Name:        "package",
					Description: `Name of the package.`,
				},
				resource.Attribute{
					Name:        "dereference_group_members",
					Description: `Indicates whether to dereference "members" field by details level for every object in reply.`,
				},
				resource.Attribute{
					Name:        "show_membership",
					Description: `Indicates whether to calculate and show "groups" field for every object in reply. ` + "`" + `filter_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "search_mode",
					Description: `When set to 'general', both the Full Text Search and Packet Search are enabled. In this mode, Packet Search will not match on 'Any' object, a negated cell or a group-with-exclusion. When the search-mode is set to 'packet', by default, the match on 'Any' object, a negated cell or a group-with-exclusion are enabled. packet-search-settings may be provided to change the default behavior.`,
				},
				resource.Attribute{
					Name:        "expand_group_members",
					Description: `(Optional, can only be used when search_mode is set to "packet") When true, if the search expression contains a UID or a name of a group object, results will include rules that match on at least one member of the group.`,
				},
				resource.Attribute{
					Name:        "expand_group_with_exclusion_members",
					Description: `(Optional, can only be used when search_mode is set to "packet") When true, if the search expression contains a UID or a name of a group-with-exclusion object, results will include rules that match at least one member of the "include" part and is not a member of the "except" part.`,
				},
				resource.Attribute{
					Name:        "match_on_any",
					Description: `(Optional, can only be used when search_mode is set to "packet") Whether to match on 'Any' object.`,
				},
				resource.Attribute{
					Name:        "match_on_group_with_exclusion",
					Description: `(Optional, can only be used when search_mode is set to "packet") Whether to match on a group-with-exclusion.`,
				},
				resource.Attribute{
					Name:        "match_on_negate",
					Description: `(Optional, can only be used when search_mode is set to "packet") Whether to match on a negated cell. ` + "`" + `order` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "asc",
					Description: `(Optional) Sorts results by the given field in ascending order.`,
				},
				resource.Attribute{
					Name:        "desc",
					Description: `(Optional) Sorts results by the given field in descending order.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "checkpoint_management_time",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Time.`,
			Description: `

Use this data source to get information on an existing Check Point Time.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
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
			Type:             "checkpoint_management_trusted_client",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing Check Point Trusted Client.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
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
			Type:             "checkpoint_management_vmware_data_center_server",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing VMware Data Center Server.`,
			Description: `

Use this data source to get information on an existing VMware Data Center Server.

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
		"checkpoint_management_access_rulebase":                           1,
		"checkpoint_management_aci_data_center_server":                    2,
		"checkpoint_management_administrator":                             3,
		"checkpoint_management_api_settings":                              4,
		"checkpoint_management_automatic_purge":                           5,
		"checkpoint_management_aws_data_center_server":                    6,
		"checkpoint_management_azure_ad":                                  7,
		"checkpoint_management_azure_ad_content":                          8,
		"checkpoint_management_azure_data_center_server":                  9,
		"checkpoint_management_checkpoint_host":                           10,
		"checkpoint_management_cloud_services":                            11,
		"checkpoint_management_cluster_member":                            12,
		"checkpoint_management_data_access_layer":                         13,
		"checkpoint_management_data_access_role":                          14,
		"checkpoint_management_data_access_rule":                          15,
		"checkpoint_management_data_access_section":                       16,
		"checkpoint_management_data_address_range":                        17,
		"checkpoint_management_data_application_site":                     18,
		"checkpoint_management_data_application_site_category":            19,
		"checkpoint_management_data_application_site_group":               20,
		"checkpoint_management_data_center_content":                       21,
		"checkpoint_management_data_center_query":                         22,
		"checkpoint_management_data_dns_domain":                           23,
		"checkpoint_management_data_dynamic_object":                       24,
		"checkpoint_management_data_exception_group":                      25,
		"checkpoint_management_data_group":                                26,
		"checkpoint_management_data_group_with_exclusion":                 27,
		"checkpoint_management_data_host":                                 28,
		"checkpoint_management_data_https_layer":                          29,
		"checkpoint_management_data_https_rule":                           30,
		"checkpoint_management_data_https_section":                        31,
		"checkpoint_management_data_multicast_address_range":              32,
		"checkpoint_management_data_network":                              33,
		"checkpoint_management_data_opsec_application":                    34,
		"checkpoint_management_data_package":                              35,
		"checkpoint_management_data_security_zone":                        36,
		"checkpoint_management_data_service_dce_rpc":                      37,
		"checkpoint_management_data_service_group":                        38,
		"checkpoint_management_data_service_icmp":                         39,
		"checkpoint_management_data_service_icmp6":                        40,
		"checkpoint_management_data_service_other":                        41,
		"checkpoint_management_data_service_rpc":                          42,
		"checkpoint_management_data_service_sctp":                         43,
		"checkpoint_management_data_service_tcp":                          44,
		"checkpoint_management_data_service_udp":                          45,
		"checkpoint_management_data_threat_indicator":                     46,
		"checkpoint_management_data_time_group":                           47,
		"checkpoint_management_data_vpn_community_meshed":                 48,
		"checkpoint_management_data_vpn_community_star":                   49,
		"checkpoint_management_data_wildcard":                             50,
		"checkpoint_management_domain":                                    51,
		"checkpoint_management_domain_permissions_profile":                52,
		"checkpoint_management_dynamic_global_network_object":             53,
		"checkpoint_management_gaia_best_practice":                        54,
		"checkpoint_management_gcp_data_center_server":                    55,
		"checkpoint_management_generic_data_center_server":                56,
		"checkpoint_management_global_assignment":                         57,
		"checkpoint_management_global_domain":                             58,
		"checkpoint_management_gsn_handover_group":                        59,
		"checkpoint_management_https_rulebase":                            60,
		"checkpoint_management_identity_tag":                              61,
		"checkpoint_management_idp_administrator_group":                   62,
		"checkpoint_management_idp_default_assignment":                    63,
		"checkpoint_management_idp_to_domain_assignment":                  64,
		"checkpoint_management_interoperable_device":                      65,
		"checkpoint_management_ips_update_schedule":                       66,
		"checkpoint_management_ise_data_center_server":                    67,
		"checkpoint_management_kubernetes_data_center_server":             68,
		"checkpoint_management_login_message":                             69,
		"checkpoint_management_lsm_cluster_profile":                       70,
		"checkpoint_management_lsv_profile":                               71,
		"checkpoint_management_md_permissions_profile":                    72,
		"checkpoint_management_mds":                                       73,
		"checkpoint_management_nat_rule":                                  74,
		"checkpoint_management_nat_rulebase":                              75,
		"checkpoint_management_nat_section":                               76,
		"checkpoint_management_network_feed":                              77,
		"checkpoint_management_nuage_data_center_server":                  78,
		"checkpoint_management_nutanix_data_center_server":                79,
		"checkpoint_management_objects":                                   80,
		"checkpoint_management_openstack_data_center_server":              81,
		"checkpoint_management_oracle_cloud_data_center_server":           82,
		"checkpoint_management_policy_settings":                           83,
		"checkpoint_management_provisioning_profile":                      84,
		"checkpoint_management_radius_group":                              85,
		"checkpoint_management_radius_server":                             86,
		"checkpoint_management_repository_package":                        87,
		"checkpoint_management_repository_script":                         88,
		"checkpoint_management_service_citrix_tcp":                        89,
		"checkpoint_management_service_compound_tcp":                      90,
		"checkpoint_management_show_objects":                              91,
		"checkpoint_management_show_updatable_objects_repository_content": 92,
		"checkpoint_management_simple_cluster":                            93,
		"checkpoint_management_simple_gateway":                            94,
		"checkpoint_management_smart_task_trigger":                        95,
		"checkpoint_management_smtp_server":                               96,
		"checkpoint_management_tacacs_server":                             97,
		"checkpoint_management_task":                                      98,
		"checkpoint_management_threat_advanced_settings":                  99,
		"checkpoint_management_threat_exception":                          100,
		"checkpoint_management_threat_ioc_feed":                           101,
		"checkpoint_management_threat_layer":                              102,
		"checkpoint_management_threat_profile":                            103,
		"checkpoint_management_threat_rule":                               104,
		"checkpoint_management_threat_rulebase":                           105,
		"checkpoint_management_time":                                      106,
		"checkpoint_management_trusted_client":                            107,
		"checkpoint_management_user":                                      108,
		"checkpoint_management_user_group":                                109,
		"checkpoint_management_user_template":                             110,
		"checkpoint_management_vmware_data_center_server":                 111,
		"checkpoint_management_vpn_community_remote_access":               112,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
