package powerdns

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "powerdns_record",
			Category:         "Resources",
			ShortDescription: `Provides a PowerDNS record resource.`,
			Description:      ``,
			Keywords: []string{
				"record",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) The name of zone to contain this record.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the record.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The record type.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Required) The TTL of the record.`,
				},
				resource.Attribute{
					Name:        "records",
					Description: `(Required) A string list of records.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "powerdns_zone",
			Category:         "Resources",
			ShortDescription: `Provides a PowerDNS zone.`,
			Description:      ``,
			Keywords: []string{
				"zone",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of zone.`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `(Required) The kind of the zone.`,
				},
				resource.Attribute{
					Name:        "nameservers",
					Description: `(Required) The zone nameservers.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"powerdns_record": 0,
		"powerdns_zone":   1,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
