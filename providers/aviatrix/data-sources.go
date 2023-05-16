package aviatrix

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_account",
			Category:         "Accounts",
			ShortDescription: `Gets an Aviatrix cloud account's details.`,
			Description:      ``,
			Keywords: []string{
				"accounts",
				"account",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_name",
					Description: `(Required) Account name. This can be used for logging in to CloudN console or UserConnect controller. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `Type of cloud service provider. ### AWS`,
				},
				resource.Attribute{
					Name:        "aws_account_number",
					Description: `AWS Account number.`,
				},
				resource.Attribute{
					Name:        "aws_role_arn",
					Description: `AWS App role ARN.`,
				},
				resource.Attribute{
					Name:        "aws_role_ec2",
					Description: `AWS EC2 role ARN.`,
				},
				resource.Attribute{
					Name:        "aws_gateway_role_app",
					Description: `A separate AWS App role ARN to assign to gateways created by the controller. Available as of provider version R2.19+.`,
				},
				resource.Attribute{
					Name:        "aws_gateway_role_ec2",
					Description: `A separate AWS EC2 role ARN to assign to gateways created by the controller. Available as of provider version R2.19+. ### Azure`,
				},
				resource.Attribute{
					Name:        "arm_subscription_id",
					Description: `Azure ARM Subscription ID. ### GCP`,
				},
				resource.Attribute{
					Name:        "gcloud_project_id",
					Description: `GCloud Project ID. ### AzureGov Cloud`,
				},
				resource.Attribute{
					Name:        "azuregov_subscription_id",
					Description: `AzureGov ARM Subscription ID. ### AWSGov Cloud`,
				},
				resource.Attribute{
					Name:        "awsgov_account_number",
					Description: `AWSGov Account number.`,
				},
				resource.Attribute{
					Name:        "awsgov_iam",
					Description: `If enabled, ` + "`" + `awsgov_role_app` + "`" + ` and ` + "`" + `awschina_role_ec2` + "`" + ` will be set. Available as of provider version R2.19+.`,
				},
				resource.Attribute{
					Name:        "awsgov_role_app",
					Description: `AWSGov App role ARN. Available as of provider version R2.19+.`,
				},
				resource.Attribute{
					Name:        "awsgov_role_ec2",
					Description: `AWSGov EC2 role ARN. Available as of provider version R2.19+. ### AWSChina Cloud`,
				},
				resource.Attribute{
					Name:        "awschina_account_number",
					Description: `AWSChina Account number. Available as of provider version R2.19+.`,
				},
				resource.Attribute{
					Name:        "awschina_iam",
					Description: `If enabled, ` + "`" + `awschina_role_app` + "`" + ` and ` + "`" + `awschina_role_ec2` + "`" + ` will be set. Available as of provider version R2.19+.`,
				},
				resource.Attribute{
					Name:        "awschina_role_app",
					Description: `AWSChina App role ARN. Available as of provider version R2.19+.`,
				},
				resource.Attribute{
					Name:        "awschina_role_ec2",
					Description: `AWSChina EC2 role ARN. Available as of provider version R2.19+. ### AzureChina Cloud`,
				},
				resource.Attribute{
					Name:        "azurechina_subscription_id",
					Description: `AzureChina ARM Subscription ID. Available as of provider version R2.19+. ### Alibaba Cloud`,
				},
				resource.Attribute{
					Name:        "alicloud_account_id",
					Description: `Alibaba Cloud Account ID. ### AWS Top Secret Cloud`,
				},
				resource.Attribute{
					Name:        "awsts_account_number",
					Description: `AWS Top Secret Region Account Number. Required when creating an account in AWS Top Secret Region. Available as of provider version R2.19.5+.`,
				},
				resource.Attribute{
					Name:        "awsts_cap_url",
					Description: `AWS Top Secret Region CAP Url. Required when creating an account in AWS Top Secret Region. Available as of provider version R2.19.5+.`,
				},
				resource.Attribute{
					Name:        "awsts_cap_agency",
					Description: `AWS Top Secret Region CAP Agency. Required when creating an account in AWS Top Secret Region. Available as of provider version R2.19.5+.`,
				},
				resource.Attribute{
					Name:        "awsts_cap_mission",
					Description: `AWS Top Secret Region Mission. Required when creating an account in AWS Top Secret Region. Available as of provider version R2.19.5+.`,
				},
				resource.Attribute{
					Name:        "awsts_cap_role_name",
					Description: `AWS Top Secret Region Role Name. Required when creating an account in AWS Top Secret Region. Available as of provider version R2.19.5+. ` + "`" + `awsts_cap_cert_path` + "`" + ` - AWS Top Secret Region CAP Certificate file name on the controller. Available as of provider R2.19.5+.`,
				},
				resource.Attribute{
					Name:        "awsts_cap_cert_key_path",
					Description: `AWS Top Secret Region CAP Certificate Key file name on the controller. Available as of provider R2.19.5+.`,
				},
				resource.Attribute{
					Name:        "aws_ca_cert_path",
					Description: `AWS Top Secret Region or Secret Region Custom Certificate Authority file name on the controller. Available as of provider R2.19.5+. ### AWS Secret Cloud`,
				},
				resource.Attribute{
					Name:        "awss_account_number",
					Description: `AWS Secret Region Account Number. Required when creating an account in AWS Secret Region. Available as of provider version R2.19.5+.`,
				},
				resource.Attribute{
					Name:        "awss_cap_url",
					Description: `AWS Secret Region CAP Url. Required when creating an account in AWS Secret Region. Available as of provider version R2.19.5+.`,
				},
				resource.Attribute{
					Name:        "awss_cap_agency",
					Description: `AWS Secret Region CAP Agency. Required when creating an account in AWS Secret Region. Available as of provider version R2.19.5+.`,
				},
				resource.Attribute{
					Name:        "awss_cap_account_name",
					Description: `AWS Secret Region Account Name. Required when creating an account in AWS Secret Region. Available as of provider version R2.19.5+.`,
				},
				resource.Attribute{
					Name:        "awss_cap_role_name",
					Description: `AWS Secret Region Role Name. Required when creating an account in AWS Secret Region. Available as of provider version R2.19.5+.`,
				},
				resource.Attribute{
					Name:        "awss_cap_cert_path",
					Description: `AWS Secret Region CAP Certificate file name on the controller. Available as of provider R2.19.5+.`,
				},
				resource.Attribute{
					Name:        "awss_cap_cert_key_path",
					Description: `AWS Secret Region CAP Certificate Key file name on the controller. Available as of provider R2.19.5+.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_type",
					Description: `Type of cloud service provider. ### AWS`,
				},
				resource.Attribute{
					Name:        "aws_account_number",
					Description: `AWS Account number.`,
				},
				resource.Attribute{
					Name:        "aws_role_arn",
					Description: `AWS App role ARN.`,
				},
				resource.Attribute{
					Name:        "aws_role_ec2",
					Description: `AWS EC2 role ARN.`,
				},
				resource.Attribute{
					Name:        "aws_gateway_role_app",
					Description: `A separate AWS App role ARN to assign to gateways created by the controller. Available as of provider version R2.19+.`,
				},
				resource.Attribute{
					Name:        "aws_gateway_role_ec2",
					Description: `A separate AWS EC2 role ARN to assign to gateways created by the controller. Available as of provider version R2.19+. ### Azure`,
				},
				resource.Attribute{
					Name:        "arm_subscription_id",
					Description: `Azure ARM Subscription ID. ### GCP`,
				},
				resource.Attribute{
					Name:        "gcloud_project_id",
					Description: `GCloud Project ID. ### AzureGov Cloud`,
				},
				resource.Attribute{
					Name:        "azuregov_subscription_id",
					Description: `AzureGov ARM Subscription ID. ### AWSGov Cloud`,
				},
				resource.Attribute{
					Name:        "awsgov_account_number",
					Description: `AWSGov Account number.`,
				},
				resource.Attribute{
					Name:        "awsgov_iam",
					Description: `If enabled, ` + "`" + `awsgov_role_app` + "`" + ` and ` + "`" + `awschina_role_ec2` + "`" + ` will be set. Available as of provider version R2.19+.`,
				},
				resource.Attribute{
					Name:        "awsgov_role_app",
					Description: `AWSGov App role ARN. Available as of provider version R2.19+.`,
				},
				resource.Attribute{
					Name:        "awsgov_role_ec2",
					Description: `AWSGov EC2 role ARN. Available as of provider version R2.19+. ### AWSChina Cloud`,
				},
				resource.Attribute{
					Name:        "awschina_account_number",
					Description: `AWSChina Account number. Available as of provider version R2.19+.`,
				},
				resource.Attribute{
					Name:        "awschina_iam",
					Description: `If enabled, ` + "`" + `awschina_role_app` + "`" + ` and ` + "`" + `awschina_role_ec2` + "`" + ` will be set. Available as of provider version R2.19+.`,
				},
				resource.Attribute{
					Name:        "awschina_role_app",
					Description: `AWSChina App role ARN. Available as of provider version R2.19+.`,
				},
				resource.Attribute{
					Name:        "awschina_role_ec2",
					Description: `AWSChina EC2 role ARN. Available as of provider version R2.19+. ### AzureChina Cloud`,
				},
				resource.Attribute{
					Name:        "azurechina_subscription_id",
					Description: `AzureChina ARM Subscription ID. Available as of provider version R2.19+. ### Alibaba Cloud`,
				},
				resource.Attribute{
					Name:        "alicloud_account_id",
					Description: `Alibaba Cloud Account ID. ### AWS Top Secret Cloud`,
				},
				resource.Attribute{
					Name:        "awsts_account_number",
					Description: `AWS Top Secret Region Account Number. Required when creating an account in AWS Top Secret Region. Available as of provider version R2.19.5+.`,
				},
				resource.Attribute{
					Name:        "awsts_cap_url",
					Description: `AWS Top Secret Region CAP Url. Required when creating an account in AWS Top Secret Region. Available as of provider version R2.19.5+.`,
				},
				resource.Attribute{
					Name:        "awsts_cap_agency",
					Description: `AWS Top Secret Region CAP Agency. Required when creating an account in AWS Top Secret Region. Available as of provider version R2.19.5+.`,
				},
				resource.Attribute{
					Name:        "awsts_cap_mission",
					Description: `AWS Top Secret Region Mission. Required when creating an account in AWS Top Secret Region. Available as of provider version R2.19.5+.`,
				},
				resource.Attribute{
					Name:        "awsts_cap_role_name",
					Description: `AWS Top Secret Region Role Name. Required when creating an account in AWS Top Secret Region. Available as of provider version R2.19.5+. ` + "`" + `awsts_cap_cert_path` + "`" + ` - AWS Top Secret Region CAP Certificate file name on the controller. Available as of provider R2.19.5+.`,
				},
				resource.Attribute{
					Name:        "awsts_cap_cert_key_path",
					Description: `AWS Top Secret Region CAP Certificate Key file name on the controller. Available as of provider R2.19.5+.`,
				},
				resource.Attribute{
					Name:        "aws_ca_cert_path",
					Description: `AWS Top Secret Region or Secret Region Custom Certificate Authority file name on the controller. Available as of provider R2.19.5+. ### AWS Secret Cloud`,
				},
				resource.Attribute{
					Name:        "awss_account_number",
					Description: `AWS Secret Region Account Number. Required when creating an account in AWS Secret Region. Available as of provider version R2.19.5+.`,
				},
				resource.Attribute{
					Name:        "awss_cap_url",
					Description: `AWS Secret Region CAP Url. Required when creating an account in AWS Secret Region. Available as of provider version R2.19.5+.`,
				},
				resource.Attribute{
					Name:        "awss_cap_agency",
					Description: `AWS Secret Region CAP Agency. Required when creating an account in AWS Secret Region. Available as of provider version R2.19.5+.`,
				},
				resource.Attribute{
					Name:        "awss_cap_account_name",
					Description: `AWS Secret Region Account Name. Required when creating an account in AWS Secret Region. Available as of provider version R2.19.5+.`,
				},
				resource.Attribute{
					Name:        "awss_cap_role_name",
					Description: `AWS Secret Region Role Name. Required when creating an account in AWS Secret Region. Available as of provider version R2.19.5+.`,
				},
				resource.Attribute{
					Name:        "awss_cap_cert_path",
					Description: `AWS Secret Region CAP Certificate file name on the controller. Available as of provider R2.19.5+.`,
				},
				resource.Attribute{
					Name:        "awss_cap_cert_key_path",
					Description: `AWS Secret Region CAP Certificate Key file name on the controller. Available as of provider R2.19.5+.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_caller_identity",
			Category:         "Settings",
			ShortDescription: `Gets the Aviatrix caller identity.`,
			Description:      ``,
			Keywords: []string{
				"settings",
				"caller",
				"identity",
			},
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
			Type:             "aviatrix_aviatrix_controller_metadata",
			Category:         "Settings",
			ShortDescription: `Gets the Aviatrix controller metadata.`,
			Description:      ``,
			Keywords: []string{
				"settings",
				"controller",
				"metadata",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `Controller region.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `Controller VPC ID.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `Controller instance ID.`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `Controller cloud type, only supported for AWS and GCP now.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `Controller region.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `Controller VPC ID.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `Controller instance ID.`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `Controller cloud type, only supported for AWS and GCP now.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_device_interfaces",
			Category:         "CloudN",
			ShortDescription: `Gets the WAN primary interfaces and IPs for a device.`,
			Description:      ``,
			Keywords: []string{
				"cloudn",
				"device",
				"interfaces",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "device_name",
					Description: `(Required) Device name. ## Attribute Reference In addition to the argument above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "wan_interfaces",
					Description: `List of WAN interfaces.`,
				},
				resource.Attribute{
					Name:        "wan_primary_interface",
					Description: `Name of the WAN primary interface.`,
				},
				resource.Attribute{
					Name:        "wan_primary_interface_public_ip",
					Description: `The WAN Primary interface public IP.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "wan_interfaces",
					Description: `List of WAN interfaces.`,
				},
				resource.Attribute{
					Name:        "wan_primary_interface",
					Description: `Name of the WAN primary interface.`,
				},
				resource.Attribute{
					Name:        "wan_primary_interface_public_ip",
					Description: `The WAN Primary interface public IP.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_edge_gateway_wan_interface_discovery",
			Category:         "Multi-Cloud Transit",
			ShortDescription: `Gets the Aviatrix Edge gateway WAN interface public IP address.`,
			Description:      ``,
			Keywords: []string{
				"multi",
				"cloud",
				"transit",
				"edge",
				"gateway",
				"wan",
				"interface",
				"discovery",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "gw_name",
					Description: `(Required) Edge gateway name.`,
				},
				resource.Attribute{
					Name:        "wan_interface_name",
					Description: `(Required) Name of the WAN interface to be discovered. ## Attribute Reference In addition to the argument above, the following attributes is exported:`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `Public IP of the Edge gateway's WAN interface.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_address",
					Description: `Public IP of the Edge gateway's WAN interface.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_firenet",
			Category:         "Firewall Network",
			ShortDescription: `Gets an Aviatrix FireNet's details.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"network",
				"firenet",
			},
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
					Name:        "egress_static_cidrs",
					Description: `List of egress static CIDRs.`,
				},
				resource.Attribute{
					Name:        "tgw_segmentation_for_egress_enabled",
					Description: `Enable TGW segmentation for egress.`,
				},
				resource.Attribute{
					Name:        "hashing_algorithm",
					Description: `(Optional) Hashing algorithm to load balance traffic across the firewall.`,
				},
				resource.Attribute{
					Name:        "keep_alive_via_lan_interface_enabled",
					Description: `(Optional) Enable Keep Alive via Firewall LAN Interface. The following arguments are deprecated:`,
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
					Name:        "egress_static_cidrs",
					Description: `List of egress static CIDRs.`,
				},
				resource.Attribute{
					Name:        "tgw_segmentation_for_egress_enabled",
					Description: `Enable TGW segmentation for egress.`,
				},
				resource.Attribute{
					Name:        "hashing_algorithm",
					Description: `(Optional) Hashing algorithm to load balance traffic across the firewall.`,
				},
				resource.Attribute{
					Name:        "keep_alive_via_lan_interface_enabled",
					Description: `(Optional) Enable Keep Alive via Firewall LAN Interface. The following arguments are deprecated:`,
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
			Type:             "aviatrix_aviatrix_firenet_firewall_manager",
			Category:         "Firewall Network",
			ShortDescription: `Performs 'save' or 'sync' for Aviatrix FireNet firewall manager.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"network",
				"firenet",
				"manager",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) VPC ID.`,
				},
				resource.Attribute{
					Name:        "gateway_name",
					Description: `(Required) The FireNet gateway name.`,
				},
				resource.Attribute{
					Name:        "vendor_type",
					Description: `(Required) Vendor type. Valid values: "Generic" and "Palo Alto Networks Panorama".`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `(Optional) The public IP address of the Panorama instance. Required for vendor type "Palo Alto Networks Panorama".`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) Panorama login name for API calls from the Controller. Required for vendor type "Palo Alto Networks Panorama".`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Panorama login password for API calls. Required for vendor type "Palo Alto Networks Panorama".`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Optional) Panorama template for each FireNet gateway. Required for vendor type "Palo Alto Networks Panorama".`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `(Optional) Panorama template stack for each FireNet gateway. Required for vendor type "Palo Alto Networks Panorama".`,
				},
				resource.Attribute{
					Name:        "route_table",
					Description: `(Optional) The name of firewall virtual router to program. If left unspecified, the Controller programs the Panorama template’s first router.`,
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
			Type:             "aviatrix_aviatrix_firenet_vendor_integration",
			Category:         "Firewall Network",
			ShortDescription: `Performs 'save' or 'sync' for vendor integration purposes for Aviatrix FireNet.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"network",
				"firenet",
				"vendor",
				"integration",
			},
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
					Description: `(Required) Select PAN. Valid values: "Generic", "Palo Alto Networks VM-Series", "Aviatrix FQDN Gateway" and "Fortinet FortiGate".`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `(Optional) The IP address of the firewall management interface for API calls from the Aviatrix Controller. If not set, the public IP of the firewall instance will be used. If the private IP is provided, please make sure that the controller can access the firewall.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) Firewall login name for API calls from the Controller. Required for vendor type "Generic", "Palo Alto Networks VM-Series" and "Aviatrix FQDN Gateway".`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Firewall login password for API calls. Required for vendor type "Generic", "Palo Alto Networks VM-Series" and "Aviatrix FQDN Gateway".`,
				},
				resource.Attribute{
					Name:        "api_token",
					Description: `(Optional) API token for API calls. Required and valid only for vendor type "Fortinet FortiGate".`,
				},
				resource.Attribute{
					Name:        "private_key_file",
					Description: `(Optional) Private key file. Valid only for vendor type "Check Point Cloud Guard". Use the ` + "`" + `file` + "`" + ` function to read from a file.`,
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
			Type:             "aviatrix_aviatrix_firewall",
			Category:         "Security",
			ShortDescription: `Gets the Aviatrix Firewall.`,
			Description:      ``,
			Keywords: []string{
				"security",
				"firewall",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "gw_name",
					Description: `(Required) Name of the gateway associated with the firewall. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "base_policy",
					Description: `The firewall's base policy.`,
				},
				resource.Attribute{
					Name:        "base_log_enabled",
					Description: `Indicates whether logging is enabled or not.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `List of policies associated with the firewall.`,
				},
				resource.Attribute{
					Name:        "src_ip",
					Description: `CIDRs separated by a comma or tag names such 'HR' or 'marketing' etc.`,
				},
				resource.Attribute{
					Name:        "dst_ip",
					Description: `CIDRs separated by a comma or tag names such 'HR' or 'marketing' etc.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `` + "`" + `all` + "`" + `, ` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `, ` + "`" + `sctp` + "`" + `, ` + "`" + `rdp` + "`" + ` or ` + "`" + `dccp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `A single port or a range of port numbers.`,
				},
				resource.Attribute{
					Name:        "log_enabled",
					Description: `Indicates whether logging is enabled or not.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "base_policy",
					Description: `The firewall's base policy.`,
				},
				resource.Attribute{
					Name:        "base_log_enabled",
					Description: `Indicates whether logging is enabled or not.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `List of policies associated with the firewall.`,
				},
				resource.Attribute{
					Name:        "src_ip",
					Description: `CIDRs separated by a comma or tag names such 'HR' or 'marketing' etc.`,
				},
				resource.Attribute{
					Name:        "dst_ip",
					Description: `CIDRs separated by a comma or tag names such 'HR' or 'marketing' etc.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `` + "`" + `all` + "`" + `, ` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `, ` + "`" + `sctp` + "`" + `, ` + "`" + `rdp` + "`" + ` or ` + "`" + `dccp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `A single port or a range of port numbers.`,
				},
				resource.Attribute{
					Name:        "log_enabled",
					Description: `Indicates whether logging is enabled or not.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_firewall_instance_images",
			Category:         "Firewall Network",
			ShortDescription: `Gets the Aviatrix firewall instance images information.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"network",
				"instance",
				"images",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) VPC ID. Example: AWS: "vpc-abcd1234", GCP: "vpc-gcp-test~-~project_id", Azure: "vnet_name:rg_name:resource_guid", OCI: "vpc-oracle-test1". ## Attribute Reference In addition to the argument above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "firewall_images",
					Description: `List of firewall images.`,
				},
				resource.Attribute{
					Name:        "firewall_image",
					Description: `Name of the firewall image.`,
				},
				resource.Attribute{
					Name:        "firewall_image_version",
					Description: `List of firewall image versions.`,
				},
				resource.Attribute{
					Name:        "firewall_size",
					Description: `List of firewall instance sizes.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "firewall_images",
					Description: `List of firewall images.`,
				},
				resource.Attribute{
					Name:        "firewall_image",
					Description: `Name of the firewall image.`,
				},
				resource.Attribute{
					Name:        "firewall_image_version",
					Description: `List of firewall image versions.`,
				},
				resource.Attribute{
					Name:        "firewall_size",
					Description: `List of firewall instance sizes.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_gateway",
			Category:         "Gateway",
			ShortDescription: `Gets an Aviatrix gateway's details.`,
			Description:      ``,
			Keywords: []string{
				"gateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "gw_name",
					Description: `(Required) Gateway name. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `Aviatrix account name.`,
				},
				resource.Attribute{
					Name:        "additional_cidrs",
					Description: `A list of destination CIDR ranges that will also go through the VPN tunnel when Split Tunnel Mode is enabled.`,
				},
				resource.Attribute{
					Name:        "additional_cidrs_designated_gateway",
					Description: `A list of CIDR ranges separated by comma to configure when 'designated_gateway' feature is enabled.`,
				},
				resource.Attribute{
					Name:        "allocate_new_eip",
					Description: `When value is false, an idle address in Elastic IP pool is reused for this gateway. Otherwise, a new Elastic IP is allocated and used for this gateway.`,
				},
				resource.Attribute{
					Name:        "cloud_instance_id",
					Description: `Instance ID of the gateway.`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `Type of cloud service provider.`,
				},
				resource.Attribute{
					Name:        "duo_api_hostname",
					Description: `API hostname for DUO auth mode.`,
				},
				resource.Attribute{
					Name:        "duo_integration_key",
					Description: `Integration key for DUO auth mode.`,
				},
				resource.Attribute{
					Name:        "duo_push_mode",
					Description: `Push mode for DUO auth.`,
				},
				resource.Attribute{
					Name:        "elb_dns_name",
					Description: `ELB DNS Name.`,
				},
				resource.Attribute{
					Name:        "elb_name",
					Description: `Name of the ELB created.`,
				},
				resource.Attribute{
					Name:        "enable_designated_gateway",
					Description: `Status of Designated Gateway feature for Gateway.`,
				},
				resource.Attribute{
					Name:        "enable_elb",
					Description: `Status of ELB for the gateway.`,
				},
				resource.Attribute{
					Name:        "enable_encrypt_volume",
					Description: `Enable encrypt gateway EBS volume. Only supported for AWS provider.`,
				},
				resource.Attribute{
					Name:        "enable_ldap",
					Description: `Status of LDAP enabled or not.`,
				},
				resource.Attribute{
					Name:        "enable_vpc_dns_server",
					Description: `Status of VPC Dns Server for Gateway.`,
				},
				resource.Attribute{
					Name:        "enable_vpn_nat",
					Description: `Status of VPN NAT.`,
				},
				resource.Attribute{
					Name:        "gw_size",
					Description: `Size of gateway Instance.`,
				},
				resource.Attribute{
					Name:        "gw_name",
					Description: `Aviatrix gateway name.`,
				},
				resource.Attribute{
					Name:        "insane_mode",
					Description: `Status of Insane Mode for Gateway.`,
				},
				resource.Attribute{
					Name:        "insane_mode_az",
					Description: `AZ of subnet being created for Insane Mode gateway.`,
				},
				resource.Attribute{
					Name:        "ldap_bind_dn",
					Description: `LDAP bind DN.`,
				},
				resource.Attribute{
					Name:        "ldap_base_dn",
					Description: `LDAP base DN.`,
				},
				resource.Attribute{
					Name:        "ldap_server",
					Description: `LDAP server address.`,
				},
				resource.Attribute{
					Name:        "ldap_username_attribute",
					Description: `LDAP user attribute.`,
				},
				resource.Attribute{
					Name:        "max_vpn_conn",
					Description: `Maximum connection of VPN access. Valid for VPN gateway only. If not set, '100' will be default value.`,
				},
				resource.Attribute{
					Name:        "name_servers",
					Description: `A list of DNS servers used to resolve domain names by a connected VPN user when Split Tunnel Mode is enabled.`,
				},
				resource.Attribute{
					Name:        "okta_url",
					Description: `URL for Okta auth mode.`,
				},
				resource.Attribute{
					Name:        "okta_username_suffix",
					Description: `Username suffix for Okta auth mode.`,
				},
				resource.Attribute{
					Name:        "otp_mode",
					Description: `Two step authentication mode.`,
				},
				resource.Attribute{
					Name:        "peering_ha_cloud_instance_id",
					Description: `Instance ID of the peering HA gateway.`,
				},
				resource.Attribute{
					Name:        "peering_ha_gw_name",
					Description: `Aviatrix gateway unique name of HA gateway.`,
				},
				resource.Attribute{
					Name:        "peering_ha_gw_size",
					Description: `Peering HA Gateway Size.`,
				},
				resource.Attribute{
					Name:        "peering_ha_insane_mode_az",
					Description: `AZ of subnet being created for Insane Mode Peering HA Gateway. Required if insane_mode is set.`,
				},
				resource.Attribute{
					Name:        "peering_ha_private_ip",
					Description: `Private IP address of HA gateway.`,
				},
				resource.Attribute{
					Name:        "peering_ha_public_ip",
					Description: `Public IP address that you want assigned to the HA peering instance.`,
				},
				resource.Attribute{
					Name:        "peering_ha_subnet",
					Description: `Public Subnet Information while creating Peering HA Gateway, only subnet is accepted. Required to create peering ha gateway if cloud_type = 1 or 8 (AWS or Azure).`,
				},
				resource.Attribute{
					Name:        "peering_ha_zone",
					Description: `Zone information for creating Peering HA Gateway. Required to create peering ha gateway if cloud_type = 4 (GCP).`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Private IP address of the Gateway created.`,
				},
				resource.Attribute{
					Name:        "public_dns_server",
					Description: `NS server used by the gateway.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP address of the Gateway created.`,
				},
				resource.Attribute{
					Name:        "saml_enabled",
					Description: `Status of SAML.`,
				},
				resource.Attribute{
					Name:        "search_domains",
					Description: `A list of domain names that will use the NameServer when a specific name is not in the destination when Split Tunnel Mode is enabled.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `Security group used for the gateway.`,
				},
				resource.Attribute{
					Name:        "single_az_ha",
					Description: `Status of Single AZ HA.`,
				},
				resource.Attribute{
					Name:        "single_ip_snat",
					Description: `Single IP Source NAT status for the container.`,
				},
				resource.Attribute{
					Name:        "split_tunnel",
					Description: `Status of split tunnel mode.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `A VPC Network address range selected from one of the available network ranges.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC-ID/VNet-Name of cloud provider.`,
				},
				resource.Attribute{
					Name:        "vpc_reg",
					Description: `Region of cloud provider.`,
				},
				resource.Attribute{
					Name:        "vpn_access",
					Description: `Status of user access through VPN to the container.`,
				},
				resource.Attribute{
					Name:        "vpn_cidr",
					Description: `VPN CIDR block for the container.`,
				},
				resource.Attribute{
					Name:        "vpn_protocol",
					Description: `ELB protocol for VPN gateway with ELB enabled.`,
				},
				resource.Attribute{
					Name:        "availability_domain",
					Description: `Availability domain for OCI.`,
				},
				resource.Attribute{
					Name:        "fault_domain",
					Description: `Fault domain for OCI.`,
				},
				resource.Attribute{
					Name:        "peering_ha_availability_domain",
					Description: `HA gateway availability domain for OCI.`,
				},
				resource.Attribute{
					Name:        "peering_ha_fault_domain",
					Description: `HA gateway fault domain for OCI.`,
				},
				resource.Attribute{
					Name:        "software_version",
					Description: `The software version of the gateway.`,
				},
				resource.Attribute{
					Name:        "image_version",
					Description: `The image version of the gateway.`,
				},
				resource.Attribute{
					Name:        "peering_ha_software_version",
					Description: `The software version of the HA gateway.`,
				},
				resource.Attribute{
					Name:        "peering_ha_image_version",
					Description: `The image version of the HA gateway. The following argument is deprecated:`,
				},
				resource.Attribute{
					Name:        "tag_list",
					Description: `Instance tag of cloud provider.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "account_name",
					Description: `Aviatrix account name.`,
				},
				resource.Attribute{
					Name:        "additional_cidrs",
					Description: `A list of destination CIDR ranges that will also go through the VPN tunnel when Split Tunnel Mode is enabled.`,
				},
				resource.Attribute{
					Name:        "additional_cidrs_designated_gateway",
					Description: `A list of CIDR ranges separated by comma to configure when 'designated_gateway' feature is enabled.`,
				},
				resource.Attribute{
					Name:        "allocate_new_eip",
					Description: `When value is false, an idle address in Elastic IP pool is reused for this gateway. Otherwise, a new Elastic IP is allocated and used for this gateway.`,
				},
				resource.Attribute{
					Name:        "cloud_instance_id",
					Description: `Instance ID of the gateway.`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `Type of cloud service provider.`,
				},
				resource.Attribute{
					Name:        "duo_api_hostname",
					Description: `API hostname for DUO auth mode.`,
				},
				resource.Attribute{
					Name:        "duo_integration_key",
					Description: `Integration key for DUO auth mode.`,
				},
				resource.Attribute{
					Name:        "duo_push_mode",
					Description: `Push mode for DUO auth.`,
				},
				resource.Attribute{
					Name:        "elb_dns_name",
					Description: `ELB DNS Name.`,
				},
				resource.Attribute{
					Name:        "elb_name",
					Description: `Name of the ELB created.`,
				},
				resource.Attribute{
					Name:        "enable_designated_gateway",
					Description: `Status of Designated Gateway feature for Gateway.`,
				},
				resource.Attribute{
					Name:        "enable_elb",
					Description: `Status of ELB for the gateway.`,
				},
				resource.Attribute{
					Name:        "enable_encrypt_volume",
					Description: `Enable encrypt gateway EBS volume. Only supported for AWS provider.`,
				},
				resource.Attribute{
					Name:        "enable_ldap",
					Description: `Status of LDAP enabled or not.`,
				},
				resource.Attribute{
					Name:        "enable_vpc_dns_server",
					Description: `Status of VPC Dns Server for Gateway.`,
				},
				resource.Attribute{
					Name:        "enable_vpn_nat",
					Description: `Status of VPN NAT.`,
				},
				resource.Attribute{
					Name:        "gw_size",
					Description: `Size of gateway Instance.`,
				},
				resource.Attribute{
					Name:        "gw_name",
					Description: `Aviatrix gateway name.`,
				},
				resource.Attribute{
					Name:        "insane_mode",
					Description: `Status of Insane Mode for Gateway.`,
				},
				resource.Attribute{
					Name:        "insane_mode_az",
					Description: `AZ of subnet being created for Insane Mode gateway.`,
				},
				resource.Attribute{
					Name:        "ldap_bind_dn",
					Description: `LDAP bind DN.`,
				},
				resource.Attribute{
					Name:        "ldap_base_dn",
					Description: `LDAP base DN.`,
				},
				resource.Attribute{
					Name:        "ldap_server",
					Description: `LDAP server address.`,
				},
				resource.Attribute{
					Name:        "ldap_username_attribute",
					Description: `LDAP user attribute.`,
				},
				resource.Attribute{
					Name:        "max_vpn_conn",
					Description: `Maximum connection of VPN access. Valid for VPN gateway only. If not set, '100' will be default value.`,
				},
				resource.Attribute{
					Name:        "name_servers",
					Description: `A list of DNS servers used to resolve domain names by a connected VPN user when Split Tunnel Mode is enabled.`,
				},
				resource.Attribute{
					Name:        "okta_url",
					Description: `URL for Okta auth mode.`,
				},
				resource.Attribute{
					Name:        "okta_username_suffix",
					Description: `Username suffix for Okta auth mode.`,
				},
				resource.Attribute{
					Name:        "otp_mode",
					Description: `Two step authentication mode.`,
				},
				resource.Attribute{
					Name:        "peering_ha_cloud_instance_id",
					Description: `Instance ID of the peering HA gateway.`,
				},
				resource.Attribute{
					Name:        "peering_ha_gw_name",
					Description: `Aviatrix gateway unique name of HA gateway.`,
				},
				resource.Attribute{
					Name:        "peering_ha_gw_size",
					Description: `Peering HA Gateway Size.`,
				},
				resource.Attribute{
					Name:        "peering_ha_insane_mode_az",
					Description: `AZ of subnet being created for Insane Mode Peering HA Gateway. Required if insane_mode is set.`,
				},
				resource.Attribute{
					Name:        "peering_ha_private_ip",
					Description: `Private IP address of HA gateway.`,
				},
				resource.Attribute{
					Name:        "peering_ha_public_ip",
					Description: `Public IP address that you want assigned to the HA peering instance.`,
				},
				resource.Attribute{
					Name:        "peering_ha_subnet",
					Description: `Public Subnet Information while creating Peering HA Gateway, only subnet is accepted. Required to create peering ha gateway if cloud_type = 1 or 8 (AWS or Azure).`,
				},
				resource.Attribute{
					Name:        "peering_ha_zone",
					Description: `Zone information for creating Peering HA Gateway. Required to create peering ha gateway if cloud_type = 4 (GCP).`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Private IP address of the Gateway created.`,
				},
				resource.Attribute{
					Name:        "public_dns_server",
					Description: `NS server used by the gateway.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP address of the Gateway created.`,
				},
				resource.Attribute{
					Name:        "saml_enabled",
					Description: `Status of SAML.`,
				},
				resource.Attribute{
					Name:        "search_domains",
					Description: `A list of domain names that will use the NameServer when a specific name is not in the destination when Split Tunnel Mode is enabled.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `Security group used for the gateway.`,
				},
				resource.Attribute{
					Name:        "single_az_ha",
					Description: `Status of Single AZ HA.`,
				},
				resource.Attribute{
					Name:        "single_ip_snat",
					Description: `Single IP Source NAT status for the container.`,
				},
				resource.Attribute{
					Name:        "split_tunnel",
					Description: `Status of split tunnel mode.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `A VPC Network address range selected from one of the available network ranges.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC-ID/VNet-Name of cloud provider.`,
				},
				resource.Attribute{
					Name:        "vpc_reg",
					Description: `Region of cloud provider.`,
				},
				resource.Attribute{
					Name:        "vpn_access",
					Description: `Status of user access through VPN to the container.`,
				},
				resource.Attribute{
					Name:        "vpn_cidr",
					Description: `VPN CIDR block for the container.`,
				},
				resource.Attribute{
					Name:        "vpn_protocol",
					Description: `ELB protocol for VPN gateway with ELB enabled.`,
				},
				resource.Attribute{
					Name:        "availability_domain",
					Description: `Availability domain for OCI.`,
				},
				resource.Attribute{
					Name:        "fault_domain",
					Description: `Fault domain for OCI.`,
				},
				resource.Attribute{
					Name:        "peering_ha_availability_domain",
					Description: `HA gateway availability domain for OCI.`,
				},
				resource.Attribute{
					Name:        "peering_ha_fault_domain",
					Description: `HA gateway fault domain for OCI.`,
				},
				resource.Attribute{
					Name:        "software_version",
					Description: `The software version of the gateway.`,
				},
				resource.Attribute{
					Name:        "image_version",
					Description: `The image version of the gateway.`,
				},
				resource.Attribute{
					Name:        "peering_ha_software_version",
					Description: `The software version of the HA gateway.`,
				},
				resource.Attribute{
					Name:        "peering_ha_image_version",
					Description: `The image version of the HA gateway. The following argument is deprecated:`,
				},
				resource.Attribute{
					Name:        "tag_list",
					Description: `Instance tag of cloud provider.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_gateway_image",
			Category:         "Gateway",
			ShortDescription: `Gets an Aviatrix gateway image version details.`,
			Description:      ``,
			Keywords: []string{
				"gateway",
				"image",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_type",
					Description: `(Required) Cloud type. Type: Integer. Example: 1 (AWS)`,
				},
				resource.Attribute{
					Name:        "software_version",
					Description: `(Required) Software version. Type: String. Example: "6.4.2487" ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "image_version",
					Description: `Image version that is compatible with the given cloud_type and software_version.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "image_version",
					Description: `Image version that is compatible with the given cloud_type and software_version.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_network_domains",
			Category:         "TGW Orchestrator",
			ShortDescription: `Gets a list of all Network Domains.`,
			Description:      ``,
			Keywords: []string{
				"tgw",
				"orchestrator",
				"network",
				"domains",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "network_domains",
					Description: `The list of all Network Domains`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Network Domain name.`,
				},
				resource.Attribute{
					Name:        "tgw_name",
					Description: `AWS TGW name.`,
				},
				resource.Attribute{
					Name:        "account",
					Description: `Access Account name.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `Route table's id.`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `Type of cloud service provider.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Region of cloud provider.`,
				},
				resource.Attribute{
					Name:        "intra_domain_inspection",
					Description: `Firewall inspection for traffic within one Security Domain.`,
				},
				resource.Attribute{
					Name:        "egress_inspection",
					Description: `Egress inspection is enable or not.`,
				},
				resource.Attribute{
					Name:        "inspection_policy",
					Description: `Inspection policy name.`,
				},
				resource.Attribute{
					Name:        "intra_domain_inspection_name",
					Description: `Intra domain inspection name.`,
				},
				resource.Attribute{
					Name:        "egress_inspection_name",
					Description: `Egress inspection name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of network domain.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_spoke_gateway",
			Category:         "Multi-Cloud Transit",
			ShortDescription: `Gets an Aviatrix spoke gateway's details.`,
			Description:      ``,
			Keywords: []string{
				"multi",
				"cloud",
				"transit",
				"spoke",
				"gateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "gw_name",
					Description: `(Required) Spoke gateway name. It can be used for getting spoke gateway. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `Aviatrix account name.`,
				},
				resource.Attribute{
					Name:        "allocate_new_eip",
					Description: `When value is false, an idle address in Elastic IP pool is reused for this gateway. Otherwise, a new Elastic IP is allocated and used for this gateway.`,
				},
				resource.Attribute{
					Name:        "cloud_instance_id",
					Description: `Cloud instance ID.`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `Type of cloud service provider.`,
				},
				resource.Attribute{
					Name:        "customized_spoke_vpc_routes",
					Description: `A list of comma separated CIDRs to be customized for the spoke VPC routes.`,
				},
				resource.Attribute{
					Name:        "enable_encrypt_volume",
					Description: `Status of Encrypt Volume of spoke gateway.`,
				},
				resource.Attribute{
					Name:        "enable_vpc_dns_server",
					Description: `Status of VPC Dns Server of spoke gateway.`,
				},
				resource.Attribute{
					Name:        "filtered_spoke_vpc_routes",
					Description: `A list of comma separated CIDRs to be filtered from the spoke VPC route table.`,
				},
				resource.Attribute{
					Name:        "ha_cloud_instance_id",
					Description: `Cloud instance ID of HA spoke gateway.`,
				},
				resource.Attribute{
					Name:        "ha_insane_mode_az",
					Description: `AZ of subnet being created for Insane Mode Spoke HA Gateway.`,
				},
				resource.Attribute{
					Name:        "ha_gw_name",
					Description: `Aviatrix spoke gateway unique name of HA spoke gateway.`,
				},
				resource.Attribute{
					Name:        "ha_gw_size",
					Description: `HA Gateway Size.`,
				},
				resource.Attribute{
					Name:        "ha_private_ip",
					Description: `Private IP address of HA spoke gateway.`,
				},
				resource.Attribute{
					Name:        "ha_public_ip",
					Description: `Public IP address of the HA spoke gateway.`,
				},
				resource.Attribute{
					Name:        "ha_subnet",
					Description: `HA Subnet.`,
				},
				resource.Attribute{
					Name:        "ha_zone",
					Description: `HA Zone.`,
				},
				resource.Attribute{
					Name:        "gw_name",
					Description: `Aviatrix spoke gateway name.`,
				},
				resource.Attribute{
					Name:        "gw_size",
					Description: `Size of spoke gateway instance.`,
				},
				resource.Attribute{
					Name:        "included_advertised_spoke_routes",
					Description: `A list of comma separated CIDRs to be advertised to on-prem as "Included CIDR List".`,
				},
				resource.Attribute{
					Name:        "insane_mode",
					Description: `Status of Insane Mode for Spoke Gateway.`,
				},
				resource.Attribute{
					Name:        "insane_mode_az",
					Description: `AZ of subnet being created for Insane Mode spoke gateway.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Private IP address of the spoke gateway.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP of spoke gateway.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `Security group used of the spoke gateway.`,
				},
				resource.Attribute{
					Name:        "single_az_ha",
					Description: `Status of Single AZ HA of spoke gateway.`,
				},
				resource.Attribute{
					Name:        "single_ip_snat",
					Description: `Status of Single IP Source NAT mode of the spoke gateway.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `A VPC Network address range selected from one of the available network ranges.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC-ID/VNet-Name of cloud provider.`,
				},
				resource.Attribute{
					Name:        "vpc_reg",
					Description: `Region of cloud provider.`,
				},
				resource.Attribute{
					Name:        "enable_private_oob",
					Description: `Status of private OOB for the spoke gateway.`,
				},
				resource.Attribute{
					Name:        "oob_management_subnet",
					Description: `OOB management subnet.`,
				},
				resource.Attribute{
					Name:        "oob_availability_zone",
					Description: `OOB availability zone.`,
				},
				resource.Attribute{
					Name:        "ha_oob_management_subnet",
					Description: `HA OOB management subnet.`,
				},
				resource.Attribute{
					Name:        "ha_oob_availability_zone",
					Description: `HA OOB availability zone.`,
				},
				resource.Attribute{
					Name:        "availability_domain",
					Description: `Availability domain for OCI.`,
				},
				resource.Attribute{
					Name:        "fault_domain",
					Description: `Fault domain for OCI.`,
				},
				resource.Attribute{
					Name:        "ha_availability_domain",
					Description: `HA gateway availability domain for OCI.`,
				},
				resource.Attribute{
					Name:        "ha_fault_domain",
					Description: `HA gateway fault domain for OCI.`,
				},
				resource.Attribute{
					Name:        "software_version",
					Description: `The software version of the gateway.`,
				},
				resource.Attribute{
					Name:        "image_version",
					Description: `The image version of the gateway.`,
				},
				resource.Attribute{
					Name:        "ha_software_version",
					Description: `The software version of the HA gateway.`,
				},
				resource.Attribute{
					Name:        "ha_image_version",
					Description: `The image version of the HA gateway.`,
				},
				resource.Attribute{
					Name:        "eip",
					Description: `The EIP address of the Spoke Gateway.`,
				},
				resource.Attribute{
					Name:        "ha_eip",
					Description: `The EIP address of the HA Spoke Gateway. The following arguments are deprecated:`,
				},
				resource.Attribute{
					Name:        "tag_list",
					Description: `Instance tag of cloud provider.`,
				},
				resource.Attribute{
					Name:        "transit_gw",
					Description: `Transit gateways attached to this spoke gateway.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "account_name",
					Description: `Aviatrix account name.`,
				},
				resource.Attribute{
					Name:        "allocate_new_eip",
					Description: `When value is false, an idle address in Elastic IP pool is reused for this gateway. Otherwise, a new Elastic IP is allocated and used for this gateway.`,
				},
				resource.Attribute{
					Name:        "cloud_instance_id",
					Description: `Cloud instance ID.`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `Type of cloud service provider.`,
				},
				resource.Attribute{
					Name:        "customized_spoke_vpc_routes",
					Description: `A list of comma separated CIDRs to be customized for the spoke VPC routes.`,
				},
				resource.Attribute{
					Name:        "enable_encrypt_volume",
					Description: `Status of Encrypt Volume of spoke gateway.`,
				},
				resource.Attribute{
					Name:        "enable_vpc_dns_server",
					Description: `Status of VPC Dns Server of spoke gateway.`,
				},
				resource.Attribute{
					Name:        "filtered_spoke_vpc_routes",
					Description: `A list of comma separated CIDRs to be filtered from the spoke VPC route table.`,
				},
				resource.Attribute{
					Name:        "ha_cloud_instance_id",
					Description: `Cloud instance ID of HA spoke gateway.`,
				},
				resource.Attribute{
					Name:        "ha_insane_mode_az",
					Description: `AZ of subnet being created for Insane Mode Spoke HA Gateway.`,
				},
				resource.Attribute{
					Name:        "ha_gw_name",
					Description: `Aviatrix spoke gateway unique name of HA spoke gateway.`,
				},
				resource.Attribute{
					Name:        "ha_gw_size",
					Description: `HA Gateway Size.`,
				},
				resource.Attribute{
					Name:        "ha_private_ip",
					Description: `Private IP address of HA spoke gateway.`,
				},
				resource.Attribute{
					Name:        "ha_public_ip",
					Description: `Public IP address of the HA spoke gateway.`,
				},
				resource.Attribute{
					Name:        "ha_subnet",
					Description: `HA Subnet.`,
				},
				resource.Attribute{
					Name:        "ha_zone",
					Description: `HA Zone.`,
				},
				resource.Attribute{
					Name:        "gw_name",
					Description: `Aviatrix spoke gateway name.`,
				},
				resource.Attribute{
					Name:        "gw_size",
					Description: `Size of spoke gateway instance.`,
				},
				resource.Attribute{
					Name:        "included_advertised_spoke_routes",
					Description: `A list of comma separated CIDRs to be advertised to on-prem as "Included CIDR List".`,
				},
				resource.Attribute{
					Name:        "insane_mode",
					Description: `Status of Insane Mode for Spoke Gateway.`,
				},
				resource.Attribute{
					Name:        "insane_mode_az",
					Description: `AZ of subnet being created for Insane Mode spoke gateway.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Private IP address of the spoke gateway.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP of spoke gateway.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `Security group used of the spoke gateway.`,
				},
				resource.Attribute{
					Name:        "single_az_ha",
					Description: `Status of Single AZ HA of spoke gateway.`,
				},
				resource.Attribute{
					Name:        "single_ip_snat",
					Description: `Status of Single IP Source NAT mode of the spoke gateway.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `A VPC Network address range selected from one of the available network ranges.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC-ID/VNet-Name of cloud provider.`,
				},
				resource.Attribute{
					Name:        "vpc_reg",
					Description: `Region of cloud provider.`,
				},
				resource.Attribute{
					Name:        "enable_private_oob",
					Description: `Status of private OOB for the spoke gateway.`,
				},
				resource.Attribute{
					Name:        "oob_management_subnet",
					Description: `OOB management subnet.`,
				},
				resource.Attribute{
					Name:        "oob_availability_zone",
					Description: `OOB availability zone.`,
				},
				resource.Attribute{
					Name:        "ha_oob_management_subnet",
					Description: `HA OOB management subnet.`,
				},
				resource.Attribute{
					Name:        "ha_oob_availability_zone",
					Description: `HA OOB availability zone.`,
				},
				resource.Attribute{
					Name:        "availability_domain",
					Description: `Availability domain for OCI.`,
				},
				resource.Attribute{
					Name:        "fault_domain",
					Description: `Fault domain for OCI.`,
				},
				resource.Attribute{
					Name:        "ha_availability_domain",
					Description: `HA gateway availability domain for OCI.`,
				},
				resource.Attribute{
					Name:        "ha_fault_domain",
					Description: `HA gateway fault domain for OCI.`,
				},
				resource.Attribute{
					Name:        "software_version",
					Description: `The software version of the gateway.`,
				},
				resource.Attribute{
					Name:        "image_version",
					Description: `The image version of the gateway.`,
				},
				resource.Attribute{
					Name:        "ha_software_version",
					Description: `The software version of the HA gateway.`,
				},
				resource.Attribute{
					Name:        "ha_image_version",
					Description: `The image version of the HA gateway.`,
				},
				resource.Attribute{
					Name:        "eip",
					Description: `The EIP address of the Spoke Gateway.`,
				},
				resource.Attribute{
					Name:        "ha_eip",
					Description: `The EIP address of the HA Spoke Gateway. The following arguments are deprecated:`,
				},
				resource.Attribute{
					Name:        "tag_list",
					Description: `Instance tag of cloud provider.`,
				},
				resource.Attribute{
					Name:        "transit_gw",
					Description: `Transit gateways attached to this spoke gateway.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_spoke_gateway_inspection_subnets",
			Category:         "Multi-Cloud Transit",
			ShortDescription: `Gets all subnets available for the subnet inspection feature.`,
			Description:      ``,
			Keywords: []string{
				"multi",
				"cloud",
				"transit",
				"spoke",
				"gateway",
				"inspection",
				"subnets",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "gw_name",
					Description: `(Required) Spoke gateway name. ## Attribute Reference In addition to the argument above, the following attribute is exported:`,
				},
				resource.Attribute{
					Name:        "subnets_for_inspection",
					Description: `The list of all subnets available for the subnet inspection feature. This attribute is only supported for Azure.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "subnets_for_inspection",
					Description: `The list of all subnets available for the subnet inspection feature. This attribute is only supported for Azure.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_spoke_gateways",
			Category:         "Multi-Cloud Transit",
			ShortDescription: `Gets a list of all Aviatrix spoke gateway's details.`,
			Description:      ``,
			Keywords: []string{
				"multi",
				"cloud",
				"transit",
				"spoke",
				"gateways",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "gateway_list",
					Description: `The list of all spoke gateways`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `Type of cloud service provider.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `Aviatrix account name.`,
				},
				resource.Attribute{
					Name:        "gw_name",
					Description: `Aviatrix spoke gateway name.`,
				},
				resource.Attribute{
					Name:        "gw_size",
					Description: `Size of spoke gateway instance.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC-ID/VNet-Name of cloud provider.`,
				},
				resource.Attribute{
					Name:        "vpc_reg",
					Description: `Region of cloud provider.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `Availability Zone. Only available for cloud_type = 8 (Azure). Must be in the form 'az-n', for example, 'az-2'.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `A VPC Network address range selected from one of the available network ranges.`,
				},
				resource.Attribute{
					Name:        "insane_mode_az",
					Description: `AZ of subnet being created for Insane Mode spoke gateway.`,
				},
				resource.Attribute{
					Name:        "single_ip_snat",
					Description: `Status of Single IP Source Nat mode of the spoke gateway.`,
				},
				resource.Attribute{
					Name:        "allocate_new_eip",
					Description: `When value is false, an idle address in Elastic IP pool is reused for this gateway. Otherwise, a new Elastic IP is allocated and used for this gateway.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP address of the Spoke Gateway.`,
				},
				resource.Attribute{
					Name:        "single_az_ha",
					Description: `Status of Single AZ HA of spoke gateway.`,
				},
				resource.Attribute{
					Name:        "transit_gw",
					Description: `Transit Gateways this spoke has joined.`,
				},
				resource.Attribute{
					Name:        "insane_mode",
					Description: `Status of Insane Mode of the spoke gateway.`,
				},
				resource.Attribute{
					Name:        "enable_vpc_dns_server",
					Description: `Status of Vpc Dns Server of the spoke Gateway.`,
				},
				resource.Attribute{
					Name:        "enable_encrypt_volume",
					Description: `Status of Encrypt Gateway EBS Volume of the spoke gateway.`,
				},
				resource.Attribute{
					Name:        "customized_spoke_vpc_routes",
					Description: `A list of comma separated CIDRs to be customized for the spoke VPC routes.`,
				},
				resource.Attribute{
					Name:        "filtered_spoke_vpc_routes",
					Description: `A list of comma separated CIDRs to be filtered from the spoke VPC route table.`,
				},
				resource.Attribute{
					Name:        "included_advertised_spoke_routes",
					Description: `A list of comma separated CIDRs to be advertised to on-prem as 'Included CIDR List'. When configured, it will replace all advertised routes from this VPC.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `Security group used for the spoke gateway.`,
				},
				resource.Attribute{
					Name:        "cloud_instance_id",
					Description: `Instance ID of the spoke gateway.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Private IP address of the spoke gateway created.`,
				},
				resource.Attribute{
					Name:        "enable_private_oob",
					Description: `Status of private OOB for the spoke gateway.`,
				},
				resource.Attribute{
					Name:        "oob_management_subnet",
					Description: `OOB management subnet.`,
				},
				resource.Attribute{
					Name:        "oob_availability_zone",
					Description: `OOB availability zone.`,
				},
				resource.Attribute{
					Name:        "tunnel_detection_time",
					Description: `The IPSec tunnel down detection time for the spoke gateway.`,
				},
				resource.Attribute{
					Name:        "availability_domain",
					Description: `Availability domain for OCI.`,
				},
				resource.Attribute{
					Name:        "fault_domain",
					Description: `Fault domain for OCI.`,
				},
				resource.Attribute{
					Name:        "software_version",
					Description: `The software version of the gateway.`,
				},
				resource.Attribute{
					Name:        "image_version",
					Description: `The image version of the gateway.`,
				},
				resource.Attribute{
					Name:        "enable_monitor_gateway_subnets",
					Description: `Enable monitor gateway subnets. Only valid for cloud_type = 1 (AWS) or 256 (AWSGov).`,
				},
				resource.Attribute{
					Name:        "monitor_exclude_list",
					Description: `A set of monitored instance ids. Only valid when 'enable_monitor_gateway_subnets' = true.`,
				},
				resource.Attribute{
					Name:        "enable_jumbo_frame",
					Description: `Enable jumbo frame support for spoke gateway.`,
				},
				resource.Attribute{
					Name:        "enable_private_vpc_default_route",
					Description: `Enable Private VPC Default Route.`,
				},
				resource.Attribute{
					Name:        "enable_skip_public_route_table_update",
					Description: `Skip Public Route Table Update.`,
				},
				resource.Attribute{
					Name:        "enable_auto_advertise_s2c_cidrs",
					Description: `Automatically advertise remote CIDR to Aviatrix Spoke Gateway when route based Site2Cloud Tunnel is created.`,
				},
				resource.Attribute{
					Name:        "spoke_bgp_manual_advertise_cidrs",
					Description: `Intended CIDR list to be advertised to external BGP router.`,
				},
				resource.Attribute{
					Name:        "enable_bgp",
					Description: `Enable BGP.`,
				},
				resource.Attribute{
					Name:        "enable_learned_cidrs_approval",
					Description: `Switch to enable/disable encrypted transit approval for BGP Spoke Gateway.`,
				},
				resource.Attribute{
					Name:        "learned_cidrs_approval_mode",
					Description: `Set the learned CIDRs approval mode.`,
				},
				resource.Attribute{
					Name:        "bgp_ecmp",
					Description: `Enable Equal Cost Multi Path (ECMP) routing for the next hop.`,
				},
				resource.Attribute{
					Name:        "enable_active_standby",
					Description: `Enables Active-Standby Mode, available only with HA enabled.`,
				},
				resource.Attribute{
					Name:        "enable_active_standby_preemptive",
					Description: `Enables Preemptive Mode for Active-Standby, available only with Active-Standby enabled.`,
				},
				resource.Attribute{
					Name:        "disable_route_propagation",
					Description: `Disables route propagation on BGP Spoke to attached Transit Gateway.`,
				},
				resource.Attribute{
					Name:        "local_as_number",
					Description: `Changes the Aviatrix Spoke Gateway ASN number before you setup Aviatrix Spoke Gateway connection configurations.`,
				},
				resource.Attribute{
					Name:        "prepend_as_path",
					Description: `List of AS numbers to populate BGP AP_PATH field when it advertises to VGW or peer devices.`,
				},
				resource.Attribute{
					Name:        "bgp_polling_time",
					Description: `BGP route polling time. Unit is in seconds.`,
				},
				resource.Attribute{
					Name:        "bgp_hold_time",
					Description: `BGP Hold Time.`,
				},
				resource.Attribute{
					Name:        "enable_spot_instance",
					Description: `Enable spot instance. NOT supported for production deployment.`,
				},
				resource.Attribute{
					Name:        "spot_price",
					Description: `Price for spot instance. NOT supported for production deployment.`,
				},
				resource.Attribute{
					Name:        "azure_eip_name_resource_group",
					Description: `The name of the public IP address and its resource group in Azure to assign to this Spoke Gateway.`,
				},
				resource.Attribute{
					Name:        "eip",
					Description: `The EIP address of the Spoke Gateway.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_transit_gateway",
			Category:         "Multi-Cloud Transit",
			ShortDescription: `Gets an Aviatrix transit gateway's details.`,
			Description:      ``,
			Keywords: []string{
				"multi",
				"cloud",
				"transit",
				"gateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "gw_name",
					Description: `(Required) Transit gateway name. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `Aviatrix account name.`,
				},
				resource.Attribute{
					Name:        "allocate_new_eip",
					Description: `When value is false, an idle address in Elastic IP pool is reused for this gateway. Otherwise, a new Elastic IP is allocated and used for this gateway.`,
				},
				resource.Attribute{
					Name:        "bgp_manual_spoke_advertise_cidrs",
					Description: `Intended CIDR list to advertise to VGW.`,
				},
				resource.Attribute{
					Name:        "cloud_instance_id",
					Description: `Instance ID of the transit gateway.`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `Type of cloud service provider.`,
				},
				resource.Attribute{
					Name:        "connected_transit",
					Description: `Status of Connected Transit of transit gateway.`,
				},
				resource.Attribute{
					Name:        "customized_spoke_vpc_routes",
					Description: `A list of comma separated CIDRs to be customized for the spoke VPC routes.`,
				},
				resource.Attribute{
					Name:        "customized_transit_vpc_routes",
					Description: `A list of CIDRs to be customized for the transit VPC routes.`,
				},
				resource.Attribute{
					Name:        "gw_size",
					Description: `Size of transit gateway instance.`,
				},
				resource.Attribute{
					Name:        "gw_name",
					Description: `Aviatrix transit gateway name.`,
				},
				resource.Attribute{
					Name:        "insane_mode_az",
					Description: `AZ of subnet being created for Insane Mode transit gateway.`,
				},
				resource.Attribute{
					Name:        "enable_advertise_transit_cidr",
					Description: `Status of Advertise Transit VPC network CIDR of the transit gateway.`,
				},
				resource.Attribute{
					Name:        "enable_encrypt_volume",
					Description: `Status of Encrypt Gateway EBS Volume of the transit gateway.`,
				},
				resource.Attribute{
					Name:        "enable_firenet",
					Description: `Status of FireNet Interfaces of the transit gateway.`,
				},
				resource.Attribute{
					Name:        "enable_hybrid_connection",
					Description: `Sign of readiness for TGW connection.`,
				},
				resource.Attribute{
					Name:        "enable_learned_cidrs_approval",
					Description: `Status of Encrypted Transit Approval for transit gateway.`,
				},
				resource.Attribute{
					Name:        "enable_vpc_dns_server",
					Description: `Status of Vpc Dns Server of the transit Gateway.`,
				},
				resource.Attribute{
					Name:        "enable_transit_firenet",
					Description: `Status of Transit FireNet Interfaces of the transit gateway.`,
				},
				resource.Attribute{
					Name:        "enable_egress_transit_firenet",
					Description: `Status of Egress Transit FireNet being enabled on the gateway.`,
				},
				resource.Attribute{
					Name:        "excluded_advertised_spoke_routes",
					Description: `A list of comma separated CIDRs to be advertised to on-prem as "Excluded CIDR List".`,
				},
				resource.Attribute{
					Name:        "filtered_spoke_vpc_routes",
					Description: `A list of comma separated CIDRs to be filtered from the spoke VPC route table.`,
				},
				resource.Attribute{
					Name:        "ha_insane_mode_az",
					Description: `AZ of subnet being created for Insane Mode Transit HA Gateway.`,
				},
				resource.Attribute{
					Name:        "ha_cloud_instance_id",
					Description: `Cloud instance ID of HA transit gateway.`,
				},
				resource.Attribute{
					Name:        "ha_gw_name",
					Description: `Aviatrix transit gateway unique name of HA transit gateway.`,
				},
				resource.Attribute{
					Name:        "ha_gw_size",
					Description: `HA Gateway Size.`,
				},
				resource.Attribute{
					Name:        "ha_private_ip",
					Description: `Private IP address that assigned to the HA Transit Gateway.`,
				},
				resource.Attribute{
					Name:        "ha_public_ip",
					Description: `Public IP address that assigned to the HA Transit Gateway.`,
				},
				resource.Attribute{
					Name:        "ha_subnet",
					Description: `HA Subnet.`,
				},
				resource.Attribute{
					Name:        "ha_zone",
					Description: `HA Zone.`,
				},
				resource.Attribute{
					Name:        "insane_mode",
					Description: `Status of Insane Mode of the transit gateway.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Private IP address of the transit gateway created.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP address of the Transit Gateway created.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `Security group used for the transit gateway.`,
				},
				resource.Attribute{
					Name:        "single_az_ha",
					Description: `Status of Single AZ HA of transit gateway.`,
				},
				resource.Attribute{
					Name:        "single_ip_snat",
					Description: `Status of Single IP Source Nat mode of the transit gateway.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `A VPC Network address range selected from one of the available network ranges.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC-ID/VNet-Name of cloud provider.`,
				},
				resource.Attribute{
					Name:        "vpc_reg",
					Description: `Region of cloud provider.`,
				},
				resource.Attribute{
					Name:        "enable_private_oob",
					Description: `Status of private OOB for the transit gateway.`,
				},
				resource.Attribute{
					Name:        "oob_management_subnet",
					Description: `OOB management subnet.`,
				},
				resource.Attribute{
					Name:        "oob_availability_zone",
					Description: `OOB availability zone.`,
				},
				resource.Attribute{
					Name:        "ha_oob_management_subnet",
					Description: `HA OOB management subnet.`,
				},
				resource.Attribute{
					Name:        "ha_oob_availability_zone",
					Description: `HA OOB availability zone.`,
				},
				resource.Attribute{
					Name:        "enable_multi_tier_transit",
					Description: `Status of multi-tier transit mode on transit gateway.`,
				},
				resource.Attribute{
					Name:        "availability_domain",
					Description: `Availability domain for OCI.`,
				},
				resource.Attribute{
					Name:        "fault_domain",
					Description: `Fault domain for OCI.`,
				},
				resource.Attribute{
					Name:        "ha_availability_domain",
					Description: `HA gateway availability domain for OCI.`,
				},
				resource.Attribute{
					Name:        "ha_fault_domain",
					Description: `HA gateway fault domain for OCI.`,
				},
				resource.Attribute{
					Name:        "software_version",
					Description: `The software version of the gateway.`,
				},
				resource.Attribute{
					Name:        "image_version",
					Description: `The image version of the gateway.`,
				},
				resource.Attribute{
					Name:        "ha_software_version",
					Description: `The software version of the HA gateway.`,
				},
				resource.Attribute{
					Name:        "ha_image_version",
					Description: `The image version of the HA gateway.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `Availability Zone for Azure.`,
				},
				resource.Attribute{
					Name:        "enable_gateway_load_balancer",
					Description: `Status of AWS Gateway Load Balancer.`,
				},
				resource.Attribute{
					Name:        "lan_vpc_id",
					Description: `LAN VPC ID for GCP Transit FireNet.`,
				},
				resource.Attribute{
					Name:        "lan_private_subnet",
					Description: `LAN Private Subnet for GCP Transit FireNet.`,
				},
				resource.Attribute{
					Name:        "learned_cidrs_approval_mode",
					Description: `Learned CIDRs approval mode.`,
				},
				resource.Attribute{
					Name:        "approved_learned_cidrs",
					Description: `Approved learned CIDRs.`,
				},
				resource.Attribute{
					Name:        "bgp_polling_time",
					Description: `BGP route polling time.`,
				},
				resource.Attribute{
					Name:        "prepend_as_path",
					Description: `List of AS numbers to populate BGP AP_PATH field when it advertises to VGW or peer devices.`,
				},
				resource.Attribute{
					Name:        "bgp_ecmp",
					Description: `Status of Equal Cost Multi Path (ECMP) routing for the next hop.`,
				},
				resource.Attribute{
					Name:        "enable_segmentation",
					Description: `Status of segmentation.`,
				},
				resource.Attribute{
					Name:        "enable_active_standby",
					Description: `Status of Active-Standby Mode.`,
				},
				resource.Attribute{
					Name:        "enable_active_standby_preemptive",
					Description: `Status of Preemptive Mode for Active-Standby.`,
				},
				resource.Attribute{
					Name:        "enable_monitor_gateway_subnets",
					Description: `Status of monitor gateway subnets.`,
				},
				resource.Attribute{
					Name:        "monitor_exclude_list",
					Description: `A set of monitored instance IDs.`,
				},
				resource.Attribute{
					Name:        "enable_bgp_over_lan",
					Description: `Status of BGP over LAN functionality.`,
				},
				resource.Attribute{
					Name:        "bgp_lan_interfaces",
					Description: `Interfaces to run BGP protocol on top of the ethernet interface, to connect to the onprem/remote peer.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC-ID of GCP cloud provider.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `Subnet Info.`,
				},
				resource.Attribute{
					Name:        "ha_bgp_lan_interfaces",
					Description: `Interfaces to run BGP protocol on top of the ethernet interface, to connect to the onprem/remote peer. Only available for GCP HA Transit.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC-ID of GCP cloud provider.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `Subnet Info.`,
				},
				resource.Attribute{
					Name:        "enable_jumbo_frame",
					Description: `Status of jumbo frame support.`,
				},
				resource.Attribute{
					Name:        "bgp_hold_time",
					Description: `BGP Hold Time.`,
				},
				resource.Attribute{
					Name:        "enable_transit_summarize_cidr_to_tgw",
					Description: `Status of transit summarize CIDR to TGW.`,
				},
				resource.Attribute{
					Name:        "enable_spot_instance",
					Description: `Status of spot instance.`,
				},
				resource.Attribute{
					Name:        "spot_price",
					Description: `Price for spot instance.`,
				},
				resource.Attribute{
					Name:        "azure_eip_name_resource_group",
					Description: `The name of the public IP address and its resource group in Azure to assign to this Transit Gateway.`,
				},
				resource.Attribute{
					Name:        "ha_azure_eip_name_resource_group",
					Description: `The name of the public IP address and its resource group in Azure to assign to the HA Transit Gateway.`,
				},
				resource.Attribute{
					Name:        "local_as_number",
					Description: `Local ASN number.`,
				},
				resource.Attribute{
					Name:        "ha_security_group_id",
					Description: `HA security group used for the transit gateway.`,
				},
				resource.Attribute{
					Name:        "lan_interface_cidr",
					Description: `Transit gateway lan interface cidr.`,
				},
				resource.Attribute{
					Name:        "ha_lan_interface_cidr",
					Description: `Transit gateway lan interface cidr for the HA gateway.`,
				},
				resource.Attribute{
					Name:        "bgp_lan_ip_list",
					Description: `List of available BGP LAN interface IPs for transit external device connection creation. Only supports GCP and Azure.`,
				},
				resource.Attribute{
					Name:        "ha_bgp_lan_ip_list",
					Description: `List of available BGP LAN interface IPs for transit external device HA connection creation. Only supports GCP and Azure.`,
				},
				resource.Attribute{
					Name:        "eip",
					Description: `The EIP address of the Transit Gateway.`,
				},
				resource.Attribute{
					Name:        "ha_eip",
					Description: `The EIP address of the HA Transit Gateway. The following argument is deprecated:`,
				},
				resource.Attribute{
					Name:        "tag_list",
					Description: `Instance tag of cloud provider.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "account_name",
					Description: `Aviatrix account name.`,
				},
				resource.Attribute{
					Name:        "allocate_new_eip",
					Description: `When value is false, an idle address in Elastic IP pool is reused for this gateway. Otherwise, a new Elastic IP is allocated and used for this gateway.`,
				},
				resource.Attribute{
					Name:        "bgp_manual_spoke_advertise_cidrs",
					Description: `Intended CIDR list to advertise to VGW.`,
				},
				resource.Attribute{
					Name:        "cloud_instance_id",
					Description: `Instance ID of the transit gateway.`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `Type of cloud service provider.`,
				},
				resource.Attribute{
					Name:        "connected_transit",
					Description: `Status of Connected Transit of transit gateway.`,
				},
				resource.Attribute{
					Name:        "customized_spoke_vpc_routes",
					Description: `A list of comma separated CIDRs to be customized for the spoke VPC routes.`,
				},
				resource.Attribute{
					Name:        "customized_transit_vpc_routes",
					Description: `A list of CIDRs to be customized for the transit VPC routes.`,
				},
				resource.Attribute{
					Name:        "gw_size",
					Description: `Size of transit gateway instance.`,
				},
				resource.Attribute{
					Name:        "gw_name",
					Description: `Aviatrix transit gateway name.`,
				},
				resource.Attribute{
					Name:        "insane_mode_az",
					Description: `AZ of subnet being created for Insane Mode transit gateway.`,
				},
				resource.Attribute{
					Name:        "enable_advertise_transit_cidr",
					Description: `Status of Advertise Transit VPC network CIDR of the transit gateway.`,
				},
				resource.Attribute{
					Name:        "enable_encrypt_volume",
					Description: `Status of Encrypt Gateway EBS Volume of the transit gateway.`,
				},
				resource.Attribute{
					Name:        "enable_firenet",
					Description: `Status of FireNet Interfaces of the transit gateway.`,
				},
				resource.Attribute{
					Name:        "enable_hybrid_connection",
					Description: `Sign of readiness for TGW connection.`,
				},
				resource.Attribute{
					Name:        "enable_learned_cidrs_approval",
					Description: `Status of Encrypted Transit Approval for transit gateway.`,
				},
				resource.Attribute{
					Name:        "enable_vpc_dns_server",
					Description: `Status of Vpc Dns Server of the transit Gateway.`,
				},
				resource.Attribute{
					Name:        "enable_transit_firenet",
					Description: `Status of Transit FireNet Interfaces of the transit gateway.`,
				},
				resource.Attribute{
					Name:        "enable_egress_transit_firenet",
					Description: `Status of Egress Transit FireNet being enabled on the gateway.`,
				},
				resource.Attribute{
					Name:        "excluded_advertised_spoke_routes",
					Description: `A list of comma separated CIDRs to be advertised to on-prem as "Excluded CIDR List".`,
				},
				resource.Attribute{
					Name:        "filtered_spoke_vpc_routes",
					Description: `A list of comma separated CIDRs to be filtered from the spoke VPC route table.`,
				},
				resource.Attribute{
					Name:        "ha_insane_mode_az",
					Description: `AZ of subnet being created for Insane Mode Transit HA Gateway.`,
				},
				resource.Attribute{
					Name:        "ha_cloud_instance_id",
					Description: `Cloud instance ID of HA transit gateway.`,
				},
				resource.Attribute{
					Name:        "ha_gw_name",
					Description: `Aviatrix transit gateway unique name of HA transit gateway.`,
				},
				resource.Attribute{
					Name:        "ha_gw_size",
					Description: `HA Gateway Size.`,
				},
				resource.Attribute{
					Name:        "ha_private_ip",
					Description: `Private IP address that assigned to the HA Transit Gateway.`,
				},
				resource.Attribute{
					Name:        "ha_public_ip",
					Description: `Public IP address that assigned to the HA Transit Gateway.`,
				},
				resource.Attribute{
					Name:        "ha_subnet",
					Description: `HA Subnet.`,
				},
				resource.Attribute{
					Name:        "ha_zone",
					Description: `HA Zone.`,
				},
				resource.Attribute{
					Name:        "insane_mode",
					Description: `Status of Insane Mode of the transit gateway.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Private IP address of the transit gateway created.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP address of the Transit Gateway created.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `Security group used for the transit gateway.`,
				},
				resource.Attribute{
					Name:        "single_az_ha",
					Description: `Status of Single AZ HA of transit gateway.`,
				},
				resource.Attribute{
					Name:        "single_ip_snat",
					Description: `Status of Single IP Source Nat mode of the transit gateway.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `A VPC Network address range selected from one of the available network ranges.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC-ID/VNet-Name of cloud provider.`,
				},
				resource.Attribute{
					Name:        "vpc_reg",
					Description: `Region of cloud provider.`,
				},
				resource.Attribute{
					Name:        "enable_private_oob",
					Description: `Status of private OOB for the transit gateway.`,
				},
				resource.Attribute{
					Name:        "oob_management_subnet",
					Description: `OOB management subnet.`,
				},
				resource.Attribute{
					Name:        "oob_availability_zone",
					Description: `OOB availability zone.`,
				},
				resource.Attribute{
					Name:        "ha_oob_management_subnet",
					Description: `HA OOB management subnet.`,
				},
				resource.Attribute{
					Name:        "ha_oob_availability_zone",
					Description: `HA OOB availability zone.`,
				},
				resource.Attribute{
					Name:        "enable_multi_tier_transit",
					Description: `Status of multi-tier transit mode on transit gateway.`,
				},
				resource.Attribute{
					Name:        "availability_domain",
					Description: `Availability domain for OCI.`,
				},
				resource.Attribute{
					Name:        "fault_domain",
					Description: `Fault domain for OCI.`,
				},
				resource.Attribute{
					Name:        "ha_availability_domain",
					Description: `HA gateway availability domain for OCI.`,
				},
				resource.Attribute{
					Name:        "ha_fault_domain",
					Description: `HA gateway fault domain for OCI.`,
				},
				resource.Attribute{
					Name:        "software_version",
					Description: `The software version of the gateway.`,
				},
				resource.Attribute{
					Name:        "image_version",
					Description: `The image version of the gateway.`,
				},
				resource.Attribute{
					Name:        "ha_software_version",
					Description: `The software version of the HA gateway.`,
				},
				resource.Attribute{
					Name:        "ha_image_version",
					Description: `The image version of the HA gateway.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `Availability Zone for Azure.`,
				},
				resource.Attribute{
					Name:        "enable_gateway_load_balancer",
					Description: `Status of AWS Gateway Load Balancer.`,
				},
				resource.Attribute{
					Name:        "lan_vpc_id",
					Description: `LAN VPC ID for GCP Transit FireNet.`,
				},
				resource.Attribute{
					Name:        "lan_private_subnet",
					Description: `LAN Private Subnet for GCP Transit FireNet.`,
				},
				resource.Attribute{
					Name:        "learned_cidrs_approval_mode",
					Description: `Learned CIDRs approval mode.`,
				},
				resource.Attribute{
					Name:        "approved_learned_cidrs",
					Description: `Approved learned CIDRs.`,
				},
				resource.Attribute{
					Name:        "bgp_polling_time",
					Description: `BGP route polling time.`,
				},
				resource.Attribute{
					Name:        "prepend_as_path",
					Description: `List of AS numbers to populate BGP AP_PATH field when it advertises to VGW or peer devices.`,
				},
				resource.Attribute{
					Name:        "bgp_ecmp",
					Description: `Status of Equal Cost Multi Path (ECMP) routing for the next hop.`,
				},
				resource.Attribute{
					Name:        "enable_segmentation",
					Description: `Status of segmentation.`,
				},
				resource.Attribute{
					Name:        "enable_active_standby",
					Description: `Status of Active-Standby Mode.`,
				},
				resource.Attribute{
					Name:        "enable_active_standby_preemptive",
					Description: `Status of Preemptive Mode for Active-Standby.`,
				},
				resource.Attribute{
					Name:        "enable_monitor_gateway_subnets",
					Description: `Status of monitor gateway subnets.`,
				},
				resource.Attribute{
					Name:        "monitor_exclude_list",
					Description: `A set of monitored instance IDs.`,
				},
				resource.Attribute{
					Name:        "enable_bgp_over_lan",
					Description: `Status of BGP over LAN functionality.`,
				},
				resource.Attribute{
					Name:        "bgp_lan_interfaces",
					Description: `Interfaces to run BGP protocol on top of the ethernet interface, to connect to the onprem/remote peer.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC-ID of GCP cloud provider.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `Subnet Info.`,
				},
				resource.Attribute{
					Name:        "ha_bgp_lan_interfaces",
					Description: `Interfaces to run BGP protocol on top of the ethernet interface, to connect to the onprem/remote peer. Only available for GCP HA Transit.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC-ID of GCP cloud provider.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `Subnet Info.`,
				},
				resource.Attribute{
					Name:        "enable_jumbo_frame",
					Description: `Status of jumbo frame support.`,
				},
				resource.Attribute{
					Name:        "bgp_hold_time",
					Description: `BGP Hold Time.`,
				},
				resource.Attribute{
					Name:        "enable_transit_summarize_cidr_to_tgw",
					Description: `Status of transit summarize CIDR to TGW.`,
				},
				resource.Attribute{
					Name:        "enable_spot_instance",
					Description: `Status of spot instance.`,
				},
				resource.Attribute{
					Name:        "spot_price",
					Description: `Price for spot instance.`,
				},
				resource.Attribute{
					Name:        "azure_eip_name_resource_group",
					Description: `The name of the public IP address and its resource group in Azure to assign to this Transit Gateway.`,
				},
				resource.Attribute{
					Name:        "ha_azure_eip_name_resource_group",
					Description: `The name of the public IP address and its resource group in Azure to assign to the HA Transit Gateway.`,
				},
				resource.Attribute{
					Name:        "local_as_number",
					Description: `Local ASN number.`,
				},
				resource.Attribute{
					Name:        "ha_security_group_id",
					Description: `HA security group used for the transit gateway.`,
				},
				resource.Attribute{
					Name:        "lan_interface_cidr",
					Description: `Transit gateway lan interface cidr.`,
				},
				resource.Attribute{
					Name:        "ha_lan_interface_cidr",
					Description: `Transit gateway lan interface cidr for the HA gateway.`,
				},
				resource.Attribute{
					Name:        "bgp_lan_ip_list",
					Description: `List of available BGP LAN interface IPs for transit external device connection creation. Only supports GCP and Azure.`,
				},
				resource.Attribute{
					Name:        "ha_bgp_lan_ip_list",
					Description: `List of available BGP LAN interface IPs for transit external device HA connection creation. Only supports GCP and Azure.`,
				},
				resource.Attribute{
					Name:        "eip",
					Description: `The EIP address of the Transit Gateway.`,
				},
				resource.Attribute{
					Name:        "ha_eip",
					Description: `The EIP address of the HA Transit Gateway. The following argument is deprecated:`,
				},
				resource.Attribute{
					Name:        "tag_list",
					Description: `Instance tag of cloud provider.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_transit_gateways",
			Category:         "Multi-Cloud Transit",
			ShortDescription: `Gets a list of all Aviatrix transit gateway's details.`,
			Description:      ``,
			Keywords: []string{
				"multi",
				"cloud",
				"transit",
				"gateways",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "gateway_list",
					Description: `The list of all transit gateways`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `Aviatrix account name.`,
				},
				resource.Attribute{
					Name:        "availability_domain",
					Description: `Availability domain for OCI.`,
				},
				resource.Attribute{
					Name:        "azure_eip_name_resource_group",
					Description: `The name of the public IP address and its resource group in Azure to assign to this Transit Gateway.`,
				},
				resource.Attribute{
					Name:        "allocate_new_eip",
					Description: `When value is false, an idle address in Elastic IP pool is reused for this gateway. Otherwise, a new Elastic IP is allocated and used for this gateway.`,
				},
				resource.Attribute{
					Name:        "bgp_ecmp",
					Description: `Enable Equal Cost Multi Path (ECMP) routing for the next hop.`,
				},
				resource.Attribute{
					Name:        "bgp_hold_time",
					Description: `BGP Hold Time.`,
				},
				resource.Attribute{
					Name:        "bgp_lan_interfaces",
					Description: `Interfaces to run BGP protocol on top of the ethernet interface, to connect to the onprem/remote peer. Only available for GCP Transit.`,
				},
				resource.Attribute{
					Name:        "bgp_lan_ip_list",
					Description: `List of available BGP LAN interface IPs for transit external device connection creation. Only supports GCP. Available as of provider version R2.21.0+.`,
				},
				resource.Attribute{
					Name:        "bgp_polling_time",
					Description: `BGP route polling time. Unit is in seconds.`,
				},
				resource.Attribute{
					Name:        "cloud_instance_id",
					Description: `Instance ID of the transit gateway.`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `Type of cloud service provider.`,
				},
				resource.Attribute{
					Name:        "connected_transit",
					Description: `Status of Connected Transit of transit gateway.`,
				},
				resource.Attribute{
					Name:        "customized_spoke_vpc_routes",
					Description: `A list of comma separated CIDRs to be customized for the spoke VPC routes.`,
				},
				resource.Attribute{
					Name:        "enable_active_standby",
					Description: `Enables Active-Standby Mode, available only with HA enabled.`,
				},
				resource.Attribute{
					Name:        "enable_active_standby_preemptive",
					Description: `Enables Preemptive Mode for Active-Standby, available only with Active-Standby enabled.`,
				},
				resource.Attribute{
					Name:        "enable_bgp_over_lan",
					Description: `Pre-allocate a network interface(eth4) for \"BGP over LAN\" functionality. Only valid for cloud_type = 4 (GCP) and 8 (Azure). Available as of provider version R2.18+`,
				},
				resource.Attribute{
					Name:        "enable_gateway_load_balancer",
					Description: `Enable firenet interfaces with AWS Gateway Load Balancer.`,
				},
				resource.Attribute{
					Name:        "enable_jumbo_frame",
					Description: `Enable jumbo frame support for transit gateway.`,
				},
				resource.Attribute{
					Name:        "enable_monitor_gateway_subnets",
					Description: `Enable [monitor gateway subnets](https://docs.aviatrix.com/HowTos/gateway.html#monitor-gateway-subnet). Only valid for cloud_type = 1 (AWS) or 256 (AWSGov).`,
				},
				resource.Attribute{
					Name:        "enable_segmentation",
					Description: `Enable segmentation to allow association of transit gateway to security domains.`,
				},
				resource.Attribute{
					Name:        "enable_spot_instance",
					Description: `Enable spot instance. NOT supported for production deployment.`,
				},
				resource.Attribute{
					Name:        "enable_transit_summarize_cidr_to_tgw",
					Description: `Enable summarize CIDR to TGW.`,
				},
				resource.Attribute{
					Name:        "enable_encrypt_volume",
					Description: `Status of Encrypt Gateway EBS Volume of the transit gateway.`,
				},
				resource.Attribute{
					Name:        "enable_hybrid_connection",
					Description: `Sign of readiness for TGW connection.`,
				},
				resource.Attribute{
					Name:        "enable_vpc_dns_server",
					Description: `Status of Vpc Dns Server of the transit Gateway.`,
				},
				resource.Attribute{
					Name:        "enable_private_oob",
					Description: `Status of private OOB for the transit gateway.`,
				},
				resource.Attribute{
					Name:        "enable_multi_tier_transit",
					Description: `Status of multi-tier transit mode on transit gateway.`,
				},
				resource.Attribute{
					Name:        "excluded_advertised_spoke_routes",
					Description: `A list of comma separated CIDRs to be advertised to on-prem as "Excluded CIDR List".`,
				},
				resource.Attribute{
					Name:        "filtered_spoke_vpc_routes",
					Description: `A list of comma separated CIDRs to be filtered from the spoke VPC route table.`,
				},
				resource.Attribute{
					Name:        "fault_domain",
					Description: `Fault domain for OCI.`,
				},
				resource.Attribute{
					Name:        "gw_size",
					Description: `Size of transit gateway instance.`,
				},
				resource.Attribute{
					Name:        "gw_name",
					Description: `Aviatrix transit gateway name.`,
				},
				resource.Attribute{
					Name:        "ha_bgp_lan_interfaces",
					Description: `Interfaces to run BGP protocol on top of the ethernet interface, to connect to the onprem/remote peer. Only available for GCP HA Transit.`,
				},
				resource.Attribute{
					Name:        "ha_bgp_lan_ip_list",
					Description: `List of available BGP LAN interface IPs for transit external device HA connection creation. Only supports GCP. Available as of provider version R2.21.0+.`,
				},
				resource.Attribute{
					Name:        "insane_mode_az",
					Description: `AZ of subnet being created for Insane Mode transit gateway.`,
				},
				resource.Attribute{
					Name:        "insane_mode",
					Description: `Status of Insane Mode of the transit gateway.`,
				},
				resource.Attribute{
					Name:        "image_version",
					Description: `The image version of the gateway.`,
				},
				resource.Attribute{
					Name:        "lan_private_subnet",
					Description: `LAN Private Subnet. Only used for GCP Transit FireNet.`,
				},
				resource.Attribute{
					Name:        "lan_vpc_id",
					Description: `LAN VPC ID. Only used for GCP Transit FireNet.`,
				},
				resource.Attribute{
					Name:        "learned_cidrs_approval_mode",
					Description: `Set the learned CIDRs approval mode.`,
				},
				resource.Attribute{
					Name:        "local_as_number",
					Description: `Changes the Aviatrix Transit Gateway ASN number before you setup Aviatrix Transit Gateway connection configurations.`,
				},
				resource.Attribute{
					Name:        "monitor_exclude_list",
					Description: `A set of monitored instance ids. Only valid when 'enable_monitor_gateway_subnets' = true.`,
				},
				resource.Attribute{
					Name:        "oob_management_subnet",
					Description: `OOB management subnet.`,
				},
				resource.Attribute{
					Name:        "oob_availability_zone",
					Description: `OOB availability zone.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Private IP address of the transit gateway created.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP address of the Transit Gateway created.`,
				},
				resource.Attribute{
					Name:        "prepend_as_path",
					Description: `List of AS numbers to populate BGP AP_PATH field when it advertises to VGW or peer devices.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `Security group used for the transit gateway.`,
				},
				resource.Attribute{
					Name:        "single_az_ha",
					Description: `Status of Single AZ HA of transit gateway.`,
				},
				resource.Attribute{
					Name:        "single_ip_snat",
					Description: `Status of Single IP Source Nat mode of the transit gateway.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `A VPC Network address range selected from one of the available network ranges.`,
				},
				resource.Attribute{
					Name:        "spot_price",
					Description: `Price for spot instance. NOT supported for production deployment.`,
				},
				resource.Attribute{
					Name:        "software_version",
					Description: `The software version of the gateway.`,
				},
				resource.Attribute{
					Name:        "tunnel_detection_time",
					Description: `The IPSec tunnel down detection time for the transit gateway.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC-ID/VNet-Name of cloud provider.`,
				},
				resource.Attribute{
					Name:        "vpc_reg",
					Description: `Region of cloud provider.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `Availability Zone. Only available for cloud_type = 8 (Azure). Must be in the form 'az-n', for example, 'az-2'.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_vpc",
			Category:         "Useful Tools",
			ShortDescription: `Gets an Aviatrix VPC's details.`,
			Description:      ``,
			Keywords: []string{
				"useful",
				"tools",
				"vpc",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Aviatrix VPC.`,
				},
				resource.Attribute{
					Name:        "route_tables_filter",
					Description: `(Optional) Filters the ` + "`" + `route_tables` + "`" + ` list to contain only public or private route tables. Valid values are 'private' or 'public'. If not set ` + "`" + `route_tables` + "`" + ` is not filtered. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `Type of cloud service provider.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `Account name of the VPC created.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Region of the VPC created.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `Subnet of the VPC created.`,
				},
				resource.Attribute{
					Name:        "subnet_size",
					Description: `Subnet size. Only supported for AWS, Azure provider.`,
				},
				resource.Attribute{
					Name:        "num_of_subnet_pairs",
					Description: `Number of public subnet and private subnet pair created. Only supported for AWS, Azure provider.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Region of the VPC created.`,
				},
				resource.Attribute{
					Name:        "aviatrix_transit_vpc",
					Description: `Switch if the VPC created is an Aviatrix Transit VPC or not.`,
				},
				resource.Attribute{
					Name:        "aviatrix_firenet_vpc",
					Description: `Switch if the VPC created is an Aviatrix FireNet VPC or not.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC created.`,
				},
				resource.Attribute{
					Name:        "resource_group",
					Description: `Resource group of the Azure VPC created.`,
				},
				resource.Attribute{
					Name:        "azure_vnet_resource_id",
					Description: `Azure vnet resource ID.`,
				},
				resource.Attribute{
					Name:        "route_tables",
					Description: `List of route table ids associated with this VPC. Only populated for AWS, AWSGov and Azure vpc.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `List of subnet of the VPC created.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `Subnet CIDR.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Subnet name.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `Subnet ID.`,
				},
				resource.Attribute{
					Name:        "public_subnets",
					Description: `List of public subnet of the VPC(AWS, Azure) created.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `Public subnet CIDR.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Public subnet name.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `Public subnet ID.`,
				},
				resource.Attribute{
					Name:        "private_subnets",
					Description: `List of private subnet of the VPC(AWS, Azure) created.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `Private subnet CIDR.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Private subnet name.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `Private subnet ID.`,
				},
				resource.Attribute{
					Name:        "availability_domains",
					Description: `List of OCI availability domains.`,
				},
				resource.Attribute{
					Name:        "fault_domains",
					Description: `List of OCI fault domains.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_type",
					Description: `Type of cloud service provider.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `Account name of the VPC created.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Region of the VPC created.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `Subnet of the VPC created.`,
				},
				resource.Attribute{
					Name:        "subnet_size",
					Description: `Subnet size. Only supported for AWS, Azure provider.`,
				},
				resource.Attribute{
					Name:        "num_of_subnet_pairs",
					Description: `Number of public subnet and private subnet pair created. Only supported for AWS, Azure provider.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Region of the VPC created.`,
				},
				resource.Attribute{
					Name:        "aviatrix_transit_vpc",
					Description: `Switch if the VPC created is an Aviatrix Transit VPC or not.`,
				},
				resource.Attribute{
					Name:        "aviatrix_firenet_vpc",
					Description: `Switch if the VPC created is an Aviatrix FireNet VPC or not.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC created.`,
				},
				resource.Attribute{
					Name:        "resource_group",
					Description: `Resource group of the Azure VPC created.`,
				},
				resource.Attribute{
					Name:        "azure_vnet_resource_id",
					Description: `Azure vnet resource ID.`,
				},
				resource.Attribute{
					Name:        "route_tables",
					Description: `List of route table ids associated with this VPC. Only populated for AWS, AWSGov and Azure vpc.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `List of subnet of the VPC created.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `Subnet CIDR.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Subnet name.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `Subnet ID.`,
				},
				resource.Attribute{
					Name:        "public_subnets",
					Description: `List of public subnet of the VPC(AWS, Azure) created.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `Public subnet CIDR.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Public subnet name.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `Public subnet ID.`,
				},
				resource.Attribute{
					Name:        "private_subnets",
					Description: `List of private subnet of the VPC(AWS, Azure) created.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `Private subnet CIDR.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Private subnet name.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `Private subnet ID.`,
				},
				resource.Attribute{
					Name:        "availability_domains",
					Description: `List of OCI availability domains.`,
				},
				resource.Attribute{
					Name:        "fault_domains",
					Description: `List of OCI fault domains.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_vpc_tracker",
			Category:         "Useful Tools",
			ShortDescription: `Gets the Aviatrix VPC Tracker information.`,
			Description:      ``,
			Keywords: []string{
				"useful",
				"tools",
				"vpc",
				"tracker",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_type",
					Description: `(Optional) Filters VPC list by cloud provider id. For example, cloud_type = 1 will give all AWS VPCs.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Optional) Filters VPC list by CIDR (AWS/Azure only).`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Filters VPC list by region (AWS/Azure only).`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `(Optional) Filters VPC list by access account name. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "vpc_list",
					Description: `List of VPCs from the VPC tracker.`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `Cloud provider id hosting this VPC.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC id.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `Aviatrix access account associated with the VPC.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `VPC region (AWS/Azure only).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `VPC name.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `VPC CIDR (AWS/Azure only).`,
				},
				resource.Attribute{
					Name:        "instance_count",
					Description: `Number of running instances in the VPC.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `List of subnets within this VPC (GCP only).`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Subnet region.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Subnet name.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `Subnet CIDR.`,
				},
				resource.Attribute{
					Name:        "gw_ip",
					Description: `Subnet gateway ip. ## Notes`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vpc_list",
					Description: `List of VPCs from the VPC tracker.`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `Cloud provider id hosting this VPC.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC id.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `Aviatrix access account associated with the VPC.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `VPC region (AWS/Azure only).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `VPC name.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `VPC CIDR (AWS/Azure only).`,
				},
				resource.Attribute{
					Name:        "instance_count",
					Description: `Number of running instances in the VPC.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `List of subnets within this VPC (GCP only).`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Subnet region.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Subnet name.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `Subnet CIDR.`,
				},
				resource.Attribute{
					Name:        "gw_ip",
					Description: `Subnet gateway ip. ## Notes`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"aviatrix_aviatrix_account":                              0,
		"aviatrix_aviatrix_caller_identity":                      1,
		"aviatrix_aviatrix_controller_metadata":                  2,
		"aviatrix_aviatrix_device_interfaces":                    3,
		"aviatrix_aviatrix_edge_gateway_wan_interface_discovery": 4,
		"aviatrix_aviatrix_firenet":                              5,
		"aviatrix_aviatrix_firenet_firewall_manager":             6,
		"aviatrix_aviatrix_firenet_vendor_integration":           7,
		"aviatrix_aviatrix_firewall":                             8,
		"aviatrix_aviatrix_firewall_instance_images":             9,
		"aviatrix_aviatrix_gateway":                              10,
		"aviatrix_aviatrix_gateway_image":                        11,
		"aviatrix_aviatrix_network_domains":                      12,
		"aviatrix_aviatrix_spoke_gateway":                        13,
		"aviatrix_aviatrix_spoke_gateway_inspection_subnets":     14,
		"aviatrix_aviatrix_spoke_gateways":                       15,
		"aviatrix_aviatrix_transit_gateway":                      16,
		"aviatrix_aviatrix_transit_gateways":                     17,
		"aviatrix_aviatrix_vpc":                                  18,
		"aviatrix_aviatrix_vpc_tracker":                          19,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
