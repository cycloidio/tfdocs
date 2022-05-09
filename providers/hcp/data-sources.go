package hcp

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "hcp_aws_network_peering",
			Category:         "Data Sources",
			ShortDescription: `The AWS network peering data source provides information about an existing network peering between an HVN and a peer AWS VPC.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcp_aws_transit_gateway_attachment",
			Category:         "Data Sources",
			ShortDescription: `The AWS transit gateway attachment data source provides information about an existing transit gateway attachment.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcp_azure_peering_connection",
			Category:         "Data Sources",
			ShortDescription: `The Azure peering connection data source provides information about a peering connection between an HVN and a peer Azure VNet.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcp_consul_agent_helm_config",
			Category:         "Data Sources",
			ShortDescription: `The Consul agent Helm config data source provides Helm values for a Consul agent running in Kubernetes.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcp_consul_agent_kubernetes_secret",
			Category:         "Data Sources",
			ShortDescription: `The agent config Kubernetes secret data source provides Consul agents running in Kubernetes the configuration needed to connect to the Consul cluster.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcp_consul_cluster",
			Category:         "Data Sources",
			ShortDescription: `The cluster data source provides information about an existing HCP Consul cluster.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcp_consul_versions",
			Category:         "Data Sources",
			ShortDescription: `The Consul versions data source provides the Consul versions supported by HCP.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcp_hvn",
			Category:         "Data Sources",
			ShortDescription: `The HVN data source provides information about an existing HashiCorp Virtual Network.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcp_hvn_peering_connection",
			Category:         "Data Sources",
			ShortDescription: `The HVN peering connection data source provides information about an existing peering connection between HVNs.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcp_hvn_route",
			Category:         "Data Sources",
			ShortDescription: `The HVN route data source provides information about an existing HVN route.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcp_packer_image",
			Category:         "Data Sources",
			ShortDescription: `The Packer Image data source iteration gets the most recent iteration (or build) of an image, given an iteration id.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcp_packer_image_iteration",
			Category:         "Data Sources",
			ShortDescription: `The Packer Image data source iteration gets the most recent iteration (or build) of an image, given a channel.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcp_packer_iteration",
			Category:         "Data Sources",
			ShortDescription: `The Packer Image data source iteration gets the most recent iteration (or build) of an image, given a channel.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcp_vault_cluster",
			Category:         "Data Sources",
			ShortDescription: `The cluster data source provides information about an existing HCP Vault cluster.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"hcp_aws_network_peering":            0,
		"hcp_aws_transit_gateway_attachment": 1,
		"hcp_azure_peering_connection":       2,
		"hcp_consul_agent_helm_config":       3,
		"hcp_consul_agent_kubernetes_secret": 4,
		"hcp_consul_cluster":                 5,
		"hcp_consul_versions":                6,
		"hcp_hvn":                            7,
		"hcp_hvn_peering_connection":         8,
		"hcp_hvn_route":                      9,
		"hcp_packer_image":                   10,
		"hcp_packer_image_iteration":         11,
		"hcp_packer_iteration":               12,
		"hcp_vault_cluster":                  13,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
