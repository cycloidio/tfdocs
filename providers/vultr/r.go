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
					Name:        "region_id",
					Description: `(Required) The ID of the region that the server is to be created in.`,
				},
				resource.Attribute{
					Name:        "plan_id",
					Description: `(Required) The ID of the plan that you want the server to subscribe to.`,
				},
				resource.Attribute{
					Name:        "os_id",
					Description: `(Optional) The ID of the operating system to be installed on the server.`,
				},
				resource.Attribute{
					Name:        "app_id",
					Description: `(Optional) The ID of the Vultr application to be installed on the server.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) The ID of the Vultr snapshot that the server will restore for the initial installation.`,
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
					Name:        "notify_activate",
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
					Description: `(Optional) A label for the server. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the server.`,
				},
				resource.Attribute{
					Name:        "region_id",
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
					Name:        "location",
					Description: `The physical location of the server.`,
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
					Name:        "v6_networks",
					Description: `A list of the server's IPv6 networks.`,
				},
				resource.Attribute{
					Name:        "plan_id",
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
					Description: `Base64 encoded generic data store, which some provisioning tools and cloud operating systems use as a configuration file. It is generally consumed only once after an instance has been launched, but individual needs may vary.`,
				},
				resource.Attribute{
					Name:        "enable_ipv6",
					Description: `Whether the server has IPv6 networking activated.`,
				},
				resource.Attribute{
					Name:        "notify_activate",
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
					Description: `A label for the server. ## Import Bare Metal Servers can be imported using the server ` + "`" + `SUBID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_bare_metal_server.my_server 1312965 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the server.`,
				},
				resource.Attribute{
					Name:        "region_id",
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
					Name:        "location",
					Description: `The physical location of the server.`,
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
					Name:        "v6_networks",
					Description: `A list of the server's IPv6 networks.`,
				},
				resource.Attribute{
					Name:        "plan_id",
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
					Description: `Base64 encoded generic data store, which some provisioning tools and cloud operating systems use as a configuration file. It is generally consumed only once after an instance has been launched, but individual needs may vary.`,
				},
				resource.Attribute{
					Name:        "enable_ipv6",
					Description: `Whether the server has IPv6 networking activated.`,
				},
				resource.Attribute{
					Name:        "notify_activate",
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
					Description: `A label for the server. ## Import Bare Metal Servers can be imported using the server ` + "`" + `SUBID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_bare_metal_server.my_server 1312965 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "region_id",
					Description: `(Required) Region in which this block storage will reside in. (Currently only NJ/NY supported region_id 1)`,
				},
				resource.Attribute{
					Name:        "attached_id",
					Description: `(Optional) VPS ID that you want to have this block storage attached to.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Optional) Label that is given to your block storage. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "size_gb",
					Description: `The size of the given block storage.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `Region in which this block storage will reside in. (Currently only NJ/NY supported region_id 1)`,
				},
				resource.Attribute{
					Name:        "attached_id",
					Description: `VPS ID that is attached to this block storage.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `Label that is given to your block storage.`,
				},
				resource.Attribute{
					Name:        "cost_per_month",
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
					Name:        "id",
					Description: `The ID for this block storage. ## Import Block Storage can be imported using the Block Storage ` + "`" + `SUBID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_block_storage.my_blockstorage 25058682 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "size_gb",
					Description: `The size of the given block storage.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `Region in which this block storage will reside in. (Currently only NJ/NY supported region_id 1)`,
				},
				resource.Attribute{
					Name:        "attached_id",
					Description: `VPS ID that is attached to this block storage.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `Label that is given to your block storage.`,
				},
				resource.Attribute{
					Name:        "cost_per_month",
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
					Name:        "id",
					Description: `The ID for this block storage. ## Import Block Storage can be imported using the Block Storage ` + "`" + `SUBID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_block_storage.my_blockstorage 25058682 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "server_ip",
					Description: `(Required) Server IP you want associated to domain. ## Attributes Reference The following attributes are exported:`,
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
					Name:        "server_ip",
					Description: `Server IP you want associated to domain. ## Import DNS Domains can be imported using the Dns Domain ` + "`" + `domain` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_dns_domain.name domain.com ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "server_ip",
					Description: `Server IP you want associated to domain. ## Import DNS Domains can be imported using the Dns Domain ` + "`" + `domain` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_dns_domain.name domain.com ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Required) IP Address of the server the domain is associated with.`,
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
					Description: `IP Address of the server the domain is associated with.`,
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
					Description: `The time to live of this record.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID associated with the record.`,
				},
				resource.Attribute{
					Name:        "data",
					Description: `IP Address of the server the domain is associated with.`,
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
					Description: `The time to live of this record.`,
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
					Description: `The number of servers that are currently using this firewall group.`,
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
					Description: `The number of servers that are currently using this firewall group.`,
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
					Description: `(Required) The type of protocol for this firewall rule. Possible values (icmp, tcp, udp, gre)`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Required) IP address that you want to define for this firewall rule.`,
				},
				resource.Attribute{
					Name:        "from_port",
					Description: `(Optional) Port that you want to define for this rule.`,
				},
				resource.Attribute{
					Name:        "to_port",
					Description: `(Optional) This can be used with the from port if you want to define multiple ports. Example from port 8085 to port 8090`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) A simple note for a given firewall rule ## Attributes Reference The following attributes are exported:`,
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
					Description: `The type of protocol for this firewall rule. Possible values (icmp, tcp, udp, gre)`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `IP address that is defined for this rule.`,
				},
				resource.Attribute{
					Name:        "from_port",
					Description: `Port that is defined for this rule.`,
				},
				resource.Attribute{
					Name:        "to_port",
					Description: `This can be used with the from port if you want to define multiple ports. Example from port 8085 to port 8090`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `A simple note for a given firewall rule`,
				},
				resource.Attribute{
					Name:        "ip_type",
					Description: `The type of ip this rule is - may be either v4 or v6.`,
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
					Description: `The type of protocol for this firewall rule. Possible values (icmp, tcp, udp, gre)`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `IP address that is defined for this rule.`,
				},
				resource.Attribute{
					Name:        "from_port",
					Description: `Port that is defined for this rule.`,
				},
				resource.Attribute{
					Name:        "to_port",
					Description: `This can be used with the from port if you want to define multiple ports. Example from port 8085 to port 8090`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `A simple note for a given firewall rule`,
				},
				resource.Attribute{
					Name:        "ip_type",
					Description: `The type of ip this rule is - may be either v4 or v6.`,
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
					Description: `The status of the ISO file. ## Import ISOs can be imported using the ISO ` + "`" + `ISOID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_iso_private.my_iso 2349859 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `The status of the ISO file. ## Import ISOs can be imported using the ISO ` + "`" + `ISOID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_iso_private.my_iso 2349859 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vultr_network",
			Category:         "Resources",
			ShortDescription: `Provides a Vultr private network resource. This can be used to create, read, and delete private networks on your Vultr account.`,
			Description:      ``,
			Keywords: []string{
				"network",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region_id",
					Description: `(Required) The region ID that you want the network to be created in.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description you want to give your network.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Optional) The IPv4 subnet and subnet mask to be used when attaching servers to this network. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the network.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `The region ID that the network operates in.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the network.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The IPv4 subnet and subnet mask to be used when attaching servers to this network.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date that the network was added to your Vultr account. ## Import Networks can be imported using the network ` + "`" + `NETWORKID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_network.my_network net539626f0798d7 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the network.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `The region ID that the network operates in.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the network.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The IPv4 subnet and subnet mask to be used when attaching servers to this network.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date that the network was added to your Vultr account. ## Import Networks can be imported using the network ` + "`" + `NETWORKID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_network.my_network net539626f0798d7 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "region_id",
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
					Name:        "attached_id",
					Description: `(Optional) The VPS ID you want this reserved IP to be attached to. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the reserved IP.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `The region ID (` + "`" + `DCID` + "`" + ` in the Vultr API) that this reserved IP belongs to.`,
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
					Name:        "attached_id",
					Description: `The ID of the server the reserved IP is attached to.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `The reserved IP's subnet.`,
				},
				resource.Attribute{
					Name:        "subnet_size",
					Description: `The reserved IP's subnet size. ## Import Reserved IPs can be imported using the reserved IP ` + "`" + `SUBID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_reserved_ip.my_reserved_ip 1313044 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the reserved IP.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `The region ID (` + "`" + `DCID` + "`" + ` in the Vultr API) that this reserved IP belongs to.`,
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
					Name:        "attached_id",
					Description: `The ID of the server the reserved IP is attached to.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `The reserved IP's subnet.`,
				},
				resource.Attribute{
					Name:        "subnet_size",
					Description: `The reserved IP's subnet size. ## Import Reserved IPs can be imported using the reserved IP ` + "`" + `SUBID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_reserved_ip.my_reserved_ip 1313044 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vultr_server",
			Category:         "Resources",
			ShortDescription: `Provides a Vultr server resource. This can be used to create, read, modify, and delete servers on your Vultr account.`,
			Description:      ``,
			Keywords: []string{
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region_id",
					Description: `(Required) The ID of the region that the server is to be created in.`,
				},
				resource.Attribute{
					Name:        "plan_id",
					Description: `(Required) The ID of the plan that you want the server to subscribe to.`,
				},
				resource.Attribute{
					Name:        "os_id",
					Description: `(Optional) The ID of the operating system to be installed on the server.`,
				},
				resource.Attribute{
					Name:        "iso_id",
					Description: `(Optional) The ID of the ISO file to be installed on the server.`,
				},
				resource.Attribute{
					Name:        "app_id",
					Description: `(Optional) The ID of the Vultr application to be installed on the server.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) The ID of the Vultr snapshot that the server will restore for the initial installation.`,
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
					Name:        "network_ids",
					Description: `(Optional) A list of private network IDs to be attached to the server.`,
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
					Name:        "auto_backup",
					Description: `(Optional) Whether automatic backups will be enabled for this server (these have an extra charge associated with them).`,
				},
				resource.Attribute{
					Name:        "enable_ipv6",
					Description: `(Optional) Whether the server has IPv6 networking activated.`,
				},
				resource.Attribute{
					Name:        "enable_private_network",
					Description: `(Optional) Whether the server has private networking support enabled.`,
				},
				resource.Attribute{
					Name:        "notify_activate",
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
					Description: `(Optional) A label for the server. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the server.`,
				},
				resource.Attribute{
					Name:        "region_id",
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
					Name:        "vps_cpu_count",
					Description: `The number of virtual CPUs available on the server.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The physical location of the server.`,
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
					Name:        "pending_charges",
					Description: `Charges pending for this server's subscription in USD.`,
				},
				resource.Attribute{
					Name:        "cost_per_month",
					Description: `The server's cost per month in USD.`,
				},
				resource.Attribute{
					Name:        "current_bandwidth",
					Description: `The server's current bandwidth usage in GB.`,
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
					Name:        "server_state",
					Description: `A more detailed server status (none, locked, installingbooting, isomounting, ok).`,
				},
				resource.Attribute{
					Name:        "v6_networks",
					Description: `A list of the server's IPv6 networks.`,
				},
				resource.Attribute{
					Name:        "internal_ip",
					Description: `The server's internal IP address.`,
				},
				resource.Attribute{
					Name:        "kvm_url",
					Description: `The server's current KVM URL. This URL will change periodically. It is not advised to cache this value.`,
				},
				resource.Attribute{
					Name:        "plan_id",
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
					Name:        "network_ids",
					Description: `A list of private network IDs attached to the server.`,
				},
				resource.Attribute{
					Name:        "ssh_key_ids",
					Description: `A list of SSH key IDs applied to the server on install.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `Base64 encoded generic data store, which some provisioning tools and cloud operating systems use as a configuration file. It is generally consumed only once after an instance has been launched, but individual needs may vary.`,
				},
				resource.Attribute{
					Name:        "auto_backup",
					Description: `Whether automatic backups are enabled for this server.`,
				},
				resource.Attribute{
					Name:        "enable_ipv6",
					Description: `Whether the server has IPv6 networking activated.`,
				},
				resource.Attribute{
					Name:        "enable_private_network",
					Description: `Whether the server has private networking support enabled.`,
				},
				resource.Attribute{
					Name:        "notify_activate",
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
					Description: `A label for the server. ## Import Servers can be imported using the server ` + "`" + `SUBID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_server.my_server 1312965 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the server.`,
				},
				resource.Attribute{
					Name:        "region_id",
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
					Name:        "vps_cpu_count",
					Description: `The number of virtual CPUs available on the server.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The physical location of the server.`,
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
					Name:        "pending_charges",
					Description: `Charges pending for this server's subscription in USD.`,
				},
				resource.Attribute{
					Name:        "cost_per_month",
					Description: `The server's cost per month in USD.`,
				},
				resource.Attribute{
					Name:        "current_bandwidth",
					Description: `The server's current bandwidth usage in GB.`,
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
					Name:        "server_state",
					Description: `A more detailed server status (none, locked, installingbooting, isomounting, ok).`,
				},
				resource.Attribute{
					Name:        "v6_networks",
					Description: `A list of the server's IPv6 networks.`,
				},
				resource.Attribute{
					Name:        "internal_ip",
					Description: `The server's internal IP address.`,
				},
				resource.Attribute{
					Name:        "kvm_url",
					Description: `The server's current KVM URL. This URL will change periodically. It is not advised to cache this value.`,
				},
				resource.Attribute{
					Name:        "plan_id",
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
					Name:        "network_ids",
					Description: `A list of private network IDs attached to the server.`,
				},
				resource.Attribute{
					Name:        "ssh_key_ids",
					Description: `A list of SSH key IDs applied to the server on install.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `Base64 encoded generic data store, which some provisioning tools and cloud operating systems use as a configuration file. It is generally consumed only once after an instance has been launched, but individual needs may vary.`,
				},
				resource.Attribute{
					Name:        "auto_backup",
					Description: `Whether automatic backups are enabled for this server.`,
				},
				resource.Attribute{
					Name:        "enable_ipv6",
					Description: `Whether the server has IPv6 networking activated.`,
				},
				resource.Attribute{
					Name:        "enable_private_network",
					Description: `Whether the server has private networking support enabled.`,
				},
				resource.Attribute{
					Name:        "notify_activate",
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
					Description: `A label for the server. ## Import Servers can be imported using the server ` + "`" + `SUBID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_server.my_server 1312965 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "vps_id",
					Description: `(Required) ID of a given server that you want to create a snapshot from.`,
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
					Name:        "vps_id",
					Description: `The ID of the server that the snapshot was created from.`,
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
					Description: `The app id which the snapshot is associated with. ## Import Snapshots can be imported using the Snapshot ` + "`" + `SNAPSHOTID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_snapshot_url.my_snapshot 9735ced831ed2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID for the given snapshot.`,
				},
				resource.Attribute{
					Name:        "vps_id",
					Description: `The ID of the server that the snapshot was created from.`,
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
					Description: `The app id which the snapshot is associated with. ## Import Snapshots can be imported using the Snapshot ` + "`" + `SNAPSHOTID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_snapshot_url.my_snapshot 9735ced831ed2 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `The app id which the snapshot is associated with. ## Import Snapshots from Url can be imported using the Snapshot ` + "`" + `SNAPSHOTID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_snapshot_from_url.my_snapshot 9735ced831ed2 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `The app id which the snapshot is associated with. ## Import Snapshots from Url can be imported using the Snapshot ` + "`" + `SNAPSHOTID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_snapshot_from_url.my_snapshot 9735ced831ed2 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `The date the SSH key was added to your Vultr account. ## Import SSH keys can be imported using the SSH key ` + "`" + `SSHKEYID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_ssh_key.my_key 541b4960f23bd ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `The date the SSH key was added to your Vultr account. ## Import SSH keys can be imported using the SSH key ` + "`" + `SSHKEYID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_ssh_key.my_key 541b4960f23bd ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Required) Contents of the startup script.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type of startup script. Default is boot. ## Attributes Reference The following attributes are exported:`,
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
					Description: `The contents of the startup script. ## Import Startup Scripts can be imported using the Startup Scripts ` + "`" + `SCRIPTID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_startup_script.my_script 537932 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `The contents of the startup script. ## Import Startup Scripts can be imported using the Startup Scripts ` + "`" + `SCRIPTID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_startup_script.my_script 537932 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `Whether API is enabled for the user.`,
				},
				resource.Attribute{
					Name:        "api_key",
					Description: `API Key that is assigned to this user. ## Import Users can be imported using the User ` + "`" + `USERID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_user.myuser cbe5ced2ae716 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `Whether API is enabled for the user.`,
				},
				resource.Attribute{
					Name:        "api_key",
					Description: `API Key that is assigned to this user. ## Import Users can be imported using the User ` + "`" + `USERID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import vultr_user.myuser cbe5ced2ae716 ` + "`" + `` + "`" + `` + "`" + ``,
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
		"vultr_iso":               6,
		"vultr_network":           7,
		"vultr_reserved_ip":       8,
		"vultr_server":            9,
		"vultr_snapshot":          10,
		"vultr_snapshot_from_url": 11,
		"vultr_ssh_key":           12,
		"vultr_startup_script":    13,
		"vultr_user":              14,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
