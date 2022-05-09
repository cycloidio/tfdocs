package dns

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "dns_dns_a_record_set",
			Category:         "Resources",
			ShortDescription: `Creates a A type DNS record set.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) DNS zone the record set belongs to. It must be an FQDN, that is, include the trailing dot.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the record set. The ` + "`" + `zone` + "`" + ` argument will be appended to this value to create the full record path.`,
				},
				resource.Attribute{
					Name:        "addresses",
					Description: `(Required) The IPv4 addresses this record set will point to.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The TTL of the record set. Defaults to ` + "`" + `3600` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "addresses",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `See Argument Reference above. ## Import Records can be imported using the FQDN, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import dns_a_record_set.www www.example.com. ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "addresses",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `See Argument Reference above. ## Import Records can be imported using the FQDN, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import dns_a_record_set.www www.example.com. ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dns_dns_aaaa_record_set",
			Category:         "Resources",
			ShortDescription: `Creates a AAAA type DNS record set.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) DNS zone the record set belongs to. It must be an FQDN, that is, include the trailing dot.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the record set. The ` + "`" + `zone` + "`" + ` argument will be appended to this value to create the full record path.`,
				},
				resource.Attribute{
					Name:        "addresses",
					Description: `(Required) The IPv6 addresses this record set will point to.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The TTL of the record set. Defaults to ` + "`" + `3600` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "addresses",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `See Argument Reference above. ## Import Records can be imported using the FQDN, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import dns_aaaa_record_set.www www.example.com. ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "addresses",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `See Argument Reference above. ## Import Records can be imported using the FQDN, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import dns_aaaa_record_set.www www.example.com. ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dns_dns_cname_record",
			Category:         "Resources",
			ShortDescription: `Creates a CNAME type DNS record.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) DNS zone the record belongs to. It must be an FQDN, that is, include the trailing dot.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the record. The ` + "`" + `zone` + "`" + ` argument will be appended to this value to create the full record path.`,
				},
				resource.Attribute{
					Name:        "cname",
					Description: `(Required) The canonical name this record will point to.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The TTL of the record set. Defaults to ` + "`" + `3600` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cname",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `See Argument Reference above. ## Import Records can be imported using the FQDN, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import dns_cname_record.foo foo.example.com. ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cname",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `See Argument Reference above. ## Import Records can be imported using the FQDN, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import dns_cname_record.foo foo.example.com. ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dns_dns_mx_record_set",
			Category:         "Resources",
			ShortDescription: `Creates an MX type DNS record set.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) DNS zone the record set belongs to. It must be an FQDN, that is, include the trailing dot.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the record set. The ` + "`" + `zone` + "`" + ` argument will be appended to this value to create the full record path.`,
				},
				resource.Attribute{
					Name:        "mx",
					Description: `(Required) Can be specified multiple times for each MX record. Each block supports fields documented below.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The TTL of the record set. Defaults to ` + "`" + `3600` + "`" + `. The ` + "`" + `mx` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "preference",
					Description: `(Required) The preference for the record.`,
				},
				resource.Attribute{
					Name:        "exchange",
					Description: `(Required) The FQDN of the mail exchange, include the trailing dot. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mx",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `See Argument Reference above. ## Import Records can be imported using the FQDN, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ terraform import dns_mx_record_set.mx example.com. ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mx",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `See Argument Reference above. ## Import Records can be imported using the FQDN, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ terraform import dns_mx_record_set.mx example.com. ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dns_dns_ns_record_set",
			Category:         "Resources",
			ShortDescription: `Creates a NS type DNS record set.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) DNS zone the record set belongs to. It must be an FQDN, that is, include the trailing dot.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the record set. The ` + "`" + `zone` + "`" + ` argument will be appended to this value to create the full record path.`,
				},
				resource.Attribute{
					Name:        "nameservers",
					Description: `(Required) The nameservers this record set will point to.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The TTL of the record set. Defaults to ` + "`" + `3600` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "nameservers",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `See Argument Reference above. ## Import Records can be imported using the FQDN, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import dns_ns_record_set.www www.example.com. ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "nameservers",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `See Argument Reference above. ## Import Records can be imported using the FQDN, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import dns_ns_record_set.www www.example.com. ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dns_dns_ptr_record",
			Category:         "Resources",
			ShortDescription: `Creates a PTR type DNS record.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) DNS zone the record belongs to. It must be an FQDN, that is, include the trailing dot.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the record. The ` + "`" + `zone` + "`" + ` argument will be appended to this value to create the full record path.`,
				},
				resource.Attribute{
					Name:        "ptr",
					Description: `(Required) The canonical name this record will point to.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The TTL of the record set. Defaults to ` + "`" + `3600` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ptr",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `See Argument Reference above. ## Import Records can be imported using the FQDN, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import dns_ptr_record.dns-sd r._dns-sd.example.com. ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ptr",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `See Argument Reference above. ## Import Records can be imported using the FQDN, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import dns_ptr_record.dns-sd r._dns-sd.example.com. ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dns_dns_srv_record_set",
			Category:         "Resources",
			ShortDescription: `Creates an SRV type DNS record set.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) DNS zone the record set belongs to. It must be an FQDN, that is, include the trailing dot.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the record set. The ` + "`" + `zone` + "`" + ` argument will be appended to this value to create the full record path.`,
				},
				resource.Attribute{
					Name:        "srv",
					Description: `(Required) Can be specified multiple times for each SRV record. Each block supports fields documented below.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The TTL of the record set. Defaults to ` + "`" + `3600` + "`" + `. The ` + "`" + `srv` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Required) The priority for the record.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Required) The weight for the record.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Required) The FQDN of the target, include the trailing dot.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) The port for the service on the target. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "srv",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `See Argument Reference above. ## Import Records can be imported using the FQDN, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import dns_srv_record_set.sip _sip._tcp.example.com. ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "srv",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `See Argument Reference above. ## Import Records can be imported using the FQDN, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import dns_srv_record_set.sip _sip._tcp.example.com. ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dns_dns_txt_record_set",
			Category:         "Resources",
			ShortDescription: `Creates a TXT type DNS record set.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) DNS zone the record set belongs to. It must be an FQDN, that is, include the trailing dot.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the record set. The ` + "`" + `zone` + "`" + ` argument will be appended to this value to create the full record path.`,
				},
				resource.Attribute{
					Name:        "txt",
					Description: `(Required) The text records this record set will be set to.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The TTL of the record set. Defaults to ` + "`" + `3600` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "txt",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `See Argument Reference above. ## Import Records can be imported using the FQDN, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import dns_txt_record_set.google example.com. ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "txt",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `See Argument Reference above. ## Import Records can be imported using the FQDN, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import dns_txt_record_set.google example.com. ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"dns_dns_a_record_set":    0,
		"dns_dns_aaaa_record_set": 1,
		"dns_dns_cname_record":    2,
		"dns_dns_mx_record_set":   3,
		"dns_dns_ns_record_set":   4,
		"dns_dns_ptr_record":      5,
		"dns_dns_srv_record_set":  6,
		"dns_dns_txt_record_set":  7,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
