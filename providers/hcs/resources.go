package hcs

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "hcs_cluster",
			Category:         "Resources",
			ShortDescription: `The cluster resource allows you to manage an HCS Azure Managed Application.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcs_cluster_root_token",
			Category:         "Resources",
			ShortDescription: `The cluster root token resource is the token used to bootstrap the cluster's ACL system. Using this resource to create a new root token for an cluster resource will invalidate the consul root token accessor id and Consul root token secret id properties of the cluster.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcs_snapshot",
			Category:         "Resources",
			ShortDescription: `The snapshot resource allows users to manage Consul snapshots of an HCS cluster. Snapshots currently have a retention policy of 30 days.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"hcs_cluster":            0,
		"hcs_cluster_root_token": 1,
		"hcs_snapshot":           2,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
