package linode

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "linode_database_access_controls",
			Category:         "Resources",
			ShortDescription: `Manages the access controls for a Linode Database.`,
			Description:      ``,
			Keywords: []string{
				"database",
				"access",
				"controls",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "database_id",
					Description: `(Required) The unique ID of the target database.`,
				},
				resource.Attribute{
					Name:        "allow_list",
					Description: `(Required) A list of IP addresses that can access the Managed Database. Each item can be a single IP address or a range in CIDR format.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_database_mysql",
			Category:         "Resources",
			ShortDescription: `Manages a Linode MySQL Database.`,
			Description:      ``,
			Keywords: []string{
				"database",
				"mysql",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "engine_id",
					Description: `(Required) The Managed Database engine in engine/version format. (e.g. ` + "`" + `mysql/8.0.26` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Required) A unique, user-defined string referring to the Managed Database.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The region to use for the Managed Database.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The Linode Instance type used for the nodes of the Managed Database instance. - - -`,
				},
				resource.Attribute{
					Name:        "allow_list",
					Description: `(Optional) A list of IP addresses that can access the Managed Database. Each item can be a single IP address or a range in CIDR format. Use ` + "`" + `linode_database_mysql_firewall` + "`" + ` to manage your allow list separately.`,
				},
				resource.Attribute{
					Name:        "cluster_size",
					Description: `(Optional) The number of Linode Instance nodes deployed to the Managed Database. (default ` + "`" + `1` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `(Optional) Whether the Managed Databases is encrypted. (default ` + "`" + `false` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "replication_type",
					Description: `(Optional) The replication method used for the Managed Database. (` + "`" + `none` + "`" + `, ` + "`" + `asynch` + "`" + `, ` + "`" + `semi_synch` + "`" + `; default ` + "`" + `none` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "ssl_connection",
					Description: `(Optional) Whether to require SSL credentials to establish a connection to the Managed Database. (default ` + "`" + `false` + "`" + `) ## Attributes In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Managed Database.`,
				},
				resource.Attribute{
					Name:        "ca_cert",
					Description: `The base64-encoded SSL CA certificate for the Managed Database instance.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `When this Managed Database was created.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `The Managed Database engine. (e.g. ` + "`" + `mysql` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "host_primary",
					Description: `The primary host for the Managed Database.`,
				},
				resource.Attribute{
					Name:        "host_secondary",
					Description: `The secondary/private network host for the Managed Database.`,
				},
				resource.Attribute{
					Name:        "root_password",
					Description: `The randomly-generated root password for the Managed Database instance.`,
				},
				resource.Attribute{
					Name:        "root_username",
					Description: `The root username for the Managed Database instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The operating status of the Managed Database.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `When this Managed Database was last updated.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The Managed Database engine version. (e.g. ` + "`" + `v8.0.26` + "`" + `) ## Import Linode MySQL Databases can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `sh terraform import linode_database_mysql.foobar 1234567 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_domain",
			Category:         "Resources",
			ShortDescription: `Manages a Linode Domain Record.`,
			Description:      ``,
			Keywords: []string{
				"domain",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) The domain this Domain represents. These must be unique in our system; you cannot have two Domains representing the same domain.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) If this Domain represents the authoritative source of information for the domain it describes, or if it is a read-only copy of a master (also called a slave).`,
				},
				resource.Attribute{
					Name:        "soa_email",
					Description: `(Required) Start of Authority email address. This is required for master Domains.`,
				},
				resource.Attribute{
					Name:        "master_ips",
					Description: `(Required for type="slave") The IP addresses representing the master DNS for this Domain. - - -`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Used to control whether this Domain is currently being rendered (defaults to "active").`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description for this Domain. This is for display purposes only.`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `(Optional) The group this Domain belongs to. This is for display purposes only.`,
				},
				resource.Attribute{
					Name:        "ttl_sec",
					Description: `(Optional) 'Time to Live' - the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers. Valid values are 0, 30, 120, 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.`,
				},
				resource.Attribute{
					Name:        "retry_sec",
					Description: `(Optional) The interval, in seconds, at which a failed refresh should be retried. Valid values are 0, 30, 120, 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.`,
				},
				resource.Attribute{
					Name:        "expire_sec",
					Description: `(Optional) The amount of time in seconds that may pass before this Domain is no longer authoritative. Valid values are 0, 30, 120, 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.`,
				},
				resource.Attribute{
					Name:        "refresh_sec",
					Description: `(Optional) The amount of time in seconds before this Domain should be refreshed. Valid values are 0, 30, 120, 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.`,
				},
				resource.Attribute{
					Name:        "axfr_ips",
					Description: `(Optional) The list of IPs that may perform a zone transfer for this Domain. This is potentially dangerous, and should be set to an empty list unless you intend to use it.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A list of tags applied to this object. Tags are for organizational purposes only. ## Attributes This resource exports no additional attributes, however ` + "`" + `status` + "`" + ` may reflect degraded states. ## Import Linodes Domains can be imported using the Linode Domain ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `sh terraform import linode_domain.foobar 1234567 ` + "`" + `` + "`" + `` + "`" + ` The Linode Guide, [Import Existing Infrastructure to Terraform](https://www.linode.com/docs/applications/configuration-management/import-existing-infrastructure-to-terraform/), offers resource importing examples for Domains and other Linode resource types.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_domain_record",
			Category:         "Resources",
			ShortDescription: `Manages a Linode Domain Record.`,
			Description:      ``,
			Keywords: []string{
				"domain",
				"record",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of this Record. Setting this is invalid for ` + "`" + `SRV` + "`" + ` records as it is generated by the API. This field's actual usage depends on the type of record this represents. For A and AAAA records, this is the subdomain being associated with an IP address.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Required) The ID of the Domain to access.`,
				},
				resource.Attribute{
					Name:        "record_type",
					Description: `(Required) The type of Record this is in the DNS system. For example, A records associate a domain name with an IPv4 address, and AAAA records associate a domain name with an IPv6 address. See all supported record types [here](https://www.linode.com/docs/api/domains/#domain-record-create__request-body-schema).`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Required) The target for this Record. This field's actual usage depends on the type of record this represents. For A and AAAA records, this is the address the named Domain should resolve to. - - -`,
				},
				resource.Attribute{
					Name:        "ttl_sec",
					Description: `(Optional) 'Time to Live' - the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers. Valid values are 30, 120, 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) The priority of the target host. Lower values are preferred.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) The protocol this Record's service communicates with. Only valid for SRV records.`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `(Optional) The service this Record identified. Only valid for SRV records.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) The tag portion of a CAA record. It is invalid to set this on other record types.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The port this Record points to.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) The relative weight of this Record. Higher values are preferred. ## Attributes This resource exports no additional attributes. ## Import Linodes Domain Records can be imported using the Linode Domain ` + "`" + `id` + "`" + ` followed by the Domain Record ` + "`" + `id` + "`" + ` separated by a comma, e.g. ` + "`" + `` + "`" + `` + "`" + `sh terraform import linode_domain_record.www-foobar 1234567,7654321 ` + "`" + `` + "`" + `` + "`" + ` The Linode Guide, [Import Existing Infrastructure to Terraform](https://www.linode.com/docs/applications/configuration-management/import-existing-infrastructure-to-terraform/), offers resource importing examples for Domain Records and other Linode resource types.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_firewall",
			Category:         "Resources",
			ShortDescription: `Manages a Linode Firewall.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "label",
					Description: `(Required) This Firewall's unique label.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, the Firewall's rules are not enforced (defaults to ` + "`" + `false` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "inbound_policy",
					Description: `(Required) The default behavior for inbound traffic. This setting can be overridden by updating the inbound.action property of the Firewall Rule. (` + "`" + `ACCEPT` + "`" + `, ` + "`" + `DROP` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "outbound_policy",
					Description: `(Required) The default behavior for outbound traffic. This setting can be overridden by updating the outbound.action property for an individual Firewall Rule. (` + "`" + `ACCEPT` + "`" + `, ` + "`" + `DROP` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "linodes",
					Description: `(Optional) A list of IDs of Linodes this Firewall should govern it's network traffic for.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A list of tags applied to the Kubernetes cluster. Tags are for organizational purposes only. ### inbound and outbound The following arguments are supported in the inbound and outbound rule blocks:`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(required) Used to identify this rule. For display purposes only.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(required) Controls whether traffic is accepted or dropped by this rule (` + "`" + `ACCEPT` + "`" + `, ` + "`" + `DROP` + "`" + `). Overrides the Firewall’s inbound_policy if this is an inbound rule, or the outbound_policy if this is an outbound rule.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The network protocol this rule controls. (` + "`" + `TCP` + "`" + `, ` + "`" + `UDP` + "`" + `, ` + "`" + `ICMP` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `(Optional) A string representation of ports and/or port ranges (i.e. "443" or "80-90, 91").`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `(Optional) A list of IPv4 addresses or networks. Must be in IP/mask format.`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `(Optional) A list of IPv6 addresses or networks. Must be in IP/mask format. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Firewall.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the Firewall.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Firewall Device.`,
				},
				resource.Attribute{
					Name:        "entity_id",
					Description: `The ID of the underlying entity this device references (i.e. the Linode's ID).`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of Firewall Device.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The label of the underlying entity this device references.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Firewall.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the Firewall.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Firewall Device.`,
				},
				resource.Attribute{
					Name:        "entity_id",
					Description: `The ID of the underlying entity this device references (i.e. the Linode's ID).`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of Firewall Device.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The label of the underlying entity this device references.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_firewall_device",
			Category:         "Resources",
			ShortDescription: `Manages a Linode Firewall Device.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"device",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "firewall_id",
					Description: `(Required) The unique ID of the target Firewall.`,
				},
				resource.Attribute{
					Name:        "entity_id",
					Description: `(Required) The unique ID of the entity to attach.`,
				},
				resource.Attribute{
					Name:        "entity_type",
					Description: `(Optional) The type of the entity to attach. (default: ` + "`" + `linode` + "`" + `) ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `When the Firewall Device was last created.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `When the Firewall Device was last updated.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created",
					Description: `When the Firewall Device was last created.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `When the Firewall Device was last updated.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_image",
			Category:         "Resources",
			ShortDescription: `Manages a Linode Image.`,
			Description:      ``,
			Keywords: []string{
				"image",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "label",
					Description: `(Required) A short description of the Image. Labels cannot contain special characters.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A detailed description of this Image. - - - The following arguments apply to creating an image from an existing Linode Instance:`,
				},
				resource.Attribute{
					Name:        "disk_id",
					Description: `(Required) The ID of the Linode Disk that this Image will be created from.`,
				},
				resource.Attribute{
					Name:        "linode_id",
					Description: `(Required) The ID of the Linode that this Image will be created from. - - - ~>`,
				},
				resource.Attribute{
					Name:        "file_path",
					Description: `(Required) The path of the image file to be uploaded.`,
				},
				resource.Attribute{
					Name:        "file_hash",
					Description: `(Optional) The MD5 hash of the file to be uploaded. This is used to trigger file updates.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The region of the image. See all regions [here](https://api.linode.com/v4/regions). ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 20 mins) Used when creating the instance image (until the instance is available) ## Attributes This resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of this Image. The ID of private images begin with ` + "`" + `private/` + "`" + ` followed by the numeric identifier of the private image, for example ` + "`" + `private/12345` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `When this Image was created.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `The name of the User who created this Image.`,
				},
				resource.Attribute{
					Name:        "deprecated",
					Description: `Whether or not this Image is deprecated. Will only be True for deprecated public Images.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `True if the Image is public.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The minimum size this Image needs to deploy. Size is in MB.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `How the Image was created. 'Manual' Images can be created at any time. 'Automatic' images are created automatically from a deleted Linode.`,
				},
				resource.Attribute{
					Name:        "expiry",
					Description: `Only Images created automatically (from a deleted Linode; type=automatic) will expire.`,
				},
				resource.Attribute{
					Name:        "vendor",
					Description: `The upstream distribution vendor. Nil for private Images. ## Import Linodes Images can be imported using the Linode Image ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `sh terraform import linode_image.myimage 1234567 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_instance",
			Category:         "Resources",
			ShortDescription: `Manages a Linode instance.`,
			Description:      ``,
			Keywords: []string{
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Required) This is the location where the Linode is deployed. Examples are ` + "`" + `"us-east"` + "`" + `, ` + "`" + `"us-west"` + "`" + `, ` + "`" + `"ap-south"` + "`" + `, etc. See all regions [here](https://api.linode.com/v4/regions).`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The Linode type defines the pricing, CPU, disk, and RAM specs of the instance. Examples are ` + "`" + `"g6-nanode-1"` + "`" + `, ` + "`" + `"g6-standard-2"` + "`" + `, ` + "`" + `"g6-highmem-16"` + "`" + `, ` + "`" + `"g6-dedicated-16"` + "`" + `, etc. See all types [here](https://api.linode.com/v4/linode/types). - - -`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Optional) The Linode's label is for display purposes only. If no label is provided for a Linode, a default will be assigned.`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `(Optional) The display group of the Linode instance.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A list of tags applied to this object. Tags are for organizational purposes only.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `(Optional) If true, the created Linode will have private networking enabled, allowing use of the 192.168.128.0/17 network within the Linode's region. It can be enabled on an existing Linode but it can't be disabled.`,
				},
				resource.Attribute{
					Name:        "shared_ipv4",
					Description: `(Optional) A set of IPv4 addresses to be shared with the Instance. These IP addresses can be both private and public, but must be in the same region as the instance.`,
				},
				resource.Attribute{
					Name:        "resize_disk",
					Description: `(Optional) If true, changes in Linode type will attempt to upsize or downsize implicitly created disks. This must be false if explicit disks are defined.`,
				},
				resource.Attribute{
					Name:        "alerts.0.cpu",
					Description: `(Optional) The percentage of CPU usage required to trigger an alert. If the average CPU usage over two hours exceeds this value, we'll send you an alert. If this is set to 0, the alert is disabled.`,
				},
				resource.Attribute{
					Name:        "alerts.0.network_in",
					Description: `(Optional) The amount of incoming traffic, in Mbit/s, required to trigger an alert. If the average incoming traffic over two hours exceeds this value, we'll send you an alert. If this is set to 0 (zero), the alert is disabled.`,
				},
				resource.Attribute{
					Name:        "alerts.0.network_out",
					Description: `(Optional) The amount of outbound traffic, in Mbit/s, required to trigger an alert. If the average outbound traffic over two hours exceeds this value, we'll send you an alert. If this is set to 0 (zero), the alert is disabled.`,
				},
				resource.Attribute{
					Name:        "alerts.0.transfer_quota",
					Description: `(Optional) The percentage of network transfer that may be used before an alert is triggered. When this value is exceeded, we'll alert you. If this is set to 0 (zero), the alert is disabled.`,
				},
				resource.Attribute{
					Name:        "alerts.0.io",
					Description: `(Optional) The amount of disk IO operation per second required to trigger an alert. If the average disk IO over two hours exceeds this value, we'll send you an alert. If set to 0, this alert is disabled.`,
				},
				resource.Attribute{
					Name:        "backups_enabled",
					Description: `(Optional) If this field is set to true, the created Linode will automatically be enrolled in the Linode Backup service. This will incur an additional charge. The cost for the Backup service is dependent on the Type of Linode deployed.`,
				},
				resource.Attribute{
					Name:        "watchdog_enabled",
					Description: `(Optional) The watchdog, named Lassie, is a Shutdown Watchdog that monitors your Linode and will reboot it if it powers off unexpectedly. It works by issuing a boot job when your Linode powers off without a shutdown job being responsible. To prevent a loop, Lassie will give up if there have been more than 5 boot jobs issued within 15 minutes.`,
				},
				resource.Attribute{
					Name:        "booted",
					Description: `(Optional) Specifies whether the Linode should be ` + "`" + `running` + "`" + ` or ` + "`" + `offline` + "`" + `. If unspecified, the Linode's power status will not be managed by the Provider.`,
				},
				resource.Attribute{
					Name:        "authorized_keys",
					Description: `(Optional with ` + "`" + `image` + "`" + `) A list of SSH public keys to deploy for the root user on the newly created Linode.`,
				},
				resource.Attribute{
					Name:        "authorized_users",
					Description: `(Optional with ` + "`" + `image` + "`" + `) A list of Linode usernames. If the usernames have associated SSH keys, the keys will be appended to the ` + "`" + `root` + "`" + ` user's ` + "`" + `~/.ssh/authorized_keys` + "`" + ` file automatically.`,
				},
				resource.Attribute{
					Name:        "root_pass",
					Description: `(Optional) The initial password for the ` + "`" + `root` + "`" + ` user account.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Optional) An Image ID to deploy the Disk from. Official Linode Images start with linode/, while your Images start with ` + "`" + `private/` + "`" + `. See [images](https://api.linode.com/v4/images) for more information on the Images available for you to use. Examples are ` + "`" + `linode/debian9` + "`" + `, ` + "`" + `linode/fedora28` + "`" + `, ` + "`" + `linode/ubuntu16.04lts` + "`" + `, ` + "`" + `linode/arch` + "`" + `, and ` + "`" + `private/12345` + "`" + `. See all images [here](https://api.linode.com/v4/linode/images) (Requires a personal access token; docs [here](https://developers.linode.com/api/v4/images)).`,
				},
				resource.Attribute{
					Name:        "stackscript_id",
					Description: `(Optional) The StackScript to deploy to the newly created Linode. If provided, 'image' must also be provided, and must be an Image that is compatible with this StackScript.`,
				},
				resource.Attribute{
					Name:        "stackscript_data",
					Description: `(Optional) An object containing responses to any User Defined Fields present in the StackScript being deployed to this Linode. Only accepted if 'stackscript_id' is given. The required values depend on the StackScript being deployed.`,
				},
				resource.Attribute{
					Name:        "swap_size",
					Description: `(Optional) When deploying from an Image, this field is optional with a Linode API default of 512mb, otherwise it is ignored. This is used to set the swap disk size for the newly-created Linode.`,
				},
				resource.Attribute{
					Name:        "backup_id",
					Description: `(Optional) A Backup ID from another Linode's available backups. Your User must have read_write access to that Linode, the Backup must have a status of successful, and the Linode must be deployed to the same region as the Backup. See /linode/instances/{linodeId}/backups for a Linode's available backups. This field and the image field are mutually exclusive.`,
				},
				resource.Attribute{
					Name:        "boot_config_label",
					Description: `(Optional) The Label of the Instance Config that should be used to boot the Linode instance. If there is only one ` + "`" + `config` + "`" + `, the ` + "`" + `label` + "`" + ` of that ` + "`" + `config` + "`" + ` will be used as the ` + "`" + `boot_config_label` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Required) The disks label, which acts as an identifier in Terraform. This must be unique within each Linode Instance.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) The size of the Disk in MB.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the disk in the Linode API.`,
				},
				resource.Attribute{
					Name:        "filesystem",
					Description: `(Optional) The Disk filesystem can be one of: ` + "`" + `"raw"` + "`" + `, ` + "`" + `"swap"` + "`" + `, ` + "`" + `"ext3"` + "`" + `, ` + "`" + `"ext4"` + "`" + `, or ` + "`" + `"initrd"` + "`" + ` which has a max size of 32mb and can be used in the config ` + "`" + `initrd` + "`" + ` (not currently supported in this Terraform Provider).`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) If true, this Disk is read-only.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Optional) An Image ID to deploy the Disk from. Official Linode Images start with linode/, while your Images start with private/. See /images for more information on the Images available for you to use. Examples are ` + "`" + `linode/debian9` + "`" + `, ` + "`" + `linode/fedora28` + "`" + `, ` + "`" + `linode/ubuntu16.04lts` + "`" + `, ` + "`" + `linode/arch` + "`" + `, and ` + "`" + `private/12345` + "`" + `. See all images [here](https://api.linode.com/v4/images).`,
				},
				resource.Attribute{
					Name:        "authorized_keys",
					Description: `(Optional with ` + "`" + `image` + "`" + `) A list of SSH public keys to deploy for the root user on the newly created Linode. Only accepted if ` + "`" + `image` + "`" + ` is provided.`,
				},
				resource.Attribute{
					Name:        "authorized_users",
					Description: `(Optional with ` + "`" + `image` + "`" + `) A list of Linode usernames. If the usernames have associated SSH keys, the keys will be appended to the ` + "`" + `root` + "`" + ` user's ` + "`" + `~/.ssh/authorized_keys` + "`" + ` file automatically.`,
				},
				resource.Attribute{
					Name:        "root_pass",
					Description: `(Optional with ` + "`" + `image` + "`" + `) The initial password for the ` + "`" + `root` + "`" + ` user account.`,
				},
				resource.Attribute{
					Name:        "stackscript_id",
					Description: `(Optional with ` + "`" + `image` + "`" + `) The StackScript to deploy to the newly created Linode. If provided, 'image' must also be provided, and must be an Image that is compatible with this StackScript.`,
				},
				resource.Attribute{
					Name:        "stackscript_data",
					Description: `(Optional with ` + "`" + `image` + "`" + `) An object containing responses to any User Defined Fields present in the StackScript being deployed to this Linode. Only accepted if 'stackscript_id' is given. The required values depend on the StackScript being deployed.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Required) The Config's label for display purposes. Also used by ` + "`" + `boot_config_label` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "helpers",
					Description: `(Options) Helpers enabled when booting to this Linode Config.`,
				},
				resource.Attribute{
					Name:        "updatedb_disabled",
					Description: `(Optional) Disables updatedb cron job to avoid disk thrashing.`,
				},
				resource.Attribute{
					Name:        "distro",
					Description: `(Optional) Controls the behavior of the Linode Config's Distribution Helper setting.`,
				},
				resource.Attribute{
					Name:        "modules_dep",
					Description: `(Optional) Creates a modules dependency file for the Kernel you run.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Optional) Controls the behavior of the Linode Config's Network Helper setting, used to automatically configure additional IP addresses assigned to this instance.`,
				},
				resource.Attribute{
					Name:        "devices",
					Description: `(Optional) A list of ` + "`" + `disk` + "`" + ` or ` + "`" + `volume` + "`" + ` attachments for this ` + "`" + `config` + "`" + `. If the ` + "`" + `boot_config_label` + "`" + ` omits a ` + "`" + `devices` + "`" + ` block, the Linode will not be booted.`,
				},
				resource.Attribute{
					Name:        "disk_label",
					Description: `(Optional) The ` + "`" + `label` + "`" + ` of the ` + "`" + `disk` + "`" + ` to map to this ` + "`" + `device` + "`" + ` slot.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `(Optional) The Volume ID to map to this ` + "`" + `device` + "`" + ` slot.`,
				},
				resource.Attribute{
					Name:        "disk_id",
					Description: `(Computed) The Disk ID of the associated ` + "`" + `disk_label` + "`" + `, if used.`,
				},
				resource.Attribute{
					Name:        "kernel",
					Description: `(Optional) - A Kernel ID to boot a Linode with. Default is based on image choice. Examples are ` + "`" + `linode/latest-64bit` + "`" + `, ` + "`" + `linode/grub2` + "`" + `, ` + "`" + `linode/direct-disk` + "`" + `, etc. See all kernels [here](https://api.linode.com/v4/linode/kernels). Note that this is a paginated API endpoint ([docs](https://developers.linode.com/api/v4/linode-kernels)).`,
				},
				resource.Attribute{
					Name:        "run_level",
					Description: `(Optional) - Defines the state of your Linode after booting. Defaults to ` + "`" + `"default"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "virt_mode",
					Description: `(Optional) - Controls the virtualization mode. Defaults to ` + "`" + `"paravirt"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "root_device",
					Description: `(Optional) - The root device to boot. The corresponding disk must be attached to a ` + "`" + `device` + "`" + ` slot. Example: ` + "`" + `"/dev/sda"` + "`" + ``,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) - Arbitrary user comments about this ` + "`" + `config` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "memory_limit",
					Description: `(Optional) - Defaults to the total RAM of the Linode`,
				},
				resource.Attribute{
					Name:        "purpose",
					Description: `(Required) The type of interface. (` + "`" + `public` + "`" + `, ` + "`" + `vlan` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Optional) The name of this interface. If the interface is a VLAN, a label is required.`,
				},
				resource.Attribute{
					Name:        "ipam_address",
					Description: `(Optional) This Network Interface’s private IP address in Classless Inter-Domain Routing (CIDR) notation. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 10 mins) Used when launching the instance (until it reaches the initial ` + "`" + `running` + "`" + ` state)`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 20 mins) Used when stopping and starting the instance when necessary during update - e.g. when changing instance type`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 10 mins) Used when terminating the instance ## Attributes This Linode Instance resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the instance, indicating the current readiness state. (` + "`" + `running` + "`" + `, ` + "`" + `offline` + "`" + `, ...)`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `A string containing the Linode's public IP address.`,
				},
				resource.Attribute{
					Name:        "private_ip_address",
					Description: `This Linode's Private IPv4 Address, if enabled. The regional private IP address range, 192.168.128.0/17, is shared by all Linode Instances in a region.`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `This Linode's IPv6 SLAAC addresses. This address is specific to a Linode, and may not be shared. The prefix (` + "`" + `/64` + "`" + `) is included in this attribute.`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `This Linode's IPv4 Addresses. Each Linode is assigned a single public IPv4 address upon creation, and may get a single private IPv4 address if needed. You may need to open a support ticket to get additional IPv4 addresses.`,
				},
				resource.Attribute{
					Name:        "specs.0.disk",
					Description: `The amount of storage space, in GB. this Linode has access to. A typical Linode will divide this space between a primary disk with an image deployed to it, and a swap disk, usually 512 MB. This is the default configuration created when deploying a Linode with an image through POST /linode/instances.`,
				},
				resource.Attribute{
					Name:        "specs.0.memory",
					Description: `The amount of RAM, in MB, this Linode has access to. Typically a Linode will choose to boot with all of its available RAM, but this can be configured in a Config profile.`,
				},
				resource.Attribute{
					Name:        "specs.0.vcpus",
					Description: `The number of vcpus this Linode has access to. Typically a Linode will choose to boot with all of its available vcpus, but this can be configured in a Config Profile.`,
				},
				resource.Attribute{
					Name:        "specs.0.transfer",
					Description: `The amount of network transfer this Linode is allotted each month.`,
				},
				resource.Attribute{
					Name:        "backups",
					Description: `Information about this Linode's backups status.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `If this Linode has the Backup service enabled.`,
				},
				resource.Attribute{
					Name:        "day",
					Description: `The day of the week that your Linode's weekly Backup is taken. If not set manually, a day will be chosen for you. Backups are taken every day, but backups taken on this day are preferred when selecting backups to retain for a longer period. If not set manually, then when backups are initially enabled, this may come back as "Scheduling" until the day is automatically selected.`,
				},
				resource.Attribute{
					Name:        "window",
					Description: `The window ('W0'-'W22') in which your backups will be taken, in UTC. A backups window is a two-hour span of time in which the backup may occur. For example, 'W10' indicates that your backups should be taken between 10:00 and 12:00. If you do not choose a backup window, one will be selected for you automatically. If not set manually, when backups are initially enabled this may come back as Scheduling until the window is automatically selected. ## Import Linodes Instances can be imported using the Linode ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `sh terraform import linode_instance.mylinode 1234567 ` + "`" + `` + "`" + `` + "`" + ` When importing an instance, all ` + "`" + `disk` + "`" + ` and ` + "`" + `config` + "`" + ` values must be represented. Imported disks must include their ` + "`" + `label` + "`" + ` value.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_instance_ip",
			Category:         "Resources",
			ShortDescription: `Manages a Linode instance IP.`,
			Description:      ``,
			Keywords: []string{
				"instance",
				"ip",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "linode_id",
					Description: `(Required) The ID of the Linode to allocate an IPv4 address for.`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `(Optional) Whether the IPv4 address is public or private. Defaults to true.`,
				},
				resource.Attribute{
					Name:        "rdns",
					Description: `(Optional) The reverse DNS assigned to this address.`,
				},
				resource.Attribute{
					Name:        "apply_immediately",
					Description: `(Optional) If true, the instance will be rebooted to update network interfaces. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `The default gateway for this address`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `The number of bits set in the subnet mask.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The resulting IPv4 address.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region this IP resides in.`,
				},
				resource.Attribute{
					Name:        "subnet_mask",
					Description: `The mask that separates host bits from network bits for this address.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of IP address. (` + "`" + `ipv4` + "`" + `, ` + "`" + `ipv6` + "`" + `, ` + "`" + `ipv6/pool` + "`" + `, ` + "`" + `ipv6/range` + "`" + `)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "gateway",
					Description: `The default gateway for this address`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `The number of bits set in the subnet mask.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The resulting IPv4 address.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region this IP resides in.`,
				},
				resource.Attribute{
					Name:        "subnet_mask",
					Description: `The mask that separates host bits from network bits for this address.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of IP address. (` + "`" + `ipv4` + "`" + `, ` + "`" + `ipv6` + "`" + `, ` + "`" + `ipv6/pool` + "`" + `, ` + "`" + `ipv6/range` + "`" + `)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_ipv6_range",
			Category:         "Resources",
			ShortDescription: `Manages a Linode IPv6 range.`,
			Description:      ``,
			Keywords: []string{
				"ipv6",
				"range",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "prefix_length",
					Description: `(Required) The prefix length of the IPv6 range.`,
				},
				resource.Attribute{
					Name:        "linode_id",
					Description: `(Required) The ID of the Linode to assign this range to. This field may be updated to reassign the IPv6 range.`,
				},
				resource.Attribute{
					Name:        "route_target",
					Description: `(Required) The IPv6 SLAAC address to assign this range to. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "is_bgp",
					Description: `Whether this IPv6 range is shared.`,
				},
				resource.Attribute{
					Name:        "linodes",
					Description: `A list of Linodes targeted by this IPv6 range. Includes Linodes with IP sharing.`,
				},
				resource.Attribute{
					Name:        "range",
					Description: `The IPv6 range of addresses in this pool.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region for this range of IPv6 addresses.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "is_bgp",
					Description: `Whether this IPv6 range is shared.`,
				},
				resource.Attribute{
					Name:        "linodes",
					Description: `A list of Linodes targeted by this IPv6 range. Includes Linodes with IP sharing.`,
				},
				resource.Attribute{
					Name:        "range",
					Description: `The IPv6 range of addresses in this pool.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region for this range of IPv6 addresses.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_lke_cluster",
			Category:         "Resources",
			ShortDescription: `Manages a Linode instance.`,
			Description:      ``,
			Keywords: []string{
				"lke",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "label",
					Description: `(Required) This Kubernetes cluster's unique label.`,
				},
				resource.Attribute{
					Name:        "k8s_version",
					Description: `(Required) The desired Kubernetes version for this Kubernetes cluster in the format of ` + "`" + `major.minor` + "`" + ` (e.g. ` + "`" + `1.21` + "`" + `), and the latest supported patch version will be deployed.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) This Kubernetes cluster's location.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) An array of tags applied to the Kubernetes cluster. Tags are for organizational purposes only. ### pool The following arguments are supported in the ` + "`" + `pool` + "`" + ` specification block:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) A Linode Type for all of the nodes in the Node Pool. See all node types [here](https://api.linode.com/v4/linode/types).`,
				},
				resource.Attribute{
					Name:        "count",
					Description: `(Required) The number of nodes in the Node Pool.`,
				},
				resource.Attribute{
					Name:        "min",
					Description: `(Required) The minimum number of nodes to autoscale to.`,
				},
				resource.Attribute{
					Name:        "max",
					Description: `(Required) The maximum number of nodes to autoscale to. ### control_plane The following arguments are supported in the ` + "`" + `control_plane` + "`" + ` specification block:`,
				},
				resource.Attribute{
					Name:        "high_availability",
					Description: `(Optional) Defines whether High Availability is enabled for the cluster Control Plane. This is an`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the cluster.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the cluster.`,
				},
				resource.Attribute{
					Name:        "api_endpoints",
					Description: `The endpoints for the Kubernetes API server.`,
				},
				resource.Attribute{
					Name:        "kubeconfig",
					Description: `The base64 encoded kubeconfig for the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "dashboard_url",
					Description: `The Kubernetes Dashboard access URL for this cluster.`,
				},
				resource.Attribute{
					Name:        "pool",
					Description: `Additional nested attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Node Pool.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the node.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of the underlying Linode instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the node. (` + "`" + `ready` + "`" + `, ` + "`" + `not_ready` + "`" + `) ## Import LKE Clusters can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `sh terraform import linode_lke_cluster.my_cluster 12345 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the cluster.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the cluster.`,
				},
				resource.Attribute{
					Name:        "api_endpoints",
					Description: `The endpoints for the Kubernetes API server.`,
				},
				resource.Attribute{
					Name:        "kubeconfig",
					Description: `The base64 encoded kubeconfig for the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "dashboard_url",
					Description: `The Kubernetes Dashboard access URL for this cluster.`,
				},
				resource.Attribute{
					Name:        "pool",
					Description: `Additional nested attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Node Pool.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the node.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of the underlying Linode instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the node. (` + "`" + `ready` + "`" + `, ` + "`" + `not_ready` + "`" + `) ## Import LKE Clusters can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `sh terraform import linode_lke_cluster.my_cluster 12345 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_nodebalancer",
			Category:         "Resources",
			ShortDescription: `Manages a Linode NodeBalancer.`,
			Description:      ``,
			Keywords: []string{
				"nodebalancer",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The region where this NodeBalancer will be deployed. Examples are ` + "`" + `"us-east"` + "`" + `, ` + "`" + `"us-west"` + "`" + `, ` + "`" + `"ap-south"` + "`" + `, etc. See all regions [here](https://api.linode.com/v4/regions).`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Optional) The label of the Linode NodeBalancer`,
				},
				resource.Attribute{
					Name:        "client_conn_throttle",
					Description: `(Optional) Throttle connections per second (0-20). Set to 0 (default) to disable throttling.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A list of tags applied to this object. Tags are for organizational purposes only. ## Attributes This resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `This NodeBalancer's hostname, ending with .nodebalancer.linode.com`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `The Public IPv4 Address of this NodeBalancer`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `The Public IPv6 Address of this NodeBalancer`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `When this NodeBalancer was created`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `When this NodeBalancer was last updated.`,
				},
				resource.Attribute{
					Name:        "in",
					Description: `The total transfer, in MB, used by this NodeBalancer for the current month`,
				},
				resource.Attribute{
					Name:        "out",
					Description: `The total inbound transfer, in MB, used for this NodeBalancer for the current month`,
				},
				resource.Attribute{
					Name:        "total",
					Description: `The total outbound transfer, in MB, used for this NodeBalancer for the current month ## Import Linodes NodeBalancers can be imported using the Linode NodeBalancer ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `sh terraform import linode_nodebalancer.mynodebalancer 1234567 ` + "`" + `` + "`" + `` + "`" + ` The Linode Guide, [Import Existing Infrastructure to Terraform](https://www.linode.com/docs/applications/configuration-management/import-existing-infrastructure-to-terraform/), offers resource importing examples for NodeBalancers and other Linode resource types.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_nodebalancer_config",
			Category:         "Resources",
			ShortDescription: `Manages a Linode NodeBalancer Config.`,
			Description:      ``,
			Keywords: []string{
				"nodebalancer",
				"config",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "nodebalancer_id",
					Description: `(Required) The ID of the NodeBalancer to access.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The region where this nodebalancer_config will be deployed. Examples are ` + "`" + `"us-east"` + "`" + `, ` + "`" + `"us-west"` + "`" + `, ` + "`" + `"ap-south"` + "`" + `, etc. See all regions [here](https://api.linode.com/v4/regions).`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) The protocol this port is configured to serve. If this is set to https you must include an ssl_cert and an ssl_key. (` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + `, ` + "`" + `tcp` + "`" + `) (Defaults to ` + "`" + `http` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "proxy_protocol",
					Description: `(Optional) The version of ProxyProtocol to use for the underlying NodeBalancer. This requires protocol to be ` + "`" + `tcp` + "`" + `. (` + "`" + `none` + "`" + `, ` + "`" + `v1` + "`" + `, ` + "`" + `v2` + "`" + `) (Defaults to ` + "`" + `none` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The TCP port this Config is for. These values must be unique across configs on a single NodeBalancer (you can't have two configs for port 80, for example). While some ports imply some protocols, no enforcement is done and you may configure your NodeBalancer however is useful to you. For example, while port 443 is generally used for HTTPS, you do not need SSL configured to have a NodeBalancer listening on port 443. (Defaults to 80)`,
				},
				resource.Attribute{
					Name:        "algorithm",
					Description: `(Optional) What algorithm this NodeBalancer should use for routing traffic to backends. (` + "`" + `roundrobin` + "`" + `, ` + "`" + `leastconn` + "`" + `, ` + "`" + `source` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "stickiness",
					Description: `(Optional) Controls how session stickiness is handled on this port. (` + "`" + `none` + "`" + `, ` + "`" + `table` + "`" + `, ` + "`" + `http_cookie` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "check",
					Description: `(Optional) The type of check to perform against backends to ensure they are serving requests. This is used to determine if backends are up or down. If none no check is performed. connection requires only a connection to the backend to succeed. http and http_body rely on the backend serving HTTP, and that the response returned matches what is expected. (` + "`" + `none` + "`" + `, ` + "`" + `connection` + "`" + `, ` + "`" + `http` + "`" + `, ` + "`" + `http_body` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "check_interval",
					Description: `(Optional) How often, in seconds, to check that backends are up and serving requests.`,
				},
				resource.Attribute{
					Name:        "check_timeout",
					Description: `(Optional) How long, in seconds, to wait for a check attempt before considering it failed. (1-30)`,
				},
				resource.Attribute{
					Name:        "check_attempts",
					Description: `(Optional) How many times to attempt a check before considering a backend to be down. (1-30)`,
				},
				resource.Attribute{
					Name:        "check_path",
					Description: `(Optional) The URL path to check on each backend. If the backend does not respond to this request it is considered to be down.`,
				},
				resource.Attribute{
					Name:        "check_passive",
					Description: `(Optional) If true, any response from this backend with a 5xx status code will be enough for it to be considered unhealthy and taken out of rotation.`,
				},
				resource.Attribute{
					Name:        "cipher_suite",
					Description: `(Optional) What ciphers to use for SSL connections served by this NodeBalancer. ` + "`" + `legacy` + "`" + ` is considered insecure and should only be used if necessary.`,
				},
				resource.Attribute{
					Name:        "ssl_cert",
					Description: `(Optional) The certificate this port is serving. This is not returned. If set, this field will come back as ` + "`" + `<REDACTED>` + "`" + `. Please use the ssl_commonname and ssl_fingerprint to identify the certificate.`,
				},
				resource.Attribute{
					Name:        "ssl_key",
					Description: `(Optional) The private key corresponding to this port's certificate. This is not returned. If set, this field will come back as ` + "`" + `<REDACTED>` + "`" + `. Please use the ssl_commonname and ssl_fingerprint to identify the certificate. ## Attributes This resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "ssl_commonname",
					Description: `The read-only common name automatically derived from the SSL certificate assigned to this NodeBalancerConfig. Please refer to this field to verify that the appropriate certificate is assigned to your NodeBalancerConfig.`,
				},
				resource.Attribute{
					Name:        "ssl_fingerprint",
					Description: `The read-only fingerprint automatically derived from the SSL certificate assigned to this NodeBalancerConfig. Please refer to this field to verify that the appropriate certificate is assigned to your NodeBalancerConfig.`,
				},
				resource.Attribute{
					Name:        "up",
					Description: `The number of backends considered to be 'UP' and healthy, and that are serving requests.`,
				},
				resource.Attribute{
					Name:        "down",
					Description: `The number of backends considered to be 'DOWN' and unhealthy. These are not in rotation, and not serving requests. ## Import NodeBalancer Configs can be imported using the NodeBalancer ` + "`" + `nodebalancer_id` + "`" + ` followed by the NodeBalancer Config ` + "`" + `id` + "`" + ` separated by a comma, e.g. ` + "`" + `` + "`" + `` + "`" + `sh terraform import linode_nodebalancer_config.http-foobar 1234567,7654321 ` + "`" + `` + "`" + `` + "`" + ` The Linode Guide, [Import Existing Infrastructure to Terraform](https://www.linode.com/docs/applications/configuration-management/import-existing-infrastructure-to-terraform/), offers resource importing examples for NodeBalancer Configs and other Linode resource types.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_nodebalancer_node",
			Category:         "Resources",
			ShortDescription: `Manages a Linode NodeBalancer Node.`,
			Description:      ``,
			Keywords: []string{
				"nodebalancer",
				"node",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "label",
					Description: `(Required) The label of the Linode NodeBalancer Node. This is for display purposes only.`,
				},
				resource.Attribute{
					Name:        "nodebalancer_id",
					Description: `(Required) The ID of the NodeBalancer to access.`,
				},
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required) The ID of the NodeBalancerConfig to access.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) The private IP Address where this backend can be reached. This must be a private IP address. - - -`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) The mode this NodeBalancer should use when sending traffic to this backend. If set to ` + "`" + `accept` + "`" + ` this backend is accepting traffic. If set to ` + "`" + `reject` + "`" + ` this backend will not receive traffic. If set to ` + "`" + `drain` + "`" + ` this backend will not receive new traffic, but connections already pinned to it will continue to be routed to it. (` + "`" + `accept` + "`" + `, ` + "`" + `reject` + "`" + `, ` + "`" + `drain` + "`" + `, ` + "`" + `backup` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) Used when picking a backend to serve a request and is not pinned to a single backend yet. Nodes with a higher weight will receive more traffic. (1-255). ## Attributes This resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current status of this node, based on the configured checks of its NodeBalancer Config. (` + "`" + `unknown` + "`" + `, ` + "`" + `UP` + "`" + `, ` + "`" + `DOWN` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "config_id",
					Description: `The ID of the NodeBalancerConfig this NodeBalancerNode is attached to.`,
				},
				resource.Attribute{
					Name:        "nodebalancer_id",
					Description: `The ID of the NodeBalancer this NodeBalancerNode is attached to. ## Import NodeBalancer Nodes can be imported using the NodeBalancer ` + "`" + `nodebalancer_id` + "`" + ` followed by the NodeBalancer Config ` + "`" + `config_id` + "`" + ` followed by the NodeBalancer Node ` + "`" + `id` + "`" + `, separated by a comma, e.g. ` + "`" + `` + "`" + `` + "`" + `sh terraform import linode_nodebalancer_node.https-foobar-1 1234567,7654321,9999999 ` + "`" + `` + "`" + `` + "`" + ` The Linode Guide, [Import Existing Infrastructure to Terraform](https://www.linode.com/docs/applications/configuration-management/import-existing-infrastructure-to-terraform/), offers resource importing examples for NodeBalancer Nodes and other Linode resource types.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_object_storage_bucket",
			Category:         "Resources",
			ShortDescription: `Manages a Linode Object Storage Bucket.`,
			Description:      ``,
			Keywords: []string{
				"object",
				"storage",
				"bucket",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster",
					Description: `(Required) The cluster of the Linode Object Storage Bucket.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Required) The label of the Linode Object Storage Bucket.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `(Optional) The Access Control Level of the bucket using a canned ACL string. See all ACL strings [in the Linode API v4 documentation](linode.com/docs/api/object-storage/#object-storage-bucket-access-update__request-body-schema).`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `(Optional) The access key to authenticate with.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Optional) The secret key to authenticate with.`,
				},
				resource.Attribute{
					Name:        "cors_enabled",
					Description: `(Optional) If true, the bucket will have CORS enabled for all origins.`,
				},
				resource.Attribute{
					Name:        "versioning",
					Description: `(Optional) Whether to enable versioning. Once you version-enable a bucket, it can never return to an unversioned state. You can, however, suspend versioning on that bucket. (Requires ` + "`" + `access_key` + "`" + ` and ` + "`" + `secret_key` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(Required) The Base64 encoded and PEM formatted SSL certificate.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Required) The private key associated with the TLS/SSL certificate. ### lifecycle_rule The following arguments are supported in the lifecycle_rule specification block:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The unique identifier for the rule.`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Optional) The object key prefix identifying one or more objects to which the rule applies.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Specifies whether the lifecycle rule is active.`,
				},
				resource.Attribute{
					Name:        "abort_incomplete_multipart_upload_days",
					Description: `(Optional) Specifies the number of days after initiating a multipart upload when the multipart upload must be completed.`,
				},
				resource.Attribute{
					Name:        "date",
					Description: `(Optional) Specifies the date after which you want the corresponding action to take effect.`,
				},
				resource.Attribute{
					Name:        "days",
					Description: `(Optional) Specifies the number of days after object creation when the specific rule action takes effect.`,
				},
				resource.Attribute{
					Name:        "expired_object_delete_marker",
					Description: `(Optional) On a versioned bucket (versioning-enabled or versioning-suspended bucket), you can add this element in the lifecycle configuration to direct Linode Object Storage to delete expired object delete markers. This cannot be specified with Days or Date in a Lifecycle Expiration Policy. ### noncurrent_version_expiration The following arguments are supported in the noncurrent_version_expiration specification block:`,
				},
				resource.Attribute{
					Name:        "days",
					Description: `(Required) Specifies the number of days non-current object versions expire. ## Import Linodes Object Storage Buckets can be imported using the resource ` + "`" + `id` + "`" + ` which is made of ` + "`" + `cluster:label` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `sh terraform import linode_object_storage_bucket.mybucket us-east-1:foobar ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_object_storage_key",
			Category:         "Resources",
			ShortDescription: `Manages a Linode Object Storage Key.`,
			Description:      ``,
			Keywords: []string{
				"object",
				"storage",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "label",
					Description: `(Required) The label given to this key. For display purposes only. - - -`,
				},
				resource.Attribute{
					Name:        "bucket_access",
					Description: `(Optional) Defines this key as a Limited Access Key. Limited Access Keys restrict this Object Storage key’s access to only the bucket(s) declared in this array and define their bucket-level permissions. Not providing this block will not limit this Object Storage Key. ### bucket_access The following arguments are supported in the bucket_access block:`,
				},
				resource.Attribute{
					Name:        "bucket_name",
					Description: `The unique label of the bucket to which the key will grant limited access.`,
				},
				resource.Attribute{
					Name:        "cluster",
					Description: `The Object Storage cluster where a bucket to which the key is granting access is hosted.`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `This Limited Access Key’s permissions for the selected bucket.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `This keypair's access key. This is not secret.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `This keypair's secret key.`,
				},
				resource.Attribute{
					Name:        "limited",
					Description: `Whether or not this key is a limited access key.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_object_storage_object",
			Category:         "Resources",
			ShortDescription: `Manages a Linode Object Storage Object.`,
			Description:      ``,
			Keywords: []string{
				"object",
				"storage",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) The name of the bucket to put the object in.`,
				},
				resource.Attribute{
					Name:        "cluster",
					Description: `(Required) The cluster the bucket is in.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) They name of the object once it is in the bucket.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Required) The secret key to authenitcate with.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `(Required) The access key to authenticate with.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional, conflicts with ` + "`" + `content` + "`" + ` and ` + "`" + `content_base64` + "`" + `) The path to a file that will be read and uploaded as raw bytes for the object content. The path must either be relative to the root module or absolute.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Optional, conflicts with ` + "`" + `source` + "`" + ` and ` + "`" + `content_base64` + "`" + `) Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.`,
				},
				resource.Attribute{
					Name:        "content_base64",
					Description: `(Optional, conflicts with ` + "`" + `source` + "`" + ` and ` + "`" + `content` + "`" + `) Base64-encoded data that will be decoded and uploaded as raw bytes for the object content. This allows safely uploading non-UTF8 binary data, but is recommended only for small content such as the result of the ` + "`" + `gzipbase64` + "`" + ` function with small text strings. For larger objects, use ` + "`" + `source` + "`" + ` to stream the content from a disk file.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `(Optional) The canned ACL to apply. (` + "`" + `private` + "`" + `, ` + "`" + `public-read` + "`" + `, ` + "`" + `authenticated-read` + "`" + `, ` + "`" + `public-read-write` + "`" + `, ` + "`" + `custom` + "`" + `) (defaults to ` + "`" + `private` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "cache_control",
					Description: `(Optional) Specifies caching behavior along the request/reply chain Read [w3c cache_control](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9) for further details.`,
				},
				resource.Attribute{
					Name:        "content_disposition",
					Description: `(Optional) Specifies presentational information for the object. Read [w3c content_disposition](http://www.w3.org/Protocols/rfc2616/rfc2616-sec19.html#sec19.5.1) for further information.`,
				},
				resource.Attribute{
					Name:        "content_encoding",
					Description: `(Optional) Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read [w3c content encoding](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.11) for further information.`,
				},
				resource.Attribute{
					Name:        "content_language",
					Description: `(Optional) The language the content is in e.g. en-US or en-GB.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `(Optional) A standard MIME type describing the format of the object data, e.g. application/octet-stream. All Valid MIME Types are valid for this input.`,
				},
				resource.Attribute{
					Name:        "website_redirect",
					Description: `(Optional) Specifies a target URL for website redirect.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `(Optional) Used to trigger updates. The only meaningful value is ` + "`" + `${filemd5("path/to/file")}` + "`" + ` (Terraform 0.11.12 or later) or ` + "`" + `${md5(file("path/to/file"))}` + "`" + ` (Terraform 0.11.11 or earlier).`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) A map of keys/values to provision metadata.`,
				},
				resource.Attribute{
					Name:        "force_destroy",
					Description: `(Optional) Allow the object to be deleted regardless of any legal hold or object lock (defaults to ` + "`" + `false` + "`" + `). ## Attributes Reference The following attributes are exported`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `A unique version ID value for the object.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "version_id",
					Description: `A unique version ID value for the object.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_rdns",
			Category:         "Resources",
			ShortDescription: `Manages the RDNS / PTR record for the IP Address associated with a Linode Instance.`,
			Description:      ``,
			Keywords: []string{
				"rdns",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "address",
					Description: `The Public IPv4 or IPv6 address that will receive the ` + "`" + `PTR` + "`" + ` record. A matching ` + "`" + `A` + "`" + ` or ` + "`" + `AAAA` + "`" + ` record must exist.`,
				},
				resource.Attribute{
					Name:        "rdns",
					Description: `The name of the RDNS address.`,
				},
				resource.Attribute{
					Name:        "wait_for_available",
					Description: `(Optional) If true, the RDNS assignment will be retried within the operation timeout period. ## Import Linodes RDNS resources can be imported using the address as the ` + "`" + `id` + "`" + `. ` + "`" + `` + "`" + `` + "`" + `sh terraform import linode_rdns.foo 123.123.123.123 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_sshkey",
			Category:         "Resources",
			ShortDescription: `Manages a Linode SSH Key.`,
			Description:      ``,
			Keywords: []string{
				"sshkey",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "label",
					Description: `A label for the SSH Key.`,
				},
				resource.Attribute{
					Name:        "ssh_key",
					Description: `The public SSH Key, which is used to authenticate to the root user of the Linodes you deploy. ## Attributes This resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The date this SSH Key was created. ## Import Linodes SSH Keys can be imported using the Linode SSH Key ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `sh terraform import linode_sshkey.mysshkey 1234567 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_stackscript",
			Category:         "Resources",
			ShortDescription: `Manages a Linode StackScript.`,
			Description:      ``,
			Keywords: []string{
				"stackscript",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "label",
					Description: `(Required) The StackScript's label is for display purposes only.`,
				},
				resource.Attribute{
					Name:        "script",
					Description: `(Required) The script to execute when provisioning a new Linode with this StackScript.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) A description for the StackScript.`,
				},
				resource.Attribute{
					Name:        "images",
					Description: `(Required) An array of Image IDs representing the Images that this StackScript is compatible for deploying with. - - -`,
				},
				resource.Attribute{
					Name:        "rev_note",
					Description: `(Optional) This field allows you to add notes for the set of revisions made to this StackScript.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `(Optional) This determines whether other users can use your StackScript. Once a StackScript is made public, it cannot be made private.`,
				},
				resource.Attribute{
					Name:        "deployments_active",
					Description: `Count of currently active, deployed Linodes created from this StackScript.`,
				},
				resource.Attribute{
					Name:        "user_gravatar_id",
					Description: `The Gravatar ID for the User who created the StackScript.`,
				},
				resource.Attribute{
					Name:        "deployments_total",
					Description: `The total number of times this StackScript has been deployed.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The User who created the StackScript.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The date this StackScript was created.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `The date this StackScript was updated.`,
				},
				resource.Attribute{
					Name:        "user_defined_fields",
					Description: `This is a list of fields defined with a special syntax inside this StackScript that allow for supplying customized parameters during deployment.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `A human-readable label for the field that will serve as the input prompt for entering the value during deployment.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the field.`,
				},
				resource.Attribute{
					Name:        "example",
					Description: `An example value for the field.`,
				},
				resource.Attribute{
					Name:        "one_of",
					Description: `A list of acceptable single values for the field.`,
				},
				resource.Attribute{
					Name:        "many_of",
					Description: `A list of acceptable values for the field in any quantity, combination or order.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `The default value. If not specified, this value will be used. ## Import Linodes StackScripts can be imported using the Linode StackScript ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `sh terraform import linode_stackscript.mystackscript 1234567 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_token",
			Category:         "Resources",
			ShortDescription: `Manages a Linode Token.`,
			Description:      ``,
			Keywords: []string{
				"token",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "label",
					Description: `A label for the Token.`,
				},
				resource.Attribute{
					Name:        "scopes",
					Description: `The scopes this token was created with. These define what parts of the Account the token can be used to access. Many command-line tools, such as the Linode CLI, require tokens with access to`,
				},
				resource.Attribute{
					Name:        "expiry",
					Description: `When this token will expire. Personal Access Tokens cannot be renewed, so after this time the token will be completely unusable and a new token will need to be generated. Tokens may be created with 'null' as their expiry and will never expire unless revoked. ## Attributes This resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `The token used to access the API.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The date this Token was created. ## Import Linodes Tokens can be imported using the Linode Token ` + "`" + `id` + "`" + `, e.g. The secret token will not be imported. ` + "`" + `` + "`" + `` + "`" + `sh terraform import linode_token.mytoken 1234567 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_user",
			Category:         "Resources",
			ShortDescription: `Manages a Linode User.`,
			Description:      ``,
			Keywords: []string{
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "username",
					Description: `(required) The username of the user.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(required) The email address of the user.`,
				},
				resource.Attribute{
					Name:        "restricted",
					Description: `(optional) If true, this user will only have explicit permissions granted.`,
				},
				resource.Attribute{
					Name:        "domain_grant",
					Description: `(optional) The domains the user has permissions access to.`,
				},
				resource.Attribute{
					Name:        "firewall_grant",
					Description: `(optional) The firewalls the user has permissions access to.`,
				},
				resource.Attribute{
					Name:        "image_grant",
					Description: `(optional) The images the user has permissions access to.`,
				},
				resource.Attribute{
					Name:        "linode_grant",
					Description: `(optional) The Linodes the user has permissions access to.`,
				},
				resource.Attribute{
					Name:        "longview_grant",
					Description: `(optional) The longview the user has permissions access to.`,
				},
				resource.Attribute{
					Name:        "nodebalancer_grant",
					Description: `(optional) The NodeBalancers the user has permissions access to.`,
				},
				resource.Attribute{
					Name:        "stackscript_grant",
					Description: `(optional) The StackScripts the user has permissions access to.`,
				},
				resource.Attribute{
					Name:        "volume_grant",
					Description: `(optional) The volumes the user has permissions access to. ## Attributes Reference In addition to all the arguments above, the following attributes are exported.`,
				},
				resource.Attribute{
					Name:        "tfa_enabled",
					Description: `Whether the user has two-factor-authentication enabled.`,
				},
				resource.Attribute{
					Name:        "ssh_keys",
					Description: `A list of the User's SSH keys. ## Global Grants`,
				},
				resource.Attribute{
					Name:        "account-access",
					Description: `(optional) The level of access this User has to Account-level actions, like billing information. (` + "`" + `read_only` + "`" + `, ` + "`" + `read_write` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "add_domains",
					Description: `(optional) If true, this User may add Domains.`,
				},
				resource.Attribute{
					Name:        "add_firewalls",
					Description: `(optional) If true, this User may add Firewalls.`,
				},
				resource.Attribute{
					Name:        "add_images",
					Description: `(optional) If true, this User may add Images.`,
				},
				resource.Attribute{
					Name:        "add_linodes",
					Description: `(optional) If true, this User may create Linodes.`,
				},
				resource.Attribute{
					Name:        "add_longview",
					Description: `(optional) If true, this User may create Longview clients and view the current plan.`,
				},
				resource.Attribute{
					Name:        "add_nodebalancers",
					Description: `(optional) If true, this User may add NodeBalancers.`,
				},
				resource.Attribute{
					Name:        "add_stackscripts",
					Description: `(optional) If true, this User may add StackScripts.`,
				},
				resource.Attribute{
					Name:        "cancel_account",
					Description: `(optional) If true, this User may cancel the entire Account.`,
				},
				resource.Attribute{
					Name:        "longview_subscription",
					Description: `(optional) If true, this User may manage the Account’s Longview subscription. ## Entity Grants`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(required) The ID of the entity this grant applies to.`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `(required) The level of access this User has to this entity. (` + "`" + `read_only` + "`" + `, ` + "`" + `read_write` + "`" + `)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "tfa_enabled",
					Description: `Whether the user has two-factor-authentication enabled.`,
				},
				resource.Attribute{
					Name:        "ssh_keys",
					Description: `A list of the User's SSH keys. ## Global Grants`,
				},
				resource.Attribute{
					Name:        "account-access",
					Description: `(optional) The level of access this User has to Account-level actions, like billing information. (` + "`" + `read_only` + "`" + `, ` + "`" + `read_write` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "add_domains",
					Description: `(optional) If true, this User may add Domains.`,
				},
				resource.Attribute{
					Name:        "add_firewalls",
					Description: `(optional) If true, this User may add Firewalls.`,
				},
				resource.Attribute{
					Name:        "add_images",
					Description: `(optional) If true, this User may add Images.`,
				},
				resource.Attribute{
					Name:        "add_linodes",
					Description: `(optional) If true, this User may create Linodes.`,
				},
				resource.Attribute{
					Name:        "add_longview",
					Description: `(optional) If true, this User may create Longview clients and view the current plan.`,
				},
				resource.Attribute{
					Name:        "add_nodebalancers",
					Description: `(optional) If true, this User may add NodeBalancers.`,
				},
				resource.Attribute{
					Name:        "add_stackscripts",
					Description: `(optional) If true, this User may add StackScripts.`,
				},
				resource.Attribute{
					Name:        "cancel_account",
					Description: `(optional) If true, this User may cancel the entire Account.`,
				},
				resource.Attribute{
					Name:        "longview_subscription",
					Description: `(optional) If true, this User may manage the Account’s Longview subscription. ## Entity Grants`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(required) The ID of the entity this grant applies to.`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `(required) The level of access this User has to this entity. (` + "`" + `read_only` + "`" + `, ` + "`" + `read_write` + "`" + `)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "linode_volume",
			Category:         "Resources",
			ShortDescription: `Manages a Linode Volume.`,
			Description:      ``,
			Keywords: []string{
				"volume",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "label",
					Description: `(Required) The label of the Linode Volume`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The region where this volume will be deployed. Examples are ` + "`" + `"us-east"` + "`" + `, ` + "`" + `"us-west"` + "`" + `, ` + "`" + `"ap-south"` + "`" + `, etc. See all regions [here](https://api.linode.com/v4/regions).`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) Size of the Volume in GB.`,
				},
				resource.Attribute{
					Name:        "linode_id",
					Description: `(Optional) The ID of a Linode Instance where the Volume should be attached.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A list of tags applied to this object. Tags are for organizational purposes only. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 10 mins) Used when creating the volume (until the volume is reaches the initial ` + "`" + `active` + "`" + ` state)`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 20 mins) Used when updating the volume when necessary during update - e.g. when resizing the volume`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 10 mins) Used when deleting the volume ## Attributes This resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the Linode Volume. (` + "`" + `creating` + "`" + `, ` + "`" + `active` + "`" + `, ` + "`" + `resizing` + "`" + `, ` + "`" + `contact_support` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "filesystem_path",
					Description: `The full filesystem path for the Volume based on the Volume's label. The path is "/dev/disk/by-id/scsi-0Linode_Volume_" + the Volume label ## Import Linodes Volumes can be imported using the Linode Volume ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `sh terraform import linode_volume.myvolume 1234567 ` + "`" + `` + "`" + `` + "`" + ` The Linode Guide, [Import Existing Infrastructure to Terraform](https://www.linode.com/docs/applications/configuration-management/import-existing-infrastructure-to-terraform/), offers resource importing examples for Block Storage Volumes and other Linode resource types.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"linode_database_access_controls": 0,
		"linode_database_mysql":           1,
		"linode_domain":                   2,
		"linode_domain_record":            3,
		"linode_firewall":                 4,
		"linode_firewall_device":          5,
		"linode_image":                    6,
		"linode_instance":                 7,
		"linode_instance_ip":              8,
		"linode_ipv6_range":               9,
		"linode_lke_cluster":              10,
		"linode_nodebalancer":             11,
		"linode_nodebalancer_config":      12,
		"linode_nodebalancer_node":        13,
		"linode_object_storage_bucket":    14,
		"linode_object_storage_key":       15,
		"linode_object_storage_object":    16,
		"linode_rdns":                     17,
		"linode_sshkey":                   18,
		"linode_stackscript":              19,
		"linode_token":                    20,
		"linode_user":                     21,
		"linode_volume":                   22,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
