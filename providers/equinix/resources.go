package equinix

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "equinix_ecx_l2_connection",
			Category:         "Resources",
			ShortDescription: `Provides Equinix Fabric Layer 2 connection resource`,
			Description: `

Resource ` + "`" + `equinix_ecx_l2_connection` + "`" + ` allows creation and management of Equinix Fabric
layer 2 connections.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_ecx_l2_connection_accepter",
			Category:         "Resources",
			ShortDescription: `Provides Equinix Fabric Layer 2 connection accepter resource`,
			Description: `

Resource ` + "`" + `equinix_ecx_l2_connection_accepter` + "`" + ` is used to accept Equinix Fabric 
layer 2 connection on provider side.

Resource leverages Equinix Fabric integration with service providers.
Currently supported providers are:

* ` + "`" + `AWS` + "`" + ` (AWS Direct Connect)

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "connection_id",
					Description: `(Required) Identifier of Layer 2 connection that will be accepted`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `(Optional) Access Key used to accept connection on provider side`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Optional) Secret Key used to accept connection on provider side`,
				},
				resource.Attribute{
					Name:        "aws_profile",
					Description: `(Optional) AWS Profile Name for retrieving credentials from shared credentials file ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "aws_connection_id",
					Description: `Identifier of a hosted Direct Connect connection on AWS side, applicable for accepter resource with connections to AWS only`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "aws_connection_id",
					Description: `Identifier of a hosted Direct Connect connection on AWS side, applicable for accepter resource with connections to AWS only`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_ecx_l2_serviceprofile",
			Category:         "Resources",
			ShortDescription: `Provides Equinix Fabric Layer 2 service profile resource`,
			Description: `

Resource ` + "`" + `equinix_ecx_l2_serviceprofile` + "`" + ` is used to manage layer 2 service profiles
in Equinix Fabric.

This resource relies on the Equinix Cloud Exchange Fabric API. The parameters
and attributes available map to the fields described at
<https://developer.equinix.com/catalog/sellerv3#operation/getProfileByIdOrNameUsingGET>.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_network_acl_template",
			Category:         "Resources",
			ShortDescription: `Provides Equinix Network Edge ACL template resource`,
			Description: `

Resource ` + "`" + `equinix_network_acl_template` + "`" + ` allows creation and management of
Equinix Network Edge device Access Control List templates.

