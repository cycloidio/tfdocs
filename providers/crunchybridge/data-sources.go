package crunchybridge

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "crunchybridge_account",
			Category:         "Data Sources",
			ShortDescription: `Data Source for retreiving Account team resource data`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "crunchybridge_cloudprovider",
			Category:         "Data Sources",
			ShortDescription: `Data Source for retreiving Cluster resource data`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "crunchybridge_cluster",
			Category:         "Data Sources",
			ShortDescription: `Data Source for retreiving Cluster resource data`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "crunchybridge_clusterids",
			Category:         "Data Sources",
			ShortDescription: `Data Source for retreiving Cluster identifiers from the user-provided label`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "crunchybridge_clusterroles",
			Category:         "Data Sources",
			ShortDescription: `Data Source for retreiving Cluster resource data`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "crunchybridge_clusterstatus",
			Category:         "Data Sources",
			ShortDescription: `Data Source for retreiving Cluster resource data`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"crunchybridge_account":       0,
		"crunchybridge_cloudprovider": 1,
		"crunchybridge_cluster":       2,
		"crunchybridge_clusterids":    3,
		"crunchybridge_clusterroles":  4,
		"crunchybridge_clusterstatus": 5,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
