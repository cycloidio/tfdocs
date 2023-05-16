package ec

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "ec_deployment",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to retrieve information about an existing Elastic Cloud deployment.`,
			Description: `

Use this data source to retrieve information about an existing Elastic Cloud deployment.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ec_deployments",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to retrieve a list of IDs for the deployment and resource kinds, based on the specified query.`,
			Description: `

Use this data source to retrieve a list of IDs for the deployment and resource kinds, based on the specified query.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ec_stack",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to retrieve information about an existing Elastic Cloud stack. -> **Note on regions** Before you start, you might want to check the [full list](https://www.elastic.co/guide/en/cloud/current/ec-regions-templates-instances.html) of regions available in Elasticsearch Service (ESS).`,
			Description: `

Use this data source to retrieve information about an existing Elastic Cloud stack.

  -> **Note on regions** Before you start, you might want to check the [full list](https://www.elastic.co/guide/en/cloud/current/ec-regions-templates-instances.html) of regions available in Elasticsearch Service (ESS).

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ec_traffic_filter",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to filter for an existing traffic filter that has been created via one of the provided filters.`,
			Description: `

Use this data source to filter for an existing traffic filter that has been created via one of the provided filters.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"ec_deployment":     0,
		"ec_deployments":    1,
		"ec_stack":          2,
		"ec_traffic_filter": 3,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
