package triton

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "triton_account",
			Category:         "Data Sources",
			ShortDescription: `The ` + "`" + `triton_account` + "`" + ` data source queries Triton for Account information.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(string) The unique identifier representing the Account in Triton.`,
				},
				resource.Attribute{
					Name:        "login",
					Description: `(string) The login name associated with the Account.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(string) An e-mail address that is current set in the Account.`,
				},
				resource.Attribute{
					Name:        "cns_enabled",
					Description: `(boolean) Whether the Container Name Service (CNS) is enabled for the Account. [1]: /docs/providers/triton/index.html`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(string) The unique identifier representing the Account in Triton.`,
				},
				resource.Attribute{
					Name:        "login",
					Description: `(string) The login name associated with the Account.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(string) An e-mail address that is current set in the Account.`,
				},
				resource.Attribute{
					Name:        "cns_enabled",
					Description: `(boolean) Whether the Container Name Service (CNS) is enabled for the Account. [1]: /docs/providers/triton/index.html`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "triton_datacenter",
			Category:         "Data Sources",
			ShortDescription: `The ` + "`" + `triton_datacenter` + "`" + ` data source queries Triton for Data Center information.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(string) The name of the Data Center.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(string) The endpoint URL of the Data Center. [1]: /docs/providers/triton/index.html`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(string) The name of the Data Center.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(string) The endpoint URL of the Data Center. [1]: /docs/providers/triton/index.html`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "triton_fabric_network",
			Category:         "Data Sources",
			ShortDescription: `The ` + "`" + `triton_fabric_network` + "`" + ` data source queries Triton for Fabric Network information (e.g., subnet CIDR, gateway, static routes, etc.) based on the name of the Fabric Network and ID of the VLAN on which the network has been created.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "vlan_id",
					Description: `(integer)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(string) The name of the Fabric Network.`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `(boolean) Whether this Fabric Network is a public or private [RFC1918][3] network.`,
				},
				resource.Attribute{
					Name:        "fabric",
					Description: `(boolean) Whether this network is created on a [Fabric][4]. This is always`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(string) The description of the Fabric Network, if any.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `(string) A [CIDR][5] block used for the Fabric Network.`,
				},
				resource.Attribute{
					Name:        "provision_start_ip",
					Description: `(string) The first IP address on this network that may be assigned.`,
				},
				resource.Attribute{
					Name:        "provision_end_ip",
					Description: `(string) The last IP address on this network that may be assigned.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `(string) An IP address of the gateway on this network, if any.`,
				},
				resource.Attribute{
					Name:        "resolvers",
					Description: `(list) A list of IP addresses of DNS resolvers on this network.`,
				},
				resource.Attribute{
					Name:        "routes",
					Description: `(map) A map of static routes (using the [CIDR][5] notation) and corresponding gateways on this network, if any.`,
				},
				resource.Attribute{
					Name:        "internet_nat",
					Description: `(boolean) Whether the gateway on this network is also provisioned with the Internet NAT zone.`,
				},
				resource.Attribute{
					Name:        "vlan_id",
					Description: `(integer) The unique identifier (VLAN ID) of the Fabric VLAN. [1]: /docs/providers/triton/d/triton_fabric_vlan.html [2]: https://docs.joyent.com/public-cloud/network/sdn#vlans [3]: https://tools.ietf.org/html/rfc1918 [4]: https://docs.joyent.com/public-cloud/network/sdn [5]: https://tools.ietf.org/html/rfc4632`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(string) The name of the Fabric Network.`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `(boolean) Whether this Fabric Network is a public or private [RFC1918][3] network.`,
				},
				resource.Attribute{
					Name:        "fabric",
					Description: `(boolean) Whether this network is created on a [Fabric][4]. This is always`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(string) The description of the Fabric Network, if any.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `(string) A [CIDR][5] block used for the Fabric Network.`,
				},
				resource.Attribute{
					Name:        "provision_start_ip",
					Description: `(string) The first IP address on this network that may be assigned.`,
				},
				resource.Attribute{
					Name:        "provision_end_ip",
					Description: `(string) The last IP address on this network that may be assigned.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `(string) An IP address of the gateway on this network, if any.`,
				},
				resource.Attribute{
					Name:        "resolvers",
					Description: `(list) A list of IP addresses of DNS resolvers on this network.`,
				},
				resource.Attribute{
					Name:        "routes",
					Description: `(map) A map of static routes (using the [CIDR][5] notation) and corresponding gateways on this network, if any.`,
				},
				resource.Attribute{
					Name:        "internet_nat",
					Description: `(boolean) Whether the gateway on this network is also provisioned with the Internet NAT zone.`,
				},
				resource.Attribute{
					Name:        "vlan_id",
					Description: `(integer) The unique identifier (VLAN ID) of the Fabric VLAN. [1]: /docs/providers/triton/d/triton_fabric_vlan.html [2]: https://docs.joyent.com/public-cloud/network/sdn#vlans [3]: https://tools.ietf.org/html/rfc1918 [4]: https://docs.joyent.com/public-cloud/network/sdn [5]: https://tools.ietf.org/html/rfc4632`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "triton_fabric_vlan",
			Category:         "Data Sources",
			ShortDescription: `The ` + "`" + `triton_fabric_vlan` + "`" + ` data source queries Triton for Fabric VLAN information (e.g., VLAN ID, etc.) based either on the name, VLAN ID or description of the Fabric VLAN.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(string) Optional. The name of the Fabric VLAN.`,
				},
				resource.Attribute{
					Name:        "vlan_id",
					Description: `(integer) Optional. The unique identifier (VLAN ID) of the Fabric VLAN.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(string) Optional. The description of the Fabric VLAN. ~>`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(string) The name of the Fabric VLAN, if any.`,
				},
				resource.Attribute{
					Name:        "vlan_id",
					Description: `(integer) The unique identifier (VLAN ID) of the Fabric VLAN.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(string) The description of the Fabric VLAN, if any. [1]: https://docs.joyent.com/public-cloud/network/sdn#vlans [2]: https://en.wikipedia.org/wiki/Glob_(programming)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(string) The name of the Fabric VLAN, if any.`,
				},
				resource.Attribute{
					Name:        "vlan_id",
					Description: `(integer) The unique identifier (VLAN ID) of the Fabric VLAN.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(string) The description of the Fabric VLAN, if any. [1]: https://docs.joyent.com/public-cloud/network/sdn#vlans [2]: https://en.wikipedia.org/wiki/Glob_(programming)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "triton_image",
			Category:         "Data Sources",
			ShortDescription: `The ` + "`" + `triton_image` + "`" + ` data source queries the Triton Image API for image IDs.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(string) The name of the image`,
				},
				resource.Attribute{
					Name:        "os",
					Description: `(string) The underlying operating system for the image`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(string) The version for the image`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `(boolean) Whether to return public as well as private images`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(string) The state of the image. By default, only ` + "`" + `active` + "`" + ` images are shown. Must be one of: ` + "`" + `active` + "`" + `, ` + "`" + `unactivated` + "`" + `, ` + "`" + `disabled` + "`" + `, ` + "`" + `creating` + "`" + `, ` + "`" + `failed` + "`" + ` or ` + "`" + `all` + "`" + `, though the default is sufficient in almost every case.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `(string) The UUID of the account which owns the image`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(string) The image type. Must be one of: ` + "`" + `zone-dataset` + "`" + `, ` + "`" + `lx-dataset` + "`" + `, ` + "`" + `zvol` + "`" + `, ` + "`" + `docker` + "`" + ` or ` + "`" + `other` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "most_recent",
					Description: `(bool) If more than one result is returned, use the most recent Image. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(string) - The identifier representing the image in Triton.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(string) - The identifier representing the image in Triton.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "triton_network",
			Category:         "Data Sources",
			ShortDescription: `The ` + "`" + `triton_network` + "`" + ` data source queries Triton for Network information (e.g., Network ID, etc.) based on the name of the Network.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(string)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(string) The unique identifier of the Network.`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `(boolean) Whether this Network is a public or private [RFC1918][1] network.`,
				},
				resource.Attribute{
					Name:        "fabric",
					Description: `(boolean) Whether this Network is created on a [Fabric][2]. [1]: https://tools.ietf.org/html/rfc1918 [2]: https://docs.joyent.com/public-cloud/network/sdn`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(string) The unique identifier of the Network.`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `(boolean) Whether this Network is a public or private [RFC1918][1] network.`,
				},
				resource.Attribute{
					Name:        "fabric",
					Description: `(boolean) Whether this Network is created on a [Fabric][2]. [1]: https://tools.ietf.org/html/rfc1918 [2]: https://docs.joyent.com/public-cloud/network/sdn`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"triton_account":        0,
		"triton_datacenter":     1,
		"triton_fabric_network": 2,
		"triton_fabric_vlan":    3,
		"triton_image":          4,
		"triton_network":        5,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
