package hcloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

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

		"hcloud_floating_ip":            0,
		"hcloud_floating_ip_assignment": 1,
		"hcloud_network":                2,
		"hcloud_network_route":          3,
		"hcloud_network_subnet":         4,
		"hcloud_rdns":                   5,
		"hcloud_server":                 6,
		"hcloud_server_network":         7,
		"hcloud_ssh_key":                8,
		"hcloud_volume":                 9,
		"hcloud_volume_attachment":      10,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
