package ns1

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "ns1_apikey",
			Category:         "Resources",
			ShortDescription: `Provides a NS1 Api Key resource.`,
			Description:      ``,
			Keywords: []string{
				"apikey",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The free form name of the apikey.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The apikeys authentication token.`,
				},
				resource.Attribute{
					Name:        "teams",
					Description: `(Optional) The teams that the apikey belongs to.`,
				},
				resource.Attribute{
					Name:        "ip_whitelist",
					Description: `(Optional) The IP addresses to whitelist for this key.`,
				},
				resource.Attribute{
					Name:        "ip_whitelist_strict",
					Description: `(Optional) Sets exclusivity on this IP whitelist.`,
				},
				resource.Attribute{
					Name:        "dns_view_zones",
					Description: `(Optional) Whether the apikey can view the accounts zones.`,
				},
				resource.Attribute{
					Name:        "dns_manage_zones",
					Description: `(Optional) Whether the apikey can modify the accounts zones.`,
				},
				resource.Attribute{
					Name:        "dns_zones_allow_by_default",
					Description: `(Optional) If true, enable the ` + "`" + `dns_zones_allow` + "`" + ` list, otherwise enable the ` + "`" + `dns_zones_deny` + "`" + ` list.`,
				},
				resource.Attribute{
					Name:        "dns_zones_allow",
					Description: `(Optional) List of zones that the apikey may access.`,
				},
				resource.Attribute{
					Name:        "dns_zones_deny",
					Description: `(Optional) List of zones that the apikey may not access.`,
				},
				resource.Attribute{
					Name:        "data_push_to_datafeeds",
					Description: `(Optional) Whether the apikey can publish to data feeds.`,
				},
				resource.Attribute{
					Name:        "data_manage_datasources",
					Description: `(Optional) Whether the apikey can modify data sources.`,
				},
				resource.Attribute{
					Name:        "data_manage_datafeeds",
					Description: `(Optional) Whether the apikey can modify data feeds.`,
				},
				resource.Attribute{
					Name:        "account_manage_users",
					Description: `(Optional) Whether the apikey can modify account users.`,
				},
				resource.Attribute{
					Name:        "account_manage_payment_methods",
					Description: `(Optional) Whether the apikey can modify account payment methods.`,
				},
				resource.Attribute{
					Name:        "account_manage_plan",
					Description: `(Optional) Whether the apikey can modify the account plan.`,
				},
				resource.Attribute{
					Name:        "account_manage_teams",
					Description: `(Optional) Whether the apikey can modify other teams in the account.`,
				},
				resource.Attribute{
					Name:        "account_manage_apikeys",
					Description: `(Optional) Whether the apikey can modify account apikeys.`,
				},
				resource.Attribute{
					Name:        "account_manage_account_settings",
					Description: `(Optional) Whether the apikey can modify account settings.`,
				},
				resource.Attribute{
					Name:        "account_view_activity_log",
					Description: `(Optional) Whether the apikey can view activity logs.`,
				},
				resource.Attribute{
					Name:        "account_view_invoices",
					Description: `(Optional) Whether the apikey can view invoices.`,
				},
				resource.Attribute{
					Name:        "monitoring_manage_lists",
					Description: `(Optional) Whether the apikey can modify notification lists.`,
				},
				resource.Attribute{
					Name:        "monitoring_manage_jobs",
					Description: `(Optional) Whether the apikey can modify monitoring jobs.`,
				},
				resource.Attribute{
					Name:        "monitoring_view_jobs",
					Description: `(Optional) Whether the apikey can view monitoring jobs.`,
				},
				resource.Attribute{
					Name:        "security_manage_global_2fa",
					Description: `(Optional) Whether the apikey can manage global two factor authentication.`,
				},
				resource.Attribute{
					Name:        "security_manage_active_directory",
					Description: `(Optional) Whether the apikey can manage global active directory. Only relevant for the DDI product.`,
				},
				resource.Attribute{
					Name:        "dhcp_manage_dhcp",
					Description: `(Optional) Whether the apikey can manage DHCP. Only relevant for the DDI product.`,
				},
				resource.Attribute{
					Name:        "dhcp_view_dhcp",
					Description: `(Optional) Whether the apikey can view DHCP. Only relevant for the DDI product.`,
				},
				resource.Attribute{
					Name:        "ipam_manage_ipam",
					Description: `(Optional) Whether the apikey can manage IPAM. Only relevant for the DDI product.`,
				},
				resource.Attribute{
					Name:        "ipam_view_ipam",
					Description: `(Optional) Whether the apikey can view IPAM. Only relevant for the DDI product. ## Attributes Reference All of the arguments listed above are exported as attributes, with no additions. ## NS1 Documentation [ApiKeys Api Doc](https://ns1.com/api#api-key)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ns1_datafeed",
			Category:         "Resources",
			ShortDescription: `Provides a NS1 Data Feed resource.`,
			Description:      ``,
			Keywords: []string{
				"datafeed",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "source_id",
					Description: `(Required) The data source id that this feed is connected to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The free form name of the data feed.`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `(Optional) The feeds configuration matching the specification in ` + "`" + `feed_config` + "`" + ` from /data/sourcetypes. ## Attributes Reference All of the arguments listed above are exported as attributes, with no additions. ## NS1 Documentation [Datafeed Api Doc](https://ns1.com/api#data-feeds)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ns1_datasource",
			Category:         "Resources",
			ShortDescription: `Provides a NS1 Data Source resource.`,
			Description:      ``,
			Keywords: []string{
				"datasource",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The free form name of the data source.`,
				},
				resource.Attribute{
					Name:        "sourcetype",
					Description: `(Required) The data sources type, listed in API endpoint https://api.nsone.net/v1/data/sourcetypes.`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `(Optional) The data source configuration, determined by its type, matching the specification in ` + "`" + `config` + "`" + ` from /data/sourcetypes. ## Attributes Reference All of the arguments listed above are exported as attributes, with no additions. ## NS1 Documentation [Datasource Api Doc](https://ns1.com/api#data-sources)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ns1_monitoringjob",
			Category:         "Resources",
			ShortDescription: `Provides a NS1 Monitoring Job resource.`,
			Description:      ``,
			Keywords: []string{
				"monitoringjob",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The free-form display name for the monitoring job.`,
				},
				resource.Attribute{
					Name:        "job_type",
					Description: `(Required) The type of monitoring job to be run. Refer to the NS1 API documentation (https://ns1.com/api#monitoring-jobs) for supported values which include ping, tcp, dns, http.`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `(Required) Indicates if the job is active or temporarily disabled.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `(Required) The list of region codes in which to run the monitoring job. See NS1 API docs for supported values.`,
				},
				resource.Attribute{
					Name:        "frequency",
					Description: `(Required) The frequency, in seconds, at which to run the monitoring job in each region.`,
				},
				resource.Attribute{
					Name:        "rapid_recheck",
					Description: `(Required) If true, on any apparent state change, the job is quickly re-run after one second to confirm the state change before notification.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Required) The policy for determining the monitor's global status based on the status of the job in all regions. See NS1 API docs for supported values.`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `(Required) A configuration dictionary with keys and values depending on the job_type. Configuration details for each job_type are found by submitting a GET request to https://api.nsone.net/v1/monitoring/jobtypes.`,
				},
				resource.Attribute{
					Name:        "notify_delay",
					Description: `(Optional) The time in seconds after a failure to wait before sending a notification.`,
				},
				resource.Attribute{
					Name:        "notify_repeat",
					Description: `(Optional) The time in seconds between repeat notifications of a failed job.`,
				},
				resource.Attribute{
					Name:        "notify_failback",
					Description: `(Optional) If true, a notification is sent when a job returns to an "up" state.`,
				},
				resource.Attribute{
					Name:        "notify_regional",
					Description: `(Optional) If true, notifications are sent for any regional failure (and failback if desired), in addition to global state notifications.`,
				},
				resource.Attribute{
					Name:        "notify_list",
					Description: `(Optional) The Terraform ID (e.g. ns1_notifylist.my_slack_notifier.id) of the notification list to which monitoring notifications should be sent.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) Freeform notes to be included in any notifications about this job.`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `(Optional) A list of rules for determining failure conditions. Each rule acts on one of the outputs from the monitoring job. You must specify key (the output key); comparison (a comparison to perform on the the output); and value (the value to compare to). For example, {"key":"rtt", "comparison":"<", "value":100} is a rule requiring the rtt from a job to be under 100ms, or the job will be marked failed. Available output keys, comparators, and value types are are found by submitting a GET request to https://api.nsone.net/v1/monitoring/jobtypes. ## Attributes Reference All of the arguments listed above are exported as attributes, with no additions. ## NS1 Documentation [MonitoringJob Api Doc](https://ns1.com/api#monitoring-jobs)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ns1_notifylist",
			Category:         "Resources",
			ShortDescription: `Provides a NS1 Notify List resource.`,
			Description:      ``,
			Keywords: []string{
				"notifylist",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The free-form display name for the notify list.`,
				},
				resource.Attribute{
					Name:        "notifications",
					Description: `(Optional) A list of notifiers. All notifiers in a notification list will receive notifications whenever an event is send to the list (e.g., when a monitoring job fails). Notifiers are documented below. Notify List Notifiers (` + "`" + `notifications` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of notifier. Available notifiers are indicated in /notifytypes endpoint.`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `(Required) Configuration details for the given notifier type. ## Attributes Reference All of the arguments listed above are exported as attributes, with no additions. ## NS1 Documentation [NotifyList Api Doc](https://ns1.com/api#notification-lists)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ns1_record",
			Category:         "Resources",
			ShortDescription: `Provides a NS1 Record resource.`,
			Description:      ``,
			Keywords: []string{
				"record",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) The zone the record belongs to. Cannot have leading or trailing dots (".") - see the example above and ` + "`" + `FQDN formatting` + "`" + ` below.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) The records' domain. Cannot have leading or trailing dots - see the example above and ` + "`" + `FQDN formatting` + "`" + ` below.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The records' RR type.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The records' time to live.`,
				},
				resource.Attribute{
					Name:        "link",
					Description: `(Optional) The target record to link to. This means this record is a 'linked' record, and it inherits all properties from its target.`,
				},
				resource.Attribute{
					Name:        "use_client_subnet",
					Description: `(Optional) Whether to use EDNS client subnet data when available(in filter chain).`,
				},
				resource.Attribute{
					Name:        "meta",
					Description: `(Optional) meta is supported at the ` + "`" + `record` + "`" + ` level. [Meta](#meta-3) is documented below.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `(Optional) One or more "regions" for the record. These are really just groupings based on metadata, and are called "Answer Groups" in the NS1 UI, but remain ` + "`" + `regions` + "`" + ` here for legacy reasons. [Regions](#regions-1) are documented below. Please note the ordering requirement!`,
				},
				resource.Attribute{
					Name:        "answers",
					Description: `(Optional) One or more NS1 answers for the records' specified type. [Answers](#answers-1) are documented below.`,
				},
				resource.Attribute{
					Name:        "filters",
					Description: `(Optional) One or more NS1 filters for the record(order matters). [Filters](#filters-1) are documented below. #### Answers ` + "`" + `answers` + "`" + ` support the following:`,
				},
				resource.Attribute{
					Name:        "answer",
					Description: `(Required) Space delimited string of RDATA fields dependent on the record type. A: answer = "1.2.3.4" CNAME: answer = "www.example.com" MX: answer = "5 mail.example.com" SRV: answer = "10 0 2380 node-1.example.com" SPF: answer = "v=DKIM1; k=rsa; p=XXXXXXXX"`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region (Answer Group really) that this answer belongs to. This should be one of the names specified in ` + "`" + `regions` + "`" + `. Only a single ` + "`" + `region` + "`" + ` per answer is currently supported. If you want an answer in multiple regions, duplicating the answer (including metadata) is the correct approach.`,
				},
				resource.Attribute{
					Name:        "meta",
					Description: `(Optional) meta is supported at the ` + "`" + `answer` + "`" + ` level. [Meta](#meta-3) is documented below. #### Filters ` + "`" + `filters` + "`" + ` support the following:`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Required) The type of filter.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional) Determines whether the filter is applied in the filter chain.`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `(Optional) The filters' configuration. Simple key/value pairs determined by the filter type. #### Regions ` + "`" + `regions` + "`" + ` support the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the region (or Answer Group).`,
				},
				resource.Attribute{
					Name:        "meta",
					Description: `(Optional) meta is supported at the ` + "`" + `regions` + "`" + ` level. [Meta](#meta-3) is documented below. Note that ` + "`" + `Meta` + "`" + ` values for ` + "`" + `country` + "`" + `, ` + "`" + `ca_province` + "`" + `, ` + "`" + `georegion` + "`" + `, and ` + "`" + `us_state` + "`" + ` should be comma separated strings, and changes in ordering will not lead to terraform detecting a change. Note: regions`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ns1_team",
			Category:         "Resources",
			ShortDescription: `Provides a NS1 Team resource.`,
			Description:      ``,
			Keywords: []string{
				"team",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The free form name of the team.`,
				},
				resource.Attribute{
					Name:        "ip_whitelist",
					Description: `(Optional) The IP addresses to whitelist for this key.`,
				},
				resource.Attribute{
					Name:        "dns_view_zones",
					Description: `(Optional) Whether the team can view the accounts zones.`,
				},
				resource.Attribute{
					Name:        "dns_manage_zones",
					Description: `(Optional) Whether the team can modify the accounts zones.`,
				},
				resource.Attribute{
					Name:        "dns_zones_allow_by_default",
					Description: `(Optional) If true, enable the ` + "`" + `dns_zones_allow` + "`" + ` list, otherwise enable the ` + "`" + `dns_zones_deny` + "`" + ` list.`,
				},
				resource.Attribute{
					Name:        "dns_zones_allow",
					Description: `(Optional) List of zones that the team may access.`,
				},
				resource.Attribute{
					Name:        "dns_zones_deny",
					Description: `(Optional) List of zones that the team may not access.`,
				},
				resource.Attribute{
					Name:        "data_push_to_datafeeds",
					Description: `(Optional) Whether the team can publish to data feeds.`,
				},
				resource.Attribute{
					Name:        "data_manage_datasources",
					Description: `(Optional) Whether the team can modify data sources.`,
				},
				resource.Attribute{
					Name:        "data_manage_datafeeds",
					Description: `(Optional) Whether the team can modify data feeds.`,
				},
				resource.Attribute{
					Name:        "account_manage_users",
					Description: `(Optional) Whether the team can modify account users.`,
				},
				resource.Attribute{
					Name:        "account_manage_payment_methods",
					Description: `(Optional) Whether the team can modify account payment methods.`,
				},
				resource.Attribute{
					Name:        "account_manage_plan",
					Description: `(Optional) Whether the team can modify the account plan.`,
				},
				resource.Attribute{
					Name:        "account_manage_teams",
					Description: `(Optional) Whether the team can modify other teams in the account.`,
				},
				resource.Attribute{
					Name:        "account_manage_apikeys",
					Description: `(Optional) Whether the team can modify account apikeys.`,
				},
				resource.Attribute{
					Name:        "account_manage_account_settings",
					Description: `(Optional) Whether the team can modify account settings.`,
				},
				resource.Attribute{
					Name:        "account_view_activity_log",
					Description: `(Optional) Whether the team can view activity logs.`,
				},
				resource.Attribute{
					Name:        "account_view_invoices",
					Description: `(Optional) Whether the team can view invoices.`,
				},
				resource.Attribute{
					Name:        "monitoring_manage_lists",
					Description: `(Optional) Whether the team can modify notification lists.`,
				},
				resource.Attribute{
					Name:        "monitoring_manage_jobs",
					Description: `(Optional) Whether the team can modify monitoring jobs.`,
				},
				resource.Attribute{
					Name:        "monitoring_view_jobs",
					Description: `(Optional) Whether the team can view monitoring jobs.`,
				},
				resource.Attribute{
					Name:        "security_manage_global_2fa",
					Description: `(Optional) Whether the team can manage global two factor authentication.`,
				},
				resource.Attribute{
					Name:        "security_manage_active_directory",
					Description: `(Optional) Whether the team can manage global active directory. Only relevant for the DDI product.`,
				},
				resource.Attribute{
					Name:        "dhcp_manage_dhcp",
					Description: `(Optional) Whether the team can manage DHCP. Only relevant for the DDI product.`,
				},
				resource.Attribute{
					Name:        "dhcp_view_dhcp",
					Description: `(Optional) Whether the team can view DHCP. Only relevant for the DDI product.`,
				},
				resource.Attribute{
					Name:        "ipam_manage_ipam",
					Description: `(Optional) Whether the team can manage IPAM. Only relevant for the DDI product.`,
				},
				resource.Attribute{
					Name:        "ipam_view_ipam",
					Description: `(Optional) Whether the team can view IPAM. Only relevant for the DDI product. ## Attributes Reference All of the arguments listed above are exported as attributes, with no additions. ## NS1 Documentation [Team Api Docs](https://ns1.com/api#team)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ns1_user",
			Category:         "Resources",
			ShortDescription: `Provides a NS1 User resource.`,
			Description:      ``,
			Keywords: []string{
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The free form name of the user.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The users login name.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Required) The email address of the user.`,
				},
				resource.Attribute{
					Name:        "notify",
					Description: `(Required) Whether or not to notify the user of specified events. Only ` + "`" + `billing` + "`" + ` is available currently.`,
				},
				resource.Attribute{
					Name:        "teams",
					Description: `(Required) The teams that the user belongs to.`,
				},
				resource.Attribute{
					Name:        "ip_whitelist",
					Description: `(Optional) The IP addresses to whitelist for this key.`,
				},
				resource.Attribute{
					Name:        "ip_whitelist_strict",
					Description: `(Optional) Sets exclusivity on this IP whitelist.`,
				},
				resource.Attribute{
					Name:        "dns_view_zones",
					Description: `(Optional) Whether the user can view the accounts zones.`,
				},
				resource.Attribute{
					Name:        "dns_manage_zones",
					Description: `(Optional) Whether the user can modify the accounts zones.`,
				},
				resource.Attribute{
					Name:        "dns_zones_allow_by_default",
					Description: `(Optional) If true, enable the ` + "`" + `dns_zones_allow` + "`" + ` list, otherwise enable the ` + "`" + `dns_zones_deny` + "`" + ` list.`,
				},
				resource.Attribute{
					Name:        "dns_zones_allow",
					Description: `(Optional) List of zones that the user may access.`,
				},
				resource.Attribute{
					Name:        "dns_zones_deny",
					Description: `(Optional) List of zones that the user may not access.`,
				},
				resource.Attribute{
					Name:        "data_push_to_datafeeds",
					Description: `(Optional) Whether the user can publish to data feeds.`,
				},
				resource.Attribute{
					Name:        "data_manage_datasources",
					Description: `(Optional) Whether the user can modify data sources.`,
				},
				resource.Attribute{
					Name:        "data_manage_datafeeds",
					Description: `(Optional) Whether the user can modify data feeds.`,
				},
				resource.Attribute{
					Name:        "account_manage_users",
					Description: `(Optional) Whether the user can modify account users.`,
				},
				resource.Attribute{
					Name:        "account_manage_payment_methods",
					Description: `(Optional) Whether the user can modify account payment methods.`,
				},
				resource.Attribute{
					Name:        "account_manage_plan",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "account_manage_teams",
					Description: `(Optional) Whether the user can modify other teams in the account.`,
				},
				resource.Attribute{
					Name:        "account_manage_apikeys",
					Description: `(Optional) Whether the user can modify account apikeys.`,
				},
				resource.Attribute{
					Name:        "account_manage_account_settings",
					Description: `(Optional) Whether the user can modify account settings.`,
				},
				resource.Attribute{
					Name:        "account_view_activity_log",
					Description: `(Optional) Whether the user can view activity logs.`,
				},
				resource.Attribute{
					Name:        "account_view_invoices",
					Description: `(Optional) Whether the user can view invoices.`,
				},
				resource.Attribute{
					Name:        "monitoring_manage_lists",
					Description: `(Optional) Whether the user can modify notification lists.`,
				},
				resource.Attribute{
					Name:        "monitoring_manage_jobs",
					Description: `(Optional) Whether the user can modify monitoring jobs.`,
				},
				resource.Attribute{
					Name:        "monitoring_view_jobs",
					Description: `(Optional) Whether the user can view monitoring jobs.`,
				},
				resource.Attribute{
					Name:        "security_manage_global_2fa",
					Description: `(Optional) Whether the user can manage global two factor authentication.`,
				},
				resource.Attribute{
					Name:        "security_manage_active_directory",
					Description: `(Optional) Whether the user can manage global active directory. Only relevant for the DDI product.`,
				},
				resource.Attribute{
					Name:        "dhcp_manage_dhcp",
					Description: `(Optional) Whether the user can manage DHCP. Only relevant for the DDI product.`,
				},
				resource.Attribute{
					Name:        "dhcp_view_dhcp",
					Description: `(Optional) Whether the user can view DHCP. Only relevant for the DDI product.`,
				},
				resource.Attribute{
					Name:        "ipam_manage_ipam",
					Description: `(Optional) Whether the user can manage IPAM. Only relevant for the DDI product. ## Attributes Reference All of the arguments listed above are exported as attributes, with no additions. ## NS1 Documentation [User Api Docs](https://ns1.com/api#user) [Managing user permissions](https://help.ns1.com/hc/en-us/articles/360024409034-Managing-user-permissions)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ns1_zone",
			Category:         "Resources",
			ShortDescription: `Provides a NS1 Zone resource.`,
			Description:      ``,
			Keywords: []string{
				"zone",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) The domain name of the zone.`,
				},
				resource.Attribute{
					Name:        "link",
					Description: `(Optional) The target zone(domain name) to link to.`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `(Optional) The primary zones' IPv4 address. This makes the zone a secondary. Conflicts with ` + "`" + `secondaries` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "additional_primaries",
					Description: `(Optional) List of additional IPv4 addresses for the primary zone. Conflicts with ` + "`" + `secondaries` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional/Computed) The SOA TTL.`,
				},
				resource.Attribute{
					Name:        "refresh",
					Description: `(Optional/Computed) The SOA Refresh. Conflicts with ` + "`" + `primary` + "`" + ` and ` + "`" + `additional_primaries` + "`" + ` (default must be accepted).`,
				},
				resource.Attribute{
					Name:        "retry",
					Description: `(Optional/Computed) The SOA Retry. Conflicts with ` + "`" + `primary` + "`" + ` and ` + "`" + `additional_primaries` + "`" + ` (default must be accepted).`,
				},
				resource.Attribute{
					Name:        "expiry",
					Description: `(Optional/Computed) The SOA Expiry. Conflicts with ` + "`" + `primary` + "`" + ` and ` + "`" + `additional_primaries` + "`" + ` (default must be accepted).`,
				},
				resource.Attribute{
					Name:        "nx_ttl",
					Description: `(Optional/Computed) The SOA NX TTL. Conflicts with ` + "`" + `primary` + "`" + ` and ` + "`" + `additional_primaries` + "`" + ` (default must be accepted).`,
				},
				resource.Attribute{
					Name:        "dnssec",
					Description: `(Optional/Computed) Whether or not DNSSEC is enabled for the zone. Note that DNSSEC must be enabled on the account by support for this to be set to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "networks",
					Description: `(Optional/Computed) List of network IDs for which the zone is available. If no network is provided, the zone will be created in network 0, the primary NS1 Global Network.`,
				},
				resource.Attribute{
					Name:        "secondaries",
					Description: `(Optional) List of secondary servers. This makes the zone a primary. Conflicts with ` + "`" + `primary` + "`" + ` and ` + "`" + `additional_primaries` + "`" + `. [Secondaries](#secondaries-1) is documented below.`,
				},
				resource.Attribute{
					Name:        "autogenerate_ns_record",
					Description: `(Optional, default true). If set to false, clears the autogenerated NS record on zone creation. This allows an automated workflow for creating zones with the NS record in terraform state. See above for an example. Note that this option only has an effect when a zone is being created. #### Secondaries A zone can have zero or more ` + "`" + `secondaries` + "`" + `. Note how this is implemented in the example above. A secondary has the following fields:`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) IPv4 address of the secondary server.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Port of the the secondary server. Default ` + "`" + `53` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "notify",
					Description: `(Optional) Whether we send ` + "`" + `NOTIFY` + "`" + ` messages to the secondary host when the zone changes. Default ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "networks",
					Description: `(Computed) - List of network IDs (` + "`" + `int` + "`" + `) for which the zone should be made available. Default is network 0, the primary NSONE Global Network. Normally, you should not have to worry about this. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "dns_servers",
					Description: `(Computed) Authoritative Name Servers.`,
				},
				resource.Attribute{
					Name:        "hostmaster",
					Description: `(Computed) The SOA Hostmaster. ## A note on making Primary or Secondary changes to zones Switching a zone to being a secondary forces a new resource. In other words, the zone will first be destroyed, then recreated as a secondary. Editing or removing the ` + "`" + `primary` + "`" + ` key, or directly changing a secondary zone to a primary (by removing the ` + "`" + `primary` + "`" + ` and ` + "`" + `additional_primaries` + "`" + ` keys, and setting ` + "`" + `secondaries` + "`" + `) is supported "in place". However, in these situations we do not alter records on the zone. You may need to amend records, or finagle them into Terraform state. As a particular example, if you change a secondary zone to be primary (or just not-a-secondary) before a zone transfer has occurred, you can end up with no records on the zone. Currently, this provider does not support zones being both Primary and Secondary. If that functionality is important for your workflow, please open an issue or contact support, so we can prioritize the work accordingly. ## Import ` + "`" + `terraform import ns1_zone.<name> <zone>` + "`" + ` So for the example above: ` + "`" + `terraform import ns1_zone.example terraform.example.io` + "`" + ` ## NS1 Documentation [Zone Api Docs](https://ns1.com/api#zones)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "dns_servers",
					Description: `(Computed) Authoritative Name Servers.`,
				},
				resource.Attribute{
					Name:        "hostmaster",
					Description: `(Computed) The SOA Hostmaster. ## A note on making Primary or Secondary changes to zones Switching a zone to being a secondary forces a new resource. In other words, the zone will first be destroyed, then recreated as a secondary. Editing or removing the ` + "`" + `primary` + "`" + ` key, or directly changing a secondary zone to a primary (by removing the ` + "`" + `primary` + "`" + ` and ` + "`" + `additional_primaries` + "`" + ` keys, and setting ` + "`" + `secondaries` + "`" + `) is supported "in place". However, in these situations we do not alter records on the zone. You may need to amend records, or finagle them into Terraform state. As a particular example, if you change a secondary zone to be primary (or just not-a-secondary) before a zone transfer has occurred, you can end up with no records on the zone. Currently, this provider does not support zones being both Primary and Secondary. If that functionality is important for your workflow, please open an issue or contact support, so we can prioritize the work accordingly. ## Import ` + "`" + `terraform import ns1_zone.<name> <zone>` + "`" + ` So for the example above: ` + "`" + `terraform import ns1_zone.example terraform.example.io` + "`" + ` ## NS1 Documentation [Zone Api Docs](https://ns1.com/api#zones)`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"ns1_apikey":        0,
		"ns1_datafeed":      1,
		"ns1_datasource":    2,
		"ns1_monitoringjob": 3,
		"ns1_notifylist":    4,
		"ns1_record":        5,
		"ns1_team":          6,
		"ns1_user":          7,
		"ns1_zone":          8,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
