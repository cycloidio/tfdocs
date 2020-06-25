package sumologic

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "sumologic_caller_identity",
			Category:         "Data Sources",
			ShortDescription: `Provides an easy way to retrieve Sumo Logic auth details.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sumologic_collector",
			Category:         "Data Sources",
			ShortDescription: `Provides a way to retrieve Sumo Logic collector details (id, names, etc) for a collector managed by another terraform stack.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sumologic_personal_folder",
			Category:         "Data Sources",
			ShortDescription: `Provides an easy way to retrieve the Personal Folder.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"sumologic_caller_identity": 0,
		"sumologic_collector":       1,
		"sumologic_personal_folder": 2,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
