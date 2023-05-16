package cpln

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "cpln_agent",
			Category:         "Agent",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"agent",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cpln_audit_context",
			Category:         "Audit Context",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"audit",
				"context",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cpln_cloud_account",
			Category:         "Cloud Account",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"account",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cpln_domain",
			Category:         "Domain",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"domain",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cpln_group",
			Category:         "Group",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"group",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cpln_gvc",
			Category:         "Global Virtual Cloud",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"global",
				"virtual",
				"cloud",
				"gvc",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cpln_identity",
			Category:         "Identity",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"identity",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cpln_org_logging",
			Category:         "Org Logging",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"org",
				"logging",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cpln_org_tracing",
			Category:         "Org Tracing",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"org",
				"tracing",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cpln_policy",
			Category:         "Policy",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cpln_secret",
			Category:         "Secret",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"secret",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cpln_service_account",
			Category:         "Service Account",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"service",
				"account",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cpln_service_account_key",
			Category:         "Service Account",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"service",
				"account",
				"key",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cpln_workload",
			Category:         "Workload",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"workload",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"cpln_agent":               0,
		"cpln_audit_context":       1,
		"cpln_cloud_account":       2,
		"cpln_domain":              3,
		"cpln_group":               4,
		"cpln_gvc":                 5,
		"cpln_identity":            6,
		"cpln_org_logging":         7,
		"cpln_org_tracing":         8,
		"cpln_policy":              9,
		"cpln_secret":              10,
		"cpln_service_account":     11,
		"cpln_service_account_key": 12,
		"cpln_workload":            13,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
