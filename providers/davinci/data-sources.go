package davinci

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "davinci_application",
			Category:         "Application",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"application",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "davinci_applications",
			Category:         "Connection",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"connection",
				"applications",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "davinci_connection",
			Category:         "Connection",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"connection",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "davinci_connections",
			Category:         "Connection",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"connection",
				"connections",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"davinci_application":  0,
		"davinci_applications": 1,
		"davinci_connection":   2,
		"davinci_connections":  3,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
