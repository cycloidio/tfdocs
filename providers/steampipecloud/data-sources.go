package steampipecloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "steampipecloud_organization",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to retrieve information about an existing organization.`,
			Description: `

Use this data source to retrieve information about an existing organization.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "steampipecloud_process",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to retrieve information about a process for an identity or identity workspace.`,
			Description: `

Use this data source to retrieve information about a process that belongs to an identity or identity workspace.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "steampipecloud_user",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to retrieve information about the user whose token is used for authentication.`,
			Description: `

Use this data source to retrieve information about the user whose token is used for authentication.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"steampipecloud_organization": 0,
		"steampipecloud_process":      1,
		"steampipecloud_user":         2,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
