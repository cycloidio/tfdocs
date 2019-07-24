package aws

import (
	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*Resource{

		&resource.Resource{
			Name:             "",
			Type:             "dyn_record",
			Category:         "Resources",
			ShortDescription: `Provides a Dyn DNS record resource.`,
			Description:      ``,
			Keywords: []string{
				"record",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the record.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of the record.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of the record.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) The DNS zone to add the record to.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The TTL of the record. Default uses the zone default. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The record ID.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `The FQDN of the record, built from the ` + "`" + `name` + "`" + ` and the ` + "`" + `zone` + "`" + `. ## Import Dyn records can be imported using a combination of the ` + "`" + `type` + "`" + `, ` + "`" + `zone` + "`" + `, ` + "`" + `fdqn` + "`" + `, and optionally ` + "`" + `id` + "`" + `. ` + "`" + `` + "`" + `` + "`" + ` $terraform import dyn_record.record {type}/{zone}/{fqdn}[/{id}] ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The record ID.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `The FQDN of the record, built from the ` + "`" + `name` + "`" + ` and the ` + "`" + `zone` + "`" + `. ## Import Dyn records can be imported using a combination of the ` + "`" + `type` + "`" + `, ` + "`" + `zone` + "`" + `, ` + "`" + `fdqn` + "`" + `, and optionally ` + "`" + `id` + "`" + `. ` + "`" + `` + "`" + `` + "`" + ` $terraform import dyn_record.record {type}/{zone}/{fqdn}[/{id}] ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"dyn_record": 0,
	}
)

func GetResource(r string) (*resouce.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs]
}
