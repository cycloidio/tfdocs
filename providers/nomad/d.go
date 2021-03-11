package nomad

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "nomad_acl_policies",
			Category:         "Data Sources",
			ShortDescription: `Retrieve a list of ACL Policies.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
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
			Type:             "nomad_acl_tokens",
			Category:         "Data Sources",
			ShortDescription: `Get a list of ACL tokens.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nomad_datacenters",
			Category:         "Data Sources",
			ShortDescription: `Retrieve a list of datacenters.`,
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
			Type:             "nomad_job_parser",
			Category:         "Data Sources",
			ShortDescription: `Parse a HCL jobspec and produce the equivalent JSON encoded job.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nomad_namespace",
			Category:         "Data Sources",
			ShortDescription: `Get information about a namespace in Nomad.`,
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
			Type:             "scaling_policies",
			Category:         "Data Sources",
			ShortDescription: `Retrieve a list of Scaling Policies.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "scaling_policy",
			Category:         "Data Sources",
			ShortDescription: `Retrieve a Scaling Policy.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nomad_scheduler_config",
			Category:         "Data Sources",
			ShortDescription: `Retrieve the cluster's scheduler configuration.`,
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

		"nomad_acl_policies":     0,
		"nomad_acl_policy":       1,
		"nomad_acl_token":        2,
		"nomad_acl_tokens":       3,
		"nomad_datacenters":      4,
		"nomad_deployments":      5,
		"nomad_job":              6,
		"nomad_job_parser":       7,
		"nomad_namespace":        8,
		"nomad_namespaces":       9,
		"nomad_plugin":           10,
		"nomad_plugins":          11,
		"nomad_regions":          12,
		"scaling_policies":       13,
		"scaling_policy":         14,
		"nomad_scheduler_config": 15,
		"nomad_volumes":          16,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
