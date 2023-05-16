package samsungcloudplatform

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "samsungcloudplatform_block_storage",
			Category:         "Resources",
			ShortDescription: `Provides a Block Storage resource.`,
			Description: `

Provides a Block Storage resource.


`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "samsungcloudplatform_file_storage",
			Category:         "Resources",
			ShortDescription: `Provides a File Storage resource.`,
			Description: `

Provides a File Storage resource.


`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "samsungcloudplatform_firewall",
			Category:         "Resources",
			ShortDescription: `Provides a Firewall resource.`,
			Description: `

Provides a Firewall resource.


`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "samsungcloudplatform_firewall_rule",
			Category:         "Resources",
			ShortDescription: `Provides a Firewall Rule resource.`,
			Description: `

Provides a Firewall Rule resource.


`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "samsungcloudplatform_internet_gateway",
			Category:         "Resources",
			ShortDescription: `Provides a Internet Gateway resource.`,
			Description: `

Provides a Internet Gateway resource.


`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "samsungcloudplatform_kubernetes_apps",
			Category:         "Resources",
			ShortDescription: `Provides a K8s Apps resource.`,
			Description: `

Provides a K8s Apps resource.


`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "samsungcloudplatform_kubernetes_engine",
			Category:         "Resources",
			ShortDescription: `Provides a K8s Engine resource.`,
			Description: `

Provides a K8s Engine resource.


`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "samsungcloudplatform_kubernetes_namespace",
			Category:         "Resources",
			ShortDescription: `Provides a K8s Namespace resource.`,
			Description: `

Provides a K8s Namespace resource.


`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "samsungcloudplatform_kubernetes_node_pool",
			Category:         "Resources",
			ShortDescription: `Provides a K8s Node Pool resource.`,
			Description: `

Provides a K8s Node Pool resource.


`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "samsungcloudplatform_lb_profile",
			Category:         "Resources",
			ShortDescription: `Provides a Load Balancer Profile resource.`,
			Description: `

Provides a Load Balancer Profile resource.


`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "samsungcloudplatform_lb_server_group",
			Category:         "Resources",
			ShortDescription: `Provides a Load Balancer Server Group resource.`,
			Description: `

Provides a Load Balancer Server Group resource.


`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "samsungcloudplatform_lb_service",
			Category:         "Resources",
			ShortDescription: `Provides a Load Balancer Service resource.`,
			Description: `

Provides a Load Balancer Service resource.


`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "samsungcloudplatform_load_balancer",
			Category:         "Resources",
			ShortDescription: `Provides a Load Balancer resource.`,
			Description: `

Provides a Load Balancer resource.


`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "samsungcloudplatform_nat_gateway",
			Category:         "Resources",
			ShortDescription: `Provides a NAT Gateway resource.`,
			Description: `

Provides a NAT Gateway resource.


`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "samsungcloudplatform_postgresql",
			Category:         "Resources",
			ShortDescription: `Provides a PostgreSQL Database resource.`,
			Description: `

Provides a PostgreSQL Database resource.


`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "samsungcloudplatform_public_ip",
			Category:         "Resources",
			ShortDescription: `Provides a Public IP resource.`,
			Description: `

Provides a Public IP resource.


`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "samsungcloudplatform_security_group",
			Category:         "Resources",
			ShortDescription: `Provides a Security Group resource.`,
			Description: `

Provides a Security Group resource.


`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "samsungcloudplatform_security_group_rule",
			Category:         "Resources",
			ShortDescription: `Provides a Security Group Rule resource.`,
			Description: `

Provides a Security Group Rule resource.


`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "samsungcloudplatform_subnet",
			Category:         "Resources",
			ShortDescription: `Provides a Subnet resource.`,
			Description: `

Provides a Subnet resource.


`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "samsungcloudplatform_virtual_server",
			Category:         "Resources",
			ShortDescription: `Provides a Virtual Server resource.`,
			Description: `

Provides a Virtual Server resource.


`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "samsungcloudplatform_vpc",
			Category:         "Resources",
			ShortDescription: `Provides a VPC resource.`,
			Description: `

Provides a VPC resource.


`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"samsungcloudplatform_block_storage":        0,
		"samsungcloudplatform_file_storage":         1,
		"samsungcloudplatform_firewall":             2,
		"samsungcloudplatform_firewall_rule":        3,
		"samsungcloudplatform_internet_gateway":     4,
		"samsungcloudplatform_kubernetes_apps":      5,
		"samsungcloudplatform_kubernetes_engine":    6,
		"samsungcloudplatform_kubernetes_namespace": 7,
		"samsungcloudplatform_kubernetes_node_pool": 8,
		"samsungcloudplatform_lb_profile":           9,
		"samsungcloudplatform_lb_server_group":      10,
		"samsungcloudplatform_lb_service":           11,
		"samsungcloudplatform_load_balancer":        12,
		"samsungcloudplatform_nat_gateway":          13,
		"samsungcloudplatform_postgresql":           14,
		"samsungcloudplatform_public_ip":            15,
		"samsungcloudplatform_security_group":       16,
		"samsungcloudplatform_security_group_rule":  17,
		"samsungcloudplatform_subnet":               18,
		"samsungcloudplatform_virtual_server":       19,
		"samsungcloudplatform_vpc":                  20,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
