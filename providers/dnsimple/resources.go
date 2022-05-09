package dnsimple

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "dnsimple_domain",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The domain name to be created ## Import DNSimple domains can be imported using their numeric record ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import dnsimple_domain.resource_name 5678 ` + "`" + `` + "`" + `` + "`" + ` The record ID can be found within [DNSimple Domains API](https://developer.dnsimple.com/v2/domains/#listDomains). Check out [Authentication](https://developer.dnsimple.com/v2/#authentication) in API Overview for available options. ` + "`" + `` + "`" + `` + "`" + ` $ curl -u 'EMAIL:PASSWORD' https://api.dnsimple.com/v2/1234/domains?name_like=example.com | jq { "data": [ { "id": 5678, "account_id": 1234, "registrant_id": null, "name": "example.com", "unicode_name": "example.com", "state": "hosted", "auto_renew": false, "private_whois": false, "expires_on": null, "expires_at": null, "created_at": "2021-10-01T00:00:00Z", "updated_at": "2021-10-01T00:00:00Z" } ], "pagination": { "current_page": 1, "per_page": 30, "total_entries": 1, "total_pages": 1 } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dnsimple_email_forward",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) The domain to add the email forwarding rule to`,
				},
				resource.Attribute{
					Name:        "alias_name",
					Description: `The name part (the part before the @) of the source email address on the domain`,
				},
				resource.Attribute{
					Name:        "destination_email",
					Description: `(Required) The destination email address on another domain ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The email forward ID`,
				},
				resource.Attribute{
					Name:        "alias_name",
					Description: `The name part (the part before the @) of the source email address on the domain`,
				},
				resource.Attribute{
					Name:        "alias_email",
					Description: `The source email address on the domain`,
				},
				resource.Attribute{
					Name:        "destination_email",
					Description: `The destination email address on another domain`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The email forward ID`,
				},
				resource.Attribute{
					Name:        "alias_name",
					Description: `The name part (the part before the @) of the source email address on the domain`,
				},
				resource.Attribute{
					Name:        "alias_email",
					Description: `The source email address on the domain`,
				},
				resource.Attribute{
					Name:        "destination_email",
					Description: `The destination email address on another domain`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dnsimple_lets_encrypt_certificate",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Required) The domain to be issued the certificate for`,
				},
				resource.Attribute{
					Name:        "contact_id",
					Description: `(Required) The contact id for the certificate ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The certificate ID`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `The domain ID`,
				},
				resource.Attribute{
					Name:        "contact_id",
					Description: `The contact ID`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The certificate name`,
				},
				resource.Attribute{
					Name:        "years",
					Description: `The years the certificate will last`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the certificate`,
				},
				resource.Attribute{
					Name:        "authority_identifier",
					Description: `The identifying certification authority (CA)`,
				},
				resource.Attribute{
					Name:        "auto_renew",
					Description: `Set to true if the certificate will auto-renew`,
				},
				resource.Attribute{
					Name:        "csr",
					Description: `The certificate signing request`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The certificate ID`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `The domain ID`,
				},
				resource.Attribute{
					Name:        "contact_id",
					Description: `The contact ID`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The certificate name`,
				},
				resource.Attribute{
					Name:        "years",
					Description: `The years the certificate will last`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the certificate`,
				},
				resource.Attribute{
					Name:        "authority_identifier",
					Description: `The identifying certification authority (CA)`,
				},
				resource.Attribute{
					Name:        "auto_renew",
					Description: `Set to true if the certificate will auto-renew`,
				},
				resource.Attribute{
					Name:        "csr",
					Description: `The certificate signing request`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dnsimple_zone_record",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_name",
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
					Name:        "zone_id",
					Description: `The domain ID of the record`,
				},
				resource.Attribute{
					Name:        "qualified_name",
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
					Name:        "zone_id",
					Description: `The domain ID of the record`,
				},
				resource.Attribute{
					Name:        "qualified_name",
					Description: `The FQDN of the record ## Import DNSimple resources can be imported using their parent zone name (domain name) and numeric record ID.`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"dnsimple_domain":                   0,
		"dnsimple_email_forward":            1,
		"dnsimple_lets_encrypt_certificate": 2,
		"dnsimple_zone_record":              3,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
