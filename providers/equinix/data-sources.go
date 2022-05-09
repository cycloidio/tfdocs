package equinix

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_ecx_l2_sellerprofile",
			Category:         "Fabric",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"fabric",
				"ecx",
				"l2",
				"sellerprofile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the seller profile.`,
				},
				resource.Attribute{
					Name:        "organization_name",
					Description: `(Optional) Name of seller's organization.`,
				},
				resource.Attribute{
					Name:        "organization_global_name",
					Description: `(Optional) Name of seller's global organization. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `Unique identifier of the seller profile.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Seller Profile text description.`,
				},
				resource.Attribute{
					Name:        "speed_from_api",
					Description: `Boolean that indicates if seller is deriving connection speed from an API call.`,
				},
				resource.Attribute{
					Name:        "speed_customization_allowed",
					Description: `Boolean that indicates if seller allows customer to enter a custom connection speed.`,
				},
				resource.Attribute{
					Name:        "redundancy_required",
					Description: `Boolean that indicate if seller requires connections to be redundant`,
				},
				resource.Attribute{
					Name:        "encapsulation",
					Description: `Seller profile's encapsulation (either Dot1q or QinQ).`,
				},
				resource.Attribute{
					Name:        "speed_band",
					Description: `One or more specifications of speed/bandwidth supported by given seller profile. See [Speed Band Attribute](#speed-band-attribute) below for more details.`,
				},
				resource.Attribute{
					Name:        "metro",
					Description: `One or more specifications of metro locations supported by seller profile. See [Metro Attribute](#metro-attribute) below for more details.`,
				},
				resource.Attribute{
					Name:        "additional_info",
					Description: `One or more specifications of additional buyer information attributes that can be provided in connection definition that uses given seller profile. See [Additional Info Attribute](#additional-info-attribute) below for more details. ### Speed Band Attribute Each element in the ` + "`" + `speed_band` + "`" + ` set exports:`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `Speed/bandwidth supported by given service profile.`,
				},
				resource.Attribute{
					Name:        "speed_unit",
					Description: `Unit of the speed/bandwidth supported by given service profile. ### Metro Attribute Each element in the ` + "`" + `metro` + "`" + ` set exports:`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `Location metro code.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Location metro name.`,
				},
				resource.Attribute{
					Name:        "ibxes",
					Description: `List of IBXes supported within given metro.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `List of regions supported within given. ### Additional Info Attribute Each element in the ` + "`" + `additional_info` + "`" + ` set exports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of additional information attribute.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Textual description of additional information attribute.`,
				},
				resource.Attribute{
					Name:        "data_type",
					Description: `Data type of additional information attribute. One of ` + "`" + `BOOLEAN` + "`" + `, ` + "`" + `INTEGER` + "`" + ` or ` + "`" + `STRING` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mandatory",
					Description: `Specifies if additional information is mandatory to create connection.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "uuid",
					Description: `Unique identifier of the seller profile.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Seller Profile text description.`,
				},
				resource.Attribute{
					Name:        "speed_from_api",
					Description: `Boolean that indicates if seller is deriving connection speed from an API call.`,
				},
				resource.Attribute{
					Name:        "speed_customization_allowed",
					Description: `Boolean that indicates if seller allows customer to enter a custom connection speed.`,
				},
				resource.Attribute{
					Name:        "redundancy_required",
					Description: `Boolean that indicate if seller requires connections to be redundant`,
				},
				resource.Attribute{
					Name:        "encapsulation",
					Description: `Seller profile's encapsulation (either Dot1q or QinQ).`,
				},
				resource.Attribute{
					Name:        "speed_band",
					Description: `One or more specifications of speed/bandwidth supported by given seller profile. See [Speed Band Attribute](#speed-band-attribute) below for more details.`,
				},
				resource.Attribute{
					Name:        "metro",
					Description: `One or more specifications of metro locations supported by seller profile. See [Metro Attribute](#metro-attribute) below for more details.`,
				},
				resource.Attribute{
					Name:        "additional_info",
					Description: `One or more specifications of additional buyer information attributes that can be provided in connection definition that uses given seller profile. See [Additional Info Attribute](#additional-info-attribute) below for more details. ### Speed Band Attribute Each element in the ` + "`" + `speed_band` + "`" + ` set exports:`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `Speed/bandwidth supported by given service profile.`,
				},
				resource.Attribute{
					Name:        "speed_unit",
					Description: `Unit of the speed/bandwidth supported by given service profile. ### Metro Attribute Each element in the ` + "`" + `metro` + "`" + ` set exports:`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `Location metro code.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Location metro name.`,
				},
				resource.Attribute{
					Name:        "ibxes",
					Description: `List of IBXes supported within given metro.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `List of regions supported within given. ### Additional Info Attribute Each element in the ` + "`" + `additional_info` + "`" + ` set exports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of additional information attribute.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Textual description of additional information attribute.`,
				},
				resource.Attribute{
					Name:        "data_type",
					Description: `Data type of additional information attribute. One of ` + "`" + `BOOLEAN` + "`" + `, ` + "`" + `INTEGER` + "`" + ` or ` + "`" + `STRING` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mandatory",
					Description: `Specifies if additional information is mandatory to create connection.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_ecx_l2_sellerprofiles",
			Category:         "Fabric",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"fabric",
				"ecx",
				"l2",
				"sellerprofiles",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to apply on returned seller profile names and filter search results.`,
				},
				resource.Attribute{
					Name:        "metro_codes",
					Description: `(Optional) List of metro codes of locations that should be served by resulting profiles.`,
				},
				resource.Attribute{
					Name:        "speed_bands",
					Description: `(Optional) List of speed bands that should be supported by resulting profiles.`,
				},
				resource.Attribute{
					Name:        "organization_name",
					Description: `(Optional) Name of seller's organization.`,
				},
				resource.Attribute{
					Name:        "organization_global_name",
					Description: `(Optional) Name of seller's global organization. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "profiles",
					Description: `List of resulting profiles. Each element in the ` + "`" + `profiles` + "`" + ` list exports all [Service Profile Attributes](./equinix_ecx_l2_sellerprofile.md#attributes-reference).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "profiles",
					Description: `List of resulting profiles. Each element in the ` + "`" + `profiles` + "`" + ` list exports all [Service Profile Attributes](./equinix_ecx_l2_sellerprofile.md#attributes-reference).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_ecx_port",
			Category:         "Fabric",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"fabric",
				"ecx",
				"port",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the port. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `Unique identifier of the port.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the port.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Port location region.`,
				},
				resource.Attribute{
					Name:        "ibx",
					Description: `Port location Equinix Business Exchange (IBX).`,
				},
				resource.Attribute{
					Name:        "metro_code",
					Description: `Port location metro code.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority of the device (primary / secondary) where the port resides.`,
				},
				resource.Attribute{
					Name:        "encapsulation",
					Description: `The VLAN encapsulation of the port (Dot1q or QinQ).`,
				},
				resource.Attribute{
					Name:        "buyout",
					Description: `Boolean value that indicates whether the port supports unlimited connections. If ` + "`" + `false` + "`" + `, the port is a standard port with limited connections. If ` + "`" + `true` + "`" + `, the port is an ` + "`" + `unlimited connections` + "`" + ` port that allows multiple connections at no additional charge.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `Port Bandwidth in bytes.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Port status that indicates whether a port has been assigned or is ready for connection.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "uuid",
					Description: `Unique identifier of the port.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the port.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Port location region.`,
				},
				resource.Attribute{
					Name:        "ibx",
					Description: `Port location Equinix Business Exchange (IBX).`,
				},
				resource.Attribute{
					Name:        "metro_code",
					Description: `Port location metro code.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority of the device (primary / secondary) where the port resides.`,
				},
				resource.Attribute{
					Name:        "encapsulation",
					Description: `The VLAN encapsulation of the port (Dot1q or QinQ).`,
				},
				resource.Attribute{
					Name:        "buyout",
					Description: `Boolean value that indicates whether the port supports unlimited connections. If ` + "`" + `false` + "`" + `, the port is a standard port with limited connections. If ` + "`" + `true` + "`" + `, the port is an ` + "`" + `unlimited connections` + "`" + ` port that allows multiple connections at no additional charge.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `Port Bandwidth in bytes.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Port status that indicates whether a port has been assigned or is ready for connection.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_metal_connection",
			Category:         "Metal",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"metal",
				"connection",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "connection_id",
					Description: `(Required) ID of the connection resource. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the connection resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the connection resource.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `String list of tags.`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `Slug of a facility to which the connection belongs.`,
				},
				resource.Attribute{
					Name:        "metro",
					Description: `Slug of a metro to which the connection belongs.`,
				},
				resource.Attribute{
					Name:        "organization_id",
					Description: `ID of organization to which the connection belongs.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of project to which the connection belongs.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the connection.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `Fabric Token for the [Equinix Fabric Portal](https://fabric.equinix.com/dashboard).`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Connection type, dedicated or shared.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Mode for connections in IBX facilities with the dedicated type - standard or tunnel.`,
				},
				resource.Attribute{
					Name:        "redundancy",
					Description: `Connection redundancy, reduntant or primary.`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `Connection speed in bits per second.`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `List of connection ports - primary (` + "`" + `ports[0]` + "`" + `) and secondary (` + "`" + `ports[1]` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Port name.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Port UUID.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `Port role - primary or secondary.`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `Port speed in bits per second.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Port status.`,
				},
				resource.Attribute{
					Name:        "link_status",
					Description: `Port link status.`,
				},
				resource.Attribute{
					Name:        "virtual_circuit_ids",
					Description: `List of IDs of virtual cicruits attached to this port.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Description of the connection resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the connection resource.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `String list of tags.`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `Slug of a facility to which the connection belongs.`,
				},
				resource.Attribute{
					Name:        "metro",
					Description: `Slug of a metro to which the connection belongs.`,
				},
				resource.Attribute{
					Name:        "organization_id",
					Description: `ID of organization to which the connection belongs.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of project to which the connection belongs.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the connection.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `Fabric Token for the [Equinix Fabric Portal](https://fabric.equinix.com/dashboard).`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Connection type, dedicated or shared.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Mode for connections in IBX facilities with the dedicated type - standard or tunnel.`,
				},
				resource.Attribute{
					Name:        "redundancy",
					Description: `Connection redundancy, reduntant or primary.`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `Connection speed in bits per second.`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `List of connection ports - primary (` + "`" + `ports[0]` + "`" + `) and secondary (` + "`" + `ports[1]` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Port name.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Port UUID.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `Port role - primary or secondary.`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `Port speed in bits per second.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Port status.`,
				},
				resource.Attribute{
					Name:        "link_status",
					Description: `Port link status.`,
				},
				resource.Attribute{
					Name:        "virtual_circuit_ids",
					Description: `List of IDs of virtual cicruits attached to this port.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_metal_device",
			Category:         "Metal",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"metal",
				"device",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional) The device name.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) The id of the project in which the devices exists.`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `(Optional) Device ID. ->`,
				},
				resource.Attribute{
					Name:        "access_private_ipv4",
					Description: `The ipv4 private IP assigned to the device.`,
				},
				resource.Attribute{
					Name:        "access_public_ipv4",
					Description: `The ipv4 management IP assigned to the device.`,
				},
				resource.Attribute{
					Name:        "access_public_ipv6",
					Description: `The ipv6 management IP assigned to the device.`,
				},
				resource.Attribute{
					Name:        "billing_cycle",
					Description: `The billing cycle of the device (monthly or hourly).`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `The facility where the device is deployed.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description string for the device.`,
				},
				resource.Attribute{
					Name:        "hardware_reservation_id",
					Description: `The id of hardware reservation which this device occupies.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the device.`,
				},
				resource.Attribute{
					Name:        "metro",
					Description: `The metro where the device is deployed`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `The device's private and public IP (v4 and v6) network details. See [Network Attribute](#network-attribute) below for more details.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `L2 network type of the device, one of ` + "`" + `layer3` + "`" + `, ` + "`" + `layer2-bonded` + "`" + `, ` + "`" + `layer2-individual` + "`" + `, ` + "`" + `hybrid` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "operating_system",
					Description: `The operating system running on the device.`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `The hardware config of the device.`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `List of ports assigned to the device. See [Ports Attribute](#ports-attribute) below for more details.`,
				},
				resource.Attribute{
					Name:        "root_password",
					Description: `Root password to the server (if still available).`,
				},
				resource.Attribute{
					Name:        "ssh_key_ids",
					Description: `List of IDs of SSH keys deployed in the device, can be both user or project SSH keys.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the device.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags attached to the device. ### Network Attribute When a device is run without any special network, it will have 3 networks:`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `IPv4 or IPv6 address string.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `Bit length of the network mask of the address.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `Address of router.`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `Whether the address is routable from the Internet.`,
				},
				resource.Attribute{
					Name:        "family",
					Description: `IP version. One of ` + "`" + `4` + "`" + `, ` + "`" + `6` + "`" + `. ### Ports Attribute Each element in the ` + "`" + `ports` + "`" + ` list exports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the port (e.g. ` + "`" + `eth0` + "`" + `, or ` + "`" + `bond0` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the port.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the port (e.g. ` + "`" + `NetworkPort` + "`" + ` or ` + "`" + `NetworkBondPort` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `MAC address assigned to the port.`,
				},
				resource.Attribute{
					Name:        "bonded",
					Description: `Whether this port is part of a bond in bonded network setup.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "access_private_ipv4",
					Description: `The ipv4 private IP assigned to the device.`,
				},
				resource.Attribute{
					Name:        "access_public_ipv4",
					Description: `The ipv4 management IP assigned to the device.`,
				},
				resource.Attribute{
					Name:        "access_public_ipv6",
					Description: `The ipv6 management IP assigned to the device.`,
				},
				resource.Attribute{
					Name:        "billing_cycle",
					Description: `The billing cycle of the device (monthly or hourly).`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `The facility where the device is deployed.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description string for the device.`,
				},
				resource.Attribute{
					Name:        "hardware_reservation_id",
					Description: `The id of hardware reservation which this device occupies.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the device.`,
				},
				resource.Attribute{
					Name:        "metro",
					Description: `The metro where the device is deployed`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `The device's private and public IP (v4 and v6) network details. See [Network Attribute](#network-attribute) below for more details.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `L2 network type of the device, one of ` + "`" + `layer3` + "`" + `, ` + "`" + `layer2-bonded` + "`" + `, ` + "`" + `layer2-individual` + "`" + `, ` + "`" + `hybrid` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "operating_system",
					Description: `The operating system running on the device.`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `The hardware config of the device.`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `List of ports assigned to the device. See [Ports Attribute](#ports-attribute) below for more details.`,
				},
				resource.Attribute{
					Name:        "root_password",
					Description: `Root password to the server (if still available).`,
				},
				resource.Attribute{
					Name:        "ssh_key_ids",
					Description: `List of IDs of SSH keys deployed in the device, can be both user or project SSH keys.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the device.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags attached to the device. ### Network Attribute When a device is run without any special network, it will have 3 networks:`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `IPv4 or IPv6 address string.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `Bit length of the network mask of the address.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `Address of router.`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `Whether the address is routable from the Internet.`,
				},
				resource.Attribute{
					Name:        "family",
					Description: `IP version. One of ` + "`" + `4` + "`" + `, ` + "`" + `6` + "`" + `. ### Ports Attribute Each element in the ` + "`" + `ports` + "`" + ` list exports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the port (e.g. ` + "`" + `eth0` + "`" + `, or ` + "`" + `bond0` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the port.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the port (e.g. ` + "`" + `NetworkPort` + "`" + ` or ` + "`" + `NetworkBondPort` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `MAC address assigned to the port.`,
				},
				resource.Attribute{
					Name:        "bonded",
					Description: `Whether this port is part of a bond in bonded network setup.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_metal_device_bgp_neighbors",
			Category:         "Metal",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"metal",
				"device",
				"bgp",
				"neighbors",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "device_id",
					Description: `(Required) UUID of BGP-enabled device whose neighbors to list. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "bgp_neighbors",
					Description: `array of BGP neighbor records with attributes:`,
				},
				resource.Attribute{
					Name:        "address_family",
					Description: `IP address version, 4 or 6.`,
				},
				resource.Attribute{
					Name:        "customer_as",
					Description: `Local autonomous system number.`,
				},
				resource.Attribute{
					Name:        "customer_ip",
					Description: `Local used peer IP address.`,
				},
				resource.Attribute{
					Name:        "md5_enabled",
					Description: `Whether BGP session is password enabled.`,
				},
				resource.Attribute{
					Name:        "md5_password",
					Description: `BGP session password in plaintext (not a checksum).`,
				},
				resource.Attribute{
					Name:        "multihop",
					Description: `Whether the neighbor is in EBGP multihop session.`,
				},
				resource.Attribute{
					Name:        "peer_as",
					Description: `Peer AS number (different than customer_as for EBGP).`,
				},
				resource.Attribute{
					Name:        "peer_ips",
					Description: `Array of IP addresses of this neighbor's peers.`,
				},
				resource.Attribute{
					Name:        "routes_in",
					Description: `Array of incoming routes.`,
				},
				resource.Attribute{
					Name:        "route",
					Description: `CIDR expression of route (IP/mask).`,
				},
				resource.Attribute{
					Name:        "exact",
					Description: `(bool) Whether the route is exact.`,
				},
				resource.Attribute{
					Name:        "routes_out",
					Description: `Array of outgoing routes in the same format.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "bgp_neighbors",
					Description: `array of BGP neighbor records with attributes:`,
				},
				resource.Attribute{
					Name:        "address_family",
					Description: `IP address version, 4 or 6.`,
				},
				resource.Attribute{
					Name:        "customer_as",
					Description: `Local autonomous system number.`,
				},
				resource.Attribute{
					Name:        "customer_ip",
					Description: `Local used peer IP address.`,
				},
				resource.Attribute{
					Name:        "md5_enabled",
					Description: `Whether BGP session is password enabled.`,
				},
				resource.Attribute{
					Name:        "md5_password",
					Description: `BGP session password in plaintext (not a checksum).`,
				},
				resource.Attribute{
					Name:        "multihop",
					Description: `Whether the neighbor is in EBGP multihop session.`,
				},
				resource.Attribute{
					Name:        "peer_as",
					Description: `Peer AS number (different than customer_as for EBGP).`,
				},
				resource.Attribute{
					Name:        "peer_ips",
					Description: `Array of IP addresses of this neighbor's peers.`,
				},
				resource.Attribute{
					Name:        "routes_in",
					Description: `Array of incoming routes.`,
				},
				resource.Attribute{
					Name:        "route",
					Description: `CIDR expression of route (IP/mask).`,
				},
				resource.Attribute{
					Name:        "exact",
					Description: `(bool) Whether the route is exact.`,
				},
				resource.Attribute{
					Name:        "routes_out",
					Description: `Array of outgoing routes in the same format.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_metal_facility",
			Category:         "Metal",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"metal",
				"facility",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "code",
					Description: `(Required) The facility code to search for facilities.`,
				},
				resource.Attribute{
					Name:        "features_required",
					Description: `(Optional) Set of feature strings that the facility must have. Some possible values are ` + "`" + `baremetal` + "`" + `, ` + "`" + `ibx` + "`" + `, ` + "`" + `storage` + "`" + `, ` + "`" + `global_ipv4` + "`" + `, ` + "`" + `backend_transfer` + "`" + `, ` + "`" + `layer_2` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `(Optional) One or more device plans for which the facility must have capacity.`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `(Required) Device plan that must be available in selected location.`,
				},
				resource.Attribute{
					Name:        "quantity",
					Description: `(Optional) Minimun number of devices that must be available in selected location. Default is ` + "`" + `1` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the facility.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the facility.`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `The features of the facility.`,
				},
				resource.Attribute{
					Name:        "metro",
					Description: `The metro code the facility is part of.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the facility.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the facility.`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `The features of the facility.`,
				},
				resource.Attribute{
					Name:        "metro",
					Description: `The metro code the facility is part of.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_metal_gateway",
			Category:         "Metal",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"metal",
				"gateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "gateway_id",
					Description: `(Required) UUID of the metal gateway resource to retrieve. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `UUID of the project where the gateway is scoped to.`,
				},
				resource.Attribute{
					Name:        "vlan_id",
					Description: `UUID of the VLAN where the gateway is scoped to.`,
				},
				resource.Attribute{
					Name:        "vrf_id",
					Description: `UUID of the VRF associated with the IP Reservation.`,
				},
				resource.Attribute{
					Name:        "ip_reservation_id",
					Description: `UUID of IP reservation block bound to the gateway.`,
				},
				resource.Attribute{
					Name:        "private_ipv4_subnet_size",
					Description: `Size of the private IPv4 subnet bound to this metal gateway. One of ` + "`" + `8` + "`" + `, ` + "`" + `16` + "`" + `, ` + "`" + `32` + "`" + `, ` + "`" + `64` + "`" + `, ` + "`" + `128` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Status of the gateway resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `UUID of the project where the gateway is scoped to.`,
				},
				resource.Attribute{
					Name:        "vlan_id",
					Description: `UUID of the VLAN where the gateway is scoped to.`,
				},
				resource.Attribute{
					Name:        "vrf_id",
					Description: `UUID of the VRF associated with the IP Reservation.`,
				},
				resource.Attribute{
					Name:        "ip_reservation_id",
					Description: `UUID of IP reservation block bound to the gateway.`,
				},
				resource.Attribute{
					Name:        "private_ipv4_subnet_size",
					Description: `Size of the private IPv4 subnet bound to this metal gateway. One of ` + "`" + `8` + "`" + `, ` + "`" + `16` + "`" + `, ` + "`" + `32` + "`" + `, ` + "`" + `64` + "`" + `, ` + "`" + `128` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Status of the gateway resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_metal_hardware_reservation",
			Category:         "Metal",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"metal",
				"hardware",
				"reservation",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the hardware reservation.`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `(Optional) UUID of device occupying the reservation. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the hardware reservation to look up.`,
				},
				resource.Attribute{
					Name:        "short_id",
					Description: `Reservation short ID.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `UUID of project this reservation is scoped to.`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `UUID of device occupying the reservation.`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `Plan type for the reservation.`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `Plan type for the reservation.`,
				},
				resource.Attribute{
					Name:        "provisionable",
					Description: `Flag indicating whether the reserved server is provisionable or not. Spare devices can't be provisioned unless they are activated first.`,
				},
				resource.Attribute{
					Name:        "spare",
					Description: `Flag indicating whether the Hardware Reservation is a spare. Spare Hardware Reservations are used when a Hardware Reservations requires service from Metal Equinix.`,
				},
				resource.Attribute{
					Name:        "switch_uuid",
					Description: `Switch short ID, can be used to determine if two devices are connected to the same switch.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the hardware reservation to look up.`,
				},
				resource.Attribute{
					Name:        "short_id",
					Description: `Reservation short ID.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `UUID of project this reservation is scoped to.`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `UUID of device occupying the reservation.`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `Plan type for the reservation.`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `Plan type for the reservation.`,
				},
				resource.Attribute{
					Name:        "provisionable",
					Description: `Flag indicating whether the reserved server is provisionable or not. Spare devices can't be provisioned unless they are activated first.`,
				},
				resource.Attribute{
					Name:        "spare",
					Description: `Flag indicating whether the Hardware Reservation is a spare. Spare Hardware Reservations are used when a Hardware Reservations requires service from Metal Equinix.`,
				},
				resource.Attribute{
					Name:        "switch_uuid",
					Description: `Switch short ID, can be used to determine if two devices are connected to the same switch.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_metal_ip_block_ranges",
			Category:         "Metal",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"metal",
				"ip",
				"block",
				"ranges",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) ID of the project from which to list the blocks.`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `(Optional) Facility code filtering the IP blocks. Global IPv4 blcoks will be listed anyway. If you omit this and metro, all the block from the project will be listed.`,
				},
				resource.Attribute{
					Name:        "metro",
					Description: `(Optional) Metro code filtering the IP blocks. Global IPv4 blcoks will be listed anyway. If you omit this and facility, all the block from the project will be listed. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "global_ipv4",
					Description: `list of CIDR expressions for Global IPv4 blocks in the project.`,
				},
				resource.Attribute{
					Name:        "public_ipv4",
					Description: `list of CIDR expressions for Public IPv4 blocks in the project.`,
				},
				resource.Attribute{
					Name:        "private_ipv4",
					Description: `list of CIDR expressions for Private IPv4 blocks in the project.`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `list of CIDR expressions for IPv6 blocks in the project.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "global_ipv4",
					Description: `list of CIDR expressions for Global IPv4 blocks in the project.`,
				},
				resource.Attribute{
					Name:        "public_ipv4",
					Description: `list of CIDR expressions for Public IPv4 blocks in the project.`,
				},
				resource.Attribute{
					Name:        "private_ipv4",
					Description: `list of CIDR expressions for Private IPv4 blocks in the project.`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `list of CIDR expressions for IPv6 blocks in the project.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_metal_metro",
			Category:         "Metal",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"metal",
				"metro",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "code",
					Description: `(Required) The facility code to search for facilities.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `(Optional) One or more device plans for which the facility must have capacity.`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `(Required) Device plan that must be available in selected location.`,
				},
				resource.Attribute{
					Name:        "quantity",
					Description: `(Optional) Minimun number of devices that must be available in selected location. Default is ` + "`" + `1` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the metro.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the metro.`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `The country of the metro.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the metro.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the metro.`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `The country of the metro.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_metal_operating_system",
			Category:         "Metal",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"metal",
				"operating",
				"system",
			},
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
					Description: `(Optional) Version of the distribution. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Operating system slug.`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `Operating system slug (same as ` + "`" + `id` + "`" + `).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Operating system slug.`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `Operating system slug (same as ` + "`" + `id` + "`" + `).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_metal_organization",
			Category:         "Metal",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"metal",
				"organization",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The organization name.`,
				},
				resource.Attribute{
					Name:        "organization_id",
					Description: `(Optional) The UUID of the organization resource. Exactly one of ` + "`" + `name` + "`" + ` or ` + "`" + `organization_id` + "`" + ` must be given. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "project_ids",
					Description: `UUIDs of project resources which belong to this organization.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description string.`,
				},
				resource.Attribute{
					Name:        "website",
					Description: `Website link.`,
				},
				resource.Attribute{
					Name:        "twitter",
					Description: `Twitter handle.`,
				},
				resource.Attribute{
					Name:        "logo",
					Description: `Logo URL.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "project_ids",
					Description: `UUIDs of project resources which belong to this organization.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description string.`,
				},
				resource.Attribute{
					Name:        "website",
					Description: `Website link.`,
				},
				resource.Attribute{
					Name:        "twitter",
					Description: `Twitter handle.`,
				},
				resource.Attribute{
					Name:        "logo",
					Description: `Logo URL.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_metal_port",
			Category:         "Metal",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"metal",
				"port",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "port_id",
					Description: `(Optional) ID of the port to read, conflicts with ` + "`" + `device_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `(Optional) Device UUID where to lookup the port.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the port to look up, i.e. ` + "`" + `bond0` + "`" + `, ` + "`" + `eth1` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `One of ` + "`" + `layer2-bonded` + "`" + `, ` + "`" + `layer2-individual` + "`" + `, ` + "`" + `layer3` + "`" + `, ` + "`" + `hybrid` + "`" + `, ` + "`" + `hybrid-bonded` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type is either ` + "`" + `NetworkBondPort` + "`" + ` for bond ports or ` + "`" + `NetworkPort` + "`" + ` for bondable ethernet ports.`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `MAC address of the port.`,
				},
				resource.Attribute{
					Name:        "bond_id",
					Description: `UUID of the bond port.`,
				},
				resource.Attribute{
					Name:        "bond_name",
					Description: `Name of the bond port.`,
				},
				resource.Attribute{
					Name:        "bonded",
					Description: `Flag indicating whether the port is bonded.`,
				},
				resource.Attribute{
					Name:        "disbond_supported",
					Description: `Flag indicating whether the port can be removed from a bond.`,
				},
				resource.Attribute{
					Name:        "native_vlan_id",
					Description: `UUID of native VLAN of the port.`,
				},
				resource.Attribute{
					Name:        "vlan_ids",
					Description: `UUIDs of attached VLANs.`,
				},
				resource.Attribute{
					Name:        "vxlan_ids",
					Description: `VXLAN ids of attached VLANs.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "network_type",
					Description: `One of ` + "`" + `layer2-bonded` + "`" + `, ` + "`" + `layer2-individual` + "`" + `, ` + "`" + `layer3` + "`" + `, ` + "`" + `hybrid` + "`" + `, ` + "`" + `hybrid-bonded` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type is either ` + "`" + `NetworkBondPort` + "`" + ` for bond ports or ` + "`" + `NetworkPort` + "`" + ` for bondable ethernet ports.`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `MAC address of the port.`,
				},
				resource.Attribute{
					Name:        "bond_id",
					Description: `UUID of the bond port.`,
				},
				resource.Attribute{
					Name:        "bond_name",
					Description: `Name of the bond port.`,
				},
				resource.Attribute{
					Name:        "bonded",
					Description: `Flag indicating whether the port is bonded.`,
				},
				resource.Attribute{
					Name:        "disbond_supported",
					Description: `Flag indicating whether the port can be removed from a bond.`,
				},
				resource.Attribute{
					Name:        "native_vlan_id",
					Description: `UUID of native VLAN of the port.`,
				},
				resource.Attribute{
					Name:        "vlan_ids",
					Description: `UUIDs of attached VLANs.`,
				},
				resource.Attribute{
					Name:        "vxlan_ids",
					Description: `VXLAN ids of attached VLANs.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_metal_precreated_ip_block",
			Category:         "Metal",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"metal",
				"precreated",
				"ip",
				"block",
			},
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
					Description: `(Optional) Facility of the searched block. (for non-global blocks).`,
				},
				resource.Attribute{
					Name:        "metro",
					Description: `(Optional) Metro of the searched block (for non-global blocks). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
			Type:             "equinix_equinix_metal_project",
			Category:         "Metal",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"metal",
				"project",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name which is used to look up the project.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) The UUID by which to look up the project. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "payment_method_id",
					Description: `The UUID of payment method for this project.`,
				},
				resource.Attribute{
					Name:        "organization_id",
					Description: `The UUID of this project's parent organization.`,
				},
				resource.Attribute{
					Name:        "backend_transfer",
					Description: `Whether Backend Transfer is enabled for this project.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The timestamp for when the project was created.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `The timestamp for the last time the project was updated.`,
				},
				resource.Attribute{
					Name:        "user_ids",
					Description: `List of UUIDs of user accounts which belong to this project.`,
				},
				resource.Attribute{
					Name:        "bgp_config",
					Description: `Optional BGP settings. Refer to [Equinix Metal guide for BGP](https://metal.equinix.com/developers/docs/networking/local-global-bgp/). The ` + "`" + `bgp_config` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "asn",
					Description: `Autonomous System Number for local BGP deployment.`,
				},
				resource.Attribute{
					Name:        "md5",
					Description: `Password for BGP session in plaintext (not a checksum).`,
				},
				resource.Attribute{
					Name:        "deployment_type",
					Description: `One of ` + "`" + `private` + "`" + `, ` + "`" + `public` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of BGP configuration in the project.`,
				},
				resource.Attribute{
					Name:        "max_prefix",
					Description: `The maximum number of route filters allowed per server.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "payment_method_id",
					Description: `The UUID of payment method for this project.`,
				},
				resource.Attribute{
					Name:        "organization_id",
					Description: `The UUID of this project's parent organization.`,
				},
				resource.Attribute{
					Name:        "backend_transfer",
					Description: `Whether Backend Transfer is enabled for this project.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The timestamp for when the project was created.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `The timestamp for the last time the project was updated.`,
				},
				resource.Attribute{
					Name:        "user_ids",
					Description: `List of UUIDs of user accounts which belong to this project.`,
				},
				resource.Attribute{
					Name:        "bgp_config",
					Description: `Optional BGP settings. Refer to [Equinix Metal guide for BGP](https://metal.equinix.com/developers/docs/networking/local-global-bgp/). The ` + "`" + `bgp_config` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "asn",
					Description: `Autonomous System Number for local BGP deployment.`,
				},
				resource.Attribute{
					Name:        "md5",
					Description: `Password for BGP session in plaintext (not a checksum).`,
				},
				resource.Attribute{
					Name:        "deployment_type",
					Description: `One of ` + "`" + `private` + "`" + `, ` + "`" + `public` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of BGP configuration in the project.`,
				},
				resource.Attribute{
					Name:        "max_prefix",
					Description: `The maximum number of route filters allowed per server.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_metal_project_ssh_key",
			Category:         "Metal",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"metal",
				"project",
				"ssh",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "search",
					Description: `(Optional) The name, fingerprint, or public_key of the SSH Key to search for in the Equinix Metal project.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The id of the SSH Key to search for in the Equinix Metal project.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) The Equinix Metal project id of the Equinix Metal SSH Key. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the key.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the SSH key.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `The text of the public key.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The ID of parent project.`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `The ID of parent project (same as project_id).`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `The fingerprint of the SSH key.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The timestamp for when the SSH key was created.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `The timestamp for the last time the SSH key was updated.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the key.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the SSH key.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `The text of the public key.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The ID of parent project.`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `The ID of parent project (same as project_id).`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `The fingerprint of the SSH key.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The timestamp for when the SSH key was created.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `The timestamp for the last time the SSH key was updated.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_metal_reserved_ip_block",
			Category:         "Metal",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"metal",
				"reserved",
				"ip",
				"block",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) UUID of the IP address block to look up.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) UUID of the project where the searched block should be.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) Block containing this IP address will be returned. ->`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `One of ` + "`" + `global_ipv4` + "`" + `, ` + "`" + `public_ipv4` + "`" + `, ` + "`" + `private_ipv4` + "`" + `, ` + "`" + `public_ipv6` + "`" + `,or ` + "`" + `vrf` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `One of ` + "`" + `global_ipv4` + "`" + `, ` + "`" + `public_ipv4` + "`" + `, ` + "`" + `private_ipv4` + "`" + `, ` + "`" + `public_ipv6` + "`" + `,or ` + "`" + `vrf` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_metal_spot_market_price",
			Category:         "Metal",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"metal",
				"spot",
				"market",
				"price",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "plan",
					Description: `(Required) Name of the plan.`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `(Optional) Name of the facility.`,
				},
				resource.Attribute{
					Name:        "metro",
					Description: `(Optional) Name of the metro. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_metal_spot_market_request",
			Category:         "Metal",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"metal",
				"spot",
				"market",
				"request",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "request_id",
					Description: `(Required) The id of the Spot Market Request. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "device_ids",
					Description: `List of IDs of devices spawned by the referenced Spot Market Request.`,
				},
				resource.Attribute{
					Name:        "devices_min",
					Description: `Miniumum number devices to be created.`,
				},
				resource.Attribute{
					Name:        "devices_max",
					Description: `Maximum number devices to be created.`,
				},
				resource.Attribute{
					Name:        "max_bid_price",
					Description: `Maximum price user is willing to pay per hour per device.`,
				},
				resource.Attribute{
					Name:        "facilities",
					Description: `Facility IDs where devices should be created.`,
				},
				resource.Attribute{
					Name:        "metro",
					Description: `Metro where devices should be created.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Project ID.`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `The device plan slug.`,
				},
				resource.Attribute{
					Name:        "end_at",
					Description: `Date and time When the spot market request will be ended.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "device_ids",
					Description: `List of IDs of devices spawned by the referenced Spot Market Request.`,
				},
				resource.Attribute{
					Name:        "devices_min",
					Description: `Miniumum number devices to be created.`,
				},
				resource.Attribute{
					Name:        "devices_max",
					Description: `Maximum number devices to be created.`,
				},
				resource.Attribute{
					Name:        "max_bid_price",
					Description: `Maximum price user is willing to pay per hour per device.`,
				},
				resource.Attribute{
					Name:        "facilities",
					Description: `Facility IDs where devices should be created.`,
				},
				resource.Attribute{
					Name:        "metro",
					Description: `Metro where devices should be created.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Project ID.`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `The device plan slug.`,
				},
				resource.Attribute{
					Name:        "end_at",
					Description: `Date and time When the spot market request will be ended.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_metal_virtual_circuit",
			Category:         "Metal",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"metal",
				"virtual",
				"circuit",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "virtual_circuit_id",
					Description: `(Required) ID of the virtual circuit resource ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the virtual circuit resource.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the virtal circuit.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of project to which the VC belongs.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description for the Virtual Circuit resource.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags for the Virtual Circuit resource.`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `Speed of the Virtual Circuit resource.`,
				},
				resource.Attribute{
					Name:        "vrf_id",
					Description: `UUID of the VLAN to associate.`,
				},
				resource.Attribute{
					Name:        "peer_asn",
					Description: `The BGP ASN of the peer. The same ASN may be the used across several VCs, but it cannot be the same as the local_asn of the VRF.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `A subnet from one of the IP blocks associated with the VRF that we will help create an IP reservation for. Can only be either a /30 or /31.`,
				},
				resource.Attribute{
					Name:        "metal_ip",
					Description: `The IP address that’s set as “our” IP that is configured on the rack_local_vlan SVI. Will default to the first usable IP in the subnet.`,
				},
				resource.Attribute{
					Name:        "customer_ip",
					Description: `The IP address set as the customer IP which the CSR switch will peer with. Will default to the other usable IP in the subnet.`,
				},
				resource.Attribute{
					Name:        "md5",
					Description: `The password that can be set for the VRF BGP peer`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name of the virtual circuit resource.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the virtal circuit.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of project to which the VC belongs.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description for the Virtual Circuit resource.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags for the Virtual Circuit resource.`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `Speed of the Virtual Circuit resource.`,
				},
				resource.Attribute{
					Name:        "vrf_id",
					Description: `UUID of the VLAN to associate.`,
				},
				resource.Attribute{
					Name:        "peer_asn",
					Description: `The BGP ASN of the peer. The same ASN may be the used across several VCs, but it cannot be the same as the local_asn of the VRF.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `A subnet from one of the IP blocks associated with the VRF that we will help create an IP reservation for. Can only be either a /30 or /31.`,
				},
				resource.Attribute{
					Name:        "metal_ip",
					Description: `The IP address that’s set as “our” IP that is configured on the rack_local_vlan SVI. Will default to the first usable IP in the subnet.`,
				},
				resource.Attribute{
					Name:        "customer_ip",
					Description: `The IP address set as the customer IP which the CSR switch will peer with. Will default to the other usable IP in the subnet.`,
				},
				resource.Attribute{
					Name:        "md5",
					Description: `The password that can be set for the VRF BGP peer`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_metal_vlan",
			Category:         "Metal",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"metal",
				"vlan",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vlan_id",
					Description: `(Optional) Metal UUID of the VLAN resource to look up.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) UUID of parent project of the VLAN. Use together with the vxlan number and metro or facility.`,
				},
				resource.Attribute{
					Name:        "vxlan",
					Description: `(Optional) vxlan number of the VLAN to look up. Use together with the project_id and metro or facility.`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `(Optional) Facility where the VLAN is deployed.`,
				},
				resource.Attribute{
					Name:        "metro",
					Description: `(Optional) Metro where the VLAN is deployed. ->`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description text of the VLAN resource.`,
				},
				resource.Attribute{
					Name:        "assigned_devices_ids",
					Description: `List of device ID to which this VLAN is assigned.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Description text of the VLAN resource.`,
				},
				resource.Attribute{
					Name:        "assigned_devices_ids",
					Description: `List of device ID to which this VLAN is assigned.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_metal_vrf",
			Category:         "Metal",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"metal",
				"vrf",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vrf_id",
					Description: `(Required) ID of the VRF resource ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `User-supplied name of the VRF, unique to the project`,
				},
				resource.Attribute{
					Name:        "metro",
					Description: `Metro ID or Code where the VRF will be deployed.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Project ID where the VRF will be deployed.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the VRF.`,
				},
				resource.Attribute{
					Name:        "local_asn",
					Description: `The 4-byte ASN set on the VRF.`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `All IPv4 and IPv6 Ranges that will be available to BGP Peers. IPv4 addresses must be /8 or smaller with a minimum size of /29. IPv6 must be /56 or smaller with a minimum size of /64. Ranges must not overlap other ranges within the VRF.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `User-supplied name of the VRF, unique to the project`,
				},
				resource.Attribute{
					Name:        "metro",
					Description: `Metro ID or Code where the VRF will be deployed.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Project ID where the VRF will be deployed.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the VRF.`,
				},
				resource.Attribute{
					Name:        "local_asn",
					Description: `The 4-byte ASN set on the VRF.`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `All IPv4 and IPv6 Ranges that will be available to BGP Peers. IPv4 addresses must be /8 or smaller with a minimum size of /29. IPv6 must be /56 or smaller with a minimum size of /64. Ranges must not overlap other ranges within the VRF.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_network_account",
			Category:         "Network Edge",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"edge",
				"account",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "metro_code",
					Description: `(Required) Account location metro code.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Account name for filtering.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Account status for filtering. Possible values are: ` + "`" + `Active` + "`" + `, ` + "`" + `Processing` + "`" + `, ` + "`" + `Submitted` + "`" + `, ` + "`" + `Staged` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "number",
					Description: `Account unique number.`,
				},
				resource.Attribute{
					Name:        "ucm_id",
					Description: `Account unique identifier.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "number",
					Description: `Account unique number.`,
				},
				resource.Attribute{
					Name:        "ucm_id",
					Description: `Account unique identifier.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_network_device_platform",
			Category:         "Network Edge",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"edge",
				"device",
				"platform",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "device_type",
					Description: `(Required) Device type code`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `(Optional) Device platform flavor that determines number of CPU cores and memory. Supported values are: ` + "`" + `small` + "`" + `, ` + "`" + `medium` + "`" + `, ` + "`" + `large` + "`" + `, ` + "`" + `xlarge` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "core_count",
					Description: `(Optional) Number of CPU cores used to limit platform search results.`,
				},
				resource.Attribute{
					Name:        "packages",
					Description: `(Optional) List of software package codes to limit platform search results.`,
				},
				resource.Attribute{
					Name:        "management_types",
					Description: `(Optional) List of device management types to limit platform search results. Supported values are: ` + "`" + `EQUINIX-CONFIGURED` + "`" + `, ` + "`" + `SELF-CONFIGURED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "license_options",
					Description: `(Optional) List of device licensing options to limit platform search result. Supported values are: ` + "`" + `BYOL` + "`" + ` (for Bring Your Own License), ` + "`" + `Sub` + "`" + ` (for license subscription). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `The amount of memory provided by device platform.`,
				},
				resource.Attribute{
					Name:        "memory_unit",
					Description: `Unit of memory provider by device platform.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "memory",
					Description: `The amount of memory provided by device platform.`,
				},
				resource.Attribute{
					Name:        "memory_unit",
					Description: `Unit of memory provider by device platform.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_network_device_software",
			Category:         "Network Edge",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"edge",
				"device",
				"software",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "device_type",
					Description: `(Required) Code of a device type.`,
				},
				resource.Attribute{
					Name:        "version_regex",
					Description: `(Optional) A regex string to apply on returned versions and filter search results.`,
				},
				resource.Attribute{
					Name:        "stable",
					Description: `(Optional) Boolean value to limit query results to stable versions only.`,
				},
				resource.Attribute{
					Name:        "packages",
					Description: `(Optional) Limits returned versions to those that are supported by given software package codes.`,
				},
				resource.Attribute{
					Name:        "most_recent",
					Description: `(Optional) Boolean value to indicate that most recent version should be used`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Version number.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `Software image name.`,
				},
				resource.Attribute{
					Name:        "date",
					Description: `Version release date.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Version status.`,
				},
				resource.Attribute{
					Name:        "release_notes_link",
					Description: `Link to version release notes.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "version",
					Description: `Version number.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `Software image name.`,
				},
				resource.Attribute{
					Name:        "date",
					Description: `Version release date.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Version status.`,
				},
				resource.Attribute{
					Name:        "release_notes_link",
					Description: `Link to version release notes.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_network_device_type",
			Category:         "Network Edge",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"edge",
				"device",
				"type",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Device type name.`,
				},
				resource.Attribute{
					Name:        "vendor",
					Description: `(Optional) Device type vendor i.e. ` + "`" + `Cisco` + "`" + `, ` + "`" + `Juniper Networks` + "`" + `, ` + "`" + `VERSA Networks` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `(Optional) Device type category. One of: ` + "`" + `Router` + "`" + `, ` + "`" + `Firewall` + "`" + `, ` + "`" + `SDWAN` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "metro_codes",
					Description: `(Optional) List of metro codes where device type has to be available ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `Device type short code, unique identifier of a network device type`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Device type textual description`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "code",
					Description: `Device type short code, unique identifier of a network device type`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Device type textual description`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"equinix_equinix_ecx_l2_sellerprofile":       0,
		"equinix_equinix_ecx_l2_sellerprofiles":      1,
		"equinix_equinix_ecx_port":                   2,
		"equinix_equinix_metal_connection":           3,
		"equinix_equinix_metal_device":               4,
		"equinix_equinix_metal_device_bgp_neighbors": 5,
		"equinix_equinix_metal_facility":             6,
		"equinix_equinix_metal_gateway":              7,
		"equinix_equinix_metal_hardware_reservation": 8,
		"equinix_equinix_metal_ip_block_ranges":      9,
		"equinix_equinix_metal_metro":                10,
		"equinix_equinix_metal_operating_system":     11,
		"equinix_equinix_metal_organization":         12,
		"equinix_equinix_metal_port":                 13,
		"equinix_equinix_metal_precreated_ip_block":  14,
		"equinix_equinix_metal_project":              15,
		"equinix_equinix_metal_project_ssh_key":      16,
		"equinix_equinix_metal_reserved_ip_block":    17,
		"equinix_equinix_metal_spot_market_price":    18,
		"equinix_equinix_metal_spot_market_request":  19,
		"equinix_equinix_metal_virtual_circuit":      20,
		"equinix_equinix_metal_vlan":                 21,
		"equinix_equinix_metal_vrf":                  22,
		"equinix_equinix_network_account":            23,
		"equinix_equinix_network_device_platform":    24,
		"equinix_equinix_network_device_software":    25,
		"equinix_equinix_network_device_type":        26,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
