package ns1

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "ns1_dnssec",
			Category:         "Data Sources",
			ShortDescription: `Provides DNSSEC details about a NS1 Zone.`,
			Description: `

Provides DNSSEC details about a NS1 Zone.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) The name of the zone to get DNSSEC details for. ## Attributes Reference In addition to the argument above, the following are exported:`,
				},
				resource.Attribute{
					Name:        "keys",
					Description: `(Computed) - [Keys](#keys-1) field is documented below.`,
				},
				resource.Attribute{
					Name:        "delegation",
					Description: `(Computed) - [Delegation](#delegation-1) field is documented below. #### Keys ` + "`" + `keys` + "`" + ` has the following fields:`,
				},
				resource.Attribute{
					Name:        "dnskey",
					Description: `(Computed) List of Keys. [Key](#key) is documented below.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Computed) TTL for the Keys (int). #### Delegation ` + "`" + `delegation` + "`" + ` has the following fields:`,
				},
				resource.Attribute{
					Name:        "dnskey",
					Description: `(Computed) List of Keys. [Key](#key) is documented below.`,
				},
				resource.Attribute{
					Name:        "ds",
					Description: `(Computed) List of Keys. [Key](#key) is documented below.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Computed) TTL for the Keys (int). #### Key A ` + "`" + `key` + "`" + ` has the following (string) fields:`,
				},
				resource.Attribute{
					Name:        "flags",
					Description: `(Computed) Flags for the key.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Computed) Protocol of the key.`,
				},
				resource.Attribute{
					Name:        "algorithm",
					Description: `(Computed) Algorithm of the key.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(Computed) Public key for the key.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "keys",
					Description: `(Computed) - [Keys](#keys-1) field is documented below.`,
				},
				resource.Attribute{
					Name:        "delegation",
					Description: `(Computed) - [Delegation](#delegation-1) field is documented below. #### Keys ` + "`" + `keys` + "`" + ` has the following fields:`,
				},
				resource.Attribute{
					Name:        "dnskey",
					Description: `(Computed) List of Keys. [Key](#key) is documented below.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Computed) TTL for the Keys (int). #### Delegation ` + "`" + `delegation` + "`" + ` has the following fields:`,
				},
				resource.Attribute{
					Name:        "dnskey",
					Description: `(Computed) List of Keys. [Key](#key) is documented below.`,
				},
				resource.Attribute{
					Name:        "ds",
					Description: `(Computed) List of Keys. [Key](#key) is documented below.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Computed) TTL for the Keys (int). #### Key A ` + "`" + `key` + "`" + ` has the following (string) fields:`,
				},
				resource.Attribute{
					Name:        "flags",
					Description: `(Computed) Flags for the key.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Computed) Protocol of the key.`,
				},
				resource.Attribute{
					Name:        "algorithm",
					Description: `(Computed) Algorithm of the key.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(Computed) Public key for the key.`,
				},
			},
		},
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
					Description: `The primary zones' IPv4 address.`,
				},
				resource.Attribute{
					Name:        "additional_primaries",
					Description: `List of additional IPv4 addresses for the primary zone.`,
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
					Name:        "dnssec",
					Description: `Whether or not DNSSEC is enabled for the zone.`,
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
				resource.Attribute{
					Name:        "secondaries",
					Description: `List of secondary servers. [Secondaries](#secondaries-1) is documented below. #### Secondaries A secondary has the following fields:`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `IPv4 address of the secondary server.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port of the the secondary server. Default ` + "`" + `53` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "notify",
					Description: `Whether we send ` + "`" + `NOTIFY` + "`" + ` messages to the secondary host when the zone changes. Default ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "networks",
					Description: `List of network IDs (` + "`" + `int` + "`" + `) for which the zone should be made available. Default is network 0, the primary NSONE Global Network.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "link",
					Description: `The linked target zone.`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `The primary zones' IPv4 address.`,
				},
				resource.Attribute{
					Name:        "additional_primaries",
					Description: `List of additional IPv4 addresses for the primary zone.`,
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
					Name:        "dnssec",
					Description: `Whether or not DNSSEC is enabled for the zone.`,
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
				resource.Attribute{
					Name:        "secondaries",
					Description: `List of secondary servers. [Secondaries](#secondaries-1) is documented below. #### Secondaries A secondary has the following fields:`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `IPv4 address of the secondary server.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port of the the secondary server. Default ` + "`" + `53` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "notify",
					Description: `Whether we send ` + "`" + `NOTIFY` + "`" + ` messages to the secondary host when the zone changes. Default ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "networks",
					Description: `List of network IDs (` + "`" + `int` + "`" + `) for which the zone should be made available. Default is network 0, the primary NSONE Global Network.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"ns1_dnssec": 0,
		"ns1_zone":   1,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
