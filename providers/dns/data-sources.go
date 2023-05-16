package dns

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "dns_a_record_set",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get DNS A records of the host.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dns_aaaa_record_set",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get DNS AAAA records of the host.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dns_cname_record_set",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get DNS CNAME record set of the host.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dns_mx_record_set",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get DNS MX records for a domain.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dns_ns_record_set",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get DNS NS records of the host.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dns_ptr_record_set",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get DNS PTR record set of the ip address.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dns_srv_record_set",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get DNS SRV records for a service.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dns_txt_record_set",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get DNS TXT record set of the host.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"dns_a_record_set":     0,
		"dns_aaaa_record_set":  1,
		"dns_cname_record_set": 2,
		"dns_mx_record_set":    3,
		"dns_ns_record_set":    4,
		"dns_ptr_record_set":   5,
		"dns_srv_record_set":   6,
		"dns_txt_record_set":   7,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
