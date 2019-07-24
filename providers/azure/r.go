package azure

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "azure_affinity_group",
			Category:         "Resources",
			ShortDescription: `Creates a new affinity group on Azure.`,
			Description:      ``,
			Keywords: []string{
				"affinity",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the affinity group. Must be unique on your Azure subscription.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Required) The location where the affinity group should be created. For a list of all Azure locations, please consult [this link](https://azure.microsoft.com/en-us/regions/).`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Required) A label to be used for tracking purposes.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description for the affinity group. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The affinity group ID. Coincides with the given ` + "`" + `name` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The affinity group ID. Coincides with the given ` + "`" + `name` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azure_data_disk",
			Category:         "Resources",
			ShortDescription: `Adds a data disk to a virtual machine. If the name of an existing disk is given, it will attach that disk. Otherwise it will create and attach a new empty disk.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"disk",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of an existing registered disk to attach to the virtual machine. If left empty, a new empty disk will be created and attached instead. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Optional) The identifier of the data disk. Changing this forces a new resource to be created (defaults to "virtual_machine-lun")`,
				},
				resource.Attribute{
					Name:        "lun",
					Description: `(Required) The Logical Unit Number (LUN) for the disk. The LUN specifies the slot in which the data drive appears when mounted for usage by the virtual machine. Valid LUN values are 0 through 31.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) The size, in GB, of an empty disk to be attached to the virtual machine. Required when creating a new disk, not used otherwise.`,
				},
				resource.Attribute{
					Name:        "caching",
					Description: `(Optional) The caching behavior of data disk. Valid options are: ` + "`" + `None` + "`" + `, ` + "`" + `ReadOnly` + "`" + ` and ` + "`" + `ReadWrite` + "`" + ` (defaults ` + "`" + `None` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "storage_service_name",
					Description: `(Optional) The name of an existing storage account within the subscription which will be used to store the VHD of this disk. Required if no value is supplied for ` + "`" + `media_link` + "`" + `. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "media_link",
					Description: `(Optional) The location of the blob in storage where the VHD of this disk will be created. The storage account where must be associated with the subscription. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "source_media_link",
					Description: `(Optional) The location of a blob in storage where a VHD file is located that is imported and registered as a disk. If a value is supplied, ` + "`" + `media_link` + "`" + ` will not be used.`,
				},
				resource.Attribute{
					Name:        "virtual_machine",
					Description: `(Required) The name of the virtual machine the disk will be attached to. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The security group ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the disk.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The identifier for the disk.`,
				},
				resource.Attribute{
					Name:        "media_link",
					Description: `The location of the blob in storage where the VHD of this disk is created.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The security group ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the disk.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The identifier for the disk.`,
				},
				resource.Attribute{
					Name:        "media_link",
					Description: `The location of the blob in storage where the VHD of this disk is created.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azure_dns_server",
			Category:         "Resources",
			ShortDescription: `Creates a new DNS server definition to be used internally in Azure.`,
			Description:      ``,
			Keywords: []string{
				"dns",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the DNS server reference. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "dns_address",
					Description: `(Required) The IP address of the DNS server. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The DNS server definition ID. Coincides with the given ` + "`" + `name` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The DNS server definition ID. Coincides with the given ` + "`" + `name` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azure_hosted_service",
			Category:         "Resources",
			ShortDescription: `Creates a new hosted service on Azure with its own .cloudapp.net domain.`,
			Description:      ``,
			Keywords: []string{
				"hosted",
				"service",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the hosted service. Must be unique on Azure.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Required) The location where the hosted service should be created. For a list of all Azure locations, please consult [this link](https://azure.microsoft.com/en-us/regions/).`,
				},
				resource.Attribute{
					Name:        "ephemeral_contents",
					Description: `(Required) A boolean value (true|false), specifying whether all the resources present in the hosted hosted service should be destroyed following the hosted service's destruction.`,
				},
				resource.Attribute{
					Name:        "reverse_dns_fqdn",
					Description: `(Optional) The reverse of the fully qualified domain name for the hosted service.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Optional) A label to be used for tracking purposes. Must be non-void. Defaults to ` + "`" + `Made by Terraform.` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description for the hosted service. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The hosted service ID. Coincides with the given ` + "`" + `name` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The hosted service ID. Coincides with the given ` + "`" + `name` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azure_instance",
			Category:         "Resources",
			ShortDescription: `Creates a hosted service, role and deployment and then creates a virtual machine in the deployment based on the specified configuration.`,
			Description:      ``,
			Keywords: []string{
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the instance. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "hosted_service_name",
					Description: `(Optional) The name of the hosted service the instance should be deployed under. If not provided; it will default to the value of ` + "`" + `name` + "`" + `. Changes to this parameter forces the creation of a new resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description for the associated hosted service. Changing this forces a new resource to be created (defaults to the instance name).`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Required) The name of an existing VM or OS image to use for this instance. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) The size of the instance.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `(Optional) The name of the subnet to connect this instance to. If a value is supplied ` + "`" + `virtual_network` + "`" + ` is required. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "virtual_network",
					Description: `(Optional) The name of the virtual network the ` + "`" + `subnet` + "`" + ` belongs to. If a value is supplied ` + "`" + `subnet` + "`" + ` is required. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "storage_service_name",
					Description: `(Optional) The name of an existing storage account within the subscription which will be used to store the VHDs of this instance. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "reverse_dns",
					Description: `(Optional) The DNS address to which the IP address of the hosted service resolves when queried using a reverse DNS query. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Required) The location/region where the cloud service is created. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "automatic_updates",
					Description: `(Optional) If true this will enable automatic updates. This attribute is only used when creating a Windows instance. Changing this forces a new resource to be created (defaults false)`,
				},
				resource.Attribute{
					Name:        "time_zone",
					Description: `(Optional) The appropriate time zone for this instance in the format 'America/Los_Angeles'. This attribute is only used when creating a Windows instance. Changing this forces a new resource to be created (defaults false)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The username of a new user that will be created while creating the instance. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The password of the new user that will be created while creating the instance. Required when creating a Windows instance or when not supplying an ` + "`" + `ssh_key_thumbprint` + "`" + ` while creating a Linux instance. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "ssh_key_thumbprint",
					Description: `(Optional) The SSH thumbprint of an existing SSH key within the subscription. This attribute is only used when creating a Linux instance. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "security_group",
					Description: `(Optional) The Network Security Group to associate with this instance.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Optional) Can be specified multiple times to define multiple endpoints. Each ` + "`" + `endpoint` + "`" + ` block supports fields documented below.`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `(Optional) The name of an Active Directory domain to join.`,
				},
				resource.Attribute{
					Name:        "domain_ou",
					Description: `(Optional) Specifies the LDAP Organizational Unit to place the instance in.`,
				},
				resource.Attribute{
					Name:        "domain_username",
					Description: `(Optional) The username of an account with permission to join the instance to the domain. Required if a domain_name is specified.`,
				},
				resource.Attribute{
					Name:        "domain_password",
					Description: `(Optional) The password for the domain_username account specified above.`,
				},
				resource.Attribute{
					Name:        "custom_data",
					Description: `(Optional) The custom data to provide when launching the instance. The ` + "`" + `endpoint` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the external endpoint.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) The transport protocol for the endpoint. Valid options are: ` + "`" + `tcp` + "`" + ` and ` + "`" + `udp` + "`" + ` (defaults ` + "`" + `tcp` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "public_port",
					Description: `(Required) The external port to use for the endpoint.`,
				},
				resource.Attribute{
					Name:        "private_port",
					Description: `(Required) The private port on which the instance is listening. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The instance ID.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description for the associated hosted service.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `The subnet the instance is connected to.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `The complete set of configured endpoints.`,
				},
				resource.Attribute{
					Name:        "security_group",
					Description: `The associated Network Security Group.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The private IP address assigned to the instance.`,
				},
				resource.Attribute{
					Name:        "vip_address",
					Description: `The public IP address assigned to the instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The instance ID.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description for the associated hosted service.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `The subnet the instance is connected to.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `The complete set of configured endpoints.`,
				},
				resource.Attribute{
					Name:        "security_group",
					Description: `The associated Network Security Group.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The private IP address assigned to the instance.`,
				},
				resource.Attribute{
					Name:        "vip_address",
					Description: `The public IP address assigned to the instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azure_local_network_connection",
			Category:         "Resources",
			ShortDescription: `Defines a new connection to a remote network through a VPN tunnel.`,
			Description:      ``,
			Keywords: []string{
				"local",
				"network",
				"connection",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name by which this local network connection will be referenced by. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "vpn_gateway_address",
					Description: `(Required) The public IPv4 of the VPN endpoint.`,
				},
				resource.Attribute{
					Name:        "address_space_prefixes",
					Description: `(Required) List of address spaces accessible through the VPN connection. The elements are in the CIDR format. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The local network connection ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The local network connection ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azure_security_group",
			Category:         "Resources",
			ShortDescription: `Creates a new network security group within the context of the specified subscription.`,
			Description:      ``,
			Keywords: []string{
				"security",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the security group. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Optional) The identifier for the security group. The label can be up to 1024 characters long. Changing this forces a new resource to be created (defaults to the security group name)`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Required) The location/region where the security group is created. Changing this forces a new resource to be created. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The security group ID.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The identifier for the security group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The security group ID.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The identifier for the security group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azure_security_group_rule",
			Category:         "Resources",
			ShortDescription: `Creates a new network security rule to be associated with a given security group.`,
			Description:      ``,
			Keywords: []string{
				"security",
				"group",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the security group rule.`,
				},
				resource.Attribute{
					Name:        "security_group_names",
					Description: `(Required) A list of the names of the security groups the rule should be applied to. Changing this list forces the creation of a new resource.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of the security rule. Valid options are: ` + "`" + `Inbound` + "`" + ` and ` + "`" + `Outbound` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Required) The priority of the network security rule. Rules with lower priority are evaluated first. This value can be between 100 and 4096.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) The action that is performed when the security rule is matched. Valid options are: ` + "`" + `Allow` + "`" + ` and ` + "`" + `Deny` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "source_address_prefix",
					Description: `(Required) The address prefix of packet sources that that should be subjected to the rule. An asterisk (\`,
				},
				resource.Attribute{
					Name:        "source_port_range",
					Description: `(Required) The source port or range. This value can be between 0 and 65535. An asterisk (\`,
				},
				resource.Attribute{
					Name:        "destination_address_prefix",
					Description: `(Required) The address prefix of packet destinations that should be subjected to the rule. An asterisk (\`,
				},
				resource.Attribute{
					Name:        "destination_port_range",
					Description: `(Required) The destination port or range. This value can be between 0 and 65535. An asterisk (\`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) The protocol of the security rule. Valid options are: ` + "`" + `TCP` + "`" + `, ` + "`" + `UDP` + "`" + ` and ` + "`" + ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The security group rule ID. Coincides with its given ` + "`" + `name` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azure_sql_database_server",
			Category:         "Resources",
			ShortDescription: `Allocates a new SQL Database Server on Azure.`,
			Description:      ``,
			Keywords: []string{
				"sql",
				"database",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Computed) The name of the database server. It is determined upon creation as it is randomly-generated per server.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Required) The location where the database server should be created. For a list of all Azure locations, please consult [this link](https://azure.microsoft.com/en-us/regions/).`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The username for the administrator of the database server.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) The password for the administrator of the database server.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) The version of the database server to be used. Can be any one of ` + "`" + `2.0` + "`" + ` or ` + "`" + `12.0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Computed) The fully qualified domain name of the database server. Will be of the form ` + "`" + `<name>.database.windows.net` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The database server ID. Coincides with the randomly-generated ` + "`" + `name` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The database server ID. Coincides with the randomly-generated ` + "`" + `name` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azure_sql_database_server_firewall_rule",
			Category:         "Resources",
			ShortDescription: `Defines a new Firewall Rule to be applied across the given Database Servers.`,
			Description:      ``,
			Keywords: []string{
				"sql",
				"database",
				"server",
				"firewall",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the rule. Changing forces the creation of a new resource.`,
				},
				resource.Attribute{
					Name:        "start_ip",
					Description: `(Required) The IPv4 which will represent the lower bound of the rule's application IP's. Traffic to/from IP's greater than or equal to this one up to the ` + "`" + `end_ip` + "`" + ` will be permitted.`,
				},
				resource.Attribute{
					Name:        "end_ip",
					Description: `(Required) The IPv4 which will represent the upper bound of the rule's application IP's. Traffic to/from IP's lesser that or equal to this one all the way down to the ` + "`" + `start_ip` + "`" + ` will be permitted.`,
				},
				resource.Attribute{
					Name:        "database_server_names",
					Description: `(Required) The set of names of the Azure SQL Database servers the rule should be enforced on. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The database server ID. Coincides with the given ` + "`" + `name` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The database server ID. Coincides with the given ` + "`" + `name` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azure_sql_database_service",
			Category:         "Resources",
			ShortDescription: `Creates a new SQL Database Service on an Azure Database Server.`,
			Description:      ``,
			Keywords: []string{
				"sql",
				"database",
				"service",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the database service.`,
				},
				resource.Attribute{
					Name:        "database_server_name",
					Description: `(Required) The name of the database server this service should run on. Changes here force the creation of a new resource.`,
				},
				resource.Attribute{
					Name:        "edition",
					Description: `(Optional) The edition of the database service. For more information on each variant, please view [this](https://msdn.microsoft.com/library/azure/dn741340.aspx) link.`,
				},
				resource.Attribute{
					Name:        "collation",
					Description: `(Optional) The collation to be used within the database service. Defaults to the standard Latin charset.`,
				},
				resource.Attribute{
					Name:        "max_size_bytes",
					Description: `(Optional) The maximum size in bytes the database service should be allowed to expand to. Range depends on the database ` + "`" + `edition` + "`" + ` selected above.`,
				},
				resource.Attribute{
					Name:        "service_level_id",
					Description: `(Optional) The ID corresponding to the service level per edition. Please refer to [this](https://msdn.microsoft.com/en-us/library/azure/dn505701.aspx) link for more details. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The database service ID. Coincides with the given ` + "`" + `name` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The database service ID. Coincides with the given ` + "`" + `name` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azure_storage_blob",
			Category:         "Resources",
			ShortDescription: `Creates a new storage blob within a given storage container on Azure.`,
			Description:      ``,
			Keywords: []string{
				"storage",
				"blob",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the storage blob. Must be unique within the storage service the blob is located.`,
				},
				resource.Attribute{
					Name:        "storage_service_name",
					Description: `(Required) The name of the storage service within which the storage container in which the blob will be created resides.`,
				},
				resource.Attribute{
					Name:        "storage_container_name",
					Description: `(Required) The name of the storage container in which this blob should be created. Must be located on the storage service given with ` + "`" + `storage_service_name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of the storage blob to be created. One of either ` + "`" + `BlockBlob` + "`" + ` or ` + "`" + `PageBlob` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) Used only for ` + "`" + `PageBlob` + "`" + `'s to specify the size in bytes of the blob to be created. Must be a multiple of 512. Defaults to 0. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The storage blob ID. Coincides with the given ` + "`" + `name` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The storage blob ID. Coincides with the given ` + "`" + `name` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azure_storage_container",
			Category:         "Resources",
			ShortDescription: `Creates a new storage container within a given storage service on Azure.`,
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
					Name:        "storage_service_name",
					Description: `(Required) The name of the storage service within which the storage container should be created.`,
				},
				resource.Attribute{
					Name:        "container_access_type",
					Description: `(Required) The 'interface' for access the container provides. Can be either ` + "`" + `blob` + "`" + `, ` + "`" + `container` + "`" + ` or ` + "`" + `` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "properties",
					Description: `(Optional) Key-value definition of additional properties associated to the storage service. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The storage container ID. Coincides with the given ` + "`" + `name` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The storage container ID. Coincides with the given ` + "`" + `name` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azure_storage_queue",
			Category:         "Resources",
			ShortDescription: `Creates a new storage queue within a given storage service on Azure.`,
			Description:      ``,
			Keywords: []string{
				"storage",
				"queue",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the storage queue. Must be unique within the storage service the queue is located.`,
				},
				resource.Attribute{
					Name:        "storage_service_name",
					Description: `(Required) The name of the storage service within which the storage queue should be created. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The storage queue ID. Coincides with the given ` + "`" + `name` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The storage queue ID. Coincides with the given ` + "`" + `name` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azure_storage_service",
			Category:         "Resources",
			ShortDescription: `Creates a new storage service on Azure in which storage containers may be created.`,
			Description:      ``,
			Keywords: []string{
				"storage",
				"service",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the storage service. Must be between 4 and 24 lowercase-only characters or digits. Must be unique on Azure.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Required) The location where the storage service should be created. For a list of all Azure locations, please consult [this link](https://azure.microsoft.com/en-us/regions/).`,
				},
				resource.Attribute{
					Name:        "account_type",
					Description: `(Required) The type of storage account to be created. Available options include ` + "`" + `Standard_LRS` + "`" + `, ` + "`" + `Standard_ZRS` + "`" + `, ` + "`" + `Standard_GRS` + "`" + `, ` + "`" + `Standard_RAGRS` + "`" + ` and ` + "`" + `Premium_LRS` + "`" + `. To learn more about the differences of each storage account type, please consult [this link](http://blogs.msdn.com/b/windowsazurestorage/archive/2013/12/11/introducing-read-access-geo-replicated-storage-ra-grs-for-windows-azure-storage.aspx).`,
				},
				resource.Attribute{
					Name:        "affinity_group",
					Description: `(Optional) The affinity group the storage service should belong to.`,
				},
				resource.Attribute{
					Name:        "properties",
					Description: `(Optional) Key-value definition of additional properties associated to the storage service. For additional information on what these properties do, please consult [this link](https://msdn.microsoft.com/en-us/library/azure/hh452235.aspx).`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Optional) A label to be used for tracking purposes. Must be non-void. Defaults to ` + "`" + `Made by Terraform.` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description for the storage service. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The storage service ID. Coincides with the given ` + "`" + `name` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The storage service ID. Coincides with the given ` + "`" + `name` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azure_virtual_network",
			Category:         "Resources",
			ShortDescription: `Creates a new virtual network including any configured subnets. Each subnet can optionally be configured with a security group to be associated with the subnet.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"network",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the virtual network. Changing this forces a new resource to be created.`,
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
					Description: `(Optional) List of names of DNS servers previously registered on Azure.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `(Required) Can be specified multiple times to define multiple subnets. Each ` + "`" + `subnet` + "`" + ` block supports fields documented below. The ` + "`" + `subnet` + "`" + ` block supports:`,
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
					Description: `(Optional) The Network Security Group to associate with the subnet. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The virtual NetworkConfiguration ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The virtual NetworkConfiguration ID.`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"azure_affinity_group":                    0,
		"azure_data_disk":                         1,
		"azure_dns_server":                        2,
		"azure_hosted_service":                    3,
		"azure_instance":                          4,
		"azure_local_network_connection":          5,
		"azure_security_group":                    6,
		"azure_security_group_rule":               7,
		"azure_sql_database_server":               8,
		"azure_sql_database_server_firewall_rule": 9,
		"azure_sql_database_service":              10,
		"azure_storage_blob":                      11,
		"azure_storage_container":                 12,
		"azure_storage_queue":                     13,
		"azure_storage_service":                   14,
		"azure_virtual_network":                   15,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
