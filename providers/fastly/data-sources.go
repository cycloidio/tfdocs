package fastly

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "fastly_ip_ranges",
			Category:         "Data Sources",
			ShortDescription: `Get information on Fastly IP ranges.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fastly_tls_activation",
			Category:         "Data Sources",
			ShortDescription: `Get information on Fastly TLS Activation.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fastly_tls_activation_ids",
			Category:         "Data Sources",
			ShortDescription: `Get the list of TLS Activation identifiers in Fastly.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fastly_tls_certificate",
			Category:         "Data Sources",
			ShortDescription: `Get information on Fastly TLS certificate.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fastly_tls_certificate_ids",
			Category:         "Data Sources",
			ShortDescription: `Get IDs of available TLS certificates.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fastly_tls_configuration",
			Category:         "Data Sources",
			ShortDescription: `Get information on Fastly TLS configuration.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fastly_tls_configuration_ids",
			Category:         "Data Sources",
			ShortDescription: `Get IDs of available TLS Configurations.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fastly_tls_domain",
			Category:         "Data Sources",
			ShortDescription: `Get IDs of activations, certificates and subscriptions associated with a domain.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fastly_tls_platform_certificate",
			Category:         "Data Sources",
			ShortDescription: `Get information on Fastly Platform TLS certificate.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fastly_tls_platform_certificate_ids",
			Category:         "Data Sources",
			ShortDescription: `Get IDs of available Platform TLS certificates.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fastly_tls_private_key",
			Category:         "Data Sources",
			ShortDescription: `Get information on a TLS Private Key.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fastly_tls_private_key_ids",
			Category:         "Data Sources",
			ShortDescription: `Get the list of TLS private key identifiers in Fastly.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fastly_tls_subscription",
			Category:         "Data Sources",
			ShortDescription: `Get information on Fastly TLS subscription.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fastly_tls_subscription_ids",
			Category:         "Data Sources",
			ShortDescription: `Get the list of TLS Subscriptions in Fastly.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fastly_waf_rules",
			Category:         "Data Sources",
			ShortDescription: `Get information on Fastly WAF rules.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "publishers",
					Description: `Inclusion filter by WAF rule's publishers.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Inclusion filter by WAF rule's tags.`,
				},
				resource.Attribute{
					Name:        "exclude_modsec_rule_ids",
					Description: `Exclusion filter by WAF rule's ModSecurity ID. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `The Web Application Firewall's rules result set. ~>`,
				},
				resource.Attribute{
					Name:        "modsec_rule_id",
					Description: `The rule's modsecurity ID.`,
				},
				resource.Attribute{
					Name:        "latest_revision_number",
					Description: `The rule's latest revision.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The rule's type. [1]: https://developer.fastly.com/reference/api/waf/rules/ <!-- schema generated by tfplugindocs --> ## Schema ### Optional -`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "rules",
					Description: `The Web Application Firewall's rules result set. ~>`,
				},
				resource.Attribute{
					Name:        "modsec_rule_id",
					Description: `The rule's modsecurity ID.`,
				},
				resource.Attribute{
					Name:        "latest_revision_number",
					Description: `The rule's latest revision.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The rule's type. [1]: https://developer.fastly.com/reference/api/waf/rules/ <!-- schema generated by tfplugindocs --> ## Schema ### Optional -`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"fastly_ip_ranges":                    0,
		"fastly_tls_activation":               1,
		"fastly_tls_activation_ids":           2,
		"fastly_tls_certificate":              3,
		"fastly_tls_certificate_ids":          4,
		"fastly_tls_configuration":            5,
		"fastly_tls_configuration_ids":        6,
		"fastly_tls_domain":                   7,
		"fastly_tls_platform_certificate":     8,
		"fastly_tls_platform_certificate_ids": 9,
		"fastly_tls_private_key":              10,
		"fastly_tls_private_key_ids":          11,
		"fastly_tls_subscription":             12,
		"fastly_tls_subscription_ids":         13,
		"fastly_waf_rules":                    14,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
