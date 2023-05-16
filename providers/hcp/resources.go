package hcp

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "hcp_aws_network_peering",
			Category:         "Resources",
			ShortDescription: `The AWS network peering resource allows you to manage a network peering between an HVN and a peer AWS VPC.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcp_aws_transit_gateway_attachment",
			Category:         "Resources",
			ShortDescription: `The AWS transit gateway attachment resource allows you to manage a transit gateway attachment. The transit gateway attachment attaches an HVN to a user-owned transit gateway in AWS. Note that the HVN and transit gateway must be located in the same AWS region.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcp_azure_peering_connection",
			Category:         "Resources",
			ShortDescription: `The Azure peering connection resource allows you to manage a peering connection between an HVN and a peer Azure VNet.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcp_boundary_cluster",
			Category:         "Resources",
			ShortDescription: `This resource allows you to manage an HCP Boundary cluster`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcp_consul_cluster",
			Category:         "Resources",
			ShortDescription: `The Consul cluster resource allows you to manage an HCP Consul cluster.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcp_consul_cluster_root_token",
			Category:         "Resources",
			ShortDescription: `The cluster root token resource is the token used to bootstrap the cluster's ACL system. You can also generate this root token from the HCP Consul UI.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcp_consul_snapshot",
			Category:         "Resources",
			ShortDescription: `The Consul snapshot resource allows users to manage Consul snapshots of an HCP Consul cluster. Snapshots currently have a retention policy of 30 days.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcp_hvn",
			Category:         "Resources",
			ShortDescription: `The HVN resource allows you to manage a HashiCorp Virtual Network in HCP.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcp_hvn_peering_connection",
			Category:         "Resources",
			ShortDescription: `The HVN peering connection resource allows you to manage a peering connection between HVNs.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcp_hvn_route",
			Category:         "Resources",
			ShortDescription: `The HVN route resource allows you to manage an HVN route.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcp_packer_channel",
			Category:         "Resources",
			ShortDescription: `The Packer Channel resource allows you to manage image bucket channels within an active HCP Packer Registry.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcp_vault_cluster",
			Category:         "Resources",
			ShortDescription: `The Vault cluster resource allows you to manage an HCP Vault cluster.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcp_vault_cluster_admin_token",
			Category:         "Resources",
			ShortDescription: `The Vault cluster admin token resource generates an admin-level token for the HCP Vault cluster.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"hcp_aws_network_peering":            0,
		"hcp_aws_transit_gateway_attachment": 1,
		"hcp_azure_peering_connection":       2,
		"hcp_boundary_cluster":               3,
		"hcp_consul_cluster":                 4,
		"hcp_consul_cluster_root_token":      5,
		"hcp_consul_snapshot":                6,
		"hcp_hvn":                            7,
		"hcp_hvn_peering_connection":         8,
		"hcp_hvn_route":                      9,
		"hcp_packer_channel":                 10,
		"hcp_vault_cluster":                  11,
		"hcp_vault_cluster_admin_token":      12,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
