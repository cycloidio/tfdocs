package upcloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "upcloud_hosts",
			Category:         "Data Sources",
			ShortDescription: `Provides an UpCloud hosts datasource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "hosts",
					Description: `A list of hosts within your private cloud that can be accessed by the account ### hosts Each entry in the ` + "`" + `hosts` + "`" + ` attribute represents an individual host from your account. The following attributes are available. Each ` + "`" + `hosts` + "`" + ` entry provides the following:`,
				},
				resource.Attribute{
					Name:        "host_id",
					Description: `The unique ID representing the host`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Freeform comment string for the host`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `The zone ID where the hosts is located.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_ip_addresses",
			Category:         "Data Sources",
			ShortDescription: `Provides an UpCloud IP Addresses datasource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_networks",
			Category:         "Data Sources",
			ShortDescription: `Get information on available UpCloud networks.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) Limit returned networks to this UpCloud zone.`,
				},
				resource.Attribute{
					Name:        "filter_name",
					Description: `(Optional) A regular-expression allowing filtering of networks by name. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "networks",
					Description: `A set of the available networks. ### networks`,
				},
				resource.Attribute{
					Name:        "ip_network",
					Description: `A set of the IP subnets within the network.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the network.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of network (one of ` + "`" + `public` + "`" + `, ` + "`" + `utility` + "`" + `, ` + "`" + `private` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier of the network.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `The zone in which this network resides.`,
				},
				resource.Attribute{
					Name:        "servers",
					Description: `A set of the servers joined to this network. ### ip_network`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The CIRD range of the subnet.`,
				},
				resource.Attribute{
					Name:        "dhcp",
					Description: `Is DHCP enabled?`,
				},
				resource.Attribute{
					Name:        "dhcp_default_route",
					Description: `Is the ` + "`" + `gateway` + "`" + ` the DHCP default route?`,
				},
				resource.Attribute{
					Name:        "dhcp_dns",
					Description: `A set of the DNS servers given by DHCP.`,
				},
				resource.Attribute{
					Name:        "family",
					Description: `The IP address familty (one of ` + "`" + `IPv4` + "`" + ` or ` + "`" + `IPv6` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `The gateway address given by DHCP. ### servers`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier of the server.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `The short description of the server. [1]: https://upcloud.com/products/software-defined-networking/`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "networks",
					Description: `A set of the available networks. ### networks`,
				},
				resource.Attribute{
					Name:        "ip_network",
					Description: `A set of the IP subnets within the network.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the network.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of network (one of ` + "`" + `public` + "`" + `, ` + "`" + `utility` + "`" + `, ` + "`" + `private` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier of the network.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `The zone in which this network resides.`,
				},
				resource.Attribute{
					Name:        "servers",
					Description: `A set of the servers joined to this network. ### ip_network`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The CIRD range of the subnet.`,
				},
				resource.Attribute{
					Name:        "dhcp",
					Description: `Is DHCP enabled?`,
				},
				resource.Attribute{
					Name:        "dhcp_default_route",
					Description: `Is the ` + "`" + `gateway` + "`" + ` the DHCP default route?`,
				},
				resource.Attribute{
					Name:        "dhcp_dns",
					Description: `A set of the DNS servers given by DHCP.`,
				},
				resource.Attribute{
					Name:        "family",
					Description: `The IP address familty (one of ` + "`" + `IPv4` + "`" + ` or ` + "`" + `IPv6` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `The gateway address given by DHCP. ### servers`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier of the server.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `The short description of the server. [1]: https://upcloud.com/products/software-defined-networking/`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_tags",
			Category:         "Data Sources",
			ShortDescription: `Provides an UpCloud tags datasource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "tags",
					Description: `A set of tags Each element of the tags set provides the following attributes`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The value representing the tag`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Free form text representing the meaning of the tag`,
				},
				resource.Attribute{
					Name:        "servers",
					Description: `A collection of servers that have been assigned the tag`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_zone",
			Category:         "Data Sources",
			ShortDescription: `Provides an UpCloud zone datasource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The UpCloud value that represents the zone. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique label for the zone (same as ` + "`" + `name` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A real world value representing the zone, for example ` + "`" + `London #1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `Identifies the zone as either public ` + "`" + `true` + "`" + ` or private ` + "`" + `false` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Unique label for the zone (same as ` + "`" + `name` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A real world value representing the zone, for example ` + "`" + `London #1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `Identifies the zone as either public ` + "`" + `true` + "`" + ` or private ` + "`" + `false` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_zones",
			Category:         "Data Sources",
			ShortDescription: `Provides an UpCloud zones datasource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter_type",
					Description: `(Optional) Allows the zone_ids to be filtered by type, (public, private or all). Defaults to ` + "`" + `all` + "`" + ` ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `UTC timestamp when the request was completed; used only for identifying the request in Terraform.`,
				},
				resource.Attribute{
					Name:        "zone_ids",
					Description: `A Collection of ids representing the available zones within UpCloud.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `UTC timestamp when the request was completed; used only for identifying the request in Terraform.`,
				},
				resource.Attribute{
					Name:        "zone_ids",
					Description: `A Collection of ids representing the available zones within UpCloud.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"upcloud_hosts":        0,
		"upcloud_ip_addresses": 1,
		"upcloud_networks":     2,
		"upcloud_tags":         3,
		"upcloud_zone":         4,
		"upcloud_zones":        5,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
