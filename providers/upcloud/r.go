package upcloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "upcloud_firewall_rules",
			Category:         "Resources",
			ShortDescription: `Provides a list UpCloud firewall rules`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"rules",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "server_id",
					Description: `(Required) The unique id of the server to be protected the firewall rules.`,
				},
				resource.Attribute{
					Name:        "firewall_rule",
					Description: `(Required) A customisable list of firewall rules that are to be attached to the server. See [Firewall Rules](#firewall-rules) below for more details. ### Firewall Rules Each of the ` + "`" + `firewall_rule` + "`" + ` blocks represent a single firewall rule that is the be attached to the server specified through ` + "`" + `server_id` + "`" + `. The order of the ` + "`" + `firewall_rule` + "`" + ` blocks is used as the order to attach each rule to the server. Moving the block will change the order they are applied in. Each ` + "`" + `firewall_rule` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) Action to take if the rule conditions are met. Accepted value ` + "`" + `accept` + "`" + ` or ` + "`" + `drop` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Freeform comment string for the rule. Accepted length 0-250 characters.`,
				},
				resource.Attribute{
					Name:        "destination_address_end",
					Description: `(Optional) The destination address range ends from this address. Required if using ` + "`" + `destination_address_start` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "destination_address_start",
					Description: `(Optional) The destination address range starts from this address. Required if using ` + "`" + `destination_address_end` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "destination_port_end",
					Description: `(Optional) The destination port range ends from this port number. Required if using ` + "`" + `destination_port_start` + "`" + `. Accepted value 1-65535.`,
				},
				resource.Attribute{
					Name:        "destination_port_start",
					Description: `(Optional) The destination port range starts from this port number. Required if using ` + "`" + `destination_port_end` + "`" + `. Accepted valye 1-65535.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Required) - The direction of network traffic this rule will be applied to.`,
				},
				resource.Attribute{
					Name:        "family",
					Description: `(Required) - The address family of new firewall rule, Accepted value ` + "`" + `IPv4` + "`" + ` or ` + "`" + `IPv6` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "icmp_type",
					Description: `(Optional) - The ICMP type. Accepted value 0-255.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) - The protocol this rule will be applied to. Accepted values ` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + ` or ` + "`" + `icmp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "source_address_end",
					Description: `(Optional) The source address range ends from this address. Required if using ` + "`" + `source_address_start` + "`" + ``,
				},
				resource.Attribute{
					Name:        "source_address_start",
					Description: `(Optional) The source address range starts from this address. Required if using ` + "`" + `source_address_end` + "`" + ``,
				},
				resource.Attribute{
					Name:        "source_port_end",
					Description: `(Optional) The source port range ends from this port number. Required if using ` + "`" + `source_port_start` + "`" + `. Accepted value 1-65535.`,
				},
				resource.Attribute{
					Name:        "source_port_start",
					Description: `(Optional) The source port range starts from this port number. Required if using ` + "`" + `source_port_end` + "`" + `. Accepted value 1-65535. ## Import Existing UpCloud firewall rules can be imported into the current Terraform state through the server id UUID. ` + "`" + `` + "`" + `` + "`" + `hcl terraform import upcloud_firewall_rules.my_example_rules 049d7ca2-757e-4fb1-a833-f87ee056547a ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_floating_ip_address",
			Category:         "Resources",
			ShortDescription: `Provides an UpCloud Floating IP Address`,
			Description:      ``,
			Keywords: []string{
				"floating",
				"ip",
				"address",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "access",
					Description: `(Optional) The IP address access type (one of ` + "`" + `utility` + "`" + ` or ` + "`" + `public` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "family",
					Description: `(Optional) The IP address family (one of ` + "`" + `IPv4` + "`" + ` or ` + "`" + `IPv6` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `(Optional) MAC address of server interface to assign address to ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `An UpCloud assigned IP Address`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `Zone of address, required when assigning a detached floating IP address. Required when defining a detached floating IP address resource. ## Import An existing UpCloud Floating IP address can be imported into the current Terraform state through the assigned IP Address. ` + "`" + `` + "`" + `` + "`" + `hcl terraform import upcloud_floating_ip_address.my_new_floating_address 94.237.114.205 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_address",
					Description: `An UpCloud assigned IP Address`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `Zone of address, required when assigning a detached floating IP address. Required when defining a detached floating IP address resource. ## Import An existing UpCloud Floating IP address can be imported into the current Terraform state through the assigned IP Address. ` + "`" + `` + "`" + `` + "`" + `hcl terraform import upcloud_floating_ip_address.my_new_floating_address 94.237.114.205 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_network",
			Category:         "Resources",
			ShortDescription: `Provides an UpCloud network`,
			Description:      ``,
			Keywords: []string{
				"network",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name to be assigned to the network.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `The UpCloud zone this network will reside in.`,
				},
				resource.Attribute{
					Name:        "router",
					Description: `A router to connect networks together.`,
				},
				resource.Attribute{
					Name:        "ip_network",
					Description: `A block containing the IP details of the network. ### ip_network`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) The CIDR range of the network. Updating forces a new ` + "`" + `network` + "`" + ` resource.`,
				},
				resource.Attribute{
					Name:        "dhcp",
					Description: `(Required) Is DHCP enabled?`,
				},
				resource.Attribute{
					Name:        "dhcp_default_route",
					Description: `Is the ` + "`" + `gateway` + "`" + ` the DHCP default route?`,
				},
				resource.Attribute{
					Name:        "dhcp_dns",
					Description: `A set of the DNS servers given by DHCP.`,
				},
				resource.Attribute{
					Name:        "family",
					Description: `(Required) The IP address family (one of ` + "`" + `IPv4` + "`" + ` or ` + "`" + `IPv6` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `The gateway address given by DHCP. ## Attributes Reference In addition to the arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type assigned to the network.`,
				},
				resource.Attribute{
					Name:        "servers",
					Description: `A set of attached servers ### servers`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier of the server.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `The short description of the server. ## Import Existing UpCloud networks can be imported into the current Terraform state through the assigned UUID. ` + "`" + `` + "`" + `` + "`" + `hcl terraform import upcloud_network.my_example_network 03e44422-07b8-4798-a597-c8eab1fa64df ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `The type assigned to the network.`,
				},
				resource.Attribute{
					Name:        "servers",
					Description: `A set of attached servers ### servers`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier of the server.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `The short description of the server. ## Import Existing UpCloud networks can be imported into the current Terraform state through the assigned UUID. ` + "`" + `` + "`" + `` + "`" + `hcl terraform import upcloud_network.my_example_network 03e44422-07b8-4798-a597-c8eab1fa64df ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_object_storage",
			Category:         "Resources",
			ShortDescription: `Provides a UpCloud Object Storage management`,
			Description:      ``,
			Keywords: []string{
				"object",
				"storage",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "size",
					Description: `(Required) The maximum amount of storage to allocate to the object storage instance in gigabytes. Valid values are ` + "`" + `250` + "`" + `, ` + "`" + `500` + "`" + ` and ` + "`" + `1000` + "`" + ``,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Object Storage instance, this value forms part of the url used to access the storage instance so must conform to host naming rules.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) - The zone in which the server will be hosted, e.g. fi-hel2. See [Zones API](https://developers.upcloud.com/1.3/5-zones/)`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `(Required) - The access key/username used to access the storage instance. Must be a string between 4 and 255 characters in length.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Required) - The secret key/password used to access the storage instance. Must be a string between 8 and 255 characters in length.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) - A user defined string containing test to associate with the storage instance. ### Bucket Each of ` + "`" + `bucket` + "`" + ` blocks represents a single bucket in the object storage instance and supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the bucket.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_router",
			Category:         "Resources",
			ShortDescription: `Provides an UpCloud router`,
			Description:      ``,
			Keywords: []string{
				"router",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name to be assigned to the router ## Attributes Reference In addition to the arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type assigned to the router.`,
				},
				resource.Attribute{
					Name:        "attached_networks",
					Description: `A list of network UUID that are attached through this router. ## Import Existing UpCloud routers can be imported into the current Terraform state through the assigned UUID. ` + "`" + `` + "`" + `` + "`" + `hcl terraform import upcloud_router.my_example_router 049d7ca2-757e-4fb1-a833-f87ee056547a ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `The type assigned to the router.`,
				},
				resource.Attribute{
					Name:        "attached_networks",
					Description: `A list of network UUID that are attached through this router. ## Import Existing UpCloud routers can be imported into the current Terraform state through the assigned UUID. ` + "`" + `` + "`" + `` + "`" + `hcl terraform import upcloud_router.my_example_router 049d7ca2-757e-4fb1-a833-f87ee056547a ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_server",
			Category:         "Resources",
			ShortDescription: `Provides an UpCloud service`,
			Description:      ``,
			Keywords: []string{
				"server",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_storage",
			Category:         "Resources",
			ShortDescription: `Provides an UpCloud Storage`,
			Description:      ``,
			Keywords: []string{
				"storage",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "size",
					Description: `(Required) The size of the storage in gigabytes`,
				},
				resource.Attribute{
					Name:        "tier",
					Description: `(Optional) The storage tier to use`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `(Required) A short, informative description`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) The zone in which the storage will be created`,
				},
				resource.Attribute{
					Name:        "backup_rule",
					Description: `(Optional) The criteria to backup the storage`,
				},
				resource.Attribute{
					Name:        "import",
					Description: `(Optional) Details of the external data to import`,
				},
				resource.Attribute{
					Name:        "clone",
					Description: `(Optional) Details of another storage device to clone The ` + "`" + `backup_rule` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Required) The weekday when the backup is created`,
				},
				resource.Attribute{
					Name:        "time",
					Description: `(Required) The time of day when the backup is created`,
				},
				resource.Attribute{
					Name:        "retention",
					Description: `(Required) The number of days before a backup is automatically deleted The ` + "`" + `import` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Required) The source type (one of ` + "`" + `direct_upload` + "`" + ` or ` + "`" + `http_import` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "source_location",
					Description: `(Required) For ` + "`" + `direct_upload` + "`" + ` the path to a local file. For ` + "`" + `http_import` + "`" + ` an accessible URL.`,
				},
				resource.Attribute{
					Name:        "source_hash",
					Description: `(Optional) The hash of ` + "`" + `source_location` + "`" + `. This is used to indicate that ` + "`" + `source_location` + "`" + ` has changed. It is not used for verification. The ` + "`" + `clone` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The unique identifier of another storage device to clone. ## Import Existing UpCloud storage can be imported into the current Terraform state through the assigned UUID. ` + "`" + `` + "`" + `` + "`" + `hcl terraform import upcloud_storage.example_storage 0128ae5a-91dd-4ebf-bd1e-304c47f2c652 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_tag",
			Category:         "Resources",
			ShortDescription: `Provides an UpCloud Tag`,
			Description:      ``,
			Keywords: []string{
				"tag",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The value representing the tag`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Freeform comment string for the host`,
				},
				resource.Attribute{
					Name:        "servers",
					Description: `A collection of servers that have been assigned the tag. ## Import An Existing UpCloud Tag can be imported into the current Terraform state through the tag name. ` + "`" + `` + "`" + `` + "`" + `hcl terraform import upcloud_tag.example_tag dev ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"upcloud_firewall_rules":      0,
		"upcloud_floating_ip_address": 1,
		"upcloud_network":             2,
		"upcloud_object_storage":      3,
		"upcloud_router":              4,
		"upcloud_server":              5,
		"upcloud_storage":             6,
		"upcloud_tag":                 7,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
