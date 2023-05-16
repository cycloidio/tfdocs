package azurestack

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "azurestack_availability_set",
			Category:         "Compute Resources",
			ShortDescription: `Manages an availability set for virtual machines.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"availability",
				"set",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the availability set. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The name of the resource group in which to create the availability set. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Required) Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "platform_update_domain_count",
					Description: `(Optional) Specifies the number of update domains that are used. Defaults to 5. ~>`,
				},
				resource.Attribute{
					Name:        "platform_fault_domain_count",
					Description: `(Optional) Specifies the number of fault domains that are used. Defaults to 3. ~>`,
				},
				resource.Attribute{
					Name:        "managed",
					Description: `(Optional) Specifies whether the availability set is managed or not. Possible values are ` + "`" + `true` + "`" + ` (to specify aligned) or ` + "`" + `false` + "`" + ` (to specify classic). Default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The virtual Availability Set ID. ## Import Availability Sets can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azurestack_availability_set.group1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/availabilitySets/webAvailSet ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The virtual Availability Set ID. ## Import Availability Sets can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azurestack_availability_set.group1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/availabilitySets/webAvailSet ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurestack_dns_a_record",
			Category:         "DNS Resources",
			ShortDescription: `Manages a DNS A Record.`,
			Description:      ``,
			Keywords: []string{
				"dns",
				"a",
				"record",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the DNS A Record.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) Specifies the resource group where the resource exists. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "zone_name",
					Description: `(Required) Specifies the DNS Zone where the resource exists. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "TTL",
					Description: `(Required) The Time To Live (TTL) of the DNS record.`,
				},
				resource.Attribute{
					Name:        "records",
					Description: `(Required) List of IPv4 Addresses.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The DNS A Record ID. ## Import A records can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azurestack_dns_a_record.test /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/dnsZones/zone1/A/myrecord1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The DNS A Record ID. ## Import A records can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azurestack_dns_a_record.test /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/dnsZones/zone1/A/myrecord1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurestack_dns_zone",
			Category:         "DNS Resources",
			ShortDescription: `Manages a DNS Zone.`,
			Description:      ``,
			Keywords: []string{
				"dns",
				"zone",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the DNS Zone. Must be a valid domain name.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) Specifies the resource group where the resource exists. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The DNS Zone ID.`,
				},
				resource.Attribute{
					Name:        "max_number_of_record_sets",
					Description: `(Optional) Maximum number of Records in the zone. Defaults to ` + "`" + `1000` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "number_of_record_sets",
					Description: `(Optional) The number of records already in the zone.`,
				},
				resource.Attribute{
					Name:        "name_servers",
					Description: `(Optional) A list of values that make up the NS record for the zone. ## Import DNS Zones can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azurestack_dns_zone.zone1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/dnsZones/zone1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The DNS Zone ID.`,
				},
				resource.Attribute{
					Name:        "max_number_of_record_sets",
					Description: `(Optional) Maximum number of Records in the zone. Defaults to ` + "`" + `1000` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "number_of_record_sets",
					Description: `(Optional) The number of records already in the zone.`,
				},
				resource.Attribute{
					Name:        "name_servers",
					Description: `(Optional) A list of values that make up the NS record for the zone. ## Import DNS Zones can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azurestack_dns_zone.zone1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/dnsZones/zone1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurestack_lb",
			Category:         "Load Balancer Resources",
			ShortDescription: `Manages a Load Balancer Resource.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
				"lb",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the LoadBalancer.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The name of the resource group in which to create the LoadBalancer.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Required) Specifies the supported Azure location where the resource exists.`,
				},
				resource.Attribute{
					Name:        "frontend_ip_configuration",
					Description: `(Optional) A frontend ip configuration block as documented below.`,
				},
				resource.Attribute{
					Name:        "sku",
					Description: `(Optional) The SKU of the Azure Load Balancer. Accepted values are ` + "`" + `Basic` + "`" + ` and ` + "`" + `Standard` + "`" + `. Defaults to ` + "`" + `Basic` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. ` + "`" + `frontend_ip_configuration` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the frontend ip configuration.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) Reference to subnet associated with the IP Configuration.`,
				},
				resource.Attribute{
					Name:        "private_ip_address",
					Description: `(Optional) Private IP Address to assign to the Load Balancer. The last one and first four IPs in any range are reserved and cannot be manually assigned.`,
				},
				resource.Attribute{
					Name:        "private_ip_address_allocation",
					Description: `(Optional) Defines how a private IP address is assigned. Options are Static or Dynamic.`,
				},
				resource.Attribute{
					Name:        "public_ip_address_id",
					Description: `(Optional) Reference to Public IP address to be associated with the Load Balancer. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The LoadBalancer ID.`,
				},
				resource.Attribute{
					Name:        "private_ip_address",
					Description: `The first private IP address assigned to the load balancer in ` + "`" + `frontend_ip_configuration` + "`" + ` blocks, if any.`,
				},
				resource.Attribute{
					Name:        "private_ip_addresses",
					Description: `The list of private IP address assigned to the load balancer in ` + "`" + `frontend_ip_configuration` + "`" + ` blocks, if any. ## Import Load Balancers can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azurestack_lb.test /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/lb1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The LoadBalancer ID.`,
				},
				resource.Attribute{
					Name:        "private_ip_address",
					Description: `The first private IP address assigned to the load balancer in ` + "`" + `frontend_ip_configuration` + "`" + ` blocks, if any.`,
				},
				resource.Attribute{
					Name:        "private_ip_addresses",
					Description: `The list of private IP address assigned to the load balancer in ` + "`" + `frontend_ip_configuration` + "`" + ` blocks, if any. ## Import Load Balancers can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azurestack_lb.test /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/lb1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurestack_lb_backend_address_pool",
			Category:         "Load Balancer Resources",
			ShortDescription: `Manages a LoadBalancer Backend Address Pool.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
				"lb",
				"backend",
				"address",
				"pool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the Backend Address Pool.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The name of the resource group in which to create the resource.`,
				},
				resource.Attribute{
					Name:        "loadbalancer_id",
					Description: `(Required) The ID of the LoadBalancer in which to create the Backend Address Pool. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the LoadBalancer to which the resource is attached. ## Import Load Balancer Backend Address Pools can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azurestack_lb_backend_address_pool.test /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/lb1/backendAddressPools/pool1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the LoadBalancer to which the resource is attached. ## Import Load Balancer Backend Address Pools can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azurestack_lb_backend_address_pool.test /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/lb1/backendAddressPools/pool1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurestack_lb_nat_pool",
			Category:         "Load Balancer Resources",
			ShortDescription: `Manages a Load Balancer NAT Pool.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
				"lb",
				"nat",
				"pool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the NAT pool.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The name of the resource group in which to create the resource.`,
				},
				resource.Attribute{
					Name:        "loadbalancer_id",
					Description: `(Required) The ID of the Load Balancer in which to create the NAT pool.`,
				},
				resource.Attribute{
					Name:        "frontend_ip_configuration_name",
					Description: `(Required) The name of the frontend IP configuration exposing this rule.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The transport protocol for the external endpoint. Possible values are ` + "`" + `Udp` + "`" + ` or ` + "`" + `Tcp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "frontend_port_start",
					Description: `(Required) The first port number in the range of external ports that will be used to provide Inbound Nat to NICs associated with this Load Balancer. Possible values range between 1 and 65534, inclusive.`,
				},
				resource.Attribute{
					Name:        "frontend_port_end",
					Description: `(Required) The last port number in the range of external ports that will be used to provide Inbound Nat to NICs associated with this Load Balancer. Possible values range between 1 and 65534, inclusive.`,
				},
				resource.Attribute{
					Name:        "backend_port",
					Description: `(Required) The port used for the internal endpoint. Possible values range between 1 and 65535, inclusive. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Load Balancer to which the resource is attached. ## Import Load Balancer NAT Pools can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azurestack_lb_nat_pool.test /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/lb1/inboundNatPools/pool1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Load Balancer to which the resource is attached. ## Import Load Balancer NAT Pools can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azurestack_lb_nat_pool.test /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/lb1/inboundNatPools/pool1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurestack_lb_nat_rule",
			Category:         "Load Balancer Resources",
			ShortDescription: `Manages a LoadBalancer NAT Rule.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
				"lb",
				"nat",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the NAT Rule.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The name of the resource group in which to create the resource.`,
				},
				resource.Attribute{
					Name:        "loadbalancer_id",
					Description: `(Required) The ID of the LoadBalancer in which to create the NAT Rule.`,
				},
				resource.Attribute{
					Name:        "frontend_ip_configuration_name",
					Description: `(Required) The name of the frontend IP configuration exposing this rule.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The transport protocol for the external endpoint. Possible values are ` + "`" + `Udp` + "`" + ` or ` + "`" + `Tcp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "frontend_port",
					Description: `(Required) The port for the external endpoint. Port numbers for each Rule must be unique within the Load Balancer. Possible values range between 1 and 65534, inclusive.`,
				},
				resource.Attribute{
					Name:        "backend_port",
					Description: `(Required) The port used for internal connections on the endpoint. Possible values range between 1 and 65535, inclusive.`,
				},
				resource.Attribute{
					Name:        "enable_floating_ip",
					Description: `(Optional) Enables the Floating IP Capacity, required to configure a SQL AlwaysOn Availability Group. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the LoadBalancer to which the resource is attached. ## Import Load Balancer NAT Rules can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azurestack_lb_nat_rule.test /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/lb1/inboundNatRules/rule1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the LoadBalancer to which the resource is attached. ## Import Load Balancer NAT Rules can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azurestack_lb_nat_rule.test /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/lb1/inboundNatRules/rule1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurestack_lb_probe",
			Category:         "Load Balancer Resources",
			ShortDescription: `Manages a LoadBalancer Probe Resource.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
				"lb",
				"probe",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the Probe.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The name of the resource group in which to create the resource.`,
				},
				resource.Attribute{
					Name:        "loadbalancer_id",
					Description: `(Required) The ID of the LoadBalancer in which to create the NAT Rule.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) Specifies the protocol of the end point. Possible values are ` + "`" + `Http` + "`" + ` or ` + "`" + `Tcp` + "`" + `. If Tcp is specified, a received ACK is required for the probe to be successful. If Http is specified, a 200 OK response from the specified URI is required for the probe to be successful.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) Port on which the Probe queries the backend endpoint. Possible values range from 1 to 65535, inclusive.`,
				},
				resource.Attribute{
					Name:        "request_path",
					Description: `(Optional) The URI used for requesting health status from the backend endpoint. Required if protocol is set to Http. Otherwise, it is not allowed.`,
				},
				resource.Attribute{
					Name:        "interval_in_seconds",
					Description: `(Optional) The interval, in seconds between probes to the backend endpoint for health status. The default value is 15, the minimum value is 5.`,
				},
				resource.Attribute{
					Name:        "number_of_probes",
					Description: `(Optional) The number of failed probe attempts after which the backend endpoint is removed from rotation. The default value is 2. NumberOfProbes multiplied by intervalInSeconds value must be greater or equal to 10.Endpoints are returned to rotation when at least one probe is successful. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the LoadBalancer to which the resource is attached. ## Import Load Balancer Probes can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azurestack_lb_probe.test /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/lb1/probes/probe1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the LoadBalancer to which the resource is attached. ## Import Load Balancer Probes can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azurestack_lb_probe.test /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/lb1/probes/probe1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurestack_lb_rule",
			Category:         "Load Balancer Resources",
			ShortDescription: `Manages a Load Balancer Rule.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
				"lb",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the LB Rule.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The name of the resource group in which to create the resource.`,
				},
				resource.Attribute{
					Name:        "loadbalancer_id",
					Description: `(Required) The ID of the Load Balancer in which to create the Rule.`,
				},
				resource.Attribute{
					Name:        "frontend_ip_configuration_name",
					Description: `(Required) The name of the frontend IP configuration to which the rule is associated.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The transport protocol for the external endpoint. Possible values are ` + "`" + `Udp` + "`" + ` or ` + "`" + `Tcp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "frontend_port",
					Description: `(Required) The port for the external endpoint. Port numbers for each Rule must be unique within the Load Balancer. Possible values range between 1 and 65534, inclusive.`,
				},
				resource.Attribute{
					Name:        "backend_port",
					Description: `(Required) The port used for internal connections on the endpoint. Possible values range between 1 and 65535, inclusive.`,
				},
				resource.Attribute{
					Name:        "backend_address_pool_id",
					Description: `(Optional) A reference to a Backend Address Pool over which this Load Balancing Rule operates.`,
				},
				resource.Attribute{
					Name:        "probe_id",
					Description: `(Optional) A reference to a Probe used by this Load Balancing Rule.`,
				},
				resource.Attribute{
					Name:        "enable_floating_ip",
					Description: `(Optional) Floating IP is pertinent to failover scenarios: a "floating” IP is reassigned to a secondary server in case the primary server fails. Floating IP is required for SQL AlwaysOn.`,
				},
				resource.Attribute{
					Name:        "idle_timeout_in_minutes",
					Description: `(Optional) Specifies the timeout for the Tcp idle connection. The value can be set between 4 and 30 minutes. The default value is 4 minutes. This element is only used when the protocol is set to Tcp.`,
				},
				resource.Attribute{
					Name:        "load_distribution",
					Description: `(Optional) Specifies the load balancing distribution type to be used by the Load Balancer. Possible values are: ` + "`" + `Default` + "`" + ` – The load balancer is configured to use a 5 tuple hash to map traffic to available servers. ` + "`" + `SourceIP` + "`" + ` – The load balancer is configured to use a 2 tuple hash to map traffic to available servers. ` + "`" + `SourceIPProtocol` + "`" + ` – The load balancer is configured to use a 3 tuple hash to map traffic to available servers. Also known as Session Persistence, where the options are called ` + "`" + `None` + "`" + `, ` + "`" + `Client IP` + "`" + ` and ` + "`" + `Client IP and Protocol` + "`" + ` respectively. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Load Balancer to which the resource is attached. ## Import Load Balancer Rules can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azurestack_lb_rule.test /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/lb1/loadBalancingRules/rule1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Load Balancer to which the resource is attached. ## Import Load Balancer Rules can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azurestack_lb_rule.test /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/loadBalancers/lb1/loadBalancingRules/rule1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurestack_local_network_gateway",
			Category:         "Network Resources",
			ShortDescription: `Manages a local network gateway connection over which specific connections can be configured.`,
			Description: `

Manages a local network gateway connection over which specific connections can be configured.

`,
			Keywords: []string{
				"network",
				"local",
				"gateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the local network gateway. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The name of the resource group in which to create the local network gateway.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Required) The location/region where the local network gateway is created. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "gateway_address",
					Description: `(Required) The IP address of the gateway to which to connect.`,
				},
				resource.Attribute{
					Name:        "address_space",
					Description: `(Required) The list of string CIDRs representing the address spaces the gateway exposes.`,
				},
				resource.Attribute{
					Name:        "bgp_settings",
					Description: `(Optional) A ` + "`" + `bgp_settings` + "`" + ` block as defined below containing the Local Network Gateway's BGP speaker settings.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. --- ` + "`" + `bgp_settings` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "asn",
					Description: `(Required) The BGP speaker's ASN.`,
				},
				resource.Attribute{
					Name:        "bgp_peering_address",
					Description: `(Required) The BGP peering address and BGP identifier of this BGP speaker.`,
				},
				resource.Attribute{
					Name:        "peer_weight",
					Description: `(Optional) The weight added to routes learned from this BGP speaker. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The local network gateway unique ID within Azure. ## Import Local Network Gateways can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azurestack_local_network_gateway.lng1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/localNetworkGateways/lng1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The local network gateway unique ID within Azure. ## Import Local Network Gateways can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azurestack_local_network_gateway.lng1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/localNetworkGateways/lng1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurestack_managed_disk",
			Category:         "Compute Resources",
			ShortDescription: `Manages a Managed Disk.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"managed",
				"disk",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the managed disk. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The name of the resource group in which to create the managed disk.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Required) Specified the supported Azure location where the resource exists. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "storage_account_type",
					Description: `(Required) The type of storage to use for the managed disk. Allowable values are ` + "`" + `Standard_LRS` + "`" + ` or ` + "`" + `Premium_LRS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "create_option",
					Description: `(Required) The method to use when creating the managed disk. Possible values include:`,
				},
				resource.Attribute{
					Name:        "Import",
					Description: `Import a VHD file in to the managed disk (VHD specified with ` + "`" + `source_uri` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "Empty",
					Description: `Create an empty managed disk.`,
				},
				resource.Attribute{
					Name:        "Copy",
					Description: `Copy an existing managed disk or snapshot (specified with ` + "`" + `source_resource_id` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "FromImage",
					Description: `Copy a Platform Image (specified with ` + "`" + `image_reference_id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "source_uri",
					Description: `(Optional) URI to a valid VHD file to be used when ` + "`" + `create_option` + "`" + ` is ` + "`" + `Import` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "source_resource_id",
					Description: `(Optional) ID of an existing managed disk to copy ` + "`" + `create_option` + "`" + ` is ` + "`" + `Copy` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "image_reference_id",
					Description: `(Optional) ID of an existing platform/marketplace disk image to copy when ` + "`" + `create_option` + "`" + ` is ` + "`" + `FromImage` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `(Optional) Specify a value when the source of an ` + "`" + `Import` + "`" + ` or ` + "`" + `Copy` + "`" + ` operation targets a source that contains an operating system. Valid values are ` + "`" + `Linux` + "`" + ` or ` + "`" + `Windows` + "`" + ``,
				},
				resource.Attribute{
					Name:        "disk_size_gb",
					Description: `(Optional, Required for a new managed disk) Specifies the size of the managed disk to create in gigabytes. If ` + "`" + `create_option` + "`" + ` is ` + "`" + `Copy` + "`" + ` or ` + "`" + `FromImage` + "`" + `, then the value must be equal to or greater than the source's size.`,
				},
				resource.Attribute{
					Name:        "encryption",
					Description: `(Optional) A ` + "`" + `encryption` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "hyper_v_generation",
					Description: `(Optional) The HyperV Generation of the Disk when the source of an ` + "`" + `Import` + "`" + ` or ` + "`" + `Copy` + "`" + ` operation targets a source that contains an operating system. Possible values are ` + "`" + `V1` + "`" + ` and ` + "`" + `V2` + "`" + `. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. --- The ` + "`" + `encryption` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Is Encryption enabled on this Managed Disk? Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "disk_encryption_key",
					Description: `(Optional) A ` + "`" + `disk_encryption_key` + "`" + ` block as defined above.`,
				},
				resource.Attribute{
					Name:        "key_encryption_key",
					Description: `(Optional) A ` + "`" + `key_encryption_key` + "`" + ` block as defined below. --- For more information on managed disks, such as sizing options and pricing, please check out the [azure documentation](https://docs.microsoft.com/en-us/azure/storage/storage-managed-disks-overview). ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The managed disk ID. ## Import Managed Disks can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azurestack_managed_disk.test /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.compute/disks/manageddisk1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The managed disk ID. ## Import Managed Disks can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azurestack_managed_disk.test /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.compute/disks/manageddisk1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurestack_network_interface",
			Category:         "Network Resources",
			ShortDescription: `Manages a Network Interface located in a Virtual Network, usually attached to a Virtual Machine.`,
			Description:      ``,
			Keywords: []string{
				"network",
				"interface",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the network interface. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The name of the resource group in which to create the network interface. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Required) The location/region where the network interface is created. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "enable_ip_forwarding",
					Description: `(Optional) Enables IP Forwarding on the NIC. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dns_servers",
					Description: `(Optional) List of DNS servers IP addresses to use for this NIC, overrides the VNet-level server list`,
				},
				resource.Attribute{
					Name:        "ip_configuration",
					Description: `(Required) One or more ` + "`" + `ip_configuration` + "`" + ` associated with this NIC as documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. The ` + "`" + `ip_configuration` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) User-defined name of the IP.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) Reference to a subnet in which this NIC has been created.`,
				},
				resource.Attribute{
					Name:        "private_ip_address",
					Description: `(Optional) Static IP Address.`,
				},
				resource.Attribute{
					Name:        "private_ip_address_allocation",
					Description: `(Required) Defines how a private IP address is assigned. Options are Static or Dynamic.`,
				},
				resource.Attribute{
					Name:        "public_ip_address_id",
					Description: `(Optional) Reference to a Public IP Address to associate with this NIC`,
				},
				resource.Attribute{
					Name:        "private_ip_address_version",
					Description: `(Optional) The IP Version to use. Possible values are ` + "`" + `IPv4` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `(Optional) Is this the Primary Network Interface? If set to ` + "`" + `true` + "`" + ` this should be the first ` + "`" + `ip_configuration` + "`" + ` in the array. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Virtual Network Interface ID.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `The media access control (MAC) address of the network interface.`,
				},
				resource.Attribute{
					Name:        "private_ip_address",
					Description: `The private ip address of the network interface.`,
				},
				resource.Attribute{
					Name:        "virtual_machine_id",
					Description: `Reference to a VM with which this NIC has been associated.`,
				},
				resource.Attribute{
					Name:        "applied_dns_servers",
					Description: `If the VM that uses this NIC is part of an Availability Set, then this list will have the union of all DNS servers from all NICs that are part of the Availability Set ## Import Network Interfaces can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azurestack_network_interface.test /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.network/networkInterfaces/nic1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Virtual Network Interface ID.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `The media access control (MAC) address of the network interface.`,
				},
				resource.Attribute{
					Name:        "private_ip_address",
					Description: `The private ip address of the network interface.`,
				},
				resource.Attribute{
					Name:        "virtual_machine_id",
					Description: `Reference to a VM with which this NIC has been associated.`,
				},
				resource.Attribute{
					Name:        "applied_dns_servers",
					Description: `If the VM that uses this NIC is part of an Availability Set, then this list will have the union of all DNS servers from all NICs that are part of the Availability Set ## Import Network Interfaces can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azurestack_network_interface.test /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.network/networkInterfaces/nic1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurestack_network_security_group",
			Category:         "Network Resources",
			ShortDescription: `Manages a network security group that contains a list of network security rules. Network security groups enable inbound or outbound traffic to be enabled or denied.`,
			Description:      ``,
			Keywords: []string{
				"network",
				"security",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the network security group. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The name of the resource group in which to create the network security group. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Required) Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "security_rule",
					Description: `(Optional) One or more ` + "`" + `security_rule` + "`" + ` blocks as defined below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. The ` + "`" + `security_rule` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the security rule.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description for this rule. Restricted to 140 characters.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Network protocol this rule applies to. Can be ` + "`" + `Tcp` + "`" + `, ` + "`" + `Udp` + "`" + ` or ` + "`" + ``,
				},
				resource.Attribute{
					Name:        "source_port_range",
					Description: `(Optional) Source Port or Range. Integer or range between ` + "`" + `0` + "`" + ` and ` + "`" + `65535` + "`" + ` or ` + "`" + ``,
				},
				resource.Attribute{
					Name:        "destination_port_range",
					Description: `(Optional) Destination Port or Range. Integer or range between ` + "`" + `0` + "`" + ` and ` + "`" + `65535` + "`" + ` or ` + "`" + ``,
				},
				resource.Attribute{
					Name:        "source_address_prefix",
					Description: `(Optional) CIDR or source IP range or`,
				},
				resource.Attribute{
					Name:        "destination_address_prefix",
					Description: `(Optional) CIDR or destination IP range or`,
				},
				resource.Attribute{
					Name:        "access",
					Description: `(Required) Specifies whether network traffic is allowed or denied. Possible values are ` + "`" + `Allow` + "`" + ` and ` + "`" + `Deny` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Required) Specifies the priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Required) The direction specifies if rule will be evaluated on incoming or outgoing traffic. Possible values are ` + "`" + `Inbound` + "`" + ` and ` + "`" + `Outbound` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Network Security Group ID. ## Import Network Security Groups can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azurestack_network_security_group.group1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/networkSecurityGroups/mySecurityGroup ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Network Security Group ID. ## Import Network Security Groups can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azurestack_network_security_group.group1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/networkSecurityGroups/mySecurityGroup ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurestack_network_security_rule",
			Category:         "Network Resources",
			ShortDescription: `Manages a Network Security Rule.`,
			Description:      ``,
			Keywords: []string{
				"network",
				"security",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the security rule. This needs to be unique across all Rules in the Network Security Group. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The name of the resource group in which to create the Network Security Rule. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "network_security_group_name",
					Description: `(Required) The name of the Network Security Group that we want to attach the rule to. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description for this rule. Restricted to 140 characters.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Network protocol this rule applies to. Possible values include ` + "`" + `Tcp` + "`" + `, ` + "`" + `Udp` + "`" + ` or ` + "`" + ``,
				},
				resource.Attribute{
					Name:        "source_port_range",
					Description: `(Optional) Source Port or Range. Integer or range between ` + "`" + `0` + "`" + ` and ` + "`" + `65535` + "`" + ` or ` + "`" + ``,
				},
				resource.Attribute{
					Name:        "destination_port_range",
					Description: `(Optional) Destination Port or Range. Integer or range between ` + "`" + `0` + "`" + ` and ` + "`" + `65535` + "`" + ` or ` + "`" + ``,
				},
				resource.Attribute{
					Name:        "source_address_prefix",
					Description: `(Optional) CIDR or source IP range or`,
				},
				resource.Attribute{
					Name:        "destination_address_prefix",
					Description: `(Optional) CIDR or destination IP range or`,
				},
				resource.Attribute{
					Name:        "access",
					Description: `(Required) Specifies whether network traffic is allowed or denied. Possible values are ` + "`" + `Allow` + "`" + ` and ` + "`" + `Deny` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Required) Specifies the priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Required) The direction specifies if rule will be evaluated on incoming or outgoing traffic. Possible values are ` + "`" + `Inbound` + "`" + ` and ` + "`" + `Outbound` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Network Security Rule ID. ## Import Network Security Rules can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azurestack_network_security_rule.rule1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/networkSecurityGroups/mySecurityGroup/securityRules/rule1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Network Security Rule ID. ## Import Network Security Rules can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azurestack_network_security_rule.rule1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/networkSecurityGroups/mySecurityGroup/securityRules/rule1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurestack_public_ip",
			Category:         "Network Resources",
			ShortDescription: `Manages a Public IP Address.`,
			Description:      ``,
			Keywords: []string{
				"network",
				"public",
				"ip",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the Public IP resource . Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The name of the resource group in which to create the public ip.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Required) Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "public_ip_address_allocation",
					Description: `(Optional) Defines whether the IP address is static or dynamic. Options are Static or Dynamic.`,
				},
				resource.Attribute{
					Name:        "allocation_method",
					Description: `(Optional) Defines the allocation method for this IP address. Possible values are ` + "`" + `Static` + "`" + ` or ` + "`" + `Dynamic` + "`" + `. ~>`,
				},
				resource.Attribute{
					Name:        "idle_timeout_in_minutes",
					Description: `(Optional) Specifies the timeout for the TCP idle connection. The value can be set between 4 and 30 minutes.`,
				},
				resource.Attribute{
					Name:        "domain_name_label",
					Description: `(Optional) Label for the Domain Name. Will be used to make up the FQDN. If a domain name label is specified, an A DNS record is created for the public IP in the Microsoft Azure DNS system.`,
				},
				resource.Attribute{
					Name:        "reverse_fqdn",
					Description: `(Optional) A fully qualified domain name that resolves to this public IP address. If the reverseFqdn is specified, then a PTR DNS record is created pointing from the IP address in the in-addr.arpa domain to the reverse FQDN.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Public IP ID.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The IP address value that was allocated. ~>`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `Fully qualified domain name of the A DNS record associated with the public IP. This is the concatenation of the domainNameLabel and the regionalized DNS zone ## Import Public IPs can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azurestack_public_ip.myPublicIp /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/publicIPAddresses/myPublicIpAddress1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Public IP ID.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The IP address value that was allocated. ~>`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `Fully qualified domain name of the A DNS record associated with the public IP. This is the concatenation of the domainNameLabel and the regionalized DNS zone ## Import Public IPs can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azurestack_public_ip.myPublicIp /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/publicIPAddresses/myPublicIpAddress1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurestack_resource_group",
			Category:         "Base Resources",
			ShortDescription: `Creates a new resource group on Azure.`,
			Description:      ``,
			Keywords: []string{
				"base",
				"resource",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the resource group. Must be unique on your Azure subscription.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Required) The location where the resource group should be created. For a list of all Azure locations, please consult [this link](http://azure.microsoft.com/en-us/regions/) or run ` + "`" + `az account list-locations --output table` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource group ID. ## Import Resource Groups can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azurestack_resource_group.mygroup /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource group ID. ## Import Resource Groups can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azurestack_resource_group.mygroup /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurestack_route",
			Category:         "Network Resources",
			ShortDescription: `Manages a Route within a Route Table.`,
			Description:      ``,
			Keywords: []string{
				"network",
				"route",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the route. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The name of the resource group in which to create the route. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "route_table_name",
					Description: `(Required) The name of the route table within which create the route. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "address_prefix",
					Description: `(Required) The destination CIDR to which the route applies, such as ` + "`" + `10.1.0.0/16` + "`" + ``,
				},
				resource.Attribute{
					Name:        "next_hop_type",
					Description: `(Required) The type of Azure hop the packet should be sent to. Possible values are ` + "`" + `VirtualNetworkGateway` + "`" + `, ` + "`" + `VnetLocal` + "`" + `, ` + "`" + `Internet` + "`" + `, ` + "`" + `VirtualAppliance` + "`" + ` and ` + "`" + `None` + "`" + ``,
				},
				resource.Attribute{
					Name:        "next_hop_in_ip_address",
					Description: `(Optional) Contains the IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is ` + "`" + `VirtualAppliance` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Route ID. ## Import Routes can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azurestack_route.testRoute /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/routeTables/mytable1/routes/myroute1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Route ID. ## Import Routes can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azurestack_route.testRoute /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/routeTables/mytable1/routes/myroute1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurestack_route_table",
			Category:         "Network Resources",
			ShortDescription: `Manages a Route Table`,
			Description:      ``,
			Keywords: []string{
				"network",
				"route",
				"table",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the route table. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The name of the resource group in which to create the route table. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Required) Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "route",
					Description: `(Optional) Can be specified multiple times to define multiple routes. Each ` + "`" + `route` + "`" + ` block supports fields documented below.`,
				},
				resource.Attribute{
					Name:        "disable_bgp_route_propagation",
					Description: `(Optional) Boolean flag which controls propagation of routes learned by BGP on that route table. True means disable.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. The ` + "`" + `route` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the route.`,
				},
				resource.Attribute{
					Name:        "address_prefix",
					Description: `(Required) The destination CIDR to which the route applies, such as 10.1.0.0/16`,
				},
				resource.Attribute{
					Name:        "next_hop_type",
					Description: `(Required) The type of Azure hop the packet should be sent to. Possible values are ` + "`" + `VirtualNetworkGateway` + "`" + `, ` + "`" + `VnetLocal` + "`" + `, ` + "`" + `Internet` + "`" + `, ` + "`" + `VirtualAppliance` + "`" + ` and ` + "`" + `None` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "next_hop_in_ip_address",
					Description: `(Optional) Contains the IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is ` + "`" + `VirtualAppliance` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Route Table ID.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `The collection of Subnets associated with this route table. ## Import Route Tables can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azurestack_route_table.test /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/routeTables/mytable1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Route Table ID.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `The collection of Subnets associated with this route table. ## Import Route Tables can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azurestack_route_table.test /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/routeTables/mytable1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurestack_storage_account",
			Category:         "Storage Resources",
			ShortDescription: `Manages a Azure Storage Account.`,
			Description:      ``,
			Keywords: []string{
				"storage",
				"account",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the storage account. Changing this forces a new resource to be created. This must be unique across the entire Azure service, not just within the resource group.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The name of the resource group in which to create the storage account. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Required) Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "account_kind",
					Description: `(Optional) Defines the Kind of account. Valid option is ` + "`" + `Storage` + "`" + `. . Changing this forces a new resource to be created. Defaults to ` + "`" + `Storage` + "`" + ` currently as per [Azure Stack Storage Differences](https://docs.microsoft.com/en-us/azure/azure-stack/user/azure-stack-acs-differences)`,
				},
				resource.Attribute{
					Name:        "account_tier",
					Description: `(Required) Defines the Tier to use for this storage account. Valid options are ` + "`" + `Standard` + "`" + ` and ` + "`" + `Premium` + "`" + `. Changing this forces a new resource to be created -`,
				},
				resource.Attribute{
					Name:        "account_replication_type",
					Description: `(Required) Defines the type of replication to use for this storage account. Valid option is ` + "`" + `LRS` + "`" + ` currently as per [Azure Stack Storage Differences](https://docs.microsoft.com/en-us/azure/azure-stack/user/azure-stack-acs-differences)`,
				},
				resource.Attribute{
					Name:        "account_encryption_source",
					Description: `(Optional) The Encryption Source for this Storage Account. Possible values are ` + "`" + `Microsoft.Keyvault` + "`" + ` and ` + "`" + `Microsoft.Storage` + "`" + `. Defaults to ` + "`" + `Microsoft.Storage` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "custom_domain",
					Description: `(Optional) A ` + "`" + `custom_domain` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "enable_https_traffic_only",
					Description: `(Optional) Boolean flag which forces HTTPS if enabled, see [here](https://docs.microsoft.com/en-us/azure/storage/storage-require-secure-transfer/) for more information. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource.\ ---`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The Custom Domain Name to use for the Storage Account, which will be validated by Azure.`,
				},
				resource.Attribute{
					Name:        "use_subdomain",
					Description: `(Optional) Should the Custom Domain Name be validated by using indirect CNAME validation? ~>`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The storage account Resource ID.`,
				},
				resource.Attribute{
					Name:        "primary_location",
					Description: `The primary location of the storage account.`,
				},
				resource.Attribute{
					Name:        "secondary_location",
					Description: `The secondary location of the storage account.`,
				},
				resource.Attribute{
					Name:        "primary_blob_endpoint",
					Description: `The endpoint URL for blob storage in the primary location.`,
				},
				resource.Attribute{
					Name:        "secondary_blob_endpoint",
					Description: `The endpoint URL for blob storage in the secondary location.`,
				},
				resource.Attribute{
					Name:        "primary_queue_endpoint",
					Description: `The endpoint URL for queue storage in the primary location.`,
				},
				resource.Attribute{
					Name:        "secondary_queue_endpoint",
					Description: `The endpoint URL for queue storage in the secondary location.`,
				},
				resource.Attribute{
					Name:        "primary_table_endpoint",
					Description: `The endpoint URL for table storage in the primary location.`,
				},
				resource.Attribute{
					Name:        "secondary_table_endpoint",
					Description: `The endpoint URL for table storage in the secondary location.`,
				},
				resource.Attribute{
					Name:        "primary_file_endpoint",
					Description: `The endpoint URL for file storage in the primary location.`,
				},
				resource.Attribute{
					Name:        "primary_access_key",
					Description: `The primary access key for the storage account`,
				},
				resource.Attribute{
					Name:        "secondary_access_key",
					Description: `The secondary access key for the storage account`,
				},
				resource.Attribute{
					Name:        "primary_connection_string",
					Description: `The connection string associated with the primary location`,
				},
				resource.Attribute{
					Name:        "secondary_connection_string",
					Description: `The connection string associated with the secondary location`,
				},
				resource.Attribute{
					Name:        "primary_blob_connection_string",
					Description: `The connection string associated with the primary blob location`,
				},
				resource.Attribute{
					Name:        "secondary_blob_connection_string",
					Description: `The connection string associated with the secondary blob location ## Import Storage Accounts can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azurestack_storage_account.storageAcc1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Storage/storageAccounts/myaccount ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The storage account Resource ID.`,
				},
				resource.Attribute{
					Name:        "primary_location",
					Description: `The primary location of the storage account.`,
				},
				resource.Attribute{
					Name:        "secondary_location",
					Description: `The secondary location of the storage account.`,
				},
				resource.Attribute{
					Name:        "primary_blob_endpoint",
					Description: `The endpoint URL for blob storage in the primary location.`,
				},
				resource.Attribute{
					Name:        "secondary_blob_endpoint",
					Description: `The endpoint URL for blob storage in the secondary location.`,
				},
				resource.Attribute{
					Name:        "primary_queue_endpoint",
					Description: `The endpoint URL for queue storage in the primary location.`,
				},
				resource.Attribute{
					Name:        "secondary_queue_endpoint",
					Description: `The endpoint URL for queue storage in the secondary location.`,
				},
				resource.Attribute{
					Name:        "primary_table_endpoint",
					Description: `The endpoint URL for table storage in the primary location.`,
				},
				resource.Attribute{
					Name:        "secondary_table_endpoint",
					Description: `The endpoint URL for table storage in the secondary location.`,
				},
				resource.Attribute{
					Name:        "primary_file_endpoint",
					Description: `The endpoint URL for file storage in the primary location.`,
				},
				resource.Attribute{
					Name:        "primary_access_key",
					Description: `The primary access key for the storage account`,
				},
				resource.Attribute{
					Name:        "secondary_access_key",
					Description: `The secondary access key for the storage account`,
				},
				resource.Attribute{
					Name:        "primary_connection_string",
					Description: `The connection string associated with the primary location`,
				},
				resource.Attribute{
					Name:        "secondary_connection_string",
					Description: `The connection string associated with the secondary location`,
				},
				resource.Attribute{
					Name:        "primary_blob_connection_string",
					Description: `The connection string associated with the primary blob location`,
				},
				resource.Attribute{
					Name:        "secondary_blob_connection_string",
					Description: `The connection string associated with the secondary blob location ## Import Storage Accounts can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azurestack_storage_account.storageAcc1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Storage/storageAccounts/myaccount ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurestack_storage_blob",
			Category:         "Storage Resources",
			ShortDescription: `Manages a Azure Storage Blob.`,
			Description:      ``,
			Keywords: []string{
				"storage",
				"blob",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the storage blob. Must be unique within the storage container the blob is located.`,
				},
				resource.Attribute{
					Name:        "storage_account_name",
					Description: `(Required) Specifies the storage account in which to create the storage container. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "storage_container_name",
					Description: `(Required) The name of the storage container in which this blob should be created.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type of the storage blob to be created. One of either ` + "`" + `block` + "`" + ` or ` + "`" + `page` + "`" + `. When not copying from an existing blob, this becomes required.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) Used only for ` + "`" + `page` + "`" + ` blobs to specify the size in bytes of the blob to be created. Must be a multiple of 512. Defaults to 0.`,
				},
				resource.Attribute{
					Name:        "cache_control",
					Description: `(Optional) Controls the [cache control header](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Cache-Control) content of the response when blob is requested .`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `(Optional) The content type of the storage blob. Cannot be defined if ` + "`" + `source_uri` + "`" + ` is defined. Defaults to ` + "`" + `application/octet-stream` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) An absolute path to a file on the local system. Cannot be defined if ` + "`" + `source_uri` + "`" + ` is defined.`,
				},
				resource.Attribute{
					Name:        "source_content",
					Description: `(Optional) The content for this blob which should be defined inline. This field can only be specified for Block blobs and cannot be specified if ` + "`" + `source` + "`" + ` or ` + "`" + `source_uri` + "`" + ` is specified.`,
				},
				resource.Attribute{
					Name:        "source_uri",
					Description: `(Optional) The URI of an existing blob, or a file in the Azure File service, to use as the source contents for the blob to be created. Changing this forces a new resource to be created. Cannot be defined if ` + "`" + `source` + "`" + ` is defined.`,
				},
				resource.Attribute{
					Name:        "content_md5",
					Description: `(Optional) The MD5 sum of the blob contents. Cannot be defined if ` + "`" + `source_uri` + "`" + ` is defined, or if blob type is Append or Page. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "parallelism",
					Description: `(Optional) The number of workers per CPU core to run for concurrent uploads. Defaults to ` + "`" + `8` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) A map of custom blob metadata. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The storage blob Resource ID.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `The URL of the blob`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The storage blob Resource ID.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `The URL of the blob`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurestack_storage_container",
			Category:         "Storage Resources",
			ShortDescription: `Manages a Azure Storage Container.`,
			Description:      ``,
			Keywords: []string{
				"storage",
				"container",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the storage container. Must be unique within the storage service the container is located.`,
				},
				resource.Attribute{
					Name:        "storage_account_name",
					Description: `(Required) Specifies the storage account in which to create the storage container. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "container_access_type",
					Description: `(Optional) The 'interface' for access the container provides. Can be either ` + "`" + `blob` + "`" + `, ` + "`" + `container` + "`" + ` or ` + "`" + `private` + "`" + `. Defaults to ` + "`" + `private` + "`" + `. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) A mapping of MetaData for this Container. All metadata keys should be lowercase. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The storage container Resource ID.`,
				},
				resource.Attribute{
					Name:        "has_immutability_policy",
					Description: `Is there an Immutability Policy configured on this Storage Container?`,
				},
				resource.Attribute{
					Name:        "has_legal_hold",
					Description: `Is there a Legal Hold configured on this Storage Container?`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The storage container Resource ID.`,
				},
				resource.Attribute{
					Name:        "has_immutability_policy",
					Description: `Is there an Immutability Policy configured on this Storage Container?`,
				},
				resource.Attribute{
					Name:        "has_legal_hold",
					Description: `Is there a Legal Hold configured on this Storage Container?`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurestack_subnet",
			Category:         "Network Resources",
			ShortDescription: `Manages a subnet. Subnets represent network segments within the IP space defined by the virtual network.`,
			Description:      ``,
			Keywords: []string{
				"network",
				"subnet",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the subnet. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The name of the resource group in which to create the subnet. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "virtual_network_name",
					Description: `(Required) The name of the virtual network to which to attach the subnet. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "address_prefix",
					Description: `(Required) The address prefix to use for the subnet.`,
				},
				resource.Attribute{
					Name:        "network_security_group_id",
					Description: `(Optional) The ID of the Network Security Group to associate with the subnet.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `(Optional) The ID of the Route Table to associate with the subnet. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The subnet ID.`,
				},
				resource.Attribute{
					Name:        "ip_configurations",
					Description: `The collection of IP Configurations with IPs within this subnet.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the subnet.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of the resource group in which the subnet is created in.`,
				},
				resource.Attribute{
					Name:        "virtual_network_name",
					Description: `The name of the virtual network in which the subnet is created in`,
				},
				resource.Attribute{
					Name:        "address_prefix",
					Description: `The address prefix for the subnet ## Import Subnets can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azurestack_subnet.testSubnet /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/virtualNetworks/myvnet1/subnets/mysubnet1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The subnet ID.`,
				},
				resource.Attribute{
					Name:        "ip_configurations",
					Description: `The collection of IP Configurations with IPs within this subnet.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the subnet.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of the resource group in which the subnet is created in.`,
				},
				resource.Attribute{
					Name:        "virtual_network_name",
					Description: `The name of the virtual network in which the subnet is created in`,
				},
				resource.Attribute{
					Name:        "address_prefix",
					Description: `The address prefix for the subnet ## Import Subnets can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azurestack_subnet.testSubnet /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/virtualNetworks/myvnet1/subnets/mysubnet1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurestack_template_deployment",
			Category:         "Template Resources",
			ShortDescription: `Manages a template deployment of resources.`,
			Description:      ``,
			Keywords: []string{
				"template",
				"deployment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the template deployment. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The name of the resource group in which to create the template deployment.`,
				},
				resource.Attribute{
					Name:        "deployment_mode",
					Description: `(Required) Specifies the mode that is used to deploy resources. This value could be either ` + "`" + `Incremental` + "`" + ` or ` + "`" + `Complete` + "`" + `. Note that you will almost`,
				},
				resource.Attribute{
					Name:        "template_body",
					Description: `(Optional) Specifies the JSON definition for the template. ~>`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `(Optional) Specifies the name and value pairs that define the deployment parameters for the template.`,
				},
				resource.Attribute{
					Name:        "parameters_body",
					Description: `(Optional) Specifies a valid Azure JSON parameters file that define the deployment parameters. It can contain KeyVault references ~>`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Template Deployment ID.`,
				},
				resource.Attribute{
					Name:        "outputs",
					Description: `A map of supported scalar output types returned from the deployment (currently, Azure Template Deployment outputs of type String, Int and Bool are supported, and are converted to strings - others will be ignored) and can be accessed using ` + "`" + `.outputs["name"]` + "`" + `. ## Note Terraform does not know about the individual resources created by Azure using a deployment template and therefore cannot delete these resources during a destroy. Destroying a template deployment removes the associated deployment operations, but will not delete the Azure resources created by the deployment. In order to delete these resources, the containing resource group must also be destroyed. [More information](https://docs.microsoft.com/en-us/rest/api/resources/deployments#Deployments_Delete).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Template Deployment ID.`,
				},
				resource.Attribute{
					Name:        "outputs",
					Description: `A map of supported scalar output types returned from the deployment (currently, Azure Template Deployment outputs of type String, Int and Bool are supported, and are converted to strings - others will be ignored) and can be accessed using ` + "`" + `.outputs["name"]` + "`" + `. ## Note Terraform does not know about the individual resources created by Azure using a deployment template and therefore cannot delete these resources during a destroy. Destroying a template deployment removes the associated deployment operations, but will not delete the Azure resources created by the deployment. In order to delete these resources, the containing resource group must also be destroyed. [More information](https://docs.microsoft.com/en-us/rest/api/resources/deployments#Deployments_Delete).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurestack_virtual_machine",
			Category:         "Compute Resources",
			ShortDescription: `Manages a Virtual Machine.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"virtual",
				"machine",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the virtual machine resource. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The name of the resource group in which to create the virtual machine.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Required) Specifies the supported Azure Stack Region where the resource exists. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `(Optional) A plan block as documented below.`,
				},
				resource.Attribute{
					Name:        "availability_set_id",
					Description: `(Optional) The Id of the Availability Set in which to create the virtual machine`,
				},
				resource.Attribute{
					Name:        "boot_diagnostics",
					Description: `(Optional) A boot diagnostics profile block as referenced below.`,
				},
				resource.Attribute{
					Name:        "vm_size",
					Description: `(Required) Specifies the [size of the virtual machine](https://azure.microsoft.com/en-us/documentation/articles/virtual-machines-size-specs/).`,
				},
				resource.Attribute{
					Name:        "storage_image_reference",
					Description: `(Optional) A Storage Image Reference block as documented below.`,
				},
				resource.Attribute{
					Name:        "storage_os_disk",
					Description: `(Required) A ` + "`" + `storage_os_disk` + "`" + ` block.`,
				},
				resource.Attribute{
					Name:        "storage_data_disk",
					Description: `(Optional) A list of Storage Data disk blocks as referenced below.`,
				},
				resource.Attribute{
					Name:        "delete_os_disk_on_termination",
					Description: `(Optional) Should the OS Disk be deleted when the Virtual Machine is destroyed? Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "delete_data_disks_on_termination",
					Description: `(Optional) Flag to enable deletion of storage data disk VHD blobs when the VM is deleted, defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "os_profile",
					Description: `(Optional) An OS Profile block as documented below. Required when ` + "`" + `create_option` + "`" + ` in the ` + "`" + `storage_os_disk` + "`" + ` block is set to ` + "`" + `FromImage` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "identity",
					Description: `(Optional) An identity block as documented below.`,
				},
				resource.Attribute{
					Name:        "license_type",
					Description: `(Optional, when a Windows machine) Specifies the Windows OS license type. If supplied, the only allowed values are ` + "`" + `Windows_Client` + "`" + ` and ` + "`" + `Windows_Server` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "os_profile_windows_config",
					Description: `(Required, when a Windows machine) A Windows config block as documented below.`,
				},
				resource.Attribute{
					Name:        "os_profile_linux_config",
					Description: `(Required, when a Linux machine) A Linux config block as documented below.`,
				},
				resource.Attribute{
					Name:        "os_profile_secrets",
					Description: `(Optional) A collection of Secret blocks as documented below.`,
				},
				resource.Attribute{
					Name:        "network_interface_ids",
					Description: `(Required) Specifies the list of resource IDs for the network interfaces associated with the virtual machine.`,
				},
				resource.Attribute{
					Name:        "primary_network_interface_id",
					Description: `(Optional) Specifies the resource ID for the primary network interface associated with the virtual machine.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. For more information on the different example configurations, please check out the [azure documentation](https://msdn.microsoft.com/en-us/library/mt163591.aspx#Anchor_2) ` + "`" + `Plan` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the image from the marketplace.`,
				},
				resource.Attribute{
					Name:        "publisher",
					Description: `(Required) Specifies the publisher of the image.`,
				},
				resource.Attribute{
					Name:        "product",
					Description: `(Required) Specifies the product of the image from the marketplace. ` + "`" + `boot_diagnostics` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Specifies the ID of the (custom) image to use to create the virtual machine, for example: ` + "`" + `` + "`" + `` + "`" + `hcl resource "azurestack_image" "test" { name = "test" #... } resource "azurestack_virtual_machine" "test" { name = "test" #... storage_image_reference { id = "${azurestack_image.test.id}" } #... } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "publisher",
					Description: `(Required, when not using image resource) Specifies the publisher of the image used to create the virtual machine. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "offer",
					Description: `(Required, when not using image resource) Specifies the offer of the image used to create the virtual machine. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "sku",
					Description: `(Required, when not using image resource) Specifies the SKU of the image used to create the virtual machine. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) Specifies the version of the image used to create the virtual machine. Changing this forces a new resource to be created. ` + "`" + `storage_os_disk` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the disk name.`,
				},
				resource.Attribute{
					Name:        "create_option",
					Description: `(Required) Specifies how the OS Disk should be created. Possible values are ` + "`" + `Attach` + "`" + ` (managed disks only) and ` + "`" + `FromImage` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "caching",
					Description: `(Optional) Specifies the caching requirements for the OS Disk. Possible values include ` + "`" + `None` + "`" + `, ` + "`" + `ReadOnly` + "`" + ` and ` + "`" + `ReadWrite` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "image_uri",
					Description: `(Optional) Specifies the image_uri in the form publisherName:offer:skus:version. ` + "`" + `image_uri` + "`" + ` can also specify the [VHD uri](https://azure.microsoft.com/en-us/documentation/articles/virtual-machines-linux-cli-deploy-templates/#create-a-custom-vm-image) of a custom VM image to clone. When cloning a custom disk image the ` + "`" + `os_type` + "`" + ` documented below becomes required.`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `(Optional) Specifies the Operating System on the OS Disk. Possible values are ` + "`" + `Linux` + "`" + ` and ` + "`" + `Windows` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disk_size_gb",
					Description: `(Optional) Specifies the size of the os disk in gigabytes. The following properties apply when using Managed Disks:`,
				},
				resource.Attribute{
					Name:        "managed_disk_id",
					Description: `(Optional) Specifies the ID of an existing Managed Disk which should be attached as the OS Disk of this Virtual Machine. If this is set then the ` + "`" + `create_option` + "`" + ` must be set to ` + "`" + `Attach` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "managed_disk_type",
					Description: `(Optional) Specifies the type of Managed Disk which should be created. Possible values are ` + "`" + `Standard_LRS` + "`" + ` or ` + "`" + `Premium_LRS` + "`" + `. The following properties apply when using Unmanaged Disks:`,
				},
				resource.Attribute{
					Name:        "vhd_uri",
					Description: `(Optional) Specifies the URI of the VHD file backing this Unmanaged OS Disk. Changing this forces a new resource to be created. ` + "`" + `storage_data_disk` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the data disk.`,
				},
				resource.Attribute{
					Name:        "create_option",
					Description: `(Required) Specifies how the data disk should be created. Possible values are ` + "`" + `Attach` + "`" + `, ` + "`" + `FromImage` + "`" + ` and ` + "`" + `Empty` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disk_size_gb",
					Description: `(Required) Specifies the size of the data disk in gigabytes.`,
				},
				resource.Attribute{
					Name:        "caching",
					Description: `(Optional) Specifies the caching requirements.`,
				},
				resource.Attribute{
					Name:        "lun",
					Description: `(Required) Specifies the logical unit number of the data disk. The following properties apply when using Managed Disks:`,
				},
				resource.Attribute{
					Name:        "managed_disk_type",
					Description: `(Optional) Specifies the type of managed disk to create. Possible values are either ` + "`" + `Standard_LRS` + "`" + ` or ` + "`" + `Premium_LRS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "managed_disk_id",
					Description: `(Optional) Specifies the ID of an Existing Managed Disk which should be attached to this Virtual Machine. When this field is set ` + "`" + `create_option` + "`" + ` must be set to ` + "`" + `Attach` + "`" + `. The following properties apply when using Unmanaged Disks:`,
				},
				resource.Attribute{
					Name:        "vhd_uri",
					Description: `(Optional) Specifies the URI of the VHD file backing this Unmanaged Data Disk. Changing this forces a new resource to be created. ` + "`" + `os_profile` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "computer_name",
					Description: `(Required) Specifies the name of the virtual machine.`,
				},
				resource.Attribute{
					Name:        "admin_username",
					Description: `(Required) Specifies the name of the administrator account.`,
				},
				resource.Attribute{
					Name:        "admin_password",
					Description: `(Required for Windows, Optional for Linux) Specifies the password of the administrator account.`,
				},
				resource.Attribute{
					Name:        "custom_data",
					Description: `(Optional) Specifies custom data to supply to the machine. On linux-based systems, this can be used as a cloud-init script. On other systems, this will be copied as a file on disk. Internally, Terraform will base64 encode this value before sending it to the API. The maximum length of the binary array is 65535 bytes. ~>`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Specifies the identity type of the virtual machine. The only allowable value is ` + "`" + `SystemAssigned` + "`" + `. To enable Managed Service Identity the virtual machine extension "ManagedIdentityExtensionForWindows" or "ManagedIdentityExtensionForLinux" must also be added to the virtual machine. The Principal ID can be retrieved after the virtual machine has been created, e.g. ` + "`" + `` + "`" + `` + "`" + `hcl resource "azurestack_virtual_machine" "test" { name = "test" identity = { type = "SystemAssigned" } } resource "azurestack_virtual_machine_extension" "test" { name = "test" resource_group_name = "${azurestack_resource_group.test.name}" location = "${azurestack_resource_group.test.location}" virtual_machine_name = "${azurestack_virtual_machine.test.name}" publisher = "Microsoft.ManagedIdentity" type = "ManagedIdentityExtensionForWindows" type_handler_version = "1.0" settings = <<SETTINGS { "port": 50342 } SETTINGS } output "principal_id" { value = "${lookup(azurestack_virtual_machine.test.identity[0], "principal_id")}" } ` + "`" + `` + "`" + `` + "`" + ` ` + "`" + `os_profile_windows_config` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "provision_vm_agent",
					Description: `(Optional) This value defaults to false.`,
				},
				resource.Attribute{
					Name:        "enable_automatic_upgrades",
					Description: `(Optional) This value defaults to false.`,
				},
				resource.Attribute{
					Name:        "winrm",
					Description: `(Optional) A collection of WinRM configuration blocks as documented below.`,
				},
				resource.Attribute{
					Name:        "additional_unattend_config",
					Description: `(Optional) An Additional Unattended Config block as documented below. ` + "`" + `winrm` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Specifies the protocol of listener`,
				},
				resource.Attribute{
					Name:        "certificate_url",
					Description: `(Optional) Specifies URL of the certificate with which new Virtual Machines is provisioned. ` + "`" + `additional_unattend_config` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "pass",
					Description: `(Required) Specifies the name of the pass that the content applies to. The only allowable value is ` + "`" + `oobeSystem` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "component",
					Description: `(Required) Specifies the name of the component to configure with the added content. The only allowable value is ` + "`" + `Microsoft-Windows-Shell-Setup` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "setting_name",
					Description: `(Required) Specifies the name of the setting to which the content applies. Possible values are: ` + "`" + `FirstLogonCommands` + "`" + ` and ` + "`" + `AutoLogon` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Optional) Specifies the base-64 encoded XML formatted content that is added to the unattend.xml file for the specified path and component. ` + "`" + `os_profile_linux_config` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "disable_password_authentication",
					Description: `(Required) Specifies whether password authentication should be disabled. If set to ` + "`" + `false` + "`" + `, an ` + "`" + `admin_password` + "`" + ` must be specified.`,
				},
				resource.Attribute{
					Name:        "ssh_keys",
					Description: `(Optional) Specifies a collection of ` + "`" + `path` + "`" + ` and ` + "`" + `key_data` + "`" + ` to be placed on the virtual machine. ~>`,
				},
				resource.Attribute{
					Name:        "source_vault_id",
					Description: `(Required) Specifies the key vault to use.`,
				},
				resource.Attribute{
					Name:        "vault_certificates",
					Description: `(Required) A collection of Vault Certificates as documented below ` + "`" + `vault_certificates` + "`" + ` support the following:`,
				},
				resource.Attribute{
					Name:        "certificate_url",
					Description: `(Required) Specifies the URI of the key vault secrets in the format of ` + "`" + `https://<vaultEndpoint>/secrets/<secretName>/<secretVersion>` + "`" + `. Stored secret is the Base64 encoding of a JSON Object that which is encoded in UTF-8 of which the contents need to be ` + "`" + `` + "`" + `` + "`" + `json { "data":"<Base64-encoded-certificate>", "dataType":"pfx", "password":"<pfx-file-password>" } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "certificate_store",
					Description: `(Required, on windows machines) Specifies the certificate store on the Virtual Machine where the certificate should be added to. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The virtual machine ID. ## Import Virtual Machines can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azurestack_virtual_machine.test /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.compute/virtualMachines/machine1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The virtual machine ID. ## Import Virtual Machines can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azurestack_virtual_machine.test /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.compute/virtualMachines/machine1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurestack_virtual_machine_extension",
			Category:         "Compute Resources",
			ShortDescription: `Creates a new Virtual Machine Extension to provide post deployment configuration and run automated tasks.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"virtual",
				"machine",
				"extension",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the virtual machine extension peering. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Required) The location where the extension is created. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The name of the resource group in which to create the virtual network. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "virtual_machine_name",
					Description: `(Required) The name of the virtual machine. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "publisher",
					Description: `(Required) The publisher of the extension, available publishers can be found by using the Azure CLI.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of extension, available types for a publisher can be found using the Azure CLI. ~>`,
				},
				resource.Attribute{
					Name:        "type_handler_version",
					Description: `(Required) Specifies the version of the extension to use, available versions can be found using the Azure CLI.`,
				},
				resource.Attribute{
					Name:        "auto_upgrade_minor_version",
					Description: `(Optional) Specifies if the platform deploys the latest minor version update to the ` + "`" + `type_handler_version` + "`" + ` specified.`,
				},
				resource.Attribute{
					Name:        "settings",
					Description: `(Required) The settings passed to the extension, these are specified as a JSON object in a string. ~>`,
				},
				resource.Attribute{
					Name:        "protected_settings",
					Description: `(Optional) The protected_settings passed to the extension, like settings, these are specified as a JSON object in a string. ~>`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Virtual Machine Extension ID. ## Import Virtual Machine Extensions can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azurestack_virtual_machine_extension.test /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/virtualMachines/myVM/extensions/hostname ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Virtual Machine Extension ID. ## Import Virtual Machine Extensions can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azurestack_virtual_machine_extension.test /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/virtualMachines/myVM/extensions/hostname ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurestack_virtual_machine_scale_set",
			Category:         "Compute Resources",
			ShortDescription: `Manages a Virtual Machine scale set.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"virtual",
				"machine",
				"scale",
				"set",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the virtual machine scale set resource. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The name of the resource group in which to create the virtual machine scale set. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Required) Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "sku",
					Description: `(Required) A sku block as documented below.`,
				},
				resource.Attribute{
					Name:        "upgrade_policy_mode",
					Description: `(Required) Specifies the mode of an upgrade to virtual machines in the scale set. Possible values, ` + "`" + `Manual` + "`" + ` or ` + "`" + `Automatic` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "overprovision",
					Description: `(Optional) Specifies whether the virtual machine scale set should be overprovisioned. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "license_type",
					Description: `(Optional, when a Windows machine) Specifies the Windows OS license type. If supplied, the only allowed values are ` + "`" + `Windows_Client` + "`" + ` and ` + "`" + `Windows_Server` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "os_profile",
					Description: `(Required) A Virtual Machine OS Profile block as documented below.`,
				},
				resource.Attribute{
					Name:        "os_profile_secrets",
					Description: `(Optional) A collection of Secret blocks as documented below.`,
				},
				resource.Attribute{
					Name:        "os_profile_windows_config",
					Description: `(Required, when a windows machine) A Windows config block as documented below.`,
				},
				resource.Attribute{
					Name:        "os_profile_linux_config",
					Description: `(Required, when a linux machine) A Linux config block as documented below.`,
				},
				resource.Attribute{
					Name:        "network_profile",
					Description: `(Required) A collection of network profile block as documented below.`,
				},
				resource.Attribute{
					Name:        "storage_profile_os_disk",
					Description: `(Required) A storage profile os disk block as documented below`,
				},
				resource.Attribute{
					Name:        "storage_profile_data_disk",
					Description: `(Optional) A storage profile data disk block as documented below`,
				},
				resource.Attribute{
					Name:        "storage_profile_image_reference",
					Description: `(Optional) A storage profile image reference block as documented below.`,
				},
				resource.Attribute{
					Name:        "extension",
					Description: `(Optional) Can be specified multiple times to add extension profiles to the scale set. Each ` + "`" + `extension` + "`" + ` block supports the fields documented below.`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `(Optional) A plan block as documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. ->`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the size of virtual machines in a scale set.`,
				},
				resource.Attribute{
					Name:        "tier",
					Description: `(Optional) Specifies the tier of virtual machines in a scale set. Possible values, ` + "`" + `standard` + "`" + ` or ` + "`" + `basic` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `(Required) Specifies the number of virtual machines in the scale set.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Specifies the identity type to be assigned to the scale set. The only allowable value is ` + "`" + `SystemAssigned` + "`" + `. To enable Managed Service Identity (MSI) on all machines in the scale set, an extension with the type "ManagedIdentityExtensionForWindows" or "ManagedIdentityExtensionForLinux" must also be added. The scale set's Service Principal ID (SPN) can be retrieved after the scale set has been created. ` + "`" + `` + "`" + `` + "`" + `hcl resource "azurestack_virtual_machine_scale_set" "test" { name = "vm-scaleset" resource_group_name = "${azurestack_resource_group.test.name}" location = "${azurestack_resource_group.test.location}" sku { name = "${var.vm_sku}" tier = "Standard" capacity = "${var.instance_count}" } identity { type = "systemAssigned" } extension { name = "MSILinuxExtension" publisher = "Microsoft.ManagedIdentity" type = "ManagedIdentityExtensionForLinux" type_handler_version = "1.0" settings = "{\"port\": 50342}" } output "principal_id" { value = "${lookup(azurestack_virtual_machine.test.identity[0], "principal_id")}" } } ` + "`" + `` + "`" + `` + "`" + ` ` + "`" + `os_profile` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "computer_name_prefix",
					Description: `(Required) Specifies the computer name prefix for all of the virtual machines in the scale set. Computer name prefixes must be 1 to 9 characters long for windows images and 1 - 58 for linux. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "admin_username",
					Description: `(Required) Specifies the administrator account name to use for all the instances of virtual machines in the scale set.`,
				},
				resource.Attribute{
					Name:        "admin_password",
					Description: `(Required) Specifies the administrator password to use for all the instances of virtual machines in a scale set.`,
				},
				resource.Attribute{
					Name:        "custom_data",
					Description: `(Optional) Specifies custom data to supply to the machine. On linux-based systems, this can be used as a cloud-init script. On other systems, this will be copied as a file on disk. Internally, Terraform will base64 encode this value before sending it to the API. The maximum length of the binary array is 65535 bytes. ` + "`" + `os_profile_secrets` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "source_vault_id",
					Description: `(Required) Specifies the key vault to use.`,
				},
				resource.Attribute{
					Name:        "vault_certificates",
					Description: `(Required, on windows machines) A collection of Vault Certificates as documented below ` + "`" + `vault_certificates` + "`" + ` support the following:`,
				},
				resource.Attribute{
					Name:        "certificate_url",
					Description: `(Required) It is the Base64 encoding of a JSON Object that which is encoded in UTF-8 of which the contents need to be ` + "`" + `data` + "`" + `, ` + "`" + `dataType` + "`" + ` and ` + "`" + `password` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "certificate_store",
					Description: `(Required, on windows machines) Specifies the certificate store on the Virtual Machine where the certificate should be added to. ` + "`" + `os_profile_windows_config` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "provision_vm_agent",
					Description: `(Optional) Indicates whether virtual machine agent should be provisioned on the virtual machines in the scale set.`,
				},
				resource.Attribute{
					Name:        "enable_automatic_upgrades",
					Description: `(Optional) Indicates whether virtual machines in the scale set are enabled for automatic updates.`,
				},
				resource.Attribute{
					Name:        "winrm",
					Description: `(Optional) A collection of WinRM configuration blocks as documented below.`,
				},
				resource.Attribute{
					Name:        "additional_unattend_config",
					Description: `(Optional) An Additional Unattended Config block as documented below. ` + "`" + `winrm` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Specifies the protocol of listener`,
				},
				resource.Attribute{
					Name:        "certificate_url",
					Description: `(Optional) Specifies URL of the certificate with which new Virtual Machines is provisioned. ` + "`" + `additional_unattend_config` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "pass",
					Description: `(Required) Specifies the name of the pass that the content applies to. The only allowable value is ` + "`" + `oobeSystem` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "component",
					Description: `(Required) Specifies the name of the component to configure with the added content. The only allowable value is ` + "`" + `Microsoft-Windows-Shell-Setup` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "setting_name",
					Description: `(Required) Specifies the name of the setting to which the content applies. Possible values are: ` + "`" + `FirstLogonCommands` + "`" + ` and ` + "`" + `AutoLogon` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Optional) Specifies the base-64 encoded XML formatted content that is added to the unattend.xml file for the specified path and component. ` + "`" + `os_profile_linux_config` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "disable_password_authentication",
					Description: `(Required) Specifies whether password authentication should be disabled. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "ssh_keys",
					Description: `(Optional) Specifies a collection of ` + "`" + `path` + "`" + ` and ` + "`" + `key_data` + "`" + ` to be placed on the virtual machine. ~> _`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the network interface configuration.`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `(Required) Indicates whether network interfaces created from the network interface configuration will be the primary NIC of the VM.`,
				},
				resource.Attribute{
					Name:        "ip_configuration",
					Description: `(Required) An ip_configuration block as documented below. ` + "`" + `public_ip_address_configuration` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the public ip address configuration`,
				},
				resource.Attribute{
					Name:        "idle_timeout",
					Description: `(Required) The idle timeout in minutes. This value must be between 4 and 32.`,
				},
				resource.Attribute{
					Name:        "domain_name_label",
					Description: `(Required) The domain name label for the dns settings. ` + "`" + `storage_profile_os_disk` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Specifies the disk name. Must be specified when using unmanaged disk ('managed_disk_type' property not set).`,
				},
				resource.Attribute{
					Name:        "vhd_containers",
					Description: `(Optional) Specifies the vhd uri. Cannot be used when ` + "`" + `image` + "`" + ` or ` + "`" + `managed_disk_type` + "`" + ` is specified.`,
				},
				resource.Attribute{
					Name:        "managed_disk_type",
					Description: `(Optional) Specifies the type of managed disk to create. Value you must be either ` + "`" + `Standard_LRS` + "`" + ` or ` + "`" + `Premium_LRS` + "`" + `. Cannot be used when ` + "`" + `vhd_containers` + "`" + ` or ` + "`" + `image` + "`" + ` is specified.`,
				},
				resource.Attribute{
					Name:        "create_option",
					Description: `(Required) Specifies how the virtual machine should be created. The only possible option is ` + "`" + `FromImage` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "caching",
					Description: `(Optional) Specifies the caching requirements. Possible values include: ` + "`" + `None` + "`" + ` (default), ` + "`" + `ReadOnly` + "`" + `, ` + "`" + `ReadWrite` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Optional) Specifies the blob uri for user image. A virtual machine scale set creates an os disk in the same container as the user image. Updating the osDisk image causes the existing disk to be deleted and a new one created with the new image. If the VM scale set is in Manual upgrade mode then the virtual machines are not updated until they have manualUpgrade applied to them. When setting this field ` + "`" + `os_type` + "`" + ` needs to be specified.`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `(Optional) Specifies the operating system Type, valid values are windows, linux. ` + "`" + `storage_profile_data_disk` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "lun",
					Description: `(Required) Specifies the Logical Unit Number of the disk in each virtual machine in the scale set.`,
				},
				resource.Attribute{
					Name:        "create_option",
					Description: `(Optional) Specifies how the data disk should be created. The only possible options are ` + "`" + `FromImage` + "`" + ` and ` + "`" + `Empty` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "caching",
					Description: `(Optional) Specifies the caching requirements. Possible values include: ` + "`" + `None` + "`" + ` (default), ` + "`" + `ReadOnly` + "`" + `, ` + "`" + `ReadWrite` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disk_size_gb",
					Description: `(Optional) Specifies the size of the disk in GB. This element is required when creating an empty disk. ` + "`" + `storage_profile_image_reference` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Specifies the ID of the (custom) image to use to create the virtual machine scale set, as in the [example below](#example-of-storage_profile_image_reference-with-id).`,
				},
				resource.Attribute{
					Name:        "publisher",
					Description: `(Optional) Specifies the publisher of the image used to create the virtual machines.`,
				},
				resource.Attribute{
					Name:        "offer",
					Description: `(Optional) Specifies the offer of the image used to create the virtual machines.`,
				},
				resource.Attribute{
					Name:        "sku",
					Description: `(Optional) Specifies the SKU of the image used to create the virtual machines.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) Specifies the version of the image used to create the virtual machines. ` + "`" + `boot_diagnostics` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the extension.`,
				},
				resource.Attribute{
					Name:        "publisher",
					Description: `(Required) The publisher of the extension, available publishers can be found by using the Azure CLI.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of extension, available types for a publisher can be found using the Azure CLI.`,
				},
				resource.Attribute{
					Name:        "type_handler_version",
					Description: `(Required) Specifies the version of the extension to use, available versions can be found using the Azure CLI.`,
				},
				resource.Attribute{
					Name:        "auto_upgrade_minor_version",
					Description: `(Optional) Specifies whether or not to use the latest minor version available.`,
				},
				resource.Attribute{
					Name:        "settings",
					Description: `(Required) The settings passed to the extension, these are specified as a JSON object in a string.`,
				},
				resource.Attribute{
					Name:        "protected_settings",
					Description: `(Optional) The protected_settings passed to the extension, like settings, these are specified as a JSON object in a string. ` + "`" + `plan` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the image from the marketplace.`,
				},
				resource.Attribute{
					Name:        "publisher",
					Description: `(Required) Specifies the publisher of the image.`,
				},
				resource.Attribute{
					Name:        "product",
					Description: `(Required) Specifies the product of the image from the marketplace. ## Example of storage_profile_image_reference with id ` + "`" + `` + "`" + `` + "`" + `hcl resource "azurestack_image" "test" { name = "test" #... } resource "azurestack_virtual_machine_scale_set" "test" { name = "test" #... storage_profile_image_reference { id = "${azurestack_image.test.id}" } #... } ` + "`" + `` + "`" + `` + "`" + ` ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The virtual machine scale set ID.`,
				},
				resource.Attribute{
					Name:        "boot_diagnostics",
					Description: `A boot diagnostics profile block as referenced below. ## Import Virtual Machine Scale Sets can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azurestack_virtual_machine_scale_set.scaleset1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/virtualMachineScaleSets/scaleset1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The virtual machine scale set ID.`,
				},
				resource.Attribute{
					Name:        "boot_diagnostics",
					Description: `A boot diagnostics profile block as referenced below. ## Import Virtual Machine Scale Sets can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azurestack_virtual_machine_scale_set.scaleset1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/virtualMachineScaleSets/scaleset1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurestack_virtual_network",
			Category:         "Network Resources",
			ShortDescription: `Creates a new virtual network including any configured subnets. Each subnet can optionally be configured with a security group to be associated with the subnet.`,
			Description:      ``,
			Keywords: []string{
				"network",
				"virtual",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the virtual network. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The name of the resource group in which to create the virtual network.`,
				},
				resource.Attribute{
					Name:        "address_space",
					Description: `(Required) The address space that is used the virtual network. You can supply more than one address space. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Required) The location/region where the virtual network is created. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "dns_servers",
					Description: `(Optional) List of IP addresses of DNS servers`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `(Optional) Can be specified multiple times to define multiple subnets. Each ` + "`" + `subnet` + "`" + ` block supports fields documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. The ` + "`" + `subnet` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the subnet.`,
				},
				resource.Attribute{
					Name:        "address_prefix",
					Description: `(Required) The address prefix to use for the subnet.`,
				},
				resource.Attribute{
					Name:        "security_group",
					Description: `(Optional) The Network Security Group to associate with the subnet. (Referenced by ` + "`" + `id` + "`" + `, ie. ` + "`" + `azurestack_network_security_group.test.id` + "`" + `) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The virtual NetworkConfiguration ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the virtual network.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of the resource group in which to create the virtual network.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location/region where the virtual network is created`,
				},
				resource.Attribute{
					Name:        "address_space",
					Description: `The address space that is used the virtual network. ## Import Virtual Networks can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azurestack_virtual_network.testNetwork /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/virtualNetworks/myvnet1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The virtual NetworkConfiguration ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the virtual network.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of the resource group in which to create the virtual network.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location/region where the virtual network is created`,
				},
				resource.Attribute{
					Name:        "address_space",
					Description: `The address space that is used the virtual network. ## Import Virtual Networks can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azurestack_virtual_network.testNetwork /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/virtualNetworks/myvnet1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurestack_virtual_network_gateway",
			Category:         "Network Resources",
			ShortDescription: `Manages a Virtual Network Gateway to establish secure, cross-premises connectivity.`,
			Description:      ``,
			Keywords: []string{
				"network",
				"virtual",
				"gateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the connection. Changing the name forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The name of the resource group in which to create the connection Changing the name forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Required) The location/region where the connection is located. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of the Virtual Network Gateway. Valid options is ` + "`" + `Vpn` + "`" + ``,
				},
				resource.Attribute{
					Name:        "vpn_type",
					Description: `(Optional) The routing type of the Virtual Network Gateway. Only valid option is ` + "`" + `RouteBased` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enable_bgp",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, BGP (Border Gateway Protocol) is enabled for this connection. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "bgp_settings",
					Description: `(Optional) A ` + "`" + `bgp_settings` + "`" + ` block which is documented below. In this block the BGP specific settings can be defined.`,
				},
				resource.Attribute{
					Name:        "sku",
					Description: `(Required) Configuration of the size and capacity of the virtual network gateway. Valid options are ` + "`" + `Basic` + "`" + `, ` + "`" + `Standard` + "`" + ` and ` + "`" + `HighPerformance` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "active_active",
					Description: `(Optional) If true, an active-active Virtual Network Gateway will be created. An active-active gateway requires the HighPerformance SKU. If false, an active-standby gateway will be created. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "ip_configuration",
					Description: `(Required) One or two ip_configuration blocks documented below. An active-standby gateway requires exactly one ip_configuration block whereas an active-active gateway requires exactly two ip_configuration blocks.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. The ` + "`" + `ip_configuration` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A user-defined name of the IP configuration. Defaults to vnetGatewayConfig.`,
				},
				resource.Attribute{
					Name:        "private_ip_address_allocation",
					Description: `(Optional) Defines how the private IP address of the gateways virtual interface is assigned. Valid options are Static or Dynamic. Defaults to Dynamic.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) The ID of the gateway subnet of a virtual network in which the virtual network gateway will be created. It is mandatory that the associated subnet is named ` + "`" + `GatewaySubnet` + "`" + `. Therefore, each virtual network can contain at most a single Virtual Network Gateway.`,
				},
				resource.Attribute{
					Name:        "public_ip_address_id",
					Description: `(Optional) The ID of the public ip address to associate with the Virtual Network Gateway. The ` + "`" + `vpn_client_configuration` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "address_space",
					Description: `(Required) The address space out of which ip addresses for vpn clients will be taken. You can provide more than one address space, e.g. in CIDR notation.`,
				},
				resource.Attribute{
					Name:        "root_certificate",
					Description: `(Optional) One or more ` + "`" + `root_certificate` + "`" + ` blocks which are defined below. These root certificates are used to sign the client certificate used by the VPN clients to connect to the gateway.`,
				},
				resource.Attribute{
					Name:        "revoked_certificate",
					Description: `(Optional) One or more ` + "`" + `revoked_certificate` + "`" + ` blocks which are defined below.`,
				},
				resource.Attribute{
					Name:        "radius_server_address",
					Description: `(Optional) The address of the Radius server.`,
				},
				resource.Attribute{
					Name:        "radius_server_secret",
					Description: `(Optional) The secret used by the Radius server.`,
				},
				resource.Attribute{
					Name:        "vpn_client_protocols",
					Description: `(Optional) List of the protocols supported by the vpn client. The supported values are ` + "`" + `SSTP` + "`" + `, ` + "`" + `IkeV2` + "`" + ` and ` + "`" + `OpenVPN` + "`" + `. The ` + "`" + `bgp_settings` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "asn",
					Description: `(Optional) The Autonomous System Number (ASN) to use as part of the BGP.`,
				},
				resource.Attribute{
					Name:        "peering_address",
					Description: `(Optional) The BGP peer IP address of the virtual network gateway. This address is needed to configure the created gateway as a BGP Peer on the on-premises VPN devices. The IP address must be part of the subnet of the Virtual Network Gateway. Changing this forces a new resource to be created`,
				},
				resource.Attribute{
					Name:        "peer_weight",
					Description: `(Optional) The weight added to routes which have been learned through BGP peering. Valid values can be between 0 and 100. The ` + "`" + `root_certificate` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A user-defined name of the root certificate.`,
				},
				resource.Attribute{
					Name:        "public_cert_data",
					Description: `(Required) The public certificate of the root certificate authority. The certificate must be provided in Base-64 encoded X.509 format (PEM). In particular, this argument`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A user-defined name of the revoked certificate.`,
				},
				resource.Attribute{
					Name:        "thumbprint",
					Description: `(Required) The SHA1 thumbprint of the certificate to be revoked. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Virtual Network Gateway. ## Import Virtual Network Gateways can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import azurestack_virtual_network_gateway.testGateway /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myGroup1/providers/Microsoft.Network/virtualNetworkGateways/myGateway1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Virtual Network Gateway. ## Import Virtual Network Gateways can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import azurestack_virtual_network_gateway.testGateway /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myGroup1/providers/Microsoft.Network/virtualNetworkGateways/myGateway1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurestack_virtual_network_gateway_connection",
			Category:         "Network Resources",
			ShortDescription: `Manages a connection in an existing Virtual Network Gateway.`,
			Description:      ``,
			Keywords: []string{
				"network",
				"virtual",
				"gateway",
				"connection",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the connection. Changing the name forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The name of the resource group in which to create the connection Changing the name forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Required) The location/region where the connection is located. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of connection. Valid options are ` + "`" + `IPsec` + "`" + ` (Site-to-Site), ` + "`" + `ExpressRoute` + "`" + ` (ExpressRoute). Each connection type requires different mandatory arguments (refer to the examples above). Changing the connection type will force a new connection to be created.`,
				},
				resource.Attribute{
					Name:        "virtual_network_gateway_id",
					Description: `(Required) The ID of the Virtual Network Gateway in which the connection will be created. Changing the gateway forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "authorization_key",
					Description: `(Optional) The authorization key associated with the Express Route Circuit. This field is required only if the type is an ExpressRoute connection.`,
				},
				resource.Attribute{
					Name:        "express_route_circuit_id",
					Description: `(Optional) The ID of the Express Route Circuit when creating an ExpressRoute connection (i.e. when ` + "`" + `type` + "`" + ` is ` + "`" + `ExpressRoute` + "`" + `). The Express Route Circuit can be in the same or in a different subscription.`,
				},
				resource.Attribute{
					Name:        "peer_virtual_network_gateway_id",
					Description: `(Optional) The ID of the peer virtual network gateway when creating a VNet-to-VNet connection (i.e. when ` + "`" + `type` + "`" + ` is ` + "`" + `Vnet2Vnet` + "`" + `). The peer Virtual Network Gateway can be in the same or in a different subscription.`,
				},
				resource.Attribute{
					Name:        "local_network_gateway_id",
					Description: `(Optional) The ID of the local network gateway when creating Site-to-Site connection (i.e. when ` + "`" + `type` + "`" + ` is ` + "`" + `IPsec` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "routing_weight",
					Description: `(Optional) The routing weight. Defaults to ` + "`" + `10` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "shared_key",
					Description: `(Optional) The shared IPSec key. A key must be provided if a Site-to-Site or VNet-to-VNet connection is created whereas ExpressRoute connections do not need a shared key.`,
				},
				resource.Attribute{
					Name:        "enable_bgp",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, BGP (Border Gateway Protocol) is enabled for this connection. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The connection ID. ## Import Virtual Network Gateway Connections can be imported using their ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import azurestack_virtual_network_gateway_connection.testConnection /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myGroup1/providers/Microsoft.Network/connections/myConnection1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The connection ID. ## Import Virtual Network Gateway Connections can be imported using their ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import azurestack_virtual_network_gateway_connection.testConnection /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myGroup1/providers/Microsoft.Network/connections/myConnection1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"azurestack_availability_set":                   0,
		"azurestack_dns_a_record":                       1,
		"azurestack_dns_zone":                           2,
		"azurestack_lb":                                 3,
		"azurestack_lb_backend_address_pool":            4,
		"azurestack_lb_nat_pool":                        5,
		"azurestack_lb_nat_rule":                        6,
		"azurestack_lb_probe":                           7,
		"azurestack_lb_rule":                            8,
		"azurestack_local_network_gateway":              9,
		"azurestack_managed_disk":                       10,
		"azurestack_network_interface":                  11,
		"azurestack_network_security_group":             12,
		"azurestack_network_security_rule":              13,
		"azurestack_public_ip":                          14,
		"azurestack_resource_group":                     15,
		"azurestack_route":                              16,
		"azurestack_route_table":                        17,
		"azurestack_storage_account":                    18,
		"azurestack_storage_blob":                       19,
		"azurestack_storage_container":                  20,
		"azurestack_subnet":                             21,
		"azurestack_template_deployment":                22,
		"azurestack_virtual_machine":                    23,
		"azurestack_virtual_machine_extension":          24,
		"azurestack_virtual_machine_scale_set":          25,
		"azurestack_virtual_network":                    26,
		"azurestack_virtual_network_gateway":            27,
		"azurestack_virtual_network_gateway_connection": 28,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
