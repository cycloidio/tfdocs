package packet

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "packet_bgp_session",
			Category:         "Resources",
			ShortDescription: `BGP session in Packet Host`,
			Description:      ``,
			Keywords: []string{
				"bgp",
				"session",
			},
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
			Type:             "packet_device",
			Category:         "Resources",
			ShortDescription: `Provides a Packet device resource. This can be used to create, modify, and delete devices.`,
			Description:      ``,
			Keywords: []string{
				"device",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required) The device name`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The ID of the project in which to create the device`,
				},
				resource.Attribute{
					Name:        "operating_system",
					Description: `(Required) The operating system slug. To find the slug, or visit [Operating Systems API docs](https://www.packet.com/developers/api/#operatingsystems), set your API auth token in the top of the page and see JSON from the API response.`,
				},
				resource.Attribute{
					Name:        "facilities",
					Description: `List of facility codes with deployment preferences. Packet API will go through the list and will deploy your device to first facility with free capacity. List items must be facility codes or ` + "`" + `any` + "`" + ` (a wildcard). To find the facility code, visit [Facilities API docs](https://www.packet.com/developers/api/#facilities), set your API auth token in the top of the page and see JSON from the API response.`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `(Required) The device plan slug. To find the plan slug, visit [Device plans API docs](https://www.packet.com/developers/api/#plans), set your auth token in the top of the page and see JSON from the API response.`,
				},
				resource.Attribute{
					Name:        "billing_cycle",
					Description: `(Required) monthly or hourly`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags attached to the device`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description string for the device`,
				},
				resource.Attribute{
					Name:        "project_ssh_key_ids",
					Description: `Array of IDs of the project SSH keys which should be added to the device. If you omit this, SSH keys of all the members of the parent project will be added to the device. If you specify this array, only the listed project SSH keys will be added. Project SSH keys can be created with the [packet_project_ssh_key][packet_project_ssh_key.html] resource.`,
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
					Name:        "description",
					Description: `Description string for the device`,
				},
				resource.Attribute{
					Name:        "hardware_reservation_id",
					Description: `The ID of hardware reservation which this device occupies`,
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
					Description: `The timestamp for the last time the device was updated`,
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
					Name:        "description",
					Description: `Description string for the device`,
				},
				resource.Attribute{
					Name:        "hardware_reservation_id",
					Description: `The ID of hardware reservation which this device occupies`,
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
					Description: `The timestamp for the last time the device was updated`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "packet_ip_attachment",
			Category:         "Resources",
			ShortDescription: `Provides a Resource for Attaching IP Subnets from a Reserved Block to a Device`,
			Description:      ``,
			Keywords: []string{
				"ip",
				"attachment",
			},
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
			Type:             "packet_organization",
			Category:         "Resources",
			ShortDescription: `Provides a Packet Organization resource.`,
			Description:      ``,
			Keywords: []string{
				"organization",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Organization.`,
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
					Description: `Logo URL. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the organization.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Organization.`,
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
					Name:        "id",
					Description: `The unique ID of the organization.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Organization.`,
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
			Type:             "packet_port_vlan_attachment",
			Category:         "Resources",
			ShortDescription: `Provides a Resource for Attaching VLANs to Device Ports`,
			Description:      ``,
			Keywords: []string{
				"port",
				"vlan",
				"attachment",
			},
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
					Description: `(Optional) Mark this VLAN a native VLAN on the port. This can be used only if this assignment assigns second or further VLAN to the port. To ensure that this attachment is not first on a port, you can use ` + "`" + `depends_on` + "`" + ` pointing to another packet_port_vlan_attachment, just like in the layer2-individual example above. ## Attribute Referece`,
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
			Type:             "packet_project",
			Category:         "Resources",
			ShortDescription: `Provides a Packet Project resource.`,
			Description:      ``,
			Keywords: []string{
				"project",
			},
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
					Description: `Enable or disable [Backend Transfer](https://support.packet.com/kb/articles/backend-transfer), default is false`,
				},
				resource.Attribute{
					Name:        "bgp_config",
					Description: `Optional BGP settings. Refer to [Packet guide for BGP](https://support.packet.com/kb/articles/bgp). Once you set the BGP config in a project, it can't be removed (due to a limitation in the Packet API). It can be updated. The ` + "`" + `bgp_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "asn",
					Description: `Autonomous System Numer for local BGP deployment`,
				},
				resource.Attribute{
					Name:        "md5",
					Description: `(Optional) Password for BGP session in plaintext (not a checksum)`,
				},
				resource.Attribute{
					Name:        "deployment_type",
					Description: `` + "`" + `private` + "`" + ` or ` + "`" + `public` + "`" + `, the ` + "`" + `private` + "`" + ` is likely to be usable immediately, the ` + "`" + `public` + "`" + ` will need to be review by Packet engineers ## Attributes Reference The following attributes are exported:`,
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
					Description: `The maximum number of route filters allowed per server`,
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
					Description: `The maximum number of route filters allowed per server`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "packet_project_ssh_key",
			Category:         "Resources",
			ShortDescription: `Provides a Packet Project SSH key resource.`,
			Description:      ``,
			Keywords: []string{
				"project",
				"ssh",
				"key",
			},
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
			Type:             "packet_reserved_ip_block",
			Category:         "Resources",
			ShortDescription: `Provides a Resource for reserving IP addresses in the Packet Host`,
			Description:      ``,
			Keywords: []string{
				"reserved",
				"ip",
				"block",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The packet project ID where to allocate the address block`,
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
					Description: `(Optional) Facility where to allocate the public IP address block, makes sense only for type==public_ipv4, must be empty for type==global_ipv4`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Arbitrary description ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `The facility where the block was allocated, empty for global blocks`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `To which project the addresses beling`,
				},
				resource.Attribute{
					Name:        "quantity",
					Description: `Number of /32 addresses in the block`,
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
					Description: `boolean flag whether addresses from a block are global (i.e. can be assigned in any facility) Idempotent reference to a first /32 address from a reserved block might look like ` + "`" + `"${cidrhost(packet_reserved_ip_block.test.cidr_notation,0)}/32"` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "facility",
					Description: `The facility where the block was allocated, empty for global blocks`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `To which project the addresses beling`,
				},
				resource.Attribute{
					Name:        "quantity",
					Description: `Number of /32 addresses in the block`,
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
					Description: `boolean flag whether addresses from a block are global (i.e. can be assigned in any facility) Idempotent reference to a first /32 address from a reserved block might look like ` + "`" + `"${cidrhost(packet_reserved_ip_block.test.cidr_notation,0)}/32"` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "packet_spot_market_request",
			Category:         "Resources",
			ShortDescription: `Provides a Packet Spot Market Request Resource.`,
			Description:      ``,
			Keywords: []string{
				"spot",
				"market",
				"request",
			},
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
					Name:        "facilities",
					Description: `(Required) Facility IDs where devices should be created`,
				},
				resource.Attribute{
					Name:        "instance_parameters",
					Description: `(Required) Device parameters. See device resource for details`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) Project ID`,
				},
				resource.Attribute{
					Name:        "wait_for_devices",
					Description: `(Optional) On resource creation - wait until all desired devices are active, on resource destruction - wait until devices are removed ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
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
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Spot Market Request`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "packet_ssh_key",
			Category:         "Resources",
			ShortDescription: `Provides a Packet SSH key resource.`,
			Description:      ``,
			Keywords: []string{
				"ssh",
				"key",
			},
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
					Description: `The UUID of the Packet API User who owns this key`,
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
					Name:        "fingerprint",
					Description: `The fingerprint of the SSH key`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `The UUID of the Packet API User who owns this key`,
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
			Type:             "packet_vlan",
			Category:         "Resources",
			ShortDescription: `Provides a resource for Packet Virtual Network.`,
			Description:      ``,
			Keywords: []string{
				"vlan",
			},
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
					Description: `Description string ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "vxlan",
					Description: `VXLAN segment ID`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vxlan",
					Description: `VXLAN segment ID`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "packet_volume",
			Category:         "Resources",
			ShortDescription: `Provides a Packet Block Storage Volume Resource.`,
			Description:      ``,
			Keywords: []string{
				"volume",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "plan",
					Description: `(Required) The service plan slug of the volume`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `(Required) The facility to create the volume in`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The packet project ID to deploy the volume in`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) The size in GB to make the volume`,
				},
				resource.Attribute{
					Name:        "billing_cycle",
					Description: `The billing cycle, defaults to "hourly"`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Optional description for the volume`,
				},
				resource.Attribute{
					Name:        "snapshot_policies",
					Description: `Optional list of snapshot policies`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `Lock or unlock the volume ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the volume`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the volume`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the volume`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size in GB of the volume`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `Performance plan the volume is on`,
				},
				resource.Attribute{
					Name:        "billing_cycle",
					Description: `The billing cycle, defaults to hourly`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `The facility slug the volume resides in`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the volume`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `Whether the volume is locked or not`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The project id the volume is in`,
				},
				resource.Attribute{
					Name:        "attachments",
					Description: `A list of attachments, each with it's own ` + "`" + `href` + "`" + ` attribute`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The timestamp for when the volume was created`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `The timestamp for the last time the volume was updated`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the volume`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the volume`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the volume`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size in GB of the volume`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `Performance plan the volume is on`,
				},
				resource.Attribute{
					Name:        "billing_cycle",
					Description: `The billing cycle, defaults to hourly`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `The facility slug the volume resides in`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the volume`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `Whether the volume is locked or not`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The project id the volume is in`,
				},
				resource.Attribute{
					Name:        "attachments",
					Description: `A list of attachments, each with it's own ` + "`" + `href` + "`" + ` attribute`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The timestamp for when the volume was created`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `The timestamp for the last time the volume was updated`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "packet_volume_attachment",
			Category:         "Resources",
			ShortDescription: `Provides attachment of volumes to devices in the Packet Host.`,
			Description:      ``,
			Keywords: []string{
				"volume",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "volume_id",
					Description: `(Required) The ID of the volume to attach`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `(Required) The ID of the device to which the volume should be attached ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the volume attachment`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the volume attachment`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"packet_bgp_session":          0,
		"packet_device":               1,
		"packet_ip_attachment":        2,
		"packet_organization":         3,
		"packet_port_vlan_attachment": 4,
		"packet_project":              5,
		"packet_project_ssh_key":      6,
		"packet_reserved_ip_block":    7,
		"packet_spot_market_request":  8,
		"packet_ssh_key":              9,
		"packet_vlan":                 10,
		"packet_volume":               11,
		"packet_volume_attachment":    12,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
