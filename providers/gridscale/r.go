package gridscale

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "gridscale_ipv4",
			Category:         "Resources",
			ShortDescription: `Manages an IPv4 address in gridscale.`,
			Description:      ``,
			Keywords: []string{
				"ipv4",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The human-readable name of the object. It supports the full UTF-8 charset, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `(Optional) Helps to identify which datacenter an object belongs to. Frankfurt is the default.`,
				},
				resource.Attribute{
					Name:        "failover",
					Description: `(Optional) Sets failover mode for this IP. If true, then this IP is no longer available for DHCP and can no longer be related to any server.`,
				},
				resource.Attribute{
					Name:        "reverse_dns",
					Description: `(Optional) Defines the reverse DNS entry for the IP Address (PTR Resource Record).`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) List of labels in the format [ "label1", "label2" ]. ## Attributes This resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "failover",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "reverse_dns",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Defines the IP Address.`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `The network address and the subnet.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `status indicates the status of the object.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time the object was created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `Defines the date and time of the last object change.`,
				},
				resource.Attribute{
					Name:        "location_country",
					Description: `Formatted by the 2 digit country code (ISO 3166-2) of the host country.`,
				},
				resource.Attribute{
					Name:        "location_iata",
					Description: `Uses IATA airport code, which works as a location identifier.`,
				},
				resource.Attribute{
					Name:        "location_name",
					Description: `The location name.`,
				},
				resource.Attribute{
					Name:        "delete_block",
					Description: `Defines if the object is administratively blocked. If true, it can not be deleted by the user.`,
				},
				resource.Attribute{
					Name:        "usage_in_minutes",
					Description: `The amount of minutes the IP address has been in use.`,
				},
				resource.Attribute{
					Name:        "current_price",
					Description: `The price for the current period since the last bill.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_ipv6",
			Category:         "Resources",
			ShortDescription: `Manages an IPv6 address in gridscale.`,
			Description:      ``,
			Keywords: []string{
				"ipv6",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The human-readable name of the object. It supports the full UTF-8 charset, with a maximum of 66 characters.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `(Optional) Helps to identify which datacenter an object belongs to. Frankfurt is the default.`,
				},
				resource.Attribute{
					Name:        "failover",
					Description: `(Optional) Sets failover mode for this IP. If true, then this IP is no longer available for DHCP and can no longer be related to any server.`,
				},
				resource.Attribute{
					Name:        "reverse_dns",
					Description: `(Optional) Defines the reverse DNS entry for the IP Address (PTR Resource Record).`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) List of labels in the format [ "label1", "label2" ]. ## Attributes This resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "failover",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "reverse_dns",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Defines the IP Address.`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `The network address and the subnet.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `status indicates the status of the object.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time the object was created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `Defines the date and time of the last object change.`,
				},
				resource.Attribute{
					Name:        "location_country",
					Description: `Formatted by the 2 digit country code (ISO 3166-2) of the host country.`,
				},
				resource.Attribute{
					Name:        "location_iata",
					Description: `Uses IATA airport code, which works as a location identifier.`,
				},
				resource.Attribute{
					Name:        "location_name",
					Description: `The location name.`,
				},
				resource.Attribute{
					Name:        "delete_block",
					Description: `Defines if the object is administratively blocked. If true, it can not be deleted by the user.`,
				},
				resource.Attribute{
					Name:        "usage_in_minutes",
					Description: `The amount of minutes the IP address has been in use.`,
				},
				resource.Attribute{
					Name:        "current_price",
					Description: `The price for the current period since the last bill.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_network",
			Category:         "Resources",
			ShortDescription: `Manages a network in gridscale.`,
			Description:      ``,
			Keywords: []string{
				"network",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The human-readable name of the object. It supports the full UTF-8 charset, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `(Optional) Helps to identify which datacenter an object belongs to. Frankfurt is the default.`,
				},
				resource.Attribute{
					Name:        "l2security",
					Description: `(Optional) Defines information about MAC spoofing protection (filters layer2 and ARP traffic based on MAC source). It can only be (de-)activated on a private network - the public network always has l2security enabled. It will be true if the network is public, and false if the network is private.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) List of labels in the format [ "label1", "label2" ]. ## Attributes This resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "l2security",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `status indicates the status of the object.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time the object was created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `Defines the date and time of the last object change.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `The type of this network, can be mpls, breakout or network.`,
				},
				resource.Attribute{
					Name:        "location_country",
					Description: `Formatted by the 2 digit country code (ISO 3166-2) of the host country.`,
				},
				resource.Attribute{
					Name:        "location_iata",
					Description: `Uses IATA airport code, which works as a location identifier.`,
				},
				resource.Attribute{
					Name:        "location_name",
					Description: `The location name.`,
				},
				resource.Attribute{
					Name:        "public_net",
					Description: `Is the network public or not.`,
				},
				resource.Attribute{
					Name:        "delete_block",
					Description: `If deleting this network is allowed.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_server",
			Category:         "Resources",
			ShortDescription: `Manages a server in gridscale.`,
			Description:      ``,
			Keywords: []string{
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The human-readable name of the object. It supports the full UTF-8 charset, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "cores",
					Description: `(Required) The number of server cores.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `(Required) The amount of server memory in GB.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `(Optional) Helps to identify which datacenter an object belongs to. Frankfurt is the default.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) List of labels in the format [ "label1", "label2" ].`,
				},
				resource.Attribute{
					Name:        "hardware_profile",
					Description: `(Optional) The hardware profile of the Server. Options are default, legacy, nested, cisco_csr, sophos_utm, f5_bigip and q35 at the moment of writing. Check the`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `(Optional) The UUID of the IPv4 address of the server. When this option is set, the server will automatically be connected to the public network, giving it access to the internet.`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `(Optional) The UUID of the IPv6 address of the server. When this option is set, the server will automatically be connected to the public network, giving it access to the internet.`,
				},
				resource.Attribute{
					Name:        "isoimage",
					Description: `(Optional) The UUID of an ISO image in gridscale. The server will automatically boot from the ISO if one was added. The UUIDs of ISO images can be found in [the expert panel](https://my.gridscale.io/Expert/ISOImage).`,
				},
				resource.Attribute{
					Name:        "power",
					Description: `(Optional) The power state of the server. Set this to true to will boot the server, false will shut it down.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) Defines which Availability-Zone the Server is placed.`,
				},
				resource.Attribute{
					Name:        "storage",
					Description: `(Optional) Connects a storage to the server.`,
				},
				resource.Attribute{
					Name:        "object_uuid",
					Description: `(Required) The object UUID or id of the storage.`,
				},
				resource.Attribute{
					Name:        "bootdevice",
					Description: `(Optional) Make this storage the boot device.`,
				},
				resource.Attribute{
					Name:        "storage",
					Description: `(Optional) Connects a storage to the server.`,
				},
				resource.Attribute{
					Name:        "object_uuid",
					Description: `(Required) The object UUID or id of the network.`,
				},
				resource.Attribute{
					Name:        "bootdevice",
					Description: `(Optional) Make this network the boot device. This can only be set for one network. ## Attributes This resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cores",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "hardware_profile",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "storages",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "networks",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "power",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `status indicates the status of the object.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time the object was created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `Defines the date and time of the last object change.`,
				},
				resource.Attribute{
					Name:        "location_country",
					Description: `Formatted by the 2 digit country code (ISO 3166-2) of the host country.`,
				},
				resource.Attribute{
					Name:        "location_iata",
					Description: `Uses IATA airport code, which works as a location identifier.`,
				},
				resource.Attribute{
					Name:        "location_name",
					Description: `The location name.`,
				},
				resource.Attribute{
					Name:        "current_price",
					Description: `The price for the current period since the last bill.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_sshkey",
			Category:         "Resources",
			ShortDescription: `Manages an SSH public key in gridscale.`,
			Description:      ``,
			Keywords: []string{
				"sshkey",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The human-readable name of the object. It supports the full UTF-8 charset, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "sshkey",
					Description: `(Required) This is the OpenSSH public key string (all key types are supported => ed25519, ecdsa, dsa, rsa, rsa1).`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) List of labels in the format [ "label1", "label2" ]. ## Attributes This resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "sshkey",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `status indicates the status of the object.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time the object was created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `Defines the date and time of the last object change.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_storage",
			Category:         "Resources",
			ShortDescription: `Manages a storage in gridscale.`,
			Description:      ``,
			Keywords: []string{
				"storage",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The human-readable name of the object. It supports the full UTF-8 charset, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `(Required) required (integer - minimum: 1 - maximum: 4096).`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `(Optional) (one of storage, storage_high, storage_insane).`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `(Optional) Helps to identify which datacenter an object belongs to. Frankfurt is the default.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) List of labels in the format [ "label1", "label2" ].`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Optional) List of labels in the format [ "label1", "label2" ].`,
				},
				resource.Attribute{
					Name:        "template_uuid",
					Description: `(Required) The UUID of a template. This can be found in the [expert panel](https://my.gridscale.io/Expert/Template) by clicking more on the template or by using a gridscale_tempalte datasource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The root (Linux) or Administrator (Windows) password to set for the installed storage. Valid only for public templates. The password has to be either plaintext or a crypt string (modular crypt format - MCF)..`,
				},
				resource.Attribute{
					Name:        "password_type",
					Description: `(Optional) (one of plain, crypt) Required if password is set (ignored for private templates and public Windows templates).`,
				},
				resource.Attribute{
					Name:        "sshkeys",
					Description: `(Optional) (array of any - minItems: 0) Public linux templates only! The UUIDs of SSHkeys to install for the root user.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional) The hostname of the installed server (ignored for private templates and public windows templates). ~>`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `status indicates the status of the object.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time the object was created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `Defines the date and time of the last object change.`,
				},
				resource.Attribute{
					Name:        "location_country",
					Description: `Formatted by the 2 digit country code (ISO 3166-2) of the host country.`,
				},
				resource.Attribute{
					Name:        "location_iata",
					Description: `Uses IATA airport code, which works as a location identifier.`,
				},
				resource.Attribute{
					Name:        "location_name",
					Description: `The location name.`,
				},
				resource.Attribute{
					Name:        "license_product_no",
					Description: `If a template has been used that requires a license key (e.g. Windows Servers) this shows the product_no of the license (see the /prices endpoint for more details).`,
				},
				resource.Attribute{
					Name:        "last_used_template",
					Description: `Indicates the UUID of the last used template on this storage (inherited from snapshots).`,
				},
				resource.Attribute{
					Name:        "usage_in_minutes",
					Description: `The amount of minutes the IP address has been in use.`,
				},
				resource.Attribute{
					Name:        "current_price",
					Description: `The price for the current period since the last bill.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"gridscale_ipv4":    0,
		"gridscale_ipv6":    1,
		"gridscale_network": 2,
		"gridscale_server":  3,
		"gridscale_sshkey":  4,
		"gridscale_storage": 5,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
