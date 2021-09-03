package hcloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "hcloud_certificate",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Hetzner Cloud Certificate.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_certificates",
			Category:         "Data Sources",
			ShortDescription: `Provides details about multiple Hetzner Cloud Certificates.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_datacenter",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Hetzner Cloud Datacenter.`,
			Description: `
Provides details about a specific Hetzner Cloud Datacenter.
Use this resource to get detailed information about specific datacenter.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_datacenters",
			Category:         "Data Sources",
			ShortDescription: `List all available Hetzner Cloud Datacenters.`,
			Description: `
Provides a list of available Hetzner Cloud Datacenters.
This resource may be useful to create highly available infrastructure, distributed across several datacenters.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_firewall",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Hetzner Cloud Firewall.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_firewalls",
			Category:         "Data Sources",
			ShortDescription: `Provides details about multiple Hetzner Cloud Firewall.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_floating_ip",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Hetzner Cloud Floating IP.`,
			Description: `

Provides details about a Hetzner Cloud Floating IP.

This resource can be useful when you need to determine a Floating IP ID based on the IP address.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_floating_ips",
			Category:         "Data Sources",
			ShortDescription: `Provides details about multiple Hetzner Cloud Floating IPs.`,
			Description: `
Provides details about multiple Hetzner Cloud Floating IPs.


`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_image",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Hetzner Cloud Image.`,
			Description: `
Provides details about a Hetzner Cloud Image.
This resource is useful if you want to use a non-terraform managed image.
`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_images",
			Category:         "Data Sources",
			ShortDescription: `Provides details about multiple Hetzner Cloud Images.`,
			Description: `
Provides details about multiple Hetzner Cloud Images.


`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_load_balancer",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Hetzner Cloud Load Balancer.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_load_balancers",
			Category:         "Data Sources",
			ShortDescription: `Provides details about multiple Hetzner Cloud Load Balancers.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_location",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Hetzner Cloud Location.`,
			Description: `
Provides details about a specific Hetzner Cloud Location.
Use this resource to get detailed information about specific location.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_locations",
			Category:         "Data Sources",
			ShortDescription: `List all available Hetzner Cloud Locations.`,
			Description: `
Provides a list of available Hetzner Cloud Locations.
This resource may be useful to create highly available infrastructure, distributed across several locations.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_network",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Hetzner Cloud network.`,
			Description: `
Provides details about a Hetzner Cloud network.
This resource is useful if you want to use a non-terraform managed network.
`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_networks",
			Category:         "Data Sources",
			ShortDescription: `Provides details about multiple Hetzner Cloud Networks.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_placement_group",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Hetzner Cloud Placement Group.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_placement_groups",
			Category:         "Data Sources",
			ShortDescription: `Provides details about multiple Hetzner Cloud Placement Groups.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_server",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Hetzner Cloud Server.`,
			Description: `
Provides details about a Hetzner Cloud Server.
This resource is useful if you want to use a non-terraform managed server.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_server_type",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Hetzner Cloud Server Type.`,
			Description: `
Provides details about a specific Hetzner Cloud Server Type.
Use this resource to get detailed information about specific Server Type.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_server_types",
			Category:         "Data Sources",
			ShortDescription: `List all available Hetzner Cloud Server Types.`,
			Description: `
Provides a list of available Hetzner Cloud Server Types.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_servers",
			Category:         "Data Sources",
			ShortDescription: `Provides details about multiple Hetzner Cloud Servers.`,
			Description: `
Provides details about multiple Hetzner Cloud Servers.
This resource is useful if you want to use non-terraform managed servers.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_ssh_key",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Hetzner Cloud SSH Key.`,
			Description: `
Provides details about a Hetzner Cloud SSH Key.
This resource is useful if you want to use a non-terraform managed SSH Key.
`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_ssh_keys",
			Category:         "Data Sources",
			ShortDescription: `Provides details about multiple Hetzner Cloud SSH Keys.`,
			Description: `

Provides details about Hetzner Cloud SSH Keys.
This resource is useful if you want to use a non-terraform managed SSH Key.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_volume",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Hetzner Cloud volume.`,
			Description: `
Provides details about a Hetzner Cloud volume.
This resource is useful if you want to use a non-terraform managed volume.
`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_volumes",
			Category:         "Data Sources",
			ShortDescription: `Provides details about multiple Hetzner Cloud volumes.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"hcloud_certificate":      0,
		"hcloud_certificates":     1,
		"hcloud_datacenter":       2,
		"hcloud_datacenters":      3,
		"hcloud_firewall":         4,
		"hcloud_firewalls":        5,
		"hcloud_floating_ip":      6,
		"hcloud_floating_ips":     7,
		"hcloud_image":            8,
		"hcloud_images":           9,
		"hcloud_load_balancer":    10,
		"hcloud_load_balancers":   11,
		"hcloud_location":         12,
		"hcloud_locations":        13,
		"hcloud_network":          14,
		"hcloud_networks":         15,
		"hcloud_placement_group":  16,
		"hcloud_placement_groups": 17,
		"hcloud_server":           18,
		"hcloud_server_type":      19,
		"hcloud_server_types":     20,
		"hcloud_servers":          21,
		"hcloud_ssh_key":          22,
		"hcloud_ssh_keys":         23,
		"hcloud_volume":           24,
		"hcloud_volumes":          25,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
