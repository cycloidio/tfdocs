package anxcloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "anxcloud_dns_record",
			Category:         "Resources",
			ShortDescription: `This resource allows you to create DNS records for a specified zone. TXT records might behave funny, we are working on it.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "anxcloud_dns_zone",
			Category:         "Resources",
			ShortDescription: `This resource allows you to create DNS zones.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "anxcloud_ip_address",
			Category:         "Resources",
			ShortDescription: `This resource allows you to create and configure IP addresses.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "anxcloud_kubernetes_cluster",
			Category:         "Resources",
			ShortDescription: `Resource to create Kubernetes clusters. Updates are currently not supported.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "anxcloud_kubernetes_kubeconfig",
			Category:         "Resources",
			ShortDescription: `Resource to create a Kubernetes kubeconfig.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "anxcloud_kubernetes_node_pool",
			Category:         "Resources",
			ShortDescription: `Resource to create Kubernetes node pools. Updates are currently not supported.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "anxcloud_lbaas_loadbalancer",
			Category:         "Resources",
			ShortDescription: `This resource allows you to create and manage LBaaS Load Balancer resources.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "anxcloud_network_prefix",
			Category:         "Resources",
			ShortDescription: `This resource allows you to create and configure network prefix.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "anxcloud_tag",
			Category:         "Resources",
			ShortDescription: `The tag resource allows you to create and configure a tag.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "anxcloud_virtual_server",
			Category:         "Resources",
			ShortDescription: `The virtual_server resource allows you to configure and run virtual machines.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "anxcloud_vlan",
			Category:         "Resources",
			ShortDescription: `The VLAN resource allows you to create and configure VLAN.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"anxcloud_dns_record":            0,
		"anxcloud_dns_zone":              1,
		"anxcloud_ip_address":            2,
		"anxcloud_kubernetes_cluster":    3,
		"anxcloud_kubernetes_kubeconfig": 4,
		"anxcloud_kubernetes_node_pool":  5,
		"anxcloud_lbaas_loadbalancer":    6,
		"anxcloud_network_prefix":        7,
		"anxcloud_tag":                   8,
		"anxcloud_virtual_server":        9,
		"anxcloud_vlan":                  10,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
