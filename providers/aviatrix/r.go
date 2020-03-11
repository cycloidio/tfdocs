package aviatrix

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_account",
			Category:         "Accounts",
			ShortDescription: `Creates and manages Aviatrix cloud accounts`,
			Description:      ``,
			Keywords: []string{
				"accounts",
				"account",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_name",
					Description: `(Required) Account name. This can be used for logging in to CloudN console or UserConnect controller.`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `(Required) Type of cloud service provider. Only AWS, GCP, ARM, OCI, and AWS Gov are supported currently. Enter 1 for AWS, 4 for GCP, 8 for ARM, 16 for OCI, 256 for AWS Gov. ### AWS`,
				},
				resource.Attribute{
					Name:        "aws_account_number",
					Description: `(Optional) AWS Account number to associate with Aviatrix account. Required when creating an account for AWS.`,
				},
				resource.Attribute{
					Name:        "aws_iam",
					Description: `(Optional) AWS IAM-role based flag, this option is for UserConnect.`,
				},
				resource.Attribute{
					Name:        "aws_access_key",
					Description: `(Optional) AWS Access Key. Required when aws_iam is "false" and when creating an account for AWS.`,
				},
				resource.Attribute{
					Name:        "aws_secret_key",
					Description: `(Optional) AWS Secret Key. Required when aws_iam is "false" and when creating an account for AWS.`,
				},
				resource.Attribute{
					Name:        "aws_role_app",
					Description: `(Optional) AWS App role ARN, this option is for UserConnect. Required when aws_iam is "true" and when creating an account for AWS.`,
				},
				resource.Attribute{
					Name:        "aws_role_ec2",
					Description: `(Optional) AWS EC2 role ARN, this option is for UserConnect. Required when aws_iam is "true" and when creating an account for AWS. ### Azure`,
				},
				resource.Attribute{
					Name:        "arm_subscription_id",
					Description: `(Optional) Azure ARM Subscription ID. Required when creating an account for ARM.`,
				},
				resource.Attribute{
					Name:        "arm_directory_id",
					Description: `(Optional) Azure ARM Directory ID. Required when creating an account for ARM.`,
				},
				resource.Attribute{
					Name:        "arm_application_id",
					Description: `(Optional) Azure ARM Application ID. Required when creating an account for ARM.`,
				},
				resource.Attribute{
					Name:        "arm_application_key",
					Description: `(Optional) Azure ARM Application key. Required when creating an account for ARM. ### Google Cloud`,
				},
				resource.Attribute{
					Name:        "gcloud_project_id",
					Description: `(Optional) GCloud Project ID.`,
				},
				resource.Attribute{
					Name:        "gcloud_project_credentials_filepath",
					Description: `(Optional) GCloud Project Credentials [local filepath].json. Required when creating an account for GCP. ### Oracle Cloud`,
				},
				resource.Attribute{
					Name:        "oci_tenancy_id",
					Description: `(Optional)Oracle OCI Tenancy ID. Required when creating an account for OCI.`,
				},
				resource.Attribute{
					Name:        "oci_user_id",
					Description: `(Optional) Oracle OCI User ID. Required when creating an account for OCI.`,
				},
				resource.Attribute{
					Name:        "oci_compartment_id",
					Description: `(Optional) Oracle OCI Compartment ID. Required when creating an account for OCI.`,
				},
				resource.Attribute{
					Name:        "oci_api_private_key_filepath",
					Description: `(Optional) Oracle OCI API Private Key local file path. Required when creating an account for OCI. ### AWS GovCloud`,
				},
				resource.Attribute{
					Name:        "awsgov_account_number",
					Description: `(Optional) AWS Gov Account number to associate with Aviatrix account. Required when creating an account for AWS Gov.`,
				},
				resource.Attribute{
					Name:        "awsgov_access_key",
					Description: `(Optional) AWS Access Key. Required when creating an account for AWS Gov.`,
				},
				resource.Attribute{
					Name:        "awsgov_secret_key",
					Description: `(Optional) AWS Secret Key. Required when creating an account for AWS Gov. ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_account_user",
			Category:         "Accounts",
			ShortDescription: `Creates and manages Aviatrix user accounts`,
			Description:      ``,
			Keywords: []string{
				"accounts",
				"account",
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "username",
					Description: `(Required) Name of account user to be created.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `(Required) Cloud account name of user to be created.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Required) Email of address of account user to be created.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Login password for the account user to be created. If password is changed, current account will be destroyed and a new account will be created. ## Import Instance account_user can be imported using the username (when doing import, needs to leave password argument blank), e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import aviatrix_account_user.test username ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_arm_peer",
			Category:         "Peering",
			ShortDescription: `Creates and manages Aviatrix ARM peerings`,
			Description:      ``,
			Keywords: []string{
				"peering",
				"arm",
				"peer",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_name1",
					Description: `(Required) This parameter represents the name of an Azure Cloud-Account in Aviatrix controller.`,
				},
				resource.Attribute{
					Name:        "account_name2",
					Description: `(Required) This parameter represents the name of an Azure Cloud-Account in Aviatrix controller.`,
				},
				resource.Attribute{
					Name:        "vnet_name_resource_group1",
					Description: `(Required) VNet-Name of Azure cloud. Example: "VNet_Name:Resource_Group_Name".`,
				},
				resource.Attribute{
					Name:        "vnet_name_resource_group2",
					Description: `(Required) VNet-Name of Azure cloud. Example: "VNet_Name:Resource_Group_Name".`,
				},
				resource.Attribute{
					Name:        "vnet_reg1",
					Description: `(Required) Region of Azure cloud. Example: "East US 2".`,
				},
				resource.Attribute{
					Name:        "vnet_reg2",
					Description: `(Required) Region of Azure cloud. Example: "East US 2". ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "vnet_cidr1",
					Description: `List of VNet CIDR of vnet_name_resource_group1.`,
				},
				resource.Attribute{
					Name:        "vnet_cidr2",
					Description: `List of VNet CIDR of vnet_name_resource_group2. ## Import Instance arm_peer can be imported using the vnet_name_resource_group1 and vnet_name_resource_group2, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import aviatrix_aws_peer.test vnet_name_resource_group1~vnet_name_resource_group2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vnet_cidr1",
					Description: `List of VNet CIDR of vnet_name_resource_group1.`,
				},
				resource.Attribute{
					Name:        "vnet_cidr2",
					Description: `List of VNet CIDR of vnet_name_resource_group2. ## Import Instance arm_peer can be imported using the vnet_name_resource_group1 and vnet_name_resource_group2, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import aviatrix_aws_peer.test vnet_name_resource_group1~vnet_name_resource_group2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aws_peer",
			Category:         "Peering",
			ShortDescription: `Creates and manages Aviatrix AWS peerings`,
			Description:      ``,
			Keywords: []string{
				"peering",
				"aws",
				"peer",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_name1",
					Description: `(Required) This parameter represents the name of an AWS Cloud-Account in Aviatrix controller.`,
				},
				resource.Attribute{
					Name:        "account_name2",
					Description: `(Required) This parameter represents the name of an AWS Cloud-Account in Aviatrix controller.`,
				},
				resource.Attribute{
					Name:        "vpc_id1",
					Description: `(Required) VPC-ID of AWS cloud. Example: AWS: "vpc-abcd1234".`,
				},
				resource.Attribute{
					Name:        "vpc_id2",
					Description: `(Required) VPC-ID of AWS cloud. Example: AWS: "vpc-abcd1234".`,
				},
				resource.Attribute{
					Name:        "vpc_reg1",
					Description: `(Required) Region of AWS cloud. Example: AWS: "us-east-1".`,
				},
				resource.Attribute{
					Name:        "vpc_reg2",
					Description: `(Required) Region of AWS cloud. Example: AWS: "us-east-1".`,
				},
				resource.Attribute{
					Name:        "rtb_list1",
					Description: `(Optional) List of Route table ID. Valid Values: ["all"], ["rtb-abcd1234"] OR ["rtb-abcd1234,rtb-wxyz5678"].`,
				},
				resource.Attribute{
					Name:        "rtb_list2",
					Description: `(Optional) List of Route table ID. Valid Values: ["all"], ["rtb-abcd1234"] OR ["rtb-abcd1234,rtb-wxyz5678"]. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "rtb_list1_output",
					Description: `List of route table ID of vpc_id1.`,
				},
				resource.Attribute{
					Name:        "rtb_list2_output",
					Description: `List of route table ID of vpc_id2. ## Import Instance aws_peer can be imported using the vpc_id1 and vpc_id2, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import aviatrix_aws_peer.test vpc_id1~vpc_id2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "rtb_list1_output",
					Description: `List of route table ID of vpc_id1.`,
				},
				resource.Attribute{
					Name:        "rtb_list2_output",
					Description: `List of route table ID of vpc_id2. ## Import Instance aws_peer can be imported using the vpc_id1 and vpc_id2, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import aviatrix_aws_peer.test vpc_id1~vpc_id2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aws_tgw",
			Category:         "TGW Orchestrator",
			ShortDescription: `Creates and manages Aviatrix AWS TGWs`,
			Description:      ``,
			Keywords: []string{
				"tgw",
				"orchestrator",
				"aws",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tgw_name",
					Description: `(Required) Name of the AWS TGW which is going to be created.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `(Required) This parameter represents the name of a Cloud-Account in Aviatrix controller.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Region of cloud provider(AWS).`,
				},
				resource.Attribute{
					Name:        "aws_side_as_number",
					Description: `(Required) BGP Local ASN (Autonomous System Number). Integer between 1-65535. Example: "65001".`,
				},
				resource.Attribute{
					Name:        "security_domains",
					Description: `(Required) Security Domains to create together with AWS TGW's creation. Three default domains are created automatically together with the AWS TGW's creation, so are the connections between any two of them. These three domains can't be deleted, but the connection between any two of them can be deleted.`,
				},
				resource.Attribute{
					Name:        "security_domain_name",
					Description: `(Required) Three default domains ("Aviatrix_Edge_Domain", "Default_Domain" and "Shared_Service_Domain") are required with AWS TGW's creation.`,
				},
				resource.Attribute{
					Name:        "aviatrix_firewall",
					Description: `(Optional) Set to true if the security domain is an aviatrix firewall domain. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "native_egress",
					Description: `(Optional) Set to true if the security domain is a native egress domain. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "native_firewall",
					Description: `(Optional) Set to true if the security domain is a native firewall domain. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "connected_domains",
					Description: `(Optional) A list of domains connected to the domain (name: ` + "`" + `security_domain_name` + "`" + `) together with its creation.`,
				},
				resource.Attribute{
					Name:        "attached_vpc",
					Description: `(Optional) A list of VPCs attached to the domain (name: ` + "`" + `security_domain_name` + "`" + `) together with its creation. This list needs to be null for "Aviatrix_Edge_Domain".`,
				},
				resource.Attribute{
					Name:        "vpc_region",
					Description: `(Required) Region of the vpc, needs to be consistent with AWS TGW's region.`,
				},
				resource.Attribute{
					Name:        "vpc_account_name",
					Description: `(Required) This parameter represents the name of a Cloud-Account in Aviatrix controller.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) This parameter represents the ID of the VPC which is going to be attached to the security domain (name: ` + "`" + `security_domain_name` + "`" + `) which is going to be created.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `(Optional) Advanced option. VPC subnets separated by ',' to attach to the VPC. If left blank, Aviatrix Controller automatically selects a subnet representing each AZ for the VPC attachment. Example: "subnet-214f5646,subnet-085e8c81a89d70846".`,
				},
				resource.Attribute{
					Name:        "route_tables",
					Description: `(Optional) Advanced option. Route tables separated by ',' to participate in TGW Orchestrator, i.e., learned routes will be propagated to these route tables. Example: "rtb-212ff547,rtb-045397874c170c745".`,
				},
				resource.Attribute{
					Name:        "customized_routes",
					Description: `(Optional) Advanced option. Customized Spoke VPC Routes. It allows the admin to enter non-RFC1918 routes in the VPC route table targeting the TGW. Example: "10.8.0.0/16,10.9.0.0/16,10.10.0.0/16".`,
				},
				resource.Attribute{
					Name:        "customized_route_advertisement",
					Description: `(Optional) Advanced option. Customized route(s) to advertise. Example: "10.8.0.0/16,10.9.0.0/16,10.10.0.0/16".`,
				},
				resource.Attribute{
					Name:        "disable_local_route_propagation",
					Description: `(Optional) Advanced option. Switch to allow admin not to propagate the VPC CIDR to the security domain/TGW route table that it is being attached to. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "attached_aviatrix_transit_gateway",
					Description: `(Optional) A list of Names of Aviatrix Transit Gateway to attach to one of the three default domains: Aviatrix_Edge_Domain.`,
				},
				resource.Attribute{
					Name:        "manage_vpc_attachment",
					Description: `(Optional) This parameter is a switch used to allow attaching VPCs to tgw using the aviatrix_aws_tgw resource. If it is set to false, attachment of VPC must be done using the aviatrix_aws_tgw_vpc_attachment resource. Valid values: true or false. Default value is true. ->`,
				},
				resource.Attribute{
					Name:        "manage_vpc_attachment",
					Description: `If you are using/upgraded to Aviatrix Terraform Provider R1.5+, and an aws_tgw resource was originally created with a provider version <R1.5, you must do ‘terraform refresh’ to update and apply the attribute’s default value (true) into the state file. ## Import Instance aws_tgw can be imported using the tgw_name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import aviatrix_aws_tgw.test tgw_name ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aws_tgw_directconnect",
			Category:         "TGW Orchestrator",
			ShortDescription: `Creates and manages Aviatrix AWS TGW DirectConnects`,
			Description:      ``,
			Keywords: []string{
				"tgw",
				"orchestrator",
				"aws",
				"directconnect",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tgw_name",
					Description: `(Required) This parameter represents the name of an AWS TGW.`,
				},
				resource.Attribute{
					Name:        "directconnect_account_name",
					Description: `(Required) This parameter represents the name of an Account in Aviatrix controller.`,
				},
				resource.Attribute{
					Name:        "dx_gateway_id",
					Description: `(Required) This parameter represents the name of a Direct Connect Gateway ID.`,
				},
				resource.Attribute{
					Name:        "security_domain_name",
					Description: `(Required) The name of a security domain, to which the direct connect gateway will be attached.`,
				},
				resource.Attribute{
					Name:        "allowed_prefix",
					Description: `(Required) A list of comma separated CIDRs for DXGW to advertise to remote(on-prem). ## Import Instance aws_tgw_directconnect can be imported using the tgw_name and dx_gateway_id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import aviatrix_aws_tgw_directconnect.test tgw_name~dx_gateway_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aws_tgw_vpc_attachment",
			Category:         "TGW Orchestrator",
			ShortDescription: `Manages attaching/detaching VPC to/from an AWS TGW, and FireNet Gateway to TGW Firewall Domain`,
			Description:      ``,
			Keywords: []string{
				"tgw",
				"orchestrator",
				"aws",
				"vpc",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tgw_name",
					Description: `(Required) Name of the AWS TGW.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Region of cloud provider(AWS).`,
				},
				resource.Attribute{
					Name:        "security_domain_name",
					Description: `(Required & ForceNew) The name of the security domain, to which the VPC will be attached. If changed, the VPC will be detached from the old domain, and attached to the new domain.`,
				},
				resource.Attribute{
					Name:        "vpc_account_name",
					Description: `(Required) This parameter represents the name of a Cloud-Account in Aviatrix controller, which is associated with the VPC.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) This parameter represents the ID of the VPC which is going to be attached to the security domain (name: ` + "`" + `security_domain_name` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `(Optional and ForceNew) Advanced option. VPC subnets separated by ',' to attach to the VPC. If left blank, Aviatrix Controller automatically selects a subnet representing each AZ for the VPC attachment. Example: "subnet-214f5646,subnet-085e8c81a89d70846".`,
				},
				resource.Attribute{
					Name:        "route_tables",
					Description: `(Optional and ForceNew) Advanced option. Route tables separated by ',' to participate in TGW Orchestrator, i.e., learned routes will be propagated to these route tables. Example: "rtb-212ff547,rtb-045397874c170c745".`,
				},
				resource.Attribute{
					Name:        "customized_routes",
					Description: `(Optional) Advanced option. Customized Spoke VPC Routes. It allows the admin to enter non-RFC1918 routes in the VPC route table targeting the TGW. Example: "10.8.0.0/16,10.9.0.0/16,10.10.0.0/16".`,
				},
				resource.Attribute{
					Name:        "customized_route_advertisement",
					Description: `(Optional and ForeNew) Advanced option. Customized route(s) to advertise. Example: "10.8.0.0/16,10.9.0.0/16,10.10.0.0/16".`,
				},
				resource.Attribute{
					Name:        "disable_local_route_propagation",
					Description: `(Optional and ForceNew) Switch to allow admin not to propagate the VPC CIDR to the security domain/TGW route table that it is being attached to. Valid values: true, false. Default value: false. ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aws_tgw_vpn_conn",
			Category:         "TGW Orchestrator",
			ShortDescription: `Creates and manages Aviatrix AWS TGW VPN connections`,
			Description:      ``,
			Keywords: []string{
				"tgw",
				"orchestrator",
				"aws",
				"vpn",
				"conn",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tgw_name",
					Description: `(Required) This parameter represents the name of an AWS TGW.`,
				},
				resource.Attribute{
					Name:        "route_domain_name",
					Description: `(Required) The name of a route domain, to which the vpn will be attached.`,
				},
				resource.Attribute{
					Name:        "connection_name",
					Description: `(Required) Unique name of the connection.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `(Required) Public IP address. Example: "40.0.0.0".`,
				},
				resource.Attribute{
					Name:        "connection_type",
					Description: `(Optional) Connection type. Valid values: 'dynamic', 'static'. 'dynamic' stands for a BGP VPN connection; 'static' stands for a static VPN connection. Default value: 'dynamic'. ->`,
				},
				resource.Attribute{
					Name:        "connection_type",
					Description: `If you are using/upgraded to Aviatrix Terraform Provider R2.11.0+, and an aws_tgw_vpn_conn resource (static VPN connection) was originally created with a provider version <R2.11.0, you must add ` + "`" + `connection_type = static` + "`" + ` into your configuration file and do ‘terraform refresh’ to update and apply the attribute’s value (static) into the state file.`,
				},
				resource.Attribute{
					Name:        "remote_as_number",
					Description: `(Optional) AWS side as a number. Integer between 1-65535. Example: "12". Required for a dynamic VPN connection.`,
				},
				resource.Attribute{
					Name:        "remote_cidr",
					Description: `(Optional) Remote CIDRs separated by ",". Example: AWS: "16.0.0.0/16,16.1.0.0/16". Required for a static VPN connection.`,
				},
				resource.Attribute{
					Name:        "inside_ip_cidr_tun_1",
					Description: `(Optional) Inside IP CIDR for Tunnel 1. A /30 CIDR in 169.254.0.0/16.`,
				},
				resource.Attribute{
					Name:        "pre_shared_key_tun_1",
					Description: `(Optional) Pre-Shared Key for Tunnel 1. A 8-64 character string with alphanumeric underscore(_) and dot(.). It cannot start with 0.`,
				},
				resource.Attribute{
					Name:        "inside_ip_cidr_tun_2",
					Description: `(Optional) Inside IP CIDR for Tunnel 2. A /30 CIDR in 169.254.0.0/16.`,
				},
				resource.Attribute{
					Name:        "pre_shared_key_tun_2",
					Description: `(Optional) Pre-Shared Key for Tunnel 2. A 8-64 character string with alphanumeric underscore(_) and dot(.). It cannot start with 0. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "vpn_id",
					Description: `ID of the VPN generated by creation of the connection. ## Import Instance aws_tgw_vpn_conn can be imported using the tgw_name and vpn_id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import aviatrix_aws_tgw_vpn_conn.test tgw_name~vpn_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vpn_id",
					Description: `ID of the VPN generated by creation of the connection. ## Import Instance aws_tgw_vpn_conn can be imported using the tgw_name and vpn_id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import aviatrix_aws_tgw_vpn_conn.test tgw_name~vpn_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_controller_config",
			Category:         "Settings",
			ShortDescription: `Creates and manages an Aviatrix controller config resource`,
			Description:      ``,
			Keywords: []string{
				"settings",
				"controller",
				"config",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "sg_management_account_name",
					Description: `(Optional) Cloud account name of user.`,
				},
				resource.Attribute{
					Name:        "http_access",
					Description: `(Optional) Switch for http access. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "fqdn_exception_rule",
					Description: `(Optional) A system-wide mode. Valida values: true, false. Defaultvalue: true.`,
				},
				resource.Attribute{
					Name:        "security_group_management",
					Description: `(Optional) Used to manage the Controller instance’s inbound rules from gateways. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "target_version",
					Description: `(Optional) The release version number to which the controller will be upgraded to. If not specified, controller will not be upgraded. If set to "latest", controller will be upgraded to the latest release. Please look at https://docs.aviatrix.com/HowTos/inline_upgrade.html for more information.`,
				},
				resource.Attribute{
					Name:        "backup_configuration",
					Description: `(Optional) Switch to enable/disable controller cloudn backup config. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "backup_cloud_type",
					Description: `(Optional) Type of cloud service provider, requires an integer value. Use 1 for AWS.`,
				},
				resource.Attribute{
					Name:        "backup_account_name",
					Description: `(Optional) This parameter represents the name of a Cloud-Account in Aviatrix controller.`,
				},
				resource.Attribute{
					Name:        "backup_bucket_name",
					Description: `(Optional) S3 Bucket Name for AWS.`,
				},
				resource.Attribute{
					Name:        "multiple_backups",
					Description: `(Optional) Switch to enable the controller to backup up to a maximum of 3 rotating backups. Valid values: true, false. Default value: false. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Current version of the controller. ## Import Instance controller_config can be imported using controller IP, e.g. controller IP is : 10.11.12.13 ` + "`" + `` + "`" + `` + "`" + ` $ terraform import aviatrix_controller_config.test 10-11-12-13 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "version",
					Description: `Current version of the controller. ## Import Instance controller_config can be imported using controller IP, e.g. controller IP is : 10.11.12.13 ` + "`" + `` + "`" + `` + "`" + ` $ terraform import aviatrix_controller_config.test 10-11-12-13 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_firenet",
			Category:         "Firewall Network",
			ShortDescription: `Creates and manages Aviatrix FireNets`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"network",
				"firenet",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) ID of the Security VPC.`,
				},
				resource.Attribute{
					Name:        "inspection_enabled",
					Description: `(Optional) Enable/Disable traffic inspection. Valid values: true, false. Default value: true.`,
				},
				resource.Attribute{
					Name:        "egress_enabled",
					Description: `(Optional) Enable/Disable egress through firewall. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "firewall_instance_association",
					Description: `(Optional) List of firewall instances to be associated with fireNet.`,
				},
				resource.Attribute{
					Name:        "firenet_gw_name",
					Description: `(Required) Name of the primary FireNet gateway.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) ID of Firewall instance, if associating FQDN gateway to fireNet, it is FQDN gateway's gw_name..`,
				},
				resource.Attribute{
					Name:        "vendor_type",
					Description: `(Optional) Type of the firewall. Valid values: "Generic", "fqdn_gateway". Default value: "Generic". Value "fqdn_gateway" is required for FQDN gateway.`,
				},
				resource.Attribute{
					Name:        "firewall_name",
					Description: `(Optional) Firewall instance name, required if it is a firewall instance.`,
				},
				resource.Attribute{
					Name:        "management_interface",
					Description: `(Optional) Management interface ID, required if it is a firewall instance.`,
				},
				resource.Attribute{
					Name:        "inspection_enabled",
					Description: `Default value is true for associating firewall instance to fireNet. Only false is supported for associating FQDN gateway to fireNet. ->`,
				},
				resource.Attribute{
					Name:        "egress_enabled",
					Description: `Default value is false for associating firewall instance to fireNet. Only true is supported for associating FQDN gateway to fireNet. ->`,
				},
				resource.Attribute{
					Name:        "firewall_instance_association",
					Description: `If associating FQDN gateway to fireNet, "single_az_ha" needs to be enabled for the FQDN gateway. ## Import Instance firenet can be imported using the vpc_id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import aviatrix_firenet.test vpc_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_firewall",
			Category:         "Security",
			ShortDescription: `Creates and manages Aviatrix Firewall Policies`,
			Description:      ``,
			Keywords: []string{
				"security",
				"firewall",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "gw_name",
					Description: `(Required) The name of gateway.`,
				},
				resource.Attribute{
					Name:        "base_policy",
					Description: `(Optional) New base policy. Valid Values: "allow-all", "deny-all".`,
				},
				resource.Attribute{
					Name:        "base_log_enabled",
					Description: `(Optional) Indicates whether enable logging or not. Valid Values: true, false.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Optional) New access policy for the gateway. Type: String (valid JSON). Seven fields are required for each policy item: src_ip, dst_ip, protocol, port, allow_deny, log_enabled and description.`,
				},
				resource.Attribute{
					Name:        "src_ip",
					Description: `(Required) CIDRs separated by comma or tag names such "HR" or "marketing" etc. Example: "10.30.0.0/16,10.45.0.0/20". The aviatrix_firewall_tag resource should be created prior to using the tag name.`,
				},
				resource.Attribute{
					Name:        "dst_ip",
					Description: `(Required) CIDRs separated by comma or tag names such "HR" or "marketing" etc. Example: "10.30.0.0/16,10.45.0.0/20". The aviatrix_firewall_tag resource should be created prior to using the tag name.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) a single port or a range of port numbers. Example: "25", "25:1024".`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_firewall_instance",
			Category:         "Firewall Network",
			ShortDescription: `Creates and deletes Aviatrix Firewall Instances`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"network",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) ID of the Security VPC.`,
				},
				resource.Attribute{
					Name:        "firenet_gw_name",
					Description: `(Required) Name of the primary FireNet gateway.`,
				},
				resource.Attribute{
					Name:        "firewall_name",
					Description: `(Required) Name of the firewall instance to be created.`,
				},
				resource.Attribute{
					Name:        "firewall_image",
					Description: `(Required) One of the AWS AMIs from Palo Alto Networks.`,
				},
				resource.Attribute{
					Name:        "firewall_size",
					Description: `(Required) Instance size of the firewall. Example: "m5.xlarge".`,
				},
				resource.Attribute{
					Name:        "management_subnet",
					Description: `(Required) Management Interface Subnet. Select the subnet whose name contains “gateway and firewall management”.`,
				},
				resource.Attribute{
					Name:        "egress_subnet",
					Description: `(Required) Egress Interface Subnet. Select the subnet whose name contains “FW-ingress-egress”.`,
				},
				resource.Attribute{
					Name:        "firewall_image_version",
					Description: `(Optional) Version of firewall image.`,
				},
				resource.Attribute{
					Name:        "iam_role",
					Description: `(Optional) In advanced mode, create an IAM Role on the AWS account that launched the FireNet gateway. Create a policy to attach to the role. The policy is to allow access to “Bootstrap Bucket”.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_firewall_tag",
			Category:         "Security",
			ShortDescription: `Creates and manages Aviatrix Firewall Tags`,
			Description:      ``,
			Keywords: []string{
				"security",
				"firewall",
				"tag",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "firewall_tag",
					Description: `(Required) This parameter represents the name of a Cloud-Account in Aviatrix controller.`,
				},
				resource.Attribute{
					Name:        "cidr_list",
					Description: `(Optional) A JSON file with the following:`,
				},
				resource.Attribute{
					Name:        "cidr_tag_name",
					Description: `(Required) The name attribute of a policy. Example: "policy1".`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Required) The CIDR attribute of a policy. Example: "10.88.88.88/32". ## Import Instance firewall_tag can be imported using the firewall_tag, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import aviatrix_firewall_tag.test firewall_tag ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_fqdn",
			Category:         "Security",
			ShortDescription: `Manages Aviatrix FQDN filtering for Gateway`,
			Description:      ``,
			Keywords: []string{
				"security",
				"fqdn",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "fqdn_tag",
					Description: `(Required) FQDN Filter Tag Name.`,
				},
				resource.Attribute{
					Name:        "fqdn_enabled",
					Description: `(Optional) FQDN Filter Tag Status. Valid values: true, false.`,
				},
				resource.Attribute{
					Name:        "fqdn_mode",
					Description: `(Optional) Specify the tag color to be a white-list tag or black-list tag. Valid values: "white", "black".`,
				},
				resource.Attribute{
					Name:        "gw_filter_tag_list",
					Description: `(Optional) A list of gateways to attach to the specific tag.`,
				},
				resource.Attribute{
					Name:        "gw_name",
					Description: `(Required) Name of the gateway to attach to the specific tag.`,
				},
				resource.Attribute{
					Name:        "source_ip_list",
					Description: `(Optional) List of source IPs in the VPC qualified for a specific tag.`,
				},
				resource.Attribute{
					Name:        "domain_names",
					Description: `(Optional) One or more domain names in a list with details as listed below:`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `(Required) FQDN. Example: "facebook.com".`,
				},
				resource.Attribute{
					Name:        "proto",
					Description: `(Required) Protocol. Valid values: "all", "tcp", "udp", "icmp".`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) Port. Example "25".`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_gateway",
			Category:         "Gateway",
			ShortDescription: `Creates and manages Aviatrix gateways`,
			Description:      ``,
			Keywords: []string{
				"gateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_type",
					Description: `(Required) Type of cloud service provider, requires an integer value. Currently only AWS(1), GCP(4), ARM(8), OCI(16), and AWSGov(256) are supported.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `(Required) Account name. This account will be used to launch Aviatrix gateway.`,
				},
				resource.Attribute{
					Name:        "gw_name",
					Description: `(Required) Aviatrix gateway unique name.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) VPC-ID/VNet-Name of cloud provider. Example: AWS: "vpc-abcd1234", GCP: "vpc-gcp-test", ARM: "vnet1:hello", OCI: "vpc-oracle-test1".`,
				},
				resource.Attribute{
					Name:        "vpc_reg",
					Description: `(Required) Region of cloud provider. Example: AWS: "us-east-1", GCP: "us-west2-a", ARM: "East US 2", Oracle: "us-ashburn-1".`,
				},
				resource.Attribute{
					Name:        "gw_size",
					Description: `(Required) Size of the gateway instance. Example: AWS: "t2.large", ARM: "Standard_B1s", Oracle: "VM.Standard2.2", GCP: "n1-standard-1".`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `(Required) A VPC Network address range selected from one of the available network ranges. Example: "172.31.0.0/20".`,
				},
				resource.Attribute{
					Name:        "peering_ha_subnet",
					Description: `(Optional) Public subnet CIDR to create Peering HA Gateway in. Required only if enabling Peering HA for AWS/ARM . Example: AWS: "10.0.0.0/16".`,
				},
				resource.Attribute{
					Name:        "peering_ha_zone",
					Description: `(Optional) Zone information for creating Peering HA Gateway, only zone is accepted. Required only if enabling Peering HA for GCP. Example: GCP: "us-west1-c".`,
				},
				resource.Attribute{
					Name:        "peering_ha_insane_mode_az",
					Description: `(Optional) AZ of subnet being created for Insane Mode Peering HA Gateway. Required for AWS only if ` + "`" + `insane_mode` + "`" + ` is set and ` + "`" + `peering_ha_subnet` + "`" + ` is set. Example: AWS: "us-west-1a".`,
				},
				resource.Attribute{
					Name:        "peering_ha_eip",
					Description: `(Optional) Public IP address that you want assigned to the HA peering instance. Only available for AWS.`,
				},
				resource.Attribute{
					Name:        "peering_ha_gw_size",
					Description: `(Optional) Size of the Peering HA Gateway to be created. Required if enabling Peering HA.`,
				},
				resource.Attribute{
					Name:        "insane_mode",
					Description: `(Optional) Enable Insane Mode for Gateway. Insane Mode Gateway size must be at least c5 (AWS) or Standard_D3_v2 (ARM). If enabled, you must specify a valid /26 CIDR segment of the VPC to create a new subnet. Only supported for AWS, AWSGov or ARM. Valid values: true, false.`,
				},
				resource.Attribute{
					Name:        "insane_mode_az",
					Description: `(Optional) AZ of subnet being created for Insane Mode Gateway. Required for AWS and AWSGov if insane_mode is set. Example: AWS: "us-west-1a". ### SNAT/DNAT`,
				},
				resource.Attribute{
					Name:        "single_ip_snat",
					Description: `(Optional) Enable Source NAT in "single ip" mode for this container. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "vpn_access",
					Description: `(Optional) Enable user access through VPN to this container. Valid values: true, false.`,
				},
				resource.Attribute{
					Name:        "vpn_cidr",
					Description: `(Optional) VPN CIDR block for the container. Required if ` + "`" + `vpn_access` + "`" + ` is true. Example: "192.168.43.0/24".`,
				},
				resource.Attribute{
					Name:        "max_vpn_conn",
					Description: `(Optional) Maximum number of active VPN users allowed to be connected to this gateway. Required if vpn_access is true. Make sure the number is smaller than the VPN CIDR block. Example: 100.`,
				},
				resource.Attribute{
					Name:        "enable_elb",
					Description: `(Optional) Specify whether to enable ELB or not. Not supported for Oracle gateways. Valid values: true, false.`,
				},
				resource.Attribute{
					Name:        "elb_name",
					Description: `(Optional) A name for the ELB that is created. If it is not specified, a name is generated automatically.`,
				},
				resource.Attribute{
					Name:        "vpn_protocol",
					Description: `(Optional) Transport mode for VPN connection. All ` + "`" + `cloud_types` + "`" + ` support TCP with ELB, and UDP without ELB. AWS(1) additionally supports UDP with ELB. Valid values: "TCP", "UDP". If not specified, "TCP" will be used.`,
				},
				resource.Attribute{
					Name:        "split_tunnel",
					Description: `(Optional) Specify split tunnel mode. Valid values: true, false.`,
				},
				resource.Attribute{
					Name:        "name_servers",
					Description: `(Optional) A list of DNS servers used to resolve domain names by a connected VPN user when Split Tunnel Mode is enabled.`,
				},
				resource.Attribute{
					Name:        "search_domains",
					Description: `(Optional) A list of domain names that will use the NameServer when a specific name is not in the destination when Split Tunnel Mode is enabled.`,
				},
				resource.Attribute{
					Name:        "additional_cidrs",
					Description: `(Optional) A list of destination CIDR ranges that will also go through the VPN tunnel when Split Tunnel Mode is enabled.`,
				},
				resource.Attribute{
					Name:        "otp_mode",
					Description: `(Optional) Two step authentication mode. "2": DUO, "3": Okta.`,
				},
				resource.Attribute{
					Name:        "saml_enabled",
					Description: `(Optional) This field indicates whether enabling SAML or not. This field is available in controller version 3.3 or later release. Valid values: true, false.`,
				},
				resource.Attribute{
					Name:        "enable_vpn_nat",
					Description: `(Optional) This field indicates whether enabling VPN NAT or not. Only supported for VPN gateway. Valid values: true, false. Default value: true.`,
				},
				resource.Attribute{
					Name:        "okta_token",
					Description: `(Optional) Token for Okta auth mode. Required if otp_mode is "3".`,
				},
				resource.Attribute{
					Name:        "okta_url",
					Description: `(Optional) URL for Okta auth mode. Required if otp_mode is "3".`,
				},
				resource.Attribute{
					Name:        "okta_username_suffix",
					Description: `(Optional) Username suffix for Okta auth mode. Example: "aviatrix.com".`,
				},
				resource.Attribute{
					Name:        "duo_integration_key",
					Description: `(Optional) Integration key for DUO auth mode. Required if otp_mode is "2".`,
				},
				resource.Attribute{
					Name:        "duo_secret_key",
					Description: `(Optional) Secret key for DUO auth mode. Required if otp_mode is "2".`,
				},
				resource.Attribute{
					Name:        "duo_api_hostname",
					Description: `(Optional) API hostname for DUO auth mode. Required: Yes if otp_mode is "2".`,
				},
				resource.Attribute{
					Name:        "duo_push_mode",
					Description: `(Optional) Push mode for DUO auth. Required if otp_mode is "2". Valid values: "auto", "selective" and "token".`,
				},
				resource.Attribute{
					Name:        "enable_ldap",
					Description: `(Optional) Specify whether to enable LDAP or not. Valid values: true, false.`,
				},
				resource.Attribute{
					Name:        "ldap_server",
					Description: `(Optional) LDAP server address. Required if enable_ldap is true.`,
				},
				resource.Attribute{
					Name:        "ldap_bind_dn",
					Description: `(Optional) LDAP bind DN. Required if enable_ldap is true.`,
				},
				resource.Attribute{
					Name:        "ldap_password",
					Description: `(Optional) LDAP password. Required if enable_ldap is true.`,
				},
				resource.Attribute{
					Name:        "ldap_base_dn",
					Description: `(Optional) LDAP base DN. Required if enable_ldap is true.`,
				},
				resource.Attribute{
					Name:        "ldap_username_attribute",
					Description: `(Optional) LDAP user attribute. Required if enable_ldap is true. ### Misc.`,
				},
				resource.Attribute{
					Name:        "allocate_new_eip",
					Description: `(Optional) When value is false, reuse an idle address in Elastic IP pool for this gateway. Otherwise, allocate a new Elastic IP and use it for this gateway. Available in Controller 2.7+. Valid values: true, false. Default: true. Option not available for GCP, ARM and OCI gateways, they will automatically allocate new EIPs.`,
				},
				resource.Attribute{
					Name:        "eip",
					Description: `(Optional) Required when ` + "`" + `allocate_new_eip` + "`" + ` is false. It uses the specified EIP for this gateway. Available in Controller version 3.5+. Only available for AWS.`,
				},
				resource.Attribute{
					Name:        "tag_list",
					Description: `(Optional) Instance tag of cloud provider. Only available for AWS and AWSGov. Example: ["key1:value1", "key2:value2"].`,
				},
				resource.Attribute{
					Name:        "enable_vpc_dns_server",
					Description: `(Optional) Enable VPC DNS Server for Gateway. Currently only supports AWS and AWSGov. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "enable_designated_gateway",
					Description: `(Optional) Enable 'designated_gateway' feature for Gateway. Only supports AWS. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "additional_cidrs_designated_gateway",
					Description: `(Optional) A list of CIDR ranges separated by comma to configure when 'designated_gateway' feature is enabled. Example: "10.8.0.0/16,10.9.0.0/16,10.10.0.0/16".`,
				},
				resource.Attribute{
					Name:        "enable_encrypt_volume",
					Description: `(Optional) Enable EBS volume encryption for Gateway. Only supports AWS. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "customer_managed_keys",
					Description: `(Optional and Sensitive) Customer managed key ID. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "elb_dns_name",
					Description: `ELB DNS name.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP address of the Gateway created.`,
				},
				resource.Attribute{
					Name:        "backup_public_ip",
					Description: `Private IP address of the Gateway created.`,
				},
				resource.Attribute{
					Name:        "public_dns_server",
					Description: `DNS server used by the gateway. Default is "8.8.8.8", can be overridden with the VPC's setting.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `Security group used for the gateway.`,
				},
				resource.Attribute{
					Name:        "cloud_instance_id",
					Description: `Instance ID of the gateway.`,
				},
				resource.Attribute{
					Name:        "cloudn_bkup_gateway_inst_id",
					Description: `Instance ID of the backup gateway. The following arguments are deprecated:`,
				},
				resource.Attribute{
					Name:        "dns_server",
					Description: `Specify the DNS IP, only required while using a custom private DNS for the VPC.`,
				},
				resource.Attribute{
					Name:        "enable_snat",
					Description: `(Optional) Enable Source NAT for this container. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "dnat_policy",
					Description: `(Optional) Policy rule applied for enabling Destination NAT (DNAT), which allows you to change the destination to a virtual address range. Currently only supports AWS(1) and ARM(8).`,
				},
				resource.Attribute{
					Name:        "src_ip",
					Description: `(Optional) A source IP address range where the policy rule applies.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `(Optional) A source port that the policy rule applies.`,
				},
				resource.Attribute{
					Name:        "dst_ip",
					Description: `(Optional) A destination IP address range where the policy rule applies.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `(Optional) A destination port where the policy rule applies.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) A destination port protocol where the policy rule applies.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(Optional) An output interface where the policy rule applies.`,
				},
				resource.Attribute{
					Name:        "connection",
					Description: `(Optional) Default value: "None".`,
				},
				resource.Attribute{
					Name:        "mark",
					Description: `(Optional) A tag or mark of a TCP session where the policy rule applies.`,
				},
				resource.Attribute{
					Name:        "new_src_ip",
					Description: `(Optional) The changed source IP address when all specified qualifier conditions meet. One of the rule fields must be specified for this rule to take effect.`,
				},
				resource.Attribute{
					Name:        "new_src_port",
					Description: `(Optional) The translated destination port when all specified qualifier conditions meet. One of the rule field must be specified for this rule to take effect.`,
				},
				resource.Attribute{
					Name:        "exclude_rtb",
					Description: `(Optional) This field specifies which VPC private route table will not be programmed with the default route entry. ## Import Instance gateway can be imported using the gw_name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import aviatrix_gateway.test gw_name ` + "`" + `` + "`" + `` + "`" + ` ## Notes ### FQDN In order for the FQDN feature to be enabled for the specified gateway, ` + "`" + `single_ip_snat` + "`" + ` must be set to true. If it is not set at gateway creation, creation of FQDN resource will automatically enable SNAT and users must rectify the diff in the Terraform state by setting ` + "`" + `single_ip_snat = true` + "`" + ` in their config file. ### insane_mode If ` + "`" + `insane_mode` + "`" + ` is enabled, you must specify a valid /26 CIDR segment of the VPC specified for the ` + "`" + `subnet` + "`" + `. This will then create a new subnet to be used for the corresponding gateway. You cannot specify an existing /26 subnet. ### max_vpn_conn If you are using/upgraded to Aviatrix Terraform Provider R1.14+, and a gateway with VPN enabled was originally created with a provider version <R1.14, you must do a ‘terraform refresh’ to update and apply the attribute’s value into the state. In addition, you must also input this attribute and its value to "100" in your ` + "`" + `.tf` + "`" + ` file. ### peering_ha_gw_size If you are using/upgraded to Aviatrix Terraform Provider R1.8+, and a peering-HA gateway was originally created with a provider version <R1.8, you must do a ‘terraform refresh’ to update and apply the attribute’s value into the state. In addition, you must also input this attribute and its value to its corresponding gateway resource in your ` + "`" + `.tf` + "`" + ` file. ### enable_snat If you are using/upgraded to Aviatrix Terraform Provider R2.10+, and a gateway with ` + "`" + `enable_snat` + "`" + ` set to true was originally created with a provider version <R2.10, you must do a ‘terraform refresh’ to update and apply the attribute’s value into the state. In addition, you must also change this attribute to ` + "`" + `single_ip_snat` + "`" + ` in your ` + "`" + `.tf` + "`" + ` file. ### dnat_policy If you are using/upgraded to Aviatrix Terraform Provider R2.10+, and a gateway with ` + "`" + `dnat_policy` + "`" + ` was originally created with a provider version <R2.10, you must do a ‘terraform refresh’ to remove attribute’s value from the state. In addition, you must transfer its corresponding values to the`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "elb_dns_name",
					Description: `ELB DNS name.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP address of the Gateway created.`,
				},
				resource.Attribute{
					Name:        "backup_public_ip",
					Description: `Private IP address of the Gateway created.`,
				},
				resource.Attribute{
					Name:        "public_dns_server",
					Description: `DNS server used by the gateway. Default is "8.8.8.8", can be overridden with the VPC's setting.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `Security group used for the gateway.`,
				},
				resource.Attribute{
					Name:        "cloud_instance_id",
					Description: `Instance ID of the gateway.`,
				},
				resource.Attribute{
					Name:        "cloudn_bkup_gateway_inst_id",
					Description: `Instance ID of the backup gateway. The following arguments are deprecated:`,
				},
				resource.Attribute{
					Name:        "dns_server",
					Description: `Specify the DNS IP, only required while using a custom private DNS for the VPC.`,
				},
				resource.Attribute{
					Name:        "enable_snat",
					Description: `(Optional) Enable Source NAT for this container. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "dnat_policy",
					Description: `(Optional) Policy rule applied for enabling Destination NAT (DNAT), which allows you to change the destination to a virtual address range. Currently only supports AWS(1) and ARM(8).`,
				},
				resource.Attribute{
					Name:        "src_ip",
					Description: `(Optional) A source IP address range where the policy rule applies.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `(Optional) A source port that the policy rule applies.`,
				},
				resource.Attribute{
					Name:        "dst_ip",
					Description: `(Optional) A destination IP address range where the policy rule applies.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `(Optional) A destination port where the policy rule applies.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) A destination port protocol where the policy rule applies.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(Optional) An output interface where the policy rule applies.`,
				},
				resource.Attribute{
					Name:        "connection",
					Description: `(Optional) Default value: "None".`,
				},
				resource.Attribute{
					Name:        "mark",
					Description: `(Optional) A tag or mark of a TCP session where the policy rule applies.`,
				},
				resource.Attribute{
					Name:        "new_src_ip",
					Description: `(Optional) The changed source IP address when all specified qualifier conditions meet. One of the rule fields must be specified for this rule to take effect.`,
				},
				resource.Attribute{
					Name:        "new_src_port",
					Description: `(Optional) The translated destination port when all specified qualifier conditions meet. One of the rule field must be specified for this rule to take effect.`,
				},
				resource.Attribute{
					Name:        "exclude_rtb",
					Description: `(Optional) This field specifies which VPC private route table will not be programmed with the default route entry. ## Import Instance gateway can be imported using the gw_name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import aviatrix_gateway.test gw_name ` + "`" + `` + "`" + `` + "`" + ` ## Notes ### FQDN In order for the FQDN feature to be enabled for the specified gateway, ` + "`" + `single_ip_snat` + "`" + ` must be set to true. If it is not set at gateway creation, creation of FQDN resource will automatically enable SNAT and users must rectify the diff in the Terraform state by setting ` + "`" + `single_ip_snat = true` + "`" + ` in their config file. ### insane_mode If ` + "`" + `insane_mode` + "`" + ` is enabled, you must specify a valid /26 CIDR segment of the VPC specified for the ` + "`" + `subnet` + "`" + `. This will then create a new subnet to be used for the corresponding gateway. You cannot specify an existing /26 subnet. ### max_vpn_conn If you are using/upgraded to Aviatrix Terraform Provider R1.14+, and a gateway with VPN enabled was originally created with a provider version <R1.14, you must do a ‘terraform refresh’ to update and apply the attribute’s value into the state. In addition, you must also input this attribute and its value to "100" in your ` + "`" + `.tf` + "`" + ` file. ### peering_ha_gw_size If you are using/upgraded to Aviatrix Terraform Provider R1.8+, and a peering-HA gateway was originally created with a provider version <R1.8, you must do a ‘terraform refresh’ to update and apply the attribute’s value into the state. In addition, you must also input this attribute and its value to its corresponding gateway resource in your ` + "`" + `.tf` + "`" + ` file. ### enable_snat If you are using/upgraded to Aviatrix Terraform Provider R2.10+, and a gateway with ` + "`" + `enable_snat` + "`" + ` set to true was originally created with a provider version <R2.10, you must do a ‘terraform refresh’ to update and apply the attribute’s value into the state. In addition, you must also change this attribute to ` + "`" + `single_ip_snat` + "`" + ` in your ` + "`" + `.tf` + "`" + ` file. ### dnat_policy If you are using/upgraded to Aviatrix Terraform Provider R2.10+, and a gateway with ` + "`" + `dnat_policy` + "`" + ` was originally created with a provider version <R2.10, you must do a ‘terraform refresh’ to remove attribute’s value from the state. In addition, you must transfer its corresponding values to the`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_gateway_dnat",
			Category:         "Gateway",
			ShortDescription: `Configure policies for destination NAT for an Aviatrix gateway`,
			Description:      ``,
			Keywords: []string{
				"gateway",
				"dnat",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "gw_name",
					Description: `(Required) Aviatrix gateway unique name.`,
				},
				resource.Attribute{
					Name:        "dnat_policy",
					Description: `(Required) Policy rule applied for enabling Destination NAT (DNAT), which allows you to change the destination to a virtual address range. Currently only supports AWS(1) and ARM(8).`,
				},
				resource.Attribute{
					Name:        "src_cidr",
					Description: `(Optional) This is a qualifier condition that specifies a source IP address range where the rule applies. When left blank, this field is not used.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `(Optional) This is a qualifier condition that specifies a source port that the rule applies. When left blank, this field is not used.`,
				},
				resource.Attribute{
					Name:        "dst_cidr",
					Description: `(Optional) This is a qualifier condition that specifies a destination IP address range where the rule applies. When left blank, this field is not used.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `(Optional) This is a qualifier condition that specifies a destination port where the rule applies. When left blank, this field is not used.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) This is a qualifier condition that specifies a destination port protocol where the rule applies. When left blank, this field is not used.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(Optional) This is a qualifier condition that specifies output interface where the rule applies. When left blank, this field is not used.`,
				},
				resource.Attribute{
					Name:        "connection",
					Description: `(Optional) Default value: "None".`,
				},
				resource.Attribute{
					Name:        "mark",
					Description: `(Optional) This is a rule field that specifies a tag or mark of a TCP session when all qualifier conditions meet. When left blank, this field is not used.`,
				},
				resource.Attribute{
					Name:        "dnat_ips",
					Description: `(Optional) This is a rule field that specifies the translated destination IP address when all specified qualifier conditions meet. When left blank, this field is not used. One of the rule field must be specified for this rule to take effect.`,
				},
				resource.Attribute{
					Name:        "dnat_port",
					Description: `(Optional) This is a rule field that specifies the translated destination port when all specified qualifier conditions meet. When left blank, this field is not used. One of the rule field must be specified for this rule to take effect.`,
				},
				resource.Attribute{
					Name:        "exclude_rtb",
					Description: `(Optional) This field specifies which VPC private route table will not be programmed with the default route entry. ## Import Instance gateway_dnat can be imported using the gw_name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import aviatrix_gateway_dnat.test gw_name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_gateway_snat",
			Category:         "Gateway",
			ShortDescription: `Configure customized SNAT policies for an Aviatrix gateway`,
			Description:      ``,
			Keywords: []string{
				"gateway",
				"snat",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "gw_name",
					Description: `(Required) Aviatrix gateway unique name.`,
				},
				resource.Attribute{
					Name:        "snat_mode",
					Description: `(Optional) NAT mode. Valid values: "customized_snat". Default value: "customized_snat".`,
				},
				resource.Attribute{
					Name:        "snat_policy",
					Description: `(Required) Policy rule applied for enabling source NAT (mode: "customized_snat"). Currently only supports AWS(1) and ARM(8).`,
				},
				resource.Attribute{
					Name:        "src_cidr",
					Description: `(Optional) This is a qualifier condition that specifies a source IP address range where the rule applies. When left blank, this field is not used.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `(Optional) This is a qualifier condition that specifies a source port that the rule applies. When left blank, this field is not used.`,
				},
				resource.Attribute{
					Name:        "dst_cidr",
					Description: `(Optional) This is a qualifier condition that specifies a destination IP address range where the rule applies. When left blank, this field is not used.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `(Optional) This is a qualifier condition that specifies a destination port where the rule applies. When left blank, this field is not used.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) This is a qualifier condition that specifies a destination port protocol where the rule applies. When left blank, this field is not used.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(Optional) This is a qualifier condition that specifies output interface where the rule applies. When left blank, this field is not used.`,
				},
				resource.Attribute{
					Name:        "connection",
					Description: `(Optional) Default value: "None".`,
				},
				resource.Attribute{
					Name:        "mark",
					Description: `(Optional) This is a qualifier condition that specifies a tag or mark of a TCP session where the rule applies. When left blank, this field is not used.`,
				},
				resource.Attribute{
					Name:        "snat_ips",
					Description: `(Optional) This is a rule field that specifies the changed source IP address when all specified qualifier conditions meet. When left blank, this field is not used. One of the rule fields must be specified for this rule to take effect.`,
				},
				resource.Attribute{
					Name:        "snat_port",
					Description: `(Optional) This is a rule field that specifies the changed source port when all specified qualifier conditions meet. When left blank, this field is not used. One of the rule fields must be specified for this rule to take effect.`,
				},
				resource.Attribute{
					Name:        "exclude_rtb",
					Description: `(Optional) This field specifies which VPC private route table will not be programmed with the default route entry. ## Import Instance gateway_snat can be imported using the gw_name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import aviatrix_gateway_snat.test gw_name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_geo_vpn",
			Category:         "OpenVPN",
			ShortDescription: `Enables and manages the Aviatrix Geo VPN`,
			Description:      ``,
			Keywords: []string{
				"openvpn",
				"geo",
				"vpn",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_type",
					Description: `(Required) Type of cloud service provider, requires an integer value. Currently only AWS(1) is supported.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `(Required) This parameter represents the name of a Cloud-Account in Aviatrix controller.`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `(Required) The hosted domain name. It must be hosted by AWS Route53 or Azure DNS in the selected account.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The hostname that users will connect to. A DNS record will be created for this name in the specified domain name.`,
				},
				resource.Attribute{
					Name:        "elb_dns_names",
					Description: `(Required) List of ELB names to attach to this Geo VPN name. ## Import Instance geo_vpn can be imported using the service name and domain name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import aviatrix_geo_vpn.test service_name~domain_name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_saml_endpoint",
			Category:         "OpenVPN",
			ShortDescription: `Creates and manages an Aviatrix SAML Endpoint`,
			Description:      ``,
			Keywords: []string{
				"openvpn",
				"saml",
				"endpoint",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "endpoint_name",
					Description: `(Required) The SAML Endpoint name.`,
				},
				resource.Attribute{
					Name:        "idp_metadata_type",
					Description: `(Required) The IDP Metadata type. At the moment only "Text" is supported.`,
				},
				resource.Attribute{
					Name:        "idp_metadata",
					Description: `(Required) The IDP Metadata from SAML provider. Normally the metadata is in XML format which may contain special characters. Best practice is encode metadata in base64 and set here ` + "`" + `${base64decode(var.idp_metadata)}` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "custom_entity_id",
					Description: `(Optional) Custom Entity ID. Required to be non-empty for 'Custom' Entity ID type, empty for 'Hostname' Entity ID type.`,
				},
				resource.Attribute{
					Name:        "custom_saml_request_template",
					Description: `(Optional) Custom SAML Request Template in string. ## Import Instance saml_endpoint can be imported using the SAML Endpoint name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import aviatrix_saml_endpoint.test saml-test ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_site2cloud",
			Category:         "Site2Cloud",
			ShortDescription: `Creates and Manages Aviatrix Site2Cloud connections`,
			Description:      ``,
			Keywords: []string{
				"site2cloud",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) VPC Id of the cloud gateway.`,
				},
				resource.Attribute{
					Name:        "connection_name",
					Description: `(Required) Site2Cloud Connection Name.`,
				},
				resource.Attribute{
					Name:        "remote_gateway_type",
					Description: `(Required) Remote Gateway Type. Valid Values: "generic", "avx", "aws", "azure", "sonicwall", "oracle".`,
				},
				resource.Attribute{
					Name:        "connection_type",
					Description: `(Required) Connection Type. Valid Values: "mapped", "unmapped".`,
				},
				resource.Attribute{
					Name:        "tunnel_type",
					Description: `(Required) Site2Cloud Tunnel Type. Valid Values: "udp", "tcp".`,
				},
				resource.Attribute{
					Name:        "primary_cloud_gateway_name",
					Description: `(Required) Primary Cloud Gateway Name.`,
				},
				resource.Attribute{
					Name:        "remote_gateway_ip",
					Description: `(Required) Remote Gateway IP.`,
				},
				resource.Attribute{
					Name:        "remote_subnet_cidr",
					Description: `(Required) Remote Subnet CIDR.`,
				},
				resource.Attribute{
					Name:        "remote_subnet_virtual",
					Description: `Remote Subnet CIDR (Virtual). Required for connection type "mapped" only.`,
				},
				resource.Attribute{
					Name:        "local_subnet_cidr",
					Description: `(Optional) Local Subnet CIDR. Required for connection type "mapped".`,
				},
				resource.Attribute{
					Name:        "local_subnet_virtual",
					Description: `Local Subnet CIDR (Virtual). Required for connection type "mapped" only. ## HA`,
				},
				resource.Attribute{
					Name:        "ha_enabled",
					Description: `(Optional) Specify whether or not to enable HA. Valid Values: true, false.`,
				},
				resource.Attribute{
					Name:        "backup_gateway_name",
					Description: `(Optional) Backup gateway name.`,
				},
				resource.Attribute{
					Name:        "backup_remote_gateway_ip",
					Description: `(Optional) Backup Remote Gateway IP.`,
				},
				resource.Attribute{
					Name:        "backup_pre_shared_key",
					Description: `(Optional) Backup Pre-Shared Key. ### Custom Algorithms`,
				},
				resource.Attribute{
					Name:        "custom_algorithms",
					Description: `(Optional) Switch to enable custom/non-default algorithms for IPSec Authentication/Encryption. Valid values: true, false.`,
				},
				resource.Attribute{
					Name:        "phase_1_authentication",
					Description: `(Optional) Phase one Authentication. Valid values: 'SHA-1', 'SHA-256', 'SHA-384' and 'SHA-512'. Default value: 'SHA-1'.`,
				},
				resource.Attribute{
					Name:        "phase_2_authentication",
					Description: `(Optional) Phase two Authentication. Valid values: 'NO-AUTH', 'HMAC-SHA-1', 'HMAC-SHA-256', 'HMAC-SHA-384' and 'HMAC-SHA-512'. Default value: 'HMAC-SHA-1'.`,
				},
				resource.Attribute{
					Name:        "phase_1_dh_groups",
					Description: `(Optional) Phase one DH Groups. Valid values: '1', '2', '5', '14', '15', '16', '17' and '18'. Default value: '2'.`,
				},
				resource.Attribute{
					Name:        "phase_2_dh_groups",
					Description: `(Optional) Phase two DH Groups. Valid values: '1', '2', '5', '14', '15', '16', '17' and '18'. Default value: '2'.`,
				},
				resource.Attribute{
					Name:        "phase_1_encryption",
					Description: `(Optional) Phase one Encryption. Valid values: '3DES', 'AES-128-CBC', 'AES-192-CBC' and 'AES-256-CBC'. Default value: 'AES-256-CBC'.`,
				},
				resource.Attribute{
					Name:        "phase_2_encryption",
					Description: `(Optional) Phase two Encryption. Valid values: '3DES', 'AES-128-CBC', 'AES-192-CBC', 'AES-256-CBC', 'AES-128-GCM-64', 'AES-128-GCM-96' and 'AES-128-GCM-128'. Default value: 'AES-256-CBC'. ### Encryption over ExpressRoute/DirectConnect`,
				},
				resource.Attribute{
					Name:        "private_route_encryption",
					Description: `(Optional) Private route encryption switch. Valid values: true, false.`,
				},
				resource.Attribute{
					Name:        "route_table_list",
					Description: `(Optional) Route tables to modify.`,
				},
				resource.Attribute{
					Name:        "remote_gateway_latitude",
					Description: `(Optional) Latitude of remote gateway. Does not support refresh.`,
				},
				resource.Attribute{
					Name:        "remote_gateway_longitude",
					Description: `(Optional) Longitude of remote gateway. Does not support refresh.`,
				},
				resource.Attribute{
					Name:        "backup_remote_gateway_latitude",
					Description: `(Optional) Latitude of backup remote gateway. Does not support refresh.`,
				},
				resource.Attribute{
					Name:        "backup_remote_gateway_longitude",
					Description: `(Optional) Longitude of backup remote gateway. Does not support refresh. ### Misc.`,
				},
				resource.Attribute{
					Name:        "pre_shared_key",
					Description: `(Optional) Pre-Shared Key. Only available for "udp" tunnel_type.`,
				},
				resource.Attribute{
					Name:        "ssl_server_pool",
					Description: `(Optional) Specify ssl_server_pool for tunnel_type "tcp". Default value: "192.168.44.0/24".`,
				},
				resource.Attribute{
					Name:        "enable_dead_peer_detection",
					Description: `(Optional) Enable/disable Deed Peer Detection for an existing site2cloud connection. Default value: true.`,
				},
				resource.Attribute{
					Name:        "enable_active_active",
					Description: `(Optional) Enable/disable active active HA for an existing site2cloud connection. Valid values: true, false. Default value: false. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "local_subnet_cidr",
					Description: `Local subnet CIDR. ## Import Instance site2cloud can be imported using the connection_name and vpc_id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import aviatrix_site2cloud.test connection_name~vpc_id ` + "`" + `` + "`" + `` + "`" + ` ## Notes ### custom_algorithms Only supported for 'udp' tunnel type. If set to true, the six algorithm arguments cannot all be default value. If set to false, default values will be used for all six algorithm arguments. ### enable_dead_peer_detection If you are using/upgraded to Aviatrix Terraform Provider R1.9+, and a site2cloud resource was originally created with a provider version <R1.9, you must do ‘terraform refresh’ to update and apply the attribute’s default value (true) into the state file. ### HA Enabled The following arguments are only supported if the backup gateway is set up by enabling peering HA through the primary gateway resource by specifying a "peering_ha_subnet" and "peering_ha_gw_size". For more information on site2cloud, please see the doc site [here](https://docs.aviatrix.com/HowTos/site2cloud.html):`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "local_subnet_cidr",
					Description: `Local subnet CIDR. ## Import Instance site2cloud can be imported using the connection_name and vpc_id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import aviatrix_site2cloud.test connection_name~vpc_id ` + "`" + `` + "`" + `` + "`" + ` ## Notes ### custom_algorithms Only supported for 'udp' tunnel type. If set to true, the six algorithm arguments cannot all be default value. If set to false, default values will be used for all six algorithm arguments. ### enable_dead_peer_detection If you are using/upgraded to Aviatrix Terraform Provider R1.9+, and a site2cloud resource was originally created with a provider version <R1.9, you must do ‘terraform refresh’ to update and apply the attribute’s default value (true) into the state file. ### HA Enabled The following arguments are only supported if the backup gateway is set up by enabling peering HA through the primary gateway resource by specifying a "peering_ha_subnet" and "peering_ha_gw_size". For more information on site2cloud, please see the doc site [here](https://docs.aviatrix.com/HowTos/site2cloud.html):`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_spoke_gateway",
			Category:         "Transit Network",
			ShortDescription: `Creates and manages Aviatrix spoke gateways`,
			Description:      ``,
			Keywords: []string{
				"transit",
				"network",
				"spoke",
				"gateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_type",
					Description: `(Required) Type of cloud service provider, requires an integer value. Currently only AWS(1), GCP(4), ARM(8), and OCI(16) are supported.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `(Required) This parameter represents the name of a Cloud-Account in Aviatrix controller.`,
				},
				resource.Attribute{
					Name:        "gw_name",
					Description: `(Required) Name of the gateway which is going to be created.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) VPC-ID/VNet-Name of cloud provider. Example: AWS: "vpc-abcd1234", GCP: "vpc-gcp-test", ARM: "vnet1:hello", OCI: "vpc-oracle-test1".`,
				},
				resource.Attribute{
					Name:        "vpc_reg",
					Description: `(Required) Region of cloud provider. Example: AWS: "us-east-1", GCP: "us-west2-a", ARM: "East US 2", Oracle: "us-ashburn-1".`,
				},
				resource.Attribute{
					Name:        "gw_size",
					Description: `(Required) Size of the gateway instance. Example: AWS: "t2.large", ARM: "Standard_B1s", Oracle: "VM.Standard2.2", GCP: "n1-standard-1".`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `(Required) A VPC Network address range selected from one of the available network ranges. Example: "172.31.0.0/20".`,
				},
				resource.Attribute{
					Name:        "ha_subnet",
					Description: `(Optional) HA Subnet. Required only if enabling HA for AWS/ARM gateway. Setting to empty/unsetting will disable HA. Setting to a valid subnet CIDR will create an HA gateway on the subnet. Example: "10.12.0.0/24"`,
				},
				resource.Attribute{
					Name:        "ha_zone",
					Description: `(Optional) HA Zone. Required only if enabling HA for GCP gateway. Setting to empty/unsetting will disable HA. Setting to a valid zone will create an HA gateway in the zone. Example: "us-west1-c".`,
				},
				resource.Attribute{
					Name:        "ha_eip",
					Description: `(Optional) Public IP address that you want to assign to the HA peering instance. If no value is given, a new EIP will automatically be allocated. Only available for AWS.`,
				},
				resource.Attribute{
					Name:        "ha_gw_size",
					Description: `(Optional) HA Gateway Size. Mandatory if enabling HA. Example: "t2.micro". ### Insane Mode`,
				},
				resource.Attribute{
					Name:        "insane_mode",
					Description: `(Optional) Enable Insane Mode for Spoke Gateway. Insane Mode gateway size has to be at least c5 (AWS) or Standard_D3_v2 (ARM). If enabled, you must specify a valid /26 CIDR segment of the VPC to create a new subnet. Only supported for AWS and ARM. Valid values: true, false.`,
				},
				resource.Attribute{
					Name:        "insane_mode_az",
					Description: `(Optional) AZ of subnet being created for Insane Mode Spoke Gateway. Required for AWS if ` + "`" + `insane_mode` + "`" + ` is enabled. Example: AWS: "us-west-1a". ### SNAT/DNAT`,
				},
				resource.Attribute{
					Name:        "single_ip_snat",
					Description: `(Optional) Specify whether to enable Source NAT feature in "single_ip" mode on the gateway or not. Please disable AWS NAT instance before enabling this feature. Currently only supports AWS(1) and ARM(8). Valid values: true, false. ->`,
				},
				resource.Attribute{
					Name:        "transit_gw",
					Description: `(Optional) Specify the transit Gateway.`,
				},
				resource.Attribute{
					Name:        "allocate_new_eip",
					Description: `(Optional) When value is false, reuse an idle address in Elastic IP pool for this gateway. Otherwise, allocate a new Elastic IP and use it for this gateway. Available in Controller 4.7+. Valid values: true, false. Default: true. Option not available for GCP, ARM and OCI gateways, they will automatically allocate new EIPs.`,
				},
				resource.Attribute{
					Name:        "eip",
					Description: `(Optional) Required when ` + "`" + `allocate_new_eip` + "`" + ` is false. It uses the specified EIP for this gateway. Available in Controller 4.7+. Only available for AWS.`,
				},
				resource.Attribute{
					Name:        "tag_list",
					Description: `(Optional) Instance tag of cloud provider. Only AWS, cloud_type is "1", is supported. Example: ["key1:value1", "key2:value2"].`,
				},
				resource.Attribute{
					Name:        "enable_active_mesh",
					Description: `(Optional) Switch to enable/disable Active Mesh Mode for Spoke Gateway. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "enable_vpc_dns_server",
					Description: `(Optional) Enable VPC DNS Server for Gateway. Currently only supports AWS. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "enable_encrypt_volume",
					Description: `(Optional) Enable EBS volume encryption for Gateway. Only supports AWS. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "customer_managed_keys",
					Description: `(Optional and Sensitive) Customer managed key ID.`,
				},
				resource.Attribute{
					Name:        "customized_spoke_vpc_routes",
					Description: `(Optional) A list of comma separated CIDRs to be customized for the spoke VPC routes. When configured, it will replace all learned routes in VPC routing tables, including RFC1918 and non-RFC1918 CIDRs. It applies to this spoke gateway only​. Example: "10.0.0.0/116,10.2.0.0/16".`,
				},
				resource.Attribute{
					Name:        "filtered_spoke_vpc_routes",
					Description: `(Optional) A list of comma separated CIDRs to be filtered from the spoke VPC route table. When configured, filtering CIDR(s) or it’s subnet will be deleted from VPC routing tables as well as from spoke gateway’s routing table. It applies to this spoke gateway only. Example: "10.2.0.0/116,10.3.0.0/16".`,
				},
				resource.Attribute{
					Name:        "included_advertised_spoke_routes",
					Description: `(Optional) A list of comma separated CIDRs to be advertised to on-prem as 'Included CIDR List'. When configured, it will replace all advertised routes from this VPC. Example: "10.4.0.0/116,10.5.0.0/16". ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "eip",
					Description: `Public IP address assigned to the gateway.`,
				},
				resource.Attribute{
					Name:        "ha_eip",
					Description: `Public IP address assigned to the HA gateway.`,
				},
				resource.Attribute{
					Name:        "cloud_instance_id",
					Description: `Cloud Instance ID. The following arguments are deprecated:`,
				},
				resource.Attribute{
					Name:        "enable_snat",
					Description: `(Optional) Specify whether enabling Source NAT feature on the gateway or not. Please disable AWS NAT instance before enabling this feature. Currently only supports AWS(1) and ARM(8). Valid values: true, false.`,
				},
				resource.Attribute{
					Name:        "snat_mode",
					Description: `(Optional) Valid values: "primary", "secondary" and "custom". Default value: "primary".`,
				},
				resource.Attribute{
					Name:        "snat_policy",
					Description: `(Optional) Policy rule applied for "snat_mode" of "custom".`,
				},
				resource.Attribute{
					Name:        "src_ip",
					Description: `(Optional) A source IP address range where the policy rule applies.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `(Optional) A source port that the policy rule applies.`,
				},
				resource.Attribute{
					Name:        "dst_ip",
					Description: `(Optional) A destination IP address range where the policy rule applies.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `(Optional) A destination port where the policy rule applies.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) A destination port protocol where the policy rule applies.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(Optional) An output interface where the policy rule applies.`,
				},
				resource.Attribute{
					Name:        "connection",
					Description: `(Optional) Default value: "None".`,
				},
				resource.Attribute{
					Name:        "mark",
					Description: `(Optional) A tag or mark of a TCP session where the policy rule applies.`,
				},
				resource.Attribute{
					Name:        "new_src_ip",
					Description: `(Optional) The changed source IP address when all specified qualifier conditions meet. One of the rule fields must be specified for this rule to take effect.`,
				},
				resource.Attribute{
					Name:        "new_src_port",
					Description: `(Optional) The translated destination port when all specified qualifier conditions meet. One of the rule field must be specified for this rule to take effect.`,
				},
				resource.Attribute{
					Name:        "exclude_rtb",
					Description: `(Optional) This field specifies which VPC private route table will not be programmed with the default route entry.`,
				},
				resource.Attribute{
					Name:        "dnat_policy",
					Description: `(Optional) Policy rule applied for enabling Destination NAT (DNAT), which allows you to change the destination to a virtual address range. Currently only supports AWS(1) and ARM(8).`,
				},
				resource.Attribute{
					Name:        "src_ip",
					Description: `(Optional) A source IP address range where the policy rule applies.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `(Optional) A source port that the policy rule applies.`,
				},
				resource.Attribute{
					Name:        "dst_ip",
					Description: `(Optional) A destination IP address range where the policy rule applies.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `(Optional) A destination port where the policy rule applies.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) A destination port protocol where the policy rule applies.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(Optional) An output interface where the policy rule applies.`,
				},
				resource.Attribute{
					Name:        "connection",
					Description: `(Optional) Default value: "None".`,
				},
				resource.Attribute{
					Name:        "mark",
					Description: `(Optional) A tag or mark of a TCP session where the policy rule applies.`,
				},
				resource.Attribute{
					Name:        "new_src_ip",
					Description: `(Optional) The changed source IP address when all specified qualifier conditions meet. One of the rule fields must be specified for this rule to take effect.`,
				},
				resource.Attribute{
					Name:        "new_src_port",
					Description: `(Optional) The translated destination port when all specified qualifier conditions meet. One of the rule field must be specified for this rule to take effect.`,
				},
				resource.Attribute{
					Name:        "exclude_rtb",
					Description: `(Optional) This field specifies which VPC private route table will not be programmed with the default route entry. ## Import Instance spoke_gateway can be imported using the gw_name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import aviatrix_spoke_gateway.test gw_name ` + "`" + `` + "`" + `` + "`" + ` ## Notes ### insane_mode If ` + "`" + `insane_mode` + "`" + ` is enabled, you must specify a valid /26 CIDR segment of the VPC specified for the ` + "`" + `subnet` + "`" + `. This will then create a new subnet to be used for the corresponding gateway. You cannot specify an existing /26 subnet. ### enable_snat If you are using/upgraded to Aviatrix Terraform Provider R2.10+, and a spoke gateway with ` + "`" + `enable_snat` + "`" + ` set to true was originally created with a provider version <R2.10, you must do a ‘terraform refresh’ to update and apply the attribute’s value into the state. In addition, you must also change this attribute to ` + "`" + `single_ip_snat` + "`" + ` in your ` + "`" + `.tf` + "`" + ` file. ### snat_mode & snat_policy If you are using/upgraded to Aviatrix Terraform Provider R2.10+, and a spoke gateway with ` + "`" + `snat_mode` + "`" + ` and ` + "`" + `snat_policy` + "`" + ` was originally created with a provider version <R2.10, you must do a ‘terraform refresh’ to remove attribute’s value from the state. In addition, you must transfer its corresponding values to the`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "eip",
					Description: `Public IP address assigned to the gateway.`,
				},
				resource.Attribute{
					Name:        "ha_eip",
					Description: `Public IP address assigned to the HA gateway.`,
				},
				resource.Attribute{
					Name:        "cloud_instance_id",
					Description: `Cloud Instance ID. The following arguments are deprecated:`,
				},
				resource.Attribute{
					Name:        "enable_snat",
					Description: `(Optional) Specify whether enabling Source NAT feature on the gateway or not. Please disable AWS NAT instance before enabling this feature. Currently only supports AWS(1) and ARM(8). Valid values: true, false.`,
				},
				resource.Attribute{
					Name:        "snat_mode",
					Description: `(Optional) Valid values: "primary", "secondary" and "custom". Default value: "primary".`,
				},
				resource.Attribute{
					Name:        "snat_policy",
					Description: `(Optional) Policy rule applied for "snat_mode" of "custom".`,
				},
				resource.Attribute{
					Name:        "src_ip",
					Description: `(Optional) A source IP address range where the policy rule applies.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `(Optional) A source port that the policy rule applies.`,
				},
				resource.Attribute{
					Name:        "dst_ip",
					Description: `(Optional) A destination IP address range where the policy rule applies.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `(Optional) A destination port where the policy rule applies.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) A destination port protocol where the policy rule applies.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(Optional) An output interface where the policy rule applies.`,
				},
				resource.Attribute{
					Name:        "connection",
					Description: `(Optional) Default value: "None".`,
				},
				resource.Attribute{
					Name:        "mark",
					Description: `(Optional) A tag or mark of a TCP session where the policy rule applies.`,
				},
				resource.Attribute{
					Name:        "new_src_ip",
					Description: `(Optional) The changed source IP address when all specified qualifier conditions meet. One of the rule fields must be specified for this rule to take effect.`,
				},
				resource.Attribute{
					Name:        "new_src_port",
					Description: `(Optional) The translated destination port when all specified qualifier conditions meet. One of the rule field must be specified for this rule to take effect.`,
				},
				resource.Attribute{
					Name:        "exclude_rtb",
					Description: `(Optional) This field specifies which VPC private route table will not be programmed with the default route entry.`,
				},
				resource.Attribute{
					Name:        "dnat_policy",
					Description: `(Optional) Policy rule applied for enabling Destination NAT (DNAT), which allows you to change the destination to a virtual address range. Currently only supports AWS(1) and ARM(8).`,
				},
				resource.Attribute{
					Name:        "src_ip",
					Description: `(Optional) A source IP address range where the policy rule applies.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `(Optional) A source port that the policy rule applies.`,
				},
				resource.Attribute{
					Name:        "dst_ip",
					Description: `(Optional) A destination IP address range where the policy rule applies.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `(Optional) A destination port where the policy rule applies.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) A destination port protocol where the policy rule applies.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(Optional) An output interface where the policy rule applies.`,
				},
				resource.Attribute{
					Name:        "connection",
					Description: `(Optional) Default value: "None".`,
				},
				resource.Attribute{
					Name:        "mark",
					Description: `(Optional) A tag or mark of a TCP session where the policy rule applies.`,
				},
				resource.Attribute{
					Name:        "new_src_ip",
					Description: `(Optional) The changed source IP address when all specified qualifier conditions meet. One of the rule fields must be specified for this rule to take effect.`,
				},
				resource.Attribute{
					Name:        "new_src_port",
					Description: `(Optional) The translated destination port when all specified qualifier conditions meet. One of the rule field must be specified for this rule to take effect.`,
				},
				resource.Attribute{
					Name:        "exclude_rtb",
					Description: `(Optional) This field specifies which VPC private route table will not be programmed with the default route entry. ## Import Instance spoke_gateway can be imported using the gw_name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import aviatrix_spoke_gateway.test gw_name ` + "`" + `` + "`" + `` + "`" + ` ## Notes ### insane_mode If ` + "`" + `insane_mode` + "`" + ` is enabled, you must specify a valid /26 CIDR segment of the VPC specified for the ` + "`" + `subnet` + "`" + `. This will then create a new subnet to be used for the corresponding gateway. You cannot specify an existing /26 subnet. ### enable_snat If you are using/upgraded to Aviatrix Terraform Provider R2.10+, and a spoke gateway with ` + "`" + `enable_snat` + "`" + ` set to true was originally created with a provider version <R2.10, you must do a ‘terraform refresh’ to update and apply the attribute’s value into the state. In addition, you must also change this attribute to ` + "`" + `single_ip_snat` + "`" + ` in your ` + "`" + `.tf` + "`" + ` file. ### snat_mode & snat_policy If you are using/upgraded to Aviatrix Terraform Provider R2.10+, and a spoke gateway with ` + "`" + `snat_mode` + "`" + ` and ` + "`" + `snat_policy` + "`" + ` was originally created with a provider version <R2.10, you must do a ‘terraform refresh’ to remove attribute’s value from the state. In addition, you must transfer its corresponding values to the`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_spoke_vpc",
			Category:         "Transit Network",
			ShortDescription: `Creates and Manages Aviatrix Spoke Gateways`,
			Description:      ``,
			Keywords: []string{
				"transit",
				"network",
				"spoke",
				"vpc",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_type",
					Description: `(Required) Type of cloud service provider. AWS=1, GCP=4, ARM=8.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `(Required) This parameter represents the name of a Cloud-Account in Aviatrix controller.`,
				},
				resource.Attribute{
					Name:        "gw_name",
					Description: `(Required) Name of the gateway which is going to be created.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) VPC-ID/VNet-Name of cloud provider. Required if cloud_type is "1" or "4". Example: AWS: "vpc-abcd1234", etc...`,
				},
				resource.Attribute{
					Name:        "vpc_reg",
					Description: `(Required) Region of cloud provider. Example: AWS: "us-east-1", GCP: "us-west1-b", ARM: "East US 2", etc...`,
				},
				resource.Attribute{
					Name:        "vpc_size",
					Description: `(Required) Size of the gateway instance. Example: AWS: "t2.large", GCP: "f1.micro", ARM: "StandardD2", etc...`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `(Required) Public Subnet Info. Example: AWS: "CIDRZONESubnetName", etc...`,
				},
				resource.Attribute{
					Name:        "ha_subnet",
					Description: `(Optional) HA Subnet. Required for enabling HA for AWS/ARM gateways. Setting to empty/unset will disable HA. Setting to a valid subnet (Example: 10.12.0.0/24) will create an HA gateway on the subnet.`,
				},
				resource.Attribute{
					Name:        "ha_zone",
					Description: `(Optional) HA Zone. Required for enabling HA for GCP gateway. Setting to empty/unset will disable HA. Setting to a valid zone will create an HA gateway in the zone. Example: "us-west1-c".`,
				},
				resource.Attribute{
					Name:        "ha_gw_size",
					Description: `(Optional) HA Gateway Size. Mandatory if HA is enabled (ha_subnet is set). Example: "t2.micro".`,
				},
				resource.Attribute{
					Name:        "enable_snat",
					Description: `(Optional) Enable Source NAT for this container. Supported values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "single_az_ha",
					Description: `(Optional) Set to "enabled" if this feature is desired.`,
				},
				resource.Attribute{
					Name:        "transit_gw",
					Description: `(Optional) Specify the transit Gateway.`,
				},
				resource.Attribute{
					Name:        "tag_list",
					Description: `(Optional) Instance tag of cloud provider. Example: key1:value1,key002:value002, etc... Only AWS (cloud_type is "1") is supported The following arguments are deprecated:`,
				},
				resource.Attribute{
					Name:        "dns_server",
					Description: `Specify the DNS IP, only required while using a custom private DNS for the VPC. ->`,
				},
				resource.Attribute{
					Name:        "vnet_and_resource_group_names",
					Description: `If you are using/upgraded to Aviatrix Terraform Provider R1.10+, and an ARM spoke_vpc resource was originally created with a provider version < R1.10, you must replace "vnet_and_resource_group_names" with "vpc_id" in your configuration file, and do ‘terraform refresh’ to set its value to "vpc_id" and apply it into the state file. ## Import Instance spoke_vpc can be imported using the gw_name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import aviatrix_spoke_vpc.test gw_name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_trans_peer",
			Category:         "Peering",
			ShortDescription: `Creates and manages Aviatrix transitive peerings`,
			Description:      ``,
			Keywords: []string{
				"peering",
				"trans",
				"peer",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "source",
					Description: `(Required) Name of Source gateway.`,
				},
				resource.Attribute{
					Name:        "nexthop",
					Description: `(Required) Name of nexthop gateway.`,
				},
				resource.Attribute{
					Name:        "reachable_cidr",
					Description: `(Required) Destination CIDR. ## Import Instance trans_peer can be imported using the source, nexthop and reachable_cidr, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import aviatrix_trans_peer.test source~nexthop~reachable_cidr ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_transit_gateway",
			Category:         "Transit Network",
			ShortDescription: `Creates and manages the Aviatrix transit network gateways`,
			Description:      ``,
			Keywords: []string{
				"transit",
				"network",
				"gateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_type",
					Description: `(Required) Type of cloud service provider, requires an integer value. Currently only AWS(1), GCP(4), ARM(8), and OCI(16) are supported.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `(Required) This parameter represents the name of a Cloud-Account in Aviatrix controller.`,
				},
				resource.Attribute{
					Name:        "gw_name",
					Description: `(Required) Name of the gateway which is going to be created.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) VPC-ID/VNet-Name of cloud provider. Example: AWS: "vpc-abcd1234", GCP: "vpc-gcp-test", ARM: "vnet1:hello", OCI: "vpc-oracle-test1".`,
				},
				resource.Attribute{
					Name:        "vpc_reg",
					Description: `(Required) Region of cloud provider. Example: AWS: "us-east-1", GCP: "us-west2-a", ARM: "East US 2", Oracle: "us-ashburn-1".`,
				},
				resource.Attribute{
					Name:        "gw_size",
					Description: `(Required) Size of the gateway instance. Example: AWS: "t2.large", ARM: "Standard_B1s", Oracle: "VM.Standard2.2", GCP: "n1-standard-1".`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `(Required) A VPC Network address range selected from one of the available network ranges. Example: "172.31.0.0/20".`,
				},
				resource.Attribute{
					Name:        "ha_subnet",
					Description: `(Optional) HA Subnet CIDR. Required only if enabling HA for AWS/ARM gateway. Setting to empty/unsetting will disable HA. Setting to a valid subnet CIDR will create an HA gateway on the subnet. Example: "10.12.0.0/24".`,
				},
				resource.Attribute{
					Name:        "ha_zone",
					Description: `(Optional) HA Zone. Required only if enabling HA for GCP gateway. Setting to empty/unsetting will disable HA. Setting to a valid zone will create an HA gateway in the zone. Example: "us-west1-c".`,
				},
				resource.Attribute{
					Name:        "ha_insane_mode_az",
					Description: `(Optional) AZ of subnet being created for Insane Mode Transit HA Gateway. Required for AWS if ` + "`" + `insane_mode` + "`" + ` is enabled and ` + "`" + `ha_subnet` + "`" + ` is set. Example: AWS: "us-west-1a".`,
				},
				resource.Attribute{
					Name:        "ha_eip",
					Description: `(Optional) Public IP address that you want to assign to the HA peering instance. If no value is given, a new EIP will automatically be allocated. Only available for AWS.`,
				},
				resource.Attribute{
					Name:        "ha_gw_size",
					Description: `(Optional) HA Gateway Size. Mandatory if enabling HA. Example: "t2.micro". ### Insane Mode`,
				},
				resource.Attribute{
					Name:        "insane_mode",
					Description: `(Optional) Specify Insane Mode high performance gateway. Insane Mode gateway size must be at least c5 size (AWS) or Standard_D3_v2 (ARM). If enabled, you must specify a valid /26 CIDR segment of the VPC to create a new subnet. Only available for AWS and ARM. Valid values: true, false.`,
				},
				resource.Attribute{
					Name:        "insane_mode_az",
					Description: `(Optional) AZ of subnet being created for Insane Mode Transit Gateway. Required for AWS if ` + "`" + `insane_mode` + "`" + ` is enabled. Example: AWS: "us-west-1a". ### SNAT`,
				},
				resource.Attribute{
					Name:        "single_ip_snat",
					Description: `(Optional) Enable "single_ip" mode Source NAT for this container. Valid values: true, false.`,
				},
				resource.Attribute{
					Name:        "allocate_new_eip",
					Description: `(Optional) When value is false, reuse an idle address in Elastic IP pool for this gateway. Otherwise, allocate a new Elastic IP and use it for this gateway. Available in Controller 4.7+. Valid values: true, false. Default: true. Option not available for GCP, ARM and OCI gateways, they will automatically allocate new EIPs.`,
				},
				resource.Attribute{
					Name:        "eip",
					Description: `(Optional) Required when ` + "`" + `allocate_new_eip` + "`" + ` is false. It uses the specified EIP for this gateway. Available in Controller version 4.7+. Only available for AWS.`,
				},
				resource.Attribute{
					Name:        "tag_list",
					Description: `(Optional) Instance tag of cloud provider. Only supported for AWS. Example: ["key1:value1","key2:value2"].`,
				},
				resource.Attribute{
					Name:        "connected_transit",
					Description: `(Optional) Specify Connected Transit status. If enabled, it allows spokes to run traffics to other spokes via transit gateway. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "enable_hybrid_connection",
					Description: `(Optional) Sign of readiness for TGW connection. Only supported for AWS. Example: false.`,
				},
				resource.Attribute{
					Name:        "enable_firenet",
					Description: `(Optional) Sign of readiness for FireNet connection. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "enable_advertise_transit_cidr",
					Description: `(Optional) Switch to enable/disable advertise transit VPC network CIDR for a vgw connection. Available as of R2.6.`,
				},
				resource.Attribute{
					Name:        "bgp_manual_spoke_advertise_cidrs",
					Description: `(Optional) Intended CIDR list to advertise to VGW. Example: "10.2.0.0/16,10.4.0.0/16". Available as of R2.6.`,
				},
				resource.Attribute{
					Name:        "enable_active_mesh",
					Description: `(Optional) Switch to enable/disable Active Mesh Mode for Transit Gateway. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "enable_vpc_dns_server",
					Description: `(Optional) Enable VPC DNS Server for Gateway. Currently only supports AWS. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "enable_encrypt_volume",
					Description: `(Optional) Enable EBS volume encryption for Gateway. Only supports AWS. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "customer_managed_keys",
					Description: `(Optional and Sensitive) Customer managed key ID.`,
				},
				resource.Attribute{
					Name:        "customized_spoke_vpc_routes",
					Description: `(Optional) A list of comma separated CIDRs to be customized for the spoke VPC routes. When configured, it will replace all learned routes in VPC routing tables, including RFC1918 and non-RFC1918 CIDRs. It applies to all spoke gateways attached to this transit gateway. Example: "10.0.0.0/116,10.2.0.0/16".`,
				},
				resource.Attribute{
					Name:        "filtered_spoke_vpc_routes",
					Description: `(Optional) A list of comma separated CIDRs to be filtered from the spoke VPC route table. When configured, filtering CIDR(s) or it’s subnet will be deleted from VPC routing tables as well as from spoke gateway’s routing table. It applies to all spoke gateways attached to this transit gateway. Example: "10.2.0.0/116,10.3.0.0/16".`,
				},
				resource.Attribute{
					Name:        "excluded_advertised_spoke_routes",
					Description: `(Optional) A list of comma separated CIDRs to be advertised to on-prem as 'Excluded CIDR List'. When configured, it inspects all the advertised CIDRs from its spoke gateways and remove those included in the 'Excluded CIDR List'. Example: "10.4.0.0/116,10.5.0.0/16". ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "eip",
					Description: `Public IP address assigned to the gateway.`,
				},
				resource.Attribute{
					Name:        "ha_eip",
					Description: `Public IP address assigned to the HA gateway. The following arguments are deprecated:`,
				},
				resource.Attribute{
					Name:        "enable_firenet_interfaces",
					Description: `(Optional) Sign of readiness for FireNet connection. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "enable_snat",
					Description: `(Optional) Enable Source NAT for this container. Valid values: true, false. ## Import Instance transit_gateway can be imported using the gw_name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import aviatrix_transit_gateway.test gw_name ` + "`" + `` + "`" + `` + "`" + ` ## Notes ### CIDR advertising ` + "`" + `enable_advertise_transit_cidr` + "`" + ` and ` + "`" + `bgp_manual_spoke_advertise_cidrs` + "`" + ` functionality has been migrated over to`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "eip",
					Description: `Public IP address assigned to the gateway.`,
				},
				resource.Attribute{
					Name:        "ha_eip",
					Description: `Public IP address assigned to the HA gateway. The following arguments are deprecated:`,
				},
				resource.Attribute{
					Name:        "enable_firenet_interfaces",
					Description: `(Optional) Sign of readiness for FireNet connection. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "enable_snat",
					Description: `(Optional) Enable Source NAT for this container. Valid values: true, false. ## Import Instance transit_gateway can be imported using the gw_name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import aviatrix_transit_gateway.test gw_name ` + "`" + `` + "`" + `` + "`" + ` ## Notes ### CIDR advertising ` + "`" + `enable_advertise_transit_cidr` + "`" + ` and ` + "`" + `bgp_manual_spoke_advertise_cidrs` + "`" + ` functionality has been migrated over to`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_transit_gateway_peering",
			Category:         "Transit Network",
			ShortDescription: `Creates and manages Aviatrix transit gateway peerings`,
			Description:      ``,
			Keywords: []string{
				"transit",
				"network",
				"gateway",
				"peering",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "transit_gateway_name1",
					Description: `(Required) The first transit gateway name to make a peer pair.`,
				},
				resource.Attribute{
					Name:        "transit_gateway_name2",
					Description: `(Required) The second transit gateway name to make a peer pair. ## Import Instance transit_gateway_peering can be imported using the transit_gateway_name1 and transit_gateway_name2, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import aviatrix_transit_gateway_peering.test transit_gateway_name1~transit_gateway_name2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_transit_vpc",
			Category:         "Transit Network",
			ShortDescription: `Creates and Manages the Aviatrix Transit Network Gateways`,
			Description:      ``,
			Keywords: []string{
				"transit",
				"network",
				"vpc",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_type",
					Description: `(Required) Type of cloud service provider, requires an integer value. Use 1 for AWS.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `(Required) This parameter represents the name of a Cloud-Account in Aviatrix controller.`,
				},
				resource.Attribute{
					Name:        "gw_name",
					Description: `(Required) Name of the gateway which is going to be created.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) VPC-ID/VNet-Name of cloud provider. Required if for aws. Example: AWS: "vpc-abcd1234", GCP: "mygooglecloudvpcname", etc...`,
				},
				resource.Attribute{
					Name:        "vpc_reg",
					Description: `(Required) Region of cloud provider. Example: AWS: "us-east-1", ARM: "East US 2", etc...`,
				},
				resource.Attribute{
					Name:        "vpc_size",
					Description: `(Required) Size of the gateway instance. Example: AWS: "t2.large", etc...`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `(Required) Public Subnet CIDR. Example: AWS: "10.0.0.0/24". Copy/paste from AWS Console to get the right subnet CIDR.`,
				},
				resource.Attribute{
					Name:        "ha_subnet",
					Description: `(Optional) HA Subnet CIDR. Example: "10.12.0.0/24".Setting to empty/unset will disable HA. Setting to a valid subnet CIDR will create an HA gateway on the subnet.`,
				},
				resource.Attribute{
					Name:        "ha_gw_size",
					Description: `(Optional) HA Gateway Size. Mandatory if HA is enabled (ha_subnet is set). Example: "t2.micro".`,
				},
				resource.Attribute{
					Name:        "enable_snat",
					Description: `(Optional) Enable Source NAT for this container. Supported values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "tag_list",
					Description: `(Optional) Instance tag of cloud provider. Only supported for aws. Example: ["key1:value1","key002:value002"]`,
				},
				resource.Attribute{
					Name:        "enable_hybrid_connection",
					Description: `(Optional) Sign of readiness for TGW connection. Only supported for aws. Example: false.`,
				},
				resource.Attribute{
					Name:        "enable_firenet_interfaces",
					Description: `(Optional) Sign of readiness for FireNet connection. Valid values: true and false. Default: false.`,
				},
				resource.Attribute{
					Name:        "connected_transit",
					Description: `(Optional) Specify Connected Transit status. Supported values: true, false.`,
				},
				resource.Attribute{
					Name:        "insane_mode",
					Description: `(Optional) Specify Insane Mode high performance gateway. Insane Mode gateway size must be at least c5 size. If enabled, will look for spare /26 segment to create a new subnet. Only available for AWS. Supported values: true, false.`,
				},
				resource.Attribute{
					Name:        "insane_mode_az",
					Description: `(Optional) AZ of subnet being created for Insane Mode Transit Gateway. Required if insane_mode is enabled.`,
				},
				resource.Attribute{
					Name:        "ha_insane_mode_az",
					Description: `(Optional) AZ of subnet being created for Insane Mode Transit HA Gateway. Required if insane_mode is enabled and ha_subnet is set. The following arguments are deprecated:`,
				},
				resource.Attribute{
					Name:        "dns_server",
					Description: `Specify the DNS IP, only required while using a custom private DNS for the VPC.`,
				},
				resource.Attribute{
					Name:        "vnet_name_resource_group",
					Description: `(Optional) VPC-ID/VNet-Name of cloud provider. Required if for azure. ARM: "VNet_Name:Resource_Group_Name". It is replaced by "vpc_id". ->`,
				},
				resource.Attribute{
					Name:        "enable_firenet_interfaces",
					Description: `If you are using/upgraded to Aviatrix Terraform Provider R1.8+, and a transit_vpc resource was originally created with a provider version < R1.8, you must do ‘terraform refresh’ to update and apply the attribute’s default value (false) into the state file. ->`,
				},
				resource.Attribute{
					Name:        "vnet_name_resource_group",
					Description: `If you are using/upgraded to Aviatrix Terraform Provider R1.10+, and an ARM transit_vpc resource was originally created with a provider version < R1.10, you must replace "vnet_name_resource_group" with "vpc_id" in your configuration file, and do ‘terraform refresh’ to set its value to "vpc_id" and apply it into the state file. ## Import Instance transit_vpc can be imported using the gw_name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import aviatrix_transit_vpc.test gw_name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_tunnel",
			Category:         "Peering",
			ShortDescription: `Creates and manages Aviatrix Tunnels`,
			Description:      ``,
			Keywords: []string{
				"peering",
				"tunnel",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "gw_name1",
					Description: `(Required) The first VPC Container name to make a peer pair.`,
				},
				resource.Attribute{
					Name:        "gw_name2",
					Description: `(Required) The second VPC Container name to make a peer pair.`,
				},
				resource.Attribute{
					Name:        "enable_ha",
					Description: `(Optional) Whether Peering HA is enabled. Valid values: true, false. Default value: false. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "peering_state",
					Description: `(Computed) Status of the tunnel.`,
				},
				resource.Attribute{
					Name:        "peering_hastatus",
					Description: `(Computed) Status of the HA tunnel.`,
				},
				resource.Attribute{
					Name:        "peering_link",
					Description: `(Computed) Name of the peering link. ## Import Instance tunnel can be imported using the gw_name1 and gw_name2, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import aviatrix_tunnel.test gw_name1~gw_name2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "peering_state",
					Description: `(Computed) Status of the tunnel.`,
				},
				resource.Attribute{
					Name:        "peering_hastatus",
					Description: `(Computed) Status of the HA tunnel.`,
				},
				resource.Attribute{
					Name:        "peering_link",
					Description: `(Computed) Name of the peering link. ## Import Instance tunnel can be imported using the gw_name1 and gw_name2, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import aviatrix_tunnel.test gw_name1~gw_name2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_vgw_conn",
			Category:         "Transit Network",
			ShortDescription: `Manages the Aviatrix Transit Gateway to VGW`,
			Description:      ``,
			Keywords: []string{
				"transit",
				"network",
				"vgw",
				"conn",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "conn_name",
					Description: `(Required) The name of for Transit GW to VGW connection connection which is going to be created. Example: "my-connection-vgw-to-tgw".`,
				},
				resource.Attribute{
					Name:        "gw_name",
					Description: `(Required) Name of the Transit Gateway. Example: "my-transit-gw".`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) VPC-ID where the Transit Gateway is located. Example: AWS: "vpc-abcd1234".`,
				},
				resource.Attribute{
					Name:        "bgp_vgw_id",
					Description: `(Required) Id of AWS's VGW that is used for this connection. Example: "vgw-abcd1234".`,
				},
				resource.Attribute{
					Name:        "bgp_vgw_account",
					Description: `(Required) Account of AWS's VGW that is used for this connection. Example: "dev-account-1".`,
				},
				resource.Attribute{
					Name:        "bgp_vgw_region",
					Description: `(Required) Region of AWS's VGW that is used for this connection. Example: "us-east-1".`,
				},
				resource.Attribute{
					Name:        "bgp_local_as_num",
					Description: `(Required) BGP Local ASN (Autonomous System Number). Integer between 1-65535. Example: "65001". The following arguments are deprecated:`,
				},
				resource.Attribute{
					Name:        "enable_advertise_transit_cidr",
					Description: `(Optional) Switch to enable/disable advertise transit VPC network CIDR for a vgw connection.`,
				},
				resource.Attribute{
					Name:        "bgp_manual_spoke_advertise_cidrs",
					Description: `(Optional) Intended CIDR list to advertise to VGW. Example: "10.2.0.0/16,10.4.0.0/16". ->`,
				},
				resource.Attribute{
					Name:        "enable_advertise_transit_cidr",
					Description: `If you are using/upgraded to Aviatrix Terraform Provider R1.9+, and a vgw_conn resource was originally created with a provider version <R1.9, you must do ‘terraform refresh’ to update and apply the attribute’s default value (false) into the state file. ~>`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_vpc",
			Category:         "Useful Tools",
			ShortDescription: `Creates and manages VPCs`,
			Description:      ``,
			Keywords: []string{
				"useful",
				"tools",
				"vpc",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_type",
					Description: `(Required) Type of cloud service provider, requires an integer value. Currently only AWS(1) and GCP(4) is supported.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `(Required) This parameter represents the name of a Cloud-Account in Aviatrix controller.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the vpc to be created.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Region of cloud provider. Required to be empty for GCP provider, and non-empty for other providers. Example: AWS: "us-east-1", ARM: "East US 2".`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Optional) VPC cidr. Required to be empty for GCP provider, and non-empty for other providers. Example: "10.11.0.0/24".`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `(Optional) List of subnets to be specify for GCP provider. Required to be non-empty for GCP provider, and empty for other providers.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Region of this subnet.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `CIDR block.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of this subnet.`,
				},
				resource.Attribute{
					Name:        "aviatrix_transit_vpc",
					Description: `(Optional) Specify whether it is an Aviatrix Transit VPC. Only supported for AWS provider, required to be false for other providers. Valid values: true, false. Default: false.`,
				},
				resource.Attribute{
					Name:        "aviatrix_firenet_vpc",
					Description: `(Optional) Specify whether it is an Aviatrix FireNet VPC. Only supported for AWS provider, required to be false for other providers. Valid values: true, false. Default: false. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the vpc to be created.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `List of subnet of the VPC to be created.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `CIDR block.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of this subnet.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of this subnet. ->`,
				},
				resource.Attribute{
					Name:        "aviatrix_firenet_vpc",
					Description: `If you are using/ upgraded to Aviatrix Terraform Provider R1.8+, and an vpc resource was originally created with a provider version < R1.8, you must do ‘terraform refresh’ to update and apply the attribute’s default value (false) into the state file. ->`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `If created as a FireNet VPC, four public subnets will be created in the following order: subnet for firewall-mgmt in the first zone, subnet for ingress-egress in the first zone, subnet for firewall-mgmt in the second zone, and subnet for ingress-egress in the second zone. ## Import Instance vpc can be imported using the name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import aviatrix_vpc.test name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the vpc to be created.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `List of subnet of the VPC to be created.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `CIDR block.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of this subnet.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of this subnet. ->`,
				},
				resource.Attribute{
					Name:        "aviatrix_firenet_vpc",
					Description: `If you are using/ upgraded to Aviatrix Terraform Provider R1.8+, and an vpc resource was originally created with a provider version < R1.8, you must do ‘terraform refresh’ to update and apply the attribute’s default value (false) into the state file. ->`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `If created as a FireNet VPC, four public subnets will be created in the following order: subnet for firewall-mgmt in the first zone, subnet for ingress-egress in the first zone, subnet for firewall-mgmt in the second zone, and subnet for ingress-egress in the second zone. ## Import Instance vpc can be imported using the name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import aviatrix_vpc.test name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_vpn_profile",
			Category:         "OpenVPN",
			ShortDescription: `Creates and manages Aviatrix VPN User Profiles`,
			Description:      ``,
			Keywords: []string{
				"openvpn",
				"vpn",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Enter any name for the VPN profile.`,
				},
				resource.Attribute{
					Name:        "base_rule",
					Description: `(Optional) Base policy rule of the profile to be added. Enter "allow_all" or "deny_all", based on whether you want a white list or black list.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `(Optional) List of VPN users to attach to this profile.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Optional) New security policy for the profile. Each policy has the following attributes:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) Should be the opposite of the base rule for correct behaviour. Valid values for action: "allow", "deny".`,
				},
				resource.Attribute{
					Name:        "proto",
					Description: `(Required) Protocol to allow or deny. Valid values for protocol: "all", "tcp", "udp", "icmp", "sctp", "rdp", "dccp".`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) Port to be allowed or denied. Valid values for port: a single port or a range of port numbers e.g.: "25", "25:1024". For "all" and "icmp", port should only be "0:65535".`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Required) CIDR to be allowed or denied. Valid values for target: IPv4 CIDRs. Example: "10.30.0.0/16". ## Import Instance vpn_profile can be imported using the name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import aviatrix_vpn_profile.test name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_vpn_user",
			Category:         "OpenVPN",
			ShortDescription: `Creates and Manages Aviatrix VPN Users`,
			Description:      ``,
			Keywords: []string{
				"openvpn",
				"vpn",
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) VPC ID of Aviatrix VPN gateway. Example: "vpc-abcd1234".`,
				},
				resource.Attribute{
					Name:        "gw_name",
					Description: `(Required) If ELB is enabled, this will be the name of the ELB, else it will be the name of the Aviatrix VPN gateway. Example: "gw1".`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `(Required) VPN user name. Example: "user".`,
				},
				resource.Attribute{
					Name:        "user_email",
					Description: `(Optional) VPN User's email. Example: "abc@xyz.com".`,
				},
				resource.Attribute{
					Name:        "saml_endpoint",
					Description: `(Optional) This is the name of the SAML endpoint to which the user is to be associated. This is required if adding user to a SAML gateway/LB. ## Import Instance vpn_user can be imported using the user_name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import aviatrix_vpn_user.test user_name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_vpn_user_accelerator",
			Category:         "OpenVPN",
			ShortDescription: `Manages the Aviatrix VPN User Accelerator`,
			Description:      ``,
			Keywords: []string{
				"openvpn",
				"vpn",
				"user",
				"accelerator",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "elb_name",
					Description: `(Required) Name of ELB to be added to VPN User Accelerator. Example: "Aviatrix-vpc-abcd2134". ## Import ` + "`" + `` + "`" + `` + "`" + ` $ terraform import aviatrix_vpn_user_acclerator.test Aviatrix-vpc-abcd1234 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"aviatrix_account":                 0,
		"aviatrix_account_user":            1,
		"aviatrix_arm_peer":                2,
		"aviatrix_aws_peer":                3,
		"aviatrix_aws_tgw":                 4,
		"aviatrix_aws_tgw_directconnect":   5,
		"aviatrix_aws_tgw_vpc_attachment":  6,
		"aviatrix_aws_tgw_vpn_conn":        7,
		"aviatrix_controller_config":       8,
		"aviatrix_firenet":                 9,
		"aviatrix_firewall":                10,
		"aviatrix_firewall_instance":       11,
		"aviatrix_firewall_tag":            12,
		"aviatrix_fqdn":                    13,
		"aviatrix_gateway":                 14,
		"aviatrix_gateway_dnat":            15,
		"aviatrix_gateway_snat":            16,
		"aviatrix_geo_vpn":                 17,
		"aviatrix_saml_endpoint":           18,
		"aviatrix_site2cloud":              19,
		"aviatrix_spoke_gateway":           20,
		"aviatrix_spoke_vpc":               21,
		"aviatrix_trans_peer":              22,
		"aviatrix_transit_gateway":         23,
		"aviatrix_transit_gateway_peering": 24,
		"aviatrix_transit_vpc":             25,
		"aviatrix_tunnel":                  26,
		"aviatrix_vgw_conn":                27,
		"aviatrix_vpc":                     28,
		"aviatrix_vpn_profile":             29,
		"aviatrix_vpn_user":                30,
		"aviatrix_vpn_user_accelerator":    31,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
