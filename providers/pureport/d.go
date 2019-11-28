package pureport

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "pureport_accounts",
			Category:         "Data Sources",
			ShortDescription: `Provides details about existing Pureport accounts.`,
			Description: `\_account

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A filter used to scope the list e.g. by tags.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the filter. The valid values are defined in the [Pureport SDK Model](https://github.com/pureport/pureport-sdk-go/blob/develop/docs/client/Account.md). Nested values are supported. E.g.("Location.DisplayName")`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) The value of the filter. Currently only regex strings are supported. ## Attributes The Pureport Account resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "accounts",
					Description: `The found list of accounts.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier for the Pureport account.`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `The unique path reference to the Pureport account. This will be used by other resources to identify the account in most cases.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name on the account.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the account.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A dictionary of user defined key/value pairs associated with this resource. The Pureport Guide, []()`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pureport_aws_connection",
			Category:         "Data Sources",
			ShortDescription: `Manages a Pureport AWS Direct Connect Connection.`,
			Description: `\_aws\_connection

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "connection_id",
					Description: `(Required) The ID of the connection. You should use the ` + "`" + `pureport_connections` + "`" + ` data source for querying the list of available connections and discover the ID for the connection. - - -`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A filter used to scope the list e.g. by tags.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the filter. The valid values are defined in the [Pureport SDK Model](https://github.com/pureport/pureport-sdk-go/blob/develop/docs/client/AwsDirectConnectConnection.md). Nested values are supported. E.g.("Location.DisplayName")`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) The value of the filter. Currently only regex strings are supported. ## Attributes`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name for the connection`,
				},
				resource.Attribute{
					Name:        "location_href",
					Description: `HREF for the Pureport Location to attach the connection.`,
				},
				resource.Attribute{
					Name:        "network_href",
					Description: `HREF for the network to associate the connection.`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `The maximum QoS for this connection. Valid values are 50, 100, 200, 300, 400, 500, 1000, 10000 in Mbps.`,
				},
				resource.Attribute{
					Name:        "aws_account_id",
					Description: `Your AWS Account ID.`,
				},
				resource.Attribute{
					Name:        "aws_region",
					Description: `The AWS region to create your connection.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description for the connection.`,
				},
				resource.Attribute{
					Name:        "customer_networks",
					Description: `A list of named CIDR block to easily identify a customer network.`,
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
					Name:        "billing_term",
					Description: `The billing term for the connection: (Currently only HOURLY is supported.)`,
				},
				resource.Attribute{
					Name:        "high_availability",
					Description: `Whether a redundant gateway is/should be provisioned for this connection.`,
				},
				resource.Attribute{
					Name:        "peering_type",
					Description: `The peering type to to use for the connection:`,
				},
				resource.Attribute{
					Name:        "cloud_service_hrefs",
					Description: `When PUBLIC peering is configured, a list of HREFs for the Public peering services to which we want access.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A dictionary of user defined key/value pairs to associate with this resource.`,
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
			Category:         "Data Sources",
			ShortDescription: `Manages a Pureport Azure Express Route Connection.`,
			Description: `\_azure\_connection

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "connection_id",
					Description: `(Required) The ID of the connection. You should use the ` + "`" + `pureport_connections` + "`" + ` data source for querying the list of available connections and discover the ID for the connection. - - -`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A filter used to scope the list e.g. by tags.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the filter. The valid values are defined in the [Pureport SDK Model](https://github.com/pureport/pureport-sdk-go/blob/develop/docs/client/AzureExpressRouteConnection.md). Nested values are supported. E.g.("Location.DisplayName")`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) The value of the filter. Currently only regex strings are supported. ## Attributes`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name for the connection`,
				},
				resource.Attribute{
					Name:        "location_href",
					Description: `HREF for the Pureport Location to attach the connection.`,
				},
				resource.Attribute{
					Name:        "network_href",
					Description: `HREF for the network to associate the connection.`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `The maximum QoS for this connection. Valid values are 50, 100, 200, 300, 400, 500, 1000, 10000 in Mbps.`,
				},
				resource.Attribute{
					Name:        "service_key",
					Description: `The Azure service key for the Express Route Circuit.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description for the connection.`,
				},
				resource.Attribute{
					Name:        "customer_networks",
					Description: `A list of named CIDR block to easily identify a customer network.`,
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
					Name:        "billing_term",
					Description: `The billing term for the connection: (Currently only HOURLY is supported.)`,
				},
				resource.Attribute{
					Name:        "high_availability",
					Description: `Whether a redundant gateway is/should be provisioned for this connection.`,
				},
				resource.Attribute{
					Name:        "peering_type",
					Description: `The peering type to to use for the connection:`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A dictionary of user defined key/value pairs to associate with this resource.`,
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
			Type:             "pureport_cloud_regions",
			Category:         "Data Sources",
			ShortDescription: `Provides details about existing Pureport cloud regions.`,
			Description: `\_cloud\_regions

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A filter used to scope the list e.g. by tags.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the filter. The valid values are defined in the [Pureport SDK Model](https://github.com/pureport/pureport-sdk-go/blob/develop/docs/client/CloudRegion.md).`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) The value of the filter. Currently only regex strings are supported. ## Attributes`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `The found list of regions.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier for the cloud region.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The display name for the cloud region.`,
				},
				resource.Attribute{
					Name:        "provider",
					Description: `The cloud provider for the cloud region.`,
				},
				resource.Attribute{
					Name:        "identifier",
					Description: `The identifier provided by the cloud provider for this region.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A dictionary of user defined key/value pairs associated with this resource. The Pureport Guide, []()`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pureport_cloud_services",
			Category:         "Data Sources",
			ShortDescription: `Provides details about existing Pureport cloud services.`,
			Description: `\_cloud\_services

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A filter used to scope the list e.g. by tags.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the filter. The valid values are defined in the [Pureport SDK Model](https://github.com/pureport/pureport-sdk-go/blob/develop/docs/client/CloudService.md).`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) The value of the filter. Currently only regex strings are supported. ## Attributes`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `The found list of cloud provider services.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier for the cloud service.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The display name for the cloud service.`,
				},
				resource.Attribute{
					Name:        "provider",
					Description: `The cloud provider for the cloud service.`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `The unique path reference to the cloud service. This will be used by other resources to identify the service in most cases.`,
				},
				resource.Attribute{
					Name:        "ipv4_prefix_count",
					Description: `The number of IPv4 prefixes associated with this cloud service.`,
				},
				resource.Attribute{
					Name:        "ipv6_prefix_count",
					Description: `The number of IPv6 prefixes associated with this cloud service.`,
				},
				resource.Attribute{
					Name:        "cloud_region_id",
					Description: `The identifier for the cloud service where this service is located.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A dictionary of user defined key/value pairs associated with this resource. The Pureport Guide, []()`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pureport_google_cloud_connection",
			Category:         "Data Sources",
			ShortDescription: `Manages a Pureport Google Cloud Connection.`,
			Description: `\_google\_cloud\_connection

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "connection_id",
					Description: `(Required) The ID of the connection. You should use the ` + "`" + `pureport_connections` + "`" + ` data source for querying the list of available connections and discover the ID for the connection. - - -`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A filter used to scope the list e.g. by tags.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the filter. The valid values are defined in the [Pureport SDK Model](https://github.com/pureport/pureport-sdk-go/blob/develop/docs/client/GoogleCloudInterconnectConnection.md).`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) The value of the filter. Currently only regex strings are supported. ## Attributes`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name for the connection`,
				},
				resource.Attribute{
					Name:        "location_href",
					Description: `HREF for the Pureport Location to attach the connection.`,
				},
				resource.Attribute{
					Name:        "network_href",
					Description: `HREF for the network to associate the connection.`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `The maximum QoS for this connection. Valid values are 50, 100, 200, 300, 400, 500, 1000, 10000 in Mbps.`,
				},
				resource.Attribute{
					Name:        "primary_pairing_key",
					Description: `The pairing key for the primary Google Cloud Interconnect Attachment.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description for the connection.`,
				},
				resource.Attribute{
					Name:        "customer_networks",
					Description: `A list of named CIDR block to easily identify a customer network.`,
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
					Name:        "billing_term",
					Description: `The billing term for the connection: (Currently only HOURLY is supported.)`,
				},
				resource.Attribute{
					Name:        "high_availability",
					Description: `Whether a redundant gateway is/should be provisioned for this connection.`,
				},
				resource.Attribute{
					Name:        "secodary_pairing_key",
					Description: `If HA is enabled, the pairing key for the backup Google Cloud Interconnect Attachment.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A dictionary of user defined key/value pairs to associate with this resource.`,
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
			Type:             "pureport_locations",
			Category:         "Data Sources",
			ShortDescription: `Provides details about existing Pureport locations.`,
			Description: `\_locations

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A filter used to scope the list e.g. by tags.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the filter. The valid values are defined in the [Pureport SDK Model](https://github.com/pureport/pureport-sdk-go/blob/develop/docs/client/Location.md).`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) The value of the filter. Currently only regex strings are supported. ## Attributes`,
				},
				resource.Attribute{
					Name:        "locations",
					Description: `A list of Pureport locations.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier for the Pureport locations.`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `The unique path reference for the Pureport locations. This will be used by other resources to identify the locations in most cases.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the location.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `The available links to other Pureport locations.`,
				},
				resource.Attribute{
					Name:        "location_href",
					Description: `The href of the linked location.`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `The connection speed between the locations.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A dictionary of user defined key/value pairs associated with this resource. The Pureport Guide, []()`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pureport_networks",
			Category:         "Data Sources",
			ShortDescription: `Provides details about existing Pureport networks.`,
			Description: `\_networks

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_href",
					Description: `(Required) The HREF for the Pureport account associated with this network. - - -`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A filter used to scope the list e.g. by tags.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the filter. The valid values are defined in the [Pureport SDK Model](https://github.com/pureport/pureport-sdk-go/blob/develop/docs/client/Network.md).`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) The value of the filter. Currently only regex strings are supported. ## Attributes`,
				},
				resource.Attribute{
					Name:        "networks",
					Description: `A list of Pureport networks.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier for the Pureport network.`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `The unique path reference for the Pureport network. This will be used by other resources to identify the locations in most cases.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the network.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description for the network.`,
				},
				resource.Attribute{
					Name:        "account_href",
					Description: `The HREF for the Pureport account associated with this network.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A dictionary of user defined key/value pairs associated with this resource. The Pureport Guide, []()`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pureport_site_vpn_connection",
			Category:         "Data Sources",
			ShortDescription: `Manages a Pureport Site VPN Connection.`,
			Description: `\_site\_vpn\_connection

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "connection_id",
					Description: `(Required) The ID of the connection. You should use the ` + "`" + `pureport_connections` + "`" + ` data source for querying the list of available connections and discover the ID for the connection. - - -`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A filter used to scope the list e.g. by tags.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the filter. The valid values are defined in the [Pureport SDK Model](https://github.com/pureport/pureport-sdk-go/blob/develop/docs/client/SiteIpSecVpnConnection.md).`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) The value of the filter. Currently only regex strings are supported. ## Attributes`,
				},
				resource.Attribute{
					Name:        "auth_type",
					Description: `The Authentication Type to use. (Currently only ` + "`" + `PSK` + "`" + ` is supported.)`,
				},
				resource.Attribute{
					Name:        "enable_bgp_password",
					Description: `Enable BGP password authentication. (Default: false)`,
				},
				resource.Attribute{
					Name:        "ike_version",
					Description: `the IKE Version to use. Valid values are ` + "`" + `V1` + "`" + `, ` + "`" + `V2` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ike_config",
					Description: `IKE Configuration to use:`,
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
					Description: ``,
				},
				resource.Attribute{
					Name:        "primary_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "routing_type",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secondary_customer_router_ip",
					Description: ``,
				},
				resource.Attribute{
					Name:        "secondary_key",
					Description: ``,
				},
				resource.Attribute{
					Name:        "traffic_selectors",
					Description: `List of Traffic Selectors for Route Based VPN`,
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
					Description: `The name for the connection`,
				},
				resource.Attribute{
					Name:        "location_href",
					Description: `HREF for the Pureport Location to attach the connection.`,
				},
				resource.Attribute{
					Name:        "network_href",
					Description: `HREF for the network to associate the connection.`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `The maximum QoS for this connection. Valid values are 50, 100, 200, 300, 400, 500, 1000, 10000 in Mbps.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description for the connection.`,
				},
				resource.Attribute{
					Name:        "customer_networks",
					Description: `A list of named CIDR block to easily identify a customer network.`,
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
					Name:        "billing_term",
					Description: `The billing term for the connection: (Currently only HOURLY is supported.)`,
				},
				resource.Attribute{
					Name:        "high_availability",
					Description: `Whether a redundant gateway is/should be provisioned for this connection.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A dictionary of user defined key/value pairs to associate with this resource.`,
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

	dataSourcesMap = map[string]int{

		"pureport_accounts":                0,
		"pureport_aws_connection":          1,
		"pureport_azure_connection":        2,
		"pureport_cloud_regions":           3,
		"pureport_cloud_services":          4,
		"pureport_google_cloud_connection": 5,
		"pureport_locations":               6,
		"pureport_networks":                7,
		"pureport_site_vpn_connection":     8,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
