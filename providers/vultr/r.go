package vultr

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "vultr_bare_metal_server",
			Category:         "Resources",
			ShortDescription: `Provides a Vultr bare metal server resource. This can be used to create, read, modify, and delete bare metal servers on your Vultr account.`,
			Description:      ``,
			Keywords: []string{
				"bare",
				"metal",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The ID of the region that the server is to be created in. [See List Regions](https://www.vultr.com/api/#operation/list-regions)`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `(Required) The ID of the plan that you want the server to subscribe to. [See List Plans](https://www.vultr.com/api/#tag/plans)`,
				},
				resource.Attribute{
					Name:        "os_id",
					Description: `(Optional) The ID of the operating system to be installed on the server. [See List OS](https://www.vultr.com/api/#operation/list-os)`,
				},
				resource.Attribute{
					Name:        "app_id",
					Description: `(Optional) The ID of the Vultr application to be installed on the server. [See List Applications](https://www.vultr.com/api/#operation/list-applications)`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Optional) The ID of the Vultr marketplace application to be installed on the server. [See List Applications](https://www.vultr.com/api/#operation/list-applications) Note marketplace applications are denoted by type: ` + "`" + `marketplace` + "`" + ` and you must use the ` + "`" + `image_id` + "`" + ` not the id.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) The ID of the Vultr snapshot that the server will restore for the initial installation. [See List Snapshots](https://www.vultr.com/api/#operation/list-snapshots)`,
				},
				resource.Attribute{
					Name:        "script_id",
					Description: `(Optional) The ID of the startup script you want added to the server.`,
				},
				resource.Attribute{
					Name:        "ssh_key_ids",
					Description: `(Optional) A list of SSH key IDs to apply to the server on install (only valid for Linux/FreeBSD).`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional) Generic data store, which some provisioning tools and cloud operating systems use as a configuration file. It is generally consumed only once after an instance has been launched, but individual needs may vary.`,
				},
				resource.Attribute{
					Name:        "enable_ipv6",
					Description: `(Optional) Whether the server has IPv6 networking activated.`,
				},
				resource.Attribute{
					Name:        "activation_email",
					Description: `(Optional) Whether an activation email will be sent when the server is ready.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional) The hostname to assign to the server.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) The tag to assign to the server.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Optional) A label for the server.`,
				},
				resource.Attribute{
					Name:        "reserved_ipv4",
					Description: `(Optional) IP address of the floating IP to use as the main IP of this server. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the server.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The ID of the region that the server is in.`,
				},
				resource.Attribute{
					Name:        "os",
					Description: `The string description of the operating system installed on the server.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `The amount of memory available on the server in MB.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `The description of the disk(s) on the server.`,
				},
				resource.Attribute{
					Name:        "main_ip",
					Description: `The server's main IP address.`,
				},
				resource.Attribute{
					Name:        "cpu_count",
					Description: `The number of CPUs available on the server.`,
				},
				resource.Attribute{
					Name:        "default_password",
					Description: `The server's default password.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date the server was added to your Vultr account.`,
				},
				resource.Attribute{
					Name:        "netmask_v4",
					Description: `The server's IPv4 netmask.`,
				},
				resource.Attribute{
					Name:        "gateway_v4",
					Description: `The server's IPv4 gateway.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the server's subscription.`,
				},
				resource.Attribute{
					Name:        "v6_network",
					Description: `The IPv6 subnet.`,
				},
				resource.Attribute{
					Name:        "v6_main_ip",
					Description: `The main IPv6 network address.`,
				},
				resource.Attribute{
					Name:        "v6_network_size",
					Description: `The IPv6 network size in bits.`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `The ID of the plan that server is subscribed to.`,
				},
				resource.Attribute{
					Name:        "os_id",
					Description: `The ID of the operating system installed on the server.`,
				},
				resource.Attribute{
					Name:        "app_id",
					Description: `The ID of the Vultr application installed on the server.`,
				},
				resource.Attribute{
					Name:        "app_id",
					Description: `The ID of the Vultr marketplace application installed on the server.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The ID of the Vultr snapshot that the server was restored from.`,
				},
				resource.Attribute{
					Name:        "script_id",
					Description: `The ID of the startup script that was added to the server.`,
				},
				resource.Attribute{
					Name:        "ssh_key_ids",
					Description: `A list of SSH key IDs applied to the server on install.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `Generic data store, which some provisioning tools and cloud operating systems use as a configuration file. It is generally consumed only once after an instance has been launched, but individual needs may vary.`,
				},
				resource.Attribute{
					Name:        "enable_ipv6",
					Description: `Whether the server has IPv6 networking activated.`,
				},
				resource.Attribute{
					Name:        "activation_email",
					Description: `Whether an activation email was sent when the server was ready.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `The hostname assigned to the server.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `The tag assigned to the server.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `A label for the server.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `The MAC address associated with the server. ## Import Bare Metal Servers can be imported using the server ` + "`" + `ID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_bare_metal_server.my_server b6a859c5-b299-49dd-8888-b1abbc517d08 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the server.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The ID of the region that the server is in.`,
				},
				resource.Attribute{
					Name:        "os",
					Description: `The string description of the operating system installed on the server.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `The amount of memory available on the server in MB.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `The description of the disk(s) on the server.`,
				},
				resource.Attribute{
					Name:        "main_ip",
					Description: `The server's main IP address.`,
				},
				resource.Attribute{
					Name:        "cpu_count",
					Description: `The number of CPUs available on the server.`,
				},
				resource.Attribute{
					Name:        "default_password",
					Description: `The server's default password.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date the server was added to your Vultr account.`,
				},
				resource.Attribute{
					Name:        "netmask_v4",
					Description: `The server's IPv4 netmask.`,
				},
				resource.Attribute{
					Name:        "gateway_v4",
					Description: `The server's IPv4 gateway.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the server's subscription.`,
				},
				resource.Attribute{
					Name:        "v6_network",
					Description: `The IPv6 subnet.`,
				},
				resource.Attribute{
					Name:        "v6_main_ip",
					Description: `The main IPv6 network address.`,
				},
				resource.Attribute{
					Name:        "v6_network_size",
					Description: `The IPv6 network size in bits.`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `The ID of the plan that server is subscribed to.`,
				},
				resource.Attribute{
					Name:        "os_id",
					Description: `The ID of the operating system installed on the server.`,
				},
				resource.Attribute{
					Name:        "app_id",
					Description: `The ID of the Vultr application installed on the server.`,
				},
				resource.Attribute{
					Name:        "app_id",
					Description: `The ID of the Vultr marketplace application installed on the server.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The ID of the Vultr snapshot that the server was restored from.`,
				},
				resource.Attribute{
					Name:        "script_id",
					Description: `The ID of the startup script that was added to the server.`,
				},
				resource.Attribute{
					Name:        "ssh_key_ids",
					Description: `A list of SSH key IDs applied to the server on install.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `Generic data store, which some provisioning tools and cloud operating systems use as a configuration file. It is generally consumed only once after an instance has been launched, but individual needs may vary.`,
				},
				resource.Attribute{
					Name:        "enable_ipv6",
					Description: `Whether the server has IPv6 networking activated.`,
				},
				resource.Attribute{
					Name:        "activation_email",
					Description: `Whether an activation email was sent when the server was ready.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `The hostname assigned to the server.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `The tag assigned to the server.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `A label for the server.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `The MAC address associated with the server. ## Import Bare Metal Servers can be imported using the server ` + "`" + `ID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_bare_metal_server.my_server b6a859c5-b299-49dd-8888-b1abbc517d08 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vultr_block_storage",
			Category:         "Resources",
			ShortDescription: `Provides a Vultr Block Storage resource. This can be used to create, read, modify, and delete Block Storage.`,
			Description:      ``,
			Keywords: []string{
				"block",
				"storage",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "size_gb",
					Description: `(Required) The size of the given block storage.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Region in which this block storage will reside in. (Currently only NJ/NY supported region "ewr")`,
				},
				resource.Attribute{
					Name:        "attached_to_instance",
					Description: `(Optional) VPS ID that you want to have this block storage attached to.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Optional) Label that is given to your block storage.`,
				},
				resource.Attribute{
					Name:        "block_type",
					Description: `(Optional) Determines on the type of block storage volume that will be created. Soon to become a required parameter. Options are ` + "`" + `high_per` + "`" + ` or ` + "`" + `storage_opt` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "live",
					Description: `(Optional) Boolean value that will allow attachment of the volume to an instance without a restart. Default is false. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID for this block storage.`,
				},
				resource.Attribute{
					Name:        "size_gb",
					Description: `The size of the given block storage.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Region in which this block storage will reside in. (Currently only NJ/NY supported region "ewr")`,
				},
				resource.Attribute{
					Name:        "attached_to_instance",
					Description: `VPS ID that is attached to this block storage.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `Label that is given to your block storage.`,
				},
				resource.Attribute{
					Name:        "cost",
					Description: `The monthly cost of this block storage.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date this block storage was created.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of your block storage.`,
				},
				resource.Attribute{
					Name:        "live",
					Description: `Flag which will determine if a volume should be attached with a restart or not.`,
				},
				resource.Attribute{
					Name:        "mount_id",
					Description: `An ID associated with the instance, when mounted the ID can be found in /dev/disk/by-id prefixed with virtio.`,
				},
				resource.Attribute{
					Name:        "block_type",
					Description: `The type of block storage volume. Values are ` + "`" + `high_per` + "`" + ` or ` + "`" + `storage_opt` + "`" + `. ## Import Block Storage can be imported using the Block Storage ` + "`" + `ID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_block_storage.my_blockstorage e315835e-d466-4e89-9b4c-dfd8788d7685 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID for this block storage.`,
				},
				resource.Attribute{
					Name:        "size_gb",
					Description: `The size of the given block storage.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Region in which this block storage will reside in. (Currently only NJ/NY supported region "ewr")`,
				},
				resource.Attribute{
					Name:        "attached_to_instance",
					Description: `VPS ID that is attached to this block storage.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `Label that is given to your block storage.`,
				},
				resource.Attribute{
					Name:        "cost",
					Description: `The monthly cost of this block storage.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date this block storage was created.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of your block storage.`,
				},
				resource.Attribute{
					Name:        "live",
					Description: `Flag which will determine if a volume should be attached with a restart or not.`,
				},
				resource.Attribute{
					Name:        "mount_id",
					Description: `An ID associated with the instance, when mounted the ID can be found in /dev/disk/by-id prefixed with virtio.`,
				},
				resource.Attribute{
					Name:        "block_type",
					Description: `The type of block storage volume. Values are ` + "`" + `high_per` + "`" + ` or ` + "`" + `storage_opt` + "`" + `. ## Import Block Storage can be imported using the Block Storage ` + "`" + `ID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_block_storage.my_blockstorage e315835e-d466-4e89-9b4c-dfd8788d7685 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vultr_dns_domain",
			Category:         "Resources",
			ShortDescription: `Provides a Vultr DNS Domain resource. This can be used to create, read, modify, and delete DNS Domains.`,
			Description:      ``,
			Keywords: []string{
				"dns",
				"domain",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) Name of domain.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Optional) Instance IP you want associated to domain. If omitted this will create a domain with no records.`,
				},
				resource.Attribute{
					Name:        "dns_sec",
					Description: `(Optional) The Domain's DNSSEC status. Valid options are ` + "`" + `enabled` + "`" + ` or ` + "`" + `disabled` + "`" + `. Note ` + "`" + `disabled` + "`" + ` is default ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID is the name of the domain.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Name of domain.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date the domain was added to your account.`,
				},
				resource.Attribute{
					Name:        "dns_sec",
					Description: `The Domain's DNSSEC status ## Import DNS Domains can be imported using the Dns Domain ` + "`" + `domain` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_dns_domain.name domain.com ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID is the name of the domain.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Name of domain.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date the domain was added to your account.`,
				},
				resource.Attribute{
					Name:        "dns_sec",
					Description: `The Domain's DNSSEC status ## Import DNS Domains can be imported using the Dns Domain ` + "`" + `domain` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_dns_domain.name domain.com ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vultr_dns_record",
			Category:         "Resources",
			ShortDescription: `Provides a Vultr DNS Record resource. This can be used to create, read, modify, and delete DNS Records.`,
			Description:      ``,
			Keywords: []string{
				"dns",
				"record",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "data",
					Description: `(Required) IP Address of the instance the domain is associated with.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) Name of the DNS Domain this record will belong to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name (subdomain) for this record.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type of record.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Priority of this record (only required for MX and SRV).`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The time to live of this record. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID associated with the record.`,
				},
				resource.Attribute{
					Name:        "data",
					Description: `IP Address of the instance the domain is associated with.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Name of the DNS Domain this record will belong to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name for this record (Can be subdomain).`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of record.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Priority of this record (only required for MX and SRV).`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `The time to live of this record. ## Import DNS Records can be imported using the Dns Domain ` + "`" + `domain` + "`" + ` and DNS Record ` + "`" + `ID` + "`" + ` e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_dns_record.rec domain.com,1a0019bd-7645-4310-81bd-03bc5906940f ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID associated with the record.`,
				},
				resource.Attribute{
					Name:        "data",
					Description: `IP Address of the instance the domain is associated with.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Name of the DNS Domain this record will belong to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name for this record (Can be subdomain).`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of record.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Priority of this record (only required for MX and SRV).`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `The time to live of this record. ## Import DNS Records can be imported using the Dns Domain ` + "`" + `domain` + "`" + ` and DNS Record ` + "`" + `ID` + "`" + ` e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_dns_record.rec domain.com,1a0019bd-7645-4310-81bd-03bc5906940f ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vultr_firewall_group",
			Category:         "Resources",
			ShortDescription: `Provides a Vultr Firewall Group resource. This can be used to create, read, modify, and delete Firewall Group.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the firewall group. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the firewall group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the firewall group.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date the firewall group was created.`,
				},
				resource.Attribute{
					Name:        "date_modified",
					Description: `The date the firewall group was modified.`,
				},
				resource.Attribute{
					Name:        "instance_count",
					Description: `The number of instances that are currently using this firewall group.`,
				},
				resource.Attribute{
					Name:        "max_rule_count",
					Description: `The number of max firewall rules this group can have.`,
				},
				resource.Attribute{
					Name:        "rule_count",
					Description: `The number of firewall rules this group currently has. ## Import Firewall Groups can be imported using the Firewall Group ` + "`" + `FIREWALLGROUPID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_firewall_group.my_firewallgroup c342f929 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the firewall group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the firewall group.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date the firewall group was created.`,
				},
				resource.Attribute{
					Name:        "date_modified",
					Description: `The date the firewall group was modified.`,
				},
				resource.Attribute{
					Name:        "instance_count",
					Description: `The number of instances that are currently using this firewall group.`,
				},
				resource.Attribute{
					Name:        "max_rule_count",
					Description: `The number of max firewall rules this group can have.`,
				},
				resource.Attribute{
					Name:        "rule_count",
					Description: `The number of firewall rules this group currently has. ## Import Firewall Groups can be imported using the Firewall Group ` + "`" + `FIREWALLGROUPID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_firewall_group.my_firewallgroup c342f929 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vultr_firewall_rule",
			Category:         "Resources",
			ShortDescription: `Provides a Vultr Firewall Rule resource. This can be used to create, read, modify, and delete Firewall rules.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "firewall_group_id",
					Description: `(Required) The firewall group that the firewall rule will belong to.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The type of protocol for this firewall rule. Possible values (icmp, tcp, udp, gre, esp, ah)`,
				},
				resource.Attribute{
					Name:        "ip_type",
					Description: `(Required) The type of ip for this firewall rule. Possible values (v4, v6)`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `(Required) IP address that you want to define for this firewall rule.`,
				},
				resource.Attribute{
					Name:        "subnet_size",
					Description: `(Required) The number of bits for the subnet in CIDR notation. Example: 32.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) TCP/UDP only. This field can be a specific port or a colon separated port range.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) A simple note for a given firewall rule`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) Possible values ("", cloudflare) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The given ID for a firewall rule.`,
				},
				resource.Attribute{
					Name:        "firewall_group_id",
					Description: `The firewall group that the firewall rule belongs to.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The type of protocol for this firewall rule. Possible values (icmp, tcp, udp, gre, esp, ah)`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `IP address that is defined for this rule.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `This field can be a specific port or a colon separated port range.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `A simple note for a given firewall rule`,
				},
				resource.Attribute{
					Name:        "ip_type",
					Description: `The type of ip this rule is - may be either v4 or v6. ## Import Firewall Rules can be imported using the Firewall Group ` + "`" + `ID` + "`" + ` and Firewall Rule ` + "`" + `ID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_firewall_rule.my_rule b6a859c5-b299-49dd-8888-b1abbc517d08,1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The given ID for a firewall rule.`,
				},
				resource.Attribute{
					Name:        "firewall_group_id",
					Description: `The firewall group that the firewall rule belongs to.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The type of protocol for this firewall rule. Possible values (icmp, tcp, udp, gre, esp, ah)`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `IP address that is defined for this rule.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `This field can be a specific port or a colon separated port range.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `A simple note for a given firewall rule`,
				},
				resource.Attribute{
					Name:        "ip_type",
					Description: `The type of ip this rule is - may be either v4 or v6. ## Import Firewall Rules can be imported using the Firewall Group ` + "`" + `ID` + "`" + ` and Firewall Rule ` + "`" + `ID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_firewall_rule.my_rule b6a859c5-b299-49dd-8888-b1abbc517d08,1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vultr_instance",
			Category:         "Resources",
			ShortDescription: `Provides a Vultr instance resource. This can be used to create, read, modify, and delete instances on your Vultr account.`,
			Description:      ``,
			Keywords: []string{
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The ID of the region that the instance is to be created in. [See List Regions](https://www.vultr.com/api/#operation/list-regions)`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `(Required) The ID of the plan that you want the instance to subscribe to. [See List Plans](https://www.vultr.com/api/#tag/plans)`,
				},
				resource.Attribute{
					Name:        "os_id",
					Description: `(Optional) The ID of the operating system to be installed on the server. [See List OS](https://www.vultr.com/api/#operation/list-os)`,
				},
				resource.Attribute{
					Name:        "iso_id",
					Description: `(Optional) The ID of the ISO file to be installed on the server. [See List ISO](https://www.vultr.com/api/#operation/list-isos)`,
				},
				resource.Attribute{
					Name:        "app_id",
					Description: `(Optional) The ID of the Vultr application to be installed on the server. [See List Applications](https://www.vultr.com/api/#operation/list-applications)`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Optional) The ID of the Vultr marketplace application to be installed on the server. [See List Applications](https://www.vultr.com/api/#operation/list-applications) Note marketplace applications are denoted by type: ` + "`" + `marketplace` + "`" + ` and you must use the ` + "`" + `image_id` + "`" + ` not the id.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) The ID of the Vultr snapshot that the server will restore for the initial installation. [See List Snapshots](https://www.vultr.com/api/#operation/list-snapshots)`,
				},
				resource.Attribute{
					Name:        "script_id",
					Description: `(Optional) The ID of the startup script you want added to the server.`,
				},
				resource.Attribute{
					Name:        "firewall_group_id",
					Description: `(Optional) The ID of the firewall group to assign to the server.`,
				},
				resource.Attribute{
					Name:        "private_network_ids",
					Description: `(Optional) (Deprecated: use ` + "`" + `vpc_ids` + "`" + ` instead) A list of private network IDs to be attached to the server.`,
				},
				resource.Attribute{
					Name:        "vpc_ids",
					Description: `(Optional) A list of VPC IDs to be attached to the server.`,
				},
				resource.Attribute{
					Name:        "ssh_key_ids",
					Description: `(Optional) A list of SSH key IDs to apply to the server on install (only valid for Linux/FreeBSD).`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional) Generic data store, which some provisioning tools and cloud operating systems use as a configuration file. It is generally consumed only once after an instance has been launched, but individual needs may vary.`,
				},
				resource.Attribute{
					Name:        "backups",
					Description: `(Optional) Whether automatic backups will be enabled for this server (these have an extra charge associated with them). Values can be enabled or disabled.`,
				},
				resource.Attribute{
					Name:        "enable_ipv6",
					Description: `(Optional) Whether the server has IPv6 networking activated.`,
				},
				resource.Attribute{
					Name:        "activation_email",
					Description: `(Optional) Whether an activation email will be sent when the server is ready.`,
				},
				resource.Attribute{
					Name:        "ddos_protection",
					Description: `(Optional) Whether DDOS protection will be enabled on the server (there is an additional charge for this).`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional) The hostname to assign to the server.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) The tag to assign to the server.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Optional) A label for the server.`,
				},
				resource.Attribute{
					Name:        "reserved_ip_id",
					Description: `(Optional) ID of the floating IP to use as the main IP of this server.`,
				},
				resource.Attribute{
					Name:        "backups_schedule",
					Description: `(Optional) A block that defines the way backups should be scheduled. While this is an optional field if ` + "`" + `backups` + "`" + ` are ` + "`" + `enabled` + "`" + ` this field is mandatory. The configuration of a ` + "`" + `backups_schedule` + "`" + ` is listed below. ` + "`" + `backups_schedule` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of backup schedule Possible values are ` + "`" + `daily` + "`" + `, ` + "`" + `weekly` + "`" + `, ` + "`" + `monthly` + "`" + `, ` + "`" + `daily_alt_event` + "`" + `, or ` + "`" + `daily_alt_odd` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "hour",
					Description: `(Optional) Hour of day to run in UTC.`,
				},
				resource.Attribute{
					Name:        "dow",
					Description: `(Optional) Day of week to run. ` + "`" + `1 = Sunday` + "`" + `, ` + "`" + `2 = Monday` + "`" + `, ` + "`" + `3 = Tuesday` + "`" + `, ` + "`" + `4 = Wednesday` + "`" + `, ` + "`" + `5 = Thursday` + "`" + `, ` + "`" + `6 = Friday` + "`" + `, ` + "`" + `7 = Saturday` + "`" + ``,
				},
				resource.Attribute{
					Name:        "dom",
					Description: `(Optional) Day of month to run. Use values between 1 and 28. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the server.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The ID of the region that the server is in.`,
				},
				resource.Attribute{
					Name:        "os",
					Description: `The string description of the operating system installed on the server.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `The amount of memory available on the server in MB.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `The description of the disk(s) on the server.`,
				},
				resource.Attribute{
					Name:        "main_ip",
					Description: `The server's main IP address.`,
				},
				resource.Attribute{
					Name:        "vcpu_count",
					Description: `The number of virtual CPUs available on the server.`,
				},
				resource.Attribute{
					Name:        "default_password",
					Description: `The server's default password.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date the server was added to your Vultr account.`,
				},
				resource.Attribute{
					Name:        "allowed_bandwidth",
					Description: `The server's allowed bandwidth usage in GB.`,
				},
				resource.Attribute{
					Name:        "netmask_v4",
					Description: `The server's IPv4 netmask.`,
				},
				resource.Attribute{
					Name:        "gateway_v4",
					Description: `The server's IPv4 gateway.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the server's subscription.`,
				},
				resource.Attribute{
					Name:        "power_status",
					Description: `Whether the server is powered on or not.`,
				},
				resource.Attribute{
					Name:        "server_status",
					Description: `A more detailed server status (none, locked, installingbooting, isomounting, ok).`,
				},
				resource.Attribute{
					Name:        "v6_network",
					Description: `The IPv6 subnet.`,
				},
				resource.Attribute{
					Name:        "v6_main_ip",
					Description: `The main IPv6 network address.`,
				},
				resource.Attribute{
					Name:        "v6_network_size",
					Description: `The IPv6 network size in bits.`,
				},
				resource.Attribute{
					Name:        "internal_ip",
					Description: `The server's internal IP address.`,
				},
				resource.Attribute{
					Name:        "kvm",
					Description: `The server's current KVM URL. This URL will change periodically. It is not advised to cache this value.`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `The ID of the plan that server is subscribed to.`,
				},
				resource.Attribute{
					Name:        "os_id",
					Description: `The ID of the operating system installed on the server.`,
				},
				resource.Attribute{
					Name:        "iso_id",
					Description: `The ID of the ISO file installed on the server.`,
				},
				resource.Attribute{
					Name:        "app_id",
					Description: `The ID of the Vultr application installed on the server.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The ID of the Vultr marketplace application installed on the server.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The ID of the Vultr snapshot that the server was restored from.`,
				},
				resource.Attribute{
					Name:        "script_id",
					Description: `The ID of the startup script that was added to the server.`,
				},
				resource.Attribute{
					Name:        "firewall_group_id",
					Description: `The ID of the firewall group assigned to the server.`,
				},
				resource.Attribute{
					Name:        "private_network_ids",
					Description: `(Deprecated: Use ` + "`" + `vpc_ids` + "`" + ` instead) A list of private network IDs attached to the server.`,
				},
				resource.Attribute{
					Name:        "vpc_ids",
					Description: `A list of VPC IDs attached to the server.`,
				},
				resource.Attribute{
					Name:        "ssh_key_ids",
					Description: `A list of SSH key IDs applied to the server on install.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `Generic data store, which some provisioning tools and cloud operating systems use as a configuration file. It is generally consumed only once after an instance has been launched, but individual needs may vary.`,
				},
				resource.Attribute{
					Name:        "backups",
					Description: `Whether automatic backups are enabled for this server.`,
				},
				resource.Attribute{
					Name:        "enable_ipv6",
					Description: `Whether the server has IPv6 networking activated.`,
				},
				resource.Attribute{
					Name:        "activation_email",
					Description: `Whether an activation email was sent when the server was ready.`,
				},
				resource.Attribute{
					Name:        "ddos_protection",
					Description: `Whether DDOS protection is enabled on the server.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `The hostname assigned to the server.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `The tag assigned to the server.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `A label for the server.`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `Array of which features are enabled.`,
				},
				resource.Attribute{
					Name:        "backups_schedule",
					Description: `(Optional) A block that defines the way backups should be scheduled. ## Import Servers can be imported using the server ` + "`" + `ID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_server.my_server b6a859c5-b299-49dd-8888-b1abbc517d08 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the server.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The ID of the region that the server is in.`,
				},
				resource.Attribute{
					Name:        "os",
					Description: `The string description of the operating system installed on the server.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `The amount of memory available on the server in MB.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `The description of the disk(s) on the server.`,
				},
				resource.Attribute{
					Name:        "main_ip",
					Description: `The server's main IP address.`,
				},
				resource.Attribute{
					Name:        "vcpu_count",
					Description: `The number of virtual CPUs available on the server.`,
				},
				resource.Attribute{
					Name:        "default_password",
					Description: `The server's default password.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date the server was added to your Vultr account.`,
				},
				resource.Attribute{
					Name:        "allowed_bandwidth",
					Description: `The server's allowed bandwidth usage in GB.`,
				},
				resource.Attribute{
					Name:        "netmask_v4",
					Description: `The server's IPv4 netmask.`,
				},
				resource.Attribute{
					Name:        "gateway_v4",
					Description: `The server's IPv4 gateway.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the server's subscription.`,
				},
				resource.Attribute{
					Name:        "power_status",
					Description: `Whether the server is powered on or not.`,
				},
				resource.Attribute{
					Name:        "server_status",
					Description: `A more detailed server status (none, locked, installingbooting, isomounting, ok).`,
				},
				resource.Attribute{
					Name:        "v6_network",
					Description: `The IPv6 subnet.`,
				},
				resource.Attribute{
					Name:        "v6_main_ip",
					Description: `The main IPv6 network address.`,
				},
				resource.Attribute{
					Name:        "v6_network_size",
					Description: `The IPv6 network size in bits.`,
				},
				resource.Attribute{
					Name:        "internal_ip",
					Description: `The server's internal IP address.`,
				},
				resource.Attribute{
					Name:        "kvm",
					Description: `The server's current KVM URL. This URL will change periodically. It is not advised to cache this value.`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `The ID of the plan that server is subscribed to.`,
				},
				resource.Attribute{
					Name:        "os_id",
					Description: `The ID of the operating system installed on the server.`,
				},
				resource.Attribute{
					Name:        "iso_id",
					Description: `The ID of the ISO file installed on the server.`,
				},
				resource.Attribute{
					Name:        "app_id",
					Description: `The ID of the Vultr application installed on the server.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The ID of the Vultr marketplace application installed on the server.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The ID of the Vultr snapshot that the server was restored from.`,
				},
				resource.Attribute{
					Name:        "script_id",
					Description: `The ID of the startup script that was added to the server.`,
				},
				resource.Attribute{
					Name:        "firewall_group_id",
					Description: `The ID of the firewall group assigned to the server.`,
				},
				resource.Attribute{
					Name:        "private_network_ids",
					Description: `(Deprecated: Use ` + "`" + `vpc_ids` + "`" + ` instead) A list of private network IDs attached to the server.`,
				},
				resource.Attribute{
					Name:        "vpc_ids",
					Description: `A list of VPC IDs attached to the server.`,
				},
				resource.Attribute{
					Name:        "ssh_key_ids",
					Description: `A list of SSH key IDs applied to the server on install.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `Generic data store, which some provisioning tools and cloud operating systems use as a configuration file. It is generally consumed only once after an instance has been launched, but individual needs may vary.`,
				},
				resource.Attribute{
					Name:        "backups",
					Description: `Whether automatic backups are enabled for this server.`,
				},
				resource.Attribute{
					Name:        "enable_ipv6",
					Description: `Whether the server has IPv6 networking activated.`,
				},
				resource.Attribute{
					Name:        "activation_email",
					Description: `Whether an activation email was sent when the server was ready.`,
				},
				resource.Attribute{
					Name:        "ddos_protection",
					Description: `Whether DDOS protection is enabled on the server.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `The hostname assigned to the server.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `The tag assigned to the server.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `A label for the server.`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `Array of which features are enabled.`,
				},
				resource.Attribute{
					Name:        "backups_schedule",
					Description: `(Optional) A block that defines the way backups should be scheduled. ## Import Servers can be imported using the server ` + "`" + `ID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_server.my_server b6a859c5-b299-49dd-8888-b1abbc517d08 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vultr_instance_ipv4",
			Category:         "Resources",
			ShortDescription: `Provides a instance IPv4 resource. This can be used to create, read, and modify a IPv4 address.`,
			Description:      ``,
			Keywords: []string{
				"instance",
				"ipv4",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) The ID of the instance you want to set an IPv4 reverse DNS record for.`,
				},
				resource.Attribute{
					Name:        "reboot",
					Description: `(Optional) Default true. Determines whether or not the server is rebooted after adding the IPv4 address. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID is the IPv4 address in canonical format.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of the server the IPv4 was set for.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `The IPv4 address in canonical format.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `The gateway IP address.`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `The IPv4 netmask in dot-decimal notation.`,
				},
				resource.Attribute{
					Name:        "reverse",
					Description: `The reverse DNS information for this IP address.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID is the IPv4 address in canonical format.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of the server the IPv4 was set for.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `The IPv4 address in canonical format.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `The gateway IP address.`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `The IPv4 netmask in dot-decimal notation.`,
				},
				resource.Attribute{
					Name:        "reverse",
					Description: `The reverse DNS information for this IP address.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vultr_iso",
			Category:         "Resources",
			ShortDescription: `Provides a Vultr ISO file resource. This can be used to create, read, and delete ISO files on your Vultr account.`,
			Description:      ``,
			Keywords: []string{
				"iso",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "url",
					Description: `(Required) URL pointing to the ISO file. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the ISO.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `URL pointing to the ISO file.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date the ISO was created.`,
				},
				resource.Attribute{
					Name:        "filename",
					Description: `The ISO filename.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The ISO size in bytes.`,
				},
				resource.Attribute{
					Name:        "md5sum",
					Description: `The md5 hash of the ISO file.`,
				},
				resource.Attribute{
					Name:        "sha512sum",
					Description: `The sha512 hash of the ISO file.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the ISO file. ## Import ISOs can be imported using the ISO ` + "`" + `ID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_iso_private.my_iso eb807c23-c311-4d35-a77f-e49bfeb2dec5 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the ISO.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `URL pointing to the ISO file.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date the ISO was created.`,
				},
				resource.Attribute{
					Name:        "filename",
					Description: `The ISO filename.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The ISO size in bytes.`,
				},
				resource.Attribute{
					Name:        "md5sum",
					Description: `The md5 hash of the ISO file.`,
				},
				resource.Attribute{
					Name:        "sha512sum",
					Description: `The sha512 hash of the ISO file.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the ISO file. ## Import ISOs can be imported using the ISO ` + "`" + `ID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_iso_private.my_iso eb807c23-c311-4d35-a77f-e49bfeb2dec5 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vultr_kubernetes",
			Category:         "Resources",
			ShortDescription: `Provides a Vultr Kubernetes Engine (VKE) resource. This can be used to create, read, modify, and delete VKE clusters on your Vultr account.`,
			Description:      ``,
			Keywords: []string{
				"kubernetes",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The region your VKE cluster will be deployed in. Currently, supported values are ` + "`" + `ewr` + "`" + ` and ` + "`" + `lax` + "`" + ``,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The version your VKE cluster you want deployed. [See Available Version](https://www.vultr.com/api/#operation/get-kubernetes-versions)`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Optional) The VKE clusters label. ` + "`" + `node_pools` + "`" + ` (Required) There must be 1 node pool with the kubernetes resource. It supports the following fields`,
				},
				resource.Attribute{
					Name:        "node_quantity",
					Description: `(Required) The number of nodes in this node pool.`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `(Required) The plan to be used in this node pool. [See Plans List](https://www.vultr.com/api/#operation/list-plans) Note the minimum plan requirements must have at least 1 core and 2 gbs of memory.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Required) The label to be used as a prefix for nodes in this node pool. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The VKE cluster ID.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The VKE clusters label.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region your VKE cluster is deployed in.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The current kubernetes version your VKE cluster is running on.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The overall status of the cluster.`,
				},
				resource.Attribute{
					Name:        "service_subnet",
					Description: `IP range that services will run on this cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_subnet",
					Description: `IP range that your pods will run on in this cluster.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `Domain for your Kubernetes clusters control plane.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `IP address of VKE cluster control plane.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `Date of VKE cluster creation.`,
				},
				resource.Attribute{
					Name:        "kube_config",
					Description: `Base64 encoded Kubeconfig for this VKE cluster.`,
				},
				resource.Attribute{
					Name:        "node_pools",
					Description: `Contains the default node pool that was deployed. ` + "`" + `node_pools` + "`" + ``,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `Date of node pool creation.`,
				},
				resource.Attribute{
					Name:        "date_updated",
					Description: `Date of node pool updates.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `Label of node pool.`,
				},
				resource.Attribute{
					Name:        "node_quantity",
					Description: `Number of nodes within node pool.`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `Node plan that nodes are using within this node pool.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of node pool.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `Tag for node pool.`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `Array that contains information about nodes within this node pool. ` + "`" + `nodes` + "`" + ``,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `Date node was created.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of node.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `Label of node.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of node.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The VKE cluster ID.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The VKE clusters label.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region your VKE cluster is deployed in.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The current kubernetes version your VKE cluster is running on.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The overall status of the cluster.`,
				},
				resource.Attribute{
					Name:        "service_subnet",
					Description: `IP range that services will run on this cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_subnet",
					Description: `IP range that your pods will run on in this cluster.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `Domain for your Kubernetes clusters control plane.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `IP address of VKE cluster control plane.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `Date of VKE cluster creation.`,
				},
				resource.Attribute{
					Name:        "kube_config",
					Description: `Base64 encoded Kubeconfig for this VKE cluster.`,
				},
				resource.Attribute{
					Name:        "node_pools",
					Description: `Contains the default node pool that was deployed. ` + "`" + `node_pools` + "`" + ``,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `Date of node pool creation.`,
				},
				resource.Attribute{
					Name:        "date_updated",
					Description: `Date of node pool updates.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `Label of node pool.`,
				},
				resource.Attribute{
					Name:        "node_quantity",
					Description: `Number of nodes within node pool.`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `Node plan that nodes are using within this node pool.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of node pool.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `Tag for node pool.`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `Array that contains information about nodes within this node pool. ` + "`" + `nodes` + "`" + ``,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `Date node was created.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of node.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `Label of node.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of node.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vultr_load_balancer",
			Category:         "Resources",
			ShortDescription: `Get information about a Vultr Load Balancer.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The region your load balancer is deployed in.`,
				},
				resource.Attribute{
					Name:        "forwarding_rules",
					Description: `(Required) List of forwarding rules for a load balancer. The configuration of a ` + "`" + `forwarding_rules` + "`" + ` is listened below.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Optional) The load balancer's label.`,
				},
				resource.Attribute{
					Name:        "balancing_algorithm",
					Description: `(Optional) The balancing algorithm for your load balancer. Options are ` + "`" + `roundrobin` + "`" + ` or ` + "`" + `leastconn` + "`" + `. Default value is ` + "`" + `roundrobin` + "`" + ``,
				},
				resource.Attribute{
					Name:        "proxy_protocol",
					Description: `(Optional) Boolean value that indicates if Proxy Protocol is enabled.`,
				},
				resource.Attribute{
					Name:        "cookie_name",
					Description: `(Optional) Name for your given sticky session.`,
				},
				resource.Attribute{
					Name:        "ssl_redirect",
					Description: `(Optional) Boolean value that indicates if HTTP calls will be redirected to HTTPS.`,
				},
				resource.Attribute{
					Name:        "attached_instances",
					Description: `(Optional) Array of instances that are currently attached to the load balancer.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `(Optional) A block that defines the way load balancers should check for health. The configuration of a ` + "`" + `health_check` + "`" + ` is listed below.`,
				},
				resource.Attribute{
					Name:        "ssl",
					Description: `(Optional) A block that supplies your ssl configuration to be used with HTTPS. The configuration of a ` + "`" + `ssl` + "`" + ` is listed below.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) The protocol used to traffic requests to the load balancer. Possible values are ` + "`" + `http` + "`" + `, or ` + "`" + `tcp` + "`" + `. Default value is ` + "`" + `http` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) The path on the attached instances that the load balancer should check against. Default value is ` + "`" + `/` + "`" + ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The assigned port (integer) on the attached instances that the load balancer should check against. Default value is ` + "`" + `80` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "check_interval",
					Description: `(Optional) Time in seconds to perform health check. Default value is 15.`,
				},
				resource.Attribute{
					Name:        "response_timeout",
					Description: `(Optional) Time in seconds to wait for a health check response. Default value is 5.`,
				},
				resource.Attribute{
					Name:        "unhealthy_threshold",
					Description: `(Optional) Number of failed attempts encountered before failover. Default value is 5.`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `(Optional) Number of failed attempts encountered before failover. Default value is 5. ` + "`" + `forwarding_rules` + "`" + ` supports the following`,
				},
				resource.Attribute{
					Name:        "frontend_protocol",
					Description: `(Required) Protocol on load balancer side. Possible values: "http", "https", "tcp".`,
				},
				resource.Attribute{
					Name:        "frontend_port",
					Description: `(Required) Port on load balancer side.`,
				},
				resource.Attribute{
					Name:        "backend_protocol",
					Description: `(Required) Protocol on instance side. Possible values: "http", "https", "tcp".`,
				},
				resource.Attribute{
					Name:        "backend_port",
					Description: `(Required) Port on instance side. ` + "`" + `ssl` + "`" + ` supports the following`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Required) The SSL certificates private key.`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(Required) The SSL Certificate.`,
				},
				resource.Attribute{
					Name:        "chain",
					Description: `(Optional) The SSL certificate chain. ` + "`" + `firewall_rules` + "`" + ` supports the following`,
				},
				resource.Attribute{
					Name:        "frontend_port",
					Description: `(Required) Port on load balancer side.`,
				},
				resource.Attribute{
					Name:        "ip_type",
					Description: `(Required) The type of ip this rule is - may be either v4 or v6.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Required) IP address with subnet that is allowed through the firewall. You may also pass in ` + "`" + `cloudflare` + "`" + ` which will allow only CloudFlares IP range. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The load balancer ID.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region your load balancer is deployed in.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The load balancer's label.`,
				},
				resource.Attribute{
					Name:        "balancing_algorithm",
					Description: `The balancing algorithm for your load balancer.`,
				},
				resource.Attribute{
					Name:        "proxy_protocol",
					Description: `Boolean value that indicates if Proxy Protocol is enabled.`,
				},
				resource.Attribute{
					Name:        "cookie_name",
					Description: `Name for your given sticky session.`,
				},
				resource.Attribute{
					Name:        "ssl_redirect",
					Description: `Boolean value that indicates if HTTP calls will be redirected to HTTPS.`,
				},
				resource.Attribute{
					Name:        "has_ssl",
					Description: `Boolean value that indicates if SSL is enabled.`,
				},
				resource.Attribute{
					Name:        "attached_instances",
					Description: `Array of instances that are currently attached to the load balancer.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status for the load balancer`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `IPv4 address for your load balancer.`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `IPv6 address for your load balancer.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `Defines the way load balancers should check for health.`,
				},
				resource.Attribute{
					Name:        "forwarding_rules",
					Description: `Defines the forwarding rules for a load balancer.`,
				},
				resource.Attribute{
					Name:        "firewall_rules",
					Description: `Defines the firewall rules for a load balancer.`,
				},
				resource.Attribute{
					Name:        "private_network",
					Description: `(Deprecated: use ` + "`" + `vpc` + "`" + ` instead) Defines the private network the load balancer is attached to.`,
				},
				resource.Attribute{
					Name:        "vpc",
					Description: `Defines the VPC the load balancer is attached to. ## Import Load Balancers can be imported using the load balancer ` + "`" + `ID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_load_balancer.lb b6a859c5-b299-49dd-8888-b1abbc517d08 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The load balancer ID.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region your load balancer is deployed in.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The load balancer's label.`,
				},
				resource.Attribute{
					Name:        "balancing_algorithm",
					Description: `The balancing algorithm for your load balancer.`,
				},
				resource.Attribute{
					Name:        "proxy_protocol",
					Description: `Boolean value that indicates if Proxy Protocol is enabled.`,
				},
				resource.Attribute{
					Name:        "cookie_name",
					Description: `Name for your given sticky session.`,
				},
				resource.Attribute{
					Name:        "ssl_redirect",
					Description: `Boolean value that indicates if HTTP calls will be redirected to HTTPS.`,
				},
				resource.Attribute{
					Name:        "has_ssl",
					Description: `Boolean value that indicates if SSL is enabled.`,
				},
				resource.Attribute{
					Name:        "attached_instances",
					Description: `Array of instances that are currently attached to the load balancer.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status for the load balancer`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `IPv4 address for your load balancer.`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `IPv6 address for your load balancer.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `Defines the way load balancers should check for health.`,
				},
				resource.Attribute{
					Name:        "forwarding_rules",
					Description: `Defines the forwarding rules for a load balancer.`,
				},
				resource.Attribute{
					Name:        "firewall_rules",
					Description: `Defines the firewall rules for a load balancer.`,
				},
				resource.Attribute{
					Name:        "private_network",
					Description: `(Deprecated: use ` + "`" + `vpc` + "`" + ` instead) Defines the private network the load balancer is attached to.`,
				},
				resource.Attribute{
					Name:        "vpc",
					Description: `Defines the VPC the load balancer is attached to. ## Import Load Balancers can be imported using the load balancer ` + "`" + `ID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_load_balancer.lb b6a859c5-b299-49dd-8888-b1abbc517d08 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vultr_object_storage",
			Category:         "Resources",
			ShortDescription: `Provides a Vultr private object storage resource. This can be used to create, read, update and delete object storage resources on your Vultr account.`,
			Description:      ``,
			Keywords: []string{
				"object",
				"storage",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) The region ID that you want the network to be created in.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Optional) The description you want to give your network. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the object storage subscription.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The label of the object storage subscription.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location which this subscription resides in.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `The identifying cluster ID.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region ID of the object storage subscription.`,
				},
				resource.Attribute{
					Name:        "s3_access_key",
					Description: `Your access key.`,
				},
				resource.Attribute{
					Name:        "s3_hostname",
					Description: `The hostname for this subscription.`,
				},
				resource.Attribute{
					Name:        "s3_secret_key",
					Description: `Your secret key.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of this object storage subscription.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `Date of creation for the object storage subscription. ## Import Object Storage can be imported using the object storage ` + "`" + `ID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_object_storage.my_s3 0e04f918-575e-41cb-86f6-d729b354a5a1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the object storage subscription.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The label of the object storage subscription.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location which this subscription resides in.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `The identifying cluster ID.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region ID of the object storage subscription.`,
				},
				resource.Attribute{
					Name:        "s3_access_key",
					Description: `Your access key.`,
				},
				resource.Attribute{
					Name:        "s3_hostname",
					Description: `The hostname for this subscription.`,
				},
				resource.Attribute{
					Name:        "s3_secret_key",
					Description: `Your secret key.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of this object storage subscription.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `Date of creation for the object storage subscription. ## Import Object Storage can be imported using the object storage ` + "`" + `ID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_object_storage.my_s3 0e04f918-575e-41cb-86f6-d729b354a5a1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vultr_private_network",
			Category:         "Resources",
			ShortDescription: `Provides a Vultr private network resource. This can be used to create, read, and delete private networks on your Vultr account.`,
			Description:      ``,
			Keywords: []string{
				"private",
				"network",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The region ID that you want the network to be created in.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description you want to give your network.`,
				},
				resource.Attribute{
					Name:        "v4_subnet",
					Description: `(Optional) The IPv4 subnet to be used when attaching instances to this network.`,
				},
				resource.Attribute{
					Name:        "v4_subnet_mask",
					Description: `The number of bits for the netmask in CIDR notation. Example: 32 ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the network.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region ID that the network operates in.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the network.`,
				},
				resource.Attribute{
					Name:        "v4_subnet",
					Description: `The IPv4 subnet used when attaching instances to this network.`,
				},
				resource.Attribute{
					Name:        "v4_subnet_mask",
					Description: `The number of bits for the netmask in CIDR notation. Example: 32`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date that the network was added to your Vultr account. ## Import Networks can be imported using the network ` + "`" + `ID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_private_network.my_network 0e04f918-575e-41cb-86f6-d729b354a5a1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the network.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region ID that the network operates in.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the network.`,
				},
				resource.Attribute{
					Name:        "v4_subnet",
					Description: `The IPv4 subnet used when attaching instances to this network.`,
				},
				resource.Attribute{
					Name:        "v4_subnet_mask",
					Description: `The number of bits for the netmask in CIDR notation. Example: 32`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date that the network was added to your Vultr account. ## Import Networks can be imported using the network ` + "`" + `ID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_private_network.my_network 0e04f918-575e-41cb-86f6-d729b354a5a1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vultr_reserved_ip",
			Category:         "Resources",
			ShortDescription: `Provides a Vultr reserved IP resource. This can be used to create, read, modify, and delete reserved IP addresses on your Vultr account.`,
			Description:      ``,
			Keywords: []string{
				"reserved",
				"ip",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The region ID that you want the reserved IP to be created in.`,
				},
				resource.Attribute{
					Name:        "ip_type",
					Description: `(Required) The type of reserved IP that you want. Either "v4" or "v6".`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Optional) The label you want to give your reserved IP.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Optional) The VPS ID you want this reserved IP to be attached to. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the reserved IP.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region ID that this reserved IP belongs to.`,
				},
				resource.Attribute{
					Name:        "ip_type",
					Description: `The reserved IP's type.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The reserved IP's label.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of the instance the reserved IP is attached to.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `The reserved IP's subnet.`,
				},
				resource.Attribute{
					Name:        "subnet_size",
					Description: `The reserved IP's subnet size. ## Import Reserved IPs can be imported using the reserved IP ` + "`" + `ID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_reserved_ip.my_reserved_ip b9cc6fad-70b1-40ee-ab6a-4d622858962f ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the reserved IP.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region ID that this reserved IP belongs to.`,
				},
				resource.Attribute{
					Name:        "ip_type",
					Description: `The reserved IP's type.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The reserved IP's label.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of the instance the reserved IP is attached to.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `The reserved IP's subnet.`,
				},
				resource.Attribute{
					Name:        "subnet_size",
					Description: `The reserved IP's subnet size. ## Import Reserved IPs can be imported using the reserved IP ` + "`" + `ID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_reserved_ip.my_reserved_ip b9cc6fad-70b1-40ee-ab6a-4d622858962f ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vultr_reverse_ipv4",
			Category:         "Resources",
			ShortDescription: `Provides a Vultr Reverse IPv4 resource. This can be used to create, read, and modify reverse DNS records for IPv4 addresses.`,
			Description:      ``,
			Keywords: []string{
				"reverse",
				"ipv4",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) The ID of the instance you want to set an IPv4 reverse DNS record for.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) The IPv4 address used in the reverse DNS record.`,
				},
				resource.Attribute{
					Name:        "reverse",
					Description: `(Required) The hostname used in the IPv4 reverse DNS record. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID is the IPv4 address in canonical format.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of the instance the IPv4 reverse DNS record was set for.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `The IPv4 address in canonical format used in the reverse DNS record.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `The gateway IP address.`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `The IPv4 netmask in dot-decimal notation.`,
				},
				resource.Attribute{
					Name:        "reverse",
					Description: `The reverse DNS information for this IP address.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID is the IPv4 address in canonical format.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of the instance the IPv4 reverse DNS record was set for.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `The IPv4 address in canonical format used in the reverse DNS record.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `The gateway IP address.`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `The IPv4 netmask in dot-decimal notation.`,
				},
				resource.Attribute{
					Name:        "reverse",
					Description: `The reverse DNS information for this IP address.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vultr_reverse_ipv6",
			Category:         "Resources",
			ShortDescription: `Provides a Vultr Reverse IPv6 resource. This can be used to create, read, modify, and delete reverse DNS records for IPv6 addresses.`,
			Description:      ``,
			Keywords: []string{
				"reverse",
				"ipv6",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) The ID of the server you want to set an IPv6 reverse DNS record for.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) The IPv6 address used in the reverse DNS record.`,
				},
				resource.Attribute{
					Name:        "reverse",
					Description: `(Required) The hostname used in the IPv6 reverse DNS record. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID is the IPv6 address in canonical format.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of the server the IPv6 reverse DNS record was set for.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `The IPv6 address in canonical format used in the reverse DNS record.`,
				},
				resource.Attribute{
					Name:        "reverse",
					Description: `The hostname used in the IPv6 reverse DNS record.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID is the IPv6 address in canonical format.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of the server the IPv6 reverse DNS record was set for.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `The IPv6 address in canonical format used in the reverse DNS record.`,
				},
				resource.Attribute{
					Name:        "reverse",
					Description: `The hostname used in the IPv6 reverse DNS record.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vultr_snapshot",
			Category:         "Resources",
			ShortDescription: `Provides a Vultr Snapshot resource. This can be used to create, read, modify, and delete Snapshot.`,
			Description:      ``,
			Keywords: []string{
				"snapshot",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) ID of a given instance that you want to create a snapshot from.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description for the given snapshot. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID for the given snapshot.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of the instance that the snapshot was created from.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description for the given snapshot.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date the snapshot was created.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the snapshot in Bytes.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status for the given snapshot.`,
				},
				resource.Attribute{
					Name:        "os_id",
					Description: `The os id which the snapshot is associated with.`,
				},
				resource.Attribute{
					Name:        "app_id",
					Description: `The app id which the snapshot is associated with. ## Import Snapshots can be imported using the Snapshot ` + "`" + `ID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_snapshot_url.my_snapshot 283941e8-0783-410e-9540-71c86b833992 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID for the given snapshot.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of the instance that the snapshot was created from.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description for the given snapshot.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date the snapshot was created.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the snapshot in Bytes.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status for the given snapshot.`,
				},
				resource.Attribute{
					Name:        "os_id",
					Description: `The os id which the snapshot is associated with.`,
				},
				resource.Attribute{
					Name:        "app_id",
					Description: `The app id which the snapshot is associated with. ## Import Snapshots can be imported using the Snapshot ` + "`" + `ID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_snapshot_url.my_snapshot 283941e8-0783-410e-9540-71c86b833992 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vultr_snapshot_from_url",
			Category:         "Resources",
			ShortDescription: `Provides a Vultr Snapshots from URL resource. This can be used to create, read, modify, and delete Snapshots from URL.`,
			Description:      ``,
			Keywords: []string{
				"snapshot",
				"from",
				"url",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "url",
					Description: `(Required) URL of the given resource you want to create a snapshot from. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID for the given snapshot.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description for the given snapshot.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `The url from where the raw image was used to create the snapshot.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date the snapshot was created.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the snapshot in Bytes.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status for the given snapshot.`,
				},
				resource.Attribute{
					Name:        "os_id",
					Description: `The os id which the snapshot is associated with.`,
				},
				resource.Attribute{
					Name:        "app_id",
					Description: `The app id which the snapshot is associated with. ## Import Snapshots from Url can be imported using the Snapshot ` + "`" + `ID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_snapshot_from_url.my_snapshot e60dc0a2-9313-4bab-bffc-57ffe33d99f6 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID for the given snapshot.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description for the given snapshot.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `The url from where the raw image was used to create the snapshot.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date the snapshot was created.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the snapshot in Bytes.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status for the given snapshot.`,
				},
				resource.Attribute{
					Name:        "os_id",
					Description: `The os id which the snapshot is associated with.`,
				},
				resource.Attribute{
					Name:        "app_id",
					Description: `The app id which the snapshot is associated with. ## Import Snapshots from Url can be imported using the Snapshot ` + "`" + `ID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_snapshot_from_url.my_snapshot e60dc0a2-9313-4bab-bffc-57ffe33d99f6 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vultr_ssh_key",
			Category:         "Resources",
			ShortDescription: `Provides a Vultr SSH key resource. This can be used to create, read, modify, and delete SSH keys.`,
			Description:      ``,
			Keywords: []string{
				"ssh",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name/label of the SSH key.`,
				},
				resource.Attribute{
					Name:        "ssh_key",
					Description: `(Required) The public SSH key. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the SSH key.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name/label of the SSH key.`,
				},
				resource.Attribute{
					Name:        "ssh_key",
					Description: `The public SSH key.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date the SSH key was added to your Vultr account. ## Import SSH keys can be imported using the SSH key ` + "`" + `ID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_ssh_key.my_key 6b0876a7-f709-41ba-aed8-abed9d38ae45 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the SSH key.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name/label of the SSH key.`,
				},
				resource.Attribute{
					Name:        "ssh_key",
					Description: `The public SSH key.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date the SSH key was added to your Vultr account. ## Import SSH keys can be imported using the SSH key ` + "`" + `ID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_ssh_key.my_key 6b0876a7-f709-41ba-aed8-abed9d38ae45 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vultr_startup_script",
			Category:         "Resources",
			ShortDescription: `Provides a Vultr Startup Script resource. This can be used to create, read, modify, and delete Startup Scripts.`,
			Description:      ``,
			Keywords: []string{
				"startup",
				"script",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the given script.`,
				},
				resource.Attribute{
					Name:        "script",
					Description: `(Required) Contents of the startup script base64 encoded.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type of startup script. Possible values are boot or pxe - default is boot. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the script.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the given script.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `Date the script was created.`,
				},
				resource.Attribute{
					Name:        "date_modified",
					Description: `Date the script was last modified.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of startup script this is.`,
				},
				resource.Attribute{
					Name:        "script",
					Description: `The contents of the startup script base64 encoded. ## Import Startup Scripts can be imported using the Startup Scripts ` + "`" + `ID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_startup_script.my_script ff8f36a8-eb86-4b8d-8667-b9d5459b6390 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the script.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the given script.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `Date the script was created.`,
				},
				resource.Attribute{
					Name:        "date_modified",
					Description: `Date the script was last modified.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of startup script this is.`,
				},
				resource.Attribute{
					Name:        "script",
					Description: `The contents of the startup script base64 encoded. ## Import Startup Scripts can be imported using the Startup Scripts ` + "`" + `ID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_startup_script.my_script ff8f36a8-eb86-4b8d-8667-b9d5459b6390 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vultr_user",
			Category:         "Resources",
			ShortDescription: `Provides a Vultr User resource. This can be used to create, read, modify, and delete Users.`,
			Description:      ``,
			Keywords: []string{
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name for this user.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Required) Email for this user.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Password for this user.`,
				},
				resource.Attribute{
					Name:        "api_enabled",
					Description: `(Optional) Whether API is enabled for the user. Default behavior is set to enabled.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `(Optional) The access control list for the user. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID associated with the user.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name for this user.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `Email for this user.`,
				},
				resource.Attribute{
					Name:        "api_enabled",
					Description: `Whether API is enabled for the user. ## Import Users can be imported using the User ` + "`" + `ID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_user.myuser 1345fef0-8ed3-4a66-bd8c-822a7b7bd05a ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID associated with the user.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name for this user.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `Email for this user.`,
				},
				resource.Attribute{
					Name:        "api_enabled",
					Description: `Whether API is enabled for the user. ## Import Users can be imported using the User ` + "`" + `ID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_user.myuser 1345fef0-8ed3-4a66-bd8c-822a7b7bd05a ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vultr_vpc",
			Category:         "Resources",
			ShortDescription: `Provides a Vultr VPC resource. This can be used to create, read, and delete VPCs on your Vultr account.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The region ID that you want the VPC to be created in.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description you want to give your VPC.`,
				},
				resource.Attribute{
					Name:        "v4_subnet",
					Description: `(Optional) The IPv4 subnet to be used when attaching instances to this VPC.`,
				},
				resource.Attribute{
					Name:        "v4_subnet_mask",
					Description: `The number of bits for the netmask in CIDR notation. Example: 32 ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the VPC.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region ID that the VPC operates in.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the VPC.`,
				},
				resource.Attribute{
					Name:        "v4_subnet",
					Description: `The IPv4 subnet used when attaching instances to this VPC.`,
				},
				resource.Attribute{
					Name:        "v4_subnet_mask",
					Description: `The number of bits for the netmask in CIDR notation. Example: 32`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date that the VPC was added to your Vultr account. ## Import VPCs can be imported using the VPC ` + "`" + `ID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_vpc.my_vpc 0e04f918-575e-41cb-86f6-d729b354a5a1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the VPC.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region ID that the VPC operates in.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the VPC.`,
				},
				resource.Attribute{
					Name:        "v4_subnet",
					Description: `The IPv4 subnet used when attaching instances to this VPC.`,
				},
				resource.Attribute{
					Name:        "v4_subnet_mask",
					Description: `The number of bits for the netmask in CIDR notation. Example: 32`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date that the VPC was added to your Vultr account. ## Import VPCs can be imported using the VPC ` + "`" + `ID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_vpc.my_vpc 0e04f918-575e-41cb-86f6-d729b354a5a1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"vultr_bare_metal_server": 0,
		"vultr_block_storage":     1,
		"vultr_dns_domain":        2,
		"vultr_dns_record":        3,
		"vultr_firewall_group":    4,
		"vultr_firewall_rule":     5,
		"vultr_instance":          6,
		"vultr_instance_ipv4":     7,
		"vultr_iso":               8,
		"vultr_kubernetes":        9,
		"vultr_load_balancer":     10,
		"vultr_object_storage":    11,
		"vultr_private_network":   12,
		"vultr_reserved_ip":       13,
		"vultr_reverse_ipv4":      14,
		"vultr_reverse_ipv6":      15,
		"vultr_snapshot":          16,
		"vultr_snapshot_from_url": 17,
		"vultr_ssh_key":           18,
		"vultr_startup_script":    19,
		"vultr_user":              20,
		"vultr_vpc":               21,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
