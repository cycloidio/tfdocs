package aws

import (
	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*Resource{

		&resource.Resource{
			Name:             "",
			Type:             "opc_compute_acl",
			Category:         "Compute Classic Resources",
			ShortDescription: `Creates and manages an ACL in an Oracle Cloud Infrastructure Compute Classic identity domain.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"classic",
				"acl",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the ACL.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Enables or disables the ACL. Set to true by default.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the ACL.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) List of tags that may be applied to the ACL. In addition to the above, the following values are exported:`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `The Uniform Resource Identifier for the ACL ## Import ACL's can be imported using the ` + "`" + `resource name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ terraform import opc_compute_acl.acl1 example ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opc_compute_image_list",
			Category:         "Compute Classic Resources",
			ShortDescription: `Creates and manages an Image List in an Oracle Cloud Infrastructure Compute Classic identity domain.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"classic",
				"image",
				"list",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Image List.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) A description of the Image List.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `(Required) The image list entry to be used, by default, when launching instances using this image list. Defaults to ` + "`" + `1` + "`" + `. ## Import Image List's can be imported using the ` + "`" + `resource name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ terraform import opc_compute_image_list.imagelist1 example ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opc_compute_image_list_entry",
			Category:         "Compute Classic Resources",
			ShortDescription: `Creates and manages an Image List Entry in an Oracle Cloud Infrastructure Compute Classic identity domain.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"classic",
				"image",
				"list",
				"entry",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Image List.`,
				},
				resource.Attribute{
					Name:        "machine_images",
					Description: `(Required) An array of machine images.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The unique version of the image list entry, as an integer.`,
				},
				resource.Attribute{
					Name:        "attributes",
					Description: `(Optional) JSON String of optional data that will be passed to an instance of this machine image when it is launched. ## Attributes Reference In addition to the above arguments, the following attributes are exported`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `The Unique Resource Identifier for the Image List Entry. ## Import Image List's can be imported using the Name of the Image List, along with the Version of the Image List Entry, delimited via the ` + "`" + `|` + "`" + ` character, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ terraform import opc_compute_image_list_entry.entry1 my_image_list|2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "uri",
					Description: `The Unique Resource Identifier for the Image List Entry. ## Import Image List's can be imported using the Name of the Image List, along with the Version of the Image List Entry, delimited via the ` + "`" + `|` + "`" + ` character, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ terraform import opc_compute_image_list_entry.entry1 my_image_list|2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opc_compute_instance",
			Category:         "Compute Classic Resources",
			ShortDescription: `Creates and manages an instance in an Oracle Cloud Infrastructure Compute Classic identity domain.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"classic",
				"instance",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the instance.`,
				},
				resource.Attribute{
					Name:        "shape",
					Description: `(Required) The shape of the instance, e.g. ` + "`" + `oc4` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_attributes",
					Description: `(Optional) A JSON string of custom attributes. See [Attributes](#attributes) below for more information.`,
				},
				resource.Attribute{
					Name:        "boot_order",
					Description: `(Optional) The index number of the bootable storage volume, presented as a list, that should be used to boot the instance. The only valid value is ` + "`" + `[1]` + "`" + `. If you set this attribute, you must also specify a bootable storage volume with index number 1 in the volume sub-parameter of storage_attachments. When you specify boot_order, you don't need to specify the imagelist attribute, because the instance is booted using the image on the specified bootable storage volume. If you specify both boot_order and imagelist, the imagelist attribute is ignored.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional) The host name assigned to the instance. On an Oracle Linux instance, this host name is displayed in response to the hostname command. Only relative DNS is supported. The domain name is suffixed to the host name that you specify. The host name must not end with a period. If you don't specify a host name, then a name is generated automatically.`,
				},
				resource.Attribute{
					Name:        "image_list",
					Description: `(Optional) The imageList of the instance, e.g. ` + "`" + `/oracle/public/oel_6.4_2GB_v1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Optional) The label to apply to the instance.`,
				},
				resource.Attribute{
					Name:        "desired_state",
					Description: `(Optional) Set the desire state of the instance to ` + "`" + `running` + "`" + ` (default) or ` + "`" + `shutdown` + "`" + `. You can use this request to shut down and restart individual instances which use a persistent bootable storage volume.`,
				},
				resource.Attribute{
					Name:        "networking_info",
					Description: `(Optional) Information pertaining to an individual network interface to be created and attached to the instance. If left unspecified, the instance will be created within the ` + "`" + `shared_network` + "`" + `. See [Networking Info](#networking-info) below for more information.`,
				},
				resource.Attribute{
					Name:        "storage",
					Description: `(Optional) Information pertaining to an individual storage attachment to be created during instance creation. Please see [Storage Attachments](#storage-attachments) below for more information.`,
				},
				resource.Attribute{
					Name:        "reverse_dns",
					Description: `(Optional) If set to ` + "`" + `true` + "`" + ` (default), then reverse DNS records are created. If set to ` + "`" + `false` + "`" + `, no reverse DNS records are created.`,
				},
				resource.Attribute{
					Name:        "ssh_keys",
					Description: `(Optional) A list of the names of the SSH Keys that can be used to log into the instance.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A list of strings that should be supplied to the instance as tags. ## Attributes During instance creation, there are several custom attributes that a user may wish to make available to the instance during instance creation. These attributes can be specified via the ` + "`" + `instance_attributes` + "`" + ` field, and must be presented as a string in JSON format. The easiest way to populate this field is with a HEREDOC: ` + "`" + `` + "`" + `` + "`" + `hcl resource "opc_compute_instance" "foo" { name = "test" label = "test" shape = "oc3" imageList = "/oracle/public/oel_6.4_2GB_v1" instance_attributes = <<JSON { "foo": "bar", "baz": 42, "my_obj": { "my_key": false, "another": true } } JSON sshKeys = ["${opc_compute_ssh_key.key1.name}"] } ` + "`" + `` + "`" + `` + "`" + ` This allows the user to have full control over the attributes supplied to an instance during instance creation. There are, as well, some attributes that get populated during instance creation, and the full attributes map can be seen via the exported ` + "`" + `attributes` + "`" + ` attribute.`,
				},
				resource.Attribute{
					Name:        "index",
					Description: `(Required) The numerical index of the network interface. Specified as an integer to allow for use of ` + "`" + `count` + "`" + `, but directly maps to ` + "`" + `ethX` + "`" + `. ie: With ` + "`" + `index` + "`" + ` set to ` + "`" + `0` + "`" + `, the interface ` + "`" + `eth0` + "`" + ` will be created. Can only be ` + "`" + `0-9` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dns",
					Description: `(Optional, IP Network Only) List of DNS A record names for the instance. You can specify up to eight DNS A record names for each interface on an IP network. These names can be queried by instances on any IP network in the same IP network exchange.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional, IP Network Only) IP Address assigned to the interface.`,
				},
				resource.Attribute{
					Name:        "ip_network",
					Description: `(Optional, IP Network Only) The IP Network assigned to the interface.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `(Optional, IP Network Only) The MAC address of the interface.`,
				},
				resource.Attribute{
					Name:        "is_default_gateway",
					Description: `(Optional, IP Network Only) Specify the interface is to be used as the default gateway for all traffic. Only one interface on an instance can be specified as the default gateway. If the instance has an interface on the shared network, that interface is always used as the default gateway.`,
				},
				resource.Attribute{
					Name:        "model",
					Description: `(Optional, Shared Network Only) The model of the NIC card used. Must be set to ` + "`" + `e1000` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name_servers",
					Description: `(Optional) Array of name servers for the interface.`,
				},
				resource.Attribute{
					Name:        "nat",
					Description: `(Optional for IP Networks, Required for the Shared Network) The IP Reservations associated with the interface (IP Network). Indicates whether a temporary or permanent public IP address should be assigned to the instance (Shared Network).`,
				},
				resource.Attribute{
					Name:        "search_domains",
					Description: `(Optional) The search domains that are sent through DHCP as option 119.`,
				},
				resource.Attribute{
					Name:        "sec_lists",
					Description: `(Optional, Shared Network Only) The security lists the interface is added to.`,
				},
				resource.Attribute{
					Name:        "shared_network",
					Description: `(Required) Whether or not the interface is inside the Shared Network or an IP Network.`,
				},
				resource.Attribute{
					Name:        "vnic",
					Description: `(Optional, IP Network Only) The name of the vNIC created for the IP Network.`,
				},
				resource.Attribute{
					Name:        "vnic_sets",
					Description: `(Optional, IP Network Only) The array of vNIC Sets the interface was added to. ## Storage Attachments Each Storage Attachment config manages a single storage attachment that is created _during instance creation_. This means that any storage attachments created during instance creation cannot be detached from the instance. Use the ` + "`" + `resource_storage_attachment` + "`" + ` resource to manage storage attachments for instances if you wish to detach the storage volumes at a later date. The following attributes are supported:`,
				},
				resource.Attribute{
					Name:        "index",
					Description: `(Required) The Index number of the volume attachment. ` + "`" + `1` + "`" + ` is the boot volume for the instance. Values ` + "`" + `1-10` + "`" + ` allowed.`,
				},
				resource.Attribute{
					Name:        "volume",
					Description: `(Required) The name of the storage volume to attach to the instance. In addition to the above attributes, the following attributes are exported for a storage volume`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the storage volume attachment. ## Attributes Reference In addition to the attributes listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `id` + "`" + ` of the instance.`,
				},
				resource.Attribute{
					Name:        "attributes",
					Description: `The full attributes of the instance, as a JSON string.`,
				},
				resource.Attribute{
					Name:        "availability_domain",
					Description: `The availability domain the instance is in.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `The default domain to use for the hostname and for DNS lookups.`,
				},
				resource.Attribute{
					Name:        "entry",
					Description: `Imagelist entry number.`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `SSH server fingerprint presented by the instance.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `The fully qualified domain name of the instance.`,
				},
				resource.Attribute{
					Name:        "image_format",
					Description: `The format of the image.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The IP Address of the instance.`,
				},
				resource.Attribute{
					Name:        "placement_requirements",
					Description: `The array of placement requirements for the instance.`,
				},
				resource.Attribute{
					Name:        "platform",
					Description: `The OS Platform of the instance.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority at which the instance was ran.`,
				},
				resource.Attribute{
					Name:        "quota_reservation",
					Description: `Reference to the QuotaReservation, to be destroyed with the instance.`,
				},
				resource.Attribute{
					Name:        "relationships",
					Description: `The array of relationship specifications to be satisfied on instance placement.`,
				},
				resource.Attribute{
					Name:        "resolvers",
					Description: `Array of resolvers to be used instead of the default resolvers.`,
				},
				resource.Attribute{
					Name:        "site",
					Description: `The site the instance is running on.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `The launch time of the instance.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The instance's state.`,
				},
				resource.Attribute{
					Name:        "vcable_id",
					Description: `vCable ID for the instance.`,
				},
				resource.Attribute{
					Name:        "virtio",
					Description: `Boolean that determines if the instance is a virtio device.`,
				},
				resource.Attribute{
					Name:        "vnc_address",
					Description: `The VNC address and port of the instance. ## Import Instances can be imported using the Instance's combined ` + "`" + `Name` + "`" + ` and ` + "`" + `ID` + "`" + ` with a ` + "`" + `/` + "`" + ` character separating them. If viewing an instance in the Oracle Web Console, the instance's ` + "`" + `name` + "`" + ` and ` + "`" + `id` + "`" + ` are the last two fields in the instances fully qualified ` + "`" + `Name` + "`" + ` For example, in the Web Console an instance's fully qualified name is: ` + "`" + `` + "`" + `` + "`" + ` /Compute-<identify>/<user>@<account>/<instance_name>/<instance_id> ` + "`" + `` + "`" + `` + "`" + ` The instance can be imported as such: ` + "`" + `` + "`" + `` + "`" + `shell $ terraform import opc_compute_instance.instance1 instance_name/instance_id ` + "`" + `` + "`" + `` + "`" + ` <a id="timeouts"></a> ## Timeouts ` + "`" + `opc_compute_instance` + "`" + ` provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `20 minutes` + "`" + `) Used for Creating Instances. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `20 minutes` + "`" + `) Used for updating Instances. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `20 minutes` + "`" + `) Used for Deleting Instances.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `id` + "`" + ` of the instance.`,
				},
				resource.Attribute{
					Name:        "attributes",
					Description: `The full attributes of the instance, as a JSON string.`,
				},
				resource.Attribute{
					Name:        "availability_domain",
					Description: `The availability domain the instance is in.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `The default domain to use for the hostname and for DNS lookups.`,
				},
				resource.Attribute{
					Name:        "entry",
					Description: `Imagelist entry number.`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `SSH server fingerprint presented by the instance.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `The fully qualified domain name of the instance.`,
				},
				resource.Attribute{
					Name:        "image_format",
					Description: `The format of the image.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The IP Address of the instance.`,
				},
				resource.Attribute{
					Name:        "placement_requirements",
					Description: `The array of placement requirements for the instance.`,
				},
				resource.Attribute{
					Name:        "platform",
					Description: `The OS Platform of the instance.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority at which the instance was ran.`,
				},
				resource.Attribute{
					Name:        "quota_reservation",
					Description: `Reference to the QuotaReservation, to be destroyed with the instance.`,
				},
				resource.Attribute{
					Name:        "relationships",
					Description: `The array of relationship specifications to be satisfied on instance placement.`,
				},
				resource.Attribute{
					Name:        "resolvers",
					Description: `Array of resolvers to be used instead of the default resolvers.`,
				},
				resource.Attribute{
					Name:        "site",
					Description: `The site the instance is running on.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `The launch time of the instance.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The instance's state.`,
				},
				resource.Attribute{
					Name:        "vcable_id",
					Description: `vCable ID for the instance.`,
				},
				resource.Attribute{
					Name:        "virtio",
					Description: `Boolean that determines if the instance is a virtio device.`,
				},
				resource.Attribute{
					Name:        "vnc_address",
					Description: `The VNC address and port of the instance. ## Import Instances can be imported using the Instance's combined ` + "`" + `Name` + "`" + ` and ` + "`" + `ID` + "`" + ` with a ` + "`" + `/` + "`" + ` character separating them. If viewing an instance in the Oracle Web Console, the instance's ` + "`" + `name` + "`" + ` and ` + "`" + `id` + "`" + ` are the last two fields in the instances fully qualified ` + "`" + `Name` + "`" + ` For example, in the Web Console an instance's fully qualified name is: ` + "`" + `` + "`" + `` + "`" + ` /Compute-<identify>/<user>@<account>/<instance_name>/<instance_id> ` + "`" + `` + "`" + `` + "`" + ` The instance can be imported as such: ` + "`" + `` + "`" + `` + "`" + `shell $ terraform import opc_compute_instance.instance1 instance_name/instance_id ` + "`" + `` + "`" + `` + "`" + ` <a id="timeouts"></a> ## Timeouts ` + "`" + `opc_compute_instance` + "`" + ` provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `20 minutes` + "`" + `) Used for Creating Instances. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `20 minutes` + "`" + `) Used for updating Instances. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `20 minutes` + "`" + `) Used for Deleting Instances.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opc_compute_ip_address_association",
			Category:         "Compute Classic Resources",
			ShortDescription: `Creates and manages an IP address association in an Oracle Cloud Infrastructure Compute Classic identity domain, for an IP Network.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"classic",
				"ip",
				"address",
				"association",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the ip address association.`,
				},
				resource.Attribute{
					Name:        "ip_address_reservation",
					Description: `(Optional) The name of the NAT IP address reservation.`,
				},
				resource.Attribute{
					Name:        "vnic",
					Description: `(Optional) The name of the virtual NIC associated with this NAT IP reservation.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the ip address association.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) List of tags that may be applied to the ip address association. In addition to the above, the following variables are exported:`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `(Computed) The Uniform Resource Identifier of the ip address association. ## Import IP Address Associations can be imported using the ` + "`" + `resource name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ terraform import opc_compute_ip_address_association.default example ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opc_compute_ip_address_prefix_set",
			Category:         "Compute Classic Resources",
			ShortDescription: `Creates and manages an IP address prefix set in an Oracle Cloud Infrastructure Compute Classic identity domain.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"classic",
				"ip",
				"address",
				"prefix",
				"set",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the ip address prefix set.`,
				},
				resource.Attribute{
					Name:        "prefixes",
					Description: `(Optional) List of CIDR IPv4 prefixes assigned in the virtual network.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the ip address prefix set.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) List of tags that may be applied to the ip address prefix set. In addition to the above, the following variables are exported:`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `(Computed) The Uniform Resource Identifier of the ip address prefix set. ## Import IP Address Prefix Set can be imported using the ` + "`" + `resource name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ terraform import opc_compute_ip_address_prefix_set.default example ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opc_compute_ip_address_reservation",
			Category:         "Compute Classic Resources",
			ShortDescription: `Creates and manages an IP address reservation in an Oracle Cloud Infrastructure Compute Classic identity domain for an IP Network.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"classic",
				"ip",
				"address",
				"reservation",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the ip address reservation.`,
				},
				resource.Attribute{
					Name:        "ip_address_pool",
					Description: `(Required) The IP address pool from which you want to reserve an IP address. Typically one of either ` + "`" + `public-ippool` + "`" + ` or ` + "`" + `cloud-ippool` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the ip address reservation.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) List of tags that may be applied to the IP address reservation. In addition to the above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `Reserved NAT IPv4 address from the IP address pool.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `The Uniform Resource Identifier of the ip address reservation ## Import IP Address Reservations can be imported using the ` + "`" + `resource name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ terraform import opc_compute_ip_address_reservation.default example ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opc_compute_ip_association",
			Category:         "Compute Classic Resources",
			ShortDescription: `Creates and manages an IP association in an Oracle Cloud Infrastructure Compute Classic identity domain for the Shared Network.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"classic",
				"ip",
				"association",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "vcable",
					Description: `(Required) The vcable of the instance to associate the IP address with.`,
				},
				resource.Attribute{
					Name:        "parent_pool",
					Description: `(Required) The pool from which to take an IP address. To associate a specific reserved IP address, use the prefix ` + "`" + `ipreservation:` + "`" + ` followed by the name of the IP reservation. To allocate an IP address from a pool, use the prefix ` + "`" + `ippool:` + "`" + `, e.g. ` + "`" + `ippool:/oracle/public/ippool` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opc_compute_ip_network",
			Category:         "Compute Classic Resources",
			ShortDescription: `Creates and manages an IP Network in an Oracle Cloud Infrastructure Compute Classic identity domain.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"classic",
				"ip",
				"network",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the IP Network. Changing this name forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "ip_address_prefix",
					Description: `(Required) The IPv4 address prefix, in CIDR format.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the IP Network.`,
				},
				resource.Attribute{
					Name:        "ip_network_exchange",
					Description: `(Optional) Specify the IP Network exchange to which the IP Network belongs to.`,
				},
				resource.Attribute{
					Name:        "public_napt_enabled",
					Description: `(Optional) If true, enable public internet access using NAPT for VNICs without any public IP Reservation. Defaults to ` + "`" + `false` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the IP Network`,
				},
				resource.Attribute{
					Name:        "ip_address_prefix",
					Description: `The IPv4 address prefix, in CIDR format.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the IP Network.`,
				},
				resource.Attribute{
					Name:        "ip_network_exchange",
					Description: `The IP Network Exchange for the IP Network`,
				},
				resource.Attribute{
					Name:        "public_napt_enabled",
					Description: `Whether public internet access using NAPT for VNICs without any public IP Reservation or not.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `Uniform Resource Identifier for the IP Network ## Import IP Networks can be imported using the ` + "`" + `resource name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ terraform import opc_compute_ip_network.default example ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the IP Network`,
				},
				resource.Attribute{
					Name:        "ip_address_prefix",
					Description: `The IPv4 address prefix, in CIDR format.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the IP Network.`,
				},
				resource.Attribute{
					Name:        "ip_network_exchange",
					Description: `The IP Network Exchange for the IP Network`,
				},
				resource.Attribute{
					Name:        "public_napt_enabled",
					Description: `Whether public internet access using NAPT for VNICs without any public IP Reservation or not.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `Uniform Resource Identifier for the IP Network ## Import IP Networks can be imported using the ` + "`" + `resource name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ terraform import opc_compute_ip_network.default example ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opc_compute_ip_network_exchange",
			Category:         "Compute Classic Resources",
			ShortDescription: `Creates and manages an IP network exchange in an Oracle Cloud Infrastructure Compute Classic identity domain.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"classic",
				"ip",
				"network",
				"exchange",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the ip network exchange.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the ip network exchange.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) List of tags that may be applied to the IP network exchange. ## Import IP Network Exchange's can be imported using the ` + "`" + `resource name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ terraform import opc_compute_ip_network_exchange.exchange1 example ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opc_compute_ip_reservation",
			Category:         "Compute Classic Resources",
			ShortDescription: `Creates and manages an IP reservation in an Oracle Cloud Infrastructure Compute Classic identity domain for the Shared Network.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"classic",
				"ip",
				"reservation",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "permanent",
					Description: `(Required) Whether the IP address remains reserved even when it is no longer associated with an instance (if true), or may be returned to the pool and replaced with a different IP address when an instance is restarted, or deleted and recreated (if false).`,
				},
				resource.Attribute{
					Name:        "parent_pool",
					Description: `(Optional) The pool from which to allocate the IP address. Defaults to ` + "`" + `/oracle/public/ippool` + "`" + `, and is currently the only acceptable input.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the IP Reservation. Will be generated if unspecified.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) List of tags that may be applied to the IP reservation. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `The Public IP address.`,
				},
				resource.Attribute{
					Name:        "used",
					Description: `indicates that the IP reservation is associated with an instance. ## Import IP Reservations can be imported using the ` + "`" + `resource name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ terraform import opc_compute_ip_reservations.reservation1 example ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "ip",
					Description: `The Public IP address.`,
				},
				resource.Attribute{
					Name:        "used",
					Description: `indicates that the IP reservation is associated with an instance. ## Import IP Reservations can be imported using the ` + "`" + `resource name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ terraform import opc_compute_ip_reservations.reservation1 example ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opc_compute_machine_image",
			Category:         "Compute Classic Resources",
			ShortDescription: `Creates and manages a Machine Image in a Oracle Cloud Infrastructure Compute Classic identity domain.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"classic",
				"machine",
				"image",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "account",
					Description: `(Required) The two part name of the compute object storage account in the format ` + "`" + `/Compute-{identity_domain}/cloud_storage` + "`" + ``,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Machine Image.`,
				},
				resource.Attribute{
					Name:        "file",
					Description: `(Required) The name of the Machine Image .tar.gz file in the ` + "`" + `compute_images` + "`" + ` storage container.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the Machine Image.`,
				},
				resource.Attribute{
					Name:        "attributes",
					Description: `(Optional) An optional JSON object of arbitrary attributes to be made available to the instance. These are user-defined tags. After defining attributes, you can view them from within an instance at http://192.0.0.192/ In addition to the above, the following values are exported:`,
				},
				resource.Attribute{
					Name:        "error_reason",
					Description: `Description of the state of the machine image if there is an error.`,
				},
				resource.Attribute{
					Name:        "hypervisor",
					Description: `Dictionary of hypervisor-specific attributes.`,
				},
				resource.Attribute{
					Name:        "image_format",
					Description: `The format of the image.`,
				},
				resource.Attribute{
					Name:        "platform",
					Description: `The OS platform of the image.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the uploaded machine image.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `The Uniform Resource Identifier for the Machine Image. ## Import Machine Images can be imported using the ` + "`" + `resource name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ terraform import opc_compute_machine_image.machine_image1 example ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opc_compute_orchestrated_instance",
			Category:         "Compute Classic Resources",
			ShortDescription: `Creates and manages an Orchestration containing a number of instances in an Oracle Cloud Infrastructure Compute Classic identity domain.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"classic",
				"orchestrated",
				"instance",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the orchestration.`,
				},
				resource.Attribute{
					Name:        "desired_state",
					Description: `(Required) The desired state of the orchestration. Permitted values are: - ` + "`" + `active` + "`" + `: all resource (instances) declared in the orchestration are created - ` + "`" + `suspend` + "`" + `: all resources (instances) declared in the orchestration are removed unless the instances has ` + "`" + `persistent = true` + "`" + ` - ` + "`" + `inactive` + "`" + `: all resources (instances) declared in the orchestration are removed including the instances that have ` + "`" + `persistent = true` + "`" + ``,
				},
				resource.Attribute{
					Name:        "instance",
					Description: `(Required) The information pertaining to creating an instance through the orchestration API.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the orchestration. ## Instance Instance supports the arguments found in [opc_compute_instance](https://www.terraform.io/docs/providers/opc/r/opc_compute_instance.html) and the following:`,
				},
				resource.Attribute{
					Name:        "persistent",
					Description: `(Optional) Determines whether the instance will persist when the orchestration is suspended. Defaults to false. In addition to the above, the following values are exported:`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `The Uniform Resource Identifier for the Orchestration`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) The version of the orchestration.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opc_compute_route",
			Category:         "Compute Classic Resources",
			ShortDescription: `Creates and manages a Route resource for an IP Network in an Oracle Cloud Infrastructure Compute Classic identity domain.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"classic",
				"route",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the route.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the route.`,
				},
				resource.Attribute{
					Name:        "admin_distance",
					Description: `(Optional) The route's administrative distance. Defaults to ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ip_address_prefix",
					Description: `(Required) The IPv4 address prefix, in CIDR format, of the external network from which to route traffic.`,
				},
				resource.Attribute{
					Name:        "next_hop_vnic_set",
					Description: `(Required) Name of the virtual NIC set to route matching packets to. Routed flows are load-balanced among all the virtual NICs in the virtual NIC set. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the route.`,
				},
				resource.Attribute{
					Name:        "admin_distance",
					Description: `The route's administrative distance. Defaults to ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ip_address_prefix",
					Description: `The IPv4 address prefix, in CIDR format, of the external network from which to route traffic.`,
				},
				resource.Attribute{
					Name:        "next_hop_vnic_set",
					Description: `Name of the virtual NIC set to route matching packets to. Routed flows are load-balanced among all the virtual NICs in the virtual NIC set. ## Import Route's can be imported using the ` + "`" + `resource name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ terraform import opc_compute_route.route1 example ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the route.`,
				},
				resource.Attribute{
					Name:        "admin_distance",
					Description: `The route's administrative distance. Defaults to ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ip_address_prefix",
					Description: `The IPv4 address prefix, in CIDR format, of the external network from which to route traffic.`,
				},
				resource.Attribute{
					Name:        "next_hop_vnic_set",
					Description: `Name of the virtual NIC set to route matching packets to. Routed flows are load-balanced among all the virtual NICs in the virtual NIC set. ## Import Route's can be imported using the ` + "`" + `resource name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ terraform import opc_compute_route.route1 example ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opc_compute_sec_rule",
			Category:         "Compute Classic Resources",
			ShortDescription: `Creates and manages a sec rule in an Oracle Cloud Infrastructure Compute Classic identity domain.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"classic",
				"sec",
				"rule",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The unique (within the identity domain) name of the security rule.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description for this security rule.`,
				},
				resource.Attribute{
					Name:        "source_list",
					Description: `(Required) The source security list (prefixed with ` + "`" + `seclist:` + "`" + `), or security IP list (prefixed with ` + "`" + `seciplist:` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "destination_list",
					Description: `(Required) The destination security list (prefixed with ` + "`" + `seclist:` + "`" + `), or security IP list (prefixed with ` + "`" + `seciplist:` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "application",
					Description: `(Required) The name of the application to which the rule applies.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) Whether to ` + "`" + `permit` + "`" + `, ` + "`" + `refuse` + "`" + ` or ` + "`" + `deny` + "`" + ` packets to which this rule applies. This will ordinarily be ` + "`" + `permit` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional) Whether to disable this security rule. This is useful if you want to temporarily disable a rule without removing it outright from your Terraform resource definition. Defaults to ` + "`" + `false` + "`" + `. In addition to the above, the following values are exported:`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `The Uniform Resource Identifier of the sec rule. ## Import Sec Rule's can be imported using the ` + "`" + `resource name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ terraform import opc_compute_sec_rule.rule1 example ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opc_compute_security_application",
			Category:         "Compute Classic Resources",
			ShortDescription: `Creates and manages a security application in an Oracle Cloud Infrastructure Compute Classic identity domain.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"classic",
				"security",
				"application",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The unique (within the identity domain) name of the application`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The protocol to enable for this application. Must be one of ` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `ah` + "`" + `, ` + "`" + `esp` + "`" + `, ` + "`" + `icmp` + "`" + `, ` + "`" + `icmpv6` + "`" + `, ` + "`" + `igmp` + "`" + `, ` + "`" + `ipip` + "`" + `, ` + "`" + `gre` + "`" + `, ` + "`" + `mplsip` + "`" + `, ` + "`" + `ospf` + "`" + `, ` + "`" + `pim` + "`" + `, ` + "`" + `rdp` + "`" + `, ` + "`" + `sctp` + "`" + ` or ` + "`" + `all` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dport",
					Description: `(Required) The port, or range of ports, to enable for this application, e.g ` + "`" + `8080` + "`" + `, ` + "`" + `6000-7000` + "`" + `. This must be set if the ` + "`" + `protocol` + "`" + ` is set to ` + "`" + `tcp` + "`" + ` or ` + "`" + `udp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "icmptype",
					Description: `(Optional) The ICMP type to enable for this application, if the ` + "`" + `protocol` + "`" + ` is ` + "`" + `icmp` + "`" + `. Must be one of ` + "`" + `echo` + "`" + `, ` + "`" + `reply` + "`" + `, ` + "`" + `ttl` + "`" + `, ` + "`" + `traceroute` + "`" + `, ` + "`" + `unreachable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "icmpcode",
					Description: `(Optional) The ICMP code to enable for this application, if the ` + "`" + `protocol` + "`" + ` is ` + "`" + `icmp` + "`" + `. Must be one of ` + "`" + `admin` + "`" + `, ` + "`" + `df` + "`" + `, ` + "`" + `host` + "`" + `, ` + "`" + `network` + "`" + `, ` + "`" + `port` + "`" + ` or ` + "`" + `protocol` + "`" + `. ## Import Security Application's can be imported using the ` + "`" + `resource name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ terraform import opc_compute_security_application.application1 example ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opc_compute_security_association",
			Category:         "Compute Classic Resources",
			ShortDescription: `Creates and manages a security association in an Oracle Cloud Infrastructure Compute Classic identity domain.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"classic",
				"security",
				"association",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The Name for the Security Association. If not specified, one is created automatically. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "vcable",
					Description: `(Required) The ` + "`" + `vcable` + "`" + ` of the instance to associate to the security list.`,
				},
				resource.Attribute{
					Name:        "seclist",
					Description: `(Required) The name of the security list to associate the instance to. ## Import Security Association's can be imported using the ` + "`" + `resource name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ terraform import opc_compute_security_association.association1 example ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opc_compute_security_ip_list",
			Category:         "Compute Classic Resources",
			ShortDescription: `Creates and manages a security IP list in an Oracle Cloud Infrastructure Compute Classic identity domain.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"classic",
				"security",
				"ip",
				"list",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The unique (within the identity domain) name of the security IP list.`,
				},
				resource.Attribute{
					Name:        "ip_entries",
					Description: `(Required) The IP addresses to include in the list.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the security ip list. ## Import IP List's can be imported using the ` + "`" + `resource name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ terraform import opc_compute_ip_list.list1 example ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opc_compute_security_list",
			Category:         "Compute Classic Resources",
			ShortDescription: `Creates and manages a security list in an Oracle Cloud Infrastructure Compute Classic identity domain.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"classic",
				"security",
				"list",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The unique (within the identity domain) name of the security list.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Required) The policy to apply to instances associated with this list. Must be one of ` + "`" + `permit` + "`" + `, ` + "`" + `reject` + "`" + ` (packets are dropped but a reply is sent) and ` + "`" + `deny` + "`" + ` (packets are dropped and no reply is sent).`,
				},
				resource.Attribute{
					Name:        "output_cidr_policy",
					Description: `(Required) The policy for outbound traffic from the security list. Must be one of ` + "`" + `permit` + "`" + `, ` + "`" + `reject` + "`" + ` (packets are dropped but a reply is sent) and ` + "`" + `deny` + "`" + ` (packets are dropped and no reply is sent). ## Import Security List's can be imported using the ` + "`" + `resource name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ terraform import opc_compute_security_list.list1 example ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opc_compute_security_protocol",
			Category:         "Compute Classic Resources",
			ShortDescription: `Creates and manages an security protocol in an Oracle Cloud Infrastructure Compute Classic identity domain.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"classic",
				"security",
				"protocol",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the security protocol.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the security protocol.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) List of tags that may be applied to the security protocol. In addition to the above, the following values are exported:`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `The Uniform Resource Identifier for the Security Protocol ## Import ACL's can be imported using the ` + "`" + `resource name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ terraform import opc_compute_security_protocol.default example ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opc_compute_security_rule",
			Category:         "Compute Classic Resources",
			ShortDescription: `Creates and manages a security rule in an Oracle Cloud Infrastructure Compute Classic identity domain.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"classic",
				"security",
				"rule",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the security rule.`,
				},
				resource.Attribute{
					Name:        "flow_direction",
					Description: `(Required) Specify the direction of flow of traffic, which is relative to the instances, for this security rule. Allowed values are ingress or egress.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional) Whether to disable this security rule. This is useful if you want to temporarily disable a rule without removing it outright from your Terraform resource definition. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `(Optional) Name of the ACL that contains this security rule.`,
				},
				resource.Attribute{
					Name:        "dst_ip_address_prefixes",
					Description: `(Optional) List of IP address prefix set names to match the packet's destination IP address.`,
				},
				resource.Attribute{
					Name:        "src_ip_address_prefixes",
					Description: `(Optional) List of names of IP address prefix set to match the packet's source IP address.`,
				},
				resource.Attribute{
					Name:        "dst_vnic_set",
					Description: `(Optional) Name of virtual NIC set containing the packet's destination virtual NIC.`,
				},
				resource.Attribute{
					Name:        "src_vnic_set",
					Description: `(Optional) Name of virtual NIC set containing the packet's source virtual NIC.`,
				},
				resource.Attribute{
					Name:        "security_protocols",
					Description: `(Optional) List of security protocol object names to match the packet's protocol and port.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the security rule.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) List of tags that may be applied to the security rule. ## Attributes Reference In addition to the above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `The Uniform Resource Identifier of the security rule. ## Import Security Rule's can be imported using the ` + "`" + `resource name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ terraform import opc_compute_security_rule.rule1 example ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "uri",
					Description: `The Uniform Resource Identifier of the security rule. ## Import Security Rule's can be imported using the ` + "`" + `resource name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ terraform import opc_compute_security_rule.rule1 example ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opc_compute_ssh_key",
			Category:         "Compute Classic Resources",
			ShortDescription: `Creates and manages an SSH key in an Oracle Cloud Infrastructure Compute Classic identity domain.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"classic",
				"ssh",
				"key",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The unique (within this identity domain) name of the SSH key.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The SSH key itself`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Whether or not the key is enabled. This is useful if you want to temporarily disable an SSH key, without removing it entirely from your Terraform resource definition. Defaults to ` + "`" + `true` + "`" + ` ## Import SSH Key's can be imported using the ` + "`" + `resource name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ terraform import opc_compute_ssh_key.key1 example ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opc_compute_storage_volume",
			Category:         "Compute Classic Resources",
			ShortDescription: `Creates and manages a storage volume in an Oracle Cloud Infrastructure Compute Classic identity domain.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"classic",
				"storage",
				"volume",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "storage_type",
					Description: `(Optional) - The Type of Storage to provision. Defaults to ` + "`" + `/oracle/public/storage/default` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "bootable",
					Description: `(Optional) Is the Volume Bootable? Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "image_list",
					Description: `(Optional) Defines an image list.`,
				},
				resource.Attribute{
					Name:        "image_list_entry",
					Description: `(Optional) Defines an image list entry.`,
				},
				resource.Attribute{
					Name:        "snapshot",
					Description: `(Optional) The name of the parent snapshot from which the storage volume is restored or cloned. See [Snapshots](#snapshots), below for more information.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) The Id of the parent snapshot from which the storage volume is restored or cloned. See [Snapshots](#snapshots), below for more information.`,
				},
				resource.Attribute{
					Name:        "snapshot_account",
					Description: `(Optional) The Account of the parent snapshot from which the storage volume is restored. See [Snapshots](#snapshots), below for more information.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Comma-separated strings that tag the storage volume. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "hypervisor",
					Description: `The hypervisor that this volume is compatible with.`,
				},
				resource.Attribute{
					Name:        "machine_image",
					Description: `Name of the Machine Image - available if the volume is a bootable storage volume.`,
				},
				resource.Attribute{
					Name:        "managed",
					Description: `Is this a Managed Volume?`,
				},
				resource.Attribute{
					Name:        "platform",
					Description: `The OS platform this volume is compatible with.`,
				},
				resource.Attribute{
					Name:        "readonly",
					Description: `Can this Volume be attached as readonly?`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current state of the storage volume.`,
				},
				resource.Attribute{
					Name:        "storage_pool",
					Description: `The storage pool from which this volume is allocated.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `Unique Resource Identifier of the Storage Volume. ## Import Storage Volume's can be imported using the ` + "`" + `resource name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ terraform import opc_compute_storage_volume.volume1 example ` + "`" + `` + "`" + `` + "`" + ` <a id="snapshots"></a> ## Snapshots Restoring a storage volume from a snapshot can happen via the use of the ` + "`" + `snapshot` + "`" + `, ` + "`" + `snapshot_id` + "`" + `, and ` + "`" + `snapshot_account` + "`" + ` attributes.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "hypervisor",
					Description: `The hypervisor that this volume is compatible with.`,
				},
				resource.Attribute{
					Name:        "machine_image",
					Description: `Name of the Machine Image - available if the volume is a bootable storage volume.`,
				},
				resource.Attribute{
					Name:        "managed",
					Description: `Is this a Managed Volume?`,
				},
				resource.Attribute{
					Name:        "platform",
					Description: `The OS platform this volume is compatible with.`,
				},
				resource.Attribute{
					Name:        "readonly",
					Description: `Can this Volume be attached as readonly?`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current state of the storage volume.`,
				},
				resource.Attribute{
					Name:        "storage_pool",
					Description: `The storage pool from which this volume is allocated.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `Unique Resource Identifier of the Storage Volume. ## Import Storage Volume's can be imported using the ` + "`" + `resource name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ terraform import opc_compute_storage_volume.volume1 example ` + "`" + `` + "`" + `` + "`" + ` <a id="snapshots"></a> ## Snapshots Restoring a storage volume from a snapshot can happen via the use of the ` + "`" + `snapshot` + "`" + `, ` + "`" + `snapshot_id` + "`" + `, and ` + "`" + `snapshot_account` + "`" + ` attributes.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opc_compute_storage_volume_attachment",
			Category:         "Compute Classic Resources",
			ShortDescription: `Creates and manages a storage volume attachment in an Oracle Cloud Infrastructure Compute Classic identity domain.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"classic",
				"storage",
				"volume",
				"attachment",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "instance",
					Description: `(Required) The name of the instance the volume will be attached to.`,
				},
				resource.Attribute{
					Name:        "storage_volume",
					Description: `(Required) The name of the storage volume that will be attached to the instance`,
				},
				resource.Attribute{
					Name:        "index",
					Description: `(Required) The index on the instance that the storage volume will be attached to.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opc_compute_storage_volume_snapshot",
			Category:         "Compute Classic Resources",
			ShortDescription: `Creates and manages a storage volume snapshot in an Oracle Cloud Infrastructure Compute Classic identity domain.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"classic",
				"storage",
				"volume",
				"snapshot",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Comma-separated strings that tag the storage volume. ## Attributes Reference In addition to the attributes above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "account",
					Description: `Account to use for snapshots.`,
				},
				resource.Attribute{
					Name:        "machine_image_name",
					Description: `The name of the machine image that's used in the boot volume from which this snapshot is taken.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the snapshot in GB.`,
				},
				resource.Attribute{
					Name:        "property",
					Description: `Where the snapshot is stored, whether collocated, or in the Oracle Storage Cloud Service instance.`,
				},
				resource.Attribute{
					Name:        "platform",
					Description: `The OS platform this snapshot is compatible with`,
				},
				resource.Attribute{
					Name:        "snapshot_timestamp",
					Description: `Timestamp of the storage snapshot, generated by storage server. The snapshot will contain data written to the original volume before this time.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The Oracle ID of the snapshot.`,
				},
				resource.Attribute{
					Name:        "start_timestamp",
					Description: `Timestamp when the snapshot was started.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the snapshot.`,
				},
				resource.Attribute{
					Name:        "status_detail",
					Description: `Details about the latest state of the storage volume snapshot.`,
				},
				resource.Attribute{
					Name:        "status_timestamp",
					Description: `Indicates the time that the current view of the storage volume snapshot was generated.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `Uniform Resource Identifier ## Import Storage Volume Snapshot's can be imported using the ` + "`" + `resource name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ terraform import opc_compute_storage_volume_snapshot.volume1 example ` + "`" + `` + "`" + `` + "`" + ` <a id="timeouts"></a> ## Timeouts ` + "`" + `opc_compute_storage_volume_snapshot` + "`" + ` provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `30 minutes` + "`" + `) Used for Creating Storage Volume Snapshots. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `30 minutes` + "`" + `) Used for Deleting Storage Volume Snapshots.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "account",
					Description: `Account to use for snapshots.`,
				},
				resource.Attribute{
					Name:        "machine_image_name",
					Description: `The name of the machine image that's used in the boot volume from which this snapshot is taken.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the snapshot in GB.`,
				},
				resource.Attribute{
					Name:        "property",
					Description: `Where the snapshot is stored, whether collocated, or in the Oracle Storage Cloud Service instance.`,
				},
				resource.Attribute{
					Name:        "platform",
					Description: `The OS platform this snapshot is compatible with`,
				},
				resource.Attribute{
					Name:        "snapshot_timestamp",
					Description: `Timestamp of the storage snapshot, generated by storage server. The snapshot will contain data written to the original volume before this time.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The Oracle ID of the snapshot.`,
				},
				resource.Attribute{
					Name:        "start_timestamp",
					Description: `Timestamp when the snapshot was started.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the snapshot.`,
				},
				resource.Attribute{
					Name:        "status_detail",
					Description: `Details about the latest state of the storage volume snapshot.`,
				},
				resource.Attribute{
					Name:        "status_timestamp",
					Description: `Indicates the time that the current view of the storage volume snapshot was generated.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `Uniform Resource Identifier ## Import Storage Volume Snapshot's can be imported using the ` + "`" + `resource name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ terraform import opc_compute_storage_volume_snapshot.volume1 example ` + "`" + `` + "`" + `` + "`" + ` <a id="timeouts"></a> ## Timeouts ` + "`" + `opc_compute_storage_volume_snapshot` + "`" + ` provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `30 minutes` + "`" + `) Used for Creating Storage Volume Snapshots. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `30 minutes` + "`" + `) Used for Deleting Storage Volume Snapshots.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opc_compute_vnic_set",
			Category:         "Compute Classic Resources",
			ShortDescription: `Creates and manages a virtual NIC set in an Oracle Cloud Infrastructure Compute Classic identity domain`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"classic",
				"vnic",
				"set",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The unique (within this identity domain) name of the virtual nic set.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the virtual nic set.`,
				},
				resource.Attribute{
					Name:        "applied_acls",
					Description: `(Optional) A list of the ACLs to apply to the virtual nics in the set.`,
				},
				resource.Attribute{
					Name:        "virtual_nics",
					Description: `(Optional) List of virtual NICs associated with this virtual NIC set.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A list of tags to apply to the storage volume. ## Import VNIC Set's can be imported using the ` + "`" + `resource name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ terraform import opc_compute_vnic_set.set1 example ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opc_compute_vpn_endpoint_v2",
			Category:         "Compute Classic Resources",
			ShortDescription: `Creates and manages a VPN Endpoint V2 in an Oracle Cloud Infrastructure Compute Classic identity domain.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"classic",
				"vpn",
				"endpoint",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the VPN Endpoint V2.`,
				},
				resource.Attribute{
					Name:        "customer_vpn_gateway",
					Description: `(Required) The ip address of the VPN gateway in your data center through which you want to connect to the Oracle Cloud VPN gateway.`,
				},
				resource.Attribute{
					Name:        "ip_network",
					Description: `(Required) The name of the IP network on which the cloud gateway is created by VPNaaS.`,
				},
				resource.Attribute{
					Name:        "pre_shared_key",
					Description: `(Required) The pre-shared VPN key`,
				},
				resource.Attribute{
					Name:        "reachable_routes",
					Description: `(Required) A list of routes (CIDR prefixes) that are reachable through this VPN tunnel.`,
				},
				resource.Attribute{
					Name:        "vnic_sets",
					Description: `(Required) A list of vnic sets that traffics is allowed to and from.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the VPN Endpoint V2.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Enables or disables the VPN Endpoint V2. Set to true by default.`,
				},
				resource.Attribute{
					Name:        "ike_identifier",
					Description: `(Optional) The Internet Key Exchange (IKE) ID. If you don't specify a value, the default value is the public IP address of the cloud gateway.`,
				},
				resource.Attribute{
					Name:        "require_perfect_forward_secrecy",
					Description: `(Optional) Boolean specifying whether Perfect Forward Secrecy is enabled. Set to true by default.`,
				},
				resource.Attribute{
					Name:        "phase_one_settings",
					Description: `(Optional) Settings for the phase one protocol (IKE). Phase One Settings are detailed below.`,
				},
				resource.Attribute{
					Name:        "phase_two_settings",
					Description: `(Optional) Settings for the phase two protocol (IPSEC). Phase Two Settings are detailed below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) List of tags that may be applied to the VPN Endpoint V2. Phase One Settings support the following:`,
				},
				resource.Attribute{
					Name:        "encryption",
					Description: `(Required) IKE Encryption. ` + "`" + `aes128` + "`" + `, ` + "`" + `aes192` + "`" + ` or ` + "`" + `aes256` + "`" + ``,
				},
				resource.Attribute{
					Name:        "hash",
					Description: `(Required) IKE Hash. ` + "`" + `sha1` + "`" + `, ` + "`" + `sha2_256` + "`" + ` or ` + "`" + `md5` + "`" + ``,
				},
				resource.Attribute{
					Name:        "dh_group",
					Description: `(Required) Diffie-Hellman group for both IKE and ESP. ` + "`" + `group2` + "`" + `, ` + "`" + `group5` + "`" + `, ` + "`" + `group14` + "`" + `, ` + "`" + `group22` + "`" + `, ` + "`" + `group23` + "`" + `, or ` + "`" + `group24` + "`" + ``,
				},
				resource.Attribute{
					Name:        "lifetime",
					Description: `(Optional) IKE Lifetime in seconds. Phase Two Settings support the following:`,
				},
				resource.Attribute{
					Name:        "encryption",
					Description: `(Required) ESP Encryption. ` + "`" + `aes128` + "`" + `, ` + "`" + `aes192` + "`" + ` or ` + "`" + `aes256` + "`" + ``,
				},
				resource.Attribute{
					Name:        "hash",
					Description: `(Required) ESP Hash. ` + "`" + `sha1` + "`" + `, ` + "`" + `sha2_256` + "`" + ` or ` + "`" + `md5` + "`" + ``,
				},
				resource.Attribute{
					Name:        "lifetime",
					Description: `(Optional) IPSEC Lifetime in seconds. In addition to the above, the following values are exported:`,
				},
				resource.Attribute{
					Name:        "local_gateway_ip_address",
					Description: `Public IP Address of the Local Gateway.`,
				},
				resource.Attribute{
					Name:        "local_gateway_private_ip_address",
					Description: `Private IP Address of the Local Gateway.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `The Uniform Resource Identifier for the VPN Endpoint V2. ## Import VPN Endpoint V2's can be imported using the ` + "`" + `resource name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ terraform import opc_compute_vpn_endpoint_v2.vpnaas1 /Compute-mydomain/user/example ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opc_lbaas_certificate",
			Category:         "Load Balancer Classic Resources",
			ShortDescription: `Creates an Oracle Load Balancer Classic SSL Certificate.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
				"classic",
				"lbaas",
				"certificate",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Certificate.`,
				},
				resource.Attribute{
					Name:        "certificate_body",
					Description: `(Required) The Certificate data in PEM format.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Sets the Certificate Type. ` + "`" + `TRUSTED` + "`" + ` or ` + "`" + `SERVER` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "certificate_chain",
					Description: `(Optional) PEM encoded bodies of all certificates in the chain up to and including the CA certificate. This is not need when the certificate is self signed.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Optional) The private key data in PEM format. Only required for Server Certificates ## Additional Attributes In addition to the above, the following values are exported:`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The State of the Digital Certificate resource.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `The Uniform Resource Identifier for the Certificate resource.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opc_lbaas_listener",
			Category:         "Load Balancer Classic Resources",
			ShortDescription: `Creates an Oracle Load Balancer Classic Listener.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
				"classic",
				"lbaas",
				"listener",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Listener.`,
				},
				resource.Attribute{
					Name:        "load_balancer",
					Description: `(Required) The parent Load Balancer the Listener.`,
				},
				resource.Attribute{
					Name:        "balancer_protocol",
					Description: `(Required) transport protocol that will be accepted for all incoming requests to the selected load balancer listener. ` + "`" + `HTTP` + "`" + ` or ` + "`" + `HTTPS` + "`" + `. If set to HTTPS then you must also set the server ` + "`" + `certificates` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) The port on which the Load Balancer is listening.`,
				},
				resource.Attribute{
					Name:        "server_protocol",
					Description: `(Required) The protocol to be used for routing traffic to the origin servers in the server pool. ` + "`" + `HTTP` + "`" + ` or ` + "`" + `HTTPS` + "`" + `. If set to ` + "`" + `HTTPS` + "`" + ` then you must include a Trusted Certificate Policy in the ` + "`" + `policies` + "`" + ``,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Boolean flag to enable or disable the Listener. Default is ` + "`" + `true` + "`" + ` (enabled).`,
				},
				resource.Attribute{
					Name:        "path_prefixes",
					Description: `(Optional) List of paths to configure the listener to accept only requests that are targeted to a specific path within the URI of the request.`,
				},
				resource.Attribute{
					Name:        "polices",
					Description: `(Optional) List of the Load Balancer Policy URIs to apply to the listener.`,
				},
				resource.Attribute{
					Name:        "server_pool",
					Description: `(Optional) URI of the Server Pool resource to which the load balancer distributes requests.`,
				},
				resource.Attribute{
					Name:        "certificates",
					Description: `(Optional) The URI of the server security certificate.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) List of tags.`,
				},
				resource.Attribute{
					Name:        "virtual_hosts",
					Description: `(Optional) Configure the listener to only accept URI requests that include the host names listed in this field. ## Additional Attributes In addition to the above, the following values are exported:`,
				},
				resource.Attribute{
					Name:        "operational_details",
					Description: `Description of the operational state.`,
				},
				resource.Attribute{
					Name:        "parent_listener",
					Description: `The parent Listener.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the Origin Server Pool.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `The Uniform Resource Identifier for the Listner. ## Import Listeners can be imported using the a combinable of the resource region, load balancer name and policy name and in the format ` + "`" + `region/loadbalancer/name` + "`" + ` ` + "`" + `` + "`" + `` + "`" + `shell $ terraform import opc_lbaas_listener.listener1 uscom-central-1/lb1/example-listener1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opc_lbaas_load_balancer",
			Category:         "Load Balancer Classic Resources",
			ShortDescription: `Creates an Oracle Load Balancer Classic instance.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
				"classic",
				"lbaas",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Load Balancer.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The region in which to create the Load Balancer, e.g. ` + "`" + `uscom-central-1` + "`" + ``,
				},
				resource.Attribute{
					Name:        "scheme",
					Description: `(Required) Set to either ` + "`" + `INTERNET_FACING` + "`" + ` or ` + "`" + `INTERNAL` + "`" + ` - ` + "`" + `INTERNET_FACING` + "`" + ` - Create an internet-facing load balancer in a given IP network. - ` + "`" + `INTERNAL` + "`" + ` - Create an internal load balancer in a given IP network for sole consumption of other clients inside the same network.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A short description for the load balancer. The description must not exceed 1000 characters.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Boolean flag to enable or disable the Load Balancer. Default is ` + "`" + `true` + "`" + ` (enabled).`,
				},
				resource.Attribute{
					Name:        "ip_network",
					Description: `(Optional) Fully qualified three part name of the IP network to be associated with the load balancer.`,
				},
				resource.Attribute{
					Name:        "parent_load_balancer",
					Description: `(Optional) Select a parent load balancer if you want to create a dependent load balancer.`,
				},
				resource.Attribute{
					Name:        "permitted_clients",
					Description: `(Optional) List of permitted client IP addresses or CIDR ranges which can connect to this load balancer on the configured Listener ports. If not set all connections are permitted.`,
				},
				resource.Attribute{
					Name:        "permitted_methods",
					Description: `(Optional) List of permitted HTTP methods. e.g. ` + "`" + `GET` + "`" + `, ` + "`" + `POST` + "`" + `, ` + "`" + `PUT` + "`" + `, ` + "`" + `PATCH` + "`" + `, ` + "`" + `DELETE` + "`" + `, ` + "`" + `HEAD` + "`" + ` or you can also create your own custom methods. Requests with methods not listed in this field will result in a 403 (unauthorized access) response.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) List of tags. ## Additional Attributes In addition to the above, the following values are exported:`,
				},
				resource.Attribute{
					Name:        "operational_details",
					Description: `Description of the operational state.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the Origin Server Pool.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the Origin Server Pool.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `The Uniform Resource Identifier for the Load Balancer. ## Import Load Balancers can be imported using the a combinable of the resource ` + "`" + `region` + "`" + ` and ` + "`" + `name` + "`" + ` in the format ` + "`" + `region/name` + "`" + ` ` + "`" + `` + "`" + `` + "`" + `shell $ terraform import opc_lbaas_load_balancer.lb1 uscom-central-1/example-lb1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opc_lbaas_policy",
			Category:         "Load Balancer Classic Resources",
			ShortDescription: `Creates an Oracle Load Balancer Classic Policy.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
				"classic",
				"lbaas",
				"policy",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Listener.`,
				},
				resource.Attribute{
					Name:        "load_balancer",
					Description: `(Required) The parent Load Balancer the Listener. Include one, and only one, of:`,
				},
				resource.Attribute{
					Name:        "application_cookie_stickiness_policy",
					Description: `see [Application Cookie Stickiness Policy](#application-cookie-stickiness-policy)`,
				},
				resource.Attribute{
					Name:        "cloudgate_policy",
					Description: `see [CloudGate Policy](#cloudgate-policy)`,
				},
				resource.Attribute{
					Name:        "load_balancer_cookie_stickiness_policy",
					Description: `see [Load Balancer Cookie Stickiness Policy](#load-balancer-cookie-stickiness-policy)`,
				},
				resource.Attribute{
					Name:        "load_balancing_mechanism_policy",
					Description: `see [Load Balancing Mechanism Policy](#load-balancing-mechanism-policy)`,
				},
				resource.Attribute{
					Name:        "rate_limiting_request_policy",
					Description: `see [Rate Limiting Request Policy](#rate-limiting-request-policy)`,
				},
				resource.Attribute{
					Name:        "redirect_policy",
					Description: `see [Redirect Policy](#redirect-policy)`,
				},
				resource.Attribute{
					Name:        "resource_access_control_policy",
					Description: `see [Resource Access Control Policy](#resource-access-control-policy)`,
				},
				resource.Attribute{
					Name:        "set_request_header_policy",
					Description: `see [Set Request Header Policy](#set-request-header-policy)`,
				},
				resource.Attribute{
					Name:        "ssl_negotiation_policy",
					Description: `see [SSL Negotiation Policy](#set-negotiation-policy)`,
				},
				resource.Attribute{
					Name:        "trusted_certificate_policy",
					Description: `see [Trusted Certificate Policy](#trusted-certificate-policy) ### Application Cookie Stickiness Policy Enable session stickiness (session affinity) for any request based on a given cookie name specified in the policy. #### Example ` + "`" + `` + "`" + `` + "`" + `hcl resource "opc_lbaas_policy" "application_cookie_stickiness_policy" { load_balancer = "${opc_lbaas_load_balancer.lb1.id}" name = "example_application_cookie_stickiness_policy" application_cookie_stickiness_policy { cookie_name = "MY_APP_COOKIE" } } ` + "`" + `` + "`" + `` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "cookie_name",
					Description: `(Required) Name of the application cookie used to control how long the load balancer will continue to route requests to the same origin server. ### CloudGate Policy protect resources/applications with the help of CloudGate module available in Load Balancer. These headers will enable CloudGate to lookup for the application and the policy present under the appropriate IDCS Tenant containing information for the protection mechanism to be enforced. #### Example ` + "`" + `` + "`" + `` + "`" + `hcl resource "opc_lbaas_policy" "cloudgate_policy" { load_balancer = "${opc_lbaas_load_balancer.lb1.id}" name = "example_cloudgate_policy" cloudgate_policy { virtual_hostname_for_policy_attribution = "host1.example.com" } } ` + "`" + `` + "`" + `` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "virtual_hostname_for_policy_attribution",
					Description: `(Required) Host name needed by CloudGate to enforce OAuth policies. ### Load Balancer Cookie Stickiness Policy Enables session stickiness (session affinity) for all requests for a given period of time specified in the policy. #### Example ` + "`" + `` + "`" + `` + "`" + `hcl resource "opc_lbaas_policy" "load_balancer_cookie_stickiness_policy" { load_balancer = "${opc_lbaas_load_balancer.lb1.id}" name = "example_load_balancer_cookie_stickiness_policy" load_balancer_cookie_stickiness_policy { cookie_expiration_period = 60 } } ` + "`" + `` + "`" + `` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "cookie_expiration_period",
					Description: `(Required) The time period, in seconds, after which the cookie should be considered stale. If the value is zero or negative the stickiness session lasts for the duration of the browser session. ### Load Balancing Mechanism Policy Specify a load balancing mechanism for distributing client requests across multiple origin servers #### Example ` + "`" + `` + "`" + `` + "`" + `hcl resource "opc_lbaas_policy" "load_balancing_mechanism_policy" { load_balancer = "${opc_lbaas_load_balancer.lb1.id}" name = "example_load_balancing_mechanism_policy" load_balancing_mechanism_policy { load_balancing_mechanism = "round_robin" } } ` + "`" + `` + "`" + `` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "load_balancing_mechanism",
					Description: `(Required) Supported options are ` + "`" + `round_robin` + "`" + `, ` + "`" + `least_conn` + "`" + `, and ` + "`" + `ip_hash` + "`" + `. ### Rate Limiting Request Policy Limits the number of requests that can be processed per second by the load balancer. #### Example ` + "`" + `` + "`" + `` + "`" + `hcl resource "opc_lbaas_policy" "rate_limiting_request_policy" { load_balancer = "${opc_lbaas_load_balancer.lb1.id}" name = "example_rate_limiting_request_policy" rate_limiting_request_policy { requests_per_second = 1 burst_size = 10 delay_excessive_requests = true zone = "examplezone" } } ` + "`" + `` + "`" + `` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "burst_size",
					Description: `(Required) The number of requests that can be delayed until it exceeds the maximum number specified as burst size in which case the request is terminated.`,
				},
				resource.Attribute{
					Name:        "delay_excessive_requests",
					Description: `(Required) delay excessive requests while requests are being limited.`,
				},
				resource.Attribute{
					Name:        "requests_per_second",
					Description: `(Required) Maximum number of requests per second.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) Name of the shared memory zone.`,
				},
				resource.Attribute{
					Name:        "http_error_code",
					Description: `(Optional) Status code to return in response to rejected requests. You can specify any status code between 405 to 599. Default is ` + "`" + `503` + "`" + ``,
				},
				resource.Attribute{
					Name:        "logging_level",
					Description: `(Optional) Logging level for cases when the server refuses to process requests due to rate exceeding, or delays request processing. Can be one of ` + "`" + `info` + "`" + `, ` + "`" + `notice` + "`" + `, ` + "`" + `warn` + "`" + `, or ` + "`" + `error` + "`" + `. Default is ` + "`" + `warn` + "`" + ``,
				},
				resource.Attribute{
					Name:        "rate_limiting_criteria",
					Description: `(Optional) Criteria based on which requests will be throttled. Default is ` + "`" + `server` + "`" + ` - ` + "`" + `server` + "`" + ` - limit the requests processed by the virtual server - ` + "`" + `remote_address` + "`" + ` - limit the processing rate of requests coming from a single IP address. - ` + "`" + `host` + "`" + ` - limit the processing rate of requests coming from a host.`,
				},
				resource.Attribute{
					Name:        "zone_memory_size",
					Description: `(Optional) Size of the shared memory occupied by the zone. Default is ` + "`" + `10` + "`" + ` ### Redirect Policy Redirects all requests to this load balancer to a specific URI. #### Example ` + "`" + `` + "`" + `` + "`" + `hcl resource "opc_lbaas_policy" "redirect_policy" { load_balancer = "${opc_lbaas_load_balancer.lb1.id}" name = "example_redirect_policy" redirect_policy { redirect_uri = "https://redirect.example.com" response_code = 306 } } ` + "`" + `` + "`" + `` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "redirect_uri",
					Description: `(Required) redirected requests to the specified URI.`,
				},
				resource.Attribute{
					Name:        "response_code",
					Description: `(Required) The exact 3xx response code to use when redirecting ### Resource Access Control Policy Controls what clients have access to the load balancer, based on the IP address or the CIDR range of the incoming request. #### Example ` + "`" + `` + "`" + `` + "`" + `hcl resource "opc_lbaas_policy" "resource_access_control_policy" { load_balancer = "${opc_lbaas_load_balancer.lb1.id}" name = "example_resource_access_control_policy" resource_access_control_policy { disposition = "DENY_ALL" permitted_clients = ["10.0.0.0/16"] } } ` + "`" + `` + "`" + `` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "disposition",
					Description: `(Required) Default policy. ` + "`" + `DENY_ALL` + "`" + ` or ` + "`" + `PERMIT_ALL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "denied_clients",
					Description: `(Optional) List of IP address or CIDR ranges identifying clients from which requests must be accepted by the Load Balancer`,
				},
				resource.Attribute{
					Name:        "permitted_clients",
					Description: `(Optional) IP address or CIDR ranges identifying clients from which requests must be denied by the Load Balancer. ### Set Request Header Policy Inserts additional information into the standard HTTP headers of requests forwarded to a server pool. #### Example ` + "`" + `` + "`" + `` + "`" + `hcl resource "opc_lbaas_policy" "set_request_header_policy" { load_balancer = "${opc_lbaas_load_balancer.lb1.id}" name = "example_set_request_header_policy" set_request_header_policy { header_name = "X-Custom-Header" value = "foo-bar " action_when_header_exists = "OVERWRITE" action_when_header_value_is = ["bar", "foo"] } } ` + "`" + `` + "`" + `` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "header_name",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "action_when_header_exists",
					Description: `(Required) action to be taken when a header exists in the request. Options: ` + "`" + `NOOP` + "`" + `, ` + "`" + `PREPEND` + "`" + `, ` + "`" + `APPEND` + "`" + `, ` + "`" + `OVERWRITE` + "`" + `, ` + "`" + `CLEAR` + "`" + ``,
				},
				resource.Attribute{
					Name:        "action_when_header_value_is",
					Description: `(Optional) List if header values. Action is taken only when the header exists in the request and the header value matches one of the values provided.`,
				},
				resource.Attribute{
					Name:        "action_when_header_value_is_not",
					Description: `(Optional) List if header values. Action is taken only when the header exists in the request and the header value does not match any of the values provided. ### SSL Negotiation Policy Define specific SSL protocols, ciphers, and server order preference for the secure connection #### Example ` + "`" + `` + "`" + `` + "`" + `hcl resource "opc_lbaas_policy" "ssl_negotiation_policy" { load_balancer = "${opc_lbaas_load_balancer.lb1.id}" name = "example_ssl_negitiation_policy" ssl_negotiation_policy { port = 8022 server_order_preference = "ENABLED" ssl_protocol = ["SSLv3", "TLSv1.2"] ssl_ciphers = ["AES256-SHA"] } } ` + "`" + `` + "`" + `` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "ssl_protocol",
					Description: `(Required) Security protocols supported for incoming secure client connections to the associated listener. Supported options are ` + "`" + `SSLv2` + "`" + `, ` + "`" + `SSLv3` + "`" + `, ` + "`" + `TLSv1` + "`" + `, ` + "`" + `TLSv1.1` + "`" + `, ` + "`" + `TLSv1.2` + "`" + ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The load balancer port for the the SSL protocols and the SSL ciphers.`,
				},
				resource.Attribute{
					Name:        "server_order_preference",
					Description: `(Optional) enable or disable the server order preference for secure connections to associated Listener. ` + "`" + `ENABLED` + "`" + ` or ` + "`" + `DISABLED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ssl_ciphers",
					Description: `(Optional) List of SSL ciphers supported for incoming secure client connections to the associated Listener. ### Trusted Certificate Policy Identifies a trusted certificate, which the load balancer uses when making a secure connection to the compute instances in the server pool. If you are configuring a secure connection (HTTPS or SSL) between the load balancer and the origin servers, you must add this policy to the load balancer. #### Example ` + "`" + `` + "`" + `` + "`" + `hcl resource "opc_lbaas_policy" "trusted_certificate_policy" { load_balancer = "${opc_lbaas_load_balancer.lb1.id}" name = "example_trusted_certificate_policy" trusted_certificate_policy { trusted_certificate = "${opc_lbaas_ssl_certificate.cert1.uri}" } } ` + "`" + `` + "`" + `` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "trusted_certificate",
					Description: `(Required) URI of the SSL Certificate ## Additional Attributes In addition to the above, the following values are exported:`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the Policy.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The Type of the Policy.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `The Uniform Resource Identifier for the Policy. ## Import Policies can be imported using the a combinable of the resource region, load balancer name and policy name and in the format ` + "`" + `region/loadbalancer/name` + "`" + ` ` + "`" + `` + "`" + `` + "`" + `shell $ terraform import opc_lbaas_policy.policy1 uscom-central-1/lb1/example-policy1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opc_lbaas_server_pool",
			Category:         "Load Balancer Classic Resources",
			ShortDescription: `Creates an Oracle Load Balancer Classic Origin Server Pool.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
				"classic",
				"lbaas",
				"server",
				"pool",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Server Pool.`,
				},
				resource.Attribute{
					Name:        "load_balancer",
					Description: `(Required) The parent Load Balancer the Origin Server Pool.`,
				},
				resource.Attribute{
					Name:        "servers",
					Description: `(Required) List of servers in the Server Pool. To define the server in the server pool, provide IP address or DNS name of the compute instances, and port for load balancer to direct traffic to, in the format ` + "`" + `host:port` + "`" + ``,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Boolean flag to enable or disable the Server Pool. Default is ` + "`" + `true` + "`" + ` (enabled).`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `(Optional) Enables Load Balancer health check, see [Health Check Attributes](#health-check-attributes)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) List of tags.`,
				},
				resource.Attribute{
					Name:        "vnic_set",
					Description: `(Optional) Fully qualified three part name of a vNICSet to be associated with the server pool vNIC. Load Balancer uses this vNICSet to set the right ACLs to allow egress traffic from the load balancer. ### Heath Check Attributes The load balancer can perform regular health checks of the origin servers and route inbound traffic to the healthy origin servers.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Health check mechanism to use to test the origin servers. Default is ` + "`" + `http` + "`" + ` - ` + "`" + `http` + "`" + ` - sends an HTTP HEAD to the set ` + "`" + `path` + "`" + ` and checks response is in one of the ` + "`" + `accepted_return_codes` + "`" + ``,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) The path to check. Set the '/' the check all paths. Default is '/'`,
				},
				resource.Attribute{
					Name:        "accepted_return_codes",
					Description: `(Optional) List of HTTP response status codes that indicate the origin server is healthy. Accepted return codes can be one or more of the ` + "`" + `2xx` + "`" + `, ` + "`" + `3xx` + "`" + `, ` + "`" + `4xx` + "`" + `, or ` + "`" + `5xx` + "`" + ` codes. Default is ` + "`" + `["2xx","3xx"]` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Boolean flag to enable or disable the Health Checks. Default is ` + "`" + `true` + "`" + ` (enabled).`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional) The approximate interval, in seconds, that the load balancer will wait before sending the target request to each origin server, in the range 5 to 300 seconds. Default is ` + "`" + `30` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) The amount of time, in seconds, that the load balancer will wait without a response before identifying the origin server as unavailable. The timeout value must be less than the ` + "`" + `interval` + "`" + ` value and it should range between 2 to 60. Default is ` + "`" + `20` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `(Optional) The number of consecutive successful health checks required before moving the origin server to the healthy state. The value of healthy threshold ranges from 2 to 10. Default is ` + "`" + `5` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "unhealthy_threshold",
					Description: `(Optional) The number of consecutive health check failures required before moving the origin server to the unhealthy state. The value of unhealthy threshold ranges from 2 to 10. Default is ` + "`" + `7` + "`" + `. ## Additional Attributes In addition to the above, the following values are exported:`,
				},
				resource.Attribute{
					Name:        "operational_details",
					Description: `Description of the operational state.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the Origin Server Pool.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the Origin Server Pool.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `The Uniform Resource Identifier for the Server Pool. ## Import Origin Server Pools can be imported using the combination of the resource region, load balancer name, and policy name and in the format ` + "`" + `region/loadbalancer/name` + "`" + ` ` + "`" + `` + "`" + `` + "`" + `shell $ terraform import opc_lbaas_server_pool.serverpool1 uscom-central-1/lb1/example-serverpool1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opc_storage_container",
			Category:         "Object Storage Classic Resources",
			ShortDescription: `Creates and manages a Container in the Oracle Cloud Infrastructure Storage Classic service. ` + "`" + `storage_endpoint` + "`" + ` must be set in the provider or environment to manage these resources.`,
			Description:      ``,
			Keywords: []string{
				"object",
				"storage",
				"classic",
				"container",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Storage Container.`,
				},
				resource.Attribute{
					Name:        "read_acls",
					Description: `(Optional) The list of ACLs that grant read access. See [Setting Container ACLs](#setting-container-acls).`,
				},
				resource.Attribute{
					Name:        "write_acls",
					Description: `(Optional) The list of ACLs that grant write access. See [Setting Container ACLs](#setting-container-acls).`,
				},
				resource.Attribute{
					Name:        "primary_key",
					Description: `(Optional) The primary secret key value for temporary URLs.`,
				},
				resource.Attribute{
					Name:        "secondary_key",
					Description: `(Optional) The secondary secret key value for temporary URLs.`,
				},
				resource.Attribute{
					Name:        "allowed_origins",
					Description: `(Optional) List of origins that are allowed to make cross-origin requests.`,
				},
				resource.Attribute{
					Name:        "exposed_headers",
					Description: `(Optional) List of headers exposed to the user agent (e.g. browser) in the actual request response`,
				},
				resource.Attribute{
					Name:        "max_age",
					Description: `(Optional) Maximum age in seconds for the origin to hold the preflight results.`,
				},
				resource.Attribute{
					Name:        "quota_bytes",
					Description: `(Optional) Maximum size of the container, in bytes`,
				},
				resource.Attribute{
					Name:        "quota_count",
					Description: `(Optional) Maximum object count of the container`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) Additional object metadata headers. See [Container Metadata ](#container-metadata) below for more information. ## Setting Container ACLs The ` + "`" + `read_acl` + "`" + ` consists of a list of`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opc_storage_object",
			Category:         "Object Storage Classic Resources",
			ShortDescription: `Creates and manages a Object in an Oracle Cloud Infrastructure Storage Classic container. ` + "`" + `storage_endpoint` + "`" + ` must be set in the provider or environment to manage these resources.`,
			Description:      ``,
			Keywords: []string{
				"object",
				"storage",
				"classic",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Storage Object.`,
				},
				resource.Attribute{
					Name:        "container",
					Description: `(Required) The name of Storage Container the store the object in.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Optional) Raw content in string-form of the data.`,
				},
				resource.Attribute{
					Name:        "file",
					Description: `(Optional) File path for the content to use for data.`,
				},
				resource.Attribute{
					Name:        "copy_from",
					Description: `(Optional) name of an existing object used to create the new object as a copy. The value is in form ` + "`" + `container/object` + "`" + `. You must UTF-8-encode and then URL-encode the names of the container and object.`,
				},
				resource.Attribute{
					Name:        "content_disposition",
					Description: `(Optional) Set the HTTP ` + "`" + `Content-Disposition` + "`" + ` header to specify the override behaviour for the browser, e.g. ` + "`" + `inline` + "`" + ` or ` + "`" + `attachment` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "content_encoding",
					Description: `(Optional) set the HTTP ` + "`" + `Content-Encoding` + "`" + ` for the object.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `(Optional) set the MIME type for the object.`,
				},
				resource.Attribute{
					Name:        "delete_at",
					Description: `(Optional) The date and time in UNIX Epoch time stamp format when the system removes the object.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `(Optional) MD5 checksum value of the request body. Strongly Recommended.`,
				},
				resource.Attribute{
					Name:        "transfer_encoding",
					Description: `(Optional) Set to ` + "`" + `chunked` + "`" + ` to enable chunked transfer encoding.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) Additional object metadata headers. See [Object Metadata ](#object-metadata) below for more information. ## Attributes In addition to the attributes listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The combined container and object name path of the object.`,
				},
				resource.Attribute{
					Name:        "accept_ranges",
					Description: `Type of ranges that the object accepts.`,
				},
				resource.Attribute{
					Name:        "content_length",
					Description: `Length of the object in bytes.`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `Date and Time that the object was created/modified in ISO 8601.`,
				},
				resource.Attribute{
					Name:        "object_manifest",
					Description: `The dynamic large-object manifest object.`,
				},
				resource.Attribute{
					Name:        "timestamp",
					Description: `Date and Time in UNIX EPOCH when the account, container, or object was initially created at the current version.`,
				},
				resource.Attribute{
					Name:        "transaction_id",
					Description: `Transaction ID of the request. ## Object Metadata The ` + "`" + `metadata` + "`" + ` config defines a map of additional meta data header name value pairs. The additional meta data items are set HTTP Headers on the object in the form ` + "`" + `X-Object-Meta-{name}: {value}` + "`" + `, where ` + "`" + `{name}` + "`" + ` is the name of the metadata item ` + "`" + `{value}` + "`" + ` is the header content. For example: ` + "`" + `` + "`" + `` + "`" + `hcl metadata { "Foo-Bar" = "barfoo", "Sha256" = "e91ed4f93637379a7539cb5d8d0b5bca3972755de4f9371ab2e123e7b4c53680" } ` + "`" + `` + "`" + `` + "`" + ` ## Import Object's can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ terraform import opc_storage_object.default container/example ` + "`" + `` + "`" + `` + "`" + ` Please note though, importing a Storage Object does _not_ allow a user to modify the content, or attributes for the Storage Object. It is, however, possible to import a Storage Object, and replace the object with new content, or a copy of another Storage Object. It is also possible to import a Storage Object into Terraform in order to delete the object.`,
				},
			},
			Attributes: []resource.Argument{},
		},
	}

	resourcesMap = map[string]int{

		"opc_compute_acl":                       0,
		"opc_compute_image_list":                1,
		"opc_compute_image_list_entry":          2,
		"opc_compute_instance":                  3,
		"opc_compute_ip_address_association":    4,
		"opc_compute_ip_address_prefix_set":     5,
		"opc_compute_ip_address_reservation":    6,
		"opc_compute_ip_association":            7,
		"opc_compute_ip_network":                8,
		"opc_compute_ip_network_exchange":       9,
		"opc_compute_ip_reservation":            10,
		"opc_compute_machine_image":             11,
		"opc_compute_orchestrated_instance":     12,
		"opc_compute_route":                     13,
		"opc_compute_sec_rule":                  14,
		"opc_compute_security_application":      15,
		"opc_compute_security_association":      16,
		"opc_compute_security_ip_list":          17,
		"opc_compute_security_list":             18,
		"opc_compute_security_protocol":         19,
		"opc_compute_security_rule":             20,
		"opc_compute_ssh_key":                   21,
		"opc_compute_storage_volume":            22,
		"opc_compute_storage_volume_attachment": 23,
		"opc_compute_storage_volume_snapshot":   24,
		"opc_compute_vnic_set":                  25,
		"opc_compute_vpn_endpoint_v2":           26,
		"opc_lbaas_certificate":                 27,
		"opc_lbaas_listener":                    28,
		"opc_lbaas_load_balancer":               29,
		"opc_lbaas_policy":                      30,
		"opc_lbaas_server_pool":                 31,
		"opc_storage_container":                 32,
		"opc_storage_object":                    33,
	}
)

func GetResource(r string) (*resouce.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs]
}
