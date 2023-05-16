package external

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "external_external",
			Category:         "Data Sources",
			ShortDescription: `The external data source allows an external program implementing a specific protocol (defined below) to act as a data source, exposing arbitrary data for use elsewhere in the Terraform configuration. Warning This mechanism is provided as an "escape hatch" for exceptional situations where a first-class Terraform provider is not more appropriate. Its capabilities are limited in comparison to a true data source, and implementing a data source via an external program is likely to hurt the portability of your Terraform configuration by creating dependencies on external programs and libraries that may not be available (or may need to be used differently) on different operating systems. Warning Terraform Enterprise does not guarantee availability of any particular language runtimes or external programs beyond standard shell utilities, so it is not recommended to use this data source within configurations that are applied within Terraform Enterprise.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"external_external": 0,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
