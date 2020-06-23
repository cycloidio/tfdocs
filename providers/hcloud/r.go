package hcloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "hcloud_certificate",
			Category:         "Resources",
			ShortDescription: `Provides a Hetzner Cloud Certificate to represent a TLS certificate in the Hetzner Cloud.`,
			Description:      ``,
			Keywords: []string{
				"certificate",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_floating_ip",
			Category:         "Resources",
			ShortDescription: `Provides a Hetzner Cloud Floating IP to represent a publicly-accessible static IP address that can be mapped to one of your servers.`,
			Description:      ``,
			Keywords: []string{
				"floating",
				"ip",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_floating_ip_assignment",
			Category:         "Resources",
			ShortDescription: `Provides a Hetzner Cloud Floating IP Assignment to assign a Floating IP to a Hetzner Cloud Server.`,
			Description:      ``,
			Keywords: []string{
				"floating",
				"ip",
				"assignment",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_load_balancer",
			Category:         "Resources",
			ShortDescription: `Provides a Hetzner Cloud Load Balancer to represent a Load Balancer in the Hetzner Cloud.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_load_balancer_network",
			Category:         "Resources",
			ShortDescription: `Provides a Hetzner Cloud Load Balancer Network to represent a private network on a Load Balancer in the Hetzner Cloud.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
				"network",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_load_balancer_service",
			Category:         "Resources",
			ShortDescription: `Define services for Hetzner Cloud Load Balancers.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
				"service",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_network",
			Category:         "Resources",
			ShortDescription: `Provides a Hetzner Cloud Network to represent a Network in the Hetzner Cloud.`,
			Description:      ``,
			Keywords: []string{
				"network",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_network_route",
			Category:         "Resources",
			ShortDescription: `Provides a Hetzner Cloud Network Route to represent a Network route in the Hetzner Cloud.`,
			Description:      ``,
			Keywords: []string{
				"network",
				"route",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_network_subnet",
			Category:         "Resources",
			ShortDescription: `Provides a Hetzner Cloud Network Subnet to represent a Subnet in the Hetzner Cloud.`,
			Description:      ``,
			Keywords: []string{
				"network",
				"subnet",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_rdns",
			Category:         "Resources",
			ShortDescription: `Provides a Hetzner Cloud Reverse DNS Entry to create, modify and reset reverse dns entries for Hetzner Cloud Floating IPs or servers.`,
			Description:      ``,
			Keywords: []string{
				"rdns",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_server",
			Category:         "Resources",
			ShortDescription: `Provides an Hetzner Cloud server resource. This can be used to create, modify, and delete servers. Servers also support provisioning.`,
			Description:      ``,
			Keywords: []string{
				"server",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_server_network",
			Category:         "Resources",
			ShortDescription: `Provides a Hetzner Cloud Server Network to represent a private network on a server in the Hetzner Cloud.`,
			Description:      ``,
			Keywords: []string{
				"server",
				"network",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_ssh_key",
			Category:         "Resources",
			ShortDescription: `Provides a Hetzner Cloud SSH key resource to manage SSH keys for server access.`,
			Description:      ``,
			Keywords: []string{
				"ssh",
				"key",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_volume",
			Category:         "Resources",
			ShortDescription: `Provides a Hetzner Cloud volume resource to manage volumes.`,
			Description:      ``,
			Keywords: []string{
				"volume",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_volume_attachment",
			Category:         "Resources",
			ShortDescription: `Provides a Hetzner Cloud Volume attachment to attach a Volume to a Hetzner Cloud Server.`,
			Description:      ``,
			Keywords: []string{
				"volume",
				"attachment",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"hcloud_certificate":            0,
		"hcloud_floating_ip":            1,
		"hcloud_floating_ip_assignment": 2,
		"hcloud_load_balancer":          3,
		"hcloud_load_balancer_network":  4,
		"hcloud_load_balancer_service":  5,
		"hcloud_network":                6,
		"hcloud_network_route":          7,
		"hcloud_network_subnet":         8,
		"hcloud_rdns":                   9,
		"hcloud_server":                 10,
		"hcloud_server_network":         11,
		"hcloud_ssh_key":                12,
		"hcloud_volume":                 13,
		"hcloud_volume_attachment":      14,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
