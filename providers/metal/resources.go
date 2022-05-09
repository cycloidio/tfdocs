package metal

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "metal_bgp_session",
			Category:         "Resources",
			ShortDescription: `BGP session in Equinix Metal Host`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "device_id",
					Description: `(Required) ID of device`,
				},
				resource.Attribute{
					Name:        "address_family",
					Description: `(Required) ` + "`" + `ipv4` + "`" + ` or ` + "`" + `ipv6` + "`" + ``,
				},
				resource.Attribute{
					Name:        "default_route",
					Description: `(Optional) Boolean flag to set the default route policy. False by default. ## Attributes Reference The following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "metal_connection",
			Category:         "Resources",
			ShortDescription: `Request/Create Equinix Fabric Connection`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the connection resource`,
				},
				resource.Attribute{
					Name:        "metro",
					Description: `(Optional) Metro where the connection will be created`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `(Optional) Facility where the connection will be created`,
				},
				resource.Attribute{
					Name:        "redundancy",
					Description: `(Required) Connection redundancy - redundant or primary`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Connection type - dedicated or shared`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) ID of the project where the connection is scoped to, must be set for shared connection`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `(Required) Connection speed - one of 50Mbps, 200Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for the connection resource`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) Mode for connections in IBX facilities with the dedicated type - standard or tunnel. Default is standard`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) String list of tags`,
				},
				resource.Attribute{
					Name:        "vlans",
					Description: `(Optional) Only used with shared connection. Vlans to attach. Pass one vlan for Primary/Single connection and two vlans for Redundant connection`,
				},
				resource.Attribute{
					Name:        "service_token_type",
					Description: `(Optional) Only used with shared connection. Type of service token to use for the connection, a_side or z_side ## Attributes Reference`,
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
					Name:        "ports",
					Description: `List of connection ports - primary (` + "`" + `ports[0]` + "`" + `) and secondary (` + "`" + `ports[1]` + "`" + `). Schema of port is described in documentation of the [metal_connection datasource](../data-sources/connection.md).`,
				},
				resource.Attribute{
					Name:        "service_tokens",
					Description: `List of connection service tokens with attributes. Scehma of service_token is described in documentation of the [metal_connection datasource](../data-sources/connection.md).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "organization_id",
					Description: `ID of the organization where the connection is scoped to`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the connection resource`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `List of connection ports - primary (` + "`" + `ports[0]` + "`" + `) and secondary (` + "`" + `ports[1]` + "`" + `). Schema of port is described in documentation of the [metal_connection datasource](../data-sources/connection.md).`,
				},
				resource.Attribute{
					Name:        "service_tokens",
					Description: `List of connection service tokens with attributes. Scehma of service_token is described in documentation of the [metal_connection datasource](../data-sources/connection.md).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "metal_device",
			Category:         "Resources",
			ShortDescription: `Provides an Equinix Metal device resource. This can be used to create, modify, and delete devices.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "always_pxe",
					Description: `(Optional) If true, a device with OS ` + "`" + `custom_ipxe` + "`" + ` will continue to boot via iPXE on reboots.`,
				},
				resource.Attribute{
					Name:        "billing_cycle",
					Description: `(Optional) monthly or hourly`,
				},
				resource.Attribute{
					Name:        "custom_data",
					Description: `(Optional) A string of the desired Custom Data for the device.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The device description.`,
				},
				resource.Attribute{
					Name:        "facilities",
					Description: `List of facility codes with deployment preferences. Equinix Metal API will go through the list and will deploy your device to first facility with free capacity. List items must be facility codes or ` + "`" + `any` + "`" + ` (a wildcard). To find the facility code, visit [Facilities API docs](https://metal.equinix.com/developers/api/facilities/), set your API auth token in the top of the page and see JSON from the API response. Conflicts with ` + "`" + `metro` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "force_detach_volumes",
					Description: `(Optional) Delete device even if it has volumes attached. Only applies for destroy action.`,
				},
				resource.Attribute{
					Name:        "hardware_reservation_id",
					Description: `(Optional) The UUID of the hardware reservation where you want this device deployed, or ` + "`" + `next-available` + "`" + ` if you want to pick your next available reservation automatically. Changing this from a reservation UUID to ` + "`" + `next-available` + "`" + ` will re-create the device in another reservation. Please be careful when using hardware reservation UUID and ` + "`" + `next-available` + "`" + ` together for the same pool of reservations. It might happen that the reservation which Equinix Metal API will pick as ` + "`" + `next-available` + "`" + ` is the reservation which you refer with UUID in another metal_device resource. If that happens, and the metal_device with the UUID is created later, resource creation will fail because the reservation is already in use (by the resource created with ` + "`" + `next-available` + "`" + `). To workaround this, have the ` + "`" + `next-available` + "`" + ` resource [explicitly depend_on](https://learn.hashicorp.com/terraform/getting-started/dependencies.html#implicit-and-explicit-dependencies) the resource with hardware reservation UUID, so that the latter is created first. For more details, see [issue #176](https://github.com/packethost/terraform-provider-packet/issues/176).`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional) The device hostname used in deployments taking advantage of Layer3 DHCP or metadata service configuration.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) A list of IP address types for the device (structure is documented below).`,
				},
				resource.Attribute{
					Name:        "ipxe_script_url",
					Description: `(Optional) URL pointing to a hosted iPXE script. More information is in the [Custom iPXE](https://metal.equinix.com/developers/docs/servers/custom-ipxe/) doc.`,
				},
				resource.Attribute{
					Name:        "metro",
					Description: `Metro area for the new device. Conflicts with ` + "`" + `facilities` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "operating_system",
					Description: `(Required) The operating system slug. To find the slug, or visit [Operating Systems API docs](https://metal.equinix.com/developers/api/operatingsystems), set your API auth token in the top of the page and see JSON from the API response.`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `(Required) The device plan slug. To find the plan slug, visit [Device plans API docs](https://metal.equinix.com/developers/api/plans), set your auth token in the top of the page and see JSON from the API response.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The ID of the project in which to create the device`,
				},
				resource.Attribute{
					Name:        "project_ssh_key_ids",
					Description: `Array of IDs of the project SSH keys which should be added to the device. If you omit this, SSH keys of all the members of the parent project will be added to the device. If you specify this array, only the listed project SSH keys will be added. Project SSH keys can be created with the [metal_project_ssh_key](project_ssh_key.md) resource.`,
				},
				resource.Attribute{
					Name:        "reinstall",
					Description: `(Optional) Whether the device should be reinstalled instead of destroyed when modifying user_data, custom_data, or operating system.`,
				},
				resource.Attribute{
					Name:        "storage",
					Description: `(Optional) JSON for custom partitioning. Only usable on reserved hardware. More information in in the [Custom Partitioning and RAID](https://metal.equinix.com/developers/docs/servers/custom-partitioning-raid/) doc. Please note that the disks.partitions.size attribute must be a string, not an integer. It can be a number string, or size notation string, e.g. "4G" or "8M" (for gigabytes and megabytes).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags attached to the device`,
				},
				resource.Attribute{
					Name:        "termination_time",
					Description: `Timestamp for device termination. For example ` + "`" + `2021-09-03T16:32:00+03:00` + "`" + `. If you don't supply timezone info, timestamp is assumed to be in UTC.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional) A string of the desired User Data for the device.`,
				},
				resource.Attribute{
					Name:        "wait_for_reservation_deprovision",
					Description: `(Optional) Only used for devices in reserved hardware. If set, the deletion of this device will block until the hardware reservation is marked provisionable (about 4 minutes in August 2019). The ` + "`" + `ip_address` + "`" + ` block has 3 fields:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `One of [` + "`" + `private_ipv4` + "`" + `, ` + "`" + `public_ipv4` + "`" + `, ` + "`" + `public_ipv6` + "`" + `]`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `CIDR suffix for IP address block to be assigned, i.e. amount of addresses.`,
				},
				resource.Attribute{
					Name:        "reservation_ids",
					Description: `List of UUIDs of [IP block reservations](reserved_ip_block.md) from which the public IPv4 address should be taken. You can supply one ` + "`" + `ip_address` + "`" + ` block per IP address type. If you use the ` + "`" + `ip_address` + "`" + ` you must always pass a block for ` + "`" + `private_ipv4` + "`" + `. To learn more about using the reserved IP addresses for new devices, see the examples in the [metal_reserved_ip_block](reserved_ip_block.md) documentation. The ` + "`" + `reinstall` + "`" + ` block has 3 fields:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Whether the provider should favour reinstall over destroy and create. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "preserve_data",
					Description: `(Optional) Whether the non-OS disks should be kept or wiped during reinstall. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "deprovision_fast",
					Description: `(Optional) Whether the OS disk should be filled with ` + "`" + `00h` + "`" + ` bytes before reinstall. Defaults to ` + "`" + `false` + "`" + `. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/configuration/resources#operation-timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 20 mins) Used when creating the Device. This includes the time to provision the OS.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 20 mins) Used when updating the Device. This includes the time needed to reprovision instances when ` + "`" + `reinstall` + "`" + ` arguments are used.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 20 mins) Used when deleting the Device. This includes the time to deprovision a hardware reservation when ` + "`" + `wait_for_reservation_deprovision` + "`" + ` is enabled. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "access_private_ipv4",
					Description: `The ipv4 private IP assigned to the device`,
				},
				resource.Attribute{
					Name:        "access_public_ipv4",
					Description: `The ipv4 maintenance IP assigned to the device`,
				},
				resource.Attribute{
					Name:        "access_public_ipv6",
					Description: `The ipv6 maintenance IP assigned to the device`,
				},
				resource.Attribute{
					Name:        "billing_cycle",
					Description: `The billing cycle of the device (monthly or hourly)`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The timestamp for when the device was created`,
				},
				resource.Attribute{
					Name:        "deployed_facility",
					Description: `The facility where the device is deployed.`,
				},
				resource.Attribute{
					Name:        "deployed_hardware_reservation_id",
					Description: `ID of hardware reservation where this device was deployed. It is useful when using the ` + "`" + `next-available` + "`" + ` hardware reservation.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description string for the device`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `The hostname of the device`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the device`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `Whether the device is locked`,
				},
				resource.Attribute{
					Name:        "metro",
					Description: `The metro area where the device is deployed`,
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
					Description: `bit length of the network mask of the address`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `address of router`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `whether the address is routable from the Internet`,
				},
				resource.Attribute{
					Name:        "family",
					Description: `IP version - "4" or "6"`,
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
					Name:        "project_id",
					Description: `The ID of the project the device belongs to`,
				},
				resource.Attribute{
					Name:        "root_password",
					Description: `Root password to the server (disabled after 24 hours)`,
				},
				resource.Attribute{
					Name:        "ssh_key_ids",
					Description: `List of IDs of SSH keys deployed in the device, can be both user and project SSH keys`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The status of the device`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags attached to the device`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `The timestamp for the last time the device was updated ## Import This resource can be imported using an existing device ID: ` + "`" + `` + "`" + `` + "`" + `sh terraform import metal_device {existing_device_id} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "access_private_ipv4",
					Description: `The ipv4 private IP assigned to the device`,
				},
				resource.Attribute{
					Name:        "access_public_ipv4",
					Description: `The ipv4 maintenance IP assigned to the device`,
				},
				resource.Attribute{
					Name:        "access_public_ipv6",
					Description: `The ipv6 maintenance IP assigned to the device`,
				},
				resource.Attribute{
					Name:        "billing_cycle",
					Description: `The billing cycle of the device (monthly or hourly)`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The timestamp for when the device was created`,
				},
				resource.Attribute{
					Name:        "deployed_facility",
					Description: `The facility where the device is deployed.`,
				},
				resource.Attribute{
					Name:        "deployed_hardware_reservation_id",
					Description: `ID of hardware reservation where this device was deployed. It is useful when using the ` + "`" + `next-available` + "`" + ` hardware reservation.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description string for the device`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `The hostname of the device`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the device`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `Whether the device is locked`,
				},
				resource.Attribute{
					Name:        "metro",
					Description: `The metro area where the device is deployed`,
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
					Description: `bit length of the network mask of the address`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `address of router`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `whether the address is routable from the Internet`,
				},
				resource.Attribute{
					Name:        "family",
					Description: `IP version - "4" or "6"`,
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
					Name:        "project_id",
					Description: `The ID of the project the device belongs to`,
				},
				resource.Attribute{
					Name:        "root_password",
					Description: `Root password to the server (disabled after 24 hours)`,
				},
				resource.Attribute{
					Name:        "ssh_key_ids",
					Description: `List of IDs of SSH keys deployed in the device, can be both user and project SSH keys`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The status of the device`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags attached to the device`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `The timestamp for the last time the device was updated ## Import This resource can be imported using an existing device ID: ` + "`" + `` + "`" + `` + "`" + `sh terraform import metal_device {existing_device_id} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "metal_device_network_type",
			Category:         "Resources",
			ShortDescription: `Provides a resource to manage network type of Equinix Metal devices.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "device_id",
					Description: `(Required) The ID of the device on which the network type should be set.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Network type to set. Must be one of ` + "`" + `layer3` + "`" + `, ` + "`" + `hybrid` + "`" + `, ` + "`" + `layer2-individual` + "`" + ` and ` + "`" + `layer2-bonded` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the controlled device. Use this in linked resources, if you need to wait for the network type change. It is the same as ` + "`" + `device_id` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the controlled device. Use this in linked resources, if you need to wait for the network type change. It is the same as ` + "`" + `device_id` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "metal_gateway",
			Category:         "Resources",
			ShortDescription: `Create Equinix Metal Gateways`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
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
					Description: `(Optional) UUID of IP reservation block to bind to the gateway, the reservation must be in the same metro as the VLAN, conflicts with ` + "`" + `private_ipv4_subnet_size` + "`" + ``,
				},
				resource.Attribute{
					Name:        "private_ipv4_subnet_size",
					Description: `(Optional) Size of the private IPv4 subnet to create for this metal gateway, must be one of (8, 16, 32, 64, 128), conflicts with ` + "`" + `ip_reservation_id` + "`" + ` ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Status of the gateway resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "state",
					Description: `Status of the gateway resource`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "metal_ip_attachment",
			Category:         "Resources",
			ShortDescription: `Provides a Resource for Attaching IP Subnets from a Reserved Block to a Device`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "device_id",
					Description: `(Required) ID of device to which to assign the subnet`,
				},
				resource.Attribute{
					Name:        "cidr_notation",
					Description: `(Required) CIDR notation of subnet from block reserved in the same project and facility as the device ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the assignment`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `ID of device to which subnet is assigned`,
				},
				resource.Attribute{
					Name:        "cidr_notation",
					Description: `Assigned subnet in CIDR notation, e.g. "147.229.15.30/31"`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `IP address of gateway for the subnet`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `Subnet network address`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `Subnet mask in decimal notation, e.g. "255.255.255.0"`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `length of CIDR prefix of the subnet as integer`,
				},
				resource.Attribute{
					Name:        "address_family",
					Description: `Address family as integer (4 or 6)`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `boolean flag whether subnet is reachable from the Internet`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the assignment`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `ID of device to which subnet is assigned`,
				},
				resource.Attribute{
					Name:        "cidr_notation",
					Description: `Assigned subnet in CIDR notation, e.g. "147.229.15.30/31"`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `IP address of gateway for the subnet`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `Subnet network address`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `Subnet mask in decimal notation, e.g. "255.255.255.0"`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `length of CIDR prefix of the subnet as integer`,
				},
				resource.Attribute{
					Name:        "address_family",
					Description: `Address family as integer (4 or 6)`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `boolean flag whether subnet is reachable from the Internet`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "metal_organization",
			Category:         "Resources",
			ShortDescription: `Provides an Equinix Metal Organization resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Organization`,
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
					Description: `Logo URL ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the organization`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Organization`,
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
					Description: `Logo URL ## Import This resource can be imported using an existing organization ID: ` + "`" + `` + "`" + `` + "`" + `sh terraform import metal_organization {existing_organization_id} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the organization`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Organization`,
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
					Description: `Logo URL ## Import This resource can be imported using an existing organization ID: ` + "`" + `` + "`" + `` + "`" + `sh terraform import metal_organization {existing_organization_id} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "metal_port",
			Category:         "Resources",
			ShortDescription: `Manipulate device ports`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "port_id",
					Description: `(Required) ID of the port to read`,
				},
				resource.Attribute{
					Name:        "bonded",
					Description: `(Required) Whether the port should be bonded`,
				},
				resource.Attribute{
					Name:        "layer2",
					Description: `(Optional) Whether to put the port to Layer 2 mode, valid only for bond ports`,
				},
				resource.Attribute{
					Name:        "vlan_ids",
					Description: `(Optional) List of VLAN UUIDs to attach to the port, valid only for L2 and Hybrid ports`,
				},
				resource.Attribute{
					Name:        "vxlan_ids",
					Description: `(Optional) List of VXLAN IDs to attach to the port, valid only for L2 and Hybrid ports`,
				},
				resource.Attribute{
					Name:        "native_vlan_id",
					Description: `(Optional) UUID of a VLAN to assign as a native VLAN. It must be one of attached VLANs (from ` + "`" + `vlan_ids` + "`" + ` parameter), valid only for physical (non-bond) ports`,
				},
				resource.Attribute{
					Name:        "reset_on_delete",
					Description: `(Optional) Behavioral setting to reset the port to default settings. For a bond port it means layer3 without vlans attached, eth ports will be bonded without native vlan and vlans attached ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the port, e.g. ` + "`" + `bond0` + "`" + ` or ` + "`" + `eth0` + "`" + ``,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `One of layer2-bonded, layer2-individual, layer3, hybrid and hybrid-bonded. This attribute is only set on bond ports.`,
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
					Description: `UUID of the bond port`,
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
					Name:        "vlan_ids",
					Description: `List of VLAN UUIDs to attach to the port`,
				},
				resource.Attribute{
					Name:        "vxlan_ids",
					Description: `List of VXLAN IDs to attach to the port`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name of the port, e.g. ` + "`" + `bond0` + "`" + ` or ` + "`" + `eth0` + "`" + ``,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `One of layer2-bonded, layer2-individual, layer3, hybrid and hybrid-bonded. This attribute is only set on bond ports.`,
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
					Description: `UUID of the bond port`,
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
					Name:        "vlan_ids",
					Description: `List of VLAN UUIDs to attach to the port`,
				},
				resource.Attribute{
					Name:        "vxlan_ids",
					Description: `List of VXLAN IDs to attach to the port`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "metal_port_vlan_attachment",
			Category:         "Resources",
			ShortDescription: `Provides a Resource for Attaching VLANs to Device Ports`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "device_id",
					Description: `(Required) ID of device to be assigned to the VLAN`,
				},
				resource.Attribute{
					Name:        "port_name",
					Description: `(Required) Name of network port to be assigned to the VLAN`,
				},
				resource.Attribute{
					Name:        "force_bond",
					Description: `Add port back to the bond when this resource is removed. Default is false.`,
				},
				resource.Attribute{
					Name:        "vlan_vnid",
					Description: `VXLAN Network Identifier, integer`,
				},
				resource.Attribute{
					Name:        "native",
					Description: `(Optional) Mark this VLAN a native VLAN on the port. This can be used only if this assignment assigns second or further VLAN to the port. To ensure that this attachment is not first on a port, you can use ` + "`" + `depends_on` + "`" + ` pointing to another metal_port_vlan_attachment, just like in the layer2-individual example above. ## Attribute Referece`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `UUID of device port used in the assignment`,
				},
				resource.Attribute{
					Name:        "vlan_id",
					Description: `UUID of VLAN API resource`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `UUID of device port`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "metal_project",
			Category:         "Resources",
			ShortDescription: `Provides an Equinix Metal Project resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the project`,
				},
				resource.Attribute{
					Name:        "organization_id",
					Description: `The UUID of organization under which you want to create the project. If you leave it out, the project will be create under your the default organization of your account.`,
				},
				resource.Attribute{
					Name:        "payment_method_id",
					Description: `The UUID of payment method for this project. The payment method and the project need to belong to the same organization (passed with ` + "`" + `organization_id` + "`" + `, or default).`,
				},
				resource.Attribute{
					Name:        "backend_transfer",
					Description: `Enable or disable [Backend Transfer](https://metal.equinix.com/developers/docs/networking/backend-transfer/), default is false`,
				},
				resource.Attribute{
					Name:        "bgp_config",
					Description: `Optional BGP settings. Refer to [Equinix Metal guide for BGP](https://metal.equinix.com/developers/docs/networking/local-global-bgp/). Once you set the BGP config in a project, it can't be removed (due to a limitation in the Equinix Metal API). It can be updated. The ` + "`" + `bgp_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "asn",
					Description: `Autonomous System Number for local BGP deployment`,
				},
				resource.Attribute{
					Name:        "md5",
					Description: `(Optional) Password for BGP session in plaintext (not a checksum)`,
				},
				resource.Attribute{
					Name:        "deployment_type",
					Description: `` + "`" + `private` + "`" + ` or ` + "`" + `public` + "`" + `, the ` + "`" + `private` + "`" + ` is likely to be usable immediately, the ` + "`" + `public` + "`" + ` will need to be review by Equinix Metal engineers ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the project`,
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
					Description: `The timestamp for when the project was created`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `The timestamp for the last time the project was updated The ` + "`" + `bgp_config` + "`" + ` block additionally exports:`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `status of BGP configuration in the project`,
				},
				resource.Attribute{
					Name:        "max_prefix",
					Description: `The maximum number of route filters allowed per server ## Import This resource can be imported using an existing project ID: ` + "`" + `` + "`" + `` + "`" + `sh terraform import metal_project {existing_project_id} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the project`,
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
					Description: `The timestamp for when the project was created`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `The timestamp for the last time the project was updated The ` + "`" + `bgp_config` + "`" + ` block additionally exports:`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `status of BGP configuration in the project`,
				},
				resource.Attribute{
					Name:        "max_prefix",
					Description: `The maximum number of route filters allowed per server ## Import This resource can be imported using an existing project ID: ` + "`" + `` + "`" + `` + "`" + `sh terraform import metal_project {existing_project_id} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "metal_project_api_key",
			Category:         "Resources",
			ShortDescription: `Create Equinix Metal Project API Keys`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `UUID of the project where the API key is scoped to`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description string for the Project API Key resource`,
				},
				resource.Attribute{
					Name:        "read-only",
					Description: `Flag indicating whether the API key shoud be read-only ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `API token which can be used in Equinix Metal API clients`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "token",
					Description: `API token which can be used in Equinix Metal API clients`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "metal_project_ssh_key",
			Category:         "Resources",
			ShortDescription: `Provides an Equinix Metal Project SSH key resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the SSH key for identification`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(Required) The public key. If this is a file, it can be read using the file interpolation function`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The ID of parent project ## Attributes Reference The following attributes are exported:`,
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
			Category:         "Resources",
			ShortDescription: `Provides a Resource for reserving IP addresses in the Equinix Metal Host`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The metal project ID where to allocate the address block`,
				},
				resource.Attribute{
					Name:        "quantity",
					Description: `(Required) The number of allocated /32 addresses, a power of 2`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Either "global_ipv4" or "public_ipv4", defaults to "public_ipv4" for backward compatibility`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `(Optional) Facility where to allocate the public IP address block, makes sense only for type==public_ipv4, must be empty for type==global_ipv4, conflicts with ` + "`" + `metro` + "`" + ``,
				},
				resource.Attribute{
					Name:        "metro",
					Description: `(Optional) Metro where to allocate the public IP address block, makes sense only for type==public_ipv4, must be empty for type==global_ipv4, conflicts with ` + "`" + `facility` + "`" + ``,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Arbitrary description`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) String list of tags`,
				},
				resource.Attribute{
					Name:        "wait_for_state",
					Description: `(Optional) Wait for the IP reservation block to reach a desired state on resource creation. One of: ` + "`" + `pending` + "`" + `, ` + "`" + `created` + "`" + `. The ` + "`" + `created` + "`" + ` state is default and recommended if the addresses are needed within the configuration. An error will be returned if a timeout or the ` + "`" + `denied` + "`" + ` state is encountered.`,
				},
				resource.Attribute{
					Name:        "custom_data",
					Description: `(Optional) Custom Data is an arbitrary object (submitted in Terraform as serialized JSON) to assign to the IP Reservation. This may be helpful for self-managed IPAM. The object must be valid JSON. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `The facility where the block was allocated, empty for global blocks`,
				},
				resource.Attribute{
					Name:        "metro",
					Description: `The metro where the block was allocated, empty for global blocks`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `To which project the addresses beling`,
				},
				resource.Attribute{
					Name:        "quantity",
					Description: `Number of "/32" addresses in the block`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the block`,
				},
				resource.Attribute{
					Name:        "cidr_notation",
					Description: `Address and mask in CIDR notation, e.g. "147.229.15.30/31"`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `Network IP address portion of the block specification`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `Mask in decimal notation, e.g. "255.255.255.0"`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `length of CIDR prefix of the block as integer`,
				},
				resource.Attribute{
					Name:        "address_family",
					Description: `Address family as integer (4 or 6)`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `boolean flag whether addresses from a block are public`,
				},
				resource.Attribute{
					Name:        "global",
					Description: `boolean flag whether addresses from a block are global (i.e. can be assigned in any facility)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `string list of tags Idempotent reference to a first "/32" address from a reserved block might look like this: ` + "`" + `join("/", [cidrhost(metal_reserved_ip_block.myblock.cidr_notation,0), "32"])` + "`" + ` ## Import This resource can be imported using an existing IP reservation ID: ` + "`" + `` + "`" + `` + "`" + `sh terraform import metal_reserved_ip_block {existing_ip_reservation_id} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "facility",
					Description: `The facility where the block was allocated, empty for global blocks`,
				},
				resource.Attribute{
					Name:        "metro",
					Description: `The metro where the block was allocated, empty for global blocks`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `To which project the addresses beling`,
				},
				resource.Attribute{
					Name:        "quantity",
					Description: `Number of "/32" addresses in the block`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the block`,
				},
				resource.Attribute{
					Name:        "cidr_notation",
					Description: `Address and mask in CIDR notation, e.g. "147.229.15.30/31"`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `Network IP address portion of the block specification`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `Mask in decimal notation, e.g. "255.255.255.0"`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `length of CIDR prefix of the block as integer`,
				},
				resource.Attribute{
					Name:        "address_family",
					Description: `Address family as integer (4 or 6)`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `boolean flag whether addresses from a block are public`,
				},
				resource.Attribute{
					Name:        "global",
					Description: `boolean flag whether addresses from a block are global (i.e. can be assigned in any facility)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `string list of tags Idempotent reference to a first "/32" address from a reserved block might look like this: ` + "`" + `join("/", [cidrhost(metal_reserved_ip_block.myblock.cidr_notation,0), "32"])` + "`" + ` ## Import This resource can be imported using an existing IP reservation ID: ` + "`" + `` + "`" + `` + "`" + `sh terraform import metal_reserved_ip_block {existing_ip_reservation_id} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "metal_spot_market_request",
			Category:         "Resources",
			ShortDescription: `Provides an Equinix Metal Spot Market Request Resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "devices_max",
					Description: `(Required) Maximum number devices to be created`,
				},
				resource.Attribute{
					Name:        "devices_min",
					Description: `(Required) Miniumum number devices to be created`,
				},
				resource.Attribute{
					Name:        "max_bid_price",
					Description: `(Required) Maximum price user is willing to pay per hour per device`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) Project ID`,
				},
				resource.Attribute{
					Name:        "wait_for_devices",
					Description: `(Optional) On resource creation - wait until all desired devices are active, on resource destruction - wait until devices are removed`,
				},
				resource.Attribute{
					Name:        "facilities",
					Description: `(Optional) Facility IDs where devices should be created`,
				},
				resource.Attribute{
					Name:        "metro",
					Description: `(Optional) Metro where devices should be created`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `(Optional) Blocks deletion of the SpotMarketRequest device until the lock is disabled`,
				},
				resource.Attribute{
					Name:        "instance_parameters",
					Description: `(Required) Parameters for devices provisioned from this request. You can find the parameter description from the [metal_device doc](device.md).`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 60 mins) Used when creating the Spot Market Request and ` + "`" + `wait_for_devices == true` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 60 mins) Used when destroying the Spot Market Request and ` + "`" + `wait for devices == true` + "`" + ` ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Spot Market Request`,
				},
				resource.Attribute{
					Name:        "facilities",
					Description: `The facilities where the Spot Market Request is applied. This is computed when ` + "`" + `metro` + "`" + ` is set or no specific location was requested. ## Import This resource can be imported using an existing spot market request ID: ` + "`" + `` + "`" + `` + "`" + `sh terraform import metal_spot_market_request {existing_spot_market_request_id} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Spot Market Request`,
				},
				resource.Attribute{
					Name:        "facilities",
					Description: `The facilities where the Spot Market Request is applied. This is computed when ` + "`" + `metro` + "`" + ` is set or no specific location was requested. ## Import This resource can be imported using an existing spot market request ID: ` + "`" + `` + "`" + `` + "`" + `sh terraform import metal_spot_market_request {existing_spot_market_request_id} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "metal_ssh_key",
			Category:         "Resources",
			ShortDescription: `Provides an Equinix Metal SSH key resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the SSH key for identification`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(Required) The public key. If this is a file, it can be read using the file interpolation function ## Attributes Reference The following attributes are exported:`,
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
					Name:        "fingerprint",
					Description: `The fingerprint of the SSH key`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `The UUID of the Equinix Metal API User who owns this key`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The timestamp for when the SSH key was created`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `The timestamp for the last time the SSH key was updated ## Import This resource can be imported using an existing SSH Key ID: ` + "`" + `` + "`" + `` + "`" + `sh terraform import metal_ssh_key {existing_sshkey_id} ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "fingerprint",
					Description: `The fingerprint of the SSH key`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `The UUID of the Equinix Metal API User who owns this key`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The timestamp for when the SSH key was created`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `The timestamp for the last time the SSH key was updated ## Import This resource can be imported using an existing SSH Key ID: ` + "`" + `` + "`" + `` + "`" + `sh terraform import metal_ssh_key {existing_sshkey_id} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "metal_user_api_key",
			Category:         "Resources",
			ShortDescription: `Create Equinix Metal User API Keys`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Description string for the User API Key resource`,
				},
				resource.Attribute{
					Name:        "read-only",
					Description: `Flag indicating whether the API key shoud be read-only ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `UUID of the owner of the API key`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `API token which can be used in Equinix Metal API clients`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "user_id",
					Description: `UUID of the owner of the API key`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `API token which can be used in Equinix Metal API clients`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "metal_virtual_circuit",
			Category:         "Resources",
			ShortDescription: `Create Equinix Fabric Virtual Circuit`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "connection_id",
					Description: `(Required) UUID of Connection where the VC is scoped to`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) UUID of the Project where the VC is scoped to`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `(Required) UUID of the Connection Port where the VC is scoped to`,
				},
				resource.Attribute{
					Name:        "nni_vlan",
					Description: `(Required) Equinix Metal network-to-network VLAN ID`,
				},
				resource.Attribute{
					Name:        "vlan_id",
					Description: `(Required) UUID of the VLAN to associate`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the Virtual Circuit resource`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `(Optional) Facility where the connection will be created`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for the Virtual Circuit resource`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags for the Virtual Circuit resource`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `(Optional) Speed of the Virtual Circuit resource ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the virtal circuit`,
				},
				resource.Attribute{
					Name:        "nni_nvid",
					Description: `VLAN parameters, see the [documentation for Equinix Fabric](https://metal.equinix.com/developers/docs/networking/fabric/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `Status of the virtal circuit`,
				},
				resource.Attribute{
					Name:        "nni_nvid",
					Description: `VLAN parameters, see the [documentation for Equinix Fabric](https://metal.equinix.com/developers/docs/networking/fabric/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "metal_vlan",
			Category:         "Resources",
			ShortDescription: `Provides a resource for Equinix Metal Virtual Network.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) ID of parent project`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `(Required) Facility where to create the VLAN`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description string`,
				},
				resource.Attribute{
					Name:        "vxlan",
					Description: `VLAN ID, must be unique in metro ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "vxlan",
					Description: `VXLAN segment ID ## Import This resource can be imported using an existing VLAN ID (UUID): ` + "`" + `` + "`" + `` + "`" + `sh terraform import metal_vlan {existing_vlan_id} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vxlan",
					Description: `VXLAN segment ID ## Import This resource can be imported using an existing VLAN ID (UUID): ` + "`" + `` + "`" + `` + "`" + `sh terraform import metal_vlan {existing_vlan_id} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "metal_volume",
			Category:         "Resources",
			ShortDescription: `(Removed) Provides an Equinix Metal Block Storage Volume Resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "metal_volume_attachment",
			Category:         "Resources",
			ShortDescription: `(Removed) Provides attachment of volumes to devices in the Equinix Metal Host.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"metal_bgp_session":          0,
		"metal_connection":           1,
		"metal_device":               2,
		"metal_device_network_type":  3,
		"metal_gateway":              4,
		"metal_ip_attachment":        5,
		"metal_organization":         6,
		"metal_port":                 7,
		"metal_port_vlan_attachment": 8,
		"metal_project":              9,
		"metal_project_api_key":      10,
		"metal_project_ssh_key":      11,
		"metal_reserved_ip_block":    12,
		"metal_spot_market_request":  13,
		"metal_ssh_key":              14,
		"metal_user_api_key":         15,
		"metal_virtual_circuit":      16,
		"metal_vlan":                 17,
		"metal_volume":               18,
		"metal_volume_attachment":    19,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