Device ACL templates give possibility to define set of rules will allowed inbound
traffic. Templates can be assigned to the network devices.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) ACL template name`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) ACL template description`,
				},
				resource.Attribute{
					Name:        "metro_code",
					Description: `(Required) ACL template location metro code`,
				},
				resource.Attribute{
					Name:        "inbound_rule",
					Description: `(Required) One or more rules to specify allowed inbound traffic. Rules are ordered, matching traffic rule stops processing subsequent ones.`,
				},
				resource.Attribute{
					Name:        "inbound_rule.#.subnets",
					Description: `(Required) Inbound traffic source IP subnets in CIDR format`,
				},
				resource.Attribute{
					Name:        "inbound_rule.#.protocol",
					Description: `(Required) Inbound traffic protocol. One of: ` + "`" + `IP` + "`" + `, ` + "`" + `TCP` + "`" + `, ` + "`" + `UDP` + "`" + ``,
				},
				resource.Attribute{
					Name:        "inbound_rule.#.src_port",
					Description: `(Required) Inbound traffic source ports. Allowed values are:`,
				},
				resource.Attribute{
					Name:        "inbound_rule.#.dst_port",
					Description: `(Required) Inbound traffic destination ports. Allowed values are:`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `Unique identifier of ACL template resource`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `Identifier of a network device where template was applied`,
				},
				resource.Attribute{
					Name:        "device_acl_status",
					Description: `Status of ACL template provisioning process on a device, where template was applied. One of:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "uuid",
					Description: `Unique identifier of ACL template resource`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `Identifier of a network device where template was applied`,
				},
				resource.Attribute{
					Name:        "device_acl_status",
					Description: `Status of ACL template provisioning process on a device, where template was applied. One of:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_network_bgp",
			Category:         "Resources",
			ShortDescription: `Provides Equinix Network Edge BGP peering resource.`,
			Description: `

Resource ` + "`" + `equinix_network_bgp` + "`" + ` allows creation and management of Equinix Network
Edge BGP peering configurations.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "connection_id",
					Description: `(Required) identifier of a connection established between network device and remote service provider that will be used for peering`,
				},
				resource.Attribute{
					Name:        "local_ip_address",
					Description: `(Required) IP address in CIDR format of a local device`,
				},
				resource.Attribute{
					Name:        "local_asn",
					Description: `(Required) Local ASN number`,
				},
				resource.Attribute{
					Name:        "remote_ip_address",
					Description: `(Required) IP address of remote peer`,
				},
				resource.Attribute{
					Name:        "remote_asn",
					Description: `(Required) Remote ASN number`,
				},
				resource.Attribute{
					Name:        "authentication_key",
					Description: `(Optional) shared key used for BGP peer authentication ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `BGP peering configuration unique identifier`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `unique identifier of a network device that is a local peer in a given BGP peering configuration`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `BGP peer state, one of:`,
				},
				resource.Attribute{
					Name:        "provisioning_status",
					Description: `BGP peering configuration provisioning status, one of:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "uuid",
					Description: `BGP peering configuration unique identifier`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `unique identifier of a network device that is a local peer in a given BGP peering configuration`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `BGP peer state, one of:`,
				},
				resource.Attribute{
					Name:        "provisioning_status",
					Description: `BGP peering configuration provisioning status, one of:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_network_device",
			Category:         "Resources",
			ShortDescription: `Provides Equinix Network Edge device resource`,
			Description: `

Resource ` + "`" + `equinix_network_device` + "`" + ` allows creation and management of Equinix Network
Edge virtual network devices.

Network Edge virtual network devices can be created in two modes:

* **managed** (default) where Equinix manages connectivity and services in the
device and customer gets limited access to the device
* **self-configured** where customer provisions and manages own services in the device
with less restricted access. Some device types are offered only in this mode

In addition to management modes, there are two software license modes available:

* **subscription**  where Equinix provides software license, including end-to-end
support, and bills for the service respectively.
* **BYOL** [bring your own license] where customer brings his own, already procured
device software license. There are no charges associated with such license.
It is the only licensing mode for *self-configured* devices

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Device name`,
				},
				resource.Attribute{
					Name:        "type_code",
					Description: `(Required) Device type code`,
				},
				resource.Attribute{
					Name:        "metro_code",
					Description: `(Required) Device location metro code`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional) Device hostname prefix`,
				},
				resource.Attribute{
					Name:        "package_code",
					Description: `(Required) Device software package code`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) Device software software version`,
				},
				resource.Attribute{
					Name:        "core_count",
					Description: `(Required) Number of CPU cores used by device,`,
				},
				resource.Attribute{
					Name:        "term_length",
					Description: `(Required) Device term length`,
				},
				resource.Attribute{
					Name:        "self_managed",
					Description: `(Optional) Boolean value that determines device management mode:`,
				},
				resource.Attribute{
					Name:        "byol",
					Description: `(Optional) Boolean value that determines device licensing mode:`,
				},
				resource.Attribute{
					Name:        "license_token",
					Description: `(Optional) License Token applicable for some device types in BYOL licensing mode`,
				},
				resource.Attribute{
					Name:        "license_file",
					Description: `(Optional) Path to the license file that will be uploaded and applied on a device. Applicable for some devices types in BYOL licensing mode`,
				},
				resource.Attribute{
					Name:        "throughput",
					Description: `(Optional) Device license throughput`,
				},
				resource.Attribute{
					Name:        "throughput_unit",
					Description: `(Optional) License throughput unit (Mbps or Gbps)`,
				},
				resource.Attribute{
					Name:        "account_number",
					Description: `(Required) Billing account number for a device`,
				},
				resource.Attribute{
					Name:        "notifications",
					Description: `(Required) List of email addresses that will receive device status notifications`,
				},
				resource.Attribute{
					Name:        "purchase_order_number",
					Description: `(Optional) Purchase order number associated with a device order`,
				},
				resource.Attribute{
					Name:        "order_reference",
					Description: `(Optional) Name/number used to identify device order on the invoice`,
				},
				resource.Attribute{
					Name:        "acl_template_id",
					Description: `(Optional) Identifier of an ACL template that will be applied on the device`,
				},
				resource.Attribute{
					Name:        "additional_bandwidth",
					Description: `(Optional) Additional Internet bandwidth, in Mbps, that will be allocated to the device (in addition to default 15Mbps)`,
				},
				resource.Attribute{
					Name:        "interface_count",
					Description: `(Optional) Number of network interfaces on a device. If not specified, default number for a given device type will be used`,
				},
				resource.Attribute{
					Name:        "vendor_configuration",
					Description: `(Optional) map of vendor specific configuration parameters for a device`,
				},
				resource.Attribute{
					Name:        "ssh-key",
					Description: `(Optional) definition of SSH key that will be provisioned on a device (max one key)`,
				},
				resource.Attribute{
					Name:        "secondary_device",
					Description: `(Optional) Definition of secondary device for redundant device configurations The ` + "`" + `secondary_device` + "`" + ` block supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Secondary device name`,
				},
				resource.Attribute{
					Name:        "metro_code",
					Description: `(Required) Metro location of a secondary device`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional) Secondary device hostname`,
				},
				resource.Attribute{
					Name:        "license_token",
					Description: `(Optional) License Token can be provided for some device types o the device`,
				},
				resource.Attribute{
					Name:        "license_file",
					Description: `(Optional) Path to the license file that will be uploaded and applied on a secondary device. Applicable for some devices types in BYOL licensing mode`,
				},
				resource.Attribute{
					Name:        "account_number",
					Description: `(Required) Billing account number for secondary device`,
				},
				resource.Attribute{
					Name:        "notifications",
					Description: `(Required) List of email addresses that will receive notifications about secondary device`,
				},
				resource.Attribute{
					Name:        "additional_bandwidth",
					Description: `(Optional) Additional Internet bandwidth, in Mbps, for a secondary device`,
				},
				resource.Attribute{
					Name:        "vendor_configuration",
					Description: `(Optional) map of vendor specific configuration parameters for a secondary device`,
				},
				resource.Attribute{
					Name:        "acl_template_id",
					Description: `Identifier of an ACL template that will be applied on a secondary device`,
				},
				resource.Attribute{
					Name:        "ssh-key",
					Description: `(Optional) up to one definition of SSH key that will be provisioned on a secondary device The ` + "`" + `ssh_key` + "`" + ` block supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) username associated with given key`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) reference by name to previously provisioned public SSH key ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `Device unique identifier`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Device provisioning status`,
				},
				resource.Attribute{
					Name:        "license_status",
					Description: `Device license registration status`,
				},
				resource.Attribute{
					Name:        "license_file_id",
					Description: `Unique identifier of applied license file`,
				},
				resource.Attribute{
					Name:        "ibx",
					Description: `Device location Equinix Business Exchange name`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Device location region`,
				},
				resource.Attribute{
					Name:        "acl_template_id",
					Description: `Unique identifier of applied ACL template`,
				},
				resource.Attribute{
					Name:        "ssh_ip_address",
					Description: `IP address of SSH enabled interface on the device`,
				},
				resource.Attribute{
					Name:        "ssh_ip_fqdn",
					Description: `FQDN of SSH enabled interface on the device`,
				},
				resource.Attribute{
					Name:        "redundancy_type",
					Description: `Device redundancy type applicable for HA devices, either primary or secondary`,
				},
				resource.Attribute{
					Name:        "redundant_id",
					Description: `Unique identifier for a redundant device applicable for HA devices`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `List of device interfaces`,
				},
				resource.Attribute{
					Name:        "interface.#.id",
					Description: `interface identifier`,
				},
				resource.Attribute{
					Name:        "interface.#.name",
					Description: `interface name`,
				},
				resource.Attribute{
					Name:        "interface.#.status",
					Description: `interface status (AVAILABLE, RESERVED, ASSIGNED)`,
				},
				resource.Attribute{
					Name:        "interface.#.operational_status",
					Description: `interface operational status (up or down)`,
				},
				resource.Attribute{
					Name:        "interface.#.mac_address",
					Description: `interface MAC address`,
				},
				resource.Attribute{
					Name:        "interface.#.ip_address",
					Description: `interface IP address`,
				},
				resource.Attribute{
					Name:        "interface.#.assigned_type",
					Description: `interface management type (Equinix Managed or empty)`,
				},
				resource.Attribute{
					Name:        "interface.#.type",
					Description: `interface type ## Timeouts This resource provides the following [Timeouts configuration](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) options:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "uuid",
					Description: `Device unique identifier`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Device provisioning status`,
				},
				resource.Attribute{
					Name:        "license_status",
					Description: `Device license registration status`,
				},
				resource.Attribute{
					Name:        "license_file_id",
					Description: `Unique identifier of applied license file`,
				},
				resource.Attribute{
					Name:        "ibx",
					Description: `Device location Equinix Business Exchange name`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Device location region`,
				},
				resource.Attribute{
					Name:        "acl_template_id",
					Description: `Unique identifier of applied ACL template`,
				},
				resource.Attribute{
					Name:        "ssh_ip_address",
					Description: `IP address of SSH enabled interface on the device`,
				},
				resource.Attribute{
					Name:        "ssh_ip_fqdn",
					Description: `FQDN of SSH enabled interface on the device`,
				},
				resource.Attribute{
					Name:        "redundancy_type",
					Description: `Device redundancy type applicable for HA devices, either primary or secondary`,
				},
				resource.Attribute{
					Name:        "redundant_id",
					Description: `Unique identifier for a redundant device applicable for HA devices`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `List of device interfaces`,
				},
				resource.Attribute{
					Name:        "interface.#.id",
					Description: `interface identifier`,
				},
				resource.Attribute{
					Name:        "interface.#.name",
					Description: `interface name`,
				},
				resource.Attribute{
					Name:        "interface.#.status",
					Description: `interface status (AVAILABLE, RESERVED, ASSIGNED)`,
				},
				resource.Attribute{
					Name:        "interface.#.operational_status",
					Description: `interface operational status (up or down)`,
				},
				resource.Attribute{
					Name:        "interface.#.mac_address",
					Description: `interface MAC address`,
				},
				resource.Attribute{
					Name:        "interface.#.ip_address",
					Description: `interface IP address`,
				},
				resource.Attribute{
					Name:        "interface.#.assigned_type",
					Description: `interface management type (Equinix Managed or empty)`,
				},
				resource.Attribute{
					Name:        "interface.#.type",
					Description: `interface type ## Timeouts This resource provides the following [Timeouts configuration](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) options:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_network_ssh_key",
			Category:         "Resources",
			ShortDescription: `Provides Equinix Network Edge SSH key resource`,
			Description: `

Resource ` + "`" + `equinix_network_ssh_key` + "`" + ` allows creation and management of Equinix Network
Edge SSH keys.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of SSH key used for identification`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(Required) The SSH public key. If this is a file, it can be read using the file interpolation function ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `The unique identifier of the key`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "uuid",
					Description: `The unique identifier of the key`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_network_ssh_user",
			Category:         "Resources",
			ShortDescription: `Provides Equinix Network Edge SSH user resource`,
			Description: `

Resource ` + "`" + `equinix_network_ssh_user` + "`" + ` allows creation and management of Equinix Network
Edge SSH users.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "username",
					Description: `(Required) SSH user login name`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) SSH user password`,
				},
				resource.Attribute{
					Name:        "device_ids",
					Description: `(Required) list of device identifiers to which user will have access ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `SSH user unique identifier`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "uuid",
					Description: `SSH user unique identifier`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"equinix_ecx_l2_connection":          0,
		"equinix_ecx_l2_connection_accepter": 1,
		"equinix_ecx_l2_serviceprofile":      2,
		"equinix_network_acl_template":       3,
		"equinix_network_bgp":                4,
		"equinix_network_device":             5,
		"equinix_network_ssh_key":            6,
		"equinix_network_ssh_user":           7,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
