package nomad

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "nomad_deployments",
			Category:         "Data Sources",
			ShortDescription: `Retrieve a list of deployments and a summary of their attributes.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nomad_job",
			Category:         "Data Sources",
			ShortDescription: `Get information on an job.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nomad_namespaces",
			Category:         "Data Sources",
			ShortDescription: `Retrieve a list of namespaces available in Nomad.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nomad_regions",
			Category:         "Data Sources",
			ShortDescription: `Retrieve a list of regions available in Nomad.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"nomad_deployments": 0,
		"nomad_job":         1,
		"nomad_namespaces":  2,
		"nomad_regions":     3,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
