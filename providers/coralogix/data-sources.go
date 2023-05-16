package coralogix

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "coralogix_action",
			Category:         "Data Sources",
			ShortDescription: `"Coralogix action. For more info please review - https://coralogix.com/docs/coralogix-action-extension/."`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "coralogix_alert",
			Category:         "Data Sources",
			ShortDescription: `"Coralogix Alert"`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "coralogix_dashboard",
			Category:         "Data Sources",
			ShortDescription: `"Coralogix Dashboard."`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "coralogix_data_set",
			Category:         "Data Sources",
			ShortDescription: `"Coralogix data-set. More info: https://coralogix.com/blog/elevate-your-event-data-with-custom-data-enrichment-in-coralogix/."`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "coralogix_enrichment",
			Category:         "Data Sources",
			ShortDescription: `"Coralogix enrichment. More info: https://coralogix.com/docs/threat-discovery/."`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "coralogix_events2metric",
			Category:         "Data Sources",
			ShortDescription: `"Coralogix events2metric. More info:-https://coralogix.com/docs/event2metrics/"`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "coralogix_hosted_dashboard",
			Category:         "Data Sources",
			ShortDescription: `"Hosted dashboard. Can be one of - ["grafana"]."`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "coralogix_recording_rules_group",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "coralogix_rules_group",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "coralogix_tco_policy",
			Category:         "Data Sources",
			ShortDescription: `"Coralogix TCO-Policy. For more information - https://coralogix.com/docs/tco-optimizer-api ."`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "coralogix_webhook",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

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

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
