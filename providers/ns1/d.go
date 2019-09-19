package ns1

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "ns1_zone",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a NS1 Zone.`,
			Description: `

Provides details about a NS1 Zone. Use this if you would simply like to read
information from NS1 into your configurations. For read/write operations, you
should use a resource.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) The domain name of the zone. ## Attributes Reference In addition to the argument above, the following are exported:`,
				},
				resource.Attribute{
					Name:        "link",
					Description: `The linked target zone.`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `The primary ip.`,
				},
				resource.Attribute{
					Name:        "additional_primaries",
					Description: `List of additional IPs for the primary zone.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `The SOA TTL.`,
				},
				resource.Attribute{
					Name:        "refresh",
					Description: `The SOA Refresh.`,
				},
				resource.Attribute{
					Name:        "retry",
					Description: `The SOA Retry.`,
				},
				resource.Attribute{
					Name:        "expiry",
					Description: `The SOA Expiry.`,
				},
				resource.Attribute{
					Name:        "nx_ttl",
					Description: `The SOA NX TTL.`,
				},
				resource.Attribute{
					Name:        "networks",
					Description: `List of network IDs for which the zone is available.`,
				},
				resource.Attribute{
					Name:        "dns_servers",
					Description: `Authoritative Name Servers.`,
				},
				resource.Attribute{
					Name:        "hostmaster",
					Description: `The SOA Hostmaster.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "link",
					Description: `The linked target zone.`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `The primary ip.`,
				},
				resource.Attribute{
					Name:        "additional_primaries",
					Description: `List of additional IPs for the primary zone.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `The SOA TTL.`,
				},
				resource.Attribute{
					Name:        "refresh",
					Description: `The SOA Refresh.`,
				},
				resource.Attribute{
					Name:        "retry",
					Description: `The SOA Retry.`,
				},
				resource.Attribute{
					Name:        "expiry",
					Description: `The SOA Expiry.`,
				},
				resource.Attribute{
					Name:        "nx_ttl",
					Description: `The SOA NX TTL.`,
				},
				resource.Attribute{
					Name:        "networks",
					Description: `List of network IDs for which the zone is available.`,
				},
				resource.Attribute{
					Name:        "dns_servers",
					Description: `Authoritative Name Servers.`,
				},
				resource.Attribute{
					Name:        "hostmaster",
					Description: `The SOA Hostmaster.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"ns1_zone": 0,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
