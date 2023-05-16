package skysql

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "skysql_availability_zones",
			Category:         "Data Sources",
			ShortDescription: `Retrieve the list of availability_zones.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "skysql_credentials",
			Category:         "Data Sources",
			ShortDescription: `Returnes a default credentials for a SkySQL service`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "skysql_projects",
			Category:         "Data Sources",
			ShortDescription: `Retrieve the list of projects. Project is a way of grouping the services.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "skysql_service",
			Category:         "Data Sources",
			ShortDescription: `Returns an full SkySQL service details`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "skysql_versions",
			Category:         "Data Sources",
			ShortDescription: `SkySQL server versions`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"skysql_availability_zones": 0,
		"skysql_credentials":        1,
		"skysql_projects":           2,
		"skysql_service":            3,
		"skysql_versions":           4,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
