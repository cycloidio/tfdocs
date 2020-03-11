package akamai

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "akamai_authorities_set",
			Category:         "Data Sources",
			ShortDescription: `DNS Authorities Set`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_contract",
			Category:         "Data Sources",
			ShortDescription: `Contract`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_cp_code",
			Category:         "Data Sources",
			ShortDescription: `CP Code`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_group",
			Category:         "Data Sources",
			ShortDescription: `Group`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_gtm_default_datacenter",
			Category:         "Data Sources",
			ShortDescription: `CP Code`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "akamai_property_rules",
			Category:         "Data Sources",
			ShortDescription: `Property Rules`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"akamai_authorities_set":        0,
		"akamai_contract":               1,
		"akamai_cp_code":                2,
		"akamai_group":                  3,
		"akamai_gtm_default_datacenter": 4,
		"akamai_property_rules":         5,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
