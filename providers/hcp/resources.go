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
			ShortDescription: `The AWS Network peering resource allows you to manage a Network peering between an HVN and a peer AWS VPC.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcp_aws_transit_gateway_attachment",
			Category:         "Resources",
			ShortDescription: `The AWS Transit Gateway Attachment resource allows you to manage a transit gateway attachment. The transit gateway attachment attaches an HVN to a user-owned transit gateway in AWS. Note that the HVN and transit gateway must be located in the same AWS region.`,
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
			ShortDescription: `The cluster root token resource is the token used to bootstrap the cluster's ACL system. Using this resource to create a new root token for a cluster will invalidate the consul root token accessor id and Consul root token secret id properties of the cluster.`,
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
	}

	resourcesMap = map[string]int{

		"hcp_aws_network_peering":            0,
		"hcp_aws_transit_gateway_attachment": 1,
		"hcp_consul_cluster":                 2,
		"hcp_consul_cluster_root_token":      3,
		"hcp_consul_snapshot":                4,
		"hcp_hvn":                            5,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
