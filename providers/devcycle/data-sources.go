package devcycle

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "devcycle_environment",
			Category:         "Data Sources",
			ShortDescription: `DevCycle Environment Data Source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "devcycle_evaluated_variable_boolean",
			Category:         "Data Sources",
			ShortDescription: `Evaluated Variable data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "devcycle_evaluated_variable_json",
			Category:         "Data Sources",
			ShortDescription: `Evaluated Variable data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "devcycle_evaluated_variable_number",
			Category:         "Data Sources",
			ShortDescription: `Evaluated Variable data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "devcycle_evaluated_variable_string",
			Category:         "Data Sources",
			ShortDescription: `Evaluated Variable data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "devcycle_feature",
			Category:         "Data Sources",
			ShortDescription: `DevCycle Feature data source`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "devcycle_project",
			Category:         "Data Sources",
			ShortDescription: `DevCycle Project data source`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "devcycle_variable",
			Category:         "Data Sources",
			ShortDescription: `DevCycle Variable data source`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"devcycle_environment":                0,
		"devcycle_evaluated_variable_boolean": 1,
		"devcycle_evaluated_variable_json":    2,
		"devcycle_evaluated_variable_number":  3,
		"devcycle_evaluated_variable_string":  4,
		"devcycle_feature":                    5,
		"devcycle_project":                    6,
		"devcycle_variable":                   7,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
