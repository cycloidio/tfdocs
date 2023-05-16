package powerflex

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "powerflex_protection_domain",
			Category:         "Data Sources",
			ShortDescription: `This datasource can be used to fetch information related to protection domains from a PowerFlex array.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "powerflex_sdc",
			Category:         "Data Sources",
			ShortDescription: `This data-source can be used to fetch information related to Storage Data Clients from a PowerFlex array.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "powerflex_sds",
			Category:         "Data Sources",
			ShortDescription: `This data-source can be used to fetch information related to Storage Data Servers from a PowerFlex array.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "powerflex_snapshot_policy",
			Category:         "Data Sources",
			ShortDescription: `This data-source can be used to fetch information related to the snapshot policies from a PowerFlex array.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "powerflex_storage_pool",
			Category:         "Data Sources",
			ShortDescription: `This data-source can be used to fetch information related to the storage pools from a PowerFlex array.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "powerflex_volume",
			Category:         "Data Sources",
			ShortDescription: `This data-source can be used to fetch information related to volumes from a PowerFlex array.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"powerflex_protection_domain": 0,
		"powerflex_sdc":               1,
		"powerflex_sds":               2,
		"powerflex_snapshot_policy":   3,
		"powerflex_storage_pool":      4,
		"powerflex_volume":            5,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
