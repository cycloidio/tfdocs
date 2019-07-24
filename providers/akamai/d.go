package aws

import (
	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*Resource{

		&resource.Resource{
			Name:             "",
			Type:             "akamai_authorities_set",
			Category:         "Data Sources",
			ShortDescription: `DNS Authorities Set`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Argument{},
			Attributes:       []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_contract",
			Category:         "Data Sources",
			ShortDescription: `Contract`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Argument{},
			Attributes:       []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_cpcode",
			Category:         "Data Sources",
			ShortDescription: `CPCode`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Argument{},
			Attributes:       []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_group",
			Category:         "Data Sources",
			ShortDescription: `Group`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Argument{},
			Attributes:       []resource.Argument{},
		},
	}

	dataSourcesMap = map[string]Resource{

		"akamai_authorities_set": 0,
		"akamai_contract":        1,
		"akamai_cpcode":          2,
		"akamai_group":           3,
	}
)

func GetDataSource(r string) (*resouce.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs]
}
