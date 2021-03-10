package limelight

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "limelight_ip_ranges",
			Category:         "Data Sources",
			ShortDescription: `An IP Ranges data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `A ` + "`" + `string` + "`" + ` list where each element is an IP address.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version for the current list of ` + "`" + `ip_ranges` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `A ` + "`" + `string` + "`" + ` list where each element is an IP address.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version for the current list of ` + "`" + `ip_ranges` + "`" + `.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"limelight_ip_ranges": 0,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
