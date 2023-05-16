package fastly

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "fastly_datacenters",
			Category:         "Data Sources",
			ShortDescription: `Get information on Fastly datacenters.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fastly_dictionaries",
			Category:         "Data Sources",
			ShortDescription: `Get information on Fastly service dictionaries.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
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
			Type:             "fastly_services",
			Category:         "Data Sources",
			ShortDescription: `Get information on Fastly services.`,
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
					Description: `The rule's type. [1]: https://developer.fastly.com/reference/api/waf/rules/ <!-- schema generated by tfplugindocs --> ## Schema ### Optional - ` + "`" + `exclude_modsec_rule_ids` + "`" + ` (List of Number) A list of modsecurity rules IDs to be excluded from the data set. - ` + "`" + `modsec_rule_ids` + "`" + ` (List of Number) A list of modsecurity rules IDs to be used as filters for the data set. - ` + "`" + `publishers` + "`" + ` (List of String) A list of publishers to be used as filters for the data set. - ` + "`" + `tags` + "`" + ` (List of String) A list of tags to be used as filters for the data set. ### Read-Only - ` + "`" + `id` + "`" + ` (String) The ID of this resource. - ` + "`" + `rules` + "`" + ` (List of Object) The list of rules that results from any given combination of filters. (see [below for nested schema](#nestedatt--rules)) <a id="nestedatt--rules"></a> ### Nested Schema for ` + "`" + `rules` + "`" + ` Read-Only: - ` + "`" + `latest_revision_number` + "`" + ` (Number) - ` + "`" + `modsec_rule_id` + "`" + ` (Number) - ` + "`" + `type` + "`" + ` (String)`,
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
					Description: `The rule's type. [1]: https://developer.fastly.com/reference/api/waf/rules/ <!-- schema generated by tfplugindocs --> ## Schema ### Optional - ` + "`" + `exclude_modsec_rule_ids` + "`" + ` (List of Number) A list of modsecurity rules IDs to be excluded from the data set. - ` + "`" + `modsec_rule_ids` + "`" + ` (List of Number) A list of modsecurity rules IDs to be used as filters for the data set. - ` + "`" + `publishers` + "`" + ` (List of String) A list of publishers to be used as filters for the data set. - ` + "`" + `tags` + "`" + ` (List of String) A list of tags to be used as filters for the data set. ### Read-Only - ` + "`" + `id` + "`" + ` (String) The ID of this resource. - ` + "`" + `rules` + "`" + ` (List of Object) The list of rules that results from any given combination of filters. (see [below for nested schema](#nestedatt--rules)) <a id="nestedatt--rules"></a> ### Nested Schema for ` + "`" + `rules` + "`" + ` Read-Only: - ` + "`" + `latest_revision_number` + "`" + ` (Number) - ` + "`" + `modsec_rule_id` + "`" + ` (Number) - ` + "`" + `type` + "`" + ` (String)`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"fastly_datacenters":                  0,
		"fastly_dictionaries":                 1,
		"fastly_ip_ranges":                    2,
		"fastly_services":                     3,
		"fastly_tls_activation":               4,
		"fastly_tls_activation_ids":           5,
		"fastly_tls_certificate":              6,
		"fastly_tls_certificate_ids":          7,
		"fastly_tls_configuration":            8,
		"fastly_tls_configuration_ids":        9,
		"fastly_tls_domain":                   10,
		"fastly_tls_platform_certificate":     11,
		"fastly_tls_platform_certificate_ids": 12,
		"fastly_tls_private_key":              13,
		"fastly_tls_private_key_ids":          14,
		"fastly_tls_subscription":             15,
		"fastly_tls_subscription_ids":         16,
		"fastly_waf_rules":                    17,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
