package dns

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "dns_dns_a_record_set",
			Category:         "Data Sources",
			ShortDescription: `Get DNS A record set.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "host",
					Description: `(required): Host to look up ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Set to ` + "`" + `host` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "addrs",
					Description: `A list of IP addresses. IP addresses are always sorted to avoid constant changing plans.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Set to ` + "`" + `host` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "addrs",
					Description: `A list of IP addresses. IP addresses are always sorted to avoid constant changing plans.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dns_dns_aaaa_record_set",
			Category:         "Data Sources",
			ShortDescription: `Get DNS AAAA record set.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "host",
					Description: `(required): Host to look up ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Set to ` + "`" + `host` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "addrs",
					Description: `A list of IP addresses. IP addresses are always sorted to avoid constant changing plans.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Set to ` + "`" + `host` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "addrs",
					Description: `A list of IP addresses. IP addresses are always sorted to avoid constant changing plans.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dns_dns_cname_record_set",
			Category:         "Data Sources",
			ShortDescription: `Get DNS CNAME record set.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "host",
					Description: `(required): Host to look up ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Set to ` + "`" + `host` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cname",
					Description: `A CNAME record associated with host.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Set to ` + "`" + `host` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cname",
					Description: `A CNAME record associated with host.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dns_dns_mx_record_set",
			Category:         "Data Sources",
			ShortDescription: `Get DNS MX record set.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `(Required): Domain to look up ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Set to ` + "`" + `service` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mx",
					Description: `A list of records. They are sorted by ascending preference then alphabetically by exchange to stay consistent across runs.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Set to ` + "`" + `service` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mx",
					Description: `A list of records. They are sorted by ascending preference then alphabetically by exchange to stay consistent across runs.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dns_dns_ns_record_set",
			Category:         "Data Sources",
			ShortDescription: `Get DNS ns record set.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "host",
					Description: `(required): Host to look up ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Set to ` + "`" + `host` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "nameservers",
					Description: `A list of nameservers. Nameservers are always sorted to avoid constant changing plans.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Set to ` + "`" + `host` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "nameservers",
					Description: `A list of nameservers. Nameservers are always sorted to avoid constant changing plans.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dns_dns_ptr_record_set",
			Category:         "Data Sources",
			ShortDescription: `Get DNS PTR record set.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_address",
					Description: `(required): IP address to look up ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Set to ` + "`" + `ip_address` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ptr",
					Description: `A PTR record associated with ` + "`" + `ip_address` + "`" + `. __NOTE__: Only the first result is taken from the query.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Set to ` + "`" + `ip_address` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ptr",
					Description: `A PTR record associated with ` + "`" + `ip_address` + "`" + `. __NOTE__: Only the first result is taken from the query.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dns_dns_srv_record_set",
			Category:         "Data Sources",
			ShortDescription: `Get DNS SRV record set.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service",
					Description: `(Required): Service to look up ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Set to ` + "`" + `service` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "srv",
					Description: `A list of records. They are sorted to stay consistent across runs.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Set to ` + "`" + `service` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "srv",
					Description: `A list of records. They are sorted to stay consistent across runs.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dns_dns_txt_record_set",
			Category:         "Data Sources",
			ShortDescription: `Get DNS TXT record set.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "host",
					Description: `(required): Host to look up ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Set to ` + "`" + `host` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "record",
					Description: `The first TXT record.`,
				},
				resource.Attribute{
					Name:        "records",
					Description: `A list of TXT records.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Set to ` + "`" + `host` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "record",
					Description: `The first TXT record.`,
				},
				resource.Attribute{
					Name:        "records",
					Description: `A list of TXT records.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"dns_dns_a_record_set":     0,
		"dns_dns_aaaa_record_set":  1,
		"dns_dns_cname_record_set": 2,
		"dns_dns_mx_record_set":    3,
		"dns_dns_ns_record_set":    4,
		"dns_dns_ptr_record_set":   5,
		"dns_dns_srv_record_set":   6,
		"dns_dns_txt_record_set":   7,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
