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
	}

	dataSourcesMap = map[string]int{

		"datadog_ip_ranges": 0,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
