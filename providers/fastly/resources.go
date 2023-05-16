package fastly

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "fastly_service_acl_entries",
			Category:         "Resources",
			ShortDescription: `Defines a set of Fastly ACL entries that can be used to populate a service ACL.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fastly_service_authorization",
			Category:         "Resources",
			ShortDescription: `Grants user access to service`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fastly_service_compute",
			Category:         "Resources",
			ShortDescription: `Provides an Fastly Compute@Edge service`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fastly_service_dictionary_items",
			Category:         "Resources",
			ShortDescription: `Provides a grouping of Fastly dictionary items that can be applied to a service.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fastly_service_dynamic_snippet_content",
			Category:         "Resources",
			ShortDescription: `Provides a means to define blocks of VCL logic that is inserted into your service through Fastly dynamic snippets.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fastly_service_vcl",
			Category:         "Resources",
			ShortDescription: `Provides an Fastly Service`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fastly_service_waf_configuration",
			Category:         "Resources",
			ShortDescription: `Provides a Web Application Firewall configuration and rules that can be applied to a service.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fastly_tls_activation",
			Category:         "Resources",
			ShortDescription: `Enables TLS on a domain`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fastly_tls_certificate",
			Category:         "Resources",
			ShortDescription: `Uploads a custom TLS certificate`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fastly_tls_platform_certificate",
			Category:         "Resources",
			ShortDescription: `Uploads a TLS certificate to the Platform TLS service`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fastly_tls_private_key",
			Category:         "Resources",
			ShortDescription: `Uploads a Custom TLS Private Key`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fastly_tls_subscription",
			Category:         "Resources",
			ShortDescription: `Enables TLS on a domain using a managed certificate`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domains",
					Description: `(Required) List of domains on which to enable TLS.`,
				},
				resource.Attribute{
					Name:        "certificate_authority",
					Description: `(Required) The entity that issues and certifies the TLS certificates for your subscription. Valid values are ` + "`" + `lets-encrypt` + "`" + ` or ` + "`" + `globalsign` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "configuration_id",
					Description: `(Optional) The ID of the set of TLS configuration options that apply to the enabled domains on this subscription.`,
				},
				resource.Attribute{
					Name:        "force_update",
					Description: `(Optional) Always update subscription, even when active domains are present. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "force_destroy",
					Description: `(Optional) Always delete subscription, even when active domains are present. Defaults to false. !>`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Timestamp (GMT) when the subscription was created.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Timestamp (GMT) when the subscription was last updated.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The current state of the subscription. The list of possible states are: ` + "`" + `pending` + "`" + `, ` + "`" + `processing` + "`" + `, ` + "`" + `issued` + "`" + `, and ` + "`" + `renewing` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "managed_dns_challenges",
					Description: `A list of options for configuring DNS to respond to ACME DNS challenge in order to verify domain ownership. See Managed DNS Challenge below for details.`,
				},
				resource.Attribute{
					Name:        "managed_http_challenges",
					Description: `A list of options for configuring DNS to respond to ACME HTTP challenge in order to verify domain ownership. See Managed HTTP Challenges below for details. ### Managed DNS Challenge The available attributes in the ` + "`" + `managed_dns_challenges` + "`" + ` block are:`,
				},
				resource.Attribute{
					Name:        "record_name",
					Description: `The name of the DNS record to add. For example ` + "`" + `_acme-challenge.example.com` + "`" + `. Accessed like this, ` + "`" + `fastly_tls_subscription.tls.managed_dns_challenges.record_name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "record_type",
					Description: `The type of DNS record to add, e.g. ` + "`" + `A` + "`" + `, or ` + "`" + `CNAME` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "record_value",
					Description: `The value to which the DNS record should point, e.g. ` + "`" + `xxxxx.fastly-validations.com` + "`" + `. ### Managed HTTP Challenges The ` + "`" + `managed_http_challenges` + "`" + ` attribute is a set of different records that could be added depending on requirements. For example, whether you are adding TLS to an apex domain, or a subdomain will determine which record you require. Please note that these records will redirect production traffic to Fastly, so make sure the service is configured correctly first. Each record in the set has the following attributes:`,
				},
				resource.Attribute{
					Name:        "record_name",
					Description: `The name of the DNS record to add. For example ` + "`" + `example.com` + "`" + `. Best accessed through a ` + "`" + `for` + "`" + ` expression to filter the relevant record.`,
				},
				resource.Attribute{
					Name:        "record_type",
					Description: `The type of DNS record to add, e.g. ` + "`" + `A` + "`" + `, or ` + "`" + `CNAME` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "record_values",
					Description: `A list with the value(s) to which the DNS record should point. ## Import A subscription can be imported using its Fastly subscription ID, e.g. ` + "`" + `` + "`" + `` + "`" + `sh $ terraform import fastly_tls_subscription.demo xxxxxxxxxxx ` + "`" + `` + "`" + `` + "`" + ` <!-- schema generated by tfplugindocs --> ## Schema ### Required - ` + "`" + `certificate_authority` + "`" + ` (String) The entity that issues and certifies the TLS certificates for your subscription. Valid values are ` + "`" + `lets-encrypt` + "`" + `, ` + "`" + `globalsign` + "`" + ` or ` + "`" + `certainly` + "`" + ` (beta). - ` + "`" + `domains` + "`" + ` (Set of String) List of domains on which to enable TLS. ### Optional - ` + "`" + `common_name` + "`" + ` (String) The common name associated with the subscription generated by Fastly TLS. If you do not pass a common name on create, we will default to the first TLS domain included. If provided, the domain chosen as the common name must be included in TLS domains. - ` + "`" + `configuration_id` + "`" + ` (String) The ID of the set of TLS configuration options that apply to the enabled domains on this subscription. - ` + "`" + `force_destroy` + "`" + ` (Boolean) Force delete the subscription even if it has active domains. Warning: this can disable production traffic if used incorrectly. Defaults to false. - ` + "`" + `force_update` + "`" + ` (Boolean) Force update the subscription even if it has active domains. Warning: this can disable production traffic if used incorrectly. ### Read-Only - ` + "`" + `certificate_id` + "`" + ` (String) The certificate ID associated with the subscription. - ` + "`" + `created_at` + "`" + ` (String) Timestamp (GMT) when the subscription was created. - ` + "`" + `id` + "`" + ` (String) The ID of this resource. - ` + "`" + `managed_dns_challenge` + "`" + ` (Map of String, Deprecated) The details required to configure DNS to respond to ACME DNS challenge in order to verify domain ownership. - ` + "`" + `managed_dns_challenges` + "`" + ` (Set of Object) A list of options for configuring DNS to respond to ACME DNS challenge in order to verify domain ownership. (see [below for nested schema](#nestedatt--managed_dns_challenges)) - ` + "`" + `managed_http_challenges` + "`" + ` (Set of Object) A list of options for configuring DNS to respond to ACME HTTP challenge in order to verify domain ownership. Best accessed through a ` + "`" + `for` + "`" + ` expression to filter the relevant record. (see [below for nested schema](#nestedatt--managed_http_challenges)) - ` + "`" + `state` + "`" + ` (String) The current state of the subscription. The list of possible states are: ` + "`" + `pending` + "`" + `, ` + "`" + `processing` + "`" + `, ` + "`" + `issued` + "`" + `, and ` + "`" + `renewing` + "`" + `. - ` + "`" + `updated_at` + "`" + ` (String) Timestamp (GMT) when the subscription was updated. <a id="nestedatt--managed_dns_challenges"></a> ### Nested Schema for ` + "`" + `managed_dns_challenges` + "`" + ` Read-Only: - ` + "`" + `record_name` + "`" + ` (String) - ` + "`" + `record_type` + "`" + ` (String) - ` + "`" + `record_value` + "`" + ` (String) <a id="nestedatt--managed_http_challenges"></a> ### Nested Schema for ` + "`" + `managed_http_challenges` + "`" + ` Read-Only: - ` + "`" + `record_name` + "`" + ` (String) - ` + "`" + `record_type` + "`" + ` (String) - ` + "`" + `record_values` + "`" + ` (Set of String)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `Timestamp (GMT) when the subscription was created.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Timestamp (GMT) when the subscription was last updated.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The current state of the subscription. The list of possible states are: ` + "`" + `pending` + "`" + `, ` + "`" + `processing` + "`" + `, ` + "`" + `issued` + "`" + `, and ` + "`" + `renewing` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "managed_dns_challenges",
					Description: `A list of options for configuring DNS to respond to ACME DNS challenge in order to verify domain ownership. See Managed DNS Challenge below for details.`,
				},
				resource.Attribute{
					Name:        "managed_http_challenges",
					Description: `A list of options for configuring DNS to respond to ACME HTTP challenge in order to verify domain ownership. See Managed HTTP Challenges below for details. ### Managed DNS Challenge The available attributes in the ` + "`" + `managed_dns_challenges` + "`" + ` block are:`,
				},
				resource.Attribute{
					Name:        "record_name",
					Description: `The name of the DNS record to add. For example ` + "`" + `_acme-challenge.example.com` + "`" + `. Accessed like this, ` + "`" + `fastly_tls_subscription.tls.managed_dns_challenges.record_name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "record_type",
					Description: `The type of DNS record to add, e.g. ` + "`" + `A` + "`" + `, or ` + "`" + `CNAME` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "record_value",
					Description: `The value to which the DNS record should point, e.g. ` + "`" + `xxxxx.fastly-validations.com` + "`" + `. ### Managed HTTP Challenges The ` + "`" + `managed_http_challenges` + "`" + ` attribute is a set of different records that could be added depending on requirements. For example, whether you are adding TLS to an apex domain, or a subdomain will determine which record you require. Please note that these records will redirect production traffic to Fastly, so make sure the service is configured correctly first. Each record in the set has the following attributes:`,
				},
				resource.Attribute{
					Name:        "record_name",
					Description: `The name of the DNS record to add. For example ` + "`" + `example.com` + "`" + `. Best accessed through a ` + "`" + `for` + "`" + ` expression to filter the relevant record.`,
				},
				resource.Attribute{
					Name:        "record_type",
					Description: `The type of DNS record to add, e.g. ` + "`" + `A` + "`" + `, or ` + "`" + `CNAME` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "record_values",
					Description: `A list with the value(s) to which the DNS record should point. ## Import A subscription can be imported using its Fastly subscription ID, e.g. ` + "`" + `` + "`" + `` + "`" + `sh $ terraform import fastly_tls_subscription.demo xxxxxxxxxxx ` + "`" + `` + "`" + `` + "`" + ` <!-- schema generated by tfplugindocs --> ## Schema ### Required - ` + "`" + `certificate_authority` + "`" + ` (String) The entity that issues and certifies the TLS certificates for your subscription. Valid values are ` + "`" + `lets-encrypt` + "`" + `, ` + "`" + `globalsign` + "`" + ` or ` + "`" + `certainly` + "`" + ` (beta). - ` + "`" + `domains` + "`" + ` (Set of String) List of domains on which to enable TLS. ### Optional - ` + "`" + `common_name` + "`" + ` (String) The common name associated with the subscription generated by Fastly TLS. If you do not pass a common name on create, we will default to the first TLS domain included. If provided, the domain chosen as the common name must be included in TLS domains. - ` + "`" + `configuration_id` + "`" + ` (String) The ID of the set of TLS configuration options that apply to the enabled domains on this subscription. - ` + "`" + `force_destroy` + "`" + ` (Boolean) Force delete the subscription even if it has active domains. Warning: this can disable production traffic if used incorrectly. Defaults to false. - ` + "`" + `force_update` + "`" + ` (Boolean) Force update the subscription even if it has active domains. Warning: this can disable production traffic if used incorrectly. ### Read-Only - ` + "`" + `certificate_id` + "`" + ` (String) The certificate ID associated with the subscription. - ` + "`" + `created_at` + "`" + ` (String) Timestamp (GMT) when the subscription was created. - ` + "`" + `id` + "`" + ` (String) The ID of this resource. - ` + "`" + `managed_dns_challenge` + "`" + ` (Map of String, Deprecated) The details required to configure DNS to respond to ACME DNS challenge in order to verify domain ownership. - ` + "`" + `managed_dns_challenges` + "`" + ` (Set of Object) A list of options for configuring DNS to respond to ACME DNS challenge in order to verify domain ownership. (see [below for nested schema](#nestedatt--managed_dns_challenges)) - ` + "`" + `managed_http_challenges` + "`" + ` (Set of Object) A list of options for configuring DNS to respond to ACME HTTP challenge in order to verify domain ownership. Best accessed through a ` + "`" + `for` + "`" + ` expression to filter the relevant record. (see [below for nested schema](#nestedatt--managed_http_challenges)) - ` + "`" + `state` + "`" + ` (String) The current state of the subscription. The list of possible states are: ` + "`" + `pending` + "`" + `, ` + "`" + `processing` + "`" + `, ` + "`" + `issued` + "`" + `, and ` + "`" + `renewing` + "`" + `. - ` + "`" + `updated_at` + "`" + ` (String) Timestamp (GMT) when the subscription was updated. <a id="nestedatt--managed_dns_challenges"></a> ### Nested Schema for ` + "`" + `managed_dns_challenges` + "`" + ` Read-Only: - ` + "`" + `record_name` + "`" + ` (String) - ` + "`" + `record_type` + "`" + ` (String) - ` + "`" + `record_value` + "`" + ` (String) <a id="nestedatt--managed_http_challenges"></a> ### Nested Schema for ` + "`" + `managed_http_challenges` + "`" + ` Read-Only: - ` + "`" + `record_name` + "`" + ` (String) - ` + "`" + `record_type` + "`" + ` (String) - ` + "`" + `record_values` + "`" + ` (Set of String)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fastly_tls_subscription_validation",
			Category:         "Resources",
			ShortDescription: `Represents a successful validation of a Fastly TLS Subscription`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fastly_user",
			Category:         "Resources",
			ShortDescription: `Provides a Fastly User`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"fastly_service_acl_entries":             0,
		"fastly_service_authorization":           1,
		"fastly_service_compute":                 2,
		"fastly_service_dictionary_items":        3,
		"fastly_service_dynamic_snippet_content": 4,
		"fastly_service_vcl":                     5,
		"fastly_service_waf_configuration":       6,
		"fastly_tls_activation":                  7,
		"fastly_tls_certificate":                 8,
		"fastly_tls_platform_certificate":        9,
		"fastly_tls_private_key":                 10,
		"fastly_tls_subscription":                11,
		"fastly_tls_subscription_validation":     12,
		"fastly_user":                            13,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
