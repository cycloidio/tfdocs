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
					Description: `(Required) A string list of records. ### Attribute Reference The id of the resource is a composite of the record name and record type, joined by a separator - ` + "`" + `:::` + "`" + `. For example, record ` + "`" + `foo.test.com.` + "`" + ` of type ` + "`" + `A` + "`" + ` will be represented with the following ` + "`" + `id` + "`" + `: ` + "`" + `foo.test.com.:::A` + "`" + ` ### Importing An existing record can be imported into this resource by supplying both the record id and zone name it belongs to. If the record or zone is not found, or if the record is of a different type or in a different zone, an error will be returned. For example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import powerdns_record.test-a '{"zone": "test.com.", "id": "foo.test.com.:::A"}' ` + "`" + `` + "`" + `` + "`" + ` For more information on how to use terraform's ` + "`" + `import` + "`" + ` command, please refer to terraform's [core documentation](https://www.terraform.io/docs/import/index.html#currently-state-only).`,
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
					Description: `(Required) The zone nameservers. ## Importing An existing zone can be imported into this resource by supplying the zone name. If the zone is not found, an error will be returned. For example, to import zone ` + "`" + `test.com.` + "`" + `: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import powerdns_zone.test test.com. ` + "`" + `` + "`" + `` + "`" + ` For more information on how to use terraform's ` + "`" + `import` + "`" + ` command, please refer to terraform's [core documentation](https://www.terraform.io/docs/import/index.html#currently-state-only).`,
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
