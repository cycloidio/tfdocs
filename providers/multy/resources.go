package multy

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "multy_database",
			Category:         "Resources",
			ShortDescription: `Provides Multy Database resource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "multy_kubernetes_cluster",
			Category:         "Resources",
			ShortDescription: `Provides Multy Kubernetes Cluster resource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "multy_kubernetes_node_pool",
			Category:         "Resources",
			ShortDescription: `Provides Multy Object Storage Object resource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "multy_network_interface",
			Category:         "Resources",
			ShortDescription: `Provides Multy Network Interface resource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "multy_network_interface_security_group_association",
			Category:         "Resources",
			ShortDescription: `Provides Multy Network Interface Security Group Association resource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "multy_network_security_group",
			Category:         "Resources",
			ShortDescription: `Provides Multy Network Security Group resource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "multy_object_storage",
			Category:         "Resources",
			ShortDescription: `Provides Multy Object Storage resource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "multy_object_storage_object",
			Category:         "Resources",
			ShortDescription: `Provides Multy Object Storage Object resource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "multy_public_ip",
			Category:         "Resources",
			ShortDescription: `Provides Multy Public IP resource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "multy_route_table",
			Category:         "Resources",
			ShortDescription: `Provides Multy Route Table resource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "multy_route_table_association",
			Category:         "Resources",
			ShortDescription: `Provides Multy Route Table Association resource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "multy_subnet",
			Category:         "Resources",
			ShortDescription: `Provides Multy Subnet resource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "multy_vault",
			Category:         "Resources",
			ShortDescription: `Provides Multy Vault resource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "multy_vault_access_policy",
			Category:         "Resources",
			ShortDescription: `Provides Multy Object Storage resource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "multy_vault_secret",
			Category:         "Resources",
			ShortDescription: `Provides Multy Object Storage resource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "multy_virtual_machine",
			Category:         "Resources",
			ShortDescription: `Provides Multy Virtual Machine resource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "multy_virtual_network",
			Category:         "Resources",
			ShortDescription: `Provides Multy Virtual Network resource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"multy_database":                                     0,
		"multy_kubernetes_cluster":                           1,
		"multy_kubernetes_node_pool":                         2,
		"multy_network_interface":                            3,
		"multy_network_interface_security_group_association": 4,
		"multy_network_security_group":                       5,
		"multy_object_storage":                               6,
		"multy_object_storage_object":                        7,
		"multy_public_ip":                                    8,
		"multy_route_table":                                  9,
		"multy_route_table_association":                      10,
		"multy_subnet":                                       11,
		"multy_vault":                                        12,
		"multy_vault_access_policy":                          13,
		"multy_vault_secret":                                 14,
		"multy_virtual_machine":                              15,
		"multy_virtual_network":                              16,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
