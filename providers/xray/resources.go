package xray

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "xray_ignore_rule",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "xray_license_policy",
			Category:         "Policies",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"policies",
				"license",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "xray_licenses_report",
			Category:         "Reports",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"reports",
				"licenses",
				"report",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "xray_operational_risk_policy",
			Category:         "Policies",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"policies",
				"operational",
				"risk",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "xray_operational_risks_report",
			Category:         "Reports",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"reports",
				"operational",
				"risks",
				"report",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "xray_repository_config",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "xray_security_policy",
			Category:         "Policies",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"policies",
				"security",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "xray_settings",
			Category:         "Settings",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"settings",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "xray_violations_report",
			Category:         "Reports",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"reports",
				"violations",
				"report",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "xray_vulnerabilities_report",
			Category:         "Reports",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"reports",
				"vulnerabilities",
				"report",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "xray_watch",
			Category:         "Watch",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"watch",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "xray_workers_count",
			Category:         "Workers Count",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"workers",
				"count",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"xray_ignore_rule":              0,
		"xray_license_policy":           1,
		"xray_licenses_report":          2,
		"xray_operational_risk_policy":  3,
		"xray_operational_risks_report": 4,
		"xray_repository_config":        5,
		"xray_security_policy":          6,
		"xray_settings":                 7,
		"xray_violations_report":        8,
		"xray_vulnerabilities_report":   9,
		"xray_watch":                    10,
		"xray_workers_count":            11,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
