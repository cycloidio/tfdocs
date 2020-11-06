package pureport

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "pureport_aws_connection",
			Category:         "Resources",
			ShortDescription: `Manages a Pureport AWS Direct Connect Connection.`,
			Description: `\_aws\_connection

`,
			Keywords: []string{
				"aws",
				"connection",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name for the connection`,
				},
				resource.Attribute{
					Name:        "location_href",
					Description: `(Required) HREF for the Pureport Location to attach the connection.`,
				},
				resource.Attribute{
					Name:        "network_href",
					Description: `(Required) HREF for the network to associate the connection.`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `(Required) The maximum QoS for this connection. Valid values are 50, 100, 200, 300, 400, 500, 1000, 10000 in Mbps.`,
				},
				resource.Attribute{
					Name:        "aws_account_id",
					Description: `(Required) Your AWS Account ID.`,
				},
				resource.Attribute{
					Name:        "aws_region",
					Description: `(Required) The AWS region to create your connection. - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description for the connection.`,
				},
				resource.Attribute{
					Name:        "customer_asn",
					Description: `(Optional) The BGP ASN number to use for the customer network`,
				},
				resource.Attribute{
					Name:        "customer_networks",
					Description: `(Optional) A list of named CIDR block to easily identify a customer network.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name for the network.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The CIDR block for the network`,
				},
				resource.Attribute{
					Name:        "nat_config",
					Description: `(Optional) The Network Address Translation configuration for the connection.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Is NAT enabled for this connection.`,
				},
				resource.Attribute{
					Name:        "mappings",
					Description: `(Optional) List of NAT mapped CIDR address`,
				},
				resource.Attribute{
					Name:        "native_cidr",
					Description: `(Required) The native CIDR block to map.`,
				},
				resource.Attribute{
					Name:        "billing_term",
					Description: `(Optional) The billing term for the connection: (Currently only HOURLY is supported.)`,
				},
				resource.Attribute{
					Name:        "high_availability",
					Description: `(Optional) Whether a redundant gateway is/should be provisioned for this connection.`,
				},
				resource.Attribute{
					Name:        "peering_type",
					Description: `(Optional) The peering type to to use for the connection:`,
				},
				resource.Attribute{
					Name:        "cloud_service_hrefs",
					Description: `(Optional) When PUBLIC peering is configured, a list of HREFs for the Public peering services to which we want access.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A dictionary of user defined key/value pairs to associate with this resource. ## Attributes`,
				},
				resource.Attribute{
					Name:        "nat_config",
					Description: `The Network Address Translation configuration for the connection.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Is NAT enabled for this connection.`,
				},
				resource.Attribute{
					Name:        "mappings",
					Description: `List of NAT mapped CIDR address`,
				},
				resource.Attribute{
					Name:        "native_cidr",
					Description: `The native CIDR block to map.`,
				},
				resource.Attribute{
					Name:        "nat_cidr",
					Description: `The CIDR block use for NAT to the associated subnet.`,
				},
				resource.Attribute{
					Name:        "blocks",
					Description: `List of reserved blocks for NAT.`,
				},
				resource.Attribute{
					Name:        "pnat_cidr",
					Description: `CIDR use for PNAT between connections.`,
				},
				resource.Attribute{
					Name:        "gateways",
					Description: `List of cloud gateways and their configurations.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the cloud gateway.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the cloud gateway.`,
				},
				resource.Attribute{
					Name:        "availability_domain",
					Description: `The availability domain of the cloud gateway. The valid values are ` + "`" + `PRIMARY` + "`" + `, ` + "`" + `SECONDARY` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "customer_asn",
					Description: `The customer ASN used for BGP Peering.`,
				},
				resource.Attribute{
					Name:        "customer_ip",
					Description: `The assigned IP address to the customer side of the BGP Config.`,
				},
				resource.Attribute{
					Name:        "pureport_asn",
					Description: `The Pureport ASN used for BGP Peering.`,
				},
				resource.Attribute{
					Name:        "pureport_ip",
					Description: `The assigned IP address to the Pureport side of the BGP Config.`,
				},
				resource.Attribute{
					Name:        "bgp_password",
					Description: `The autogenerated BGP password used for authentication.`,
				},
				resource.Attribute{
					Name:        "peering_subnet",
					Description: `The BGP Config subnet assigned to establish BGP peering.`,
				},
				resource.Attribute{
					Name:        "public_nat_ip",
					Description: `The public facing IP Address for NAT used by this connection.`,
				},
				resource.Attribute{
					Name:        "remote_id",
					Description: `The ID of the AWS Direct Connect Connection.`,
				},
				resource.Attribute{
					Name:        "vlan",
					Description: `The VLAN id for the connection to cloud services. The Pureport Guide, []()`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pureport_azure_connection",
			Category:         "Resources",
			ShortDescription: `Manages a Pureport Azure Express Route Connection.`,
			Description: `\_azure\_connection

`,
			Keywords: []string{
				"azure",
				"connection",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name for the connection`,
				},
				resource.Attribute{
					Name:        "location_href",
					Description: `(Required) HREF for the Pureport Location to attach the connection.`,
				},
				resource.Attribute{
					Name:        "network_href",
					Description: `(Required) HREF for the network to associate the connection.`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `(Required) The maximum QoS for this connection. Valid values are 50, 100, 200, 300, 400, 500, 1000, 10000 in Mbps.`,
				},
				resource.Attribute{
					Name:        "service_key",
					Description: `(Required) The Azure service key for the Express Route Circuit. - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description for the connection.`,
				},
				resource.Attribute{
					Name:        "customer_networks",
					Description: `(Optional) A list of named CIDR block to easily identify a customer network.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name for the network.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The CIDR block for the network`,
				},
				resource.Attribute{
					Name:        "nat_config",
					Description: `(Optional) The Network Address Translation configuration for the connection.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Is NAT enabled for this connection.`,
				},
				resource.Attribute{
					Name:        "mappings",
					Description: `(Optional) List of NAT mapped CIDR address`,
				},
				resource.Attribute{
					Name:        "native_cidr",
					Description: `(Required) The native CIDR block to map.`,
				},
				resource.Attribute{
					Name:        "billing_term",
					Description: `(Optional) The billing term for the connection: (Currently only HOURLY is supported.)`,
				},
				resource.Attribute{
					Name:        "high_availability",
					Description: `(Optional) Whether a redundant gateway is/should be provisioned for this connection.`,
				},
				resource.Attribute{
					Name:        "peering_type",
					Description: `(Optional) The peering type to to use for the connection:`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A dictionary of user defined key/value pairs to associate with this resource. ## Attributes`,
				},
				resource.Attribute{
					Name:        "nat_config",
					Description: `The Network Address Translation configuration for the connection.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Is NAT enabled for this connection.`,
				},
				resource.Attribute{
					Name:        "mappings",
					Description: `List of NAT mapped CIDR address`,
				},
				resource.Attribute{
					Name:        "native_cidr",
					Description: `(Required) The native CIDR block to map.`,
				},
				resource.Attribute{
					Name:        "nat_cidr",
					Description: `The CIDR block use for NAT to the associated subnet.`,
				},
				resource.Attribute{
					Name:        "blocks",
					Description: `List of reserved blocks for NAT.`,
				},
				resource.Attribute{
					Name:        "pnat_cidr",
					Description: `CIDR use for PNAT between connections.`,
				},
				resource.Attribute{
					Name:        "gateways",
					Description: `List of cloud gateways and their configurations.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the cloud gateway.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the cloud gateway.`,
				},
				resource.Attribute{
					Name:        "availability_domain",
					Description: `The availability domain of the cloud gateway. The valid values are ` + "`" + `PRIMARY` + "`" + `, ` + "`" + `SECONDARY` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "customer_asn",
					Description: `The customer ASN used for BGP Peering.`,
				},
				resource.Attribute{
					Name:        "customer_ip",
					Description: `The assigned IP address to the customer side of the BGP Config.`,
				},
				resource.Attribute{
					Name:        "pureport_asn",
					Description: `The Pureport ASN used for BGP Peering.`,
				},
				resource.Attribute{
					Name:        "pureport_ip",
					Description: `The assigned IP address to the Pureport side of the BGP Config.`,
				},
				resource.Attribute{
					Name:        "bgp_password",
					Description: `The autogenerated BGP password used for authentication.`,
				},
				resource.Attribute{
					Name:        "peering_subnet",
					Description: `The BGP Config subnet assigned to establish BGP peering.`,
				},
				resource.Attribute{
					Name:        "public_nat_ip",
					Description: `The public facing IP Address for NAT used by this connection.`,
				},
				resource.Attribute{
					Name:        "remote_id",
					Description: `The ID of the Azure Express Route.`,
				},
				resource.Attribute{
					Name:        "vlan",
					Description: `The VLAN id for the connection to cloud services. The Pureport Guide, []()`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pureport_google_cloud_connection",
			Category:         "Resources",
			ShortDescription: `Manages a Pureport Google Cloud Connection.`,
			Description: `\_google\_cloud\_connection

`,
			Keywords: []string{
				"google",
				"cloud",
				"connection",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name for the connection`,
				},
				resource.Attribute{
					Name:        "location_href",
					Description: `(Required) HREF for the Pureport Location to attach the connection.`,
				},
				resource.Attribute{
					Name:        "network_href",
					Description: `(Required) HREF for the network to associate the connection.`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `(Required) The maximum QoS for this connection. Valid values are 50, 100, 200, 300, 400, 500, 1000, 10000 in Mbps.`,
				},
				resource.Attribute{
					Name:        "primary_pairing_key",
					Description: `(Required) The pairing key for the primary Google Cloud Interconnect Attachment. - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description for the connection.`,
				},
				resource.Attribute{
					Name:        "customer_networks",
					Description: `(Optional) A list of named CIDR block to easily identify a customer network.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name for the network.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The CIDR block for the network`,
				},
				resource.Attribute{
					Name:        "nat_config",
					Description: `(Optional) The Network Address Translation configuration for the connection.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Is NAT enabled for this connection.`,
				},
				resource.Attribute{
					Name:        "mappings",
					Description: `(Optional) List of NAT mapped CIDR address`,
				},
				resource.Attribute{
					Name:        "native_cidr",
					Description: `(Required) The native CIDR block to map.`,
				},
				resource.Attribute{
					Name:        "billing_term",
					Description: `(Optional) The billing term for the connection: (Currently only HOURLY is supported.)`,
				},
				resource.Attribute{
					Name:        "high_availability",
					Description: `(Optional) Whether a redundant gateway is/should be provisioned for this connection.`,
				},
				resource.Attribute{
					Name:        "secodary_pairing_key",
					Description: `(Optional) If HA is enabled, the pairing key for the backup Google Cloud Interconnect Attachment.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A dictionary of user defined key/value pairs to associate with this resource. ## Attributes`,
				},
				resource.Attribute{
					Name:        "nat_config",
					Description: `The Network Address Translation configuration for the connection.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Is NAT enabled for this connection.`,
				},
				resource.Attribute{
					Name:        "mappings",
					Description: `List of NAT mapped CIDR address`,
				},
				resource.Attribute{
					Name:        "native_cidr",
					Description: `(Required) The native CIDR block to map.`,
				},
				resource.Attribute{
					Name:        "nat_cidr",
					Description: `The CIDR block use for NAT to the associated subnet.`,
				},
				resource.Attribute{
					Name:        "blocks",
					Description: `List of reserved blocks for NAT.`,
				},
				resource.Attribute{
					Name:        "pnat_cidr",
					Description: `CIDR use for PNAT between connections.`,
				},
				resource.Attribute{
					Name:        "gateways",
					Description: `List of cloud gateways and their configurations.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the cloud gateway.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the cloud gateway.`,
				},
				resource.Attribute{
					Name:        "availability_domain",
					Description: `The availability domain of the cloud gateway. The valid values are ` + "`" + `PRIMARY` + "`" + `, ` + "`" + `SECONDARY` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "customer_asn",
					Description: `The customer ASN used for BGP Peering.`,
				},
				resource.Attribute{
					Name:        "customer_ip",
					Description: `The assigned IP address to the customer side of the BGP Config.`,
				},
				resource.Attribute{
					Name:        "pureport_asn",
					Description: `The Pureport ASN used for BGP Peering.`,
				},
				resource.Attribute{
					Name:        "pureport_ip",
					Description: `The assigned IP address to the Pureport side of the BGP Config.`,
				},
				resource.Attribute{
					Name:        "bgp_password",
					Description: `The autogenerated BGP password used for authentication.`,
				},
				resource.Attribute{
					Name:        "peering_subnet",
					Description: `The BGP Config subnet assigned to establish BGP peering.`,
				},
				resource.Attribute{
					Name:        "public_nat_ip",
					Description: `N/A`,
				},
				resource.Attribute{
					Name:        "remote_id",
					Description: `The ID of the Google Cloud Interconnect.`,
				},
				resource.Attribute{
					Name:        "remote_id",
					Description: `The ID of the Google Cloud Interconnect.`,
				},
				resource.Attribute{
					Name:        "vlan",
					Description: `The VLAN id for the connection to cloud services. The Pureport Guide, []()`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pureport_network",
			Category:         "Resources",
			ShortDescription: `Manages a Pureport Network.`,
			Description: `\_network

`,
			Keywords: []string{
				"network",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name used for the Network.`,
				},
				resource.Attribute{
					Name:        "account_href",
					Description: `(Required) HREF for the Account associated with the Network. - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description for the Network.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A dictionary of user defined key/value pairs to associate with this resource. ## Attributes`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `The HREF to reference this Network. The Pureport Guide, []()`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pureport_site_vpn_connection",
			Category:         "Resources",
			ShortDescription: `Manages a Pureport Site VPN Connection.`,
			Description: `\_site\_vpn\_connection

`,
			Keywords: []string{
				"site",
				"vpn",
				"connection",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "auth_type",
					Description: `(Optional) The Authentication Type to use. (Currently only ` + "`" + `PSK` + "`" + ` is supported.)`,
				},
				resource.Attribute{
					Name:        "enable_bgp_password",
					Description: `(Optional) Enable BGP password authentication. (Default: false)`,
				},
				resource.Attribute{
					Name:        "ike_version",
					Description: `(Required) the IKE Version to use. Valid values are ` + "`" + `V1` + "`" + `, ` + "`" + `V2` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ike_config",
					Description: `(Optional) IKE Configuration to use:`,
				},
				resource.Attribute{
					Name:        "esp",
					Description: `Encapsulating Security Payload`,
				},
				resource.Attribute{
					Name:        "dh_group",
					Description: `Diffie-Hellman Group`,
				},
				resource.Attribute{
					Name:        "encryption",
					Description: `Encryption Algorithm`,
				},
				resource.Attribute{
					Name:        "integrity",
					Description: `Integrity Algorithm`,
				},
				resource.Attribute{
					Name:        "ike",
					Description: `Internet Key Exchange`,
				},
				resource.Attribute{
					Name:        "dh_group",
					Description: `Diffie-Hellman Group`,
				},
				resource.Attribute{
					Name:        "encryption",
					Description: `Encryption Algorithm`,
				},
				resource.Attribute{
					Name:        "integrity",
					Description: `Integrity Algorithm`,
				},
				resource.Attribute{
					Name:        "prf",
					Description: `Pseudo Random Function`,
				},
				resource.Attribute{
					Name:        "primary_customer_router_ip",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "primary_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "routing_type",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "secondary_customer_router_ip",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "secondary_key",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "traffic_selectors",
					Description: `(Optional) List of Traffic Selectors for Route Based VPN`,
				},
				resource.Attribute{
					Name:        "customer_side",
					Description: `The customer side CIDR block`,
				},
				resource.Attribute{
					Name:        "pureport_side",
					Description: `The Pureport side CIDR block`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name for the connection`,
				},
				resource.Attribute{
					Name:        "location_href",
					Description: `(Required) HREF for the Pureport Location to attach the connection.`,
				},
				resource.Attribute{
					Name:        "network_href",
					Description: `(Required) HREF for the network to associate the connection.`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `(Required) The maximum QoS for this connection. Valid values are 50, 100, 200, 300, 400, 500, 1000, 10000 in Mbps. - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description for the connection.`,
				},
				resource.Attribute{
					Name:        "customer_networks",
					Description: `(Optional) A list of named CIDR block to easily identify a customer network.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name for the network.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The CIDR block for the network`,
				},
				resource.Attribute{
					Name:        "nat_config",
					Description: `(Optional) The Network Address Translation configuration for the connection.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Is NAT enabled for this connection.`,
				},
				resource.Attribute{
					Name:        "mappings",
					Description: `(Optional) List of NAT mapped CIDR address`,
				},
				resource.Attribute{
					Name:        "native_cidr",
					Description: `(Required) The native CIDR block to map.`,
				},
				resource.Attribute{
					Name:        "billing_term",
					Description: `(Optional) The billing term for the connection: (Currently only HOURLY is supported.)`,
				},
				resource.Attribute{
					Name:        "high_availability",
					Description: `(Optional) Whether a redundant gateway is/should be provisioned for this connection.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A dictionary of user defined key/value pairs to associate with this resource. ## Attributes`,
				},
				resource.Attribute{
					Name:        "nat_config",
					Description: `The Network Address Translation configuration for the connection.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Is NAT enabled for this connection.`,
				},
				resource.Attribute{
					Name:        "mappings",
					Description: `List of NAT mapped CIDR address`,
				},
				resource.Attribute{
					Name:        "native_cidr",
					Description: `(Required) The native CIDR block to map.`,
				},
				resource.Attribute{
					Name:        "nat_cidr",
					Description: `The CIDR block use for NAT to the associated subnet.`,
				},
				resource.Attribute{
					Name:        "blocks",
					Description: `List of reserved blocks for NAT.`,
				},
				resource.Attribute{
					Name:        "pnat_cidr",
					Description: `CIDR use for PNAT between connections.`,
				},
				resource.Attribute{
					Name:        "gateways",
					Description: `List of cloud gateways and their configurations.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the cloud gateway.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the cloud gateway.`,
				},
				resource.Attribute{
					Name:        "availability_domain",
					Description: `The availability domain of the cloud gateway. The valid values are ` + "`" + `PRIMARY` + "`" + `, ` + "`" + `SECONDARY` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "customer_asn",
					Description: `The customer ASN used for BGP Peering.`,
				},
				resource.Attribute{
					Name:        "customer_ip",
					Description: `The assigned IP address to the customer side of the BGP Config.`,
				},
				resource.Attribute{
					Name:        "pureport_asn",
					Description: `The Pureport ASN used for BGP Peering.`,
				},
				resource.Attribute{
					Name:        "pureport_ip",
					Description: `The assigned IP address to the Pureport side of the BGP Config.`,
				},
				resource.Attribute{
					Name:        "bgp_password",
					Description: `The autogenerated BGP password used for authentication.`,
				},
				resource.Attribute{
					Name:        "peering_subnet",
					Description: `The BGP Config subnet assigned to establish BGP peering.`,
				},
				resource.Attribute{
					Name:        "public_nat_ip",
					Description: `The public facing IP Address for NAT used by this connection.`,
				},
				resource.Attribute{
					Name:        "customer_gateway_ip",
					Description: `The public IP address of the customers VPN equipment.`,
				},
				resource.Attribute{
					Name:        "customer_vti_ip",
					Description: `The assigned IP address to the customer side of the VTI tunnel.`,
				},
				resource.Attribute{
					Name:        "pureport_gateway_ip",
					Description: `The public IP address of the Pureport VPN gateway.`,
				},
				resource.Attribute{
					Name:        "pureport_vti_ip",
					Description: `The assigned IP address to the Pureport side of the VPN VTI tunnel.`,
				},
				resource.Attribute{
					Name:        "vpn_auth_type",
					Description: `The type of authentication used for the VPN Connection.`,
				},
				resource.Attribute{
					Name:        "vpn_auth_key",
					Description: `The Authentication Key used for the VPN Connection. The Pureport Guide, []()`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"pureport_aws_connection":          0,
		"pureport_azure_connection":        1,
		"pureport_google_cloud_connection": 2,
		"pureport_network":                 3,
		"pureport_site_vpn_connection":     4,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
