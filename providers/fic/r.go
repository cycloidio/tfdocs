package fic

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "fic_eri_global_ip_address_set_v1",
			Category:         "ERI(Enterprise Reliable InterConnect) Resources",
			ShortDescription: `Manages a V1 NAT Global IP Address Set resource within Flexible InterConnect.`,
			Description:      ``,
			Keywords: []string{
				"eri",
				"enterprise",
				"reliable",
				"interconnect",
				"global",
				"ip",
				"address",
				"set",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "router_id",
					Description: `(Required) The router ID this global ip address set belongs to.`,
				},
				resource.Attribute{
					Name:        "nat_id",
					Description: `(Required) The nat component ID this global ip address set belongs to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the global ip address set.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Address type of the global ip address set. "sourceNapt" or "destinationNat" can be specified.`,
				},
				resource.Attribute{
					Name:        "number_of_addresses",
					Description: `(Required) Number of the IP addresses. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "addresses",
					Description: `Created global IP addresses.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "addresses",
					Description: `Created global IP addresses.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fic_eri_port_to_azure_microsoft_connection_v1",
			Category:         "ERI(Enterprise Reliable InterConnect) Resources",
			ShortDescription: `Manages a V1 Port to Azure Microsoft Connection resource within Flexible InterConnect.`,
			Description:      ``,
			Keywords: []string{
				"eri",
				"enterprise",
				"reliable",
				"interconnect",
				"port",
				"to",
				"azure",
				"microsoft",
				"connection",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the connection.`,
				},
				resource.Attribute{
					Name:        "source_primary_port_id",
					Description: `(Required) Primary source port's ID of the connection.`,
				},
				resource.Attribute{
					Name:        "source_primary_vlan",
					Description: `(Required) Primary source VLAN ID of the connection.`,
				},
				resource.Attribute{
					Name:        "source_secondary_port_id",
					Description: `(Required) Secondary source port's ID of the connection.`,
				},
				resource.Attribute{
					Name:        "source_secondary_vlan",
					Description: `(Required) Secondary source VLAN ID of the connection.`,
				},
				resource.Attribute{
					Name:        "source_asn",
					Description: `(Required) AS Number of the connection.`,
				},
				resource.Attribute{
					Name:        "destination_interconnect",
					Description: `(Required) Target cloud name of the connection.`,
				},
				resource.Attribute{
					Name:        "destination_qos_type",
					Description: `(Required) QOS Type of the connection. Currently only "guarantee" is supported.`,
				},
				resource.Attribute{
					Name:        "destination_service_key",
					Description: `(Required) Service key of the target cloud.`,
				},
				resource.Attribute{
					Name:        "destination_shared_key",
					Description: `(Optional) BGP MD5 key.`,
				},
				resource.Attribute{
					Name:        "destination_advertised_public_prefixes",
					Description: `(Required) Advertised Public Prefixes.`,
				},
				resource.Attribute{
					Name:        "destination_routing_registry_name",
					Description: `(Optional) Routing Registry Name. Allowed values are: "ARIN", "APNIC", "AFRINIC", "LACNIC", "RIPE", "NCC", "RADB", "ALTDB"`,
				},
				resource.Attribute{
					Name:        "primary_connected_network_address",
					Description: `(Required) Primary network address of the connection.`,
				},
				resource.Attribute{
					Name:        "secondary_connected_network_address",
					Description: `(Required) Secondary network address of the connection.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Required) Bandwidth of the connection. Allowed values are: "10M", "20M", "30M", "40M", "50M", "100M", "200M", "300M", "400M", "500M", "1G", "2G", "3G", "4G", "5G", "10G" ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "redundant",
					Description: `Redundancy of the connection.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `Tenant ID of the connection.`,
				},
				resource.Attribute{
					Name:        "area",
					Description: `Area name of the connection.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "redundant",
					Description: `Redundancy of the connection.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `Tenant ID of the connection.`,
				},
				resource.Attribute{
					Name:        "area",
					Description: `Area name of the connection.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fic_eri_port_to_azure_private_connection_v1",
			Category:         "ERI(Enterprise Reliable InterConnect) Resources",
			ShortDescription: `Manages a V1 Port to Azure Private Connection resource within Flexible InterConnect.`,
			Description:      ``,
			Keywords: []string{
				"eri",
				"enterprise",
				"reliable",
				"interconnect",
				"port",
				"to",
				"azure",
				"private",
				"connection",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the connection.`,
				},
				resource.Attribute{
					Name:        "source_primary_port_id",
					Description: `(Required) Primary source port's ID of the connection.`,
				},
				resource.Attribute{
					Name:        "source_primary_vlan",
					Description: `(Required) Primary source VLAN ID of the connection.`,
				},
				resource.Attribute{
					Name:        "source_secondary_port_id",
					Description: `(Required) Secondary source port's ID of the connection.`,
				},
				resource.Attribute{
					Name:        "source_secondary_vlan",
					Description: `(Required) Secondary source VLAN ID of the connection.`,
				},
				resource.Attribute{
					Name:        "source_asn",
					Description: `(Required) AS Number of the connection.`,
				},
				resource.Attribute{
					Name:        "destination_interconnect",
					Description: `(Required) Target cloud name of the connection.`,
				},
				resource.Attribute{
					Name:        "destination_qos_type",
					Description: `(Required) QOS Type of the connection. Currently only "guarantee" is supported.`,
				},
				resource.Attribute{
					Name:        "destination_service_key",
					Description: `(Required) Service key of the target cloud.`,
				},
				resource.Attribute{
					Name:        "destination_shared_key",
					Description: `(Optional) BGP MD5 key.`,
				},
				resource.Attribute{
					Name:        "primary_connected_network_address",
					Description: `(Required) Primary network address of the connection.`,
				},
				resource.Attribute{
					Name:        "secondary_connected_network_address",
					Description: `(Required) Secondary network address of the connection.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Required) Bandwidth of the connection. Allowed values are: "10M", "20M", "30M", "40M", "50M", "100M", "200M", "300M", "400M", "500M", "1G", "2G", "3G", "4G", "5G", "10G" ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "redundant",
					Description: `Redundancy of the connection.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `Tenant ID of the connection.`,
				},
				resource.Attribute{
					Name:        "area",
					Description: `Area name of the connection.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "redundant",
					Description: `Redundancy of the connection.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `Tenant ID of the connection.`,
				},
				resource.Attribute{
					Name:        "area",
					Description: `Area name of the connection.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fic_eri_port_to_port_connection_v1",
			Category:         "ERI(Enterprise Reliable InterConnect) Resources",
			ShortDescription: `Manages a V1 Port to Port Connection resource within Flexible InterConnect.`,
			Description:      ``,
			Keywords: []string{
				"eri",
				"enterprise",
				"reliable",
				"interconnect",
				"port",
				"to",
				"connection",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the resource.`,
				},
				resource.Attribute{
					Name:        "source_port_id",
					Description: `(Required) Source port ID of the connection.`,
				},
				resource.Attribute{
					Name:        "source_vlan",
					Description: `(Required) Source VLAN ID of the connection.`,
				},
				resource.Attribute{
					Name:        "destination_port_id",
					Description: `(Required) Destination port ID of the connection.`,
				},
				resource.Attribute{
					Name:        "destination_vlan",
					Description: `(Required) Destination VLAN ID of the connection.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Optional) Bandwidth of the connection. Allowed values are "10M", "20M", "30M", "40M", "50M", "100M", "200M", "300M", "400M", "500M", "1G", "2G", "3G", "4G", "5G" and "10G" . ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "redundant",
					Description: `Redundancy of the connection.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `Tenant ID of the connection.`,
				},
				resource.Attribute{
					Name:        "area",
					Description: `Area name of the connection.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "redundant",
					Description: `Redundancy of the connection.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `Tenant ID of the connection.`,
				},
				resource.Attribute{
					Name:        "area",
					Description: `Area name of the connection.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fic_eri_port_v1",
			Category:         "ERI(Enterprise Reliable InterConnect) Resources",
			ShortDescription: `Manages a V1 Port resource within Flexible InterConnect.`,
			Description:      ``,
			Keywords: []string{
				"eri",
				"enterprise",
				"reliable",
				"interconnect",
				"port",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the resource.`,
				},
				resource.Attribute{
					Name:        "switch_name",
					Description: `(Required) Switch name you create a port.`,
				},
				resource.Attribute{
					Name:        "number_of_vlans",
					Description: `(Optional; Required if ` + "`" + `vlan_ranges` + "`" + ` is empty) The number of VLANs used by port.`,
				},
				resource.Attribute{
					Name:        "vlan_ranges",
					Description: `(Optional; Required if ` + "`" + `number_of_vlans` + "`" + ` is empty) The list of VLAN ranges object.`,
				},
				resource.Attribute{
					Name:        "port_type",
					Description: `(Optional) Type of port either "1G" or "10G".`,
				},
				resource.Attribute{
					Name:        "is_acivated",
					Description: `(Optional) Activate status of the port. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "switch_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "is_activated",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `Tenant ID the port belongs to.`,
				},
				resource.Attribute{
					Name:        "area",
					Description: `Area name the port belongs to.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Location name the port belongs to.`,
				},
				resource.Attribute{
					Name:        "vlans/vid",
					Description: `VLAN ID of the router.`,
				},
				resource.Attribute{
					Name:        "vlans/status",
					Description: `VLAN status of the port.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "switch_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "is_activated",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `Tenant ID the port belongs to.`,
				},
				resource.Attribute{
					Name:        "area",
					Description: `Area name the port belongs to.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Location name the port belongs to.`,
				},
				resource.Attribute{
					Name:        "vlans/vid",
					Description: `VLAN ID of the router.`,
				},
				resource.Attribute{
					Name:        "vlans/status",
					Description: `VLAN status of the port.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fic_eri_router_paired_to_gcp_connection_v1",
			Category:         "ERI(Enterprise Reliable InterConnect) Resources",
			ShortDescription: `Manages a V1 Router(Paired) to GCP Connection resource within Flexible InterConnect.`,
			Description:      ``,
			Keywords: []string{
				"eri",
				"enterprise",
				"reliable",
				"interconnect",
				"router",
				"paired",
				"to",
				"gcp",
				"connection",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the connection. It must be less than 64 characters in half-width alphanumeric characters and some symbols &()-_.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Required) Bandwidth of the connection. Either "10M", "50M", "100M", "200M", "300M", "400M", "500M", "1G", "2G", "5G" or "10G".`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Required) Source of the connection. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `(Required) Destination of the connection. Structure is documented below. The ` + "`" + `source` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `(Required) Router ID. It must be a F + 12-digit number.`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `(Required) Group name. Either "group_1", "group_2", "group_3", "group_4", "group_5", "group_6", "group_7" or "group_8".`,
				},
				resource.Attribute{
					Name:        "route_fileter",
					Description: `(Required) Route filter. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "primary_med_out",
					Description: `(Required) MED egress value of primary. Either 10 or 30. The ` + "`" + `route_filter` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "in",
					Description: `(Required) BGP filter ingress value. Either "fullRoute" or "noRoute".`,
				},
				resource.Attribute{
					Name:        "out",
					Description: `(Required) BGP filter egress value. Either "fullRoute", "fullRouteWithDefaultRoute", "defaultRoute", "privateRoute" or "noRoute". The ` + "`" + `destination` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `(Required) Primary interconnect of destination.`,
				},
				resource.Attribute{
					Name:        "secondary",
					Description: `(Required) Secondary interconnect of destination. The ` + "`" + `primary` + "`" + ` and ` + "`" + `secondary` + "`" + ` blocks support:`,
				},
				resource.Attribute{
					Name:        "interconnect",
					Description: `(Required) Connecting point. See "1.3. FIC-Connection Google Cloud" in [FIC-Connection Google Cloud](https://sdpf.ntt.com/services/docs/fic/service-descriptions/connection-gcp/connection-gcp.html#id4).`,
				},
				resource.Attribute{
					Name:        "pairing_key",
					Description: `(Required) Paring key of google hybrid interconnect. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "source.0.secondary_med_out",
					Description: `MED egress value of secondary. It would be source.primary_med_out plus 10.`,
				},
				resource.Attribute{
					Name:        "destination.0.qos_type",
					Description: `QoS type. It would be "guarantee".`,
				},
				resource.Attribute{
					Name:        "redundant",
					Description: `Redundant flag of the connection. It would be true.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `Tenant ID where the connection belongs.`,
				},
				resource.Attribute{
					Name:        "area",
					Description: `Area name of the connection.`,
				},
				resource.Attribute{
					Name:        "operation_id",
					Description: `ID of the last operation.`,
				},
				resource.Attribute{
					Name:        "operation_status",
					Description: `Status of the last operation.`,
				},
				resource.Attribute{
					Name:        "primary_connected_network_address",
					Description: `Primary connected network address. It would be "<network_address>/29".`,
				},
				resource.Attribute{
					Name:        "secondary_connected_network_address",
					Description: `Secondary connected network address. It would be "<network_address>/29". ## Timeouts This resource provides the following Timeout configuration options: - ` + "`" + `create` + "`" + ` - Default is 10 minutes. - ` + "`" + `update` + "`" + ` - Default is 10 minutes. - ` + "`" + `delete` + "`" + ` - Default is 10 minutes. ## Import Connections can be imported using the ID: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fic_eri_router_paired_to_gcp_connection_v1.connection F030123456789 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "source.0.secondary_med_out",
					Description: `MED egress value of secondary. It would be source.primary_med_out plus 10.`,
				},
				resource.Attribute{
					Name:        "destination.0.qos_type",
					Description: `QoS type. It would be "guarantee".`,
				},
				resource.Attribute{
					Name:        "redundant",
					Description: `Redundant flag of the connection. It would be true.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `Tenant ID where the connection belongs.`,
				},
				resource.Attribute{
					Name:        "area",
					Description: `Area name of the connection.`,
				},
				resource.Attribute{
					Name:        "operation_id",
					Description: `ID of the last operation.`,
				},
				resource.Attribute{
					Name:        "operation_status",
					Description: `Status of the last operation.`,
				},
				resource.Attribute{
					Name:        "primary_connected_network_address",
					Description: `Primary connected network address. It would be "<network_address>/29".`,
				},
				resource.Attribute{
					Name:        "secondary_connected_network_address",
					Description: `Secondary connected network address. It would be "<network_address>/29". ## Timeouts This resource provides the following Timeout configuration options: - ` + "`" + `create` + "`" + ` - Default is 10 minutes. - ` + "`" + `update` + "`" + ` - Default is 10 minutes. - ` + "`" + `delete` + "`" + ` - Default is 10 minutes. ## Import Connections can be imported using the ID: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fic_eri_router_paired_to_gcp_connection_v1.connection F030123456789 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fic_eri_router_to_azure_microsoft_connection_v1",
			Category:         "ERI(Enterprise Reliable InterConnect) Resources",
			ShortDescription: `Manages a V1 Router to Azure Microsoft Connection resource within Flexible InterConnect.`,
			Description:      ``,
			Keywords: []string{
				"eri",
				"enterprise",
				"reliable",
				"interconnect",
				"router",
				"to",
				"azure",
				"microsoft",
				"connection",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the connection.`,
				},
				resource.Attribute{
					Name:        "source_router_id",
					Description: `(Required) Source router ID of the connection.`,
				},
				resource.Attribute{
					Name:        "source_group_name",
					Description: `(Required) Source group name of the connection. Allowed values are: "group_1", "group_2", "group_3", "group_4", "group_5", "group_6", "group_7" and "group_8"`,
				},
				resource.Attribute{
					Name:        "source_route_filter_in",
					Description: `(Required) Ingress value of BGP Filter. Allowed values are: "fullRoute", "noRoute"`,
				},
				resource.Attribute{
					Name:        "source_route_filter_out",
					Description: `(Required) Egress value of BGP Filter. Allowed values are: "natRoute", "noRoute"`,
				},
				resource.Attribute{
					Name:        "destination_interconnect",
					Description: `(Required) Target cloud name of the connection.`,
				},
				resource.Attribute{
					Name:        "destination_qos_type",
					Description: `(Required) QOS Type of the connection. Currently only "guarantee" is supported.`,
				},
				resource.Attribute{
					Name:        "destination_service_key",
					Description: `(Required) Service key of the target cloud.`,
				},
				resource.Attribute{
					Name:        "destination_advertised_public_prefixes",
					Description: `(Required) Advertised Public Prefixes.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Required) Bandwidth of the connection. Allowed values are: "10M", "20M", "30M", "40M", "50M", "100M", "200M", "300M", "400M", "500M", "1G", "2G", "3G", "4G", "5G", "10G" ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "primary_connected_network_address",
					Description: `Primary network address of the connection.`,
				},
				resource.Attribute{
					Name:        "secondary_connected_network_address",
					Description: `Secondary network address of the connection.`,
				},
				resource.Attribute{
					Name:        "redundant",
					Description: `Redundancy of the connection.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `Tenant ID of the connection.`,
				},
				resource.Attribute{
					Name:        "area",
					Description: `Area name of the connection.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "primary_connected_network_address",
					Description: `Primary network address of the connection.`,
				},
				resource.Attribute{
					Name:        "secondary_connected_network_address",
					Description: `Secondary network address of the connection.`,
				},
				resource.Attribute{
					Name:        "redundant",
					Description: `Redundancy of the connection.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `Tenant ID of the connection.`,
				},
				resource.Attribute{
					Name:        "area",
					Description: `Area name of the connection.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fic_eri_router_to_azure_private_connection_v1",
			Category:         "ERI(Enterprise Reliable InterConnect) Resources",
			ShortDescription: `Manages a V1 Router to Azure Private Connection resource within Flexible InterConnect.`,
			Description:      ``,
			Keywords: []string{
				"eri",
				"enterprise",
				"reliable",
				"interconnect",
				"router",
				"to",
				"azure",
				"private",
				"connection",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the connection.`,
				},
				resource.Attribute{
					Name:        "source_router_id",
					Description: `(Required) Source router ID of the connection.`,
				},
				resource.Attribute{
					Name:        "source_group_name",
					Description: `(Required) Source group name of the connection. Allowed values are: "group_1", "group_2", "group_3", "group_4", "group_5", "group_6", "group_7" and "group_8"`,
				},
				resource.Attribute{
					Name:        "source_route_filter_in",
					Description: `(Required) Ingress value of BGP Filter. Allowed values are: "fullRoute", "noRoute"`,
				},
				resource.Attribute{
					Name:        "source_route_filter_out",
					Description: `(Required) Egress value of BGP Filter. Allowed values are: "fullRoute", "fullRouteWithDefaultRoute", "defaultRoute", "privateRoute", "noRoute"`,
				},
				resource.Attribute{
					Name:        "destination_interconnect",
					Description: `(Required) Target cloud name of the connection.`,
				},
				resource.Attribute{
					Name:        "destination_qos_type",
					Description: `(Required) QOS Type of the connection. Currently only "guarantee" is supported.`,
				},
				resource.Attribute{
					Name:        "destination_service_key",
					Description: `(Required) Service key of the target cloud.`,
				},
				resource.Attribute{
					Name:        "primary_connected_network_address",
					Description: `(Required) Primary network address of the connection.`,
				},
				resource.Attribute{
					Name:        "secondary_connected_network_address",
					Description: `(Required) Secondary network address of the connection.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Required) Bandwidth of the connection. Allowed values are: "10M", "20M", "30M", "40M", "50M", "100M", "200M", "300M", "400M", "500M", "1G", "2G", "3G", "4G", "5G", "10G" ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "redundant",
					Description: `Redundancy of the connection.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `Tenant ID of the connection.`,
				},
				resource.Attribute{
					Name:        "area",
					Description: `Area name of the connection.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "redundant",
					Description: `Redundancy of the connection.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `Tenant ID of the connection.`,
				},
				resource.Attribute{
					Name:        "area",
					Description: `Area name of the connection.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fic_eri_router_to_ecl_connection_v1",
			Category:         "ERI(Enterprise Reliable InterConnect) Resources",
			ShortDescription: `Manages a V1 Router to ECL Connection resource within Flexible InterConnect.`,
			Description:      ``,
			Keywords: []string{
				"eri",
				"enterprise",
				"reliable",
				"interconnect",
				"router",
				"to",
				"ecl",
				"connection",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the resource.`,
				},
				resource.Attribute{
					Name:        "source_router_id",
					Description: `(Required) Source router ID of the connection.`,
				},
				resource.Attribute{
					Name:        "source_group_name",
					Description: `(Required) Source group name of the connection. Allowed values are "group_1", "group_2", "group_3", "group_4", "group_5", "group_6", "group_7" and "group_8".`,
				},
				resource.Attribute{
					Name:        "source_route_filter_in",
					Description: `(Required) Ingress value of BGP Filter. Either "fullRoute" or "noRoute" is allowed.`,
				},
				resource.Attribute{
					Name:        "source_route_filter_out",
					Description: `(Required) Egress value of BGP Filter. Allowed values are "fullRoute", "noRoute" and "fullRouteWithDefaultRoute" .`,
				},
				resource.Attribute{
					Name:        "destination_interconnect",
					Description: `(Required) Target cloud of the connection. See "1.3. FIC-Connection SDPF クラウド/サーバー" in [FIC-Connection SDPF クラウド/サーバー](https://sdpf.ntt.com/services/docs/fic/service-descriptions/connection-ecl/connection-ecl.html#id4)).`,
				},
				resource.Attribute{
					Name:        "destination_c_number",
					Description: `(Required) Destination C number of the connection.`,
				},
				resource.Attribute{
					Name:        "destination_parent_contract_number",
					Description: `(Required) Destination parent contract number of the connection.`,
				},
				resource.Attribute{
					Name:        "destination_vpn_number",
					Description: `(Required) Destination VPN number of the connection.`,
				},
				resource.Attribute{
					Name:        "destination_qos_type",
					Description: `(Required) QOS Type of the connection. Currently only "guarantee" is supported.`,
				},
				resource.Attribute{
					Name:        "source_route_filter_out",
					Description: `(Required) Egress value of BGP Filter. Allowed values are "fullRoute", "fullRouteWithDefaultRoute", "defaultRoute" and "privateRoute".`,
				},
				resource.Attribute{
					Name:        "connected_network_address",
					Description: `(Required) Network address of the connection.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Optional) Bandwidth of the connection. Allowed values are "10M", "20M", "30M", "40M", "50M", "100M", "200M", "300M", "400M", "500M" and "1G" . ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "destination_contract_number",
					Description: `Destination contract number of the connection.`,
				},
				resource.Attribute{
					Name:        "redundant",
					Description: `Redundancy of the connection.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `Tenant ID of the connection.`,
				},
				resource.Attribute{
					Name:        "area",
					Description: `Area name of the connection.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "destination_contract_number",
					Description: `Destination contract number of the connection.`,
				},
				resource.Attribute{
					Name:        "redundant",
					Description: `Redundancy of the connection.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `Tenant ID of the connection.`,
				},
				resource.Attribute{
					Name:        "area",
					Description: `Area name of the connection.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fic_eri_router_to_uno_connection_v1",
			Category:         "ERI(Enterprise Reliable InterConnect) Resources",
			ShortDescription: `Manages a V1 Router to ECL Connection resource within Flexible InterConnect.`,
			Description:      ``,
			Keywords: []string{
				"eri",
				"enterprise",
				"reliable",
				"interconnect",
				"router",
				"to",
				"uno",
				"connection",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the resource.`,
				},
				resource.Attribute{
					Name:        "source_router_id",
					Description: `(Required) Source router ID of the connection.`,
				},
				resource.Attribute{
					Name:        "source_group_name",
					Description: `(Required) Source group name of the connection. Allowed values are "group_1", "group_2", "group_3", "group_4", "group_5", "group_6", "group_7" and "group_8".`,
				},
				resource.Attribute{
					Name:        "source_route_filter_in",
					Description: `(Required) Source ingress value of BGP Filter. Either "fullRoute" or "noRoute" is allowed.`,
				},
				resource.Attribute{
					Name:        "source_route_filter_out",
					Description: `(Required) Source egress value of BGP Filter. Allowed values are "fullRoute", "fullRouteWithDefaultRoute" and "noRoute".`,
				},
				resource.Attribute{
					Name:        "destination_interconnect",
					Description: `(Required) Target cloud of the connection. See "1.3. FIC-Connection Arcstar Universal One" in [FIC-Connection Arcstar Universal One](https://sdpf.ntt.com/services/docs/fic/service-descriptions/connection-uno/connection-uno.html#id4).`,
				},
				resource.Attribute{
					Name:        "destination_c_number",
					Description: `(Optional) Destination C number of the connection.`,
				},
				resource.Attribute{
					Name:        "destination_parent_contract_number",
					Description: `(Required) Destination parent contract number of the connection.`,
				},
				resource.Attribute{
					Name:        "destination_vpn_number",
					Description: `(Required) Destination VPN number of the connection.`,
				},
				resource.Attribute{
					Name:        "destination_qos_type",
					Description: `(Required) QOS Type of the connection. Currently only "guarantee" is supported.`,
				},
				resource.Attribute{
					Name:        "destination_route_filter_out",
					Description: `(Required) Destination egress value of BGP Filter. Allowed values are "fullRoute", "fullRouteWithDefaultRoute", "defaultRoute" and "privateRoute".`,
				},
				resource.Attribute{
					Name:        "connected_network_address",
					Description: `(Required) Network address of the connection.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Required) Bandwidth of the connection. Allowed values are "10M", "20M", "30M", "40M", "50M", "100M", "200M", "300M", "400M", "500M" and "1G" . ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "destination_contract_number",
					Description: `Destination contract number of the connection.`,
				},
				resource.Attribute{
					Name:        "redundant",
					Description: `Redundancy of the connection.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `Tenant ID of the connection.`,
				},
				resource.Attribute{
					Name:        "area",
					Description: `Area name of the connection.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "destination_contract_number",
					Description: `Destination contract number of the connection.`,
				},
				resource.Attribute{
					Name:        "redundant",
					Description: `Redundancy of the connection.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `Tenant ID of the connection.`,
				},
				resource.Attribute{
					Name:        "area",
					Description: `Area name of the connection.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fic_eri_router_v1",
			Category:         "ERI(Enterprise Reliable InterConnect) Resources",
			ShortDescription: `Manages a V1 Router resource within Flexible InterConnect.`,
			Description:      ``,
			Keywords: []string{
				"eri",
				"enterprise",
				"reliable",
				"interconnect",
				"router",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the resource.`,
				},
				resource.Attribute{
					Name:        "area",
					Description: `(Required) Area name you create a router.`,
				},
				resource.Attribute{
					Name:        "user_ip_address",
					Description: `(Required; Required if ` + "`" + `vlan_ranges` + "`" + ` is empy) The IP Address of the router. It must have prefix 27.`,
				},
				resource.Attribute{
					Name:        "redundant",
					Description: `(Required) The redundant option of the router. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "area",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "user_ip_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "redundant",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `Tenant ID the router belongs to.`,
				},
				resource.Attribute{
					Name:        "firewalls/id",
					Description: `Firewall ID.`,
				},
				resource.Attribute{
					Name:        "firewalls/is_activated",
					Description: `Activate status of the Firewall.`,
				},
				resource.Attribute{
					Name:        "nats/id",
					Description: `NAT component ID.`,
				},
				resource.Attribute{
					Name:        "nats/is_activated",
					Description: `Activate status of the NAT.`,
				},
				resource.Attribute{
					Name:        "routing_groups/name",
					Description: `Routing group name of the router.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "area",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "user_ip_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "redundant",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `Tenant ID the router belongs to.`,
				},
				resource.Attribute{
					Name:        "firewalls/id",
					Description: `Firewall ID.`,
				},
				resource.Attribute{
					Name:        "firewalls/is_activated",
					Description: `Activate status of the Firewall.`,
				},
				resource.Attribute{
					Name:        "nats/id",
					Description: `NAT component ID.`,
				},
				resource.Attribute{
					Name:        "nats/is_activated",
					Description: `Activate status of the NAT.`,
				},
				resource.Attribute{
					Name:        "routing_groups/name",
					Description: `Routing group name of the router.`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"fic_eri_global_ip_address_set_v1":                0,
		"fic_eri_port_to_azure_microsoft_connection_v1":   1,
		"fic_eri_port_to_azure_private_connection_v1":     2,
		"fic_eri_port_to_port_connection_v1":              3,
		"fic_eri_port_v1":                                 4,
		"fic_eri_router_paired_to_gcp_connection_v1":      5,
		"fic_eri_router_to_azure_microsoft_connection_v1": 6,
		"fic_eri_router_to_azure_private_connection_v1":   7,
		"fic_eri_router_to_ecl_connection_v1":             8,
		"fic_eri_router_to_uno_connection_v1":             9,
		"fic_eri_router_v1":                               10,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
