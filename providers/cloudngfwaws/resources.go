package cloudngfwaws

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "cloudngfwaws_certificate",
			Category:         "Resources",
			ShortDescription: `Resource for certificate manipulation.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudngfwaws_commit_rulestack",
			Category:         "Resources",
			ShortDescription: `Resource for committing the rulestack config.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudngfwaws_custom_url_category",
			Category:         "Resources",
			ShortDescription: `Resource for custom url category manipulation.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudngfwaws_fqdn_list",
			Category:         "Resources",
			ShortDescription: `Resource for fqdn list manipulation.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudngfwaws_intelligent_feed",
			Category:         "Resources",
			ShortDescription: `Resource for intelligent feed manipulation.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudngfwaws_ngfw",
			Category:         "Resources",
			ShortDescription: `Resource for NGFW manipulation.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudngfwaws_ngfw_log_profile",
			Category:         "Resources",
			ShortDescription: `Resource for NGFW log profile manipulation.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudngfwaws_predefined_url_category_override",
			Category:         "Resources",
			ShortDescription: `Resource for predefined URL category override management.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudngfwaws_prefix_list",
			Category:         "Resources",
			ShortDescription: `Resource for prefix list manipulation.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudngfwaws_rulestack",
			Category:         "Resources",
			ShortDescription: `Resource for rulestack manipulation.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudngfwaws_security_rule",
			Category:         "Resources",
			ShortDescription: `Resource for security rule manipulation.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"cloudngfwaws_certificate":                      0,
		"cloudngfwaws_commit_rulestack":                 1,
		"cloudngfwaws_custom_url_category":              2,
		"cloudngfwaws_fqdn_list":                        3,
		"cloudngfwaws_intelligent_feed":                 4,
		"cloudngfwaws_ngfw":                             5,
		"cloudngfwaws_ngfw_log_profile":                 6,
		"cloudngfwaws_predefined_url_category_override": 7,
		"cloudngfwaws_prefix_list":                      8,
		"cloudngfwaws_rulestack":                        9,
		"cloudngfwaws_security_rule":                    10,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
