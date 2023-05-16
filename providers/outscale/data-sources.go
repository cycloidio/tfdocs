package outscale

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "outscale_access_key",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about an access key.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "access_key_ids",
					Description: `(Optional) The IDs of the access keys.`,
				},
				resource.Attribute{
					Name:        "states",
					Description: `(Optional) The states of the access keys (` + "`" + `ACTIVE` + "`" + ` \| ` + "`" + `INACTIVE` + "`" + `). ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "access_key_id",
					Description: `The ID of the access key.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `The date and time (UTC) of creation of the access key.`,
				},
				resource.Attribute{
					Name:        "expiration_date",
					Description: `The date (UTC) at which the access key expires.`,
				},
				resource.Attribute{
					Name:        "last_modification_date",
					Description: `The date and time (UTC) of the last modification of the access key.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the access key (` + "`" + `ACTIVE` + "`" + ` if the key is valid for API calls, or ` + "`" + `INACTIVE` + "`" + ` if not).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "access_key_id",
					Description: `The ID of the access key.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `The date and time (UTC) of creation of the access key.`,
				},
				resource.Attribute{
					Name:        "expiration_date",
					Description: `The date (UTC) at which the access key expires.`,
				},
				resource.Attribute{
					Name:        "last_modification_date",
					Description: `The date and time (UTC) of the last modification of the access key.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the access key (` + "`" + `ACTIVE` + "`" + ` if the key is valid for API calls, or ` + "`" + `INACTIVE` + "`" + ` if not).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_access_keys",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about access keys.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "access_key_ids",
					Description: `(Optional) The IDs of the access keys.`,
				},
				resource.Attribute{
					Name:        "states",
					Description: `(Optional) The states of the access keys (` + "`" + `ACTIVE` + "`" + ` \| ` + "`" + `INACTIVE` + "`" + `). ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "access_keys",
					Description: `A list of access keys.`,
				},
				resource.Attribute{
					Name:        "access_key_id",
					Description: `The ID of the access key.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `The date and time (UTC) of creation of the access key.`,
				},
				resource.Attribute{
					Name:        "expiration_date",
					Description: `The date (UTC) at which the access key expires.`,
				},
				resource.Attribute{
					Name:        "last_modification_date",
					Description: `The date and time (UTC) of the last modification of the access key.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the access key (` + "`" + `ACTIVE` + "`" + ` if the key is valid for API calls, or ` + "`" + `INACTIVE` + "`" + ` if not).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "access_keys",
					Description: `A list of access keys.`,
				},
				resource.Attribute{
					Name:        "access_key_id",
					Description: `The ID of the access key.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `The date and time (UTC) of creation of the access key.`,
				},
				resource.Attribute{
					Name:        "expiration_date",
					Description: `The date (UTC) at which the access key expires.`,
				},
				resource.Attribute{
					Name:        "last_modification_date",
					Description: `The date and time (UTC) of the last modification of the access key.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the access key (` + "`" + `ACTIVE` + "`" + ` if the key is valid for API calls, or ` + "`" + `INACTIVE` + "`" + ` if not).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_account",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about an account.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_id",
					Description: `The ID of the account.`,
				},
				resource.Attribute{
					Name:        "additional_emails",
					Description: `One or more additional email addresses for the account. These addresses are used for notifications only.`,
				},
				resource.Attribute{
					Name:        "city",
					Description: `The city of the account owner.`,
				},
				resource.Attribute{
					Name:        "company_name",
					Description: `The name of the company for the account.`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `The country of the account owner.`,
				},
				resource.Attribute{
					Name:        "customer_id",
					Description: `The ID of the customer.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `The main email address for the account. This address is used for your credentials and for notifications.`,
				},
				resource.Attribute{
					Name:        "first_name",
					Description: `The first name of the account owner.`,
				},
				resource.Attribute{
					Name:        "job_title",
					Description: `The job title of the account owner.`,
				},
				resource.Attribute{
					Name:        "last_name",
					Description: `The last name of the account owner.`,
				},
				resource.Attribute{
					Name:        "mobile_number",
					Description: `The mobile phone number of the account owner.`,
				},
				resource.Attribute{
					Name:        "phone_number",
					Description: `The landline phone number of the account owner.`,
				},
				resource.Attribute{
					Name:        "state_province",
					Description: `The state/province of the account.`,
				},
				resource.Attribute{
					Name:        "vat_number",
					Description: `The value added tax (VAT) number for the account.`,
				},
				resource.Attribute{
					Name:        "zip_code",
					Description: `The ZIP code of the city.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "account_id",
					Description: `The ID of the account.`,
				},
				resource.Attribute{
					Name:        "additional_emails",
					Description: `One or more additional email addresses for the account. These addresses are used for notifications only.`,
				},
				resource.Attribute{
					Name:        "city",
					Description: `The city of the account owner.`,
				},
				resource.Attribute{
					Name:        "company_name",
					Description: `The name of the company for the account.`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `The country of the account owner.`,
				},
				resource.Attribute{
					Name:        "customer_id",
					Description: `The ID of the customer.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `The main email address for the account. This address is used for your credentials and for notifications.`,
				},
				resource.Attribute{
					Name:        "first_name",
					Description: `The first name of the account owner.`,
				},
				resource.Attribute{
					Name:        "job_title",
					Description: `The job title of the account owner.`,
				},
				resource.Attribute{
					Name:        "last_name",
					Description: `The last name of the account owner.`,
				},
				resource.Attribute{
					Name:        "mobile_number",
					Description: `The mobile phone number of the account owner.`,
				},
				resource.Attribute{
					Name:        "phone_number",
					Description: `The landline phone number of the account owner.`,
				},
				resource.Attribute{
					Name:        "state_province",
					Description: `The state/province of the account.`,
				},
				resource.Attribute{
					Name:        "vat_number",
					Description: `The value added tax (VAT) number for the account.`,
				},
				resource.Attribute{
					Name:        "zip_code",
					Description: `The ZIP code of the city.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_accounts",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about accounts.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "accounts",
					Description: `The list of the accounts.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The ID of the account.`,
				},
				resource.Attribute{
					Name:        "additional_emails",
					Description: `One or more additional email addresses for the account. These addresses are used for notifications only.`,
				},
				resource.Attribute{
					Name:        "city",
					Description: `The city of the account owner.`,
				},
				resource.Attribute{
					Name:        "company_name",
					Description: `The name of the company for the account.`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `The country of the account owner.`,
				},
				resource.Attribute{
					Name:        "customer_id",
					Description: `The ID of the customer.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `The main email address for the account. This address is used for your credentials and for notifications.`,
				},
				resource.Attribute{
					Name:        "first_name",
					Description: `The first name of the account owner.`,
				},
				resource.Attribute{
					Name:        "job_title",
					Description: `The job title of the account owner.`,
				},
				resource.Attribute{
					Name:        "last_name",
					Description: `The last name of the account owner.`,
				},
				resource.Attribute{
					Name:        "mobile_number",
					Description: `The mobile phone number of the account owner.`,
				},
				resource.Attribute{
					Name:        "phone_number",
					Description: `The landline phone number of the account owner.`,
				},
				resource.Attribute{
					Name:        "state_province",
					Description: `The state/province of the account.`,
				},
				resource.Attribute{
					Name:        "vat_number",
					Description: `The value added tax (VAT) number for the account.`,
				},
				resource.Attribute{
					Name:        "zip_code",
					Description: `The ZIP code of the city.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "accounts",
					Description: `The list of the accounts.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The ID of the account.`,
				},
				resource.Attribute{
					Name:        "additional_emails",
					Description: `One or more additional email addresses for the account. These addresses are used for notifications only.`,
				},
				resource.Attribute{
					Name:        "city",
					Description: `The city of the account owner.`,
				},
				resource.Attribute{
					Name:        "company_name",
					Description: `The name of the company for the account.`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `The country of the account owner.`,
				},
				resource.Attribute{
					Name:        "customer_id",
					Description: `The ID of the customer.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `The main email address for the account. This address is used for your credentials and for notifications.`,
				},
				resource.Attribute{
					Name:        "first_name",
					Description: `The first name of the account owner.`,
				},
				resource.Attribute{
					Name:        "job_title",
					Description: `The job title of the account owner.`,
				},
				resource.Attribute{
					Name:        "last_name",
					Description: `The last name of the account owner.`,
				},
				resource.Attribute{
					Name:        "mobile_number",
					Description: `The mobile phone number of the account owner.`,
				},
				resource.Attribute{
					Name:        "phone_number",
					Description: `The landline phone number of the account owner.`,
				},
				resource.Attribute{
					Name:        "state_province",
					Description: `The state/province of the account.`,
				},
				resource.Attribute{
					Name:        "vat_number",
					Description: `The value added tax (VAT) number for the account.`,
				},
				resource.Attribute{
					Name:        "zip_code",
					Description: `The ZIP code of the city.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_api_access_policy",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about the API access policy.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "max_access_key_expiration_seconds",
					Description: `The maximum possible lifetime for your access keys, in seconds. If ` + "`" + `0` + "`" + `, your access keys can have unlimited lifetimes.`,
				},
				resource.Attribute{
					Name:        "require_trusted_env",
					Description: `If true, a trusted session is activated, allowing you to bypass Certificate Authorities (CAs) enforcement. For more information, see the ` + "`" + `ApiKeyAuth` + "`" + ` authentication scheme in the [Authentication](https://docs.outscale.com/api#authentication) section.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "max_access_key_expiration_seconds",
					Description: `The maximum possible lifetime for your access keys, in seconds. If ` + "`" + `0` + "`" + `, your access keys can have unlimited lifetimes.`,
				},
				resource.Attribute{
					Name:        "require_trusted_env",
					Description: `If true, a trusted session is activated, allowing you to bypass Certificate Authorities (CAs) enforcement. For more information, see the ` + "`" + `ApiKeyAuth` + "`" + ` authentication scheme in the [Authentication](https://docs.outscale.com/api#authentication) section.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_api_access_rule",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about an API access rule.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "api_access_rule_ids",
					Description: `(Optional) One or more IDs of API access rules.`,
				},
				resource.Attribute{
					Name:        "ca_ids",
					Description: `(Optional) One or more IDs of Client Certificate Authorities (CAs).`,
				},
				resource.Attribute{
					Name:        "cns",
					Description: `(Optional) One or more Client Certificate Common Names (CNs).`,
				},
				resource.Attribute{
					Name:        "descriptions",
					Description: `(Optional) One or more descriptions of API access rules.`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `(Optional) One or more IP addresses or CIDR blocks (for example, ` + "`" + `192.0.2.0/16` + "`" + `). ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "api_access_rule_id",
					Description: `The ID of the API access rule.`,
				},
				resource.Attribute{
					Name:        "ca_ids",
					Description: `One or more IDs of Client Certificate Authorities (CAs) used for the API access rule.`,
				},
				resource.Attribute{
					Name:        "cns",
					Description: `One or more Client Certificate Common Names (CNs).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the API access rule.`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `One or more IP ranges used for the API access rule, in CIDR notation (for example, ` + "`" + `192.0.2.0/16` + "`" + `).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "api_access_rule_id",
					Description: `The ID of the API access rule.`,
				},
				resource.Attribute{
					Name:        "ca_ids",
					Description: `One or more IDs of Client Certificate Authorities (CAs) used for the API access rule.`,
				},
				resource.Attribute{
					Name:        "cns",
					Description: `One or more Client Certificate Common Names (CNs).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the API access rule.`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `One or more IP ranges used for the API access rule, in CIDR notation (for example, ` + "`" + `192.0.2.0/16` + "`" + `).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_api_access_rules",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about API access rules.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "api_access_rule_ids",
					Description: `(Optional) One or more IDs of API access rules.`,
				},
				resource.Attribute{
					Name:        "ca_ids",
					Description: `(Optional) One or more IDs of Client Certificate Authorities (CAs).`,
				},
				resource.Attribute{
					Name:        "cns",
					Description: `(Optional) One or more Client Certificate Common Names (CNs).`,
				},
				resource.Attribute{
					Name:        "descriptions",
					Description: `(Optional) One or more descriptions of API access rules.`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `(Optional) One or more IP addresses or CIDR blocks (for example, ` + "`" + `192.0.2.0/16` + "`" + `). ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "api_access_rules",
					Description: `A list of API access rules.`,
				},
				resource.Attribute{
					Name:        "api_access_rule_id",
					Description: `The ID of the API access rule.`,
				},
				resource.Attribute{
					Name:        "ca_ids",
					Description: `One or more IDs of Client Certificate Authorities (CAs) used for the API access rule.`,
				},
				resource.Attribute{
					Name:        "cns",
					Description: `One or more Client Certificate Common Names (CNs).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the API access rule.`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `One or more IP ranges used for the API access rule, in CIDR notation (for example, ` + "`" + `192.0.2.0/16` + "`" + `).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "api_access_rules",
					Description: `A list of API access rules.`,
				},
				resource.Attribute{
					Name:        "api_access_rule_id",
					Description: `The ID of the API access rule.`,
				},
				resource.Attribute{
					Name:        "ca_ids",
					Description: `One or more IDs of Client Certificate Authorities (CAs) used for the API access rule.`,
				},
				resource.Attribute{
					Name:        "cns",
					Description: `One or more Client Certificate Common Names (CNs).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the API access rule.`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `One or more IP ranges used for the API access rule, in CIDR notation (for example, ` + "`" + `192.0.2.0/16` + "`" + `).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_ca",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a Certificate Authority (CA).]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "ca_fingerprints",
					Description: `(Optional) The fingerprints of the CAs.`,
				},
				resource.Attribute{
					Name:        "ca_ids",
					Description: `(Optional) The IDs of the CAs.`,
				},
				resource.Attribute{
					Name:        "descriptions",
					Description: `(Optional) The descriptions of the CAs. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "ca_fingerprint",
					Description: `The fingerprint of the CA.`,
				},
				resource.Attribute{
					Name:        "ca_id",
					Description: `The ID of the CA.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the CA.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ca_fingerprint",
					Description: `The fingerprint of the CA.`,
				},
				resource.Attribute{
					Name:        "ca_id",
					Description: `The ID of the CA.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the CA.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_cas",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about Certificate Authorities (CAs).]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "ca_fingerprints",
					Description: `(Optional) The fingerprints of the CAs.`,
				},
				resource.Attribute{
					Name:        "ca_ids",
					Description: `(Optional) The IDs of the CAs.`,
				},
				resource.Attribute{
					Name:        "descriptions",
					Description: `(Optional) The descriptions of the CAs. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cas",
					Description: `Information about one or more CAs.`,
				},
				resource.Attribute{
					Name:        "ca_fingerprint",
					Description: `The fingerprint of the CA.`,
				},
				resource.Attribute{
					Name:        "ca_id",
					Description: `The ID of the CA.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the CA.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cas",
					Description: `Information about one or more CAs.`,
				},
				resource.Attribute{
					Name:        "ca_fingerprint",
					Description: `The fingerprint of the CA.`,
				},
				resource.Attribute{
					Name:        "ca_id",
					Description: `The ID of the CA.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the CA.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_client_gateway",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a client gateway.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "bgp_asns",
					Description: `(Optional) The Border Gateway Protocol (BGP) Autonomous System Numbers (ASNs) of the connections.`,
				},
				resource.Attribute{
					Name:        "client_gateway_ids",
					Description: `(Optional) The IDs of the client gateways.`,
				},
				resource.Attribute{
					Name:        "connection_types",
					Description: `(Optional) The types of communication tunnels used by the client gateways (only ` + "`" + `ipsec.1` + "`" + ` is supported).`,
				},
				resource.Attribute{
					Name:        "public_ips",
					Description: `(Optional) The public IPv4 addresses of the client gateways.`,
				},
				resource.Attribute{
					Name:        "states",
					Description: `(Optional) The states of the client gateways (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the client gateways.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the client gateways.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combinations of the tags associated with the client gateways, in the following format: ` + "`" + `TAGKEY=TAGVALUE` + "`" + `. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "bgp_asn",
					Description: `The Autonomous System Number (ASN) used by the Border Gateway Protocol (BGP) to find the path to your client gateway through the Internet.`,
				},
				resource.Attribute{
					Name:        "client_gateway_id",
					Description: `The ID of the client gateway.`,
				},
				resource.Attribute{
					Name:        "connection_type",
					Description: `The type of communication tunnel used by the client gateway (only ` + "`" + `ipsec.1` + "`" + ` is supported).`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IPv4 address of the client gateway (must be a fixed address into a NATed network).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the client gateway (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the client gateway.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "bgp_asn",
					Description: `The Autonomous System Number (ASN) used by the Border Gateway Protocol (BGP) to find the path to your client gateway through the Internet.`,
				},
				resource.Attribute{
					Name:        "client_gateway_id",
					Description: `The ID of the client gateway.`,
				},
				resource.Attribute{
					Name:        "connection_type",
					Description: `The type of communication tunnel used by the client gateway (only ` + "`" + `ipsec.1` + "`" + ` is supported).`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IPv4 address of the client gateway (must be a fixed address into a NATed network).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the client gateway (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the client gateway.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_client_gateways",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about client gateways.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "bgp_asns",
					Description: `(Optional) The Border Gateway Protocol (BGP) Autonomous System Numbers (ASNs) of the connections.`,
				},
				resource.Attribute{
					Name:        "client_gateway_ids",
					Description: `(Optional) The IDs of the client gateways.`,
				},
				resource.Attribute{
					Name:        "connection_types",
					Description: `(Optional) The types of communication tunnels used by the client gateways (only ` + "`" + `ipsec.1` + "`" + ` is supported).`,
				},
				resource.Attribute{
					Name:        "public_ips",
					Description: `(Optional) The public IPv4 addresses of the client gateways.`,
				},
				resource.Attribute{
					Name:        "states",
					Description: `(Optional) The states of the client gateways (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the client gateways.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the client gateways.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combinations of the tags associated with the client gateways, in the following format: ` + "`" + `TAGKEY=TAGVALUE` + "`" + `. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "client_gateways",
					Description: `Information about one or more client gateways.`,
				},
				resource.Attribute{
					Name:        "bgp_asn",
					Description: `The Autonomous System Number (ASN) used by the Border Gateway Protocol (BGP) to find the path to your client gateway through the Internet.`,
				},
				resource.Attribute{
					Name:        "client_gateway_id",
					Description: `The ID of the client gateway.`,
				},
				resource.Attribute{
					Name:        "connection_type",
					Description: `The type of communication tunnel used by the client gateway (only ` + "`" + `ipsec.1` + "`" + ` is supported).`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IPv4 address of the client gateway (must be a fixed address into a NATed network).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the client gateway (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the client gateway.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "client_gateways",
					Description: `Information about one or more client gateways.`,
				},
				resource.Attribute{
					Name:        "bgp_asn",
					Description: `The Autonomous System Number (ASN) used by the Border Gateway Protocol (BGP) to find the path to your client gateway through the Internet.`,
				},
				resource.Attribute{
					Name:        "client_gateway_id",
					Description: `The ID of the client gateway.`,
				},
				resource.Attribute{
					Name:        "connection_type",
					Description: `The type of communication tunnel used by the client gateway (only ` + "`" + `ipsec.1` + "`" + ` is supported).`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IPv4 address of the client gateway (must be a fixed address into a NATed network).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the client gateway (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the client gateway.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_dhcp_option",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a DHCP option.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `(Optional) If true, lists all default DHCP options set. If false, lists all non-default DHCP options set.`,
				},
				resource.Attribute{
					Name:        "dhcp_options_set_ids",
					Description: `(Optional) The IDs of the DHCP options sets.`,
				},
				resource.Attribute{
					Name:        "domain_name_servers",
					Description: `(Optional) The IPs of the domain name servers used for the DHCP options sets.`,
				},
				resource.Attribute{
					Name:        "domain_names",
					Description: `(Optional) The domain names used for the DHCP options sets.`,
				},
				resource.Attribute{
					Name:        "log_servers",
					Description: `(Optional) The IPs of the log servers used for the DHCP options sets.`,
				},
				resource.Attribute{
					Name:        "ntp_servers",
					Description: `(Optional) The IPs of the Network Time Protocol (NTP) servers used for the DHCP options sets.`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the DHCP options sets.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the DHCP options sets.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combinations of the tags associated with the DHCP options sets, in the following format: ` + "`" + `TAGKEY=TAGVALUE` + "`" + `. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `If true, the DHCP options set is a default one. If false, it is not.`,
				},
				resource.Attribute{
					Name:        "dhcp_options_set_id",
					Description: `The ID of the DHCP options set.`,
				},
				resource.Attribute{
					Name:        "domain_name_servers",
					Description: `One or more IPs for the domain name servers.`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `The domain name.`,
				},
				resource.Attribute{
					Name:        "log_servers",
					Description: `One or more IPs for the log servers.`,
				},
				resource.Attribute{
					Name:        "ntp_servers",
					Description: `One or more IPs for the NTP servers.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the DHCP options set.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "default",
					Description: `If true, the DHCP options set is a default one. If false, it is not.`,
				},
				resource.Attribute{
					Name:        "dhcp_options_set_id",
					Description: `The ID of the DHCP options set.`,
				},
				resource.Attribute{
					Name:        "domain_name_servers",
					Description: `One or more IPs for the domain name servers.`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `The domain name.`,
				},
				resource.Attribute{
					Name:        "log_servers",
					Description: `One or more IPs for the log servers.`,
				},
				resource.Attribute{
					Name:        "ntp_servers",
					Description: `One or more IPs for the NTP servers.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the DHCP options set.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_dhcp_options",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about DHCP options.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `(Optional) If true, lists all default DHCP options set. If false, lists all non-default DHCP options set.`,
				},
				resource.Attribute{
					Name:        "dhcp_options_set_ids",
					Description: `(Optional) The IDs of the DHCP options sets.`,
				},
				resource.Attribute{
					Name:        "domain_name_servers",
					Description: `(Optional) The IPs of the domain name servers used for the DHCP options sets.`,
				},
				resource.Attribute{
					Name:        "domain_names",
					Description: `(Optional) The domain names used for the DHCP options sets.`,
				},
				resource.Attribute{
					Name:        "log_servers",
					Description: `(Optional) The IPs of the log servers used for the DHCP options sets.`,
				},
				resource.Attribute{
					Name:        "ntp_servers",
					Description: `(Optional) The IPs of the Network Time Protocol (NTP) servers used for the DHCP options sets.`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the DHCP options sets.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the DHCP options sets.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combinations of the tags associated with the DHCP options sets, in the following format: ` + "`" + `TAGKEY=TAGVALUE` + "`" + `. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "dhcp_options_sets",
					Description: `Information about one or more DHCP options sets.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `If true, the DHCP options set is a default one. If false, it is not.`,
				},
				resource.Attribute{
					Name:        "dhcp_options_set_id",
					Description: `The ID of the DHCP options set.`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `The domain name.`,
				},
				resource.Attribute{
					Name:        "domain_name_servers",
					Description: `One or more IPs for the domain name servers.`,
				},
				resource.Attribute{
					Name:        "log_servers",
					Description: `One or more IPs for the log servers.`,
				},
				resource.Attribute{
					Name:        "ntp_servers",
					Description: `One or more IPs for the NTP servers.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the DHCP options set.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "dhcp_options_sets",
					Description: `Information about one or more DHCP options sets.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `If true, the DHCP options set is a default one. If false, it is not.`,
				},
				resource.Attribute{
					Name:        "dhcp_options_set_id",
					Description: `The ID of the DHCP options set.`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `The domain name.`,
				},
				resource.Attribute{
					Name:        "domain_name_servers",
					Description: `One or more IPs for the domain name servers.`,
				},
				resource.Attribute{
					Name:        "log_servers",
					Description: `One or more IPs for the log servers.`,
				},
				resource.Attribute{
					Name:        "ntp_servers",
					Description: `One or more IPs for the NTP servers.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the DHCP options set.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_flexible_gpu",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a flexible GPU.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `(Optional) Indicates whether the fGPU is deleted when terminating the VM.`,
				},
				resource.Attribute{
					Name:        "flexible_gpu_ids",
					Description: `(Optional) One or more IDs of fGPUs.`,
				},
				resource.Attribute{
					Name:        "generations",
					Description: `(Optional) The processor generations that the fGPUs are compatible with.`,
				},
				resource.Attribute{
					Name:        "model_names",
					Description: `(Optional) One or more models of fGPUs. For more information, see [About Flexible GPUs](https://docs.outscale.com/en/userguide/About-Flexible-GPUs.html).`,
				},
				resource.Attribute{
					Name:        "states",
					Description: `(Optional) The states of the fGPUs (` + "`" + `allocated` + "`" + ` \| ` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subregion_names",
					Description: `(Optional) The Subregions where the fGPUs are located.`,
				},
				resource.Attribute{
					Name:        "vm_ids",
					Description: `(Optional) One or more IDs of VMs. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If true, the fGPU is deleted when the VM is terminated.`,
				},
				resource.Attribute{
					Name:        "flexible_gpu_id",
					Description: `The ID of the fGPU.`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `The compatible processor generation.`,
				},
				resource.Attribute{
					Name:        "model_name",
					Description: `The model of fGPU. For more information, see [About Flexible GPUs](https://docs.outscale.com/en/userguide/About-Flexible-GPUs.html).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the fGPU (` + "`" + `allocated` + "`" + ` \| ` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The Subregion where the fGPU is located.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM the fGPU is attached to, if any.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If true, the fGPU is deleted when the VM is terminated.`,
				},
				resource.Attribute{
					Name:        "flexible_gpu_id",
					Description: `The ID of the fGPU.`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `The compatible processor generation.`,
				},
				resource.Attribute{
					Name:        "model_name",
					Description: `The model of fGPU. For more information, see [About Flexible GPUs](https://docs.outscale.com/en/userguide/About-Flexible-GPUs.html).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the fGPU (` + "`" + `allocated` + "`" + ` \| ` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The Subregion where the fGPU is located.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM the fGPU is attached to, if any.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_flexible_gpu_catalog",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about the flexible GPU catalog.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "flexible_gpu_catalog",
					Description: `Information about one or more fGPUs available in the public catalog.`,
				},
				resource.Attribute{
					Name:        "generations",
					Description: `The generations of VMs that the fGPU is compatible with.`,
				},
				resource.Attribute{
					Name:        "max_cpu",
					Description: `The maximum number of VM vCores that the fGPU is compatible with.`,
				},
				resource.Attribute{
					Name:        "max_ram",
					Description: `The maximum amount of VM memory that the fGPU is compatible with.`,
				},
				resource.Attribute{
					Name:        "model_name",
					Description: `The model of fGPU.`,
				},
				resource.Attribute{
					Name:        "v_ram",
					Description: `The amount of video RAM (VRAM) of the fGPU.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "flexible_gpu_catalog",
					Description: `Information about one or more fGPUs available in the public catalog.`,
				},
				resource.Attribute{
					Name:        "generations",
					Description: `The generations of VMs that the fGPU is compatible with.`,
				},
				resource.Attribute{
					Name:        "max_cpu",
					Description: `The maximum number of VM vCores that the fGPU is compatible with.`,
				},
				resource.Attribute{
					Name:        "max_ram",
					Description: `The maximum amount of VM memory that the fGPU is compatible with.`,
				},
				resource.Attribute{
					Name:        "model_name",
					Description: `The model of fGPU.`,
				},
				resource.Attribute{
					Name:        "v_ram",
					Description: `The amount of video RAM (VRAM) of the fGPU.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_flexible_gpus",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about flexible GPUs.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `(Optional) Indicates whether the fGPU is deleted when terminating the VM.`,
				},
				resource.Attribute{
					Name:        "flexible_gpu_ids",
					Description: `(Optional) One or more IDs of fGPUs.`,
				},
				resource.Attribute{
					Name:        "generations",
					Description: `(Optional) The processor generations that the fGPUs are compatible with.`,
				},
				resource.Attribute{
					Name:        "model_names",
					Description: `(Optional) One or more models of fGPUs. For more information, see [About Flexible GPUs](https://docs.outscale.com/en/userguide/About-Flexible-GPUs.html).`,
				},
				resource.Attribute{
					Name:        "states",
					Description: `(Optional) The states of the fGPUs (` + "`" + `allocated` + "`" + ` \| ` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subregion_names",
					Description: `(Optional) The Subregions where the fGPUs are located.`,
				},
				resource.Attribute{
					Name:        "vm_ids",
					Description: `(Optional) One or more IDs of VMs. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "flexible_gpus",
					Description: `Information about one or more fGPUs.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If true, the fGPU is deleted when the VM is terminated.`,
				},
				resource.Attribute{
					Name:        "flexible_gpu_id",
					Description: `The ID of the fGPU.`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `The compatible processor generation.`,
				},
				resource.Attribute{
					Name:        "model_name",
					Description: `The model of fGPU. For more information, see [About Flexible GPUs](https://docs.outscale.com/en/userguide/About-Flexible-GPUs.html).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the fGPU (` + "`" + `allocated` + "`" + ` \| ` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The Subregion where the fGPU is located.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM the fGPU is attached to, if any.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "flexible_gpus",
					Description: `Information about one or more fGPUs.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If true, the fGPU is deleted when the VM is terminated.`,
				},
				resource.Attribute{
					Name:        "flexible_gpu_id",
					Description: `The ID of the fGPU.`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `The compatible processor generation.`,
				},
				resource.Attribute{
					Name:        "model_name",
					Description: `The model of fGPU. For more information, see [About Flexible GPUs](https://docs.outscale.com/en/userguide/About-Flexible-GPUs.html).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the fGPU (` + "`" + `allocated` + "`" + ` \| ` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The Subregion where the fGPU is located.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM the fGPU is attached to, if any.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_image",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about an image.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "account_aliases",
					Description: `(Optional) The account aliases of the owners of the OMIs.`,
				},
				resource.Attribute{
					Name:        "account_ids",
					Description: `(Optional) The account IDs of the owners of the OMIs. By default, all the OMIs for which you have launch permissions are described.`,
				},
				resource.Attribute{
					Name:        "architectures",
					Description: `(Optional) The architectures of the OMIs (` + "`" + `i386` + "`" + ` \| ` + "`" + `x86_64` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "block_device_mapping_delete_on_vm_deletion",
					Description: `(Optional) Whether the volumes are deleted or not when terminating the VM.`,
				},
				resource.Attribute{
					Name:        "block_device_mapping_device_names",
					Description: `(Optional) The device names for the volumes.`,
				},
				resource.Attribute{
					Name:        "block_device_mapping_snapshot_ids",
					Description: `(Optional) The IDs of the snapshots used to create the volumes.`,
				},
				resource.Attribute{
					Name:        "block_device_mapping_volume_sizes",
					Description: `(Optional) The sizes of the volumes, in gibibytes (GiB).`,
				},
				resource.Attribute{
					Name:        "block_device_mapping_volume_types",
					Description: `(Optional) The types of volumes (` + "`" + `standard` + "`" + ` \| ` + "`" + `gp2` + "`" + ` \| ` + "`" + `io1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "descriptions",
					Description: `(Optional) The descriptions of the OMIs, provided when they were created.`,
				},
				resource.Attribute{
					Name:        "file_locations",
					Description: `(Optional) The locations of the buckets where the OMI files are stored.`,
				},
				resource.Attribute{
					Name:        "hypervisors",
					Description: `(Optional) The hypervisor type of the OMI (always ` + "`" + `xen` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "image_ids",
					Description: `(Optional) The IDs of the OMIs.`,
				},
				resource.Attribute{
					Name:        "image_names",
					Description: `(Optional) The names of the OMIs, provided when they were created.`,
				},
				resource.Attribute{
					Name:        "permissions_to_launch_account_ids",
					Description: `(Optional) The account IDs of the users who have launch permissions for the OMIs.`,
				},
				resource.Attribute{
					Name:        "permissions_to_launch_global_permission",
					Description: `(Optional) If true, lists all public OMIs. If false, lists all private OMIs.`,
				},
				resource.Attribute{
					Name:        "product_codes",
					Description: `(Optional) The product code associated with the OMI (` + "`" + `0001` + "`" + ` Linux/Unix \| ` + "`" + `0002` + "`" + ` Windows \| ` + "`" + `0004` + "`" + ` Linux/Oracle \| ` + "`" + `0005` + "`" + ` Windows 10).`,
				},
				resource.Attribute{
					Name:        "root_device_names",
					Description: `(Optional) The name of the root device. This value must be /dev/sda1.`,
				},
				resource.Attribute{
					Name:        "root_device_types",
					Description: `(Optional) The types of root device used by the OMIs (always ` + "`" + `bsu` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "states",
					Description: `(Optional) The states of the OMIs (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `failed` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the OMIs.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the OMIs.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combinations of the tags associated with the OMIs, in the following format: ` + "`" + `TAGKEY=TAGVALUE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "virtualization_types",
					Description: `(Optional) The virtualization types (always ` + "`" + `hvm` + "`" + `). ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "account_alias",
					Description: `The account alias of the owner of the OMI.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the OMI.`,
				},
				resource.Attribute{
					Name:        "architecture",
					Description: `The architecture of the OMI (by default, ` + "`" + `i386` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "block_device_mappings",
					Description: `One or more block device mappings.`,
				},
				resource.Attribute{
					Name:        "bsu",
					Description: `Information about the BSU volume to create.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `By default or if set to true, the volume is deleted when terminating the VM. If false, the volume is not deleted when terminating the VM.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `The number of I/O operations per second (IOPS). This parameter must be specified only if you create an ` + "`" + `io1` + "`" + ` volume. The maximum number of IOPS allowed for ` + "`" + `io1` + "`" + ` volumes is ` + "`" + `13000` + "`" + ` with a maximum performance ratio of 300 IOPS per gibibyte.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The ID of the snapshot used to create the volume.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `The size of the volume, in gibibytes (GiB).<br /> If you specify a snapshot ID, the volume size must be at least equal to the snapshot size.<br /> If you specify a snapshot ID but no volume size, the volume is created with a size similar to the snapshot one.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `The type of the volume (` + "`" + `standard` + "`" + ` \| ` + "`" + `io1` + "`" + ` \| ` + "`" + `gp2` + "`" + `). If not specified in the request, a ` + "`" + `standard` + "`" + ` volume is created.<br /> For more information about volume types, see [About Volumes > Volume Types and IOPS](https://docs.outscale.com/en/userguide/About-Volumes.html#_volume_types_and_iops).`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `The device name for the volume. For a root device, you must use ` + "`" + `/dev/sda1` + "`" + `. For other volumes, you must use ` + "`" + `/dev/sdX` + "`" + `, ` + "`" + `/dev/sdXX` + "`" + `, ` + "`" + `/dev/xvdX` + "`" + `, or ` + "`" + `/dev/xvdXX` + "`" + ` (where the first ` + "`" + `X` + "`" + ` is a letter between ` + "`" + `b` + "`" + ` and ` + "`" + `z` + "`" + `, and the second ` + "`" + `X` + "`" + ` is a letter between ` + "`" + `a` + "`" + ` and ` + "`" + `z` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "virtual_device_name",
					Description: `The name of the virtual device (` + "`" + `ephemeralN` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `The date and time of creation of the OMI.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the OMI.`,
				},
				resource.Attribute{
					Name:        "file_location",
					Description: `The location of the bucket where the OMI files are stored.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The ID of the OMI.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `The name of the OMI.`,
				},
				resource.Attribute{
					Name:        "image_type",
					Description: `The type of the OMI.`,
				},
				resource.Attribute{
					Name:        "permissions_to_launch",
					Description: `Information about the users who have permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "account_ids",
					Description: `The account ID of one or more users who have permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "global_permission",
					Description: `If true, the resource is public. If false, the resource is private.`,
				},
				resource.Attribute{
					Name:        "product_codes",
					Description: `The product code associated with the OMI (` + "`" + `0001` + "`" + ` Linux/Unix \| ` + "`" + `0002` + "`" + ` Windows \| ` + "`" + `0004` + "`" + ` Linux/Oracle \| ` + "`" + `0005` + "`" + ` Windows 10).`,
				},
				resource.Attribute{
					Name:        "root_device_name",
					Description: `The name of the root device.`,
				},
				resource.Attribute{
					Name:        "root_device_type",
					Description: `The type of root device used by the OMI (always ` + "`" + `bsu` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state_comment",
					Description: `Information about the change of state.`,
				},
				resource.Attribute{
					Name:        "state_code",
					Description: `The code of the change of state.`,
				},
				resource.Attribute{
					Name:        "state_message",
					Description: `A message explaining the change of state.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the OMI (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `failed` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the OMI.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "account_alias",
					Description: `The account alias of the owner of the OMI.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the OMI.`,
				},
				resource.Attribute{
					Name:        "architecture",
					Description: `The architecture of the OMI (by default, ` + "`" + `i386` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "block_device_mappings",
					Description: `One or more block device mappings.`,
				},
				resource.Attribute{
					Name:        "bsu",
					Description: `Information about the BSU volume to create.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `By default or if set to true, the volume is deleted when terminating the VM. If false, the volume is not deleted when terminating the VM.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `The number of I/O operations per second (IOPS). This parameter must be specified only if you create an ` + "`" + `io1` + "`" + ` volume. The maximum number of IOPS allowed for ` + "`" + `io1` + "`" + ` volumes is ` + "`" + `13000` + "`" + ` with a maximum performance ratio of 300 IOPS per gibibyte.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The ID of the snapshot used to create the volume.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `The size of the volume, in gibibytes (GiB).<br /> If you specify a snapshot ID, the volume size must be at least equal to the snapshot size.<br /> If you specify a snapshot ID but no volume size, the volume is created with a size similar to the snapshot one.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `The type of the volume (` + "`" + `standard` + "`" + ` \| ` + "`" + `io1` + "`" + ` \| ` + "`" + `gp2` + "`" + `). If not specified in the request, a ` + "`" + `standard` + "`" + ` volume is created.<br /> For more information about volume types, see [About Volumes > Volume Types and IOPS](https://docs.outscale.com/en/userguide/About-Volumes.html#_volume_types_and_iops).`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `The device name for the volume. For a root device, you must use ` + "`" + `/dev/sda1` + "`" + `. For other volumes, you must use ` + "`" + `/dev/sdX` + "`" + `, ` + "`" + `/dev/sdXX` + "`" + `, ` + "`" + `/dev/xvdX` + "`" + `, or ` + "`" + `/dev/xvdXX` + "`" + ` (where the first ` + "`" + `X` + "`" + ` is a letter between ` + "`" + `b` + "`" + ` and ` + "`" + `z` + "`" + `, and the second ` + "`" + `X` + "`" + ` is a letter between ` + "`" + `a` + "`" + ` and ` + "`" + `z` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "virtual_device_name",
					Description: `The name of the virtual device (` + "`" + `ephemeralN` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `The date and time of creation of the OMI.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the OMI.`,
				},
				resource.Attribute{
					Name:        "file_location",
					Description: `The location of the bucket where the OMI files are stored.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The ID of the OMI.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `The name of the OMI.`,
				},
				resource.Attribute{
					Name:        "image_type",
					Description: `The type of the OMI.`,
				},
				resource.Attribute{
					Name:        "permissions_to_launch",
					Description: `Information about the users who have permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "account_ids",
					Description: `The account ID of one or more users who have permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "global_permission",
					Description: `If true, the resource is public. If false, the resource is private.`,
				},
				resource.Attribute{
					Name:        "product_codes",
					Description: `The product code associated with the OMI (` + "`" + `0001` + "`" + ` Linux/Unix \| ` + "`" + `0002` + "`" + ` Windows \| ` + "`" + `0004` + "`" + ` Linux/Oracle \| ` + "`" + `0005` + "`" + ` Windows 10).`,
				},
				resource.Attribute{
					Name:        "root_device_name",
					Description: `The name of the root device.`,
				},
				resource.Attribute{
					Name:        "root_device_type",
					Description: `The type of root device used by the OMI (always ` + "`" + `bsu` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state_comment",
					Description: `Information about the change of state.`,
				},
				resource.Attribute{
					Name:        "state_code",
					Description: `The code of the change of state.`,
				},
				resource.Attribute{
					Name:        "state_message",
					Description: `A message explaining the change of state.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the OMI (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `failed` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the OMI.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_image_export_task",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about an image export task.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "task_ids",
					Description: `(Optional) The IDs of the export tasks. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `If the OMI export task fails, an error message appears.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The ID of the OMI to be exported.`,
				},
				resource.Attribute{
					Name:        "osu_export",
					Description: `Information about the OMI export task.`,
				},
				resource.Attribute{
					Name:        "disk_image_format",
					Description: `The format of the export disk (` + "`" + `qcow2` + "`" + ` \| ` + "`" + `raw` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "osu_bucket",
					Description: `The name of the OOS bucket the OMI is exported to.`,
				},
				resource.Attribute{
					Name:        "osu_manifest_url",
					Description: `The URL of the manifest file.`,
				},
				resource.Attribute{
					Name:        "osu_prefix",
					Description: `The prefix for the key of the OOS object corresponding to the image.`,
				},
				resource.Attribute{
					Name:        "progress",
					Description: `The progress of the OMI export task, as a percentage.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the OMI export task (` + "`" + `pending/queued` + "`" + ` \| ` + "`" + `pending` + "`" + ` \| ` + "`" + `completed` + "`" + ` \| ` + "`" + `failed` + "`" + ` \| ` + "`" + `cancelled` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the image export task.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "task_id",
					Description: `The ID of the OMI export task.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "comment",
					Description: `If the OMI export task fails, an error message appears.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The ID of the OMI to be exported.`,
				},
				resource.Attribute{
					Name:        "osu_export",
					Description: `Information about the OMI export task.`,
				},
				resource.Attribute{
					Name:        "disk_image_format",
					Description: `The format of the export disk (` + "`" + `qcow2` + "`" + ` \| ` + "`" + `raw` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "osu_bucket",
					Description: `The name of the OOS bucket the OMI is exported to.`,
				},
				resource.Attribute{
					Name:        "osu_manifest_url",
					Description: `The URL of the manifest file.`,
				},
				resource.Attribute{
					Name:        "osu_prefix",
					Description: `The prefix for the key of the OOS object corresponding to the image.`,
				},
				resource.Attribute{
					Name:        "progress",
					Description: `The progress of the OMI export task, as a percentage.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the OMI export task (` + "`" + `pending/queued` + "`" + ` \| ` + "`" + `pending` + "`" + ` \| ` + "`" + `completed` + "`" + ` \| ` + "`" + `failed` + "`" + ` \| ` + "`" + `cancelled` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the image export task.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "task_id",
					Description: `The ID of the OMI export task.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_image_export_tasks",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about image export tasks.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "task_ids",
					Description: `(Optional) The IDs of the export tasks. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "image_export_tasks",
					Description: `Information about one or more image export tasks.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `If the OMI export task fails, an error message appears.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The ID of the OMI to be exported.`,
				},
				resource.Attribute{
					Name:        "osu_export",
					Description: `Information about the OMI export task.`,
				},
				resource.Attribute{
					Name:        "disk_image_format",
					Description: `The format of the export disk (` + "`" + `qcow2` + "`" + ` \| ` + "`" + `raw` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "osu_bucket",
					Description: `The name of the OOS bucket the OMI is exported to.`,
				},
				resource.Attribute{
					Name:        "osu_manifest_url",
					Description: `The URL of the manifest file.`,
				},
				resource.Attribute{
					Name:        "osu_prefix",
					Description: `The prefix for the key of the OOS object corresponding to the image.`,
				},
				resource.Attribute{
					Name:        "progress",
					Description: `The progress of the OMI export task, as a percentage.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the OMI export task (` + "`" + `pending/queued` + "`" + ` \| ` + "`" + `pending` + "`" + ` \| ` + "`" + `completed` + "`" + ` \| ` + "`" + `failed` + "`" + ` \| ` + "`" + `cancelled` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the image export task.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "task_id",
					Description: `The ID of the OMI export task.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "image_export_tasks",
					Description: `Information about one or more image export tasks.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `If the OMI export task fails, an error message appears.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The ID of the OMI to be exported.`,
				},
				resource.Attribute{
					Name:        "osu_export",
					Description: `Information about the OMI export task.`,
				},
				resource.Attribute{
					Name:        "disk_image_format",
					Description: `The format of the export disk (` + "`" + `qcow2` + "`" + ` \| ` + "`" + `raw` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "osu_bucket",
					Description: `The name of the OOS bucket the OMI is exported to.`,
				},
				resource.Attribute{
					Name:        "osu_manifest_url",
					Description: `The URL of the manifest file.`,
				},
				resource.Attribute{
					Name:        "osu_prefix",
					Description: `The prefix for the key of the OOS object corresponding to the image.`,
				},
				resource.Attribute{
					Name:        "progress",
					Description: `The progress of the OMI export task, as a percentage.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the OMI export task (` + "`" + `pending/queued` + "`" + ` \| ` + "`" + `pending` + "`" + ` \| ` + "`" + `completed` + "`" + ` \| ` + "`" + `failed` + "`" + ` \| ` + "`" + `cancelled` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the image export task.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "task_id",
					Description: `The ID of the OMI export task.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_images",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about images.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "account_aliases",
					Description: `(Optional) The account aliases of the owners of the OMIs.`,
				},
				resource.Attribute{
					Name:        "account_ids",
					Description: `(Optional) The account IDs of the owners of the OMIs. By default, all the OMIs for which you have launch permissions are described.`,
				},
				resource.Attribute{
					Name:        "architectures",
					Description: `(Optional) The architectures of the OMIs (` + "`" + `i386` + "`" + ` \| ` + "`" + `x86_64` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "block_device_mapping_delete_on_vm_deletion",
					Description: `(Optional) Whether the volumes are deleted or not when terminating the VM.`,
				},
				resource.Attribute{
					Name:        "block_device_mapping_device_names",
					Description: `(Optional) The device names for the volumes.`,
				},
				resource.Attribute{
					Name:        "block_device_mapping_snapshot_ids",
					Description: `(Optional) The IDs of the snapshots used to create the volumes.`,
				},
				resource.Attribute{
					Name:        "block_device_mapping_volume_sizes",
					Description: `(Optional) The sizes of the volumes, in gibibytes (GiB).`,
				},
				resource.Attribute{
					Name:        "block_device_mapping_volume_types",
					Description: `(Optional) The types of volumes (` + "`" + `standard` + "`" + ` \| ` + "`" + `gp2` + "`" + ` \| ` + "`" + `io1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "descriptions",
					Description: `(Optional) The descriptions of the OMIs, provided when they were created.`,
				},
				resource.Attribute{
					Name:        "file_locations",
					Description: `(Optional) The locations of the buckets where the OMI files are stored.`,
				},
				resource.Attribute{
					Name:        "hypervisors",
					Description: `(Optional) The hypervisor type of the OMI (always ` + "`" + `xen` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "image_ids",
					Description: `(Optional) The IDs of the OMIs.`,
				},
				resource.Attribute{
					Name:        "image_names",
					Description: `(Optional) The names of the OMIs, provided when they were created.`,
				},
				resource.Attribute{
					Name:        "permissions_to_launch_account_ids",
					Description: `(Optional) The account IDs of the users who have launch permissions for the OMIs.`,
				},
				resource.Attribute{
					Name:        "permissions_to_launch_global_permission",
					Description: `(Optional) If true, lists all public OMIs. If false, lists all private OMIs.`,
				},
				resource.Attribute{
					Name:        "product_codes",
					Description: `(Optional) The product code associated with the OMI (` + "`" + `0001` + "`" + ` Linux/Unix \| ` + "`" + `0002` + "`" + ` Windows \| ` + "`" + `0004` + "`" + ` Linux/Oracle \| ` + "`" + `0005` + "`" + ` Windows 10).`,
				},
				resource.Attribute{
					Name:        "root_device_names",
					Description: `(Optional) The name of the root device. This value must be /dev/sda1.`,
				},
				resource.Attribute{
					Name:        "root_device_types",
					Description: `(Optional) The types of root device used by the OMIs (always ` + "`" + `bsu` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "states",
					Description: `(Optional) The states of the OMIs (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `failed` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the OMIs.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the OMIs.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combinations of the tags associated with the OMIs, in the following format: ` + "`" + `TAGKEY=TAGVALUE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "virtualization_types",
					Description: `(Optional) The virtualization types (always ` + "`" + `hvm` + "`" + `). ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "images",
					Description: `Information about one or more OMIs.`,
				},
				resource.Attribute{
					Name:        "account_alias",
					Description: `The account alias of the owner of the OMI.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the OMI.`,
				},
				resource.Attribute{
					Name:        "architecture",
					Description: `The architecture of the OMI (by default, ` + "`" + `i386` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "block_device_mappings",
					Description: `One or more block device mappings.`,
				},
				resource.Attribute{
					Name:        "bsu",
					Description: `Information about the BSU volume to create.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `By default or if set to true, the volume is deleted when terminating the VM. If false, the volume is not deleted when terminating the VM.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `The number of I/O operations per second (IOPS). This parameter must be specified only if you create an ` + "`" + `io1` + "`" + ` volume. The maximum number of IOPS allowed for ` + "`" + `io1` + "`" + ` volumes is ` + "`" + `13000` + "`" + ` with a maximum performance ratio of 300 IOPS per gibibyte.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The ID of the snapshot used to create the volume.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `The size of the volume, in gibibytes (GiB).<br /> If you specify a snapshot ID, the volume size must be at least equal to the snapshot size.<br /> If you specify a snapshot ID but no volume size, the volume is created with a size similar to the snapshot one.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `The type of the volume (` + "`" + `standard` + "`" + ` \| ` + "`" + `io1` + "`" + ` \| ` + "`" + `gp2` + "`" + `). If not specified in the request, a ` + "`" + `standard` + "`" + ` volume is created.<br /> For more information about volume types, see [About Volumes > Volume Types and IOPS](https://docs.outscale.com/en/userguide/About-Volumes.html#_volume_types_and_iops).`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `The device name for the volume. For a root device, you must use ` + "`" + `/dev/sda1` + "`" + `. For other volumes, you must use ` + "`" + `/dev/sdX` + "`" + `, ` + "`" + `/dev/sdXX` + "`" + `, ` + "`" + `/dev/xvdX` + "`" + `, or ` + "`" + `/dev/xvdXX` + "`" + ` (where the first ` + "`" + `X` + "`" + ` is a letter between ` + "`" + `b` + "`" + ` and ` + "`" + `z` + "`" + `, and the second ` + "`" + `X` + "`" + ` is a letter between ` + "`" + `a` + "`" + ` and ` + "`" + `z` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "virtual_device_name",
					Description: `The name of the virtual device (` + "`" + `ephemeralN` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `The date and time of creation of the OMI.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the OMI.`,
				},
				resource.Attribute{
					Name:        "file_location",
					Description: `The location of the bucket where the OMI files are stored.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The ID of the OMI.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `The name of the OMI.`,
				},
				resource.Attribute{
					Name:        "image_type",
					Description: `The type of the OMI.`,
				},
				resource.Attribute{
					Name:        "permissions_to_launch",
					Description: `Information about the users who have permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "account_ids",
					Description: `The account ID of one or more users who have permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "global_permission",
					Description: `If true, the resource is public. If false, the resource is private.`,
				},
				resource.Attribute{
					Name:        "product_codes",
					Description: `The product code associated with the OMI (` + "`" + `0001` + "`" + ` Linux/Unix \| ` + "`" + `0002` + "`" + ` Windows \| ` + "`" + `0004` + "`" + ` Linux/Oracle \| ` + "`" + `0005` + "`" + ` Windows 10).`,
				},
				resource.Attribute{
					Name:        "root_device_name",
					Description: `The name of the root device.`,
				},
				resource.Attribute{
					Name:        "root_device_type",
					Description: `The type of root device used by the OMI (always ` + "`" + `bsu` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the OMI (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `failed` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state_comment",
					Description: `Information about the change of state.`,
				},
				resource.Attribute{
					Name:        "state_code",
					Description: `The code of the change of state.`,
				},
				resource.Attribute{
					Name:        "state_message",
					Description: `A message explaining the change of state.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the OMI.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "images",
					Description: `Information about one or more OMIs.`,
				},
				resource.Attribute{
					Name:        "account_alias",
					Description: `The account alias of the owner of the OMI.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the OMI.`,
				},
				resource.Attribute{
					Name:        "architecture",
					Description: `The architecture of the OMI (by default, ` + "`" + `i386` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "block_device_mappings",
					Description: `One or more block device mappings.`,
				},
				resource.Attribute{
					Name:        "bsu",
					Description: `Information about the BSU volume to create.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `By default or if set to true, the volume is deleted when terminating the VM. If false, the volume is not deleted when terminating the VM.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `The number of I/O operations per second (IOPS). This parameter must be specified only if you create an ` + "`" + `io1` + "`" + ` volume. The maximum number of IOPS allowed for ` + "`" + `io1` + "`" + ` volumes is ` + "`" + `13000` + "`" + ` with a maximum performance ratio of 300 IOPS per gibibyte.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The ID of the snapshot used to create the volume.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `The size of the volume, in gibibytes (GiB).<br /> If you specify a snapshot ID, the volume size must be at least equal to the snapshot size.<br /> If you specify a snapshot ID but no volume size, the volume is created with a size similar to the snapshot one.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `The type of the volume (` + "`" + `standard` + "`" + ` \| ` + "`" + `io1` + "`" + ` \| ` + "`" + `gp2` + "`" + `). If not specified in the request, a ` + "`" + `standard` + "`" + ` volume is created.<br /> For more information about volume types, see [About Volumes > Volume Types and IOPS](https://docs.outscale.com/en/userguide/About-Volumes.html#_volume_types_and_iops).`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `The device name for the volume. For a root device, you must use ` + "`" + `/dev/sda1` + "`" + `. For other volumes, you must use ` + "`" + `/dev/sdX` + "`" + `, ` + "`" + `/dev/sdXX` + "`" + `, ` + "`" + `/dev/xvdX` + "`" + `, or ` + "`" + `/dev/xvdXX` + "`" + ` (where the first ` + "`" + `X` + "`" + ` is a letter between ` + "`" + `b` + "`" + ` and ` + "`" + `z` + "`" + `, and the second ` + "`" + `X` + "`" + ` is a letter between ` + "`" + `a` + "`" + ` and ` + "`" + `z` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "virtual_device_name",
					Description: `The name of the virtual device (` + "`" + `ephemeralN` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `The date and time of creation of the OMI.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the OMI.`,
				},
				resource.Attribute{
					Name:        "file_location",
					Description: `The location of the bucket where the OMI files are stored.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The ID of the OMI.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `The name of the OMI.`,
				},
				resource.Attribute{
					Name:        "image_type",
					Description: `The type of the OMI.`,
				},
				resource.Attribute{
					Name:        "permissions_to_launch",
					Description: `Information about the users who have permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "account_ids",
					Description: `The account ID of one or more users who have permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "global_permission",
					Description: `If true, the resource is public. If false, the resource is private.`,
				},
				resource.Attribute{
					Name:        "product_codes",
					Description: `The product code associated with the OMI (` + "`" + `0001` + "`" + ` Linux/Unix \| ` + "`" + `0002` + "`" + ` Windows \| ` + "`" + `0004` + "`" + ` Linux/Oracle \| ` + "`" + `0005` + "`" + ` Windows 10).`,
				},
				resource.Attribute{
					Name:        "root_device_name",
					Description: `The name of the root device.`,
				},
				resource.Attribute{
					Name:        "root_device_type",
					Description: `The type of root device used by the OMI (always ` + "`" + `bsu` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the OMI (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `failed` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state_comment",
					Description: `Information about the change of state.`,
				},
				resource.Attribute{
					Name:        "state_code",
					Description: `The code of the change of state.`,
				},
				resource.Attribute{
					Name:        "state_message",
					Description: `A message explaining the change of state.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the OMI.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_internet_service",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about an Internet service.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "internet_service_ids",
					Description: `(Optional) The IDs of the Internet services.`,
				},
				resource.Attribute{
					Name:        "link_net_ids",
					Description: `(Optional) The IDs of the Nets the Internet services are attached to.`,
				},
				resource.Attribute{
					Name:        "link_states",
					Description: `(Optional) The current states of the attachments between the Internet services and the Nets (only ` + "`" + `available` + "`" + `, if the Internet gateway is attached to a VPC).`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the Internet services.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the Internet services.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combinations of the tags associated with the Internet services, in the following format: ` + "`" + `TAGKEY=TAGVALUE` + "`" + `. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "internet_service_id",
					Description: `The ID of the Internet service.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net attached to the Internet service.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment of the Internet service to the Net (always ` + "`" + `available` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Internet service.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "internet_service_id",
					Description: `The ID of the Internet service.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net attached to the Internet service.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment of the Internet service to the Net (always ` + "`" + `available` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Internet service.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_internet_services",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about Internet services.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "internet_service_ids",
					Description: `(Optional) The IDs of the Internet services.`,
				},
				resource.Attribute{
					Name:        "link_net_ids",
					Description: `(Optional) The IDs of the Nets the Internet services are attached to.`,
				},
				resource.Attribute{
					Name:        "link_states",
					Description: `(Optional) The current states of the attachments between the Internet services and the Nets (only ` + "`" + `available` + "`" + `, if the Internet gateway is attached to a VPC).`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the Internet services.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the Internet services.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combinations of the tags associated with the Internet services, in the following format: ` + "`" + `TAGKEY=TAGVALUE` + "`" + `. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "internet_services",
					Description: `Information about one or more Internet services.`,
				},
				resource.Attribute{
					Name:        "internet_service_id",
					Description: `The ID of the Internet service.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net attached to the Internet service.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment of the Internet service to the Net (always ` + "`" + `available` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Internet service.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "internet_services",
					Description: `Information about one or more Internet services.`,
				},
				resource.Attribute{
					Name:        "internet_service_id",
					Description: `The ID of the Internet service.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net attached to the Internet service.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment of the Internet service to the Net (always ` + "`" + `available` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Internet service.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_keypair",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a keypair.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "keypair_fingerprints",
					Description: `(Optional) The fingerprints of the keypairs.`,
				},
				resource.Attribute{
					Name:        "keypair_names",
					Description: `(Optional) The names of the keypairs. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "keypair_fingerprint",
					Description: `The MD5 public key fingerprint as specified in section 4 of RFC 4716.`,
				},
				resource.Attribute{
					Name:        "keypair_name",
					Description: `The name of the keypair.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "keypair_fingerprint",
					Description: `The MD5 public key fingerprint as specified in section 4 of RFC 4716.`,
				},
				resource.Attribute{
					Name:        "keypair_name",
					Description: `The name of the keypair.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_keypairs",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about keypairs.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "keypair_fingerprints",
					Description: `(Optional) The fingerprints of the keypairs.`,
				},
				resource.Attribute{
					Name:        "keypair_names",
					Description: `(Optional) The names of the keypairs. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "keypairs",
					Description: `Information about one or more keypairs.`,
				},
				resource.Attribute{
					Name:        "keypair_fingerprint",
					Description: `The MD5 public key fingerprint as specified in section 4 of RFC 4716.`,
				},
				resource.Attribute{
					Name:        "keypair_name",
					Description: `The name of the keypair.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "keypairs",
					Description: `Information about one or more keypairs.`,
				},
				resource.Attribute{
					Name:        "keypair_fingerprint",
					Description: `The MD5 public key fingerprint as specified in section 4 of RFC 4716.`,
				},
				resource.Attribute{
					Name:        "keypair_name",
					Description: `The name of the keypair.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_load_balancer",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a load balancer.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "load_balancer_names",
					Description: `(Optional) The names of the load balancers. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "access_log",
					Description: `Information about access logs.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `If true, access logs are enabled for your load balancer. If false, they are not. If you set this to true in your request, the ` + "`" + `osu_bucket_name` + "`" + ` parameter is required.`,
				},
				resource.Attribute{
					Name:        "osu_bucket_name",
					Description: `The name of the OOS bucket for the access logs.`,
				},
				resource.Attribute{
					Name:        "osu_bucket_prefix",
					Description: `The path to the folder of the access logs in your OOS bucket (by default, the ` + "`" + `root` + "`" + ` level of your bucket).`,
				},
				resource.Attribute{
					Name:        "publication_interval",
					Description: `The time interval for the publication of access logs in the OOS bucket, in minutes. This value can be either ` + "`" + `5` + "`" + ` or ` + "`" + `60` + "`" + ` (by default, ` + "`" + `60` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "application_sticky_cookie_policies",
					Description: `The stickiness policies defined for the load balancer.`,
				},
				resource.Attribute{
					Name:        "cookie_name",
					Description: `The name of the application cookie used for stickiness.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `The mnemonic name for the policy being created. The name must be unique within a set of policies for this load balancer.`,
				},
				resource.Attribute{
					Name:        "backend_vm_ids",
					Description: `One or more IDs of back-end VMs for the load balancer.`,
				},
				resource.Attribute{
					Name:        "dns_name",
					Description: `The DNS name of the load balancer.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `Information about the health check configuration.`,
				},
				resource.Attribute{
					Name:        "check_interval",
					Description: `The number of seconds between two pings (between ` + "`" + `5` + "`" + ` and ` + "`" + `600` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `The number of consecutive successful pings before considering the VM as healthy (between ` + "`" + `2` + "`" + ` and ` + "`" + `10` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `If you use the HTTP or HTTPS protocols, the ping path.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port number (between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol for the URL of the VM (` + "`" + `HTTP` + "`" + ` \| ` + "`" + `HTTPS` + "`" + ` \| ` + "`" + `TCP` + "`" + ` \| ` + "`" + `SSL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `The maximum waiting time for a response before considering the VM as unhealthy, in seconds (between ` + "`" + `2` + "`" + ` and ` + "`" + `60` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "unhealthy_threshold",
					Description: `The number of consecutive failed pings before considering the VM as unhealthy (between ` + "`" + `2` + "`" + ` and ` + "`" + `10` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "listeners",
					Description: `The listeners for the load balancer.`,
				},
				resource.Attribute{
					Name:        "backend_port",
					Description: `The port on which the back-end VM is listening (between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "backend_protocol",
					Description: `The protocol for routing traffic to back-end VMs (` + "`" + `HTTP` + "`" + ` \| ` + "`" + `HTTPS` + "`" + ` \| ` + "`" + `TCP` + "`" + ` \| ` + "`" + `SSL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "load_balancer_port",
					Description: `The port on which the load balancer is listening (between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "load_balancer_protocol",
					Description: `The routing protocol (` + "`" + `HTTP` + "`" + ` \| ` + "`" + `HTTPS` + "`" + ` \| ` + "`" + `TCP` + "`" + ` \| ` + "`" + `SSL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "policy_names",
					Description: `The names of the policies. If there are no policies enabled, the list is empty.`,
				},
				resource.Attribute{
					Name:        "server_certificate_id",
					Description: `The OUTSCALE Resource Name (ORN) of the server certificate. For more information, see [Resource Identifiers > OUTSCALE Resource Names (ORNs)](https://docs.outscale.com/en/userguide/Resource-Identifiers.html#_outscale_resource_names_orns).`,
				},
				resource.Attribute{
					Name:        "load_balancer_name",
					Description: `The name of the load balancer.`,
				},
				resource.Attribute{
					Name:        "load_balancer_sticky_cookie_policies",
					Description: `The policies defined for the load balancer.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `The name of the stickiness policy.`,
				},
				resource.Attribute{
					Name:        "load_balancer_type",
					Description: `The type of load balancer. Valid only for load balancers in a Net.<br /> If ` + "`" + `load_balancer_type` + "`" + ` is ` + "`" + `internet-facing` + "`" + `, the load balancer has a public DNS name that resolves to a public IP.<br /> If ` + "`" + `load_balancer_type` + "`" + ` is ` + "`" + `internal` + "`" + `, the load balancer has a public DNS name that resolves to a private IP.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the load balancer.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `(internet-facing only) The public IP associated with the load balancer.`,
				},
				resource.Attribute{
					Name:        "secured_cookies",
					Description: `Whether secure cookies are enabled for the load balancer.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `One or more IDs of security groups for the load balancers. Valid only for load balancers in a Net.`,
				},
				resource.Attribute{
					Name:        "source_security_group",
					Description: `Information about the source security group of the load balancer, which you can use as part of your inbound rules for your registered VMs.<br /> To only allow traffic from load balancers, add a security group rule that specifies this source security group as the inbound source.`,
				},
				resource.Attribute{
					Name:        "security_group_account_id",
					Description: `The account ID of the owner of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `The ID of the Subnet in which the load balancer was created.`,
				},
				resource.Attribute{
					Name:        "subregion_names",
					Description: `The ID of the Subregion in which the load balancer was created.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the load balancer.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "access_log",
					Description: `Information about access logs.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `If true, access logs are enabled for your load balancer. If false, they are not. If you set this to true in your request, the ` + "`" + `osu_bucket_name` + "`" + ` parameter is required.`,
				},
				resource.Attribute{
					Name:        "osu_bucket_name",
					Description: `The name of the OOS bucket for the access logs.`,
				},
				resource.Attribute{
					Name:        "osu_bucket_prefix",
					Description: `The path to the folder of the access logs in your OOS bucket (by default, the ` + "`" + `root` + "`" + ` level of your bucket).`,
				},
				resource.Attribute{
					Name:        "publication_interval",
					Description: `The time interval for the publication of access logs in the OOS bucket, in minutes. This value can be either ` + "`" + `5` + "`" + ` or ` + "`" + `60` + "`" + ` (by default, ` + "`" + `60` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "application_sticky_cookie_policies",
					Description: `The stickiness policies defined for the load balancer.`,
				},
				resource.Attribute{
					Name:        "cookie_name",
					Description: `The name of the application cookie used for stickiness.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `The mnemonic name for the policy being created. The name must be unique within a set of policies for this load balancer.`,
				},
				resource.Attribute{
					Name:        "backend_vm_ids",
					Description: `One or more IDs of back-end VMs for the load balancer.`,
				},
				resource.Attribute{
					Name:        "dns_name",
					Description: `The DNS name of the load balancer.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `Information about the health check configuration.`,
				},
				resource.Attribute{
					Name:        "check_interval",
					Description: `The number of seconds between two pings (between ` + "`" + `5` + "`" + ` and ` + "`" + `600` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `The number of consecutive successful pings before considering the VM as healthy (between ` + "`" + `2` + "`" + ` and ` + "`" + `10` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `If you use the HTTP or HTTPS protocols, the ping path.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port number (between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol for the URL of the VM (` + "`" + `HTTP` + "`" + ` \| ` + "`" + `HTTPS` + "`" + ` \| ` + "`" + `TCP` + "`" + ` \| ` + "`" + `SSL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `The maximum waiting time for a response before considering the VM as unhealthy, in seconds (between ` + "`" + `2` + "`" + ` and ` + "`" + `60` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "unhealthy_threshold",
					Description: `The number of consecutive failed pings before considering the VM as unhealthy (between ` + "`" + `2` + "`" + ` and ` + "`" + `10` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "listeners",
					Description: `The listeners for the load balancer.`,
				},
				resource.Attribute{
					Name:        "backend_port",
					Description: `The port on which the back-end VM is listening (between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "backend_protocol",
					Description: `The protocol for routing traffic to back-end VMs (` + "`" + `HTTP` + "`" + ` \| ` + "`" + `HTTPS` + "`" + ` \| ` + "`" + `TCP` + "`" + ` \| ` + "`" + `SSL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "load_balancer_port",
					Description: `The port on which the load balancer is listening (between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "load_balancer_protocol",
					Description: `The routing protocol (` + "`" + `HTTP` + "`" + ` \| ` + "`" + `HTTPS` + "`" + ` \| ` + "`" + `TCP` + "`" + ` \| ` + "`" + `SSL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "policy_names",
					Description: `The names of the policies. If there are no policies enabled, the list is empty.`,
				},
				resource.Attribute{
					Name:        "server_certificate_id",
					Description: `The OUTSCALE Resource Name (ORN) of the server certificate. For more information, see [Resource Identifiers > OUTSCALE Resource Names (ORNs)](https://docs.outscale.com/en/userguide/Resource-Identifiers.html#_outscale_resource_names_orns).`,
				},
				resource.Attribute{
					Name:        "load_balancer_name",
					Description: `The name of the load balancer.`,
				},
				resource.Attribute{
					Name:        "load_balancer_sticky_cookie_policies",
					Description: `The policies defined for the load balancer.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `The name of the stickiness policy.`,
				},
				resource.Attribute{
					Name:        "load_balancer_type",
					Description: `The type of load balancer. Valid only for load balancers in a Net.<br /> If ` + "`" + `load_balancer_type` + "`" + ` is ` + "`" + `internet-facing` + "`" + `, the load balancer has a public DNS name that resolves to a public IP.<br /> If ` + "`" + `load_balancer_type` + "`" + ` is ` + "`" + `internal` + "`" + `, the load balancer has a public DNS name that resolves to a private IP.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the load balancer.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `(internet-facing only) The public IP associated with the load balancer.`,
				},
				resource.Attribute{
					Name:        "secured_cookies",
					Description: `Whether secure cookies are enabled for the load balancer.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `One or more IDs of security groups for the load balancers. Valid only for load balancers in a Net.`,
				},
				resource.Attribute{
					Name:        "source_security_group",
					Description: `Information about the source security group of the load balancer, which you can use as part of your inbound rules for your registered VMs.<br /> To only allow traffic from load balancers, add a security group rule that specifies this source security group as the inbound source.`,
				},
				resource.Attribute{
					Name:        "security_group_account_id",
					Description: `The account ID of the owner of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `The ID of the Subnet in which the load balancer was created.`,
				},
				resource.Attribute{
					Name:        "subregion_names",
					Description: `The ID of the Subregion in which the load balancer was created.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the load balancer.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_load_balancer_listener_rule",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a load balancer listener rule.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "listener_rule_names",
					Description: `(Optional) The names of the listener rules. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `The type of action for the rule (always ` + "`" + `forward` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "host_name_pattern",
					Description: `A host-name pattern for the rule, with a maximum length of 128 characters. This host-name pattern supports maximum three wildcards, and must not contain any special characters except [-.?].`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `The ID of the listener.`,
				},
				resource.Attribute{
					Name:        "listener_rule_id",
					Description: `The ID of the listener rule.`,
				},
				resource.Attribute{
					Name:        "listener_rule_name",
					Description: `A human-readable name for the listener rule.`,
				},
				resource.Attribute{
					Name:        "path_pattern",
					Description: `A path pattern for the rule, with a maximum length of 128 characters. This path pattern supports maximum three wildcards, and must not contain any special characters except [_-.$/~&quot;'@:+?].`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority level of the listener rule, between ` + "`" + `1` + "`" + ` and ` + "`" + `19999` + "`" + ` both included. Each rule must have a unique priority level. Otherwise, an error is returned.`,
				},
				resource.Attribute{
					Name:        "vm_ids",
					Description: `The IDs of the backend VMs.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "action",
					Description: `The type of action for the rule (always ` + "`" + `forward` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "host_name_pattern",
					Description: `A host-name pattern for the rule, with a maximum length of 128 characters. This host-name pattern supports maximum three wildcards, and must not contain any special characters except [-.?].`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `The ID of the listener.`,
				},
				resource.Attribute{
					Name:        "listener_rule_id",
					Description: `The ID of the listener rule.`,
				},
				resource.Attribute{
					Name:        "listener_rule_name",
					Description: `A human-readable name for the listener rule.`,
				},
				resource.Attribute{
					Name:        "path_pattern",
					Description: `A path pattern for the rule, with a maximum length of 128 characters. This path pattern supports maximum three wildcards, and must not contain any special characters except [_-.$/~&quot;'@:+?].`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority level of the listener rule, between ` + "`" + `1` + "`" + ` and ` + "`" + `19999` + "`" + ` both included. Each rule must have a unique priority level. Otherwise, an error is returned.`,
				},
				resource.Attribute{
					Name:        "vm_ids",
					Description: `The IDs of the backend VMs.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_load_balancer_listener_rules",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about load balancer listener rules.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "listener_rule_names",
					Description: `(Optional) The names of the listener rules. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "listener_rules",
					Description: `The list of the rules to describe.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `The type of action for the rule (always ` + "`" + `forward` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "host_name_pattern",
					Description: `A host-name pattern for the rule, with a maximum length of 128 characters. This host-name pattern supports maximum three wildcards, and must not contain any special characters except [-.?].`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `The ID of the listener.`,
				},
				resource.Attribute{
					Name:        "listener_rule_id",
					Description: `The ID of the listener rule.`,
				},
				resource.Attribute{
					Name:        "listener_rule_name",
					Description: `A human-readable name for the listener rule.`,
				},
				resource.Attribute{
					Name:        "path_pattern",
					Description: `A path pattern for the rule, with a maximum length of 128 characters. This path pattern supports maximum three wildcards, and must not contain any special characters except [_-.$/~&quot;'@:+?].`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority level of the listener rule, between ` + "`" + `1` + "`" + ` and ` + "`" + `19999` + "`" + ` both included. Each rule must have a unique priority level. Otherwise, an error is returned.`,
				},
				resource.Attribute{
					Name:        "vm_ids",
					Description: `The IDs of the backend VMs.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "listener_rules",
					Description: `The list of the rules to describe.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `The type of action for the rule (always ` + "`" + `forward` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "host_name_pattern",
					Description: `A host-name pattern for the rule, with a maximum length of 128 characters. This host-name pattern supports maximum three wildcards, and must not contain any special characters except [-.?].`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `The ID of the listener.`,
				},
				resource.Attribute{
					Name:        "listener_rule_id",
					Description: `The ID of the listener rule.`,
				},
				resource.Attribute{
					Name:        "listener_rule_name",
					Description: `A human-readable name for the listener rule.`,
				},
				resource.Attribute{
					Name:        "path_pattern",
					Description: `A path pattern for the rule, with a maximum length of 128 characters. This path pattern supports maximum three wildcards, and must not contain any special characters except [_-.$/~&quot;'@:+?].`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority level of the listener rule, between ` + "`" + `1` + "`" + ` and ` + "`" + `19999` + "`" + ` both included. Each rule must have a unique priority level. Otherwise, an error is returned.`,
				},
				resource.Attribute{
					Name:        "vm_ids",
					Description: `The IDs of the backend VMs.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_load_balancer_vm_health",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about the health of one or more back-end VMs registered with a specific load balancer.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "backend_vm_ids",
					Description: `(Optional) One or more IDs of back-end VMs.`,
				},
				resource.Attribute{
					Name:        "load_balancer_name",
					Description: `(Required) The name of the load balancer. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "backend_vm_health",
					Description: `Information about the health of one or more back-end VMs.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the state of the back-end VM.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the back-end VM (` + "`" + `InService` + "`" + ` \| ` + "`" + `OutOfService` + "`" + ` \| ` + "`" + `Unknown` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state_reason",
					Description: `Information about the cause of ` + "`" + `OutOfService` + "`" + ` VMs.<br /> Specifically, whether the cause is Elastic Load Balancing or the VM (` + "`" + `ELB` + "`" + ` \| ` + "`" + `Instance` + "`" + ` \| ` + "`" + `N/A` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the back-end VM.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "backend_vm_health",
					Description: `Information about the health of one or more back-end VMs.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the state of the back-end VM.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the back-end VM (` + "`" + `InService` + "`" + ` \| ` + "`" + `OutOfService` + "`" + ` \| ` + "`" + `Unknown` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state_reason",
					Description: `Information about the cause of ` + "`" + `OutOfService` + "`" + ` VMs.<br /> Specifically, whether the cause is Elastic Load Balancing or the VM (` + "`" + `ELB` + "`" + ` \| ` + "`" + `Instance` + "`" + ` \| ` + "`" + `N/A` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the back-end VM.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_load_balancers",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about load balancers.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "load_balancer_names",
					Description: `(Optional) The names of the load balancers. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "load_balancers",
					Description: `Information about one or more load balancers.`,
				},
				resource.Attribute{
					Name:        "access_log",
					Description: `Information about access logs.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `If true, access logs are enabled for your load balancer. If false, they are not. If you set this to true in your request, the ` + "`" + `osu_bucket_name` + "`" + ` parameter is required.`,
				},
				resource.Attribute{
					Name:        "osu_bucket_name",
					Description: `The name of the OOS bucket for the access logs.`,
				},
				resource.Attribute{
					Name:        "osu_bucket_prefix",
					Description: `The path to the folder of the access logs in your OOS bucket (by default, the ` + "`" + `root` + "`" + ` level of your bucket).`,
				},
				resource.Attribute{
					Name:        "publication_interval",
					Description: `The time interval for the publication of access logs in the OOS bucket, in minutes. This value can be either ` + "`" + `5` + "`" + ` or ` + "`" + `60` + "`" + ` (by default, ` + "`" + `60` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "application_sticky_cookie_policies",
					Description: `The stickiness policies defined for the load balancer.`,
				},
				resource.Attribute{
					Name:        "cookie_name",
					Description: `The name of the application cookie used for stickiness.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `The mnemonic name for the policy being created. The name must be unique within a set of policies for this load balancer.`,
				},
				resource.Attribute{
					Name:        "backend_vm_ids",
					Description: `One or more IDs of back-end VMs for the load balancer.`,
				},
				resource.Attribute{
					Name:        "dns_name",
					Description: `The DNS name of the load balancer.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `Information about the health check configuration.`,
				},
				resource.Attribute{
					Name:        "check_interval",
					Description: `The number of seconds between two pings (between ` + "`" + `5` + "`" + ` and ` + "`" + `600` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `The number of consecutive successful pings before considering the VM as healthy (between ` + "`" + `2` + "`" + ` and ` + "`" + `10` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `If you use the HTTP or HTTPS protocols, the ping path.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port number (between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol for the URL of the VM (` + "`" + `HTTP` + "`" + ` \| ` + "`" + `HTTPS` + "`" + ` \| ` + "`" + `TCP` + "`" + ` \| ` + "`" + `SSL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `The maximum waiting time for a response before considering the VM as unhealthy, in seconds (between ` + "`" + `2` + "`" + ` and ` + "`" + `60` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "unhealthy_threshold",
					Description: `The number of consecutive failed pings before considering the VM as unhealthy (between ` + "`" + `2` + "`" + ` and ` + "`" + `10` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "listeners",
					Description: `The listeners for the load balancer.`,
				},
				resource.Attribute{
					Name:        "backend_port",
					Description: `The port on which the back-end VM is listening (between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "backend_protocol",
					Description: `The protocol for routing traffic to back-end VMs (` + "`" + `HTTP` + "`" + ` \| ` + "`" + `HTTPS` + "`" + ` \| ` + "`" + `TCP` + "`" + ` \| ` + "`" + `SSL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "load_balancer_port",
					Description: `The port on which the load balancer is listening (between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "load_balancer_protocol",
					Description: `The routing protocol (` + "`" + `HTTP` + "`" + ` \| ` + "`" + `HTTPS` + "`" + ` \| ` + "`" + `TCP` + "`" + ` \| ` + "`" + `SSL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "policy_names",
					Description: `The names of the policies. If there are no policies enabled, the list is empty.`,
				},
				resource.Attribute{
					Name:        "server_certificate_id",
					Description: `The OUTSCALE Resource Name (ORN) of the server certificate. For more information, see [Resource Identifiers > OUTSCALE Resource Names (ORNs)](https://docs.outscale.com/en/userguide/Resource-Identifiers.html#_outscale_resource_names_orns).`,
				},
				resource.Attribute{
					Name:        "load_balancer_name",
					Description: `The name of the load balancer.`,
				},
				resource.Attribute{
					Name:        "load_balancer_sticky_cookie_policies",
					Description: `The policies defined for the load balancer.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `The name of the stickiness policy.`,
				},
				resource.Attribute{
					Name:        "load_balancer_type",
					Description: `The type of load balancer. Valid only for load balancers in a Net.<br /> If ` + "`" + `load_balancer_type` + "`" + ` is ` + "`" + `internet-facing` + "`" + `, the load balancer has a public DNS name that resolves to a public IP.<br /> If ` + "`" + `load_balancer_type` + "`" + ` is ` + "`" + `internal` + "`" + `, the load balancer has a public DNS name that resolves to a private IP.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the load balancer.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `(internet-facing only) The public IP associated with the load balancer.`,
				},
				resource.Attribute{
					Name:        "secured_cookies",
					Description: `Whether secure cookies are enabled for the load balancer.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `One or more IDs of security groups for the load balancers. Valid only for load balancers in a Net.`,
				},
				resource.Attribute{
					Name:        "source_security_group",
					Description: `Information about the source security group of the load balancer, which you can use as part of your inbound rules for your registered VMs.<br /> To only allow traffic from load balancers, add a security group rule that specifies this source security group as the inbound source.`,
				},
				resource.Attribute{
					Name:        "security_group_account_id",
					Description: `The account ID of the owner of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `The ID of the Subnet in which the load balancer was created.`,
				},
				resource.Attribute{
					Name:        "subregion_names",
					Description: `The ID of the Subregion in which the load balancer was created.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the load balancer.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "load_balancers",
					Description: `Information about one or more load balancers.`,
				},
				resource.Attribute{
					Name:        "access_log",
					Description: `Information about access logs.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `If true, access logs are enabled for your load balancer. If false, they are not. If you set this to true in your request, the ` + "`" + `osu_bucket_name` + "`" + ` parameter is required.`,
				},
				resource.Attribute{
					Name:        "osu_bucket_name",
					Description: `The name of the OOS bucket for the access logs.`,
				},
				resource.Attribute{
					Name:        "osu_bucket_prefix",
					Description: `The path to the folder of the access logs in your OOS bucket (by default, the ` + "`" + `root` + "`" + ` level of your bucket).`,
				},
				resource.Attribute{
					Name:        "publication_interval",
					Description: `The time interval for the publication of access logs in the OOS bucket, in minutes. This value can be either ` + "`" + `5` + "`" + ` or ` + "`" + `60` + "`" + ` (by default, ` + "`" + `60` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "application_sticky_cookie_policies",
					Description: `The stickiness policies defined for the load balancer.`,
				},
				resource.Attribute{
					Name:        "cookie_name",
					Description: `The name of the application cookie used for stickiness.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `The mnemonic name for the policy being created. The name must be unique within a set of policies for this load balancer.`,
				},
				resource.Attribute{
					Name:        "backend_vm_ids",
					Description: `One or more IDs of back-end VMs for the load balancer.`,
				},
				resource.Attribute{
					Name:        "dns_name",
					Description: `The DNS name of the load balancer.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `Information about the health check configuration.`,
				},
				resource.Attribute{
					Name:        "check_interval",
					Description: `The number of seconds between two pings (between ` + "`" + `5` + "`" + ` and ` + "`" + `600` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `The number of consecutive successful pings before considering the VM as healthy (between ` + "`" + `2` + "`" + ` and ` + "`" + `10` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `If you use the HTTP or HTTPS protocols, the ping path.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port number (between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol for the URL of the VM (` + "`" + `HTTP` + "`" + ` \| ` + "`" + `HTTPS` + "`" + ` \| ` + "`" + `TCP` + "`" + ` \| ` + "`" + `SSL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `The maximum waiting time for a response before considering the VM as unhealthy, in seconds (between ` + "`" + `2` + "`" + ` and ` + "`" + `60` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "unhealthy_threshold",
					Description: `The number of consecutive failed pings before considering the VM as unhealthy (between ` + "`" + `2` + "`" + ` and ` + "`" + `10` + "`" + ` both included).`,
				},
				resource.Attribute{
					Name:        "listeners",
					Description: `The listeners for the load balancer.`,
				},
				resource.Attribute{
					Name:        "backend_port",
					Description: `The port on which the back-end VM is listening (between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "backend_protocol",
					Description: `The protocol for routing traffic to back-end VMs (` + "`" + `HTTP` + "`" + ` \| ` + "`" + `HTTPS` + "`" + ` \| ` + "`" + `TCP` + "`" + ` \| ` + "`" + `SSL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "load_balancer_port",
					Description: `The port on which the load balancer is listening (between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "load_balancer_protocol",
					Description: `The routing protocol (` + "`" + `HTTP` + "`" + ` \| ` + "`" + `HTTPS` + "`" + ` \| ` + "`" + `TCP` + "`" + ` \| ` + "`" + `SSL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "policy_names",
					Description: `The names of the policies. If there are no policies enabled, the list is empty.`,
				},
				resource.Attribute{
					Name:        "server_certificate_id",
					Description: `The OUTSCALE Resource Name (ORN) of the server certificate. For more information, see [Resource Identifiers > OUTSCALE Resource Names (ORNs)](https://docs.outscale.com/en/userguide/Resource-Identifiers.html#_outscale_resource_names_orns).`,
				},
				resource.Attribute{
					Name:        "load_balancer_name",
					Description: `The name of the load balancer.`,
				},
				resource.Attribute{
					Name:        "load_balancer_sticky_cookie_policies",
					Description: `The policies defined for the load balancer.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `The name of the stickiness policy.`,
				},
				resource.Attribute{
					Name:        "load_balancer_type",
					Description: `The type of load balancer. Valid only for load balancers in a Net.<br /> If ` + "`" + `load_balancer_type` + "`" + ` is ` + "`" + `internet-facing` + "`" + `, the load balancer has a public DNS name that resolves to a public IP.<br /> If ` + "`" + `load_balancer_type` + "`" + ` is ` + "`" + `internal` + "`" + `, the load balancer has a public DNS name that resolves to a private IP.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the load balancer.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `(internet-facing only) The public IP associated with the load balancer.`,
				},
				resource.Attribute{
					Name:        "secured_cookies",
					Description: `Whether secure cookies are enabled for the load balancer.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `One or more IDs of security groups for the load balancers. Valid only for load balancers in a Net.`,
				},
				resource.Attribute{
					Name:        "source_security_group",
					Description: `Information about the source security group of the load balancer, which you can use as part of your inbound rules for your registered VMs.<br /> To only allow traffic from load balancers, add a security group rule that specifies this source security group as the inbound source.`,
				},
				resource.Attribute{
					Name:        "security_group_account_id",
					Description: `The account ID of the owner of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `The ID of the Subnet in which the load balancer was created.`,
				},
				resource.Attribute{
					Name:        "subregion_names",
					Description: `The ID of the Subregion in which the load balancer was created.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the load balancer.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_nat_service",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a NAT service.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "nat_service_ids",
					Description: `(Optional) The IDs of the NAT services.`,
				},
				resource.Attribute{
					Name:        "net_ids",
					Description: `(Optional) The IDs of the Nets in which the NAT services are.`,
				},
				resource.Attribute{
					Name:        "states",
					Description: `(Optional) The states of the NAT services (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `(Optional) The IDs of the Subnets in which the NAT services are.`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the NAT services.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the NAT services.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combinations of the tags associated with the NAT services, in the following format: ` + "`" + `TAGKEY=TAGVALUE` + "`" + `. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "nat_service_id",
					Description: `The ID of the NAT service.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net in which the NAT service is.`,
				},
				resource.Attribute{
					Name:        "public_ips",
					Description: `Information about the public IP or IPs associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `The allocation ID of the public IP associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the NAT service (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet in which the NAT service is.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "nat_service_id",
					Description: `The ID of the NAT service.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net in which the NAT service is.`,
				},
				resource.Attribute{
					Name:        "public_ips",
					Description: `Information about the public IP or IPs associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `The allocation ID of the public IP associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the NAT service (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet in which the NAT service is.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_nat_services",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about NAT services.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "nat_service_ids",
					Description: `(Optional) The IDs of the NAT services.`,
				},
				resource.Attribute{
					Name:        "net_ids",
					Description: `(Optional) The IDs of the Nets in which the NAT services are.`,
				},
				resource.Attribute{
					Name:        "states",
					Description: `(Optional) The states of the NAT services (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `(Optional) The IDs of the Subnets in which the NAT services are.`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the NAT services.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the NAT services.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combinations of the tags associated with the NAT services, in the following format: ` + "`" + `TAGKEY=TAGVALUE` + "`" + `. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "nat_services",
					Description: `Information about one or more NAT services.`,
				},
				resource.Attribute{
					Name:        "nat_service_id",
					Description: `The ID of the NAT service.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net in which the NAT service is.`,
				},
				resource.Attribute{
					Name:        "public_ips",
					Description: `Information about the public IP or IPs associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `The allocation ID of the public IP associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the NAT service (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet in which the NAT service is.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "nat_services",
					Description: `Information about one or more NAT services.`,
				},
				resource.Attribute{
					Name:        "nat_service_id",
					Description: `The ID of the NAT service.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net in which the NAT service is.`,
				},
				resource.Attribute{
					Name:        "public_ips",
					Description: `Information about the public IP or IPs associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `The allocation ID of the public IP associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the NAT service (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet in which the NAT service is.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the NAT service.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_net",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a Net.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "dhcp_options_set_ids",
					Description: `(Optional) The IDs of the DHCP options sets.`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `(Optional) The IP ranges for the Nets, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "net_ids",
					Description: `(Optional) The IDs of the Nets.`,
				},
				resource.Attribute{
					Name:        "states",
					Description: `(Optional) The states of the Nets (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the Nets.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the Nets.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combinations of the tags associated with the Nets, in the following format: ` + "`" + `TAGKEY=TAGVALUE` + "`" + `. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "dhcp_options_set_id",
					Description: `The ID of the DHCP options set (or ` + "`" + `default` + "`" + ` if you want to associate the default one).`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range for the Net, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Net (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Net.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "tenancy",
					Description: `The VM tenancy in a Net.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "dhcp_options_set_id",
					Description: `The ID of the DHCP options set (or ` + "`" + `default` + "`" + ` if you want to associate the default one).`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range for the Net, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Net (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Net.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "tenancy",
					Description: `The VM tenancy in a Net.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_net_access_point",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a Net access point.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "net_access_point_ids",
					Description: `(Optional) The IDs of the Net access points.`,
				},
				resource.Attribute{
					Name:        "net_ids",
					Description: `(Optional) The IDs of the Nets.`,
				},
				resource.Attribute{
					Name:        "service_names",
					Description: `(Optional) The names of the services. For more information, see [ReadNetAccessPointServices](https://docs.outscale.com/api#readnetaccesspointservices).`,
				},
				resource.Attribute{
					Name:        "states",
					Description: `(Optional) The states of the Net access points (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the Net access points.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the Net access points.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combinations of the tags associated with the Net access points, in the following format: ` + "`" + `TAGKEY=TAGVALUE` + "`" + `. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "net_access_point_id",
					Description: `The ID of the Net access point.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net with which the Net access point is associated.`,
				},
				resource.Attribute{
					Name:        "route_table_ids",
					Description: `The ID of the route tables associated with the Net access point.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `The name of the service with which the Net access point is associated.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Net access point (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Net access point.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "net_access_point_id",
					Description: `The ID of the Net access point.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net with which the Net access point is associated.`,
				},
				resource.Attribute{
					Name:        "route_table_ids",
					Description: `The ID of the route tables associated with the Net access point.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `The name of the service with which the Net access point is associated.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Net access point (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Net access point.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_net_access_point_services",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about Net access point services.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "service_ids",
					Description: `(Optional) The IDs of the services.`,
				},
				resource.Attribute{
					Name:        "service_names",
					Description: `(Optional) The names of the services. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `The names of the services you can use for Net access points.`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `The list of network prefixes used by the service, in CIDR notation.`,
				},
				resource.Attribute{
					Name:        "service_id",
					Description: `The ID of the service.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `The name of the service.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "services",
					Description: `The names of the services you can use for Net access points.`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `The list of network prefixes used by the service, in CIDR notation.`,
				},
				resource.Attribute{
					Name:        "service_id",
					Description: `The ID of the service.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `The name of the service.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_net_access_points",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about Net access points.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "net_access_point_ids",
					Description: `(Optional) The IDs of the Net access points.`,
				},
				resource.Attribute{
					Name:        "net_ids",
					Description: `(Optional) The IDs of the Nets.`,
				},
				resource.Attribute{
					Name:        "service_names",
					Description: `(Optional) The names of the services. For more information, see [ReadNetAccessPointServices](https://docs.outscale.com/api#readnetaccesspointservices).`,
				},
				resource.Attribute{
					Name:        "states",
					Description: `(Optional) The states of the Net access points (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the Net access points.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the Net access points.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combinations of the tags associated with the Net access points, in the following format: ` + "`" + `TAGKEY=TAGVALUE` + "`" + `. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "net_access_points",
					Description: `One or more Net access points.`,
				},
				resource.Attribute{
					Name:        "net_access_point_id",
					Description: `The ID of the Net access point.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net with which the Net access point is associated.`,
				},
				resource.Attribute{
					Name:        "route_table_ids",
					Description: `The ID of the route tables associated with the Net access point.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `The name of the service with which the Net access point is associated.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Net access point (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Net access point.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "net_access_points",
					Description: `One or more Net access points.`,
				},
				resource.Attribute{
					Name:        "net_access_point_id",
					Description: `The ID of the Net access point.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net with which the Net access point is associated.`,
				},
				resource.Attribute{
					Name:        "route_table_ids",
					Description: `The ID of the route tables associated with the Net access point.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `The name of the service with which the Net access point is associated.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Net access point (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Net access point.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_net_attributes",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about the attributes of a Net.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "net_id",
					Description: `(Optional) The ID of the Net. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "dhcp_options_set_id",
					Description: `The ID of the DHCP options set (or ` + "`" + `default` + "`" + ` if you want to associate the default one).`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range for the Net, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Net (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Net.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "tenancy",
					Description: `The VM tenancy in a Net.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "dhcp_options_set_id",
					Description: `The ID of the DHCP options set (or ` + "`" + `default` + "`" + ` if you want to associate the default one).`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range for the Net, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Net (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Net.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "tenancy",
					Description: `The VM tenancy in a Net.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_net_peering",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a Net peering.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "accepter_net_account_ids",
					Description: `(Optional) The account IDs of the owners of the peer Nets.`,
				},
				resource.Attribute{
					Name:        "accepter_net_ip_ranges",
					Description: `(Optional) The IP ranges of the peer Nets, in CIDR notation (for example, ` + "`" + `10.0.0.0/24` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "accepter_net_net_ids",
					Description: `(Optional) The IDs of the peer Nets.`,
				},
				resource.Attribute{
					Name:        "net_peering_ids",
					Description: `(Optional) The IDs of the Net peerings.`,
				},
				resource.Attribute{
					Name:        "source_net_account_ids",
					Description: `(Optional) The account IDs of the owners of the peer Nets.`,
				},
				resource.Attribute{
					Name:        "source_net_ip_ranges",
					Description: `(Optional) The IP ranges of the peer Nets.`,
				},
				resource.Attribute{
					Name:        "source_net_net_ids",
					Description: `(Optional) The IDs of the peer Nets.`,
				},
				resource.Attribute{
					Name:        "state_messages",
					Description: `(Optional) Additional information about the states of the Net peerings.`,
				},
				resource.Attribute{
					Name:        "state_names",
					Description: `(Optional) The states of the Net peerings (` + "`" + `pending-acceptance` + "`" + ` \| ` + "`" + `active` + "`" + ` \| ` + "`" + `rejected` + "`" + ` \| ` + "`" + `failed` + "`" + ` \| ` + "`" + `expired` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the Net peerings.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the Net peerings.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combinations of the tags associated with the Net peerings, in the following format: ` + "`" + `TAGKEY=TAGVALUE` + "`" + `. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "accepter_net",
					Description: `Information about the accepter Net.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the accepter Net.`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range for the accepter Net, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the accepter Net.`,
				},
				resource.Attribute{
					Name:        "net_peering_id",
					Description: `The ID of the Net peering.`,
				},
				resource.Attribute{
					Name:        "source_net",
					Description: `Information about the source Net.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the source Net.`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range for the source Net, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the source Net.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Information about the state of the Net peering.`,
				},
				resource.Attribute{
					Name:        "message",
					Description: `Additional information about the state of the Net peering.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The state of the Net peering (` + "`" + `pending-acceptance` + "`" + ` \| ` + "`" + `active` + "`" + ` \| ` + "`" + `rejected` + "`" + ` \| ` + "`" + `failed` + "`" + ` \| ` + "`" + `expired` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Net peering.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "accepter_net",
					Description: `Information about the accepter Net.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the accepter Net.`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range for the accepter Net, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the accepter Net.`,
				},
				resource.Attribute{
					Name:        "net_peering_id",
					Description: `The ID of the Net peering.`,
				},
				resource.Attribute{
					Name:        "source_net",
					Description: `Information about the source Net.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the source Net.`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range for the source Net, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the source Net.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Information about the state of the Net peering.`,
				},
				resource.Attribute{
					Name:        "message",
					Description: `Additional information about the state of the Net peering.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The state of the Net peering (` + "`" + `pending-acceptance` + "`" + ` \| ` + "`" + `active` + "`" + ` \| ` + "`" + `rejected` + "`" + ` \| ` + "`" + `failed` + "`" + ` \| ` + "`" + `expired` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Net peering.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_net_peerings",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about Net peerings.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "accepter_net_account_ids",
					Description: `(Optional) The account IDs of the owners of the peer Nets.`,
				},
				resource.Attribute{
					Name:        "accepter_net_ip_ranges",
					Description: `(Optional) The IP ranges of the peer Nets, in CIDR notation (for example, ` + "`" + `10.0.0.0/24` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "accepter_net_net_ids",
					Description: `(Optional) The IDs of the peer Nets.`,
				},
				resource.Attribute{
					Name:        "net_peering_ids",
					Description: `(Optional) The IDs of the Net peerings.`,
				},
				resource.Attribute{
					Name:        "source_net_account_ids",
					Description: `(Optional) The account IDs of the owners of the peer Nets.`,
				},
				resource.Attribute{
					Name:        "source_net_ip_ranges",
					Description: `(Optional) The IP ranges of the peer Nets.`,
				},
				resource.Attribute{
					Name:        "source_net_net_ids",
					Description: `(Optional) The IDs of the peer Nets.`,
				},
				resource.Attribute{
					Name:        "state_messages",
					Description: `(Optional) Additional information about the states of the Net peerings.`,
				},
				resource.Attribute{
					Name:        "state_names",
					Description: `(Optional) The states of the Net peerings (` + "`" + `pending-acceptance` + "`" + ` \| ` + "`" + `active` + "`" + ` \| ` + "`" + `rejected` + "`" + ` \| ` + "`" + `failed` + "`" + ` \| ` + "`" + `expired` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the Net peerings.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the Net peerings.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combinations of the tags associated with the Net peerings, in the following format: ` + "`" + `TAGKEY=TAGVALUE` + "`" + `. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "net_peerings",
					Description: `Information about one or more Net peerings.`,
				},
				resource.Attribute{
					Name:        "accepter_net",
					Description: `Information about the accepter Net.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the accepter Net.`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range for the accepter Net, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the accepter Net.`,
				},
				resource.Attribute{
					Name:        "net_peering_id",
					Description: `The ID of the Net peering.`,
				},
				resource.Attribute{
					Name:        "source_net",
					Description: `Information about the source Net.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the source Net.`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range for the source Net, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the source Net.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Information about the state of the Net peering.`,
				},
				resource.Attribute{
					Name:        "message",
					Description: `Additional information about the state of the Net peering.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The state of the Net peering (` + "`" + `pending-acceptance` + "`" + ` \| ` + "`" + `active` + "`" + ` \| ` + "`" + `rejected` + "`" + ` \| ` + "`" + `failed` + "`" + ` \| ` + "`" + `expired` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Net peering.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "net_peerings",
					Description: `Information about one or more Net peerings.`,
				},
				resource.Attribute{
					Name:        "accepter_net",
					Description: `Information about the accepter Net.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the accepter Net.`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range for the accepter Net, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the accepter Net.`,
				},
				resource.Attribute{
					Name:        "net_peering_id",
					Description: `The ID of the Net peering.`,
				},
				resource.Attribute{
					Name:        "source_net",
					Description: `Information about the source Net.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the source Net.`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range for the source Net, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the source Net.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Information about the state of the Net peering.`,
				},
				resource.Attribute{
					Name:        "message",
					Description: `Additional information about the state of the Net peering.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The state of the Net peering (` + "`" + `pending-acceptance` + "`" + ` \| ` + "`" + `active` + "`" + ` \| ` + "`" + `rejected` + "`" + ` \| ` + "`" + `failed` + "`" + ` \| ` + "`" + `expired` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Net peering.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_nets",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about Nets.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "dhcp_options_set_ids",
					Description: `(Optional) The IDs of the DHCP options sets.`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `(Optional) The IP ranges for the Nets, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "net_ids",
					Description: `(Optional) The IDs of the Nets.`,
				},
				resource.Attribute{
					Name:        "states",
					Description: `(Optional) The states of the Nets (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the Nets.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the Nets.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combinations of the tags associated with the Nets, in the following format: ` + "`" + `TAGKEY=TAGVALUE` + "`" + `. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "nets",
					Description: `Information about the described Nets.`,
				},
				resource.Attribute{
					Name:        "dhcp_options_set_id",
					Description: `The ID of the DHCP options set (or ` + "`" + `default` + "`" + ` if you want to associate the default one).`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range for the Net, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Net (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Net.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "tenancy",
					Description: `The VM tenancy in a Net.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "nets",
					Description: `Information about the described Nets.`,
				},
				resource.Attribute{
					Name:        "dhcp_options_set_id",
					Description: `The ID of the DHCP options set (or ` + "`" + `default` + "`" + ` if you want to associate the default one).`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range for the Net, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Net (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Net.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "tenancy",
					Description: `The VM tenancy in a Net.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_nic",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a network interface card (NIC).]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "descriptions",
					Description: `(Optional) The descriptions of the NICs.`,
				},
				resource.Attribute{
					Name:        "is_source_dest_check",
					Description: `(Optional) Whether the source/destination checking is enabled (true) or disabled (false).`,
				},
				resource.Attribute{
					Name:        "link_nic_delete_on_vm_deletion",
					Description: `(Optional) Whether the NICs are deleted when the VMs they are attached to are terminated.`,
				},
				resource.Attribute{
					Name:        "link_nic_device_numbers",
					Description: `(Optional) The device numbers the NICs are attached to.`,
				},
				resource.Attribute{
					Name:        "link_nic_link_nic_ids",
					Description: `(Optional) The attachment IDs of the NICs.`,
				},
				resource.Attribute{
					Name:        "link_nic_states",
					Description: `(Optional) The states of the attachments.`,
				},
				resource.Attribute{
					Name:        "link_nic_vm_account_ids",
					Description: `(Optional) The account IDs of the owners of the VMs the NICs are attached to.`,
				},
				resource.Attribute{
					Name:        "link_nic_vm_ids",
					Description: `(Optional) The IDs of the VMs the NICs are attached to.`,
				},
				resource.Attribute{
					Name:        "link_public_ip_account_ids",
					Description: `(Optional) The account IDs of the owners of the public IPs associated with the NICs.`,
				},
				resource.Attribute{
					Name:        "link_public_ip_link_public_ip_ids",
					Description: `(Optional) The association IDs returned when the public IPs were associated with the NICs.`,
				},
				resource.Attribute{
					Name:        "link_public_ip_public_ip_ids",
					Description: `(Optional) The allocation IDs returned when the public IPs were allocated to their accounts.`,
				},
				resource.Attribute{
					Name:        "link_public_ip_public_ips",
					Description: `(Optional) The public IPs associated with the NICs.`,
				},
				resource.Attribute{
					Name:        "mac_addresses",
					Description: `(Optional) The Media Access Control (MAC) addresses of the NICs.`,
				},
				resource.Attribute{
					Name:        "net_ids",
					Description: `(Optional) The IDs of the Nets where the NICs are located.`,
				},
				resource.Attribute{
					Name:        "nic_ids",
					Description: `(Optional) The IDs of the NICs.`,
				},
				resource.Attribute{
					Name:        "private_dns_names",
					Description: `(Optional) The private DNS names associated with the primary private IPs.`,
				},
				resource.Attribute{
					Name:        "private_ips_link_public_ip_account_ids",
					Description: `(Optional) The account IDs of the owner of the public IPs associated with the private IPs.`,
				},
				resource.Attribute{
					Name:        "private_ips_link_public_ip_public_ips",
					Description: `(Optional) The public IPs associated with the private IPs.`,
				},
				resource.Attribute{
					Name:        "private_ips_primary_ip",
					Description: `(Optional) Whether the private IP is the primary IP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "private_ips_private_ips",
					Description: `(Optional) The private IPs of the NICs.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Optional) The IDs of the security groups associated with the NICs.`,
				},
				resource.Attribute{
					Name:        "security_group_names",
					Description: `(Optional) The names of the security groups associated with the NICs.`,
				},
				resource.Attribute{
					Name:        "states",
					Description: `(Optional) The states of the NICs.`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `(Optional) The IDs of the Subnets for the NICs.`,
				},
				resource.Attribute{
					Name:        "subregion_names",
					Description: `(Optional) The Subregions where the NICs are located.`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the NICs.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the NICs.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combinations of the tags associated with the NICs, in the following format: ` + "`" + `TAGKEY=TAGVALUE` + "`" + `. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the NIC.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the NIC.`,
				},
				resource.Attribute{
					Name:        "is_source_dest_checked",
					Description: `(Net only) If true, the source/destination check is enabled. If false, it is disabled. This value must be false for a NAT VM to perform network address translation (NAT) in a Net.`,
				},
				resource.Attribute{
					Name:        "link_nic",
					Description: `Information about the NIC attachment.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If true, the NIC is deleted when the VM is terminated.`,
				},
				resource.Attribute{
					Name:        "device_number",
					Description: `The device index for the NIC attachment (between ` + "`" + `1` + "`" + ` and ` + "`" + `7` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "link_nic_id",
					Description: `The ID of the NIC to attach.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment (` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "vm_account_id",
					Description: `The account ID of the owner of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM.`,
				},
				resource.Attribute{
					Name:        "link_public_ip",
					Description: `Information about the public IP association.`,
				},
				resource.Attribute{
					Name:        "link_public_ip_id",
					Description: `(Required in a Net) The ID representing the association of the public IP with the VM or the NIC.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_ip_account_id",
					Description: `The account ID of the owner of the public IP.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `The allocation ID of the public IP.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `The Media Access Control (MAC) address of the NIC.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the NIC.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC.`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `The private IPs of the NIC.`,
				},
				resource.Attribute{
					Name:        "is_primary",
					Description: `If true, the IP is the primary private IP of the NIC.`,
				},
				resource.Attribute{
					Name:        "link_public_ip",
					Description: `Information about the public IP association.`,
				},
				resource.Attribute{
					Name:        "link_public_ip_id",
					Description: `(Required in a Net) The ID representing the association of the public IP with the VM or the NIC.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_ip_account_id",
					Description: `The account ID of the owner of the public IP.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `The allocation ID of the public IP.`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP of the NIC.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `One or more IDs of security groups for the NIC.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the NIC (` + "`" + `available` + "`" + ` \| ` + "`" + `attaching` + "`" + ` \| ` + "`" + `in-use` + "`" + ` \| ` + "`" + `detaching` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet.`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The Subregion in which the NIC is located.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the NIC.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the NIC.`,
				},
				resource.Attribute{
					Name:        "is_source_dest_checked",
					Description: `(Net only) If true, the source/destination check is enabled. If false, it is disabled. This value must be false for a NAT VM to perform network address translation (NAT) in a Net.`,
				},
				resource.Attribute{
					Name:        "link_nic",
					Description: `Information about the NIC attachment.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If true, the NIC is deleted when the VM is terminated.`,
				},
				resource.Attribute{
					Name:        "device_number",
					Description: `The device index for the NIC attachment (between ` + "`" + `1` + "`" + ` and ` + "`" + `7` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "link_nic_id",
					Description: `The ID of the NIC to attach.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment (` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "vm_account_id",
					Description: `The account ID of the owner of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM.`,
				},
				resource.Attribute{
					Name:        "link_public_ip",
					Description: `Information about the public IP association.`,
				},
				resource.Attribute{
					Name:        "link_public_ip_id",
					Description: `(Required in a Net) The ID representing the association of the public IP with the VM or the NIC.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_ip_account_id",
					Description: `The account ID of the owner of the public IP.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `The allocation ID of the public IP.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `The Media Access Control (MAC) address of the NIC.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the NIC.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC.`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `The private IPs of the NIC.`,
				},
				resource.Attribute{
					Name:        "is_primary",
					Description: `If true, the IP is the primary private IP of the NIC.`,
				},
				resource.Attribute{
					Name:        "link_public_ip",
					Description: `Information about the public IP association.`,
				},
				resource.Attribute{
					Name:        "link_public_ip_id",
					Description: `(Required in a Net) The ID representing the association of the public IP with the VM or the NIC.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_ip_account_id",
					Description: `The account ID of the owner of the public IP.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `The allocation ID of the public IP.`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP of the NIC.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `One or more IDs of security groups for the NIC.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the NIC (` + "`" + `available` + "`" + ` \| ` + "`" + `attaching` + "`" + ` \| ` + "`" + `in-use` + "`" + ` \| ` + "`" + `detaching` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet.`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The Subregion in which the NIC is located.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_nics",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about network interface cards (NICs).]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "descriptions",
					Description: `(Optional) The descriptions of the NICs.`,
				},
				resource.Attribute{
					Name:        "is_source_dest_check",
					Description: `(Optional) Whether the source/destination checking is enabled (true) or disabled (false).`,
				},
				resource.Attribute{
					Name:        "link_nic_delete_on_vm_deletion",
					Description: `(Optional) Whether the NICs are deleted when the VMs they are attached to are terminated.`,
				},
				resource.Attribute{
					Name:        "link_nic_device_numbers",
					Description: `(Optional) The device numbers the NICs are attached to.`,
				},
				resource.Attribute{
					Name:        "link_nic_link_nic_ids",
					Description: `(Optional) The attachment IDs of the NICs.`,
				},
				resource.Attribute{
					Name:        "link_nic_states",
					Description: `(Optional) The states of the attachments.`,
				},
				resource.Attribute{
					Name:        "link_nic_vm_account_ids",
					Description: `(Optional) The account IDs of the owners of the VMs the NICs are attached to.`,
				},
				resource.Attribute{
					Name:        "link_nic_vm_ids",
					Description: `(Optional) The IDs of the VMs the NICs are attached to.`,
				},
				resource.Attribute{
					Name:        "link_public_ip_account_ids",
					Description: `(Optional) The account IDs of the owners of the public IPs associated with the NICs.`,
				},
				resource.Attribute{
					Name:        "link_public_ip_link_public_ip_ids",
					Description: `(Optional) The association IDs returned when the public IPs were associated with the NICs.`,
				},
				resource.Attribute{
					Name:        "link_public_ip_public_ip_ids",
					Description: `(Optional) The allocation IDs returned when the public IPs were allocated to their accounts.`,
				},
				resource.Attribute{
					Name:        "link_public_ip_public_ips",
					Description: `(Optional) The public IPs associated with the NICs.`,
				},
				resource.Attribute{
					Name:        "mac_addresses",
					Description: `(Optional) The Media Access Control (MAC) addresses of the NICs.`,
				},
				resource.Attribute{
					Name:        "net_ids",
					Description: `(Optional) The IDs of the Nets where the NICs are located.`,
				},
				resource.Attribute{
					Name:        "nic_ids",
					Description: `(Optional) The IDs of the NICs.`,
				},
				resource.Attribute{
					Name:        "private_dns_names",
					Description: `(Optional) The private DNS names associated with the primary private IPs.`,
				},
				resource.Attribute{
					Name:        "private_ips_link_public_ip_account_ids",
					Description: `(Optional) The account IDs of the owner of the public IPs associated with the private IPs.`,
				},
				resource.Attribute{
					Name:        "private_ips_link_public_ip_public_ips",
					Description: `(Optional) The public IPs associated with the private IPs.`,
				},
				resource.Attribute{
					Name:        "private_ips_primary_ip",
					Description: `(Optional) Whether the private IP is the primary IP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "private_ips_private_ips",
					Description: `(Optional) The private IPs of the NICs.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Optional) The IDs of the security groups associated with the NICs.`,
				},
				resource.Attribute{
					Name:        "security_group_names",
					Description: `(Optional) The names of the security groups associated with the NICs.`,
				},
				resource.Attribute{
					Name:        "states",
					Description: `(Optional) The states of the NICs.`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `(Optional) The IDs of the Subnets for the NICs.`,
				},
				resource.Attribute{
					Name:        "subregion_names",
					Description: `(Optional) The Subregions where the NICs are located.`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the NICs.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the NICs.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combinations of the tags associated with the NICs, in the following format: ` + "`" + `TAGKEY=TAGVALUE` + "`" + `. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "nics",
					Description: `Information about one or more NICs.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the NIC.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the NIC.`,
				},
				resource.Attribute{
					Name:        "is_source_dest_checked",
					Description: `(Net only) If true, the source/destination check is enabled. If false, it is disabled. This value must be false for a NAT VM to perform network address translation (NAT) in a Net.`,
				},
				resource.Attribute{
					Name:        "link_nic",
					Description: `Information about the NIC attachment.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If true, the NIC is deleted when the VM is terminated.`,
				},
				resource.Attribute{
					Name:        "device_number",
					Description: `The device index for the NIC attachment (between ` + "`" + `1` + "`" + ` and ` + "`" + `7` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "link_nic_id",
					Description: `The ID of the NIC to attach.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment (` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "vm_account_id",
					Description: `The account ID of the owner of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM.`,
				},
				resource.Attribute{
					Name:        "link_public_ip",
					Description: `Information about the public IP association.`,
				},
				resource.Attribute{
					Name:        "link_public_ip_id",
					Description: `(Required in a Net) The ID representing the association of the public IP with the VM or the NIC.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_ip_account_id",
					Description: `The account ID of the owner of the public IP.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `The allocation ID of the public IP.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `The Media Access Control (MAC) address of the NIC.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the NIC.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC.`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `The private IPs of the NIC.`,
				},
				resource.Attribute{
					Name:        "is_primary",
					Description: `If true, the IP is the primary private IP of the NIC.`,
				},
				resource.Attribute{
					Name:        "link_public_ip",
					Description: `Information about the public IP association.`,
				},
				resource.Attribute{
					Name:        "link_public_ip_id",
					Description: `(Required in a Net) The ID representing the association of the public IP with the VM or the NIC.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_ip_account_id",
					Description: `The account ID of the owner of the public IP.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `The allocation ID of the public IP.`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP of the NIC.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `One or more IDs of security groups for the NIC.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the NIC (` + "`" + `available` + "`" + ` \| ` + "`" + `attaching` + "`" + ` \| ` + "`" + `in-use` + "`" + ` \| ` + "`" + `detaching` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet.`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The Subregion in which the NIC is located.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "nics",
					Description: `Information about one or more NICs.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the NIC.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the NIC.`,
				},
				resource.Attribute{
					Name:        "is_source_dest_checked",
					Description: `(Net only) If true, the source/destination check is enabled. If false, it is disabled. This value must be false for a NAT VM to perform network address translation (NAT) in a Net.`,
				},
				resource.Attribute{
					Name:        "link_nic",
					Description: `Information about the NIC attachment.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If true, the NIC is deleted when the VM is terminated.`,
				},
				resource.Attribute{
					Name:        "device_number",
					Description: `The device index for the NIC attachment (between ` + "`" + `1` + "`" + ` and ` + "`" + `7` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "link_nic_id",
					Description: `The ID of the NIC to attach.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment (` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "vm_account_id",
					Description: `The account ID of the owner of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM.`,
				},
				resource.Attribute{
					Name:        "link_public_ip",
					Description: `Information about the public IP association.`,
				},
				resource.Attribute{
					Name:        "link_public_ip_id",
					Description: `(Required in a Net) The ID representing the association of the public IP with the VM or the NIC.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_ip_account_id",
					Description: `The account ID of the owner of the public IP.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `The allocation ID of the public IP.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `The Media Access Control (MAC) address of the NIC.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the NIC.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC.`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `The private IPs of the NIC.`,
				},
				resource.Attribute{
					Name:        "is_primary",
					Description: `If true, the IP is the primary private IP of the NIC.`,
				},
				resource.Attribute{
					Name:        "link_public_ip",
					Description: `Information about the public IP association.`,
				},
				resource.Attribute{
					Name:        "link_public_ip_id",
					Description: `(Required in a Net) The ID representing the association of the public IP with the VM or the NIC.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_ip_account_id",
					Description: `The account ID of the owner of the public IP.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `The allocation ID of the public IP.`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP of the NIC.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `One or more IDs of security groups for the NIC.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the NIC (` + "`" + `available` + "`" + ` \| ` + "`" + `attaching` + "`" + ` \| ` + "`" + `in-use` + "`" + ` \| ` + "`" + `detaching` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet.`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The Subregion in which the NIC is located.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_product_type",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a product type.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "product_type_ids",
					Description: `(Optional) The IDs of the product types. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the product type.`,
				},
				resource.Attribute{
					Name:        "product_type_id",
					Description: `The ID of the product type.`,
				},
				resource.Attribute{
					Name:        "vendor",
					Description: `The vendor of the product type.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the product type.`,
				},
				resource.Attribute{
					Name:        "product_type_id",
					Description: `The ID of the product type.`,
				},
				resource.Attribute{
					Name:        "vendor",
					Description: `The vendor of the product type.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_product_types",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about product types.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "product_type_ids",
					Description: `(Optional) The IDs of the product types. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "product_types",
					Description: `Information about one or more product types.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the product type.`,
				},
				resource.Attribute{
					Name:        "product_type_id",
					Description: `The ID of the product type.`,
				},
				resource.Attribute{
					Name:        "vendor",
					Description: `The vendor of the product type.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "product_types",
					Description: `Information about one or more product types.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the product type.`,
				},
				resource.Attribute{
					Name:        "product_type_id",
					Description: `The ID of the product type.`,
				},
				resource.Attribute{
					Name:        "vendor",
					Description: `The vendor of the product type.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_public_ip",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a public IP.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "link_public_ip_ids",
					Description: `(Optional) The IDs representing the associations of public IPs with VMs or NICs.`,
				},
				resource.Attribute{
					Name:        "nic_account_ids",
					Description: `(Optional) The account IDs of the owners of the NICs.`,
				},
				resource.Attribute{
					Name:        "nic_ids",
					Description: `(Optional) The IDs of the NICs.`,
				},
				resource.Attribute{
					Name:        "placements",
					Description: `(Optional) Whether the public IPs are for use in the public Cloud or in a Net.`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `(Optional) The private IPs associated with the public IPs.`,
				},
				resource.Attribute{
					Name:        "public_ip_ids",
					Description: `(Optional) The IDs of the public IPs.`,
				},
				resource.Attribute{
					Name:        "public_ips",
					Description: `(Optional) The public IPs.`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the public IPs.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the public IPs.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combinations of the tags associated with the public IPs, in the following format: ` + "`" + `TAGKEY=TAGVALUE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vm_ids",
					Description: `(Optional) The IDs of the VMs. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "link_public_ip_id",
					Description: `(Required in a Net) The ID representing the association of the public IP with the VM or the NIC.`,
				},
				resource.Attribute{
					Name:        "nic_account_id",
					Description: `The account ID of the owner of the NIC.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC the public IP is associated with (if any).`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP associated with the public IP.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `The allocation ID of the public IP.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the public IP.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM the public IP is associated with (if any).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "link_public_ip_id",
					Description: `(Required in a Net) The ID representing the association of the public IP with the VM or the NIC.`,
				},
				resource.Attribute{
					Name:        "nic_account_id",
					Description: `The account ID of the owner of the NIC.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC the public IP is associated with (if any).`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP associated with the public IP.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `The allocation ID of the public IP.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the public IP.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM the public IP is associated with (if any).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_public_ips",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about public IPs.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "link_public_ip_ids",
					Description: `(Optional) The IDs representing the associations of public IPs with VMs or NICs.`,
				},
				resource.Attribute{
					Name:        "nic_account_ids",
					Description: `(Optional) The account IDs of the owners of the NICs.`,
				},
				resource.Attribute{
					Name:        "nic_ids",
					Description: `(Optional) The IDs of the NICs.`,
				},
				resource.Attribute{
					Name:        "placements",
					Description: `(Optional) Whether the public IPs are for use in the public Cloud or in a Net.`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `(Optional) The private IPs associated with the public IPs.`,
				},
				resource.Attribute{
					Name:        "public_ip_ids",
					Description: `(Optional) The IDs of the public IPs.`,
				},
				resource.Attribute{
					Name:        "public_ips",
					Description: `(Optional) The public IPs.`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the public IPs.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the public IPs.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combinations of the tags associated with the public IPs, in the following format: ` + "`" + `TAGKEY=TAGVALUE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vm_ids",
					Description: `(Optional) The IDs of the VMs. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "public_ips",
					Description: `Information about one or more public IPs.`,
				},
				resource.Attribute{
					Name:        "link_public_ip_id",
					Description: `(Required in a Net) The ID representing the association of the public IP with the VM or the NIC.`,
				},
				resource.Attribute{
					Name:        "nic_account_id",
					Description: `The account ID of the owner of the NIC.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC the public IP is associated with (if any).`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP associated with the public IP.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `The allocation ID of the public IP.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the public IP.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM the public IP is associated with (if any).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "public_ips",
					Description: `Information about one or more public IPs.`,
				},
				resource.Attribute{
					Name:        "link_public_ip_id",
					Description: `(Required in a Net) The ID representing the association of the public IP with the VM or the NIC.`,
				},
				resource.Attribute{
					Name:        "nic_account_id",
					Description: `The account ID of the owner of the NIC.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC the public IP is associated with (if any).`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP associated with the public IP.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP.`,
				},
				resource.Attribute{
					Name:        "public_ip_id",
					Description: `The allocation ID of the public IP.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the public IP.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM the public IP is associated with (if any).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_quota",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a quota.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "collections",
					Description: `(Optional) The group names of the quotas.`,
				},
				resource.Attribute{
					Name:        "quota_names",
					Description: `(Optional) The names of the quotas.`,
				},
				resource.Attribute{
					Name:        "quota_types",
					Description: `(Optional) The resource IDs if these are resource-specific quotas, ` + "`" + `global` + "`" + ` if they are not.`,
				},
				resource.Attribute{
					Name:        "short_descriptions",
					Description: `(Optional) The description of the quotas. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the quotas.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the quota.`,
				},
				resource.Attribute{
					Name:        "max_value",
					Description: `The maximum value of the quota for the OUTSCALE user account (if there is no limit, ` + "`" + `0` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The unique name of the quota.`,
				},
				resource.Attribute{
					Name:        "quota_collection",
					Description: `The group name of the quota.`,
				},
				resource.Attribute{
					Name:        "quota_type",
					Description: `The resource ID if it is a resource-specific quota, ` + "`" + `global` + "`" + ` if it is not.`,
				},
				resource.Attribute{
					Name:        "short_description",
					Description: `The description of the quota.`,
				},
				resource.Attribute{
					Name:        "used_value",
					Description: `The limit value currently used by the OUTSCALE user account.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the quotas.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the quota.`,
				},
				resource.Attribute{
					Name:        "max_value",
					Description: `The maximum value of the quota for the OUTSCALE user account (if there is no limit, ` + "`" + `0` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The unique name of the quota.`,
				},
				resource.Attribute{
					Name:        "quota_collection",
					Description: `The group name of the quota.`,
				},
				resource.Attribute{
					Name:        "quota_type",
					Description: `The resource ID if it is a resource-specific quota, ` + "`" + `global` + "`" + ` if it is not.`,
				},
				resource.Attribute{
					Name:        "short_description",
					Description: `The description of the quota.`,
				},
				resource.Attribute{
					Name:        "used_value",
					Description: `The limit value currently used by the OUTSCALE user account.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_quotas",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about quotas.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "collections",
					Description: `(Optional) The group names of the quotas.`,
				},
				resource.Attribute{
					Name:        "quota_names",
					Description: `(Optional) The names of the quotas.`,
				},
				resource.Attribute{
					Name:        "quota_types",
					Description: `(Optional) The resource IDs if these are resource-specific quotas, ` + "`" + `global` + "`" + ` if they are not.`,
				},
				resource.Attribute{
					Name:        "short_descriptions",
					Description: `(Optional) The description of the quotas. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "quotas",
					Description: `One or more quotas associated with the user.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the quotas.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the quota.`,
				},
				resource.Attribute{
					Name:        "max_value",
					Description: `The maximum value of the quota for the OUTSCALE user account (if there is no limit, ` + "`" + `0` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The unique name of the quota.`,
				},
				resource.Attribute{
					Name:        "quota_collection",
					Description: `The group name of the quota.`,
				},
				resource.Attribute{
					Name:        "quota_type",
					Description: `The ressource ID if it is a resource-specific quota, ` + "`" + `global` + "`" + ` if it is not.`,
				},
				resource.Attribute{
					Name:        "short_description",
					Description: `The description of the quota.`,
				},
				resource.Attribute{
					Name:        "used_value",
					Description: `The limit value currently used by the OUTSCALE user account.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "quotas",
					Description: `One or more quotas associated with the user.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the quotas.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the quota.`,
				},
				resource.Attribute{
					Name:        "max_value",
					Description: `The maximum value of the quota for the OUTSCALE user account (if there is no limit, ` + "`" + `0` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The unique name of the quota.`,
				},
				resource.Attribute{
					Name:        "quota_collection",
					Description: `The group name of the quota.`,
				},
				resource.Attribute{
					Name:        "quota_type",
					Description: `The ressource ID if it is a resource-specific quota, ` + "`" + `global` + "`" + ` if it is not.`,
				},
				resource.Attribute{
					Name:        "short_description",
					Description: `The description of the quota.`,
				},
				resource.Attribute{
					Name:        "used_value",
					Description: `The limit value currently used by the OUTSCALE user account.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_regions",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about Regions.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "regions",
					Description: `Information about one or more Regions.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `The hostname of the gateway to access the Region.`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `The administrative name of the Region.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "regions",
					Description: `Information about one or more Regions.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `The hostname of the gateway to access the Region.`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `The administrative name of the Region.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_route_table",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a route table.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "link_route_table_ids",
					Description: `(Optional) The IDs of the route tables involved in the associations.`,
				},
				resource.Attribute{
					Name:        "link_route_table_link_route_table_ids",
					Description: `(Optional) The IDs of the associations between the route tables and the Subnets.`,
				},
				resource.Attribute{
					Name:        "link_route_table_main",
					Description: `(Optional) If true, the route tables are the main ones for their Nets.`,
				},
				resource.Attribute{
					Name:        "link_subnet_ids",
					Description: `(Optional) The IDs of the Subnets involved in the associations.`,
				},
				resource.Attribute{
					Name:        "net_ids",
					Description: `(Optional) The IDs of the Nets for the route tables.`,
				},
				resource.Attribute{
					Name:        "route_creation_methods",
					Description: `(Optional) The methods used to create a route.`,
				},
				resource.Attribute{
					Name:        "route_destination_ip_ranges",
					Description: `(Optional) The IP ranges specified in routes in the tables.`,
				},
				resource.Attribute{
					Name:        "route_destination_service_ids",
					Description: `(Optional) The service IDs specified in routes in the tables.`,
				},
				resource.Attribute{
					Name:        "route_gateway_ids",
					Description: `(Optional) The IDs of the gateways specified in routes in the tables.`,
				},
				resource.Attribute{
					Name:        "route_nat_service_ids",
					Description: `(Optional) The IDs of the NAT services specified in routes in the tables.`,
				},
				resource.Attribute{
					Name:        "route_net_peering_ids",
					Description: `(Optional) The IDs of the Net peerings specified in routes in the tables.`,
				},
				resource.Attribute{
					Name:        "route_states",
					Description: `(Optional) The states of routes in the route tables (always ` + "`" + `active` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "route_table_ids",
					Description: `(Optional) The IDs of the route tables.`,
				},
				resource.Attribute{
					Name:        "route_vm_ids",
					Description: `(Optional) The IDs of the VMs specified in routes in the tables.`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the route tables.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the route tables.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combinations of the tags associated with the route tables, in the following format: ` + "`" + `TAGKEY=TAGVALUE` + "`" + `. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "link_route_tables",
					Description: `One or more associations between the route table and Subnets.`,
				},
				resource.Attribute{
					Name:        "link_route_table_id",
					Description: `The ID of the association between the route table and the Subnet.`,
				},
				resource.Attribute{
					Name:        "main",
					Description: `If true, the route table is the main one.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The ID of the route table.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the route table.`,
				},
				resource.Attribute{
					Name:        "route_propagating_virtual_gateways",
					Description: `Information about virtual gateways propagating routes.`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_id",
					Description: `The ID of the virtual gateway.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The ID of the route table.`,
				},
				resource.Attribute{
					Name:        "routes",
					Description: `One or more routes in the route table.`,
				},
				resource.Attribute{
					Name:        "creation_method",
					Description: `The method used to create the route.`,
				},
				resource.Attribute{
					Name:        "destination_ip_range",
					Description: `The IP range used for the destination match, in CIDR notation (for example, ` + "`" + `10.0.0.0/24` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "destination_service_id",
					Description: `The ID of the OUTSCALE service.`,
				},
				resource.Attribute{
					Name:        "gateway_id",
					Description: `The ID of the Internet service or virtual gateway attached to the Net.`,
				},
				resource.Attribute{
					Name:        "nat_service_id",
					Description: `The ID of a NAT service attached to the Net.`,
				},
				resource.Attribute{
					Name:        "net_access_point_id",
					Description: `The ID of the Net access point.`,
				},
				resource.Attribute{
					Name:        "net_peering_id",
					Description: `The ID of the Net peering.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of a route in the route table (always ` + "`" + `active` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "vm_account_id",
					Description: `The account ID of the owner of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of a VM specified in a route in the table.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the route table.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "link_route_tables",
					Description: `One or more associations between the route table and Subnets.`,
				},
				resource.Attribute{
					Name:        "link_route_table_id",
					Description: `The ID of the association between the route table and the Subnet.`,
				},
				resource.Attribute{
					Name:        "main",
					Description: `If true, the route table is the main one.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The ID of the route table.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the route table.`,
				},
				resource.Attribute{
					Name:        "route_propagating_virtual_gateways",
					Description: `Information about virtual gateways propagating routes.`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_id",
					Description: `The ID of the virtual gateway.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The ID of the route table.`,
				},
				resource.Attribute{
					Name:        "routes",
					Description: `One or more routes in the route table.`,
				},
				resource.Attribute{
					Name:        "creation_method",
					Description: `The method used to create the route.`,
				},
				resource.Attribute{
					Name:        "destination_ip_range",
					Description: `The IP range used for the destination match, in CIDR notation (for example, ` + "`" + `10.0.0.0/24` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "destination_service_id",
					Description: `The ID of the OUTSCALE service.`,
				},
				resource.Attribute{
					Name:        "gateway_id",
					Description: `The ID of the Internet service or virtual gateway attached to the Net.`,
				},
				resource.Attribute{
					Name:        "nat_service_id",
					Description: `The ID of a NAT service attached to the Net.`,
				},
				resource.Attribute{
					Name:        "net_access_point_id",
					Description: `The ID of the Net access point.`,
				},
				resource.Attribute{
					Name:        "net_peering_id",
					Description: `The ID of the Net peering.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of a route in the route table (always ` + "`" + `active` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "vm_account_id",
					Description: `The account ID of the owner of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of a VM specified in a route in the table.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the route table.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_route_tables",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about route tables.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "link_route_table_ids",
					Description: `(Optional) The IDs of the route tables involved in the associations.`,
				},
				resource.Attribute{
					Name:        "link_route_table_link_route_table_ids",
					Description: `(Optional) The IDs of the associations between the route tables and the Subnets.`,
				},
				resource.Attribute{
					Name:        "link_route_table_main",
					Description: `(Optional) If true, the route tables are the main ones for their Nets.`,
				},
				resource.Attribute{
					Name:        "link_subnet_ids",
					Description: `(Optional) The IDs of the Subnets involved in the associations.`,
				},
				resource.Attribute{
					Name:        "net_ids",
					Description: `(Optional) The IDs of the Nets for the route tables.`,
				},
				resource.Attribute{
					Name:        "route_creation_methods",
					Description: `(Optional) The methods used to create a route.`,
				},
				resource.Attribute{
					Name:        "route_destination_ip_ranges",
					Description: `(Optional) The IP ranges specified in routes in the tables.`,
				},
				resource.Attribute{
					Name:        "route_destination_service_ids",
					Description: `(Optional) The service IDs specified in routes in the tables.`,
				},
				resource.Attribute{
					Name:        "route_gateway_ids",
					Description: `(Optional) The IDs of the gateways specified in routes in the tables.`,
				},
				resource.Attribute{
					Name:        "route_nat_service_ids",
					Description: `(Optional) The IDs of the NAT services specified in routes in the tables.`,
				},
				resource.Attribute{
					Name:        "route_net_peering_ids",
					Description: `(Optional) The IDs of the Net peerings specified in routes in the tables.`,
				},
				resource.Attribute{
					Name:        "route_states",
					Description: `(Optional) The states of routes in the route tables (always ` + "`" + `active` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "route_table_ids",
					Description: `(Optional) The IDs of the route tables.`,
				},
				resource.Attribute{
					Name:        "route_vm_ids",
					Description: `(Optional) The IDs of the VMs specified in routes in the tables.`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the route tables.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the route tables.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combinations of the tags associated with the route tables, in the following format: ` + "`" + `TAGKEY=TAGVALUE` + "`" + `. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "route_tables",
					Description: `Information about one or more route tables.`,
				},
				resource.Attribute{
					Name:        "link_route_tables",
					Description: `One or more associations between the route table and Subnets.`,
				},
				resource.Attribute{
					Name:        "link_route_table_id",
					Description: `The ID of the association between the route table and the Subnet.`,
				},
				resource.Attribute{
					Name:        "main",
					Description: `If true, the route table is the main one.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The ID of the route table.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the route table.`,
				},
				resource.Attribute{
					Name:        "route_propagating_virtual_gateways",
					Description: `Information about virtual gateways propagating routes.`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_id",
					Description: `The ID of the virtual gateway.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The ID of the route table.`,
				},
				resource.Attribute{
					Name:        "routes",
					Description: `One or more routes in the route table.`,
				},
				resource.Attribute{
					Name:        "creation_method",
					Description: `The method used to create the route.`,
				},
				resource.Attribute{
					Name:        "destination_ip_range",
					Description: `The IP range used for the destination match, in CIDR notation (for example, ` + "`" + `10.0.0.0/24` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "destination_service_id",
					Description: `The ID of the OUTSCALE service.`,
				},
				resource.Attribute{
					Name:        "gateway_id",
					Description: `The ID of the Internet service or virtual gateway attached to the Net.`,
				},
				resource.Attribute{
					Name:        "nat_service_id",
					Description: `The ID of a NAT service attached to the Net.`,
				},
				resource.Attribute{
					Name:        "net_access_point_id",
					Description: `The ID of the Net access point.`,
				},
				resource.Attribute{
					Name:        "net_peering_id",
					Description: `The ID of the Net peering.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of a route in the route table (always ` + "`" + `active` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "vm_account_id",
					Description: `The account ID of the owner of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of a VM specified in a route in the table.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the route table.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "route_tables",
					Description: `Information about one or more route tables.`,
				},
				resource.Attribute{
					Name:        "link_route_tables",
					Description: `One or more associations between the route table and Subnets.`,
				},
				resource.Attribute{
					Name:        "link_route_table_id",
					Description: `The ID of the association between the route table and the Subnet.`,
				},
				resource.Attribute{
					Name:        "main",
					Description: `If true, the route table is the main one.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The ID of the route table.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the route table.`,
				},
				resource.Attribute{
					Name:        "route_propagating_virtual_gateways",
					Description: `Information about virtual gateways propagating routes.`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_id",
					Description: `The ID of the virtual gateway.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The ID of the route table.`,
				},
				resource.Attribute{
					Name:        "routes",
					Description: `One or more routes in the route table.`,
				},
				resource.Attribute{
					Name:        "creation_method",
					Description: `The method used to create the route.`,
				},
				resource.Attribute{
					Name:        "destination_ip_range",
					Description: `The IP range used for the destination match, in CIDR notation (for example, ` + "`" + `10.0.0.0/24` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "destination_service_id",
					Description: `The ID of the OUTSCALE service.`,
				},
				resource.Attribute{
					Name:        "gateway_id",
					Description: `The ID of the Internet service or virtual gateway attached to the Net.`,
				},
				resource.Attribute{
					Name:        "nat_service_id",
					Description: `The ID of a NAT service attached to the Net.`,
				},
				resource.Attribute{
					Name:        "net_access_point_id",
					Description: `The ID of the Net access point.`,
				},
				resource.Attribute{
					Name:        "net_peering_id",
					Description: `The ID of the Net peering.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of a route in the route table (always ` + "`" + `active` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "vm_account_id",
					Description: `The account ID of the owner of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of a VM specified in a route in the table.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the route table.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_security_group",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a security group.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "account_ids",
					Description: `(Optional) The account IDs of the owners of the security groups.`,
				},
				resource.Attribute{
					Name:        "descriptions",
					Description: `(Optional) The descriptions of the security groups.`,
				},
				resource.Attribute{
					Name:        "inbound_rule_account_ids",
					Description: `(Optional) The account IDs that have been granted permissions.`,
				},
				resource.Attribute{
					Name:        "inbound_rule_from_port_ranges",
					Description: `(Optional) The beginnings of the port ranges for the TCP and UDP protocols, or the ICMP type numbers.`,
				},
				resource.Attribute{
					Name:        "inbound_rule_ip_ranges",
					Description: `(Optional) The IP ranges that have been granted permissions, in CIDR notation (for example, ` + "`" + `10.0.0.0/24` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "inbound_rule_protocols",
					Description: `(Optional) The IP protocols for the permissions (` + "`" + `tcp` + "`" + ` \| ` + "`" + `udp` + "`" + ` \| ` + "`" + `icmp` + "`" + `, or a protocol number, or ` + "`" + `-1` + "`" + ` for all protocols).`,
				},
				resource.Attribute{
					Name:        "inbound_rule_security_group_ids",
					Description: `(Optional) The IDs of the security groups that have been granted permissions.`,
				},
				resource.Attribute{
					Name:        "inbound_rule_security_group_names",
					Description: `(Optional) The names of the security groups that have been granted permissions.`,
				},
				resource.Attribute{
					Name:        "inbound_rule_to_port_ranges",
					Description: `(Optional) The ends of the port ranges for the TCP and UDP protocols, or the ICMP code numbers.`,
				},
				resource.Attribute{
					Name:        "net_ids",
					Description: `(Optional) The IDs of the Nets specified when the security groups were created.`,
				},
				resource.Attribute{
					Name:        "outbound_rule_account_ids",
					Description: `(Optional) The account IDs that have been granted permissions.`,
				},
				resource.Attribute{
					Name:        "outbound_rule_from_port_ranges",
					Description: `(Optional) The beginnings of the port ranges for the TCP and UDP protocols, or the ICMP type numbers.`,
				},
				resource.Attribute{
					Name:        "outbound_rule_ip_ranges",
					Description: `(Optional) The IP ranges that have been granted permissions, in CIDR notation (for example, ` + "`" + `10.0.0.0/24` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "outbound_rule_protocols",
					Description: `(Optional) The IP protocols for the permissions (` + "`" + `tcp` + "`" + ` \| ` + "`" + `udp` + "`" + ` \| ` + "`" + `icmp` + "`" + `, or a protocol number, or ` + "`" + `-1` + "`" + ` for all protocols).`,
				},
				resource.Attribute{
					Name:        "outbound_rule_security_group_ids",
					Description: `(Optional) The IDs of the security groups that have been granted permissions.`,
				},
				resource.Attribute{
					Name:        "outbound_rule_security_group_names",
					Description: `(Optional) The names of the security groups that have been granted permissions.`,
				},
				resource.Attribute{
					Name:        "outbound_rule_to_port_ranges",
					Description: `(Optional) The ends of the port ranges for the TCP and UDP protocols, or the ICMP code numbers.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Optional) The IDs of the security groups.`,
				},
				resource.Attribute{
					Name:        "security_group_names",
					Description: `(Optional) The names of the security groups.`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the security groups.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the security groups.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combinations of the tags associated with the security groups, in the following format: ` + "`" + `TAGKEY=TAGVALUE` + "`" + `. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of a user that has been granted permission.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the security group.`,
				},
				resource.Attribute{
					Name:        "inbound_rules",
					Description: `The inbound rules associated with the security group.`,
				},
				resource.Attribute{
					Name:        "from_port_range",
					Description: `The beginning of the port range for the TCP and UDP protocols, or an ICMP type number.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `The IP protocol name (` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `, or ` + "`" + `-1` + "`" + ` for all protocols). By default, ` + "`" + `-1` + "`" + `. In a Net, this can also be an IP protocol number. For more information, see the [IANA.org website](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml).`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `One or more IP ranges for the security group rules, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "security_groups_members",
					Description: `Information about one or more members of a security group.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of a user.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "service_ids",
					Description: `One or more service IDs to allow traffic from a Net to access the corresponding OUTSCALE services. For more information, see [ReadNetAccessPointServices](https://docs.outscale.com/api#readnetaccesspointservices).`,
				},
				resource.Attribute{
					Name:        "to_port_range",
					Description: `The end of the port range for the TCP and UDP protocols, or an ICMP code number.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the security group.`,
				},
				resource.Attribute{
					Name:        "outbound_rules",
					Description: `The outbound rules associated with the security group.`,
				},
				resource.Attribute{
					Name:        "from_port_range",
					Description: `The beginning of the port range for the TCP and UDP protocols, or an ICMP type number.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `The IP protocol name (` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `, or ` + "`" + `-1` + "`" + ` for all protocols). By default, ` + "`" + `-1` + "`" + `. In a Net, this can also be an IP protocol number. For more information, see the [IANA.org website](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml).`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `One or more IP ranges for the security group rules, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "security_groups_members",
					Description: `Information about one or more members of a security group.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of a user.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "service_ids",
					Description: `One or more service IDs to allow traffic from a Net to access the corresponding OUTSCALE services. For more information, see [ReadNetAccessPointServices](https://docs.outscale.com/api#readnetaccesspointservices).`,
				},
				resource.Attribute{
					Name:        "to_port_range",
					Description: `The end of the port range for the TCP and UDP protocols, or an ICMP code number.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the security group.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of a user that has been granted permission.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the security group.`,
				},
				resource.Attribute{
					Name:        "inbound_rules",
					Description: `The inbound rules associated with the security group.`,
				},
				resource.Attribute{
					Name:        "from_port_range",
					Description: `The beginning of the port range for the TCP and UDP protocols, or an ICMP type number.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `The IP protocol name (` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `, or ` + "`" + `-1` + "`" + ` for all protocols). By default, ` + "`" + `-1` + "`" + `. In a Net, this can also be an IP protocol number. For more information, see the [IANA.org website](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml).`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `One or more IP ranges for the security group rules, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "security_groups_members",
					Description: `Information about one or more members of a security group.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of a user.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "service_ids",
					Description: `One or more service IDs to allow traffic from a Net to access the corresponding OUTSCALE services. For more information, see [ReadNetAccessPointServices](https://docs.outscale.com/api#readnetaccesspointservices).`,
				},
				resource.Attribute{
					Name:        "to_port_range",
					Description: `The end of the port range for the TCP and UDP protocols, or an ICMP code number.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the security group.`,
				},
				resource.Attribute{
					Name:        "outbound_rules",
					Description: `The outbound rules associated with the security group.`,
				},
				resource.Attribute{
					Name:        "from_port_range",
					Description: `The beginning of the port range for the TCP and UDP protocols, or an ICMP type number.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `The IP protocol name (` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `, or ` + "`" + `-1` + "`" + ` for all protocols). By default, ` + "`" + `-1` + "`" + `. In a Net, this can also be an IP protocol number. For more information, see the [IANA.org website](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml).`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `One or more IP ranges for the security group rules, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "security_groups_members",
					Description: `Information about one or more members of a security group.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of a user.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "service_ids",
					Description: `One or more service IDs to allow traffic from a Net to access the corresponding OUTSCALE services. For more information, see [ReadNetAccessPointServices](https://docs.outscale.com/api#readnetaccesspointservices).`,
				},
				resource.Attribute{
					Name:        "to_port_range",
					Description: `The end of the port range for the TCP and UDP protocols, or an ICMP code number.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the security group.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_security_groups",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about security groups.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "account_ids",
					Description: `(Optional) The account IDs of the owners of the security groups.`,
				},
				resource.Attribute{
					Name:        "descriptions",
					Description: `(Optional) The descriptions of the security groups.`,
				},
				resource.Attribute{
					Name:        "inbound_rule_account_ids",
					Description: `(Optional) The account IDs that have been granted permissions.`,
				},
				resource.Attribute{
					Name:        "inbound_rule_from_port_ranges",
					Description: `(Optional) The beginnings of the port ranges for the TCP and UDP protocols, or the ICMP type numbers.`,
				},
				resource.Attribute{
					Name:        "inbound_rule_ip_ranges",
					Description: `(Optional) The IP ranges that have been granted permissions, in CIDR notation (for example, ` + "`" + `10.0.0.0/24` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "inbound_rule_protocols",
					Description: `(Optional) The IP protocols for the permissions (` + "`" + `tcp` + "`" + ` \| ` + "`" + `udp` + "`" + ` \| ` + "`" + `icmp` + "`" + `, or a protocol number, or ` + "`" + `-1` + "`" + ` for all protocols).`,
				},
				resource.Attribute{
					Name:        "inbound_rule_security_group_ids",
					Description: `(Optional) The IDs of the security groups that have been granted permissions.`,
				},
				resource.Attribute{
					Name:        "inbound_rule_security_group_names",
					Description: `(Optional) The names of the security groups that have been granted permissions.`,
				},
				resource.Attribute{
					Name:        "inbound_rule_to_port_ranges",
					Description: `(Optional) The ends of the port ranges for the TCP and UDP protocols, or the ICMP code numbers.`,
				},
				resource.Attribute{
					Name:        "net_ids",
					Description: `(Optional) The IDs of the Nets specified when the security groups were created.`,
				},
				resource.Attribute{
					Name:        "outbound_rule_account_ids",
					Description: `(Optional) The account IDs that have been granted permissions.`,
				},
				resource.Attribute{
					Name:        "outbound_rule_from_port_ranges",
					Description: `(Optional) The beginnings of the port ranges for the TCP and UDP protocols, or the ICMP type numbers.`,
				},
				resource.Attribute{
					Name:        "outbound_rule_ip_ranges",
					Description: `(Optional) The IP ranges that have been granted permissions, in CIDR notation (for example, ` + "`" + `10.0.0.0/24` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "outbound_rule_protocols",
					Description: `(Optional) The IP protocols for the permissions (` + "`" + `tcp` + "`" + ` \| ` + "`" + `udp` + "`" + ` \| ` + "`" + `icmp` + "`" + `, or a protocol number, or ` + "`" + `-1` + "`" + ` for all protocols).`,
				},
				resource.Attribute{
					Name:        "outbound_rule_security_group_ids",
					Description: `(Optional) The IDs of the security groups that have been granted permissions.`,
				},
				resource.Attribute{
					Name:        "outbound_rule_security_group_names",
					Description: `(Optional) The names of the security groups that have been granted permissions.`,
				},
				resource.Attribute{
					Name:        "outbound_rule_to_port_ranges",
					Description: `(Optional) The ends of the port ranges for the TCP and UDP protocols, or the ICMP code numbers.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Optional) The IDs of the security groups.`,
				},
				resource.Attribute{
					Name:        "security_group_names",
					Description: `(Optional) The names of the security groups.`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the security groups.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the security groups.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combinations of the tags associated with the security groups, in the following format: ` + "`" + `TAGKEY=TAGVALUE` + "`" + `. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `Information about one or more security groups.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of a user that has been granted permission.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the security group.`,
				},
				resource.Attribute{
					Name:        "inbound_rules",
					Description: `The inbound rules associated with the security group.`,
				},
				resource.Attribute{
					Name:        "from_port_range",
					Description: `The beginning of the port range for the TCP and UDP protocols, or an ICMP type number.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `The IP protocol name (` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `, or ` + "`" + `-1` + "`" + ` for all protocols). By default, ` + "`" + `-1` + "`" + `. In a Net, this can also be an IP protocol number. For more information, see the [IANA.org website](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml).`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `One or more IP ranges for the security group rules, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "security_groups_members",
					Description: `Information about one or more members of a security group.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of a user.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "service_ids",
					Description: `One or more service IDs to allow traffic from a Net to access the corresponding OUTSCALE services. For more information, see [ReadNetAccessPointServices](https://docs.outscale.com/api#readnetaccesspointservices).`,
				},
				resource.Attribute{
					Name:        "to_port_range",
					Description: `The end of the port range for the TCP and UDP protocols, or an ICMP code number.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the security group.`,
				},
				resource.Attribute{
					Name:        "outbound_rules",
					Description: `The outbound rules associated with the security group.`,
				},
				resource.Attribute{
					Name:        "from_port_range",
					Description: `The beginning of the port range for the TCP and UDP protocols, or an ICMP type number.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `The IP protocol name (` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `, or ` + "`" + `-1` + "`" + ` for all protocols). By default, ` + "`" + `-1` + "`" + `. In a Net, this can also be an IP protocol number. For more information, see the [IANA.org website](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml).`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `One or more IP ranges for the security group rules, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "security_groups_members",
					Description: `Information about one or more members of a security group.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of a user.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "service_ids",
					Description: `One or more service IDs to allow traffic from a Net to access the corresponding OUTSCALE services. For more information, see [ReadNetAccessPointServices](https://docs.outscale.com/api#readnetaccesspointservices).`,
				},
				resource.Attribute{
					Name:        "to_port_range",
					Description: `The end of the port range for the TCP and UDP protocols, or an ICMP code number.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the security group.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "security_groups",
					Description: `Information about one or more security groups.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of a user that has been granted permission.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the security group.`,
				},
				resource.Attribute{
					Name:        "inbound_rules",
					Description: `The inbound rules associated with the security group.`,
				},
				resource.Attribute{
					Name:        "from_port_range",
					Description: `The beginning of the port range for the TCP and UDP protocols, or an ICMP type number.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `The IP protocol name (` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `, or ` + "`" + `-1` + "`" + ` for all protocols). By default, ` + "`" + `-1` + "`" + `. In a Net, this can also be an IP protocol number. For more information, see the [IANA.org website](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml).`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `One or more IP ranges for the security group rules, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "security_groups_members",
					Description: `Information about one or more members of a security group.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of a user.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "service_ids",
					Description: `One or more service IDs to allow traffic from a Net to access the corresponding OUTSCALE services. For more information, see [ReadNetAccessPointServices](https://docs.outscale.com/api#readnetaccesspointservices).`,
				},
				resource.Attribute{
					Name:        "to_port_range",
					Description: `The end of the port range for the TCP and UDP protocols, or an ICMP code number.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the security group.`,
				},
				resource.Attribute{
					Name:        "outbound_rules",
					Description: `The outbound rules associated with the security group.`,
				},
				resource.Attribute{
					Name:        "from_port_range",
					Description: `The beginning of the port range for the TCP and UDP protocols, or an ICMP type number.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `The IP protocol name (` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `, or ` + "`" + `-1` + "`" + ` for all protocols). By default, ` + "`" + `-1` + "`" + `. In a Net, this can also be an IP protocol number. For more information, see the [IANA.org website](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml).`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `One or more IP ranges for the security group rules, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "security_groups_members",
					Description: `Information about one or more members of a security group.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of a user.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "service_ids",
					Description: `One or more service IDs to allow traffic from a Net to access the corresponding OUTSCALE services. For more information, see [ReadNetAccessPointServices](https://docs.outscale.com/api#readnetaccesspointservices).`,
				},
				resource.Attribute{
					Name:        "to_port_range",
					Description: `The end of the port range for the TCP and UDP protocols, or an ICMP code number.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the security group.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_server_certificate",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a server certificate.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "paths",
					Description: `(Optional) The paths to the server certificates. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "expiration_date",
					Description: `The date at which the server certificate expires.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the server certificate.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the server certificate.`,
				},
				resource.Attribute{
					Name:        "orn",
					Description: `The Outscale Resource Name (ORN) of the server certificate. For more information, see [Resource Identifiers > Outscale Resource Names (ORNs)](https://docs.outscale.com/en/userguide/Resource-Identifiers.html#_outscale_resource_names_orns).`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The path to the server certificate.`,
				},
				resource.Attribute{
					Name:        "upload_date",
					Description: `The date at which the server certificate has been uploaded.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "expiration_date",
					Description: `The date at which the server certificate expires.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the server certificate.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the server certificate.`,
				},
				resource.Attribute{
					Name:        "orn",
					Description: `The Outscale Resource Name (ORN) of the server certificate. For more information, see [Resource Identifiers > Outscale Resource Names (ORNs)](https://docs.outscale.com/en/userguide/Resource-Identifiers.html#_outscale_resource_names_orns).`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The path to the server certificate.`,
				},
				resource.Attribute{
					Name:        "upload_date",
					Description: `The date at which the server certificate has been uploaded.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_server_certificates",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about server certificates.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "paths",
					Description: `(Optional) The paths to the server certificates. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "server_certificates",
					Description: `Information about one or more server certificates.`,
				},
				resource.Attribute{
					Name:        "expiration_date",
					Description: `The date at which the server certificate expires.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the server certificate.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the server certificate.`,
				},
				resource.Attribute{
					Name:        "orn",
					Description: `The Outscale Resource Name (ORN) of the server certificate. For more information, see [Resource Identifiers > Outscale Resource Names (ORNs)](https://docs.outscale.com/en/userguide/Resource-Identifiers.html#_outscale_resource_names_orns).`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The path to the server certificate.`,
				},
				resource.Attribute{
					Name:        "upload_date",
					Description: `The date at which the server certificate has been uploaded.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "server_certificates",
					Description: `Information about one or more server certificates.`,
				},
				resource.Attribute{
					Name:        "expiration_date",
					Description: `The date at which the server certificate expires.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the server certificate.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the server certificate.`,
				},
				resource.Attribute{
					Name:        "orn",
					Description: `The Outscale Resource Name (ORN) of the server certificate. For more information, see [Resource Identifiers > Outscale Resource Names (ORNs)](https://docs.outscale.com/en/userguide/Resource-Identifiers.html#_outscale_resource_names_orns).`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The path to the server certificate.`,
				},
				resource.Attribute{
					Name:        "upload_date",
					Description: `The date at which the server certificate has been uploaded.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_snapshot",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a snapshot.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "account_aliases",
					Description: `(Optional) The account aliases of the owners of the snapshots.`,
				},
				resource.Attribute{
					Name:        "account_ids",
					Description: `(Optional) The account IDs of the owners of the snapshots.`,
				},
				resource.Attribute{
					Name:        "descriptions",
					Description: `(Optional) The descriptions of the snapshots.`,
				},
				resource.Attribute{
					Name:        "permissions_to_create_volume_account_ids",
					Description: `(Optional) The account IDs of one or more users who have permissions to create volumes.`,
				},
				resource.Attribute{
					Name:        "permissions_to_create_volume_global_permission",
					Description: `(Optional) If true, lists all public volumes. If false, lists all private volumes.`,
				},
				resource.Attribute{
					Name:        "progresses",
					Description: `(Optional) The progresses of the snapshots, as a percentage.`,
				},
				resource.Attribute{
					Name:        "snapshot_ids",
					Description: `(Optional) The IDs of the snapshots.`,
				},
				resource.Attribute{
					Name:        "states",
					Description: `(Optional) The states of the snapshots (` + "`" + `in-queue` + "`" + ` \| ` + "`" + `completed` + "`" + ` \| ` + "`" + `error` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the snapshots.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the snapshots.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combinations of the tags associated with the snapshots, in the following format: ` + "`" + `TAGKEY=TAGVALUE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "volume_ids",
					Description: `(Optional) The IDs of the volumes used to create the snapshots.`,
				},
				resource.Attribute{
					Name:        "volume_sizes",
					Description: `(Optional) The sizes of the volumes used to create the snapshots, in gibibytes (GiB). ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "account_alias",
					Description: `The account alias of the owner of the snapshot.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the snapshot.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `The date and time of creation of the snapshot.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the snapshot.`,
				},
				resource.Attribute{
					Name:        "permissions_to_create_volume",
					Description: `Information about the users who have permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "account_ids",
					Description: `The account ID of one or more users who have permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "global_permission",
					Description: `If true, the resource is public. If false, the resource is private.`,
				},
				resource.Attribute{
					Name:        "progress",
					Description: `The progress of the snapshot, as a percentage.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The ID of the snapshot.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the snapshot (` + "`" + `in-queue` + "`" + ` \| ` + "`" + `completed` + "`" + ` \| ` + "`" + `error` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the snapshot.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The ID of the volume used to create the snapshot.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `The size of the volume used to create the snapshot, in gibibytes (GiB).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "account_alias",
					Description: `The account alias of the owner of the snapshot.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the snapshot.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `The date and time of creation of the snapshot.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the snapshot.`,
				},
				resource.Attribute{
					Name:        "permissions_to_create_volume",
					Description: `Information about the users who have permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "account_ids",
					Description: `The account ID of one or more users who have permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "global_permission",
					Description: `If true, the resource is public. If false, the resource is private.`,
				},
				resource.Attribute{
					Name:        "progress",
					Description: `The progress of the snapshot, as a percentage.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The ID of the snapshot.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the snapshot (` + "`" + `in-queue` + "`" + ` \| ` + "`" + `completed` + "`" + ` \| ` + "`" + `error` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the snapshot.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The ID of the volume used to create the snapshot.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `The size of the volume used to create the snapshot, in gibibytes (GiB).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_snapshot_export_task",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a snapshot export task.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "task_ids",
					Description: `(Optional) The IDs of the export tasks. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `If the snapshot export task fails, an error message appears.`,
				},
				resource.Attribute{
					Name:        "osu_export",
					Description: `Information about the snapshot export task.`,
				},
				resource.Attribute{
					Name:        "disk_image_format",
					Description: `The format of the export disk (` + "`" + `qcow2` + "`" + ` \| ` + "`" + `raw` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "osu_bucket",
					Description: `The name of the OOS bucket the snapshot is exported to.`,
				},
				resource.Attribute{
					Name:        "osu_prefix",
					Description: `The prefix for the key of the OOS object corresponding to the snapshot.`,
				},
				resource.Attribute{
					Name:        "progress",
					Description: `The progress of the snapshot export task, as a percentage.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The ID of the snapshot to be exported.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the snapshot export task (` + "`" + `pending` + "`" + ` \| ` + "`" + `active` + "`" + ` \| ` + "`" + `completed` + "`" + ` \| ` + "`" + `failed` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the snapshot export task.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "task_id",
					Description: `The ID of the snapshot export task.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "comment",
					Description: `If the snapshot export task fails, an error message appears.`,
				},
				resource.Attribute{
					Name:        "osu_export",
					Description: `Information about the snapshot export task.`,
				},
				resource.Attribute{
					Name:        "disk_image_format",
					Description: `The format of the export disk (` + "`" + `qcow2` + "`" + ` \| ` + "`" + `raw` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "osu_bucket",
					Description: `The name of the OOS bucket the snapshot is exported to.`,
				},
				resource.Attribute{
					Name:        "osu_prefix",
					Description: `The prefix for the key of the OOS object corresponding to the snapshot.`,
				},
				resource.Attribute{
					Name:        "progress",
					Description: `The progress of the snapshot export task, as a percentage.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The ID of the snapshot to be exported.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the snapshot export task (` + "`" + `pending` + "`" + ` \| ` + "`" + `active` + "`" + ` \| ` + "`" + `completed` + "`" + ` \| ` + "`" + `failed` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the snapshot export task.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "task_id",
					Description: `The ID of the snapshot export task.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_snapshot_export_tasks",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about snapshot export tasks.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "task_ids",
					Description: `(Optional) The IDs of the export tasks. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "snapshot_export_tasks",
					Description: `Information about one or more snapshot export tasks.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `If the snapshot export task fails, an error message appears.`,
				},
				resource.Attribute{
					Name:        "osu_export",
					Description: `Information about the snapshot export task.`,
				},
				resource.Attribute{
					Name:        "disk_image_format",
					Description: `The format of the export disk (` + "`" + `qcow2` + "`" + ` \| ` + "`" + `raw` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "osu_bucket",
					Description: `The name of the OOS bucket the snapshot is exported to.`,
				},
				resource.Attribute{
					Name:        "osu_prefix",
					Description: `The prefix for the key of the OOS object corresponding to the snapshot.`,
				},
				resource.Attribute{
					Name:        "progress",
					Description: `The progress of the snapshot export task, as a percentage.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The ID of the snapshot to be exported.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the snapshot export task (` + "`" + `pending` + "`" + ` \| ` + "`" + `active` + "`" + ` \| ` + "`" + `completed` + "`" + ` \| ` + "`" + `failed` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the snapshot export task.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "task_id",
					Description: `The ID of the snapshot export task.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "snapshot_export_tasks",
					Description: `Information about one or more snapshot export tasks.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `If the snapshot export task fails, an error message appears.`,
				},
				resource.Attribute{
					Name:        "osu_export",
					Description: `Information about the snapshot export task.`,
				},
				resource.Attribute{
					Name:        "disk_image_format",
					Description: `The format of the export disk (` + "`" + `qcow2` + "`" + ` \| ` + "`" + `raw` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "osu_bucket",
					Description: `The name of the OOS bucket the snapshot is exported to.`,
				},
				resource.Attribute{
					Name:        "osu_prefix",
					Description: `The prefix for the key of the OOS object corresponding to the snapshot.`,
				},
				resource.Attribute{
					Name:        "progress",
					Description: `The progress of the snapshot export task, as a percentage.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The ID of the snapshot to be exported.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the snapshot export task (` + "`" + `pending` + "`" + ` \| ` + "`" + `active` + "`" + ` \| ` + "`" + `completed` + "`" + ` \| ` + "`" + `failed` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the snapshot export task.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "task_id",
					Description: `The ID of the snapshot export task.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_snapshots",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about snapshots.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "account_aliases",
					Description: `(Optional) The account aliases of the owners of the snapshots.`,
				},
				resource.Attribute{
					Name:        "account_ids",
					Description: `(Optional) The account IDs of the owners of the snapshots.`,
				},
				resource.Attribute{
					Name:        "descriptions",
					Description: `(Optional) The descriptions of the snapshots.`,
				},
				resource.Attribute{
					Name:        "permissions_to_create_volume_account_ids",
					Description: `(Optional) The account IDs of one or more users who have permissions to create volumes.`,
				},
				resource.Attribute{
					Name:        "permissions_to_create_volume_global_permission",
					Description: `(Optional) If true, lists all public volumes. If false, lists all private volumes.`,
				},
				resource.Attribute{
					Name:        "progresses",
					Description: `(Optional) The progresses of the snapshots, as a percentage.`,
				},
				resource.Attribute{
					Name:        "snapshot_ids",
					Description: `(Optional) The IDs of the snapshots.`,
				},
				resource.Attribute{
					Name:        "states",
					Description: `(Optional) The states of the snapshots (` + "`" + `in-queue` + "`" + ` \| ` + "`" + `completed` + "`" + ` \| ` + "`" + `error` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the snapshots.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the snapshots.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combinations of the tags associated with the snapshots, in the following format: ` + "`" + `TAGKEY=TAGVALUE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "volume_ids",
					Description: `(Optional) The IDs of the volumes used to create the snapshots.`,
				},
				resource.Attribute{
					Name:        "volume_sizes",
					Description: `(Optional) The sizes of the volumes used to create the snapshots, in gibibytes (GiB). ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "snapshots",
					Description: `Information about one or more snapshots and their permissions.`,
				},
				resource.Attribute{
					Name:        "account_alias",
					Description: `The account alias of the owner of the snapshot.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the snapshot.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `The date and time of creation of the snapshot.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the snapshot.`,
				},
				resource.Attribute{
					Name:        "permissions_to_create_volume",
					Description: `Information about the users who have permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "account_ids",
					Description: `The account ID of one or more users who have permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "global_permission",
					Description: `If true, the resource is public. If false, the resource is private.`,
				},
				resource.Attribute{
					Name:        "progress",
					Description: `The progress of the snapshot, as a percentage.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The ID of the snapshot.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the snapshot (` + "`" + `in-queue` + "`" + ` \| ` + "`" + `completed` + "`" + ` \| ` + "`" + `error` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the snapshot.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The ID of the volume used to create the snapshot.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `The size of the volume used to create the snapshot, in gibibytes (GiB).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "snapshots",
					Description: `Information about one or more snapshots and their permissions.`,
				},
				resource.Attribute{
					Name:        "account_alias",
					Description: `The account alias of the owner of the snapshot.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the snapshot.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `The date and time of creation of the snapshot.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the snapshot.`,
				},
				resource.Attribute{
					Name:        "permissions_to_create_volume",
					Description: `Information about the users who have permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "account_ids",
					Description: `The account ID of one or more users who have permissions for the resource.`,
				},
				resource.Attribute{
					Name:        "global_permission",
					Description: `If true, the resource is public. If false, the resource is private.`,
				},
				resource.Attribute{
					Name:        "progress",
					Description: `The progress of the snapshot, as a percentage.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The ID of the snapshot.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the snapshot (` + "`" + `in-queue` + "`" + ` \| ` + "`" + `completed` + "`" + ` \| ` + "`" + `error` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the snapshot.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The ID of the volume used to create the snapshot.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `The size of the volume used to create the snapshot, in gibibytes (GiB).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_subnet",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a Subnet.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "available_ips_counts",
					Description: `(Optional) The number of available IPs.`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `(Optional) The IP ranges in the Subnets, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "net_ids",
					Description: `(Optional) The IDs of the Nets in which the Subnets are.`,
				},
				resource.Attribute{
					Name:        "states",
					Description: `(Optional) The states of the Subnets (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `(Optional) The IDs of the Subnets.`,
				},
				resource.Attribute{
					Name:        "subregion_names",
					Description: `(Optional) The names of the Subregions in which the Subnets are located.`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the Subnets.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the Subnets.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combinations of the tags associated with the Subnets, in the following format: ` + "`" + `TAGKEY=TAGVALUE` + "`" + `. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "available_ips_count",
					Description: `The number of available IPs in the Subnets.`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range in the Subnet, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "map_public_ip_on_launch",
					Description: `If true, a public IP is assigned to the network interface cards (NICs) created in the specified Subnet.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net in which the Subnet is.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Subnet (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet.`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The name of the Subregion in which the Subnet is located.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Subnet.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "available_ips_count",
					Description: `The number of available IPs in the Subnets.`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range in the Subnet, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "map_public_ip_on_launch",
					Description: `If true, a public IP is assigned to the network interface cards (NICs) created in the specified Subnet.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net in which the Subnet is.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Subnet (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet.`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The name of the Subregion in which the Subnet is located.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Subnet.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_subnets",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about Subnets.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "available_ips_counts",
					Description: `(Optional) The number of available IPs.`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `(Optional) The IP ranges in the Subnets, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "net_ids",
					Description: `(Optional) The IDs of the Nets in which the Subnets are.`,
				},
				resource.Attribute{
					Name:        "states",
					Description: `(Optional) The states of the Subnets (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `(Optional) The IDs of the Subnets.`,
				},
				resource.Attribute{
					Name:        "subregion_names",
					Description: `(Optional) The names of the Subregions in which the Subnets are located.`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the Subnets.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the Subnets.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combinations of the tags associated with the Subnets, in the following format: ` + "`" + `TAGKEY=TAGVALUE` + "`" + `. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `Information about one or more Subnets.`,
				},
				resource.Attribute{
					Name:        "available_ips_count",
					Description: `The number of available IPs in the Subnets.`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range in the Subnet, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "map_public_ip_on_launch",
					Description: `If true, a public IP is assigned to the network interface cards (NICs) created in the specified Subnet.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net in which the Subnet is.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Subnet (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet.`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The name of the Subregion in which the Subnet is located.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Subnet.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "subnets",
					Description: `Information about one or more Subnets.`,
				},
				resource.Attribute{
					Name:        "available_ips_count",
					Description: `The number of available IPs in the Subnets.`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The IP range in the Subnet, in CIDR notation (for example, ` + "`" + `10.0.0.0/16` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "map_public_ip_on_launch",
					Description: `If true, a public IP is assigned to the network interface cards (NICs) created in the specified Subnet.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net in which the Subnet is.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Subnet (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet.`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The name of the Subregion in which the Subnet is located.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the Subnet.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_subregions",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about subregions.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "subregion_names",
					Description: `(Optional) The names of the Subregions. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "subregions",
					Description: `Information about one or more Subregions.`,
				},
				resource.Attribute{
					Name:        "location_code",
					Description: `The location code of the Subregion.`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `The name of the Region containing the Subregion.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Subregion (` + "`" + `available` + "`" + ` \| ` + "`" + `information` + "`" + ` \| ` + "`" + `impaired` + "`" + ` \| ` + "`" + `unavailable` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The name of the Subregion.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "subregions",
					Description: `Information about one or more Subregions.`,
				},
				resource.Attribute{
					Name:        "location_code",
					Description: `The location code of the Subregion.`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `The name of the Region containing the Subregion.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Subregion (` + "`" + `available` + "`" + ` \| ` + "`" + `information` + "`" + ` \| ` + "`" + `impaired` + "`" + ` \| ` + "`" + `unavailable` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The name of the Subregion.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_virtual_gateway",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a virtual gateway.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "connection_types",
					Description: `(Optional) The types of the virtual gateways (only ` + "`" + `ipsec.1` + "`" + ` is supported).`,
				},
				resource.Attribute{
					Name:        "link_net_ids",
					Description: `(Optional) The IDs of the Nets the virtual gateways are attached to.`,
				},
				resource.Attribute{
					Name:        "link_states",
					Description: `(Optional) The current states of the attachments between the virtual gateways and the Nets (` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "states",
					Description: `(Optional) The states of the virtual gateways (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the virtual gateways.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the virtual gateways.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combinations of the tags associated with the virtual gateways, in the following format: ` + "`" + `TAGKEY=TAGVALUE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_ids",
					Description: `(Optional) The IDs of the virtual gateways. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "connection_type",
					Description: `The type of VPN connection supported by the virtual gateway (only ` + "`" + `ipsec.1` + "`" + ` is supported).`,
				},
				resource.Attribute{
					Name:        "net_to_virtual_gateway_links",
					Description: `The Net to which the virtual gateway is attached.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net to which the virtual gateway is attached.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment (` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the virtual gateway (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the virtual gateway.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_id",
					Description: `The ID of the virtual gateway.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "connection_type",
					Description: `The type of VPN connection supported by the virtual gateway (only ` + "`" + `ipsec.1` + "`" + ` is supported).`,
				},
				resource.Attribute{
					Name:        "net_to_virtual_gateway_links",
					Description: `The Net to which the virtual gateway is attached.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net to which the virtual gateway is attached.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment (` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the virtual gateway (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the virtual gateway.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_id",
					Description: `The ID of the virtual gateway.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_virtual_gateways",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about virtual gateways.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "connection_types",
					Description: `(Optional) The types of the virtual gateways (only ` + "`" + `ipsec.1` + "`" + ` is supported).`,
				},
				resource.Attribute{
					Name:        "link_net_ids",
					Description: `(Optional) The IDs of the Nets the virtual gateways are attached to.`,
				},
				resource.Attribute{
					Name:        "link_states",
					Description: `(Optional) The current states of the attachments between the virtual gateways and the Nets (` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "states",
					Description: `(Optional) The states of the virtual gateways (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the virtual gateways.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the virtual gateways.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combinations of the tags associated with the virtual gateways, in the following format: ` + "`" + `TAGKEY=TAGVALUE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_ids",
					Description: `(Optional) The IDs of the virtual gateways. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "virtual_gateways",
					Description: `Information about one or more virtual gateways.`,
				},
				resource.Attribute{
					Name:        "connection_type",
					Description: `The type of VPN connection supported by the virtual gateway (only ` + "`" + `ipsec.1` + "`" + ` is supported).`,
				},
				resource.Attribute{
					Name:        "net_to_virtual_gateway_links",
					Description: `The Net to which the virtual gateway is attached.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net to which the virtual gateway is attached.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment (` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the virtual gateway (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the virtual gateway.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_id",
					Description: `The ID of the virtual gateway.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "virtual_gateways",
					Description: `Information about one or more virtual gateways.`,
				},
				resource.Attribute{
					Name:        "connection_type",
					Description: `The type of VPN connection supported by the virtual gateway (only ` + "`" + `ipsec.1` + "`" + ` is supported).`,
				},
				resource.Attribute{
					Name:        "net_to_virtual_gateway_links",
					Description: `The Net to which the virtual gateway is attached.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net to which the virtual gateway is attached.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment (` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the virtual gateway (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the virtual gateway.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_id",
					Description: `The ID of the virtual gateway.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_vm",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a virtual machine (VM).]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the VMs.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the VMs.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combinations of the tags associated with the VMs, in the following format: ` + "`" + `TAGKEY=TAGVALUE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vm_ids",
					Description: `(Optional) One or more IDs of VMs. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "architecture",
					Description: `The architecture of the VM (` + "`" + `i386` + "`" + ` \| ` + "`" + `x86_64` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "block_device_mappings_created",
					Description: `The block device mapping of the VM.`,
				},
				resource.Attribute{
					Name:        "bsu",
					Description: `Information about the created BSU volume.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If true, the volume is deleted when terminating the VM. If false, the volume is not deleted when terminating the VM.`,
				},
				resource.Attribute{
					Name:        "link_date",
					Description: `The time and date of attachment of the volume to the VM.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the volume.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The ID of the volume.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `The name of the device.`,
				},
				resource.Attribute{
					Name:        "client_token",
					Description: `The idempotency token provided when launching the VM.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `The date and time of creation of the VM.`,
				},
				resource.Attribute{
					Name:        "deletion_protection",
					Description: `If true, you cannot delete the VM unless you change this parameter back to false.`,
				},
				resource.Attribute{
					Name:        "hypervisor",
					Description: `The hypervisor type of the VMs (` + "`" + `ovm` + "`" + ` \| ` + "`" + `xen` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The ID of the OMI used to create the VM.`,
				},
				resource.Attribute{
					Name:        "is_source_dest_checked",
					Description: `(Net only) If true, the source/destination check is enabled. If false, it is disabled. This value must be false for a NAT VM to perform network address translation (NAT) in a Net.`,
				},
				resource.Attribute{
					Name:        "keypair_name",
					Description: `The name of the keypair used when launching the VM.`,
				},
				resource.Attribute{
					Name:        "launch_number",
					Description: `The number for the VM when launching a group of several VMs (for example, ` + "`" + `0` + "`" + `, ` + "`" + `1` + "`" + `, ` + "`" + `2` + "`" + `, and so on).`,
				},
				resource.Attribute{
					Name:        "nested_virtualization",
					Description: `If true, nested virtualization is enabled. If false, it is disabled.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net in which the VM is running.`,
				},
				resource.Attribute{
					Name:        "nics",
					Description: `(Net only) The network interface cards (NICs) the VMs are attached to.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the NIC.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the NIC.`,
				},
				resource.Attribute{
					Name:        "is_source_dest_checked",
					Description: `(Net only) If true, the source/destination check is enabled. If false, it is disabled. This value must be false for a NAT VM to perform network address translation (NAT) in a Net.`,
				},
				resource.Attribute{
					Name:        "link_nic",
					Description: `Information about the network interface card (NIC).`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If true, the NIC is deleted when the VM is terminated.`,
				},
				resource.Attribute{
					Name:        "device_number",
					Description: `The device index for the NIC attachment (between ` + "`" + `1` + "`" + ` and ` + "`" + `7` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "link_nic_id",
					Description: `The ID of the NIC to attach.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment (` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "link_public_ip",
					Description: `Information about the public IP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_ip_account_id",
					Description: `The account ID of the owner of the public IP.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `The Media Access Control (MAC) address of the NIC.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the NIC.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC.`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `The private IP or IPs of the NIC.`,
				},
				resource.Attribute{
					Name:        "is_primary",
					Description: `If true, the IP is the primary private IP of the NIC.`,
				},
				resource.Attribute{
					Name:        "link_public_ip",
					Description: `Information about the public IP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_ip_account_id",
					Description: `The account ID of the owner of the public IP.`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `One or more IDs of security groups for the NIC.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the NIC (` + "`" + `available` + "`" + ` \| ` + "`" + `attaching` + "`" + ` \| ` + "`" + `in-use` + "`" + ` \| ` + "`" + `detaching` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet for the NIC.`,
				},
				resource.Attribute{
					Name:        "os_family",
					Description: `Indicates the operating system (OS) of the VM.`,
				},
				resource.Attribute{
					Name:        "performance",
					Description: `The performance of the VM (` + "`" + `medium` + "`" + ` \| ` + "`" + `high` + "`" + ` \| ` + "`" + `highest` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "placement",
					Description: `Information about the placement of the VM.`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The name of the Subregion. If you specify this parameter, you must not specify the ` + "`" + `nics` + "`" + ` parameter.`,
				},
				resource.Attribute{
					Name:        "tenancy",
					Description: `The tenancy of the VM (` + "`" + `default` + "`" + ` \| ` + "`" + `dedicated` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The primary private IP of the VM.`,
				},
				resource.Attribute{
					Name:        "product_codes",
					Description: `The product code associated with the OMI used to create the VM (` + "`" + `0001` + "`" + ` Linux/Unix \| ` + "`" + `0002` + "`" + ` Windows \| ` + "`" + `0004` + "`" + ` Linux/Oracle \| ` + "`" + `0005` + "`" + ` Windows 10).`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP of the VM.`,
				},
				resource.Attribute{
					Name:        "reservation_id",
					Description: `The reservation ID of the VM.`,
				},
				resource.Attribute{
					Name:        "root_device_name",
					Description: `The name of the root device for the VM (for example, ` + "`" + `/dev/vda1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "root_device_type",
					Description: `The type of root device used by the VM (always ` + "`" + `bsu` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `One or more security groups associated with the VM.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "state_reason",
					Description: `The reason explaining the current state of the VM.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the VM (` + "`" + `pending` + "`" + ` \| ` + "`" + `running` + "`" + ` \| ` + "`" + `stopping` + "`" + ` \| ` + "`" + `stopped` + "`" + ` \| ` + "`" + `shutting-down` + "`" + ` \| ` + "`" + `terminated` + "`" + ` \| ` + "`" + `quarantine` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet for the VM.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the VM.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `The Base64-encoded MIME user data.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_initiated_shutdown_behavior",
					Description: `The VM behavior when you stop it. If set to ` + "`" + `stop` + "`" + `, the VM stops. If set to ` + "`" + `restart` + "`" + `, the VM stops then automatically restarts. If set to ` + "`" + `terminate` + "`" + `, the VM stops and is deleted.`,
				},
				resource.Attribute{
					Name:        "vm_type",
					Description: `The type of VM. For more information, see [Instance Types](https://docs.outscale.com/en/userguide/Instance-Types.html).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "architecture",
					Description: `The architecture of the VM (` + "`" + `i386` + "`" + ` \| ` + "`" + `x86_64` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "block_device_mappings_created",
					Description: `The block device mapping of the VM.`,
				},
				resource.Attribute{
					Name:        "bsu",
					Description: `Information about the created BSU volume.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If true, the volume is deleted when terminating the VM. If false, the volume is not deleted when terminating the VM.`,
				},
				resource.Attribute{
					Name:        "link_date",
					Description: `The time and date of attachment of the volume to the VM.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the volume.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The ID of the volume.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `The name of the device.`,
				},
				resource.Attribute{
					Name:        "client_token",
					Description: `The idempotency token provided when launching the VM.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `The date and time of creation of the VM.`,
				},
				resource.Attribute{
					Name:        "deletion_protection",
					Description: `If true, you cannot delete the VM unless you change this parameter back to false.`,
				},
				resource.Attribute{
					Name:        "hypervisor",
					Description: `The hypervisor type of the VMs (` + "`" + `ovm` + "`" + ` \| ` + "`" + `xen` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The ID of the OMI used to create the VM.`,
				},
				resource.Attribute{
					Name:        "is_source_dest_checked",
					Description: `(Net only) If true, the source/destination check is enabled. If false, it is disabled. This value must be false for a NAT VM to perform network address translation (NAT) in a Net.`,
				},
				resource.Attribute{
					Name:        "keypair_name",
					Description: `The name of the keypair used when launching the VM.`,
				},
				resource.Attribute{
					Name:        "launch_number",
					Description: `The number for the VM when launching a group of several VMs (for example, ` + "`" + `0` + "`" + `, ` + "`" + `1` + "`" + `, ` + "`" + `2` + "`" + `, and so on).`,
				},
				resource.Attribute{
					Name:        "nested_virtualization",
					Description: `If true, nested virtualization is enabled. If false, it is disabled.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net in which the VM is running.`,
				},
				resource.Attribute{
					Name:        "nics",
					Description: `(Net only) The network interface cards (NICs) the VMs are attached to.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the NIC.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the NIC.`,
				},
				resource.Attribute{
					Name:        "is_source_dest_checked",
					Description: `(Net only) If true, the source/destination check is enabled. If false, it is disabled. This value must be false for a NAT VM to perform network address translation (NAT) in a Net.`,
				},
				resource.Attribute{
					Name:        "link_nic",
					Description: `Information about the network interface card (NIC).`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If true, the NIC is deleted when the VM is terminated.`,
				},
				resource.Attribute{
					Name:        "device_number",
					Description: `The device index for the NIC attachment (between ` + "`" + `1` + "`" + ` and ` + "`" + `7` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "link_nic_id",
					Description: `The ID of the NIC to attach.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment (` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "link_public_ip",
					Description: `Information about the public IP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_ip_account_id",
					Description: `The account ID of the owner of the public IP.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `The Media Access Control (MAC) address of the NIC.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the NIC.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC.`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `The private IP or IPs of the NIC.`,
				},
				resource.Attribute{
					Name:        "is_primary",
					Description: `If true, the IP is the primary private IP of the NIC.`,
				},
				resource.Attribute{
					Name:        "link_public_ip",
					Description: `Information about the public IP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_ip_account_id",
					Description: `The account ID of the owner of the public IP.`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `One or more IDs of security groups for the NIC.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the NIC (` + "`" + `available` + "`" + ` \| ` + "`" + `attaching` + "`" + ` \| ` + "`" + `in-use` + "`" + ` \| ` + "`" + `detaching` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet for the NIC.`,
				},
				resource.Attribute{
					Name:        "os_family",
					Description: `Indicates the operating system (OS) of the VM.`,
				},
				resource.Attribute{
					Name:        "performance",
					Description: `The performance of the VM (` + "`" + `medium` + "`" + ` \| ` + "`" + `high` + "`" + ` \| ` + "`" + `highest` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "placement",
					Description: `Information about the placement of the VM.`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The name of the Subregion. If you specify this parameter, you must not specify the ` + "`" + `nics` + "`" + ` parameter.`,
				},
				resource.Attribute{
					Name:        "tenancy",
					Description: `The tenancy of the VM (` + "`" + `default` + "`" + ` \| ` + "`" + `dedicated` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The primary private IP of the VM.`,
				},
				resource.Attribute{
					Name:        "product_codes",
					Description: `The product code associated with the OMI used to create the VM (` + "`" + `0001` + "`" + ` Linux/Unix \| ` + "`" + `0002` + "`" + ` Windows \| ` + "`" + `0004` + "`" + ` Linux/Oracle \| ` + "`" + `0005` + "`" + ` Windows 10).`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP of the VM.`,
				},
				resource.Attribute{
					Name:        "reservation_id",
					Description: `The reservation ID of the VM.`,
				},
				resource.Attribute{
					Name:        "root_device_name",
					Description: `The name of the root device for the VM (for example, ` + "`" + `/dev/vda1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "root_device_type",
					Description: `The type of root device used by the VM (always ` + "`" + `bsu` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `One or more security groups associated with the VM.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "state_reason",
					Description: `The reason explaining the current state of the VM.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the VM (` + "`" + `pending` + "`" + ` \| ` + "`" + `running` + "`" + ` \| ` + "`" + `stopping` + "`" + ` \| ` + "`" + `stopped` + "`" + ` \| ` + "`" + `shutting-down` + "`" + ` \| ` + "`" + `terminated` + "`" + ` \| ` + "`" + `quarantine` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet for the VM.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the VM.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `The Base64-encoded MIME user data.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_initiated_shutdown_behavior",
					Description: `The VM behavior when you stop it. If set to ` + "`" + `stop` + "`" + `, the VM stops. If set to ` + "`" + `restart` + "`" + `, the VM stops then automatically restarts. If set to ` + "`" + `terminate` + "`" + `, the VM stops and is deleted.`,
				},
				resource.Attribute{
					Name:        "vm_type",
					Description: `The type of VM. For more information, see [Instance Types](https://docs.outscale.com/en/userguide/Instance-Types.html).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_vm_state",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a VM state.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "all_vms",
					Description: `(Optional) If true, includes the status of all VMs. By default or if set to false, only includes the status of running VMs.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "maintenance_event_codes",
					Description: `(Optional) The code for the scheduled event (` + "`" + `system-reboot` + "`" + ` \| ` + "`" + `system-maintenance` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "maintenance_event_descriptions",
					Description: `(Optional) The description of the scheduled event.`,
				},
				resource.Attribute{
					Name:        "maintenance_events_not_after",
					Description: `(Optional) The latest time the event can end.`,
				},
				resource.Attribute{
					Name:        "maintenance_events_not_before",
					Description: `(Optional) The earliest time the event can start.`,
				},
				resource.Attribute{
					Name:        "subregion_names",
					Description: `(Optional) The names of the Subregions of the VMs.`,
				},
				resource.Attribute{
					Name:        "vm_ids",
					Description: `(Optional) One or more IDs of VMs.`,
				},
				resource.Attribute{
					Name:        "vm_states",
					Description: `(Optional) The states of the VMs (` + "`" + `pending` + "`" + ` \| ` + "`" + `running` + "`" + ` \| ` + "`" + `stopping` + "`" + ` \| ` + "`" + `stopped` + "`" + ` \| ` + "`" + `shutting-down` + "`" + ` \| ` + "`" + `terminated` + "`" + ` \| ` + "`" + `quarantine` + "`" + `). ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "maintenance_events",
					Description: `One or more scheduled events associated with the VM.`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `The code of the event (` + "`" + `system-reboot` + "`" + ` \| ` + "`" + `system-maintenance` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the event.`,
				},
				resource.Attribute{
					Name:        "not_after",
					Description: `The latest scheduled end time for the event.`,
				},
				resource.Attribute{
					Name:        "not_before",
					Description: `The earliest scheduled start time for the event.`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The name of the Subregion of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_state",
					Description: `The state of the VM (` + "`" + `pending` + "`" + ` \| ` + "`" + `running` + "`" + ` \| ` + "`" + `stopping` + "`" + ` \| ` + "`" + `stopped` + "`" + ` \| ` + "`" + `shutting-down` + "`" + ` \| ` + "`" + `terminated` + "`" + ` \| ` + "`" + `quarantine` + "`" + `).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "maintenance_events",
					Description: `One or more scheduled events associated with the VM.`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `The code of the event (` + "`" + `system-reboot` + "`" + ` \| ` + "`" + `system-maintenance` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the event.`,
				},
				resource.Attribute{
					Name:        "not_after",
					Description: `The latest scheduled end time for the event.`,
				},
				resource.Attribute{
					Name:        "not_before",
					Description: `The earliest scheduled start time for the event.`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The name of the Subregion of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_state",
					Description: `The state of the VM (` + "`" + `pending` + "`" + ` \| ` + "`" + `running` + "`" + ` \| ` + "`" + `stopping` + "`" + ` \| ` + "`" + `stopped` + "`" + ` \| ` + "`" + `shutting-down` + "`" + ` \| ` + "`" + `terminated` + "`" + ` \| ` + "`" + `quarantine` + "`" + `).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_vm_states",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about VM states.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "all_vms",
					Description: `(Optional) If true, includes the status of all VMs. By default or if set to false, only includes the status of running VMs.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "maintenance_event_codes",
					Description: `(Optional) The code for the scheduled event (` + "`" + `system-reboot` + "`" + ` \| ` + "`" + `system-maintenance` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "maintenance_event_descriptions",
					Description: `(Optional) The description of the scheduled event.`,
				},
				resource.Attribute{
					Name:        "maintenance_events_not_after",
					Description: `(Optional) The latest time the event can end.`,
				},
				resource.Attribute{
					Name:        "maintenance_events_not_before",
					Description: `(Optional) The earliest time the event can start.`,
				},
				resource.Attribute{
					Name:        "subregion_names",
					Description: `(Optional) The names of the Subregions of the VMs.`,
				},
				resource.Attribute{
					Name:        "vm_ids",
					Description: `(Optional) One or more IDs of VMs.`,
				},
				resource.Attribute{
					Name:        "vm_states",
					Description: `(Optional) The states of the VMs (` + "`" + `pending` + "`" + ` \| ` + "`" + `running` + "`" + ` \| ` + "`" + `stopping` + "`" + ` \| ` + "`" + `stopped` + "`" + ` \| ` + "`" + `shutting-down` + "`" + ` \| ` + "`" + `terminated` + "`" + ` \| ` + "`" + `quarantine` + "`" + `). ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "vm_states",
					Description: `Information about one or more VM states.`,
				},
				resource.Attribute{
					Name:        "maintenance_events",
					Description: `One or more scheduled events associated with the VM.`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `The code of the event (` + "`" + `system-reboot` + "`" + ` \| ` + "`" + `system-maintenance` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the event.`,
				},
				resource.Attribute{
					Name:        "not_after",
					Description: `The latest scheduled end time for the event.`,
				},
				resource.Attribute{
					Name:        "not_before",
					Description: `The earliest scheduled start time for the event.`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The name of the Subregion of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_state",
					Description: `The state of the VM (` + "`" + `pending` + "`" + ` \| ` + "`" + `running` + "`" + ` \| ` + "`" + `stopping` + "`" + ` \| ` + "`" + `stopped` + "`" + ` \| ` + "`" + `shutting-down` + "`" + ` \| ` + "`" + `terminated` + "`" + ` \| ` + "`" + `quarantine` + "`" + `).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vm_states",
					Description: `Information about one or more VM states.`,
				},
				resource.Attribute{
					Name:        "maintenance_events",
					Description: `One or more scheduled events associated with the VM.`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `The code of the event (` + "`" + `system-reboot` + "`" + ` \| ` + "`" + `system-maintenance` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the event.`,
				},
				resource.Attribute{
					Name:        "not_after",
					Description: `The latest scheduled end time for the event.`,
				},
				resource.Attribute{
					Name:        "not_before",
					Description: `The earliest scheduled start time for the event.`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The name of the Subregion of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_state",
					Description: `The state of the VM (` + "`" + `pending` + "`" + ` \| ` + "`" + `running` + "`" + ` \| ` + "`" + `stopping` + "`" + ` \| ` + "`" + `stopped` + "`" + ` \| ` + "`" + `shutting-down` + "`" + ` \| ` + "`" + `terminated` + "`" + ` \| ` + "`" + `quarantine` + "`" + `).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_vm_types",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about VM types.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "bsu_optimized",
					Description: `(Optional) This parameter is not available. It is present in our API for the sake of historical compatibility with AWS.`,
				},
				resource.Attribute{
					Name:        "memory_sizes",
					Description: `(Optional) The amounts of memory, in gibibytes (GiB).`,
				},
				resource.Attribute{
					Name:        "vcore_counts",
					Description: `(Optional) The numbers of vCores.`,
				},
				resource.Attribute{
					Name:        "vm_type_names",
					Description: `(Optional) The names of the VM types. For more information, see [Instance Types](https://docs.outscale.com/en/userguide/Instance-Types.html).`,
				},
				resource.Attribute{
					Name:        "volume_counts",
					Description: `(Optional) The maximum number of ephemeral storage disks.`,
				},
				resource.Attribute{
					Name:        "volume_sizes",
					Description: `(Optional) The size of one ephemeral storage disk, in gibibytes (GiB). ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "vm_types",
					Description: `Information about one or more VM types.`,
				},
				resource.Attribute{
					Name:        "bsu_optimized",
					Description: `This parameter is not available. It is present in our API for the sake of historical compatibility with AWS.`,
				},
				resource.Attribute{
					Name:        "max_private_ips",
					Description: `The maximum number of private IPs per network interface card (NIC).`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `The amount of memory, in gibibytes.`,
				},
				resource.Attribute{
					Name:        "vcore_count",
					Description: `The number of vCores.`,
				},
				resource.Attribute{
					Name:        "vm_type_name",
					Description: `The name of the VM type.`,
				},
				resource.Attribute{
					Name:        "volume_count",
					Description: `The maximum number of ephemeral storage disks.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `The size of one ephemeral storage disk, in gibibytes (GiB).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vm_types",
					Description: `Information about one or more VM types.`,
				},
				resource.Attribute{
					Name:        "bsu_optimized",
					Description: `This parameter is not available. It is present in our API for the sake of historical compatibility with AWS.`,
				},
				resource.Attribute{
					Name:        "max_private_ips",
					Description: `The maximum number of private IPs per network interface card (NIC).`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `The amount of memory, in gibibytes.`,
				},
				resource.Attribute{
					Name:        "vcore_count",
					Description: `The number of vCores.`,
				},
				resource.Attribute{
					Name:        "vm_type_name",
					Description: `The name of the VM type.`,
				},
				resource.Attribute{
					Name:        "volume_count",
					Description: `The maximum number of ephemeral storage disks.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `The size of one ephemeral storage disk, in gibibytes (GiB).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_vms",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about virtual machines (VMs).]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the VMs.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the VMs.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combinations of the tags associated with the VMs, in the following format: ` + "`" + `TAGKEY=TAGVALUE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vm_ids",
					Description: `(Optional) One or more IDs of VMs. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "vms",
					Description: `Information about one or more VMs.`,
				},
				resource.Attribute{
					Name:        "architecture",
					Description: `The architecture of the VM (` + "`" + `i386` + "`" + ` \| ` + "`" + `x86_64` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "block_device_mappings_created",
					Description: `The block device mapping of the VM.`,
				},
				resource.Attribute{
					Name:        "bsu",
					Description: `Information about the created BSU volume.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If true, the volume is deleted when terminating the VM. If false, the volume is not deleted when terminating the VM.`,
				},
				resource.Attribute{
					Name:        "link_date",
					Description: `The time and date of attachment of the volume to the VM.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the volume.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The ID of the volume.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `The name of the device.`,
				},
				resource.Attribute{
					Name:        "client_token",
					Description: `The idempotency token provided when launching the VM.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `The date and time of creation of the VM.`,
				},
				resource.Attribute{
					Name:        "deletion_protection",
					Description: `If true, you cannot delete the VM unless you change this parameter back to false.`,
				},
				resource.Attribute{
					Name:        "hypervisor",
					Description: `The hypervisor type of the VMs (` + "`" + `ovm` + "`" + ` \| ` + "`" + `xen` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The ID of the OMI used to create the VM.`,
				},
				resource.Attribute{
					Name:        "is_source_dest_checked",
					Description: `(Net only) If true, the source/destination check is enabled. If false, it is disabled. This value must be false for a NAT VM to perform network address translation (NAT) in a Net.`,
				},
				resource.Attribute{
					Name:        "keypair_name",
					Description: `The name of the keypair used when launching the VM.`,
				},
				resource.Attribute{
					Name:        "launch_number",
					Description: `The number for the VM when launching a group of several VMs (for example, ` + "`" + `0` + "`" + `, ` + "`" + `1` + "`" + `, ` + "`" + `2` + "`" + `, and so on).`,
				},
				resource.Attribute{
					Name:        "nested_virtualization",
					Description: `If true, nested virtualization is enabled. If false, it is disabled.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net in which the VM is running.`,
				},
				resource.Attribute{
					Name:        "nics",
					Description: `(Net only) The network interface cards (NICs) the VMs are attached to.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the NIC.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the NIC.`,
				},
				resource.Attribute{
					Name:        "is_source_dest_checked",
					Description: `(Net only) If true, the source/destination check is enabled. If false, it is disabled. This value must be false for a NAT VM to perform network address translation (NAT) in a Net.`,
				},
				resource.Attribute{
					Name:        "link_nic",
					Description: `Information about the network interface card (NIC).`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If true, the NIC is deleted when the VM is terminated.`,
				},
				resource.Attribute{
					Name:        "device_number",
					Description: `The device index for the NIC attachment (between ` + "`" + `1` + "`" + ` and ` + "`" + `7` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "link_nic_id",
					Description: `The ID of the NIC to attach.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment (` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "link_public_ip",
					Description: `Information about the public IP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_ip_account_id",
					Description: `The account ID of the owner of the public IP.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `The Media Access Control (MAC) address of the NIC.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the NIC.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC.`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `The private IP or IPs of the NIC.`,
				},
				resource.Attribute{
					Name:        "is_primary",
					Description: `If true, the IP is the primary private IP of the NIC.`,
				},
				resource.Attribute{
					Name:        "link_public_ip",
					Description: `Information about the public IP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_ip_account_id",
					Description: `The account ID of the owner of the public IP.`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `One or more IDs of security groups for the NIC.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the NIC (` + "`" + `available` + "`" + ` \| ` + "`" + `attaching` + "`" + ` \| ` + "`" + `in-use` + "`" + ` \| ` + "`" + `detaching` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet for the NIC.`,
				},
				resource.Attribute{
					Name:        "os_family",
					Description: `Indicates the operating system (OS) of the VM.`,
				},
				resource.Attribute{
					Name:        "performance",
					Description: `The performance of the VM (` + "`" + `medium` + "`" + ` \| ` + "`" + `high` + "`" + ` \| ` + "`" + `highest` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "placement",
					Description: `Information about the placement of the VM.`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The name of the Subregion. If you specify this parameter, you must not specify the ` + "`" + `nics` + "`" + ` parameter.`,
				},
				resource.Attribute{
					Name:        "tenancy",
					Description: `The tenancy of the VM (` + "`" + `default` + "`" + ` \| ` + "`" + `dedicated` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The primary private IP of the VM.`,
				},
				resource.Attribute{
					Name:        "product_codes",
					Description: `The product code associated with the OMI used to create the VM (` + "`" + `0001` + "`" + ` Linux/Unix \| ` + "`" + `0002` + "`" + ` Windows \| ` + "`" + `0004` + "`" + ` Linux/Oracle \| ` + "`" + `0005` + "`" + ` Windows 10).`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP of the VM.`,
				},
				resource.Attribute{
					Name:        "reservation_id",
					Description: `The reservation ID of the VM.`,
				},
				resource.Attribute{
					Name:        "root_device_name",
					Description: `The name of the root device for the VM (for example, ` + "`" + `/dev/vda1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "root_device_type",
					Description: `The type of root device used by the VM (always ` + "`" + `bsu` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `One or more security groups associated with the VM.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the VM (` + "`" + `pending` + "`" + ` \| ` + "`" + `running` + "`" + ` \| ` + "`" + `stopping` + "`" + ` \| ` + "`" + `stopped` + "`" + ` \| ` + "`" + `shutting-down` + "`" + ` \| ` + "`" + `terminated` + "`" + ` \| ` + "`" + `quarantine` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state_reason",
					Description: `The reason explaining the current state of the VM.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet for the VM.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the VM.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `The Base64-encoded MIME user data.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_initiated_shutdown_behavior",
					Description: `The VM behavior when you stop it. If set to ` + "`" + `stop` + "`" + `, the VM stops. If set to ` + "`" + `restart` + "`" + `, the VM stops then automatically restarts. If set to ` + "`" + `terminate` + "`" + `, the VM stops and is deleted.`,
				},
				resource.Attribute{
					Name:        "vm_type",
					Description: `The type of VM. For more information, see [Instance Types](https://docs.outscale.com/en/userguide/Instance-Types.html).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vms",
					Description: `Information about one or more VMs.`,
				},
				resource.Attribute{
					Name:        "architecture",
					Description: `The architecture of the VM (` + "`" + `i386` + "`" + ` \| ` + "`" + `x86_64` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "block_device_mappings_created",
					Description: `The block device mapping of the VM.`,
				},
				resource.Attribute{
					Name:        "bsu",
					Description: `Information about the created BSU volume.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If true, the volume is deleted when terminating the VM. If false, the volume is not deleted when terminating the VM.`,
				},
				resource.Attribute{
					Name:        "link_date",
					Description: `The time and date of attachment of the volume to the VM.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the volume.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The ID of the volume.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `The name of the device.`,
				},
				resource.Attribute{
					Name:        "client_token",
					Description: `The idempotency token provided when launching the VM.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `The date and time of creation of the VM.`,
				},
				resource.Attribute{
					Name:        "deletion_protection",
					Description: `If true, you cannot delete the VM unless you change this parameter back to false.`,
				},
				resource.Attribute{
					Name:        "hypervisor",
					Description: `The hypervisor type of the VMs (` + "`" + `ovm` + "`" + ` \| ` + "`" + `xen` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The ID of the OMI used to create the VM.`,
				},
				resource.Attribute{
					Name:        "is_source_dest_checked",
					Description: `(Net only) If true, the source/destination check is enabled. If false, it is disabled. This value must be false for a NAT VM to perform network address translation (NAT) in a Net.`,
				},
				resource.Attribute{
					Name:        "keypair_name",
					Description: `The name of the keypair used when launching the VM.`,
				},
				resource.Attribute{
					Name:        "launch_number",
					Description: `The number for the VM when launching a group of several VMs (for example, ` + "`" + `0` + "`" + `, ` + "`" + `1` + "`" + `, ` + "`" + `2` + "`" + `, and so on).`,
				},
				resource.Attribute{
					Name:        "nested_virtualization",
					Description: `If true, nested virtualization is enabled. If false, it is disabled.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net in which the VM is running.`,
				},
				resource.Attribute{
					Name:        "nics",
					Description: `(Net only) The network interface cards (NICs) the VMs are attached to.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The account ID of the owner of the NIC.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the NIC.`,
				},
				resource.Attribute{
					Name:        "is_source_dest_checked",
					Description: `(Net only) If true, the source/destination check is enabled. If false, it is disabled. This value must be false for a NAT VM to perform network address translation (NAT) in a Net.`,
				},
				resource.Attribute{
					Name:        "link_nic",
					Description: `Information about the network interface card (NIC).`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If true, the NIC is deleted when the VM is terminated.`,
				},
				resource.Attribute{
					Name:        "device_number",
					Description: `The device index for the NIC attachment (between ` + "`" + `1` + "`" + ` and ` + "`" + `7` + "`" + `, both included).`,
				},
				resource.Attribute{
					Name:        "link_nic_id",
					Description: `The ID of the NIC to attach.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment (` + "`" + `attaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "link_public_ip",
					Description: `Information about the public IP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_ip_account_id",
					Description: `The account ID of the owner of the public IP.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `The Media Access Control (MAC) address of the NIC.`,
				},
				resource.Attribute{
					Name:        "net_id",
					Description: `The ID of the Net for the NIC.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `The ID of the NIC.`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `The private IP or IPs of the NIC.`,
				},
				resource.Attribute{
					Name:        "is_primary",
					Description: `If true, the IP is the primary private IP of the NIC.`,
				},
				resource.Attribute{
					Name:        "link_public_ip",
					Description: `Information about the public IP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP associated with the NIC.`,
				},
				resource.Attribute{
					Name:        "public_ip_account_id",
					Description: `The account ID of the owner of the public IP.`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `One or more IDs of security groups for the NIC.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the NIC (` + "`" + `available` + "`" + ` \| ` + "`" + `attaching` + "`" + ` \| ` + "`" + `in-use` + "`" + ` \| ` + "`" + `detaching` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet for the NIC.`,
				},
				resource.Attribute{
					Name:        "os_family",
					Description: `Indicates the operating system (OS) of the VM.`,
				},
				resource.Attribute{
					Name:        "performance",
					Description: `The performance of the VM (` + "`" + `medium` + "`" + ` \| ` + "`" + `high` + "`" + ` \| ` + "`" + `highest` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "placement",
					Description: `Information about the placement of the VM.`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The name of the Subregion. If you specify this parameter, you must not specify the ` + "`" + `nics` + "`" + ` parameter.`,
				},
				resource.Attribute{
					Name:        "tenancy",
					Description: `The tenancy of the VM (` + "`" + `default` + "`" + ` \| ` + "`" + `dedicated` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The name of the private DNS.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The primary private IP of the VM.`,
				},
				resource.Attribute{
					Name:        "product_codes",
					Description: `The product code associated with the OMI used to create the VM (` + "`" + `0001` + "`" + ` Linux/Unix \| ` + "`" + `0002` + "`" + ` Windows \| ` + "`" + `0004` + "`" + ` Linux/Oracle \| ` + "`" + `0005` + "`" + ` Windows 10).`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The name of the public DNS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP of the VM.`,
				},
				resource.Attribute{
					Name:        "reservation_id",
					Description: `The reservation ID of the VM.`,
				},
				resource.Attribute{
					Name:        "root_device_name",
					Description: `The name of the root device for the VM (for example, ` + "`" + `/dev/vda1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "root_device_type",
					Description: `The type of root device used by the VM (always ` + "`" + `bsu` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `One or more security groups associated with the VM.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the VM (` + "`" + `pending` + "`" + ` \| ` + "`" + `running` + "`" + ` \| ` + "`" + `stopping` + "`" + ` \| ` + "`" + `stopped` + "`" + ` \| ` + "`" + `shutting-down` + "`" + ` \| ` + "`" + `terminated` + "`" + ` \| ` + "`" + `quarantine` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state_reason",
					Description: `The reason explaining the current state of the VM.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the Subnet for the VM.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the VM.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `The Base64-encoded MIME user data.`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM.`,
				},
				resource.Attribute{
					Name:        "vm_initiated_shutdown_behavior",
					Description: `The VM behavior when you stop it. If set to ` + "`" + `stop` + "`" + `, the VM stops. If set to ` + "`" + `restart` + "`" + `, the VM stops then automatically restarts. If set to ` + "`" + `terminate` + "`" + `, the VM stops and is deleted.`,
				},
				resource.Attribute{
					Name:        "vm_type",
					Description: `The type of VM. For more information, see [Instance Types](https://docs.outscale.com/en/userguide/Instance-Types.html).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_volume",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a volume.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "creation_dates",
					Description: `(Optional) The dates and times of creation of the volumes.`,
				},
				resource.Attribute{
					Name:        "link_volume_delete_on_vm_deletion",
					Description: `(Optional) Whether the volumes are deleted or not when terminating the VMs.`,
				},
				resource.Attribute{
					Name:        "link_volume_device_names",
					Description: `(Optional) The VM device names.`,
				},
				resource.Attribute{
					Name:        "link_volume_link_dates",
					Description: `(Optional) The dates and times of creation of the volumes.`,
				},
				resource.Attribute{
					Name:        "link_volume_link_states",
					Description: `(Optional) The attachment states of the volumes (` + "`" + `attaching` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "link_volume_vm_ids",
					Description: `(Optional) One or more IDs of VMs.`,
				},
				resource.Attribute{
					Name:        "snapshot_ids",
					Description: `(Optional) The snapshots from which the volumes were created.`,
				},
				resource.Attribute{
					Name:        "subregion_names",
					Description: `(Optional) The names of the Subregions in which the volumes were created.`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the volumes.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the volumes.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combinations of the tags associated with the volumes, in the following format: ` + "`" + `TAGKEY=TAGVALUE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "volume_ids",
					Description: `(Optional) The IDs of the volumes.`,
				},
				resource.Attribute{
					Name:        "volume_sizes",
					Description: `(Optional) The sizes of the volumes, in gibibytes (GiB).`,
				},
				resource.Attribute{
					Name:        "volume_states",
					Description: `(Optional) The states of the volumes (` + "`" + `creating` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `in-use` + "`" + ` \| ` + "`" + `updating` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `error` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "volume_types",
					Description: `(Optional) The types of the volumes (` + "`" + `standard` + "`" + ` \| ` + "`" + `gp2` + "`" + ` \| ` + "`" + `io1` + "`" + `). ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `The date and time of creation of the volume.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `The number of I/O operations per second (IOPS):<br />- For ` + "`" + `io1` + "`" + ` volumes, the number of provisioned IOPS.<br />- For ` + "`" + `gp2` + "`" + ` volumes, the baseline performance of the volume.`,
				},
				resource.Attribute{
					Name:        "linked_volumes",
					Description: `Information about your volume attachment.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If true, the volume is deleted when terminating the VM. If false, the volume is not deleted when terminating the VM.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `The name of the device.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment of the volume (` + "`" + `attaching` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The ID of the volume.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the volume, in gibibytes (GiB).`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The snapshot from which the volume was created.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the volume (` + "`" + `creating` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `in-use` + "`" + ` \| ` + "`" + `updating` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `error` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The Subregion in which the volume was created.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the volume.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The ID of the volume.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `The type of the volume (` + "`" + `standard` + "`" + ` \| ` + "`" + `gp2` + "`" + ` \| ` + "`" + `io1` + "`" + `).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "creation_date",
					Description: `The date and time of creation of the volume.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `The number of I/O operations per second (IOPS):<br />- For ` + "`" + `io1` + "`" + ` volumes, the number of provisioned IOPS.<br />- For ` + "`" + `gp2` + "`" + ` volumes, the baseline performance of the volume.`,
				},
				resource.Attribute{
					Name:        "linked_volumes",
					Description: `Information about your volume attachment.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If true, the volume is deleted when terminating the VM. If false, the volume is not deleted when terminating the VM.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `The name of the device.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment of the volume (` + "`" + `attaching` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The ID of the volume.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the volume, in gibibytes (GiB).`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The snapshot from which the volume was created.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the volume (` + "`" + `creating` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `in-use` + "`" + ` \| ` + "`" + `updating` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `error` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The Subregion in which the volume was created.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the volume.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The ID of the volume.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `The type of the volume (` + "`" + `standard` + "`" + ` \| ` + "`" + `gp2` + "`" + ` \| ` + "`" + `io1` + "`" + `).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_volumes",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about volumes.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "creation_dates",
					Description: `(Optional) The dates and times of creation of the volumes.`,
				},
				resource.Attribute{
					Name:        "link_volume_delete_on_vm_deletion",
					Description: `(Optional) Whether the volumes are deleted or not when terminating the VMs.`,
				},
				resource.Attribute{
					Name:        "link_volume_device_names",
					Description: `(Optional) The VM device names.`,
				},
				resource.Attribute{
					Name:        "link_volume_link_dates",
					Description: `(Optional) The dates and times of creation of the volumes.`,
				},
				resource.Attribute{
					Name:        "link_volume_link_states",
					Description: `(Optional) The attachment states of the volumes (` + "`" + `attaching` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "link_volume_vm_ids",
					Description: `(Optional) One or more IDs of VMs.`,
				},
				resource.Attribute{
					Name:        "snapshot_ids",
					Description: `(Optional) The snapshots from which the volumes were created.`,
				},
				resource.Attribute{
					Name:        "subregion_names",
					Description: `(Optional) The names of the Subregions in which the volumes were created.`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the volumes.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the volumes.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combinations of the tags associated with the volumes, in the following format: ` + "`" + `TAGKEY=TAGVALUE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "volume_ids",
					Description: `(Optional) The IDs of the volumes.`,
				},
				resource.Attribute{
					Name:        "volume_sizes",
					Description: `(Optional) The sizes of the volumes, in gibibytes (GiB).`,
				},
				resource.Attribute{
					Name:        "volume_states",
					Description: `(Optional) The states of the volumes (` + "`" + `creating` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `in-use` + "`" + ` \| ` + "`" + `updating` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `error` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "volume_types",
					Description: `(Optional) The types of the volumes (` + "`" + `standard` + "`" + ` \| ` + "`" + `gp2` + "`" + ` \| ` + "`" + `io1` + "`" + `). ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "volumes",
					Description: `Information about one or more volumes.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `The date and time of creation of the volume.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `The number of I/O operations per second (IOPS):<br />- For ` + "`" + `io1` + "`" + ` volumes, the number of provisioned IOPS.<br />- For ` + "`" + `gp2` + "`" + ` volumes, the baseline performance of the volume.`,
				},
				resource.Attribute{
					Name:        "linked_volumes",
					Description: `Information about your volume attachment.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If true, the volume is deleted when terminating the VM. If false, the volume is not deleted when terminating the VM.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `The name of the device.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment of the volume (` + "`" + `attaching` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The ID of the volume.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the volume, in gibibytes (GiB).`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The snapshot from which the volume was created.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the volume (` + "`" + `creating` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `in-use` + "`" + ` \| ` + "`" + `updating` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `error` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The Subregion in which the volume was created.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the volume.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The ID of the volume.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `The type of the volume (` + "`" + `standard` + "`" + ` \| ` + "`" + `gp2` + "`" + ` \| ` + "`" + `io1` + "`" + `).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "volumes",
					Description: `Information about one or more volumes.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `The date and time of creation of the volume.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `The number of I/O operations per second (IOPS):<br />- For ` + "`" + `io1` + "`" + ` volumes, the number of provisioned IOPS.<br />- For ` + "`" + `gp2` + "`" + ` volumes, the baseline performance of the volume.`,
				},
				resource.Attribute{
					Name:        "linked_volumes",
					Description: `Information about your volume attachment.`,
				},
				resource.Attribute{
					Name:        "delete_on_vm_deletion",
					Description: `If true, the volume is deleted when terminating the VM. If false, the volume is not deleted when terminating the VM.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `The name of the device.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the attachment of the volume (` + "`" + `attaching` + "`" + ` \| ` + "`" + `detaching` + "`" + ` \| ` + "`" + `attached` + "`" + ` \| ` + "`" + `detached` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "vm_id",
					Description: `The ID of the VM.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The ID of the volume.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the volume, in gibibytes (GiB).`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The snapshot from which the volume was created.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the volume (` + "`" + `creating` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `in-use` + "`" + ` \| ` + "`" + `updating` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `error` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subregion_name",
					Description: `The Subregion in which the volume was created.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the volume.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The ID of the volume.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `The type of the volume (` + "`" + `standard` + "`" + ` \| ` + "`" + `gp2` + "`" + ` \| ` + "`" + `io1` + "`" + `).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_vpn_connection",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about a VPN connection.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "bgp_asns",
					Description: `(Optional) The Border Gateway Protocol (BGP) Autonomous System Numbers (ASNs) of the connections.`,
				},
				resource.Attribute{
					Name:        "client_gateway_ids",
					Description: `(Optional) The IDs of the client gateways.`,
				},
				resource.Attribute{
					Name:        "connection_types",
					Description: `(Optional) The types of the VPN connections (only ` + "`" + `ipsec.1` + "`" + ` is supported).`,
				},
				resource.Attribute{
					Name:        "route_destination_ip_ranges",
					Description: `(Optional) The destination IP ranges.`,
				},
				resource.Attribute{
					Name:        "states",
					Description: `(Optional) The states of the VPN connections (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "static_routes_only",
					Description: `(Optional) If false, the VPN connection uses dynamic routing with Border Gateway Protocol (BGP). If true, routing is controlled using static routes. For more information about how to create and delete static routes, see [CreateVpnConnectionRoute](https://docs.outscale.com/api#createvpnconnectionroute) and [DeleteVpnConnectionRoute](https://docs.outscale.com/api#deletevpnconnectionroute).`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the VPN connections.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the VPN connections.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combinations of the tags associated with the VPN connections, in the following format: ` + "`" + `TAGKEY=TAGVALUE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_ids",
					Description: `(Optional) The IDs of the virtual gateways.`,
				},
				resource.Attribute{
					Name:        "vpn_connection_ids",
					Description: `(Optional) The IDs of the VPN connections. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "client_gateway_configuration",
					Description: `Example configuration for the client gateway.`,
				},
				resource.Attribute{
					Name:        "client_gateway_id",
					Description: `The ID of the client gateway used on the client end of the connection.`,
				},
				resource.Attribute{
					Name:        "connection_type",
					Description: `The type of VPN connection (always ` + "`" + `ipsec.1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "routes",
					Description: `Information about one or more static routes associated with the VPN connection, if any.`,
				},
				resource.Attribute{
					Name:        "destination_ip_range",
					Description: `The IP range used for the destination match, in CIDR notation (for example, ` + "`" + `10.0.0.0/24` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "route_type",
					Description: `The type of route (always ` + "`" + `static` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The current state of the static route (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the VPN connection (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "static_routes_only",
					Description: `If false, the VPN connection uses dynamic routing with Border Gateway Protocol (BGP). If true, routing is controlled using static routes. For more information about how to create and delete static routes, see [CreateVpnConnectionRoute](https://docs.outscale.com/api#createvpnconnectionroute) and [DeleteVpnConnectionRoute](https://docs.outscale.com/api#deletevpnconnectionroute).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the VPN connection.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "vgw_telemetries",
					Description: `Information about the current state of one or more of the VPN tunnels.`,
				},
				resource.Attribute{
					Name:        "accepted_route_count",
					Description: `The number of routes accepted through BGP (Border Gateway Protocol) route exchanges.`,
				},
				resource.Attribute{
					Name:        "last_state_change_date",
					Description: `The date and time (UTC) of the latest state update.`,
				},
				resource.Attribute{
					Name:        "outside_ip_address",
					Description: `The IP on the OUTSCALE side of the tunnel.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the IPSEC tunnel (` + "`" + `UP` + "`" + ` \| ` + "`" + `DOWN` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state_description",
					Description: `A description of the current state of the tunnel.`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_id",
					Description: `The ID of the virtual gateway used on the OUTSCALE end of the connection.`,
				},
				resource.Attribute{
					Name:        "vpn_connection_id",
					Description: `The ID of the VPN connection.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "client_gateway_configuration",
					Description: `Example configuration for the client gateway.`,
				},
				resource.Attribute{
					Name:        "client_gateway_id",
					Description: `The ID of the client gateway used on the client end of the connection.`,
				},
				resource.Attribute{
					Name:        "connection_type",
					Description: `The type of VPN connection (always ` + "`" + `ipsec.1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "routes",
					Description: `Information about one or more static routes associated with the VPN connection, if any.`,
				},
				resource.Attribute{
					Name:        "destination_ip_range",
					Description: `The IP range used for the destination match, in CIDR notation (for example, ` + "`" + `10.0.0.0/24` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "route_type",
					Description: `The type of route (always ` + "`" + `static` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The current state of the static route (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the VPN connection (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "static_routes_only",
					Description: `If false, the VPN connection uses dynamic routing with Border Gateway Protocol (BGP). If true, routing is controlled using static routes. For more information about how to create and delete static routes, see [CreateVpnConnectionRoute](https://docs.outscale.com/api#createvpnconnectionroute) and [DeleteVpnConnectionRoute](https://docs.outscale.com/api#deletevpnconnectionroute).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the VPN connection.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "vgw_telemetries",
					Description: `Information about the current state of one or more of the VPN tunnels.`,
				},
				resource.Attribute{
					Name:        "accepted_route_count",
					Description: `The number of routes accepted through BGP (Border Gateway Protocol) route exchanges.`,
				},
				resource.Attribute{
					Name:        "last_state_change_date",
					Description: `The date and time (UTC) of the latest state update.`,
				},
				resource.Attribute{
					Name:        "outside_ip_address",
					Description: `The IP on the OUTSCALE side of the tunnel.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the IPSEC tunnel (` + "`" + `UP` + "`" + ` \| ` + "`" + `DOWN` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state_description",
					Description: `A description of the current state of the tunnel.`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_id",
					Description: `The ID of the virtual gateway used on the OUTSCALE end of the connection.`,
				},
				resource.Attribute{
					Name:        "vpn_connection_id",
					Description: `The ID of the VPN connection.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "outscale_vpn_connections",
			Category:         "Data Sources",
			ShortDescription: `[Provides information about VPN connections.]`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A combination of a filter name and one or more filter values. You can specify this argument for as many filter names as you need. The filter name can be any of the following:`,
				},
				resource.Attribute{
					Name:        "bgp_asns",
					Description: `(Optional) The Border Gateway Protocol (BGP) Autonomous System Numbers (ASNs) of the connections.`,
				},
				resource.Attribute{
					Name:        "client_gateway_ids",
					Description: `(Optional) The IDs of the client gateways.`,
				},
				resource.Attribute{
					Name:        "connection_types",
					Description: `(Optional) The types of the VPN connections (only ` + "`" + `ipsec.1` + "`" + ` is supported).`,
				},
				resource.Attribute{
					Name:        "route_destination_ip_ranges",
					Description: `(Optional) The destination IP ranges.`,
				},
				resource.Attribute{
					Name:        "states",
					Description: `(Optional) The states of the VPN connections (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "static_routes_only",
					Description: `(Optional) If false, the VPN connection uses dynamic routing with Border Gateway Protocol (BGP). If true, routing is controlled using static routes. For more information about how to create and delete static routes, see [CreateVpnConnectionRoute](https://docs.outscale.com/api#createvpnconnectionroute) and [DeleteVpnConnectionRoute](https://docs.outscale.com/api#deletevpnconnectionroute).`,
				},
				resource.Attribute{
					Name:        "tag_keys",
					Description: `(Optional) The keys of the tags associated with the VPN connections.`,
				},
				resource.Attribute{
					Name:        "tag_values",
					Description: `(Optional) The values of the tags associated with the VPN connections.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value combinations of the tags associated with the VPN connections, in the following format: ` + "`" + `TAGKEY=TAGVALUE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_ids",
					Description: `(Optional) The IDs of the virtual gateways.`,
				},
				resource.Attribute{
					Name:        "vpn_connection_ids",
					Description: `(Optional) The IDs of the VPN connections. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "vpn_connections",
					Description: `Information about one or more VPN connections.`,
				},
				resource.Attribute{
					Name:        "client_gateway_configuration",
					Description: `Example configuration for the client gateway.`,
				},
				resource.Attribute{
					Name:        "client_gateway_id",
					Description: `The ID of the client gateway used on the client end of the connection.`,
				},
				resource.Attribute{
					Name:        "connection_type",
					Description: `The type of VPN connection (always ` + "`" + `ipsec.1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "routes",
					Description: `Information about one or more static routes associated with the VPN connection, if any.`,
				},
				resource.Attribute{
					Name:        "destination_ip_range",
					Description: `The IP range used for the destination match, in CIDR notation (for example, ` + "`" + `10.0.0.0/24` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "route_type",
					Description: `The type of route (always ` + "`" + `static` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The current state of the static route (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the VPN connection (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "static_routes_only",
					Description: `If false, the VPN connection uses dynamic routing with Border Gateway Protocol (BGP). If true, routing is controlled using static routes. For more information about how to create and delete static routes, see [CreateVpnConnectionRoute](https://docs.outscale.com/api#createvpnconnectionroute) and [DeleteVpnConnectionRoute](https://docs.outscale.com/api#deletevpnconnectionroute).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the VPN connection.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "vgw_telemetries",
					Description: `Information about the current state of one or more of the VPN tunnels.`,
				},
				resource.Attribute{
					Name:        "accepted_route_count",
					Description: `The number of routes accepted through BGP (Border Gateway Protocol) route exchanges.`,
				},
				resource.Attribute{
					Name:        "last_state_change_date",
					Description: `The date and time (UTC) of the latest state update.`,
				},
				resource.Attribute{
					Name:        "outside_ip_address",
					Description: `The IP on the OUTSCALE side of the tunnel.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the IPSEC tunnel (` + "`" + `UP` + "`" + ` \| ` + "`" + `DOWN` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state_description",
					Description: `A description of the current state of the tunnel.`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_id",
					Description: `The ID of the virtual gateway used on the OUTSCALE end of the connection.`,
				},
				resource.Attribute{
					Name:        "vpn_connection_id",
					Description: `The ID of the VPN connection.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vpn_connections",
					Description: `Information about one or more VPN connections.`,
				},
				resource.Attribute{
					Name:        "client_gateway_configuration",
					Description: `Example configuration for the client gateway.`,
				},
				resource.Attribute{
					Name:        "client_gateway_id",
					Description: `The ID of the client gateway used on the client end of the connection.`,
				},
				resource.Attribute{
					Name:        "connection_type",
					Description: `The type of VPN connection (always ` + "`" + `ipsec.1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "routes",
					Description: `Information about one or more static routes associated with the VPN connection, if any.`,
				},
				resource.Attribute{
					Name:        "destination_ip_range",
					Description: `The IP range used for the destination match, in CIDR notation (for example, ` + "`" + `10.0.0.0/24` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "route_type",
					Description: `The type of route (always ` + "`" + `static` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The current state of the static route (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the VPN connection (` + "`" + `pending` + "`" + ` \| ` + "`" + `available` + "`" + ` \| ` + "`" + `deleting` + "`" + ` \| ` + "`" + `deleted` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "static_routes_only",
					Description: `If false, the VPN connection uses dynamic routing with Border Gateway Protocol (BGP). If true, routing is controlled using static routes. For more information about how to create and delete static routes, see [CreateVpnConnectionRoute](https://docs.outscale.com/api#createvpnconnectionroute) and [DeleteVpnConnectionRoute](https://docs.outscale.com/api#deletevpnconnectionroute).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `One or more tags associated with the VPN connection.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the tag, with a minimum of 1 character.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag, between 0 and 255 characters.`,
				},
				resource.Attribute{
					Name:        "vgw_telemetries",
					Description: `Information about the current state of one or more of the VPN tunnels.`,
				},
				resource.Attribute{
					Name:        "accepted_route_count",
					Description: `The number of routes accepted through BGP (Border Gateway Protocol) route exchanges.`,
				},
				resource.Attribute{
					Name:        "last_state_change_date",
					Description: `The date and time (UTC) of the latest state update.`,
				},
				resource.Attribute{
					Name:        "outside_ip_address",
					Description: `The IP on the OUTSCALE side of the tunnel.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the IPSEC tunnel (` + "`" + `UP` + "`" + ` \| ` + "`" + `DOWN` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state_description",
					Description: `A description of the current state of the tunnel.`,
				},
				resource.Attribute{
					Name:        "virtual_gateway_id",
					Description: `The ID of the virtual gateway used on the OUTSCALE end of the connection.`,
				},
				resource.Attribute{
					Name:        "vpn_connection_id",
					Description: `The ID of the VPN connection.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"outscale_access_key":                   0,
		"outscale_access_keys":                  1,
		"outscale_account":                      2,
		"outscale_accounts":                     3,
		"outscale_api_access_policy":            4,
		"outscale_api_access_rule":              5,
		"outscale_api_access_rules":             6,
		"outscale_ca":                           7,
		"outscale_cas":                          8,
		"outscale_client_gateway":               9,
		"outscale_client_gateways":              10,
		"outscale_dhcp_option":                  11,
		"outscale_dhcp_options":                 12,
		"outscale_flexible_gpu":                 13,
		"outscale_flexible_gpu_catalog":         14,
		"outscale_flexible_gpus":                15,
		"outscale_image":                        16,
		"outscale_image_export_task":            17,
		"outscale_image_export_tasks":           18,
		"outscale_images":                       19,
		"outscale_internet_service":             20,
		"outscale_internet_services":            21,
		"outscale_keypair":                      22,
		"outscale_keypairs":                     23,
		"outscale_load_balancer":                24,
		"outscale_load_balancer_listener_rule":  25,
		"outscale_load_balancer_listener_rules": 26,
		"outscale_load_balancer_vm_health":      27,
		"outscale_load_balancers":               28,
		"outscale_nat_service":                  29,
		"outscale_nat_services":                 30,
		"outscale_net":                          31,
		"outscale_net_access_point":             32,
		"outscale_net_access_point_services":    33,
		"outscale_net_access_points":            34,
		"outscale_net_attributes":               35,
		"outscale_net_peering":                  36,
		"outscale_net_peerings":                 37,
		"outscale_nets":                         38,
		"outscale_nic":                          39,
		"outscale_nics":                         40,
		"outscale_product_type":                 41,
		"outscale_product_types":                42,
		"outscale_public_ip":                    43,
		"outscale_public_ips":                   44,
		"outscale_quota":                        45,
		"outscale_quotas":                       46,
		"outscale_regions":                      47,
		"outscale_route_table":                  48,
		"outscale_route_tables":                 49,
		"outscale_security_group":               50,
		"outscale_security_groups":              51,
		"outscale_server_certificate":           52,
		"outscale_server_certificates":          53,
		"outscale_snapshot":                     54,
		"outscale_snapshot_export_task":         55,
		"outscale_snapshot_export_tasks":        56,
		"outscale_snapshots":                    57,
		"outscale_subnet":                       58,
		"outscale_subnets":                      59,
		"outscale_subregions":                   60,
		"outscale_virtual_gateway":              61,
		"outscale_virtual_gateways":             62,
		"outscale_vm":                           63,
		"outscale_vm_state":                     64,
		"outscale_vm_states":                    65,
		"outscale_vm_types":                     66,
		"outscale_vms":                          67,
		"outscale_volume":                       68,
		"outscale_volumes":                      69,
		"outscale_vpn_connection":               70,
		"outscale_vpn_connections":              71,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
