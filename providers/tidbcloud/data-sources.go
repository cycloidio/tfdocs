package tidbcloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "tidbcloud_backups",
			Category:         "Data Sources",
			ShortDescription: `backups data source`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tidbcloud_cluster_specs",
			Category:         "Data Sources",
			ShortDescription: `cluster_specs data source`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tidbcloud_clusters",
			Category:         "Data Sources",
			ShortDescription: `clusters data source`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tidbcloud_projects",
			Category:         "Data Sources",
			ShortDescription: `projects data source`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tidbcloud_restores",
			Category:         "Data Sources",
			ShortDescription: `restores data source`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"tidbcloud_backups":       0,
		"tidbcloud_cluster_specs": 1,
		"tidbcloud_clusters":      2,
		"tidbcloud_projects":      3,
		"tidbcloud_restores":      4,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
