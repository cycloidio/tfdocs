package azurestack

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "azurestack_network_interface",
			Category:         "Data Sources",
			ShortDescription: `Get information about the specified Network Interface.`,
			Description: `

Use this data source to access the properties of an Azure Network Interface.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the Network Interface.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) Specifies the name of the resource group the Network Interface is located in. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "applied_dns_servers",
					Description: `List of DNS servers applied to the specified network interface.`,
				},
				resource.Attribute{
					Name:        "dns_servers",
					Description: `The list of DNS servers used by the specified network interface.`,
				},
				resource.Attribute{
					Name:        "enable_ip_forwarding",
					Description: `Indicate if IP forwarding is set on the specified network interface.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the virtual network that the specified network interface is associated to.`,
				},
				resource.Attribute{
					Name:        "internal_dns_name_label",
					Description: `The internal dns name label of the specified network interface.`,
				},
				resource.Attribute{
					Name:        "internal_fqdn",
					Description: `The internal FQDN associated to the specified network interface.`,
				},
				resource.Attribute{
					Name:        "ip_configuration",
					Description: `The list of IP configurations associated to the specified network interface.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location of the specified network interface.`,
				},
				resource.Attribute{
					Name:        "network_security_group_id",
					Description: `The ID of the network security group associated to the specified network interface.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `The MAC address used by the specified network interface.`,
				},
				resource.Attribute{
					Name:        "private_ip_address",
					Description: `The primary private ip address associated to the specified network interface.`,
				},
				resource.Attribute{
					Name:        "private_ip_addresses",
					Description: `The list of private ip addresses associates to the specified network interface.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `List the tags assocatied to the specified network interface.`,
				},
				resource.Attribute{
					Name:        "virtual_machine_id",
					Description: `The ID of the virtual machine that the specified network interface is attached to.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "applied_dns_servers",
					Description: `List of DNS servers applied to the specified network interface.`,
				},
				resource.Attribute{
					Name:        "dns_servers",
					Description: `The list of DNS servers used by the specified network interface.`,
				},
				resource.Attribute{
					Name:        "enable_ip_forwarding",
					Description: `Indicate if IP forwarding is set on the specified network interface.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the virtual network that the specified network interface is associated to.`,
				},
				resource.Attribute{
					Name:        "internal_dns_name_label",
					Description: `The internal dns name label of the specified network interface.`,
				},
				resource.Attribute{
					Name:        "internal_fqdn",
					Description: `The internal FQDN associated to the specified network interface.`,
				},
				resource.Attribute{
					Name:        "ip_configuration",
					Description: `The list of IP configurations associated to the specified network interface.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location of the specified network interface.`,
				},
				resource.Attribute{
					Name:        "network_security_group_id",
					Description: `The ID of the network security group associated to the specified network interface.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `The MAC address used by the specified network interface.`,
				},
				resource.Attribute{
					Name:        "private_ip_address",
					Description: `The primary private ip address associated to the specified network interface.`,
				},
				resource.Attribute{
					Name:        "private_ip_addresses",
					Description: `The list of private ip addresses associates to the specified network interface.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `List the tags assocatied to the specified network interface.`,
				},
				resource.Attribute{
					Name:        "virtual_machine_id",
					Description: `The ID of the virtual machine that the specified network interface is attached to.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurestack_network_security_group",
			Category:         "Data Sources",
			ShortDescription: `Get information about the specified Network Security Group.`,
			Description: `

Use this data source to access the properties of a Network Security Group.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the Name of the Network Security Group.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) Specifies the Name of the Resource Group within which the Network Security Group exists ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Network Security Group.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The supported Azure location where the resource exists.`,
				},
				resource.Attribute{
					Name:        "security_rule",
					Description: `One or more ` + "`" + `security_rule` + "`" + ` blocks as defined below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource. The ` + "`" + `security_rule` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the security rule.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description for this rule.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The network protocol this rule applies to.`,
				},
				resource.Attribute{
					Name:        "source_port_range",
					Description: `The Source Port or Range.`,
				},
				resource.Attribute{
					Name:        "destination_port_range",
					Description: `The Destination Port or Range.`,
				},
				resource.Attribute{
					Name:        "source_address_prefix",
					Description: `CIDR or source IP range or`,
				},
				resource.Attribute{
					Name:        "destination_address_prefix",
					Description: `CIDR or destination IP range or`,
				},
				resource.Attribute{
					Name:        "access",
					Description: `Is network traffic is allowed or denied?`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority of the rule`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `The direction specifies if rule will be evaluated on incoming or outgoing traffic.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Network Security Group.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The supported Azure location where the resource exists.`,
				},
				resource.Attribute{
					Name:        "security_rule",
					Description: `One or more ` + "`" + `security_rule` + "`" + ` blocks as defined below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource. The ` + "`" + `security_rule` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the security rule.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description for this rule.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The network protocol this rule applies to.`,
				},
				resource.Attribute{
					Name:        "source_port_range",
					Description: `The Source Port or Range.`,
				},
				resource.Attribute{
					Name:        "destination_port_range",
					Description: `The Destination Port or Range.`,
				},
				resource.Attribute{
					Name:        "source_address_prefix",
					Description: `CIDR or source IP range or`,
				},
				resource.Attribute{
					Name:        "destination_address_prefix",
					Description: `CIDR or destination IP range or`,
				},
				resource.Attribute{
					Name:        "access",
					Description: `Is network traffic is allowed or denied?`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority of the rule`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `The direction specifies if rule will be evaluated on incoming or outgoing traffic.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurestack_platform_image",
			Category:         "Data Sources",
			ShortDescription: `Gets information about a Platform Image.`,
			Description: `

Use this data source to access information about a Platform Image.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "location",
					Description: `(Required) Specifies the Location to pull information about this Platform Image from.`,
				},
				resource.Attribute{
					Name:        "publisher",
					Description: `(Required) Specifies the Publisher associated with the Platform Image.`,
				},
				resource.Attribute{
					Name:        "offer",
					Description: `(Required) Specifies the Offer associated with the Platform Image.`,
				},
				resource.Attribute{
					Name:        "sku",
					Description: `(Required) Specifies the SKU of the Platform Image. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Platform Image.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The latest version of the Platform Image.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Platform Image.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The latest version of the Platform Image.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurestack_public_ip",
			Category:         "Data Sources",
			ShortDescription: `Retrieves information about the specified public IP address.`,
			Description: `

Use this data source to access the properties of an existing Azure Public IP Address.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the public IP address.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) Specifies the name of the resource group. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "domain_name_label",
					Description: `The label for the Domain Name.`,
				},
				resource.Attribute{
					Name:        "idle_timeout_in_minutes",
					Description: `Specifies the timeout for the TCP idle connection.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `Fully qualified domain name of the A DNS record associated with the public IP. This is the concatenation of the domainNameLabel and the regionalized DNS zone.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The IP address value that was allocated.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assigned to the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "domain_name_label",
					Description: `The label for the Domain Name.`,
				},
				resource.Attribute{
					Name:        "idle_timeout_in_minutes",
					Description: `Specifies the timeout for the TCP idle connection.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `Fully qualified domain name of the A DNS record associated with the public IP. This is the concatenation of the domainNameLabel and the regionalized DNS zone.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The IP address value that was allocated.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assigned to the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurestack_resource_group",
			Category:         "Data Sources",
			ShortDescription: `Get information about the specified resource group.`,
			Description: `

Use this data source to access the properties of an Azure resource group.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the resource group. ~>`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location of the resource group.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "location",
					Description: `The location of the resource group.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurestack_route_table",
			Category:         "Data Sources",
			ShortDescription: `Gets information about a Route Table`,
			Description: `

Gets information about a Route Table

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Route Table.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The name of the Resource Group in which the Route Table exists. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Route Table ID.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure Region in which the Route Table exists.`,
				},
				resource.Attribute{
					Name:        "route",
					Description: `One or more ` + "`" + `route` + "`" + ` blocks as documented below.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `The collection of Subnets associated with this route table.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the Route Table. The ` + "`" + `route` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Route.`,
				},
				resource.Attribute{
					Name:        "address_prefix",
					Description: `The destination CIDR to which the route applies.`,
				},
				resource.Attribute{
					Name:        "next_hop_type",
					Description: `The type of Azure hop the packet should be sent to.`,
				},
				resource.Attribute{
					Name:        "next_hop_in_ip_address",
					Description: `Contains the IP address packets should be forwarded to.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Route Table ID.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure Region in which the Route Table exists.`,
				},
				resource.Attribute{
					Name:        "route",
					Description: `One or more ` + "`" + `route` + "`" + ` blocks as documented below.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `The collection of Subnets associated with this route table.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the Route Table. The ` + "`" + `route` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Route.`,
				},
				resource.Attribute{
					Name:        "address_prefix",
					Description: `The destination CIDR to which the route applies.`,
				},
				resource.Attribute{
					Name:        "next_hop_type",
					Description: `The type of Azure hop the packet should be sent to.`,
				},
				resource.Attribute{
					Name:        "next_hop_in_ip_address",
					Description: `Contains the IP address packets should be forwarded to.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurestack_storage_account",
			Category:         "Data Sources",
			ShortDescription: `Get information about the specified Storage Account.`,
			Description: `

Gets information about the specified Storage Account.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the Storage Account`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) Specifies the name of the resource group the Storage Account is located in. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Storage Account.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the Storage Account exists`,
				},
				resource.Attribute{
					Name:        "account_kind",
					Description: `(Optional) Defines the Kind of account. Valid option is ` + "`" + `Storage` + "`" + `. . Changing this forces a new resource to be created. Defaults to ` + "`" + `Storage` + "`" + ` currently as per [Azure Stack Storage Differences](https://docs.microsoft.com/en-us/azure/azure-stack/user/azure-stack-acs-differences)`,
				},
				resource.Attribute{
					Name:        "account_tier",
					Description: `Defines the Tier of this storage account.`,
				},
				resource.Attribute{
					Name:        "account_replication_type",
					Description: `Defines the type of replication used for this storage account.`,
				},
				resource.Attribute{
					Name:        "access_tier",
					Description: `(Required for ` + "`" + `BlobStorage` + "`" + ` accounts) Defines the access tier for ` + "`" + `BlobStorage` + "`" + ` accounts. Valid options are ` + "`" + `Hot` + "`" + ` and ` + "`" + `Cold` + "`" + `, defaults to ` + "`" + `Hot` + "`" + `. -`,
				},
				resource.Attribute{
					Name:        "account_encryption_source",
					Description: `The Encryption Source for this Storage Account.`,
				},
				resource.Attribute{
					Name:        "custom_domain",
					Description: `A ` + "`" + `custom_domain` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assigned to the resource.`,
				},
				resource.Attribute{
					Name:        "primary_location",
					Description: `The primary location of the Storage Account.`,
				},
				resource.Attribute{
					Name:        "secondary_location",
					Description: `The secondary location of the Storage Account.`,
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
					Description: `The primary access key for the Storage Account.`,
				},
				resource.Attribute{
					Name:        "secondary_access_key",
					Description: `The secondary access key for the Storage Account.`,
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
					Description: `The connection string associated with the secondary blob location ---`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The Custom Domain Name used for the Storage Account.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Storage Account.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure location where the Storage Account exists`,
				},
				resource.Attribute{
					Name:        "account_kind",
					Description: `(Optional) Defines the Kind of account. Valid option is ` + "`" + `Storage` + "`" + `. . Changing this forces a new resource to be created. Defaults to ` + "`" + `Storage` + "`" + ` currently as per [Azure Stack Storage Differences](https://docs.microsoft.com/en-us/azure/azure-stack/user/azure-stack-acs-differences)`,
				},
				resource.Attribute{
					Name:        "account_tier",
					Description: `Defines the Tier of this storage account.`,
				},
				resource.Attribute{
					Name:        "account_replication_type",
					Description: `Defines the type of replication used for this storage account.`,
				},
				resource.Attribute{
					Name:        "access_tier",
					Description: `(Required for ` + "`" + `BlobStorage` + "`" + ` accounts) Defines the access tier for ` + "`" + `BlobStorage` + "`" + ` accounts. Valid options are ` + "`" + `Hot` + "`" + ` and ` + "`" + `Cold` + "`" + `, defaults to ` + "`" + `Hot` + "`" + `. -`,
				},
				resource.Attribute{
					Name:        "account_encryption_source",
					Description: `The Encryption Source for this Storage Account.`,
				},
				resource.Attribute{
					Name:        "custom_domain",
					Description: `A ` + "`" + `custom_domain` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assigned to the resource.`,
				},
				resource.Attribute{
					Name:        "primary_location",
					Description: `The primary location of the Storage Account.`,
				},
				resource.Attribute{
					Name:        "secondary_location",
					Description: `The secondary location of the Storage Account.`,
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
					Description: `The primary access key for the Storage Account.`,
				},
				resource.Attribute{
					Name:        "secondary_access_key",
					Description: `The secondary access key for the Storage Account.`,
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
					Description: `The connection string associated with the secondary blob location ---`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The Custom Domain Name used for the Storage Account.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurestack_subnet",
			Category:         "Data Sources",
			ShortDescription: `Get information about the specified Subnet located within a Virtual Network.`,
			Description: `

Use this data source to access the properties of an Azure Subnet located within a Virtual Network.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the Subnet.`,
				},
				resource.Attribute{
					Name:        "virtual_network_name",
					Description: `(Required) Specifies the name of the Virtual Network this Subnet is located within.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) Specifies the name of the resource group the Virtual Network is located in. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Subnet.`,
				},
				resource.Attribute{
					Name:        "address_prefix",
					Description: `The address prefix used for the subnet.`,
				},
				resource.Attribute{
					Name:        "network_security_group_id",
					Description: `The ID of the Network Security Group associated with the subnet.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The ID of the Route Table associated with this subnet.`,
				},
				resource.Attribute{
					Name:        "ip_configurations",
					Description: `The collection of IP Configurations with IPs within this subnet.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Subnet.`,
				},
				resource.Attribute{
					Name:        "address_prefix",
					Description: `The address prefix used for the subnet.`,
				},
				resource.Attribute{
					Name:        "network_security_group_id",
					Description: `The ID of the Network Security Group associated with the subnet.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The ID of the Route Table associated with this subnet.`,
				},
				resource.Attribute{
					Name:        "ip_configurations",
					Description: `The collection of IP Configurations with IPs within this subnet.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurestack_virtual_network",
			Category:         "Data Sources",
			ShortDescription: `Get information about the specified Virtual Network.`,
			Description: `

Use this data source to access the properties of an Azure Virtual Network.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the Virtual Network.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) Specifies the name of the resource group the Virtual Network is located in. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the virtual network.`,
				},
				resource.Attribute{
					Name:        "address_spaces",
					Description: `The list of address spaces used by the virtual network.`,
				},
				resource.Attribute{
					Name:        "dns_servers",
					Description: `The list of DNS servers used by the virtual network.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `The list of name of the subnets that are attached to this virtual network.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the virtual network.`,
				},
				resource.Attribute{
					Name:        "address_spaces",
					Description: `The list of address spaces used by the virtual network.`,
				},
				resource.Attribute{
					Name:        "dns_servers",
					Description: `The list of DNS servers used by the virtual network.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `The list of name of the subnets that are attached to this virtual network.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azurestack_virtual_network_gateway",
			Category:         "Data Sources",
			ShortDescription: `Get information about the specified Virtual Network Gateway.`,
			Description: `

Use this data source to access the properties of an Azure Virtual Network Gateway.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the Virtual Network Gateway.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) Specifies the name of the resource group the Virtual Network Gateway is located in. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Virtual Network Gateway.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location/region where the Virtual Network Gateway is located.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the Virtual Network Gateway.`,
				},
				resource.Attribute{
					Name:        "vpn_type",
					Description: `The routing type of the Virtual Network Gateway.`,
				},
				resource.Attribute{
					Name:        "enable_bgp",
					Description: `Will BGP (Border Gateway Protocol) will be enabled for this Virtual Network Gateway.`,
				},
				resource.Attribute{
					Name:        "default_local_network_gateway_id",
					Description: `The ID of the local network gateway through which outbound Internet traffic from the virtual network in which the gateway is created will be routed (`,
				},
				resource.Attribute{
					Name:        "sku",
					Description: `Configuration of the size and capacity of the Virtual Network Gateway.`,
				},
				resource.Attribute{
					Name:        "ip_configuration",
					Description: `One or two ` + "`" + `ip_configuration` + "`" + ` blocks documented below.`,
				},
				resource.Attribute{
					Name:        "vpn_client_configuration",
					Description: `A ` + "`" + `vpn_client_configuration` + "`" + ` block which is documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource. The ` + "`" + `ip_configuration` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A user-defined name of the IP configuration.`,
				},
				resource.Attribute{
					Name:        "private_ip_address_allocation",
					Description: `Defines how the private IP address of the gateways virtual interface is assigned.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the gateway subnet of a virtual network in which the virtual network gateway will be created. It is mandatory that the associated subnet is named ` + "`" + `GatewaySubnet` + "`" + `. Therefore, each virtual network can contain at most a single Virtual Network Gateway.`,
				},
				resource.Attribute{
					Name:        "public_ip_address_id",
					Description: `The ID of the Public IP Address associated with the Virtual Network Gateway. The ` + "`" + `vpn_client_configuration` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "address_space",
					Description: `The address space out of which ip addresses for vpn clients will be taken. You can provide more than one address space, e.g. in CIDR notation.`,
				},
				resource.Attribute{
					Name:        "root_certificate",
					Description: `One or more ` + "`" + `root_certificate` + "`" + ` blocks which are defined below. These root certificates are used to sign the client certificate used by the VPN clients to connect to the gateway.`,
				},
				resource.Attribute{
					Name:        "revoked_certificate",
					Description: `One or more ` + "`" + `revoked_certificate` + "`" + ` blocks which are defined below. The ` + "`" + `bgp_settings` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "asn",
					Description: `The Autonomous System Number (ASN) to use as part of the BGP.`,
				},
				resource.Attribute{
					Name:        "peering_address",
					Description: `The BGP peer IP address of the virtual network gateway. This address is needed to configure the created gateway as a BGP Peer on the on-premises VPN devices.`,
				},
				resource.Attribute{
					Name:        "peer_weight",
					Description: `The weight added to routes which have been learned through BGP peering. The ` + "`" + `root_certificate` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The user-defined name of the root certificate.`,
				},
				resource.Attribute{
					Name:        "public_cert_data",
					Description: `The public certificate of the root certificate authority. The certificate must be provided in Base-64 encoded X.509 format (PEM). The ` + "`" + `root_revoked_certificate` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The user-defined name of the revoked certificate.`,
				},
				resource.Attribute{
					Name:        "public_cert_data",
					Description: `The SHA1 thumbprint of the certificate to be revoked.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Virtual Network Gateway.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location/region where the Virtual Network Gateway is located.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the Virtual Network Gateway.`,
				},
				resource.Attribute{
					Name:        "vpn_type",
					Description: `The routing type of the Virtual Network Gateway.`,
				},
				resource.Attribute{
					Name:        "enable_bgp",
					Description: `Will BGP (Border Gateway Protocol) will be enabled for this Virtual Network Gateway.`,
				},
				resource.Attribute{
					Name:        "default_local_network_gateway_id",
					Description: `The ID of the local network gateway through which outbound Internet traffic from the virtual network in which the gateway is created will be routed (`,
				},
				resource.Attribute{
					Name:        "sku",
					Description: `Configuration of the size and capacity of the Virtual Network Gateway.`,
				},
				resource.Attribute{
					Name:        "ip_configuration",
					Description: `One or two ` + "`" + `ip_configuration` + "`" + ` blocks documented below.`,
				},
				resource.Attribute{
					Name:        "vpn_client_configuration",
					Description: `A ` + "`" + `vpn_client_configuration` + "`" + ` block which is documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags assigned to the resource. The ` + "`" + `ip_configuration` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A user-defined name of the IP configuration.`,
				},
				resource.Attribute{
					Name:        "private_ip_address_allocation",
					Description: `Defines how the private IP address of the gateways virtual interface is assigned.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the gateway subnet of a virtual network in which the virtual network gateway will be created. It is mandatory that the associated subnet is named ` + "`" + `GatewaySubnet` + "`" + `. Therefore, each virtual network can contain at most a single Virtual Network Gateway.`,
				},
				resource.Attribute{
					Name:        "public_ip_address_id",
					Description: `The ID of the Public IP Address associated with the Virtual Network Gateway. The ` + "`" + `vpn_client_configuration` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "address_space",
					Description: `The address space out of which ip addresses for vpn clients will be taken. You can provide more than one address space, e.g. in CIDR notation.`,
				},
				resource.Attribute{
					Name:        "root_certificate",
					Description: `One or more ` + "`" + `root_certificate` + "`" + ` blocks which are defined below. These root certificates are used to sign the client certificate used by the VPN clients to connect to the gateway.`,
				},
				resource.Attribute{
					Name:        "revoked_certificate",
					Description: `One or more ` + "`" + `revoked_certificate` + "`" + ` blocks which are defined below. The ` + "`" + `bgp_settings` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "asn",
					Description: `The Autonomous System Number (ASN) to use as part of the BGP.`,
				},
				resource.Attribute{
					Name:        "peering_address",
					Description: `The BGP peer IP address of the virtual network gateway. This address is needed to configure the created gateway as a BGP Peer on the on-premises VPN devices.`,
				},
				resource.Attribute{
					Name:        "peer_weight",
					Description: `The weight added to routes which have been learned through BGP peering. The ` + "`" + `root_certificate` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The user-defined name of the root certificate.`,
				},
				resource.Attribute{
					Name:        "public_cert_data",
					Description: `The public certificate of the root certificate authority. The certificate must be provided in Base-64 encoded X.509 format (PEM). The ` + "`" + `root_revoked_certificate` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The user-defined name of the revoked certificate.`,
				},
				resource.Attribute{
					Name:        "public_cert_data",
					Description: `The SHA1 thumbprint of the certificate to be revoked.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"azurestack_network_interface":       0,
		"azurestack_network_security_group":  1,
		"azurestack_platform_image":          2,
		"azurestack_public_ip":               3,
		"azurestack_resource_group":          4,
		"azurestack_route_table":             5,
		"azurestack_storage_account":         6,
		"azurestack_subnet":                  7,
		"azurestack_virtual_network":         8,
		"azurestack_virtual_network_gateway": 9,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
