package fortios

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "fortios_fmg_devicemanager_device",
			Category:         "FortiManager Resources",
			ShortDescription: `Provides a resource to add/delete online FortiGate to/from FortiManager`,
			Description:      ``,
			Keywords: []string{
				"fortimanager",
				"fmg",
				"devicemanager",
				"device",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "userid",
					Description: `(Required) User name.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Password.`,
				},
				resource.Attribute{
					Name:        "ipaddr",
					Description: `(Required) Fortigate's ipaddress.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `(Required) Fortigate's device name.`,
				},
				resource.Attribute{
					Name:        "adom",
					Description: `Name or ID of the ADOM where the command is to be executed on. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "userid",
					Description: `User name.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Password.`,
				},
				resource.Attribute{
					Name:        "ipaddr",
					Description: `Fortigate's ipaddress.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `Fortigate's device name.`,
				},
				resource.Attribute{
					Name:        "adom",
					Description: `Name or ID of the ADOM where the command is to be executed on.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "userid",
					Description: `User name.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Password.`,
				},
				resource.Attribute{
					Name:        "ipaddr",
					Description: `Fortigate's ipaddress.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `Fortigate's device name.`,
				},
				resource.Attribute{
					Name:        "adom",
					Description: `Name or ID of the ADOM where the command is to be executed on.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_fmg_devicemanager_install_device",
			Category:         "FortiManager Resources",
			ShortDescription: `Provides a resource to install devicemanager script from FortiManager to the related device`,
			Description:      ``,
			Keywords: []string{
				"fortimanager",
				"fmg",
				"devicemanager",
				"install",
				"device",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "target_devname",
					Description: `(Required) Target device name.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `Timeout for installing the script to the target, default: 3 minutes.`,
				},
				resource.Attribute{
					Name:        "adom",
					Description: `Source ADOM name. default is 'root'`,
				},
				resource.Attribute{
					Name:        "vdom",
					Description: `Vdom of managed device. default is 'root' ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "target_devname",
					Description: `Target device name.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `Timeout for installing the script to the target.`,
				},
				resource.Attribute{
					Name:        "adom",
					Description: `Source ADOM name.`,
				},
				resource.Attribute{
					Name:        "vdom",
					Description: `Vdom of managed device.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "target_devname",
					Description: `Target device name.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `Timeout for installing the script to the target.`,
				},
				resource.Attribute{
					Name:        "adom",
					Description: `Source ADOM name.`,
				},
				resource.Attribute{
					Name:        "vdom",
					Description: `Vdom of managed device.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_fmg_devicemanager_install_policypackage",
			Category:         "FortiManager Resources",
			ShortDescription: `Provides a resource to install devicemanager policy package from FortiManager to the related FortiGate`,
			Description:      ``,
			Keywords: []string{
				"fortimanager",
				"fmg",
				"devicemanager",
				"install",
				"policypackage",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "package_name",
					Description: `(Required) The installation package name.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `Timeout for installing the package to the target, default: 3 minutes.`,
				},
				resource.Attribute{
					Name:        "adom",
					Description: `Source ADOM name. default is 'root' ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "package_name",
					Description: `The installation package name.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `Timeout for installing the package to the target, default: 3 minutes.`,
				},
				resource.Attribute{
					Name:        "adom",
					Description: `Source ADOM name.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "package_name",
					Description: `The installation package name.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `Timeout for installing the package to the target, default: 3 minutes.`,
				},
				resource.Attribute{
					Name:        "adom",
					Description: `Source ADOM name.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_fmg_devicemanager_script",
			Category:         "FortiManager Resources",
			ShortDescription: `Provides a resource to configure devicemanager script for FortiManager.`,
			Description:      ``,
			Keywords: []string{
				"fortimanager",
				"fmg",
				"devicemanager",
				"script",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Script name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `Script content, only cli script is supported now`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `Script target, Enum: ["device_database", "remote_device", "adom_database"]`,
				},
				resource.Attribute{
					Name:        "adom",
					Description: `ADOM name. default is 'root'. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Script name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `Script content.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `Script target.`,
				},
				resource.Attribute{
					Name:        "adom",
					Description: `ADOM name.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Script name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `Script content.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `Script target.`,
				},
				resource.Attribute{
					Name:        "adom",
					Description: `ADOM name.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_fmg_devicemanager_script_execute",
			Category:         "FortiManager Resources",
			ShortDescription: `Provides a resource to execute devicemanager script on FortiManager.`,
			Description:      ``,
			Keywords: []string{
				"fortimanager",
				"fmg",
				"devicemanager",
				"script",
				"execute",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "script_name",
					Description: `(Required) Script name.`,
				},
				resource.Attribute{
					Name:        "target_devname",
					Description: `(Required) Target device name, which the script will be installed.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `Timeout(minute) for executing the script, default is 3 minutes.`,
				},
				resource.Attribute{
					Name:        "adom",
					Description: `Source ADOM name. default is 'root'`,
				},
				resource.Attribute{
					Name:        "vdom",
					Description: `Vdom of managed device. default is 'root' ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "script_name",
					Description: `Script name.`,
				},
				resource.Attribute{
					Name:        "target_devname",
					Description: `Target device name, which the script should be installed.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `Timeout(minute) for executing the script, default is 3 minutes.`,
				},
				resource.Attribute{
					Name:        "adom",
					Description: `Source ADOM name.`,
				},
				resource.Attribute{
					Name:        "vdom",
					Description: `Vdom of managed device.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "script_name",
					Description: `Script name.`,
				},
				resource.Attribute{
					Name:        "target_devname",
					Description: `Target device name, which the script should be installed.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `Timeout(minute) for executing the script, default is 3 minutes.`,
				},
				resource.Attribute{
					Name:        "adom",
					Description: `Source ADOM name.`,
				},
				resource.Attribute{
					Name:        "vdom",
					Description: `Vdom of managed device.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_fmg_firewall_object_address",
			Category:         "FortiManager Resources",
			ShortDescription: `Provides a resource to configure firewall object address for FortiManager.`,
			Description:      ``,
			Keywords: []string{
				"fortimanager",
				"fmg",
				"firewall",
				"object",
				"address",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Address name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of address, Enum: ["ipmask", "iprange", "fqdn"].`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comments.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `Fully Qualified Domain Name address.`,
				},
				resource.Attribute{
					Name:        "associated_intf",
					Description: `Network interface associated with address.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `IPv4 address/mask`,
				},
				resource.Attribute{
					Name:        "start_ip",
					Description: `First IP address (inclusive) in the range for the address.`,
				},
				resource.Attribute{
					Name:        "end_ip",
					Description: `Final IP address (inclusive) in the range for the address.`,
				},
				resource.Attribute{
					Name:        "allow_routing",
					Description: `Enable/disable use of this address in the static route configuration. default is "disable".`,
				},
				resource.Attribute{
					Name:        "adom",
					Description: `ADOM name. default is 'root'. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Address name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of address.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comments.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `Fully Qualified Domain Name address.`,
				},
				resource.Attribute{
					Name:        "associated_intf",
					Description: `Network interface associated with address.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `IPv4 address/mask`,
				},
				resource.Attribute{
					Name:        "start_ip",
					Description: `First IP address (inclusive) in the range for the address.`,
				},
				resource.Attribute{
					Name:        "end_ip",
					Description: `Final IP address (inclusive) in the range for the address.`,
				},
				resource.Attribute{
					Name:        "allow_routing",
					Description: `Enable/disable use of this address in the static route configuration.`,
				},
				resource.Attribute{
					Name:        "adom",
					Description: `ADOM name.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Address name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of address.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comments.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `Fully Qualified Domain Name address.`,
				},
				resource.Attribute{
					Name:        "associated_intf",
					Description: `Network interface associated with address.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `IPv4 address/mask`,
				},
				resource.Attribute{
					Name:        "start_ip",
					Description: `First IP address (inclusive) in the range for the address.`,
				},
				resource.Attribute{
					Name:        "end_ip",
					Description: `Final IP address (inclusive) in the range for the address.`,
				},
				resource.Attribute{
					Name:        "allow_routing",
					Description: `Enable/disable use of this address in the static route configuration.`,
				},
				resource.Attribute{
					Name:        "adom",
					Description: `ADOM name.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_fmg_firewall_object_ippool",
			Category:         "FortiManager Resources",
			ShortDescription: `Provides a resource to configure firewall object ippool for FortiManager.`,
			Description:      ``,
			Keywords: []string{
				"fortimanager",
				"fmg",
				"firewall",
				"object",
				"ippool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Ippool name.`,
				},
				resource.Attribute{
					Name:        "startip",
					Description: `(Required) First IPv4 address (inclusive) in the range for the address pool.`,
				},
				resource.Attribute{
					Name:        "endip",
					Description: `(Required) Final IPv4 address (inclusive) in the range for the address pool.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comments.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Ip pool type, Enum: ["overload", "one-to-one"].`,
				},
				resource.Attribute{
					Name:        "arp_intf",
					Description: `Select an interface that will reply to ARP requests.`,
				},
				resource.Attribute{
					Name:        "arp_reply",
					Description: `Enable/disable replying to ARP request, default is "enable".`,
				},
				resource.Attribute{
					Name:        "associated_intf",
					Description: `Associated interface name.`,
				},
				resource.Attribute{
					Name:        "adom",
					Description: `ADOM name. default is 'root'. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Ippool name.`,
				},
				resource.Attribute{
					Name:        "startip",
					Description: `First IPv4 address (inclusive) in the range for the address pool.`,
				},
				resource.Attribute{
					Name:        "endip",
					Description: `Final IPv4 address (inclusive) in the range for the address pool.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comments.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Ip pool type, Enum: ["overload", "one-to-one"].`,
				},
				resource.Attribute{
					Name:        "arp_intf",
					Description: `Select an interface that will reply to ARP requests.`,
				},
				resource.Attribute{
					Name:        "arp_reply",
					Description: `Enable/disable replying to ARP request.`,
				},
				resource.Attribute{
					Name:        "associated_intf",
					Description: `Associated interface name.`,
				},
				resource.Attribute{
					Name:        "adom",
					Description: `ADOM name.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Ippool name.`,
				},
				resource.Attribute{
					Name:        "startip",
					Description: `First IPv4 address (inclusive) in the range for the address pool.`,
				},
				resource.Attribute{
					Name:        "endip",
					Description: `Final IPv4 address (inclusive) in the range for the address pool.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comments.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Ip pool type, Enum: ["overload", "one-to-one"].`,
				},
				resource.Attribute{
					Name:        "arp_intf",
					Description: `Select an interface that will reply to ARP requests.`,
				},
				resource.Attribute{
					Name:        "arp_reply",
					Description: `Enable/disable replying to ARP request.`,
				},
				resource.Attribute{
					Name:        "associated_intf",
					Description: `Associated interface name.`,
				},
				resource.Attribute{
					Name:        "adom",
					Description: `ADOM name.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_fmg_firewall_object_service",
			Category:         "FortiManager Resources",
			ShortDescription: `Provides a resource to configure firewall object service for FortiManager.`,
			Description:      ``,
			Keywords: []string{
				"fortimanager",
				"fmg",
				"firewall",
				"object",
				"service",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Custom service name.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comments.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol type. Enum: ["TCP/UDP/SCTP", "ICMP", "ICMP6", "IP"]`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `Service category. Enum: ["", "File Access", "Authentication", "Email", "General", "Network Services", "Remote Access", "Tunneling", "VoIP, Messaging & Other Applications", "Web Access", "Web Proxy"]`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `Fully qualified domain name.`,
				},
				resource.Attribute{
					Name:        "iprange",
					Description: `Start and end of the IP range associated with service. Ip or Ip range(eg: 1.1.1.1-1.1.1.100).`,
				},
				resource.Attribute{
					Name:        "tcp_portrange",
					Description: `TCP port ranges. e.g: ["dst-port-range:src-port-range"]`,
				},
				resource.Attribute{
					Name:        "udp_portrange",
					Description: `UDP port ranges. e.g: ["dst-port-range:src-port-range"]`,
				},
				resource.Attribute{
					Name:        "sctp_portrange",
					Description: `SCTP port ranges. e.g: ["dst-port-range:src-port-range"]`,
				},
				resource.Attribute{
					Name:        "icmp_type",
					Description: `ICMP Type.`,
				},
				resource.Attribute{
					Name:        "icmp_code",
					Description: `ICMP Code.`,
				},
				resource.Attribute{
					Name:        "protocol_number",
					Description: `IP protocol number.`,
				},
				resource.Attribute{
					Name:        "adom",
					Description: `ADOM name. default is 'root'. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Custom service name.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comments.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol type.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `Service category.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `Fully qualified domain name.`,
				},
				resource.Attribute{
					Name:        "iprange",
					Description: `Start and end of the IP range associated with service.`,
				},
				resource.Attribute{
					Name:        "tcp_portrange",
					Description: `TCP port ranges.`,
				},
				resource.Attribute{
					Name:        "udp_portrange",
					Description: `UDP port ranges.`,
				},
				resource.Attribute{
					Name:        "sctp_portrange",
					Description: `SCTP port ranges.`,
				},
				resource.Attribute{
					Name:        "icmp_type",
					Description: `ICMP Type.`,
				},
				resource.Attribute{
					Name:        "icmp_code",
					Description: `ICMP Code.`,
				},
				resource.Attribute{
					Name:        "protocol_number",
					Description: `IP protocol number.`,
				},
				resource.Attribute{
					Name:        "adom",
					Description: `ADOM name.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Custom service name.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comments.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol type.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `Service category.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `Fully qualified domain name.`,
				},
				resource.Attribute{
					Name:        "iprange",
					Description: `Start and end of the IP range associated with service.`,
				},
				resource.Attribute{
					Name:        "tcp_portrange",
					Description: `TCP port ranges.`,
				},
				resource.Attribute{
					Name:        "udp_portrange",
					Description: `UDP port ranges.`,
				},
				resource.Attribute{
					Name:        "sctp_portrange",
					Description: `SCTP port ranges.`,
				},
				resource.Attribute{
					Name:        "icmp_type",
					Description: `ICMP Type.`,
				},
				resource.Attribute{
					Name:        "icmp_code",
					Description: `ICMP Code.`,
				},
				resource.Attribute{
					Name:        "protocol_number",
					Description: `IP protocol number.`,
				},
				resource.Attribute{
					Name:        "adom",
					Description: `ADOM name.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_fmg_firewall_object_vip",
			Category:         "FortiManager Resources",
			ShortDescription: `Provides a resource to configure firewall object virtual ip for FortiManager.`,
			Description:      ``,
			Keywords: []string{
				"fortimanager",
				"fmg",
				"firewall",
				"object",
				"vip",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Virtual IP name.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comments.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Virtual IP type, Enum: ["static-nat", "dns-translation", "fqdn"]`,
				},
				resource.Attribute{
					Name:        "ext_ip",
					Description: `IP address or address range on the external interface that you want to map to an address or address range on the destination network.`,
				},
				resource.Attribute{
					Name:        "ext_intf",
					Description: `Interface connected to the source network that receives the packets that will be forwarded to the destination network.`,
				},
				resource.Attribute{
					Name:        "arp_reply",
					Description: `Enable to respond to ARP requests for this virtual IP address. Enabled by default.`,
				},
				resource.Attribute{
					Name:        "mapped_ip",
					Description: `Mapped Ip address.`,
				},
				resource.Attribute{
					Name:        "config_default",
					Description: `Enable/Disable config default value. enabled by default.`,
				},
				resource.Attribute{
					Name:        "mapped_addr",
					Description: `Mapped FQDN address name.`,
				},
				resource.Attribute{
					Name:        "adom",
					Description: `ADOM name. default is 'root'. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Virtual IP name.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comments.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Virtual IP type.`,
				},
				resource.Attribute{
					Name:        "ext_ip",
					Description: `IP address or address range on the external interface that you want to map to an address or address range on the destination network.`,
				},
				resource.Attribute{
					Name:        "ext_intf",
					Description: `Interface connected to the source network that receives the packets that will be forwarded to the destination network.`,
				},
				resource.Attribute{
					Name:        "arp_reply",
					Description: `Enable to respond to ARP requests for this virtual IP address.`,
				},
				resource.Attribute{
					Name:        "mapped_ip",
					Description: `Mapped Ip address.`,
				},
				resource.Attribute{
					Name:        "config_default",
					Description: `Enable/Disable config default value.`,
				},
				resource.Attribute{
					Name:        "mapped_addr",
					Description: `Mapped FQDN address name.`,
				},
				resource.Attribute{
					Name:        "adom",
					Description: `ADOM name.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Virtual IP name.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comments.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Virtual IP type.`,
				},
				resource.Attribute{
					Name:        "ext_ip",
					Description: `IP address or address range on the external interface that you want to map to an address or address range on the destination network.`,
				},
				resource.Attribute{
					Name:        "ext_intf",
					Description: `Interface connected to the source network that receives the packets that will be forwarded to the destination network.`,
				},
				resource.Attribute{
					Name:        "arp_reply",
					Description: `Enable to respond to ARP requests for this virtual IP address.`,
				},
				resource.Attribute{
					Name:        "mapped_ip",
					Description: `Mapped Ip address.`,
				},
				resource.Attribute{
					Name:        "config_default",
					Description: `Enable/Disable config default value.`,
				},
				resource.Attribute{
					Name:        "mapped_addr",
					Description: `Mapped FQDN address name.`,
				},
				resource.Attribute{
					Name:        "adom",
					Description: `ADOM name.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_fmg_firewall_security_policy",
			Category:         "FortiManager Resources",
			ShortDescription: `Provides a resource to configure firewall security policy on FortiManager which could be installed to the FortiGate later`,
			Description:      ``,
			Keywords: []string{
				"fortimanager",
				"fmg",
				"firewall",
				"security",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Policy name.`,
				},
				resource.Attribute{
					Name:        "srcaddr",
					Description: `(Required) Source address and adress group names.`,
				},
				resource.Attribute{
					Name:        "srcintf",
					Description: `(Required) Incoming interface.`,
				},
				resource.Attribute{
					Name:        "dstaddr",
					Description: `(Required) Destination address and adress group names.`,
				},
				resource.Attribute{
					Name:        "dstintf",
					Description: `(Required) Outgoing interface.`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `(Required) Service and service group names.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) Policy action, default is deny. Enum: [allow, deny, ipsec].`,
				},
				resource.Attribute{
					Name:        "schedule",
					Description: `(Required) Schedule name.`,
				},
				resource.Attribute{
					Name:        "internet_service",
					Description: `Enable/disable use of Destination Internet Services for this policy.`,
				},
				resource.Attribute{
					Name:        "internet_service_id",
					Description: `Destination Internet Service ID.`,
				},
				resource.Attribute{
					Name:        "internet_service_src",
					Description: `Enable/disable use of Source Internet Services for this policy.`,
				},
				resource.Attribute{
					Name:        "internet_service_src_id",
					Description: `Source Internet Service ID.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `Names of individual users that can authenticate with this policy.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `Names of user groups that can authenticate with this policy.`,
				},
				resource.Attribute{
					Name:        "fsso",
					Description: `Enable/disable Fortinet Single Sign-On.`,
				},
				resource.Attribute{
					Name:        "rsso",
					Description: `Enable/disable RADIUS Single Sign-On.`,
				},
				resource.Attribute{
					Name:        "logtraffic",
					Description: `Enable or disable logging. Enum: [disable, all, utm]`,
				},
				resource.Attribute{
					Name:        "logtraffic_start",
					Description: `Record logs when a session starts and ends. Enum: [disable, enable]`,
				},
				resource.Attribute{
					Name:        "capture_packet",
					Description: `Enable/disable capture packets.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments.`,
				},
				resource.Attribute{
					Name:        "nat",
					Description: `Enable/disable source NAT.`,
				},
				resource.Attribute{
					Name:        "ippool",
					Description: `Enable/disable to use IP Pools for source NAT.`,
				},
				resource.Attribute{
					Name:        "poolname",
					Description: `IP Pool names.`,
				},
				resource.Attribute{
					Name:        "fixedport",
					Description: `Enable/disable to prevent source NAT from changing a session's source port.`,
				},
				resource.Attribute{
					Name:        "vpntunnel",
					Description: `Policy-based IPsec VPN: name of the IPsec VPN Phase 1.`,
				},
				resource.Attribute{
					Name:        "inbound",
					Description: `Policy-based IPsec VPN: only traffic from the remote network can initiate a VPN. Enum: [disable, enable]`,
				},
				resource.Attribute{
					Name:        "utm_status",
					Description: `Enable/disable to add one or more security profiles (AV, IPS, etc.) to the firewall policy.`,
				},
				resource.Attribute{
					Name:        "profile_type",
					Description: `Determine whether the firewall policy allows security profile groups or single profiles only. Enum: [single, group]`,
				},
				resource.Attribute{
					Name:        "av_profile",
					Description: `Name of an existing Antivirus profile.`,
				},
				resource.Attribute{
					Name:        "webfilter_profile",
					Description: `Name of an existing Web filter profile.`,
				},
				resource.Attribute{
					Name:        "application_list",
					Description: `Name of an existing Application list.`,
				},
				resource.Attribute{
					Name:        "ips_sensor",
					Description: `Name of an existing IPS sensor.`,
				},
				resource.Attribute{
					Name:        "waf_profile",
					Description: `Name of an existing Web application firewall profile.`,
				},
				resource.Attribute{
					Name:        "dnsfilter_profile",
					Description: `Name of an existing DNS filter profile.`,
				},
				resource.Attribute{
					Name:        "profile_protocol_options",
					Description: `Name of an existing Protocol options profile.`,
				},
				resource.Attribute{
					Name:        "profile_group",
					Description: `Name of profile group.`,
				},
				resource.Attribute{
					Name:        "traffic_shaper",
					Description: `Traffic shaper.`,
				},
				resource.Attribute{
					Name:        "traffic_shaper_reverse",
					Description: `Reverse traffic shaper.`,
				},
				resource.Attribute{
					Name:        "per_ip_shaper",
					Description: `Per-IP traffic shaper.`,
				},
				resource.Attribute{
					Name:        "package_name",
					Description: `The package name which the policy will be added to.`,
				},
				resource.Attribute{
					Name:        "adom",
					Description: `ADOM name. default is 'root'. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Policy name.`,
				},
				resource.Attribute{
					Name:        "srcaddr",
					Description: `Source address and adress group names.`,
				},
				resource.Attribute{
					Name:        "srcintf",
					Description: `Incoming interface.`,
				},
				resource.Attribute{
					Name:        "dstaddr",
					Description: `Destination address and adress group names.`,
				},
				resource.Attribute{
					Name:        "dstintf",
					Description: `Outgoing interface.`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `Service and service group names.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Policy action.`,
				},
				resource.Attribute{
					Name:        "schedule",
					Description: `Schedule name.`,
				},
				resource.Attribute{
					Name:        "internet_service",
					Description: `Enable/disable use of Destination Internet Services for this policy.`,
				},
				resource.Attribute{
					Name:        "internet_service_id",
					Description: `Destination Internet Service ID.`,
				},
				resource.Attribute{
					Name:        "internet_service_src",
					Description: `Enable/disable use of Source Internet Services for this policy.`,
				},
				resource.Attribute{
					Name:        "internet_service_src_id",
					Description: `Source Internet Service ID.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `Names of individual users that can authenticate with this policy.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `Names of user groups that can authenticate with this policy.`,
				},
				resource.Attribute{
					Name:        "fsso",
					Description: `Enable/disable Fortinet Single Sign-On.`,
				},
				resource.Attribute{
					Name:        "rsso",
					Description: `Enable/disable RADIUS Single Sign-On.`,
				},
				resource.Attribute{
					Name:        "logtraffic",
					Description: `Enable or disable logging.`,
				},
				resource.Attribute{
					Name:        "logtraffic_start",
					Description: `Record logs when a session starts and ends.`,
				},
				resource.Attribute{
					Name:        "capture_packet",
					Description: `Enable/disable capture packets.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments.`,
				},
				resource.Attribute{
					Name:        "nat",
					Description: `Enable/disable source NAT.`,
				},
				resource.Attribute{
					Name:        "ippool",
					Description: `Enable/disable to use IP Pools for source NAT.`,
				},
				resource.Attribute{
					Name:        "poolname",
					Description: `IP Pool names.`,
				},
				resource.Attribute{
					Name:        "fixedport",
					Description: `Enable/disable to prevent source NAT from changing a session's source port.`,
				},
				resource.Attribute{
					Name:        "vpntunnel",
					Description: `Policy-based IPsec VPN: name of the IPsec VPN Phase 1.`,
				},
				resource.Attribute{
					Name:        "inbound",
					Description: `Policy-based IPsec VPN: only traffic from the remote network can initiate a VPN.`,
				},
				resource.Attribute{
					Name:        "utm_status",
					Description: `Enable/disable to add one or more security profiles (AV, IPS, etc.) to the firewall policy.`,
				},
				resource.Attribute{
					Name:        "profile_type",
					Description: `Determine whether the firewall policy allows security profile groups or single profiles only.`,
				},
				resource.Attribute{
					Name:        "av_profile",
					Description: `Name of an existing Antivirus profile.`,
				},
				resource.Attribute{
					Name:        "webfilter_profile",
					Description: `Name of an existing Web filter profile.`,
				},
				resource.Attribute{
					Name:        "application_list",
					Description: `Name of an existing Application list.`,
				},
				resource.Attribute{
					Name:        "ips_sensor",
					Description: `Name of an existing IPS sensor.`,
				},
				resource.Attribute{
					Name:        "waf_profile",
					Description: `Name of an existing Web application firewall profile.`,
				},
				resource.Attribute{
					Name:        "dnsfilter_profile",
					Description: `Name of an existing DNS filter profile.`,
				},
				resource.Attribute{
					Name:        "profile_protocol_options",
					Description: `Name of an existing Protocol options profile.`,
				},
				resource.Attribute{
					Name:        "profile_group",
					Description: `Name of profile group.`,
				},
				resource.Attribute{
					Name:        "traffic_shaper",
					Description: `Traffic shaper.`,
				},
				resource.Attribute{
					Name:        "traffic_shaper_reverse",
					Description: `Reverse traffic shaper.`,
				},
				resource.Attribute{
					Name:        "per_ip_shaper",
					Description: `Per-IP traffic shaper.`,
				},
				resource.Attribute{
					Name:        "package_name",
					Description: `The package name which the policy will be added to.`,
				},
				resource.Attribute{
					Name:        "adom",
					Description: `ADOM name.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Policy name.`,
				},
				resource.Attribute{
					Name:        "srcaddr",
					Description: `Source address and adress group names.`,
				},
				resource.Attribute{
					Name:        "srcintf",
					Description: `Incoming interface.`,
				},
				resource.Attribute{
					Name:        "dstaddr",
					Description: `Destination address and adress group names.`,
				},
				resource.Attribute{
					Name:        "dstintf",
					Description: `Outgoing interface.`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `Service and service group names.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Policy action.`,
				},
				resource.Attribute{
					Name:        "schedule",
					Description: `Schedule name.`,
				},
				resource.Attribute{
					Name:        "internet_service",
					Description: `Enable/disable use of Destination Internet Services for this policy.`,
				},
				resource.Attribute{
					Name:        "internet_service_id",
					Description: `Destination Internet Service ID.`,
				},
				resource.Attribute{
					Name:        "internet_service_src",
					Description: `Enable/disable use of Source Internet Services for this policy.`,
				},
				resource.Attribute{
					Name:        "internet_service_src_id",
					Description: `Source Internet Service ID.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `Names of individual users that can authenticate with this policy.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `Names of user groups that can authenticate with this policy.`,
				},
				resource.Attribute{
					Name:        "fsso",
					Description: `Enable/disable Fortinet Single Sign-On.`,
				},
				resource.Attribute{
					Name:        "rsso",
					Description: `Enable/disable RADIUS Single Sign-On.`,
				},
				resource.Attribute{
					Name:        "logtraffic",
					Description: `Enable or disable logging.`,
				},
				resource.Attribute{
					Name:        "logtraffic_start",
					Description: `Record logs when a session starts and ends.`,
				},
				resource.Attribute{
					Name:        "capture_packet",
					Description: `Enable/disable capture packets.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments.`,
				},
				resource.Attribute{
					Name:        "nat",
					Description: `Enable/disable source NAT.`,
				},
				resource.Attribute{
					Name:        "ippool",
					Description: `Enable/disable to use IP Pools for source NAT.`,
				},
				resource.Attribute{
					Name:        "poolname",
					Description: `IP Pool names.`,
				},
				resource.Attribute{
					Name:        "fixedport",
					Description: `Enable/disable to prevent source NAT from changing a session's source port.`,
				},
				resource.Attribute{
					Name:        "vpntunnel",
					Description: `Policy-based IPsec VPN: name of the IPsec VPN Phase 1.`,
				},
				resource.Attribute{
					Name:        "inbound",
					Description: `Policy-based IPsec VPN: only traffic from the remote network can initiate a VPN.`,
				},
				resource.Attribute{
					Name:        "utm_status",
					Description: `Enable/disable to add one or more security profiles (AV, IPS, etc.) to the firewall policy.`,
				},
				resource.Attribute{
					Name:        "profile_type",
					Description: `Determine whether the firewall policy allows security profile groups or single profiles only.`,
				},
				resource.Attribute{
					Name:        "av_profile",
					Description: `Name of an existing Antivirus profile.`,
				},
				resource.Attribute{
					Name:        "webfilter_profile",
					Description: `Name of an existing Web filter profile.`,
				},
				resource.Attribute{
					Name:        "application_list",
					Description: `Name of an existing Application list.`,
				},
				resource.Attribute{
					Name:        "ips_sensor",
					Description: `Name of an existing IPS sensor.`,
				},
				resource.Attribute{
					Name:        "waf_profile",
					Description: `Name of an existing Web application firewall profile.`,
				},
				resource.Attribute{
					Name:        "dnsfilter_profile",
					Description: `Name of an existing DNS filter profile.`,
				},
				resource.Attribute{
					Name:        "profile_protocol_options",
					Description: `Name of an existing Protocol options profile.`,
				},
				resource.Attribute{
					Name:        "profile_group",
					Description: `Name of profile group.`,
				},
				resource.Attribute{
					Name:        "traffic_shaper",
					Description: `Traffic shaper.`,
				},
				resource.Attribute{
					Name:        "traffic_shaper_reverse",
					Description: `Reverse traffic shaper.`,
				},
				resource.Attribute{
					Name:        "per_ip_shaper",
					Description: `Per-IP traffic shaper.`,
				},
				resource.Attribute{
					Name:        "package_name",
					Description: `The package name which the policy will be added to.`,
				},
				resource.Attribute{
					Name:        "adom",
					Description: `ADOM name.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_fmg_firewall_security_policypackage",
			Category:         "FortiManager Resources",
			ShortDescription: `Provides a resource to configure firewall security policypackage on FortiManager which could be installed to the FortiGate later`,
			Description:      ``,
			Keywords: []string{
				"fortimanager",
				"fmg",
				"firewall",
				"security",
				"policypackage",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Security policy package name.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `The installation target.`,
				},
				resource.Attribute{
					Name:        "adom",
					Description: `Source ADOM name. default is 'root'`,
				},
				resource.Attribute{
					Name:        "vdom",
					Description: `Vdom of managed device. default is 'root'`,
				},
				resource.Attribute{
					Name:        "inspection_mode",
					Description: `Inspection Mode. Enum:[flow, proxy]. default is 'flow' ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Security policy package name.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `The installation target.`,
				},
				resource.Attribute{
					Name:        "adom",
					Description: `Source ADOM name.`,
				},
				resource.Attribute{
					Name:        "vdom",
					Description: `Vdom of managed device.`,
				},
				resource.Attribute{
					Name:        "inspection_mode",
					Description: `Inspection Mode.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Security policy package name.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `The installation target.`,
				},
				resource.Attribute{
					Name:        "adom",
					Description: `Source ADOM name.`,
				},
				resource.Attribute{
					Name:        "vdom",
					Description: `Vdom of managed device.`,
				},
				resource.Attribute{
					Name:        "inspection_mode",
					Description: `Inspection Mode.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_fmg_jsonrpc_request",
			Category:         "FortiManager Resources",
			ShortDescription: `Provides a resource to handle JSON RPC request for FortiManager. Only used for debugging and testing.`,
			Description:      ``,
			Keywords: []string{
				"fortimanager",
				"fmg",
				"jsonrpc",
				"request",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "json_content",
					Description: `(required) JSON RPC request, which should contain 'method' and 'params' parameters. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "json_content",
					Description: `JSON RPC request, which should contain 'method' and 'params' parameters.`,
				},
				resource.Attribute{
					Name:        "response",
					Description: `JSON RPC request response data.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "json_content",
					Description: `JSON RPC request, which should contain 'method' and 'params' parameters.`,
				},
				resource.Attribute{
					Name:        "response",
					Description: `JSON RPC request response data.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_fmg_object_adom_revision",
			Category:         "FortiManager Resources",
			ShortDescription: `Provides a resource to configure object adom revision for FortiManager.`,
			Description:      ``,
			Keywords: []string{
				"fortimanager",
				"fmg",
				"object",
				"adom",
				"revision",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Adom revision name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `Who created this adom revision.`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `lock. 0 means unlock and 1 means locked.`,
				},
				resource.Attribute{
					Name:        "adom",
					Description: `ADOM name. default is 'root'. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Adom revision name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `Who created this adom revision.`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `lock. 0 means unlock and 1 means locked.`,
				},
				resource.Attribute{
					Name:        "adom",
					Description: `ADOM name.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Adom revision name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `Who created this adom revision.`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `lock. 0 means unlock and 1 means locked.`,
				},
				resource.Attribute{
					Name:        "adom",
					Description: `ADOM name.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_fmg_system_admin",
			Category:         "FortiManager Resources",
			ShortDescription: `Provides a resource to configure system admin setting for FortiManager.`,
			Description:      ``,
			Keywords: []string{
				"fortimanager",
				"fmg",
				"system",
				"admin",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "http_port",
					Description: `Http port.`,
				},
				resource.Attribute{
					Name:        "https_port",
					Description: `Https port.`,
				},
				resource.Attribute{
					Name:        "idle_timeout",
					Description: `Idle Timeout(1-480 minute). ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "http_port",
					Description: `Http port.`,
				},
				resource.Attribute{
					Name:        "https_port",
					Description: `Https port.`,
				},
				resource.Attribute{
					Name:        "idle_timeout",
					Description: `Idle Timeout(1-480 minute).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "http_port",
					Description: `Http port.`,
				},
				resource.Attribute{
					Name:        "https_port",
					Description: `Https port.`,
				},
				resource.Attribute{
					Name:        "idle_timeout",
					Description: `Idle Timeout(1-480 minute).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_fmg_system_admin_profiles",
			Category:         "FortiManager Resources",
			ShortDescription: `Provides a resource to configure system admin profiles for FortiManager.`,
			Description:      ``,
			Keywords: []string{
				"fortimanager",
				"fmg",
				"system",
				"admin",
				"profiles",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "profileid",
					Description: `(Required) Profile name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description.`,
				},
				resource.Attribute{
					Name:        "system_setting",
					Description: `System Setting.`,
				},
				resource.Attribute{
					Name:        "adom_switch",
					Description: `Administrator Domain.`,
				},
				resource.Attribute{
					Name:        "deploy_management",
					Description: `Install to devices.`,
				},
				resource.Attribute{
					Name:        "import_policy_packages",
					Description: `Import Policy Package.`,
				},
				resource.Attribute{
					Name:        "intf_mapping",
					Description: `Interface Mapping.`,
				},
				resource.Attribute{
					Name:        "device_ap",
					Description: `Manage AP.`,
				},
				resource.Attribute{
					Name:        "device_forticlient",
					Description: `Manage FortiClient.`,
				},
				resource.Attribute{
					Name:        "device_fortiswitch",
					Description: `Manage FortiSwitch.`,
				},
				resource.Attribute{
					Name:        "vpn_manager",
					Description: `VPN Manager.`,
				},
				resource.Attribute{
					Name:        "log_viewer",
					Description: `Log Viewer.`,
				},
				resource.Attribute{
					Name:        "fortiguard_center",
					Description: `FortiGuard Center.`,
				},
				resource.Attribute{
					Name:        "fortiguard_center_advanced",
					Description: `FortiGuard Center Advanced.`,
				},
				resource.Attribute{
					Name:        "fortiguard_center_firmware_managerment",
					Description: `FortiGuard Center Firmware Managerment.`,
				},
				resource.Attribute{
					Name:        "fortiguard_center_licensing",
					Description: `FortiGuard Center Licensing.`,
				},
				resource.Attribute{
					Name:        "device_manager",
					Description: `Device Manager.`,
				},
				resource.Attribute{
					Name:        "device_operation",
					Description: `Device add/delete/edit.`,
				},
				resource.Attribute{
					Name:        "config_retrieve",
					Description: `Configuration Retrieve.`,
				},
				resource.Attribute{
					Name:        "config_revert",
					Description: `Revert Configuration from Revision History.`,
				},
				resource.Attribute{
					Name:        "device_revision_deletion",
					Description: `Delete device revision.`,
				},
				resource.Attribute{
					Name:        "terminal_access",
					Description: `Terminal access.`,
				},
				resource.Attribute{
					Name:        "device_config",
					Description: `Manage device configurations.`,
				},
				resource.Attribute{
					Name:        "device_profile",
					Description: `Device profile permission.`,
				},
				resource.Attribute{
					Name:        "device_wan_link_load_balance",
					Description: `Manage WAN link load balance.`,
				},
				resource.Attribute{
					Name:        "policy_objects",
					Description: `Policy objects permission.`,
				},
				resource.Attribute{
					Name:        "global_policy_packages",
					Description: `Global policy packages.`,
				},
				resource.Attribute{
					Name:        "assignment",
					Description: `Assignment Permission.`,
				},
				resource.Attribute{
					Name:        "adom_policy_packages",
					Description: `Adom policy packages.`,
				},
				resource.Attribute{
					Name:        "consistency_check",
					Description: `Consistency check.`,
				},
				resource.Attribute{
					Name:        "set_install_targets",
					Description: `Edit installation targets. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "profileid",
					Description: `Profile name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description.`,
				},
				resource.Attribute{
					Name:        "system_setting",
					Description: `System Setting.`,
				},
				resource.Attribute{
					Name:        "adom_switch",
					Description: `Administrator Domain.`,
				},
				resource.Attribute{
					Name:        "deploy_management",
					Description: `Install to devices.`,
				},
				resource.Attribute{
					Name:        "import_policy_packages",
					Description: `Import Policy Package.`,
				},
				resource.Attribute{
					Name:        "intf_mapping",
					Description: `Interface Mapping.`,
				},
				resource.Attribute{
					Name:        "device_ap",
					Description: `Manage AP.`,
				},
				resource.Attribute{
					Name:        "device_forticlient",
					Description: `Manage FortiClient.`,
				},
				resource.Attribute{
					Name:        "device_fortiswitch",
					Description: `Manage FortiSwitch.`,
				},
				resource.Attribute{
					Name:        "vpn_manager",
					Description: `VPN Manager.`,
				},
				resource.Attribute{
					Name:        "log_viewer",
					Description: `Log Viewer.`,
				},
				resource.Attribute{
					Name:        "fortiguard_center",
					Description: `FortiGuard Center.`,
				},
				resource.Attribute{
					Name:        "fortiguard_center_advanced",
					Description: `FortiGuard Center Advanced.`,
				},
				resource.Attribute{
					Name:        "fortiguard_center_firmware_managerment",
					Description: `FortiGuard Center Firmware Managerment.`,
				},
				resource.Attribute{
					Name:        "fortiguard_center_licensing",
					Description: `FortiGuard Center Licensing.`,
				},
				resource.Attribute{
					Name:        "device_manager",
					Description: `Device Manager.`,
				},
				resource.Attribute{
					Name:        "device_operation",
					Description: `Device add/delete/edit.`,
				},
				resource.Attribute{
					Name:        "config_retrieve",
					Description: `Configuration Retrieve.`,
				},
				resource.Attribute{
					Name:        "config_revert",
					Description: `Revert Configuration from Revision History.`,
				},
				resource.Attribute{
					Name:        "device_revision_deletion",
					Description: `Delete device revision.`,
				},
				resource.Attribute{
					Name:        "terminal_access",
					Description: `Terminal access.`,
				},
				resource.Attribute{
					Name:        "device_config",
					Description: `Manage device configurations.`,
				},
				resource.Attribute{
					Name:        "device_profile",
					Description: `Device profile permission.`,
				},
				resource.Attribute{
					Name:        "device_wan_link_load_balance",
					Description: `Manage WAN link load balance.`,
				},
				resource.Attribute{
					Name:        "policy_objects",
					Description: `Policy objects permission.`,
				},
				resource.Attribute{
					Name:        "global_policy_packages",
					Description: `Global policy packages.`,
				},
				resource.Attribute{
					Name:        "assignment",
					Description: `Assignment Permission.`,
				},
				resource.Attribute{
					Name:        "adom_policy_packages",
					Description: `Adom policy packages.`,
				},
				resource.Attribute{
					Name:        "consistency_check",
					Description: `Consistency check.`,
				},
				resource.Attribute{
					Name:        "set_install_targets",
					Description: `Edit installation targets.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "profileid",
					Description: `Profile name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description.`,
				},
				resource.Attribute{
					Name:        "system_setting",
					Description: `System Setting.`,
				},
				resource.Attribute{
					Name:        "adom_switch",
					Description: `Administrator Domain.`,
				},
				resource.Attribute{
					Name:        "deploy_management",
					Description: `Install to devices.`,
				},
				resource.Attribute{
					Name:        "import_policy_packages",
					Description: `Import Policy Package.`,
				},
				resource.Attribute{
					Name:        "intf_mapping",
					Description: `Interface Mapping.`,
				},
				resource.Attribute{
					Name:        "device_ap",
					Description: `Manage AP.`,
				},
				resource.Attribute{
					Name:        "device_forticlient",
					Description: `Manage FortiClient.`,
				},
				resource.Attribute{
					Name:        "device_fortiswitch",
					Description: `Manage FortiSwitch.`,
				},
				resource.Attribute{
					Name:        "vpn_manager",
					Description: `VPN Manager.`,
				},
				resource.Attribute{
					Name:        "log_viewer",
					Description: `Log Viewer.`,
				},
				resource.Attribute{
					Name:        "fortiguard_center",
					Description: `FortiGuard Center.`,
				},
				resource.Attribute{
					Name:        "fortiguard_center_advanced",
					Description: `FortiGuard Center Advanced.`,
				},
				resource.Attribute{
					Name:        "fortiguard_center_firmware_managerment",
					Description: `FortiGuard Center Firmware Managerment.`,
				},
				resource.Attribute{
					Name:        "fortiguard_center_licensing",
					Description: `FortiGuard Center Licensing.`,
				},
				resource.Attribute{
					Name:        "device_manager",
					Description: `Device Manager.`,
				},
				resource.Attribute{
					Name:        "device_operation",
					Description: `Device add/delete/edit.`,
				},
				resource.Attribute{
					Name:        "config_retrieve",
					Description: `Configuration Retrieve.`,
				},
				resource.Attribute{
					Name:        "config_revert",
					Description: `Revert Configuration from Revision History.`,
				},
				resource.Attribute{
					Name:        "device_revision_deletion",
					Description: `Delete device revision.`,
				},
				resource.Attribute{
					Name:        "terminal_access",
					Description: `Terminal access.`,
				},
				resource.Attribute{
					Name:        "device_config",
					Description: `Manage device configurations.`,
				},
				resource.Attribute{
					Name:        "device_profile",
					Description: `Device profile permission.`,
				},
				resource.Attribute{
					Name:        "device_wan_link_load_balance",
					Description: `Manage WAN link load balance.`,
				},
				resource.Attribute{
					Name:        "policy_objects",
					Description: `Policy objects permission.`,
				},
				resource.Attribute{
					Name:        "global_policy_packages",
					Description: `Global policy packages.`,
				},
				resource.Attribute{
					Name:        "assignment",
					Description: `Assignment Permission.`,
				},
				resource.Attribute{
					Name:        "adom_policy_packages",
					Description: `Adom policy packages.`,
				},
				resource.Attribute{
					Name:        "consistency_check",
					Description: `Consistency check.`,
				},
				resource.Attribute{
					Name:        "set_install_targets",
					Description: `Edit installation targets.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_fmg_system_admin_user",
			Category:         "FortiManager Resources",
			ShortDescription: `Provides a resource to configure system admin user for FortiManager.`,
			Description:      ``,
			Keywords: []string{
				"fortimanager",
				"fmg",
				"system",
				"admin",
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "userid",
					Description: `(Required) User name.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Password.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description.`,
				},
				resource.Attribute{
					Name:        "user_type",
					Description: `User type, Enum: ["local", "radius", "ldap", "tacacs-plus", "pki-auth", "group"]`,
				},
				resource.Attribute{
					Name:        "profileid",
					Description: `Profile id.`,
				},
				resource.Attribute{
					Name:        "rpc_permit",
					Description: `Rpc permit, Enum: ["read-write", "none", "read"]`,
				},
				resource.Attribute{
					Name:        "trusthost1",
					Description: `Admin user trusted host IP, default 0.0.0.0 0.0.0.0 for all.`,
				},
				resource.Attribute{
					Name:        "trusthost2",
					Description: `Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.`,
				},
				resource.Attribute{
					Name:        "trusthost3",
					Description: `Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "userid",
					Description: `User name.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Password.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description.`,
				},
				resource.Attribute{
					Name:        "user_type",
					Description: `User type, Enum: ["local", "radius", "ldap", "tacacs-plus", "pki-auth", "group"]`,
				},
				resource.Attribute{
					Name:        "profileid",
					Description: `Profile id.`,
				},
				resource.Attribute{
					Name:        "rpc_permit",
					Description: `Rpc permit, Enum: ["read-write", "none", "read"]`,
				},
				resource.Attribute{
					Name:        "trusthost1",
					Description: `Admin user trusted host IP, default 0.0.0.0 0.0.0.0 for all.`,
				},
				resource.Attribute{
					Name:        "trusthost2",
					Description: `Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.`,
				},
				resource.Attribute{
					Name:        "trusthost3",
					Description: `Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "userid",
					Description: `User name.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Password.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description.`,
				},
				resource.Attribute{
					Name:        "user_type",
					Description: `User type, Enum: ["local", "radius", "ldap", "tacacs-plus", "pki-auth", "group"]`,
				},
				resource.Attribute{
					Name:        "profileid",
					Description: `Profile id.`,
				},
				resource.Attribute{
					Name:        "rpc_permit",
					Description: `Rpc permit, Enum: ["read-write", "none", "read"]`,
				},
				resource.Attribute{
					Name:        "trusthost1",
					Description: `Admin user trusted host IP, default 0.0.0.0 0.0.0.0 for all.`,
				},
				resource.Attribute{
					Name:        "trusthost2",
					Description: `Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.`,
				},
				resource.Attribute{
					Name:        "trusthost3",
					Description: `Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_fmg_system_adom",
			Category:         "FortiManager Resources",
			ShortDescription: `Provides a resource to configure system adom for FortiManager.`,
			Description:      ``,
			Keywords: []string{
				"fortimanager",
				"fmg",
				"system",
				"adom",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Administrative Domain name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Domain type, Enum: ["FortiGate", "FortiCarrier], default is "FortiCarrier".`,
				},
				resource.Attribute{
					Name:        "central_management_vpn",
					Description: `True or False.`,
				},
				resource.Attribute{
					Name:        "central_management_fortiap",
					Description: `True or False.`,
				},
				resource.Attribute{
					Name:        "central_management_sdwan",
					Description: `True or False.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Adom mode: Normal or Backup.`,
				},
				resource.Attribute{
					Name:        "perform_policy_check_before_every_install",
					Description: `True or False.`,
				},
				resource.Attribute{
					Name:        "action_when_conflicts_occur_during_policy_check",
					Description: `True or False.`,
				},
				resource.Attribute{
					Name:        "auto_push_policy_packages_when_device_back_online",
					Description: `True or False.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Adom status. 0 means off and 1 means on. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Administrative Domain name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Domain type.`,
				},
				resource.Attribute{
					Name:        "central_management_vpn",
					Description: `True or False.`,
				},
				resource.Attribute{
					Name:        "central_management_fortiap",
					Description: `True or False.`,
				},
				resource.Attribute{
					Name:        "central_management_sdwan",
					Description: `True or False.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Adom mode: Normal or Backup.`,
				},
				resource.Attribute{
					Name:        "perform_policy_check_before_every_install",
					Description: `True or False.`,
				},
				resource.Attribute{
					Name:        "action_when_conflicts_occur_during_policy_check",
					Description: `True or False.`,
				},
				resource.Attribute{
					Name:        "auto_push_policy_packages_when_device_back_online",
					Description: `True or False.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Adom status.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Administrative Domain name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Domain type.`,
				},
				resource.Attribute{
					Name:        "central_management_vpn",
					Description: `True or False.`,
				},
				resource.Attribute{
					Name:        "central_management_fortiap",
					Description: `True or False.`,
				},
				resource.Attribute{
					Name:        "central_management_sdwan",
					Description: `True or False.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Adom mode: Normal or Backup.`,
				},
				resource.Attribute{
					Name:        "perform_policy_check_before_every_install",
					Description: `True or False.`,
				},
				resource.Attribute{
					Name:        "action_when_conflicts_occur_during_policy_check",
					Description: `True or False.`,
				},
				resource.Attribute{
					Name:        "auto_push_policy_packages_when_device_back_online",
					Description: `True or False.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Adom status.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_fmg_system_dns",
			Category:         "FortiManager Resources",
			ShortDescription: `Provides a resource to configure system dns setting for FortiManager.`,
			Description:      ``,
			Keywords: []string{
				"fortimanager",
				"fmg",
				"system",
				"dns",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "primary",
					Description: `Primary DNS IP.`,
				},
				resource.Attribute{
					Name:        "secondary",
					Description: `Secondary DNS IP. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `Primary DNS IP.`,
				},
				resource.Attribute{
					Name:        "secondary",
					Description: `Secondary DNS IP.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `Primary DNS IP.`,
				},
				resource.Attribute{
					Name:        "secondary",
					Description: `Secondary DNS IP.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_fmg_system_global",
			Category:         "FortiManager Resources",
			ShortDescription: `Provides a resource to configure system global setting for FortiManager.`,
			Description:      ``,
			Keywords: []string{
				"fortimanager",
				"fmg",
				"system",
				"global",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "fortianalyzer_status",
					Description: `Enable/Disable FortiAnalyzer feature.`,
				},
				resource.Attribute{
					Name:        "adom_status",
					Description: `Enable/Disable ADOM.`,
				},
				resource.Attribute{
					Name:        "adom_mode",
					Description: `Adom Mode.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `Hostname.`,
				},
				resource.Attribute{
					Name:        "timezone",
					Description: `TimeZone. 00 - (GMT-12:00) Eniwetak, Kwajalein.01 - (GMT-11:00) Midway Island, Samoa.02 - (GMT-10:00) Hawaii.03 - (GMT-9:00) Alaska.04 - (GMT-8:00) Pacific Time (US & Canada).05 - (GMT-7:00) Arizona.06 - (GMT-7:00) Mountain Time (US & Canada).07 - (GMT-6:00) Central America.08 - (GMT-6:00) Central Time (US & Canada).09 - (GMT-6:00) Mexico City.10 - (GMT-6:00) Saskatchewan.11 - (GMT-5:00) Bogota, Lima, Quito.12 - (GMT-5:00) Eastern Time (US & Canada).13 - (GMT-5:00) Indiana (East).14 - (GMT-4:00) Atlantic Time (Canada).15 - (GMT-4:00) La Paz.16 - (GMT-4:00) Santiago.17 - (GMT-3:30) Newfoundland.18 - (GMT-3:00) Brasilia.19 - (GMT-3:00) Buenos Aires, Georgetown.20 - (GMT-3:00) Nuuk (Greenland).21 - (GMT-2:00) Mid-Atlantic.22 - (GMT-1:00) Azores.23 - (GMT-1:00) Cape Verde Is.24 - (GMT) Monrovia.25 - (GMT) Greenwich Mean Time:Dublin, Edinburgh, Lisbon, London.26 - (GMT+1:00) Amsterdam, Berlin, Bern, Rome, Stockholm, Vienna.27 - (GMT+1:00) Belgrade, Bratislava, Budapest, Ljubljana, Prague.28 - (GMT+1:00) Brussels, Copenhagen, Madrid, Paris.29 - (GMT+1:00) Sarajevo, Skopje, Warsaw, Zagreb.30 - (GMT+1:00) West Central Africa.31 - (GMT+2:00) Athens, Sofia, Vilnius.32 - (GMT+2:00) Bucharest.33 - (GMT+2:00) Cairo.34 - (GMT+2:00) Harare, Pretoria.35 - (GMT+2:00) Helsinki, Riga,Tallinn.36 - (GMT+2:00) Jerusalem.37 - (GMT+3:00) Baghdad.38 - (GMT+3:00) Kuwait, Riyadh.39 - (GMT+3:00) St.Petersburg, Volgograd.40 - (GMT+3:00) Nairobi.41 - (GMT+3:30) Tehran.42 - (GMT+4:00) Abu Dhabi, Muscat.43 - (GMT+4:00) Baku.44 - (GMT+4:30) Kabul.45 - (GMT+5:00) Ekaterinburg.46 - (GMT+5:00) Islamabad, Karachi,Tashkent.47 - (GMT+5:30) Calcutta, Chennai, Mumbai, New Delhi.48 - (GMT+5:45) Kathmandu.49 - (GMT+6:00) Almaty, Novosibirsk.50 - (GMT+6:00) Astana, Dhaka.51 - (GMT+6:00) Sri Jayawardenapura.52 - (GMT+6:30) Rangoon.53 - (GMT+7:00) Bangkok, Hanoi, Jakarta.54 - (GMT+7:00) Krasnoyarsk.55 - (GMT+8:00) Beijing,ChongQing, HongKong,Urumqi.56 - (GMT+8:00) Irkutsk, Ulaanbaatar.57 - (GMT+8:00) Kuala Lumpur, Singapore.58 - (GMT+8:00) Perth.59 - (GMT+8:00) Taipei.60 - (GMT+9:00) Osaka, Sapporo, Tokyo, Seoul.61 - (GMT+9:00) Yakutsk.62 - (GMT+9:30) Adelaide.63 - (GMT+9:30) Darwin.64 - (GMT+10:00) Brisbane.65 - (GMT+10:00) Canberra, Melbourne, Sydney.66 - (GMT+10:00) Guam, Port Moresby.67 - (GMT+10:00) Hobart.68 - (GMT+10:00) Vladivostok.69 - (GMT+11:00) Magadan.70 - (GMT+11:00) Solomon Is., New Caledonia.71 - (GMT+12:00) Auckland, Wellington.72 - (GMT+12:00) Fiji, Kamchatka, Marshall Is.73 - (GMT+13:00) Nuku'alofa.74 - (GMT-4:30) Caracas.75 - (GMT+1:00) Namibia.76 - (GMT-5:00) Brazil-Acre.77 - (GMT-4:00) Brazil-West.78 - (GMT-3:00) Brazil-East.79 - (GMT-2:00) Brazil-DeNoronha.80 - (GMT+14:00) Kiritimati.81 - (GMT-7:00) Baja California Sur, Chihuahua.82 - (GMT+12:45) Chatham Islands.83 - (GMT+3:00) Minsk.84 - (GMT+13:00) Samoa.85 - (GMT+3:00) Istanbul.86 - (GMT-4:00) Paraguay.87 - (GMT) Casablanca.88 - (GMT+3:00) Moscow.89 - (GMT) Greenwich Mean Time. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "fortianalyzer_status",
					Description: `Enable/Disable FortiAnalyzer feature.`,
				},
				resource.Attribute{
					Name:        "adom_status",
					Description: `Enable/Disable ADOM.`,
				},
				resource.Attribute{
					Name:        "adom_mode",
					Description: `Adom Mode.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `Hostname.`,
				},
				resource.Attribute{
					Name:        "timezone",
					Description: `TimeZone.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "fortianalyzer_status",
					Description: `Enable/Disable FortiAnalyzer feature.`,
				},
				resource.Attribute{
					Name:        "adom_status",
					Description: `Enable/Disable ADOM.`,
				},
				resource.Attribute{
					Name:        "adom_mode",
					Description: `Adom Mode.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `Hostname.`,
				},
				resource.Attribute{
					Name:        "timezone",
					Description: `TimeZone.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_fmg_system_license_forticare",
			Category:         "FortiManager Resources",
			ShortDescription: `Provides a resource to upload FortiCare registration code to FortiGate through FortiManager.`,
			Description:      ``,
			Keywords: []string{
				"fortimanager",
				"fmg",
				"system",
				"license",
				"forticare",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "target",
					Description: `(Required) Target name, which is managed by FortiManager.`,
				},
				resource.Attribute{
					Name:        "registration_code",
					Description: `(Required) Registration code.`,
				},
				resource.Attribute{
					Name:        "adom",
					Description: `ADOM that the target device belongs to. default is 'root'. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `Target name, which is managed by FortiManager.`,
				},
				resource.Attribute{
					Name:        "registration_code",
					Description: `Registration code.`,
				},
				resource.Attribute{
					Name:        "adom",
					Description: `ADOM name.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `Target name, which is managed by FortiManager.`,
				},
				resource.Attribute{
					Name:        "registration_code",
					Description: `Registration code.`,
				},
				resource.Attribute{
					Name:        "adom",
					Description: `ADOM name.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_fmg_system_license_vm",
			Category:         "FortiManager Resources",
			ShortDescription: `Provides a resource to upload VM license to FortiGate through FortiManager.`,
			Description:      ``,
			Keywords: []string{
				"fortimanager",
				"fmg",
				"system",
				"license",
				"vm",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "target",
					Description: `(Required) Target name, which is managed by FortiManager.`,
				},
				resource.Attribute{
					Name:        "file_content",
					Description: `(Required) The license file, it needs to be base64 encoded.`,
				},
				resource.Attribute{
					Name:        "adom",
					Description: `ADOM that the target device belongs to. default is 'root'. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `Target name, which is managed by FortiManager.`,
				},
				resource.Attribute{
					Name:        "file_content",
					Description: `The license file content, it needs to be base64 encoded.`,
				},
				resource.Attribute{
					Name:        "adom",
					Description: `ADOM name.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `Target name, which is managed by FortiManager.`,
				},
				resource.Attribute{
					Name:        "file_content",
					Description: `The license file content, it needs to be base64 encoded.`,
				},
				resource.Attribute{
					Name:        "adom",
					Description: `ADOM name.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_fmg_system_network_interface",
			Category:         "FortiManager Resources",
			ShortDescription: `Provides a resource to configure system network interface for FortiManager.`,
			Description:      ``,
			Keywords: []string{
				"fortimanager",
				"fmg",
				"system",
				"network",
				"interface",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Interface port.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Interface Ipaddress.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Interface status, Enum: ["down", "up"]`,
				},
				resource.Attribute{
					Name:        "allow_access",
					Description: `Allow managerment access to interface, Enum: ["ping", "https", "ssh", "snmp", "telnet", "http", "web"]`,
				},
				resource.Attribute{
					Name:        "sevice_access",
					Description: `Allow service access to interface, Enum: ["fgtupdates", "webfilter"] ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Interface port.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Interface Ipaddress.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Interface status.`,
				},
				resource.Attribute{
					Name:        "allow_access",
					Description: `Allow managerment access to interface.`,
				},
				resource.Attribute{
					Name:        "sevice_access",
					Description: `Allow service access to interface.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Interface port.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Interface Ipaddress.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Interface status.`,
				},
				resource.Attribute{
					Name:        "allow_access",
					Description: `Allow managerment access to interface.`,
				},
				resource.Attribute{
					Name:        "sevice_access",
					Description: `Allow service access to interface.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_fmg_system_network_route",
			Category:         "FortiManager Resources",
			ShortDescription: `Provides a resource to configure system network route for FortiManager.`,
			Description:      ``,
			Keywords: []string{
				"fortimanager",
				"fmg",
				"system",
				"network",
				"route",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "route_id",
					Description: `(Required) Route id.`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `(Required) Destination Ip/Mask.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `(Required) Gateway Ip.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `(Required) Gateway out interface. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "route_id",
					Description: `Route id.`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `Destination Ip/Mask.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `Gateway Ip.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `Gateway out interface.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "route_id",
					Description: `Route id.`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `Destination Ip/Mask.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `Gateway Ip.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `Gateway out interface.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_fmg_system_ntp",
			Category:         "FortiManager Resources",
			ShortDescription: `Provides a resource to configure system ntp setting for FortiManager.`,
			Description:      ``,
			Keywords: []string{
				"fortimanager",
				"fmg",
				"system",
				"ntp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "server",
					Description: `IP address/hostname of NTP Server.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable NTP.`,
				},
				resource.Attribute{
					Name:        "sync_interval",
					Description: `NTP sync interval (minute). ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `IP address/hostname of NTP Server.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable NTP.`,
				},
				resource.Attribute{
					Name:        "sync_interval",
					Description: `NTP sync interval.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `IP address/hostname of NTP Server.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable NTP.`,
				},
				resource.Attribute{
					Name:        "sync_interval",
					Description: `NTP sync interval.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_fmg_system_syslogserver",
			Category:         "FortiManager Resources",
			ShortDescription: `Provides a resource to configure system syslog server for FortiManager.`,
			Description:      ``,
			Keywords: []string{
				"fortimanager",
				"fmg",
				"system",
				"syslogserver",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Syslog server name.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Ipaddress of the syslog server.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port of the syslog server. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Syslog server name.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Ipaddress of the syslog server.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port of the syslog server.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Syslog server name.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Ipaddress of the syslog server.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port of the syslog server.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_firewall_object_address",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to configure firewall addresses used in firewall policies of FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"firewall",
				"object",
				"address",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Address name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type of address(Support ipmask, iprange, fqdn and geography).`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `IP address and subnet mask of address.`,
				},
				resource.Attribute{
					Name:        "start_ip",
					Description: `First IP address (inclusive) in the range for the address.`,
				},
				resource.Attribute{
					Name:        "end_ip",
					Description: `Final IP address (inclusive) in the range for the address.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `Fully Qualified Domain Name address.`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `IP addresses associated to a specific country.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
				resource.Attribute{
					Name:        "associated_interface",
					Description: `Network interface associated with address.`,
				},
				resource.Attribute{
					Name:        "show_in_address_list",
					Description: `Enable/disable address visibility in the GUI. default is "enable".`,
				},
				resource.Attribute{
					Name:        "static_route_configure",
					Description: `Enable/disable use of this address in the static route configuration. default is "disable". ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the address item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Address name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of address(Support ipmask, iprange, fqdn and geography).`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `IP address and subnet mask of address.`,
				},
				resource.Attribute{
					Name:        "start_ip",
					Description: `First IP address (inclusive) in the range for the address.`,
				},
				resource.Attribute{
					Name:        "end_ip",
					Description: `Final IP address (inclusive) in the range for the address.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `Fully Qualified Domain Name address.`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `IP addresses associated to a specific country.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
				resource.Attribute{
					Name:        "associated_interface",
					Description: `Network interface associated with address.`,
				},
				resource.Attribute{
					Name:        "show_in_address_list",
					Description: `Enable/disable address visibility in the GUI. default is "enable".`,
				},
				resource.Attribute{
					Name:        "static_route_configure",
					Description: `Enable/disable use of this address in the static route configuration. default is "disable".`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the address item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Address name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of address(Support ipmask, iprange, fqdn and geography).`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `IP address and subnet mask of address.`,
				},
				resource.Attribute{
					Name:        "start_ip",
					Description: `First IP address (inclusive) in the range for the address.`,
				},
				resource.Attribute{
					Name:        "end_ip",
					Description: `Final IP address (inclusive) in the range for the address.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `Fully Qualified Domain Name address.`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `IP addresses associated to a specific country.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
				resource.Attribute{
					Name:        "associated_interface",
					Description: `Network interface associated with address.`,
				},
				resource.Attribute{
					Name:        "show_in_address_list",
					Description: `Enable/disable address visibility in the GUI. default is "enable".`,
				},
				resource.Attribute{
					Name:        "static_route_configure",
					Description: `Enable/disable use of this address in the static route configuration. default is "disable".`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_firewall_object_addressgroup",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to configure firewall address group used in firewall policies of FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"firewall",
				"object",
				"addressgroup",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Address group name.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Required) Address objects contained within the group.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the firewall address group item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Address group name.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `Address objects contained within the group.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the firewall address group item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Address group name.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `Address objects contained within the group.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_firewall_object_ippool",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to configure IPv4 IP address pools of FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"firewall",
				"object",
				"ippool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) IP pool name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) IP pool type(Support overload and one-to-one).`,
				},
				resource.Attribute{
					Name:        "startip",
					Description: `(Required) First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).`,
				},
				resource.Attribute{
					Name:        "endip",
					Description: `(Required) Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).`,
				},
				resource.Attribute{
					Name:        "arp_reply",
					Description: `Enable/disable replying to ARP requests when an IP Pool is added to a policy.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the IP pool item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `IP pool name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `IP pool type(Support overload and one-to-one).`,
				},
				resource.Attribute{
					Name:        "startip",
					Description: `First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).`,
				},
				resource.Attribute{
					Name:        "endip",
					Description: `Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).`,
				},
				resource.Attribute{
					Name:        "arp_reply",
					Description: `Enable/disable replying to ARP requests when an IP Pool is added to a policy.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the IP pool item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `IP pool name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `IP pool type(Support overload and one-to-one).`,
				},
				resource.Attribute{
					Name:        "startip",
					Description: `First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).`,
				},
				resource.Attribute{
					Name:        "endip",
					Description: `Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).`,
				},
				resource.Attribute{
					Name:        "arp_reply",
					Description: `Enable/disable replying to ARP requests when an IP Pool is added to a policy.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_firewall_object_service",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to configure firewall service of FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"firewall",
				"object",
				"service",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Number of minutes before an idle administrator session time out.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `(Required) Service category.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol type based on IANA numbers.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `Fully qualified domain name.`,
				},
				resource.Attribute{
					Name:        "iprange",
					Description: `Start and end of the IP range associated with service.`,
				},
				resource.Attribute{
					Name:        "tcp_portrange",
					Description: `Multiple TCP port ranges.`,
				},
				resource.Attribute{
					Name:        "udp_portrange",
					Description: `Multiple UDP port ranges.`,
				},
				resource.Attribute{
					Name:        "sctp_portrange",
					Description: `Multiple SCTP port ranges.`,
				},
				resource.Attribute{
					Name:        "icmptype",
					Description: `ICMP type.`,
				},
				resource.Attribute{
					Name:        "icmpcode",
					Description: `ICMP code.`,
				},
				resource.Attribute{
					Name:        "protocol_number",
					Description: `IP protocol number.`,
				},
				resource.Attribute{
					Name:        "session_ttl",
					Description: `Custom tcp session TTL.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the firewall service item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Number of minutes before an idle administrator session time out.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `Service category.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol type based on IANA numbers.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `Fully qualified domain name.`,
				},
				resource.Attribute{
					Name:        "iprange",
					Description: `Start and end of the IP range associated with service.`,
				},
				resource.Attribute{
					Name:        "tcp_portrange",
					Description: `Multiple TCP port ranges.`,
				},
				resource.Attribute{
					Name:        "udp_portrange",
					Description: `Multiple UDP port ranges.`,
				},
				resource.Attribute{
					Name:        "sctp_portrange",
					Description: `Multiple SCTP port ranges.`,
				},
				resource.Attribute{
					Name:        "icmptype",
					Description: `ICMP type.`,
				},
				resource.Attribute{
					Name:        "icmpcode",
					Description: `ICMP code.`,
				},
				resource.Attribute{
					Name:        "protocol_number",
					Description: `IP protocol number.`,
				},
				resource.Attribute{
					Name:        "session_ttl",
					Description: `Custom tcp session TTL.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the firewall service item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Number of minutes before an idle administrator session time out.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `Service category.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol type based on IANA numbers.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `Fully qualified domain name.`,
				},
				resource.Attribute{
					Name:        "iprange",
					Description: `Start and end of the IP range associated with service.`,
				},
				resource.Attribute{
					Name:        "tcp_portrange",
					Description: `Multiple TCP port ranges.`,
				},
				resource.Attribute{
					Name:        "udp_portrange",
					Description: `Multiple UDP port ranges.`,
				},
				resource.Attribute{
					Name:        "sctp_portrange",
					Description: `Multiple SCTP port ranges.`,
				},
				resource.Attribute{
					Name:        "icmptype",
					Description: `ICMP type.`,
				},
				resource.Attribute{
					Name:        "icmpcode",
					Description: `ICMP code.`,
				},
				resource.Attribute{
					Name:        "protocol_number",
					Description: `IP protocol number.`,
				},
				resource.Attribute{
					Name:        "session_ttl",
					Description: `Custom tcp session TTL.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_firewall_object_servicecategory",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to configure firewall service categories of FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"firewall",
				"object",
				"servicecategory",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Category name.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the firewall service item.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the firewall service item.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_firewall_object_servicegroup",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to configure firewall service group of FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"firewall",
				"object",
				"servicegroup",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Service group name.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Required) Service objects contained within the group.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the firewall service group item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Service group name.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `Service objects contained within the group.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the firewall service group item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Service group name.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `Service objects contained within the group.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_firewall_object_vip",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to configure firewall virtual IPs (VIPs) of FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"firewall",
				"object",
				"vip",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Virtual IP name.`,
				},
				resource.Attribute{
					Name:        "extip",
					Description: `(Required) IP address or address range on the external interface that you want to map to an address or address range on the destination network.`,
				},
				resource.Attribute{
					Name:        "mappedip",
					Description: `(Required) IP address or address range on the destination network to which the external IP address is mapped.`,
				},
				resource.Attribute{
					Name:        "extintf",
					Description: `Interface connected to the source network that receives the packets that will be forwarded to the destination network.`,
				},
				resource.Attribute{
					Name:        "portforward",
					Description: `Enable/disable port forwarding.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol to use when forwarding packets.`,
				},
				resource.Attribute{
					Name:        "extport",
					Description: `Incoming port number range that you want to map to a port number range on the destination network.`,
				},
				resource.Attribute{
					Name:        "mappedport",
					Description: `Port number range on the destination network to which the external port number range is mapped.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the firewall virtual IPs item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Virtual IP name.`,
				},
				resource.Attribute{
					Name:        "extip",
					Description: `IP address or address range on the external interface that you want to map to an address or address range on the destination network.`,
				},
				resource.Attribute{
					Name:        "mappedip",
					Description: `IP address or address range on the destination network to which the external IP address is mapped.`,
				},
				resource.Attribute{
					Name:        "extintf",
					Description: `Interface connected to the source network that receives the packets that will be forwarded to the destination network.`,
				},
				resource.Attribute{
					Name:        "portforward",
					Description: `Enable/disable port forwarding.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol to use when forwarding packets.`,
				},
				resource.Attribute{
					Name:        "extport",
					Description: `Incoming port number range that you want to map to a port number range on the destination network.`,
				},
				resource.Attribute{
					Name:        "mappedport",
					Description: `Port number range on the destination network to which the external port number range is mapped.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the firewall virtual IPs item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Virtual IP name.`,
				},
				resource.Attribute{
					Name:        "extip",
					Description: `IP address or address range on the external interface that you want to map to an address or address range on the destination network.`,
				},
				resource.Attribute{
					Name:        "mappedip",
					Description: `IP address or address range on the destination network to which the external IP address is mapped.`,
				},
				resource.Attribute{
					Name:        "extintf",
					Description: `Interface connected to the source network that receives the packets that will be forwarded to the destination network.`,
				},
				resource.Attribute{
					Name:        "portforward",
					Description: `Enable/disable port forwarding.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol to use when forwarding packets.`,
				},
				resource.Attribute{
					Name:        "extport",
					Description: `Incoming port number range that you want to map to a port number range on the destination network.`,
				},
				resource.Attribute{
					Name:        "mappedport",
					Description: `Port number range on the destination network to which the external port number range is mapped.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_firewall_object_vipgroup",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to configure virtual IP groups of FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"firewall",
				"object",
				"vipgroup",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) VIP group name.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `Interface name.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Required) Member VIP objects of the group.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the virtual IP groups item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `VIP group name.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `Interface name.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `Member VIP objects of the group.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the virtual IP groups item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `VIP group name.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `Interface name.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `Member VIP objects of the group.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_firewall_security_policy",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to configure firewall policies of FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"firewall",
				"security",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Policy name.`,
				},
				resource.Attribute{
					Name:        "srcintf",
					Description: `(Required) Incoming (ingress) interface.`,
				},
				resource.Attribute{
					Name:        "dstintf",
					Description: `(Required) Outgoing (egress) interface.`,
				},
				resource.Attribute{
					Name:        "srcaddr",
					Description: `(Required) Source address and address group names.`,
				},
				resource.Attribute{
					Name:        "dstaddr",
					Description: `(Required) Destination address and address group names.`,
				},
				resource.Attribute{
					Name:        "internet_service",
					Description: `Enable/disable use of Internet Services for this policy. If enabled, destination address and service are not used.`,
				},
				resource.Attribute{
					Name:        "internet_service_id",
					Description: `Internet Service ID.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) Policy action.`,
				},
				resource.Attribute{
					Name:        "schedule",
					Description: `(Required) Schedule name.`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `(Required) Service and service group names..`,
				},
				resource.Attribute{
					Name:        "utm_status",
					Description: `Enable to add one or more security profiles (AV, IPS, etc.) to the firewall policy.`,
				},
				resource.Attribute{
					Name:        "logtraffic",
					Description: `Enable or disable logging. Log all sessions or security profile sessions.`,
				},
				resource.Attribute{
					Name:        "logtraffic_start",
					Description: `Record logs when a session starts and ends.`,
				},
				resource.Attribute{
					Name:        "capture_packet",
					Description: `Enable/disable capture packets.`,
				},
				resource.Attribute{
					Name:        "ippool",
					Description: `Enable to use IP Pools for source NAT.`,
				},
				resource.Attribute{
					Name:        "poolname",
					Description: `IP Pool names.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `Names of user groups that can authenticate with this policy.`,
				},
				resource.Attribute{
					Name:        "devices",
					Description: `Device type category.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
				resource.Attribute{
					Name:        "av_profile",
					Description: `Name of an existing Antivirus profile.`,
				},
				resource.Attribute{
					Name:        "webfilter_profile",
					Description: `Name of an existing Web filter profile.`,
				},
				resource.Attribute{
					Name:        "dnsfilter_profile",
					Description: `Name of an existing DNS filter profile.`,
				},
				resource.Attribute{
					Name:        "ips_sensor",
					Description: `Name of an existing IPS sensor.`,
				},
				resource.Attribute{
					Name:        "application_list",
					Description: `Name of an existing Application list.`,
				},
				resource.Attribute{
					Name:        "ssl_ssh_profile",
					Description: `Name of an existing SSL SSH profile.`,
				},
				resource.Attribute{
					Name:        "nat",
					Description: `Enable/disable source NAT.`,
				},
				resource.Attribute{
					Name:        "internet_service_src",
					Description: `Enable/disable use of Internet Services in source for this policy. If enabled, source address is not used.`,
				},
				resource.Attribute{
					Name:        "internet_service_src_id",
					Description: `Internet Service source ID.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `Names of individual users that can authenticate with this policy.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable or disable this policy.`,
				},
				resource.Attribute{
					Name:        "profile_protocol_options",
					Description: `Name of an existing Protocol options profile. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the firewall policy item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Policy name.`,
				},
				resource.Attribute{
					Name:        "srcintf",
					Description: `Incoming (ingress) interface.`,
				},
				resource.Attribute{
					Name:        "dstintf",
					Description: `Outgoing (egress) interface.`,
				},
				resource.Attribute{
					Name:        "srcaddr",
					Description: `Source address and address group names.`,
				},
				resource.Attribute{
					Name:        "dstaddr",
					Description: `Destination address and address group names.`,
				},
				resource.Attribute{
					Name:        "internet_service",
					Description: `Enable/disable use of Internet Services for this policy. If enabled, destination address and service are not used.`,
				},
				resource.Attribute{
					Name:        "internet_service_id",
					Description: `Internet Service ID.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Policy action.`,
				},
				resource.Attribute{
					Name:        "schedule",
					Description: `Schedule name.`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `Service and service group names..`,
				},
				resource.Attribute{
					Name:        "utm_status",
					Description: `Enable to add one or more security profiles (AV, IPS, etc.) to the firewall policy.`,
				},
				resource.Attribute{
					Name:        "logtraffic",
					Description: `Enable or disable logging. Log all sessions or security profile sessions.`,
				},
				resource.Attribute{
					Name:        "logtraffic_start",
					Description: `Record logs when a session starts and ends.`,
				},
				resource.Attribute{
					Name:        "capture_packet",
					Description: `Enable/disable capture packets.`,
				},
				resource.Attribute{
					Name:        "ippool",
					Description: `Enable to use IP Pools for source NAT.`,
				},
				resource.Attribute{
					Name:        "poolname",
					Description: `IP Pool names.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `Names of user groups that can authenticate with this policy.`,
				},
				resource.Attribute{
					Name:        "devices",
					Description: `Device type category.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
				resource.Attribute{
					Name:        "av_profile",
					Description: `Name of an existing Antivirus profile.`,
				},
				resource.Attribute{
					Name:        "webfilter_profile",
					Description: `Name of an existing Web filter profile.`,
				},
				resource.Attribute{
					Name:        "dnsfilter_profile",
					Description: `Name of an existing DNS filter profile.`,
				},
				resource.Attribute{
					Name:        "ips_sensor",
					Description: `Name of an existing IPS sensor.`,
				},
				resource.Attribute{
					Name:        "application_list",
					Description: `Name of an existing Application list.`,
				},
				resource.Attribute{
					Name:        "ssl_ssh_profile",
					Description: `Name of an existing SSL SSH profile.`,
				},
				resource.Attribute{
					Name:        "nat",
					Description: `Enable/disable source NAT.`,
				},
				resource.Attribute{
					Name:        "internet_service_src",
					Description: `Enable/disable use of Internet Services in source for this policy. If enabled, source address is not used.`,
				},
				resource.Attribute{
					Name:        "internet_service_src_id",
					Description: `Internet Service source ID.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `Names of individual users that can authenticate with this policy.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable or disable this policy.`,
				},
				resource.Attribute{
					Name:        "profile_protocol_options",
					Description: `Name of an existing Protocol options profile.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the firewall policy item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Policy name.`,
				},
				resource.Attribute{
					Name:        "srcintf",
					Description: `Incoming (ingress) interface.`,
				},
				resource.Attribute{
					Name:        "dstintf",
					Description: `Outgoing (egress) interface.`,
				},
				resource.Attribute{
					Name:        "srcaddr",
					Description: `Source address and address group names.`,
				},
				resource.Attribute{
					Name:        "dstaddr",
					Description: `Destination address and address group names.`,
				},
				resource.Attribute{
					Name:        "internet_service",
					Description: `Enable/disable use of Internet Services for this policy. If enabled, destination address and service are not used.`,
				},
				resource.Attribute{
					Name:        "internet_service_id",
					Description: `Internet Service ID.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Policy action.`,
				},
				resource.Attribute{
					Name:        "schedule",
					Description: `Schedule name.`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `Service and service group names..`,
				},
				resource.Attribute{
					Name:        "utm_status",
					Description: `Enable to add one or more security profiles (AV, IPS, etc.) to the firewall policy.`,
				},
				resource.Attribute{
					Name:        "logtraffic",
					Description: `Enable or disable logging. Log all sessions or security profile sessions.`,
				},
				resource.Attribute{
					Name:        "logtraffic_start",
					Description: `Record logs when a session starts and ends.`,
				},
				resource.Attribute{
					Name:        "capture_packet",
					Description: `Enable/disable capture packets.`,
				},
				resource.Attribute{
					Name:        "ippool",
					Description: `Enable to use IP Pools for source NAT.`,
				},
				resource.Attribute{
					Name:        "poolname",
					Description: `IP Pool names.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `Names of user groups that can authenticate with this policy.`,
				},
				resource.Attribute{
					Name:        "devices",
					Description: `Device type category.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
				resource.Attribute{
					Name:        "av_profile",
					Description: `Name of an existing Antivirus profile.`,
				},
				resource.Attribute{
					Name:        "webfilter_profile",
					Description: `Name of an existing Web filter profile.`,
				},
				resource.Attribute{
					Name:        "dnsfilter_profile",
					Description: `Name of an existing DNS filter profile.`,
				},
				resource.Attribute{
					Name:        "ips_sensor",
					Description: `Name of an existing IPS sensor.`,
				},
				resource.Attribute{
					Name:        "application_list",
					Description: `Name of an existing Application list.`,
				},
				resource.Attribute{
					Name:        "ssl_ssh_profile",
					Description: `Name of an existing SSL SSH profile.`,
				},
				resource.Attribute{
					Name:        "nat",
					Description: `Enable/disable source NAT.`,
				},
				resource.Attribute{
					Name:        "internet_service_src",
					Description: `Enable/disable use of Internet Services in source for this policy. If enabled, source address is not used.`,
				},
				resource.Attribute{
					Name:        "internet_service_src_id",
					Description: `Internet Service source ID.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `Names of individual users that can authenticate with this policy.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable or disable this policy.`,
				},
				resource.Attribute{
					Name:        "profile_protocol_options",
					Description: `Name of an existing Protocol options profile.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_firewall_security_policyseq",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to alter firewall security policy sequence`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"firewall",
				"security",
				"policyseq",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_src_id",
					Description: `(Required) The policy id which you want to alter`,
				},
				resource.Attribute{
					Name:        "policy_dst_id",
					Description: `(Required) The dest policy id which you want to alter`,
				},
				resource.Attribute{
					Name:        "alter_position",
					Description: `(Required) The alter position: should only be "before" or "after" ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "policy_src_id",
					Description: `The policy id which you want to alter`,
				},
				resource.Attribute{
					Name:        "policy_dst_id",
					Description: `The dest policy id which you want to alter`,
				},
				resource.Attribute{
					Name:        "alter_position",
					Description: `The alter position: should only be "before" or "after"`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_src_id",
					Description: `The policy id which you want to alter`,
				},
				resource.Attribute{
					Name:        "policy_dst_id",
					Description: `The dest policy id which you want to alter`,
				},
				resource.Attribute{
					Name:        "alter_position",
					Description: `The alter position: should only be "before" or "after"`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_log_fortianalyzer_setting",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to configure configure logging to FortiAnalyzer log management devices.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"log",
				"fortianalyzer",
				"setting",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `(Required) Enable/disable logging to FortiAnalyzer.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `The remote FortiAnalyzer.`,
				},
				resource.Attribute{
					Name:        "source_ip",
					Description: `Source IPv4 or IPv6 address used to communicate with FortiAnalyzer.`,
				},
				resource.Attribute{
					Name:        "upload_option",
					Description: `Enable/disable logging to hard disk and then uploading to FortiAnalyzer.`,
				},
				resource.Attribute{
					Name:        "reliable",
					Description: `Enable/disable reliable logging to FortiAnalyzer.`,
				},
				resource.Attribute{
					Name:        "hmac_algorithm",
					Description: `FortiAnalyzer IPsec tunnel HMAC algorithm.`,
				},
				resource.Attribute{
					Name:        "enc_algorithm",
					Description: `Enable/disable sending FortiAnalyzer log data with SSL encryption. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable logging to FortiAnalyzer.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `The remote FortiAnalyzer.`,
				},
				resource.Attribute{
					Name:        "source_ip",
					Description: `Source IPv4 or IPv6 address used to communicate with FortiAnalyzer.`,
				},
				resource.Attribute{
					Name:        "upload_option",
					Description: `Enable/disable logging to hard disk and then uploading to FortiAnalyzer.`,
				},
				resource.Attribute{
					Name:        "reliable",
					Description: `Enable/disable reliable logging to FortiAnalyzer.`,
				},
				resource.Attribute{
					Name:        "hmac_algorithm",
					Description: `FortiAnalyzer IPsec tunnel HMAC algorithm.`,
				},
				resource.Attribute{
					Name:        "enc_algorithm",
					Description: `Enable/disable sending FortiAnalyzer log data with SSL encryption.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable logging to FortiAnalyzer.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `The remote FortiAnalyzer.`,
				},
				resource.Attribute{
					Name:        "source_ip",
					Description: `Source IPv4 or IPv6 address used to communicate with FortiAnalyzer.`,
				},
				resource.Attribute{
					Name:        "upload_option",
					Description: `Enable/disable logging to hard disk and then uploading to FortiAnalyzer.`,
				},
				resource.Attribute{
					Name:        "reliable",
					Description: `Enable/disable reliable logging to FortiAnalyzer.`,
				},
				resource.Attribute{
					Name:        "hmac_algorithm",
					Description: `FortiAnalyzer IPsec tunnel HMAC algorithm.`,
				},
				resource.Attribute{
					Name:        "enc_algorithm",
					Description: `Enable/disable sending FortiAnalyzer log data with SSL encryption.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_log_syslog_setting",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to configure logging to remote Syslog logging servers.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"log",
				"syslog",
				"setting",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `(Required) Enable/disable remote syslog logging.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `Address of remote syslog server.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Remote syslog logging over UDP/Reliable TCP.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Server listen port.`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `Remote syslog facility.`,
				},
				resource.Attribute{
					Name:        "source_ip",
					Description: `Source IP address of syslog.`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `Log format. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable remote syslog logging.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `Address of remote syslog server.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Remote syslog logging over UDP/Reliable TCP.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Server listen port.`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `Remote syslog facility.`,
				},
				resource.Attribute{
					Name:        "source_ip",
					Description: `Source IP address of syslog.`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `Log format.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable remote syslog logging.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `Address of remote syslog server.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Remote syslog logging over UDP/Reliable TCP.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Server listen port.`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `Remote syslog facility.`,
				},
				resource.Attribute{
					Name:        "source_ip",
					Description: `Source IP address of syslog.`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `Log format.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_networking_interface_port",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to configure interface settings of FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"networking",
				"interface",
				"port",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) If the interface is physical, the argument is the name of the interface.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Interface type (support physical, vlan, loopback).`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Interface IPv4 address and subnet mask, syntax` + "`" + ` - X.X.X.X X.X.X.X.`,
				},
				resource.Attribute{
					Name:        "alias",
					Description: `Alias will be displayed with the interface name to make it easier to distinguish.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Bring the interface up or shut the interface down.`,
				},
				resource.Attribute{
					Name:        "device_identification",
					Description: `Enable/disable passively gathering of device identity information about the devices on the network connected to this interface.`,
				},
				resource.Attribute{
					Name:        "tcp_mss",
					Description: `TCP maximum segment size. 0 means do not change segment size.`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `Interface speed. The default setting and the options available depend on the interface hardware.`,
				},
				resource.Attribute{
					Name:        "mtu_override",
					Description: `Enable to set a custom MTU for this interface.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `MTU value for this interface.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `Interface role.`,
				},
				resource.Attribute{
					Name:        "allowaccess",
					Description: `Permitted types of management access to this interface.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Required) Addressing mode.`,
				},
				resource.Attribute{
					Name:        "dns_server_override",
					Description: `Enable/disable use DNS acquired by DHCP or PPPoE.`,
				},
				resource.Attribute{
					Name:        "defaultgw",
					Description: `Enable to get the gateway IP from the DHCP or PPPoE server.`,
				},
				resource.Attribute{
					Name:        "distance",
					Description: `Distance for routes learned through PPPoE or DHCP, lower distance indicates preferred route.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `Interface name.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name.`,
				},
				resource.Attribute{
					Name:        "vdom",
					Description: `Interface is in this virtual domain (VDOM).`,
				},
				resource.Attribute{
					Name:        "vlanid",
					Description: `VLAN ID. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Name of the interface.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Interface IPv4 address and subnet mask, syntax` + "`" + ` - X.X.X.X X.X.X.X.`,
				},
				resource.Attribute{
					Name:        "alias",
					Description: `Alias will be displayed with the interface name to make it easier to distinguish.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Bring the interface up or shut the interface down.`,
				},
				resource.Attribute{
					Name:        "device_identification",
					Description: `Enable/disable passively gathering of device identity information about the devices on the network connected to this interface.`,
				},
				resource.Attribute{
					Name:        "tcp_mss",
					Description: `TCP maximum segment size. 0 means do not change segment size.`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `Interface speed. The default setting and the options available depend on the interface hardware.`,
				},
				resource.Attribute{
					Name:        "mtu_override",
					Description: `Enable to set a custom MTU for this interface.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `MTU value for this interface.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `Interface role.`,
				},
				resource.Attribute{
					Name:        "allowaccess",
					Description: `Permitted types of management access to this interface.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Addressing mode.`,
				},
				resource.Attribute{
					Name:        "dns_server_override",
					Description: `Enable/disable use DNS acquired by DHCP or PPPoE.`,
				},
				resource.Attribute{
					Name:        "defaultgw",
					Description: `Enable to get the gateway IP from the DHCP or PPPoE server.`,
				},
				resource.Attribute{
					Name:        "distance",
					Description: `Distance for routes learned through PPPoE or DHCP, lower distance indicates preferred route.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Interface type (support physical, vlan, loopback).`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `Interface name.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name.`,
				},
				resource.Attribute{
					Name:        "vdom",
					Description: `Interface is in this virtual domain (VDOM).`,
				},
				resource.Attribute{
					Name:        "vlanid",
					Description: `VLAN ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Name of the interface.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Interface IPv4 address and subnet mask, syntax` + "`" + ` - X.X.X.X X.X.X.X.`,
				},
				resource.Attribute{
					Name:        "alias",
					Description: `Alias will be displayed with the interface name to make it easier to distinguish.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Bring the interface up or shut the interface down.`,
				},
				resource.Attribute{
					Name:        "device_identification",
					Description: `Enable/disable passively gathering of device identity information about the devices on the network connected to this interface.`,
				},
				resource.Attribute{
					Name:        "tcp_mss",
					Description: `TCP maximum segment size. 0 means do not change segment size.`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `Interface speed. The default setting and the options available depend on the interface hardware.`,
				},
				resource.Attribute{
					Name:        "mtu_override",
					Description: `Enable to set a custom MTU for this interface.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `MTU value for this interface.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `Interface role.`,
				},
				resource.Attribute{
					Name:        "allowaccess",
					Description: `Permitted types of management access to this interface.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Addressing mode.`,
				},
				resource.Attribute{
					Name:        "dns_server_override",
					Description: `Enable/disable use DNS acquired by DHCP or PPPoE.`,
				},
				resource.Attribute{
					Name:        "defaultgw",
					Description: `Enable to get the gateway IP from the DHCP or PPPoE server.`,
				},
				resource.Attribute{
					Name:        "distance",
					Description: `Distance for routes learned through PPPoE or DHCP, lower distance indicates preferred route.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Interface type (support physical, vlan, loopback).`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `Interface name.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name.`,
				},
				resource.Attribute{
					Name:        "vdom",
					Description: `Interface is in this virtual domain (VDOM).`,
				},
				resource.Attribute{
					Name:        "vlanid",
					Description: `VLAN ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_networking_route_static",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to configure static route of FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"networking",
				"route",
				"static",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dst",
					Description: `(Required) Destination IP and mask for this route.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `(Required) Gateway IP for this route.`,
				},
				resource.Attribute{
					Name:        "blackhole",
					Description: `Enable/disable black hole.`,
				},
				resource.Attribute{
					Name:        "distance",
					Description: `Administrative distance.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Administrative weight.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Administrative priority.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `(Required) Gateway out interface or tunnel.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Optional comments.`,
				},
				resource.Attribute{
					Name:        "internet_service",
					Description: `Application ID in the Internet service database.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable this static route. default is "enable". ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the static route item.`,
				},
				resource.Attribute{
					Name:        "dst",
					Description: `Destination IP and mask for this route.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `Gateway IP for this route.`,
				},
				resource.Attribute{
					Name:        "blackhole",
					Description: `Enable/disable black hole.`,
				},
				resource.Attribute{
					Name:        "distance",
					Description: `Administrative distance.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Administrative weight.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Administrative priority.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `Gateway out interface or tunnel.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Optional comments.`,
				},
				resource.Attribute{
					Name:        "internet_service",
					Description: `Application ID in the Internet service database.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable this static route.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the static route item.`,
				},
				resource.Attribute{
					Name:        "dst",
					Description: `Destination IP and mask for this route.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `Gateway IP for this route.`,
				},
				resource.Attribute{
					Name:        "blackhole",
					Description: `Enable/disable black hole.`,
				},
				resource.Attribute{
					Name:        "distance",
					Description: `Administrative distance.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Administrative weight.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Administrative priority.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `Gateway out interface or tunnel.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Optional comments.`,
				},
				resource.Attribute{
					Name:        "internet_service",
					Description: `Application ID in the Internet service database.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable this static route.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_system_admin_administrator",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to configure administrator accounts of FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"system",
				"admin",
				"administrator",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) User name.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Admin user password.`,
				},
				resource.Attribute{
					Name:        "trusthostN",
					Description: `Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit.`,
				},
				resource.Attribute{
					Name:        "vdom",
					Description: `Virtual domain(s) that the administrator can access.`,
				},
				resource.Attribute{
					Name:        "accprofile",
					Description: `(Required) Access profile for this administrator. Access profiles control administrator access to FortiGate features.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the administrator account item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `User name.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Admin user password.`,
				},
				resource.Attribute{
					Name:        "trusthostN",
					Description: `Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit.`,
				},
				resource.Attribute{
					Name:        "vdom",
					Description: `Virtual domain(s) that the administrator can access.`,
				},
				resource.Attribute{
					Name:        "accprofile",
					Description: `Access profile for this administrator. Access profiles control administrator access to FortiGate features.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the administrator account item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `User name.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Admin user password.`,
				},
				resource.Attribute{
					Name:        "trusthostN",
					Description: `Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit.`,
				},
				resource.Attribute{
					Name:        "vdom",
					Description: `Virtual domain(s) that the administrator can access.`,
				},
				resource.Attribute{
					Name:        "accprofile",
					Description: `Access profile for this administrator. Access profiles control administrator access to FortiGate features.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_system_admin_profiles",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to configure access profiles of FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"system",
				"admin",
				"profiles",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Profile name.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `Scope of admin access.`,
				},
				resource.Attribute{
					Name:        "secfabgrp",
					Description: `Security Fabric.`,
				},
				resource.Attribute{
					Name:        "ftviewgrp",
					Description: `FortiView.`,
				},
				resource.Attribute{
					Name:        "authgrp",
					Description: `Administrator access to Users and Devices.`,
				},
				resource.Attribute{
					Name:        "sysgrp",
					Description: `System Configuration.`,
				},
				resource.Attribute{
					Name:        "netgrp",
					Description: `Network Configuration.`,
				},
				resource.Attribute{
					Name:        "loggrp",
					Description: `Administrator access to Logging and Reporting including viewing log messages.`,
				},
				resource.Attribute{
					Name:        "fwgrp",
					Description: `Administrator access to the Firewall configuration.`,
				},
				resource.Attribute{
					Name:        "vpngrp",
					Description: `Administrator access to IPsec, SSL, PPTP, and L2TP VPN.`,
				},
				resource.Attribute{
					Name:        "utmgrp",
					Description: `Administrator access to Security Profiles.`,
				},
				resource.Attribute{
					Name:        "wanoptgrp",
					Description: `Administrator access to WAN Opt & Cache.`,
				},
				resource.Attribute{
					Name:        "wifi",
					Description: `Administrator access to the WiFi controller and Switch controller.`,
				},
				resource.Attribute{
					Name:        "admintimeout_override",
					Description: `Enable/disable overriding the global administrator idle timeout.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the access profile item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Profile name.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `Scope of admin access.`,
				},
				resource.Attribute{
					Name:        "secfabgrp",
					Description: `Security Fabric.`,
				},
				resource.Attribute{
					Name:        "ftviewgrp",
					Description: `FortiView.`,
				},
				resource.Attribute{
					Name:        "authgrp",
					Description: `Administrator access to Users and Devices.`,
				},
				resource.Attribute{
					Name:        "sysgrp",
					Description: `System Configuration.`,
				},
				resource.Attribute{
					Name:        "netgrp",
					Description: `Network Configuration.`,
				},
				resource.Attribute{
					Name:        "loggrp",
					Description: `Administrator access to Logging and Reporting including viewing log messages.`,
				},
				resource.Attribute{
					Name:        "fwgrp",
					Description: `Administrator access to the Firewall configuration.`,
				},
				resource.Attribute{
					Name:        "vpngrp",
					Description: `Administrator access to IPsec, SSL, PPTP, and L2TP VPN.`,
				},
				resource.Attribute{
					Name:        "utmgrp",
					Description: `Administrator access to Security Profiles.`,
				},
				resource.Attribute{
					Name:        "wanoptgrp",
					Description: `Administrator access to WAN Opt & Cache.`,
				},
				resource.Attribute{
					Name:        "wifi",
					Description: `Administrator access to the WiFi controller and Switch controller.`,
				},
				resource.Attribute{
					Name:        "admintimeout_override",
					Description: `Enable/disable overriding the global administrator idle timeout.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the access profile item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Profile name.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `Scope of admin access.`,
				},
				resource.Attribute{
					Name:        "secfabgrp",
					Description: `Security Fabric.`,
				},
				resource.Attribute{
					Name:        "ftviewgrp",
					Description: `FortiView.`,
				},
				resource.Attribute{
					Name:        "authgrp",
					Description: `Administrator access to Users and Devices.`,
				},
				resource.Attribute{
					Name:        "sysgrp",
					Description: `System Configuration.`,
				},
				resource.Attribute{
					Name:        "netgrp",
					Description: `Network Configuration.`,
				},
				resource.Attribute{
					Name:        "loggrp",
					Description: `Administrator access to Logging and Reporting including viewing log messages.`,
				},
				resource.Attribute{
					Name:        "fwgrp",
					Description: `Administrator access to the Firewall configuration.`,
				},
				resource.Attribute{
					Name:        "vpngrp",
					Description: `Administrator access to IPsec, SSL, PPTP, and L2TP VPN.`,
				},
				resource.Attribute{
					Name:        "utmgrp",
					Description: `Administrator access to Security Profiles.`,
				},
				resource.Attribute{
					Name:        "wanoptgrp",
					Description: `Administrator access to WAN Opt & Cache.`,
				},
				resource.Attribute{
					Name:        "wifi",
					Description: `Administrator access to the WiFi controller and Switch controller.`,
				},
				resource.Attribute{
					Name:        "admintimeout_override",
					Description: `Enable/disable overriding the global administrator idle timeout.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_system_apiuser_setting",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to configure API users of FortiOS. The API user of the token for this feature should have a super admin profile, It can be set in CLI while GUI does not allow.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"system",
				"apiuser",
				"setting",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) User name.`,
				},
				resource.Attribute{
					Name:        "accprofile",
					Description: `(Required) Admin user access profile.`,
				},
				resource.Attribute{
					Name:        "vdom",
					Description: `(Required) Virtual domains.`,
				},
				resource.Attribute{
					Name:        "trusthost-Type",
					Description: `(Required) Trusthost type.`,
				},
				resource.Attribute{
					Name:        "trusthost-ipv4_trusthost",
					Description: `(Required) IPv4 trusted host address.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the API user.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `User name.`,
				},
				resource.Attribute{
					Name:        "accprofile",
					Description: `Admin user access profile.`,
				},
				resource.Attribute{
					Name:        "vdom",
					Description: `Virtual domains.`,
				},
				resource.Attribute{
					Name:        "trusthost-Type",
					Description: `Trusthost type.`,
				},
				resource.Attribute{
					Name:        "trusthost-ipv4_trusthost",
					Description: `IPv4 trusted host address.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the API user.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `User name.`,
				},
				resource.Attribute{
					Name:        "accprofile",
					Description: `Admin user access profile.`,
				},
				resource.Attribute{
					Name:        "vdom",
					Description: `Virtual domains.`,
				},
				resource.Attribute{
					Name:        "trusthost-Type",
					Description: `Trusthost type.`,
				},
				resource.Attribute{
					Name:        "trusthost-ipv4_trusthost",
					Description: `IPv4 trusted host address.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_system_license_forticare",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to add a FortiCare license for FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"system",
				"license",
				"forticare",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "registration_code",
					Description: `(Required) Registration code.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_system_license_vdom",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to add a VDOM license for FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"system",
				"license",
				"vdom",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "license",
					Description: `(Required) Registration code.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_system_license_vm",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to update VM license using uploaded file for FortiOS. Reboots immediately if successful.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"system",
				"license",
				"vm",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "file_content",
					Description: `(Required) The license file, it needs to be base64 encoded, must not contain whitespace or other invalid base64 characters, and must be included in HTTP body.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_system_setting_dns",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to configure DNS of FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"system",
				"setting",
				"dns",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "primary",
					Description: `Primary DNS server IP address.`,
				},
				resource.Attribute{
					Name:        "secondary",
					Description: `Secondary DNS server IP address.`,
				},
				resource.Attribute{
					Name:        "dns_over_tls",
					Description: `Enable/disable/enforce DNS over TLS(available since v6.2.0). Enum: [ disable, enable, enforce ] ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `Primary DNS server IP address.`,
				},
				resource.Attribute{
					Name:        "secondary",
					Description: `Secondary DNS server IP address.`,
				},
				resource.Attribute{
					Name:        "dns_over_tls",
					Description: `Enable/disable/enforce DNS over TLS(available since v6.2.0).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "primary",
					Description: `Primary DNS server IP address.`,
				},
				resource.Attribute{
					Name:        "secondary",
					Description: `Secondary DNS server IP address.`,
				},
				resource.Attribute{
					Name:        "dns_over_tls",
					Description: `Enable/disable/enforce DNS over TLS(available since v6.2.0).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_system_setting_global",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to configure options related to the overall operation of FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"system",
				"setting",
				"global",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required) FortiGate unit's hostname.`,
				},
				resource.Attribute{
					Name:        "admintimeout",
					Description: `Number of minutes before an idle administrator session time out.`,
				},
				resource.Attribute{
					Name:        "timezone",
					Description: `Number corresponding to your time zone from 00 to 86.`,
				},
				resource.Attribute{
					Name:        "admin_sport",
					Description: `Administrative access port for HTTPS.`,
				},
				resource.Attribute{
					Name:        "admin_ssh_port",
					Description: `Administrative access port for SSH.`,
				},
				resource.Attribute{
					Name:        "admin_scp",
					Description: `Enable SCP over SSH ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "admintimeout",
					Description: `Number of minutes before an idle administrator session time out.`,
				},
				resource.Attribute{
					Name:        "timezone",
					Description: `Number corresponding to your time zone from 00 to 86.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `FortiGate unit's hostname.`,
				},
				resource.Attribute{
					Name:        "admin_sport",
					Description: `Administrative access port for HTTPS.`,
				},
				resource.Attribute{
					Name:        "admin_ssh_port",
					Description: `Administrative access port for SSH.`,
				},
				resource.Attribute{
					Name:        "admin_scp",
					Description: `Enable SCP over SSH`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "admintimeout",
					Description: `Number of minutes before an idle administrator session time out.`,
				},
				resource.Attribute{
					Name:        "timezone",
					Description: `Number corresponding to your time zone from 00 to 86.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `FortiGate unit's hostname.`,
				},
				resource.Attribute{
					Name:        "admin_sport",
					Description: `Administrative access port for HTTPS.`,
				},
				resource.Attribute{
					Name:        "admin_ssh_port",
					Description: `Administrative access port for SSH.`,
				},
				resource.Attribute{
					Name:        "admin_scp",
					Description: `Enable SCP over SSH`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_system_setting_ntp",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to configure Network Time Protocol (NTP) servers of FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"system",
				"setting",
				"ntp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Use the FortiGuard NTP server or any other available NTP Server.`,
				},
				resource.Attribute{
					Name:        "ntpserver",
					Description: `Configure the FortiGate to connect to any available third-party NTP server.`,
				},
				resource.Attribute{
					Name:        "ntpsync",
					Description: `Enable/disable setting the FortiGate system time by synchronizing with an NTP Server. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Use the FortiGuard NTP server or any other available NTP Server.`,
				},
				resource.Attribute{
					Name:        "ntpserver",
					Description: `Configure the FortiGate to connect to any available third-party NTP server.`,
				},
				resource.Attribute{
					Name:        "ntpsync",
					Description: `Enable/disable setting the FortiGate system time by synchronizing with an NTP Server.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `Use the FortiGuard NTP server or any other available NTP Server.`,
				},
				resource.Attribute{
					Name:        "ntpserver",
					Description: `Configure the FortiGate to connect to any available third-party NTP server.`,
				},
				resource.Attribute{
					Name:        "ntpsync",
					Description: `Enable/disable setting the FortiGate system time by synchronizing with an NTP Server.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_system_vdom_setting",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to configure VDOM of FortiOS. The API user of the token for this feature should have a super admin profile, It can be set in CLI while GUI does not allow.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"system",
				"vdom",
				"setting",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) VDOM name.`,
				},
				resource.Attribute{
					Name:        "short_name",
					Description: `VDOM short name.`,
				},
				resource.Attribute{
					Name:        "temporary",
					Description: `Temporary. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the VDOM.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `VDOM name.`,
				},
				resource.Attribute{
					Name:        "short_name",
					Description: `VDOM short name.`,
				},
				resource.Attribute{
					Name:        "temporary",
					Description: `Temporary.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the VDOM.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `VDOM name.`,
				},
				resource.Attribute{
					Name:        "short_name",
					Description: `VDOM short name.`,
				},
				resource.Attribute{
					Name:        "temporary",
					Description: `Temporary.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_vpn_ipsec_phase1interface",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to use phase1-interface to define a phase 1 definition for a route-based (interface mode) IPsec VPN tunnel that generates authentication and encryption keys automatically.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"vpn",
				"ipsec",
				"phase1interface",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) IPsec remote gateway name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Remote gateway type.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(Required) Local physical, aggregate, or VLAN outgoing interface.`,
				},
				resource.Attribute{
					Name:        "peertype",
					Description: `Accept this peer type.`,
				},
				resource.Attribute{
					Name:        "proposal",
					Description: `Phase1 proposal.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
				resource.Attribute{
					Name:        "wizard_type",
					Description: `GUI VPN Wizard Type.`,
				},
				resource.Attribute{
					Name:        "remote_gw",
					Description: `(Required) IPv4 address of the remote gateway's external interface.`,
				},
				resource.Attribute{
					Name:        "psksecret",
					Description: `(Required) Pre-shared secret for PSK authentication.`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `Names of signed personal certificates.`,
				},
				resource.Attribute{
					Name:        "peerid",
					Description: `Accept this peer identity.`,
				},
				resource.Attribute{
					Name:        "peer",
					Description: `Accept this peer certificate.`,
				},
				resource.Attribute{
					Name:        "peergrp",
					Description: `Accept this peer certificate group.`,
				},
				resource.Attribute{
					Name:        "ipv4_split_include",
					Description: `IPv4 split-include subnets.`,
				},
				resource.Attribute{
					Name:        "split_include_service",
					Description: `Split-include services.`,
				},
				resource.Attribute{
					Name:        "ipv4_split_exclude",
					Description: `IPv4 subnets that should not be sent over the IPsec tunnel.`,
				},
				resource.Attribute{
					Name:        "authmethod",
					Description: `Authentication method.`,
				},
				resource.Attribute{
					Name:        "authmethod_remote",
					Description: `Authentication method (remote side).`,
				},
				resource.Attribute{
					Name:        "mode_cfg",
					Description: `Enable/disable configuration method. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the phase1-interface item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `IPsec remote gateway name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Remote gateway type.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `Local physical, aggregate, or VLAN outgoing interface.`,
				},
				resource.Attribute{
					Name:        "peertype",
					Description: `Accept this peer type.`,
				},
				resource.Attribute{
					Name:        "proposal",
					Description: `Phase1 proposal.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
				resource.Attribute{
					Name:        "wizard_type",
					Description: `GUI VPN Wizard Type.`,
				},
				resource.Attribute{
					Name:        "remote_gw",
					Description: `IPv4 address of the remote gateway's external interface.`,
				},
				resource.Attribute{
					Name:        "psksecret",
					Description: `Pre-shared secret for PSK authentication.`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `Names of signed personal certificates.`,
				},
				resource.Attribute{
					Name:        "peerid",
					Description: `Accept this peer identity.`,
				},
				resource.Attribute{
					Name:        "peer",
					Description: `Accept this peer certificate.`,
				},
				resource.Attribute{
					Name:        "peergrp",
					Description: `Accept this peer certificate group.`,
				},
				resource.Attribute{
					Name:        "ipv4_split_include",
					Description: `IPv4 split-include subnets.`,
				},
				resource.Attribute{
					Name:        "split_include_service",
					Description: `Split-include services.`,
				},
				resource.Attribute{
					Name:        "ipv4_split_exclude",
					Description: `IPv4 subnets that should not be sent over the IPsec tunnel.`,
				},
				resource.Attribute{
					Name:        "authmethod",
					Description: `Authentication method.`,
				},
				resource.Attribute{
					Name:        "authmethod_remote",
					Description: `Authentication method (remote side).`,
				},
				resource.Attribute{
					Name:        "mode_cfg",
					Description: `Enable/disable configuration method.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the phase1-interface item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `IPsec remote gateway name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Remote gateway type.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `Local physical, aggregate, or VLAN outgoing interface.`,
				},
				resource.Attribute{
					Name:        "peertype",
					Description: `Accept this peer type.`,
				},
				resource.Attribute{
					Name:        "proposal",
					Description: `Phase1 proposal.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
				resource.Attribute{
					Name:        "wizard_type",
					Description: `GUI VPN Wizard Type.`,
				},
				resource.Attribute{
					Name:        "remote_gw",
					Description: `IPv4 address of the remote gateway's external interface.`,
				},
				resource.Attribute{
					Name:        "psksecret",
					Description: `Pre-shared secret for PSK authentication.`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `Names of signed personal certificates.`,
				},
				resource.Attribute{
					Name:        "peerid",
					Description: `Accept this peer identity.`,
				},
				resource.Attribute{
					Name:        "peer",
					Description: `Accept this peer certificate.`,
				},
				resource.Attribute{
					Name:        "peergrp",
					Description: `Accept this peer certificate group.`,
				},
				resource.Attribute{
					Name:        "ipv4_split_include",
					Description: `IPv4 split-include subnets.`,
				},
				resource.Attribute{
					Name:        "split_include_service",
					Description: `Split-include services.`,
				},
				resource.Attribute{
					Name:        "ipv4_split_exclude",
					Description: `IPv4 subnets that should not be sent over the IPsec tunnel.`,
				},
				resource.Attribute{
					Name:        "authmethod",
					Description: `Authentication method.`,
				},
				resource.Attribute{
					Name:        "authmethod_remote",
					Description: `Authentication method (remote side).`,
				},
				resource.Attribute{
					Name:        "mode_cfg",
					Description: `Enable/disable configuration method.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_vpn_ipsec_phase2interface",
			Category:         "FortiGate Resources",
			ShortDescription: `Provides a resource to use phase2-interface to add or edit a phase 2 configuration on a route-based (interface mode) IPsec tunnel.`,
			Description:      ``,
			Keywords: []string{
				"fortigate",
				"vpn",
				"ipsec",
				"phase2interface",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) IPsec tunnel name.`,
				},
				resource.Attribute{
					Name:        "phase1name",
					Description: `(Required) Phase 1 determines the options required for phase 2.`,
				},
				resource.Attribute{
					Name:        "proposal",
					Description: `Phase2 proposal.`,
				},
				resource.Attribute{
					Name:        "src_addr_type",
					Description: `Local proxy ID type.`,
				},
				resource.Attribute{
					Name:        "src_start_ip",
					Description: `Local proxy ID start.`,
				},
				resource.Attribute{
					Name:        "src_end_ip",
					Description: `Local proxy ID end.`,
				},
				resource.Attribute{
					Name:        "src_subnet",
					Description: `Local proxy ID subnet.`,
				},
				resource.Attribute{
					Name:        "dst_addr_type",
					Description: `Local proxy ID type.`,
				},
				resource.Attribute{
					Name:        "src_name",
					Description: `Local proxy ID name.`,
				},
				resource.Attribute{
					Name:        "dst_name",
					Description: `Remote proxy ID name.`,
				},
				resource.Attribute{
					Name:        "dst_start_ip",
					Description: `Remote proxy ID IPv4 start.`,
				},
				resource.Attribute{
					Name:        "dst_end_ip",
					Description: `Remote proxy ID IPv4 end.`,
				},
				resource.Attribute{
					Name:        "dst_subnet",
					Description: `Remote proxy ID IPv4 subnet.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the phase2-interface.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `IPsec tunnel name.`,
				},
				resource.Attribute{
					Name:        "phase1name",
					Description: `Phase 1 determines the options required for phase 2.`,
				},
				resource.Attribute{
					Name:        "proposal",
					Description: `Phase2 proposal.`,
				},
				resource.Attribute{
					Name:        "src_addr_type",
					Description: `Local proxy ID type.`,
				},
				resource.Attribute{
					Name:        "src_start_ip",
					Description: `Local proxy ID start.`,
				},
				resource.Attribute{
					Name:        "src_end_ip",
					Description: `Local proxy ID end.`,
				},
				resource.Attribute{
					Name:        "src_subnet",
					Description: `Local proxy ID subnet.`,
				},
				resource.Attribute{
					Name:        "dst_addr_type",
					Description: `Local proxy ID type.`,
				},
				resource.Attribute{
					Name:        "src_name",
					Description: `Local proxy ID name.`,
				},
				resource.Attribute{
					Name:        "dst_name",
					Description: `Remote proxy ID name.`,
				},
				resource.Attribute{
					Name:        "dst_start_ip",
					Description: `Remote proxy ID IPv4 start.`,
				},
				resource.Attribute{
					Name:        "dst_end_ip",
					Description: `Remote proxy ID IPv4 end.`,
				},
				resource.Attribute{
					Name:        "dst_subnet",
					Description: `Remote proxy ID IPv4 subnet.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the phase2-interface.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `IPsec tunnel name.`,
				},
				resource.Attribute{
					Name:        "phase1name",
					Description: `Phase 1 determines the options required for phase 2.`,
				},
				resource.Attribute{
					Name:        "proposal",
					Description: `Phase2 proposal.`,
				},
				resource.Attribute{
					Name:        "src_addr_type",
					Description: `Local proxy ID type.`,
				},
				resource.Attribute{
					Name:        "src_start_ip",
					Description: `Local proxy ID start.`,
				},
				resource.Attribute{
					Name:        "src_end_ip",
					Description: `Local proxy ID end.`,
				},
				resource.Attribute{
					Name:        "src_subnet",
					Description: `Local proxy ID subnet.`,
				},
				resource.Attribute{
					Name:        "dst_addr_type",
					Description: `Local proxy ID type.`,
				},
				resource.Attribute{
					Name:        "src_name",
					Description: `Local proxy ID name.`,
				},
				resource.Attribute{
					Name:        "dst_name",
					Description: `Remote proxy ID name.`,
				},
				resource.Attribute{
					Name:        "dst_start_ip",
					Description: `Remote proxy ID IPv4 start.`,
				},
				resource.Attribute{
					Name:        "dst_end_ip",
					Description: `Remote proxy ID IPv4 end.`,
				},
				resource.Attribute{
					Name:        "dst_subnet",
					Description: `Remote proxy ID IPv4 subnet.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"fortios_fmg_devicemanager_device":                0,
		"fortios_fmg_devicemanager_install_device":        1,
		"fortios_fmg_devicemanager_install_policypackage": 2,
		"fortios_fmg_devicemanager_script":                3,
		"fortios_fmg_devicemanager_script_execute":        4,
		"fortios_fmg_firewall_object_address":             5,
		"fortios_fmg_firewall_object_ippool":              6,
		"fortios_fmg_firewall_object_service":             7,
		"fortios_fmg_firewall_object_vip":                 8,
		"fortios_fmg_firewall_security_policy":            9,
		"fortios_fmg_firewall_security_policypackage":     10,
		"fortios_fmg_jsonrpc_request":                     11,
		"fortios_fmg_object_adom_revision":                12,
		"fortios_fmg_system_admin":                        13,
		"fortios_fmg_system_admin_profiles":               14,
		"fortios_fmg_system_admin_user":                   15,
		"fortios_fmg_system_adom":                         16,
		"fortios_fmg_system_dns":                          17,
		"fortios_fmg_system_global":                       18,
		"fortios_fmg_system_license_forticare":            19,
		"fortios_fmg_system_license_vm":                   20,
		"fortios_fmg_system_network_interface":            21,
		"fortios_fmg_system_network_route":                22,
		"fortios_fmg_system_ntp":                          23,
		"fortios_fmg_system_syslogserver":                 24,
		"fortios_firewall_object_address":                 25,
		"fortios_firewall_object_addressgroup":            26,
		"fortios_firewall_object_ippool":                  27,
		"fortios_firewall_object_service":                 28,
		"fortios_firewall_object_servicecategory":         29,
		"fortios_firewall_object_servicegroup":            30,
		"fortios_firewall_object_vip":                     31,
		"fortios_firewall_object_vipgroup":                32,
		"fortios_firewall_security_policy":                33,
		"fortios_firewall_security_policyseq":             34,
		"fortios_log_fortianalyzer_setting":               35,
		"fortios_log_syslog_setting":                      36,
		"fortios_networking_interface_port":               37,
		"fortios_networking_route_static":                 38,
		"fortios_system_admin_administrator":              39,
		"fortios_system_admin_profiles":                   40,
		"fortios_system_apiuser_setting":                  41,
		"fortios_system_license_forticare":                42,
		"fortios_system_license_vdom":                     43,
		"fortios_system_license_vm":                       44,
		"fortios_system_setting_dns":                      45,
		"fortios_system_setting_global":                   46,
		"fortios_system_setting_ntp":                      47,
		"fortios_system_vdom_setting":                     48,
		"fortios_vpn_ipsec_phase1interface":               49,
		"fortios_vpn_ipsec_phase2interface":               50,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
