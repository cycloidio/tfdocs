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
					Description: `(Required) The teams that the apikey belongs to.`,
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
					Description: `(Optional) The feeds configuration matching the specification in 'feed\_config' from /data/sourcetypes.`,
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
					Description: `(Optional) The data source configuration, determined by its type.`,
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
					Description: `(Required) The type of monitoring job to be run.`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `(Required) Indicates if the job is active or temporaril.y disabled.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `(Required) The list of region codes in which to run the monitoring job.`,
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
					Description: `(Required) The policy for determining the monitor's global status based on the status of the job in all regions.`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `(Required) A configuration dictionary with keys and values depending on the jobs' type.`,
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
					Description: `(Optional) The id of the notification list to send notifications to.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) Freeform notes to be included in any notifications about this job.`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `(Optional) A list of rules for determining failure conditions. Job Rules are documented below. Monitoring Job Rules (` + "`" + `rules` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The output key.`,
				},
				resource.Attribute{
					Name:        "comparison",
					Description: `(Required) The comparison to perform on the the output.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value to compare to.`,
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
					Description: `(Required) Configuration details for the given notifier type.`,
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
					Description: `(Required) The zone the record belongs to.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) The records' domain.`,
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
					Name:        "answers",
					Description: `(Optional) One or more NS1 answers for the records' specified type. Answers are documented below.`,
				},
				resource.Attribute{
					Name:        "filters",
					Description: `(Optional) One or more NS1 filters for the record(order matters). Filters are documented below. Answers (` + "`" + `answers` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "answer",
					Description: `(Required) Space delimited string of RDATA fields dependent on the record type. A: answer = "1.2.3.4" CNAME: answer = "www.example.com" MX: answer = "5 mail.example.com" SRV: answer = "10 0 2380 node-1.example.com" SPF: answer = "v=DKIM1; k=rsa; p=XXXXXXXX"`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region (or group) name that this answer belongs to. Regions must be sorted alphanumberically by name, otherwise Terraform will detect changes to the record when none actually exist. Filters (` + "`" + `filters` + "`" + `) support the following:`,
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
					Description: `(Optional) The filters' configuration. Simple key/value pairs determined by the filter type.`,
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
					Description: `(Required) The Whether or not to notify the user of specified events. Only ` + "`" + `billing` + "`" + ` is available currently.`,
				},
				resource.Attribute{
					Name:        "teams",
					Description: `(Required) The teams that the user belongs to.`,
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
					Description: `(Optional) Whether the user can modify the account plan.`,
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
					Name:        "ttl",
					Description: `(Optional) The SOA TTL.`,
				},
				resource.Attribute{
					Name:        "refresh",
					Description: `(Optional) The SOA Refresh.`,
				},
				resource.Attribute{
					Name:        "retry",
					Description: `(Optional) The SOA Retry.`,
				},
				resource.Attribute{
					Name:        "expiry",
					Description: `(Optional) The SOA Expiry.`,
				},
				resource.Attribute{
					Name:        "nx_ttl",
					Description: `(Optional) The SOA NX TTL.`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `(Optional) The primary zones' ip. This makes the zone a secondary.`,
				},
			},
			Attributes: []resource.Attribute{},
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
