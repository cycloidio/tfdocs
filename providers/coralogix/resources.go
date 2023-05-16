package coralogix

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "coralogix_action",
			Category:         "Resources",
			ShortDescription: `"Coralogix action. For more info please review - https://coralogix.com/docs/coralogix-action-extension/."`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "coralogix_alert",
			Category:         "Resources",
			ShortDescription: `"Coralogix alert. More info: https://coralogix.com/docs/alerts-api/."`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "coralogix_dashboard",
			Category:         "Resources",
			ShortDescription: `"Coralogix Dashboard."`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "coralogix_data_set",
			Category:         "Resources",
			ShortDescription: `"Coralogix data-set. More info: https://coralogix.com/blog/elevate-your-event-data-with-custom-data-enrichment-in-coralogix/"`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "coralogix_enrichment",
			Category:         "Resources",
			ShortDescription: `"Coralogix enrichment. More info: https://coralogix.com/docs/threat-discovery/."`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "coralogix_events2metric",
			Category:         "Resources",
			ShortDescription: `"Coralogix events2metric. More info:-https://coralogix.com/docs/event2metrics/"`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "coralogix_hosted_dashboard",
			Category:         "Resources",
			ShortDescription: `"Hosted dashboard. Can be one of - ["grafana"]."`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "coralogix_recording_rules_group",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "coralogix_rules_group",
			Category:         "Resources",
			ShortDescription: `"Rule-group is list of rule-subgroups with 'and' (&&) operation between. More info: https://coralogix.com/docs/rules-api/."`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "coralogix_tco_policy",
			Category:         "Resources",
			ShortDescription: `"Coralogix TCO-Policy. For more information - https://coralogix.com/docs/tco-optimizer-api ."`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "coralogix_webhook",
			Category:         "Resources",
			ShortDescription: `"Webhook defines integration. More info - https://coralogix.com/integrations/ (Alerting section)."`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"coralogix_action":                0,
		"coralogix_alert":                 1,
		"coralogix_dashboard":             2,
		"coralogix_data_set":              3,
		"coralogix_enrichment":            4,
		"coralogix_events2metric":         5,
		"coralogix_hosted_dashboard":      6,
		"coralogix_recording_rules_group": 7,
		"coralogix_rules_group":           8,
		"coralogix_tco_policy":            9,
		"coralogix_webhook":               10,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
