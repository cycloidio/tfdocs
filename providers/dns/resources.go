package dns

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "dns_a_record_set",
			Category:         "Resources",
			ShortDescription: `Creates an A type DNS record set.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dns_aaaa_record_set",
			Category:         "Resources",
			ShortDescription: `Creates an AAAA type DNS record set.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dns_cname_record",
			Category:         "Resources",
			ShortDescription: `Creates a CNAME type DNS record.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dns_mx_record_set",
			Category:         "Resources",
			ShortDescription: `Creates an MX type DNS record set.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dns_ns_record_set",
			Category:         "Resources",
			ShortDescription: `Creates an NS type DNS record set.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dns_ptr_record",
			Category:         "Resources",
			ShortDescription: `Creates a PTR type DNS record.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dns_srv_record_set",
			Category:         "Resources",
			ShortDescription: `Creates an SRV type DNS record set.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dns_txt_record_set",
			Category:         "Resources",
			ShortDescription: `Creates a TXT type DNS record set.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"dns_a_record_set":    0,
		"dns_aaaa_record_set": 1,
		"dns_cname_record":    2,
		"dns_mx_record_set":   3,
		"dns_ns_record_set":   4,
		"dns_ptr_record":      5,
		"dns_srv_record_set":  6,
		"dns_txt_record_set":  7,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
