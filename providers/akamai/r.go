package akamai

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "akamai_dns_record",
			Category:         "Resources",
			ShortDescription: `DNS Record`,
			Description:      ``,
			Keywords: []string{
				"dns",
				"record",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_dns_zone",
			Category:         "Resources",
			ShortDescription: `DNS Zone`,
			Description:      ``,
			Keywords: []string{
				"dns",
				"zone",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_edge_hostname",
			Category:         "Resources",
			ShortDescription: `Edge Hostname`,
			Description:      ``,
			Keywords: []string{
				"edge",
				"hostname",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_property",
			Category:         "Resources",
			ShortDescription: `Create and update Akamai Properties`,
			Description:      ``,
			Keywords: []string{
				"property",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_property_activation",
			Category:         "Resources",
			ShortDescription: `Property Activation`,
			Description:      ``,
			Keywords: []string{
				"property",
				"activation",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_property_variables",
			Category:         "Resources",
			ShortDescription: `Property Variables`,
			Description:      ``,
			Keywords: []string{
				"property",
				"variables",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"akamai_dns_record":          0,
		"akamai_dns_zone":            1,
		"akamai_edge_hostname":       2,
		"akamai_property":            3,
		"akamai_property_activation": 4,
		"akamai_property_variables":  5,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
