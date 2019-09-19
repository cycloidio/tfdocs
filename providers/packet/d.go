package packet

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "packet_device",
			Category:         "Data Sources",
			ShortDescription: `Provides a Packet device datasource. This can be used to read existing devices.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required) The device name`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The id of the project in which the devices exists ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "access_private_ipv4",
					Description: `The ipv4 private IP assigned to the device`,
				},
				resource.Attribute{
					Name:        "access_public_ipv4",
					Description: `The ipv4 management IP assigned to the device`,
				},
				resource.Attribute{
					Name:        "access_public_ipv6",
					Description: `The ipv6 management IP assigned to the device`,
				},
				resource.Attribute{
					Name:        "billing_cycle",
					Description: `The billing cycle of the device (monthly or hourly)`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `The facility where the device is deployed.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description string for the device`,
				},
				resource.Attribute{
					Name:        "hardware_reservation_id",
					Description: `The id of hardware reservation which this device occupies`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the device`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `The device's private and public IP (v4 and v6) network details. When a device is run without any special network configuration, it will have 3 networks:`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `IPv4 or IPv6 address string`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `Bit length of the network mask of the address`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `Address of router`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `Whether the address is routable from the Internet`,
				},
				resource.Attribute{
					Name:        "family",
					Description: `IP version - "4" or "6"`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `L2 network type of the device, one of "layer3", "layer2-bonded", "layer2-individual", "hybrid"`,
				},
				resource.Attribute{
					Name:        "operating_system",
					Description: `The operating system running on the device`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `The hardware config of the device`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `Ports assigned to the device`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the port (e.g. ` + "`" + `eth0` + "`" + `, or ` + "`" + `bond0` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the port`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the port (e.g. ` + "`" + `NetworkPort` + "`" + ` or ` + "`" + `NetworkBondPort` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `MAC address assigned to the port`,
				},
				resource.Attribute{
					Name:        "bonded",
					Description: `Whether this port is part of a bond in bonded network setup`,
				},
				resource.Attribute{
					Name:        "root_password",
					Description: `Root password to the server (if still available)`,
				},
				resource.Attribute{
					Name:        "ssh_key_ids",
					Description: `List of IDs of SSH keys deployed in the device, can be both user or project SSH keys`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the device`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags attached to the device`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "access_private_ipv4",
					Description: `The ipv4 private IP assigned to the device`,
				},
				resource.Attribute{
					Name:        "access_public_ipv4",
					Description: `The ipv4 management IP assigned to the device`,
				},
				resource.Attribute{
					Name:        "access_public_ipv6",
					Description: `The ipv6 management IP assigned to the device`,
				},
				resource.Attribute{
					Name:        "billing_cycle",
					Description: `The billing cycle of the device (monthly or hourly)`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `The facility where the device is deployed.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description string for the device`,
				},
				resource.Attribute{
					Name:        "hardware_reservation_id",
					Description: `The id of hardware reservation which this device occupies`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the device`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `The device's private and public IP (v4 and v6) network details. When a device is run without any special network configuration, it will have 3 networks:`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `IPv4 or IPv6 address string`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `Bit length of the network mask of the address`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `Address of router`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `Whether the address is routable from the Internet`,
				},
				resource.Attribute{
					Name:        "family",
					Description: `IP version - "4" or "6"`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `L2 network type of the device, one of "layer3", "layer2-bonded", "layer2-individual", "hybrid"`,
				},
				resource.Attribute{
					Name:        "operating_system",
					Description: `The operating system running on the device`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `The hardware config of the device`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `Ports assigned to the device`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the port (e.g. ` + "`" + `eth0` + "`" + `, or ` + "`" + `bond0` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the port`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the port (e.g. ` + "`" + `NetworkPort` + "`" + ` or ` + "`" + `NetworkBondPort` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `MAC address assigned to the port`,
				},
				resource.Attribute{
					Name:        "bonded",
					Description: `Whether this port is part of a bond in bonded network setup`,
				},
				resource.Attribute{
					Name:        "root_password",
					Description: `Root password to the server (if still available)`,
				},
				resource.Attribute{
					Name:        "ssh_key_ids",
					Description: `List of IDs of SSH keys deployed in the device, can be both user or project SSH keys`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the device`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags attached to the device`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "packet_operating_system",
			Category:         "Data Sources",
			ShortDescription: `Get a Packet operating system image`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "distro",
					Description: `(Optional) Name of the OS distribution.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name or part of the name of the distribution. Case insensitive.`,
				},
				resource.Attribute{
					Name:        "provisionable_on",
					Description: `(Optional) Plan name.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) Version of the distribution ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Operating system slug`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `Operating system slug (same as ` + "`" + `id` + "`" + `)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Operating system slug`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `Operating system slug (same as ` + "`" + `id` + "`" + `)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "packet_precreated_ip_block",
			Category:         "Data Sources",
			ShortDescription: `Load automatically created IP blocks from your Packet project`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) ID of the project where the searched block should be.`,
				},
				resource.Attribute{
					Name:        "address_family",
					Description: `(Required) 4 or 6, depending on which block you are looking for.`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `(Required) Whether to look for public or private block.`,
				},
				resource.Attribute{
					Name:        "global",
					Description: `(Optional) Whether to look for global block. Default is false for backward compatibility.`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `(Optional) Facility of the searched block. (Optional) Only allowed for non-global blocks. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "cidr_notation",
					Description: `CIDR notation of the looked up block.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cidr_notation",
					Description: `CIDR notation of the looked up block.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "packet_spot_market_price",
			Category:         "Data Sources",
			ShortDescription: `Get a Packet Spot Market Price`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "facility",
					Description: `(Required) Name of the facility.`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `(Required) Name of the plan. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "price",
					Description: `Current spot market price for given plan in given facility.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "price",
					Description: `Current spot market price for given plan in given facility.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"packet_device":              0,
		"packet_operating_system":    1,
		"packet_precreated_ip_block": 2,
		"packet_spot_market_price":   3,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
