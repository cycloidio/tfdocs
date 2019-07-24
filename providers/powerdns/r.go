package aws

import (
	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*Resource{

		&resource.Resource{
			Name:             "",
			Type:             "powerdns_record",
			Category:         "Resources",
			ShortDescription: `Provides a PowerDNS record resource.`,
			Description:      ``,
			Keywords: []string{
				"record",
			},
			Arguments: []resource.Argument{
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
			Attributes: []resource.Argument{},
		},
	}

	resourcesMap = map[string]int{

		"powerdns_record": 0,
	}
)

func GetResource(r string) (*resouce.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs]
}
