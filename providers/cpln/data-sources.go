package cpln

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "cpln_gvc",
			Category:         "Global Virtual Cloud",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"global",
				"virtual",
				"cloud",
				"gvc",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cpln_location",
			Category:         "Location",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"location",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cpln_locations",
			Category:         "Location",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"location",
				"locations",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cpln_org",
			Category:         "Org",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"org",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"cpln_gvc":       0,
		"cpln_location":  1,
		"cpln_locations": 2,
		"cpln_org":       3,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
