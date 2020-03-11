package aviatrix

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_account",
			Category:         "Data Sources",
			ShortDescription: `Gets the an Aviatrix cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_name",
					Description: `(Required) Account name. This can be used for logging in to CloudN console or UserConnect controller. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `Type of cloud service provider. (Only AWS is supported currently. Value of 1 for AWS.)`,
				},
				resource.Attribute{
					Name:        "aws_account_number",
					Description: `AWS Account number to associate with Aviatrix account.`,
				},
				resource.Attribute{
					Name:        "aws_access_key",
					Description: `AWS Access Key.`,
				},
				resource.Attribute{
					Name:        "aws_role_app",
					Description: `AWS App role ARN.`,
				},
				resource.Attribute{
					Name:        "aws_role_ec2",
					Description: `AWS EC2 role ARN.`,
				},
				resource.Attribute{
					Name:        "gcloud_project_id",
					Description: `GCloud Project ID.`,
				},
				resource.Attribute{
					Name:        "gcloud_project_credentials_filepath",
					Description: `GCloud Project Credentials.`,
				},
				resource.Attribute{
					Name:        "arm_subscription_id",
					Description: `Azure ARM Subscription ID.`,
				},
				resource.Attribute{
					Name:        "arm_directory_id",
					Description: `Azure ARM Directory ID.`,
				},
				resource.Attribute{
					Name:        "arm_application_id",
					Description: `Azure ARM Application ID.`,
				},
				resource.Attribute{
					Name:        "arm_application_key",
					Description: `Azure ARM Application key.`,
				},
				resource.Attribute{
					Name:        "oci_tenancy_id",
					Description: `Oracle OCI Tenancy ID.`,
				},
				resource.Attribute{
					Name:        "oci_user_id",
					Description: `Oracle OCI User ID.`,
				},
				resource.Attribute{
					Name:        "oci_compartment_id",
					Description: `Oracle OCI Compartment ID.`,
				},
				resource.Attribute{
					Name:        "oci_api_private_key_filepath",
					Description: `Oracle OCI API Private Key local file path.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_type",
					Description: `Type of cloud service provider. (Only AWS is supported currently. Value of 1 for AWS.)`,
				},
				resource.Attribute{
					Name:        "aws_account_number",
					Description: `AWS Account number to associate with Aviatrix account.`,
				},
				resource.Attribute{
					Name:        "aws_access_key",
					Description: `AWS Access Key.`,
				},
				resource.Attribute{
					Name:        "aws_role_app",
					Description: `AWS App role ARN.`,
				},
				resource.Attribute{
					Name:        "aws_role_ec2",
					Description: `AWS EC2 role ARN.`,
				},
				resource.Attribute{
					Name:        "gcloud_project_id",
					Description: `GCloud Project ID.`,
				},
				resource.Attribute{
					Name:        "gcloud_project_credentials_filepath",
					Description: `GCloud Project Credentials.`,
				},
				resource.Attribute{
					Name:        "arm_subscription_id",
					Description: `Azure ARM Subscription ID.`,
				},
				resource.Attribute{
					Name:        "arm_directory_id",
					Description: `Azure ARM Directory ID.`,
				},
				resource.Attribute{
					Name:        "arm_application_id",
					Description: `Azure ARM Application ID.`,
				},
				resource.Attribute{
					Name:        "arm_application_key",
					Description: `Azure ARM Application key.`,
				},
				resource.Attribute{
					Name:        "oci_tenancy_id",
					Description: `Oracle OCI Tenancy ID.`,
				},
				resource.Attribute{
					Name:        "oci_user_id",
					Description: `Oracle OCI User ID.`,
				},
				resource.Attribute{
					Name:        "oci_compartment_id",
					Description: `Oracle OCI Compartment ID.`,
				},
				resource.Attribute{
					Name:        "oci_api_private_key_filepath",
					Description: `Oracle OCI API Private Key local file path.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_caller_identity",
			Category:         "Data Sources",
			ShortDescription: `Gets the an Aviatrix caller identity.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cid",
					Description: `Aviatrix caller identity.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cid",
					Description: `Aviatrix caller identity.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_firenet",
			Category:         "Data Sources",
			ShortDescription: `Gets the Aviatrix FireNet.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) ID of the Security VPC. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the Security VPC.`,
				},
				resource.Attribute{
					Name:        "inspection_enabled",
					Description: `Enable/Disable traffic inspection.`,
				},
				resource.Attribute{
					Name:        "egress_enabled",
					Description: `Enable/Disable egress through firewall.`,
				},
				resource.Attribute{
					Name:        "firewall_instance_association",
					Description: `List of firewall instances associated with fireNet.`,
				},
				resource.Attribute{
					Name:        "firenet_gw_name",
					Description: `Name of the primary FireNet gateway.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of Firewall instance.`,
				},
				resource.Attribute{
					Name:        "vendor_type",
					Description: `Type of the firewall.`,
				},
				resource.Attribute{
					Name:        "firewall_name",
					Description: `Firewall instance name.`,
				},
				resource.Attribute{
					Name:        "management_interface",
					Description: `Management interface ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the Security VPC.`,
				},
				resource.Attribute{
					Name:        "inspection_enabled",
					Description: `Enable/Disable traffic inspection.`,
				},
				resource.Attribute{
					Name:        "egress_enabled",
					Description: `Enable/Disable egress through firewall.`,
				},
				resource.Attribute{
					Name:        "firewall_instance_association",
					Description: `List of firewall instances associated with fireNet.`,
				},
				resource.Attribute{
					Name:        "firenet_gw_name",
					Description: `Name of the primary FireNet gateway.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of Firewall instance.`,
				},
				resource.Attribute{
					Name:        "vendor_type",
					Description: `Type of the firewall.`,
				},
				resource.Attribute{
					Name:        "firewall_name",
					Description: `Firewall instance name.`,
				},
				resource.Attribute{
					Name:        "management_interface",
					Description: `Management interface ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_firenet_vendor_integration",
			Category:         "Data Sources",
			ShortDescription: `Does 'save' or 'sync' for vendor integration purpose for Aviatrix FireNet.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) VPC ID.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) ID of Firewall instance.`,
				},
				resource.Attribute{
					Name:        "vendor_type",
					Description: `(Required) Select PAN. Valid values: "Generic", "Palo Alto Networks VM-Series", "Aviatrix FQDN Gateway".`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `(Required) The public IP address of the firewall management interface for API calls from the Aviatrix Controller.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) Firewall login name for API calls from the Controller.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Firewall login password for API calls.`,
				},
				resource.Attribute{
					Name:        "firewall_name",
					Description: `(Optional) Name of firewall instance.`,
				},
				resource.Attribute{
					Name:        "route_table",
					Description: `(Optional) Specify the firewall virtual Router name you wish the Controller to program. If left unspecified, the Controller programs the firewall’s default router.`,
				},
				resource.Attribute{
					Name:        "number_of_retries",
					Description: `(Optional) Number of retries for ` + "`" + `save` + "`" + ` or ` + "`" + `synchronize` + "`" + `. Example: 1. Default value: 0.`,
				},
				resource.Attribute{
					Name:        "retry_interval",
					Description: `(Optional) Retry interval in seconds for ` + "`" + `save` + "`" + ` or ` + "`" + `synchronize` + "`" + `. Example: 120. Default value: 300.`,
				},
				resource.Attribute{
					Name:        "save",
					Description: `(Optional) Switch to save or not.`,
				},
				resource.Attribute{
					Name:        "synchronize",
					Description: `(Optional) Switch to sync or not.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_gateway",
			Category:         "Data Sources",
			ShortDescription: `Gets the Aviatrix gateway.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "gw_name",
					Description: `(Required) Gateway name. This can be used for getting gateway.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `(Optional) Account name. This can be used for logging in to CloudN console or UserConnect controller. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `Aviatrix account name.`,
				},
				resource.Attribute{
					Name:        "gw_name",
					Description: `Aviatrix gateway name.`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `Type of cloud service provider.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC ID.`,
				},
				resource.Attribute{
					Name:        "vpc_reg",
					Description: `VPC Region.`,
				},
				resource.Attribute{
					Name:        "vpc_size",
					Description: `Instance type.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP address of the Gateway created.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "account_name",
					Description: `Aviatrix account name.`,
				},
				resource.Attribute{
					Name:        "gw_name",
					Description: `Aviatrix gateway name.`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `Type of cloud service provider.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC ID.`,
				},
				resource.Attribute{
					Name:        "vpc_reg",
					Description: `VPC Region.`,
				},
				resource.Attribute{
					Name:        "vpc_size",
					Description: `Instance type.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP address of the Gateway created.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_spoke_gateway",
			Category:         "Data Sources",
			ShortDescription: `Gets the Aviatrix spoke gateway.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "gw_name",
					Description: `(Required) Spoke gateway name. This can be used for getting spoke gateway.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `(Optional) Account name. This can be used for logging in to CloudN console or UserConnect controller. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "gw_name",
					Description: `Aviatrix spoke gateway name.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `Aviatrix account name.`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `Type of cloud service provider.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC ID.`,
				},
				resource.Attribute{
					Name:        "vpc_reg",
					Description: `VPC Region.`,
				},
				resource.Attribute{
					Name:        "gw_size",
					Description: `Instance type.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `Range of the subnet where the spoke gateway is launched.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP address of the spoke gateway created.`,
				},
				resource.Attribute{
					Name:        "allocate_new_eip",
					Description: `Description: "Whether the eip is newly allocated or not.`,
				},
				resource.Attribute{
					Name:        "transit_gw",
					Description: `The transit gateway that the spoke gateway is attached to.`,
				},
				resource.Attribute{
					Name:        "tag_list",
					Description: `Instance tag of cloud provider. Only supported for AWS provider.`,
				},
				resource.Attribute{
					Name:        "insane_mode",
					Description: `Enable/Disable Insane Mode for Spoke Gateway.`,
				},
				resource.Attribute{
					Name:        "insane_mode_az",
					Description: `AZ of subnet being created for Insane Mode Spoke Gateway. Required if insane_mode is enabled for aws cloud.`,
				},
				resource.Attribute{
					Name:        "enable_active_mesh",
					Description: `Enable/Disable Active Mesh Mode for Spoke Gateway.`,
				},
				resource.Attribute{
					Name:        "enable_vpc_dns_server",
					Description: `Enable/Disalbe vpc_dns_server for Gateway.`,
				},
				resource.Attribute{
					Name:        "enable_encrypt_volume",
					Description: `Enable encrypt gateway EBS volume. Only supported for AWS provider.`,
				},
				resource.Attribute{
					Name:        "customized_spoke_vpc_routes",
					Description: `A list of comma separated CIDRs to be customized for the spoke VPC routes. When configured, it will replace all learned routes in VPC routing tables, including RFC1918 and non-RFC1918 CIDRs. It applies to this spoke gateway only​.`,
				},
				resource.Attribute{
					Name:        "filtered_spoke_vpc_routes",
					Description: `A list of comma separated CIDRs to be filtered from the spoke VPC route table. When configured, filtering CIDR(s) or it’s subnet will be deleted from VPC routing tables as well as from spoke gateway’s routing table. It applies to this spoke gateway only.`,
				},
				resource.Attribute{
					Name:        "included_advertised_spoke_routes",
					Description: `A list of comma separated CIDRs to be advertised to on-prem as 'Included CIDR List'. When configured, it will replace all advertised routes from this VPC.`,
				},
				resource.Attribute{
					Name:        "cloud_instance_id",
					Description: `Cloud instance ID`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "gw_name",
					Description: `Aviatrix spoke gateway name.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `Aviatrix account name.`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `Type of cloud service provider.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC ID.`,
				},
				resource.Attribute{
					Name:        "vpc_reg",
					Description: `VPC Region.`,
				},
				resource.Attribute{
					Name:        "gw_size",
					Description: `Instance type.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `Range of the subnet where the spoke gateway is launched.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP address of the spoke gateway created.`,
				},
				resource.Attribute{
					Name:        "allocate_new_eip",
					Description: `Description: "Whether the eip is newly allocated or not.`,
				},
				resource.Attribute{
					Name:        "transit_gw",
					Description: `The transit gateway that the spoke gateway is attached to.`,
				},
				resource.Attribute{
					Name:        "tag_list",
					Description: `Instance tag of cloud provider. Only supported for AWS provider.`,
				},
				resource.Attribute{
					Name:        "insane_mode",
					Description: `Enable/Disable Insane Mode for Spoke Gateway.`,
				},
				resource.Attribute{
					Name:        "insane_mode_az",
					Description: `AZ of subnet being created for Insane Mode Spoke Gateway. Required if insane_mode is enabled for aws cloud.`,
				},
				resource.Attribute{
					Name:        "enable_active_mesh",
					Description: `Enable/Disable Active Mesh Mode for Spoke Gateway.`,
				},
				resource.Attribute{
					Name:        "enable_vpc_dns_server",
					Description: `Enable/Disalbe vpc_dns_server for Gateway.`,
				},
				resource.Attribute{
					Name:        "enable_encrypt_volume",
					Description: `Enable encrypt gateway EBS volume. Only supported for AWS provider.`,
				},
				resource.Attribute{
					Name:        "customized_spoke_vpc_routes",
					Description: `A list of comma separated CIDRs to be customized for the spoke VPC routes. When configured, it will replace all learned routes in VPC routing tables, including RFC1918 and non-RFC1918 CIDRs. It applies to this spoke gateway only​.`,
				},
				resource.Attribute{
					Name:        "filtered_spoke_vpc_routes",
					Description: `A list of comma separated CIDRs to be filtered from the spoke VPC route table. When configured, filtering CIDR(s) or it’s subnet will be deleted from VPC routing tables as well as from spoke gateway’s routing table. It applies to this spoke gateway only.`,
				},
				resource.Attribute{
					Name:        "included_advertised_spoke_routes",
					Description: `A list of comma separated CIDRs to be advertised to on-prem as 'Included CIDR List'. When configured, it will replace all advertised routes from this VPC.`,
				},
				resource.Attribute{
					Name:        "cloud_instance_id",
					Description: `Cloud instance ID`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_transit_gateway",
			Category:         "Data Sources",
			ShortDescription: `Gets the Aviatrix transit gateway.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "gw_name",
					Description: `(Required) Transit gateway name. This can be used for getting transit gateway.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `(Optional) Account name. This can be used for logging in to CloudN console or UserConnect controller. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "gw_name",
					Description: `Aviatrix transit gateway name.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `Aviatrix account name.`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `Type of cloud service provider.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC ID.`,
				},
				resource.Attribute{
					Name:        "vpc_reg",
					Description: `VPC Region.`,
				},
				resource.Attribute{
					Name:        "gw_size",
					Description: `Instance type.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `Range of the subnet where the transit gateway is launched.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP address of the Gateway created.`,
				},
				resource.Attribute{
					Name:        "allocate_new_eip",
					Description: `Whether the eip is newly allocated or not.`,
				},
				resource.Attribute{
					Name:        "single_az_ha",
					Description: `Enable/Disable this feature.`,
				},
				resource.Attribute{
					Name:        "tag_list",
					Description: `Instance tag of cloud provider. Only supported for AWS provider.`,
				},
				resource.Attribute{
					Name:        "enable_hybrid_connection",
					Description: `Sign of readiness for TGW connection.`,
				},
				resource.Attribute{
					Name:        "connected_transit",
					Description: `Connected Transit status.`,
				},
				resource.Attribute{
					Name:        "insane_mode",
					Description: `Enable/Disable Insane Mode for Spoke Gateway.`,
				},
				resource.Attribute{
					Name:        "insane_mode_az",
					Description: `AZ of subnet being created for Insane Mode Spoke Gateway. Required if insane_mode is enabled for aws cloud.`,
				},
				resource.Attribute{
					Name:        "enable_firenet",
					Description: `Whether firenet interfaces is enabled.`,
				},
				resource.Attribute{
					Name:        "enable_active_mesh",
					Description: `Enable/Disable active mesh mode for Transit Gateway.`,
				},
				resource.Attribute{
					Name:        "enable_vpc_dns_server",
					Description: `Enable/Disable vpc_dns_server for Gateway. Only supports AWS.`,
				},
				resource.Attribute{
					Name:        "enable_advertise_transit_cidr",
					Description: `Enable/Disable advertise transit VPC network CIDR.`,
				},
				resource.Attribute{
					Name:        "bgp_manual_spoke_advertise_cidrs",
					Description: `Intended CIDR list to advertise to VGW.`,
				},
				resource.Attribute{
					Name:        "enable_encrypt_volume",
					Description: `Enable/Disable encrypt gateway EBS volume. Only supported for AWS provider.`,
				},
				resource.Attribute{
					Name:        "customized_spoke_vpc_routes",
					Description: `A list of comma separated CIDRs to be customized for the spoke VPC routes. When configured, it will replace all learned routes in VPC routing tables, including RFC1918 and non-RFC1918 CIDRs. It applies to all spoke gateways attached to this transit gateway.`,
				},
				resource.Attribute{
					Name:        "filtered_spoke_vpc_routes",
					Description: `A list of comma separated CIDRs to be filtered from the spoke VPC route table. When configured, filtering CIDR(s) or it’s subnet will be deleted from VPC routing tables as well as from spoke gateway’s routing table. It applies to all spoke gateways attached to this transit gateway.`,
				},
				resource.Attribute{
					Name:        "excluded_advertised_spoke_routes",
					Description: `A list of comma separated CIDRs to be advertised to on-prem as 'Excluded CIDR List'. When configured, it inspects all the advertised CIDRs from its spoke gateways and remove those included in the 'Excluded CIDR List'.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "gw_name",
					Description: `Aviatrix transit gateway name.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `Aviatrix account name.`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `Type of cloud service provider.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC ID.`,
				},
				resource.Attribute{
					Name:        "vpc_reg",
					Description: `VPC Region.`,
				},
				resource.Attribute{
					Name:        "gw_size",
					Description: `Instance type.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `Range of the subnet where the transit gateway is launched.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP address of the Gateway created.`,
				},
				resource.Attribute{
					Name:        "allocate_new_eip",
					Description: `Whether the eip is newly allocated or not.`,
				},
				resource.Attribute{
					Name:        "single_az_ha",
					Description: `Enable/Disable this feature.`,
				},
				resource.Attribute{
					Name:        "tag_list",
					Description: `Instance tag of cloud provider. Only supported for AWS provider.`,
				},
				resource.Attribute{
					Name:        "enable_hybrid_connection",
					Description: `Sign of readiness for TGW connection.`,
				},
				resource.Attribute{
					Name:        "connected_transit",
					Description: `Connected Transit status.`,
				},
				resource.Attribute{
					Name:        "insane_mode",
					Description: `Enable/Disable Insane Mode for Spoke Gateway.`,
				},
				resource.Attribute{
					Name:        "insane_mode_az",
					Description: `AZ of subnet being created for Insane Mode Spoke Gateway. Required if insane_mode is enabled for aws cloud.`,
				},
				resource.Attribute{
					Name:        "enable_firenet",
					Description: `Whether firenet interfaces is enabled.`,
				},
				resource.Attribute{
					Name:        "enable_active_mesh",
					Description: `Enable/Disable active mesh mode for Transit Gateway.`,
				},
				resource.Attribute{
					Name:        "enable_vpc_dns_server",
					Description: `Enable/Disable vpc_dns_server for Gateway. Only supports AWS.`,
				},
				resource.Attribute{
					Name:        "enable_advertise_transit_cidr",
					Description: `Enable/Disable advertise transit VPC network CIDR.`,
				},
				resource.Attribute{
					Name:        "bgp_manual_spoke_advertise_cidrs",
					Description: `Intended CIDR list to advertise to VGW.`,
				},
				resource.Attribute{
					Name:        "enable_encrypt_volume",
					Description: `Enable/Disable encrypt gateway EBS volume. Only supported for AWS provider.`,
				},
				resource.Attribute{
					Name:        "customized_spoke_vpc_routes",
					Description: `A list of comma separated CIDRs to be customized for the spoke VPC routes. When configured, it will replace all learned routes in VPC routing tables, including RFC1918 and non-RFC1918 CIDRs. It applies to all spoke gateways attached to this transit gateway.`,
				},
				resource.Attribute{
					Name:        "filtered_spoke_vpc_routes",
					Description: `A list of comma separated CIDRs to be filtered from the spoke VPC route table. When configured, filtering CIDR(s) or it’s subnet will be deleted from VPC routing tables as well as from spoke gateway’s routing table. It applies to all spoke gateways attached to this transit gateway.`,
				},
				resource.Attribute{
					Name:        "excluded_advertised_spoke_routes",
					Description: `A list of comma separated CIDRs to be advertised to on-prem as 'Excluded CIDR List'. When configured, it inspects all the advertised CIDRs from its spoke gateways and remove those included in the 'Excluded CIDR List'.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"aviatrix_account":                    0,
		"aviatrix_caller_identity":            1,
		"aviatrix_firenet":                    2,
		"aviatrix_firenet_vendor_integration": 3,
		"aviatrix_gateway":                    4,
		"aviatrix_spoke_gateway":              5,
		"aviatrix_transit_gateway":            6,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
