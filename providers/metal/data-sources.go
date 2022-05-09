package metal

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "metal_connection",
			Category:         "Data Sources",
			ShortDescription: `Retrieve Equinix Fabric Connection`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "connection_id",
					Description: `(Required) ID of the connection resource ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the connection resource`,
				},
				resource.Attribute{
					Name:        "metro",
					Description: `Slug of a metro to which the connection belongs`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `Slug of a facility to which the connection belongs`,
				},
				resource.Attribute{
					Name:        "redundancy",
					Description: `Connection redundancy, reduntant or primary`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Connection type, dedicated or shared`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of project to which the connection belongs`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `Connection speed, one of 50Mbps, 200Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the connection resource`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Mode for connections in IBX facilities with the dedicated type - standard or tunnel`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `String list of tags`,
				},
				resource.Attribute{
					Name:        "vlans",
					Description: `Attached VLANs. Only available in shared connection. One vlan for Primary/Single connection and two vlans for Redundant connection`,
				},
				resource.Attribute{
					Name:        "service_token_type",
					Description: `Type of service token, a_side or z_side. One available in shared connection.`,
				},
				resource.Attribute{
					Name:        "organization_id",
					Description: `ID of the organization where the connection is scoped to`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the connection resource`,
				},
				resource.Attribute{
					Name:        "service_tokens",
					Description: `List of connection service tokens with attributes`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `UUID of the service token`,
				},
				resource.Attribute{
					Name:        "expires_at",
					Description: `Expiration date of the service token`,
				},
				resource.Attribute{
					Name:        "max_allowed_speed",
					Description: `Maximum allowed speed for the service token, string like in the ` + "`" + `speed` + "`" + ` attribute`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Token type, ` + "`" + `a_side` + "`" + ` or ` + "`" + `z_side` + "`" + ``,
				},
				resource.Attribute{
					Name:        "role",
					Description: `Token role, ` + "`" + `primary` + "`" + ` or ` + "`" + `secondary` + "`" + ``,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `List of connection ports - primary (` + "`" + `ports[0]` + "`" + `) and secondary (` + "`" + `ports[1]` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Port name`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Port UUID`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `Port role - primary or secondary`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `Port speed in bits per second`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Port status`,
				},
				resource.Attribute{
					Name:        "link_status",
					Description: `Port link status`,
				},
				resource.Attribute{
					Name:        "virtual_circuit_ids",
					Description: `List of IDs of virtual cicruits attached to this port`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Deprecated) Fabric Token from the [Equinix Fabric Portal](https://ecxfabric.equinix.com/dashboard)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name of the connection resource`,
				},
				resource.Attribute{
					Name:        "metro",
					Description: `Slug of a metro to which the connection belongs`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `Slug of a facility to which the connection belongs`,
				},
				resource.Attribute{
					Name:        "redundancy",
					Description: `Connection redundancy, reduntant or primary`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Connection type, dedicated or shared`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of project to which the connection belongs`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `Connection speed, one of 50Mbps, 200Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the connection resource`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Mode for connections in IBX facilities with the dedicated type - standard or tunnel`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `String list of tags`,
				},
				resource.Attribute{
					Name:        "vlans",
					Description: `Attached VLANs. Only available in shared connection. One vlan for Primary/Single connection and two vlans for Redundant connection`,
				},
				resource.Attribute{
					Name:        "service_token_type",
					Description: `Type of service token, a_side or z_side. One available in shared connection.`,
				},
				resource.Attribute{
					Name:        "organization_id",
					Description: `ID of the organization where the connection is scoped to`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the connection resource`,
				},
				resource.Attribute{
					Name:        "service_tokens",
					Description: `List of connection service tokens with attributes`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `UUID of the service token`,
				},
				resource.Attribute{
					Name:        "expires_at",
					Description: `Expiration date of the service token`,
				},
				resource.Attribute{
					Name:        "max_allowed_speed",
					Description: `Maximum allowed speed for the service token, string like in the ` + "`" + `speed` + "`" + ` attribute`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Token type, ` + "`" + `a_side` + "`" + ` or ` + "`" + `z_side` + "`" + ``,
				},
				resource.Attribute{
					Name:        "role",
					Description: `Token role, ` + "`" + `primary` + "`" + ` or ` + "`" + `secondary` + "`" + ``,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `List of connection ports - primary (` + "`" + `ports[0]` + "`" + `) and secondary (` + "`" + `ports[1]` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Port name`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Port UUID`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `Port role - primary or secondary`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `Port speed in bits per second`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Port status`,
				},
				resource.Attribute{
					Name:        "link_status",
					Description: `Port link status`,
				},
				resource.Attribute{
					Name:        "virtual_circuit_ids",
					Description: `List of IDs of virtual cicruits attached to this port`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Deprecated) Fabric Token from the [Equinix Fabric Portal](https://ecxfabric.equinix.com/dashboard)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "metal_device",
			Category:         "Data Sources",
			ShortDescription: `Provides an Equinix Metal device datasource. This can be used to read existing devices.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "hostname",
					Description: `The device name`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The id of the project in which the devices exists`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `Device ID User can lookup devices either by ` + "`" + `device_id` + "`" + ` or ` + "`" + `project_id` + "`" + ` and ` + "`" + `hostname` + "`" + `. ## Attributes Reference The following attributes are exported:`,
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
					Name:        "metro",
					Description: `The metro where the device is deployed`,
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
					Name:        "metro",
					Description: `The metro where the device is deployed`,
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
			Type:             "metal_device_bgp_neighbors",
			Category:         "Data Sources",
			ShortDescription: `Provides a datasource for listing BGP neighbors of an Equinix Metal device`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "device_id",
					Description: `UUID of BGP-enabled device whose neighbors to list ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "bgp_neighbors",
					Description: `array of BGP neighbor records with attributes:`,
				},
				resource.Attribute{
					Name:        "address_family",
					Description: `IP address version, 4 or 6`,
				},
				resource.Attribute{
					Name:        "customer_as",
					Description: `Local autonomous system number`,
				},
				resource.Attribute{
					Name:        "customer_ip",
					Description: `Local used peer IP address`,
				},
				resource.Attribute{
					Name:        "md5_enabled",
					Description: `Whether BGP session is password enabled`,
				},
				resource.Attribute{
					Name:        "md5_password",
					Description: `BGP session password in plaintext (not a checksum)`,
				},
				resource.Attribute{
					Name:        "multihop",
					Description: `Whether the neighbor is in EBGP multihop session`,
				},
				resource.Attribute{
					Name:        "peer_as",
					Description: `Peer AS number (different than customer_as for EBGP)`,
				},
				resource.Attribute{
					Name:        "peer_ips",
					Description: `Array of IP addresses of this neighbor's peers`,
				},
				resource.Attribute{
					Name:        "routes_in",
					Description: `Array of incoming routes. Each route has attributes:`,
				},
				resource.Attribute{
					Name:        "route",
					Description: `CIDR expression of route (IP/mask)`,
				},
				resource.Attribute{
					Name:        "exact",
					Description: `(bool) Whether the route is exact`,
				},
				resource.Attribute{
					Name:        "routes_out",
					Description: `Array of outgoing routes in the same format`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "bgp_neighbors",
					Description: `array of BGP neighbor records with attributes:`,
				},
				resource.Attribute{
					Name:        "address_family",
					Description: `IP address version, 4 or 6`,
				},
				resource.Attribute{
					Name:        "customer_as",
					Description: `Local autonomous system number`,
				},
				resource.Attribute{
					Name:        "customer_ip",
					Description: `Local used peer IP address`,
				},
				resource.Attribute{
					Name:        "md5_enabled",
					Description: `Whether BGP session is password enabled`,
				},
				resource.Attribute{
					Name:        "md5_password",
					Description: `BGP session password in plaintext (not a checksum)`,
				},
				resource.Attribute{
					Name:        "multihop",
					Description: `Whether the neighbor is in EBGP multihop session`,
				},
				resource.Attribute{
					Name:        "peer_as",
					Description: `Peer AS number (different than customer_as for EBGP)`,
				},
				resource.Attribute{
					Name:        "peer_ips",
					Description: `Array of IP addresses of this neighbor's peers`,
				},
				resource.Attribute{
					Name:        "routes_in",
					Description: `Array of incoming routes. Each route has attributes:`,
				},
				resource.Attribute{
					Name:        "route",
					Description: `CIDR expression of route (IP/mask)`,
				},
				resource.Attribute{
					Name:        "exact",
					Description: `(bool) Whether the route is exact`,
				},
				resource.Attribute{
					Name:        "routes_out",
					Description: `Array of outgoing routes in the same format`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "metal_facility",
			Category:         "Data Sources",
			ShortDescription: `Provides an Equinix Metal facility datasource. This can be used to read facilities.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "code",
					Description: `The facility code`,
				},
				resource.Attribute{
					Name:        "features_required",
					Description: `Set of feature strings that the facility must have Facilities can be looked up by ` + "`" + `code` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the facility`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the facility`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `The features of the facility`,
				},
				resource.Attribute{
					Name:        "metro",
					Description: `The metro code the facility is part of`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `(Optional) Ensure that queried facility has capacity for specified number of given plans - ` + "`" + `plan` + "`" + ` - device plan to check - ` + "`" + `quantity` + "`" + ` - number of device to check`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the facility`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the facility`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `The features of the facility`,
				},
				resource.Attribute{
					Name:        "metro",
					Description: `The metro code the facility is part of`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `(Optional) Ensure that queried facility has capacity for specified number of given plans - ` + "`" + `plan` + "`" + ` - device plan to check - ` + "`" + `quantity` + "`" + ` - number of device to check`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "metal_gateway",
			Category:         "Data Sources",
			ShortDescription: `Retrieve Equinix Metal Gateways`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "gateway_id",
					Description: `(Required) UUID of the metal gateway resource to retrieve ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `UUID of the project where the gateway is scoped to`,
				},
				resource.Attribute{
					Name:        "vlan_id",
					Description: `UUID of the VLAN where the gateway is scoped to`,
				},
				resource.Attribute{
					Name:        "ip_reservation_id",
					Description: `UUID of IP reservation block bound to the gateway`,
				},
				resource.Attribute{
					Name:        "private_ipv4_subnet_size",
					Description: `Size of the private IPv4 subnet bound to this metal gateway, one of (8, 16, 32, 64, 128)` + "`" + ``,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Status of the gateway resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `UUID of the project where the gateway is scoped to`,
				},
				resource.Attribute{
					Name:        "vlan_id",
					Description: `UUID of the VLAN where the gateway is scoped to`,
				},
				resource.Attribute{
					Name:        "ip_reservation_id",
					Description: `UUID of IP reservation block bound to the gateway`,
				},
				resource.Attribute{
					Name:        "private_ipv4_subnet_size",
					Description: `Size of the private IPv4 subnet bound to this metal gateway, one of (8, 16, 32, 64, 128)` + "`" + ``,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Status of the gateway resource`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "metal_hardware_reservation",
			Category:         "Data Sources",
			ShortDescription: `Retrieve Equinix Metal Hardware Reservation`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the hardware reservation`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `UUID of device occupying the reservation ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the hardware reservation to look up`,
				},
				resource.Attribute{
					Name:        "short_id",
					Description: `Reservation short ID`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `UUID of project this reservation is scoped to`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `UUID of device occupying the reservation`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `Plan type for the reservation`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `Plan type for the reservation`,
				},
				resource.Attribute{
					Name:        "provisionable",
					Description: `Flag indicating whether the reserved server is provisionable or not. Spare devices can't be provisioned unless they are activated first`,
				},
				resource.Attribute{
					Name:        "spare",
					Description: `Flag indicating whether the Hardware Reservation is a spare. Spare Hardware Reservations are used when a Hardware Reservations requires service from Metal Equinix`,
				},
				resource.Attribute{
					Name:        "switch_uuid",
					Description: `Switch short ID, can be used to determine if two devices are connected to the same switch`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the hardware reservation to look up`,
				},
				resource.Attribute{
					Name:        "short_id",
					Description: `Reservation short ID`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `UUID of project this reservation is scoped to`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `UUID of device occupying the reservation`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `Plan type for the reservation`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `Plan type for the reservation`,
				},
				resource.Attribute{
					Name:        "provisionable",
					Description: `Flag indicating whether the reserved server is provisionable or not. Spare devices can't be provisioned unless they are activated first`,
				},
				resource.Attribute{
					Name:        "spare",
					Description: `Flag indicating whether the Hardware Reservation is a spare. Spare Hardware Reservations are used when a Hardware Reservations requires service from Metal Equinix`,
				},
				resource.Attribute{
					Name:        "switch_uuid",
					Description: `Switch short ID, can be used to determine if two devices are connected to the same switch`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "metal_ip_block_ranges",
			Category:         "Data Sources",
			ShortDescription: `List IP address ranges allocated to a project`,
			Description:      ``,
			Keywords:         []string{},
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
					Description: `(Optional) Metro code filtering the IP blocks. Global IPv4 blcoks will be listed anyway. If you omit this and facility, all the block from the project will be listed. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "global_ipv4",
					Description: `list of CIDR expressions for Global IPv4 blocks in the project`,
				},
				resource.Attribute{
					Name:        "public_ipv4",
					Description: `list of CIDR expressions for Public IPv4 blocks in the project`,
				},
				resource.Attribute{
					Name:        "private_ipv4",
					Description: `list of CIDR expressions for Private IPv4 blocks in the project`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `list of CIDR expressions for IPv6 blocks in the project`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "global_ipv4",
					Description: `list of CIDR expressions for Global IPv4 blocks in the project`,
				},
				resource.Attribute{
					Name:        "public_ipv4",
					Description: `list of CIDR expressions for Public IPv4 blocks in the project`,
				},
				resource.Attribute{
					Name:        "private_ipv4",
					Description: `list of CIDR expressions for Private IPv4 blocks in the project`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `list of CIDR expressions for IPv6 blocks in the project`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "metal_metro",
			Category:         "Data Sources",
			ShortDescription: `Provides an Equinix Metal metro datasource. This can be used to read metros.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "code",
					Description: `The metro code Metros can be looked up by ` + "`" + `code` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the metro`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `The code of the metro`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `The country of the metro`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the metro`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `(Optional) Ensure that queried metro has capacity for specified number of given plans - ` + "`" + `plan` + "`" + ` - device plan to check - ` + "`" + `quantity` + "`" + ` - number of device to check`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the metro`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `The code of the metro`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `The country of the metro`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the metro`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `(Optional) Ensure that queried metro has capacity for specified number of given plans - ` + "`" + `plan` + "`" + ` - device plan to check - ` + "`" + `quantity` + "`" + ` - number of device to check`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "metal_operating_system",
			Category:         "Data Sources",
			ShortDescription: `Get an Equinix Metal operating system image`,
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
			Type:             "metal_organization",
			Category:         "Data Sources",
			ShortDescription: `Provides an Equinix Metal Organization datasource. This can be used to read existing Organizations.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The organization name`,
				},
				resource.Attribute{
					Name:        "organization_id",
					Description: `The UUID of the organization resource Exactly one of ` + "`" + `name` + "`" + ` or ` + "`" + `organization_id` + "`" + ` must be given. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "project_ids",
					Description: `UUIDs of project resources which belong to this organization`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description string`,
				},
				resource.Attribute{
					Name:        "website",
					Description: `Website link`,
				},
				resource.Attribute{
					Name:        "twitter",
					Description: `Twitter handle`,
				},
				resource.Attribute{
					Name:        "logo",
					Description: `Logo URL`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "project_ids",
					Description: `UUIDs of project resources which belong to this organization`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description string`,
				},
				resource.Attribute{
					Name:        "website",
					Description: `Website link`,
				},
				resource.Attribute{
					Name:        "twitter",
					Description: `Twitter handle`,
				},
				resource.Attribute{
					Name:        "logo",
					Description: `Logo URL`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "metal_port",
			Category:         "Data Sources",
			ShortDescription: `Fetch device ports`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) ID of the port to read, conflicts with device_id.`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Whether to look for public or private block. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `One of layer2-bonded, layer2-individual, layer3, hybrid, hybrid-bonded`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type is either "NetworkBondPort" for bond ports or "NetworkPort" for bondable ethernet ports`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `MAC address of the port`,
				},
				resource.Attribute{
					Name:        "bond_id",
					Description: `UUID of the bond port"`,
				},
				resource.Attribute{
					Name:        "bond_name",
					Description: `Name of the bond port`,
				},
				resource.Attribute{
					Name:        "bonded",
					Description: `Flag indicating whether the port is bonded`,
				},
				resource.Attribute{
					Name:        "disbond_supported",
					Description: `Flag indicating whether the port can be removed from a bond`,
				},
				resource.Attribute{
					Name:        "native_vlan_id",
					Description: `UUID of native VLAN of the port`,
				},
				resource.Attribute{
					Name:        "vlan_ids",
					Description: `UUIDs of attached VLANs`,
				},
				resource.Attribute{
					Name:        "vxlan_ids",
					Description: `VXLAN ids of attached VLANs`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "network_type",
					Description: `One of layer2-bonded, layer2-individual, layer3, hybrid, hybrid-bonded`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type is either "NetworkBondPort" for bond ports or "NetworkPort" for bondable ethernet ports`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `MAC address of the port`,
				},
				resource.Attribute{
					Name:        "bond_id",
					Description: `UUID of the bond port"`,
				},
				resource.Attribute{
					Name:        "bond_name",
					Description: `Name of the bond port`,
				},
				resource.Attribute{
					Name:        "bonded",
					Description: `Flag indicating whether the port is bonded`,
				},
				resource.Attribute{
					Name:        "disbond_supported",
					Description: `Flag indicating whether the port can be removed from a bond`,
				},
				resource.Attribute{
					Name:        "native_vlan_id",
					Description: `UUID of native VLAN of the port`,
				},
				resource.Attribute{
					Name:        "vlan_ids",
					Description: `UUIDs of attached VLANs`,
				},
				resource.Attribute{
					Name:        "vxlan_ids",
					Description: `VXLAN ids of attached VLANs`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "metal_precreated_ip_block",
			Category:         "Data Sources",
			ShortDescription: `Load automatically created IP blocks from your Equinix Metal project`,
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
					Description: `(Optional) Facility of the searched block. (for non-global blocks).`,
				},
				resource.Attribute{
					Name:        "metro",
					Description: `(Optional) Metro of the searched block (for non-global blocks). ## Attributes Reference`,
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
			Type:             "metal_project",
			Category:         "Data Sources",
			ShortDescription: `Provides an Equinix Metal Project datasource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name which is used to look up the project`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The UUID by which to look up the project ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "payment_method_id",
					Description: `The UUID of payment method for this project`,
				},
				resource.Attribute{
					Name:        "organization_id",
					Description: `The UUID of this project's parent organization`,
				},
				resource.Attribute{
					Name:        "backend_transfer",
					Description: `Whether Backend Transfer is enabled for this project`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The timestamp for when the project was created`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `The timestamp for the last time the project was updated`,
				},
				resource.Attribute{
					Name:        "user_ids",
					Description: `List of UUIDs of user accounts which belong to this project`,
				},
				resource.Attribute{
					Name:        "bgp_config",
					Description: `Optional BGP settings. Refer to [Equinix Metal guide for BGP](https://metal.equinix.com/developers/docs/networking/local-global-bgp/). The ` + "`" + `bgp_config` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "asn",
					Description: `Autonomous System Number for local BGP deployment`,
				},
				resource.Attribute{
					Name:        "md5",
					Description: `Password for BGP session in plaintext (not a checksum)`,
				},
				resource.Attribute{
					Name:        "deployment_type",
					Description: `` + "`" + `private` + "`" + ` or ` + "`" + `public` + "`" + `, the ` + "`" + `private` + "`" + ` is likely to be usable immediately, the ` + "`" + `public` + "`" + ` will need to be review by Equinix Metal engineers`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `status of BGP configuration in the project`,
				},
				resource.Attribute{
					Name:        "max_prefix",
					Description: `The maximum number of route filters allowed per server`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "payment_method_id",
					Description: `The UUID of payment method for this project`,
				},
				resource.Attribute{
					Name:        "organization_id",
					Description: `The UUID of this project's parent organization`,
				},
				resource.Attribute{
					Name:        "backend_transfer",
					Description: `Whether Backend Transfer is enabled for this project`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The timestamp for when the project was created`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `The timestamp for the last time the project was updated`,
				},
				resource.Attribute{
					Name:        "user_ids",
					Description: `List of UUIDs of user accounts which belong to this project`,
				},
				resource.Attribute{
					Name:        "bgp_config",
					Description: `Optional BGP settings. Refer to [Equinix Metal guide for BGP](https://metal.equinix.com/developers/docs/networking/local-global-bgp/). The ` + "`" + `bgp_config` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "asn",
					Description: `Autonomous System Number for local BGP deployment`,
				},
				resource.Attribute{
					Name:        "md5",
					Description: `Password for BGP session in plaintext (not a checksum)`,
				},
				resource.Attribute{
					Name:        "deployment_type",
					Description: `` + "`" + `private` + "`" + ` or ` + "`" + `public` + "`" + `, the ` + "`" + `private` + "`" + ` is likely to be usable immediately, the ` + "`" + `public` + "`" + ` will need to be review by Equinix Metal engineers`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `status of BGP configuration in the project`,
				},
				resource.Attribute{
					Name:        "max_prefix",
					Description: `The maximum number of route filters allowed per server`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "metal_project_ssh_key",
			Category:         "Data Sources",
			ShortDescription: `Provides an Equinix Metal Project SSH Key datasource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "search",
					Description: `(Optional) The name, fingerprint, or public_key of the SSH Key to search for in the Equinix Metal project`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The id of the SSH Key to search for in the Equinix Metal project`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The Equinix Metal project id of the Equinix Metal SSH Key One of either ` + "`" + `search` + "`" + ` or ` + "`" + `id` + "`" + ` must be provided along with ` + "`" + `project_id` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the key`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the SSH key`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `The text of the public key`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The ID of parent project`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `The ID of parent project (same as project_id)`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `The fingerprint of the SSH key`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The timestamp for when the SSH key was created`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `The timestamp for the last time the SSH key was updated`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the key`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the SSH key`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `The text of the public key`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The ID of parent project`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `The ID of parent project (same as project_id)`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `The fingerprint of the SSH key`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The timestamp for when the SSH key was created`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `The timestamp for the last time the SSH key was updated`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "metal_reserved_ip_block",
			Category:         "Data Sources",
			ShortDescription: `Look up an IP address block`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) UUID of the IP address block to look up`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) UUID of the project where the searched block should be`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Required) Block containing this IP address will be returned ## Attributes Reference This datasource exposes the same attributes as the [metal_reserved_ip_block resource](../resources/reserved_ip_block.md).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "metal_spot_market_price",
			Category:         "Data Sources",
			ShortDescription: `Get an Equinix Metal Spot Market Price`,
			Description:      ``,
			Keywords:         []string{},
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
					Description: `(Optional) Name of the metro. ## Attributes Reference`,
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
			Type:             "metal_spot_market_request",
			Category:         "Data Sources",
			ShortDescription: `Provides a datasource for existing Spot Market Requests in the Equinix Metal host.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "request_id",
					Description: `(Required) The id of the Spot Market Request ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "device_ids",
					Description: `List of IDs of devices spawned by the referenced Spot Market Request`,
				},
				resource.Attribute{
					Name:        "devices_min",
					Description: `Miniumum number devices to be created`,
				},
				resource.Attribute{
					Name:        "devices_max",
					Description: `Maximum number devices to be created`,
				},
				resource.Attribute{
					Name:        "max_bid_price",
					Description: `Maximum price user is willing to pay per hour per device`,
				},
				resource.Attribute{
					Name:        "facilities",
					Description: `Facility IDs where devices should be created`,
				},
				resource.Attribute{
					Name:        "metro",
					Description: `Metro where devices should be created.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Project ID`,
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
					Description: `List of IDs of devices spawned by the referenced Spot Market Request`,
				},
				resource.Attribute{
					Name:        "devices_min",
					Description: `Miniumum number devices to be created`,
				},
				resource.Attribute{
					Name:        "devices_max",
					Description: `Maximum number devices to be created`,
				},
				resource.Attribute{
					Name:        "max_bid_price",
					Description: `Maximum price user is willing to pay per hour per device`,
				},
				resource.Attribute{
					Name:        "facilities",
					Description: `Facility IDs where devices should be created`,
				},
				resource.Attribute{
					Name:        "metro",
					Description: `Metro where devices should be created.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Project ID`,
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
			Type:             "metal_virtual_circuit",
			Category:         "Data Sources",
			ShortDescription: `Retrieve Equinix Fabric Virtual Circuit`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "virtual_circuit_id",
					Description: `(Required) ID of the virtual circuit resource ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the virtual circuit resource`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the virtal circuit`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of project to which the VC belongs`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description for the Virtual Circuit resource`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags for the Virtual Circuit resource`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `Speed of the Virtual Circuit resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name of the virtual circuit resource`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the virtal circuit`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of project to which the VC belongs`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description for the Virtual Circuit resource`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags for the Virtual Circuit resource`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `Speed of the Virtual Circuit resource`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "metal_vlan",
			Category:         "Data Sources",
			ShortDescription: `Provides an Equinix Metal Virtual Network datasource. This can be used to read the attributes of existing VLANs.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vlan_id",
					Description: `Metal UUID of the VLAN resource to look up`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `UUID of parent project of the VLAN. Use together with the vxlan number and metro or facility`,
				},
				resource.Attribute{
					Name:        "vxlan",
					Description: `vxlan number of the VLAN to look up. Use together with the project_id and metro or facility`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `Facility where the VLAN is deployed`,
				},
				resource.Attribute{
					Name:        "metro",
					Description: `Metro where the VLAN is deployed ## Attributes Reference The following attributes are exported, in addition to any unspecified arguments.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description text of the VLAN resource`,
				},
				resource.Attribute{
					Name:        "assigned_devices_ids",
					Description: `List of device ID to which this VLAN is assigned`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Description text of the VLAN resource`,
				},
				resource.Attribute{
					Name:        "assigned_devices_ids",
					Description: `List of device ID to which this VLAN is assigned`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "metal_volume",
			Category:         "Data Sources",
			ShortDescription: `(Removed) Provides an Equinix Metal Block Storage Volume Datasource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"metal_connection":           0,
		"metal_device":               1,
		"metal_device_bgp_neighbors": 2,
		"metal_facility":             3,
		"metal_gateway":              4,
		"metal_hardware_reservation": 5,
		"metal_ip_block_ranges":      6,
		"metal_metro":                7,
		"metal_operating_system":     8,
		"metal_organization":         9,
		"metal_port":                 10,
		"metal_precreated_ip_block":  11,
		"metal_project":              12,
		"metal_project_ssh_key":      13,
		"metal_reserved_ip_block":    14,
		"metal_spot_market_price":    15,
		"metal_spot_market_request":  16,
		"metal_virtual_circuit":      17,
		"metal_vlan":                 18,
		"metal_volume":               19,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
