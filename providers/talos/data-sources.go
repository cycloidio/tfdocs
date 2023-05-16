package talos

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "talos_client_configuration",
			Category:         "Data Sources",
			ShortDescription: `Generate client configuration for a Talos cluster`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "talos_cluster_kubeconfig",
			Category:         "Data Sources",
			ShortDescription: `Retrieves the kubeconfig for a Talos cluster`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "talos_machine_configuration",
			Category:         "Data Sources",
			ShortDescription: `Generate a machine configuration for a node type`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "talos_machine_disks",
			Category:         "Data Sources",
			ShortDescription: `Generate a machine configuration for a node type`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"talos_client_configuration":  0,
		"talos_cluster_kubeconfig":    1,
		"talos_machine_configuration": 2,
		"talos_machine_disks":         3,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
