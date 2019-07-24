package aws

import (
	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*Resource{

		&resource.Resource{
			Name:             "",
			Type:             "triton_fabric",
			Category:         "Resources",
			ShortDescription: `The ` + "`" + `triton_fabric` + "`" + ` resource represents an SSH fabric for a Triton account.`,
			Description:      ``,
			Keywords: []string{
				"fabric",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(String, Required, Change forces new resource) Network name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(String, Optional, Change forces new resource) Optional description of network.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `(String, Required, Change forces new resource) CIDR formatted string describing network.`,
				},
				resource.Attribute{
					Name:        "provision_start_ip",
					Description: `(String, Required, Change forces new resource) First IP on the network that can be assigned.`,
				},
				resource.Attribute{
					Name:        "provision_end_ip",
					Description: `(String, Required, Change forces new resource) Last assignable IP on the network.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `(String, Optional, Change forces new resource) Optional gateway IP.`,
				},
				resource.Attribute{
					Name:        "resolvers",
					Description: `(List, Optional) Array of IP addresses for resolvers.`,
				},
				resource.Attribute{
					Name:        "routes",
					Description: `(Map, Optional, Change forces new resource) Map of CIDR block to Gateway IP address.`,
				},
				resource.Attribute{
					Name:        "internet_nat",
					Description: `(Bool, Optional, Change forces new resource) If a NAT zone is provisioned at Gateway IP address. Default is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vlan_id",
					Description: `(Int, Required, Change forces new resource) VLAN id the network is on. Number between 0-4095 indicating VLAN ID. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(string) - The identifier representing the network in Triton.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(String) - Network name.`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `(Bool) - Whether or not this is an RFC1918 network.`,
				},
				resource.Attribute{
					Name:        "fabric",
					Description: `(Bool) - Whether or not this network is on a fabric.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(String) - Optional description of network.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `(String) - CIDR formatted string describing network.`,
				},
				resource.Attribute{
					Name:        "provision_start_ip",
					Description: `(String) - First IP on the network that can be assigned.`,
				},
				resource.Attribute{
					Name:        "provision_end_ip",
					Description: `(String) - Last assignable IP on the network.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `(String) - Optional gateway IP.`,
				},
				resource.Attribute{
					Name:        "resolvers",
					Description: `(List) - Array of IP addresses for resolvers.`,
				},
				resource.Attribute{
					Name:        "routes",
					Description: `(Map) - Map of CIDR block to Gateway IP address.`,
				},
				resource.Attribute{
					Name:        "internet_nat",
					Description: `(Bool) - If a NAT zone is provisioned at Gateway IP address.`,
				},
				resource.Attribute{
					Name:        "vlan_id",
					Description: `(Int) - VLAN id the network is on. Number between 0-4095 indicating VLAN ID.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `(string) - The identifier representing the network in Triton.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(String) - Network name.`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `(Bool) - Whether or not this is an RFC1918 network.`,
				},
				resource.Attribute{
					Name:        "fabric",
					Description: `(Bool) - Whether or not this network is on a fabric.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(String) - Optional description of network.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `(String) - CIDR formatted string describing network.`,
				},
				resource.Attribute{
					Name:        "provision_start_ip",
					Description: `(String) - First IP on the network that can be assigned.`,
				},
				resource.Attribute{
					Name:        "provision_end_ip",
					Description: `(String) - Last assignable IP on the network.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `(String) - Optional gateway IP.`,
				},
				resource.Attribute{
					Name:        "resolvers",
					Description: `(List) - Array of IP addresses for resolvers.`,
				},
				resource.Attribute{
					Name:        "routes",
					Description: `(Map) - Map of CIDR block to Gateway IP address.`,
				},
				resource.Attribute{
					Name:        "internet_nat",
					Description: `(Bool) - If a NAT zone is provisioned at Gateway IP address.`,
				},
				resource.Attribute{
					Name:        "vlan_id",
					Description: `(Int) - VLAN id the network is on. Number between 0-4095 indicating VLAN ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "triton_firewall_rule",
			Category:         "Resources",
			ShortDescription: `The ` + "`" + `triton_firewall_rule` + "`" + ` resource represents a rule for the Triton cloud firewall.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"rule",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "rule",
					Description: `(string, Required) The firewall rule described using the Cloud API rule syntax defined at https://docs.joyent.com/public-cloud/network/firewall/cloud-firewall-rules-reference. Note: Cloud API will normalize rules based on case-sensitivity, parentheses, ordering of IP addresses, etc. This can result in Terraform updating rules repeatedly if the rule definition differs from the normalized value.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(boolean, Optional) Default: ` + "`" + `false` + "`" + ` Whether the rule should be effective.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(string, Optional) Description of the firewall rule ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(string) - The identifier representing the firewall rule in Triton.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `(string) - The identifier representing the firewall rule in Triton.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "triton_instance_template",
			Category:         "Resources",
			ShortDescription: `The ` + "`" + `triton_instance_template` + "`" + ` resource represents a Triton Service Group instance template.`,
			Description:      ``,
			Keywords: []string{
				"instance",
				"template",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "template_name",
					Description: `(string, Required) Friendly name for the instance template.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(string, Required) UUID of the image.`,
				},
				resource.Attribute{
					Name:        "package",
					Description: `(string, Required) Package name used for provisioning.`,
				},
				resource.Attribute{
					Name:        "firewall_enabled",
					Description: `(boolean, Optional) Whether to enable the firewall for group instances. Default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(map, Optional) Tags for group instances.`,
				},
				resource.Attribute{
					Name:        "networks",
					Description: `(list, Optional) Network IDs for group instances.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(map, Optional) Metadata for group instances.`,
				},
				resource.Attribute{
					Name:        "userdata",
					Description: `(string, Optional) Data copied to instance on boot. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(string) - The identifier representing the Triton Service Group instance template.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `(string) - The identifier representing the Triton Service Group instance template.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "triton_key",
			Category:         "Resources",
			ShortDescription: `The ` + "`" + `triton_key` + "`" + ` resource represents an SSH key for a Triton account.`,
			Description:      ``,
			Keywords: []string{
				"key",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(string, Change forces new resource) The name of the key. If this is left empty, the name is inferred from the comment in the SSH key material.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(string, Required, Change forces new resource) The SSH key material. In order to read this from a file, use the ` + "`" + `file` + "`" + ` interpolation.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "triton_machine",
			Category:         "Resources",
			ShortDescription: `The ` + "`" + `triton_machine` + "`" + ` resource represents a virtual machine or infrastructure container running in Triton.`,
			Description:      ``,
			Keywords: []string{
				"machine",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(string) The friendly name for the machine. Triton will generate a name if one is not specified.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(map) A mapping of tags to apply to the machine.`,
				},
				resource.Attribute{
					Name:        "cns",
					Description: `(map of CNS attributes, Optional) A mapping of [CNS](https://docs.joyent.com/public-cloud/network/cns) attributes to apply to the machine.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(map, optional) A mapping of metadata to apply to the machine.`,
				},
				resource.Attribute{
					Name:        "package",
					Description: `(string, Required) The name of the package to use for provisioning.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(string, Required) The UUID of the image to provision.`,
				},
				resource.Attribute{
					Name:        "networks",
					Description: `(list, optional) The list of networks to associate with the machine. The network ID will be in hex form, e.g ` + "`" + `xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "affinity",
					Description: `(list of Affinity rules, Optional) A list of valid [Affinity Rules](https://apidocs.joyent.com/cloudapi/#affinity-rules) to apply to the machine which assist in data center placement. Using this attribute will force resource creation to be serial. NOTE: Affinity rules are best guess and assist in placing instances across a data center. They're used at creation and not referenced after.`,
				},
				resource.Attribute{
					Name:        "(Deprecated) locality",
					Description: `(map of Locality hints, Optional) A mapping of [Locality](https://apidocs.joyent.com/cloudapi/#CreateMachine) attributes to apply to the machine that assist in data center placement. NOTE: Locality hints are only used at the time of machine creation and not referenced after. Locality is deprecated as of [CloudAPI v8.3.0](https://apidocs.joyent.com/cloudapi/#830).`,
				},
				resource.Attribute{
					Name:        "firewall_enabled",
					Description: `(boolean) Default: ` + "`" + `false` + "`" + ` Whether the cloud firewall should be enabled for this machine.`,
				},
				resource.Attribute{
					Name:        "root_authorized_keys",
					Description: `(string) The public keys authorized for root access via SSH to the machine.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(string) Data to be copied to the machine on boot.`,
				},
				resource.Attribute{
					Name:        "user_script",
					Description: `(string) The user script to run on boot (every boot on SmartMachines). To learn more about both the user script and user data see the [metadata API][2] documentation and the [Joyent Metadata Data Dictionary][1] specification.`,
				},
				resource.Attribute{
					Name:        "administrator_pw",
					Description: `(string) The initial password for the Administrator user. Only used for Windows virtual machines.`,
				},
				resource.Attribute{
					Name:        "cloud_config",
					Description: `(string) Cloud-init configuration for Linux brand machines, used instead of ` + "`" + `user_data` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "deletion_protection_enabled",
					Description: `(bool) Whether an instance is destroyable. Default is ` + "`" + `false` + "`" + `. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(string) - The identifier representing the machine in Triton.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(string) - The type of the machine (` + "`" + `smartmachine` + "`" + ` or ` + "`" + `virtualmachine` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(string) - The current state of the machine.`,
				},
				resource.Attribute{
					Name:        "dataset",
					Description: `(string) - The dataset URN with which the machine was provisioned.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `(int) - The amount of memory the machine has (in Mb).`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `(int) - The amount of disk the machine has (in Gb).`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `(list of strings) - IP addresses of the machine.`,
				},
				resource.Attribute{
					Name:        "primaryip",
					Description: `(string) - The primary (public) IP address for the machine.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `(string) - The time at which the machine was created.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `(string) - The time at which the machine was last updated.`,
				},
				resource.Attribute{
					Name:        "compute_node",
					Description: `(string) - UUID of the server on which the instance is located.`,
				},
				resource.Attribute{
					Name:        "nic",
					Description: `A list of the networks that the machine is attached to. Each network is represented by a ` + "`" + `nic` + "`" + `, each of which has the following properties:`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `The NIC's IPv4 address`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `The NIC's MAC address`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `Whether this is the machine's primary NIC`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `IPv4 netmask`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `IPv4 Gateway`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `The ID of the network to which the NIC is attached`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The provisioning state of the NIC The following attributes are used by ` + "`" + `cns` + "`" + `:`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `(list of strings) - The list of services that group this instance with others under a shared domain name.`,
				},
				resource.Attribute{
					Name:        "disable",
					Description: `(boolean) - The ability to temporarily disable CNS services domains (optional). The following attributes are used as ` + "`" + `locality` + "`" + ` hints:`,
				},
				resource.Attribute{
					Name:        "close_to",
					Description: `(list of strings) - List of container UUIDs that a new instance should be placed alongside, on the same host.`,
				},
				resource.Attribute{
					Name:        "far_from",
					Description: `(list of strings) - List of container UUIDs that a new instance should not be placed onto the same host. [1]: https://eng.joyent.com/mdata/datadict.html [2]: https://docs.joyent.com/private-cloud/instances/using-mdata`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `(string) - The identifier representing the machine in Triton.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(string) - The type of the machine (` + "`" + `smartmachine` + "`" + ` or ` + "`" + `virtualmachine` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(string) - The current state of the machine.`,
				},
				resource.Attribute{
					Name:        "dataset",
					Description: `(string) - The dataset URN with which the machine was provisioned.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `(int) - The amount of memory the machine has (in Mb).`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `(int) - The amount of disk the machine has (in Gb).`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `(list of strings) - IP addresses of the machine.`,
				},
				resource.Attribute{
					Name:        "primaryip",
					Description: `(string) - The primary (public) IP address for the machine.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `(string) - The time at which the machine was created.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `(string) - The time at which the machine was last updated.`,
				},
				resource.Attribute{
					Name:        "compute_node",
					Description: `(string) - UUID of the server on which the instance is located.`,
				},
				resource.Attribute{
					Name:        "nic",
					Description: `A list of the networks that the machine is attached to. Each network is represented by a ` + "`" + `nic` + "`" + `, each of which has the following properties:`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `The NIC's IPv4 address`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `The NIC's MAC address`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `Whether this is the machine's primary NIC`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `IPv4 netmask`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `IPv4 Gateway`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `The ID of the network to which the NIC is attached`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The provisioning state of the NIC The following attributes are used by ` + "`" + `cns` + "`" + `:`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `(list of strings) - The list of services that group this instance with others under a shared domain name.`,
				},
				resource.Attribute{
					Name:        "disable",
					Description: `(boolean) - The ability to temporarily disable CNS services domains (optional). The following attributes are used as ` + "`" + `locality` + "`" + ` hints:`,
				},
				resource.Attribute{
					Name:        "close_to",
					Description: `(list of strings) - List of container UUIDs that a new instance should be placed alongside, on the same host.`,
				},
				resource.Attribute{
					Name:        "far_from",
					Description: `(list of strings) - List of container UUIDs that a new instance should not be placed onto the same host. [1]: https://eng.joyent.com/mdata/datadict.html [2]: https://docs.joyent.com/private-cloud/instances/using-mdata`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "triton_service_group",
			Category:         "Resources",
			ShortDescription: `The ` + "`" + `triton_service_group` + "`" + ` resource represents a Triton Service Group.`,
			Description:      ``,
			Keywords: []string{
				"service",
				"group",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "group_name",
					Description: `(string, Required) Friendly name for the service group.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(string, Required) Identifier of an instance template.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `(int, Optional) Number of instances to launch and monitor. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(string) - The identifier representing the Triton Service Group.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `(string) - The identifier representing the Triton Service Group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "triton_snapshot",
			Category:         "Resources",
			ShortDescription: `The ` + "`" + `triton_snapshot` + "`" + ` resource represents a snapshot of a Triton machine.`,
			Description:      ``,
			Keywords: []string{
				"snapshot",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(string, Required) The name for the snapshot.`,
				},
				resource.Attribute{
					Name:        "machine_id",
					Description: `(string, Required) The ID of the machine of which to take a snapshot. ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(string) - The identifier representing the snapshot in Triton.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(string) - The current state of the snapshot.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `(string) - The identifier representing the snapshot in Triton.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(string) - The current state of the snapshot.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "triton_vlan",
			Category:         "Resources",
			ShortDescription: `The ` + "`" + `triton_vlan` + "`" + ` resource represents an VLAN for a Triton account.`,
			Description:      ``,
			Keywords: []string{
				"vlan",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "vlan_id",
					Description: `(int, Required, Change forces new resource) Number between 0-4095 indicating VLAN ID`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(string, Required) Unique name to identify VLAN`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(string, Optional) Description of the VLAN`,
				},
			},
			Attributes: []resource.Argument{},
		},
	}

	resourcesMap = map[string]int{

		"triton_fabric":            0,
		"triton_firewall_rule":     1,
		"triton_instance_template": 2,
		"triton_key":               3,
		"triton_machine":           4,
		"triton_service_group":     5,
		"triton_snapshot":          6,
		"triton_vlan":              7,
	}
)

func GetResource(r string) (*resouce.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs]
}
