package nomad

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "nomad_acl_policy",
			Category:         "Data Sources",
			ShortDescription: `Retrieve information on an ACL Policy.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nomad_acl_token",
			Category:         "Data Sources",
			ShortDescription: `Get information on an ACL token.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
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
			Type:             "nomad_plugin",
			Category:         "Data Sources",
			ShortDescription: `Get information on a specific CSI plugin.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nomad_plugins",
			Category:         "Data Sources",
			ShortDescription: `Retrieve a list of plugins.`,
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
		&resource.Resource{
			Name:             "",
			Type:             "nomad_volumes",
			Category:         "Data Sources",
			ShortDescription: `Retrieve a list of volumes.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"nomad_acl_policy":  0,
		"nomad_acl_token":   1,
		"nomad_deployments": 2,
		"nomad_job":         3,
		"nomad_namespaces":  4,
		"nomad_plugin":      5,
		"nomad_plugins":     6,
		"nomad_regions":     7,
		"nomad_volumes":     8,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
