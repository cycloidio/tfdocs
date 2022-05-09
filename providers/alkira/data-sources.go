package alkira

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "alkira_billing_tag",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing billing tag.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_credential",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing credential.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_group",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing group.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_group_connector",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get an existing connector group.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_group_user",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get an existing user group.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_policy_prefix_list",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get an existing policy prefix list.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_segment",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing segment.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"alkira_billing_tag":        0,
		"alkira_credential":         1,
		"alkira_group":              2,
		"alkira_group_connector":    3,
		"alkira_group_user":         4,
		"alkira_policy_prefix_list": 5,
		"alkira_segment":            6,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
