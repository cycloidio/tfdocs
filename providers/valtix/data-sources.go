package valtix

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "valtix_policy_rule_set",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the rule set ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "rule_set_id",
					Description: `Set to the RuleSetID of the found resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "rule_set_id",
					Description: `Set to the RuleSetID of the found resource`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"valtix_policy_rule_set": 0,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
