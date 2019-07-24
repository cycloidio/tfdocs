package cloudstack

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "cloudstack_affinity_group",
			Category:         "Resources",
			ShortDescription: `Creates an affinity group.`,
			Description:      ``,
			Keywords: []string{
				"affinity",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the affinity group. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the affinity group.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The affinity group type. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The name or ID of the project to register this affinity group to. Changing this forces a new resource to be created. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the affinity group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the affinity group. ## Import Affinity groups can be imported; use ` + "`" + `<AFFINITY GROUP ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudstack_affinity_group.default 6226ea4d-9cbe-4cc9-b30c-b9532146da5b ` + "`" + `` + "`" + `` + "`" + ` When importing into a project you need to prefix the import ID with the project name: ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudstack_affinity_group.default my-project/6226ea4d-9cbe-4cc9-b30c-b9532146da5b ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the affinity group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the affinity group. ## Import Affinity groups can be imported; use ` + "`" + `<AFFINITY GROUP ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudstack_affinity_group.default 6226ea4d-9cbe-4cc9-b30c-b9532146da5b ` + "`" + `` + "`" + `` + "`" + ` When importing into a project you need to prefix the import ID with the project name: ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudstack_affinity_group.default my-project/6226ea4d-9cbe-4cc9-b30c-b9532146da5b ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudstack_disk",
			Category:         "Resources",
			ShortDescription: `Creates a disk volume from a disk offering. This disk volume will be attached to a virtual machine if the optional parameters are configured.`,
			Description:      ``,
			Keywords: []string{
				"disk",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the disk volume. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "attach",
					Description: `(Optional) Determines whether or not to attach the disk volume to a virtual machine (defaults false).`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `(Optional) The device ID to map the disk volume to within the guest OS.`,
				},
				resource.Attribute{
					Name:        "disk_offering",
					Description: `(Required) The name or ID of the disk offering to use for this disk volume.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) The size of the disk volume in gigabytes.`,
				},
				resource.Attribute{
					Name:        "shrink_ok",
					Description: `(Optional) Verifies if the disk volume is allowed to shrink when resizing (defaults false).`,
				},
				resource.Attribute{
					Name:        "virtual_machine_id",
					Description: `(Optional) The ID of the virtual machine to which you want to attach the disk volume.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The name or ID of the project to deploy this instance to. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) The name or ID of the zone where this disk volume will be available. Changing this forces a new resource to be created. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the disk volume.`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `The device ID the disk volume is mapped to within the guest OS. ## Import Disks can be imported; use ` + "`" + `<DISK ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudstack_disk.default 6f3ee798-d417-4e7a-92bc-95ad41cf1244 ` + "`" + `` + "`" + `` + "`" + ` When importing into a project you need to prefix the import ID with the project name: ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudstack_disk.default my-project/6f3ee798-d417-4e7a-92bc-95ad41cf1244 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the disk volume.`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `The device ID the disk volume is mapped to within the guest OS. ## Import Disks can be imported; use ` + "`" + `<DISK ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudstack_disk.default 6f3ee798-d417-4e7a-92bc-95ad41cf1244 ` + "`" + `` + "`" + `` + "`" + ` When importing into a project you need to prefix the import ID with the project name: ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudstack_disk.default my-project/6f3ee798-d417-4e7a-92bc-95ad41cf1244 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudstack_egress_firewall",
			Category:         "Resources",
			ShortDescription: `Creates egress firewall rules for a given network.`,
			Description:      ``,
			Keywords: []string{
				"egress",
				"firewall",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required) The network ID for which to create the egress firewall rules. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "managed",
					Description: `(Optional) USE WITH CAUTION! If enabled all the egress firewall rules for this network will be managed by this resource. This means it will delete all firewall rules that are not in your config! (defaults false)`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Optional) Can be specified multiple times. Each rule block supports fields documented below. If ` + "`" + `managed = false` + "`" + ` at least one rule is required!`,
				},
				resource.Attribute{
					Name:        "cidr_list",
					Description: `(Required) A CIDR list to allow access to the given ports.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The name of the protocol to allow. Valid options are: ` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + ` and ` + "`" + `icmp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "icmp_type",
					Description: `(Optional) The ICMP type to allow. This can only be specified if the protocol is ICMP.`,
				},
				resource.Attribute{
					Name:        "icmp_code",
					Description: `(Optional) The ICMP code to allow. This can only be specified if the protocol is ICMP.`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `(Optional) List of ports and/or port ranges to allow. This can only be specified if the protocol is TCP or UDP. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The network ID for which the egress firewall rules are created.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The network ID for which the egress firewall rules are created.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudstack_firewall",
			Category:         "Resources",
			ShortDescription: `Creates firewall rules for a given IP address.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_address_id",
					Description: `(Required) The IP address ID for which to create the firewall rules. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "managed",
					Description: `(Optional) USE WITH CAUTION! If enabled all the firewall rules for this IP address will be managed by this resource. This means it will delete all firewall rules that are not in your config! (defaults false)`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Optional) Can be specified multiple times. Each rule block supports fields documented below. If ` + "`" + `managed = false` + "`" + ` at least one rule is required!`,
				},
				resource.Attribute{
					Name:        "cidr_list",
					Description: `(Required) A CIDR list to allow access to the given ports.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The name of the protocol to allow. Valid options are: ` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + ` and ` + "`" + `icmp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "icmp_type",
					Description: `(Optional) The ICMP type to allow. This can only be specified if the protocol is ICMP.`,
				},
				resource.Attribute{
					Name:        "icmp_code",
					Description: `(Optional) The ICMP code to allow. This can only be specified if the protocol is ICMP.`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `(Optional) List of ports and/or port ranges to allow. This can only be specified if the protocol is TCP or UDP. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The IP address ID for which the firewall rules are created.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The IP address ID for which the firewall rules are created.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudstack_instance",
			Category:         "Resources",
			ShortDescription: `Creates and automatically starts a virtual machine based on a service offering, disk offering, and template.`,
			Description:      ``,
			Keywords: []string{
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the instance.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of the instance.`,
				},
				resource.Attribute{
					Name:        "service_offering",
					Description: `(Required) The name or ID of the service offering used for this instance.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Optional) The ID of the network to connect this instance to. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) The IP address to assign to this instance. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Required) The name or ID of the template used for this instance. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "root_disk_size",
					Description: `(Optional) The size of the root disk in gigabytes. The root disk is resized on deploy. Only applies to template-based deployments. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `(Optional) The group name of the instance.`,
				},
				resource.Attribute{
					Name:        "affinity_group_ids",
					Description: `(Optional) List of affinity group IDs to apply to this instance.`,
				},
				resource.Attribute{
					Name:        "affinity_group_names",
					Description: `(Optional) List of affinity group names to apply to this instance.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Optional) List of security group IDs to apply to this instance. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "security_group_names",
					Description: `(Optional) List of security group names to apply to this instance. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The name or ID of the project to deploy this instance to. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) The name or ID of the zone where this instance will be created. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "start_vm",
					Description: `(Optional) This determines if the instances is started after it is created (defaults true)`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional) The user data to provide when launching the instance. This can be either plain text or base64 encoded text.`,
				},
				resource.Attribute{
					Name:        "keypair",
					Description: `(Optional) The name of the SSH key pair that will be used to access this instance.`,
				},
				resource.Attribute{
					Name:        "expunge",
					Description: `(Optional) This determines if the instance is expunged when it is destroyed (defaults false) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The instance ID.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name of the instance. ## Import Instances can be imported; use ` + "`" + `<INSTANCE ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudstack_instance.default 5cf69677-7e4b-4bf4-b868-f0b02bb72ee0 ` + "`" + `` + "`" + `` + "`" + ` When importing into a project you need to prefix the import ID with the project name: ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudstack_instance.default my-project/5cf69677-7e4b-4bf4-b868-f0b02bb72ee0 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The instance ID.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name of the instance. ## Import Instances can be imported; use ` + "`" + `<INSTANCE ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudstack_instance.default 5cf69677-7e4b-4bf4-b868-f0b02bb72ee0 ` + "`" + `` + "`" + `` + "`" + ` When importing into a project you need to prefix the import ID with the project name: ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudstack_instance.default my-project/5cf69677-7e4b-4bf4-b868-f0b02bb72ee0 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudstack_ipaddress",
			Category:         "Resources",
			ShortDescription: `Acquires and associates a public IP.`,
			Description:      ``,
			Keywords: []string{
				"ipaddress",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "is_portable",
					Description: `(Optional) This determines if the IP address should be transferable across zones (defaults false)`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Optional) The ID of the network for which an IP address should be acquired and associated. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) The ID of the VPC for which an IP address should be acquired and associated. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) The name or ID of the zone for which an IP address should be acquired and associated. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The name or ID of the project to deploy this instance to. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the acquired and associated IP address.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The IP address that was acquired and associated.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the acquired and associated IP address.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The IP address that was acquired and associated.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudstack_loadbalancer_rule",
			Category:         "Resources",
			ShortDescription: `Creates a load balancer rule.`,
			Description:      ``,
			Keywords: []string{
				"loadbalancer",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the loadbalancer rule. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the load balancer rule.`,
				},
				resource.Attribute{
					Name:        "ip_address_id",
					Description: `(Required) Public IP address ID from where the network traffic will be load balanced from. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Optional) The network ID this rule will be created for. Required when public IP address is not associated with any network yet (VPC case).`,
				},
				resource.Attribute{
					Name:        "algorithm",
					Description: `(Required) Load balancer rule algorithm (source, roundrobin, leastconn). Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "private_port",
					Description: `(Required) The private port of the private IP address (virtual machine) where the network traffic will be load balanced to. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "public_port",
					Description: `(Required) The public port from where the network traffic will be load balanced from. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) Load balancer protocol (tcp, udp, tcp-proxy). Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "member_ids",
					Description: `(Required) List of instance IDs to assign to the load balancer rule. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The name or ID of the project to deploy this instance to. Changing this forces a new resource to be created. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The load balancer rule ID.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the load balancer rule.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The load balancer rule ID.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the load balancer rule.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudstack_network",
			Category:         "Resources",
			ShortDescription: `Creates a network.`,
			Description:      ``,
			Keywords: []string{
				"network",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the network.`,
				},
				resource.Attribute{
					Name:        "display_text",
					Description: `(Optional) The display text of the network.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Required) The CIDR block for the network. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `(Optional) Gateway that will be provided to the instances in this network. Defaults to the first usable IP in the range.`,
				},
				resource.Attribute{
					Name:        "startip",
					Description: `(Optional) Start of the IP block that will be available on the network. Defaults to the second available IP in the range.`,
				},
				resource.Attribute{
					Name:        "endip",
					Description: `(Optional) End of the IP block that will be available on the network. Defaults to the last available IP in the range.`,
				},
				resource.Attribute{
					Name:        "network_domain",
					Description: `(Optional) DNS domain for the network.`,
				},
				resource.Attribute{
					Name:        "network_offering",
					Description: `(Required) The name or ID of the network offering to use for this network.`,
				},
				resource.Attribute{
					Name:        "vlan",
					Description: `(Optional) The VLAN number (1-4095) the network will use. This might be required by the Network Offering if specifyVlan=true is set. Only the ROOT admin can set this value.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) The VPC ID in which to create this network. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "acl_id",
					Description: `(Optional) The ACL ID that should be attached to the network or ` + "`" + `none` + "`" + ` if you do not want to attach an ACL. You can dynamically attach and swap ACL's, but if you want to detach an attached ACL and revert to using ` + "`" + `none` + "`" + `, this will force a new resource to be created. (defaults ` + "`" + `none` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The name or ID of the project to deploy this instance to. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "source_nat_ip",
					Description: `(Optional) If set to ` + "`" + `true` + "`" + ` a public IP will be associated with the network. This is mainly used when the network supports the source NAT service which claims the first associated IP address. This prevents the ability to manage the IP address as an independent entity.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) The name or ID of the zone where this network will be available. Changing this forces a new resource to be created. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the network.`,
				},
				resource.Attribute{
					Name:        "display_text",
					Description: `The display text of the network.`,
				},
				resource.Attribute{
					Name:        "network_domain",
					Description: `DNS domain for the network.`,
				},
				resource.Attribute{
					Name:        "source_nat_ip_id",
					Description: `The ID of the associated source NAT IP. ## Import Networks can be imported; use ` + "`" + `<NETWORK ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudstack_network.default 36619b20-5584-43bf-9a84-e242bacd5582 ` + "`" + `` + "`" + `` + "`" + ` When importing into a project you need to prefix the import ID with the project name: ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudstack_network.default my-project/36619b20-5584-43bf-9a84-e242bacd5582 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the network.`,
				},
				resource.Attribute{
					Name:        "display_text",
					Description: `The display text of the network.`,
				},
				resource.Attribute{
					Name:        "network_domain",
					Description: `DNS domain for the network.`,
				},
				resource.Attribute{
					Name:        "source_nat_ip_id",
					Description: `The ID of the associated source NAT IP. ## Import Networks can be imported; use ` + "`" + `<NETWORK ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudstack_network.default 36619b20-5584-43bf-9a84-e242bacd5582 ` + "`" + `` + "`" + `` + "`" + ` When importing into a project you need to prefix the import ID with the project name: ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudstack_network.default my-project/36619b20-5584-43bf-9a84-e242bacd5582 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudstack_network_acl",
			Category:         "Resources",
			ShortDescription: `Creates a Network ACL for the given VPC.`,
			Description:      ``,
			Keywords: []string{
				"network",
				"acl",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the ACL. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the ACL. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The name or ID of the project to deploy this instance to. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) The ID of the VPC to create this ACL for. Changing this forces a new resource to be created. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Network ACL ## Import Network ACLs can be imported; use ` + "`" + `<NETWORK ACL ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudstack_network_acl.default e8b5982a-1b50-4ea9-9920-6ea2290c7359 ` + "`" + `` + "`" + `` + "`" + ` When importing into a project you need to prefix the import ID with the project name: ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudstack_network_acl.default my-project/e8b5982a-1b50-4ea9-9920-6ea2290c7359 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Network ACL ## Import Network ACLs can be imported; use ` + "`" + `<NETWORK ACL ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudstack_network_acl.default e8b5982a-1b50-4ea9-9920-6ea2290c7359 ` + "`" + `` + "`" + `` + "`" + ` When importing into a project you need to prefix the import ID with the project name: ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudstack_network_acl.default my-project/e8b5982a-1b50-4ea9-9920-6ea2290c7359 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudstack_network_acl_rule",
			Category:         "Resources",
			ShortDescription: `Creates network ACL rules for a given network ACL.`,
			Description:      ``,
			Keywords: []string{
				"network",
				"acl",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "acl_id",
					Description: `(Required) The network ACL ID for which to create the rules. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "managed",
					Description: `(Optional) USE WITH CAUTION! If enabled all the firewall rules for this network ACL will be managed by this resource. This means it will delete all firewall rules that are not in your config! (defaults false)`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Optional) Can be specified multiple times. Each rule block supports fields documented below. If ` + "`" + `managed = false` + "`" + ` at least one rule is required!`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The name or ID of the project to deploy this instance to. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) The action for the rule. Valid options are: ` + "`" + `allow` + "`" + ` and ` + "`" + `deny` + "`" + ` (defaults allow).`,
				},
				resource.Attribute{
					Name:        "cidr_list",
					Description: `(Required) A CIDR list to allow access to the given ports.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The name of the protocol to allow. Valid options are: ` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `, ` + "`" + `all` + "`" + ` or a valid protocol number.`,
				},
				resource.Attribute{
					Name:        "icmp_type",
					Description: `(Optional) The ICMP type to allow, or ` + "`" + `-1` + "`" + ` to allow ` + "`" + `any` + "`" + `. This can only be specified if the protocol is ICMP. (defaults 0)`,
				},
				resource.Attribute{
					Name:        "icmp_code",
					Description: `(Optional) The ICMP code to allow, or ` + "`" + `-1` + "`" + ` to allow ` + "`" + `any` + "`" + `. This can only be specified if the protocol is ICMP. (defaults 0)`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `(Optional) List of ports and/or port ranges to allow. This can only be specified if the protocol is TCP, UDP, ALL or a valid protocol number.`,
				},
				resource.Attribute{
					Name:        "traffic_type",
					Description: `(Optional) The traffic type for the rule. Valid options are: ` + "`" + `ingress` + "`" + ` or ` + "`" + `egress` + "`" + ` (defaults ingress). ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ACL ID for which the rules are created.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ACL ID for which the rules are created.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudstack_nic",
			Category:         "Resources",
			ShortDescription: `Creates an additional NIC to add a VM to the specified network.`,
			Description:      ``,
			Keywords: []string{
				"nic",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required) The ID of the network to plug the NIC into. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) The IP address to assign to the NIC. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "virtual_machine_id",
					Description: `(Required) The ID of the virtual machine to which to attach the NIC. Changing this forces a new resource to be created. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the NIC.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The assigned IP address.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the NIC.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The assigned IP address.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudstack_port_forward",
			Category:         "Resources",
			ShortDescription: `Creates port forwards.`,
			Description:      ``,
			Keywords: []string{
				"port",
				"forward",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_address_id",
					Description: `(Required) The IP address ID for which to create the port forwards. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "managed",
					Description: `(Optional) USE WITH CAUTION! If enabled all the port forwards for this IP address will be managed by this resource. This means it will delete all port forwards that are not in your config! (defaults false)`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The name or ID of the project to create this port forward in. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "forward",
					Description: `(Required) Can be specified multiple times. Each forward block supports fields documented below. The ` + "`" + `forward` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The name of the protocol to allow. Valid options are: ` + "`" + `tcp` + "`" + ` and ` + "`" + `udp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "private_port",
					Description: `(Required) The private port to forward to.`,
				},
				resource.Attribute{
					Name:        "public_port",
					Description: `(Required) The public port to forward from.`,
				},
				resource.Attribute{
					Name:        "virtual_machine_id",
					Description: `(Required) The ID of the virtual machine to forward to.`,
				},
				resource.Attribute{
					Name:        "vm_guest_ip",
					Description: `(Optional) The virtual machine IP address for the port forwarding rule (useful when the virtual machine has secondairy NICs or IP addresses). ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the IP address for which the port forwards are created.`,
				},
				resource.Attribute{
					Name:        "vm_guest_ip",
					Description: `The IP address of the virtual machine that is used for the port forwarding rule.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the IP address for which the port forwards are created.`,
				},
				resource.Attribute{
					Name:        "vm_guest_ip",
					Description: `The IP address of the virtual machine that is used for the port forwarding rule.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudstack_private_gateway",
			Category:         "Resources",
			ShortDescription: `Creates a private gateway.`,
			Description:      ``,
			Keywords: []string{
				"private",
				"gateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "gateway",
					Description: `(Required) the gateway of the Private gateway. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Required) the IP address of the Private gateway. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `(Required) The netmask of the Private gateway. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "vlan",
					Description: `(Required) The VLAN number (1-4095) the network will use.`,
				},
				resource.Attribute{
					Name:        "physical_network_id",
					Description: `(Optional) The ID of the physical network this private gateway belongs to.`,
				},
				resource.Attribute{
					Name:        "network_offering",
					Description: `(Optional) The name or ID of the network offering to use for the private gateways network connection.`,
				},
				resource.Attribute{
					Name:        "acl_id",
					Description: `(Required) The ACL ID that should be attached to the network.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) The VPC ID in which to create this Private gateway. Changing this forces a new resource to be created. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the private gateway. ## Import Private gateways can be imported; use ` + "`" + `<PRIVATE GATEWAY ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudstack_private_gateway.default e42a24d2-46cb-4b18-9d41-382582fad309 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the private gateway. ## Import Private gateways can be imported; use ` + "`" + `<PRIVATE GATEWAY ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudstack_private_gateway.default e42a24d2-46cb-4b18-9d41-382582fad309 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudstack_secondary_ipaddress",
			Category:         "Resources",
			ShortDescription: `Assigns a secondary IP to a NIC.`,
			Description:      ``,
			Keywords: []string{
				"secondary",
				"ipaddress",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) The IP address to bind the to NIC. If not supplied an IP address will be selected randomly. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `(Optional) The NIC ID to which you want to attach the secondary IP address. Changing this forces a new resource to be created (defaults to the ID of the primary NIC)`,
				},
				resource.Attribute{
					Name:        "virtual_machine_id",
					Description: `(Required) The ID of the virtual machine to which you want to attach the secondary IP address. Changing this forces a new resource to be created. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The secondary IP address ID.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The IP address that was acquired and associated.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The secondary IP address ID.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The IP address that was acquired and associated.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudstack_security_group",
			Category:         "Resources",
			ShortDescription: `Creates a security group.`,
			Description:      ``,
			Keywords: []string{
				"security",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the security group. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the security group. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The name or ID of the project to create this security group in. Changing this forces a new resource to be created. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the security group. ## Import Security groups can be imported; use ` + "`" + `<SECURITY GROUP ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudstack_security_group.default e54970f1-f563-46dd-a365-2b2e9b78c54b ` + "`" + `` + "`" + `` + "`" + ` When importing into a project you need to prefix the import ID with the project name: ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudstack_security_group.default my-project/e54970f1-f563-46dd-a365-2b2e9b78c54b ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the security group. ## Import Security groups can be imported; use ` + "`" + `<SECURITY GROUP ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudstack_security_group.default e54970f1-f563-46dd-a365-2b2e9b78c54b ` + "`" + `` + "`" + `` + "`" + ` When importing into a project you need to prefix the import ID with the project name: ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudstack_security_group.default my-project/e54970f1-f563-46dd-a365-2b2e9b78c54b ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudstack_security_group_rule",
			Category:         "Resources",
			ShortDescription: `Authorizes and revokes both ingress and egress rulea for a given security group.`,
			Description:      ``,
			Keywords: []string{
				"security",
				"group",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Required) The security group ID for which to create the rules. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Required) Can be specified multiple times. Each rule block supports fields documented below.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The name or ID of the project in which the security group is created. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "cidr_list",
					Description: `(Optional) A CIDR list to allow access to the given ports.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The name of the protocol to allow. Valid options are: ` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `, ` + "`" + `all` + "`" + ` or a valid protocol number.`,
				},
				resource.Attribute{
					Name:        "icmp_type",
					Description: `(Optional) The ICMP type to allow, or ` + "`" + `-1` + "`" + ` to allow ` + "`" + `any` + "`" + `. This can only be specified if the protocol is ICMP. (defaults 0)`,
				},
				resource.Attribute{
					Name:        "icmp_code",
					Description: `(Optional) The ICMP code to allow, or ` + "`" + `-1` + "`" + ` to allow ` + "`" + `any` + "`" + `. This can only be specified if the protocol is ICMP. (defaults 0)`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `(Optional) List of ports and/or port ranges to allow. This can only be specified if the protocol is TCP, UDP, ALL or a valid protocol number.`,
				},
				resource.Attribute{
					Name:        "traffic_type",
					Description: `(Optional) The traffic type for the rule. Valid options are: ` + "`" + `ingress` + "`" + ` or ` + "`" + `egress` + "`" + `. (defaults ingress)`,
				},
				resource.Attribute{
					Name:        "user_security_group_list",
					Description: `(Optional) A list of security groups to apply the rules to. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The security group ID for which the rules are created.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The security group ID for which the rules are created.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudstack_ssh_keypair",
			Category:         "Resources",
			ShortDescription: `Creates or registers an SSH key pair.`,
			Description:      ``,
			Keywords: []string{
				"ssh",
				"keypair",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the SSH key pair. This is a unique value within a CloudStack account. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(Optional) The public key to register with CloudStack. If this is omitted, CloudStack will generate a new key pair. The key can be loaded from a file on disk using the [` + "`" + `file()` + "`" + ` function](https://www.terraform.io/docs/configuration/functions/file.html). Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The name or ID of the project to register this key to. Changing this forces a new resource to be created. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The key pair ID.`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `The fingerprint of the public key specified or created.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `The private key generated by CloudStack. Only available if CloudStack generated the key pair.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The key pair ID.`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `The fingerprint of the public key specified or created.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `The private key generated by CloudStack. Only available if CloudStack generated the key pair.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudstack_static_nat",
			Category:         "Resources",
			ShortDescription: `Enables static NAT for a given IP address.`,
			Description:      ``,
			Keywords: []string{
				"static",
				"nat",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_address_id",
					Description: `(Required) The public IP address ID for which static NAT will be enabled. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "virtual_machine_id",
					Description: `(Required) The virtual machine ID to enable the static NAT feature for. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "vm_guest_ip",
					Description: `(Optional) The virtual machine IP address to forward the static NAT traffic to (useful when the virtual machine has secondary NICs or IP addresses). Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The name or ID of the project to deploy this instance to. Changing this forces a new resource to be created. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The static nat ID.`,
				},
				resource.Attribute{
					Name:        "vm_guest_ip",
					Description: `The IP address of the virtual machine that is used to forward the static NAT traffic to.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The static nat ID.`,
				},
				resource.Attribute{
					Name:        "vm_guest_ip",
					Description: `The IP address of the virtual machine that is used to forward the static NAT traffic to.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudstack_static_route",
			Category:         "Resources",
			ShortDescription: `Creates a static route.`,
			Description:      ``,
			Keywords: []string{
				"static",
				"route",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cidr",
					Description: `(Required) The CIDR for the static route. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "gateway_id",
					Description: `(Required) The ID of the Private gateway. Changing this forces a new resource to be created. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the static route.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the static route.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudstack_template",
			Category:         "Resources",
			ShortDescription: `Registers an existing template into the CloudStack cloud.`,
			Description:      ``,
			Keywords: []string{
				"template",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the template.`,
				},
				resource.Attribute{
					Name:        "display_text",
					Description: `(Optional) The display name of the template.`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Required) The format of the template. Valid values are ` + "`" + `QCOW2` + "`" + `, ` + "`" + `RAW` + "`" + `, and ` + "`" + `VHD` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "hypervisor",
					Description: `(Required) The target hypervisor for the template. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `(Required) The OS Type that best represents the OS of this template.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The URL of where the template is hosted. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The name or ID of the project to create this template for. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) The name or ID of the zone where this template will be created. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "is_dynamically_scalable",
					Description: `(Optional) Set to indicate if the template contains tools to support dynamic scaling of VM cpu/memory (defaults false)`,
				},
				resource.Attribute{
					Name:        "is_extractable",
					Description: `(Optional) Set to indicate if the template is extractable (defaults false)`,
				},
				resource.Attribute{
					Name:        "is_featured",
					Description: `(Optional) Set to indicate if the template is featured (defaults false)`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `(Optional) Set to indicate if the template is available for all accounts (defaults true)`,
				},
				resource.Attribute{
					Name:        "password_enabled",
					Description: `(Optional) Set to indicate if the template should be password enabled (defaults false)`,
				},
				resource.Attribute{
					Name:        "is_ready_timeout",
					Description: `(Optional) The maximum time in seconds to wait until the template is ready for use (defaults 300 seconds) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The template ID.`,
				},
				resource.Attribute{
					Name:        "display_text",
					Description: `The display text of the template.`,
				},
				resource.Attribute{
					Name:        "is_dynamically_scalable",
					Description: `Set to "true" if the template is dynamically scalable.`,
				},
				resource.Attribute{
					Name:        "is_extractable",
					Description: `Set to "true" if the template is extractable.`,
				},
				resource.Attribute{
					Name:        "is_featured",
					Description: `Set to "true" if the template is featured.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `Set to "true" if the template is public.`,
				},
				resource.Attribute{
					Name:        "password_enabled",
					Description: `Set to "true" if the template is password enabled.`,
				},
				resource.Attribute{
					Name:        "is_ready",
					Description: `Set to "true" once the template is ready for use.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The template ID.`,
				},
				resource.Attribute{
					Name:        "display_text",
					Description: `The display text of the template.`,
				},
				resource.Attribute{
					Name:        "is_dynamically_scalable",
					Description: `Set to "true" if the template is dynamically scalable.`,
				},
				resource.Attribute{
					Name:        "is_extractable",
					Description: `Set to "true" if the template is extractable.`,
				},
				resource.Attribute{
					Name:        "is_featured",
					Description: `Set to "true" if the template is featured.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `Set to "true" if the template is public.`,
				},
				resource.Attribute{
					Name:        "password_enabled",
					Description: `Set to "true" if the template is password enabled.`,
				},
				resource.Attribute{
					Name:        "is_ready",
					Description: `Set to "true" once the template is ready for use.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudstack_vpc",
			Category:         "Resources",
			ShortDescription: `Creates a VPC.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the VPC.`,
				},
				resource.Attribute{
					Name:        "display_text",
					Description: `(Optional) The display text of the VPC.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Required) The CIDR block for the VPC. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "vpc_offering",
					Description: `(Required) The name or ID of the VPC offering to use for this VPC. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "network_domain",
					Description: `(Optional) The default DNS domain for networks created in this VPC. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The name or ID of the project to deploy this instance to. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) The name or ID of the zone where this disk volume will be available. Changing this forces a new resource to be created. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the VPC.`,
				},
				resource.Attribute{
					Name:        "display_text",
					Description: `The display text of the VPC.`,
				},
				resource.Attribute{
					Name:        "source_nat_ip",
					Description: `The source NAT IP assigned to the VPC. ## Import VPCs can be imported; use ` + "`" + `<VPC ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudstack_vpc.default 84b23264-917a-4712-b8bf-cd7604db43b0 ` + "`" + `` + "`" + `` + "`" + ` When importing into a project you need to prefix the import ID with the project name: ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudstack_vpc.default my-project/84b23264-917a-4712-b8bf-cd7604db43b0 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the VPC.`,
				},
				resource.Attribute{
					Name:        "display_text",
					Description: `The display text of the VPC.`,
				},
				resource.Attribute{
					Name:        "source_nat_ip",
					Description: `The source NAT IP assigned to the VPC. ## Import VPCs can be imported; use ` + "`" + `<VPC ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudstack_vpc.default 84b23264-917a-4712-b8bf-cd7604db43b0 ` + "`" + `` + "`" + `` + "`" + ` When importing into a project you need to prefix the import ID with the project name: ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudstack_vpc.default my-project/84b23264-917a-4712-b8bf-cd7604db43b0 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudstack_vpn_connection",
			Category:         "Resources",
			ShortDescription: `Creates a site to site VPN connection.`,
			Description:      ``,
			Keywords: []string{
				"vpn",
				"connection",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "customer_gateway_id",
					Description: `(Required) The Customer Gateway ID to connect. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "vpn_gateway_id",
					Description: `(Required) The VPN Gateway ID to connect. Changing this forces a new resource to be created. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the VPN Connection.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the VPN Connection.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudstack_vpn_customer_gateway",
			Category:         "Resources",
			ShortDescription: `Creates a site to site VPN local customer gateway.`,
			Description:      ``,
			Keywords: []string{
				"vpn",
				"customer",
				"gateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the VPN Customer Gateway.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Required) The CIDR block that needs to be routed through this gateway.`,
				},
				resource.Attribute{
					Name:        "esp_policy",
					Description: `(Required) The ESP policy to use for this VPN Customer Gateway.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `(Required) The public IP address of the related VPN Gateway.`,
				},
				resource.Attribute{
					Name:        "ike_policy",
					Description: `(Required) The IKE policy to use for this VPN Customer Gateway.`,
				},
				resource.Attribute{
					Name:        "ipsec_psk",
					Description: `(Required) The IPSEC pre-shared key used for this gateway.`,
				},
				resource.Attribute{
					Name:        "dpd",
					Description: `(Optional) If DPD is enabled for the related VPN connection (defaults false)`,
				},
				resource.Attribute{
					Name:        "esp_lifetime",
					Description: `(Optional) The ESP lifetime of phase 2 VPN connection to this VPN Customer Gateway in seconds (defaults 86400)`,
				},
				resource.Attribute{
					Name:        "ike_lifetime",
					Description: `(Optional) The IKE lifetime of phase 2 VPN connection to this VPN Customer Gateway in seconds (defaults 86400)`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The name or ID of the project to create this VPN Customer Gateway in. Changing this forces a new resource to be created. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the VPN Customer Gateway.`,
				},
				resource.Attribute{
					Name:        "dpd",
					Description: `Enable or disable DPD is enabled for the related VPN connection.`,
				},
				resource.Attribute{
					Name:        "esp_lifetime",
					Description: `The ESP lifetime of phase 2 VPN connection to this VPN Customer Gateway.`,
				},
				resource.Attribute{
					Name:        "ike_lifetime",
					Description: `The IKE lifetime of phase 2 VPN connection to this VPN Customer Gateway. ## Import VPN customer gateways can be imported; use ` + "`" + `<VPN CUSTOMER GATEWAY ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudstack_vpn_customer_gateway.default 741a7fca-1d05-4bb6-9290-1008300f0e5a ` + "`" + `` + "`" + `` + "`" + ` When importing into a project you need to prefix the import ID with the project name: ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudstack_vpn_customer_gateway.default my-project/741a7fca-1d05-4bb6-9290-1008300f0e5a ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the VPN Customer Gateway.`,
				},
				resource.Attribute{
					Name:        "dpd",
					Description: `Enable or disable DPD is enabled for the related VPN connection.`,
				},
				resource.Attribute{
					Name:        "esp_lifetime",
					Description: `The ESP lifetime of phase 2 VPN connection to this VPN Customer Gateway.`,
				},
				resource.Attribute{
					Name:        "ike_lifetime",
					Description: `The IKE lifetime of phase 2 VPN connection to this VPN Customer Gateway. ## Import VPN customer gateways can be imported; use ` + "`" + `<VPN CUSTOMER GATEWAY ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudstack_vpn_customer_gateway.default 741a7fca-1d05-4bb6-9290-1008300f0e5a ` + "`" + `` + "`" + `` + "`" + ` When importing into a project you need to prefix the import ID with the project name: ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudstack_vpn_customer_gateway.default my-project/741a7fca-1d05-4bb6-9290-1008300f0e5a ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudstack_vpn_gateway",
			Category:         "Resources",
			ShortDescription: `Creates a site to site VPN local gateway.`,
			Description:      ``,
			Keywords: []string{
				"vpn",
				"gateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) The ID of the VPC for which to create the VPN Gateway. Changing this forces a new resource to be created. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the VPN Gateway.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP address associated with the VPN Gateway. ## Import VPC gateways can be imported; use ` + "`" + `<VPN GATEWAY ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudstack_vpn_gateway.default 49cf1821-3b9f-4627-be19-8a15ffec508d ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the VPN Gateway.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP address associated with the VPN Gateway. ## Import VPC gateways can be imported; use ` + "`" + `<VPN GATEWAY ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import cloudstack_vpn_gateway.default 49cf1821-3b9f-4627-be19-8a15ffec508d ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"cloudstack_affinity_group":       0,
		"cloudstack_disk":                 1,
		"cloudstack_egress_firewall":      2,
		"cloudstack_firewall":             3,
		"cloudstack_instance":             4,
		"cloudstack_ipaddress":            5,
		"cloudstack_loadbalancer_rule":    6,
		"cloudstack_network":              7,
		"cloudstack_network_acl":          8,
		"cloudstack_network_acl_rule":     9,
		"cloudstack_nic":                  10,
		"cloudstack_port_forward":         11,
		"cloudstack_private_gateway":      12,
		"cloudstack_secondary_ipaddress":  13,
		"cloudstack_security_group":       14,
		"cloudstack_security_group_rule":  15,
		"cloudstack_ssh_keypair":          16,
		"cloudstack_static_nat":           17,
		"cloudstack_static_route":         18,
		"cloudstack_template":             19,
		"cloudstack_vpc":                  20,
		"cloudstack_vpn_connection":       21,
		"cloudstack_vpn_customer_gateway": 22,
		"cloudstack_vpn_gateway":          23,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
