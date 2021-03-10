package hcs

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "hcs_agent_helm_config",
			Category:         "Data Sources",
			ShortDescription: `The agent Helm config data source provides Helm values for a Consul agent running in Kubernetes.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcs_agent_kubernetes_secret",
			Category:         "Data Sources",
			ShortDescription: `The agent config Kubernetes secret data source provides Consul agents running in Kubernetes the configuration needed to connect to the Consul cluster.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcs_cluster",
			Category:         "Data Sources",
			ShortDescription: `The cluster data source provides information about an existing HCS cluster.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcs_consul_versions",
			Category:         "Data Sources",
			ShortDescription: `The Consul versions data source provides the Consul versions supported by HCS.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcs_federation_token",
			Category:         "Data Sources",
			ShortDescription: `The federation token data source can be used during HCS cluster creation to join the cluster to a federation.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcs_plan_defaults",
			Category:         "Data Sources",
			ShortDescription: `The plan defaults data source provides details about the current Azure Marketplace Plan defaults for the HCS offering. The plan defaults are useful when accepting the Azure Marketplace Agreement for the HCS Azure Managed Application.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"hcs_agent_helm_config":       0,
		"hcs_agent_kubernetes_secret": 1,
		"hcs_cluster":                 2,
		"hcs_consul_versions":         3,
		"hcs_federation_token":        4,
		"hcs_plan_defaults":           5,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
