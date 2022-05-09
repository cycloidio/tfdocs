package civo

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "civo_dns_domain_name",
			Category:         "Resources",
			ShortDescription: `Provides a Civo DNS domain name resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_dns_domain_record",
			Category:         "Resources",
			ShortDescription: `Provides a Civo DNS domain record resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_firewall",
			Category:         "Resources",
			ShortDescription: `Provides a Civo firewall resource. This can be used to create, modify, and delete firewalls.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_firewall_rule",
			Category:         "Resources",
			ShortDescription: `Provides a Civo firewall rule resource. This can be used to create, modify, and delete firewalls rules. This resource don't have an update option because Civo backend doesn't support it at this moment. In that case, we use ForceNew for all object in the resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_instance",
			Category:         "Resources",
			ShortDescription: `Provides a Civo instance resource. This can be used to create, modify, and delete instances.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_kubernetes_cluster",
			Category:         "Resources",
			ShortDescription: `Provides a Civo Kubernetes cluster resource. This can be used to create, delete, and modify clusters.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_kubernetes_node_pool",
			Category:         "Resources",
			ShortDescription: `Provides a Civo Kubernetes node pool resource. While the default node pool must be defined in the civo_kubernetes_cluster resource, this resource can be used to add additional ones to a cluster.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_network",
			Category:         "Resources",
			ShortDescription: `Provides a Civo network resource. This can be used to create, modify, and delete networks.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_ssh_key",
			Category:         "Resources",
			ShortDescription: `Provides a Civo SSH key resource to allow you to manage SSH keys for instance access. Keys created with this resource can be referenced in your instance configuration via their ID.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_volume",
			Category:         "Resources",
			ShortDescription: `Provides a Civo volume which can be attached to an instance in order to provide expanded storage.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_volume_attachment",
			Category:         "Resources",
			ShortDescription: `Manages volume attachment/detachment to an instance.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"civo_dns_domain_name":      0,
		"civo_dns_domain_record":    1,
		"civo_firewall":             2,
		"civo_firewall_rule":        3,
		"civo_instance":             4,
		"civo_kubernetes_cluster":   5,
		"civo_kubernetes_node_pool": 6,
		"civo_network":              7,
		"civo_ssh_key":              8,
		"civo_volume":               9,
		"civo_volume_attachment":    10,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
