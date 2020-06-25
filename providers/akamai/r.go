package akamai

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "cp_code",
			Category:         "Properties",
			ShortDescription: `CP Code`,
			Description:      ``,
			Keywords: []string{
				"properties",
				"cp",
				"code",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_dns_record",
			Category:         "Edge DNS",
			ShortDescription: `DNS Record`,
			Description:      ``,
			Keywords: []string{
				"edge",
				"dns",
				"record",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_dns_zone",
			Category:         "Edge DNS",
			ShortDescription: `DNS Zone`,
			Description:      ``,
			Keywords: []string{
				"edge",
				"dns",
				"zone",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `key name`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_edge_hostname",
			Category:         "Properties",
			ShortDescription: `Edge Hostname`,
			Description:      ``,
			Keywords: []string{
				"properties",
				"edge",
				"hostname",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_gtm_asmap",
			Category:         "Traffic Management",
			ShortDescription: `GTM AS Map`,
			Description:      ``,
			Keywords: []string{
				"traffic",
				"management",
				"gtm",
				"asmap",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_gtm_cidrmap",
			Category:         "Traffic Management",
			ShortDescription: `GTM Cidr Map`,
			Description:      ``,
			Keywords: []string{
				"traffic",
				"management",
				"gtm",
				"cidrmap",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_gtm_datacenter",
			Category:         "Traffic Management",
			ShortDescription: `GTM Datacenter`,
			Description:      ``,
			Keywords: []string{
				"traffic",
				"management",
				"gtm",
				"datacenter",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_gtm_domain",
			Category:         "Traffic Management",
			ShortDescription: `GTM Domain`,
			Description:      ``,
			Keywords: []string{
				"traffic",
				"management",
				"gtm",
				"domain",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_gtm_geomap",
			Category:         "Traffic Management",
			ShortDescription: `GTM Geographic Map`,
			Description:      ``,
			Keywords: []string{
				"traffic",
				"management",
				"gtm",
				"geomap",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_gtm_property",
			Category:         "Traffic Management",
			ShortDescription: `GTM Property`,
			Description:      ``,
			Keywords: []string{
				"traffic",
				"management",
				"gtm",
				"property",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_gtm_resource",
			Category:         "Traffic Management",
			ShortDescription: `GTM Resource`,
			Description:      ``,
			Keywords: []string{
				"traffic",
				"management",
				"gtm",
				"resource",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_property",
			Category:         "Properties",
			ShortDescription: `Create and update Akamai Properties`,
			Description:      ``,
			Keywords: []string{
				"properties",
				"property",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_property_activation",
			Category:         "Properties",
			ShortDescription: `Property Activation`,
			Description:      ``,
			Keywords: []string{
				"properties",
				"property",
				"activation",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_property_variables",
			Category:         "Properties",
			ShortDescription: `Property Variables`,
			Description:      ``,
			Keywords: []string{
				"properties",
				"property",
				"variables",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"cp_code":                    0,
		"akamai_dns_record":          1,
		"akamai_dns_zone":            2,
		"akamai_edge_hostname":       3,
		"akamai_gtm_asmap":           4,
		"akamai_gtm_cidrmap":         5,
		"akamai_gtm_datacenter":      6,
		"akamai_gtm_domain":          7,
		"akamai_gtm_geomap":          8,
		"akamai_gtm_property":        9,
		"akamai_gtm_resource":        10,
		"akamai_property":            11,
		"akamai_property_activation": 12,
		"akamai_property_variables":  13,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
