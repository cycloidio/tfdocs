package dnsimple

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "dnsimple_record",
			Category:         "Resources",
			ShortDescription: `Provides a DNSimple record resource.`,
			Description:      ``,
			Keywords: []string{
				"record",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) The domain to add the record to`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the record`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of the record`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of the record`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The TTL of the record`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) The priority of the record - only useful for some record types ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The record ID`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the record`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the record`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the record`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `The TTL of the record`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority of the record`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `The domain ID of the record`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `The FQDN of the record ## Import DNSimple resources can be imported using their parent zone name (domain name) and numeric record ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The record ID`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the record`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the record`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the record`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `The TTL of the record`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority of the record`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `The domain ID of the record`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `The FQDN of the record ## Import DNSimple resources can be imported using their parent zone name (domain name) and numeric record ID.`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"dnsimple_record": 0,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
