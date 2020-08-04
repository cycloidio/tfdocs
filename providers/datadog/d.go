package datadog

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "datadog_ip_ranges",
			Category:         "Data Sources",
			ShortDescription: `Get information on Datadog's IP addresses.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "agents_ipv4",
					Description: `An Array of IPv4 addresses in CIDR format specifying the A records for the agent endpoint.`,
				},
				resource.Attribute{
					Name:        "api_ipv4",
					Description: `An Array of IPv4 addresses in CIDR format specifying the A records for the api endpoint.`,
				},
				resource.Attribute{
					Name:        "apm_ipv4",
					Description: `An Array of IPv4 addresses in CIDR format specifying the A records for the apm endpoint.`,
				},
				resource.Attribute{
					Name:        "logs_ipv4",
					Description: `An Array of IPv4 addresses in CIDR format specifying the A records for the logs endpoint.`,
				},
				resource.Attribute{
					Name:        "process_ipv4",
					Description: `An Array of IPv4 addresses in CIDR format specifying the A records for the process endpoint.`,
				},
				resource.Attribute{
					Name:        "synthetics_ipv4",
					Description: `An Array of IPv4 addresses in CIDR format specifying the A records for the synthetics endpoint.`,
				},
				resource.Attribute{
					Name:        "webhooks_ipv4",
					Description: `An Array of IPv4 addresses in CIDR format specifying the A records for the webhooks endpoint.`,
				},
				resource.Attribute{
					Name:        "agents_ipv6",
					Description: `An Array of IPv6 addresses in CIDR format specifying the A records for the agent endpoint.`,
				},
				resource.Attribute{
					Name:        "api_ipv6",
					Description: `An Array of IPv6 addresses in CIDR format specifying the A records for the api endpoint.`,
				},
				resource.Attribute{
					Name:        "apm_ipv6",
					Description: `An Array of IPv6 addresses in CIDR format specifying the A records for the apm endpoint.`,
				},
				resource.Attribute{
					Name:        "logs_ipv6",
					Description: `An Array of IPv6 addresses in CIDR format specifying the A records for the logs endpoint.`,
				},
				resource.Attribute{
					Name:        "process_ipv6",
					Description: `An Array of IPv6 addresses in CIDR format specifying the A records for the process endpoint.`,
				},
				resource.Attribute{
					Name:        "synthetics_ipv6",
					Description: `An Array of IPv6 addresses in CIDR format specifying the A records for the synthetics endpoint.`,
				},
				resource.Attribute{
					Name:        "webhooks_ipv6",
					Description: `An Array of IPv6 addresses in CIDR format specifying the A records for the webhooks endpoint.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_monitor",
			Category:         "Data Sources",
			ShortDescription: `Retrieve information about an existing monitor.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_filter",
					Description: `(Optional) A monitor name to limit the search.`,
				},
				resource.Attribute{
					Name:        "tags_filter",
					Description: `(Optional) A list of tags to limit the search. This filters on the monitor scope.`,
				},
				resource.Attribute{
					Name:        "monitor_tags_filter",
					Description: `(Optional) A list of monitor tags to limit the search. This filters on the tags set on the monitor itself. ~>`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Datadog monitor`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the monitor.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the monitor.`,
				},
				resource.Attribute{
					Name:        "query",
					Description: `Query of the monitor.`,
				},
				resource.Attribute{
					Name:        "message",
					Description: `Message included with notifications for this monitor.`,
				},
				resource.Attribute{
					Name:        "escalation_message",
					Description: `Message included with a re-notification for this monitor.`,
				},
				resource.Attribute{
					Name:        "thresholds",
					Description: `Alert thresholds of the monitor.`,
				},
				resource.Attribute{
					Name:        "notify_no_data",
					Description: `Whether or not this monitor notifies when data stops reporting.`,
				},
				resource.Attribute{
					Name:        "new_host_delay",
					Description: `Time (in seconds) allowing a host to boot and applications to fully start before starting the evaluation of monitor results.`,
				},
				resource.Attribute{
					Name:        "evaluation_delay",
					Description: `Time (in seconds) for which evaluation is delayed. This is only used by metric monitors.`,
				},
				resource.Attribute{
					Name:        "no_data_timeframe",
					Description: `The number of minutes before the monitor notifies when data stops reporting.`,
				},
				resource.Attribute{
					Name:        "renotify_interval",
					Description: `The number of minutes after the last notification before the monitor re-notifies on the current status.`,
				},
				resource.Attribute{
					Name:        "notify_audit",
					Description: `Whether or not tagged users are notified on changes to the monitor.`,
				},
				resource.Attribute{
					Name:        "timeout_h",
					Description: `Number of hours of the monitor not reporting data before it automatically resolves from a triggered state.`,
				},
				resource.Attribute{
					Name:        "include_tags",
					Description: `Whether or not notifications from the monitor automatically inserts its triggering tags into the title.`,
				},
				resource.Attribute{
					Name:        "enable_logs_sample",
					Description: `Whether or not a list of log values which triggered the alert is included. This is only used by log monitors.`,
				},
				resource.Attribute{
					Name:        "require_full_window",
					Description: `Whether or not the monitor needs a full window of data before it is evaluated.`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `Whether or not changes to the monitor are restricted to the creator or admins.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `List of tags associated with the monitor.`,
				},
				resource.Attribute{
					Name:        "threshold_windows",
					Description: `Mapping containing ` + "`" + `recovery_window` + "`" + ` and ` + "`" + `trigger_window` + "`" + ` values, e.g. ` + "`" + `last_15m` + "`" + `. This is only used by anomaly monitors.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Datadog monitor`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the monitor.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the monitor.`,
				},
				resource.Attribute{
					Name:        "query",
					Description: `Query of the monitor.`,
				},
				resource.Attribute{
					Name:        "message",
					Description: `Message included with notifications for this monitor.`,
				},
				resource.Attribute{
					Name:        "escalation_message",
					Description: `Message included with a re-notification for this monitor.`,
				},
				resource.Attribute{
					Name:        "thresholds",
					Description: `Alert thresholds of the monitor.`,
				},
				resource.Attribute{
					Name:        "notify_no_data",
					Description: `Whether or not this monitor notifies when data stops reporting.`,
				},
				resource.Attribute{
					Name:        "new_host_delay",
					Description: `Time (in seconds) allowing a host to boot and applications to fully start before starting the evaluation of monitor results.`,
				},
				resource.Attribute{
					Name:        "evaluation_delay",
					Description: `Time (in seconds) for which evaluation is delayed. This is only used by metric monitors.`,
				},
				resource.Attribute{
					Name:        "no_data_timeframe",
					Description: `The number of minutes before the monitor notifies when data stops reporting.`,
				},
				resource.Attribute{
					Name:        "renotify_interval",
					Description: `The number of minutes after the last notification before the monitor re-notifies on the current status.`,
				},
				resource.Attribute{
					Name:        "notify_audit",
					Description: `Whether or not tagged users are notified on changes to the monitor.`,
				},
				resource.Attribute{
					Name:        "timeout_h",
					Description: `Number of hours of the monitor not reporting data before it automatically resolves from a triggered state.`,
				},
				resource.Attribute{
					Name:        "include_tags",
					Description: `Whether or not notifications from the monitor automatically inserts its triggering tags into the title.`,
				},
				resource.Attribute{
					Name:        "enable_logs_sample",
					Description: `Whether or not a list of log values which triggered the alert is included. This is only used by log monitors.`,
				},
				resource.Attribute{
					Name:        "require_full_window",
					Description: `Whether or not the monitor needs a full window of data before it is evaluated.`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `Whether or not changes to the monitor are restricted to the creator or admins.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `List of tags associated with the monitor.`,
				},
				resource.Attribute{
					Name:        "threshold_windows",
					Description: `Mapping containing ` + "`" + `recovery_window` + "`" + ` and ` + "`" + `trigger_window` + "`" + ` values, e.g. ` + "`" + `last_15m` + "`" + `. This is only used by anomaly monitors.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"datadog_ip_ranges": 0,
		"datadog_monitor":   1,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
