package dnsimple

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "dnsimple_contact",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "label",
					Description: `Label`,
				},
				resource.Attribute{
					Name:        "first_name",
					Description: `(Required) First name`,
				},
				resource.Attribute{
					Name:        "last_name",
					Description: `(Required) Last name`,
				},
				resource.Attribute{
					Name:        "organization_name",
					Description: `Organization name`,
				},
				resource.Attribute{
					Name:        "job_title",
					Description: `Job title`,
				},
				resource.Attribute{
					Name:        "address1",
					Description: `(Required) Address line 1`,
				},
				resource.Attribute{
					Name:        "address2",
					Description: `Address line 2`,
				},
				resource.Attribute{
					Name:        "city",
					Description: `(Required) City`,
				},
				resource.Attribute{
					Name:        "state_province",
					Description: `(Required) State province`,
				},
				resource.Attribute{
					Name:        "postal_code",
					Description: `(Required) Postal code`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `(Required) Country`,
				},
				resource.Attribute{
					Name:        "phone",
					Description: `(Required) Phone`,
				},
				resource.Attribute{
					Name:        "fax",
					Description: `Fax`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Required) Email # Attributes Reference - ` + "`" + `id` + "`" + ` - The ID of this resource. - ` + "`" + `account_id` + "`" + ` - The account ID for the contact. - ` + "`" + `phone_normalized` + "`" + ` - The phone number, normalized. - ` + "`" + `fax_normalized` + "`" + ` - The fax number, normalized. - ` + "`" + `created_at` + "`" + ` - Timestamp representing when this contact was created. - ` + "`" + `updated_at` + "`" + ` - Timestamp representing when this contact was updated. ## Import DNSimple contacts can be imported using their numeric ID. ` + "`" + `` + "`" + `` + "`" + `bash terraform import dnsimple_contact.resource_name 5678 ` + "`" + `` + "`" + `` + "`" + ` The ID can be found within [DNSimple Contacts API](https://developer.dnsimple.com/v2/contacts/#listContacts). Check out [Authentication](https://developer.dnsimple.com/v2/#authentication) in API Overview for available options. ` + "`" + `` + "`" + `` + "`" + `bash curl -u 'EMAIL:PASSWORD' https://api.dnsimple.com/v2/1234/contacts?label_like=example.com | jq { "data": [ { "id": 1, "account_id": 1010, "label": "Default", "first_name": "First", "last_name": "User", "job_title": "CEO", "organization_name": "Awesome Company", "email": "first@example.com", "phone": "+18001234567", "fax": "+18011234567", "address1": "Italian Street, 10", "address2": "", "city": "Roma", "state_province": "RM", "postal_code": "00100", "country": "IT", "created_at": "2013-11-08T17:23:15Z", "updated_at": "2015-01-08T21:30:50Z" }, { "id": 2, "account_id": 1010, "label": "", "first_name": "Second", "last_name": "User", "job_title": "", "organization_name": "", "email": "second@example.com", "phone": "+18881234567", "fax": "", "address1": "French Street", "address2": "c/o Someone", "city": "Paris", "state_province": "XY", "postal_code": "00200", "country": "FR", "created_at": "2014-12-06T15:46:18Z", "updated_at": "2014-12-06T15:46:18Z" } ], "pagination": { "current_page": 1, "per_page": 30, "total_entries": 2, "total_pages": 1 } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
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
					Description: `(Required) The domain name to be created # Attributes Reference - ` + "`" + `id` + "`" + ` - The ID of this resource. - ` + "`" + `account_id` + "`" + ` - The account ID for the domain. - ` + "`" + `auto_renew` + "`" + ` - Whether the domain is set to auto-renew. - ` + "`" + `private_whois` + "`" + ` - Whether the domain has WhoIs privacy enabled. - ` + "`" + `registrant_id` + "`" + ` - The ID of the registrant (contact) for the domain. - ` + "`" + `state` + "`" + ` - The state of the domain. - ` + "`" + `unicode_name` + "`" + ` - The domain name in Unicode format. ## Import DNSimple domains can be imported using their numeric record ID. ` + "`" + `` + "`" + `` + "`" + `bash terraform import dnsimple_domain.resource_name 5678 ` + "`" + `` + "`" + `` + "`" + ` The record ID can be found within [DNSimple Domains API](https://developer.dnsimple.com/v2/domains/#listDomains). Check out [Authentication](https://developer.dnsimple.com/v2/#authentication) in API Overview for available options. ` + "`" + `` + "`" + `` + "`" + `bash curl -u 'EMAIL:PASSWORD' https://api.dnsimple.com/v2/1234/domains?name_like=example.com | jq { "data": [ { "id": 5678, "account_id": 1234, "registrant_id": null, "name": "example.com", "unicode_name": "example.com", "state": "hosted", "auto_renew": false, "private_whois": false, "expires_on": null, "expires_at": null, "created_at": "2021-10-01T00:00:00Z", "updated_at": "2021-10-01T00:00:00Z" } ], "pagination": { "current_page": 1, "per_page": 30, "total_entries": 1, "total_pages": 1 } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dnsimple_domain_delegation",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) The domain name.`,
				},
				resource.Attribute{
					Name:        "name_servers",
					Description: `(Required) The list of name servers to delegate to. # Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The domain name. ## Import DNSimple domain delegations can be imported using the domain name.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dnsimple_ds_record",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) The domain name or numeric ID to create the delegation signer record for.`,
				},
				resource.Attribute{
					Name:        "algorithm",
					Description: `(Required) DNSSEC algorithm number as a string.`,
				},
				resource.Attribute{
					Name:        "digest",
					Description: `(Optional) The hexidecimal representation of the digest of the corresponding DNSKEY record.`,
				},
				resource.Attribute{
					Name:        "digest_type",
					Description: `(Optional) DNSSEC digest type number as a string.`,
				},
				resource.Attribute{
					Name:        "keytag",
					Description: `(Optional) A keytag that references the corresponding DNSKEY record.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(Optional) A public key that references the corresponding DNSKEY record. # Attributes Reference - ` + "`" + `id` + "`" + ` - The ID of this resource. - ` + "`" + `created_at` + "`" + ` - The time the DS record was created at. - ` + "`" + `updated_at` + "`" + ` - The time the DS record was last updated at. ## Import DNSimple DS record resources can be imported using their domain ID and numeric record ID. ` + "`" + `` + "`" + `` + "`" + `bash terraform import dnsimple_domain_ds_signer.resource_name example.com_5678 ` + "`" + `` + "`" + `` + "`" + ` The record ID can be found within [DNSimple DNSSEC API](https://developer.dnsimple.com/v2/domains/dnssec/#listDomainDelegationSignerRecords). Check out [Authentication](https://developer.dnsimple.com/v2/#authentication) in API Overview for available options. ` + "`" + `` + "`" + `` + "`" + `bash curl -u 'EMAIL:PASSWORD' https://api.dnsimple.com/v2/1010/domains/example.com/ds_records | jq { "data": [ { "id": 24, "domain_id": 1010, "algorithm": "8", "digest": "C1F6E04A5A61FBF65BF9DC8294C363CF11C89E802D926BDAB79C55D27BEFA94F", "digest_type": "2", "keytag": "44620", "public_key": null, "created_at": "2017-03-03T13:49:58Z", "updated_at": "2017-03-03T13:49:58Z" } ], "pagination": { "current_page": 1, "per_page": 30, "total_entries": 1, "total_pages": 1 } } ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Required) The domain name to add the email forwarding rule to`,
				},
				resource.Attribute{
					Name:        "alias_name",
					Description: `The name part (the part before the @) of the source email address on the domain`,
				},
				resource.Attribute{
					Name:        "destination_email",
					Description: `(Required) The destination email address ## Attributes Reference The following additional attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The email forward ID`,
				},
				resource.Attribute{
					Name:        "alias_email",
					Description: `The source email address on the domain, in full form. This is a computed attribute. ## Import DNSimple resources can be imported using the domain name and numeric email forward ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The email forward ID`,
				},
				resource.Attribute{
					Name:        "alias_email",
					Description: `The source email address on the domain, in full form. This is a computed attribute. ## Import DNSimple resources can be imported using the domain name and numeric email forward ID.`,
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
					Name:        "name",
					Description: `(Required) The certificate name`,
				},
				resource.Attribute{
					Name:        "alternate_names",
					Description: `(Optional) The certificate alternate names`,
				},
				resource.Attribute{
					Name:        "auto_renew",
					Description: `(Required) True if the certificate should auto-renew`,
				},
				resource.Attribute{
					Name:        "signature_algorithm",
					Description: `(Optional) The signature algorithm to use for the certificate`,
				},
				resource.Attribute{
					Name:        "timeouts",
					Description: `(Block, Optional) (see [below for nested schema](#nestedblock--timeouts)) ## Attribute Reference The following additional attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The certificate ID`,
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
					Name:        "csr",
					Description: `The certificate signing request`,
				},
				resource.Attribute{
					Name:        "expires_at",
					Description: `The datetime the certificate will expire`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The datetime the certificate was created`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The datetime the certificate was last updated <a id="nestedblock--timeouts"></a> ### Nested Schema for ` + "`" + `timeouts` + "`" + ` Optional: - ` + "`" + `read` + "`" + ` (String) - The timeout for the read operation`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The certificate ID`,
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
					Name:        "csr",
					Description: `The certificate signing request`,
				},
				resource.Attribute{
					Name:        "expires_at",
					Description: `The datetime the certificate will expire`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The datetime the certificate was created`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The datetime the certificate was last updated <a id="nestedblock--timeouts"></a> ### Nested Schema for ` + "`" + `timeouts` + "`" + ` Optional: - ` + "`" + `read` + "`" + ` (String) - The timeout for the read operation`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dnsimple_registered_domain",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The domain name to be registered`,
				},
				resource.Attribute{
					Name:        "contact_id",
					Description: `(Required) The ID of the contact to be used for the domain registration`,
				},
				resource.Attribute{
					Name:        "auto_renew_enabled",
					Description: `(Optional) Whether the domain should be set to auto-renew (default: ` + "`" + `false` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "whois_privacy_enabled",
					Description: `(Optional) Whether the domain should have WhoIs privacy enabled (default: ` + "`" + `false` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "dnssec_enabled",
					Description: `(Optional) Whether the domain should have DNSSEC enabled (default: ` + "`" + `false` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "premium_price",
					Description: `(Optional) The premium price for the domain registration. This is only required if the domain is a premium domain. You can use our [Check domain API](https://developer.dnsimple.com/v2/registrar/#checkDomain) to check if a domain is premium. And [Retrieve domain prices API](https://developer.dnsimple.com/v2/registrar/#getDomainPrices) to retrieve the premium price for a domain.`,
				},
				resource.Attribute{
					Name:        "extended_attributes",
					Description: `(Optional) A map of extended attributes to be set for the domain registration. To see if there are any required extended attributes for any TLD use our [Lists the TLD Extended Attributes API](https://developer.dnsimple.com/v2/tlds/#getTldExtendedAttributes).`,
				},
				resource.Attribute{
					Name:        "timeouts",
					Description: `(Optional) (see [below for nested schema](#nestedblock--timeouts)) # Attributes Reference - ` + "`" + `id` + "`" + ` - The ID of this resource. - ` + "`" + `unicode_name` + "`" + ` - The domain name in Unicode format. - ` + "`" + `state` + "`" + ` - The state of the domain. - ` + "`" + `domain_registration` + "`" + ` - (Block) The domain registration details. (see [below for nested schema](#nestedblock--domain_registration)) <a id="nestedblock--timeouts"></a> ### Nested Schema for ` + "`" + `timeouts` + "`" + ` Optional: - ` + "`" + `create` + "`" + ` (String) - The timeout for the read operation e.g. ` + "`" + `5m` + "`" + ` - ` + "`" + `update` + "`" + ` (String) - The timeout for the read operation e.g. ` + "`" + `5m` + "`" + ` <a id="nestedblock--domain_registration"></a> ### Nested Schema for ` + "`" + `domain_registration` + "`" + ` Attributes Reference: - ` + "`" + `id` + "`" + ` (Number) - The ID of the domain registration. - ` + "`" + `state` + "`" + ` (String) - The state of the domain registration. - ` + "`" + `period` + "`" + ` (Number) - The registration period in years. ## Import DNSimple registered domains can be imported using their domain name and`,
				},
			},
			Attributes: []resource.Attribute{},
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
					Description: `(Required) The zone name to add the record to`,
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
					Description: `(Optional) The TTL of the record - defaults to 3600`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) The priority of the record - only useful for some record types ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The record ID`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `The zone ID of the record`,
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
					Name:        "zone_id",
					Description: `The zone ID of the record`,
				},
				resource.Attribute{
					Name:        "qualified_name",
					Description: `The FQDN of the record ## Import DNSimple resources can be imported using their parent zone name (domain name) and numeric record ID.`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"dnsimple_contact":                  0,
		"dnsimple_domain":                   1,
		"dnsimple_domain_delegation":        2,
		"dnsimple_ds_record":                3,
		"dnsimple_email_forward":            4,
		"dnsimple_lets_encrypt_certificate": 5,
		"dnsimple_registered_domain":        6,
		"dnsimple_zone_record":              7,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
