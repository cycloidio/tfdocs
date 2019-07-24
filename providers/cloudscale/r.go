package aws

import (
	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*Resource{

		&resource.Resource{
			Name:             "",
			Type:             "cloudscale_floating_ip",
			Category:         "Resources",
			ShortDescription: `Provides a cloudscale.ch Floating IP resource. This can be used to create, modify, and delete Floating IPs.`,
			Description:      ``,
			Keywords: []string{
				"floating",
				"ip",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "server",
					Description: `(Required) Assign the Floating IP to this server (UUID).`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `(Required) ` + "`" + `4` + "`" + ` or ` + "`" + `6` + "`" + `, for an IPv4 or IPv6 address or network respectively.`,
				},
				resource.Attribute{
					Name:        "prefix_length",
					Description: `(Optional) If you want to assign an entire network instead of a single IP address to your server, you must specify the prefix length. Currently, there is only support for ` + "`" + `ip_version=6` + "`" + ` and ` + "`" + `prefix_length=56` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "reverse_ptr",
					Description: `(Optional) You can specify the PTR record (reverse DNS pointer) in case of a single Floating IP address. The following arguments are supported when updating Floating IPs:`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `(Required) (Re-)Assign the Floating IP to this server (UUID). ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current resource.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `The CIDR notation of the Floating IP address or network, e.g. ` + "`" + `192.0.2.123/32` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `The IP address of the server that your Floating IP is currently assigned to.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current resource.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `The CIDR notation of the Floating IP address or network, e.g. ` + "`" + `192.0.2.123/32` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `The IP address of the server that your Floating IP is currently assigned to.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudscale_server",
			Category:         "Resources",
			ShortDescription: `Provides a cloudscale.ch Server resource. This can be used to create, modify, and delete servers.`,
			Description:      ``,
			Keywords: []string{
				"server",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the new server. The name has to be a valid host name or a fully qualified domain name (FQDN).`,
				},
				resource.Attribute{
					Name:        "flavor_slug",
					Description: `(Required) The slug (name) of the flavor to use for the new server. Possible values can be found in our [API documentation](https://www.cloudscale.ch/en/api/v1#flavors).`,
				},
				resource.Attribute{
					Name:        "image_slug",
					Description: `(Required) The slug (name) of the image to use for the new server. Possible values can be found in our [API documentation](https://www.cloudscale.ch/en/api/v1#images).`,
				},
				resource.Attribute{
					Name:        "ssh_keys",
					Description: `(Required) A list of SSH public keys. Use the full content of your \`,
				},
				resource.Attribute{
					Name:        "volume_size_gb",
					Description: `(Optional) The size in GB of the SSD root volume of the new server. If this parameter is not specified, the value will be set to 10. Valid values are either 10 or multiples of 50.`,
				},
				resource.Attribute{
					Name:        "bulk_volume_size_gb",
					Description: `(Optional, Deprecated) The size in GB of the bulk storage volume of the new server. If this parameter is not specified, no bulk storage volume will be attached to the server. Valid values are multiples of 100.`,
				},
				resource.Attribute{
					Name:        "use_public_network",
					Description: `(Optional) Attach/detach the public network interface to/from the new server. Can be ` + "`" + `true` + "`" + ` (default) or ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "use_private_network",
					Description: `(Optional) Attach/detach the private network interface to/from the new server. Can be ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + ` (default).`,
				},
				resource.Attribute{
					Name:        "use_ipv6",
					Description: `(Optional) Enable/disable IPv6 on the public network interface of the new server. Can be ` + "`" + `true` + "`" + ` (default) or ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional) User data (custom cloud-config settings) to use for the new server. Needs to be valid YAML. A default configuration will be used if this parameter is not specified or set to null. Use only if you are an advanced user with knowledge of cloud-config and cloud-init.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The desired state of a server. Can be ` + "`" + `running` + "`" + ` (default) or ` + "`" + `stopped` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "allow_stopping_for_update",
					Description: `(Optional) If true, allows Terraform to stop the instance to update its properties. If you try to update a property that requires stopping the instance without setting this field, the update will fail. The following arguments are supported when updating servers:`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The desired state of a server. Can be ` + "`" + `running` + "`" + ` (default) or ` + "`" + `stopped` + "`" + `. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of this server.`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current resource.`,
				},
				resource.Attribute{
					Name:        "ssh_fingerprints",
					Description: `A list of SSH host key fingerprints (strings) of this server.`,
				},
				resource.Attribute{
					Name:        "ssh_host_keys",
					Description: `A list of SSH host keys (strings) of this server.`,
				},
				resource.Attribute{
					Name:        "volumes",
					Description: `A list of volume objects attached to this server. Each volume object has three attributes:`,
				},
				resource.Attribute{
					Name:        "device_path",
					Description: `The path (string) to the volume on your server (e.g. ` + "`" + `/dev/vda` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "size_gb",
					Description: `The size (int) of the volume in GB. Typically matches ` + "`" + `volume_size_gb` + "`" + ` or ` + "`" + `bulk_volume_size_gb` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `A string. Either ` + "`" + `ssd` + "`" + ` or ` + "`" + `bulk` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "public_ipv4",
					Description: `The first ` + "`" + `public` + "`" + ` IPv4 address of this server. The returned IP address may be ` + "`" + `""` + "`" + ` if the server does not have a public IPv4.`,
				},
				resource.Attribute{
					Name:        "private_ipv4",
					Description: `The first ` + "`" + `private` + "`" + ` IPv4 address of this server. The returned IP address may be ` + "`" + `""` + "`" + ` if the server does not have private networking enabled.`,
				},
				resource.Attribute{
					Name:        "public_ipv6",
					Description: `The first ` + "`" + `public` + "`" + ` IPv6 address of this server. The returned IP address may be ` + "`" + `""` + "`" + ` if the server does not have a public IPv6.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `A list of interface objects attached to this server. Each interface object has two attributes:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Either ` + "`" + `public` + "`" + ` or ` + "`" + `private` + "`" + `. Public interfaces are connected to the Internet, while private interfaces are not.`,
				},
				resource.Attribute{
					Name:        "addresses",
					Description: `A list of address objects:`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `An IPv4 or IPv6 address that has been assigned to this server.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `An IPv4 or IPv6 address that represents the default gateway for this interface.`,
				},
				resource.Attribute{
					Name:        "prefix_length",
					Description: `The prefix length for this IP address, typically 24 for IPv4 and 128 for IPv6.`,
				},
				resource.Attribute{
					Name:        "reverse_ptr",
					Description: `The PTR record (reverse DNS pointer) for this IP address. If you use an FQDN as your server name it will automatically be used here.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The IP version, either ` + "`" + `4` + "`" + ` or ` + "`" + `6` + "`" + `.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of this server.`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current resource.`,
				},
				resource.Attribute{
					Name:        "ssh_fingerprints",
					Description: `A list of SSH host key fingerprints (strings) of this server.`,
				},
				resource.Attribute{
					Name:        "ssh_host_keys",
					Description: `A list of SSH host keys (strings) of this server.`,
				},
				resource.Attribute{
					Name:        "volumes",
					Description: `A list of volume objects attached to this server. Each volume object has three attributes:`,
				},
				resource.Attribute{
					Name:        "device_path",
					Description: `The path (string) to the volume on your server (e.g. ` + "`" + `/dev/vda` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "size_gb",
					Description: `The size (int) of the volume in GB. Typically matches ` + "`" + `volume_size_gb` + "`" + ` or ` + "`" + `bulk_volume_size_gb` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `A string. Either ` + "`" + `ssd` + "`" + ` or ` + "`" + `bulk` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "public_ipv4",
					Description: `The first ` + "`" + `public` + "`" + ` IPv4 address of this server. The returned IP address may be ` + "`" + `""` + "`" + ` if the server does not have a public IPv4.`,
				},
				resource.Attribute{
					Name:        "private_ipv4",
					Description: `The first ` + "`" + `private` + "`" + ` IPv4 address of this server. The returned IP address may be ` + "`" + `""` + "`" + ` if the server does not have private networking enabled.`,
				},
				resource.Attribute{
					Name:        "public_ipv6",
					Description: `The first ` + "`" + `public` + "`" + ` IPv6 address of this server. The returned IP address may be ` + "`" + `""` + "`" + ` if the server does not have a public IPv6.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `A list of interface objects attached to this server. Each interface object has two attributes:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Either ` + "`" + `public` + "`" + ` or ` + "`" + `private` + "`" + `. Public interfaces are connected to the Internet, while private interfaces are not.`,
				},
				resource.Attribute{
					Name:        "addresses",
					Description: `A list of address objects:`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `An IPv4 or IPv6 address that has been assigned to this server.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `An IPv4 or IPv6 address that represents the default gateway for this interface.`,
				},
				resource.Attribute{
					Name:        "prefix_length",
					Description: `The prefix length for this IP address, typically 24 for IPv4 and 128 for IPv6.`,
				},
				resource.Attribute{
					Name:        "reverse_ptr",
					Description: `The PTR record (reverse DNS pointer) for this IP address. If you use an FQDN as your server name it will automatically be used here.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The IP version, either ` + "`" + `4` + "`" + ` or ` + "`" + `6` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudscale_server_group",
			Category:         "Resources",
			ShortDescription: `Provides a cloudscale.ch server group resource. This can be used to create, modify, and delete server groups.`,
			Description:      ``,
			Keywords: []string{
				"server",
				"group",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the new server group.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of the server group can currently only be ` + "`" + `"anti-affinity"` + "`" + `. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current resource.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudscale_volume",
			Category:         "Resources",
			ShortDescription: `Provides a cloudscale.ch Volume (block storage) resource. This can be used to create, modify, and delete Volumes.`,
			Description:      ``,
			Keywords: []string{
				"volume",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the new volume.`,
				},
				resource.Attribute{
					Name:        "size_gb",
					Description: `(Required) The volume size in GB. Valid values are multiples of 1 for type "ssd" and multiples of 100 for type "bulk".`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) For SSD/NVMe volumes specify "ssd" (default) or use "bulk" for our HDD cluster with NVMe caching. This is the only attribute that cannot be altered.`,
				},
				resource.Attribute{
					Name:        "server_uuids",
					Description: `(Optional) A list of server UUIDs. Default to an empty list. Currently a volume can only be attached to one server UUID. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current resource.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current resource.`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"cloudscale_floating_ip":  0,
		"cloudscale_server":       1,
		"cloudscale_server_group": 2,
		"cloudscale_volume":       3,
	}
)

func GetResource(r string) (*resouce.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs]
}
